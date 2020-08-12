// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package reentrancer

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

// ReentrancerABI is the input ABI used to generate the binding from.
const ReentrancerABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"called\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fallbackGiveAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fallbackGiveTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fallbackGiveToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fallbackState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"fallbackUnlockAssets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"locker\",\"type\":\"address\"}],\"name\":\"pauseLocker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fallbackState\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startCall\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"locker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"giveTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"giveToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"giveAmount\",\"type\":\"uint256\"}],\"name\":\"reGive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fallbackState\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startCall\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"locker\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"unlockAssets\",\"type\":\"address[]\"}],\"name\":\"reUnlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"giveTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"giveToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"giveAmount\",\"type\":\"uint256\"}],\"name\":\"setFallbackGiveData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"unlockAssets\",\"type\":\"address[]\"}],\"name\":\"setFallbackUnlockData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"locker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nextLocker\",\"type\":\"address\"}],\"name\":\"setNextLocker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ReentrancerFuncSigs maps the 4-byte function signature to its string representation.
var ReentrancerFuncSigs = map[string]string{
	"50f9b6cd": "called()",
	"d346b9bd": "fallbackGiveAmount()",
	"6878e2e9": "fallbackGiveTo()",
	"9a78519f": "fallbackGiveToken()",
	"3bc6be0e": "fallbackState()",
	"d9f1f12b": "fallbackUnlockAssets(uint256)",
	"9d9682e5": "pauseLocker(address)",
	"7d7ad598": "reGive(uint256,uint256,address,address,address,uint256)",
	"402343fe": "reUnlock(uint256,uint256,address,address[])",
	"17ece175": "setFallbackGiveData(address,address,uint256)",
	"5728b5ef": "setFallbackUnlockData(address[])",
	"012043b8": "setNextLocker(address,address)",
}

