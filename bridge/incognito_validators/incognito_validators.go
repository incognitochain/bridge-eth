// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package incognito_validators

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

// IncognitoValidatorsABI is the input ABI used to generate the binding from.
const IncognitoValidatorsABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"beaconCommittee\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"bridgeCommittee\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ndays\",\"type\":\"uint256\"}],\"name\":\"Extend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"changeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"committee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"oldGID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newGID\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expire\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"extend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"groupAddrs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isBeacon\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"swapID\",\"type\":\"uint256\"}],\"name\":\"loadCandidates\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isBeacon\",\"type\":\"bool\"}],\"name\":\"loadCommittee\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isBeacon\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"oldGID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newGID\",\"type\":\"uint256\"}],\"name\":\"loadGroups\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_successor\",\"type\":\"address\"}],\"name\":\"retire\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isBeacon\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"swapID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldGID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newGID\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"newAddrs\",\"type\":\"address[]\"}],\"name\":\"storeCandidates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isBeacon\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"gID1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gID2\",\"type\":\"uint256\"}],\"name\":\"storeCommittee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"successor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"oldGID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newGID\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IncognitoValidatorsFuncSigs maps the 4-byte function signature to its string representation.
var IncognitoValidatorsFuncSigs = map[string]string{
	"f851a440": "admin()",
	"06394c9b": "changeOperator(address)",
	"4e71d92d": "claim()",
	"19b52009": "committee(bool)",
	"79599f96": "expire()",
	"9714378c": "extend(uint256)",
	"e4712f36": "groupAddrs(bool,uint256,uint256)",
	"8126af55": "loadCandidates(bool,uint256)",
	"cbffbe41": "loadCommittee(bool)",
	"f3810d15": "loadGroups(bool,uint256,uint256)",
	"570ca735": "operator()",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"9e6371ba": "retire(address)",
	"48f2503d": "storeCandidates(bool,uint256,uint256,uint256,address[])",
	"e1b357e3": "storeCommittee(bool,uint256,uint256)",
	"6ff968c3": "successor()",
	"3f4ba83a": "unpause()",
	"e09366a7": "validators(bool,uint256)",
}

// IncognitoValidatorsBin is the compiled bytecode used for deploying new contracts.
var IncognitoValidatorsBin = "0x60806040523480156200001157600080fd5b506040516200112e3803806200112e83398101604081905262000034916200013f565b5050600080546001600160a01b039384166001600160a01b0319918216179091556001805460ff60a01b191690556301e13380420160025560038054929093169116179055620001ce565b80516001600160a01b03811681146200009757600080fd5b92915050565b600082601f830112620000ae578081fd5b81516001600160401b0380821115620000c5578283fd5b602080830260405182828201018181108582111715620000e3578687fd5b6040528481529450818501925085820181870183018810156200010557600080fd5b600091505b8482101562000134576200011f88826200007f565b8452928201926001919091019082016200010a565b505050505092915050565b6000806000806080858703121562000155578384fd5b6200016186866200007f565b93506200017286602087016200007f565b60408601519093506001600160401b03808211156200018f578384fd5b6200019d888389016200009d565b93506060870151915080821115620001b3578283fd5b50620001c2878288016200009d565b91505092959194509250565b610f5080620001de6000396000f3fe608060405234801561001057600080fd5b50600436106101215760003560e01c80638126af55116100ad578063e09366a711610071578063e09366a714610230578063e1b357e314610243578063e4712f3614610256578063f3810d1514610269578063f851a4401461027c57610121565b80638126af55146101cf5780638456cb59146101ef5780639714378c146101f75780639e6371ba1461020a578063cbffbe411461021d57610121565b80634e71d92d116100f45780634e71d92d14610180578063570ca735146101885780635c975abb1461019d5780636ff968c3146101b257806379599f96146101ba57610121565b806306394c9b1461012657806319b520091461013b5780633f4ba83a1461016557806348f2503d1461016d575b600080fd5b610139610134366004610b95565b610284565b005b61014e610149366004610bb0565b6102d9565b60405161015c929190610eb4565b60405180910390f35b6101396102f2565b61013961017b366004610c36565b61038c565b610139610473565b61019061050f565b60405161015c9190610d20565b6101a561051e565b60405161015c9190610d81565b61019061052e565b6101c261053d565b60405161015c9190610eab565b6101e26101dd366004610bd3565b610543565b60405161015c9190610d34565b61013961057f565b610139610205366004610d08565b610637565b610139610218366004610b95565b6106eb565b6101e261022b366004610bb0565b610758565b61014e61023e366004610bd3565b610780565b610139610251366004610c02565b6107a4565b610190610264366004610c02565b61083a565b6101e2610277366004610c02565b610885565b610190610ae3565b6000546001600160a01b031633146102b75760405162461bcd60e51b81526004016102ae90610e62565b60405180910390fd5b600380546001600160a01b0319166001600160a01b0392909216919091179055565b6006602052600090815260409020805460019091015482565b6000546001600160a01b0316331461031c5760405162461bcd60e51b81526004016102ae90610e62565b600154600160a01b900460ff166103455760405162461bcd60e51b81526004016102ae90610d8c565b6001805460ff60a01b191690556040517f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa90610382903390610d20565b60405180910390a1565b6003546001600160a01b031633146103b65760405162461bcd60e51b81526004016102ae90610e85565b84151560009081526004602052604090205483106103d357600080fd5b8115806103f0575084151560009081526004602052604090205482145b6103f957600080fd5b811561043657841515600090815260046020908152604082208054600181018255908352918190208351610434939190910191840190610af2565b505b5060408051808201825292835260208084019283529415156000908152600586528181209481529390945292909120905181559051600190910155565b60025442106104945760405162461bcd60e51b81526004016102ae90610dba565b6001546001600160a01b031633146104be5760405162461bcd60e51b81526004016102ae90610e12565b600154600080546001600160a01b0319166001600160a01b0392831617908190556040517f0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc92610382921690610d20565b6003546001600160a01b031681565b600154600160a01b900460ff1681565b6001546001600160a01b031681565b60025481565b81151560009081526005602090815260408083208484529091529020805460019091015460609161057691859190610885565b90505b92915050565b6000546001600160a01b031633146105a95760405162461bcd60e51b81526004016102ae90610e62565b600154600160a01b900460ff16156105d35760405162461bcd60e51b81526004016102ae90610e38565b60025442106105f45760405162461bcd60e51b81526004016102ae90610dba565b6001805460ff60a01b1916600160a01b1790556040517f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25890610382903390610d20565b6000546001600160a01b031633146106615760405162461bcd60e51b81526004016102ae90610e62565b60025442106106825760405162461bcd60e51b81526004016102ae90610dba565b61016e81106106a35760405162461bcd60e51b81526004016102ae90610ddb565b600280546201518083020190556040517f02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8906106e0908390610eab565b60405180910390a150565b6000546001600160a01b031633146107155760405162461bcd60e51b81526004016102ae90610e62565b60025442106107365760405162461bcd60e51b81526004016102ae90610dba565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b8015156000908152600660205260409020805460019091015460609161057991849190610885565b60056020908152600092835260408084209091529082529020805460019091015482565b6003546001600160a01b031633146107ce5760405162461bcd60e51b81526004016102ae90610e85565b82151560009081526004602052604090205482106107eb57600080fd5b821515600090815260046020526040902054811061080857600080fd5b604080518082018252928352602080840192835293151560009081526006909452909220905181559051600190910155565b6004602052826000526040600020828154811061085357fe5b90600052602060002001818154811061086857fe5b6000918252602090912001546001600160a01b0316925083915050565b6060600083156108d75784151560009081526004602052604090205484106108ac57600080fd5b84151560009081526004602052604090208054859081106108c957fe5b600091825260209091200154015b82156109255784151560009081526004602052604090205483106108fa57600080fd5b841515600090815260046020526040902080548490811061091757fe5b600091825260209091200154015b60608167ffffffffffffffff8111801561093e57600080fd5b50604051908082528060200260200182016040528015610968578160200160208202803683370190505b50905060008515610a235760005b871515600090815260046020526040902080548890811061099357fe5b600091825260209091200154811015610a215787151560009081526004602052604090208054889081106109c357fe5b9060005260206000200181815481106109d857fe5b600091825260209091200154835160018401936001600160a01b039092169185918110610a0157fe5b6001600160a01b0390921660209283029190910190910152600101610976565b505b8415610ad95760005b8715156000908152600460205260409020805487908110610a4957fe5b600091825260209091200154811015610ad7578715156000908152600460205260409020805487908110610a7957fe5b906000526020600020018181548110610a8e57fe5b600091825260209091200154835160018401936001600160a01b039092169185918110610ab757fe5b6001600160a01b0390921660209283029190910190910152600101610a2c565b505b5095945050505050565b6000546001600160a01b031681565b828054828255906000526020600020908101928215610b47579160200282015b82811115610b4757825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190610b12565b50610b53929150610b57565b5090565b610b7b91905b80821115610b535780546001600160a01b0319168155600101610b5d565b90565b80356001600160a01b038116811461057957600080fd5b600060208284031215610ba6578081fd5b6105768383610b7e565b600060208284031215610bc1578081fd5b8135610bcc81610f09565b9392505050565b60008060408385031215610be5578081fd5b82358015158114610bf4578182fd5b946020939093013593505050565b600080600060608486031215610c16578081fd5b8335610c2181610f09565b95602085013595506040909401359392505050565b600080600080600060a08688031215610c4d578081fd5b8535610c5881610f09565b945060208681013594506040870135935060608701359250608087013567ffffffffffffffff811115610c89578283fd5b80880189601f820112610c9a578384fd5b80359150610caf610caa83610ee9565b610ec2565b82815283810190828501858502840186018d1015610ccb578687fd5b8693505b84841015610cf557610ce18d82610b7e565b835260019390930192918501918501610ccf565b5080955050505050509295509295909350565b600060208284031215610d19578081fd5b5035919050565b6001600160a01b0391909116815260200190565b6020808252825182820181905260009190848201906040850190845b81811015610d755783516001600160a01b031683529284019291840191600101610d50565b50909695505050505050565b901515815260200190565b6020808252601490820152736e6f7420706175736564207269676874206e6f7760601b604082015260600190565b602080825260079082015266195e1c1a5c995960ca1b604082015260600190565b6020808252601a908201527f63616e6e6f7420657874656e6420666f7220746f6f206c6f6e67000000000000604082015260600190565b6020808252600c908201526b1d5b985d5d1a1bdc9a5e995960a21b604082015260600190565b60208082526010908201526f706175736564207269676874206e6f7760801b604082015260600190565b6020808252600990820152683737ba1030b236b4b760b91b604082015260600190565b6020808252600c908201526b3737ba1037b832b930ba37b960a11b604082015260600190565b90815260200190565b918252602082015260400190565b60405181810167ffffffffffffffff81118282101715610ee157600080fd5b604052919050565b600067ffffffffffffffff821115610eff578081fd5b5060209081020190565b8015158114610f1757600080fd5b5056fea2646970667358221220199332a6664ebda56e8d2f56cd84864cb286a9b453dc13cece3dafcb2538f27e64736f6c63430006060033"

