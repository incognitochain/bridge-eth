const { expect } = require("chai");
const hre = require("hardhat");
const { deployments, ethers } = hre;
const BN = ethers.BigNumber;
const {tokenAddresses} = require('../scripts/constants');
const {getPartOf, chooseOneFrom, confirm, fromIncDecimals, toIncDecimals, getInstance, getImplementation, generateTestIncTokenID} = require('../scripts/utils');
const { Inc } = require('../scripts/external');
const {proveEth, formatBurnProof} = require('../scripts/prove');
const path = require('path');

let setupTest = () => {
    return async function() {
        this.nConfirm = 15;
        this.signers = await ethers.getSigners();
        this.accounts = this.signers.map(s => s.address);
        // console.log('Signers :', this.accounts);
        // 0 is deployer, we use 1 2 3 5 for named accounts, skipping 4
        [this.deployer, this.vaultAdmin, this.ethUser, this.tokenUser, _, this.unshieldSender] = this.signers;

        this.ethGuy = {signer: this.ethUser, inc: '12sxXUjkMJZHz6diDB6yYnSjyYcDYiT5QygUYFsUbGUqK8PH8uhxf4LePiAE8UYoDcNkHAdJJtT1J6T8hcvpZoWLHAp8g6h1BQEfp4h5LQgEPuhMpnVMquvr1xXZZueLhTNCXc8fkVXseeVAGCt8', incPrivate: '112t8rnXoBXrThDTACHx2rbEq7nBgrzcZhVZV4fvNEcGJetQ13spZRMuW5ncvsKA1KvtkauZuK2jV8pxEZLpiuHtKX3FkKv2uC5ZeRC8L6we'};
        this.tokenGuy = {signer: this.tokenUser, inc: '12sttFKciCWyRbNsK1yD1mWEwZoeWi1JtWJZ7gKTbx5eB25U4FnrfkxgxbnZ8zDn2QNhhW44HBZJ1EnfwVBueR44D5ucWxGNpXZMawoCmv6G2cwKi4xkasuysu3WtpV5ZMSYgaJ1mwe9fqgVD9mh', incPrivate: '112t8rnZUQXxcbayAZvyyZyKDhwVJBLkHuTKMhrS51nQZcXKYXGopUTj22JtZ8KxYQcak54KUQLhimv1GLLPFk1cc8JCHZ2JwxCRXGsg4gXU'};

        this.recoveryAddress = '12sxXUjkMJZHz6diDB6yYnSjyYcDYiT5QygUYFsUbGUqK8PH8uhxf4LePiAE8UYoDcNkHAdJJtT1J6T8hcvpZoWLHAp8g6h1BQEfp4h5LQgEPuhMpnVMquvr1xXZZueLhTNCXc8fkVXseeVAGCt8';
        this.recoveryPrivateKey = '112t8rnXoBXrThDTACHx2rbEq7nBgrzcZhVZV4fvNEcGJetQ13spZRMuW5ncvsKA1KvtkauZuK2jV8pxEZLpiuHtKX3FkKv2uC5ZeRC8L6we';
        // TestingVault is never in live deployment, so it is deployed in test setup
        const { deploy } = deployments;
        const deployResult = await deploy('TestingVault', {from: this.deployer.address, args: [], log: true});
        [this.implementation, this.proxy, this.upgradedImpl, this.vaultHelper] = await Promise.all(['Vault', 'TransparentUpgradeableProxy', 'TestingVault', 'VaultHelper'].map(name => getInstance(name, null)));
        this.vault = await getInstance('Vault', 'TransparentUpgradeableProxy');
        this.upgradedVault = await getInstance('TestingVault', 'TransparentUpgradeableProxy');

        await Inc.init();
        this.inc = new Inc.SimpleWallet();
        this.incBridge = new Inc.SimpleWallet();
        
        this.inc.setProvider(hre.networkCfg().providers[0]);
        this.incBridge.setProvider(hre.networkCfg().providers[1]);
        this.transactor = await this.inc.NewTransactor(this.ethGuy.incPrivate);

        const names = ['Token1', 'Token2', 'Token3'];
        if (hre.networkCfg().deployed.testingTokens) {
            // for testnet, make sure testingTokens in network cfg has 3 or more token addresses
            // unless it's a forked network, we assume all tx senders have Ether, and token actor is funded with Token1 (MKR)
            this.testingTokens = await Promise.all(hre.networkCfg().deployed.testingTokens.map(t =>
                getInstance('contracts/IERC20.sol:IERC20', null, t)))
            if (hre.networkCfg().deployed.tokenFunder) {
                let b = await this.testingTokens[0].balanceOf(this.tokenGuy.signer.address);
                if (b.lt(ethers.utils.parseUnits('1', 'ether'))) {
                    // when in forked network, fund token 1 using impersonated address
                    const tokenFunder = await ethers.getSigner(hre.networkCfg().deployed.tokenFunder);
                    await confirm(this.testingTokens[0].connect(tokenFunder).transfer(this.tokenGuy.signer.address, ethers.utils.parseUnits('10', 'ether')));
                    b = await this.testingTokens[0].balanceOf(this.tokenGuy.signer.address);
                    console.log(`Got ${b.toString()} of token ${this.testingTokens[0].address} funded to ${this.tokenGuy.signer.address}`);
                }
            }
        } else {
            this.testingTokens = await Promise.all(names.map(n => getInstance(n)));
            // reserve one for testing shield / unshield; others are for trading into
            // [ this.currentToken, this.tradingTokens ] = chooseOneFrom(this.testingTokens);
        }
        // store token name map
        this.tokens = {};
        names.forEach((n, index) => {
            this.tokens[n] = this.testingTokens[index];
        })
        console.log(`using ERC20 tokens ${this.testingTokens.map(t => t.address)}`);
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
        const proveResult = await proveEth(tx.hash, Inc.utils.base64Encode);
        const txor = await this.inc.NewTransactor(this.ethGuy.incPrivate);
        const issuingResponse = await txor.shield({ transfer: { fee: 10, tokenID: tokenAddresses.pETH }, extra: { ethBlockHash: proveResult.ethBlockHash, ethDepositProof: proveResult.encodedProof, txIndex: proveResult.txIndex }});
        // const issuingResponse = await this.inc.rpc.issueIncToken(this.ethGuy.incPrivate, tokenAddresses.pETH, proveResult.ethBlockHash, proveResult.encodedProof, proveResult.txIndex);
        await expect(issuingResponse).to.have.nested.property('Response.txId').that.is.a('string').with.lengthOf(64, "sending issuing TX failed");
        let change = await txor.waitBalanceChange(tokenAddresses.pETH);
        change = BN.from(change.balance) - BN.from(change.oldBalance);
        change = await fromIncDecimals(change, tokenAddresses.ETH);
        await expect(change.toString()).to.equal(ethIn.toString(), 'issuing amount mismatch');
    }
}

