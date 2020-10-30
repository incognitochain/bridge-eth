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
const VaultproxyABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_delegator\",\"type\":\"address\"},{\"internalType\":\"contractIncognito\",\"name\":\"_incognito\",\"type\":\"address\"},{\"internalType\":\"contractWithdrawable\",\"name\":\"_prevault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ndays\",\"type\":\"uint256\"}],\"name\":\"Extend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"ETH_TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expire\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"extend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incognito\",\"outputs\":[{\"internalType\":\"contractIncognito\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"preVault\",\"outputs\":[{\"internalType\":\"contractWithdrawable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_successor\",\"type\":\"address\"}],\"name\":\"retire\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"successor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// VaultproxyBin is the compiled bytecode used for deploying new contracts.
var VaultproxyBin = "0x60806040526005805460ff60a01b1916600160a01b17905534801561002357600080fd5b506040516109f03803806109f08339818101604052608081101561004657600080fd5b50805160208201516040830151606090930151919290916001600160a01b0384161580159061007d57506001600160a01b03831615155b801561009157506001600160a01b03821615155b61009a57600080fd5b600080546001600160a01b039586166001600160a01b0319918216179091556305a39a80420160025560038054948616948216949094179093556005805492851692841692909217909155600480549190931691161790556108ef806101016000396000f3fe6080604052600436106100c65760003560e01c806379599f961161007f5780639714378c116100595780639714378c1461030d5780639e6371ba14610337578063ce9b79301461036a578063f851a4401461037f576100cd565b806379599f96146102bc5780638456cb59146102e35780638a984538146102f8576100cd565b806319dc32a51461020c5780633f4ba83a1461023d5780634e71d92d1461025457806358bc8337146102695780635c975abb1461027e5780636ff968c3146102a7576100cd565b366100cd57005b600554600160a01b900460ff1661011f576040805162461bcd60e51b815260206004820152601160248201527018d85b881b9bdd081c99595b9d1c985b9d607a1b604482015290519081900360640190fd5b6005805460ff60a01b191690556003546040516000916001600160a01b031690829036908083838082843760405192019450600093509091505080830381855af49150503d806000811461018f576040519150601f19603f3d011682016040523d82523d6000602084013e610194565b606091505b50509050806101e3576040805162461bcd60e51b815260206004820152601660248201527519195b1959d85d194818d85b1b081c995d995c9d195960521b604482015290519081900360640190fd5b6005805460ff60a01b1916600160a01b179055604080513d80820190925290806000833e503d81f35b34801561021857600080fd5b50610221610394565b604080516001600160a01b039092168252519081900360200190f35b34801561024957600080fd5b506102526103a3565b005b34801561026057600080fd5b50610252610485565b34801561027557600080fd5b5061022161056c565b34801561028a57600080fd5b50610293610571565b604080519115158252519081900360200190f35b3480156102b357600080fd5b50610221610581565b3480156102c857600080fd5b506102d1610590565b60408051918252519081900360200190f35b3480156102ef57600080fd5b50610252610596565b34801561030457600080fd5b506102216106bb565b34801561031957600080fd5b506102526004803603602081101561033057600080fd5b50356106ca565b34801561034357600080fd5b506102526004803603602081101561035a57600080fd5b50356001600160a01b03166107ee565b34801561037657600080fd5b5061022161089b565b34801561038b57600080fd5b506102216108aa565b6004546001600160a01b031681565b6000546001600160a01b031633146103ee576040805162461bcd60e51b81526020600482015260096024820152683737ba1030b236b4b760b91b604482015290519081900360640190fd5b600154600160a01b900460ff16610443576040805162461bcd60e51b81526020600482015260146024820152736e6f7420706175736564207269676874206e6f7760601b604482015290519081900360640190fd5b6001805460ff60a01b191690556040805133815290517f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa9181900360200190a1565b60025442106104c5576040805162461bcd60e51b8152602060048201526007602482015266195e1c1a5c995960ca1b604482015290519081900360640190fd5b6001546001600160a01b03163314610513576040805162461bcd60e51b815260206004820152600c60248201526b1d5b985d5d1a1bdc9a5e995960a21b604482015290519081900360640190fd5b600154600080546001600160a01b0319166001600160a01b0392831617908190556040805191909216815290517f0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc9181900360200190a1565b600081565b600154600160a01b900460ff1681565b6001546001600160a01b031681565b60025481565b6000546001600160a01b031633146105e1576040805162461bcd60e51b81526020600482015260096024820152683737ba1030b236b4b760b91b604482015290519081900360640190fd5b600154600160a01b900460ff1615610633576040805162461bcd60e51b815260206004820152601060248201526f706175736564207269676874206e6f7760801b604482015290519081900360640190fd5b6002544210610673576040805162461bcd60e51b8152602060048201526007602482015266195e1c1a5c995960ca1b604482015290519081900360640190fd5b6001805460ff60a01b1916600160a01b1790556040805133815290517f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2589181900360200190a1565b6005546001600160a01b031681565b6000546001600160a01b03163314610715576040805162461bcd60e51b81526020600482015260096024820152683737ba1030b236b4b760b91b604482015290519081900360640190fd5b6002544210610755576040805162461bcd60e51b8152602060048201526007602482015266195e1c1a5c995960ca1b604482015290519081900360640190fd5b61016e81106107ab576040805162461bcd60e51b815260206004820152601a60248201527f63616e6e6f7420657874656e6420666f7220746f6f206c6f6e67000000000000604482015290519081900360640190fd5b600280546201518083020190556040805182815290517f02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e89181900360200190a150565b6000546001600160a01b03163314610839576040805162461bcd60e51b81526020600482015260096024820152683737ba1030b236b4b760b91b604482015290519081900360640190fd5b6002544210610879576040805162461bcd60e51b8152602060048201526007602482015266195e1c1a5c995960ca1b604482015290519081900360640190fd5b600180546001600160a01b0319166001600160a01b0392909216919091179055565b6003546001600160a01b031681565b6000546001600160a01b03168156fea2646970667358221220850fd8c782203dcaa55d6b0563d49a910701cf5643b1588f01c68cf35ec17b4764736f6c63430006060033"

