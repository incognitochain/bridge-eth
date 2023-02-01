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
	Bin: "0x608060405234801561001057600080fd5b506115bb806100206000396000f3fe60806040526004361061007b5760003560e01c8063a64940211161004e578063a64940211461014b578063b4b5fc421461015e578063cef44c861461018a578063ed485a9c1461019d57600080fd5b80630514e66714610080578063150b7a021461009557806356323472146100df5780639fa207601461011f575b600080fd5b61009361008e36600461084c565b6101c5565b005b3480156100a157600080fd5b506100c16100b036600461096c565b630a85bd0160e11b95945050505050565b6040516001600160e01b031990911681526020015b60405180910390f35b3480156100eb57600080fd5b506101077339da41747a83aee658334415666f3ef92dd0d54181565b6040516001600160a01b0390911681526020016100d6565b34801561012b57600080fd5b506100c161013a36600461096c565b63bc197c8160e01b95945050505050565b610093610159366004610a22565b6102d8565b34801561016a57600080fd5b506100c161017936600461096c565b63f23a6e6160e01b95945050505050565b610093610198366004610a99565b61037d565b3480156101a957600080fd5b50610107739a886f7cb772fc990dd8265cb1a9547db6cfe2e381565b604051639a2b811560e01b81527339da41747a83aee658334415666f3ef92dd0d54190639a2b81159034906102009088908890600401610b52565b6000604051808303818588803b15801561021957600080fd5b505af115801561022d573d6000803e3d6000fd5b505050505060005b81518110156102d1576102bf82828151811061025357610253610beb565b6020026020010151600001518484848151811061027257610272610beb565b60200260200101516020015185858151811061029057610290610beb565b6020026020010151604001518686815181106102ae576102ae610beb565b602002602001015160600151610453565b806102c981610c01565b915050610235565b5050505050565b604051639a1fc3a760e01b8152739a886f7cb772fc990dd8265cb1a9547db6cfe2e390639a1fc3a79034906103139087908790600401610eed565b6000604051808303818588803b15801561032c57600080fd5b505af1158015610340573d6000803e3d6000fd5b5061037893506103569250869150819050610f1b565b61035f9061114f565b6103698480610f1b565b6103729061114f565b8361057f565b505050565b604051631677caff60e31b8152739a886f7cb772fc990dd8265cb1a9547db6cfe2e39063b3be57f89034906103b89087908790600401611180565b6000604051808303818588803b1580156103d157600080fd5b505af11580156103e5573d6000803e3d6000fd5b505050505060005b8281101561044d57600084848381811061040957610409610beb565b905060200281019061041b9190611215565b610424906112d9565b80515160208201515191925061043a918561057f565b508061044581610c01565b9150506103ed565b50505050565b600085600181111561046757610467610c28565b036104da576040516323b872dd60e01b81523060048201526001600160a01b038581166024830152604482018490528416906323b872dd906064015b600060405180830381600087803b1580156104bd57600080fd5b505af11580156104d1573d6000803e3d6000fd5b505050506102d1565b60018560018111156104ee576104ee610c28565b036105335760408051602081018252600081529051637921219560e11b81526001600160a01b0385169163f242432a916104a39130918991889188919060040161138b565b60405162461bcd60e51b815260206004820152601f60248201527f50423a20746865206974656d2074797065206e6f7420737570706f7274656400604482015260640160405180910390fd5b600080600061058e86866105ad565b9250925092506105a5818588606001518686610453565b505050505050565b6000806000836101000151856101000151116106455784604001516001600160a01b031663d5ec8c7786866040518363ffffffff1660e01b81526004016105f5929190611503565b60a060405180830381865afa158015610612573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106369190611528565b919650945092506106c2915050565b83604001516001600160a01b0316630813a76685876040518363ffffffff1660e01b8152600401610677929190611503565b60a060405180830381865afa158015610694573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106b89190611528565b9196509450925050505b9250925092565b60008083601f8401126106db57600080fd5b5081356001600160401b038111156106f257600080fd5b6020830191508360208260051b850101111561070d57600080fd5b9250929050565b6001600160a01b038116811461072957600080fd5b50565b803561073781610714565b919050565b634e487b7160e01b600052604160045260246000fd5b604051608081016001600160401b03811182821017156107745761077461073c565b60405290565b604080519081016001600160401b03811182821017156107745761077461073c565b6040516101a081016001600160401b03811182821017156107745761077461073c565b60405160e081016001600160401b03811182821017156107745761077461073c565b604051601f8201601f191681016001600160401b03811182821017156108095761080961073c565b604052919050565b60006001600160401b0382111561082a5761082a61073c565b5060051b60200190565b6002811061072957600080fd5b803561073781610834565b600080600080606080868803121561086357600080fd5b85356001600160401b038082111561087a57600080fd5b61088689838a016106c9565b9097509550602091508782013561089c81610714565b9450604088810135828111156108b157600080fd5b89019150601f82018a136108c457600080fd5b81356108d76108d282610811565b6107e1565b81815260079190911b8301840190848101908c8311156108f657600080fd5b938501935b8285101561095b576080858e0312156109145760008081fd5b61091c610752565b853561092781610834565b81528587013561093681610714565b81880152858501358582015287860135888201528252608090940193908501906108fb565b999c989b5096995050505050505050565b60008060008060006080868803121561098457600080fd5b853561098f81610714565b9450602086013561099f81610714565b93506040860135925060608601356001600160401b03808211156109c257600080fd5b818801915088601f8301126109d657600080fd5b8135818111156109e557600080fd5b8960208285010111156109f757600080fd5b9699959850939650602001949392505050565b600060e08284031215610a1c57600080fd5b50919050565b600080600060608486031215610a3757600080fd5b83356001600160401b0380821115610a4e57600080fd5b610a5a87838801610a0a565b94506020860135915080821115610a7057600080fd5b50610a7d86828701610a0a565b9250506040840135610a8e81610714565b809150509250925092565b600080600060408486031215610aae57600080fd5b83356001600160401b03811115610ac457600080fd5b610ad0868287016106c9565b9094509250506020840135610a8e81610714565b6000808335601e19843603018112610afb57600080fd5b83016020810192503590506001600160401b03811115610b1a57600080fd5b80360382131561070d57600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60208082528181018390526000906040808401600586901b850182018785805b89811015610bdc57888403603f190185528235368c9003605e19018112610b97578283fd5b8b018035855287810135888601526060610bb388830183610ae4565b92508189880152610bc78288018483610b29565b978a0197965050509287019250600101610b72565b50919998505050505050505050565b634e487b7160e01b600052603260045260246000fd5b600060018201610c2157634e487b7160e01b600052601160045260246000fd5b5060010190565b634e487b7160e01b600052602160045260246000fd5b60028110610c5c57634e487b7160e01b600052602160045260246000fd5b9052565b6000808335601e19843603018112610c7757600080fd5b83016020810192503590506001600160401b03811115610c9657600080fd5b8060061b360382131561070d57600080fd5b803561ffff8116811461073757600080fd5b8183526000602080850194508260005b85811015610d0e5761ffff610cde83610ca8565b16875282820135610cee81610714565b6001600160a01b0316878401526040968701969190910190600101610cca565b509495945050505050565b60006101a0610d3884610d2b8561072c565b6001600160a01b03169052565b610d4460208401610841565b610d516020860182610c3e565b50610d5e6040840161072c565b6001600160a01b03166040850152610d786060840161072c565b6001600160a01b0381166060860152506080830135608085015260a083013560a0850152610da860c0840161072c565b6001600160a01b031660c085015260e0838101359085015261010080840135908501526101208084013590850152610140610de581850185610c60565b8383880152610df78488018284610cba565b9350505050610160808401358186015250610180610e1781850185610ae4565b86840383880152610e29848284610b29565b979650505050505050565b803560ff8116811461073757600080fd5b6000813561019e19833603018112610e5c57600080fd5b60e08452610e6f60e08501848301610d19565b9050610e7d60208401610e34565b60ff81166020860152506040830135604085015260608301356060850152610ea86080840184610ae4565b8583036080870152610ebb838284610b29565b92505050610ecb60a08401610841565b610ed860a0860182610c3e565b5060c083013560c08501528091505092915050565b604081526000610f006040830185610e45565b8281036020840152610f128185610e45565b95945050505050565b6000823561019e19833603018112610f3257600080fd5b9190910192915050565b600082601f830112610f4d57600080fd5b81356020610f5d6108d283610811565b82815260069290921b84018101918181019086841115610f7c57600080fd5b8286015b84811015610fcb5760408189031215610f995760008081fd5b610fa161077a565b610faa82610ca8565b815284820135610fb981610714565b81860152835291830191604001610f80565b509695505050505050565b600082601f830112610fe757600080fd5b81356001600160401b038111156110005761100061073c565b611013601f8201601f19166020016107e1565b81815284602083860101111561102857600080fd5b816020850160208301376000918101602001919091529392505050565b60006101a0828403121561105857600080fd5b61106061079c565b905061106b8261072c565b815261107960208301610841565b602082015261108a6040830161072c565b604082015261109b6060830161072c565b60608201526080820135608082015260a082013560a08201526110c060c0830161072c565b60c082015260e0828101359082015261010080830135908201526101208083013590820152610140808301356001600160401b038082111561110157600080fd5b61110d86838701610f3c565b838501526101609250828501358385015261018092508285013591508082111561113657600080fd5b5061114385828601610fd6565b82840152505092915050565b600061115b3683611045565b92915050565b6000823560de1983360301811261117757600080fd5b90910192915050565b60208082528181018390526000906040808401600586901b850182018785805b89811015610bdc57888403603f190185528235368c9003603e190181126111c5578283fd5b8b016111d18180611161565b8786526111e088870182610e45565b90506111ee89830183611161565b9150858103898701526112018183610e45565b9689019695505050918601916001016111a0565b60008235603e19833603018112610f3257600080fd5b600060e0828403121561123d57600080fd5b6112456107bf565b905081356001600160401b038082111561125e57600080fd5b61126a85838601611045565b835261127860208501610e34565b6020840152604084013560408401526060840135606084015260808401359150808211156112a557600080fd5b506112b284828501610fd6565b6080830152506112c460a08301610841565b60a082015260c082013560c082015292915050565b6000604082360312156112eb57600080fd5b6112f361077a565b82356001600160401b038082111561130a57600080fd5b6113163683870161122b565b8352602085013591508082111561132c57600080fd5b506113393682860161122b565b60208301525092915050565b6000815180845260005b8181101561136b5760208185018101518683018201520161134f565b506000602082860101526020601f19601f83011685010191505092915050565b6001600160a01b03868116825285166020820152604081018490526060810183905260a060808201819052600090610e2990830184611345565b600081518084526020808501945080840160005b83811015610d0e578151805161ffff1688528301516001600160a01b031683880152604090960195908201906001016113d9565b80516001600160a01b0316825260006101a060208301516114316020860182610c3e565b50604083015161144c60408601826001600160a01b03169052565b50606083015161146760608601826001600160a01b03169052565b506080830151608085015260a083015160a085015260c083015161149660c08601826001600160a01b03169052565b5060e083015160e08501526101008084015181860152506101208084015181860152506101408084015182828701526114d1838701826113c5565b9250505061016080840151818601525061018080840151858303828701526114f98382611345565b9695505050505050565b604081526000611516604083018561140d565b8281036020840152610f12818561140d565b600080600080600060a0868803121561154057600080fd5b8551801515811461155057600080fd5b80955050602086015193506040860151925060608601519150608086015161157781610834565b80915050929550929590935056fea264697066735822122074e94d636d6f965e3637f210f7544942ff6df723fb6aaf1ea4e7a9eb049dec4664736f6c63430008110033",
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
