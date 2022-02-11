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

// PcurveMetaData contains all meta data concerning the Pcurve contract.
var PcurveMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ETH_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETH_CURVE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"j\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"contractICurveSwap\",\"name\":\"curvePool\",\"type\":\"address\"}],\"name\":\"exchange\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"j\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"contractICurveSwap\",\"name\":\"curvePool\",\"type\":\"address\"}],\"name\":\"exchangeUnderlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052604051610d3d380380610d3d8339818101604052602081101561002657600080fd5b8101908080519060200190929190505050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610cb6806100876000396000f3fe6080604052600436106100595760003560e01c806372e94bf6146100655780638c0b6593146100a6578063a64833a014610143578063ad5c4648146101e0578063d49d518114610221578063fc0db3191461024c57610060565b3661006057005b600080fd5b34801561007157600080fd5b5061007a61028d565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610110600480360360a08110156100bc57600080fd5b8101908080359060200190929190803590602001909291908035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610292565b604051808373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b6101ad600480360360a081101561015957600080fd5b8101908080359060200190929190803590602001909291908035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506105cc565b604051808373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b3480156101ec57600080fd5b506101f5610906565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561022d57600080fd5b5061023661092a565b6040518082815260200191505060405180910390f35b34801561025857600080fd5b5061026161094e565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b600081565b6000806000851161030b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f696e76616c6964207377617020616d6f756e740000000000000000000000000081525060200191505060405180910390fd5b60008373ffffffffffffffffffffffffffffffffffffffff1663b9947eb0896040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b15801561036057600080fd5b505af1158015610374573d6000803e3d6000fd5b505050506040513d602081101561038a57600080fd5b8101908080519060200190929190505050905060008473ffffffffffffffffffffffffffffffffffffffff1663c6610657896040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b1580156103f257600080fd5b505af1158015610406573d6000803e3d6000fd5b505050506040513d602081101561041c57600080fd5b81019080805190602001909291905050509050600073eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161461047e5781610481565b60005b905061048e838988610966565b60008673ffffffffffffffffffffffffffffffffffffffff166365b2489b8c8c8c8c6040518563ffffffff1660e01b815260040180858152602001848152602001838152602001828152602001945050505050602060405180830381600087803b1580156104fb57600080fd5b505af115801561050f573d6000803e3d6000fd5b505050506040513d602081101561052557600080fd5b81019080805190602001909291905050509050878110156105ae576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f4e6f7420656e6f75676820636f696e000000000000000000000000000000000081525060200191505060405180910390fd5b6105b88282610ae8565b818195509550505050509550959350505050565b60008060008511610645576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f696e76616c6964207377617020616d6f756e740000000000000000000000000081525060200191505060405180910390fd5b60008373ffffffffffffffffffffffffffffffffffffffff1663c6610657896040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b15801561069a57600080fd5b505af11580156106ae573d6000803e3d6000fd5b505050506040513d60208110156106c457600080fd5b8101908080519060200190929190505050905060008473ffffffffffffffffffffffffffffffffffffffff1663c6610657896040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b15801561072c57600080fd5b505af1158015610740573d6000803e3d6000fd5b505050506040513d602081101561075657600080fd5b81019080805190602001909291905050509050600073eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16146107b857816107bb565b60005b90506107c8838988610966565b60008673ffffffffffffffffffffffffffffffffffffffff16635b41b9088c8c8c8c6040518563ffffffff1660e01b815260040180858152602001848152602001838152602001828152602001945050505050602060405180830381600087803b15801561083557600080fd5b505af1158015610849573d6000803e3d6000fd5b505050506040513d602081101561085f57600080fd5b81019080805190602001909291905050509050878110156108e8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f4e6f7420656e6f75676820636f696e000000000000000000000000000000000081525060200191505060405180910390fd5b6108f28282610ae8565b818195509550505050509550959350505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee81565b600034148015610a345750818373ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e30846040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1681526020019250505060206040518083038186803b1580156109f757600080fd5b505afa158015610a0b573d6000803e3d6000fd5b505050506040513d6020811015610a2157600080fd5b8101908080519060200190929190505050105b15610ae3578273ffffffffffffffffffffffffffffffffffffffff1663095ea7b3827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b158015610aca57600080fd5b505af1158015610ade573d6000803e3d6000fd5b505050505b505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610ba35780471015610b2a57600080fd5b60003373ffffffffffffffffffffffffffffffffffffffff168260405180600001905060006040518083038185875af1925050503d8060008114610b8a576040519150601f19603f3d011682016040523d82523d6000602084013e610b8f565b606091505b5050905080610b9d57600080fd5b50610c3e565b8173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b158015610c1457600080fd5b505af1158015610c28573d6000803e3d6000fd5b50505050610c34610c42565b610c3d57600080fd5b5b5050565b600080600090503d60008114610c5f5760208114610c6857610c74565b60019150610c74565b60206000803e60005191505b5060008114159150509056fea2646970667358221220a7460028085c6e11e372c16d6920e4d6da54f17c255235d4ad961fb9b2e189b764736f6c63430007060033",
}

