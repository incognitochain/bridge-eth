const { deployments, ethers } = require('hardhat');

module.exports = async({
    getNamedAccounts,
    deployments,
    getChainId,
    getUnnamedAccounts,
}) => {
    let addresses = hre.networkCfg().deployed || {};
    // transform dict into flat array of addresses
    addresses = Object.values(addresses);
    addresses = addresses.flat(3);
    console.log(`Impersonating addresses...`);
    console.log(addresses);
    await Promise.all(addresses.map(_addr => hre.network.provider.request({
        method: "hardhat_impersonateAccount",
        params: [_addr]
    })));
}

module.exports.tags = ['1', 'fork'];
// always skip for public networks
module.exports.skip = env => Promise.resolve(!process.env.FORK);