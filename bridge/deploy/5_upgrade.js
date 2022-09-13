const { deployments, ethers } = require('hardhat');
const { getInstance, confirm, getCodeSlot } = require('../scripts/utils');
const BN = ethers.BigNumber;

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
    const regAddr = '0xdbeCBd9F55922e6487b24B4Fed572D5BF4982562'; // regulator address TBD
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

        const tx = await proxy.connect(vaultAdminSigner).upgradeToAndCall(vaultImpl.address, upgradeData);
        log('upgraded existing proxy to new implementation with params', vaultImpl.address, vaultAdminSigner.address, regAddr, ex.address);
        const { gasUsed } = await tx.wait();
        log('gas used for upgrade:', gasUsed.toString());
    }
    // await confirm(vault.setRegulator(regAddr));
};

module.exports.tags = ['5', 'upgrade'];
// need more editing to run on a public network
// module.exports.skip = env => Promise.resolve(hre.network.name == 'localhost' && !process.env.FORK);