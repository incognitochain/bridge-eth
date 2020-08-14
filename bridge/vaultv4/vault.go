// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vaultv4

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

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
var AdminPausableBin = "0x608060405234801561001057600080fd5b506040516106fc3803806106fc8339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b03199092169190911790556001805460ff60a01b191690556301e1338042016002556106808061007c6000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c806379599f961161006657806379599f96146100ea5780638456cb59146101045780639714378c1461010c5780639e6371ba14610129578063f851a4401461014f57610093565b80633f4ba83a146100985780634e71d92d146100a25780635c975abb146100aa5780636ff968c3146100c6575b600080fd5b6100a0610157565b005b6100a0610239565b6100b2610320565b604080519115158252519081900360200190f35b6100ce610330565b604080516001600160a01b039092168252519081900360200190f35b6100f261033f565b60408051918252519081900360200190f35b6100a0610345565b6100a06004803603602081101561012257600080fd5b503561046a565b6100a06004803603602081101561013f57600080fd5b50356001600160a01b031661058e565b6100ce61063b565b6000546001600160a01b031633146101a2576040805162461bcd60e51b81526020600482015260096024820152683737ba1030b236b4b760b91b604482015290519081900360640190fd5b600154600160a01b900460ff166101f7576040805162461bcd60e51b81526020600482015260146024820152736e6f7420706175736564207269676874206e6f7760601b604482015290519081900360640190fd5b6001805460ff60a01b191690556040805133815290517f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa9181900360200190a1565b6002544210610279576040805162461bcd60e51b8152602060048201526007602482015266195e1c1a5c995960ca1b604482015290519081900360640190fd5b6001546001600160a01b031633146102c7576040805162461bcd60e51b815260206004820152600c60248201526b1d5b985d5d1a1bdc9a5e995960a21b604482015290519081900360640190fd5b600154600080546001600160a01b0319166001600160a01b0392831617908190556040805191909216815290517f0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc9181900360200190a1565b600154600160a01b900460ff1681565b6001546001600160a01b031681565b60025481565b6000546001600160a01b03163314610390576040805162461bcd60e51b81526020600482015260096024820152683737ba1030b236b4b760b91b604482015290519081900360640190fd5b600154600160a01b900460ff16156103e2576040805162461bcd60e51b815260206004820152601060248201526f706175736564207269676874206e6f7760801b604482015290519081900360640190fd5b6002544210610422576040805162461bcd60e51b8152602060048201526007602482015266195e1c1a5c995960ca1b604482015290519081900360640190fd5b6001805460ff60a01b1916600160a01b1790556040805133815290517f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2589181900360200190a1565b6000546001600160a01b031633146104b5576040805162461bcd60e51b81526020600482015260096024820152683737ba1030b236b4b760b91b604482015290519081900360640190fd5b60025442106104f5576040805162461bcd60e51b8152602060048201526007602482015266195e1c1a5c995960ca1b604482015290519081900360640190fd5b61016e811061054b576040805162461bcd60e51b815260206004820152601a60248201527f63616e6e6f7420657874656e6420666f7220746f6f206c6f6e67000000000000604482015290519081900360640190fd5b600280546201518083020190556040805182815290517f02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e89181900360200190a150565b6000546001600160a01b031633146105d9576040805162461bcd60e51b81526020600482015260096024820152683737ba1030b236b4b760b91b604482015290519081900360640190fd5b6002544210610619576040805162461bcd60e51b8152602060048201526007602482015266195e1c1a5c995960ca1b604482015290519081900360640190fd5b600180546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b03168156fea2646970667358221220a2ca9ea2af34ec0d679e90a5a8e57199554aa8802806fba59805f7a240718bb964736f6c634300060c0033"

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
// Solidity: function admin() constant returns(address)
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
// Solidity: function admin() constant returns(address)
func (_AdminPausable *AdminPausableSession) Admin() (common.Address, error) {
	return _AdminPausable.Contract.Admin(&_AdminPausable.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_AdminPausable *AdminPausableCallerSession) Admin() (common.Address, error) {
	return _AdminPausable.Contract.Admin(&_AdminPausable.CallOpts)
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() constant returns(uint256)
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
// Solidity: function expire() constant returns(uint256)
func (_AdminPausable *AdminPausableSession) Expire() (*big.Int, error) {
	return _AdminPausable.Contract.Expire(&_AdminPausable.CallOpts)
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() constant returns(uint256)
func (_AdminPausable *AdminPausableCallerSession) Expire() (*big.Int, error) {
	return _AdminPausable.Contract.Expire(&_AdminPausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
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
// Solidity: function paused() constant returns(bool)
func (_AdminPausable *AdminPausableSession) Paused() (bool, error) {
	return _AdminPausable.Contract.Paused(&_AdminPausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_AdminPausable *AdminPausableCallerSession) Paused() (bool, error) {
	return _AdminPausable.Contract.Paused(&_AdminPausable.CallOpts)
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() constant returns(address)
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
// Solidity: function successor() constant returns(address)
func (_AdminPausable *AdminPausableSession) Successor() (common.Address, error) {
	return _AdminPausable.Contract.Successor(&_AdminPausable.CallOpts)
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() constant returns(address)
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

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_IERC20 *IERC20Caller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_IERC20 *IERC20Session) Decimals() (*big.Int, error) {
	return _IERC20.Contract.Decimals(&_IERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_IERC20 *IERC20CallerSession) Decimals() (*big.Int, error) {
	return _IERC20.Contract.Decimals(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns()
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns()
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns()
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns()
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns()
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns()
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns()
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns()
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns()
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncognitoABI is the input ABI used to generate the binding from.
const IncognitoABI = "[{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"name\":\"instructionApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IncognitoFuncSigs maps the 4-byte function signature to its string representation.
var IncognitoFuncSigs = map[string]string{
	"f65d2116": "instructionApproved(bool,bytes32,uint256,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])",
}

// Incognito is an auto generated Go binding around an Ethereum contract.
type Incognito struct {
	IncognitoCaller     // Read-only binding to the contract
	IncognitoTransactor // Write-only binding to the contract
	IncognitoFilterer   // Log filterer for contract events
}

// IncognitoCaller is an auto generated read-only Go binding around an Ethereum contract.
type IncognitoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncognitoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IncognitoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncognitoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IncognitoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncognitoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IncognitoSession struct {
	Contract     *Incognito        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IncognitoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IncognitoCallerSession struct {
	Contract *IncognitoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IncognitoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IncognitoTransactorSession struct {
	Contract     *IncognitoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IncognitoRaw is an auto generated low-level Go binding around an Ethereum contract.
type IncognitoRaw struct {
	Contract *Incognito // Generic contract binding to access the raw methods on
}

// IncognitoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IncognitoCallerRaw struct {
	Contract *IncognitoCaller // Generic read-only contract binding to access the raw methods on
}

// IncognitoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IncognitoTransactorRaw struct {
	Contract *IncognitoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIncognito creates a new instance of Incognito, bound to a specific deployed contract.
func NewIncognito(address common.Address, backend bind.ContractBackend) (*Incognito, error) {
	contract, err := bindIncognito(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Incognito{IncognitoCaller: IncognitoCaller{contract: contract}, IncognitoTransactor: IncognitoTransactor{contract: contract}, IncognitoFilterer: IncognitoFilterer{contract: contract}}, nil
}

// NewIncognitoCaller creates a new read-only instance of Incognito, bound to a specific deployed contract.
func NewIncognitoCaller(address common.Address, caller bind.ContractCaller) (*IncognitoCaller, error) {
	contract, err := bindIncognito(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IncognitoCaller{contract: contract}, nil
}

// NewIncognitoTransactor creates a new write-only instance of Incognito, bound to a specific deployed contract.
func NewIncognitoTransactor(address common.Address, transactor bind.ContractTransactor) (*IncognitoTransactor, error) {
	contract, err := bindIncognito(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IncognitoTransactor{contract: contract}, nil
}

// NewIncognitoFilterer creates a new log filterer instance of Incognito, bound to a specific deployed contract.
func NewIncognitoFilterer(address common.Address, filterer bind.ContractFilterer) (*IncognitoFilterer, error) {
	contract, err := bindIncognito(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IncognitoFilterer{contract: contract}, nil
}

// bindIncognito binds a generic wrapper to an already deployed contract.
func bindIncognito(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IncognitoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Incognito *IncognitoRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Incognito.Contract.IncognitoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Incognito *IncognitoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Incognito.Contract.IncognitoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Incognito *IncognitoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Incognito.Contract.IncognitoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Incognito *IncognitoCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Incognito.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Incognito *IncognitoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Incognito.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Incognito *IncognitoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Incognito.Contract.contract.Transact(opts, method, params...)
}

// InstructionApproved is a free data retrieval call binding the contract method 0xf65d2116.
//
// Solidity: function instructionApproved(bool , bytes32 , uint256 , bytes32[] , bool[] , bytes32 , bytes32 , uint256[] , uint8[] , bytes32[] , bytes32[] ) constant returns(bool)
func (_Incognito *IncognitoCaller) InstructionApproved(opts *bind.CallOpts, arg0 bool, arg1 [32]byte, arg2 *big.Int, arg3 [][32]byte, arg4 []bool, arg5 [32]byte, arg6 [32]byte, arg7 []*big.Int, arg8 []uint8, arg9 [][32]byte, arg10 [][32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Incognito.contract.Call(opts, out, "instructionApproved", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
	return *ret0, err
}

// InstructionApproved is a free data retrieval call binding the contract method 0xf65d2116.
//
// Solidity: function instructionApproved(bool , bytes32 , uint256 , bytes32[] , bool[] , bytes32 , bytes32 , uint256[] , uint8[] , bytes32[] , bytes32[] ) constant returns(bool)
func (_Incognito *IncognitoSession) InstructionApproved(arg0 bool, arg1 [32]byte, arg2 *big.Int, arg3 [][32]byte, arg4 []bool, arg5 [32]byte, arg6 [32]byte, arg7 []*big.Int, arg8 []uint8, arg9 [][32]byte, arg10 [][32]byte) (bool, error) {
	return _Incognito.Contract.InstructionApproved(&_Incognito.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
}

// InstructionApproved is a free data retrieval call binding the contract method 0xf65d2116.
//
// Solidity: function instructionApproved(bool , bytes32 , uint256 , bytes32[] , bool[] , bytes32 , bytes32 , uint256[] , uint8[] , bytes32[] , bytes32[] ) constant returns(bool)
func (_Incognito *IncognitoCallerSession) InstructionApproved(arg0 bool, arg1 [32]byte, arg2 *big.Int, arg3 [][32]byte, arg4 []bool, arg5 [32]byte, arg6 [32]byte, arg7 []*big.Int, arg8 []uint8, arg9 [][32]byte, arg10 [][32]byte) (bool, error) {
	return _Incognito.Contract.InstructionApproved(&_Incognito.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220d809b993fc3dff01408e5d7de00e443e3d8fb8a6a506486c66a3140dbe39ac3d64736f6c634300060c0033"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// VaultABI is the input ABI used to generate the binding from.
const VaultABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"incognitoProxyAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_prevVault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ndays\",\"type\":\"uint256\"}],\"name\":\"Extend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVault\",\"type\":\"address\"}],\"name\":\"Migrate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"}],\"name\":\"MoveAssets\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newIncognitoProxy\",\"type\":\"address\"}],\"name\":\"UpdateIncognitoProxy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"UpdateTokenTotal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ETH_TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipientToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"exchangeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"timestamp\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signData\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"recipientTokens\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"exchangeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"timestamp\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signData\",\"type\":\"bytes\"}],\"name\":\"executeMulti\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expire\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"extend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getDepositedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incognito\",\"outputs\":[{\"internalType\":\"contractIncognito\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"isSigDataUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"isWithdrawed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_newVault\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"migration\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"}],\"name\":\"moveAssets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newVault\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"notEntered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"}],\"name\":\"parseBurnInst\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prevVault\",\"outputs\":[{\"internalType\":\"contractWithdrawable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"timestamp\",\"type\":\"bytes\"}],\"name\":\"requestWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_successor\",\"type\":\"address\"}],\"name\":\"retire\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"sigDataUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"sigToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"uint256[2]\",\"name\":\"heights\",\"type\":\"uint256[2]\"},{\"internalType\":\"bytes32[][2]\",\"name\":\"instPaths\",\"type\":\"bytes32[][2]\"},{\"internalType\":\"bool[][2]\",\"name\":\"instPathIsLefts\",\"type\":\"bool[][2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"instRoots\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"blkData\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint256[][2]\",\"name\":\"sigIdxs\",\"type\":\"uint256[][2]\"},{\"internalType\":\"uint8[][2]\",\"name\":\"sigVs\",\"type\":\"uint8[][2]\"},{\"internalType\":\"bytes32[][2]\",\"name\":\"sigRs\",\"type\":\"bytes32[][2]\"},{\"internalType\":\"bytes32[][2]\",\"name\":\"sigSs\",\"type\":\"bytes32[][2]\"}],\"name\":\"submitBurnProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"successor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalDepositedToSCAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"updateAssets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newIncognitoProxy\",\"type\":\"address\"}],\"name\":\"updateIncognitoProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"uint256[2]\",\"name\":\"heights\",\"type\":\"uint256[2]\"},{\"internalType\":\"bytes32[][2]\",\"name\":\"instPaths\",\"type\":\"bytes32[][2]\"},{\"internalType\":\"bool[][2]\",\"name\":\"instPathIsLefts\",\"type\":\"bool[][2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"instRoots\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"blkData\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint256[][2]\",\"name\":\"sigIdxs\",\"type\":\"uint256[][2]\"},{\"internalType\":\"uint8[][2]\",\"name\":\"sigVs\",\"type\":\"uint8[][2]\"},{\"internalType\":\"bytes32[][2]\",\"name\":\"sigRs\",\"type\":\"bytes32[][2]\"},{\"internalType\":\"bytes32[][2]\",\"name\":\"sigSs\",\"type\":\"bytes32[][2]\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawRequests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"withdrawed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// VaultFuncSigs maps the 4-byte function signature to its string representation.
var VaultFuncSigs = map[string]string{
	"58bc8337": "ETH_TOKEN()",
	"f851a440": "admin()",
	"70a08231": "balanceOf(address)",
	"4e71d92d": "claim()",
	"a26e1186": "deposit(string)",
	"5a67cb87": "depositERC20(address,uint256,string)",
	"8588ccd6": "execute(address,uint256,address,address,bytes,bytes,bytes)",
	"b69f6511": "executeMulti(address[],uint256[],address[],address,bytes,bytes,bytes)",
	"79599f96": "expire()",
	"9714378c": "extend(uint256)",
	"cf54aaa0": "getDecimals(address)",
	"f75b98ce": "getDepositedBalance(address,address)",
	"8a984538": "incognito()",
	"e4bd7074": "isSigDataUsed(bytes32)",
	"749c5f86": "isWithdrawed(bytes32)",
	"ce5494bb": "migrate(address)",
	"995fac11": "migration(address,address)",
	"0c4f5039": "moveAssets(address[])",
	"88aaf0c8": "newVault()",
	"a3f5d8cc": "notEntered()",
	"7e16e6e1": "parseBurnInst(bytes)",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"fa84702e": "prevVault()",
	"87add440": "requestWithdraw(string,address,uint256,bytes,bytes)",
	"9e6371ba": "retire(address)",
	"1ea1940e": "sigDataUsed(bytes32)",
	"3fec6b40": "sigToAddress(bytes,bytes32)",
	"3cb4b58a": "submitBurnProof(bytes,uint256[2],bytes32[][2],bool[][2],bytes32[2],bytes32[2],uint256[][2],uint8[][2],bytes32[][2],bytes32[][2])",
	"6ff968c3": "successor()",
	"6304541c": "totalDepositedToSCAmount(address)",
	"3f4ba83a": "unpause()",
	"1ed4276d": "updateAssets(address[],uint256[])",
	"3a51913d": "updateIncognitoProxy(address)",
	"f8b39745": "withdraw(bytes,uint256[2],bytes32[][2],bool[][2],bytes32[2],bytes32[2],uint256[][2],uint8[][2],bytes32[][2],bytes32[][2])",
	"65b5a00f": "withdrawRequests(address,address)",
	"dca40d9e": "withdrawed(bytes32)",
}

// VaultBin is the compiled bytecode used for deploying new contracts.
var VaultBin = "0x6080604052600a805460ff60a01b1916600160a01b1790553480156200002457600080fd5b506040516200499e3803806200499e8339810160408190526200004791620000aa565b600080546001600160a01b03199081166001600160a01b03958616179091556001805460ff60a01b19169055426301e1338001600255600880548216938516939093179092556009805483169190931617909155600a8054909116905562000116565b600080600060608486031215620000bf578283fd5b8351620000cc81620000fd565b6020850151909350620000df81620000fd565b6040850151909250620000f281620000fd565b809150509250925092565b6001600160a01b03811681146200011357600080fd5b50565b61487880620001266000396000f3fe60806040526004361061021e5760003560e01c80638456cb5911610123578063a3f5d8cc116100ab578063e4bd70741161006f578063e4bd7074146105d8578063f75b98ce146105f8578063f851a44014610618578063f8b397451461062d578063fa84702e1461064d57610225565b8063a3f5d8cc14610543578063b69f651114610558578063ce5494bb1461056b578063cf54aaa01461058b578063dca40d9e146105b857610225565b80638a984538116100f25780638a984538146104bb5780639714378c146104d0578063995fac11146104f05780639e6371ba14610510578063a26e11861461053057610225565b80638456cb591461045e5780638588ccd61461047357806387add4401461048657806388aaf0c8146104a657610225565b80635a67cb87116101a65780636ff968c3116101755780636ff968c3146103c357806370a08231146103d8578063749c5f86146103f857806379599f96146104185780637e16e6e11461042d57610225565b80635a67cb871461034e5780635c975abb146103615780636304541c1461037657806365b5a00f146103a357610225565b80633cb4b58a116101ed5780633cb4b58a146102c25780633f4ba83a146102e25780633fec6b40146102f75780634e71d92d1461032457806358bc83371461033957610225565b80630c4f50391461022a5780631ea1940e1461024c5780631ed4276d146102825780633a51913d146102a257610225565b3661022557005b600080fd5b34801561023657600080fd5b5061024a610245366004613d60565b610662565b005b34801561025857600080fd5b5061026c610267366004613f3a565b610aed565b604051610279919061457a565b60405180910390f35b34801561028e57600080fd5b5061026c61029d366004613cf8565b610b02565b3480156102ae57600080fd5b5061024a6102bd366004613b5d565b610d23565b3480156102ce57600080fd5b5061024a6102dd366004613f84565b610e01565b3480156102ee57600080fd5b5061024a6110bd565b34801561030357600080fd5b506103176103123660046140cf565b611157565b604051610279919061441b565b34801561033057600080fd5b5061024a6111e6565b34801561034557600080fd5b50610317611282565b61024a61035c366004613ca2565b611287565b34801561036d57600080fd5b5061026c61154a565b34801561038257600080fd5b50610396610391366004613b5d565b61155a565b6040516102799190614632565b3480156103af57600080fd5b506103966103be366004613ba6565b61156c565b3480156103cf57600080fd5b50610317611589565b3480156103e457600080fd5b506103966103f3366004613b5d565b611598565b34801561040457600080fd5b5061026c610413366004613f3a565b611633565b34801561042457600080fd5b506103966116ea565b34801561043957600080fd5b5061044d610448366004613f52565b6116f0565b604051610279959493929190614773565b34801561046a57600080fd5b5061024a611752565b61024a610481366004613bde565b61180a565b34801561049257600080fd5b5061024a6104a1366004614111565b611b9f565b3480156104b257600080fd5b50610317611db5565b3480156104c757600080fd5b50610317611dc4565b3480156104dc57600080fd5b5061024a6104eb366004613f3a565b611dd3565b3480156104fc57600080fd5b5061026c61050b366004613ba6565b611e7c565b34801561051c57600080fd5b5061024a61052b366004613b5d565b611e9c565b61024a61053e366004613f52565b611f09565b34801561054f57600080fd5b5061026c612001565b61024a610566366004613e5b565b612011565b34801561057757600080fd5b5061024a610586366004613b5d565b612816565b34801561059757600080fd5b506105ab6105a6366004613b5d565b6128e9565b6040516102799190614765565b3480156105c457600080fd5b5061026c6105d3366004613f3a565b612967565b3480156105e457600080fd5b5061026c6105f3366004613f3a565b61297c565b34801561060457600080fd5b50610396610613366004613ba6565b6129e3565b34801561062457600080fd5b50610317612b0f565b34801561063957600080fd5b5061024a610648366004613f84565b612b1e565b34801561065957600080fd5b50610317612ebb565b6000546001600160a01b031633146106955760405162461bcd60e51b815260040161068c90614742565b60405180910390fd5b600154600160a01b900460ff166106be5760405162461bcd60e51b815260040161068c9061466c565b600a546001600160a01b031615156106d6600b612eca565b906106f45760405162461bcd60e51b815260040161068c9190614659565b50606081516001600160401b038111801561070e57600080fd5b50604051908082528060200260200182016040528015610738578160200160208202803683370190505b50905060005b8251811015610a045760006001600160a01b031683828151811061075e57fe5b60200260200101516001600160a01b0316141561084d576000805260076020527f6d5257204ebe7d88fd91ae87941cb2dd9d8062b64ae5a2bd2d28ec40b9fbf6df5482518390839081106107ae57fe5b6020908102919091010152600a546040516000916001600160a01b03169047906107d790614418565b60006040518083038185875af1925050503d8060008114610814576040519150601f19603f3d011682016040523d82523d6000602084013e610819565b606091505b50509050806108286004612eca565b906108465760405162461bcd60e51b815260040161068c9190614659565b50506109c6565b600083828151811061085b57fe5b60200260200101516001600160a01b03166370a08231306040518263ffffffff1660e01b815260040161088e919061441b565b60206040518083038186803b1580156108a657600080fd5b505afa1580156108ba573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108de91906141b3565b90508015610973578382815181106108f257fe5b6020908102919091010151600a5460405163a9059cbb60e01b81526001600160a01b039283169263a9059cbb92610930929116908590600401614453565b600060405180830381600087803b15801561094a57600080fd5b505af115801561095e573d6000803e3d6000fd5b5050505061096a612ff3565b61097357600080fd5b6007600085848151811061098357fe5b60200260200101516001600160a01b03166001600160a01b03168152602001908152602001600020548383815181106109b857fe5b602002602001018181525050505b6000600760008584815181106109d857fe5b6020908102919091018101516001600160a01b031682528101919091526040016000205560010161073e565b50600a54604051631ed4276d60e01b81526001600160a01b0390911690631ed4276d90610a37908590859060040161454c565b602060405180830381600087803b158015610a5157600080fd5b505af1158015610a65573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a899190613f1e565b610a936004612eca565b90610ab15760405162461bcd60e51b815260040161068c9190614659565b507f492fc8b292f2a2a9b328a366b83745f30c024056d12aa118a15966d26a8ce65882604051610ae19190614539565b60405180910390a15050565b60046020526000908152604090205460ff1681565b6009546000906001600160a01b031615801590610b2957506009546001600160a01b031633145b610b33600c612eca565b90610b515760405162461bcd60e51b815260040161068c9190614659565b50838214610b5f600a612eca565b90610b7d5760405162461bcd60e51b815260040161068c9190614659565b50600960009054906101000a90046001600160a01b03166001600160a01b0316635c975abb6040518163ffffffff1660e01b815260040160206040518083038186803b158015610bcc57600080fd5b505afa158015610be0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c049190613f1e565b610c0e600d612eca565b90610c2c5760405162461bcd60e51b815260040161068c9190614659565b5060005b84811015610cda57610c93848483818110610c4757fe5b9050602002013560076000898986818110610c5e57fe5b9050602002016020810190610c739190613b5d565b6001600160a01b0316815260208101919091526040016000205490613027565b60076000888885818110610ca357fe5b9050602002016020810190610cb89190613b5d565b6001600160a01b03168152602081019190915260400160002055600101610c30565b507f6a7fbbcddfd518bb8c56b28ac6c7acb0f7ca093ed232eb3306e53d14e469895f85858585604051610d1094939291906144ba565b60405180910390a1506001949350505050565b6000546001600160a01b03163314610d4d5760405162461bcd60e51b815260040161068c90614742565b600154600160a01b900460ff16610d765760405162461bcd60e51b815260040161068c9061466c565b6001600160a01b0381161515610d8c600b612eca565b90610daa5760405162461bcd60e51b815260040161068c9190614659565b50600880546001600160a01b0319166001600160a01b0383161790556040517f204252dfe190ad6ef63db40a490f048b39f661de74628408f13cd0bb2d4c344690610df690839061441b565b60405180910390a150565b600154600160a01b900460ff1615610e2b5760405162461bcd60e51b815260040161068c90614718565b600a54600160a01b900460ff16610e426001612eca565b90610e605760405162461bcd60e51b815260040161068c9190614659565b50600a805460ff60a01b19169055600080808080610e7d8f6116f0565b945094509450945094508460ff166061148015610e9d57508360ff166001145b610ea657600080fd5b6001600160a01b038316610f09576001600160a01b038316600090815260076020526040902054610ed8908290613027565b471015610ee56007612eca565b90610f035760405162461bcd60e51b815260040161068c9190614659565b50610ffe565b6000610f14846128e9565b905060098160ff161115610f315760098160ff1603600a0a820291505b6001600160a01b038416600090815260076020526040902054610f55908390613027565b6040516370a0823160e01b81526001600160a01b038616906370a0823190610f8190309060040161441b565b60206040518083038186803b158015610f9957600080fd5b505afa158015610fad573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fd191906141b3565b1015610fdd6007612eca565b90610ffb5760405162461bcd60e51b815260040161068c9190614659565b50505b6110108f8f8f8f8f8f8f8f8f8f613045565b6001600160a01b038083166000908152600560209081526040808320938716835292905220546110409082613027565b6001600160a01b038084166000908152600560209081526040808320938816835292815282822093909355600790925290205461107d9082613027565b6001600160a01b039093166000908152600760205260409020929092555050600a805460ff60a01b1916600160a01b179055505050505050505050505050565b6000546001600160a01b031633146110e75760405162461bcd60e51b815260040161068c90614742565b600154600160a01b900460ff166111105760405162461bcd60e51b815260040161068c9061466c565b6001805460ff60a01b191690556040517f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa9061114d90339061441b565b60405180910390a1565b60008060008060208601519150604086015192508560408151811061117857fe5b602001015160f81c60f81b60f81c601b019050600185828486604051600081526020016040526040516111ae949392919061463b565b6020604051602081039080840390855afa1580156111d0573d6000803e3d6000fd5b5050506020604051035193505050505b92915050565b60025442106112075760405162461bcd60e51b815260040161068c9061469a565b6001546001600160a01b031633146112315760405162461bcd60e51b815260040161068c906146f2565b600154600080546001600160a01b0319166001600160a01b0392831617908190556040517f0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc9261114d92169061441b565b600081565b600154600160a01b900460ff16156112b15760405162461bcd60e51b815260040161068c90614718565b600a54600160a01b900460ff166112c86001612eca565b906112e65760405162461bcd60e51b815260040161068c9190614659565b50600a805460ff60a01b19169055826000611300826128e9565b90506000826001600160a01b03166370a08231306040518263ffffffff1660e01b8152600401611330919061441b565b60206040518083038186803b15801561134857600080fd5b505afa15801561135c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061138091906141b3565b90508085600960ff851611156113bb5760098460ff1603600a0a81816113a257fe5b04905060098460ff1603600a0a83816113b757fe5b0492505b670de0b6b3a764000081111580156113db5750670de0b6b3a76400008311155b80156113f85750670de0b6b3a76400006113f58285613027565b11155b6114026003612eca565b906114205760405162461bcd60e51b815260040161068c9190614659565b506040516323b872dd60e01b81526001600160a01b038616906323b872dd9061145190339030908c9060040161442f565b600060405180830381600087803b15801561146b57600080fd5b505af115801561147f573d6000803e3d6000fd5b5050505061148b612ff3565b6114956004612eca565b906114b35760405162461bcd60e51b815260040161068c9190614659565b50866114c8836114c28b611598565b906132ba565b146114d3600a612eca565b906114f15760405162461bcd60e51b815260040161068c9190614659565b507f2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e88878360405161152593929190614486565b60405180910390a15050600a805460ff60a01b1916600160a01b179055505050505050565b600154600160a01b900460ff1681565b60076020526000908152604090205481565b600560209081526000928352604080842090915290825290205481565b6001546001600160a01b031681565b60006001600160a01b0382166115af57504761162e565b6040516370a0823160e01b81526001600160a01b038316906370a08231906115db90309060040161441b565b60206040518083038186803b1580156115f357600080fd5b505afa158015611607573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061162b91906141b3565b90505b919050565b60008181526003602052604081205460ff16156116525750600161162e565b6009546001600160a01b031661166a5750600061162e565b600954604051633a4e2fc360e11b81526001600160a01b039091169063749c5f869061169a908590600401614632565b60206040518083038186803b1580156116b257600080fd5b505afa1580156116c6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061162b9190613f1e565b60025481565b6000806000806000808660008151811061170657fe5b602001015160f81c60f81b60f81c905060008760018151811061172557fe5b0160200151602289015160428a01516062909a0151939a60f89290921c9990985096509194509092505050565b6000546001600160a01b0316331461177c5760405162461bcd60e51b815260040161068c90614742565b600154600160a01b900460ff16156117a65760405162461bcd60e51b815260040161068c90614718565b60025442106117c75760405162461bcd60e51b815260040161068c9061469a565b6001805460ff60a01b1916600160a01b1790556040517f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2589061114d90339061441b565b600154600160a01b900460ff16156118345760405162461bcd60e51b815260040161068c90614718565b600a54600160a01b900460ff1661184b6001612eca565b906118695760405162461bcd60e51b815260040161068c9190614659565b50600a805460ff60a01b191690556040516000906118aa90611895908790879087908c90602001614345565b604051602081830303815290604052836132cf565b90506118b6818961333b565b6001600160a01b038082166000908152600560209081526040808320938c16835292905220548711156118e96008612eca565b906119075760405162461bcd60e51b815260040161068c9190614659565b506001600160a01b03881660009081526007602052604090205461192b90886132ba565b6001600160a01b03808a166000818152600760209081526040808320959095559285168152600583528381209181529152205461196890886132ba565b6001600160a01b038083166000908152600560209081526040808320938d168084529390915290209190915534906119ab576119a48189613027565b9050611ae6565b6040516370a0823160e01b815288906001600160a01b038b16906370a08231906119d990309060040161441b565b60206040518083038186803b1580156119f157600080fd5b505afa158015611a05573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a2991906141b3565b1015611a356007612eca565b90611a535760405162461bcd60e51b815260040161068c9190614659565b5060405163a9059cbb60e01b81526001600160a01b038a169063a9059cbb90611a829089908c90600401614453565b600060405180830381600087803b158015611a9c57600080fd5b505af1158015611ab0573d6000803e3d6000fd5b50505050611abc612ff3565b611ac66004612eca565b90611ae45760405162461bcd60e51b815260040161068c9190614659565b505b6000611af48883888a61347d565b6001600160a01b038085166000908152600560209081526040808320938d1683529290522054909150611b279082613027565b6001600160a01b038085166000908152600560209081526040808320938d168352928152828220939093556007909252902054611b649082613027565b6001600160a01b039098166000908152600760205260409020979097555050600a805460ff60a01b1916600160a01b17905550505050505050565b600154600160a01b900460ff1615611bc95760405162461bcd60e51b815260040161068c90614718565b600a54600160a01b900460ff16611be06001612eca565b90611bfe5760405162461bcd60e51b815260040161068c9190614659565b50600a805460ff60a01b19169055604051600090611c3f90611c2a9088908890869089906020016143d9565b604051602081830303815290604052846132cf565b9050611c4b818661333b565b6001600160a01b03808216600090815260056020908152604080832093891683529290522054841115611c7e6008612eca565b90611c9c5760405162461bcd60e51b815260040161068c9190614659565b506001600160a01b03808216600090815260056020908152604080832093891683529290522054611ccd90856132ba565b6001600160a01b038083166000908152600560209081526040808320938a168352928152828220939093556007909252902054611d0a90856132ba565b6001600160a01b038616600081815260076020526040902091909155849015611d5e576000611d38876128e9565b905060098160ff161115611d5c5760098160ff1603600a0a8681611d5857fe5b0491505b505b7f2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e868883604051611d9193929190614486565b60405180910390a15050600a805460ff60a01b1916600160a01b1790555050505050565b600a546001600160a01b031681565b6008546001600160a01b031681565b6000546001600160a01b03163314611dfd5760405162461bcd60e51b815260040161068c90614742565b6002544210611e1e5760405162461bcd60e51b815260040161068c9061469a565b61016e8110611e3f5760405162461bcd60e51b815260040161068c906146bb565b600280546201518083020190556040517f02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e890610df6908390614632565b600660209081526000928352604080842090915290825290205460ff1681565b6000546001600160a01b03163314611ec65760405162461bcd60e51b815260040161068c90614742565b6002544210611ee75760405162461bcd60e51b815260040161068c9061469a565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b600154600160a01b900460ff1615611f335760405162461bcd60e51b815260040161068c90614718565b600a54600160a01b900460ff16611f4a6001612eca565b90611f685760405162461bcd60e51b815260040161068c9190614659565b50600a805460ff60a01b191690556b033b2e3c9fd0803ce8000000471115611f906002612eca565b90611fae5760405162461bcd60e51b815260040161068c9190614659565b507f2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e60008234604051611fe393929190614486565b60405180910390a150600a805460ff60a01b1916600160a01b179055565b600a54600160a01b900460ff1681565b600154600160a01b900460ff161561203b5760405162461bcd60e51b815260040161068c90614718565b600a54600160a01b900460ff166120526001612eca565b906120705760405162461bcd60e51b815260040161068c9190614659565b50600a805460ff60a01b1916905585518751148015612090575060008551115b61209a6006612eca565b906120b85760405162461bcd60e51b815260040161068c9190614659565b5060006120d58585858a60405160200161189594939291906142cc565b90503460005b89518110156124be57612101838b83815181106120f457fe5b602002602001015161333b565b88818151811061210d57fe5b602002602001015160056000856001600160a01b03166001600160a01b0316815260200190815260200160002060008c848151811061214857fe5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054101561217d6008612eca565b9061219b5760405162461bcd60e51b815260040161068c9190614659565b506121fa8982815181106121ab57fe5b6020026020010151600760008d85815181106121c357fe5b60200260200101516001600160a01b03166001600160a01b03168152602001908152602001600020546132ba90919063ffffffff16565b600760008c848151811061220a57fe5b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000208190555061228089828151811061224557fe5b602002602001015160056000866001600160a01b03166001600160a01b0316815260200190815260200160002060008d85815181106121c357fe5b6001600160a01b03841660009081526005602052604081208c519091908d90859081106122a957fe5b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000208190555060006001600160a01b03168a82815181106122ec57fe5b60200260200101516001600160a01b031614156123305761232989828151811061231257fe5b60200260200101518361302790919063ffffffff16565b91506124b6565b88818151811061233c57fe5b60200260200101518a828151811061235057fe5b60200260200101516001600160a01b03166370a08231306040518263ffffffff1660e01b8152600401612383919061441b565b60206040518083038186803b15801561239b57600080fd5b505afa1580156123af573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123d391906141b3565b10156123df6007612eca565b906123fd5760405162461bcd60e51b815260040161068c9190614659565b5089818151811061240a57fe5b60200260200101516001600160a01b031663a9059cbb888b848151811061242d57fe5b60200260200101516040518363ffffffff1660e01b8152600401612452929190614453565b600060405180830381600087803b15801561246c57600080fd5b505af1158015612480573d6000803e3d6000fd5b5050505061248c612ff3565b6124966004612eca565b906124b45760405162461bcd60e51b815260040161068c9190614659565b505b6001016120db565b50606087516001600160401b03811180156124d857600080fd5b50604051908082528060200260200182016040528015612502578160200160208202803683370190505b50905060005b88518110156125ba5761252d89828151811061252057fe5b6020026020010151611598565b82828151811061253957fe5b60200260200101818152505060006001600160a01b031689828151811061255c57fe5b60200260200101516001600160a01b031614156125b2576125993483838151811061258357fe5b60200260200101516132ba90919063ffffffff16565b8282815181106125a557fe5b6020026020010181815250505b600101612508565b506060806125c984898b6135be565b91509150895182511480156125df575081518151145b6125e96009612eca565b906126075760405162461bcd60e51b815260040161068c9190614659565b5060005b81518110156127f4578a818151811061262057fe5b60200260200101516001600160a01b031683828151811061263d57fe5b60200260200101516001600160a01b0316148015612692575081818151811061266257fe5b602002602001015161269085838151811061267957fe5b60200260200101516114c28e858151811061252057fe5b145b61269c6009612eca565b906126ba5760405162461bcd60e51b815260040161068c9190614659565b5061273c8282815181106126ca57fe5b602002602001015160056000896001600160a01b03166001600160a01b0316815260200190815260200160002060008e858151811061270557fe5b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000205461302790919063ffffffff16565b6001600160a01b03871660009081526005602052604081208d519091908e908590811061276557fe5b60200260200101516001600160a01b03166001600160a01b03168152602001908152602001600020819055506127b88282815181106127a057fe5b6020026020010151600760008e858151811061270557fe5b600760008d84815181106127c857fe5b6020908102919091018101516001600160a01b031682528101919091526040016000205560010161260b565b5050600a805460ff60a01b1916600160a01b1790555050505050505050505050565b6000546001600160a01b031633146128405760405162461bcd60e51b815260040161068c90614742565b600154600160a01b900460ff166128695760405162461bcd60e51b815260040161068c9061466c565b6001600160a01b038116151561287f600b612eca565b9061289d5760405162461bcd60e51b815260040161068c9190614659565b50600a80546001600160a01b0319166001600160a01b0383161790556040517fd58a618a39de682696ea37dd9a6bf9c793afa426fa1438e75c3966e3b541e45a90610df690839061441b565b600080829050806001600160a01b031663313ce5676040518163ffffffff1660e01b815260040160206040518083038186803b15801561292857600080fd5b505afa15801561293c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061296091906141b3565b9392505050565b60036020526000908152604090205460ff1681565b60008181526004602052604081205460ff161561299b5750600161162e565b6009546001600160a01b03166129b35750600061162e565b60095460405163392f5c1d60e21b81526001600160a01b039091169063e4bd70749061169a908590600401614632565b6009546000906001600160a01b031615801590612a2657506001600160a01b0380831660009081526006602090815260408083209387168352929052205460ff16155b15612ae257600954604051637badcc6760e11b8152612adb916001600160a01b03169063f75b98ce90612a5f908790879060040161446c565b60206040518083038186803b158015612a7757600080fd5b505afa158015612a8b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612aaf91906141b3565b6001600160a01b0380851660009081526005602090815260408083209389168352929052205490613027565b90506111e0565b506001600160a01b0380821660009081526005602090815260408083209386168352929052205492915050565b6000546001600160a01b031681565b600154600160a01b900460ff1615612b485760405162461bcd60e51b815260040161068c90614718565b600a54600160a01b900460ff16612b5f6001612eca565b90612b7d5760405162461bcd60e51b815260040161068c9190614659565b50600a805460ff60a01b19169055600080808080612b9a8f6116f0565b945094509450945094508460ff166048148015612bba57508360ff166001145b612bc357600080fd5b6001600160a01b038316612c26576001600160a01b038316600090815260076020526040902054612bf5908290613027565b471015612c026007612eca565b90612c205760405162461bcd60e51b815260040161068c9190614659565b50612d1b565b6000612c31846128e9565b905060098160ff161115612c4e5760098160ff1603600a0a820291505b6001600160a01b038416600090815260076020526040902054612c72908390613027565b6040516370a0823160e01b81526001600160a01b038616906370a0823190612c9e90309060040161441b565b60206040518083038186803b158015612cb657600080fd5b505afa158015612cca573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612cee91906141b3565b1015612cfa6007612eca565b90612d185760405162461bcd60e51b815260040161068c9190614659565b50505b612d2d8f8f8f8f8f8f8f8f8f8f613045565b6001600160a01b038316612dca576000826001600160a01b031682604051612d5490614418565b60006040518083038185875af1925050503d8060008114612d91576040519150601f19603f3d011682016040523d82523d6000602084013e612d96565b606091505b5050905080612da56004612eca565b90612dc35760405162461bcd60e51b815260040161068c9190614659565b5050612e5c565b60405163a9059cbb60e01b81526001600160a01b0384169063a9059cbb90612df89085908590600401614453565b600060405180830381600087803b158015612e1257600080fd5b505af1158015612e26573d6000803e3d6000fd5b50505050612e32612ff3565b612e3c6004612eca565b90612e5a5760405162461bcd60e51b815260040161068c9190614659565b505b7f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb838383604051612e8f9392919061442f565b60405180910390a15050600a805460ff60a01b1916600160a01b17905550505050505050505050505050565b6009546001600160a01b031681565b6060600082600d811115612eda57fe5b60408051600a808252818301909252919250906060908260208201818036833701905050905060005b60ff841615612f51578151600a60ff959095168581049560018401939106916030830160f81b9185918110612f3457fe5b60200101906001600160f81b031916908160001a90535050612f03565b6060816001016001600160401b0381118015612f6c57600080fd5b506040519080825280601f01601f191660200182016040528015612f97576020820181803683370190505b50905060005b828111612fe8578381840381518110612fb257fe5b602001015160f81c60f81b828281518110612fc957fe5b60200101906001600160f81b031916908160001a905350600101612f9d565b509695505050505050565b6000803d801561300a57602081146130135761301f565b6001915061301f565b60206000803e60005191505b501515905090565b600082820183811080159061303c5750828110155b61296057600080fd5b895160208b012060008b8b8260200201516040516020016130679291906143b7565b60408051601f198184030181529190528051602090910120905060008c8c6001602002015160405160200161309d9291906143b7565b6040516020818303038152906040528051906020012090506130be83611633565b156130c96005612eca565b906130e75760405162461bcd60e51b815260040161068c9190614659565b506008548c518c518c518c518c518c518c518c518c51604051637b2e908b60e11b81526001600160a01b03909a169963f65d211699613139996001998f99929891979096929591949093600401614585565b60206040518083038186803b15801561315157600080fd5b505afa158015613165573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131899190613f1e565b6131936006612eca565b906131b15760405162461bcd60e51b815260040161068c9190614659565b506008546020808e01518d8201518d8301518d8401518d8501518d8601518d8701518d880151978d0151604051637b2e908b60e11b81526001600160a01b03909a169963f65d21169961321a996000998e99919890979196909591949093919291600401614585565b60206040518083038186803b15801561323257600080fd5b505afa158015613246573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061326a9190613f1e565b6132746006612eca565b906132925760405162461bcd60e51b815260040161068c9190614659565b5050506000908152600360205260409020805460ff1916600117905550505050505050505050565b6000828211156132c957600080fd5b50900390565b815160208301206000906132e28161297c565b156132ed6005612eca565b9061330b5760405162461bcd60e51b815260040161068c9190614659565b5060006133188483611157565b600092835260046020526040909220805460ff1916600117905550905092915050565b6009546001600160a01b03161580159061337b57506001600160a01b0380831660009081526006602090815260408083209385168352929052205460ff16155b1561347957600954604051637badcc6760e11b8152613430916001600160a01b03169063f75b98ce906133b4908590879060040161446c565b60206040518083038186803b1580156133cc57600080fd5b505afa1580156133e0573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061340491906141b3565b6001600160a01b0380851660009081526005602090815260408083209387168352929052205490613027565b6001600160a01b038084166000818152600560209081526040808320948716808452948252808320959095559181526006825283812092815291905220805460ff191660011790555b5050565b60008061348986611598565b90506001600160a01b0386166134a6576134a381346132ba565b90505b844710156134b46007612eca565b906134d25760405162461bcd60e51b815260040161068c9190614659565b5060006060846001600160a01b031687876040516134f0919061439b565b60006040518083038185875af1925050503d806000811461352d576040519150601f19603f3d011682016040523d82523d6000602084013e613532565b606091505b50915091508161354157600080fd5b600080828060200190518101906135589190613b79565b91509150896001600160a01b0316826001600160a01b0316148015613588575080613586866114c28d611598565b145b6135926009612eca565b906135b05760405162461bcd60e51b815260040161068c9190614659565b509998505050505050505050565b606080844710156135cf6007612eca565b906135ed5760405162461bcd60e51b815260040161068c9190614659565b5060006060846001600160a01b0316878760405161360b919061439b565b60006040518083038185875af1925050503d8060008114613648576040519150601f19603f3d011682016040523d82523d6000602084013e61364d565b606091505b50915091508161365d6004612eca565b9061367b5760405162461bcd60e51b815260040161068c9190614659565b50808060200190518101906136909190613d9a565b935093505050935093915050565b80356111e08161481c565b80516111e08161481c565b60008083601f8401126136c5578182fd5b5081356001600160401b038111156136db578182fd5b60208301915083602080830285010111156136f557600080fd5b9250929050565b600082601f83011261370c578081fd5b813561371f61371a826147cd565b6147a7565b81815291506020808301908481018184028601820187101561374057600080fd5b60005b848110156137685781356137568161481c565b84529282019290820190600101613743565b505050505092915050565b6000601f8381840112613784578182fd5b61378e60406147a7565b9150818360005b6002811015613822578135860187858201126137b057600080fd5b80356137be61371a826147cd565b80828252602080830192508085018c8283870288010111156137df57600080fd5b600095505b8486101561380a576137f68d82613aea565b8452600195909501949281019281016137e4565b50908752958601959490940193505050600101613795565b5050505092915050565b6000601f838184011261383d578182fd5b61384760406147a7565b9150818360005b60028110156138225781358601878582011261386957600080fd5b803561387761371a826147cd565b80828252602080830192508085018c82838702880101111561389857600080fd5b600095505b848610156138bb57803584526001959095019492810192810161389d565b5090875295860195949094019350505060010161384e565b600082601f8301126138e3578081fd5b6138ed60406147a7565b9050808260005b600281101561391f5761390a8683358701613a2e565b835260209283019291909101906001016138f4565b50505092915050565b6000601f8381840112613939578182fd5b61394360406147a7565b9150818360005b60028110156138225781358601878582011261396557600080fd5b803561397361371a826147cd565b80828252602080830192508085018c82838702880101111561399457600080fd5b600095505b848610156139c657803560ff811681146139b257600080fd5b845260019590950194928101928101613999565b5090875295860195949094019350505060010161394a565b600082601f8301126139ee578081fd5b6139f860406147a7565b9050808284604085011115613a0c57600080fd5b60005b600281101561391f578135835260209283019290910190600101613a0f565b600082601f830112613a3e578081fd5b8135613a4c61371a826147cd565b818152915060208083019084810181840286018201871015613a6d57600080fd5b60005b8481101561376857813584529282019290820190600101613a70565b600082601f830112613a9c578081fd5b8151613aaa61371a826147cd565b818152915060208083019084810181840286018201871015613acb57600080fd5b60005b8481101561376857815184529282019290820190600101613ace565b80356111e081614834565b600082601f830112613b05578081fd5b81356001600160401b03811115613b1a578182fd5b613b2d601f8201601f19166020016147a7565b9150808252836020828501011115613b4457600080fd5b8060208401602084013760009082016020015292915050565b600060208284031215613b6e578081fd5b81356129608161481c565b60008060408385031215613b8b578081fd5b8251613b968161481c565b6020939093015192949293505050565b60008060408385031215613bb8578182fd5b8235613bc38161481c565b91506020830135613bd38161481c565b809150509250929050565b600080600080600080600060e0888a031215613bf8578283fd5b8735613c038161481c565b9650602088013595506040880135613c1a8161481c565b9450613c298960608a0161369e565b935060808801356001600160401b0380821115613c44578485fd5b613c508b838c01613af5565b945060a08a0135915080821115613c65578384fd5b613c718b838c01613af5565b935060c08a0135915080821115613c86578283fd5b50613c938a828b01613af5565b91505092959891949750929550565b600080600060608486031215613cb6578081fd5b8335613cc18161481c565b92506020840135915060408401356001600160401b03811115613ce2578182fd5b613cee86828701613af5565b9150509250925092565b60008060008060408587031215613d0d578182fd5b84356001600160401b0380821115613d23578384fd5b613d2f888389016136b4565b90965094506020870135915080821115613d47578384fd5b50613d54878288016136b4565b95989497509550505050565b600060208284031215613d71578081fd5b81356001600160401b03811115613d86578182fd5b613d92848285016136fc565b949350505050565b60008060408385031215613dac578182fd5b82516001600160401b0380821115613dc2578384fd5b818501915085601f830112613dd5578384fd5b8151613de361371a826147cd565b80828252602080830192508086018a828387028901011115613e03578889fd5b8896505b84871015613e2d57613e198b826136a9565b845260019690960195928101928101613e07565b508801519096509350505080821115613e44578283fd5b50613e5185828601613a8c565b9150509250929050565b600080600080600080600060e0888a031215613e75578081fd5b87356001600160401b0380821115613e8b578283fd5b613e978b838c016136fc565b985060208a0135915080821115613eac578283fd5b613eb88b838c01613a2e565b975060408a0135915080821115613ecd578283fd5b613ed98b838c016136fc565b9650613ee88b60608c0161369e565b955060808a0135915080821115613efd578283fd5b613f098b838c01613af5565b945060a08a0135915080821115613c65578283fd5b600060208284031215613f2f578081fd5b815161296081614834565b600060208284031215613f4b578081fd5b5035919050565b600060208284031215613f63578081fd5b81356001600160401b03811115613f78578182fd5b613d9284828501613af5565b6000806000806000806000806000806101a08b8d031215613fa3578384fd5b8a356001600160401b0380821115613fb9578586fd5b613fc58e838f01613af5565b9b50613fd48e60208f016139de565b9a5060608d0135915080821115613fe9578586fd5b613ff58e838f0161382c565b995060808d013591508082111561400a578586fd5b6140168e838f01613773565b98506140258e60a08f016139de565b97506140348e60e08f016139de565b96506101208d013591508082111561404a578586fd5b6140568e838f016138d3565b95506101408d013591508082111561406c578485fd5b6140788e838f01613928565b94506101608d013591508082111561408e578384fd5b61409a8e838f0161382c565b93506101808d01359150808211156140b0578283fd5b506140bd8d828e0161382c565b9150509295989b9194979a5092959850565b600080604083850312156140e1578182fd5b82356001600160401b038111156140f6578283fd5b61410285828601613af5565b95602094909401359450505050565b600080600080600060a08688031215614128578283fd5b85356001600160401b038082111561413e578485fd5b61414a89838a01613af5565b96506020880135915061415c8261481c565b9094506040870135935060608701359080821115614178578283fd5b61418489838a01613af5565b93506080880135915080821115614199578283fd5b506141a688828901613af5565b9150509295509295909350565b6000602082840312156141c4578081fd5b5051919050565b6000815180845260208085019450808401835b838110156142035781516001600160a01b0316875295820195908201906001016141de565b509495945050505050565b6000815180845260208085019450808401835b83811015614203578151151587529582019590820190600101614221565b6000815180845260208085019450808401835b8381101561420357815187529582019590820190600101614252565b6000815180845260208085019450808401835b8381101561420357815160ff1687529582019590820190600101614281565b600081518084526142b88160208601602086016147ec565b601f01601f19169290920160200192915050565b60006bffffffffffffffffffffffff198660601b168252845160206142f78260148601838a016147ec565b85519184019161430d8160148501848a016147ec565b8551920160140191818601845b828110156143365781518552938301939083019060010161431a565b50929998505050505050505050565b60006bffffffffffffffffffffffff198660601b168252845161436f8160148501602089016147ec565b8451908301906143868160148401602089016147ec565b01601481019390935250506034019392505050565b600082516143ad8184602087016147ec565b9190910192915050565b600083516143c98184602088016147ec565b9190910191825250602001919050565b600085516143eb818460208a016147ec565b606086901b6bffffffffffffffffffffffff191690830190815284516143868160148401602089016147ec565b90565b6001600160a01b0391909116815260200190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b6001600160a01b03929092168252602082015260400190565b6001600160a01b0392831681529116602082015260400190565b6001600160a01b03841681526060602082018190526000906144aa908301856142a0565b9050826040830152949350505050565b6040808252810184905260008560608301825b878110156144fd57602083356144e28161481c565b6001600160a01b0316835292830192909101906001016144cd565b5083810360208501528481526001600160fb1b0385111561451c578283fd5b602085029150818660208301370160200190815295945050505050565b60006020825261296060208301846141cb565b60006040825261455f60408301856141cb565b8281036020840152614571818561423f565b95945050505050565b901515815260200190565b60006101608d151583528c60208401528b60408401528060608401526145ad8184018c61423f565b905082810360808401526145c1818b61420e565b90508860a08401528760c084015282810360e08401526145e1818861423f565b90508281036101008401526145f6818761426e565b905082810361012084015261460b818661423f565b9050828103610140840152614620818561423f565b9e9d5050505050505050505050505050565b90815260200190565b93845260ff9290921660208401526040830152606082015260800190565b60006020825261296060208301846142a0565b6020808252601490820152736e6f7420706175736564207269676874206e6f7760601b604082015260600190565b602080825260079082015266195e1c1a5c995960ca1b604082015260600190565b6020808252601a908201527f63616e6e6f7420657874656e6420666f7220746f6f206c6f6e67000000000000604082015260600190565b6020808252600c908201526b1d5b985d5d1a1bdc9a5e995960a21b604082015260600190565b60208082526010908201526f706175736564207269676874206e6f7760801b604082015260600190565b6020808252600990820152683737ba1030b236b4b760b91b604082015260600190565b60ff91909116815260200190565b60ff95861681529390941660208401526001600160a01b039182166040840152166060820152608081019190915260a00190565b6040518181016001600160401b03811182821017156147c557600080fd5b604052919050565b60006001600160401b038211156147e2578081fd5b5060209081020190565b60005b838110156148075781810151838201526020016147ef565b83811115614816576000848401525b50505050565b6001600160a01b038116811461483157600080fd5b50565b801515811461483157600080fdfea2646970667358221220b704232ff84355e3ee42ca061d0e5c8749c293de20baf079b79fd9e65532702764736f6c634300060c0033"

// DeployVault deploys a new Ethereum contract, binding an instance of Vault to it.
func DeployVault(auth *bind.TransactOpts, backend bind.ContractBackend, admin common.Address, incognitoProxyAddress common.Address, _prevVault common.Address) (common.Address, *types.Transaction, *Vault, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VaultBin), backend, admin, incognitoProxyAddress, _prevVault)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Vault{VaultCaller: VaultCaller{contract: contract}, VaultTransactor: VaultTransactor{contract: contract}, VaultFilterer: VaultFilterer{contract: contract}}, nil
}

// Vault is an auto generated Go binding around an Ethereum contract.
type Vault struct {
	VaultCaller     // Read-only binding to the contract
	VaultTransactor // Write-only binding to the contract
	VaultFilterer   // Log filterer for contract events
}

// VaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultSession struct {
	Contract     *Vault            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultCallerSession struct {
	Contract *VaultCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultTransactorSession struct {
	Contract     *VaultTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultRaw struct {
	Contract *Vault // Generic contract binding to access the raw methods on
}

// VaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultCallerRaw struct {
	Contract *VaultCaller // Generic read-only contract binding to access the raw methods on
}

// VaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultTransactorRaw struct {
	Contract *VaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVault creates a new instance of Vault, bound to a specific deployed contract.
func NewVault(address common.Address, backend bind.ContractBackend) (*Vault, error) {
	contract, err := bindVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vault{VaultCaller: VaultCaller{contract: contract}, VaultTransactor: VaultTransactor{contract: contract}, VaultFilterer: VaultFilterer{contract: contract}}, nil
}

// NewVaultCaller creates a new read-only instance of Vault, bound to a specific deployed contract.
func NewVaultCaller(address common.Address, caller bind.ContractCaller) (*VaultCaller, error) {
	contract, err := bindVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultCaller{contract: contract}, nil
}

// NewVaultTransactor creates a new write-only instance of Vault, bound to a specific deployed contract.
func NewVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultTransactor, error) {
	contract, err := bindVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultTransactor{contract: contract}, nil
}

// NewVaultFilterer creates a new log filterer instance of Vault, bound to a specific deployed contract.
func NewVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultFilterer, error) {
	contract, err := bindVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultFilterer{contract: contract}, nil
}

// bindVault binds a generic wrapper to an already deployed contract.
func bindVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.VaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transact(opts, method, params...)
}

// ETHTOKEN is a free data retrieval call binding the contract method 0x58bc8337.
//
// Solidity: function ETH_TOKEN() constant returns(address)
func (_Vault *VaultCaller) ETHTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "ETH_TOKEN")
	return *ret0, err
}

// ETHTOKEN is a free data retrieval call binding the contract method 0x58bc8337.
//
// Solidity: function ETH_TOKEN() constant returns(address)
func (_Vault *VaultSession) ETHTOKEN() (common.Address, error) {
	return _Vault.Contract.ETHTOKEN(&_Vault.CallOpts)
}

// ETHTOKEN is a free data retrieval call binding the contract method 0x58bc8337.
//
// Solidity: function ETH_TOKEN() constant returns(address)
func (_Vault *VaultCallerSession) ETHTOKEN() (common.Address, error) {
	return _Vault.Contract.ETHTOKEN(&_Vault.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Vault *VaultCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Vault *VaultSession) Admin() (common.Address, error) {
	return _Vault.Contract.Admin(&_Vault.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Vault *VaultCallerSession) Admin() (common.Address, error) {
	return _Vault.Contract.Admin(&_Vault.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address token) constant returns(uint256)
func (_Vault *VaultCaller) BalanceOf(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "balanceOf", token)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address token) constant returns(uint256)
func (_Vault *VaultSession) BalanceOf(token common.Address) (*big.Int, error) {
	return _Vault.Contract.BalanceOf(&_Vault.CallOpts, token)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address token) constant returns(uint256)
func (_Vault *VaultCallerSession) BalanceOf(token common.Address) (*big.Int, error) {
	return _Vault.Contract.BalanceOf(&_Vault.CallOpts, token)
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() constant returns(uint256)
func (_Vault *VaultCaller) Expire(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "expire")
	return *ret0, err
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() constant returns(uint256)
func (_Vault *VaultSession) Expire() (*big.Int, error) {
	return _Vault.Contract.Expire(&_Vault.CallOpts)
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() constant returns(uint256)
func (_Vault *VaultCallerSession) Expire() (*big.Int, error) {
	return _Vault.Contract.Expire(&_Vault.CallOpts)
}

// GetDecimals is a free data retrieval call binding the contract method 0xcf54aaa0.
//
// Solidity: function getDecimals(address token) constant returns(uint8)
func (_Vault *VaultCaller) GetDecimals(opts *bind.CallOpts, token common.Address) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "getDecimals", token)
	return *ret0, err
}

// GetDecimals is a free data retrieval call binding the contract method 0xcf54aaa0.
//
// Solidity: function getDecimals(address token) constant returns(uint8)
func (_Vault *VaultSession) GetDecimals(token common.Address) (uint8, error) {
	return _Vault.Contract.GetDecimals(&_Vault.CallOpts, token)
}

// GetDecimals is a free data retrieval call binding the contract method 0xcf54aaa0.
//
// Solidity: function getDecimals(address token) constant returns(uint8)
func (_Vault *VaultCallerSession) GetDecimals(token common.Address) (uint8, error) {
	return _Vault.Contract.GetDecimals(&_Vault.CallOpts, token)
}

// GetDepositedBalance is a free data retrieval call binding the contract method 0xf75b98ce.
//
// Solidity: function getDepositedBalance(address token, address owner) constant returns(uint256)
func (_Vault *VaultCaller) GetDepositedBalance(opts *bind.CallOpts, token common.Address, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "getDepositedBalance", token, owner)
	return *ret0, err
}

// GetDepositedBalance is a free data retrieval call binding the contract method 0xf75b98ce.
//
// Solidity: function getDepositedBalance(address token, address owner) constant returns(uint256)
func (_Vault *VaultSession) GetDepositedBalance(token common.Address, owner common.Address) (*big.Int, error) {
	return _Vault.Contract.GetDepositedBalance(&_Vault.CallOpts, token, owner)
}

// GetDepositedBalance is a free data retrieval call binding the contract method 0xf75b98ce.
//
// Solidity: function getDepositedBalance(address token, address owner) constant returns(uint256)
func (_Vault *VaultCallerSession) GetDepositedBalance(token common.Address, owner common.Address) (*big.Int, error) {
	return _Vault.Contract.GetDepositedBalance(&_Vault.CallOpts, token, owner)
}

// Incognito is a free data retrieval call binding the contract method 0x8a984538.
//
// Solidity: function incognito() constant returns(address)
func (_Vault *VaultCaller) Incognito(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "incognito")
	return *ret0, err
}

// Incognito is a free data retrieval call binding the contract method 0x8a984538.
//
// Solidity: function incognito() constant returns(address)
func (_Vault *VaultSession) Incognito() (common.Address, error) {
	return _Vault.Contract.Incognito(&_Vault.CallOpts)
}

// Incognito is a free data retrieval call binding the contract method 0x8a984538.
//
// Solidity: function incognito() constant returns(address)
func (_Vault *VaultCallerSession) Incognito() (common.Address, error) {
	return _Vault.Contract.Incognito(&_Vault.CallOpts)
}

// IsSigDataUsed is a free data retrieval call binding the contract method 0xe4bd7074.
//
// Solidity: function isSigDataUsed(bytes32 hash) constant returns(bool)
func (_Vault *VaultCaller) IsSigDataUsed(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "isSigDataUsed", hash)
	return *ret0, err
}

// IsSigDataUsed is a free data retrieval call binding the contract method 0xe4bd7074.
//
// Solidity: function isSigDataUsed(bytes32 hash) constant returns(bool)
func (_Vault *VaultSession) IsSigDataUsed(hash [32]byte) (bool, error) {
	return _Vault.Contract.IsSigDataUsed(&_Vault.CallOpts, hash)
}

// IsSigDataUsed is a free data retrieval call binding the contract method 0xe4bd7074.
//
// Solidity: function isSigDataUsed(bytes32 hash) constant returns(bool)
func (_Vault *VaultCallerSession) IsSigDataUsed(hash [32]byte) (bool, error) {
	return _Vault.Contract.IsSigDataUsed(&_Vault.CallOpts, hash)
}

// IsWithdrawed is a free data retrieval call binding the contract method 0x749c5f86.
//
// Solidity: function isWithdrawed(bytes32 hash) constant returns(bool)
func (_Vault *VaultCaller) IsWithdrawed(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "isWithdrawed", hash)
	return *ret0, err
}

// IsWithdrawed is a free data retrieval call binding the contract method 0x749c5f86.
//
// Solidity: function isWithdrawed(bytes32 hash) constant returns(bool)
func (_Vault *VaultSession) IsWithdrawed(hash [32]byte) (bool, error) {
	return _Vault.Contract.IsWithdrawed(&_Vault.CallOpts, hash)
}

// IsWithdrawed is a free data retrieval call binding the contract method 0x749c5f86.
//
// Solidity: function isWithdrawed(bytes32 hash) constant returns(bool)
func (_Vault *VaultCallerSession) IsWithdrawed(hash [32]byte) (bool, error) {
	return _Vault.Contract.IsWithdrawed(&_Vault.CallOpts, hash)
}

// Migration is a free data retrieval call binding the contract method 0x995fac11.
//
// Solidity: function migration(address , address ) constant returns(bool)
func (_Vault *VaultCaller) Migration(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "migration", arg0, arg1)
	return *ret0, err
}

// Migration is a free data retrieval call binding the contract method 0x995fac11.
//
// Solidity: function migration(address , address ) constant returns(bool)
func (_Vault *VaultSession) Migration(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Vault.Contract.Migration(&_Vault.CallOpts, arg0, arg1)
}

// Migration is a free data retrieval call binding the contract method 0x995fac11.
//
// Solidity: function migration(address , address ) constant returns(bool)
func (_Vault *VaultCallerSession) Migration(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Vault.Contract.Migration(&_Vault.CallOpts, arg0, arg1)
}

// NewVault is a free data retrieval call binding the contract method 0x88aaf0c8.
//
// Solidity: function newVault() constant returns(address)
func (_Vault *VaultCaller) NewVault(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "newVault")
	return *ret0, err
}

// NewVault is a free data retrieval call binding the contract method 0x88aaf0c8.
//
// Solidity: function newVault() constant returns(address)
func (_Vault *VaultSession) NewVault() (common.Address, error) {
	return _Vault.Contract.NewVault(&_Vault.CallOpts)
}

// NewVault is a free data retrieval call binding the contract method 0x88aaf0c8.
//
// Solidity: function newVault() constant returns(address)
func (_Vault *VaultCallerSession) NewVault() (common.Address, error) {
	return _Vault.Contract.NewVault(&_Vault.CallOpts)
}

// NotEntered is a free data retrieval call binding the contract method 0xa3f5d8cc.
//
// Solidity: function notEntered() constant returns(bool)
func (_Vault *VaultCaller) NotEntered(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "notEntered")
	return *ret0, err
}

// NotEntered is a free data retrieval call binding the contract method 0xa3f5d8cc.
//
// Solidity: function notEntered() constant returns(bool)
func (_Vault *VaultSession) NotEntered() (bool, error) {
	return _Vault.Contract.NotEntered(&_Vault.CallOpts)
}

// NotEntered is a free data retrieval call binding the contract method 0xa3f5d8cc.
//
// Solidity: function notEntered() constant returns(bool)
func (_Vault *VaultCallerSession) NotEntered() (bool, error) {
	return _Vault.Contract.NotEntered(&_Vault.CallOpts)
}

// ParseBurnInst is a free data retrieval call binding the contract method 0x7e16e6e1.
//
// Solidity: function parseBurnInst(bytes inst) constant returns(uint8, uint8, address, address, uint256)
func (_Vault *VaultCaller) ParseBurnInst(opts *bind.CallOpts, inst []byte) (uint8, uint8, common.Address, common.Address, *big.Int, error) {
	var (
		ret0 = new(uint8)
		ret1 = new(uint8)
		ret2 = new(common.Address)
		ret3 = new(common.Address)
		ret4 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _Vault.contract.Call(opts, out, "parseBurnInst", inst)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// ParseBurnInst is a free data retrieval call binding the contract method 0x7e16e6e1.
//
// Solidity: function parseBurnInst(bytes inst) constant returns(uint8, uint8, address, address, uint256)
func (_Vault *VaultSession) ParseBurnInst(inst []byte) (uint8, uint8, common.Address, common.Address, *big.Int, error) {
	return _Vault.Contract.ParseBurnInst(&_Vault.CallOpts, inst)
}

// ParseBurnInst is a free data retrieval call binding the contract method 0x7e16e6e1.
//
// Solidity: function parseBurnInst(bytes inst) constant returns(uint8, uint8, address, address, uint256)
func (_Vault *VaultCallerSession) ParseBurnInst(inst []byte) (uint8, uint8, common.Address, common.Address, *big.Int, error) {
	return _Vault.Contract.ParseBurnInst(&_Vault.CallOpts, inst)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Vault *VaultCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Vault *VaultSession) Paused() (bool, error) {
	return _Vault.Contract.Paused(&_Vault.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Vault *VaultCallerSession) Paused() (bool, error) {
	return _Vault.Contract.Paused(&_Vault.CallOpts)
}

// PrevVault is a free data retrieval call binding the contract method 0xfa84702e.
//
// Solidity: function prevVault() constant returns(address)
func (_Vault *VaultCaller) PrevVault(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "prevVault")
	return *ret0, err
}

// PrevVault is a free data retrieval call binding the contract method 0xfa84702e.
//
// Solidity: function prevVault() constant returns(address)
func (_Vault *VaultSession) PrevVault() (common.Address, error) {
	return _Vault.Contract.PrevVault(&_Vault.CallOpts)
}

// PrevVault is a free data retrieval call binding the contract method 0xfa84702e.
//
// Solidity: function prevVault() constant returns(address)
func (_Vault *VaultCallerSession) PrevVault() (common.Address, error) {
	return _Vault.Contract.PrevVault(&_Vault.CallOpts)
}

// SigDataUsed is a free data retrieval call binding the contract method 0x1ea1940e.
//
// Solidity: function sigDataUsed(bytes32 ) constant returns(bool)
func (_Vault *VaultCaller) SigDataUsed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "sigDataUsed", arg0)
	return *ret0, err
}

// SigDataUsed is a free data retrieval call binding the contract method 0x1ea1940e.
//
// Solidity: function sigDataUsed(bytes32 ) constant returns(bool)
func (_Vault *VaultSession) SigDataUsed(arg0 [32]byte) (bool, error) {
	return _Vault.Contract.SigDataUsed(&_Vault.CallOpts, arg0)
}

// SigDataUsed is a free data retrieval call binding the contract method 0x1ea1940e.
//
// Solidity: function sigDataUsed(bytes32 ) constant returns(bool)
func (_Vault *VaultCallerSession) SigDataUsed(arg0 [32]byte) (bool, error) {
	return _Vault.Contract.SigDataUsed(&_Vault.CallOpts, arg0)
}

// SigToAddress is a free data retrieval call binding the contract method 0x3fec6b40.
//
// Solidity: function sigToAddress(bytes signData, bytes32 hash) constant returns(address)
func (_Vault *VaultCaller) SigToAddress(opts *bind.CallOpts, signData []byte, hash [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "sigToAddress", signData, hash)
	return *ret0, err
}

// SigToAddress is a free data retrieval call binding the contract method 0x3fec6b40.
//
// Solidity: function sigToAddress(bytes signData, bytes32 hash) constant returns(address)
func (_Vault *VaultSession) SigToAddress(signData []byte, hash [32]byte) (common.Address, error) {
	return _Vault.Contract.SigToAddress(&_Vault.CallOpts, signData, hash)
}

// SigToAddress is a free data retrieval call binding the contract method 0x3fec6b40.
//
// Solidity: function sigToAddress(bytes signData, bytes32 hash) constant returns(address)
func (_Vault *VaultCallerSession) SigToAddress(signData []byte, hash [32]byte) (common.Address, error) {
	return _Vault.Contract.SigToAddress(&_Vault.CallOpts, signData, hash)
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() constant returns(address)
func (_Vault *VaultCaller) Successor(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "successor")
	return *ret0, err
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() constant returns(address)
func (_Vault *VaultSession) Successor() (common.Address, error) {
	return _Vault.Contract.Successor(&_Vault.CallOpts)
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() constant returns(address)
func (_Vault *VaultCallerSession) Successor() (common.Address, error) {
	return _Vault.Contract.Successor(&_Vault.CallOpts)
}

// TotalDepositedToSCAmount is a free data retrieval call binding the contract method 0x6304541c.
//
// Solidity: function totalDepositedToSCAmount(address ) constant returns(uint256)
func (_Vault *VaultCaller) TotalDepositedToSCAmount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "totalDepositedToSCAmount", arg0)
	return *ret0, err
}

// TotalDepositedToSCAmount is a free data retrieval call binding the contract method 0x6304541c.
//
// Solidity: function totalDepositedToSCAmount(address ) constant returns(uint256)
func (_Vault *VaultSession) TotalDepositedToSCAmount(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.TotalDepositedToSCAmount(&_Vault.CallOpts, arg0)
}

// TotalDepositedToSCAmount is a free data retrieval call binding the contract method 0x6304541c.
//
// Solidity: function totalDepositedToSCAmount(address ) constant returns(uint256)
func (_Vault *VaultCallerSession) TotalDepositedToSCAmount(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.TotalDepositedToSCAmount(&_Vault.CallOpts, arg0)
}

// WithdrawRequests is a free data retrieval call binding the contract method 0x65b5a00f.
//
// Solidity: function withdrawRequests(address , address ) constant returns(uint256)
func (_Vault *VaultCaller) WithdrawRequests(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "withdrawRequests", arg0, arg1)
	return *ret0, err
}

// WithdrawRequests is a free data retrieval call binding the contract method 0x65b5a00f.
//
// Solidity: function withdrawRequests(address , address ) constant returns(uint256)
func (_Vault *VaultSession) WithdrawRequests(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Vault.Contract.WithdrawRequests(&_Vault.CallOpts, arg0, arg1)
}

// WithdrawRequests is a free data retrieval call binding the contract method 0x65b5a00f.
//
// Solidity: function withdrawRequests(address , address ) constant returns(uint256)
func (_Vault *VaultCallerSession) WithdrawRequests(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Vault.Contract.WithdrawRequests(&_Vault.CallOpts, arg0, arg1)
}

// Withdrawed is a free data retrieval call binding the contract method 0xdca40d9e.
//
// Solidity: function withdrawed(bytes32 ) constant returns(bool)
func (_Vault *VaultCaller) Withdrawed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "withdrawed", arg0)
	return *ret0, err
}

// Withdrawed is a free data retrieval call binding the contract method 0xdca40d9e.
//
// Solidity: function withdrawed(bytes32 ) constant returns(bool)
func (_Vault *VaultSession) Withdrawed(arg0 [32]byte) (bool, error) {
	return _Vault.Contract.Withdrawed(&_Vault.CallOpts, arg0)
}

// Withdrawed is a free data retrieval call binding the contract method 0xdca40d9e.
//
// Solidity: function withdrawed(bytes32 ) constant returns(bool)
func (_Vault *VaultCallerSession) Withdrawed(arg0 [32]byte) (bool, error) {
	return _Vault.Contract.Withdrawed(&_Vault.CallOpts, arg0)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Vault *VaultTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Vault *VaultSession) Claim() (*types.Transaction, error) {
	return _Vault.Contract.Claim(&_Vault.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Vault *VaultTransactorSession) Claim() (*types.Transaction, error) {
	return _Vault.Contract.Claim(&_Vault.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xa26e1186.
//
// Solidity: function deposit(string incognitoAddress) returns()
func (_Vault *VaultTransactor) Deposit(opts *bind.TransactOpts, incognitoAddress string) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "deposit", incognitoAddress)
}

// Deposit is a paid mutator transaction binding the contract method 0xa26e1186.
//
// Solidity: function deposit(string incognitoAddress) returns()
func (_Vault *VaultSession) Deposit(incognitoAddress string) (*types.Transaction, error) {
	return _Vault.Contract.Deposit(&_Vault.TransactOpts, incognitoAddress)
}

// Deposit is a paid mutator transaction binding the contract method 0xa26e1186.
//
// Solidity: function deposit(string incognitoAddress) returns()
func (_Vault *VaultTransactorSession) Deposit(incognitoAddress string) (*types.Transaction, error) {
	return _Vault.Contract.Deposit(&_Vault.TransactOpts, incognitoAddress)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x5a67cb87.
//
// Solidity: function depositERC20(address token, uint256 amount, string incognitoAddress) returns()
func (_Vault *VaultTransactor) DepositERC20(opts *bind.TransactOpts, token common.Address, amount *big.Int, incognitoAddress string) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "depositERC20", token, amount, incognitoAddress)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x5a67cb87.
//
// Solidity: function depositERC20(address token, uint256 amount, string incognitoAddress) returns()
func (_Vault *VaultSession) DepositERC20(token common.Address, amount *big.Int, incognitoAddress string) (*types.Transaction, error) {
	return _Vault.Contract.DepositERC20(&_Vault.TransactOpts, token, amount, incognitoAddress)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x5a67cb87.
//
// Solidity: function depositERC20(address token, uint256 amount, string incognitoAddress) returns()
func (_Vault *VaultTransactorSession) DepositERC20(token common.Address, amount *big.Int, incognitoAddress string) (*types.Transaction, error) {
	return _Vault.Contract.DepositERC20(&_Vault.TransactOpts, token, amount, incognitoAddress)
}

// Execute is a paid mutator transaction binding the contract method 0x8588ccd6.
//
// Solidity: function execute(address token, uint256 amount, address recipientToken, address exchangeAddress, bytes callData, bytes timestamp, bytes signData) returns()
func (_Vault *VaultTransactor) Execute(opts *bind.TransactOpts, token common.Address, amount *big.Int, recipientToken common.Address, exchangeAddress common.Address, callData []byte, timestamp []byte, signData []byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "execute", token, amount, recipientToken, exchangeAddress, callData, timestamp, signData)
}

// Execute is a paid mutator transaction binding the contract method 0x8588ccd6.
//
// Solidity: function execute(address token, uint256 amount, address recipientToken, address exchangeAddress, bytes callData, bytes timestamp, bytes signData) returns()
func (_Vault *VaultSession) Execute(token common.Address, amount *big.Int, recipientToken common.Address, exchangeAddress common.Address, callData []byte, timestamp []byte, signData []byte) (*types.Transaction, error) {
	return _Vault.Contract.Execute(&_Vault.TransactOpts, token, amount, recipientToken, exchangeAddress, callData, timestamp, signData)
}

// Execute is a paid mutator transaction binding the contract method 0x8588ccd6.
//
// Solidity: function execute(address token, uint256 amount, address recipientToken, address exchangeAddress, bytes callData, bytes timestamp, bytes signData) returns()
func (_Vault *VaultTransactorSession) Execute(token common.Address, amount *big.Int, recipientToken common.Address, exchangeAddress common.Address, callData []byte, timestamp []byte, signData []byte) (*types.Transaction, error) {
	return _Vault.Contract.Execute(&_Vault.TransactOpts, token, amount, recipientToken, exchangeAddress, callData, timestamp, signData)
}

// ExecuteMulti is a paid mutator transaction binding the contract method 0xb69f6511.
//
// Solidity: function executeMulti(address[] tokens, uint256[] amounts, address[] recipientTokens, address exchangeAddress, bytes callData, bytes timestamp, bytes signData) returns()
func (_Vault *VaultTransactor) ExecuteMulti(opts *bind.TransactOpts, tokens []common.Address, amounts []*big.Int, recipientTokens []common.Address, exchangeAddress common.Address, callData []byte, timestamp []byte, signData []byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "executeMulti", tokens, amounts, recipientTokens, exchangeAddress, callData, timestamp, signData)
}

// ExecuteMulti is a paid mutator transaction binding the contract method 0xb69f6511.
//
// Solidity: function executeMulti(address[] tokens, uint256[] amounts, address[] recipientTokens, address exchangeAddress, bytes callData, bytes timestamp, bytes signData) returns()
func (_Vault *VaultSession) ExecuteMulti(tokens []common.Address, amounts []*big.Int, recipientTokens []common.Address, exchangeAddress common.Address, callData []byte, timestamp []byte, signData []byte) (*types.Transaction, error) {
	return _Vault.Contract.ExecuteMulti(&_Vault.TransactOpts, tokens, amounts, recipientTokens, exchangeAddress, callData, timestamp, signData)
}

// ExecuteMulti is a paid mutator transaction binding the contract method 0xb69f6511.
//
// Solidity: function executeMulti(address[] tokens, uint256[] amounts, address[] recipientTokens, address exchangeAddress, bytes callData, bytes timestamp, bytes signData) returns()
func (_Vault *VaultTransactorSession) ExecuteMulti(tokens []common.Address, amounts []*big.Int, recipientTokens []common.Address, exchangeAddress common.Address, callData []byte, timestamp []byte, signData []byte) (*types.Transaction, error) {
	return _Vault.Contract.ExecuteMulti(&_Vault.TransactOpts, tokens, amounts, recipientTokens, exchangeAddress, callData, timestamp, signData)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_Vault *VaultTransactor) Extend(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "extend", n)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_Vault *VaultSession) Extend(n *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Extend(&_Vault.TransactOpts, n)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_Vault *VaultTransactorSession) Extend(n *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Extend(&_Vault.TransactOpts, n)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _newVault) returns()
func (_Vault *VaultTransactor) Migrate(opts *bind.TransactOpts, _newVault common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "migrate", _newVault)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _newVault) returns()
func (_Vault *VaultSession) Migrate(_newVault common.Address) (*types.Transaction, error) {
	return _Vault.Contract.Migrate(&_Vault.TransactOpts, _newVault)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _newVault) returns()
func (_Vault *VaultTransactorSession) Migrate(_newVault common.Address) (*types.Transaction, error) {
	return _Vault.Contract.Migrate(&_Vault.TransactOpts, _newVault)
}

// MoveAssets is a paid mutator transaction binding the contract method 0x0c4f5039.
//
// Solidity: function moveAssets(address[] assets) returns()
func (_Vault *VaultTransactor) MoveAssets(opts *bind.TransactOpts, assets []common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "moveAssets", assets)
}

// MoveAssets is a paid mutator transaction binding the contract method 0x0c4f5039.
//
// Solidity: function moveAssets(address[] assets) returns()
func (_Vault *VaultSession) MoveAssets(assets []common.Address) (*types.Transaction, error) {
	return _Vault.Contract.MoveAssets(&_Vault.TransactOpts, assets)
}

// MoveAssets is a paid mutator transaction binding the contract method 0x0c4f5039.
//
// Solidity: function moveAssets(address[] assets) returns()
func (_Vault *VaultTransactorSession) MoveAssets(assets []common.Address) (*types.Transaction, error) {
	return _Vault.Contract.MoveAssets(&_Vault.TransactOpts, assets)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Vault *VaultTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Vault *VaultSession) Pause() (*types.Transaction, error) {
	return _Vault.Contract.Pause(&_Vault.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Vault *VaultTransactorSession) Pause() (*types.Transaction, error) {
	return _Vault.Contract.Pause(&_Vault.TransactOpts)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x87add440.
//
// Solidity: function requestWithdraw(string incognitoAddress, address token, uint256 amount, bytes signData, bytes timestamp) returns()
func (_Vault *VaultTransactor) RequestWithdraw(opts *bind.TransactOpts, incognitoAddress string, token common.Address, amount *big.Int, signData []byte, timestamp []byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "requestWithdraw", incognitoAddress, token, amount, signData, timestamp)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x87add440.
//
// Solidity: function requestWithdraw(string incognitoAddress, address token, uint256 amount, bytes signData, bytes timestamp) returns()
func (_Vault *VaultSession) RequestWithdraw(incognitoAddress string, token common.Address, amount *big.Int, signData []byte, timestamp []byte) (*types.Transaction, error) {
	return _Vault.Contract.RequestWithdraw(&_Vault.TransactOpts, incognitoAddress, token, amount, signData, timestamp)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x87add440.
//
// Solidity: function requestWithdraw(string incognitoAddress, address token, uint256 amount, bytes signData, bytes timestamp) returns()
func (_Vault *VaultTransactorSession) RequestWithdraw(incognitoAddress string, token common.Address, amount *big.Int, signData []byte, timestamp []byte) (*types.Transaction, error) {
	return _Vault.Contract.RequestWithdraw(&_Vault.TransactOpts, incognitoAddress, token, amount, signData, timestamp)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_Vault *VaultTransactor) Retire(opts *bind.TransactOpts, _successor common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "retire", _successor)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_Vault *VaultSession) Retire(_successor common.Address) (*types.Transaction, error) {
	return _Vault.Contract.Retire(&_Vault.TransactOpts, _successor)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_Vault *VaultTransactorSession) Retire(_successor common.Address) (*types.Transaction, error) {
	return _Vault.Contract.Retire(&_Vault.TransactOpts, _successor)
}

// SubmitBurnProof is a paid mutator transaction binding the contract method 0x3cb4b58a.
//
// Solidity: function submitBurnProof(bytes inst, uint256[2] heights, bytes32[][2] instPaths, bool[][2] instPathIsLefts, bytes32[2] instRoots, bytes32[2] blkData, uint256[][2] sigIdxs, uint8[][2] sigVs, bytes32[][2] sigRs, bytes32[][2] sigSs) returns()
func (_Vault *VaultTransactor) SubmitBurnProof(opts *bind.TransactOpts, inst []byte, heights [2]*big.Int, instPaths [2][][32]byte, instPathIsLefts [2][]bool, instRoots [2][32]byte, blkData [2][32]byte, sigIdxs [2][]*big.Int, sigVs [2][]uint8, sigRs [2][][32]byte, sigSs [2][][32]byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "submitBurnProof", inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// SubmitBurnProof is a paid mutator transaction binding the contract method 0x3cb4b58a.
//
// Solidity: function submitBurnProof(bytes inst, uint256[2] heights, bytes32[][2] instPaths, bool[][2] instPathIsLefts, bytes32[2] instRoots, bytes32[2] blkData, uint256[][2] sigIdxs, uint8[][2] sigVs, bytes32[][2] sigRs, bytes32[][2] sigSs) returns()
func (_Vault *VaultSession) SubmitBurnProof(inst []byte, heights [2]*big.Int, instPaths [2][][32]byte, instPathIsLefts [2][]bool, instRoots [2][32]byte, blkData [2][32]byte, sigIdxs [2][]*big.Int, sigVs [2][]uint8, sigRs [2][][32]byte, sigSs [2][][32]byte) (*types.Transaction, error) {
	return _Vault.Contract.SubmitBurnProof(&_Vault.TransactOpts, inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// SubmitBurnProof is a paid mutator transaction binding the contract method 0x3cb4b58a.
//
// Solidity: function submitBurnProof(bytes inst, uint256[2] heights, bytes32[][2] instPaths, bool[][2] instPathIsLefts, bytes32[2] instRoots, bytes32[2] blkData, uint256[][2] sigIdxs, uint8[][2] sigVs, bytes32[][2] sigRs, bytes32[][2] sigSs) returns()
func (_Vault *VaultTransactorSession) SubmitBurnProof(inst []byte, heights [2]*big.Int, instPaths [2][][32]byte, instPathIsLefts [2][]bool, instRoots [2][32]byte, blkData [2][32]byte, sigIdxs [2][]*big.Int, sigVs [2][]uint8, sigRs [2][][32]byte, sigSs [2][][32]byte) (*types.Transaction, error) {
	return _Vault.Contract.SubmitBurnProof(&_Vault.TransactOpts, inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Vault *VaultTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Vault *VaultSession) Unpause() (*types.Transaction, error) {
	return _Vault.Contract.Unpause(&_Vault.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Vault *VaultTransactorSession) Unpause() (*types.Transaction, error) {
	return _Vault.Contract.Unpause(&_Vault.TransactOpts)
}

// UpdateAssets is a paid mutator transaction binding the contract method 0x1ed4276d.
//
// Solidity: function updateAssets(address[] assets, uint256[] amounts) returns(bool)
func (_Vault *VaultTransactor) UpdateAssets(opts *bind.TransactOpts, assets []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "updateAssets", assets, amounts)
}

// UpdateAssets is a paid mutator transaction binding the contract method 0x1ed4276d.
//
// Solidity: function updateAssets(address[] assets, uint256[] amounts) returns(bool)
func (_Vault *VaultSession) UpdateAssets(assets []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Vault.Contract.UpdateAssets(&_Vault.TransactOpts, assets, amounts)
}

// UpdateAssets is a paid mutator transaction binding the contract method 0x1ed4276d.
//
// Solidity: function updateAssets(address[] assets, uint256[] amounts) returns(bool)
func (_Vault *VaultTransactorSession) UpdateAssets(assets []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Vault.Contract.UpdateAssets(&_Vault.TransactOpts, assets, amounts)
}

// UpdateIncognitoProxy is a paid mutator transaction binding the contract method 0x3a51913d.
//
// Solidity: function updateIncognitoProxy(address newIncognitoProxy) returns()
func (_Vault *VaultTransactor) UpdateIncognitoProxy(opts *bind.TransactOpts, newIncognitoProxy common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "updateIncognitoProxy", newIncognitoProxy)
}

// UpdateIncognitoProxy is a paid mutator transaction binding the contract method 0x3a51913d.
//
// Solidity: function updateIncognitoProxy(address newIncognitoProxy) returns()
func (_Vault *VaultSession) UpdateIncognitoProxy(newIncognitoProxy common.Address) (*types.Transaction, error) {
	return _Vault.Contract.UpdateIncognitoProxy(&_Vault.TransactOpts, newIncognitoProxy)
}

// UpdateIncognitoProxy is a paid mutator transaction binding the contract method 0x3a51913d.
//
// Solidity: function updateIncognitoProxy(address newIncognitoProxy) returns()
func (_Vault *VaultTransactorSession) UpdateIncognitoProxy(newIncognitoProxy common.Address) (*types.Transaction, error) {
	return _Vault.Contract.UpdateIncognitoProxy(&_Vault.TransactOpts, newIncognitoProxy)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf8b39745.
//
// Solidity: function withdraw(bytes inst, uint256[2] heights, bytes32[][2] instPaths, bool[][2] instPathIsLefts, bytes32[2] instRoots, bytes32[2] blkData, uint256[][2] sigIdxs, uint8[][2] sigVs, bytes32[][2] sigRs, bytes32[][2] sigSs) returns()
func (_Vault *VaultTransactor) Withdraw(opts *bind.TransactOpts, inst []byte, heights [2]*big.Int, instPaths [2][][32]byte, instPathIsLefts [2][]bool, instRoots [2][32]byte, blkData [2][32]byte, sigIdxs [2][]*big.Int, sigVs [2][]uint8, sigRs [2][][32]byte, sigSs [2][][32]byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "withdraw", inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf8b39745.
//
// Solidity: function withdraw(bytes inst, uint256[2] heights, bytes32[][2] instPaths, bool[][2] instPathIsLefts, bytes32[2] instRoots, bytes32[2] blkData, uint256[][2] sigIdxs, uint8[][2] sigVs, bytes32[][2] sigRs, bytes32[][2] sigSs) returns()
func (_Vault *VaultSession) Withdraw(inst []byte, heights [2]*big.Int, instPaths [2][][32]byte, instPathIsLefts [2][]bool, instRoots [2][32]byte, blkData [2][32]byte, sigIdxs [2][]*big.Int, sigVs [2][]uint8, sigRs [2][][32]byte, sigSs [2][][32]byte) (*types.Transaction, error) {
	return _Vault.Contract.Withdraw(&_Vault.TransactOpts, inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf8b39745.
//
// Solidity: function withdraw(bytes inst, uint256[2] heights, bytes32[][2] instPaths, bool[][2] instPathIsLefts, bytes32[2] instRoots, bytes32[2] blkData, uint256[][2] sigIdxs, uint8[][2] sigVs, bytes32[][2] sigRs, bytes32[][2] sigSs) returns()
func (_Vault *VaultTransactorSession) Withdraw(inst []byte, heights [2]*big.Int, instPaths [2][][32]byte, instPathIsLefts [2][]bool, instRoots [2][32]byte, blkData [2][32]byte, sigIdxs [2][]*big.Int, sigVs [2][]uint8, sigRs [2][][32]byte, sigSs [2][][32]byte) (*types.Transaction, error) {
	return _Vault.Contract.Withdraw(&_Vault.TransactOpts, inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// VaultClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the Vault contract.
type VaultClaimIterator struct {
	Event *VaultClaim // Event containing the contract specifics and raw log

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
func (it *VaultClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultClaim)
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
		it.Event = new(VaultClaim)
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
func (it *VaultClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultClaim represents a Claim event raised by the Vault contract.
type VaultClaim struct {
	Claimer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_Vault *VaultFilterer) FilterClaim(opts *bind.FilterOpts) (*VaultClaimIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return &VaultClaimIterator{contract: _Vault.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_Vault *VaultFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *VaultClaim) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultClaim)
				if err := _Vault.contract.UnpackLog(event, "Claim", log); err != nil {
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
func (_Vault *VaultFilterer) ParseClaim(log types.Log) (*VaultClaim, error) {
	event := new(VaultClaim)
	if err := _Vault.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VaultDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Vault contract.
type VaultDepositIterator struct {
	Event *VaultDeposit // Event containing the contract specifics and raw log

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
func (it *VaultDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultDeposit)
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
		it.Event = new(VaultDeposit)
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
func (it *VaultDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultDeposit represents a Deposit event raised by the Vault contract.
type VaultDeposit struct {
	Token            common.Address
	IncognitoAddress string
	Amount           *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e.
//
// Solidity: event Deposit(address token, string incognitoAddress, uint256 amount)
func (_Vault *VaultFilterer) FilterDeposit(opts *bind.FilterOpts) (*VaultDepositIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &VaultDepositIterator{contract: _Vault.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e.
//
// Solidity: event Deposit(address token, string incognitoAddress, uint256 amount)
func (_Vault *VaultFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *VaultDeposit) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultDeposit)
				if err := _Vault.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e.
//
// Solidity: event Deposit(address token, string incognitoAddress, uint256 amount)
func (_Vault *VaultFilterer) ParseDeposit(log types.Log) (*VaultDeposit, error) {
	event := new(VaultDeposit)
	if err := _Vault.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VaultExtendIterator is returned from FilterExtend and is used to iterate over the raw logs and unpacked data for Extend events raised by the Vault contract.
type VaultExtendIterator struct {
	Event *VaultExtend // Event containing the contract specifics and raw log

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
func (it *VaultExtendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultExtend)
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
		it.Event = new(VaultExtend)
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
func (it *VaultExtendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultExtendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultExtend represents a Extend event raised by the Vault contract.
type VaultExtend struct {
	Ndays *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterExtend is a free log retrieval operation binding the contract event 0x02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8.
//
// Solidity: event Extend(uint256 ndays)
func (_Vault *VaultFilterer) FilterExtend(opts *bind.FilterOpts) (*VaultExtendIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Extend")
	if err != nil {
		return nil, err
	}
	return &VaultExtendIterator{contract: _Vault.contract, event: "Extend", logs: logs, sub: sub}, nil
}

// WatchExtend is a free log subscription operation binding the contract event 0x02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8.
//
// Solidity: event Extend(uint256 ndays)
func (_Vault *VaultFilterer) WatchExtend(opts *bind.WatchOpts, sink chan<- *VaultExtend) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Extend")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultExtend)
				if err := _Vault.contract.UnpackLog(event, "Extend", log); err != nil {
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
func (_Vault *VaultFilterer) ParseExtend(log types.Log) (*VaultExtend, error) {
	event := new(VaultExtend)
	if err := _Vault.contract.UnpackLog(event, "Extend", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VaultMigrateIterator is returned from FilterMigrate and is used to iterate over the raw logs and unpacked data for Migrate events raised by the Vault contract.
type VaultMigrateIterator struct {
	Event *VaultMigrate // Event containing the contract specifics and raw log

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
func (it *VaultMigrateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultMigrate)
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
		it.Event = new(VaultMigrate)
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
func (it *VaultMigrateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultMigrateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultMigrate represents a Migrate event raised by the Vault contract.
type VaultMigrate struct {
	NewVault common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMigrate is a free log retrieval operation binding the contract event 0xd58a618a39de682696ea37dd9a6bf9c793afa426fa1438e75c3966e3b541e45a.
//
// Solidity: event Migrate(address newVault)
func (_Vault *VaultFilterer) FilterMigrate(opts *bind.FilterOpts) (*VaultMigrateIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Migrate")
	if err != nil {
		return nil, err
	}
	return &VaultMigrateIterator{contract: _Vault.contract, event: "Migrate", logs: logs, sub: sub}, nil
}

// WatchMigrate is a free log subscription operation binding the contract event 0xd58a618a39de682696ea37dd9a6bf9c793afa426fa1438e75c3966e3b541e45a.
//
// Solidity: event Migrate(address newVault)
func (_Vault *VaultFilterer) WatchMigrate(opts *bind.WatchOpts, sink chan<- *VaultMigrate) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Migrate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultMigrate)
				if err := _Vault.contract.UnpackLog(event, "Migrate", log); err != nil {
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

// ParseMigrate is a log parse operation binding the contract event 0xd58a618a39de682696ea37dd9a6bf9c793afa426fa1438e75c3966e3b541e45a.
//
// Solidity: event Migrate(address newVault)
func (_Vault *VaultFilterer) ParseMigrate(log types.Log) (*VaultMigrate, error) {
	event := new(VaultMigrate)
	if err := _Vault.contract.UnpackLog(event, "Migrate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VaultMoveAssetsIterator is returned from FilterMoveAssets and is used to iterate over the raw logs and unpacked data for MoveAssets events raised by the Vault contract.
type VaultMoveAssetsIterator struct {
	Event *VaultMoveAssets // Event containing the contract specifics and raw log

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
func (it *VaultMoveAssetsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultMoveAssets)
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
		it.Event = new(VaultMoveAssets)
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
func (it *VaultMoveAssetsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultMoveAssetsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultMoveAssets represents a MoveAssets event raised by the Vault contract.
type VaultMoveAssets struct {
	Assets []common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMoveAssets is a free log retrieval operation binding the contract event 0x492fc8b292f2a2a9b328a366b83745f30c024056d12aa118a15966d26a8ce658.
//
// Solidity: event MoveAssets(address[] assets)
func (_Vault *VaultFilterer) FilterMoveAssets(opts *bind.FilterOpts) (*VaultMoveAssetsIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "MoveAssets")
	if err != nil {
		return nil, err
	}
	return &VaultMoveAssetsIterator{contract: _Vault.contract, event: "MoveAssets", logs: logs, sub: sub}, nil
}

// WatchMoveAssets is a free log subscription operation binding the contract event 0x492fc8b292f2a2a9b328a366b83745f30c024056d12aa118a15966d26a8ce658.
//
// Solidity: event MoveAssets(address[] assets)
func (_Vault *VaultFilterer) WatchMoveAssets(opts *bind.WatchOpts, sink chan<- *VaultMoveAssets) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "MoveAssets")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultMoveAssets)
				if err := _Vault.contract.UnpackLog(event, "MoveAssets", log); err != nil {
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

// ParseMoveAssets is a log parse operation binding the contract event 0x492fc8b292f2a2a9b328a366b83745f30c024056d12aa118a15966d26a8ce658.
//
// Solidity: event MoveAssets(address[] assets)
func (_Vault *VaultFilterer) ParseMoveAssets(log types.Log) (*VaultMoveAssets, error) {
	event := new(VaultMoveAssets)
	if err := _Vault.contract.UnpackLog(event, "MoveAssets", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VaultPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Vault contract.
type VaultPausedIterator struct {
	Event *VaultPaused // Event containing the contract specifics and raw log

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
func (it *VaultPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultPaused)
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
		it.Event = new(VaultPaused)
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
func (it *VaultPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultPaused represents a Paused event raised by the Vault contract.
type VaultPaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address pauser)
func (_Vault *VaultFilterer) FilterPaused(opts *bind.FilterOpts) (*VaultPausedIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &VaultPausedIterator{contract: _Vault.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address pauser)
func (_Vault *VaultFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *VaultPaused) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultPaused)
				if err := _Vault.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Vault *VaultFilterer) ParsePaused(log types.Log) (*VaultPaused, error) {
	event := new(VaultPaused)
	if err := _Vault.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VaultUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Vault contract.
type VaultUnpausedIterator struct {
	Event *VaultUnpaused // Event containing the contract specifics and raw log

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
func (it *VaultUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultUnpaused)
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
		it.Event = new(VaultUnpaused)
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
func (it *VaultUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultUnpaused represents a Unpaused event raised by the Vault contract.
type VaultUnpaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address pauser)
func (_Vault *VaultFilterer) FilterUnpaused(opts *bind.FilterOpts) (*VaultUnpausedIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &VaultUnpausedIterator{contract: _Vault.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address pauser)
func (_Vault *VaultFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *VaultUnpaused) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultUnpaused)
				if err := _Vault.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Vault *VaultFilterer) ParseUnpaused(log types.Log) (*VaultUnpaused, error) {
	event := new(VaultUnpaused)
	if err := _Vault.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VaultUpdateIncognitoProxyIterator is returned from FilterUpdateIncognitoProxy and is used to iterate over the raw logs and unpacked data for UpdateIncognitoProxy events raised by the Vault contract.
type VaultUpdateIncognitoProxyIterator struct {
	Event *VaultUpdateIncognitoProxy // Event containing the contract specifics and raw log

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
func (it *VaultUpdateIncognitoProxyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultUpdateIncognitoProxy)
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
		it.Event = new(VaultUpdateIncognitoProxy)
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
func (it *VaultUpdateIncognitoProxyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultUpdateIncognitoProxyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultUpdateIncognitoProxy represents a UpdateIncognitoProxy event raised by the Vault contract.
type VaultUpdateIncognitoProxy struct {
	NewIncognitoProxy common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUpdateIncognitoProxy is a free log retrieval operation binding the contract event 0x204252dfe190ad6ef63db40a490f048b39f661de74628408f13cd0bb2d4c3446.
//
// Solidity: event UpdateIncognitoProxy(address newIncognitoProxy)
func (_Vault *VaultFilterer) FilterUpdateIncognitoProxy(opts *bind.FilterOpts) (*VaultUpdateIncognitoProxyIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "UpdateIncognitoProxy")
	if err != nil {
		return nil, err
	}
	return &VaultUpdateIncognitoProxyIterator{contract: _Vault.contract, event: "UpdateIncognitoProxy", logs: logs, sub: sub}, nil
}

// WatchUpdateIncognitoProxy is a free log subscription operation binding the contract event 0x204252dfe190ad6ef63db40a490f048b39f661de74628408f13cd0bb2d4c3446.
//
// Solidity: event UpdateIncognitoProxy(address newIncognitoProxy)
func (_Vault *VaultFilterer) WatchUpdateIncognitoProxy(opts *bind.WatchOpts, sink chan<- *VaultUpdateIncognitoProxy) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "UpdateIncognitoProxy")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultUpdateIncognitoProxy)
				if err := _Vault.contract.UnpackLog(event, "UpdateIncognitoProxy", log); err != nil {
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

// ParseUpdateIncognitoProxy is a log parse operation binding the contract event 0x204252dfe190ad6ef63db40a490f048b39f661de74628408f13cd0bb2d4c3446.
//
// Solidity: event UpdateIncognitoProxy(address newIncognitoProxy)
func (_Vault *VaultFilterer) ParseUpdateIncognitoProxy(log types.Log) (*VaultUpdateIncognitoProxy, error) {
	event := new(VaultUpdateIncognitoProxy)
	if err := _Vault.contract.UnpackLog(event, "UpdateIncognitoProxy", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VaultUpdateTokenTotalIterator is returned from FilterUpdateTokenTotal and is used to iterate over the raw logs and unpacked data for UpdateTokenTotal events raised by the Vault contract.
type VaultUpdateTokenTotalIterator struct {
	Event *VaultUpdateTokenTotal // Event containing the contract specifics and raw log

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
func (it *VaultUpdateTokenTotalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultUpdateTokenTotal)
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
		it.Event = new(VaultUpdateTokenTotal)
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
func (it *VaultUpdateTokenTotalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultUpdateTokenTotalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultUpdateTokenTotal represents a UpdateTokenTotal event raised by the Vault contract.
type VaultUpdateTokenTotal struct {
	Assets  []common.Address
	Amounts []*big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateTokenTotal is a free log retrieval operation binding the contract event 0x6a7fbbcddfd518bb8c56b28ac6c7acb0f7ca093ed232eb3306e53d14e469895f.
//
// Solidity: event UpdateTokenTotal(address[] assets, uint256[] amounts)
func (_Vault *VaultFilterer) FilterUpdateTokenTotal(opts *bind.FilterOpts) (*VaultUpdateTokenTotalIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "UpdateTokenTotal")
	if err != nil {
		return nil, err
	}
	return &VaultUpdateTokenTotalIterator{contract: _Vault.contract, event: "UpdateTokenTotal", logs: logs, sub: sub}, nil
}

// WatchUpdateTokenTotal is a free log subscription operation binding the contract event 0x6a7fbbcddfd518bb8c56b28ac6c7acb0f7ca093ed232eb3306e53d14e469895f.
//
// Solidity: event UpdateTokenTotal(address[] assets, uint256[] amounts)
func (_Vault *VaultFilterer) WatchUpdateTokenTotal(opts *bind.WatchOpts, sink chan<- *VaultUpdateTokenTotal) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "UpdateTokenTotal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultUpdateTokenTotal)
				if err := _Vault.contract.UnpackLog(event, "UpdateTokenTotal", log); err != nil {
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

// ParseUpdateTokenTotal is a log parse operation binding the contract event 0x6a7fbbcddfd518bb8c56b28ac6c7acb0f7ca093ed232eb3306e53d14e469895f.
//
// Solidity: event UpdateTokenTotal(address[] assets, uint256[] amounts)
func (_Vault *VaultFilterer) ParseUpdateTokenTotal(log types.Log) (*VaultUpdateTokenTotal, error) {
	event := new(VaultUpdateTokenTotal)
	if err := _Vault.contract.UnpackLog(event, "UpdateTokenTotal", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VaultWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Vault contract.
type VaultWithdrawIterator struct {
	Event *VaultWithdraw // Event containing the contract specifics and raw log

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
func (it *VaultWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultWithdraw)
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
		it.Event = new(VaultWithdraw)
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
func (it *VaultWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultWithdraw represents a Withdraw event raised by the Vault contract.
type VaultWithdraw struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address token, address to, uint256 amount)
func (_Vault *VaultFilterer) FilterWithdraw(opts *bind.FilterOpts) (*VaultWithdrawIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &VaultWithdrawIterator{contract: _Vault.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address token, address to, uint256 amount)
func (_Vault *VaultFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *VaultWithdraw) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultWithdraw)
				if err := _Vault.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address token, address to, uint256 amount)
func (_Vault *VaultFilterer) ParseWithdraw(log types.Log) (*VaultWithdraw, error) {
	event := new(VaultWithdraw)
	if err := _Vault.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WithdrawableABI is the input ABI used to generate the binding from.
const WithdrawableABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getDepositedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isSigDataUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isWithdrawed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"name\":\"updateAssets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// WithdrawableFuncSigs maps the 4-byte function signature to its string representation.
var WithdrawableFuncSigs = map[string]string{
	"f75b98ce": "getDepositedBalance(address,address)",
	"e4bd7074": "isSigDataUsed(bytes32)",
	"749c5f86": "isWithdrawed(bytes32)",
	"5c975abb": "paused()",
	"1ed4276d": "updateAssets(address[],uint256[])",
}

// Withdrawable is an auto generated Go binding around an Ethereum contract.
type Withdrawable struct {
	WithdrawableCaller     // Read-only binding to the contract
	WithdrawableTransactor // Write-only binding to the contract
	WithdrawableFilterer   // Log filterer for contract events
}

// WithdrawableCaller is an auto generated read-only Go binding around an Ethereum contract.
type WithdrawableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WithdrawableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WithdrawableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WithdrawableSession struct {
	Contract     *Withdrawable     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WithdrawableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WithdrawableCallerSession struct {
	Contract *WithdrawableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// WithdrawableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WithdrawableTransactorSession struct {
	Contract     *WithdrawableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// WithdrawableRaw is an auto generated low-level Go binding around an Ethereum contract.
type WithdrawableRaw struct {
	Contract *Withdrawable // Generic contract binding to access the raw methods on
}

// WithdrawableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WithdrawableCallerRaw struct {
	Contract *WithdrawableCaller // Generic read-only contract binding to access the raw methods on
}

// WithdrawableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WithdrawableTransactorRaw struct {
	Contract *WithdrawableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWithdrawable creates a new instance of Withdrawable, bound to a specific deployed contract.
func NewWithdrawable(address common.Address, backend bind.ContractBackend) (*Withdrawable, error) {
	contract, err := bindWithdrawable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Withdrawable{WithdrawableCaller: WithdrawableCaller{contract: contract}, WithdrawableTransactor: WithdrawableTransactor{contract: contract}, WithdrawableFilterer: WithdrawableFilterer{contract: contract}}, nil
}

// NewWithdrawableCaller creates a new read-only instance of Withdrawable, bound to a specific deployed contract.
func NewWithdrawableCaller(address common.Address, caller bind.ContractCaller) (*WithdrawableCaller, error) {
	contract, err := bindWithdrawable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WithdrawableCaller{contract: contract}, nil
}

// NewWithdrawableTransactor creates a new write-only instance of Withdrawable, bound to a specific deployed contract.
func NewWithdrawableTransactor(address common.Address, transactor bind.ContractTransactor) (*WithdrawableTransactor, error) {
	contract, err := bindWithdrawable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WithdrawableTransactor{contract: contract}, nil
}

// NewWithdrawableFilterer creates a new log filterer instance of Withdrawable, bound to a specific deployed contract.
func NewWithdrawableFilterer(address common.Address, filterer bind.ContractFilterer) (*WithdrawableFilterer, error) {
	contract, err := bindWithdrawable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WithdrawableFilterer{contract: contract}, nil
}

// bindWithdrawable binds a generic wrapper to an already deployed contract.
func bindWithdrawable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WithdrawableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Withdrawable *WithdrawableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Withdrawable.Contract.WithdrawableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Withdrawable *WithdrawableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Withdrawable.Contract.WithdrawableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Withdrawable *WithdrawableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Withdrawable.Contract.WithdrawableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Withdrawable *WithdrawableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Withdrawable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Withdrawable *WithdrawableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Withdrawable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Withdrawable *WithdrawableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Withdrawable.Contract.contract.Transact(opts, method, params...)
}

// GetDepositedBalance is a free data retrieval call binding the contract method 0xf75b98ce.
//
// Solidity: function getDepositedBalance(address , address ) constant returns(uint256)
func (_Withdrawable *WithdrawableCaller) GetDepositedBalance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Withdrawable.contract.Call(opts, out, "getDepositedBalance", arg0, arg1)
	return *ret0, err
}

// GetDepositedBalance is a free data retrieval call binding the contract method 0xf75b98ce.
//
// Solidity: function getDepositedBalance(address , address ) constant returns(uint256)
func (_Withdrawable *WithdrawableSession) GetDepositedBalance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Withdrawable.Contract.GetDepositedBalance(&_Withdrawable.CallOpts, arg0, arg1)
}

// GetDepositedBalance is a free data retrieval call binding the contract method 0xf75b98ce.
//
// Solidity: function getDepositedBalance(address , address ) constant returns(uint256)
func (_Withdrawable *WithdrawableCallerSession) GetDepositedBalance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Withdrawable.Contract.GetDepositedBalance(&_Withdrawable.CallOpts, arg0, arg1)
}

// IsSigDataUsed is a free data retrieval call binding the contract method 0xe4bd7074.
//
// Solidity: function isSigDataUsed(bytes32 ) constant returns(bool)
func (_Withdrawable *WithdrawableCaller) IsSigDataUsed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Withdrawable.contract.Call(opts, out, "isSigDataUsed", arg0)
	return *ret0, err
}

// IsSigDataUsed is a free data retrieval call binding the contract method 0xe4bd7074.
//
// Solidity: function isSigDataUsed(bytes32 ) constant returns(bool)
func (_Withdrawable *WithdrawableSession) IsSigDataUsed(arg0 [32]byte) (bool, error) {
	return _Withdrawable.Contract.IsSigDataUsed(&_Withdrawable.CallOpts, arg0)
}

// IsSigDataUsed is a free data retrieval call binding the contract method 0xe4bd7074.
//
// Solidity: function isSigDataUsed(bytes32 ) constant returns(bool)
func (_Withdrawable *WithdrawableCallerSession) IsSigDataUsed(arg0 [32]byte) (bool, error) {
	return _Withdrawable.Contract.IsSigDataUsed(&_Withdrawable.CallOpts, arg0)
}

// IsWithdrawed is a free data retrieval call binding the contract method 0x749c5f86.
//
// Solidity: function isWithdrawed(bytes32 ) constant returns(bool)
func (_Withdrawable *WithdrawableCaller) IsWithdrawed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Withdrawable.contract.Call(opts, out, "isWithdrawed", arg0)
	return *ret0, err
}

// IsWithdrawed is a free data retrieval call binding the contract method 0x749c5f86.
//
// Solidity: function isWithdrawed(bytes32 ) constant returns(bool)
func (_Withdrawable *WithdrawableSession) IsWithdrawed(arg0 [32]byte) (bool, error) {
	return _Withdrawable.Contract.IsWithdrawed(&_Withdrawable.CallOpts, arg0)
}

// IsWithdrawed is a free data retrieval call binding the contract method 0x749c5f86.
//
// Solidity: function isWithdrawed(bytes32 ) constant returns(bool)
func (_Withdrawable *WithdrawableCallerSession) IsWithdrawed(arg0 [32]byte) (bool, error) {
	return _Withdrawable.Contract.IsWithdrawed(&_Withdrawable.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Withdrawable *WithdrawableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Withdrawable.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Withdrawable *WithdrawableSession) Paused() (bool, error) {
	return _Withdrawable.Contract.Paused(&_Withdrawable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Withdrawable *WithdrawableCallerSession) Paused() (bool, error) {
	return _Withdrawable.Contract.Paused(&_Withdrawable.CallOpts)
}

// UpdateAssets is a paid mutator transaction binding the contract method 0x1ed4276d.
//
// Solidity: function updateAssets(address[] , uint256[] ) returns(bool)
func (_Withdrawable *WithdrawableTransactor) UpdateAssets(opts *bind.TransactOpts, arg0 []common.Address, arg1 []*big.Int) (*types.Transaction, error) {
	return _Withdrawable.contract.Transact(opts, "updateAssets", arg0, arg1)
}

// UpdateAssets is a paid mutator transaction binding the contract method 0x1ed4276d.
//
// Solidity: function updateAssets(address[] , uint256[] ) returns(bool)
func (_Withdrawable *WithdrawableSession) UpdateAssets(arg0 []common.Address, arg1 []*big.Int) (*types.Transaction, error) {
	return _Withdrawable.Contract.UpdateAssets(&_Withdrawable.TransactOpts, arg0, arg1)
}

// UpdateAssets is a paid mutator transaction binding the contract method 0x1ed4276d.
//
// Solidity: function updateAssets(address[] , uint256[] ) returns(bool)
func (_Withdrawable *WithdrawableTransactorSession) UpdateAssets(arg0 []common.Address, arg1 []*big.Int) (*types.Transaction, error) {
	return _Withdrawable.Contract.UpdateAssets(&_Withdrawable.TransactOpts, arg0, arg1)
}
