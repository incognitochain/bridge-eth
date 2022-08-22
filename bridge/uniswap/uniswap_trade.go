// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswap

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

// UniswapV2MetaData contains all meta data concerning the UniswapV2 contract.
var UniswapV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"ad5c4648": "WETH()",
		"c45a0155": "factory()",
		"d06ca61f": "getAmountsOut(uint256,address[])",
		"7ff36ab5": "swapExactETHForTokens(uint256,address[],address,uint256)",
		"18cbafe5": "swapExactTokensForETH(uint256,uint256,address[],address,uint256)",
		"38ed1739": "swapExactTokensForTokens(uint256,uint256,address[],address,uint256)",
	},
}

// UniswapV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapV2MetaData.ABI instead.
var UniswapV2ABI = UniswapV2MetaData.ABI

// Deprecated: Use UniswapV2MetaData.Sigs instead.
// UniswapV2FuncSigs maps the 4-byte function signature to its string representation.
var UniswapV2FuncSigs = UniswapV2MetaData.Sigs

// UniswapV2 is an auto generated Go binding around an Ethereum contract.
type UniswapV2 struct {
	UniswapV2Caller     // Read-only binding to the contract
	UniswapV2Transactor // Write-only binding to the contract
	UniswapV2Filterer   // Log filterer for contract events
}

// UniswapV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapV2Session struct {
	Contract     *UniswapV2        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniswapV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapV2CallerSession struct {
	Contract *UniswapV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// UniswapV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapV2TransactorSession struct {
	Contract     *UniswapV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// UniswapV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapV2Raw struct {
	Contract *UniswapV2 // Generic contract binding to access the raw methods on
}

// UniswapV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapV2CallerRaw struct {
	Contract *UniswapV2Caller // Generic read-only contract binding to access the raw methods on
}

// UniswapV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapV2TransactorRaw struct {
	Contract *UniswapV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapV2 creates a new instance of UniswapV2, bound to a specific deployed contract.
func NewUniswapV2(address common.Address, backend bind.ContractBackend) (*UniswapV2, error) {
	contract, err := bindUniswapV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapV2{UniswapV2Caller: UniswapV2Caller{contract: contract}, UniswapV2Transactor: UniswapV2Transactor{contract: contract}, UniswapV2Filterer: UniswapV2Filterer{contract: contract}}, nil
}

// NewUniswapV2Caller creates a new read-only instance of UniswapV2, bound to a specific deployed contract.
func NewUniswapV2Caller(address common.Address, caller bind.ContractCaller) (*UniswapV2Caller, error) {
	contract, err := bindUniswapV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV2Caller{contract: contract}, nil
}

// NewUniswapV2Transactor creates a new write-only instance of UniswapV2, bound to a specific deployed contract.
func NewUniswapV2Transactor(address common.Address, transactor bind.ContractTransactor) (*UniswapV2Transactor, error) {
	contract, err := bindUniswapV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV2Transactor{contract: contract}, nil
}

// NewUniswapV2Filterer creates a new log filterer instance of UniswapV2, bound to a specific deployed contract.
func NewUniswapV2Filterer(address common.Address, filterer bind.ContractFilterer) (*UniswapV2Filterer, error) {
	contract, err := bindUniswapV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapV2Filterer{contract: contract}, nil
}

