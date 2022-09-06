// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vault

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

// VaultBurnInstData is an auto generated low-level Go binding around an user-defined struct.
type VaultBurnInstData struct {
	Meta   uint8
	Shard  uint8
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Itx    [32]byte
}

// VaultRedepositOptions is an auto generated low-level Go binding around an user-defined struct.
type VaultRedepositOptions struct {
	RedepositToken      common.Address
	RedepositIncAddress []byte
	WithdrawAddress     common.Address
}

// AddressMetaData contains all meta data concerning the Address contract.
var AddressMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122056923f7f32633c310e8d7f6c1b5fdd5c2fb7886b521493707effd7774c4b7fce64736f6c634300060c0033",
}

// AddressABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressMetaData.ABI instead.
var AddressABI = AddressMetaData.ABI

// AddressBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressMetaData.Bin instead.
var AddressBin = AddressMetaData.Bin

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := AddressMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// CountersMetaData contains all meta data concerning the Counters contract.
var CountersMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a14b78cb1265d7d5c5037533362bf7f7d5a737b90423f1a0e4ec6e0beaa907df64736f6c634300060c0033",
}

// CountersABI is the input ABI used to generate the binding from.
// Deprecated: Use CountersMetaData.ABI instead.
var CountersABI = CountersMetaData.ABI

// CountersBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CountersMetaData.Bin instead.
var CountersBin = CountersMetaData.Bin

// DeployCounters deploys a new Ethereum contract, binding an instance of Counters to it.
func DeployCounters(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Counters, error) {
	parsed, err := CountersMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CountersBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Counters{CountersCaller: CountersCaller{contract: contract}, CountersTransactor: CountersTransactor{contract: contract}, CountersFilterer: CountersFilterer{contract: contract}}, nil
}

// Counters is an auto generated Go binding around an Ethereum contract.
type Counters struct {
	CountersCaller     // Read-only binding to the contract
	CountersTransactor // Write-only binding to the contract
	CountersFilterer   // Log filterer for contract events
}

