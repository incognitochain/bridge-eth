// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package executiondelegate

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

// ExecutiondelegateMetaData contains all meta data concerning the Executiondelegate contract.
var ExecutiondelegateMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"ApproveContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"DenyContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"GrantApproval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"RevokeApproval\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"approveContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"contracts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"denyContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"grantApproval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokeApproval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"revokedApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferERC721Unsafe\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061001a3361001f565b61006f565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b610bd98061007e6000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c80638da5cb5b1161008c578063b7e2077e11610066578063b7e2077e1461019d578063ca72da8e146101b0578063da3e8ce4146101c3578063f2fde38b146101d657600080fd5b80638da5cb5b1461017257806390d02b3c1461018d578063a8034df11461019557600080fd5b806307f7aafb146100d45780634a3e3a1f146100e957806369dc9ff314610121578063715018a61461014457806374a9402e1461014c578063789f93f61461015f575b600080fd5b6100e76100e23660046109d2565b6101e9565b005b61010c6100f73660046109d2565b60026020526000908152604090205460ff1681565b60405190151581526020015b60405180910390f35b61010c61012f3660046109d2565b60016020526000908152604090205460ff1681565b6100e7610240565b6100e761015a3660046109ed565b610254565b6100e761016d366004610a42565b61034b565b6000546040516001600160a01b039091168152602001610118565b6100e7610424565b6100e7610466565b6100e76101ab3660046109d2565b6104a5565b6100e76101be366004610a42565b6104f6565b6100e76101d1366004610a42565b61059b565b6100e76101e43660046109d2565b6106cd565b6101f1610746565b6001600160a01b0381166000818152600160208190526040808320805460ff1916909217909155517f283ffe02a14663588cf87ba17adbc21c9ce0d0cdb15655926bf2b987af3075fe9190a250565b610248610746565b61025260006107a0565b565b3360009081526001602052604090205460ff1661028c5760405162461bcd60e51b815260040161028390610a8d565b60405180910390fd5b6001600160a01b03841660009081526002602052604090205460ff16156102c55760405162461bcd60e51b815260040161028390610ad7565b604051637921219560e11b81526001600160a01b0385811660048301528481166024830152604482018490526064820183905260a06084830152600060a483015286169063f242432a9060c401600060405180830381600087803b15801561032c57600080fd5b505af1158015610340573d6000803e3d6000fd5b505050505050505050565b3360009081526001602052604090205460ff1661037a5760405162461bcd60e51b815260040161028390610a8d565b6001600160a01b03831660009081526002602052604090205460ff16156103b35760405162461bcd60e51b815260040161028390610ad7565b604051632142170760e11b81526001600160a01b0384811660048301528381166024830152604482018390528516906342842e0e906064015b600060405180830381600087803b15801561040657600080fd5b505af115801561041a573d6000803e3d6000fd5b5050505050505050565b33600081815260026020526040808220805460ff19166001179055517fdddeac663983b1e35153215a4578fecbb5921d12e660b3c4259aa7d9dbb9709f9190a2565b33600081815260026020526040808220805460ff19169055517f120d91a0121c2d5a7ce9638fce4bd262d4b443568fce40f681f50dca814a629a9190a2565b6104ad610746565b6001600160a01b038116600081815260016020526040808220805460ff19169055517f2b35b0a030b4f4cef0a9e8d01828235bb82a11ec4e37c11bd6d8770d9aafb17c9190a250565b3360009081526001602052604090205460ff166105255760405162461bcd60e51b815260040161028390610a8d565b6001600160a01b03831660009081526002602052604090205460ff161561055e5760405162461bcd60e51b815260040161028390610ad7565b6040516323b872dd60e01b81526001600160a01b0384811660048301528381166024830152604482018390528516906323b872dd906064016103ec565b3360009081526001602052604090205460ff166105ca5760405162461bcd60e51b815260040161028390610a8d565b6001600160a01b03831660009081526002602052604090205460ff16156106035760405162461bcd60e51b815260040161028390610ad7565b604080516001600160a01b038581166024830152848116604483015260648083018590528351808403909101815260849092019092526020810180516001600160e01b03166323b872dd60e01b17905290600090610663908716836107f0565b8051909150156106c557808060200190518101906106819190610b0e565b6106c55760405162461bcd60e51b8152602060048201526015602482015274115490cc8c081d1c985b9cd9995c8819985a5b1959605a1b6044820152606401610283565b505050505050565b6106d5610746565b6001600160a01b03811661073a5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610283565b610743816107a0565b50565b6000546001600160a01b031633146102525760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610283565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6060610834838360006040518060400160405280601e81526020017f416464726573733a206c6f772d6c6576656c2063616c6c206661696c6564000081525061083b565b9392505050565b60608247101561089c5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610283565b600080866001600160a01b031685876040516108b89190610b54565b60006040518083038185875af1925050503d80600081146108f5576040519150601f19603f3d011682016040523d82523d6000602084013e6108fa565b606091505b509150915061090b87838387610918565b925050505b949350505050565b60608315610987578251600003610980576001600160a01b0385163b6109805760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610283565b5081610910565b610910838381511561099c5781518083602001fd5b8060405162461bcd60e51b81526004016102839190610b70565b80356001600160a01b03811681146109cd57600080fd5b919050565b6000602082840312156109e457600080fd5b610834826109b6565b600080600080600060a08688031215610a0557600080fd5b610a0e866109b6565b9450610a1c602087016109b6565b9350610a2a604087016109b6565b94979396509394606081013594506080013592915050565b60008060008060808587031215610a5857600080fd5b610a61856109b6565b9350610a6f602086016109b6565b9250610a7d604086016109b6565b9396929550929360600135925050565b6020808252602a908201527f436f6e7472616374206973206e6f7420617070726f76656420746f206d616b65604082015269207472616e736665727360b01b606082015260800190565b60208082526019908201527f5573657220686173207265766f6b656420617070726f76616c00000000000000604082015260600190565b600060208284031215610b2057600080fd5b8151801515811461083457600080fd5b60005b83811015610b4b578181015183820152602001610b33565b50506000910152565b60008251610b66818460208701610b30565b9190910192915050565b6020815260008251806020840152610b8f816040850160208701610b30565b601f01601f1916919091016040019291505056fea264697066735822122065fbfb1c1dcccb1086f249ca02cb9407958694b6e90973ae97913b96e487470564736f6c63430008110033",
}

