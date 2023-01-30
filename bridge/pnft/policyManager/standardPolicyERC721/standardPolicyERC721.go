// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package standardPolicyERC721

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

// Fee is an auto generated low-level Go binding around an user-defined struct.
type Fee struct {
	Rate      uint16
	Recipient common.Address
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

// StandardPolicyERC721MetaData contains all meta data concerning the StandardPolicyERC721 contract.
var StandardPolicyERC721MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"makerAsk\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"takerBid\",\"type\":\"tuple\"}],\"name\":\"canMatchMakerAsk\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"enumAssetType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"makerBid\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"takerAsk\",\"type\":\"tuple\"}],\"name\":\"canMatchMakerBid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"enumAssetType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506106f5806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80630813a7661461003b578063d5ec8c771461006f575b600080fd5b6100556004803603810190610050919061046c565b6100a3565b60405161006695949392919061058f565b60405180910390f35b6100896004803603810190610084919061046c565b610270565b60405161009a95949392919061058f565b60405180910390f35b60008060008060008560200160208101906100be9190610607565b60018111156100d0576100cf610518565b5b8760200160208101906100e39190610607565b60018111156100f5576100f4610518565b5b1415801561015257508560c00160208101906101119190610692565b73ffffffffffffffffffffffffffffffffffffffff168760c001602081019061013a9190610692565b73ffffffffffffffffffffffffffffffffffffffff16145b80156101ad575085606001602081019061016c9190610692565b73ffffffffffffffffffffffffffffffffffffffff168760600160208101906101959190610692565b73ffffffffffffffffffffffffffffffffffffffff16145b80156101c0575085608001358760800135145b80156101d0575060018760a00135145b80156101e0575060018660a00135145b801561023b57508560400160208101906101fa9190610692565b73ffffffffffffffffffffffffffffffffffffffff168760400160208101906102239190610692565b73ffffffffffffffffffffffffffffffffffffffff16145b801561024e57508560e001358760e00135145b8760e00135886080013560016000945094509450945094509295509295909350565b600080600080600085602001602081019061028b9190610607565b600181111561029d5761029c610518565b5b8760200160208101906102b09190610607565b60018111156102c2576102c1610518565b5b1415801561031f57508560c00160208101906102de9190610692565b73ffffffffffffffffffffffffffffffffffffffff168760c00160208101906103079190610692565b73ffffffffffffffffffffffffffffffffffffffff16145b801561037a57508560600160208101906103399190610692565b73ffffffffffffffffffffffffffffffffffffffff168760600160208101906103629190610692565b73ffffffffffffffffffffffffffffffffffffffff16145b801561038d575085608001358760800135145b801561039d575060018760a00135145b80156103ad575060018660a00135145b801561040857508560400160208101906103c79190610692565b73ffffffffffffffffffffffffffffffffffffffff168760400160208101906103f09190610692565b73ffffffffffffffffffffffffffffffffffffffff16145b801561041b57508560e001358760e00135145b8760e00135886080013560016000945094509450945094509295509295909350565b600080fd5b600080fd5b600080fd5b60006101a0828403121561046357610462610447565b5b81905092915050565b600080604083850312156104835761048261043d565b5b600083013567ffffffffffffffff8111156104a1576104a0610442565b5b6104ad8582860161044c565b925050602083013567ffffffffffffffff8111156104ce576104cd610442565b5b6104da8582860161044c565b9150509250929050565b60008115159050919050565b6104f9816104e4565b82525050565b6000819050919050565b610512816104ff565b82525050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6002811061055857610557610518565b5b50565b600081905061056982610547565b919050565b60006105798261055b565b9050919050565b6105898161056e565b82525050565b600060a0820190506105a460008301886104f0565b6105b16020830187610509565b6105be6040830186610509565b6105cb6060830185610509565b6105d86080830184610580565b9695505050505050565b600281106105ef57600080fd5b50565b600081359050610601816105e2565b92915050565b60006020828403121561061d5761061c61043d565b5b600061062b848285016105f2565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061065f82610634565b9050919050565b61066f81610654565b811461067a57600080fd5b50565b60008135905061068c81610666565b92915050565b6000602082840312156106a8576106a761043d565b5b60006106b68482850161067d565b9150509291505056fea2646970667358221220665552b20c7e105b8036c9081bad836ae38bbcba25b9d518edd1527b72ce494364736f6c63430008110033",
}

