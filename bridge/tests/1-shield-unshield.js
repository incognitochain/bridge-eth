const { expect } = require("chai");
const hre = require("hardhat");
const { deployments, ethers } = hre;
const BN = ethers.BigNumber;
const {getPartOf, chooseOneFrom, confirm, fromIncDecimals, toIncDecimals, getInstance, getImplementation, generateTestIncTokenID} = require('../scripts/utils');
const { setupTest, shield, unshield, unshieldRevertRecover, shieldToken, unshieldToken } = require('./0-basic');

describe("Vault - Shield & Unshield", () => {
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
            return confirm(signers[0].sendTransaction({to: accounts[1], value: ethers.utils.parseUnits("0.6969", 'gwei')}), 15)
            .then(console.log)
            .then(_ => ethers.provider.getBlockNumber())
            .then(console.log)
        });
    });

    describe('(UN)SHIELD ETHER', async function() {
        it(`should receive ${ethers.utils.formatUnits(startEth, 'gwei')} (gwei) ETH deposit`, shield(startEth));
        const amount = getPartOf(startEth, 50);
        it(`should withdraw ${ethers.utils.formatUnits(amount, 'gwei')}`, unshield(amount));
        // it(`should fail to withdraw ${ethers.utils.formatUnits(startEth, 'gwei')}, then recover`, unshieldRevertRecover(amount));
    });

    describe('(UN)SHIELD TOKEN', async function() {
        it(`should receive ${ethers.utils.formatUnits(startToken, 'gwei')} (gwei) Token1 deposit`, shieldToken(startToken, 'Token1'));
        const amount = getPartOf(startToken, 50);
        it(`should withdraw ${ethers.utils.formatUnits(amount, 'gwei')}`, unshieldToken(amount, 'Token1'));
    });

    describe('RECOVER FAILED UNSHIELD', async function() {
        const amount = BN.from(ethers.utils.parseUnits('0.00001', 'ether'));
        it(`should submit a failing unshield, then recover`, unshieldRevertRecover(amount, null));
    })
})
