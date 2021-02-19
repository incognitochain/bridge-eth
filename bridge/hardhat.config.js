require("@nomiclabs/hardhat-waffle");
require("@nomiclabs/hardhat-web3");
require('hardhat-deploy');
require("hardhat-deploy-ethers");

// This is a sample Hardhat task. To learn how to create your own go to
// https://hardhat.org/guides/create-task.html
task("accounts", "Prints the list of accounts", async() => {
    const accounts = await ethers.getSigners();

    for (const account of accounts) {
        console.log(account.address);
    }
});

let getNetwork = (obj, id) => {
    let names = Object.keys(obj);
    let theKey = names.filter(name => obj[name].chainId==id)[0];
    return Object.assign({name: theKey}, obj[theKey]);
}

// You need to export an object to set up your config
// Go to https://hardhat.org/config/ to learn more

/**
 * @type import('hardhat/config').HardhatUserConfig
 */
module.exports = {
    solidity: "0.6.6",
    chiAddress: '0x0000000000004946c0e9F43F4Dee607b0eF1fA1c',
    paths: {
        tests: './tests'
    },
    namedAccounts: {
        deployer: {
            default: 0
        },
        vaultAdmin: {
            default: 1
        },
        tokenUser: {
            default: 3
        }
    },
    mocha: {
        bail: true,
        timeout: 1200000,
    },
    getNetwork,
    networks: {
        localhost: {
            chainId: 31337,
            geth_name_conf: '127.0.0.1',
            geth_port_conf: '8545',
            geth_protocol_conf: 'HTTP'
        },
        rinkeby: {
            chainId: 4,
            url: 'https://rinkeby.infura.io/v3/557131bed4eb4d09857f9c74445fc930',
            accounts: {
                mnemonic: 'swing ethics siren solution laundry movie arrange race sun hungry scout arrow'
            },
            geth_name_conf: 'rinkeby.infura.io/v3/557131bed4eb4d09857f9c74445fc930',
            geth_port_conf: '',
            geth_protocol_conf: 'HTTPS'
        },
        mainnet: {
            chainId: 1,
            url: 'https://mainnet.infura.io/v3/557131bed4eb4d09857f9c74445fc930',
            accounts: {
                mnemonic: 'swing ethics siren solution laundry movie arrange race sun hungry scout arrow'
            },
            geth_name_conf: 'mainnet.infura.io/v3/557131bed4eb4d09857f9c74445fc930',
            geth_port_conf: '',
            geth_protocol_conf: 'HTTPS'
        }
    }
};