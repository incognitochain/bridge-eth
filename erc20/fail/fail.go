// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fail

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

// FAILMetaData contains all meta data concerning the FAIL contract.
var FAILMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"unfreeze\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"freezeOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"freeze\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialSupply\",\"type\":\"uint256\"},{\"name\":\"tokenName\",\"type\":\"string\"},{\"name\":\"decimalUnits\",\"type\":\"uint8\"},{\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Freeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Unfreeze\",\"type\":\"event\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"42966c68": "burn(uint256)",
		"313ce567": "decimals()",
		"d7a78db8": "freeze(uint256)",
		"cd4217c1": "freezeOf(address)",
		"06fdde03": "name()",
		"8da5cb5b": "owner()",
		"95d89b41": "symbol()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
		"6623fc46": "unfreeze(uint256)",
		"3bed33ce": "withdrawEther(uint256)",
	},
	Bin: "0x60806040523480156200001157600080fd5b5060405162000d2938038062000d29833981018060405260808110156200003757600080fd5b8151602083018051919392830192916401000000008111156200005957600080fd5b820160208101848111156200006d57600080fd5b81516401000000008111828201871017156200008857600080fd5b50506020820151604090920180519194929391640100000000811115620000ae57600080fd5b82016020810184811115620000c257600080fd5b8151640100000000811182820187101715620000dd57600080fd5b50503360009081526005602090815260408220899055600389905587519295506200010f945090925086019062000156565b5080516200012590600190602084019062000156565b50506002805460ff90921660ff199092169190911790555050600480546001600160a01b03191633179055620001fb565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200019957805160ff1916838001178555620001c9565b82800160010185558215620001c9579182015b82811115620001c9578251825591602001919060010190620001ac565b50620001d7929150620001db565b5090565b620001f891905b80821115620001d75760008155600101620001e2565b90565b610b1e806200020b6000396000f3fe6080604052600436106100e85760003560e01c80636623fc461161008a578063a9059cbb11610059578063a9059cbb1461034d578063cd4217c114610386578063d7a78db8146103b9578063dd62ed3e146103e3576100e8565b80636623fc46146102aa57806370a08231146102d45780638da5cb5b1461030757806395d89b4114610338576100e8565b806323b872dd116100c657806323b872dd146101e8578063313ce5671461022b5780633bed33ce1461025657806342966c6814610280576100e8565b806306fdde03146100ea578063095ea7b31461017457806318160ddd146101c1575b005b3480156100f657600080fd5b506100ff61041e565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610139578181015183820152602001610121565b50505050905090810190601f1680156101665780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561018057600080fd5b506101ad6004803603604081101561019757600080fd5b506001600160a01b0381351690602001356104ac565b604080519115158252519081900360200190f35b3480156101cd57600080fd5b506101d66104e8565b60408051918252519081900360200190f35b3480156101f457600080fd5b506101ad6004803603606081101561020b57600080fd5b506001600160a01b038135811691602081013590911690604001356104ee565b34801561023757600080fd5b50610240610687565b6040805160ff9092168252519081900360200190f35b34801561026257600080fd5b506100e86004803603602081101561027957600080fd5b5035610690565b34801561028c57600080fd5b506101ad600480360360208110156102a357600080fd5b50356106e5565b3480156102b657600080fd5b506101ad600480360360208110156102cd57600080fd5b5035610786565b3480156102e057600080fd5b506101d6600480360360208110156102f757600080fd5b50356001600160a01b0316610840565b34801561031357600080fd5b5061031c610852565b604080516001600160a01b039092168252519081900360200190f35b34801561034457600080fd5b506100ff610861565b34801561035957600080fd5b506101ad6004803603604081101561037057600080fd5b506001600160a01b0381351690602001356108bb565b34801561039257600080fd5b506101d6600480360360208110156103a957600080fd5b50356001600160a01b03166109c4565b3480156103c557600080fd5b506101ad600480360360208110156103dc57600080fd5b50356109d6565b3480156103ef57600080fd5b506101d66004803603604081101561040657600080fd5b506001600160a01b0381358116916020013516610a90565b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156104a45780601f10610479576101008083540402835291602001916104a4565b820191906000526020600020905b81548152906001019060200180831161048757829003601f168201915b505050505081565b60008082116104ba57600080fd5b503360009081526007602090815260408083206001600160a01b039590951683529390529190912055600190565b60035481565b60006001600160a01b03831661050357600080fd5b6000821161051057600080fd5b6001600160a01b03841660009081526005602052604090205482111561053557600080fd5b6001600160a01b038316600090815260056020526040902054828101101561055c57600080fd5b6001600160a01b038416600090815260076020908152604080832033845290915290205482111561058c57600080fd5b6001600160a01b0384166000908152600560205260409020546105af9083610aad565b6001600160a01b0380861660009081526005602052604080822093909355908516815220546105de9083610ac1565b6001600160a01b03808516600090815260056020908152604080832094909455918716815260078252828120338252909152205461061c9083610aad565b6001600160a01b03808616600081815260076020908152604080832033845282529182902094909455805186815290519287169391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a35060009392505050565b60025460ff1681565b6004546001600160a01b031633146106a757600080fd5b6004546040516001600160a01b039091169082156108fc029083906000818181858888f193505050501580156106e1573d6000803e3d6000fd5b5050565b3360009081526005602052604081205482111561070157600080fd5b6000821161070e57600080fd5b336000908152600560205260409020546107289083610aad565b336000908152600560205260409020556003546107459083610aad565b60035560408051838152905133917fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5919081900360200190a2506001919050565b336000908152600660205260408120548211156107a257600080fd5b600082116107af57600080fd5b336000908152600660205260409020546107c99083610aad565b336000908152600660209081526040808320939093556005905220546107ef9083610ac1565b33600081815260056020908152604091829020939093558051858152905191927f2cfce4af01bcb9d6cf6c84ee1b7c491100b8695368264146a94d71e10a63083f92918290030190a2506001919050565b60056020526000908152604090205481565b6004546001600160a01b031681565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156104a45780601f10610479576101008083540402835291602001916104a4565b60006001600160a01b0383166108d057600080fd5b600082116108dd57600080fd5b336000908152600560205260409020548211156108f957600080fd5b6001600160a01b038316600090815260056020526040902054828101101561092057600080fd5b3360009081526005602052604090205461093a9083610aad565b33600090815260056020526040808220929092556001600160a01b038516815220546109669083610ac1565b6001600160a01b0384166000818152600560209081526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600092915050565b60066020526000908152604090205481565b336000908152600560205260408120548211156109f257600080fd5b600082116109ff57600080fd5b33600090815260056020526040902054610a199083610aad565b33600090815260056020908152604080832093909355600690522054610a3f9083610ac1565b33600081815260066020908152604091829020939093558051858152905191927ff97a274face0b5517365ad396b1fdba6f68bd3135ef603e44272adba3af5a1e092918290030190a2506001919050565b600760209081526000928352604080842090915290825290205481565b6000610abb83831115610ae5565b50900390565b6000828201610ade848210801590610ad95750838210155b610ae5565b9392505050565b80610aef57600080fd5b5056fea165627a7a72305820e13a3a20704341b26392f45d9be312643a54355446f0f0f41cac074bb4f3fa540029",
}

