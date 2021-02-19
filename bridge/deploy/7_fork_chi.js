const { deployments, ethers } = require('hardhat');

module.exports = async({
    getNamedAccounts,
    deployments,
    getChainId,
    getUnnamedAccounts,
}) => {
    const { deploy, log, execute } = deployments;
    const { deployer, tokenUser } = await getNamedAccounts();
    const cfg = require('../hardhat.config');
    await hre.network.provider.request({
        method: "hardhat_impersonateAccount",
        params: [cfg.chiAddress]
    });
}

module.exports.tags = ['7', 'fork-chi'];