// StandardPolicyERC721ABI is the input ABI used to generate the binding from.
// Deprecated: Use StandardPolicyERC721MetaData.ABI instead.
var StandardPolicyERC721ABI = StandardPolicyERC721MetaData.ABI

// StandardPolicyERC721Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StandardPolicyERC721MetaData.Bin instead.
var StandardPolicyERC721Bin = StandardPolicyERC721MetaData.Bin

// DeployStandardPolicyERC721 deploys a new Ethereum contract, binding an instance of StandardPolicyERC721 to it.
func DeployStandardPolicyERC721(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StandardPolicyERC721, error) {
	parsed, err := StandardPolicyERC721MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StandardPolicyERC721Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StandardPolicyERC721{StandardPolicyERC721Caller: StandardPolicyERC721Caller{contract: contract}, StandardPolicyERC721Transactor: StandardPolicyERC721Transactor{contract: contract}, StandardPolicyERC721Filterer: StandardPolicyERC721Filterer{contract: contract}}, nil
}

// StandardPolicyERC721 is an auto generated Go binding around an Ethereum contract.
type StandardPolicyERC721 struct {
	StandardPolicyERC721Caller     // Read-only binding to the contract
	StandardPolicyERC721Transactor // Write-only binding to the contract
	StandardPolicyERC721Filterer   // Log filterer for contract events
}

// StandardPolicyERC721Caller is an auto generated read-only Go binding around an Ethereum contract.
type StandardPolicyERC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardPolicyERC721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type StandardPolicyERC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardPolicyERC721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StandardPolicyERC721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardPolicyERC721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StandardPolicyERC721Session struct {
	Contract     *StandardPolicyERC721 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// StandardPolicyERC721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StandardPolicyERC721CallerSession struct {
	Contract *StandardPolicyERC721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// StandardPolicyERC721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StandardPolicyERC721TransactorSession struct {
	Contract     *StandardPolicyERC721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// StandardPolicyERC721Raw is an auto generated low-level Go binding around an Ethereum contract.
type StandardPolicyERC721Raw struct {
	Contract *StandardPolicyERC721 // Generic contract binding to access the raw methods on
}

// StandardPolicyERC721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StandardPolicyERC721CallerRaw struct {
	Contract *StandardPolicyERC721Caller // Generic read-only contract binding to access the raw methods on
}

// StandardPolicyERC721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StandardPolicyERC721TransactorRaw struct {
	Contract *StandardPolicyERC721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewStandardPolicyERC721 creates a new instance of StandardPolicyERC721, bound to a specific deployed contract.
func NewStandardPolicyERC721(address common.Address, backend bind.ContractBackend) (*StandardPolicyERC721, error) {
	contract, err := bindStandardPolicyERC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StandardPolicyERC721{StandardPolicyERC721Caller: StandardPolicyERC721Caller{contract: contract}, StandardPolicyERC721Transactor: StandardPolicyERC721Transactor{contract: contract}, StandardPolicyERC721Filterer: StandardPolicyERC721Filterer{contract: contract}}, nil
}

// NewStandardPolicyERC721Caller creates a new read-only instance of StandardPolicyERC721, bound to a specific deployed contract.
func NewStandardPolicyERC721Caller(address common.Address, caller bind.ContractCaller) (*StandardPolicyERC721Caller, error) {
	contract, err := bindStandardPolicyERC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StandardPolicyERC721Caller{contract: contract}, nil
}

// NewStandardPolicyERC721Transactor creates a new write-only instance of StandardPolicyERC721, bound to a specific deployed contract.
func NewStandardPolicyERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*StandardPolicyERC721Transactor, error) {
	contract, err := bindStandardPolicyERC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StandardPolicyERC721Transactor{contract: contract}, nil
}

// NewStandardPolicyERC721Filterer creates a new log filterer instance of StandardPolicyERC721, bound to a specific deployed contract.
func NewStandardPolicyERC721Filterer(address common.Address, filterer bind.ContractFilterer) (*StandardPolicyERC721Filterer, error) {
	contract, err := bindStandardPolicyERC721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StandardPolicyERC721Filterer{contract: contract}, nil
}

