// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blur

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

// Execution is an auto generated low-level Go binding around an user-defined struct.
type Execution struct {
	Sell Input
	Buy  Input
}

// Fee is an auto generated low-level Go binding around an user-defined struct.
type Fee struct {
	Rate      uint16
	Recipient common.Address
}

// Input is an auto generated low-level Go binding around an user-defined struct.
type Input struct {
	Order            Order
	V                uint8
	R                [32]byte
	S                [32]byte
	ExtraSignature   []byte
	SignatureVersion uint8
	BlockNumber      *big.Int
}

// Order is an auto generated low-level Go binding around an user-defined struct.
type Order struct {
	Trader         common.Address
	Side           uint8
	MatchingPolicy common.Address
	Collection     common.Address
	TokenId        *big.Int
	Amount         *big.Int
	PaymentToken   common.Address
	Price          *big.Int
	ListingTime    *big.Int
	ExpirationTime *big.Int
	Fees           []Fee
	Salt           *big.Int
	ExtraParams    []byte
}

// TokenInfo is an auto generated low-level Go binding around an user-defined struct.
type TokenInfo struct {
	TokenType  uint8
	Collection common.Address
	Identifier *big.Int
	Amount     *big.Int
}

// TradeDetails is an auto generated low-level Go binding around an user-defined struct.
type TradeDetails struct {
	MarketId  *big.Int
	Value     *big.Int
	TradeData []byte
}

