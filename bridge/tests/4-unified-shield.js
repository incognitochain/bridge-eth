const { expect } = require("chai");
const hre = require("hardhat");
const { deployments, ethers } = hre;
const BN = ethers.BigNumber;
const { getPartOf, chooseOneFrom, confirm, fromIncDecimals, toIncDecimals, getInstance, getImplementation, generateTestIncTokenID, prepareExternalCallByVault, getBridgedIncTokenInfo } = require('../scripts/utils');
const { tokenAddresses } = require('../scripts/constants');
const { proveEth, formatBurnProof } = require('../scripts/prove');
const { inc } = require('../scripts/external');
const { setupTest, shield, unshield, shieldToken, unshieldToken } = require('./0-basic');

let tradeSetup =  () => {
    return async function() {
        console.log(`setup for trading in network ${hre.network.name}`);
        this.kyber = await getInstance('KBNTrade');
        this.uniswap = await getInstance('UniswapV2Trade');
        this.testingExchange = await getInstance('TestingExchange');
    }
}

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
        let change = await txor.waitBalanceChange(tokenAddresses.unify[tokenAddresses.ETH]);
        change = BN.from(change.balance) - BN.from(change.oldBalance);
        change = await fromIncDecimals(change, tokenAddresses.ETH);
        await expect(change).to.gte(ethIn, 'issuing amount mismatch');
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
        const incAmount = await toIncDecimals(_amount, tokenContract.address);
        const tx = await confirm(this.vault.connect(this.tokenGuy.signer).depositERC20(tokenContract.address, _amount, this.tokenGuy.inc), this.nConfirm);
        await expect(tx).to.emit(this.vault, 'Deposit')
            .withArgs(ethers.utils.getAddress(tokenContract.address), this.tokenGuy.inc, incAmount);
        const proveResult = await proveEth(tx.hash, inc.utils.base64Encode);
        const txor = await inc.NewTransactor(this.tokenGuy.incPrivate);
        const { result: issuingResponse } = await inc.Tx(txor).unified_shield(tokenContract.inc, proveResult.ethBlockHash, proveResult.encodedProof, proveResult.txIndex, 1, tokenContract.unify).send().catch(console.error);
        await expect(issuingResponse).to.have.nested.property('txId').that.is.a('string').with.lengthOf(64, "sending issuing TX failed");
        let change = await txor.waitBalanceChange(tokenContract.unify);
        await expect(change).to.gte(incAmount, 'token issuing amount mismatch');
    }
}

