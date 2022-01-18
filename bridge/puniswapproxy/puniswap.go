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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"sellToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyToken\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"amount\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isNative\",\"type\":\"bool\"}],\"name\":\"multiTrades\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"}],\"internalType\":\"structISwapRouter2.ExactInputParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"isNative\",\"type\":\"bool\"}],\"name\":\"tradeInput\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structISwapRouter2.ExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"isNative\",\"type\":\"bool\"}],\"name\":\"tradeInputSingle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISwapRouter2\",\"name\":\"_swaproute02\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"ETH_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swaprouter02\",\"outputs\":[{\"internalType\":\"contractISwapRouter2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wmatic\",\"outputs\":[{\"internalType\":\"contractWmatic\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405260405162001df938038062001df98339818101604052810190620000299190620001ae565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634aa4a4fc6040518163ffffffff1660e01b8152600401602060405180830381600087803b158015620000d257600080fd5b505af1158015620000e7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200010d919062000182565b600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505062000256565b600081519050620001658162000222565b92915050565b6000815190506200017c816200023c565b92915050565b6000602082840312156200019557600080fd5b6000620001a58482850162000154565b91505092915050565b600060208284031215620001c157600080fd5b6000620001d1848285016200016b565b91505092915050565b6000620001e78262000202565b9050919050565b6000620001fb82620001da565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200022d81620001da565b81146200023957600080fd5b50565b6200024781620001ee565b81146200025357600080fd5b50565b611b9380620002666000396000f3fe6080604052600436106100745760003560e01c806372e94bf61161004e57806372e94bf61461010d578063c8dc75e614610138578063d49d518114610169578063fb41be16146101945761007b565b806306229c8014610080578063421f4388146100b15780636b150c3c146100e25761007b565b3661007b57005b600080fd5b61009a6004803603810190610095919061126d565b6101bf565b6040516100a892919061167b565b60405180910390f35b6100cb60048036038101906100c69190611207565b6102f5565b6040516100d992919061167b565b60405180910390f35b3480156100ee57600080fd5b506100f7610440565b60405161010491906116a4565b60405180910390f35b34801561011957600080fd5b50610122610464565b60405161012f919061160e565b60405180910390f35b610152600480360381019061014d91906111b3565b610469565b60405161016092919061167b565b60405180910390f35b34801561017557600080fd5b5061017e61064b565b60405161018b9190611737565b60405180910390f35b3480156101a057600080fd5b506101a961066f565b6040516101b691906116bf565b60405180910390f35b6000806101d2868563ffffffff16610695565b60008060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635ae401dc348d8d8d6040518563ffffffff1660e01b815260040161023393929190611752565b6000604051808303818588803b15801561024c57600080fd5b505af1158015610260573d6000803e3d6000fd5b50505050506040513d6000823e3d601f19601f8201168201806040525081019061028a9190611172565b905060005b81518110156102d0578181815181106102a457fe5b60200260200101518060200190518101906102bf9190611244565b83019250808060010191505061028f565b5060006102de888488610805565b905080839450945050505097509795505050505050565b60008061031884600001602081019061030e9190611149565b8560800135610695565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166304e45aaf34876040518363ffffffff1660e01b8152600401610375919061171c565b6020604051808303818588803b15801561038e57600080fd5b505af11580156103a2573d6000803e3d6000fd5b50505050506040513d601f19601f820116820180604052508101906103c79190611244565b90508460a00135811015610410576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610407906116da565b60405180910390fd5b600061042f8660200160208101906104289190611149565b8387610805565b905080829350935050509250929050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600081565b60008060006104c98580600001906104819190611784565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050610911565b505090506104db818660400135610695565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b858183f34886040518363ffffffff1660e01b815260040161053891906116fa565b6020604051808303818588803b15801561055157600080fd5b505af1158015610565573d6000803e3d6000fd5b50505050506040513d601f19601f8201168201806040525081019061058a9190611244565b9050600086806000019061059e9190611784565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050905060005b60011561062d5760006105f883610962565b90508015610610576106098361097d565b9250610627565b61061983610911565b90915050809250505061062d565b506105e6565b610638818489610805565b9050808395509550505050509250929050565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60003414801561074e5750808273ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e3060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518363ffffffff1660e01b81526004016106fc929190611629565b60206040518083038186803b15801561071457600080fd5b505afa158015610728573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061074c9190611244565b105b15610801578173ffffffffffffffffffffffffffffffffffffffff1663095ea7b360008054906101000a900473ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6040518363ffffffff1660e01b81526004016107ce92919061167b565b600060405180830381600087803b1580156107e857600080fd5b505af11580156107fc573d6000803e3d6000fd5b505050505b5050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161480156108615750815b1561090657600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632e1a7d4d846040518263ffffffff1660e01b81526004016108c19190611737565b600060405180830381600087803b1580156108db57600080fd5b505af11580156108ef573d6000803e3d6000fd5b505050506000905061090181846109a6565b61090a565b8390505b9392505050565b600080600061092a600085610ae790919063ffffffff16565b9250610940601485610c0090919063ffffffff16565b9050610959600360140185610ae790919063ffffffff16565b91509193909250565b60006003601401601460036014010101825110159050919050565b606061099f6003601401600360140184510384610d0a9092919063ffffffff16565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610a6457804710156109e857600080fd5b60003373ffffffffffffffffffffffffffffffffffffffff1682604051610a0e906115f9565b60006040518083038185875af1925050503d8060008114610a4b576040519150601f19603f3d011682016040523d82523d6000602084013e610a50565b606091505b5050905080610a5e57600080fd5b50610ae3565b8173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b8152600401610a9f929190611652565b600060405180830381600087803b158015610ab957600080fd5b505af1158015610acd573d6000803e3d6000fd5b50505050610ad9610ef4565b610ae257600080fd5b5b5050565b600081601483011015610b62576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f746f416464726573735f6f766572666c6f77000000000000000000000000000081525060200191505060405180910390fd5b6014820183511015610bdc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f746f416464726573735f6f75744f66426f756e6473000000000000000000000081525060200191505060405180910390fd5b60006c01000000000000000000000000836020860101510490508091505092915050565b600081600383011015610c7b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f746f55696e7432345f6f766572666c6f7700000000000000000000000000000081525060200191505060405180910390fd5b6003820183511015610cf5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f746f55696e7432345f6f75744f66426f756e647300000000000000000000000081525060200191505060405180910390fd5b60008260038501015190508091505092915050565b606081601f83011015610d85576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f736c6963655f6f766572666c6f7700000000000000000000000000000000000081525060200191505060405180910390fd5b828284011015610dfd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f736c6963655f6f766572666c6f7700000000000000000000000000000000000081525060200191505060405180910390fd5b81830184511015610e76576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f736c6963655f6f75744f66426f756e647300000000000000000000000000000081525060200191505060405180910390fd5b6060821560008114610e975760405191506000825260208201604052610ee8565b6040519150601f8416801560200281840101858101878315602002848b0101015b81831015610ed55780518352602083019250602081019050610eb8565b50868552601f19601f8301166040525050505b50809150509392505050565b600080600090503d60008114610f115760208114610f1a57610f26565b60019150610f26565b60206000803e60005191505b50600081141591505090565b6000610f45610f408461180c565b6117db565b9050808382526020820190508260005b85811015610f855781518501610f6b888261106b565b845260208401935060208301925050600181019050610f55565b5050509392505050565b6000610fa2610f9d84611838565b6117db565b905082815260208101848484011115610fba57600080fd5b610fc5848285611a76565b509392505050565b600081359050610fdc81611abc565b92915050565b60008083601f840112610ff457600080fd5b8235905067ffffffffffffffff81111561100d57600080fd5b60208301915083602082028301111561102557600080fd5b9250929050565b600082601f83011261103d57600080fd5b815161104d848260208601610f32565b91505092915050565b60008135905061106581611ad3565b92915050565b600082601f83011261107c57600080fd5b815161108c848260208601610f8f565b91505092915050565b6000813590506110a481611aea565b92915050565b6000608082840312156110bc57600080fd5b81905092915050565b600060e082840312156110d757600080fd5b81905092915050565b6000813590506110ef81611b01565b92915050565b60008135905061110481611b18565b92915050565b60008135905061111981611b2f565b92915050565b60008151905061112e81611b2f565b92915050565b60008135905061114381611b46565b92915050565b60006020828403121561115b57600080fd5b600061116984828501610fcd565b91505092915050565b60006020828403121561118457600080fd5b600082015167ffffffffffffffff81111561119e57600080fd5b6111aa8482850161102c565b91505092915050565b600080604083850312156111c657600080fd5b600083013567ffffffffffffffff8111156111e057600080fd5b6111ec858286016110aa565b92505060206111fd85828601611056565b9150509250929050565b600080610100838503121561121b57600080fd5b6000611229858286016110c5565b92505060e061123a85828601611056565b9150509250929050565b60006020828403121561125657600080fd5b60006112648482850161111f565b91505092915050565b600080600080600080600060c0888a03121561128857600080fd5b60006112968a828b0161110a565b975050602088013567ffffffffffffffff8111156112b357600080fd5b6112bf8a828b01610fe2565b965096505060406112d28a828b01611095565b94505060606112e38a828b01610fcd565b93505060806112f48a828b01611134565b92505060a06113058a828b01611056565b91505092959891949750929550565b60006113218484846113cd565b90509392505050565b611333816119e9565b82525050565b61134281611970565b82525050565b61135181611970565b82525050565b6000611363838561187f565b93508360208402850161137584611868565b8060005b878110156113bb57848403895261139082846118d4565b61139b868284611314565b95506113a684611872565b935060208b019a505050600181019050611379565b50829750879450505050509392505050565b60006113d98385611890565b93506113e6838584611a67565b6113ef83611aab565b840190509392505050565b611403816119fb565b82525050565b61141281611a1f565b82525050565b6000611425601a836118ac565b91507f6c6f776572207468616e206578706563746564206f75747075740000000000006000830152602082019050919050565b60006114656000836118a1565b9150600082019050919050565b60006080830161148560008401846118d4565b85830360008701526114988382846113cd565b925050506114a960208401846118bd565b6114b66020860182611339565b506114c46040840184611959565b6114d160408601826115db565b506114df6060840184611959565b6114ec60608601826115db565b508091505092915050565b60e0820161150860008301836118bd565b6115156000850182611339565b5061152360208301836118bd565b6115306020850182611339565b5061153e6040830183611942565b61154b60408501826115cc565b5061155960608301836118bd565b6115666060850182611339565b506115746080830183611959565b61158160808501826115db565b5061158f60a0830183611959565b61159c60a08501826115db565b506115aa60c083018361192b565b6115b760c08501826115bd565b50505050565b6115c6816119a0565b82525050565b6115d5816119c0565b82525050565b6115e4816119cf565b82525050565b6115f3816119cf565b82525050565b600061160482611458565b9150819050919050565b60006020820190506116236000830184611348565b92915050565b600060408201905061163e600083018561132a565b61164b6020830184611348565b9392505050565b6000604082019050611667600083018561132a565b61167460208301846115ea565b9392505050565b60006040820190506116906000830185611348565b61169d60208301846115ea565b9392505050565b60006020820190506116b960008301846113fa565b92915050565b60006020820190506116d46000830184611409565b92915050565b600060208201905081810360008301526116f381611418565b9050919050565b600060208201905081810360008301526117148184611472565b905092915050565b600060e08201905061173160008301846114f7565b92915050565b600060208201905061174c60008301846115ea565b92915050565b600060408201905061176760008301866115ea565b818103602083015261177a818486611357565b9050949350505050565b6000808335600160200384360303811261179d57600080fd5b80840192508235915067ffffffffffffffff8211156117bb57600080fd5b6020830192506001820236038313156117d357600080fd5b509250929050565b6000604051905081810181811067ffffffffffffffff8211171561180257611801611aa9565b5b8060405250919050565b600067ffffffffffffffff82111561182757611826611aa9565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561185357611852611aa9565b5b601f19601f8301169050602081019050919050565b6000819050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b60006118cc6020840184610fcd565b905092915050565b600080833560016020038436030381126118ed57600080fd5b83810192508235915060208301925067ffffffffffffffff82111561191157600080fd5b60018202360384131561192357600080fd5b509250929050565b600061193a60208401846110e0565b905092915050565b600061195160208401846110f5565b905092915050565b6000611968602084018461110a565b905092915050565b600061197b826119a0565b9050919050565b60008115159050919050565b600061199982611970565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600062ffffff82169050919050565b6000819050919050565b600063ffffffff82169050919050565b60006119f482611a43565b9050919050565b6000611a0682611a0d565b9050919050565b6000611a18826119a0565b9050919050565b6000611a2a82611a31565b9050919050565b6000611a3c826119a0565b9050919050565b6000611a4e82611a55565b9050919050565b6000611a60826119a0565b9050919050565b82818337600083830152505050565b60005b83811015611a94578082015181840152602081019050611a79565b83811115611aa3576000848401525b50505050565bfe5b6000601f19601f8301169050919050565b611ac581611970565b8114611ad057600080fd5b50565b611adc81611982565b8114611ae757600080fd5b50565b611af38161198e565b8114611afe57600080fd5b50565b611b0a816119a0565b8114611b1557600080fd5b50565b611b21816119c0565b8114611b2c57600080fd5b50565b611b38816119cf565b8114611b4357600080fd5b50565b611b4f816119d9565b8114611b5a57600080fd5b5056fea26469706673582212209d70fe756993e39074e0bf9fe824d116f0c95d49e975e024d86ea096dd8092ec64736f6c63430007060033",
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

