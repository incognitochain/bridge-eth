const { expect } = require("chai");
const hre = require("hardhat");
const { deployments, ethers } = hre;
const BN = ethers.BigNumber;
const {tokenAddresses} = require('../scripts/constants');
const { setupTest, shieldToken, unshieldToken } = require('./0-basic');
const {getPartOf, confirm, fromIncDecimals, toIncDecimals, getInstance, getImplementation} = require('../scripts/utils');
const { Inc } = require('../scripts/external');

describe('Vault - Upgrade & Pausing', () => {
    const startEth = ethers.utils.parseUnits('0.00005', 'ether');
    const startToken = ethers.utils.parseUnits('0.00001', 'ether');
    before(setupTest());
    const amount = getPartOf(startToken, 25);
    it('should shield some ERC20 token', shieldToken(startToken, 'Token1'));

    it('should have correct storageLayoutVersion', async function() {
        const v = await this.vault.storageLayoutVersion();
        expect(v).to.equal(2);
        await expect(this.vault.upgradeVaultStorageLayout("")).to.be.reverted;
    })

    it('should upgrade to future implementation, preserving balances', async function() {
        let myTokenAddresses = this.testingTokens.map(t => t.address);
        myTokenAddresses.push(tokenAddresses.ETH);
        // proxy starts off pointing to Vault
        let currentImpl = await getImplementation(this.proxy.address);
        await expect(currentImpl).to.equal(this.implementation.address);
        // get total balance to compare later
        const balancesBeforeUpgrade = await Promise.all(myTokenAddresses.map(_tokenAddr => this.vault.balanceOf(_tokenAddr)));

        const upgradeData = this.upgradedImpl.interface.encodeFunctionData('upgradeVaultStorageLayout', []);
        console.log('will upgrade proxy to new implementation with params', this.upgradedImpl.address, this.vaultAdmin.address, upgradeData);
        await this.proxy.connect(this.vaultAdmin).upgradeToAndCall(this.upgradedImpl.address, upgradeData);
        await confirm(this.upgradedVault.setAllowWithdraw(true));

        currentImpl = await getImplementation(this.proxy.address);
        await expect(currentImpl).to.equal(this.upgradedImpl.address);
        const v = await this.upgradedVault.storageLayoutVersion();
        expect(v).to.equal(3);
        await expect(this.upgradedVault.upgradeVaultStorageLayout()).to.be.reverted;

        // plug ABI of new vault to the proxy i just upgraded
        const balances = await Promise.all(myTokenAddresses.map(_tokenAddr => this.upgradedVault.balanceOf(_tokenAddr)));
        console.log('balances before upgrade', balancesBeforeUpgrade.map(b => b.toString()));
        console.log('after upgrade', balances.map(b => b.toString()));
        await expect(balances).to.deep.equal(balancesBeforeUpgrade);

    });
    // withdraw less than half of the deposited "Token 1"

    it('should unshield token with this new implementation', unshieldToken(amount, 'Token1', {balanceChange : true}));

    it('suppose the new implementation acts maliciously', async function() {
        await confirm(this.upgradedVault.setAllowWithdraw(false));
    });
    it('unshield token success but no token payout', unshieldToken(amount, 'Token1', {balanceChange : false}));

    it('should pause, rejecting all calls', async function() {
        const tokenContract = this.testingTokens[0];
        // "false" means we expect no balance change, since this implementation we just upgraded to is flawed
        // now we need to pause the proxy to avoid more damage
        await confirm(this.proxy.connect(this.vaultAdmin).pause());
        // after pause, all calls should revert
        const appr = await confirm(tokenContract.connect(this.tokenGuy.signer).approve(this.upgradedVault.address, amount));
        await expect(appr).to.emit(tokenContract, 'Approval')
        .withArgs(this.tokenGuy.signer.address, this.upgradedVault.address, amount);
        try {
            await expect(this.upgradedVault.connect(this.unshieldSender).estimateGas.depositERC20(tokenContract.address, amount, this.tokenGuy.inc)).to.be.reverted;
        } catch(e) {
            if (hre.network.name == 'localhost') throw e;
            else {
                console.error("Skipping known issue: testnet provider does not register Transaction Revert");
                console.error(e);
            }
        }
    });

    it('should upgrade to previous safe implementation & unpause', async function() {
        await confirm(this.proxy.connect(this.vaultAdmin).upgradeTo(this.implementation.address));
        await confirm(this.proxy.connect(this.vaultAdmin).unpause(), this.nConfirm + 1);

        const currentImpl = await getImplementation(this.proxy.address);
        await expect(currentImpl).to.equal(this.implementation.address);

        const v = await this.vault.storageLayoutVersion();
        expect(v).to.equal(3);
        await expect(this.vault.upgradeVaultStorageLayout("")).to.be.reverted;
    });
    // shield some token & make sure it succeeds
    it('should shield normally again', shieldToken(amount, 'Token1'));

});