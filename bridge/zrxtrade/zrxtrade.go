// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zrxtrade

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

// ZrxtradeABI is the input ABI used to generate the binding from.
const ZrxtradeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wETH\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_zeroProxy\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_incognitoSmartContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ETH_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incognitoSmartContract\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callDataHex\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_forwarder\",\"type\":\"address\"}],\"name\":\"trade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawWrapETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// ZrxtradeBin is the compiled bytecode used for deploying new contracts.
var ZrxtradeBin = "0x608060405234801561001057600080fd5b50604051610b64380380610b648339818101604052606081101561003357600080fd5b8101908080519060200190929190805190602001909291908051906020019092919050505081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050610a388061012c6000396000f3fe6080604052600436106100435760003560e01c80630e0569a61461004f57806372e94bf61461008a57806398c36152146100e1578063b42a644b1461024d5761004a565b3661004a57005b600080fd5b34801561005b57600080fd5b506100886004803603602081101561007257600080fd5b81019080803590602001909291905050506102a4565b005b34801561009657600080fd5b5061009f610334565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610204600480360360a08110156100f757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019064010000000081111561015e57600080fd5b82018360208201111561017057600080fd5b8035906020019184600183028401116401000000008311171561019257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610339565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b34801561025957600080fd5b50610262610508565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632e1a7d4d826040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b15801561031957600080fd5b505af115801561032d573d6000803e3d6000fd5b5050505050565b600081565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461039557600080fd5b6103c287600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168861052d565b60008373ffffffffffffffffffffffffffffffffffffffff1634866040518082805190602001908083835b6020831061041057805182526020820191506020810190506020830392506103ed565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114610472576040519150601f19603f3d011682016040523d82523d6000602084013e610477565b606091505b505090508061048557600080fd5b60008073ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff1614156104e0576104c560006106c8565b90506104d0816102a4565b6104db600082610741565b6104f6565b6104e9876106c8565b90506104f58782610741565b5b86819350935050509550959350505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16146106c3578273ffffffffffffffffffffffffffffffffffffffff1663095ea7b38360006040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b1580156105e957600080fd5b505af11580156105fd573d6000803e3d6000fd5b505050506106096108c6565b61061257600080fd5b8273ffffffffffffffffffffffffffffffffffffffff1663095ea7b383836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b15801561069957600080fd5b505af11580156106ad573d6000803e3d6000fd5b505050506106b96108c6565b6106c257600080fd5b5b505050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561073057610729600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16610904565b905061073c565b61073982610904565b90505b919050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156107f0578047101561078357600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156107ea573d6000803e3d6000fd5b506108c2565b8173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b15801561089857600080fd5b505af11580156108ac573d6000803e3d6000fd5b505050506108b86108c6565b6108c157600080fd5b5b5050565b600080600090503d600081146108e357602081146108ec576108f8565b600191506108f8565b60206000803e60005191505b50600081141591505090565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610942574790506109fd565b8173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b1580156109bf57600080fd5b505afa1580156109d3573d6000803e3d6000fd5b505050506040513d60208110156109e957600080fd5b810190808051906020019092919050505090505b91905056fea26469706673582212200c039363bbc7231ba8f4837df0e406d8150639bad07bd30ce5881ebc3264660964736f6c63430006060033"

