// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vault

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

// VaultHelperPreSignData is an auto generated low-level Go binding around an user-defined struct.
type VaultHelperPreSignData struct {
	Prefix    uint8
	Token     common.Address
	Timestamp []byte
	Amount    *big.Int
}

// VaultHelperMetaData contains all meta data concerning the VaultHelper contract.
var VaultHelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"enumVaultHelper.Prefix\",\"name\":\"prefix\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"timestamp\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"_buildPreSignData\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumVaultHelper.Prefix\",\"name\":\"prefix\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"timestamp\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structVaultHelper.PreSignData\",\"name\":\"psd\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"recipientToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"exchangeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"name\":\"_buildSignExecute\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumVaultHelper.Prefix\",\"name\":\"prefix\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"timestamp\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structVaultHelper.PreSignData\",\"name\":\"psd\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"}],\"name\":\"_buildSignRequestWithdraw\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumVaultHelper.Prefix\",\"name\":\"prefix\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"timestamp\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"newPreSignData\",\"outputs\":[{\"components\":[{\"internalType\":\"enumVaultHelper.Prefix\",\"name\":\"prefix\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"timestamp\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structVaultHelper.PreSignData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"741d7a50": "_buildPreSignData(uint8,address,bytes,uint256)",
		"41f59b6f": "_buildSignExecute((uint8,address,bytes,uint256),address,address,bytes)",
		"3cf6e1ca": "_buildSignRequestWithdraw((uint8,address,bytes,uint256),string)",
		"b95494e3": "newPreSignData(uint8,address,bytes,uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b5061068d806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80633cf6e1ca1461005157806341f59b6f1461007a578063741d7a501461008d578063b95494e3146100a2575b600080fd5b61006461005f366004610464565b6100c2565b604051610071919061057c565b60405180910390f35b6100646100883660046103d3565b610115565b6100a061009b36600461036a565b61016d565b005b6100b56100b036600461036a565b610174565b6040516100719190610585565b60006001845160018111156100d357fe5b146100dd57600080fd5b60008484846040516020016100f4939291906105e8565b60408051808303601f19018152919052805160209091012095945050505050565b6000808651600181111561012557fe5b1461012f57600080fd5b6000868686868660405160200161014a95949392919061059f565b60408051808303601f190181529190528051602090910120979650505050505050565b5050505050565b61017c6101fa565b6101846101fa565b604051806080016040528088600181111561019b57fe5b8152602001876001600160a01b0316815260200186868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060200184905291505095945050505050565b6040805160808101909152806000815260200160006001600160a01b0316815260200160608152602001600081525090565b80356001600160a01b038116811461024357600080fd5b92915050565b60008083601f84011261025a578182fd5b50813567ffffffffffffffff811115610271578182fd5b60208301915083602082850101111561028957600080fd5b9250929050565b80356002811061024357600080fd5b6000608082840312156102b0578081fd5b6102ba6080610618565b90506102c68383610290565b815260206102d68482850161022c565b81830152604083013567ffffffffffffffff808211156102f557600080fd5b818501915085601f83011261030957600080fd5b81358181111561031857600080fd5b61032a601f8201601f19168501610618565b9150808252868482850101111561034057600080fd5b80848401858401376000848284010152508060408501525050506060820135606082015292915050565b600080600080600060808688031215610381578081fd5b61038b8787610290565b945061039a876020880161022c565b9350604086013567ffffffffffffffff8111156103b5578182fd5b6103c188828901610249565b96999598509660600135949350505050565b6000806000806000608086880312156103ea578081fd5b853567ffffffffffffffff80821115610401578283fd5b61040d89838a0161029f565b96506020880135915061041f8261063f565b9094506040870135906104318261063f565b90935060608701359080821115610446578283fd5b5061045388828901610249565b969995985093965092949392505050565b600080600060408486031215610478578283fd5b833567ffffffffffffffff8082111561048f578485fd5b61049b8783880161029f565b945060208601359150808211156104b0578384fd5b506104bd86828701610249565b9497909650939450505050565b60008284528282602086013780602084860101526020601f19601f85011685010190509392505050565b600081516002811061050257fe5b83526020828101516001600160a01b0316818501526040808401516080918601829052805191860182905290835b8181101561054c5782810184015187820160a001528301610530565b8181111561055d578460a083890101525b5060609485015194860194909452505050601f01601f19160160a00190565b90815260200190565b60006020825261059860208301846104f4565b9392505050565b6000608082526105b260808301886104f4565b6001600160a01b0387811660208501528616604084015282810360608401526105dc8185876104ca565b98975050505050505050565b6000604082526105fb60408301866104f4565b828103602084015261060e8185876104ca565b9695505050505050565b60405181810167ffffffffffffffff8111828210171561063757600080fd5b604052919050565b6001600160a01b038116811461065457600080fd5b5056fea26469706673582212204984ec2f7cee86d1cffa0cd0d4ec4f015cb7f73ca00c5f036461c97e3d363d4064736f6c634300060c0033",
}

// VaultHelperABI is the input ABI used to generate the binding from.
// Deprecated: Use VaultHelperMetaData.ABI instead.
var VaultHelperABI = VaultHelperMetaData.ABI

// Deprecated: Use VaultHelperMetaData.Sigs instead.
// VaultHelperFuncSigs maps the 4-byte function signature to its string representation.
var VaultHelperFuncSigs = VaultHelperMetaData.Sigs

// VaultHelperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VaultHelperMetaData.Bin instead.
var VaultHelperBin = VaultHelperMetaData.Bin

