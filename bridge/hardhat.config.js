require("@nomiclabs/hardhat-waffle");
// require("@nomiclabs/hardhat-web3");
require("hardhat-deploy");
require("hardhat-deploy-ethers");
const fs = require('fs');

// This is a sample Hardhat task. To learn how to create your own go to
// https://hardhat.org/guides/create-task.html
task("accounts", "Prints the list of accounts", async() => {
    const accounts = await ethers.getSigners();

    for (const account of accounts) {
        console.log(account.address);
    }
});

task("list-contracts", "Exports & prints the list of deployed contracts")
    .setAction(async taskArgs => {
        const filename = `./${hre.network.name}-exports.json`;
        await hre.run('export', {export: filename});
        const obj = require(filename).contracts;
        let result = {};
        Object.keys(obj).map(k => {
            result[k] = obj[k].address;
        });
        console.log(result);
        const { getImplementation, getInstance } = require('./scripts/utils');
        let addr = await getImplementation(result.TransparentUpgradeableProxy);
        console.log(`Vault Proxy's current implementation is ${ethers.utils.hexlify(addr)}`);
        const kyb = await getInstance('KBNTrade');
        addr = await kyb.kyberNetworkProxyContract();
        console.log(`KBNTrade is using KyberNetworkRouter ${ethers.utils.hexlify(addr)}`);
        const uni = await getInstance('UniswapV2Trade');
        addr = await uni.uniswapV2();
        console.log(`UniswapV2Trade is using UniswapRouter ${ethers.utils.hexlify(addr)}`);
    });

const devCommittees = {
    beacons : [
        "0x3cD69B1A595B7A9589391538d29ee7663326e4d3"
    ],
    bridges : [
        "0xD2902ab2F5dF2b17C5A5aa380f511F04a2542E10"
    ],
};
const devCommitteesBig = {
    beacons: ["0xD7d93b7fa42b60b6076f3017fCA99b69257A912D",
        "0xf25ee30cfed2d2768C51A6Eb6787890C1c364cA4",
        "0x0D8c517557f3edE116988DD7EC0bAF83b96fe0Cb",
        "0xc225fcd5CE8Ad42863182Ab71acb6abD9C4ddCbE"
    ],
    bridges: ["0x3c78124783E8e39D1E084FdDD0E097334ba2D945",
        "0x76E34d8a527961286E55532620Af5b84F3C6538F",
        "0x68686dB6874588D2404155D00A73F82a50FDd190",
        "0x1533ac4d2922C150551f2F5dc2b0c1eDE382b890"
    ]
}
const mainnetCommittees = {
    beacons: ["0x77b9F481f979e16Cf4234866C0803c2e65106862",
        "0x8F4aEd5Adb0347eF2Db3bDce9f0DF2747D0107E8",
        "0xcf836142D459B1257ed52aef34E62F6F4e7eF894",
        "0x4C265F5eb8Eb68A1eA99626cD838558836438f80",
        "0xc7CC1EE53eF28CC551e97476F0dB01596D945fE0",
        "0xFd119B9DBEb478154E3650F14f55db3787E1bd38",
        "0x780329f064F0BE00FdbDeA4bA8A5b04C7AB2866c"
    ],
    bridges: ["0xaced3cf99897a55057B7513b8505c22DaF9378D9",
        "0x8CbDC490c4188b721e210622027a54E14f27CA7F",
        "0x74d0C0A0A5f89527b4e2850EA09d4F6cE9BBb3bB",
        "0x6964D5c5A7C1503E2228852d1EC115c0e7a20593",
        "0x803c90C23a8a34a676B57CaF0372026C988B416d",
        "0x17E21A7a018046ab3cAE7Aab4215087a2497a7D7",
        "0x5bA8281b5BE1F864E52B3ef8FcEF80560e41005C",
        "0x9Ee3002A85701ae62B16e92e0d8F2044D79a35B6",
        "0xb7eF123cc555cA977aE2fbB5A3690ce57628C664",
        "0x211880118421A814Da0151A4bd06be703DB3654e",
        "0x6A25Ed4Ef6Fa034c895D5721D73dBEC5163Ad4f9",
        "0x1A7232f56F4D71e794D8Bbfc5fa5991d544e1C9f",
        "0xcFcFc3A2CC9DFdF98aC075441E45818C7A70a29e",
        "0x99118446796dFa58d8327834347806711f67Cb79",
        "0x8986acdde31E4519acFcabb139Fd2A2B1da274b2",
        "0xE59C59D87f52D39B1BB8136966e0E1817D7a845A",
        "0x604589220D909878ebDC906d0b33b433Fc3cc0a3",
        "0xf069494c92A85DD31FE6850D8EfE6F2398Ea072c",
        "0x93E1b517726d05c235AE3AF53fa84C34d400Cae9",
        "0x6284C7FD0F623E705d0e0a2D4621299B98eA3895",
        "0xf57Ac7832b1C8F7f5C3E228eF7811D58647A70BF",
        "0x8fa98CBa06b199922E9Acc5749F25FF549e5eEbd"
    ]
};
const testnetCommittees = {
    beacons: ["0x3cD69B1A595B7A9589391538d29ee7663326e4d3",
        "0xc687470342f4E80ECEf6bBd25e276266d40b8429",
        "0x2A40c96b41AdEc5641F28eF923e270B73e29bb53",
        "0x131B772A9ADe1793F000024eAb23b77bEd3BFe64"
    ],
    bridges: ["0x28655822DAf6c4B32303B06e875F92dC6e242cE4",
        "0xD2902ab2F5dF2b17C5A5aa380f511F04a2542E10",
        "0xB67376ad63EAdC22f05efE428e93f09D4f13B4fD",
        "0x40bAA64EAFbD355f5427d127979f377cfA48cc10"
    ]
};

