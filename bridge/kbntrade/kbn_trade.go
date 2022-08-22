// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kbntrade

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"313ce567": "decimals()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
}

// IERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetaData.ABI instead.
var IERC20ABI = IERC20MetaData.ABI

// Deprecated: Use IERC20MetaData.Sigs instead.
// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = IERC20MetaData.Sigs

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
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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
	event.Raw = log
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
	event.Raw = log
	return event, nil
}

// KBNTradeMetaData contains all meta data concerning the KBNTrade contract.
var KBNTradeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractKyberNetwork\",\"name\":\"_kyberNetworkProxyContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ETH_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcQty\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"destToken\",\"type\":\"address\"}],\"name\":\"getConversionRates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kyberNetworkProxyContract\",\"outputs\":[{\"internalType\":\"contractKyberNetwork\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcQty\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minConversionRate\",\"type\":\"uint256\"}],\"name\":\"trade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"72e94bf6": "ETH_CONTRACT_ADDRESS()",
		"0aea8188": "getConversionRates(address,uint256,address)",
		"785250da": "kyberNetworkProxyContract()",
		"bb39a960": "trade(address,uint256,address,uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b506040516107f73803806107f78339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b0319909216919091179055610792806100656000396000f3fe6080604052600436106100435760003560e01c80630aea81881461004f57806372e94bf6146100ab578063785250da146100dc578063bb39a960146100f15761004a565b3661004a57005b600080fd5b34801561005b57600080fd5b506100926004803603606081101561007257600080fd5b506001600160a01b0381358116916020810135916040909101351661014e565b6040805192835260208301919091528051918290030190f35b3480156100b757600080fd5b506100c06101e8565b604080516001600160a01b039092168252519081900360200190f35b3480156100e857600080fd5b506100c06101ed565b61012b6004803603608081101561010757600080fd5b506001600160a01b03813581169160208101359160408201351690606001356101fc565b604080516001600160a01b03909316835260208301919091528051918290030190f35b600080546040805163809a9e5560e01b81526001600160a01b03878116600483015285811660248301526044820187905282518594919091169263809a9e559260648082019391829003018186803b1580156101a957600080fd5b505afa1580156101bd573d6000803e3d6000fd5b505050506040513d60408110156101d357600080fd5b50805160209091015190969095509350505050565b600081565b6000546001600160a01b031681565b60008084610209876102dd565b101561021457600080fd5b836001600160a01b0316866001600160a01b0316141561023357600080fd5b60006001600160a01b038716156102a45760005461025c9088906001600160a01b03168861036e565b6001600160a01b038516156102885760006102798888888861047a565b1161028357600080fd5b61029f565b6000610295888887610514565b1161029f57600080fd5b6102bb565b60006102b18688876105a5565b116102bb57600080fd5b6102c4856102dd565b90506102d08582610637565b9396939550929350505050565b60006001600160a01b0382166102f4575047610369565b604080516370a0823160e01b815230600482015290516001600160a01b038416916370a08231916024808301926020929190829003018186803b15801561033a57600080fd5b505afa15801561034e573d6000803e3d6000fd5b505050506040513d602081101561036457600080fd5b505190505b919050565b6001600160a01b03831615610475576040805163095ea7b360e01b81526001600160a01b03848116600483015260006024830181905292519086169263095ea7b3926044808201939182900301818387803b1580156103cc57600080fd5b505af11580156103e0573d6000803e3d6000fd5b505050506103ec610728565b6103f557600080fd5b826001600160a01b031663095ea7b383836040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050600060405180830381600087803b15801561044c57600080fd5b505af1158015610460573d6000803e3d6000fd5b5050505061046c610728565b61047557600080fd5b505050565b6000805460408051637409e2eb60e01b81526001600160a01b0388811660048301526024820188905286811660448301526064820186905291519190921691637409e2eb91608480830192602092919082900301818787803b1580156104df57600080fd5b505af11580156104f3573d6000803e3d6000fd5b505050506040513d602081101561050957600080fd5b505195945050505050565b6000805460408051630eee887760e21b81526001600160a01b038781166004830152602482018790526044820186905291519190921691633bba21dc91606480830192602092919082900301818787803b15801561057157600080fd5b505af1158015610585573d6000803e3d6000fd5b505050506040513d602081101561059b57600080fd5b5051949350505050565b6000824710156105b457600080fd5b60005460408051633d15022b60e11b81526001600160a01b0387811660048301526024820186905291519190921691637a2a045691869160448082019260209290919082900301818588803b15801561060c57600080fd5b505af1158015610620573d6000803e3d6000fd5b50505050506040513d602081101561059b57600080fd5b6001600160a01b0382166106ad578047101561065257600080fd5b604051600090339083908381818185875af1925050503d8060008114610694576040519150601f19603f3d011682016040523d82523d6000602084013e610699565b606091505b50509050806106a757600080fd5b50610724565b6040805163a9059cbb60e01b81523360048201526024810183905290516001600160a01b0384169163a9059cbb91604480830192600092919082900301818387803b1580156106fb57600080fd5b505af115801561070f573d6000803e3d6000fd5b5050505061071b610728565b61072457600080fd5b5050565b6000803d801561073f576020811461074857610754565b60019150610754565b60206000803e60005191505b50151590509056fea26469706673582212206b10dd3c84675d44f68b944a8b1d5521d7ff0d97f70f3a374b676fcb66e5a33164736f6c634300060c0033",
}

