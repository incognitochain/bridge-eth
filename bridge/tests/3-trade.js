const { expect } = require("chai");
const hre = require("hardhat");
const { deployments, ethers } = hre;
const BN = ethers.BigNumber;
const {tokenAddresses} = require('../scripts/constants');
const { setupTest, shield, shieldToken, unshield, unshieldToken } = require('./0-basic');
const {getPartOf, chooseOneFrom, confirm, fromIncDecimals, toIncDecimals, getInstance, getImplementation, generateTestIncTokenID, getReqWithdrawSignMessage, getExecuteSignMessage, prepareExternalCallByVault, getBridgedIncTokenInfo, getEmittedDepositNumber} = require('../scripts/utils');
const { inc } = require('../scripts/external');
const { proveEth, formatBurnProof } = require('../scripts/prove');

let tradeSetup =  () => {
    return async function() {
        console.log(`setup for trading in network ${hre.network.name}`);
        this.kyber = await getInstance('KBNTrade');
        this.uniswap = await getInstance('UniswapV2Trade');
        // this.testingExchange = await getInstance('TestingExchange');
    }
}

// blank token name means ether
// move some Ether or ERC20 token from Incognito chain to Vault contract (using submitBurnProof())
let moveToContract = (ethOut, privKey, tokenName) => {
    return async function() {
        const myToken = getBridgedIncTokenInfo(this, tokenName);
        const actionSigner = new ethers.Wallet(privKey);
        const incAmount = await toIncDecimals(ethOut, myToken.address);
        const txor = await inc.NewTransactor(myToken.sender.incPrivate);
        const { result: burnResult } = await inc.Tx(txor).withTokenID(myToken.inc).withInfo("BURN ETHER TO CONTRACT").burningRequest(incAmount.toString(), actionSigner.address, inc.constants.BurningRequestToSCMeta).send();
        await expect(burnResult).to.have.nested.property('txId').that.is.a('string').with.lengthOf(64, "sending withdraw TX failed");
        await this.transactor.waitTx(burnResult.txId);
        const burnProof = await inc.rpc.getBurnProof(burnResult.txId, true);
        let params = await formatBurnProof(burnProof);
        const balanceBefore = await this.vault.getDepositedBalance(myToken.address, actionSigner.address);
        const tx = await confirm(this.vault.connect(this.unshieldSender).submitBurnProof(...params, {gasLimit: 300000}));
        const receipt = await tx.wait();

        const balance = await this.vault.getDepositedBalance(myToken.address, actionSigner.address);
        console.log(`After submitting burn proof, ${myToken.address} balance of ${actionSigner.address} is ${balance.toString()}`);
        // calculate the ETH balance increase after withdrawal
        await expect(BN.from(balance).sub(BN.from(balanceBefore))).to.equal(ethOut);
    }
}

// move some Ether or ERC20 token to Incognito chain from Vault contract (using requestWithdraw())
let moveToIncognitoChain = (amount, privKey, tokenName) => {
    return async function() {
        const myToken = getBridgedIncTokenInfo(this, tokenName);
        const actionSigner = new ethers.Wallet(privKey);
        // console.log(actionSigner._signingKey());
        console.log(`address for request-withdraw : ${actionSigner.address}`);
        let timestamp = ethers.utils.hexlify(ethers.utils.randomBytes(6));
        let incAmount;
        if (amount) {
            incAmount = await toIncDecimals(amount, myToken.address);
        } else {
            // 0 means move all
            amount = await this.vault.getDepositedBalance(myToken.address, actionSigner.address);
            incAmount = await toIncDecimals(amount, myToken.address);
        }
        const { data, hashedData, compare } = await getReqWithdrawSignMessage(this.vaultHelper, myToken.address, timestamp, amount, myToken.sender.inc);
        const signatureObj = await actionSigner._signingKey().signDigest(ethers.utils.arrayify(hashedData));
        const signature = customJoinSignature(signatureObj);
        // console.log(`sign ${hashedData} -> ${signature}`);
        const recoveredAddress = ethers.utils.recoverAddress(ethers.utils.arrayify(hashedData), signature);
        // console.log(`recover into ${recoveredAddress} vs...`);
        await expect(recoveredAddress).to.equal(actionSigner.address)
        // as per Incognito specs, deposit and wait more than 15 blocks
        const tx = await confirm(this.vault.connect(this.unshieldSender).requestWithdraw(myToken.sender.inc, myToken.address, amount, signature, timestamp), this.nConfirm);
        // in this test case, this call should emit only one event
        // console.log(`Deposit : ${tx.hash} => achieved ${nConfirm + 1} confirmations`);
        await expect(tx).to.emit(this.vault, 'Deposit')
        .withArgs(ethers.utils.getAddress(myToken.address), myToken.sender.inc, await getEmittedDepositNumber(this, tokenName, amount));
        const proveResult = await proveEth(tx.hash, inc.utils.base64Encode);
        const txor = await inc.NewTransactor(myToken.sender.incPrivate);
        const { result: issuingResponse } = await inc.Tx(txor).shield(myToken.inc, proveResult.ethBlockHash, proveResult.encodedProof, proveResult.txIndex).send();
        await expect(issuingResponse).to.have.nested.property('txId').that.is.a('string').with.lengthOf(64, "sending issuing TX failed");
        let change = await txor.waitBalanceChange(myToken.inc);

        await expect(incAmount).to.equal(change.balance - change.oldBalance, 'issuing amount mismatch');
    }
}