// ExecutiondelegateABI is the input ABI used to generate the binding from.
// Deprecated: Use ExecutiondelegateMetaData.ABI instead.
var ExecutiondelegateABI = ExecutiondelegateMetaData.ABI

// ExecutiondelegateBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExecutiondelegateMetaData.Bin instead.
var ExecutiondelegateBin = ExecutiondelegateMetaData.Bin

// DeployExecutiondelegate deploys a new Ethereum contract, binding an instance of Executiondelegate to it.
func DeployExecutiondelegate(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Executiondelegate, error) {
	parsed, err := ExecutiondelegateMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExecutiondelegateBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Executiondelegate{ExecutiondelegateCaller: ExecutiondelegateCaller{contract: contract}, ExecutiondelegateTransactor: ExecutiondelegateTransactor{contract: contract}, ExecutiondelegateFilterer: ExecutiondelegateFilterer{contract: contract}}, nil
}

// Executiondelegate is an auto generated Go binding around an Ethereum contract.
type Executiondelegate struct {
	ExecutiondelegateCaller     // Read-only binding to the contract
	ExecutiondelegateTransactor // Write-only binding to the contract
	ExecutiondelegateFilterer   // Log filterer for contract events
}

// ExecutiondelegateCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExecutiondelegateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutiondelegateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExecutiondelegateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutiondelegateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExecutiondelegateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutiondelegateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExecutiondelegateSession struct {
	Contract     *Executiondelegate // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ExecutiondelegateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExecutiondelegateCallerSession struct {
	Contract *ExecutiondelegateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ExecutiondelegateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExecutiondelegateTransactorSession struct {
	Contract     *ExecutiondelegateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ExecutiondelegateRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExecutiondelegateRaw struct {
	Contract *Executiondelegate // Generic contract binding to access the raw methods on
}

// ExecutiondelegateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExecutiondelegateCallerRaw struct {
	Contract *ExecutiondelegateCaller // Generic read-only contract binding to access the raw methods on
}

// ExecutiondelegateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExecutiondelegateTransactorRaw struct {
	Contract *ExecutiondelegateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExecutiondelegate creates a new instance of Executiondelegate, bound to a specific deployed contract.
func NewExecutiondelegate(address common.Address, backend bind.ContractBackend) (*Executiondelegate, error) {
	contract, err := bindExecutiondelegate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Executiondelegate{ExecutiondelegateCaller: ExecutiondelegateCaller{contract: contract}, ExecutiondelegateTransactor: ExecutiondelegateTransactor{contract: contract}, ExecutiondelegateFilterer: ExecutiondelegateFilterer{contract: contract}}, nil
}

