// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bnb

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

// BnbABI is the input ABI used to generate the binding from.
const BnbABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawEther\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"unfreeze\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"freezeOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"freeze\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialSupply\",\"type\":\"uint256\"},{\"name\":\"tokenName\",\"type\":\"string\"},{\"name\":\"decimalUnits\",\"type\":\"uint8\"},{\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"payable\":false,\"type\":\"constructor\"},{\"payable\":true,\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Freeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Unfreeze\",\"type\":\"event\"}]"

// BnbBin is the compiled bytecode used for deploying new contracts.
var BnbBin = "0x606060405234156200001057600080fd5b604051620016d8380380620016d8833981016040528080519060200190919080518201919060200180519060200190919080518201919050505b83600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550836003819055508260009080519060200190620000ad9291906200012e565b508060019080519060200190620000c69291906200012e565b5081600260006101000a81548160ff021916908360ff16021790555033600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50505050620001dd565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200017157805160ff1916838001178555620001a2565b82800160010185558215620001a2579182015b82811115620001a157825182559160200191906001019062000184565b5b509050620001b19190620001b5565b5090565b620001da91905b80821115620001d6576000816000905550600101620001bc565b5090565b90565b6114eb80620001ed6000396000f300606060405236156100d9576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306fdde03146100e2578063095ea7b31461017157806318160ddd146101cb57806323b872dd146101f4578063313ce5671461026d5780633bed33ce1461029c57806342966c68146102bf5780636623fc46146102fa57806370a08231146103355780638da5cb5b1461038257806395d89b41146103d7578063a9059cbb14610466578063cd4217c1146104a8578063d7a78db8146104f5578063dd62ed3e14610530575b6100e05b5b565b005b34156100ed57600080fd5b6100f561059c565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101365780820151818401525b60208101905061011a565b50505050905090810190601f1680156101635780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561017c57600080fd5b6101b1600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001909190505061063a565b604051808215151515815260200191505060405180910390f35b34156101d657600080fd5b6101de6106d6565b6040518082815260200191505060405180910390f35b34156101ff57600080fd5b610253600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803573ffffffffffffffffffffffffffffffffffffffff169060200190919080359060200190919050506106dc565b604051808215151515815260200191505060405180910390f35b341561027857600080fd5b610280610b01565b604051808260ff1660ff16815260200191505060405180910390f35b34156102a757600080fd5b6102bd6004808035906020019091905050610b14565b005b34156102ca57600080fd5b6102e06004808035906020019091905050610bd6565b604051808215151515815260200191505060405180910390f35b341561030557600080fd5b61031b6004808035906020019091905050610d29565b604051808215151515815260200191505060405180910390f35b341561034057600080fd5b61036c600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610ef6565b6040518082815260200191505060405180910390f35b341561038d57600080fd5b610395610f0e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156103e257600080fd5b6103ea610f34565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561042b5780820151818401525b60208101905061040f565b50505050905090810190601f1680156104585780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561047157600080fd5b6104a6600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091908035906020019091905050610fd2565b005b34156104b357600080fd5b6104df600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050611260565b6040518082815260200191505060405180910390f35b341561050057600080fd5b6105166004808035906020019091905050611278565b604051808215151515815260200191505060405180910390f35b341561053b57600080fd5b610586600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050611445565b6040518082815260200191505060405180910390f35b60008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106325780601f1061060757610100808354040283529160200191610632565b820191906000526020600020905b81548152906001019060200180831161061557829003601f168201915b505050505081565b6000808211151561064a57600080fd5b81600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600190505b92915050565b60035481565b6000808373ffffffffffffffffffffffffffffffffffffffff16141561070157600080fd5b60008211151561071057600080fd5b81600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101561075c57600080fd5b600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020540110156107e957600080fd5b600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482111561087257600080fd5b6108bb600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548361146a565b600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610947600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205483611484565b600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610a10600760008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548361146a565b600760008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a3600190505b9392505050565b600260009054906101000a900460ff1681565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610b7057600080fd5b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501515610bd257600080fd5b5b50565b600081600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610c2457600080fd5b600082111515610c3357600080fd5b610c7c600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548361146a565b600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610ccb6003548361146a565b6003819055503373ffffffffffffffffffffffffffffffffffffffff167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5836040518082815260200191505060405180910390a2600190505b919050565b600081600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610d7757600080fd5b600082111515610d8657600080fd5b610dcf600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548361146a565b600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610e5b600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205483611484565b600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055503373ffffffffffffffffffffffffffffffffffffffff167f2cfce4af01bcb9d6cf6c84ee1b7c491100b8695368264146a94d71e10a63083f836040518082815260200191505060405180910390a2600190505b919050565b60056020528060005260406000206000915090505481565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610fca5780601f10610f9f57610100808354040283529160200191610fca565b820191906000526020600020905b815481529060010190602001808311610fad57829003601f168201915b505050505081565b60008273ffffffffffffffffffffffffffffffffffffffff161415610ff657600080fd5b60008111151561100557600080fd5b80600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101561105157600080fd5b600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205481600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020540110156110de57600080fd5b611127600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548261146a565b600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506111b3600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482611484565b600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35b5050565b60066020528060005260406000206000915090505481565b600081600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410156112c657600080fd5b6000821115156112d557600080fd5b61131e600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548361146a565b600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506113aa600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205483611484565b600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055503373ffffffffffffffffffffffffffffffffffffffff167ff97a274face0b5517365ad396b1fdba6f68bd3135ef603e44272adba3af5a1e0836040518082815260200191505060405180910390a2600190505b919050565b6007602052816000526040600020602052806000526040600020600091509150505481565b6000611478838311156114af565b81830390505b92915050565b60008082840190506114a484821015801561149f5750838210155b6114af565b8091505b5092915050565b8015156114bb57600080fd5b5b505600a165627a7a72305820082734e053ffbdf2a3195354a3210dff3723c239a1e76ae3be0936f6aed31bee0029000000000000000000000000000000000000000000a56fa5b99019a5c80000000000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000001200000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000000000000003424e4200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003424e420000000000000000000000000000000000000000000000000000000000"

