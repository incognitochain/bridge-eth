// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pnft

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

// Execution is an auto generated low-level Go binding around an user-defined struct.
type Execution struct {
	Sell Input
	Buy  Input
}

// Fee is an auto generated low-level Go binding around an user-defined struct.
type Fee struct {
	Rate      uint16
	Recipient common.Address
}

// Input is an auto generated low-level Go binding around an user-defined struct.
type Input struct {
	Order            Order
	V                uint8
	R                [32]byte
	S                [32]byte
	ExtraSignature   []byte
	SignatureVersion uint8
	BlockNumber      *big.Int
}

// Order is an auto generated low-level Go binding around an user-defined struct.
type Order struct {
	Trader         common.Address
	Side           uint8
	MatchingPolicy common.Address
	Collection     common.Address
	TokenId        *big.Int
	Amount         *big.Int
	PaymentToken   common.Address
	Price          *big.Int
	ListingTime    *big.Int
	ExpirationTime *big.Int
	Fees           []Fee
	Salt           *big.Int
	ExtraParams    []byte
}

// AddressUpgradeableMetaData contains all meta data concerning the AddressUpgradeable contract.
var AddressUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220de8d927f760ccf5be4034d995054165792552cfb0d05de6428d1b32ab83cf22364736f6c63430008110033",
}

// AddressUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressUpgradeableMetaData.ABI instead.
var AddressUpgradeableABI = AddressUpgradeableMetaData.ABI

// AddressUpgradeableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressUpgradeableMetaData.Bin instead.
var AddressUpgradeableBin = AddressUpgradeableMetaData.Bin

// DeployAddressUpgradeable deploys a new Ethereum contract, binding an instance of AddressUpgradeable to it.
func DeployAddressUpgradeable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AddressUpgradeable, error) {
	parsed, err := AddressUpgradeableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressUpgradeableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AddressUpgradeable{AddressUpgradeableCaller: AddressUpgradeableCaller{contract: contract}, AddressUpgradeableTransactor: AddressUpgradeableTransactor{contract: contract}, AddressUpgradeableFilterer: AddressUpgradeableFilterer{contract: contract}}, nil
}

// AddressUpgradeable is an auto generated Go binding around an Ethereum contract.
type AddressUpgradeable struct {
	AddressUpgradeableCaller     // Read-only binding to the contract
	AddressUpgradeableTransactor // Write-only binding to the contract
	AddressUpgradeableFilterer   // Log filterer for contract events
}

// AddressUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressUpgradeableSession struct {
	Contract     *AddressUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AddressUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressUpgradeableCallerSession struct {
	Contract *AddressUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// AddressUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressUpgradeableTransactorSession struct {
	Contract     *AddressUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// AddressUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressUpgradeableRaw struct {
	Contract *AddressUpgradeable // Generic contract binding to access the raw methods on
}

// AddressUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressUpgradeableCallerRaw struct {
	Contract *AddressUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// AddressUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressUpgradeableTransactorRaw struct {
	Contract *AddressUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressUpgradeable creates a new instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeable(address common.Address, backend bind.ContractBackend) (*AddressUpgradeable, error) {
	contract, err := bindAddressUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeable{AddressUpgradeableCaller: AddressUpgradeableCaller{contract: contract}, AddressUpgradeableTransactor: AddressUpgradeableTransactor{contract: contract}, AddressUpgradeableFilterer: AddressUpgradeableFilterer{contract: contract}}, nil
}

// NewAddressUpgradeableCaller creates a new read-only instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*AddressUpgradeableCaller, error) {
	contract, err := bindAddressUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeableCaller{contract: contract}, nil
}

// NewAddressUpgradeableTransactor creates a new write-only instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressUpgradeableTransactor, error) {
	contract, err := bindAddressUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeableTransactor{contract: contract}, nil
}

// NewAddressUpgradeableFilterer creates a new log filterer instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressUpgradeableFilterer, error) {
	contract, err := bindAddressUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeableFilterer{contract: contract}, nil
}

// bindAddressUpgradeable binds a generic wrapper to an already deployed contract.
func bindAddressUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressUpgradeable *AddressUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressUpgradeable.Contract.AddressUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressUpgradeable *AddressUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.AddressUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressUpgradeable *AddressUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.AddressUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressUpgradeable *AddressUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressUpgradeable *AddressUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressUpgradeable *AddressUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// BlurExchangeMetaData contains all meta data concerning the BlurExchange contract.
var BlurExchangeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Closed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockRange\",\"type\":\"uint256\"}],\"name\":\"NewBlockRange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIExecutionDelegate\",\"name\":\"executionDelegate\",\"type\":\"address\"}],\"name\":\"NewExecutionDelegate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"NewOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIPolicyManager\",\"name\":\"policyManager\",\"type\":\"address\"}],\"name\":\"NewPolicyManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNonce\",\"type\":\"uint256\"}],\"name\":\"NonceIncremented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Opened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structOrder\",\"name\":\"sell\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sellHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structOrder\",\"name\":\"buy\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"buyHash\",\"type\":\"bytes32\"}],\"name\":\"OrdersMatched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEE_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INVERSE_BASIS_POINT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ORACLE_ORDER_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ORDER_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POOL\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROOT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"sell\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"buy\",\"type\":\"tuple\"}],\"name\":\"_execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockRange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"sell\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"buy\",\"type\":\"tuple\"}],\"internalType\":\"structExecution[]\",\"name\":\"executions\",\"type\":\"tuple[]\"}],\"name\":\"bulkExecute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"name\":\"cancelOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cancelledOrFilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"close\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"sell\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"buy\",\"type\":\"tuple\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"executionDelegate\",\"outputs\":[{\"internalType\":\"contractIExecutionDelegate\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"getOrderHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incrementNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIExecutionDelegate\",\"name\":\"_executionDelegate\",\"type\":\"address\"},{\"internalType\":\"contractIPolicyManager\",\"name\":\"_policyManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockRange\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInternal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOpen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"open\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"policyManager\",\"outputs\":[{\"internalType\":\"contractIPolicyManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remainingETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockRange\",\"type\":\"uint256\"}],\"name\":\"setBlockRange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIExecutionDelegate\",\"name\":\"_executionDelegate\",\"type\":\"address\"}],\"name\":\"setExecutionDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPolicyManager\",\"name\":\"_policyManager\",\"type\":\"address\"}],\"name\":\"setPolicyManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60a06040526000805460ff1990811682553060805261012f80549091169055610130553480156200002f57600080fd5b506200003a62000040565b62000102565b605f54610100900460ff1615620000ad5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b605f5460ff90811610156200010057605f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6080516139eb6200013a600039600081816107ef0152818161082f015281816109070152818161094701526109da01526139eb6000f3fe6080604052600436106102305760003560e01c80637ecebe001161012e578063b3be57f8116100ab578063f2fde38b1161006f578063f2fde38b146106ab578063f4acd740146106cb578063f973a209146106eb578063fcfff16f1461071f578063ffa1ad741461073457600080fd5b8063b3be57f81461062f578063cae6047f14610642578063cf756fdf14610658578063e04d94ae14610678578063e9f713ad1461068b57600080fd5b8063a4b2c674116100f2578063a4b2c6741461058f578063ab3dbf3b146105a6578063ab7e8cba146105c7578063ad5c4648146105e7578063adde41e11461060f57600080fd5b80637ecebe00146104c95780638da5cb5b146104f7578063986c9b20146105155780639a1fc3a714610536578063a3f4df7e1461054957600080fd5b80634832ede1116101bc5780636992aa36116101805780636992aa3614610418578063715018a6146104385780637535d2461461044d5780637adbf973146104885780637dc0d1d0146104a857600080fd5b80634832ede1146103765780634f1ef286146103aa57806352d1902d146103bd5780635511f319146103d2578063627cdcb91461040357600080fd5b806331e6d0fe1161020357806331e6d0fe146102e05780633644e515146103145780633659cfe61461032a57806343d726d61461034a57806347535d7b1461035f57600080fd5b8063037c9be21461023557806316e29d71146102575780631d97c9bb146102875780632c7acf8c146102c9575b600080fd5b34801561024157600080fd5b50610255610250366004612da0565b610763565b005b34801561026357600080fd5b5061012f546102729060ff1681565b60405190151581526020015b60405180910390f35b34801561029357600080fd5b506102bb7fd71080023d2f293ed0723dc287d6b2d4e7d27d0b6c12928e300721b7c78c748581565b60405190815260200161027e565b3480156102d557600080fd5b506102bb6101305481565b3480156102ec57600080fd5b506102bb7f5bcf4b2eaff7fcdeb49f0bda53026b9ebdd93db566fe4c447125cb899e598c9081565b34801561032057600080fd5b506102bb60325481565b34801561033657600080fd5b50610255610345366004612da0565b6107e5565b34801561035657600080fd5b506102556108c4565b34801561036b57600080fd5b506102bb6101285481565b34801561038257600080fd5b506102bb7f05b43f730f67de334a342883f867101fc7ef3361dfdff4a29a7aa97e0920ef7a81565b6102556103b8366004612e14565b6108fd565b3480156103c957600080fd5b506102bb6109cd565b3480156103de57600080fd5b506102726103ed366004612ebc565b61012d6020526000908152604090205460ff1681565b34801561040f57600080fd5b50610255610a80565b34801561042457600080fd5b50610255610433366004612ebc565b610aed565b34801561044457600080fd5b50610255610b31565b34801561045957600080fd5b506104706ea39bb272e79075ade125fd351887ac81565b6040516001600160a01b03909116815260200161027e565b34801561049457600080fd5b506102556104a3366004612da0565b610b45565b3480156104b457600080fd5b5061012b54610470906001600160a01b031681565b3480156104d557600080fd5b506102bb6104e4366004612da0565b61012e6020526000908152604090205481565b34801561050357600080fd5b506092546001600160a01b0316610470565b34801561052157600080fd5b5061012954610470906001600160a01b031681565b610255610544366004612eed565b610bbe565b34801561055557600080fd5b506105826040518060400160405280600d81526020016c426c75722045786368616e676560981b81525081565b60405161027e9190612f75565b34801561059b57600080fd5b506102bb61012c5481565b3480156105b257600080fd5b5061012a54610470906001600160a01b031681565b3480156105d357600080fd5b506102556105e2366004612ff4565b610c58565b3480156105f357600080fd5b5061047073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc281565b34801561061b57600080fd5b5061025561062a366004612da0565b610ca7565b61025561063d366004612ff4565b610d20565b34801561064e57600080fd5b506102bb61271081565b34801561066457600080fd5b50610255610673366004613036565b610e75565b610255610686366004612eed565b611039565b34801561069757600080fd5b506102bb6106a6366004613087565b6114fe565b3480156106b757600080fd5b506102556106c6366004612da0565b61151b565b3480156106d757600080fd5b506102556106e6366004613087565b611591565b3480156106f757600080fd5b506102bb7f376bfbc394a7ba7fdf10f224572cef371358e3053e362f4554fcd2ad56329b3f81565b34801561072b57600080fd5b506102556116c5565b34801561074057600080fd5b50610582604051806040016040528060038152602001620312e360ec1b81525081565b61076b6116fe565b6001600160a01b03811661079a5760405162461bcd60e51b8152600401610791906130c3565b60405180910390fd5b61012980546001600160a01b0319166001600160a01b0383169081179091556040517ff9a0f356a7ef079355de09d32ce45cc3cfabc8f118681c19a17501f005a376ac90600090a250565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016300361082d5760405162461bcd60e51b8152600401610791906130f3565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661087660008051602061396f833981519152546001600160a01b031690565b6001600160a01b03161461089c5760405162461bcd60e51b81526004016107919061313f565b6108a581611758565b604080516000808252602082019092526108c191839190611760565b50565b6108cc6116fe565b60006101288190556040517f1cdde67b72a90f19919ac732a437ac2f7a10fc128d28c2a6e525d89ce5cd9d3a9190a1565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036109455760405162461bcd60e51b8152600401610791906130f3565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661098e60008051602061396f833981519152546001600160a01b031690565b6001600160a01b0316146109b45760405162461bcd60e51b81526004016107919061313f565b6109bd82611758565b6109c982826001611760565b5050565b6000306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610a6d5760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c00000000000000006064820152608401610791565b5060008051602061396f83398151915290565b33600090815261012e60205260408120805460019290610aa19084906131a1565b909155505033600081815261012e60209081526040918290205491519182527fa82a649bbd060c9099cd7b7326e2b0dc9e9af0836480e0f849dc9eaa79710b3b910160405180910390a2565b610af56116fe565b61012c8190556040518181527f7706177c541ba1b858371bfc568aa77450b4713bbdbba62c730d4484ab6c12519060200160405180910390a150565b610b396116fe565b610b4360006118cb565b565b610b4d6116fe565b6001600160a01b038116610b735760405162461bcd60e51b8152600401610791906130c3565b61012b80546001600160a01b0319166001600160a01b0383169081179091556040517fb3eacd0e351fafdfefdec84e1cd19679b38dbcd63ea7c2c24da17fd2bc3b3c0e90600090a250565b61012854600114610bfa5760405162461bcd60e51b815260206004820152600660248201526510db1bdcd95960d21b6044820152606401610791565b61012f5460ff1615610c1e5760405162461bcd60e51b8152600401610791906131b4565b346101305561012f805460ff19166001179055610c3b8282611039565b610c4361191d565b505060006101305561012f805460ff19169055565b60005b60ff8116821115610ca257610c9083838360ff16818110610c7e57610c7e6131d9565b90506020028101906106e691906131ef565b80610c9a81613210565b915050610c5b565b505050565b610caf6116fe565b6001600160a01b038116610cd55760405162461bcd60e51b8152600401610791906130c3565b61012a80546001600160a01b0319166001600160a01b0383169081179091556040517fdbe18f3fd927cc2aefe380ffd2abfdb8e13f0239c0258ccfc84c3d8fdd8c041890600090a250565b61012854600114610d5c5760405162461bcd60e51b815260206004820152600660248201526510db1bdcd95960d21b6044820152606401610791565b61012f5460ff1615610d805760405162461bcd60e51b8152600401610791906131b4565b346101305561012f805460ff19166001179055806000819003610ddc5760405162461bcd60e51b81526020600482015260146024820152734e6f206f726465727320746f206578656375746560601b6044820152606401610791565b60005b818160ff161015610e5657604051602082028501358501600060018085018614908114610e1c576001850160200288013588018390039150610e22565b82360391505b50637026ca5760e11b8352808284600401376000806004830185305af4505050508080610e4e90613210565b915050610ddf565b50610e5f61191d565b50506000610130555061012f805460ff19169055565b605f54610100900460ff1615808015610e955750605f54600160ff909116105b80610eaf5750303b158015610eaf5750605f5460ff166001145b610f125760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610791565b605f805460ff191660011790558015610f3557605f805461ff0019166101001790555b610f3d61193b565b6001610128556040805160c081018252600d608082019081526c426c75722045786368616e676560981b60a083015281528151808301835260038152620312e360ec1b6020828101919091528201524691810191909152306060820152610fa39061196a565b60325561012980546001600160a01b038088166001600160a01b03199283161790925561012a805487841690831617905561012b80549286169290911691909117905561012c829055801561103257605f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b61012f5460ff1661105c5760405162461bcd60e51b8152600401610791906131b4565b60005460ff16156110a55760405162461bcd60e51b81526020600482015260136024820152721499595b9d1c985b98de4819195d1958dd1959606a1b6044820152606401610791565b6000805460ff191660019081179091556110bf83806131ef565b6110d090604081019060200161325d565b60018111156110e1576110e161322f565b146110eb57600080fd5b600061113d6110fa84806131ef565b61012e600061110987806131ef565b611117906020810190612da0565b6001600160a01b03166001600160a01b0316815260200190815260200160002054611a07565b9050600061114e6110fa84806131ef565b905061116361115d85806131ef565b83611b47565b6111af5760405162461bcd60e51b815260206004820152601b60248201527f53656c6c2068617320696e76616c696420706172616d657465727300000000006044820152606401610791565b6111c26111bc84806131ef565b82611b47565b61120e5760405162461bcd60e51b815260206004820152601a60248201527f4275792068617320696e76616c696420706172616d65746572730000000000006044820152606401610791565b6112188483611ba5565b6112645760405162461bcd60e51b815260206004820152601960248201527f53656c6c206661696c656420617574686f72697a6174696f6e000000000000006044820152606401610791565b61126e8382611ba5565b6112ba5760405162461bcd60e51b815260206004820152601860248201527f427579206661696c656420617574686f72697a6174696f6e00000000000000006044820152606401610791565b60008080806112db6112cc89806131ef565b6112d689806131ef565b611d47565b60008a815261012d60205260408082208054600160ff1991821681179092558c845291909220805490911690911790559296509094509250905061138261132289806131ef565b611330906020810190612da0565b61133a89806131ef565b611348906020810190612da0565b6113528b806131ef565b6113639060e081019060c001612da0565b61136d8c806131ef565b61137c9061014081019061327a565b89612075565b6113d861138f89806131ef565b6113a0906080810190606001612da0565b6113aa8a806131ef565b6113b8906020810190612da0565b6113c28a806131ef565b6113d0906020810190612da0565b868686612153565b6113e287806131ef565b61010001356113f189806131ef565b6101000135116114185761140587806131ef565b611413906020810190612da0565b611430565b61142288806131ef565b611430906020810190612da0565b6001600160a01b031661144388806131ef565b61010001356114528a806131ef565b6101000135111561147a5761146788806131ef565b611475906020810190612da0565b611492565b61148489806131ef565b611492906020810190612da0565b6001600160a01b03167f61cbb2a3dee0b6064c2e681aadd61677fb4ef319f0b547508d495626f5a62f646114c68b806131ef565b896114d18c806131ef565b8a6040516114e2949392919061352a565b60405180910390a350506000805460ff19169055505050505050565b60006115158261012e836111176020840184612da0565b92915050565b6115236116fe565b6001600160a01b0381166115885760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610791565b6108c1816118cb565b61159e6020820182612da0565b6001600160a01b0316336001600160a01b0316146115f35760405162461bcd60e51b81526020600482015260126024820152712737ba1039b2b73a10313c903a3930b232b960711b6044820152606401610791565b600061160a8261012e836111176020840184612da0565b600081815261012d602052604090205490915060ff161561166d5760405162461bcd60e51b815260206004820152601960248201527f4f726465722063616e63656c6c6564206f722066696c6c6564000000000000006044820152606401610791565b600081815261012d602052604090819020805460ff19166001179055517f5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d906116b99083815260200190565b60405180910390a15050565b6116cd6116fe565b6001610128556040517fd1dcd00534373f20882b79e6ab6875a5c358c5bd576448757ed50e63069ab51890600090a1565b6092546001600160a01b03163314610b435760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610791565b6108c16116fe565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff161561179357610ca283612288565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156117ed575060408051601f3d908101601f191682019092526117ea91810190613567565b60015b6118505760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b6064820152608401610791565b60008051602061396f83398151915281146118bf5760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b6064820152608401610791565b50610ca2838383612324565b609280546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6101305480156108c15760008060008084335af1806109c957600080fd5b605f54610100900460ff166119625760405162461bcd60e51b815260040161079190613580565b610b4361234f565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f826000015180519060200120836020015180519060200120846040015185606001516040516020016119ea9594939291909485526020850193909352604084019190915260608301526001600160a01b0316608082015260a00190565b604051602081830303815290604052805190602001209050919050565b60007f376bfbc394a7ba7fdf10f224572cef371358e3053e362f4554fcd2ad56329b3f611a376020850185612da0565b611a47604086016020870161325d565b611a576060870160408801612da0565b611a676080880160608901612da0565b608088013560a0890135611a8160e08b0160c08c01612da0565b8a60e001358b61010001358c6101200135611aab8e806101400190611aa6919061327a565b61237f565b8e61016001358f806101800190611ac291906135cb565b604051611ad0929190613612565b604051908190038120611af29e9d9c9b9a999897969594939291602001613622565b60408051601f1981840301815282825260208301859052910160408051601f1981840301815290829052611b2992916020016136ad565b60405160208183030381529060405280519060200120905092915050565b600080611b576020850185612da0565b6001600160a01b031614158015611b7e5750600082815261012d602052604090205460ff16155b8015611b8e575042836101000135105b8015611b9e575082610120013542105b9392505050565b600080611bb284806131ef565b611bc1906101808101906135cb565b9050118015611c105750611bd583806131ef565b611be4906101808101906135cb565b6000818110611bf557611bf56131d9565b9050013560f81c60f81b6001600160f81b031916600160f81b145b15611cab5761012c54611c2760c0850135436136dc565b10611c745760405162461bcd60e51b815260206004820181905260248201527f5369676e656420626c6f636b206e756d626572206f7574206f662072616e67656044820152606401610791565b611c9f82611c8860c0860160a0870161325d565b611c9560808701876135cb565b8760c00135612452565b611cab57506000611515565b33611cb684806131ef565b611cc4906020810190612da0565b6001600160a01b031603611cda57506001611515565b611d3282611ce885806131ef565b611cf6906020810190612da0565b611d0660408701602088016136ef565b60408701356060880135611d2060c08a0160a08b0161325d565b611d2d60808b018b6135cb565b6124e8565b611d3e57506000611515565b50600192915050565b600080600080600085610100013587610100013511611ec15761012a546001600160a01b031663874516cd611d8260608a0160408b01612da0565b6040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602401602060405180830381865afa158015611dc6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611dea9190613722565b611e325760405162461bcd60e51b8152602060048201526019602482015278141bdb1a58de481a5cc81b9bdd081dda1a5d195b1a5cdd1959603a1b6044820152606401610791565b611e426060880160408901612da0565b6001600160a01b031663d5ec8c7788886040518363ffffffff1660e01b8152600401611e6f92919061373d565b60a060405180830381865afa158015611e8c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611eb09190613762565b92985090965094509250905061201e565b61012a546001600160a01b031663874516cd611ee36060890160408a01612da0565b6040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602401602060405180830381865afa158015611f27573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f4b9190613722565b611f935760405162461bcd60e51b8152602060048201526019602482015278141bdb1a58de481a5cc81b9bdd081dda1a5d195b1a5cdd1959603a1b6044820152606401610791565b611fa36060870160408801612da0565b6001600160a01b0316630813a76687896040518363ffffffff1660e01b8152600401611fd092919061373d565b60a060405180830381865afa158015611fed573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120119190613762565b9298509096509450925090505b8061206b5760405162461bcd60e51b815260206004820152601860248201527f4f72646572732063616e6e6f74206265206d61746368656400000000000000006044820152606401610791565b5092959194509250565b6001600160a01b03841661212d57336001600160a01b038616146120cc5760405162461bcd60e51b815260206004820152600e60248201526d086c2dcdcdee840eae6ca408aa8960931b6044820152606401610791565b806101305410156121145760405162461bcd60e51b8152602060048201526012602482015271496e73756666696369656e742076616c756560701b6044820152606401610791565b80610130600082825461212791906136dc565b90915550505b600061213c84848789866125e3565b905061214a858789846126fe565b50505050505050565b60008160018111156121675761216761322f565b036121e85761012954604051633c4fc9fb60e11b81526001600160a01b03888116600483015287811660248301528681166044830152606482018690529091169063789f93f690608401600060405180830381600087803b1580156121cb57600080fd5b505af11580156121df573d6000803e3d6000fd5b50505050612280565b60018160018111156121fc576121fc61322f565b036122805761012954604051633a54a01760e11b81526001600160a01b038881166004830152878116602483015286811660448301526064820186905260848201859052909116906374a9402e9060a401600060405180830381600087803b15801561226757600080fd5b505af115801561227b573d6000803e3d6000fd5b505050505b505050505050565b6001600160a01b0381163b6122f55760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401610791565b60008051602061396f83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b61232d836129eb565b60008251118061233a5750805b15610ca2576123498383612a2b565b50505050565b605f54610100900460ff166123765760405162461bcd60e51b815260040161079190613580565b610b43336118cb565b6000808267ffffffffffffffff81111561239b5761239b612dcd565b6040519080825280602002602001820160405280156123c4578160200160208202803683370190505b50905060005b83811015612421576123f28585838181106123e7576123e76131d9565b905060400201612b16565b828281518110612404576124046131d9565b602090810291909101015280612419816137b6565b9150506123ca565b508060405160200161243391906137cf565b6040516020818303038152906040528051906020012091505092915050565b60008061245f8784612b86565b905060008080808960018111156124785761247861322f565b0361249257505085359050602086013560408701356124be565b60018960018111156124a6576124a661322f565b036124be575050506020850135604086013560608701355b61012b546124d8906001600160a01b031685858585612bee565b9450505050505b95945050505050565b600080808560018111156124fe576124fe61322f565b036125135761250c8a612cd5565b90506125c8565b60018560018111156125275761252761322f565b036125c857600061253a84860186613805565b9050600073__$ca4f43907f6e08ecee9109e3f06ef9ebdf$__639c7bf9388d846040518363ffffffff1660e01b815260040161257792919061389f565b602060405180830381865af4158015612594573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125b89190613567565b90506125c381612cec565b925050505b6125d589828a8a8a612bee565b9a9950505050505050505050565b600080805b60ff811687111561269557600061271089898460ff1681811061260d5761260d6131d9565b61262392602060409092020190810191506138ed565b6126319061ffff1687613908565b61263b919061391f565b905061267587878b8b8660ff16818110612657576126576131d9565b905060400201602001602081019061266f9190612da0565b846126fe565b61267f81846131a1565b925050808061268d90613210565b9150506125e8565b50828111156126e65760405162461bcd60e51b815260206004820152601c60248201527f4665657320617265206d6f7265207468616e20746865207072696365000000006044820152606401610791565b60006126f282856136dc565b98975050505050505050565b8015612349576001600160a01b038416612807576001600160a01b0382166127685760405162461bcd60e51b815260206004820152601860248201527f5472616e7366657220746f207a65726f206164647265737300000000000000006044820152606401610791565b6000826001600160a01b03168260405160006040518083038185875af1925050503d80600081146127b5576040519150601f19603f3d011682016040523d82523d6000602084013e6127ba565b606091505b50509050806128015760405162461bcd60e51b8152602060048201526013602482015272115512081d1c985b9cd9995c8819985a5b1959606a1b6044820152606401610791565b50612349565b6ea39bb272e79075ade125fd351887ab196001600160a01b038516016128f8576040516323b872dd60e01b81526001600160a01b03808516600483015283166024820152604481018290526000906ea39bb272e79075ade125fd351887ac906323b872dd906064016020604051808303816000875af115801561288e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128b29190613722565b9050806128015760405162461bcd60e51b8152602060048201526014602482015273141bdbdb081d1c985b9cd9995c8819985a5b195960621b6044820152606401610791565b73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc1196001600160a01b038516016129ab576101295460405163368fa33960e21b815273c02aaa39b223fe8d0a0e5c4f27ead9083c756cc260048201526001600160a01b0385811660248301528481166044830152606482018490529091169063da3e8ce490608401600060405180830381600087803b15801561298e57600080fd5b505af11580156129a2573d6000803e3d6000fd5b50505050612349565b60405162461bcd60e51b815260206004820152601560248201527424b73b30b634b2103830bcb6b2b73a103a37b5b2b760591b6044820152606401610791565b6129f481612288565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b60606001600160a01b0383163b612a935760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b6064820152608401610791565b600080846001600160a01b031684604051612aae9190613941565b600060405180830381855af49150503d8060008114612ae9576040519150601f19603f3d011682016040523d82523d6000602084013e612aee565b606091505b50915091506124df828260405180606001604052806027815260200161398f60279139612d4d565b60007f05b43f730f67de334a342883f867101fc7ef3361dfdff4a29a7aa97e0920ef7a612b4660208401846138ed565b612b566040850160208601612da0565b6040516020016119ea9392919092835261ffff9190911660208301526001600160a01b0316604082015260600190565b603254604080517fd71080023d2f293ed0723dc287d6b2d4e7d27d0b6c12928e300721b7c78c74856020820152908101849052606081018390526000919060800160405160208183030381529060405280519060200120604051602001611b29929190613953565b60008360ff16601b1480612c0557508360ff16601c145b612c475760405162461bcd60e51b815260206004820152601360248201527224b73b30b634b2103b103830b930b6b2ba32b960691b6044820152606401610791565b6040805160008082526020820180845288905260ff871692820192909252606081018590526080810184905260019060a0016020604051602081039080840390855afa158015612c9b573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116612cc05760009150506124df565b6001600160a01b0387811691161490506124df565b6000603254826040516020016119ea929190613953565b603254604080517f5bcf4b2eaff7fcdeb49f0bda53026b9ebdd93db566fe4c447125cb899e598c90602082015290810183905260009190606001604051602081830303815290604052805190602001206040516020016119ea929190613953565b60608315612d5c575081611b9e565b611b9e8383815115612d715781518083602001fd5b8060405162461bcd60e51b81526004016107919190612f75565b6001600160a01b03811681146108c157600080fd5b600060208284031215612db257600080fd5b8135611b9e81612d8b565b8035612dc881612d8b565b919050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715612e0c57612e0c612dcd565b604052919050565b60008060408385031215612e2757600080fd5b8235612e3281612d8b565b915060208381013567ffffffffffffffff80821115612e5057600080fd5b818601915086601f830112612e6457600080fd5b813581811115612e7657612e76612dcd565b612e88601f8201601f19168501612de3565b91508082528784828501011115612e9e57600080fd5b80848401858401376000848284010152508093505050509250929050565b600060208284031215612ece57600080fd5b5035919050565b600060e08284031215612ee757600080fd5b50919050565b60008060408385031215612f0057600080fd5b823567ffffffffffffffff80821115612f1857600080fd5b612f2486838701612ed5565b93506020850135915080821115612f3a57600080fd5b50612f4785828601612ed5565b9150509250929050565b60005b83811015612f6c578181015183820152602001612f54565b50506000910152565b6020815260008251806020840152612f94816040850160208701612f51565b601f01601f19169190910160400192915050565b60008083601f840112612fba57600080fd5b50813567ffffffffffffffff811115612fd257600080fd5b6020830191508360208260051b8501011115612fed57600080fd5b9250929050565b6000806020838503121561300757600080fd5b823567ffffffffffffffff81111561301e57600080fd5b61302a85828601612fa8565b90969095509350505050565b6000806000806080858703121561304c57600080fd5b843561305781612d8b565b9350602085013561306781612d8b565b9250604085013561307781612d8b565b9396929550929360600135925050565b60006020828403121561309957600080fd5b813567ffffffffffffffff8111156130b057600080fd5b82016101a08185031215611b9e57600080fd5b602080825260169082015275416464726573732063616e6e6f74206265207a65726f60501b604082015260600190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b808201808211156115155761151561318b565b6020808252600b908201526a155b9cd859994818d85b1b60aa1b604082015260600190565b634e487b7160e01b600052603260045260246000fd5b6000823561019e1983360301811261320657600080fd5b9190910192915050565b600060ff821660ff81036132265761322661318b565b60010192915050565b634e487b7160e01b600052602160045260246000fd5b600281106108c157600080fd5b8035612dc881613245565b60006020828403121561326f57600080fd5b8135611b9e81613245565b6000808335601e1984360301811261329157600080fd5b83018035915067ffffffffffffffff8211156132ac57600080fd5b6020019150600681901b3603821315612fed57600080fd5b600281106132e257634e487b7160e01b600052602160045260246000fd5b9052565b6000808335601e198436030181126132fd57600080fd5b830160208101925035905067ffffffffffffffff81111561331d57600080fd5b8060061b3603821315612fed57600080fd5b803561ffff81168114612dc857600080fd5b8183526000602080850194508260005b858110156133955761ffff6133658361332f565b1687528282013561337581612d8b565b6001600160a01b0316878401526040968701969190910190600101613351565b509495945050505050565b6000808335601e198436030181126133b757600080fd5b830160208101925035905067ffffffffffffffff8111156133d757600080fd5b803603821315612fed57600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60006101a061342e8461342185612dbd565b6001600160a01b03169052565b61343a60208401613252565b61344760208601826132c4565b5061345460408401612dbd565b6001600160a01b0316604085015261346e60608401612dbd565b6001600160a01b0381166060860152506080830135608085015260a083013560a085015261349e60c08401612dbd565b6001600160a01b031660c085015260e08381013590850152610100808401359085015261012080840135908501526101406134db818501856132e6565b83838801526134ed8488018284613341565b935050505061016080840135818601525061018061350d818501856133a0565b8684038388015261351f8482846133e6565b979650505050505050565b60808152600061353d608083018761340f565b8560208401528281036040840152613555818661340f565b91505082606083015295945050505050565b60006020828403121561357957600080fd5b5051919050565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b6000808335601e198436030181126135e257600080fd5b83018035915067ffffffffffffffff8211156135fd57600080fd5b602001915036819003821315612fed57600080fd5b8183823760009101908152919050565b8e81526001600160a01b038e1660208201526101c08101613646604083018f6132c4565b6001600160a01b039c8d1660608301529a8c16608082015260a081019990995260c08901979097529490981660e08701526101008601929092526101208501526101408401526101608301949094526101808201939093526101a001919091529392505050565b600083516136bf818460208801612f51565b8351908301906136d3818360208801612f51565b01949350505050565b818103818111156115155761151561318b565b60006020828403121561370157600080fd5b813560ff81168114611b9e57600080fd5b80518015158114612dc857600080fd5b60006020828403121561373457600080fd5b611b9e82613712565b604081526000613750604083018561340f565b82810360208401526124df818561340f565b600080600080600060a0868803121561377a57600080fd5b61378386613712565b945060208601519350604086015192506060860151915060808601516137a881613245565b809150509295509295909350565b6000600182016137c8576137c861318b565b5060010190565b815160009082906020808601845b838110156137f9578151855293820193908201906001016137dd565b50929695505050505050565b6000602080838503121561381857600080fd5b823567ffffffffffffffff8082111561383057600080fd5b818501915085601f83011261384457600080fd5b81358181111561385657613856612dcd565b8060051b9150613867848301612de3565b818152918301840191848101908884111561388157600080fd5b938501935b838510156126f257843582529385019390850190613886565b6000604082018483526020604081850152818551808452606086019150828701935060005b818110156138e0578451835293830193918301916001016138c4565b5090979650505050505050565b6000602082840312156138ff57600080fd5b611b9e8261332f565b80820281158282048414176115155761151561318b565b60008261393c57634e487b7160e01b600052601260045260246000fd5b500490565b60008251613206818460208701612f51565b61190160f01b8152600281019290925260228201526042019056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220b5fa38681d7c4e0454f857b9d297090856ddf139acf755d603909ef2def755bd64736f6c63430008110033",
}