// BlurMetaData contains all meta data concerning the Blur contract.
var BlurMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"marketId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"tradeData\",\"type\":\"bytes\"}],\"internalType\":\"structTradeDetails[]\",\"name\":\"tradeDetails\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAssetType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTokenInfo[]\",\"name\":\"listTokens\",\"type\":\"tuple[]\"}],\"name\":\"batchBuyWithETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blur1\",\"outputs\":[{\"internalType\":\"contractIBlur\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blur2\",\"outputs\":[{\"internalType\":\"contractIBlur\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"sell\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"buy\",\"type\":\"tuple\"}],\"internalType\":\"structExecution[]\",\"name\":\"executions\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"bulkExecute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"sell\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"buy\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506115a9806100206000396000f3fe60806040526004361061007b5760003560e01c8063a64940211161004e578063a64940211461014b578063b4b5fc421461015e578063cef44c861461018a578063ed485a9c1461019d57600080fd5b80630514e66714610080578063150b7a021461009557806356323472146100df5780639fa207601461011f575b600080fd5b61009361008e36600461083a565b6101bf565b005b3480156100a157600080fd5b506100c16100b036600461095a565b630a85bd0160e11b95945050505050565b6040516001600160e01b031990911681526020015b60405180910390f35b3480156100eb57600080fd5b506101077339da41747a83aee658334415666f3ef92dd0d54181565b6040516001600160a01b0390911681526020016100d6565b34801561012b57600080fd5b506100c161013a36600461095a565b63bc197c8160e01b95945050505050565b610093610159366004610a10565b6102d2565b34801561016a57600080fd5b506100c161017936600461095a565b63f23a6e6160e01b95945050505050565b610093610198366004610a87565b610371565b3480156101a957600080fd5b506101076dad05ccc4f10045630fb830b9512781565b604051639a2b811560e01b81527339da41747a83aee658334415666f3ef92dd0d54190639a2b81159034906101fa9088908890600401610b40565b6000604051808303818588803b15801561021357600080fd5b505af1158015610227573d6000803e3d6000fd5b505050505060005b81518110156102cb576102b982828151811061024d5761024d610bd9565b6020026020010151600001518484848151811061026c5761026c610bd9565b60200260200101516020015185858151811061028a5761028a610bd9565b6020026020010151604001518686815181106102a8576102a8610bd9565b602002602001015160600151610441565b806102c381610bef565b91505061022f565b5050505050565b604051639a1fc3a760e01b81526dad05ccc4f10045630fb830b9512790639a1fc3a79034906103079087908790600401610edb565b6000604051808303818588803b15801561032057600080fd5b505af1158015610334573d6000803e3d6000fd5b5061036c935061034a9250869150819050610f09565b6103539061113d565b61035d8480610f09565b6103669061113d565b8361056d565b505050565b604051631677caff60e31b81526dad05ccc4f10045630fb830b951279063b3be57f89034906103a6908790879060040161116e565b6000604051808303818588803b1580156103bf57600080fd5b505af11580156103d3573d6000803e3d6000fd5b505050505060005b8281101561043b5760008484838181106103f7576103f7610bd9565b90506020028101906104099190611203565b610412906112c7565b805151602082015151919250610428918561056d565b508061043381610bef565b9150506103db565b50505050565b600085600181111561045557610455610c16565b036104c8576040516323b872dd60e01b81523060048201526001600160a01b038581166024830152604482018490528416906323b872dd906064015b600060405180830381600087803b1580156104ab57600080fd5b505af11580156104bf573d6000803e3d6000fd5b505050506102cb565b60018560018111156104dc576104dc610c16565b036105215760408051602081018252600081529051637921219560e11b81526001600160a01b0385169163f242432a9161049191309189918891889190600401611379565b60405162461bcd60e51b815260206004820152601f60248201527f50423a20746865206974656d2074797065206e6f7420737570706f7274656400604482015260640160405180910390fd5b600080600061057c868661059b565b925092509250610593818588606001518686610441565b505050505050565b6000806000836101000151856101000151116106335784604001516001600160a01b031663d5ec8c7786866040518363ffffffff1660e01b81526004016105e39291906114f1565b60a060405180830381865afa158015610600573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106249190611516565b919650945092506106b0915050565b83604001516001600160a01b0316630813a76685876040518363ffffffff1660e01b81526004016106659291906114f1565b60a060405180830381865afa158015610682573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106a69190611516565b9196509450925050505b9250925092565b60008083601f8401126106c957600080fd5b5081356001600160401b038111156106e057600080fd5b6020830191508360208260051b85010111156106fb57600080fd5b9250929050565b6001600160a01b038116811461071757600080fd5b50565b803561072581610702565b919050565b634e487b7160e01b600052604160045260246000fd5b604051608081016001600160401b03811182821017156107625761076261072a565b60405290565b604080519081016001600160401b03811182821017156107625761076261072a565b6040516101a081016001600160401b03811182821017156107625761076261072a565b60405160e081016001600160401b03811182821017156107625761076261072a565b604051601f8201601f191681016001600160401b03811182821017156107f7576107f761072a565b604052919050565b60006001600160401b038211156108185761081861072a565b5060051b60200190565b6002811061071757600080fd5b803561072581610822565b600080600080606080868803121561085157600080fd5b85356001600160401b038082111561086857600080fd5b61087489838a016106b7565b9097509550602091508782013561088a81610702565b94506040888101358281111561089f57600080fd5b89019150601f82018a136108b257600080fd5b81356108c56108c0826107ff565b6107cf565b81815260079190911b8301840190848101908c8311156108e457600080fd5b938501935b82851015610949576080858e0312156109025760008081fd5b61090a610740565b853561091581610822565b81528587013561092481610702565b81880152858501358582015287860135888201528252608090940193908501906108e9565b999c989b5096995050505050505050565b60008060008060006080868803121561097257600080fd5b853561097d81610702565b9450602086013561098d81610702565b93506040860135925060608601356001600160401b03808211156109b057600080fd5b818801915088601f8301126109c457600080fd5b8135818111156109d357600080fd5b8960208285010111156109e557600080fd5b9699959850939650602001949392505050565b600060e08284031215610a0a57600080fd5b50919050565b600080600060608486031215610a2557600080fd5b83356001600160401b0380821115610a3c57600080fd5b610a48878388016109f8565b94506020860135915080821115610a5e57600080fd5b50610a6b868287016109f8565b9250506040840135610a7c81610702565b809150509250925092565b600080600060408486031215610a9c57600080fd5b83356001600160401b03811115610ab257600080fd5b610abe868287016106b7565b9094509250506020840135610a7c81610702565b6000808335601e19843603018112610ae957600080fd5b83016020810192503590506001600160401b03811115610b0857600080fd5b8036038213156106fb57600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60208082528181018390526000906040808401600586901b850182018785805b89811015610bca57888403603f190185528235368c9003605e19018112610b85578283fd5b8b018035855287810135888601526060610ba188830183610ad2565b92508189880152610bb58288018483610b17565b978a0197965050509287019250600101610b60565b50919998505050505050505050565b634e487b7160e01b600052603260045260246000fd5b600060018201610c0f57634e487b7160e01b600052601160045260246000fd5b5060010190565b634e487b7160e01b600052602160045260246000fd5b60028110610c4a57634e487b7160e01b600052602160045260246000fd5b9052565b6000808335601e19843603018112610c6557600080fd5b83016020810192503590506001600160401b03811115610c8457600080fd5b8060061b36038213156106fb57600080fd5b803561ffff8116811461072557600080fd5b8183526000602080850194508260005b85811015610cfc5761ffff610ccc83610c96565b16875282820135610cdc81610702565b6001600160a01b0316878401526040968701969190910190600101610cb8565b509495945050505050565b60006101a0610d2684610d198561071a565b6001600160a01b03169052565b610d326020840161082f565b610d3f6020860182610c2c565b50610d4c6040840161071a565b6001600160a01b03166040850152610d666060840161071a565b6001600160a01b0381166060860152506080830135608085015260a083013560a0850152610d9660c0840161071a565b6001600160a01b031660c085015260e0838101359085015261010080840135908501526101208084013590850152610140610dd381850185610c4e565b8383880152610de58488018284610ca8565b9350505050610160808401358186015250610180610e0581850185610ad2565b86840383880152610e17848284610b17565b979650505050505050565b803560ff8116811461072557600080fd5b6000813561019e19833603018112610e4a57600080fd5b60e08452610e5d60e08501848301610d07565b9050610e6b60208401610e22565b60ff81166020860152506040830135604085015260608301356060850152610e966080840184610ad2565b8583036080870152610ea9838284610b17565b92505050610eb960a0840161082f565b610ec660a0860182610c2c565b5060c083013560c08501528091505092915050565b604081526000610eee6040830185610e33565b8281036020840152610f008185610e33565b95945050505050565b6000823561019e19833603018112610f2057600080fd5b9190910192915050565b600082601f830112610f3b57600080fd5b81356020610f4b6108c0836107ff565b82815260069290921b84018101918181019086841115610f6a57600080fd5b8286015b84811015610fb95760408189031215610f875760008081fd5b610f8f610768565b610f9882610c96565b815284820135610fa781610702565b81860152835291830191604001610f6e565b509695505050505050565b600082601f830112610fd557600080fd5b81356001600160401b03811115610fee57610fee61072a565b611001601f8201601f19166020016107cf565b81815284602083860101111561101657600080fd5b816020850160208301376000918101602001919091529392505050565b60006101a0828403121561104657600080fd5b61104e61078a565b90506110598261071a565b81526110676020830161082f565b60208201526110786040830161071a565b60408201526110896060830161071a565b60608201526080820135608082015260a082013560a08201526110ae60c0830161071a565b60c082015260e0828101359082015261010080830135908201526101208083013590820152610140808301356001600160401b03808211156110ef57600080fd5b6110fb86838701610f2a565b838501526101609250828501358385015261018092508285013591508082111561112457600080fd5b5061113185828601610fc4565b82840152505092915050565b60006111493683611033565b92915050565b6000823560de1983360301811261116557600080fd5b90910192915050565b60208082528181018390526000906040808401600586901b850182018785805b89811015610bca57888403603f190185528235368c9003603e190181126111b3578283fd5b8b016111bf818061114f565b8786526111ce88870182610e33565b90506111dc8983018361114f565b9150858103898701526111ef8183610e33565b96890196955050509186019160010161118e565b60008235603e19833603018112610f2057600080fd5b600060e0828403121561122b57600080fd5b6112336107ad565b905081356001600160401b038082111561124c57600080fd5b61125885838601611033565b835261126660208501610e22565b60208401526040840135604084015260608401356060840152608084013591508082111561129357600080fd5b506112a084828501610fc4565b6080830152506112b260a0830161082f565b60a082015260c082013560c082015292915050565b6000604082360312156112d957600080fd5b6112e1610768565b82356001600160401b03808211156112f857600080fd5b61130436838701611219565b8352602085013591508082111561131a57600080fd5b5061132736828601611219565b60208301525092915050565b6000815180845260005b818110156113595760208185018101518683018201520161133d565b506000602082860101526020601f19601f83011685010191505092915050565b6001600160a01b03868116825285166020820152604081018490526060810183905260a060808201819052600090610e1790830184611333565b600081518084526020808501945080840160005b83811015610cfc578151805161ffff1688528301516001600160a01b031683880152604090960195908201906001016113c7565b80516001600160a01b0316825260006101a0602083015161141f6020860182610c2c565b50604083015161143a60408601826001600160a01b03169052565b50606083015161145560608601826001600160a01b03169052565b506080830151608085015260a083015160a085015260c083015161148460c08601826001600160a01b03169052565b5060e083015160e08501526101008084015181860152506101208084015181860152506101408084015182828701526114bf838701826113b3565b9250505061016080840151818601525061018080840151858303828701526114e78382611333565b9695505050505050565b60408152600061150460408301856113fb565b8281036020840152610f0081856113fb565b600080600080600060a0868803121561152e57600080fd5b8551801515811461153e57600080fd5b80955050602086015193506040860151925060608601519150608086015161156581610822565b80915050929550929590935056fea2646970667358221220fbc0d216d86572c52636252030498c80cb5f39d379a34f530fc3657067f181eb64736f6c63430008110033",
}

