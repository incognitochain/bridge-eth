// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package opensea

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

// OpenseaMetaData contains all meta data concerning the Opensea contract.
var OpenseaMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"callee\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"forward\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610209806100206000396000f3fe60806040526004361061001e5760003560e01c80636fadcf7214610023575b600080fd5b610036610031366004610132565b610059565b604080516001600160a01b03909316835260208301919091520160405180910390f35b6000806000856001600160a01b03163486866040516100799291906101c3565b60006040518083038185875af1925050503d80600081146100b6576040519150601f19603f3d011682016040523d82523d6000602084013e6100bb565b606091505b50509050806101105760405162461bcd60e51b815260206004820181905260248201527f50726f78793a207265717565737420746f206f70656e736561206661696c6564604482015260640160405180910390fd5b50734cb607c24ac252a0ce4b2e987ec4413da0f1e3ae95600095509350505050565b60008060006040848603121561014757600080fd5b83356001600160a01b038116811461015e57600080fd5b9250602084013567ffffffffffffffff8082111561017b57600080fd5b818601915086601f83011261018f57600080fd5b81358181111561019e57600080fd5b8760208285010111156101b057600080fd5b6020830194508093505050509250925092565b818382376000910190815291905056fea26469706673582212205ec92c45e4d69e21fd412b429d87353ed1745bc742986a0f87d233c141393c8c64736f6c63430008110033",
}

// OpenseaABI is the input ABI used to generate the binding from.
// Deprecated: Use OpenseaMetaData.ABI instead.
var OpenseaABI = OpenseaMetaData.ABI

// OpenseaBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OpenseaMetaData.Bin instead.
var OpenseaBin = OpenseaMetaData.Bin

// DeployOpensea deploys a new Ethereum contract, binding an instance of Opensea to it.
func DeployOpensea(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Opensea, error) {
	parsed, err := OpenseaMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OpenseaBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Opensea{OpenseaCaller: OpenseaCaller{contract: contract}, OpenseaTransactor: OpenseaTransactor{contract: contract}, OpenseaFilterer: OpenseaFilterer{contract: contract}}, nil
}

// Opensea is an auto generated Go binding around an Ethereum contract.
type Opensea struct {
	OpenseaCaller     // Read-only binding to the contract
	OpenseaTransactor // Write-only binding to the contract
	OpenseaFilterer   // Log filterer for contract events
}

// OpenseaCaller is an auto generated read-only Go binding around an Ethereum contract.
type OpenseaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpenseaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OpenseaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpenseaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OpenseaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpenseaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OpenseaSession struct {
	Contract     *Opensea          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OpenseaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OpenseaCallerSession struct {
	Contract *OpenseaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OpenseaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OpenseaTransactorSession struct {
	Contract     *OpenseaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OpenseaRaw is an auto generated low-level Go binding around an Ethereum contract.
type OpenseaRaw struct {
	Contract *Opensea // Generic contract binding to access the raw methods on
}

// OpenseaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OpenseaCallerRaw struct {
	Contract *OpenseaCaller // Generic read-only contract binding to access the raw methods on
}

// OpenseaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OpenseaTransactorRaw struct {
	Contract *OpenseaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOpensea creates a new instance of Opensea, bound to a specific deployed contract.
func NewOpensea(address common.Address, backend bind.ContractBackend) (*Opensea, error) {
	contract, err := bindOpensea(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Opensea{OpenseaCaller: OpenseaCaller{contract: contract}, OpenseaTransactor: OpenseaTransactor{contract: contract}, OpenseaFilterer: OpenseaFilterer{contract: contract}}, nil
}

// NewOpenseaCaller creates a new read-only instance of Opensea, bound to a specific deployed contract.
func NewOpenseaCaller(address common.Address, caller bind.ContractCaller) (*OpenseaCaller, error) {
	contract, err := bindOpensea(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OpenseaCaller{contract: contract}, nil
}

// NewOpenseaTransactor creates a new write-only instance of Opensea, bound to a specific deployed contract.
func NewOpenseaTransactor(address common.Address, transactor bind.ContractTransactor) (*OpenseaTransactor, error) {
	contract, err := bindOpensea(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OpenseaTransactor{contract: contract}, nil
}

// NewOpenseaFilterer creates a new log filterer instance of Opensea, bound to a specific deployed contract.
func NewOpenseaFilterer(address common.Address, filterer bind.ContractFilterer) (*OpenseaFilterer, error) {
	contract, err := bindOpensea(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OpenseaFilterer{contract: contract}, nil
}

// bindOpensea binds a generic wrapper to an already deployed contract.
func bindOpensea(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OpenseaABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Opensea *OpenseaRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Opensea.Contract.OpenseaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Opensea *OpenseaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Opensea.Contract.OpenseaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Opensea *OpenseaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Opensea.Contract.OpenseaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Opensea *OpenseaCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Opensea.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Opensea *OpenseaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Opensea.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Opensea *OpenseaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Opensea.Contract.contract.Transact(opts, method, params...)
}

// Forward is a paid mutator transaction binding the contract method 0x6fadcf72.
//
// Solidity: function forward(address callee, bytes message) payable returns(address, uint256)
func (_Opensea *OpenseaTransactor) Forward(opts *bind.TransactOpts, callee common.Address, message []byte) (*types.Transaction, error) {
	return _Opensea.contract.Transact(opts, "forward", callee, message)
}

// Forward is a paid mutator transaction binding the contract method 0x6fadcf72.
//
// Solidity: function forward(address callee, bytes message) payable returns(address, uint256)
func (_Opensea *OpenseaSession) Forward(callee common.Address, message []byte) (*types.Transaction, error) {
	return _Opensea.Contract.Forward(&_Opensea.TransactOpts, callee, message)
}

// Forward is a paid mutator transaction binding the contract method 0x6fadcf72.
//
// Solidity: function forward(address callee, bytes message) payable returns(address, uint256)
func (_Opensea *OpenseaTransactorSession) Forward(callee common.Address, message []byte) (*types.Transaction, error) {
	return _Opensea.Contract.Forward(&_Opensea.TransactOpts, callee, message)
}
