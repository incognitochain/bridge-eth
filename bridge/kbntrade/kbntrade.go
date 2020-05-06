// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kbntrade

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

// KbntradeABI is the input ABI used to generate the binding from.
const KbntradeABI = "[{\"inputs\":[{\"internalType\":\"contractKyberNetwork\",\"name\":\"_kyberNetworkProxyContract\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_incognitoSmartContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ETH_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcQty\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"destToken\",\"type\":\"address\"}],\"name\":\"getConversionRates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incognitoSmartContract\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kyberNetworkProxyContract\",\"outputs\":[{\"internalType\":\"contractKyberNetwork\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcQty\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minConversionRate\",\"type\":\"uint256\"}],\"name\":\"trade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// KbntradeBin is the compiled bytecode used for deploying new contracts.
var KbntradeBin = "0x608060405234801561001057600080fd5b50604051610e7f380380610e7f8339818101604052604081101561003357600080fd5b81019080805190602001909291908051906020019092919050505081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050610d9f806100e06000396000f3fe60806040526004361061004e5760003560e01c80630aea81881461005a57806372e94bf6146100f0578063785250da14610147578063b42a644b1461019e578063bb39a960146101f557610055565b3661005557005b600080fd5b34801561006657600080fd5b506100d36004803603606081101561007d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506102b4565b604051808381526020018281526020019250505060405180910390f35b3480156100fc57600080fd5b506101056103e2565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561015357600080fd5b5061015c6103e7565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156101aa57600080fd5b506101b361040d565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61026b6004803603608081101561020b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610432565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b600080600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663809a9e558685876040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050604080518083038186803b15801561039157600080fd5b505afa1580156103a5573d6000803e3d6000fd5b505050506040513d60408110156103bb57600080fd5b81019080805190602001909291908051906020019092919050505091509150935093915050565b600081565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461048e57600080fd5b84610498876105ed565b10156104a357600080fd5b8373ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1614156104dc57600080fd5b6000809050600073ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff16146105b05761054287600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16886106eb565b600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff161461059357600061058488888888610886565b1161058e57600080fd5b6105ab565b60006105a08888876109b2565b116105aa57600080fd5b5b6105c8565b60006105bd868887610aa9565b116105c757600080fd5b5b6105d1856105ed565b90506105dd8582610ba6565b8481925092505094509492505050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561062b574790506106e6565b8173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b1580156106a857600080fd5b505afa1580156106bc573d6000803e3d6000fd5b505050506040513d60208110156106d257600080fd5b810190808051906020019092919050505090505b919050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614610881578273ffffffffffffffffffffffffffffffffffffffff1663095ea7b38360006040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b1580156107a757600080fd5b505af11580156107bb573d6000803e3d6000fd5b505050506107c7610d2b565b6107d057600080fd5b8273ffffffffffffffffffffffffffffffffffffffff1663095ea7b383836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b15801561085757600080fd5b505af115801561086b573d6000803e3d6000fd5b50505050610877610d2b565b61088057600080fd5b5b505050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637409e2eb868686866040518563ffffffff1660e01b8152600401808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018481526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001945050505050602060405180830381600087803b15801561096d57600080fd5b505af1158015610981573d6000803e3d6000fd5b505050506040513d602081101561099757600080fd5b81019080805190602001909291905050509050949350505050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633bba21dc8585856040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018281526020019350505050602060405180830381600087803b158015610a6557600080fd5b505af1158015610a79573d6000803e3d6000fd5b505050506040513d6020811015610a8f57600080fd5b810190808051906020019092919050505090509392505050565b600082471015610ab857600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637a2a04568486856040518463ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001925050506020604051808303818588803b158015610b6157600080fd5b505af1158015610b75573d6000803e3d6000fd5b50505050506040513d6020811015610b8c57600080fd5b810190808051906020019092919050505090509392505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610c555780471015610be857600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015610c4f573d6000803e3d6000fd5b50610d27565b8173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b158015610cfd57600080fd5b505af1158015610d11573d6000803e3d6000fd5b50505050610d1d610d2b565b610d2657600080fd5b5b5050565b600080600090503d60008114610d485760208114610d5157610d5d565b60019150610d5d565b60206000803e60005191505b5060008114159150509056fea26469706673582212207401a6a4e9f58f14349867a5c8b979105c920ec788bd82832151c50b13c8dc9b64736f6c63430006060033"

