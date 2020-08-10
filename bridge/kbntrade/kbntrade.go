// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kbntrade

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

// GetLockerABI is the input ABI used to generate the binding from.
const GetLockerABI = "[{\"inputs\":[],\"name\":\"locker\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GetLockerFuncSigs maps the 4-byte function signature to its string representation.
var GetLockerFuncSigs = map[string]string{
	"d7b96d4e": "locker()",
}

// GetLocker is an auto generated Go binding around an Ethereum contract.
type GetLocker struct {
	GetLockerCaller     // Read-only binding to the contract
	GetLockerTransactor // Write-only binding to the contract
	GetLockerFilterer   // Log filterer for contract events
}

// GetLockerCaller is an auto generated read-only Go binding around an Ethereum contract.
type GetLockerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GetLockerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GetLockerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GetLockerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GetLockerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GetLockerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GetLockerSession struct {
	Contract     *GetLocker        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GetLockerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GetLockerCallerSession struct {
	Contract *GetLockerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// GetLockerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GetLockerTransactorSession struct {
	Contract     *GetLockerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// GetLockerRaw is an auto generated low-level Go binding around an Ethereum contract.
type GetLockerRaw struct {
	Contract *GetLocker // Generic contract binding to access the raw methods on
}

// GetLockerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GetLockerCallerRaw struct {
	Contract *GetLockerCaller // Generic read-only contract binding to access the raw methods on
}

// GetLockerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GetLockerTransactorRaw struct {
	Contract *GetLockerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGetLocker creates a new instance of GetLocker, bound to a specific deployed contract.
func NewGetLocker(address common.Address, backend bind.ContractBackend) (*GetLocker, error) {
	contract, err := bindGetLocker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GetLocker{GetLockerCaller: GetLockerCaller{contract: contract}, GetLockerTransactor: GetLockerTransactor{contract: contract}, GetLockerFilterer: GetLockerFilterer{contract: contract}}, nil
}

// NewGetLockerCaller creates a new read-only instance of GetLocker, bound to a specific deployed contract.
func NewGetLockerCaller(address common.Address, caller bind.ContractCaller) (*GetLockerCaller, error) {
	contract, err := bindGetLocker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GetLockerCaller{contract: contract}, nil
}

// NewGetLockerTransactor creates a new write-only instance of GetLocker, bound to a specific deployed contract.
func NewGetLockerTransactor(address common.Address, transactor bind.ContractTransactor) (*GetLockerTransactor, error) {
	contract, err := bindGetLocker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GetLockerTransactor{contract: contract}, nil
}

// NewGetLockerFilterer creates a new log filterer instance of GetLocker, bound to a specific deployed contract.
func NewGetLockerFilterer(address common.Address, filterer bind.ContractFilterer) (*GetLockerFilterer, error) {
	contract, err := bindGetLocker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GetLockerFilterer{contract: contract}, nil
}

// bindGetLocker binds a generic wrapper to an already deployed contract.
func bindGetLocker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GetLockerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GetLocker *GetLockerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GetLocker.Contract.GetLockerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GetLocker *GetLockerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GetLocker.Contract.GetLockerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GetLocker *GetLockerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GetLocker.Contract.GetLockerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GetLocker *GetLockerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GetLocker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GetLocker *GetLockerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GetLocker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GetLocker *GetLockerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GetLocker.Contract.contract.Transact(opts, method, params...)
}

// Locker is a paid mutator transaction binding the contract method 0xd7b96d4e.
//
// Solidity: function locker() returns(address)
func (_GetLocker *GetLockerTransactor) Locker(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GetLocker.contract.Transact(opts, "locker")
}

// Locker is a paid mutator transaction binding the contract method 0xd7b96d4e.
//
// Solidity: function locker() returns(address)
func (_GetLocker *GetLockerSession) Locker() (*types.Transaction, error) {
	return _GetLocker.Contract.Locker(&_GetLocker.TransactOpts)
}

