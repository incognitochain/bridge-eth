const { deployments, ethers } = require('hardhat');

module.exports = async({
    getNamedAccounts,
    deployments,
    getChainId,
    getUnnamedAccounts,
}) => {
    const { deploy, log } = deployments;
    const { deployer, vaultAdmin } = await getNamedAccounts();

    const ip = await deployments.get('IncognitoProxy');
    let vaultResult = await deploy('Vault', {
        from: deployer,
        args: [],
    });
    if (vaultResult.newlyDeployed) {
        log(`Vault deployed at ${vaultResult.address}`);
    }
    const chainId = await getChainId();
    const cfg = require('../hardhat.config');
    let chiAddress = cfg.chiAddress;
    const chi = await deployments.getOrNull('ChiToken');
    if (chi) {
        // when in dev environment, use the local CHI instance
        chiAddress = chi.address;
    }
    console.log('chi is',chiAddress);
    const vaultFac = await ethers.getContractFactory('Vault');
    const vault = await vaultFac.attach(vaultResult.address);
    const initializeData = vault.interface.encodeFunctionData('initialize', ['0x0000000000000000000000000000000000000000', chiAddress]);
    // log(initializeData);
    let proxyResult = await deploy('TransparentUpgradeableProxy', {
        from: deployer,
        gasLimit: 4000000,
        args: [vaultResult.address, vaultAdmin, ip.address, initializeData],
    });
    if (proxyResult.newlyDeployed) {
        log(`Proxy deployed at ${proxyResult.address}`);
    }
};

module.exports.tags = ['3', 'vault', 'proxy'];
module.exports.dependencies = ['incognito-proxy'];