// ReentrancerBin is the compiled bytecode used for deploying new contracts.
var ReentrancerBin = "0x608060405234801561001057600080fd5b506000600555610ca6806100256000396000f3fe6080604052600436106100a75760003560e01c80636878e2e9116100645780636878e2e9146105495780637d7ad5981461057a5780639a78519f146105cf5780639d9682e5146105e4578063d346b9bd14610617578063d9f1f12b1461062c576100a7565b8063012043b81461031357806317ece1751461034e5780633bc6be0e14610391578063402343fe146103b857806350f9b6cd146104845780635728b5ef14610499575b6004546101f657600554156100bb57610311565b606060006040516024018080602001828103825283818154815260200191508054801561011157602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116100f3575b505060408051601f198184030181529181526020820180516001600160e01b031663d791de6360e01b17815260016005559051825192965060009550339450869350918291908083835b6020831061017a5780518252601f19909201916020918201910161015b565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d80600081146101dc576040519150601f19603f3d011682016040523d82523d6000602084013e6101e1565b606091505b50509050806101ef57600080fd5b5050610311565b6005541561020357610311565b60018054600254600354604080516001600160a01b03948516602482015292909316604483015260648083019190915282518083039091018152608490910182526020810180516001600160e01b0316633b22ca0f60e21b1781526005939093559051815191926000923392859290918291908083835b602083106102995780518252601f19909201916020918201910161027a565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d80600081146102fb576040519150601f19603f3d011682016040523d82523d6000602084013e610300565b606091505b505090508061030e57600080fd5b50505b005b34801561031f57600080fd5b506103116004803603604081101561033657600080fd5b506001600160a01b0381358116916020013516610656565b34801561035a57600080fd5b506103116004803603606081101561037157600080fd5b506001600160a01b03813581169160208101359091169060400135610749565b34801561039d57600080fd5b506103a661077e565b60408051918252519081900360200190f35b3480156103c457600080fd5b50610311600480360360808110156103db57600080fd5b8135916020810135916001600160a01b03604083013516919081019060808101606082013564010000000081111561041257600080fd5b82018360208201111561042457600080fd5b8035906020019184602083028401116401000000008311171561044657600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610784945050505050565b34801561049057600080fd5b506103a661094c565b3480156104a557600080fd5b50610311600480360360208110156104bc57600080fd5b8101906020810181356401000000008111156104d757600080fd5b8201836020820111156104e957600080fd5b8035906020019184602083028401116401000000008311171561050b57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610952945050505050565b34801561055557600080fd5b5061055e610969565b604080516001600160a01b039092168252519081900360200190f35b34801561058657600080fd5b50610311600480360360c081101561059d57600080fd5b508035906020810135906001600160a01b03604082013581169160608101358216916080820135169060a00135610978565b3480156105db57600080fd5b5061055e610ad0565b3480156105f057600080fd5b506103116004803603602081101561060757600080fd5b50356001600160a01b0316610adf565b34801561062357600080fd5b506103a6610bbf565b34801561063857600080fd5b5061055e6004803603602081101561064f57600080fd5b5035610bc5565b604080516001600160a01b0380841660248084019190915283518084039091018152604490920183526020820180516001600160e01b031663277cd36160e01b1781529251825192936000939287169285929182918083835b602083106106ce5780518252601f1990920191602091820191016106af565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610730576040519150601f19603f3d011682016040523d82523d6000602084013e610735565b606091505b505090508061074357600080fd5b50505050565b600180546001600160a01b039485166001600160a01b0319918216179091556002805493909416921691909117909155600355565b60045481565b6004849055600583905560405160206024820181815283516044840152835160609385938392606490920191818601910280838360005b838110156107d35781810151838201526020016107bb565b50506040805193909501838103601f190184528552505060208101805163d791de6360e01b6001600160e01b039091161781529251815191975060009650606095506001600160a01b038a169450879390925082918083835b6020831061084b5780518252601f19909201916020918201910161082c565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d80600081146108ad576040519150601f19603f3d011682016040523d82523d6000602084013e6108b2565b606091505b50915091508181906109425760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156109075781810151838201526020016108ef565b50505050905090810190601f1680156109345780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5050505050505050565b60055481565b8051610965906000906020840190610bec565b5050565b6001546001600160a01b031681565b60048690556005859055604080516001600160a01b0380861660248301528085166044830152606480830185905283518084039091018152608490920183526020820180516001600160e01b0316633b22ca0f60e21b1781529251825192936000936060938a1692869290918291908083835b60208310610a0a5780518252601f1990920191602091820191016109eb565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610a6c576040519150601f19603f3d011682016040523d82523d6000602084013e610a71565b606091505b5091509150818190610ac45760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156109075781810151838201526020016108ef565b50505050505050505050565b6002546001600160a01b031681565b60408051600481526024810182526020810180516001600160e01b0316638456cb5960e01b1781529151815191926000926001600160a01b03861692859290918291908083835b60208310610b455780518252601f199092019160209182019101610b26565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610ba7576040519150601f19603f3d011682016040523d82523d6000602084013e610bac565b606091505b5050905080610bba57600080fd5b505050565b60035481565b60008181548110610bd257fe5b6000918252602090912001546001600160a01b0316905081565b828054828255906000526020600020908101928215610c41579160200282015b82811115610c4157825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190610c0c565b50610c4d929150610c51565b5090565b5b80821115610c4d5780546001600160a01b0319168155600101610c5256fea2646970667358221220281a1c79c7ccb15ec954d3f2b878e54d0da4d75e690df99eea18ec60457a509a64736f6c634300060c0033"

