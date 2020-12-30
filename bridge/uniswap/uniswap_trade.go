// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswap

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

// UniswapABI is the input ABI used to generate the binding from.
const UniswapABI = "[{\"inputs\":[{\"internalType\":\"contractUniswapV2\",\"name\":\"_uniswapV2\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ETH_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"paths\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"srcQty\",\"type\":\"uint256\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"paths\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"srcQty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"}],\"name\":\"trade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV2\",\"outputs\":[{\"internalType\":\"contractUniswapV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// UniswapBin is the compiled bytecode used for deploying new contracts.
var UniswapBin = "0x608060405234801561001057600080fd5b506040516111683803806111688339818101604052602081101561003357600080fd5b8101908080519060200190929190505050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ad5c46486040518163ffffffff1660e01b815260040160206040518083038186803b1580156100eb57600080fd5b505afa1580156100ff573d6000803e3d6000fd5b505050506040513d602081101561011557600080fd5b8101908080519060200190929190505050600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610ff2806101766000396000f3fe60806040526004361061004e5760003560e01c806311897c9a1461005a5780635187c0911461013f57806372e94bf614610196578063e5d8a39d146101ed578063f2428621146102cb57610055565b3661005557005b600080fd5b34801561006657600080fd5b506100e86004803603604081101561007d57600080fd5b810190808035906020019064010000000081111561009a57600080fd5b8201836020820111156100ac57600080fd5b803590602001918460208302840111640100000000831117156100ce57600080fd5b909192939192939080359060200190929190505050610322565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561012b578082015181840152602081019050610110565b505050509050019250505060405180910390f35b34801561014b57600080fd5b506101546104ab565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156101a257600080fd5b506101ab6104d0565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6102826004803603608081101561020357600080fd5b810190808035906020019064010000000081111561022057600080fd5b82018360208201111561023257600080fd5b8035906020019184602083028401116401000000008311171561025457600080fd5b90919293919293908035906020019092919080359060200190929190803590602001909291905050506104d5565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b3480156102d757600080fd5b506102e061081e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b60606000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d06ca61f8386866040518463ffffffff1660e01b815260040180848152602001806020018281038252848482818152602001925060200280828437600081840152601f19601f82011690508083019250505094505050505060006040518083038186803b1580156103cd57600080fd5b505afa1580156103e1573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250602081101561040b57600080fd5b810190808051604051939291908464010000000082111561042b57600080fd5b8382019150602082018581111561044157600080fd5b825186602082028301116401000000008211171561045e57600080fd5b8083526020830192505050908051906020019060200280838360005b8381101561049557808201518184015260208101905061047a565b5050505090500160405250505090509392505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600081565b600080600187879050116104e857600080fd5b6000878760018a8a9050038181106104fc57fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff168888600081811061053f57fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561057d57600080fd5b6060600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16898960008181106105c557fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610754576106528989600081811061060d57fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1689610844565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16146106fb576106f4898980806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050508888886109df565b905061074f565b610748898980806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f82011690508083019250505050505050888888610bc1565b9050600091505b6107a4565b6107a1898980806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f82011690508083019250505050505050888888610da3565b90505b60018151116107b257600080fd5b85816001835103815181106107c357fe5b6020026020010151101580156107ec575086816000815181106107e257fe5b6020026020010151145b6107f557600080fd5b818160018351038151811061080657fe5b60200260200101519350935050509550959350505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16146109da578273ffffffffffffffffffffffffffffffffffffffff1663095ea7b38360006040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b15801561090057600080fd5b505af1158015610914573d6000803e3d6000fd5b50505050610920610f7e565b61092957600080fd5b8273ffffffffffffffffffffffffffffffffffffffff1663095ea7b383836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b1580156109b057600080fd5b505af11580156109c4573d6000803e3d6000fd5b505050506109d0610f7e565b6109d957600080fd5b5b505050565b60606000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166338ed173985858833876040518663ffffffff1660e01b815260040180868152602001858152602001806020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828103825285818151815260200191508051906020019060200280838360005b83811015610ab9578082015181840152602081019050610a9e565b505050509050019650505050505050600060405180830381600087803b158015610ae257600080fd5b505af1158015610af6573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052506020811015610b2057600080fd5b8101908080516040519392919084640100000000821115610b4057600080fd5b83820191506020820185811115610b5657600080fd5b8251866020820283011164010000000082111715610b7357600080fd5b8083526020830192505050908051906020019060200280838360005b83811015610baa578082015181840152602081019050610b8f565b505050509050016040525050509050949350505050565b60606000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166318cbafe585858833876040518663ffffffff1660e01b815260040180868152602001858152602001806020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828103825285818151815260200191508051906020019060200280838360005b83811015610c9b578082015181840152602081019050610c80565b505050509050019650505050505050600060405180830381600087803b158015610cc457600080fd5b505af1158015610cd8573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052506020811015610d0257600080fd5b8101908080516040519392919084640100000000821115610d2257600080fd5b83820191506020820185811115610d3857600080fd5b8251866020820283011164010000000082111715610d5557600080fd5b8083526020830192505050908051906020019060200280838360005b83811015610d8c578082015181840152602081019050610d71565b505050509050016040525050509050949350505050565b60606000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637ff36ab585858833876040518663ffffffff1660e01b815260040180858152602001806020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828103825285818151815260200191508051906020019060200280838360005b83811015610e77578082015181840152602081019050610e5c565b50505050905001955050505050506000604051808303818588803b158015610e9e57600080fd5b505af1158015610eb2573d6000803e3d6000fd5b50505050506040513d6000823e3d601f19601f820116820180604052506020811015610edd57600080fd5b8101908080516040519392919084640100000000821115610efd57600080fd5b83820191506020820185811115610f1357600080fd5b8251866020820283011164010000000082111715610f3057600080fd5b8083526020830192505050908051906020019060200280838360005b83811015610f67578082015181840152602081019050610f4c565b505050509050016040525050509050949350505050565b600080600090503d60008114610f9b5760208114610fa457610fb0565b60019150610fb0565b60206000803e60005191505b5060008114159150509056fea26469706673582212207b862a09ce908c4c6585a7f2fb31fbe280b50dd943cc2dcdacd303e3789a3e8b64736f6c63430006060033"