// CountersCaller is an auto generated read-only Go binding around an Ethereum contract.
type CountersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CountersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CountersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CountersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CountersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CountersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CountersSession struct {
	Contract     *Counters         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CountersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CountersCallerSession struct {
	Contract *CountersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CountersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CountersTransactorSession struct {
	Contract     *CountersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CountersRaw is an auto generated low-level Go binding around an Ethereum contract.
type CountersRaw struct {
	Contract *Counters // Generic contract binding to access the raw methods on
}

// CountersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CountersCallerRaw struct {
	Contract *CountersCaller // Generic read-only contract binding to access the raw methods on
}

// CountersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CountersTransactorRaw struct {
	Contract *CountersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCounters creates a new instance of Counters, bound to a specific deployed contract.
func NewCounters(address common.Address, backend bind.ContractBackend) (*Counters, error) {
	contract, err := bindCounters(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Counters{CountersCaller: CountersCaller{contract: contract}, CountersTransactor: CountersTransactor{contract: contract}, CountersFilterer: CountersFilterer{contract: contract}}, nil
}

// NewCountersCaller creates a new read-only instance of Counters, bound to a specific deployed contract.
func NewCountersCaller(address common.Address, caller bind.ContractCaller) (*CountersCaller, error) {
	contract, err := bindCounters(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CountersCaller{contract: contract}, nil
}

// NewCountersTransactor creates a new write-only instance of Counters, bound to a specific deployed contract.
func NewCountersTransactor(address common.Address, transactor bind.ContractTransactor) (*CountersTransactor, error) {
	contract, err := bindCounters(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CountersTransactor{contract: contract}, nil
}

// NewCountersFilterer creates a new log filterer instance of Counters, bound to a specific deployed contract.
func NewCountersFilterer(address common.Address, filterer bind.ContractFilterer) (*CountersFilterer, error) {
	contract, err := bindCounters(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CountersFilterer{contract: contract}, nil
}

// bindCounters binds a generic wrapper to an already deployed contract.
func bindCounters(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CountersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counters *CountersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counters.Contract.CountersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counters *CountersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counters.Contract.CountersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counters *CountersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counters.Contract.CountersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counters *CountersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counters.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counters *CountersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counters.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counters *CountersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counters.Contract.contract.Transact(opts, method, params...)
}

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

// IncognitoMetaData contains all meta data concerning the Incognito contract.
var IncognitoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"name\":\"instructionApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"f65d2116": "instructionApproved(bool,bytes32,uint256,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])",
	},
}

// IncognitoABI is the input ABI used to generate the binding from.
// Deprecated: Use IncognitoMetaData.ABI instead.
var IncognitoABI = IncognitoMetaData.ABI

// Deprecated: Use IncognitoMetaData.Sigs instead.
// IncognitoFuncSigs maps the 4-byte function signature to its string representation.
var IncognitoFuncSigs = IncognitoMetaData.Sigs

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
func (_Incognito *IncognitoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_Incognito *IncognitoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
	var out []interface{}
	err := _Incognito.contract.Call(opts, &out, "instructionApproved", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

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

// SafeMathMetaData contains all meta data concerning the SafeMath contract.
var SafeMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220703b71cb9c1ee26e8419ec9abe19fc4e186c1f845638eb18662a9b5f8336b20c64736f6c634300060c0033",
}

// SafeMathABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeMathMetaData.ABI instead.
var SafeMathABI = SafeMathMetaData.ABI

// SafeMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeMathMetaData.Bin instead.
var SafeMathBin = SafeMathMetaData.Bin

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := SafeMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeMathBin), backend)
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
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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

// VaultMetaData contains all meta data concerning the Vault contract.
var VaultMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositID\",\"type\":\"uint256\"}],\"name\":\"DepositV2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"phaseID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"errorData\",\"type\":\"bytes\"}],\"name\":\"ExecuteFnLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"redepositIncAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"itx\",\"type\":\"bytes32\"}],\"name\":\"Redeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newIncognitoProxy\",\"type\":\"address\"}],\"name\":\"UpdateIncognitoProxy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"UpdateTokenTotal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BURN_CALL_REQUEST_METADATA_TYPE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BURN_REQUEST_METADATA_TYPE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BURN_TO_CONTRACT_REQUEST_METADATA_TYPE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CURRENT_NETWORK_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETH_TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"externalCalldata\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"redepositToken\",\"type\":\"address\"}],\"name\":\"_callExternal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"_transferExternal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"txId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signData\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"txId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signData\",\"type\":\"bytes\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"}],\"name\":\"depositERC20_V2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"}],\"name\":\"deposit_V2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipientToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"exchangeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"timestamp\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signData\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"heights\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"instPaths\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"instPathIsLefts\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"instRoots\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"sigIdxs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"sigVs\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigRs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigSs\",\"type\":\"bytes32[]\"}],\"name\":\"executeWithBurnProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getDepositedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_prevVault\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"isSigDataUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"isWithdrawed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"migration\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"notEntered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"}],\"name\":\"parseBurnInst\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"meta\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"shard\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"itx\",\"type\":\"bytes32\"}],\"internalType\":\"structVault.BurnInstData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"}],\"name\":\"parseCalldataFromBurnInst\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"meta\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"shard\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"itx\",\"type\":\"bytes32\"}],\"internalType\":\"structVault.BurnInstData\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"redepositToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"redepositIncAddress\",\"type\":\"bytes\"},{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"internalType\":\"structVault.RedepositOptions\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prevVault\",\"outputs\":[{\"internalType\":\"contractWithdrawable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"regulator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"timestamp\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"txId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"regulatorSig\",\"type\":\"bytes\"}],\"name\":\"requestWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_regulator\",\"type\":\"address\"}],\"name\":\"setRegulator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"sigDataUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"sigToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"heights\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"instPaths\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"instPathIsLefts\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"instRoots\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"sigIdxs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"sigVs\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigRs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigSs\",\"type\":\"bytes32[]\"}],\"name\":\"submitBurnProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalDepositedToSCAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"updateAssets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"heights\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"instPaths\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"instPathIsLefts\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"instRoots\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"sigIdxs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"sigVs\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigRs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigSs\",\"type\":\"bytes32[]\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawRequests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"withdrawed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"bd835c42": "BURN_CALL_REQUEST_METADATA_TYPE()",
		"568c04fd": "BURN_REQUEST_METADATA_TYPE()",
		"6f2cbc48": "BURN_TO_CONTRACT_REQUEST_METADATA_TYPE()",
		"d7200eb1": "CURRENT_NETWORK_ID()",
		"58bc8337": "ETH_TOKEN()",
		"bda9b509": "_callExternal(address,address,uint256,bytes,address)",
		"145e2a6b": "_transferExternal(address,address,uint256)",
		"70a08231": "balanceOf(address)",
		"c791d705": "deposit(string,bytes32,bytes)",
		"a807b5bb": "depositERC20(address,uint256,string,bytes32,bytes)",
		"5c421b2f": "depositERC20_V2(address,uint256,string)",
		"49c5ba23": "deposit_V2(string)",
		"8588ccd6": "execute(address,uint256,address,address,bytes,bytes,bytes)",
		"3ed1b376": "executeWithBurnProof(bytes,uint256,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])",
		"cf54aaa0": "getDecimals(address)",
		"f75b98ce": "getDepositedBalance(address,address)",
		"c4d66de8": "initialize(address)",
		"392e53cd": "isInitialized()",
		"e4bd7074": "isSigDataUsed(bytes32)",
		"749c5f86": "isWithdrawed(bytes32)",
		"995fac11": "migration(address,address)",
		"a3f5d8cc": "notEntered()",
		"7e16e6e1": "parseBurnInst(bytes)",
		"66945b31": "parseCalldataFromBurnInst(bytes)",
		"fa84702e": "prevVault()",
		"dd8fee14": "regulator()",
		"fee8efda": "requestWithdraw(string,address,uint256,bytes,bytes,bytes32,bytes)",
		"cde0a4f8": "setRegulator(address)",
		"1ea1940e": "sigDataUsed(bytes32)",
		"3fec6b40": "sigToAddress(bytes,bytes32)",
		"73bf9651": "submitBurnProof(bytes,uint256,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])",
		"6304541c": "totalDepositedToSCAmount(address)",
		"1ed4276d": "updateAssets(address[],uint256[])",
		"1beb7de2": "withdraw(bytes,uint256,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])",
		"65b5a00f": "withdrawRequests(address,address)",
		"dca40d9e": "withdrawed(bytes32)",
	},
	Bin: "0x608060405234801561001057600080fd5b50614e5e806100206000396000f3fe6080604052600436106102135760003560e01c80637e16e6e111610118578063cde0a4f8116100a0578063dd8fee141161006f578063dd8fee14146105e6578063e4bd7074146105fb578063f75b98ce1461061b578063fa84702e1461063b578063fee8efda146106505761021a565b8063cde0a4f814610571578063cf54aaa014610591578063d7200eb1146105b1578063dca40d9e146105c65761021a565b8063a807b5bb116100e7578063a807b5bb146104e9578063bd835c4214610509578063bda9b5091461051e578063c4d66de81461053e578063c791d7051461055e5761021a565b80637e16e6e1146104745780638588ccd6146104a1578063995fac11146104b4578063a3f5d8cc146104d45761021a565b806358bc83371161019b57806366945b311161016a57806366945b31146103d05780636f2cbc48146103ff57806370a082311461041457806373bf965114610434578063749c5f86146104545761021a565b806358bc83371461034e5780635c421b2f146103635780636304541c1461038357806365b5a00f146103b05761021a565b8063392e53cd116101e2578063392e53cd146102b75780633ed1b376146102cc5780633fec6b40146102ec57806349c5ba2314610319578063568c04fd1461032c5761021a565b8063145e2a6b1461021f5780631beb7de2146102415780631ea1940e146102615780631ed4276d146102975761021a565b3661021a57005b600080fd5b34801561022b57600080fd5b5061023f61023a366004613dcf565b610670565b005b34801561024d57600080fd5b5061023f61025c3660046142d9565b610759565b34801561026d57600080fd5b5061028161027c3660046140c8565b610bb6565b60405161028e9190614928565b60405180910390f35b3480156102a357600080fd5b506102816102b2366004614050565b610bcb565b3480156102c357600080fd5b50610281610dec565b3480156102d857600080fd5b5061023f6102e736600461411f565b610dfc565b3480156102f857600080fd5b5061030c610307366004614297565b61132a565b60405161028e9190614753565b61023f6103273660046140e0565b6113b9565b34801561033857600080fd5b5061034161149f565b60405161028e9190614d32565b34801561035a57600080fd5b5061030c6114a4565b34801561036f57600080fd5b5061023f61037e366004613f65565b6114a9565b34801561038f57600080fd5b506103a361039e366004613cdf565b61175a565b60405161028e91906149e0565b3480156103bc57600080fd5b506103a36103cb366004613d97565b61176c565b3480156103dc57600080fd5b506103f06103eb3660046140e0565b611789565b60405161028e93929190614c36565b34801561040b57600080fd5b5061034161196c565b34801561042057600080fd5b506103a361042f366004613cdf565b611971565b34801561044057600080fd5b5061023f61044f3660046142d9565b611a45565b34801561046057600080fd5b5061028161046f3660046140c8565b611dbd565b34801561048057600080fd5b5061049461048f36600461425d565b611e74565b60405161028e9190614c28565b61023f6104af366004613e8d565b611ef6565b3480156104c057600080fd5b506102816104cf366004613d97565b61238e565b3480156104e057600080fd5b506102816123ae565b3480156104f557600080fd5b5061023f610504366004613fbe565b6123be565b34801561051557600080fd5b50610341612697565b34801561052a57600080fd5b506103a3610539366004613e0f565b61269c565b34801561054a57600080fd5b5061023f610559366004613cdf565b612842565b61023f61056c3660046144f1565b6128ba565b34801561057d57600080fd5b5061023f61058c366004613cdf565b6129ce565b34801561059d57600080fd5b506103416105ac366004613cdf565b612a4f565b3480156105bd57600080fd5b50610341612b09565b3480156105d257600080fd5b506102816105e13660046140c8565b612b0e565b3480156105f257600080fd5b5061030c612b23565b34801561060757600080fd5b506102816106163660046140c8565b612b32565b34801561062757600080fd5b506103a3610636366004613d97565b612b99565b34801561064757600080fd5b5061030c612cc5565b34801561065c57600080fd5b5061023f61066b36600461440a565b612cd4565b30331461067d6014612f40565b906106a45760405162461bcd60e51b815260040161069b9190614a2f565b60405180910390fd5b506001600160a01b0383166106c2576106bd8282613069565b610754565b60405163a9059cbb60e01b81526001600160a01b0384169063a9059cbb906106f0908590859060040161478b565b600060405180830381600087803b15801561070a57600080fd5b505af115801561071e573d6000803e3d6000fd5b5050505061072a613105565b6107346004612f40565b906107525760405162461bcd60e51b815260040161069b9190614a2f565b505b505050565b600554600160a01b900460ff166107706001612f40565b9061078e5760405162461bcd60e51b815260040161069b9190614a2f565b506005805460ff60a01b191690558951608211156107ac6006612f40565b906107ca5760405162461bcd60e51b815260040161069b9190614a2f565b506107d36139fd565b6107dc8b611e74565b805190915060ff1660f11480156107fa5750806020015160ff166001145b6108046006612f40565b906108225760405162461bcd60e51b815260040161069b9190614a2f565b506108308160a00151611dbd565b1561083b6005612f40565b906108595760405162461bcd60e51b815260040161069b9190614a2f565b5060a081015160009081526020819052604090819020805460ff191660011790558101516001600160a01b03166108e4576040808201516001600160a01b031660009081526004602052205460808201516108b391613139565b4710156108c06007612f40565b906108de5760405162461bcd60e51b815260040161069b9190614a2f565b506109f6565b60006108f38260400151612a4f565b905060098160ff16111561092057608082015161091a9060081960ff841601600a0a61319f565b60808301525b6040808301516001600160a01b0316600090815260046020522054608083015161094991613139565b82604001516001600160a01b03166370a08231306040518263ffffffff1660e01b81526004016109799190614753565b60206040518083038186803b15801561099157600080fd5b505afa1580156109a5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109c99190614567565b10156109d56007612f40565b906109f35760405162461bcd60e51b815260040161069b9190614a2f565b50505b610a088b8b8b8b8b8b8b8b8b8b613201565b60408101516001600160a01b0316610ab157600081606001516001600160a01b03168260800151604051610a3b90614750565b60006040518083038185875af1925050503d8060008114610a78576040519150601f19603f3d011682016040523d82523d6000602084013e610a7d565b606091505b5050905080610a8c6004612f40565b90610aaa5760405162461bcd60e51b815260040161069b9190614a2f565b5050610b4f565b80604001516001600160a01b031663a9059cbb826060015183608001516040518363ffffffff1660e01b8152600401610aeb92919061478b565b600060405180830381600087803b158015610b0557600080fd5b505af1158015610b19573d6000803e3d6000fd5b50505050610b25613105565b610b2f6004612f40565b90610b4d5760405162461bcd60e51b815260040161069b9190614a2f565b505b7f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb816040015182606001518360800151604051610b8e93929190614767565b60405180910390a150506005805460ff60a01b1916600160a01b179055505050505050505050565b60016020526000908152604090205460ff1681565b6005546000906001600160a01b031615801590610bf257506005546001600160a01b031633145b610bfc600c612f40565b90610c1a5760405162461bcd60e51b815260040161069b9190614a2f565b50838214610c28600a612f40565b90610c465760405162461bcd60e51b815260040161069b9190614a2f565b50600560009054906101000a90046001600160a01b03166001600160a01b0316635c975abb6040518163ffffffff1660e01b815260040160206040518083038186803b158015610c9557600080fd5b505afa158015610ca9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ccd91906140ac565b610cd7600d612f40565b90610cf55760405162461bcd60e51b815260040161069b9190614a2f565b5060005b84811015610da357610d5c848483818110610d1057fe5b9050602002013560046000898986818110610d2757fe5b9050602002016020810190610d3c9190613cdf565b6001600160a01b0316815260208101919091526040016000205490613139565b60046000888885818110610d6c57fe5b9050602002016020810190610d819190613cdf565b6001600160a01b03168152602081019190915260400160002055600101610cf9565b507f6a7fbbcddfd518bb8c56b28ac6c7acb0f7ca093ed232eb3306e53d14e469895f85858585604051610dd994939291906148a9565b60405180910390a1506001949350505050565b600554600160a81b900460ff1681565b600554600160a01b900460ff16610e136001612f40565b90610e315760405162461bcd60e51b815260040161069b9190614a2f565b506005805460ff60a01b19169055610e476139fd565b610e4f613a32565b6060610e5b8e8e611789565b925092509250610e6e8360a00151611dbd565b15610e796005612f40565b90610e975760405162461bcd60e51b815260040161069b9190614a2f565b5060a083015160009081526020819052604090819020805460ff191660011790558301516001600160a01b0316610f22576040808401516001600160a01b03166000908152600460205220546080840151610ef191613139565b471015610efe6007612f40565b90610f1c5760405162461bcd60e51b815260040161069b9190614a2f565b50611034565b6000610f318460400151612a4f565b905060098160ff161115610f5e576080840151610f589060081960ff841601600a0a61319f565b60808501525b6040808501516001600160a01b03166000908152600460205220546080850151610f8791613139565b84604001516001600160a01b03166370a08231306040518263ffffffff1660e01b8152600401610fb79190614753565b60206040518083038186803b158015610fcf57600080fd5b505afa158015610fe3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110079190614567565b10156110136007612f40565b906110315760405162461bcd60e51b815260040161069b9190614a2f565b50505b61108a8e8e8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508d8d8d8d8d8d8d8d8d613201565b604080840151606085015160808601518551935163bda9b50960e01b8152309463bda9b509946110c49490939092909188916004016147be565b602060405180830381600087803b1580156110de57600080fd5b505af192505050801561110e575060408051601f3d908101601f1916820190925261110b91810190614567565b60015b6111a7573d80801561113c576040519150601f19603f3d011682016040523d82523d6000602084013e611141565b606091505b5061115e8460400151846020015186608001518760a001516132fc565b7fdbbb883f24557adf486292429863dcfd4ac5d4db168ae94921da8e3d9a95d4168460a00151600083604051611196939291906149e9565b60405180910390a15050505061130a565b60408301516001600160a01b03166111d6576111d183600001518460200151838760a001516132fc565b611305565b8251604080850151905163145e2a6b60e01b8152309263145e2a6b92611200928690600401614767565b600060405180830381600087803b15801561121a57600080fd5b505af192505050801561122b575060015b6112c1573d808015611259576040519150601f19603f3d011682016040523d82523d6000602084013e61125e565b606091505b5061127784600001518560200151848860a001516132fc565b7fdbbb883f24557adf486292429863dcfd4ac5d4db168ae94921da8e3d9a95d4168560a001516001836040516112af939291906149e9565b60405180910390a1505050505061130a565b7f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb83600001518460400151836040516112fc93929190614767565b60405180910390a15b505050505b50506005805460ff60a01b1916600160a01b179055505050505050505050565b60008060008060208601519150604086015192508560408151811061134b57fe5b602001015160f81c60f81b60f81c601b019050600185828486604051600081526020016040526040516113819493929190614a11565b6020604051602081039080840390855afa1580156113a3573d6000803e3d6000fd5b5050506020604051035193505050505b92915050565b600554600160a01b900460ff166113d06001612f40565b906113ee5760405162461bcd60e51b815260040161069b9190614a2f565b506005805460ff60a01b191690556b033b2e3c9fd0803ce80000004711156114166002612f40565b906114345760405162461bcd60e51b815260040161069b9190614a2f565b507fd30df8040a1092415b49422a02dbd8cdd5915a596abcba02cd0f65dd86ab3851600083833461146560066133f2565b604051611476959493929190614870565b60405180910390a161148860066133f6565b50506005805460ff60a01b1916600160a01b179055565b60f181565b600081565b600554600160a01b900460ff166114c06001612f40565b906114de5760405162461bcd60e51b815260040161069b9190614a2f565b506005805460ff60a01b191690558360006114f882612a4f565b90506000826001600160a01b03166370a08231306040518263ffffffff1660e01b81526004016115289190614753565b60206040518083038186803b15801561154057600080fd5b505afa158015611554573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115789190614567565b90508086600960ff851611156115b35760098460ff1603600a0a818161159a57fe5b04905060098460ff1603600a0a83816115af57fe5b0492505b670de0b6b3a764000081111580156115d35750670de0b6b3a76400008311155b80156115f05750670de0b6b3a76400006115ed8285613139565b11155b6115fa6003612f40565b906116185760405162461bcd60e51b815260040161069b9190614a2f565b506040516323b872dd60e01b81526001600160a01b038616906323b872dd9061164990339030908d90600401614767565b600060405180830381600087803b15801561166357600080fd5b505af1158015611677573d6000803e3d6000fd5b50505050611683613105565b61168d6004612f40565b906116ab5760405162461bcd60e51b815260040161069b9190614a2f565b50876116c0836116ba8c611971565b906133ff565b146116cb600a612f40565b906116e95760405162461bcd60e51b815260040161069b9190614a2f565b507fd30df8040a1092415b49422a02dbd8cdd5915a596abcba02cd0f65dd86ab38518988888461171960066133f2565b60405161172a959493929190614870565b60405180910390a161173c60066133f6565b50506005805460ff60a01b1916600160a01b17905550505050505050565b60046020526000908152604090205481565b600260209081526000928352604080842090915290825290205481565b6117916139fd565b611799613a32565b60606101288410156117ab6013612f40565b906117c95760405162461bcd60e51b815260040161069b9190614a2f565b506117d26139fd565b858560008181106117df57fe5b919091013560f81c825250858560018181106117f757fe5b919091013560f81c60208301525060008686600281811061181457fe5b845192013560f81c92505060ff16609e1480156118385750816020015160ff166001145b8015611847575060ff81166001145b6118516013612f40565b9061186f5760405162461bcd60e51b815260040161069b9190614a2f565b5050611879613a32565b61188760c36003888a614d85565b8101906118949190613cfb565b6001600160a01b03908116604088810191909152918116875260a0880192909252608087019290925291821660608601529116908301526118da61012860c3888a614d85565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505050506020820152818161192388610128818c614d85565b81818080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250969e959d50919b50939950505050505050505050565b60f381565b60006001600160a01b038216611988575047611a40565b61199182613455565b6119ad5760405162461bcd60e51b815260040161069b90614b53565b6040516370a0823160e01b81526001600160a01b038316906370a08231906119d9903090600401614753565b60206040518083038186803b1580156119f157600080fd5b505afa925050508015611a21575060408051601f3d908101601f19168201909252611a1e91810190614567565b60015b611a3d5760405162461bcd60e51b815260040161069b90614bf1565b90505b919050565b600554600160a01b900460ff16611a5c6001612f40565b90611a7a5760405162461bcd60e51b815260040161069b9190614a2f565b506005805460ff60a01b19169055895160821115611a986006612f40565b90611ab65760405162461bcd60e51b815260040161069b9190614a2f565b50611abf6139fd565b611ac88b611e74565b805190915060ff1660f3148015611ae65750806020015160ff166001145b611af06006612f40565b90611b0e5760405162461bcd60e51b815260040161069b9190614a2f565b50611b1c8160a00151611dbd565b15611b276005612f40565b90611b455760405162461bcd60e51b815260040161069b9190614a2f565b5060a081015160009081526020819052604090819020805460ff191660011790558101516001600160a01b0316611bd0576040808201516001600160a01b03166000908152600460205220546080820151611b9f91613139565b471015611bac6007612f40565b90611bca5760405162461bcd60e51b815260040161069b9190614a2f565b50611ce2565b6000611bdf8260400151612a4f565b905060098160ff161115611c0c576080820151611c069060081960ff841601600a0a61319f565b60808301525b6040808301516001600160a01b03166000908152600460205220546080830151611c3591613139565b82604001516001600160a01b03166370a08231306040518263ffffffff1660e01b8152600401611c659190614753565b60206040518083038186803b158015611c7d57600080fd5b505afa158015611c91573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611cb59190614567565b1015611cc16007612f40565b90611cdf5760405162461bcd60e51b815260040161069b9190614a2f565b50505b611cf48b8b8b8b8b8b8b8b8b8b613201565b608081015160608201516001600160a01b0390811660009081526002602090815260408083208187015190941683529290522054611d3191613139565b60608201516001600160a01b03908116600090815260026020908152604080832081870180518616855290835281842095909555608086015194519093168252600490522054611d8091613139565b6040918201516001600160a01b031660009081526004602052919091205550506005805460ff60a01b1916600160a01b1790555050505050505050565b60008181526020819052604081205460ff1615611ddc57506001611a40565b6005546001600160a01b0316611df457506000611a40565b600554604051633a4e2fc360e11b81526001600160a01b039091169063749c5f8690611e249085906004016149e0565b60206040518083038186803b158015611e3c57600080fd5b505afa158015611e50573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a3d91906140ac565b611e7c6139fd565b611e846139fd565b82600081518110611e9157fe5b016020015160f81c8152825183906001908110611eaa57fe5b0160209081015160f81c9082015260228301516042840151606285015160828601516001600160a01b039384166040860152929091166060840152608083015260a08201529050919050565b600554600160a01b900460ff16611f0d6001612f40565b90611f2b5760405162461bcd60e51b815260040161069b9190614a2f565b506005805460ff60a01b191690556000611fa7611f4b828d88888f61345b565b8a8a8a8a604051602001611f63959493929190614ca3565b60408051601f198184030181526020601f8701819004810284018101909252858352919086908690819084018382808284376000920191909152506134e192505050565b9050611fb3818c613585565b6001600160a01b038082166000908152600260209081526040808320938f16835292905220548a1115611fe66008612f40565b906120045760405162461bcd60e51b815260040161069b9190614a2f565b506001600160a01b038b16600090815260046020526040902054612028908b6133ff565b6001600160a01b03808d1660008181526004602090815260408083209590955592851681526002835283812091815291522054612065908b6133ff565b60026000836001600160a01b03166001600160a01b0316815260200190815260200160002060008d6001600160a01b03166001600160a01b0316815260200190815260200160002081905550600034905060006001600160a01b03168c6001600160a01b031614156120e2576120db818c613139565b905061221d565b6040516370a0823160e01b81528b906001600160a01b038e16906370a0823190612110903090600401614753565b60206040518083038186803b15801561212857600080fd5b505afa15801561213c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121609190614567565b101561216c6007612f40565b9061218a5760405162461bcd60e51b815260040161069b9190614a2f565b5060405163a9059cbb60e01b81526001600160a01b038d169063a9059cbb906121b9908c908f9060040161478b565b600060405180830381600087803b1580156121d357600080fd5b505af11580156121e7573d6000803e3d6000fd5b505050506121f3613105565b6121fd6004612f40565b9061221b5760405162461bcd60e51b815260040161069b9190614a2f565b505b600061226f8b838b8b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508d6136c7565b90506122cc8160026000866001600160a01b03166001600160a01b0316815260200190815260200160002060008e6001600160a01b03166001600160a01b031681526020019081526020016000205461313990919063ffffffff16565b60026000856001600160a01b03166001600160a01b0316815260200190815260200160002060008d6001600160a01b03166001600160a01b031681526020019081526020016000208190555061235081600460008e6001600160a01b03166001600160a01b031681526020019081526020016000205461313990919063ffffffff16565b6001600160a01b03909b1660009081526004602052604090209a909a5550506005805460ff60a01b1916600160a01b17905550505050505050505050565b600360209081526000928352604080842090915290825290205460ff1681565b600554600160a01b900460ff1681565b600554600160a01b900460ff166123d56001612f40565b906123f35760405162461bcd60e51b815260040161069b9190614a2f565b506005805460ff60a01b19169055604080516020601f84018190048102820181019092528281526124419185919085908590819084018382808284376000920191909152506137f092505050565b86600061244d82612a4f565b90506000826001600160a01b03166370a08231306040518263ffffffff1660e01b815260040161247d9190614753565b60206040518083038186803b15801561249557600080fd5b505afa1580156124a9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124cd9190614567565b90508089600960ff851611156125085760098460ff1603600a0a81816124ef57fe5b04905060098460ff1603600a0a838161250457fe5b0492505b670de0b6b3a764000081111580156125285750670de0b6b3a76400008311155b80156125455750670de0b6b3a76400006125428285613139565b11155b61254f6003612f40565b9061256d5760405162461bcd60e51b815260040161069b9190614a2f565b50846001600160a01b03166323b872dd33308e6040518463ffffffff1660e01b815260040161259e93929190614767565b600060405180830381600087803b1580156125b857600080fd5b505af11580156125cc573d6000803e3d6000fd5b505050506125d8613105565b6125e26004612f40565b906126005760405162461bcd60e51b815260040161069b9190614a2f565b508a61260f836116ba8f611971565b1461261a600a612f40565b906126385760405162461bcd60e51b815260040161069b9190614a2f565b507f2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e8c8b8b8460405161266e949392919061483a565b60405180910390a150506005805460ff60a01b1916600160a01b17905550505050505050505050565b609e81565b60003033146126ab6014612f40565b906126c95760405162461bcd60e51b815260040161069b9190614a2f565b5060006126d583611971565b905060606001600160a01b0388166126f9576126f287868861385d565b9050612797565b60405163a9059cbb60e01b81526001600160a01b0389169063a9059cbb90612727908a908a9060040161478b565b600060405180830381600087803b15801561274157600080fd5b505af1158015612755573d6000803e3d6000fd5b50505050612761613105565b61276b6004612f40565b906127895760405162461bcd60e51b815260040161069b9190614a2f565b50612794878661388d565b90505b80516040146127a66009612f40565b906127c45760405162461bcd60e51b815260040161069b9190614a2f565b50600080828060200190518101906127dc9190613d6a565b91509150856001600160a01b0316826001600160a01b031614801561280c57508061280a856116ba89611971565b145b6128166009612f40565b906128345760405162461bcd60e51b815260040161069b9190614a2f565b509998505050505050505050565b600554600160a81b900460ff161561285a600f612f40565b906128785760405162461bcd60e51b815260040161069b9190614a2f565b5060058054600160a01b600160a81b6001600160a01b03199092166001600160a01b03949094169390931760ff60a81b19161760ff60a01b1916919091179055565b600554600160a01b900460ff166128d16001612f40565b906128ef5760405162461bcd60e51b815260040161069b9190614a2f565b506005805460ff60a01b191690556b033b2e3c9fd0803ce80000004711156129176002612f40565b906129355760405162461bcd60e51b815260040161069b9190614a2f565b506129768383838080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506137f092505050565b7f2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e60008686346040516129ac949392919061483a565b60405180910390a150506005805460ff60a01b1916600160a01b179055505050565b6007546001600160a01b031615806129f057506007546001600160a01b031633145b8015612a0457506001600160a01b03811615155b612a0e6011612f40565b90612a2c5760405162461bcd60e51b815260040161069b9190614a2f565b50600780546001600160a01b0319166001600160a01b0392909216919091179055565b6000612a5a82613455565b612a765760405162461bcd60e51b815260040161069b90614bba565b6000829050806001600160a01b031663313ce5676040518163ffffffff1660e01b815260040160206040518083038186803b158015612ab457600080fd5b505afa925050508015612ae4575060408051601f3d908101601f19168201909252612ae191810190614567565b60015b612b005760405162461bcd60e51b815260040161069b90614a42565b9150611a409050565b600181565b60006020819052908152604090205460ff1681565b6007546001600160a01b031681565b60008181526001602052604081205460ff1615612b5157506001611a40565b6005546001600160a01b0316612b6957506000611a40565b60055460405163392f5c1d60e21b81526001600160a01b039091169063e4bd707490611e249085906004016149e0565b6005546000906001600160a01b031615801590612bdc57506001600160a01b0380831660009081526003602090815260408083209387168352929052205460ff16155b15612c9857600554604051637badcc6760e11b8152612c91916001600160a01b03169063f75b98ce90612c1590879087906004016147a4565b60206040518083038186803b158015612c2d57600080fd5b505afa158015612c41573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c659190614567565b6001600160a01b0380851660009081526002602090815260408083209389168352929052205490613139565b90506113b3565b506001600160a01b0380821660009081526002602090815260408083209386168352929052205492915050565b6005546001600160a01b031681565b600554600160a01b900460ff16612ceb6001612f40565b90612d095760405162461bcd60e51b815260040161069b9190614a2f565b506005805460ff60a01b19169055604080516020601f8401819004810282018101909252828152612d579185919085908590819084018382808284376000920191909152506137f092505050565b6000612dc2612d6a60018c89898e61345b565b8d8d604051602001612d7e93929190614cec565b60408051601f198184030181526020601f8c018190048102840181019092528a835291908b908b90819084018382808284376000920191909152506134e192505050565b9050612dce818b613585565b6001600160a01b038082166000908152600260209081526040808320938e1683529290522054891115612e016008612f40565b90612e1f5760405162461bcd60e51b815260040161069b9190614a2f565b506001600160a01b038082166000908152600260209081526040808320938e1683529290522054612e50908a6133ff565b6001600160a01b038083166000908152600260209081526040808320938f168352928152828220939093556004909252902054612e8d908a6133ff565b6001600160a01b038b16600081815260046020526040902091909155899015612ee1576000612ebb8c612a4f565b905060098160ff161115612edf5760098160ff1603600a0a8b81612edb57fe5b0491505b505b7f2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e8b8e8e84604051612f16949392919061483a565b60405180910390a150506005805460ff60a01b1916600160a01b1790555050505050505050505050565b60606000826014811115612f5057fe5b60408051600a808252818301909252919250906060908260208201818036833701905050905060005b60ff841615612fc7578151600a60ff959095168581049560018401939106916030830160f81b9185918110612faa57fe5b60200101906001600160f81b031916908160001a90535050612f79565b6060816001016001600160401b0381118015612fe257600080fd5b506040519080825280601f01601f19166020018201604052801561300d576020820181803683370190505b50905060005b82811161305e57838184038151811061302857fe5b602001015160f81c60f81b82828151811061303f57fe5b60200101906001600160f81b031916908160001a905350600101613013565b509695505050505050565b804710156130895760405162461bcd60e51b815260040161069b90614ad6565b6000826001600160a01b0316826040516130a290614750565b60006040518083038185875af1925050503d80600081146130df576040519150601f19603f3d011682016040523d82523d6000602084013e6130e4565b606091505b50509050806107545760405162461bcd60e51b815260040161069b90614a79565b6000803d801561311c576020811461312557613131565b60019150613131565b60206000803e60005191505b501515905090565b600082820183811080159061314e5750828110155b6040518060400160405280601281526020017129b0b332a6b0ba341032bc31b2b83a34b7b760711b815250906131975760405162461bcd60e51b815260040161069b9190614a2f565b509392505050565b600082820283158061314e5750828482816131b657fe5b04146040518060400160405280601281526020017129b0b332a6b0ba341032bc31b2b83a34b7b760711b815250906131975760405162461bcd60e51b815260040161069b9190614a2f565b60008a8a60405160200161321692919061472e565b6040516020818303038152906040528051906020012090506132366138cf565b6001600160a01b031663f65d21166001838d8d8d8d8d8d8d8d8d6040518c63ffffffff1660e01b81526004016132769b9a99989796959493929190614933565b60206040518083038186803b15801561328e57600080fd5b505afa1580156132a2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132c691906140ac565b6132d06006612f40565b906132ee5760405162461bcd60e51b815260040161069b9190614a2f565b505050505050505050505050565b816001600160a01b038516613349576b033b2e3c9fd0803ce80000004711156133256002612f40565b906133435760405162461bcd60e51b815260040161069b9190614a2f565b506133af565b600061335486612a4f565b905060098160ff1611156133785760098160ff1603600a0a828161337457fe5b0491505b670de0b6b3a764000082111561338e6003612f40565b906133ac5760405162461bcd60e51b815260040161069b9190614a2f565b50505b7eb45d95b5117447e2fafe7f34def913ff3ba220e4b8688acf37ae2328af7a3d858583856040516133e39493929190614803565b60405180910390a15050505050565b5490565b80546001019055565b6000828211156040518060400160405280601281526020017129b0b332a6b0ba341032bc31b2b83a34b7b760711b8152509061344e5760405162461bcd60e51b815260040161069b9190614a2f565b5050900390565b3b151590565b613463613a65565b61346b613a65565b604051806080016040528088600781111561348257fe5b8152602001876001600160a01b0316815260200186868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060200184905291505095945050505050565b815160208301206000906134f481612b32565b156134ff6005612f40565b9061351d5760405162461bcd60e51b815260040161069b9190614a2f565b50600061352a848361132a565b90506001600160a01b03811615156135426010612f40565b906135605760405162461bcd60e51b815260040161069b9190614a2f565b506000918252600160208190526040909220805460ff19169092179091559392505050565b6005546001600160a01b0316158015906135c557506001600160a01b0380831660009081526003602090815260408083209385168352929052205460ff16155b156136c357600554604051637badcc6760e11b815261367a916001600160a01b03169063f75b98ce906135fe90859087906004016147a4565b60206040518083038186803b15801561361657600080fd5b505afa15801561362a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061364e9190614567565b6001600160a01b0380851660009081526002602090815260408083209387168352929052205490613139565b6001600160a01b038084166000818152600260209081526040808320948716808452948252808320959095559181526003825283812092815291905220805460ff191660011790555b5050565b6000806136d386611971565b90506001600160a01b0386166136f0576136ed81346133ff565b90505b844710156136fe6007612f40565b9061371c5760405162461bcd60e51b815260040161069b9190614a2f565b5060006060846001600160a01b0316878760405161373a9190614712565b60006040518083038185875af1925050503d8060008114613777576040519150601f19603f3d011682016040523d82523d6000602084013e61377c565b606091505b50915091508161378c6004612f40565b906137aa5760405162461bcd60e51b815260040161069b9190614a2f565b50600080828060200190518101906137c29190613d6a565b91509150896001600160a01b0316826001600160a01b031614801561280c57508061280a866116ba8d611971565b600061383f826040518060400160405280336001600160a01b03168152602001868152506040516020016138249190614d12565b6040516020818303038152906040528051906020012061132a565b6007549091506001600160a01b038083169116146107346010612f40565b6060613883848484604051806060016040528060298152602001614e00602991396138f4565b90505b9392505050565b606061388683836040518060400160405280601e81526020017f416464726573733a206c6f772d6c6576656c2063616c6c206661696c656400008152506139b5565b7f62135fc083646fdb4e1a9d700e351b886a4a5a39da980650269edd1ade91ffd25490565b6060824710156139165760405162461bcd60e51b815260040161069b90614b0d565b61391f85613455565b61393b5760405162461bcd60e51b815260040161069b90614b83565b60006060866001600160a01b031685876040516139589190614712565b60006040518083038185875af1925050503d8060008114613995576040519150601f19603f3d011682016040523d82523d6000602084013e61399a565b606091505b50915091506139aa8282866139c4565b979650505050505050565b606061388384846000856138f4565b606083156139d3575081613886565b8251156139e35782518084602001fd5b8160405162461bcd60e51b815260040161069b9190614a2f565b6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a081019190915290565b604051806060016040528060006001600160a01b031681526020016060815260200160006001600160a01b031681525090565b6040805160808101909152806000815260200160006001600160a01b0316815260200160608152602001600081525090565b80356113b381614dd9565b60008083601f840112613ab3578182fd5b5081356001600160401b03811115613ac9578182fd5b6020830191508360208083028501011115613ae357600080fd5b9250929050565b600082601f830112613afa578081fd5b8135613b0d613b0882614d66565b614d40565b818152915060208083019084810181840286018201871015613b2e57600080fd5b60005b84811015613b56578135613b4481614df1565b84529282019290820190600101613b31565b505050505092915050565b600082601f830112613b71578081fd5b8135613b7f613b0882614d66565b818152915060208083019084810181840286018201871015613ba057600080fd5b60005b84811015613b5657813584529282019290820190600101613ba3565b600082601f830112613bcf578081fd5b8135613bdd613b0882614d66565b818152915060208083019084810181840286018201871015613bfe57600080fd5b6000805b85811015613c2c57823560ff81168114613c1a578283fd5b85529383019391830191600101613c02565b50505050505092915050565b60008083601f840112613c49578182fd5b5081356001600160401b03811115613c5f578182fd5b602083019150836020828501011115613ae357600080fd5b600082601f830112613c87578081fd5b81356001600160401b03811115613c9c578182fd5b613caf601f8201601f1916602001614d40565b9150808252836020828501011115613cc657600080fd5b8060208401602084013760009082016020015292915050565b600060208284031215613cf0578081fd5b813561388681614dd9565b60008060008060008060c08789031215613d13578182fd5b8635613d1e81614dd9565b95506020870135613d2e81614dd9565b945060408701359350606087013592506080870135613d4c81614dd9565b915060a0870135613d5c81614dd9565b809150509295509295509295565b60008060408385031215613d7c578182fd5b8251613d8781614dd9565b6020939093015192949293505050565b60008060408385031215613da9578182fd5b8235613db481614dd9565b91506020830135613dc481614dd9565b809150509250929050565b600080600060608486031215613de3578283fd5b8335613dee81614dd9565b92506020840135613dfe81614dd9565b929592945050506040919091013590565b600080600080600060a08688031215613e26578283fd5b8535613e3181614dd9565b94506020860135613e4181614dd9565b93506040860135925060608601356001600160401b03811115613e62578182fd5b613e6e88828901613c77565b9250506080860135613e7f81614dd9565b809150509295509295909350565b60008060008060008060008060008060e08b8d031215613eab578788fd5b8a35613eb681614dd9565b995060208b0135985060408b0135613ecd81614dd9565b9750613edc8c60608d01613a97565b965060808b01356001600160401b0380821115613ef7578586fd5b613f038e838f01613c38565b909850965060a08d0135915080821115613f1b578586fd5b613f278e838f01613c38565b909650945060c08d0135915080821115613f3f578384fd5b50613f4c8d828e01613c38565b915080935050809150509295989b9194979a5092959850565b60008060008060608587031215613f7a578182fd5b8435613f8581614dd9565b93506020850135925060408501356001600160401b03811115613fa6578283fd5b613fb287828801613c38565b95989497509550505050565b600080600080600080600060a0888a031215613fd8578081fd5b8735613fe381614dd9565b96506020880135955060408801356001600160401b0380821115614005578283fd5b6140118b838c01613c38565b909750955060608a0135945060808a0135915080821115614030578283fd5b5061403d8a828b01613c38565b989b979a50959850939692959293505050565b60008060008060408587031215614065578182fd5b84356001600160401b038082111561407b578384fd5b61408788838901613aa2565b9096509450602087013591508082111561409f578384fd5b50613fb287828801613aa2565b6000602082840312156140bd578081fd5b815161388681614df1565b6000602082840312156140d9578081fd5b5035919050565b600080602083850312156140f2578182fd5b82356001600160401b03811115614107578283fd5b61411385828601613c38565b90969095509350505050565b60008060008060008060008060008060006101408c8e031215614140578485fd5b6001600160401b03808d351115614155578586fd5b6141628e8e358f01613c38565b909c509a5060208d0135995060408d013581101561417e578586fd5b61418e8e60408f01358f01613b61565b98508060608e013511156141a0578586fd5b6141b08e60608f01358f01613aea565b975060808d0135965060a08d013595508060c08e013511156141d0578182fd5b6141e08e60c08f01358f01613b61565b94508060e08e013511156141f2578182fd5b6142028e60e08f01358f01613bbf565b9350806101008e01351115614215578182fd5b6142268e6101008f01358f01613b61565b9250806101208e01351115614239578182fd5b5061424b8d6101208e01358e01613b61565b90509295989b509295989b9093969950565b60006020828403121561426e578081fd5b81356001600160401b03811115614283578182fd5b61428f84828501613c77565b949350505050565b600080604083850312156142a9578182fd5b82356001600160401b038111156142be578283fd5b6142ca85828601613c77565b95602094909401359450505050565b6000806000806000806000806000806101408b8d0312156142f8578384fd5b8a356001600160401b038082111561430e578586fd5b61431a8e838f01613c77565b9b5060208d01359a5060408d0135915080821115614336578586fd5b6143428e838f01613b61565b995060608d0135915080821115614357578586fd5b6143638e838f01613aea565b985060808d0135975060a08d0135965060c08d0135915080821115614386578586fd5b6143928e838f01613b61565b955060e08d01359150808211156143a7578485fd5b6143b38e838f01613bbf565b94506101008d01359150808211156143c9578384fd5b6143d58e838f01613b61565b93506101208d01359150808211156143eb578283fd5b506143f88d828e01613b61565b9150509295989b9194979a5092959850565b600080600080600080600080600080600060e08c8e03121561442a578485fd5b6001600160401b03808d35111561443f578586fd5b61444c8e8e358f01613c38565b909c509a5061445e8e60208f01613a97565b995060408d013598508060608e01351115614477578586fd5b6144878e60608f01358f01613c38565b909850965060808d013581101561449c578586fd5b6144ac8e60808f01358f01613c38565b909650945060a08d0135935060c08d01358110156144c8578182fd5b506144d98d60c08e01358e01613c38565b81935080925050509295989b509295989b9093969950565b600080600080600060608688031215614508578283fd5b85356001600160401b038082111561451e578485fd5b61452a89838a01613c38565b9097509550602088013594506040880135915080821115614549578283fd5b5061455688828901613c38565b969995985093965092949392505050565b600060208284031215614578578081fd5b5051919050565b6000815180845260208085019450808401835b838110156145b0578151151587529582019590820190600101614592565b509495945050505050565b6000815180845260208085019450808401835b838110156145b0578151875295820195908201906001016145ce565b6000815180845260208085019450808401835b838110156145b057815160ff16875295820195908201906001016145fd565b60008284528282602086013780602084860101526020601f19601f85011685010190509392505050565b6000815180845261465e816020860160208601614dad565b601f01601f19169290920160200192915050565b60ff815116825260ff6020820151166020830152604081015160018060a01b03808216604085015280606084015116606085015250506080810151608083015260a081015160a08301525050565b60008151600881106146ce57fe5b83526020828101516001600160a01b0316908401526040808301516080918501829052906146fe90850182614646565b606093840151949093019390935250919050565b60008251614724818460208701614dad565b9190910192915050565b60008351614740818460208801614dad565b9190910191825250602001919050565b90565b6001600160a01b0391909116815260200190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b6001600160a01b03929092168252602082015260400190565b6001600160a01b0392831681529116602082015260400190565b600060018060a01b038088168352808716602084015285604084015260a060608401526147ee60a0840186614646565b91508084166080840152509695505050505050565b6001600160a01b038516815260806020820181905260009061482790830186614646565b6040830194909452506060015292915050565b6001600160a01b038516815260606020820181905260009061485f908301858761461c565b905082604083015295945050505050565b6001600160a01b0386168152608060208201819052600090614895908301868861461c565b604083019490945250606001529392505050565b6040808252810184905260008560608301825b878110156148ec57602083356148d181614dd9565b6001600160a01b0316835292830192909101906001016148bc565b5083810360208501528481526001600160fb1b0385111561490b578283fd5b602085029150818660208301370160200190815295945050505050565b901515815260200190565b60006101608d151583528c60208401528b604084015280606084015261495b8184018c6145bb565b9050828103608084015261496f818b61457f565b90508860a08401528760c084015282810360e084015261498f81886145bb565b90508281036101008401526149a481876145ea565b90508281036101208401526149b981866145bb565b90508281036101408401526149ce81856145bb565b9e9d5050505050505050505050505050565b90815260200190565b600084825283602083015260606040830152614a086060830184614646565b95945050505050565b93845260ff9290921660208401526040830152606082015260800190565b6000602082526138866020830184614646565b60208082526018908201527f67657420455243323020646563696d616c206661696c65640000000000000000604082015260600190565b6020808252603a908201527f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260408201527f6563697069656e74206d61792068617665207265766572746564000000000000606082015260800190565b6020808252601d908201527f416464726573733a20696e73756666696369656e742062616c616e6365000000604082015260600190565b60208082526026908201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6040820152651c8818d85b1b60d21b606082015260800190565b60208082526016908201527518985b185b98d953d9881b9bdb8b58dbdb9d1c9858dd60521b604082015260600190565b6020808252601d908201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604082015260600190565b60208082526018908201527f676574446563696d616c73206e6f6e2d636f6e74726163740000000000000000604082015260600190565b60208082526018908201527f6765742045524332302062616c616e6365206661696c65640000000000000000604082015260600190565b60c081016113b38284614672565b6000610100614c458387614672565b8060c084015260018060a01b038086511682850152602086015191506060610120850152614c77610160850183614646565b9150806040870151166101408501525082810360e0840152614c998185614646565b9695505050505050565b600060808252614cb660808301886146c0565b6001600160a01b038781166020850152861660408401528281036060840152614ce081858761461c565b98975050505050505050565b600060408252614cff60408301866146c0565b8281036020840152614c9981858761461c565b81516001600160a01b031681526020918201519181019190915260400190565b60ff91909116815260200190565b6040518181016001600160401b0381118282101715614d5e57600080fd5b604052919050565b60006001600160401b03821115614d7b578081fd5b5060209081020190565b60008085851115614d94578182fd5b83861115614da0578182fd5b5050820193919092039150565b60005b83811015614dc8578181015183820152602001614db0565b838111156107525750506000910152565b6001600160a01b0381168114614dee57600080fd5b50565b8015158114614dee57600080fdfe416464726573733a206c6f772d6c6576656c2063616c6c20776974682076616c7565206661696c6564a2646970667358221220222191c6435a76a38262cc3285674d98e209845da2bd564dc78a473938536bf964736f6c634300060c0033",
}

