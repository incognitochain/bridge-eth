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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"sellToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sellAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isNative\",\"type\":\"bool\"}],\"name\":\"multiTrades\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"}],\"internalType\":\"structISwapRouter2.ExactInputParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"isNative\",\"type\":\"bool\"}],\"name\":\"tradeInput\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structISwapRouter2.ExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"isNative\",\"type\":\"bool\"}],\"name\":\"tradeInputSingle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISwapRouter2\",\"name\":\"_swaproute02\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"ETH_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swaprouter02\",\"outputs\":[{\"internalType\":\"contractISwapRouter2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wmatic\",\"outputs\":[{\"internalType\":\"contractWmatic\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405260405162001db738038062001db78339818101604052810190620000299190620001ae565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634aa4a4fc6040518163ffffffff1660e01b8152600401602060405180830381600087803b158015620000d257600080fd5b505af1158015620000e7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200010d919062000182565b600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505062000256565b600081519050620001658162000222565b92915050565b6000815190506200017c816200023c565b92915050565b6000602082840312156200019557600080fd5b6000620001a58482850162000154565b91505092915050565b600060208284031215620001c157600080fd5b6000620001d1848285016200016b565b91505092915050565b6000620001e78262000202565b9050919050565b6000620001fb82620001da565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200022d81620001da565b81146200023957600080fd5b50565b6200024781620001ee565b81146200025357600080fd5b50565b611b5180620002666000396000f3fe6080604052600436106100745760003560e01c806392171fd81161004e57806392171fd814610107578063c8dc75e614610138578063d49d518114610169578063fb41be16146101945761007b565b8063421f4388146100805780636b150c3c146100b157806372e94bf6146100dc5761007b565b3661007b57005b600080fd5b61009a600480360381019061009591906111ec565b6101bf565b6040516100a8929190611660565b60405180910390f35b3480156100bd57600080fd5b506100c661030a565b6040516100d39190611689565b60405180910390f35b3480156100e857600080fd5b506100f161032e565b6040516100fe91906115f3565b60405180910390f35b610121600480360381019061011c9190611252565b610333565b60405161012f929190611660565b60405180910390f35b610152600480360381019061014d9190611198565b610463565b604051610160929190611660565b60405180910390f35b34801561017557600080fd5b5061017e610645565b60405161018b919061171c565b60405180910390f35b3480156101a057600080fd5b506101a9610669565b6040516101b691906116a4565b60405180910390f35b6000806101e28460000160208101906101d8919061112e565b856080013561068f565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166304e45aaf34876040518363ffffffff1660e01b815260040161023f9190611701565b6020604051808303818588803b15801561025857600080fd5b505af115801561026c573d6000803e3d6000fd5b50505050506040513d601f19601f820116820180604052508101906102919190611229565b90508460a001358110156102da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102d1906116bf565b60405180910390fd5b60006102f98660200160208101906102f2919061112e565b83876107ff565b905080829350935050509250929050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600081565b600080610340868561068f565b60008060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635ae401dc348d8d8d6040518563ffffffff1660e01b81526004016103a193929190611737565b6000604051808303818588803b1580156103ba57600080fd5b505af11580156103ce573d6000803e3d6000fd5b50505050506040513d6000823e3d601f19601f820116820180604052508101906103f89190611157565b905060005b815181101561043e5781818151811061041257fe5b602002602001015180602001905181019061042d9190611229565b8301925080806001019150506103fd565b50600061044c8884886107ff565b905080839450945050505097509795505050505050565b60008060006104c385806000019061047b9190611769565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505061090b565b505090506104d581866040013561068f565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b858183f34886040518363ffffffff1660e01b815260040161053291906116df565b6020604051808303818588803b15801561054b57600080fd5b505af115801561055f573d6000803e3d6000fd5b50505050506040513d601f19601f820116820180604052508101906105849190611229565b905060008680600001906105989190611769565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050905060005b6001156106275760006105f28361095c565b9050801561060a5761060383610977565b9250610621565b6106138361090b565b909150508092505050610627565b506105e0565b6106328184896107ff565b9050808395509550505050509250929050565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000341480156107485750808273ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e3060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518363ffffffff1660e01b81526004016106f692919061160e565b60206040518083038186803b15801561070e57600080fd5b505afa158015610722573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107469190611229565b105b156107fb578173ffffffffffffffffffffffffffffffffffffffff1663095ea7b360008054906101000a900473ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6040518363ffffffff1660e01b81526004016107c8929190611660565b600060405180830381600087803b1580156107e257600080fd5b505af11580156107f6573d6000803e3d6000fd5b505050505b5050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614801561085b5750815b1561090057600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632e1a7d4d846040518263ffffffff1660e01b81526004016108bb919061171c565b600060405180830381600087803b1580156108d557600080fd5b505af11580156108e9573d6000803e3d6000fd5b50505050600090506108fb81846109a0565b610904565b8390505b9392505050565b6000806000610924600085610ae190919063ffffffff16565b925061093a601485610bfa90919063ffffffff16565b9050610953600360140185610ae190919063ffffffff16565b91509193909250565b60006003601401601460036014010101825110159050919050565b60606109996003601401600360140184510384610d049092919063ffffffff16565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610a5e57804710156109e257600080fd5b60003373ffffffffffffffffffffffffffffffffffffffff1682604051610a08906115de565b60006040518083038185875af1925050503d8060008114610a45576040519150601f19603f3d011682016040523d82523d6000602084013e610a4a565b606091505b5050905080610a5857600080fd5b50610add565b8173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b8152600401610a99929190611637565b600060405180830381600087803b158015610ab357600080fd5b505af1158015610ac7573d6000803e3d6000fd5b50505050610ad3610eee565b610adc57600080fd5b5b5050565b600081601483011015610b5c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f746f416464726573735f6f766572666c6f77000000000000000000000000000081525060200191505060405180910390fd5b6014820183511015610bd6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f746f416464726573735f6f75744f66426f756e6473000000000000000000000081525060200191505060405180910390fd5b60006c01000000000000000000000000836020860101510490508091505092915050565b600081600383011015610c75576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f746f55696e7432345f6f766572666c6f7700000000000000000000000000000081525060200191505060405180910390fd5b6003820183511015610cef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f746f55696e7432345f6f75744f66426f756e647300000000000000000000000081525060200191505060405180910390fd5b60008260038501015190508091505092915050565b606081601f83011015610d7f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f736c6963655f6f766572666c6f7700000000000000000000000000000000000081525060200191505060405180910390fd5b828284011015610df7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f736c6963655f6f766572666c6f7700000000000000000000000000000000000081525060200191505060405180910390fd5b81830184511015610e70576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f736c6963655f6f75744f66426f756e647300000000000000000000000000000081525060200191505060405180910390fd5b6060821560008114610e915760405191506000825260208201604052610ee2565b6040519150601f8416801560200281840101858101878315602002848b0101015b81831015610ecf5780518352602083019250602081019050610eb2565b50868552601f19601f8301166040525050505b50809150509392505050565b600080600090503d60008114610f0b5760208114610f1457610f20565b60019150610f20565b60206000803e60005191505b50600081141591505090565b6000610f3f610f3a846117f1565b6117c0565b9050808382526020820190508260005b85811015610f7f5781518501610f658882611065565b845260208401935060208301925050600181019050610f4f565b5050509392505050565b6000610f9c610f978461181d565b6117c0565b905082815260208101848484011115610fb457600080fd5b610fbf848285611a4b565b509392505050565b600081359050610fd681611a91565b92915050565b60008083601f840112610fee57600080fd5b8235905067ffffffffffffffff81111561100757600080fd5b60208301915083602082028301111561101f57600080fd5b9250929050565b600082601f83011261103757600080fd5b8151611047848260208601610f2c565b91505092915050565b60008135905061105f81611aa8565b92915050565b600082601f83011261107657600080fd5b8151611086848260208601610f89565b91505092915050565b60008135905061109e81611abf565b92915050565b6000608082840312156110b657600080fd5b81905092915050565b600060e082840312156110d157600080fd5b81905092915050565b6000813590506110e981611ad6565b92915050565b6000813590506110fe81611aed565b92915050565b60008135905061111381611b04565b92915050565b60008151905061112881611b04565b92915050565b60006020828403121561114057600080fd5b600061114e84828501610fc7565b91505092915050565b60006020828403121561116957600080fd5b600082015167ffffffffffffffff81111561118357600080fd5b61118f84828501611026565b91505092915050565b600080604083850312156111ab57600080fd5b600083013567ffffffffffffffff8111156111c557600080fd5b6111d1858286016110a4565b92505060206111e285828601611050565b9150509250929050565b600080610100838503121561120057600080fd5b600061120e858286016110bf565b92505060e061121f85828601611050565b9150509250929050565b60006020828403121561123b57600080fd5b600061124984828501611119565b91505092915050565b600080600080600080600060c0888a03121561126d57600080fd5b600061127b8a828b01611104565b975050602088013567ffffffffffffffff81111561129857600080fd5b6112a48a828b01610fdc565b965096505060406112b78a828b0161108f565b94505060606112c88a828b01610fc7565b93505060806112d98a828b01611104565b92505060a06112ea8a828b01611050565b91505092959891949750929550565b60006113068484846113b2565b90509392505050565b611318816119be565b82525050565b61132781611955565b82525050565b61133681611955565b82525050565b60006113488385611864565b93508360208402850161135a8461184d565b8060005b878110156113a057848403895261137582846118b9565b6113808682846112f9565b955061138b84611857565b935060208b019a50505060018101905061135e565b50829750879450505050509392505050565b60006113be8385611875565b93506113cb838584611a3c565b6113d483611a80565b840190509392505050565b6113e8816119d0565b82525050565b6113f7816119f4565b82525050565b600061140a601a83611891565b91507f6c6f776572207468616e206578706563746564206f75747075740000000000006000830152602082019050919050565b600061144a600083611886565b9150600082019050919050565b60006080830161146a60008401846118b9565b858303600087015261147d8382846113b2565b9250505061148e60208401846118a2565b61149b602086018261131e565b506114a9604084018461193e565b6114b660408601826115c0565b506114c4606084018461193e565b6114d160608601826115c0565b508091505092915050565b60e082016114ed60008301836118a2565b6114fa600085018261131e565b5061150860208301836118a2565b611515602085018261131e565b506115236040830183611927565b61153060408501826115b1565b5061153e60608301836118a2565b61154b606085018261131e565b50611559608083018361193e565b61156660808501826115c0565b5061157460a083018361193e565b61158160a08501826115c0565b5061158f60c0830183611910565b61159c60c08501826115a2565b50505050565b6115ab81611985565b82525050565b6115ba816119a5565b82525050565b6115c9816119b4565b82525050565b6115d8816119b4565b82525050565b60006115e98261143d565b9150819050919050565b6000602082019050611608600083018461132d565b92915050565b6000604082019050611623600083018561130f565b611630602083018461132d565b9392505050565b600060408201905061164c600083018561130f565b61165960208301846115cf565b9392505050565b6000604082019050611675600083018561132d565b61168260208301846115cf565b9392505050565b600060208201905061169e60008301846113df565b92915050565b60006020820190506116b960008301846113ee565b92915050565b600060208201905081810360008301526116d8816113fd565b9050919050565b600060208201905081810360008301526116f98184611457565b905092915050565b600060e08201905061171660008301846114dc565b92915050565b600060208201905061173160008301846115cf565b92915050565b600060408201905061174c60008301866115cf565b818103602083015261175f81848661133c565b9050949350505050565b6000808335600160200384360303811261178257600080fd5b80840192508235915067ffffffffffffffff8211156117a057600080fd5b6020830192506001820236038313156117b857600080fd5b509250929050565b6000604051905081810181811067ffffffffffffffff821117156117e7576117e6611a7e565b5b8060405250919050565b600067ffffffffffffffff82111561180c5761180b611a7e565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561183857611837611a7e565b5b601f19601f8301169050602081019050919050565b6000819050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b60006118b16020840184610fc7565b905092915050565b600080833560016020038436030381126118d257600080fd5b83810192508235915060208301925067ffffffffffffffff8211156118f657600080fd5b60018202360384131561190857600080fd5b509250929050565b600061191f60208401846110da565b905092915050565b600061193660208401846110ef565b905092915050565b600061194d6020840184611104565b905092915050565b600061196082611985565b9050919050565b60008115159050919050565b600061197e82611955565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600062ffffff82169050919050565b6000819050919050565b60006119c982611a18565b9050919050565b60006119db826119e2565b9050919050565b60006119ed82611985565b9050919050565b60006119ff82611a06565b9050919050565b6000611a1182611985565b9050919050565b6000611a2382611a2a565b9050919050565b6000611a3582611985565b9050919050565b82818337600083830152505050565b60005b83811015611a69578082015181840152602081019050611a4e565b83811115611a78576000848401525b50505050565bfe5b6000601f19601f8301169050919050565b611a9a81611955565b8114611aa557600080fd5b50565b611ab181611967565b8114611abc57600080fd5b50565b611ac881611973565b8114611ad357600080fd5b50565b611adf81611985565b8114611aea57600080fd5b50565b611af6816119a5565b8114611b0157600080fd5b50565b611b0d816119b4565b8114611b1857600080fd5b5056fea2646970667358221220c1296113f5dc327f301aa83501dbb1c0f7f62925f91879612ed97a3d75e8bd3964736f6c63430007060033",
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

