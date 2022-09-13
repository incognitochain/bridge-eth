task("list-contracts", "Exports & prints the list of deployed contracts")
    .setAction(async (taskArgs, hre) => {
        const filename = `${hre.network.name}-exports.json`;
        await hre.run('export', {export: filename});
        const obj = require(`../${filename}`).contracts;
        let result = {};
        Object.keys(obj).map(k => {
            result[k] = obj[k].address;
        });
        console.log(result);
        const { getImplementation, getInstance, getCodeSlot } = require('./utils');
        let addr = await getImplementation(result.TransparentUpgradeableProxy);
        console.log(`Vault Proxy's current implementation is ${ethers.utils.hexlify(addr)}`);
        // const kyb = await getInstance('KBNTrade');
        // addr = await kyb.kyberNetworkProxyContract();
        // console.log(`KBNTrade is using KyberNetworkRouter ${ethers.utils.hexlify(addr)}`);
        // const uni = await getInstance('UniswapV2Trade');
        // addr = await uni.uniswapV2();
        // console.log(`UniswapV2Trade is using UniswapRouter ${ethers.utils.hexlify(addr)}`);

        // get data from unstructured storage
        addr = await getCodeSlot(result.TransparentUpgradeableProxy, '0x62135fc083646fdb4e1a9d700e351b886a4a5a39da980650269edd1ade91ffd2')
        console.log(`Using IncognitoProxy instance at ${ethers.utils.hexlify(addr)}`);
        addr = await getCodeSlot(result.TransparentUpgradeableProxy, '0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103')
        console.log(`Current Vault Admin : ${ethers.utils.hexlify(addr)}`);
        addr = await getCodeSlot(result.TransparentUpgradeableProxy, '0x7b13fc932b1063ca775d428558b73e20eab6804d4d9b5a148d7cbae4488973f8')
        console.log(`Next Vault Admin : ${ethers.utils.hexlify(addr)}`);
    });

task("unshield-status", "Asks if the contract has processed an unshield")
    .addOptionalParam("id", "The unshield id", "", types.string)
    .setAction(async (taskArgs, hre) => {
        const incTxHash = taskArgs.id;
        const vault = await ethers.getContractAt('Vault', await hre.deployments.get('Vault'));
        const res = await vault.isWithdrawed(incTxHash);
        console.log(res);
    });

task("vault-balance-status", "Balance by token")
    .addOptionalParam("token", "The ERC20 token", "0x0000000000000000000000000000000000000000", types.string)
    .setAction(async taskArgs => {
        const { getInstance } = require('./utils');
        let token = taskArgs.token;
        const vault = await getInstance('Vault', 'TransparentUpgradeableProxy');
        const b = await vault.balanceOf(token);
        const res = await vault.totalDepositedToSCAmount(token);
        console.log(`Contract actual balance ${b.toString()}`);
        console.log(`Contract deposited amount ${res.toString()}`);
    });

task("show-committees", "Asks the IncProxy contract for its committee at height")
    .addOptionalParam("height", "The height", 0, types.int)
    .addOptionalParam("address", "The address for IncProxy", "0x0000000000000000000000000000000000000000", types.string)
    .setAction(async taskArgs => {
        const { getInstance } = require('./utils');
        const ip = await getInstance('IncognitoProxy', null, taskArgs.address);
        const res = await ip.findBeaconCommitteeFromHeight(taskArgs.height);
        console.log(res);
    });

subtask("contract-call-setup", "Unshield (withdraw) from Incognito")
    .setAction(async (taskArgs, hre) => {
        const { inc } = require('./external');
        const { getInstance, confirm, toIncDecimals } = require('./utils');
        const { proveEth, formatBurnProof } = require('./prove');
        const { tokenAddresses } = require('./constants');
        const BN = ethers.BigNumber;
        const signers = await ethers.getSigners();
        const defaultSigner = signers[3];
        await inc.init(null, hre.networkCfg().providers[0], taskArgs.shards ? taskArgs.shards : 8);
        const transactor = taskArgs.privkey ? await inc.NewTransactor(taskArgs.privkey) : null;
        const vault = await getInstance('Vault', 'TransparentUpgradeableProxy');
        const proxy = await getInstance('TransparentUpgradeableProxy');

        return {
            getInstance, confirm, proveEth, formatBurnProof,
            Inc, inc, transactor, tokenAddresses,
            BN, toIncDecimals,
            defaultSigner,
            vault, proxy,
        }
    });

