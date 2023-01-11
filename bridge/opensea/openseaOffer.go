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
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"counter\",\"type\":\"uint256\"}],\"internalType\":\"structOrderComponents\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"offerSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"txId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"regulatorSignData\",\"type\":\"bytes\"}],\"name\":\"cancelOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"counter\",\"type\":\"uint256\"}],\"internalType\":\"structOrderComponents\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"otaKey\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"offer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"offers\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"otaKey\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offerAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"seaport\",\"outputs\":[{\"internalType\":\"contractConsiderationInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"domainSeparator_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"structHash_\",\"type\":\"bytes32\"}],\"name\":\"toTypedDataHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"contractVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"contractWETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b506122b6806100206000396000f3fe60806040526004361061008a5760003560e01c8063b453966111610059578063b453966114610174578063d34517ed14610189578063f5c7bd70146101a9578063f698da25146101cc578063fbfa77cf1461020057600080fd5b80631626ba7e146100965780633fc8cef3146100d4578063474d3ff0146101145780637df7a71c1461014657600080fd5b3661009157005b600080fd5b3480156100a257600080fd5b506100b66100b1366004611492565b610228565b6040516001600160e01b031990911681526020015b60405180910390f35b3480156100e057600080fd5b506100fc73b4fbf271143f4fbf7b91a5ded31805e42b2208d681565b6040516001600160a01b0390911681526020016100cb565b34801561012057600080fd5b5061013461012f3660046114dd565b61025a565b6040516100cb9695949392919061153c565b34801561015257600080fd5b5061016661016136600461158f565b6103ad565b6040519081526020016100cb565b6101876101823660046115f2565b6103ee565b005b34801561019557600080fd5b506101876101a4366004611698565b61098b565b3480156101b557600080fd5b506100fc6e6c3852cbef3e08e8df289169ede58181565b3480156101d857600080fd5b506101667f712fdde1f147adcbb0fabb1aeb9c2d317530a46d266db095bc40c606fe94f0c281565b34801561020c57600080fd5b506100fc73c157cc3077ddfa425bae12d2f3002668971a4e3d81565b6000610235848484610f9e565b156102485750630b135d3f60e11b610253565b506001600160e01b03195b9392505050565b60006020819052908152604090208054819061027590611735565b80601f01602080910402602001604051908101604052809291908181526020018280546102a190611735565b80156102ee5780601f106102c3576101008083540402835291602001916102ee565b820191906000526020600020905b8154815290600101906020018083116102d157829003601f168201915b505050600184015460028501805494956001600160a01b0390921694919350915061031890611735565b80601f016020809104026020016040519081016040528092919081815260200182805461034490611735565b80156103915780601f1061036657610100808354040283529160200191610391565b820191906000526020600020905b81548152906001019060200180831161037457829003601f168201915b5050505050908060030154908060040154908060050154905086565b60405161190160f01b602082015260228101839052604281018290526000906062016040516020818303038152906040528051906020012090505b92915050565b306103fc6020880188611769565b6001600160a01b0316148015610420575061041a6040870187611786565b90506001145b801561042f575060a086013515155b6104805760405162461bcd60e51b815260206004820152601d60248201527f4f70656e7365614f666665723a20696e76616c6964206f66666572657200000060448201526064015b60405180910390fd5b6040516379df72bd60e01b815260009061051c907f712fdde1f147adcbb0fabb1aeb9c2d317530a46d266db095bc40c606fe94f0c2906e6c3852cbef3e08e8df289169ede581906379df72bd906104db908c906004016119bc565b602060405180830381865afa1580156104f8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101619190611ab5565b6000818152602081905260409020600301549091501561057e5760405162461bcd60e51b815260206004820152601b60248201527f4f70656e7365614f666665723a206f66666572206578697374656400000000006044820152606401610477565b60006105c08286868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061117692505050565b90506000805b6105d360408b018b611786565b90508160ff1610156106a45760006105ee60408c018c611786565b8360ff1681811061060157610601611ace565b905060a002018036038101906106179190611bf4565b9050806080015181606001511461067f5760405162461bcd60e51b815260206004820152602660248201527f4f70656e7365614f666665723a20696e76616c696420737461727420656e6420604482015265185b5bdd5b9d60d21b6064820152608401610477565b606081015161068e9084611c26565b925050808061069c90611c39565b9150506105c6565b508034146106ff5760405162461bcd60e51b815260206004820152602260248201527f4f70656e7365614f666665723a20696e76616c6964206f6666657220616d6f756044820152611b9d60f21b6064820152608401610477565b6040518060c0016040528089898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505050908252506001600160a01b03841660208083019190915260408051601f8a018390048302810183018252898152920191908990899081908401838280828437600092018290525093855250505060a08c013560208084019190915260c08d0135604080850191909152606090930185905286825281905220815181906107c59082611ca7565b5060208201516001820180546001600160a01b0319166001600160a01b03909216919091179055604082015160028201906108009082611ca7565b50606082015181600301556080820151816004015560a0820151816005015590505073b4fbf271143f4fbf7b91a5ded31805e42b2208d66001600160a01b031663d0e30db0346040518263ffffffff1660e01b81526004016000604051808303818588803b15801561087157600080fd5b505af1158015610885573d6000803e3d6000fd5b5050604051636eb1769f60e11b81523060048201526001600160a01b038816602482015273b4fbf271143f4fbf7b91a5ded31805e42b2208d6935063095ea7b392508791508490849063dd62ed3e90604401602060405180830381865afa1580156108f4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109189190611ab5565b6109229190611c26565b6040516001600160e01b031960e085901b1681526001600160a01b0390921660048301526024820152604401600060405180830381600087803b15801561096857600080fd5b505af115801561097c573d6000803e3d6000fd5b50505050505050505050505050565b6040516379df72bd60e01b81526000906e6c3852cbef3e08e8df289169ede581906379df72bd906109c0908a906004016119bc565b602060405180830381865afa1580156109dd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a019190611ab5565b6040516346423aa760e01b815260048101829052909150600090819081906e6c3852cbef3e08e8df289169ede581906346423aa790602401608060405180830381865afa158015610a56573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a7a9190611d76565b935093509350506000806000610ab37f712fdde1f147adcbb0fabb1aeb9c2d317530a46d266db095bc40c606fe94f0c260001b886103ad565b81526020019081526020016000206040518060c0016040529081600082018054610adc90611735565b80601f0160208091040260200160405190810160405280929190818152602001828054610b0890611735565b8015610b555780601f10610b2a57610100808354040283529160200191610b55565b820191906000526020600020905b815481529060010190602001808311610b3857829003601f168201915b505050918352505060018201546001600160a01b03166020820152600282018054604090920191610b8590611735565b80601f0160208091040260200160405190810160405280929190818152602001828054610bb190611735565b8015610bfe5780601f10610bd357610100808354040283529160200191610bfe565b820191906000526020600020905b815481529060010190602001808311610be157829003601f168201915b50505050508152602001600382015481526020016004820154815260200160058201548152505090508060600151600014158015610c3f575060a081015115155b8015610c49575083155b8015610c5c575082821180610c5c575081155b610ca85760405162461bcd60e51b815260206004820152601b60248201527f4f70656e7365614f666665723a20696e76616c6964206f6666657200000000006044820152606401610477565b8060800151421015610d5857610cf4858b8b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061117692505050565b6001600160a01b031681602001516001600160a01b031614610d585760405162461bcd60e51b815260206004820152601f60248201527f4f70656e7365614f666665723a20696e76616c6964207369676e6174757265006044820152606401610477565b604080516001808252818301909252600091816020015b610d776113cd565b815260200190600190039081610d6f579050509050610d958c611f0f565b81600081518110610da857610da8611ace565b6020908102919091010152604051630fd9f1e160e41b81526e6c3852cbef3e08e8df289169ede5819063fd9f1e1090610de59084906004016120d0565b6020604051808303816000875af1158015610e04573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e2891906121d2565b610e8b5760405162461bcd60e51b815260206004820152602e60248201527f4f70656e7365614f666665723a20657865637574652063616e63656c206f6e2060448201526d1cd9585c1bdc9d0819985a5b195960921b6064820152608401610477565b60008315610eb057838585038460a001510281610eaa57610eaa6121ed565b04610eb6565b8260a001515b604051632e1a7d4d60e01b81526004810182905290915073b4fbf271143f4fbf7b91a5ded31805e42b2208d690632e1a7d4d90602401600060405180830381600087803b158015610f0657600080fd5b505af1158015610f1a573d6000803e3d6000fd5b5050845160405163c791d70560e01b815273c157cc3077ddfa425bae12d2f3002668971a4e3d935063c791d70592508491610f5d918f908f908f90600401612203565b6000604051808303818588803b158015610f7657600080fd5b505af1158015610f8a573d6000803e3d6000fd5b505050505050505050505050505050505050565b6000806000808681526020019081526020016000206040518060c0016040529081600082018054610fce90611735565b80601f0160208091040260200160405190810160405280929190818152602001828054610ffa90611735565b80156110475780601f1061101c57610100808354040283529160200191611047565b820191906000526020600020905b81548152906001019060200180831161102a57829003601f168201915b505050918352505060018201546001600160a01b0316602082015260028201805460409092019161107790611735565b80601f01602080910402602001604051908101604052809291908181526020018280546110a390611735565b80156110f05780601f106110c5576101008083540402835291602001916110f0565b820191906000526020600020905b8154815290600101906020018083116110d357829003601f168201915b505050505081526020016003820154815260200160048201548152602001600582015481525050905080606001516000148061114e57508383604051611137929190612250565b604051809103902081604001518051906020012014155b8061115c5750428160800151105b1561116b576000915050610253565b506001949350505050565b600081516041146111dd5760405162461bcd60e51b815260206004820152603a602482015260008051602061226183398151915260448201527f3a20696e76616c6964207369676e6174757265206c656e6774680000000000006064820152608401610477565b6000826040815181106111f2576111f2611ace565b0160209081015190840151604085015160f89290921c92507f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a082111561128e5760405162461bcd60e51b815260206004820152603d602482015260008051602061226183398151915260448201527f3a20696e76616c6964207369676e6174757265202773272076616c75650000006064820152608401610477565b8260ff16601b141580156112a657508260ff16601c14155b156113075760405162461bcd60e51b815260206004820152603d602482015260008051602061226183398151915260448201527f3a20696e76616c6964207369676e6174757265202776272076616c75650000006064820152608401610477565b60408051600081526020810180835288905260ff851691810191909152606081018290526080810183905260019060a0016020604051602081039080840390855afa15801561135a573d6000803e3d6000fd5b5050604051601f1901519450506001600160a01b0384166113c45760405162461bcd60e51b8152602060048201526030602482015260008051602061226183398151915260448201526f1d1024a72b20a624a22fa9a4a3a722a960811b6064820152608401610477565b50505092915050565b60405180610160016040528060006001600160a01b0316815260200160006001600160a01b0316815260200160608152602001606081526020016000600381111561141a5761141a611825565b815260006020820181905260408201819052606082018190526080820181905260a0820181905260c09091015290565b60008083601f84011261145c57600080fd5b5081356001600160401b0381111561147357600080fd5b60208301915083602082850101111561148b57600080fd5b9250929050565b6000806000604084860312156114a757600080fd5b8335925060208401356001600160401b038111156114c457600080fd5b6114d08682870161144a565b9497909650939450505050565b6000602082840312156114ef57600080fd5b5035919050565b6000815180845260005b8181101561151c57602081850181015186830182015201611500565b506000602082860101526020601f19601f83011685010191505092915050565b60c08152600061154f60c08301896114f6565b6001600160a01b0388166020840152828103604084015261157081886114f6565b60608401969096525050608081019290925260a0909101529392505050565b600080604083850312156115a257600080fd5b50508035926020909101359150565b600061016082840312156115c457600080fd5b50919050565b6001600160a01b03811681146115df57600080fd5b50565b80356115ed816115ca565b919050565b6000806000806000806080878903121561160b57600080fd5b86356001600160401b038082111561162257600080fd5b61162e8a838b016115b1565b9750602089013591508082111561164457600080fd5b6116508a838b0161144a565b9097509550604089013591508082111561166957600080fd5b5061167689828a0161144a565b909450925050606087013561168a816115ca565b809150509295509295509295565b600080600080600080608087890312156116b157600080fd5b86356001600160401b03808211156116c857600080fd5b6116d48a838b016115b1565b975060208901359150808211156116ea57600080fd5b6116f68a838b0161144a565b909750955060408901359450606089013591508082111561171657600080fd5b5061172389828a0161144a565b979a9699509497509295939492505050565b600181811c9082168061174957607f821691505b6020821081036115c457634e487b7160e01b600052602260045260246000fd5b60006020828403121561177b57600080fd5b8135610253816115ca565b6000808335601e1984360301811261179d57600080fd5b8301803591506001600160401b038211156117b757600080fd5b602001915060a08102360382131561148b57600080fd5b6000808335601e198436030181126117e557600080fd5b83016020810192503590506001600160401b0381111561180457600080fd5b60a08102360382131561148b57600080fd5b8035600681106115ed57600080fd5b634e487b7160e01b600052602160045260246000fd5b6006811061184b5761184b611825565b9052565b8183526000602080850194508260005b858110156118c3576118798761187484611816565b61183b565b82820135611886816115ca565b6001600160a01b03168388015260408281013590880152606080830135908801526080808301359088015260a0968701969091019060010161185f565b509495945050505050565b6000808335601e198436030181126118e557600080fd5b83016020810192503590506001600160401b0381111561190457600080fd5b60c08102360382131561148b57600080fd5b8183526000602080850194508260005b858110156118c35761193b8761187484611816565b82820135611948816115ca565b6001600160a01b039081168885015260408381013590890152606080840135908901526080808401359089015260a09083820135611985816115ca565b169088015260c0968701969190910190600101611926565b8035600481106115ed57600080fd5b6004811061184b5761184b611825565b602081526119dd602082016119d0846115e2565b6001600160a01b03169052565b60006119eb602084016115e2565b6001600160a01b038116604084015250611a0860408401846117ce565b610160806060860152611a206101808601838561184f565b9250611a2f60608701876118ce565b868503601f190160808801529250611a48848483611916565b935050611a576080870161199d565b9150611a6660a08601836119ac565b60a086013560c086015260c086013560e0860152610100915060e08601358286015261012082870135818701526101409250808701358387015250818601358186015250508091505092915050565b600060208284031215611ac757600080fd5b5051919050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b60405160c081016001600160401b0381118282101715611b1c57611b1c611ae4565b60405290565b60405161016081016001600160401b0381118282101715611b1c57611b1c611ae4565b604051601f8201601f191681016001600160401b0381118282101715611b6d57611b6d611ae4565b604052919050565b600060a08284031215611b8757600080fd5b60405160a081018181106001600160401b0382111715611ba957611ba9611ae4565b604052905080611bb883611816565b81526020830135611bc8816115ca565b806020830152506040830135604082015260608301356060820152608083013560808201525092915050565b600060a08284031215611c0657600080fd5b6102538383611b75565b634e487b7160e01b600052601160045260246000fd5b808201808211156103e8576103e8611c10565b600060ff821660ff8103611c4f57611c4f611c10565b60010192915050565b601f821115611ca257600081815260208120601f850160051c81016020861015611c7f5750805b601f850160051c820191505b81811015611c9e57828155600101611c8b565b5050505b505050565b81516001600160401b03811115611cc057611cc0611ae4565b611cd481611cce8454611735565b84611c58565b602080601f831160018114611d095760008415611cf15750858301515b600019600386901b1c1916600185901b178555611c9e565b600085815260208120601f198616915b82811015611d3857888601518255948401946001909101908401611d19565b5085821015611d565787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b805180151581146115ed57600080fd5b60008060008060808587031215611d8c57600080fd5b611d9585611d66565b9350611da360208601611d66565b6040860151606090960151949790965092505050565b60006001600160401b03821115611dd257611dd2611ae4565b5060051b60200190565b600082601f830112611ded57600080fd5b81356020611e02611dfd83611db9565b611b45565b82815260a09283028501820192828201919087851115611e2157600080fd5b8387015b85811015611e4457611e378982611b75565b8452928401928101611e25565b5090979650505050505050565b600082601f830112611e6257600080fd5b81356020611e72611dfd83611db9565b82815260c09283028501820192828201919087851115611e9157600080fd5b8387015b85811015611e445781818a031215611ead5760008081fd5b611eb5611afa565b611ebe82611816565b815285820135611ecd816115ca565b8187015260408281013590820152606080830135908201526080808301359082015260a080830135611efe816115ca565b908201528452928401928101611e95565b60006101608236031215611f2257600080fd5b611f2a611b22565b611f33836115e2565b8152611f41602084016115e2565b602082015260408301356001600160401b0380821115611f6057600080fd5b611f6c36838701611ddc565b60408401526060850135915080821115611f8557600080fd5b50611f9236828601611e51565b606083015250611fa46080840161199d565b608082015260a0838101359082015260c0808401359082015260e080840135908201526101008084013590820152610120808401359082015261014092830135928101929092525090565b600081518084526020808501945080840160005b838110156118c357815161201888825161183b565b838101516001600160a01b03168885015260408082015190890152606080820151908901526080908101519088015260a09096019590820190600101612003565b600081518084526020808501945080840160005b838110156118c357815161208288825161183b565b808401516001600160a01b0390811689860152604080830151908a0152606080830151908a0152608080830151908a015260a091820151169088015260c0909601959082019060010161206d565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b838110156121c457888303603f19018552815180516001600160a01b03168452610160818901516001600160a01b038116868b01525087820151818987015261214282870182611fef565b9150506060808301518683038288015261215c8382612059565b92505050608080830151612172828801826119ac565b505060a0828101519086015260c0808301519086015260e080830151908601526101008083015190860152610120808301519086015261014091820151919094015293860193908601906001016120f7565b509098975050505050505050565b6000602082840312156121e457600080fd5b61025382611d66565b634e487b7160e01b600052601260045260246000fd5b60608152600061221660608301876114f6565b8560208401528281036040840152838152838560208301376000602085830101526020601f19601f86011682010191505095945050505050565b818382376000910190815291905056fe5369676e617475726556616c696461746f72237265636f7665725369676e6572a264697066735822122093ecc746abfa95bb7970736afe2755b668fdfdc7d61848867a72a310375d773464736f6c63430008110033",
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

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OpenseaOffer *OpenseaOfferTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpenseaOffer.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OpenseaOffer *OpenseaOfferSession) Receive() (*types.Transaction, error) {
	return _OpenseaOffer.Contract.Receive(&_OpenseaOffer.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OpenseaOffer *OpenseaOfferTransactorSession) Receive() (*types.Transaction, error) {
	return _OpenseaOffer.Contract.Receive(&_OpenseaOffer.TransactOpts)
}
