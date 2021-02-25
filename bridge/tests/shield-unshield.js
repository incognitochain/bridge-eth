const { expect } = require("chai");
const hre = require("hardhat");
const { deployments, ethers } = hre;

// const cfg = require('../hardhat.config');
const {tokenAddresses} = require('../scripts/constants');
const {BN, getPartOf, confirm, fromIncDecimals, toIncDecimals, getInstance, getImplementation, Inc} = require('../scripts/external');
const {proveEth, formatBurnProof, db} = require('../scripts/prove');
const {getChi, chiify, mintChi} = require('../scripts/chi');
const path = require('path');

describe("Vault (Proxy) - Unshield with CHI", () => {
    let signers, accounts, vaultAdmin, ethUser, tokenUser, chiFunder, ethGuy, tokenGuy;
    const startEth = ethers.utils.parseUnits('0.00005', 'ether');
    const startToken = ethers.utils.parseUnits('0.00001', 'ether');
    const nConfirm = 15;
    let inc, incBridge, transactor;
    let vault, proxy, implementation, upgradedImpl, upgradedVault, testingTokens, chi; // contract instances
    // let network = 'development';

    before(async () => {
        console.log("Network", await ethers.provider.getNetwork());
        chi = await getChi();
        console.log("CHI Token contract", chi.address);
        // console.log(await chi.balanceOf(chiFunder.address));
        signers = await ethers.getSigners();
        accounts = signers.map(s => s.address);
        console.log('Signers :', accounts);
        // 0 is deployer, we use 1 2 3 5 for named accounts, skipping 4
        [_, vaultAdmin, ethUser, tokenUser, _, chiFunder] = signers;

        ethGuy = {signer: ethUser, inc: '12sxXUjkMJZHz6diDB6yYnSjyYcDYiT5QygUYFsUbGUqK8PH8uhxf4LePiAE8UYoDcNkHAdJJtT1J6T8hcvpZoWLHAp8g6h1BQEfp4h5LQgEPuhMpnVMquvr1xXZZueLhTNCXc8fkVXseeVAGCt8', incPrivate: '112t8rnXoBXrThDTACHx2rbEq7nBgrzcZhVZV4fvNEcGJetQ13spZRMuW5ncvsKA1KvtkauZuK2jV8pxEZLpiuHtKX3FkKv2uC5ZeRC8L6we'};
        tokenGuy = {signer: tokenUser, inc: '12sttFKciCWyRbNsK1yD1mWEwZoeWi1JtWJZ7gKTbx5eB25U4FnrfkxgxbnZ8zDn2QNhhW44HBZJ1EnfwVBueR44D5ucWxGNpXZMawoCmv6G2cwKi4xkasuysu3WtpV5ZMSYgaJ1mwe9fqgVD9mh', incPrivate: '112t8rnZUQXxcbayAZvyyZyKDhwVJBLkHuTKMhrS51nQZcXKYXGopUTj22JtZ8KxYQcak54KUQLhimv1GLLPFk1cc8JCHZ2JwxCRXGsg4gXU'};
        // await deployments.fixture('Testing', {fallbackToGlobal: false});
        
        [implementation, proxy, upgradedImpl] = await Promise.all(['Vault', 'TransparentUpgradeableProxy', 'TestingVault'].map(name => getInstance(name, null)));
        vault = await getInstance('Vault', 'TransparentUpgradeableProxy');
        upgradedVault = await getInstance('TestingVault', 'TransparentUpgradeableProxy');

        inc = new Inc.Lib();
        incBridge = new Inc.Lib();
        // only load WASM once
        await inc.init();
        incBridge.setProvider('http://localhost:9338');
        transactor = await inc.NewTransactor(ethGuy.incPrivate);

        // check initial token balances
        const startBalance = await ethGuy.signer.getBalance();
        expect(startBalance).to.be.at.least(startEth);
        const token1 = await getInstance('Token1');
        const t1StartBalance = await token1.balanceOf(tokenGuy.signer.address);
        expect(t1StartBalance).to.be.at.least(startToken);
        const token2 = await getInstance('Token2');
        testingTokens = [token1, token2];
        console.log(`Test ERC20 tokens ${token1.address} and ${token2.address}`)
        const t2StartBalance = await token2.balanceOf(tokenGuy.signer.address);
        expect(t2StartBalance).to.be.at.least(startToken);
        console.log(`Start balances are ${startBalance.toString()}, ${t1StartBalance.toString()}, ${t2StartBalance.toString()}`);
    });

    // for checking that the networks are alive
    describe.skip('Incognito & ETH RPC query', () => {
        it('should have TX methods', () => {
            expect(transactor).to.be.an('object', 'missing transacting functions')
                .with.property('doPRVTransactionV2');
        });
        it('should return a Burning Address', () => {
            expect(incBridge.rpc).to.be.an('object', 'missing Incognito RPC for Bridge node')
                .with.property('getBurningAddress');
            expect(inc.rpc).to.be.an('object', 'missing Incognito RPC')
                .with.property('getBurningAddress');
            return inc.rpc.getBurningAddress(0)
                .then(resp => {
                    expect(resp).to.be.a('string', 'burning address must be a string');
                    expect(resp).to.have.lengthOf.above(10, 'burning address must be longer than this');
                });
        });
        it('should send ETH and wait for 15 blocks', () => {
            const blocksToWait = 15;
            return confirm(signers[0].sendTransaction({to: accounts[1], value: ethers.utils.parseUnits("0.0006969", 'ether')}), blocksToWait, inc.sleep, 15)
            .then(console.log)
            .then(_ => ethers.provider.getBlockNumber())
            .then(console.log)
        });
    });

    let shield = (ethIn) =>
        it(`should accept ${ethers.utils.formatUnits(ethIn, 'gwei')}Gwei in ETH deposit from Ethereum to Incognito chain`, async function() {
            // as per Incognito specs, deposit and wait more than 15 blocks
            const tx = await confirm(vault.connect(ethGuy.signer).deposit(ethGuy.inc, {value: ethIn}), nConfirm + 1);
            // in this test case, this call should emit only one event
            // console.log(`Deposit : ${tx.hash} => achieved ${nConfirm + 1} confirmations`);
            await expect(tx).to.emit(vault, 'Deposit')
            .withArgs(tokenAddresses.ETH, ethGuy.inc, ethIn);
            const proveResult = await proveEth(Inc.utils.Trie, db, tx.hash, Inc.utils.base64Encode);
            const issuingResponse = await inc.rpc.issueIncToken(ethGuy.incPrivate, tokenAddresses.pETH, proveResult.ethBlockHash, proveResult.encodedProof, proveResult.txIndex);
            await expect(issuingResponse).to.have.property('TxID').that.is.a('string').with.lengthOf(64, "sending issuing TX failed");

            const t = await inc.NewTransactor(ethGuy.incPrivate);
            let change = await t.waitBalanceChange(tokenAddresses.pETH);
            change = await fromIncDecimals(change.balance-change.oldBalance, tokenAddresses.ETH);
            await expect(ethIn).to.equal(change, 'issuing amount mismatch');
        });

    let testMintChi = (chiToMint) => {
        chiToMint = BN.from(chiToMint);
        // minting CHI takes a bunch of gas, so each mint call we send is capped of amount by the gas limit
        it(`should mint ${chiToMint.toNumber()} CHI token`, async function() {
            await mintChi(chiFunder, chiToMint, chi);
            // we only need to make sure we have at least this much CHI
            const chiBalance = await chi.balanceOf(chiFunder.address);
            // console.log(`CHI balance is ${chiBalance.toString()}`);
            await expect(chiBalance).to.be.at.least(chiToMint);
        })
    }

    let unshield = (ethOut, chiToApprove) => {
        chiToApprove = BN.from(chiToApprove);
        it(`should withdraw ${ethers.utils.formatUnits(ethOut, 'gwei')}Gwei in ETH back to Ethereum account`, async function() {
            const amount = await toIncDecimals(ethOut, tokenAddresses.ETH);
            const burnResult = await inc.rpc.burnTokenToContract(ethGuy.incPrivate, tokenAddresses.pETH, amount.toNumber(), ethGuy.signer.address, false);
            await expect(burnResult).to.have.property('TxID').that.is.a('string').with.lengthOf(64, "sending withdraw TX failed");
            await transactor.waitTx(burnResult.TxID, nConfirm + 1);
            const burnProof = await incBridge.rpc.getBurnProof(burnResult.TxID);
            let params = await formatBurnProof(burnProof);

            const balanceBefore = await ethGuy.signer.getBalance();
            const result = await chiify(chiFunder, chiToApprove, vault, 'withdraw', params);
            await expect(result.approveTx).to.emit(chi, 'Approval')
            .withArgs(chiFunder.address, vault.address, chiToApprove);
            
            console.log('CHI analysis : %O', result);
            if (chiToApprove.eq(0)) {
                expect(result.chiUsed).to.equal(0);
            } else {
                expect(result.chiUsed).to.be.above(0).and.at.most(chiToApprove);
            }
            const balance = await ethGuy.signer.getBalance();
            // calculate the ETH balance increase after withdrawal
            await expect(BN.from(balance).sub(BN.from(balanceBefore))).to.equal(ethOut);
        })
    }

    describe('(UN)SHIELD ETHER', () => {
        shield(startEth);
        testMintChi(140);
        const amount = getPartOf(startEth, 50);
        unshield(amount, 0);
        unshield(amount, 50);
    })

    // our test tokens both have decimal 18 (in any case it's not relevant enough to change from 18)
    let shieldToken = (_amount, _tokenIndex) =>
        it(`should accept ${ethers.utils.formatUnits(_amount, 'gwei')}Gwei in ERC20 token deposit`, async function() {
            const tokenContract = testingTokens[_tokenIndex];
            // approve the token amount we're about to spend
            // console.log(tokenContract.address);
            const appr = await confirm(tokenContract.connect(tokenGuy.signer).approve(vault.address, _amount));
            await expect(appr).to.emit(tokenContract, 'Approval')
            .withArgs(tokenGuy.signer.address, vault.address, _amount);
            // as per Incognito specs, deposit and wait more than 15 blocks
            const tx = await confirm(vault.connect(tokenGuy.signer).depositERC20(tokenContract.address, _amount, tokenGuy.inc), nConfirm + 1);
            await expect(tx).to.emit(vault, 'Deposit')
            .withArgs(tokenContract.address, tokenGuy.inc, await toIncDecimals(_amount, tokenContract.address));
            const proveResult = await proveEth(Inc.utils.Trie, db, tx.hash, Inc.utils.base64Encode);
            const issuingResponse = await inc.rpc.issueIncToken(tokenGuy.incPrivate, tokenAddresses.pTokens[_tokenIndex], proveResult.ethBlockHash, proveResult.encodedProof, proveResult.txIndex);
            await expect(issuingResponse).to.have.property('TxID').that.is.a('string').with.lengthOf(64, "sending issuing TX failed");
            const t = await inc.NewTransactor(tokenGuy.incPrivate);
            let change = await t.waitBalanceChange(tokenAddresses.pTokens[_tokenIndex]);
            change = await fromIncDecimals(change.balance-change.oldBalance, tokenContract.address);
            await expect(_amount).to.equal(change, 'token issuing amount mismatch');
        });

    let unshieldToken = (_amount, _tokenIndex, _chiToApprove, _expectations = {balanceChange : true, chiUse: true}) => {
        const chiToApprove = BN.from(_chiToApprove);
        it(`should withdraw ${ethers.utils.formatUnits(_amount, 'gwei')}Gwei in ERC20 token`, async function() {
            const tokenContract = testingTokens[_tokenIndex];
            // switch to Incognito's decimal
            const amount = await toIncDecimals(_amount, tokenContract.address);
            const burnResult = await inc.rpc.burnTokenToContract(tokenGuy.incPrivate, tokenAddresses.pTokens[_tokenIndex], amount.toNumber(), tokenGuy.signer.address, false);
            await expect(burnResult).to.have.property('TxID').that.is.a('string').with.lengthOf(64, "sending withdraw TX failed");
            await transactor.waitTx(burnResult.TxID, nConfirm + 1);
            const burnProof = await incBridge.rpc.getBurnProof(burnResult.TxID);
            let params = await formatBurnProof(burnProof);

            const balanceBefore = await tokenContract.balanceOf(tokenGuy.signer.address);
            const result = await chiify(chiFunder, chiToApprove, vault, 'withdraw', params);
            await expect(result.approveTx).to.emit(chi, 'Approval')
            .withArgs(chiFunder.address, vault.address, chiToApprove);
            
            console.log('CHI analysis : %O', result);
            if (!_expectations.chiUse || chiToApprove.eq(0)) {
                expect(result.chiUsed).to.equal(0);
            } else {
                expect(result.chiUsed).to.be.above(0).and.at.most(chiToApprove);
            }
            const balance = await tokenContract.balanceOf(tokenGuy.signer.address);
            
            if (_expectations.balanceChange) {
                await expect(BN.from(balance).sub(BN.from(balanceBefore))).to.equal(_amount, "withdrawn balance must go to Ethereum address");
            } else {
                await expect(BN.from(balance).sub(BN.from(balanceBefore))).to.equal(0, "there must be no balance change");
            }
        })
    }

    describe('(UN)SHIELD TOKEN', () => {
        shieldToken(startToken, 1);
        shieldToken(startToken, 0);
        testMintChi(140);
        // withdraw less than half of the deposited "Token 2" without any CHI, then do it again after approving 140 CHI
        const amount = getPartOf(startToken, 50);
        unshieldToken(amount, 1, 0);
        unshieldToken(amount, 1, 140);
    })

    // it(`should trade`);

    describe('Test Upgrade & Pausing Feature', () => {
        it('should upgrade to new implementation, preserving balances', async function() {
            let myTokenAddresses = testingTokens.map(t => t.address);
            myTokenAddresses.push(tokenAddresses.ETH);
            // proxy starts off pointing to Vault
            let currentImpl = await getImplementation(proxy.address);
            await expect(currentImpl).to.equal(implementation.address);
            // get total balance to compare later
            const balancesBeforeUpgrade = await Promise.all(myTokenAddresses.map(_tokenAddr => vault.balanceOf(_tokenAddr)));

            // new upgrade method (Proxy Pattern) means prevVault no longer changes
            // const data = upgradedImpl.interface.encodeFunctionData('callAfterUpgrade', [true]);
            // this proxy was initialized upon creation, so state changes that need to be done after upgrade must be via another function
            // console.log("will upgrade to", upgradedImpl.address);
            await proxy.connect(vaultAdmin).upgradeTo(upgradedImpl.address);
            currentImpl = await getImplementation(proxy.address);
            await expect(currentImpl).to.equal(upgradedImpl.address);

            // plug ABI of new vault to the proxy i just upgraded
            const balances = await Promise.all(myTokenAddresses.map(_tokenAddr => upgradedVault.balanceOf(_tokenAddr)));
            console.log('balances before upgrade', balancesBeforeUpgrade);
            console.log('after upgrade', balances);
            await expect(balances).to.deep.equal(balancesBeforeUpgrade);
            console.log('New implementation sets CHI to an EOA');
            await confirm(upgradedVault.setChi(ethGuy.signer.address), 2);
        });
        // withdraw less than half of the deposited "Token 1"
        testMintChi(140);
        const amount = getPartOf(startToken, 25);
        // expect to withdraw successfully but no CHI used
        unshieldToken(amount, 0, 20, {chiUse : false, balanceChange : true});

        it('suppose the new implementation acts maliciously', async function() {
            // make the new vault use a non-CHI contract for CHI, as well as not paying withdrawn balance
            await confirm(upgradedVault.callAfterUpgrade(true), 2);
            await confirm(upgradedVault.setChi(testingTokens[0].address), 2);
        });
        unshieldToken(amount, 0, 20, {chiUse : false, balanceChange : false});

        it('should pause, rejecting all calls', async function() {
            const tokenContract = testingTokens[0];
            // "false" means we expect no balance change, since this implementation we just upgraded to is flawed
            // now we need to pause the proxy to avoid more damage
            await confirm(proxy.connect(vaultAdmin).pause(), 2);
            // after pause, all calls should revert
            const appr = await confirm(tokenContract.connect(tokenGuy.signer).approve(upgradedVault.address, amount));
            await expect(appr).to.emit(tokenContract, 'Approval')
            .withArgs(tokenGuy.signer.address, upgradedVault.address, amount);
            await expect(upgradedVault.connect(chiFunder).depositERC20(tokenContract.address, amount, tokenGuy.inc)).to.be.reverted;
        });

        it('should upgrade to previous safe implementation & unpause', async function() {
            await confirm(proxy.connect(vaultAdmin).upgradeTo(implementation.address), 2);
            await confirm(proxy.connect(vaultAdmin).unpause(), 15);

            const currentImpl = await getImplementation(proxy.address);
            await expect(currentImpl).to.equal(implementation.address);
        });
        // shield some token & make sure it succeeds
        shieldToken(amount, 0);
    });
})