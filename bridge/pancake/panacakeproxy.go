// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pancakeproxy

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

// PancakeproxyMetaData contains all meta data concerning the Pancakeproxy contract.
var PancakeproxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPancakeRouter02\",\"name\":\"_pancake02\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ETH_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WBNB_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pancakeRouter02\",\"outputs\":[{\"internalType\":\"contractIPancakeRouter02\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"srcQty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isNative\",\"type\":\"bool\"}],\"name\":\"trade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isNative\",\"type\":\"bool\"}],\"name\":\"tradeTokensSupportingFee\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506040516118d53803806118d58339818101604052602081101561003357600080fd5b810190808051906020019092919050505080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ad5c46486040518163ffffffff1660e01b815260040160206040518083038186803b1580156100ed57600080fd5b505afa158015610101573d6000803e3d6000fd5b505050506040513d602081101561011757600080fd5b81019080805190602001909291905050506000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505061175e806101776000396000f3fe6080604052600436106100555760003560e01c806371e13d241461005a57806372e94bf61461012e578063c83d788b1461016f578063cd2239fb146101b0578063d49d5181146101f1578063f87485361461021c575b600080fd5b6100fb600480360360a081101561007057600080fd5b810190808035906020019064010000000081111561008d57600080fd5b82018360208201111561009f57600080fd5b803590602001918460208302840111640100000000831117156100c157600080fd5b90919293919293908035906020019092919080359060200190929190803590602001909291908035151590602001909291905050506102e6565b604051808373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b34801561013a57600080fd5b50610143610cae565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561017b57600080fd5b50610184610cb3565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156101bc57600080fd5b506101c5610cd9565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156101fd57600080fd5b50610206610cfd565b6040518082815260200191505060405180910390f35b6102b36004803603608081101561023257600080fd5b810190808035906020019064010000000081111561024f57600080fd5b82018360208201111561026157600080fd5b8035906020019184602083028401116401000000008311171561028357600080fd5b90919293919293908035906020019092919080359060200190929190803515159060200190929190505050610d21565b604051808373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b60008060008888905011610362576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f696e76616c69642070617468000000000000000000000000000000000000000081525060200191505060405180910390fd5b6000888890509050606060008034141561097c5760008b8b600081811061038557fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff169050898173ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e30600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1681526020019250505060206040518083038186803b15801561044c57600080fd5b505afa158015610460573d6000803e3d6000fd5b505050506040513d602081101561047657600080fd5b81019080805190602001909291905050501015610605578073ffffffffffffffffffffffffffffffffffffffff1663095ea7b3600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660006040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b15801561052157600080fd5b505af1158015610535573d6000803e3d6000fd5b505050508073ffffffffffffffffffffffffffffffffffffffff1663095ea7b3600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b1580156105ec57600080fd5b505af1158015610600573d6000803e3d6000fd5b505050505b866107c057600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166338ed17398b8b8f8f338e6040518763ffffffff1660e01b815260040180878152602001868152602001806020018473ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018281038252868682818152602001925060200280828437600081840152601f19601f820116905080830192505050975050505050505050600060405180830381600087803b1580156106e457600080fd5b505af11580156106f8573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250602081101561072257600080fd5b810190808051604051939291908464010000000082111561074257600080fd5b8382019150602082018581111561075857600080fd5b825186602082028301116401000000008211171561077557600080fd5b8083526020830192505050908051906020019060200280838360005b838110156107ac578082015181840152602081019050610791565b505050509050016040525050509250610976565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166318cbafe58b8b8f8f338e6040518763ffffffff1660e01b815260040180878152602001868152602001806020018473ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018281038252868682818152602001925060200280828437600081840152601f19601f820116905080830192505050975050505050505050600060405180830381600087803b15801561089a57600080fd5b505af11580156108ae573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525060208110156108d857600080fd5b81019080805160405193929190846401000000008211156108f857600080fd5b8382019150602082018581111561090e57600080fd5b825186602082028301116401000000008211171561092b57600080fd5b8083526020830192505050908051906020019060200280838360005b83811015610962578082015181840152602081019050610947565b505050509050016040525050509250600191505b50610b27565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637ff36ab58a8a8e8e338d6040518763ffffffff1660e01b815260040180868152602001806020018473ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018281038252868682818152602001925060200280828437600081840152601f19601f82011690508083019250505096505050505050506000604051808303818588803b158015610a4e57600080fd5b505af1158015610a62573d6000803e3d6000fd5b50505050506040513d6000823e3d601f19601f820116820180604052506020811015610a8d57600080fd5b8101908080516040519392919084640100000000821115610aad57600080fd5b83820191506020820185811115610ac357600080fd5b8251866020820283011164010000000082111715610ae057600080fd5b8083526020830192505050908051906020019060200280838360005b83811015610b17578082015181840152602081019050610afc565b5050505090500160405250505091505b600282511015610b9f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f696e76616c6964206f7574707574732076616c7565000000000000000000000081525060200191505060405180910390fd5b8782600184510381518110610bb057fe5b602002602001015110158015610bd957508882600081518110610bcf57fe5b6020026020010151145b610c4b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f657870656374656420616d6f756e74206e6f742072656163680000000000000081525060200191505060405180910390fd5b80610c81578a8a60018503818110610c5f57fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff16610c84565b60005b82600184510381518110610c9457fe5b602002602001015194509450505050965096945050505050565b600081565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81565b60008060008787905011610d9d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f696e76616c69642070617468000000000000000000000000000000000000000081525060200191505060405180910390fd5b60008787905090506000803414156112de57600089896000818110610dbe57fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff16905060008173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015610e4657600080fd5b505afa158015610e5a573d6000803e3d6000fd5b505050506040513d6020811015610e7057600080fd5b81019080805190602001909291905050509050808273ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e30600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1681526020019250505060206040518083038186803b158015610f2b57600080fd5b505afa158015610f3f573d6000803e3d6000fd5b505050506040513d6020811015610f5557600080fd5b810190808051906020019092919050505010156110e4578173ffffffffffffffffffffffffffffffffffffffff1663095ea7b3600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660006040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b15801561100057600080fd5b505af1158015611014573d6000803e3d6000fd5b505050508173ffffffffffffffffffffffffffffffffffffffff1663095ea7b3600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b1580156110cb57600080fd5b505af11580156110df573d6000803e3d6000fd5b505050505b866111e057600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635c11d795828b8e8e308e6040518763ffffffff1660e01b815260040180878152602001868152602001806020018473ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018281038252868682818152602001925060200280828437600081840152601f19601f820116905080830192505050975050505050505050600060405180830381600087803b1580156111c357600080fd5b505af11580156111d7573d6000803e3d6000fd5b505050506112d7565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663791ac947828b8e8e308e6040518763ffffffff1660e01b815260040180878152602001868152602001806020018473ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018281038252868682818152602001925060200280828437600081840152601f19601f820116905080830192505050975050505050505050600060405180830381600087803b1580156112ba57600080fd5b505af11580156112ce573d6000803e3d6000fd5b50505050600192505b50506113ca565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b6f9de9534898c8c308c6040518763ffffffff1660e01b815260040180868152602001806020018473ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018281038252868682818152602001925060200280828437600081840152601f19601f82011690508083019250505096505050505050506000604051808303818588803b1580156113b057600080fd5b505af11580156113c4573d6000803e3d6000fd5b50505050505b600081611402578989600185038181106113e057fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff16611405565b60005b90506000611412826114a8565b90508881101561148a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f657870656374656420616d6f756e74206e6f742072656163680000000000000081525060200191505060405180910390fd5b6114948282611590565b818195509550505050509550959350505050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156114e65747905061158b565b8173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561154d57600080fd5b505afa158015611561573d6000803e3d6000fd5b505050506040513d602081101561157757600080fd5b810190808051906020019092919050505090505b919050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561164b57804710156115d257600080fd5b60003373ffffffffffffffffffffffffffffffffffffffff168260405180600001905060006040518083038185875af1925050503d8060008114611632576040519150601f19603f3d011682016040523d82523d6000602084013e611637565b606091505b505090508061164557600080fd5b506116e6565b8173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b1580156116bc57600080fd5b505af11580156116d0573d6000803e3d6000fd5b505050506116dc6116ea565b6116e557600080fd5b5b5050565b600080600090503d6000811461170757602081146117105761171c565b6001915061171c565b60206000803e60005191505b5060008114159150509056fea26469706673582212207cc913d8aa0fd60fdf4b7eb355a35d24fce32fc6ee2ca7d52622236a61bcd0b964736f6c634300060c0033",
}

