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
    const vaultAdminSigner = await ethers.getSigner(vaultAdmin);
    const getAdmin = (pVault) => {
        return getCodeSlot(pVault.address, '0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103');
    }
    const ip = await deployments.get('IncognitoProxy');

    const vaultContractName = hre.networkCfg().vaultContractName || 'Vault';
    console.log('Deploy: ', vaultContractName);
    let vaultResult = await deploy(vaultContractName, {
        from: deployer,
        args: [],
        skipIfAlreadyDeployed: true,
        log: true,
        // waitConfirmations: 6,
    });
    const vaultFactory = await ethers.getContract(vaultContractName);
    const vault = await vaultFactory.attach(vaultResult.address);

    let previousVault;
    try {
        previousVault = await ethers.getContract('PrevVault');
    } catch (e) {
        previousVault = {address: '0x0000000000000000000000000000000000000000'};
    }

    const initializeData = vaultFactory.interface.encodeFunctionData('initialize', [previousVault.address]);
    log('will deploy proxy & upgrade with params', vault.address, vaultAdmin, ip.address, initializeData);
    let proxyResult = await deploy('TransparentUpgradeableProxy', {
        from: deployer,
        args: [vault.address, vaultAdmin, ip.address, initializeData],
        skipIfAlreadyDeployed: true,
        log: true,
    });

    const proxy = await ethers.getContract('TransparentUpgradeableProxy');
    let borrowedPermissionFrom = null;
    try {
        // when in fork environment, simulate the real prev admin yielding permission
        if (process.env.FORK) {
            log(`FORK-ENV: "Borrowing" Admin permission and money`);
            const realAdm = await ethers.provider.getSigner(hre.networkCfg().deployed.adminToBorrowFrom);
            const moneyAcc = await ethers.getSigner(deployer);
            await confirm(moneyAcc.sendTransaction({to: prevVaultAdmin, value: ethers.utils.parseUnits('1', 'ether')}));
            await confirm(vault.connect(realAdm).retire(prevVaultAdmin));
        }
        const adm = await getAdmin(vault);
        if (BN.from(vaultAdmin).eq(BN.from(adm))) {
            console.log(vaultAdmin);
            console.log(adm)
            throw 'Need admin permission to upgrade';
        } else {
            log('Proceeding with Admin role');
        }
    } catch (e) {
        throw e;
    }

    if (!proxyResult.newlyDeployed) {
        // const upgradeData = vaultFactory.interface.encodeFunctionData('upgradeVaultStorageLayout', [hre.networkCfg().recoveryAddress]);
        log('will upgrade proxy to new implementation with params', vault.address, vaultAdmin);
        await proxy.connect(vaultAdminSigner).upgradeTo(vault.address);
        // await proxy.connect(vaultAdminSigner).upgradeToAndCall(vault.address, upgradeData);
    }
    log(`DEV : Incognito nodes should use ${proxy.address} as EthVaultContract`);

    if (borrowedPermissionFrom) {
        const theVaultProxy = await getInstance('TransparentUpgradeableProxy');
        let tx = await confirm(theVaultProxy.connect(vaultAdminSigner).retire(borrowedPermissionFrom));
        let rc = await tx.wait();
        log(`Gas used: ${rc.gasUsed.toString()}`);
    }
};

module.exports.tags = ['4', 'vault', 'proxy'];
module.exports.dependencies = ['fork', 'incognito-proxy'];