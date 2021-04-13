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

const mainnetTokenList = [
    "0xecF51a98B71f0421151a1d45E033Ab8B88665221", "0xD7F8032777C50aFD2e7AFa41912a4d8038127271", "0x426CA1eA2406c07d75Db9585F22781c096e3d0E0", "0x70Ec7702ADA8530d8f7332f7f3700099553D772D", "0xA587469eE454A0911C5adF761754320319E7bD2F", "0x2505a3c035291c05cb78cb43cff39564637e1dd9", "0xB351dA6ffEbd5DddD1dA037929FCf334d6B4A8D5", "0x2f141ce366a2462f02cea3d12cf93e4dca49e4fd", "0xf34845b76015d2952b6e39436bc59cae3c9ba17d", "0xF1d50e8299687FC17378c6D5e2B25a7a0e07DcB4", "0x5d3dc0fdba0477b906ad4a36f95035b252060434", "0xfF95Ea9eBeFf204A95954bb1Ed76175354914Ea1", "0xe0b9bcd54bf8a730ea5d3f1ffce0885e911a502c", "0x08130635368AA28b217a4dfb68E1bF8dC525621C", "0xc3761eb917cd790b30dad99f6cc5b4ff93c4f9ea", "0x2b591e99afe9f32eaa6214f7b7629768c40eeb39", "0x5228a22e72ccc52d415ecfd199f99d0665e7733b", "0x16ea01acb4b0bca2000ee5473348b6937ee6f72f", "0x66fd97a78d8854fec445cd1c80a07896b0b4851f", "0x2d184014b5658c453443aa87c8e9c4d57285620b", "0x0000000000000000000000000000000000000000", "0x4cc19356f2d37338b9802aa8e8fc58b0373296e7", "0x4f9254c83eb525f9fcf346490bbb3ed28a81c667", "0xead7f3ae4e0bb0d8785852cc37cc9d0b5e75c06a", "0x4575f41308EC1483f3d399aa9a2826d74Da13Deb", "0x799a4202c12ca952cb311598a024c80ed371a41e", "0x607F4C5BB672230e8672085532f7e901544a7375", "0xdac17f958d2ee523a2206206994597c13d831ec7", "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48", "0x6b175474e89094c44da98b954eedeac495271d0f", "0x7d1afa7b718fb893db30a3abc0cfc608aacfebb0", "0x4fabb145d64652a948d72533023f6e7a623c7c53", "0x0d8775f648430679a709e98d2b0cb6250d2887ef", "0x514910771af9ca656af840dff83e8264ecf986ca", "0x89d24a6b4ccb1b6faa2625fe562bdd9a23260359", "0xb7cb1c96db6b22b0d3d9536e0108d062bd488f74", "0x8762db106b2c2a0bccb3a80d1ed41273552616e8", "0x408e41876cccdc0f92210600ef50372656052a38", "0xf0ee6b27b759c9893ce4f094b49ad28fd15a23e4", "0xb63b606ac810a52cca15e44bb630fd42d8d1d83d", "0xdd974d5c2e2928dea5f71b9825b8b646686bd200", "0x8400d94a5cb0fa0d041a3788e395285d61c9ee5e", "0x998FFE1E43fAcffb941dc337dD0468d52bA5b48A", "0x056fd409e1d7a124bd7017459dfea2f387b6d5cd", "0xae746520FfDB15d0505e32f1d6e9a2b4ab866572", "0x41e5560054824ea6b0732e656e3ad64e20e94e45", "0x5432c580e34f590f4dd901b825ddeb92e905e826", "0xb683d83a532e2cb7dfa5275eed3698436371cc9f", "0xd26114cd6EE289AccF82350c8d8487fedB8A0C07", "0x286BDA1413a2Df81731D4930ce2F862a35A609fE", "0xf629cbd94d3791c9250152bd8dfbdf380e2a3b9c", "0x595832f8fc6bf59c85c527fec3740a1b7a361269", "0x1d287cc25dad7ccaf76a26bc660c5f7c8e2a05bd", "0xba11d00c5f74255f56a5e366f4f77f5a186d7f55", "0x6c6ee5e31d828de241282b9606c8e98ea48526e2", "0x8e870d67f660d95d5be530380d0ec0bd388289e1", "0x1f573d6fb3f13d689ff844b4ce37794d79a7ff1c", "0x50d1c9771902476076ecfc8b2a83ad6b9355a4c9", "0x4f3afec4e5a3f2a6a1a411def7d7dfe50ee057bf", "0xbbbbca6a901c926f240b89eacb641d8aec7aeafd", "0xe41d2489571d322189246dafa5ebde1f4699f498", "0x0000000000085d4780B73119b644AE5ecd22b376", "0x1c5857e110cd8411054660f60b5de6a6958cfae2", "0x5d3a536E4D6DbD6114cc1Ead35777bAB948E3643", "0x5C406D99E04B8494dc253FCc52943Ef82bcA7D75", "0x55296f69f40ea6d20e478533c15a6b08b654e758", "0xa0b73e1ff0b80914ab6fe0444e65848c4c34450b", "0x4e15361fd6b4bb609fa63c81a2be19d873717870", "0xAC8491258d2D93228E8b49aAC2e332A96f04E56C", "0x8E9c3D1F30904E91764B7b8bBFDa3a429b886874", "0x71E4B8DE109428f999391eB3424D2CC87192e8bA", "0x91f5f9c36B093907B0F99d1dBcf41aAF1db28916", "0x035df12e0f3ac6671126525f1015e47d79dfeddf", "0x7db5454f3500f28171d1f9c7a38527c9cf94e6b2", "0xAbf14EAc02407842a0140AD012239a03F8985404", "0xf222Ba8Af81d799C565241B0D3eEDF9Bdc4Fc462", "0xc30951ff31c04a47b26ce496b0510a3b2d709e92", "0xcea83bc11a51cf07eea1286eee099871b428e613", "0x3917E933bd430C08304cae2AA6d9746b806406c2", "0xe172f366678ec7b559f6c2913a437baadfd4e6c8", "0xeD79E6dd91324F6Af138f01967BD24233d642a24", "0x1e24f42cb509e2af5675c0f5b529fc0b4c1a112a", "0x9D494A823Fc3E852f8fF92f36A05662A46de0381", "0x159A1dFAe19057de57dFfFcbB3DA1aE784678965", "0x4df76a9dab9bb8310e4ad3dc4336a8e26ed24ebb", "0xaa7fb1c8ce6f18d4fd4aabb61a2193d4d441c54f", "0xaC9Bb427953aC7FDDC562ADcA86CF42D988047Fd", "0xec91fcca41e8ab83dd5bc2bbcc2ffb71e314ba25", "0xf816725650497630642b52dbc3a734e118cf2ed2", "0xc00b89fc3a7ee7043629d8f9a79abfef55634960", "0x3e3Aafa44d6E122b07d329b992F0DF62CF82b1e7", "0x86C31E6da2190a1FFd39A36990a44174D0A8be15", "0xdf413690fb785e56895551cc21086a15b6e90386", "0x719035ac096b12aad44578a2db8a352ad874892d", "0x0f90842e0b39fdc014dfe9daf5835c54ef894bf0", "0x7b53B2C4B2F495d843a4e92e5c5511034d32bd15", "0x716523231368d43BDfe1F06AfE1C62930731aB13", "0x767EE3150Ac31f982190Ef41728Cf9a969355286"
];

