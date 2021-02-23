// const cfg = require('../hardhat.config.js');
// const { deployments, ethers } = require('hardhat');
const { confirm, BN } = require('./external');
const { MaxChiEachMint } = require('./constants');

let chiAddress = '0x0000000000004946c0e9F43F4Dee607b0eF1fA1c';
let getChi = async (_chiAddress = chiAddress) => {
    let chi = await ethers.getContractAt('ChiToken', _chiAddress);
    try {
        await chi.balanceOf('0x0000000000000000000000000000000000000000');
    } catch(e) {
        // not mainnet -> use locally deployed CHI Token contract
        // console.error(e);
        chi = await ethers.getContract('ChiToken');
    }
    return chi;
}

let mintChi = async (chiFunder, chiToMint, chiInstance = null) => {
    chiToMint = BN.from(chiToMint);
    if (!chiInstance) chiInstance = await getChi();
    let total = BN.from(0);
    console.log(`CHI contract ${chiInstance.address}`);
    while (total.lt(chiToMint)) {
        const tx = await chiInstance.connect(chiFunder).mint(BN.from(MaxChiEachMint));
        const receipt = await tx.wait();
        console.log(`Minted ${MaxChiEachMint} CHI using ${ethers.utils.formatUnits(receipt.gasUsed, 'wei')} gas`);
        total = total.add(MaxChiEachMint);
    }
}

// turn normal contract method call into one that uses some CHI to pay gas.
// this does NOT reduce gas in test networks since the CHI contract feature is bound to its mainnet address.
// to calculate the gas saving, run the tests in a mainnet fork.
let chiify = async (_funder, _chiApproved, _contract, _method, _args) => {
    const chi = await getChi();
    const startBalance = await chi.balanceOf(_funder.address);
    const approveTx = await chi.connect(_funder).approve(_contract.address, _chiApproved)
    await approveTx.wait();
    // now we can send the tx to use CHI for gas
    const actualTx = await _contract.connect(_funder)[_method](..._args)
    let receipt = await actualTx.wait();

    const balance = await chi.balanceOf(_funder.address);
    return {
        approveTx,
        actualTx,
        chiUsed: startBalance.sub(balance),
        gas: receipt.gasUsed.toNumber()
    }
}

// // predict the gas saving by CHI
// let esCHImate = (_funder, _chiApproved, _contract, _method, _args, _network) => getChi(_network)
//     .then(_chi =>
//     // first, un-approve all CHI to get `real` gas cost
//         _chi.approve(_contract.address, 0, {from: _funder})
//         .then(_tx =>
//               // match Approval event
//               expect(grabEvent(_tx, chiEvents, 'Approval', _chi.address)).to.deep.include({owner: _funder, spender: _contract.address})
//               .with.property('value').that.bignumber.equals('0')
//         )
//         // now we can estimate gas
//         .then(_ => _method.estimateGas(..._args, {from: _funder}))
//         .then(_gasNoChi => {
//             return _chi.approve(_contract.address, _chiApproved, {from: _funder})
//             // next, approve some CHI
//             .then(_tx => {
//                 expect(grabEvent(_tx, chiEvents, 'Approval', _chi.address)).to.deep.include({owner: _funder, spender: _contract.address})
//                   .with.property('value').that.bignumber.equals(_chiApproved)
//                 // send the actual TX. Note: the contract always spends CHI from msg.sender
//                 // we can also toggle off _doSend to estimate again instead of sending
//                 return _method.estimateGas(..._args, {from: _funder}).then(g => ({gasUsed: g}))
//             })
//             .then(_result => {
//                 // handle for both sending & estimating promise resolves
//                 let gasUsed = new BN(_result.gasUsed);
//                 _gasNoChi = new BN(_gasNoChi);

//                 // return the different and total for calculations
//                 return {
//                     saved : _gasNoChi.sub(gasUsed).toString(),
//                     total : _gasNoChi.toString(),
//                     receipt : _result.receipt
//                 }
//             })
//         })
//     )

module.exports = {
    chiAddress,
    getChi,
    mintChi,
    chiify
}