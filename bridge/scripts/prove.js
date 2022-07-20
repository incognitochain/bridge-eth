const { ethers } = require('hardhat');
const BN = ethers.BigNumber;
const level = require('level');
const rlp = require('rlp');
const os = require('os');
const db = level(os.tmpdir());
const { Trie, intToBuffer } = require('./external');

let flattenLog = (l) => {
    let res = [];
    res.push(Buffer.from(ethers.utils.arrayify(l.address)));
    res.push(l.topics.map(_t => Buffer.from(ethers.utils.arrayify(_t))));
    res.push(Buffer.from(ethers.utils.arrayify(l.data)));
    // console.debug(res[1]);
    return res;
}

let getTrieReceipt = (_r) => {
    let g = _r.cumulativeGasUsed;
    let receipt = {
        status: _r.status ? 1 : 0,
        gasUsed: Buffer.from(ethers.utils.arrayify(g.toHexString())),
        bitvector: Buffer.from(ethers.utils.arrayify(_r.logsBloom)),
        logs: _r.logs.map(flattenLog),
    }
    return receipt;
}

let sleep = (ms) => new Promise(resolve => setTimeout(resolve, ms));

let prove = async (txh, encoder) => {
    let trie = new Trie(db);
    let tx = await ethers.provider.getTransaction(txh);

    let bh = tx.blockHash;
    let myIndex = tx.transactionIndex;
    let block = await ethers.provider.getBlock(bh)
    // grab all receipts in that block via web3
    let receipts = await Promise.all(block.transactions.map(ethers.provider.getTransactionReceipt));
    let encodedReceipts = [];
    for (let i=0; i < receipts.length; i++) {
        const r = receipts[i];
        const currentTxh = block.transactions[i];
        const encoded = rlp.encode(Object.values(getTrieReceipt(r)));
        const tx = await ethers.provider.getTransaction(currentTxh);
        if (i == receipts.length && tx.to == '0x0000000000000000000000000000000000000000' && tx.from == '0x0000000000000000000000000000000000000000') {
            continue;
        }
        if (tx.type != 0) {
            encodedReceipts.push(Buffer.concat([intToBuffer(tx.type), encoded]));
        } else encodedReceipts.push(encoded);
        await sleep(500); // space out requests to avoid getting rate-limited
    };
  
    const bump = encodedReceipts.length < 128 ? encodedReceipts.length : 128;
    for (let i=1; i<bump; i++) {
        await trie.put(rlp.encode(i), encodedReceipts[i]);
    }
    await trie.put(rlp.encode(0), encodedReceipts[0]);
    for (let i=128; i<encodedReceipts.length; i++) {
        await trie.put(rlp.encode(i), encodedReceipts[i]);
    }
    let proof = await Trie.createProof(trie, rlp.encode(myIndex));

    // encode base64 to include in Incognito TX
    return { root: trie.root, key: rlp.encode(myIndex), proof, ethBlockHash: bh, txIndex: myIndex, encodedProof: encoder ? proof.map(encoder) : [] }
}

// read gasUsed from rlp-encoded receipt
// let getGas = (s) => {
//     if (s[1]!='x') s = '0x' + s;
//     let obj = rlp.decode(s);
//     let res = web3.utils.hexToNumber(ethers.utils.hexlify(obj[1]));
//     return res;
// }

// add a '0x' prefix if it's missing. Leave non-string arguments intact
let maybeAddPrefix = (_s) => (typeof(_s) != 'string' || _s.startsWith('0x') || _s.length == 0) ? _s : '0x' + _s;
let deepAddPrefix = (_nestedArr) => (!_nestedArr.map) ? maybeAddPrefix(_nestedArr) : _nestedArr.map(deepAddPrefix);

let unpackSigs = (_sigs, _blk, _instRoot) => {
    // const msg = web3.utils.soliditySha3(web3.utils.soliditySha3(_blk, _instRoot));
    let v = [],
        r = [],
        s = [];
    _sigs.forEach(sig => {
        sig = maybeAddPrefix(sig);
        // optional : can recover address here to compare with 'dev' committee when debugging
        // let acc = eth.accounts.recover(msg, sig, true)
        // debug(acc);
        let arr = ethers.utils.arrayify(sig);
        v.push(arr[64] + 27);
        r.push(ethers.utils.hexlify(arr.slice(0, 32)));
        s.push(ethers.utils.hexlify(arr.slice(32, 64)));
    })
    return { v, r, s };
}

let formatBurnProof = (obj) => {
    let sigs = unpackSigs(...deepAddPrefix([obj.BeaconSigs, obj.BeaconBlkData, obj.BeaconInstRoot]));
    //inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs
    return deepAddPrefix([obj.Instruction, obj.BeaconHeight, obj.BeaconInstPath, obj.BeaconInstPathIsLeft, obj.BeaconInstRoot, obj.BeaconBlkData, obj.BeaconSigIdxs, sigs.v, sigs.r, sigs.s]);
}

module.exports = {
    proveEth: prove,
    toPrefixed: deepAddPrefix,
    formatBurnProof
}