let unshield = (ethOut) => {
    return async function() {
        const amount = await toIncDecimals(ethOut, tokenAddresses.ETH);
        const txor = await this.inc.NewTransactor(this.ethGuy.incPrivate);
        const burnResult = await txor.burn({ transfer: { fee: 10, tokenID: tokenAddresses.pETH, info: "BURN ETHER" }, extra: { remoteAddress: this.ethGuy.signer.address, burnAmount: amount.toString(),  burningType: Inc.constants.BurningRequestMeta }});
        await expect(burnResult).to.have.nested.property('Response.txId').that.is.a('string').with.lengthOf(64, "sending withdraw TX failed");
        await this.transactor.waitTx(burnResult.Response.txId, this.nConfirm + 1);
        const burnProof = await this.incBridge.rpc.getBurnProof(burnResult.Response.txId, false);
        let params = await formatBurnProof(burnProof);

        const balanceBefore = await this.ethGuy.signer.getBalance();
        const tx = await confirm(this.vault.connect(this.unshieldSender).withdraw(...params));
        await expect(tx).to.emit(this.vault, 'Withdraw')
        .withArgs(tokenAddresses.ETH, this.ethGuy.signer.address, ethOut);
        const receipt = await tx.wait();
        console.log(`Unshield gas : ${receipt.gasUsed.toString()}`);

        const balance = await this.ethGuy.signer.getBalance();
        // calculate the ETH balance increase after withdrawal
        await expect(BN.from(balance).sub(BN.from(balanceBefore))).to.equal(ethOut);
    }
}

let unshieldRevertRecover = (ethAmount, fixedTestcase = null) => {
    return async function() {
        const amount = await toIncDecimals(ethAmount);
        let txor = await this.inc.NewTransactor(this.ethGuy.incPrivate);
        let burnProof;
        if (!fixedTestcase) {
            // unshield pETH to a contract address that cannot receive ETH
            const burnResult = await txor.burn({transfer: { fee: 10, tokenID: tokenAddresses.pETH, info: "BURN ETHER" }, extra: { remoteAddress: this.testingTokens[0].address, burnAmount: amount.toString(), burningType: Inc.constants.BurningRequestMeta }});
            await expect(burnResult).to.have.nested.property('Response.txId').that.is.a('string').with.lengthOf(64, "sending withdraw TX failed");
            await this.transactor.waitTx(burnResult.Response.txId, this.nConfirm + 1);
            burnProof = await this.incBridge.rpc.getBurnProof(burnResult.Response.txId, false);
        } else {
            // make sure the testing network was deployed with "mainnet" committees
            console.log('Simulating a recovery on mainnet');
            burnProof = fixedTestcase;
        }

        let params = await formatBurnProof(burnProof);
        const balanceBefore = await this.ethGuy.signer.getBalance();
        const tx = await confirm(this.vault.connect(this.unshieldSender).withdraw(...params), this.nConfirm + 1);
        await expect(tx).to.not.emit(this.vault, 'Withdraw');
        await expect(tx).to.emit(this.vault, 'Deposit')
        .withArgs(tokenAddresses.ETH, this.recoveryAddress, ethAmount);
        
        const balance = await ethers.provider.getBalance(this.testingTokens[0].address);
        // calculate the ETH balance increase after withdrawal
        await expect(balance).to.equal(BN.from(0));

        // issue pETH back to the recovery address
        if (this.recoveryPrivateKey) {
            const proveResult = await proveEth(tx.hash, Inc.utils.base64Encode);
            txor = await this.inc.NewTransactor(this.recoveryPrivateKey);
            const issuingResponse = await txor.shield({ transfer: { fee: 10, tokenID: tokenAddresses.pETH }, extra: { ethBlockHash: proveResult.ethBlockHash, ethDepositProof: proveResult.encodedProof, txIndex: proveResult.txIndex }});
            await expect(issuingResponse).to.have.nested.property('Response.txId').that.is.a('string').with.lengthOf(64, "sending issuing TX failed");
            let change = await txor.waitBalanceChange(tokenAddresses.pETH);
            change = BN.from(change.balance) - BN.from(change.oldBalance);
            await expect(change.toString()).to.equal(amount.toString(), 'issuing amount mismatch');
        }
    }
}

