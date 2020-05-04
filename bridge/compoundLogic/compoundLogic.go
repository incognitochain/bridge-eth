// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package compoundLogic

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CompoundLogicABI is the input ABI used to generate the binding from.
const CompoundLogicABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ETH_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"addToMarkets\",\"type\":\"address[]\"}],\"name\":\"borrow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"addToMarkets\",\"type\":\"address[]\"}],\"name\":\"borrowByMultiCollateral\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cEther\",\"outputs\":[{\"internalType\":\"contractCEther\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"comptroller\",\"outputs\":[{\"internalType\":\"contractComptroller\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incognitoSmartContract\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"}],\"name\":\"liquidateBorrow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isredeemUnderlying\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"exitToMarkets\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"repayBorrow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// CompoundLogicBin is the compiled bytecode used for deploying new contracts.
var CompoundLogicBin = "0x608060405234801561001057600080fd5b506120af806100206000396000f3fe6080604052600436106100955760003560e01c806372e94bf61161005957806372e94bf6146103b05780637520f7ed14610407578063abdb5ea8146104d5578063b42a644b1461056a578063b6dbc8ed146105c15761009c565b806319b68c00146100a157806326d5e251146100f857806340c10f19146101ef5780635fe3b5671461028457806364fd7078146102db5761009c565b3661009c57005b600080fd5b3480156100ad57600080fd5b506100b661070e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561010457600080fd5b506101a66004803603606081101561011b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561016257600080fd5b82018360208201111561017457600080fd5b8035906020019184602083028401116401000000008311171561019657600080fd5b9091929391929390505050610726565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b61023b6004803603604081101561020557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610b59565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b34801561029057600080fd5b50610299610d64565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610367600480360360808110156102f157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610d7c565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b3480156103bc57600080fd5b506103c561105c565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561041357600080fd5b5061048c6004803603608081101561042a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803515159060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611061565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b610521600480360360408110156104eb57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506114c5565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b34801561057657600080fd5b5061057f6116af565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156105cd57600080fd5b5061066f600480360360608110156105e457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561062b57600080fd5b82018360208201111561063d57600080fd5b8035906020019184602083028401116401000000008311171561065f57600080fd5b90919293919293905050506116c7565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b838110156106b657808201518184015260208101905061069b565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156106f85780820151818401526020810190506106dd565b5050505090500194505050505060405180910390f35b73f92fbe0d3c0dcdae407923b2ac17ec223b1084e481565b60008060008484905011156108a257731f5d7f3caac149fe41b8bd62a3673fe6ec0ab73b73ffffffffffffffffffffffffffffffffffffffff1663c299823885856040518363ffffffff1660e01b815260040180806020018281038252848482818152602001925060200280828437600081840152601f19601f8201169050808301925050509350505050600060405180830381600087803b1580156107cb57600080fd5b505af11580156107df573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250602081101561080957600080fd5b810190808051604051939291908464010000000082111561082957600080fd5b8382019150602082018581111561083f57600080fd5b825186602082028301116401000000008211171561085c57600080fd5b8083526020830192505050908051906020019060200280838360005b83811015610893578082015181840152602081019050610878565b50505050905001604052505050505b60008673ffffffffffffffffffffffffffffffffffffffff1663c5ebeaec876040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b1580156108f757600080fd5b505af115801561090b573d6000803e3d6000fd5b505050506040513d602081101561092157600080fd5b81019080805190602001909291905050501461093c57600080fd5b600080905073f92fbe0d3c0dcdae407923b2ac17ec223b1084e473ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff1614156109ad576109936000611c37565b90506109a0600082611d35565b6000819250925050610b50565b610a388773ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156109f857600080fd5b505af1158015610a0c573d6000803e3d6000fd5b505050506040513d6020811015610a2257600080fd5b8101908080519060200190929190505050611c37565b9050610ac68773ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b8152600401602060405180830381600087803b158015610a8557600080fd5b505af1158015610a99573d6000803e3d6000fd5b505050506040513d6020811015610aaf57600080fd5b810190808051906020019092919050505082611d35565b8673ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b8152600401602060405180830381600087803b158015610b0e57600080fd5b505af1158015610b22573d6000803e3d6000fd5b505050506040513d6020811015610b3857600080fd5b81019080805190602001909291905050508192509250505b94509492505050565b6000806000610b6785611c37565b905073f92fbe0d3c0dcdae407923b2ac17ec223b1084e473ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff161415610c17578473ffffffffffffffffffffffffffffffffffffffff16631249c58b346040518263ffffffff1660e01b81526004016000604051808303818588803b158015610bf957600080fd5b505af1158015610c0d573d6000803e3d6000fd5b5050505050610d3f565b610ca48573ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b8152600401602060405180830381600087803b158015610c6257600080fd5b505af1158015610c76573d6000803e3d6000fd5b505050506040513d6020811015610c8c57600080fd5b81019080805190602001909291905050508686611ea0565b60008573ffffffffffffffffffffffffffffffffffffffff1663a0712d68866040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b158015610cf957600080fd5b505af1158015610d0d573d6000803e3d6000fd5b505050506040513d6020811015610d2357600080fd5b810190808051906020019092919050505014610d3e57600080fd5b5b80610d4986611c37565b039050610d568582611d35565b848192509250509250929050565b731f5d7f3caac149fe41b8bd62a3673fe6ec0ab73b81565b6000806000610d8a84611c37565b905073f92fbe0d3c0dcdae407923b2ac17ec223b1084e473ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff161415610ea5578673ffffffffffffffffffffffffffffffffffffffff1663aae40a2a3488876040518463ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001925050506000604051808303818588803b158015610e8757600080fd5b505af1158015610e9b573d6000803e3d6000fd5b5050505050611035565b610f328773ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b8152600401602060405180830381600087803b158015610ef057600080fd5b505af1158015610f04573d6000803e3d6000fd5b505050506040513d6020811015610f1a57600080fd5b81019080805190602001909291905050508887611ea0565b60008773ffffffffffffffffffffffffffffffffffffffff1663f5e3c4628888886040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019350505050602060405180830381600087803b158015610fef57600080fd5b505af1158015611003573d6000803e3d6000fd5b505050506040513d602081101561101957600080fd5b81019080805190602001909291905050501461103457600080fd5b5b8061103f85611c37565b03905061104c8482611d35565b8381925092505094509492505050565b600081565b600080600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161461116857731f5d7f3caac149fe41b8bd62a3673fe6ec0ab73b73ffffffffffffffffffffffffffffffffffffffff1663ede4edd0846040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b15801561112b57600080fd5b505af115801561113f573d6000803e3d6000fd5b505050506040513d602081101561115557600080fd5b8101908080519060200190929190505050505b831561120d5760008673ffffffffffffffffffffffffffffffffffffffff1663852a12e3876040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b1580156111c357600080fd5b505af11580156111d7573d6000803e3d6000fd5b505050506040513d60208110156111ed57600080fd5b81019080805190602001909291905050501461120857600080fd5b6112a8565b60008673ffffffffffffffffffffffffffffffffffffffff1663db006a75876040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b15801561126257600080fd5b505af1158015611276573d6000803e3d6000fd5b505050506040513d602081101561128c57600080fd5b8101908080519060200190929190505050146112a757600080fd5b5b600080905073f92fbe0d3c0dcdae407923b2ac17ec223b1084e473ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff161415611319576112ff6000611c37565b905061130c600082611d35565b60008192509250506114bc565b6113a48773ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561136457600080fd5b505af1158015611378573d6000803e3d6000fd5b505050506040513d602081101561138e57600080fd5b8101908080519060200190929190505050611c37565b90506114328773ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156113f157600080fd5b505af1158015611405573d6000803e3d6000fd5b505050506040513d602081101561141b57600080fd5b810190808051906020019092919050505082611d35565b8673ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561147a57600080fd5b505af115801561148e573d6000803e3d6000fd5b505050506040513d60208110156114a457600080fd5b81019080805190602001909291905050508192509250505b94509492505050565b60008073f92fbe0d3c0dcdae407923b2ac17ec223b1084e473ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415611576578373ffffffffffffffffffffffffffffffffffffffff16634e4d9fea346040518263ffffffff1660e01b81526004016000604051808303818588803b15801561155857600080fd5b505af115801561156c573d6000803e3d6000fd5b505050505061169e565b6116038473ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156115c157600080fd5b505af11580156115d5573d6000803e3d6000fd5b505050506040513d60208110156115eb57600080fd5b81019080805190602001909291905050508585611ea0565b60008473ffffffffffffffffffffffffffffffffffffffff16630e752702856040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b15801561165857600080fd5b505af115801561166c573d6000803e3d6000fd5b505050506040513d602081101561168257600080fd5b81019080805190602001909291905050501461169d57600080fd5b5b836000809050915091509250929050565b738c4b922e2f7d6d1aba41d79c47f497f6f54e0af881565b606080600084849050111561184357731f5d7f3caac149fe41b8bd62a3673fe6ec0ab73b73ffffffffffffffffffffffffffffffffffffffff1663c299823885856040518363ffffffff1660e01b815260040180806020018281038252848482818152602001925060200280828437600081840152601f19601f8201169050808301925050509350505050600060405180830381600087803b15801561176c57600080fd5b505af1158015611780573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525060208110156117aa57600080fd5b81019080805160405193929190846401000000008211156117ca57600080fd5b838201915060208201858111156117e057600080fd5b82518660208202830111640100000000821117156117fd57600080fd5b8083526020830192505050908051906020019060200280838360005b83811015611834578082015181840152602081019050611819565b50505050905001604052505050505b6060600167ffffffffffffffff8111801561185d57600080fd5b5060405190808252806020026020018201604052801561188c5781602001602082028036833780820191505090505b50905060008773ffffffffffffffffffffffffffffffffffffffff1663c5ebeaec886040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b1580156118e457600080fd5b505af11580156118f8573d6000803e3d6000fd5b505050506040513d602081101561190e57600080fd5b81019080805190602001909291905050501461192957600080fd5b600080905073f92fbe0d3c0dcdae407923b2ac17ec223b1084e473ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff1614156119db576119806000611c37565b905061198d600082611d35565b60008260008151811061199c57fe5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050611bbf565b611a668873ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b8152600401602060405180830381600087803b158015611a2657600080fd5b505af1158015611a3a573d6000803e3d6000fd5b505050506040513d6020811015611a5057600080fd5b8101908080519060200190929190505050611c37565b9050611af48873ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b8152600401602060405180830381600087803b158015611ab357600080fd5b505af1158015611ac7573d6000803e3d6000fd5b505050506040513d6020811015611add57600080fd5b810190808051906020019092919050505082611d35565b8773ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b8152600401602060405180830381600087803b158015611b3c57600080fd5b505af1158015611b50573d6000803e3d6000fd5b505050506040513d6020811015611b6657600080fd5b810190808051906020019092919050505082600081518110611b8457fe5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250505b6060600167ffffffffffffffff81118015611bd957600080fd5b50604051908082528060200260200182016040528015611c085781602001602082028036833780820191505090505b5090508181600081518110611c1957fe5b60200260200101818152505082819450945050505094509492505050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611c7557479050611d30565b8173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015611cf257600080fd5b505afa158015611d06573d6000803e3d6000fd5b505050506040513d6020811015611d1c57600080fd5b810190808051906020019092919050505090505b919050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611dd75780471015611d7757600080fd5b738c4b922e2f7d6d1aba41d79c47f497f6f54e0af873ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015611dd1573d6000803e3d6000fd5b50611e9c565b8173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb738c4b922e2f7d6d1aba41d79c47f497f6f54e0af8836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b158015611e7257600080fd5b505af1158015611e86573d6000803e3d6000fd5b50505050611e9261203b565b611e9b57600080fd5b5b5050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614612036578273ffffffffffffffffffffffffffffffffffffffff1663095ea7b38360006040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b158015611f5c57600080fd5b505af1158015611f70573d6000803e3d6000fd5b50505050611f7c61203b565b611f8557600080fd5b8273ffffffffffffffffffffffffffffffffffffffff1663095ea7b383836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b15801561200c57600080fd5b505af1158015612020573d6000803e3d6000fd5b5050505061202c61203b565b61203557600080fd5b5b505050565b600080600090503d6000811461205857602081146120615761206d565b6001915061206d565b60206000803e60005191505b5060008114159150509056fea2646970667358221220967659ac861952a9aa0050995c174f4a8403162a7e76b358899bb954b12a859664736f6c63430006060033"

