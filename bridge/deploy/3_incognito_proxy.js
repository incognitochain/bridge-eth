const { deployments } = require('hardhat');

module.exports = async({
    getNamedAccounts,
    deployments,
    getChainId,
    getUnnamedAccounts,
}) => {
    const { deploy, log } = deployments;
    const { deployer } = await getNamedAccounts();
    const cfg = require('../hardhat.config');
    const { beacons, bridges } = hre.networkCfg().committees;

    const result = await deploy('IncognitoProxy', {
        from: deployer,
        args: [deployer, beacons, bridges],
        skipIfAlreadyDeployed: true,
        log: true
    });
};

module.exports.tags = ['2', 'incognito-proxy'];
module.exports.dependencies = [];