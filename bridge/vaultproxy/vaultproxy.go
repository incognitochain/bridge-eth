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

// ProxyABI is the input ABI used to generate the binding from.
const ProxyABI = "[{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// Proxy is an auto generated Go binding around an Ethereum contract.
type Proxy struct {
	ProxyCaller     // Read-only binding to the contract
	ProxyTransactor // Write-only binding to the contract
	ProxyFilterer   // Log filterer for contract events
}

// ProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxySession struct {
	Contract     *Proxy            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxyCallerSession struct {
	Contract *ProxyCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxyTransactorSession struct {
	Contract     *ProxyTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxyRaw struct {
	Contract *Proxy // Generic contract binding to access the raw methods on
}

// ProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxyCallerRaw struct {
	Contract *ProxyCaller // Generic read-only contract binding to access the raw methods on
}

// ProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxyTransactorRaw struct {
	Contract *ProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxy creates a new instance of Proxy, bound to a specific deployed contract.
func NewProxy(address common.Address, backend bind.ContractBackend) (*Proxy, error) {
	contract, err := bindProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Proxy{ProxyCaller: ProxyCaller{contract: contract}, ProxyTransactor: ProxyTransactor{contract: contract}, ProxyFilterer: ProxyFilterer{contract: contract}}, nil
}

// NewProxyCaller creates a new read-only instance of Proxy, bound to a specific deployed contract.
func NewProxyCaller(address common.Address, caller bind.ContractCaller) (*ProxyCaller, error) {
	contract, err := bindProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyCaller{contract: contract}, nil
}

// NewProxyTransactor creates a new write-only instance of Proxy, bound to a specific deployed contract.
func NewProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyTransactor, error) {
	contract, err := bindProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyTransactor{contract: contract}, nil
}

// NewProxyFilterer creates a new log filterer instance of Proxy, bound to a specific deployed contract.
func NewProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxyFilterer, error) {
	contract, err := bindProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyFilterer{contract: contract}, nil
}

