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
const VaultproxyABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_incognito\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousIncognito\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newIncognito\",\"type\":\"address\"}],\"name\":\"IncognitoChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousSuccessor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newSuccessor\",\"type\":\"address\"}],\"name\":\"SuccessorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incognito\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newSuccessor\",\"type\":\"address\"}],\"name\":\"retire\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"successor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newIncognito\",\"type\":\"address\"}],\"name\":\"upgradeIncognito\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// VaultproxyBin is the compiled bytecode used for deploying new contracts.
var VaultproxyBin = "0x6080604052604051620016d1380380620016d1833981810160405260808110156200002957600080fd5b81019080805190602001909291908051906020019092919080519060200190929190805160405193929190846401000000008211156200006857600080fd5b838201915060208201858111156200007f57600080fd5b82518660018202830111640100000000821117156200009d57600080fd5b8083526020830192505050908051906020019080838360005b83811015620000d3578082015181840152602081019050620000b6565b50505050905090810190601f168015620001015780820380516001836020036101000a031916815260200191505b50604052505050838160017f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbd60001c0360001b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b146200016057fe5b62000171826200043260201b60201c565b600081511115620002ac5760008273ffffffffffffffffffffffffffffffffffffffff16826040518082805190602001908083835b60208310620001cb5780518252602082019150602081019050602083039250620001a6565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d80600081146200022d576040519150601f19603f3d011682016040523d82523d6000602084013e62000232565b606091505b5050905080620002aa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f44454c454741544543414c4c206661696c65640000000000000000000000000081525060200191505060405180910390fd5b505b505060017fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610460001c0360001b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610360001b146200030457fe5b60017f7b13fc932b1063ca775d428558b73e20eab6804d4d9b5a148d7cbae4488973f960001c0360001b7f7b13fc932b1063ca775d428558b73e20eab6804d4d9b5a148d7cbae4488973f860001b146200035a57fe5b60017f8dea8703c3cf94703383ce38a9c894669dccd4ca8e65ddb43267aa024871145160001c0360001b7f8dea8703c3cf94703383ce38a9c894669dccd4ca8e65ddb43267aa024871145060001b14620003b057fe5b60017f62135fc083646fdb4e1a9d700e351b886a4a5a39da980650269edd1ade91ffd360001c0360001b7f62135fc083646fdb4e1a9d700e351b886a4a5a39da980650269edd1ade91ffd260001b146200040657fe5b6200041783620004c960201b60201c565b6200042882620004f860201b60201c565b505050506200053a565b62000443816200052760201b60201c565b6200049a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260368152602001806200169b6036913960400191505060405180910390fd5b60007f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b90508181555050565b60007fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610360001b90508181555050565b60007f62135fc083646fdb4e1a9d700e351b886a4a5a39da980650269edd1ade91ffd260001b90508181555050565b600080823b905060008111915050919050565b611151806200054a6000396000f3fe6080604052600436106100ab5760003560e01c80635c975abb116100645780635c975abb1461026e5780636ff968c31461029b5780638456cb59146102dc5780638a984538146102f35780639e6371ba14610334578063f851a44014610385576100ba565b80631c587771146100c45780633659cfe6146101155780633f4ba83a146101665780634e71d92d1461017d5780634f1ef286146101945780635c60da1b1461022d576100ba565b366100ba576100b86103c6565b005b6100c26103c6565b005b3480156100d057600080fd5b50610113600480360360208110156100e757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506103e0565b005b34801561012157600080fd5b506101646004803603602081101561013857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061052d565b005b34801561017257600080fd5b5061017b610582565b005b34801561018957600080fd5b50610192610634565b005b61022b600480360360408110156101aa57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001906401000000008111156101e757600080fd5b8201836020820111156101f957600080fd5b8035906020019184600183028401116401000000008311171561021b57600080fd5b90919293919293905050506106e3565b005b34801561023957600080fd5b50610242610822565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561027a57600080fd5b5061028361087a565b60405180821515815260200191505060405180910390f35b3480156102a757600080fd5b506102b06108d2565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156102e857600080fd5b506102f161092a565b005b3480156102ff57600080fd5b506103086109dd565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561034057600080fd5b506103836004803603602081101561035757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610a35565b005b34801561039157600080fd5b5061039a610b82565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6103ce610bda565b6103de6103d9610cce565b610cff565b565b6103e8610d25565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561052157600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156104a1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260408152602001806110076040913960400191505060405180910390fd5b7f86d392a76e88298144124db3dd7265135d76810f52d747dc329a0f7722135e5c6104ca610d56565b82604051808373ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a161051c81610d87565b61052a565b6105296103c6565b5b50565b610535610d25565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156105765761057181610db6565b61057f565b61057e6103c6565b5b50565b61058a610d25565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415610629576105c5610e05565b61061a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001806110aa6030913960400191505060405180910390fd5b6106246000610e36565b610632565b6106316103c6565b5b565b61063c610e65565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156106d8577f0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc610698610e65565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a16106d36106ce610e65565b610e96565b6106e1565b6106e06103c6565b5b565b6106eb610d25565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156108145761072783610db6565b60008373ffffffffffffffffffffffffffffffffffffffff168383604051808383808284378083019250505092505050600060405180830381855af49150503d8060008114610792576040519150601f19603f3d011682016040523d82523d6000602084013e610797565b606091505b505090508061080e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f44454c454741544543414c4c206661696c65640000000000000000000000000081525060200191505060405180910390fd5b5061081d565b61081c6103c6565b5b505050565b600061082c610d25565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561086e57610867610cce565b9050610877565b6108766103c6565b5b90565b6000610884610d25565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156108c6576108bf610e05565b90506108cf565b6108ce6103c6565b5b90565b60006108dc610d25565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561091e57610917610e65565b9050610927565b6109266103c6565b5b90565b610932610d25565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156109d25761096d610e05565b156109c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260348152602001806110476034913960400191505060405180910390fd5b6109cd6001610e36565b6109db565b6109da6103c6565b5b565b60006109e7610d25565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415610a2957610a22610d56565b9050610a32565b610a316103c6565b5b90565b610a3d610d25565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415610b7657600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610af6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603a815260200180610f97603a913960400191505060405180910390fd5b7ff966f857c3c376f2e1df873bbe2596a18675dc056dc3465dfbbe8fe9ac02c974610b1f610e65565b82604051808373ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a1610b7181610ec5565b610b7f565b610b7e6103c6565b5b50565b6000610b8c610d25565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415610bce57610bc7610d25565b9050610bd7565b610bd66103c6565b5b90565b610be2610d25565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415610c66576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260428152602001806110da6042913960600191505060405180910390fd5b610c6e610e05565b15610cc4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f81526020018061107b602f913960400191505060405180910390fd5b610ccc610ef4565b565b6000807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b9050805491505090565b3660008037600080366000845af43d6000803e8060008114610d20573d6000f35b3d6000fd5b6000807fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610360001b9050805491505090565b6000807f62135fc083646fdb4e1a9d700e351b886a4a5a39da980650269edd1ade91ffd260001b9050805491505090565b60007f62135fc083646fdb4e1a9d700e351b886a4a5a39da980650269edd1ade91ffd260001b90508181555050565b610dbf81610ef6565b8073ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a250565b6000807f8dea8703c3cf94703383ce38a9c894669dccd4ca8e65ddb43267aa024871145060001b9050805491505090565b60007f8dea8703c3cf94703383ce38a9c894669dccd4ca8e65ddb43267aa024871145060001b90508181555050565b6000807f7b13fc932b1063ca775d428558b73e20eab6804d4d9b5a148d7cbae4488973f860001b9050805491505090565b60007fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610360001b90508181555050565b60007f7b13fc932b1063ca775d428558b73e20eab6804d4d9b5a148d7cbae4488973f860001b90508181555050565b565b610eff81610f83565b610f54576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526036815260200180610fd16036913960400191505060405180910390fd5b60007f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b90508181555050565b600080823b90506000811191505091905056fe5472616e73706172656e745570677261646561626c6550726f78793a20737563636573736f7220697320746865207a65726f20616464726573735570677261646561626c6550726f78793a206e657720696d706c656d656e746174696f6e206973206e6f74206120636f6e74726163745472616e73706172656e745570677261646561626c6550726f78793a20696e636f676e69746f2070726f787920697320746865207a65726f20616464726573735472616e73706172656e745570677261646561626c6550726f78793a20636f6e74726163742070617573656420616c72656164795472616e73706172656e745570677261646561626c6550726f78793a20636f6e7472616374206973207061757365645472616e73706172656e745570677261646561626c6550726f78793a20636f6e7472616374206e6f74207061757365645472616e73706172656e745570677261646561626c6550726f78793a2061646d696e2063616e6e6f742066616c6c6261636b20746f2070726f787920746172676574a2646970667358221220636d47e7b862098a3c6dd14eaea292252692bd2db499654e35cd2907543f533a64736f6c634300060c00335570677261646561626c6550726f78793a206e657720696d706c656d656e746174696f6e206973206e6f74206120636f6e7472616374"

