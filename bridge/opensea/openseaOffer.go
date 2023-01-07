// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package opensea

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

// OpenseaOfferMetaData contains all meta data concerning the OpenseaOffer contract.
var OpenseaOfferMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"counter\",\"type\":\"uint256\"}],\"internalType\":\"structOrderComponents\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"offerSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"txId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"regulatorSignData\",\"type\":\"bytes\"}],\"name\":\"cancelOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"counter\",\"type\":\"uint256\"}],\"internalType\":\"structOrderComponents\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"otaKey\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"offer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"offers\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"otaKey\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offerAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"seaport\",\"outputs\":[{\"internalType\":\"contractConsiderationInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"domainSeparator_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"structHash_\",\"type\":\"bytes32\"}],\"name\":\"toTypedDataHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"contractVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"contractWETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50612159806100206000396000f3fe6080604052600436106100865760003560e01c8063b453966111610059578063b453966114610169578063d34517ed1461017e578063f5c7bd701461019e578063f698da25146101c1578063fbfa77cf146101f557600080fd5b80631626ba7e1461008b5780633fc8cef3146100c9578063474d3ff0146101095780637df7a71c1461013b575b600080fd5b34801561009757600080fd5b506100ab6100a6366004611397565b61021d565b6040516001600160e01b031990911681526020015b60405180910390f35b3480156100d557600080fd5b506100f173c02aaa39b223fe8d0a0e5c4f27ead9083c756cc281565b6040516001600160a01b0390911681526020016100c0565b34801561011557600080fd5b506101296101243660046113e2565b61024f565b6040516100c096959493929190611441565b34801561014757600080fd5b5061015b610156366004611494565b6103a2565b6040519081526020016100c0565b61017c6101773660046114f7565b6103e3565b005b34801561018a57600080fd5b5061017c61019936600461159d565b610957565b3480156101aa57600080fd5b506100f16e6c3852cbef3e08e8df289169ede58181565b3480156101cd57600080fd5b5061015b7f712fdde1f147adcbb0fabb1aeb9c2d317530a46d266db095bc40c606fe94f0c281565b34801561020157600080fd5b506100f173c157cc3077ddfa425bae12d2f3002668971a4e3d81565b600061022a848484610ea3565b1561023d5750630b135d3f60e11b610248565b506001600160e01b03195b9392505050565b60006020819052908152604090208054819061026a9061163a565b80601f01602080910402602001604051908101604052809291908181526020018280546102969061163a565b80156102e35780601f106102b8576101008083540402835291602001916102e3565b820191906000526020600020905b8154815290600101906020018083116102c657829003601f168201915b505050600184015460028501805494956001600160a01b0390921694919350915061030d9061163a565b80601f01602080910402602001604051908101604052809291908181526020018280546103399061163a565b80156103865780601f1061035b57610100808354040283529160200191610386565b820191906000526020600020905b81548152906001019060200180831161036957829003601f168201915b5050505050908060030154908060040154908060050154905086565b60405161190160f01b602082015260228101839052604281018290526000906062016040516020818303038152906040528051906020012090505b92915050565b306103f1602088018861166e565b6001600160a01b03161461044c5760405162461bcd60e51b815260206004820152601d60248201527f4f70656e7365614f666665723a20696e76616c6964206f66666572657200000060448201526064015b60405180910390fd5b6040516379df72bd60e01b81526000906104e8907f712fdde1f147adcbb0fabb1aeb9c2d317530a46d266db095bc40c606fe94f0c2906e6c3852cbef3e08e8df289169ede581906379df72bd906104a7908c90600401611879565b602060405180830381865afa1580156104c4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101569190611972565b6000818152602081905260409020600301549091501561054a5760405162461bcd60e51b815260206004820152601b60248201527f4f70656e7365614f666665723a206f66666572206578697374656400000000006044820152606401610443565b600061058c8286868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061107b92505050565b90506000805b61059f60408b018b61198b565b90508160ff1610156106705760006105ba60408c018c61198b565b8360ff168181106105cd576105cd6119d3565b905060a002018036038101906105e39190611af9565b9050806080015181606001511461064b5760405162461bcd60e51b815260206004820152602660248201527f4f70656e7365614f666665723a20696e76616c696420737461727420656e6420604482015265185b5bdd5b9d60d21b6064820152608401610443565b606081015161065a9084611b2b565b925050808061066890611b3e565b915050610592565b508034146106cb5760405162461bcd60e51b815260206004820152602260248201527f4f70656e7365614f666665723a20696e76616c6964206f6666657220616d6f756044820152611b9d60f21b6064820152608401610443565b6040518060c0016040528089898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505050908252506001600160a01b03841660208083019190915260408051601f8a018390048302810183018252898152920191908990899081908401838280828437600092018290525093855250505060a08c013560208084019190915260c08d0135604080850191909152606090930185905286825281905220815181906107919082611bac565b5060208201516001820180546001600160a01b0319166001600160a01b03909216919091179055604082015160028201906107cc9082611bac565b50606082015181600301556080820151816004015560a0820151816005015590505073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc26001600160a01b031663d0e30db0346040518263ffffffff1660e01b81526004016000604051808303818588803b15801561083d57600080fd5b505af1158015610851573d6000803e3d6000fd5b5050604051636eb1769f60e11b81523060048201526001600160a01b038816602482015273c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2935063095ea7b392508791508490849063dd62ed3e90604401602060405180830381865afa1580156108c0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108e49190611972565b6108ee9190611b2b565b6040516001600160e01b031960e085901b1681526001600160a01b0390921660048301526024820152604401600060405180830381600087803b15801561093457600080fd5b505af1158015610948573d6000803e3d6000fd5b50505050505050505050505050565b6040516379df72bd60e01b81526000906e6c3852cbef3e08e8df289169ede581906379df72bd9061098c908a90600401611879565b602060405180830381865afa1580156109a9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109cd9190611972565b9050600080806109fd7f712fdde1f147adcbb0fabb1aeb9c2d317530a46d266db095bc40c606fe94f0c2856103a2565b81526020019081526020016000206040518060c0016040529081600082018054610a269061163a565b80601f0160208091040260200160405190810160405280929190818152602001828054610a529061163a565b8015610a9f5780601f10610a7457610100808354040283529160200191610a9f565b820191906000526020600020905b815481529060010190602001808311610a8257829003601f168201915b505050918352505060018201546001600160a01b03166020820152600282018054604090920191610acf9061163a565b80601f0160208091040260200160405190810160405280929190818152602001828054610afb9061163a565b8015610b485780601f10610b1d57610100808354040283529160200191610b48565b820191906000526020600020905b815481529060010190602001808311610b2b57829003601f168201915b50505050508152602001600382015481526020016004820154815260200160058201548152505090508060600151600014158015610b89575060a081015115155b610bd55760405162461bcd60e51b815260206004820152601b60248201527f4f70656e7365614f666665723a20696e76616c6964206f6666657200000000006044820152606401610443565b8060800151421015610c8557610c218288888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061107b92505050565b6001600160a01b031681602001516001600160a01b031614610c855760405162461bcd60e51b815260206004820152601f60248201527f4f70656e7365614f666665723a20696e76616c6964207369676e6174757265006044820152606401610443565b604080516001808252818301909252600091816020015b610ca46112d2565b815260200190600190039081610c9c579050509050610cc289611dc1565b81600081518110610cd557610cd56119d3565b6020908102919091010152604051630fd9f1e160e41b81526e6c3852cbef3e08e8df289169ede5819063fd9f1e1090610d12908490600401611f82565b6020604051808303816000875af1158015610d31573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d559190612084565b610db85760405162461bcd60e51b815260206004820152602e60248201527f4f70656e7365614f666665723a20657865637574652063616e63656c206f6e2060448201526d1cd9585c1bdc9d0819985a5b195960921b6064820152608401610443565b60a0820151604051632e1a7d4d60e01b8152600481019190915273c02aaa39b223fe8d0a0e5c4f27ead9083c756cc290632e1a7d4d90602401600060405180830381600087803b158015610e0b57600080fd5b505af1158015610e1f573d6000803e3d6000fd5b50505060a0830151835160405163c791d70560e01b815273c157cc3077ddfa425bae12d2f3002668971a4e3d935063c791d7059291610e66918b908b908b906004016120a6565b6000604051808303818588803b158015610e7f57600080fd5b505af1158015610e93573d6000803e3d6000fd5b5050505050505050505050505050565b6000806000808681526020019081526020016000206040518060c0016040529081600082018054610ed39061163a565b80601f0160208091040260200160405190810160405280929190818152602001828054610eff9061163a565b8015610f4c5780601f10610f2157610100808354040283529160200191610f4c565b820191906000526020600020905b815481529060010190602001808311610f2f57829003601f168201915b505050918352505060018201546001600160a01b03166020820152600282018054604090920191610f7c9061163a565b80601f0160208091040260200160405190810160405280929190818152602001828054610fa89061163a565b8015610ff55780601f10610fca57610100808354040283529160200191610ff5565b820191906000526020600020905b815481529060010190602001808311610fd857829003601f168201915b50505050508152602001600382015481526020016004820154815260200160058201548152505090508060600151600014806110535750838360405161103c9291906120f3565b604051809103902081604001518051906020012014155b806110615750428160800151105b15611070576000915050610248565b506001949350505050565b600081516041146110e25760405162461bcd60e51b815260206004820152603a602482015260008051602061210483398151915260448201527f3a20696e76616c6964207369676e6174757265206c656e6774680000000000006064820152608401610443565b6000826040815181106110f7576110f76119d3565b0160209081015190840151604085015160f89290921c92507f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08211156111935760405162461bcd60e51b815260206004820152603d602482015260008051602061210483398151915260448201527f3a20696e76616c6964207369676e6174757265202773272076616c75650000006064820152608401610443565b8260ff16601b141580156111ab57508260ff16601c14155b1561120c5760405162461bcd60e51b815260206004820152603d602482015260008051602061210483398151915260448201527f3a20696e76616c6964207369676e6174757265202776272076616c75650000006064820152608401610443565b60408051600081526020810180835288905260ff851691810191909152606081018290526080810183905260019060a0016020604051602081039080840390855afa15801561125f573d6000803e3d6000fd5b5050604051601f1901519450506001600160a01b0384166112c95760405162461bcd60e51b8152602060048201526030602482015260008051602061210483398151915260448201526f1d1024a72b20a624a22fa9a4a3a722a960811b6064820152608401610443565b50505092915050565b60405180610160016040528060006001600160a01b0316815260200160006001600160a01b0316815260200160608152602001606081526020016000600381111561131f5761131f6116e2565b815260006020820181905260408201819052606082018190526080820181905260a0820181905260c09091015290565b60008083601f84011261136157600080fd5b5081356001600160401b0381111561137857600080fd5b60208301915083602082850101111561139057600080fd5b9250929050565b6000806000604084860312156113ac57600080fd5b8335925060208401356001600160401b038111156113c957600080fd5b6113d58682870161134f565b9497909650939450505050565b6000602082840312156113f457600080fd5b5035919050565b6000815180845260005b8181101561142157602081850181015186830182015201611405565b506000602082860101526020601f19601f83011685010191505092915050565b60c08152600061145460c08301896113fb565b6001600160a01b0388166020840152828103604084015261147581886113fb565b60608401969096525050608081019290925260a0909101529392505050565b600080604083850312156114a757600080fd5b50508035926020909101359150565b600061016082840312156114c957600080fd5b50919050565b6001600160a01b03811681146114e457600080fd5b50565b80356114f2816114cf565b919050565b6000806000806000806080878903121561151057600080fd5b86356001600160401b038082111561152757600080fd5b6115338a838b016114b6565b9750602089013591508082111561154957600080fd5b6115558a838b0161134f565b9097509550604089013591508082111561156e57600080fd5b5061157b89828a0161134f565b909450925050606087013561158f816114cf565b809150509295509295509295565b600080600080600080608087890312156115b657600080fd5b86356001600160401b03808211156115cd57600080fd5b6115d98a838b016114b6565b975060208901359150808211156115ef57600080fd5b6115fb8a838b0161134f565b909750955060408901359450606089013591508082111561161b57600080fd5b5061162889828a0161134f565b979a9699509497509295939492505050565b600181811c9082168061164e57607f821691505b6020821081036114c957634e487b7160e01b600052602260045260246000fd5b60006020828403121561168057600080fd5b8135610248816114cf565b6000808335601e198436030181126116a257600080fd5b83016020810192503590506001600160401b038111156116c157600080fd5b60a08102360382131561139057600080fd5b8035600681106114f257600080fd5b634e487b7160e01b600052602160045260246000fd5b60068110611708576117086116e2565b9052565b8183526000602080850194508260005b858110156117805761173687611731846116d3565b6116f8565b82820135611743816114cf565b6001600160a01b03168388015260408281013590880152606080830135908801526080808301359088015260a0968701969091019060010161171c565b509495945050505050565b6000808335601e198436030181126117a257600080fd5b83016020810192503590506001600160401b038111156117c157600080fd5b60c08102360382131561139057600080fd5b8183526000602080850194508260005b85811015611780576117f887611731846116d3565b82820135611805816114cf565b6001600160a01b039081168885015260408381013590890152606080840135908901526080808401359089015260a09083820135611842816114cf565b169088015260c09687019691909101906001016117e3565b8035600481106114f257600080fd5b60048110611708576117086116e2565b6020815261189a6020820161188d846114e7565b6001600160a01b03169052565b60006118a8602084016114e7565b6001600160a01b0381166040840152506118c5604084018461168b565b6101608060608601526118dd6101808601838561170c565b92506118ec606087018761178b565b868503601f1901608088015292506119058484836117d3565b9350506119146080870161185a565b915061192360a0860183611869565b60a086013560c086015260c086013560e0860152610100915060e08601358286015261012082870135818701526101409250808701358387015250818601358186015250508091505092915050565b60006020828403121561198457600080fd5b5051919050565b6000808335601e198436030181126119a257600080fd5b8301803591506001600160401b038211156119bc57600080fd5b602001915060a08102360382131561139057600080fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b60405160c081016001600160401b0381118282101715611a2157611a216119e9565b60405290565b60405161016081016001600160401b0381118282101715611a2157611a216119e9565b604051601f8201601f191681016001600160401b0381118282101715611a7257611a726119e9565b604052919050565b600060a08284031215611a8c57600080fd5b60405160a081018181106001600160401b0382111715611aae57611aae6119e9565b604052905080611abd836116d3565b81526020830135611acd816114cf565b806020830152506040830135604082015260608301356060820152608083013560808201525092915050565b600060a08284031215611b0b57600080fd5b6102488383611a7a565b634e487b7160e01b600052601160045260246000fd5b808201808211156103dd576103dd611b15565b600060ff821660ff8103611b5457611b54611b15565b60010192915050565b601f821115611ba757600081815260208120601f850160051c81016020861015611b845750805b601f850160051c820191505b81811015611ba357828155600101611b90565b5050505b505050565b81516001600160401b03811115611bc557611bc56119e9565b611bd981611bd3845461163a565b84611b5d565b602080601f831160018114611c0e5760008415611bf65750858301515b600019600386901b1c1916600185901b178555611ba3565b600085815260208120601f198616915b82811015611c3d57888601518255948401946001909101908401611c1e565b5085821015611c5b5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60006001600160401b03821115611c8457611c846119e9565b5060051b60200190565b600082601f830112611c9f57600080fd5b81356020611cb4611caf83611c6b565b611a4a565b82815260a09283028501820192828201919087851115611cd357600080fd5b8387015b85811015611cf657611ce98982611a7a565b8452928401928101611cd7565b5090979650505050505050565b600082601f830112611d1457600080fd5b81356020611d24611caf83611c6b565b82815260c09283028501820192828201919087851115611d4357600080fd5b8387015b85811015611cf65781818a031215611d5f5760008081fd5b611d676119ff565b611d70826116d3565b815285820135611d7f816114cf565b8187015260408281013590820152606080830135908201526080808301359082015260a080830135611db0816114cf565b908201528452928401928101611d47565b60006101608236031215611dd457600080fd5b611ddc611a27565b611de5836114e7565b8152611df3602084016114e7565b602082015260408301356001600160401b0380821115611e1257600080fd5b611e1e36838701611c8e565b60408401526060850135915080821115611e3757600080fd5b50611e4436828601611d03565b606083015250611e566080840161185a565b608082015260a0838101359082015260c0808401359082015260e080840135908201526101008084013590820152610120808401359082015261014092830135928101929092525090565b600081518084526020808501945080840160005b83811015611780578151611eca8882516116f8565b838101516001600160a01b03168885015260408082015190890152606080820151908901526080908101519088015260a09096019590820190600101611eb5565b600081518084526020808501945080840160005b83811015611780578151611f348882516116f8565b808401516001600160a01b0390811689860152604080830151908a0152606080830151908a0152608080830151908a015260a091820151169088015260c09096019590820190600101611f1f565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b8381101561207657888303603f19018552815180516001600160a01b03168452610160818901516001600160a01b038116868b015250878201518189870152611ff482870182611ea1565b9150506060808301518683038288015261200e8382611f0b565b9250505060808083015161202482880182611869565b505060a0828101519086015260c0808301519086015260e08083015190860152610100808301519086015261012080830151908601526101409182015191909401529386019390860190600101611fa9565b509098975050505050505050565b60006020828403121561209657600080fd5b8151801515811461024857600080fd5b6060815260006120b960608301876113fb565b8560208401528281036040840152838152838560208301376000602085830101526020601f19601f86011682010191505095945050505050565b818382376000910190815291905056fe5369676e617475726556616c696461746f72237265636f7665725369676e6572a2646970667358221220da0c23efd8067d8482aa017ee499310a80c18d81acddd4348da095a07ca6153f64736f6c63430008110033",
}