// bindProxy binds a generic wrapper to an already deployed contract.
func bindProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxy *ProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.ProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxy *ProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxy *ProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxy *ProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxy *ProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxy *ProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transact(opts, method, params...)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Proxy *ProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Proxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Proxy *ProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Proxy.Contract.Fallback(&_Proxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Proxy *ProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Proxy.Contract.Fallback(&_Proxy.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Proxy *ProxyTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Proxy *ProxySession) Receive() (*types.Transaction, error) {
	return _Proxy.Contract.Receive(&_Proxy.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Proxy *ProxyTransactorSession) Receive() (*types.Transaction, error) {
	return _Proxy.Contract.Receive(&_Proxy.TransactOpts)
}

// TransparentUpgradeableProxyABI is the input ABI used to generate the binding from.
const TransparentUpgradeableProxyABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_incognito\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousIncognito\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newIncognito\",\"type\":\"address\"}],\"name\":\"IncognitoChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousSuccessor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newSuccessor\",\"type\":\"address\"}],\"name\":\"SuccessorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incognito\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newSuccessor\",\"type\":\"address\"}],\"name\":\"retire\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"successor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newIncognito\",\"type\":\"address\"}],\"name\":\"upgradeIncognito\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// TransparentUpgradeableProxyFuncSigs maps the 4-byte function signature to its string representation.
var TransparentUpgradeableProxyFuncSigs = map[string]string{
	"f851a440": "admin()",
	"4e71d92d": "claim()",
	"5c60da1b": "implementation()",
	"8a984538": "incognito()",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"9e6371ba": "retire(address)",
	"6ff968c3": "successor()",
	"3f4ba83a": "unpause()",
	"1c587771": "upgradeIncognito(address)",
	"3659cfe6": "upgradeTo(address)",
	"4f1ef286": "upgradeToAndCall(address,bytes)",
}

// TransparentUpgradeableProxyBin is the compiled bytecode used for deploying new contracts.
var TransparentUpgradeableProxyBin = "0x608060405260405162000f5838038062000f58833981810160405260808110156200002957600080fd5b8151602083015160408085015160608601805192519496939591949391820192846401000000008211156200005d57600080fd5b9083019060208201858111156200007357600080fd5b82516401000000008111828201881017156200008e57600080fd5b82525081516020918201929091019080838360005b83811015620000bd578181015183820152602001620000a3565b50505050905090810190601f168015620000eb5780820380516001836020036101000a031916815260200191505b5060405250859150829050620001018262000230565b80511562000207576000826001600160a01b0316826040518082805190602001908083835b60208310620001475780518252601f19909201916020918201910162000126565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d8060008114620001a9576040519150601f19603f3d011682016040523d82523d6000602084013e620001ae565b606091505b505090508062000205576040805162461bcd60e51b815260206004820152601360248201527f44454c454741544543414c4c206661696c656400000000000000000000000000604482015290519081900360640190fd5b505b50620002109050565b6200021b836200029c565b6200022682620002c0565b50505050620002ea565b6200023b81620002e4565b620002785760405162461bcd60e51b815260040180806020018281038252603681526020018062000f226036913960400191505060405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc55565b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610355565b7f62135fc083646fdb4e1a9d700e351b886a4a5a39da980650269edd1ade91ffd255565b3b151590565b610c2880620002fa6000396000f3fe6080604052600436106100ab5760003560e01c80635c975abb116100645780635c975abb146102035780636ff968c31461022c5780638456cb59146102415780638a984538146102565780639e6371ba1461026b578063f851a4401461029e576100ba565b80631c587771146100c25780633659cfe6146100f55780633f4ba83a146101285780634e71d92d1461013d5780634f1ef286146101525780635c60da1b146101d2576100ba565b366100ba576100b86102b3565b005b6100b86102b3565b3480156100ce57600080fd5b506100b8600480360360208110156100e557600080fd5b50356001600160a01b03166102cd565b34801561010157600080fd5b506100b86004803603602081101561011857600080fd5b50356001600160a01b0316610397565b34801561013457600080fd5b506100b86103c1565b34801561014957600080fd5b506100b861043c565b6100b86004803603604081101561016857600080fd5b6001600160a01b03823516919081019060408101602082013564010000000081111561019357600080fd5b8201836020820111156101a557600080fd5b803590602001918460018302840111640100000000831117156101c757600080fd5b5090925090506104b1565b3480156101de57600080fd5b506101e761059c565b604080516001600160a01b039092168252519081900360200190f35b34801561020f57600080fd5b506102186105d9565b604080519115158252519081900360200190f35b34801561023857600080fd5b506101e7610604565b34801561024d57600080fd5b506100b861062f565b34801561026257600080fd5b506101e761069e565b34801561027757600080fd5b506100b86004803603602081101561028e57600080fd5b50356001600160a01b03166106c9565b3480156102aa57600080fd5b506101e7610783565b6102bb6107ae565b6102cb6102c6610852565b610877565b565b6102d561089b565b6001600160a01b0316336001600160a01b0316141561038c576001600160a01b0381166103335760405162461bcd60e51b8152600401808060200182810382526040815260200180610ade6040913960400191505060405180910390fd5b7f86d392a76e88298144124db3dd7265135d76810f52d747dc329a0f7722135e5c61035c6108c0565b604080516001600160a01b03928316815291841660208301528051918290030190a1610387816108e5565b610394565b6103946102b3565b50565b61039f61089b565b6001600160a01b0316336001600160a01b0316141561038c5761038781610909565b6103c961089b565b6001600160a01b0316336001600160a01b03161415610434576103ea610949565b6104255760405162461bcd60e51b8152600401808060200182810382526030815260200180610b816030913960400191505060405180910390fd5b61042f600061096e565b6102cb565b6102cb6102b3565b610444610992565b6001600160a01b0316336001600160a01b03161415610434577f0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc610486610992565b604080516001600160a01b039092168252519081900360200190a161042f6104ac610992565b6109b7565b6104b961089b565b6001600160a01b0316336001600160a01b0316141561058f576104db83610909565b6000836001600160a01b031683836040518083838082843760405192019450600093509091505080830381855af49150503d8060008114610538576040519150601f19603f3d011682016040523d82523d6000602084013e61053d565b606091505b5050905080610589576040805162461bcd60e51b81526020600482015260136024820152721111531151d0551150d053130819985a5b1959606a1b604482015290519081900360640190fd5b50610597565b6105976102b3565b505050565b60006105a661089b565b6001600160a01b0316336001600160a01b031614156105ce576105c7610852565b90506105d6565b6105d66102b3565b90565b60006105e361089b565b6001600160a01b0316336001600160a01b031614156105ce576105c7610949565b600061060e61089b565b6001600160a01b0316336001600160a01b031614156105ce576105c7610992565b61063761089b565b6001600160a01b0316336001600160a01b0316141561043457610658610949565b156106945760405162461bcd60e51b8152600401808060200182810382526034815260200180610b1e6034913960400191505060405180910390fd5b61042f600161096e565b60006106a861089b565b6001600160a01b0316336001600160a01b031614156105ce576105c76108c0565b6106d161089b565b6001600160a01b0316336001600160a01b0316141561038c576001600160a01b03811661072f5760405162461bcd60e51b815260040180806020018281038252603a815260200180610a6e603a913960400191505060405180910390fd5b7ff966f857c3c376f2e1df873bbe2596a18675dc056dc3465dfbbe8fe9ac02c974610758610992565b604080516001600160a01b03928316815291841660208301528051918290030190a1610387816109db565b600061078d61089b565b6001600160a01b0316336001600160a01b031614156105ce576105c761089b565b6107b661089b565b6001600160a01b0316336001600160a01b031614156108065760405162461bcd60e51b8152600401808060200182810382526042815260200180610bb16042913960600191505060405180910390fd5b61080e610949565b1561084a5760405162461bcd60e51b815260040180806020018281038252602f815260200180610b52602f913960400191505060405180910390fd5b6102cb6102cb565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5490565b3660008037600080366000845af43d6000803e808015610896573d6000f35b3d6000fd5b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035490565b7f62135fc083646fdb4e1a9d700e351b886a4a5a39da980650269edd1ade91ffd25490565b7f62135fc083646fdb4e1a9d700e351b886a4a5a39da980650269edd1ade91ffd255565b610912816109ff565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b7f8dea8703c3cf94703383ce38a9c894669dccd4ca8e65ddb43267aa02487114505490565b7f8dea8703c3cf94703383ce38a9c894669dccd4ca8e65ddb43267aa024871145055565b7f7b13fc932b1063ca775d428558b73e20eab6804d4d9b5a148d7cbae4488973f85490565b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610355565b7f7b13fc932b1063ca775d428558b73e20eab6804d4d9b5a148d7cbae4488973f855565b610a0881610a67565b610a435760405162461bcd60e51b8152600401808060200182810382526036815260200180610aa86036913960400191505060405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc55565b3b15159056fe5472616e73706172656e745570677261646561626c6550726f78793a20737563636573736f7220697320746865207a65726f20616464726573735570677261646561626c6550726f78793a206e657720696d706c656d656e746174696f6e206973206e6f74206120636f6e74726163745472616e73706172656e745570677261646561626c6550726f78793a20696e636f676e69746f2070726f787920697320746865207a65726f20616464726573735472616e73706172656e745570677261646561626c6550726f78793a20636f6e74726163742070617573656420616c72656164795472616e73706172656e745570677261646561626c6550726f78793a20636f6e7472616374206973207061757365645472616e73706172656e745570677261646561626c6550726f78793a20636f6e7472616374206e6f74207061757365645472616e73706172656e745570677261646561626c6550726f78793a2061646d696e2063616e6e6f742066616c6c6261636b20746f2070726f787920746172676574a264697066735822122047ec491508f68ed23aba597fef69ed3072d05f20af1cf4eeee23330c738f6cae64736f6c634300060c00335570677261646561626c6550726f78793a206e657720696d706c656d656e746174696f6e206973206e6f74206120636f6e7472616374"

// DeployTransparentUpgradeableProxy deploys a new Ethereum contract, binding an instance of TransparentUpgradeableProxy to it.
func DeployTransparentUpgradeableProxy(auth *bind.TransactOpts, backend bind.ContractBackend, _logic common.Address, _admin common.Address, _incognito common.Address, _data []byte) (common.Address, *types.Transaction, *TransparentUpgradeableProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(TransparentUpgradeableProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TransparentUpgradeableProxyBin), backend, _logic, _admin, _incognito, _data)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TransparentUpgradeableProxy{TransparentUpgradeableProxyCaller: TransparentUpgradeableProxyCaller{contract: contract}, TransparentUpgradeableProxyTransactor: TransparentUpgradeableProxyTransactor{contract: contract}, TransparentUpgradeableProxyFilterer: TransparentUpgradeableProxyFilterer{contract: contract}}, nil
}

// TransparentUpgradeableProxy is an auto generated Go binding around an Ethereum contract.
type TransparentUpgradeableProxy struct {
	TransparentUpgradeableProxyCaller     // Read-only binding to the contract
	TransparentUpgradeableProxyTransactor // Write-only binding to the contract
	TransparentUpgradeableProxyFilterer   // Log filterer for contract events
}

// TransparentUpgradeableProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransparentUpgradeableProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransparentUpgradeableProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransparentUpgradeableProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransparentUpgradeableProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransparentUpgradeableProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransparentUpgradeableProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransparentUpgradeableProxySession struct {
	Contract     *TransparentUpgradeableProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// TransparentUpgradeableProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransparentUpgradeableProxyCallerSession struct {
	Contract *TransparentUpgradeableProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// TransparentUpgradeableProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransparentUpgradeableProxyTransactorSession struct {
	Contract     *TransparentUpgradeableProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// TransparentUpgradeableProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransparentUpgradeableProxyRaw struct {
	Contract *TransparentUpgradeableProxy // Generic contract binding to access the raw methods on
}

// TransparentUpgradeableProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransparentUpgradeableProxyCallerRaw struct {
	Contract *TransparentUpgradeableProxyCaller // Generic read-only contract binding to access the raw methods on
}

// TransparentUpgradeableProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransparentUpgradeableProxyTransactorRaw struct {
	Contract *TransparentUpgradeableProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransparentUpgradeableProxy creates a new instance of TransparentUpgradeableProxy, bound to a specific deployed contract.
func NewTransparentUpgradeableProxy(address common.Address, backend bind.ContractBackend) (*TransparentUpgradeableProxy, error) {
	contract, err := bindTransparentUpgradeableProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TransparentUpgradeableProxy{TransparentUpgradeableProxyCaller: TransparentUpgradeableProxyCaller{contract: contract}, TransparentUpgradeableProxyTransactor: TransparentUpgradeableProxyTransactor{contract: contract}, TransparentUpgradeableProxyFilterer: TransparentUpgradeableProxyFilterer{contract: contract}}, nil
}

// NewTransparentUpgradeableProxyCaller creates a new read-only instance of TransparentUpgradeableProxy, bound to a specific deployed contract.
func NewTransparentUpgradeableProxyCaller(address common.Address, caller bind.ContractCaller) (*TransparentUpgradeableProxyCaller, error) {
	contract, err := bindTransparentUpgradeableProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransparentUpgradeableProxyCaller{contract: contract}, nil
}

