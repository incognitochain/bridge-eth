// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapter

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

// AdapterMetaData contains all meta data concerning the Adapter contract.
var AdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"blur1\",\"outputs\":[{\"internalType\":\"contractIBlur\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"sell\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"buy\",\"type\":\"tuple\"}],\"internalType\":\"structExecution[]\",\"name\":\"executions\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"bulkExecute\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumPNFTAdapter.PNFTMarket[]\",\"name\":\"_markets\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes[]\",\"name\":\"messages\",\"type\":\"bytes[]\"}],\"name\":\"buyBatchETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"sell\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"buy\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"opensea\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611651806100206000396000f3fe60806040526004361061007f5760003560e01c8063b4b5fc421161004e578063b4b5fc421461016e578063cd22e5631461019a578063cef44c86146101ad578063ed485a9c146101c057600080fd5b8063150b7a021461008b578063511ed382146100d55780639fa2076014610110578063a64940211461013c57600080fd5b3661008657005b600080fd5b34801561009757600080fd5b506100b76100a6366004610878565b630a85bd0160e11b95945050505050565b6040516001600160e01b031990911681526020015b60405180910390f35b3480156100e157600080fd5b506100f86e6c3852cbef3e08e8df289169ede58181565b6040516001600160a01b0390911681526020016100cc565b34801561011c57600080fd5b506100b761012b366004610878565b63bc197c8160e01b95945050505050565b61014f61014a36600461092e565b6101e8565b604080516001600160a01b0390931683526020830191909152016100cc565b34801561017a57600080fd5b506100b7610189366004610878565b63f23a6e6160e01b95945050505050565b61014f6101a83660046109f0565b6102ae565b61014f6101bb366004610a5b565b6104e0565b3480156101cc57600080fd5b506100f87387e5ffa37503487691c75359401080b1e2fbde5e81565b604051639a1fc3a760e01b815260009081907387e5ffa37503487691c75359401080b1e2fbde5e90639a1fc3a79034906102289089908990600401610df1565b6000604051808303818588803b15801561024157600080fd5b505af1158015610255573d6000803e3d6000fd5b5061028d935061026b9250889150819050610e1f565b6102749061111f565b61027e8680610e1f565b6102879061111f565b856105d7565b50734cb607c24ac252a0ce4b2e987ec4413da0f1e3ae946000945092505050565b6000808483146103055760405162461bcd60e51b815260206004820152601b60248201527f416461707465723a20696e76616c696420696e7075742064617461000000000060448201526064015b60405180910390fd5b6000805b8681101561045a57600188888381811061032557610325611131565b905060200201602081019061033a9190611147565b600181111561034b5761034b610abe565b03610367576e6c3852cbef3e08e8df289169ede581915061036b565b3091505b6000826001600160a01b03164788888581811061038a5761038a611131565b905060200281019061039c919061116b565b6040516103aa9291906111b1565b60006040518083038185875af1925050503d80600081146103e7576040519150601f19603f3d011682016040523d82523d6000602084013e6103ec565b606091505b50509050806104475760405162461bcd60e51b815260206004820152602160248201527f416461707465723a207265717565737420746f206d61726b6574206661696c656044820152601960fa1b60648201526084016102fc565b5080610452816111c1565b915050610309565b5047156104bd5760405162461bcd60e51b815260206004820152602b60248201527f416461707465723a2062616c616e63652061667465722065786563757465206d60448201526a757374206265207a65726f60a81b60648201526084016102fc565b50734cb607c24ac252a0ce4b2e987ec4413da0f1e3ae9660009650945050505050565b604051631677caff60e31b815260009081907387e5ffa37503487691c75359401080b1e2fbde5e9063b3be57f89034906105209089908990600401611207565b6000604051808303818588803b15801561053957600080fd5b505af115801561054d573d6000803e3d6000fd5b505050505060005b848110156105b557600086868381811061057157610571611131565b905060200281019061058391906112ab565b61058c9061136f565b8051516020820151519192506105a291876105d7565b50806105ad816111c1565b915050610555565b50734cb607c24ac252a0ce4b2e987ec4413da0f1e3ae95600095509350505050565b60008060006105e68686610605565b9250925092506105fd818588606001518686610721565b505050505050565b60008060008361010001518561010001511161069d5784604001516001600160a01b031663d5ec8c7786866040518363ffffffff1660e01b815260040161064d92919061155f565b60a060405180830381865afa15801561066a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061068e9190611584565b9196509450925061071a915050565b83604001516001600160a01b0316630813a76685876040518363ffffffff1660e01b81526004016106cf92919061155f565b60a060405180830381865afa1580156106ec573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107109190611584565b9196509450925050505b9250925092565b600085600181111561073557610735610abe565b036107a8576040516323b872dd60e01b81523060048201526001600160a01b038581166024830152604482018490528416906323b872dd906064015b600060405180830381600087803b15801561078b57600080fd5b505af115801561079f573d6000803e3d6000fd5b50505050610849565b60018560018111156107bc576107bc610abe565b036108015760408051602081018252600081529051637921219560e11b81526001600160a01b0385169163f242432a91610771913091899188918891906004016115e1565b60405162461bcd60e51b815260206004820152601f60248201527f50423a20746865206974656d2074797065206e6f7420737570706f727465640060448201526064016102fc565b5050505050565b6001600160a01b038116811461086557600080fd5b50565b803561087381610850565b919050565b60008060008060006080868803121561089057600080fd5b853561089b81610850565b945060208601356108ab81610850565b93506040860135925060608601356001600160401b03808211156108ce57600080fd5b818801915088601f8301126108e257600080fd5b8135818111156108f157600080fd5b89602082850101111561090357600080fd5b9699959850939650602001949392505050565b600060e0828403121561092857600080fd5b50919050565b60008060006060848603121561094357600080fd5b83356001600160401b038082111561095a57600080fd5b61096687838801610916565b9450602086013591508082111561097c57600080fd5b5061098986828701610916565b925050604084013561099a81610850565b809150509250925092565b60008083601f8401126109b757600080fd5b5081356001600160401b038111156109ce57600080fd5b6020830191508360208260051b85010111156109e957600080fd5b9250929050565b60008060008060408587031215610a0657600080fd5b84356001600160401b0380821115610a1d57600080fd5b610a29888389016109a5565b90965094506020870135915080821115610a4257600080fd5b50610a4f878288016109a5565b95989497509550505050565b600080600060408486031215610a7057600080fd5b83356001600160401b03811115610a8657600080fd5b610a92868287016109a5565b909450925050602084013561099a81610850565b6002811061086557600080fd5b803561087381610aa6565b634e487b7160e01b600052602160045260246000fd5b60028110610af257634e487b7160e01b600052602160045260246000fd5b9052565b6000808335601e19843603018112610b0d57600080fd5b83016020810192503590506001600160401b03811115610b2c57600080fd5b8060061b36038213156109e957600080fd5b803561ffff8116811461087357600080fd5b8183526000602080850194508260005b85811015610ba45761ffff610b7483610b3e565b16875282820135610b8481610850565b6001600160a01b0316878401526040968701969190910190600101610b60565b509495945050505050565b6000808335601e19843603018112610bc657600080fd5b83016020810192503590506001600160401b03811115610be557600080fd5b8036038213156109e957600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60006101a0610c3c84610c2f85610868565b6001600160a01b03169052565b610c4860208401610ab3565b610c556020860182610ad4565b50610c6260408401610868565b6001600160a01b03166040850152610c7c60608401610868565b6001600160a01b0381166060860152506080830135608085015260a083013560a0850152610cac60c08401610868565b6001600160a01b031660c085015260e0838101359085015261010080840135908501526101208084013590850152610140610ce981850185610af6565b8383880152610cfb8488018284610b50565b9350505050610160808401358186015250610180610d1b81850185610baf565b86840383880152610d2d848284610bf4565b979650505050505050565b803560ff8116811461087357600080fd5b6000813561019e19833603018112610d6057600080fd5b60e08452610d7360e08501848301610c1d565b9050610d8160208401610d38565b60ff81166020860152506040830135604085015260608301356060850152610dac6080840184610baf565b8583036080870152610dbf838284610bf4565b92505050610dcf60a08401610ab3565b610ddc60a0860182610ad4565b5060c083013560c08501528091505092915050565b604081526000610e046040830185610d49565b8281036020840152610e168185610d49565b95945050505050565b6000823561019e19833603018112610e3657600080fd5b9190910192915050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715610e7857610e78610e40565b60405290565b6040516101a081016001600160401b0381118282101715610e7857610e78610e40565b60405160e081016001600160401b0381118282101715610e7857610e78610e40565b604051601f8201601f191681016001600160401b0381118282101715610eeb57610eeb610e40565b604052919050565b600082601f830112610f0457600080fd5b813560206001600160401b03821115610f1f57610f1f610e40565b610f2d818360051b01610ec3565b82815260069290921b84018101918181019086841115610f4c57600080fd5b8286015b84811015610f9b5760408189031215610f695760008081fd5b610f71610e56565b610f7a82610b3e565b815284820135610f8981610850565b81860152835291830191604001610f50565b509695505050505050565b600082601f830112610fb757600080fd5b81356001600160401b03811115610fd057610fd0610e40565b610fe3601f8201601f1916602001610ec3565b818152846020838601011115610ff857600080fd5b816020850160208301376000918101602001919091529392505050565b60006101a0828403121561102857600080fd5b611030610e7e565b905061103b82610868565b815261104960208301610ab3565b602082015261105a60408301610868565b604082015261106b60608301610868565b60608201526080820135608082015260a082013560a082015261109060c08301610868565b60c082015260e0828101359082015261010080830135908201526101208083013590820152610140808301356001600160401b03808211156110d157600080fd5b6110dd86838701610ef3565b838501526101609250828501358385015261018092508285013591508082111561110657600080fd5b5061111385828601610fa6565b82840152505092915050565b600061112b3683611015565b92915050565b634e487b7160e01b600052603260045260246000fd5b60006020828403121561115957600080fd5b813561116481610aa6565b9392505050565b6000808335601e1984360301811261118257600080fd5b8301803591506001600160401b0382111561119c57600080fd5b6020019150368190038213156109e957600080fd5b8183823760009101908152919050565b6000600182016111e157634e487b7160e01b600052601160045260246000fd5b5060010190565b6000823560de198336030181126111fe57600080fd5b90910192915050565b60208082528181018390526000906040808401600586901b850182018785805b8981101561129c57888403603f190185528235368c9003603e1901811261124c578283fd5b8b0161125881806111e8565b87865261126788870182610d49565b9050611275898301836111e8565b9150858103898701526112888183610d49565b968901969550505091860191600101611227565b50919998505050505050505050565b60008235603e19833603018112610e3657600080fd5b600060e082840312156112d357600080fd5b6112db610ea1565b905081356001600160401b03808211156112f457600080fd5b61130085838601611015565b835261130e60208501610d38565b60208401526040840135604084015260608401356060840152608084013591508082111561133b57600080fd5b5061134884828501610fa6565b60808301525061135a60a08301610ab3565b60a082015260c082013560c082015292915050565b60006040823603121561138157600080fd5b611389610e56565b82356001600160401b03808211156113a057600080fd5b6113ac368387016112c1565b835260208501359150808211156113c257600080fd5b506113cf368286016112c1565b60208301525092915050565b600081518084526020808501945080840160005b83811015610ba4578151805161ffff1688528301516001600160a01b031683880152604090960195908201906001016113ef565b6000815180845260005b818110156114495760208185018101518683018201520161142d565b506000602082860101526020601f19601f83011685010191505092915050565b80516001600160a01b0316825260006101a0602083015161148d6020860182610ad4565b5060408301516114a860408601826001600160a01b03169052565b5060608301516114c360608601826001600160a01b03169052565b506080830151608085015260a083015160a085015260c08301516114f260c08601826001600160a01b03169052565b5060e083015160e085015261010080840151818601525061012080840151818601525061014080840151828287015261152d838701826113db565b9250505061016080840151818601525061018080840151858303828701526115558382611423565b9695505050505050565b6040815260006115726040830185611469565b8281036020840152610e168185611469565b600080600080600060a0868803121561159c57600080fd5b855180151581146115ac57600080fd5b8095505060208601519350604086015192506060860151915060808601516115d381610aa6565b809150509295509295909350565b6001600160a01b03868116825285166020820152604081018490526060810183905260a060808201819052600090610d2d9083018461142356fea2646970667358221220d746cd5a3bc693027991cee7eb3b2cf53db919b5875c675dabfd56b1b44c722964736f6c63430008110033",
}

