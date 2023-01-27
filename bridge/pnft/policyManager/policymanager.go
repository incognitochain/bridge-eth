// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package policymanager

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

// PolicymanagerMetaData contains all meta data concerning the Policymanager contract.
var PolicymanagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"policy\",\"type\":\"address\"}],\"name\":\"PolicyRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"policy\",\"type\":\"address\"}],\"name\":\"PolicyWhitelisted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"policy\",\"type\":\"address\"}],\"name\":\"addPolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"policy\",\"type\":\"address\"}],\"name\":\"isPolicyWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"policy\",\"type\":\"address\"}],\"name\":\"removePolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"viewCountWhitelistedPolicies\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cursor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"viewWhitelistedPolicies\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061002d61002261003260201b60201c565b61003a60201b60201c565b6100fe565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b610f048061010d6000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80638da5cb5b1161005b5780638da5cb5b14610114578063b84ef08114610132578063d6f0d88e1461014e578063f2fde38b1461016c57610088565b8063715018a61461008d578063747f54ae1461009757806378d38d82146100c8578063874516cd146100e4575b600080fd5b610095610188565b005b6100b160048036038101906100ac9190610938565b61019c565b6040516100bf929190610a77565b60405180910390f35b6100e260048036038101906100dd9190610ad3565b6102ca565b005b6100fe60048036038101906100f99190610ad3565b610380565b60405161010b9190610b1b565b60405180910390f35b61011c61039d565b6040516101299190610b45565b60405180910390f35b61014c60048036038101906101479190610ad3565b6103c6565b005b61015661047d565b6040516101639190610b60565b60405180910390f35b61018660048036038101906101819190610ad3565b61048e565b005b610190610511565b61019a600061058f565b565b6060600080839050846101af6001610653565b6101b99190610baa565b8111156101d857846101cb6001610653565b6101d59190610baa565b90505b60008167ffffffffffffffff8111156101f4576101f3610bde565b5b6040519080825280602002602001820160405280156102225781602001602082028036833780820191505090505b50905060005b828110156102af5761024f818861023f9190610c0d565b600161066890919063ffffffff16565b82828151811061026257610261610c41565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505080806102a790610c70565b915050610228565b508082876102bd9190610c0d565b9350935050509250929050565b6102d2610511565b6102e681600161068290919063ffffffff16565b610325576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161031c90610d15565b60405180910390fd5b6103398160016106b290919063ffffffff16565b508073ffffffffffffffffffffffffffffffffffffffff167f7d69543e56287cc6159e1cd00d200f2d86996b94ef9a170c4ec9a31daa6dd2ce60405160405180910390a250565b600061039682600161068290919063ffffffff16565b9050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6103ce610511565b6103e281600161068290919063ffffffff16565b15610422576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161041990610d81565b60405180910390fd5b6104368160016106e290919063ffffffff16565b508073ffffffffffffffffffffffffffffffffffffffff167f5b5592d50e950152eab424f0fde17ba8b36801c96694a656ca54c6ffd149980860405160405180910390a250565b60006104896001610653565b905090565b610496610511565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610505576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104fc90610e13565b60405180910390fd5b61050e8161058f565b50565b610519610712565b73ffffffffffffffffffffffffffffffffffffffff1661053761039d565b73ffffffffffffffffffffffffffffffffffffffff161461058d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161058490610e7f565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b60006106618260000161071a565b9050919050565b6000610677836000018361072b565b60001c905092915050565b60006106aa836000018373ffffffffffffffffffffffffffffffffffffffff1660001b610756565b905092915050565b60006106da836000018373ffffffffffffffffffffffffffffffffffffffff1660001b610779565b905092915050565b600061070a836000018373ffffffffffffffffffffffffffffffffffffffff1660001b61088d565b905092915050565b600033905090565b600081600001805490509050919050565b600082600001828154811061074357610742610c41565b5b9060005260206000200154905092915050565b600080836001016000848152602001908152602001600020541415905092915050565b600080836001016000848152602001908152602001600020549050600081146108815760006001826107ab9190610baa565b90506000600186600001805490506107c39190610baa565b90508181146108325760008660000182815481106107e4576107e3610c41565b5b906000526020600020015490508087600001848154811061080857610807610c41565b5b90600052602060002001819055508387600101600083815260200190815260200160002081905550505b8560000180548061084657610845610e9f565b5b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610887565b60009150505b92915050565b60006108998383610756565b6108f25782600001829080600181540180825580915050600190039060005260206000200160009091909190915055826000018054905083600101600084815260200190815260200160002081905550600190506108f7565b600090505b92915050565b600080fd5b6000819050919050565b61091581610902565b811461092057600080fd5b50565b6000813590506109328161090c565b92915050565b6000806040838503121561094f5761094e6108fd565b5b600061095d85828601610923565b925050602061096e85828601610923565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006109cf826109a4565b9050919050565b6109df816109c4565b82525050565b60006109f183836109d6565b60208301905092915050565b6000602082019050919050565b6000610a1582610978565b610a1f8185610983565b9350610a2a83610994565b8060005b83811015610a5b578151610a4288826109e5565b9750610a4d836109fd565b925050600181019050610a2e565b5085935050505092915050565b610a7181610902565b82525050565b60006040820190508181036000830152610a918185610a0a565b9050610aa06020830184610a68565b9392505050565b610ab0816109c4565b8114610abb57600080fd5b50565b600081359050610acd81610aa7565b92915050565b600060208284031215610ae957610ae86108fd565b5b6000610af784828501610abe565b91505092915050565b60008115159050919050565b610b1581610b00565b82525050565b6000602082019050610b306000830184610b0c565b92915050565b610b3f816109c4565b82525050565b6000602082019050610b5a6000830184610b36565b92915050565b6000602082019050610b756000830184610a68565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610bb582610902565b9150610bc083610902565b9250828203905081811115610bd857610bd7610b7b565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000610c1882610902565b9150610c2383610902565b9250828201905080821115610c3b57610c3a610b7b565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000610c7b82610902565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610cad57610cac610b7b565b5b600182019050919050565b600082825260208201905092915050565b7f4e6f742077686974656c69737465640000000000000000000000000000000000600082015250565b6000610cff600f83610cb8565b9150610d0a82610cc9565b602082019050919050565b60006020820190508181036000830152610d2e81610cf2565b9050919050565b7f416c72656164792077686974656c697374656400000000000000000000000000600082015250565b6000610d6b601383610cb8565b9150610d7682610d35565b602082019050919050565b60006020820190508181036000830152610d9a81610d5e565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000610dfd602683610cb8565b9150610e0882610da1565b604082019050919050565b60006020820190508181036000830152610e2c81610df0565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000610e69602083610cb8565b9150610e7482610e33565b602082019050919050565b60006020820190508181036000830152610e9881610e5c565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea264697066735822122010e62540dcc55a662e1422e433009a05685928c687aa619c5947b75ef11cb46e64736f6c63430008110033",
}

