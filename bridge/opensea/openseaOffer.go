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
	Bin: "0x608060405234801561001057600080fd5b506122d0806100206000396000f3fe60806040526004361061008a5760003560e01c8063b453966111610059578063b453966114610174578063d34517ed14610189578063f5c7bd70146101a9578063f698da25146101cc578063fbfa77cf1461020057600080fd5b80631626ba7e146100965780633fc8cef3146100d4578063474d3ff0146101145780637df7a71c1461014657600080fd5b3661009157005b600080fd5b3480156100a257600080fd5b506100b66100b1366004611476565b610228565b6040516001600160e01b031990911681526020015b60405180910390f35b3480156100e057600080fd5b506100fc73b4fbf271143f4fbf7b91a5ded31805e42b2208d681565b6040516001600160a01b0390911681526020016100cb565b34801561012057600080fd5b5061013461012f3660046114c1565b61025a565b6040516100cb96959493929190611520565b34801561015257600080fd5b50610166610161366004611573565b6103ad565b6040519081526020016100cb565b6101876101823660046115d6565b6103ee565b005b34801561019557600080fd5b506101876101a436600461167c565b61097c565b3480156101b557600080fd5b506100fc6e6c3852cbef3e08e8df289169ede58181565b3480156101d857600080fd5b506101667f712fdde1f147adcbb0fabb1aeb9c2d317530a46d266db095bc40c606fe94f0c281565b34801561020c57600080fd5b506100fc73c157cc3077ddfa425bae12d2f3002668971a4e3d81565b6000610235848484610f82565b156102485750630b135d3f60e11b610253565b506001600160e01b03195b9392505050565b60006020819052908152604090208054819061027590611719565b80601f01602080910402602001604051908101604052809291908181526020018280546102a190611719565b80156102ee5780601f106102c3576101008083540402835291602001916102ee565b820191906000526020600020905b8154815290600101906020018083116102d157829003601f168201915b505050600184015460028501805494956001600160a01b0390921694919350915061031890611719565b80601f016020809104026020016040519081016040528092919081815260200182805461034490611719565b80156103915780601f1061036657610100808354040283529160200191610391565b820191906000526020600020905b81548152906001019060200180831161037457829003601f168201915b5050505050908060030154908060040154908060050154905086565b60405161190160f01b602082015260228101839052604281018290526000906062016040516020818303038152906040528051906020012090505b92915050565b306103fc602088018861174d565b6001600160a01b0316148015610420575061041a604087018761176a565b90506001145b6104715760405162461bcd60e51b815260206004820152601d60248201527f4f70656e7365614f666665723a20696e76616c6964206f66666572657200000060448201526064015b60405180910390fd5b6040516379df72bd60e01b815260009061050d907f712fdde1f147adcbb0fabb1aeb9c2d317530a46d266db095bc40c606fe94f0c2906e6c3852cbef3e08e8df289169ede581906379df72bd906104cc908c906004016119a0565b602060405180830381865afa1580156104e9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101619190611a99565b6000818152602081905260409020600301549091501561056f5760405162461bcd60e51b815260206004820152601b60248201527f4f70656e7365614f666665723a206f66666572206578697374656400000000006044820152606401610468565b60006105b18286868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061115a92505050565b90506000805b6105c460408b018b61176a565b90508160ff1610156106955760006105df60408c018c61176a565b8360ff168181106105f2576105f2611ab2565b905060a002018036038101906106089190611bd8565b905080608001518160600151146106705760405162461bcd60e51b815260206004820152602660248201527f4f70656e7365614f666665723a20696e76616c696420737461727420656e6420604482015265185b5bdd5b9d60d21b6064820152608401610468565b606081015161067f9084611c0a565b925050808061068d90611c1d565b9150506105b7565b508034146106f05760405162461bcd60e51b815260206004820152602260248201527f4f70656e7365614f666665723a20696e76616c6964206f6666657220616d6f756044820152611b9d60f21b6064820152608401610468565b6040518060c0016040528089898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505050908252506001600160a01b03841660208083019190915260408051601f8a018390048302810183018252898152920191908990899081908401838280828437600092018290525093855250505060a08c013560208084019190915260c08d0135604080850191909152606090930185905286825281905220815181906107b69082611c8b565b5060208201516001820180546001600160a01b0319166001600160a01b03909216919091179055604082015160028201906107f19082611c8b565b50606082015181600301556080820151816004015560a0820151816005015590505073b4fbf271143f4fbf7b91a5ded31805e42b2208d66001600160a01b031663d0e30db0346040518263ffffffff1660e01b81526004016000604051808303818588803b15801561086257600080fd5b505af1158015610876573d6000803e3d6000fd5b5050604051636eb1769f60e11b81523060048201526001600160a01b038816602482015273b4fbf271143f4fbf7b91a5ded31805e42b2208d6935063095ea7b392508791508490849063dd62ed3e90604401602060405180830381865afa1580156108e5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109099190611a99565b6109139190611c0a565b6040516001600160e01b031960e085901b1681526001600160a01b0390921660048301526024820152604401600060405180830381600087803b15801561095957600080fd5b505af115801561096d573d6000803e3d6000fd5b50505050505050505050505050565b6040516379df72bd60e01b81526000906e6c3852cbef3e08e8df289169ede581906379df72bd906109b1908a906004016119a0565b602060405180830381865afa1580156109ce573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109f29190611a99565b6040516346423aa760e01b815260048101829052909150600090819081906e6c3852cbef3e08e8df289169ede581906346423aa790602401608060405180830381865afa158015610a47573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a6b9190611d5a565b935093509350506000806000610aa47f712fdde1f147adcbb0fabb1aeb9c2d317530a46d266db095bc40c606fe94f0c260001b886103ad565b81526020019081526020016000206040518060c0016040529081600082018054610acd90611719565b80601f0160208091040260200160405190810160405280929190818152602001828054610af990611719565b8015610b465780601f10610b1b57610100808354040283529160200191610b46565b820191906000526020600020905b815481529060010190602001808311610b2957829003601f168201915b505050918352505060018201546001600160a01b03166020820152600282018054604090920191610b7690611719565b80601f0160208091040260200160405190810160405280929190818152602001828054610ba290611719565b8015610bef5780601f10610bc457610100808354040283529160200191610bef565b820191906000526020600020905b815481529060010190602001808311610bd257829003601f168201915b50505050508152602001600382015481526020016004820154815260200160058201548152505090508060600151600014158015610c30575060a081015115155b8015610c3a575083155b8015610c4557508282115b610c915760405162461bcd60e51b815260206004820152601b60248201527f4f70656e7365614f666665723a20696e76616c6964206f6666657200000000006044820152606401610468565b8060800151421015610d4157610cdd858b8b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061115a92505050565b6001600160a01b031681602001516001600160a01b031614610d415760405162461bcd60e51b815260206004820152601f60248201527f4f70656e7365614f666665723a20696e76616c6964207369676e6174757265006044820152606401610468565b604080516001808252818301909252600091816020015b610d606113b1565b815260200190600190039081610d58579050509050610d7e8c611ef3565b81600081518110610d9157610d91611ab2565b6020908102919091010152604051630fd9f1e160e41b81526e6c3852cbef3e08e8df289169ede5819063fd9f1e1090610dce9084906004016120b4565b6020604051808303816000875af1158015610ded573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e1191906121b6565b610e745760405162461bcd60e51b815260206004820152602e60248201527f4f70656e7365614f666665723a20657865637574652063616e63656c206f6e2060448201526d1cd9585c1bdc9d0819985a5b195960921b6064820152608401610468565b600083610e8186826121d1565b8460a00151610e9091906121e4565b610e9a91906121fb565b604051632e1a7d4d60e01b81526004810182905290915073b4fbf271143f4fbf7b91a5ded31805e42b2208d690632e1a7d4d90602401600060405180830381600087803b158015610eea57600080fd5b505af1158015610efe573d6000803e3d6000fd5b5050845160405163c791d70560e01b815273c157cc3077ddfa425bae12d2f3002668971a4e3d935063c791d70592508491610f41918f908f908f9060040161221d565b6000604051808303818588803b158015610f5a57600080fd5b505af1158015610f6e573d6000803e3d6000fd5b505050505050505050505050505050505050565b6000806000808681526020019081526020016000206040518060c0016040529081600082018054610fb290611719565b80601f0160208091040260200160405190810160405280929190818152602001828054610fde90611719565b801561102b5780601f106110005761010080835404028352916020019161102b565b820191906000526020600020905b81548152906001019060200180831161100e57829003601f168201915b505050918352505060018201546001600160a01b0316602082015260028201805460409092019161105b90611719565b80601f016020809104026020016040519081016040528092919081815260200182805461108790611719565b80156110d45780601f106110a9576101008083540402835291602001916110d4565b820191906000526020600020905b8154815290600101906020018083116110b757829003601f168201915b50505050508152602001600382015481526020016004820154815260200160058201548152505090508060600151600014806111325750838360405161111b92919061226a565b604051809103902081604001518051906020012014155b806111405750428160800151105b1561114f576000915050610253565b506001949350505050565b600081516041146111c15760405162461bcd60e51b815260206004820152603a602482015260008051602061227b83398151915260448201527f3a20696e76616c6964207369676e6174757265206c656e6774680000000000006064820152608401610468565b6000826040815181106111d6576111d6611ab2565b0160209081015190840151604085015160f89290921c92507f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08211156112725760405162461bcd60e51b815260206004820152603d602482015260008051602061227b83398151915260448201527f3a20696e76616c6964207369676e6174757265202773272076616c75650000006064820152608401610468565b8260ff16601b1415801561128a57508260ff16601c14155b156112eb5760405162461bcd60e51b815260206004820152603d602482015260008051602061227b83398151915260448201527f3a20696e76616c6964207369676e6174757265202776272076616c75650000006064820152608401610468565b60408051600081526020810180835288905260ff851691810191909152606081018290526080810183905260019060a0016020604051602081039080840390855afa15801561133e573d6000803e3d6000fd5b5050604051601f1901519450506001600160a01b0384166113a85760405162461bcd60e51b8152602060048201526030602482015260008051602061227b83398151915260448201526f1d1024a72b20a624a22fa9a4a3a722a960811b6064820152608401610468565b50505092915050565b60405180610160016040528060006001600160a01b0316815260200160006001600160a01b031681526020016060815260200160608152602001600060038111156113fe576113fe611809565b815260006020820181905260408201819052606082018190526080820181905260a0820181905260c09091015290565b60008083601f84011261144057600080fd5b5081356001600160401b0381111561145757600080fd5b60208301915083602082850101111561146f57600080fd5b9250929050565b60008060006040848603121561148b57600080fd5b8335925060208401356001600160401b038111156114a857600080fd5b6114b48682870161142e565b9497909650939450505050565b6000602082840312156114d357600080fd5b5035919050565b6000815180845260005b81811015611500576020818501810151868301820152016114e4565b506000602082860101526020601f19601f83011685010191505092915050565b60c08152600061153360c08301896114da565b6001600160a01b0388166020840152828103604084015261155481886114da565b60608401969096525050608081019290925260a0909101529392505050565b6000806040838503121561158657600080fd5b50508035926020909101359150565b600061016082840312156115a857600080fd5b50919050565b6001600160a01b03811681146115c357600080fd5b50565b80356115d1816115ae565b919050565b600080600080600080608087890312156115ef57600080fd5b86356001600160401b038082111561160657600080fd5b6116128a838b01611595565b9750602089013591508082111561162857600080fd5b6116348a838b0161142e565b9097509550604089013591508082111561164d57600080fd5b5061165a89828a0161142e565b909450925050606087013561166e816115ae565b809150509295509295509295565b6000806000806000806080878903121561169557600080fd5b86356001600160401b03808211156116ac57600080fd5b6116b88a838b01611595565b975060208901359150808211156116ce57600080fd5b6116da8a838b0161142e565b90975095506040890135945060608901359150808211156116fa57600080fd5b5061170789828a0161142e565b979a9699509497509295939492505050565b600181811c9082168061172d57607f821691505b6020821081036115a857634e487b7160e01b600052602260045260246000fd5b60006020828403121561175f57600080fd5b8135610253816115ae565b6000808335601e1984360301811261178157600080fd5b8301803591506001600160401b0382111561179b57600080fd5b602001915060a08102360382131561146f57600080fd5b6000808335601e198436030181126117c957600080fd5b83016020810192503590506001600160401b038111156117e857600080fd5b60a08102360382131561146f57600080fd5b8035600681106115d157600080fd5b634e487b7160e01b600052602160045260246000fd5b6006811061182f5761182f611809565b9052565b8183526000602080850194508260005b858110156118a75761185d87611858846117fa565b61181f565b8282013561186a816115ae565b6001600160a01b03168388015260408281013590880152606080830135908801526080808301359088015260a09687019690910190600101611843565b509495945050505050565b6000808335601e198436030181126118c957600080fd5b83016020810192503590506001600160401b038111156118e857600080fd5b60c08102360382131561146f57600080fd5b8183526000602080850194508260005b858110156118a75761191f87611858846117fa565b8282013561192c816115ae565b6001600160a01b039081168885015260408381013590890152606080840135908901526080808401359089015260a09083820135611969816115ae565b169088015260c096870196919091019060010161190a565b8035600481106115d157600080fd5b6004811061182f5761182f611809565b602081526119c1602082016119b4846115c6565b6001600160a01b03169052565b60006119cf602084016115c6565b6001600160a01b0381166040840152506119ec60408401846117b2565b610160806060860152611a0461018086018385611833565b9250611a1360608701876118b2565b868503601f190160808801529250611a2c8484836118fa565b935050611a3b60808701611981565b9150611a4a60a0860183611990565b60a086013560c086015260c086013560e0860152610100915060e08601358286015261012082870135818701526101409250808701358387015250818601358186015250508091505092915050565b600060208284031215611aab57600080fd5b5051919050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b60405160c081016001600160401b0381118282101715611b0057611b00611ac8565b60405290565b60405161016081016001600160401b0381118282101715611b0057611b00611ac8565b604051601f8201601f191681016001600160401b0381118282101715611b5157611b51611ac8565b604052919050565b600060a08284031215611b6b57600080fd5b60405160a081018181106001600160401b0382111715611b8d57611b8d611ac8565b604052905080611b9c836117fa565b81526020830135611bac816115ae565b806020830152506040830135604082015260608301356060820152608083013560808201525092915050565b600060a08284031215611bea57600080fd5b6102538383611b59565b634e487b7160e01b600052601160045260246000fd5b808201808211156103e8576103e8611bf4565b600060ff821660ff8103611c3357611c33611bf4565b60010192915050565b601f821115611c8657600081815260208120601f850160051c81016020861015611c635750805b601f850160051c820191505b81811015611c8257828155600101611c6f565b5050505b505050565b81516001600160401b03811115611ca457611ca4611ac8565b611cb881611cb28454611719565b84611c3c565b602080601f831160018114611ced5760008415611cd55750858301515b600019600386901b1c1916600185901b178555611c82565b600085815260208120601f198616915b82811015611d1c57888601518255948401946001909101908401611cfd565b5085821015611d3a5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b805180151581146115d157600080fd5b60008060008060808587031215611d7057600080fd5b611d7985611d4a565b9350611d8760208601611d4a565b6040860151606090960151949790965092505050565b60006001600160401b03821115611db657611db6611ac8565b5060051b60200190565b600082601f830112611dd157600080fd5b81356020611de6611de183611d9d565b611b29565b82815260a09283028501820192828201919087851115611e0557600080fd5b8387015b85811015611e2857611e1b8982611b59565b8452928401928101611e09565b5090979650505050505050565b600082601f830112611e4657600080fd5b81356020611e56611de183611d9d565b82815260c09283028501820192828201919087851115611e7557600080fd5b8387015b85811015611e285781818a031215611e915760008081fd5b611e99611ade565b611ea2826117fa565b815285820135611eb1816115ae565b8187015260408281013590820152606080830135908201526080808301359082015260a080830135611ee2816115ae565b908201528452928401928101611e79565b60006101608236031215611f0657600080fd5b611f0e611b06565b611f17836115c6565b8152611f25602084016115c6565b602082015260408301356001600160401b0380821115611f4457600080fd5b611f5036838701611dc0565b60408401526060850135915080821115611f6957600080fd5b50611f7636828601611e35565b606083015250611f8860808401611981565b608082015260a0838101359082015260c0808401359082015260e080840135908201526101008084013590820152610120808401359082015261014092830135928101929092525090565b600081518084526020808501945080840160005b838110156118a7578151611ffc88825161181f565b838101516001600160a01b03168885015260408082015190890152606080820151908901526080908101519088015260a09096019590820190600101611fe7565b600081518084526020808501945080840160005b838110156118a757815161206688825161181f565b808401516001600160a01b0390811689860152604080830151908a0152606080830151908a0152608080830151908a015260a091820151169088015260c09096019590820190600101612051565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b838110156121a857888303603f19018552815180516001600160a01b03168452610160818901516001600160a01b038116868b01525087820151818987015261212682870182611fd3565b91505060608083015186830382880152612140838261203d565b9250505060808083015161215682880182611990565b505060a0828101519086015260c0808301519086015260e080830151908601526101008083015190860152610120808301519086015261014091820151919094015293860193908601906001016120db565b509098975050505050505050565b6000602082840312156121c857600080fd5b61025382611d4a565b818103818111156103e8576103e8611bf4565b80820281158282048414176103e8576103e8611bf4565b60008261221857634e487b7160e01b600052601260045260246000fd5b500490565b60608152600061223060608301876114da565b8560208401528281036040840152838152838560208301376000602085830101526020601f19601f86011682010191505095945050505050565b818382376000910190815291905056fe5369676e617475726556616c696461746f72237265636f7665725369676e6572a264697066735822122020f0f7b0a3f1a9b06761f4541b1dfb2985900c649851999c4c57dad341b3e90a64736f6c63430008110033",
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
