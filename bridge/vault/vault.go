// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vault

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

// VaultBurnInstData is an auto generated low-level Go binding around an user-defined struct.
type VaultBurnInstData struct {
	Meta   uint8
	Shard  uint8
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Itx    [32]byte
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
var AdminPausableBin = "0x608060405234801561001057600080fd5b506040516106fc3803806106fc8339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b03199092169190911790556001805460ff60a01b191690556305a39a8042016002556106808061007c6000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c806379599f961161006657806379599f96146100ea5780638456cb59146101045780639714378c1461010c5780639e6371ba14610129578063f851a4401461014f57610093565b80633f4ba83a146100985780634e71d92d146100a25780635c975abb146100aa5780636ff968c3146100c6575b600080fd5b6100a0610157565b005b6100a0610239565b6100b2610320565b604080519115158252519081900360200190f35b6100ce610330565b604080516001600160a01b039092168252519081900360200190f35b6100f261033f565b60408051918252519081900360200190f35b6100a0610345565b6100a06004803603602081101561012257600080fd5b503561046a565b6100a06004803603602081101561013f57600080fd5b50356001600160a01b031661058e565b6100ce61063b565b6000546001600160a01b031633146101a2576040805162461bcd60e51b81526020600482015260096024820152683737ba1030b236b4b760b91b604482015290519081900360640190fd5b600154600160a01b900460ff166101f7576040805162461bcd60e51b81526020600482015260146024820152736e6f7420706175736564207269676874206e6f7760601b604482015290519081900360640190fd5b6001805460ff60a01b191690556040805133815290517f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa9181900360200190a1565b6002544210610279576040805162461bcd60e51b8152602060048201526007602482015266195e1c1a5c995960ca1b604482015290519081900360640190fd5b6001546001600160a01b031633146102c7576040805162461bcd60e51b815260206004820152600c60248201526b1d5b985d5d1a1bdc9a5e995960a21b604482015290519081900360640190fd5b600154600080546001600160a01b0319166001600160a01b0392831617908190556040805191909216815290517f0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc9181900360200190a1565b600154600160a01b900460ff1681565b6001546001600160a01b031681565b60025481565b6000546001600160a01b03163314610390576040805162461bcd60e51b81526020600482015260096024820152683737ba1030b236b4b760b91b604482015290519081900360640190fd5b600154600160a01b900460ff16156103e2576040805162461bcd60e51b815260206004820152601060248201526f706175736564207269676874206e6f7760801b604482015290519081900360640190fd5b6002544210610422576040805162461bcd60e51b8152602060048201526007602482015266195e1c1a5c995960ca1b604482015290519081900360640190fd5b6001805460ff60a01b1916600160a01b1790556040805133815290517f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2589181900360200190a1565b6000546001600160a01b031633146104b5576040805162461bcd60e51b81526020600482015260096024820152683737ba1030b236b4b760b91b604482015290519081900360640190fd5b60025442106104f5576040805162461bcd60e51b8152602060048201526007602482015266195e1c1a5c995960ca1b604482015290519081900360640190fd5b61016e811061054b576040805162461bcd60e51b815260206004820152601a60248201527f63616e6e6f7420657874656e6420666f7220746f6f206c6f6e67000000000000604482015290519081900360640190fd5b600280546201518083020190556040805182815290517f02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e89181900360200190a150565b6000546001600160a01b031633146105d9576040805162461bcd60e51b81526020600482015260096024820152683737ba1030b236b4b760b91b604482015290519081900360640190fd5b6002544210610619576040805162461bcd60e51b8152602060048201526007602482015266195e1c1a5c995960ca1b604482015290519081900360640190fd5b600180546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b03168156fea2646970667358221220dd1d0e57e511ff44730bf8bfce7030942947570ebce1cea5e5ddb685c1dbc7e664736f6c63430006060033"

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
// Solidity: function allowance(address owner, address spender) view returns(uint256)
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
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
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
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
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
// Solidity: function decimals() view returns(uint256)
func (_IERC20 *IERC20Session) Decimals() (*big.Int, error) {
	return _IERC20.Contract.Decimals(&_IERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_IERC20 *IERC20CallerSession) Decimals() (*big.Int, error) {
	return _IERC20.Contract.Decimals(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
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
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
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
// Solidity: function instructionApproved(bool , bytes32 , uint256 , bytes32[] , bool[] , bytes32 , bytes32 , uint256[] , uint8[] , bytes32[] , bytes32[] ) view returns(bool)
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
// Solidity: function instructionApproved(bool , bytes32 , uint256 , bytes32[] , bool[] , bytes32 , bytes32 , uint256[] , uint8[] , bytes32[] , bytes32[] ) view returns(bool)
func (_Incognito *IncognitoSession) InstructionApproved(arg0 bool, arg1 [32]byte, arg2 *big.Int, arg3 [][32]byte, arg4 []bool, arg5 [32]byte, arg6 [32]byte, arg7 []*big.Int, arg8 []uint8, arg9 [][32]byte, arg10 [][32]byte) (bool, error) {
	return _Incognito.Contract.InstructionApproved(&_Incognito.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
}

// InstructionApproved is a free data retrieval call binding the contract method 0xf65d2116.
//
// Solidity: function instructionApproved(bool , bytes32 , uint256 , bytes32[] , bool[] , bytes32 , bytes32 , uint256[] , uint8[] , bytes32[] , bytes32[] ) view returns(bool)
func (_Incognito *IncognitoCallerSession) InstructionApproved(arg0 bool, arg1 [32]byte, arg2 *big.Int, arg3 [][32]byte, arg4 []bool, arg5 [32]byte, arg6 [32]byte, arg7 []*big.Int, arg8 []uint8, arg9 [][32]byte, arg10 [][32]byte) (bool, error) {
	return _Incognito.Contract.InstructionApproved(&_Incognito.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
}

// LockerABI is the input ABI used to generate the binding from.
const LockerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"give\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// LockerFuncSigs maps the 4-byte function signature to its string representation.
var LockerFuncSigs = map[string]string{
	"ec8b283c": "give(address,address,uint256)",
}

// Locker is an auto generated Go binding around an Ethereum contract.
type Locker struct {
	LockerCaller     // Read-only binding to the contract
	LockerTransactor // Write-only binding to the contract
	LockerFilterer   // Log filterer for contract events
}

// LockerCaller is an auto generated read-only Go binding around an Ethereum contract.
type LockerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LockerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LockerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LockerSession struct {
	Contract     *Locker           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LockerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LockerCallerSession struct {
	Contract *LockerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LockerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LockerTransactorSession struct {
	Contract     *LockerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LockerRaw is an auto generated low-level Go binding around an Ethereum contract.
type LockerRaw struct {
	Contract *Locker // Generic contract binding to access the raw methods on
}

// LockerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LockerCallerRaw struct {
	Contract *LockerCaller // Generic read-only contract binding to access the raw methods on
}

// LockerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LockerTransactorRaw struct {
	Contract *LockerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLocker creates a new instance of Locker, bound to a specific deployed contract.
func NewLocker(address common.Address, backend bind.ContractBackend) (*Locker, error) {
	contract, err := bindLocker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Locker{LockerCaller: LockerCaller{contract: contract}, LockerTransactor: LockerTransactor{contract: contract}, LockerFilterer: LockerFilterer{contract: contract}}, nil
}

// NewLockerCaller creates a new read-only instance of Locker, bound to a specific deployed contract.
func NewLockerCaller(address common.Address, caller bind.ContractCaller) (*LockerCaller, error) {
	contract, err := bindLocker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LockerCaller{contract: contract}, nil
}

// NewLockerTransactor creates a new write-only instance of Locker, bound to a specific deployed contract.
func NewLockerTransactor(address common.Address, transactor bind.ContractTransactor) (*LockerTransactor, error) {
	contract, err := bindLocker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LockerTransactor{contract: contract}, nil
}

// NewLockerFilterer creates a new log filterer instance of Locker, bound to a specific deployed contract.
func NewLockerFilterer(address common.Address, filterer bind.ContractFilterer) (*LockerFilterer, error) {
	contract, err := bindLocker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LockerFilterer{contract: contract}, nil
}

// bindLocker binds a generic wrapper to an already deployed contract.
func bindLocker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LockerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Locker *LockerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Locker.Contract.LockerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Locker *LockerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Locker.Contract.LockerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Locker *LockerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Locker.Contract.LockerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Locker *LockerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Locker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Locker *LockerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Locker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Locker *LockerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Locker.Contract.contract.Transact(opts, method, params...)
}

// Give is a paid mutator transaction binding the contract method 0xec8b283c.
//
// Solidity: function give(address to, address token, uint256 amount) returns()
func (_Locker *LockerTransactor) Give(opts *bind.TransactOpts, to common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Locker.contract.Transact(opts, "give", to, token, amount)
}

// Give is a paid mutator transaction binding the contract method 0xec8b283c.
//
// Solidity: function give(address to, address token, uint256 amount) returns()
func (_Locker *LockerSession) Give(to common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Locker.Contract.Give(&_Locker.TransactOpts, to, token, amount)
}

// Give is a paid mutator transaction binding the contract method 0xec8b283c.
//
// Solidity: function give(address to, address token, uint256 amount) returns()
func (_Locker *LockerTransactorSession) Give(to common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Locker.Contract.Give(&_Locker.TransactOpts, to, token, amount)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212207d8f0bf289eb365016e4d221ce5f59395c0cb9c18a598fd58bcac1a08c85485464736f6c63430006060033"

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
const VaultABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"incognitoProxyAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_prevVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_locker\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ndays\",\"type\":\"uint256\"}],\"name\":\"Extend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVault\",\"type\":\"address\"}],\"name\":\"Migrate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"}],\"name\":\"MoveAssets\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newIncognitoProxy\",\"type\":\"address\"}],\"name\":\"UpdateIncognitoProxy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"UpdateTokenTotal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ETH_TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"balanceOfLocker\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipientToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"exchangeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"timestamp\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signData\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"recipientTokens\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"exchangeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"timestamp\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signData\",\"type\":\"bytes\"}],\"name\":\"executeMulti\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expire\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"extend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getDepositedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incognito\",\"outputs\":[{\"internalType\":\"contractIncognito\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"isSigDataUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"isWithdrawed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"locker\",\"outputs\":[{\"internalType\":\"contractLocker\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_newVault\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"migration\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"}],\"name\":\"moveAssets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newVault\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"notEntered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"}],\"name\":\"parseBurnInst\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"meta\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"shard\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"itx\",\"type\":\"bytes32\"}],\"internalType\":\"structVault.BurnInstData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prevVault\",\"outputs\":[{\"internalType\":\"contractWithdrawable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"timestamp\",\"type\":\"bytes\"}],\"name\":\"requestWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_successor\",\"type\":\"address\"}],\"name\":\"retire\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"sigDataUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"sigToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"heights\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"instPaths\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"instPathIsLefts\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"instRoots\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"sigIdxs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"sigVs\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigRs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigSs\",\"type\":\"bytes32[]\"}],\"name\":\"submitBurnProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"successor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalDepositedToSCAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"updateAssets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newIncognitoProxy\",\"type\":\"address\"}],\"name\":\"updateIncognitoProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newLocker\",\"type\":\"address\"}],\"name\":\"updateLocker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"heights\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"instPaths\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"instPathIsLefts\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"instRoots\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"sigIdxs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"sigVs\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigRs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigSs\",\"type\":\"bytes32[]\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawRequests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"withdrawed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// VaultFuncSigs maps the 4-byte function signature to its string representation.
var VaultFuncSigs = map[string]string{
	"58bc8337": "ETH_TOKEN()",
	"f851a440": "admin()",
	"9b780b2a": "balanceOfLocker(address)",
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
	"d7b96d4e": "locker()",
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
	"73bf9651": "submitBurnProof(bytes,uint256,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])",
	"6ff968c3": "successor()",
	"6304541c": "totalDepositedToSCAmount(address)",
	"3f4ba83a": "unpause()",
	"1ed4276d": "updateAssets(address[],uint256[])",
	"3a51913d": "updateIncognitoProxy(address)",
	"e1d5ea8a": "updateLocker(address)",
	"1beb7de2": "withdraw(bytes,uint256,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])",
	"65b5a00f": "withdrawRequests(address,address)",
	"dca40d9e": "withdrawed(bytes32)",
}

// VaultBin is the compiled bytecode used for deploying new contracts.
var VaultBin = "0x6080604052600b805460ff60a01b1916600160a01b1790553480156200002457600080fd5b50604051620048f6380380620048f68339810160408190526200004791620000d8565b600080546001600160a01b03199081166001600160a01b03968716179091556001805460ff60a01b19169055426305a39a800160025560098054821694861694909417909355600a80548416928516929092179091556008805483169190931617909155600b8054909116905562000138565b80516001600160a01b0381168114620000d257600080fd5b92915050565b60008060008060808587031215620000ee578384fd5b620000fa8686620000ba565b93506200010b8660208701620000ba565b92506200011c8660408701620000ba565b91506200012d8660608701620000ba565b905092959194509250565b6147ae80620001486000396000f3fe6080604052600436106102345760003560e01c80638588ccd61161012e578063b69f6511116100ab578063e1d5ea8a1161006f578063e1d5ea8a1461061f578063e4bd70741461063f578063f75b98ce1461065f578063f851a4401461067f578063fa84702e146106945761023b565b8063b69f65111461058a578063ce5494bb1461059d578063cf54aaa0146105bd578063d7b96d4e146105ea578063dca40d9e146105ff5761023b565b8063995fac11116100f2578063995fac11146105025780639b780b2a146105225780639e6371ba14610542578063a26e118614610562578063a3f5d8cc146105755761023b565b80638588ccd61461048557806387add4401461049857806388aaf0c8146104b85780638a984538146104cd5780639714378c146104e25761023b565b80635a67cb87116101bc57806373bf96511161018057806373bf9651146103ee578063749c5f861461040e57806379599f961461042e5780637e16e6e1146104435780638456cb59146104705761023b565b80635a67cb87146103645780635c975abb146103775780636304541c1461038c57806365b5a00f146103b95780636ff968c3146103d95761023b565b80633a51913d116102035780633a51913d146102d85780633f4ba83a146102f85780633fec6b401461030d5780634e71d92d1461033a57806358bc83371461034f5761023b565b80630c4f5039146102405780631beb7de2146102625780631ea1940e146102825780631ed4276d146102b85761023b565b3661023b57005b600080fd5b34801561024c57600080fd5b5061026061025b36600461378c565b6106a9565b005b34801561026e57600080fd5b5061026061027d3660046139f2565b6108ea565b34801561028e57600080fd5b506102a261029d366004613966565b610c2f565b6040516102af9190613f64565b60405180910390f35b3480156102c457600080fd5b506102a26102d3366004613724565b610c44565b3480156102e457600080fd5b506102606102f3366004613589565b6110b7565b34801561030457600080fd5b50610260611186565b34801561031957600080fd5b5061032d6103283660046139b0565b611220565b6040516102af9190613def565b34801561034657600080fd5b506102606112af565b34801561035b57600080fd5b5061032d61134b565b6102606103723660046136ce565b611350565b34801561038357600080fd5b506102a2611671565b34801561039857600080fd5b506103ac6103a7366004613589565b611681565b6040516102af9190614019565b3480156103c557600080fd5b506103ac6103d43660046135d2565b611693565b3480156103e557600080fd5b5061032d6116b0565b3480156103fa57600080fd5b506102606104093660046139f2565b6116bf565b34801561041a57600080fd5b506102a2610429366004613966565b611a00565b34801561043a57600080fd5b506103ac611abf565b34801561044f57600080fd5b5061046361045e36600461397e565b611ac5565b6040516102af9190614677565b34801561047c57600080fd5b50610260611b47565b61026061049336600461360a565b611bff565b3480156104a457600080fd5b506102606104b3366004613b23565b611e9a565b3480156104c457600080fd5b5061032d6120a2565b3480156104d957600080fd5b5061032d6120b1565b3480156104ee57600080fd5b506102606104fd366004613966565b6120c0565b34801561050e57600080fd5b506102a261051d3660046135d2565b612169565b34801561052e57600080fd5b506103ac61053d366004613589565b612189565b34801561054e57600080fd5b5061026061055d366004613589565b61222e565b61026061057036600461397e565b61229b565b34801561058157600080fd5b506102a2612409565b610260610598366004613887565b612419565b3480156105a957600080fd5b506102606105b8366004613589565b612a2c565b3480156105c957600080fd5b506105dd6105d8366004613589565b612af0565b6040516102af91906146cf565b3480156105f657600080fd5b5061032d612b6e565b34801561060b57600080fd5b506102a261061a366004613966565b612b7d565b34801561062b57600080fd5b5061026061063a366004613589565b612b92565b34801561064b57600080fd5b506102a261065a366004613966565b612c2d565b34801561066b57600080fd5b506103ac61067a3660046135d2565b612c94565b34801561068b57600080fd5b5061032d612dc6565b3480156106a057600080fd5b5061032d612dd5565b6000546001600160a01b031633146106dc5760405162461bcd60e51b81526004016106d3906145db565b60405180910390fd5b600154600160a01b900460ff166107055760405162461bcd60e51b81526004016106d390614040565b600b546001600160a01b031661072d5760405162461bcd60e51b81526004016106d3906144d5565b606081516001600160401b038111801561074657600080fd5b50604051908082528060200260200182016040528015610770578160200160208202803683370190505b50905060005b825181101561080e576007600084838151811061078f57fe5b60200260200101516001600160a01b03166001600160a01b03168152602001908152602001600020548282815181106107c457fe5b6020026020010181815250506000600760008584815181106107e257fe5b6020908102919091018101516001600160a01b0316825281019190915260400160002055600101610776565b50600b54604051631ed4276d60e01b81526001600160a01b0390911690631ed4276d906108419085908590600401613f36565b602060405180830381600087803b15801561085b57600080fd5b505af115801561086f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610893919061394a565b6108af5760405162461bcd60e51b81526004016106d3906140f0565b7f492fc8b292f2a2a9b328a366b83745f30c024056d12aa118a15966d26a8ce658826040516108de9190613f23565b60405180910390a15050565b600154600160a01b900460ff16156109145760405162461bcd60e51b81526004016106d3906145b1565b600b54600160a01b900460ff1661093d5760405162461bcd60e51b81526004016106d39061406e565b600b805460ff60a01b19169055610952613282565b61095b8b611ac5565b9050806000015160ff1660f114801561097b5750806020015160ff166001145b61098457600080fd5b6109918160a00151611a00565b156109ae5760405162461bcd60e51b81526004016106d390614651565b60a081015160009081526003602052604090819020805460ff191660011790558101516001600160a01b0316610a3d576040808201516001600160a01b03166000908152600760205220546080820151610a0d9163ffffffff612de416565b6008546001600160a01b0316311015610a385760405162461bcd60e51b81526004016106d3906142ce565b610b44565b6000610a4c8260400151612af0565b905060098160ff161115610a6f5760808201805160081960ff841601600a0a0290525b6040808301516001600160a01b03166000908152600760205220546080830151610a9e9163ffffffff612de416565b60408084015160085491516370a0823160e01b81526001600160a01b03918216926370a0823192610ad492911690600401613def565b60206040518083038186803b158015610aec57600080fd5b505afa158015610b00573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b249190613bc5565b1015610b425760405162461bcd60e51b81526004016106d3906142ce565b505b610b568b8b8b8b8b8b8b8b8b8b612e02565b600854606082015160408084015160808501519151633b22ca0f60e21b81526001600160a01b039094169363ec8b283c93610b9693909291600401613e03565b600060405180830381600087803b158015610bb057600080fd5b505af1158015610bc4573d6000803e3d6000fd5b505050507f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb816040015182606001518360800151604051610c0793929190613e03565b60405180910390a15050600b805460ff60a01b1916600160a01b179055505050505050505050565b60046020526000908152604090205460ff1681565b600a546000906001600160a01b031615801590610c6b5750600a546001600160a01b031633145b610c875760405162461bcd60e51b81526004016106d3906143cf565b838214610ca65760405162461bcd60e51b81526004016106d3906141ec565b600a60009054906101000a90046001600160a01b03166001600160a01b0316635c975abb6040518163ffffffff1660e01b815260040160206040518083038186803b158015610cf457600080fd5b505afa158015610d08573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d2c919061394a565b610d485760405162461bcd60e51b81526004016106d39061427f565b6008546001600160a01b0316610d705760405162461bcd60e51b81526004016106d390614223565b60005b8481101561106e57610ddc848483818110610d8a57fe5b9050602002013560076000898986818110610da157fe5b9050602002016020810190610db69190613589565b6001600160a01b031681526020810191909152604001600020549063ffffffff612de416565b60076000888885818110610dec57fe5b9050602002016020810190610e019190613589565b6001600160a01b031681526020810191909152604001600090812091909155868683818110610e2c57fe5b9050602002016020810190610e419190613589565b6001600160a01b0316148015610e575750600047115b15610ee3576008546040516000916001600160a01b0316904790610e7a90613dec565b60006040518083038185875af1925050503d8060008114610eb7576040519150601f19603f3d011682016040523d82523d6000602084013e610ebc565b606091505b5050905080610edd5760405162461bcd60e51b81526004016106d390614093565b50611066565b6000868683818110610ef157fe5b9050602002016020810190610f069190613589565b6001600160a01b031614611066576000868683818110610f2257fe5b9050602002016020810190610f379190613589565b6001600160a01b03166370a08231306040518263ffffffff1660e01b8152600401610f629190613def565b60206040518083038186803b158015610f7a57600080fd5b505afa158015610f8e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fb29190613bc5565b9050801561106457868683818110610fc657fe5b9050602002016020810190610fdb9190613589565b60085460405163a9059cbb60e01b81526001600160a01b039283169263a9059cbb9261100e929116908590600401613e8b565b600060405180830381600087803b15801561102857600080fd5b505af115801561103c573d6000803e3d6000fd5b50505050611048612eec565b6110645760405162461bcd60e51b81526004016106d390614558565b505b600101610d73565b507f6a7fbbcddfd518bb8c56b28ac6c7acb0f7ca093ed232eb3306e53d14e469895f858585856040516110a49493929190613ea4565b60405180910390a1506001949350505050565b6000546001600160a01b031633146110e15760405162461bcd60e51b81526004016106d3906145db565b600154600160a01b900460ff1661110a5760405162461bcd60e51b81526004016106d390614040565b6001600160a01b0381166111305760405162461bcd60e51b81526004016106d39061442d565b600980546001600160a01b0319166001600160a01b0383161790556040517f204252dfe190ad6ef63db40a490f048b39f661de74628408f13cd0bb2d4c34469061117b908390613def565b60405180910390a150565b6000546001600160a01b031633146111b05760405162461bcd60e51b81526004016106d3906145db565b600154600160a01b900460ff166111d95760405162461bcd60e51b81526004016106d390614040565b6001805460ff60a01b191690556040517f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa90611216903390613def565b60405180910390a1565b60008060008060208601519150604086015192508560408151811061124157fe5b602001015160f81c60f81b60f81c601b019050600185828486604051600081526020016040526040516112779493929190614022565b6020604051602081039080840390855afa158015611299573d6000803e3d6000fd5b5050506020604051035193505050505b92915050565b60025442106112d05760405162461bcd60e51b81526004016106d3906142ad565b6001546001600160a01b031633146112fa5760405162461bcd60e51b81526004016106d3906143a9565b600154600080546001600160a01b0319166001600160a01b0392831617908190556040517f0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc92611216921690613def565b600081565b600154600160a01b900460ff161561137a5760405162461bcd60e51b81526004016106d3906145b1565b600b54600160a01b900460ff166113a35760405162461bcd60e51b81526004016106d39061406e565b600b805460ff60a01b191690558260006113bc82612af0565b6008546040516370a0823160e01b81529192506000916001600160a01b03858116926370a08231926113f49290911690600401613def565b60206040518083038186803b15801561140c57600080fd5b505afa158015611420573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114449190613bc5565b90508085600960ff8516111561147f5760098460ff1603600a0a818161146657fe5b04905060098460ff1603600a0a838161147b57fe5b0492505b670de0b6b3a7640000811115801561149f5750670de0b6b3a76400008311155b80156114c25750670de0b6b3a76400006114bf828563ffffffff612de416565b11155b6114de5760405162461bcd60e51b81526004016106d3906145fe565b6008546040516323b872dd60e01b81526001600160a01b03808816926323b872dd9261151292339216908c90600401613e03565b600060405180830381600087803b15801561152c57600080fd5b505af1158015611540573d6000803e3d6000fd5b5050505061154c612eec565b6115685760405162461bcd60e51b81526004016106d3906142f8565b6008546040516370a0823160e01b815288916115fc9185916001600160a01b03808b16926370a08231926115a0921690600401613def565b60206040518083038186803b1580156115b857600080fd5b505afa1580156115cc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115f09190613bc5565b9063ffffffff612f2016565b146116195760405162461bcd60e51b81526004016106d3906140c1565b7f2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e88878360405161164c93929190613e41565b60405180910390a15050600b805460ff60a01b1916600160a01b179055505050505050565b600154600160a01b900460ff1681565b60076020526000908152604090205481565b600560209081526000928352604080842090915290825290205481565b6001546001600160a01b031681565b600154600160a01b900460ff16156116e95760405162461bcd60e51b81526004016106d3906145b1565b600b54600160a01b900460ff166117125760405162461bcd60e51b81526004016106d39061406e565b600b805460ff60a01b19169055611727613282565b6117308b611ac5565b9050806000015160ff1660f31480156117505750806020015160ff166001145b61175957600080fd5b6117668160a00151611a00565b156117835760405162461bcd60e51b81526004016106d390614325565b60a081015160009081526003602052604090819020805460ff191660011790558101516001600160a01b0316611812576040808201516001600160a01b031660009081526007602052205460808201516117e29163ffffffff612de416565b6008546001600160a01b031631101561180d5760405162461bcd60e51b81526004016106d390614629565b611919565b60006118218260400151612af0565b905060098160ff1611156118445760808201805160081960ff841601600a0a0290525b6040808301516001600160a01b031660009081526007602052205460808301516118739163ffffffff612de416565b60408084015160085491516370a0823160e01b81526001600160a01b03918216926370a08231926118a992911690600401613def565b60206040518083038186803b1580156118c157600080fd5b505afa1580156118d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118f99190613bc5565b10156119175760405162461bcd60e51b81526004016106d390614464565b505b61192b8b8b8b8b8b8b8b8b8b8b612e02565b608081015160608201516001600160a01b039081166000908152600560209081526040808320818701519094168352929052205461196e9163ffffffff612de416565b60608201516001600160a01b039081166000908152600560209081526040808320818701805186168552908352818420959095556080860151945190931682526007905220546119c39163ffffffff612de416565b6040918201516001600160a01b03166000908152600760205291909120555050600b805460ff60a01b1916600160a01b1790555050505050505050565b60008181526003602052604081205460ff1615611a1f57506001611aba565b600a546001600160a01b0316611a3757506000611aba565b600a54604051633a4e2fc360e11b81526001600160a01b039091169063749c5f8690611a67908590600401614019565b60206040518083038186803b158015611a7f57600080fd5b505afa158015611a93573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ab7919061394a565b90505b919050565b60025481565b611acd613282565b611ad5613282565b82600081518110611ae257fe5b016020015160f81c8152825183906001908110611afb57fe5b0160209081015160f81c9082015260228301516042840151606285015160828601516001600160a01b039384166040860152929091166060840152608083015260a08201529050919050565b6000546001600160a01b03163314611b715760405162461bcd60e51b81526004016106d3906145db565b600154600160a01b900460ff1615611b9b5760405162461bcd60e51b81526004016106d3906145b1565b6002544210611bbc5760405162461bcd60e51b81526004016106d3906142ad565b6001805460ff60a01b1916600160a01b1790556040517f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25890611216903390613def565b600154600160a01b900460ff1615611c295760405162461bcd60e51b81526004016106d3906145b1565b600b54600160a01b900460ff16611c525760405162461bcd60e51b81526004016106d39061406e565b600b805460ff60a01b19169055604051600090611c9290611c7d908790879087908c90602001613d24565b60405160208183030381529060405283612f35565b9050611c9e8189612f94565b6001600160a01b038082166000908152600560209081526040808320938c1683529290522054871115611ce35760405162461bcd60e51b81526004016106d39061418b565b6001600160a01b038816600090815260076020526040902054611d0c908863ffffffff612f2016565b6001600160a01b03808a1660008181526007602090815260408083209590955592851681526005835283812091815291522054611d4f908863ffffffff612f2016565b6001600160a01b0380831660009081526005602090815260408083208d8516845290915290819020929092556008549151633b22ca0f60e21b815291169063ec8b283c90611da59088908c908c90600401613e03565b600060405180830381600087803b158015611dbf57600080fd5b505af1158015611dd3573d6000803e3d6000fd5b505050506000611de48786886130dc565b6001600160a01b038084166000908152600560209081526040808320938c1683529290522054909150611e1d908263ffffffff612de416565b6001600160a01b038084166000908152600560209081526040808320938c168352928152828220939093556007909252902054611e60908263ffffffff612de416565b6001600160a01b039097166000908152600760205260409020969096555050600b805460ff60a01b1916600160a01b179055505050505050565b600154600160a01b900460ff1615611ec45760405162461bcd60e51b81526004016106d3906145b1565b600b54600160a01b900460ff16611eed5760405162461bcd60e51b81526004016106d39061406e565b600b805460ff60a01b19169055604051600090611f2d90611f18908890889086908990602001613db2565b60405160208183030381529060405284612f35565b9050611f398186612f94565b6001600160a01b03808216600090815260056020908152604080832093891683529290522054841115611f7e5760405162461bcd60e51b81526004016106d39061418b565b6001600160a01b03808216600090815260056020908152604080832093891683529290522054611fb4908563ffffffff612f2016565b6001600160a01b038083166000908152600560209081526040808320938a168352928152828220939093556007909252902054611ff7908563ffffffff612f2016565b6001600160a01b03861660008181526007602052604090209190915584901561204b57600061202587612af0565b905060098160ff1611156120495760098160ff1603600a0a868161204557fe5b0491505b505b7f2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e86888360405161207e93929190613e41565b60405180910390a15050600b805460ff60a01b1916600160a01b1790555050505050565b600b546001600160a01b031681565b6009546001600160a01b031681565b6000546001600160a01b031633146120ea5760405162461bcd60e51b81526004016106d3906145db565b600254421061210b5760405162461bcd60e51b81526004016106d3906142ad565b61016e811061212c5760405162461bcd60e51b81526004016106d390614347565b600280546201518083020190556040517f02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e89061117b908390614019565b600660209081526000928352604080842090915290825290205460ff1681565b60006001600160a01b0382166121ac57506008546001600160a01b031631611aba565b6008546040516370a0823160e01b81526001600160a01b03848116926370a08231926121de9290911690600401613def565b60206040518083038186803b1580156121f657600080fd5b505afa15801561220a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ab79190613bc5565b6000546001600160a01b031633146122585760405162461bcd60e51b81526004016106d3906145db565b60025442106122795760405162461bcd60e51b81526004016106d3906142ad565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b600154600160a01b900460ff16156122c55760405162461bcd60e51b81526004016106d3906145b1565b600b54600160a01b900460ff166122ee5760405162461bcd60e51b81526004016106d39061406e565b600b805460ff60a01b191690556008546040516000916001600160a01b031690349061231990613dec565b60006040518083038185875af1925050503d8060008114612356576040519150601f19603f3d011682016040523d82523d6000602084013e61235b565b606091505b505090508061237c5760405162461bcd60e51b81526004016106d390614253565b6008546b033b2e3c9fd0803ce80000006001600160a01b039091163111156123b65760405162461bcd60e51b81526004016106d390614588565b7f2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e600083346040516123ea93929190613e41565b60405180910390a15050600b805460ff60a01b1916600160a01b179055565b600b54600160a01b900460ff1681565b600154600160a01b900460ff16156124435760405162461bcd60e51b81526004016106d3906145b1565b600b54600160a01b900460ff1661246c5760405162461bcd60e51b81526004016106d39061406e565b600b805460ff60a01b191690558551875114801561248b575060008551115b6124a75760405162461bcd60e51b81526004016106d3906141b5565b60006124c38585858a604051602001611c7d9493929190613cb2565b905060005b885181101561275e576124ee828a83815181106124e157fe5b6020026020010151612f94565b8781815181106124fa57fe5b602002602001015160056000846001600160a01b03166001600160a01b0316815260200190815260200160002060008b848151811061253557fe5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054101561257c5760405162461bcd60e51b81526004016106d39061418b565b6125da88828151811061258b57fe5b6020026020010151600760008c85815181106125a357fe5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054612f2090919063ffffffff16565b600760008b84815181106125ea57fe5b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000208190555061266088828151811061262557fe5b602002602001015160056000856001600160a01b03166001600160a01b0316815260200190815260200160002060008c85815181106125a357fe5b6001600160a01b03831660009081526005602052604081208b519091908c908590811061268957fe5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002081905550600860009054906101000a90046001600160a01b03166001600160a01b031663ec8b283c878b84815181106126e657fe5b60200260200101518b85815181106126fa57fe5b60200260200101516040518463ffffffff1660e01b815260040161272093929190613e03565b600060405180830381600087803b15801561273a57600080fd5b505af115801561274e573d6000803e3d6000fd5b5050600190920191506124c89050565b50606086516001600160401b038111801561277857600080fd5b506040519080825280602002602001820160405280156127a2578160200160208202803683370190505b50905060005b87518110156127ec576127cd8882815181106127c057fe5b6020026020010151612189565b8282815181106127d957fe5b60209081029190910101526001016127a8565b506060806127fa87896131db565b9150915088518251148015612810575081518151145b61282c5760405162461bcd60e51b81526004016106d39061448e565b60005b8151811015612a0b5789818151811061284457fe5b60200260200101516001600160a01b031683828151811061286157fe5b60200260200101516001600160a01b03161480156128b6575081818151811061288657fe5b60200260200101516128b485838151811061289d57fe5b60200260200101516115f08d85815181106127c057fe5b145b6128d25760405162461bcd60e51b81526004016106d3906143f6565b6129538282815181106128e157fe5b602002602001015160056000886001600160a01b03166001600160a01b0316815260200190815260200160002060008d858151811061291c57fe5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054612de490919063ffffffff16565b6001600160a01b03861660009081526005602052604081208c519091908d908590811061297c57fe5b60200260200101516001600160a01b03166001600160a01b03168152602001908152602001600020819055506129cf8282815181106129b757fe5b6020026020010151600760008d858151811061291c57fe5b600760008c84815181106129df57fe5b6020908102919091018101516001600160a01b031682528101919091526040016000205560010161282f565b5050600b805460ff60a01b1916600160a01b17905550505050505050505050565b6000546001600160a01b03163314612a565760405162461bcd60e51b81526004016106d3906145db565b600154600160a01b900460ff16612a7f5760405162461bcd60e51b81526004016106d390614040565b6001600160a01b038116612aa55760405162461bcd60e51b81526004016106d3906144d5565b600b80546001600160a01b0319166001600160a01b0383161790556040517fd58a618a39de682696ea37dd9a6bf9c793afa426fa1438e75c3966e3b541e45a9061117b908390613def565b600080829050806001600160a01b031663313ce5676040518163ffffffff1660e01b815260040160206040518083038186803b158015612b2f57600080fd5b505afa158015612b43573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b679190613bc5565b9392505050565b6008546001600160a01b031681565b60036020526000908152604090205460ff1681565b6000546001600160a01b03163314612bbc5760405162461bcd60e51b81526004016106d3906145db565b600154600160a01b900460ff16612be55760405162461bcd60e51b81526004016106d390614040565b6001600160a01b038116612c0b5760405162461bcd60e51b81526004016106d39061437e565b600880546001600160a01b0319166001600160a01b0392909216919091179055565b60008181526004602052604081205460ff1615612c4c57506001611aba565b600a546001600160a01b0316612c6457506000611aba565b600a5460405163392f5c1d60e21b81526001600160a01b039091169063e4bd707490611a67908590600401614019565b600a546000906001600160a01b031615801590612cd757506001600160a01b0380831660009081526006602090815260408083209387168352929052205460ff16155b15612d9957600a54604051637badcc6760e11b8152612d92916001600160a01b03169063f75b98ce90612d109087908790600401613e27565b60206040518083038186803b158015612d2857600080fd5b505afa158015612d3c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d609190613bc5565b6001600160a01b038085166000908152600560209081526040808320938916835292905220549063ffffffff612de416565b90506112a9565b506001600160a01b0380821660009081526005602090815260408083209386168352929052205492915050565b6000546001600160a01b031681565b600a546001600160a01b031681565b6000828201838110801590612df95750828110155b612b6757600080fd5b60008a8a604051602001612e17929190613d90565b60408051601f19818403018152908290528051602090910120600954637b2e908b60e11b83529092506001600160a01b03169063f65d211690612e739060019085908f908f908f908f908f908f908f908f908f90600401613f6f565b60206040518083038186803b158015612e8b57600080fd5b505afa158015612e9f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ec3919061394a565b612edf5760405162461bcd60e51b81526004016106d39061411d565b5050505050505050505050565b6000803d8015612f035760208114612f0c57612f18565b60019150612f18565b60206000803e60005191505b501515905090565b600082821115612f2f57600080fd5b50900390565b81516020830120600090612f4881612c2d565b15612f655760405162461bcd60e51b81526004016106d390614536565b6000612f718483611220565b600092835260046020526040909220805460ff1916600117905550905092915050565b600a546001600160a01b031615801590612fd457506001600160a01b0380831660009081526006602090815260408083209385168352929052205460ff16155b156130d857600a54604051637badcc6760e11b815261308f916001600160a01b03169063f75b98ce9061300d9085908790600401613e27565b60206040518083038186803b15801561302557600080fd5b505afa158015613039573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061305d9190613bc5565b6001600160a01b038085166000908152600560209081526040808320938716835292905220549063ffffffff612de416565b6001600160a01b038084166000818152600560209081526040808320948716808452948252808320959095559181526006825283812092815291905220805460ff191660011790555b5050565b6000806130e885612189565b905060006060846001600160a01b03166000876040516131089190613d74565b60006040518083038185875af1925050503d8060008114613145576040519150601f19603f3d011682016040523d82523d6000602084013e61314a565b606091505b50915091508161316c5760405162461bcd60e51b81526004016106d3906144ff565b6000808280602001905181019061318391906135a5565b91509150886001600160a01b0316826001600160a01b03161480156131b35750806131b1866115f08c612189565b145b6131cf5760405162461bcd60e51b81526004016106d3906143f6565b98975050505050505050565b60608060006060846001600160a01b03166000876040516131fc9190613d74565b60006040518083038185875af1925050503d8060008114613239576040519150601f19603f3d011682016040523d82523d6000602084013e61323e565b606091505b5091509150816132605760405162461bcd60e51b81526004016106d390614154565b8080602001905181019061327491906137c6565b9350935050505b9250929050565b6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a081019190915290565b80356112a981614752565b80516112a981614752565b60008083601f8401126132de578182fd5b5081356001600160401b038111156132f4578182fd5b602083019150836020808302850101111561327b57600080fd5b600082601f83011261331e578081fd5b813561333161332c82614703565b6146dd565b81815291506020808301908481018184028601820187101561335257600080fd5b60005b8481101561337a57813561336881614752565b84529282019290820190600101613355565b505050505092915050565b600082601f830112613395578081fd5b81356133a361332c82614703565b8181529150602080830190848101818402860182018710156133c457600080fd5b60005b8481101561337a5781356133da8161476a565b845292820192908201906001016133c7565b600082601f8301126133fc578081fd5b813561340a61332c82614703565b81815291506020808301908481018184028601820187101561342b57600080fd5b60005b8481101561337a5781358452928201929082019060010161342e565b600082601f83011261345a578081fd5b815161346861332c82614703565b81815291506020808301908481018184028601820187101561348957600080fd5b60005b8481101561337a5781518452928201929082019060010161348c565b600082601f8301126134b8578081fd5b81356134c661332c82614703565b8181529150602080830190848101818402860182018710156134e757600080fd5b6000805b8581101561351557823560ff81168114613503578283fd5b855293830193918301916001016134eb565b50505050505092915050565b600082601f830112613531578081fd5b81356001600160401b03811115613546578182fd5b613559601f8201601f19166020016146dd565b915080825283602082850101111561357057600080fd5b8060208401602084013760009082016020015292915050565b60006020828403121561359a578081fd5b8135612b6781614752565b600080604083850312156135b7578081fd5b82516135c281614752565b6020939093015192949293505050565b600080604083850312156135e4578182fd5b82356135ef81614752565b915060208301356135ff81614752565b809150509250929050565b600080600080600080600060e0888a031215613624578283fd5b873561362f81614752565b965060208801359550604088013561364681614752565b94506136558960608a016132b7565b935060808801356001600160401b0380821115613670578485fd5b61367c8b838c01613521565b945060a08a0135915080821115613691578384fd5b61369d8b838c01613521565b935060c08a01359150808211156136b2578283fd5b506136bf8a828b01613521565b91505092959891949750929550565b6000806000606084860312156136e2578081fd5b83356136ed81614752565b92506020840135915060408401356001600160401b0381111561370e578182fd5b61371a86828701613521565b9150509250925092565b60008060008060408587031215613739578182fd5b84356001600160401b038082111561374f578384fd5b61375b888389016132cd565b90965094506020870135915080821115613773578384fd5b50613780878288016132cd565b95989497509550505050565b60006020828403121561379d578081fd5b81356001600160401b038111156137b2578182fd5b6137be8482850161330e565b949350505050565b600080604083850312156137d8578182fd5b82516001600160401b03808211156137ee578384fd5b81850186601f8201126137ff578485fd5b8051925061380f61332c84614703565b80848252602080830192508084018a82838902870101111561382f578889fd5b8894505b86851015613859576138458b826132c2565b845260019490940193928101928101613833565b508801519096509350505080821115613870578283fd5b5061387d8582860161344a565b9150509250929050565b600080600080600080600060e0888a0312156138a1578081fd5b87356001600160401b03808211156138b7578283fd5b6138c38b838c0161330e565b985060208a01359150808211156138d8578283fd5b6138e48b838c016133ec565b975060408a01359150808211156138f9578283fd5b6139058b838c0161330e565b96506139148b60608c016132b7565b955060808a0135915080821115613929578283fd5b6139358b838c01613521565b945060a08a0135915080821115613691578283fd5b60006020828403121561395b578081fd5b8151612b678161476a565b600060208284031215613977578081fd5b5035919050565b60006020828403121561398f578081fd5b81356001600160401b038111156139a4578182fd5b6137be84828501613521565b600080604083850312156139c2578182fd5b82356001600160401b038111156139d7578283fd5b6139e385828601613521565b95602094909401359450505050565b6000806000806000806000806000806101408b8d031215613a11578384fd5b8a356001600160401b0380821115613a27578586fd5b613a338e838f01613521565b9b5060208d01359a5060408d0135915080821115613a4f578586fd5b613a5b8e838f016133ec565b995060608d0135915080821115613a70578586fd5b613a7c8e838f01613385565b985060808d0135975060a08d0135965060c08d0135915080821115613a9f578586fd5b613aab8e838f016133ec565b955060e08d0135915080821115613ac0578485fd5b613acc8e838f016134a8565b94506101008d0135915080821115613ae2578384fd5b613aee8e838f016133ec565b93506101208d0135915080821115613b04578283fd5b50613b118d828e016133ec565b9150509295989b9194979a5092959850565b600080600080600060a08688031215613b3a578283fd5b85356001600160401b0380821115613b50578485fd5b613b5c89838a01613521565b965060208801359150613b6e82614752565b9094506040870135935060608701359080821115613b8a578283fd5b613b9689838a01613521565b93506080880135915080821115613bab578283fd5b50613bb888828901613521565b9150509295509295909350565b600060208284031215613bd6578081fd5b5051919050565b6000815180845260208085019450808401835b83811015613c155781516001600160a01b031687529582019590820190600101613bf0565b509495945050505050565b6000815180845260208085019450808401835b83811015613c15578151151587529582019590820190600101613c33565b6000815180845260208085019450808401835b83811015613c1557815187529582019590820190600101613c64565b6000815180845260208085019450808401835b83811015613c1557815160ff1687529582019590820190600101613c93565b60006001600160601b03198660601b16825284516020613cd88260148601838a01614722565b8551918401613ced8360148301848a01614722565b8551928287019101601401845b84811015613d1657825182529183019190830190600101613cfa565b509998505050505050505050565b60006001600160601b03198660601b1682528451613d49816014850160208901614722565b8451908301613d5f826014830160208901614722565b01601481019390935250506034019392505050565b60008251613d86818460208701614722565b9190910192915050565b60008351613da2818460208801614722565b9190910191825250602001919050565b60008551613dc4818460208a01614722565b8083016001600160601b03198760601b16815285519150613d5f826014830160208901614722565b90565b6001600160a01b0391909116815260200190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b6001600160a01b0392831681529116602082015260400190565b600060018060a01b0385168252606060208301528351806060840152613e6e816080850160208801614722565b604083019390935250601f91909101601f19160160800192915050565b6001600160a01b03929092168252602082015260400190565b6040808252810184905260008560608301825b87811015613ee75760208335613ecc81614752565b6001600160a01b031683529283019290910190600101613eb7565b5083810360208501528481526001600160fb1b03851115613f06578283fd5b602085029150818660208301370160200190815295945050505050565b600060208252612b676020830184613bdd565b600060408252613f496040830185613bdd565b8281036020840152613f5b8185613c51565b95945050505050565b901515815260200190565b60006101608d151583528c60208401528b6040840152806060840152613f978184018c613c51565b8381036080850152613fa9818c613c20565b9150508860a08401528760c084015282810360e0840152613fca8188613c51565b838103610100850152613fdd8188613c80565b915050828103610120840152613ff38186613c51565b8381036101408501526140068186613c51565b9f9e505050505050505050505050505050565b90815260200190565b93845260ff9290921660208401526040830152606082015260800190565b6020808252601490820152736e6f7420706175736564207269676874206e6f7760601b604082015260600190565b6020808252600b908201526a1c99595b9d1c985b9d195960aa1b604082015260600190565b6020808252601490820152730ccc2d2d8cac840dad2cee4c2e8d2dcce40cae8d60631b604082015260600190565b6020808252601590820152741b995dc818985b185b98d9481a5b98dbdc9c9958dd605a1b604082015260600190565b6020808252601390820152726661696c65642075706461746541737365747360681b604082015260600190565b60208082526018908201527f696e737472756374696f6e206e6f7420617070726f7665640000000000000000604082015260600190565b60208082526017908201527f6661696c65642063616c6c696e672065786368616e6765000000000000000000604082015260600190565b60208082526010908201526f3737ba1032b737bab3b4103a37b5b2b760811b604082015260600190565b60208082526019908201527f6d69736d6174636820746f6b656e20616e6420616d6f756e7400000000000000604082015260600190565b6020808252601b908201527f6d69736d617463682061737365747320616e6420616d6f756e74730000000000604082015260600190565b6020808252601690820152751b1bd8dad95c881859191c995cdcc81a5cc81b9d5b1b60521b604082015260600190565b6020808252601290820152710ccc2d2d8cac840c8cae0dee6d2e840cae8d60731b604082015260600190565b6020808252601490820152731c1c995d95985d5b1d081b9bdd081c185d5cd95960621b604082015260600190565b602080825260079082015266195e1c1a5c995960ca1b604082015260600190565b60208082526010908201526f0e8ded6cadc40dcdee840cadcdeeaced60831b604082015260600190565b6020808252601390820152726661696c6564207472616e7366657246726f6d60681b604082015260600190565b6020808252600890820152671a5d1e081d5cd95960c21b604082015260600190565b6020808252601a908201527f63616e6e6f7420657874656e6420666f7220746f6f206c6f6e67000000000000604082015260600190565b60208082526011908201527034b73b30b634b2103732bba637b1b5b2b960791b604082015260600190565b6020808252600c908201526b1d5b985d5d1a1bdc9a5e995960a21b604082015260600190565b6020808252600d908201526c1b9bdd081c1c995d95985d5b1d609a1b604082015260600190565b60208082526017908201527f72657475726e20616d6f756e7420696e636f7272656374000000000000000000604082015260600190565b60208082526019908201527f696e76616c6964206e6577496e636f676e69746f50726f787900000000000000604082015260600190565b60208082526010908201526f06e6f7420656e6f7567682065726332360841b604082015260600190565b60208082526027908201527f6d69736d617463682072657475726e20746f6b656e20616e642072657475726e60408201526608185b5bdd5b9d60ca1b606082015260800190565b60208082526010908201526f1a5b9d985b1a59081b995dd5985d5b1d60821b604082015260600190565b60208082526017908201527f6661696c65642063616c6c696e67206578742066756e63000000000000000000604082015260600190565b6020808252600890820152671cda59c81d5cd95960c21b604082015260600190565b6020808252601690820152753330b4b632b21036b4b3b930ba34b733903a37b5b2b760511b604082015260600190565b6020808252600f908201526e1b585e08195d1a081c995858da1959608a1b604082015260600190565b60208082526010908201526f706175736564207269676874206e6f7760801b604082015260600190565b6020808252600990820152683737ba1030b236b4b760b91b604082015260600190565b6020808252601190820152701b585e08195c98cc8c081c995858da1959607a1b604082015260600190565b6020808252600e908201526d0dcdee840cadcdeeaced040cae8d60931b604082015260600190565b6020808252600c908201526b185b1c9958591e481d5cd95960a21b604082015260600190565b600060c08201905060ff835116825260ff6020840151166020830152604083015160018060a01b03808216604085015280606086015116606085015250506080830151608083015260a083015160a083015292915050565b60ff91909116815260200190565b6040518181016001600160401b03811182821017156146fb57600080fd5b604052919050565b60006001600160401b03821115614718578081fd5b5060209081020190565b60005b8381101561473d578181015183820152602001614725565b8381111561474c576000848401525b50505050565b6001600160a01b038116811461476757600080fd5b50565b801515811461476757600080fdfea26469706673582212201fa0b2107fbd92a8ad288933480def9e65389f7648cdd8cbbb7b35c8e02b445564736f6c63430006060033"

// DeployVault deploys a new Ethereum contract, binding an instance of Vault to it.
func DeployVault(auth *bind.TransactOpts, backend bind.ContractBackend, admin common.Address, incognitoProxyAddress common.Address, _prevVault common.Address, _locker common.Address) (common.Address, *types.Transaction, *Vault, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VaultBin), backend, admin, incognitoProxyAddress, _prevVault, _locker)
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
// Solidity: function ETH_TOKEN() view returns(address)
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
// Solidity: function ETH_TOKEN() view returns(address)
func (_Vault *VaultSession) ETHTOKEN() (common.Address, error) {
	return _Vault.Contract.ETHTOKEN(&_Vault.CallOpts)
}

// ETHTOKEN is a free data retrieval call binding the contract method 0x58bc8337.
//
// Solidity: function ETH_TOKEN() view returns(address)
func (_Vault *VaultCallerSession) ETHTOKEN() (common.Address, error) {
	return _Vault.Contract.ETHTOKEN(&_Vault.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
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
// Solidity: function admin() view returns(address)
func (_Vault *VaultSession) Admin() (common.Address, error) {
	return _Vault.Contract.Admin(&_Vault.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Vault *VaultCallerSession) Admin() (common.Address, error) {
	return _Vault.Contract.Admin(&_Vault.CallOpts)
}

// BalanceOfLocker is a free data retrieval call binding the contract method 0x9b780b2a.
//
// Solidity: function balanceOfLocker(address token) view returns(uint256)
func (_Vault *VaultCaller) BalanceOfLocker(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "balanceOfLocker", token)
	return *ret0, err
}

// BalanceOfLocker is a free data retrieval call binding the contract method 0x9b780b2a.
//
// Solidity: function balanceOfLocker(address token) view returns(uint256)
func (_Vault *VaultSession) BalanceOfLocker(token common.Address) (*big.Int, error) {
	return _Vault.Contract.BalanceOfLocker(&_Vault.CallOpts, token)
}

// BalanceOfLocker is a free data retrieval call binding the contract method 0x9b780b2a.
//
// Solidity: function balanceOfLocker(address token) view returns(uint256)
func (_Vault *VaultCallerSession) BalanceOfLocker(token common.Address) (*big.Int, error) {
	return _Vault.Contract.BalanceOfLocker(&_Vault.CallOpts, token)
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() view returns(uint256)
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
// Solidity: function expire() view returns(uint256)
func (_Vault *VaultSession) Expire() (*big.Int, error) {
	return _Vault.Contract.Expire(&_Vault.CallOpts)
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() view returns(uint256)
func (_Vault *VaultCallerSession) Expire() (*big.Int, error) {
	return _Vault.Contract.Expire(&_Vault.CallOpts)
}

// GetDecimals is a free data retrieval call binding the contract method 0xcf54aaa0.
//
// Solidity: function getDecimals(address token) view returns(uint8)
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
// Solidity: function getDecimals(address token) view returns(uint8)
func (_Vault *VaultSession) GetDecimals(token common.Address) (uint8, error) {
	return _Vault.Contract.GetDecimals(&_Vault.CallOpts, token)
}

// GetDecimals is a free data retrieval call binding the contract method 0xcf54aaa0.
//
// Solidity: function getDecimals(address token) view returns(uint8)
func (_Vault *VaultCallerSession) GetDecimals(token common.Address) (uint8, error) {
	return _Vault.Contract.GetDecimals(&_Vault.CallOpts, token)
}

// GetDepositedBalance is a free data retrieval call binding the contract method 0xf75b98ce.
//
// Solidity: function getDepositedBalance(address token, address owner) view returns(uint256)
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
// Solidity: function getDepositedBalance(address token, address owner) view returns(uint256)
func (_Vault *VaultSession) GetDepositedBalance(token common.Address, owner common.Address) (*big.Int, error) {
	return _Vault.Contract.GetDepositedBalance(&_Vault.CallOpts, token, owner)
}

// GetDepositedBalance is a free data retrieval call binding the contract method 0xf75b98ce.
//
// Solidity: function getDepositedBalance(address token, address owner) view returns(uint256)
func (_Vault *VaultCallerSession) GetDepositedBalance(token common.Address, owner common.Address) (*big.Int, error) {
	return _Vault.Contract.GetDepositedBalance(&_Vault.CallOpts, token, owner)
}

// Incognito is a free data retrieval call binding the contract method 0x8a984538.
//
// Solidity: function incognito() view returns(address)
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
// Solidity: function incognito() view returns(address)
func (_Vault *VaultSession) Incognito() (common.Address, error) {
	return _Vault.Contract.Incognito(&_Vault.CallOpts)
}

// Incognito is a free data retrieval call binding the contract method 0x8a984538.
//
// Solidity: function incognito() view returns(address)
func (_Vault *VaultCallerSession) Incognito() (common.Address, error) {
	return _Vault.Contract.Incognito(&_Vault.CallOpts)
}

// IsSigDataUsed is a free data retrieval call binding the contract method 0xe4bd7074.
//
// Solidity: function isSigDataUsed(bytes32 hash) view returns(bool)
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
// Solidity: function isSigDataUsed(bytes32 hash) view returns(bool)
func (_Vault *VaultSession) IsSigDataUsed(hash [32]byte) (bool, error) {
	return _Vault.Contract.IsSigDataUsed(&_Vault.CallOpts, hash)
}

// IsSigDataUsed is a free data retrieval call binding the contract method 0xe4bd7074.
//
// Solidity: function isSigDataUsed(bytes32 hash) view returns(bool)
func (_Vault *VaultCallerSession) IsSigDataUsed(hash [32]byte) (bool, error) {
	return _Vault.Contract.IsSigDataUsed(&_Vault.CallOpts, hash)
}

// IsWithdrawed is a free data retrieval call binding the contract method 0x749c5f86.
//
// Solidity: function isWithdrawed(bytes32 hash) view returns(bool)
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
// Solidity: function isWithdrawed(bytes32 hash) view returns(bool)
func (_Vault *VaultSession) IsWithdrawed(hash [32]byte) (bool, error) {
	return _Vault.Contract.IsWithdrawed(&_Vault.CallOpts, hash)
}

// IsWithdrawed is a free data retrieval call binding the contract method 0x749c5f86.
//
// Solidity: function isWithdrawed(bytes32 hash) view returns(bool)
func (_Vault *VaultCallerSession) IsWithdrawed(hash [32]byte) (bool, error) {
	return _Vault.Contract.IsWithdrawed(&_Vault.CallOpts, hash)
}

// Locker is a free data retrieval call binding the contract method 0xd7b96d4e.
//
// Solidity: function locker() view returns(address)
func (_Vault *VaultCaller) Locker(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "locker")
	return *ret0, err
}

// Locker is a free data retrieval call binding the contract method 0xd7b96d4e.
//
// Solidity: function locker() view returns(address)
func (_Vault *VaultSession) Locker() (common.Address, error) {
	return _Vault.Contract.Locker(&_Vault.CallOpts)
}

// Locker is a free data retrieval call binding the contract method 0xd7b96d4e.
//
// Solidity: function locker() view returns(address)
func (_Vault *VaultCallerSession) Locker() (common.Address, error) {
	return _Vault.Contract.Locker(&_Vault.CallOpts)
}

// Migration is a free data retrieval call binding the contract method 0x995fac11.
//
// Solidity: function migration(address , address ) view returns(bool)
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
// Solidity: function migration(address , address ) view returns(bool)
func (_Vault *VaultSession) Migration(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Vault.Contract.Migration(&_Vault.CallOpts, arg0, arg1)
}

// Migration is a free data retrieval call binding the contract method 0x995fac11.
//
// Solidity: function migration(address , address ) view returns(bool)
func (_Vault *VaultCallerSession) Migration(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Vault.Contract.Migration(&_Vault.CallOpts, arg0, arg1)
}

// NewVault is a free data retrieval call binding the contract method 0x88aaf0c8.
//
// Solidity: function newVault() view returns(address)
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
// Solidity: function newVault() view returns(address)
func (_Vault *VaultSession) NewVault() (common.Address, error) {
	return _Vault.Contract.NewVault(&_Vault.CallOpts)
}

// NewVault is a free data retrieval call binding the contract method 0x88aaf0c8.
//
// Solidity: function newVault() view returns(address)
func (_Vault *VaultCallerSession) NewVault() (common.Address, error) {
	return _Vault.Contract.NewVault(&_Vault.CallOpts)
}

// NotEntered is a free data retrieval call binding the contract method 0xa3f5d8cc.
//
// Solidity: function notEntered() view returns(bool)
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
// Solidity: function notEntered() view returns(bool)
func (_Vault *VaultSession) NotEntered() (bool, error) {
	return _Vault.Contract.NotEntered(&_Vault.CallOpts)
}

// NotEntered is a free data retrieval call binding the contract method 0xa3f5d8cc.
//
// Solidity: function notEntered() view returns(bool)
func (_Vault *VaultCallerSession) NotEntered() (bool, error) {
	return _Vault.Contract.NotEntered(&_Vault.CallOpts)
}

// ParseBurnInst is a free data retrieval call binding the contract method 0x7e16e6e1.
//
// Solidity: function parseBurnInst(bytes inst) pure returns((uint8,uint8,address,address,uint256,bytes32))
func (_Vault *VaultCaller) ParseBurnInst(opts *bind.CallOpts, inst []byte) (VaultBurnInstData, error) {
	var (
		ret0 = new(VaultBurnInstData)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "parseBurnInst", inst)
	return *ret0, err
}

// ParseBurnInst is a free data retrieval call binding the contract method 0x7e16e6e1.
//
// Solidity: function parseBurnInst(bytes inst) pure returns((uint8,uint8,address,address,uint256,bytes32))
func (_Vault *VaultSession) ParseBurnInst(inst []byte) (VaultBurnInstData, error) {
	return _Vault.Contract.ParseBurnInst(&_Vault.CallOpts, inst)
}

// ParseBurnInst is a free data retrieval call binding the contract method 0x7e16e6e1.
//
// Solidity: function parseBurnInst(bytes inst) pure returns((uint8,uint8,address,address,uint256,bytes32))
func (_Vault *VaultCallerSession) ParseBurnInst(inst []byte) (VaultBurnInstData, error) {
	return _Vault.Contract.ParseBurnInst(&_Vault.CallOpts, inst)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
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
// Solidity: function paused() view returns(bool)
func (_Vault *VaultSession) Paused() (bool, error) {
	return _Vault.Contract.Paused(&_Vault.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Vault *VaultCallerSession) Paused() (bool, error) {
	return _Vault.Contract.Paused(&_Vault.CallOpts)
}

// PrevVault is a free data retrieval call binding the contract method 0xfa84702e.
//
// Solidity: function prevVault() view returns(address)
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
// Solidity: function prevVault() view returns(address)
func (_Vault *VaultSession) PrevVault() (common.Address, error) {
	return _Vault.Contract.PrevVault(&_Vault.CallOpts)
}

// PrevVault is a free data retrieval call binding the contract method 0xfa84702e.
//
// Solidity: function prevVault() view returns(address)
func (_Vault *VaultCallerSession) PrevVault() (common.Address, error) {
	return _Vault.Contract.PrevVault(&_Vault.CallOpts)
}

// SigDataUsed is a free data retrieval call binding the contract method 0x1ea1940e.
//
// Solidity: function sigDataUsed(bytes32 ) view returns(bool)
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
// Solidity: function sigDataUsed(bytes32 ) view returns(bool)
func (_Vault *VaultSession) SigDataUsed(arg0 [32]byte) (bool, error) {
	return _Vault.Contract.SigDataUsed(&_Vault.CallOpts, arg0)
}

// SigDataUsed is a free data retrieval call binding the contract method 0x1ea1940e.
//
// Solidity: function sigDataUsed(bytes32 ) view returns(bool)
func (_Vault *VaultCallerSession) SigDataUsed(arg0 [32]byte) (bool, error) {
	return _Vault.Contract.SigDataUsed(&_Vault.CallOpts, arg0)
}

// SigToAddress is a free data retrieval call binding the contract method 0x3fec6b40.
//
// Solidity: function sigToAddress(bytes signData, bytes32 hash) pure returns(address)
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
// Solidity: function sigToAddress(bytes signData, bytes32 hash) pure returns(address)
func (_Vault *VaultSession) SigToAddress(signData []byte, hash [32]byte) (common.Address, error) {
	return _Vault.Contract.SigToAddress(&_Vault.CallOpts, signData, hash)
}

// SigToAddress is a free data retrieval call binding the contract method 0x3fec6b40.
//
// Solidity: function sigToAddress(bytes signData, bytes32 hash) pure returns(address)
func (_Vault *VaultCallerSession) SigToAddress(signData []byte, hash [32]byte) (common.Address, error) {
	return _Vault.Contract.SigToAddress(&_Vault.CallOpts, signData, hash)
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() view returns(address)
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
// Solidity: function successor() view returns(address)
func (_Vault *VaultSession) Successor() (common.Address, error) {
	return _Vault.Contract.Successor(&_Vault.CallOpts)
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() view returns(address)
func (_Vault *VaultCallerSession) Successor() (common.Address, error) {
	return _Vault.Contract.Successor(&_Vault.CallOpts)
}

// TotalDepositedToSCAmount is a free data retrieval call binding the contract method 0x6304541c.
//
// Solidity: function totalDepositedToSCAmount(address ) view returns(uint256)
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
// Solidity: function totalDepositedToSCAmount(address ) view returns(uint256)
func (_Vault *VaultSession) TotalDepositedToSCAmount(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.TotalDepositedToSCAmount(&_Vault.CallOpts, arg0)
}

// TotalDepositedToSCAmount is a free data retrieval call binding the contract method 0x6304541c.
//
// Solidity: function totalDepositedToSCAmount(address ) view returns(uint256)
func (_Vault *VaultCallerSession) TotalDepositedToSCAmount(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.TotalDepositedToSCAmount(&_Vault.CallOpts, arg0)
}

// WithdrawRequests is a free data retrieval call binding the contract method 0x65b5a00f.
//
// Solidity: function withdrawRequests(address , address ) view returns(uint256)
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
// Solidity: function withdrawRequests(address , address ) view returns(uint256)
func (_Vault *VaultSession) WithdrawRequests(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Vault.Contract.WithdrawRequests(&_Vault.CallOpts, arg0, arg1)
}

// WithdrawRequests is a free data retrieval call binding the contract method 0x65b5a00f.
//
// Solidity: function withdrawRequests(address , address ) view returns(uint256)
func (_Vault *VaultCallerSession) WithdrawRequests(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Vault.Contract.WithdrawRequests(&_Vault.CallOpts, arg0, arg1)
}

// Withdrawed is a free data retrieval call binding the contract method 0xdca40d9e.
//
// Solidity: function withdrawed(bytes32 ) view returns(bool)
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
// Solidity: function withdrawed(bytes32 ) view returns(bool)
func (_Vault *VaultSession) Withdrawed(arg0 [32]byte) (bool, error) {
	return _Vault.Contract.Withdrawed(&_Vault.CallOpts, arg0)
}

// Withdrawed is a free data retrieval call binding the contract method 0xdca40d9e.
//
// Solidity: function withdrawed(bytes32 ) view returns(bool)
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
// Solidity: function deposit(string incognitoAddress) payable returns()
func (_Vault *VaultTransactor) Deposit(opts *bind.TransactOpts, incognitoAddress string) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "deposit", incognitoAddress)
}

// Deposit is a paid mutator transaction binding the contract method 0xa26e1186.
//
// Solidity: function deposit(string incognitoAddress) payable returns()
func (_Vault *VaultSession) Deposit(incognitoAddress string) (*types.Transaction, error) {
	return _Vault.Contract.Deposit(&_Vault.TransactOpts, incognitoAddress)
}

// Deposit is a paid mutator transaction binding the contract method 0xa26e1186.
//
// Solidity: function deposit(string incognitoAddress) payable returns()
func (_Vault *VaultTransactorSession) Deposit(incognitoAddress string) (*types.Transaction, error) {
	return _Vault.Contract.Deposit(&_Vault.TransactOpts, incognitoAddress)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x5a67cb87.
//
// Solidity: function depositERC20(address token, uint256 amount, string incognitoAddress) payable returns()
func (_Vault *VaultTransactor) DepositERC20(opts *bind.TransactOpts, token common.Address, amount *big.Int, incognitoAddress string) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "depositERC20", token, amount, incognitoAddress)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x5a67cb87.
//
// Solidity: function depositERC20(address token, uint256 amount, string incognitoAddress) payable returns()
func (_Vault *VaultSession) DepositERC20(token common.Address, amount *big.Int, incognitoAddress string) (*types.Transaction, error) {
	return _Vault.Contract.DepositERC20(&_Vault.TransactOpts, token, amount, incognitoAddress)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x5a67cb87.
//
// Solidity: function depositERC20(address token, uint256 amount, string incognitoAddress) payable returns()
func (_Vault *VaultTransactorSession) DepositERC20(token common.Address, amount *big.Int, incognitoAddress string) (*types.Transaction, error) {
	return _Vault.Contract.DepositERC20(&_Vault.TransactOpts, token, amount, incognitoAddress)
}

// Execute is a paid mutator transaction binding the contract method 0x8588ccd6.
//
// Solidity: function execute(address token, uint256 amount, address recipientToken, address exchangeAddress, bytes callData, bytes timestamp, bytes signData) payable returns()
func (_Vault *VaultTransactor) Execute(opts *bind.TransactOpts, token common.Address, amount *big.Int, recipientToken common.Address, exchangeAddress common.Address, callData []byte, timestamp []byte, signData []byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "execute", token, amount, recipientToken, exchangeAddress, callData, timestamp, signData)
}

// Execute is a paid mutator transaction binding the contract method 0x8588ccd6.
//
// Solidity: function execute(address token, uint256 amount, address recipientToken, address exchangeAddress, bytes callData, bytes timestamp, bytes signData) payable returns()
func (_Vault *VaultSession) Execute(token common.Address, amount *big.Int, recipientToken common.Address, exchangeAddress common.Address, callData []byte, timestamp []byte, signData []byte) (*types.Transaction, error) {
	return _Vault.Contract.Execute(&_Vault.TransactOpts, token, amount, recipientToken, exchangeAddress, callData, timestamp, signData)
}

// Execute is a paid mutator transaction binding the contract method 0x8588ccd6.
//
// Solidity: function execute(address token, uint256 amount, address recipientToken, address exchangeAddress, bytes callData, bytes timestamp, bytes signData) payable returns()
func (_Vault *VaultTransactorSession) Execute(token common.Address, amount *big.Int, recipientToken common.Address, exchangeAddress common.Address, callData []byte, timestamp []byte, signData []byte) (*types.Transaction, error) {
	return _Vault.Contract.Execute(&_Vault.TransactOpts, token, amount, recipientToken, exchangeAddress, callData, timestamp, signData)
}

// ExecuteMulti is a paid mutator transaction binding the contract method 0xb69f6511.
//
// Solidity: function executeMulti(address[] tokens, uint256[] amounts, address[] recipientTokens, address exchangeAddress, bytes callData, bytes timestamp, bytes signData) payable returns()
func (_Vault *VaultTransactor) ExecuteMulti(opts *bind.TransactOpts, tokens []common.Address, amounts []*big.Int, recipientTokens []common.Address, exchangeAddress common.Address, callData []byte, timestamp []byte, signData []byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "executeMulti", tokens, amounts, recipientTokens, exchangeAddress, callData, timestamp, signData)
}

// ExecuteMulti is a paid mutator transaction binding the contract method 0xb69f6511.
//
// Solidity: function executeMulti(address[] tokens, uint256[] amounts, address[] recipientTokens, address exchangeAddress, bytes callData, bytes timestamp, bytes signData) payable returns()
func (_Vault *VaultSession) ExecuteMulti(tokens []common.Address, amounts []*big.Int, recipientTokens []common.Address, exchangeAddress common.Address, callData []byte, timestamp []byte, signData []byte) (*types.Transaction, error) {
	return _Vault.Contract.ExecuteMulti(&_Vault.TransactOpts, tokens, amounts, recipientTokens, exchangeAddress, callData, timestamp, signData)
}

// ExecuteMulti is a paid mutator transaction binding the contract method 0xb69f6511.
//
// Solidity: function executeMulti(address[] tokens, uint256[] amounts, address[] recipientTokens, address exchangeAddress, bytes callData, bytes timestamp, bytes signData) payable returns()
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

// SubmitBurnProof is a paid mutator transaction binding the contract method 0x73bf9651.
//
// Solidity: function submitBurnProof(bytes inst, uint256 heights, bytes32[] instPaths, bool[] instPathIsLefts, bytes32 instRoots, bytes32 blkData, uint256[] sigIdxs, uint8[] sigVs, bytes32[] sigRs, bytes32[] sigSs) returns()
func (_Vault *VaultTransactor) SubmitBurnProof(opts *bind.TransactOpts, inst []byte, heights *big.Int, instPaths [][32]byte, instPathIsLefts []bool, instRoots [32]byte, blkData [32]byte, sigIdxs []*big.Int, sigVs []uint8, sigRs [][32]byte, sigSs [][32]byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "submitBurnProof", inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// SubmitBurnProof is a paid mutator transaction binding the contract method 0x73bf9651.
//
// Solidity: function submitBurnProof(bytes inst, uint256 heights, bytes32[] instPaths, bool[] instPathIsLefts, bytes32 instRoots, bytes32 blkData, uint256[] sigIdxs, uint8[] sigVs, bytes32[] sigRs, bytes32[] sigSs) returns()
func (_Vault *VaultSession) SubmitBurnProof(inst []byte, heights *big.Int, instPaths [][32]byte, instPathIsLefts []bool, instRoots [32]byte, blkData [32]byte, sigIdxs []*big.Int, sigVs []uint8, sigRs [][32]byte, sigSs [][32]byte) (*types.Transaction, error) {
	return _Vault.Contract.SubmitBurnProof(&_Vault.TransactOpts, inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// SubmitBurnProof is a paid mutator transaction binding the contract method 0x73bf9651.
//
// Solidity: function submitBurnProof(bytes inst, uint256 heights, bytes32[] instPaths, bool[] instPathIsLefts, bytes32 instRoots, bytes32 blkData, uint256[] sigIdxs, uint8[] sigVs, bytes32[] sigRs, bytes32[] sigSs) returns()
func (_Vault *VaultTransactorSession) SubmitBurnProof(inst []byte, heights *big.Int, instPaths [][32]byte, instPathIsLefts []bool, instRoots [32]byte, blkData [32]byte, sigIdxs []*big.Int, sigVs []uint8, sigRs [][32]byte, sigSs [][32]byte) (*types.Transaction, error) {
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

// UpdateLocker is a paid mutator transaction binding the contract method 0xe1d5ea8a.
//
// Solidity: function updateLocker(address newLocker) returns()
func (_Vault *VaultTransactor) UpdateLocker(opts *bind.TransactOpts, newLocker common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "updateLocker", newLocker)
}

// UpdateLocker is a paid mutator transaction binding the contract method 0xe1d5ea8a.
//
// Solidity: function updateLocker(address newLocker) returns()
func (_Vault *VaultSession) UpdateLocker(newLocker common.Address) (*types.Transaction, error) {
	return _Vault.Contract.UpdateLocker(&_Vault.TransactOpts, newLocker)
}

// UpdateLocker is a paid mutator transaction binding the contract method 0xe1d5ea8a.
//
// Solidity: function updateLocker(address newLocker) returns()
func (_Vault *VaultTransactorSession) UpdateLocker(newLocker common.Address) (*types.Transaction, error) {
	return _Vault.Contract.UpdateLocker(&_Vault.TransactOpts, newLocker)
}

// Withdraw is a paid mutator transaction binding the contract method 0x1beb7de2.
//
// Solidity: function withdraw(bytes inst, uint256 heights, bytes32[] instPaths, bool[] instPathIsLefts, bytes32 instRoots, bytes32 blkData, uint256[] sigIdxs, uint8[] sigVs, bytes32[] sigRs, bytes32[] sigSs) returns()
func (_Vault *VaultTransactor) Withdraw(opts *bind.TransactOpts, inst []byte, heights *big.Int, instPaths [][32]byte, instPathIsLefts []bool, instRoots [32]byte, blkData [32]byte, sigIdxs []*big.Int, sigVs []uint8, sigRs [][32]byte, sigSs [][32]byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "withdraw", inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// Withdraw is a paid mutator transaction binding the contract method 0x1beb7de2.
//
// Solidity: function withdraw(bytes inst, uint256 heights, bytes32[] instPaths, bool[] instPathIsLefts, bytes32 instRoots, bytes32 blkData, uint256[] sigIdxs, uint8[] sigVs, bytes32[] sigRs, bytes32[] sigSs) returns()
func (_Vault *VaultSession) Withdraw(inst []byte, heights *big.Int, instPaths [][32]byte, instPathIsLefts []bool, instRoots [32]byte, blkData [32]byte, sigIdxs []*big.Int, sigVs []uint8, sigRs [][32]byte, sigSs [][32]byte) (*types.Transaction, error) {
	return _Vault.Contract.Withdraw(&_Vault.TransactOpts, inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// Withdraw is a paid mutator transaction binding the contract method 0x1beb7de2.
//
// Solidity: function withdraw(bytes inst, uint256 heights, bytes32[] instPaths, bool[] instPathIsLefts, bytes32 instRoots, bytes32 blkData, uint256[] sigIdxs, uint8[] sigVs, bytes32[] sigRs, bytes32[] sigSs) returns()
func (_Vault *VaultTransactorSession) Withdraw(inst []byte, heights *big.Int, instPaths [][32]byte, instPathIsLefts []bool, instRoots [32]byte, blkData [32]byte, sigIdxs []*big.Int, sigVs []uint8, sigRs [][32]byte, sigSs [][32]byte) (*types.Transaction, error) {
	return _Vault.Contract.Withdraw(&_Vault.TransactOpts, inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

//// Receive is a paid mutator transaction binding the contract receive function.
////
//// Solidity: receive() payable returns()
//func (_Vault *VaultTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
//	return _Vault.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
//}

//// Receive is a paid mutator transaction binding the contract receive function.
////
//// Solidity: receive() payable returns()
//func (_Vault *VaultSession) Receive() (*types.Transaction, error) {
//	return _Vault.Contract.Receive(&_Vault.TransactOpts)
//}

//// Receive is a paid mutator transaction binding the contract receive function.
////
//// Solidity: receive() payable returns()
//func (_Vault *VaultTransactorSession) Receive() (*types.Transaction, error) {
//	return _Vault.Contract.Receive(&_Vault.TransactOpts)
//}

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
// Solidity: function getDepositedBalance(address , address ) view returns(uint256)
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
// Solidity: function getDepositedBalance(address , address ) view returns(uint256)
func (_Withdrawable *WithdrawableSession) GetDepositedBalance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Withdrawable.Contract.GetDepositedBalance(&_Withdrawable.CallOpts, arg0, arg1)
}

// GetDepositedBalance is a free data retrieval call binding the contract method 0xf75b98ce.
//
// Solidity: function getDepositedBalance(address , address ) view returns(uint256)
func (_Withdrawable *WithdrawableCallerSession) GetDepositedBalance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Withdrawable.Contract.GetDepositedBalance(&_Withdrawable.CallOpts, arg0, arg1)
}

// IsSigDataUsed is a free data retrieval call binding the contract method 0xe4bd7074.
//
// Solidity: function isSigDataUsed(bytes32 ) view returns(bool)
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
// Solidity: function isSigDataUsed(bytes32 ) view returns(bool)
func (_Withdrawable *WithdrawableSession) IsSigDataUsed(arg0 [32]byte) (bool, error) {
	return _Withdrawable.Contract.IsSigDataUsed(&_Withdrawable.CallOpts, arg0)
}

// IsSigDataUsed is a free data retrieval call binding the contract method 0xe4bd7074.
//
// Solidity: function isSigDataUsed(bytes32 ) view returns(bool)
func (_Withdrawable *WithdrawableCallerSession) IsSigDataUsed(arg0 [32]byte) (bool, error) {
	return _Withdrawable.Contract.IsSigDataUsed(&_Withdrawable.CallOpts, arg0)
}

// IsWithdrawed is a free data retrieval call binding the contract method 0x749c5f86.
//
// Solidity: function isWithdrawed(bytes32 ) view returns(bool)
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
// Solidity: function isWithdrawed(bytes32 ) view returns(bool)
func (_Withdrawable *WithdrawableSession) IsWithdrawed(arg0 [32]byte) (bool, error) {
	return _Withdrawable.Contract.IsWithdrawed(&_Withdrawable.CallOpts, arg0)
}

// IsWithdrawed is a free data retrieval call binding the contract method 0x749c5f86.
//
// Solidity: function isWithdrawed(bytes32 ) view returns(bool)
func (_Withdrawable *WithdrawableCallerSession) IsWithdrawed(arg0 [32]byte) (bool, error) {
	return _Withdrawable.Contract.IsWithdrawed(&_Withdrawable.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
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
// Solidity: function paused() view returns(bool)
func (_Withdrawable *WithdrawableSession) Paused() (bool, error) {
	return _Withdrawable.Contract.Paused(&_Withdrawable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
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
