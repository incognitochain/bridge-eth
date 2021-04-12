const { deployments, ethers } = require('hardhat');
const { getInstance, confirm } = require('../scripts/utils');

let tokenList = [
    "0xecF51a98B71f0421151a1d45E033Ab8B88665221", "0xD7F8032777C50aFD2e7AFa41912a4d8038127271", "0x426CA1eA2406c07d75Db9585F22781c096e3d0E0", "0x70Ec7702ADA8530d8f7332f7f3700099553D772D", "0xA587469eE454A0911C5adF761754320319E7bD2F", "0x2505a3c035291c05cb78cb43cff39564637e1dd9", "0xB351dA6ffEbd5DddD1dA037929FCf334d6B4A8D5", "0x2f141ce366a2462f02cea3d12cf93e4dca49e4fd", "0xf34845b76015d2952b6e39436bc59cae3c9ba17d", "0xF1d50e8299687FC17378c6D5e2B25a7a0e07DcB4", "0x5d3dc0fdba0477b906ad4a36f95035b252060434", "0xfF95Ea9eBeFf204A95954bb1Ed76175354914Ea1", "0xe0b9bcd54bf8a730ea5d3f1ffce0885e911a502c", "0x08130635368AA28b217a4dfb68E1bF8dC525621C", "0xc3761eb917cd790b30dad99f6cc5b4ff93c4f9ea", "0x2b591e99afe9f32eaa6214f7b7629768c40eeb39", "0x5228a22e72ccc52d415ecfd199f99d0665e7733b", "0x16ea01acb4b0bca2000ee5473348b6937ee6f72f", "0x66fd97a78d8854fec445cd1c80a07896b0b4851f", "0x2d184014b5658c453443aa87c8e9c4d57285620b", "0x0000000000000000000000000000000000000000", "0x4cc19356f2d37338b9802aa8e8fc58b0373296e7", "0x4f9254c83eb525f9fcf346490bbb3ed28a81c667", "0xead7f3ae4e0bb0d8785852cc37cc9d0b5e75c06a", "0x4575f41308EC1483f3d399aa9a2826d74Da13Deb", "0x799a4202c12ca952cb311598a024c80ed371a41e", "0x607F4C5BB672230e8672085532f7e901544a7375", "0xdac17f958d2ee523a2206206994597c13d831ec7", "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48", "0x6b175474e89094c44da98b954eedeac495271d0f", "0x7d1afa7b718fb893db30a3abc0cfc608aacfebb0", "0x4fabb145d64652a948d72533023f6e7a623c7c53", "0x0d8775f648430679a709e98d2b0cb6250d2887ef", "0x514910771af9ca656af840dff83e8264ecf986ca", "0x89d24a6b4ccb1b6faa2625fe562bdd9a23260359", "0xb7cb1c96db6b22b0d3d9536e0108d062bd488f74", "0x8762db106b2c2a0bccb3a80d1ed41273552616e8", "0x408e41876cccdc0f92210600ef50372656052a38", "0xf0ee6b27b759c9893ce4f094b49ad28fd15a23e4", "0xb63b606ac810a52cca15e44bb630fd42d8d1d83d", "0xdd974d5c2e2928dea5f71b9825b8b646686bd200", "0x8400d94a5cb0fa0d041a3788e395285d61c9ee5e", "0x998FFE1E43fAcffb941dc337dD0468d52bA5b48A", "0x056fd409e1d7a124bd7017459dfea2f387b6d5cd", "0xae746520FfDB15d0505e32f1d6e9a2b4ab866572", "0x41e5560054824ea6b0732e656e3ad64e20e94e45", "0x5432c580e34f590f4dd901b825ddeb92e905e826", "0xb683d83a532e2cb7dfa5275eed3698436371cc9f", "0xd26114cd6EE289AccF82350c8d8487fedB8A0C07", "0x286BDA1413a2Df81731D4930ce2F862a35A609fE", "0xf629cbd94d3791c9250152bd8dfbdf380e2a3b9c", "0x595832f8fc6bf59c85c527fec3740a1b7a361269", "0x1d287cc25dad7ccaf76a26bc660c5f7c8e2a05bd", "0xba11d00c5f74255f56a5e366f4f77f5a186d7f55", "0x6c6ee5e31d828de241282b9606c8e98ea48526e2", "0x8e870d67f660d95d5be530380d0ec0bd388289e1", "0x1f573d6fb3f13d689ff844b4ce37794d79a7ff1c", "0x50d1c9771902476076ecfc8b2a83ad6b9355a4c9", "0x4f3afec4e5a3f2a6a1a411def7d7dfe50ee057bf", "0xbbbbca6a901c926f240b89eacb641d8aec7aeafd", "0xe41d2489571d322189246dafa5ebde1f4699f498", "0x0000000000085d4780B73119b644AE5ecd22b376", "0x1c5857e110cd8411054660f60b5de6a6958cfae2", "0x5d3a536E4D6DbD6114cc1Ead35777bAB948E3643", "0x5C406D99E04B8494dc253FCc52943Ef82bcA7D75", "0x55296f69f40ea6d20e478533c15a6b08b654e758", "0xa0b73e1ff0b80914ab6fe0444e65848c4c34450b", "0x4e15361fd6b4bb609fa63c81a2be19d873717870", "0xAC8491258d2D93228E8b49aAC2e332A96f04E56C", "0x8E9c3D1F30904E91764B7b8bBFDa3a429b886874", "0x71E4B8DE109428f999391eB3424D2CC87192e8bA", "0x91f5f9c36B093907B0F99d1dBcf41aAF1db28916", "0x035df12e0f3ac6671126525f1015e47d79dfeddf", "0x7db5454f3500f28171d1f9c7a38527c9cf94e6b2", "0xAbf14EAc02407842a0140AD012239a03F8985404", "0xf222Ba8Af81d799C565241B0D3eEDF9Bdc4Fc462", "0xc30951ff31c04a47b26ce496b0510a3b2d709e92", "0xcea83bc11a51cf07eea1286eee099871b428e613", "0x3917E933bd430C08304cae2AA6d9746b806406c2", "0xe172f366678ec7b559f6c2913a437baadfd4e6c8", "0xeD79E6dd91324F6Af138f01967BD24233d642a24", "0x1e24f42cb509e2af5675c0f5b529fc0b4c1a112a", "0x9D494A823Fc3E852f8fF92f36A05662A46de0381", "0x159A1dFAe19057de57dFfFcbB3DA1aE784678965", "0x4df76a9dab9bb8310e4ad3dc4336a8e26ed24ebb", "0xaa7fb1c8ce6f18d4fd4aabb61a2193d4d441c54f", "0xaC9Bb427953aC7FDDC562ADcA86CF42D988047Fd", "0xec91fcca41e8ab83dd5bc2bbcc2ffb71e314ba25", "0xf816725650497630642b52dbc3a734e118cf2ed2", "0xc00b89fc3a7ee7043629d8f9a79abfef55634960", "0x3e3Aafa44d6E122b07d329b992F0DF62CF82b1e7", "0x86C31E6da2190a1FFd39A36990a44174D0A8be15", "0xdf413690fb785e56895551cc21086a15b6e90386", "0x719035ac096b12aad44578a2db8a352ad874892d", "0x0f90842e0b39fdc014dfe9daf5835c54ef894bf0", "0x7b53B2C4B2F495d843a4e92e5c5511034d32bd15", "0x716523231368d43BDfe1F06AfE1C62930731aB13", "0x767EE3150Ac31f982190Ef41728Cf9a969355286"
    ];