// PancakeproxyABI is the input ABI used to generate the binding from.
// Deprecated: Use PancakeproxyMetaData.ABI instead.
var PancakeproxyABI = PancakeproxyMetaData.ABI

// PancakeproxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PancakeproxyMetaData.Bin instead.
var PancakeproxyBin = PancakeproxyMetaData.Bin

// DeployPancakeproxy deploys a new Ethereum contract, binding an instance of Pancakeproxy to it.
func DeployPancakeproxy(auth *bind.TransactOpts, backend bind.ContractBackend, _pancake02 common.Address) (common.Address, *types.Transaction, *Pancakeproxy, error) {
	parsed, err := PancakeproxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PancakeproxyBin), backend, _pancake02)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pancakeproxy{PancakeproxyCaller: PancakeproxyCaller{contract: contract}, PancakeproxyTransactor: PancakeproxyTransactor{contract: contract}, PancakeproxyFilterer: PancakeproxyFilterer{contract: contract}}, nil
}

// Pancakeproxy is an auto generated Go binding around an Ethereum contract.
type Pancakeproxy struct {
	PancakeproxyCaller     // Read-only binding to the contract
	PancakeproxyTransactor // Write-only binding to the contract
	PancakeproxyFilterer   // Log filterer for contract events
}

// PancakeproxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type PancakeproxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeproxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PancakeproxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeproxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PancakeproxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeproxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PancakeproxySession struct {
	Contract     *Pancakeproxy     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PancakeproxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PancakeproxyCallerSession struct {
	Contract *PancakeproxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PancakeproxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PancakeproxyTransactorSession struct {
	Contract     *PancakeproxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PancakeproxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type PancakeproxyRaw struct {
	Contract *Pancakeproxy // Generic contract binding to access the raw methods on
}

// PancakeproxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PancakeproxyCallerRaw struct {
	Contract *PancakeproxyCaller // Generic read-only contract binding to access the raw methods on
}

// PancakeproxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PancakeproxyTransactorRaw struct {
	Contract *PancakeproxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPancakeproxy creates a new instance of Pancakeproxy, bound to a specific deployed contract.
func NewPancakeproxy(address common.Address, backend bind.ContractBackend) (*Pancakeproxy, error) {
	contract, err := bindPancakeproxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pancakeproxy{PancakeproxyCaller: PancakeproxyCaller{contract: contract}, PancakeproxyTransactor: PancakeproxyTransactor{contract: contract}, PancakeproxyFilterer: PancakeproxyFilterer{contract: contract}}, nil
}

// NewPancakeproxyCaller creates a new read-only instance of Pancakeproxy, bound to a specific deployed contract.
func NewPancakeproxyCaller(address common.Address, caller bind.ContractCaller) (*PancakeproxyCaller, error) {
	contract, err := bindPancakeproxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeproxyCaller{contract: contract}, nil
}

// NewPancakeproxyTransactor creates a new write-only instance of Pancakeproxy, bound to a specific deployed contract.
func NewPancakeproxyTransactor(address common.Address, transactor bind.ContractTransactor) (*PancakeproxyTransactor, error) {
	contract, err := bindPancakeproxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeproxyTransactor{contract: contract}, nil
}

// NewPancakeproxyFilterer creates a new log filterer instance of Pancakeproxy, bound to a specific deployed contract.
func NewPancakeproxyFilterer(address common.Address, filterer bind.ContractFilterer) (*PancakeproxyFilterer, error) {
	contract, err := bindPancakeproxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PancakeproxyFilterer{contract: contract}, nil
}

// bindPancakeproxy binds a generic wrapper to an already deployed contract.
func bindPancakeproxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PancakeproxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pancakeproxy *PancakeproxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pancakeproxy.Contract.PancakeproxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pancakeproxy *PancakeproxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pancakeproxy.Contract.PancakeproxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pancakeproxy *PancakeproxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pancakeproxy.Contract.PancakeproxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pancakeproxy *PancakeproxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pancakeproxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pancakeproxy *PancakeproxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pancakeproxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pancakeproxy *PancakeproxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pancakeproxy.Contract.contract.Transact(opts, method, params...)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() view returns(address)
func (_Pancakeproxy *PancakeproxyCaller) ETHCONTRACTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pancakeproxy.contract.Call(opts, &out, "ETH_CONTRACT_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() view returns(address)
func (_Pancakeproxy *PancakeproxySession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _Pancakeproxy.Contract.ETHCONTRACTADDRESS(&_Pancakeproxy.CallOpts)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() view returns(address)
func (_Pancakeproxy *PancakeproxyCallerSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _Pancakeproxy.Contract.ETHCONTRACTADDRESS(&_Pancakeproxy.CallOpts)
}

// MAX is a free data retrieval call binding the contract method 0xd49d5181.
//
// Solidity: function MAX() view returns(uint256)
func (_Pancakeproxy *PancakeproxyCaller) MAX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pancakeproxy.contract.Call(opts, &out, "MAX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAX is a free data retrieval call binding the contract method 0xd49d5181.
//
// Solidity: function MAX() view returns(uint256)
func (_Pancakeproxy *PancakeproxySession) MAX() (*big.Int, error) {
	return _Pancakeproxy.Contract.MAX(&_Pancakeproxy.CallOpts)
}

// MAX is a free data retrieval call binding the contract method 0xd49d5181.
//
// Solidity: function MAX() view returns(uint256)
func (_Pancakeproxy *PancakeproxyCallerSession) MAX() (*big.Int, error) {
	return _Pancakeproxy.Contract.MAX(&_Pancakeproxy.CallOpts)
}

// WBNBCONTRACTADDRESS is a free data retrieval call binding the contract method 0xcd2239fb.
//
// Solidity: function WBNB_CONTRACT_ADDRESS() view returns(address)
func (_Pancakeproxy *PancakeproxyCaller) WBNBCONTRACTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pancakeproxy.contract.Call(opts, &out, "WBNB_CONTRACT_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WBNBCONTRACTADDRESS is a free data retrieval call binding the contract method 0xcd2239fb.
//
// Solidity: function WBNB_CONTRACT_ADDRESS() view returns(address)
func (_Pancakeproxy *PancakeproxySession) WBNBCONTRACTADDRESS() (common.Address, error) {
	return _Pancakeproxy.Contract.WBNBCONTRACTADDRESS(&_Pancakeproxy.CallOpts)
}

// WBNBCONTRACTADDRESS is a free data retrieval call binding the contract method 0xcd2239fb.
//
// Solidity: function WBNB_CONTRACT_ADDRESS() view returns(address)
func (_Pancakeproxy *PancakeproxyCallerSession) WBNBCONTRACTADDRESS() (common.Address, error) {
	return _Pancakeproxy.Contract.WBNBCONTRACTADDRESS(&_Pancakeproxy.CallOpts)
}

// PancakeRouter02 is a free data retrieval call binding the contract method 0xc83d788b.
//
// Solidity: function pancakeRouter02() view returns(address)
func (_Pancakeproxy *PancakeproxyCaller) PancakeRouter02(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pancakeproxy.contract.Call(opts, &out, "pancakeRouter02")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PancakeRouter02 is a free data retrieval call binding the contract method 0xc83d788b.
//
// Solidity: function pancakeRouter02() view returns(address)
func (_Pancakeproxy *PancakeproxySession) PancakeRouter02() (common.Address, error) {
	return _Pancakeproxy.Contract.PancakeRouter02(&_Pancakeproxy.CallOpts)
}

// PancakeRouter02 is a free data retrieval call binding the contract method 0xc83d788b.
//
// Solidity: function pancakeRouter02() view returns(address)
func (_Pancakeproxy *PancakeproxyCallerSession) PancakeRouter02() (common.Address, error) {
	return _Pancakeproxy.Contract.PancakeRouter02(&_Pancakeproxy.CallOpts)
}

// Trade is a paid mutator transaction binding the contract method 0x71e13d24.
//
// Solidity: function trade(address[] path, uint256 srcQty, uint256 amountOutMin, uint256 deadline, bool isNative) payable returns(address, uint256)
func (_Pancakeproxy *PancakeproxyTransactor) Trade(opts *bind.TransactOpts, path []common.Address, srcQty *big.Int, amountOutMin *big.Int, deadline *big.Int, isNative bool) (*types.Transaction, error) {
	return _Pancakeproxy.contract.Transact(opts, "trade", path, srcQty, amountOutMin, deadline, isNative)
}

// Trade is a paid mutator transaction binding the contract method 0x71e13d24.
//
// Solidity: function trade(address[] path, uint256 srcQty, uint256 amountOutMin, uint256 deadline, bool isNative) payable returns(address, uint256)
func (_Pancakeproxy *PancakeproxySession) Trade(path []common.Address, srcQty *big.Int, amountOutMin *big.Int, deadline *big.Int, isNative bool) (*types.Transaction, error) {
	return _Pancakeproxy.Contract.Trade(&_Pancakeproxy.TransactOpts, path, srcQty, amountOutMin, deadline, isNative)
}

// Trade is a paid mutator transaction binding the contract method 0x71e13d24.
//
// Solidity: function trade(address[] path, uint256 srcQty, uint256 amountOutMin, uint256 deadline, bool isNative) payable returns(address, uint256)
func (_Pancakeproxy *PancakeproxyTransactorSession) Trade(path []common.Address, srcQty *big.Int, amountOutMin *big.Int, deadline *big.Int, isNative bool) (*types.Transaction, error) {
	return _Pancakeproxy.Contract.Trade(&_Pancakeproxy.TransactOpts, path, srcQty, amountOutMin, deadline, isNative)
}

// TradeTokensSupportingFee is a paid mutator transaction binding the contract method 0xf8748536.
//
// Solidity: function tradeTokensSupportingFee(address[] path, uint256 amountOutMin, uint256 deadline, bool isNative) payable returns(address, uint256)
func (_Pancakeproxy *PancakeproxyTransactor) TradeTokensSupportingFee(opts *bind.TransactOpts, path []common.Address, amountOutMin *big.Int, deadline *big.Int, isNative bool) (*types.Transaction, error) {
	return _Pancakeproxy.contract.Transact(opts, "tradeTokensSupportingFee", path, amountOutMin, deadline, isNative)
}

// TradeTokensSupportingFee is a paid mutator transaction binding the contract method 0xf8748536.
//
// Solidity: function tradeTokensSupportingFee(address[] path, uint256 amountOutMin, uint256 deadline, bool isNative) payable returns(address, uint256)
func (_Pancakeproxy *PancakeproxySession) TradeTokensSupportingFee(path []common.Address, amountOutMin *big.Int, deadline *big.Int, isNative bool) (*types.Transaction, error) {
	return _Pancakeproxy.Contract.TradeTokensSupportingFee(&_Pancakeproxy.TransactOpts, path, amountOutMin, deadline, isNative)
}

// TradeTokensSupportingFee is a paid mutator transaction binding the contract method 0xf8748536.
//
// Solidity: function tradeTokensSupportingFee(address[] path, uint256 amountOutMin, uint256 deadline, bool isNative) payable returns(address, uint256)
func (_Pancakeproxy *PancakeproxyTransactorSession) TradeTokensSupportingFee(path []common.Address, amountOutMin *big.Int, deadline *big.Int, isNative bool) (*types.Transaction, error) {
	return _Pancakeproxy.Contract.TradeTokensSupportingFee(&_Pancakeproxy.TransactOpts, path, amountOutMin, deadline, isNative)
}
