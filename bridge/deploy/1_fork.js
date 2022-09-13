// const { deployments, ethers } = require('hardhat');
const { getInstance, confirm, getCodeSlot } = require('../scripts/utils');

module.exports = async({
    getNamedAccounts,
    deployments,
    getChainId,
    getUnnamedAccounts,
}) => {
    const { deploy, log, execute, rawTx } = deployments;
    const signers = (await ethers.getSigners()).slice(0, 5);
    let addresses = hre.networkCfg().deployed || {};
    const { tokenFunders, testingTokens } = addresses;
    // transform dict into flat array of addresses
    addresses = Object.values(addresses);
    addresses = addresses.flat(3);
    addresses = addresses.map(e => typeof(e) == 'object' ? Object.values(e) : e);
    addresses = addresses.flat(3);
    console.log(`Impersonating addresses...`);
    console.log(addresses);
    await Promise.all(addresses.map(async _addr => {
        await hre.network.provider.request({
            method: "hardhat_impersonateAccount",
            params: [_addr]
        });
    }));

    if (Boolean(testingTokens)) {
        for (let i=0; i < testingTokens.length; i++) {
            const token = await getInstance('contracts/IERC20.sol:IERC20', null, testingTokens[i]);
            for (const s of signers) {
                const f = await ethers.getImpersonatedSigner(tokenFunders[i]);
                await signers[0].sendTransaction({value: ethers.utils.parseUnits('0.1', 'ether'), to: tokenFunders[i]});
                // console.log(f.address, testingTokens[i])
                await token.connect(f).transfer(s.address, ethers.utils.parseUnits('1', 'ether'));
            }
        }
    }
}

module.exports.tags = ['1', 'fork'];
// always skip for public networks
module.exports.skip = env => Promise.resolve(!process.env.FORK)// && hre.network.name != 'hardhat' && hre.network.name != 'localhost');