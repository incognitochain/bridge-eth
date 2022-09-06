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
	ABI: "[{\"inputs\":[],\"name\":\"ETH_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETH_CURVE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"source\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address[6]\",\"name\":\"route\",\"type\":\"address[6]\"},{\"internalType\":\"uint256[8]\",\"name\":\"indices\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256\",\"name\":\"mintReceived\",\"type\":\"uint256\"},{\"internalType\":\"contractICurveSwap\",\"name\":\"curvePool\",\"type\":\"address\"}],\"name\":\"exchange\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"j\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"contractICurveSwap\",\"name\":\"curvePool\",\"type\":\"address\"}],\"name\":\"exchange\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"j\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"contractICurveSwap\",\"name\":\"curvePool\",\"type\":\"address\"}],\"name\":\"exchangeUnderlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b506110c7806100206000396000f3fe6080604052600436106100745760003560e01c8063a64833a01161004e578063a64833a01461024a578063ad5c4648146102f4578063d49d518114610335578063fc0db319146103605761007b565b806323bdc7b41461008057806372e94bf61461015f5780638c0b6593146101a05761007b565b3661007b57005b600080fd5b61012c600480360361026081101561009757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908060c001909192919290806101000190919291929080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506103a1565b604051808373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b34801561016b57600080fd5b506101746105f2565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156101ac57600080fd5b50610217600480360360a08110156101c357600080fd5b8101908080359060200190929190803590602001909291908035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506105f7565b604051808373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b34801561025657600080fd5b506102c1600480360360a081101561026d57600080fd5b8101908080359060200190929190803590602001909291908035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610913565b604051808373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b34801561030057600080fd5b50610309610c2f565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561034157600080fd5b5061034a610c53565b6040518082815260200191505060405180910390f35b34801561036c57600080fd5b50610375610c77565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6000806000871161041a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f696e76616c6964207377617020616d6f756e740000000000000000000000000081525060200191505060405180910390fd5b600073eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee73ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff1614610469578861046c565b60005b90506104798a8986610c8f565b8373ffffffffffffffffffffffffffffffffffffffff1663d6db9993348a8a8a8a306040518763ffffffff1660e01b81526004018086815260200185600660200280828437600081840152601f19601f82011690508083019250505084600860200280828437600081840152601f19601f8201169050808301925050508381526020018273ffffffffffffffffffffffffffffffffffffffff168152602001955050505050506000604051808303818588803b15801561053857600080fd5b505af115801561054c573d6000803e3d6000fd5b5050505050600061055c8a610e11565b9050858110156105d4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f4e6f7420656e6f75676820636f696e000000000000000000000000000000000081525060200191505060405180910390fd5b6105de8282610ef9565b818193509350505097509795505050505050565b600081565b60008060008511610670576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f696e76616c6964207377617020616d6f756e740000000000000000000000000081525060200191505060405180910390fd5b60008373ffffffffffffffffffffffffffffffffffffffff1663b9947eb0896040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156106c357600080fd5b505afa1580156106d7573d6000803e3d6000fd5b505050506040513d60208110156106ed57600080fd5b8101908080519060200190929190505050905060008473ffffffffffffffffffffffffffffffffffffffff1663b9947eb0896040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561075357600080fd5b505afa158015610767573d6000803e3d6000fd5b505050506040513d602081101561077d57600080fd5b81019080805190602001909291905050509050600073eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16146107df57816107e2565b60005b90506107ef838988610c8f565b8573ffffffffffffffffffffffffffffffffffffffff166365b2489b8b8b8b8b6040518563ffffffff1660e01b815260040180858152602001848152602001838152602001828152602001945050505050600060405180830381600087803b15801561085a57600080fd5b505af115801561086e573d6000803e3d6000fd5b50505050600061087d82610e11565b9050878110156108f5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f4e6f7420656e6f75676820636f696e000000000000000000000000000000000081525060200191505060405180910390fd5b6108ff8282610ef9565b818195509550505050509550959350505050565b6000806000851161098c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f696e76616c6964207377617020616d6f756e740000000000000000000000000081525060200191505060405180910390fd5b60008373ffffffffffffffffffffffffffffffffffffffff1663c6610657896040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156109df57600080fd5b505afa1580156109f3573d6000803e3d6000fd5b505050506040513d6020811015610a0957600080fd5b8101908080519060200190929190505050905060008473ffffffffffffffffffffffffffffffffffffffff1663c6610657896040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015610a6f57600080fd5b505afa158015610a83573d6000803e3d6000fd5b505050506040513d6020811015610a9957600080fd5b81019080805190602001909291905050509050600073eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614610afb5781610afe565b60005b9050610b0b838988610c8f565b8573ffffffffffffffffffffffffffffffffffffffff16635b41b9088b8b8b8b6040518563ffffffff1660e01b815260040180858152602001848152602001838152602001828152602001945050505050600060405180830381600087803b158015610b7657600080fd5b505af1158015610b8a573d6000803e3d6000fd5b505050506000610b9982610e11565b905087811015610c11576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f4e6f7420656e6f75676820636f696e000000000000000000000000000000000081525060200191505060405180910390fd5b610c1b8282610ef9565b818195509550505050509550959350505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee81565b600034148015610d5d5750818373ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e30846040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1681526020019250505060206040518083038186803b158015610d2057600080fd5b505afa158015610d34573d6000803e3d6000fd5b505050506040513d6020811015610d4a57600080fd5b8101908080519060200190929190505050105b15610e0c578273ffffffffffffffffffffffffffffffffffffffff1663095ea7b3827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b158015610df357600080fd5b505af1158015610e07573d6000803e3d6000fd5b505050505b505050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610e4f57479050610ef4565b8173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015610eb657600080fd5b505afa158015610eca573d6000803e3d6000fd5b505050506040513d6020811015610ee057600080fd5b810190808051906020019092919050505090505b919050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610fb45780471015610f3b57600080fd5b60003373ffffffffffffffffffffffffffffffffffffffff168260405180600001905060006040518083038185875af1925050503d8060008114610f9b576040519150601f19603f3d011682016040523d82523d6000602084013e610fa0565b606091505b5050905080610fae57600080fd5b5061104f565b8173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b15801561102557600080fd5b505af1158015611039573d6000803e3d6000fd5b50505050611045611053565b61104e57600080fd5b5b5050565b600080600090503d60008114611070576020811461107957611085565b60019150611085565b60206000803e60005191505b5060008114159150509056fea2646970667358221220416518f94439b75267e97339b54271d8f31e01142a27cb36efed51b3717f8a1f64736f6c63430007060033",
}