// VaultABI is the input ABI used to generate the binding from.
// Deprecated: Use VaultMetaData.ABI instead.
var VaultABI = VaultMetaData.ABI

// Deprecated: Use VaultMetaData.Sigs instead.
// VaultFuncSigs maps the 4-byte function signature to its string representation.
var VaultFuncSigs = VaultMetaData.Sigs

// VaultBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VaultMetaData.Bin instead.
var VaultBin = VaultMetaData.Bin

// DeployVault deploys a new Ethereum contract, binding an instance of Vault to it.
func DeployVault(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Vault, error) {
	parsed, err := VaultMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VaultBin), backend)
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
func (_Vault *VaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_Vault *VaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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

// BURNCALLREQUESTMETADATATYPE is a free data retrieval call binding the contract method 0xbd835c42.
//
// Solidity: function BURN_CALL_REQUEST_METADATA_TYPE() view returns(uint8)
func (_Vault *VaultCaller) BURNCALLREQUESTMETADATATYPE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "BURN_CALL_REQUEST_METADATA_TYPE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BURNCALLREQUESTMETADATATYPE is a free data retrieval call binding the contract method 0xbd835c42.
//
// Solidity: function BURN_CALL_REQUEST_METADATA_TYPE() view returns(uint8)
func (_Vault *VaultSession) BURNCALLREQUESTMETADATATYPE() (uint8, error) {
	return _Vault.Contract.BURNCALLREQUESTMETADATATYPE(&_Vault.CallOpts)
}

// BURNCALLREQUESTMETADATATYPE is a free data retrieval call binding the contract method 0xbd835c42.
//
// Solidity: function BURN_CALL_REQUEST_METADATA_TYPE() view returns(uint8)
func (_Vault *VaultCallerSession) BURNCALLREQUESTMETADATATYPE() (uint8, error) {
	return _Vault.Contract.BURNCALLREQUESTMETADATATYPE(&_Vault.CallOpts)
}

// BURNREQUESTMETADATATYPE is a free data retrieval call binding the contract method 0x568c04fd.
//
// Solidity: function BURN_REQUEST_METADATA_TYPE() view returns(uint8)
func (_Vault *VaultCaller) BURNREQUESTMETADATATYPE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "BURN_REQUEST_METADATA_TYPE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BURNREQUESTMETADATATYPE is a free data retrieval call binding the contract method 0x568c04fd.
//
// Solidity: function BURN_REQUEST_METADATA_TYPE() view returns(uint8)
func (_Vault *VaultSession) BURNREQUESTMETADATATYPE() (uint8, error) {
	return _Vault.Contract.BURNREQUESTMETADATATYPE(&_Vault.CallOpts)
}

// BURNREQUESTMETADATATYPE is a free data retrieval call binding the contract method 0x568c04fd.
//
// Solidity: function BURN_REQUEST_METADATA_TYPE() view returns(uint8)
func (_Vault *VaultCallerSession) BURNREQUESTMETADATATYPE() (uint8, error) {
	return _Vault.Contract.BURNREQUESTMETADATATYPE(&_Vault.CallOpts)
}

// BURNTOCONTRACTREQUESTMETADATATYPE is a free data retrieval call binding the contract method 0x6f2cbc48.
//
// Solidity: function BURN_TO_CONTRACT_REQUEST_METADATA_TYPE() view returns(uint8)
func (_Vault *VaultCaller) BURNTOCONTRACTREQUESTMETADATATYPE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "BURN_TO_CONTRACT_REQUEST_METADATA_TYPE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BURNTOCONTRACTREQUESTMETADATATYPE is a free data retrieval call binding the contract method 0x6f2cbc48.
//
// Solidity: function BURN_TO_CONTRACT_REQUEST_METADATA_TYPE() view returns(uint8)
func (_Vault *VaultSession) BURNTOCONTRACTREQUESTMETADATATYPE() (uint8, error) {
	return _Vault.Contract.BURNTOCONTRACTREQUESTMETADATATYPE(&_Vault.CallOpts)
}

// BURNTOCONTRACTREQUESTMETADATATYPE is a free data retrieval call binding the contract method 0x6f2cbc48.
//
// Solidity: function BURN_TO_CONTRACT_REQUEST_METADATA_TYPE() view returns(uint8)
func (_Vault *VaultCallerSession) BURNTOCONTRACTREQUESTMETADATATYPE() (uint8, error) {
	return _Vault.Contract.BURNTOCONTRACTREQUESTMETADATATYPE(&_Vault.CallOpts)
}

// CURRENTNETWORKID is a free data retrieval call binding the contract method 0xd7200eb1.
//
// Solidity: function CURRENT_NETWORK_ID() view returns(uint8)
func (_Vault *VaultCaller) CURRENTNETWORKID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "CURRENT_NETWORK_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CURRENTNETWORKID is a free data retrieval call binding the contract method 0xd7200eb1.
//
// Solidity: function CURRENT_NETWORK_ID() view returns(uint8)
func (_Vault *VaultSession) CURRENTNETWORKID() (uint8, error) {
	return _Vault.Contract.CURRENTNETWORKID(&_Vault.CallOpts)
}

// CURRENTNETWORKID is a free data retrieval call binding the contract method 0xd7200eb1.
//
// Solidity: function CURRENT_NETWORK_ID() view returns(uint8)
func (_Vault *VaultCallerSession) CURRENTNETWORKID() (uint8, error) {
	return _Vault.Contract.CURRENTNETWORKID(&_Vault.CallOpts)
}

// ETHTOKEN is a free data retrieval call binding the contract method 0x58bc8337.
//
// Solidity: function ETH_TOKEN() view returns(address)
func (_Vault *VaultCaller) ETHTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "ETH_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

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

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address token) view returns(uint256)
func (_Vault *VaultCaller) BalanceOf(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "balanceOf", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address token) view returns(uint256)
func (_Vault *VaultSession) BalanceOf(token common.Address) (*big.Int, error) {
	return _Vault.Contract.BalanceOf(&_Vault.CallOpts, token)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address token) view returns(uint256)
