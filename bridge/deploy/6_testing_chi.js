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
    await deploy('ChiToken', {from: deployer, args: []});
}

module.exports.tags = ['6', 'test-chi'];