// PolicymanagerABI is the input ABI used to generate the binding from.
// Deprecated: Use PolicymanagerMetaData.ABI instead.
var PolicymanagerABI = PolicymanagerMetaData.ABI

// PolicymanagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PolicymanagerMetaData.Bin instead.
var PolicymanagerBin = PolicymanagerMetaData.Bin

// DeployPolicymanager deploys a new Ethereum contract, binding an instance of Policymanager to it.
func DeployPolicymanager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Policymanager, error) {
	parsed, err := PolicymanagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PolicymanagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Policymanager{PolicymanagerCaller: PolicymanagerCaller{contract: contract}, PolicymanagerTransactor: PolicymanagerTransactor{contract: contract}, PolicymanagerFilterer: PolicymanagerFilterer{contract: contract}}, nil
}

// Policymanager is an auto generated Go binding around an Ethereum contract.
type Policymanager struct {
	PolicymanagerCaller     // Read-only binding to the contract
	PolicymanagerTransactor // Write-only binding to the contract
	PolicymanagerFilterer   // Log filterer for contract events
}

// PolicymanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolicymanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolicymanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolicymanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolicymanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolicymanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolicymanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolicymanagerSession struct {
	Contract     *Policymanager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PolicymanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolicymanagerCallerSession struct {
	Contract *PolicymanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PolicymanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolicymanagerTransactorSession struct {
	Contract     *PolicymanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PolicymanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolicymanagerRaw struct {
	Contract *Policymanager // Generic contract binding to access the raw methods on
}

// PolicymanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolicymanagerCallerRaw struct {
	Contract *PolicymanagerCaller // Generic read-only contract binding to access the raw methods on
}

// PolicymanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolicymanagerTransactorRaw struct {
	Contract *PolicymanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolicymanager creates a new instance of Policymanager, bound to a specific deployed contract.
func NewPolicymanager(address common.Address, backend bind.ContractBackend) (*Policymanager, error) {
	contract, err := bindPolicymanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Policymanager{PolicymanagerCaller: PolicymanagerCaller{contract: contract}, PolicymanagerTransactor: PolicymanagerTransactor{contract: contract}, PolicymanagerFilterer: PolicymanagerFilterer{contract: contract}}, nil
}

// NewPolicymanagerCaller creates a new read-only instance of Policymanager, bound to a specific deployed contract.
func NewPolicymanagerCaller(address common.Address, caller bind.ContractCaller) (*PolicymanagerCaller, error) {
	contract, err := bindPolicymanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolicymanagerCaller{contract: contract}, nil
}

// NewPolicymanagerTransactor creates a new write-only instance of Policymanager, bound to a specific deployed contract.
func NewPolicymanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*PolicymanagerTransactor, error) {
	contract, err := bindPolicymanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolicymanagerTransactor{contract: contract}, nil
}