// DeployUniswap deploys a new Ethereum contract, binding an instance of Uniswap to it.
func DeployUniswap(auth *bind.TransactOpts, backend bind.ContractBackend, _uniswapV2 common.Address) (common.Address, *types.Transaction, *Uniswap, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UniswapBin), backend, _uniswapV2)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Uniswap{UniswapCaller: UniswapCaller{contract: contract}, UniswapTransactor: UniswapTransactor{contract: contract}, UniswapFilterer: UniswapFilterer{contract: contract}}, nil
}

// Uniswap is an auto generated Go binding around an Ethereum contract.
type Uniswap struct {
	UniswapCaller     // Read-only binding to the contract
	UniswapTransactor // Write-only binding to the contract
	UniswapFilterer   // Log filterer for contract events
}

// UniswapCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapSession struct {
	Contract     *Uniswap          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniswapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapCallerSession struct {
	Contract *UniswapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// UniswapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapTransactorSession struct {
	Contract     *UniswapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// UniswapRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapRaw struct {
	Contract *Uniswap // Generic contract binding to access the raw methods on
}

// UniswapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapCallerRaw struct {
	Contract *UniswapCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapTransactorRaw struct {
	Contract *UniswapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswap creates a new instance of Uniswap, bound to a specific deployed contract.
func NewUniswap(address common.Address, backend bind.ContractBackend) (*Uniswap, error) {
	contract, err := bindUniswap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Uniswap{UniswapCaller: UniswapCaller{contract: contract}, UniswapTransactor: UniswapTransactor{contract: contract}, UniswapFilterer: UniswapFilterer{contract: contract}}, nil
}

// NewUniswapCaller creates a new read-only instance of Uniswap, bound to a specific deployed contract.
func NewUniswapCaller(address common.Address, caller bind.ContractCaller) (*UniswapCaller, error) {
	contract, err := bindUniswap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapCaller{contract: contract}, nil
}

// NewUniswapTransactor creates a new write-only instance of Uniswap, bound to a specific deployed contract.
func NewUniswapTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapTransactor, error) {
	contract, err := bindUniswap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapTransactor{contract: contract}, nil
}

// NewUniswapFilterer creates a new log filterer instance of Uniswap, bound to a specific deployed contract.
func NewUniswapFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapFilterer, error) {
	contract, err := bindUniswap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapFilterer{contract: contract}, nil
}

