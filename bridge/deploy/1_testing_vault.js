const { deployments, ethers } = require('hardhat');

module.exports = async({
    getNamedAccounts,
    deployments,
    getChainId,
    getUnnamedAccounts,
}) => {
    const { deploy, log, execute } = deployments;
    const { deployer, tokenUser } = await getNamedAccounts();
    const chainId = await getChainId();
    const cfg = require('../hardhat.config');
    if (chainId!=31337 && chainId!=4) {
        // only applicable for dev & test networks
        return;
    }

    await deploy('WrongChi', {from: deployer, args: []});
    await deploy('TestingVault', {from: deployer, args: []});
    // deploy 2 test ERC20 token contracts and mint 10000 of each to tokenUser
    await deploy('Token1', {from: deployer, args: []});
    await execute('Token1', {from: deployer}, 'mint', tokenUser, ethers.utils.parseUnits('10', 'ether'));
    await deploy('Token2', {from: deployer, args: []});
    await execute('Token2', {from: deployer}, 'mint', tokenUser, ethers.utils.parseUnits('10', 'ether'));
    log('Deployed some testing contracts for Hardhat network');
}

module.exports.tags = ['1', 'testing', 'local'];
