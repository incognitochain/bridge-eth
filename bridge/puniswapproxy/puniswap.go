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

// ISwapRouterExactInputParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactInputParams struct {
	Path             []byte
	Recipient        common.Address
	Deadline         *big.Int
	AmountIn         *big.Int
	AmountOutMinimum *big.Int
}

// ISwapRouterExactInputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactInputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	Deadline          *big.Int
	AmountIn          *big.Int
	AmountOutMinimum  *big.Int
	SqrtPriceLimitX96 *big.Int
}

// PuniswapMetaData contains all meta data concerning the Puniswap contract.
var PuniswapMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractISwapRouter2\",\"name\":\"_swaproute02\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ETH_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swaprouter02\",\"outputs\":[{\"internalType\":\"contractISwapRouter2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"}],\"internalType\":\"structISwapRouter.ExactInputParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"isNative\",\"type\":\"bool\"}],\"name\":\"tradeInput\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structISwapRouter.ExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"isNative\",\"type\":\"bool\"}],\"name\":\"tradeInputSingle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wmatic\",\"outputs\":[{\"internalType\":\"contractWmatic\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040526040516200186e3803806200186e8339818101604052810190620000299190620001ae565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634aa4a4fc6040518163ffffffff1660e01b8152600401602060405180830381600087803b158015620000d257600080fd5b505af1158015620000e7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200010d919062000182565b600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505062000256565b600081519050620001658162000222565b92915050565b6000815190506200017c816200023c565b92915050565b6000602082840312156200019557600080fd5b6000620001a58482850162000154565b91505092915050565b600060208284031215620001c157600080fd5b6000620001d1848285016200016b565b91505092915050565b6000620001e78262000202565b9050919050565b6000620001fb82620001da565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200022d81620001da565b81146200023957600080fd5b50565b6200024781620001ee565b81146200025357600080fd5b50565b61160880620002666000396000f3fe6080604052600436106100595760003560e01c80631c63c8cb146100655780636b150c3c1461009657806372e94bf6146100c1578063d49d5181146100ec578063dae580b814610117578063fb41be161461014857610060565b3661006057005b600080fd5b61007f600480360381019061007a9190610ee2565b610173565b60405161008d92919061125b565b60405180910390f35b3480156100a257600080fd5b506100ab6102be565b6040516100b89190611284565b60405180910390f35b3480156100cd57600080fd5b506100d66102e2565b6040516100e391906111ee565b60405180910390f35b3480156100f857600080fd5b506101016102e7565b60405161010e9190611318565b60405180910390f35b610131600480360381019061012c9190610e8e565b61030b565b60405161013f92919061125b565b60405180910390f35b34801561015457600080fd5b5061015d6104ed565b60405161016a919061129f565b60405180910390f35b60008061019684600001602081019061018c9190610e65565b8560a00135610513565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663414bf38934876040518363ffffffff1660e01b81526004016101f391906112fc565b6020604051808303818588803b15801561020c57600080fd5b505af1158015610220573d6000803e3d6000fd5b50505050506040513d601f19601f820116820180604052508101906102459190610f20565b90508460c0013581101561028e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610285906112ba565b60405180910390fd5b60006102ad8660200160208101906102a69190610e65565b8387610683565b905080829350935050509250929050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600081565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81565b600080600061036b8580600001906103239190611333565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505061078f565b5050905061037d818660600135610513565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c04b8d5934886040518363ffffffff1660e01b81526004016103da91906112da565b6020604051808303818588803b1580156103f357600080fd5b505af1158015610407573d6000803e3d6000fd5b50505050506040513d601f19601f8201168201806040525081019061042c9190610f20565b905060008680600001906104409190611333565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050905060005b6001156104cf57600061049a836107e0565b905080156104b2576104ab836107fb565b92506104c9565b6104bb8361078f565b9091505080925050506104cf565b50610488565b6104da818489610683565b9050808395509550505050509250929050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000341480156105cc5750808273ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e3060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518363ffffffff1660e01b815260040161057a929190611209565b60206040518083038186803b15801561059257600080fd5b505afa1580156105a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105ca9190610f20565b105b1561067f578173ffffffffffffffffffffffffffffffffffffffff1663095ea7b360008054906101000a900473ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6040518363ffffffff1660e01b815260040161064c92919061125b565b600060405180830381600087803b15801561066657600080fd5b505af115801561067a573d6000803e3d6000fd5b505050505b5050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161480156106df5750815b1561078457600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632e1a7d4d846040518263ffffffff1660e01b815260040161073f9190611318565b600060405180830381600087803b15801561075957600080fd5b505af115801561076d573d6000803e3d6000fd5b505050506000905061077f8184610824565b610788565b8390505b9392505050565b60008060006107a860008561096590919063ffffffff16565b92506107be601485610a7e90919063ffffffff16565b90506107d760036014018561096590919063ffffffff16565b91509193909250565b60006003601401601460036014010101825110159050919050565b606061081d6003601401600360140184510384610b889092919063ffffffff16565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156108e2578047101561086657600080fd5b60003373ffffffffffffffffffffffffffffffffffffffff168260405161088c906111d9565b60006040518083038185875af1925050503d80600081146108c9576040519150601f19603f3d011682016040523d82523d6000602084013e6108ce565b606091505b50509050806108dc57600080fd5b50610961565b8173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b815260040161091d929190611232565b600060405180830381600087803b15801561093757600080fd5b505af115801561094b573d6000803e3d6000fd5b50505050610957610d72565b61096057600080fd5b5b5050565b6000816014830110156109e0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f746f416464726573735f6f766572666c6f77000000000000000000000000000081525060200191505060405180910390fd5b6014820183511015610a5a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f746f416464726573735f6f75744f66426f756e6473000000000000000000000081525060200191505060405180910390fd5b60006c01000000000000000000000000836020860101510490508091505092915050565b600081600383011015610af9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f746f55696e7432345f6f766572666c6f7700000000000000000000000000000081525060200191505060405180910390fd5b6003820183511015610b73576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f746f55696e7432345f6f75744f66426f756e647300000000000000000000000081525060200191505060405180910390fd5b60008260038501015190508091505092915050565b606081601f83011015610c03576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f736c6963655f6f766572666c6f7700000000000000000000000000000000000081525060200191505060405180910390fd5b828284011015610c7b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f736c6963655f6f766572666c6f7700000000000000000000000000000000000081525060200191505060405180910390fd5b81830184511015610cf4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f736c6963655f6f75744f66426f756e647300000000000000000000000000000081525060200191505060405180910390fd5b6060821560008114610d155760405191506000825260208201604052610d66565b6040519150601f8416801560200281840101858101878315602002848b0101015b81831015610d535780518352602083019250602081019050610d36565b50868552601f19601f8301166040525050505b50809150509392505050565b600080600090503d60008114610d8f5760208114610d9857610da4565b60019150610da4565b60206000803e60005191505b50600081141591505090565b600081359050610dbf8161155f565b92915050565b600081359050610dd481611576565b92915050565b600060a08284031215610dec57600080fd5b81905092915050565b60006101008284031215610e0857600080fd5b81905092915050565b600081359050610e208161158d565b92915050565b600081359050610e35816115a4565b92915050565b600081359050610e4a816115bb565b92915050565b600081519050610e5f816115bb565b92915050565b600060208284031215610e7757600080fd5b6000610e8584828501610db0565b91505092915050565b60008060408385031215610ea157600080fd5b600083013567ffffffffffffffff811115610ebb57600080fd5b610ec785828601610dda565b9250506020610ed885828601610dc5565b9150509250929050565b6000806101208385031215610ef657600080fd5b6000610f0485828601610df5565b925050610100610f1685828601610dc5565b9150509250929050565b600060208284031215610f3257600080fd5b6000610f4084828501610e50565b91505092915050565b610f52816114c1565b82525050565b610f618161146a565b82525050565b610f708161146a565b82525050565b6000610f82838561138a565b9350610f8f83858461153f565b610f988361154e565b840190509392505050565b610fac816114d3565b82525050565b610fbb816114f7565b82525050565b6000610fce601a836113a6565b91507f6c6f776572207468616e206578706563746564206f75747075740000000000006000830152602082019050919050565b600061100e60008361139b565b9150600082019050919050565b600060a0830161102e60008401846113ce565b8583036000870152611041838284610f76565b9250505061105260208401846113b7565b61105f6020860182610f58565b5061106d6040840184611453565b61107a60408601826111bb565b506110886060840184611453565b61109560608601826111bb565b506110a36080840184611453565b6110b060808601826111bb565b508091505092915050565b61010082016110cd60008301836113b7565b6110da6000850182610f58565b506110e860208301836113b7565b6110f56020850182610f58565b50611103604083018361143c565b61111060408501826111ac565b5061111e60608301836113b7565b61112b6060850182610f58565b506111396080830183611453565b61114660808501826111bb565b5061115460a0830183611453565b61116160a08501826111bb565b5061116f60c0830183611453565b61117c60c08501826111bb565b5061118a60e0830183611425565b61119760e085018261119d565b50505050565b6111a681611488565b82525050565b6111b5816114a8565b82525050565b6111c4816114b7565b82525050565b6111d3816114b7565b82525050565b60006111e482611001565b9150819050919050565b60006020820190506112036000830184610f67565b92915050565b600060408201905061121e6000830185610f49565b61122b6020830184610f67565b9392505050565b60006040820190506112476000830185610f49565b61125460208301846111ca565b9392505050565b60006040820190506112706000830185610f67565b61127d60208301846111ca565b9392505050565b60006020820190506112996000830184610fa3565b92915050565b60006020820190506112b46000830184610fb2565b92915050565b600060208201905081810360008301526112d381610fc1565b9050919050565b600060208201905081810360008301526112f4818461101b565b905092915050565b60006101008201905061131260008301846110bb565b92915050565b600060208201905061132d60008301846111ca565b92915050565b6000808335600160200384360303811261134c57600080fd5b80840192508235915067ffffffffffffffff82111561136a57600080fd5b60208301925060018202360383131561138257600080fd5b509250929050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b60006113c66020840184610db0565b905092915050565b600080833560016020038436030381126113e757600080fd5b83810192508235915060208301925067ffffffffffffffff82111561140b57600080fd5b60018202360384131561141d57600080fd5b509250929050565b60006114346020840184610e11565b905092915050565b600061144b6020840184610e26565b905092915050565b60006114626020840184610e3b565b905092915050565b600061147582611488565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600062ffffff82169050919050565b6000819050919050565b60006114cc8261151b565b9050919050565b60006114de826114e5565b9050919050565b60006114f082611488565b9050919050565b600061150282611509565b9050919050565b600061151482611488565b9050919050565b60006115268261152d565b9050919050565b600061153882611488565b9050919050565b82818337600083830152505050565b6000601f19601f8301169050919050565b6115688161146a565b811461157357600080fd5b50565b61157f8161147c565b811461158a57600080fd5b50565b61159681611488565b81146115a157600080fd5b50565b6115ad816114a8565b81146115b857600080fd5b50565b6115c4816114b7565b81146115cf57600080fd5b5056fea264697066735822122073327fe73d784527982862c909ff708ce1c80a79077fe06b7f1a8b1e55bf7c5e64736f6c63430007060033",
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

