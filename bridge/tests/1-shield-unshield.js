const { expect } = require("chai");
const hre = require("hardhat");
const { deployments, ethers } = hre;

// const cfg = require('../hardhat.config');
const {tokenAddresses} = require('../scripts/constants');
const {BN, getPartOf, confirm, fromIncDecimals, toIncDecimals, getInstance, getImplementation, Inc} = require('../scripts/external');
const {proveEth, formatBurnProof, db} = require('../scripts/prove');
const {getChi, chiify, mintChi} = require('../scripts/chi');
const path = require('path');

let setupTest = () => {
    return async function() {
        this.nConfirm = 15;
        console.log("Network", await ethers.provider.getNetwork());
        this.chi = await getChi();
        console.log("CHI Token contract", this.chi.address);
        // console.log(await chi.balanceOf(chiFunder.address));
        this.signers = await ethers.getSigners();
        this.accounts = this.signers.map(s => s.address);
        // console.log('Signers :', this.accounts);
        // 0 is deployer, we use 1 2 3 5 for named accounts, skipping 4
        [_, this.vaultAdmin, this.ethUser, this.tokenUser, _, this.chiFunder] = this.signers;

        this.ethGuy = {signer: this.ethUser, inc: '12sxXUjkMJZHz6diDB6yYnSjyYcDYiT5QygUYFsUbGUqK8PH8uhxf4LePiAE8UYoDcNkHAdJJtT1J6T8hcvpZoWLHAp8g6h1BQEfp4h5LQgEPuhMpnVMquvr1xXZZueLhTNCXc8fkVXseeVAGCt8', incPrivate: '112t8rnXoBXrThDTACHx2rbEq7nBgrzcZhVZV4fvNEcGJetQ13spZRMuW5ncvsKA1KvtkauZuK2jV8pxEZLpiuHtKX3FkKv2uC5ZeRC8L6we'};
        this.tokenGuy = {signer: this.tokenUser, inc: '12sttFKciCWyRbNsK1yD1mWEwZoeWi1JtWJZ7gKTbx5eB25U4FnrfkxgxbnZ8zDn2QNhhW44HBZJ1EnfwVBueR44D5ucWxGNpXZMawoCmv6G2cwKi4xkasuysu3WtpV5ZMSYgaJ1mwe9fqgVD9mh', incPrivate: '112t8rnZUQXxcbayAZvyyZyKDhwVJBLkHuTKMhrS51nQZcXKYXGopUTj22JtZ8KxYQcak54KUQLhimv1GLLPFk1cc8JCHZ2JwxCRXGsg4gXU'};

        [this.implementation, this.proxy, this.upgradedImpl] = await Promise.all(['Vault', 'TransparentUpgradeableProxy', 'TestingVault'].map(name => getInstance(name, null)));
        this.vault = await getInstance('Vault', 'TransparentUpgradeableProxy');
        this.upgradedVault = await getInstance('TestingVault', 'TransparentUpgradeableProxy');

        this.inc = new Inc.Lib();
        this.incBridge = new Inc.Lib();
        // only load WASM once
        await this.inc.init();
        this.incBridge.setProvider('http://localhost:9338');
        this.transactor = await this.inc.NewTransactor(this.ethGuy.incPrivate);

        // check initial token balances
        const token1 = await getInstance('Token1');
        const token2 = await getInstance('Token2');
        this.testingTokens = [token1, token2];
        console.log(`Test ERC20 tokens ${token1.address} and ${token2.address}`)
    }
}

let testMintChi = (chiToMint) => {
    chiToMint = BN.from(chiToMint);
    return async function() {
        // minting CHI takes a bunch of gas, so each mint call we send is capped of amount by the gas limit
        await mintChi(this.chiFunder, chiToMint, this.chi);
        // we only need to make sure we have at least this much CHI
        const chiBalance = await this.chi.balanceOf(this.chiFunder.address);
        // console.log(`CHI balance is ${chiBalance.toString()}`);
        await expect(chiBalance).to.be.at.least(chiToMint);
    }
}