// NewPolicymanagerFilterer creates a new log filterer instance of Policymanager, bound to a specific deployed contract.
func NewPolicymanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*PolicymanagerFilterer, error) {
	contract, err := bindPolicymanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolicymanagerFilterer{contract: contract}, nil
}

// bindPolicymanager binds a generic wrapper to an already deployed contract.
func bindPolicymanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PolicymanagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Policymanager *PolicymanagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Policymanager.Contract.PolicymanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Policymanager *PolicymanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Policymanager.Contract.PolicymanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Policymanager *PolicymanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Policymanager.Contract.PolicymanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Policymanager *PolicymanagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Policymanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Policymanager *PolicymanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Policymanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Policymanager *PolicymanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Policymanager.Contract.contract.Transact(opts, method, params...)
}

// IsPolicyWhitelisted is a free data retrieval call binding the contract method 0x874516cd.
//
// Solidity: function isPolicyWhitelisted(address policy) view returns(bool)
func (_Policymanager *PolicymanagerCaller) IsPolicyWhitelisted(opts *bind.CallOpts, policy common.Address) (bool, error) {
	var out []interface{}
	err := _Policymanager.contract.Call(opts, &out, "isPolicyWhitelisted", policy)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPolicyWhitelisted is a free data retrieval call binding the contract method 0x874516cd.
//
// Solidity: function isPolicyWhitelisted(address policy) view returns(bool)
func (_Policymanager *PolicymanagerSession) IsPolicyWhitelisted(policy common.Address) (bool, error) {
	return _Policymanager.Contract.IsPolicyWhitelisted(&_Policymanager.CallOpts, policy)
}

// IsPolicyWhitelisted is a free data retrieval call binding the contract method 0x874516cd.
//
// Solidity: function isPolicyWhitelisted(address policy) view returns(bool)
func (_Policymanager *PolicymanagerCallerSession) IsPolicyWhitelisted(policy common.Address) (bool, error) {
	return _Policymanager.Contract.IsPolicyWhitelisted(&_Policymanager.CallOpts, policy)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Policymanager *PolicymanagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Policymanager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Policymanager *PolicymanagerSession) Owner() (common.Address, error) {
	return _Policymanager.Contract.Owner(&_Policymanager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Policymanager *PolicymanagerCallerSession) Owner() (common.Address, error) {
	return _Policymanager.Contract.Owner(&_Policymanager.CallOpts)
}

// ViewCountWhitelistedPolicies is a free data retrieval call binding the contract method 0xd6f0d88e.
//
// Solidity: function viewCountWhitelistedPolicies() view returns(uint256)
func (_Policymanager *PolicymanagerCaller) ViewCountWhitelistedPolicies(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Policymanager.contract.Call(opts, &out, "viewCountWhitelistedPolicies")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ViewCountWhitelistedPolicies is a free data retrieval call binding the contract method 0xd6f0d88e.
//
// Solidity: function viewCountWhitelistedPolicies() view returns(uint256)
func (_Policymanager *PolicymanagerSession) ViewCountWhitelistedPolicies() (*big.Int, error) {
	return _Policymanager.Contract.ViewCountWhitelistedPolicies(&_Policymanager.CallOpts)
}

// ViewCountWhitelistedPolicies is a free data retrieval call binding the contract method 0xd6f0d88e.
//
// Solidity: function viewCountWhitelistedPolicies() view returns(uint256)
func (_Policymanager *PolicymanagerCallerSession) ViewCountWhitelistedPolicies() (*big.Int, error) {
	return _Policymanager.Contract.ViewCountWhitelistedPolicies(&_Policymanager.CallOpts)
}

// ViewWhitelistedPolicies is a free data retrieval call binding the contract method 0x747f54ae.
//
// Solidity: function viewWhitelistedPolicies(uint256 cursor, uint256 size) view returns(address[], uint256)
func (_Policymanager *PolicymanagerCaller) ViewWhitelistedPolicies(opts *bind.CallOpts, cursor *big.Int, size *big.Int) ([]common.Address, *big.Int, error) {
	var out []interface{}
	err := _Policymanager.contract.Call(opts, &out, "viewWhitelistedPolicies", cursor, size)

	if err != nil {
		return *new([]common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// ViewWhitelistedPolicies is a free data retrieval call binding the contract method 0x747f54ae.
//
// Solidity: function viewWhitelistedPolicies(uint256 cursor, uint256 size) view returns(address[], uint256)
func (_Policymanager *PolicymanagerSession) ViewWhitelistedPolicies(cursor *big.Int, size *big.Int) ([]common.Address, *big.Int, error) {
	return _Policymanager.Contract.ViewWhitelistedPolicies(&_Policymanager.CallOpts, cursor, size)
}

// ViewWhitelistedPolicies is a free data retrieval call binding the contract method 0x747f54ae.
//
// Solidity: function viewWhitelistedPolicies(uint256 cursor, uint256 size) view returns(address[], uint256)
func (_Policymanager *PolicymanagerCallerSession) ViewWhitelistedPolicies(cursor *big.Int, size *big.Int) ([]common.Address, *big.Int, error) {
	return _Policymanager.Contract.ViewWhitelistedPolicies(&_Policymanager.CallOpts, cursor, size)
}

// AddPolicy is a paid mutator transaction binding the contract method 0xb84ef081.
//
// Solidity: function addPolicy(address policy) returns()
func (_Policymanager *PolicymanagerTransactor) AddPolicy(opts *bind.TransactOpts, policy common.Address) (*types.Transaction, error) {
	return _Policymanager.contract.Transact(opts, "addPolicy", policy)
}

// AddPolicy is a paid mutator transaction binding the contract method 0xb84ef081.
//
// Solidity: function addPolicy(address policy) returns()
func (_Policymanager *PolicymanagerSession) AddPolicy(policy common.Address) (*types.Transaction, error) {
	return _Policymanager.Contract.AddPolicy(&_Policymanager.TransactOpts, policy)
}

// AddPolicy is a paid mutator transaction binding the contract method 0xb84ef081.
//
// Solidity: function addPolicy(address policy) returns()
func (_Policymanager *PolicymanagerTransactorSession) AddPolicy(policy common.Address) (*types.Transaction, error) {
	return _Policymanager.Contract.AddPolicy(&_Policymanager.TransactOpts, policy)
}

// RemovePolicy is a paid mutator transaction binding the contract method 0x78d38d82.
//
// Solidity: function removePolicy(address policy) returns()
func (_Policymanager *PolicymanagerTransactor) RemovePolicy(opts *bind.TransactOpts, policy common.Address) (*types.Transaction, error) {
	return _Policymanager.contract.Transact(opts, "removePolicy", policy)
}

// RemovePolicy is a paid mutator transaction binding the contract method 0x78d38d82.
//
// Solidity: function removePolicy(address policy) returns()
func (_Policymanager *PolicymanagerSession) RemovePolicy(policy common.Address) (*types.Transaction, error) {
	return _Policymanager.Contract.RemovePolicy(&_Policymanager.TransactOpts, policy)
}

// RemovePolicy is a paid mutator transaction binding the contract method 0x78d38d82.
//
// Solidity: function removePolicy(address policy) returns()
func (_Policymanager *PolicymanagerTransactorSession) RemovePolicy(policy common.Address) (*types.Transaction, error) {
	return _Policymanager.Contract.RemovePolicy(&_Policymanager.TransactOpts, policy)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Policymanager *PolicymanagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Policymanager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Policymanager *PolicymanagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Policymanager.Contract.RenounceOwnership(&_Policymanager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Policymanager *PolicymanagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Policymanager.Contract.RenounceOwnership(&_Policymanager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Policymanager *PolicymanagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Policymanager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Policymanager *PolicymanagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Policymanager.Contract.TransferOwnership(&_Policymanager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Policymanager *PolicymanagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Policymanager.Contract.TransferOwnership(&_Policymanager.TransactOpts, newOwner)
}

// PolicymanagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Policymanager contract.
type PolicymanagerOwnershipTransferredIterator struct {
	Event *PolicymanagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PolicymanagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolicymanagerOwnershipTransferred)
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
		it.Event = new(PolicymanagerOwnershipTransferred)
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
func (it *PolicymanagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolicymanagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolicymanagerOwnershipTransferred represents a OwnershipTransferred event raised by the Policymanager contract.
type PolicymanagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Policymanager *PolicymanagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PolicymanagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Policymanager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PolicymanagerOwnershipTransferredIterator{contract: _Policymanager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Policymanager *PolicymanagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PolicymanagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Policymanager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolicymanagerOwnershipTransferred)
				if err := _Policymanager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Policymanager *PolicymanagerFilterer) ParseOwnershipTransferred(log types.Log) (*PolicymanagerOwnershipTransferred, error) {
	event := new(PolicymanagerOwnershipTransferred)
	if err := _Policymanager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolicymanagerPolicyRemovedIterator is returned from FilterPolicyRemoved and is used to iterate over the raw logs and unpacked data for PolicyRemoved events raised by the Policymanager contract.
type PolicymanagerPolicyRemovedIterator struct {
	Event *PolicymanagerPolicyRemoved // Event containing the contract specifics and raw log

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
func (it *PolicymanagerPolicyRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolicymanagerPolicyRemoved)
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
		it.Event = new(PolicymanagerPolicyRemoved)
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
func (it *PolicymanagerPolicyRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolicymanagerPolicyRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolicymanagerPolicyRemoved represents a PolicyRemoved event raised by the Policymanager contract.
type PolicymanagerPolicyRemoved struct {
	Policy common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPolicyRemoved is a free log retrieval operation binding the contract event 0x7d69543e56287cc6159e1cd00d200f2d86996b94ef9a170c4ec9a31daa6dd2ce.
//
// Solidity: event PolicyRemoved(address indexed policy)
func (_Policymanager *PolicymanagerFilterer) FilterPolicyRemoved(opts *bind.FilterOpts, policy []common.Address) (*PolicymanagerPolicyRemovedIterator, error) {

	var policyRule []interface{}
	for _, policyItem := range policy {
		policyRule = append(policyRule, policyItem)
	}

	logs, sub, err := _Policymanager.contract.FilterLogs(opts, "PolicyRemoved", policyRule)
	if err != nil {
		return nil, err
	}
	return &PolicymanagerPolicyRemovedIterator{contract: _Policymanager.contract, event: "PolicyRemoved", logs: logs, sub: sub}, nil
}

// WatchPolicyRemoved is a free log subscription operation binding the contract event 0x7d69543e56287cc6159e1cd00d200f2d86996b94ef9a170c4ec9a31daa6dd2ce.
//
// Solidity: event PolicyRemoved(address indexed policy)
func (_Policymanager *PolicymanagerFilterer) WatchPolicyRemoved(opts *bind.WatchOpts, sink chan<- *PolicymanagerPolicyRemoved, policy []common.Address) (event.Subscription, error) {

	var policyRule []interface{}
	for _, policyItem := range policy {
		policyRule = append(policyRule, policyItem)
	}

	logs, sub, err := _Policymanager.contract.WatchLogs(opts, "PolicyRemoved", policyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolicymanagerPolicyRemoved)
				if err := _Policymanager.contract.UnpackLog(event, "PolicyRemoved", log); err != nil {
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

// ParsePolicyRemoved is a log parse operation binding the contract event 0x7d69543e56287cc6159e1cd00d200f2d86996b94ef9a170c4ec9a31daa6dd2ce.
//
// Solidity: event PolicyRemoved(address indexed policy)
func (_Policymanager *PolicymanagerFilterer) ParsePolicyRemoved(log types.Log) (*PolicymanagerPolicyRemoved, error) {
	event := new(PolicymanagerPolicyRemoved)
	if err := _Policymanager.contract.UnpackLog(event, "PolicyRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolicymanagerPolicyWhitelistedIterator is returned from FilterPolicyWhitelisted and is used to iterate over the raw logs and unpacked data for PolicyWhitelisted events raised by the Policymanager contract.
type PolicymanagerPolicyWhitelistedIterator struct {
	Event *PolicymanagerPolicyWhitelisted // Event containing the contract specifics and raw log

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
func (it *PolicymanagerPolicyWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolicymanagerPolicyWhitelisted)
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
		it.Event = new(PolicymanagerPolicyWhitelisted)
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
func (it *PolicymanagerPolicyWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolicymanagerPolicyWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolicymanagerPolicyWhitelisted represents a PolicyWhitelisted event raised by the Policymanager contract.
type PolicymanagerPolicyWhitelisted struct {
	Policy common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPolicyWhitelisted is a free log retrieval operation binding the contract event 0x5b5592d50e950152eab424f0fde17ba8b36801c96694a656ca54c6ffd1499808.
//
// Solidity: event PolicyWhitelisted(address indexed policy)
func (_Policymanager *PolicymanagerFilterer) FilterPolicyWhitelisted(opts *bind.FilterOpts, policy []common.Address) (*PolicymanagerPolicyWhitelistedIterator, error) {

	var policyRule []interface{}
	for _, policyItem := range policy {
		policyRule = append(policyRule, policyItem)
	}

	logs, sub, err := _Policymanager.contract.FilterLogs(opts, "PolicyWhitelisted", policyRule)
	if err != nil {
		return nil, err
	}
	return &PolicymanagerPolicyWhitelistedIterator{contract: _Policymanager.contract, event: "PolicyWhitelisted", logs: logs, sub: sub}, nil
}

// WatchPolicyWhitelisted is a free log subscription operation binding the contract event 0x5b5592d50e950152eab424f0fde17ba8b36801c96694a656ca54c6ffd1499808.
//
// Solidity: event PolicyWhitelisted(address indexed policy)
func (_Policymanager *PolicymanagerFilterer) WatchPolicyWhitelisted(opts *bind.WatchOpts, sink chan<- *PolicymanagerPolicyWhitelisted, policy []common.Address) (event.Subscription, error) {

	var policyRule []interface{}
	for _, policyItem := range policy {
		policyRule = append(policyRule, policyItem)
	}

	logs, sub, err := _Policymanager.contract.WatchLogs(opts, "PolicyWhitelisted", policyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolicymanagerPolicyWhitelisted)
				if err := _Policymanager.contract.UnpackLog(event, "PolicyWhitelisted", log); err != nil {
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

// ParsePolicyWhitelisted is a log parse operation binding the contract event 0x5b5592d50e950152eab424f0fde17ba8b36801c96694a656ca54c6ffd1499808.
//
// Solidity: event PolicyWhitelisted(address indexed policy)
func (_Policymanager *PolicymanagerFilterer) ParsePolicyWhitelisted(log types.Log) (*PolicymanagerPolicyWhitelisted, error) {
	event := new(PolicymanagerPolicyWhitelisted)
	if err := _Policymanager.contract.UnpackLog(event, "PolicyWhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