// PcurveABI is the input ABI used to generate the binding from.
// Deprecated: Use PcurveMetaData.ABI instead.
var PcurveABI = PcurveMetaData.ABI

// PcurveBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PcurveMetaData.Bin instead.
var PcurveBin = PcurveMetaData.Bin

// DeployPcurve deploys a new Ethereum contract, binding an instance of Pcurve to it.
func DeployPcurve(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pcurve, error) {
	parsed, err := PcurveMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PcurveBin), backend)
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

// Exchange is a paid mutator transaction binding the contract method 0x23bdc7b4.
//
// Solidity: function exchange(address source, address dest, uint256 amount, address[6] route, uint256[8] indices, uint256 mintReceived, address curvePool) payable returns(address, uint256)
func (_Pcurve *PcurveTransactor) Exchange(opts *bind.TransactOpts, source common.Address, dest common.Address, amount *big.Int, route [6]common.Address, indices [8]*big.Int, mintReceived *big.Int, curvePool common.Address) (*types.Transaction, error) {
	return _Pcurve.contract.Transact(opts, "exchange", source, dest, amount, route, indices, mintReceived, curvePool)
}

// Exchange is a paid mutator transaction binding the contract method 0x23bdc7b4.
//
// Solidity: function exchange(address source, address dest, uint256 amount, address[6] route, uint256[8] indices, uint256 mintReceived, address curvePool) payable returns(address, uint256)
func (_Pcurve *PcurveSession) Exchange(source common.Address, dest common.Address, amount *big.Int, route [6]common.Address, indices [8]*big.Int, mintReceived *big.Int, curvePool common.Address) (*types.Transaction, error) {
	return _Pcurve.Contract.Exchange(&_Pcurve.TransactOpts, source, dest, amount, route, indices, mintReceived, curvePool)
}

