require("@nomiclabs/hardhat-ethers");
require("@nomiclabs/hardhat-waffle");
// require("@nomiclabs/hardhat-web3");
require("hardhat-deploy");
// require("hardhat-deploy-ethers");

const fs = require('fs');

// This is a sample Hardhat task. To learn how to create your own go to
// https://hardhat.org/guides/create-task.html
task("accounts", "Prints the list of accounts", async () => {
    const accounts = await ethers.getSigners();

    for (const account of accounts) {
        console.log(account.address);
    }
});
require('./scripts/tasks');

const devCommittees = {
    beacons: [
        "0x3cD69B1A595B7A9589391538d29ee7663326e4d3"
    ],
    bridges: [
        "0xD2902ab2F5dF2b17C5A5aa380f511F04a2542E10"
    ],
};
const devCommitteesBig = {
    beacons: [
        "0x3cD69B1A595B7A9589391538d29ee7663326e4d3",
        "0xc687470342f4E80ECEf6bBd25e276266d40b8429",
        "0x2A40c96b41AdEc5641F28eF923e270B73e29bb53",
        "0x131B772A9ADe1793F000024eAb23b77bEd3BFe64",
    ],
    bridges: ["0x3c78124783E8e39D1E084FdDD0E097334ba2D945",
        "0x76E34d8a527961286E55532620Af5b84F3C6538F",
        "0x68686dB6874588D2404155D00A73F82a50FDd190",
        "0x1533ac4d2922C150551f2F5dc2b0c1eDE382b890"
    ]
}
const mainnetCommittees = {
    beacons: [
        '0xe1fe6bdb4FB5f80801D242480c5467c1de94719c',
        '0xD57Dc32f9753a20Af166F9Dc48dE22355A9F7c83',
        '0x44b39171D742C2CdFdA0EBb6226f8584CAfc57FC',
        '0x4c8b59d24f07192B9095DA1EAE9af5c890413A71',
        '0x6d678311c5DAf5F8c8c48223C7Aea2A49D8d8B12',
        '0x93114859F53F98dC2a1FA6be9340Ce3B1D74722B',
        '0x0c7d24b75bEc5E94924e8e5d6c793609e48e7FF6',
    ],
    bridges: []
};
const testnetCommittees = {
    beacons: [
        '0x7ef17C60cAa1c5C43d2Af62726c8f7c14000AB02',
        '0xFe30C03E5Db66236a82b0Dd204E811444Ba7761E',
        '0xa357789d21e217FE3a30c7320A867B8B47C793A4',
        '0xcc817963abe49569Ac58f1BC047B38cDA95832a1',
    ],
    bridges: []
};

const devProviders = ['http://localhost:8334', 'http://localhost:8334'];

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

const mainnetTokenList = [];

const testnetTokenList = [];

const mainnetForkNetwork = {
    chainId: 31337,
    accounts: [deployerPrivateKey, vaultAdminPrivateKey],
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
        previousVaultAdmin: '0xE2516f0F38d9400a8ceCe5672884De72FBD01cE1',
        adminToBorrowFrom: '0x037ac7fFfC1C52Cf6351e33A77eDBdd14CE35040'
    },
    committees: mainnetCommittees,
    providers: devProviders,
    numShards: 2,
};
const forkcfg = {
    plg: {
        forking: {
            url: `https://polygon-mumbai.g.alchemy.com/v2/...`,
            blockNumber: 27826242,
        },
        accounts: {
            mnemonic: mnemonic,
            count: 4
        },
        deployed: {
            testingTokens: ['0x0b3F868E0BE5597D5DB7fEB59E1CADBb0fdDa50a', '0x2C89bbc92BD86F8075d1DEcc58C7F4E0107f286b', '0xb33EaAd8d922B1083446DC23f610c2567fB5180f'], 
            tokenFunders: ['0x0f0c716b007c289c0011e470cc7f14de4fe9fc80', '0xd839db910fd2ae169f4d53a84eb182a2e50760c3', '0x06959153b974d0d5fdfd87d561db6d8d4fa0bb0b']
        },
        testingTokenNames: ['SUSHI', 'AVAX', 'USDT'],
        committees: devCommitteesBig,
        providers: devProviders,
        numShards: 2,
        vaultContractName: 'VaultPLG',
    },
    eth: {
        forking: {
            url: `https://eth-mainnet.g.alchemy.com/v2/...`,
            blockNumber:  15414100,
        },
        accounts: {
            mnemonic: mnemonic,
            count: 4
        },
        deployed: {
            kyber: '0xdeFA4e8a7bcBA345F687a2f1456F5Edd9CE97202',
            uniswap: '0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45', // swap router 02
            testingTokens: ['0x6B175474E89094C44Da98b954EedeAC495271d0F', '0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2', '0x514910771AF9Ca656af840dff83E8264EcF986CA'], 
            tokenFunders: ['0xf977814e90da44bfa03b6295a0616a897441acec', '0x8eb8a3b98659cce290402893d0123abb75e3ab28', '0xf977814e90da44bfa03b6295a0616a897441acec']
        },
        testingTokenNames: ['DAI', 'WETH', 'LINK'],
        committees: devCommitteesBig,
        providers: devProviders,
        numShards: 2,
        vaultContractName: 'Vault',
    },
    bsc: {
        forking: {
            url: `https://bsc-dataseed.binance.org`,
            blockNumber:  20690000,
        },
        accounts: {
            mnemonic: mnemonic,
            count: 4
        },
        deployed: {
            testingTokens: ['0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82', '0x9f589e3eabe42ebC94A44727b3f3531C0c877809', '0xbA2aE424d960c26247Dd6c32edC70B295c744C43'], 
            tokenFunders: ['0xf977814e90da44bfa03b6295a0616a897441acec', '0xf977814e90da44bfa03b6295a0616a897441acec', '0xf977814e90da44bfa03b6295a0616a897441acec']
        },
        testingTokenNames: ['CAKE', 'TKO', 'DOGE'],
        committees: devCommitteesBig,
        providers: devProviders,
        numShards: 2,
        vaultContractName: 'VaultBSC',
    }
}


