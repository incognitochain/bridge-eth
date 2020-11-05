// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vaultproxy

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

// VaultproxyABI is the input ABI used to generate the binding from.
const VaultproxyABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousSuccessor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newSuccessor\",\"type\":\"address\"}],\"name\":\"SuccessorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newSuccessor\",\"type\":\"address\"}],\"name\":\"retire\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"successor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// VaultproxyBin is the compiled bytecode used for deploying new contracts.
var VaultproxyBin = "0x6080604052604051620010f6380380620010f6833981810160405260608110156200002957600080fd5b810190808051906020019092919080519060200190929190805160405193929190846401000000008211156200005e57600080fd5b838201915060208201858111156200007557600080fd5b82518660018202830111640100000000821117156200009357600080fd5b8083526020830192505050908051906020019080838360005b83811015620000c9578082015181840152602081019050620000ac565b50505050905090810190601f168015620000f75780820380516001836020036101000a031916815260200191505b506040525050508281600160405180807f656970313936372e70726f78792e696d706c656d656e746174696f6e00000000815250601c019050604051809103902060001c0360001b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b146200016b57fe5b6200017c826200034060201b60201c565b6000815111156200024e5760008273ffffffffffffffffffffffffffffffffffffffff16826040518082805190602001908083835b60208310620001d65780518252602082019150602081019050602083039250620001b1565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d806000811462000238576040519150601f19603f3d011682016040523d82523d6000602084013e6200023d565b606091505b50509050806200024c57600080fd5b505b5050600160405180807f656970313936372e70726f78792e61646d696e000000000000000000000000008152506013019050604051809103902060001c0360001b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610360001b14620002bb57fe5b600160405180807f656970313936372e70726f78792e737563636573736f720000000000000000008152506017019050604051809103902060001c0360001b7f7b13fc932b1063ca775d428558b73e20eab6804d4d9b5a148d7cbae4488973f860001b146200032657fe5b6200033782620003d760201b60201c565b50505062000419565b62000351816200040660201b60201c565b620003a8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526036815260200180620010c06036913960400191505060405180910390fd5b60007f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b90508181555050565b60007fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610360001b90508181555050565b600080823b905060008111915050919050565b610c9780620004296000396000f3fe60806040526004361061007f5760003560e01c80636ff968c31161004e5780636ff968c3146101f05780638f283970146102475780639e6371ba14610298578063f851a440146102e95761008e565b80633659cfe6146100985780634e71d92d146100e95780634f1ef286146101005780635c60da1b146101995761008e565b3661008e5761008c610340565b005b610096610340565b005b3480156100a457600080fd5b506100e7600480360360208110156100bb57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061035a565b005b3480156100f557600080fd5b506100fe6103af565b005b6101976004803603604081101561011657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019064010000000081111561015357600080fd5b82018360208201111561016557600080fd5b8035906020019184600183028401116401000000008311171561018757600080fd5b90919293919293905050506104b6565b005b3480156101a557600080fd5b506101ae61058c565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156101fc57600080fd5b506102056105e4565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561025357600080fd5b506102966004803603602081101561026a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061063c565b005b3480156102a457600080fd5b506102e7600480360360208110156102bb57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506107b5565b005b3480156102f557600080fd5b506102fe61092e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610348610986565b610358610353610988565b6109b9565b565b6103626109df565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156103a35761039e81610a10565b6103ac565b6103ab610340565b5b50565b6103b7610a5f565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461043a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526029815260200180610b8f6029913960400191505060405180910390fd5b7f0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc610463610a5f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a16104b46104af610a5f565b610a90565b565b6104be6109df565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561057e576104fa83610a10565b60008373ffffffffffffffffffffffffffffffffffffffff168383604051808383808284378083019250505092505050600060405180830381855af49150503d8060008114610565576040519150601f19603f3d011682016040523d82523d6000602084013e61056a565b606091505b505090508061057857600080fd5b50610587565b610586610340565b5b505050565b60006105966109df565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156105d8576105d1610988565b90506105e1565b6105e0610340565b5b90565b60006105ee6109df565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561063057610629610a5f565b9050610639565b610638610340565b5b90565b6106446109df565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156107a957600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156106fd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603a815260200180610bb8603a913960400191505060405180910390fd5b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f6107266109df565b82604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a16107a481610a90565b6107b2565b6107b1610340565b5b50565b6107bd6109df565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561092257600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610876576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603a815260200180610bf2603a913960400191505060405180910390fd5b7ff966f857c3c376f2e1df873bbe2596a18675dc056dc3465dfbbe8fe9ac02c97461089f610a5f565b82604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a161091d81610abf565b61092b565b61092a610340565b5b50565b60006109386109df565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561097a576109736109df565b9050610983565b610982610340565b5b90565b565b6000807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b9050805491505090565b3660008037600080366000845af43d6000803e80600081146109da573d6000f35b3d6000fd5b6000807fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610360001b9050805491505090565b610a1981610aee565b8073ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a250565b6000807f7b13fc932b1063ca775d428558b73e20eab6804d4d9b5a148d7cbae4488973f860001b9050805491505090565b60007fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610360001b90508181555050565b60007f7b13fc932b1063ca775d428558b73e20eab6804d4d9b5a148d7cbae4488973f860001b90508181555050565b610af781610b7b565b610b4c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526036815260200180610c2c6036913960400191505060405180910390fd5b60007f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b90508181555050565b600080823b90506000811191505091905056fe5472616e73706172656e745570677261646561626c6550726f78793a20756e617574686f72697a65645472616e73706172656e745570677261646561626c6550726f78793a206e65772061646d696e20697320746865207a65726f20616464726573735472616e73706172656e745570677261646561626c6550726f78793a20737563636573736f7220697320746865207a65726f20616464726573735570677261646561626c6550726f78793a206e657720696d706c656d656e746174696f6e206973206e6f74206120636f6e7472616374a2646970667358221220444f6485acc12dc715ce74bdbc644b5deefe257e3e5e057b4e2a5ade820457dc64736f6c634300060600335570677261646561626c6550726f78793a206e657720696d706c656d656e746174696f6e206973206e6f74206120636f6e7472616374"

