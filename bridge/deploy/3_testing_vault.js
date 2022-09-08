const { deployments, ethers } = require('hardhat');

module.exports = async({
    getNamedAccounts,
    deployments,
    getChainId,
    getUnnamedAccounts,
}) => {
    const { deploy, log, execute, rawTx } = deployments;
    const { deployer, tokenUser, vaultAdmin } = await getNamedAccounts();
    const addresses = hre.networkCfg().deployed || {};
    let testingExchange = null;

    if (hre.networkCfg().chainId == 31337) { // for local
        const ip = await deployments.get('IncognitoProxy');
        // deploy previous implementation of Vault to local network
        let vaultResult = await deploy('PreviousVersionVault', {
            from: deployer,
            args: [],
            skipIfAlreadyDeployed: true,
            log: true
        });

        const vaultFactory = await ethers.getContractFactory('PreviousVersionVault');
        const vault = await vaultFactory.attach(vaultResult.address);
        const initializeData = vaultFactory.interface.encodeFunctionData('initialize', ['0x0000000000000000000000000000000000000000']);
        
        let proxyResult = await deploy('TransparentUpgradeableProxy', {
            from: deployer,
            args: [vault.address, vaultAdmin, ip.address, initializeData],
            skipIfAlreadyDeployed: true,
            log: true,
        });
        if (proxyResult.newlyDeployed) {
            log('deployed new proxy & initialized with params', vault.address, vaultAdmin, ip.address, initializeData);
        }
        log('Deployed previous vault implementation for testing network');

    
        let res = await deploy('TestingExchange', {from: deployer, args: [], skipIfAlreadyDeployed: true,
            log: true,
        });
        if (res.newlyDeployed) {
            log(`Testing Exchange deployed at ${res.address}`);
        }
        const kb = await deployments.get('TestingExchange');
        testingExchange = kb.address;

        res = await deploy('VaultHelper', {from: deployer, args: []});
        // await rawTx({from:deployer, to: testingExchange, value: ethers.utils.parseUnits('1', 'ether')}); // fund testingExchange with ether
        let tokenNames = ['Token1', 'Token2', 'Token3'];
        // deploy test ERC20 token contracts and mint them to tokenUser & testingExchange
        for (const tokenName of tokenNames) {
            const res = await deploy(tokenName, {from: deployer, args: [], skipIfAlreadyDeployed: true});
            console.log(`Token ${res.address}`);
            await execute(tokenName, {from: deployer}, 'mint', tokenUser, ethers.utils.parseUnits('10', 'ether'));
            await execute(tokenName, {from: deployer}, 'mint', testingExchange, ethers.utils.parseUnits('1', 'ether'));
        }
    }

    const routers = addresses.routers || {};
    log("Using trade router addresses", routers);
    for (const pname in routers) {
        const extRouterAddr = routers[pname] == 1 ? testingExchange : routers[pname];
        await deploy(pname, {
            from: deployer,
            args: [extRouterAddr],
            skipIfAlreadyDeployed: true,
            log: true
        });
    }
    log('Deployed some testing contracts for localhost network');
}

module.exports.tags = ['3', 'testing'];
module.exports.dependencies = ['incognito-proxy'];
// always skip for public networks
module.exports.skip = env => Promise.resolve(!(env.network.name == 'localhost' || env.network.name == 'hardhat' || env.network.name == 'kovan' || env.network.name == 'goerli') || process.env.FORK);