// DeployZrxtrade deploys a new Ethereum contract, binding an instance of Zrxtrade to it.
func DeployZrxtrade(auth *bind.TransactOpts, backend bind.ContractBackend, _wETH common.Address, _zeroProxy common.Address, _incognitoSmartContract common.Address) (common.Address, *types.Transaction, *Zrxtrade, error) {
	parsed, err := abi.JSON(strings.NewReader(ZrxtradeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZrxtradeBin), backend, _wETH, _zeroProxy, _incognitoSmartContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Zrxtrade{ZrxtradeCaller: ZrxtradeCaller{contract: contract}, ZrxtradeTransactor: ZrxtradeTransactor{contract: contract}, ZrxtradeFilterer: ZrxtradeFilterer{contract: contract}}, nil
}

// Zrxtrade is an auto generated Go binding around an Ethereum contract.
type Zrxtrade struct {
	ZrxtradeCaller     // Read-only binding to the contract
	ZrxtradeTransactor // Write-only binding to the contract
	ZrxtradeFilterer   // Log filterer for contract events
}

// ZrxtradeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZrxtradeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZrxtradeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZrxtradeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZrxtradeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZrxtradeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZrxtradeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZrxtradeSession struct {
	Contract     *Zrxtrade         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZrxtradeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZrxtradeCallerSession struct {
	Contract *ZrxtradeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ZrxtradeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZrxtradeTransactorSession struct {
	Contract     *ZrxtradeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ZrxtradeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZrxtradeRaw struct {
	Contract *Zrxtrade // Generic contract binding to access the raw methods on
}

// ZrxtradeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZrxtradeCallerRaw struct {
	Contract *ZrxtradeCaller // Generic read-only contract binding to access the raw methods on
}

// ZrxtradeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZrxtradeTransactorRaw struct {
	Contract *ZrxtradeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZrxtrade creates a new instance of Zrxtrade, bound to a specific deployed contract.
func NewZrxtrade(address common.Address, backend bind.ContractBackend) (*Zrxtrade, error) {
	contract, err := bindZrxtrade(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Zrxtrade{ZrxtradeCaller: ZrxtradeCaller{contract: contract}, ZrxtradeTransactor: ZrxtradeTransactor{contract: contract}, ZrxtradeFilterer: ZrxtradeFilterer{contract: contract}}, nil
}

// NewZrxtradeCaller creates a new read-only instance of Zrxtrade, bound to a specific deployed contract.
func NewZrxtradeCaller(address common.Address, caller bind.ContractCaller) (*ZrxtradeCaller, error) {
	contract, err := bindZrxtrade(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZrxtradeCaller{contract: contract}, nil
}

// NewZrxtradeTransactor creates a new write-only instance of Zrxtrade, bound to a specific deployed contract.
func NewZrxtradeTransactor(address common.Address, transactor bind.ContractTransactor) (*ZrxtradeTransactor, error) {
	contract, err := bindZrxtrade(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZrxtradeTransactor{contract: contract}, nil
}

// NewZrxtradeFilterer creates a new log filterer instance of Zrxtrade, bound to a specific deployed contract.
func NewZrxtradeFilterer(address common.Address, filterer bind.ContractFilterer) (*ZrxtradeFilterer, error) {
	contract, err := bindZrxtrade(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZrxtradeFilterer{contract: contract}, nil
}

// bindZrxtrade binds a generic wrapper to an already deployed contract.
func bindZrxtrade(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZrxtradeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zrxtrade *ZrxtradeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Zrxtrade.Contract.ZrxtradeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zrxtrade *ZrxtradeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zrxtrade.Contract.ZrxtradeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zrxtrade *ZrxtradeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zrxtrade.Contract.ZrxtradeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zrxtrade *ZrxtradeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Zrxtrade.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zrxtrade *ZrxtradeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zrxtrade.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zrxtrade *ZrxtradeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zrxtrade.Contract.contract.Transact(opts, method, params...)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_Zrxtrade *ZrxtradeCaller) ETHCONTRACTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Zrxtrade.contract.Call(opts, out, "ETH_CONTRACT_ADDRESS")
	return *ret0, err
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_Zrxtrade *ZrxtradeSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _Zrxtrade.Contract.ETHCONTRACTADDRESS(&_Zrxtrade.CallOpts)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_Zrxtrade *ZrxtradeCallerSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _Zrxtrade.Contract.ETHCONTRACTADDRESS(&_Zrxtrade.CallOpts)
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_Zrxtrade *ZrxtradeCaller) IncognitoSmartContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Zrxtrade.contract.Call(opts, out, "incognitoSmartContract")
	return *ret0, err
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_Zrxtrade *ZrxtradeSession) IncognitoSmartContract() (common.Address, error) {
	return _Zrxtrade.Contract.IncognitoSmartContract(&_Zrxtrade.CallOpts)
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_Zrxtrade *ZrxtradeCallerSession) IncognitoSmartContract() (common.Address, error) {
	return _Zrxtrade.Contract.IncognitoSmartContract(&_Zrxtrade.CallOpts)
}

// Trade is a paid mutator transaction binding the contract method 0x98c36152.
//
// Solidity: function trade(address srcToken, uint256 amount, address destToken, bytes callDataHex, address _forwarder) returns(address, uint256)
func (_Zrxtrade *ZrxtradeTransactor) Trade(opts *bind.TransactOpts, srcToken common.Address, amount *big.Int, destToken common.Address, callDataHex []byte, _forwarder common.Address) (*types.Transaction, error) {
	return _Zrxtrade.contract.Transact(opts, "trade", srcToken, amount, destToken, callDataHex, _forwarder)
}

// Trade is a paid mutator transaction binding the contract method 0x98c36152.
//
// Solidity: function trade(address srcToken, uint256 amount, address destToken, bytes callDataHex, address _forwarder) returns(address, uint256)
func (_Zrxtrade *ZrxtradeSession) Trade(srcToken common.Address, amount *big.Int, destToken common.Address, callDataHex []byte, _forwarder common.Address) (*types.Transaction, error) {
	return _Zrxtrade.Contract.Trade(&_Zrxtrade.TransactOpts, srcToken, amount, destToken, callDataHex, _forwarder)
}

// Trade is a paid mutator transaction binding the contract method 0x98c36152.
//
// Solidity: function trade(address srcToken, uint256 amount, address destToken, bytes callDataHex, address _forwarder) returns(address, uint256)
func (_Zrxtrade *ZrxtradeTransactorSession) Trade(srcToken common.Address, amount *big.Int, destToken common.Address, callDataHex []byte, _forwarder common.Address) (*types.Transaction, error) {
	return _Zrxtrade.Contract.Trade(&_Zrxtrade.TransactOpts, srcToken, amount, destToken, callDataHex, _forwarder)
}

// WithdrawWrapETH is a paid mutator transaction binding the contract method 0x0e0569a6.
//
// Solidity: function withdrawWrapETH(uint256 amount) returns()
func (_Zrxtrade *ZrxtradeTransactor) WithdrawWrapETH(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Zrxtrade.contract.Transact(opts, "withdrawWrapETH", amount)
}

// WithdrawWrapETH is a paid mutator transaction binding the contract method 0x0e0569a6.
//
// Solidity: function withdrawWrapETH(uint256 amount) returns()
func (_Zrxtrade *ZrxtradeSession) WithdrawWrapETH(amount *big.Int) (*types.Transaction, error) {
	return _Zrxtrade.Contract.WithdrawWrapETH(&_Zrxtrade.TransactOpts, amount)
}

// WithdrawWrapETH is a paid mutator transaction binding the contract method 0x0e0569a6.
//
// Solidity: function withdrawWrapETH(uint256 amount) returns()
func (_Zrxtrade *ZrxtradeTransactorSession) WithdrawWrapETH(amount *big.Int) (*types.Transaction, error) {
	return _Zrxtrade.Contract.WithdrawWrapETH(&_Zrxtrade.TransactOpts, amount)
}