let shieldToken = (_amount, _tokenName) => {
    return async function() {
        const tokenContract = this.tokens[_tokenName];
        // approve the token amount we're about to spend
        // console.log(tokenContract.address);
        const appr = await confirm(tokenContract.connect(this.tokenGuy.signer).approve(this.vault.address, _amount));
        await expect(appr).to.emit(tokenContract, 'Approval')
        .withArgs(this.tokenGuy.signer.address, this.vault.address, _amount);
        // as per Incognito specs, deposit and wait more than 15 blocks
        const incAmount = await toIncDecimals(_amount, tokenContract.address);
        const tx = await confirm(this.vault.connect(this.tokenGuy.signer).depositERC20(tokenContract.address, _amount, this.tokenGuy.inc), this.nConfirm + 1);
        await expect(tx).to.emit(this.vault, 'Deposit')
        .withArgs(ethers.utils.getAddress(tokenContract.address), this.tokenGuy.inc, incAmount);
        const proveResult = await proveEth(tx.hash, Inc.utils.base64Encode);
        const txor = await this.inc.NewTransactor(this.tokenGuy.incPrivate);
        const issuingResponse = await txor.shield({ transfer: { fee: 10, tokenID: generateTestIncTokenID(tokenContract.address) }, extra: { ethBlockHash: proveResult.ethBlockHash, ethDepositProof: proveResult.encodedProof, txIndex: proveResult.txIndex }});
        await expect(issuingResponse).to.have.nested.property('Response.txId').that.is.a('string').with.lengthOf(64, "sending issuing TX failed");
        let change = await txor.waitBalanceChange(generateTestIncTokenID(tokenContract.address));
        await expect(incAmount).to.equal(change.balance - change.oldBalance, 'token issuing amount mismatch');
    }
}

let unshieldToken = (_amount, _tokenName, _expectations = {balanceChange : true}) => {
    return async function() {
        const tokenContract = this.tokens[_tokenName];
        // switch to Incognito's decimal
        const incAmount = await toIncDecimals(_amount, tokenContract.address);
        const txor = await this.inc.NewTransactor(this.tokenGuy.incPrivate);
        const burnResult = await txor.burn({ transfer: { fee: 10, tokenID: generateTestIncTokenID(tokenContract.address), info: "BURN TOKEN" }, extra: { remoteAddress: this.tokenGuy.signer.address, burnAmount: incAmount.toString(), burningType: Inc.constants.BurningRequestMeta }});
        await expect(burnResult).to.have.nested.property('Response.txId').that.is.a('string').with.lengthOf(64, "sending withdraw TX failed");
        await this.transactor.waitTx(burnResult.Response.txId, this.nConfirm + 1);
        const burnProof = await this.incBridge.rpc.getBurnProof(burnResult.Response.txId, false);
        let params = await formatBurnProof(burnProof);

        const balanceBefore = await tokenContract.balanceOf(this.tokenGuy.signer.address);
        const tx = await confirm(this.vault.connect(this.unshieldSender).withdraw(...params));

        const balance = await tokenContract.balanceOf(this.tokenGuy.signer.address);
        if (_expectations.balanceChange) {
            await expect(tx).to.emit(this.vault, 'Withdraw')
            .withArgs(tokenContract.address, this.tokenGuy.signer.address, _amount);
            await expect(BN.from(balance).sub(BN.from(balanceBefore))).to.equal(_amount, "withdrawn balance must go to Ethereum address");
        } else {
            await expect(BN.from(balance).sub(BN.from(balanceBefore))).to.equal(0, "there must be no balance change");
        }
    }
}

module.exports = {
    setupTest,
    shield,
    shieldToken,
    unshield,
    unshieldRevertRecover,
    unshieldToken
}