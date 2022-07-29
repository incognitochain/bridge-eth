const { deployments, ethers } = require('hardhat');

module.exports = async({
    getNamedAccounts,
    deployments,
    getChainId,
    getUnnamedAccounts,
}) => {
    const { deploy, log, execute, rawTx } = deployments;
    const { deployer } = await getNamedAccounts();
    const addresses = hre.networkCfg().deployed;
    let kyberAddr, uniswapAddr;
    if (addresses.kyber && addresses.uniswap) {
        log("Using trade router addresses");
        log(addresses);
        // use real trade routers
        kyberAddr = addresses.kyber;
        uniswapAddr = addresses.uniswap;
    } else {
        // use testing exchange for local
        const c = await deployments.get('TestingExchange');
        kyberAddr = c.address;
        uniswapAddr = c.address;
    }

    res = await deploy('KBNTrade', {
        from: deployer,
        args: [kyberAddr],
        skipIfAlreadyDeployed: true,
        log: true
    });
    // const k = await deployments.read('KBNTrade', 'kyberNetworkProxyContract');
    // const kbr = await deployments.get('KyberNetworkProxy');
    // const mgp = await deployments.read('KyberNetworkProxy', 'maxGasPrice');
    // log(`max gas price ${mgp}`);
    // log(`mainnet Kyber router ${k}`);
    res = await deploy('UniswapV2Trade', {
        from: deployer,
        args: [uniswapAddr],
        skipIfAlreadyDeployed: true,
        log: true
    });
}

module.exports.tags = ['5', 'trade-testing'];
// need more editing to run on a public network
// module.exports.skip = env => Promise.resolve(hre.network.name == 'localhost' && !process.env.FORK);