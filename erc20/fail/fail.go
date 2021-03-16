// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fail

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// FAILABI is the input ABI used to generate the binding from.
const FAILABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimalUnits\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Freeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Unfreeze\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"freeze\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"freezeOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receive\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"unfreeze\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawEther\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FAILFuncSigs maps the 4-byte function signature to its string representation.
var FAILFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"42966c68": "burn(uint256)",
	"313ce567": "decimals()",
	"d7a78db8": "freeze(uint256)",
	"cd4217c1": "freezeOf(address)",
	"06fdde03": "name()",
	"8da5cb5b": "owner()",
	"a3e76c0f": "receive()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"6623fc46": "unfreeze(uint256)",
	"3bed33ce": "withdrawEther(uint256)",
}

// FAILBin is the compiled bytecode used for deploying new contracts.
var FAILBin = "0x60806040523480156200001157600080fd5b5060405162000e0d38038062000e0d833981810160405260808110156200003757600080fd5b8151602083018051604051929492938301929190846401000000008211156200005f57600080fd5b9083019060208201858111156200007557600080fd5b82516401000000008111828201881017156200009057600080fd5b82525081516020918201929091019080838360005b83811015620000bf578181015183820152602001620000a5565b50505050905090810190601f168015620000ed5780820380516001836020036101000a031916815260200191505b506040818152602083015192018051929491939192846401000000008211156200011657600080fd5b9083019060208201858111156200012c57600080fd5b82516401000000008111828201881017156200014757600080fd5b82525081516020918201929091019080838360005b83811015620001765781810151838201526020016200015c565b50505050905090810190601f168015620001a45780820380516001836020036101000a031916815260200191505b5060409081523360009081526005602090815291812089905560038990558751620001d8955090935090870191506200021f565b508051620001ee9060019060208401906200021f565b50506002805460ff90921660ff199092169190911790555050600480546001600160a01b03191633179055620002bb565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200026257805160ff191683800117855562000292565b8280016001018555821562000292579182015b828111156200029257825182559160200191906001019062000275565b50620002a0929150620002a4565b5090565b5b80821115620002a05760008155600101620002a5565b610b4280620002cb6000396000f3fe6080604052600436106100f35760003560e01c806370a082311161008a578063a9059cbb11610059578063a9059cbb14610365578063cd4217c11461039e578063d7a78db8146103d1578063dd62ed3e146103fb576100f3565b806370a08231146102e45780638da5cb5b1461031757806395d89b4114610348578063a3e76c0f1461035d576100f3565b8063313ce567116100c6578063313ce567146102395780633bed33ce1461026457806342966c68146102905780636623fc46146102ba576100f3565b806306fdde03146100f8578063095ea7b31461018257806318160ddd146101cf57806323b872dd146101f6575b600080fd5b34801561010457600080fd5b5061010d610436565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561014757818101518382015260200161012f565b50505050905090810190601f1680156101745780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561018e57600080fd5b506101bb600480360360408110156101a557600080fd5b506001600160a01b0381351690602001356104c4565b604080519115158252519081900360200190f35b3480156101db57600080fd5b506101e4610500565b60408051918252519081900360200190f35b34801561020257600080fd5b506101bb6004803603606081101561021957600080fd5b506001600160a01b03813581169160208101359091169060400135610506565b34801561024557600080fd5b5061024e61069f565b6040805160ff9092168252519081900360200190f35b34801561027057600080fd5b5061028e6004803603602081101561028757600080fd5b50356106a8565b005b34801561029c57600080fd5b506101bb600480360360208110156102b357600080fd5b50356106fd565b3480156102c657600080fd5b506101bb600480360360208110156102dd57600080fd5b503561079e565b3480156102f057600080fd5b506101e46004803603602081101561030757600080fd5b50356001600160a01b0316610858565b34801561032357600080fd5b5061032c61086a565b604080516001600160a01b039092168252519081900360200190f35b34801561035457600080fd5b5061010d610879565b61028e6108d3565b34801561037157600080fd5b506101bb6004803603604081101561038857600080fd5b506001600160a01b0381351690602001356108d5565b3480156103aa57600080fd5b506101e4600480360360208110156103c157600080fd5b50356001600160a01b03166109de565b3480156103dd57600080fd5b506101bb600480360360208110156103f457600080fd5b50356109f0565b34801561040757600080fd5b506101e46004803603604081101561041e57600080fd5b506001600160a01b0381358116916020013516610aaa565b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156104bc5780601f10610491576101008083540402835291602001916104bc565b820191906000526020600020905b81548152906001019060200180831161049f57829003601f168201915b505050505081565b60008082116104d257600080fd5b503360009081526007602090815260408083206001600160a01b039590951683529390529190912055600190565b60035481565b60006001600160a01b03831661051b57600080fd5b6000821161052857600080fd5b6001600160a01b03841660009081526005602052604090205482111561054d57600080fd5b6001600160a01b038316600090815260056020526040902054828101101561057457600080fd5b6001600160a01b03841660009081526007602090815260408083203384529091529020548211156105a457600080fd5b6001600160a01b0384166000908152600560205260409020546105c79083610ac7565b6001600160a01b0380861660009081526005602052604080822093909355908516815220546105f69083610adb565b6001600160a01b0380851660009081526005602090815260408083209490945591871681526007825282812033825290915220546106349083610ac7565b6001600160a01b03808616600081815260076020908152604080832033845282529182902094909455805186815290519287169391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a35060009392505050565b60025460ff1681565b6004546001600160a01b031633146106bf57600080fd5b6004546040516001600160a01b039091169082156108fc029083906000818181858888f193505050501580156106f9573d6000803e3d6000fd5b5050565b3360009081526005602052604081205482111561071957600080fd5b6000821161072657600080fd5b336000908152600560205260409020546107409083610ac7565b3360009081526005602052604090205560035461075d9083610ac7565b60035560408051838152905133917fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5919081900360200190a2506001919050565b336000908152600660205260408120548211156107ba57600080fd5b600082116107c757600080fd5b336000908152600660205260409020546107e19083610ac7565b336000908152600660209081526040808320939093556005905220546108079083610adb565b33600081815260056020908152604091829020939093558051858152905191927f2cfce4af01bcb9d6cf6c84ee1b7c491100b8695368264146a94d71e10a63083f92918290030190a2506001919050565b60056020526000908152604090205481565b6004546001600160a01b031681565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156104bc5780601f10610491576101008083540402835291602001916104bc565b565b60006001600160a01b0383166108ea57600080fd5b600082116108f757600080fd5b3360009081526005602052604090205482111561091357600080fd5b6001600160a01b038316600090815260056020526040902054828101101561093a57600080fd5b336000908152600560205260409020546109549083610ac7565b33600090815260056020526040808220929092556001600160a01b038516815220546109809083610adb565b6001600160a01b0384166000818152600560209081526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600092915050565b60066020526000908152604090205481565b33600090815260056020526040812054821115610a0c57600080fd5b60008211610a1957600080fd5b33600090815260056020526040902054610a339083610ac7565b33600090815260056020908152604080832093909355600690522054610a599083610adb565b33600081815260066020908152604091829020939093558051858152905191927ff97a274face0b5517365ad396b1fdba6f68bd3135ef603e44272adba3af5a1e092918290030190a2506001919050565b600760209081526000928352604080842090915290825290205481565b6000610ad583831115610aff565b50900390565b6000828201610af8848210801590610af35750838210155b610aff565b9392505050565b80610b0957600080fd5b5056fea26469706673582212201ce1a209406b8986ca3e105dd9cdfb7864b4dafd9524dc37c570ece048d0e8cb64736f6c634300060c0033"