// BlurExchangeABI is the input ABI used to generate the binding from.
// Deprecated: Use BlurExchangeMetaData.ABI instead.
var BlurExchangeABI = BlurExchangeMetaData.ABI

// BlurExchangeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BlurExchangeMetaData.Bin instead.
var BlurExchangeBin = BlurExchangeMetaData.Bin

// DeployBlurExchange deploys a new Ethereum contract, binding an instance of BlurExchange to it.
func DeployBlurExchange(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BlurExchange, error) {
	parsed, err := BlurExchangeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	merkleVerifierAddr, _, _, _ := DeployMerkleVerifier(auth, backend)
	BlurExchangeBin = strings.ReplaceAll(BlurExchangeBin, "__$ca4f43907f6e08ecee9109e3f06ef9ebdf$__", merkleVerifierAddr.String()[2:])

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BlurExchangeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BlurExchange{BlurExchangeCaller: BlurExchangeCaller{contract: contract}, BlurExchangeTransactor: BlurExchangeTransactor{contract: contract}, BlurExchangeFilterer: BlurExchangeFilterer{contract: contract}}, nil
}

// BlurExchange is an auto generated Go binding around an Ethereum contract.
type BlurExchange struct {
	BlurExchangeCaller     // Read-only binding to the contract
	BlurExchangeTransactor // Write-only binding to the contract
	BlurExchangeFilterer   // Log filterer for contract events
}

// BlurExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlurExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlurExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlurExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlurExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlurExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlurExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlurExchangeSession struct {
	Contract     *BlurExchange     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlurExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlurExchangeCallerSession struct {
	Contract *BlurExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BlurExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlurExchangeTransactorSession struct {
	Contract     *BlurExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BlurExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlurExchangeRaw struct {
	Contract *BlurExchange // Generic contract binding to access the raw methods on
}

// BlurExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlurExchangeCallerRaw struct {
	Contract *BlurExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// BlurExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlurExchangeTransactorRaw struct {
	Contract *BlurExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlurExchange creates a new instance of BlurExchange, bound to a specific deployed contract.
func NewBlurExchange(address common.Address, backend bind.ContractBackend) (*BlurExchange, error) {
	contract, err := bindBlurExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlurExchange{BlurExchangeCaller: BlurExchangeCaller{contract: contract}, BlurExchangeTransactor: BlurExchangeTransactor{contract: contract}, BlurExchangeFilterer: BlurExchangeFilterer{contract: contract}}, nil
}

// NewBlurExchangeCaller creates a new read-only instance of BlurExchange, bound to a specific deployed contract.
func NewBlurExchangeCaller(address common.Address, caller bind.ContractCaller) (*BlurExchangeCaller, error) {
	contract, err := bindBlurExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlurExchangeCaller{contract: contract}, nil
}

// NewBlurExchangeTransactor creates a new write-only instance of BlurExchange, bound to a specific deployed contract.
func NewBlurExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*BlurExchangeTransactor, error) {
	contract, err := bindBlurExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlurExchangeTransactor{contract: contract}, nil
}

// NewBlurExchangeFilterer creates a new log filterer instance of BlurExchange, bound to a specific deployed contract.
func NewBlurExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*BlurExchangeFilterer, error) {
	contract, err := bindBlurExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlurExchangeFilterer{contract: contract}, nil
}

// bindBlurExchange binds a generic wrapper to an already deployed contract.
func bindBlurExchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BlurExchangeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlurExchange *BlurExchangeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlurExchange.Contract.BlurExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlurExchange *BlurExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlurExchange.Contract.BlurExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlurExchange *BlurExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlurExchange.Contract.BlurExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlurExchange *BlurExchangeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlurExchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlurExchange *BlurExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlurExchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlurExchange *BlurExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlurExchange.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BlurExchange *BlurExchangeCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BlurExchange.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BlurExchange *BlurExchangeSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _BlurExchange.Contract.DOMAINSEPARATOR(&_BlurExchange.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BlurExchange *BlurExchangeCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _BlurExchange.Contract.DOMAINSEPARATOR(&_BlurExchange.CallOpts)
}

