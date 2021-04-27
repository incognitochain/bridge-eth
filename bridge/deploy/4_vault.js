const { deployments, ethers } = require('hardhat');
const { getInstance, confirm, getCodeSlot } = require('../scripts/utils');

module.exports = async({
    getNamedAccounts,
    deployments,
    getChainId,
    getUnnamedAccounts,
}) => {
    const { deploy, log } = deployments;
    let { deployer, vaultAdmin , previousVaultAdmin: prevVaultAdmin } = await getNamedAccounts();
    log(`Actors ${deployer}, ${vaultAdmin}`);
    const getAdmin = async (pVault) => {
        return await pVault.admin();
    }
    const getSuccessor = async (pVault) => {
        return await pVault.successor();
    }

    const newVaultAdminSigner = await ethers.getSigner(vaultAdmin);
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

    const needRestore = Boolean(process.env.RESTORE);
    const restoreAmount = ethers.utils.parseUnits(process.env.RESTORE, 'ether');
    log(`Amount to restore: ${restoreAmount.toString()}`);
    let vaultResult = await deploy('Vault', {
        from: deployer,
        args: [],
        skipIfAlreadyDeployed: true,
        log: true
    });

    const vaultFactory = await ethers.getContractFactory('Vault');
    const vault = await vaultFactory.attach(vaultResult.address);
    let previousVault, needMoving = false, borrowedPermissionFrom = null;
    try {
        previousVault = await ethers.getContract('PrevVault');
        let isPaused = true;
        try {
            isPaused = await previousVault.paused();
        } catch {}
        needMoving = !isPaused;
        // when in fork environment, simulate the real prev admin yielding permission
        if (process.env.FORK) {
            log(`FORK-ENV: "Borrowing" Admin permission and money`);
            const realAdm = await ethers.provider.getSigner(hre.networkCfg().deployed.adminToBorrowFrom);
            const moneyAcc = await ethers.getSigner(deployer);
            await confirm(moneyAcc.sendTransaction({to: prevVaultAdmin, value: ethers.utils.parseUnits('1', 'ether')}));
            await confirm(previousVault.connect(realAdm).retire(prevVaultAdmin));
        }
        const adm = await getAdmin(previousVault);
        const succ = await getSuccessor(previousVault);
        if (prevVaultAdmin != adm) {
            if (prevVaultAdmin == succ) {
                borrowedPermissionFrom = adm;
                log(`Claiming borrowed Admin permission from ${borrowedPermissionFrom}`);
                tx = await confirm(previousVault.connect(prevVaultAdminSigner).claim());
                rc = await tx.wait();
                log(`Gas used: ${rc.gasUsed.toString()}`);
            } else {
                throw 'Need Admin permission to upgrade vault';
            }
        } else {
            log('Proceeding with Admin role');
        }
    } catch (e) {
        log(e);
        log('No previous vault on this network. Skip upgrade...');
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

    let restoringVaultResult = null, restoringVault;
    if (needRestore) {
        restoringVaultResult = await deploy('RestoringVault', {
            from: deployer,
            args: [restoreAmount, proxyResult.address],
            skipIfAlreadyDeployed: true,
            log: true
        });
        restoringVault = await ethers.getContract('RestoringVault');
    }
    log(`DEV : Incognito nodes should use ${proxyResult.address} as EthVaultContract`);
    if (needMoving) {
        let tokenList = hre.networkCfg().tokenList;
        const migrateTask = async (src, dstAddress, tokens) => {
            let depositsBeforeMigrate = await Promise.all(tokens.map(_tokenAddr => src.totalDepositedToSCAmount(_tokenAddr)));
            depositsBeforeMigrate = depositsBeforeMigrate.map(d => d.toString());
            let balancesBeforeMigrate = await Promise.all(tokens.map(_tokenAddr => src.balanceOf(_tokenAddr)));
            balancesBeforeMigrate = balancesBeforeMigrate.map(b => b.toString());
            log(`admin ${prevVaultAdmin} will upgrade ${src.address} to ${dstAddress}`);
            
            try {
                let tx = await confirm(src.connect(prevVaultAdminSigner).pause());
                let rc = await tx.wait();
                log(`Gas used: ${rc.gasUsed.toString()}`);
            } catch(e) {
                log(e);
                log(`Skipping pause() step`);
            }

            try {
                let tx = await confirm(src.connect(prevVaultAdminSigner).migrate(dstAddress));
                let rc = await tx.wait();
                log(`Gas used: ${rc.gasUsed.toString()}`);
            } catch(e) {
                throw e;
            }
            tx = await confirm(src.connect(prevVaultAdminSigner).moveAssets(tokens));
            rc = await tx.wait();
            log(`Gas used: ${rc.gasUsed.toString()}`);
            const theNewVault = await getInstance('Vault', 'TransparentUpgradeableProxy');
            const dummySigner = new ethers.Wallet('0x0000000000000000000000000000000000000055000000000000000000000000', ethers.provider);
            // dummySigner.connect(ethers.provider);
            let deposits = await Promise.all(tokens.map(_tokenAddr => theNewVault.connect(dummySigner).totalDepositedToSCAmount(_tokenAddr)));
            deposits = deposits.map(d => d.toString());
            let balances = await Promise.all(tokens.map(_tokenAddr => theNewVault.connect(dummySigner).balanceOf(_tokenAddr)));
            balances = balances.map(b => b.toString());

            const compare = (arr1, arr2, keys) => {
                let obj = {};
                keys.forEach((k, i) => obj[k] = [arr1[i], arr2[i]]);
                return obj;
            }
            console.log(`Migrated vault ${src.address} to ${dstAddress}, then ${proxyResult.address} with tokens ${tokens}`);
            console.log({"Balance Comparison" : compare(balancesBeforeMigrate, balances, tokens)});
            console.log({"Deposit Comparison" : compare(depositsBeforeMigrate, deposits, tokens)});
        }

        if (needRestore) {
            console.log('Performing restoration of funds');
            const tokensToRestore = [tokenList[0]];
            tokenList = tokenList.slice(1);
            await migrateTask(previousVault, proxyResult.address, tokenList);
            await migrateTask(previousVault, restoringVault.address, tokensToRestore);        
        } else {
            await migrateTask(previousVault, proxyResult.address, tokenList);
        }      
    }
    if (borrowedPermissionFrom) {
        let tx = await confirm(previousVault.connect(prevVaultAdminSigner).migrate(proxyResult.address));
        let rc = await tx.wait();
        log(`Gas used: ${rc.gasUsed.toString()}`);
        log(`Yielding Admin role of both vaults to ${borrowedPermissionFrom}`);
        tx = await confirm(previousVault.connect(prevVaultAdminSigner).retire(borrowedPermissionFrom));
        rc = await tx.wait();
        log(`Gas used: ${rc.gasUsed.toString()}`);
        const theVaultProxy = await getInstance('TransparentUpgradeableProxy');
        tx = await confirm(theVaultProxy.connect(newVaultAdminSigner).retire(borrowedPermissionFrom));
        rc = await tx.wait();
        log(`Gas used: ${rc.gasUsed.toString()}`);
    }
};

module.exports.tags = ['3', 'vault', 'proxy'];
module.exports.dependencies = ['fork', 'incognito-proxy'];