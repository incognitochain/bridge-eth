// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vault

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

// VaultHelperPreSignData is an auto generated low-level Go binding around an user-defined struct.
type VaultHelperPreSignData struct {
	Prefix    uint8
	Token     common.Address
	Timestamp []byte
	Amount    *big.Int
}

// VaultHelperABI is the input ABI used to generate the binding from.
const VaultHelperABI = "[{\"inputs\":[{\"internalType\":\"enumVaultHelper.Prefix\",\"name\":\"prefix\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"timestamp\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"_buildPreSignData\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumVaultHelper.Prefix\",\"name\":\"prefix\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"timestamp\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structVaultHelper.PreSignData\",\"name\":\"psd\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"recipientToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"exchangeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"name\":\"_buildSignExecute\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumVaultHelper.Prefix\",\"name\":\"prefix\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"timestamp\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structVaultHelper.PreSignData\",\"name\":\"psd\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"}],\"name\":\"_buildSignRequestWithdraw\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// VaultHelperFuncSigs maps the 4-byte function signature to its string representation.
var VaultHelperFuncSigs = map[string]string{
	"741d7a50": "_buildPreSignData(uint8,address,bytes,uint256)",
	"41f59b6f": "_buildSignExecute((uint8,address,bytes,uint256),address,address,bytes)",
	"3cf6e1ca": "_buildSignRequestWithdraw((uint8,address,bytes,uint256),string)",
}

// VaultHelperBin is the compiled bytecode used for deploying new contracts.
var VaultHelperBin = "0x608060405234801561001057600080fd5b50610573806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80633cf6e1ca1461004657806341f59b6f1461006f578063741d7a5014610082575b600080fd5b61005961005436600461033e565b610097565b604051610066919061046b565b60405180910390f35b61005961007d3660046102ad565b6100c8565b610095610090366004610244565b6100ff565b005b6060808484846040516020016100af939291906104ce565b60408051808303601f1901815291905295945050505050565b60608086868686866040516020016100e4959493929190610485565b60408051808303601f19018152919052979650505050505050565b5050505050565b80356001600160a01b038116811461011d57600080fd5b92915050565b60008083601f840112610134578182fd5b50813567ffffffffffffffff81111561014b578182fd5b60208301915083602082850101111561016357600080fd5b9250929050565b80356002811061011d57600080fd5b60006080828403121561018a578081fd5b61019460806104fe565b90506101a0838361016a565b815260206101b084828501610106565b81830152604083013567ffffffffffffffff808211156101cf57600080fd5b818501915085601f8301126101e357600080fd5b8135818111156101f257600080fd5b610204601f8201601f191685016104fe565b9150808252868482850101111561021a57600080fd5b80848401858401376000848284010152508060408501525050506060820135606082015292915050565b60008060008060006080868803121561025b578081fd5b610265878761016a565b94506102748760208801610106565b9350604086013567ffffffffffffffff81111561028f578182fd5b61029b88828901610123565b96999598509660600135949350505050565b6000806000806000608086880312156102c4578081fd5b853567ffffffffffffffff808211156102db578283fd5b6102e789838a01610179565b9650602088013591506102f982610525565b90945060408701359061030b82610525565b90935060608701359080821115610320578283fd5b5061032d88828901610123565b969995985093965092949392505050565b600080600060408486031215610352578283fd5b833567ffffffffffffffff80821115610369578485fd5b61037587838801610179565b9450602086013591508082111561038a578384fd5b5061039786828701610123565b9497909650939450505050565b60008284528282602086013780602084860101526020601f19601f85011685010190509392505050565b60008151808452815b818110156103f3576020818501810151868301820152016103d7565b818111156104045782602083870101525b50601f01601f19169290920160200192915050565b600081516002811061042757fe5b83526020828101516001600160a01b031690840152604080830151608091850182905290610457908501826103ce565b606093840151949093019390935250919050565b60006020825261047e60208301846103ce565b9392505050565b6000608082526104986080830188610419565b6001600160a01b0387811660208501528616604084015282810360608401526104c28185876103a4565b98975050505050505050565b6000604082526104e16040830186610419565b82810360208401526104f48185876103a4565b9695505050505050565b60405181810167ffffffffffffffff8111828210171561051d57600080fd5b604052919050565b6001600160a01b038116811461053a57600080fd5b5056fea26469706673582212202ed63d797852ecf48c69cc6b84dd7747ed8328795a87ab4bd71d1e76f7aa14c764736f6c634300060c0033"