// FAILABI is the input ABI used to generate the binding from.
// Deprecated: Use FAILMetaData.ABI instead.
var FAILABI = FAILMetaData.ABI

// Deprecated: Use FAILMetaData.Sigs instead.
// FAILFuncSigs maps the 4-byte function signature to its string representation.
var FAILFuncSigs = FAILMetaData.Sigs

// FAILBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FAILMetaData.Bin instead.
var FAILBin = FAILMetaData.Bin

// DeployFAIL deploys a new Ethereum contract, binding an instance of FAIL to it.
func DeployFAIL(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int, tokenName string, decimalUnits uint8, tokenSymbol string) (common.Address, *types.Transaction, *FAIL, error) {
	parsed, err := FAILMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FAILBin), backend, initialSupply, tokenName, decimalUnits, tokenSymbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FAIL{FAILCaller: FAILCaller{contract: contract}, FAILTransactor: FAILTransactor{contract: contract}, FAILFilterer: FAILFilterer{contract: contract}}, nil
}

// FAIL is an auto generated Go binding around an Ethereum contract.
type FAIL struct {
	FAILCaller     // Read-only binding to the contract
	FAILTransactor // Write-only binding to the contract
	FAILFilterer   // Log filterer for contract events
}

// FAILCaller is an auto generated read-only Go binding around an Ethereum contract.
type FAILCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FAILTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FAILTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FAILFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FAILFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FAILSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FAILSession struct {
	Contract     *FAIL             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FAILCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FAILCallerSession struct {
	Contract *FAILCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FAILTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FAILTransactorSession struct {
	Contract     *FAILTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FAILRaw is an auto generated low-level Go binding around an Ethereum contract.
type FAILRaw struct {
	Contract *FAIL // Generic contract binding to access the raw methods on
}

// FAILCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FAILCallerRaw struct {
	Contract *FAILCaller // Generic read-only contract binding to access the raw methods on
}

// FAILTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FAILTransactorRaw struct {
	Contract *FAILTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFAIL creates a new instance of FAIL, bound to a specific deployed contract.
func NewFAIL(address common.Address, backend bind.ContractBackend) (*FAIL, error) {
	contract, err := bindFAIL(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FAIL{FAILCaller: FAILCaller{contract: contract}, FAILTransactor: FAILTransactor{contract: contract}, FAILFilterer: FAILFilterer{contract: contract}}, nil
}

// NewFAILCaller creates a new read-only instance of FAIL, bound to a specific deployed contract.
func NewFAILCaller(address common.Address, caller bind.ContractCaller) (*FAILCaller, error) {
	contract, err := bindFAIL(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FAILCaller{contract: contract}, nil
}

// NewFAILTransactor creates a new write-only instance of FAIL, bound to a specific deployed contract.
func NewFAILTransactor(address common.Address, transactor bind.ContractTransactor) (*FAILTransactor, error) {
	contract, err := bindFAIL(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FAILTransactor{contract: contract}, nil
}

// NewFAILFilterer creates a new log filterer instance of FAIL, bound to a specific deployed contract.
func NewFAILFilterer(address common.Address, filterer bind.ContractFilterer) (*FAILFilterer, error) {
	contract, err := bindFAIL(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FAILFilterer{contract: contract}, nil
}

// bindFAIL binds a generic wrapper to an already deployed contract.
func bindFAIL(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FAILABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FAIL *FAILRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FAIL.Contract.FAILCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FAIL *FAILRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FAIL.Contract.FAILTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FAIL *FAILRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FAIL.Contract.FAILTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FAIL *FAILCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FAIL.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FAIL *FAILTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FAIL.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FAIL *FAILTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FAIL.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_FAIL *FAILCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FAIL.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_FAIL *FAILSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _FAIL.Contract.Allowance(&_FAIL.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_FAIL *FAILCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _FAIL.Contract.Allowance(&_FAIL.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_FAIL *FAILCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FAIL.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_FAIL *FAILSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _FAIL.Contract.BalanceOf(&_FAIL.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_FAIL *FAILCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _FAIL.Contract.BalanceOf(&_FAIL.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FAIL *FAILCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _FAIL.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FAIL *FAILSession) Decimals() (uint8, error) {
	return _FAIL.Contract.Decimals(&_FAIL.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FAIL *FAILCallerSession) Decimals() (uint8, error) {
	return _FAIL.Contract.Decimals(&_FAIL.CallOpts)
}

// FreezeOf is a free data retrieval call binding the contract method 0xcd4217c1.
//
// Solidity: function freezeOf(address ) view returns(uint256)
func (_FAIL *FAILCaller) FreezeOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FAIL.contract.Call(opts, &out, "freezeOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FreezeOf is a free data retrieval call binding the contract method 0xcd4217c1.
//
// Solidity: function freezeOf(address ) view returns(uint256)
func (_FAIL *FAILSession) FreezeOf(arg0 common.Address) (*big.Int, error) {
	return _FAIL.Contract.FreezeOf(&_FAIL.CallOpts, arg0)
}

// FreezeOf is a free data retrieval call binding the contract method 0xcd4217c1.
//
// Solidity: function freezeOf(address ) view returns(uint256)
func (_FAIL *FAILCallerSession) FreezeOf(arg0 common.Address) (*big.Int, error) {
	return _FAIL.Contract.FreezeOf(&_FAIL.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FAIL *FAILCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FAIL.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FAIL *FAILSession) Name() (string, error) {
	return _FAIL.Contract.Name(&_FAIL.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FAIL *FAILCallerSession) Name() (string, error) {
	return _FAIL.Contract.Name(&_FAIL.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FAIL *FAILCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FAIL.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FAIL *FAILSession) Owner() (common.Address, error) {
	return _FAIL.Contract.Owner(&_FAIL.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FAIL *FAILCallerSession) Owner() (common.Address, error) {
	return _FAIL.Contract.Owner(&_FAIL.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FAIL *FAILCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FAIL.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FAIL *FAILSession) Symbol() (string, error) {
	return _FAIL.Contract.Symbol(&_FAIL.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FAIL *FAILCallerSession) Symbol() (string, error) {
	return _FAIL.Contract.Symbol(&_FAIL.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FAIL *FAILCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FAIL.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FAIL *FAILSession) TotalSupply() (*big.Int, error) {
	return _FAIL.Contract.TotalSupply(&_FAIL.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FAIL *FAILCallerSession) TotalSupply() (*big.Int, error) {
	return _FAIL.Contract.TotalSupply(&_FAIL.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_FAIL *FAILTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FAIL.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_FAIL *FAILSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FAIL.Contract.Approve(&_FAIL.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_FAIL *FAILTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FAIL.Contract.Approve(&_FAIL.TransactOpts, _spender, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns(bool success)
func (_FAIL *FAILTransactor) Burn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _FAIL.contract.Transact(opts, "burn", _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns(bool success)
func (_FAIL *FAILSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _FAIL.Contract.Burn(&_FAIL.TransactOpts, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns(bool success)
func (_FAIL *FAILTransactorSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _FAIL.Contract.Burn(&_FAIL.TransactOpts, _value)
}

// Freeze is a paid mutator transaction binding the contract method 0xd7a78db8.
//
// Solidity: function freeze(uint256 _value) returns(bool success)
func (_FAIL *FAILTransactor) Freeze(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _FAIL.contract.Transact(opts, "freeze", _value)
}

// Freeze is a paid mutator transaction binding the contract method 0xd7a78db8.
//
// Solidity: function freeze(uint256 _value) returns(bool success)
func (_FAIL *FAILSession) Freeze(_value *big.Int) (*types.Transaction, error) {
	return _FAIL.Contract.Freeze(&_FAIL.TransactOpts, _value)
}

// Freeze is a paid mutator transaction binding the contract method 0xd7a78db8.
//
// Solidity: function freeze(uint256 _value) returns(bool success)
func (_FAIL *FAILTransactorSession) Freeze(_value *big.Int) (*types.Transaction, error) {
	return _FAIL.Contract.Freeze(&_FAIL.TransactOpts, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_FAIL *FAILTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FAIL.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_FAIL *FAILSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FAIL.Contract.Transfer(&_FAIL.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_FAIL *FAILTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FAIL.Contract.Transfer(&_FAIL.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_FAIL *FAILTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FAIL.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_FAIL *FAILSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FAIL.Contract.TransferFrom(&_FAIL.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_FAIL *FAILTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FAIL.Contract.TransferFrom(&_FAIL.TransactOpts, _from, _to, _value)
}

// Unfreeze is a paid mutator transaction binding the contract method 0x6623fc46.
//
// Solidity: function unfreeze(uint256 _value) returns(bool success)
func (_FAIL *FAILTransactor) Unfreeze(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _FAIL.contract.Transact(opts, "unfreeze", _value)
}

// Unfreeze is a paid mutator transaction binding the contract method 0x6623fc46.
//
// Solidity: function unfreeze(uint256 _value) returns(bool success)
func (_FAIL *FAILSession) Unfreeze(_value *big.Int) (*types.Transaction, error) {
	return _FAIL.Contract.Unfreeze(&_FAIL.TransactOpts, _value)
}

// Unfreeze is a paid mutator transaction binding the contract method 0x6623fc46.
//
// Solidity: function unfreeze(uint256 _value) returns(bool success)
func (_FAIL *FAILTransactorSession) Unfreeze(_value *big.Int) (*types.Transaction, error) {
	return _FAIL.Contract.Unfreeze(&_FAIL.TransactOpts, _value)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x3bed33ce.
//
// Solidity: function withdrawEther(uint256 amount) returns()
func (_FAIL *FAILTransactor) WithdrawEther(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _FAIL.contract.Transact(opts, "withdrawEther", amount)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x3bed33ce.
//
// Solidity: function withdrawEther(uint256 amount) returns()
func (_FAIL *FAILSession) WithdrawEther(amount *big.Int) (*types.Transaction, error) {
	return _FAIL.Contract.WithdrawEther(&_FAIL.TransactOpts, amount)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x3bed33ce.
//
// Solidity: function withdrawEther(uint256 amount) returns()
func (_FAIL *FAILTransactorSession) WithdrawEther(amount *big.Int) (*types.Transaction, error) {
	return _FAIL.Contract.WithdrawEther(&_FAIL.TransactOpts, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_FAIL *FAILTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _FAIL.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_FAIL *FAILSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _FAIL.Contract.Fallback(&_FAIL.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_FAIL *FAILTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _FAIL.Contract.Fallback(&_FAIL.TransactOpts, calldata)
}

// FAILBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the FAIL contract.
type FAILBurnIterator struct {
	Event *FAILBurn // Event containing the contract specifics and raw log

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
func (it *FAILBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FAILBurn)
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
		it.Event = new(FAILBurn)
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
func (it *FAILBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FAILBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FAILBurn represents a Burn event raised by the FAIL contract.
type FAILBurn struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_FAIL *FAILFilterer) FilterBurn(opts *bind.FilterOpts, from []common.Address) (*FAILBurnIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _FAIL.contract.FilterLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return &FAILBurnIterator{contract: _FAIL.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_FAIL *FAILFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *FAILBurn, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _FAIL.contract.WatchLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FAILBurn)
				if err := _FAIL.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_FAIL *FAILFilterer) ParseBurn(log types.Log) (*FAILBurn, error) {
	event := new(FAILBurn)
	if err := _FAIL.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FAILFreezeIterator is returned from FilterFreeze and is used to iterate over the raw logs and unpacked data for Freeze events raised by the FAIL contract.
type FAILFreezeIterator struct {
	Event *FAILFreeze // Event containing the contract specifics and raw log

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
func (it *FAILFreezeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FAILFreeze)
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
		it.Event = new(FAILFreeze)
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
func (it *FAILFreezeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FAILFreezeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FAILFreeze represents a Freeze event raised by the FAIL contract.
type FAILFreeze struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterFreeze is a free log retrieval operation binding the contract event 0xf97a274face0b5517365ad396b1fdba6f68bd3135ef603e44272adba3af5a1e0.
//
// Solidity: event Freeze(address indexed from, uint256 value)
func (_FAIL *FAILFilterer) FilterFreeze(opts *bind.FilterOpts, from []common.Address) (*FAILFreezeIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _FAIL.contract.FilterLogs(opts, "Freeze", fromRule)
	if err != nil {
		return nil, err
	}
	return &FAILFreezeIterator{contract: _FAIL.contract, event: "Freeze", logs: logs, sub: sub}, nil
}

// WatchFreeze is a free log subscription operation binding the contract event 0xf97a274face0b5517365ad396b1fdba6f68bd3135ef603e44272adba3af5a1e0.
//
// Solidity: event Freeze(address indexed from, uint256 value)
func (_FAIL *FAILFilterer) WatchFreeze(opts *bind.WatchOpts, sink chan<- *FAILFreeze, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _FAIL.contract.WatchLogs(opts, "Freeze", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FAILFreeze)
				if err := _FAIL.contract.UnpackLog(event, "Freeze", log); err != nil {
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

// ParseFreeze is a log parse operation binding the contract event 0xf97a274face0b5517365ad396b1fdba6f68bd3135ef603e44272adba3af5a1e0.
//
// Solidity: event Freeze(address indexed from, uint256 value)
func (_FAIL *FAILFilterer) ParseFreeze(log types.Log) (*FAILFreeze, error) {
	event := new(FAILFreeze)
	if err := _FAIL.contract.UnpackLog(event, "Freeze", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FAILTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the FAIL contract.
type FAILTransferIterator struct {
	Event *FAILTransfer // Event containing the contract specifics and raw log

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
func (it *FAILTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FAILTransfer)
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
		it.Event = new(FAILTransfer)
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
func (it *FAILTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FAILTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FAILTransfer represents a Transfer event raised by the FAIL contract.
type FAILTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_FAIL *FAILFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FAILTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FAIL.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FAILTransferIterator{contract: _FAIL.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_FAIL *FAILFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *FAILTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FAIL.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FAILTransfer)
				if err := _FAIL.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_FAIL *FAILFilterer) ParseTransfer(log types.Log) (*FAILTransfer, error) {
	event := new(FAILTransfer)
	if err := _FAIL.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FAILUnfreezeIterator is returned from FilterUnfreeze and is used to iterate over the raw logs and unpacked data for Unfreeze events raised by the FAIL contract.
type FAILUnfreezeIterator struct {
	Event *FAILUnfreeze // Event containing the contract specifics and raw log

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
func (it *FAILUnfreezeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FAILUnfreeze)
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
		it.Event = new(FAILUnfreeze)
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
func (it *FAILUnfreezeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FAILUnfreezeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FAILUnfreeze represents a Unfreeze event raised by the FAIL contract.
type FAILUnfreeze struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnfreeze is a free log retrieval operation binding the contract event 0x2cfce4af01bcb9d6cf6c84ee1b7c491100b8695368264146a94d71e10a63083f.
//
// Solidity: event Unfreeze(address indexed from, uint256 value)
func (_FAIL *FAILFilterer) FilterUnfreeze(opts *bind.FilterOpts, from []common.Address) (*FAILUnfreezeIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _FAIL.contract.FilterLogs(opts, "Unfreeze", fromRule)
	if err != nil {
		return nil, err
	}
	return &FAILUnfreezeIterator{contract: _FAIL.contract, event: "Unfreeze", logs: logs, sub: sub}, nil
}

// WatchUnfreeze is a free log subscription operation binding the contract event 0x2cfce4af01bcb9d6cf6c84ee1b7c491100b8695368264146a94d71e10a63083f.
//
// Solidity: event Unfreeze(address indexed from, uint256 value)
func (_FAIL *FAILFilterer) WatchUnfreeze(opts *bind.WatchOpts, sink chan<- *FAILUnfreeze, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _FAIL.contract.WatchLogs(opts, "Unfreeze", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FAILUnfreeze)
				if err := _FAIL.contract.UnpackLog(event, "Unfreeze", log); err != nil {
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

// ParseUnfreeze is a log parse operation binding the contract event 0x2cfce4af01bcb9d6cf6c84ee1b7c491100b8695368264146a94d71e10a63083f.
//
// Solidity: event Unfreeze(address indexed from, uint256 value)
func (_FAIL *FAILFilterer) ParseUnfreeze(log types.Log) (*FAILUnfreeze, error) {
	event := new(FAILUnfreeze)
	if err := _FAIL.contract.UnpackLog(event, "Unfreeze", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeMathMetaData contains all meta data concerning the SafeMath contract.
var SafeMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080604052348015600f57600080fd5b50603580601d6000396000f3fe6080604052600080fdfea165627a7a7230582046754e83e4c8d5016b0fc42c835943a106d4509d54322d7fb57ab8d3f41317f00029",
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
