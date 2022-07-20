// const timeMachine = require('ganache-time-traveler');
const { tokenAddresses } = require('./constants');
const hre = require('hardhat');
const { ethers, deployments } = hre;
const BN = ethers.BigNumber;
// const eth = web3.eth;

let chooseOneFrom = (arr) => {
    if (arr.length == 0) throw 'no element left to choose';
    let clonedArr = arr.map(x => x);
    const num = BN.from(ethers.utils.randomBytes(6)).toNumber();
    return [ clonedArr.splice(num % arr.length, 1)[0], clonedArr ];
}

let getPartOf = (amount, cap) => {
    cap = BN.from(cap);
    let num = BN.from(amount);
    let ahundred = BN.from(100);
    // range from 1 to 100 percent, then multiply by cap%
    let percent = BN.from(ethers.utils.randomBytes(6));
    percent = percent.mod(ahundred).add(1);
    const result = num.mul(percent).div(100).mul(cap).div(100);
    // console.log(amount, 'decay by', percent.toString(), '% =>', result.toString());
    return result;
}

// we call tx.wait for block confirmations, but additionally handle for local network which does not have a defined block time
let confirm = async (txp, target = 2, network = 'development') => {
    const tx = await txp;
    console.log(`${tx.hash} => waiting for ${target} confirmations`);
    if (hre.network.name == 'localhost') {
        for (let i=0; i<=target; i++) {
            await hre.network.provider.request({
                method: "evm_mine",
                params: []
            })
        }
    }
    await tx.wait(target);
    return tx;
}

let getDecimals = (_addr) =>
    // amounts of ETH will not have their decimals changed when shielded to Incognito, so we pretend ETH's decimal is 9 here
    (!_addr || _addr==tokenAddresses.ETH) ? Promise.resolve(BN.from(18)) : ethers.getContractAt('contracts/IERC20.sol:IERC20', _addr).then(_c => _c.decimals())

let toIncDecimals = (_amount, _addr,  _useUnified = false) => getDecimals(_addr)
    .then(_d => {
        let result = BN.from(_amount);
        if (!_useUnified && _d.lte(9)) return result;
        const ten = BN.from(10);
        // console.log(`${result} * ${ten.pow(_d.subn(9))}`);
        return result.mul(ten.pow(9)).div(ten.pow(_d));
    })

let fromIncDecimals = (_amount, _addr, _useUnified = false) => getDecimals(_addr)
    .then(_d => {
        let result = BN.from(_amount);
        if (!_useUnified && _d.lte(9)) return result;
        const ten = BN.from(10);
        // console.log(`${result} / ${ten.pow(_d.subn(9))}`);
        return result.mul(ten.pow(_d)).div(ten.pow(9));
    })

// useful for binding Vault ABI to proxy address, or get an IERC20 token contract instance...
let getInstance = async (abiname, deployedName = null, deployedAddress = null) => {
    let fac;
    try {
        fac = await ethers.getContractFactory(abiname);
    } catch(e) {
        // handle for abstract contract
        const result = await ethers.getContractAt(abiname, ethers.utils.getAddress(deployedAddress));
        return result;
    }
    // try to get by ABI & address
    if (deployedAddress) {
        // change to checksum address
        deployedAddress = ethers.utils.getAddress(deployedAddress);
        // make sure it is indeed a contract address
        const res = await ethers.provider.getCode(deployedAddress);
        console.log(`Name : ${abiname} - Code : ${res.slice(0, 64)}...`);
        if (res && res.length>2) {
            try {
                const inst = await fac.attach(deployedAddress);
                return inst;
            } catch(e) {} // ignore errors
        }
    }

    // if that fails, get by ABI & saved deployment
    let c = await deployments.getOrNull(deployedName ? deployedName : abiname);
    // contract has not been deployed => return factory
    if (!c) return fac;
    return await fac.attach(c.address);
}

let generateTestIncTokenID = (tokenAddress) => {
    return tokenAddresses.pTokens[tokenAddress] || ethers.utils.keccak256(tokenAddress).slice(2);
}

// return useful information for testing : deterministic Incognito Bridged Token ID, current test sender
let getBridgedIncTokenInfo = (dict, tokenName) => {
    if (dict.tokens && dict.tokens[tokenName]) {
        const addr = dict.tokens[tokenName].address;
        const inc = generateTestIncTokenID(addr);
        return Object.assign(dict.tokens[tokenName], {inc, sender: dict.tokenGuy, unify: tokenAddresses.unify[addr]});
    }
    return {address: tokenAddresses.ETH, inc: tokenAddresses.pETH, sender: dict.ethGuy, unify: tokenAddresses.unify[tokenAddresses.ETH]};
}

let getEmittedDepositNumber = async (dict, tokenName, amountForContract) => {
    if (dict.tokens && dict.tokens[tokenName]) {
        const addr = dict.tokens[tokenName].address;
        return await toIncDecimals(amountForContract, addr);
    }
    return amountForContract;
}

let getImplementation = async (contractAddress) => {
    const slot = '0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc'; //'eip1967.proxy.implementation';
    const result = await ethers.provider.getStorageAt(contractAddress, BN.from(slot));
    return BN.from(result);
}

