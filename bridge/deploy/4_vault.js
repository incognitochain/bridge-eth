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
    console.log('Deploy:', vaultContractName);
    let vaultResult = await deploy(vaultContractName, {
        from: deployer,
        args: [],
        // skipIfAlreadyDeployed: true,
        log: true,
        // waitConfirmations: 6,
    });
    const vaultFactory = await ethers.getContractFactory(vaultContractName);
    const vaultImpl = await vaultFactory.attach(vaultResult.address);

    let previousVault;
    try {
        previousVault = await getInstance('PrevVault');
    } catch (e) {
        previousVault = {address: '0x0000000000000000000000000000000000000000'};
    }

    // '0xdbeCBd9F55922e6487b24B4Fed572D5BF4982562'
    const regAddr = '0xdbeCBd9F55922e6487b24B4Fed572D5BF4982562'; // regulator address TBD
    const executorContractName = hre.networkCfg().executorContractName || 'UniswapV2Trade';
    const ex = await deployments.get(executorContractName);
    const initializeData = vaultImpl.interface.encodeFunctionData('initialize', [previousVault.address, regAddr, ex.address]);
    
    let proxyResult = await deploy('TransparentUpgradeableProxy', {
        from: deployer,
        args: [vaultImpl.address, vaultAdmin, ip.address, initializeData],
        skipIfAlreadyDeployed: true,
        log: true,
    });

    const proxy = await getInstance('TransparentUpgradeableProxy');
    const vault = await vaultFactory.attach(proxyResult.address);
    let borrowedPermissionFrom = null;
    const prevVaultAdmin = hre.networkCfg().deployed.adminToBorrowFrom;
    try {
        // when in fork environment, simulate the real prev admin yielding permission
        if (process.env.FORK && prevVaultAdmin) {
            log(`FORK-ENV: "Borrowing" Admin permission`);
            const prevAdm = await ethers.getImpersonatedSigner(prevVaultAdmin);
            // const moneyAcc = await ethers.getSigner(deployer);
            // await confirm(moneyAcc.sendTransaction({to: prevVaultAdmin, value: ethers.utils.parseUnits('1', 'ether')}));
            await confirm(proxy.connect(prevAdm).retire(vaultAdmin));
        }
        const adm = await getAdmin(proxy);
        if (!BN.from(vaultAdmin).eq(BN.from(adm))) {
            console.log(BN.from(vaultAdmin).toHexString());
            console.log(BN.from(adm).toHexString())
            console.log('Claim()-ing admin permission to upgrade');
            await confirm(proxy.connect(vaultAdminSigner).claim());
        } else {
            log('Proceeding with Admin role');
        }
    } catch (e) {
        throw e;
    }

    if (proxyResult.newlyDeployed) {
        log('deployed new proxy & initialized with params', vault.address, vaultAdmin, ip.address, initializeData);
    } else {
        // await proxy.connect(vaultAdminSigner).upgradeTo(vault.address);
        // await proxy.connect(vaultAdminSigner).upgradeIncognito(ip.address);
        const upgradeData = vaultFactory.interface.encodeFunctionData('upgradeVaultStorage', [regAddr, ex.address]);

        await proxy.connect(vaultAdminSigner).upgradeToAndCall(vaultImpl.address, upgradeData);
        log('upgraded existing proxy to new implementation with params', vaultImpl.address, vaultAdmin, regAddr, ex.address);
    }
    log(`DEV : Incognito nodes should use ${proxy.address} as EthVaultContract`);
    // await confirm(vault.setRegulator(regAddr));

    if (borrowedPermissionFrom) {
        const theVaultProxy = await getInstance('TransparentUpgradeableProxy');
        let tx = await confirm(theVaultProxy.connect(vaultAdminSigner).retire(borrowedPermissionFrom));
        let rc = await tx.wait();
        log(`Gas used: ${rc.gasUsed.toString()}`);
    }
};

module.exports.tags = ['4', 'vault', 'proxy'];
module.exports.dependencies = ['incognito-proxy', 'testing'];