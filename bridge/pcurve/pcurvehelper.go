// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pcurve

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

// PcurvehelperMetaData contains all meta data concerning the Pcurvehelper contract.
var PcurvehelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"j\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"}],\"name\":\"exchange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"j\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"}],\"name\":\"exchange_underlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"j\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"get_dy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"j\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"get_dy_underlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"underlying_coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PcurvehelperABI is the input ABI used to generate the binding from.
// Deprecated: Use PcurvehelperMetaData.ABI instead.
var PcurvehelperABI = PcurvehelperMetaData.ABI

// Pcurvehelper is an auto generated Go binding around an Ethereum contract.
type Pcurvehelper struct {
	PcurvehelperCaller     // Read-only binding to the contract
	PcurvehelperTransactor // Write-only binding to the contract
	PcurvehelperFilterer   // Log filterer for contract events
}

// PcurvehelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type PcurvehelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PcurvehelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PcurvehelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PcurvehelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PcurvehelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PcurvehelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PcurvehelperSession struct {
	Contract     *Pcurvehelper     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PcurvehelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PcurvehelperCallerSession struct {
	Contract *PcurvehelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PcurvehelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PcurvehelperTransactorSession struct {
	Contract     *PcurvehelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PcurvehelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type PcurvehelperRaw struct {
	Contract *Pcurvehelper // Generic contract binding to access the raw methods on
}

// PcurvehelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PcurvehelperCallerRaw struct {
	Contract *PcurvehelperCaller // Generic read-only contract binding to access the raw methods on
}

// PcurvehelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PcurvehelperTransactorRaw struct {
	Contract *PcurvehelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPcurvehelper creates a new instance of Pcurvehelper, bound to a specific deployed contract.
func NewPcurvehelper(address common.Address, backend bind.ContractBackend) (*Pcurvehelper, error) {
	contract, err := bindPcurvehelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pcurvehelper{PcurvehelperCaller: PcurvehelperCaller{contract: contract}, PcurvehelperTransactor: PcurvehelperTransactor{contract: contract}, PcurvehelperFilterer: PcurvehelperFilterer{contract: contract}}, nil
}

// NewPcurvehelperCaller creates a new read-only instance of Pcurvehelper, bound to a specific deployed contract.
func NewPcurvehelperCaller(address common.Address, caller bind.ContractCaller) (*PcurvehelperCaller, error) {
	contract, err := bindPcurvehelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PcurvehelperCaller{contract: contract}, nil
}

// NewPcurvehelperTransactor creates a new write-only instance of Pcurvehelper, bound to a specific deployed contract.
func NewPcurvehelperTransactor(address common.Address, transactor bind.ContractTransactor) (*PcurvehelperTransactor, error) {
	contract, err := bindPcurvehelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PcurvehelperTransactor{contract: contract}, nil
}

// NewPcurvehelperFilterer creates a new log filterer instance of Pcurvehelper, bound to a specific deployed contract.
func NewPcurvehelperFilterer(address common.Address, filterer bind.ContractFilterer) (*PcurvehelperFilterer, error) {
	contract, err := bindPcurvehelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PcurvehelperFilterer{contract: contract}, nil
}

// bindPcurvehelper binds a generic wrapper to an already deployed contract.
func bindPcurvehelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PcurvehelperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pcurvehelper *PcurvehelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pcurvehelper.Contract.PcurvehelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pcurvehelper *PcurvehelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pcurvehelper.Contract.PcurvehelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pcurvehelper *PcurvehelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pcurvehelper.Contract.PcurvehelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pcurvehelper *PcurvehelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pcurvehelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pcurvehelper *PcurvehelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pcurvehelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pcurvehelper *PcurvehelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pcurvehelper.Contract.contract.Transact(opts, method, params...)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 index) view returns(address)
func (_Pcurvehelper *PcurvehelperCaller) Coins(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Pcurvehelper.contract.Call(opts, &out, "coins", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 index) view returns(address)
func (_Pcurvehelper *PcurvehelperSession) Coins(index *big.Int) (common.Address, error) {
	return _Pcurvehelper.Contract.Coins(&_Pcurvehelper.CallOpts, index)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 index) view returns(address)
func (_Pcurvehelper *PcurvehelperCallerSession) Coins(index *big.Int) (common.Address, error) {
	return _Pcurvehelper.Contract.Coins(&_Pcurvehelper.CallOpts, index)
}

// GetDy is a free data retrieval call binding the contract method 0x556d6e9f.
//
// Solidity: function get_dy(uint256 i, uint256 j, uint256 amount) view returns(uint256)
func (_Pcurvehelper *PcurvehelperCaller) GetDy(opts *bind.CallOpts, i *big.Int, j *big.Int, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pcurvehelper.contract.Call(opts, &out, "get_dy", i, j, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy is a free data retrieval call binding the contract method 0x556d6e9f.
//
// Solidity: function get_dy(uint256 i, uint256 j, uint256 amount) view returns(uint256)
func (_Pcurvehelper *PcurvehelperSession) GetDy(i *big.Int, j *big.Int, amount *big.Int) (*big.Int, error) {
	return _Pcurvehelper.Contract.GetDy(&_Pcurvehelper.CallOpts, i, j, amount)
}

// GetDy is a free data retrieval call binding the contract method 0x556d6e9f.
//
// Solidity: function get_dy(uint256 i, uint256 j, uint256 amount) view returns(uint256)
func (_Pcurvehelper *PcurvehelperCallerSession) GetDy(i *big.Int, j *big.Int, amount *big.Int) (*big.Int, error) {
	return _Pcurvehelper.Contract.GetDy(&_Pcurvehelper.CallOpts, i, j, amount)
}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x85f11d1e.
//
// Solidity: function get_dy_underlying(uint256 i, uint256 j, uint256 amount) view returns(uint256)
func (_Pcurvehelper *PcurvehelperCaller) GetDyUnderlying(opts *bind.CallOpts, i *big.Int, j *big.Int, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pcurvehelper.contract.Call(opts, &out, "get_dy_underlying", i, j, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x85f11d1e.
//
// Solidity: function get_dy_underlying(uint256 i, uint256 j, uint256 amount) view returns(uint256)
func (_Pcurvehelper *PcurvehelperSession) GetDyUnderlying(i *big.Int, j *big.Int, amount *big.Int) (*big.Int, error) {
	return _Pcurvehelper.Contract.GetDyUnderlying(&_Pcurvehelper.CallOpts, i, j, amount)
}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x85f11d1e.
//
// Solidity: function get_dy_underlying(uint256 i, uint256 j, uint256 amount) view returns(uint256)
func (_Pcurvehelper *PcurvehelperCallerSession) GetDyUnderlying(i *big.Int, j *big.Int, amount *big.Int) (*big.Int, error) {
	return _Pcurvehelper.Contract.GetDyUnderlying(&_Pcurvehelper.CallOpts, i, j, amount)
}

// UnderlyingCoins is a free data retrieval call binding the contract method 0xb9947eb0.
//
// Solidity: function underlying_coins(uint256 index) view returns(address)
func (_Pcurvehelper *PcurvehelperCaller) UnderlyingCoins(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Pcurvehelper.contract.Call(opts, &out, "underlying_coins", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingCoins is a free data retrieval call binding the contract method 0xb9947eb0.
//
// Solidity: function underlying_coins(uint256 index) view returns(address)
func (_Pcurvehelper *PcurvehelperSession) UnderlyingCoins(index *big.Int) (common.Address, error) {
	return _Pcurvehelper.Contract.UnderlyingCoins(&_Pcurvehelper.CallOpts, index)
}

// UnderlyingCoins is a free data retrieval call binding the contract method 0xb9947eb0.
//
// Solidity: function underlying_coins(uint256 index) view returns(address)
func (_Pcurvehelper *PcurvehelperCallerSession) UnderlyingCoins(index *big.Int) (common.Address, error) {
	return _Pcurvehelper.Contract.UnderlyingCoins(&_Pcurvehelper.CallOpts, index)
}

// Exchange is a paid mutator transaction binding the contract method 0x5b41b908.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 amount, uint256 minAmount) payable returns(uint256)
func (_Pcurvehelper *PcurvehelperTransactor) Exchange(opts *bind.TransactOpts, i *big.Int, j *big.Int, amount *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _Pcurvehelper.contract.Transact(opts, "exchange", i, j, amount, minAmount)
}

// Exchange is a paid mutator transaction binding the contract method 0x5b41b908.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 amount, uint256 minAmount) payable returns(uint256)
func (_Pcurvehelper *PcurvehelperSession) Exchange(i *big.Int, j *big.Int, amount *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _Pcurvehelper.Contract.Exchange(&_Pcurvehelper.TransactOpts, i, j, amount, minAmount)
}

// Exchange is a paid mutator transaction binding the contract method 0x5b41b908.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 amount, uint256 minAmount) payable returns(uint256)
func (_Pcurvehelper *PcurvehelperTransactorSession) Exchange(i *big.Int, j *big.Int, amount *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _Pcurvehelper.Contract.Exchange(&_Pcurvehelper.TransactOpts, i, j, amount, minAmount)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0x65b2489b.
//
// Solidity: function exchange_underlying(uint256 i, uint256 j, uint256 amount, uint256 minAmount) payable returns(uint256)
func (_Pcurvehelper *PcurvehelperTransactor) ExchangeUnderlying(opts *bind.TransactOpts, i *big.Int, j *big.Int, amount *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _Pcurvehelper.contract.Transact(opts, "exchange_underlying", i, j, amount, minAmount)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0x65b2489b.
//
// Solidity: function exchange_underlying(uint256 i, uint256 j, uint256 amount, uint256 minAmount) payable returns(uint256)
func (_Pcurvehelper *PcurvehelperSession) ExchangeUnderlying(i *big.Int, j *big.Int, amount *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _Pcurvehelper.Contract.ExchangeUnderlying(&_Pcurvehelper.TransactOpts, i, j, amount, minAmount)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0x65b2489b.
//
// Solidity: function exchange_underlying(uint256 i, uint256 j, uint256 amount, uint256 minAmount) payable returns(uint256)
func (_Pcurvehelper *PcurvehelperTransactorSession) ExchangeUnderlying(i *big.Int, j *big.Int, amount *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _Pcurvehelper.Contract.ExchangeUnderlying(&_Pcurvehelper.TransactOpts, i, j, amount, minAmount)
}