// MultiTrades is a paid mutator transaction binding the contract method 0x92171fd8.
//
// Solidity: function multiTrades(uint256 deadline, bytes[] data, address sellToken, address buyToken, uint256 sellAmount, bool isNative) payable returns(address, uint256)
func (_Puniswap *PuniswapTransactor) MultiTrades(opts *bind.TransactOpts, deadline *big.Int, data [][]byte, sellToken common.Address, buyToken common.Address, sellAmount *big.Int, isNative bool) (*types.Transaction, error) {
	return _Puniswap.contract.Transact(opts, "multiTrades", deadline, data, sellToken, buyToken, sellAmount, isNative)
}

// MultiTrades is a paid mutator transaction binding the contract method 0x92171fd8.
//
// Solidity: function multiTrades(uint256 deadline, bytes[] data, address sellToken, address buyToken, uint256 sellAmount, bool isNative) payable returns(address, uint256)
func (_Puniswap *PuniswapSession) MultiTrades(deadline *big.Int, data [][]byte, sellToken common.Address, buyToken common.Address, sellAmount *big.Int, isNative bool) (*types.Transaction, error) {
	return _Puniswap.Contract.MultiTrades(&_Puniswap.TransactOpts, deadline, data, sellToken, buyToken, sellAmount, isNative)
}

// MultiTrades is a paid mutator transaction binding the contract method 0x92171fd8.
//
// Solidity: function multiTrades(uint256 deadline, bytes[] data, address sellToken, address buyToken, uint256 sellAmount, bool isNative) payable returns(address, uint256)
func (_Puniswap *PuniswapTransactorSession) MultiTrades(deadline *big.Int, data [][]byte, sellToken common.Address, buyToken common.Address, sellAmount *big.Int, isNative bool) (*types.Transaction, error) {
	return _Puniswap.Contract.MultiTrades(&_Puniswap.TransactOpts, deadline, data, sellToken, buyToken, sellAmount, isNative)
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
