const { expect } = require("chai");
const hre = require("hardhat");
const { deployments, ethers } = hre;
const BN = ethers.BigNumber;
const {tokenAddresses} = require('../scripts/constants');
const { setupTest, shieldToken, unshieldToken } = require('./0-basic');
const {getPartOf, confirm, fromIncDecimals, toIncDecimals, getInstance, getImplementation, getCodeSlot, getBridgedIncTokenInfo} = require('../scripts/utils');
const { inc } = require('../scripts/external');
const { proveEth, formatBurnProof } = require('../scripts/prove');

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

describe('Vault - Upgrade & Pausing', async function(){
    const startEth = ethers.utils.parseUnits('0.00005', 'ether');
    const startToken = ethers.utils.parseUnits('0.00001', 'ether');
    before(setupTest());
    
    describe('Move some funds to contract', async function() {
        let mySigningKey = ethers.utils.randomBytes(32);
        const amount = getPartOf(startToken, 25);
        it.skip('should have shielded Ether', shield(startEth, 'Token1'));
        it.skip('should have shielded ERC20 token', shieldToken(startToken, 'Token1'));
        it(`should move ${ethers.utils.formatUnits(amount, 'gwei')} to contract`, moveToContract(amount, mySigningKey));
        it(`should move ${ethers.utils.formatUnits(amount, 'gwei')} to contract`, moveToContract(amount, mySigningKey, 'Token1'));
    })    

    const amount = getPartOf(startToken, 25);
    describe('Vault can upgrade implementation', async function() {
        it('should have valid state', async function() {
            const v = await this.vault.connect(this.ethUser);
            expect(await v.isInitialized()).to.equal(true);
            this.previousVaultAddress = (await deployments.get('PreviousVersionVault')).address;
            this.latestVaultAddress = (await deployments.get('Vault')).address;
            await expect(BN.from(this.previousVaultAddress)).to.not.equal(BN.from(0), 'need non-zero previous vault');
        });
        it('should pause, rejecting all calls', async function() {
            const tokenContract = this.testingTokens[0];
            let myTokenAddresses = [tokenAddresses.ETH, ...this.testingTokens.map(t => t.address)];
            this.balancesBeforeUpgrade = await Promise.all(myTokenAddresses.map(_tokenAddr => this.vault.totalDepositedToSCAmount(_tokenAddr)));
            const getAdmin = (pVault) => {
                return getCodeSlot(pVault.address, '0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103');
            }
            console.info('admin', await getAdmin(this.vault), this.vaultAdmin.address);
            await confirm(this.proxy.connect(this.vaultAdmin).pause());
            // after pause, all calls should revert
            const appr = await confirm(tokenContract.connect(this.tokenGuy.signer).approve(this.vault.address, amount));
            await expect(appr).to.emit(tokenContract, 'Approval')
            .withArgs(this.tokenGuy.signer.address, this.vault.address, amount);
            try {
                await expect(this.vault.connect(this.unshieldSender).estimateGas.depositERC20(tokenContract.address, amount, this.tokenGuy.inc, "0x0000000000000000000000000000000000000000000000000000000000000000", "0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")).to.be.reverted;
            } catch(e) {
                if (hre.network.name == 'localhost') throw e; else {
                    console.error("Skipping known issue: testnet provider does not register Transaction Revert");
                    console.error(e);
                }
            }
        });

        it('should upgrade implementation & unpause', async function() {
            let currentImpl = await getImplementation(this.proxy.address);
            let u;
            if (currentImpl.eq(BN.from(this.previousVaultAddress))) {
                u = this.latestVaultAddress;
                console.info(`upgrade to ${u}`);
            } else {
                u = this.previousVaultAddress;
                console.info(`downgrade to ${u}`);
            }
            let myTokenAddresses = [tokenAddresses.ETH, ...this.testingTokens.map(t => t.address)];
            await confirm(this.proxy.connect(this.vaultAdmin).upgradeTo(u));
            await confirm(this.proxy.connect(this.vaultAdmin).unpause());
            currentImpl = await getImplementation(this.proxy.address);
            await expect(currentImpl).to.equal(u);
            const balances = await Promise.all(myTokenAddresses.map(_tokenAddr => this.vault.totalDepositedToSCAmount(_tokenAddr)));
            console.info('balances before upgrade', this.balancesBeforeUpgrade.map(b => b.toString()));
            console.info('after upgrade', balances.map(b => b.toString()));
            await expect(balances).to.deep.equal(this.balancesBeforeUpgrade);
        });

        it.skip('should shield normally again', shieldToken(amount, 'Token1'));
    });

});