const networks = {
    hardhat: Object.assign({
        accounts:
        {
            mnemonic: mnemonic,
            count: 8
        },
        // forking: {
        //     url: forkcfg[process.env.FORK]?.furl,
        //     blockNumber: forkcfg[process.env.FORK]?.blockNumber,
        // },
        committees: devCommitteesBig,
        providers: devProviders,
        chainId: 31337,
        numShards: 2,
    }, process.env.FORK ? forkcfg[process.env.FORK] : {}),
    localhost: process.env.FORK ? forkcfg[process.env.FORK] : {
        committees: devCommitteesBig,
        providers: devProviders,
        chainId: 31337,
        numShards: 2,
        deployed: {
        }
    },
    kovan: {
        chainId: 42,
        url: `https://kovan.infura.io/v3/${infuraApiKey}`,
        accounts: [deployerPrivateKey, vaultAdminPrivateKey],
        // {
        //     mnemonic: mnemonic
        // },
        geth_name_conf: `kovan.infura.io/v3/${infuraApiKey}`,
        geth_port_conf: '',
        geth_protocol_conf: 'HTTPS',
        deployed: {
            kyber: '0x692f391bCc85cefCe8C237C01e1f636BbD70EA4D',
            kyberEtherAddress: '0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee',
            uniswap: '0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D',
            // MKR, KNC, UNI
            testingTokens: ['0xaaf64bfcc32d0f15873a02163e7e500671a4ffcd', '0xdB7ec4E4784118D9733710e46F7C83fE7889596a', '0x1f9840a85d5af5bf1d1762f925bdaddc4201f984']
        },
        committees: testnetCommittees,
        providers: devProviders,
        numShards: 2,
    },
    goerli: {
        url: `https://eth-goerli.g.alchemy.com/v2/${alchemyApiKey}`,
        accounts: {
            mnemonic: mnemonic
        },
        deployed: {
            kyber: '0x692f391bCc85cefCe8C237C01e1f636BbD70EA4D',
            kyberEtherAddress: '0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee',
            uniswap: '0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D',
            // MKR, KNC, UNI
            // testingTokens: ['0xaaf64bfcc32d0f15873a02163e7e500671a4ffcd', '0xdB7ec4E4784118D9733710e46F7C83fE7889596a', '0x1f9840a85d5af5bf1d1762f925bdaddc4201f984']
        },
        committees: devCommitteesBig,
        providers: devProviders,
        numShards: 2,
    },
    bsctest: {
        url: "https://data-seed-prebsc-1-s1.binance.org:8545",
        accounts: {
            mnemonic: mnemonic
        },
        deployed: {},
        committees: devCommitteesBig,
        providers: devProviders,
        numShards: 2,
        vaultContractName: 'VaultBSC',
    },
    mumbai: {
        url: "https://polygon-mumbai.g.alchemy.com/v2/...",
        accounts: {
            mnemonic: mnemonic
        },
        deployed: {
            uniswap: '0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45',
            kyber: '0xddff'
        },
        committees: devCommitteesBig,
        providers: devProviders,
        numShards: 2,
        vaultContractName: 'VaultPLG',
    },
    ftmtest: {
        url: "https://rpc.testnet.fantom.network",
        accounts: {
            mnemonic: mnemonic
        },
        deployed: {},
        committees: devCommitteesBig,
        providers: devProviders,
        numShards: 2,
        vaultContractName: 'VaultFTM',
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
        providers: devProviders,
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
    solidity: {
        compilers: [
            {
                version: "0.6.12",
                settings: {
                    optimizer: {
                        enabled: true,
                        runs: 1000
                    }
                },
            },
            {
                version: "0.7.6",
                settings: {
                    optimizer: {
                        enabled: true,
                        runs: 1000
                    }
                },
            },
            {
                version: "0.8.9",
                settings: {
                    optimizer: {
                        enabled: true,
                        runs: 1000
                    }
                },
            },
        ]
    },
    paths: {
        tests: './tests'
    },
    mocha: {
        bail: true,
        timeout: 12000000,
    },
    networks,
    namedAccounts: {
        deployer: {
            default: 0
        },
        vaultAdmin: {
            // can fallback to deployer key
            default: vaultAdminPrivateKey == deployerPrivateKey ? 0 : 1
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
            default: 5,
            goerli: 0,
        }
    },
};