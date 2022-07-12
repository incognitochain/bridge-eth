const { expect } = require("chai");
const hre = require("hardhat");
const { deployments, ethers } = hre;
const BN = ethers.BigNumber;
const { getPartOf, chooseOneFrom, confirm, fromIncDecimals, toIncDecimals, getInstance, getImplementation, generateTestIncTokenID, prepareExternalCallByVault, getBridgedIncTokenInfo } = require('../scripts/utils');
const { tokenAddresses } = require('../scripts/constants');
const { proveEth, formatBurnProof } = require('../scripts/prove');
const { inc } = require('../scripts/external');

const unified_shield = (ethIn) => {
    return async function() {
        // as per Incognito specs, deposit and wait more than 15 blocks
        const tx = await confirm(this.vault.connect(this.ethGuy.signer).deposit(this.ethGuy.inc, { value: ethIn }), this.nConfirm);
        // in this test case, this call should emit only one event
        // console.log(`Deposit : ${tx.hash} => achieved ${nConfirm + 1} confirmations`);
        await expect(tx).to.emit(this.vault, 'Deposit')
            .withArgs(tokenAddresses.ETH, this.ethGuy.inc, ethIn);
        const proveResult = await proveEth(tx.hash, inc.utils.base64Encode);
        const txor = await inc.NewTransactor(this.ethGuy.incPrivate);
        const { result: issuingResponse } = await inc.Tx(txor).unified_shield(tokenAddresses.pETH, proveResult.ethBlockHash, proveResult.encodedProof, proveResult.txIndex, 1, tokenAddresses.unify[tokenAddresses.ETH]).send().catch(console.error);
        await expect(issuingResponse).to.have.nested.property('txId').that.is.a('string').with.lengthOf(64, "sending issuing TX failed");
        let [change] = await Promise.all([txor.waitBalanceChange(tokenAddresses.unify[tokenAddresses.ETH]), txor.waitTx(issuingResponse.txId)]);
        change = BN.from(change.balance).sub(BN.from(change.oldBalance));
        change = await fromIncDecimals(change, tokenAddresses.ETH);
        await expect(change).to.equal(ethIn, 'issuing amount mismatch');
    }
}

const unified_unshield = (amount) => {
    return async function() {
        const incAmount = await toIncDecimals(amount, tokenAddresses.ETH);
        const txor = await inc.NewTransactor(this.ethGuy.incPrivate);
        const { result: burnResult } = await inc.Tx(txor).withTokenID(tokenAddresses.unify[tokenAddresses.ETH]).withInfo("BURN ETHER").unified_unshield([{
            BurningAmount: incAmount.toString(),
            RemoteAddress: this.ethGuy.signer.address,
            MinExpectedAmount: incAmount.toString(),
            IncTokenID: tokenAddresses.pETH,
        }], false).send().catch(console.error);
        await expect(burnResult).to.have.nested.property('txId').that.is.a('string').with.lengthOf(64, "sending withdraw TX failed");
        await this.transactor.waitTx(burnResult.txId);
        const burnProof = await inc.rpc.unified_getBurnProof(burnResult.txId);
        let params = await formatBurnProof(burnProof);

        const balanceBefore = await this.ethGuy.signer.getBalance();
        const tx = await confirm(this.vault.connect(this.unshieldSender).withdraw(...params));
        await expect(tx).to.emit(this.vault, 'Withdraw')
        // .withArgs(tokenAddresses.ETH, this.ethGuy.signer.address, ethOut);
        const receipt = await tx.wait();
        console.log(`Unshield gas : ${receipt.gasUsed.toString()}`);

        const balance = await this.ethGuy.signer.getBalance();
        // calculate the ETH balance increase after withdrawal
        console.log(amount.toString());
        console.log(BN.from(balance).sub(BN.from(balanceBefore)).toString());
        await expect(BN.from(balance).sub(BN.from(balanceBefore))).to.equal(amount);
    }
}

