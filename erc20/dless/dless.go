// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dless

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

// DLESSABI is the input ABI used to generate the binding from.
const DLESSABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Freeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Unfreeze\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"freeze\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"freezeOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"unfreeze\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DLESSFuncSigs maps the 4-byte function signature to its string representation.
var DLESSFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"42966c68": "burn(uint256)",
	"d7a78db8": "freeze(uint256)",
	"cd4217c1": "freezeOf(address)",
	"06fdde03": "name()",
	"8da5cb5b": "owner()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"6623fc46": "unfreeze(uint256)",
	"3bed33ce": "withdrawEther(uint256)",
}

// DLESSBin is the compiled bytecode used for deploying new contracts.
var DLESSBin = "0x60806040523480156200001157600080fd5b5060405162000da338038062000da3833981810160405260608110156200003757600080fd5b8151602083018051604051929492938301929190846401000000008211156200005f57600080fd5b9083019060208201858111156200007557600080fd5b82516401000000008111828201881017156200009057600080fd5b82525081516020918201929091019080838360005b83811015620000bf578181015183820152602001620000a5565b50505050905090810190601f168015620000ed5780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200011157600080fd5b9083019060208201858111156200012757600080fd5b82516401000000008111828201881017156200014257600080fd5b82525081516020918201929091019080838360005b838110156200017157818101518382015260200162000157565b50505050905090810190601f1680156200019f5780820380516001836020036101000a031916815260200191505b5060409081523360009081526004602090815291812088905560028890558651620001d39550909350908601915062000206565b508051620001e990600190602084019062000206565b5050600380546001600160a01b0319163317905550620002ab9050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200024957805160ff191683800117855562000279565b8280016001018555821562000279579182015b82811115620002795782518255916020019190600101906200025c565b50620002879291506200028b565b5090565b620002a891905b8082111562000287576000815560010162000292565b90565b610ae880620002bb6000396000f3fe6080604052600436106100dd5760003560e01c806370a082311161007f578063a9059cbb11610059578063a9059cbb14610317578063cd4217c114610350578063d7a78db814610383578063dd62ed3e146103ad576100dd565b806370a082311461029e5780638da5cb5b146102d157806395d89b4114610302576100dd565b806323b872dd116100bb57806323b872dd146101dd5780633bed33ce1461022057806342966c681461024a5780636623fc4614610274576100dd565b806306fdde03146100df578063095ea7b31461016957806318160ddd146101b6575b005b3480156100eb57600080fd5b506100f46103e8565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561012e578181015183820152602001610116565b50505050905090810190601f16801561015b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561017557600080fd5b506101a26004803603604081101561018c57600080fd5b506001600160a01b038135169060200135610476565b604080519115158252519081900360200190f35b3480156101c257600080fd5b506101cb6104b2565b60408051918252519081900360200190f35b3480156101e957600080fd5b506101a26004803603606081101561020057600080fd5b506001600160a01b038135811691602081013590911690604001356104b8565b34801561022c57600080fd5b506100dd6004803603602081101561024357600080fd5b5035610651565b34801561025657600080fd5b506101a26004803603602081101561026d57600080fd5b50356106a6565b34801561028057600080fd5b506101a26004803603602081101561029757600080fd5b5035610747565b3480156102aa57600080fd5b506101cb600480360360208110156102c157600080fd5b50356001600160a01b0316610801565b3480156102dd57600080fd5b506102e6610813565b604080516001600160a01b039092168252519081900360200190f35b34801561030e57600080fd5b506100f4610822565b34801561032357600080fd5b506101a26004803603604081101561033a57600080fd5b506001600160a01b03813516906020013561087c565b34801561035c57600080fd5b506101cb6004803603602081101561037357600080fd5b50356001600160a01b0316610985565b34801561038f57600080fd5b506101a2600480360360208110156103a657600080fd5b5035610997565b3480156103b957600080fd5b506101cb600480360360408110156103d057600080fd5b506001600160a01b0381358116916020013516610a51565b6000805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561046e5780601f106104435761010080835404028352916020019161046e565b820191906000526020600020905b81548152906001019060200180831161045157829003601f168201915b505050505081565b600080821161048457600080fd5b503360009081526006602090815260408083206001600160a01b039590951683529390529190912055600190565b60025481565b60006001600160a01b0383166104cd57600080fd5b600082116104da57600080fd5b6001600160a01b0384166000908152600460205260409020548211156104ff57600080fd5b6001600160a01b038316600090815260046020526040902054828101101561052657600080fd5b6001600160a01b038416600090815260066020908152604080832033845290915290205482111561055657600080fd5b6001600160a01b0384166000908152600460205260409020546105799083610a6e565b6001600160a01b0380861660009081526004602052604080822093909355908516815220546105a89083610a82565b6001600160a01b0380851660009081526004602090815260408083209490945591871681526006825282812033825290915220546105e69083610a6e565b6001600160a01b03808616600081815260066020908152604080832033845282529182902094909455805186815290519287169391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a35060019392505050565b6003546001600160a01b0316331461066857600080fd5b6003546040516001600160a01b039091169082156108fc029083906000818181858888f193505050501580156106a2573d6000803e3d6000fd5b5050565b336000908152600460205260408120548211156106c257600080fd5b600082116106cf57600080fd5b336000908152600460205260409020546106e99083610a6e565b336000908152600460205260409020556002546107069083610a6e565b60025560408051838152905133917fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5919081900360200190a2506001919050565b3360009081526005602052604081205482111561076357600080fd5b6000821161077057600080fd5b3360009081526005602052604090205461078a9083610a6e565b336000908152600560209081526040808320939093556004905220546107b09083610a82565b33600081815260046020908152604091829020939093558051858152905191927f2cfce4af01bcb9d6cf6c84ee1b7c491100b8695368264146a94d71e10a63083f92918290030190a2506001919050565b60046020526000908152604090205481565b6003546001600160a01b031681565b60018054604080516020600284861615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561046e5780601f106104435761010080835404028352916020019161046e565b60006001600160a01b03831661089157600080fd5b6000821161089e57600080fd5b336000908152600460205260409020548211156108ba57600080fd5b6001600160a01b03831660009081526004602052604090205482810110156108e157600080fd5b336000908152600460205260409020546108fb9083610a6e565b33600090815260046020526040808220929092556001600160a01b038516815220546109279083610a82565b6001600160a01b0384166000818152600460209081526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b60056020526000908152604090205481565b336000908152600460205260408120548211156109b357600080fd5b600082116109c057600080fd5b336000908152600460205260409020546109da9083610a6e565b33600090815260046020908152604080832093909355600590522054610a009083610a82565b33600081815260056020908152604091829020939093558051858152905191927ff97a274face0b5517365ad396b1fdba6f68bd3135ef603e44272adba3af5a1e092918290030190a2506001919050565b600660209081526000928352604080842090915290825290205481565b6000610a7c83831115610aa6565b50900390565b6000828201610a9f848210801590610a9a5750838210155b610aa6565b9392505050565b80610ab057600080fd5b5056fea265627a7a72315820a64a75b298fddb7407fb4aa168996f516a51e80727331277d9ec603d328558e064736f6c634300050c0032"