task("shield", "Shield (deposit) to Incognito")
    .addOptionalParam("token", "Token to shield", "0x0000000000000000000000000000000000000000", types.string)
    .addOptionalParam("amount", "Amount to shield (in wei)", "0", types.string)
    .addOptionalParam("to", "Receiver Incognito address", "", types.string)
    .addOptionalParam("privkey", "Your private key (if you're submitting proof from this CLI)", "", types.string)
    .setAction(async taskArgs => {
        let { token, amount, to, privkey } = taskArgs;
        const { Inc, BN, vault, defaultSigner, confirm, transactor, proveEth, tokenAddresses, toIncDecimals, getInstance } = await run('contract-call-setup', { privkey });
        try {
            amount = BN.from(amount);
            const temp = Inc.utils.base58CheckDeserialize(to);
        } catch (e) {
            console.error(e);
            throw 'Invalid Parameter';
        }
        let tx;
        try {
            if (token == tokenAddresses.ETH) {
                tx = await confirm(vault.connect(defaultSigner).deposit(to, {value: amount}), 16);
            } else {
                const tokenInst = await getInstance('contracts/IERC20.sol:IERC20', null, token);
                await confirm(tokenInst.connect(defaultSigner).approve(vault.address, amount));
                tx = await confirm(vault.connect(defaultSigner).depositERC20(token, amount, to), 16);
            }
        } catch(e) {
            console.error('Shielding failed at deposit step');
            throw e;
        }
        if (privkey) {
            if (!tokenAddresses.pTokens[token]) throw `TokenID unavailable for ${token}`;
            console.log('Now submitting deposit proof to issue pETH');
            const proveResult = await proveEth(tx.hash, Inc.utils.base64Encode);
            const issuingResponse = await transactor.shield({ transfer: { fee: 10, tokenID: tokenAddresses.pTokens[token] }, extra: { ethBlockHash: proveResult.ethBlockHash, ethDepositProof: proveResult.encodedProof, txIndex: proveResult.txIndex }});
            let change = await transactor.waitBalanceChange(tokenAddresses.pTokens[token]);
            console.log(`Detected balance change from ${BN.from(change.oldBalance)} to ${BN.from(change.balance)}`);
        }
        console.log('Done !');
        return tx;
    });

task("unshield", "Unshield (withdraw) from Incognito")
    .addOptionalParam("token", "Token to burn", "0x0000000000000000000000000000000000000000", types.string)
    .addOptionalParam("amount", "Amount to burn", "0", types.string)
    .addOptionalParam("to", "Receiver Ethereum address", "", types.string)
    .addOptionalParam("privkey", "Your private key (if you're submitting burn proof from this CLI)", "", types.string)
    .setAction(async taskArgs => {
        try {
            let { token, amount, to, privkey } = taskArgs;
            const { getInstance, toIncDecimals, transactor, Inc, inc, formatBurnProof, defaultSigner, BN, tokenAddresses, confirm, vault } = await run('contract-call-setup', { privkey });
            amount = BN.from(amount);
            if (to.length != 42) to = defaultSigner.address;
            const incAmount = await toIncDecimals(amount, token);
            const burnResult = await transactor.burn({ transfer: { fee: 10, tokenID: tokenAddresses.pTokens[token], info: "BURN TOKEN" }, extra: { remoteAddress: to, burnAmount: incAmount.toString(), burningType: Inc.constants.BurningRequestMeta }});
            await transactor.waitTx(burnResult.Response.txId, 16);
            const burnProof = await inc.rpc.getBurnProof(burnResult.Response.txId, false);
            let params = await formatBurnProof(burnProof);

            let getBalance = (token, addr) => token == tokenAddresses.ETH ? ethers.provider.getBalance(addr) : getInstance('contracts/IERC20.sol:IERC20', null, token).then(t => t.balanceOf(addr));
            const balanceBefore = await getBalance(token, to);
            const tx = await confirm(vault.connect(defaultSigner).withdraw(...params));
            const balance = await getBalance(token, to);
            console.log(`Receiver balance in Ethereum changed from ${balanceBefore.toString()} to ${balance.toString()}`);
            console.log('Done !');
            return tx;
        } catch (e) {
            console.error('Unshield failed')
            throw e;
        }
    });

task("query-var", "Display the vault's Incognito recovery address")
    .addOptionalParam("name", "name of public storage variable to get", "isInitialized")
    .setAction(async taskArgs => {
        const { getInstance } = require('./utils');
        const vault = await getInstance('Vault', 'TransparentUpgradeableProxy');
        const res = await vault[taskArgs.name]();
        console.log(`Vault ${vault.address}: ${taskArgs.name} = ${res}`);
    });