// bindUniswapV2 binds a generic wrapper to an already deployed contract.
func bindUniswapV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV2 *UniswapV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV2.Contract.UniswapV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV2 *UniswapV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV2.Contract.UniswapV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV2 *UniswapV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV2.Contract.UniswapV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV2 *UniswapV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV2 *UniswapV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV2 *UniswapV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV2.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() pure returns(address)
func (_UniswapV2 *UniswapV2Caller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV2.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() pure returns(address)
func (_UniswapV2 *UniswapV2Session) WETH() (common.Address, error) {
	return _UniswapV2.Contract.WETH(&_UniswapV2.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() pure returns(address)
func (_UniswapV2 *UniswapV2CallerSession) WETH() (common.Address, error) {
	return _UniswapV2.Contract.WETH(&_UniswapV2.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_UniswapV2 *UniswapV2Caller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV2.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_UniswapV2 *UniswapV2Session) Factory() (common.Address, error) {
	return _UniswapV2.Contract.Factory(&_UniswapV2.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_UniswapV2 *UniswapV2CallerSession) Factory() (common.Address, error) {
	return _UniswapV2.Contract.Factory(&_UniswapV2.CallOpts)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_UniswapV2 *UniswapV2Caller) GetAmountsOut(opts *bind.CallOpts, amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _UniswapV2.contract.Call(opts, &out, "getAmountsOut", amountIn, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_UniswapV2 *UniswapV2Session) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _UniswapV2.Contract.GetAmountsOut(&_UniswapV2.CallOpts, amountIn, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_UniswapV2 *UniswapV2CallerSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _UniswapV2.Contract.GetAmountsOut(&_UniswapV2.CallOpts, amountIn, path)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_UniswapV2 *UniswapV2Transactor) SwapExactETHForTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2.contract.Transact(opts, "swapExactETHForTokens", amountOutMin, path, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_UniswapV2 *UniswapV2Session) SwapExactETHForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2.Contract.SwapExactETHForTokens(&_UniswapV2.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_UniswapV2 *UniswapV2TransactorSession) SwapExactETHForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2.Contract.SwapExactETHForTokens(&_UniswapV2.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_UniswapV2 *UniswapV2Transactor) SwapExactTokensForETH(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2.contract.Transact(opts, "swapExactTokensForETH", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_UniswapV2 *UniswapV2Session) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2.Contract.SwapExactTokensForETH(&_UniswapV2.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_UniswapV2 *UniswapV2TransactorSession) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2.Contract.SwapExactTokensForETH(&_UniswapV2.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_UniswapV2 *UniswapV2Transactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_UniswapV2 *UniswapV2Session) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2.Contract.SwapExactTokensForTokens(&_UniswapV2.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_UniswapV2 *UniswapV2TransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2.Contract.SwapExactTokensForTokens(&_UniswapV2.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// UniswapV2TradeMetaData contains all meta data concerning the UniswapV2Trade contract.
var UniswapV2TradeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractUniswapV2\",\"name\":\"_uniswapV2\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ETH_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcQty\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcQty\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"}],\"name\":\"trade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV2\",\"outputs\":[{\"internalType\":\"contractUniswapV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"72e94bf6": "ETH_CONTRACT_ADDRESS()",
		"cba7064f": "getAmountsOut(address,uint256,address)",
		"bb39a960": "trade(address,uint256,address,uint256)",
		"5187c091": "uniswapV2()",
		"f2428621": "wETH()",
	},
	Bin: "0x608060405234801561001057600080fd5b50604051610bec380380610bec8339818101604052602081101561003357600080fd5b5051600080546001600160a01b0319166001600160a01b038084169190911791829055604080516315ab88c960e31b81529051929091169163ad5c464891600480820192602092909190829003018186803b15801561009157600080fd5b505afa1580156100a5573d6000803e3d6000fd5b505050506040513d60208110156100bb57600080fd5b5051600180546001600160a01b0319166001600160a01b0390921691909117905550610b00806100ec6000396000f3fe60806040526004361061004e5760003560e01c80635187c0911461005a57806372e94bf61461008b578063bb39a960146100a0578063cba7064f146100fd578063f24286211461019057610055565b3661005557005b600080fd5b34801561006657600080fd5b5061006f6101a5565b604080516001600160a01b039092168252519081900360200190f35b34801561009757600080fd5b5061006f6101b4565b6100da600480360360808110156100b657600080fd5b506001600160a01b03813581169160208101359160408201351690606001356101b9565b604080516001600160a01b03909316835260208301919091528051918290030190f35b34801561010957600080fd5b506101406004803603606081101561012057600080fd5b506001600160a01b038135811691602081013591604090910135166103e5565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561017c578181015183820152602001610164565b505050509050019250505060405180910390f35b34801561019c57600080fd5b5061006f6105ba565b6000546001600160a01b031681565b600081565b600080846101c6876105c9565b10156101d157600080fd5b836001600160a01b0316866001600160a01b031614156101f057600080fd5b604080516002808252606080830184529260208301908036833701905050905060606001600160a01b038816156102f357878260008151811061022f57fe5b6001600160a01b039283166020918202929092010152600054610255918a91168961065a565b6001600160a01b038616156102a457858260018151811061027257fe5b60200260200101906001600160a01b031690816001600160a01b03168152505061029d828887610766565b90506102ee565b6001805483516001600160a01b0390911691849181106102c057fe5b60200260200101906001600160a01b031690816001600160a01b0316815250506102eb8288876108fe565b90505b61036b565b60015482516001600160a01b0390911690839060009061030f57fe5b60200260200101906001600160a01b031690816001600160a01b031681525050858260018151811061033d57fe5b60200260200101906001600160a01b031690816001600160a01b03168152505061036882888761099a565b90505b60028151101561037a57600080fd5b848160018351038151811061038b57fe5b6020026020010151101580156103b4575086816000815181106103aa57fe5b6020026020010151145b6103bd57600080fd5b85816001835103815181106103ce57fe5b602002602001015193509350505094509492505050565b604080516002808252606080830184529283929190602083019080368337019050509050848160008151811061041757fe5b60200260200101906001600160a01b031690816001600160a01b031681525050828160018151811061044557fe5b6001600160a01b03928316602091820292909201810191909152600080546040805163d06ca61f60e01b8152600481018a815260248201928352875160448301528751939096169563d06ca61f958b958995929493606401928683019202908190849084905b838110156104c35781810151838201526020016104ab565b50505050905001935050505060006040518083038186803b1580156104e757600080fd5b505afa1580156104fb573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561052457600080fd5b810190808051604051939291908464010000000082111561054457600080fd5b90830190602082018581111561055957600080fd5b825186602082028301116401000000008211171561057657600080fd5b82525081516020918201928201910280838360005b838110156105a357818101518382015260200161058b565b505050509050016040525050509150509392505050565b6001546001600160a01b031681565b60006001600160a01b0382166105e0575047610655565b604080516370a0823160e01b815230600482015290516001600160a01b038416916370a08231916024808301926020929190829003018186803b15801561062657600080fd5b505afa15801561063a573d6000803e3d6000fd5b505050506040513d602081101561065057600080fd5b505190505b919050565b6001600160a01b03831615610761576040805163095ea7b360e01b81526001600160a01b03848116600483015260006024830181905292519086169263095ea7b3926044808201939182900301818387803b1580156106b857600080fd5b505af11580156106cc573d6000803e3d6000fd5b505050506106d8610a96565b6106e157600080fd5b826001600160a01b031663095ea7b383836040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050600060405180830381600087803b15801561073857600080fd5b505af115801561074c573d6000803e3d6000fd5b50505050610758610a96565b61076157600080fd5b505050565b606060008054906101000a90046001600160a01b03166001600160a01b03166338ed17398484873342610258016040518663ffffffff1660e01b81526004018086815260200185815260200180602001846001600160a01b03168152602001838152602001828103825285818151815260200191508051906020019060200280838360005b838110156108035781810151838201526020016107eb565b505050509050019650505050505050600060405180830381600087803b15801561082c57600080fd5b505af1158015610840573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561086957600080fd5b810190808051604051939291908464010000000082111561088957600080fd5b90830190602082018581111561089e57600080fd5b82518660208202830111640100000000821117156108bb57600080fd5b82525081516020918201928201910280838360005b838110156108e85781810151838201526020016108d0565b5050505090500160405250505090509392505050565b606060008054906101000a90046001600160a01b03166001600160a01b03166318cbafe58484873342610258016040518663ffffffff1660e01b81526004018086815260200185815260200180602001846001600160a01b0316815260200183815260200182810382528581815181526020019150805190602001906020028083836000838110156108035781810151838201526020016107eb565b606060008054906101000a90046001600160a01b03166001600160a01b0316637ff36ab58484873342610258016040518663ffffffff1660e01b81526004018085815260200180602001846001600160a01b03168152602001838152602001828103825285818151815260200191508051906020019060200280838360005b83811015610a31578181015183820152602001610a19565b50505050905001955050505050506000604051808303818588803b158015610a5857600080fd5b505af1158015610a6c573d6000803e3d6000fd5b50505050506040513d6000823e601f3d908101601f19168201604052602081101561086957600080fd5b6000803d8015610aad5760208114610ab657610ac2565b60019150610ac2565b60206000803e60005191505b50151590509056fea2646970667358221220312842f05194bfba0d1722229c971e9265fc021924dff05da73c73479596f36c64736f6c634300060c0033",
}

// UniswapV2TradeABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapV2TradeMetaData.ABI instead.
var UniswapV2TradeABI = UniswapV2TradeMetaData.ABI

// Deprecated: Use UniswapV2TradeMetaData.Sigs instead.
// UniswapV2TradeFuncSigs maps the 4-byte function signature to its string representation.
var UniswapV2TradeFuncSigs = UniswapV2TradeMetaData.Sigs

// UniswapV2TradeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UniswapV2TradeMetaData.Bin instead.
var UniswapV2TradeBin = UniswapV2TradeMetaData.Bin

// DeployUniswapV2Trade deploys a new Ethereum contract, binding an instance of UniswapV2Trade to it.
func DeployUniswapV2Trade(auth *bind.TransactOpts, backend bind.ContractBackend, _uniswapV2 common.Address) (common.Address, *types.Transaction, *UniswapV2Trade, error) {
	parsed, err := UniswapV2TradeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UniswapV2TradeBin), backend, _uniswapV2)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UniswapV2Trade{UniswapV2TradeCaller: UniswapV2TradeCaller{contract: contract}, UniswapV2TradeTransactor: UniswapV2TradeTransactor{contract: contract}, UniswapV2TradeFilterer: UniswapV2TradeFilterer{contract: contract}}, nil
}

// UniswapV2Trade is an auto generated Go binding around an Ethereum contract.
type UniswapV2Trade struct {
	UniswapV2TradeCaller     // Read-only binding to the contract
	UniswapV2TradeTransactor // Write-only binding to the contract
	UniswapV2TradeFilterer   // Log filterer for contract events
}

// UniswapV2TradeCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapV2TradeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV2TradeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapV2TradeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV2TradeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapV2TradeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV2TradeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapV2TradeSession struct {
	Contract     *UniswapV2Trade   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniswapV2TradeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapV2TradeCallerSession struct {
	Contract *UniswapV2TradeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// UniswapV2TradeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapV2TradeTransactorSession struct {
	Contract     *UniswapV2TradeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// UniswapV2TradeRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapV2TradeRaw struct {
	Contract *UniswapV2Trade // Generic contract binding to access the raw methods on
}

// UniswapV2TradeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapV2TradeCallerRaw struct {
	Contract *UniswapV2TradeCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapV2TradeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapV2TradeTransactorRaw struct {
	Contract *UniswapV2TradeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapV2Trade creates a new instance of UniswapV2Trade, bound to a specific deployed contract.
func NewUniswapV2Trade(address common.Address, backend bind.ContractBackend) (*UniswapV2Trade, error) {
	contract, err := bindUniswapV2Trade(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapV2Trade{UniswapV2TradeCaller: UniswapV2TradeCaller{contract: contract}, UniswapV2TradeTransactor: UniswapV2TradeTransactor{contract: contract}, UniswapV2TradeFilterer: UniswapV2TradeFilterer{contract: contract}}, nil
}

// NewUniswapV2TradeCaller creates a new read-only instance of UniswapV2Trade, bound to a specific deployed contract.
func NewUniswapV2TradeCaller(address common.Address, caller bind.ContractCaller) (*UniswapV2TradeCaller, error) {
	contract, err := bindUniswapV2Trade(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV2TradeCaller{contract: contract}, nil
}

// NewUniswapV2TradeTransactor creates a new write-only instance of UniswapV2Trade, bound to a specific deployed contract.
func NewUniswapV2TradeTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapV2TradeTransactor, error) {
	contract, err := bindUniswapV2Trade(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV2TradeTransactor{contract: contract}, nil
}

// NewUniswapV2TradeFilterer creates a new log filterer instance of UniswapV2Trade, bound to a specific deployed contract.
func NewUniswapV2TradeFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapV2TradeFilterer, error) {
	contract, err := bindUniswapV2Trade(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapV2TradeFilterer{contract: contract}, nil
}

// bindUniswapV2Trade binds a generic wrapper to an already deployed contract.
func bindUniswapV2Trade(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapV2TradeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV2Trade *UniswapV2TradeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV2Trade.Contract.UniswapV2TradeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV2Trade *UniswapV2TradeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV2Trade.Contract.UniswapV2TradeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV2Trade *UniswapV2TradeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV2Trade.Contract.UniswapV2TradeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV2Trade *UniswapV2TradeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV2Trade.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV2Trade *UniswapV2TradeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV2Trade.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV2Trade *UniswapV2TradeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV2Trade.Contract.contract.Transact(opts, method, params...)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() view returns(address)
func (_UniswapV2Trade *UniswapV2TradeCaller) ETHCONTRACTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV2Trade.contract.Call(opts, &out, "ETH_CONTRACT_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() view returns(address)
func (_UniswapV2Trade *UniswapV2TradeSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _UniswapV2Trade.Contract.ETHCONTRACTADDRESS(&_UniswapV2Trade.CallOpts)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() view returns(address)
func (_UniswapV2Trade *UniswapV2TradeCallerSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _UniswapV2Trade.Contract.ETHCONTRACTADDRESS(&_UniswapV2Trade.CallOpts)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xcba7064f.
//
// Solidity: function getAmountsOut(address srcToken, uint256 srcQty, address destToken) view returns(uint256[])
func (_UniswapV2Trade *UniswapV2TradeCaller) GetAmountsOut(opts *bind.CallOpts, srcToken common.Address, srcQty *big.Int, destToken common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _UniswapV2Trade.contract.Call(opts, &out, "getAmountsOut", srcToken, srcQty, destToken)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0xcba7064f.
//
// Solidity: function getAmountsOut(address srcToken, uint256 srcQty, address destToken) view returns(uint256[])
func (_UniswapV2Trade *UniswapV2TradeSession) GetAmountsOut(srcToken common.Address, srcQty *big.Int, destToken common.Address) ([]*big.Int, error) {
	return _UniswapV2Trade.Contract.GetAmountsOut(&_UniswapV2Trade.CallOpts, srcToken, srcQty, destToken)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xcba7064f.
//
// Solidity: function getAmountsOut(address srcToken, uint256 srcQty, address destToken) view returns(uint256[])
func (_UniswapV2Trade *UniswapV2TradeCallerSession) GetAmountsOut(srcToken common.Address, srcQty *big.Int, destToken common.Address) ([]*big.Int, error) {
	return _UniswapV2Trade.Contract.GetAmountsOut(&_UniswapV2Trade.CallOpts, srcToken, srcQty, destToken)
}

// UniswapV2 is a free data retrieval call binding the contract method 0x5187c091.
//
// Solidity: function uniswapV2() view returns(address)
func (_UniswapV2Trade *UniswapV2TradeCaller) UniswapV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV2Trade.contract.Call(opts, &out, "uniswapV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV2 is a free data retrieval call binding the contract method 0x5187c091.
//
// Solidity: function uniswapV2() view returns(address)
func (_UniswapV2Trade *UniswapV2TradeSession) UniswapV2() (common.Address, error) {
	return _UniswapV2Trade.Contract.UniswapV2(&_UniswapV2Trade.CallOpts)
}

// UniswapV2 is a free data retrieval call binding the contract method 0x5187c091.
//
// Solidity: function uniswapV2() view returns(address)
func (_UniswapV2Trade *UniswapV2TradeCallerSession) UniswapV2() (common.Address, error) {
	return _UniswapV2Trade.Contract.UniswapV2(&_UniswapV2Trade.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xf2428621.
//
// Solidity: function wETH() view returns(address)
func (_UniswapV2Trade *UniswapV2TradeCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV2Trade.contract.Call(opts, &out, "wETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xf2428621.
//
// Solidity: function wETH() view returns(address)
func (_UniswapV2Trade *UniswapV2TradeSession) WETH() (common.Address, error) {
	return _UniswapV2Trade.Contract.WETH(&_UniswapV2Trade.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xf2428621.
//
// Solidity: function wETH() view returns(address)
func (_UniswapV2Trade *UniswapV2TradeCallerSession) WETH() (common.Address, error) {
	return _UniswapV2Trade.Contract.WETH(&_UniswapV2Trade.CallOpts)
}

// Trade is a paid mutator transaction binding the contract method 0xbb39a960.
//
// Solidity: function trade(address srcToken, uint256 srcQty, address destToken, uint256 amountOutMin) payable returns(address, uint256)
func (_UniswapV2Trade *UniswapV2TradeTransactor) Trade(opts *bind.TransactOpts, srcToken common.Address, srcQty *big.Int, destToken common.Address, amountOutMin *big.Int) (*types.Transaction, error) {
	return _UniswapV2Trade.contract.Transact(opts, "trade", srcToken, srcQty, destToken, amountOutMin)
}

// Trade is a paid mutator transaction binding the contract method 0xbb39a960.
//
// Solidity: function trade(address srcToken, uint256 srcQty, address destToken, uint256 amountOutMin) payable returns(address, uint256)
func (_UniswapV2Trade *UniswapV2TradeSession) Trade(srcToken common.Address, srcQty *big.Int, destToken common.Address, amountOutMin *big.Int) (*types.Transaction, error) {
	return _UniswapV2Trade.Contract.Trade(&_UniswapV2Trade.TransactOpts, srcToken, srcQty, destToken, amountOutMin)
}

// Trade is a paid mutator transaction binding the contract method 0xbb39a960.
//
// Solidity: function trade(address srcToken, uint256 srcQty, address destToken, uint256 amountOutMin) payable returns(address, uint256)
func (_UniswapV2Trade *UniswapV2TradeTransactorSession) Trade(srcToken common.Address, srcQty *big.Int, destToken common.Address, amountOutMin *big.Int) (*types.Transaction, error) {
	return _UniswapV2Trade.Contract.Trade(&_UniswapV2Trade.TransactOpts, srcToken, srcQty, destToken, amountOutMin)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniswapV2Trade *UniswapV2TradeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV2Trade.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniswapV2Trade *UniswapV2TradeSession) Receive() (*types.Transaction, error) {
	return _UniswapV2Trade.Contract.Receive(&_UniswapV2Trade.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniswapV2Trade *UniswapV2TradeTransactorSession) Receive() (*types.Transaction, error) {
	return _UniswapV2Trade.Contract.Receive(&_UniswapV2Trade.TransactOpts)
}
