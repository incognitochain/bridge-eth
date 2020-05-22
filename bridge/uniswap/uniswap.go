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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// UniswapABI is the input ABI used to generate the binding from.
const UniswapABI = "[{\"inputs\":[{\"internalType\":\"contractUniswapV2\",\"name\":\"_uniswapV2\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_incognitoSmartContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ETH_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcQty\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incognitoSmartContract\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcQty\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"}],\"name\":\"trade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV2\",\"outputs\":[{\"internalType\":\"contractUniswapV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// UniswapBin is the compiled bytecode used for deploying new contracts.
var UniswapBin = "0x608060405234801561001057600080fd5b506040516114f33803806114f38339818101604052604081101561003357600080fd5b81019080805190602001909291908051906020019092919050505081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ad5c46486040518163ffffffff1660e01b815260040160206040518083038186803b15801561013757600080fd5b505afa15801561014b573d6000803e3d6000fd5b505050506040513d602081101561016157600080fd5b8101908080519060200190929190505050600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050611330806101c36000396000f3fe6080604052600436106100595760003560e01c80635187c0911461006557806372e94bf6146100bc578063b42a644b14610113578063bb39a9601461016a578063cba7064f14610229578063f2428621146102f957610060565b3661006057005b600080fd5b34801561007157600080fd5b5061007a610350565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156100c857600080fd5b506100d1610376565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561011f57600080fd5b5061012861037b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101e06004803603608081101561018057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506103a0565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b34801561023557600080fd5b506102a26004803603606081101561024c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610777565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156102e55780820151818401526020810190506102ca565b505050509050019250505060405180910390f35b34801561030557600080fd5b5061030e6109ef565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600081565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103fc57600080fd5b8461040687610a15565b101561041157600080fd5b8373ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16141561044a57600080fd5b6060600267ffffffffffffffff8111801561046457600080fd5b506040519080825280602002602001820160405280156104935781602001602082028036833780820191505090505b5090506060600073ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff161461064c5787826000815181106104da57fe5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505061054188600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1689610b13565b600073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16146105cf57858260018151811061058357fe5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506105c8828887610cae565b9050610647565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16826001815181106105ff57fe5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050610644828887610eb5565b90505b61070c565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168260008151811061067c57fe5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505085826001815181106106c457fe5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506107098288876110bc565b90505b848160018351038151811061071d57fe5b6020026020010151101580156107465750868160008151811061073c57fe5b6020026020010151145b61074f57600080fd5b858160018351038151811061076057fe5b602002602001015193509350505094509492505050565b606080600267ffffffffffffffff8111801561079257600080fd5b506040519080825280602002602001820160405280156107c15781602001602082028036833780820191505090505b50905084816000815181106107d257fe5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050828160018151811061081a57fe5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d06ca61f85836040518363ffffffff1660e01b81526004018083815260200180602001828103825283818151815260200191508051906020019060200280838360005b838110156108ec5780820151818401526020810190506108d1565b50505050905001935050505060006040518083038186803b15801561091057600080fd5b505afa158015610924573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250602081101561094e57600080fd5b810190808051604051939291908464010000000082111561096e57600080fd5b8382019150602082018581111561098457600080fd5b82518660208202830111640100000000821117156109a157600080fd5b8083526020830192505050908051906020019060200280838360005b838110156109d85780820151818401526020810190506109bd565b505050509050016040525050509150509392505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610a5357479050610b0e565b8173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015610ad057600080fd5b505afa158015610ae4573d6000803e3d6000fd5b505050506040513d6020811015610afa57600080fd5b810190808051906020019092919050505090505b919050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614610ca9578273ffffffffffffffffffffffffffffffffffffffff1663095ea7b38360006040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b158015610bcf57600080fd5b505af1158015610be3573d6000803e3d6000fd5b50505050610bef6112bc565b610bf857600080fd5b8273ffffffffffffffffffffffffffffffffffffffff1663095ea7b383836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b158015610c7f57600080fd5b505af1158015610c93573d6000803e3d6000fd5b50505050610c9f6112bc565b610ca857600080fd5b5b505050565b6060600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166338ed17398484876000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1661025842016040518663ffffffff1660e01b815260040180868152602001858152602001806020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828103825285818151815260200191508051906020019060200280838360005b83811015610dae578082015181840152602081019050610d93565b505050509050019650505050505050600060405180830381600087803b158015610dd757600080fd5b505af1158015610deb573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052506020811015610e1557600080fd5b8101908080516040519392919084640100000000821115610e3557600080fd5b83820191506020820185811115610e4b57600080fd5b8251866020820283011164010000000082111715610e6857600080fd5b8083526020830192505050908051906020019060200280838360005b83811015610e9f578082015181840152602081019050610e84565b5050505090500160405250505090509392505050565b6060600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166318cbafe58484876000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1661025842016040518663ffffffff1660e01b815260040180868152602001858152602001806020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828103825285818151815260200191508051906020019060200280838360005b83811015610fb5578082015181840152602081019050610f9a565b505050509050019650505050505050600060405180830381600087803b158015610fde57600080fd5b505af1158015610ff2573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250602081101561101c57600080fd5b810190808051604051939291908464010000000082111561103c57600080fd5b8382019150602082018581111561105257600080fd5b825186602082028301116401000000008211171561106f57600080fd5b8083526020830192505050908051906020019060200280838360005b838110156110a657808201518184015260208101905061108b565b5050505090500160405250505090509392505050565b6060600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637ff36ab58484876000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1661025842016040518663ffffffff1660e01b815260040180858152602001806020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828103825285818151815260200191508051906020019060200280838360005b838110156111b657808201518184015260208101905061119b565b50505050905001955050505050506000604051808303818588803b1580156111dd57600080fd5b505af11580156111f1573d6000803e3d6000fd5b50505050506040513d6000823e3d601f19601f82011682018060405250602081101561121c57600080fd5b810190808051604051939291908464010000000082111561123c57600080fd5b8382019150602082018581111561125257600080fd5b825186602082028301116401000000008211171561126f57600080fd5b8083526020830192505050908051906020019060200280838360005b838110156112a657808201518184015260208101905061128b565b5050505090500160405250505090509392505050565b600080600090503d600081146112d957602081146112e2576112ee565b600191506112ee565b60206000803e60005191505b5060008114159150509056fea26469706673582212200626d9c07c63d9e24c5f20b22ccad38eabd2d493761420be4d2e4aed5e55a5ba64736f6c63430006060033"