task("testing-balance", "Display the testing accounts' token balances")
    .setAction(async taskArgs => {
        const { getInstance } = require('./utils');
        const signers = await ethers.getSigners();
        const { testingTokenNames, deployed } = hre.networkCfg();
        const { tokenFunders, testingTokens } = deployed || {};  
        console.log(tokenFunders, testingTokens)      
        for (let i=0; i < testingTokens.length; i++) {
            console.log(testingTokenNames[i], '\t', testingTokens[i], 'balances');
            const token = await getInstance('contracts/IERC20.sol:IERC20', null, testingTokens[i]);
            for (const s of signers) {
                const b = await token.balanceOf(s.address);
                console.log('...', s.address, ':', b.toString());
            }
        }
    });

task("vault-evid", "Display the a vault's event ID (topic)")
    .addOptionalParam("name", "event name", "Deposit", types.string)
    .setAction(async taskArgs => {
        const { getInstance } = require('./utils');
        const vault = await getInstance('Vault', 'TransparentUpgradeableProxy');
        const res = await vault.interface.getEventTopic(taskArgs.name);
        console.log(res);
        // console.log(vault.interface.getEvent('0x00b45d95b5117447e2fafe7f34def913ff3ba220e4b8688acf37ae2328af7a3d'));
    });

task("decode-event", "Decode event data by name")
    .addOptionalParam("name", "name", "Redeposit", types.string)
    .addOptionalParam("data", "data", "", types.string)
    .setAction(async (taskArgs, hre) => {
        const [signer, vaultAdmin] = await ethers.getSigners();
        const { getInstance, getImplementation, confirm } = require('./utils');
        const vault = await getInstance('Vault', 'TransparentUpgradeableProxy');
        const evdata = taskArgs.data;
        res = vault.interface.decodeEventLog(taskArgs.name, evdata, [vault.interface.getEventTopic(taskArgs.name)]);
        console.log('decoded', taskArgs.name)
        console.dir(res);
        if (res.errorData) {
            // when errorData is available, parse it as string
            const s = '0x' + res.errorData.slice(10);
            console.log('error', ethers.utils.defaultAbiCoder.decode(['string'], s));
        }
    })

task("decode-submit", "Decode proof submission call")
    .addOptionalParam("name", "name", "executeWithBurnProof", types.string)
    .addOptionalParam("inspector", "inspector function name", "parseCalldataFromBurnInst", types.string)
    .addOptionalParam("data", "data", "", types.string)
    .setAction(async (taskArgs, hre) => {
        const [signer, vaultAdmin] = await ethers.getSigners();
        const { getInstance, getImplementation, confirm } = require('./utils');
        const vault = await getInstance(hre.networkCfg().vaultContractName || 'Vault');
        const params = vault.interface.decodeFunctionData(taskArgs.name, taskArgs.data);
        const res = await vault.functions[taskArgs.inspector](params.inst);
        console.log(taskArgs.name)
        console.dir(res);
    })    

task("inspect-call", "Decode (swap) calldata by name")
    .addOptionalParam("name", "original function name", "tradeInput", types.string)
    .addOptionalParam("inspector", "inspector function name", "_inspectTradeInput", types.string)
    .addOptionalParam("data", "data", "", types.string)
    .setAction(async (taskArgs, hre) => {
        const [signer, vaultAdmin] = await ethers.getSigners();
        const { getInstance, getImplementation, confirm } = require('./utils');
        const c = await getInstance(hre.networkCfg().executorContractName || 'UniswapProxy');
        const params = c.interface.decodeFunctionData(taskArgs.name, taskArgs.data);
        const replayResult = await c.functions[taskArgs.inspector](...params);
        console.log(replayResult);
    })

task("decode-fn-v2", "Decode function data by name")
    .addOptionalParam("name", "name", "trade", types.string)
    .addOptionalParam("data", "data", "", types.string)
    .setAction(async (taskArgs, hre) => {
        const [signer, vaultAdmin] = await ethers.getSigners();
        const { getInstance, getImplementation, confirm } = require('./utils');
        const abi = require('../artifacts/contracts/bsc/pancake_proxy.sol/PancakeProxy.json').abi;
        const iface = new ethers.utils.Interface(abi);
        const fdata = taskArgs.data;
        res = iface.decodeFunctionData(taskArgs.name, fdata);
        console.log('decoded func', taskArgs.name)
        console.dir(res);
        // console.log(Object.keys(res.params))
    })

