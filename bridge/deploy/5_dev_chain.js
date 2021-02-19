const { deployments, ethers } = require('hardhat');
const cfg = require('../hardhat.config');

const path = require('path');
const util = require('util');
let {exec} = require('child_process');
exec = util.promisify(exec);


module.exports = async({
    getNamedAccounts,
    deployments,
    getChainId,
    getUnnamedAccounts,
}) => {
    const { log } = deployments;
    const { address: vaultAddress } = await deployments.get('TransparentUpgradeableProxy');
    let myEnv = process.env;

    const chainId = await getChainId();
    const network = cfg.getNetwork(cfg.networks, chainId);
    myEnv.CONTRACT = vaultAddress;
    myEnv.GETH_NAME = network.geth_name_conf;
    myEnv.GETH_PORT = network.geth_port_conf || '';
    myEnv.GETH_PROTOCOL = network.geth_protocol_conf || '';
    // log(__dirname);
    const localNetPath = path.join(__dirname, '..', 'scripts', 'inc-local-net');
    // console.debug(localNetPath);
    outputs = await exec('./start-network.sh', {env: myEnv, cwd: localNetPath});
    log(outputs.stdout);
    const preload = require(path.join(localNetPath, 'preload'));
    const {inc, Inc, senders, addresses} = await preload();
    // wait for the new chain to stabilize, convert to v2 and fund acc1 to acc5
    log("Allocating funds...");
    await senders[0].waitHeight(10);
    let tx = await senders[0].convert([], 15);
    await senders[0].waitTx(tx.Response.txId, 2);
    const payments = [1, 2, 3, 4, 5].map(i => ({PaymentAddress: addresses[i], Amount: 8000000}));
    tx = await senders[0].prv(payments, 10);
    await senders[0].waitTx(tx.Response.txId, 2);
};

module.exports.tags = ['5', 'dev-chain'];
module.exports.dependencies = ['vault'];