let shield = (ethIn) => {
    return async function() {
        // as per Incognito specs, deposit and wait more than 15 blocks
        const tx = await confirm(this.vault.connect(this.ethGuy.signer).deposit(this.ethGuy.inc, {value: ethIn}), this.nConfirm + 1);
        // in this test case, this call should emit only one event
        // console.log(`Deposit : ${tx.hash} => achieved ${nConfirm + 1} confirmations`);
        await expect(tx).to.emit(this.vault, 'Deposit')
        .withArgs(tokenAddresses.ETH, this.ethGuy.inc, ethIn);
        const proveResult = await proveEth(Inc.utils.Trie, db, tx.hash, Inc.utils.base64Encode);
        const issuingResponse = await this.inc.rpc.issueIncToken(this.ethGuy.incPrivate, tokenAddresses.pETH, proveResult.ethBlockHash, proveResult.encodedProof, proveResult.txIndex);
        await expect(issuingResponse).to.have.property('TxID').that.is.a('string').with.lengthOf(64, "sending issuing TX failed");

        const t = await this.inc.NewTransactor(this.ethGuy.incPrivate);
        let change = await t.waitBalanceChange(tokenAddresses.pETH);
        change = await fromIncDecimals(change.balance - change.oldBalance, tokenAddresses.ETH);
        await expect(ethIn).to.equal(change, 'issuing amount mismatch');
    }
}

let unshield = (ethOut, chiToApprove) => {
    chiToApprove = BN.from(chiToApprove);
    return async function() {
        const amount = await toIncDecimals(ethOut, tokenAddresses.ETH);
        const burnResult = await this.inc.rpc.burnTokenToContract(this.ethGuy.incPrivate, tokenAddresses.pETH, amount.toNumber(), this.ethGuy.signer.address, false);
        await expect(burnResult).to.have.property('TxID').that.is.a('string').with.lengthOf(64, "sending withdraw TX failed");
        await this.transactor.waitTx(burnResult.TxID, this.nConfirm + 1);
        const burnProof = await this.incBridge.rpc.getBurnProof(burnResult.TxID);
        let params = await formatBurnProof(burnProof);

        const balanceBefore = await this.ethGuy.signer.getBalance();
        const result = await chiify(this.chiFunder, chiToApprove, this.vault, 'withdraw', params);
        await expect(result.approveTx).to.emit(this.chi, 'Approval')
        .withArgs(this.chiFunder.address, this.vault.address, chiToApprove);

        console.log('CHI analysis : %O', result);
        if (chiToApprove.eq(0)) {
            expect(result.chiUsed).to.equal(0);
        } else {
            expect(result.chiUsed).to.be.above(0).and.at.most(chiToApprove);
        }
        const balance = await this.ethGuy.signer.getBalance();
        // calculate the ETH balance increase after withdrawal
        await expect(BN.from(balance).sub(BN.from(balanceBefore))).to.equal(ethOut);
    }
}

let shieldToken = (_amount, _tokenIndex) => {
    return async function() {
        const tokenContract = this.testingTokens[_tokenIndex];
        // approve the token amount we're about to spend
        // console.log(tokenContract.address);
        const appr = await confirm(tokenContract.connect(this.tokenGuy.signer).approve(this.vault.address, _amount));
        await expect(appr).to.emit(tokenContract, 'Approval')
        .withArgs(this.tokenGuy.signer.address, this.vault.address, _amount);
        // as per Incognito specs, deposit and wait more than 15 blocks
        const tx = await confirm(this.vault.connect(this.tokenGuy.signer).depositERC20(tokenContract.address, _amount, this.tokenGuy.inc), this.nConfirm + 1);
        await expect(tx).to.emit(this.vault, 'Deposit')
        .withArgs(tokenContract.address, this.tokenGuy.inc, await toIncDecimals(_amount, tokenContract.address));
        const proveResult = await proveEth(Inc.utils.Trie, db, tx.hash, Inc.utils.base64Encode);
        const issuingResponse = await this.inc.rpc.issueIncToken(this.tokenGuy.incPrivate, tokenAddresses.pTokens[_tokenIndex], proveResult.ethBlockHash, proveResult.encodedProof, proveResult.txIndex);
        await expect(issuingResponse).to.have.property('TxID').that.is.a('string').with.lengthOf(64, "sending issuing TX failed");
        const t = await this.inc.NewTransactor(this.tokenGuy.incPrivate)
        let change = await t.waitBalanceChange(tokenAddresses.pTokens[_tokenIndex]);
        change = await fromIncDecimals(change.balance - change.oldBalance, tokenContract.address);
        await expect(_amount).to.equal(change, 'token issuing amount mismatch');
    }
}