func (_Vault *VaultCallerSession) BalanceOf(token common.Address) (*big.Int, error) {
	return _Vault.Contract.BalanceOf(&_Vault.CallOpts, token)
}

// GetDecimals is a free data retrieval call binding the contract method 0xcf54aaa0.
//
// Solidity: function getDecimals(address token) view returns(uint8)
func (_Vault *VaultCaller) GetDecimals(opts *bind.CallOpts, token common.Address) (uint8, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getDecimals", token)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

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
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getDepositedBalance", token, owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_Vault *VaultCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_Vault *VaultSession) IsInitialized() (bool, error) {
	return _Vault.Contract.IsInitialized(&_Vault.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_Vault *VaultCallerSession) IsInitialized() (bool, error) {
	return _Vault.Contract.IsInitialized(&_Vault.CallOpts)
}

// IsSigDataUsed is a free data retrieval call binding the contract method 0xe4bd7074.
//
// Solidity: function isSigDataUsed(bytes32 hash) view returns(bool)
func (_Vault *VaultCaller) IsSigDataUsed(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "isSigDataUsed", hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

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
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "isWithdrawed", hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

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

// Migration is a free data retrieval call binding the contract method 0x995fac11.
//
// Solidity: function migration(address , address ) view returns(bool)
func (_Vault *VaultCaller) Migration(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "migration", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

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

// NotEntered is a free data retrieval call binding the contract method 0xa3f5d8cc.
//
// Solidity: function notEntered() view returns(bool)
func (_Vault *VaultCaller) NotEntered(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "notEntered")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

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
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "parseBurnInst", inst)

	if err != nil {
		return *new(VaultBurnInstData), err
	}

	out0 := *abi.ConvertType(out[0], new(VaultBurnInstData)).(*VaultBurnInstData)

	return out0, err

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

// ParseCalldataFromBurnInst is a free data retrieval call binding the contract method 0x66945b31.
//
// Solidity: function parseCalldataFromBurnInst(bytes inst) pure returns((uint8,uint8,address,address,uint256,bytes32), (address,bytes,address), bytes)
func (_Vault *VaultCaller) ParseCalldataFromBurnInst(opts *bind.CallOpts, inst []byte) (VaultBurnInstData, VaultRedepositOptions, []byte, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "parseCalldataFromBurnInst", inst)

	if err != nil {
		return *new(VaultBurnInstData), *new(VaultRedepositOptions), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(VaultBurnInstData)).(*VaultBurnInstData)
	out1 := *abi.ConvertType(out[1], new(VaultRedepositOptions)).(*VaultRedepositOptions)
	out2 := *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return out0, out1, out2, err

}

// ParseCalldataFromBurnInst is a free data retrieval call binding the contract method 0x66945b31.
//
// Solidity: function parseCalldataFromBurnInst(bytes inst) pure returns((uint8,uint8,address,address,uint256,bytes32), (address,bytes,address), bytes)
func (_Vault *VaultSession) ParseCalldataFromBurnInst(inst []byte) (VaultBurnInstData, VaultRedepositOptions, []byte, error) {
	return _Vault.Contract.ParseCalldataFromBurnInst(&_Vault.CallOpts, inst)
}

// ParseCalldataFromBurnInst is a free data retrieval call binding the contract method 0x66945b31.
//
// Solidity: function parseCalldataFromBurnInst(bytes inst) pure returns((uint8,uint8,address,address,uint256,bytes32), (address,bytes,address), bytes)
func (_Vault *VaultCallerSession) ParseCalldataFromBurnInst(inst []byte) (VaultBurnInstData, VaultRedepositOptions, []byte, error) {
	return _Vault.Contract.ParseCalldataFromBurnInst(&_Vault.CallOpts, inst)
}

// PrevVault is a free data retrieval call binding the contract method 0xfa84702e.
//
// Solidity: function prevVault() view returns(address)
func (_Vault *VaultCaller) PrevVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "prevVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

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

// Regulator is a free data retrieval call binding the contract method 0xdd8fee14.
//
// Solidity: function regulator() view returns(address)
func (_Vault *VaultCaller) Regulator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "regulator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Regulator is a free data retrieval call binding the contract method 0xdd8fee14.
//
// Solidity: function regulator() view returns(address)
func (_Vault *VaultSession) Regulator() (common.Address, error) {
	return _Vault.Contract.Regulator(&_Vault.CallOpts)
}

// Regulator is a free data retrieval call binding the contract method 0xdd8fee14.
//
// Solidity: function regulator() view returns(address)
func (_Vault *VaultCallerSession) Regulator() (common.Address, error) {
	return _Vault.Contract.Regulator(&_Vault.CallOpts)
}

// SigDataUsed is a free data retrieval call binding the contract method 0x1ea1940e.
//
// Solidity: function sigDataUsed(bytes32 ) view returns(bool)
func (_Vault *VaultCaller) SigDataUsed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "sigDataUsed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

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
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "sigToAddress", signData, hash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

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

// TotalDepositedToSCAmount is a free data retrieval call binding the contract method 0x6304541c.
//
// Solidity: function totalDepositedToSCAmount(address ) view returns(uint256)
func (_Vault *VaultCaller) TotalDepositedToSCAmount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "totalDepositedToSCAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "withdrawRequests", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "withdrawed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

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

// CallExternal is a paid mutator transaction binding the contract method 0xbda9b509.
//
// Solidity: function _callExternal(address token, address to, uint256 amount, bytes externalCalldata, address redepositToken) returns(uint256)
func (_Vault *VaultTransactor) CallExternal(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int, externalCalldata []byte, redepositToken common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "_callExternal", token, to, amount, externalCalldata, redepositToken)
}

