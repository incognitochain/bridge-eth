// import some functions from web-js for unit tests. Tasks should not need this
let Inc, Trie;
try {
    Inc = require('incognito');
    const path = require('path');
    require = require('esm')(module/*, options*/)
    Trie = require('./trie').Trie;
} catch(e) {
    console.error('WARNING: some optional modules are missing. Proceeding...');
    console.error(e);
}

module.exports = {
    Inc,
    Trie
}