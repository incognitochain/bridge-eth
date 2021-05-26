const { deployments, ethers } = require('hardhat');

module.exports = async({
    getNamedAccounts,
    deployments,
    getChainId,
    getUnnamedAccounts,
}) => {
    const { deploy, log, execute, rawTx } = deployments;
    const { deployer, tokenUser, vaultAdmin , previousVaultAdmin: prevVaultAdmin } = await getNamedAccounts();
    const addresses = hre.networkCfg().deployed || {};

    const ip = await deployments.get('IncognitoProxy');
    // deploy previous implementation of Vault to local network
    let vaultResult = await deploy('PreviousVersionVault', {
        from: deployer,
        args: [],
        skipIfAlreadyDeployed: true,
        log: true
    });

    const vaultFactory = await ethers.getContractFactory('Vault');
    const vault = await vaultFactory.attach(vaultResult.address);
    const initializeData = vaultFactory.interface.encodeFunctionData('initialize', ['0x0000000000000000000000000000000000000000']);
    log('will deploy proxy & upgrade with params', vault.address, vaultAdmin, ip.address, initializeData);
    let proxyResult = await deploy('TransparentUpgradeableProxy', {
        from: deployer,
        args: [vault.address, vaultAdmin, ip.address, initializeData],
        skipIfAlreadyDeployed: true,
        log: true,
    });
    log('Deployed previous vault implementation for testing network');

    // LOCAL-ONLY contract
    if (hre.networkCfg().chainId != 31337) return;
    let res = await deploy('TestingExchange', {from: deployer, args: []});
    if (res.newlyDeployed) {
        log(`Testing Exchange deployed at ${res.address}`);
    }
    const kb = await deployments.get('TestingExchange');
    let testingExchange = kb.address;

    res = await deploy('VaultHelper', {from: deployer, args: []});
    // fund testingExchange with ether
    await rawTx({from:deployer, to: testingExchange, value: ethers.utils.parseUnits('1', 'ether')});
    let tokenNames = ['Token1', 'Token2', 'Token3'];
    // deploy test ERC20 token contracts and mint them to tokenUser & testingExchange
    tokenNames.forEach(async tokenName => {
        const res = await deploy(tokenName, {from: deployer, args: []});
        console.log(`Token ${res.address}`);
        await execute(tokenName, {from: deployer}, 'mint', tokenUser, ethers.utils.parseUnits('10', 'ether'));
        await execute(tokenName, {from: deployer}, 'mint', testingExchange, ethers.utils.parseUnits('1', 'ether'));
    })
    log('Deployed some testing contracts for localhost network');
}

module.exports.tags = ['1', 'testing', 'local'];
module.exports.dependencies = ['incognito-proxy'];
// always skip for public networks
module.exports.skip = env => Promise.resolve(!(env.network.name == 'localhost' || env.network.name == 'kovan') || process.env.FORK);