// NewTransparentUpgradeableProxyTransactor creates a new write-only instance of TransparentUpgradeableProxy, bound to a specific deployed contract.
func NewTransparentUpgradeableProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*TransparentUpgradeableProxyTransactor, error) {
	contract, err := bindTransparentUpgradeableProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransparentUpgradeableProxyTransactor{contract: contract}, nil
}

// NewTransparentUpgradeableProxyFilterer creates a new log filterer instance of TransparentUpgradeableProxy, bound to a specific deployed contract.
func NewTransparentUpgradeableProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*TransparentUpgradeableProxyFilterer, error) {
	contract, err := bindTransparentUpgradeableProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransparentUpgradeableProxyFilterer{contract: contract}, nil
}

// bindTransparentUpgradeableProxy binds a generic wrapper to an already deployed contract.
func bindTransparentUpgradeableProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TransparentUpgradeableProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransparentUpgradeableProxy.Contract.TransparentUpgradeableProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.TransparentUpgradeableProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.TransparentUpgradeableProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransparentUpgradeableProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.contract.Transact(opts, method, params...)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactor) Admin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.contract.Transact(opts, "admin")
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxySession) Admin() (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.Admin(&_TransparentUpgradeableProxy.TransactOpts)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactorSession) Admin() (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.Admin(&_TransparentUpgradeableProxy.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxySession) Claim() (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.Claim(&_TransparentUpgradeableProxy.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactorSession) Claim() (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.Claim(&_TransparentUpgradeableProxy.TransactOpts)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactor) Implementation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.contract.Transact(opts, "implementation")
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxySession) Implementation() (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.Implementation(&_TransparentUpgradeableProxy.TransactOpts)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactorSession) Implementation() (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.Implementation(&_TransparentUpgradeableProxy.TransactOpts)
}

// Incognito is a paid mutator transaction binding the contract method 0x8a984538.
//
// Solidity: function incognito() returns(address)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactor) Incognito(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.contract.Transact(opts, "incognito")
}

// Incognito is a paid mutator transaction binding the contract method 0x8a984538.
//
// Solidity: function incognito() returns(address)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxySession) Incognito() (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.Incognito(&_TransparentUpgradeableProxy.TransactOpts)
}

// Incognito is a paid mutator transaction binding the contract method 0x8a984538.
//
// Solidity: function incognito() returns(address)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactorSession) Incognito() (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.Incognito(&_TransparentUpgradeableProxy.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxySession) Pause() (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.Pause(&_TransparentUpgradeableProxy.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactorSession) Pause() (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.Pause(&_TransparentUpgradeableProxy.TransactOpts)
}

// Paused is a paid mutator transaction binding the contract method 0x5c975abb.
//
// Solidity: function paused() returns(bool)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactor) Paused(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.contract.Transact(opts, "paused")
}

// Paused is a paid mutator transaction binding the contract method 0x5c975abb.
//
// Solidity: function paused() returns(bool)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxySession) Paused() (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.Paused(&_TransparentUpgradeableProxy.TransactOpts)
}

// Paused is a paid mutator transaction binding the contract method 0x5c975abb.
//
// Solidity: function paused() returns(bool)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactorSession) Paused() (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.Paused(&_TransparentUpgradeableProxy.TransactOpts)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address newSuccessor) returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactor) Retire(opts *bind.TransactOpts, newSuccessor common.Address) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.contract.Transact(opts, "retire", newSuccessor)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address newSuccessor) returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxySession) Retire(newSuccessor common.Address) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.Retire(&_TransparentUpgradeableProxy.TransactOpts, newSuccessor)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address newSuccessor) returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactorSession) Retire(newSuccessor common.Address) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.Retire(&_TransparentUpgradeableProxy.TransactOpts, newSuccessor)
}

// Successor is a paid mutator transaction binding the contract method 0x6ff968c3.
//
// Solidity: function successor() returns(address)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactor) Successor(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.contract.Transact(opts, "successor")
}

// Successor is a paid mutator transaction binding the contract method 0x6ff968c3.
//
// Solidity: function successor() returns(address)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxySession) Successor() (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.Successor(&_TransparentUpgradeableProxy.TransactOpts)
}

// Successor is a paid mutator transaction binding the contract method 0x6ff968c3.
//
// Solidity: function successor() returns(address)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactorSession) Successor() (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.Successor(&_TransparentUpgradeableProxy.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxySession) Unpause() (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.Unpause(&_TransparentUpgradeableProxy.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactorSession) Unpause() (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.Unpause(&_TransparentUpgradeableProxy.TransactOpts)
}

// UpgradeIncognito is a paid mutator transaction binding the contract method 0x1c587771.
//
// Solidity: function upgradeIncognito(address newIncognito) returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactor) UpgradeIncognito(opts *bind.TransactOpts, newIncognito common.Address) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.contract.Transact(opts, "upgradeIncognito", newIncognito)
}

