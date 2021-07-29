// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package usdc

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

// UsdcWrapABI is the input ABI used to generate the binding from.
const UsdcWrapABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_implementation\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"}]"

// UsdcWrapBin is the compiled bytecode used for deploying new contracts.
var UsdcWrapBin = "0x608060405234801561001057600080fd5b50604051602080610b2983398101806040528101908080519060200190929190505050808060405180807f6f72672e7a657070656c696e6f732e70726f78792e696d706c656d656e74617481526020017f696f6e000000000000000000000000000000000000000000000000000000000081525060230190506040518091039020600019167f7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c3600102600019161415156100c657fe5b6100de81610169640100000000026401000000009004565b5060405180807f6f72672e7a657070656c696e6f732e70726f78792e61646d696e000000000000815250601a0190506040518091039020600019167f10d6a54a4754c8869d6886b5f5d7fbfa5b4522237ea5c60d11bc4e7a1ff9390b6001026000191614151561014a57fe5b6101623361024e640100000000026401000000009004565b5050610290565b60006101878261027d6401000000000261084b176401000000009004565b1515610221576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603b8152602001807f43616e6e6f742073657420612070726f787920696d706c656d656e746174696f81526020017f6e20746f2061206e6f6e2d636f6e74726163742061646472657373000000000081525060400191505060405180910390fd5b7f7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c360010290508181555050565b60007f10d6a54a4754c8869d6886b5f5d7fbfa5b4522237ea5c60d11bc4e7a1ff9390b60010290508181555050565b600080823b905060008111915050919050565b61088a8061029f6000396000f30060806040526004361061006d576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680633659cfe6146100775780634f1ef286146100ba5780635c60da1b146101085780638f2839701461015f578063f851a440146101a2575b6100756101f9565b005b34801561008357600080fd5b506100b8600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610213565b005b610106600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001908201803590602001919091929391929390505050610268565b005b34801561011457600080fd5b5061011d610308565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561016b57600080fd5b506101a0600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610360565b005b3480156101ae57600080fd5b506101b761051e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610201610576565b61021161020c610651565b610682565b565b61021b6106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561025c57610257816106d9565b610265565b6102646101f9565b5b50565b6102706106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156102fa576102ac836106d9565b3073ffffffffffffffffffffffffffffffffffffffff163483836040518083838082843782019150509250505060006040518083038185875af19250505015156102f557600080fd5b610303565b6103026101f9565b5b505050565b60006103126106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156103545761034d610651565b905061035d565b61035c6101f9565b5b90565b6103686106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561051257600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610466576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260368152602001807f43616e6e6f74206368616e6765207468652061646d696e206f6620612070726f81526020017f787920746f20746865207a65726f20616464726573730000000000000000000081525060400191505060405180910390fd5b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f61048f6106a8565b82604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a161050d81610748565b61051b565b61051a6101f9565b5b50565b60006105286106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561056a576105636106a8565b9050610573565b6105726101f9565b5b90565b61057e6106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151515610647576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260328152602001807f43616e6e6f742063616c6c2066616c6c6261636b2066756e6374696f6e20667281526020017f6f6d207468652070726f78792061646d696e000000000000000000000000000081525060400191505060405180910390fd5b61064f610777565b565b6000807f7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c36001029050805491505090565b3660008037600080366000845af43d6000803e80600081146106a3573d6000f35b3d6000fd5b6000807f10d6a54a4754c8869d6886b5f5d7fbfa5b4522237ea5c60d11bc4e7a1ff9390b6001029050805491505090565b6106e281610779565b7fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b81604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b60007f10d6a54a4754c8869d6886b5f5d7fbfa5b4522237ea5c60d11bc4e7a1ff9390b60010290508181555050565b565b60006107848261084b565b151561081e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603b8152602001807f43616e6e6f742073657420612070726f787920696d706c656d656e746174696f81526020017f6e20746f2061206e6f6e2d636f6e74726163742061646472657373000000000081525060400191505060405180910390fd5b7f7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c360010290508181555050565b600080823b9050600081119150509190505600a165627a7a72305820a4a547cfc7202c5acaaae74d428e988bc62ad5024eb0165532d3a8f91db4ed2400290000000000000000000000000882477e7895bdc5cea7cb1552ed914ab157fe56"

