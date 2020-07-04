// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package incognito_proxy

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IncognitoProxyCommittee is an auto generated low-level Go binding around an user-defined struct.
type IncognitoProxyCommittee struct {
	Pubkeys    []common.Address
	StartBlock *big.Int
}

// IncognitoProxyInstructionProof is an auto generated low-level Go binding around an user-defined struct.
type IncognitoProxyInstructionProof struct {
	Inst    []byte
	Path    [][32]byte
	IsLeft  []bool
	Root    [32]byte
	BlkData [32]byte
	SigIdx  []*big.Int
	SigV    []uint8
	SigR    [][32]byte
	SigS    [][32]byte
}

// AdminPausableABI is the input ABI used to generate the binding from.
const AdminPausableABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ndays\",\"type\":\"uint256\"}],\"name\":\"Extend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expire\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"extend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_successor\",\"type\":\"address\"}],\"name\":\"retire\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"successor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AdminPausableFuncSigs maps the 4-byte function signature to its string representation.
var AdminPausableFuncSigs = map[string]string{
	"f851a440": "admin()",
	"4e71d92d": "claim()",
	"79599f96": "expire()",
	"9714378c": "extend(uint256)",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"9e6371ba": "retire(address)",
	"6ff968c3": "successor()",
	"3f4ba83a": "unpause()",
}

// AdminPausableBin is the compiled bytecode used for deploying new contracts.
var AdminPausableBin = "0x608060405234801561001057600080fd5b506040516106fc3803806106fc8339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b03199092169190911790556001805460ff60a01b191690556301e1338042016002556106808061007c6000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c806379599f961161006657806379599f96146100ea5780638456cb59146101045780639714378c1461010c5780639e6371ba14610129578063f851a4401461014f57610093565b80633f4ba83a146100985780634e71d92d146100a25780635c975abb146100aa5780636ff968c3146100c6575b600080fd5b6100a0610157565b005b6100a0610239565b6100b2610320565b604080519115158252519081900360200190f35b6100ce610330565b604080516001600160a01b039092168252519081900360200190f35b6100f261033f565b60408051918252519081900360200190f35b6100a0610345565b6100a06004803603602081101561012257600080fd5b503561046a565b6100a06004803603602081101561013f57600080fd5b50356001600160a01b031661058e565b6100ce61063b565b6000546001600160a01b031633146101a2576040805162461bcd60e51b81526020600482015260096024820152683737ba1030b236b4b760b91b604482015290519081900360640190fd5b600154600160a01b900460ff166101f7576040805162461bcd60e51b81526020600482015260146024820152736e6f7420706175736564207269676874206e6f7760601b604482015290519081900360640190fd5b6001805460ff60a01b191690556040805133815290517f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa9181900360200190a1565b6002544210610279576040805162461bcd60e51b8152602060048201526007602482015266195e1c1a5c995960ca1b604482015290519081900360640190fd5b6001546001600160a01b031633146102c7576040805162461bcd60e51b815260206004820152600c60248201526b1d5b985d5d1a1bdc9a5e995960a21b604482015290519081900360640190fd5b600154600080546001600160a01b0319166001600160a01b0392831617908190556040805191909216815290517f0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc9181900360200190a1565b600154600160a01b900460ff1681565b6001546001600160a01b031681565b60025481565b6000546001600160a01b03163314610390576040805162461bcd60e51b81526020600482015260096024820152683737ba1030b236b4b760b91b604482015290519081900360640190fd5b600154600160a01b900460ff16156103e2576040805162461bcd60e51b815260206004820152601060248201526f706175736564207269676874206e6f7760801b604482015290519081900360640190fd5b6002544210610422576040805162461bcd60e51b8152602060048201526007602482015266195e1c1a5c995960ca1b604482015290519081900360640190fd5b6001805460ff60a01b1916600160a01b1790556040805133815290517f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2589181900360200190a1565b6000546001600160a01b031633146104b5576040805162461bcd60e51b81526020600482015260096024820152683737ba1030b236b4b760b91b604482015290519081900360640190fd5b60025442106104f5576040805162461bcd60e51b8152602060048201526007602482015266195e1c1a5c995960ca1b604482015290519081900360640190fd5b61016e811061054b576040805162461bcd60e51b815260206004820152601a60248201527f63616e6e6f7420657874656e6420666f7220746f6f206c6f6e67000000000000604482015290519081900360640190fd5b600280546201518083020190556040805182815290517f02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e89181900360200190a150565b6000546001600160a01b031633146105d9576040805162461bcd60e51b81526020600482015260096024820152683737ba1030b236b4b760b91b604482015290519081900360640190fd5b6002544210610619576040805162461bcd60e51b8152602060048201526007602482015266195e1c1a5c995960ca1b604482015290519081900360640190fd5b600180546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b03168156fea2646970667358221220f07a79af805ec091ff5b9b9d4b8853e8c61f0e8db1f210d302f28d54ee5b38a364736f6c63430006060033"