// DeployVaultproxy deploys a new Ethereum contract, binding an instance of Vaultproxy to it.
func DeployVaultproxy(auth *bind.TransactOpts, backend bind.ContractBackend, _logic common.Address, _admin common.Address, _data []byte) (common.Address, *types.Transaction, *Vaultproxy, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultproxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VaultproxyBin), backend, _logic, _admin, _data)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Vaultproxy{VaultproxyCaller: VaultproxyCaller{contract: contract}, VaultproxyTransactor: VaultproxyTransactor{contract: contract}, VaultproxyFilterer: VaultproxyFilterer{contract: contract}}, nil
}

// Vaultproxy is an auto generated Go binding around an Ethereum contract.
type Vaultproxy struct {
	VaultproxyCaller     // Read-only binding to the contract
	VaultproxyTransactor // Write-only binding to the contract
	VaultproxyFilterer   // Log filterer for contract events
}

// VaultproxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultproxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultproxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultproxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultproxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultproxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultproxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultproxySession struct {
	Contract     *Vaultproxy       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultproxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultproxyCallerSession struct {
	Contract *VaultproxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// VaultproxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultproxyTransactorSession struct {
	Contract     *VaultproxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// VaultproxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultproxyRaw struct {
	Contract *Vaultproxy // Generic contract binding to access the raw methods on
}

// VaultproxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultproxyCallerRaw struct {
	Contract *VaultproxyCaller // Generic read-only contract binding to access the raw methods on
}

// VaultproxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultproxyTransactorRaw struct {
	Contract *VaultproxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVaultproxy creates a new instance of Vaultproxy, bound to a specific deployed contract.