// need private key to sign the action
let trade = (exchangeName, amount, privKey, srcTokenName, tradedTokenName, useWorstRate = false) => {
    return async function() {
        const actionSigner = new ethers.Wallet(privKey);
        const [srcToken, tradedToken] = [srcTokenName, tradedTokenName].map(n => getBridgedIncTokenInfo(this, n));
        // console.log(srcTokenName, tradedTokenName);
        // console.log(srcToken, tradedToken);
        console.log(`address for execute() : ${actionSigner.address}`);
        let timestamp = ethers.utils.hexlify(ethers.utils.randomBytes(6));
        let incAmount = await toIncDecimals(amount, srcToken.address);
        const {encodedTradeCall, hashedData, exchange, tradeReturn} = await prepareExternalCallByVault(this, exchangeName, srcToken.address, tradedToken.address, timestamp, amount, useWorstRate);
        const signatureObj = await actionSigner._signingKey().signDigest(ethers.utils.arrayify(hashedData));
        const signature = customJoinSignature(signatureObj);
        // console.log(`sign ${hashedData} -> ${signature}`);
        const recoveredAddress = ethers.utils.recoverAddress(ethers.utils.arrayify(hashedData), signature);
        expect(recoveredAddress).to.equal(actionSigner.address);

        const balanceBefore = await this.vault.getDepositedBalance(tradedToken.address, actionSigner.address);
        // as per Incognito specs, deposit and wait more than 15 blocks
        // console.log(`before execute call, ${tradedToken.address} balance of ${actionSigner.address} is ${balanceBefore.toString()}`);
        // const temp = await this.vault.getDepositedBalance(srcToken.address, actionSigner.address);
        // console.log(`before execute call, ${srcToken.address} balance of ${actionSigner.address} is ${temp.toString()}`);
        const tx = await confirm(this.vault.connect(this.unshieldSender).execute(srcToken.address, amount, tradedToken.address, exchange.address, encodedTradeCall, timestamp, signature), this.nConfirm);
        const balance = await this.vault.getDepositedBalance(tradedToken.address, actionSigner.address);
        // console.log(tradeReturn, balance, balanceBefore);
        await expect(tradeReturn).to.lte(balance.sub(balanceBefore), 'traded amount mismatch');
    }
}

// match the format in crypto.Sign from Geth, which marshals recoveryId to byte without adding 27
let customJoinSignature = (signature) => {
    signature = ethers.utils.splitSignature(signature);
    return ethers.utils.hexlify(ethers.utils.concat([
         signature.r,
         signature.s,
         (signature.recoveryParam ? "0x01": "0x00")
    ]));
}