// CallExternal is a paid mutator transaction binding the contract method 0xbda9b509.
//
// Solidity: function _callExternal(address token, address to, uint256 amount, bytes externalCalldata, address redepositToken) returns(uint256)
func (_Vault *VaultSession) CallExternal(token common.Address, to common.Address, amount *big.Int, externalCalldata []byte, redepositToken common.Address) (*types.Transaction, error) {
	return _Vault.Contract.CallExternal(&_Vault.TransactOpts, token, to, amount, externalCalldata, redepositToken)
}

// CallExternal is a paid mutator transaction binding the contract method 0xbda9b509.
//
// Solidity: function _callExternal(address token, address to, uint256 amount, bytes externalCalldata, address redepositToken) returns(uint256)
func (_Vault *VaultTransactorSession) CallExternal(token common.Address, to common.Address, amount *big.Int, externalCalldata []byte, redepositToken common.Address) (*types.Transaction, error) {
	return _Vault.Contract.CallExternal(&_Vault.TransactOpts, token, to, amount, externalCalldata, redepositToken)
}

// TransferExternal is a paid mutator transaction binding the contract method 0x145e2a6b.
//
// Solidity: function _transferExternal(address token, address to, uint256 amount) returns()
func (_Vault *VaultTransactor) TransferExternal(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "_transferExternal", token, to, amount)
}