// TradeInput is a paid mutator transaction binding the contract method 0xdae580b8.
//
// Solidity: function tradeInput((bytes,address,uint256,uint256,uint256) params, bool isNative) payable returns(address, uint256)
func (_Puniswap *PuniswapTransactor) TradeInput(opts *bind.TransactOpts, params ISwapRouterExactInputParams, isNative bool) (*types.Transaction, error) {
	return _Puniswap.contract.Transact(opts, "tradeInput", params, isNative)
}

// TradeInput is a paid mutator transaction binding the contract method 0xdae580b8.
//
// Solidity: function tradeInput((bytes,address,uint256,uint256,uint256) params, bool isNative) payable returns(address, uint256)
func (_Puniswap *PuniswapSession) TradeInput(params ISwapRouterExactInputParams, isNative bool) (*types.Transaction, error) {
	return _Puniswap.Contract.TradeInput(&_Puniswap.TransactOpts, params, isNative)
}

// TradeInput is a paid mutator transaction binding the contract method 0xdae580b8.
//
// Solidity: function tradeInput((bytes,address,uint256,uint256,uint256) params, bool isNative) payable returns(address, uint256)
func (_Puniswap *PuniswapTransactorSession) TradeInput(params ISwapRouterExactInputParams, isNative bool) (*types.Transaction, error) {
	return _Puniswap.Contract.TradeInput(&_Puniswap.TransactOpts, params, isNative)
}

// TradeInputSingle is a paid mutator transaction binding the contract method 0x1c63c8cb.
//
// Solidity: function tradeInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params, bool isNative) payable returns(address, uint256)
func (_Puniswap *PuniswapTransactor) TradeInputSingle(opts *bind.TransactOpts, params ISwapRouterExactInputSingleParams, isNative bool) (*types.Transaction, error) {
	return _Puniswap.contract.Transact(opts, "tradeInputSingle", params, isNative)
}

// TradeInputSingle is a paid mutator transaction binding the contract method 0x1c63c8cb.
//
// Solidity: function tradeInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params, bool isNative) payable returns(address, uint256)
func (_Puniswap *PuniswapSession) TradeInputSingle(params ISwapRouterExactInputSingleParams, isNative bool) (*types.Transaction, error) {
	return _Puniswap.Contract.TradeInputSingle(&_Puniswap.TransactOpts, params, isNative)
}

// TradeInputSingle is a paid mutator transaction binding the contract method 0x1c63c8cb.
//
// Solidity: function tradeInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params, bool isNative) payable returns(address, uint256)
func (_Puniswap *PuniswapTransactorSession) TradeInputSingle(params ISwapRouterExactInputSingleParams, isNative bool) (*types.Transaction, error) {
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