// PcurveABI is the input ABI used to generate the binding from.
// Deprecated: Use PcurveMetaData.ABI instead.
var PcurveABI = PcurveMetaData.ABI

// PcurveBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PcurveMetaData.Bin instead.
var PcurveBin = PcurveMetaData.Bin

// DeployPcurve deploys a new Ethereum contract, binding an instance of Pcurve to it.
func DeployPcurve(auth *bind.TransactOpts, backend bind.ContractBackend, _weth common.Address) (common.Address, *types.Transaction, *Pcurve, error) {
	parsed, err := PcurveMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PcurveBin), backend, _weth)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pcurve{PcurveCaller: PcurveCaller{contract: contract}, PcurveTransactor: PcurveTransactor{contract: contract}, PcurveFilterer: PcurveFilterer{contract: contract}}, nil
}

// Pcurve is an auto generated Go binding around an Ethereum contract.
type Pcurve struct {
	PcurveCaller     // Read-only binding to the contract
	PcurveTransactor // Write-only binding to the contract
	PcurveFilterer   // Log filterer for contract events
}

// PcurveCaller is an auto generated read-only Go binding around an Ethereum contract.
type PcurveCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PcurveTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PcurveTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PcurveFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PcurveFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PcurveSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PcurveSession struct {
	Contract     *Pcurve           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PcurveCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PcurveCallerSession struct {
	Contract *PcurveCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PcurveTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PcurveTransactorSession struct {
	Contract     *PcurveTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PcurveRaw is an auto generated low-level Go binding around an Ethereum contract.
type PcurveRaw struct {
	Contract *Pcurve // Generic contract binding to access the raw methods on
}

// PcurveCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PcurveCallerRaw struct {
	Contract *PcurveCaller // Generic read-only contract binding to access the raw methods on
}

// PcurveTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PcurveTransactorRaw struct {
	Contract *PcurveTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPcurve creates a new instance of Pcurve, bound to a specific deployed contract.
func NewPcurve(address common.Address, backend bind.ContractBackend) (*Pcurve, error) {
	contract, err := bindPcurve(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pcurve{PcurveCaller: PcurveCaller{contract: contract}, PcurveTransactor: PcurveTransactor{contract: contract}, PcurveFilterer: PcurveFilterer{contract: contract}}, nil
}

// NewPcurveCaller creates a new read-only instance of Pcurve, bound to a specific deployed contract.
func NewPcurveCaller(address common.Address, caller bind.ContractCaller) (*PcurveCaller, error) {
	contract, err := bindPcurve(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PcurveCaller{contract: contract}, nil
}

// NewPcurveTransactor creates a new write-only instance of Pcurve, bound to a specific deployed contract.
func NewPcurveTransactor(address common.Address, transactor bind.ContractTransactor) (*PcurveTransactor, error) {
	contract, err := bindPcurve(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PcurveTransactor{contract: contract}, nil
}

// NewPcurveFilterer creates a new log filterer instance of Pcurve, bound to a specific deployed contract.
func NewPcurveFilterer(address common.Address, filterer bind.ContractFilterer) (*PcurveFilterer, error) {
	contract, err := bindPcurve(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PcurveFilterer{contract: contract}, nil
}

// bindPcurve binds a generic wrapper to an already deployed contract.
func bindPcurve(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PcurveABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pcurve *PcurveRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pcurve.Contract.PcurveCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pcurve *PcurveRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pcurve.Contract.PcurveTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pcurve *PcurveRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pcurve.Contract.PcurveTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pcurve *PcurveCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pcurve.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pcurve *PcurveTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pcurve.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pcurve *PcurveTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pcurve.Contract.contract.Transact(opts, method, params...)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() view returns(address)
func (_Pcurve *PcurveCaller) ETHCONTRACTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pcurve.contract.Call(opts, &out, "ETH_CONTRACT_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() view returns(address)
func (_Pcurve *PcurveSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _Pcurve.Contract.ETHCONTRACTADDRESS(&_Pcurve.CallOpts)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() view returns(address)
func (_Pcurve *PcurveCallerSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _Pcurve.Contract.ETHCONTRACTADDRESS(&_Pcurve.CallOpts)
}

// ETHCURVEADDRESS is a free data retrieval call binding the contract method 0xfc0db319.
//
// Solidity: function ETH_CURVE_ADDRESS() view returns(address)
func (_Pcurve *PcurveCaller) ETHCURVEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pcurve.contract.Call(opts, &out, "ETH_CURVE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ETHCURVEADDRESS is a free data retrieval call binding the contract method 0xfc0db319.
//
// Solidity: function ETH_CURVE_ADDRESS() view returns(address)
func (_Pcurve *PcurveSession) ETHCURVEADDRESS() (common.Address, error) {
	return _Pcurve.Contract.ETHCURVEADDRESS(&_Pcurve.CallOpts)
}

// ETHCURVEADDRESS is a free data retrieval call binding the contract method 0xfc0db319.
//
// Solidity: function ETH_CURVE_ADDRESS() view returns(address)
func (_Pcurve *PcurveCallerSession) ETHCURVEADDRESS() (common.Address, error) {
	return _Pcurve.Contract.ETHCURVEADDRESS(&_Pcurve.CallOpts)
}

// MAX is a free data retrieval call binding the contract method 0xd49d5181.
//
// Solidity: function MAX() view returns(uint256)
func (_Pcurve *PcurveCaller) MAX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pcurve.contract.Call(opts, &out, "MAX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAX is a free data retrieval call binding the contract method 0xd49d5181.
//
// Solidity: function MAX() view returns(uint256)
func (_Pcurve *PcurveSession) MAX() (*big.Int, error) {
	return _Pcurve.Contract.MAX(&_Pcurve.CallOpts)
}

// MAX is a free data retrieval call binding the contract method 0xd49d5181.
//
// Solidity: function MAX() view returns(uint256)
func (_Pcurve *PcurveCallerSession) MAX() (*big.Int, error) {
	return _Pcurve.Contract.MAX(&_Pcurve.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Pcurve *PcurveCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pcurve.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Pcurve *PcurveSession) WETH() (common.Address, error) {
	return _Pcurve.Contract.WETH(&_Pcurve.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Pcurve *PcurveCallerSession) WETH() (common.Address, error) {
	return _Pcurve.Contract.WETH(&_Pcurve.CallOpts)
}

// Exchange is a paid mutator transaction binding the contract method 0xa64833a0.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 amount, uint256 minAmount, address curvePool) payable returns(address, uint256)
func (_Pcurve *PcurveTransactor) Exchange(opts *bind.TransactOpts, i *big.Int, j *big.Int, amount *big.Int, minAmount *big.Int, curvePool common.Address) (*types.Transaction, error) {
	return _Pcurve.contract.Transact(opts, "exchange", i, j, amount, minAmount, curvePool)
}

// Exchange is a paid mutator transaction binding the contract method 0xa64833a0.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 amount, uint256 minAmount, address curvePool) payable returns(address, uint256)
func (_Pcurve *PcurveSession) Exchange(i *big.Int, j *big.Int, amount *big.Int, minAmount *big.Int, curvePool common.Address) (*types.Transaction, error) {
	return _Pcurve.Contract.Exchange(&_Pcurve.TransactOpts, i, j, amount, minAmount, curvePool)
}

// Exchange is a paid mutator transaction binding the contract method 0xa64833a0.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 amount, uint256 minAmount, address curvePool) payable returns(address, uint256)
func (_Pcurve *PcurveTransactorSession) Exchange(i *big.Int, j *big.Int, amount *big.Int, minAmount *big.Int, curvePool common.Address) (*types.Transaction, error) {
	return _Pcurve.Contract.Exchange(&_Pcurve.TransactOpts, i, j, amount, minAmount, curvePool)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0x8c0b6593.
//
// Solidity: function exchangeUnderlying(uint256 i, uint256 j, uint256 amount, uint256 minAmount, address curvePool) payable returns(address, uint256)
func (_Pcurve *PcurveTransactor) ExchangeUnderlying(opts *bind.TransactOpts, i *big.Int, j *big.Int, amount *big.Int, minAmount *big.Int, curvePool common.Address) (*types.Transaction, error) {
	return _Pcurve.contract.Transact(opts, "exchangeUnderlying", i, j, amount, minAmount, curvePool)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0x8c0b6593.
//
// Solidity: function exchangeUnderlying(uint256 i, uint256 j, uint256 amount, uint256 minAmount, address curvePool) payable returns(address, uint256)
func (_Pcurve *PcurveSession) ExchangeUnderlying(i *big.Int, j *big.Int, amount *big.Int, minAmount *big.Int, curvePool common.Address) (*types.Transaction, error) {
	return _Pcurve.Contract.ExchangeUnderlying(&_Pcurve.TransactOpts, i, j, amount, minAmount, curvePool)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0x8c0b6593.
//
// Solidity: function exchangeUnderlying(uint256 i, uint256 j, uint256 amount, uint256 minAmount, address curvePool) payable returns(address, uint256)
func (_Pcurve *PcurveTransactorSession) ExchangeUnderlying(i *big.Int, j *big.Int, amount *big.Int, minAmount *big.Int, curvePool common.Address) (*types.Transaction, error) {
	return _Pcurve.Contract.ExchangeUnderlying(&_Pcurve.TransactOpts, i, j, amount, minAmount, curvePool)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Pcurve *PcurveTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pcurve.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Pcurve *PcurveSession) Receive() (*types.Transaction, error) {
	return _Pcurve.Contract.Receive(&_Pcurve.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Pcurve *PcurveTransactorSession) Receive() (*types.Transaction, error) {
	return _Pcurve.Contract.Receive(&_Pcurve.TransactOpts)
}