// TransferExternal is a paid mutator transaction binding the contract method 0x145e2a6b.
//
// Solidity: function _transferExternal(address token, address to, uint256 amount) returns()
func (_Vault *VaultSession) TransferExternal(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.TransferExternal(&_Vault.TransactOpts, token, to, amount)
}

// TransferExternal is a paid mutator transaction binding the contract method 0x145e2a6b.
//
// Solidity: function _transferExternal(address token, address to, uint256 amount) returns()
func (_Vault *VaultTransactorSession) TransferExternal(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.TransferExternal(&_Vault.TransactOpts, token, to, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xc791d705.
//
// Solidity: function deposit(string incognitoAddress, bytes32 txId, bytes signData) payable returns()
func (_Vault *VaultTransactor) Deposit(opts *bind.TransactOpts, incognitoAddress string, txId [32]byte, signData []byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "deposit", incognitoAddress, txId, signData)
}

// Deposit is a paid mutator transaction binding the contract method 0xc791d705.
//
// Solidity: function deposit(string incognitoAddress, bytes32 txId, bytes signData) payable returns()
func (_Vault *VaultSession) Deposit(incognitoAddress string, txId [32]byte, signData []byte) (*types.Transaction, error) {
	return _Vault.Contract.Deposit(&_Vault.TransactOpts, incognitoAddress, txId, signData)
}

// Deposit is a paid mutator transaction binding the contract method 0xc791d705.
//
// Solidity: function deposit(string incognitoAddress, bytes32 txId, bytes signData) payable returns()
func (_Vault *VaultTransactorSession) Deposit(incognitoAddress string, txId [32]byte, signData []byte) (*types.Transaction, error) {
	return _Vault.Contract.Deposit(&_Vault.TransactOpts, incognitoAddress, txId, signData)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xa807b5bb.
//
// Solidity: function depositERC20(address token, uint256 amount, string incognitoAddress, bytes32 txId, bytes signData) returns()
func (_Vault *VaultTransactor) DepositERC20(opts *bind.TransactOpts, token common.Address, amount *big.Int, incognitoAddress string, txId [32]byte, signData []byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "depositERC20", token, amount, incognitoAddress, txId, signData)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xa807b5bb.
//
// Solidity: function depositERC20(address token, uint256 amount, string incognitoAddress, bytes32 txId, bytes signData) returns()
func (_Vault *VaultSession) DepositERC20(token common.Address, amount *big.Int, incognitoAddress string, txId [32]byte, signData []byte) (*types.Transaction, error) {
	return _Vault.Contract.DepositERC20(&_Vault.TransactOpts, token, amount, incognitoAddress, txId, signData)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xa807b5bb.
//
// Solidity: function depositERC20(address token, uint256 amount, string incognitoAddress, bytes32 txId, bytes signData) returns()
func (_Vault *VaultTransactorSession) DepositERC20(token common.Address, amount *big.Int, incognitoAddress string, txId [32]byte, signData []byte) (*types.Transaction, error) {
	return _Vault.Contract.DepositERC20(&_Vault.TransactOpts, token, amount, incognitoAddress, txId, signData)
}

// DepositERC20V2 is a paid mutator transaction binding the contract method 0x5c421b2f.
//
// Solidity: function depositERC20_V2(address token, uint256 amount, string incognitoAddress) returns()
func (_Vault *VaultTransactor) DepositERC20V2(opts *bind.TransactOpts, token common.Address, amount *big.Int, incognitoAddress string) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "depositERC20_V2", token, amount, incognitoAddress)
}

// DepositERC20V2 is a paid mutator transaction binding the contract method 0x5c421b2f.
//
// Solidity: function depositERC20_V2(address token, uint256 amount, string incognitoAddress) returns()
func (_Vault *VaultSession) DepositERC20V2(token common.Address, amount *big.Int, incognitoAddress string) (*types.Transaction, error) {
	return _Vault.Contract.DepositERC20V2(&_Vault.TransactOpts, token, amount, incognitoAddress)
}

// DepositERC20V2 is a paid mutator transaction binding the contract method 0x5c421b2f.
//
// Solidity: function depositERC20_V2(address token, uint256 amount, string incognitoAddress) returns()
func (_Vault *VaultTransactorSession) DepositERC20V2(token common.Address, amount *big.Int, incognitoAddress string) (*types.Transaction, error) {
	return _Vault.Contract.DepositERC20V2(&_Vault.TransactOpts, token, amount, incognitoAddress)
}

// DepositV2 is a paid mutator transaction binding the contract method 0x49c5ba23.
//
// Solidity: function deposit_V2(string incognitoAddress) payable returns()
func (_Vault *VaultTransactor) DepositV2(opts *bind.TransactOpts, incognitoAddress string) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "deposit_V2", incognitoAddress)
}