let unified_shieldToken = (_amount, _tokenName) => {
    return async function() {
        const tokenContract = getBridgedIncTokenInfo(this, _tokenName);
        // approve the token amount we're about to spend
        // console.log(tokenContract.address);
        const appr = await confirm(tokenContract.connect(this.tokenGuy.signer).approve(this.vault.address, _amount));
        await expect(appr).to.emit(tokenContract, 'Approval')
            .withArgs(this.tokenGuy.signer.address, this.vault.address, _amount);
        // as per Incognito specs, deposit and wait more than 15 blocks
        const incAmount = await toIncDecimals(_amount, tokenContract.address, true);
        const emitAmount = await toIncDecimals(_amount, tokenContract.address, false);
        const tx = await confirm(this.vault.connect(this.tokenGuy.signer).depositERC20(tokenContract.address, _amount, this.tokenGuy.inc), this.nConfirm);
        await expect(tx).to.emit(this.vault, 'Deposit')
            .withArgs(ethers.utils.getAddress(tokenContract.address), this.tokenGuy.inc, emitAmount);
        const proveResult = await proveEth(tx.hash, inc.utils.base64Encode);
        const txor = await inc.NewTransactor(this.tokenGuy.incPrivate);
        const { result: issuingResponse } = await inc.Tx(txor).unified_shield(tokenContract.inc, proveResult.ethBlockHash, proveResult.encodedProof, proveResult.txIndex, 1, tokenContract.unify).send().catch(console.error);
        await expect(issuingResponse).to.have.nested.property('txId').that.is.a('string').with.lengthOf(64, "sending issuing TX failed");
        let [change] = await Promise.all([txor.waitBalanceChange(tokenContract.unify), txor.waitTx(issuingResponse.txId)]);
        change = BN.from(change.balance).sub(BN.from(change.oldBalance));
        await expect(change).to.equal(incAmount, 'token issuing amount mismatch');
    }
}