// AdapterABI is the input ABI used to generate the binding from.
// Deprecated: Use AdapterMetaData.ABI instead.
var AdapterABI = AdapterMetaData.ABI

// AdapterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AdapterMetaData.Bin instead.
var AdapterBin = AdapterMetaData.Bin

// DeployAdapter deploys a new Ethereum contract, binding an instance of Adapter to it.
func DeployAdapter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Adapter, error) {
	parsed, err := AdapterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AdapterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Adapter{AdapterCaller: AdapterCaller{contract: contract}, AdapterTransactor: AdapterTransactor{contract: contract}, AdapterFilterer: AdapterFilterer{contract: contract}}, nil
}

// Adapter is an auto generated Go binding around an Ethereum contract.
type Adapter struct {
	AdapterCaller     // Read-only binding to the contract
	AdapterTransactor // Write-only binding to the contract
	AdapterFilterer   // Log filterer for contract events
}

// AdapterCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdapterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdapterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdapterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdapterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AdapterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdapterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdapterSession struct {
	Contract     *Adapter          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdapterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdapterCallerSession struct {
	Contract *AdapterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AdapterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdapterTransactorSession struct {
	Contract     *AdapterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AdapterRaw is an auto generated low-level Go binding around an Ethereum contract.
type AdapterRaw struct {
	Contract *Adapter // Generic contract binding to access the raw methods on
}

// AdapterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AdapterCallerRaw struct {
	Contract *AdapterCaller // Generic read-only contract binding to access the raw methods on
}

// AdapterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AdapterTransactorRaw struct {
	Contract *AdapterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAdapter creates a new instance of Adapter, bound to a specific deployed contract.
func NewAdapter(address common.Address, backend bind.ContractBackend) (*Adapter, error) {
	contract, err := bindAdapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Adapter{AdapterCaller: AdapterCaller{contract: contract}, AdapterTransactor: AdapterTransactor{contract: contract}, AdapterFilterer: AdapterFilterer{contract: contract}}, nil
}

// NewAdapterCaller creates a new read-only instance of Adapter, bound to a specific deployed contract.
func NewAdapterCaller(address common.Address, caller bind.ContractCaller) (*AdapterCaller, error) {
	contract, err := bindAdapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdapterCaller{contract: contract}, nil
}

// NewAdapterTransactor creates a new write-only instance of Adapter, bound to a specific deployed contract.
func NewAdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*AdapterTransactor, error) {
	contract, err := bindAdapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdapterTransactor{contract: contract}, nil
}

