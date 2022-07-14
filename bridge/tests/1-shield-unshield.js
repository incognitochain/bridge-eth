const { expect } = require("chai");
const hre = require("hardhat");
const { deployments, ethers } = hre;
const BN = ethers.BigNumber;
const {getPartOf, chooseOneFrom, confirm, fromIncDecimals, toIncDecimals, getInstance, getImplementation, generateTestIncTokenID} = require('../scripts/utils');
const { setupTest, shield, unshield, shieldToken, unshieldToken } = require('./0-basic');

describe("Vault - Shield & Unshield", () => {
    const startEth = ethers.utils.parseUnits('0.00005', 'ether');
    const startToken = ethers.utils.parseUnits('0.00001', 'ether');
    before(setupTest());

    describe('(UN)SHIELD ETHER', async function() {
        it(`should receive ${ethers.utils.formatUnits(startEth, 'gwei')} (gwei) ETH deposit`, shield(startEth));
        const amount = getPartOf(startEth, 50);
        it(`should withdraw ${ethers.utils.formatUnits(amount, 'gwei')}`, unshield(amount));
    });

    describe('(UN)SHIELD TOKEN', async function() {
        it(`should receive ${ethers.utils.formatUnits(startToken, 'gwei')} (gwei) Token1 deposit`, shieldToken(startToken, 'Token1'));
        const amount = getPartOf(startToken, 50);
        it(`should withdraw ${ethers.utils.formatUnits(amount, 'gwei')}`, unshieldToken(amount, 'Token1'));
    });
})
