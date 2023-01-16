// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blur

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

// ConverstionDetails is an auto generated low-level Go binding around an user-defined struct.
type ConverstionDetails struct {
	ConversionData []byte
}

// ERC20Details is an auto generated low-level Go binding around an user-defined struct.
type ERC20Details struct {
	TokenAddrs []common.Address
	Amounts    []*big.Int
}

// BlurInterfaceMetaData contains all meta data concerning the BlurInterface contract.
var BlurInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"tokenAddrs\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"internalType\":\"structERC20Details\",\"name\":\"erc20Details\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"marketId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"tradeData\",\"type\":\"bytes\"}],\"internalType\":\"structTradeDetails[]\",\"name\":\"tradeDetails\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"conversionData\",\"type\":\"bytes\"}],\"internalType\":\"structConverstionDetails[]\",\"name\":\"converstionDetails\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"dustTokens\",\"type\":\"address[]\"}],\"name\":\"batchBuyWithERC20s\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"marketId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"tradeData\",\"type\":\"bytes\"}],\"internalType\":\"structTradeDetails[]\",\"name\":\"tradeDetails\",\"type\":\"tuple[]\"}],\"name\":\"batchBuyWithETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"sell\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"buy\",\"type\":\"tuple\"}],\"internalType\":\"structExecution[]\",\"name\":\"executions\",\"type\":\"tuple[]\"}],\"name\":\"bulkExecute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"sell\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"buy\",\"type\":\"tuple\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// BlurInterfaceABI is the input ABI used to generate the binding from.
// Deprecated: Use BlurInterfaceMetaData.ABI instead.
var BlurInterfaceABI = BlurInterfaceMetaData.ABI

// BlurInterface is an auto generated Go binding around an Ethereum contract.
type BlurInterface struct {
	BlurInterfaceCaller     // Read-only binding to the contract
	BlurInterfaceTransactor // Write-only binding to the contract
	BlurInterfaceFilterer   // Log filterer for contract events
}

// BlurInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlurInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlurInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlurInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlurInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlurInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlurInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlurInterfaceSession struct {
	Contract     *BlurInterface    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlurInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlurInterfaceCallerSession struct {
	Contract *BlurInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BlurInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlurInterfaceTransactorSession struct {
	Contract     *BlurInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BlurInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlurInterfaceRaw struct {
	Contract *BlurInterface // Generic contract binding to access the raw methods on
}

// BlurInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlurInterfaceCallerRaw struct {
	Contract *BlurInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// BlurInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlurInterfaceTransactorRaw struct {
	Contract *BlurInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlurInterface creates a new instance of BlurInterface, bound to a specific deployed contract.
func NewBlurInterface(address common.Address, backend bind.ContractBackend) (*BlurInterface, error) {
	contract, err := bindBlurInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlurInterface{BlurInterfaceCaller: BlurInterfaceCaller{contract: contract}, BlurInterfaceTransactor: BlurInterfaceTransactor{contract: contract}, BlurInterfaceFilterer: BlurInterfaceFilterer{contract: contract}}, nil
}

// NewBlurInterfaceCaller creates a new read-only instance of BlurInterface, bound to a specific deployed contract.
func NewBlurInterfaceCaller(address common.Address, caller bind.ContractCaller) (*BlurInterfaceCaller, error) {
	contract, err := bindBlurInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlurInterfaceCaller{contract: contract}, nil
}

// NewBlurInterfaceTransactor creates a new write-only instance of BlurInterface, bound to a specific deployed contract.
func NewBlurInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*BlurInterfaceTransactor, error) {
	contract, err := bindBlurInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlurInterfaceTransactor{contract: contract}, nil
}

// NewBlurInterfaceFilterer creates a new log filterer instance of BlurInterface, bound to a specific deployed contract.
func NewBlurInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*BlurInterfaceFilterer, error) {
	contract, err := bindBlurInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlurInterfaceFilterer{contract: contract}, nil
}

// bindBlurInterface binds a generic wrapper to an already deployed contract.
func bindBlurInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BlurInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlurInterface *BlurInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlurInterface.Contract.BlurInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlurInterface *BlurInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlurInterface.Contract.BlurInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlurInterface *BlurInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlurInterface.Contract.BlurInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlurInterface *BlurInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlurInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlurInterface *BlurInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlurInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlurInterface *BlurInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlurInterface.Contract.contract.Transact(opts, method, params...)
}

// BatchBuyWithERC20s is a paid mutator transaction binding the contract method 0x09ba153d.
//
// Solidity: function batchBuyWithERC20s((address[],uint256[]) erc20Details, (uint256,uint256,bytes)[] tradeDetails, (bytes)[] converstionDetails, address[] dustTokens) payable returns()
func (_BlurInterface *BlurInterfaceTransactor) BatchBuyWithERC20s(opts *bind.TransactOpts, erc20Details ERC20Details, tradeDetails []TradeDetails, converstionDetails []ConverstionDetails, dustTokens []common.Address) (*types.Transaction, error) {
	return _BlurInterface.contract.Transact(opts, "batchBuyWithERC20s", erc20Details, tradeDetails, converstionDetails, dustTokens)
}