// DeployIncognitoValidators deploys a new Ethereum contract, binding an instance of IncognitoValidators to it.
func DeployIncognitoValidators(auth *bind.TransactOpts, backend bind.ContractBackend, _admin common.Address, _operator common.Address, beaconCommittee []common.Address, bridgeCommittee []common.Address) (common.Address, *types.Transaction, *IncognitoValidators, error) {
	parsed, err := abi.JSON(strings.NewReader(IncognitoValidatorsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IncognitoValidatorsBin), backend, _admin, _operator, beaconCommittee, bridgeCommittee)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IncognitoValidators{IncognitoValidatorsCaller: IncognitoValidatorsCaller{contract: contract}, IncognitoValidatorsTransactor: IncognitoValidatorsTransactor{contract: contract}, IncognitoValidatorsFilterer: IncognitoValidatorsFilterer{contract: contract}}, nil
}

// IncognitoValidators is an auto generated Go binding around an Ethereum contract.
type IncognitoValidators struct {
	IncognitoValidatorsCaller     // Read-only binding to the contract
	IncognitoValidatorsTransactor // Write-only binding to the contract
	IncognitoValidatorsFilterer   // Log filterer for contract events
}

// IncognitoValidatorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IncognitoValidatorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncognitoValidatorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IncognitoValidatorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncognitoValidatorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IncognitoValidatorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncognitoValidatorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IncognitoValidatorsSession struct {
	Contract     *IncognitoValidators // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IncognitoValidatorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IncognitoValidatorsCallerSession struct {
	Contract *IncognitoValidatorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IncognitoValidatorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IncognitoValidatorsTransactorSession struct {
	Contract     *IncognitoValidatorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IncognitoValidatorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IncognitoValidatorsRaw struct {
	Contract *IncognitoValidators // Generic contract binding to access the raw methods on
}

// IncognitoValidatorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IncognitoValidatorsCallerRaw struct {
	Contract *IncognitoValidatorsCaller // Generic read-only contract binding to access the raw methods on
}

// IncognitoValidatorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IncognitoValidatorsTransactorRaw struct {
	Contract *IncognitoValidatorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIncognitoValidators creates a new instance of IncognitoValidators, bound to a specific deployed contract.