// FEETYPEHASH is a free data retrieval call binding the contract method 0x4832ede1.
//
// Solidity: function FEE_TYPEHASH() view returns(bytes32)
func (_BlurExchange *BlurExchangeCaller) FEETYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BlurExchange.contract.Call(opts, &out, "FEE_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FEETYPEHASH is a free data retrieval call binding the contract method 0x4832ede1.
//
// Solidity: function FEE_TYPEHASH() view returns(bytes32)
func (_BlurExchange *BlurExchangeSession) FEETYPEHASH() ([32]byte, error) {
	return _BlurExchange.Contract.FEETYPEHASH(&_BlurExchange.CallOpts)
}

// FEETYPEHASH is a free data retrieval call binding the contract method 0x4832ede1.
//
// Solidity: function FEE_TYPEHASH() view returns(bytes32)
func (_BlurExchange *BlurExchangeCallerSession) FEETYPEHASH() ([32]byte, error) {
	return _BlurExchange.Contract.FEETYPEHASH(&_BlurExchange.CallOpts)
}

// INVERSEBASISPOINT is a free data retrieval call binding the contract method 0xcae6047f.
//
// Solidity: function INVERSE_BASIS_POINT() view returns(uint256)
func (_BlurExchange *BlurExchangeCaller) INVERSEBASISPOINT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlurExchange.contract.Call(opts, &out, "INVERSE_BASIS_POINT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INVERSEBASISPOINT is a free data retrieval call binding the contract method 0xcae6047f.
//
// Solidity: function INVERSE_BASIS_POINT() view returns(uint256)
func (_BlurExchange *BlurExchangeSession) INVERSEBASISPOINT() (*big.Int, error) {
	return _BlurExchange.Contract.INVERSEBASISPOINT(&_BlurExchange.CallOpts)
}

// INVERSEBASISPOINT is a free data retrieval call binding the contract method 0xcae6047f.
//
// Solidity: function INVERSE_BASIS_POINT() view returns(uint256)
func (_BlurExchange *BlurExchangeCallerSession) INVERSEBASISPOINT() (*big.Int, error) {
	return _BlurExchange.Contract.INVERSEBASISPOINT(&_BlurExchange.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_BlurExchange *BlurExchangeCaller) NAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BlurExchange.contract.Call(opts, &out, "NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_BlurExchange *BlurExchangeSession) NAME() (string, error) {
	return _BlurExchange.Contract.NAME(&_BlurExchange.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_BlurExchange *BlurExchangeCallerSession) NAME() (string, error) {
	return _BlurExchange.Contract.NAME(&_BlurExchange.CallOpts)
}

// ORACLEORDERTYPEHASH is a free data retrieval call binding the contract method 0x1d97c9bb.
//
// Solidity: function ORACLE_ORDER_TYPEHASH() view returns(bytes32)
func (_BlurExchange *BlurExchangeCaller) ORACLEORDERTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BlurExchange.contract.Call(opts, &out, "ORACLE_ORDER_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ORACLEORDERTYPEHASH is a free data retrieval call binding the contract method 0x1d97c9bb.
//
// Solidity: function ORACLE_ORDER_TYPEHASH() view returns(bytes32)
func (_BlurExchange *BlurExchangeSession) ORACLEORDERTYPEHASH() ([32]byte, error) {
	return _BlurExchange.Contract.ORACLEORDERTYPEHASH(&_BlurExchange.CallOpts)
}

// ORACLEORDERTYPEHASH is a free data retrieval call binding the contract method 0x1d97c9bb.
//
// Solidity: function ORACLE_ORDER_TYPEHASH() view returns(bytes32)
func (_BlurExchange *BlurExchangeCallerSession) ORACLEORDERTYPEHASH() ([32]byte, error) {
	return _BlurExchange.Contract.ORACLEORDERTYPEHASH(&_BlurExchange.CallOpts)
}

// ORDERTYPEHASH is a free data retrieval call binding the contract method 0xf973a209.
//
// Solidity: function ORDER_TYPEHASH() view returns(bytes32)
func (_BlurExchange *BlurExchangeCaller) ORDERTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BlurExchange.contract.Call(opts, &out, "ORDER_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ORDERTYPEHASH is a free data retrieval call binding the contract method 0xf973a209.
//
// Solidity: function ORDER_TYPEHASH() view returns(bytes32)
func (_BlurExchange *BlurExchangeSession) ORDERTYPEHASH() ([32]byte, error) {
	return _BlurExchange.Contract.ORDERTYPEHASH(&_BlurExchange.CallOpts)
}

// ORDERTYPEHASH is a free data retrieval call binding the contract method 0xf973a209.
//
// Solidity: function ORDER_TYPEHASH() view returns(bytes32)
func (_BlurExchange *BlurExchangeCallerSession) ORDERTYPEHASH() ([32]byte, error) {
	return _BlurExchange.Contract.ORDERTYPEHASH(&_BlurExchange.CallOpts)
}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_BlurExchange *BlurExchangeCaller) POOL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlurExchange.contract.Call(opts, &out, "POOL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_BlurExchange *BlurExchangeSession) POOL() (common.Address, error) {
	return _BlurExchange.Contract.POOL(&_BlurExchange.CallOpts)
}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_BlurExchange *BlurExchangeCallerSession) POOL() (common.Address, error) {
	return _BlurExchange.Contract.POOL(&_BlurExchange.CallOpts)
}

// ROOTTYPEHASH is a free data retrieval call binding the contract method 0x31e6d0fe.
//
// Solidity: function ROOT_TYPEHASH() view returns(bytes32)
func (_BlurExchange *BlurExchangeCaller) ROOTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BlurExchange.contract.Call(opts, &out, "ROOT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ROOTTYPEHASH is a free data retrieval call binding the contract method 0x31e6d0fe.
//
// Solidity: function ROOT_TYPEHASH() view returns(bytes32)
func (_BlurExchange *BlurExchangeSession) ROOTTYPEHASH() ([32]byte, error) {
	return _BlurExchange.Contract.ROOTTYPEHASH(&_BlurExchange.CallOpts)
}

// ROOTTYPEHASH is a free data retrieval call binding the contract method 0x31e6d0fe.
//
// Solidity: function ROOT_TYPEHASH() view returns(bytes32)
func (_BlurExchange *BlurExchangeCallerSession) ROOTTYPEHASH() ([32]byte, error) {
	return _BlurExchange.Contract.ROOTTYPEHASH(&_BlurExchange.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_BlurExchange *BlurExchangeCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BlurExchange.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_BlurExchange *BlurExchangeSession) VERSION() (string, error) {
	return _BlurExchange.Contract.VERSION(&_BlurExchange.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_BlurExchange *BlurExchangeCallerSession) VERSION() (string, error) {
	return _BlurExchange.Contract.VERSION(&_BlurExchange.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_BlurExchange *BlurExchangeCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlurExchange.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_BlurExchange *BlurExchangeSession) WETH() (common.Address, error) {
	return _BlurExchange.Contract.WETH(&_BlurExchange.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_BlurExchange *BlurExchangeCallerSession) WETH() (common.Address, error) {
	return _BlurExchange.Contract.WETH(&_BlurExchange.CallOpts)
}

// BlockRange is a free data retrieval call binding the contract method 0xa4b2c674.
//
// Solidity: function blockRange() view returns(uint256)
func (_BlurExchange *BlurExchangeCaller) BlockRange(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlurExchange.contract.Call(opts, &out, "blockRange")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockRange is a free data retrieval call binding the contract method 0xa4b2c674.
//
// Solidity: function blockRange() view returns(uint256)
func (_BlurExchange *BlurExchangeSession) BlockRange() (*big.Int, error) {
	return _BlurExchange.Contract.BlockRange(&_BlurExchange.CallOpts)
}

// BlockRange is a free data retrieval call binding the contract method 0xa4b2c674.
//
// Solidity: function blockRange() view returns(uint256)
func (_BlurExchange *BlurExchangeCallerSession) BlockRange() (*big.Int, error) {
	return _BlurExchange.Contract.BlockRange(&_BlurExchange.CallOpts)
}

// CancelledOrFilled is a free data retrieval call binding the contract method 0x5511f319.
//
// Solidity: function cancelledOrFilled(bytes32 ) view returns(bool)
func (_BlurExchange *BlurExchangeCaller) CancelledOrFilled(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _BlurExchange.contract.Call(opts, &out, "cancelledOrFilled", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CancelledOrFilled is a free data retrieval call binding the contract method 0x5511f319.
//
// Solidity: function cancelledOrFilled(bytes32 ) view returns(bool)
func (_BlurExchange *BlurExchangeSession) CancelledOrFilled(arg0 [32]byte) (bool, error) {
	return _BlurExchange.Contract.CancelledOrFilled(&_BlurExchange.CallOpts, arg0)
}

// CancelledOrFilled is a free data retrieval call binding the contract method 0x5511f319.
//
// Solidity: function cancelledOrFilled(bytes32 ) view returns(bool)
func (_BlurExchange *BlurExchangeCallerSession) CancelledOrFilled(arg0 [32]byte) (bool, error) {
	return _BlurExchange.Contract.CancelledOrFilled(&_BlurExchange.CallOpts, arg0)
}

// ExecutionDelegate is a free data retrieval call binding the contract method 0x986c9b20.
//
// Solidity: function executionDelegate() view returns(address)
func (_BlurExchange *BlurExchangeCaller) ExecutionDelegate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlurExchange.contract.Call(opts, &out, "executionDelegate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExecutionDelegate is a free data retrieval call binding the contract method 0x986c9b20.
//
// Solidity: function executionDelegate() view returns(address)
func (_BlurExchange *BlurExchangeSession) ExecutionDelegate() (common.Address, error) {
	return _BlurExchange.Contract.ExecutionDelegate(&_BlurExchange.CallOpts)
}

// ExecutionDelegate is a free data retrieval call binding the contract method 0x986c9b20.
//
// Solidity: function executionDelegate() view returns(address)
func (_BlurExchange *BlurExchangeCallerSession) ExecutionDelegate() (common.Address, error) {
	return _BlurExchange.Contract.ExecutionDelegate(&_BlurExchange.CallOpts)
}

// GetOrderHash is a free data retrieval call binding the contract method 0xe9f713ad.
//
// Solidity: function getOrderHash((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) order) view returns(bytes32)
func (_BlurExchange *BlurExchangeCaller) GetOrderHash(opts *bind.CallOpts, order Order) ([32]byte, error) {
	var out []interface{}
	err := _BlurExchange.contract.Call(opts, &out, "getOrderHash", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOrderHash is a free data retrieval call binding the contract method 0xe9f713ad.
//
// Solidity: function getOrderHash((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) order) view returns(bytes32)
func (_BlurExchange *BlurExchangeSession) GetOrderHash(order Order) ([32]byte, error) {
	return _BlurExchange.Contract.GetOrderHash(&_BlurExchange.CallOpts, order)
}

// GetOrderHash is a free data retrieval call binding the contract method 0xe9f713ad.
//
// Solidity: function getOrderHash((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) order) view returns(bytes32)
func (_BlurExchange *BlurExchangeCallerSession) GetOrderHash(order Order) ([32]byte, error) {
	return _BlurExchange.Contract.GetOrderHash(&_BlurExchange.CallOpts, order)
}

// IsInternal is a free data retrieval call binding the contract method 0x16e29d71.
//
// Solidity: function isInternal() view returns(bool)
func (_BlurExchange *BlurExchangeCaller) IsInternal(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BlurExchange.contract.Call(opts, &out, "isInternal")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInternal is a free data retrieval call binding the contract method 0x16e29d71.
//
// Solidity: function isInternal() view returns(bool)
func (_BlurExchange *BlurExchangeSession) IsInternal() (bool, error) {
	return _BlurExchange.Contract.IsInternal(&_BlurExchange.CallOpts)
}

// IsInternal is a free data retrieval call binding the contract method 0x16e29d71.
//
// Solidity: function isInternal() view returns(bool)
func (_BlurExchange *BlurExchangeCallerSession) IsInternal() (bool, error) {
	return _BlurExchange.Contract.IsInternal(&_BlurExchange.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(uint256)
func (_BlurExchange *BlurExchangeCaller) IsOpen(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlurExchange.contract.Call(opts, &out, "isOpen")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(uint256)
func (_BlurExchange *BlurExchangeSession) IsOpen() (*big.Int, error) {
	return _BlurExchange.Contract.IsOpen(&_BlurExchange.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(uint256)
func (_BlurExchange *BlurExchangeCallerSession) IsOpen() (*big.Int, error) {
	return _BlurExchange.Contract.IsOpen(&_BlurExchange.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_BlurExchange *BlurExchangeCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BlurExchange.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_BlurExchange *BlurExchangeSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _BlurExchange.Contract.Nonces(&_BlurExchange.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_BlurExchange *BlurExchangeCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _BlurExchange.Contract.Nonces(&_BlurExchange.CallOpts, arg0)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_BlurExchange *BlurExchangeCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlurExchange.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_BlurExchange *BlurExchangeSession) Oracle() (common.Address, error) {
	return _BlurExchange.Contract.Oracle(&_BlurExchange.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_BlurExchange *BlurExchangeCallerSession) Oracle() (common.Address, error) {
	return _BlurExchange.Contract.Oracle(&_BlurExchange.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BlurExchange *BlurExchangeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlurExchange.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BlurExchange *BlurExchangeSession) Owner() (common.Address, error) {
	return _BlurExchange.Contract.Owner(&_BlurExchange.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BlurExchange *BlurExchangeCallerSession) Owner() (common.Address, error) {
	return _BlurExchange.Contract.Owner(&_BlurExchange.CallOpts)
}

// PolicyManager is a free data retrieval call binding the contract method 0xab3dbf3b.
//
// Solidity: function policyManager() view returns(address)
func (_BlurExchange *BlurExchangeCaller) PolicyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlurExchange.contract.Call(opts, &out, "policyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PolicyManager is a free data retrieval call binding the contract method 0xab3dbf3b.
//
// Solidity: function policyManager() view returns(address)
func (_BlurExchange *BlurExchangeSession) PolicyManager() (common.Address, error) {
	return _BlurExchange.Contract.PolicyManager(&_BlurExchange.CallOpts)
}

// PolicyManager is a free data retrieval call binding the contract method 0xab3dbf3b.
//
// Solidity: function policyManager() view returns(address)
func (_BlurExchange *BlurExchangeCallerSession) PolicyManager() (common.Address, error) {
	return _BlurExchange.Contract.PolicyManager(&_BlurExchange.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BlurExchange *BlurExchangeCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BlurExchange.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BlurExchange *BlurExchangeSession) ProxiableUUID() ([32]byte, error) {
	return _BlurExchange.Contract.ProxiableUUID(&_BlurExchange.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BlurExchange *BlurExchangeCallerSession) ProxiableUUID() ([32]byte, error) {
	return _BlurExchange.Contract.ProxiableUUID(&_BlurExchange.CallOpts)
}

// RemainingETH is a free data retrieval call binding the contract method 0x2c7acf8c.
//
// Solidity: function remainingETH() view returns(uint256)
func (_BlurExchange *BlurExchangeCaller) RemainingETH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlurExchange.contract.Call(opts, &out, "remainingETH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemainingETH is a free data retrieval call binding the contract method 0x2c7acf8c.
//
// Solidity: function remainingETH() view returns(uint256)
func (_BlurExchange *BlurExchangeSession) RemainingETH() (*big.Int, error) {
	return _BlurExchange.Contract.RemainingETH(&_BlurExchange.CallOpts)
}

// RemainingETH is a free data retrieval call binding the contract method 0x2c7acf8c.
//
// Solidity: function remainingETH() view returns(uint256)
func (_BlurExchange *BlurExchangeCallerSession) RemainingETH() (*big.Int, error) {
	return _BlurExchange.Contract.RemainingETH(&_BlurExchange.CallOpts)
}

// Execute2 is a paid mutator transaction binding the contract method 0xe04d94ae.
//
// Solidity: function _execute(((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) sell, ((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) buy) payable returns()
func (_BlurExchange *BlurExchangeTransactor) Execute2(opts *bind.TransactOpts, sell Input, buy Input) (*types.Transaction, error) {
	return _BlurExchange.contract.Transact(opts, "_execute", sell, buy)
}

// Execute2 is a paid mutator transaction binding the contract method 0xe04d94ae.
//
// Solidity: function _execute(((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) sell, ((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) buy) payable returns()
func (_BlurExchange *BlurExchangeSession) Execute2(sell Input, buy Input) (*types.Transaction, error) {
	return _BlurExchange.Contract.Execute2(&_BlurExchange.TransactOpts, sell, buy)
}

// Execute2 is a paid mutator transaction binding the contract method 0xe04d94ae.
//
// Solidity: function _execute(((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) sell, ((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) buy) payable returns()
func (_BlurExchange *BlurExchangeTransactorSession) Execute2(sell Input, buy Input) (*types.Transaction, error) {
	return _BlurExchange.Contract.Execute2(&_BlurExchange.TransactOpts, sell, buy)
}

// BulkExecute is a paid mutator transaction binding the contract method 0xb3be57f8.
//
// Solidity: function bulkExecute((((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256),((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256))[] executions) payable returns()
func (_BlurExchange *BlurExchangeTransactor) BulkExecute(opts *bind.TransactOpts, executions []Execution) (*types.Transaction, error) {
	return _BlurExchange.contract.Transact(opts, "bulkExecute", executions)
}

// BulkExecute is a paid mutator transaction binding the contract method 0xb3be57f8.
//
// Solidity: function bulkExecute((((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256),((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256))[] executions) payable returns()
func (_BlurExchange *BlurExchangeSession) BulkExecute(executions []Execution) (*types.Transaction, error) {
	return _BlurExchange.Contract.BulkExecute(&_BlurExchange.TransactOpts, executions)
}

// BulkExecute is a paid mutator transaction binding the contract method 0xb3be57f8.
//
// Solidity: function bulkExecute((((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256),((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256))[] executions) payable returns()
func (_BlurExchange *BlurExchangeTransactorSession) BulkExecute(executions []Execution) (*types.Transaction, error) {
	return _BlurExchange.Contract.BulkExecute(&_BlurExchange.TransactOpts, executions)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xf4acd740.
//
// Solidity: function cancelOrder((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) order) returns()
func (_BlurExchange *BlurExchangeTransactor) CancelOrder(opts *bind.TransactOpts, order Order) (*types.Transaction, error) {
	return _BlurExchange.contract.Transact(opts, "cancelOrder", order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xf4acd740.
//
// Solidity: function cancelOrder((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) order) returns()
func (_BlurExchange *BlurExchangeSession) CancelOrder(order Order) (*types.Transaction, error) {
	return _BlurExchange.Contract.CancelOrder(&_BlurExchange.TransactOpts, order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xf4acd740.
//
// Solidity: function cancelOrder((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) order) returns()
func (_BlurExchange *BlurExchangeTransactorSession) CancelOrder(order Order) (*types.Transaction, error) {
	return _BlurExchange.Contract.CancelOrder(&_BlurExchange.TransactOpts, order)
}

// CancelOrders is a paid mutator transaction binding the contract method 0xab7e8cba.
//
// Solidity: function cancelOrders((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes)[] orders) returns()
func (_BlurExchange *BlurExchangeTransactor) CancelOrders(opts *bind.TransactOpts, orders []Order) (*types.Transaction, error) {
	return _BlurExchange.contract.Transact(opts, "cancelOrders", orders)
}

// CancelOrders is a paid mutator transaction binding the contract method 0xab7e8cba.
//
// Solidity: function cancelOrders((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes)[] orders) returns()
func (_BlurExchange *BlurExchangeSession) CancelOrders(orders []Order) (*types.Transaction, error) {
	return _BlurExchange.Contract.CancelOrders(&_BlurExchange.TransactOpts, orders)
}

// CancelOrders is a paid mutator transaction binding the contract method 0xab7e8cba.
//
// Solidity: function cancelOrders((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes)[] orders) returns()
func (_BlurExchange *BlurExchangeTransactorSession) CancelOrders(orders []Order) (*types.Transaction, error) {
	return _BlurExchange.Contract.CancelOrders(&_BlurExchange.TransactOpts, orders)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_BlurExchange *BlurExchangeTransactor) Close(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlurExchange.contract.Transact(opts, "close")
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_BlurExchange *BlurExchangeSession) Close() (*types.Transaction, error) {
	return _BlurExchange.Contract.Close(&_BlurExchange.TransactOpts)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_BlurExchange *BlurExchangeTransactorSession) Close() (*types.Transaction, error) {
	return _BlurExchange.Contract.Close(&_BlurExchange.TransactOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x9a1fc3a7.
//
// Solidity: function execute(((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) sell, ((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) buy) payable returns()
func (_BlurExchange *BlurExchangeTransactor) Execute(opts *bind.TransactOpts, sell Input, buy Input) (*types.Transaction, error) {
	return _BlurExchange.contract.Transact(opts, "execute", sell, buy)
}

// Execute is a paid mutator transaction binding the contract method 0x9a1fc3a7.
//
// Solidity: function execute(((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) sell, ((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) buy) payable returns()
func (_BlurExchange *BlurExchangeSession) Execute(sell Input, buy Input) (*types.Transaction, error) {
	return _BlurExchange.Contract.Execute(&_BlurExchange.TransactOpts, sell, buy)
}

// Execute is a paid mutator transaction binding the contract method 0x9a1fc3a7.
//
// Solidity: function execute(((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) sell, ((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) buy) payable returns()
func (_BlurExchange *BlurExchangeTransactorSession) Execute(sell Input, buy Input) (*types.Transaction, error) {
	return _BlurExchange.Contract.Execute(&_BlurExchange.TransactOpts, sell, buy)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns()
func (_BlurExchange *BlurExchangeTransactor) IncrementNonce(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlurExchange.contract.Transact(opts, "incrementNonce")
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns()
func (_BlurExchange *BlurExchangeSession) IncrementNonce() (*types.Transaction, error) {
	return _BlurExchange.Contract.IncrementNonce(&_BlurExchange.TransactOpts)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns()
func (_BlurExchange *BlurExchangeTransactorSession) IncrementNonce() (*types.Transaction, error) {
	return _BlurExchange.Contract.IncrementNonce(&_BlurExchange.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address _executionDelegate, address _policyManager, address _oracle, uint256 _blockRange) returns()
func (_BlurExchange *BlurExchangeTransactor) Initialize(opts *bind.TransactOpts, _executionDelegate common.Address, _policyManager common.Address, _oracle common.Address, _blockRange *big.Int) (*types.Transaction, error) {
	return _BlurExchange.contract.Transact(opts, "initialize", _executionDelegate, _policyManager, _oracle, _blockRange)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address _executionDelegate, address _policyManager, address _oracle, uint256 _blockRange) returns()
func (_BlurExchange *BlurExchangeSession) Initialize(_executionDelegate common.Address, _policyManager common.Address, _oracle common.Address, _blockRange *big.Int) (*types.Transaction, error) {
	return _BlurExchange.Contract.Initialize(&_BlurExchange.TransactOpts, _executionDelegate, _policyManager, _oracle, _blockRange)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address _executionDelegate, address _policyManager, address _oracle, uint256 _blockRange) returns()
func (_BlurExchange *BlurExchangeTransactorSession) Initialize(_executionDelegate common.Address, _policyManager common.Address, _oracle common.Address, _blockRange *big.Int) (*types.Transaction, error) {
	return _BlurExchange.Contract.Initialize(&_BlurExchange.TransactOpts, _executionDelegate, _policyManager, _oracle, _blockRange)
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() returns()
func (_BlurExchange *BlurExchangeTransactor) Open(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlurExchange.contract.Transact(opts, "open")
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() returns()
func (_BlurExchange *BlurExchangeSession) Open() (*types.Transaction, error) {
	return _BlurExchange.Contract.Open(&_BlurExchange.TransactOpts)
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() returns()
func (_BlurExchange *BlurExchangeTransactorSession) Open() (*types.Transaction, error) {
	return _BlurExchange.Contract.Open(&_BlurExchange.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BlurExchange *BlurExchangeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlurExchange.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BlurExchange *BlurExchangeSession) RenounceOwnership() (*types.Transaction, error) {
	return _BlurExchange.Contract.RenounceOwnership(&_BlurExchange.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BlurExchange *BlurExchangeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BlurExchange.Contract.RenounceOwnership(&_BlurExchange.TransactOpts)
}

// SetBlockRange is a paid mutator transaction binding the contract method 0x6992aa36.
//
// Solidity: function setBlockRange(uint256 _blockRange) returns()
func (_BlurExchange *BlurExchangeTransactor) SetBlockRange(opts *bind.TransactOpts, _blockRange *big.Int) (*types.Transaction, error) {
	return _BlurExchange.contract.Transact(opts, "setBlockRange", _blockRange)
}

// SetBlockRange is a paid mutator transaction binding the contract method 0x6992aa36.
//
// Solidity: function setBlockRange(uint256 _blockRange) returns()
func (_BlurExchange *BlurExchangeSession) SetBlockRange(_blockRange *big.Int) (*types.Transaction, error) {
	return _BlurExchange.Contract.SetBlockRange(&_BlurExchange.TransactOpts, _blockRange)
}

// SetBlockRange is a paid mutator transaction binding the contract method 0x6992aa36.
//
// Solidity: function setBlockRange(uint256 _blockRange) returns()
func (_BlurExchange *BlurExchangeTransactorSession) SetBlockRange(_blockRange *big.Int) (*types.Transaction, error) {
	return _BlurExchange.Contract.SetBlockRange(&_BlurExchange.TransactOpts, _blockRange)
}

// SetExecutionDelegate is a paid mutator transaction binding the contract method 0x037c9be2.
//
// Solidity: function setExecutionDelegate(address _executionDelegate) returns()
func (_BlurExchange *BlurExchangeTransactor) SetExecutionDelegate(opts *bind.TransactOpts, _executionDelegate common.Address) (*types.Transaction, error) {
	return _BlurExchange.contract.Transact(opts, "setExecutionDelegate", _executionDelegate)
}

// SetExecutionDelegate is a paid mutator transaction binding the contract method 0x037c9be2.
//
// Solidity: function setExecutionDelegate(address _executionDelegate) returns()
func (_BlurExchange *BlurExchangeSession) SetExecutionDelegate(_executionDelegate common.Address) (*types.Transaction, error) {
	return _BlurExchange.Contract.SetExecutionDelegate(&_BlurExchange.TransactOpts, _executionDelegate)
}

// SetExecutionDelegate is a paid mutator transaction binding the contract method 0x037c9be2.
//
// Solidity: function setExecutionDelegate(address _executionDelegate) returns()
func (_BlurExchange *BlurExchangeTransactorSession) SetExecutionDelegate(_executionDelegate common.Address) (*types.Transaction, error) {
	return _BlurExchange.Contract.SetExecutionDelegate(&_BlurExchange.TransactOpts, _executionDelegate)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_BlurExchange *BlurExchangeTransactor) SetOracle(opts *bind.TransactOpts, _oracle common.Address) (*types.Transaction, error) {
	return _BlurExchange.contract.Transact(opts, "setOracle", _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_BlurExchange *BlurExchangeSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _BlurExchange.Contract.SetOracle(&_BlurExchange.TransactOpts, _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_BlurExchange *BlurExchangeTransactorSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _BlurExchange.Contract.SetOracle(&_BlurExchange.TransactOpts, _oracle)
}

// SetPolicyManager is a paid mutator transaction binding the contract method 0xadde41e1.
//
// Solidity: function setPolicyManager(address _policyManager) returns()
func (_BlurExchange *BlurExchangeTransactor) SetPolicyManager(opts *bind.TransactOpts, _policyManager common.Address) (*types.Transaction, error) {
	return _BlurExchange.contract.Transact(opts, "setPolicyManager", _policyManager)
}

// SetPolicyManager is a paid mutator transaction binding the contract method 0xadde41e1.
//
// Solidity: function setPolicyManager(address _policyManager) returns()
func (_BlurExchange *BlurExchangeSession) SetPolicyManager(_policyManager common.Address) (*types.Transaction, error) {
	return _BlurExchange.Contract.SetPolicyManager(&_BlurExchange.TransactOpts, _policyManager)
}

// SetPolicyManager is a paid mutator transaction binding the contract method 0xadde41e1.
//
// Solidity: function setPolicyManager(address _policyManager) returns()
func (_BlurExchange *BlurExchangeTransactorSession) SetPolicyManager(_policyManager common.Address) (*types.Transaction, error) {
	return _BlurExchange.Contract.SetPolicyManager(&_BlurExchange.TransactOpts, _policyManager)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BlurExchange *BlurExchangeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BlurExchange.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BlurExchange *BlurExchangeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BlurExchange.Contract.TransferOwnership(&_BlurExchange.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BlurExchange *BlurExchangeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BlurExchange.Contract.TransferOwnership(&_BlurExchange.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_BlurExchange *BlurExchangeTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _BlurExchange.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_BlurExchange *BlurExchangeSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _BlurExchange.Contract.UpgradeTo(&_BlurExchange.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_BlurExchange *BlurExchangeTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _BlurExchange.Contract.UpgradeTo(&_BlurExchange.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BlurExchange *BlurExchangeTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BlurExchange.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BlurExchange *BlurExchangeSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BlurExchange.Contract.UpgradeToAndCall(&_BlurExchange.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BlurExchange *BlurExchangeTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BlurExchange.Contract.UpgradeToAndCall(&_BlurExchange.TransactOpts, newImplementation, data)
}

// BlurExchangeAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the BlurExchange contract.
type BlurExchangeAdminChangedIterator struct {
	Event *BlurExchangeAdminChanged // Event containing the contract specifics and raw log

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
func (it *BlurExchangeAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurExchangeAdminChanged)
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
		it.Event = new(BlurExchangeAdminChanged)
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
func (it *BlurExchangeAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurExchangeAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurExchangeAdminChanged represents a AdminChanged event raised by the BlurExchange contract.
type BlurExchangeAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_BlurExchange *BlurExchangeFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*BlurExchangeAdminChangedIterator, error) {

	logs, sub, err := _BlurExchange.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &BlurExchangeAdminChangedIterator{contract: _BlurExchange.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_BlurExchange *BlurExchangeFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *BlurExchangeAdminChanged) (event.Subscription, error) {

	logs, sub, err := _BlurExchange.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurExchangeAdminChanged)
				if err := _BlurExchange.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_BlurExchange *BlurExchangeFilterer) ParseAdminChanged(log types.Log) (*BlurExchangeAdminChanged, error) {
	event := new(BlurExchangeAdminChanged)
	if err := _BlurExchange.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurExchangeBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the BlurExchange contract.
type BlurExchangeBeaconUpgradedIterator struct {
	Event *BlurExchangeBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *BlurExchangeBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurExchangeBeaconUpgraded)
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
		it.Event = new(BlurExchangeBeaconUpgraded)
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
func (it *BlurExchangeBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurExchangeBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurExchangeBeaconUpgraded represents a BeaconUpgraded event raised by the BlurExchange contract.
type BlurExchangeBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_BlurExchange *BlurExchangeFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*BlurExchangeBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _BlurExchange.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &BlurExchangeBeaconUpgradedIterator{contract: _BlurExchange.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_BlurExchange *BlurExchangeFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *BlurExchangeBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _BlurExchange.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurExchangeBeaconUpgraded)
				if err := _BlurExchange.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_BlurExchange *BlurExchangeFilterer) ParseBeaconUpgraded(log types.Log) (*BlurExchangeBeaconUpgraded, error) {
	event := new(BlurExchangeBeaconUpgraded)
	if err := _BlurExchange.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurExchangeClosedIterator is returned from FilterClosed and is used to iterate over the raw logs and unpacked data for Closed events raised by the BlurExchange contract.
type BlurExchangeClosedIterator struct {
	Event *BlurExchangeClosed // Event containing the contract specifics and raw log

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
func (it *BlurExchangeClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurExchangeClosed)
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
		it.Event = new(BlurExchangeClosed)
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
func (it *BlurExchangeClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurExchangeClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurExchangeClosed represents a Closed event raised by the BlurExchange contract.
type BlurExchangeClosed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterClosed is a free log retrieval operation binding the contract event 0x1cdde67b72a90f19919ac732a437ac2f7a10fc128d28c2a6e525d89ce5cd9d3a.
//
// Solidity: event Closed()
func (_BlurExchange *BlurExchangeFilterer) FilterClosed(opts *bind.FilterOpts) (*BlurExchangeClosedIterator, error) {

	logs, sub, err := _BlurExchange.contract.FilterLogs(opts, "Closed")
	if err != nil {
		return nil, err
	}
	return &BlurExchangeClosedIterator{contract: _BlurExchange.contract, event: "Closed", logs: logs, sub: sub}, nil
}

// WatchClosed is a free log subscription operation binding the contract event 0x1cdde67b72a90f19919ac732a437ac2f7a10fc128d28c2a6e525d89ce5cd9d3a.
//
// Solidity: event Closed()
func (_BlurExchange *BlurExchangeFilterer) WatchClosed(opts *bind.WatchOpts, sink chan<- *BlurExchangeClosed) (event.Subscription, error) {

	logs, sub, err := _BlurExchange.contract.WatchLogs(opts, "Closed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurExchangeClosed)
				if err := _BlurExchange.contract.UnpackLog(event, "Closed", log); err != nil {
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

// ParseClosed is a log parse operation binding the contract event 0x1cdde67b72a90f19919ac732a437ac2f7a10fc128d28c2a6e525d89ce5cd9d3a.
//
// Solidity: event Closed()
func (_BlurExchange *BlurExchangeFilterer) ParseClosed(log types.Log) (*BlurExchangeClosed, error) {
	event := new(BlurExchangeClosed)
	if err := _BlurExchange.contract.UnpackLog(event, "Closed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurExchangeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BlurExchange contract.
type BlurExchangeInitializedIterator struct {
	Event *BlurExchangeInitialized // Event containing the contract specifics and raw log

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
func (it *BlurExchangeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurExchangeInitialized)
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
		it.Event = new(BlurExchangeInitialized)
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
func (it *BlurExchangeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurExchangeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurExchangeInitialized represents a Initialized event raised by the BlurExchange contract.
type BlurExchangeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BlurExchange *BlurExchangeFilterer) FilterInitialized(opts *bind.FilterOpts) (*BlurExchangeInitializedIterator, error) {

	logs, sub, err := _BlurExchange.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BlurExchangeInitializedIterator{contract: _BlurExchange.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BlurExchange *BlurExchangeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BlurExchangeInitialized) (event.Subscription, error) {

	logs, sub, err := _BlurExchange.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurExchangeInitialized)
				if err := _BlurExchange.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BlurExchange *BlurExchangeFilterer) ParseInitialized(log types.Log) (*BlurExchangeInitialized, error) {
	event := new(BlurExchangeInitialized)
	if err := _BlurExchange.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurExchangeNewBlockRangeIterator is returned from FilterNewBlockRange and is used to iterate over the raw logs and unpacked data for NewBlockRange events raised by the BlurExchange contract.
type BlurExchangeNewBlockRangeIterator struct {
	Event *BlurExchangeNewBlockRange // Event containing the contract specifics and raw log

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
func (it *BlurExchangeNewBlockRangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurExchangeNewBlockRange)
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
		it.Event = new(BlurExchangeNewBlockRange)
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
func (it *BlurExchangeNewBlockRangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurExchangeNewBlockRangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurExchangeNewBlockRange represents a NewBlockRange event raised by the BlurExchange contract.
type BlurExchangeNewBlockRange struct {
	BlockRange *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewBlockRange is a free log retrieval operation binding the contract event 0x7706177c541ba1b858371bfc568aa77450b4713bbdbba62c730d4484ab6c1251.
//
// Solidity: event NewBlockRange(uint256 blockRange)
func (_BlurExchange *BlurExchangeFilterer) FilterNewBlockRange(opts *bind.FilterOpts) (*BlurExchangeNewBlockRangeIterator, error) {

	logs, sub, err := _BlurExchange.contract.FilterLogs(opts, "NewBlockRange")
	if err != nil {
		return nil, err
	}
	return &BlurExchangeNewBlockRangeIterator{contract: _BlurExchange.contract, event: "NewBlockRange", logs: logs, sub: sub}, nil
}

// WatchNewBlockRange is a free log subscription operation binding the contract event 0x7706177c541ba1b858371bfc568aa77450b4713bbdbba62c730d4484ab6c1251.
//
// Solidity: event NewBlockRange(uint256 blockRange)
func (_BlurExchange *BlurExchangeFilterer) WatchNewBlockRange(opts *bind.WatchOpts, sink chan<- *BlurExchangeNewBlockRange) (event.Subscription, error) {

	logs, sub, err := _BlurExchange.contract.WatchLogs(opts, "NewBlockRange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurExchangeNewBlockRange)
				if err := _BlurExchange.contract.UnpackLog(event, "NewBlockRange", log); err != nil {
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

// ParseNewBlockRange is a log parse operation binding the contract event 0x7706177c541ba1b858371bfc568aa77450b4713bbdbba62c730d4484ab6c1251.
//
// Solidity: event NewBlockRange(uint256 blockRange)
func (_BlurExchange *BlurExchangeFilterer) ParseNewBlockRange(log types.Log) (*BlurExchangeNewBlockRange, error) {
	event := new(BlurExchangeNewBlockRange)
	if err := _BlurExchange.contract.UnpackLog(event, "NewBlockRange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurExchangeNewExecutionDelegateIterator is returned from FilterNewExecutionDelegate and is used to iterate over the raw logs and unpacked data for NewExecutionDelegate events raised by the BlurExchange contract.
type BlurExchangeNewExecutionDelegateIterator struct {
	Event *BlurExchangeNewExecutionDelegate // Event containing the contract specifics and raw log

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
func (it *BlurExchangeNewExecutionDelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurExchangeNewExecutionDelegate)
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
		it.Event = new(BlurExchangeNewExecutionDelegate)
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
func (it *BlurExchangeNewExecutionDelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurExchangeNewExecutionDelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurExchangeNewExecutionDelegate represents a NewExecutionDelegate event raised by the BlurExchange contract.
type BlurExchangeNewExecutionDelegate struct {
	ExecutionDelegate common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewExecutionDelegate is a free log retrieval operation binding the contract event 0xf9a0f356a7ef079355de09d32ce45cc3cfabc8f118681c19a17501f005a376ac.
//
// Solidity: event NewExecutionDelegate(address indexed executionDelegate)
func (_BlurExchange *BlurExchangeFilterer) FilterNewExecutionDelegate(opts *bind.FilterOpts, executionDelegate []common.Address) (*BlurExchangeNewExecutionDelegateIterator, error) {

	var executionDelegateRule []interface{}
	for _, executionDelegateItem := range executionDelegate {
		executionDelegateRule = append(executionDelegateRule, executionDelegateItem)
	}

	logs, sub, err := _BlurExchange.contract.FilterLogs(opts, "NewExecutionDelegate", executionDelegateRule)
	if err != nil {
		return nil, err
	}
	return &BlurExchangeNewExecutionDelegateIterator{contract: _BlurExchange.contract, event: "NewExecutionDelegate", logs: logs, sub: sub}, nil
}

// WatchNewExecutionDelegate is a free log subscription operation binding the contract event 0xf9a0f356a7ef079355de09d32ce45cc3cfabc8f118681c19a17501f005a376ac.
//
// Solidity: event NewExecutionDelegate(address indexed executionDelegate)
func (_BlurExchange *BlurExchangeFilterer) WatchNewExecutionDelegate(opts *bind.WatchOpts, sink chan<- *BlurExchangeNewExecutionDelegate, executionDelegate []common.Address) (event.Subscription, error) {

	var executionDelegateRule []interface{}
	for _, executionDelegateItem := range executionDelegate {
		executionDelegateRule = append(executionDelegateRule, executionDelegateItem)
	}

	logs, sub, err := _BlurExchange.contract.WatchLogs(opts, "NewExecutionDelegate", executionDelegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurExchangeNewExecutionDelegate)
				if err := _BlurExchange.contract.UnpackLog(event, "NewExecutionDelegate", log); err != nil {
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

// ParseNewExecutionDelegate is a log parse operation binding the contract event 0xf9a0f356a7ef079355de09d32ce45cc3cfabc8f118681c19a17501f005a376ac.
//
// Solidity: event NewExecutionDelegate(address indexed executionDelegate)
func (_BlurExchange *BlurExchangeFilterer) ParseNewExecutionDelegate(log types.Log) (*BlurExchangeNewExecutionDelegate, error) {
	event := new(BlurExchangeNewExecutionDelegate)
	if err := _BlurExchange.contract.UnpackLog(event, "NewExecutionDelegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurExchangeNewOracleIterator is returned from FilterNewOracle and is used to iterate over the raw logs and unpacked data for NewOracle events raised by the BlurExchange contract.
type BlurExchangeNewOracleIterator struct {
	Event *BlurExchangeNewOracle // Event containing the contract specifics and raw log

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
func (it *BlurExchangeNewOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurExchangeNewOracle)
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
		it.Event = new(BlurExchangeNewOracle)
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
func (it *BlurExchangeNewOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurExchangeNewOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurExchangeNewOracle represents a NewOracle event raised by the BlurExchange contract.
type BlurExchangeNewOracle struct {
	Oracle common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewOracle is a free log retrieval operation binding the contract event 0xb3eacd0e351fafdfefdec84e1cd19679b38dbcd63ea7c2c24da17fd2bc3b3c0e.
//
// Solidity: event NewOracle(address indexed oracle)
func (_BlurExchange *BlurExchangeFilterer) FilterNewOracle(opts *bind.FilterOpts, oracle []common.Address) (*BlurExchangeNewOracleIterator, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _BlurExchange.contract.FilterLogs(opts, "NewOracle", oracleRule)
	if err != nil {
		return nil, err
	}
	return &BlurExchangeNewOracleIterator{contract: _BlurExchange.contract, event: "NewOracle", logs: logs, sub: sub}, nil
}

// WatchNewOracle is a free log subscription operation binding the contract event 0xb3eacd0e351fafdfefdec84e1cd19679b38dbcd63ea7c2c24da17fd2bc3b3c0e.
//
// Solidity: event NewOracle(address indexed oracle)
func (_BlurExchange *BlurExchangeFilterer) WatchNewOracle(opts *bind.WatchOpts, sink chan<- *BlurExchangeNewOracle, oracle []common.Address) (event.Subscription, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _BlurExchange.contract.WatchLogs(opts, "NewOracle", oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurExchangeNewOracle)
				if err := _BlurExchange.contract.UnpackLog(event, "NewOracle", log); err != nil {
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

// ParseNewOracle is a log parse operation binding the contract event 0xb3eacd0e351fafdfefdec84e1cd19679b38dbcd63ea7c2c24da17fd2bc3b3c0e.
//
// Solidity: event NewOracle(address indexed oracle)
func (_BlurExchange *BlurExchangeFilterer) ParseNewOracle(log types.Log) (*BlurExchangeNewOracle, error) {
	event := new(BlurExchangeNewOracle)
	if err := _BlurExchange.contract.UnpackLog(event, "NewOracle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurExchangeNewPolicyManagerIterator is returned from FilterNewPolicyManager and is used to iterate over the raw logs and unpacked data for NewPolicyManager events raised by the BlurExchange contract.
type BlurExchangeNewPolicyManagerIterator struct {
	Event *BlurExchangeNewPolicyManager // Event containing the contract specifics and raw log

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
func (it *BlurExchangeNewPolicyManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurExchangeNewPolicyManager)
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
		it.Event = new(BlurExchangeNewPolicyManager)
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
func (it *BlurExchangeNewPolicyManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurExchangeNewPolicyManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurExchangeNewPolicyManager represents a NewPolicyManager event raised by the BlurExchange contract.
type BlurExchangeNewPolicyManager struct {
	PolicyManager common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewPolicyManager is a free log retrieval operation binding the contract event 0xdbe18f3fd927cc2aefe380ffd2abfdb8e13f0239c0258ccfc84c3d8fdd8c0418.
//
// Solidity: event NewPolicyManager(address indexed policyManager)
func (_BlurExchange *BlurExchangeFilterer) FilterNewPolicyManager(opts *bind.FilterOpts, policyManager []common.Address) (*BlurExchangeNewPolicyManagerIterator, error) {

	var policyManagerRule []interface{}
	for _, policyManagerItem := range policyManager {
		policyManagerRule = append(policyManagerRule, policyManagerItem)
	}

	logs, sub, err := _BlurExchange.contract.FilterLogs(opts, "NewPolicyManager", policyManagerRule)
	if err != nil {
		return nil, err
	}
	return &BlurExchangeNewPolicyManagerIterator{contract: _BlurExchange.contract, event: "NewPolicyManager", logs: logs, sub: sub}, nil
}

// WatchNewPolicyManager is a free log subscription operation binding the contract event 0xdbe18f3fd927cc2aefe380ffd2abfdb8e13f0239c0258ccfc84c3d8fdd8c0418.
//
// Solidity: event NewPolicyManager(address indexed policyManager)
func (_BlurExchange *BlurExchangeFilterer) WatchNewPolicyManager(opts *bind.WatchOpts, sink chan<- *BlurExchangeNewPolicyManager, policyManager []common.Address) (event.Subscription, error) {

	var policyManagerRule []interface{}
	for _, policyManagerItem := range policyManager {
		policyManagerRule = append(policyManagerRule, policyManagerItem)
	}

	logs, sub, err := _BlurExchange.contract.WatchLogs(opts, "NewPolicyManager", policyManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurExchangeNewPolicyManager)
				if err := _BlurExchange.contract.UnpackLog(event, "NewPolicyManager", log); err != nil {
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

// ParseNewPolicyManager is a log parse operation binding the contract event 0xdbe18f3fd927cc2aefe380ffd2abfdb8e13f0239c0258ccfc84c3d8fdd8c0418.
//
// Solidity: event NewPolicyManager(address indexed policyManager)
func (_BlurExchange *BlurExchangeFilterer) ParseNewPolicyManager(log types.Log) (*BlurExchangeNewPolicyManager, error) {
	event := new(BlurExchangeNewPolicyManager)
	if err := _BlurExchange.contract.UnpackLog(event, "NewPolicyManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurExchangeNonceIncrementedIterator is returned from FilterNonceIncremented and is used to iterate over the raw logs and unpacked data for NonceIncremented events raised by the BlurExchange contract.
type BlurExchangeNonceIncrementedIterator struct {
	Event *BlurExchangeNonceIncremented // Event containing the contract specifics and raw log

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
func (it *BlurExchangeNonceIncrementedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurExchangeNonceIncremented)
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
		it.Event = new(BlurExchangeNonceIncremented)
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
func (it *BlurExchangeNonceIncrementedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurExchangeNonceIncrementedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurExchangeNonceIncremented represents a NonceIncremented event raised by the BlurExchange contract.
type BlurExchangeNonceIncremented struct {
	Trader   common.Address
	NewNonce *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNonceIncremented is a free log retrieval operation binding the contract event 0xa82a649bbd060c9099cd7b7326e2b0dc9e9af0836480e0f849dc9eaa79710b3b.
//
// Solidity: event NonceIncremented(address indexed trader, uint256 newNonce)
func (_BlurExchange *BlurExchangeFilterer) FilterNonceIncremented(opts *bind.FilterOpts, trader []common.Address) (*BlurExchangeNonceIncrementedIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _BlurExchange.contract.FilterLogs(opts, "NonceIncremented", traderRule)
	if err != nil {
		return nil, err
	}
	return &BlurExchangeNonceIncrementedIterator{contract: _BlurExchange.contract, event: "NonceIncremented", logs: logs, sub: sub}, nil
}

// WatchNonceIncremented is a free log subscription operation binding the contract event 0xa82a649bbd060c9099cd7b7326e2b0dc9e9af0836480e0f849dc9eaa79710b3b.
//
// Solidity: event NonceIncremented(address indexed trader, uint256 newNonce)
func (_BlurExchange *BlurExchangeFilterer) WatchNonceIncremented(opts *bind.WatchOpts, sink chan<- *BlurExchangeNonceIncremented, trader []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _BlurExchange.contract.WatchLogs(opts, "NonceIncremented", traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurExchangeNonceIncremented)
				if err := _BlurExchange.contract.UnpackLog(event, "NonceIncremented", log); err != nil {
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

// ParseNonceIncremented is a log parse operation binding the contract event 0xa82a649bbd060c9099cd7b7326e2b0dc9e9af0836480e0f849dc9eaa79710b3b.
//
// Solidity: event NonceIncremented(address indexed trader, uint256 newNonce)
func (_BlurExchange *BlurExchangeFilterer) ParseNonceIncremented(log types.Log) (*BlurExchangeNonceIncremented, error) {
	event := new(BlurExchangeNonceIncremented)
	if err := _BlurExchange.contract.UnpackLog(event, "NonceIncremented", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurExchangeOpenedIterator is returned from FilterOpened and is used to iterate over the raw logs and unpacked data for Opened events raised by the BlurExchange contract.
type BlurExchangeOpenedIterator struct {
	Event *BlurExchangeOpened // Event containing the contract specifics and raw log

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
func (it *BlurExchangeOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurExchangeOpened)
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
		it.Event = new(BlurExchangeOpened)
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
func (it *BlurExchangeOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurExchangeOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurExchangeOpened represents a Opened event raised by the BlurExchange contract.
type BlurExchangeOpened struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOpened is a free log retrieval operation binding the contract event 0xd1dcd00534373f20882b79e6ab6875a5c358c5bd576448757ed50e63069ab518.
//
// Solidity: event Opened()
func (_BlurExchange *BlurExchangeFilterer) FilterOpened(opts *bind.FilterOpts) (*BlurExchangeOpenedIterator, error) {

	logs, sub, err := _BlurExchange.contract.FilterLogs(opts, "Opened")
	if err != nil {
		return nil, err
	}
	return &BlurExchangeOpenedIterator{contract: _BlurExchange.contract, event: "Opened", logs: logs, sub: sub}, nil
}

// WatchOpened is a free log subscription operation binding the contract event 0xd1dcd00534373f20882b79e6ab6875a5c358c5bd576448757ed50e63069ab518.
//
// Solidity: event Opened()
func (_BlurExchange *BlurExchangeFilterer) WatchOpened(opts *bind.WatchOpts, sink chan<- *BlurExchangeOpened) (event.Subscription, error) {

	logs, sub, err := _BlurExchange.contract.WatchLogs(opts, "Opened")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurExchangeOpened)
				if err := _BlurExchange.contract.UnpackLog(event, "Opened", log); err != nil {
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

// ParseOpened is a log parse operation binding the contract event 0xd1dcd00534373f20882b79e6ab6875a5c358c5bd576448757ed50e63069ab518.
//
// Solidity: event Opened()
func (_BlurExchange *BlurExchangeFilterer) ParseOpened(log types.Log) (*BlurExchangeOpened, error) {
	event := new(BlurExchangeOpened)
	if err := _BlurExchange.contract.UnpackLog(event, "Opened", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurExchangeOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the BlurExchange contract.
type BlurExchangeOrderCancelledIterator struct {
	Event *BlurExchangeOrderCancelled // Event containing the contract specifics and raw log

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
func (it *BlurExchangeOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurExchangeOrderCancelled)
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
		it.Event = new(BlurExchangeOrderCancelled)
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
func (it *BlurExchangeOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurExchangeOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurExchangeOrderCancelled represents a OrderCancelled event raised by the BlurExchange contract.
type BlurExchangeOrderCancelled struct {
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 hash)
func (_BlurExchange *BlurExchangeFilterer) FilterOrderCancelled(opts *bind.FilterOpts) (*BlurExchangeOrderCancelledIterator, error) {

	logs, sub, err := _BlurExchange.contract.FilterLogs(opts, "OrderCancelled")
	if err != nil {
		return nil, err
	}
	return &BlurExchangeOrderCancelledIterator{contract: _BlurExchange.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 hash)
func (_BlurExchange *BlurExchangeFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *BlurExchangeOrderCancelled) (event.Subscription, error) {

	logs, sub, err := _BlurExchange.contract.WatchLogs(opts, "OrderCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurExchangeOrderCancelled)
				if err := _BlurExchange.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
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

// ParseOrderCancelled is a log parse operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 hash)
func (_BlurExchange *BlurExchangeFilterer) ParseOrderCancelled(log types.Log) (*BlurExchangeOrderCancelled, error) {
	event := new(BlurExchangeOrderCancelled)
	if err := _BlurExchange.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurExchangeOrdersMatchedIterator is returned from FilterOrdersMatched and is used to iterate over the raw logs and unpacked data for OrdersMatched events raised by the BlurExchange contract.
type BlurExchangeOrdersMatchedIterator struct {
	Event *BlurExchangeOrdersMatched // Event containing the contract specifics and raw log

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
func (it *BlurExchangeOrdersMatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurExchangeOrdersMatched)
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
		it.Event = new(BlurExchangeOrdersMatched)
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
func (it *BlurExchangeOrdersMatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurExchangeOrdersMatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurExchangeOrdersMatched represents a OrdersMatched event raised by the BlurExchange contract.
type BlurExchangeOrdersMatched struct {
	Maker    common.Address
	Taker    common.Address
	Sell     Order
	SellHash [32]byte
	Buy      Order
	BuyHash  [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOrdersMatched is a free log retrieval operation binding the contract event 0x61cbb2a3dee0b6064c2e681aadd61677fb4ef319f0b547508d495626f5a62f64.
//
// Solidity: event OrdersMatched(address indexed maker, address indexed taker, (address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) sell, bytes32 sellHash, (address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) buy, bytes32 buyHash)
func (_BlurExchange *BlurExchangeFilterer) FilterOrdersMatched(opts *bind.FilterOpts, maker []common.Address, taker []common.Address) (*BlurExchangeOrdersMatchedIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _BlurExchange.contract.FilterLogs(opts, "OrdersMatched", makerRule, takerRule)
	if err != nil {
		return nil, err
	}
	return &BlurExchangeOrdersMatchedIterator{contract: _BlurExchange.contract, event: "OrdersMatched", logs: logs, sub: sub}, nil
}

// WatchOrdersMatched is a free log subscription operation binding the contract event 0x61cbb2a3dee0b6064c2e681aadd61677fb4ef319f0b547508d495626f5a62f64.
//
// Solidity: event OrdersMatched(address indexed maker, address indexed taker, (address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) sell, bytes32 sellHash, (address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) buy, bytes32 buyHash)
func (_BlurExchange *BlurExchangeFilterer) WatchOrdersMatched(opts *bind.WatchOpts, sink chan<- *BlurExchangeOrdersMatched, maker []common.Address, taker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _BlurExchange.contract.WatchLogs(opts, "OrdersMatched", makerRule, takerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurExchangeOrdersMatched)
				if err := _BlurExchange.contract.UnpackLog(event, "OrdersMatched", log); err != nil {
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

// ParseOrdersMatched is a log parse operation binding the contract event 0x61cbb2a3dee0b6064c2e681aadd61677fb4ef319f0b547508d495626f5a62f64.
//
// Solidity: event OrdersMatched(address indexed maker, address indexed taker, (address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) sell, bytes32 sellHash, (address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) buy, bytes32 buyHash)
func (_BlurExchange *BlurExchangeFilterer) ParseOrdersMatched(log types.Log) (*BlurExchangeOrdersMatched, error) {
	event := new(BlurExchangeOrdersMatched)
	if err := _BlurExchange.contract.UnpackLog(event, "OrdersMatched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurExchangeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BlurExchange contract.
type BlurExchangeOwnershipTransferredIterator struct {
	Event *BlurExchangeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BlurExchangeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurExchangeOwnershipTransferred)
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
		it.Event = new(BlurExchangeOwnershipTransferred)
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
func (it *BlurExchangeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurExchangeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurExchangeOwnershipTransferred represents a OwnershipTransferred event raised by the BlurExchange contract.
type BlurExchangeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BlurExchange *BlurExchangeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BlurExchangeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BlurExchange.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BlurExchangeOwnershipTransferredIterator{contract: _BlurExchange.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BlurExchange *BlurExchangeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BlurExchangeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BlurExchange.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurExchangeOwnershipTransferred)
				if err := _BlurExchange.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BlurExchange *BlurExchangeFilterer) ParseOwnershipTransferred(log types.Log) (*BlurExchangeOwnershipTransferred, error) {
	event := new(BlurExchangeOwnershipTransferred)
	if err := _BlurExchange.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurExchangeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the BlurExchange contract.
type BlurExchangeUpgradedIterator struct {
	Event *BlurExchangeUpgraded // Event containing the contract specifics and raw log

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
func (it *BlurExchangeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurExchangeUpgraded)
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
		it.Event = new(BlurExchangeUpgraded)
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
func (it *BlurExchangeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurExchangeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurExchangeUpgraded represents a Upgraded event raised by the BlurExchange contract.
type BlurExchangeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BlurExchange *BlurExchangeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BlurExchangeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BlurExchange.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BlurExchangeUpgradedIterator{contract: _BlurExchange.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BlurExchange *BlurExchangeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BlurExchangeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BlurExchange.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurExchangeUpgraded)
				if err := _BlurExchange.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BlurExchange *BlurExchangeFilterer) ParseUpgraded(log types.Log) (*BlurExchangeUpgraded, error) {
	event := new(BlurExchangeUpgraded)
	if err := _BlurExchange.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContextUpgradeableMetaData contains all meta data concerning the ContextUpgradeable contract.
var ContextUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"}]",
}

// ContextUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use ContextUpgradeableMetaData.ABI instead.
var ContextUpgradeableABI = ContextUpgradeableMetaData.ABI

// ContextUpgradeable is an auto generated Go binding around an Ethereum contract.
type ContextUpgradeable struct {
	ContextUpgradeableCaller     // Read-only binding to the contract
	ContextUpgradeableTransactor // Write-only binding to the contract
	ContextUpgradeableFilterer   // Log filterer for contract events
}

// ContextUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextUpgradeableSession struct {
	Contract     *ContextUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContextUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextUpgradeableCallerSession struct {
	Contract *ContextUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ContextUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextUpgradeableTransactorSession struct {
	Contract     *ContextUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ContextUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextUpgradeableRaw struct {
	Contract *ContextUpgradeable // Generic contract binding to access the raw methods on
}

// ContextUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextUpgradeableCallerRaw struct {
	Contract *ContextUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// ContextUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextUpgradeableTransactorRaw struct {
	Contract *ContextUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContextUpgradeable creates a new instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeable(address common.Address, backend bind.ContractBackend) (*ContextUpgradeable, error) {
	contract, err := bindContextUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeable{ContextUpgradeableCaller: ContextUpgradeableCaller{contract: contract}, ContextUpgradeableTransactor: ContextUpgradeableTransactor{contract: contract}, ContextUpgradeableFilterer: ContextUpgradeableFilterer{contract: contract}}, nil
}

// NewContextUpgradeableCaller creates a new read-only instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*ContextUpgradeableCaller, error) {
	contract, err := bindContextUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableCaller{contract: contract}, nil
}

// NewContextUpgradeableTransactor creates a new write-only instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextUpgradeableTransactor, error) {
	contract, err := bindContextUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableTransactor{contract: contract}, nil
}

// NewContextUpgradeableFilterer creates a new log filterer instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextUpgradeableFilterer, error) {
	contract, err := bindContextUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableFilterer{contract: contract}, nil
}

// bindContextUpgradeable binds a generic wrapper to an already deployed contract.
func bindContextUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContextUpgradeable *ContextUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContextUpgradeable.Contract.ContextUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContextUpgradeable *ContextUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.ContextUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContextUpgradeable *ContextUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.ContextUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContextUpgradeable *ContextUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContextUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContextUpgradeable *ContextUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContextUpgradeable *ContextUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// ContextUpgradeableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContextUpgradeable contract.
type ContextUpgradeableInitializedIterator struct {
	Event *ContextUpgradeableInitialized // Event containing the contract specifics and raw log

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
func (it *ContextUpgradeableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContextUpgradeableInitialized)
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
		it.Event = new(ContextUpgradeableInitialized)
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
func (it *ContextUpgradeableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContextUpgradeableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContextUpgradeableInitialized represents a Initialized event raised by the ContextUpgradeable contract.
type ContextUpgradeableInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContextUpgradeable *ContextUpgradeableFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContextUpgradeableInitializedIterator, error) {

	logs, sub, err := _ContextUpgradeable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableInitializedIterator{contract: _ContextUpgradeable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContextUpgradeable *ContextUpgradeableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContextUpgradeableInitialized) (event.Subscription, error) {

	logs, sub, err := _ContextUpgradeable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContextUpgradeableInitialized)
				if err := _ContextUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContextUpgradeable *ContextUpgradeableFilterer) ParseInitialized(log types.Log) (*ContextUpgradeableInitialized, error) {
	event := new(ContextUpgradeableInitialized)
	if err := _ContextUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EIP712MetaData contains all meta data concerning the EIP712 contract.
var EIP712MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEE_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ORACLE_ORDER_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ORDER_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROOT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610149806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80631d97c9bb1461005c57806331e6d0fe146100955780633644e515146100bc5780634832ede1146100c5578063f973a209146100ec575b600080fd5b6100837fd71080023d2f293ed0723dc287d6b2d4e7d27d0b6c12928e300721b7c78c748581565b60405190815260200160405180910390f35b6100837f5bcf4b2eaff7fcdeb49f0bda53026b9ebdd93db566fe4c447125cb899e598c9081565b61008360005481565b6100837f05b43f730f67de334a342883f867101fc7ef3361dfdff4a29a7aa97e0920ef7a81565b6100837f376bfbc394a7ba7fdf10f224572cef371358e3053e362f4554fcd2ad56329b3f8156fea2646970667358221220fe30d6e3feda75d1628570872801e1854145310c0ab83b35a6f888c9f8c2890c64736f6c63430008110033",
}

// EIP712ABI is the input ABI used to generate the binding from.
// Deprecated: Use EIP712MetaData.ABI instead.
var EIP712ABI = EIP712MetaData.ABI

// EIP712Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EIP712MetaData.Bin instead.
var EIP712Bin = EIP712MetaData.Bin

// DeployEIP712 deploys a new Ethereum contract, binding an instance of EIP712 to it.
func DeployEIP712(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EIP712, error) {
	parsed, err := EIP712MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EIP712Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EIP712{EIP712Caller: EIP712Caller{contract: contract}, EIP712Transactor: EIP712Transactor{contract: contract}, EIP712Filterer: EIP712Filterer{contract: contract}}, nil
}

// EIP712 is an auto generated Go binding around an Ethereum contract.
type EIP712 struct {
	EIP712Caller     // Read-only binding to the contract
	EIP712Transactor // Write-only binding to the contract
	EIP712Filterer   // Log filterer for contract events
}

// EIP712Caller is an auto generated read-only Go binding around an Ethereum contract.
type EIP712Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP712Transactor is an auto generated write-only Go binding around an Ethereum contract.
type EIP712Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP712Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EIP712Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP712Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EIP712Session struct {
	Contract     *EIP712           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EIP712CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EIP712CallerSession struct {
	Contract *EIP712Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EIP712TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EIP712TransactorSession struct {
	Contract     *EIP712Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EIP712Raw is an auto generated low-level Go binding around an Ethereum contract.
type EIP712Raw struct {
	Contract *EIP712 // Generic contract binding to access the raw methods on
}

// EIP712CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EIP712CallerRaw struct {
	Contract *EIP712Caller // Generic read-only contract binding to access the raw methods on
}

// EIP712TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EIP712TransactorRaw struct {
	Contract *EIP712Transactor // Generic write-only contract binding to access the raw methods on
}

// NewEIP712 creates a new instance of EIP712, bound to a specific deployed contract.
func NewEIP712(address common.Address, backend bind.ContractBackend) (*EIP712, error) {
	contract, err := bindEIP712(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EIP712{EIP712Caller: EIP712Caller{contract: contract}, EIP712Transactor: EIP712Transactor{contract: contract}, EIP712Filterer: EIP712Filterer{contract: contract}}, nil
}

// NewEIP712Caller creates a new read-only instance of EIP712, bound to a specific deployed contract.
func NewEIP712Caller(address common.Address, caller bind.ContractCaller) (*EIP712Caller, error) {
	contract, err := bindEIP712(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EIP712Caller{contract: contract}, nil
}

// NewEIP712Transactor creates a new write-only instance of EIP712, bound to a specific deployed contract.
func NewEIP712Transactor(address common.Address, transactor bind.ContractTransactor) (*EIP712Transactor, error) {
	contract, err := bindEIP712(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EIP712Transactor{contract: contract}, nil
}

// NewEIP712Filterer creates a new log filterer instance of EIP712, bound to a specific deployed contract.
func NewEIP712Filterer(address common.Address, filterer bind.ContractFilterer) (*EIP712Filterer, error) {
	contract, err := bindEIP712(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EIP712Filterer{contract: contract}, nil
}

// bindEIP712 binds a generic wrapper to an already deployed contract.
func bindEIP712(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EIP712ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EIP712 *EIP712Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EIP712.Contract.EIP712Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EIP712 *EIP712Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EIP712.Contract.EIP712Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EIP712 *EIP712Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EIP712.Contract.EIP712Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EIP712 *EIP712CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EIP712.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EIP712 *EIP712TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EIP712.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EIP712 *EIP712TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EIP712.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_EIP712 *EIP712Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EIP712.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_EIP712 *EIP712Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _EIP712.Contract.DOMAINSEPARATOR(&_EIP712.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_EIP712 *EIP712CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _EIP712.Contract.DOMAINSEPARATOR(&_EIP712.CallOpts)
}

// FEETYPEHASH is a free data retrieval call binding the contract method 0x4832ede1.
//
// Solidity: function FEE_TYPEHASH() view returns(bytes32)
func (_EIP712 *EIP712Caller) FEETYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EIP712.contract.Call(opts, &out, "FEE_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FEETYPEHASH is a free data retrieval call binding the contract method 0x4832ede1.
//
// Solidity: function FEE_TYPEHASH() view returns(bytes32)
func (_EIP712 *EIP712Session) FEETYPEHASH() ([32]byte, error) {
	return _EIP712.Contract.FEETYPEHASH(&_EIP712.CallOpts)
}

// FEETYPEHASH is a free data retrieval call binding the contract method 0x4832ede1.
//
// Solidity: function FEE_TYPEHASH() view returns(bytes32)
func (_EIP712 *EIP712CallerSession) FEETYPEHASH() ([32]byte, error) {
	return _EIP712.Contract.FEETYPEHASH(&_EIP712.CallOpts)
}

// ORACLEORDERTYPEHASH is a free data retrieval call binding the contract method 0x1d97c9bb.
//
// Solidity: function ORACLE_ORDER_TYPEHASH() view returns(bytes32)
func (_EIP712 *EIP712Caller) ORACLEORDERTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EIP712.contract.Call(opts, &out, "ORACLE_ORDER_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ORACLEORDERTYPEHASH is a free data retrieval call binding the contract method 0x1d97c9bb.
//
// Solidity: function ORACLE_ORDER_TYPEHASH() view returns(bytes32)
func (_EIP712 *EIP712Session) ORACLEORDERTYPEHASH() ([32]byte, error) {
	return _EIP712.Contract.ORACLEORDERTYPEHASH(&_EIP712.CallOpts)
}

// ORACLEORDERTYPEHASH is a free data retrieval call binding the contract method 0x1d97c9bb.
//
// Solidity: function ORACLE_ORDER_TYPEHASH() view returns(bytes32)
func (_EIP712 *EIP712CallerSession) ORACLEORDERTYPEHASH() ([32]byte, error) {
	return _EIP712.Contract.ORACLEORDERTYPEHASH(&_EIP712.CallOpts)
}

// ORDERTYPEHASH is a free data retrieval call binding the contract method 0xf973a209.
//
// Solidity: function ORDER_TYPEHASH() view returns(bytes32)
func (_EIP712 *EIP712Caller) ORDERTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EIP712.contract.Call(opts, &out, "ORDER_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ORDERTYPEHASH is a free data retrieval call binding the contract method 0xf973a209.
//
// Solidity: function ORDER_TYPEHASH() view returns(bytes32)
func (_EIP712 *EIP712Session) ORDERTYPEHASH() ([32]byte, error) {
	return _EIP712.Contract.ORDERTYPEHASH(&_EIP712.CallOpts)
}

// ORDERTYPEHASH is a free data retrieval call binding the contract method 0xf973a209.
//
// Solidity: function ORDER_TYPEHASH() view returns(bytes32)
func (_EIP712 *EIP712CallerSession) ORDERTYPEHASH() ([32]byte, error) {
	return _EIP712.Contract.ORDERTYPEHASH(&_EIP712.CallOpts)
}

// ROOTTYPEHASH is a free data retrieval call binding the contract method 0x31e6d0fe.
//
// Solidity: function ROOT_TYPEHASH() view returns(bytes32)
func (_EIP712 *EIP712Caller) ROOTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EIP712.contract.Call(opts, &out, "ROOT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ROOTTYPEHASH is a free data retrieval call binding the contract method 0x31e6d0fe.
//
// Solidity: function ROOT_TYPEHASH() view returns(bytes32)
func (_EIP712 *EIP712Session) ROOTTYPEHASH() ([32]byte, error) {
	return _EIP712.Contract.ROOTTYPEHASH(&_EIP712.CallOpts)
}

// ROOTTYPEHASH is a free data retrieval call binding the contract method 0x31e6d0fe.
//
// Solidity: function ROOT_TYPEHASH() view returns(bytes32)
func (_EIP712 *EIP712CallerSession) ROOTTYPEHASH() ([32]byte, error) {
	return _EIP712.Contract.ROOTTYPEHASH(&_EIP712.CallOpts)
}

// ERC1967UpgradeUpgradeableMetaData contains all meta data concerning the ERC1967UpgradeUpgradeable contract.
var ERC1967UpgradeUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"}]",
}

// ERC1967UpgradeUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC1967UpgradeUpgradeableMetaData.ABI instead.
var ERC1967UpgradeUpgradeableABI = ERC1967UpgradeUpgradeableMetaData.ABI

// ERC1967UpgradeUpgradeable is an auto generated Go binding around an Ethereum contract.
type ERC1967UpgradeUpgradeable struct {
	ERC1967UpgradeUpgradeableCaller     // Read-only binding to the contract
	ERC1967UpgradeUpgradeableTransactor // Write-only binding to the contract
	ERC1967UpgradeUpgradeableFilterer   // Log filterer for contract events
}

// ERC1967UpgradeUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC1967UpgradeUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1967UpgradeUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC1967UpgradeUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1967UpgradeUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC1967UpgradeUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1967UpgradeUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC1967UpgradeUpgradeableSession struct {
	Contract     *ERC1967UpgradeUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ERC1967UpgradeUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC1967UpgradeUpgradeableCallerSession struct {
	Contract *ERC1967UpgradeUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// ERC1967UpgradeUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC1967UpgradeUpgradeableTransactorSession struct {
	Contract     *ERC1967UpgradeUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// ERC1967UpgradeUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC1967UpgradeUpgradeableRaw struct {
	Contract *ERC1967UpgradeUpgradeable // Generic contract binding to access the raw methods on
}

// ERC1967UpgradeUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC1967UpgradeUpgradeableCallerRaw struct {
	Contract *ERC1967UpgradeUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// ERC1967UpgradeUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC1967UpgradeUpgradeableTransactorRaw struct {
	Contract *ERC1967UpgradeUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC1967UpgradeUpgradeable creates a new instance of ERC1967UpgradeUpgradeable, bound to a specific deployed contract.
func NewERC1967UpgradeUpgradeable(address common.Address, backend bind.ContractBackend) (*ERC1967UpgradeUpgradeable, error) {
	contract, err := bindERC1967UpgradeUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeUpgradeable{ERC1967UpgradeUpgradeableCaller: ERC1967UpgradeUpgradeableCaller{contract: contract}, ERC1967UpgradeUpgradeableTransactor: ERC1967UpgradeUpgradeableTransactor{contract: contract}, ERC1967UpgradeUpgradeableFilterer: ERC1967UpgradeUpgradeableFilterer{contract: contract}}, nil
}

// NewERC1967UpgradeUpgradeableCaller creates a new read-only instance of ERC1967UpgradeUpgradeable, bound to a specific deployed contract.
func NewERC1967UpgradeUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*ERC1967UpgradeUpgradeableCaller, error) {
	contract, err := bindERC1967UpgradeUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeUpgradeableCaller{contract: contract}, nil
}

// NewERC1967UpgradeUpgradeableTransactor creates a new write-only instance of ERC1967UpgradeUpgradeable, bound to a specific deployed contract.
func NewERC1967UpgradeUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC1967UpgradeUpgradeableTransactor, error) {
	contract, err := bindERC1967UpgradeUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeUpgradeableTransactor{contract: contract}, nil
}

// NewERC1967UpgradeUpgradeableFilterer creates a new log filterer instance of ERC1967UpgradeUpgradeable, bound to a specific deployed contract.
func NewERC1967UpgradeUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC1967UpgradeUpgradeableFilterer, error) {
	contract, err := bindERC1967UpgradeUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeUpgradeableFilterer{contract: contract}, nil
}

// bindERC1967UpgradeUpgradeable binds a generic wrapper to an already deployed contract.
func bindERC1967UpgradeUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC1967UpgradeUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1967UpgradeUpgradeable.Contract.ERC1967UpgradeUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1967UpgradeUpgradeable.Contract.ERC1967UpgradeUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1967UpgradeUpgradeable.Contract.ERC1967UpgradeUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1967UpgradeUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1967UpgradeUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1967UpgradeUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// ERC1967UpgradeUpgradeableAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the ERC1967UpgradeUpgradeable contract.
type ERC1967UpgradeUpgradeableAdminChangedIterator struct {
	Event *ERC1967UpgradeUpgradeableAdminChanged // Event containing the contract specifics and raw log

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
func (it *ERC1967UpgradeUpgradeableAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1967UpgradeUpgradeableAdminChanged)
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
		it.Event = new(ERC1967UpgradeUpgradeableAdminChanged)
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
func (it *ERC1967UpgradeUpgradeableAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1967UpgradeUpgradeableAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1967UpgradeUpgradeableAdminChanged represents a AdminChanged event raised by the ERC1967UpgradeUpgradeable contract.
type ERC1967UpgradeUpgradeableAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*ERC1967UpgradeUpgradeableAdminChangedIterator, error) {

	logs, sub, err := _ERC1967UpgradeUpgradeable.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeUpgradeableAdminChangedIterator{contract: _ERC1967UpgradeUpgradeable.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ERC1967UpgradeUpgradeableAdminChanged) (event.Subscription, error) {

	logs, sub, err := _ERC1967UpgradeUpgradeable.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1967UpgradeUpgradeableAdminChanged)
				if err := _ERC1967UpgradeUpgradeable.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) ParseAdminChanged(log types.Log) (*ERC1967UpgradeUpgradeableAdminChanged, error) {
	event := new(ERC1967UpgradeUpgradeableAdminChanged)
	if err := _ERC1967UpgradeUpgradeable.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1967UpgradeUpgradeableBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the ERC1967UpgradeUpgradeable contract.
type ERC1967UpgradeUpgradeableBeaconUpgradedIterator struct {
	Event *ERC1967UpgradeUpgradeableBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *ERC1967UpgradeUpgradeableBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1967UpgradeUpgradeableBeaconUpgraded)
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
		it.Event = new(ERC1967UpgradeUpgradeableBeaconUpgraded)
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
func (it *ERC1967UpgradeUpgradeableBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1967UpgradeUpgradeableBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1967UpgradeUpgradeableBeaconUpgraded represents a BeaconUpgraded event raised by the ERC1967UpgradeUpgradeable contract.
type ERC1967UpgradeUpgradeableBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*ERC1967UpgradeUpgradeableBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ERC1967UpgradeUpgradeable.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeUpgradeableBeaconUpgradedIterator{contract: _ERC1967UpgradeUpgradeable.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *ERC1967UpgradeUpgradeableBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ERC1967UpgradeUpgradeable.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1967UpgradeUpgradeableBeaconUpgraded)
				if err := _ERC1967UpgradeUpgradeable.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) ParseBeaconUpgraded(log types.Log) (*ERC1967UpgradeUpgradeableBeaconUpgraded, error) {
	event := new(ERC1967UpgradeUpgradeableBeaconUpgraded)
	if err := _ERC1967UpgradeUpgradeable.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1967UpgradeUpgradeableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ERC1967UpgradeUpgradeable contract.
type ERC1967UpgradeUpgradeableInitializedIterator struct {
	Event *ERC1967UpgradeUpgradeableInitialized // Event containing the contract specifics and raw log

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
func (it *ERC1967UpgradeUpgradeableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1967UpgradeUpgradeableInitialized)
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
		it.Event = new(ERC1967UpgradeUpgradeableInitialized)
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
func (it *ERC1967UpgradeUpgradeableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1967UpgradeUpgradeableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1967UpgradeUpgradeableInitialized represents a Initialized event raised by the ERC1967UpgradeUpgradeable contract.
type ERC1967UpgradeUpgradeableInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) FilterInitialized(opts *bind.FilterOpts) (*ERC1967UpgradeUpgradeableInitializedIterator, error) {

	logs, sub, err := _ERC1967UpgradeUpgradeable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeUpgradeableInitializedIterator{contract: _ERC1967UpgradeUpgradeable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ERC1967UpgradeUpgradeableInitialized) (event.Subscription, error) {

	logs, sub, err := _ERC1967UpgradeUpgradeable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1967UpgradeUpgradeableInitialized)
				if err := _ERC1967UpgradeUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) ParseInitialized(log types.Log) (*ERC1967UpgradeUpgradeableInitialized, error) {
	event := new(ERC1967UpgradeUpgradeableInitialized)
	if err := _ERC1967UpgradeUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1967UpgradeUpgradeableUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ERC1967UpgradeUpgradeable contract.
type ERC1967UpgradeUpgradeableUpgradedIterator struct {
	Event *ERC1967UpgradeUpgradeableUpgraded // Event containing the contract specifics and raw log

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
func (it *ERC1967UpgradeUpgradeableUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1967UpgradeUpgradeableUpgraded)
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
		it.Event = new(ERC1967UpgradeUpgradeableUpgraded)
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
func (it *ERC1967UpgradeUpgradeableUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1967UpgradeUpgradeableUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1967UpgradeUpgradeableUpgraded represents a Upgraded event raised by the ERC1967UpgradeUpgradeable contract.
type ERC1967UpgradeUpgradeableUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ERC1967UpgradeUpgradeableUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ERC1967UpgradeUpgradeable.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeUpgradeableUpgradedIterator{contract: _ERC1967UpgradeUpgradeable.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ERC1967UpgradeUpgradeableUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ERC1967UpgradeUpgradeable.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1967UpgradeUpgradeableUpgraded)
				if err := _ERC1967UpgradeUpgradeable.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) ParseUpgraded(log types.Log) (*ERC1967UpgradeUpgradeableUpgraded, error) {
	event := new(ERC1967UpgradeUpgradeableUpgraded)
	if err := _ERC1967UpgradeUpgradeable.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBeaconUpgradeableMetaData contains all meta data concerning the IBeaconUpgradeable contract.
var IBeaconUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IBeaconUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use IBeaconUpgradeableMetaData.ABI instead.
var IBeaconUpgradeableABI = IBeaconUpgradeableMetaData.ABI

// IBeaconUpgradeable is an auto generated Go binding around an Ethereum contract.
type IBeaconUpgradeable struct {
	IBeaconUpgradeableCaller     // Read-only binding to the contract
	IBeaconUpgradeableTransactor // Write-only binding to the contract
	IBeaconUpgradeableFilterer   // Log filterer for contract events
}

// IBeaconUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBeaconUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBeaconUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBeaconUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBeaconUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBeaconUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBeaconUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBeaconUpgradeableSession struct {
	Contract     *IBeaconUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IBeaconUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBeaconUpgradeableCallerSession struct {
	Contract *IBeaconUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IBeaconUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBeaconUpgradeableTransactorSession struct {
	Contract     *IBeaconUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IBeaconUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBeaconUpgradeableRaw struct {
	Contract *IBeaconUpgradeable // Generic contract binding to access the raw methods on
}

// IBeaconUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBeaconUpgradeableCallerRaw struct {
	Contract *IBeaconUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// IBeaconUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBeaconUpgradeableTransactorRaw struct {
	Contract *IBeaconUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBeaconUpgradeable creates a new instance of IBeaconUpgradeable, bound to a specific deployed contract.
func NewIBeaconUpgradeable(address common.Address, backend bind.ContractBackend) (*IBeaconUpgradeable, error) {
	contract, err := bindIBeaconUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBeaconUpgradeable{IBeaconUpgradeableCaller: IBeaconUpgradeableCaller{contract: contract}, IBeaconUpgradeableTransactor: IBeaconUpgradeableTransactor{contract: contract}, IBeaconUpgradeableFilterer: IBeaconUpgradeableFilterer{contract: contract}}, nil
}

// NewIBeaconUpgradeableCaller creates a new read-only instance of IBeaconUpgradeable, bound to a specific deployed contract.
func NewIBeaconUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*IBeaconUpgradeableCaller, error) {
	contract, err := bindIBeaconUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBeaconUpgradeableCaller{contract: contract}, nil
}

// NewIBeaconUpgradeableTransactor creates a new write-only instance of IBeaconUpgradeable, bound to a specific deployed contract.
func NewIBeaconUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*IBeaconUpgradeableTransactor, error) {
	contract, err := bindIBeaconUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBeaconUpgradeableTransactor{contract: contract}, nil
}

// NewIBeaconUpgradeableFilterer creates a new log filterer instance of IBeaconUpgradeable, bound to a specific deployed contract.
func NewIBeaconUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*IBeaconUpgradeableFilterer, error) {
	contract, err := bindIBeaconUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBeaconUpgradeableFilterer{contract: contract}, nil
}

// bindIBeaconUpgradeable binds a generic wrapper to an already deployed contract.
func bindIBeaconUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBeaconUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBeaconUpgradeable *IBeaconUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBeaconUpgradeable.Contract.IBeaconUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBeaconUpgradeable *IBeaconUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBeaconUpgradeable.Contract.IBeaconUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBeaconUpgradeable *IBeaconUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBeaconUpgradeable.Contract.IBeaconUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBeaconUpgradeable *IBeaconUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBeaconUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBeaconUpgradeable *IBeaconUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBeaconUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBeaconUpgradeable *IBeaconUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBeaconUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_IBeaconUpgradeable *IBeaconUpgradeableCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IBeaconUpgradeable.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_IBeaconUpgradeable *IBeaconUpgradeableSession) Implementation() (common.Address, error) {
	return _IBeaconUpgradeable.Contract.Implementation(&_IBeaconUpgradeable.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_IBeaconUpgradeable *IBeaconUpgradeableCallerSession) Implementation() (common.Address, error) {
	return _IBeaconUpgradeable.Contract.Implementation(&_IBeaconUpgradeable.CallOpts)
}

// IBlurExchangeMetaData contains all meta data concerning the IBlurExchange contract.
var IBlurExchangeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"name\":\"cancelOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"close\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"sell\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"buy\",\"type\":\"tuple\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incrementNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIExecutionDelegate\",\"name\":\"_executionDelegate\",\"type\":\"address\"},{\"internalType\":\"contractIPolicyManager\",\"name\":\"_policyManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockRange\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockRange\",\"type\":\"uint256\"}],\"name\":\"setBlockRange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIExecutionDelegate\",\"name\":\"_executionDelegate\",\"type\":\"address\"}],\"name\":\"setExecutionDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPolicyManager\",\"name\":\"_policyManager\",\"type\":\"address\"}],\"name\":\"setPolicyManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IBlurExchangeABI is the input ABI used to generate the binding from.
// Deprecated: Use IBlurExchangeMetaData.ABI instead.
var IBlurExchangeABI = IBlurExchangeMetaData.ABI

// IBlurExchange is an auto generated Go binding around an Ethereum contract.
type IBlurExchange struct {
	IBlurExchangeCaller     // Read-only binding to the contract
	IBlurExchangeTransactor // Write-only binding to the contract
	IBlurExchangeFilterer   // Log filterer for contract events
}

// IBlurExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBlurExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBlurExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBlurExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBlurExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBlurExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBlurExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBlurExchangeSession struct {
	Contract     *IBlurExchange    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBlurExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBlurExchangeCallerSession struct {
	Contract *IBlurExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IBlurExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBlurExchangeTransactorSession struct {
	Contract     *IBlurExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IBlurExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBlurExchangeRaw struct {
	Contract *IBlurExchange // Generic contract binding to access the raw methods on
}

// IBlurExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBlurExchangeCallerRaw struct {
	Contract *IBlurExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// IBlurExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBlurExchangeTransactorRaw struct {
	Contract *IBlurExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBlurExchange creates a new instance of IBlurExchange, bound to a specific deployed contract.
func NewIBlurExchange(address common.Address, backend bind.ContractBackend) (*IBlurExchange, error) {
	contract, err := bindIBlurExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBlurExchange{IBlurExchangeCaller: IBlurExchangeCaller{contract: contract}, IBlurExchangeTransactor: IBlurExchangeTransactor{contract: contract}, IBlurExchangeFilterer: IBlurExchangeFilterer{contract: contract}}, nil
}

// NewIBlurExchangeCaller creates a new read-only instance of IBlurExchange, bound to a specific deployed contract.
func NewIBlurExchangeCaller(address common.Address, caller bind.ContractCaller) (*IBlurExchangeCaller, error) {
	contract, err := bindIBlurExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBlurExchangeCaller{contract: contract}, nil
}

// NewIBlurExchangeTransactor creates a new write-only instance of IBlurExchange, bound to a specific deployed contract.
func NewIBlurExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*IBlurExchangeTransactor, error) {
	contract, err := bindIBlurExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBlurExchangeTransactor{contract: contract}, nil
}

// NewIBlurExchangeFilterer creates a new log filterer instance of IBlurExchange, bound to a specific deployed contract.
func NewIBlurExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*IBlurExchangeFilterer, error) {
	contract, err := bindIBlurExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBlurExchangeFilterer{contract: contract}, nil
}

// bindIBlurExchange binds a generic wrapper to an already deployed contract.
func bindIBlurExchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBlurExchangeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBlurExchange *IBlurExchangeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBlurExchange.Contract.IBlurExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBlurExchange *IBlurExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBlurExchange.Contract.IBlurExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBlurExchange *IBlurExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBlurExchange.Contract.IBlurExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBlurExchange *IBlurExchangeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBlurExchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBlurExchange *IBlurExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBlurExchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBlurExchange *IBlurExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBlurExchange.Contract.contract.Transact(opts, method, params...)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_IBlurExchange *IBlurExchangeCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IBlurExchange.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_IBlurExchange *IBlurExchangeSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _IBlurExchange.Contract.Nonces(&_IBlurExchange.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_IBlurExchange *IBlurExchangeCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _IBlurExchange.Contract.Nonces(&_IBlurExchange.CallOpts, arg0)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xf4acd740.
//
// Solidity: function cancelOrder((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) order) returns()
func (_IBlurExchange *IBlurExchangeTransactor) CancelOrder(opts *bind.TransactOpts, order Order) (*types.Transaction, error) {
	return _IBlurExchange.contract.Transact(opts, "cancelOrder", order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xf4acd740.
//
// Solidity: function cancelOrder((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) order) returns()
func (_IBlurExchange *IBlurExchangeSession) CancelOrder(order Order) (*types.Transaction, error) {
	return _IBlurExchange.Contract.CancelOrder(&_IBlurExchange.TransactOpts, order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xf4acd740.
//
// Solidity: function cancelOrder((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) order) returns()
func (_IBlurExchange *IBlurExchangeTransactorSession) CancelOrder(order Order) (*types.Transaction, error) {
	return _IBlurExchange.Contract.CancelOrder(&_IBlurExchange.TransactOpts, order)
}

// CancelOrders is a paid mutator transaction binding the contract method 0xab7e8cba.
//
// Solidity: function cancelOrders((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes)[] orders) returns()
func (_IBlurExchange *IBlurExchangeTransactor) CancelOrders(opts *bind.TransactOpts, orders []Order) (*types.Transaction, error) {
	return _IBlurExchange.contract.Transact(opts, "cancelOrders", orders)
}

// CancelOrders is a paid mutator transaction binding the contract method 0xab7e8cba.
//
// Solidity: function cancelOrders((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes)[] orders) returns()
func (_IBlurExchange *IBlurExchangeSession) CancelOrders(orders []Order) (*types.Transaction, error) {
	return _IBlurExchange.Contract.CancelOrders(&_IBlurExchange.TransactOpts, orders)
}

// CancelOrders is a paid mutator transaction binding the contract method 0xab7e8cba.
//
// Solidity: function cancelOrders((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes)[] orders) returns()
func (_IBlurExchange *IBlurExchangeTransactorSession) CancelOrders(orders []Order) (*types.Transaction, error) {
	return _IBlurExchange.Contract.CancelOrders(&_IBlurExchange.TransactOpts, orders)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_IBlurExchange *IBlurExchangeTransactor) Close(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBlurExchange.contract.Transact(opts, "close")
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_IBlurExchange *IBlurExchangeSession) Close() (*types.Transaction, error) {
	return _IBlurExchange.Contract.Close(&_IBlurExchange.TransactOpts)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_IBlurExchange *IBlurExchangeTransactorSession) Close() (*types.Transaction, error) {
	return _IBlurExchange.Contract.Close(&_IBlurExchange.TransactOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x9a1fc3a7.
//
// Solidity: function execute(((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) sell, ((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) buy) payable returns()
func (_IBlurExchange *IBlurExchangeTransactor) Execute(opts *bind.TransactOpts, sell Input, buy Input) (*types.Transaction, error) {
	return _IBlurExchange.contract.Transact(opts, "execute", sell, buy)
}

// Execute is a paid mutator transaction binding the contract method 0x9a1fc3a7.
//
// Solidity: function execute(((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) sell, ((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) buy) payable returns()
func (_IBlurExchange *IBlurExchangeSession) Execute(sell Input, buy Input) (*types.Transaction, error) {
	return _IBlurExchange.Contract.Execute(&_IBlurExchange.TransactOpts, sell, buy)
}

// Execute is a paid mutator transaction binding the contract method 0x9a1fc3a7.
//
// Solidity: function execute(((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) sell, ((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) buy) payable returns()
func (_IBlurExchange *IBlurExchangeTransactorSession) Execute(sell Input, buy Input) (*types.Transaction, error) {
	return _IBlurExchange.Contract.Execute(&_IBlurExchange.TransactOpts, sell, buy)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns()
func (_IBlurExchange *IBlurExchangeTransactor) IncrementNonce(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBlurExchange.contract.Transact(opts, "incrementNonce")
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns()
func (_IBlurExchange *IBlurExchangeSession) IncrementNonce() (*types.Transaction, error) {
	return _IBlurExchange.Contract.IncrementNonce(&_IBlurExchange.TransactOpts)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns()
func (_IBlurExchange *IBlurExchangeTransactorSession) IncrementNonce() (*types.Transaction, error) {
	return _IBlurExchange.Contract.IncrementNonce(&_IBlurExchange.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address _executionDelegate, address _policyManager, address _oracle, uint256 _blockRange) returns()
func (_IBlurExchange *IBlurExchangeTransactor) Initialize(opts *bind.TransactOpts, _executionDelegate common.Address, _policyManager common.Address, _oracle common.Address, _blockRange *big.Int) (*types.Transaction, error) {
	return _IBlurExchange.contract.Transact(opts, "initialize", _executionDelegate, _policyManager, _oracle, _blockRange)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address _executionDelegate, address _policyManager, address _oracle, uint256 _blockRange) returns()
func (_IBlurExchange *IBlurExchangeSession) Initialize(_executionDelegate common.Address, _policyManager common.Address, _oracle common.Address, _blockRange *big.Int) (*types.Transaction, error) {
	return _IBlurExchange.Contract.Initialize(&_IBlurExchange.TransactOpts, _executionDelegate, _policyManager, _oracle, _blockRange)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address _executionDelegate, address _policyManager, address _oracle, uint256 _blockRange) returns()
func (_IBlurExchange *IBlurExchangeTransactorSession) Initialize(_executionDelegate common.Address, _policyManager common.Address, _oracle common.Address, _blockRange *big.Int) (*types.Transaction, error) {
	return _IBlurExchange.Contract.Initialize(&_IBlurExchange.TransactOpts, _executionDelegate, _policyManager, _oracle, _blockRange)
}

// SetBlockRange is a paid mutator transaction binding the contract method 0x6992aa36.
//
// Solidity: function setBlockRange(uint256 _blockRange) returns()
func (_IBlurExchange *IBlurExchangeTransactor) SetBlockRange(opts *bind.TransactOpts, _blockRange *big.Int) (*types.Transaction, error) {
	return _IBlurExchange.contract.Transact(opts, "setBlockRange", _blockRange)
}

// SetBlockRange is a paid mutator transaction binding the contract method 0x6992aa36.
//
// Solidity: function setBlockRange(uint256 _blockRange) returns()
func (_IBlurExchange *IBlurExchangeSession) SetBlockRange(_blockRange *big.Int) (*types.Transaction, error) {
	return _IBlurExchange.Contract.SetBlockRange(&_IBlurExchange.TransactOpts, _blockRange)
}

// SetBlockRange is a paid mutator transaction binding the contract method 0x6992aa36.
//
// Solidity: function setBlockRange(uint256 _blockRange) returns()
func (_IBlurExchange *IBlurExchangeTransactorSession) SetBlockRange(_blockRange *big.Int) (*types.Transaction, error) {
	return _IBlurExchange.Contract.SetBlockRange(&_IBlurExchange.TransactOpts, _blockRange)
}

// SetExecutionDelegate is a paid mutator transaction binding the contract method 0x037c9be2.
//
// Solidity: function setExecutionDelegate(address _executionDelegate) returns()
func (_IBlurExchange *IBlurExchangeTransactor) SetExecutionDelegate(opts *bind.TransactOpts, _executionDelegate common.Address) (*types.Transaction, error) {
	return _IBlurExchange.contract.Transact(opts, "setExecutionDelegate", _executionDelegate)
}

// SetExecutionDelegate is a paid mutator transaction binding the contract method 0x037c9be2.
//
// Solidity: function setExecutionDelegate(address _executionDelegate) returns()
func (_IBlurExchange *IBlurExchangeSession) SetExecutionDelegate(_executionDelegate common.Address) (*types.Transaction, error) {
	return _IBlurExchange.Contract.SetExecutionDelegate(&_IBlurExchange.TransactOpts, _executionDelegate)
}

// SetExecutionDelegate is a paid mutator transaction binding the contract method 0x037c9be2.
//
// Solidity: function setExecutionDelegate(address _executionDelegate) returns()
func (_IBlurExchange *IBlurExchangeTransactorSession) SetExecutionDelegate(_executionDelegate common.Address) (*types.Transaction, error) {
	return _IBlurExchange.Contract.SetExecutionDelegate(&_IBlurExchange.TransactOpts, _executionDelegate)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_IBlurExchange *IBlurExchangeTransactor) SetOracle(opts *bind.TransactOpts, _oracle common.Address) (*types.Transaction, error) {
	return _IBlurExchange.contract.Transact(opts, "setOracle", _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_IBlurExchange *IBlurExchangeSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _IBlurExchange.Contract.SetOracle(&_IBlurExchange.TransactOpts, _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_IBlurExchange *IBlurExchangeTransactorSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _IBlurExchange.Contract.SetOracle(&_IBlurExchange.TransactOpts, _oracle)
}

// SetPolicyManager is a paid mutator transaction binding the contract method 0xadde41e1.
//
// Solidity: function setPolicyManager(address _policyManager) returns()
func (_IBlurExchange *IBlurExchangeTransactor) SetPolicyManager(opts *bind.TransactOpts, _policyManager common.Address) (*types.Transaction, error) {
	return _IBlurExchange.contract.Transact(opts, "setPolicyManager", _policyManager)
}

// SetPolicyManager is a paid mutator transaction binding the contract method 0xadde41e1.
//
// Solidity: function setPolicyManager(address _policyManager) returns()
func (_IBlurExchange *IBlurExchangeSession) SetPolicyManager(_policyManager common.Address) (*types.Transaction, error) {
	return _IBlurExchange.Contract.SetPolicyManager(&_IBlurExchange.TransactOpts, _policyManager)
}

// SetPolicyManager is a paid mutator transaction binding the contract method 0xadde41e1.
//
// Solidity: function setPolicyManager(address _policyManager) returns()
func (_IBlurExchange *IBlurExchangeTransactorSession) SetPolicyManager(_policyManager common.Address) (*types.Transaction, error) {
	return _IBlurExchange.Contract.SetPolicyManager(&_IBlurExchange.TransactOpts, _policyManager)
}

// IBlurPoolMetaData contains all meta data concerning the IBlurPool contract.
var IBlurPoolMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IBlurPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use IBlurPoolMetaData.ABI instead.
var IBlurPoolABI = IBlurPoolMetaData.ABI

// IBlurPool is an auto generated Go binding around an Ethereum contract.
type IBlurPool struct {
	IBlurPoolCaller     // Read-only binding to the contract
	IBlurPoolTransactor // Write-only binding to the contract
	IBlurPoolFilterer   // Log filterer for contract events
}

// IBlurPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBlurPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBlurPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBlurPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBlurPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBlurPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBlurPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBlurPoolSession struct {
	Contract     *IBlurPool        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBlurPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBlurPoolCallerSession struct {
	Contract *IBlurPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IBlurPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBlurPoolTransactorSession struct {
	Contract     *IBlurPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IBlurPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBlurPoolRaw struct {
	Contract *IBlurPool // Generic contract binding to access the raw methods on
}

// IBlurPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBlurPoolCallerRaw struct {
	Contract *IBlurPoolCaller // Generic read-only contract binding to access the raw methods on
}

// IBlurPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBlurPoolTransactorRaw struct {
	Contract *IBlurPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBlurPool creates a new instance of IBlurPool, bound to a specific deployed contract.
func NewIBlurPool(address common.Address, backend bind.ContractBackend) (*IBlurPool, error) {
	contract, err := bindIBlurPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBlurPool{IBlurPoolCaller: IBlurPoolCaller{contract: contract}, IBlurPoolTransactor: IBlurPoolTransactor{contract: contract}, IBlurPoolFilterer: IBlurPoolFilterer{contract: contract}}, nil
}

// NewIBlurPoolCaller creates a new read-only instance of IBlurPool, bound to a specific deployed contract.
func NewIBlurPoolCaller(address common.Address, caller bind.ContractCaller) (*IBlurPoolCaller, error) {
	contract, err := bindIBlurPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBlurPoolCaller{contract: contract}, nil
}

// NewIBlurPoolTransactor creates a new write-only instance of IBlurPool, bound to a specific deployed contract.
func NewIBlurPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*IBlurPoolTransactor, error) {
	contract, err := bindIBlurPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBlurPoolTransactor{contract: contract}, nil
}

// NewIBlurPoolFilterer creates a new log filterer instance of IBlurPool, bound to a specific deployed contract.
func NewIBlurPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*IBlurPoolFilterer, error) {
	contract, err := bindIBlurPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBlurPoolFilterer{contract: contract}, nil
}

// bindIBlurPool binds a generic wrapper to an already deployed contract.
func bindIBlurPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBlurPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBlurPool *IBlurPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBlurPool.Contract.IBlurPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBlurPool *IBlurPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBlurPool.Contract.IBlurPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBlurPool *IBlurPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBlurPool.Contract.IBlurPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBlurPool *IBlurPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBlurPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBlurPool *IBlurPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBlurPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBlurPool *IBlurPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBlurPool.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_IBlurPool *IBlurPoolCaller) BalanceOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IBlurPool.contract.Call(opts, &out, "balanceOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_IBlurPool *IBlurPoolSession) BalanceOf(user common.Address) (*big.Int, error) {
	return _IBlurPool.Contract.BalanceOf(&_IBlurPool.CallOpts, user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_IBlurPool *IBlurPoolCallerSession) BalanceOf(user common.Address) (*big.Int, error) {
	return _IBlurPool.Contract.BalanceOf(&_IBlurPool.CallOpts, user)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IBlurPool *IBlurPoolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IBlurPool.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IBlurPool *IBlurPoolSession) TotalSupply() (*big.Int, error) {
	return _IBlurPool.Contract.TotalSupply(&_IBlurPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IBlurPool *IBlurPoolCallerSession) TotalSupply() (*big.Int, error) {
	return _IBlurPool.Contract.TotalSupply(&_IBlurPool.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IBlurPool *IBlurPoolTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBlurPool.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IBlurPool *IBlurPoolSession) Deposit() (*types.Transaction, error) {
	return _IBlurPool.Contract.Deposit(&_IBlurPool.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IBlurPool *IBlurPoolTransactorSession) Deposit() (*types.Transaction, error) {
	return _IBlurPool.Contract.Deposit(&_IBlurPool.TransactOpts)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IBlurPool *IBlurPoolTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBlurPool.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IBlurPool *IBlurPoolSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBlurPool.Contract.TransferFrom(&_IBlurPool.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IBlurPool *IBlurPoolTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBlurPool.Contract.TransferFrom(&_IBlurPool.TransactOpts, from, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IBlurPool *IBlurPoolTransactor) Withdraw(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _IBlurPool.contract.Transact(opts, "withdraw", arg0)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IBlurPool *IBlurPoolSession) Withdraw(arg0 *big.Int) (*types.Transaction, error) {
	return _IBlurPool.Contract.Withdraw(&_IBlurPool.TransactOpts, arg0)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IBlurPool *IBlurPoolTransactorSession) Withdraw(arg0 *big.Int) (*types.Transaction, error) {
	return _IBlurPool.Contract.Withdraw(&_IBlurPool.TransactOpts, arg0)
}

// IBlurPoolTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IBlurPool contract.
type IBlurPoolTransferIterator struct {
	Event *IBlurPoolTransfer // Event containing the contract specifics and raw log

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
func (it *IBlurPoolTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBlurPoolTransfer)
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
		it.Event = new(IBlurPoolTransfer)
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
func (it *IBlurPoolTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBlurPoolTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBlurPoolTransfer represents a Transfer event raised by the IBlurPool contract.
type IBlurPoolTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_IBlurPool *IBlurPoolFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IBlurPoolTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IBlurPool.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IBlurPoolTransferIterator{contract: _IBlurPool.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_IBlurPool *IBlurPoolFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IBlurPoolTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IBlurPool.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBlurPoolTransfer)
				if err := _IBlurPool.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_IBlurPool *IBlurPoolFilterer) ParseTransfer(log types.Log) (*IBlurPoolTransfer, error) {
	event := new(IBlurPoolTransfer)
	if err := _IBlurPool.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC1822ProxiableUpgradeableMetaData contains all meta data concerning the IERC1822ProxiableUpgradeable contract.
var IERC1822ProxiableUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IERC1822ProxiableUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC1822ProxiableUpgradeableMetaData.ABI instead.
var IERC1822ProxiableUpgradeableABI = IERC1822ProxiableUpgradeableMetaData.ABI

// IERC1822ProxiableUpgradeable is an auto generated Go binding around an Ethereum contract.
type IERC1822ProxiableUpgradeable struct {
	IERC1822ProxiableUpgradeableCaller     // Read-only binding to the contract
	IERC1822ProxiableUpgradeableTransactor // Write-only binding to the contract
	IERC1822ProxiableUpgradeableFilterer   // Log filterer for contract events
}

// IERC1822ProxiableUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC1822ProxiableUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1822ProxiableUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC1822ProxiableUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1822ProxiableUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC1822ProxiableUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1822ProxiableUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC1822ProxiableUpgradeableSession struct {
	Contract     *IERC1822ProxiableUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IERC1822ProxiableUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC1822ProxiableUpgradeableCallerSession struct {
	Contract *IERC1822ProxiableUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// IERC1822ProxiableUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC1822ProxiableUpgradeableTransactorSession struct {
	Contract     *IERC1822ProxiableUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// IERC1822ProxiableUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC1822ProxiableUpgradeableRaw struct {
	Contract *IERC1822ProxiableUpgradeable // Generic contract binding to access the raw methods on
}

// IERC1822ProxiableUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC1822ProxiableUpgradeableCallerRaw struct {
	Contract *IERC1822ProxiableUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// IERC1822ProxiableUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC1822ProxiableUpgradeableTransactorRaw struct {
	Contract *IERC1822ProxiableUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC1822ProxiableUpgradeable creates a new instance of IERC1822ProxiableUpgradeable, bound to a specific deployed contract.
func NewIERC1822ProxiableUpgradeable(address common.Address, backend bind.ContractBackend) (*IERC1822ProxiableUpgradeable, error) {
	contract, err := bindIERC1822ProxiableUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC1822ProxiableUpgradeable{IERC1822ProxiableUpgradeableCaller: IERC1822ProxiableUpgradeableCaller{contract: contract}, IERC1822ProxiableUpgradeableTransactor: IERC1822ProxiableUpgradeableTransactor{contract: contract}, IERC1822ProxiableUpgradeableFilterer: IERC1822ProxiableUpgradeableFilterer{contract: contract}}, nil
}

// NewIERC1822ProxiableUpgradeableCaller creates a new read-only instance of IERC1822ProxiableUpgradeable, bound to a specific deployed contract.
func NewIERC1822ProxiableUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*IERC1822ProxiableUpgradeableCaller, error) {
	contract, err := bindIERC1822ProxiableUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1822ProxiableUpgradeableCaller{contract: contract}, nil
}

// NewIERC1822ProxiableUpgradeableTransactor creates a new write-only instance of IERC1822ProxiableUpgradeable, bound to a specific deployed contract.
func NewIERC1822ProxiableUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC1822ProxiableUpgradeableTransactor, error) {
	contract, err := bindIERC1822ProxiableUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1822ProxiableUpgradeableTransactor{contract: contract}, nil
}

// NewIERC1822ProxiableUpgradeableFilterer creates a new log filterer instance of IERC1822ProxiableUpgradeable, bound to a specific deployed contract.
func NewIERC1822ProxiableUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC1822ProxiableUpgradeableFilterer, error) {
	contract, err := bindIERC1822ProxiableUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC1822ProxiableUpgradeableFilterer{contract: contract}, nil
}

// bindIERC1822ProxiableUpgradeable binds a generic wrapper to an already deployed contract.
func bindIERC1822ProxiableUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC1822ProxiableUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1822ProxiableUpgradeable *IERC1822ProxiableUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC1822ProxiableUpgradeable.Contract.IERC1822ProxiableUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1822ProxiableUpgradeable *IERC1822ProxiableUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1822ProxiableUpgradeable.Contract.IERC1822ProxiableUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1822ProxiableUpgradeable *IERC1822ProxiableUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1822ProxiableUpgradeable.Contract.IERC1822ProxiableUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1822ProxiableUpgradeable *IERC1822ProxiableUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC1822ProxiableUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1822ProxiableUpgradeable *IERC1822ProxiableUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1822ProxiableUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1822ProxiableUpgradeable *IERC1822ProxiableUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1822ProxiableUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_IERC1822ProxiableUpgradeable *IERC1822ProxiableUpgradeableCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IERC1822ProxiableUpgradeable.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_IERC1822ProxiableUpgradeable *IERC1822ProxiableUpgradeableSession) ProxiableUUID() ([32]byte, error) {
	return _IERC1822ProxiableUpgradeable.Contract.ProxiableUUID(&_IERC1822ProxiableUpgradeable.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_IERC1822ProxiableUpgradeable *IERC1822ProxiableUpgradeableCallerSession) ProxiableUUID() ([32]byte, error) {
	return _IERC1822ProxiableUpgradeable.Contract.ProxiableUUID(&_IERC1822ProxiableUpgradeable.CallOpts)
}

// IExecutionDelegateMetaData contains all meta data concerning the IExecutionDelegate contract.
var IExecutionDelegateMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"approveContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"denyContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"grantApproval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokeApproval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferERC721Unsafe\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IExecutionDelegateABI is the input ABI used to generate the binding from.
// Deprecated: Use IExecutionDelegateMetaData.ABI instead.
var IExecutionDelegateABI = IExecutionDelegateMetaData.ABI

// IExecutionDelegate is an auto generated Go binding around an Ethereum contract.
type IExecutionDelegate struct {
	IExecutionDelegateCaller     // Read-only binding to the contract
	IExecutionDelegateTransactor // Write-only binding to the contract
	IExecutionDelegateFilterer   // Log filterer for contract events
}

// IExecutionDelegateCaller is an auto generated read-only Go binding around an Ethereum contract.
type IExecutionDelegateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExecutionDelegateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IExecutionDelegateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExecutionDelegateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IExecutionDelegateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExecutionDelegateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IExecutionDelegateSession struct {
	Contract     *IExecutionDelegate // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IExecutionDelegateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IExecutionDelegateCallerSession struct {
	Contract *IExecutionDelegateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IExecutionDelegateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IExecutionDelegateTransactorSession struct {
	Contract     *IExecutionDelegateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IExecutionDelegateRaw is an auto generated low-level Go binding around an Ethereum contract.
type IExecutionDelegateRaw struct {
	Contract *IExecutionDelegate // Generic contract binding to access the raw methods on
}

// IExecutionDelegateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IExecutionDelegateCallerRaw struct {
	Contract *IExecutionDelegateCaller // Generic read-only contract binding to access the raw methods on
}

// IExecutionDelegateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IExecutionDelegateTransactorRaw struct {
	Contract *IExecutionDelegateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIExecutionDelegate creates a new instance of IExecutionDelegate, bound to a specific deployed contract.
func NewIExecutionDelegate(address common.Address, backend bind.ContractBackend) (*IExecutionDelegate, error) {
	contract, err := bindIExecutionDelegate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IExecutionDelegate{IExecutionDelegateCaller: IExecutionDelegateCaller{contract: contract}, IExecutionDelegateTransactor: IExecutionDelegateTransactor{contract: contract}, IExecutionDelegateFilterer: IExecutionDelegateFilterer{contract: contract}}, nil
}

// NewIExecutionDelegateCaller creates a new read-only instance of IExecutionDelegate, bound to a specific deployed contract.
func NewIExecutionDelegateCaller(address common.Address, caller bind.ContractCaller) (*IExecutionDelegateCaller, error) {
	contract, err := bindIExecutionDelegate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IExecutionDelegateCaller{contract: contract}, nil
}

// NewIExecutionDelegateTransactor creates a new write-only instance of IExecutionDelegate, bound to a specific deployed contract.
func NewIExecutionDelegateTransactor(address common.Address, transactor bind.ContractTransactor) (*IExecutionDelegateTransactor, error) {
	contract, err := bindIExecutionDelegate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IExecutionDelegateTransactor{contract: contract}, nil
}

// NewIExecutionDelegateFilterer creates a new log filterer instance of IExecutionDelegate, bound to a specific deployed contract.
func NewIExecutionDelegateFilterer(address common.Address, filterer bind.ContractFilterer) (*IExecutionDelegateFilterer, error) {
	contract, err := bindIExecutionDelegate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IExecutionDelegateFilterer{contract: contract}, nil
}

// bindIExecutionDelegate binds a generic wrapper to an already deployed contract.
func bindIExecutionDelegate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IExecutionDelegateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IExecutionDelegate *IExecutionDelegateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IExecutionDelegate.Contract.IExecutionDelegateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IExecutionDelegate *IExecutionDelegateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExecutionDelegate.Contract.IExecutionDelegateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IExecutionDelegate *IExecutionDelegateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IExecutionDelegate.Contract.IExecutionDelegateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IExecutionDelegate *IExecutionDelegateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IExecutionDelegate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IExecutionDelegate *IExecutionDelegateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExecutionDelegate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IExecutionDelegate *IExecutionDelegateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IExecutionDelegate.Contract.contract.Transact(opts, method, params...)
}

// ApproveContract is a paid mutator transaction binding the contract method 0x07f7aafb.
//
// Solidity: function approveContract(address _contract) returns()
func (_IExecutionDelegate *IExecutionDelegateTransactor) ApproveContract(opts *bind.TransactOpts, _contract common.Address) (*types.Transaction, error) {
	return _IExecutionDelegate.contract.Transact(opts, "approveContract", _contract)
}

// ApproveContract is a paid mutator transaction binding the contract method 0x07f7aafb.
//
// Solidity: function approveContract(address _contract) returns()
func (_IExecutionDelegate *IExecutionDelegateSession) ApproveContract(_contract common.Address) (*types.Transaction, error) {
	return _IExecutionDelegate.Contract.ApproveContract(&_IExecutionDelegate.TransactOpts, _contract)
}

// ApproveContract is a paid mutator transaction binding the contract method 0x07f7aafb.
//
// Solidity: function approveContract(address _contract) returns()
func (_IExecutionDelegate *IExecutionDelegateTransactorSession) ApproveContract(_contract common.Address) (*types.Transaction, error) {
	return _IExecutionDelegate.Contract.ApproveContract(&_IExecutionDelegate.TransactOpts, _contract)
}

// DenyContract is a paid mutator transaction binding the contract method 0xb7e2077e.
//
// Solidity: function denyContract(address _contract) returns()
func (_IExecutionDelegate *IExecutionDelegateTransactor) DenyContract(opts *bind.TransactOpts, _contract common.Address) (*types.Transaction, error) {
	return _IExecutionDelegate.contract.Transact(opts, "denyContract", _contract)
}

// DenyContract is a paid mutator transaction binding the contract method 0xb7e2077e.
//
// Solidity: function denyContract(address _contract) returns()
func (_IExecutionDelegate *IExecutionDelegateSession) DenyContract(_contract common.Address) (*types.Transaction, error) {
	return _IExecutionDelegate.Contract.DenyContract(&_IExecutionDelegate.TransactOpts, _contract)
}

// DenyContract is a paid mutator transaction binding the contract method 0xb7e2077e.
//
// Solidity: function denyContract(address _contract) returns()
func (_IExecutionDelegate *IExecutionDelegateTransactorSession) DenyContract(_contract common.Address) (*types.Transaction, error) {
	return _IExecutionDelegate.Contract.DenyContract(&_IExecutionDelegate.TransactOpts, _contract)
}

// GrantApproval is a paid mutator transaction binding the contract method 0xa8034df1.
//
// Solidity: function grantApproval() returns()
func (_IExecutionDelegate *IExecutionDelegateTransactor) GrantApproval(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExecutionDelegate.contract.Transact(opts, "grantApproval")
}

// GrantApproval is a paid mutator transaction binding the contract method 0xa8034df1.
//
// Solidity: function grantApproval() returns()
func (_IExecutionDelegate *IExecutionDelegateSession) GrantApproval() (*types.Transaction, error) {
	return _IExecutionDelegate.Contract.GrantApproval(&_IExecutionDelegate.TransactOpts)
}

// GrantApproval is a paid mutator transaction binding the contract method 0xa8034df1.
//
// Solidity: function grantApproval() returns()
func (_IExecutionDelegate *IExecutionDelegateTransactorSession) GrantApproval() (*types.Transaction, error) {
	return _IExecutionDelegate.Contract.GrantApproval(&_IExecutionDelegate.TransactOpts)
}

// RevokeApproval is a paid mutator transaction binding the contract method 0x90d02b3c.
//
// Solidity: function revokeApproval() returns()
func (_IExecutionDelegate *IExecutionDelegateTransactor) RevokeApproval(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExecutionDelegate.contract.Transact(opts, "revokeApproval")
}

// RevokeApproval is a paid mutator transaction binding the contract method 0x90d02b3c.
//
// Solidity: function revokeApproval() returns()
func (_IExecutionDelegate *IExecutionDelegateSession) RevokeApproval() (*types.Transaction, error) {
	return _IExecutionDelegate.Contract.RevokeApproval(&_IExecutionDelegate.TransactOpts)
}

// RevokeApproval is a paid mutator transaction binding the contract method 0x90d02b3c.
//
// Solidity: function revokeApproval() returns()
func (_IExecutionDelegate *IExecutionDelegateTransactorSession) RevokeApproval() (*types.Transaction, error) {
	return _IExecutionDelegate.Contract.RevokeApproval(&_IExecutionDelegate.TransactOpts)
}

// TransferERC1155 is a paid mutator transaction binding the contract method 0x74a9402e.
//
// Solidity: function transferERC1155(address collection, address from, address to, uint256 tokenId, uint256 amount) returns()
func (_IExecutionDelegate *IExecutionDelegateTransactor) TransferERC1155(opts *bind.TransactOpts, collection common.Address, from common.Address, to common.Address, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _IExecutionDelegate.contract.Transact(opts, "transferERC1155", collection, from, to, tokenId, amount)
}

// TransferERC1155 is a paid mutator transaction binding the contract method 0x74a9402e.
//
// Solidity: function transferERC1155(address collection, address from, address to, uint256 tokenId, uint256 amount) returns()
func (_IExecutionDelegate *IExecutionDelegateSession) TransferERC1155(collection common.Address, from common.Address, to common.Address, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _IExecutionDelegate.Contract.TransferERC1155(&_IExecutionDelegate.TransactOpts, collection, from, to, tokenId, amount)
}

// TransferERC1155 is a paid mutator transaction binding the contract method 0x74a9402e.
//
// Solidity: function transferERC1155(address collection, address from, address to, uint256 tokenId, uint256 amount) returns()
func (_IExecutionDelegate *IExecutionDelegateTransactorSession) TransferERC1155(collection common.Address, from common.Address, to common.Address, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _IExecutionDelegate.Contract.TransferERC1155(&_IExecutionDelegate.TransactOpts, collection, from, to, tokenId, amount)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0xda3e8ce4.
//
// Solidity: function transferERC20(address token, address from, address to, uint256 amount) returns()
func (_IExecutionDelegate *IExecutionDelegateTransactor) TransferERC20(opts *bind.TransactOpts, token common.Address, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IExecutionDelegate.contract.Transact(opts, "transferERC20", token, from, to, amount)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0xda3e8ce4.
//
// Solidity: function transferERC20(address token, address from, address to, uint256 amount) returns()
func (_IExecutionDelegate *IExecutionDelegateSession) TransferERC20(token common.Address, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IExecutionDelegate.Contract.TransferERC20(&_IExecutionDelegate.TransactOpts, token, from, to, amount)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0xda3e8ce4.
//
// Solidity: function transferERC20(address token, address from, address to, uint256 amount) returns()
func (_IExecutionDelegate *IExecutionDelegateTransactorSession) TransferERC20(token common.Address, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IExecutionDelegate.Contract.TransferERC20(&_IExecutionDelegate.TransactOpts, token, from, to, amount)
}

// TransferERC721 is a paid mutator transaction binding the contract method 0x789f93f6.
//
// Solidity: function transferERC721(address collection, address from, address to, uint256 tokenId) returns()
func (_IExecutionDelegate *IExecutionDelegateTransactor) TransferERC721(opts *bind.TransactOpts, collection common.Address, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IExecutionDelegate.contract.Transact(opts, "transferERC721", collection, from, to, tokenId)
}

// TransferERC721 is a paid mutator transaction binding the contract method 0x789f93f6.
//
// Solidity: function transferERC721(address collection, address from, address to, uint256 tokenId) returns()
func (_IExecutionDelegate *IExecutionDelegateSession) TransferERC721(collection common.Address, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IExecutionDelegate.Contract.TransferERC721(&_IExecutionDelegate.TransactOpts, collection, from, to, tokenId)
}

// TransferERC721 is a paid mutator transaction binding the contract method 0x789f93f6.
//
// Solidity: function transferERC721(address collection, address from, address to, uint256 tokenId) returns()
func (_IExecutionDelegate *IExecutionDelegateTransactorSession) TransferERC721(collection common.Address, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IExecutionDelegate.Contract.TransferERC721(&_IExecutionDelegate.TransactOpts, collection, from, to, tokenId)
}

// TransferERC721Unsafe is a paid mutator transaction binding the contract method 0xca72da8e.
//
// Solidity: function transferERC721Unsafe(address collection, address from, address to, uint256 tokenId) returns()
func (_IExecutionDelegate *IExecutionDelegateTransactor) TransferERC721Unsafe(opts *bind.TransactOpts, collection common.Address, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IExecutionDelegate.contract.Transact(opts, "transferERC721Unsafe", collection, from, to, tokenId)
}

// TransferERC721Unsafe is a paid mutator transaction binding the contract method 0xca72da8e.
//
// Solidity: function transferERC721Unsafe(address collection, address from, address to, uint256 tokenId) returns()
func (_IExecutionDelegate *IExecutionDelegateSession) TransferERC721Unsafe(collection common.Address, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IExecutionDelegate.Contract.TransferERC721Unsafe(&_IExecutionDelegate.TransactOpts, collection, from, to, tokenId)
}

// TransferERC721Unsafe is a paid mutator transaction binding the contract method 0xca72da8e.
//
// Solidity: function transferERC721Unsafe(address collection, address from, address to, uint256 tokenId) returns()
func (_IExecutionDelegate *IExecutionDelegateTransactorSession) TransferERC721Unsafe(collection common.Address, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IExecutionDelegate.Contract.TransferERC721Unsafe(&_IExecutionDelegate.TransactOpts, collection, from, to, tokenId)
}

// IMatchingPolicyMetaData contains all meta data concerning the IMatchingPolicy contract.
var IMatchingPolicyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"makerAsk\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"takerBid\",\"type\":\"tuple\"}],\"name\":\"canMatchMakerAsk\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"enumAssetType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"makerBid\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"takerAsk\",\"type\":\"tuple\"}],\"name\":\"canMatchMakerBid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"enumAssetType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IMatchingPolicyABI is the input ABI used to generate the binding from.
// Deprecated: Use IMatchingPolicyMetaData.ABI instead.
var IMatchingPolicyABI = IMatchingPolicyMetaData.ABI

// IMatchingPolicy is an auto generated Go binding around an Ethereum contract.
type IMatchingPolicy struct {
	IMatchingPolicyCaller     // Read-only binding to the contract
	IMatchingPolicyTransactor // Write-only binding to the contract
	IMatchingPolicyFilterer   // Log filterer for contract events
}

// IMatchingPolicyCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMatchingPolicyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMatchingPolicyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMatchingPolicyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMatchingPolicyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMatchingPolicyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMatchingPolicySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMatchingPolicySession struct {
	Contract     *IMatchingPolicy  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IMatchingPolicyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMatchingPolicyCallerSession struct {
	Contract *IMatchingPolicyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IMatchingPolicyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMatchingPolicyTransactorSession struct {
	Contract     *IMatchingPolicyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IMatchingPolicyRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMatchingPolicyRaw struct {
	Contract *IMatchingPolicy // Generic contract binding to access the raw methods on
}

// IMatchingPolicyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMatchingPolicyCallerRaw struct {
	Contract *IMatchingPolicyCaller // Generic read-only contract binding to access the raw methods on
}

// IMatchingPolicyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMatchingPolicyTransactorRaw struct {
	Contract *IMatchingPolicyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMatchingPolicy creates a new instance of IMatchingPolicy, bound to a specific deployed contract.
func NewIMatchingPolicy(address common.Address, backend bind.ContractBackend) (*IMatchingPolicy, error) {
	contract, err := bindIMatchingPolicy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMatchingPolicy{IMatchingPolicyCaller: IMatchingPolicyCaller{contract: contract}, IMatchingPolicyTransactor: IMatchingPolicyTransactor{contract: contract}, IMatchingPolicyFilterer: IMatchingPolicyFilterer{contract: contract}}, nil
}

// NewIMatchingPolicyCaller creates a new read-only instance of IMatchingPolicy, bound to a specific deployed contract.
func NewIMatchingPolicyCaller(address common.Address, caller bind.ContractCaller) (*IMatchingPolicyCaller, error) {
	contract, err := bindIMatchingPolicy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMatchingPolicyCaller{contract: contract}, nil
}

// NewIMatchingPolicyTransactor creates a new write-only instance of IMatchingPolicy, bound to a specific deployed contract.
func NewIMatchingPolicyTransactor(address common.Address, transactor bind.ContractTransactor) (*IMatchingPolicyTransactor, error) {
	contract, err := bindIMatchingPolicy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMatchingPolicyTransactor{contract: contract}, nil
}

// NewIMatchingPolicyFilterer creates a new log filterer instance of IMatchingPolicy, bound to a specific deployed contract.
func NewIMatchingPolicyFilterer(address common.Address, filterer bind.ContractFilterer) (*IMatchingPolicyFilterer, error) {
	contract, err := bindIMatchingPolicy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMatchingPolicyFilterer{contract: contract}, nil
}

// bindIMatchingPolicy binds a generic wrapper to an already deployed contract.
func bindIMatchingPolicy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IMatchingPolicyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMatchingPolicy *IMatchingPolicyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMatchingPolicy.Contract.IMatchingPolicyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMatchingPolicy *IMatchingPolicyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMatchingPolicy.Contract.IMatchingPolicyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMatchingPolicy *IMatchingPolicyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMatchingPolicy.Contract.IMatchingPolicyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMatchingPolicy *IMatchingPolicyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMatchingPolicy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMatchingPolicy *IMatchingPolicyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMatchingPolicy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMatchingPolicy *IMatchingPolicyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMatchingPolicy.Contract.contract.Transact(opts, method, params...)
}

// CanMatchMakerAsk is a free data retrieval call binding the contract method 0xd5ec8c77.
//
// Solidity: function canMatchMakerAsk((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) makerAsk, (address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) takerBid) view returns(bool, uint256, uint256, uint256, uint8)
func (_IMatchingPolicy *IMatchingPolicyCaller) CanMatchMakerAsk(opts *bind.CallOpts, makerAsk Order, takerBid Order) (bool, *big.Int, *big.Int, *big.Int, uint8, error) {
	var out []interface{}
	err := _IMatchingPolicy.contract.Call(opts, &out, "canMatchMakerAsk", makerAsk, takerBid)

	if err != nil {
		return *new(bool), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(uint8)).(*uint8)

	return out0, out1, out2, out3, out4, err

}

// CanMatchMakerAsk is a free data retrieval call binding the contract method 0xd5ec8c77.
//
// Solidity: function canMatchMakerAsk((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) makerAsk, (address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) takerBid) view returns(bool, uint256, uint256, uint256, uint8)
func (_IMatchingPolicy *IMatchingPolicySession) CanMatchMakerAsk(makerAsk Order, takerBid Order) (bool, *big.Int, *big.Int, *big.Int, uint8, error) {
	return _IMatchingPolicy.Contract.CanMatchMakerAsk(&_IMatchingPolicy.CallOpts, makerAsk, takerBid)
}

// CanMatchMakerAsk is a free data retrieval call binding the contract method 0xd5ec8c77.
//
// Solidity: function canMatchMakerAsk((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) makerAsk, (address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) takerBid) view returns(bool, uint256, uint256, uint256, uint8)
func (_IMatchingPolicy *IMatchingPolicyCallerSession) CanMatchMakerAsk(makerAsk Order, takerBid Order) (bool, *big.Int, *big.Int, *big.Int, uint8, error) {
	return _IMatchingPolicy.Contract.CanMatchMakerAsk(&_IMatchingPolicy.CallOpts, makerAsk, takerBid)
}

// CanMatchMakerBid is a free data retrieval call binding the contract method 0x0813a766.
//
// Solidity: function canMatchMakerBid((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) makerBid, (address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) takerAsk) view returns(bool, uint256, uint256, uint256, uint8)
func (_IMatchingPolicy *IMatchingPolicyCaller) CanMatchMakerBid(opts *bind.CallOpts, makerBid Order, takerAsk Order) (bool, *big.Int, *big.Int, *big.Int, uint8, error) {
	var out []interface{}
	err := _IMatchingPolicy.contract.Call(opts, &out, "canMatchMakerBid", makerBid, takerAsk)

	if err != nil {
		return *new(bool), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(uint8)).(*uint8)

	return out0, out1, out2, out3, out4, err

}

// CanMatchMakerBid is a free data retrieval call binding the contract method 0x0813a766.
//
// Solidity: function canMatchMakerBid((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) makerBid, (address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) takerAsk) view returns(bool, uint256, uint256, uint256, uint8)
func (_IMatchingPolicy *IMatchingPolicySession) CanMatchMakerBid(makerBid Order, takerAsk Order) (bool, *big.Int, *big.Int, *big.Int, uint8, error) {
	return _IMatchingPolicy.Contract.CanMatchMakerBid(&_IMatchingPolicy.CallOpts, makerBid, takerAsk)
}

// CanMatchMakerBid is a free data retrieval call binding the contract method 0x0813a766.
//
// Solidity: function canMatchMakerBid((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) makerBid, (address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) takerAsk) view returns(bool, uint256, uint256, uint256, uint8)
func (_IMatchingPolicy *IMatchingPolicyCallerSession) CanMatchMakerBid(makerBid Order, takerAsk Order) (bool, *big.Int, *big.Int, *big.Int, uint8, error) {
	return _IMatchingPolicy.Contract.CanMatchMakerBid(&_IMatchingPolicy.CallOpts, makerBid, takerAsk)
}

// IPolicyManagerMetaData contains all meta data concerning the IPolicyManager contract.
var IPolicyManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"policy\",\"type\":\"address\"}],\"name\":\"addPolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"policy\",\"type\":\"address\"}],\"name\":\"isPolicyWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"policy\",\"type\":\"address\"}],\"name\":\"removePolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"viewCountWhitelistedPolicies\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cursor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"viewWhitelistedPolicies\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IPolicyManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IPolicyManagerMetaData.ABI instead.
var IPolicyManagerABI = IPolicyManagerMetaData.ABI

// IPolicyManager is an auto generated Go binding around an Ethereum contract.
type IPolicyManager struct {
	IPolicyManagerCaller     // Read-only binding to the contract
	IPolicyManagerTransactor // Write-only binding to the contract
	IPolicyManagerFilterer   // Log filterer for contract events
}

// IPolicyManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPolicyManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPolicyManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPolicyManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPolicyManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPolicyManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPolicyManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPolicyManagerSession struct {
	Contract     *IPolicyManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPolicyManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPolicyManagerCallerSession struct {
	Contract *IPolicyManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IPolicyManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPolicyManagerTransactorSession struct {
	Contract     *IPolicyManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IPolicyManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPolicyManagerRaw struct {
	Contract *IPolicyManager // Generic contract binding to access the raw methods on
}

// IPolicyManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPolicyManagerCallerRaw struct {
	Contract *IPolicyManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IPolicyManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPolicyManagerTransactorRaw struct {
	Contract *IPolicyManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPolicyManager creates a new instance of IPolicyManager, bound to a specific deployed contract.
func NewIPolicyManager(address common.Address, backend bind.ContractBackend) (*IPolicyManager, error) {
	contract, err := bindIPolicyManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPolicyManager{IPolicyManagerCaller: IPolicyManagerCaller{contract: contract}, IPolicyManagerTransactor: IPolicyManagerTransactor{contract: contract}, IPolicyManagerFilterer: IPolicyManagerFilterer{contract: contract}}, nil
}

// NewIPolicyManagerCaller creates a new read-only instance of IPolicyManager, bound to a specific deployed contract.
func NewIPolicyManagerCaller(address common.Address, caller bind.ContractCaller) (*IPolicyManagerCaller, error) {
	contract, err := bindIPolicyManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPolicyManagerCaller{contract: contract}, nil
}

// NewIPolicyManagerTransactor creates a new write-only instance of IPolicyManager, bound to a specific deployed contract.
func NewIPolicyManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IPolicyManagerTransactor, error) {
	contract, err := bindIPolicyManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPolicyManagerTransactor{contract: contract}, nil
}

// NewIPolicyManagerFilterer creates a new log filterer instance of IPolicyManager, bound to a specific deployed contract.
func NewIPolicyManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IPolicyManagerFilterer, error) {
	contract, err := bindIPolicyManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPolicyManagerFilterer{contract: contract}, nil
}

// bindIPolicyManager binds a generic wrapper to an already deployed contract.
func bindIPolicyManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPolicyManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPolicyManager *IPolicyManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPolicyManager.Contract.IPolicyManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPolicyManager *IPolicyManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPolicyManager.Contract.IPolicyManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPolicyManager *IPolicyManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPolicyManager.Contract.IPolicyManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPolicyManager *IPolicyManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPolicyManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPolicyManager *IPolicyManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPolicyManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPolicyManager *IPolicyManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPolicyManager.Contract.contract.Transact(opts, method, params...)
}

// IsPolicyWhitelisted is a free data retrieval call binding the contract method 0x874516cd.
//
// Solidity: function isPolicyWhitelisted(address policy) view returns(bool)
func (_IPolicyManager *IPolicyManagerCaller) IsPolicyWhitelisted(opts *bind.CallOpts, policy common.Address) (bool, error) {
	var out []interface{}
	err := _IPolicyManager.contract.Call(opts, &out, "isPolicyWhitelisted", policy)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPolicyWhitelisted is a free data retrieval call binding the contract method 0x874516cd.
//
// Solidity: function isPolicyWhitelisted(address policy) view returns(bool)
func (_IPolicyManager *IPolicyManagerSession) IsPolicyWhitelisted(policy common.Address) (bool, error) {
	return _IPolicyManager.Contract.IsPolicyWhitelisted(&_IPolicyManager.CallOpts, policy)
}

// IsPolicyWhitelisted is a free data retrieval call binding the contract method 0x874516cd.
//
// Solidity: function isPolicyWhitelisted(address policy) view returns(bool)
func (_IPolicyManager *IPolicyManagerCallerSession) IsPolicyWhitelisted(policy common.Address) (bool, error) {
	return _IPolicyManager.Contract.IsPolicyWhitelisted(&_IPolicyManager.CallOpts, policy)
}

// ViewCountWhitelistedPolicies is a free data retrieval call binding the contract method 0xd6f0d88e.
//
// Solidity: function viewCountWhitelistedPolicies() view returns(uint256)
func (_IPolicyManager *IPolicyManagerCaller) ViewCountWhitelistedPolicies(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPolicyManager.contract.Call(opts, &out, "viewCountWhitelistedPolicies")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ViewCountWhitelistedPolicies is a free data retrieval call binding the contract method 0xd6f0d88e.
//
// Solidity: function viewCountWhitelistedPolicies() view returns(uint256)
func (_IPolicyManager *IPolicyManagerSession) ViewCountWhitelistedPolicies() (*big.Int, error) {
	return _IPolicyManager.Contract.ViewCountWhitelistedPolicies(&_IPolicyManager.CallOpts)
}

// ViewCountWhitelistedPolicies is a free data retrieval call binding the contract method 0xd6f0d88e.
//
// Solidity: function viewCountWhitelistedPolicies() view returns(uint256)
func (_IPolicyManager *IPolicyManagerCallerSession) ViewCountWhitelistedPolicies() (*big.Int, error) {
	return _IPolicyManager.Contract.ViewCountWhitelistedPolicies(&_IPolicyManager.CallOpts)
}

// ViewWhitelistedPolicies is a free data retrieval call binding the contract method 0x747f54ae.
//
// Solidity: function viewWhitelistedPolicies(uint256 cursor, uint256 size) view returns(address[], uint256)
func (_IPolicyManager *IPolicyManagerCaller) ViewWhitelistedPolicies(opts *bind.CallOpts, cursor *big.Int, size *big.Int) ([]common.Address, *big.Int, error) {
	var out []interface{}
	err := _IPolicyManager.contract.Call(opts, &out, "viewWhitelistedPolicies", cursor, size)

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
func (_IPolicyManager *IPolicyManagerSession) ViewWhitelistedPolicies(cursor *big.Int, size *big.Int) ([]common.Address, *big.Int, error) {
	return _IPolicyManager.Contract.ViewWhitelistedPolicies(&_IPolicyManager.CallOpts, cursor, size)
}

// ViewWhitelistedPolicies is a free data retrieval call binding the contract method 0x747f54ae.
//
// Solidity: function viewWhitelistedPolicies(uint256 cursor, uint256 size) view returns(address[], uint256)
func (_IPolicyManager *IPolicyManagerCallerSession) ViewWhitelistedPolicies(cursor *big.Int, size *big.Int) ([]common.Address, *big.Int, error) {
	return _IPolicyManager.Contract.ViewWhitelistedPolicies(&_IPolicyManager.CallOpts, cursor, size)
}

// AddPolicy is a paid mutator transaction binding the contract method 0xb84ef081.
//
// Solidity: function addPolicy(address policy) returns()
func (_IPolicyManager *IPolicyManagerTransactor) AddPolicy(opts *bind.TransactOpts, policy common.Address) (*types.Transaction, error) {
	return _IPolicyManager.contract.Transact(opts, "addPolicy", policy)
}

// AddPolicy is a paid mutator transaction binding the contract method 0xb84ef081.
//
// Solidity: function addPolicy(address policy) returns()
func (_IPolicyManager *IPolicyManagerSession) AddPolicy(policy common.Address) (*types.Transaction, error) {
	return _IPolicyManager.Contract.AddPolicy(&_IPolicyManager.TransactOpts, policy)
}

// AddPolicy is a paid mutator transaction binding the contract method 0xb84ef081.
//
// Solidity: function addPolicy(address policy) returns()
func (_IPolicyManager *IPolicyManagerTransactorSession) AddPolicy(policy common.Address) (*types.Transaction, error) {
	return _IPolicyManager.Contract.AddPolicy(&_IPolicyManager.TransactOpts, policy)
}

// RemovePolicy is a paid mutator transaction binding the contract method 0x78d38d82.
//
// Solidity: function removePolicy(address policy) returns()
func (_IPolicyManager *IPolicyManagerTransactor) RemovePolicy(opts *bind.TransactOpts, policy common.Address) (*types.Transaction, error) {
	return _IPolicyManager.contract.Transact(opts, "removePolicy", policy)
}

// RemovePolicy is a paid mutator transaction binding the contract method 0x78d38d82.
//
// Solidity: function removePolicy(address policy) returns()
func (_IPolicyManager *IPolicyManagerSession) RemovePolicy(policy common.Address) (*types.Transaction, error) {
	return _IPolicyManager.Contract.RemovePolicy(&_IPolicyManager.TransactOpts, policy)
}

// RemovePolicy is a paid mutator transaction binding the contract method 0x78d38d82.
//
// Solidity: function removePolicy(address policy) returns()
func (_IPolicyManager *IPolicyManagerTransactorSession) RemovePolicy(policy common.Address) (*types.Transaction, error) {
	return _IPolicyManager.Contract.RemovePolicy(&_IPolicyManager.TransactOpts, policy)
}

// InitializableMetaData contains all meta data concerning the Initializable contract.
var InitializableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"}]",
}

// InitializableABI is the input ABI used to generate the binding from.
// Deprecated: Use InitializableMetaData.ABI instead.
var InitializableABI = InitializableMetaData.ABI

// Initializable is an auto generated Go binding around an Ethereum contract.
type Initializable struct {
	InitializableCaller     // Read-only binding to the contract
	InitializableTransactor // Write-only binding to the contract
	InitializableFilterer   // Log filterer for contract events
}

// InitializableCaller is an auto generated read-only Go binding around an Ethereum contract.
type InitializableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InitializableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InitializableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InitializableSession struct {
	Contract     *Initializable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InitializableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InitializableCallerSession struct {
	Contract *InitializableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// InitializableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InitializableTransactorSession struct {
	Contract     *InitializableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// InitializableRaw is an auto generated low-level Go binding around an Ethereum contract.
type InitializableRaw struct {
	Contract *Initializable // Generic contract binding to access the raw methods on
}

// InitializableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InitializableCallerRaw struct {
	Contract *InitializableCaller // Generic read-only contract binding to access the raw methods on
}

// InitializableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InitializableTransactorRaw struct {
	Contract *InitializableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInitializable creates a new instance of Initializable, bound to a specific deployed contract.
func NewInitializable(address common.Address, backend bind.ContractBackend) (*Initializable, error) {
	contract, err := bindInitializable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Initializable{InitializableCaller: InitializableCaller{contract: contract}, InitializableTransactor: InitializableTransactor{contract: contract}, InitializableFilterer: InitializableFilterer{contract: contract}}, nil
}

// NewInitializableCaller creates a new read-only instance of Initializable, bound to a specific deployed contract.
func NewInitializableCaller(address common.Address, caller bind.ContractCaller) (*InitializableCaller, error) {
	contract, err := bindInitializable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableCaller{contract: contract}, nil
}

// NewInitializableTransactor creates a new write-only instance of Initializable, bound to a specific deployed contract.
func NewInitializableTransactor(address common.Address, transactor bind.ContractTransactor) (*InitializableTransactor, error) {
	contract, err := bindInitializable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableTransactor{contract: contract}, nil
}

// NewInitializableFilterer creates a new log filterer instance of Initializable, bound to a specific deployed contract.
func NewInitializableFilterer(address common.Address, filterer bind.ContractFilterer) (*InitializableFilterer, error) {
	contract, err := bindInitializable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InitializableFilterer{contract: contract}, nil
}

// bindInitializable binds a generic wrapper to an already deployed contract.
func bindInitializable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InitializableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Initializable *InitializableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Initializable.Contract.InitializableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Initializable *InitializableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initializable.Contract.InitializableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Initializable *InitializableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Initializable.Contract.InitializableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Initializable *InitializableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Initializable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Initializable *InitializableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initializable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Initializable *InitializableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Initializable.Contract.contract.Transact(opts, method, params...)
}

// InitializableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Initializable contract.
type InitializableInitializedIterator struct {
	Event *InitializableInitialized // Event containing the contract specifics and raw log

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
func (it *InitializableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InitializableInitialized)
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
		it.Event = new(InitializableInitialized)
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
func (it *InitializableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InitializableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InitializableInitialized represents a Initialized event raised by the Initializable contract.
type InitializableInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Initializable *InitializableFilterer) FilterInitialized(opts *bind.FilterOpts) (*InitializableInitializedIterator, error) {

	logs, sub, err := _Initializable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &InitializableInitializedIterator{contract: _Initializable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Initializable *InitializableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *InitializableInitialized) (event.Subscription, error) {

	logs, sub, err := _Initializable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InitializableInitialized)
				if err := _Initializable.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Initializable *InitializableFilterer) ParseInitialized(log types.Log) (*InitializableInitialized, error) {
	event := new(InitializableInitialized)
	if err := _Initializable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerkleVerifierMetaData contains all meta data concerning the MerkleVerifier contract.
var MerkleVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"_computeRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"_verifyProof\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x6102f461003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100405760003560e01c806372aaadac146100455780639c7bf9381461005a575b600080fd5b6100586100533660046101ea565b61007f565b005b61006d61006836600461023a565b6100b3565b60405190815260200160405180910390f35b600061008b84836100b3565b90508281146100ad576040516309bde33960e01b815260040160405180910390fd5b50505050565b600082815b83518110156100ff5760008482815181106100d5576100d5610281565b602002602001015190506100e98382610107565b92505080806100f790610297565b9150506100b8565b509392505050565b6000818310610123576000828152602084905260409020610132565b60008381526020839052604090205b9392505050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261016057600080fd5b8135602067ffffffffffffffff8083111561017d5761017d610139565b8260051b604051601f19603f830116810181811084821117156101a2576101a2610139565b6040529384528581018301938381019250878511156101c057600080fd5b83870191505b848210156101df578135835291830191908301906101c6565b979650505050505050565b6000806000606084860312156101ff57600080fd5b8335925060208401359150604084013567ffffffffffffffff81111561022457600080fd5b6102308682870161014f565b9150509250925092565b6000806040838503121561024d57600080fd5b82359150602083013567ffffffffffffffff81111561026b57600080fd5b6102778582860161014f565b9150509250929050565b634e487b7160e01b600052603260045260246000fd5b6000600182016102b757634e487b7160e01b600052601160045260246000fd5b506001019056fea264697066735822122095b583675eb46a37009578a0f3bb6b1d018c53b09d2ea49e2d14cfe9a17e135764736f6c63430008110033",
}

// MerkleVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use MerkleVerifierMetaData.ABI instead.
var MerkleVerifierABI = MerkleVerifierMetaData.ABI

// MerkleVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MerkleVerifierMetaData.Bin instead.
var MerkleVerifierBin = MerkleVerifierMetaData.Bin

// DeployMerkleVerifier deploys a new Ethereum contract, binding an instance of MerkleVerifier to it.
func DeployMerkleVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MerkleVerifier, error) {
	parsed, err := MerkleVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MerkleVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MerkleVerifier{MerkleVerifierCaller: MerkleVerifierCaller{contract: contract}, MerkleVerifierTransactor: MerkleVerifierTransactor{contract: contract}, MerkleVerifierFilterer: MerkleVerifierFilterer{contract: contract}}, nil
}

// MerkleVerifier is an auto generated Go binding around an Ethereum contract.
type MerkleVerifier struct {
	MerkleVerifierCaller     // Read-only binding to the contract
	MerkleVerifierTransactor // Write-only binding to the contract
	MerkleVerifierFilterer   // Log filterer for contract events
}

// MerkleVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleVerifierSession struct {
	Contract     *MerkleVerifier   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleVerifierCallerSession struct {
	Contract *MerkleVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// MerkleVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleVerifierTransactorSession struct {
	Contract     *MerkleVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MerkleVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleVerifierRaw struct {
	Contract *MerkleVerifier // Generic contract binding to access the raw methods on
}

// MerkleVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleVerifierCallerRaw struct {
	Contract *MerkleVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleVerifierTransactorRaw struct {
	Contract *MerkleVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkleVerifier creates a new instance of MerkleVerifier, bound to a specific deployed contract.
func NewMerkleVerifier(address common.Address, backend bind.ContractBackend) (*MerkleVerifier, error) {
	contract, err := bindMerkleVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerkleVerifier{MerkleVerifierCaller: MerkleVerifierCaller{contract: contract}, MerkleVerifierTransactor: MerkleVerifierTransactor{contract: contract}, MerkleVerifierFilterer: MerkleVerifierFilterer{contract: contract}}, nil
}

// NewMerkleVerifierCaller creates a new read-only instance of MerkleVerifier, bound to a specific deployed contract.
func NewMerkleVerifierCaller(address common.Address, caller bind.ContractCaller) (*MerkleVerifierCaller, error) {
	contract, err := bindMerkleVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleVerifierCaller{contract: contract}, nil
}

// NewMerkleVerifierTransactor creates a new write-only instance of MerkleVerifier, bound to a specific deployed contract.
func NewMerkleVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleVerifierTransactor, error) {
	contract, err := bindMerkleVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleVerifierTransactor{contract: contract}, nil
}

// NewMerkleVerifierFilterer creates a new log filterer instance of MerkleVerifier, bound to a specific deployed contract.
func NewMerkleVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkleVerifierFilterer, error) {
	contract, err := bindMerkleVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleVerifierFilterer{contract: contract}, nil
}

// bindMerkleVerifier binds a generic wrapper to an already deployed contract.
func bindMerkleVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleVerifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleVerifier *MerkleVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleVerifier.Contract.MerkleVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleVerifier *MerkleVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleVerifier.Contract.MerkleVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleVerifier *MerkleVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleVerifier.Contract.MerkleVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleVerifier *MerkleVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleVerifier *MerkleVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleVerifier *MerkleVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleVerifier.Contract.contract.Transact(opts, method, params...)
}

// ComputeRoot is a free data retrieval call binding the contract method 0x9c7bf938.
//
// Solidity: function _computeRoot(bytes32 leaf, bytes32[] proof) pure returns(bytes32)
func (_MerkleVerifier *MerkleVerifierCaller) ComputeRoot(opts *bind.CallOpts, leaf [32]byte, proof [][32]byte) ([32]byte, error) {
	var out []interface{}
	err := _MerkleVerifier.contract.Call(opts, &out, "_computeRoot", leaf, proof)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ComputeRoot is a free data retrieval call binding the contract method 0x9c7bf938.
//
// Solidity: function _computeRoot(bytes32 leaf, bytes32[] proof) pure returns(bytes32)
func (_MerkleVerifier *MerkleVerifierSession) ComputeRoot(leaf [32]byte, proof [][32]byte) ([32]byte, error) {
	return _MerkleVerifier.Contract.ComputeRoot(&_MerkleVerifier.CallOpts, leaf, proof)
}

// ComputeRoot is a free data retrieval call binding the contract method 0x9c7bf938.
//
// Solidity: function _computeRoot(bytes32 leaf, bytes32[] proof) pure returns(bytes32)
func (_MerkleVerifier *MerkleVerifierCallerSession) ComputeRoot(leaf [32]byte, proof [][32]byte) ([32]byte, error) {
	return _MerkleVerifier.Contract.ComputeRoot(&_MerkleVerifier.CallOpts, leaf, proof)
}

// VerifyProof is a free data retrieval call binding the contract method 0x72aaadac.
//
// Solidity: function _verifyProof(bytes32 leaf, bytes32 root, bytes32[] proof) pure returns()
func (_MerkleVerifier *MerkleVerifierCaller) VerifyProof(opts *bind.CallOpts, leaf [32]byte, root [32]byte, proof [][32]byte) error {
	var out []interface{}
	err := _MerkleVerifier.contract.Call(opts, &out, "_verifyProof", leaf, root, proof)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0x72aaadac.
//
// Solidity: function _verifyProof(bytes32 leaf, bytes32 root, bytes32[] proof) pure returns()
func (_MerkleVerifier *MerkleVerifierSession) VerifyProof(leaf [32]byte, root [32]byte, proof [][32]byte) error {
	return _MerkleVerifier.Contract.VerifyProof(&_MerkleVerifier.CallOpts, leaf, root, proof)
}

// VerifyProof is a free data retrieval call binding the contract method 0x72aaadac.
//
// Solidity: function _verifyProof(bytes32 leaf, bytes32 root, bytes32[] proof) pure returns()
func (_MerkleVerifier *MerkleVerifierCallerSession) VerifyProof(leaf [32]byte, root [32]byte, proof [][32]byte) error {
	return _MerkleVerifier.Contract.VerifyProof(&_MerkleVerifier.CallOpts, leaf, root, proof)
}

// OwnableUpgradeableMetaData contains all meta data concerning the OwnableUpgradeable contract.
var OwnableUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OwnableUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnableUpgradeableMetaData.ABI instead.
var OwnableUpgradeableABI = OwnableUpgradeableMetaData.ABI

// OwnableUpgradeable is an auto generated Go binding around an Ethereum contract.
type OwnableUpgradeable struct {
	OwnableUpgradeableCaller     // Read-only binding to the contract
	OwnableUpgradeableTransactor // Write-only binding to the contract
	OwnableUpgradeableFilterer   // Log filterer for contract events
}

// OwnableUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableUpgradeableSession struct {
	Contract     *OwnableUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// OwnableUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableUpgradeableCallerSession struct {
	Contract *OwnableUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// OwnableUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableUpgradeableTransactorSession struct {
	Contract     *OwnableUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// OwnableUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableUpgradeableRaw struct {
	Contract *OwnableUpgradeable // Generic contract binding to access the raw methods on
}

// OwnableUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableUpgradeableCallerRaw struct {
	Contract *OwnableUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableUpgradeableTransactorRaw struct {
	Contract *OwnableUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnableUpgradeable creates a new instance of OwnableUpgradeable, bound to a specific deployed contract.
func NewOwnableUpgradeable(address common.Address, backend bind.ContractBackend) (*OwnableUpgradeable, error) {
	contract, err := bindOwnableUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeable{OwnableUpgradeableCaller: OwnableUpgradeableCaller{contract: contract}, OwnableUpgradeableTransactor: OwnableUpgradeableTransactor{contract: contract}, OwnableUpgradeableFilterer: OwnableUpgradeableFilterer{contract: contract}}, nil
}

// NewOwnableUpgradeableCaller creates a new read-only instance of OwnableUpgradeable, bound to a specific deployed contract.
func NewOwnableUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*OwnableUpgradeableCaller, error) {
	contract, err := bindOwnableUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableCaller{contract: contract}, nil
}

// NewOwnableUpgradeableTransactor creates a new write-only instance of OwnableUpgradeable, bound to a specific deployed contract.
func NewOwnableUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableUpgradeableTransactor, error) {
	contract, err := bindOwnableUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableTransactor{contract: contract}, nil
}

// NewOwnableUpgradeableFilterer creates a new log filterer instance of OwnableUpgradeable, bound to a specific deployed contract.
func NewOwnableUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableUpgradeableFilterer, error) {
	contract, err := bindOwnableUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableFilterer{contract: contract}, nil
}

// bindOwnableUpgradeable binds a generic wrapper to an already deployed contract.
func bindOwnableUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnableUpgradeable *OwnableUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnableUpgradeable.Contract.OwnableUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnableUpgradeable *OwnableUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.OwnableUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnableUpgradeable *OwnableUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.OwnableUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnableUpgradeable *OwnableUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnableUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnableUpgradeable *OwnableUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnableUpgradeable *OwnableUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnableUpgradeable *OwnableUpgradeableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OwnableUpgradeable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnableUpgradeable *OwnableUpgradeableSession) Owner() (common.Address, error) {
	return _OwnableUpgradeable.Contract.Owner(&_OwnableUpgradeable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnableUpgradeable *OwnableUpgradeableCallerSession) Owner() (common.Address, error) {
	return _OwnableUpgradeable.Contract.Owner(&_OwnableUpgradeable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnableUpgradeable *OwnableUpgradeableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableUpgradeable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnableUpgradeable *OwnableUpgradeableSession) RenounceOwnership() (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.RenounceOwnership(&_OwnableUpgradeable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnableUpgradeable *OwnableUpgradeableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.RenounceOwnership(&_OwnableUpgradeable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnableUpgradeable *OwnableUpgradeableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OwnableUpgradeable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnableUpgradeable *OwnableUpgradeableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.TransferOwnership(&_OwnableUpgradeable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnableUpgradeable *OwnableUpgradeableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.TransferOwnership(&_OwnableUpgradeable.TransactOpts, newOwner)
}

// OwnableUpgradeableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the OwnableUpgradeable contract.
type OwnableUpgradeableInitializedIterator struct {
	Event *OwnableUpgradeableInitialized // Event containing the contract specifics and raw log

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
func (it *OwnableUpgradeableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableUpgradeableInitialized)
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
		it.Event = new(OwnableUpgradeableInitialized)
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
func (it *OwnableUpgradeableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableUpgradeableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableUpgradeableInitialized represents a Initialized event raised by the OwnableUpgradeable contract.
type OwnableUpgradeableInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OwnableUpgradeable *OwnableUpgradeableFilterer) FilterInitialized(opts *bind.FilterOpts) (*OwnableUpgradeableInitializedIterator, error) {

	logs, sub, err := _OwnableUpgradeable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableInitializedIterator{contract: _OwnableUpgradeable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OwnableUpgradeable *OwnableUpgradeableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OwnableUpgradeableInitialized) (event.Subscription, error) {

	logs, sub, err := _OwnableUpgradeable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableUpgradeableInitialized)
				if err := _OwnableUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OwnableUpgradeable *OwnableUpgradeableFilterer) ParseInitialized(log types.Log) (*OwnableUpgradeableInitialized, error) {
	event := new(OwnableUpgradeableInitialized)
	if err := _OwnableUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnableUpgradeableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OwnableUpgradeable contract.
type OwnableUpgradeableOwnershipTransferredIterator struct {
	Event *OwnableUpgradeableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableUpgradeableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableUpgradeableOwnershipTransferred)
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
		it.Event = new(OwnableUpgradeableOwnershipTransferred)
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
func (it *OwnableUpgradeableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableUpgradeableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableUpgradeableOwnershipTransferred represents a OwnershipTransferred event raised by the OwnableUpgradeable contract.
type OwnableUpgradeableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnableUpgradeable *OwnableUpgradeableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableUpgradeableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OwnableUpgradeable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableOwnershipTransferredIterator{contract: _OwnableUpgradeable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnableUpgradeable *OwnableUpgradeableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableUpgradeableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OwnableUpgradeable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableUpgradeableOwnershipTransferred)
				if err := _OwnableUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OwnableUpgradeable *OwnableUpgradeableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableUpgradeableOwnershipTransferred, error) {
	event := new(OwnableUpgradeableOwnershipTransferred)
	if err := _OwnableUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReentrancyGuardedMetaData contains all meta data concerning the ReentrancyGuarded contract.
var ReentrancyGuardedMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60806040526000805460ff19169055348015601957600080fd5b50603f8060276000396000f3fe6080604052600080fdfea26469706673582212205083eccff38dd4087dae34ed3d1ebfa6605fbd83643e6850a819cf6ad2fc552e64736f6c63430008110033",
}

// ReentrancyGuardedABI is the input ABI used to generate the binding from.
// Deprecated: Use ReentrancyGuardedMetaData.ABI instead.
var ReentrancyGuardedABI = ReentrancyGuardedMetaData.ABI

// ReentrancyGuardedBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ReentrancyGuardedMetaData.Bin instead.
var ReentrancyGuardedBin = ReentrancyGuardedMetaData.Bin

// DeployReentrancyGuarded deploys a new Ethereum contract, binding an instance of ReentrancyGuarded to it.
func DeployReentrancyGuarded(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ReentrancyGuarded, error) {
	parsed, err := ReentrancyGuardedMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReentrancyGuardedBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReentrancyGuarded{ReentrancyGuardedCaller: ReentrancyGuardedCaller{contract: contract}, ReentrancyGuardedTransactor: ReentrancyGuardedTransactor{contract: contract}, ReentrancyGuardedFilterer: ReentrancyGuardedFilterer{contract: contract}}, nil
}

// ReentrancyGuarded is an auto generated Go binding around an Ethereum contract.
type ReentrancyGuarded struct {
	ReentrancyGuardedCaller     // Read-only binding to the contract
	ReentrancyGuardedTransactor // Write-only binding to the contract
	ReentrancyGuardedFilterer   // Log filterer for contract events
}

// ReentrancyGuardedCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReentrancyGuardedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReentrancyGuardedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReentrancyGuardedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReentrancyGuardedSession struct {
	Contract     *ReentrancyGuarded // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ReentrancyGuardedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReentrancyGuardedCallerSession struct {
	Contract *ReentrancyGuardedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ReentrancyGuardedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReentrancyGuardedTransactorSession struct {
	Contract     *ReentrancyGuardedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ReentrancyGuardedRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReentrancyGuardedRaw struct {
	Contract *ReentrancyGuarded // Generic contract binding to access the raw methods on
}

// ReentrancyGuardedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReentrancyGuardedCallerRaw struct {
	Contract *ReentrancyGuardedCaller // Generic read-only contract binding to access the raw methods on
}

// ReentrancyGuardedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReentrancyGuardedTransactorRaw struct {
	Contract *ReentrancyGuardedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReentrancyGuarded creates a new instance of ReentrancyGuarded, bound to a specific deployed contract.
func NewReentrancyGuarded(address common.Address, backend bind.ContractBackend) (*ReentrancyGuarded, error) {
	contract, err := bindReentrancyGuarded(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuarded{ReentrancyGuardedCaller: ReentrancyGuardedCaller{contract: contract}, ReentrancyGuardedTransactor: ReentrancyGuardedTransactor{contract: contract}, ReentrancyGuardedFilterer: ReentrancyGuardedFilterer{contract: contract}}, nil
}

// NewReentrancyGuardedCaller creates a new read-only instance of ReentrancyGuarded, bound to a specific deployed contract.
func NewReentrancyGuardedCaller(address common.Address, caller bind.ContractCaller) (*ReentrancyGuardedCaller, error) {
	contract, err := bindReentrancyGuarded(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardedCaller{contract: contract}, nil
}

// NewReentrancyGuardedTransactor creates a new write-only instance of ReentrancyGuarded, bound to a specific deployed contract.
func NewReentrancyGuardedTransactor(address common.Address, transactor bind.ContractTransactor) (*ReentrancyGuardedTransactor, error) {
	contract, err := bindReentrancyGuarded(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardedTransactor{contract: contract}, nil
}

// NewReentrancyGuardedFilterer creates a new log filterer instance of ReentrancyGuarded, bound to a specific deployed contract.
func NewReentrancyGuardedFilterer(address common.Address, filterer bind.ContractFilterer) (*ReentrancyGuardedFilterer, error) {
	contract, err := bindReentrancyGuarded(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardedFilterer{contract: contract}, nil
}

// bindReentrancyGuarded binds a generic wrapper to an already deployed contract.
func bindReentrancyGuarded(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReentrancyGuardedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuarded *ReentrancyGuardedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuarded.Contract.ReentrancyGuardedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuarded *ReentrancyGuardedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuarded.Contract.ReentrancyGuardedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuarded *ReentrancyGuardedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuarded.Contract.ReentrancyGuardedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuarded *ReentrancyGuardedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuarded.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuarded *ReentrancyGuardedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuarded.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuarded *ReentrancyGuardedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuarded.Contract.contract.Transact(opts, method, params...)
}

// StorageSlotUpgradeableMetaData contains all meta data concerning the StorageSlotUpgradeable contract.
var StorageSlotUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122099a0199b5af84da2fcdc58f0554a8af6b5fadfa473719e9ad9e13441ded2343964736f6c63430008110033",
}

// StorageSlotUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageSlotUpgradeableMetaData.ABI instead.
var StorageSlotUpgradeableABI = StorageSlotUpgradeableMetaData.ABI

// StorageSlotUpgradeableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StorageSlotUpgradeableMetaData.Bin instead.
var StorageSlotUpgradeableBin = StorageSlotUpgradeableMetaData.Bin

// DeployStorageSlotUpgradeable deploys a new Ethereum contract, binding an instance of StorageSlotUpgradeable to it.
func DeployStorageSlotUpgradeable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StorageSlotUpgradeable, error) {
	parsed, err := StorageSlotUpgradeableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StorageSlotUpgradeableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StorageSlotUpgradeable{StorageSlotUpgradeableCaller: StorageSlotUpgradeableCaller{contract: contract}, StorageSlotUpgradeableTransactor: StorageSlotUpgradeableTransactor{contract: contract}, StorageSlotUpgradeableFilterer: StorageSlotUpgradeableFilterer{contract: contract}}, nil
}

// StorageSlotUpgradeable is an auto generated Go binding around an Ethereum contract.
type StorageSlotUpgradeable struct {
	StorageSlotUpgradeableCaller     // Read-only binding to the contract
	StorageSlotUpgradeableTransactor // Write-only binding to the contract
	StorageSlotUpgradeableFilterer   // Log filterer for contract events
}

// StorageSlotUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageSlotUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSlotUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageSlotUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSlotUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageSlotUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSlotUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageSlotUpgradeableSession struct {
	Contract     *StorageSlotUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StorageSlotUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageSlotUpgradeableCallerSession struct {
	Contract *StorageSlotUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// StorageSlotUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageSlotUpgradeableTransactorSession struct {
	Contract     *StorageSlotUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// StorageSlotUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageSlotUpgradeableRaw struct {
	Contract *StorageSlotUpgradeable // Generic contract binding to access the raw methods on
}

// StorageSlotUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageSlotUpgradeableCallerRaw struct {
	Contract *StorageSlotUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// StorageSlotUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageSlotUpgradeableTransactorRaw struct {
	Contract *StorageSlotUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorageSlotUpgradeable creates a new instance of StorageSlotUpgradeable, bound to a specific deployed contract.
func NewStorageSlotUpgradeable(address common.Address, backend bind.ContractBackend) (*StorageSlotUpgradeable, error) {
	contract, err := bindStorageSlotUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StorageSlotUpgradeable{StorageSlotUpgradeableCaller: StorageSlotUpgradeableCaller{contract: contract}, StorageSlotUpgradeableTransactor: StorageSlotUpgradeableTransactor{contract: contract}, StorageSlotUpgradeableFilterer: StorageSlotUpgradeableFilterer{contract: contract}}, nil
}

// NewStorageSlotUpgradeableCaller creates a new read-only instance of StorageSlotUpgradeable, bound to a specific deployed contract.
func NewStorageSlotUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*StorageSlotUpgradeableCaller, error) {
	contract, err := bindStorageSlotUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageSlotUpgradeableCaller{contract: contract}, nil
}

// NewStorageSlotUpgradeableTransactor creates a new write-only instance of StorageSlotUpgradeable, bound to a specific deployed contract.
func NewStorageSlotUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageSlotUpgradeableTransactor, error) {
	contract, err := bindStorageSlotUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageSlotUpgradeableTransactor{contract: contract}, nil
}

// NewStorageSlotUpgradeableFilterer creates a new log filterer instance of StorageSlotUpgradeable, bound to a specific deployed contract.
func NewStorageSlotUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageSlotUpgradeableFilterer, error) {
	contract, err := bindStorageSlotUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageSlotUpgradeableFilterer{contract: contract}, nil
}

// bindStorageSlotUpgradeable binds a generic wrapper to an already deployed contract.
func bindStorageSlotUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StorageSlotUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageSlotUpgradeable *StorageSlotUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageSlotUpgradeable.Contract.StorageSlotUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageSlotUpgradeable *StorageSlotUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageSlotUpgradeable.Contract.StorageSlotUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageSlotUpgradeable *StorageSlotUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageSlotUpgradeable.Contract.StorageSlotUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageSlotUpgradeable *StorageSlotUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageSlotUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageSlotUpgradeable *StorageSlotUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageSlotUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageSlotUpgradeable *StorageSlotUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageSlotUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// UUPSUpgradeableMetaData contains all meta data concerning the UUPSUpgradeable contract.
var UUPSUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// UUPSUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use UUPSUpgradeableMetaData.ABI instead.
var UUPSUpgradeableABI = UUPSUpgradeableMetaData.ABI

// UUPSUpgradeable is an auto generated Go binding around an Ethereum contract.
type UUPSUpgradeable struct {
	UUPSUpgradeableCaller     // Read-only binding to the contract
	UUPSUpgradeableTransactor // Write-only binding to the contract
	UUPSUpgradeableFilterer   // Log filterer for contract events
}

// UUPSUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type UUPSUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UUPSUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UUPSUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UUPSUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UUPSUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UUPSUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UUPSUpgradeableSession struct {
	Contract     *UUPSUpgradeable  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UUPSUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UUPSUpgradeableCallerSession struct {
	Contract *UUPSUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// UUPSUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UUPSUpgradeableTransactorSession struct {
	Contract     *UUPSUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// UUPSUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type UUPSUpgradeableRaw struct {
	Contract *UUPSUpgradeable // Generic contract binding to access the raw methods on
}

// UUPSUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UUPSUpgradeableCallerRaw struct {
	Contract *UUPSUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// UUPSUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UUPSUpgradeableTransactorRaw struct {
	Contract *UUPSUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUUPSUpgradeable creates a new instance of UUPSUpgradeable, bound to a specific deployed contract.
func NewUUPSUpgradeable(address common.Address, backend bind.ContractBackend) (*UUPSUpgradeable, error) {
	contract, err := bindUUPSUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UUPSUpgradeable{UUPSUpgradeableCaller: UUPSUpgradeableCaller{contract: contract}, UUPSUpgradeableTransactor: UUPSUpgradeableTransactor{contract: contract}, UUPSUpgradeableFilterer: UUPSUpgradeableFilterer{contract: contract}}, nil
}

// NewUUPSUpgradeableCaller creates a new read-only instance of UUPSUpgradeable, bound to a specific deployed contract.
func NewUUPSUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*UUPSUpgradeableCaller, error) {
	contract, err := bindUUPSUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UUPSUpgradeableCaller{contract: contract}, nil
}

// NewUUPSUpgradeableTransactor creates a new write-only instance of UUPSUpgradeable, bound to a specific deployed contract.
func NewUUPSUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*UUPSUpgradeableTransactor, error) {
	contract, err := bindUUPSUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UUPSUpgradeableTransactor{contract: contract}, nil
}

// NewUUPSUpgradeableFilterer creates a new log filterer instance of UUPSUpgradeable, bound to a specific deployed contract.
func NewUUPSUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*UUPSUpgradeableFilterer, error) {
	contract, err := bindUUPSUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UUPSUpgradeableFilterer{contract: contract}, nil
}

// bindUUPSUpgradeable binds a generic wrapper to an already deployed contract.
func bindUUPSUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UUPSUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UUPSUpgradeable *UUPSUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UUPSUpgradeable.Contract.UUPSUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UUPSUpgradeable *UUPSUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UUPSUpgradeable.Contract.UUPSUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UUPSUpgradeable *UUPSUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UUPSUpgradeable.Contract.UUPSUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UUPSUpgradeable *UUPSUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UUPSUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UUPSUpgradeable *UUPSUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UUPSUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UUPSUpgradeable *UUPSUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UUPSUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_UUPSUpgradeable *UUPSUpgradeableCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UUPSUpgradeable.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_UUPSUpgradeable *UUPSUpgradeableSession) ProxiableUUID() ([32]byte, error) {
	return _UUPSUpgradeable.Contract.ProxiableUUID(&_UUPSUpgradeable.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_UUPSUpgradeable *UUPSUpgradeableCallerSession) ProxiableUUID() ([32]byte, error) {
	return _UUPSUpgradeable.Contract.ProxiableUUID(&_UUPSUpgradeable.CallOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_UUPSUpgradeable *UUPSUpgradeableTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _UUPSUpgradeable.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_UUPSUpgradeable *UUPSUpgradeableSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _UUPSUpgradeable.Contract.UpgradeTo(&_UUPSUpgradeable.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_UUPSUpgradeable *UUPSUpgradeableTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _UUPSUpgradeable.Contract.UpgradeTo(&_UUPSUpgradeable.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_UUPSUpgradeable *UUPSUpgradeableTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _UUPSUpgradeable.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_UUPSUpgradeable *UUPSUpgradeableSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _UUPSUpgradeable.Contract.UpgradeToAndCall(&_UUPSUpgradeable.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_UUPSUpgradeable *UUPSUpgradeableTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _UUPSUpgradeable.Contract.UpgradeToAndCall(&_UUPSUpgradeable.TransactOpts, newImplementation, data)
}

// UUPSUpgradeableAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the UUPSUpgradeable contract.
type UUPSUpgradeableAdminChangedIterator struct {
	Event *UUPSUpgradeableAdminChanged // Event containing the contract specifics and raw log

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
func (it *UUPSUpgradeableAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UUPSUpgradeableAdminChanged)
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
		it.Event = new(UUPSUpgradeableAdminChanged)
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
func (it *UUPSUpgradeableAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UUPSUpgradeableAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UUPSUpgradeableAdminChanged represents a AdminChanged event raised by the UUPSUpgradeable contract.
type UUPSUpgradeableAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*UUPSUpgradeableAdminChangedIterator, error) {

	logs, sub, err := _UUPSUpgradeable.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &UUPSUpgradeableAdminChangedIterator{contract: _UUPSUpgradeable.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *UUPSUpgradeableAdminChanged) (event.Subscription, error) {

	logs, sub, err := _UUPSUpgradeable.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UUPSUpgradeableAdminChanged)
				if err := _UUPSUpgradeable.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) ParseAdminChanged(log types.Log) (*UUPSUpgradeableAdminChanged, error) {
	event := new(UUPSUpgradeableAdminChanged)
	if err := _UUPSUpgradeable.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UUPSUpgradeableBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the UUPSUpgradeable contract.
type UUPSUpgradeableBeaconUpgradedIterator struct {
	Event *UUPSUpgradeableBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *UUPSUpgradeableBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UUPSUpgradeableBeaconUpgraded)
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
		it.Event = new(UUPSUpgradeableBeaconUpgraded)
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
func (it *UUPSUpgradeableBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UUPSUpgradeableBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UUPSUpgradeableBeaconUpgraded represents a BeaconUpgraded event raised by the UUPSUpgradeable contract.
type UUPSUpgradeableBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*UUPSUpgradeableBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _UUPSUpgradeable.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &UUPSUpgradeableBeaconUpgradedIterator{contract: _UUPSUpgradeable.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *UUPSUpgradeableBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _UUPSUpgradeable.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UUPSUpgradeableBeaconUpgraded)
				if err := _UUPSUpgradeable.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) ParseBeaconUpgraded(log types.Log) (*UUPSUpgradeableBeaconUpgraded, error) {
	event := new(UUPSUpgradeableBeaconUpgraded)
	if err := _UUPSUpgradeable.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UUPSUpgradeableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the UUPSUpgradeable contract.
type UUPSUpgradeableInitializedIterator struct {
	Event *UUPSUpgradeableInitialized // Event containing the contract specifics and raw log

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
func (it *UUPSUpgradeableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UUPSUpgradeableInitialized)
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
		it.Event = new(UUPSUpgradeableInitialized)
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
func (it *UUPSUpgradeableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UUPSUpgradeableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UUPSUpgradeableInitialized represents a Initialized event raised by the UUPSUpgradeable contract.
type UUPSUpgradeableInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) FilterInitialized(opts *bind.FilterOpts) (*UUPSUpgradeableInitializedIterator, error) {

	logs, sub, err := _UUPSUpgradeable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &UUPSUpgradeableInitializedIterator{contract: _UUPSUpgradeable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *UUPSUpgradeableInitialized) (event.Subscription, error) {

	logs, sub, err := _UUPSUpgradeable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UUPSUpgradeableInitialized)
				if err := _UUPSUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) ParseInitialized(log types.Log) (*UUPSUpgradeableInitialized, error) {
	event := new(UUPSUpgradeableInitialized)
	if err := _UUPSUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UUPSUpgradeableUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the UUPSUpgradeable contract.
type UUPSUpgradeableUpgradedIterator struct {
	Event *UUPSUpgradeableUpgraded // Event containing the contract specifics and raw log

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
func (it *UUPSUpgradeableUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UUPSUpgradeableUpgraded)
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
		it.Event = new(UUPSUpgradeableUpgraded)
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
func (it *UUPSUpgradeableUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UUPSUpgradeableUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UUPSUpgradeableUpgraded represents a Upgraded event raised by the UUPSUpgradeable contract.
type UUPSUpgradeableUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*UUPSUpgradeableUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _UUPSUpgradeable.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &UUPSUpgradeableUpgradedIterator{contract: _UUPSUpgradeable.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *UUPSUpgradeableUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _UUPSUpgradeable.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UUPSUpgradeableUpgraded)
				if err := _UUPSUpgradeable.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) ParseUpgraded(log types.Log) (*UUPSUpgradeableUpgraded, error) {
	event := new(UUPSUpgradeableUpgraded)
	if err := _UUPSUpgradeable.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
