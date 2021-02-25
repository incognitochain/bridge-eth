const level = require('level');
// if space is relevant, consider using os.tmpdir() to create leveldb object (since this trie does not need to persist)
const rlp = require('rlp');
const timeMachine = require('ganache-time-traveler');
const { tokenAddresses } = require('./constants');

const BN = ethers.BigNumber;
// const eth = web3.eth;

let chooseOneFrom = (arr) => {
    const num = BN.from(ethers.utils.randomBytes(6)).toNumber();
    return arr[num % arr.length];
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

// we call tx.wait, but additionally handle for local network which does not have a defined block time
let confirm = async (txp, target = 2, network = 'development') => {
    const tx = await txp;
    console.log(`${tx.hash} => waiting for ${target} confirmations`);
    if (network=='development') {
        for (let i=0; i<=target; i++) {
            await timeMachine.advanceBlock();
        }
    }
    await tx.wait(target);
    return tx;
}

let getDecimals = (_addr) =>
    (!_addr || _addr==tokenAddresses.ETH) ? Promise.resolve(BN.from(18)) : ethers.getContractAt('contracts/external/IERC20.sol:IERC20', _addr).then(_c => _c.decimals())

let toIncDecimals = (_amount, _addr) => getDecimals(_addr)
    .then(_d => {
        let result = BN.from(_amount);
        if (_d.lte(9)) return result;
        const ten = BN.from(10);
        // console.log(`${result} * ${ten.pow(_d.subn(9))}`);
        return result.div(ten.pow(_d.sub(9)));
    })

let fromIncDecimals = (_amount, _addr) => getDecimals(_addr)
    .then(_d => {
        let result = BN.from(_amount);
        if (_d.lte(9)) return result;
        const ten = BN.from(10);
        // console.log(`${result} / ${ten.pow(_d.subn(9))}`);
        return result.mul(ten.pow(_d.sub(9)));
    })

let getInstance = async (abiname, deployedName = null) => {
    let fac = await ethers.getContractFactory(abiname);
    let c = await deployments.get(deployedName ? deployedName : abiname);
    return await fac.attach(c.address);
}

let getImplementation = async (contractAddress) => {
    const slot = '0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc'; //'eip1967.proxy.implementation';
    const result = await ethers.provider.getStorageAt(contractAddress, BN.from(slot));
    return BN.from(result);
}

// import some functions from web-js for unit tests. Tasks should not need this
let Inc;
try {
    const path = require('path');
    require = require("esm")(module/*, options*/)
    Inc = require("./lib.js");

    const webJsPath = path.join(__dirname, '../../../../incognitochain/incognito-chain-web-js');
    const {load} = require(path.join(webJsPath, 'lib/wasm/loader-node.js'));
    const wasmDefaultPath = path.join(webJsPath, 'privacy.wasm');
    Inc.Lib.prototype.init = function(fileName = wasmDefaultPath, incNodeUrl){
        this.setProvider(incNodeUrl);
        return load(fileName);
    }
} catch(e) {
    console.log('web-js not found, skipping...');
}

module.exports = {
    rlp,
    level,
    BN,
    getPartOf,
    chooseOneFrom,
    confirm,
    toIncDecimals,
    fromIncDecimals,
    getInstance,
    getImplementation,
    Inc
}