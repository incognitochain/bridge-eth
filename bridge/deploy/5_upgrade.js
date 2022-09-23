const { deployments, ethers } = require('hardhat');
const { getInstance, confirm, getCodeSlot } = require('../scripts/utils');
const BN = ethers.BigNumber;
const latestStorageLayoutVersion = 2;

module.exports = async({
    getNamedAccounts,
    deployments,
    getChainId,
    getUnnamedAccounts,
}) => {
    const { deploy, log } = deployments;
    let { deployer, vaultAdmin } = await getNamedAccounts();
    let vaultAdminSigner = await ethers.getSigner(vaultAdmin);
    try {
        // when in fork environment, test-upgrade vault using impersonated signer
        const impAdm = hre.networkCfg().deployed.adminToBorrowFrom;
        if (process.env.FORK && impAdm) {
            log(`FORK-ENV: sending txs as mainnet vaultAdmin`);
            vaultAdminSigner = await ethers.getImpersonatedSigner(impAdm);
            const d = await ethers.getSigner(deployer);
            await d.sendTransaction({value: ethers.utils.parseUnits('5', 'ether'), to: impAdm});
        }
    } catch (e) {
        throw e;
    }
    const getAdmin = (pVault) => {
        return getCodeSlot(pVault.address, '0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103');
    }
    const regAddr = '0x86a879cF735B0F462F4659bFAADD91F84Ed61aB3'; // mainnet regulator address
    const executorContractName = hre.networkCfg().executorContractName || 'UniswapV2Trade';
    const ex = await deployments.get(executorContractName);
    const proxy = await getInstance('TransparentUpgradeableProxy');
    const adminAddrCompare = await getAdmin(proxy);
    if (BN.from(vaultAdminSigner.address).eq(BN.from(adminAddrCompare))) {
        log('Proceeding with Admin role');
        // await proxy.connect(vaultAdminSigner).upgradeTo(vault.address);
        // await proxy.connect(vaultAdminSigner).upgradeIncognito(ip.address);
        const vaultContractName = hre.networkCfg().vaultContractName || 'Vault';
        const vaultFactory = await ethers.getContractFactory(vaultContractName);
        const vaultImpl = await deployments.get(vaultContractName);
        const upgradeData = vaultFactory.interface.encodeFunctionData('upgradeVaultStorage', [regAddr, ex.address]);
        const v = await vaultFactory.attach(proxy.address);
        let slver = BN.from(0); 
        try {
            slver = await v.storageLayoutVersion();
        } catch (e){ console.log('...vault has no storage layout version, use zero') } // likely upgrading from an old version
        console.log('storage layout version', slver.toString());
        if (slver.eq(latestStorageLayoutVersion)) {
            console.log('Upgrade without call');
            const tx = await proxy.connect(vaultAdminSigner).upgradeTo(vaultImpl.address);
        } else {
            const tx = await proxy.connect(vaultAdminSigner).upgradeToAndCall(vaultImpl.address, upgradeData);
            log('upgraded existing proxy to new implementation with params', vaultImpl.address, vaultAdminSigner.address, regAddr, ex.address);
            const { gasUsed } = await tx.wait();
            log('gas used for upgrade:', gasUsed.toString());
        }
    } else {
        throw `Unable to upgrade. Your signer ${vaultAdminSigner.address} does not match Vault Admin`;
    }
    // await confirm(vault.setRegulator(regAddr));
};

module.exports.tags = ['5', 'upgrade'];
// need more editing to run on a public network
// module.exports.skip = env => Promise.resolve(hre.network.name == 'localhost' && !process.env.FORK);