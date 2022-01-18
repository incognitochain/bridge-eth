// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package puniswap

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

// ISwapRouter2ExactInputParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouter2ExactInputParams struct {
	Path             []byte
	Recipient        common.Address
	AmountIn         *big.Int
	AmountOutMinimum *big.Int
}

// ISwapRouter2ExactInputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouter2ExactInputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	AmountIn          *big.Int
	AmountOutMinimum  *big.Int
	SqrtPriceLimitX96 *big.Int
}

// PuniswapMetaData contains all meta data concerning the Puniswap contract.
var PuniswapMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractISwapRouter2\",\"name\":\"_swaproute02\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ETH_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swaprouter02\",\"outputs\":[{\"internalType\":\"contractISwapRouter2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"}],\"internalType\":\"structISwapRouter2.ExactInputParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"isNative\",\"type\":\"bool\"}],\"name\":\"tradeInput\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structISwapRouter2.ExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"isNative\",\"type\":\"bool\"}],\"name\":\"tradeInputSingle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wmatic\",\"outputs\":[{\"internalType\":\"contractWmatic\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405260405162001834380380620018348339818101604052810190620000299190620001ae565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634aa4a4fc6040518163ffffffff1660e01b8152600401602060405180830381600087803b158015620000d257600080fd5b505af1158015620000e7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200010d919062000182565b600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505062000256565b600081519050620001658162000222565b92915050565b6000815190506200017c816200023c565b92915050565b6000602082840312156200019557600080fd5b6000620001a58482850162000154565b91505092915050565b600060208284031215620001c157600080fd5b6000620001d1848285016200016b565b91505092915050565b6000620001e78262000202565b9050919050565b6000620001fb82620001da565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200022d81620001da565b81146200023957600080fd5b50565b6200024781620001ee565b81146200025357600080fd5b50565b6115ce80620002666000396000f3fe6080604052600436106100595760003560e01c8063421f4388146100655780636b150c3c1461009657806372e94bf6146100c1578063c8dc75e6146100ec578063d49d51811461011d578063fb41be161461014857610060565b3661006057005b600080fd5b61007f600480360381019061007a9190610ee1565b610173565b60405161008d929190611222565b60405180910390f35b3480156100a257600080fd5b506100ab6102be565b6040516100b8919061124b565b60405180910390f35b3480156100cd57600080fd5b506100d66102e2565b6040516100e391906111b5565b60405180910390f35b61010660048036038101906101019190610e8d565b6102e7565b604051610114929190611222565b60405180910390f35b34801561012957600080fd5b506101326104c9565b60405161013f91906112de565b60405180910390f35b34801561015457600080fd5b5061015d6104ed565b60405161016a9190611266565b60405180910390f35b60008061019684600001602081019061018c9190610e64565b8560800135610513565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166304e45aaf34876040518363ffffffff1660e01b81526004016101f391906112c3565b6020604051808303818588803b15801561020c57600080fd5b505af1158015610220573d6000803e3d6000fd5b50505050506040513d601f19601f820116820180604052508101906102459190610f1e565b90508460a0013581101561028e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161028590611281565b60405180910390fd5b60006102ad8660200160208101906102a69190610e64565b8387610683565b905080829350935050509250929050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600081565b60008060006103478580600001906102ff91906112f9565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505061078f565b50509050610359818660400135610513565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b858183f34886040518363ffffffff1660e01b81526004016103b691906112a1565b6020604051808303818588803b1580156103cf57600080fd5b505af11580156103e3573d6000803e3d6000fd5b50505050506040513d601f19601f820116820180604052508101906104089190610f1e565b9050600086806000019061041c91906112f9565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050905060005b6001156104ab576000610476836107e0565b9050801561048e57610487836107fb565b92506104a5565b6104978361078f565b9091505080925050506104ab565b50610464565b6104b6818489610683565b9050808395509550505050509250929050565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000341480156105cc5750808273ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e3060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518363ffffffff1660e01b815260040161057a9291906111d0565b60206040518083038186803b15801561059257600080fd5b505afa1580156105a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105ca9190610f1e565b105b1561067f578173ffffffffffffffffffffffffffffffffffffffff1663095ea7b360008054906101000a900473ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6040518363ffffffff1660e01b815260040161064c929190611222565b600060405180830381600087803b15801561066657600080fd5b505af115801561067a573d6000803e3d6000fd5b505050505b5050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161480156106df5750815b1561078457600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632e1a7d4d846040518263ffffffff1660e01b815260040161073f91906112de565b600060405180830381600087803b15801561075957600080fd5b505af115801561076d573d6000803e3d6000fd5b505050506000905061077f8184610824565b610788565b8390505b9392505050565b60008060006107a860008561096590919063ffffffff16565b92506107be601485610a7e90919063ffffffff16565b90506107d760036014018561096590919063ffffffff16565b91509193909250565b60006003601401601460036014010101825110159050919050565b606061081d6003601401600360140184510384610b889092919063ffffffff16565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156108e2578047101561086657600080fd5b60003373ffffffffffffffffffffffffffffffffffffffff168260405161088c906111a0565b60006040518083038185875af1925050503d80600081146108c9576040519150601f19603f3d011682016040523d82523d6000602084013e6108ce565b606091505b50509050806108dc57600080fd5b50610961565b8173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b815260040161091d9291906111f9565b600060405180830381600087803b15801561093757600080fd5b505af115801561094b573d6000803e3d6000fd5b50505050610957610d72565b61096057600080fd5b5b5050565b6000816014830110156109e0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f746f416464726573735f6f766572666c6f77000000000000000000000000000081525060200191505060405180910390fd5b6014820183511015610a5a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f746f416464726573735f6f75744f66426f756e6473000000000000000000000081525060200191505060405180910390fd5b60006c01000000000000000000000000836020860101510490508091505092915050565b600081600383011015610af9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f746f55696e7432345f6f766572666c6f7700000000000000000000000000000081525060200191505060405180910390fd5b6003820183511015610b73576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f746f55696e7432345f6f75744f66426f756e647300000000000000000000000081525060200191505060405180910390fd5b60008260038501015190508091505092915050565b606081601f83011015610c03576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f736c6963655f6f766572666c6f7700000000000000000000000000000000000081525060200191505060405180910390fd5b828284011015610c7b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f736c6963655f6f766572666c6f7700000000000000000000000000000000000081525060200191505060405180910390fd5b81830184511015610cf4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f736c6963655f6f75744f66426f756e647300000000000000000000000000000081525060200191505060405180910390fd5b6060821560008114610d155760405191506000825260208201604052610d66565b6040519150601f8416801560200281840101858101878315602002848b0101015b81831015610d535780518352602083019250602081019050610d36565b50868552601f19601f8301166040525050505b50809150509392505050565b600080600090503d60008114610d8f5760208114610d9857610da4565b60019150610da4565b60206000803e60005191505b50600081141591505090565b600081359050610dbf81611525565b92915050565b600081359050610dd48161153c565b92915050565b600060808284031215610dec57600080fd5b81905092915050565b600060e08284031215610e0757600080fd5b81905092915050565b600081359050610e1f81611553565b92915050565b600081359050610e348161156a565b92915050565b600081359050610e4981611581565b92915050565b600081519050610e5e81611581565b92915050565b600060208284031215610e7657600080fd5b6000610e8484828501610db0565b91505092915050565b60008060408385031215610ea057600080fd5b600083013567ffffffffffffffff811115610eba57600080fd5b610ec685828601610dda565b9250506020610ed785828601610dc5565b9150509250929050565b6000806101008385031215610ef557600080fd5b6000610f0385828601610df5565b92505060e0610f1485828601610dc5565b9150509250929050565b600060208284031215610f3057600080fd5b6000610f3e84828501610e4f565b91505092915050565b610f5081611487565b82525050565b610f5f81611430565b82525050565b610f6e81611430565b82525050565b6000610f808385611350565b9350610f8d838584611505565b610f9683611514565b840190509392505050565b610faa81611499565b82525050565b610fb9816114bd565b82525050565b6000610fcc601a8361136c565b91507f6c6f776572207468616e206578706563746564206f75747075740000000000006000830152602082019050919050565b600061100c600083611361565b9150600082019050919050565b60006080830161102c6000840184611394565b858303600087015261103f838284610f74565b92505050611050602084018461137d565b61105d6020860182610f56565b5061106b6040840184611419565b6110786040860182611182565b506110866060840184611419565b6110936060860182611182565b508091505092915050565b60e082016110af600083018361137d565b6110bc6000850182610f56565b506110ca602083018361137d565b6110d76020850182610f56565b506110e56040830183611402565b6110f26040850182611173565b50611100606083018361137d565b61110d6060850182610f56565b5061111b6080830183611419565b6111286080850182611182565b5061113660a0830183611419565b61114360a0850182611182565b5061115160c08301836113eb565b61115e60c0850182611164565b50505050565b61116d8161144e565b82525050565b61117c8161146e565b82525050565b61118b8161147d565b82525050565b61119a8161147d565b82525050565b60006111ab82610fff565b9150819050919050565b60006020820190506111ca6000830184610f65565b92915050565b60006040820190506111e56000830185610f47565b6111f26020830184610f65565b9392505050565b600060408201905061120e6000830185610f47565b61121b6020830184611191565b9392505050565b60006040820190506112376000830185610f65565b6112446020830184611191565b9392505050565b60006020820190506112606000830184610fa1565b92915050565b600060208201905061127b6000830184610fb0565b92915050565b6000602082019050818103600083015261129a81610fbf565b9050919050565b600060208201905081810360008301526112bb8184611019565b905092915050565b600060e0820190506112d8600083018461109e565b92915050565b60006020820190506112f36000830184611191565b92915050565b6000808335600160200384360303811261131257600080fd5b80840192508235915067ffffffffffffffff82111561133057600080fd5b60208301925060018202360383131561134857600080fd5b509250929050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600061138c6020840184610db0565b905092915050565b600080833560016020038436030381126113ad57600080fd5b83810192508235915060208301925067ffffffffffffffff8211156113d157600080fd5b6001820236038413156113e357600080fd5b509250929050565b60006113fa6020840184610e10565b905092915050565b60006114116020840184610e25565b905092915050565b60006114286020840184610e3a565b905092915050565b600061143b8261144e565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600062ffffff82169050919050565b6000819050919050565b6000611492826114e1565b9050919050565b60006114a4826114ab565b9050919050565b60006114b68261144e565b9050919050565b60006114c8826114cf565b9050919050565b60006114da8261144e565b9050919050565b60006114ec826114f3565b9050919050565b60006114fe8261144e565b9050919050565b82818337600083830152505050565b6000601f19601f8301169050919050565b61152e81611430565b811461153957600080fd5b50565b61154581611442565b811461155057600080fd5b50565b61155c8161144e565b811461156757600080fd5b50565b6115738161146e565b811461157e57600080fd5b50565b61158a8161147d565b811461159557600080fd5b5056fea2646970667358221220af82c4559b7c1104fd6cd303a4fe372fe7f596252ec82770e4271b7b1154a69e64736f6c63430007060033",
}