const devProviders = ['http://localhost:9334', 'http://localhost:9338'];

const devMnemonic = 'test test test test test test test test test test test junk';
let readKey = (filename, defaultValue = '') => {
    try {
        const result = fs.readFileSync(`./${filename}`).toString();
        // console.log(`read ${filename}`);
        return result;
    } catch {
        console.log(`WARNING: ${filename} not found`);
        return defaultValue;
    }
}
let mnemonic = readKey('.mnemonic', devMnemonic);
let infuraApiKey = readKey('.infuraKey');
let alchemyApiKey = readKey('.alchemyKey');
let deployerPrivateKey = readKey('.deployerPrivateKey', '0x00');
let vaultAdminPrivateKey = readKey('.vaultAdminPrivateKey', '0x00');


const mainnetForkNetwork = {
    chainId: 31337,
    // can use mainnet keys when present (if so, remove previousVaultAdmin)
    // accounts: [deployerPrivateKey, vaultAdminPrivateKey],
    // reference parameters to use when starting Incognito nodes
    geth_name_conf: '127.0.0.1',
    geth_port_conf: '8545',
    geth_protocol_conf: 'HTTP',
    deployed: {
        kyber: '0x9AAb3f75489902f3a48495025729a0AF77d4b11e', // previously '0x818E6FECD516Ecc3849DAf6845e3EC868087B755',
        kyberEtherAddress: '0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee',
        uniswap: '0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D',
        // configs for tests : tokens [MKR, LINK, DAI]
        testingTokens: ['0x9f8F72aA9304c8B593d555F12eF6589cC3A579A2', '0x514910771af9ca656af840dff83e8264ecf986ca', '0x6b175474e89094c44da98b954eedeac495271d0f'],
        // configs for forked network
        // MKR holder to impersonate in fork
        tokenFunder: '0x05E793cE0C6027323Ac150F6d45C2344d28B6019',
        // mainnet vault admin address to impersonate in fork
        previousVaultAdmin: '0x037ac7fFfC1C52Cf6351e33A77eDBdd14CE35040'
    },
    committees: mainnetCommittees,
    providers: devProviders
}

const networks = {
    hardhat: {
        accounts: {
            mnemonic: devMnemonic,
            count: 8
        },
        forking: {
            url: `https://eth-mainnet.alchemyapi.io/v2/${alchemyApiKey}`,
            blockNumber: 12000000,
            enabled: Boolean(process.env.FORK)
        }
    },
    localhost: process.env.FORK ? mainnetForkNetwork : {
        chainId: 31337,
        geth_name_conf: '127.0.0.1',
        geth_port_conf: '8545',
        geth_protocol_conf: 'HTTP',
        deployed: {
            kyberEtherAddress: '0x0000000000000000000000000000000000000000',
        },
        committees: devCommittees,
        providers: devProviders
    },
    kovan: {
        chainId: 42,
        url: `https://kovan.infura.io/v3/${infuraApiKey}`,
        accounts: {
            mnemonic: mnemonic
        },
        geth_name_conf: `kovan.infura.io/v3/${infuraApiKey}`,
        geth_port_conf: '',
        geth_protocol_conf: 'HTTPS',
        deployed: {
            kyber: '0x692f391bCc85cefCe8C237C01e1f636BbD70EA4D',
            kyberEtherAddress: '0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee',
            uniswap: '0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D',
            // MKR, KNC, DAI
            testingTokens: ['0xaaf64bfcc32d0f15873a02163e7e500671a4ffcd', '0xdB7ec4E4784118D9733710e46F7C83fE7889596a', '0x1f9840a85d5af5bf1d1762f925bdaddc4201f984']
        },
        committees: testnetCommittees,
        providers: devProviders
    },
    mainnet: {
        chainId: 1,
        accounts: [deployerPrivateKey, vaultAdminPrivateKey],
        url: `https://mainnet.infura.io/v3/${infuraApiKey}`,
        geth_name_conf: `mainnet.infura.io/v3/${infuraApiKey}`,
        geth_port_conf: '',
        geth_protocol_conf: 'HTTPS',
        deployed: {
            // make sure this matches vaultAdminPrivateKey
            // previousVaultAdmin: '0x037ac7fFfC1C52Cf6351e33A77eDBdd14CE35040',
            kyber: '0x9AAb3f75489902f3a48495025729a0AF77d4b11e', // previously '0x818E6FECD516Ecc3849DAf6845e3EC868087B755',
            kyberEtherAddress: '0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee',
            uniswap: '0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D',
        },
        committees: mainnetCommittees,
        // providers are only connected to inside Hardhat tests, which do NOT run on mainnet
        providers: devProviders
    }
};

extendEnvironment(hre => {
    // add our custom config to the HRE as a getter function
    hre.networkCfg = () => networks[hre.network.name];
});

// You need to export an object to set up your config
// Go to https://hardhat.org/config/ to learn more

/**
 * @type import('hardhat/config').HardhatUserConfig
 */
module.exports = {
    solidity: "0.6.12",
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
        ethUser: {
            default: 2
        },
        tokenUser: {
            default: 3
        },
        previousVaultAdmin: {
            default: 4
        },
        unshieldSender: {
            default: 5
        }
    },
    mocha: {
        bail: true,
        timeout: 1200000,
    },
    networks,
};