task("decode-fn", "Decode function data by name")
    .addOptionalParam("name", "name", "exactInputSingle", types.string)
    .addOptionalParam("data", "data", "", types.string)
    .setAction(async (taskArgs, hre) => {
        const [signer, vaultAdmin] = await ethers.getSigners();
        const { getInstance, getImplementation, confirm } = require('./utils');
        const abi = require('../artifacts/contracts/polygon/uniswap_proxy.sol/ISwapRouter2.json').abi;
        const iface = new ethers.utils.Interface(abi);
        const fdata = taskArgs.data;
        res = iface.decodeFunctionData(taskArgs.name, fdata);
        console.log('decoded func', taskArgs.name)
        console.dir(res);
        console.log(Object.keys(res.params))
    })

task("func-id", "Display function ID by name")
    .addOptionalParam("name", "function name", "Deposit", types.string)
    .setAction(async taskArgs => {
        const { getInstance } = require('./utils');
        // const iface = await ethers.getContractFactory(hre.networkCfg().executorContractName || 'UniswapProxy');
        const abi = require('../artifacts/contracts/polygon/uniswap_proxy.sol/ISwapRouter2.json').abi;
        const iface = new ethers.utils.Interface(abi);
        const res = await iface.getSighash(taskArgs.name);
        console.log(res);
    });

task("vault-balance", "Display the vault's ether balance")
    .addOptionalParam("token", "Token to burn", "0x0000000000000000000000000000000000000000", types.string)
    .setAction(async taskArgs => {
        const { getInstance } = require('./utils');
        const vault = await getInstance('Vault', 'TransparentUpgradeableProxy');
        const res = await vault.balanceOf(taskArgs.token);
        const bn = await ethers.provider.getBlockNumber();
        console.log(`Vault ${vault.address} at block #${bn.toString()}: ${res.toString()} gwei of ETH`);
    });