module.exports = async({
    getNamedAccounts,
    deployments,
    getChainId,
    getUnnamedAccounts,
}) => {
    const { deploy, log } = deployments;
    let { deployer, vaultAdmin } = await getNamedAccounts();
    // vaultAdmin is a named account in Hardhat config
    let prevVaultAdmin = vaultAdmin;
    // prevVaultAdmin is specified in the "deployed" contracts object in config. It can fallback to vaultAdmin if not found, in case we choose the same EOA as admin
    if (hre.networkCfg().deployed.previousVaultAdmin) {
        prevVaultAdmin = hre.networkCfg().deployed.previousVaultAdmin;
    }
    let prevVaultAdminSigner = await ethers.provider.getSigner(prevVaultAdmin);

    const ip = await deployments.get('IncognitoProxy');
    let vaultResult = await deploy('Vault', {
        from: deployer,
        args: [],
        skipIfAlreadyDeployed: true,
        log: true
    });

    const vaultFactory = await ethers.getContractFactory('Vault');
    const vault = await vaultFactory.attach(vaultResult.address);
    let previousVault, needMoving = false;
    try {
        previousVault = await ethers.getContract('PrevVault');
        let isPaused = true;
        try {
            isPaused = await previousVault.paused();
        } catch {}
        needMoving = !isPaused;
    } catch (e) {
        previousVault = {address: '0x0000000000000000000000000000000000000000'};
    }
    const initializeData = vaultFactory.interface.encodeFunctionData('initialize', [previousVault.address]);
    log('will deploy proxy & upgrade with params', vault.address, vaultAdmin, ip.address, initializeData);
    let proxyResult = await deploy('TransparentUpgradeableProxy', {
        from: deployer,
        args: [vault.address, vaultAdmin, ip.address, initializeData],
        skipIfAlreadyDeployed: true,
        log: true,
    });
    log(`DEV : Incognito nodes should use ${proxyResult.address} as EthVaultContract`);
    if (needMoving) {
        let depositsBeforeMigrate = await Promise.all(tokenList.map(_tokenAddr => previousVault.totalDepositedToSCAmount(_tokenAddr)));
        depositsBeforeMigrate = depositsBeforeMigrate.map(d => d.toString());
        let balancesBeforeMigrate = await Promise.all(tokenList.map(_tokenAddr => previousVault.balanceOf(_tokenAddr)));
        balancesBeforeMigrate = balancesBeforeMigrate.map(b => b.toString());
        log(`admin ${prevVaultAdmin} will upgrade ${previousVault.address} to ${proxyResult.address}`);
        let tx = await confirm(previousVault.connect(prevVaultAdminSigner).pause());
        let rc = await tx.wait();
        log(`Gas used: ${rc.gasUsed.toString()}`);
        tx = await confirm(previousVault.connect(prevVaultAdminSigner).migrate(proxyResult.address));
        rc = await tx.wait();
        log(`Gas used: ${rc.gasUsed.toString()}`);
        tx = await confirm(previousVault.connect(prevVaultAdminSigner).moveAssets(tokenList));
        rc = await tx.wait();
        log(`Gas used: ${rc.gasUsed.toString()}`);

        const theNewVault = await getInstance('Vault', 'TransparentUpgradeableProxy');
        let deposits = await Promise.all(tokenList.map(_tokenAddr => theNewVault.totalDepositedToSCAmount(_tokenAddr)));
        deposits = deposits.map(d => d.toString());
        let balances = await Promise.all(tokenList.map(_tokenAddr => theNewVault.balanceOf(_tokenAddr)));
        balances = balances.map(b => b.toString());
        const compare = (arr1, arr2, keys) => {
            let obj = {};
            keys.forEach((k, i) => obj[k] = [arr1[i], arr2[i]]);
            return obj;
        }
        console.log({"Balance Comparison" : compare(balancesBeforeMigrate, balances, tokenList)});
        console.log({"Deposit Comparison" : compare(depositsBeforeMigrate, deposits, tokenList)});
    }
};

module.exports.tags = ['3', 'vault', 'proxy'];
module.exports.dependencies = ['fork', 'incognito-proxy'];