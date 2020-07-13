// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dappMulti

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

// DappMultiABI is the input ABI used to generate the binding from.
const DappMultiABI = "[{\"inputs\":[{\"internalType\":\"contractKyberNetwork\",\"name\":\"_kyberNetworkProxyContract\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_incognitoSmartContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ETH_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcQty\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"destToken\",\"type\":\"address\"}],\"name\":\"getConversionRates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incognitoSmartContract\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kyberNetworkProxyContract\",\"outputs\":[{\"internalType\":\"contractKyberNetwork\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"srcTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"srcQties\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"destTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"minConversionRates\",\"type\":\"uint256[]\"}],\"name\":\"returnAmountWithoutTransfer\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"srcTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"srcQties\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"destTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"minConversionRates\",\"type\":\"uint256[]\"}],\"name\":\"simplePass\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"testReentrance\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"srcTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"srcQties\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"destTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"minConversionRates\",\"type\":\"uint256[]\"}],\"name\":\"trade\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// DappMultiBin is the compiled bytecode used for deploying new contracts.
var DappMultiBin = "0x608060405234801561001057600080fd5b506040516114063803806114068339818101604052604081101561003357600080fd5b508051602090910151600180546001600160a01b039384166001600160a01b0319918216179091556000805493909216921691909117905561138c8061007a6000396000f3fe60806040526004361061007f5760003560e01c806372e94bf61161004e57806372e94bf61461080f578063785250da14610840578063b42a644b14610855578063bb0fd0401461086a57610086565b80630598e9e81461008b5780630aea818814610358578063398b4d9b146103b457806363f89c8e146105db57610086565b3661008657005b600080fd5b34801561009757600080fd5b506102bf600480360360808110156100ae57600080fd5b810190602081018135600160201b8111156100c857600080fd5b8201836020820111156100da57600080fd5b803590602001918460208302840111600160201b831117156100fb57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561014a57600080fd5b82018360208201111561015c57600080fd5b803590602001918460208302840111600160201b8311171561017d57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156101cc57600080fd5b8201836020820111156101de57600080fd5b803590602001918460208302840111600160201b831117156101ff57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561024e57600080fd5b82018360208201111561026057600080fd5b803590602001918460208302840111600160201b8311171561028157600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506108e5945050505050565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b838110156103035781810151838201526020016102eb565b50505050905001838103825284818151815260200191508051906020019060200280838360005b8381101561034257818101518382015260200161032a565b5050505090500194505050505060405180910390f35b34801561036457600080fd5b5061039b6004803603606081101561037b57600080fd5b506001600160a01b03813581169160208101359160409091013516610970565b6040805192835260208301919091528051918290030190f35b6102bf600480360360808110156103ca57600080fd5b810190602081018135600160201b8111156103e457600080fd5b8201836020820111156103f657600080fd5b803590602001918460208302840111600160201b8311171561041757600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561046657600080fd5b82018360208201111561047857600080fd5b803590602001918460208302840111600160201b8311171561049957600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156104e857600080fd5b8201836020820111156104fa57600080fd5b803590602001918460208302840111600160201b8311171561051b57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561056a57600080fd5b82018360208201111561057c57600080fd5b803590602001918460208302840111600160201b8311171561059d57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610a0b945050505050565b3480156105e757600080fd5b506102bf600480360360808110156105fe57600080fd5b810190602081018135600160201b81111561061857600080fd5b82018360208201111561062a57600080fd5b803590602001918460208302840111600160201b8311171561064b57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561069a57600080fd5b8201836020820111156106ac57600080fd5b803590602001918460208302840111600160201b831117156106cd57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561071c57600080fd5b82018360208201111561072e57600080fd5b803590602001918460208302840111600160201b8311171561074f57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561079e57600080fd5b8201836020820111156107b057600080fd5b803590602001918460208302840111600160201b831117156107d157600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610d3e945050505050565b34801561081b57600080fd5b50610824610db1565b604080516001600160a01b039092168252519081900360200190f35b34801561084c57600080fd5b50610824610db6565b34801561086157600080fd5b50610824610dc5565b34801561087657600080fd5b506102bf6004803603602081101561088d57600080fd5b810190602081018135600160201b8111156108a757600080fd5b8201836020820111156108b957600080fd5b803590602001918460018302840111600160201b831117156108da57600080fd5b509092509050610dd4565b6040805160018082528183019092526060918291829160208083019080368337505060408051600180825281830190925292935060609291506020808301908036833701905050905060008160008151811061093d57fe5b60200260200101818152505060008160008151811061095857fe5b60209081029190910101529097909650945050505050565b6001546040805163809a9e5560e01b81526001600160a01b0386811660048301528481166024830152604482018690528251600094859492169263809a9e55926064808301939192829003018186803b1580156109cc57600080fd5b505afa1580156109e0573d6000803e3d6000fd5b505050506040513d60408110156109f657600080fd5b50805160209091015190969095509350505050565b60005460609081906001600160a01b03163314610a2757600080fd5b84518651148015610a39575085518451145b610a4257600080fd5b8251845114610a5057600080fd5b6060845167ffffffffffffffff81118015610a6a57600080fd5b50604051908082528060200260200182016040528015610a94578160200160208202803683370190505b50905060005b8751811015610d3057868181518110610aaf57fe5b6020026020010151610ad3898381518110610ac657fe5b6020026020010151610ee2565b1015610ade57600080fd5b858181518110610aea57fe5b60200260200101516001600160a01b0316888281518110610b0757fe5b60200260200101516001600160a01b03161415610b2357600080fd5b60006001600160a01b0316888281518110610b3a57fe5b60200260200101516001600160a01b031614610c8157610b96888281518110610b5f57fe5b6020026020010151600160009054906101000a90046001600160a01b0316898481518110610b8957fe5b6020026020010151610f73565b60006001600160a01b0316868281518110610bad57fe5b60200260200101516001600160a01b031614610c2c576000610c1d898381518110610bd457fe5b6020026020010151898481518110610be857fe5b6020026020010151898581518110610bfc57fe5b6020026020010151898681518110610c1057fe5b6020026020010151611088565b11610c2757600080fd5b610c7c565b6000610c72898381518110610c3d57fe5b6020026020010151898481518110610c5157fe5b6020026020010151888581518110610c6557fe5b6020026020010151611125565b11610c7c57600080fd5b610cd1565b6000610cc7878381518110610c9257fe5b6020026020010151898481518110610ca657fe5b6020026020010151888581518110610cba57fe5b60200260200101516111b9565b11610cd157600080fd5b610ce0868281518110610ac657fe5b828281518110610cec57fe5b602002602001018181525050610d28868281518110610d0757fe5b6020026020010151838381518110610d1b57fe5b602002602001015161124b565b600101610a9a565b509396939550929350505050565b60408051600180825281830190925260609182918291602080830190803683375050604080516001808252818301909252929350606092915060208083019080368337019050509050600081600081518110610d9657fe5b60200260200101818152505060018160008151811061095857fe5b600081565b6001546001600160a01b031681565b6000546001600160a01b031681565b60608060008060009054906101000a90046001600160a01b03166001600160a01b03168585604051808383808284376040519201945060009350909150508083038183865af19150503d8060008114610e49576040519150601f19603f3d011682016040523d82523d6000602084013e610e4e565b606091505b5050905080610e5c57600080fd5b604080516001808252818301909252606091602080830190803683375050604080516001808252818301909252929350606092915060208083019080368337019050509050600081600081518110610eb057fe5b602002602001018181525050600081600081518110610ecb57fe5b602090810291909101015290969095509350505050565b60006001600160a01b038216610ef9575047610f6e565b604080516370a0823160e01b815230600482015290516001600160a01b038416916370a08231916024808301926020929190829003018186803b158015610f3f57600080fd5b505afa158015610f53573d6000803e3d6000fd5b505050506040513d6020811015610f6957600080fd5b505190505b919050565b6001600160a01b03831615611083576040805163095ea7b360e01b81526001600160a01b03848116600483015260006024830181905292519086169263095ea7b3926044808201939182900301818387803b158015610fd157600080fd5b505af1158015610fe5573d6000803e3d6000fd5b50505050610ff1611322565b610ffa57600080fd5b826001600160a01b031663095ea7b383836040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050600060405180830381600087803b15801561105a57600080fd5b505af115801561106e573d6000803e3d6000fd5b5050505061107a611322565b61108357600080fd5b505050565b60015460408051637409e2eb60e01b81526001600160a01b0387811660048301526024820187905285811660448301526064820185905291516000939290921691637409e2eb9160848082019260209290919082900301818787803b1580156110f057600080fd5b505af1158015611104573d6000803e3d6000fd5b505050506040513d602081101561111a57600080fd5b505195945050505050565b60015460408051630eee887760e21b81526001600160a01b038681166004830152602482018690526044820185905291516000939290921691633bba21dc9160648082019260209290919082900301818787803b15801561118557600080fd5b505af1158015611199573d6000803e3d6000fd5b505050506040513d60208110156111af57600080fd5b5051949350505050565b6000824710156111c857600080fd5b60015460408051633d15022b60e11b81526001600160a01b0387811660048301526024820186905291519190921691637a2a045691869160448082019260209290919082900301818588803b15801561122057600080fd5b505af1158015611234573d6000803e3d6000fd5b50505050506040513d60208110156111af57600080fd5b6001600160a01b0382166112a5578047101561126657600080fd5b600080546040516001600160a01b039091169183156108fc02918491818181858888f1935050505015801561129f573d6000803e3d6000fd5b5061131e565b600080546040805163a9059cbb60e01b81526001600160a01b0392831660048201526024810185905290519185169263a9059cbb9260448084019382900301818387803b1580156112f557600080fd5b505af1158015611309573d6000803e3d6000fd5b50505050611315611322565b61131e57600080fd5b5050565b6000803d801561133957602081146113425761134e565b6001915061134e565b60206000803e60005191505b50151590509056fea2646970667358221220dea4629f038423d35cdd115d2426f169988fd37a7a7852574e5043e3701c9cc364736f6c63430006060033"