// DeployVaultHelper deploys a new Ethereum contract, binding an instance of VaultHelper to it.
func DeployVaultHelper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VaultHelper, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultHelperABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VaultHelperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VaultHelper{VaultHelperCaller: VaultHelperCaller{contract: contract}, VaultHelperTransactor: VaultHelperTransactor{contract: contract}, VaultHelperFilterer: VaultHelperFilterer{contract: contract}}, nil
}

// VaultHelper is an auto generated Go binding around an Ethereum contract.
type VaultHelper struct {
	VaultHelperCaller     // Read-only binding to the contract
	VaultHelperTransactor // Write-only binding to the contract
	VaultHelperFilterer   // Log filterer for contract events
}

// VaultHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultHelperSession struct {
	Contract     *VaultHelper      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultHelperCallerSession struct {
	Contract *VaultHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// VaultHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultHelperTransactorSession struct {
	Contract     *VaultHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// VaultHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultHelperRaw struct {
	Contract *VaultHelper // Generic contract binding to access the raw methods on
}

// VaultHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultHelperCallerRaw struct {
	Contract *VaultHelperCaller // Generic read-only contract binding to access the raw methods on
}

// VaultHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultHelperTransactorRaw struct {
	Contract *VaultHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVaultHelper creates a new instance of VaultHelper, bound to a specific deployed contract.
func NewVaultHelper(address common.Address, backend bind.ContractBackend) (*VaultHelper, error) {
	contract, err := bindVaultHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VaultHelper{VaultHelperCaller: VaultHelperCaller{contract: contract}, VaultHelperTransactor: VaultHelperTransactor{contract: contract}, VaultHelperFilterer: VaultHelperFilterer{contract: contract}}, nil
}

// NewVaultHelperCaller creates a new read-only instance of VaultHelper, bound to a specific deployed contract.
func NewVaultHelperCaller(address common.Address, caller bind.ContractCaller) (*VaultHelperCaller, error) {
	contract, err := bindVaultHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultHelperCaller{contract: contract}, nil
}

// NewVaultHelperTransactor creates a new write-only instance of VaultHelper, bound to a specific deployed contract.
func NewVaultHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultHelperTransactor, error) {
	contract, err := bindVaultHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultHelperTransactor{contract: contract}, nil
}

// NewVaultHelperFilterer creates a new log filterer instance of VaultHelper, bound to a specific deployed contract.
func NewVaultHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultHelperFilterer, error) {
	contract, err := bindVaultHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultHelperFilterer{contract: contract}, nil
}

// bindVaultHelper binds a generic wrapper to an already deployed contract.
func bindVaultHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultHelperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultHelper *VaultHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultHelper.Contract.VaultHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultHelper *VaultHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultHelper.Contract.VaultHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultHelper *VaultHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultHelper.Contract.VaultHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultHelper *VaultHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultHelper *VaultHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultHelper *VaultHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultHelper.Contract.contract.Transact(opts, method, params...)
}

// BuildPreSignData is a free data retrieval call binding the contract method 0x741d7a50.
//
// Solidity: function _buildPreSignData(uint8 prefix, address token, bytes timestamp, uint256 amount) pure returns()
func (_VaultHelper *VaultHelperCaller) BuildPreSignData(opts *bind.CallOpts, prefix uint8, token common.Address, timestamp []byte, amount *big.Int) error {
	var out []interface{}
	err := _VaultHelper.contract.Call(opts, &out, "_buildPreSignData", prefix, token, timestamp, amount)

	if err != nil {
		return err
	}

	return err

}

// BuildPreSignData is a free data retrieval call binding the contract method 0x741d7a50.
//
// Solidity: function _buildPreSignData(uint8 prefix, address token, bytes timestamp, uint256 amount) pure returns()
func (_VaultHelper *VaultHelperSession) BuildPreSignData(prefix uint8, token common.Address, timestamp []byte, amount *big.Int) error {
	return _VaultHelper.Contract.BuildPreSignData(&_VaultHelper.CallOpts, prefix, token, timestamp, amount)
}

