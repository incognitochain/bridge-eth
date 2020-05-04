// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kbnmultiTrade

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

// KbnmultiTradeABI is the input ABI used to generate the binding from.
const KbnmultiTradeABI = "[{\"inputs\":[{\"internalType\":\"contractKyberNetwork\",\"name\":\"_kyberNetworkProxyContract\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_incognitoSmartContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ETH_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcQty\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"destToken\",\"type\":\"address\"}],\"name\":\"getConversionRates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incognitoSmartContract\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kyberNetworkProxyContract\",\"outputs\":[{\"internalType\":\"contractKyberNetwork\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"srcTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"srcQties\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"destTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"minConversionRates\",\"type\":\"uint256[]\"}],\"name\":\"trade\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// KbnmultiTradeBin is the compiled bytecode used for deploying new contracts.
var KbnmultiTradeBin = "0x608060405234801561001057600080fd5b506040516113033803806113038339818101604052604081101561003357600080fd5b81019080805190602001909291908051906020019092919050505081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050611223806100e06000396000f3fe60806040526004361061004e5760003560e01c80630aea81881461005a578063398b4d9b146100f057806372e94bf614610401578063785250da14610458578063b42a644b146104af57610055565b3661005557005b600080fd5b34801561006657600080fd5b506100d36004803603606081101561007d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610506565b604051808381526020018281526020019250505060405180910390f35b6103626004803603608081101561010657600080fd5b810190808035906020019064010000000081111561012357600080fd5b82018360208201111561013557600080fd5b8035906020019184602083028401116401000000008311171561015757600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050509192919290803590602001906401000000008111156101b757600080fd5b8201836020820111156101c957600080fd5b803590602001918460208302840111640100000000831117156101eb57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561024b57600080fd5b82018360208201111561025d57600080fd5b8035906020019184602083028401116401000000008311171561027f57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050509192919290803590602001906401000000008111156102df57600080fd5b8201836020820111156102f157600080fd5b8035906020019184602083028401116401000000008311171561031357600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050509192919290505050610634565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b838110156103a957808201518184015260208101905061038e565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156103eb5780820151818401526020810190506103d0565b5050505090500194505050505060405180910390f35b34801561040d57600080fd5b50610416610a21565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561046457600080fd5b5061046d610a26565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156104bb57600080fd5b506104c4610a4c565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b600080600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663809a9e558685876040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050604080518083038186803b1580156105e357600080fd5b505afa1580156105f7573d6000803e3d6000fd5b505050506040513d604081101561060d57600080fd5b81019080805190602001909291908051906020019092919050505091509150935093915050565b6060806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461069057600080fd5b845186511480156106a2575085518451145b6106ab57600080fd5b82518451146106b957600080fd5b6060845167ffffffffffffffff811180156106d357600080fd5b506040519080825280602002602001820160405280156107025781602001602082028036833780820191505090505b50905060008090505b8751811015610a105786818151811061072057fe5b602002602001015161074489838151811061073757fe5b6020026020010151610a71565b101561074f57600080fd5b85818151811061075b57fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1688828151811061078557fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1614156107ae57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168882815181106107d257fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff161461094e5761084888828151811061080457fe5b6020026020010151600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1689848151811061083b57fe5b6020026020010151610b6f565b600073ffffffffffffffffffffffffffffffffffffffff1686828151811061086c57fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff16146108f85760006108e98983815181106108a057fe5b60200260200101518984815181106108b457fe5b60200260200101518985815181106108c857fe5b60200260200101518986815181106108dc57fe5b6020026020010151610d0a565b116108f357600080fd5b610949565b600061093e89838151811061090957fe5b602002602001015189848151811061091d57fe5b602002602001015188858151811061093157fe5b6020026020010151610e36565b1161094857600080fd5b5b61099f565b600061099487838151811061095f57fe5b602002602001015189848151811061097357fe5b602002602001015188858151811061098757fe5b6020026020010151610f2d565b1161099e57600080fd5b5b6109bb8682815181106109ae57fe5b6020026020010151610a71565b8282815181106109c757fe5b602002602001018181525050610a038682815181106109e257fe5b60200260200101518383815181106109f657fe5b602002602001015161102a565b808060010191505061070b565b508481925092505094509492505050565b600081565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610aaf57479050610b6a565b8173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015610b2c57600080fd5b505afa158015610b40573d6000803e3d6000fd5b505050506040513d6020811015610b5657600080fd5b810190808051906020019092919050505090505b919050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614610d05578273ffffffffffffffffffffffffffffffffffffffff1663095ea7b38360006040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b158015610c2b57600080fd5b505af1158015610c3f573d6000803e3d6000fd5b50505050610c4b6111af565b610c5457600080fd5b8273ffffffffffffffffffffffffffffffffffffffff1663095ea7b383836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b158015610cdb57600080fd5b505af1158015610cef573d6000803e3d6000fd5b50505050610cfb6111af565b610d0457600080fd5b5b505050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637409e2eb868686866040518563ffffffff1660e01b8152600401808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018481526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001945050505050602060405180830381600087803b158015610df157600080fd5b505af1158015610e05573d6000803e3d6000fd5b505050506040513d6020811015610e1b57600080fd5b81019080805190602001909291905050509050949350505050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633bba21dc8585856040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018281526020019350505050602060405180830381600087803b158015610ee957600080fd5b505af1158015610efd573d6000803e3d6000fd5b505050506040513d6020811015610f1357600080fd5b810190808051906020019092919050505090509392505050565b600082471015610f3c57600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637a2a04568486856040518463ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001925050506020604051808303818588803b158015610fe557600080fd5b505af1158015610ff9573d6000803e3d6000fd5b50505050506040513d602081101561101057600080fd5b810190808051906020019092919050505090509392505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156110d9578047101561106c57600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156110d3573d6000803e3d6000fd5b506111ab565b8173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b15801561118157600080fd5b505af1158015611195573d6000803e3d6000fd5b505050506111a16111af565b6111aa57600080fd5b5b5050565b600080600090503d600081146111cc57602081146111d5576111e1565b600191506111e1565b60206000803e60005191505b5060008114159150509056fea264697066735822122025e1110adb0c884395a0a6fdb3e27d13d525dbd22bf817876f548315e19cfa2064736f6c63430006060033"