// OpenseaOfferABI is the input ABI used to generate the binding from.
// Deprecated: Use OpenseaOfferMetaData.ABI instead.
var OpenseaOfferABI = OpenseaOfferMetaData.ABI

// OpenseaOfferBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OpenseaOfferMetaData.Bin instead.
var OpenseaOfferBin = OpenseaOfferMetaData.Bin

// DeployOpenseaOffer deploys a new Ethereum contract, binding an instance of OpenseaOffer to it.
func DeployOpenseaOffer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OpenseaOffer, error) {
	parsed, err := OpenseaOfferMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OpenseaOfferBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OpenseaOffer{OpenseaOfferCaller: OpenseaOfferCaller{contract: contract}, OpenseaOfferTransactor: OpenseaOfferTransactor{contract: contract}, OpenseaOfferFilterer: OpenseaOfferFilterer{contract: contract}}, nil
}

// OpenseaOffer is an auto generated Go binding around an Ethereum contract.
type OpenseaOffer struct {
	OpenseaOfferCaller     // Read-only binding to the contract
	OpenseaOfferTransactor // Write-only binding to the contract
	OpenseaOfferFilterer   // Log filterer for contract events
}

// OpenseaOfferCaller is an auto generated read-only Go binding around an Ethereum contract.
type OpenseaOfferCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpenseaOfferTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OpenseaOfferTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpenseaOfferFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OpenseaOfferFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpenseaOfferSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OpenseaOfferSession struct {
	Contract     *OpenseaOffer     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OpenseaOfferCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OpenseaOfferCallerSession struct {
	Contract *OpenseaOfferCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// OpenseaOfferTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OpenseaOfferTransactorSession struct {
	Contract     *OpenseaOfferTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// OpenseaOfferRaw is an auto generated low-level Go binding around an Ethereum contract.
type OpenseaOfferRaw struct {
	Contract *OpenseaOffer // Generic contract binding to access the raw methods on
}

// OpenseaOfferCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OpenseaOfferCallerRaw struct {
	Contract *OpenseaOfferCaller // Generic read-only contract binding to access the raw methods on
}

// OpenseaOfferTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OpenseaOfferTransactorRaw struct {
	Contract *OpenseaOfferTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOpenseaOffer creates a new instance of OpenseaOffer, bound to a specific deployed contract.
func NewOpenseaOffer(address common.Address, backend bind.ContractBackend) (*OpenseaOffer, error) {
	contract, err := bindOpenseaOffer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OpenseaOffer{OpenseaOfferCaller: OpenseaOfferCaller{contract: contract}, OpenseaOfferTransactor: OpenseaOfferTransactor{contract: contract}, OpenseaOfferFilterer: OpenseaOfferFilterer{contract: contract}}, nil
}

// NewOpenseaOfferCaller creates a new read-only instance of OpenseaOffer, bound to a specific deployed contract.
func NewOpenseaOfferCaller(address common.Address, caller bind.ContractCaller) (*OpenseaOfferCaller, error) {
	contract, err := bindOpenseaOffer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OpenseaOfferCaller{contract: contract}, nil
}

// NewOpenseaOfferTransactor creates a new write-only instance of OpenseaOffer, bound to a specific deployed contract.
func NewOpenseaOfferTransactor(address common.Address, transactor bind.ContractTransactor) (*OpenseaOfferTransactor, error) {
	contract, err := bindOpenseaOffer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OpenseaOfferTransactor{contract: contract}, nil
}

// NewOpenseaOfferFilterer creates a new log filterer instance of OpenseaOffer, bound to a specific deployed contract.
func NewOpenseaOfferFilterer(address common.Address, filterer bind.ContractFilterer) (*OpenseaOfferFilterer, error) {
	contract, err := bindOpenseaOffer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OpenseaOfferFilterer{contract: contract}, nil
}

// bindOpenseaOffer binds a generic wrapper to an already deployed contract.
func bindOpenseaOffer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OpenseaOfferABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OpenseaOffer *OpenseaOfferRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OpenseaOffer.Contract.OpenseaOfferCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OpenseaOffer *OpenseaOfferRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpenseaOffer.Contract.OpenseaOfferTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OpenseaOffer *OpenseaOfferRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OpenseaOffer.Contract.OpenseaOfferTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OpenseaOffer *OpenseaOfferCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OpenseaOffer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OpenseaOffer *OpenseaOfferTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpenseaOffer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OpenseaOffer *OpenseaOfferTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OpenseaOffer.Contract.contract.Transact(opts, method, params...)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_OpenseaOffer *OpenseaOfferCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OpenseaOffer.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_OpenseaOffer *OpenseaOfferSession) DomainSeparator() ([32]byte, error) {
	return _OpenseaOffer.Contract.DomainSeparator(&_OpenseaOffer.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_OpenseaOffer *OpenseaOfferCallerSession) DomainSeparator() ([32]byte, error) {
	return _OpenseaOffer.Contract.DomainSeparator(&_OpenseaOffer.CallOpts)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_OpenseaOffer *OpenseaOfferCaller) IsValidSignature(opts *bind.CallOpts, _hash [32]byte, _signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _OpenseaOffer.contract.Call(opts, &out, "isValidSignature", _hash, _signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_OpenseaOffer *OpenseaOfferSession) IsValidSignature(_hash [32]byte, _signature []byte) ([4]byte, error) {
	return _OpenseaOffer.Contract.IsValidSignature(&_OpenseaOffer.CallOpts, _hash, _signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_OpenseaOffer *OpenseaOfferCallerSession) IsValidSignature(_hash [32]byte, _signature []byte) ([4]byte, error) {
	return _OpenseaOffer.Contract.IsValidSignature(&_OpenseaOffer.CallOpts, _hash, _signature)
}

// Offers is a free data retrieval call binding the contract method 0x474d3ff0.
//
// Solidity: function offers(bytes32 ) view returns(string otaKey, address signer, bytes signature, uint256 startTime, uint256 endTime, uint256 offerAmount)
func (_OpenseaOffer *OpenseaOfferCaller) Offers(opts *bind.CallOpts, arg0 [32]byte) (struct {
	OtaKey      string
	Signer      common.Address
	Signature   []byte
	StartTime   *big.Int
	EndTime     *big.Int
	OfferAmount *big.Int
}, error) {
	var out []interface{}
	err := _OpenseaOffer.contract.Call(opts, &out, "offers", arg0)

	outstruct := new(struct {
		OtaKey      string
		Signer      common.Address
		Signature   []byte
		StartTime   *big.Int
		EndTime     *big.Int
		OfferAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OtaKey = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Signer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Signature = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.StartTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.OfferAmount = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Offers is a free data retrieval call binding the contract method 0x474d3ff0.
//
// Solidity: function offers(bytes32 ) view returns(string otaKey, address signer, bytes signature, uint256 startTime, uint256 endTime, uint256 offerAmount)
func (_OpenseaOffer *OpenseaOfferSession) Offers(arg0 [32]byte) (struct {
	OtaKey      string
	Signer      common.Address
	Signature   []byte
	StartTime   *big.Int
	EndTime     *big.Int
	OfferAmount *big.Int
}, error) {
	return _OpenseaOffer.Contract.Offers(&_OpenseaOffer.CallOpts, arg0)
}

// Offers is a free data retrieval call binding the contract method 0x474d3ff0.
//
// Solidity: function offers(bytes32 ) view returns(string otaKey, address signer, bytes signature, uint256 startTime, uint256 endTime, uint256 offerAmount)
func (_OpenseaOffer *OpenseaOfferCallerSession) Offers(arg0 [32]byte) (struct {
	OtaKey      string
	Signer      common.Address
	Signature   []byte
	StartTime   *big.Int
	EndTime     *big.Int
	OfferAmount *big.Int
}, error) {
	return _OpenseaOffer.Contract.Offers(&_OpenseaOffer.CallOpts, arg0)
}

// Seaport is a free data retrieval call binding the contract method 0xf5c7bd70.
//
// Solidity: function seaport() view returns(address)
func (_OpenseaOffer *OpenseaOfferCaller) Seaport(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OpenseaOffer.contract.Call(opts, &out, "seaport")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Seaport is a free data retrieval call binding the contract method 0xf5c7bd70.
//
// Solidity: function seaport() view returns(address)
func (_OpenseaOffer *OpenseaOfferSession) Seaport() (common.Address, error) {
	return _OpenseaOffer.Contract.Seaport(&_OpenseaOffer.CallOpts)
}

// Seaport is a free data retrieval call binding the contract method 0xf5c7bd70.
//
// Solidity: function seaport() view returns(address)
func (_OpenseaOffer *OpenseaOfferCallerSession) Seaport() (common.Address, error) {
	return _OpenseaOffer.Contract.Seaport(&_OpenseaOffer.CallOpts)
}

// ToTypedDataHash is a free data retrieval call binding the contract method 0x7df7a71c.
//
// Solidity: function toTypedDataHash(bytes32 domainSeparator_, bytes32 structHash_) pure returns(bytes32)
func (_OpenseaOffer *OpenseaOfferCaller) ToTypedDataHash(opts *bind.CallOpts, domainSeparator_ [32]byte, structHash_ [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _OpenseaOffer.contract.Call(opts, &out, "toTypedDataHash", domainSeparator_, structHash_)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ToTypedDataHash is a free data retrieval call binding the contract method 0x7df7a71c.
//
// Solidity: function toTypedDataHash(bytes32 domainSeparator_, bytes32 structHash_) pure returns(bytes32)
func (_OpenseaOffer *OpenseaOfferSession) ToTypedDataHash(domainSeparator_ [32]byte, structHash_ [32]byte) ([32]byte, error) {
	return _OpenseaOffer.Contract.ToTypedDataHash(&_OpenseaOffer.CallOpts, domainSeparator_, structHash_)
}

// ToTypedDataHash is a free data retrieval call binding the contract method 0x7df7a71c.
//
// Solidity: function toTypedDataHash(bytes32 domainSeparator_, bytes32 structHash_) pure returns(bytes32)
func (_OpenseaOffer *OpenseaOfferCallerSession) ToTypedDataHash(domainSeparator_ [32]byte, structHash_ [32]byte) ([32]byte, error) {
	return _OpenseaOffer.Contract.ToTypedDataHash(&_OpenseaOffer.CallOpts, domainSeparator_, structHash_)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_OpenseaOffer *OpenseaOfferCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OpenseaOffer.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_OpenseaOffer *OpenseaOfferSession) Vault() (common.Address, error) {
	return _OpenseaOffer.Contract.Vault(&_OpenseaOffer.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_OpenseaOffer *OpenseaOfferCallerSession) Vault() (common.Address, error) {
	return _OpenseaOffer.Contract.Vault(&_OpenseaOffer.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_OpenseaOffer *OpenseaOfferCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OpenseaOffer.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_OpenseaOffer *OpenseaOfferSession) Weth() (common.Address, error) {
	return _OpenseaOffer.Contract.Weth(&_OpenseaOffer.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_OpenseaOffer *OpenseaOfferCallerSession) Weth() (common.Address, error) {
	return _OpenseaOffer.Contract.Weth(&_OpenseaOffer.CallOpts)
}

// CancelOffer is a paid mutator transaction binding the contract method 0xd34517ed.
//
// Solidity: function cancelOffer((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256) order, bytes offerSignature, bytes32 txId, bytes regulatorSignData) returns()
func (_OpenseaOffer *OpenseaOfferTransactor) CancelOffer(opts *bind.TransactOpts, order OrderComponents, offerSignature []byte, txId [32]byte, regulatorSignData []byte) (*types.Transaction, error) {
	return _OpenseaOffer.contract.Transact(opts, "cancelOffer", order, offerSignature, txId, regulatorSignData)
}

// CancelOffer is a paid mutator transaction binding the contract method 0xd34517ed.
//
// Solidity: function cancelOffer((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256) order, bytes offerSignature, bytes32 txId, bytes regulatorSignData) returns()
func (_OpenseaOffer *OpenseaOfferSession) CancelOffer(order OrderComponents, offerSignature []byte, txId [32]byte, regulatorSignData []byte) (*types.Transaction, error) {
	return _OpenseaOffer.Contract.CancelOffer(&_OpenseaOffer.TransactOpts, order, offerSignature, txId, regulatorSignData)
}

// CancelOffer is a paid mutator transaction binding the contract method 0xd34517ed.
//
// Solidity: function cancelOffer((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256) order, bytes offerSignature, bytes32 txId, bytes regulatorSignData) returns()
func (_OpenseaOffer *OpenseaOfferTransactorSession) CancelOffer(order OrderComponents, offerSignature []byte, txId [32]byte, regulatorSignData []byte) (*types.Transaction, error) {
	return _OpenseaOffer.Contract.CancelOffer(&_OpenseaOffer.TransactOpts, order, offerSignature, txId, regulatorSignData)
}

// Offer is a paid mutator transaction binding the contract method 0xb4539661.
//
// Solidity: function offer((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256) order, string otaKey, bytes signature, address conduit) payable returns()
func (_OpenseaOffer *OpenseaOfferTransactor) Offer(opts *bind.TransactOpts, order OrderComponents, otaKey string, signature []byte, conduit common.Address) (*types.Transaction, error) {
	return _OpenseaOffer.contract.Transact(opts, "offer", order, otaKey, signature, conduit)
}

// Offer is a paid mutator transaction binding the contract method 0xb4539661.
//
// Solidity: function offer((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256) order, string otaKey, bytes signature, address conduit) payable returns()
func (_OpenseaOffer *OpenseaOfferSession) Offer(order OrderComponents, otaKey string, signature []byte, conduit common.Address) (*types.Transaction, error) {
	return _OpenseaOffer.Contract.Offer(&_OpenseaOffer.TransactOpts, order, otaKey, signature, conduit)
}

// Offer is a paid mutator transaction binding the contract method 0xb4539661.
//
// Solidity: function offer((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256) order, string otaKey, bytes signature, address conduit) payable returns()
func (_OpenseaOffer *OpenseaOfferTransactorSession) Offer(order OrderComponents, otaKey string, signature []byte, conduit common.Address) (*types.Transaction, error) {
	return _OpenseaOffer.Contract.Offer(&_OpenseaOffer.TransactOpts, order, otaKey, signature, conduit)
}
