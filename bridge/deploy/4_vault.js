const { deployments, ethers } = require('hardhat');
const { getInstance, confirm } = require('../scripts/utils');

module.exports = async({
    getNamedAccounts,
    deployments,
    getChainId,
    getUnnamedAccounts,
}) => {
    const { deploy, log } = deployments;
    let { deployer, vaultAdmin , previousVaultAdmin: prevVaultAdmin } = await getNamedAccounts();
    // vaultAdmin is a named account in Hardhat config
    let prevVaultAdminSigner;
    // prevVaultAdmin is specified in the "deployed" contracts object in config. It can fallback to vaultAdmin if not found, in case we choose the same EOA as admin
    if (hre.networkCfg().deployed.previousVaultAdmin) {
        // fork network config
        prevVaultAdmin = hre.networkCfg().deployed.previousVaultAdmin;
        prevVaultAdminSigner = await ethers.provider.getSigner(prevVaultAdmin);
    } else if (prevVaultAdmin) {
        // separate admins between Vault versions
        prevVaultAdminSigner = await ethers.getSigner(prevVaultAdmin);
    } else {
        prevVaultAdmin = vaultAdmin;
        prevVaultAdminSigner = await ethers.getSigner(prevVaultAdmin);
    }

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
        let tokenList = hre.networkCfg().tokenList;
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