func NewIncognitoValidators(address common.Address, backend bind.ContractBackend) (*IncognitoValidators, error) {
	contract, err := bindIncognitoValidators(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IncognitoValidators{IncognitoValidatorsCaller: IncognitoValidatorsCaller{contract: contract}, IncognitoValidatorsTransactor: IncognitoValidatorsTransactor{contract: contract}, IncognitoValidatorsFilterer: IncognitoValidatorsFilterer{contract: contract}}, nil
}

// NewIncognitoValidatorsCaller creates a new read-only instance of IncognitoValidators, bound to a specific deployed contract.
func NewIncognitoValidatorsCaller(address common.Address, caller bind.ContractCaller) (*IncognitoValidatorsCaller, error) {
	contract, err := bindIncognitoValidators(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IncognitoValidatorsCaller{contract: contract}, nil
}

// NewIncognitoValidatorsTransactor creates a new write-only instance of IncognitoValidators, bound to a specific deployed contract.
func NewIncognitoValidatorsTransactor(address common.Address, transactor bind.ContractTransactor) (*IncognitoValidatorsTransactor, error) {
	contract, err := bindIncognitoValidators(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IncognitoValidatorsTransactor{contract: contract}, nil
}

// NewIncognitoValidatorsFilterer creates a new log filterer instance of IncognitoValidators, bound to a specific deployed contract.
func NewIncognitoValidatorsFilterer(address common.Address, filterer bind.ContractFilterer) (*IncognitoValidatorsFilterer, error) {
	contract, err := bindIncognitoValidators(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IncognitoValidatorsFilterer{contract: contract}, nil
}

// bindIncognitoValidators binds a generic wrapper to an already deployed contract.
func bindIncognitoValidators(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IncognitoValidatorsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IncognitoValidators *IncognitoValidatorsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IncognitoValidators.Contract.IncognitoValidatorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IncognitoValidators *IncognitoValidatorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncognitoValidators.Contract.IncognitoValidatorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IncognitoValidators *IncognitoValidatorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IncognitoValidators.Contract.IncognitoValidatorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IncognitoValidators *IncognitoValidatorsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IncognitoValidators.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IncognitoValidators *IncognitoValidatorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncognitoValidators.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IncognitoValidators *IncognitoValidatorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IncognitoValidators.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_IncognitoValidators *IncognitoValidatorsCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IncognitoValidators.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_IncognitoValidators *IncognitoValidatorsSession) Admin() (common.Address, error) {
	return _IncognitoValidators.Contract.Admin(&_IncognitoValidators.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_IncognitoValidators *IncognitoValidatorsCallerSession) Admin() (common.Address, error) {
	return _IncognitoValidators.Contract.Admin(&_IncognitoValidators.CallOpts)
}

// Committee is a free data retrieval call binding the contract method 0x19b52009.
//
// Solidity: function committee(bool ) view returns(uint256 oldGID, uint256 newGID)
func (_IncognitoValidators *IncognitoValidatorsCaller) Committee(opts *bind.CallOpts, arg0 bool) (struct {
	OldGID *big.Int
	NewGID *big.Int
}, error) {
	ret := new(struct {
		OldGID *big.Int
		NewGID *big.Int
	})
	out := ret
	err := _IncognitoValidators.contract.Call(opts, out, "committee", arg0)
	return *ret, err
}

// Committee is a free data retrieval call binding the contract method 0x19b52009.
//
// Solidity: function committee(bool ) view returns(uint256 oldGID, uint256 newGID)
func (_IncognitoValidators *IncognitoValidatorsSession) Committee(arg0 bool) (struct {
	OldGID *big.Int
	NewGID *big.Int
}, error) {
	return _IncognitoValidators.Contract.Committee(&_IncognitoValidators.CallOpts, arg0)
}

// Committee is a free data retrieval call binding the contract method 0x19b52009.
//
// Solidity: function committee(bool ) view returns(uint256 oldGID, uint256 newGID)
func (_IncognitoValidators *IncognitoValidatorsCallerSession) Committee(arg0 bool) (struct {
	OldGID *big.Int
	NewGID *big.Int
}, error) {
	return _IncognitoValidators.Contract.Committee(&_IncognitoValidators.CallOpts, arg0)
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() view returns(uint256)
func (_IncognitoValidators *IncognitoValidatorsCaller) Expire(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IncognitoValidators.contract.Call(opts, out, "expire")
	return *ret0, err
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() view returns(uint256)
func (_IncognitoValidators *IncognitoValidatorsSession) Expire() (*big.Int, error) {
	return _IncognitoValidators.Contract.Expire(&_IncognitoValidators.CallOpts)
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() view returns(uint256)
func (_IncognitoValidators *IncognitoValidatorsCallerSession) Expire() (*big.Int, error) {
	return _IncognitoValidators.Contract.Expire(&_IncognitoValidators.CallOpts)
}

// GroupAddrs is a free data retrieval call binding the contract method 0xe4712f36.
//
// Solidity: function groupAddrs(bool , uint256 , uint256 ) view returns(address)
func (_IncognitoValidators *IncognitoValidatorsCaller) GroupAddrs(opts *bind.CallOpts, arg0 bool, arg1 *big.Int, arg2 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IncognitoValidators.contract.Call(opts, out, "groupAddrs", arg0, arg1, arg2)
	return *ret0, err
}

// GroupAddrs is a free data retrieval call binding the contract method 0xe4712f36.
//
// Solidity: function groupAddrs(bool , uint256 , uint256 ) view returns(address)
func (_IncognitoValidators *IncognitoValidatorsSession) GroupAddrs(arg0 bool, arg1 *big.Int, arg2 *big.Int) (common.Address, error) {
	return _IncognitoValidators.Contract.GroupAddrs(&_IncognitoValidators.CallOpts, arg0, arg1, arg2)
}

// GroupAddrs is a free data retrieval call binding the contract method 0xe4712f36.
//
// Solidity: function groupAddrs(bool , uint256 , uint256 ) view returns(address)
func (_IncognitoValidators *IncognitoValidatorsCallerSession) GroupAddrs(arg0 bool, arg1 *big.Int, arg2 *big.Int) (common.Address, error) {
	return _IncognitoValidators.Contract.GroupAddrs(&_IncognitoValidators.CallOpts, arg0, arg1, arg2)
}

// LoadCandidates is a free data retrieval call binding the contract method 0x8126af55.
//
// Solidity: function loadCandidates(bool isBeacon, uint256 swapID) view returns(address[])
func (_IncognitoValidators *IncognitoValidatorsCaller) LoadCandidates(opts *bind.CallOpts, isBeacon bool, swapID *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _IncognitoValidators.contract.Call(opts, out, "loadCandidates", isBeacon, swapID)
	return *ret0, err
}

// LoadCandidates is a free data retrieval call binding the contract method 0x8126af55.
//
// Solidity: function loadCandidates(bool isBeacon, uint256 swapID) view returns(address[])
func (_IncognitoValidators *IncognitoValidatorsSession) LoadCandidates(isBeacon bool, swapID *big.Int) ([]common.Address, error) {
	return _IncognitoValidators.Contract.LoadCandidates(&_IncognitoValidators.CallOpts, isBeacon, swapID)
}

// LoadCandidates is a free data retrieval call binding the contract method 0x8126af55.
//
// Solidity: function loadCandidates(bool isBeacon, uint256 swapID) view returns(address[])
func (_IncognitoValidators *IncognitoValidatorsCallerSession) LoadCandidates(isBeacon bool, swapID *big.Int) ([]common.Address, error) {
	return _IncognitoValidators.Contract.LoadCandidates(&_IncognitoValidators.CallOpts, isBeacon, swapID)
}

// LoadCommittee is a free data retrieval call binding the contract method 0xcbffbe41.
//
// Solidity: function loadCommittee(bool isBeacon) view returns(address[])
func (_IncognitoValidators *IncognitoValidatorsCaller) LoadCommittee(opts *bind.CallOpts, isBeacon bool) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _IncognitoValidators.contract.Call(opts, out, "loadCommittee", isBeacon)
	return *ret0, err
}

// LoadCommittee is a free data retrieval call binding the contract method 0xcbffbe41.
//
// Solidity: function loadCommittee(bool isBeacon) view returns(address[])
func (_IncognitoValidators *IncognitoValidatorsSession) LoadCommittee(isBeacon bool) ([]common.Address, error) {
	return _IncognitoValidators.Contract.LoadCommittee(&_IncognitoValidators.CallOpts, isBeacon)
}

// LoadCommittee is a free data retrieval call binding the contract method 0xcbffbe41.
//
// Solidity: function loadCommittee(bool isBeacon) view returns(address[])
func (_IncognitoValidators *IncognitoValidatorsCallerSession) LoadCommittee(isBeacon bool) ([]common.Address, error) {
	return _IncognitoValidators.Contract.LoadCommittee(&_IncognitoValidators.CallOpts, isBeacon)
}

// LoadGroups is a free data retrieval call binding the contract method 0xf3810d15.
//
// Solidity: function loadGroups(bool isBeacon, uint256 oldGID, uint256 newGID) view returns(address[])
func (_IncognitoValidators *IncognitoValidatorsCaller) LoadGroups(opts *bind.CallOpts, isBeacon bool, oldGID *big.Int, newGID *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _IncognitoValidators.contract.Call(opts, out, "loadGroups", isBeacon, oldGID, newGID)
	return *ret0, err
}

// LoadGroups is a free data retrieval call binding the contract method 0xf3810d15.
//
// Solidity: function loadGroups(bool isBeacon, uint256 oldGID, uint256 newGID) view returns(address[])
func (_IncognitoValidators *IncognitoValidatorsSession) LoadGroups(isBeacon bool, oldGID *big.Int, newGID *big.Int) ([]common.Address, error) {
	return _IncognitoValidators.Contract.LoadGroups(&_IncognitoValidators.CallOpts, isBeacon, oldGID, newGID)
}

// LoadGroups is a free data retrieval call binding the contract method 0xf3810d15.
//
// Solidity: function loadGroups(bool isBeacon, uint256 oldGID, uint256 newGID) view returns(address[])
func (_IncognitoValidators *IncognitoValidatorsCallerSession) LoadGroups(isBeacon bool, oldGID *big.Int, newGID *big.Int) ([]common.Address, error) {
	return _IncognitoValidators.Contract.LoadGroups(&_IncognitoValidators.CallOpts, isBeacon, oldGID, newGID)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_IncognitoValidators *IncognitoValidatorsCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IncognitoValidators.contract.Call(opts, out, "operator")
	return *ret0, err
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_IncognitoValidators *IncognitoValidatorsSession) Operator() (common.Address, error) {
	return _IncognitoValidators.Contract.Operator(&_IncognitoValidators.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_IncognitoValidators *IncognitoValidatorsCallerSession) Operator() (common.Address, error) {
	return _IncognitoValidators.Contract.Operator(&_IncognitoValidators.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IncognitoValidators *IncognitoValidatorsCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IncognitoValidators.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IncognitoValidators *IncognitoValidatorsSession) Paused() (bool, error) {
	return _IncognitoValidators.Contract.Paused(&_IncognitoValidators.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IncognitoValidators *IncognitoValidatorsCallerSession) Paused() (bool, error) {
	return _IncognitoValidators.Contract.Paused(&_IncognitoValidators.CallOpts)
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() view returns(address)
func (_IncognitoValidators *IncognitoValidatorsCaller) Successor(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IncognitoValidators.contract.Call(opts, out, "successor")
	return *ret0, err
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() view returns(address)
func (_IncognitoValidators *IncognitoValidatorsSession) Successor() (common.Address, error) {
	return _IncognitoValidators.Contract.Successor(&_IncognitoValidators.CallOpts)
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() view returns(address)
func (_IncognitoValidators *IncognitoValidatorsCallerSession) Successor() (common.Address, error) {
	return _IncognitoValidators.Contract.Successor(&_IncognitoValidators.CallOpts)
}

// Validators is a free data retrieval call binding the contract method 0xe09366a7.
//
// Solidity: function validators(bool , uint256 ) view returns(uint256 oldGID, uint256 newGID)
func (_IncognitoValidators *IncognitoValidatorsCaller) Validators(opts *bind.CallOpts, arg0 bool, arg1 *big.Int) (struct {
	OldGID *big.Int
	NewGID *big.Int
}, error) {
	ret := new(struct {
		OldGID *big.Int
		NewGID *big.Int
	})
	out := ret
	err := _IncognitoValidators.contract.Call(opts, out, "validators", arg0, arg1)
	return *ret, err
}

// Validators is a free data retrieval call binding the contract method 0xe09366a7.
//
// Solidity: function validators(bool , uint256 ) view returns(uint256 oldGID, uint256 newGID)
func (_IncognitoValidators *IncognitoValidatorsSession) Validators(arg0 bool, arg1 *big.Int) (struct {
	OldGID *big.Int
	NewGID *big.Int
}, error) {
	return _IncognitoValidators.Contract.Validators(&_IncognitoValidators.CallOpts, arg0, arg1)
}

// Validators is a free data retrieval call binding the contract method 0xe09366a7.
//
// Solidity: function validators(bool , uint256 ) view returns(uint256 oldGID, uint256 newGID)
func (_IncognitoValidators *IncognitoValidatorsCallerSession) Validators(arg0 bool, arg1 *big.Int) (struct {
	OldGID *big.Int
	NewGID *big.Int
}, error) {
	return _IncognitoValidators.Contract.Validators(&_IncognitoValidators.CallOpts, arg0, arg1)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(address newOperator) returns()
func (_IncognitoValidators *IncognitoValidatorsTransactor) ChangeOperator(opts *bind.TransactOpts, newOperator common.Address) (*types.Transaction, error) {
	return _IncognitoValidators.contract.Transact(opts, "changeOperator", newOperator)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(address newOperator) returns()
func (_IncognitoValidators *IncognitoValidatorsSession) ChangeOperator(newOperator common.Address) (*types.Transaction, error) {
	return _IncognitoValidators.Contract.ChangeOperator(&_IncognitoValidators.TransactOpts, newOperator)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(address newOperator) returns()
func (_IncognitoValidators *IncognitoValidatorsTransactorSession) ChangeOperator(newOperator common.Address) (*types.Transaction, error) {
	return _IncognitoValidators.Contract.ChangeOperator(&_IncognitoValidators.TransactOpts, newOperator)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_IncognitoValidators *IncognitoValidatorsTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncognitoValidators.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_IncognitoValidators *IncognitoValidatorsSession) Claim() (*types.Transaction, error) {
	return _IncognitoValidators.Contract.Claim(&_IncognitoValidators.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_IncognitoValidators *IncognitoValidatorsTransactorSession) Claim() (*types.Transaction, error) {
	return _IncognitoValidators.Contract.Claim(&_IncognitoValidators.TransactOpts)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_IncognitoValidators *IncognitoValidatorsTransactor) Extend(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _IncognitoValidators.contract.Transact(opts, "extend", n)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_IncognitoValidators *IncognitoValidatorsSession) Extend(n *big.Int) (*types.Transaction, error) {
	return _IncognitoValidators.Contract.Extend(&_IncognitoValidators.TransactOpts, n)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_IncognitoValidators *IncognitoValidatorsTransactorSession) Extend(n *big.Int) (*types.Transaction, error) {
	return _IncognitoValidators.Contract.Extend(&_IncognitoValidators.TransactOpts, n)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IncognitoValidators *IncognitoValidatorsTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncognitoValidators.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IncognitoValidators *IncognitoValidatorsSession) Pause() (*types.Transaction, error) {
	return _IncognitoValidators.Contract.Pause(&_IncognitoValidators.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IncognitoValidators *IncognitoValidatorsTransactorSession) Pause() (*types.Transaction, error) {
	return _IncognitoValidators.Contract.Pause(&_IncognitoValidators.TransactOpts)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_IncognitoValidators *IncognitoValidatorsTransactor) Retire(opts *bind.TransactOpts, _successor common.Address) (*types.Transaction, error) {
	return _IncognitoValidators.contract.Transact(opts, "retire", _successor)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_IncognitoValidators *IncognitoValidatorsSession) Retire(_successor common.Address) (*types.Transaction, error) {
	return _IncognitoValidators.Contract.Retire(&_IncognitoValidators.TransactOpts, _successor)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_IncognitoValidators *IncognitoValidatorsTransactorSession) Retire(_successor common.Address) (*types.Transaction, error) {
	return _IncognitoValidators.Contract.Retire(&_IncognitoValidators.TransactOpts, _successor)
}

// StoreCandidates is a paid mutator transaction binding the contract method 0x48f2503d.
//
// Solidity: function storeCandidates(bool isBeacon, uint256 swapID, uint256 oldGID, uint256 newGID, address[] newAddrs) returns()
func (_IncognitoValidators *IncognitoValidatorsTransactor) StoreCandidates(opts *bind.TransactOpts, isBeacon bool, swapID *big.Int, oldGID *big.Int, newGID *big.Int, newAddrs []common.Address) (*types.Transaction, error) {
	return _IncognitoValidators.contract.Transact(opts, "storeCandidates", isBeacon, swapID, oldGID, newGID, newAddrs)
}

// StoreCandidates is a paid mutator transaction binding the contract method 0x48f2503d.
//
// Solidity: function storeCandidates(bool isBeacon, uint256 swapID, uint256 oldGID, uint256 newGID, address[] newAddrs) returns()
func (_IncognitoValidators *IncognitoValidatorsSession) StoreCandidates(isBeacon bool, swapID *big.Int, oldGID *big.Int, newGID *big.Int, newAddrs []common.Address) (*types.Transaction, error) {
	return _IncognitoValidators.Contract.StoreCandidates(&_IncognitoValidators.TransactOpts, isBeacon, swapID, oldGID, newGID, newAddrs)
}

// StoreCandidates is a paid mutator transaction binding the contract method 0x48f2503d.
//
// Solidity: function storeCandidates(bool isBeacon, uint256 swapID, uint256 oldGID, uint256 newGID, address[] newAddrs) returns()
func (_IncognitoValidators *IncognitoValidatorsTransactorSession) StoreCandidates(isBeacon bool, swapID *big.Int, oldGID *big.Int, newGID *big.Int, newAddrs []common.Address) (*types.Transaction, error) {
	return _IncognitoValidators.Contract.StoreCandidates(&_IncognitoValidators.TransactOpts, isBeacon, swapID, oldGID, newGID, newAddrs)
}

// StoreCommittee is a paid mutator transaction binding the contract method 0xe1b357e3.
//
// Solidity: function storeCommittee(bool isBeacon, uint256 gID1, uint256 gID2) returns()
func (_IncognitoValidators *IncognitoValidatorsTransactor) StoreCommittee(opts *bind.TransactOpts, isBeacon bool, gID1 *big.Int, gID2 *big.Int) (*types.Transaction, error) {
	return _IncognitoValidators.contract.Transact(opts, "storeCommittee", isBeacon, gID1, gID2)
}

// StoreCommittee is a paid mutator transaction binding the contract method 0xe1b357e3.
//
// Solidity: function storeCommittee(bool isBeacon, uint256 gID1, uint256 gID2) returns()
func (_IncognitoValidators *IncognitoValidatorsSession) StoreCommittee(isBeacon bool, gID1 *big.Int, gID2 *big.Int) (*types.Transaction, error) {
	return _IncognitoValidators.Contract.StoreCommittee(&_IncognitoValidators.TransactOpts, isBeacon, gID1, gID2)
}

// StoreCommittee is a paid mutator transaction binding the contract method 0xe1b357e3.
//
// Solidity: function storeCommittee(bool isBeacon, uint256 gID1, uint256 gID2) returns()
func (_IncognitoValidators *IncognitoValidatorsTransactorSession) StoreCommittee(isBeacon bool, gID1 *big.Int, gID2 *big.Int) (*types.Transaction, error) {
	return _IncognitoValidators.Contract.StoreCommittee(&_IncognitoValidators.TransactOpts, isBeacon, gID1, gID2)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IncognitoValidators *IncognitoValidatorsTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncognitoValidators.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IncognitoValidators *IncognitoValidatorsSession) Unpause() (*types.Transaction, error) {
	return _IncognitoValidators.Contract.Unpause(&_IncognitoValidators.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IncognitoValidators *IncognitoValidatorsTransactorSession) Unpause() (*types.Transaction, error) {
	return _IncognitoValidators.Contract.Unpause(&_IncognitoValidators.TransactOpts)
}

// IncognitoValidatorsClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the IncognitoValidators contract.
type IncognitoValidatorsClaimIterator struct {
	Event *IncognitoValidatorsClaim // Event containing the contract specifics and raw log

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
func (it *IncognitoValidatorsClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncognitoValidatorsClaim)
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
		it.Event = new(IncognitoValidatorsClaim)
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
func (it *IncognitoValidatorsClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncognitoValidatorsClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncognitoValidatorsClaim represents a Claim event raised by the IncognitoValidators contract.
type IncognitoValidatorsClaim struct {
	Claimer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_IncognitoValidators *IncognitoValidatorsFilterer) FilterClaim(opts *bind.FilterOpts) (*IncognitoValidatorsClaimIterator, error) {

	logs, sub, err := _IncognitoValidators.contract.FilterLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return &IncognitoValidatorsClaimIterator{contract: _IncognitoValidators.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_IncognitoValidators *IncognitoValidatorsFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *IncognitoValidatorsClaim) (event.Subscription, error) {

	logs, sub, err := _IncognitoValidators.contract.WatchLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncognitoValidatorsClaim)
				if err := _IncognitoValidators.contract.UnpackLog(event, "Claim", log); err != nil {
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
func (_IncognitoValidators *IncognitoValidatorsFilterer) ParseClaim(log types.Log) (*IncognitoValidatorsClaim, error) {
	event := new(IncognitoValidatorsClaim)
	if err := _IncognitoValidators.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncognitoValidatorsExtendIterator is returned from FilterExtend and is used to iterate over the raw logs and unpacked data for Extend events raised by the IncognitoValidators contract.
type IncognitoValidatorsExtendIterator struct {
	Event *IncognitoValidatorsExtend // Event containing the contract specifics and raw log

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
func (it *IncognitoValidatorsExtendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncognitoValidatorsExtend)
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
		it.Event = new(IncognitoValidatorsExtend)
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
func (it *IncognitoValidatorsExtendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncognitoValidatorsExtendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncognitoValidatorsExtend represents a Extend event raised by the IncognitoValidators contract.
type IncognitoValidatorsExtend struct {
	Ndays *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterExtend is a free log retrieval operation binding the contract event 0x02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8.
//
// Solidity: event Extend(uint256 ndays)
func (_IncognitoValidators *IncognitoValidatorsFilterer) FilterExtend(opts *bind.FilterOpts) (*IncognitoValidatorsExtendIterator, error) {

	logs, sub, err := _IncognitoValidators.contract.FilterLogs(opts, "Extend")
	if err != nil {
		return nil, err
	}
	return &IncognitoValidatorsExtendIterator{contract: _IncognitoValidators.contract, event: "Extend", logs: logs, sub: sub}, nil
}

// WatchExtend is a free log subscription operation binding the contract event 0x02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8.
//
// Solidity: event Extend(uint256 ndays)
func (_IncognitoValidators *IncognitoValidatorsFilterer) WatchExtend(opts *bind.WatchOpts, sink chan<- *IncognitoValidatorsExtend) (event.Subscription, error) {

	logs, sub, err := _IncognitoValidators.contract.WatchLogs(opts, "Extend")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncognitoValidatorsExtend)
				if err := _IncognitoValidators.contract.UnpackLog(event, "Extend", log); err != nil {
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
func (_IncognitoValidators *IncognitoValidatorsFilterer) ParseExtend(log types.Log) (*IncognitoValidatorsExtend, error) {
	event := new(IncognitoValidatorsExtend)
	if err := _IncognitoValidators.contract.UnpackLog(event, "Extend", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncognitoValidatorsPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the IncognitoValidators contract.
type IncognitoValidatorsPausedIterator struct {
	Event *IncognitoValidatorsPaused // Event containing the contract specifics and raw log

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
func (it *IncognitoValidatorsPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncognitoValidatorsPaused)
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
		it.Event = new(IncognitoValidatorsPaused)
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
func (it *IncognitoValidatorsPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncognitoValidatorsPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncognitoValidatorsPaused represents a Paused event raised by the IncognitoValidators contract.
type IncognitoValidatorsPaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address pauser)
func (_IncognitoValidators *IncognitoValidatorsFilterer) FilterPaused(opts *bind.FilterOpts) (*IncognitoValidatorsPausedIterator, error) {

	logs, sub, err := _IncognitoValidators.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &IncognitoValidatorsPausedIterator{contract: _IncognitoValidators.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address pauser)
func (_IncognitoValidators *IncognitoValidatorsFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *IncognitoValidatorsPaused) (event.Subscription, error) {

	logs, sub, err := _IncognitoValidators.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncognitoValidatorsPaused)
				if err := _IncognitoValidators.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_IncognitoValidators *IncognitoValidatorsFilterer) ParsePaused(log types.Log) (*IncognitoValidatorsPaused, error) {
	event := new(IncognitoValidatorsPaused)
	if err := _IncognitoValidators.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncognitoValidatorsUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the IncognitoValidators contract.
type IncognitoValidatorsUnpausedIterator struct {
	Event *IncognitoValidatorsUnpaused // Event containing the contract specifics and raw log

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
func (it *IncognitoValidatorsUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncognitoValidatorsUnpaused)
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
		it.Event = new(IncognitoValidatorsUnpaused)
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
func (it *IncognitoValidatorsUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncognitoValidatorsUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncognitoValidatorsUnpaused represents a Unpaused event raised by the IncognitoValidators contract.
type IncognitoValidatorsUnpaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address pauser)
func (_IncognitoValidators *IncognitoValidatorsFilterer) FilterUnpaused(opts *bind.FilterOpts) (*IncognitoValidatorsUnpausedIterator, error) {

	logs, sub, err := _IncognitoValidators.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &IncognitoValidatorsUnpausedIterator{contract: _IncognitoValidators.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address pauser)
func (_IncognitoValidators *IncognitoValidatorsFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *IncognitoValidatorsUnpaused) (event.Subscription, error) {

	logs, sub, err := _IncognitoValidators.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncognitoValidatorsUnpaused)
				if err := _IncognitoValidators.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_IncognitoValidators *IncognitoValidatorsFilterer) ParseUnpaused(log types.Log) (*IncognitoValidatorsUnpaused, error) {
	event := new(IncognitoValidatorsUnpaused)
	if err := _IncognitoValidators.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OperableABI is the input ABI used to generate the binding from.
const OperableABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ndays\",\"type\":\"uint256\"}],\"name\":\"Extend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"changeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expire\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"extend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_successor\",\"type\":\"address\"}],\"name\":\"retire\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"successor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OperableFuncSigs maps the 4-byte function signature to its string representation.
var OperableFuncSigs = map[string]string{
	"f851a440": "admin()",
	"06394c9b": "changeOperator(address)",
	"4e71d92d": "claim()",
	"79599f96": "expire()",
	"9714378c": "extend(uint256)",
	"570ca735": "operator()",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"9e6371ba": "retire(address)",
	"6ff968c3": "successor()",
	"3f4ba83a": "unpause()",
}

// OperableBin is the compiled bytecode used for deploying new contracts.
var OperableBin = "0x608060405234801561001057600080fd5b5060405161077438038061077483398101604081905261002f91610077565b600080546001600160a01b039384166001600160a01b0319918216179091556001805460ff60a01b191690556301e133804201600255600380549290931691161790556100c8565b60008060408385031215610089578182fd5b8251610094816100b0565b60208401519092506100a5816100b0565b809150509250929050565b6001600160a01b03811681146100c557600080fd5b50565b61069d806100d76000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80636ff968c3116100715780636ff968c31461010657806379599f961461010e5780638456cb59146101235780639714378c1461012b5780639e6371ba1461013e578063f851a44014610151576100a9565b806306394c9b146100ae5780633f4ba83a146100c35780634e71d92d146100cb578063570ca735146100d35780635c975abb146100f1575b600080fd5b6100c16100bc366004610500565b610159565b005b6100c16101ae565b6100c1610248565b6100db6102e4565b6040516100e89190610546565b60405180910390f35b6100f96102f3565b6040516100e8919061055a565b6100db610303565b610116610312565b6040516100e8919061065e565b6100c1610318565b6100c161013936600461052e565b6103d0565b6100c161014c366004610500565b610484565b6100db6104f1565b6000546001600160a01b0316331461018c5760405162461bcd60e51b81526004016101839061063b565b60405180910390fd5b600380546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b031633146101d85760405162461bcd60e51b81526004016101839061063b565b600154600160a01b900460ff166102015760405162461bcd60e51b815260040161018390610565565b6001805460ff60a01b191690556040517f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa9061023e903390610546565b60405180910390a1565b60025442106102695760405162461bcd60e51b815260040161018390610593565b6001546001600160a01b031633146102935760405162461bcd60e51b8152600401610183906105eb565b600154600080546001600160a01b0319166001600160a01b0392831617908190556040517f0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc9261023e921690610546565b6003546001600160a01b031681565b600154600160a01b900460ff1681565b6001546001600160a01b031681565b60025481565b6000546001600160a01b031633146103425760405162461bcd60e51b81526004016101839061063b565b600154600160a01b900460ff161561036c5760405162461bcd60e51b815260040161018390610611565b600254421061038d5760405162461bcd60e51b815260040161018390610593565b6001805460ff60a01b1916600160a01b1790556040517f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2589061023e903390610546565b6000546001600160a01b031633146103fa5760405162461bcd60e51b81526004016101839061063b565b600254421061041b5760405162461bcd60e51b815260040161018390610593565b61016e811061043c5760405162461bcd60e51b8152600401610183906105b4565b600280546201518083020190556040517f02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e89061047990839061065e565b60405180910390a150565b6000546001600160a01b031633146104ae5760405162461bcd60e51b81526004016101839061063b565b60025442106104cf5760405162461bcd60e51b815260040161018390610593565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b031681565b600060208284031215610511578081fd5b81356001600160a01b0381168114610527578182fd5b9392505050565b60006020828403121561053f578081fd5b5035919050565b6001600160a01b0391909116815260200190565b901515815260200190565b6020808252601490820152736e6f7420706175736564207269676874206e6f7760601b604082015260600190565b602080825260079082015266195e1c1a5c995960ca1b604082015260600190565b6020808252601a908201527f63616e6e6f7420657874656e6420666f7220746f6f206c6f6e67000000000000604082015260600190565b6020808252600c908201526b1d5b985d5d1a1bdc9a5e995960a21b604082015260600190565b60208082526010908201526f706175736564207269676874206e6f7760801b604082015260600190565b6020808252600990820152683737ba1030b236b4b760b91b604082015260600190565b9081526020019056fea26469706673582212200e549cff2d1f75820f77ce5c7f62b4cdc4592e9bb12440082cf5ca477eafdb5264736f6c63430006060033"

// DeployOperable deploys a new Ethereum contract, binding an instance of Operable to it.
func DeployOperable(auth *bind.TransactOpts, backend bind.ContractBackend, _admin common.Address, _operator common.Address) (common.Address, *types.Transaction, *Operable, error) {
	parsed, err := abi.JSON(strings.NewReader(OperableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OperableBin), backend, _admin, _operator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Operable{OperableCaller: OperableCaller{contract: contract}, OperableTransactor: OperableTransactor{contract: contract}, OperableFilterer: OperableFilterer{contract: contract}}, nil
}

// Operable is an auto generated Go binding around an Ethereum contract.
type Operable struct {
	OperableCaller     // Read-only binding to the contract
	OperableTransactor // Write-only binding to the contract
	OperableFilterer   // Log filterer for contract events
}

// OperableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OperableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OperableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OperableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OperableSession struct {
	Contract     *Operable         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OperableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OperableCallerSession struct {
	Contract *OperableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// OperableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OperableTransactorSession struct {
	Contract     *OperableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// OperableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OperableRaw struct {
	Contract *Operable // Generic contract binding to access the raw methods on
}

// OperableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OperableCallerRaw struct {
	Contract *OperableCaller // Generic read-only contract binding to access the raw methods on
}

// OperableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OperableTransactorRaw struct {
	Contract *OperableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOperable creates a new instance of Operable, bound to a specific deployed contract.
func NewOperable(address common.Address, backend bind.ContractBackend) (*Operable, error) {
	contract, err := bindOperable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Operable{OperableCaller: OperableCaller{contract: contract}, OperableTransactor: OperableTransactor{contract: contract}, OperableFilterer: OperableFilterer{contract: contract}}, nil
}

// NewOperableCaller creates a new read-only instance of Operable, bound to a specific deployed contract.
func NewOperableCaller(address common.Address, caller bind.ContractCaller) (*OperableCaller, error) {
	contract, err := bindOperable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OperableCaller{contract: contract}, nil
}

// NewOperableTransactor creates a new write-only instance of Operable, bound to a specific deployed contract.
func NewOperableTransactor(address common.Address, transactor bind.ContractTransactor) (*OperableTransactor, error) {
	contract, err := bindOperable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OperableTransactor{contract: contract}, nil
}

// NewOperableFilterer creates a new log filterer instance of Operable, bound to a specific deployed contract.
func NewOperableFilterer(address common.Address, filterer bind.ContractFilterer) (*OperableFilterer, error) {
	contract, err := bindOperable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OperableFilterer{contract: contract}, nil
}

// bindOperable binds a generic wrapper to an already deployed contract.
func bindOperable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OperableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Operable *OperableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Operable.Contract.OperableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Operable *OperableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Operable.Contract.OperableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Operable *OperableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Operable.Contract.OperableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Operable *OperableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Operable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Operable *OperableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Operable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Operable *OperableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Operable.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Operable *OperableCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Operable.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Operable *OperableSession) Admin() (common.Address, error) {
	return _Operable.Contract.Admin(&_Operable.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Operable *OperableCallerSession) Admin() (common.Address, error) {
	return _Operable.Contract.Admin(&_Operable.CallOpts)
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() view returns(uint256)
func (_Operable *OperableCaller) Expire(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Operable.contract.Call(opts, out, "expire")
	return *ret0, err
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() view returns(uint256)
func (_Operable *OperableSession) Expire() (*big.Int, error) {
	return _Operable.Contract.Expire(&_Operable.CallOpts)
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() view returns(uint256)
func (_Operable *OperableCallerSession) Expire() (*big.Int, error) {
	return _Operable.Contract.Expire(&_Operable.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Operable *OperableCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Operable.contract.Call(opts, out, "operator")
	return *ret0, err
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Operable *OperableSession) Operator() (common.Address, error) {
	return _Operable.Contract.Operator(&_Operable.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Operable *OperableCallerSession) Operator() (common.Address, error) {
	return _Operable.Contract.Operator(&_Operable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Operable *OperableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Operable.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Operable *OperableSession) Paused() (bool, error) {
	return _Operable.Contract.Paused(&_Operable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Operable *OperableCallerSession) Paused() (bool, error) {
	return _Operable.Contract.Paused(&_Operable.CallOpts)
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() view returns(address)
func (_Operable *OperableCaller) Successor(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Operable.contract.Call(opts, out, "successor")
	return *ret0, err
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() view returns(address)
func (_Operable *OperableSession) Successor() (common.Address, error) {
	return _Operable.Contract.Successor(&_Operable.CallOpts)
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() view returns(address)
func (_Operable *OperableCallerSession) Successor() (common.Address, error) {
	return _Operable.Contract.Successor(&_Operable.CallOpts)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(address newOperator) returns()
func (_Operable *OperableTransactor) ChangeOperator(opts *bind.TransactOpts, newOperator common.Address) (*types.Transaction, error) {
	return _Operable.contract.Transact(opts, "changeOperator", newOperator)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(address newOperator) returns()
func (_Operable *OperableSession) ChangeOperator(newOperator common.Address) (*types.Transaction, error) {
	return _Operable.Contract.ChangeOperator(&_Operable.TransactOpts, newOperator)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(address newOperator) returns()
func (_Operable *OperableTransactorSession) ChangeOperator(newOperator common.Address) (*types.Transaction, error) {
	return _Operable.Contract.ChangeOperator(&_Operable.TransactOpts, newOperator)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Operable *OperableTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Operable.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Operable *OperableSession) Claim() (*types.Transaction, error) {
	return _Operable.Contract.Claim(&_Operable.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Operable *OperableTransactorSession) Claim() (*types.Transaction, error) {
	return _Operable.Contract.Claim(&_Operable.TransactOpts)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_Operable *OperableTransactor) Extend(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _Operable.contract.Transact(opts, "extend", n)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_Operable *OperableSession) Extend(n *big.Int) (*types.Transaction, error) {
	return _Operable.Contract.Extend(&_Operable.TransactOpts, n)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_Operable *OperableTransactorSession) Extend(n *big.Int) (*types.Transaction, error) {
	return _Operable.Contract.Extend(&_Operable.TransactOpts, n)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Operable *OperableTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Operable.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Operable *OperableSession) Pause() (*types.Transaction, error) {
	return _Operable.Contract.Pause(&_Operable.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Operable *OperableTransactorSession) Pause() (*types.Transaction, error) {
	return _Operable.Contract.Pause(&_Operable.TransactOpts)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_Operable *OperableTransactor) Retire(opts *bind.TransactOpts, _successor common.Address) (*types.Transaction, error) {
	return _Operable.contract.Transact(opts, "retire", _successor)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_Operable *OperableSession) Retire(_successor common.Address) (*types.Transaction, error) {
	return _Operable.Contract.Retire(&_Operable.TransactOpts, _successor)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_Operable *OperableTransactorSession) Retire(_successor common.Address) (*types.Transaction, error) {
	return _Operable.Contract.Retire(&_Operable.TransactOpts, _successor)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Operable *OperableTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Operable.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Operable *OperableSession) Unpause() (*types.Transaction, error) {
	return _Operable.Contract.Unpause(&_Operable.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Operable *OperableTransactorSession) Unpause() (*types.Transaction, error) {
	return _Operable.Contract.Unpause(&_Operable.TransactOpts)
}

// OperableClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the Operable contract.
type OperableClaimIterator struct {
	Event *OperableClaim // Event containing the contract specifics and raw log

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
func (it *OperableClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperableClaim)
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
		it.Event = new(OperableClaim)
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
func (it *OperableClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperableClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperableClaim represents a Claim event raised by the Operable contract.
type OperableClaim struct {
	Claimer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_Operable *OperableFilterer) FilterClaim(opts *bind.FilterOpts) (*OperableClaimIterator, error) {

	logs, sub, err := _Operable.contract.FilterLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return &OperableClaimIterator{contract: _Operable.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_Operable *OperableFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *OperableClaim) (event.Subscription, error) {

	logs, sub, err := _Operable.contract.WatchLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperableClaim)
				if err := _Operable.contract.UnpackLog(event, "Claim", log); err != nil {
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
func (_Operable *OperableFilterer) ParseClaim(log types.Log) (*OperableClaim, error) {
	event := new(OperableClaim)
	if err := _Operable.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OperableExtendIterator is returned from FilterExtend and is used to iterate over the raw logs and unpacked data for Extend events raised by the Operable contract.
type OperableExtendIterator struct {
	Event *OperableExtend // Event containing the contract specifics and raw log

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
func (it *OperableExtendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperableExtend)
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
		it.Event = new(OperableExtend)
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
func (it *OperableExtendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperableExtendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperableExtend represents a Extend event raised by the Operable contract.
type OperableExtend struct {
	Ndays *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterExtend is a free log retrieval operation binding the contract event 0x02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8.
//
// Solidity: event Extend(uint256 ndays)
func (_Operable *OperableFilterer) FilterExtend(opts *bind.FilterOpts) (*OperableExtendIterator, error) {

	logs, sub, err := _Operable.contract.FilterLogs(opts, "Extend")
	if err != nil {
		return nil, err
	}
	return &OperableExtendIterator{contract: _Operable.contract, event: "Extend", logs: logs, sub: sub}, nil
}

// WatchExtend is a free log subscription operation binding the contract event 0x02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8.
//
// Solidity: event Extend(uint256 ndays)
func (_Operable *OperableFilterer) WatchExtend(opts *bind.WatchOpts, sink chan<- *OperableExtend) (event.Subscription, error) {

	logs, sub, err := _Operable.contract.WatchLogs(opts, "Extend")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperableExtend)
				if err := _Operable.contract.UnpackLog(event, "Extend", log); err != nil {
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
func (_Operable *OperableFilterer) ParseExtend(log types.Log) (*OperableExtend, error) {
	event := new(OperableExtend)
	if err := _Operable.contract.UnpackLog(event, "Extend", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OperablePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Operable contract.
type OperablePausedIterator struct {
	Event *OperablePaused // Event containing the contract specifics and raw log

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
func (it *OperablePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperablePaused)
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
		it.Event = new(OperablePaused)
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
func (it *OperablePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperablePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperablePaused represents a Paused event raised by the Operable contract.
type OperablePaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address pauser)
func (_Operable *OperableFilterer) FilterPaused(opts *bind.FilterOpts) (*OperablePausedIterator, error) {

	logs, sub, err := _Operable.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &OperablePausedIterator{contract: _Operable.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address pauser)
func (_Operable *OperableFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *OperablePaused) (event.Subscription, error) {

	logs, sub, err := _Operable.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperablePaused)
				if err := _Operable.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Operable *OperableFilterer) ParsePaused(log types.Log) (*OperablePaused, error) {
	event := new(OperablePaused)
	if err := _Operable.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OperableUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Operable contract.
type OperableUnpausedIterator struct {
	Event *OperableUnpaused // Event containing the contract specifics and raw log

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
func (it *OperableUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperableUnpaused)
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
		it.Event = new(OperableUnpaused)
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
func (it *OperableUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperableUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperableUnpaused represents a Unpaused event raised by the Operable contract.
type OperableUnpaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address pauser)
func (_Operable *OperableFilterer) FilterUnpaused(opts *bind.FilterOpts) (*OperableUnpausedIterator, error) {

	logs, sub, err := _Operable.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &OperableUnpausedIterator{contract: _Operable.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address pauser)
func (_Operable *OperableFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *OperableUnpaused) (event.Subscription, error) {

	logs, sub, err := _Operable.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperableUnpaused)
				if err := _Operable.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Operable *OperableFilterer) ParseUnpaused(log types.Log) (*OperableUnpaused, error) {
	event := new(OperableUnpaused)
	if err := _Operable.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}