// KBNTradeABI is the input ABI used to generate the binding from.
// Deprecated: Use KBNTradeMetaData.ABI instead.
var KBNTradeABI = KBNTradeMetaData.ABI

// Deprecated: Use KBNTradeMetaData.Sigs instead.
// KBNTradeFuncSigs maps the 4-byte function signature to its string representation.
var KBNTradeFuncSigs = KBNTradeMetaData.Sigs

// KBNTradeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use KBNTradeMetaData.Bin instead.
var KBNTradeBin = KBNTradeMetaData.Bin

// DeployKBNTrade deploys a new Ethereum contract, binding an instance of KBNTrade to it.
func DeployKBNTrade(auth *bind.TransactOpts, backend bind.ContractBackend, _kyberNetworkProxyContract common.Address) (common.Address, *types.Transaction, *KBNTrade, error) {
	parsed, err := KBNTradeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(KBNTradeBin), backend, _kyberNetworkProxyContract)
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
func (_KBNTrade *KBNTradeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_KBNTrade *KBNTradeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
// Solidity: function ETH_CONTRACT_ADDRESS() view returns(address)
func (_KBNTrade *KBNTradeCaller) ETHCONTRACTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KBNTrade.contract.Call(opts, &out, "ETH_CONTRACT_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() view returns(address)
func (_KBNTrade *KBNTradeSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _KBNTrade.Contract.ETHCONTRACTADDRESS(&_KBNTrade.CallOpts)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() view returns(address)
func (_KBNTrade *KBNTradeCallerSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _KBNTrade.Contract.ETHCONTRACTADDRESS(&_KBNTrade.CallOpts)
}

// GetConversionRates is a free data retrieval call binding the contract method 0x0aea8188.
//
// Solidity: function getConversionRates(address srcToken, uint256 srcQty, address destToken) view returns(uint256, uint256)
func (_KBNTrade *KBNTradeCaller) GetConversionRates(opts *bind.CallOpts, srcToken common.Address, srcQty *big.Int, destToken common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _KBNTrade.contract.Call(opts, &out, "getConversionRates", srcToken, srcQty, destToken)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetConversionRates is a free data retrieval call binding the contract method 0x0aea8188.
//
// Solidity: function getConversionRates(address srcToken, uint256 srcQty, address destToken) view returns(uint256, uint256)
func (_KBNTrade *KBNTradeSession) GetConversionRates(srcToken common.Address, srcQty *big.Int, destToken common.Address) (*big.Int, *big.Int, error) {
	return _KBNTrade.Contract.GetConversionRates(&_KBNTrade.CallOpts, srcToken, srcQty, destToken)
}

// GetConversionRates is a free data retrieval call binding the contract method 0x0aea8188.
//
// Solidity: function getConversionRates(address srcToken, uint256 srcQty, address destToken) view returns(uint256, uint256)
func (_KBNTrade *KBNTradeCallerSession) GetConversionRates(srcToken common.Address, srcQty *big.Int, destToken common.Address) (*big.Int, *big.Int, error) {
	return _KBNTrade.Contract.GetConversionRates(&_KBNTrade.CallOpts, srcToken, srcQty, destToken)
}

// KyberNetworkProxyContract is a free data retrieval call binding the contract method 0x785250da.
//
// Solidity: function kyberNetworkProxyContract() view returns(address)
func (_KBNTrade *KBNTradeCaller) KyberNetworkProxyContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KBNTrade.contract.Call(opts, &out, "kyberNetworkProxyContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KyberNetworkProxyContract is a free data retrieval call binding the contract method 0x785250da.
//
// Solidity: function kyberNetworkProxyContract() view returns(address)
func (_KBNTrade *KBNTradeSession) KyberNetworkProxyContract() (common.Address, error) {
	return _KBNTrade.Contract.KyberNetworkProxyContract(&_KBNTrade.CallOpts)
}

// KyberNetworkProxyContract is a free data retrieval call binding the contract method 0x785250da.
//
// Solidity: function kyberNetworkProxyContract() view returns(address)
func (_KBNTrade *KBNTradeCallerSession) KyberNetworkProxyContract() (common.Address, error) {
	return _KBNTrade.Contract.KyberNetworkProxyContract(&_KBNTrade.CallOpts)
}

// Trade is a paid mutator transaction binding the contract method 0xbb39a960.
//
// Solidity: function trade(address srcToken, uint256 srcQty, address destToken, uint256 minConversionRate) payable returns(address, uint256)
func (_KBNTrade *KBNTradeTransactor) Trade(opts *bind.TransactOpts, srcToken common.Address, srcQty *big.Int, destToken common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KBNTrade.contract.Transact(opts, "trade", srcToken, srcQty, destToken, minConversionRate)
}

// Trade is a paid mutator transaction binding the contract method 0xbb39a960.
//
// Solidity: function trade(address srcToken, uint256 srcQty, address destToken, uint256 minConversionRate) payable returns(address, uint256)
func (_KBNTrade *KBNTradeSession) Trade(srcToken common.Address, srcQty *big.Int, destToken common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KBNTrade.Contract.Trade(&_KBNTrade.TransactOpts, srcToken, srcQty, destToken, minConversionRate)
}

// Trade is a paid mutator transaction binding the contract method 0xbb39a960.
//
// Solidity: function trade(address srcToken, uint256 srcQty, address destToken, uint256 minConversionRate) payable returns(address, uint256)
func (_KBNTrade *KBNTradeTransactorSession) Trade(srcToken common.Address, srcQty *big.Int, destToken common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KBNTrade.Contract.Trade(&_KBNTrade.TransactOpts, srcToken, srcQty, destToken, minConversionRate)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_KBNTrade *KBNTradeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KBNTrade.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_KBNTrade *KBNTradeSession) Receive() (*types.Transaction, error) {
	return _KBNTrade.Contract.Receive(&_KBNTrade.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_KBNTrade *KBNTradeTransactorSession) Receive() (*types.Transaction, error) {
	return _KBNTrade.Contract.Receive(&_KBNTrade.TransactOpts)
}

// KyberNetworkMetaData contains all meta data concerning the KyberNetwork contract.
var KyberNetworkMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcQty\",\"type\":\"uint256\"}],\"name\":\"getExpectedRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slippageRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minConversionRate\",\"type\":\"uint256\"}],\"name\":\"swapEtherToToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minConversionRate\",\"type\":\"uint256\"}],\"name\":\"swapTokenToEther\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minConversionRate\",\"type\":\"uint256\"}],\"name\":\"swapTokenToToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxDestAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minConversionRate\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"walletId\",\"type\":\"address\"}],\"name\":\"trade\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"809a9e55": "getExpectedRate(address,address,uint256)",
		"7a2a0456": "swapEtherToToken(address,uint256)",
		"3bba21dc": "swapTokenToEther(address,uint256,uint256)",
		"7409e2eb": "swapTokenToToken(address,uint256,address,uint256)",
		"cb3c28c7": "trade(address,uint256,address,address,uint256,uint256,address)",
	},
}

// KyberNetworkABI is the input ABI used to generate the binding from.
// Deprecated: Use KyberNetworkMetaData.ABI instead.
var KyberNetworkABI = KyberNetworkMetaData.ABI

// Deprecated: Use KyberNetworkMetaData.Sigs instead.
// KyberNetworkFuncSigs maps the 4-byte function signature to its string representation.
var KyberNetworkFuncSigs = KyberNetworkMetaData.Sigs

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
func (_KyberNetwork *KyberNetworkRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_KyberNetwork *KyberNetworkCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
// Solidity: function getExpectedRate(address src, address dest, uint256 srcQty) view returns(uint256 expectedRate, uint256 slippageRate)
func (_KyberNetwork *KyberNetworkCaller) GetExpectedRate(opts *bind.CallOpts, src common.Address, dest common.Address, srcQty *big.Int) (struct {
	ExpectedRate *big.Int
	SlippageRate *big.Int
}, error) {
	var out []interface{}
	err := _KyberNetwork.contract.Call(opts, &out, "getExpectedRate", src, dest, srcQty)

	outstruct := new(struct {
		ExpectedRate *big.Int
		SlippageRate *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ExpectedRate = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SlippageRate = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetExpectedRate is a free data retrieval call binding the contract method 0x809a9e55.
//
// Solidity: function getExpectedRate(address src, address dest, uint256 srcQty) view returns(uint256 expectedRate, uint256 slippageRate)
func (_KyberNetwork *KyberNetworkSession) GetExpectedRate(src common.Address, dest common.Address, srcQty *big.Int) (struct {
	ExpectedRate *big.Int
	SlippageRate *big.Int
}, error) {
	return _KyberNetwork.Contract.GetExpectedRate(&_KyberNetwork.CallOpts, src, dest, srcQty)
}

// GetExpectedRate is a free data retrieval call binding the contract method 0x809a9e55.
//
// Solidity: function getExpectedRate(address src, address dest, uint256 srcQty) view returns(uint256 expectedRate, uint256 slippageRate)
func (_KyberNetwork *KyberNetworkCallerSession) GetExpectedRate(src common.Address, dest common.Address, srcQty *big.Int) (struct {
	ExpectedRate *big.Int
	SlippageRate *big.Int
}, error) {
	return _KyberNetwork.Contract.GetExpectedRate(&_KyberNetwork.CallOpts, src, dest, srcQty)
}

// SwapEtherToToken is a paid mutator transaction binding the contract method 0x7a2a0456.
//
// Solidity: function swapEtherToToken(address token, uint256 minConversionRate) payable returns(uint256)
func (_KyberNetwork *KyberNetworkTransactor) SwapEtherToToken(opts *bind.TransactOpts, token common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetwork.contract.Transact(opts, "swapEtherToToken", token, minConversionRate)
}

// SwapEtherToToken is a paid mutator transaction binding the contract method 0x7a2a0456.
//
// Solidity: function swapEtherToToken(address token, uint256 minConversionRate) payable returns(uint256)
func (_KyberNetwork *KyberNetworkSession) SwapEtherToToken(token common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetwork.Contract.SwapEtherToToken(&_KyberNetwork.TransactOpts, token, minConversionRate)
}

// SwapEtherToToken is a paid mutator transaction binding the contract method 0x7a2a0456.
//
// Solidity: function swapEtherToToken(address token, uint256 minConversionRate) payable returns(uint256)
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
// Solidity: function trade(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId) payable returns(uint256)
func (_KyberNetwork *KyberNetworkTransactor) Trade(opts *bind.TransactOpts, src common.Address, srcAmount *big.Int, dest common.Address, destAddress common.Address, maxDestAmount *big.Int, minConversionRate *big.Int, walletId common.Address) (*types.Transaction, error) {
	return _KyberNetwork.contract.Transact(opts, "trade", src, srcAmount, dest, destAddress, maxDestAmount, minConversionRate, walletId)
}

// Trade is a paid mutator transaction binding the contract method 0xcb3c28c7.
//
// Solidity: function trade(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId) payable returns(uint256)
func (_KyberNetwork *KyberNetworkSession) Trade(src common.Address, srcAmount *big.Int, dest common.Address, destAddress common.Address, maxDestAmount *big.Int, minConversionRate *big.Int, walletId common.Address) (*types.Transaction, error) {
	return _KyberNetwork.Contract.Trade(&_KyberNetwork.TransactOpts, src, srcAmount, dest, destAddress, maxDestAmount, minConversionRate, walletId)
}

// Trade is a paid mutator transaction binding the contract method 0xcb3c28c7.
//
// Solidity: function trade(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId) payable returns(uint256)
func (_KyberNetwork *KyberNetworkTransactorSession) Trade(src common.Address, srcAmount *big.Int, dest common.Address, destAddress common.Address, maxDestAmount *big.Int, minConversionRate *big.Int, walletId common.Address) (*types.Transaction, error) {
	return _KyberNetwork.Contract.Trade(&_KyberNetwork.TransactOpts, src, srcAmount, dest, destAddress, maxDestAmount, minConversionRate, walletId)
}

// TradeUtilsMetaData contains all meta data concerning the TradeUtils contract.
var TradeUtilsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ETH_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"72e94bf6": "ETH_CONTRACT_ADDRESS()",
	},
	Bin: "0x6080604052348015600f57600080fd5b50608a8061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c806372e94bf614602d575b600080fd5b6033604f565b604080516001600160a01b039092168252519081900360200190f35b60008156fea2646970667358221220b7777546a30e2b201149c5afb6d1db6bd399fcccbc04741c03e2ec6057762ba564736f6c634300060c0033",
}

// TradeUtilsABI is the input ABI used to generate the binding from.
// Deprecated: Use TradeUtilsMetaData.ABI instead.
var TradeUtilsABI = TradeUtilsMetaData.ABI

// Deprecated: Use TradeUtilsMetaData.Sigs instead.
// TradeUtilsFuncSigs maps the 4-byte function signature to its string representation.
var TradeUtilsFuncSigs = TradeUtilsMetaData.Sigs

// TradeUtilsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TradeUtilsMetaData.Bin instead.
var TradeUtilsBin = TradeUtilsMetaData.Bin

// DeployTradeUtils deploys a new Ethereum contract, binding an instance of TradeUtils to it.
func DeployTradeUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TradeUtils, error) {
	parsed, err := TradeUtilsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TradeUtilsBin), backend)
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
func (_TradeUtils *TradeUtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_TradeUtils *TradeUtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
// Solidity: function ETH_CONTRACT_ADDRESS() view returns(address)
func (_TradeUtils *TradeUtilsCaller) ETHCONTRACTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TradeUtils.contract.Call(opts, &out, "ETH_CONTRACT_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() view returns(address)
func (_TradeUtils *TradeUtilsSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _TradeUtils.Contract.ETHCONTRACTADDRESS(&_TradeUtils.CallOpts)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() view returns(address)
func (_TradeUtils *TradeUtilsCallerSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _TradeUtils.Contract.ETHCONTRACTADDRESS(&_TradeUtils.CallOpts)
}