// MultiTrades is a paid mutator transaction binding the contract method 0x06229c80.
//
// Solidity: function multiTrades(uint256 deadline, bytes[] data, address sellToken, address buyToken, uint32 amount, bool isNative) payable returns(address, uint256)
func (_Puniswap *PuniswapTransactor) MultiTrades(opts *bind.TransactOpts, deadline *big.Int, data [][]byte, sellToken common.Address, buyToken common.Address, amount uint32, isNative bool) (*types.Transaction, error) {
	return _Puniswap.contract.Transact(opts, "multiTrades", deadline, data, sellToken, buyToken, amount, isNative)
}

// MultiTrades is a paid mutator transaction binding the contract method 0x06229c80.
//
// Solidity: function multiTrades(uint256 deadline, bytes[] data, address sellToken, address buyToken, uint32 amount, bool isNative) payable returns(address, uint256)
func (_Puniswap *PuniswapSession) MultiTrades(deadline *big.Int, data [][]byte, sellToken common.Address, buyToken common.Address, amount uint32, isNative bool) (*types.Transaction, error) {
	return _Puniswap.Contract.MultiTrades(&_Puniswap.TransactOpts, deadline, data, sellToken, buyToken, amount, isNative)
}

// MultiTrades is a paid mutator transaction binding the contract method 0x06229c80.
//
// Solidity: function multiTrades(uint256 deadline, bytes[] data, address sellToken, address buyToken, uint32 amount, bool isNative) payable returns(address, uint256)
func (_Puniswap *PuniswapTransactorSession) MultiTrades(deadline *big.Int, data [][]byte, sellToken common.Address, buyToken common.Address, amount uint32, isNative bool) (*types.Transaction, error) {
	return _Puniswap.Contract.MultiTrades(&_Puniswap.TransactOpts, deadline, data, sellToken, buyToken, amount, isNative)
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