let unshieldToken = (_amount, _tokenIndex, _chiToApprove, _expectations = {balanceChange : true, chiUse: true}) => {
    const chiToApprove = BN.from(_chiToApprove);
    return async function() {
        const tokenContract = this.testingTokens[_tokenIndex];
        // switch to Incognito's decimal
        const amount = await toIncDecimals(_amount, tokenContract.address);
        const burnResult = await this.inc.rpc.burnTokenToContract(this.tokenGuy.incPrivate, tokenAddresses.pTokens[_tokenIndex], amount.toNumber(), this.tokenGuy.signer.address, false);
        await expect(burnResult).to.have.property('TxID').that.is.a('string').with.lengthOf(64, "sending withdraw TX failed");
        await this.transactor.waitTx(burnResult.TxID, this.nConfirm + 1);
        const burnProof = await this.incBridge.rpc.getBurnProof(burnResult.TxID);
        let params = await formatBurnProof(burnProof);

        const balanceBefore = await tokenContract.balanceOf(this.tokenGuy.signer.address);
        const result = await chiify(this.chiFunder, chiToApprove, this.vault, 'withdraw', params);
        await expect(result.approveTx).to.emit(this.chi, 'Approval')
        .withArgs(this.chiFunder.address, this.vault.address, chiToApprove);

        console.log('CHI analysis : %O', result);
        if (!_expectations.chiUse || chiToApprove.eq(0)) {
            expect(result.chiUsed).to.equal(0);
        } else {
            expect(result.chiUsed).to.be.above(0).and.at.most(chiToApprove);
        }
        const balance = await tokenContract.balanceOf(this.tokenGuy.signer.address);

        if (_expectations.balanceChange) {
            await expect(BN.from(balance).sub(BN.from(balanceBefore))).to.equal(_amount, "withdrawn balance must go to Ethereum address");
        } else {
            await expect(BN.from(balance).sub(BN.from(balanceBefore))).to.equal(0, "there must be no balance change");
        }
    }
}

describe("Vault - Shield & Unshield with CHI", () => {
    const startEth = ethers.utils.parseUnits('0.00005', 'ether');
    const startToken = ethers.utils.parseUnits('0.00001', 'ether');
    before(setupTest());

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

    describe('(UN)SHIELD ETHER', async function() {
        it(`should receive ${ethers.utils.formatUnits(startEth, 'gwei')} (gwei) ETH deposit`, shield(startEth));
        it('should mint CHI', testMintChi(140));
        const amount = getPartOf(startEth, 50);
        it(`should withdraw ${ethers.utils.formatUnits(amount, 'gwei')} using no CHI`, unshield(amount, 0));
        it(`should withdraw ${ethers.utils.formatUnits(amount, 'gwei')} using some CHI`, unshield(amount, 50));
    });

    describe('(UN)SHIELD TOKEN', async function() {
        it(`should receive ${ethers.utils.formatUnits(startEth, 'gwei')} (gwei) Token1 deposit`, shieldToken(startToken, 0));
        it(`should receive ${ethers.utils.formatUnits(startEth, 'gwei')} (gwei) Token2 deposit`, shieldToken(startToken, 1));
        it('should mint CHI', testMintChi(140));
        // withdraw less than half of the deposited "Token 2" without any CHI, then do it again after approving 140 CHI
        const amount = getPartOf(startToken, 50);
        it(`should withdraw ${ethers.utils.formatUnits(amount, 'gwei')} using no CHI`, unshieldToken(amount, 1, 0));
        it(`should withdraw ${ethers.utils.formatUnits(amount, 'gwei')} using some CHI`, unshieldToken(amount, 1, 50));
    });

    // it(`should trade`);
})

module.exports = {
    setupTest,
    testMintChi,
    shield,
    shieldToken,
    unshield,
    unshieldToken
}