// PuniswapABI is the input ABI used to generate the binding from.
// Deprecated: Use PuniswapMetaData.ABI instead.
var PuniswapABI = PuniswapMetaData.ABI

// PuniswapBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PuniswapMetaData.Bin instead.
var PuniswapBin = PuniswapMetaData.Bin

// DeployPuniswap deploys a new Ethereum contract, binding an instance of Puniswap to it.
func DeployPuniswap(auth *bind.TransactOpts, backend bind.ContractBackend, _swaproute02 common.Address) (common.Address, *types.Transaction, *Puniswap, error) {
	parsed, err := PuniswapMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PuniswapBin), backend, _swaproute02)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Puniswap{PuniswapCaller: PuniswapCaller{contract: contract}, PuniswapTransactor: PuniswapTransactor{contract: contract}, PuniswapFilterer: PuniswapFilterer{contract: contract}}, nil
}

// Puniswap is an auto generated Go binding around an Ethereum contract.
type Puniswap struct {
	PuniswapCaller     // Read-only binding to the contract
	PuniswapTransactor // Write-only binding to the contract
	PuniswapFilterer   // Log filterer for contract events
}

// PuniswapCaller is an auto generated read-only Go binding around an Ethereum contract.
type PuniswapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PuniswapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PuniswapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PuniswapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PuniswapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PuniswapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PuniswapSession struct {
	Contract     *Puniswap         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PuniswapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PuniswapCallerSession struct {
	Contract *PuniswapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PuniswapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PuniswapTransactorSession struct {
	Contract     *PuniswapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PuniswapRaw is an auto generated low-level Go binding around an Ethereum contract.
type PuniswapRaw struct {
	Contract *Puniswap // Generic contract binding to access the raw methods on
}

// PuniswapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PuniswapCallerRaw struct {
	Contract *PuniswapCaller // Generic read-only contract binding to access the raw methods on
}

// PuniswapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PuniswapTransactorRaw struct {
	Contract *PuniswapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPuniswap creates a new instance of Puniswap, bound to a specific deployed contract.
func NewPuniswap(address common.Address, backend bind.ContractBackend) (*Puniswap, error) {
	contract, err := bindPuniswap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Puniswap{PuniswapCaller: PuniswapCaller{contract: contract}, PuniswapTransactor: PuniswapTransactor{contract: contract}, PuniswapFilterer: PuniswapFilterer{contract: contract}}, nil
}

// NewPuniswapCaller creates a new read-only instance of Puniswap, bound to a specific deployed contract.
func NewPuniswapCaller(address common.Address, caller bind.ContractCaller) (*PuniswapCaller, error) {
	contract, err := bindPuniswap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PuniswapCaller{contract: contract}, nil
}

// NewPuniswapTransactor creates a new write-only instance of Puniswap, bound to a specific deployed contract.
func NewPuniswapTransactor(address common.Address, transactor bind.ContractTransactor) (*PuniswapTransactor, error) {
	contract, err := bindPuniswap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PuniswapTransactor{contract: contract}, nil
}

// NewPuniswapFilterer creates a new log filterer instance of Puniswap, bound to a specific deployed contract.
func NewPuniswapFilterer(address common.Address, filterer bind.ContractFilterer) (*PuniswapFilterer, error) {
	contract, err := bindPuniswap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PuniswapFilterer{contract: contract}, nil
}

// bindPuniswap binds a generic wrapper to an already deployed contract.
func bindPuniswap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PuniswapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Puniswap *PuniswapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Puniswap.Contract.PuniswapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Puniswap *PuniswapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Puniswap.Contract.PuniswapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Puniswap *PuniswapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Puniswap.Contract.PuniswapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Puniswap *PuniswapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Puniswap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Puniswap *PuniswapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Puniswap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Puniswap *PuniswapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Puniswap.Contract.contract.Transact(opts, method, params...)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() view returns(address)
func (_Puniswap *PuniswapCaller) ETHCONTRACTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Puniswap.contract.Call(opts, &out, "ETH_CONTRACT_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() view returns(address)
func (_Puniswap *PuniswapSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _Puniswap.Contract.ETHCONTRACTADDRESS(&_Puniswap.CallOpts)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() view returns(address)
func (_Puniswap *PuniswapCallerSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _Puniswap.Contract.ETHCONTRACTADDRESS(&_Puniswap.CallOpts)
}

// MAX is a free data retrieval call binding the contract method 0xd49d5181.
//
// Solidity: function MAX() view returns(uint256)
func (_Puniswap *PuniswapCaller) MAX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Puniswap.contract.Call(opts, &out, "MAX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAX is a free data retrieval call binding the contract method 0xd49d5181.
//
// Solidity: function MAX() view returns(uint256)
func (_Puniswap *PuniswapSession) MAX() (*big.Int, error) {
	return _Puniswap.Contract.MAX(&_Puniswap.CallOpts)
}

// MAX is a free data retrieval call binding the contract method 0xd49d5181.
//
// Solidity: function MAX() view returns(uint256)
func (_Puniswap *PuniswapCallerSession) MAX() (*big.Int, error) {
	return _Puniswap.Contract.MAX(&_Puniswap.CallOpts)
}

// Swaprouter02 is a free data retrieval call binding the contract method 0x6b150c3c.
//
// Solidity: function swaprouter02() view returns(address)
func (_Puniswap *PuniswapCaller) Swaprouter02(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Puniswap.contract.Call(opts, &out, "swaprouter02")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Swaprouter02 is a free data retrieval call binding the contract method 0x6b150c3c.
//
// Solidity: function swaprouter02() view returns(address)
func (_Puniswap *PuniswapSession) Swaprouter02() (common.Address, error) {
	return _Puniswap.Contract.Swaprouter02(&_Puniswap.CallOpts)
}

// Swaprouter02 is a free data retrieval call binding the contract method 0x6b150c3c.
//
// Solidity: function swaprouter02() view returns(address)
func (_Puniswap *PuniswapCallerSession) Swaprouter02() (common.Address, error) {
	return _Puniswap.Contract.Swaprouter02(&_Puniswap.CallOpts)
}

// Wmatic is a free data retrieval call binding the contract method 0xfb41be16.
//
// Solidity: function wmatic() view returns(address)
func (_Puniswap *PuniswapCaller) Wmatic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Puniswap.contract.Call(opts, &out, "wmatic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wmatic is a free data retrieval call binding the contract method 0xfb41be16.
//
// Solidity: function wmatic() view returns(address)
func (_Puniswap *PuniswapSession) Wmatic() (common.Address, error) {
	return _Puniswap.Contract.Wmatic(&_Puniswap.CallOpts)
}

// Wmatic is a free data retrieval call binding the contract method 0xfb41be16.
//
// Solidity: function wmatic() view returns(address)
func (_Puniswap *PuniswapCallerSession) Wmatic() (common.Address, error) {
	return _Puniswap.Contract.Wmatic(&_Puniswap.CallOpts)
}

// TradeInput is a paid mutator transaction binding the contract method 0xc8dc75e6.
//
// Solidity: function tradeInput((bytes,address,uint256,uint256) params, bool isNative) payable returns(address, uint256)
func (_Puniswap *PuniswapTransactor) TradeInput(opts *bind.TransactOpts, params ISwapRouter2ExactInputParams, isNative bool) (*types.Transaction, error) {
	return _Puniswap.contract.Transact(opts, "tradeInput", params, isNative)
}

// TradeInput is a paid mutator transaction binding the contract method 0xc8dc75e6.
//
// Solidity: function tradeInput((bytes,address,uint256,uint256) params, bool isNative) payable returns(address, uint256)
func (_Puniswap *PuniswapSession) TradeInput(params ISwapRouter2ExactInputParams, isNative bool) (*types.Transaction, error) {
	return _Puniswap.Contract.TradeInput(&_Puniswap.TransactOpts, params, isNative)
}

// TradeInput is a paid mutator transaction binding the contract method 0xc8dc75e6.
//
// Solidity: function tradeInput((bytes,address,uint256,uint256) params, bool isNative) payable returns(address, uint256)
func (_Puniswap *PuniswapTransactorSession) TradeInput(params ISwapRouter2ExactInputParams, isNative bool) (*types.Transaction, error) {
	return _Puniswap.Contract.TradeInput(&_Puniswap.TransactOpts, params, isNative)
}

// TradeInputSingle is a paid mutator transaction binding the contract method 0x421f4388.
//
// Solidity: function tradeInputSingle((address,address,uint24,address,uint256,uint256,uint160) params, bool isNative) payable returns(address, uint256)
func (_Puniswap *PuniswapTransactor) TradeInputSingle(opts *bind.TransactOpts, params ISwapRouter2ExactInputSingleParams, isNative bool) (*types.Transaction, error) {
	return _Puniswap.contract.Transact(opts, "tradeInputSingle", params, isNative)
}

// TradeInputSingle is a paid mutator transaction binding the contract method 0x421f4388.
//
// Solidity: function tradeInputSingle((address,address,uint24,address,uint256,uint256,uint160) params, bool isNative) payable returns(address, uint256)
func (_Puniswap *PuniswapSession) TradeInputSingle(params ISwapRouter2ExactInputSingleParams, isNative bool) (*types.Transaction, error) {
	return _Puniswap.Contract.TradeInputSingle(&_Puniswap.TransactOpts, params, isNative)
}

// TradeInputSingle is a paid mutator transaction binding the contract method 0x421f4388.
//
// Solidity: function tradeInputSingle((address,address,uint24,address,uint256,uint256,uint160) params, bool isNative) payable returns(address, uint256)
func (_Puniswap *PuniswapTransactorSession) TradeInputSingle(params ISwapRouter2ExactInputSingleParams, isNative bool) (*types.Transaction, error) {
	return _Puniswap.Contract.TradeInputSingle(&_Puniswap.TransactOpts, params, isNative)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Puniswap *PuniswapTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Puniswap.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Puniswap *PuniswapSession) Receive() (*types.Transaction, error) {
	return _Puniswap.Contract.Receive(&_Puniswap.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Puniswap *PuniswapTransactorSession) Receive() (*types.Transaction, error) {
	return _Puniswap.Contract.Receive(&_Puniswap.TransactOpts)
}