// BuildPreSignData is a free data retrieval call binding the contract method 0x741d7a50.
//
// Solidity: function _buildPreSignData(uint8 prefix, address token, bytes timestamp, uint256 amount) pure returns()
func (_VaultHelper *VaultHelperCallerSession) BuildPreSignData(prefix uint8, token common.Address, timestamp []byte, amount *big.Int) error {
	return _VaultHelper.Contract.BuildPreSignData(&_VaultHelper.CallOpts, prefix, token, timestamp, amount)
}

// BuildSignExecute is a free data retrieval call binding the contract method 0x41f59b6f.
//
// Solidity: function _buildSignExecute((uint8,address,bytes,uint256) psd, address recipientToken, address exchangeAddress, bytes callData) pure returns(bytes)
func (_VaultHelper *VaultHelperCaller) BuildSignExecute(opts *bind.CallOpts, psd VaultHelperPreSignData, recipientToken common.Address, exchangeAddress common.Address, callData []byte) ([]byte, error) {
	var out []interface{}
	err := _VaultHelper.contract.Call(opts, &out, "_buildSignExecute", psd, recipientToken, exchangeAddress, callData)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// BuildSignExecute is a free data retrieval call binding the contract method 0x41f59b6f.
//
// Solidity: function _buildSignExecute((uint8,address,bytes,uint256) psd, address recipientToken, address exchangeAddress, bytes callData) pure returns(bytes)
func (_VaultHelper *VaultHelperSession) BuildSignExecute(psd VaultHelperPreSignData, recipientToken common.Address, exchangeAddress common.Address, callData []byte) ([]byte, error) {
	return _VaultHelper.Contract.BuildSignExecute(&_VaultHelper.CallOpts, psd, recipientToken, exchangeAddress, callData)
}

// BuildSignExecute is a free data retrieval call binding the contract method 0x41f59b6f.
//
// Solidity: function _buildSignExecute((uint8,address,bytes,uint256) psd, address recipientToken, address exchangeAddress, bytes callData) pure returns(bytes)
func (_VaultHelper *VaultHelperCallerSession) BuildSignExecute(psd VaultHelperPreSignData, recipientToken common.Address, exchangeAddress common.Address, callData []byte) ([]byte, error) {
	return _VaultHelper.Contract.BuildSignExecute(&_VaultHelper.CallOpts, psd, recipientToken, exchangeAddress, callData)
}

// BuildSignRequestWithdraw is a free data retrieval call binding the contract method 0x3cf6e1ca.
//
// Solidity: function _buildSignRequestWithdraw((uint8,address,bytes,uint256) psd, string incognitoAddress) pure returns(bytes)
func (_VaultHelper *VaultHelperCaller) BuildSignRequestWithdraw(opts *bind.CallOpts, psd VaultHelperPreSignData, incognitoAddress string) ([]byte, error) {
	var out []interface{}
	err := _VaultHelper.contract.Call(opts, &out, "_buildSignRequestWithdraw", psd, incognitoAddress)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// BuildSignRequestWithdraw is a free data retrieval call binding the contract method 0x3cf6e1ca.
//
// Solidity: function _buildSignRequestWithdraw((uint8,address,bytes,uint256) psd, string incognitoAddress) pure returns(bytes)
func (_VaultHelper *VaultHelperSession) BuildSignRequestWithdraw(psd VaultHelperPreSignData, incognitoAddress string) ([]byte, error) {
	return _VaultHelper.Contract.BuildSignRequestWithdraw(&_VaultHelper.CallOpts, psd, incognitoAddress)
}

// BuildSignRequestWithdraw is a free data retrieval call binding the contract method 0x3cf6e1ca.
//
// Solidity: function _buildSignRequestWithdraw((uint8,address,bytes,uint256) psd, string incognitoAddress) pure returns(bytes)
func (_VaultHelper *VaultHelperCallerSession) BuildSignRequestWithdraw(psd VaultHelperPreSignData, incognitoAddress string) ([]byte, error) {
	return _VaultHelper.Contract.BuildSignRequestWithdraw(&_VaultHelper.CallOpts, psd, incognitoAddress)
}