// DeployCompoundLogic deploys a new Ethereum contract, binding an instance of CompoundLogic to it.
func DeployCompoundLogic(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CompoundLogic, error) {
	parsed, err := abi.JSON(strings.NewReader(CompoundLogicABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CompoundLogicBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CompoundLogic{CompoundLogicCaller: CompoundLogicCaller{contract: contract}, CompoundLogicTransactor: CompoundLogicTransactor{contract: contract}, CompoundLogicFilterer: CompoundLogicFilterer{contract: contract}}, nil
}

// CompoundLogic is an auto generated Go binding around an Ethereum contract.
type CompoundLogic struct {
	CompoundLogicCaller     // Read-only binding to the contract
	CompoundLogicTransactor // Write-only binding to the contract
	CompoundLogicFilterer   // Log filterer for contract events
}

// CompoundLogicCaller is an auto generated read-only Go binding around an Ethereum contract.
type CompoundLogicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundLogicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CompoundLogicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundLogicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CompoundLogicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundLogicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CompoundLogicSession struct {
	Contract     *CompoundLogic    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CompoundLogicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CompoundLogicCallerSession struct {
	Contract *CompoundLogicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// CompoundLogicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CompoundLogicTransactorSession struct {
	Contract     *CompoundLogicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CompoundLogicRaw is an auto generated low-level Go binding around an Ethereum contract.
type CompoundLogicRaw struct {
	Contract *CompoundLogic // Generic contract binding to access the raw methods on
}

// CompoundLogicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CompoundLogicCallerRaw struct {
	Contract *CompoundLogicCaller // Generic read-only contract binding to access the raw methods on
}

// CompoundLogicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CompoundLogicTransactorRaw struct {
	Contract *CompoundLogicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompoundLogic creates a new instance of CompoundLogic, bound to a specific deployed contract.
func NewCompoundLogic(address common.Address, backend bind.ContractBackend) (*CompoundLogic, error) {
	contract, err := bindCompoundLogic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CompoundLogic{CompoundLogicCaller: CompoundLogicCaller{contract: contract}, CompoundLogicTransactor: CompoundLogicTransactor{contract: contract}, CompoundLogicFilterer: CompoundLogicFilterer{contract: contract}}, nil
}

// NewCompoundLogicCaller creates a new read-only instance of CompoundLogic, bound to a specific deployed contract.
func NewCompoundLogicCaller(address common.Address, caller bind.ContractCaller) (*CompoundLogicCaller, error) {
	contract, err := bindCompoundLogic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CompoundLogicCaller{contract: contract}, nil
}

// NewCompoundLogicTransactor creates a new write-only instance of CompoundLogic, bound to a specific deployed contract.
func NewCompoundLogicTransactor(address common.Address, transactor bind.ContractTransactor) (*CompoundLogicTransactor, error) {
	contract, err := bindCompoundLogic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CompoundLogicTransactor{contract: contract}, nil
}

// NewCompoundLogicFilterer creates a new log filterer instance of CompoundLogic, bound to a specific deployed contract.
func NewCompoundLogicFilterer(address common.Address, filterer bind.ContractFilterer) (*CompoundLogicFilterer, error) {
	contract, err := bindCompoundLogic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CompoundLogicFilterer{contract: contract}, nil
}

// bindCompoundLogic binds a generic wrapper to an already deployed contract.
func bindCompoundLogic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CompoundLogicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompoundLogic *CompoundLogicRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CompoundLogic.Contract.CompoundLogicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompoundLogic *CompoundLogicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompoundLogic.Contract.CompoundLogicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompoundLogic *CompoundLogicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompoundLogic.Contract.CompoundLogicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompoundLogic *CompoundLogicCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CompoundLogic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompoundLogic *CompoundLogicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompoundLogic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompoundLogic *CompoundLogicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompoundLogic.Contract.contract.Transact(opts, method, params...)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_CompoundLogic *CompoundLogicCaller) ETHCONTRACTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CompoundLogic.contract.Call(opts, out, "ETH_CONTRACT_ADDRESS")
	return *ret0, err
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_CompoundLogic *CompoundLogicSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _CompoundLogic.Contract.ETHCONTRACTADDRESS(&_CompoundLogic.CallOpts)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_CompoundLogic *CompoundLogicCallerSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _CompoundLogic.Contract.ETHCONTRACTADDRESS(&_CompoundLogic.CallOpts)
}

// CEther is a free data retrieval call binding the contract method 0x19b68c00.
//
// Solidity: function cEther() constant returns(address)
func (_CompoundLogic *CompoundLogicCaller) CEther(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CompoundLogic.contract.Call(opts, out, "cEther")
	return *ret0, err
}

// CEther is a free data retrieval call binding the contract method 0x19b68c00.
//
// Solidity: function cEther() constant returns(address)
func (_CompoundLogic *CompoundLogicSession) CEther() (common.Address, error) {
	return _CompoundLogic.Contract.CEther(&_CompoundLogic.CallOpts)
}

// CEther is a free data retrieval call binding the contract method 0x19b68c00.
//
// Solidity: function cEther() constant returns(address)
func (_CompoundLogic *CompoundLogicCallerSession) CEther() (common.Address, error) {
	return _CompoundLogic.Contract.CEther(&_CompoundLogic.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() constant returns(address)
func (_CompoundLogic *CompoundLogicCaller) Comptroller(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CompoundLogic.contract.Call(opts, out, "comptroller")
	return *ret0, err
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() constant returns(address)
func (_CompoundLogic *CompoundLogicSession) Comptroller() (common.Address, error) {
	return _CompoundLogic.Contract.Comptroller(&_CompoundLogic.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() constant returns(address)
func (_CompoundLogic *CompoundLogicCallerSession) Comptroller() (common.Address, error) {
	return _CompoundLogic.Contract.Comptroller(&_CompoundLogic.CallOpts)
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_CompoundLogic *CompoundLogicCaller) IncognitoSmartContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CompoundLogic.contract.Call(opts, out, "incognitoSmartContract")
	return *ret0, err
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_CompoundLogic *CompoundLogicSession) IncognitoSmartContract() (common.Address, error) {
	return _CompoundLogic.Contract.IncognitoSmartContract(&_CompoundLogic.CallOpts)
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_CompoundLogic *CompoundLogicCallerSession) IncognitoSmartContract() (common.Address, error) {
	return _CompoundLogic.Contract.IncognitoSmartContract(&_CompoundLogic.CallOpts)
}

// Borrow is a paid mutator transaction binding the contract method 0x26d5e251.
//
// Solidity: function borrow(address cToken, uint256 amount, address[] addToMarkets) returns(address, uint256)
func (_CompoundLogic *CompoundLogicTransactor) Borrow(opts *bind.TransactOpts, cToken common.Address, amount *big.Int, addToMarkets []common.Address) (*types.Transaction, error) {
	return _CompoundLogic.contract.Transact(opts, "borrow", cToken, amount, addToMarkets)
}

// Borrow is a paid mutator transaction binding the contract method 0x26d5e251.
//
// Solidity: function borrow(address cToken, uint256 amount, address[] addToMarkets) returns(address, uint256)
func (_CompoundLogic *CompoundLogicSession) Borrow(cToken common.Address, amount *big.Int, addToMarkets []common.Address) (*types.Transaction, error) {
	return _CompoundLogic.Contract.Borrow(&_CompoundLogic.TransactOpts, cToken, amount, addToMarkets)
}

// Borrow is a paid mutator transaction binding the contract method 0x26d5e251.
//
// Solidity: function borrow(address cToken, uint256 amount, address[] addToMarkets) returns(address, uint256)
func (_CompoundLogic *CompoundLogicTransactorSession) Borrow(cToken common.Address, amount *big.Int, addToMarkets []common.Address) (*types.Transaction, error) {
	return _CompoundLogic.Contract.Borrow(&_CompoundLogic.TransactOpts, cToken, amount, addToMarkets)
}

// BorrowByMultiCollateral is a paid mutator transaction binding the contract method 0xb6dbc8ed.
//
// Solidity: function borrowByMultiCollateral(address cToken, uint256 amount, address[] addToMarkets) returns(address[], uint256[])
func (_CompoundLogic *CompoundLogicTransactor) BorrowByMultiCollateral(opts *bind.TransactOpts, cToken common.Address, amount *big.Int, addToMarkets []common.Address) (*types.Transaction, error) {
	return _CompoundLogic.contract.Transact(opts, "borrowByMultiCollateral", cToken, amount, addToMarkets)
}

// BorrowByMultiCollateral is a paid mutator transaction binding the contract method 0xb6dbc8ed.
//
// Solidity: function borrowByMultiCollateral(address cToken, uint256 amount, address[] addToMarkets) returns(address[], uint256[])
func (_CompoundLogic *CompoundLogicSession) BorrowByMultiCollateral(cToken common.Address, amount *big.Int, addToMarkets []common.Address) (*types.Transaction, error) {
	return _CompoundLogic.Contract.BorrowByMultiCollateral(&_CompoundLogic.TransactOpts, cToken, amount, addToMarkets)
}

// BorrowByMultiCollateral is a paid mutator transaction binding the contract method 0xb6dbc8ed.
//
// Solidity: function borrowByMultiCollateral(address cToken, uint256 amount, address[] addToMarkets) returns(address[], uint256[])
func (_CompoundLogic *CompoundLogicTransactorSession) BorrowByMultiCollateral(cToken common.Address, amount *big.Int, addToMarkets []common.Address) (*types.Transaction, error) {
	return _CompoundLogic.Contract.BorrowByMultiCollateral(&_CompoundLogic.TransactOpts, cToken, amount, addToMarkets)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0x64fd7078.
//
// Solidity: function liquidateBorrow(address cToken, address borrower, uint256 repayAmount, address cTokenCollateral) returns(address, uint256)
func (_CompoundLogic *CompoundLogicTransactor) LiquidateBorrow(opts *bind.TransactOpts, cToken common.Address, borrower common.Address, repayAmount *big.Int, cTokenCollateral common.Address) (*types.Transaction, error) {
	return _CompoundLogic.contract.Transact(opts, "liquidateBorrow", cToken, borrower, repayAmount, cTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0x64fd7078.
//
// Solidity: function liquidateBorrow(address cToken, address borrower, uint256 repayAmount, address cTokenCollateral) returns(address, uint256)
func (_CompoundLogic *CompoundLogicSession) LiquidateBorrow(cToken common.Address, borrower common.Address, repayAmount *big.Int, cTokenCollateral common.Address) (*types.Transaction, error) {
	return _CompoundLogic.Contract.LiquidateBorrow(&_CompoundLogic.TransactOpts, cToken, borrower, repayAmount, cTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0x64fd7078.
//
// Solidity: function liquidateBorrow(address cToken, address borrower, uint256 repayAmount, address cTokenCollateral) returns(address, uint256)
func (_CompoundLogic *CompoundLogicTransactorSession) LiquidateBorrow(cToken common.Address, borrower common.Address, repayAmount *big.Int, cTokenCollateral common.Address) (*types.Transaction, error) {
	return _CompoundLogic.Contract.LiquidateBorrow(&_CompoundLogic.TransactOpts, cToken, borrower, repayAmount, cTokenCollateral)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address cToken, uint256 amount) returns(address, uint256)
func (_CompoundLogic *CompoundLogicTransactor) Mint(opts *bind.TransactOpts, cToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CompoundLogic.contract.Transact(opts, "mint", cToken, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address cToken, uint256 amount) returns(address, uint256)
func (_CompoundLogic *CompoundLogicSession) Mint(cToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CompoundLogic.Contract.Mint(&_CompoundLogic.TransactOpts, cToken, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address cToken, uint256 amount) returns(address, uint256)
func (_CompoundLogic *CompoundLogicTransactorSession) Mint(cToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CompoundLogic.Contract.Mint(&_CompoundLogic.TransactOpts, cToken, amount)
}

// Redeem is a paid mutator transaction binding the contract method 0x7520f7ed.
//
// Solidity: function redeem(address cToken, uint256 amount, bool isredeemUnderlying, address exitToMarkets) returns(address, uint256)
func (_CompoundLogic *CompoundLogicTransactor) Redeem(opts *bind.TransactOpts, cToken common.Address, amount *big.Int, isredeemUnderlying bool, exitToMarkets common.Address) (*types.Transaction, error) {
	return _CompoundLogic.contract.Transact(opts, "redeem", cToken, amount, isredeemUnderlying, exitToMarkets)
}

// Redeem is a paid mutator transaction binding the contract method 0x7520f7ed.
//
// Solidity: function redeem(address cToken, uint256 amount, bool isredeemUnderlying, address exitToMarkets) returns(address, uint256)
func (_CompoundLogic *CompoundLogicSession) Redeem(cToken common.Address, amount *big.Int, isredeemUnderlying bool, exitToMarkets common.Address) (*types.Transaction, error) {
	return _CompoundLogic.Contract.Redeem(&_CompoundLogic.TransactOpts, cToken, amount, isredeemUnderlying, exitToMarkets)
}

// Redeem is a paid mutator transaction binding the contract method 0x7520f7ed.
//
// Solidity: function redeem(address cToken, uint256 amount, bool isredeemUnderlying, address exitToMarkets) returns(address, uint256)
func (_CompoundLogic *CompoundLogicTransactorSession) Redeem(cToken common.Address, amount *big.Int, isredeemUnderlying bool, exitToMarkets common.Address) (*types.Transaction, error) {
	return _CompoundLogic.Contract.Redeem(&_CompoundLogic.TransactOpts, cToken, amount, isredeemUnderlying, exitToMarkets)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0xabdb5ea8.
//
// Solidity: function repayBorrow(address cToken, uint256 amount) returns(address, uint256)
func (_CompoundLogic *CompoundLogicTransactor) RepayBorrow(opts *bind.TransactOpts, cToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CompoundLogic.contract.Transact(opts, "repayBorrow", cToken, amount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0xabdb5ea8.
//
// Solidity: function repayBorrow(address cToken, uint256 amount) returns(address, uint256)
func (_CompoundLogic *CompoundLogicSession) RepayBorrow(cToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CompoundLogic.Contract.RepayBorrow(&_CompoundLogic.TransactOpts, cToken, amount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0xabdb5ea8.
//
// Solidity: function repayBorrow(address cToken, uint256 amount) returns(address, uint256)
func (_CompoundLogic *CompoundLogicTransactorSession) RepayBorrow(cToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CompoundLogic.Contract.RepayBorrow(&_CompoundLogic.TransactOpts, cToken, amount)
}
