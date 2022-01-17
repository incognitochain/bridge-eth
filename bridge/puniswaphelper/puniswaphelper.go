// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pUniswapHelper

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

// IUinswpaHelperExactInputParams is an auto generated low-level Go binding around an user-defined struct.
type IUinswpaHelperExactInputParams struct {
	Path             []byte
	Recipient        common.Address
	Deadline         *big.Int
	AmountIn         *big.Int
	AmountOutMinimum *big.Int
}

// IUinswpaHelperExactInputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IUinswpaHelperExactInputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	Deadline          *big.Int
	AmountIn          *big.Int
	AmountOutMinimum  *big.Int
	SqrtPriceLimitX96 *big.Int
}

// IUinswpaHelperQuoteExactInputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IUinswpaHelperQuoteExactInputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	AmountIn          *big.Int
	Fee               *big.Int
	SqrtPriceLimitX96 *big.Int
}

// IUinswpaHelperQuoteExactOutputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IUinswpaHelperQuoteExactOutputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Amount            *big.Int
	Fee               *big.Int
	SqrtPriceLimitX96 *big.Int
}

// PUniswapHelperMetaData contains all meta data concerning the PUniswapHelper contract.
var PUniswapHelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"}],\"internalType\":\"structIUinswpaHelper.ExactInputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structIUinswpaHelper.ExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"quoteExactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint160[]\",\"name\":\"sqrtPriceX96AfterList\",\"type\":\"uint160[]\"},{\"internalType\":\"uint32[]\",\"name\":\"initializedTicksCrossedList\",\"type\":\"uint32[]\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structIUinswpaHelper.QuoteExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"quoteExactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96After\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"initializedTicksCrossed\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"quoteExactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint160[]\",\"name\":\"sqrtPriceX96AfterList\",\"type\":\"uint160[]\"},{\"internalType\":\"uint32[]\",\"name\":\"initializedTicksCrossedList\",\"type\":\"uint32[]\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structIUinswpaHelper.QuoteExactOutputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"quoteExactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96After\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"initializedTicksCrossed\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PUniswapHelperABI is the input ABI used to generate the binding from.
// Deprecated: Use PUniswapHelperMetaData.ABI instead.
var PUniswapHelperABI = PUniswapHelperMetaData.ABI

// PUniswapHelper is an auto generated Go binding around an Ethereum contract.
type PUniswapHelper struct {
	PUniswapHelperCaller     // Read-only binding to the contract
	PUniswapHelperTransactor // Write-only binding to the contract
	PUniswapHelperFilterer   // Log filterer for contract events
}

// PUniswapHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type PUniswapHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PUniswapHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PUniswapHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PUniswapHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PUniswapHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PUniswapHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PUniswapHelperSession struct {
	Contract     *PUniswapHelper   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PUniswapHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PUniswapHelperCallerSession struct {
	Contract *PUniswapHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PUniswapHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PUniswapHelperTransactorSession struct {
	Contract     *PUniswapHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PUniswapHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type PUniswapHelperRaw struct {
	Contract *PUniswapHelper // Generic contract binding to access the raw methods on
}

// PUniswapHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PUniswapHelperCallerRaw struct {
	Contract *PUniswapHelperCaller // Generic read-only contract binding to access the raw methods on
}

// PUniswapHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PUniswapHelperTransactorRaw struct {
	Contract *PUniswapHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPUniswapHelper creates a new instance of PUniswapHelper, bound to a specific deployed contract.
func NewPUniswapHelper(address common.Address, backend bind.ContractBackend) (*PUniswapHelper, error) {
	contract, err := bindPUniswapHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PUniswapHelper{PUniswapHelperCaller: PUniswapHelperCaller{contract: contract}, PUniswapHelperTransactor: PUniswapHelperTransactor{contract: contract}, PUniswapHelperFilterer: PUniswapHelperFilterer{contract: contract}}, nil
}

// NewPUniswapHelperCaller creates a new read-only instance of PUniswapHelper, bound to a specific deployed contract.
func NewPUniswapHelperCaller(address common.Address, caller bind.ContractCaller) (*PUniswapHelperCaller, error) {
	contract, err := bindPUniswapHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PUniswapHelperCaller{contract: contract}, nil
}

// NewPUniswapHelperTransactor creates a new write-only instance of PUniswapHelper, bound to a specific deployed contract.
func NewPUniswapHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*PUniswapHelperTransactor, error) {
	contract, err := bindPUniswapHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PUniswapHelperTransactor{contract: contract}, nil
}

