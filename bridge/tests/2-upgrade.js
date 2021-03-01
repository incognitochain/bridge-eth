const { expect } = require("chai");
const hre = require("hardhat");
const { deployments, ethers } = hre;

const {tokenAddresses} = require('../scripts/constants');
const {chiAddress} = require('../scripts/chi');
const { setupTest, shieldToken, unshieldToken, testMintChi } = require('./1-shield-unshield');
const {BN, getPartOf, confirm, fromIncDecimals, toIncDecimals, getInstance, getImplementation, Inc} = require('../scripts/external');

describe('Vault - Upgrade & Pausing', () => {
    const startEth = ethers.utils.parseUnits('0.00005', 'ether');
    const startToken = ethers.utils.parseUnits('0.00001', 'ether');
    before(setupTest());
    it('should upgrade to new implementation, preserving balances', async function() {
        let myTokenAddresses = this.testingTokens.map(t => t.address);
        myTokenAddresses.push(tokenAddresses.ETH);
        // proxy starts off pointing to Vault
        let currentImpl = await getImplementation(this.proxy.address);
        await expect(currentImpl).to.equal(this.implementation.address);
        // get total balance to compare later
        const balancesBeforeUpgrade = await Promise.all(myTokenAddresses.map(_tokenAddr => this.vault.balanceOf(_tokenAddr)));

        // new upgrade method (Proxy Pattern) means prevVault no longer changes
        // const data = upgradedImpl.interface.encodeFunctionData('callAfterUpgrade', [true]);
        // this proxy was initialized upon creation, so state changes that need to be done after upgrade must be via another function
        // console.log("will upgrade to", upgradedImpl.address);
        await this.proxy.connect(this.vaultAdmin).upgradeTo(this.upgradedImpl.address);
        await confirm(this.upgradedVault.callAfterUpgrade(false));
        currentImpl = await getImplementation(this.proxy.address);
        await expect(currentImpl).to.equal(this.upgradedImpl.address);

        // plug ABI of new vault to the proxy i just upgraded
        const balances = await Promise.all(myTokenAddresses.map(_tokenAddr => this.upgradedVault.balanceOf(_tokenAddr)));
        console.log('balances before upgrade', balancesBeforeUpgrade);
        console.log('after upgrade', balances);
        await expect(balances).to.deep.equal(balancesBeforeUpgrade);
        console.log('New implementation sets CHI to an EOA');
        await confirm(this.upgradedVault.setChi(this.ethGuy.signer.address));
    });
    // withdraw less than half of the deposited "Token 1"
    it('mint more CHI', testMintChi(140));
    const amount = getPartOf(startToken, 25);
    // expect to withdraw successfully but no CHI used
    it('should still unshield token with non-contract CHI address', unshieldToken(amount, 0, 20, {chiUse : false, balanceChange : true}));

    it('suppose the new implementation acts maliciously', async function() {
        // make the new vault use a non-CHI contract for CHI, as well as not paying withdrawn balance
        await confirm(this.upgradedVault.callAfterUpgrade(true));
        await confirm(this.upgradedVault.setChi(this.testingTokens[0].address));
    });
    it('unshield token success but no token payout', unshieldToken(amount, 0, 20, {chiUse : false, balanceChange : false}));

    it('should pause, rejecting all calls', async function() {
        const tokenContract = this.testingTokens[0];
        // "false" means we expect no balance change, since this implementation we just upgraded to is flawed
        // now we need to pause the proxy to avoid more damage
        await confirm(this.upgradedVault.setChi(this.chi.address));
        await confirm(this.proxy.connect(this.vaultAdmin).pause());
        // after pause, all calls should revert
        const appr = await confirm(tokenContract.connect(this.tokenGuy.signer).approve(this.upgradedVault.address, amount));
        await expect(appr).to.emit(tokenContract, 'Approval')
        .withArgs(this.tokenGuy.signer.address, this.upgradedVault.address, amount);
        await expect(this.upgradedVault.connect(this.chiFunder).depositERC20(tokenContract.address, amount, this.tokenGuy.inc)).to.be.reverted;
    });

    it('should upgrade to previous safe implementation & unpause', async function() {
        await confirm(this.proxy.connect(this.vaultAdmin).upgradeTo(this.implementation.address));
        await confirm(this.proxy.connect(this.vaultAdmin).unpause(), this.nConfirm + 1);

        const currentImpl = await getImplementation(this.proxy.address);
        await expect(currentImpl).to.equal(this.implementation.address);
    });
    // shield some token & make sure it succeeds
    it('should shield normally again', shieldToken(amount, 0));

});