let unified_unshieldToken = (_amount, _tokenName, _expectations = { balanceChange: true }) => {
    return async function() {
        const tokenContract = getBridgedIncTokenInfo(this, _tokenName);
        // switch to Incognito's decimal
        const incAmount = await toIncDecimals(_amount, tokenContract.address, true);
        const txor = await inc.NewTransactor(this.tokenGuy.incPrivate);
        const { result: burnResult } = await inc.Tx(txor).withTokenID(tokenContract.unify).withInfo("BURN TOKEN").unified_unshield([{
            BurningAmount: incAmount.toString(),
            RemoteAddress: this.tokenGuy.signer.address,
            MinExpectedAmount: incAmount.toString(),
            IncTokenID: tokenContract.inc,
        }], false).send().catch(console.error);
        await expect(burnResult).to.have.nested.property('txId').that.is.a('string').with.lengthOf(64, "sending withdraw TX failed");
        await this.transactor.waitTx(burnResult.txId);
        const burnProof = await inc.rpc.unified_getBurnProof(burnResult.txId);
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

const burnForCall = (exchangeName, amount, srcTokenName = null, tradedTokenName = null, options = { useWithdraw: false, useUnified: { burn: true, reshield: true }}) => {
    return async function() {
        const [srcToken, tradedToken] = [srcTokenName, tradedTokenName].map(n => getBridgedIncTokenInfo(this, n));
        const txor = await inc.NewTransactor(srcToken.sender.incPrivate);
        const { useWithdraw, useUnified } = options;
        const withdrawAddress = useWithdraw ? srcToken.sender.signer.address : undefined;
        const incAmount = await toIncDecimals(amount, srcToken.address, useUnified.burn);
        // encode calldata
        const timestamp = ethers.utils.hexlify(ethers.utils.randomBytes(6));
        const { encodedTradeCall, hashedData, exchange, tradeReturn } = await prepareExternalCallByVault(this, exchangeName, srcToken.address, tradedToken.address, timestamp, amount);
        const externalNetworkID = 1; // test networkID: Ethereum
        // create & send burn tx
        const { result: burnResult } = await inc.Tx(txor).withTokenID(useUnified.burn ? srcToken.unify : srcToken.inc).withInfo("BURN ETHER")
            .burnForCall(incAmount.toString(), srcToken.inc, externalNetworkID, exchange.address,
                encodedTradeCall, tradedToken.address, withdrawAddress)
            .send().catch(console.error);
        await expect(burnResult).to.have.nested.property('txId').that.is.a('string').with.lengthOf(64, "sending withdraw TX failed");
        await this.transactor.waitTx(burnResult.txId);
        const burnProof = await inc.rpc.unified_getBurnProof(burnResult.txId);
        let params = await formatBurnProof(burnProof);
        
        const balanceBefore =  tradedTokenName ? await tradedToken.balanceOf(srcToken.sender.signer.address) : await srcToken.sender.signer.getBalance();
        // submit burn proof
        const tx = await confirm(this.vault.connect(this.unshieldSender).executeWithBurnProof(...params), this.nConfirm);
        const receipt = await tx.wait();
        const tradeReturnIncAmount = await toIncDecimals(tradeReturn, tradedToken.address, useUnified.reshield);
        const emitAmount = tradedToken.address == tokenAddresses.ETH ? tradeReturn : await toIncDecimals(tradeReturn, tradedToken.address, false);

        if (useWithdraw) {
            await expect(tx).to.emit(this.vault, 'Withdraw')
                .withArgs(tradedToken.address, srcToken.sender.signer.address, tradeReturn);
            const balanceAfter = tradedTokenName ? await tradedToken.balanceOf(srcToken.sender.signer.address) : await srcToken.sender.signer.getBalance();
            expect(BN.from(balanceAfter).sub(BN.from(balanceBefore))).to.equal(BN.from(tradeReturn));
        } else {
            const { amount: redepositAmount, token: redepositToken, redepositIncAddress, itx } = await receipt.events.filter(e => e.event == 'Redeposit')[0].args;
            await expect(tx).to.emit(this.vault, 'Redeposit')
                .withArgs(tradedToken.address, redepositIncAddress, emitAmount, itx);
            // perform reshield using event
            const proveResult = await proveEth(tx.hash, inc.utils.base64Encode);
            const { result: issuingResponse } = useUnified.reshield ?
                await inc.Tx(txor).unified_shield(tradedToken.inc, proveResult.ethBlockHash, proveResult.encodedProof, proveResult.txIndex, 1, tradedToken.unify).send().catch(console.error) :
                await inc.Tx(txor).shield(tradedToken.inc, proveResult.ethBlockHash, proveResult.encodedProof, proveResult.txIndex).send().catch(console.error);
            await expect(issuingResponse).to.have.nested.property('txId').that.is.a('string').with.lengthOf(64, "sending issuing TX failed");
            let [change] = await Promise.all([txor.waitBalanceChange(useUnified.reshield ? tradedToken.unify : tradedToken.inc), txor.waitTx(issuingResponse.txId)]);
            change = BN.from(change.balance).sub(BN.from(change.oldBalance));
            await expect(change).to.equal(BN.from(tradeReturnIncAmount), 'redeposit amount mismatch');
        }
        this.tradeReturnAmount = tradeReturn;
    }
}

const burnForCallRefundedInstruction = (exchangeName, amount, srcTokenName = null, tradedTokenName = null, options = { useWithdraw: false, useUnified: { burn: true }, changeIndex: 0 }) => {
    return async function() {
        const [srcToken, tradedToken] = [srcTokenName, tradedTokenName].map(n => getBridgedIncTokenInfo(this, n));
        const txor = await inc.NewTransactor(srcToken.sender.incPrivate);
        const { useWithdraw, useUnified, changeIndex } = options;
        const withdrawAddress = useWithdraw ? srcToken.sender.signer.address : undefined;
        const incAmount = await toIncDecimals(amount, srcToken.address, useUnified.burn);
        // encode calldata
        const timestamp = ethers.utils.hexlify(ethers.utils.randomBytes(6));
        const { encodedTradeCall, hashedData, exchange, tradeReturn } = await prepareExternalCallByVault(this, exchangeName, srcToken.address, tradedToken.address, timestamp, amount);
        const externalNetworkID = 1; // test networkID: Ethereum
        // create & send burn tx
        const { result: burnResult } = await inc.Tx(txor).withTokenID(useUnified.burn ? srcToken.unify : srcToken.inc).withInfo("BURN ETHER")
            .burnForCall(...makeInvalidBurnParams(changeIndex, incAmount.toString(), srcToken.inc, externalNetworkID, exchange.address,
                encodedTradeCall, tradedToken.address), withdrawAddress)
            .send().catch(console.error);
        
        let burnBalanceChange = await txor.waitBalanceChange(useUnified.burn ? srcToken.unify : srcToken.inc);
        burnBalanceChange = BN.from(burnBalanceChange.oldBalance) - BN.from(burnBalanceChange.balance);
        await expect(burnBalanceChange).to.equal(incAmount, 'burn amount mismatch');
        let [change] = await Promise.all([txor.waitBalanceChange(useUnified.burn ? srcToken.unify : srcToken.inc), txor.waitTx(burnResult.txId)]);
        change = BN.from(change.balance).sub(BN.from(change.oldBalance));
        await expect(change).to.equal(incAmount, 'instruction refund amount mismatch');
    }
}

const burnForCallRefundedEvent = (exchangeName, amount, srcTokenName = null, tradedTokenName = null, options = { useWithdraw: false, useUnified: { burn: true, reshield: true }, changeIndex: 0 }) => {
    return async function() {
        const [srcToken, tradedToken] = [srcTokenName, tradedTokenName].map(n => getBridgedIncTokenInfo(this, n));
        const txor = await inc.NewTransactor(srcToken.sender.incPrivate);
        const { useWithdraw, useUnified, changeIndex } = options;
        const withdrawAddress = useWithdraw ? srcToken.sender.signer.address : undefined;
        const incAmount = await toIncDecimals(amount, srcToken.address, useUnified.burn);
        // encode calldata
        const timestamp = ethers.utils.hexlify(ethers.utils.randomBytes(6));
        const { encodedTradeCall, hashedData, exchange, tradeReturn } = await prepareExternalCallByVault(this, exchangeName, srcToken.address, tradedToken.address, timestamp, amount);
        const externalNetworkID = 1; // test networkID: Ethereum
        // create & send burn tx
        const incParams = makeInvalidBurnParams(changeIndex, incAmount.toString(), srcToken.inc, externalNetworkID, exchange.address,
                encodedTradeCall, tradedToken.address);
        const { result: burnResult } = await inc.Tx(txor).withTokenID(useUnified.burn ? srcToken.unify : srcToken.inc).withInfo("BURN ETHER")
            .burnForCall(...incParams, withdrawAddress)
            .send().catch(console.error);
        await expect(burnResult).to.have.nested.property('txId').that.is.a('string').with.lengthOf(64, "sending withdraw TX failed");
        await this.transactor.waitTx(burnResult.txId);
        const burnProof = await inc.rpc.unified_getBurnProof(burnResult.txId);
        let params = await formatBurnProof(burnProof);
        
        // submit burn proof
        const tx = await confirm(this.vault.connect(this.unshieldSender).executeWithBurnProof(...params), this.nConfirm);
        const receipt = await tx.wait();
        const emitAmount = srcToken.address == tokenAddresses.ETH ? await fromIncDecimals(incParams[0], srcToken.address, useUnified.burn) : incParams[0];
        const { amount: redepositAmount, token: redepositToken, redepositIncAddress, itx } = await receipt.events.filter(e => e.event == 'Redeposit')[0].args;
        await expect(tx).to.emit(this.vault, 'Redeposit')
            .withArgs(srcToken.address, redepositIncAddress, emitAmount, itx); // expect refund
        // // perform reshield using event
        const proveResult = await proveEth(tx.hash, inc.utils.base64Encode);
        const { result: issuingResponse } = useUnified.reshield ?
            await inc.Tx(txor).unified_shield(srcToken.inc, proveResult.ethBlockHash, proveResult.encodedProof, proveResult.txIndex, 1, srcToken.unify).send().catch(console.error) :
            await inc.Tx(txor).shield(srcToken.inc, proveResult.ethBlockHash, proveResult.encodedProof, proveResult.txIndex).send().catch(console.error);
        await expect(issuingResponse).to.have.nested.property('txId').that.is.a('string').with.lengthOf(64, "sending issuing TX failed");
        let [change] = await Promise.all([txor.waitBalanceChange(useUnified.reshield ? srcToken.unify : srcToken.inc), txor.waitTx(issuingResponse.txId)]);
        change = BN.from(change.balance).sub(BN.from(change.oldBalance));
        await expect(change).to.equal(BN.from(incParams[0]), 'redeposit amount mismatch');
    }
}

const burnForCallWithdrawException = (exchangeName, amount, srcTokenName = null, tradedTokenName = null, options = { useUnified: { burn: true }}) => {
    return async function() {
        const [srcToken, tradedToken] = [srcTokenName, tradedTokenName].map(n => getBridgedIncTokenInfo(this, n));
        const txor = await inc.NewTransactor(srcToken.sender.incPrivate);
        const { useWithdraw, useUnified } = options;
        const withdrawAddress = (await getInstance('IncognitoProxy')).address; // withdraw to contract that can't receive ETH
        const incAmount = await toIncDecimals(amount, srcToken.address, useUnified.burn);
        // encode calldata
        const timestamp = ethers.utils.hexlify(ethers.utils.randomBytes(6));
        const { encodedTradeCall, hashedData, exchange, tradeReturn } = await prepareExternalCallByVault(this, exchangeName, srcToken.address, tradedToken.address, timestamp, amount);
        const externalNetworkID = 1; // test networkID: Ethereum
        // create & send burn tx
        const { result: burnResult } = await inc.Tx(txor).withTokenID(useUnified.burn ? srcToken.unify : srcToken.inc).withInfo("BURN ETHER")
            .burnForCall(incAmount.toString(), srcToken.inc, externalNetworkID, exchange.address,
                encodedTradeCall, tradedToken.address, withdrawAddress)
            .send().catch(console.error);
        await expect(burnResult).to.have.nested.property('txId').that.is.a('string').with.lengthOf(64, "sending withdraw TX failed");
        await this.transactor.waitTx(burnResult.txId);
        const burnProof = await inc.rpc.unified_getBurnProof(burnResult.txId);
        let params = await formatBurnProof(burnProof);
        // submit burn proof
        const tx = await confirm(this.vault.connect(this.unshieldSender).executeWithBurnProof(...params), this.nConfirm);
        const receipt = await tx.wait();
        const emitAmount = tradedToken.address == tokenAddresses.ETH ? amount : await toIncDecimals(amount, tradedToken.address, false);
        const { amount: redepositAmount, token: redepositToken, redepositIncAddress, itx } = await receipt.events.filter(e => e.event == 'Redeposit')[0].args;
        await expect(tx).to.emit(this.vault, 'Redeposit')
            .withArgs(tradedToken.address, redepositIncAddress, emitAmount, itx); // withdraw exception -> expect fallback to reshield
    }
}

const invalidFunctionSignature = ethers.utils.arrayify(ethers.utils.keccak256(ethers.utils.toUtf8Bytes('nonexistentFunction'))).slice(0, 4);

const changeByte = (s) => {
    if (typeof(s) == 'string' && !s.startsWith('0x')) s = '0x' + s;
    let arr = ethers.utils.arrayify(s);
    const changeInd = Math.floor(Math.random() * arr.length);
    let b = ethers.utils.randomBytes(1)[0];
    while (b == arr[changeInd]) b = ethers.utils.randomBytes(1)[0];
    // console.log(`${arr[changeInd]} -> ${b}`);
    arr[changeInd] = b;
    return arr;
}

const makeInvalidBurnParams = (ind, ...params) => {
    let [incAmount, srcToken, externalNetworkID, exchangeAddress, encodedTradeCall, tradedToken] = params;
    if (ind == 0) {
        console.log('use invalid (reduced) burn amount');
        params[0] = BN.from(incAmount).sub(1000).toString();
    } else if (ind == 1) {
        console.log('use invalid burned IncTokenID');
        params[1] = ethers.utils.hexlify(changeByte(params[1])).slice(2);
    } else if (ind == 2) {
        console.log('use invalid ExternalCallAddress');
        params[3] = ethers.utils.hexlify(changeByte(params[3])).slice(2);
    } else if (ind == 3) {
        console.log('use invalid function signature for call');
        params[4] = ethers.utils.hexlify(ethers.utils.concat([invalidFunctionSignature, ethers.utils.arrayify(params[4]).slice(4)]));
    } else if (ind == 4) {
        console.log('use invalid EncodedCall inputs');
        let data = changeByte(ethers.utils.arrayify(params[4]).slice(4));
        params[4] = ethers.utils.hexlify(ethers.utils.concat([ethers.utils.arrayify(params[4]).slice(0, 4), data])).slice(2);
    } else if (ind == 5) {
        console.log('use invalid ReceiveToken address');
        params[5] = ethers.utils.hexlify(changeByte(params[5])).slice(2);
    }
    // console.log('final params', params);
    return params;
}

module.exports = {
    unified_shield,
    unified_unshield,
    unified_shieldToken,
    unified_unshieldToken,
    burnForCall,
    burnForCallRefundedInstruction,
    burnForCallRefundedEvent,
    burnForCallWithdrawException,
    invalidFunctionSignature,
    changeByte,
    makeInvalidBurnParams,
}