// bindStandardPolicyERC721 binds a generic wrapper to an already deployed contract.
func bindStandardPolicyERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardPolicyERC721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardPolicyERC721 *StandardPolicyERC721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StandardPolicyERC721.Contract.StandardPolicyERC721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardPolicyERC721 *StandardPolicyERC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardPolicyERC721.Contract.StandardPolicyERC721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardPolicyERC721 *StandardPolicyERC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardPolicyERC721.Contract.StandardPolicyERC721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardPolicyERC721 *StandardPolicyERC721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StandardPolicyERC721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardPolicyERC721 *StandardPolicyERC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardPolicyERC721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardPolicyERC721 *StandardPolicyERC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardPolicyERC721.Contract.contract.Transact(opts, method, params...)
}

// CanMatchMakerAsk is a free data retrieval call binding the contract method 0xd5ec8c77.
//
// Solidity: function canMatchMakerAsk((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) makerAsk, (address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) takerBid) pure returns(bool, uint256, uint256, uint256, uint8)
func (_StandardPolicyERC721 *StandardPolicyERC721Caller) CanMatchMakerAsk(opts *bind.CallOpts, makerAsk Order, takerBid Order) (bool, *big.Int, *big.Int, *big.Int, uint8, error) {
	var out []interface{}
	err := _StandardPolicyERC721.contract.Call(opts, &out, "canMatchMakerAsk", makerAsk, takerBid)

	if err != nil {
		return *new(bool), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(uint8)).(*uint8)

	return out0, out1, out2, out3, out4, err

}

// CanMatchMakerAsk is a free data retrieval call binding the contract method 0xd5ec8c77.
//
// Solidity: function canMatchMakerAsk((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) makerAsk, (address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) takerBid) pure returns(bool, uint256, uint256, uint256, uint8)
func (_StandardPolicyERC721 *StandardPolicyERC721Session) CanMatchMakerAsk(makerAsk Order, takerBid Order) (bool, *big.Int, *big.Int, *big.Int, uint8, error) {
	return _StandardPolicyERC721.Contract.CanMatchMakerAsk(&_StandardPolicyERC721.CallOpts, makerAsk, takerBid)
}

// CanMatchMakerAsk is a free data retrieval call binding the contract method 0xd5ec8c77.
//
// Solidity: function canMatchMakerAsk((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) makerAsk, (address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) takerBid) pure returns(bool, uint256, uint256, uint256, uint8)
func (_StandardPolicyERC721 *StandardPolicyERC721CallerSession) CanMatchMakerAsk(makerAsk Order, takerBid Order) (bool, *big.Int, *big.Int, *big.Int, uint8, error) {
	return _StandardPolicyERC721.Contract.CanMatchMakerAsk(&_StandardPolicyERC721.CallOpts, makerAsk, takerBid)
}

// CanMatchMakerBid is a free data retrieval call binding the contract method 0x0813a766.
//
// Solidity: function canMatchMakerBid((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) makerBid, (address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) takerAsk) pure returns(bool, uint256, uint256, uint256, uint8)
func (_StandardPolicyERC721 *StandardPolicyERC721Caller) CanMatchMakerBid(opts *bind.CallOpts, makerBid Order, takerAsk Order) (bool, *big.Int, *big.Int, *big.Int, uint8, error) {
	var out []interface{}
	err := _StandardPolicyERC721.contract.Call(opts, &out, "canMatchMakerBid", makerBid, takerAsk)

	if err != nil {
		return *new(bool), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(uint8)).(*uint8)

	return out0, out1, out2, out3, out4, err

}

// CanMatchMakerBid is a free data retrieval call binding the contract method 0x0813a766.
//
// Solidity: function canMatchMakerBid((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) makerBid, (address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) takerAsk) pure returns(bool, uint256, uint256, uint256, uint8)
func (_StandardPolicyERC721 *StandardPolicyERC721Session) CanMatchMakerBid(makerBid Order, takerAsk Order) (bool, *big.Int, *big.Int, *big.Int, uint8, error) {
	return _StandardPolicyERC721.Contract.CanMatchMakerBid(&_StandardPolicyERC721.CallOpts, makerBid, takerAsk)
}

// CanMatchMakerBid is a free data retrieval call binding the contract method 0x0813a766.
//
// Solidity: function canMatchMakerBid((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) makerBid, (address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) takerAsk) pure returns(bool, uint256, uint256, uint256, uint8)
func (_StandardPolicyERC721 *StandardPolicyERC721CallerSession) CanMatchMakerBid(makerBid Order, takerAsk Order) (bool, *big.Int, *big.Int, *big.Int, uint8, error) {
	return _StandardPolicyERC721.Contract.CanMatchMakerBid(&_StandardPolicyERC721.CallOpts, makerBid, takerAsk)
}