const testnetTokenList = [
    '0x0000000000000000000000000000000000000000', '0xE94747D0CDE29064e543B682D2E928795F86bEDe', '0x7ecd5a768add5c4b620c17d169f91c537cec8b4c', '0x42886047bffcc8207688d7f480a569f24779f406', '0x7079f3762805cff9c979a5bdc6f5648bcfee76c8', '0x628796a2d192b3c595cacf03ecaf1f51a78198e3', '0x4c6e1efc12fdfd568186b7baec0a43fffb4bcccf', '0x64f42416317fe7996b0c547d823effb7d8bf2988', '0xe6a51cf10d43365c179ab54b78e0d875b1364dc0', '0x1ff4bc2720e26128287782e3a549879e89714f88', '0x5d661f79c7fb221a962db80ada71e20ccae588e1', '0xafbbb0875bb95c4983253effe4e0fffe1c544633', '0xbe2aea7e0c174aea3027ae88cc33990a960ad512', '0x25e91e333fb7c8993db8ba70ece97a00cf78ee39', '0xd17ae18e75b7f908672caca4b3387910e1603a62', '0x2517dcc7b71d76fa4dbd8bd7d63be2954afc071d', '0x6911a48876cdc65eaed95569e9bd64cb8701aaa2'
];

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
    providers: devProviders,
    tokenList: mainnetTokenList
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
        tokenList: testnetTokenList
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
        tokenList: mainnetTokenList
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