// DeployUniswap deploys a new Ethereum contract, binding an instance of Uniswap to it.
func DeployUniswap(auth *bind.TransactOpts, backend bind.ContractBackend, _uniswapV2 common.Address, _incognitoSmartContract common.Address) (common.Address, *types.Transaction, *Uniswap, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UniswapBin), backend, _uniswapV2, _incognitoSmartContract)
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

// GetAmountsOut is a free data retrieval call binding the contract method 0xcba7064f.
//
// Solidity: function getAmountsOut(address srcToken, uint256 srcQty, address destToken) view returns(uint256[])
func (_Uniswap *UniswapCaller) GetAmountsOut(opts *bind.CallOpts, srcToken common.Address, srcQty *big.Int, destToken common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Uniswap.contract.Call(opts, out, "getAmountsOut", srcToken, srcQty, destToken)
	return *ret0, err
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xcba7064f.
//
// Solidity: function getAmountsOut(address srcToken, uint256 srcQty, address destToken) view returns(uint256[])
func (_Uniswap *UniswapSession) GetAmountsOut(srcToken common.Address, srcQty *big.Int, destToken common.Address) ([]*big.Int, error) {
	return _Uniswap.Contract.GetAmountsOut(&_Uniswap.CallOpts, srcToken, srcQty, destToken)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xcba7064f.
//
// Solidity: function getAmountsOut(address srcToken, uint256 srcQty, address destToken) view returns(uint256[])
func (_Uniswap *UniswapCallerSession) GetAmountsOut(srcToken common.Address, srcQty *big.Int, destToken common.Address) ([]*big.Int, error) {
	return _Uniswap.Contract.GetAmountsOut(&_Uniswap.CallOpts, srcToken, srcQty, destToken)
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() view returns(address)
func (_Uniswap *UniswapCaller) IncognitoSmartContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Uniswap.contract.Call(opts, out, "incognitoSmartContract")
	return *ret0, err
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() view returns(address)
func (_Uniswap *UniswapSession) IncognitoSmartContract() (common.Address, error) {
	return _Uniswap.Contract.IncognitoSmartContract(&_Uniswap.CallOpts)
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() view returns(address)
func (_Uniswap *UniswapCallerSession) IncognitoSmartContract() (common.Address, error) {
	return _Uniswap.Contract.IncognitoSmartContract(&_Uniswap.CallOpts)
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

// Trade is a paid mutator transaction binding the contract method 0xbb39a960.
//
// Solidity: function trade(address srcToken, uint256 srcQty, address destToken, uint256 amountOutMin) payable returns(address, uint256)
func (_Uniswap *UniswapTransactor) Trade(opts *bind.TransactOpts, srcToken common.Address, srcQty *big.Int, destToken common.Address, amountOutMin *big.Int) (*types.Transaction, error) {
	return _Uniswap.contract.Transact(opts, "trade", srcToken, srcQty, destToken, amountOutMin)
}

// Trade is a paid mutator transaction binding the contract method 0xbb39a960.
//
// Solidity: function trade(address srcToken, uint256 srcQty, address destToken, uint256 amountOutMin) payable returns(address, uint256)
func (_Uniswap *UniswapSession) Trade(srcToken common.Address, srcQty *big.Int, destToken common.Address, amountOutMin *big.Int) (*types.Transaction, error) {
	return _Uniswap.Contract.Trade(&_Uniswap.TransactOpts, srcToken, srcQty, destToken, amountOutMin)
}

// Trade is a paid mutator transaction binding the contract method 0xbb39a960.
//
// Solidity: function trade(address srcToken, uint256 srcQty, address destToken, uint256 amountOutMin) payable returns(address, uint256)
func (_Uniswap *UniswapTransactorSession) Trade(srcToken common.Address, srcQty *big.Int, destToken common.Address, amountOutMin *big.Int) (*types.Transaction, error) {
	return _Uniswap.Contract.Trade(&_Uniswap.TransactOpts, srcToken, srcQty, destToken, amountOutMin)
}