// DeployReentrancer deploys a new Ethereum contract, binding an instance of Reentrancer to it.
func DeployReentrancer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Reentrancer, error) {
	parsed, err := abi.JSON(strings.NewReader(ReentrancerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ReentrancerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Reentrancer{ReentrancerCaller: ReentrancerCaller{contract: contract}, ReentrancerTransactor: ReentrancerTransactor{contract: contract}, ReentrancerFilterer: ReentrancerFilterer{contract: contract}}, nil
}

// Reentrancer is an auto generated Go binding around an Ethereum contract.
type Reentrancer struct {
	ReentrancerCaller     // Read-only binding to the contract
	ReentrancerTransactor // Write-only binding to the contract
	ReentrancerFilterer   // Log filterer for contract events
}

// ReentrancerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReentrancerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReentrancerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReentrancerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReentrancerSession struct {
	Contract     *Reentrancer      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReentrancerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReentrancerCallerSession struct {
	Contract *ReentrancerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ReentrancerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReentrancerTransactorSession struct {
	Contract     *ReentrancerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ReentrancerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReentrancerRaw struct {
	Contract *Reentrancer // Generic contract binding to access the raw methods on
}

// ReentrancerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReentrancerCallerRaw struct {
	Contract *ReentrancerCaller // Generic read-only contract binding to access the raw methods on
}

// ReentrancerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReentrancerTransactorRaw struct {
	Contract *ReentrancerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReentrancer creates a new instance of Reentrancer, bound to a specific deployed contract.
func NewReentrancer(address common.Address, backend bind.ContractBackend) (*Reentrancer, error) {
	contract, err := bindReentrancer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Reentrancer{ReentrancerCaller: ReentrancerCaller{contract: contract}, ReentrancerTransactor: ReentrancerTransactor{contract: contract}, ReentrancerFilterer: ReentrancerFilterer{contract: contract}}, nil
}

// NewReentrancerCaller creates a new read-only instance of Reentrancer, bound to a specific deployed contract.
func NewReentrancerCaller(address common.Address, caller bind.ContractCaller) (*ReentrancerCaller, error) {
	contract, err := bindReentrancer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancerCaller{contract: contract}, nil
}

// NewReentrancerTransactor creates a new write-only instance of Reentrancer, bound to a specific deployed contract.
func NewReentrancerTransactor(address common.Address, transactor bind.ContractTransactor) (*ReentrancerTransactor, error) {
	contract, err := bindReentrancer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancerTransactor{contract: contract}, nil
}

// NewReentrancerFilterer creates a new log filterer instance of Reentrancer, bound to a specific deployed contract.
func NewReentrancerFilterer(address common.Address, filterer bind.ContractFilterer) (*ReentrancerFilterer, error) {
	contract, err := bindReentrancer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReentrancerFilterer{contract: contract}, nil
}

// bindReentrancer binds a generic wrapper to an already deployed contract.
func bindReentrancer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReentrancerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reentrancer *ReentrancerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Reentrancer.Contract.ReentrancerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reentrancer *ReentrancerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reentrancer.Contract.ReentrancerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reentrancer *ReentrancerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reentrancer.Contract.ReentrancerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reentrancer *ReentrancerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Reentrancer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reentrancer *ReentrancerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reentrancer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reentrancer *ReentrancerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reentrancer.Contract.contract.Transact(opts, method, params...)
}

// Called is a free data retrieval call binding the contract method 0x50f9b6cd.
//
// Solidity: function called() constant returns(uint256)
func (_Reentrancer *ReentrancerCaller) Called(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Reentrancer.contract.Call(opts, out, "called")
	return *ret0, err
}

// Called is a free data retrieval call binding the contract method 0x50f9b6cd.
//
// Solidity: function called() constant returns(uint256)
func (_Reentrancer *ReentrancerSession) Called() (*big.Int, error) {
	return _Reentrancer.Contract.Called(&_Reentrancer.CallOpts)
}

// Called is a free data retrieval call binding the contract method 0x50f9b6cd.
//
// Solidity: function called() constant returns(uint256)
func (_Reentrancer *ReentrancerCallerSession) Called() (*big.Int, error) {
	return _Reentrancer.Contract.Called(&_Reentrancer.CallOpts)
}

// FallbackGiveAmount is a free data retrieval call binding the contract method 0xd346b9bd.
//
// Solidity: function fallbackGiveAmount() constant returns(uint256)
func (_Reentrancer *ReentrancerCaller) FallbackGiveAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Reentrancer.contract.Call(opts, out, "fallbackGiveAmount")
	return *ret0, err
}

// FallbackGiveAmount is a free data retrieval call binding the contract method 0xd346b9bd.
//
// Solidity: function fallbackGiveAmount() constant returns(uint256)
func (_Reentrancer *ReentrancerSession) FallbackGiveAmount() (*big.Int, error) {
	return _Reentrancer.Contract.FallbackGiveAmount(&_Reentrancer.CallOpts)
}

// FallbackGiveAmount is a free data retrieval call binding the contract method 0xd346b9bd.
//
// Solidity: function fallbackGiveAmount() constant returns(uint256)
func (_Reentrancer *ReentrancerCallerSession) FallbackGiveAmount() (*big.Int, error) {
	return _Reentrancer.Contract.FallbackGiveAmount(&_Reentrancer.CallOpts)
}

// FallbackGiveTo is a free data retrieval call binding the contract method 0x6878e2e9.
//
// Solidity: function fallbackGiveTo() constant returns(address)
func (_Reentrancer *ReentrancerCaller) FallbackGiveTo(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Reentrancer.contract.Call(opts, out, "fallbackGiveTo")
	return *ret0, err
}

// FallbackGiveTo is a free data retrieval call binding the contract method 0x6878e2e9.
//
// Solidity: function fallbackGiveTo() constant returns(address)
func (_Reentrancer *ReentrancerSession) FallbackGiveTo() (common.Address, error) {
	return _Reentrancer.Contract.FallbackGiveTo(&_Reentrancer.CallOpts)
}

// FallbackGiveTo is a free data retrieval call binding the contract method 0x6878e2e9.
//
// Solidity: function fallbackGiveTo() constant returns(address)
func (_Reentrancer *ReentrancerCallerSession) FallbackGiveTo() (common.Address, error) {
	return _Reentrancer.Contract.FallbackGiveTo(&_Reentrancer.CallOpts)
}

// FallbackGiveToken is a free data retrieval call binding the contract method 0x9a78519f.
//
// Solidity: function fallbackGiveToken() constant returns(address)
func (_Reentrancer *ReentrancerCaller) FallbackGiveToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Reentrancer.contract.Call(opts, out, "fallbackGiveToken")
	return *ret0, err
}