// NewPUniswapHelperFilterer creates a new log filterer instance of PUniswapHelper, bound to a specific deployed contract.
func NewPUniswapHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*PUniswapHelperFilterer, error) {
	contract, err := bindPUniswapHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PUniswapHelperFilterer{contract: contract}, nil
}

// bindPUniswapHelper binds a generic wrapper to an already deployed contract.
func bindPUniswapHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PUniswapHelperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PUniswapHelper *PUniswapHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PUniswapHelper.Contract.PUniswapHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PUniswapHelper *PUniswapHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PUniswapHelper.Contract.PUniswapHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PUniswapHelper *PUniswapHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PUniswapHelper.Contract.PUniswapHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PUniswapHelper *PUniswapHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PUniswapHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PUniswapHelper *PUniswapHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PUniswapHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PUniswapHelper *PUniswapHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PUniswapHelper.Contract.contract.Transact(opts, method, params...)
}

// QuoteExactInput is a free data retrieval call binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) view returns(uint256 amountOut, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_PUniswapHelper *PUniswapHelperCaller) QuoteExactInput(opts *bind.CallOpts, path []byte, amountIn *big.Int) (struct {
	AmountOut                   *big.Int
	SqrtPriceX96AfterList       []*big.Int
	InitializedTicksCrossedList []uint32
	GasEstimate                 *big.Int
}, error) {
	var out []interface{}
	err := _PUniswapHelper.contract.Call(opts, &out, "quoteExactInput", path, amountIn)

	outstruct := new(struct {
		AmountOut                   *big.Int
		SqrtPriceX96AfterList       []*big.Int
		InitializedTicksCrossedList []uint32
		GasEstimate                 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountOut = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SqrtPriceX96AfterList = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.InitializedTicksCrossedList = *abi.ConvertType(out[2], new([]uint32)).(*[]uint32)
	outstruct.GasEstimate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// QuoteExactInput is a free data retrieval call binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) view returns(uint256 amountOut, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_PUniswapHelper *PUniswapHelperSession) QuoteExactInput(path []byte, amountIn *big.Int) (struct {
	AmountOut                   *big.Int
	SqrtPriceX96AfterList       []*big.Int
	InitializedTicksCrossedList []uint32
	GasEstimate                 *big.Int
}, error) {
	return _PUniswapHelper.Contract.QuoteExactInput(&_PUniswapHelper.CallOpts, path, amountIn)
}

// QuoteExactInput is a free data retrieval call binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) view returns(uint256 amountOut, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_PUniswapHelper *PUniswapHelperCallerSession) QuoteExactInput(path []byte, amountIn *big.Int) (struct {
	AmountOut                   *big.Int
	SqrtPriceX96AfterList       []*big.Int
	InitializedTicksCrossedList []uint32
	GasEstimate                 *big.Int
}, error) {
	return _PUniswapHelper.Contract.QuoteExactInput(&_PUniswapHelper.CallOpts, path, amountIn)
}

