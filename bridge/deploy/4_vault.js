const { deployments, ethers } = require('hardhat');
const { getInstance, confirm } = require('../scripts/utils');

module.exports = async({
    getNamedAccounts,
    deployments,
    getChainId,
    getUnnamedAccounts,
}) => {
    const { deploy, log } = deployments;
    let { deployer, vaultAdmin } = await getNamedAccounts();
    let vaultAdminSigner = await ethers.getSigner(vaultAdmin);
    const ip = await deployments.get('IncognitoProxy');

    const vaultContractName = hre.networkCfg().vaultContractName || 'Vault';
    log('Deploy:', vaultContractName);
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

    const regAddr = '0xdbeCBd9F55922e6487b24B4Fed572D5BF4982562'; // regulator address TBD
    const executorContractName = hre.networkCfg().executorContractName || 'UniswapV2Trade';
    const ex = await deployments.get(executorContractName);
    const initializeData = vaultImpl.interface.encodeFunctionData('initialize', [previousVault.address, regAddr, ex.address]);
    
    // NOTE: on mainnet, we will be upgrading via proxy instead of deploying & initializing
    let proxyResult = await deploy('TransparentUpgradeableProxy', {
        from: deployer,
        args: [vaultImpl.address, vaultAdminSigner.address, ip.address, initializeData],
        skipIfAlreadyDeployed: true,
        log: true,
    });
    if (proxyResult.newlyDeployed) {
        log('deployed new proxy & initialized with params', vaultImpl.address, vaultAdminSigner.address, ip.address, initializeData);
    }

    const proxy = await getInstance('TransparentUpgradeableProxy');
    log(`DEV : Incognito nodes should use ${proxy.address} as EthVaultContract`);
    // await confirm(vault.setRegulator(regAddr));
};

module.exports.tags = ['4', 'vault', 'proxy'];
module.exports.dependencies = ['incognito-proxy', 'testing'];