// FallbackGiveToken is a free data retrieval call binding the contract method 0x9a78519f.
//
// Solidity: function fallbackGiveToken() constant returns(address)
func (_Reentrancer *ReentrancerSession) FallbackGiveToken() (common.Address, error) {
	return _Reentrancer.Contract.FallbackGiveToken(&_Reentrancer.CallOpts)
}

// FallbackGiveToken is a free data retrieval call binding the contract method 0x9a78519f.
//
// Solidity: function fallbackGiveToken() constant returns(address)
func (_Reentrancer *ReentrancerCallerSession) FallbackGiveToken() (common.Address, error) {
	return _Reentrancer.Contract.FallbackGiveToken(&_Reentrancer.CallOpts)
}

// FallbackState is a free data retrieval call binding the contract method 0x3bc6be0e.
//
// Solidity: function fallbackState() constant returns(uint256)
func (_Reentrancer *ReentrancerCaller) FallbackState(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Reentrancer.contract.Call(opts, out, "fallbackState")
	return *ret0, err
}

// FallbackState is a free data retrieval call binding the contract method 0x3bc6be0e.
//
// Solidity: function fallbackState() constant returns(uint256)
func (_Reentrancer *ReentrancerSession) FallbackState() (*big.Int, error) {
	return _Reentrancer.Contract.FallbackState(&_Reentrancer.CallOpts)
}

// FallbackState is a free data retrieval call binding the contract method 0x3bc6be0e.
//
// Solidity: function fallbackState() constant returns(uint256)
func (_Reentrancer *ReentrancerCallerSession) FallbackState() (*big.Int, error) {
	return _Reentrancer.Contract.FallbackState(&_Reentrancer.CallOpts)
}

// FallbackUnlockAssets is a free data retrieval call binding the contract method 0xd9f1f12b.
//
// Solidity: function fallbackUnlockAssets(uint256 ) constant returns(address)
func (_Reentrancer *ReentrancerCaller) FallbackUnlockAssets(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Reentrancer.contract.Call(opts, out, "fallbackUnlockAssets", arg0)
	return *ret0, err
}

// FallbackUnlockAssets is a free data retrieval call binding the contract method 0xd9f1f12b.
//
// Solidity: function fallbackUnlockAssets(uint256 ) constant returns(address)
func (_Reentrancer *ReentrancerSession) FallbackUnlockAssets(arg0 *big.Int) (common.Address, error) {
	return _Reentrancer.Contract.FallbackUnlockAssets(&_Reentrancer.CallOpts, arg0)
}

// FallbackUnlockAssets is a free data retrieval call binding the contract method 0xd9f1f12b.
//
// Solidity: function fallbackUnlockAssets(uint256 ) constant returns(address)
func (_Reentrancer *ReentrancerCallerSession) FallbackUnlockAssets(arg0 *big.Int) (common.Address, error) {
	return _Reentrancer.Contract.FallbackUnlockAssets(&_Reentrancer.CallOpts, arg0)
}

// PauseLocker is a paid mutator transaction binding the contract method 0x9d9682e5.
//
// Solidity: function pauseLocker(address locker) returns()
func (_Reentrancer *ReentrancerTransactor) PauseLocker(opts *bind.TransactOpts, locker common.Address) (*types.Transaction, error) {
	return _Reentrancer.contract.Transact(opts, "pauseLocker", locker)
}