// DeployFAIL deploys a new Ethereum contract, binding an instance of FAIL to it.
func DeployFAIL(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int, tokenName string, decimalUnits uint8, tokenSymbol string) (common.Address, *types.Transaction, *FAIL, error) {
	parsed, err := abi.JSON(strings.NewReader(FAILABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FAILBin), backend, initialSupply, tokenName, decimalUnits, tokenSymbol)
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
func (_FAIL *FAILRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_FAIL *FAILCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
// Solidity: function allowance(address , address ) constant returns(uint256)
func (_FAIL *FAILCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FAIL.contract.Call(opts, out, "allowance", arg0, arg1)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) constant returns(uint256)
func (_FAIL *FAILSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _FAIL.Contract.Allowance(&_FAIL.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) constant returns(uint256)
func (_FAIL *FAILCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _FAIL.Contract.Allowance(&_FAIL.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) constant returns(uint256)
func (_FAIL *FAILCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FAIL.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) constant returns(uint256)
func (_FAIL *FAILSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _FAIL.Contract.BalanceOf(&_FAIL.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) constant returns(uint256)
func (_FAIL *FAILCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _FAIL.Contract.BalanceOf(&_FAIL.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_FAIL *FAILCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _FAIL.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_FAIL *FAILSession) Decimals() (uint8, error) {
	return _FAIL.Contract.Decimals(&_FAIL.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_FAIL *FAILCallerSession) Decimals() (uint8, error) {
	return _FAIL.Contract.Decimals(&_FAIL.CallOpts)
}

// FreezeOf is a free data retrieval call binding the contract method 0xcd4217c1.
//
// Solidity: function freezeOf(address ) constant returns(uint256)
func (_FAIL *FAILCaller) FreezeOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FAIL.contract.Call(opts, out, "freezeOf", arg0)
	return *ret0, err
}

// FreezeOf is a free data retrieval call binding the contract method 0xcd4217c1.
//
// Solidity: function freezeOf(address ) constant returns(uint256)
func (_FAIL *FAILSession) FreezeOf(arg0 common.Address) (*big.Int, error) {
	return _FAIL.Contract.FreezeOf(&_FAIL.CallOpts, arg0)
}

// FreezeOf is a free data retrieval call binding the contract method 0xcd4217c1.
//
// Solidity: function freezeOf(address ) constant returns(uint256)
func (_FAIL *FAILCallerSession) FreezeOf(arg0 common.Address) (*big.Int, error) {
	return _FAIL.Contract.FreezeOf(&_FAIL.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_FAIL *FAILCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FAIL.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_FAIL *FAILSession) Name() (string, error) {
	return _FAIL.Contract.Name(&_FAIL.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_FAIL *FAILCallerSession) Name() (string, error) {
	return _FAIL.Contract.Name(&_FAIL.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FAIL *FAILCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FAIL.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FAIL *FAILSession) Owner() (common.Address, error) {
	return _FAIL.Contract.Owner(&_FAIL.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FAIL *FAILCallerSession) Owner() (common.Address, error) {
	return _FAIL.Contract.Owner(&_FAIL.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_FAIL *FAILCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FAIL.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_FAIL *FAILSession) Symbol() (string, error) {
	return _FAIL.Contract.Symbol(&_FAIL.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_FAIL *FAILCallerSession) Symbol() (string, error) {
	return _FAIL.Contract.Symbol(&_FAIL.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_FAIL *FAILCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FAIL.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_FAIL *FAILSession) TotalSupply() (*big.Int, error) {
	return _FAIL.Contract.TotalSupply(&_FAIL.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
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

// Receive is a paid mutator transaction binding the contract method 0xa3e76c0f.
//
// Solidity: function receive() returns()
func (_FAIL *FAILTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FAIL.contract.Transact(opts, "receive")
}

// Receive is a paid mutator transaction binding the contract method 0xa3e76c0f.
//
// Solidity: function receive() returns()
func (_FAIL *FAILSession) Receive() (*types.Transaction, error) {
	return _FAIL.Contract.Receive(&_FAIL.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract method 0xa3e76c0f.
//
// Solidity: function receive() returns()
func (_FAIL *FAILTransactorSession) Receive() (*types.Transaction, error) {
	return _FAIL.Contract.Receive(&_FAIL.TransactOpts)
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
	return event, nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea26469706673582212201af705d885f676e1d14ca60eec16c920a58eca3bb2edc5347426a16c8f3d640d64736f6c634300060c0033"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
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
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