// DeployVaultproxy deploys a new Ethereum contract, binding an instance of Vaultproxy to it.
func DeployVaultproxy(auth *bind.TransactOpts, backend bind.ContractBackend, _admin common.Address, _delegator common.Address, _incognito common.Address, _prevault common.Address) (common.Address, *types.Transaction, *Vaultproxy, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultproxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VaultproxyBin), backend, _admin, _delegator, _incognito, _prevault)
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

// ETHTOKEN is a free data retrieval call binding the contract method 0x58bc8337.
//
// Solidity: function ETH_TOKEN() view returns(address)
func (_Vaultproxy *VaultproxyCaller) ETHTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vaultproxy.contract.Call(opts, out, "ETH_TOKEN")
	return *ret0, err
}

// ETHTOKEN is a free data retrieval call binding the contract method 0x58bc8337.
//
// Solidity: function ETH_TOKEN() view returns(address)
func (_Vaultproxy *VaultproxySession) ETHTOKEN() (common.Address, error) {
	return _Vaultproxy.Contract.ETHTOKEN(&_Vaultproxy.CallOpts)
}

// ETHTOKEN is a free data retrieval call binding the contract method 0x58bc8337.
//
// Solidity: function ETH_TOKEN() view returns(address)
func (_Vaultproxy *VaultproxyCallerSession) ETHTOKEN() (common.Address, error) {
	return _Vaultproxy.Contract.ETHTOKEN(&_Vaultproxy.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Vaultproxy *VaultproxyCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vaultproxy.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Vaultproxy *VaultproxySession) Admin() (common.Address, error) {
	return _Vaultproxy.Contract.Admin(&_Vaultproxy.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Vaultproxy *VaultproxyCallerSession) Admin() (common.Address, error) {
	return _Vaultproxy.Contract.Admin(&_Vaultproxy.CallOpts)
}

// Delegator is a free data retrieval call binding the contract method 0xce9b7930.
//
// Solidity: function delegator() view returns(address)
func (_Vaultproxy *VaultproxyCaller) Delegator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vaultproxy.contract.Call(opts, out, "delegator")
	return *ret0, err
}

// Delegator is a free data retrieval call binding the contract method 0xce9b7930.
//
// Solidity: function delegator() view returns(address)
func (_Vaultproxy *VaultproxySession) Delegator() (common.Address, error) {
	return _Vaultproxy.Contract.Delegator(&_Vaultproxy.CallOpts)
}

// Delegator is a free data retrieval call binding the contract method 0xce9b7930.
//
// Solidity: function delegator() view returns(address)
func (_Vaultproxy *VaultproxyCallerSession) Delegator() (common.Address, error) {
	return _Vaultproxy.Contract.Delegator(&_Vaultproxy.CallOpts)
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() view returns(uint256)
func (_Vaultproxy *VaultproxyCaller) Expire(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vaultproxy.contract.Call(opts, out, "expire")
	return *ret0, err
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() view returns(uint256)
func (_Vaultproxy *VaultproxySession) Expire() (*big.Int, error) {
	return _Vaultproxy.Contract.Expire(&_Vaultproxy.CallOpts)
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() view returns(uint256)
func (_Vaultproxy *VaultproxyCallerSession) Expire() (*big.Int, error) {
	return _Vaultproxy.Contract.Expire(&_Vaultproxy.CallOpts)
}

// Incognito is a free data retrieval call binding the contract method 0x8a984538.
//
// Solidity: function incognito() view returns(address)
func (_Vaultproxy *VaultproxyCaller) Incognito(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vaultproxy.contract.Call(opts, out, "incognito")
	return *ret0, err
}

// Incognito is a free data retrieval call binding the contract method 0x8a984538.
//
// Solidity: function incognito() view returns(address)
func (_Vaultproxy *VaultproxySession) Incognito() (common.Address, error) {
	return _Vaultproxy.Contract.Incognito(&_Vaultproxy.CallOpts)
}

// Incognito is a free data retrieval call binding the contract method 0x8a984538.
//
// Solidity: function incognito() view returns(address)
func (_Vaultproxy *VaultproxyCallerSession) Incognito() (common.Address, error) {
	return _Vaultproxy.Contract.Incognito(&_Vaultproxy.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Vaultproxy *VaultproxyCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Vaultproxy.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Vaultproxy *VaultproxySession) Paused() (bool, error) {
	return _Vaultproxy.Contract.Paused(&_Vaultproxy.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Vaultproxy *VaultproxyCallerSession) Paused() (bool, error) {
	return _Vaultproxy.Contract.Paused(&_Vaultproxy.CallOpts)
}

// PreVault is a free data retrieval call binding the contract method 0x19dc32a5.
//
// Solidity: function preVault() view returns(address)
func (_Vaultproxy *VaultproxyCaller) PreVault(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vaultproxy.contract.Call(opts, out, "preVault")
	return *ret0, err
}

// PreVault is a free data retrieval call binding the contract method 0x19dc32a5.
//
// Solidity: function preVault() view returns(address)
func (_Vaultproxy *VaultproxySession) PreVault() (common.Address, error) {
	return _Vaultproxy.Contract.PreVault(&_Vaultproxy.CallOpts)
}

// PreVault is a free data retrieval call binding the contract method 0x19dc32a5.
//
// Solidity: function preVault() view returns(address)
func (_Vaultproxy *VaultproxyCallerSession) PreVault() (common.Address, error) {
	return _Vaultproxy.Contract.PreVault(&_Vaultproxy.CallOpts)
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() view returns(address)
func (_Vaultproxy *VaultproxyCaller) Successor(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vaultproxy.contract.Call(opts, out, "successor")
	return *ret0, err
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() view returns(address)
func (_Vaultproxy *VaultproxySession) Successor() (common.Address, error) {
	return _Vaultproxy.Contract.Successor(&_Vaultproxy.CallOpts)
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() view returns(address)
func (_Vaultproxy *VaultproxyCallerSession) Successor() (common.Address, error) {
	return _Vaultproxy.Contract.Successor(&_Vaultproxy.CallOpts)
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

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_Vaultproxy *VaultproxyTransactor) Extend(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _Vaultproxy.contract.Transact(opts, "extend", n)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_Vaultproxy *VaultproxySession) Extend(n *big.Int) (*types.Transaction, error) {
	return _Vaultproxy.Contract.Extend(&_Vaultproxy.TransactOpts, n)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_Vaultproxy *VaultproxyTransactorSession) Extend(n *big.Int) (*types.Transaction, error) {
	return _Vaultproxy.Contract.Extend(&_Vaultproxy.TransactOpts, n)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Vaultproxy *VaultproxyTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vaultproxy.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Vaultproxy *VaultproxySession) Pause() (*types.Transaction, error) {
	return _Vaultproxy.Contract.Pause(&_Vaultproxy.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Vaultproxy *VaultproxyTransactorSession) Pause() (*types.Transaction, error) {
	return _Vaultproxy.Contract.Pause(&_Vaultproxy.TransactOpts)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_Vaultproxy *VaultproxyTransactor) Retire(opts *bind.TransactOpts, _successor common.Address) (*types.Transaction, error) {
	return _Vaultproxy.contract.Transact(opts, "retire", _successor)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_Vaultproxy *VaultproxySession) Retire(_successor common.Address) (*types.Transaction, error) {
	return _Vaultproxy.Contract.Retire(&_Vaultproxy.TransactOpts, _successor)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_Vaultproxy *VaultproxyTransactorSession) Retire(_successor common.Address) (*types.Transaction, error) {
	return _Vaultproxy.Contract.Retire(&_Vaultproxy.TransactOpts, _successor)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Vaultproxy *VaultproxyTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vaultproxy.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Vaultproxy *VaultproxySession) Unpause() (*types.Transaction, error) {
	return _Vaultproxy.Contract.Unpause(&_Vaultproxy.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Vaultproxy *VaultproxyTransactorSession) Unpause() (*types.Transaction, error) {
	return _Vaultproxy.Contract.Unpause(&_Vaultproxy.TransactOpts)
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

// VaultproxyExtendIterator is returned from FilterExtend and is used to iterate over the raw logs and unpacked data for Extend events raised by the Vaultproxy contract.
type VaultproxyExtendIterator struct {
	Event *VaultproxyExtend // Event containing the contract specifics and raw log

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
func (it *VaultproxyExtendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultproxyExtend)
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
		it.Event = new(VaultproxyExtend)
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
func (it *VaultproxyExtendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultproxyExtendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultproxyExtend represents a Extend event raised by the Vaultproxy contract.
type VaultproxyExtend struct {
	Ndays *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterExtend is a free log retrieval operation binding the contract event 0x02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8.
//
// Solidity: event Extend(uint256 ndays)
func (_Vaultproxy *VaultproxyFilterer) FilterExtend(opts *bind.FilterOpts) (*VaultproxyExtendIterator, error) {

	logs, sub, err := _Vaultproxy.contract.FilterLogs(opts, "Extend")
	if err != nil {
		return nil, err
	}
	return &VaultproxyExtendIterator{contract: _Vaultproxy.contract, event: "Extend", logs: logs, sub: sub}, nil
}

// WatchExtend is a free log subscription operation binding the contract event 0x02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8.
//
// Solidity: event Extend(uint256 ndays)
func (_Vaultproxy *VaultproxyFilterer) WatchExtend(opts *bind.WatchOpts, sink chan<- *VaultproxyExtend) (event.Subscription, error) {

	logs, sub, err := _Vaultproxy.contract.WatchLogs(opts, "Extend")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultproxyExtend)
				if err := _Vaultproxy.contract.UnpackLog(event, "Extend", log); err != nil {
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

// ParseExtend is a log parse operation binding the contract event 0x02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8.
//
// Solidity: event Extend(uint256 ndays)
func (_Vaultproxy *VaultproxyFilterer) ParseExtend(log types.Log) (*VaultproxyExtend, error) {
	event := new(VaultproxyExtend)
	if err := _Vaultproxy.contract.UnpackLog(event, "Extend", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VaultproxyPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Vaultproxy contract.
type VaultproxyPausedIterator struct {
	Event *VaultproxyPaused // Event containing the contract specifics and raw log

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
func (it *VaultproxyPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultproxyPaused)
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
		it.Event = new(VaultproxyPaused)
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
func (it *VaultproxyPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultproxyPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultproxyPaused represents a Paused event raised by the Vaultproxy contract.
type VaultproxyPaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address pauser)
func (_Vaultproxy *VaultproxyFilterer) FilterPaused(opts *bind.FilterOpts) (*VaultproxyPausedIterator, error) {

	logs, sub, err := _Vaultproxy.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &VaultproxyPausedIterator{contract: _Vaultproxy.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address pauser)
func (_Vaultproxy *VaultproxyFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *VaultproxyPaused) (event.Subscription, error) {

	logs, sub, err := _Vaultproxy.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultproxyPaused)
				if err := _Vaultproxy.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address pauser)
func (_Vaultproxy *VaultproxyFilterer) ParsePaused(log types.Log) (*VaultproxyPaused, error) {
	event := new(VaultproxyPaused)
	if err := _Vaultproxy.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VaultproxyUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Vaultproxy contract.
type VaultproxyUnpausedIterator struct {
	Event *VaultproxyUnpaused // Event containing the contract specifics and raw log

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
func (it *VaultproxyUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultproxyUnpaused)
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
		it.Event = new(VaultproxyUnpaused)
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
func (it *VaultproxyUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultproxyUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultproxyUnpaused represents a Unpaused event raised by the Vaultproxy contract.
type VaultproxyUnpaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address pauser)
func (_Vaultproxy *VaultproxyFilterer) FilterUnpaused(opts *bind.FilterOpts) (*VaultproxyUnpausedIterator, error) {

	logs, sub, err := _Vaultproxy.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &VaultproxyUnpausedIterator{contract: _Vaultproxy.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address pauser)
func (_Vaultproxy *VaultproxyFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *VaultproxyUnpaused) (event.Subscription, error) {

	logs, sub, err := _Vaultproxy.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultproxyUnpaused)
				if err := _Vaultproxy.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address pauser)
func (_Vaultproxy *VaultproxyFilterer) ParseUnpaused(log types.Log) (*VaultproxyUnpaused, error) {
	event := new(VaultproxyUnpaused)
	if err := _Vaultproxy.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}