// bindUniswap binds a generic wrapper to an already deployed contract.
func bindUniswap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswap *UniswapRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Uniswap.Contract.UniswapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswap *UniswapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswap.Contract.UniswapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswap *UniswapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswap.Contract.UniswapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswap *UniswapCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Uniswap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswap *UniswapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswap *UniswapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswap.Contract.contract.Transact(opts, method, params...)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() view returns(address)
func (_Uniswap *UniswapCaller) ETHCONTRACTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Uniswap.contract.Call(opts, out, "ETH_CONTRACT_ADDRESS")
	return *ret0, err
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() view returns(address)
func (_Uniswap *UniswapSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _Uniswap.Contract.ETHCONTRACTADDRESS(&_Uniswap.CallOpts)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() view returns(address)
func (_Uniswap *UniswapCallerSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _Uniswap.Contract.ETHCONTRACTADDRESS(&_Uniswap.CallOpts)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0x11897c9a.
//
// Solidity: function getAmountsOut(address[] paths, uint256 srcQty) view returns(uint256[])
func (_Uniswap *UniswapCaller) GetAmountsOut(opts *bind.CallOpts, paths []common.Address, srcQty *big.Int) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Uniswap.contract.Call(opts, out, "getAmountsOut", paths, srcQty)
	return *ret0, err
}

// GetAmountsOut is a free data retrieval call binding the contract method 0x11897c9a.
//
// Solidity: function getAmountsOut(address[] paths, uint256 srcQty) view returns(uint256[])
func (_Uniswap *UniswapSession) GetAmountsOut(paths []common.Address, srcQty *big.Int) ([]*big.Int, error) {
	return _Uniswap.Contract.GetAmountsOut(&_Uniswap.CallOpts, paths, srcQty)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0x11897c9a.
//
// Solidity: function getAmountsOut(address[] paths, uint256 srcQty) view returns(uint256[])
func (_Uniswap *UniswapCallerSession) GetAmountsOut(paths []common.Address, srcQty *big.Int) ([]*big.Int, error) {
	return _Uniswap.Contract.GetAmountsOut(&_Uniswap.CallOpts, paths, srcQty)
}

// UniswapV2 is a free data retrieval call binding the contract method 0x5187c091.
//
// Solidity: function uniswapV2() view returns(address)
func (_Uniswap *UniswapCaller) UniswapV2(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Uniswap.contract.Call(opts, out, "uniswapV2")
	return *ret0, err
}

// UniswapV2 is a free data retrieval call binding the contract method 0x5187c091.
//
// Solidity: function uniswapV2() view returns(address)
func (_Uniswap *UniswapSession) UniswapV2() (common.Address, error) {
	return _Uniswap.Contract.UniswapV2(&_Uniswap.CallOpts)
}

// UniswapV2 is a free data retrieval call binding the contract method 0x5187c091.
//
// Solidity: function uniswapV2() view returns(address)
func (_Uniswap *UniswapCallerSession) UniswapV2() (common.Address, error) {
	return _Uniswap.Contract.UniswapV2(&_Uniswap.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xf2428621.
//
// Solidity: function wETH() view returns(address)
func (_Uniswap *UniswapCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Uniswap.contract.Call(opts, out, "wETH")
	return *ret0, err
}

// WETH is a free data retrieval call binding the contract method 0xf2428621.
//
// Solidity: function wETH() view returns(address)
func (_Uniswap *UniswapSession) WETH() (common.Address, error) {
	return _Uniswap.Contract.WETH(&_Uniswap.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xf2428621.
//
// Solidity: function wETH() view returns(address)
func (_Uniswap *UniswapCallerSession) WETH() (common.Address, error) {
	return _Uniswap.Contract.WETH(&_Uniswap.CallOpts)
}

// Trade is a paid mutator transaction binding the contract method 0xe5d8a39d.
//
// Solidity: function trade(address[] paths, uint256 srcQty, uint256 amountOutMin, uint256 timeout) payable returns(address, uint256)
func (_Uniswap *UniswapTransactor) Trade(opts *bind.TransactOpts, paths []common.Address, srcQty *big.Int, amountOutMin *big.Int, timeout *big.Int) (*types.Transaction, error) {
	return _Uniswap.contract.Transact(opts, "trade", paths, srcQty, amountOutMin, timeout)
}

// Trade is a paid mutator transaction binding the contract method 0xe5d8a39d.
//
// Solidity: function trade(address[] paths, uint256 srcQty, uint256 amountOutMin, uint256 timeout) payable returns(address, uint256)
func (_Uniswap *UniswapSession) Trade(paths []common.Address, srcQty *big.Int, amountOutMin *big.Int, timeout *big.Int) (*types.Transaction, error) {
	return _Uniswap.Contract.Trade(&_Uniswap.TransactOpts, paths, srcQty, amountOutMin, timeout)
}

// Trade is a paid mutator transaction binding the contract method 0xe5d8a39d.
//
// Solidity: function trade(address[] paths, uint256 srcQty, uint256 amountOutMin, uint256 timeout) payable returns(address, uint256)
func (_Uniswap *UniswapTransactorSession) Trade(paths []common.Address, srcQty *big.Int, amountOutMin *big.Int, timeout *big.Int) (*types.Transaction, error) {
	return _Uniswap.Contract.Trade(&_Uniswap.TransactOpts, paths, srcQty, amountOutMin, timeout)
}