describe('Trading Test', async function() {
    const startEth = BN.from(ethers.utils.parseUnits('0.00005', 'ether'));
    const startToken = BN.from(ethers.utils.parseUnits('0.00001', 'ether'));
    before(setupTest());
    before(tradeSetup())

    it.skip(`should check ABI encoding of signed data`, async function(){
        // const encodedPreSignData = this.vault.interface.encodeFunctionData('_buildPreSignData', [0, this.testingTokens[0].address, '0x1354', 502020]);
        // const psd = this.vault.interface.decodeFunctionData('_buildPreSignData', encodedPreSignData);
        {
            const { hashedData } = await getReqWithdrawSignMessage(this.vaultHelper, this.testingTokens[1].address, '0x1414', 3000, 'SOME ADDRESS');
            const psdCompare = await this.vaultHelper.newPreSignData(1, this.testingTokens[1].address, '0x1414', 3000);
            const compare = await this.vaultHelper._buildSignRequestWithdraw(psdCompare, 'SOME ADDRESS');
            console.log(hashedData, compare);
            await expect(hashedData).to.have.length.above(10);
            await expect(hashedData).to.equal(compare);
        }
        {
            const { hashedData } = await getExecuteSignMessage(this.vaultHelper, this.testingTokens[0].address, '0x1412', 32000, this.ethGuy.signer.address, this.tokenGuy.signer.address, '0xADADAD');
            const psdCompare = await this.vaultHelper.newPreSignData(0, this.testingTokens[0].address, '0x1412', 32000);
            const compare = await this.vaultHelper._buildSignExecute(psdCompare, this.ethGuy.signer.address, this.tokenGuy.signer.address, '0xADADAD');
            console.log(hashedData, compare);
            await expect(hashedData).to.have.length.above(10);
            await expect(hashedData).to.equal(compare);
        }
    });
    describe('Move Ether & ERC20 token around', async function() {
        // it(`should receive ${ethers.utils.formatUnits(startEth, 'gwei')} (gwei) ETH deposit`, shield(startEth));
        let amount = getPartOf(startEth, 50);
        let mySigningKey = ethers.utils.randomBytes(32);
        it(`should move ${ethers.utils.formatUnits(amount, 'gwei')} to contract`, moveToContract(amount, mySigningKey));
        it(`should move ${ethers.utils.formatUnits(amount, 'gwei')} back to Incognito chain`, moveToIncognitoChain(amount, mySigningKey));
        it(`should withdraw ${ethers.utils.formatUnits(startEth, 'gwei')}`, unshield(amount));

        it(`should receive ${ethers.utils.formatUnits(startEth, 'gwei')} token deposit`, shieldToken(startToken, 'Token1'));
        amount = getPartOf(startToken, 50);
        mySigningKey = ethers.utils.randomBytes(32);
        it(`should move ${ethers.utils.formatUnits(amount, 'gwei')} to contract`, moveToContract(amount, mySigningKey, 'Token1'));
        it(`should move ${ethers.utils.formatUnits(amount, 'gwei')} back to Incognito chain`, moveToIncognitoChain(amount, mySigningKey, 'Token1'));
        it(`should withdraw ${ethers.utils.formatUnits(startToken, 'gwei')}`, unshieldToken(amount, 'Token1'));
    })

    describe('Trade Ether', async function() {
        it(`should receive ${ethers.utils.formatUnits(startEth, 'gwei')} (gwei) ETH deposit`, shield(startEth));
        let amount = getPartOf(startEth, 50);
        const mySigningKey = ethers.utils.randomBytes(32);
        it(`should move ${ethers.utils.formatUnits(startEth, 'gwei')} to contract`, moveToContract(startEth, mySigningKey));
        it(`should trade ${ethers.utils.formatUnits(amount, 'gwei')} ETH for token2 via Kyber`, trade('kyber', amount, mySigningKey, null, 'Token2'));
        amount = getPartOf(startEth, 50);
        it(`should trade ${ethers.utils.formatUnits(amount, 'gwei')} ETH for token3 via Uniswap`, trade('uniswap', amount, mySigningKey, null, 'Token3'));
        it(`should move all token3 back to Incognito chain`, moveToIncognitoChain(0, mySigningKey, 'Token3'));
    })
    describe('Trade Token', async function() {
        it(`should receive ${ethers.utils.formatUnits(startToken, 'gwei')} (gwei) token1 deposit`, shieldToken(startToken, 'Token1'));
        let amount = getPartOf(startToken, 25);
        const mySigningKey = ethers.utils.randomBytes(32);
        it(`should move ${ethers.utils.formatUnits(startToken, 'gwei')} to contract`, moveToContract(startToken, mySigningKey, 'Token1'));
        // this trade tends to fail, so we use worstRate & trade in small increments for token
        it(`should trade ${ethers.utils.formatUnits(amount, 'gwei')} token1 for token2 via Kyber`, trade('kyber', amount, mySigningKey, 'Token1', 'Token2', true));
        amount = getPartOf(startToken, 25);
        it(`should trade ${ethers.utils.formatUnits(amount, 'gwei')} token1 for token3 via Uniswap`, trade('uniswap', amount, mySigningKey, 'Token1', 'Token3'));
        amount = getPartOf(startToken, 25);
        it(`should trade ${ethers.utils.formatUnits(amount, 'gwei')} token1 for ether via Kyber`, trade('kyber', amount, mySigningKey, 'Token1', null, true));
        amount = getPartOf(startToken, 25);
        it(`should trade ${ethers.utils.formatUnits(amount, 'gwei')} token1 for ether via Uniswap`, trade('uniswap', amount, mySigningKey, 'Token1', null));
        it(`should move all ether back to Incognito chain`, moveToIncognitoChain(0, mySigningKey));
        // it(`should move all token2 back to Incognito chain`, moveToIncognitoChain(0, mySigningKey, 'Token2'));
        // it(`should move all token3 back to Incognito chain`, moveToIncognitoChain(0, mySigningKey, 'Token3'));
    })
})

module.exports = {
    tradeSetup,
}
