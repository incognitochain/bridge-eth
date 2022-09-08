const { deployments, ethers } = require('hardhat');

module.exports = async({
    getNamedAccounts,
    deployments,
    getChainId,
    getUnnamedAccounts,
}) => {
    const { deploy, log, execute, rawTx } = deployments;
    const { deployer } = await getNamedAccounts();
    const addresses = hre.networkCfg().deployed || {};
    
}

module.exports.tags = ['5'];
// need more editing to run on a public network
// module.exports.skip = env => Promise.resolve(hre.network.name == 'localhost' && !process.env.FORK);