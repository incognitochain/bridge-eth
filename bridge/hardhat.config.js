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


/* test using the command:
* PRIVKEY=0x... hh mint-watch --chi 280 --gas 100 --network localhost
* (signer #6 in 'localhost' Hardhat provider)
* ('localhost' provider uses gasPrice=8gwei, so try --gas 7 for example)
*/
task("mint-watch", "Keep minting CHI token to sustain a minimum")
    .addOptionalParam("chi", "The minimum CHI to maintain", 280, types.int)
    .addOptionalParam("gas", "The gas price (gwei) above which we do NOT mint", 100, types.int)
    .setAction(async taskArgs => {
        // const accounts = await ethers.getSigners();
        // const chiFunder = accounts[6];
        // mne = "test test test test test test test test test test test junk";
        // const chiFunder = ethers.Wallet.fromMnemonic(mne, "m/44'/60'/0'/0/6");
        const privateKey = process.env.PRIVKEY;
        
        
        const chiFunder = new ethers.Wallet(privateKey, ethers.provider);
        console.log('Mint CHI to address', chiFunder.address);
        const { getChi, mintChi, chiAddress } = require('./scripts/chi');
        console.log('Parameters', taskArgs, chiAddress)
        const chi = await getChi();
        let sleep = (ms) => {
            return new Promise(resolve => setTimeout(resolve, ms));
        }
        // can only be stopped by an interrupt
        while (true){
            try {
                const gasPrice = await ethers.provider.getGasPrice();
                const mintBelowPrice = ethers.utils.parseUnits(String(taskArgs.gas), 'gwei');
                const balance = await chi.balanceOf(chiFunder.address);
                // console.log('gas price at', gasPrice.toNumber(), 'vs', mintBelowPrice.toNumber());

                if (gasPrice.gt(mintBelowPrice)) throw 'Gas is too expensive';
                if (balance.gte(taskArgs.chi)) throw 'Balance is not low enough';
                await mintChi(chiFunder, taskArgs.chi - balance, chi);
            } catch(e){ console.error(e) } // ignore errors
            // try every 3 mins (dev: 3s)
            await sleep(3000);
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