// DeployKbnmultiTrade deploys a new Ethereum contract, binding an instance of KbnmultiTrade to it.
func DeployKbnmultiTrade(auth *bind.TransactOpts, backend bind.ContractBackend, _kyberNetworkProxyContract common.Address, _incognitoSmartContract common.Address) (common.Address, *types.Transaction, *KbnmultiTrade, error) {
	parsed, err := abi.JSON(strings.NewReader(KbnmultiTradeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KbnmultiTradeBin), backend, _kyberNetworkProxyContract, _incognitoSmartContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KbnmultiTrade{KbnmultiTradeCaller: KbnmultiTradeCaller{contract: contract}, KbnmultiTradeTransactor: KbnmultiTradeTransactor{contract: contract}, KbnmultiTradeFilterer: KbnmultiTradeFilterer{contract: contract}}, nil
}

// KbnmultiTrade is an auto generated Go binding around an Ethereum contract.
type KbnmultiTrade struct {
	KbnmultiTradeCaller     // Read-only binding to the contract
	KbnmultiTradeTransactor // Write-only binding to the contract
	KbnmultiTradeFilterer   // Log filterer for contract events
}

// KbnmultiTradeCaller is an auto generated read-only Go binding around an Ethereum contract.
type KbnmultiTradeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KbnmultiTradeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KbnmultiTradeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KbnmultiTradeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KbnmultiTradeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KbnmultiTradeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KbnmultiTradeSession struct {
	Contract     *KbnmultiTrade    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KbnmultiTradeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KbnmultiTradeCallerSession struct {
	Contract *KbnmultiTradeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// KbnmultiTradeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KbnmultiTradeTransactorSession struct {
	Contract     *KbnmultiTradeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// KbnmultiTradeRaw is an auto generated low-level Go binding around an Ethereum contract.
type KbnmultiTradeRaw struct {
	Contract *KbnmultiTrade // Generic contract binding to access the raw methods on
}

// KbnmultiTradeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KbnmultiTradeCallerRaw struct {
	Contract *KbnmultiTradeCaller // Generic read-only contract binding to access the raw methods on
}

// KbnmultiTradeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KbnmultiTradeTransactorRaw struct {
	Contract *KbnmultiTradeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKbnmultiTrade creates a new instance of KbnmultiTrade, bound to a specific deployed contract.
func NewKbnmultiTrade(address common.Address, backend bind.ContractBackend) (*KbnmultiTrade, error) {
	contract, err := bindKbnmultiTrade(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KbnmultiTrade{KbnmultiTradeCaller: KbnmultiTradeCaller{contract: contract}, KbnmultiTradeTransactor: KbnmultiTradeTransactor{contract: contract}, KbnmultiTradeFilterer: KbnmultiTradeFilterer{contract: contract}}, nil
}

// NewKbnmultiTradeCaller creates a new read-only instance of KbnmultiTrade, bound to a specific deployed contract.
func NewKbnmultiTradeCaller(address common.Address, caller bind.ContractCaller) (*KbnmultiTradeCaller, error) {
	contract, err := bindKbnmultiTrade(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KbnmultiTradeCaller{contract: contract}, nil
}

// NewKbnmultiTradeTransactor creates a new write-only instance of KbnmultiTrade, bound to a specific deployed contract.
func NewKbnmultiTradeTransactor(address common.Address, transactor bind.ContractTransactor) (*KbnmultiTradeTransactor, error) {
	contract, err := bindKbnmultiTrade(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KbnmultiTradeTransactor{contract: contract}, nil
}

// NewKbnmultiTradeFilterer creates a new log filterer instance of KbnmultiTrade, bound to a specific deployed contract.
func NewKbnmultiTradeFilterer(address common.Address, filterer bind.ContractFilterer) (*KbnmultiTradeFilterer, error) {
	contract, err := bindKbnmultiTrade(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KbnmultiTradeFilterer{contract: contract}, nil
}

// bindKbnmultiTrade binds a generic wrapper to an already deployed contract.
func bindKbnmultiTrade(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KbnmultiTradeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KbnmultiTrade *KbnmultiTradeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KbnmultiTrade.Contract.KbnmultiTradeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KbnmultiTrade *KbnmultiTradeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KbnmultiTrade.Contract.KbnmultiTradeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KbnmultiTrade *KbnmultiTradeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KbnmultiTrade.Contract.KbnmultiTradeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KbnmultiTrade *KbnmultiTradeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KbnmultiTrade.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KbnmultiTrade *KbnmultiTradeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KbnmultiTrade.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KbnmultiTrade *KbnmultiTradeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KbnmultiTrade.Contract.contract.Transact(opts, method, params...)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_KbnmultiTrade *KbnmultiTradeCaller) ETHCONTRACTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KbnmultiTrade.contract.Call(opts, out, "ETH_CONTRACT_ADDRESS")
	return *ret0, err
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_KbnmultiTrade *KbnmultiTradeSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _KbnmultiTrade.Contract.ETHCONTRACTADDRESS(&_KbnmultiTrade.CallOpts)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_KbnmultiTrade *KbnmultiTradeCallerSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _KbnmultiTrade.Contract.ETHCONTRACTADDRESS(&_KbnmultiTrade.CallOpts)
}

// GetConversionRates is a free data retrieval call binding the contract method 0x0aea8188.
//
// Solidity: function getConversionRates(address srcToken, uint256 srcQty, address destToken) constant returns(uint256, uint256)
func (_KbnmultiTrade *KbnmultiTradeCaller) GetConversionRates(opts *bind.CallOpts, srcToken common.Address, srcQty *big.Int, destToken common.Address) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _KbnmultiTrade.contract.Call(opts, out, "getConversionRates", srcToken, srcQty, destToken)
	return *ret0, *ret1, err
}

// GetConversionRates is a free data retrieval call binding the contract method 0x0aea8188.
//
// Solidity: function getConversionRates(address srcToken, uint256 srcQty, address destToken) constant returns(uint256, uint256)
func (_KbnmultiTrade *KbnmultiTradeSession) GetConversionRates(srcToken common.Address, srcQty *big.Int, destToken common.Address) (*big.Int, *big.Int, error) {
	return _KbnmultiTrade.Contract.GetConversionRates(&_KbnmultiTrade.CallOpts, srcToken, srcQty, destToken)
}

// GetConversionRates is a free data retrieval call binding the contract method 0x0aea8188.
//
// Solidity: function getConversionRates(address srcToken, uint256 srcQty, address destToken) constant returns(uint256, uint256)
func (_KbnmultiTrade *KbnmultiTradeCallerSession) GetConversionRates(srcToken common.Address, srcQty *big.Int, destToken common.Address) (*big.Int, *big.Int, error) {
	return _KbnmultiTrade.Contract.GetConversionRates(&_KbnmultiTrade.CallOpts, srcToken, srcQty, destToken)
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_KbnmultiTrade *KbnmultiTradeCaller) IncognitoSmartContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KbnmultiTrade.contract.Call(opts, out, "incognitoSmartContract")
	return *ret0, err
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_KbnmultiTrade *KbnmultiTradeSession) IncognitoSmartContract() (common.Address, error) {
	return _KbnmultiTrade.Contract.IncognitoSmartContract(&_KbnmultiTrade.CallOpts)
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_KbnmultiTrade *KbnmultiTradeCallerSession) IncognitoSmartContract() (common.Address, error) {
	return _KbnmultiTrade.Contract.IncognitoSmartContract(&_KbnmultiTrade.CallOpts)
}

// KyberNetworkProxyContract is a free data retrieval call binding the contract method 0x785250da.
//
// Solidity: function kyberNetworkProxyContract() constant returns(address)
func (_KbnmultiTrade *KbnmultiTradeCaller) KyberNetworkProxyContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KbnmultiTrade.contract.Call(opts, out, "kyberNetworkProxyContract")
	return *ret0, err
}

// KyberNetworkProxyContract is a free data retrieval call binding the contract method 0x785250da.
//
// Solidity: function kyberNetworkProxyContract() constant returns(address)
func (_KbnmultiTrade *KbnmultiTradeSession) KyberNetworkProxyContract() (common.Address, error) {
	return _KbnmultiTrade.Contract.KyberNetworkProxyContract(&_KbnmultiTrade.CallOpts)
}

// KyberNetworkProxyContract is a free data retrieval call binding the contract method 0x785250da.
//
// Solidity: function kyberNetworkProxyContract() constant returns(address)
func (_KbnmultiTrade *KbnmultiTradeCallerSession) KyberNetworkProxyContract() (common.Address, error) {
	return _KbnmultiTrade.Contract.KyberNetworkProxyContract(&_KbnmultiTrade.CallOpts)
}

// Trade is a paid mutator transaction binding the contract method 0x398b4d9b.
//
// Solidity: function trade(address[] srcTokens, uint256[] srcQties, address[] destTokens, uint256[] minConversionRates) returns(address[], uint256[])
func (_KbnmultiTrade *KbnmultiTradeTransactor) Trade(opts *bind.TransactOpts, srcTokens []common.Address, srcQties []*big.Int, destTokens []common.Address, minConversionRates []*big.Int) (*types.Transaction, error) {
	return _KbnmultiTrade.contract.Transact(opts, "trade", srcTokens, srcQties, destTokens, minConversionRates)
}

// Trade is a paid mutator transaction binding the contract method 0x398b4d9b.
//
// Solidity: function trade(address[] srcTokens, uint256[] srcQties, address[] destTokens, uint256[] minConversionRates) returns(address[], uint256[])
func (_KbnmultiTrade *KbnmultiTradeSession) Trade(srcTokens []common.Address, srcQties []*big.Int, destTokens []common.Address, minConversionRates []*big.Int) (*types.Transaction, error) {
	return _KbnmultiTrade.Contract.Trade(&_KbnmultiTrade.TransactOpts, srcTokens, srcQties, destTokens, minConversionRates)
}

// Trade is a paid mutator transaction binding the contract method 0x398b4d9b.
//
// Solidity: function trade(address[] srcTokens, uint256[] srcQties, address[] destTokens, uint256[] minConversionRates) returns(address[], uint256[])
func (_KbnmultiTrade *KbnmultiTradeTransactorSession) Trade(srcTokens []common.Address, srcQties []*big.Int, destTokens []common.Address, minConversionRates []*big.Int) (*types.Transaction, error) {
	return _KbnmultiTrade.Contract.Trade(&_KbnmultiTrade.TransactOpts, srcTokens, srcQties, destTokens, minConversionRates)
}