// Locker is a paid mutator transaction binding the contract method 0xd7b96d4e.
//
// Solidity: function locker() returns(address)
func (_GetLocker *GetLockerTransactorSession) Locker() (*types.Transaction, error) {
	return _GetLocker.Contract.Locker(&_GetLocker.TransactOpts)
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

// KBNTradeABI is the input ABI used to generate the binding from.
const KBNTradeABI = "[{\"inputs\":[{\"internalType\":\"contractKyberNetwork\",\"name\":\"_kyberNetworkProxyContract\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_incognitoSmartContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ETH_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcQty\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"destToken\",\"type\":\"address\"}],\"name\":\"getConversionRates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incognitoSmartContract\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kyberNetworkProxyContract\",\"outputs\":[{\"internalType\":\"contractKyberNetwork\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcQty\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minConversionRate\",\"type\":\"uint256\"}],\"name\":\"trade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// KBNTradeFuncSigs maps the 4-byte function signature to its string representation.
var KBNTradeFuncSigs = map[string]string{
	"72e94bf6": "ETH_CONTRACT_ADDRESS()",
	"0aea8188": "getConversionRates(address,uint256,address)",
	"b42a644b": "incognitoSmartContract()",
	"785250da": "kyberNetworkProxyContract()",
	"bb39a960": "trade(address,uint256,address,uint256)",
}

// KBNTradeBin is the compiled bytecode used for deploying new contracts.
var KBNTradeBin = "0x608060405234801561001057600080fd5b506040516108953803806108958339818101604052604081101561003357600080fd5b508051602090910151600180546001600160a01b039384166001600160a01b0319918216179091556000805493909216921691909117905561081b8061007a6000396000f3fe60806040526004361061004e5760003560e01c80630aea81881461005a57806372e94bf6146100b6578063785250da146100e7578063b42a644b146100fc578063bb39a9601461011157610055565b3661005557005b600080fd5b34801561006657600080fd5b5061009d6004803603606081101561007d57600080fd5b506001600160a01b0381358116916020810135916040909101351661016e565b6040805192835260208301919091528051918290030190f35b3480156100c257600080fd5b506100cb610209565b604080516001600160a01b039092168252519081900360200190f35b3480156100f357600080fd5b506100cb61020e565b34801561010857600080fd5b506100cb61021d565b61014b6004803603608081101561012757600080fd5b506001600160a01b038135811691602081013591604082013516906060013561022c565b604080516001600160a01b03909316835260208301919091528051918290030190f35b6001546040805163809a9e5560e01b81526001600160a01b0386811660048301528481166024830152604482018690528251600094859492169263809a9e55926064808301939192829003018186803b1580156101ca57600080fd5b505afa1580156101de573d6000803e3d6000fd5b505050506040513d60408110156101f457600080fd5b50805160209091015190969095509350505050565b600081565b6001546001600160a01b031681565b6000546001600160a01b031681565b6000805481906001600160a01b0316331461024657600080fd5b8461025087610324565b101561025b57600080fd5b836001600160a01b0316866001600160a01b0316141561027a57600080fd5b60006001600160a01b038716156102eb576001546102a39088906001600160a01b0316886103b5565b6001600160a01b038516156102cf5760006102c0888888886104c1565b116102ca57600080fd5b6102e6565b60006102dc88888761055e565b116102e657600080fd5b610302565b60006102f88688876105f2565b1161030257600080fd5b61030b85610324565b90506103178582610684565b9396939550929350505050565b60006001600160a01b03821661033b5750476103b0565b604080516370a0823160e01b815230600482015290516001600160a01b038416916370a08231916024808301926020929190829003018186803b15801561038157600080fd5b505afa158015610395573d6000803e3d6000fd5b505050506040513d60208110156103ab57600080fd5b505190505b919050565b6001600160a01b038316156104bc576040805163095ea7b360e01b81526001600160a01b03848116600483015260006024830181905292519086169263095ea7b3926044808201939182900301818387803b15801561041357600080fd5b505af1158015610427573d6000803e3d6000fd5b505050506104336107b1565b61043c57600080fd5b826001600160a01b031663095ea7b383836040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050600060405180830381600087803b15801561049357600080fd5b505af11580156104a7573d6000803e3d6000fd5b505050506104b36107b1565b6104bc57600080fd5b505050565b60015460408051637409e2eb60e01b81526001600160a01b0387811660048301526024820187905285811660448301526064820185905291516000939290921691637409e2eb9160848082019260209290919082900301818787803b15801561052957600080fd5b505af115801561053d573d6000803e3d6000fd5b505050506040513d602081101561055357600080fd5b505195945050505050565b60015460408051630eee887760e21b81526001600160a01b038681166004830152602482018690526044820185905291516000939290921691633bba21dc9160648082019260209290919082900301818787803b1580156105be57600080fd5b505af11580156105d2573d6000803e3d6000fd5b505050506040513d60208110156105e857600080fd5b5051949350505050565b60008247101561060157600080fd5b60015460408051633d15022b60e11b81526001600160a01b0387811660048301526024820186905291519190921691637a2a045691869160448082019260209290919082900301818588803b15801561065957600080fd5b505af115801561066d573d6000803e3d6000fd5b50505050506040513d60208110156105e857600080fd5b60008060009054906101000a90046001600160a01b03166001600160a01b031663d7b96d4e6040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156106d557600080fd5b505af11580156106e9573d6000803e3d6000fd5b505050506040513d60208110156106ff57600080fd5b505190506001600160a01b03831661075a578147101561071e57600080fd5b6040516001600160a01b0382169083156108fc029084906000818181858888f19350505050158015610754573d6000803e3d6000fd5b506104bc565b826001600160a01b031663a9059cbb82846040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050600060405180830381600087803b15801561049357600080fd5b6000803d80156107c857602081146107d1576107dd565b600191506107dd565b60206000803e60005191505b50151590509056fea2646970667358221220f63f5821597d27a2b0ffda690045ee9ce1b770dca8440dbcc49ceee5e7d86d4b64736f6c634300060c0033"

// DeployKBNTrade deploys a new Ethereum contract, binding an instance of KBNTrade to it.
func DeployKBNTrade(auth *bind.TransactOpts, backend bind.ContractBackend, _kyberNetworkProxyContract common.Address, _incognitoSmartContract common.Address) (common.Address, *types.Transaction, *KBNTrade, error) {
	parsed, err := abi.JSON(strings.NewReader(KBNTradeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KBNTradeBin), backend, _kyberNetworkProxyContract, _incognitoSmartContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KBNTrade{KBNTradeCaller: KBNTradeCaller{contract: contract}, KBNTradeTransactor: KBNTradeTransactor{contract: contract}, KBNTradeFilterer: KBNTradeFilterer{contract: contract}}, nil
}

// KBNTrade is an auto generated Go binding around an Ethereum contract.
type KBNTrade struct {
	KBNTradeCaller     // Read-only binding to the contract
	KBNTradeTransactor // Write-only binding to the contract
	KBNTradeFilterer   // Log filterer for contract events
}

// KBNTradeCaller is an auto generated read-only Go binding around an Ethereum contract.
type KBNTradeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KBNTradeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KBNTradeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KBNTradeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KBNTradeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KBNTradeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KBNTradeSession struct {
	Contract     *KBNTrade         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KBNTradeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KBNTradeCallerSession struct {
	Contract *KBNTradeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// KBNTradeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KBNTradeTransactorSession struct {
	Contract     *KBNTradeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// KBNTradeRaw is an auto generated low-level Go binding around an Ethereum contract.
type KBNTradeRaw struct {
	Contract *KBNTrade // Generic contract binding to access the raw methods on
}

// KBNTradeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KBNTradeCallerRaw struct {
	Contract *KBNTradeCaller // Generic read-only contract binding to access the raw methods on
}

// KBNTradeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KBNTradeTransactorRaw struct {
	Contract *KBNTradeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKBNTrade creates a new instance of KBNTrade, bound to a specific deployed contract.
func NewKBNTrade(address common.Address, backend bind.ContractBackend) (*KBNTrade, error) {
	contract, err := bindKBNTrade(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KBNTrade{KBNTradeCaller: KBNTradeCaller{contract: contract}, KBNTradeTransactor: KBNTradeTransactor{contract: contract}, KBNTradeFilterer: KBNTradeFilterer{contract: contract}}, nil
}

// NewKBNTradeCaller creates a new read-only instance of KBNTrade, bound to a specific deployed contract.
func NewKBNTradeCaller(address common.Address, caller bind.ContractCaller) (*KBNTradeCaller, error) {
	contract, err := bindKBNTrade(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KBNTradeCaller{contract: contract}, nil
}

// NewKBNTradeTransactor creates a new write-only instance of KBNTrade, bound to a specific deployed contract.
func NewKBNTradeTransactor(address common.Address, transactor bind.ContractTransactor) (*KBNTradeTransactor, error) {
	contract, err := bindKBNTrade(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KBNTradeTransactor{contract: contract}, nil
}

// NewKBNTradeFilterer creates a new log filterer instance of KBNTrade, bound to a specific deployed contract.
func NewKBNTradeFilterer(address common.Address, filterer bind.ContractFilterer) (*KBNTradeFilterer, error) {
	contract, err := bindKBNTrade(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KBNTradeFilterer{contract: contract}, nil
}

// bindKBNTrade binds a generic wrapper to an already deployed contract.
func bindKBNTrade(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KBNTradeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KBNTrade *KBNTradeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KBNTrade.Contract.KBNTradeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KBNTrade *KBNTradeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KBNTrade.Contract.KBNTradeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KBNTrade *KBNTradeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KBNTrade.Contract.KBNTradeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KBNTrade *KBNTradeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KBNTrade.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KBNTrade *KBNTradeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KBNTrade.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KBNTrade *KBNTradeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KBNTrade.Contract.contract.Transact(opts, method, params...)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_KBNTrade *KBNTradeCaller) ETHCONTRACTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KBNTrade.contract.Call(opts, out, "ETH_CONTRACT_ADDRESS")
	return *ret0, err
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_KBNTrade *KBNTradeSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _KBNTrade.Contract.ETHCONTRACTADDRESS(&_KBNTrade.CallOpts)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_KBNTrade *KBNTradeCallerSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _KBNTrade.Contract.ETHCONTRACTADDRESS(&_KBNTrade.CallOpts)
}

// GetConversionRates is a free data retrieval call binding the contract method 0x0aea8188.
//
// Solidity: function getConversionRates(address srcToken, uint256 srcQty, address destToken) constant returns(uint256, uint256)
func (_KBNTrade *KBNTradeCaller) GetConversionRates(opts *bind.CallOpts, srcToken common.Address, srcQty *big.Int, destToken common.Address) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _KBNTrade.contract.Call(opts, out, "getConversionRates", srcToken, srcQty, destToken)
	return *ret0, *ret1, err
}

// GetConversionRates is a free data retrieval call binding the contract method 0x0aea8188.
//
// Solidity: function getConversionRates(address srcToken, uint256 srcQty, address destToken) constant returns(uint256, uint256)
func (_KBNTrade *KBNTradeSession) GetConversionRates(srcToken common.Address, srcQty *big.Int, destToken common.Address) (*big.Int, *big.Int, error) {
	return _KBNTrade.Contract.GetConversionRates(&_KBNTrade.CallOpts, srcToken, srcQty, destToken)
}

// GetConversionRates is a free data retrieval call binding the contract method 0x0aea8188.
//
// Solidity: function getConversionRates(address srcToken, uint256 srcQty, address destToken) constant returns(uint256, uint256)
func (_KBNTrade *KBNTradeCallerSession) GetConversionRates(srcToken common.Address, srcQty *big.Int, destToken common.Address) (*big.Int, *big.Int, error) {
	return _KBNTrade.Contract.GetConversionRates(&_KBNTrade.CallOpts, srcToken, srcQty, destToken)
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_KBNTrade *KBNTradeCaller) IncognitoSmartContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KBNTrade.contract.Call(opts, out, "incognitoSmartContract")
	return *ret0, err
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_KBNTrade *KBNTradeSession) IncognitoSmartContract() (common.Address, error) {
	return _KBNTrade.Contract.IncognitoSmartContract(&_KBNTrade.CallOpts)
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_KBNTrade *KBNTradeCallerSession) IncognitoSmartContract() (common.Address, error) {
	return _KBNTrade.Contract.IncognitoSmartContract(&_KBNTrade.CallOpts)
}

// KyberNetworkProxyContract is a free data retrieval call binding the contract method 0x785250da.
//
// Solidity: function kyberNetworkProxyContract() constant returns(address)
func (_KBNTrade *KBNTradeCaller) KyberNetworkProxyContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KBNTrade.contract.Call(opts, out, "kyberNetworkProxyContract")
	return *ret0, err
}

// KyberNetworkProxyContract is a free data retrieval call binding the contract method 0x785250da.
//
// Solidity: function kyberNetworkProxyContract() constant returns(address)
func (_KBNTrade *KBNTradeSession) KyberNetworkProxyContract() (common.Address, error) {
	return _KBNTrade.Contract.KyberNetworkProxyContract(&_KBNTrade.CallOpts)
}

// KyberNetworkProxyContract is a free data retrieval call binding the contract method 0x785250da.
//
// Solidity: function kyberNetworkProxyContract() constant returns(address)
func (_KBNTrade *KBNTradeCallerSession) KyberNetworkProxyContract() (common.Address, error) {
	return _KBNTrade.Contract.KyberNetworkProxyContract(&_KBNTrade.CallOpts)
}

// Trade is a paid mutator transaction binding the contract method 0xbb39a960.
//
// Solidity: function trade(address srcToken, uint256 srcQty, address destToken, uint256 minConversionRate) returns(address, uint256)
func (_KBNTrade *KBNTradeTransactor) Trade(opts *bind.TransactOpts, srcToken common.Address, srcQty *big.Int, destToken common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KBNTrade.contract.Transact(opts, "trade", srcToken, srcQty, destToken, minConversionRate)
}

// Trade is a paid mutator transaction binding the contract method 0xbb39a960.
//
// Solidity: function trade(address srcToken, uint256 srcQty, address destToken, uint256 minConversionRate) returns(address, uint256)
func (_KBNTrade *KBNTradeSession) Trade(srcToken common.Address, srcQty *big.Int, destToken common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KBNTrade.Contract.Trade(&_KBNTrade.TransactOpts, srcToken, srcQty, destToken, minConversionRate)
}

// Trade is a paid mutator transaction binding the contract method 0xbb39a960.
//
// Solidity: function trade(address srcToken, uint256 srcQty, address destToken, uint256 minConversionRate) returns(address, uint256)
func (_KBNTrade *KBNTradeTransactorSession) Trade(srcToken common.Address, srcQty *big.Int, destToken common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KBNTrade.Contract.Trade(&_KBNTrade.TransactOpts, srcToken, srcQty, destToken, minConversionRate)
}

// KyberNetworkABI is the input ABI used to generate the binding from.
const KyberNetworkABI = "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcQty\",\"type\":\"uint256\"}],\"name\":\"getExpectedRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slippageRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minConversionRate\",\"type\":\"uint256\"}],\"name\":\"swapEtherToToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minConversionRate\",\"type\":\"uint256\"}],\"name\":\"swapTokenToEther\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minConversionRate\",\"type\":\"uint256\"}],\"name\":\"swapTokenToToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxDestAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minConversionRate\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"walletId\",\"type\":\"address\"}],\"name\":\"trade\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// KyberNetworkFuncSigs maps the 4-byte function signature to its string representation.
var KyberNetworkFuncSigs = map[string]string{
	"809a9e55": "getExpectedRate(address,address,uint256)",
	"7a2a0456": "swapEtherToToken(address,uint256)",
	"3bba21dc": "swapTokenToEther(address,uint256,uint256)",
	"7409e2eb": "swapTokenToToken(address,uint256,address,uint256)",
	"cb3c28c7": "trade(address,uint256,address,address,uint256,uint256,address)",
}

// KyberNetwork is an auto generated Go binding around an Ethereum contract.
type KyberNetwork struct {
	KyberNetworkCaller     // Read-only binding to the contract
	KyberNetworkTransactor // Write-only binding to the contract
	KyberNetworkFilterer   // Log filterer for contract events
}

// KyberNetworkCaller is an auto generated read-only Go binding around an Ethereum contract.
type KyberNetworkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberNetworkTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KyberNetworkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberNetworkFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KyberNetworkFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberNetworkSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KyberNetworkSession struct {
	Contract     *KyberNetwork     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KyberNetworkCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KyberNetworkCallerSession struct {
	Contract *KyberNetworkCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// KyberNetworkTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KyberNetworkTransactorSession struct {
	Contract     *KyberNetworkTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// KyberNetworkRaw is an auto generated low-level Go binding around an Ethereum contract.
type KyberNetworkRaw struct {
	Contract *KyberNetwork // Generic contract binding to access the raw methods on
}

// KyberNetworkCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KyberNetworkCallerRaw struct {
	Contract *KyberNetworkCaller // Generic read-only contract binding to access the raw methods on
}

// KyberNetworkTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KyberNetworkTransactorRaw struct {
	Contract *KyberNetworkTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKyberNetwork creates a new instance of KyberNetwork, bound to a specific deployed contract.
func NewKyberNetwork(address common.Address, backend bind.ContractBackend) (*KyberNetwork, error) {
	contract, err := bindKyberNetwork(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KyberNetwork{KyberNetworkCaller: KyberNetworkCaller{contract: contract}, KyberNetworkTransactor: KyberNetworkTransactor{contract: contract}, KyberNetworkFilterer: KyberNetworkFilterer{contract: contract}}, nil
}

// NewKyberNetworkCaller creates a new read-only instance of KyberNetwork, bound to a specific deployed contract.
func NewKyberNetworkCaller(address common.Address, caller bind.ContractCaller) (*KyberNetworkCaller, error) {
	contract, err := bindKyberNetwork(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KyberNetworkCaller{contract: contract}, nil
}

// NewKyberNetworkTransactor creates a new write-only instance of KyberNetwork, bound to a specific deployed contract.
func NewKyberNetworkTransactor(address common.Address, transactor bind.ContractTransactor) (*KyberNetworkTransactor, error) {
	contract, err := bindKyberNetwork(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KyberNetworkTransactor{contract: contract}, nil
}

// NewKyberNetworkFilterer creates a new log filterer instance of KyberNetwork, bound to a specific deployed contract.
func NewKyberNetworkFilterer(address common.Address, filterer bind.ContractFilterer) (*KyberNetworkFilterer, error) {
	contract, err := bindKyberNetwork(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KyberNetworkFilterer{contract: contract}, nil
}

// bindKyberNetwork binds a generic wrapper to an already deployed contract.
func bindKyberNetwork(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KyberNetworkABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KyberNetwork *KyberNetworkRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KyberNetwork.Contract.KyberNetworkCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KyberNetwork *KyberNetworkRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KyberNetwork.Contract.KyberNetworkTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KyberNetwork *KyberNetworkRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KyberNetwork.Contract.KyberNetworkTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KyberNetwork *KyberNetworkCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KyberNetwork.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KyberNetwork *KyberNetworkTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KyberNetwork.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KyberNetwork *KyberNetworkTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KyberNetwork.Contract.contract.Transact(opts, method, params...)
}

// GetExpectedRate is a free data retrieval call binding the contract method 0x809a9e55.
//
// Solidity: function getExpectedRate(address src, address dest, uint256 srcQty) constant returns(uint256 expectedRate, uint256 slippageRate)
func (_KyberNetwork *KyberNetworkCaller) GetExpectedRate(opts *bind.CallOpts, src common.Address, dest common.Address, srcQty *big.Int) (struct {
	ExpectedRate *big.Int
	SlippageRate *big.Int
}, error) {
	ret := new(struct {
		ExpectedRate *big.Int
		SlippageRate *big.Int
	})
	out := ret
	err := _KyberNetwork.contract.Call(opts, out, "getExpectedRate", src, dest, srcQty)
	return *ret, err
}

// GetExpectedRate is a free data retrieval call binding the contract method 0x809a9e55.
//
// Solidity: function getExpectedRate(address src, address dest, uint256 srcQty) constant returns(uint256 expectedRate, uint256 slippageRate)
func (_KyberNetwork *KyberNetworkSession) GetExpectedRate(src common.Address, dest common.Address, srcQty *big.Int) (struct {
	ExpectedRate *big.Int
	SlippageRate *big.Int
}, error) {
	return _KyberNetwork.Contract.GetExpectedRate(&_KyberNetwork.CallOpts, src, dest, srcQty)
}

// GetExpectedRate is a free data retrieval call binding the contract method 0x809a9e55.
//
// Solidity: function getExpectedRate(address src, address dest, uint256 srcQty) constant returns(uint256 expectedRate, uint256 slippageRate)
func (_KyberNetwork *KyberNetworkCallerSession) GetExpectedRate(src common.Address, dest common.Address, srcQty *big.Int) (struct {
	ExpectedRate *big.Int
	SlippageRate *big.Int
}, error) {
	return _KyberNetwork.Contract.GetExpectedRate(&_KyberNetwork.CallOpts, src, dest, srcQty)
}

// SwapEtherToToken is a paid mutator transaction binding the contract method 0x7a2a0456.
//
// Solidity: function swapEtherToToken(address token, uint256 minConversionRate) returns(uint256)
func (_KyberNetwork *KyberNetworkTransactor) SwapEtherToToken(opts *bind.TransactOpts, token common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetwork.contract.Transact(opts, "swapEtherToToken", token, minConversionRate)
}

// SwapEtherToToken is a paid mutator transaction binding the contract method 0x7a2a0456.
//
// Solidity: function swapEtherToToken(address token, uint256 minConversionRate) returns(uint256)
func (_KyberNetwork *KyberNetworkSession) SwapEtherToToken(token common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetwork.Contract.SwapEtherToToken(&_KyberNetwork.TransactOpts, token, minConversionRate)
}

// SwapEtherToToken is a paid mutator transaction binding the contract method 0x7a2a0456.
//
// Solidity: function swapEtherToToken(address token, uint256 minConversionRate) returns(uint256)
func (_KyberNetwork *KyberNetworkTransactorSession) SwapEtherToToken(token common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetwork.Contract.SwapEtherToToken(&_KyberNetwork.TransactOpts, token, minConversionRate)
}

// SwapTokenToEther is a paid mutator transaction binding the contract method 0x3bba21dc.
//
// Solidity: function swapTokenToEther(address token, uint256 srcAmount, uint256 minConversionRate) returns(uint256)
func (_KyberNetwork *KyberNetworkTransactor) SwapTokenToEther(opts *bind.TransactOpts, token common.Address, srcAmount *big.Int, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetwork.contract.Transact(opts, "swapTokenToEther", token, srcAmount, minConversionRate)
}

// SwapTokenToEther is a paid mutator transaction binding the contract method 0x3bba21dc.
//
// Solidity: function swapTokenToEther(address token, uint256 srcAmount, uint256 minConversionRate) returns(uint256)
func (_KyberNetwork *KyberNetworkSession) SwapTokenToEther(token common.Address, srcAmount *big.Int, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetwork.Contract.SwapTokenToEther(&_KyberNetwork.TransactOpts, token, srcAmount, minConversionRate)
}

// SwapTokenToEther is a paid mutator transaction binding the contract method 0x3bba21dc.
//
// Solidity: function swapTokenToEther(address token, uint256 srcAmount, uint256 minConversionRate) returns(uint256)
func (_KyberNetwork *KyberNetworkTransactorSession) SwapTokenToEther(token common.Address, srcAmount *big.Int, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetwork.Contract.SwapTokenToEther(&_KyberNetwork.TransactOpts, token, srcAmount, minConversionRate)
}

// SwapTokenToToken is a paid mutator transaction binding the contract method 0x7409e2eb.
//
// Solidity: function swapTokenToToken(address src, uint256 srcAmount, address dest, uint256 minConversionRate) returns(uint256)
func (_KyberNetwork *KyberNetworkTransactor) SwapTokenToToken(opts *bind.TransactOpts, src common.Address, srcAmount *big.Int, dest common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetwork.contract.Transact(opts, "swapTokenToToken", src, srcAmount, dest, minConversionRate)
}

// SwapTokenToToken is a paid mutator transaction binding the contract method 0x7409e2eb.
//
// Solidity: function swapTokenToToken(address src, uint256 srcAmount, address dest, uint256 minConversionRate) returns(uint256)
func (_KyberNetwork *KyberNetworkSession) SwapTokenToToken(src common.Address, srcAmount *big.Int, dest common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetwork.Contract.SwapTokenToToken(&_KyberNetwork.TransactOpts, src, srcAmount, dest, minConversionRate)
}

// SwapTokenToToken is a paid mutator transaction binding the contract method 0x7409e2eb.
//
// Solidity: function swapTokenToToken(address src, uint256 srcAmount, address dest, uint256 minConversionRate) returns(uint256)
func (_KyberNetwork *KyberNetworkTransactorSession) SwapTokenToToken(src common.Address, srcAmount *big.Int, dest common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetwork.Contract.SwapTokenToToken(&_KyberNetwork.TransactOpts, src, srcAmount, dest, minConversionRate)
}

// Trade is a paid mutator transaction binding the contract method 0xcb3c28c7.
//
// Solidity: function trade(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId) returns(uint256)
func (_KyberNetwork *KyberNetworkTransactor) Trade(opts *bind.TransactOpts, src common.Address, srcAmount *big.Int, dest common.Address, destAddress common.Address, maxDestAmount *big.Int, minConversionRate *big.Int, walletId common.Address) (*types.Transaction, error) {
	return _KyberNetwork.contract.Transact(opts, "trade", src, srcAmount, dest, destAddress, maxDestAmount, minConversionRate, walletId)
}

// Trade is a paid mutator transaction binding the contract method 0xcb3c28c7.
//
// Solidity: function trade(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId) returns(uint256)
func (_KyberNetwork *KyberNetworkSession) Trade(src common.Address, srcAmount *big.Int, dest common.Address, destAddress common.Address, maxDestAmount *big.Int, minConversionRate *big.Int, walletId common.Address) (*types.Transaction, error) {
	return _KyberNetwork.Contract.Trade(&_KyberNetwork.TransactOpts, src, srcAmount, dest, destAddress, maxDestAmount, minConversionRate, walletId)
}

// Trade is a paid mutator transaction binding the contract method 0xcb3c28c7.
//
// Solidity: function trade(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId) returns(uint256)
func (_KyberNetwork *KyberNetworkTransactorSession) Trade(src common.Address, srcAmount *big.Int, dest common.Address, destAddress common.Address, maxDestAmount *big.Int, minConversionRate *big.Int, walletId common.Address) (*types.Transaction, error) {
	return _KyberNetwork.Contract.Trade(&_KyberNetwork.TransactOpts, src, srcAmount, dest, destAddress, maxDestAmount, minConversionRate, walletId)
}

// TradeUtilsABI is the input ABI used to generate the binding from.
const TradeUtilsABI = "[{\"inputs\":[],\"name\":\"ETH_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incognitoSmartContract\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TradeUtilsFuncSigs maps the 4-byte function signature to its string representation.
var TradeUtilsFuncSigs = map[string]string{
	"72e94bf6": "ETH_CONTRACT_ADDRESS()",
	"b42a644b": "incognitoSmartContract()",
}

// TradeUtilsBin is the compiled bytecode used for deploying new contracts.
var TradeUtilsBin = "0x6080604052348015600f57600080fd5b5060a98061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c806372e94bf6146037578063b42a644b146059575b600080fd5b603d605f565b604080516001600160a01b039092168252519081900360200190f35b603d6064565b600081565b6000546001600160a01b03168156fea26469706673582212203aef2a1e85171534925f45f87d8dc763082616487e030a48bce77e9f5ef1fe0964736f6c634300060c0033"

// DeployTradeUtils deploys a new Ethereum contract, binding an instance of TradeUtils to it.
func DeployTradeUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TradeUtils, error) {
	parsed, err := abi.JSON(strings.NewReader(TradeUtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TradeUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TradeUtils{TradeUtilsCaller: TradeUtilsCaller{contract: contract}, TradeUtilsTransactor: TradeUtilsTransactor{contract: contract}, TradeUtilsFilterer: TradeUtilsFilterer{contract: contract}}, nil
}

// TradeUtils is an auto generated Go binding around an Ethereum contract.
type TradeUtils struct {
	TradeUtilsCaller     // Read-only binding to the contract
	TradeUtilsTransactor // Write-only binding to the contract
	TradeUtilsFilterer   // Log filterer for contract events
}

// TradeUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type TradeUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradeUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TradeUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradeUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TradeUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradeUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TradeUtilsSession struct {
	Contract     *TradeUtils       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TradeUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TradeUtilsCallerSession struct {
	Contract *TradeUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TradeUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TradeUtilsTransactorSession struct {
	Contract     *TradeUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TradeUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type TradeUtilsRaw struct {
	Contract *TradeUtils // Generic contract binding to access the raw methods on
}

// TradeUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TradeUtilsCallerRaw struct {
	Contract *TradeUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// TradeUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TradeUtilsTransactorRaw struct {
	Contract *TradeUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTradeUtils creates a new instance of TradeUtils, bound to a specific deployed contract.
func NewTradeUtils(address common.Address, backend bind.ContractBackend) (*TradeUtils, error) {
	contract, err := bindTradeUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TradeUtils{TradeUtilsCaller: TradeUtilsCaller{contract: contract}, TradeUtilsTransactor: TradeUtilsTransactor{contract: contract}, TradeUtilsFilterer: TradeUtilsFilterer{contract: contract}}, nil
}

// NewTradeUtilsCaller creates a new read-only instance of TradeUtils, bound to a specific deployed contract.
func NewTradeUtilsCaller(address common.Address, caller bind.ContractCaller) (*TradeUtilsCaller, error) {
	contract, err := bindTradeUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TradeUtilsCaller{contract: contract}, nil
}

// NewTradeUtilsTransactor creates a new write-only instance of TradeUtils, bound to a specific deployed contract.
func NewTradeUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*TradeUtilsTransactor, error) {
	contract, err := bindTradeUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TradeUtilsTransactor{contract: contract}, nil
}

// NewTradeUtilsFilterer creates a new log filterer instance of TradeUtils, bound to a specific deployed contract.
func NewTradeUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*TradeUtilsFilterer, error) {
	contract, err := bindTradeUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TradeUtilsFilterer{contract: contract}, nil
}

// bindTradeUtils binds a generic wrapper to an already deployed contract.
func bindTradeUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TradeUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TradeUtils *TradeUtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TradeUtils.Contract.TradeUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TradeUtils *TradeUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TradeUtils.Contract.TradeUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TradeUtils *TradeUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TradeUtils.Contract.TradeUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TradeUtils *TradeUtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TradeUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TradeUtils *TradeUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TradeUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TradeUtils *TradeUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TradeUtils.Contract.contract.Transact(opts, method, params...)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_TradeUtils *TradeUtilsCaller) ETHCONTRACTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TradeUtils.contract.Call(opts, out, "ETH_CONTRACT_ADDRESS")
	return *ret0, err
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_TradeUtils *TradeUtilsSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _TradeUtils.Contract.ETHCONTRACTADDRESS(&_TradeUtils.CallOpts)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_TradeUtils *TradeUtilsCallerSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _TradeUtils.Contract.ETHCONTRACTADDRESS(&_TradeUtils.CallOpts)
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_TradeUtils *TradeUtilsCaller) IncognitoSmartContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TradeUtils.contract.Call(opts, out, "incognitoSmartContract")
	return *ret0, err
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_TradeUtils *TradeUtilsSession) IncognitoSmartContract() (common.Address, error) {
	return _TradeUtils.Contract.IncognitoSmartContract(&_TradeUtils.CallOpts)
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_TradeUtils *TradeUtilsCallerSession) IncognitoSmartContract() (common.Address, error) {
	return _TradeUtils.Contract.IncognitoSmartContract(&_TradeUtils.CallOpts)
}