// NewExecutiondelegateCaller creates a new read-only instance of Executiondelegate, bound to a specific deployed contract.
func NewExecutiondelegateCaller(address common.Address, caller bind.ContractCaller) (*ExecutiondelegateCaller, error) {
	contract, err := bindExecutiondelegate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutiondelegateCaller{contract: contract}, nil
}

// NewExecutiondelegateTransactor creates a new write-only instance of Executiondelegate, bound to a specific deployed contract.
func NewExecutiondelegateTransactor(address common.Address, transactor bind.ContractTransactor) (*ExecutiondelegateTransactor, error) {
	contract, err := bindExecutiondelegate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutiondelegateTransactor{contract: contract}, nil
}

// NewExecutiondelegateFilterer creates a new log filterer instance of Executiondelegate, bound to a specific deployed contract.
func NewExecutiondelegateFilterer(address common.Address, filterer bind.ContractFilterer) (*ExecutiondelegateFilterer, error) {
	contract, err := bindExecutiondelegate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExecutiondelegateFilterer{contract: contract}, nil
}

// bindExecutiondelegate binds a generic wrapper to an already deployed contract.
func bindExecutiondelegate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExecutiondelegateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Executiondelegate *ExecutiondelegateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Executiondelegate.Contract.ExecutiondelegateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Executiondelegate *ExecutiondelegateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Executiondelegate.Contract.ExecutiondelegateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Executiondelegate *ExecutiondelegateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Executiondelegate.Contract.ExecutiondelegateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Executiondelegate *ExecutiondelegateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Executiondelegate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Executiondelegate *ExecutiondelegateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Executiondelegate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Executiondelegate *ExecutiondelegateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Executiondelegate.Contract.contract.Transact(opts, method, params...)
}

