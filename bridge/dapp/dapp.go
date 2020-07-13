// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dapp

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

// DappABI is the input ABI used to generate the binding from.
const DappABI = "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_vault\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"datacall\",\"type\":\"bytes\"}],\"name\":\"ReEntranceAttack\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"}],\"name\":\"ReturnAmountWithoutTranfer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reentrance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"}],\"name\":\"revertCall\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"}],\"name\":\"simpleCall\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"test_amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DappBin is the compiled bytecode used for deploying new contracts.
var DappBin = "0x608060405234801561001057600080fd5b506040516103493803806103498339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b03199092169190911790556102e4806100656000396000f3fe60806040526004361061006f5760003560e01c806380f9c6851161004e57806380f9c685146100f657806384ea73e71461011c578063babd22a414610142578063fbfa77cf146101c25761006f565b80622a3bb1146100715780636807676314610098578063788c8081146100e1575b005b34801561007d57600080fd5b506100866101f3565b60408051918252519081900360200190f35b6100be600480360360208110156100ae57600080fd5b50356001600160a01b03166101f9565b604080516001600160a01b03909316835260208301919091528051918290030190f35b3480156100ed57600080fd5b50610086610202565b6100be6004803603602081101561010c57600080fd5b50356001600160a01b031661020a565b6100be6004803603602081101561013257600080fd5b50356001600160a01b0316610212565b6100be6004803603604081101561015857600080fd5b6001600160a01b03823516919081019060408101602082013564010000000081111561018357600080fd5b82018360208201111561019557600080fd5b803590602001918460018302840111640100000000831117156101b757600080fd5b509092509050610218565b3480156101ce57600080fd5b506101d76102a0565b604080516001600160a01b039092168252519081900360200190f35b60015481565b90633b9aca0090565b633b9aca0081565b600080600080fd5b90600090565b60008054604051829182916001600160a01b039091169086908690808383808284376040519201945060009350909150508083038183865af19150503d8060008114610280576040519150601f19603f3d011682016040523d82523d6000602084013e610285565b606091505b505090508061029357600080fd5b5093946000945092505050565b6000546001600160a01b03168156fea265627a7a72315820047b424f1699edee0df8ed29dca109a90dabdeba3e2cfedc3038399b061842d364736f6c63430005100032"

// DeployDapp deploys a new Ethereum contract, binding an instance of Dapp to it.
func DeployDapp(auth *bind.TransactOpts, backend bind.ContractBackend, _vault common.Address) (common.Address, *types.Transaction, *Dapp, error) {
	parsed, err := abi.JSON(strings.NewReader(DappABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DappBin), backend, _vault)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Dapp{DappCaller: DappCaller{contract: contract}, DappTransactor: DappTransactor{contract: contract}, DappFilterer: DappFilterer{contract: contract}}, nil
}

// Dapp is an auto generated Go binding around an Ethereum contract.
type Dapp struct {
	DappCaller     // Read-only binding to the contract
	DappTransactor // Write-only binding to the contract
	DappFilterer   // Log filterer for contract events
}

