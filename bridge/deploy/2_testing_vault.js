const { deployments, ethers } = require('hardhat');

module.exports = async({
    getNamedAccounts,
    deployments,
    getChainId,
    getUnnamedAccounts,
}) => {
    const { deploy, log, execute, rawTx } = deployments;
    const { deployer, tokenUser } = await getNamedAccounts();
    const addresses = hre.networkCfg().deployed || {};

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
        await deploy(tokenName, {from: deployer, args: []});
        await execute(tokenName, {from: deployer}, 'mint', tokenUser, ethers.utils.parseUnits('10', 'ether'));
        await execute(tokenName, {from: deployer}, 'mint', testingExchange, ethers.utils.parseUnits('1', 'ether'));
    })
    log('Deployed some testing contracts for localhost network');
}

module.exports.tags = ['1', 'testing', 'local'];
// always skip for public networks
module.exports.skip = env => Promise.resolve(env.network.name != 'localhost');