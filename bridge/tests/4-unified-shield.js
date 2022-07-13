const { expect } = require("chai");
const hre = require("hardhat");
const { deployments, ethers } = hre;
const BN = ethers.BigNumber;
const { getPartOf, chooseOneFrom, confirm, fromIncDecimals, toIncDecimals, getInstance, getImplementation, generateTestIncTokenID, prepareExternalCallByVault, getBridgedIncTokenInfo } = require('../scripts/utils');
const { tokenAddresses } = require('../scripts/constants');
const { proveEth, formatBurnProof } = require('../scripts/prove');
const { inc } = require('../scripts/external');
const { setupTest, shield, unshield, shieldToken, unshieldToken } = require('./0-basic');
const { unified_shield, unified_unshield, unified_shieldToken, unified_unshieldToken, burnForCall, burnForCallRefundedInstruction, burnForCallRefundedEvent, burnForCallWithdrawException, burnForCallRevert } = require('./unified-token-helpers');

let tradeSetup = () => {
    return async function() {
        console.log(`setup for trading in network ${hre.network.name}`);
        this.kyber = await getInstance('KBNTrade');
        this.uniswap = await getInstance('UniswapV2Trade');
        this.testingExchange = await getInstance('TestingExchange');
    }
}

describe("Unified Shield & Unshield", async function() {
    const startEth = ethers.utils.parseUnits('0.00005', 'ether');
    const startToken = ethers.utils.parseUnits('0.00001', 'ether');
    before(setupTest());
    before(tradeSetup());

    describe('pApp: swap pETH for upToken1', async function() {
        it(`should receive ${ethers.utils.formatUnits(startEth, 'gwei')} (gwei) ETH deposit`, shield(startEth));
        let amount = getPartOf(startEth, 50);
        it('should burn pETH, call SWAP & reshield upToken1', burnForCall('kyber', amount, null, 'Token1', { useWithdraw: false, useUnified: { burn: false, reshield: true }}));
        it(`should unshield ${ethers.utils.formatUnits(this.tradeReturnAmount, 'gwei')} of upToken1`, unified_unshieldToken('Token1', this.tradeReturnAmount));
    });

    describe('pApp: swap upETH for pToken2', async function() {
        it(`should receive ${ethers.utils.formatUnits(startEth, 'gwei')} (gwei) ETH unified deposit`, unified_shield(startEth));
        let amount = getPartOf(startEth, 50);
        it('should burn upETH, call SWAP & reshield pToken2', burnForCall('kyber', amount, null, 'Token2', { useWithdraw: false, useUnified: { burn: true, reshield: false }}));
        it(`should unshield ${ethers.utils.formatUnits(this.tradeReturnAmount, 'gwei')} of pToken2`, unshieldToken('Token2', this.tradeReturnAmount));
    });

    describe('pApp: swap upToken2 for upETH', async function() {
        it(`should receive ${ethers.utils.formatUnits(startToken, 'gwei')} (gwei) Token2 unified deposit`, unified_shieldToken(startToken, 'Token2'));
        let amount = getPartOf(startToken, 50);
        it(`should burn upToken2, call SWAP & reshield upETH`, burnForCall('kyber', amount, 'Token2', null, { useWithdraw: false, useUnified: { burn: true, reshield: true }}));
        it(`should unshield ${ethers.utils.formatUnits(this.tradeReturnAmount, 'gwei')} of upETH`, unified_unshield(this.tradeReturnAmount));
    });

    describe('pApp: swap pToken1 for pETH', async function() {
        it(`should receive ${ethers.utils.formatUnits(startToken, 'gwei')} (gwei) Token1 deposit`, shieldToken(startToken, 'Token1'));
        let amount = getPartOf(startToken, 50);
        it(`should burn pToken1, call SWAP & reshield pETH`, burnForCall('kyber', amount, 'Token1', null, { useWithdraw: false, useUnified: { burn: false, reshield: false }}));
        it(`should unshield ${ethers.utils.formatUnits(this.tradeReturnAmount, 'gwei')} of pETH`, unshield(this.tradeReturnAmount));
    });

    describe('pApp: swap pETH for Token1', async function() {
        it(`should receive ${ethers.utils.formatUnits(startEth, 'gwei')} (gwei) ETH deposit`, shield(startEth));
        let amount = getPartOf(startEth, 50);
        it('should burn pETH, call SWAP & withdraw Token1', burnForCall('kyber', amount, null, 'Token1', { useWithdraw: true, useUnified: { burn: false }}));
    });

    describe('pApp: swap upETH for Token2', async function() {
        it(`should receive ${ethers.utils.formatUnits(startEth, 'gwei')} (gwei) ETH unified deposit`, unified_shield(startEth));
        let amount = getPartOf(startEth, 50);
        it('should burn upETH, call SWAP & withdraw Token2', burnForCall('kyber', amount, null, 'Token2', { useWithdraw: true, useUnified: { burn: true }}));
    });

    describe('pApp: swap pToken1 for ETH', async function() {
        it(`should receive ${ethers.utils.formatUnits(startToken, 'gwei')} (gwei) Token1 deposit`, shieldToken(startToken, 'Token1'));
        let amount = getPartOf(startToken, 50);
        it(`should burn pToken1, call SWAP & withdraw ETH`, burnForCall('kyber', amount, 'Token1', null, { useWithdraw: true, useUnified: { burn: false }}));
    });

    describe('pApp: swap upToken2 for ETH', async function() {
        it(`should receive ${ethers.utils.formatUnits(startToken, 'gwei')} (gwei) Token2 unified deposit`, unified_shieldToken(startToken, 'Token2'));
        let amount = getPartOf(startToken, 50);
        it(`should burn upToken2, call SWAP & withdraw ETH`, burnForCall('kyber', amount, 'Token2', null, { useWithdraw: true, useUnified: { burn: true }}));
    });

    describe('prepare funds for recovery cases', async function() {
        it(`should shield 20x ETH`, unified_shield(startEth.mul(20)));
        it(`should shield 20x Token1`, unified_shieldToken(startToken.mul(20), 'Token1'));
    });

    describe('pApp: burn upETH -> refund', async function() {
        it(`refund event: wrong burn amount`, burnForCallRefundedEvent('kyber', startEth, null, 'Token1', { useWithdraw: false, useUnified: { burn: true, reshield: true }, changeIndex: 0 }));
        it(`refund instruction: wrong IncTokenID`, burnForCallRefundedInstruction('kyber', startEth, null, 'Token1', { useWithdraw: true, useUnified: { burn: true, reshield: true }, changeIndex: 1 }));
        it(`refund event: random invalid exchange address`, burnForCallRefundedEvent('kyber', startEth, null, 'Token1', { useWithdraw: false, useUnified: { burn: true, reshield: true }, changeIndex: 2 }));
        it(`refund event: call nonexistent function`, burnForCallRefundedEvent('kyber', startEth, null, 'Token1', { useWithdraw: true, useUnified: { burn: true, reshield: true }, changeIndex: 3 }));
        it(`refund event: call with changed inputs`, burnForCallRefundedEvent('kyber', startEth, null, 'Token1', { useWithdraw: false, useUnified: { burn: true, reshield: true }, changeIndex: 4 }));
        it(`refund event: wrong ReceiveToken`, burnForCallRefundedEvent('kyber', startEth, null, 'Token1', { useWithdraw: true, useUnified: { burn: true, reshield: true }, changeIndex: 5 }));
        // it(`revert: submit call to wrong network`, burnForCallRevert('kyber', startEth, null, 'Token1', { useWithdraw: false, useUnified: { burn: true }, changeIndex: 6 }));
    });

    describe('pApp: burn upToken1 -> refund', async function() {
        it(`refund event: wrong burn amount`, burnForCallRefundedEvent('kyber', startToken, 'Token1', null, { useWithdraw: false, useUnified: { burn: true, reshield: true }, changeIndex: 0 }));
        it(`refund instruction: wrong IncTokenID`, burnForCallRefundedInstruction('kyber', startToken, 'Token1', null, { useWithdraw: true, useUnified: { burn: true, reshield: true }, changeIndex: 1 }));
        it(`refund event: random invalid exchange address`, burnForCallRefundedEvent('kyber', startToken, 'Token1', null, { useWithdraw: false, useUnified: { burn: true, reshield: true }, changeIndex: 2 }));
        it(`refund event: call nonexistent function`, burnForCallRefundedEvent('kyber', startToken, 'Token1', null, { useWithdraw: true, useUnified: { burn: true, reshield: true }, changeIndex: 3 }));
        it(`refund event: call with changed inputs`, burnForCallRefundedEvent('kyber', startToken, 'Token1', null, { useWithdraw: false, useUnified: { burn: true, reshield: true }, changeIndex: 4 }));
        it(`refund event: wrong ReceiveToken`, burnForCallRefundedEvent('kyber', startToken, 'Token1', null, { useWithdraw: true, useUnified: { burn: true, reshield: true }, changeIndex: 5 }));
        it(`reshield event: recover from withdraw exception`, burnForCallWithdrawException('kyber', startToken, 'Token1', null, { useUnified: { burn: true }}));
    });
})
