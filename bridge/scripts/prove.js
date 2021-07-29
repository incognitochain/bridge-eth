const { ethers } = require('hardhat');
const BN = ethers.BigNumber;
const level = require('level');
const rlp = require('rlp');
const os = require('os');
const db = level(os.tmpdir());
const { Trie } = require('./external');

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
    // _gasCumul.iadd(g);
    let receipt = {
        status: _r.status ? 1 : 0,
        gasUsed: Buffer.from(ethers.utils.arrayify(g.toHexString())),
        bitvector: Buffer.from(ethers.utils.arrayify(_r.logsBloom)),
        logs: _r.logs.map(flattenLog),
    }
    // console.debug('receipt', receipt);
    return receipt;
}

let prove = (txh, encoder) => {
    let trie = new Trie(db);
    return ethers.provider.getTransaction(txh)
    .then(tx => {
        let bh = tx.blockHash;
        let myIndex = tx.transactionIndex;
        return ethers.provider.getBlock(bh)
        .then(block => {
            // grab all receipts in that block via web3
            return Promise.all(block.transactions.map(ethers.provider.getTransactionReceipt))
        })
        .then(receipts => {
            // rebuild the receipt trie
            // let gasCumulator = BN.from(0);
            let sequence = Promise.resolve();
            receipts.forEach((r, i) => {
                // cumulating gas is totally synchronous
                let encoded = rlp.encode(Object.values(getTrieReceipt(r)));
                // console.debug(ethers.utils.hexlify(encoded));
                // these trie-writing promises need to be executed sequentially
                sequence = sequence.then(_ => {
                    return trie.put(rlp.encode(i), encoded)
                })
            })
            return sequence;
        })
        .then(_ => Trie.createProof(trie, rlp.encode(myIndex)))
        // encode base64 to include in Incognito TX
        .then(_proof => ({root: trie.root, key: rlp.encode(myIndex), proof: _proof, ethBlockHash: bh, txIndex: myIndex, encodedProof: encoder ? _proof.map(encoder) : []}))
    })
}

// read gasUsed from rlp-encoded receipt
// let getGas = (s) => {
//     if (s[1]!='x') s = '0x' + s;
//     let obj = rlp.decode(s);
//     let res = web3.utils.hexToNumber(ethers.utils.hexlify(obj[1]));
//     return res;
// }

// add a '0x' prefix if it's missing. Leave non-string arguments intact
let maybeAddPrefix = (_s) => (typeof(_s)!='string' || _s.startsWith('0x') || _s.length==0) ? _s : '0x' + _s;
let deepAddPrefix = (_nestedArr) => (!_nestedArr.map) ? maybeAddPrefix(_nestedArr) : _nestedArr.map(deepAddPrefix);

let unpackSigs = (_sigs, _blk, _instRoot) => {
    // const msg = web3.utils.soliditySha3(web3.utils.soliditySha3(_blk, _instRoot));
    let v = [], r = [], s = [];
    _sigs.forEach(sig => {
        sig = maybeAddPrefix(sig);
        // optional : can recover address here to compare with 'dev' committee when debugging
        // let acc = eth.accounts.recover(msg, sig, true)
        // debug(acc);
        let arr = ethers.utils.arrayify(sig);
        v.push(arr[64]+27);
        r.push(ethers.utils.hexlify(arr.slice(0, 32)));
        s.push(ethers.utils.hexlify(arr.slice(32, 64)));
    })
    return {v,r,s};
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