// QuoteExactInputSingle is a free data retrieval call binding the contract method 0xc6a5026a.
//
// Solidity: function quoteExactInputSingle((address,address,uint256,uint24,uint160) params) view returns(uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_PUniswapHelper *PUniswapHelperCaller) QuoteExactInputSingle(opts *bind.CallOpts, params IUinswpaHelperQuoteExactInputSingleParams) (struct {
	AmountOut               *big.Int
	SqrtPriceX96After       *big.Int
	InitializedTicksCrossed uint32
	GasEstimate             *big.Int
}, error) {
	var out []interface{}
	err := _PUniswapHelper.contract.Call(opts, &out, "quoteExactInputSingle", params)

	outstruct := new(struct {
		AmountOut               *big.Int
		SqrtPriceX96After       *big.Int
		InitializedTicksCrossed uint32
		GasEstimate             *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountOut = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SqrtPriceX96After = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.InitializedTicksCrossed = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.GasEstimate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// QuoteExactInputSingle is a free data retrieval call binding the contract method 0xc6a5026a.
//
// Solidity: function quoteExactInputSingle((address,address,uint256,uint24,uint160) params) view returns(uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_PUniswapHelper *PUniswapHelperSession) QuoteExactInputSingle(params IUinswpaHelperQuoteExactInputSingleParams) (struct {
	AmountOut               *big.Int
	SqrtPriceX96After       *big.Int
	InitializedTicksCrossed uint32
	GasEstimate             *big.Int
}, error) {
	return _PUniswapHelper.Contract.QuoteExactInputSingle(&_PUniswapHelper.CallOpts, params)
}

// QuoteExactInputSingle is a free data retrieval call binding the contract method 0xc6a5026a.
//
// Solidity: function quoteExactInputSingle((address,address,uint256,uint24,uint160) params) view returns(uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_PUniswapHelper *PUniswapHelperCallerSession) QuoteExactInputSingle(params IUinswpaHelperQuoteExactInputSingleParams) (struct {
	AmountOut               *big.Int
	SqrtPriceX96After       *big.Int
	InitializedTicksCrossed uint32
	GasEstimate             *big.Int
}, error) {
	return _PUniswapHelper.Contract.QuoteExactInputSingle(&_PUniswapHelper.CallOpts, params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_PUniswapHelper *PUniswapHelperTransactor) ExactInput(opts *bind.TransactOpts, params IUinswpaHelperExactInputParams) (*types.Transaction, error) {
	return _PUniswapHelper.contract.Transact(opts, "exactInput", params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_PUniswapHelper *PUniswapHelperSession) ExactInput(params IUinswpaHelperExactInputParams) (*types.Transaction, error) {
	return _PUniswapHelper.Contract.ExactInput(&_PUniswapHelper.TransactOpts, params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_PUniswapHelper *PUniswapHelperTransactorSession) ExactInput(params IUinswpaHelperExactInputParams) (*types.Transaction, error) {
	return _PUniswapHelper.Contract.ExactInput(&_PUniswapHelper.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x414bf389.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_PUniswapHelper *PUniswapHelperTransactor) ExactInputSingle(opts *bind.TransactOpts, params IUinswpaHelperExactInputSingleParams) (*types.Transaction, error) {
	return _PUniswapHelper.contract.Transact(opts, "exactInputSingle", params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x414bf389.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_PUniswapHelper *PUniswapHelperSession) ExactInputSingle(params IUinswpaHelperExactInputSingleParams) (*types.Transaction, error) {
	return _PUniswapHelper.Contract.ExactInputSingle(&_PUniswapHelper.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x414bf389.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_PUniswapHelper *PUniswapHelperTransactorSession) ExactInputSingle(params IUinswpaHelperExactInputSingleParams) (*types.Transaction, error) {
	return _PUniswapHelper.Contract.ExactInputSingle(&_PUniswapHelper.TransactOpts, params)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_PUniswapHelper *PUniswapHelperTransactor) QuoteExactOutput(opts *bind.TransactOpts, path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _PUniswapHelper.contract.Transact(opts, "quoteExactOutput", path, amountOut)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_PUniswapHelper *PUniswapHelperSession) QuoteExactOutput(path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _PUniswapHelper.Contract.QuoteExactOutput(&_PUniswapHelper.TransactOpts, path, amountOut)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_PUniswapHelper *PUniswapHelperTransactorSession) QuoteExactOutput(path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _PUniswapHelper.Contract.QuoteExactOutput(&_PUniswapHelper.TransactOpts, path, amountOut)
}

// QuoteExactOutputSingle is a paid mutator transaction binding the contract method 0xbd21704a.
//
// Solidity: function quoteExactOutputSingle((address,address,uint256,uint24,uint160) params) returns(uint256 amountIn, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_PUniswapHelper *PUniswapHelperTransactor) QuoteExactOutputSingle(opts *bind.TransactOpts, params IUinswpaHelperQuoteExactOutputSingleParams) (*types.Transaction, error) {
	return _PUniswapHelper.contract.Transact(opts, "quoteExactOutputSingle", params)
}

// QuoteExactOutputSingle is a paid mutator transaction binding the contract method 0xbd21704a.
//
// Solidity: function quoteExactOutputSingle((address,address,uint256,uint24,uint160) params) returns(uint256 amountIn, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_PUniswapHelper *PUniswapHelperSession) QuoteExactOutputSingle(params IUinswpaHelperQuoteExactOutputSingleParams) (*types.Transaction, error) {
	return _PUniswapHelper.Contract.QuoteExactOutputSingle(&_PUniswapHelper.TransactOpts, params)
}

// QuoteExactOutputSingle is a paid mutator transaction binding the contract method 0xbd21704a.
//
// Solidity: function quoteExactOutputSingle((address,address,uint256,uint24,uint160) params) returns(uint256 amountIn, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_PUniswapHelper *PUniswapHelperTransactorSession) QuoteExactOutputSingle(params IUinswpaHelperQuoteExactOutputSingleParams) (*types.Transaction, error) {
	return _PUniswapHelper.Contract.QuoteExactOutputSingle(&_PUniswapHelper.TransactOpts, params)
}