// BlurABI is the input ABI used to generate the binding from.
// Deprecated: Use BlurMetaData.ABI instead.
var BlurABI = BlurMetaData.ABI

// BlurBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BlurMetaData.Bin instead.
var BlurBin = BlurMetaData.Bin

// DeployBlur deploys a new Ethereum contract, binding an instance of Blur to it.
func DeployBlur(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Blur, error) {
	parsed, err := BlurMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BlurBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Blur{BlurCaller: BlurCaller{contract: contract}, BlurTransactor: BlurTransactor{contract: contract}, BlurFilterer: BlurFilterer{contract: contract}}, nil
}

// Blur is an auto generated Go binding around an Ethereum contract.
type Blur struct {
	BlurCaller     // Read-only binding to the contract
	BlurTransactor // Write-only binding to the contract
	BlurFilterer   // Log filterer for contract events
}

// BlurCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlurCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlurTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlurTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlurFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlurFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlurSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlurSession struct {
	Contract     *Blur             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlurCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlurCallerSession struct {
	Contract *BlurCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BlurTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlurTransactorSession struct {
	Contract     *BlurTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlurRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlurRaw struct {
	Contract *Blur // Generic contract binding to access the raw methods on
}

// BlurCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlurCallerRaw struct {
	Contract *BlurCaller // Generic read-only contract binding to access the raw methods on
}

// BlurTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlurTransactorRaw struct {
	Contract *BlurTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlur creates a new instance of Blur, bound to a specific deployed contract.
func NewBlur(address common.Address, backend bind.ContractBackend) (*Blur, error) {
	contract, err := bindBlur(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Blur{BlurCaller: BlurCaller{contract: contract}, BlurTransactor: BlurTransactor{contract: contract}, BlurFilterer: BlurFilterer{contract: contract}}, nil
}

// NewBlurCaller creates a new read-only instance of Blur, bound to a specific deployed contract.
func NewBlurCaller(address common.Address, caller bind.ContractCaller) (*BlurCaller, error) {
	contract, err := bindBlur(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlurCaller{contract: contract}, nil
}

// NewBlurTransactor creates a new write-only instance of Blur, bound to a specific deployed contract.
func NewBlurTransactor(address common.Address, transactor bind.ContractTransactor) (*BlurTransactor, error) {
	contract, err := bindBlur(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlurTransactor{contract: contract}, nil
}

// NewBlurFilterer creates a new log filterer instance of Blur, bound to a specific deployed contract.
func NewBlurFilterer(address common.Address, filterer bind.ContractFilterer) (*BlurFilterer, error) {
	contract, err := bindBlur(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlurFilterer{contract: contract}, nil
}

// bindBlur binds a generic wrapper to an already deployed contract.
func bindBlur(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BlurABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blur *BlurRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blur.Contract.BlurCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blur *BlurRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blur.Contract.BlurTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blur *BlurRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blur.Contract.BlurTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blur *BlurCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blur.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blur *BlurTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blur.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blur *BlurTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blur.Contract.contract.Transact(opts, method, params...)
}

// Blur1 is a free data retrieval call binding the contract method 0xed485a9c.
//
// Solidity: function blur1() view returns(address)
func (_Blur *BlurCaller) Blur1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Blur.contract.Call(opts, &out, "blur1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Blur1 is a free data retrieval call binding the contract method 0xed485a9c.
//
// Solidity: function blur1() view returns(address)
func (_Blur *BlurSession) Blur1() (common.Address, error) {
	return _Blur.Contract.Blur1(&_Blur.CallOpts)
}

// Blur1 is a free data retrieval call binding the contract method 0xed485a9c.
//
// Solidity: function blur1() view returns(address)
func (_Blur *BlurCallerSession) Blur1() (common.Address, error) {
	return _Blur.Contract.Blur1(&_Blur.CallOpts)
}

// Blur2 is a free data retrieval call binding the contract method 0x56323472.
//
// Solidity: function blur2() view returns(address)
func (_Blur *BlurCaller) Blur2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Blur.contract.Call(opts, &out, "blur2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Blur2 is a free data retrieval call binding the contract method 0x56323472.
//
// Solidity: function blur2() view returns(address)
func (_Blur *BlurSession) Blur2() (common.Address, error) {
	return _Blur.Contract.Blur2(&_Blur.CallOpts)
}

// Blur2 is a free data retrieval call binding the contract method 0x56323472.
//
// Solidity: function blur2() view returns(address)
func (_Blur *BlurCallerSession) Blur2() (common.Address, error) {
	return _Blur.Contract.Blur2(&_Blur.CallOpts)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0x9fa20760.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Blur *BlurCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Blur.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0x9fa20760.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Blur *BlurSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Blur.Contract.OnERC1155BatchReceived(&_Blur.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0x9fa20760.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Blur *BlurCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Blur.Contract.OnERC1155BatchReceived(&_Blur.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xb4b5fc42.
//
// Solidity: function onERC1155Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Blur *BlurCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Blur.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xb4b5fc42.
//
// Solidity: function onERC1155Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Blur *BlurSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Blur.Contract.OnERC1155Received(&_Blur.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xb4b5fc42.
//
// Solidity: function onERC1155Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Blur *BlurCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Blur.Contract.OnERC1155Received(&_Blur.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Blur *BlurCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Blur.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Blur *BlurSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Blur.Contract.OnERC721Received(&_Blur.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Blur *BlurCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Blur.Contract.OnERC721Received(&_Blur.CallOpts, arg0, arg1, arg2, arg3)
}

// BatchBuyWithETH is a paid mutator transaction binding the contract method 0x0514e667.
//
// Solidity: function batchBuyWithETH((uint256,uint256,bytes)[] tradeDetails, address recipient, (uint8,address,uint256,uint256)[] listTokens) payable returns()
func (_Blur *BlurTransactor) BatchBuyWithETH(opts *bind.TransactOpts, tradeDetails []TradeDetails, recipient common.Address, listTokens []TokenInfo) (*types.Transaction, error) {
	return _Blur.contract.Transact(opts, "batchBuyWithETH", tradeDetails, recipient, listTokens)
}

// BatchBuyWithETH is a paid mutator transaction binding the contract method 0x0514e667.
//
// Solidity: function batchBuyWithETH((uint256,uint256,bytes)[] tradeDetails, address recipient, (uint8,address,uint256,uint256)[] listTokens) payable returns()
func (_Blur *BlurSession) BatchBuyWithETH(tradeDetails []TradeDetails, recipient common.Address, listTokens []TokenInfo) (*types.Transaction, error) {
	return _Blur.Contract.BatchBuyWithETH(&_Blur.TransactOpts, tradeDetails, recipient, listTokens)
}

// BatchBuyWithETH is a paid mutator transaction binding the contract method 0x0514e667.
//
// Solidity: function batchBuyWithETH((uint256,uint256,bytes)[] tradeDetails, address recipient, (uint8,address,uint256,uint256)[] listTokens) payable returns()
func (_Blur *BlurTransactorSession) BatchBuyWithETH(tradeDetails []TradeDetails, recipient common.Address, listTokens []TokenInfo) (*types.Transaction, error) {
	return _Blur.Contract.BatchBuyWithETH(&_Blur.TransactOpts, tradeDetails, recipient, listTokens)
}

// BulkExecute is a paid mutator transaction binding the contract method 0xcef44c86.
//
// Solidity: function bulkExecute((((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256),((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256))[] executions, address recipient) payable returns()
func (_Blur *BlurTransactor) BulkExecute(opts *bind.TransactOpts, executions []Execution, recipient common.Address) (*types.Transaction, error) {
	return _Blur.contract.Transact(opts, "bulkExecute", executions, recipient)
}

// BulkExecute is a paid mutator transaction binding the contract method 0xcef44c86.
//
// Solidity: function bulkExecute((((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256),((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256))[] executions, address recipient) payable returns()
func (_Blur *BlurSession) BulkExecute(executions []Execution, recipient common.Address) (*types.Transaction, error) {
	return _Blur.Contract.BulkExecute(&_Blur.TransactOpts, executions, recipient)
}

// BulkExecute is a paid mutator transaction binding the contract method 0xcef44c86.
//
// Solidity: function bulkExecute((((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256),((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256))[] executions, address recipient) payable returns()
func (_Blur *BlurTransactorSession) BulkExecute(executions []Execution, recipient common.Address) (*types.Transaction, error) {
	return _Blur.Contract.BulkExecute(&_Blur.TransactOpts, executions, recipient)
}

// Execute is a paid mutator transaction binding the contract method 0xa6494021.
//
// Solidity: function execute(((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) sell, ((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) buy, address recipient) payable returns()
func (_Blur *BlurTransactor) Execute(opts *bind.TransactOpts, sell Input, buy Input, recipient common.Address) (*types.Transaction, error) {
	return _Blur.contract.Transact(opts, "execute", sell, buy, recipient)
}

// Execute is a paid mutator transaction binding the contract method 0xa6494021.
//
// Solidity: function execute(((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) sell, ((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) buy, address recipient) payable returns()
func (_Blur *BlurSession) Execute(sell Input, buy Input, recipient common.Address) (*types.Transaction, error) {
	return _Blur.Contract.Execute(&_Blur.TransactOpts, sell, buy, recipient)
}

// Execute is a paid mutator transaction binding the contract method 0xa6494021.
//
// Solidity: function execute(((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) sell, ((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) buy, address recipient) payable returns()
func (_Blur *BlurTransactorSession) Execute(sell Input, buy Input, recipient common.Address) (*types.Transaction, error) {
	return _Blur.Contract.Execute(&_Blur.TransactOpts, sell, buy, recipient)
}
