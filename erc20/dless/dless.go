// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dless

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

// DlessABI is the input ABI used to generate the binding from.
const DlessABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Freeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Unfreeze\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"freeze\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"freezeOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receive\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"unfreeze\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawEther\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DlessBin is the compiled bytecode used for deploying new contracts.
var DlessBin = "0x60806040523480156200001157600080fd5b5060405162000db538038062000db5833981810160405260608110156200003757600080fd5b8151602083018051604051929492938301929190846401000000008211156200005f57600080fd5b9083019060208201858111156200007557600080fd5b82516401000000008111828201881017156200009057600080fd5b82525081516020918201929091019080838360005b83811015620000bf578181015183820152602001620000a5565b50505050905090810190601f168015620000ed5780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200011157600080fd5b9083019060208201858111156200012757600080fd5b82516401000000008111828201881017156200014257600080fd5b82525081516020918201929091019080838360005b838110156200017157818101518382015260200162000157565b50505050905090810190601f1680156200019f5780820380516001836020036101000a031916815260200191505b5060409081523360009081526004602090815291812088905560028890558651620001d39550909350908601915062000206565b508051620001e990600190602084019062000206565b5050600380546001600160a01b0319163317905550620002a29050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200024957805160ff191683800117855562000279565b8280016001018555821562000279579182015b82811115620002795782518255916020019190600101906200025c565b50620002879291506200028b565b5090565b5b808211156200028757600081556001016200028c565b610b0380620002b26000396000f3fe6080604052600436106100e85760003560e01c806370a082311161008a578063a9059cbb11610059578063a9059cbb1461032f578063cd4217c114610368578063d7a78db81461039b578063dd62ed3e146103c5576100e8565b806370a08231146102ae5780638da5cb5b146102e157806395d89b4114610312578063a3e76c0f14610327576100e8565b806323b872dd116100c657806323b872dd146101eb5780633bed33ce1461022e57806342966c681461025a5780636623fc4614610284576100e8565b806306fdde03146100ed578063095ea7b31461017757806318160ddd146101c4575b600080fd5b3480156100f957600080fd5b50610102610400565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561013c578181015183820152602001610124565b50505050905090810190601f1680156101695780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561018357600080fd5b506101b06004803603604081101561019a57600080fd5b506001600160a01b03813516906020013561048e565b604080519115158252519081900360200190f35b3480156101d057600080fd5b506101d96104ca565b60408051918252519081900360200190f35b3480156101f757600080fd5b506101b06004803603606081101561020e57600080fd5b506001600160a01b038135811691602081013590911690604001356104d0565b34801561023a57600080fd5b506102586004803603602081101561025157600080fd5b5035610669565b005b34801561026657600080fd5b506101b06004803603602081101561027d57600080fd5b50356106be565b34801561029057600080fd5b506101b0600480360360208110156102a757600080fd5b503561075f565b3480156102ba57600080fd5b506101d9600480360360208110156102d157600080fd5b50356001600160a01b0316610819565b3480156102ed57600080fd5b506102f661082b565b604080516001600160a01b039092168252519081900360200190f35b34801561031e57600080fd5b5061010261083a565b610258610894565b34801561033b57600080fd5b506101b06004803603604081101561035257600080fd5b506001600160a01b038135169060200135610896565b34801561037457600080fd5b506101d96004803603602081101561038b57600080fd5b50356001600160a01b031661099f565b3480156103a757600080fd5b506101b0600480360360208110156103be57600080fd5b50356109b1565b3480156103d157600080fd5b506101d9600480360360408110156103e857600080fd5b506001600160a01b0381358116916020013516610a6b565b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156104865780601f1061045b57610100808354040283529160200191610486565b820191906000526020600020905b81548152906001019060200180831161046957829003601f168201915b505050505081565b600080821161049c57600080fd5b503360009081526006602090815260408083206001600160a01b039590951683529390529190912055600190565b60025481565b60006001600160a01b0383166104e557600080fd5b600082116104f257600080fd5b6001600160a01b03841660009081526004602052604090205482111561051757600080fd5b6001600160a01b038316600090815260046020526040902054828101101561053e57600080fd5b6001600160a01b038416600090815260066020908152604080832033845290915290205482111561056e57600080fd5b6001600160a01b0384166000908152600460205260409020546105919083610a88565b6001600160a01b0380861660009081526004602052604080822093909355908516815220546105c09083610a9c565b6001600160a01b0380851660009081526004602090815260408083209490945591871681526006825282812033825290915220546105fe9083610a88565b6001600160a01b03808616600081815260066020908152604080832033845282529182902094909455805186815290519287169391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a35060019392505050565b6003546001600160a01b0316331461068057600080fd5b6003546040516001600160a01b039091169082156108fc029083906000818181858888f193505050501580156106ba573d6000803e3d6000fd5b5050565b336000908152600460205260408120548211156106da57600080fd5b600082116106e757600080fd5b336000908152600460205260409020546107019083610a88565b3360009081526004602052604090205560025461071e9083610a88565b60025560408051838152905133917fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5919081900360200190a2506001919050565b3360009081526005602052604081205482111561077b57600080fd5b6000821161078857600080fd5b336000908152600560205260409020546107a29083610a88565b336000908152600560209081526040808320939093556004905220546107c89083610a9c565b33600081815260046020908152604091829020939093558051858152905191927f2cfce4af01bcb9d6cf6c84ee1b7c491100b8695368264146a94d71e10a63083f92918290030190a2506001919050565b60046020526000908152604090205481565b6003546001600160a01b031681565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156104865780601f1061045b57610100808354040283529160200191610486565b565b60006001600160a01b0383166108ab57600080fd5b600082116108b857600080fd5b336000908152600460205260409020548211156108d457600080fd5b6001600160a01b03831660009081526004602052604090205482810110156108fb57600080fd5b336000908152600460205260409020546109159083610a88565b33600090815260046020526040808220929092556001600160a01b038516815220546109419083610a9c565b6001600160a01b0384166000818152600460209081526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b60056020526000908152604090205481565b336000908152600460205260408120548211156109cd57600080fd5b600082116109da57600080fd5b336000908152600460205260409020546109f49083610a88565b33600090815260046020908152604080832093909355600590522054610a1a9083610a9c565b33600081815260056020908152604091829020939093558051858152905191927ff97a274face0b5517365ad396b1fdba6f68bd3135ef603e44272adba3af5a1e092918290030190a2506001919050565b600660209081526000928352604080842090915290825290205481565b6000610a9683831115610ac0565b50900390565b6000828201610ab9848210801590610ab45750838210155b610ac0565b9392505050565b80610aca57600080fd5b5056fea264697066735822122099332c6723ec3f109547e259dcec5d0ac533ab63a39ddfe375bc628eeb88fae464736f6c634300060c0033"