// UpgradeIncognito is a paid mutator transaction binding the contract method 0x1c587771.
//
// Solidity: function upgradeIncognito(address newIncognito) returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxySession) UpgradeIncognito(newIncognito common.Address) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.UpgradeIncognito(&_TransparentUpgradeableProxy.TransactOpts, newIncognito)
}

// UpgradeIncognito is a paid mutator transaction binding the contract method 0x1c587771.
//
// Solidity: function upgradeIncognito(address newIncognito) returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactorSession) UpgradeIncognito(newIncognito common.Address) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.UpgradeIncognito(&_TransparentUpgradeableProxy.TransactOpts, newIncognito)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxySession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.UpgradeTo(&_TransparentUpgradeableProxy.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.UpgradeTo(&_TransparentUpgradeableProxy.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.UpgradeToAndCall(&_TransparentUpgradeableProxy.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.UpgradeToAndCall(&_TransparentUpgradeableProxy.TransactOpts, newImplementation, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.Fallback(&_TransparentUpgradeableProxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.Fallback(&_TransparentUpgradeableProxy.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxySession) Receive() (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.Receive(&_TransparentUpgradeableProxy.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactorSession) Receive() (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.Receive(&_TransparentUpgradeableProxy.TransactOpts)
}

// TransparentUpgradeableProxyClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the TransparentUpgradeableProxy contract.
type TransparentUpgradeableProxyClaimIterator struct {
	Event *TransparentUpgradeableProxyClaim // Event containing the contract specifics and raw log

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
func (it *TransparentUpgradeableProxyClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransparentUpgradeableProxyClaim)
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
		it.Event = new(TransparentUpgradeableProxyClaim)
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
func (it *TransparentUpgradeableProxyClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransparentUpgradeableProxyClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransparentUpgradeableProxyClaim represents a Claim event raised by the TransparentUpgradeableProxy contract.
type TransparentUpgradeableProxyClaim struct {
	Claimer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) FilterClaim(opts *bind.FilterOpts) (*TransparentUpgradeableProxyClaimIterator, error) {

	logs, sub, err := _TransparentUpgradeableProxy.contract.FilterLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return &TransparentUpgradeableProxyClaimIterator{contract: _TransparentUpgradeableProxy.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *TransparentUpgradeableProxyClaim) (event.Subscription, error) {

	logs, sub, err := _TransparentUpgradeableProxy.contract.WatchLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransparentUpgradeableProxyClaim)
				if err := _TransparentUpgradeableProxy.contract.UnpackLog(event, "Claim", log); err != nil {
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
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) ParseClaim(log types.Log) (*TransparentUpgradeableProxyClaim, error) {
	event := new(TransparentUpgradeableProxyClaim)
	if err := _TransparentUpgradeableProxy.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransparentUpgradeableProxyIncognitoChangedIterator is returned from FilterIncognitoChanged and is used to iterate over the raw logs and unpacked data for IncognitoChanged events raised by the TransparentUpgradeableProxy contract.
type TransparentUpgradeableProxyIncognitoChangedIterator struct {
	Event *TransparentUpgradeableProxyIncognitoChanged // Event containing the contract specifics and raw log

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
func (it *TransparentUpgradeableProxyIncognitoChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransparentUpgradeableProxyIncognitoChanged)
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
		it.Event = new(TransparentUpgradeableProxyIncognitoChanged)
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
func (it *TransparentUpgradeableProxyIncognitoChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransparentUpgradeableProxyIncognitoChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransparentUpgradeableProxyIncognitoChanged represents a IncognitoChanged event raised by the TransparentUpgradeableProxy contract.
type TransparentUpgradeableProxyIncognitoChanged struct {
	PreviousIncognito common.Address
	NewIncognito      common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterIncognitoChanged is a free log retrieval operation binding the contract event 0x86d392a76e88298144124db3dd7265135d76810f52d747dc329a0f7722135e5c.
//
// Solidity: event IncognitoChanged(address previousIncognito, address newIncognito)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) FilterIncognitoChanged(opts *bind.FilterOpts) (*TransparentUpgradeableProxyIncognitoChangedIterator, error) {

	logs, sub, err := _TransparentUpgradeableProxy.contract.FilterLogs(opts, "IncognitoChanged")
	if err != nil {
		return nil, err
	}
	return &TransparentUpgradeableProxyIncognitoChangedIterator{contract: _TransparentUpgradeableProxy.contract, event: "IncognitoChanged", logs: logs, sub: sub}, nil
}

// WatchIncognitoChanged is a free log subscription operation binding the contract event 0x86d392a76e88298144124db3dd7265135d76810f52d747dc329a0f7722135e5c.
//
// Solidity: event IncognitoChanged(address previousIncognito, address newIncognito)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) WatchIncognitoChanged(opts *bind.WatchOpts, sink chan<- *TransparentUpgradeableProxyIncognitoChanged) (event.Subscription, error) {

	logs, sub, err := _TransparentUpgradeableProxy.contract.WatchLogs(opts, "IncognitoChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransparentUpgradeableProxyIncognitoChanged)
				if err := _TransparentUpgradeableProxy.contract.UnpackLog(event, "IncognitoChanged", log); err != nil {
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

// ParseIncognitoChanged is a log parse operation binding the contract event 0x86d392a76e88298144124db3dd7265135d76810f52d747dc329a0f7722135e5c.
//
// Solidity: event IncognitoChanged(address previousIncognito, address newIncognito)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) ParseIncognitoChanged(log types.Log) (*TransparentUpgradeableProxyIncognitoChanged, error) {
	event := new(TransparentUpgradeableProxyIncognitoChanged)
	if err := _TransparentUpgradeableProxy.contract.UnpackLog(event, "IncognitoChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransparentUpgradeableProxyPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the TransparentUpgradeableProxy contract.
type TransparentUpgradeableProxyPausedIterator struct {
	Event *TransparentUpgradeableProxyPaused // Event containing the contract specifics and raw log

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
func (it *TransparentUpgradeableProxyPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransparentUpgradeableProxyPaused)
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
		it.Event = new(TransparentUpgradeableProxyPaused)
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
func (it *TransparentUpgradeableProxyPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransparentUpgradeableProxyPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransparentUpgradeableProxyPaused represents a Paused event raised by the TransparentUpgradeableProxy contract.
type TransparentUpgradeableProxyPaused struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address admin)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) FilterPaused(opts *bind.FilterOpts) (*TransparentUpgradeableProxyPausedIterator, error) {

	logs, sub, err := _TransparentUpgradeableProxy.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &TransparentUpgradeableProxyPausedIterator{contract: _TransparentUpgradeableProxy.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address admin)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *TransparentUpgradeableProxyPaused) (event.Subscription, error) {

	logs, sub, err := _TransparentUpgradeableProxy.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransparentUpgradeableProxyPaused)
				if err := _TransparentUpgradeableProxy.contract.UnpackLog(event, "Paused", log); err != nil {
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
// Solidity: event Paused(address admin)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) ParsePaused(log types.Log) (*TransparentUpgradeableProxyPaused, error) {
	event := new(TransparentUpgradeableProxyPaused)
	if err := _TransparentUpgradeableProxy.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransparentUpgradeableProxySuccessorChangedIterator is returned from FilterSuccessorChanged and is used to iterate over the raw logs and unpacked data for SuccessorChanged events raised by the TransparentUpgradeableProxy contract.
type TransparentUpgradeableProxySuccessorChangedIterator struct {
	Event *TransparentUpgradeableProxySuccessorChanged // Event containing the contract specifics and raw log

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
func (it *TransparentUpgradeableProxySuccessorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransparentUpgradeableProxySuccessorChanged)
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
		it.Event = new(TransparentUpgradeableProxySuccessorChanged)
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
func (it *TransparentUpgradeableProxySuccessorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransparentUpgradeableProxySuccessorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransparentUpgradeableProxySuccessorChanged represents a SuccessorChanged event raised by the TransparentUpgradeableProxy contract.
type TransparentUpgradeableProxySuccessorChanged struct {
	PreviousSuccessor common.Address
	NewSuccessor      common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSuccessorChanged is a free log retrieval operation binding the contract event 0xf966f857c3c376f2e1df873bbe2596a18675dc056dc3465dfbbe8fe9ac02c974.
//
// Solidity: event SuccessorChanged(address previousSuccessor, address newSuccessor)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) FilterSuccessorChanged(opts *bind.FilterOpts) (*TransparentUpgradeableProxySuccessorChangedIterator, error) {

	logs, sub, err := _TransparentUpgradeableProxy.contract.FilterLogs(opts, "SuccessorChanged")
	if err != nil {
		return nil, err
	}
	return &TransparentUpgradeableProxySuccessorChangedIterator{contract: _TransparentUpgradeableProxy.contract, event: "SuccessorChanged", logs: logs, sub: sub}, nil
}

// WatchSuccessorChanged is a free log subscription operation binding the contract event 0xf966f857c3c376f2e1df873bbe2596a18675dc056dc3465dfbbe8fe9ac02c974.
//
// Solidity: event SuccessorChanged(address previousSuccessor, address newSuccessor)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) WatchSuccessorChanged(opts *bind.WatchOpts, sink chan<- *TransparentUpgradeableProxySuccessorChanged) (event.Subscription, error) {

	logs, sub, err := _TransparentUpgradeableProxy.contract.WatchLogs(opts, "SuccessorChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransparentUpgradeableProxySuccessorChanged)
				if err := _TransparentUpgradeableProxy.contract.UnpackLog(event, "SuccessorChanged", log); err != nil {
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
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) ParseSuccessorChanged(log types.Log) (*TransparentUpgradeableProxySuccessorChanged, error) {
	event := new(TransparentUpgradeableProxySuccessorChanged)
	if err := _TransparentUpgradeableProxy.contract.UnpackLog(event, "SuccessorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransparentUpgradeableProxyUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the TransparentUpgradeableProxy contract.
type TransparentUpgradeableProxyUnpausedIterator struct {
	Event *TransparentUpgradeableProxyUnpaused // Event containing the contract specifics and raw log

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
func (it *TransparentUpgradeableProxyUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransparentUpgradeableProxyUnpaused)
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
		it.Event = new(TransparentUpgradeableProxyUnpaused)
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
func (it *TransparentUpgradeableProxyUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransparentUpgradeableProxyUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransparentUpgradeableProxyUnpaused represents a Unpaused event raised by the TransparentUpgradeableProxy contract.
type TransparentUpgradeableProxyUnpaused struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address admin)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) FilterUnpaused(opts *bind.FilterOpts) (*TransparentUpgradeableProxyUnpausedIterator, error) {

	logs, sub, err := _TransparentUpgradeableProxy.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &TransparentUpgradeableProxyUnpausedIterator{contract: _TransparentUpgradeableProxy.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address admin)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *TransparentUpgradeableProxyUnpaused) (event.Subscription, error) {

	logs, sub, err := _TransparentUpgradeableProxy.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransparentUpgradeableProxyUnpaused)
				if err := _TransparentUpgradeableProxy.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
// Solidity: event Unpaused(address admin)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) ParseUnpaused(log types.Log) (*TransparentUpgradeableProxyUnpaused, error) {
	event := new(TransparentUpgradeableProxyUnpaused)
	if err := _TransparentUpgradeableProxy.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransparentUpgradeableProxyUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the TransparentUpgradeableProxy contract.
type TransparentUpgradeableProxyUpgradedIterator struct {
	Event *TransparentUpgradeableProxyUpgraded // Event containing the contract specifics and raw log

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
func (it *TransparentUpgradeableProxyUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransparentUpgradeableProxyUpgraded)
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
		it.Event = new(TransparentUpgradeableProxyUpgraded)
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
func (it *TransparentUpgradeableProxyUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransparentUpgradeableProxyUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransparentUpgradeableProxyUpgraded represents a Upgraded event raised by the TransparentUpgradeableProxy contract.
type TransparentUpgradeableProxyUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TransparentUpgradeableProxyUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TransparentUpgradeableProxy.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TransparentUpgradeableProxyUpgradedIterator{contract: _TransparentUpgradeableProxy.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TransparentUpgradeableProxyUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TransparentUpgradeableProxy.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransparentUpgradeableProxyUpgraded)
				if err := _TransparentUpgradeableProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) ParseUpgraded(log types.Log) (*TransparentUpgradeableProxyUpgraded, error) {
	event := new(TransparentUpgradeableProxyUpgraded)
	if err := _TransparentUpgradeableProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeableProxyABI is the input ABI used to generate the binding from.
const UpgradeableProxyABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// UpgradeableProxyBin is the compiled bytecode used for deploying new contracts.
var UpgradeableProxyBin = "0x608060405260405161034a38038061034a8339818101604052604081101561002657600080fd5b81516020830180516040519294929383019291908464010000000082111561004d57600080fd5b90830190602082018581111561006257600080fd5b825164010000000081118282018810171561007c57600080fd5b82525081516020918201929091019080838360005b838110156100a9578181015183820152602001610091565b50505050905090810190601f1680156100d65780820380516001836020036101000a031916815260200191505b50604052506100e3915050565b6100ec826101f3565b8051156101ec576000826001600160a01b0316826040518082805190602001908083835b6020831061012f5780518252601f199092019160209182019101610110565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d806000811461018f576040519150601f19603f3d011682016040523d82523d6000602084013e610194565b606091505b50509050806101ea576040805162461bcd60e51b815260206004820152601360248201527f44454c454741544543414c4c206661696c656400000000000000000000000000604482015290519081900360640190fd5b505b5050610261565b6101fc8161025b565b6102375760405162461bcd60e51b81526004018080602001828103825260368152602001806103146036913960400191505060405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc55565b3b151590565b60a58061026f6000396000f3fe608060405236601057600e6013565b005b600e5b60196025565b602560216027565b604c565b565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5490565b3660008037600080366000845af43d6000803e808015606a573d6000f35b3d6000fdfea26469706673582212204bd40f776036c57969bd612afceddfe75b0ecaa871c34eab82d7265ff9740ab364736f6c634300060c00335570677261646561626c6550726f78793a206e657720696d706c656d656e746174696f6e206973206e6f74206120636f6e7472616374"

// DeployUpgradeableProxy deploys a new Ethereum contract, binding an instance of UpgradeableProxy to it.
func DeployUpgradeableProxy(auth *bind.TransactOpts, backend bind.ContractBackend, _logic common.Address, _data []byte) (common.Address, *types.Transaction, *UpgradeableProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(UpgradeableProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UpgradeableProxyBin), backend, _logic, _data)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UpgradeableProxy{UpgradeableProxyCaller: UpgradeableProxyCaller{contract: contract}, UpgradeableProxyTransactor: UpgradeableProxyTransactor{contract: contract}, UpgradeableProxyFilterer: UpgradeableProxyFilterer{contract: contract}}, nil
}

// UpgradeableProxy is an auto generated Go binding around an Ethereum contract.
type UpgradeableProxy struct {
	UpgradeableProxyCaller     // Read-only binding to the contract
	UpgradeableProxyTransactor // Write-only binding to the contract
	UpgradeableProxyFilterer   // Log filterer for contract events
}

// UpgradeableProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type UpgradeableProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeableProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UpgradeableProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeableProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UpgradeableProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeableProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UpgradeableProxySession struct {
	Contract     *UpgradeableProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UpgradeableProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UpgradeableProxyCallerSession struct {
	Contract *UpgradeableProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// UpgradeableProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UpgradeableProxyTransactorSession struct {
	Contract     *UpgradeableProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// UpgradeableProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type UpgradeableProxyRaw struct {
	Contract *UpgradeableProxy // Generic contract binding to access the raw methods on
}

// UpgradeableProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UpgradeableProxyCallerRaw struct {
	Contract *UpgradeableProxyCaller // Generic read-only contract binding to access the raw methods on
}

// UpgradeableProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UpgradeableProxyTransactorRaw struct {
	Contract *UpgradeableProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUpgradeableProxy creates a new instance of UpgradeableProxy, bound to a specific deployed contract.
func NewUpgradeableProxy(address common.Address, backend bind.ContractBackend) (*UpgradeableProxy, error) {
	contract, err := bindUpgradeableProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UpgradeableProxy{UpgradeableProxyCaller: UpgradeableProxyCaller{contract: contract}, UpgradeableProxyTransactor: UpgradeableProxyTransactor{contract: contract}, UpgradeableProxyFilterer: UpgradeableProxyFilterer{contract: contract}}, nil
}

// NewUpgradeableProxyCaller creates a new read-only instance of UpgradeableProxy, bound to a specific deployed contract.
func NewUpgradeableProxyCaller(address common.Address, caller bind.ContractCaller) (*UpgradeableProxyCaller, error) {
	contract, err := bindUpgradeableProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeableProxyCaller{contract: contract}, nil
}

// NewUpgradeableProxyTransactor creates a new write-only instance of UpgradeableProxy, bound to a specific deployed contract.
func NewUpgradeableProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*UpgradeableProxyTransactor, error) {
	contract, err := bindUpgradeableProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeableProxyTransactor{contract: contract}, nil
}

// NewUpgradeableProxyFilterer creates a new log filterer instance of UpgradeableProxy, bound to a specific deployed contract.
func NewUpgradeableProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*UpgradeableProxyFilterer, error) {
	contract, err := bindUpgradeableProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UpgradeableProxyFilterer{contract: contract}, nil
}

// bindUpgradeableProxy binds a generic wrapper to an already deployed contract.
func bindUpgradeableProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UpgradeableProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradeableProxy *UpgradeableProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpgradeableProxy.Contract.UpgradeableProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradeableProxy *UpgradeableProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeableProxy.Contract.UpgradeableProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradeableProxy *UpgradeableProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradeableProxy.Contract.UpgradeableProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradeableProxy *UpgradeableProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpgradeableProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradeableProxy *UpgradeableProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeableProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradeableProxy *UpgradeableProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradeableProxy.Contract.contract.Transact(opts, method, params...)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_UpgradeableProxy *UpgradeableProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _UpgradeableProxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_UpgradeableProxy *UpgradeableProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _UpgradeableProxy.Contract.Fallback(&_UpgradeableProxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_UpgradeableProxy *UpgradeableProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _UpgradeableProxy.Contract.Fallback(&_UpgradeableProxy.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UpgradeableProxy *UpgradeableProxyTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeableProxy.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UpgradeableProxy *UpgradeableProxySession) Receive() (*types.Transaction, error) {
	return _UpgradeableProxy.Contract.Receive(&_UpgradeableProxy.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UpgradeableProxy *UpgradeableProxyTransactorSession) Receive() (*types.Transaction, error) {
	return _UpgradeableProxy.Contract.Receive(&_UpgradeableProxy.TransactOpts)
}

// UpgradeableProxyUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the UpgradeableProxy contract.
type UpgradeableProxyUpgradedIterator struct {
	Event *UpgradeableProxyUpgraded // Event containing the contract specifics and raw log

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
func (it *UpgradeableProxyUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeableProxyUpgraded)
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
		it.Event = new(UpgradeableProxyUpgraded)
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
func (it *UpgradeableProxyUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeableProxyUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeableProxyUpgraded represents a Upgraded event raised by the UpgradeableProxy contract.
type UpgradeableProxyUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_UpgradeableProxy *UpgradeableProxyFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*UpgradeableProxyUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _UpgradeableProxy.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeableProxyUpgradedIterator{contract: _UpgradeableProxy.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_UpgradeableProxy *UpgradeableProxyFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *UpgradeableProxyUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _UpgradeableProxy.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeableProxyUpgraded)
				if err := _UpgradeableProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_UpgradeableProxy *UpgradeableProxyFilterer) ParseUpgraded(log types.Log) (*UpgradeableProxyUpgraded, error) {
	event := new(UpgradeableProxyUpgraded)
	if err := _UpgradeableProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