// DeployDappMulti deploys a new Ethereum contract, binding an instance of DappMulti to it.
func DeployDappMulti(auth *bind.TransactOpts, backend bind.ContractBackend, _kyberNetworkProxyContract common.Address, _incognitoSmartContract common.Address) (common.Address, *types.Transaction, *DappMulti, error) {
	parsed, err := abi.JSON(strings.NewReader(DappMultiABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DappMultiBin), backend, _kyberNetworkProxyContract, _incognitoSmartContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DappMulti{DappMultiCaller: DappMultiCaller{contract: contract}, DappMultiTransactor: DappMultiTransactor{contract: contract}, DappMultiFilterer: DappMultiFilterer{contract: contract}}, nil
}

// DappMulti is an auto generated Go binding around an Ethereum contract.
type DappMulti struct {
	DappMultiCaller     // Read-only binding to the contract
	DappMultiTransactor // Write-only binding to the contract
	DappMultiFilterer   // Log filterer for contract events
}

// DappMultiCaller is an auto generated read-only Go binding around an Ethereum contract.
type DappMultiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DappMultiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DappMultiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DappMultiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DappMultiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DappMultiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DappMultiSession struct {
	Contract     *DappMulti        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DappMultiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DappMultiCallerSession struct {
	Contract *DappMultiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DappMultiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DappMultiTransactorSession struct {
	Contract     *DappMultiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DappMultiRaw is an auto generated low-level Go binding around an Ethereum contract.
type DappMultiRaw struct {
	Contract *DappMulti // Generic contract binding to access the raw methods on
}

// DappMultiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DappMultiCallerRaw struct {
	Contract *DappMultiCaller // Generic read-only contract binding to access the raw methods on
}

// DappMultiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DappMultiTransactorRaw struct {
	Contract *DappMultiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDappMulti creates a new instance of DappMulti, bound to a specific deployed contract.
func NewDappMulti(address common.Address, backend bind.ContractBackend) (*DappMulti, error) {
	contract, err := bindDappMulti(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DappMulti{DappMultiCaller: DappMultiCaller{contract: contract}, DappMultiTransactor: DappMultiTransactor{contract: contract}, DappMultiFilterer: DappMultiFilterer{contract: contract}}, nil
}

// NewDappMultiCaller creates a new read-only instance of DappMulti, bound to a specific deployed contract.
func NewDappMultiCaller(address common.Address, caller bind.ContractCaller) (*DappMultiCaller, error) {
	contract, err := bindDappMulti(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DappMultiCaller{contract: contract}, nil
}

// NewDappMultiTransactor creates a new write-only instance of DappMulti, bound to a specific deployed contract.
func NewDappMultiTransactor(address common.Address, transactor bind.ContractTransactor) (*DappMultiTransactor, error) {
	contract, err := bindDappMulti(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DappMultiTransactor{contract: contract}, nil
}

// NewDappMultiFilterer creates a new log filterer instance of DappMulti, bound to a specific deployed contract.
func NewDappMultiFilterer(address common.Address, filterer bind.ContractFilterer) (*DappMultiFilterer, error) {
	contract, err := bindDappMulti(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DappMultiFilterer{contract: contract}, nil
}

// bindDappMulti binds a generic wrapper to an already deployed contract.
func bindDappMulti(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DappMultiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DappMulti *DappMultiRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DappMulti.Contract.DappMultiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DappMulti *DappMultiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappMulti.Contract.DappMultiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DappMulti *DappMultiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DappMulti.Contract.DappMultiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DappMulti *DappMultiCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DappMulti.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DappMulti *DappMultiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappMulti.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DappMulti *DappMultiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DappMulti.Contract.contract.Transact(opts, method, params...)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() view returns(address)
func (_DappMulti *DappMultiCaller) ETHCONTRACTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DappMulti.contract.Call(opts, out, "ETH_CONTRACT_ADDRESS")
	return *ret0, err
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() view returns(address)
func (_DappMulti *DappMultiSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _DappMulti.Contract.ETHCONTRACTADDRESS(&_DappMulti.CallOpts)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() view returns(address)
func (_DappMulti *DappMultiCallerSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _DappMulti.Contract.ETHCONTRACTADDRESS(&_DappMulti.CallOpts)
}

// GetConversionRates is a free data retrieval call binding the contract method 0x0aea8188.
//
// Solidity: function getConversionRates(address srcToken, uint256 srcQty, address destToken) view returns(uint256, uint256)
func (_DappMulti *DappMultiCaller) GetConversionRates(opts *bind.CallOpts, srcToken common.Address, srcQty *big.Int, destToken common.Address) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _DappMulti.contract.Call(opts, out, "getConversionRates", srcToken, srcQty, destToken)
	return *ret0, *ret1, err
}

// GetConversionRates is a free data retrieval call binding the contract method 0x0aea8188.
//
// Solidity: function getConversionRates(address srcToken, uint256 srcQty, address destToken) view returns(uint256, uint256)
func (_DappMulti *DappMultiSession) GetConversionRates(srcToken common.Address, srcQty *big.Int, destToken common.Address) (*big.Int, *big.Int, error) {
	return _DappMulti.Contract.GetConversionRates(&_DappMulti.CallOpts, srcToken, srcQty, destToken)
}

// GetConversionRates is a free data retrieval call binding the contract method 0x0aea8188.
//
// Solidity: function getConversionRates(address srcToken, uint256 srcQty, address destToken) view returns(uint256, uint256)
func (_DappMulti *DappMultiCallerSession) GetConversionRates(srcToken common.Address, srcQty *big.Int, destToken common.Address) (*big.Int, *big.Int, error) {
	return _DappMulti.Contract.GetConversionRates(&_DappMulti.CallOpts, srcToken, srcQty, destToken)
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() view returns(address)
func (_DappMulti *DappMultiCaller) IncognitoSmartContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DappMulti.contract.Call(opts, out, "incognitoSmartContract")
	return *ret0, err
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() view returns(address)
func (_DappMulti *DappMultiSession) IncognitoSmartContract() (common.Address, error) {
	return _DappMulti.Contract.IncognitoSmartContract(&_DappMulti.CallOpts)
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() view returns(address)
func (_DappMulti *DappMultiCallerSession) IncognitoSmartContract() (common.Address, error) {
	return _DappMulti.Contract.IncognitoSmartContract(&_DappMulti.CallOpts)
}

// KyberNetworkProxyContract is a free data retrieval call binding the contract method 0x785250da.
//
// Solidity: function kyberNetworkProxyContract() view returns(address)
func (_DappMulti *DappMultiCaller) KyberNetworkProxyContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DappMulti.contract.Call(opts, out, "kyberNetworkProxyContract")
	return *ret0, err
}

// KyberNetworkProxyContract is a free data retrieval call binding the contract method 0x785250da.
//
// Solidity: function kyberNetworkProxyContract() view returns(address)
func (_DappMulti *DappMultiSession) KyberNetworkProxyContract() (common.Address, error) {
	return _DappMulti.Contract.KyberNetworkProxyContract(&_DappMulti.CallOpts)
}

// KyberNetworkProxyContract is a free data retrieval call binding the contract method 0x785250da.
//
// Solidity: function kyberNetworkProxyContract() view returns(address)
func (_DappMulti *DappMultiCallerSession) KyberNetworkProxyContract() (common.Address, error) {
	return _DappMulti.Contract.KyberNetworkProxyContract(&_DappMulti.CallOpts)
}

// ReturnAmountWithoutTransfer is a paid mutator transaction binding the contract method 0x63f89c8e.
//
// Solidity: function returnAmountWithoutTransfer(address[] srcTokens, uint256[] srcQties, address[] destTokens, uint256[] minConversionRates) returns(address[], uint256[])
func (_DappMulti *DappMultiTransactor) ReturnAmountWithoutTransfer(opts *bind.TransactOpts, srcTokens []common.Address, srcQties []*big.Int, destTokens []common.Address, minConversionRates []*big.Int) (*types.Transaction, error) {
	return _DappMulti.contract.Transact(opts, "returnAmountWithoutTransfer", srcTokens, srcQties, destTokens, minConversionRates)
}

// ReturnAmountWithoutTransfer is a paid mutator transaction binding the contract method 0x63f89c8e.
//
// Solidity: function returnAmountWithoutTransfer(address[] srcTokens, uint256[] srcQties, address[] destTokens, uint256[] minConversionRates) returns(address[], uint256[])
func (_DappMulti *DappMultiSession) ReturnAmountWithoutTransfer(srcTokens []common.Address, srcQties []*big.Int, destTokens []common.Address, minConversionRates []*big.Int) (*types.Transaction, error) {
	return _DappMulti.Contract.ReturnAmountWithoutTransfer(&_DappMulti.TransactOpts, srcTokens, srcQties, destTokens, minConversionRates)
}

// ReturnAmountWithoutTransfer is a paid mutator transaction binding the contract method 0x63f89c8e.
//
// Solidity: function returnAmountWithoutTransfer(address[] srcTokens, uint256[] srcQties, address[] destTokens, uint256[] minConversionRates) returns(address[], uint256[])
func (_DappMulti *DappMultiTransactorSession) ReturnAmountWithoutTransfer(srcTokens []common.Address, srcQties []*big.Int, destTokens []common.Address, minConversionRates []*big.Int) (*types.Transaction, error) {
	return _DappMulti.Contract.ReturnAmountWithoutTransfer(&_DappMulti.TransactOpts, srcTokens, srcQties, destTokens, minConversionRates)
}

// SimplePass is a paid mutator transaction binding the contract method 0x0598e9e8.
//
// Solidity: function simplePass(address[] srcTokens, uint256[] srcQties, address[] destTokens, uint256[] minConversionRates) returns(address[], uint256[])
func (_DappMulti *DappMultiTransactor) SimplePass(opts *bind.TransactOpts, srcTokens []common.Address, srcQties []*big.Int, destTokens []common.Address, minConversionRates []*big.Int) (*types.Transaction, error) {
	return _DappMulti.contract.Transact(opts, "simplePass", srcTokens, srcQties, destTokens, minConversionRates)
}

// SimplePass is a paid mutator transaction binding the contract method 0x0598e9e8.
//
// Solidity: function simplePass(address[] srcTokens, uint256[] srcQties, address[] destTokens, uint256[] minConversionRates) returns(address[], uint256[])
func (_DappMulti *DappMultiSession) SimplePass(srcTokens []common.Address, srcQties []*big.Int, destTokens []common.Address, minConversionRates []*big.Int) (*types.Transaction, error) {
	return _DappMulti.Contract.SimplePass(&_DappMulti.TransactOpts, srcTokens, srcQties, destTokens, minConversionRates)
}

// SimplePass is a paid mutator transaction binding the contract method 0x0598e9e8.
//
// Solidity: function simplePass(address[] srcTokens, uint256[] srcQties, address[] destTokens, uint256[] minConversionRates) returns(address[], uint256[])
func (_DappMulti *DappMultiTransactorSession) SimplePass(srcTokens []common.Address, srcQties []*big.Int, destTokens []common.Address, minConversionRates []*big.Int) (*types.Transaction, error) {
	return _DappMulti.Contract.SimplePass(&_DappMulti.TransactOpts, srcTokens, srcQties, destTokens, minConversionRates)
}

// TestReentrance is a paid mutator transaction binding the contract method 0xbb0fd040.
//
// Solidity: function testReentrance(bytes data) returns(address[], uint256[])
func (_DappMulti *DappMultiTransactor) TestReentrance(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _DappMulti.contract.Transact(opts, "testReentrance", data)
}

// TestReentrance is a paid mutator transaction binding the contract method 0xbb0fd040.
//
// Solidity: function testReentrance(bytes data) returns(address[], uint256[])
func (_DappMulti *DappMultiSession) TestReentrance(data []byte) (*types.Transaction, error) {
	return _DappMulti.Contract.TestReentrance(&_DappMulti.TransactOpts, data)
}

// TestReentrance is a paid mutator transaction binding the contract method 0xbb0fd040.
//
// Solidity: function testReentrance(bytes data) returns(address[], uint256[])
func (_DappMulti *DappMultiTransactorSession) TestReentrance(data []byte) (*types.Transaction, error) {
	return _DappMulti.Contract.TestReentrance(&_DappMulti.TransactOpts, data)
}

// Trade is a paid mutator transaction binding the contract method 0x398b4d9b.
//
// Solidity: function trade(address[] srcTokens, uint256[] srcQties, address[] destTokens, uint256[] minConversionRates) payable returns(address[], uint256[])
func (_DappMulti *DappMultiTransactor) Trade(opts *bind.TransactOpts, srcTokens []common.Address, srcQties []*big.Int, destTokens []common.Address, minConversionRates []*big.Int) (*types.Transaction, error) {
	return _DappMulti.contract.Transact(opts, "trade", srcTokens, srcQties, destTokens, minConversionRates)
}

// Trade is a paid mutator transaction binding the contract method 0x398b4d9b.
//
// Solidity: function trade(address[] srcTokens, uint256[] srcQties, address[] destTokens, uint256[] minConversionRates) payable returns(address[], uint256[])
func (_DappMulti *DappMultiSession) Trade(srcTokens []common.Address, srcQties []*big.Int, destTokens []common.Address, minConversionRates []*big.Int) (*types.Transaction, error) {
	return _DappMulti.Contract.Trade(&_DappMulti.TransactOpts, srcTokens, srcQties, destTokens, minConversionRates)
}

// Trade is a paid mutator transaction binding the contract method 0x398b4d9b.
//
// Solidity: function trade(address[] srcTokens, uint256[] srcQties, address[] destTokens, uint256[] minConversionRates) payable returns(address[], uint256[])
func (_DappMulti *DappMultiTransactorSession) Trade(srcTokens []common.Address, srcQties []*big.Int, destTokens []common.Address, minConversionRates []*big.Int) (*types.Transaction, error) {
	return _DappMulti.Contract.Trade(&_DappMulti.TransactOpts, srcTokens, srcQties, destTokens, minConversionRates)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
// func (_DappMulti *DappMultiTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
// 	return _DappMulti.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
// }

// // Receive is a paid mutator transaction binding the contract receive function.
// //
// // Solidity: receive() payable returns()
// func (_DappMulti *DappMultiSession) Receive() (*types.Transaction, error) {
// 	return _DappMulti.Contract.Receive(&_DappMulti.TransactOpts)
// }

// // Receive is a paid mutator transaction binding the contract receive function.
// //
// // Solidity: receive() payable returns()
// func (_DappMulti *DappMultiTransactorSession) Receive() (*types.Transaction, error) {
// 	return _DappMulti.Contract.Receive(&_DappMulti.TransactOpts)
// }