// NewAdapterFilterer creates a new log filterer instance of Adapter, bound to a specific deployed contract.
func NewAdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*AdapterFilterer, error) {
	contract, err := bindAdapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdapterFilterer{contract: contract}, nil
}

// bindAdapter binds a generic wrapper to an already deployed contract.
func bindAdapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AdapterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Adapter *AdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Adapter.Contract.AdapterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Adapter *AdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Adapter.Contract.AdapterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Adapter *AdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Adapter.Contract.AdapterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Adapter *AdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Adapter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Adapter *AdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Adapter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Adapter *AdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Adapter.Contract.contract.Transact(opts, method, params...)
}

// Blur1 is a free data retrieval call binding the contract method 0xed485a9c.
//
// Solidity: function blur1() view returns(address)
func (_Adapter *AdapterCaller) Blur1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Adapter.contract.Call(opts, &out, "blur1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Blur1 is a free data retrieval call binding the contract method 0xed485a9c.
//
// Solidity: function blur1() view returns(address)
func (_Adapter *AdapterSession) Blur1() (common.Address, error) {
	return _Adapter.Contract.Blur1(&_Adapter.CallOpts)
}

// Blur1 is a free data retrieval call binding the contract method 0xed485a9c.
//
// Solidity: function blur1() view returns(address)
func (_Adapter *AdapterCallerSession) Blur1() (common.Address, error) {
	return _Adapter.Contract.Blur1(&_Adapter.CallOpts)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0x9fa20760.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Adapter *AdapterCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Adapter.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0x9fa20760.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Adapter *AdapterSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Adapter.Contract.OnERC1155BatchReceived(&_Adapter.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0x9fa20760.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Adapter *AdapterCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Adapter.Contract.OnERC1155BatchReceived(&_Adapter.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xb4b5fc42.
//
// Solidity: function onERC1155Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Adapter *AdapterCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Adapter.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xb4b5fc42.
//
// Solidity: function onERC1155Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Adapter *AdapterSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Adapter.Contract.OnERC1155Received(&_Adapter.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xb4b5fc42.
//
// Solidity: function onERC1155Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Adapter *AdapterCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Adapter.Contract.OnERC1155Received(&_Adapter.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Adapter *AdapterCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Adapter.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Adapter *AdapterSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Adapter.Contract.OnERC721Received(&_Adapter.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Adapter *AdapterCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Adapter.Contract.OnERC721Received(&_Adapter.CallOpts, arg0, arg1, arg2, arg3)
}

// Opensea is a free data retrieval call binding the contract method 0x511ed382.
//
// Solidity: function opensea() view returns(address)
func (_Adapter *AdapterCaller) Opensea(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Adapter.contract.Call(opts, &out, "opensea")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Opensea is a free data retrieval call binding the contract method 0x511ed382.
//
// Solidity: function opensea() view returns(address)
func (_Adapter *AdapterSession) Opensea() (common.Address, error) {
	return _Adapter.Contract.Opensea(&_Adapter.CallOpts)
}

// Opensea is a free data retrieval call binding the contract method 0x511ed382.
//
// Solidity: function opensea() view returns(address)
func (_Adapter *AdapterCallerSession) Opensea() (common.Address, error) {
	return _Adapter.Contract.Opensea(&_Adapter.CallOpts)
}

// BulkExecute is a paid mutator transaction binding the contract method 0xcef44c86.
//
// Solidity: function bulkExecute((((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256),((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256))[] executions, address recipient) payable returns(address, uint256)
func (_Adapter *AdapterTransactor) BulkExecute(opts *bind.TransactOpts, executions []Execution, recipient common.Address) (*types.Transaction, error) {
	return _Adapter.contract.Transact(opts, "bulkExecute", executions, recipient)
}

// BulkExecute is a paid mutator transaction binding the contract method 0xcef44c86.
//
// Solidity: function bulkExecute((((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256),((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256))[] executions, address recipient) payable returns(address, uint256)
func (_Adapter *AdapterSession) BulkExecute(executions []Execution, recipient common.Address) (*types.Transaction, error) {
	return _Adapter.Contract.BulkExecute(&_Adapter.TransactOpts, executions, recipient)
}

// BulkExecute is a paid mutator transaction binding the contract method 0xcef44c86.
//
// Solidity: function bulkExecute((((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256),((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256))[] executions, address recipient) payable returns(address, uint256)
func (_Adapter *AdapterTransactorSession) BulkExecute(executions []Execution, recipient common.Address) (*types.Transaction, error) {
	return _Adapter.Contract.BulkExecute(&_Adapter.TransactOpts, executions, recipient)
}

// BuyBatchETH is a paid mutator transaction binding the contract method 0xcd22e563.
//
// Solidity: function buyBatchETH(uint8[] _markets, bytes[] messages) payable returns(address, uint256)
func (_Adapter *AdapterTransactor) BuyBatchETH(opts *bind.TransactOpts, _markets []uint8, messages [][]byte) (*types.Transaction, error) {
	return _Adapter.contract.Transact(opts, "buyBatchETH", _markets, messages)
}

// BuyBatchETH is a paid mutator transaction binding the contract method 0xcd22e563.
//
// Solidity: function buyBatchETH(uint8[] _markets, bytes[] messages) payable returns(address, uint256)
func (_Adapter *AdapterSession) BuyBatchETH(_markets []uint8, messages [][]byte) (*types.Transaction, error) {
	return _Adapter.Contract.BuyBatchETH(&_Adapter.TransactOpts, _markets, messages)
}

// BuyBatchETH is a paid mutator transaction binding the contract method 0xcd22e563.
//
// Solidity: function buyBatchETH(uint8[] _markets, bytes[] messages) payable returns(address, uint256)
func (_Adapter *AdapterTransactorSession) BuyBatchETH(_markets []uint8, messages [][]byte) (*types.Transaction, error) {
	return _Adapter.Contract.BuyBatchETH(&_Adapter.TransactOpts, _markets, messages)
}

// Execute is a paid mutator transaction binding the contract method 0xa6494021.
//
// Solidity: function execute(((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) sell, ((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) buy, address recipient) payable returns(address, uint256)
func (_Adapter *AdapterTransactor) Execute(opts *bind.TransactOpts, sell Input, buy Input, recipient common.Address) (*types.Transaction, error) {
	return _Adapter.contract.Transact(opts, "execute", sell, buy, recipient)
}

// Execute is a paid mutator transaction binding the contract method 0xa6494021.
//
// Solidity: function execute(((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) sell, ((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) buy, address recipient) payable returns(address, uint256)
func (_Adapter *AdapterSession) Execute(sell Input, buy Input, recipient common.Address) (*types.Transaction, error) {
	return _Adapter.Contract.Execute(&_Adapter.TransactOpts, sell, buy, recipient)
}

// Execute is a paid mutator transaction binding the contract method 0xa6494021.
//
// Solidity: function execute(((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) sell, ((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) buy, address recipient) payable returns(address, uint256)
func (_Adapter *AdapterTransactorSession) Execute(sell Input, buy Input, recipient common.Address) (*types.Transaction, error) {
	return _Adapter.Contract.Execute(&_Adapter.TransactOpts, sell, buy, recipient)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Adapter *AdapterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Adapter.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Adapter *AdapterSession) Receive() (*types.Transaction, error) {
	return _Adapter.Contract.Receive(&_Adapter.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Adapter *AdapterTransactorSession) Receive() (*types.Transaction, error) {
	return _Adapter.Contract.Receive(&_Adapter.TransactOpts)
}
