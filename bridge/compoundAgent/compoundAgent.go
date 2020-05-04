// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package compoundAgent

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

// CompoundAgentABI is the input ABI used to generate the binding from.
const CompoundAgentABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxyCompound\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_compoundAgentLogic\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"compoundAgentLogic\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"delegateCall\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxyCompound\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_compoundAgentLogic\",\"type\":\"address\"}],\"name\":\"updateAgentLogic\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxyCompound\",\"type\":\"address\"}],\"name\":\"updateProxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// CompoundAgent is an auto generated Go binding around an Ethereum contract.
type CompoundAgent struct {
	CompoundAgentCaller     // Read-only binding to the contract
	CompoundAgentTransactor // Write-only binding to the contract
	CompoundAgentFilterer   // Log filterer for contract events
}

// CompoundAgentCaller is an auto generated read-only Go binding around an Ethereum contract.
type CompoundAgentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundAgentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CompoundAgentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundAgentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CompoundAgentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundAgentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CompoundAgentSession struct {
	Contract     *CompoundAgent    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CompoundAgentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CompoundAgentCallerSession struct {
	Contract *CompoundAgentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// CompoundAgentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CompoundAgentTransactorSession struct {
	Contract     *CompoundAgentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CompoundAgentRaw is an auto generated low-level Go binding around an Ethereum contract.
type CompoundAgentRaw struct {
	Contract *CompoundAgent // Generic contract binding to access the raw methods on
}

// CompoundAgentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CompoundAgentCallerRaw struct {
	Contract *CompoundAgentCaller // Generic read-only contract binding to access the raw methods on
}

// CompoundAgentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CompoundAgentTransactorRaw struct {
	Contract *CompoundAgentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompoundAgent creates a new instance of CompoundAgent, bound to a specific deployed contract.
func NewCompoundAgent(address common.Address, backend bind.ContractBackend) (*CompoundAgent, error) {
	contract, err := bindCompoundAgent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CompoundAgent{CompoundAgentCaller: CompoundAgentCaller{contract: contract}, CompoundAgentTransactor: CompoundAgentTransactor{contract: contract}, CompoundAgentFilterer: CompoundAgentFilterer{contract: contract}}, nil
}

// NewCompoundAgentCaller creates a new read-only instance of CompoundAgent, bound to a specific deployed contract.
func NewCompoundAgentCaller(address common.Address, caller bind.ContractCaller) (*CompoundAgentCaller, error) {
	contract, err := bindCompoundAgent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CompoundAgentCaller{contract: contract}, nil
}

// NewCompoundAgentTransactor creates a new write-only instance of CompoundAgent, bound to a specific deployed contract.
func NewCompoundAgentTransactor(address common.Address, transactor bind.ContractTransactor) (*CompoundAgentTransactor, error) {
	contract, err := bindCompoundAgent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CompoundAgentTransactor{contract: contract}, nil
}

// NewCompoundAgentFilterer creates a new log filterer instance of CompoundAgent, bound to a specific deployed contract.
func NewCompoundAgentFilterer(address common.Address, filterer bind.ContractFilterer) (*CompoundAgentFilterer, error) {
	contract, err := bindCompoundAgent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CompoundAgentFilterer{contract: contract}, nil
}

// bindCompoundAgent binds a generic wrapper to an already deployed contract.
func bindCompoundAgent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CompoundAgentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompoundAgent *CompoundAgentRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CompoundAgent.Contract.CompoundAgentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompoundAgent *CompoundAgentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompoundAgent.Contract.CompoundAgentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompoundAgent *CompoundAgentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompoundAgent.Contract.CompoundAgentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompoundAgent *CompoundAgentCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CompoundAgent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompoundAgent *CompoundAgentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompoundAgent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompoundAgent *CompoundAgentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompoundAgent.Contract.contract.Transact(opts, method, params...)
}

// CompoundAgentLogic is a free data retrieval call binding the contract method 0xfee8a5c1.
//
// Solidity: function compoundAgentLogic() constant returns(address)
func (_CompoundAgent *CompoundAgentCaller) CompoundAgentLogic(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CompoundAgent.contract.Call(opts, out, "compoundAgentLogic")
	return *ret0, err
}

// CompoundAgentLogic is a free data retrieval call binding the contract method 0xfee8a5c1.
//
// Solidity: function compoundAgentLogic() constant returns(address)
func (_CompoundAgent *CompoundAgentSession) CompoundAgentLogic() (common.Address, error) {
	return _CompoundAgent.Contract.CompoundAgentLogic(&_CompoundAgent.CallOpts)
}

// CompoundAgentLogic is a free data retrieval call binding the contract method 0xfee8a5c1.
//
// Solidity: function compoundAgentLogic() constant returns(address)
func (_CompoundAgent *CompoundAgentCallerSession) CompoundAgentLogic() (common.Address, error) {
	return _CompoundAgent.Contract.CompoundAgentLogic(&_CompoundAgent.CallOpts)
}