func NewVaultproxy(address common.Address, backend bind.ContractBackend) (*Vaultproxy, error) {
	contract, err := bindVaultproxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vaultproxy{VaultproxyCaller: VaultproxyCaller{contract: contract}, VaultproxyTransactor: VaultproxyTransactor{contract: contract}, VaultproxyFilterer: VaultproxyFilterer{contract: contract}}, nil
}

// NewVaultproxyCaller creates a new read-only instance of Vaultproxy, bound to a specific deployed contract.
func NewVaultproxyCaller(address common.Address, caller bind.ContractCaller) (*VaultproxyCaller, error) {
	contract, err := bindVaultproxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultproxyCaller{contract: contract}, nil
}

// NewVaultproxyTransactor creates a new write-only instance of Vaultproxy, bound to a specific deployed contract.
func NewVaultproxyTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultproxyTransactor, error) {
	contract, err := bindVaultproxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultproxyTransactor{contract: contract}, nil
}

// NewVaultproxyFilterer creates a new log filterer instance of Vaultproxy, bound to a specific deployed contract.
func NewVaultproxyFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultproxyFilterer, error) {
	contract, err := bindVaultproxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultproxyFilterer{contract: contract}, nil
}

// bindVaultproxy binds a generic wrapper to an already deployed contract.
func bindVaultproxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultproxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vaultproxy *VaultproxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Vaultproxy.Contract.VaultproxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vaultproxy *VaultproxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vaultproxy.Contract.VaultproxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vaultproxy *VaultproxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vaultproxy.Contract.VaultproxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vaultproxy *VaultproxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Vaultproxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vaultproxy *VaultproxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vaultproxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vaultproxy *VaultproxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vaultproxy.Contract.contract.Transact(opts, method, params...)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_Vaultproxy *VaultproxyTransactor) Admin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vaultproxy.contract.Transact(opts, "admin")
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_Vaultproxy *VaultproxySession) Admin() (*types.Transaction, error) {
	return _Vaultproxy.Contract.Admin(&_Vaultproxy.TransactOpts)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_Vaultproxy *VaultproxyTransactorSession) Admin() (*types.Transaction, error) {
	return _Vaultproxy.Contract.Admin(&_Vaultproxy.TransactOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_Vaultproxy *VaultproxyTransactor) ChangeAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Vaultproxy.contract.Transact(opts, "changeAdmin", newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_Vaultproxy *VaultproxySession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Vaultproxy.Contract.ChangeAdmin(&_Vaultproxy.TransactOpts, newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_Vaultproxy *VaultproxyTransactorSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Vaultproxy.Contract.ChangeAdmin(&_Vaultproxy.TransactOpts, newAdmin)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Vaultproxy *VaultproxyTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vaultproxy.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Vaultproxy *VaultproxySession) Claim() (*types.Transaction, error) {
	return _Vaultproxy.Contract.Claim(&_Vaultproxy.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Vaultproxy *VaultproxyTransactorSession) Claim() (*types.Transaction, error) {
	return _Vaultproxy.Contract.Claim(&_Vaultproxy.TransactOpts)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_Vaultproxy *VaultproxyTransactor) Implementation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vaultproxy.contract.Transact(opts, "implementation")
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_Vaultproxy *VaultproxySession) Implementation() (*types.Transaction, error) {
	return _Vaultproxy.Contract.Implementation(&_Vaultproxy.TransactOpts)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_Vaultproxy *VaultproxyTransactorSession) Implementation() (*types.Transaction, error) {
	return _Vaultproxy.Contract.Implementation(&_Vaultproxy.TransactOpts)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address newSuccessor) returns()
func (_Vaultproxy *VaultproxyTransactor) Retire(opts *bind.TransactOpts, newSuccessor common.Address) (*types.Transaction, error) {
	return _Vaultproxy.contract.Transact(opts, "retire", newSuccessor)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address newSuccessor) returns()
func (_Vaultproxy *VaultproxySession) Retire(newSuccessor common.Address) (*types.Transaction, error) {
	return _Vaultproxy.Contract.Retire(&_Vaultproxy.TransactOpts, newSuccessor)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address newSuccessor) returns()
func (_Vaultproxy *VaultproxyTransactorSession) Retire(newSuccessor common.Address) (*types.Transaction, error) {
	return _Vaultproxy.Contract.Retire(&_Vaultproxy.TransactOpts, newSuccessor)
}

// Successor is a paid mutator transaction binding the contract method 0x6ff968c3.
//
// Solidity: function successor() returns(address)
func (_Vaultproxy *VaultproxyTransactor) Successor(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vaultproxy.contract.Transact(opts, "successor")
}

// Successor is a paid mutator transaction binding the contract method 0x6ff968c3.
//
// Solidity: function successor() returns(address)
func (_Vaultproxy *VaultproxySession) Successor() (*types.Transaction, error) {
	return _Vaultproxy.Contract.Successor(&_Vaultproxy.TransactOpts)
}

// Successor is a paid mutator transaction binding the contract method 0x6ff968c3.
//
// Solidity: function successor() returns(address)
func (_Vaultproxy *VaultproxyTransactorSession) Successor() (*types.Transaction, error) {
	return _Vaultproxy.Contract.Successor(&_Vaultproxy.TransactOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Vaultproxy *VaultproxyTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Vaultproxy.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Vaultproxy *VaultproxySession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Vaultproxy.Contract.UpgradeTo(&_Vaultproxy.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Vaultproxy *VaultproxyTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Vaultproxy.Contract.UpgradeTo(&_Vaultproxy.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Vaultproxy *VaultproxyTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Vaultproxy.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Vaultproxy *VaultproxySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Vaultproxy.Contract.UpgradeToAndCall(&_Vaultproxy.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Vaultproxy *VaultproxyTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Vaultproxy.Contract.UpgradeToAndCall(&_Vaultproxy.TransactOpts, newImplementation, data)
}

// VaultproxyAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Vaultproxy contract.
type VaultproxyAdminChangedIterator struct {
	Event *VaultproxyAdminChanged // Event containing the contract specifics and raw log

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
func (it *VaultproxyAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultproxyAdminChanged)
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
		it.Event = new(VaultproxyAdminChanged)
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
func (it *VaultproxyAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultproxyAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultproxyAdminChanged represents a AdminChanged event raised by the Vaultproxy contract.
type VaultproxyAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Vaultproxy *VaultproxyFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*VaultproxyAdminChangedIterator, error) {

	logs, sub, err := _Vaultproxy.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &VaultproxyAdminChangedIterator{contract: _Vaultproxy.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Vaultproxy *VaultproxyFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *VaultproxyAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Vaultproxy.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultproxyAdminChanged)
				if err := _Vaultproxy.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Vaultproxy *VaultproxyFilterer) ParseAdminChanged(log types.Log) (*VaultproxyAdminChanged, error) {
	event := new(VaultproxyAdminChanged)
	if err := _Vaultproxy.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VaultproxyClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the Vaultproxy contract.
type VaultproxyClaimIterator struct {
	Event *VaultproxyClaim // Event containing the contract specifics and raw log

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
func (it *VaultproxyClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultproxyClaim)
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
		it.Event = new(VaultproxyClaim)
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
func (it *VaultproxyClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultproxyClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultproxyClaim represents a Claim event raised by the Vaultproxy contract.
type VaultproxyClaim struct {
	Claimer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_Vaultproxy *VaultproxyFilterer) FilterClaim(opts *bind.FilterOpts) (*VaultproxyClaimIterator, error) {

	logs, sub, err := _Vaultproxy.contract.FilterLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return &VaultproxyClaimIterator{contract: _Vaultproxy.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_Vaultproxy *VaultproxyFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *VaultproxyClaim) (event.Subscription, error) {

	logs, sub, err := _Vaultproxy.contract.WatchLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultproxyClaim)
				if err := _Vaultproxy.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_Vaultproxy *VaultproxyFilterer) ParseClaim(log types.Log) (*VaultproxyClaim, error) {
	event := new(VaultproxyClaim)
	if err := _Vaultproxy.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VaultproxySuccessorChangedIterator is returned from FilterSuccessorChanged and is used to iterate over the raw logs and unpacked data for SuccessorChanged events raised by the Vaultproxy contract.
type VaultproxySuccessorChangedIterator struct {
	Event *VaultproxySuccessorChanged // Event containing the contract specifics and raw log

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
func (it *VaultproxySuccessorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultproxySuccessorChanged)
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
		it.Event = new(VaultproxySuccessorChanged)
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
func (it *VaultproxySuccessorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultproxySuccessorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultproxySuccessorChanged represents a SuccessorChanged event raised by the Vaultproxy contract.
type VaultproxySuccessorChanged struct {
	PreviousSuccessor common.Address
	NewSuccessor      common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSuccessorChanged is a free log retrieval operation binding the contract event 0xf966f857c3c376f2e1df873bbe2596a18675dc056dc3465dfbbe8fe9ac02c974.
//
// Solidity: event SuccessorChanged(address previousSuccessor, address newSuccessor)
func (_Vaultproxy *VaultproxyFilterer) FilterSuccessorChanged(opts *bind.FilterOpts) (*VaultproxySuccessorChangedIterator, error) {

	logs, sub, err := _Vaultproxy.contract.FilterLogs(opts, "SuccessorChanged")
	if err != nil {
		return nil, err
	}
	return &VaultproxySuccessorChangedIterator{contract: _Vaultproxy.contract, event: "SuccessorChanged", logs: logs, sub: sub}, nil
}

// WatchSuccessorChanged is a free log subscription operation binding the contract event 0xf966f857c3c376f2e1df873bbe2596a18675dc056dc3465dfbbe8fe9ac02c974.
//
// Solidity: event SuccessorChanged(address previousSuccessor, address newSuccessor)
func (_Vaultproxy *VaultproxyFilterer) WatchSuccessorChanged(opts *bind.WatchOpts, sink chan<- *VaultproxySuccessorChanged) (event.Subscription, error) {

	logs, sub, err := _Vaultproxy.contract.WatchLogs(opts, "SuccessorChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultproxySuccessorChanged)
				if err := _Vaultproxy.contract.UnpackLog(event, "SuccessorChanged", log); err != nil {
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

// ParseSuccessorChanged is a log parse operation binding the contract event 0xf966f857c3c376f2e1df873bbe2596a18675dc056dc3465dfbbe8fe9ac02c974.
//
// Solidity: event SuccessorChanged(address previousSuccessor, address newSuccessor)
func (_Vaultproxy *VaultproxyFilterer) ParseSuccessorChanged(log types.Log) (*VaultproxySuccessorChanged, error) {
	event := new(VaultproxySuccessorChanged)
	if err := _Vaultproxy.contract.UnpackLog(event, "SuccessorChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VaultproxyUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Vaultproxy contract.
type VaultproxyUpgradedIterator struct {
	Event *VaultproxyUpgraded // Event containing the contract specifics and raw log

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
func (it *VaultproxyUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultproxyUpgraded)
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
		it.Event = new(VaultproxyUpgraded)
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
func (it *VaultproxyUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultproxyUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultproxyUpgraded represents a Upgraded event raised by the Vaultproxy contract.
type VaultproxyUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Vaultproxy *VaultproxyFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*VaultproxyUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Vaultproxy.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &VaultproxyUpgradedIterator{contract: _Vaultproxy.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Vaultproxy *VaultproxyFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *VaultproxyUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Vaultproxy.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultproxyUpgraded)
				if err := _Vaultproxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
// Solidity: event Upgraded(address indexed implementation)
func (_Vaultproxy *VaultproxyFilterer) ParseUpgraded(log types.Log) (*VaultproxyUpgraded, error) {
	event := new(VaultproxyUpgraded)
	if err := _Vaultproxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	return event, nil
}