let getCodeSlot = async (contractAddress, slot) => {
    const result = await ethers.provider.getStorageAt(contractAddress, BN.from(slot));
    return BN.from(result);
}

let getReqWithdrawSignMessage = async (vaultHelper, token, timestamp, amount, incAddress) => {
    const psd = {prefix: 1, token: token, timestamp: timestamp, amount: amount};
    let data = vaultHelper.interface.encodeFunctionData('_buildSignRequestWithdraw', [psd, incAddress]);
    // discard the first 4 bytes
    data = '0x'+data.slice(10);
    const hashedData = ethers.utils.keccak256(data);
    return {hashedData , data};
}

let getExecuteSignMessage = async (vaultHelper, token, timestamp, amount, receivingToken, exchangeAddress, inputData) => {
    const psd = {prefix: 0, token: token, timestamp: timestamp, amount: amount};
    let data = vaultHelper.interface.encodeFunctionData('_buildSignExecute', [psd, receivingToken, exchangeAddress, inputData]);
    // discard the first 4 bytes
    data = '0x'+data.slice(10);
    const hashedData = ethers.utils.keccak256(data);
    return {hashedData, data};
}

let getTokenAddressForKyber = (t) => {
    if (t == tokenAddresses.ETH) {
        const addresses = hre.networkCfg().deployed;
        if (addresses.kyberEtherAddress) return addresses.kyberEtherAddress;
    }
    return t;
}

let getTokenAddressForUniswap = (uniswapContract, t) =>
    t == tokenAddresses.ETH ? uniswapContract.wETH() : Promise.resolve(t);

// Vault interactions use Ethereum decimals, except for when emitting Deposit event
let prepareExternalCallByVault = async (ctx, exchangeName, srcToken, dstToken, timestamp, amount, useWorstRate = false) => {
    const exchange = ctx[exchangeName];
    if (exchangeName == 'kyber') {
        const tokensToKyber = [srcToken, dstToken].map(getTokenAddressForKyber);
        let [ expectedRate, worstRate ] = await exchange.getConversionRates(tokensToKyber[0], amount, tokensToKyber[1]);
        let chosenRate = useWorstRate ? worstRate : expectedRate;
        // console.log(`Kyber pair ${tokensToKyber} -> rate ${expectedRate.toString()} / ${worstRate.toString()}\nChoose ${chosenRate.toString()}`);
        tradeReturn = amount.mul(chosenRate).div("1000000000000000000");
        console.log(`via Kyber, trade ${ethers.utils.formatUnits(amount, 'gwei')} of ${srcToken} for ${ethers.utils.formatUnits(tradeReturn, 'gwei')} of ${dstToken}`);
        const encodedTradeCall = exchange.interface.encodeFunctionData('trade', [srcToken, amount, dstToken, chosenRate]);
        const { data, hashedData, compare } = await getExecuteSignMessage(ctx.vaultHelper, srcToken, timestamp, amount, dstToken, exchange.address, encodedTradeCall);
        // debug
        // if (srcToken != tokenAddresses.ETH) {
        //     const knr = await ethers.getContract('KyberNetworkProxy');
        //     let tokenInst = await getInstance('contracts/IERC20.sol:IERC20', null, srcToken);
        //     let tx = await confirm(tokenInst.connect(ctx.tokenGuy.signer).approve(knr.address, amount));
        //     tx = await confirm(knr.connect(ctx.tokenGuy.signer).swapTokenToToken(tokensToKyber[0], amount, tokensToKyber[1], chosenRate));
        //     console.log(tx);
        // }
        // END debug
        return {encodedTradeCall, hashedData, exchange, tradeReturn};
    } else if (exchangeName == 'uniswap') {
        const tokensToUniswap = await Promise.all([srcToken, dstToken].map(t => getTokenAddressForUniswap(exchange, t)));
        console.log(`Uniswap pair ${tokensToUniswap}`);
        let tradeAmounts = await exchange.getAmountsOut(tokensToUniswap[0], amount, tokensToUniswap[1]);
        console.log(`via Uniswap, trade ${ethers.utils.formatUnits(tradeAmounts[0], 'gwei')} of ${srcToken} for ${ethers.utils.formatUnits(tradeAmounts[1], 'gwei')} of ${dstToken}`);
        // const tx = await exchange.connect(ctx.unshieldSender).trade(srcToken, amount, dstToken, tradeAmounts[1], {value: amount});
        const encodedTradeCall = exchange.interface.encodeFunctionData('trade', [srcToken, tradeAmounts[0], dstToken, tradeAmounts[1]]);
        const { data, hashedData, compare } = await getExecuteSignMessage(ctx.vaultHelper, srcToken, timestamp, amount, dstToken, exchange.address, encodedTradeCall);
        return {encodedTradeCall, hashedData, exchange, tradeReturn: tradeAmounts[1]};
    } else throw `Exchange ${exchangeName} not supported`;
}

module.exports = {
    getPartOf,
    chooseOneFrom,
    confirm,
    toIncDecimals,
    fromIncDecimals,
    getInstance,
    getImplementation,
    getCodeSlot,
    generateTestIncTokenID,
    getBridgedIncTokenInfo,
    getEmittedDepositNumber,
    getReqWithdrawSignMessage,
    getExecuteSignMessage,
    prepareExternalCallByVault,
}