// DappCaller is an auto generated read-only Go binding around an Ethereum contract.
type DappCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DappTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DappTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DappFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DappFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DappSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DappSession struct {
	Contract     *Dapp             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DappCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DappCallerSession struct {
	Contract *DappCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DappTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DappTransactorSession struct {
	Contract     *DappTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DappRaw is an auto generated low-level Go binding around an Ethereum contract.
type DappRaw struct {
	Contract *Dapp // Generic contract binding to access the raw methods on
}

// DappCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DappCallerRaw struct {
	Contract *DappCaller // Generic read-only contract binding to access the raw methods on
}

// DappTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DappTransactorRaw struct {
	Contract *DappTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDapp creates a new instance of Dapp, bound to a specific deployed contract.
func NewDapp(address common.Address, backend bind.ContractBackend) (*Dapp, error) {
	contract, err := bindDapp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dapp{DappCaller: DappCaller{contract: contract}, DappTransactor: DappTransactor{contract: contract}, DappFilterer: DappFilterer{contract: contract}}, nil
}

// NewDappCaller creates a new read-only instance of Dapp, bound to a specific deployed contract.
func NewDappCaller(address common.Address, caller bind.ContractCaller) (*DappCaller, error) {
	contract, err := bindDapp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DappCaller{contract: contract}, nil
}

// NewDappTransactor creates a new write-only instance of Dapp, bound to a specific deployed contract.
func NewDappTransactor(address common.Address, transactor bind.ContractTransactor) (*DappTransactor, error) {
	contract, err := bindDapp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DappTransactor{contract: contract}, nil
}

// NewDappFilterer creates a new log filterer instance of Dapp, bound to a specific deployed contract.
func NewDappFilterer(address common.Address, filterer bind.ContractFilterer) (*DappFilterer, error) {
	contract, err := bindDapp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DappFilterer{contract: contract}, nil
}

// bindDapp binds a generic wrapper to an already deployed contract.
func bindDapp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DappABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dapp *DappRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dapp.Contract.DappCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dapp *DappRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dapp.Contract.DappTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dapp *DappRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dapp.Contract.DappTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dapp *DappCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dapp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dapp *DappTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dapp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dapp *DappTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dapp.Contract.contract.Transact(opts, method, params...)
}

// Reentrance is a free data retrieval call binding the contract method 0x002a3bb1.
//
// Solidity: function reentrance() constant returns(uint256)
func (_Dapp *DappCaller) Reentrance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dapp.contract.Call(opts, out, "reentrance")
	return *ret0, err
}

// Reentrance is a free data retrieval call binding the contract method 0x002a3bb1.
//
// Solidity: function reentrance() constant returns(uint256)
func (_Dapp *DappSession) Reentrance() (*big.Int, error) {
	return _Dapp.Contract.Reentrance(&_Dapp.CallOpts)
}

// Reentrance is a free data retrieval call binding the contract method 0x002a3bb1.
//
// Solidity: function reentrance() constant returns(uint256)
func (_Dapp *DappCallerSession) Reentrance() (*big.Int, error) {
	return _Dapp.Contract.Reentrance(&_Dapp.CallOpts)
}

// TestAmount is a free data retrieval call binding the contract method 0x788c8081.
//
// Solidity: function test_amount() constant returns(uint256)
func (_Dapp *DappCaller) TestAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dapp.contract.Call(opts, out, "test_amount")
	return *ret0, err
}

// TestAmount is a free data retrieval call binding the contract method 0x788c8081.
//
// Solidity: function test_amount() constant returns(uint256)
func (_Dapp *DappSession) TestAmount() (*big.Int, error) {
	return _Dapp.Contract.TestAmount(&_Dapp.CallOpts)
}