// DeployVaultHelper deploys a new Ethereum contract, binding an instance of VaultHelper to it.
func DeployVaultHelper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VaultHelper, error) {
	parsed, err := VaultHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VaultHelperBin), backend)
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
// Solidity: function _buildSignExecute((uint8,address,bytes,uint256) psd, address recipientToken, address exchangeAddress, bytes callData) pure returns(bytes32)
func (_VaultHelper *VaultHelperCaller) BuildSignExecute(opts *bind.CallOpts, psd VaultHelperPreSignData, recipientToken common.Address, exchangeAddress common.Address, callData []byte) ([32]byte, error) {
	var out []interface{}
	err := _VaultHelper.contract.Call(opts, &out, "_buildSignExecute", psd, recipientToken, exchangeAddress, callData)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BuildSignExecute is a free data retrieval call binding the contract method 0x41f59b6f.
//
// Solidity: function _buildSignExecute((uint8,address,bytes,uint256) psd, address recipientToken, address exchangeAddress, bytes callData) pure returns(bytes32)
func (_VaultHelper *VaultHelperSession) BuildSignExecute(psd VaultHelperPreSignData, recipientToken common.Address, exchangeAddress common.Address, callData []byte) ([32]byte, error) {
	return _VaultHelper.Contract.BuildSignExecute(&_VaultHelper.CallOpts, psd, recipientToken, exchangeAddress, callData)
}

// BuildSignExecute is a free data retrieval call binding the contract method 0x41f59b6f.
//
// Solidity: function _buildSignExecute((uint8,address,bytes,uint256) psd, address recipientToken, address exchangeAddress, bytes callData) pure returns(bytes32)
func (_VaultHelper *VaultHelperCallerSession) BuildSignExecute(psd VaultHelperPreSignData, recipientToken common.Address, exchangeAddress common.Address, callData []byte) ([32]byte, error) {
	return _VaultHelper.Contract.BuildSignExecute(&_VaultHelper.CallOpts, psd, recipientToken, exchangeAddress, callData)
}

// BuildSignRequestWithdraw is a free data retrieval call binding the contract method 0x3cf6e1ca.
//
// Solidity: function _buildSignRequestWithdraw((uint8,address,bytes,uint256) psd, string incognitoAddress) pure returns(bytes32)
func (_VaultHelper *VaultHelperCaller) BuildSignRequestWithdraw(opts *bind.CallOpts, psd VaultHelperPreSignData, incognitoAddress string) ([32]byte, error) {
	var out []interface{}
	err := _VaultHelper.contract.Call(opts, &out, "_buildSignRequestWithdraw", psd, incognitoAddress)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BuildSignRequestWithdraw is a free data retrieval call binding the contract method 0x3cf6e1ca.
//
// Solidity: function _buildSignRequestWithdraw((uint8,address,bytes,uint256) psd, string incognitoAddress) pure returns(bytes32)
func (_VaultHelper *VaultHelperSession) BuildSignRequestWithdraw(psd VaultHelperPreSignData, incognitoAddress string) ([32]byte, error) {
	return _VaultHelper.Contract.BuildSignRequestWithdraw(&_VaultHelper.CallOpts, psd, incognitoAddress)
}

// BuildSignRequestWithdraw is a free data retrieval call binding the contract method 0x3cf6e1ca.
//
// Solidity: function _buildSignRequestWithdraw((uint8,address,bytes,uint256) psd, string incognitoAddress) pure returns(bytes32)
func (_VaultHelper *VaultHelperCallerSession) BuildSignRequestWithdraw(psd VaultHelperPreSignData, incognitoAddress string) ([32]byte, error) {
	return _VaultHelper.Contract.BuildSignRequestWithdraw(&_VaultHelper.CallOpts, psd, incognitoAddress)
}

// NewPreSignData is a free data retrieval call binding the contract method 0xb95494e3.
//
// Solidity: function newPreSignData(uint8 prefix, address token, bytes timestamp, uint256 amount) pure returns((uint8,address,bytes,uint256))
func (_VaultHelper *VaultHelperCaller) NewPreSignData(opts *bind.CallOpts, prefix uint8, token common.Address, timestamp []byte, amount *big.Int) (VaultHelperPreSignData, error) {
	var out []interface{}
	err := _VaultHelper.contract.Call(opts, &out, "newPreSignData", prefix, token, timestamp, amount)

	if err != nil {
		return *new(VaultHelperPreSignData), err
	}

	out0 := *abi.ConvertType(out[0], new(VaultHelperPreSignData)).(*VaultHelperPreSignData)

	return out0, err

}

// NewPreSignData is a free data retrieval call binding the contract method 0xb95494e3.
//
// Solidity: function newPreSignData(uint8 prefix, address token, bytes timestamp, uint256 amount) pure returns((uint8,address,bytes,uint256))
func (_VaultHelper *VaultHelperSession) NewPreSignData(prefix uint8, token common.Address, timestamp []byte, amount *big.Int) (VaultHelperPreSignData, error) {
	return _VaultHelper.Contract.NewPreSignData(&_VaultHelper.CallOpts, prefix, token, timestamp, amount)
}

// NewPreSignData is a free data retrieval call binding the contract method 0xb95494e3.
//
// Solidity: function newPreSignData(uint8 prefix, address token, bytes timestamp, uint256 amount) pure returns((uint8,address,bytes,uint256))
func (_VaultHelper *VaultHelperCallerSession) NewPreSignData(prefix uint8, token common.Address, timestamp []byte, amount *big.Int) (VaultHelperPreSignData, error) {
	return _VaultHelper.Contract.NewPreSignData(&_VaultHelper.CallOpts, prefix, token, timestamp, amount)
}