// DeployDLESS deploys a new Ethereum contract, binding an instance of DLESS to it.
func DeployDLESS(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int, tokenName string, tokenSymbol string) (common.Address, *types.Transaction, *DLESS, error) {
	parsed, err := abi.JSON(strings.NewReader(DLESSABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DLESSBin), backend, initialSupply, tokenName, tokenSymbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DLESS{DLESSCaller: DLESSCaller{contract: contract}, DLESSTransactor: DLESSTransactor{contract: contract}, DLESSFilterer: DLESSFilterer{contract: contract}}, nil
}

// DLESS is an auto generated Go binding around an Ethereum contract.
type DLESS struct {
	DLESSCaller     // Read-only binding to the contract
	DLESSTransactor // Write-only binding to the contract
	DLESSFilterer   // Log filterer for contract events
}

// DLESSCaller is an auto generated read-only Go binding around an Ethereum contract.
type DLESSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DLESSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DLESSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DLESSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DLESSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DLESSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DLESSSession struct {
	Contract     *DLESS            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DLESSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DLESSCallerSession struct {
	Contract *DLESSCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DLESSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DLESSTransactorSession struct {
	Contract     *DLESSTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DLESSRaw is an auto generated low-level Go binding around an Ethereum contract.
type DLESSRaw struct {
	Contract *DLESS // Generic contract binding to access the raw methods on
}

// DLESSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DLESSCallerRaw struct {
	Contract *DLESSCaller // Generic read-only contract binding to access the raw methods on
}

// DLESSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DLESSTransactorRaw struct {
	Contract *DLESSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDLESS creates a new instance of DLESS, bound to a specific deployed contract.
func NewDLESS(address common.Address, backend bind.ContractBackend) (*DLESS, error) {
	contract, err := bindDLESS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DLESS{DLESSCaller: DLESSCaller{contract: contract}, DLESSTransactor: DLESSTransactor{contract: contract}, DLESSFilterer: DLESSFilterer{contract: contract}}, nil
}

// NewDLESSCaller creates a new read-only instance of DLESS, bound to a specific deployed contract.
func NewDLESSCaller(address common.Address, caller bind.ContractCaller) (*DLESSCaller, error) {
	contract, err := bindDLESS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DLESSCaller{contract: contract}, nil
}

// NewDLESSTransactor creates a new write-only instance of DLESS, bound to a specific deployed contract.
func NewDLESSTransactor(address common.Address, transactor bind.ContractTransactor) (*DLESSTransactor, error) {
	contract, err := bindDLESS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DLESSTransactor{contract: contract}, nil
}

// NewDLESSFilterer creates a new log filterer instance of DLESS, bound to a specific deployed contract.
func NewDLESSFilterer(address common.Address, filterer bind.ContractFilterer) (*DLESSFilterer, error) {
	contract, err := bindDLESS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DLESSFilterer{contract: contract}, nil
}

// bindDLESS binds a generic wrapper to an already deployed contract.
func bindDLESS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DLESSABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DLESS *DLESSRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DLESS.Contract.DLESSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DLESS *DLESSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DLESS.Contract.DLESSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DLESS *DLESSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DLESS.Contract.DLESSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DLESS *DLESSCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DLESS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DLESS *DLESSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DLESS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DLESS *DLESSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DLESS.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) constant returns(uint256)
func (_DLESS *DLESSCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DLESS.contract.Call(opts, out, "allowance", arg0, arg1)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) constant returns(uint256)
func (_DLESS *DLESSSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DLESS.Contract.Allowance(&_DLESS.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) constant returns(uint256)
func (_DLESS *DLESSCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DLESS.Contract.Allowance(&_DLESS.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) constant returns(uint256)
func (_DLESS *DLESSCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DLESS.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) constant returns(uint256)
func (_DLESS *DLESSSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _DLESS.Contract.BalanceOf(&_DLESS.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) constant returns(uint256)
func (_DLESS *DLESSCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _DLESS.Contract.BalanceOf(&_DLESS.CallOpts, arg0)
}

// FreezeOf is a free data retrieval call binding the contract method 0xcd4217c1.
//
// Solidity: function freezeOf(address ) constant returns(uint256)
func (_DLESS *DLESSCaller) FreezeOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DLESS.contract.Call(opts, out, "freezeOf", arg0)
	return *ret0, err
}

// FreezeOf is a free data retrieval call binding the contract method 0xcd4217c1.
//
// Solidity: function freezeOf(address ) constant returns(uint256)
func (_DLESS *DLESSSession) FreezeOf(arg0 common.Address) (*big.Int, error) {
	return _DLESS.Contract.FreezeOf(&_DLESS.CallOpts, arg0)
}

// FreezeOf is a free data retrieval call binding the contract method 0xcd4217c1.
//
// Solidity: function freezeOf(address ) constant returns(uint256)
func (_DLESS *DLESSCallerSession) FreezeOf(arg0 common.Address) (*big.Int, error) {
	return _DLESS.Contract.FreezeOf(&_DLESS.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_DLESS *DLESSCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DLESS.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_DLESS *DLESSSession) Name() (string, error) {
	return _DLESS.Contract.Name(&_DLESS.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_DLESS *DLESSCallerSession) Name() (string, error) {
	return _DLESS.Contract.Name(&_DLESS.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DLESS *DLESSCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DLESS.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DLESS *DLESSSession) Owner() (common.Address, error) {
	return _DLESS.Contract.Owner(&_DLESS.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DLESS *DLESSCallerSession) Owner() (common.Address, error) {
	return _DLESS.Contract.Owner(&_DLESS.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_DLESS *DLESSCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DLESS.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_DLESS *DLESSSession) Symbol() (string, error) {
	return _DLESS.Contract.Symbol(&_DLESS.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_DLESS *DLESSCallerSession) Symbol() (string, error) {
	return _DLESS.Contract.Symbol(&_DLESS.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_DLESS *DLESSCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DLESS.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_DLESS *DLESSSession) TotalSupply() (*big.Int, error) {
	return _DLESS.Contract.TotalSupply(&_DLESS.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_DLESS *DLESSCallerSession) TotalSupply() (*big.Int, error) {
	return _DLESS.Contract.TotalSupply(&_DLESS.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_DLESS *DLESSTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DLESS.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_DLESS *DLESSSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DLESS.Contract.Approve(&_DLESS.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_DLESS *DLESSTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DLESS.Contract.Approve(&_DLESS.TransactOpts, _spender, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns(bool success)
func (_DLESS *DLESSTransactor) Burn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _DLESS.contract.Transact(opts, "burn", _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns(bool success)
func (_DLESS *DLESSSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _DLESS.Contract.Burn(&_DLESS.TransactOpts, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns(bool success)
func (_DLESS *DLESSTransactorSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _DLESS.Contract.Burn(&_DLESS.TransactOpts, _value)
}

// Freeze is a paid mutator transaction binding the contract method 0xd7a78db8.
//
// Solidity: function freeze(uint256 _value) returns(bool success)
func (_DLESS *DLESSTransactor) Freeze(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _DLESS.contract.Transact(opts, "freeze", _value)
}

// Freeze is a paid mutator transaction binding the contract method 0xd7a78db8.
//
// Solidity: function freeze(uint256 _value) returns(bool success)
func (_DLESS *DLESSSession) Freeze(_value *big.Int) (*types.Transaction, error) {
	return _DLESS.Contract.Freeze(&_DLESS.TransactOpts, _value)
}

// Freeze is a paid mutator transaction binding the contract method 0xd7a78db8.
//
// Solidity: function freeze(uint256 _value) returns(bool success)
func (_DLESS *DLESSTransactorSession) Freeze(_value *big.Int) (*types.Transaction, error) {
	return _DLESS.Contract.Freeze(&_DLESS.TransactOpts, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_DLESS *DLESSTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DLESS.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_DLESS *DLESSSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DLESS.Contract.Transfer(&_DLESS.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_DLESS *DLESSTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DLESS.Contract.Transfer(&_DLESS.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_DLESS *DLESSTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DLESS.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_DLESS *DLESSSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DLESS.Contract.TransferFrom(&_DLESS.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_DLESS *DLESSTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DLESS.Contract.TransferFrom(&_DLESS.TransactOpts, _from, _to, _value)
}

// Unfreeze is a paid mutator transaction binding the contract method 0x6623fc46.
//
// Solidity: function unfreeze(uint256 _value) returns(bool success)
func (_DLESS *DLESSTransactor) Unfreeze(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _DLESS.contract.Transact(opts, "unfreeze", _value)
}

// Unfreeze is a paid mutator transaction binding the contract method 0x6623fc46.
//
// Solidity: function unfreeze(uint256 _value) returns(bool success)
func (_DLESS *DLESSSession) Unfreeze(_value *big.Int) (*types.Transaction, error) {
	return _DLESS.Contract.Unfreeze(&_DLESS.TransactOpts, _value)
}

// Unfreeze is a paid mutator transaction binding the contract method 0x6623fc46.
//
// Solidity: function unfreeze(uint256 _value) returns(bool success)
func (_DLESS *DLESSTransactorSession) Unfreeze(_value *big.Int) (*types.Transaction, error) {
	return _DLESS.Contract.Unfreeze(&_DLESS.TransactOpts, _value)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x3bed33ce.
//
// Solidity: function withdrawEther(uint256 amount) returns()
func (_DLESS *DLESSTransactor) WithdrawEther(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _DLESS.contract.Transact(opts, "withdrawEther", amount)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x3bed33ce.
//
// Solidity: function withdrawEther(uint256 amount) returns()
func (_DLESS *DLESSSession) WithdrawEther(amount *big.Int) (*types.Transaction, error) {
	return _DLESS.Contract.WithdrawEther(&_DLESS.TransactOpts, amount)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x3bed33ce.
//
// Solidity: function withdrawEther(uint256 amount) returns()
func (_DLESS *DLESSTransactorSession) WithdrawEther(amount *big.Int) (*types.Transaction, error) {
	return _DLESS.Contract.WithdrawEther(&_DLESS.TransactOpts, amount)
}

// DLESSBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the DLESS contract.
type DLESSBurnIterator struct {
	Event *DLESSBurn // Event containing the contract specifics and raw log

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
func (it *DLESSBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DLESSBurn)
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
		it.Event = new(DLESSBurn)
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
func (it *DLESSBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DLESSBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DLESSBurn represents a Burn event raised by the DLESS contract.
type DLESSBurn struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_DLESS *DLESSFilterer) FilterBurn(opts *bind.FilterOpts, from []common.Address) (*DLESSBurnIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _DLESS.contract.FilterLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return &DLESSBurnIterator{contract: _DLESS.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_DLESS *DLESSFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *DLESSBurn, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _DLESS.contract.WatchLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DLESSBurn)
				if err := _DLESS.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_DLESS *DLESSFilterer) ParseBurn(log types.Log) (*DLESSBurn, error) {
	event := new(DLESSBurn)
	if err := _DLESS.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DLESSFreezeIterator is returned from FilterFreeze and is used to iterate over the raw logs and unpacked data for Freeze events raised by the DLESS contract.
type DLESSFreezeIterator struct {
	Event *DLESSFreeze // Event containing the contract specifics and raw log

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
func (it *DLESSFreezeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DLESSFreeze)
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
		it.Event = new(DLESSFreeze)
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
func (it *DLESSFreezeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DLESSFreezeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DLESSFreeze represents a Freeze event raised by the DLESS contract.
type DLESSFreeze struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterFreeze is a free log retrieval operation binding the contract event 0xf97a274face0b5517365ad396b1fdba6f68bd3135ef603e44272adba3af5a1e0.
//
// Solidity: event Freeze(address indexed from, uint256 value)
func (_DLESS *DLESSFilterer) FilterFreeze(opts *bind.FilterOpts, from []common.Address) (*DLESSFreezeIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _DLESS.contract.FilterLogs(opts, "Freeze", fromRule)
	if err != nil {
		return nil, err
	}
	return &DLESSFreezeIterator{contract: _DLESS.contract, event: "Freeze", logs: logs, sub: sub}, nil
}

// WatchFreeze is a free log subscription operation binding the contract event 0xf97a274face0b5517365ad396b1fdba6f68bd3135ef603e44272adba3af5a1e0.
//
// Solidity: event Freeze(address indexed from, uint256 value)
func (_DLESS *DLESSFilterer) WatchFreeze(opts *bind.WatchOpts, sink chan<- *DLESSFreeze, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _DLESS.contract.WatchLogs(opts, "Freeze", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DLESSFreeze)
				if err := _DLESS.contract.UnpackLog(event, "Freeze", log); err != nil {
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

// ParseFreeze is a log parse operation binding the contract event 0xf97a274face0b5517365ad396b1fdba6f68bd3135ef603e44272adba3af5a1e0.
//
// Solidity: event Freeze(address indexed from, uint256 value)
func (_DLESS *DLESSFilterer) ParseFreeze(log types.Log) (*DLESSFreeze, error) {
	event := new(DLESSFreeze)
	if err := _DLESS.contract.UnpackLog(event, "Freeze", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DLESSTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the DLESS contract.
type DLESSTransferIterator struct {
	Event *DLESSTransfer // Event containing the contract specifics and raw log

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
func (it *DLESSTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DLESSTransfer)
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
		it.Event = new(DLESSTransfer)
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
func (it *DLESSTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DLESSTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DLESSTransfer represents a Transfer event raised by the DLESS contract.
type DLESSTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_DLESS *DLESSFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DLESSTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DLESS.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DLESSTransferIterator{contract: _DLESS.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_DLESS *DLESSFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DLESSTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DLESS.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DLESSTransfer)
				if err := _DLESS.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_DLESS *DLESSFilterer) ParseTransfer(log types.Log) (*DLESSTransfer, error) {
	event := new(DLESSTransfer)
	if err := _DLESS.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DLESSUnfreezeIterator is returned from FilterUnfreeze and is used to iterate over the raw logs and unpacked data for Unfreeze events raised by the DLESS contract.
type DLESSUnfreezeIterator struct {
	Event *DLESSUnfreeze // Event containing the contract specifics and raw log

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
func (it *DLESSUnfreezeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DLESSUnfreeze)
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
		it.Event = new(DLESSUnfreeze)
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
func (it *DLESSUnfreezeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DLESSUnfreezeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DLESSUnfreeze represents a Unfreeze event raised by the DLESS contract.
type DLESSUnfreeze struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnfreeze is a free log retrieval operation binding the contract event 0x2cfce4af01bcb9d6cf6c84ee1b7c491100b8695368264146a94d71e10a63083f.
//
// Solidity: event Unfreeze(address indexed from, uint256 value)
func (_DLESS *DLESSFilterer) FilterUnfreeze(opts *bind.FilterOpts, from []common.Address) (*DLESSUnfreezeIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _DLESS.contract.FilterLogs(opts, "Unfreeze", fromRule)
	if err != nil {
		return nil, err
	}
	return &DLESSUnfreezeIterator{contract: _DLESS.contract, event: "Unfreeze", logs: logs, sub: sub}, nil
}

// WatchUnfreeze is a free log subscription operation binding the contract event 0x2cfce4af01bcb9d6cf6c84ee1b7c491100b8695368264146a94d71e10a63083f.
//
// Solidity: event Unfreeze(address indexed from, uint256 value)
func (_DLESS *DLESSFilterer) WatchUnfreeze(opts *bind.WatchOpts, sink chan<- *DLESSUnfreeze, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _DLESS.contract.WatchLogs(opts, "Unfreeze", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DLESSUnfreeze)
				if err := _DLESS.contract.UnpackLog(event, "Unfreeze", log); err != nil {
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

// ParseUnfreeze is a log parse operation binding the contract event 0x2cfce4af01bcb9d6cf6c84ee1b7c491100b8695368264146a94d71e10a63083f.
//
// Solidity: event Unfreeze(address indexed from, uint256 value)
func (_DLESS *DLESSFilterer) ParseUnfreeze(log types.Log) (*DLESSUnfreeze, error) {
	event := new(DLESSUnfreeze)
	if err := _DLESS.contract.UnpackLog(event, "Unfreeze", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x6080604052348015600f57600080fd5b50603e80601d6000396000f3fe6080604052600080fdfea265627a7a72315820301debc6645148e77406913398a15a344f69e10efdc59867b47fa5184806a2d764736f6c634300050c0032"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}
