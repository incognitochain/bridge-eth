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
	Bin: "0x608060405234801561001057600080fd5b506122a7806100206000396000f3fe60806040526004361061008a5760003560e01c8063b453966111610059578063b453966114610174578063d34517ed14610189578063f5c7bd70146101a9578063f698da25146101cc578063fbfa77cf1461020057600080fd5b80631626ba7e146100965780633fc8cef3146100d4578063474d3ff0146101145780637df7a71c1461014657600080fd5b3661009157005b600080fd5b3480156100a257600080fd5b506100b66100b1366004611483565b610228565b6040516001600160e01b031990911681526020015b60405180910390f35b3480156100e057600080fd5b506100fc73b4fbf271143f4fbf7b91a5ded31805e42b2208d681565b6040516001600160a01b0390911681526020016100cb565b34801561012057600080fd5b5061013461012f3660046114ce565b61025a565b6040516100cb9695949392919061152d565b34801561015257600080fd5b50610166610161366004611580565b6103ad565b6040519081526020016100cb565b6101876101823660046115e3565b6103ee565b005b34801561019557600080fd5b506101876101a4366004611689565b61097c565b3480156101b557600080fd5b506100fc6e6c3852cbef3e08e8df289169ede58181565b3480156101d857600080fd5b506101667f712fdde1f147adcbb0fabb1aeb9c2d317530a46d266db095bc40c606fe94f0c281565b34801561020c57600080fd5b506100fc73c157cc3077ddfa425bae12d2f3002668971a4e3d81565b6000610235848484610f8f565b156102485750630b135d3f60e11b610253565b506001600160e01b03195b9392505050565b60006020819052908152604090208054819061027590611726565b80601f01602080910402602001604051908101604052809291908181526020018280546102a190611726565b80156102ee5780601f106102c3576101008083540402835291602001916102ee565b820191906000526020600020905b8154815290600101906020018083116102d157829003601f168201915b505050600184015460028501805494956001600160a01b0390921694919350915061031890611726565b80601f016020809104026020016040519081016040528092919081815260200182805461034490611726565b80156103915780601f1061036657610100808354040283529160200191610391565b820191906000526020600020905b81548152906001019060200180831161037457829003601f168201915b5050505050908060030154908060040154908060050154905086565b60405161190160f01b602082015260228101839052604281018290526000906062016040516020818303038152906040528051906020012090505b92915050565b306103fc602088018861175a565b6001600160a01b0316148015610420575061041a6040870187611777565b90506001145b6104715760405162461bcd60e51b815260206004820152601d60248201527f4f70656e7365614f666665723a20696e76616c6964206f66666572657200000060448201526064015b60405180910390fd5b6040516379df72bd60e01b815260009061050d907f712fdde1f147adcbb0fabb1aeb9c2d317530a46d266db095bc40c606fe94f0c2906e6c3852cbef3e08e8df289169ede581906379df72bd906104cc908c906004016119ad565b602060405180830381865afa1580156104e9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101619190611aa6565b6000818152602081905260409020600301549091501561056f5760405162461bcd60e51b815260206004820152601b60248201527f4f70656e7365614f666665723a206f66666572206578697374656400000000006044820152606401610468565b60006105b18286868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061116792505050565b90506000805b6105c460408b018b611777565b90508160ff1610156106955760006105df60408c018c611777565b8360ff168181106105f2576105f2611abf565b905060a002018036038101906106089190611be5565b905080608001518160600151146106705760405162461bcd60e51b815260206004820152602660248201527f4f70656e7365614f666665723a20696e76616c696420737461727420656e6420604482015265185b5bdd5b9d60d21b6064820152608401610468565b606081015161067f9084611c17565b925050808061068d90611c2a565b9150506105b7565b508034146106f05760405162461bcd60e51b815260206004820152602260248201527f4f70656e7365614f666665723a20696e76616c6964206f6666657220616d6f756044820152611b9d60f21b6064820152608401610468565b6040518060c0016040528089898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505050908252506001600160a01b03841660208083019190915260408051601f8a018390048302810183018252898152920191908990899081908401838280828437600092018290525093855250505060a08c013560208084019190915260c08d0135604080850191909152606090930185905286825281905220815181906107b69082611c98565b5060208201516001820180546001600160a01b0319166001600160a01b03909216919091179055604082015160028201906107f19082611c98565b50606082015181600301556080820151816004015560a0820151816005015590505073b4fbf271143f4fbf7b91a5ded31805e42b2208d66001600160a01b031663d0e30db0346040518263ffffffff1660e01b81526004016000604051808303818588803b15801561086257600080fd5b505af1158015610876573d6000803e3d6000fd5b5050604051636eb1769f60e11b81523060048201526001600160a01b038816602482015273b4fbf271143f4fbf7b91a5ded31805e42b2208d6935063095ea7b392508791508490849063dd62ed3e90604401602060405180830381865afa1580156108e5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109099190611aa6565b6109139190611c17565b6040516001600160e01b031960e085901b1681526001600160a01b0390921660048301526024820152604401600060405180830381600087803b15801561095957600080fd5b505af115801561096d573d6000803e3d6000fd5b50505050505050505050505050565b6040516379df72bd60e01b81526000906e6c3852cbef3e08e8df289169ede581906379df72bd906109b1908a906004016119ad565b602060405180830381865afa1580156109ce573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109f29190611aa6565b6040516346423aa760e01b815260048101829052909150600090819081906e6c3852cbef3e08e8df289169ede581906346423aa790602401608060405180830381865afa158015610a47573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a6b9190611d67565b935093509350506000806000610aa47f712fdde1f147adcbb0fabb1aeb9c2d317530a46d266db095bc40c606fe94f0c260001b886103ad565b81526020019081526020016000206040518060c0016040529081600082018054610acd90611726565b80601f0160208091040260200160405190810160405280929190818152602001828054610af990611726565b8015610b465780601f10610b1b57610100808354040283529160200191610b46565b820191906000526020600020905b815481529060010190602001808311610b2957829003601f168201915b505050918352505060018201546001600160a01b03166020820152600282018054604090920191610b7690611726565b80601f0160208091040260200160405190810160405280929190818152602001828054610ba290611726565b8015610bef5780601f10610bc457610100808354040283529160200191610bef565b820191906000526020600020905b815481529060010190602001808311610bd257829003601f168201915b50505050508152602001600382015481526020016004820154815260200160058201548152505090508060600151600014158015610c30575060a081015115155b8015610c3a575083155b8015610c4d575082821180610c4d575081155b610c995760405162461bcd60e51b815260206004820152601b60248201527f4f70656e7365614f666665723a20696e76616c6964206f6666657200000000006044820152606401610468565b8060800151421015610d4957610ce5858b8b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061116792505050565b6001600160a01b031681602001516001600160a01b031614610d495760405162461bcd60e51b815260206004820152601f60248201527f4f70656e7365614f666665723a20696e76616c6964207369676e6174757265006044820152606401610468565b604080516001808252818301909252600091816020015b610d686113be565b815260200190600190039081610d60579050509050610d868c611f00565b81600081518110610d9957610d99611abf565b6020908102919091010152604051630fd9f1e160e41b81526e6c3852cbef3e08e8df289169ede5819063fd9f1e1090610dd69084906004016120c1565b6020604051808303816000875af1158015610df5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e1991906121c3565b610e7c5760405162461bcd60e51b815260206004820152602e60248201527f4f70656e7365614f666665723a20657865637574652063616e63656c206f6e2060448201526d1cd9585c1bdc9d0819985a5b195960921b6064820152608401610468565b60008315610ea157838585038460a001510281610e9b57610e9b6121de565b04610ea7565b8260a001515b604051632e1a7d4d60e01b81526004810182905290915073b4fbf271143f4fbf7b91a5ded31805e42b2208d690632e1a7d4d90602401600060405180830381600087803b158015610ef757600080fd5b505af1158015610f0b573d6000803e3d6000fd5b5050845160405163c791d70560e01b815273c157cc3077ddfa425bae12d2f3002668971a4e3d935063c791d70592508491610f4e918f908f908f906004016121f4565b6000604051808303818588803b158015610f6757600080fd5b505af1158015610f7b573d6000803e3d6000fd5b505050505050505050505050505050505050565b6000806000808681526020019081526020016000206040518060c0016040529081600082018054610fbf90611726565b80601f0160208091040260200160405190810160405280929190818152602001828054610feb90611726565b80156110385780601f1061100d57610100808354040283529160200191611038565b820191906000526020600020905b81548152906001019060200180831161101b57829003601f168201915b505050918352505060018201546001600160a01b0316602082015260028201805460409092019161106890611726565b80601f016020809104026020016040519081016040528092919081815260200182805461109490611726565b80156110e15780601f106110b6576101008083540402835291602001916110e1565b820191906000526020600020905b8154815290600101906020018083116110c457829003601f168201915b505050505081526020016003820154815260200160048201548152602001600582015481525050905080606001516000148061113f57508383604051611128929190612241565b604051809103902081604001518051906020012014155b8061114d5750428160800151105b1561115c576000915050610253565b506001949350505050565b600081516041146111ce5760405162461bcd60e51b815260206004820152603a602482015260008051602061225283398151915260448201527f3a20696e76616c6964207369676e6174757265206c656e6774680000000000006064820152608401610468565b6000826040815181106111e3576111e3611abf565b0160209081015190840151604085015160f89290921c92507f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a082111561127f5760405162461bcd60e51b815260206004820152603d602482015260008051602061225283398151915260448201527f3a20696e76616c6964207369676e6174757265202773272076616c75650000006064820152608401610468565b8260ff16601b1415801561129757508260ff16601c14155b156112f85760405162461bcd60e51b815260206004820152603d602482015260008051602061225283398151915260448201527f3a20696e76616c6964207369676e6174757265202776272076616c75650000006064820152608401610468565b60408051600081526020810180835288905260ff851691810191909152606081018290526080810183905260019060a0016020604051602081039080840390855afa15801561134b573d6000803e3d6000fd5b5050604051601f1901519450506001600160a01b0384166113b55760405162461bcd60e51b8152602060048201526030602482015260008051602061225283398151915260448201526f1d1024a72b20a624a22fa9a4a3a722a960811b6064820152608401610468565b50505092915050565b60405180610160016040528060006001600160a01b0316815260200160006001600160a01b0316815260200160608152602001606081526020016000600381111561140b5761140b611816565b815260006020820181905260408201819052606082018190526080820181905260a0820181905260c09091015290565b60008083601f84011261144d57600080fd5b5081356001600160401b0381111561146457600080fd5b60208301915083602082850101111561147c57600080fd5b9250929050565b60008060006040848603121561149857600080fd5b8335925060208401356001600160401b038111156114b557600080fd5b6114c18682870161143b565b9497909650939450505050565b6000602082840312156114e057600080fd5b5035919050565b6000815180845260005b8181101561150d576020818501810151868301820152016114f1565b506000602082860101526020601f19601f83011685010191505092915050565b60c08152600061154060c08301896114e7565b6001600160a01b0388166020840152828103604084015261156181886114e7565b60608401969096525050608081019290925260a0909101529392505050565b6000806040838503121561159357600080fd5b50508035926020909101359150565b600061016082840312156115b557600080fd5b50919050565b6001600160a01b03811681146115d057600080fd5b50565b80356115de816115bb565b919050565b600080600080600080608087890312156115fc57600080fd5b86356001600160401b038082111561161357600080fd5b61161f8a838b016115a2565b9750602089013591508082111561163557600080fd5b6116418a838b0161143b565b9097509550604089013591508082111561165a57600080fd5b5061166789828a0161143b565b909450925050606087013561167b816115bb565b809150509295509295509295565b600080600080600080608087890312156116a257600080fd5b86356001600160401b03808211156116b957600080fd5b6116c58a838b016115a2565b975060208901359150808211156116db57600080fd5b6116e78a838b0161143b565b909750955060408901359450606089013591508082111561170757600080fd5b5061171489828a0161143b565b979a9699509497509295939492505050565b600181811c9082168061173a57607f821691505b6020821081036115b557634e487b7160e01b600052602260045260246000fd5b60006020828403121561176c57600080fd5b8135610253816115bb565b6000808335601e1984360301811261178e57600080fd5b8301803591506001600160401b038211156117a857600080fd5b602001915060a08102360382131561147c57600080fd5b6000808335601e198436030181126117d657600080fd5b83016020810192503590506001600160401b038111156117f557600080fd5b60a08102360382131561147c57600080fd5b8035600681106115de57600080fd5b634e487b7160e01b600052602160045260246000fd5b6006811061183c5761183c611816565b9052565b8183526000602080850194508260005b858110156118b45761186a8761186584611807565b61182c565b82820135611877816115bb565b6001600160a01b03168388015260408281013590880152606080830135908801526080808301359088015260a09687019690910190600101611850565b509495945050505050565b6000808335601e198436030181126118d657600080fd5b83016020810192503590506001600160401b038111156118f557600080fd5b60c08102360382131561147c57600080fd5b8183526000602080850194508260005b858110156118b45761192c8761186584611807565b82820135611939816115bb565b6001600160a01b039081168885015260408381013590890152606080840135908901526080808401359089015260a09083820135611976816115bb565b169088015260c0968701969190910190600101611917565b8035600481106115de57600080fd5b6004811061183c5761183c611816565b602081526119ce602082016119c1846115d3565b6001600160a01b03169052565b60006119dc602084016115d3565b6001600160a01b0381166040840152506119f960408401846117bf565b610160806060860152611a1161018086018385611840565b9250611a2060608701876118bf565b868503601f190160808801529250611a39848483611907565b935050611a486080870161198e565b9150611a5760a086018361199d565b60a086013560c086015260c086013560e0860152610100915060e08601358286015261012082870135818701526101409250808701358387015250818601358186015250508091505092915050565b600060208284031215611ab857600080fd5b5051919050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b60405160c081016001600160401b0381118282101715611b0d57611b0d611ad5565b60405290565b60405161016081016001600160401b0381118282101715611b0d57611b0d611ad5565b604051601f8201601f191681016001600160401b0381118282101715611b5e57611b5e611ad5565b604052919050565b600060a08284031215611b7857600080fd5b60405160a081018181106001600160401b0382111715611b9a57611b9a611ad5565b604052905080611ba983611807565b81526020830135611bb9816115bb565b806020830152506040830135604082015260608301356060820152608083013560808201525092915050565b600060a08284031215611bf757600080fd5b6102538383611b66565b634e487b7160e01b600052601160045260246000fd5b808201808211156103e8576103e8611c01565b600060ff821660ff8103611c4057611c40611c01565b60010192915050565b601f821115611c9357600081815260208120601f850160051c81016020861015611c705750805b601f850160051c820191505b81811015611c8f57828155600101611c7c565b5050505b505050565b81516001600160401b03811115611cb157611cb1611ad5565b611cc581611cbf8454611726565b84611c49565b602080601f831160018114611cfa5760008415611ce25750858301515b600019600386901b1c1916600185901b178555611c8f565b600085815260208120601f198616915b82811015611d2957888601518255948401946001909101908401611d0a565b5085821015611d475787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b805180151581146115de57600080fd5b60008060008060808587031215611d7d57600080fd5b611d8685611d57565b9350611d9460208601611d57565b6040860151606090960151949790965092505050565b60006001600160401b03821115611dc357611dc3611ad5565b5060051b60200190565b600082601f830112611dde57600080fd5b81356020611df3611dee83611daa565b611b36565b82815260a09283028501820192828201919087851115611e1257600080fd5b8387015b85811015611e3557611e288982611b66565b8452928401928101611e16565b5090979650505050505050565b600082601f830112611e5357600080fd5b81356020611e63611dee83611daa565b82815260c09283028501820192828201919087851115611e8257600080fd5b8387015b85811015611e355781818a031215611e9e5760008081fd5b611ea6611aeb565b611eaf82611807565b815285820135611ebe816115bb565b8187015260408281013590820152606080830135908201526080808301359082015260a080830135611eef816115bb565b908201528452928401928101611e86565b60006101608236031215611f1357600080fd5b611f1b611b13565b611f24836115d3565b8152611f32602084016115d3565b602082015260408301356001600160401b0380821115611f5157600080fd5b611f5d36838701611dcd565b60408401526060850135915080821115611f7657600080fd5b50611f8336828601611e42565b606083015250611f956080840161198e565b608082015260a0838101359082015260c0808401359082015260e080840135908201526101008084013590820152610120808401359082015261014092830135928101929092525090565b600081518084526020808501945080840160005b838110156118b457815161200988825161182c565b838101516001600160a01b03168885015260408082015190890152606080820151908901526080908101519088015260a09096019590820190600101611ff4565b600081518084526020808501945080840160005b838110156118b457815161207388825161182c565b808401516001600160a01b0390811689860152604080830151908a0152606080830151908a0152608080830151908a015260a091820151169088015260c0909601959082019060010161205e565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b838110156121b557888303603f19018552815180516001600160a01b03168452610160818901516001600160a01b038116868b01525087820151818987015261213382870182611fe0565b9150506060808301518683038288015261214d838261204a565b925050506080808301516121638288018261199d565b505060a0828101519086015260c0808301519086015260e080830151908601526101008083015190860152610120808301519086015261014091820151919094015293860193908601906001016120e8565b509098975050505050505050565b6000602082840312156121d557600080fd5b61025382611d57565b634e487b7160e01b600052601260045260246000fd5b60608152600061220760608301876114e7565b8560208401528281036040840152838152838560208301376000602085830101526020601f19601f86011682010191505095945050505050565b818382376000910190815291905056fe5369676e617475726556616c696461746f72237265636f7665725369676e6572a2646970667358221220125ba56e62663f76f84d0e0172d9f8a1f73e001c76f6f6b15dd3662756dacba664736f6c63430008110033",
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