// DepositV2 is a paid mutator transaction binding the contract method 0x49c5ba23.
//
// Solidity: function deposit_V2(string incognitoAddress) payable returns()
func (_Vault *VaultSession) DepositV2(incognitoAddress string) (*types.Transaction, error) {
	return _Vault.Contract.DepositV2(&_Vault.TransactOpts, incognitoAddress)
}

// DepositV2 is a paid mutator transaction binding the contract method 0x49c5ba23.
//
// Solidity: function deposit_V2(string incognitoAddress) payable returns()
func (_Vault *VaultTransactorSession) DepositV2(incognitoAddress string) (*types.Transaction, error) {
	return _Vault.Contract.DepositV2(&_Vault.TransactOpts, incognitoAddress)
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

// ExecuteWithBurnProof is a paid mutator transaction binding the contract method 0x3ed1b376.
//
// Solidity: function executeWithBurnProof(bytes inst, uint256 heights, bytes32[] instPaths, bool[] instPathIsLefts, bytes32 instRoots, bytes32 blkData, uint256[] sigIdxs, uint8[] sigVs, bytes32[] sigRs, bytes32[] sigSs) returns()
func (_Vault *VaultTransactor) ExecuteWithBurnProof(opts *bind.TransactOpts, inst []byte, heights *big.Int, instPaths [][32]byte, instPathIsLefts []bool, instRoots [32]byte, blkData [32]byte, sigIdxs []*big.Int, sigVs []uint8, sigRs [][32]byte, sigSs [][32]byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "executeWithBurnProof", inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// ExecuteWithBurnProof is a paid mutator transaction binding the contract method 0x3ed1b376.
//
// Solidity: function executeWithBurnProof(bytes inst, uint256 heights, bytes32[] instPaths, bool[] instPathIsLefts, bytes32 instRoots, bytes32 blkData, uint256[] sigIdxs, uint8[] sigVs, bytes32[] sigRs, bytes32[] sigSs) returns()
func (_Vault *VaultSession) ExecuteWithBurnProof(inst []byte, heights *big.Int, instPaths [][32]byte, instPathIsLefts []bool, instRoots [32]byte, blkData [32]byte, sigIdxs []*big.Int, sigVs []uint8, sigRs [][32]byte, sigSs [][32]byte) (*types.Transaction, error) {
	return _Vault.Contract.ExecuteWithBurnProof(&_Vault.TransactOpts, inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// ExecuteWithBurnProof is a paid mutator transaction binding the contract method 0x3ed1b376.
//
// Solidity: function executeWithBurnProof(bytes inst, uint256 heights, bytes32[] instPaths, bool[] instPathIsLefts, bytes32 instRoots, bytes32 blkData, uint256[] sigIdxs, uint8[] sigVs, bytes32[] sigRs, bytes32[] sigSs) returns()
func (_Vault *VaultTransactorSession) ExecuteWithBurnProof(inst []byte, heights *big.Int, instPaths [][32]byte, instPathIsLefts []bool, instRoots [32]byte, blkData [32]byte, sigIdxs []*big.Int, sigVs []uint8, sigRs [][32]byte, sigSs [][32]byte) (*types.Transaction, error) {
	return _Vault.Contract.ExecuteWithBurnProof(&_Vault.TransactOpts, inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _prevVault) returns()
func (_Vault *VaultTransactor) Initialize(opts *bind.TransactOpts, _prevVault common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "initialize", _prevVault)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _prevVault) returns()
func (_Vault *VaultSession) Initialize(_prevVault common.Address) (*types.Transaction, error) {
	return _Vault.Contract.Initialize(&_Vault.TransactOpts, _prevVault)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _prevVault) returns()
func (_Vault *VaultTransactorSession) Initialize(_prevVault common.Address) (*types.Transaction, error) {
	return _Vault.Contract.Initialize(&_Vault.TransactOpts, _prevVault)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0xfee8efda.
//
// Solidity: function requestWithdraw(string incognitoAddress, address token, uint256 amount, bytes signData, bytes timestamp, bytes32 txId, bytes regulatorSig) returns()
func (_Vault *VaultTransactor) RequestWithdraw(opts *bind.TransactOpts, incognitoAddress string, token common.Address, amount *big.Int, signData []byte, timestamp []byte, txId [32]byte, regulatorSig []byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "requestWithdraw", incognitoAddress, token, amount, signData, timestamp, txId, regulatorSig)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0xfee8efda.
//
// Solidity: function requestWithdraw(string incognitoAddress, address token, uint256 amount, bytes signData, bytes timestamp, bytes32 txId, bytes regulatorSig) returns()
func (_Vault *VaultSession) RequestWithdraw(incognitoAddress string, token common.Address, amount *big.Int, signData []byte, timestamp []byte, txId [32]byte, regulatorSig []byte) (*types.Transaction, error) {
	return _Vault.Contract.RequestWithdraw(&_Vault.TransactOpts, incognitoAddress, token, amount, signData, timestamp, txId, regulatorSig)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0xfee8efda.
//
// Solidity: function requestWithdraw(string incognitoAddress, address token, uint256 amount, bytes signData, bytes timestamp, bytes32 txId, bytes regulatorSig) returns()
func (_Vault *VaultTransactorSession) RequestWithdraw(incognitoAddress string, token common.Address, amount *big.Int, signData []byte, timestamp []byte, txId [32]byte, regulatorSig []byte) (*types.Transaction, error) {
	return _Vault.Contract.RequestWithdraw(&_Vault.TransactOpts, incognitoAddress, token, amount, signData, timestamp, txId, regulatorSig)
}

// SetRegulator is a paid mutator transaction binding the contract method 0xcde0a4f8.
//
// Solidity: function setRegulator(address _regulator) returns()
func (_Vault *VaultTransactor) SetRegulator(opts *bind.TransactOpts, _regulator common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setRegulator", _regulator)
}

// SetRegulator is a paid mutator transaction binding the contract method 0xcde0a4f8.
//
// Solidity: function setRegulator(address _regulator) returns()
func (_Vault *VaultSession) SetRegulator(_regulator common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetRegulator(&_Vault.TransactOpts, _regulator)
}

// SetRegulator is a paid mutator transaction binding the contract method 0xcde0a4f8.
//
// Solidity: function setRegulator(address _regulator) returns()
func (_Vault *VaultTransactorSession) SetRegulator(_regulator common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetRegulator(&_Vault.TransactOpts, _regulator)
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

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Vault *VaultTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Vault *VaultSession) Receive() (*types.Transaction, error) {
	return _Vault.Contract.Receive(&_Vault.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Vault *VaultTransactorSession) Receive() (*types.Transaction, error) {
	return _Vault.Contract.Receive(&_Vault.TransactOpts)
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
	event.Raw = log
	return event, nil
}

// VaultDepositV2Iterator is returned from FilterDepositV2 and is used to iterate over the raw logs and unpacked data for DepositV2 events raised by the Vault contract.
type VaultDepositV2Iterator struct {
	Event *VaultDepositV2 // Event containing the contract specifics and raw log

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
func (it *VaultDepositV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultDepositV2)
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
		it.Event = new(VaultDepositV2)
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
func (it *VaultDepositV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultDepositV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultDepositV2 represents a DepositV2 event raised by the Vault contract.
type VaultDepositV2 struct {
	Token            common.Address
	IncognitoAddress string
	Amount           *big.Int
	DepositID        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDepositV2 is a free log retrieval operation binding the contract event 0xd30df8040a1092415b49422a02dbd8cdd5915a596abcba02cd0f65dd86ab3851.
//
// Solidity: event DepositV2(address token, string incognitoAddress, uint256 amount, uint256 depositID)
func (_Vault *VaultFilterer) FilterDepositV2(opts *bind.FilterOpts) (*VaultDepositV2Iterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "DepositV2")
	if err != nil {
		return nil, err
	}
	return &VaultDepositV2Iterator{contract: _Vault.contract, event: "DepositV2", logs: logs, sub: sub}, nil
}

// WatchDepositV2 is a free log subscription operation binding the contract event 0xd30df8040a1092415b49422a02dbd8cdd5915a596abcba02cd0f65dd86ab3851.
//
// Solidity: event DepositV2(address token, string incognitoAddress, uint256 amount, uint256 depositID)
func (_Vault *VaultFilterer) WatchDepositV2(opts *bind.WatchOpts, sink chan<- *VaultDepositV2) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "DepositV2")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultDepositV2)
				if err := _Vault.contract.UnpackLog(event, "DepositV2", log); err != nil {
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

// ParseDepositV2 is a log parse operation binding the contract event 0xd30df8040a1092415b49422a02dbd8cdd5915a596abcba02cd0f65dd86ab3851.
//
// Solidity: event DepositV2(address token, string incognitoAddress, uint256 amount, uint256 depositID)
func (_Vault *VaultFilterer) ParseDepositV2(log types.Log) (*VaultDepositV2, error) {
	event := new(VaultDepositV2)
	if err := _Vault.contract.UnpackLog(event, "DepositV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultExecuteFnLogIterator is returned from FilterExecuteFnLog and is used to iterate over the raw logs and unpacked data for ExecuteFnLog events raised by the Vault contract.
type VaultExecuteFnLogIterator struct {
	Event *VaultExecuteFnLog // Event containing the contract specifics and raw log

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
func (it *VaultExecuteFnLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultExecuteFnLog)
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
		it.Event = new(VaultExecuteFnLog)
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
func (it *VaultExecuteFnLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultExecuteFnLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultExecuteFnLog represents a ExecuteFnLog event raised by the Vault contract.
type VaultExecuteFnLog struct {
	Id        [32]byte
	PhaseID   *big.Int
	ErrorData []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExecuteFnLog is a free log retrieval operation binding the contract event 0xdbbb883f24557adf486292429863dcfd4ac5d4db168ae94921da8e3d9a95d416.
//
// Solidity: event ExecuteFnLog(bytes32 id, uint256 phaseID, bytes errorData)
func (_Vault *VaultFilterer) FilterExecuteFnLog(opts *bind.FilterOpts) (*VaultExecuteFnLogIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "ExecuteFnLog")
	if err != nil {
		return nil, err
	}
	return &VaultExecuteFnLogIterator{contract: _Vault.contract, event: "ExecuteFnLog", logs: logs, sub: sub}, nil
}

// WatchExecuteFnLog is a free log subscription operation binding the contract event 0xdbbb883f24557adf486292429863dcfd4ac5d4db168ae94921da8e3d9a95d416.
//
// Solidity: event ExecuteFnLog(bytes32 id, uint256 phaseID, bytes errorData)
func (_Vault *VaultFilterer) WatchExecuteFnLog(opts *bind.WatchOpts, sink chan<- *VaultExecuteFnLog) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "ExecuteFnLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultExecuteFnLog)
				if err := _Vault.contract.UnpackLog(event, "ExecuteFnLog", log); err != nil {
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

// ParseExecuteFnLog is a log parse operation binding the contract event 0xdbbb883f24557adf486292429863dcfd4ac5d4db168ae94921da8e3d9a95d416.
//
// Solidity: event ExecuteFnLog(bytes32 id, uint256 phaseID, bytes errorData)
func (_Vault *VaultFilterer) ParseExecuteFnLog(log types.Log) (*VaultExecuteFnLog, error) {
	event := new(VaultExecuteFnLog)
	if err := _Vault.contract.UnpackLog(event, "ExecuteFnLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultRedepositIterator is returned from FilterRedeposit and is used to iterate over the raw logs and unpacked data for Redeposit events raised by the Vault contract.
type VaultRedepositIterator struct {
	Event *VaultRedeposit // Event containing the contract specifics and raw log

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
func (it *VaultRedepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultRedeposit)
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
		it.Event = new(VaultRedeposit)
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
func (it *VaultRedepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultRedepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultRedeposit represents a Redeposit event raised by the Vault contract.
type VaultRedeposit struct {
	Token               common.Address
	RedepositIncAddress []byte
	Amount              *big.Int
	Itx                 [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterRedeposit is a free log retrieval operation binding the contract event 0x00b45d95b5117447e2fafe7f34def913ff3ba220e4b8688acf37ae2328af7a3d.
//
// Solidity: event Redeposit(address token, bytes redepositIncAddress, uint256 amount, bytes32 itx)
func (_Vault *VaultFilterer) FilterRedeposit(opts *bind.FilterOpts) (*VaultRedepositIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Redeposit")
	if err != nil {
		return nil, err
	}
	return &VaultRedepositIterator{contract: _Vault.contract, event: "Redeposit", logs: logs, sub: sub}, nil
}

// WatchRedeposit is a free log subscription operation binding the contract event 0x00b45d95b5117447e2fafe7f34def913ff3ba220e4b8688acf37ae2328af7a3d.
//
// Solidity: event Redeposit(address token, bytes redepositIncAddress, uint256 amount, bytes32 itx)
func (_Vault *VaultFilterer) WatchRedeposit(opts *bind.WatchOpts, sink chan<- *VaultRedeposit) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Redeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultRedeposit)
				if err := _Vault.contract.UnpackLog(event, "Redeposit", log); err != nil {
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

// ParseRedeposit is a log parse operation binding the contract event 0x00b45d95b5117447e2fafe7f34def913ff3ba220e4b8688acf37ae2328af7a3d.
//
// Solidity: event Redeposit(address token, bytes redepositIncAddress, uint256 amount, bytes32 itx)
func (_Vault *VaultFilterer) ParseRedeposit(log types.Log) (*VaultRedeposit, error) {
	event := new(VaultRedeposit)
	if err := _Vault.contract.UnpackLog(event, "Redeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
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
	event.Raw = log
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
	event.Raw = log
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
	event.Raw = log
	return event, nil
}

// WithdrawableMetaData contains all meta data concerning the Withdrawable contract.
var WithdrawableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getDepositedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isSigDataUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isWithdrawed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"name\":\"updateAssets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"f75b98ce": "getDepositedBalance(address,address)",
		"e4bd7074": "isSigDataUsed(bytes32)",
		"749c5f86": "isWithdrawed(bytes32)",
		"5c975abb": "paused()",
		"1ed4276d": "updateAssets(address[],uint256[])",
	},
}

// WithdrawableABI is the input ABI used to generate the binding from.
// Deprecated: Use WithdrawableMetaData.ABI instead.
var WithdrawableABI = WithdrawableMetaData.ABI

// Deprecated: Use WithdrawableMetaData.Sigs instead.
// WithdrawableFuncSigs maps the 4-byte function signature to its string representation.
var WithdrawableFuncSigs = WithdrawableMetaData.Sigs

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
func (_Withdrawable *WithdrawableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_Withdrawable *WithdrawableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
	var out []interface{}
	err := _Withdrawable.contract.Call(opts, &out, "getDepositedBalance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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
	var out []interface{}
	err := _Withdrawable.contract.Call(opts, &out, "isSigDataUsed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

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
	var out []interface{}
	err := _Withdrawable.contract.Call(opts, &out, "isWithdrawed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

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
	var out []interface{}
	err := _Withdrawable.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

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
