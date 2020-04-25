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
const KbnmultiTradeABI = "[{\"inputs\":[{\"internalType\":\"contractKyberNetwork\",\"name\":\"_kyberNetworkProxyContract\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_incognitoSmartContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"ETH_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcQty\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"destToken\",\"type\":\"address\"}],\"name\":\"getConversionRates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"incognitoSmartContract\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kyberNetworkProxyContract\",\"outputs\":[{\"internalType\":\"contractKyberNetwork\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"srcTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"srcQties\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"destTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"minConversionRates\",\"type\":\"uint256[]\"}],\"name\":\"trade\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// KbnmultiTradeBin is the compiled bytecode used for deploying new contracts.
var KbnmultiTradeBin = "0x608060405234801561001057600080fd5b506040516112dd3803806112dd8339818101604052604081101561003357600080fd5b81019080805190602001909291908051906020019092919050505081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050506111fd806100e06000396000f3fe60806040526004361061004a5760003560e01c80630aea81881461004c578063398b4d9b146100e257806372e94bf6146103f3578063785250da1461044a578063b42a644b146104a1575b005b34801561005857600080fd5b506100c56004803603606081101561006f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506104f8565b604051808381526020018281526020019250505060405180910390f35b610354600480360360808110156100f857600080fd5b810190808035906020019064010000000081111561011557600080fd5b82018360208201111561012757600080fd5b8035906020019184602083028401116401000000008311171561014957600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050509192919290803590602001906401000000008111156101a957600080fd5b8201836020820111156101bb57600080fd5b803590602001918460208302840111640100000000831117156101dd57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561023d57600080fd5b82018360208201111561024f57600080fd5b8035906020019184602083028401116401000000008311171561027157600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050509192919290803590602001906401000000008111156102d157600080fd5b8201836020820111156102e357600080fd5b8035906020019184602083028401116401000000008311171561030557600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050509192919290505050610626565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b8381101561039b578082015181840152602081019050610380565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156103dd5780820151818401526020810190506103c2565b5050505090500194505050505060405180910390f35b3480156103ff57600080fd5b506104086109fc565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561045657600080fd5b5061045f610a01565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156104ad57600080fd5b506104b6610a27565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b600080600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663809a9e558685876040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050604080518083038186803b1580156105d557600080fd5b505afa1580156105e9573d6000803e3d6000fd5b505050506040513d60408110156105ff57600080fd5b81019080805190602001909291908051906020019092919050505091509150935093915050565b6060806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461068257600080fd5b84518651148015610694575085518451145b61069d57600080fd5b82518451146106ab57600080fd5b606084516040519080825280602002602001820160405280156106dd5781602001602082028038833980820191505090505b50905060008090505b87518110156109eb578681815181106106fb57fe5b602002602001015161071f89838151811061071257fe5b6020026020010151610a4c565b101561072a57600080fd5b85818151811061073657fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1688828151811061076057fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff16141561078957600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168882815181106107ad57fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1614610929576108238882815181106107df57fe5b6020026020010151600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1689848151811061081657fe5b6020026020010151610b4a565b600073ffffffffffffffffffffffffffffffffffffffff1686828151811061084757fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff16146108d35760006108c489838151811061087b57fe5b602002602001015189848151811061088f57fe5b60200260200101518985815181106108a357fe5b60200260200101518986815181106108b757fe5b6020026020010151610ce5565b116108ce57600080fd5b610924565b60006109198983815181106108e457fe5b60200260200101518984815181106108f857fe5b602002602001015188858151811061090c57fe5b6020026020010151610e11565b1161092357600080fd5b5b61097a565b600061096f87838151811061093a57fe5b602002602001015189848151811061094e57fe5b602002602001015188858151811061096257fe5b6020026020010151610f08565b1161097957600080fd5b5b61099686828151811061098957fe5b6020026020010151610a4c565b8282815181106109a257fe5b6020026020010181815250506109de8682815181106109bd57fe5b60200260200101518383815181106109d157fe5b6020026020010151611005565b80806001019150506106e6565b508481925092505094509492505050565b600081565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610a8a57479050610b45565b8173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015610b0757600080fd5b505afa158015610b1b573d6000803e3d6000fd5b505050506040513d6020811015610b3157600080fd5b810190808051906020019092919050505090505b919050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614610ce0578273ffffffffffffffffffffffffffffffffffffffff1663095ea7b38360006040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b158015610c0657600080fd5b505af1158015610c1a573d6000803e3d6000fd5b50505050610c2661118a565b610c2f57600080fd5b8273ffffffffffffffffffffffffffffffffffffffff1663095ea7b383836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b158015610cb657600080fd5b505af1158015610cca573d6000803e3d6000fd5b50505050610cd661118a565b610cdf57600080fd5b5b505050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637409e2eb868686866040518563ffffffff1660e01b8152600401808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018481526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001945050505050602060405180830381600087803b158015610dcc57600080fd5b505af1158015610de0573d6000803e3d6000fd5b505050506040513d6020811015610df657600080fd5b81019080805190602001909291905050509050949350505050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633bba21dc8585856040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018281526020019350505050602060405180830381600087803b158015610ec457600080fd5b505af1158015610ed8573d6000803e3d6000fd5b505050506040513d6020811015610eee57600080fd5b810190808051906020019092919050505090509392505050565b600082471015610f1757600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637a2a04568486856040518463ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001925050506020604051808303818588803b158015610fc057600080fd5b505af1158015610fd4573d6000803e3d6000fd5b50505050506040513d6020811015610feb57600080fd5b810190808051906020019092919050505090509392505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156110b4578047101561104757600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156110ae573d6000803e3d6000fd5b50611186565b8173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b15801561115c57600080fd5b505af1158015611170573d6000803e3d6000fd5b5050505061117c61118a565b61118557600080fd5b5b5050565b600080600090503d600081146111a757602081146111b0576111bc565b600191506111bc565b60206000803e60005191505b5060008114159150509056fea265627a7a72315820cd25eee25e6c0cf073a5f60557661d4a43252f501cf5fcccfb5eecf44039e11564736f6c63430005100032"

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
