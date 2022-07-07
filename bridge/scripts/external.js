// import some functions from web-js for unit tests. Tasks should not need this
let inc, Trie, encodeReceipt;
try {
    inc = require('incognito');
    const path = require('path');
    require = require('esm')(module/*, options*/)
    Inc = require('incognito');
    const t = require('./trie');
    Trie = t.Trie;
    intToBuffer = t.intToBuffer;
} catch(e) {
    console.error('WARNING: some optional modules are missing. Proceeding...');
    console.error(e);
}

module.exports = {
    inc,
    Trie,
    intToBuffer
}