// DeployAdminPausable deploys a new Ethereum contract, binding an instance of AdminPausable to it.
func DeployAdminPausable(auth *bind.TransactOpts, backend bind.ContractBackend, _admin common.Address) (common.Address, *types.Transaction, *AdminPausable, error) {
	parsed, err := abi.JSON(strings.NewReader(AdminPausableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AdminPausableBin), backend, _admin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AdminPausable{AdminPausableCaller: AdminPausableCaller{contract: contract}, AdminPausableTransactor: AdminPausableTransactor{contract: contract}, AdminPausableFilterer: AdminPausableFilterer{contract: contract}}, nil
}

// AdminPausable is an auto generated Go binding around an Ethereum contract.
type AdminPausable struct {
	AdminPausableCaller     // Read-only binding to the contract
	AdminPausableTransactor // Write-only binding to the contract
	AdminPausableFilterer   // Log filterer for contract events
}

// AdminPausableCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdminPausableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminPausableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdminPausableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminPausableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AdminPausableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminPausableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdminPausableSession struct {
	Contract     *AdminPausable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdminPausableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdminPausableCallerSession struct {
	Contract *AdminPausableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AdminPausableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdminPausableTransactorSession struct {
	Contract     *AdminPausableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AdminPausableRaw is an auto generated low-level Go binding around an Ethereum contract.
type AdminPausableRaw struct {
	Contract *AdminPausable // Generic contract binding to access the raw methods on
}

// AdminPausableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AdminPausableCallerRaw struct {
	Contract *AdminPausableCaller // Generic read-only contract binding to access the raw methods on
}

// AdminPausableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AdminPausableTransactorRaw struct {
	Contract *AdminPausableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAdminPausable creates a new instance of AdminPausable, bound to a specific deployed contract.
func NewAdminPausable(address common.Address, backend bind.ContractBackend) (*AdminPausable, error) {
	contract, err := bindAdminPausable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AdminPausable{AdminPausableCaller: AdminPausableCaller{contract: contract}, AdminPausableTransactor: AdminPausableTransactor{contract: contract}, AdminPausableFilterer: AdminPausableFilterer{contract: contract}}, nil
}

// NewAdminPausableCaller creates a new read-only instance of AdminPausable, bound to a specific deployed contract.
func NewAdminPausableCaller(address common.Address, caller bind.ContractCaller) (*AdminPausableCaller, error) {
	contract, err := bindAdminPausable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdminPausableCaller{contract: contract}, nil
}

// NewAdminPausableTransactor creates a new write-only instance of AdminPausable, bound to a specific deployed contract.
func NewAdminPausableTransactor(address common.Address, transactor bind.ContractTransactor) (*AdminPausableTransactor, error) {
	contract, err := bindAdminPausable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdminPausableTransactor{contract: contract}, nil
}

// NewAdminPausableFilterer creates a new log filterer instance of AdminPausable, bound to a specific deployed contract.
func NewAdminPausableFilterer(address common.Address, filterer bind.ContractFilterer) (*AdminPausableFilterer, error) {
	contract, err := bindAdminPausable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdminPausableFilterer{contract: contract}, nil
}

// bindAdminPausable binds a generic wrapper to an already deployed contract.
func bindAdminPausable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AdminPausableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdminPausable *AdminPausableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AdminPausable.Contract.AdminPausableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdminPausable *AdminPausableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminPausable.Contract.AdminPausableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdminPausable *AdminPausableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdminPausable.Contract.AdminPausableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdminPausable *AdminPausableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AdminPausable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdminPausable *AdminPausableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminPausable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdminPausable *AdminPausableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdminPausable.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_AdminPausable *AdminPausableCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AdminPausable.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_AdminPausable *AdminPausableSession) Admin() (common.Address, error) {
	return _AdminPausable.Contract.Admin(&_AdminPausable.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_AdminPausable *AdminPausableCallerSession) Admin() (common.Address, error) {
	return _AdminPausable.Contract.Admin(&_AdminPausable.CallOpts)
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() view returns(uint256)
func (_AdminPausable *AdminPausableCaller) Expire(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AdminPausable.contract.Call(opts, out, "expire")
	return *ret0, err
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() view returns(uint256)
func (_AdminPausable *AdminPausableSession) Expire() (*big.Int, error) {
	return _AdminPausable.Contract.Expire(&_AdminPausable.CallOpts)
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() view returns(uint256)
func (_AdminPausable *AdminPausableCallerSession) Expire() (*big.Int, error) {
	return _AdminPausable.Contract.Expire(&_AdminPausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AdminPausable *AdminPausableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AdminPausable.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AdminPausable *AdminPausableSession) Paused() (bool, error) {
	return _AdminPausable.Contract.Paused(&_AdminPausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AdminPausable *AdminPausableCallerSession) Paused() (bool, error) {
	return _AdminPausable.Contract.Paused(&_AdminPausable.CallOpts)
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() view returns(address)
func (_AdminPausable *AdminPausableCaller) Successor(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AdminPausable.contract.Call(opts, out, "successor")
	return *ret0, err
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() view returns(address)
func (_AdminPausable *AdminPausableSession) Successor() (common.Address, error) {
	return _AdminPausable.Contract.Successor(&_AdminPausable.CallOpts)
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() view returns(address)
func (_AdminPausable *AdminPausableCallerSession) Successor() (common.Address, error) {
	return _AdminPausable.Contract.Successor(&_AdminPausable.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_AdminPausable *AdminPausableTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminPausable.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_AdminPausable *AdminPausableSession) Claim() (*types.Transaction, error) {
	return _AdminPausable.Contract.Claim(&_AdminPausable.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_AdminPausable *AdminPausableTransactorSession) Claim() (*types.Transaction, error) {
	return _AdminPausable.Contract.Claim(&_AdminPausable.TransactOpts)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_AdminPausable *AdminPausableTransactor) Extend(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _AdminPausable.contract.Transact(opts, "extend", n)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_AdminPausable *AdminPausableSession) Extend(n *big.Int) (*types.Transaction, error) {
	return _AdminPausable.Contract.Extend(&_AdminPausable.TransactOpts, n)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_AdminPausable *AdminPausableTransactorSession) Extend(n *big.Int) (*types.Transaction, error) {
	return _AdminPausable.Contract.Extend(&_AdminPausable.TransactOpts, n)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AdminPausable *AdminPausableTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminPausable.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AdminPausable *AdminPausableSession) Pause() (*types.Transaction, error) {
	return _AdminPausable.Contract.Pause(&_AdminPausable.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AdminPausable *AdminPausableTransactorSession) Pause() (*types.Transaction, error) {
	return _AdminPausable.Contract.Pause(&_AdminPausable.TransactOpts)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_AdminPausable *AdminPausableTransactor) Retire(opts *bind.TransactOpts, _successor common.Address) (*types.Transaction, error) {
	return _AdminPausable.contract.Transact(opts, "retire", _successor)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_AdminPausable *AdminPausableSession) Retire(_successor common.Address) (*types.Transaction, error) {
	return _AdminPausable.Contract.Retire(&_AdminPausable.TransactOpts, _successor)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_AdminPausable *AdminPausableTransactorSession) Retire(_successor common.Address) (*types.Transaction, error) {
	return _AdminPausable.Contract.Retire(&_AdminPausable.TransactOpts, _successor)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AdminPausable *AdminPausableTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminPausable.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AdminPausable *AdminPausableSession) Unpause() (*types.Transaction, error) {
	return _AdminPausable.Contract.Unpause(&_AdminPausable.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AdminPausable *AdminPausableTransactorSession) Unpause() (*types.Transaction, error) {
	return _AdminPausable.Contract.Unpause(&_AdminPausable.TransactOpts)
}

// AdminPausableClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the AdminPausable contract.
type AdminPausableClaimIterator struct {
	Event *AdminPausableClaim // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AdminPausableClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminPausableClaim)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AdminPausableClaim)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AdminPausableClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminPausableClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminPausableClaim represents a Claim event raised by the AdminPausable contract.
type AdminPausableClaim struct {
	Claimer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_AdminPausable *AdminPausableFilterer) FilterClaim(opts *bind.FilterOpts) (*AdminPausableClaimIterator, error) {

	logs, sub, err := _AdminPausable.contract.FilterLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return &AdminPausableClaimIterator{contract: _AdminPausable.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_AdminPausable *AdminPausableFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *AdminPausableClaim) (event.Subscription, error) {

	logs, sub, err := _AdminPausable.contract.WatchLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminPausableClaim)
				if err := _AdminPausable.contract.UnpackLog(event, "Claim", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaim is a log parse operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_AdminPausable *AdminPausableFilterer) ParseClaim(log types.Log) (*AdminPausableClaim, error) {
	event := new(AdminPausableClaim)
	if err := _AdminPausable.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AdminPausableExtendIterator is returned from FilterExtend and is used to iterate over the raw logs and unpacked data for Extend events raised by the AdminPausable contract.
type AdminPausableExtendIterator struct {
	Event *AdminPausableExtend // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AdminPausableExtendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminPausableExtend)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AdminPausableExtend)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AdminPausableExtendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminPausableExtendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminPausableExtend represents a Extend event raised by the AdminPausable contract.
type AdminPausableExtend struct {
	Ndays *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterExtend is a free log retrieval operation binding the contract event 0x02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8.
//
// Solidity: event Extend(uint256 ndays)
func (_AdminPausable *AdminPausableFilterer) FilterExtend(opts *bind.FilterOpts) (*AdminPausableExtendIterator, error) {

	logs, sub, err := _AdminPausable.contract.FilterLogs(opts, "Extend")
	if err != nil {
		return nil, err
	}
	return &AdminPausableExtendIterator{contract: _AdminPausable.contract, event: "Extend", logs: logs, sub: sub}, nil
}

// WatchExtend is a free log subscription operation binding the contract event 0x02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8.
//
// Solidity: event Extend(uint256 ndays)
func (_AdminPausable *AdminPausableFilterer) WatchExtend(opts *bind.WatchOpts, sink chan<- *AdminPausableExtend) (event.Subscription, error) {

	logs, sub, err := _AdminPausable.contract.WatchLogs(opts, "Extend")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminPausableExtend)
				if err := _AdminPausable.contract.UnpackLog(event, "Extend", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExtend is a log parse operation binding the contract event 0x02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8.
//
// Solidity: event Extend(uint256 ndays)
func (_AdminPausable *AdminPausableFilterer) ParseExtend(log types.Log) (*AdminPausableExtend, error) {
	event := new(AdminPausableExtend)
	if err := _AdminPausable.contract.UnpackLog(event, "Extend", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AdminPausablePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the AdminPausable contract.
type AdminPausablePausedIterator struct {
	Event *AdminPausablePaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AdminPausablePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminPausablePaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AdminPausablePaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AdminPausablePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminPausablePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminPausablePaused represents a Paused event raised by the AdminPausable contract.
type AdminPausablePaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address pauser)
func (_AdminPausable *AdminPausableFilterer) FilterPaused(opts *bind.FilterOpts) (*AdminPausablePausedIterator, error) {

	logs, sub, err := _AdminPausable.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AdminPausablePausedIterator{contract: _AdminPausable.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address pauser)
func (_AdminPausable *AdminPausableFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AdminPausablePaused) (event.Subscription, error) {

	logs, sub, err := _AdminPausable.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminPausablePaused)
				if err := _AdminPausable.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address pauser)
func (_AdminPausable *AdminPausableFilterer) ParsePaused(log types.Log) (*AdminPausablePaused, error) {
	event := new(AdminPausablePaused)
	if err := _AdminPausable.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AdminPausableUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the AdminPausable contract.
type AdminPausableUnpausedIterator struct {
	Event *AdminPausableUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AdminPausableUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminPausableUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AdminPausableUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AdminPausableUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminPausableUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminPausableUnpaused represents a Unpaused event raised by the AdminPausable contract.
type AdminPausableUnpaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address pauser)
func (_AdminPausable *AdminPausableFilterer) FilterUnpaused(opts *bind.FilterOpts) (*AdminPausableUnpausedIterator, error) {

	logs, sub, err := _AdminPausable.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AdminPausableUnpausedIterator{contract: _AdminPausable.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address pauser)
func (_AdminPausable *AdminPausableFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AdminPausableUnpaused) (event.Subscription, error) {

	logs, sub, err := _AdminPausable.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminPausableUnpaused)
				if err := _AdminPausable.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address pauser)
func (_AdminPausable *AdminPausableFilterer) ParseUnpaused(log types.Log) (*AdminPausableUnpaused, error) {
	event := new(AdminPausableUnpaused)
	if err := _AdminPausable.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncognitoProxyABI is the input ABI used to generate the binding from.
const IncognitoProxyABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"beaconCommittee\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"bridgeCommittee\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startHeight\",\"type\":\"uint256\"}],\"name\":\"BeaconCommitteeSwapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startHeight\",\"type\":\"uint256\"}],\"name\":\"BridgeCommitteeSwapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isBeacon\",\"type\":\"bool\"}],\"name\":\"ChainFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ndays\",\"type\":\"uint256\"}],\"name\":\"Extend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startHeight\",\"type\":\"uint256\"}],\"name\":\"SubmittedBridgeCandidate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"beaconCandidates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"beaconCommittees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beaconFinality\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bridgeCandidates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bridgeCommittees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeFinality\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expire\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"extend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"numVals\",\"type\":\"uint256\"}],\"name\":\"extractCommitteeFromInstruction\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"}],\"name\":\"extractDataFromBlockMerkleInstruction\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"}],\"name\":\"extractMetaFromInstruction\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blkHeight\",\"type\":\"uint256\"}],\"name\":\"findBeaconCommitteeFromHeight\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blkHeight\",\"type\":\"uint256\"}],\"name\":\"findBridgeCommitteeFromHeight\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getBeaconCommittee\",\"outputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"pubkeys\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIncognitoProxy.Committee\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getBridgeCommittee\",\"outputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"pubkeys\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIncognitoProxy.Committee\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isBeacon\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"instHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blkHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"instPath\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"instPathIsLeft\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"instRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"sigIdx\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"sigV\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigR\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigS\",\"type\":\"bytes32[]\"}],\"name\":\"instructionApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"instHash\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"instPath\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"instPathIsLeft\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"instRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint8[]\",\"name\":\"sigV\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigR\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigS\",\"type\":\"bytes32[]\"}],\"name\":\"instructionApprovedByCommittee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"left\",\"type\":\"bool[]\"}],\"name\":\"instructionInMerkleTree\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_successor\",\"type\":\"address\"}],\"name\":\"retire\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"isLeft\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"sigIdx\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"sigV\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigR\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigS\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIncognitoProxy.InstructionProof\",\"name\":\"instProof\",\"type\":\"tuple\"}],\"name\":\"submitBeaconCandidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"isLeft\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"sigIdx\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"sigV\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigR\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigS\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIncognitoProxy.InstructionProof[2]\",\"name\":\"instProofs\",\"type\":\"tuple[2]\"}],\"name\":\"submitBridgeCandidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"isLeft\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"sigIdx\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"sigV\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigR\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigS\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIncognitoProxy.InstructionProof[]\",\"name\":\"instProofs\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"swapID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isCandidate\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isBeacon\",\"type\":\"bool\"}],\"name\":\"submitFinalityProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"successor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"committee\",\"type\":\"address[]\"},{\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"verifySig\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// IncognitoProxyFuncSigs maps the 4-byte function signature to its string representation.
var IncognitoProxyFuncSigs = map[string]string{
	"f851a440": "admin()",
	"62532738": "beaconCandidates(uint256)",
	"f203a5ed": "beaconCommittees(uint256)",
	"5d7f86d9": "beaconFinality()",
	"e83d31ec": "bridgeCandidates(uint256)",
	"9b30b637": "bridgeCommittees(uint256)",
	"2c1c9cac": "bridgeFinality()",
	"4e71d92d": "claim()",
	"79599f96": "expire()",
	"9714378c": "extend(uint256)",
	"8eb60066": "extractCommitteeFromInstruction(bytes,uint256)",
	"f179dc78": "extractDataFromBlockMerkleInstruction(bytes)",
	"90500bae": "extractMetaFromInstruction(bytes)",
	"b600ffdb": "findBeaconCommitteeFromHeight(uint256)",
	"f5205fde": "findBridgeCommitteeFromHeight(uint256)",
	"faea3167": "getBeaconCommittee(uint256)",
	"8ceb69c3": "getBridgeCommittee(uint256)",
	"f65d2116": "instructionApproved(bool,bytes32,uint256,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])",
	"64a72862": "instructionApprovedByCommittee(bytes32,address[],bytes32[],bool[],bytes32,bytes32,uint8[],bytes32[],bytes32[])",
	"47c4b328": "instructionInMerkleTree(bytes32,bytes32,bytes32[],bool[])",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"9e6371ba": "retire(address)",
	"52d45239": "submitBeaconCandidate(bytes,(bytes,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[]))",
	"7e59f5bf": "submitBridgeCandidate(bytes,(bytes,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])[2])",
	"cc554f00": "submitFinalityProof((bytes,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])[],uint256,bool,bool)",
	"6ff968c3": "successor()",
	"3f4ba83a": "unpause()",
	"3aacfdad": "verifySig(address[],bytes32,uint8[],bytes32[],bytes32[])",
}

// IncognitoProxyBin is the compiled bytecode used for deploying new contracts.
var IncognitoProxyBin = "0x60806040523480156200001157600080fd5b5060405162002ab838038062002ab883398101604081905262000034916200029e565b600080546001600160a01b0385166001600160a01b03199091161781556001805460ff60a01b191681556301e1338042016002908155604080518082019091528581526020808201859052600380549485018155909452805180519194939092027fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0192620000c89284929101906200014a565b50602091820151600191820155604080518082019091528381526000818401819052600480549384018155905280518051919360029093027f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b019262000134928492909101906200014a565b5060208201518160010155505050505062000319565b828054828255906000526020600020908101928215620001a2579160200282015b82811115620001a257825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906200016b565b50620001b0929150620001b4565b5090565b620001db91905b80821115620001b05780546001600160a01b0319168155600101620001bb565b90565b80516001600160a01b0381168114620001f657600080fd5b92915050565b600082601f8301126200020d578081fd5b81516001600160401b038082111562000224578283fd5b60208083026040518282820101818110858211171562000242578687fd5b6040528481529450818501925085820181870183018810156200026457600080fd5b600091505b8482101562000293576200027e8882620001de565b84529282019260019190910190820162000269565b505050505092915050565b600080600060608486031215620002b3578283fd5b620002bf8585620001de565b60208501519093506001600160401b0380821115620002dc578384fd5b620002ea87838801620001fc565b9350604086015191508082111562000300578283fd5b506200030f86828701620001fc565b9150509250925092565b61278f80620003296000396000f3fe608060405234801561001057600080fd5b50600436106101cf5760003560e01c80638ceb69c311610104578063cc554f00116100a2578063f5205fde11610071578063f5205fde146103df578063f65d2116146103f2578063f851a44014610405578063faea31671461040d576101cf565b8063cc554f0014610384578063e83d31ec14610397578063f179dc78146103aa578063f203a5ed146103cc576101cf565b80639714378c116100de5780639714378c1461032a5780639b30b6371461033d5780639e6371ba14610350578063b600ffdb14610363576101cf565b80638ceb69c3146102c65780638eb60066146102e657806390500bae14610306576101cf565b80635d7f86d9116101715780636ff968c31161014b5780636ff968c31461028e57806379599f96146102a35780637e59f5bf146102ab5780638456cb59146102be576101cf565b80635d7f86d914610253578063625327381461025b57806364a728621461027b576101cf565b806347c4b328116101ad57806347c4b3281461021d5780634e71d92d1461023057806352d45239146102385780635c975abb1461024b576101cf565b80632c1c9cac146101d45780633aacfdad146101f35780633f4ba83a14610213575b600080fd5b6101dc610420565b6040516101ea9291906123c7565b60405180910390f35b610206610201366004611dc4565b610429565b6040516101ea919061241e565b61021b610527565b005b61020661022b36600461216f565b6105ca565b61021b6106c7565b61021b6102463660046122c1565b610763565b61020661090f565b6101dc61091f565b61026e610269366004612363565b610928565b6040516101ea91906123be565b610206610289366004612063565b61093d565b6102966109ef565b6040516101ea91906123d5565b61026e6109fe565b61021b6102b936600461221c565b610a04565b61021b610c65565b6102d96102d4366004612363565b610d1d565b6040516101ea9190612681565b6102f96102f4366004612321565b610dbe565b6040516101ea91906123e9565b6103196103143660046121e2565b610e69565b6040516101ea9594939291906126cd565b61021b610338366004612363565b610edb565b61026e61034b366004612363565b610f8f565b61021b61035e366004611da2565b610fb5565b610376610371366004612363565b611022565b6040516101ea9291906123fc565b61021b610392366004611e75565b611110565b61026e6103a5366004612363565b611736565b6103bd6103b83660046121e2565b61174b565b6040516101ea939291906126b2565b61026e6103da366004612363565b61178e565b6103766103ed366004612363565b61179b565b610206610400366004611f2f565b611810565b610296611956565b6102d961041b366004612363565b611965565b600954600a5482565b6000825184511461043957600080fd5b815184511461044757600080fd5b60005b84518110156105185786818151811061045f57fe5b60200260200101516001600160a01b031660018787848151811061047f57fe5b602002602001015187858151811061049357fe5b60200260200101518786815181106104a757fe5b6020026020010151604051600081526020016040526040516104cc9493929190612429565b6020604051602081039080840390855afa1580156104ee573d6000803e3d6000fd5b505050602060405103516001600160a01b03161461051057600091505061051e565b60010161044a565b50600190505b95945050505050565b6000546001600160a01b0316331461055a5760405162461bcd60e51b815260040161055190612631565b60405180910390fd5b600154600160a01b900460ff166105835760405162461bcd60e51b815260040161055190612447565b6001805460ff60a01b191690556040517f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa906105c09033906123d5565b60405180910390a1565b600082518251146105da57600080fd5b8460005b84518110156106bb578381815181106105f357fe5b6020026020010151156106435784818151811061060c57fe5b6020026020010151826040516020016106269291906123c7565b6040516020818303038152906040528051906020012091506106b3565b84818151811061064f57fe5b60200260200101516000801b14156106745781826040516020016106269291906123c7565b8185828151811061068157fe5b602002602001015160405160200161069a9291906123c7565b6040516020818303038152906040528051906020012091505b6001016105de565b50909314949350505050565b60025442106106e85760405162461bcd60e51b8152600401610551906124ff565b6001546001600160a01b031633146107125760405162461bcd60e51b81526004016105519061257d565b600154600080546001600160a01b0319166001600160a01b0392831617908190556040517f0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc926105c09216906123d5565b600154600160a01b900460ff161561078d5760405162461bcd60e51b815260040161055190612607565b81516020830120600380546107f09160019184919060001981019081106107b057fe5b90600052602060002090600202016001015485602001518660400151876060015188608001518960a001518a60c001518b60e001518c6101000151611810565b6107f957600080fd5b600080600080600061080a88610e69565b945094509450945094508460ff16604614801561082a57508360ff166001145b61083357600080fd5b60038054600019810190811061084557fe5b90600052602060002090600202016001015483116108755760405162461bcd60e51b8152600401610551906125a3565b60606108818984610dbe565b60408051808201825282815260208082018890526000868152600682529290922081518051949550919390926108bb92849291019061197a565b50602091909101516001909101556003546040517fc918d4d247e9af7d4a8ef8ac6f0078262006c87209d062a6d846f5b547dcff66916108fc9187906123c7565b60405180910390a1505050505050505050565b600154600160a01b900460ff1681565b60075460085482565b60066020526000908152604090206001015481565b60008085876040516020016109539291906123c7565b6040516020818303038152906040528051906020012060405160200161097991906123be565b60405160208183030381529060405280519060200120905060038a516002028161099f57fe5b048551116109b15760009150506109e2565b6109be8a82878787610429565b6109c757600080fd5b6109d38b888b8b6105ca565b6109dc57600080fd5b60019150505b9998505050505050505050565b6001546001600160a01b031681565b60025481565b600154600160a01b900460ff1615610a2e5760405162461bcd60e51b815260040161055190612607565b8151602083012060038054610ac5916001918491906000198101908110610a5157fe5b90600052602060002090600202016001015485600060028110610a7057fe5b60200201516020015186600060028110610a8657fe5b60200201516040015187516060810151608082015160a083015160c084015160e090940151929391929091908c60005b60200201516101000151611810565b610ace57600080fd5b610b59600082600460016004805490500381548110610ae957fe5b90600052602060002090600202016001015485600160028110610b0857fe5b60200201516020015186600160028110610b1e57fe5b6020908102919091015160400151908801516060810151608082015160a083015160c084015160e090940151929391929091908c6001610ab6565b610b6257600080fd5b6000806000806000610b7388610e69565b945094509450945094508460ff166047148015610b9357508360ff166001145b610b9c57600080fd5b600480546000198101908110610bae57fe5b9060005260206000209060020201600101548311610bde5760405162461bcd60e51b8152600401610551906125a3565b6060610bea8984610dbe565b6040805180820182528281526020808201889052600086815260058252929092208151805194955091939092610c2492849291019061197a565b50602091909101516001909101556004546040517fc918d4d247e9af7d4a8ef8ac6f0078262006c87209d062a6d846f5b547dcff66916108fc9187906123c7565b6000546001600160a01b03163314610c8f5760405162461bcd60e51b815260040161055190612631565b600154600160a01b900460ff1615610cb95760405162461bcd60e51b815260040161055190612607565b6002544210610cda5760405162461bcd60e51b8152600401610551906124ff565b6001805460ff60a01b1916600160a01b1790556040517f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258906105c09033906123d5565b610d256119df565b60048281548110610d3257fe5b906000526020600020906002020160405180604001604052908160008201805480602002602001604051908101604052809291908181526020018280548015610da457602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610d86575b505050505081526020016001820154815250509050919050565b606081602002606201835114610dd357600080fd5b6060826001600160401b0381118015610deb57600080fd5b50604051908082528060200260200182016040528015610e15578160200160208202803683370190505b5090506000805b84811015610e5d5760208102608287010151915081838281518110610e3d57fe5b6001600160a01b0390921660209283029190910190910152600101610e1c565b50909150505b92915050565b6000806000806000604286511015610e8057600080fd5b600086600081518110610e8f57fe5b602001015160f81c60f81b60f81c9050600087600181518110610eae57fe5b0160200151602289015160428a01516062909a0151939a60f89290921c9990985096509194509092505050565b6000546001600160a01b03163314610f055760405162461bcd60e51b815260040161055190612631565b6002544210610f265760405162461bcd60e51b8152600401610551906124ff565b61016e8110610f475760405162461bcd60e51b815260040161055190612546565b600280546201518083020190556040517f02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e890610f849083906123be565b60405180910390a150565b60048181548110610f9c57fe5b6000918252602090912060016002909202010154905081565b6000546001600160a01b03163314610fdf5760405162461bcd60e51b815260040161055190612631565b60025442106110005760405162461bcd60e51b8152600401610551906124ff565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b60035460609060009081908061103757600080fd5b600019015b80821461108a5760006002600184840101049050856003828154811061105e57fe5b9060005260206000209060020201600101541161107d57809250611084565b6001810391505b5061103c565b6003828154811061109757fe5b906000526020600020906002020160000182818054806020026020016040519081016040528092919081815260200182805480156110fe57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116110e0575b50505050509150935093505050915091565b600154600160a01b900460ff161561113a5760405162461bcd60e51b815260040161055190612607565b835160021461115b5760405162461bcd60e51b815260040161055190612475565b6060821561123e5781156111d357600084815260066020908152604091829020805483518184028101840190945280845290918301828280156111c757602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116111a9575b50505050509050611239565b6000848152600560209081526040918290208054835181840281018401909452808452909183018282801561123157602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611213575b505050505090505b611336565b81156112be5760038054600019810190811061125657fe5b60009182526020918290206002909102018054604080518285028101850190915281815292830182828015611231576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116112135750505050509050611336565b6004805460001981019081106112d057fe5b6000918252602091829020600290910201805460408051828502810185019091528181529283018282801561132e57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611310575b505050505090505b606081516001600160401b038111801561134f57600080fd5b50604051908082528060200260200182016040528015611379578160200160208202803683370190505b50905060005b60028110156115fe5786818151811061139457fe5b602002602001015160a00151518782815181106113ad57fe5b602002602001015160c0015151146113c457600080fd5b60005b8782815181106113d357fe5b602002602001015160a001515181101561151b5783518883815181106113f557fe5b602002602001015160a00151828151811061140c57fe5b6020026020010151106114315760405162461bcd60e51b8152600401610551906125da565b801580611495575087828151811061144557fe5b602002602001015160a00151600182038151811061145f57fe5b602002602001015188838151811061147357fe5b602002602001015160a00151828151811061148a57fe5b602002602001015111155b6114b15760405162461bcd60e51b8152600401610551906124bd565b838883815181106114be57fe5b602002602001015160a0015182815181106114d557fe5b6020026020010151815181106114e757fe5b60200260200101518382815181106114fb57fe5b6001600160a01b03909216602092830291909101909101526001016113c7565b506115ed87828151811061152b57fe5b602002602001015160000151805190602001208389848151811061154b57fe5b6020026020010151602001518a858151811061156357fe5b6020026020010151604001518b868151811061157b57fe5b6020026020010151606001518c878151811061159357fe5b6020026020010151608001518d88815181106115ab57fe5b602002602001015160c001518e89815181106115c357fe5b602002602001015160e001518f8a815181106115db57fe5b6020026020010151610100015161093d565b6115f657600080fd5b60010161137f565b5060008060006116258960008151811061161457fe5b60200260200101516000015161174b565b92509250925060008060006116408c60018151811061161457fe5b925092509250600a818161165057fe5b04600a8504600101146116755760405162461bcd60e51b815260040161055190612654565b8560ff16600114801561168b57508260ff166001145b6116a75760405162461bcd60e51b815260040161055190612520565b88156116d157604080518082019091526000808252602090910186905260075560088590556116f1565b6040805180820190915260008082526020909101869052600955600a8590555b7f045c41374537196136848ce62859f8aa22ae91b490103bd2bed75eac6a4d518089604051611720919061241e565b60405180910390a1505050505050505050505050565b60056020526000908152604090206001015481565b600080600060418451101561175f57600080fd5b60008460008151811061176e57fe5b0160200151602186015160419096015160f89190911c9690945092505050565b60038181548110610f9c57fe5b6004546060906000908190806117b057600080fd5b600019015b808214611803576000600260018484010104905085600482815481106117d757fe5b906000526020600020906002020160010154116117f6578092506117fd565b6001810391505b506117b5565b6004828154811061109757fe5b6000606060008d1561182f576118258c611022565b909250905061183e565b6118388c61179b565b90925090505b865186511461184c57600080fd5b845186511461185a57600080fd5b835186511461186857600080fd5b60005b8751811015611930576000811180156118ad575087600182038151811061188e57fe5b60200260200101518882815181106118a257fe5b602002602001015111155b806118cc575082518882815181106118c157fe5b602002602001015110155b156118dd5760009350505050611947565b828882815181106118ea57fe5b6020026020010151815181106118fc57fe5b602002602001015183828151811061191057fe5b6001600160a01b039092166020928302919091019091015260010161186b565b506119428d838d8d8d8d8c8c8c61093d565b925050505b9b9a5050505050505050505050565b6000546001600160a01b031681565b61196d6119df565b60038281548110610d3257fe5b8280548282559060005260206000209081019282156119cf579160200282015b828111156119cf57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019061199a565b506119db9291506119f9565b5090565b604051806040016040528060608152602001600081525090565b611a1d91905b808211156119db5780546001600160a01b03191681556001016119ff565b90565b80356001600160a01b0381168114610e6357600080fd5b600082601f830112611a47578081fd5b8135611a5a611a558261271e565b6126f8565b818152915060208083019084810181840286018201871015611a7b57600080fd5b60005b84811015611aa257611a908883611a20565b84529282019290820190600101611a7e565b505050505092915050565b600082601f830112611abd578081fd5b8135611acb611a558261271e565b818152915060208083019084810181840286018201871015611aec57600080fd5b60005b84811015611aa257611b018883611bea565b84529282019290820190600101611aef565b600082601f830112611b23578081fd5b8135611b31611a558261271e565b818152915060208083019084810181840286018201871015611b5257600080fd5b60005b84811015611aa257813584529282019290820190600101611b55565b600082601f830112611b81578081fd5b8135611b8f611a558261271e565b818152915060208083019084810181840286018201871015611bb057600080fd5b6000805b85811015611bde57823560ff81168114611bcc578283fd5b85529383019391830191600101611bb4565b50505050505092915050565b80358015158114610e6357600080fd5b600082601f830112611c0a578081fd5b81356001600160401b03811115611c1f578182fd5b611c32601f8201601f19166020016126f8565b9150808252836020828501011115611c4957600080fd5b8060208401602084013760009082016020015292915050565b6000610120808385031215611c75578182fd5b611c7e816126f8565b91505081356001600160401b0380821115611c9857600080fd5b611ca485838601611bfa565b83526020840135915080821115611cba57600080fd5b611cc685838601611b13565b60208401526040840135915080821115611cdf57600080fd5b611ceb85838601611aad565b6040840152606084013560608401526080840135608084015260a0840135915080821115611d1857600080fd5b611d2485838601611b13565b60a084015260c0840135915080821115611d3d57600080fd5b611d4985838601611b71565b60c084015260e0840135915080821115611d6257600080fd5b611d6e85838601611b13565b60e084015261010091508184013581811115611d8957600080fd5b611d9586828701611b13565b8385015250505092915050565b600060208284031215611db3578081fd5b611dbd8383611a20565b9392505050565b600080600080600060a08688031215611ddb578081fd5b85356001600160401b0380821115611df1578283fd5b611dfd89838a01611a37565b9650602088013595506040880135915080821115611e19578283fd5b611e2589838a01611b71565b94506060880135915080821115611e3a578283fd5b611e4689838a01611b13565b93506080880135915080821115611e5b578283fd5b50611e6888828901611b13565b9150509295509295909350565b60008060008060808587031215611e8a578384fd5b84356001600160401b03811115611e9f578485fd5b80860187601f820112611eb0578586fd5b80359150611ec0611a558361271e565b82815260208082019190838101895b86811015611ef857611ee68d848435890101611c62565b85529382019390820190600101611ecf565b5050819850808a013597505050505050611f158660408701611bea565b9150611f248660608701611bea565b905092959194509250565b60008060008060008060008060008060006101608c8e031215611f50578889fd5b611f5a8d8d611bea565b9a5060208c0135995060408c013598506001600160401b038060608e01351115611f82578687fd5b611f928e60608f01358f01611b13565b98508060808e01351115611fa4578687fd5b611fb48e60808f01358f01611aad565b975060a08d0135965060c08d013595508060e08e01351115611fd4578485fd5b611fe48e60e08f01358f01611b13565b9450806101008e01351115611ff7578384fd5b6120088e6101008f01358f01611b71565b9350806101208e0135111561201b578283fd5b61202c8e6101208f01358f01611b13565b9250806101408e0135111561203f578182fd5b506120518d6101408e01358e01611b13565b90509295989b509295989b9093969950565b60008060008060008060008060006101208a8c031215612081578283fd5b8935985060208a01356001600160401b038082111561209e578485fd5b6120aa8d838e01611a37565b995060408c01359150808211156120bf578485fd5b6120cb8d838e01611b13565b985060608c01359150808211156120e0578485fd5b6120ec8d838e01611aad565b975060808c0135965060a08c0135955060c08c013591508082111561210f578485fd5b61211b8d838e01611b71565b945060e08c0135915080821115612130578384fd5b61213c8d838e01611b13565b93506101008c0135915080821115612152578283fd5b5061215f8c828d01611b13565b9150509295985092959850929598565b60008060008060808587031215612184578182fd5b843593506020850135925060408501356001600160401b03808211156121a8578384fd5b6121b488838901611b13565b935060608701359150808211156121c9578283fd5b506121d687828801611aad565b91505092959194509250565b6000602082840312156121f3578081fd5b81356001600160401b03811115612208578182fd5b61221484828501611bfa565b949350505050565b6000806040838503121561222e578182fd5b82356001600160401b0380821115612244578384fd5b61225086838701611bfa565b9350602091508185013581811115612266578384fd5b85019050601f81018613612278578283fd5b6002612286611a558261273d565b8083865b848110156122b15761229f8b83358801611c62565b8452928601929086019060010161228a565b5096999098509650505050505050565b600080604083850312156122d3578182fd5b82356001600160401b03808211156122e9578384fd5b6122f586838701611bfa565b9350602085013591508082111561230a578283fd5b5061231785828601611c62565b9150509250929050565b60008060408385031215612333578182fd5b82356001600160401b03811115612348578283fd5b61235485828601611bfa565b95602094909401359450505050565b600060208284031215612374578081fd5b5035919050565b6000815180845260208085019450808401835b838110156123b35781516001600160a01b03168752958201959082019060010161238e565b509495945050505050565b90815260200190565b918252602082015260400190565b6001600160a01b0391909116815260200190565b600060208252611dbd602083018461237b565b60006040825261240f604083018561237b565b90508260208301529392505050565b901515815260200190565b93845260ff9290921660208401526040830152606082015260800190565b6020808252601490820152736e6f7420706175736564207269676874206e6f7760601b604082015260600190565b60208082526028908201527f6d7573742070726f76696465203220626c6f636b7320666f722066696e616c696040820152673a3c90383937b7b360c11b606082015260800190565b60208082526022908201527f736967496478206d757374206265207374726963746c7920696e6372656173696040820152616e6760f01b606082015260800190565b602080825260079082015266195e1c1a5c995960ca1b604082015260600190565b6020808252600c908201526b696e76616c6964206d65746160a01b604082015260600190565b6020808252601a908201527f63616e6e6f7420657874656e6420666f7220746f6f206c6f6e67000000000000604082015260600190565b6020808252600c908201526b1d5b985d5d1a1bdc9a5e995960a21b604082015260600190565b6020808252601b908201527f63616e6e6f74206368616e6765206f6c6420636f6d6d69747465650000000000604082015260600190565b602080825260139082015272736967496478206e6f7420696e2072616e676560681b604082015260600190565b60208082526010908201526f706175736564207269676874206e6f7760801b604082015260600190565b6020808252600990820152683737ba1030b236b4b760b91b604082015260600190565b6020808252601390820152721c1c9bdc1bdcd9551a5b59481a5b9d985b1a59606a1b604082015260600190565b60006020825282516040602084015261269d606084018261237b565b60208501516040850152809250505092915050565b60ff9390931683526020830191909152604082015260600190565b60ff958616815293909416602084015260408301919091526060820152608081019190915260a00190565b6040518181016001600160401b038111828210171561271657600080fd5b604052919050565b60006001600160401b03821115612733578081fd5b5060209081020190565b60006001600160401b03821115612752578081fd5b506020029056fea26469706673582212206cb9311bc80c1ee3245f8a0bacdbc0a6c9e15babd9d40e3acdc5116e2974dec664736f6c63430006060033"

// DeployIncognitoProxy deploys a new Ethereum contract, binding an instance of IncognitoProxy to it.
func DeployIncognitoProxy(auth *bind.TransactOpts, backend bind.ContractBackend, admin common.Address, beaconCommittee []common.Address, bridgeCommittee []common.Address) (common.Address, *types.Transaction, *IncognitoProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(IncognitoProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IncognitoProxyBin), backend, admin, beaconCommittee, bridgeCommittee)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IncognitoProxy{IncognitoProxyCaller: IncognitoProxyCaller{contract: contract}, IncognitoProxyTransactor: IncognitoProxyTransactor{contract: contract}, IncognitoProxyFilterer: IncognitoProxyFilterer{contract: contract}}, nil
}

// IncognitoProxy is an auto generated Go binding around an Ethereum contract.
type IncognitoProxy struct {
	IncognitoProxyCaller     // Read-only binding to the contract
	IncognitoProxyTransactor // Write-only binding to the contract
	IncognitoProxyFilterer   // Log filterer for contract events
}

// IncognitoProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type IncognitoProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncognitoProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IncognitoProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncognitoProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IncognitoProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncognitoProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IncognitoProxySession struct {
	Contract     *IncognitoProxy   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IncognitoProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IncognitoProxyCallerSession struct {
	Contract *IncognitoProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IncognitoProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IncognitoProxyTransactorSession struct {
	Contract     *IncognitoProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IncognitoProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type IncognitoProxyRaw struct {
	Contract *IncognitoProxy // Generic contract binding to access the raw methods on
}

// IncognitoProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IncognitoProxyCallerRaw struct {
	Contract *IncognitoProxyCaller // Generic read-only contract binding to access the raw methods on
}

// IncognitoProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IncognitoProxyTransactorRaw struct {
	Contract *IncognitoProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIncognitoProxy creates a new instance of IncognitoProxy, bound to a specific deployed contract.
func NewIncognitoProxy(address common.Address, backend bind.ContractBackend) (*IncognitoProxy, error) {
	contract, err := bindIncognitoProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IncognitoProxy{IncognitoProxyCaller: IncognitoProxyCaller{contract: contract}, IncognitoProxyTransactor: IncognitoProxyTransactor{contract: contract}, IncognitoProxyFilterer: IncognitoProxyFilterer{contract: contract}}, nil
}

// NewIncognitoProxyCaller creates a new read-only instance of IncognitoProxy, bound to a specific deployed contract.
func NewIncognitoProxyCaller(address common.Address, caller bind.ContractCaller) (*IncognitoProxyCaller, error) {
	contract, err := bindIncognitoProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IncognitoProxyCaller{contract: contract}, nil
}

// NewIncognitoProxyTransactor creates a new write-only instance of IncognitoProxy, bound to a specific deployed contract.
func NewIncognitoProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*IncognitoProxyTransactor, error) {
	contract, err := bindIncognitoProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IncognitoProxyTransactor{contract: contract}, nil
}

// NewIncognitoProxyFilterer creates a new log filterer instance of IncognitoProxy, bound to a specific deployed contract.
func NewIncognitoProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*IncognitoProxyFilterer, error) {
	contract, err := bindIncognitoProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IncognitoProxyFilterer{contract: contract}, nil
}

// bindIncognitoProxy binds a generic wrapper to an already deployed contract.
func bindIncognitoProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IncognitoProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IncognitoProxy *IncognitoProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IncognitoProxy.Contract.IncognitoProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IncognitoProxy *IncognitoProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.IncognitoProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IncognitoProxy *IncognitoProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.IncognitoProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IncognitoProxy *IncognitoProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IncognitoProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IncognitoProxy *IncognitoProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IncognitoProxy *IncognitoProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_IncognitoProxy *IncognitoProxyCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_IncognitoProxy *IncognitoProxySession) Admin() (common.Address, error) {
	return _IncognitoProxy.Contract.Admin(&_IncognitoProxy.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_IncognitoProxy *IncognitoProxyCallerSession) Admin() (common.Address, error) {
	return _IncognitoProxy.Contract.Admin(&_IncognitoProxy.CallOpts)
}

// BeaconCandidates is a free data retrieval call binding the contract method 0x62532738.
//
// Solidity: function beaconCandidates(uint256 ) view returns(uint256 startBlock)
func (_IncognitoProxy *IncognitoProxyCaller) BeaconCandidates(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "beaconCandidates", arg0)
	return *ret0, err
}

// BeaconCandidates is a free data retrieval call binding the contract method 0x62532738.
//
// Solidity: function beaconCandidates(uint256 ) view returns(uint256 startBlock)
func (_IncognitoProxy *IncognitoProxySession) BeaconCandidates(arg0 *big.Int) (*big.Int, error) {
	return _IncognitoProxy.Contract.BeaconCandidates(&_IncognitoProxy.CallOpts, arg0)
}

// BeaconCandidates is a free data retrieval call binding the contract method 0x62532738.
//
// Solidity: function beaconCandidates(uint256 ) view returns(uint256 startBlock)
func (_IncognitoProxy *IncognitoProxyCallerSession) BeaconCandidates(arg0 *big.Int) (*big.Int, error) {
	return _IncognitoProxy.Contract.BeaconCandidates(&_IncognitoProxy.CallOpts, arg0)
}

// BeaconCommittees is a free data retrieval call binding the contract method 0xf203a5ed.
//
// Solidity: function beaconCommittees(uint256 ) view returns(uint256 startBlock)
func (_IncognitoProxy *IncognitoProxyCaller) BeaconCommittees(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "beaconCommittees", arg0)
	return *ret0, err
}

// BeaconCommittees is a free data retrieval call binding the contract method 0xf203a5ed.
//
// Solidity: function beaconCommittees(uint256 ) view returns(uint256 startBlock)
func (_IncognitoProxy *IncognitoProxySession) BeaconCommittees(arg0 *big.Int) (*big.Int, error) {
	return _IncognitoProxy.Contract.BeaconCommittees(&_IncognitoProxy.CallOpts, arg0)
}

// BeaconCommittees is a free data retrieval call binding the contract method 0xf203a5ed.
//
// Solidity: function beaconCommittees(uint256 ) view returns(uint256 startBlock)
func (_IncognitoProxy *IncognitoProxyCallerSession) BeaconCommittees(arg0 *big.Int) (*big.Int, error) {
	return _IncognitoProxy.Contract.BeaconCommittees(&_IncognitoProxy.CallOpts, arg0)
}

// BeaconFinality is a free data retrieval call binding the contract method 0x5d7f86d9.
//
// Solidity: function beaconFinality() view returns(uint256 blockHeight, bytes32 rootHash)
func (_IncognitoProxy *IncognitoProxyCaller) BeaconFinality(opts *bind.CallOpts) (struct {
	BlockHeight *big.Int
	RootHash    [32]byte
}, error) {
	ret := new(struct {
		BlockHeight *big.Int
		RootHash    [32]byte
	})
	out := ret
	err := _IncognitoProxy.contract.Call(opts, out, "beaconFinality")
	return *ret, err
}

// BeaconFinality is a free data retrieval call binding the contract method 0x5d7f86d9.
//
// Solidity: function beaconFinality() view returns(uint256 blockHeight, bytes32 rootHash)
func (_IncognitoProxy *IncognitoProxySession) BeaconFinality() (struct {
	BlockHeight *big.Int
	RootHash    [32]byte
}, error) {
	return _IncognitoProxy.Contract.BeaconFinality(&_IncognitoProxy.CallOpts)
}

// BeaconFinality is a free data retrieval call binding the contract method 0x5d7f86d9.
//
// Solidity: function beaconFinality() view returns(uint256 blockHeight, bytes32 rootHash)
func (_IncognitoProxy *IncognitoProxyCallerSession) BeaconFinality() (struct {
	BlockHeight *big.Int
	RootHash    [32]byte
}, error) {
	return _IncognitoProxy.Contract.BeaconFinality(&_IncognitoProxy.CallOpts)
}

// BridgeCandidates is a free data retrieval call binding the contract method 0xe83d31ec.
//
// Solidity: function bridgeCandidates(uint256 ) view returns(uint256 startBlock)
func (_IncognitoProxy *IncognitoProxyCaller) BridgeCandidates(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "bridgeCandidates", arg0)
	return *ret0, err
}

// BridgeCandidates is a free data retrieval call binding the contract method 0xe83d31ec.
//
// Solidity: function bridgeCandidates(uint256 ) view returns(uint256 startBlock)
func (_IncognitoProxy *IncognitoProxySession) BridgeCandidates(arg0 *big.Int) (*big.Int, error) {
	return _IncognitoProxy.Contract.BridgeCandidates(&_IncognitoProxy.CallOpts, arg0)
}

// BridgeCandidates is a free data retrieval call binding the contract method 0xe83d31ec.
//
// Solidity: function bridgeCandidates(uint256 ) view returns(uint256 startBlock)
func (_IncognitoProxy *IncognitoProxyCallerSession) BridgeCandidates(arg0 *big.Int) (*big.Int, error) {
	return _IncognitoProxy.Contract.BridgeCandidates(&_IncognitoProxy.CallOpts, arg0)
}

// BridgeCommittees is a free data retrieval call binding the contract method 0x9b30b637.
//
// Solidity: function bridgeCommittees(uint256 ) view returns(uint256 startBlock)
func (_IncognitoProxy *IncognitoProxyCaller) BridgeCommittees(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "bridgeCommittees", arg0)
	return *ret0, err
}

// BridgeCommittees is a free data retrieval call binding the contract method 0x9b30b637.
//
// Solidity: function bridgeCommittees(uint256 ) view returns(uint256 startBlock)
func (_IncognitoProxy *IncognitoProxySession) BridgeCommittees(arg0 *big.Int) (*big.Int, error) {
	return _IncognitoProxy.Contract.BridgeCommittees(&_IncognitoProxy.CallOpts, arg0)
}

// BridgeCommittees is a free data retrieval call binding the contract method 0x9b30b637.
//
// Solidity: function bridgeCommittees(uint256 ) view returns(uint256 startBlock)
func (_IncognitoProxy *IncognitoProxyCallerSession) BridgeCommittees(arg0 *big.Int) (*big.Int, error) {
	return _IncognitoProxy.Contract.BridgeCommittees(&_IncognitoProxy.CallOpts, arg0)
}

// BridgeFinality is a free data retrieval call binding the contract method 0x2c1c9cac.
//
// Solidity: function bridgeFinality() view returns(uint256 blockHeight, bytes32 rootHash)
func (_IncognitoProxy *IncognitoProxyCaller) BridgeFinality(opts *bind.CallOpts) (struct {
	BlockHeight *big.Int
	RootHash    [32]byte
}, error) {
	ret := new(struct {
		BlockHeight *big.Int
		RootHash    [32]byte
	})
	out := ret
	err := _IncognitoProxy.contract.Call(opts, out, "bridgeFinality")
	return *ret, err
}

// BridgeFinality is a free data retrieval call binding the contract method 0x2c1c9cac.
//
// Solidity: function bridgeFinality() view returns(uint256 blockHeight, bytes32 rootHash)
func (_IncognitoProxy *IncognitoProxySession) BridgeFinality() (struct {
	BlockHeight *big.Int
	RootHash    [32]byte
}, error) {
	return _IncognitoProxy.Contract.BridgeFinality(&_IncognitoProxy.CallOpts)
}

// BridgeFinality is a free data retrieval call binding the contract method 0x2c1c9cac.
//
// Solidity: function bridgeFinality() view returns(uint256 blockHeight, bytes32 rootHash)
func (_IncognitoProxy *IncognitoProxyCallerSession) BridgeFinality() (struct {
	BlockHeight *big.Int
	RootHash    [32]byte
}, error) {
	return _IncognitoProxy.Contract.BridgeFinality(&_IncognitoProxy.CallOpts)
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() view returns(uint256)
func (_IncognitoProxy *IncognitoProxyCaller) Expire(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "expire")
	return *ret0, err
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() view returns(uint256)
func (_IncognitoProxy *IncognitoProxySession) Expire() (*big.Int, error) {
	return _IncognitoProxy.Contract.Expire(&_IncognitoProxy.CallOpts)
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() view returns(uint256)
func (_IncognitoProxy *IncognitoProxyCallerSession) Expire() (*big.Int, error) {
	return _IncognitoProxy.Contract.Expire(&_IncognitoProxy.CallOpts)
}

// ExtractCommitteeFromInstruction is a free data retrieval call binding the contract method 0x8eb60066.
//
// Solidity: function extractCommitteeFromInstruction(bytes inst, uint256 numVals) pure returns(address[])
func (_IncognitoProxy *IncognitoProxyCaller) ExtractCommitteeFromInstruction(opts *bind.CallOpts, inst []byte, numVals *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "extractCommitteeFromInstruction", inst, numVals)
	return *ret0, err
}

// ExtractCommitteeFromInstruction is a free data retrieval call binding the contract method 0x8eb60066.
//
// Solidity: function extractCommitteeFromInstruction(bytes inst, uint256 numVals) pure returns(address[])
func (_IncognitoProxy *IncognitoProxySession) ExtractCommitteeFromInstruction(inst []byte, numVals *big.Int) ([]common.Address, error) {
	return _IncognitoProxy.Contract.ExtractCommitteeFromInstruction(&_IncognitoProxy.CallOpts, inst, numVals)
}

// ExtractCommitteeFromInstruction is a free data retrieval call binding the contract method 0x8eb60066.
//
// Solidity: function extractCommitteeFromInstruction(bytes inst, uint256 numVals) pure returns(address[])
func (_IncognitoProxy *IncognitoProxyCallerSession) ExtractCommitteeFromInstruction(inst []byte, numVals *big.Int) ([]common.Address, error) {
	return _IncognitoProxy.Contract.ExtractCommitteeFromInstruction(&_IncognitoProxy.CallOpts, inst, numVals)
}

// ExtractDataFromBlockMerkleInstruction is a free data retrieval call binding the contract method 0xf179dc78.
//
// Solidity: function extractDataFromBlockMerkleInstruction(bytes inst) pure returns(uint8, bytes32, uint256)
func (_IncognitoProxy *IncognitoProxyCaller) ExtractDataFromBlockMerkleInstruction(opts *bind.CallOpts, inst []byte) (uint8, [32]byte, *big.Int, error) {
	var (
		ret0 = new(uint8)
		ret1 = new([32]byte)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _IncognitoProxy.contract.Call(opts, out, "extractDataFromBlockMerkleInstruction", inst)
	return *ret0, *ret1, *ret2, err
}

// ExtractDataFromBlockMerkleInstruction is a free data retrieval call binding the contract method 0xf179dc78.
//
// Solidity: function extractDataFromBlockMerkleInstruction(bytes inst) pure returns(uint8, bytes32, uint256)
func (_IncognitoProxy *IncognitoProxySession) ExtractDataFromBlockMerkleInstruction(inst []byte) (uint8, [32]byte, *big.Int, error) {
	return _IncognitoProxy.Contract.ExtractDataFromBlockMerkleInstruction(&_IncognitoProxy.CallOpts, inst)
}

// ExtractDataFromBlockMerkleInstruction is a free data retrieval call binding the contract method 0xf179dc78.
//
// Solidity: function extractDataFromBlockMerkleInstruction(bytes inst) pure returns(uint8, bytes32, uint256)
func (_IncognitoProxy *IncognitoProxyCallerSession) ExtractDataFromBlockMerkleInstruction(inst []byte) (uint8, [32]byte, *big.Int, error) {
	return _IncognitoProxy.Contract.ExtractDataFromBlockMerkleInstruction(&_IncognitoProxy.CallOpts, inst)
}

// ExtractMetaFromInstruction is a free data retrieval call binding the contract method 0x90500bae.
//
// Solidity: function extractMetaFromInstruction(bytes inst) pure returns(uint8, uint8, uint256, uint256, uint256)
func (_IncognitoProxy *IncognitoProxyCaller) ExtractMetaFromInstruction(opts *bind.CallOpts, inst []byte) (uint8, uint8, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(uint8)
		ret1 = new(uint8)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
		ret4 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _IncognitoProxy.contract.Call(opts, out, "extractMetaFromInstruction", inst)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// ExtractMetaFromInstruction is a free data retrieval call binding the contract method 0x90500bae.
//
// Solidity: function extractMetaFromInstruction(bytes inst) pure returns(uint8, uint8, uint256, uint256, uint256)
func (_IncognitoProxy *IncognitoProxySession) ExtractMetaFromInstruction(inst []byte) (uint8, uint8, *big.Int, *big.Int, *big.Int, error) {
	return _IncognitoProxy.Contract.ExtractMetaFromInstruction(&_IncognitoProxy.CallOpts, inst)
}

// ExtractMetaFromInstruction is a free data retrieval call binding the contract method 0x90500bae.
//
// Solidity: function extractMetaFromInstruction(bytes inst) pure returns(uint8, uint8, uint256, uint256, uint256)
func (_IncognitoProxy *IncognitoProxyCallerSession) ExtractMetaFromInstruction(inst []byte) (uint8, uint8, *big.Int, *big.Int, *big.Int, error) {
	return _IncognitoProxy.Contract.ExtractMetaFromInstruction(&_IncognitoProxy.CallOpts, inst)
}

// FindBeaconCommitteeFromHeight is a free data retrieval call binding the contract method 0xb600ffdb.
//
// Solidity: function findBeaconCommitteeFromHeight(uint256 blkHeight) view returns(address[], uint256)
func (_IncognitoProxy *IncognitoProxyCaller) FindBeaconCommitteeFromHeight(opts *bind.CallOpts, blkHeight *big.Int) ([]common.Address, *big.Int, error) {
	var (
		ret0 = new([]common.Address)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _IncognitoProxy.contract.Call(opts, out, "findBeaconCommitteeFromHeight", blkHeight)
	return *ret0, *ret1, err
}

// FindBeaconCommitteeFromHeight is a free data retrieval call binding the contract method 0xb600ffdb.
//
// Solidity: function findBeaconCommitteeFromHeight(uint256 blkHeight) view returns(address[], uint256)
func (_IncognitoProxy *IncognitoProxySession) FindBeaconCommitteeFromHeight(blkHeight *big.Int) ([]common.Address, *big.Int, error) {
	return _IncognitoProxy.Contract.FindBeaconCommitteeFromHeight(&_IncognitoProxy.CallOpts, blkHeight)
}

// FindBeaconCommitteeFromHeight is a free data retrieval call binding the contract method 0xb600ffdb.
//
// Solidity: function findBeaconCommitteeFromHeight(uint256 blkHeight) view returns(address[], uint256)
func (_IncognitoProxy *IncognitoProxyCallerSession) FindBeaconCommitteeFromHeight(blkHeight *big.Int) ([]common.Address, *big.Int, error) {
	return _IncognitoProxy.Contract.FindBeaconCommitteeFromHeight(&_IncognitoProxy.CallOpts, blkHeight)
}

// FindBridgeCommitteeFromHeight is a free data retrieval call binding the contract method 0xf5205fde.
//
// Solidity: function findBridgeCommitteeFromHeight(uint256 blkHeight) view returns(address[], uint256)
func (_IncognitoProxy *IncognitoProxyCaller) FindBridgeCommitteeFromHeight(opts *bind.CallOpts, blkHeight *big.Int) ([]common.Address, *big.Int, error) {
	var (
		ret0 = new([]common.Address)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _IncognitoProxy.contract.Call(opts, out, "findBridgeCommitteeFromHeight", blkHeight)
	return *ret0, *ret1, err
}

// FindBridgeCommitteeFromHeight is a free data retrieval call binding the contract method 0xf5205fde.
//
// Solidity: function findBridgeCommitteeFromHeight(uint256 blkHeight) view returns(address[], uint256)
func (_IncognitoProxy *IncognitoProxySession) FindBridgeCommitteeFromHeight(blkHeight *big.Int) ([]common.Address, *big.Int, error) {
	return _IncognitoProxy.Contract.FindBridgeCommitteeFromHeight(&_IncognitoProxy.CallOpts, blkHeight)
}

// FindBridgeCommitteeFromHeight is a free data retrieval call binding the contract method 0xf5205fde.
//
// Solidity: function findBridgeCommitteeFromHeight(uint256 blkHeight) view returns(address[], uint256)
func (_IncognitoProxy *IncognitoProxyCallerSession) FindBridgeCommitteeFromHeight(blkHeight *big.Int) ([]common.Address, *big.Int, error) {
	return _IncognitoProxy.Contract.FindBridgeCommitteeFromHeight(&_IncognitoProxy.CallOpts, blkHeight)
}

// GetBeaconCommittee is a free data retrieval call binding the contract method 0xfaea3167.
//
// Solidity: function getBeaconCommittee(uint256 i) view returns((address[],uint256))
func (_IncognitoProxy *IncognitoProxyCaller) GetBeaconCommittee(opts *bind.CallOpts, i *big.Int) (IncognitoProxyCommittee, error) {
	var (
		ret0 = new(IncognitoProxyCommittee)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "getBeaconCommittee", i)
	return *ret0, err
}

// GetBeaconCommittee is a free data retrieval call binding the contract method 0xfaea3167.
//
// Solidity: function getBeaconCommittee(uint256 i) view returns((address[],uint256))
func (_IncognitoProxy *IncognitoProxySession) GetBeaconCommittee(i *big.Int) (IncognitoProxyCommittee, error) {
	return _IncognitoProxy.Contract.GetBeaconCommittee(&_IncognitoProxy.CallOpts, i)
}

// GetBeaconCommittee is a free data retrieval call binding the contract method 0xfaea3167.
//
// Solidity: function getBeaconCommittee(uint256 i) view returns((address[],uint256))
func (_IncognitoProxy *IncognitoProxyCallerSession) GetBeaconCommittee(i *big.Int) (IncognitoProxyCommittee, error) {
	return _IncognitoProxy.Contract.GetBeaconCommittee(&_IncognitoProxy.CallOpts, i)
}

// GetBridgeCommittee is a free data retrieval call binding the contract method 0x8ceb69c3.
//
// Solidity: function getBridgeCommittee(uint256 i) view returns((address[],uint256))
func (_IncognitoProxy *IncognitoProxyCaller) GetBridgeCommittee(opts *bind.CallOpts, i *big.Int) (IncognitoProxyCommittee, error) {
	var (
		ret0 = new(IncognitoProxyCommittee)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "getBridgeCommittee", i)
	return *ret0, err
}

// GetBridgeCommittee is a free data retrieval call binding the contract method 0x8ceb69c3.
//
// Solidity: function getBridgeCommittee(uint256 i) view returns((address[],uint256))
func (_IncognitoProxy *IncognitoProxySession) GetBridgeCommittee(i *big.Int) (IncognitoProxyCommittee, error) {
	return _IncognitoProxy.Contract.GetBridgeCommittee(&_IncognitoProxy.CallOpts, i)
}

// GetBridgeCommittee is a free data retrieval call binding the contract method 0x8ceb69c3.
//
// Solidity: function getBridgeCommittee(uint256 i) view returns((address[],uint256))
func (_IncognitoProxy *IncognitoProxyCallerSession) GetBridgeCommittee(i *big.Int) (IncognitoProxyCommittee, error) {
	return _IncognitoProxy.Contract.GetBridgeCommittee(&_IncognitoProxy.CallOpts, i)
}

// InstructionApproved is a free data retrieval call binding the contract method 0xf65d2116.
//
// Solidity: function instructionApproved(bool isBeacon, bytes32 instHash, uint256 blkHeight, bytes32[] instPath, bool[] instPathIsLeft, bytes32 instRoot, bytes32 blkData, uint256[] sigIdx, uint8[] sigV, bytes32[] sigR, bytes32[] sigS) view returns(bool)
func (_IncognitoProxy *IncognitoProxyCaller) InstructionApproved(opts *bind.CallOpts, isBeacon bool, instHash [32]byte, blkHeight *big.Int, instPath [][32]byte, instPathIsLeft []bool, instRoot [32]byte, blkData [32]byte, sigIdx []*big.Int, sigV []uint8, sigR [][32]byte, sigS [][32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "instructionApproved", isBeacon, instHash, blkHeight, instPath, instPathIsLeft, instRoot, blkData, sigIdx, sigV, sigR, sigS)
	return *ret0, err
}

// InstructionApproved is a free data retrieval call binding the contract method 0xf65d2116.
//
// Solidity: function instructionApproved(bool isBeacon, bytes32 instHash, uint256 blkHeight, bytes32[] instPath, bool[] instPathIsLeft, bytes32 instRoot, bytes32 blkData, uint256[] sigIdx, uint8[] sigV, bytes32[] sigR, bytes32[] sigS) view returns(bool)
func (_IncognitoProxy *IncognitoProxySession) InstructionApproved(isBeacon bool, instHash [32]byte, blkHeight *big.Int, instPath [][32]byte, instPathIsLeft []bool, instRoot [32]byte, blkData [32]byte, sigIdx []*big.Int, sigV []uint8, sigR [][32]byte, sigS [][32]byte) (bool, error) {
	return _IncognitoProxy.Contract.InstructionApproved(&_IncognitoProxy.CallOpts, isBeacon, instHash, blkHeight, instPath, instPathIsLeft, instRoot, blkData, sigIdx, sigV, sigR, sigS)
}

// InstructionApproved is a free data retrieval call binding the contract method 0xf65d2116.
//
// Solidity: function instructionApproved(bool isBeacon, bytes32 instHash, uint256 blkHeight, bytes32[] instPath, bool[] instPathIsLeft, bytes32 instRoot, bytes32 blkData, uint256[] sigIdx, uint8[] sigV, bytes32[] sigR, bytes32[] sigS) view returns(bool)
func (_IncognitoProxy *IncognitoProxyCallerSession) InstructionApproved(isBeacon bool, instHash [32]byte, blkHeight *big.Int, instPath [][32]byte, instPathIsLeft []bool, instRoot [32]byte, blkData [32]byte, sigIdx []*big.Int, sigV []uint8, sigR [][32]byte, sigS [][32]byte) (bool, error) {
	return _IncognitoProxy.Contract.InstructionApproved(&_IncognitoProxy.CallOpts, isBeacon, instHash, blkHeight, instPath, instPathIsLeft, instRoot, blkData, sigIdx, sigV, sigR, sigS)
}

// InstructionApprovedByCommittee is a free data retrieval call binding the contract method 0x64a72862.
//
// Solidity: function instructionApprovedByCommittee(bytes32 instHash, address[] signers, bytes32[] instPath, bool[] instPathIsLeft, bytes32 instRoot, bytes32 blkData, uint8[] sigV, bytes32[] sigR, bytes32[] sigS) view returns(bool)
func (_IncognitoProxy *IncognitoProxyCaller) InstructionApprovedByCommittee(opts *bind.CallOpts, instHash [32]byte, signers []common.Address, instPath [][32]byte, instPathIsLeft []bool, instRoot [32]byte, blkData [32]byte, sigV []uint8, sigR [][32]byte, sigS [][32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "instructionApprovedByCommittee", instHash, signers, instPath, instPathIsLeft, instRoot, blkData, sigV, sigR, sigS)
	return *ret0, err
}

// InstructionApprovedByCommittee is a free data retrieval call binding the contract method 0x64a72862.
//
// Solidity: function instructionApprovedByCommittee(bytes32 instHash, address[] signers, bytes32[] instPath, bool[] instPathIsLeft, bytes32 instRoot, bytes32 blkData, uint8[] sigV, bytes32[] sigR, bytes32[] sigS) view returns(bool)
func (_IncognitoProxy *IncognitoProxySession) InstructionApprovedByCommittee(instHash [32]byte, signers []common.Address, instPath [][32]byte, instPathIsLeft []bool, instRoot [32]byte, blkData [32]byte, sigV []uint8, sigR [][32]byte, sigS [][32]byte) (bool, error) {
	return _IncognitoProxy.Contract.InstructionApprovedByCommittee(&_IncognitoProxy.CallOpts, instHash, signers, instPath, instPathIsLeft, instRoot, blkData, sigV, sigR, sigS)
}

// InstructionApprovedByCommittee is a free data retrieval call binding the contract method 0x64a72862.
//
// Solidity: function instructionApprovedByCommittee(bytes32 instHash, address[] signers, bytes32[] instPath, bool[] instPathIsLeft, bytes32 instRoot, bytes32 blkData, uint8[] sigV, bytes32[] sigR, bytes32[] sigS) view returns(bool)
func (_IncognitoProxy *IncognitoProxyCallerSession) InstructionApprovedByCommittee(instHash [32]byte, signers []common.Address, instPath [][32]byte, instPathIsLeft []bool, instRoot [32]byte, blkData [32]byte, sigV []uint8, sigR [][32]byte, sigS [][32]byte) (bool, error) {
	return _IncognitoProxy.Contract.InstructionApprovedByCommittee(&_IncognitoProxy.CallOpts, instHash, signers, instPath, instPathIsLeft, instRoot, blkData, sigV, sigR, sigS)
}

// InstructionInMerkleTree is a free data retrieval call binding the contract method 0x47c4b328.
//
// Solidity: function instructionInMerkleTree(bytes32 leaf, bytes32 root, bytes32[] path, bool[] left) pure returns(bool)
func (_IncognitoProxy *IncognitoProxyCaller) InstructionInMerkleTree(opts *bind.CallOpts, leaf [32]byte, root [32]byte, path [][32]byte, left []bool) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "instructionInMerkleTree", leaf, root, path, left)
	return *ret0, err
}

// InstructionInMerkleTree is a free data retrieval call binding the contract method 0x47c4b328.
//
// Solidity: function instructionInMerkleTree(bytes32 leaf, bytes32 root, bytes32[] path, bool[] left) pure returns(bool)
func (_IncognitoProxy *IncognitoProxySession) InstructionInMerkleTree(leaf [32]byte, root [32]byte, path [][32]byte, left []bool) (bool, error) {
	return _IncognitoProxy.Contract.InstructionInMerkleTree(&_IncognitoProxy.CallOpts, leaf, root, path, left)
}

// InstructionInMerkleTree is a free data retrieval call binding the contract method 0x47c4b328.
//
// Solidity: function instructionInMerkleTree(bytes32 leaf, bytes32 root, bytes32[] path, bool[] left) pure returns(bool)
func (_IncognitoProxy *IncognitoProxyCallerSession) InstructionInMerkleTree(leaf [32]byte, root [32]byte, path [][32]byte, left []bool) (bool, error) {
	return _IncognitoProxy.Contract.InstructionInMerkleTree(&_IncognitoProxy.CallOpts, leaf, root, path, left)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IncognitoProxy *IncognitoProxyCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IncognitoProxy *IncognitoProxySession) Paused() (bool, error) {
	return _IncognitoProxy.Contract.Paused(&_IncognitoProxy.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IncognitoProxy *IncognitoProxyCallerSession) Paused() (bool, error) {
	return _IncognitoProxy.Contract.Paused(&_IncognitoProxy.CallOpts)
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() view returns(address)
func (_IncognitoProxy *IncognitoProxyCaller) Successor(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "successor")
	return *ret0, err
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() view returns(address)
func (_IncognitoProxy *IncognitoProxySession) Successor() (common.Address, error) {
	return _IncognitoProxy.Contract.Successor(&_IncognitoProxy.CallOpts)
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() view returns(address)
func (_IncognitoProxy *IncognitoProxyCallerSession) Successor() (common.Address, error) {
	return _IncognitoProxy.Contract.Successor(&_IncognitoProxy.CallOpts)
}

// VerifySig is a free data retrieval call binding the contract method 0x3aacfdad.
//
// Solidity: function verifySig(address[] committee, bytes32 msgHash, uint8[] v, bytes32[] r, bytes32[] s) pure returns(bool)
func (_IncognitoProxy *IncognitoProxyCaller) VerifySig(opts *bind.CallOpts, committee []common.Address, msgHash [32]byte, v []uint8, r [][32]byte, s [][32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "verifySig", committee, msgHash, v, r, s)
	return *ret0, err
}

// VerifySig is a free data retrieval call binding the contract method 0x3aacfdad.
//
// Solidity: function verifySig(address[] committee, bytes32 msgHash, uint8[] v, bytes32[] r, bytes32[] s) pure returns(bool)
func (_IncognitoProxy *IncognitoProxySession) VerifySig(committee []common.Address, msgHash [32]byte, v []uint8, r [][32]byte, s [][32]byte) (bool, error) {
	return _IncognitoProxy.Contract.VerifySig(&_IncognitoProxy.CallOpts, committee, msgHash, v, r, s)
}

// VerifySig is a free data retrieval call binding the contract method 0x3aacfdad.
//
// Solidity: function verifySig(address[] committee, bytes32 msgHash, uint8[] v, bytes32[] r, bytes32[] s) pure returns(bool)
func (_IncognitoProxy *IncognitoProxyCallerSession) VerifySig(committee []common.Address, msgHash [32]byte, v []uint8, r [][32]byte, s [][32]byte) (bool, error) {
	return _IncognitoProxy.Contract.VerifySig(&_IncognitoProxy.CallOpts, committee, msgHash, v, r, s)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_IncognitoProxy *IncognitoProxyTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncognitoProxy.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_IncognitoProxy *IncognitoProxySession) Claim() (*types.Transaction, error) {
	return _IncognitoProxy.Contract.Claim(&_IncognitoProxy.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_IncognitoProxy *IncognitoProxyTransactorSession) Claim() (*types.Transaction, error) {
	return _IncognitoProxy.Contract.Claim(&_IncognitoProxy.TransactOpts)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_IncognitoProxy *IncognitoProxyTransactor) Extend(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _IncognitoProxy.contract.Transact(opts, "extend", n)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_IncognitoProxy *IncognitoProxySession) Extend(n *big.Int) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.Extend(&_IncognitoProxy.TransactOpts, n)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_IncognitoProxy *IncognitoProxyTransactorSession) Extend(n *big.Int) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.Extend(&_IncognitoProxy.TransactOpts, n)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IncognitoProxy *IncognitoProxyTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncognitoProxy.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IncognitoProxy *IncognitoProxySession) Pause() (*types.Transaction, error) {
	return _IncognitoProxy.Contract.Pause(&_IncognitoProxy.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IncognitoProxy *IncognitoProxyTransactorSession) Pause() (*types.Transaction, error) {
	return _IncognitoProxy.Contract.Pause(&_IncognitoProxy.TransactOpts)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_IncognitoProxy *IncognitoProxyTransactor) Retire(opts *bind.TransactOpts, _successor common.Address) (*types.Transaction, error) {
	return _IncognitoProxy.contract.Transact(opts, "retire", _successor)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_IncognitoProxy *IncognitoProxySession) Retire(_successor common.Address) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.Retire(&_IncognitoProxy.TransactOpts, _successor)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_IncognitoProxy *IncognitoProxyTransactorSession) Retire(_successor common.Address) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.Retire(&_IncognitoProxy.TransactOpts, _successor)
}

// SubmitBeaconCandidate is a paid mutator transaction binding the contract method 0x52d45239.
//
// Solidity: function submitBeaconCandidate(bytes inst, (bytes,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[]) instProof) returns()
func (_IncognitoProxy *IncognitoProxyTransactor) SubmitBeaconCandidate(opts *bind.TransactOpts, inst []byte, instProof IncognitoProxyInstructionProof) (*types.Transaction, error) {
	return _IncognitoProxy.contract.Transact(opts, "submitBeaconCandidate", inst, instProof)
}

// SubmitBeaconCandidate is a paid mutator transaction binding the contract method 0x52d45239.
//
// Solidity: function submitBeaconCandidate(bytes inst, (bytes,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[]) instProof) returns()
func (_IncognitoProxy *IncognitoProxySession) SubmitBeaconCandidate(inst []byte, instProof IncognitoProxyInstructionProof) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.SubmitBeaconCandidate(&_IncognitoProxy.TransactOpts, inst, instProof)
}

// SubmitBeaconCandidate is a paid mutator transaction binding the contract method 0x52d45239.
//
// Solidity: function submitBeaconCandidate(bytes inst, (bytes,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[]) instProof) returns()
func (_IncognitoProxy *IncognitoProxyTransactorSession) SubmitBeaconCandidate(inst []byte, instProof IncognitoProxyInstructionProof) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.SubmitBeaconCandidate(&_IncognitoProxy.TransactOpts, inst, instProof)
}

// SubmitBridgeCandidate is a paid mutator transaction binding the contract method 0x7e59f5bf.
//
// Solidity: function submitBridgeCandidate(bytes inst, (bytes,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])[2] instProofs) returns()
func (_IncognitoProxy *IncognitoProxyTransactor) SubmitBridgeCandidate(opts *bind.TransactOpts, inst []byte, instProofs [2]IncognitoProxyInstructionProof) (*types.Transaction, error) {
	return _IncognitoProxy.contract.Transact(opts, "submitBridgeCandidate", inst, instProofs)
}

// SubmitBridgeCandidate is a paid mutator transaction binding the contract method 0x7e59f5bf.
//
// Solidity: function submitBridgeCandidate(bytes inst, (bytes,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])[2] instProofs) returns()
func (_IncognitoProxy *IncognitoProxySession) SubmitBridgeCandidate(inst []byte, instProofs [2]IncognitoProxyInstructionProof) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.SubmitBridgeCandidate(&_IncognitoProxy.TransactOpts, inst, instProofs)
}

// SubmitBridgeCandidate is a paid mutator transaction binding the contract method 0x7e59f5bf.
//
// Solidity: function submitBridgeCandidate(bytes inst, (bytes,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])[2] instProofs) returns()
func (_IncognitoProxy *IncognitoProxyTransactorSession) SubmitBridgeCandidate(inst []byte, instProofs [2]IncognitoProxyInstructionProof) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.SubmitBridgeCandidate(&_IncognitoProxy.TransactOpts, inst, instProofs)
}

// SubmitFinalityProof is a paid mutator transaction binding the contract method 0xcc554f00.
//
// Solidity: function submitFinalityProof((bytes,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])[] instProofs, uint256 swapID, bool isCandidate, bool isBeacon) returns()
func (_IncognitoProxy *IncognitoProxyTransactor) SubmitFinalityProof(opts *bind.TransactOpts, instProofs []IncognitoProxyInstructionProof, swapID *big.Int, isCandidate bool, isBeacon bool) (*types.Transaction, error) {
	return _IncognitoProxy.contract.Transact(opts, "submitFinalityProof", instProofs, swapID, isCandidate, isBeacon)
}

// SubmitFinalityProof is a paid mutator transaction binding the contract method 0xcc554f00.
//
// Solidity: function submitFinalityProof((bytes,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])[] instProofs, uint256 swapID, bool isCandidate, bool isBeacon) returns()
func (_IncognitoProxy *IncognitoProxySession) SubmitFinalityProof(instProofs []IncognitoProxyInstructionProof, swapID *big.Int, isCandidate bool, isBeacon bool) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.SubmitFinalityProof(&_IncognitoProxy.TransactOpts, instProofs, swapID, isCandidate, isBeacon)
}

// SubmitFinalityProof is a paid mutator transaction binding the contract method 0xcc554f00.
//
// Solidity: function submitFinalityProof((bytes,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])[] instProofs, uint256 swapID, bool isCandidate, bool isBeacon) returns()
func (_IncognitoProxy *IncognitoProxyTransactorSession) SubmitFinalityProof(instProofs []IncognitoProxyInstructionProof, swapID *big.Int, isCandidate bool, isBeacon bool) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.SubmitFinalityProof(&_IncognitoProxy.TransactOpts, instProofs, swapID, isCandidate, isBeacon)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IncognitoProxy *IncognitoProxyTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncognitoProxy.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IncognitoProxy *IncognitoProxySession) Unpause() (*types.Transaction, error) {
	return _IncognitoProxy.Contract.Unpause(&_IncognitoProxy.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IncognitoProxy *IncognitoProxyTransactorSession) Unpause() (*types.Transaction, error) {
	return _IncognitoProxy.Contract.Unpause(&_IncognitoProxy.TransactOpts)
}

// IncognitoProxyBeaconCommitteeSwappedIterator is returned from FilterBeaconCommitteeSwapped and is used to iterate over the raw logs and unpacked data for BeaconCommitteeSwapped events raised by the IncognitoProxy contract.
type IncognitoProxyBeaconCommitteeSwappedIterator struct {
	Event *IncognitoProxyBeaconCommitteeSwapped // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncognitoProxyBeaconCommitteeSwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncognitoProxyBeaconCommitteeSwapped)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncognitoProxyBeaconCommitteeSwapped)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncognitoProxyBeaconCommitteeSwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncognitoProxyBeaconCommitteeSwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncognitoProxyBeaconCommitteeSwapped represents a BeaconCommitteeSwapped event raised by the IncognitoProxy contract.
type IncognitoProxyBeaconCommitteeSwapped struct {
	Id          *big.Int
	StartHeight *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBeaconCommitteeSwapped is a free log retrieval operation binding the contract event 0xe15e1a9dec6ad906dd5985b062bfa5ee8bc5d5738e46e4deb8a2df2fbbbb59d1.
//
// Solidity: event BeaconCommitteeSwapped(uint256 id, uint256 startHeight)
func (_IncognitoProxy *IncognitoProxyFilterer) FilterBeaconCommitteeSwapped(opts *bind.FilterOpts) (*IncognitoProxyBeaconCommitteeSwappedIterator, error) {

	logs, sub, err := _IncognitoProxy.contract.FilterLogs(opts, "BeaconCommitteeSwapped")
	if err != nil {
		return nil, err
	}
	return &IncognitoProxyBeaconCommitteeSwappedIterator{contract: _IncognitoProxy.contract, event: "BeaconCommitteeSwapped", logs: logs, sub: sub}, nil
}

// WatchBeaconCommitteeSwapped is a free log subscription operation binding the contract event 0xe15e1a9dec6ad906dd5985b062bfa5ee8bc5d5738e46e4deb8a2df2fbbbb59d1.
//
// Solidity: event BeaconCommitteeSwapped(uint256 id, uint256 startHeight)
func (_IncognitoProxy *IncognitoProxyFilterer) WatchBeaconCommitteeSwapped(opts *bind.WatchOpts, sink chan<- *IncognitoProxyBeaconCommitteeSwapped) (event.Subscription, error) {

	logs, sub, err := _IncognitoProxy.contract.WatchLogs(opts, "BeaconCommitteeSwapped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncognitoProxyBeaconCommitteeSwapped)
				if err := _IncognitoProxy.contract.UnpackLog(event, "BeaconCommitteeSwapped", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconCommitteeSwapped is a log parse operation binding the contract event 0xe15e1a9dec6ad906dd5985b062bfa5ee8bc5d5738e46e4deb8a2df2fbbbb59d1.
//
// Solidity: event BeaconCommitteeSwapped(uint256 id, uint256 startHeight)
func (_IncognitoProxy *IncognitoProxyFilterer) ParseBeaconCommitteeSwapped(log types.Log) (*IncognitoProxyBeaconCommitteeSwapped, error) {
	event := new(IncognitoProxyBeaconCommitteeSwapped)
	if err := _IncognitoProxy.contract.UnpackLog(event, "BeaconCommitteeSwapped", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncognitoProxyBridgeCommitteeSwappedIterator is returned from FilterBridgeCommitteeSwapped and is used to iterate over the raw logs and unpacked data for BridgeCommitteeSwapped events raised by the IncognitoProxy contract.
type IncognitoProxyBridgeCommitteeSwappedIterator struct {
	Event *IncognitoProxyBridgeCommitteeSwapped // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncognitoProxyBridgeCommitteeSwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncognitoProxyBridgeCommitteeSwapped)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncognitoProxyBridgeCommitteeSwapped)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncognitoProxyBridgeCommitteeSwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncognitoProxyBridgeCommitteeSwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncognitoProxyBridgeCommitteeSwapped represents a BridgeCommitteeSwapped event raised by the IncognitoProxy contract.
type IncognitoProxyBridgeCommitteeSwapped struct {
	Id          *big.Int
	StartHeight *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBridgeCommitteeSwapped is a free log retrieval operation binding the contract event 0xfd354f8fe3f6f03db5436879221c99b65c610515bb33434539060150d61f8a44.
//
// Solidity: event BridgeCommitteeSwapped(uint256 id, uint256 startHeight)
func (_IncognitoProxy *IncognitoProxyFilterer) FilterBridgeCommitteeSwapped(opts *bind.FilterOpts) (*IncognitoProxyBridgeCommitteeSwappedIterator, error) {

	logs, sub, err := _IncognitoProxy.contract.FilterLogs(opts, "BridgeCommitteeSwapped")
	if err != nil {
		return nil, err
	}
	return &IncognitoProxyBridgeCommitteeSwappedIterator{contract: _IncognitoProxy.contract, event: "BridgeCommitteeSwapped", logs: logs, sub: sub}, nil
}

// WatchBridgeCommitteeSwapped is a free log subscription operation binding the contract event 0xfd354f8fe3f6f03db5436879221c99b65c610515bb33434539060150d61f8a44.
//
// Solidity: event BridgeCommitteeSwapped(uint256 id, uint256 startHeight)
func (_IncognitoProxy *IncognitoProxyFilterer) WatchBridgeCommitteeSwapped(opts *bind.WatchOpts, sink chan<- *IncognitoProxyBridgeCommitteeSwapped) (event.Subscription, error) {

	logs, sub, err := _IncognitoProxy.contract.WatchLogs(opts, "BridgeCommitteeSwapped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncognitoProxyBridgeCommitteeSwapped)
				if err := _IncognitoProxy.contract.UnpackLog(event, "BridgeCommitteeSwapped", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBridgeCommitteeSwapped is a log parse operation binding the contract event 0xfd354f8fe3f6f03db5436879221c99b65c610515bb33434539060150d61f8a44.
//
// Solidity: event BridgeCommitteeSwapped(uint256 id, uint256 startHeight)
func (_IncognitoProxy *IncognitoProxyFilterer) ParseBridgeCommitteeSwapped(log types.Log) (*IncognitoProxyBridgeCommitteeSwapped, error) {
	event := new(IncognitoProxyBridgeCommitteeSwapped)
	if err := _IncognitoProxy.contract.UnpackLog(event, "BridgeCommitteeSwapped", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncognitoProxyChainFinalizedIterator is returned from FilterChainFinalized and is used to iterate over the raw logs and unpacked data for ChainFinalized events raised by the IncognitoProxy contract.
type IncognitoProxyChainFinalizedIterator struct {
	Event *IncognitoProxyChainFinalized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncognitoProxyChainFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncognitoProxyChainFinalized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncognitoProxyChainFinalized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncognitoProxyChainFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncognitoProxyChainFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncognitoProxyChainFinalized represents a ChainFinalized event raised by the IncognitoProxy contract.
type IncognitoProxyChainFinalized struct {
	IsBeacon bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChainFinalized is a free log retrieval operation binding the contract event 0x045c41374537196136848ce62859f8aa22ae91b490103bd2bed75eac6a4d5180.
//
// Solidity: event ChainFinalized(bool isBeacon)
func (_IncognitoProxy *IncognitoProxyFilterer) FilterChainFinalized(opts *bind.FilterOpts) (*IncognitoProxyChainFinalizedIterator, error) {

	logs, sub, err := _IncognitoProxy.contract.FilterLogs(opts, "ChainFinalized")
	if err != nil {
		return nil, err
	}
	return &IncognitoProxyChainFinalizedIterator{contract: _IncognitoProxy.contract, event: "ChainFinalized", logs: logs, sub: sub}, nil
}

// WatchChainFinalized is a free log subscription operation binding the contract event 0x045c41374537196136848ce62859f8aa22ae91b490103bd2bed75eac6a4d5180.
//
// Solidity: event ChainFinalized(bool isBeacon)
func (_IncognitoProxy *IncognitoProxyFilterer) WatchChainFinalized(opts *bind.WatchOpts, sink chan<- *IncognitoProxyChainFinalized) (event.Subscription, error) {

	logs, sub, err := _IncognitoProxy.contract.WatchLogs(opts, "ChainFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncognitoProxyChainFinalized)
				if err := _IncognitoProxy.contract.UnpackLog(event, "ChainFinalized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChainFinalized is a log parse operation binding the contract event 0x045c41374537196136848ce62859f8aa22ae91b490103bd2bed75eac6a4d5180.
//
// Solidity: event ChainFinalized(bool isBeacon)
func (_IncognitoProxy *IncognitoProxyFilterer) ParseChainFinalized(log types.Log) (*IncognitoProxyChainFinalized, error) {
	event := new(IncognitoProxyChainFinalized)
	if err := _IncognitoProxy.contract.UnpackLog(event, "ChainFinalized", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncognitoProxyClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the IncognitoProxy contract.
type IncognitoProxyClaimIterator struct {
	Event *IncognitoProxyClaim // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncognitoProxyClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncognitoProxyClaim)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncognitoProxyClaim)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncognitoProxyClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncognitoProxyClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncognitoProxyClaim represents a Claim event raised by the IncognitoProxy contract.
type IncognitoProxyClaim struct {
	Claimer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_IncognitoProxy *IncognitoProxyFilterer) FilterClaim(opts *bind.FilterOpts) (*IncognitoProxyClaimIterator, error) {

	logs, sub, err := _IncognitoProxy.contract.FilterLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return &IncognitoProxyClaimIterator{contract: _IncognitoProxy.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_IncognitoProxy *IncognitoProxyFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *IncognitoProxyClaim) (event.Subscription, error) {

	logs, sub, err := _IncognitoProxy.contract.WatchLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncognitoProxyClaim)
				if err := _IncognitoProxy.contract.UnpackLog(event, "Claim", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaim is a log parse operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_IncognitoProxy *IncognitoProxyFilterer) ParseClaim(log types.Log) (*IncognitoProxyClaim, error) {
	event := new(IncognitoProxyClaim)
	if err := _IncognitoProxy.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncognitoProxyExtendIterator is returned from FilterExtend and is used to iterate over the raw logs and unpacked data for Extend events raised by the IncognitoProxy contract.
type IncognitoProxyExtendIterator struct {
	Event *IncognitoProxyExtend // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncognitoProxyExtendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncognitoProxyExtend)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncognitoProxyExtend)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncognitoProxyExtendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncognitoProxyExtendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncognitoProxyExtend represents a Extend event raised by the IncognitoProxy contract.
type IncognitoProxyExtend struct {
	Ndays *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterExtend is a free log retrieval operation binding the contract event 0x02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8.
//
// Solidity: event Extend(uint256 ndays)
func (_IncognitoProxy *IncognitoProxyFilterer) FilterExtend(opts *bind.FilterOpts) (*IncognitoProxyExtendIterator, error) {

	logs, sub, err := _IncognitoProxy.contract.FilterLogs(opts, "Extend")
	if err != nil {
		return nil, err
	}
	return &IncognitoProxyExtendIterator{contract: _IncognitoProxy.contract, event: "Extend", logs: logs, sub: sub}, nil
}

// WatchExtend is a free log subscription operation binding the contract event 0x02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8.
//
// Solidity: event Extend(uint256 ndays)
func (_IncognitoProxy *IncognitoProxyFilterer) WatchExtend(opts *bind.WatchOpts, sink chan<- *IncognitoProxyExtend) (event.Subscription, error) {

	logs, sub, err := _IncognitoProxy.contract.WatchLogs(opts, "Extend")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncognitoProxyExtend)
				if err := _IncognitoProxy.contract.UnpackLog(event, "Extend", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExtend is a log parse operation binding the contract event 0x02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8.
//
// Solidity: event Extend(uint256 ndays)
func (_IncognitoProxy *IncognitoProxyFilterer) ParseExtend(log types.Log) (*IncognitoProxyExtend, error) {
	event := new(IncognitoProxyExtend)
	if err := _IncognitoProxy.contract.UnpackLog(event, "Extend", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncognitoProxyPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the IncognitoProxy contract.
type IncognitoProxyPausedIterator struct {
	Event *IncognitoProxyPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncognitoProxyPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncognitoProxyPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncognitoProxyPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncognitoProxyPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncognitoProxyPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncognitoProxyPaused represents a Paused event raised by the IncognitoProxy contract.
type IncognitoProxyPaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address pauser)
func (_IncognitoProxy *IncognitoProxyFilterer) FilterPaused(opts *bind.FilterOpts) (*IncognitoProxyPausedIterator, error) {

	logs, sub, err := _IncognitoProxy.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &IncognitoProxyPausedIterator{contract: _IncognitoProxy.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address pauser)
func (_IncognitoProxy *IncognitoProxyFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *IncognitoProxyPaused) (event.Subscription, error) {

	logs, sub, err := _IncognitoProxy.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncognitoProxyPaused)
				if err := _IncognitoProxy.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address pauser)
func (_IncognitoProxy *IncognitoProxyFilterer) ParsePaused(log types.Log) (*IncognitoProxyPaused, error) {
	event := new(IncognitoProxyPaused)
	if err := _IncognitoProxy.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncognitoProxySubmittedBridgeCandidateIterator is returned from FilterSubmittedBridgeCandidate and is used to iterate over the raw logs and unpacked data for SubmittedBridgeCandidate events raised by the IncognitoProxy contract.
type IncognitoProxySubmittedBridgeCandidateIterator struct {
	Event *IncognitoProxySubmittedBridgeCandidate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncognitoProxySubmittedBridgeCandidateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncognitoProxySubmittedBridgeCandidate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncognitoProxySubmittedBridgeCandidate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncognitoProxySubmittedBridgeCandidateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncognitoProxySubmittedBridgeCandidateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncognitoProxySubmittedBridgeCandidate represents a SubmittedBridgeCandidate event raised by the IncognitoProxy contract.
type IncognitoProxySubmittedBridgeCandidate struct {
	Id          *big.Int
	StartHeight *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSubmittedBridgeCandidate is a free log retrieval operation binding the contract event 0xc918d4d247e9af7d4a8ef8ac6f0078262006c87209d062a6d846f5b547dcff66.
//
// Solidity: event SubmittedBridgeCandidate(uint256 id, uint256 startHeight)
func (_IncognitoProxy *IncognitoProxyFilterer) FilterSubmittedBridgeCandidate(opts *bind.FilterOpts) (*IncognitoProxySubmittedBridgeCandidateIterator, error) {

	logs, sub, err := _IncognitoProxy.contract.FilterLogs(opts, "SubmittedBridgeCandidate")
	if err != nil {
		return nil, err
	}
	return &IncognitoProxySubmittedBridgeCandidateIterator{contract: _IncognitoProxy.contract, event: "SubmittedBridgeCandidate", logs: logs, sub: sub}, nil
}

// WatchSubmittedBridgeCandidate is a free log subscription operation binding the contract event 0xc918d4d247e9af7d4a8ef8ac6f0078262006c87209d062a6d846f5b547dcff66.
//
// Solidity: event SubmittedBridgeCandidate(uint256 id, uint256 startHeight)
func (_IncognitoProxy *IncognitoProxyFilterer) WatchSubmittedBridgeCandidate(opts *bind.WatchOpts, sink chan<- *IncognitoProxySubmittedBridgeCandidate) (event.Subscription, error) {

	logs, sub, err := _IncognitoProxy.contract.WatchLogs(opts, "SubmittedBridgeCandidate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncognitoProxySubmittedBridgeCandidate)
				if err := _IncognitoProxy.contract.UnpackLog(event, "SubmittedBridgeCandidate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSubmittedBridgeCandidate is a log parse operation binding the contract event 0xc918d4d247e9af7d4a8ef8ac6f0078262006c87209d062a6d846f5b547dcff66.
//
// Solidity: event SubmittedBridgeCandidate(uint256 id, uint256 startHeight)
func (_IncognitoProxy *IncognitoProxyFilterer) ParseSubmittedBridgeCandidate(log types.Log) (*IncognitoProxySubmittedBridgeCandidate, error) {
	event := new(IncognitoProxySubmittedBridgeCandidate)
	if err := _IncognitoProxy.contract.UnpackLog(event, "SubmittedBridgeCandidate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncognitoProxyUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the IncognitoProxy contract.
type IncognitoProxyUnpausedIterator struct {
	Event *IncognitoProxyUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncognitoProxyUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncognitoProxyUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncognitoProxyUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncognitoProxyUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncognitoProxyUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncognitoProxyUnpaused represents a Unpaused event raised by the IncognitoProxy contract.
type IncognitoProxyUnpaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address pauser)
func (_IncognitoProxy *IncognitoProxyFilterer) FilterUnpaused(opts *bind.FilterOpts) (*IncognitoProxyUnpausedIterator, error) {

	logs, sub, err := _IncognitoProxy.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &IncognitoProxyUnpausedIterator{contract: _IncognitoProxy.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address pauser)
func (_IncognitoProxy *IncognitoProxyFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *IncognitoProxyUnpaused) (event.Subscription, error) {

	logs, sub, err := _IncognitoProxy.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncognitoProxyUnpaused)
				if err := _IncognitoProxy.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address pauser)
func (_IncognitoProxy *IncognitoProxyFilterer) ParseUnpaused(log types.Log) (*IncognitoProxyUnpaused, error) {
	event := new(IncognitoProxyUnpaused)
	if err := _IncognitoProxy.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}