task("send-test", "Display the a vault's event ID (topic)")
    .addOptionalParam("privkey", "Your private key (if you're submitting burn proof from this CLI)", "", types.string)
    .setAction(async (taskArgs, hre) => {
        const [signer, vaultAdmin] = await ethers.getSigners();
        const { getInstance, getImplementation, confirm } = require('./utils');
        // const vault = await getInstance('VaultPLG', 'TransparentUpgradeableProxy');
        // const proxy = await getInstance('TransparentUpgradeableProxy');
        const inputData = '0x3ed1b376000000000000000000000000000000000000000000000000000000000000014000000000000000000000000000000000000000000000000000000000005165c500000000000000000000000000000000000000000000000000000000000003c00000000000000000000000000000000000000000000000000000000000000460c9b0ef63bf2e885982bcb7948b97252586df561ea4d3d73ef54278dd46487117720198ef11dae07ce39e124ff6803ee1d85c0760de86033b0b6286d0eb11cac00000000000000000000000000000000000000000000000000000000000000500000000000000000000000000000000000000000000000000000000000000058000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000680000000000000000000000000000000000000000000000000000000000000024c9e010200000000000000000000000084b9b910527ad5c03a9ca831909e21e236ea7b060000000000000000000000000e2923c21e2c5a2bdd18aa460b3fddddadb0ae1800000000000000000000000000000000000000000000000000000000000f4240f4c6eaca0907a34780afa5a0b2d077746f18e587a139b0b36d7fcf88df1ec1ff00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000045c9a9a53cdbce19576d06057df4ab9b7fa6195e1e630ce83489323418fa589c09490a0e86a186ad88cb360a49eb34b4e32cab1be10818a812e81128fdd3a46a60000007fa6b5ba2c76840a50fe7cd4fb437f71484f07547322e7883cfdbe83d01734f2c471e13d2400000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000e8d4a51000000000000000000000000000000000000000000000000000000e062addd6f6f300000000000000000000000000000000000000000000000000000000632123520000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000300000000000000000000000084b9b910527ad5c03a9ca831909e21e236ea7b0600000000000000000000000078867bbeef44f2326bf8ddd1941a4439382ef2a7000000000000000000000000ae13d989dac2f0debff460ac112a837c89baa7cd00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004f2b402d02a89e8fa73bb0690f358a0509c16f8b6016fe6275538232d63b10d8938e7497ffe223f662f9f54f555b7dbbede3412b10301e8b0fe0d65c3a0c48ff23f02ce6e890b53fe2a2d81dc9694f87ff5ce38bcd5e9790ad9419341d75a9432449547208dd1b0e6889f36cbaa6a34dd9718863b222883ce1cc60136db52f0450000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000030000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000030000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000001b000000000000000000000000000000000000000000000000000000000000001c000000000000000000000000000000000000000000000000000000000000001b00000000000000000000000000000000000000000000000000000000000000039d9d077a7a5caa7df8eefcb22fde14a4c64dea7fa7b4994d07d8ffb7205d1171d161b7ec6f1ebfe0c3aa06ea97be599ee16b6495c0e5d005c161a8ad9d92c01d757b80cfce340a401008b035afd0ab6d168ec3427a833d6b2bc3b53715d6e8a9000000000000000000000000000000000000000000000000000000000000000372bffd284cf20f12194a2b2ea5001034addc5c3ecf1c7cdf2e28937ae7819b31573ed27d5ecd82bb6091d33b114f8438d328cf73986a5860e2285adf2f82592331ea35ea6e2440a7590b64fde6fc6eefeabb8fdf0b7fc748a952fc5856d8cfc8';
        const extCalldata = '0x5ae401dc000000000000000000000000000000000000000000000000000000009eb48eb6000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000000e404e45aaf000000000000000000000000a6fa4fb5f76172d178d61b04b0ecd319c5d1c0aa0000000000000000000000009c3c9283d3e44854697cd22d3faa240cfb03288900000000000000000000000000000000000000000000000000000000000001f40000000000000000000000001d413f78f061209dfe1d8a71bf893939c36f6d7f000000000000000000000000000000000000000000000000000000a2fb4058000000000000000000000000000000000000000000000000000000547517447aa100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000e404e45aaf000000000000000000000000a6fa4fb5f76172d178d61b04b0ecd319c5d1c0aa0000000000000000000000009c3c9283d3e44854697cd22d3faa240cfb0328890000000000000000000000000000000000000000000000000000000000000bb80000000000000000000000001d413f78f061209dfe1d8a71bf893939c36f6d7f00000000000000000000000000000000000000000000000000000045d964b80000000000000000000000000000000000000000000000000000004985c940b621000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000';
        const { deploy } = hre.deployments;
        // await deploy('VaultPLGTest', { from: signer.address, args: [], log: true, skipIfAlreadyDeployed: true });
        // const tv = await getInstance('VaultPLGTest');
        // const cur = await getImplementation(vault.address);
        // console.log(cur, 'to', tv.address);
        // p 0x76318093c374e39B260120EBFCe6aBF7f75c8D28
        //  0x64b464037ef0aa3d1a95a5c04bc77e8667870e90
        // await confirm(proxy.connect(vaultAdmin).upgradeTo(tv.address));
        // await signer.sendTransaction({ value: ethers.utils.parseUnits('0.4', 'ether'), to: vault.address });
        // const uni = await getInstance('UniswapV2Trade', null, '0x0c61c7f99dee0270a5934f1520d109dc44a51d6c');
        // const tx = await signer.sendTransaction({ value: 0, to: vault.address, data: inputData });
        const tok = await getInstance('contracts/IERC20.sol:IERC20', null, '0xA6FA4fB5f76172d178d61B04b0ecd319C5d1C0aa'); // WETH
        // to: '0x9093BeC7D4a1666F0fD316072C3feA57E6521a40',
        // token: '0xA6FA4fB5f76172d178d61B04b0ecd319C5d1C0aa',
        // amount: BigNumber { _hex: '0xe8d4a51000', _isBigNumber: true },
        // data: '0x5ae401dc00000000000000000000000000000000000000000000000000000000630e873d000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000000e404e45aaf000000000000000000000000a6fa4fb5f76172d178d61b04b0ecd319c5d1c0aa0000000000000000000000009c3c9283d3e44854697cd22d3faa240cfb03288900000000000000000000000000000000000000000000000000000000000001f40000000000000000000000001d413f78f061209dfe1d8a71bf893939c36f6d7f000000000000000000000000000000000000000000000000000000a2fb4058000000000000000000000000000000000000000000000000000000547517447aa100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000e404e45aaf000000000000000000000000a6fa4fb5f76172d178d61b04b0ecd319c5d1c0aa0000000000000000000000009c3c9283d3e44854697cd22d3faa240cfb0328890000000000000000000000000000000000000000000000000000000000000bb80000000000000000000000001d413f78f061209dfe1d8a71bf893939c36f6d7f00000000000000000000000000000000000000000000000000000045d964b80000000000000000000000000000000000000000000000000000004985c940b621000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000'
        // const appr = await confirm(tok.approve('0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45', '0xe8d4a51000'));
        const tx = await signer.sendTransaction({ value: 0, to: '0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45', data: extCalldata, gasLimit: 2000000 });
        // const tx = await signer.sendTransaction({ value: ethers.utils.parseUnits('0.001', 'ether'), to: signer.address });
        // console.log(tx.hash)
        const r = await tx.wait();
        console.dir(r, {depth: null})
        // console.log(r.logs.map(l => l.topics))        
        // console.log(b)
    });