// Exchange is a paid mutator transaction binding the contract method 0x23bdc7b4.
//
// Solidity: function exchange(address source, address dest, uint256 amount, address[6] route, uint256[8] indices, uint256 mintReceived, address curvePool) payable returns(address, uint256)
func (_Pcurve *PcurveTransactorSession) Exchange(source common.Address, dest common.Address, amount *big.Int, route [6]common.Address, indices [8]*big.Int, mintReceived *big.Int, curvePool common.Address) (*types.Transaction, error) {
	return _Pcurve.Contract.Exchange(&_Pcurve.TransactOpts, source, dest, amount, route, indices, mintReceived, curvePool)
}

// Exchange0 is a paid mutator transaction binding the contract method 0xa64833a0.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 amount, uint256 minAmount, address curvePool) returns(address, uint256)
func (_Pcurve *PcurveTransactor) Exchange0(opts *bind.TransactOpts, i *big.Int, j *big.Int, amount *big.Int, minAmount *big.Int, curvePool common.Address) (*types.Transaction, error) {
	return _Pcurve.contract.Transact(opts, "exchange0", i, j, amount, minAmount, curvePool)
}

// Exchange0 is a paid mutator transaction binding the contract method 0xa64833a0.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 amount, uint256 minAmount, address curvePool) returns(address, uint256)
func (_Pcurve *PcurveSession) Exchange0(i *big.Int, j *big.Int, amount *big.Int, minAmount *big.Int, curvePool common.Address) (*types.Transaction, error) {
	return _Pcurve.Contract.Exchange0(&_Pcurve.TransactOpts, i, j, amount, minAmount, curvePool)
}

// Exchange0 is a paid mutator transaction binding the contract method 0xa64833a0.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 amount, uint256 minAmount, address curvePool) returns(address, uint256)
func (_Pcurve *PcurveTransactorSession) Exchange0(i *big.Int, j *big.Int, amount *big.Int, minAmount *big.Int, curvePool common.Address) (*types.Transaction, error) {
	return _Pcurve.Contract.Exchange0(&_Pcurve.TransactOpts, i, j, amount, minAmount, curvePool)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0x8c0b6593.
//
// Solidity: function exchangeUnderlying(uint256 i, uint256 j, uint256 amount, uint256 minAmount, address curvePool) returns(address, uint256)
func (_Pcurve *PcurveTransactor) ExchangeUnderlying(opts *bind.TransactOpts, i *big.Int, j *big.Int, amount *big.Int, minAmount *big.Int, curvePool common.Address) (*types.Transaction, error) {
	return _Pcurve.contract.Transact(opts, "exchangeUnderlying", i, j, amount, minAmount, curvePool)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0x8c0b6593.
//
// Solidity: function exchangeUnderlying(uint256 i, uint256 j, uint256 amount, uint256 minAmount, address curvePool) returns(address, uint256)
func (_Pcurve *PcurveSession) ExchangeUnderlying(i *big.Int, j *big.Int, amount *big.Int, minAmount *big.Int, curvePool common.Address) (*types.Transaction, error) {
	return _Pcurve.Contract.ExchangeUnderlying(&_Pcurve.TransactOpts, i, j, amount, minAmount, curvePool)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0x8c0b6593.
//
// Solidity: function exchangeUnderlying(uint256 i, uint256 j, uint256 amount, uint256 minAmount, address curvePool) returns(address, uint256)
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