// TestAmount is a free data retrieval call binding the contract method 0x788c8081.
//
// Solidity: function test_amount() constant returns(uint256)
func (_Dapp *DappCallerSession) TestAmount() (*big.Int, error) {
	return _Dapp.Contract.TestAmount(&_Dapp.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() constant returns(address)
func (_Dapp *DappCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dapp.contract.Call(opts, out, "vault")
	return *ret0, err
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() constant returns(address)
func (_Dapp *DappSession) Vault() (common.Address, error) {
	return _Dapp.Contract.Vault(&_Dapp.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() constant returns(address)
func (_Dapp *DappCallerSession) Vault() (common.Address, error) {
	return _Dapp.Contract.Vault(&_Dapp.CallOpts)
}

// ReEntranceAttack is a paid mutator transaction binding the contract method 0xbabd22a4.
//
// Solidity: function ReEntranceAttack(address destToken, bytes datacall) returns(address, uint256)
func (_Dapp *DappTransactor) ReEntranceAttack(opts *bind.TransactOpts, destToken common.Address, datacall []byte) (*types.Transaction, error) {
	return _Dapp.contract.Transact(opts, "ReEntranceAttack", destToken, datacall)
}

// ReEntranceAttack is a paid mutator transaction binding the contract method 0xbabd22a4.
//
// Solidity: function ReEntranceAttack(address destToken, bytes datacall) returns(address, uint256)
func (_Dapp *DappSession) ReEntranceAttack(destToken common.Address, datacall []byte) (*types.Transaction, error) {
	return _Dapp.Contract.ReEntranceAttack(&_Dapp.TransactOpts, destToken, datacall)
}

// ReEntranceAttack is a paid mutator transaction binding the contract method 0xbabd22a4.
//
// Solidity: function ReEntranceAttack(address destToken, bytes datacall) returns(address, uint256)
func (_Dapp *DappTransactorSession) ReEntranceAttack(destToken common.Address, datacall []byte) (*types.Transaction, error) {
	return _Dapp.Contract.ReEntranceAttack(&_Dapp.TransactOpts, destToken, datacall)
}

// ReturnAmountWithoutTranfer is a paid mutator transaction binding the contract method 0x68076763.
//
// Solidity: function ReturnAmountWithoutTranfer(address destToken) returns(address, uint256)
func (_Dapp *DappTransactor) ReturnAmountWithoutTranfer(opts *bind.TransactOpts, destToken common.Address) (*types.Transaction, error) {
	return _Dapp.contract.Transact(opts, "ReturnAmountWithoutTranfer", destToken)
}

// ReturnAmountWithoutTranfer is a paid mutator transaction binding the contract method 0x68076763.
//
// Solidity: function ReturnAmountWithoutTranfer(address destToken) returns(address, uint256)
func (_Dapp *DappSession) ReturnAmountWithoutTranfer(destToken common.Address) (*types.Transaction, error) {
	return _Dapp.Contract.ReturnAmountWithoutTranfer(&_Dapp.TransactOpts, destToken)
}

// ReturnAmountWithoutTranfer is a paid mutator transaction binding the contract method 0x68076763.
//
// Solidity: function ReturnAmountWithoutTranfer(address destToken) returns(address, uint256)
func (_Dapp *DappTransactorSession) ReturnAmountWithoutTranfer(destToken common.Address) (*types.Transaction, error) {
	return _Dapp.Contract.ReturnAmountWithoutTranfer(&_Dapp.TransactOpts, destToken)
}

// RevertCall is a paid mutator transaction binding the contract method 0x80f9c685.
//
// Solidity: function revertCall(address destToken) returns(address, uint256)
func (_Dapp *DappTransactor) RevertCall(opts *bind.TransactOpts, destToken common.Address) (*types.Transaction, error) {
	return _Dapp.contract.Transact(opts, "revertCall", destToken)
}

// RevertCall is a paid mutator transaction binding the contract method 0x80f9c685.
//
// Solidity: function revertCall(address destToken) returns(address, uint256)
func (_Dapp *DappSession) RevertCall(destToken common.Address) (*types.Transaction, error) {
	return _Dapp.Contract.RevertCall(&_Dapp.TransactOpts, destToken)
}

// RevertCall is a paid mutator transaction binding the contract method 0x80f9c685.
//
// Solidity: function revertCall(address destToken) returns(address, uint256)
func (_Dapp *DappTransactorSession) RevertCall(destToken common.Address) (*types.Transaction, error) {
	return _Dapp.Contract.RevertCall(&_Dapp.TransactOpts, destToken)
}

// SimpleCall is a paid mutator transaction binding the contract method 0x84ea73e7.
//
// Solidity: function simpleCall(address destToken) returns(address, uint256)
func (_Dapp *DappTransactor) SimpleCall(opts *bind.TransactOpts, destToken common.Address) (*types.Transaction, error) {
	return _Dapp.contract.Transact(opts, "simpleCall", destToken)
}

// SimpleCall is a paid mutator transaction binding the contract method 0x84ea73e7.
//
// Solidity: function simpleCall(address destToken) returns(address, uint256)
func (_Dapp *DappSession) SimpleCall(destToken common.Address) (*types.Transaction, error) {
	return _Dapp.Contract.SimpleCall(&_Dapp.TransactOpts, destToken)
}

// SimpleCall is a paid mutator transaction binding the contract method 0x84ea73e7.
//
// Solidity: function simpleCall(address destToken) returns(address, uint256)
func (_Dapp *DappTransactorSession) SimpleCall(destToken common.Address) (*types.Transaction, error) {
	return _Dapp.Contract.SimpleCall(&_Dapp.TransactOpts, destToken)
}