// BatchBuyWithERC20s is a paid mutator transaction binding the contract method 0x09ba153d.
//
// Solidity: function batchBuyWithERC20s((address[],uint256[]) erc20Details, (uint256,uint256,bytes)[] tradeDetails, (bytes)[] converstionDetails, address[] dustTokens) payable returns()
func (_BlurInterface *BlurInterfaceSession) BatchBuyWithERC20s(erc20Details ERC20Details, tradeDetails []TradeDetails, converstionDetails []ConverstionDetails, dustTokens []common.Address) (*types.Transaction, error) {
	return _BlurInterface.Contract.BatchBuyWithERC20s(&_BlurInterface.TransactOpts, erc20Details, tradeDetails, converstionDetails, dustTokens)
}

// BatchBuyWithERC20s is a paid mutator transaction binding the contract method 0x09ba153d.
//
// Solidity: function batchBuyWithERC20s((address[],uint256[]) erc20Details, (uint256,uint256,bytes)[] tradeDetails, (bytes)[] converstionDetails, address[] dustTokens) payable returns()
func (_BlurInterface *BlurInterfaceTransactorSession) BatchBuyWithERC20s(erc20Details ERC20Details, tradeDetails []TradeDetails, converstionDetails []ConverstionDetails, dustTokens []common.Address) (*types.Transaction, error) {
	return _BlurInterface.Contract.BatchBuyWithERC20s(&_BlurInterface.TransactOpts, erc20Details, tradeDetails, converstionDetails, dustTokens)
}

// BatchBuyWithETH is a paid mutator transaction binding the contract method 0x9a2b8115.
//
// Solidity: function batchBuyWithETH((uint256,uint256,bytes)[] tradeDetails) payable returns()
func (_BlurInterface *BlurInterfaceTransactor) BatchBuyWithETH(opts *bind.TransactOpts, tradeDetails []TradeDetails) (*types.Transaction, error) {
	return _BlurInterface.contract.Transact(opts, "batchBuyWithETH", tradeDetails)
}

// BatchBuyWithETH is a paid mutator transaction binding the contract method 0x9a2b8115.
//
// Solidity: function batchBuyWithETH((uint256,uint256,bytes)[] tradeDetails) payable returns()
func (_BlurInterface *BlurInterfaceSession) BatchBuyWithETH(tradeDetails []TradeDetails) (*types.Transaction, error) {
	return _BlurInterface.Contract.BatchBuyWithETH(&_BlurInterface.TransactOpts, tradeDetails)
}

// BatchBuyWithETH is a paid mutator transaction binding the contract method 0x9a2b8115.
//
// Solidity: function batchBuyWithETH((uint256,uint256,bytes)[] tradeDetails) payable returns()
func (_BlurInterface *BlurInterfaceTransactorSession) BatchBuyWithETH(tradeDetails []TradeDetails) (*types.Transaction, error) {
	return _BlurInterface.Contract.BatchBuyWithETH(&_BlurInterface.TransactOpts, tradeDetails)
}

// BulkExecute is a paid mutator transaction binding the contract method 0xb3be57f8.
//
// Solidity: function bulkExecute((((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256),((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256))[] executions) payable returns()
func (_BlurInterface *BlurInterfaceTransactor) BulkExecute(opts *bind.TransactOpts, executions []Execution) (*types.Transaction, error) {
	return _BlurInterface.contract.Transact(opts, "bulkExecute", executions)
}

// BulkExecute is a paid mutator transaction binding the contract method 0xb3be57f8.
//
// Solidity: function bulkExecute((((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256),((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256))[] executions) payable returns()
func (_BlurInterface *BlurInterfaceSession) BulkExecute(executions []Execution) (*types.Transaction, error) {
	return _BlurInterface.Contract.BulkExecute(&_BlurInterface.TransactOpts, executions)
}

// BulkExecute is a paid mutator transaction binding the contract method 0xb3be57f8.
//
// Solidity: function bulkExecute((((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256),((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256))[] executions) payable returns()
func (_BlurInterface *BlurInterfaceTransactorSession) BulkExecute(executions []Execution) (*types.Transaction, error) {
	return _BlurInterface.Contract.BulkExecute(&_BlurInterface.TransactOpts, executions)
}

// Execute is a paid mutator transaction binding the contract method 0x9a1fc3a7.
//
// Solidity: function execute(((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) sell, ((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) buy) payable returns()
func (_BlurInterface *BlurInterfaceTransactor) Execute(opts *bind.TransactOpts, sell Input, buy Input) (*types.Transaction, error) {
	return _BlurInterface.contract.Transact(opts, "execute", sell, buy)
}

// Execute is a paid mutator transaction binding the contract method 0x9a1fc3a7.
//
// Solidity: function execute(((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) sell, ((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) buy) payable returns()
func (_BlurInterface *BlurInterfaceSession) Execute(sell Input, buy Input) (*types.Transaction, error) {
	return _BlurInterface.Contract.Execute(&_BlurInterface.TransactOpts, sell, buy)
}

// Execute is a paid mutator transaction binding the contract method 0x9a1fc3a7.
//
// Solidity: function execute(((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) sell, ((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) buy) payable returns()
func (_BlurInterface *BlurInterfaceTransactorSession) Execute(sell Input, buy Input) (*types.Transaction, error) {
	return _BlurInterface.Contract.Execute(&_BlurInterface.TransactOpts, sell, buy)
}