// DeployDless deploys a new Ethereum contract, binding an instance of Dless to it.
func DeployDless(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int, tokenName string, tokenSymbol string) (common.Address, *types.Transaction, *Dless, error) {
	parsed, err := abi.JSON(strings.NewReader(DlessABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DlessBin), backend, initialSupply, tokenName, tokenSymbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Dless{DlessCaller: DlessCaller{contract: contract}, DlessTransactor: DlessTransactor{contract: contract}, DlessFilterer: DlessFilterer{contract: contract}}, nil
}

// Dless is an auto generated Go binding around an Ethereum contract.
type Dless struct {
	DlessCaller     // Read-only binding to the contract
	DlessTransactor // Write-only binding to the contract
	DlessFilterer   // Log filterer for contract events
}

// DlessCaller is an auto generated read-only Go binding around an Ethereum contract.
type DlessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DlessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DlessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DlessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DlessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DlessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DlessSession struct {
	Contract     *Dless            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DlessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DlessCallerSession struct {
	Contract *DlessCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DlessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DlessTransactorSession struct {
	Contract     *DlessTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DlessRaw is an auto generated low-level Go binding around an Ethereum contract.
type DlessRaw struct {
	Contract *Dless // Generic contract binding to access the raw methods on
}

// DlessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DlessCallerRaw struct {
	Contract *DlessCaller // Generic read-only contract binding to access the raw methods on
}

// DlessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DlessTransactorRaw struct {
	Contract *DlessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDless creates a new instance of Dless, bound to a specific deployed contract.
func NewDless(address common.Address, backend bind.ContractBackend) (*Dless, error) {
	contract, err := bindDless(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dless{DlessCaller: DlessCaller{contract: contract}, DlessTransactor: DlessTransactor{contract: contract}, DlessFilterer: DlessFilterer{contract: contract}}, nil
}

// NewDlessCaller creates a new read-only instance of Dless, bound to a specific deployed contract.
func NewDlessCaller(address common.Address, caller bind.ContractCaller) (*DlessCaller, error) {
	contract, err := bindDless(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DlessCaller{contract: contract}, nil
}

// NewDlessTransactor creates a new write-only instance of Dless, bound to a specific deployed contract.
func NewDlessTransactor(address common.Address, transactor bind.ContractTransactor) (*DlessTransactor, error) {
	contract, err := bindDless(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DlessTransactor{contract: contract}, nil
}

// NewDlessFilterer creates a new log filterer instance of Dless, bound to a specific deployed contract.
func NewDlessFilterer(address common.Address, filterer bind.ContractFilterer) (*DlessFilterer, error) {
	contract, err := bindDless(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DlessFilterer{contract: contract}, nil
}

// bindDless binds a generic wrapper to an already deployed contract.
func bindDless(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DlessABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dless *DlessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dless.Contract.DlessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dless *DlessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dless.Contract.DlessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dless *DlessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dless.Contract.DlessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dless *DlessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dless.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dless *DlessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dless.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dless *DlessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dless.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Dless *DlessCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dless.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Dless *DlessSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Dless.Contract.Allowance(&_Dless.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Dless *DlessCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Dless.Contract.Allowance(&_Dless.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Dless *DlessCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dless.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Dless *DlessSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Dless.Contract.BalanceOf(&_Dless.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Dless *DlessCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Dless.Contract.BalanceOf(&_Dless.CallOpts, arg0)
}

// FreezeOf is a free data retrieval call binding the contract method 0xcd4217c1.
//
// Solidity: function freezeOf(address ) view returns(uint256)
func (_Dless *DlessCaller) FreezeOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dless.contract.Call(opts, &out, "freezeOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FreezeOf is a free data retrieval call binding the contract method 0xcd4217c1.
//
// Solidity: function freezeOf(address ) view returns(uint256)
func (_Dless *DlessSession) FreezeOf(arg0 common.Address) (*big.Int, error) {
	return _Dless.Contract.FreezeOf(&_Dless.CallOpts, arg0)
}

// FreezeOf is a free data retrieval call binding the contract method 0xcd4217c1.
//
// Solidity: function freezeOf(address ) view returns(uint256)
func (_Dless *DlessCallerSession) FreezeOf(arg0 common.Address) (*big.Int, error) {
	return _Dless.Contract.FreezeOf(&_Dless.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Dless *DlessCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Dless.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Dless *DlessSession) Name() (string, error) {
	return _Dless.Contract.Name(&_Dless.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Dless *DlessCallerSession) Name() (string, error) {
	return _Dless.Contract.Name(&_Dless.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dless *DlessCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dless.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dless *DlessSession) Owner() (common.Address, error) {
	return _Dless.Contract.Owner(&_Dless.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dless *DlessCallerSession) Owner() (common.Address, error) {
	return _Dless.Contract.Owner(&_Dless.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Dless *DlessCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Dless.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Dless *DlessSession) Symbol() (string, error) {
	return _Dless.Contract.Symbol(&_Dless.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Dless *DlessCallerSession) Symbol() (string, error) {
	return _Dless.Contract.Symbol(&_Dless.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Dless *DlessCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dless.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Dless *DlessSession) TotalSupply() (*big.Int, error) {
	return _Dless.Contract.TotalSupply(&_Dless.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Dless *DlessCallerSession) TotalSupply() (*big.Int, error) {
	return _Dless.Contract.TotalSupply(&_Dless.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Dless *DlessTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Dless.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Dless *DlessSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Dless.Contract.Approve(&_Dless.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Dless *DlessTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Dless.Contract.Approve(&_Dless.TransactOpts, _spender, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns(bool success)
func (_Dless *DlessTransactor) Burn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Dless.contract.Transact(opts, "burn", _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns(bool success)
func (_Dless *DlessSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _Dless.Contract.Burn(&_Dless.TransactOpts, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns(bool success)
func (_Dless *DlessTransactorSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _Dless.Contract.Burn(&_Dless.TransactOpts, _value)
}

// Freeze is a paid mutator transaction binding the contract method 0xd7a78db8.
//
// Solidity: function freeze(uint256 _value) returns(bool success)
func (_Dless *DlessTransactor) Freeze(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Dless.contract.Transact(opts, "freeze", _value)
}

// Freeze is a paid mutator transaction binding the contract method 0xd7a78db8.
//
// Solidity: function freeze(uint256 _value) returns(bool success)
func (_Dless *DlessSession) Freeze(_value *big.Int) (*types.Transaction, error) {
	return _Dless.Contract.Freeze(&_Dless.TransactOpts, _value)
}

// Freeze is a paid mutator transaction binding the contract method 0xd7a78db8.
//
// Solidity: function freeze(uint256 _value) returns(bool success)
func (_Dless *DlessTransactorSession) Freeze(_value *big.Int) (*types.Transaction, error) {
	return _Dless.Contract.Freeze(&_Dless.TransactOpts, _value)
}

// Receive is a paid mutator transaction binding the contract method 0xa3e76c0f.
//
// Solidity: function receive() payable returns()
func (_Dless *DlessTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dless.contract.Transact(opts, "receive")
}

// Receive is a paid mutator transaction binding the contract method 0xa3e76c0f.
//
// Solidity: function receive() payable returns()
func (_Dless *DlessSession) Receive() (*types.Transaction, error) {
	return _Dless.Contract.Receive(&_Dless.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract method 0xa3e76c0f.
//
// Solidity: function receive() payable returns()
func (_Dless *DlessTransactorSession) Receive() (*types.Transaction, error) {
	return _Dless.Contract.Receive(&_Dless.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_Dless *DlessTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Dless.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_Dless *DlessSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Dless.Contract.Transfer(&_Dless.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_Dless *DlessTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Dless.Contract.Transfer(&_Dless.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_Dless *DlessTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Dless.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_Dless *DlessSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Dless.Contract.TransferFrom(&_Dless.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_Dless *DlessTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Dless.Contract.TransferFrom(&_Dless.TransactOpts, _from, _to, _value)
}

// Unfreeze is a paid mutator transaction binding the contract method 0x6623fc46.
//
// Solidity: function unfreeze(uint256 _value) returns(bool success)
func (_Dless *DlessTransactor) Unfreeze(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Dless.contract.Transact(opts, "unfreeze", _value)
}

// Unfreeze is a paid mutator transaction binding the contract method 0x6623fc46.
//
// Solidity: function unfreeze(uint256 _value) returns(bool success)
func (_Dless *DlessSession) Unfreeze(_value *big.Int) (*types.Transaction, error) {
	return _Dless.Contract.Unfreeze(&_Dless.TransactOpts, _value)
}

// Unfreeze is a paid mutator transaction binding the contract method 0x6623fc46.
//
// Solidity: function unfreeze(uint256 _value) returns(bool success)
func (_Dless *DlessTransactorSession) Unfreeze(_value *big.Int) (*types.Transaction, error) {
	return _Dless.Contract.Unfreeze(&_Dless.TransactOpts, _value)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x3bed33ce.
//
// Solidity: function withdrawEther(uint256 amount) returns()
func (_Dless *DlessTransactor) WithdrawEther(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Dless.contract.Transact(opts, "withdrawEther", amount)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x3bed33ce.
//
// Solidity: function withdrawEther(uint256 amount) returns()
func (_Dless *DlessSession) WithdrawEther(amount *big.Int) (*types.Transaction, error) {
	return _Dless.Contract.WithdrawEther(&_Dless.TransactOpts, amount)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x3bed33ce.
//
// Solidity: function withdrawEther(uint256 amount) returns()
func (_Dless *DlessTransactorSession) WithdrawEther(amount *big.Int) (*types.Transaction, error) {
	return _Dless.Contract.WithdrawEther(&_Dless.TransactOpts, amount)
}

// DlessBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Dless contract.
type DlessBurnIterator struct {
	Event *DlessBurn // Event containing the contract specifics and raw log

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
func (it *DlessBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DlessBurn)
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
		it.Event = new(DlessBurn)
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
func (it *DlessBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DlessBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DlessBurn represents a Burn event raised by the Dless contract.
type DlessBurn struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_Dless *DlessFilterer) FilterBurn(opts *bind.FilterOpts, from []common.Address) (*DlessBurnIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Dless.contract.FilterLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return &DlessBurnIterator{contract: _Dless.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_Dless *DlessFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *DlessBurn, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Dless.contract.WatchLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DlessBurn)
				if err := _Dless.contract.UnpackLog(event, "Burn", log); err != nil {
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
func (_Dless *DlessFilterer) ParseBurn(log types.Log) (*DlessBurn, error) {
	event := new(DlessBurn)
	if err := _Dless.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DlessFreezeIterator is returned from FilterFreeze and is used to iterate over the raw logs and unpacked data for Freeze events raised by the Dless contract.
type DlessFreezeIterator struct {
	Event *DlessFreeze // Event containing the contract specifics and raw log

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
func (it *DlessFreezeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DlessFreeze)
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
		it.Event = new(DlessFreeze)
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
func (it *DlessFreezeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DlessFreezeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DlessFreeze represents a Freeze event raised by the Dless contract.
type DlessFreeze struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterFreeze is a free log retrieval operation binding the contract event 0xf97a274face0b5517365ad396b1fdba6f68bd3135ef603e44272adba3af5a1e0.
//
// Solidity: event Freeze(address indexed from, uint256 value)
func (_Dless *DlessFilterer) FilterFreeze(opts *bind.FilterOpts, from []common.Address) (*DlessFreezeIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Dless.contract.FilterLogs(opts, "Freeze", fromRule)
	if err != nil {
		return nil, err
	}
	return &DlessFreezeIterator{contract: _Dless.contract, event: "Freeze", logs: logs, sub: sub}, nil
}

// WatchFreeze is a free log subscription operation binding the contract event 0xf97a274face0b5517365ad396b1fdba6f68bd3135ef603e44272adba3af5a1e0.
//
// Solidity: event Freeze(address indexed from, uint256 value)
func (_Dless *DlessFilterer) WatchFreeze(opts *bind.WatchOpts, sink chan<- *DlessFreeze, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Dless.contract.WatchLogs(opts, "Freeze", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DlessFreeze)
				if err := _Dless.contract.UnpackLog(event, "Freeze", log); err != nil {
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
func (_Dless *DlessFilterer) ParseFreeze(log types.Log) (*DlessFreeze, error) {
	event := new(DlessFreeze)
	if err := _Dless.contract.UnpackLog(event, "Freeze", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DlessTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Dless contract.
type DlessTransferIterator struct {
	Event *DlessTransfer // Event containing the contract specifics and raw log

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
func (it *DlessTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DlessTransfer)
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
		it.Event = new(DlessTransfer)
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
func (it *DlessTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DlessTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DlessTransfer represents a Transfer event raised by the Dless contract.
type DlessTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Dless *DlessFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DlessTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Dless.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DlessTransferIterator{contract: _Dless.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Dless *DlessFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DlessTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Dless.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DlessTransfer)
				if err := _Dless.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Dless *DlessFilterer) ParseTransfer(log types.Log) (*DlessTransfer, error) {
	event := new(DlessTransfer)
	if err := _Dless.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DlessUnfreezeIterator is returned from FilterUnfreeze and is used to iterate over the raw logs and unpacked data for Unfreeze events raised by the Dless contract.
type DlessUnfreezeIterator struct {
	Event *DlessUnfreeze // Event containing the contract specifics and raw log

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
func (it *DlessUnfreezeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DlessUnfreeze)
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
		it.Event = new(DlessUnfreeze)
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
func (it *DlessUnfreezeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DlessUnfreezeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DlessUnfreeze represents a Unfreeze event raised by the Dless contract.
type DlessUnfreeze struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnfreeze is a free log retrieval operation binding the contract event 0x2cfce4af01bcb9d6cf6c84ee1b7c491100b8695368264146a94d71e10a63083f.
//
// Solidity: event Unfreeze(address indexed from, uint256 value)
func (_Dless *DlessFilterer) FilterUnfreeze(opts *bind.FilterOpts, from []common.Address) (*DlessUnfreezeIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Dless.contract.FilterLogs(opts, "Unfreeze", fromRule)
	if err != nil {
		return nil, err
	}
	return &DlessUnfreezeIterator{contract: _Dless.contract, event: "Unfreeze", logs: logs, sub: sub}, nil
}

// WatchUnfreeze is a free log subscription operation binding the contract event 0x2cfce4af01bcb9d6cf6c84ee1b7c491100b8695368264146a94d71e10a63083f.
//
// Solidity: event Unfreeze(address indexed from, uint256 value)
func (_Dless *DlessFilterer) WatchUnfreeze(opts *bind.WatchOpts, sink chan<- *DlessUnfreeze, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Dless.contract.WatchLogs(opts, "Unfreeze", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DlessUnfreeze)
				if err := _Dless.contract.UnpackLog(event, "Unfreeze", log); err != nil {
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
func (_Dless *DlessFilterer) ParseUnfreeze(log types.Log) (*DlessUnfreeze, error) {
	event := new(DlessUnfreeze)
	if err := _Dless.contract.UnpackLog(event, "Unfreeze", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