// DeployVaultproxy deploys a new Ethereum contract, binding an instance of Vaultproxy to it.
func DeployVaultproxy(auth *bind.TransactOpts, backend bind.ContractBackend, _logic common.Address, _admin common.Address, _incognito common.Address, _data []byte) (common.Address, *types.Transaction, *Vaultproxy, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultproxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VaultproxyBin), backend, _logic, _admin, _incognito, _data)
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

// Incognito is a paid mutator transaction binding the contract method 0x8a984538.
//
// Solidity: function incognito() returns(address)
func (_Vaultproxy *VaultproxyTransactor) Incognito(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vaultproxy.contract.Transact(opts, "incognito")
}

// Incognito is a paid mutator transaction binding the contract method 0x8a984538.
//
// Solidity: function incognito() returns(address)
func (_Vaultproxy *VaultproxySession) Incognito() (*types.Transaction, error) {
	return _Vaultproxy.Contract.Incognito(&_Vaultproxy.TransactOpts)
}

// Incognito is a paid mutator transaction binding the contract method 0x8a984538.
//
// Solidity: function incognito() returns(address)
func (_Vaultproxy *VaultproxyTransactorSession) Incognito() (*types.Transaction, error) {
	return _Vaultproxy.Contract.Incognito(&_Vaultproxy.TransactOpts)
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

// Paused is a paid mutator transaction binding the contract method 0x5c975abb.
//
// Solidity: function paused() returns(bool)
func (_Vaultproxy *VaultproxyTransactor) Paused(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vaultproxy.contract.Transact(opts, "paused")
}

// Paused is a paid mutator transaction binding the contract method 0x5c975abb.
//
// Solidity: function paused() returns(bool)
func (_Vaultproxy *VaultproxySession) Paused() (*types.Transaction, error) {
	return _Vaultproxy.Contract.Paused(&_Vaultproxy.TransactOpts)
}

// Paused is a paid mutator transaction binding the contract method 0x5c975abb.
//
// Solidity: function paused() returns(bool)
func (_Vaultproxy *VaultproxyTransactorSession) Paused() (*types.Transaction, error) {
	return _Vaultproxy.Contract.Paused(&_Vaultproxy.TransactOpts)
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

// UpgradeIncognito is a paid mutator transaction binding the contract method 0x1c587771.
//
// Solidity: function upgradeIncognito(address newIncognito) returns()
func (_Vaultproxy *VaultproxyTransactor) UpgradeIncognito(opts *bind.TransactOpts, newIncognito common.Address) (*types.Transaction, error) {
	return _Vaultproxy.contract.Transact(opts, "upgradeIncognito", newIncognito)
}

// UpgradeIncognito is a paid mutator transaction binding the contract method 0x1c587771.
//
// Solidity: function upgradeIncognito(address newIncognito) returns()
func (_Vaultproxy *VaultproxySession) UpgradeIncognito(newIncognito common.Address) (*types.Transaction, error) {
	return _Vaultproxy.Contract.UpgradeIncognito(&_Vaultproxy.TransactOpts, newIncognito)
}

// UpgradeIncognito is a paid mutator transaction binding the contract method 0x1c587771.
//
// Solidity: function upgradeIncognito(address newIncognito) returns()
func (_Vaultproxy *VaultproxyTransactorSession) UpgradeIncognito(newIncognito common.Address) (*types.Transaction, error) {
	return _Vaultproxy.Contract.UpgradeIncognito(&_Vaultproxy.TransactOpts, newIncognito)
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

// VaultproxyIncognitoChangedIterator is returned from FilterIncognitoChanged and is used to iterate over the raw logs and unpacked data for IncognitoChanged events raised by the Vaultproxy contract.
type VaultproxyIncognitoChangedIterator struct {
	Event *VaultproxyIncognitoChanged // Event containing the contract specifics and raw log

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
func (it *VaultproxyIncognitoChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultproxyIncognitoChanged)
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
		it.Event = new(VaultproxyIncognitoChanged)
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
func (it *VaultproxyIncognitoChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultproxyIncognitoChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultproxyIncognitoChanged represents a IncognitoChanged event raised by the Vaultproxy contract.
type VaultproxyIncognitoChanged struct {
	PreviousIncognito common.Address
	NewIncognito      common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterIncognitoChanged is a free log retrieval operation binding the contract event 0x86d392a76e88298144124db3dd7265135d76810f52d747dc329a0f7722135e5c.
//
// Solidity: event IncognitoChanged(address previousIncognito, address newIncognito)
func (_Vaultproxy *VaultproxyFilterer) FilterIncognitoChanged(opts *bind.FilterOpts) (*VaultproxyIncognitoChangedIterator, error) {

	logs, sub, err := _Vaultproxy.contract.FilterLogs(opts, "IncognitoChanged")
	if err != nil {
		return nil, err
	}
	return &VaultproxyIncognitoChangedIterator{contract: _Vaultproxy.contract, event: "IncognitoChanged", logs: logs, sub: sub}, nil
}

// WatchIncognitoChanged is a free log subscription operation binding the contract event 0x86d392a76e88298144124db3dd7265135d76810f52d747dc329a0f7722135e5c.
//
// Solidity: event IncognitoChanged(address previousIncognito, address newIncognito)
func (_Vaultproxy *VaultproxyFilterer) WatchIncognitoChanged(opts *bind.WatchOpts, sink chan<- *VaultproxyIncognitoChanged) (event.Subscription, error) {

	logs, sub, err := _Vaultproxy.contract.WatchLogs(opts, "IncognitoChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultproxyIncognitoChanged)
				if err := _Vaultproxy.contract.UnpackLog(event, "IncognitoChanged", log); err != nil {
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
func (_Vaultproxy *VaultproxyFilterer) ParseIncognitoChanged(log types.Log) (*VaultproxyIncognitoChanged, error) {
	event := new(VaultproxyIncognitoChanged)
	if err := _Vaultproxy.contract.UnpackLog(event, "IncognitoChanged", log); err != nil {
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
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address admin)
func (_Vaultproxy *VaultproxyFilterer) FilterPaused(opts *bind.FilterOpts) (*VaultproxyPausedIterator, error) {

	logs, sub, err := _Vaultproxy.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &VaultproxyPausedIterator{contract: _Vaultproxy.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address admin)
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
// Solidity: event Paused(address admin)
func (_Vaultproxy *VaultproxyFilterer) ParsePaused(log types.Log) (*VaultproxyPaused, error) {
	event := new(VaultproxyPaused)
	if err := _Vaultproxy.contract.UnpackLog(event, "Paused", log); err != nil {
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
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address admin)
func (_Vaultproxy *VaultproxyFilterer) FilterUnpaused(opts *bind.FilterOpts) (*VaultproxyUnpausedIterator, error) {

	logs, sub, err := _Vaultproxy.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &VaultproxyUnpausedIterator{contract: _Vaultproxy.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address admin)
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
// Solidity: event Unpaused(address admin)
func (_Vaultproxy *VaultproxyFilterer) ParseUnpaused(log types.Log) (*VaultproxyUnpaused, error) {
	event := new(VaultproxyUnpaused)
	if err := _Vaultproxy.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