// DeployUsdcWrap deploys a new Ethereum contract, binding an instance of UsdcWrap to it.
func DeployUsdcWrap(auth *bind.TransactOpts, backend bind.ContractBackend, _implementation common.Address) (common.Address, *types.Transaction, *UsdcWrap, error) {
	parsed, err := abi.JSON(strings.NewReader(UsdcWrapABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UsdcWrapBin), backend, _implementation)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UsdcWrap{UsdcWrapCaller: UsdcWrapCaller{contract: contract}, UsdcWrapTransactor: UsdcWrapTransactor{contract: contract}, UsdcWrapFilterer: UsdcWrapFilterer{contract: contract}}, nil
}

// UsdcWrap is an auto generated Go binding around an Ethereum contract.
type UsdcWrap struct {
	UsdcWrapCaller     // Read-only binding to the contract
	UsdcWrapTransactor // Write-only binding to the contract
	UsdcWrapFilterer   // Log filterer for contract events
}

// UsdcWrapCaller is an auto generated read-only Go binding around an Ethereum contract.
type UsdcWrapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsdcWrapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UsdcWrapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsdcWrapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UsdcWrapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsdcWrapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UsdcWrapSession struct {
	Contract     *UsdcWrap         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UsdcWrapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UsdcWrapCallerSession struct {
	Contract *UsdcWrapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// UsdcWrapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UsdcWrapTransactorSession struct {
	Contract     *UsdcWrapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// UsdcWrapRaw is an auto generated low-level Go binding around an Ethereum contract.
type UsdcWrapRaw struct {
	Contract *UsdcWrap // Generic contract binding to access the raw methods on
}

// UsdcWrapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UsdcWrapCallerRaw struct {
	Contract *UsdcWrapCaller // Generic read-only contract binding to access the raw methods on
}

// UsdcWrapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UsdcWrapTransactorRaw struct {
	Contract *UsdcWrapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUsdcWrap creates a new instance of UsdcWrap, bound to a specific deployed contract.
func NewUsdcWrap(address common.Address, backend bind.ContractBackend) (*UsdcWrap, error) {
	contract, err := bindUsdcWrap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UsdcWrap{UsdcWrapCaller: UsdcWrapCaller{contract: contract}, UsdcWrapTransactor: UsdcWrapTransactor{contract: contract}, UsdcWrapFilterer: UsdcWrapFilterer{contract: contract}}, nil
}

// NewUsdcWrapCaller creates a new read-only instance of UsdcWrap, bound to a specific deployed contract.
func NewUsdcWrapCaller(address common.Address, caller bind.ContractCaller) (*UsdcWrapCaller, error) {
	contract, err := bindUsdcWrap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UsdcWrapCaller{contract: contract}, nil
}

// NewUsdcWrapTransactor creates a new write-only instance of UsdcWrap, bound to a specific deployed contract.
func NewUsdcWrapTransactor(address common.Address, transactor bind.ContractTransactor) (*UsdcWrapTransactor, error) {
	contract, err := bindUsdcWrap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UsdcWrapTransactor{contract: contract}, nil
}

// NewUsdcWrapFilterer creates a new log filterer instance of UsdcWrap, bound to a specific deployed contract.
func NewUsdcWrapFilterer(address common.Address, filterer bind.ContractFilterer) (*UsdcWrapFilterer, error) {
	contract, err := bindUsdcWrap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UsdcWrapFilterer{contract: contract}, nil
}

// bindUsdcWrap binds a generic wrapper to an already deployed contract.
func bindUsdcWrap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UsdcWrapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UsdcWrap *UsdcWrapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UsdcWrap.Contract.UsdcWrapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UsdcWrap *UsdcWrapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UsdcWrap.Contract.UsdcWrapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UsdcWrap *UsdcWrapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UsdcWrap.Contract.UsdcWrapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UsdcWrap *UsdcWrapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UsdcWrap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UsdcWrap *UsdcWrapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UsdcWrap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UsdcWrap *UsdcWrapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UsdcWrap.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_UsdcWrap *UsdcWrapCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UsdcWrap.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_UsdcWrap *UsdcWrapSession) Admin() (common.Address, error) {
	return _UsdcWrap.Contract.Admin(&_UsdcWrap.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_UsdcWrap *UsdcWrapCallerSession) Admin() (common.Address, error) {
	return _UsdcWrap.Contract.Admin(&_UsdcWrap.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_UsdcWrap *UsdcWrapCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UsdcWrap.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_UsdcWrap *UsdcWrapSession) Implementation() (common.Address, error) {
	return _UsdcWrap.Contract.Implementation(&_UsdcWrap.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_UsdcWrap *UsdcWrapCallerSession) Implementation() (common.Address, error) {
	return _UsdcWrap.Contract.Implementation(&_UsdcWrap.CallOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_UsdcWrap *UsdcWrapTransactor) ChangeAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _UsdcWrap.contract.Transact(opts, "changeAdmin", newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_UsdcWrap *UsdcWrapSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _UsdcWrap.Contract.ChangeAdmin(&_UsdcWrap.TransactOpts, newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_UsdcWrap *UsdcWrapTransactorSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _UsdcWrap.Contract.ChangeAdmin(&_UsdcWrap.TransactOpts, newAdmin)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_UsdcWrap *UsdcWrapTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _UsdcWrap.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_UsdcWrap *UsdcWrapSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _UsdcWrap.Contract.UpgradeTo(&_UsdcWrap.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_UsdcWrap *UsdcWrapTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _UsdcWrap.Contract.UpgradeTo(&_UsdcWrap.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_UsdcWrap *UsdcWrapTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _UsdcWrap.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_UsdcWrap *UsdcWrapSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _UsdcWrap.Contract.UpgradeToAndCall(&_UsdcWrap.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_UsdcWrap *UsdcWrapTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _UsdcWrap.Contract.UpgradeToAndCall(&_UsdcWrap.TransactOpts, newImplementation, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_UsdcWrap *UsdcWrapTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _UsdcWrap.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_UsdcWrap *UsdcWrapSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _UsdcWrap.Contract.Fallback(&_UsdcWrap.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_UsdcWrap *UsdcWrapTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _UsdcWrap.Contract.Fallback(&_UsdcWrap.TransactOpts, calldata)
}

// UsdcWrapAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the UsdcWrap contract.
type UsdcWrapAdminChangedIterator struct {
	Event *UsdcWrapAdminChanged // Event containing the contract specifics and raw log

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
func (it *UsdcWrapAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdcWrapAdminChanged)
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
		it.Event = new(UsdcWrapAdminChanged)
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
func (it *UsdcWrapAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdcWrapAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdcWrapAdminChanged represents a AdminChanged event raised by the UsdcWrap contract.
type UsdcWrapAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_UsdcWrap *UsdcWrapFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*UsdcWrapAdminChangedIterator, error) {

	logs, sub, err := _UsdcWrap.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &UsdcWrapAdminChangedIterator{contract: _UsdcWrap.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_UsdcWrap *UsdcWrapFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *UsdcWrapAdminChanged) (event.Subscription, error) {

	logs, sub, err := _UsdcWrap.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdcWrapAdminChanged)
				if err := _UsdcWrap.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_UsdcWrap *UsdcWrapFilterer) ParseAdminChanged(log types.Log) (*UsdcWrapAdminChanged, error) {
	event := new(UsdcWrapAdminChanged)
	if err := _UsdcWrap.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UsdcWrapUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the UsdcWrap contract.
type UsdcWrapUpgradedIterator struct {
	Event *UsdcWrapUpgraded // Event containing the contract specifics and raw log

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
func (it *UsdcWrapUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdcWrapUpgraded)
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
		it.Event = new(UsdcWrapUpgraded)
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
func (it *UsdcWrapUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdcWrapUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdcWrapUpgraded represents a Upgraded event raised by the UsdcWrap contract.
type UsdcWrapUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address implementation)
func (_UsdcWrap *UsdcWrapFilterer) FilterUpgraded(opts *bind.FilterOpts) (*UsdcWrapUpgradedIterator, error) {

	logs, sub, err := _UsdcWrap.contract.FilterLogs(opts, "Upgraded")
	if err != nil {
		return nil, err
	}
	return &UsdcWrapUpgradedIterator{contract: _UsdcWrap.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address implementation)
func (_UsdcWrap *UsdcWrapFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *UsdcWrapUpgraded) (event.Subscription, error) {

	logs, sub, err := _UsdcWrap.contract.WatchLogs(opts, "Upgraded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdcWrapUpgraded)
				if err := _UsdcWrap.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address implementation)
func (_UsdcWrap *UsdcWrapFilterer) ParseUpgraded(log types.Log) (*UsdcWrapUpgraded, error) {
	event := new(UsdcWrapUpgraded)
	if err := _UsdcWrap.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
