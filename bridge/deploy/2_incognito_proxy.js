const { deployments } = require('hardhat');

module.exports = async({
    getNamedAccounts,
    deployments,
    getChainId,
    getUnnamedAccounts,
}) => {
    const { deploy } = deployments;
    const { deployer } = await getNamedAccounts();
    const beacons = [
        "0x3cD69B1A595B7A9589391538d29ee7663326e4d3", // only this committee is relevant
    ];
    const bridges = [
        "0xD2902ab2F5dF2b17C5A5aa380f511F04a2542E10" // deprecated
    ];

    await deploy('IncognitoProxy', {
        from: deployer,
        gasLimit: 4000000,
        args: [deployer, beacons, bridges],
    });
};

module.exports.tags = ['2', 'incognito-proxy'];
module.exports.dependencies = ['testing'];