// DeployBnb deploys a new Ethereum contract, binding an instance of Bnb to it.
func DeployBnb(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int, tokenName string, decimalUnits uint8, tokenSymbol string) (common.Address, *types.Transaction, *Bnb, error) {
	parsed, err := abi.JSON(strings.NewReader(BnbABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BnbBin), backend, initialSupply, tokenName, decimalUnits, tokenSymbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bnb{BnbCaller: BnbCaller{contract: contract}, BnbTransactor: BnbTransactor{contract: contract}, BnbFilterer: BnbFilterer{contract: contract}}, nil
}

// Bnb is an auto generated Go binding around an Ethereum contract.
type Bnb struct {
	BnbCaller     // Read-only binding to the contract
	BnbTransactor // Write-only binding to the contract
	BnbFilterer   // Log filterer for contract events
}

// BnbCaller is an auto generated read-only Go binding around an Ethereum contract.
type BnbCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BnbTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BnbTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BnbFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BnbFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BnbSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BnbSession struct {
	Contract     *Bnb              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BnbCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BnbCallerSession struct {
	Contract *BnbCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BnbTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BnbTransactorSession struct {
	Contract     *BnbTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BnbRaw is an auto generated low-level Go binding around an Ethereum contract.
type BnbRaw struct {
	Contract *Bnb // Generic contract binding to access the raw methods on
}

// BnbCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BnbCallerRaw struct {
	Contract *BnbCaller // Generic read-only contract binding to access the raw methods on
}

// BnbTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BnbTransactorRaw struct {
	Contract *BnbTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBnb creates a new instance of Bnb, bound to a specific deployed contract.
func NewBnb(address common.Address, backend bind.ContractBackend) (*Bnb, error) {
	contract, err := bindBnb(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bnb{BnbCaller: BnbCaller{contract: contract}, BnbTransactor: BnbTransactor{contract: contract}, BnbFilterer: BnbFilterer{contract: contract}}, nil
}

// NewBnbCaller creates a new read-only instance of Bnb, bound to a specific deployed contract.
func NewBnbCaller(address common.Address, caller bind.ContractCaller) (*BnbCaller, error) {
	contract, err := bindBnb(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BnbCaller{contract: contract}, nil
}

// NewBnbTransactor creates a new write-only instance of Bnb, bound to a specific deployed contract.
func NewBnbTransactor(address common.Address, transactor bind.ContractTransactor) (*BnbTransactor, error) {
	contract, err := bindBnb(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BnbTransactor{contract: contract}, nil
}

// NewBnbFilterer creates a new log filterer instance of Bnb, bound to a specific deployed contract.
func NewBnbFilterer(address common.Address, filterer bind.ContractFilterer) (*BnbFilterer, error) {
	contract, err := bindBnb(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BnbFilterer{contract: contract}, nil
}

// bindBnb binds a generic wrapper to an already deployed contract.
func bindBnb(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BnbABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bnb *BnbRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bnb.Contract.BnbCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bnb *BnbRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bnb.Contract.BnbTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bnb *BnbRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bnb.Contract.BnbTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bnb *BnbCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bnb.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bnb *BnbTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bnb.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bnb *BnbTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bnb.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) returns(uint256)
func (_Bnb *BnbCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bnb.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) returns(uint256)
func (_Bnb *BnbSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Bnb.Contract.Allowance(&_Bnb.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) returns(uint256)
func (_Bnb *BnbCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Bnb.Contract.Allowance(&_Bnb.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) returns(uint256)
func (_Bnb *BnbCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bnb.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) returns(uint256)
func (_Bnb *BnbSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Bnb.Contract.BalanceOf(&_Bnb.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) returns(uint256)
func (_Bnb *BnbCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Bnb.Contract.BalanceOf(&_Bnb.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint8)
func (_Bnb *BnbCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Bnb.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint8)
func (_Bnb *BnbSession) Decimals() (uint8, error) {
	return _Bnb.Contract.Decimals(&_Bnb.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint8)
func (_Bnb *BnbCallerSession) Decimals() (uint8, error) {
	return _Bnb.Contract.Decimals(&_Bnb.CallOpts)
}

// FreezeOf is a free data retrieval call binding the contract method 0xcd4217c1.
//
// Solidity: function freezeOf(address ) returns(uint256)
func (_Bnb *BnbCaller) FreezeOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bnb.contract.Call(opts, &out, "freezeOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FreezeOf is a free data retrieval call binding the contract method 0xcd4217c1.
//
// Solidity: function freezeOf(address ) returns(uint256)
func (_Bnb *BnbSession) FreezeOf(arg0 common.Address) (*big.Int, error) {
	return _Bnb.Contract.FreezeOf(&_Bnb.CallOpts, arg0)
}

// FreezeOf is a free data retrieval call binding the contract method 0xcd4217c1.
//
// Solidity: function freezeOf(address ) returns(uint256)
func (_Bnb *BnbCallerSession) FreezeOf(arg0 common.Address) (*big.Int, error) {
	return _Bnb.Contract.FreezeOf(&_Bnb.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string)
func (_Bnb *BnbCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bnb.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string)
func (_Bnb *BnbSession) Name() (string, error) {
	return _Bnb.Contract.Name(&_Bnb.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string)
func (_Bnb *BnbCallerSession) Name() (string, error) {
	return _Bnb.Contract.Name(&_Bnb.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address)
func (_Bnb *BnbCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bnb.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address)
func (_Bnb *BnbSession) Owner() (common.Address, error) {
	return _Bnb.Contract.Owner(&_Bnb.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address)
func (_Bnb *BnbCallerSession) Owner() (common.Address, error) {
	return _Bnb.Contract.Owner(&_Bnb.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(string)
func (_Bnb *BnbCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bnb.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(string)
func (_Bnb *BnbSession) Symbol() (string, error) {
	return _Bnb.Contract.Symbol(&_Bnb.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(string)
func (_Bnb *BnbCallerSession) Symbol() (string, error) {
	return _Bnb.Contract.Symbol(&_Bnb.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_Bnb *BnbCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bnb.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_Bnb *BnbSession) TotalSupply() (*big.Int, error) {
	return _Bnb.Contract.TotalSupply(&_Bnb.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_Bnb *BnbCallerSession) TotalSupply() (*big.Int, error) {
	return _Bnb.Contract.TotalSupply(&_Bnb.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Bnb *BnbTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Bnb.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Bnb *BnbSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Bnb.Contract.Approve(&_Bnb.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Bnb *BnbTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Bnb.Contract.Approve(&_Bnb.TransactOpts, _spender, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns(bool success)
func (_Bnb *BnbTransactor) Burn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Bnb.contract.Transact(opts, "burn", _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns(bool success)
func (_Bnb *BnbSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _Bnb.Contract.Burn(&_Bnb.TransactOpts, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns(bool success)
func (_Bnb *BnbTransactorSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _Bnb.Contract.Burn(&_Bnb.TransactOpts, _value)
}

// Freeze is a paid mutator transaction binding the contract method 0xd7a78db8.
//
// Solidity: function freeze(uint256 _value) returns(bool success)
func (_Bnb *BnbTransactor) Freeze(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Bnb.contract.Transact(opts, "freeze", _value)
}

// Freeze is a paid mutator transaction binding the contract method 0xd7a78db8.
//
// Solidity: function freeze(uint256 _value) returns(bool success)
func (_Bnb *BnbSession) Freeze(_value *big.Int) (*types.Transaction, error) {
	return _Bnb.Contract.Freeze(&_Bnb.TransactOpts, _value)
}

// Freeze is a paid mutator transaction binding the contract method 0xd7a78db8.
//
// Solidity: function freeze(uint256 _value) returns(bool success)
func (_Bnb *BnbTransactorSession) Freeze(_value *big.Int) (*types.Transaction, error) {
	return _Bnb.Contract.Freeze(&_Bnb.TransactOpts, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns()
func (_Bnb *BnbTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Bnb.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns()
func (_Bnb *BnbSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Bnb.Contract.Transfer(&_Bnb.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns()
func (_Bnb *BnbTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Bnb.Contract.Transfer(&_Bnb.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_Bnb *BnbTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Bnb.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_Bnb *BnbSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Bnb.Contract.TransferFrom(&_Bnb.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_Bnb *BnbTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Bnb.Contract.TransferFrom(&_Bnb.TransactOpts, _from, _to, _value)
}

// Unfreeze is a paid mutator transaction binding the contract method 0x6623fc46.
//
// Solidity: function unfreeze(uint256 _value) returns(bool success)
func (_Bnb *BnbTransactor) Unfreeze(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Bnb.contract.Transact(opts, "unfreeze", _value)
}

// Unfreeze is a paid mutator transaction binding the contract method 0x6623fc46.
//
// Solidity: function unfreeze(uint256 _value) returns(bool success)
func (_Bnb *BnbSession) Unfreeze(_value *big.Int) (*types.Transaction, error) {
	return _Bnb.Contract.Unfreeze(&_Bnb.TransactOpts, _value)
}

// Unfreeze is a paid mutator transaction binding the contract method 0x6623fc46.
//
// Solidity: function unfreeze(uint256 _value) returns(bool success)
func (_Bnb *BnbTransactorSession) Unfreeze(_value *big.Int) (*types.Transaction, error) {
	return _Bnb.Contract.Unfreeze(&_Bnb.TransactOpts, _value)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x3bed33ce.
//
// Solidity: function withdrawEther(uint256 amount) returns()
func (_Bnb *BnbTransactor) WithdrawEther(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Bnb.contract.Transact(opts, "withdrawEther", amount)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x3bed33ce.
//
// Solidity: function withdrawEther(uint256 amount) returns()
func (_Bnb *BnbSession) WithdrawEther(amount *big.Int) (*types.Transaction, error) {
	return _Bnb.Contract.WithdrawEther(&_Bnb.TransactOpts, amount)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x3bed33ce.
//
// Solidity: function withdrawEther(uint256 amount) returns()
func (_Bnb *BnbTransactorSession) WithdrawEther(amount *big.Int) (*types.Transaction, error) {
	return _Bnb.Contract.WithdrawEther(&_Bnb.TransactOpts, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Bnb *BnbTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Bnb.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Bnb *BnbSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Bnb.Contract.Fallback(&_Bnb.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Bnb *BnbTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Bnb.Contract.Fallback(&_Bnb.TransactOpts, calldata)
}

// BnbBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Bnb contract.
type BnbBurnIterator struct {
	Event *BnbBurn // Event containing the contract specifics and raw log

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
func (it *BnbBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BnbBurn)
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
		it.Event = new(BnbBurn)
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
func (it *BnbBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BnbBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BnbBurn represents a Burn event raised by the Bnb contract.
type BnbBurn struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_Bnb *BnbFilterer) FilterBurn(opts *bind.FilterOpts, from []common.Address) (*BnbBurnIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Bnb.contract.FilterLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return &BnbBurnIterator{contract: _Bnb.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_Bnb *BnbFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *BnbBurn, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Bnb.contract.WatchLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BnbBurn)
				if err := _Bnb.contract.UnpackLog(event, "Burn", log); err != nil {
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
func (_Bnb *BnbFilterer) ParseBurn(log types.Log) (*BnbBurn, error) {
	event := new(BnbBurn)
	if err := _Bnb.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BnbFreezeIterator is returned from FilterFreeze and is used to iterate over the raw logs and unpacked data for Freeze events raised by the Bnb contract.
type BnbFreezeIterator struct {
	Event *BnbFreeze // Event containing the contract specifics and raw log

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
func (it *BnbFreezeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BnbFreeze)
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
		it.Event = new(BnbFreeze)
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
func (it *BnbFreezeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BnbFreezeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BnbFreeze represents a Freeze event raised by the Bnb contract.
type BnbFreeze struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterFreeze is a free log retrieval operation binding the contract event 0xf97a274face0b5517365ad396b1fdba6f68bd3135ef603e44272adba3af5a1e0.
//
// Solidity: event Freeze(address indexed from, uint256 value)
func (_Bnb *BnbFilterer) FilterFreeze(opts *bind.FilterOpts, from []common.Address) (*BnbFreezeIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Bnb.contract.FilterLogs(opts, "Freeze", fromRule)
	if err != nil {
		return nil, err
	}
	return &BnbFreezeIterator{contract: _Bnb.contract, event: "Freeze", logs: logs, sub: sub}, nil
}

// WatchFreeze is a free log subscription operation binding the contract event 0xf97a274face0b5517365ad396b1fdba6f68bd3135ef603e44272adba3af5a1e0.
//
// Solidity: event Freeze(address indexed from, uint256 value)
func (_Bnb *BnbFilterer) WatchFreeze(opts *bind.WatchOpts, sink chan<- *BnbFreeze, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Bnb.contract.WatchLogs(opts, "Freeze", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BnbFreeze)
				if err := _Bnb.contract.UnpackLog(event, "Freeze", log); err != nil {
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
func (_Bnb *BnbFilterer) ParseFreeze(log types.Log) (*BnbFreeze, error) {
	event := new(BnbFreeze)
	if err := _Bnb.contract.UnpackLog(event, "Freeze", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BnbTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Bnb contract.
type BnbTransferIterator struct {
	Event *BnbTransfer // Event containing the contract specifics and raw log

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
func (it *BnbTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BnbTransfer)
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
		it.Event = new(BnbTransfer)
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
func (it *BnbTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BnbTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BnbTransfer represents a Transfer event raised by the Bnb contract.
type BnbTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Bnb *BnbFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BnbTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bnb.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BnbTransferIterator{contract: _Bnb.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Bnb *BnbFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BnbTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bnb.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BnbTransfer)
				if err := _Bnb.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Bnb *BnbFilterer) ParseTransfer(log types.Log) (*BnbTransfer, error) {
	event := new(BnbTransfer)
	if err := _Bnb.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BnbUnfreezeIterator is returned from FilterUnfreeze and is used to iterate over the raw logs and unpacked data for Unfreeze events raised by the Bnb contract.
type BnbUnfreezeIterator struct {
	Event *BnbUnfreeze // Event containing the contract specifics and raw log

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
func (it *BnbUnfreezeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BnbUnfreeze)
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
		it.Event = new(BnbUnfreeze)
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
func (it *BnbUnfreezeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BnbUnfreezeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BnbUnfreeze represents a Unfreeze event raised by the Bnb contract.
type BnbUnfreeze struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnfreeze is a free log retrieval operation binding the contract event 0x2cfce4af01bcb9d6cf6c84ee1b7c491100b8695368264146a94d71e10a63083f.
//
// Solidity: event Unfreeze(address indexed from, uint256 value)
func (_Bnb *BnbFilterer) FilterUnfreeze(opts *bind.FilterOpts, from []common.Address) (*BnbUnfreezeIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Bnb.contract.FilterLogs(opts, "Unfreeze", fromRule)
	if err != nil {
		return nil, err
	}
	return &BnbUnfreezeIterator{contract: _Bnb.contract, event: "Unfreeze", logs: logs, sub: sub}, nil
}

// WatchUnfreeze is a free log subscription operation binding the contract event 0x2cfce4af01bcb9d6cf6c84ee1b7c491100b8695368264146a94d71e10a63083f.
//
// Solidity: event Unfreeze(address indexed from, uint256 value)
func (_Bnb *BnbFilterer) WatchUnfreeze(opts *bind.WatchOpts, sink chan<- *BnbUnfreeze, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Bnb.contract.WatchLogs(opts, "Unfreeze", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BnbUnfreeze)
				if err := _Bnb.contract.UnpackLog(event, "Unfreeze", log); err != nil {
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
func (_Bnb *BnbFilterer) ParseUnfreeze(log types.Log) (*BnbUnfreeze, error) {
	event := new(BnbUnfreeze)
	if err := _Bnb.contract.UnpackLog(event, "Unfreeze", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