// PauseLocker is a paid mutator transaction binding the contract method 0x9d9682e5.
//
// Solidity: function pauseLocker(address locker) returns()
func (_Reentrancer *ReentrancerSession) PauseLocker(locker common.Address) (*types.Transaction, error) {
	return _Reentrancer.Contract.PauseLocker(&_Reentrancer.TransactOpts, locker)
}

// PauseLocker is a paid mutator transaction binding the contract method 0x9d9682e5.
//
// Solidity: function pauseLocker(address locker) returns()
func (_Reentrancer *ReentrancerTransactorSession) PauseLocker(locker common.Address) (*types.Transaction, error) {
	return _Reentrancer.Contract.PauseLocker(&_Reentrancer.TransactOpts, locker)
}

// ReGive is a paid mutator transaction binding the contract method 0x7d7ad598.
//
// Solidity: function reGive(uint256 _fallbackState, uint256 _startCall, address locker, address giveTo, address giveToken, uint256 giveAmount) returns()
func (_Reentrancer *ReentrancerTransactor) ReGive(opts *bind.TransactOpts, _fallbackState *big.Int, _startCall *big.Int, locker common.Address, giveTo common.Address, giveToken common.Address, giveAmount *big.Int) (*types.Transaction, error) {
	return _Reentrancer.contract.Transact(opts, "reGive", _fallbackState, _startCall, locker, giveTo, giveToken, giveAmount)
}

// ReGive is a paid mutator transaction binding the contract method 0x7d7ad598.
//
// Solidity: function reGive(uint256 _fallbackState, uint256 _startCall, address locker, address giveTo, address giveToken, uint256 giveAmount) returns()
func (_Reentrancer *ReentrancerSession) ReGive(_fallbackState *big.Int, _startCall *big.Int, locker common.Address, giveTo common.Address, giveToken common.Address, giveAmount *big.Int) (*types.Transaction, error) {
	return _Reentrancer.Contract.ReGive(&_Reentrancer.TransactOpts, _fallbackState, _startCall, locker, giveTo, giveToken, giveAmount)
}

// ReGive is a paid mutator transaction binding the contract method 0x7d7ad598.
//
// Solidity: function reGive(uint256 _fallbackState, uint256 _startCall, address locker, address giveTo, address giveToken, uint256 giveAmount) returns()
func (_Reentrancer *ReentrancerTransactorSession) ReGive(_fallbackState *big.Int, _startCall *big.Int, locker common.Address, giveTo common.Address, giveToken common.Address, giveAmount *big.Int) (*types.Transaction, error) {
	return _Reentrancer.Contract.ReGive(&_Reentrancer.TransactOpts, _fallbackState, _startCall, locker, giveTo, giveToken, giveAmount)
}

// ReUnlock is a paid mutator transaction binding the contract method 0x402343fe.
//
// Solidity: function reUnlock(uint256 _fallbackState, uint256 _startCall, address locker, address[] unlockAssets) returns()
func (_Reentrancer *ReentrancerTransactor) ReUnlock(opts *bind.TransactOpts, _fallbackState *big.Int, _startCall *big.Int, locker common.Address, unlockAssets []common.Address) (*types.Transaction, error) {
	return _Reentrancer.contract.Transact(opts, "reUnlock", _fallbackState, _startCall, locker, unlockAssets)
}

// ReUnlock is a paid mutator transaction binding the contract method 0x402343fe.
//
// Solidity: function reUnlock(uint256 _fallbackState, uint256 _startCall, address locker, address[] unlockAssets) returns()
func (_Reentrancer *ReentrancerSession) ReUnlock(_fallbackState *big.Int, _startCall *big.Int, locker common.Address, unlockAssets []common.Address) (*types.Transaction, error) {
	return _Reentrancer.Contract.ReUnlock(&_Reentrancer.TransactOpts, _fallbackState, _startCall, locker, unlockAssets)
}

// ReUnlock is a paid mutator transaction binding the contract method 0x402343fe.
//
// Solidity: function reUnlock(uint256 _fallbackState, uint256 _startCall, address locker, address[] unlockAssets) returns()
func (_Reentrancer *ReentrancerTransactorSession) ReUnlock(_fallbackState *big.Int, _startCall *big.Int, locker common.Address, unlockAssets []common.Address) (*types.Transaction, error) {
	return _Reentrancer.Contract.ReUnlock(&_Reentrancer.TransactOpts, _fallbackState, _startCall, locker, unlockAssets)
}