// DeployKbntrade deploys a new Ethereum contract, binding an instance of Kbntrade to it.
func DeployKbntrade(auth *bind.TransactOpts, backend bind.ContractBackend, _kyberNetworkProxyContract common.Address, _incognitoSmartContract common.Address) (common.Address, *types.Transaction, *Kbntrade, error) {
	parsed, err := abi.JSON(strings.NewReader(KbntradeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KbntradeBin), backend, _kyberNetworkProxyContract, _incognitoSmartContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Kbntrade{KbntradeCaller: KbntradeCaller{contract: contract}, KbntradeTransactor: KbntradeTransactor{contract: contract}, KbntradeFilterer: KbntradeFilterer{contract: contract}}, nil
}

// Kbntrade is an auto generated Go binding around an Ethereum contract.
type Kbntrade struct {
	KbntradeCaller     // Read-only binding to the contract
	KbntradeTransactor // Write-only binding to the contract
	KbntradeFilterer   // Log filterer for contract events
}

// KbntradeCaller is an auto generated read-only Go binding around an Ethereum contract.
type KbntradeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KbntradeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KbntradeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KbntradeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KbntradeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KbntradeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KbntradeSession struct {
	Contract     *Kbntrade         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KbntradeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KbntradeCallerSession struct {
	Contract *KbntradeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// KbntradeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KbntradeTransactorSession struct {
	Contract     *KbntradeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// KbntradeRaw is an auto generated low-level Go binding around an Ethereum contract.
type KbntradeRaw struct {
	Contract *Kbntrade // Generic contract binding to access the raw methods on
}

// KbntradeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KbntradeCallerRaw struct {
	Contract *KbntradeCaller // Generic read-only contract binding to access the raw methods on
}

// KbntradeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KbntradeTransactorRaw struct {
	Contract *KbntradeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKbntrade creates a new instance of Kbntrade, bound to a specific deployed contract.
func NewKbntrade(address common.Address, backend bind.ContractBackend) (*Kbntrade, error) {
	contract, err := bindKbntrade(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Kbntrade{KbntradeCaller: KbntradeCaller{contract: contract}, KbntradeTransactor: KbntradeTransactor{contract: contract}, KbntradeFilterer: KbntradeFilterer{contract: contract}}, nil
}

// NewKbntradeCaller creates a new read-only instance of Kbntrade, bound to a specific deployed contract.
func NewKbntradeCaller(address common.Address, caller bind.ContractCaller) (*KbntradeCaller, error) {
	contract, err := bindKbntrade(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KbntradeCaller{contract: contract}, nil
}

// NewKbntradeTransactor creates a new write-only instance of Kbntrade, bound to a specific deployed contract.
func NewKbntradeTransactor(address common.Address, transactor bind.ContractTransactor) (*KbntradeTransactor, error) {
	contract, err := bindKbntrade(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KbntradeTransactor{contract: contract}, nil
}

// NewKbntradeFilterer creates a new log filterer instance of Kbntrade, bound to a specific deployed contract.
func NewKbntradeFilterer(address common.Address, filterer bind.ContractFilterer) (*KbntradeFilterer, error) {
	contract, err := bindKbntrade(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KbntradeFilterer{contract: contract}, nil
}

// bindKbntrade binds a generic wrapper to an already deployed contract.
func bindKbntrade(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KbntradeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kbntrade *KbntradeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Kbntrade.Contract.KbntradeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kbntrade *KbntradeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kbntrade.Contract.KbntradeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kbntrade *KbntradeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kbntrade.Contract.KbntradeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kbntrade *KbntradeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Kbntrade.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kbntrade *KbntradeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kbntrade.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kbntrade *KbntradeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kbntrade.Contract.contract.Transact(opts, method, params...)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_Kbntrade *KbntradeCaller) ETHCONTRACTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Kbntrade.contract.Call(opts, out, "ETH_CONTRACT_ADDRESS")
	return *ret0, err
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_Kbntrade *KbntradeSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _Kbntrade.Contract.ETHCONTRACTADDRESS(&_Kbntrade.CallOpts)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_Kbntrade *KbntradeCallerSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _Kbntrade.Contract.ETHCONTRACTADDRESS(&_Kbntrade.CallOpts)
}

// GetConversionRates is a free data retrieval call binding the contract method 0x0aea8188.
//
// Solidity: function getConversionRates(address srcToken, uint256 srcQty, address destToken) constant returns(uint256, uint256)
func (_Kbntrade *KbntradeCaller) GetConversionRates(opts *bind.CallOpts, srcToken common.Address, srcQty *big.Int, destToken common.Address) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Kbntrade.contract.Call(opts, out, "getConversionRates", srcToken, srcQty, destToken)
	return *ret0, *ret1, err
}

// GetConversionRates is a free data retrieval call binding the contract method 0x0aea8188.
//
// Solidity: function getConversionRates(address srcToken, uint256 srcQty, address destToken) constant returns(uint256, uint256)
func (_Kbntrade *KbntradeSession) GetConversionRates(srcToken common.Address, srcQty *big.Int, destToken common.Address) (*big.Int, *big.Int, error) {
	return _Kbntrade.Contract.GetConversionRates(&_Kbntrade.CallOpts, srcToken, srcQty, destToken)
}

// GetConversionRates is a free data retrieval call binding the contract method 0x0aea8188.
//
// Solidity: function getConversionRates(address srcToken, uint256 srcQty, address destToken) constant returns(uint256, uint256)
func (_Kbntrade *KbntradeCallerSession) GetConversionRates(srcToken common.Address, srcQty *big.Int, destToken common.Address) (*big.Int, *big.Int, error) {
	return _Kbntrade.Contract.GetConversionRates(&_Kbntrade.CallOpts, srcToken, srcQty, destToken)
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_Kbntrade *KbntradeCaller) IncognitoSmartContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Kbntrade.contract.Call(opts, out, "incognitoSmartContract")
	return *ret0, err
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_Kbntrade *KbntradeSession) IncognitoSmartContract() (common.Address, error) {
	return _Kbntrade.Contract.IncognitoSmartContract(&_Kbntrade.CallOpts)
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_Kbntrade *KbntradeCallerSession) IncognitoSmartContract() (common.Address, error) {
	return _Kbntrade.Contract.IncognitoSmartContract(&_Kbntrade.CallOpts)
}

// KyberNetworkProxyContract is a free data retrieval call binding the contract method 0x785250da.
//
// Solidity: function kyberNetworkProxyContract() constant returns(address)
func (_Kbntrade *KbntradeCaller) KyberNetworkProxyContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Kbntrade.contract.Call(opts, out, "kyberNetworkProxyContract")
	return *ret0, err
}

// KyberNetworkProxyContract is a free data retrieval call binding the contract method 0x785250da.
//
// Solidity: function kyberNetworkProxyContract() constant returns(address)
func (_Kbntrade *KbntradeSession) KyberNetworkProxyContract() (common.Address, error) {
	return _Kbntrade.Contract.KyberNetworkProxyContract(&_Kbntrade.CallOpts)
}

// KyberNetworkProxyContract is a free data retrieval call binding the contract method 0x785250da.
//
// Solidity: function kyberNetworkProxyContract() constant returns(address)
func (_Kbntrade *KbntradeCallerSession) KyberNetworkProxyContract() (common.Address, error) {
	return _Kbntrade.Contract.KyberNetworkProxyContract(&_Kbntrade.CallOpts)
}

// Trade is a paid mutator transaction binding the contract method 0xbb39a960.
//
// Solidity: function trade(address srcToken, uint256 srcQty, address destToken, uint256 minConversionRate) returns(address, uint256)
func (_Kbntrade *KbntradeTransactor) Trade(opts *bind.TransactOpts, srcToken common.Address, srcQty *big.Int, destToken common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _Kbntrade.contract.Transact(opts, "trade", srcToken, srcQty, destToken, minConversionRate)
}

// Trade is a paid mutator transaction binding the contract method 0xbb39a960.
//
// Solidity: function trade(address srcToken, uint256 srcQty, address destToken, uint256 minConversionRate) returns(address, uint256)
func (_Kbntrade *KbntradeSession) Trade(srcToken common.Address, srcQty *big.Int, destToken common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _Kbntrade.Contract.Trade(&_Kbntrade.TransactOpts, srcToken, srcQty, destToken, minConversionRate)
}

// Trade is a paid mutator transaction binding the contract method 0xbb39a960.
//
// Solidity: function trade(address srcToken, uint256 srcQty, address destToken, uint256 minConversionRate) returns(address, uint256)
func (_Kbntrade *KbntradeTransactorSession) Trade(srcToken common.Address, srcQty *big.Int, destToken common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _Kbntrade.Contract.Trade(&_Kbntrade.TransactOpts, srcToken, srcQty, destToken, minConversionRate)
}