// Contracts is a free data retrieval call binding the contract method 0x69dc9ff3.
//
// Solidity: function contracts(address ) view returns(bool)
func (_Executiondelegate *ExecutiondelegateCaller) Contracts(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Executiondelegate.contract.Call(opts, &out, "contracts", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Contracts is a free data retrieval call binding the contract method 0x69dc9ff3.
//
// Solidity: function contracts(address ) view returns(bool)
func (_Executiondelegate *ExecutiondelegateSession) Contracts(arg0 common.Address) (bool, error) {
	return _Executiondelegate.Contract.Contracts(&_Executiondelegate.CallOpts, arg0)
}

// Contracts is a free data retrieval call binding the contract method 0x69dc9ff3.
//
// Solidity: function contracts(address ) view returns(bool)
func (_Executiondelegate *ExecutiondelegateCallerSession) Contracts(arg0 common.Address) (bool, error) {
	return _Executiondelegate.Contract.Contracts(&_Executiondelegate.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Executiondelegate *ExecutiondelegateCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Executiondelegate.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Executiondelegate *ExecutiondelegateSession) Owner() (common.Address, error) {
	return _Executiondelegate.Contract.Owner(&_Executiondelegate.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Executiondelegate *ExecutiondelegateCallerSession) Owner() (common.Address, error) {
	return _Executiondelegate.Contract.Owner(&_Executiondelegate.CallOpts)
}

// RevokedApproval is a free data retrieval call binding the contract method 0x4a3e3a1f.
//
// Solidity: function revokedApproval(address ) view returns(bool)
func (_Executiondelegate *ExecutiondelegateCaller) RevokedApproval(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Executiondelegate.contract.Call(opts, &out, "revokedApproval", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RevokedApproval is a free data retrieval call binding the contract method 0x4a3e3a1f.
//
// Solidity: function revokedApproval(address ) view returns(bool)
func (_Executiondelegate *ExecutiondelegateSession) RevokedApproval(arg0 common.Address) (bool, error) {
	return _Executiondelegate.Contract.RevokedApproval(&_Executiondelegate.CallOpts, arg0)
}

// RevokedApproval is a free data retrieval call binding the contract method 0x4a3e3a1f.
//
// Solidity: function revokedApproval(address ) view returns(bool)
func (_Executiondelegate *ExecutiondelegateCallerSession) RevokedApproval(arg0 common.Address) (bool, error) {
	return _Executiondelegate.Contract.RevokedApproval(&_Executiondelegate.CallOpts, arg0)
}

// ApproveContract is a paid mutator transaction binding the contract method 0x07f7aafb.
//
// Solidity: function approveContract(address _contract) returns()
func (_Executiondelegate *ExecutiondelegateTransactor) ApproveContract(opts *bind.TransactOpts, _contract common.Address) (*types.Transaction, error) {
	return _Executiondelegate.contract.Transact(opts, "approveContract", _contract)
}

// ApproveContract is a paid mutator transaction binding the contract method 0x07f7aafb.
//
// Solidity: function approveContract(address _contract) returns()
func (_Executiondelegate *ExecutiondelegateSession) ApproveContract(_contract common.Address) (*types.Transaction, error) {
	return _Executiondelegate.Contract.ApproveContract(&_Executiondelegate.TransactOpts, _contract)
}

// ApproveContract is a paid mutator transaction binding the contract method 0x07f7aafb.
//
// Solidity: function approveContract(address _contract) returns()
func (_Executiondelegate *ExecutiondelegateTransactorSession) ApproveContract(_contract common.Address) (*types.Transaction, error) {
	return _Executiondelegate.Contract.ApproveContract(&_Executiondelegate.TransactOpts, _contract)
}

// DenyContract is a paid mutator transaction binding the contract method 0xb7e2077e.
//
// Solidity: function denyContract(address _contract) returns()
func (_Executiondelegate *ExecutiondelegateTransactor) DenyContract(opts *bind.TransactOpts, _contract common.Address) (*types.Transaction, error) {
	return _Executiondelegate.contract.Transact(opts, "denyContract", _contract)
}

// DenyContract is a paid mutator transaction binding the contract method 0xb7e2077e.
//
// Solidity: function denyContract(address _contract) returns()
func (_Executiondelegate *ExecutiondelegateSession) DenyContract(_contract common.Address) (*types.Transaction, error) {
	return _Executiondelegate.Contract.DenyContract(&_Executiondelegate.TransactOpts, _contract)
}

// DenyContract is a paid mutator transaction binding the contract method 0xb7e2077e.
//
// Solidity: function denyContract(address _contract) returns()
func (_Executiondelegate *ExecutiondelegateTransactorSession) DenyContract(_contract common.Address) (*types.Transaction, error) {
	return _Executiondelegate.Contract.DenyContract(&_Executiondelegate.TransactOpts, _contract)
}

// GrantApproval is a paid mutator transaction binding the contract method 0xa8034df1.
//
// Solidity: function grantApproval() returns()
func (_Executiondelegate *ExecutiondelegateTransactor) GrantApproval(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Executiondelegate.contract.Transact(opts, "grantApproval")
}

// GrantApproval is a paid mutator transaction binding the contract method 0xa8034df1.
//
// Solidity: function grantApproval() returns()
func (_Executiondelegate *ExecutiondelegateSession) GrantApproval() (*types.Transaction, error) {
	return _Executiondelegate.Contract.GrantApproval(&_Executiondelegate.TransactOpts)
}

// GrantApproval is a paid mutator transaction binding the contract method 0xa8034df1.
//
// Solidity: function grantApproval() returns()
func (_Executiondelegate *ExecutiondelegateTransactorSession) GrantApproval() (*types.Transaction, error) {
	return _Executiondelegate.Contract.GrantApproval(&_Executiondelegate.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Executiondelegate *ExecutiondelegateTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Executiondelegate.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Executiondelegate *ExecutiondelegateSession) RenounceOwnership() (*types.Transaction, error) {
	return _Executiondelegate.Contract.RenounceOwnership(&_Executiondelegate.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Executiondelegate *ExecutiondelegateTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Executiondelegate.Contract.RenounceOwnership(&_Executiondelegate.TransactOpts)
}

// RevokeApproval is a paid mutator transaction binding the contract method 0x90d02b3c.
//
// Solidity: function revokeApproval() returns()
func (_Executiondelegate *ExecutiondelegateTransactor) RevokeApproval(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Executiondelegate.contract.Transact(opts, "revokeApproval")
}

// RevokeApproval is a paid mutator transaction binding the contract method 0x90d02b3c.
//
// Solidity: function revokeApproval() returns()
func (_Executiondelegate *ExecutiondelegateSession) RevokeApproval() (*types.Transaction, error) {
	return _Executiondelegate.Contract.RevokeApproval(&_Executiondelegate.TransactOpts)
}

// RevokeApproval is a paid mutator transaction binding the contract method 0x90d02b3c.
//
// Solidity: function revokeApproval() returns()
func (_Executiondelegate *ExecutiondelegateTransactorSession) RevokeApproval() (*types.Transaction, error) {
	return _Executiondelegate.Contract.RevokeApproval(&_Executiondelegate.TransactOpts)
}

// TransferERC1155 is a paid mutator transaction binding the contract method 0x74a9402e.
//
// Solidity: function transferERC1155(address collection, address from, address to, uint256 tokenId, uint256 amount) returns()
func (_Executiondelegate *ExecutiondelegateTransactor) TransferERC1155(opts *bind.TransactOpts, collection common.Address, from common.Address, to common.Address, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Executiondelegate.contract.Transact(opts, "transferERC1155", collection, from, to, tokenId, amount)
}

// TransferERC1155 is a paid mutator transaction binding the contract method 0x74a9402e.
//
// Solidity: function transferERC1155(address collection, address from, address to, uint256 tokenId, uint256 amount) returns()
func (_Executiondelegate *ExecutiondelegateSession) TransferERC1155(collection common.Address, from common.Address, to common.Address, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Executiondelegate.Contract.TransferERC1155(&_Executiondelegate.TransactOpts, collection, from, to, tokenId, amount)
}

// TransferERC1155 is a paid mutator transaction binding the contract method 0x74a9402e.
//
// Solidity: function transferERC1155(address collection, address from, address to, uint256 tokenId, uint256 amount) returns()
func (_Executiondelegate *ExecutiondelegateTransactorSession) TransferERC1155(collection common.Address, from common.Address, to common.Address, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Executiondelegate.Contract.TransferERC1155(&_Executiondelegate.TransactOpts, collection, from, to, tokenId, amount)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0xda3e8ce4.
//
// Solidity: function transferERC20(address token, address from, address to, uint256 amount) returns()
func (_Executiondelegate *ExecutiondelegateTransactor) TransferERC20(opts *bind.TransactOpts, token common.Address, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Executiondelegate.contract.Transact(opts, "transferERC20", token, from, to, amount)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0xda3e8ce4.
//
// Solidity: function transferERC20(address token, address from, address to, uint256 amount) returns()
func (_Executiondelegate *ExecutiondelegateSession) TransferERC20(token common.Address, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Executiondelegate.Contract.TransferERC20(&_Executiondelegate.TransactOpts, token, from, to, amount)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0xda3e8ce4.
//
// Solidity: function transferERC20(address token, address from, address to, uint256 amount) returns()
func (_Executiondelegate *ExecutiondelegateTransactorSession) TransferERC20(token common.Address, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Executiondelegate.Contract.TransferERC20(&_Executiondelegate.TransactOpts, token, from, to, amount)
}

// TransferERC721 is a paid mutator transaction binding the contract method 0x789f93f6.
//
// Solidity: function transferERC721(address collection, address from, address to, uint256 tokenId) returns()
func (_Executiondelegate *ExecutiondelegateTransactor) TransferERC721(opts *bind.TransactOpts, collection common.Address, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Executiondelegate.contract.Transact(opts, "transferERC721", collection, from, to, tokenId)
}

// TransferERC721 is a paid mutator transaction binding the contract method 0x789f93f6.
//
// Solidity: function transferERC721(address collection, address from, address to, uint256 tokenId) returns()
func (_Executiondelegate *ExecutiondelegateSession) TransferERC721(collection common.Address, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Executiondelegate.Contract.TransferERC721(&_Executiondelegate.TransactOpts, collection, from, to, tokenId)
}

// TransferERC721 is a paid mutator transaction binding the contract method 0x789f93f6.
//
// Solidity: function transferERC721(address collection, address from, address to, uint256 tokenId) returns()
func (_Executiondelegate *ExecutiondelegateTransactorSession) TransferERC721(collection common.Address, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Executiondelegate.Contract.TransferERC721(&_Executiondelegate.TransactOpts, collection, from, to, tokenId)
}

// TransferERC721Unsafe is a paid mutator transaction binding the contract method 0xca72da8e.
//
// Solidity: function transferERC721Unsafe(address collection, address from, address to, uint256 tokenId) returns()
func (_Executiondelegate *ExecutiondelegateTransactor) TransferERC721Unsafe(opts *bind.TransactOpts, collection common.Address, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Executiondelegate.contract.Transact(opts, "transferERC721Unsafe", collection, from, to, tokenId)
}

// TransferERC721Unsafe is a paid mutator transaction binding the contract method 0xca72da8e.
//
// Solidity: function transferERC721Unsafe(address collection, address from, address to, uint256 tokenId) returns()
func (_Executiondelegate *ExecutiondelegateSession) TransferERC721Unsafe(collection common.Address, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Executiondelegate.Contract.TransferERC721Unsafe(&_Executiondelegate.TransactOpts, collection, from, to, tokenId)
}

// TransferERC721Unsafe is a paid mutator transaction binding the contract method 0xca72da8e.
//
// Solidity: function transferERC721Unsafe(address collection, address from, address to, uint256 tokenId) returns()
func (_Executiondelegate *ExecutiondelegateTransactorSession) TransferERC721Unsafe(collection common.Address, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Executiondelegate.Contract.TransferERC721Unsafe(&_Executiondelegate.TransactOpts, collection, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Executiondelegate *ExecutiondelegateTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Executiondelegate.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Executiondelegate *ExecutiondelegateSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Executiondelegate.Contract.TransferOwnership(&_Executiondelegate.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Executiondelegate *ExecutiondelegateTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Executiondelegate.Contract.TransferOwnership(&_Executiondelegate.TransactOpts, newOwner)
}

// ExecutiondelegateApproveContractIterator is returned from FilterApproveContract and is used to iterate over the raw logs and unpacked data for ApproveContract events raised by the Executiondelegate contract.
type ExecutiondelegateApproveContractIterator struct {
	Event *ExecutiondelegateApproveContract // Event containing the contract specifics and raw log

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
func (it *ExecutiondelegateApproveContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutiondelegateApproveContract)
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
		it.Event = new(ExecutiondelegateApproveContract)
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
func (it *ExecutiondelegateApproveContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutiondelegateApproveContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutiondelegateApproveContract represents a ApproveContract event raised by the Executiondelegate contract.
type ExecutiondelegateApproveContract struct {
	Contract common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproveContract is a free log retrieval operation binding the contract event 0x283ffe02a14663588cf87ba17adbc21c9ce0d0cdb15655926bf2b987af3075fe.
//
// Solidity: event ApproveContract(address indexed _contract)
func (_Executiondelegate *ExecutiondelegateFilterer) FilterApproveContract(opts *bind.FilterOpts, _contract []common.Address) (*ExecutiondelegateApproveContractIterator, error) {

	var _contractRule []interface{}
	for _, _contractItem := range _contract {
		_contractRule = append(_contractRule, _contractItem)
	}

	logs, sub, err := _Executiondelegate.contract.FilterLogs(opts, "ApproveContract", _contractRule)
	if err != nil {
		return nil, err
	}
	return &ExecutiondelegateApproveContractIterator{contract: _Executiondelegate.contract, event: "ApproveContract", logs: logs, sub: sub}, nil
}

// WatchApproveContract is a free log subscription operation binding the contract event 0x283ffe02a14663588cf87ba17adbc21c9ce0d0cdb15655926bf2b987af3075fe.
//
// Solidity: event ApproveContract(address indexed _contract)
func (_Executiondelegate *ExecutiondelegateFilterer) WatchApproveContract(opts *bind.WatchOpts, sink chan<- *ExecutiondelegateApproveContract, _contract []common.Address) (event.Subscription, error) {

	var _contractRule []interface{}
	for _, _contractItem := range _contract {
		_contractRule = append(_contractRule, _contractItem)
	}

	logs, sub, err := _Executiondelegate.contract.WatchLogs(opts, "ApproveContract", _contractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutiondelegateApproveContract)
				if err := _Executiondelegate.contract.UnpackLog(event, "ApproveContract", log); err != nil {
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

// ParseApproveContract is a log parse operation binding the contract event 0x283ffe02a14663588cf87ba17adbc21c9ce0d0cdb15655926bf2b987af3075fe.
//
// Solidity: event ApproveContract(address indexed _contract)
func (_Executiondelegate *ExecutiondelegateFilterer) ParseApproveContract(log types.Log) (*ExecutiondelegateApproveContract, error) {
	event := new(ExecutiondelegateApproveContract)
	if err := _Executiondelegate.contract.UnpackLog(event, "ApproveContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutiondelegateDenyContractIterator is returned from FilterDenyContract and is used to iterate over the raw logs and unpacked data for DenyContract events raised by the Executiondelegate contract.
type ExecutiondelegateDenyContractIterator struct {
	Event *ExecutiondelegateDenyContract // Event containing the contract specifics and raw log

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
func (it *ExecutiondelegateDenyContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutiondelegateDenyContract)
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
		it.Event = new(ExecutiondelegateDenyContract)
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
func (it *ExecutiondelegateDenyContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutiondelegateDenyContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutiondelegateDenyContract represents a DenyContract event raised by the Executiondelegate contract.
type ExecutiondelegateDenyContract struct {
	Contract common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDenyContract is a free log retrieval operation binding the contract event 0x2b35b0a030b4f4cef0a9e8d01828235bb82a11ec4e37c11bd6d8770d9aafb17c.
//
// Solidity: event DenyContract(address indexed _contract)
func (_Executiondelegate *ExecutiondelegateFilterer) FilterDenyContract(opts *bind.FilterOpts, _contract []common.Address) (*ExecutiondelegateDenyContractIterator, error) {

	var _contractRule []interface{}
	for _, _contractItem := range _contract {
		_contractRule = append(_contractRule, _contractItem)
	}

	logs, sub, err := _Executiondelegate.contract.FilterLogs(opts, "DenyContract", _contractRule)
	if err != nil {
		return nil, err
	}
	return &ExecutiondelegateDenyContractIterator{contract: _Executiondelegate.contract, event: "DenyContract", logs: logs, sub: sub}, nil
}

// WatchDenyContract is a free log subscription operation binding the contract event 0x2b35b0a030b4f4cef0a9e8d01828235bb82a11ec4e37c11bd6d8770d9aafb17c.
//
// Solidity: event DenyContract(address indexed _contract)
func (_Executiondelegate *ExecutiondelegateFilterer) WatchDenyContract(opts *bind.WatchOpts, sink chan<- *ExecutiondelegateDenyContract, _contract []common.Address) (event.Subscription, error) {

	var _contractRule []interface{}
	for _, _contractItem := range _contract {
		_contractRule = append(_contractRule, _contractItem)
	}

	logs, sub, err := _Executiondelegate.contract.WatchLogs(opts, "DenyContract", _contractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutiondelegateDenyContract)
				if err := _Executiondelegate.contract.UnpackLog(event, "DenyContract", log); err != nil {
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

// ParseDenyContract is a log parse operation binding the contract event 0x2b35b0a030b4f4cef0a9e8d01828235bb82a11ec4e37c11bd6d8770d9aafb17c.
//
// Solidity: event DenyContract(address indexed _contract)
func (_Executiondelegate *ExecutiondelegateFilterer) ParseDenyContract(log types.Log) (*ExecutiondelegateDenyContract, error) {
	event := new(ExecutiondelegateDenyContract)
	if err := _Executiondelegate.contract.UnpackLog(event, "DenyContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutiondelegateGrantApprovalIterator is returned from FilterGrantApproval and is used to iterate over the raw logs and unpacked data for GrantApproval events raised by the Executiondelegate contract.
type ExecutiondelegateGrantApprovalIterator struct {
	Event *ExecutiondelegateGrantApproval // Event containing the contract specifics and raw log

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
func (it *ExecutiondelegateGrantApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutiondelegateGrantApproval)
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
		it.Event = new(ExecutiondelegateGrantApproval)
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
func (it *ExecutiondelegateGrantApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutiondelegateGrantApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutiondelegateGrantApproval represents a GrantApproval event raised by the Executiondelegate contract.
type ExecutiondelegateGrantApproval struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterGrantApproval is a free log retrieval operation binding the contract event 0x120d91a0121c2d5a7ce9638fce4bd262d4b443568fce40f681f50dca814a629a.
//
// Solidity: event GrantApproval(address indexed user)
func (_Executiondelegate *ExecutiondelegateFilterer) FilterGrantApproval(opts *bind.FilterOpts, user []common.Address) (*ExecutiondelegateGrantApprovalIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Executiondelegate.contract.FilterLogs(opts, "GrantApproval", userRule)
	if err != nil {
		return nil, err
	}
	return &ExecutiondelegateGrantApprovalIterator{contract: _Executiondelegate.contract, event: "GrantApproval", logs: logs, sub: sub}, nil
}

// WatchGrantApproval is a free log subscription operation binding the contract event 0x120d91a0121c2d5a7ce9638fce4bd262d4b443568fce40f681f50dca814a629a.
//
// Solidity: event GrantApproval(address indexed user)
func (_Executiondelegate *ExecutiondelegateFilterer) WatchGrantApproval(opts *bind.WatchOpts, sink chan<- *ExecutiondelegateGrantApproval, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Executiondelegate.contract.WatchLogs(opts, "GrantApproval", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutiondelegateGrantApproval)
				if err := _Executiondelegate.contract.UnpackLog(event, "GrantApproval", log); err != nil {
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

// ParseGrantApproval is a log parse operation binding the contract event 0x120d91a0121c2d5a7ce9638fce4bd262d4b443568fce40f681f50dca814a629a.
//
// Solidity: event GrantApproval(address indexed user)
func (_Executiondelegate *ExecutiondelegateFilterer) ParseGrantApproval(log types.Log) (*ExecutiondelegateGrantApproval, error) {
	event := new(ExecutiondelegateGrantApproval)
	if err := _Executiondelegate.contract.UnpackLog(event, "GrantApproval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutiondelegateOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Executiondelegate contract.
type ExecutiondelegateOwnershipTransferredIterator struct {
	Event *ExecutiondelegateOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ExecutiondelegateOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutiondelegateOwnershipTransferred)
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
		it.Event = new(ExecutiondelegateOwnershipTransferred)
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
func (it *ExecutiondelegateOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutiondelegateOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutiondelegateOwnershipTransferred represents a OwnershipTransferred event raised by the Executiondelegate contract.
type ExecutiondelegateOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Executiondelegate *ExecutiondelegateFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ExecutiondelegateOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Executiondelegate.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ExecutiondelegateOwnershipTransferredIterator{contract: _Executiondelegate.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Executiondelegate *ExecutiondelegateFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ExecutiondelegateOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Executiondelegate.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutiondelegateOwnershipTransferred)
				if err := _Executiondelegate.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Executiondelegate *ExecutiondelegateFilterer) ParseOwnershipTransferred(log types.Log) (*ExecutiondelegateOwnershipTransferred, error) {
	event := new(ExecutiondelegateOwnershipTransferred)
	if err := _Executiondelegate.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutiondelegateRevokeApprovalIterator is returned from FilterRevokeApproval and is used to iterate over the raw logs and unpacked data for RevokeApproval events raised by the Executiondelegate contract.
type ExecutiondelegateRevokeApprovalIterator struct {
	Event *ExecutiondelegateRevokeApproval // Event containing the contract specifics and raw log

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
func (it *ExecutiondelegateRevokeApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutiondelegateRevokeApproval)
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
		it.Event = new(ExecutiondelegateRevokeApproval)
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
func (it *ExecutiondelegateRevokeApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutiondelegateRevokeApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutiondelegateRevokeApproval represents a RevokeApproval event raised by the Executiondelegate contract.
type ExecutiondelegateRevokeApproval struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRevokeApproval is a free log retrieval operation binding the contract event 0xdddeac663983b1e35153215a4578fecbb5921d12e660b3c4259aa7d9dbb9709f.
//
// Solidity: event RevokeApproval(address indexed user)
func (_Executiondelegate *ExecutiondelegateFilterer) FilterRevokeApproval(opts *bind.FilterOpts, user []common.Address) (*ExecutiondelegateRevokeApprovalIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Executiondelegate.contract.FilterLogs(opts, "RevokeApproval", userRule)
	if err != nil {
		return nil, err
	}
	return &ExecutiondelegateRevokeApprovalIterator{contract: _Executiondelegate.contract, event: "RevokeApproval", logs: logs, sub: sub}, nil
}

// WatchRevokeApproval is a free log subscription operation binding the contract event 0xdddeac663983b1e35153215a4578fecbb5921d12e660b3c4259aa7d9dbb9709f.
//
// Solidity: event RevokeApproval(address indexed user)
func (_Executiondelegate *ExecutiondelegateFilterer) WatchRevokeApproval(opts *bind.WatchOpts, sink chan<- *ExecutiondelegateRevokeApproval, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Executiondelegate.contract.WatchLogs(opts, "RevokeApproval", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutiondelegateRevokeApproval)
				if err := _Executiondelegate.contract.UnpackLog(event, "RevokeApproval", log); err != nil {
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

// ParseRevokeApproval is a log parse operation binding the contract event 0xdddeac663983b1e35153215a4578fecbb5921d12e660b3c4259aa7d9dbb9709f.
//
// Solidity: event RevokeApproval(address indexed user)
func (_Executiondelegate *ExecutiondelegateFilterer) ParseRevokeApproval(log types.Log) (*ExecutiondelegateRevokeApproval, error) {
	event := new(ExecutiondelegateRevokeApproval)
	if err := _Executiondelegate.contract.UnpackLog(event, "RevokeApproval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