let unified_unshieldToken = (_amount, _tokenName, _expectations = { balanceChange: true }) => {
    return async function() {
        const tokenContract = getBridgedIncTokenInfo(this, _tokenName);
        // switch to Incognito's decimal
        const incAmount = await toIncDecimals(_amount, tokenContract.address);
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

const burnUnifiedForCall = (exchangeName, amount, srcTokenName = null, tradedTokenName = null) => {
    return async function() {
        const [srcToken, tradedToken] = [srcTokenName, tradedTokenName].map(n => getBridgedIncTokenInfo(this, n));
        const txor = await inc.NewTransactor(srcToken.sender.incPrivate);
        const incAmount = await toIncDecimals(amount, srcToken.address);
        const timestamp = ethers.utils.hexlify(ethers.utils.randomBytes(6));
        const { encodedTradeCall, hashedData, exchange, tradeReturn } = await prepareExternalCallByVault(this, exchangeName, srcToken.address, tradedToken.address, timestamp, amount);
        const externalNetworkID = 1; // test networkID: Ethereum
        const { result: burnResult } = await inc.Tx(txor).withTokenID(srcToken.unify).withInfo("BURN ETHER").burnForCall(incAmount.toString(), srcToken.inc, 1, exchange.address, encodedTradeCall, tradedToken.address).send().catch(console.error);
        await expect(burnResult).to.have.nested.property('txId').that.is.a('string').with.lengthOf(64, "sending withdraw TX failed");
        await this.transactor.waitTx(burnResult.txId);
        const burnProof = await inc.rpc.unified_getBurnProof(burnResult.txId);
        let params = await formatBurnProof(burnProof);

        const balanceBefore = await srcToken.sender.signer.getBalance();
        const tx = await confirm(this.vault.connect(this.unshieldSender).executeWithBurnProof(...params), this.nConfirm);
        const receipt = await tx.wait();
        const tradeReturnIncAmount = await toIncDecimals(tradeReturn, tradedToken.address);
        const emitAmount = tradedToken.address == tokenAddresses.ETH ? tradeReturn : tradeReturnIncAmount;
        const { amount: redepositAmount, token: redepositToken, redepositIncAddress, itx } = await receipt.events.filter(e => e.event=='Redeposit')[0].args;
        await expect(tx).to.emit(this.vault, 'Redeposit')
        .withArgs(tradedToken.address, redepositIncAddress, emitAmount, itx);
        
       const proveResult = await proveEth(tx.hash, inc.utils.base64Encode);
        const { result: issuingResponse } = await inc.Tx(txor).shield(tradedToken.inc, proveResult.ethBlockHash, proveResult.encodedProof, proveResult.txIndex).send().catch(console.error);
        await expect(issuingResponse).to.have.nested.property('txId').that.is.a('string').with.lengthOf(64, "sending issuing TX failed");
        let change = await txor.waitBalanceChange(tradedToken.inc);
        change = BN.from(change.balance) - BN.from(change.oldBalance);
        await expect(change.toString()).to.equal(tradeReturnIncAmount.toString(), 'redeposit amount mismatch');
    }
}

const burnForCall = (exchangeName, amount, srcTokenName = null, tradedTokenName = null, useUnified = { burn: true, reshield: true}) => {
    return async function() {
        const [srcToken, tradedToken] = [srcTokenName, tradedTokenName].map(n => getBridgedIncTokenInfo(this, n));
        const txor = await inc.NewTransactor(srcToken.sender.incPrivate);
        const incAmount = await toIncDecimals(amount, srcToken.address);
        const timestamp = ethers.utils.hexlify(ethers.utils.randomBytes(6));
        const { encodedTradeCall, hashedData, exchange, tradeReturn } = await prepareExternalCallByVault(this, exchangeName, srcToken.address, tradedToken.address, timestamp, amount);
        const externalNetworkID = 1; // test networkID: Ethereum
        const { result: burnResult } = await inc.Tx(txor).withTokenID(useUnified.burn ? srcToken.unify : srcToken.inc).withInfo("BURN ETHER").burnForCall(incAmount.toString(), srcToken.inc, 1, exchange.address, encodedTradeCall, tradedToken.address).send().catch(console.error);
        await expect(burnResult).to.have.nested.property('txId').that.is.a('string').with.lengthOf(64, "sending withdraw TX failed");
        await this.transactor.waitTx(burnResult.txId);
        const burnProof = await inc.rpc.unified_getBurnProof(burnResult.txId);
        let params = await formatBurnProof(burnProof);

        const balanceBefore = await srcToken.sender.signer.getBalance();
        const tx = await confirm(this.vault.connect(this.unshieldSender).executeWithBurnProof(...params), this.nConfirm);
        const receipt = await tx.wait();
        const tradeReturnIncAmount = await toIncDecimals(tradeReturn, tradedToken.address);
        const emitAmount = tradedToken.address == tokenAddresses.ETH ? tradeReturn : tradeReturnIncAmount;
        const { amount: redepositAmount, token: redepositToken, redepositIncAddress, itx } = await receipt.events.filter(e => e.event=='Redeposit')[0].args;
        await expect(tx).to.emit(this.vault, 'Redeposit')
        .withArgs(tradedToken.address, redepositIncAddress, emitAmount, itx);
        
        const proveResult = await proveEth(tx.hash, inc.utils.base64Encode);
        const { result: issuingResponse } = useUnified.reshield ? 
            await inc.Tx(txor).shield(tradedToken.inc, proveResult.ethBlockHash, proveResult.encodedProof, proveResult.txIndex).send().catch(console.error)
            : await inc.Tx(txor).unified_shield(tradedToken.inc, proveResult.ethBlockHash, proveResult.encodedProof, proveResult.txIndex, 1, tradedToken.unify).send().catch(console.error);
        await expect(issuingResponse).to.have.nested.property('txId').that.is.a('string').with.lengthOf(64, "sending issuing TX failed");
        let change = await txor.waitBalanceChange(tradedToken.inc);
        change = BN.from(change.balance) - BN.from(change.oldBalance);
        await expect(change).to.gte(tradeReturnIncAmount, 'redeposit amount mismatch');
        this.tradeReturnAmount = tradeReturn;
    }
}

describe("Unified Shield & Unshield", async function() {
    const startEth = ethers.utils.parseUnits('0.00005', 'ether');
    const startToken = ethers.utils.parseUnits('0.00001', 'ether');
    before(setupTest());
    before(tradeSetup());

    describe.skip('pApp: swap pETH for upToken1', async function() {
        it(`should receive ${ethers.utils.formatUnits(startEth, 'gwei')} (gwei) ETH deposit`, shield(startEth));
        let amount = getPartOf(startEth, 50);
        it('should burn pETH, call SWAP & reshield Token1', burnForCall('kyber', amount, null, 'Token1', { burn: false, reshield: true }));
        it(`should unshield ${ethers.utils.formatUnits(this.tradeReturnAmount, 'gwei')} of UPT1`, unified_unshieldToken('Token1', this.tradeReturnAmount));
    });

    describe('pApp: swap upETH for pToken3', async function() {
        it(`should receive ${ethers.utils.formatUnits(startEth, 'gwei')} (gwei) ETH unified deposit`, unified_shield(startEth));
        let amount = getPartOf(startEth, 50);
        it('should burn upETH, call SWAP & reshield pToken3', burnForCall('kyber', amount, null, 'Token3', { burn: true, reshield: false }));
        it(`should unshield ${ethers.utils.formatUnits(this.tradeReturnAmount, 'gwei')} of UPT3`, unshieldToken('Token3', this.tradeReturnAmount));
    });

    describe('pApp: swap upToken2 for upETH', async function() {
        it(`should receive ${ethers.utils.formatUnits(startToken, 'gwei')} (gwei) Token2 unified deposit`, unified_shieldToken(startToken, 'Token2'));
        let amount = getPartOf(startToken, 50);
        it(`should burn upToken2, call SWAP & reshield upETH`, burnForCall('kyber', amount, 'Token2', null, { burn: true, reshield: true }));
        it(`should unshield ${ethers.utils.formatUnits(this.tradeReturnAmount, 'gwei')} of upETH`, unified_unshield(this.tradeReturnAmount));
    });

    describe('pApp: swap pToken1 for pETH', async function() {
        it(`should receive ${ethers.utils.formatUnits(startToken, 'gwei')} (gwei) Token1 deposit`, unified_shieldToken(startToken, 'Token1'));
        let amount = getPartOf(startToken, 50);
        it(`should burn pToken1, call SWAP & reshield pETH`, burnForCall('kyber', amount, 'Token1', null, { burn: false, reshield: false }));
        it(`should unshield ${ethers.utils.formatUnits(this.tradeReturnAmount, 'gwei')} of pETH`, unshield(this.tradeReturnAmount));
    });
})