// SetFallbackGiveData is a paid mutator transaction binding the contract method 0x17ece175.
//
// Solidity: function setFallbackGiveData(address giveTo, address giveToken, uint256 giveAmount) returns()
func (_Reentrancer *ReentrancerTransactor) SetFallbackGiveData(opts *bind.TransactOpts, giveTo common.Address, giveToken common.Address, giveAmount *big.Int) (*types.Transaction, error) {
	return _Reentrancer.contract.Transact(opts, "setFallbackGiveData", giveTo, giveToken, giveAmount)
}

// SetFallbackGiveData is a paid mutator transaction binding the contract method 0x17ece175.
//
// Solidity: function setFallbackGiveData(address giveTo, address giveToken, uint256 giveAmount) returns()
func (_Reentrancer *ReentrancerSession) SetFallbackGiveData(giveTo common.Address, giveToken common.Address, giveAmount *big.Int) (*types.Transaction, error) {
	return _Reentrancer.Contract.SetFallbackGiveData(&_Reentrancer.TransactOpts, giveTo, giveToken, giveAmount)
}

// SetFallbackGiveData is a paid mutator transaction binding the contract method 0x17ece175.
//
// Solidity: function setFallbackGiveData(address giveTo, address giveToken, uint256 giveAmount) returns()
func (_Reentrancer *ReentrancerTransactorSession) SetFallbackGiveData(giveTo common.Address, giveToken common.Address, giveAmount *big.Int) (*types.Transaction, error) {
	return _Reentrancer.Contract.SetFallbackGiveData(&_Reentrancer.TransactOpts, giveTo, giveToken, giveAmount)
}

// SetFallbackUnlockData is a paid mutator transaction binding the contract method 0x5728b5ef.
//
// Solidity: function setFallbackUnlockData(address[] unlockAssets) returns()
func (_Reentrancer *ReentrancerTransactor) SetFallbackUnlockData(opts *bind.TransactOpts, unlockAssets []common.Address) (*types.Transaction, error) {
	return _Reentrancer.contract.Transact(opts, "setFallbackUnlockData", unlockAssets)
}

// SetFallbackUnlockData is a paid mutator transaction binding the contract method 0x5728b5ef.
//
// Solidity: function setFallbackUnlockData(address[] unlockAssets) returns()
func (_Reentrancer *ReentrancerSession) SetFallbackUnlockData(unlockAssets []common.Address) (*types.Transaction, error) {
	return _Reentrancer.Contract.SetFallbackUnlockData(&_Reentrancer.TransactOpts, unlockAssets)
}

// SetFallbackUnlockData is a paid mutator transaction binding the contract method 0x5728b5ef.
//
// Solidity: function setFallbackUnlockData(address[] unlockAssets) returns()
func (_Reentrancer *ReentrancerTransactorSession) SetFallbackUnlockData(unlockAssets []common.Address) (*types.Transaction, error) {
	return _Reentrancer.Contract.SetFallbackUnlockData(&_Reentrancer.TransactOpts, unlockAssets)
}

// SetNextLocker is a paid mutator transaction binding the contract method 0x012043b8.
//
// Solidity: function setNextLocker(address locker, address nextLocker) returns()
func (_Reentrancer *ReentrancerTransactor) SetNextLocker(opts *bind.TransactOpts, locker common.Address, nextLocker common.Address) (*types.Transaction, error) {
	return _Reentrancer.contract.Transact(opts, "setNextLocker", locker, nextLocker)
}

// SetNextLocker is a paid mutator transaction binding the contract method 0x012043b8.
//
// Solidity: function setNextLocker(address locker, address nextLocker) returns()
func (_Reentrancer *ReentrancerSession) SetNextLocker(locker common.Address, nextLocker common.Address) (*types.Transaction, error) {
	return _Reentrancer.Contract.SetNextLocker(&_Reentrancer.TransactOpts, locker, nextLocker)
}

// SetNextLocker is a paid mutator transaction binding the contract method 0x012043b8.
//
// Solidity: function setNextLocker(address locker, address nextLocker) returns()
func (_Reentrancer *ReentrancerTransactorSession) SetNextLocker(locker common.Address, nextLocker common.Address) (*types.Transaction, error) {
	return _Reentrancer.Contract.SetNextLocker(&_Reentrancer.TransactOpts, locker, nextLocker)
}