// ProxyCompound is a free data retrieval call binding the contract method 0x370f81fe.
//
// Solidity: function proxyCompound() constant returns(address)
func (_CompoundAgent *CompoundAgentCaller) ProxyCompound(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CompoundAgent.contract.Call(opts, out, "proxyCompound")
	return *ret0, err
}

// ProxyCompound is a free data retrieval call binding the contract method 0x370f81fe.
//
// Solidity: function proxyCompound() constant returns(address)
func (_CompoundAgent *CompoundAgentSession) ProxyCompound() (common.Address, error) {
	return _CompoundAgent.Contract.ProxyCompound(&_CompoundAgent.CallOpts)
}

// ProxyCompound is a free data retrieval call binding the contract method 0x370f81fe.
//
// Solidity: function proxyCompound() constant returns(address)
func (_CompoundAgent *CompoundAgentCallerSession) ProxyCompound() (common.Address, error) {
	return _CompoundAgent.Contract.ProxyCompound(&_CompoundAgent.CallOpts)
}

// DelegateCall is a paid mutator transaction binding the contract method 0xda67bcc4.
//
// Solidity: function delegateCall(bytes data) returns(bytes)
func (_CompoundAgent *CompoundAgentTransactor) DelegateCall(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _CompoundAgent.contract.Transact(opts, "delegateCall", data)
}

// DelegateCall is a paid mutator transaction binding the contract method 0xda67bcc4.
//
// Solidity: function delegateCall(bytes data) returns(bytes)
func (_CompoundAgent *CompoundAgentSession) DelegateCall(data []byte) (*types.Transaction, error) {
	return _CompoundAgent.Contract.DelegateCall(&_CompoundAgent.TransactOpts, data)
}

// DelegateCall is a paid mutator transaction binding the contract method 0xda67bcc4.
//
// Solidity: function delegateCall(bytes data) returns(bytes)
func (_CompoundAgent *CompoundAgentTransactorSession) DelegateCall(data []byte) (*types.Transaction, error) {
	return _CompoundAgent.Contract.DelegateCall(&_CompoundAgent.TransactOpts, data)
}

// UpdateAgentLogic is a paid mutator transaction binding the contract method 0xea1efaf4.
//
// Solidity: function updateAgentLogic(address _compoundAgentLogic) returns(address, uint256)
func (_CompoundAgent *CompoundAgentTransactor) UpdateAgentLogic(opts *bind.TransactOpts, _compoundAgentLogic common.Address) (*types.Transaction, error) {
	return _CompoundAgent.contract.Transact(opts, "updateAgentLogic", _compoundAgentLogic)
}

// UpdateAgentLogic is a paid mutator transaction binding the contract method 0xea1efaf4.
//
// Solidity: function updateAgentLogic(address _compoundAgentLogic) returns(address, uint256)
func (_CompoundAgent *CompoundAgentSession) UpdateAgentLogic(_compoundAgentLogic common.Address) (*types.Transaction, error) {
	return _CompoundAgent.Contract.UpdateAgentLogic(&_CompoundAgent.TransactOpts, _compoundAgentLogic)
}

// UpdateAgentLogic is a paid mutator transaction binding the contract method 0xea1efaf4.
//
// Solidity: function updateAgentLogic(address _compoundAgentLogic) returns(address, uint256)
func (_CompoundAgent *CompoundAgentTransactorSession) UpdateAgentLogic(_compoundAgentLogic common.Address) (*types.Transaction, error) {
	return _CompoundAgent.Contract.UpdateAgentLogic(&_CompoundAgent.TransactOpts, _compoundAgentLogic)
}

// UpdateProxy is a paid mutator transaction binding the contract method 0x9e955149.
//
// Solidity: function updateProxy(address _proxyCompound) returns(address, uint256)
func (_CompoundAgent *CompoundAgentTransactor) UpdateProxy(opts *bind.TransactOpts, _proxyCompound common.Address) (*types.Transaction, error) {
	return _CompoundAgent.contract.Transact(opts, "updateProxy", _proxyCompound)
}

// UpdateProxy is a paid mutator transaction binding the contract method 0x9e955149.
//
// Solidity: function updateProxy(address _proxyCompound) returns(address, uint256)
func (_CompoundAgent *CompoundAgentSession) UpdateProxy(_proxyCompound common.Address) (*types.Transaction, error) {
	return _CompoundAgent.Contract.UpdateProxy(&_CompoundAgent.TransactOpts, _proxyCompound)
}

// UpdateProxy is a paid mutator transaction binding the contract method 0x9e955149.
//
// Solidity: function updateProxy(address _proxyCompound) returns(address, uint256)
func (_CompoundAgent *CompoundAgentTransactorSession) UpdateProxy(_proxyCompound common.Address) (*types.Transaction, error) {
	return _CompoundAgent.Contract.UpdateProxy(&_CompoundAgent.TransactOpts, _proxyCompound)
}
