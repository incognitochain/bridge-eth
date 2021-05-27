const { deployments, ethers } = require('hardhat');
const path = require('path');
const util = require('util');
let { exec } = require('child_process');
exec = util.promisify(exec);


module.exports = async ({
    getNamedAccounts,
    deployments,
    getChainId,
    getUnnamedAccounts,
}) => {
    const { log } = deployments;
    const { address: vaultAddress } = await deployments.get('TransparentUpgradeableProxy');
    const network = hre.networkCfg();
    const publishedPorts = { 1: 9334, 2: 9338, 3: 9350, 4: 9330 };
    // clear all previous data
    await exec(`docker kill n1 n2 n3 n4; docker rm n1 n2 n3 n4; docker network rm incognito-local`).catch(e => console.error('Could not clear previous chain data'));
    // set up a network for containers
    await exec(`docker network create --subnet 192.168.1.0/24 --gateway 192.168.1.1 incognito-local`).catch(console.error);
    // via Docker, fire up a highway node + 3 Incognito nodes (2 shards, 1 beacon). Map their listening ports to the host. Set their Ethereum endpoints to Hardhat's network
    await Promise.all(
        [1, 2, 3, 4].map(i => exec(`docker run -d -it -p ${publishedPorts[i]}:${publishedPorts[i]}/tcp -p ${publishedPorts[i]}:${publishedPorts[i]}/udp --env TYPE=${i} --env CONTRACT=${vaultAddress} --env GETH_NAME=${network.geth_name_conf} --env GETH_PORT=${network.geth_port_conf || ''} --env GETH_PROTOCOL=${network.geth_protocol_conf || ''} --name n${i} --network incognito-local --ip 192.168.1.${100+i} dariuslt/incognito-local-dev-network:prebuilt; sleep 2`)
            .then(o => console.error(o.stderr))
            .catch(console.error))
    );
    let loadDevEnv = (prov) => {
        const { Inc } = require('../scripts/external');
        const privateKeys = [
            '112t8roafGgHL1rhAP9632Yef3sx5k8xgp8cwK4MCJsCL1UWcxXvpzg97N4dwvcD735iKf31Q2ZgrAvKfVjeSUEvnzKJyyJD3GqqSZdxN4or',
            '112t8rnZDRztVgPjbYQiXS7mJgaTzn66NvHD7Vus2SrhSAY611AzADsPFzKjKQCKWTgbkgYrCPo9atvSMoCf9KT23Sc7Js9RKhzbNJkxpJU6',
            '112t8rne7fpTVvSgZcSgyFV23FYEv3sbRRJZzPscRcTo8DsdZwstgn6UyHbnKHmyLJrSkvF13fzkZ4e8YD5A2wg8jzUZx6Yscdr4NuUUQDAt',
            '112t8rnXoBXrThDTACHx2rbEq7nBgrzcZhVZV4fvNEcGJetQ13spZRMuW5ncvsKA1KvtkauZuK2jV8pxEZLpiuHtKX3FkKv2uC5ZeRC8L6we',
            '112t8rnbcZ92v5omVfbXf1gu7j7S1xxr2eppxitbHfjAMHWdLLBjBcQSv1X1cKjarJLffrPGwBhqZzBvEeA9PhtKeM8ALWiWjhUzN5Fi6WVC',
            '112t8rnZUQXxcbayAZvyyZyKDhwVJBLkHuTKMhrS51nQZcXKYXGopUTj22JtZ8KxYQcak54KUQLhimv1GLLPFk1cc8JCHZ2JwxCRXGsg4gXU',
            '112t8rnXDS4cAjFVgCDEw4sWGdaqQSbKLRH1Hu4nUPBFPJdn29YgUei2KXNEtC8mhi1sEZb1V3gnXdAXjmCuxPa49rbHcH9uNaf85cnF3tMw',
            '112t8rnYoioTRNsM8gnUYt54ThWWrRnG4e1nRX147MWGbEazYP7RWrEUB58JLnBjKhh49FMS5o5ttypZucfw5dFYMAsgDUsHPa9BAasY8U1i',
            '112t8rnXtw6pWwowv88Ry4XxukFNLfbbY2PLh2ph38ixbCbZKwf9ZxVjd4s7jU3RSdKctC7gGZp9piy8nZoLqHwqDBWcsMHWsQg27S5WCdm4',
        ]
        let inc = new Inc.SimpleWallet();
        inc.setProvider(prov);
        return Inc.init()
            .then(_ => {
                // make a transactor for each private key
                return Promise.all(privateKeys.map(_k => inc.NewTransactor(_k)))
                    .then(_senders =>
                        Promise.all(_senders.map(s => s.submitKeyAndSync()))
                        .then(_ => _senders))
                    // aggregate relevant variables and unfold them onto global
                    .then(_senders => {
                        const res = {
                            privateKeys,
                            addresses: _senders.map(_t => _t.key.base58CheckSerialize(Inc.constants.PaymentAddressType)),
                            Inc,
                            inc,
                            senders: _senders,
                            rpc: inc.rpc,
                        }
                        return res;
                    })
            })
    }
    const { inc, Inc, senders, addresses } = await loadDevEnv(network.providers[0]);
    // wait for the new chain to stabilize, convert to v2 and fund acc1 to acc5
    log("Allocating funds...");
    await senders[0].waitHeight(10);
    let tx = await senders[0].convert({ transfer: { fee: 15 } });
    await senders[0].waitTx(tx.Response.txId, 2);
    const payments = [1, 2, 3, 4, 5].map(i => ({ PaymentAddress: addresses[i], Amount: 8000000 }));
    tx = await senders[0].prv({ transfer: { prvPayments: payments, fee: 10 } });
    await senders[0].waitTx(tx.Response.txId, 2);
};

module.exports.tags = ['6', 'dev-chain'];
module.exports.dependencies = ['vault'];
// use for dev/test networks only
module.exports.skip = env => Promise.resolve(env.network.name != 'localhost' && env.network.name != 'kovan');