// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package incognito_proxy

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IncognitoProxyCommittee is an auto generated low-level Go binding around an user-defined struct.
type IncognitoProxyCommittee struct {
	Pubkeys    []common.Address
	StartBlock *big.Int
	SwapID     *big.Int
}

// IncognitoProxyInstructionProof is an auto generated low-level Go binding around an user-defined struct.
type IncognitoProxyInstructionProof struct {
	Inst    []byte
	Path    [][32]byte
	IsLeft  []bool
	Root    [32]byte
	BlkData [32]byte
	SigIdx  []*big.Int
	SigV    []uint8
	SigR    [][32]byte
	SigS    [][32]byte
}

// IncognitoProxyMerkleProof is an auto generated low-level Go binding around an user-defined struct.
type IncognitoProxyMerkleProof struct {
	Path   [][32]byte
	IsLeft []bool
}

// AdminPausableABI is the input ABI used to generate the binding from.
const AdminPausableABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ndays\",\"type\":\"uint256\"}],\"name\":\"Extend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expire\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"extend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_successor\",\"type\":\"address\"}],\"name\":\"retire\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"successor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AdminPausableFuncSigs maps the 4-byte function signature to its string representation.
var AdminPausableFuncSigs = map[string]string{
	"f851a440": "admin()",
	"4e71d92d": "claim()",
	"79599f96": "expire()",
	"9714378c": "extend(uint256)",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"9e6371ba": "retire(address)",
	"6ff968c3": "successor()",
	"3f4ba83a": "unpause()",
}

// AdminPausableBin is the compiled bytecode used for deploying new contracts.
var AdminPausableBin = "0x608060405234801561001057600080fd5b506040516106fc3803806106fc8339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b03199092169190911790556001805460ff60a01b191690556301e1338042016002556106808061007c6000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c806379599f961161006657806379599f96146100ea5780638456cb59146101045780639714378c1461010c5780639e6371ba14610129578063f851a4401461014f57610093565b80633f4ba83a146100985780634e71d92d146100a25780635c975abb146100aa5780636ff968c3146100c6575b600080fd5b6100a0610157565b005b6100a0610239565b6100b2610320565b604080519115158252519081900360200190f35b6100ce610330565b604080516001600160a01b039092168252519081900360200190f35b6100f261033f565b60408051918252519081900360200190f35b6100a0610345565b6100a06004803603602081101561012257600080fd5b503561046a565b6100a06004803603602081101561013f57600080fd5b50356001600160a01b031661058e565b6100ce61063b565b6000546001600160a01b031633146101a2576040805162461bcd60e51b81526020600482015260096024820152683737ba1030b236b4b760b91b604482015290519081900360640190fd5b600154600160a01b900460ff166101f7576040805162461bcd60e51b81526020600482015260146024820152736e6f7420706175736564207269676874206e6f7760601b604482015290519081900360640190fd5b6001805460ff60a01b191690556040805133815290517f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa9181900360200190a1565b6002544210610279576040805162461bcd60e51b8152602060048201526007602482015266195e1c1a5c995960ca1b604482015290519081900360640190fd5b6001546001600160a01b031633146102c7576040805162461bcd60e51b815260206004820152600c60248201526b1d5b985d5d1a1bdc9a5e995960a21b604482015290519081900360640190fd5b600154600080546001600160a01b0319166001600160a01b0392831617908190556040805191909216815290517f0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc9181900360200190a1565b600154600160a01b900460ff1681565b6001546001600160a01b031681565b60025481565b6000546001600160a01b03163314610390576040805162461bcd60e51b81526020600482015260096024820152683737ba1030b236b4b760b91b604482015290519081900360640190fd5b600154600160a01b900460ff16156103e2576040805162461bcd60e51b815260206004820152601060248201526f706175736564207269676874206e6f7760801b604482015290519081900360640190fd5b6002544210610422576040805162461bcd60e51b8152602060048201526007602482015266195e1c1a5c995960ca1b604482015290519081900360640190fd5b6001805460ff60a01b1916600160a01b1790556040805133815290517f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2589181900360200190a1565b6000546001600160a01b031633146104b5576040805162461bcd60e51b81526020600482015260096024820152683737ba1030b236b4b760b91b604482015290519081900360640190fd5b60025442106104f5576040805162461bcd60e51b8152602060048201526007602482015266195e1c1a5c995960ca1b604482015290519081900360640190fd5b61016e811061054b576040805162461bcd60e51b815260206004820152601a60248201527f63616e6e6f7420657874656e6420666f7220746f6f206c6f6e67000000000000604482015290519081900360640190fd5b600280546201518083020190556040805182815290517f02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e89181900360200190a150565b6000546001600160a01b031633146105d9576040805162461bcd60e51b81526020600482015260096024820152683737ba1030b236b4b760b91b604482015290519081900360640190fd5b6002544210610619576040805162461bcd60e51b8152602060048201526007602482015266195e1c1a5c995960ca1b604482015290519081900360640190fd5b600180546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b03168156fea2646970667358221220f07a79af805ec091ff5b9b9d4b8853e8c61f0e8db1f210d302f28d54ee5b38a364736f6c63430006060033"

// DeployAdminPausable deploys a new Ethereum contract, binding an instance of AdminPausable to it.
func DeployAdminPausable(auth *bind.TransactOpts, backend bind.ContractBackend, _admin common.Address) (common.Address, *types.Transaction, *AdminPausable, error) {
	parsed, err := abi.JSON(strings.NewReader(AdminPausableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AdminPausableBin), backend, _admin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AdminPausable{AdminPausableCaller: AdminPausableCaller{contract: contract}, AdminPausableTransactor: AdminPausableTransactor{contract: contract}, AdminPausableFilterer: AdminPausableFilterer{contract: contract}}, nil
}

// AdminPausable is an auto generated Go binding around an Ethereum contract.
type AdminPausable struct {
	AdminPausableCaller     // Read-only binding to the contract
	AdminPausableTransactor // Write-only binding to the contract
	AdminPausableFilterer   // Log filterer for contract events
}

// AdminPausableCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdminPausableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminPausableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdminPausableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminPausableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AdminPausableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminPausableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdminPausableSession struct {
	Contract     *AdminPausable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdminPausableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdminPausableCallerSession struct {
	Contract *AdminPausableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AdminPausableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdminPausableTransactorSession struct {
	Contract     *AdminPausableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AdminPausableRaw is an auto generated low-level Go binding around an Ethereum contract.
type AdminPausableRaw struct {
	Contract *AdminPausable // Generic contract binding to access the raw methods on
}

// AdminPausableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AdminPausableCallerRaw struct {
	Contract *AdminPausableCaller // Generic read-only contract binding to access the raw methods on
}

// AdminPausableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AdminPausableTransactorRaw struct {
	Contract *AdminPausableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAdminPausable creates a new instance of AdminPausable, bound to a specific deployed contract.
func NewAdminPausable(address common.Address, backend bind.ContractBackend) (*AdminPausable, error) {
	contract, err := bindAdminPausable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AdminPausable{AdminPausableCaller: AdminPausableCaller{contract: contract}, AdminPausableTransactor: AdminPausableTransactor{contract: contract}, AdminPausableFilterer: AdminPausableFilterer{contract: contract}}, nil
}

// NewAdminPausableCaller creates a new read-only instance of AdminPausable, bound to a specific deployed contract.
func NewAdminPausableCaller(address common.Address, caller bind.ContractCaller) (*AdminPausableCaller, error) {
	contract, err := bindAdminPausable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdminPausableCaller{contract: contract}, nil
}

// NewAdminPausableTransactor creates a new write-only instance of AdminPausable, bound to a specific deployed contract.
func NewAdminPausableTransactor(address common.Address, transactor bind.ContractTransactor) (*AdminPausableTransactor, error) {
	contract, err := bindAdminPausable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdminPausableTransactor{contract: contract}, nil
}

// NewAdminPausableFilterer creates a new log filterer instance of AdminPausable, bound to a specific deployed contract.
func NewAdminPausableFilterer(address common.Address, filterer bind.ContractFilterer) (*AdminPausableFilterer, error) {
	contract, err := bindAdminPausable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdminPausableFilterer{contract: contract}, nil
}

// bindAdminPausable binds a generic wrapper to an already deployed contract.
func bindAdminPausable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AdminPausableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdminPausable *AdminPausableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AdminPausable.Contract.AdminPausableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdminPausable *AdminPausableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminPausable.Contract.AdminPausableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdminPausable *AdminPausableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdminPausable.Contract.AdminPausableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdminPausable *AdminPausableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AdminPausable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdminPausable *AdminPausableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminPausable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdminPausable *AdminPausableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdminPausable.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_AdminPausable *AdminPausableCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AdminPausable.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_AdminPausable *AdminPausableSession) Admin() (common.Address, error) {
	return _AdminPausable.Contract.Admin(&_AdminPausable.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_AdminPausable *AdminPausableCallerSession) Admin() (common.Address, error) {
	return _AdminPausable.Contract.Admin(&_AdminPausable.CallOpts)
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() view returns(uint256)
func (_AdminPausable *AdminPausableCaller) Expire(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AdminPausable.contract.Call(opts, out, "expire")
	return *ret0, err
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() view returns(uint256)
func (_AdminPausable *AdminPausableSession) Expire() (*big.Int, error) {
	return _AdminPausable.Contract.Expire(&_AdminPausable.CallOpts)
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() view returns(uint256)
func (_AdminPausable *AdminPausableCallerSession) Expire() (*big.Int, error) {
	return _AdminPausable.Contract.Expire(&_AdminPausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AdminPausable *AdminPausableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AdminPausable.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AdminPausable *AdminPausableSession) Paused() (bool, error) {
	return _AdminPausable.Contract.Paused(&_AdminPausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AdminPausable *AdminPausableCallerSession) Paused() (bool, error) {
	return _AdminPausable.Contract.Paused(&_AdminPausable.CallOpts)
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() view returns(address)
func (_AdminPausable *AdminPausableCaller) Successor(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AdminPausable.contract.Call(opts, out, "successor")
	return *ret0, err
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() view returns(address)
func (_AdminPausable *AdminPausableSession) Successor() (common.Address, error) {
	return _AdminPausable.Contract.Successor(&_AdminPausable.CallOpts)
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() view returns(address)
func (_AdminPausable *AdminPausableCallerSession) Successor() (common.Address, error) {
	return _AdminPausable.Contract.Successor(&_AdminPausable.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_AdminPausable *AdminPausableTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminPausable.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_AdminPausable *AdminPausableSession) Claim() (*types.Transaction, error) {
	return _AdminPausable.Contract.Claim(&_AdminPausable.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_AdminPausable *AdminPausableTransactorSession) Claim() (*types.Transaction, error) {
	return _AdminPausable.Contract.Claim(&_AdminPausable.TransactOpts)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_AdminPausable *AdminPausableTransactor) Extend(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _AdminPausable.contract.Transact(opts, "extend", n)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_AdminPausable *AdminPausableSession) Extend(n *big.Int) (*types.Transaction, error) {
	return _AdminPausable.Contract.Extend(&_AdminPausable.TransactOpts, n)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_AdminPausable *AdminPausableTransactorSession) Extend(n *big.Int) (*types.Transaction, error) {
	return _AdminPausable.Contract.Extend(&_AdminPausable.TransactOpts, n)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AdminPausable *AdminPausableTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminPausable.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AdminPausable *AdminPausableSession) Pause() (*types.Transaction, error) {
	return _AdminPausable.Contract.Pause(&_AdminPausable.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AdminPausable *AdminPausableTransactorSession) Pause() (*types.Transaction, error) {
	return _AdminPausable.Contract.Pause(&_AdminPausable.TransactOpts)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_AdminPausable *AdminPausableTransactor) Retire(opts *bind.TransactOpts, _successor common.Address) (*types.Transaction, error) {
	return _AdminPausable.contract.Transact(opts, "retire", _successor)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_AdminPausable *AdminPausableSession) Retire(_successor common.Address) (*types.Transaction, error) {
	return _AdminPausable.Contract.Retire(&_AdminPausable.TransactOpts, _successor)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_AdminPausable *AdminPausableTransactorSession) Retire(_successor common.Address) (*types.Transaction, error) {
	return _AdminPausable.Contract.Retire(&_AdminPausable.TransactOpts, _successor)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AdminPausable *AdminPausableTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminPausable.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AdminPausable *AdminPausableSession) Unpause() (*types.Transaction, error) {
	return _AdminPausable.Contract.Unpause(&_AdminPausable.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AdminPausable *AdminPausableTransactorSession) Unpause() (*types.Transaction, error) {
	return _AdminPausable.Contract.Unpause(&_AdminPausable.TransactOpts)
}

// AdminPausableClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the AdminPausable contract.
type AdminPausableClaimIterator struct {
	Event *AdminPausableClaim // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AdminPausableClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminPausableClaim)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AdminPausableClaim)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AdminPausableClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminPausableClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminPausableClaim represents a Claim event raised by the AdminPausable contract.
type AdminPausableClaim struct {
	Claimer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_AdminPausable *AdminPausableFilterer) FilterClaim(opts *bind.FilterOpts) (*AdminPausableClaimIterator, error) {

	logs, sub, err := _AdminPausable.contract.FilterLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return &AdminPausableClaimIterator{contract: _AdminPausable.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_AdminPausable *AdminPausableFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *AdminPausableClaim) (event.Subscription, error) {

	logs, sub, err := _AdminPausable.contract.WatchLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminPausableClaim)
				if err := _AdminPausable.contract.UnpackLog(event, "Claim", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaim is a log parse operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_AdminPausable *AdminPausableFilterer) ParseClaim(log types.Log) (*AdminPausableClaim, error) {
	event := new(AdminPausableClaim)
	if err := _AdminPausable.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AdminPausableExtendIterator is returned from FilterExtend and is used to iterate over the raw logs and unpacked data for Extend events raised by the AdminPausable contract.
type AdminPausableExtendIterator struct {
	Event *AdminPausableExtend // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AdminPausableExtendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminPausableExtend)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AdminPausableExtend)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AdminPausableExtendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminPausableExtendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminPausableExtend represents a Extend event raised by the AdminPausable contract.
type AdminPausableExtend struct {
	Ndays *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterExtend is a free log retrieval operation binding the contract event 0x02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8.
//
// Solidity: event Extend(uint256 ndays)
func (_AdminPausable *AdminPausableFilterer) FilterExtend(opts *bind.FilterOpts) (*AdminPausableExtendIterator, error) {

	logs, sub, err := _AdminPausable.contract.FilterLogs(opts, "Extend")
	if err != nil {
		return nil, err
	}
	return &AdminPausableExtendIterator{contract: _AdminPausable.contract, event: "Extend", logs: logs, sub: sub}, nil
}

// WatchExtend is a free log subscription operation binding the contract event 0x02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8.
//
// Solidity: event Extend(uint256 ndays)
func (_AdminPausable *AdminPausableFilterer) WatchExtend(opts *bind.WatchOpts, sink chan<- *AdminPausableExtend) (event.Subscription, error) {

	logs, sub, err := _AdminPausable.contract.WatchLogs(opts, "Extend")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminPausableExtend)
				if err := _AdminPausable.contract.UnpackLog(event, "Extend", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExtend is a log parse operation binding the contract event 0x02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8.
//
// Solidity: event Extend(uint256 ndays)
func (_AdminPausable *AdminPausableFilterer) ParseExtend(log types.Log) (*AdminPausableExtend, error) {
	event := new(AdminPausableExtend)
	if err := _AdminPausable.contract.UnpackLog(event, "Extend", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AdminPausablePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the AdminPausable contract.
type AdminPausablePausedIterator struct {
	Event *AdminPausablePaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AdminPausablePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminPausablePaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AdminPausablePaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AdminPausablePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminPausablePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminPausablePaused represents a Paused event raised by the AdminPausable contract.
type AdminPausablePaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address pauser)
func (_AdminPausable *AdminPausableFilterer) FilterPaused(opts *bind.FilterOpts) (*AdminPausablePausedIterator, error) {

	logs, sub, err := _AdminPausable.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AdminPausablePausedIterator{contract: _AdminPausable.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address pauser)
func (_AdminPausable *AdminPausableFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AdminPausablePaused) (event.Subscription, error) {

	logs, sub, err := _AdminPausable.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminPausablePaused)
				if err := _AdminPausable.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address pauser)
func (_AdminPausable *AdminPausableFilterer) ParsePaused(log types.Log) (*AdminPausablePaused, error) {
	event := new(AdminPausablePaused)
	if err := _AdminPausable.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AdminPausableUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the AdminPausable contract.
type AdminPausableUnpausedIterator struct {
	Event *AdminPausableUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AdminPausableUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminPausableUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AdminPausableUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AdminPausableUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminPausableUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminPausableUnpaused represents a Unpaused event raised by the AdminPausable contract.
type AdminPausableUnpaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address pauser)
func (_AdminPausable *AdminPausableFilterer) FilterUnpaused(opts *bind.FilterOpts) (*AdminPausableUnpausedIterator, error) {

	logs, sub, err := _AdminPausable.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AdminPausableUnpausedIterator{contract: _AdminPausable.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address pauser)
func (_AdminPausable *AdminPausableFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AdminPausableUnpaused) (event.Subscription, error) {

	logs, sub, err := _AdminPausable.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminPausableUnpaused)
				if err := _AdminPausable.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address pauser)
func (_AdminPausable *AdminPausableFilterer) ParseUnpaused(log types.Log) (*AdminPausableUnpaused, error) {
	event := new(AdminPausableUnpaused)
	if err := _AdminPausable.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncognitoProxyABI is the input ABI used to generate the binding from.
const IncognitoProxyABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"beaconCommittee\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"bridgeCommittee\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startHeight\",\"type\":\"uint256\"}],\"name\":\"BeaconCommitteeSwapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isBeacon\",\"type\":\"bool\"}],\"name\":\"CandidatePromoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isBeacon\",\"type\":\"bool\"}],\"name\":\"ChainFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ndays\",\"type\":\"uint256\"}],\"name\":\"Extend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startHeight\",\"type\":\"uint256\"}],\"name\":\"SubmittedBridgeCandidate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"beaconCandidates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"beaconBlockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"beaconCommittees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapID\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beaconFinality\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bridgeCandidates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"beaconBlockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bridgeCommittees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapID\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeFinality\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expire\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"extend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"numVals\",\"type\":\"uint256\"}],\"name\":\"extractCommitteeFromInstruction\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"}],\"name\":\"extractDataFromBlockMerkleInstruction\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"}],\"name\":\"extractMetaFromInstruction\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"sigIdx\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"}],\"name\":\"filterSigners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blkHeight\",\"type\":\"uint256\"}],\"name\":\"findBeaconCommitteeFromHeight\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blkHeight\",\"type\":\"uint256\"}],\"name\":\"findBridgeCommitteeFromHeight\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getBeaconCommittee\",\"outputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"pubkeys\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapID\",\"type\":\"uint256\"}],\"internalType\":\"structIncognitoProxy.Committee\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getBridgeCommittee\",\"outputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"pubkeys\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapID\",\"type\":\"uint256\"}],\"internalType\":\"structIncognitoProxy.Committee\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isBeacon\",\"type\":\"bool\"}],\"name\":\"getLatestSwapID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isBeacon\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"instHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blkHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"instPath\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"instPathIsLeft\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"instRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"sigIdx\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"sigV\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigR\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigS\",\"type\":\"bytes32[]\"}],\"name\":\"instructionApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"instHash\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"instPath\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"instPathIsLeft\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"instRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint8[]\",\"name\":\"sigV\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigR\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigS\",\"type\":\"bytes32[]\"}],\"name\":\"instructionApprovedBySigners\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"left\",\"type\":\"bool[]\"}],\"name\":\"leafInMerkleTree\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"swapID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isBeacon\",\"type\":\"bool\"}],\"name\":\"loadCandidates\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"swapID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isBeacon\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"isLeft\",\"type\":\"bool[]\"}],\"internalType\":\"structIncognitoProxy.MerkleProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"name\":\"promoteCandidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_successor\",\"type\":\"address\"}],\"name\":\"retire\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"isLeft\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"sigIdx\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"sigV\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigR\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigS\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIncognitoProxy.InstructionProof\",\"name\":\"instProof\",\"type\":\"tuple\"}],\"name\":\"submitBeaconCandidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"isLeft\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"sigIdx\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"sigV\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigR\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigS\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIncognitoProxy.InstructionProof[2]\",\"name\":\"instProofs\",\"type\":\"tuple[2]\"}],\"name\":\"submitBridgeCandidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"isLeft\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"sigIdx\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"sigV\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigR\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigS\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIncognitoProxy.InstructionProof[2]\",\"name\":\"instProofs\",\"type\":\"tuple[2]\"},{\"internalType\":\"uint256\",\"name\":\"swapID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isBeacon\",\"type\":\"bool\"}],\"name\":\"submitFinalityProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"successor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"committee\",\"type\":\"address[]\"},{\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"verifySig\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// IncognitoProxyFuncSigs maps the 4-byte function signature to its string representation.
var IncognitoProxyFuncSigs = map[string]string{
	"f851a440": "admin()",
	"62532738": "beaconCandidates(uint256)",
	"f203a5ed": "beaconCommittees(uint256)",
	"5d7f86d9": "beaconFinality()",
	"e83d31ec": "bridgeCandidates(uint256)",
	"9b30b637": "bridgeCommittees(uint256)",
	"2c1c9cac": "bridgeFinality()",
	"4e71d92d": "claim()",
	"79599f96": "expire()",
	"9714378c": "extend(uint256)",
	"8eb60066": "extractCommitteeFromInstruction(bytes,uint256)",
	"f179dc78": "extractDataFromBlockMerkleInstruction(bytes)",
	"90500bae": "extractMetaFromInstruction(bytes)",
	"6ae474d8": "filterSigners(uint256[],address[])",
	"b600ffdb": "findBeaconCommitteeFromHeight(uint256)",
	"f5205fde": "findBridgeCommitteeFromHeight(uint256)",
	"faea3167": "getBeaconCommittee(uint256)",
	"8ceb69c3": "getBridgeCommittee(uint256)",
	"722a5d7b": "getLatestSwapID(bool)",
	"f65d2116": "instructionApproved(bool,bytes32,uint256,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])",
	"3e379d05": "instructionApprovedBySigners(bytes32,address[],bytes32[],bool[],bytes32,bytes32,uint8[],bytes32[],bytes32[])",
	"1217a3c7": "leafInMerkleTree(bytes32,bytes32,bytes32[],bool[])",
	"6e845d19": "loadCandidates(uint256,bool)",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"ea200467": "promoteCandidate(uint256,bool,(bytes32[],bool[]))",
	"9e6371ba": "retire(address)",
	"52d45239": "submitBeaconCandidate(bytes,(bytes,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[]))",
	"7e59f5bf": "submitBridgeCandidate(bytes,(bytes,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])[2])",
	"c37cfbf3": "submitFinalityProof((bytes,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])[2],uint256,bool)",
	"6ff968c3": "successor()",
	"3f4ba83a": "unpause()",
	"3aacfdad": "verifySig(address[],bytes32,uint8[],bytes32[],bytes32[])",
}

// IncognitoProxyBin is the compiled bytecode used for deploying new contracts.
var IncognitoProxyBin = "0x60806040523480156200001157600080fd5b50604051620030b9380380620030b98339810160408190526200003491620002bf565b600080546001600160a01b0319166001600160a01b0385161781556001805460ff60a01b19168155426301e1338001600255604080516060810182528581526020808201859052918101849052600380549384018155938490528051805191947fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b94029390930192620000cc9284929101906200016b565b50602082810151600180840191909155604093840151600290930192909255825160608101845284815260008183018190529381018490526004805493840181559093528251805160039093027f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b01926200014b92849201906200016b565b50602082015181600101556040820151816002015550505050506200033a565b828054828255906000526020600020908101928215620001c3579160200282015b82811115620001c357825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906200018c565b50620001d1929150620001d5565b5090565b620001fc91905b80821115620001d15780546001600160a01b0319168155600101620001dc565b90565b80516001600160a01b03811681146200021757600080fd5b92915050565b600082601f8301126200022e578081fd5b81516001600160401b038082111562000245578283fd5b60208083026040518282820101818110858211171562000263578687fd5b6040528481529450818501925085820181870183018810156200028557600080fd5b600091505b84821015620002b4576200029f8882620001ff565b8452928201926001919091019082016200028a565b505050505092915050565b600080600060608486031215620002d4578283fd5b620002e08585620001ff565b60208501519093506001600160401b0380821115620002fd578384fd5b6200030b878388016200021d565b9350604086015191508082111562000321578283fd5b5062000330868287016200021d565b9150509250925092565b612d6f806200034a6000396000f3fe608060405234801561001057600080fd5b50600436106101fb5760003560e01c80638456cb591161011a578063c37cfbf3116100ad578063f203a5ed1161007c578063f203a5ed14610444578063f5205fde14610457578063f65d21161461046a578063f851a4401461047d578063faea316714610485576101fb565b8063c37cfbf3146103e9578063e83d31ec146103fc578063ea2004671461040f578063f179dc7814610422576101fb565b80639714378c116100e95780639714378c1461038f5780639b30b637146103a25780639e6371ba146103b5578063b600ffdb146103c8576101fb565b80638456cb59146103305780638ceb69c3146103385780638eb600661461035857806390500bae1461036b576101fb565b80635d7f86d9116101925780636ff968c3116101615780636ff968c3146102e0578063722a5d7b146102f557806379599f96146103155780637e59f5bf1461031d576101fb565b80635d7f86d914610292578063625327381461029a5780636ae474d8146102ad5780636e845d19146102cd576101fb565b80633f4ba83a116101ce5780633f4ba83a146102655780634e71d92d1461026f57806352d45239146102775780635c975abb1461028a576101fb565b80631217a3c7146102005780632c1c9cac146102295780633aacfdad1461023f5780633e379d0514610252575b600080fd5b61021361020e36600461264f565b610498565b6040516102209190612980565b60405180910390f35b610231610595565b604051610220929190612929565b61021361024d36600461228a565b61059e565b610213610260366004612543565b61069c565b61026d61074e565b005b61026d6107f1565b61026d610285366004612752565b61088d565b610213610bb1565b610231610bc1565b6102316102a83660046127ea565b610bca565b6102c06102bb366004612393565b610be6565b604051610220919061294b565b6102c06102db366004612802565b610d0b565b6102e8610e44565b6040516102209190612937565b6103086103033660046123f3565b610e53565b6040516102209190612920565b610308610eb2565b61026d61032b3660046126fc565b610eb8565b61026d611226565b61034b6103463660046127ea565b6112de565b6040516102209190612c52565b6102c06103663660046127a8565b611389565b61037e6103793660046126c2565b611432565b604051610220959493929190612cb8565b61026d61039d3660046127ea565b6114a4565b6102316103b03660046127ea565b611558565b61026d6103c3366004612268565b611589565b6103db6103d63660046127ea565b6115f6565b60405161022092919061295e565b61026d6103f736600461233b565b6116e4565b61023161040a3660046127ea565b61194e565b61026d61041d366004612831565b61196a565b6104356104303660046126c2565b611c69565b60405161022093929190612c9d565b6102316104523660046127ea565b611cac565b6103db6104653660046127ea565b611cb9565b61021361047836600461240f565b611d2e565b6102e8611d8b565b61034b6104933660046127ea565b611d9a565b600082518251146104a857600080fd5b8460005b8451811015610589578381815181106104c157fe5b602002602001015115610511578481815181106104da57fe5b6020026020010151826040516020016104f4929190612929565b604051602081830303815290604052805190602001209150610581565b84818151811061051d57fe5b60200260200101516000801b14156105425781826040516020016104f4929190612929565b8185828151811061054f57fe5b6020026020010151604051602001610568929190612929565b6040516020818303038152906040528051906020012091505b6001016104ac565b50909314949350505050565b600954600a5482565b600082518451146105ae57600080fd5b81518451146105bc57600080fd5b60005b845181101561068d578681815181106105d457fe5b60200260200101516001600160a01b03166001878784815181106105f457fe5b602002602001015187858151811061060857fe5b602002602001015187868151811061061c57fe5b602002602001015160405160008152602001604052604051610641949392919061298b565b6020604051602081039080840390855afa158015610663573d6000803e3d6000fd5b505050602060405103516001600160a01b031614610685576000915050610693565b6001016105bf565b50600190505b95945050505050565b60008085876040516020016106b2929190612929565b604051602081830303815290604052805190602001206040516020016106d89190612920565b60405160208183030381529060405280519060200120905060038a51600202816106fe57fe5b04855111610710576000915050610741565b61071d8a8287878761059e565b61072657600080fd5b6107328b888b8b610498565b61073b57600080fd5b60019150505b9998505050505050505050565b6000546001600160a01b031633146107815760405162461bcd60e51b815260040161077890612c02565b60405180910390fd5b600154600160a01b900460ff166107aa5760405162461bcd60e51b8152600401610778906129a9565b6001805460ff60a01b191690556040517f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa906107e7903390612937565b60405180910390a1565b60025442106108125760405162461bcd60e51b815260040161077890612a87565b6001546001600160a01b0316331461083c5760405162461bcd60e51b815260040161077890612b05565b600154600080546001600160a01b0319166001600160a01b0392831617908190556040517f0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc926107e7921690612937565b600154600160a01b900460ff16156108b75760405162461bcd60e51b815260040161077890612ba9565b6108bf611daf565b6108c883611432565b60808601526060850152604084015260ff90811660208401521680825260461480156108fb5750806020015160ff166001145b61090457600080fd5b8251602084012060038054606091600091600019810190811061092357fe5b90600052602060002090600302016002015490508084608001511161095a5760405162461bcd60e51b8152600401610778906129ff565b80600101846080015114156109f5576109ee8560a0015160036001600380549050038154811061098657fe5b600091825260209182902060039091020180546040805182850281018501909152818152928301828280156109e457602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116109c6575b5050505050610be6565b9150610a70565b60a08501516080850151600019016000908152600660209081526040918290208054835181840281018401909452808452610a6d9493928301828280156109e4576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116109c6575050505050610be6565b91505b610a9e83838760200151886040015189606001518a608001518b60c001518c60e001518d610100015161069c565b610aa757600080fd5b6060610ab7878660600151611389565b9050600086608001518760600151604051602001610ad6929190612929565b60405160208183030381529060405280519060200120604051602001610afc9190612920565b60408051601f198184030181528282528051602091820120606084018352858452898301518483015283830181905260808a015160009081526006835292909220835180519395509092610b539284920190611de4565b50602082015160018201556040918201516002909101556003548782015191517fc918d4d247e9af7d4a8ef8ac6f0078262006c87209d062a6d846f5b547dcff6692610b9f9291612929565b60405180910390a15050505050505050565b600154600160a01b900460ff1681565b60075460085482565b6006602052600090815260409020600181015460029091015482565b60608082516001600160401b0381118015610c0057600080fd5b50604051908082528060200260200182016040528015610c2a578160200160208202803683370190505b50905060005b8451811015610d0157600081118015610c725750846001820381518110610c5357fe5b6020026020010151858281518110610c6757fe5b602002602001015111155b80610c9157508351858281518110610c8657fe5b602002602001015110155b15610cae5760405162461bcd60e51b8152600401610778906129d7565b83858281518110610cbb57fe5b602002602001015181518110610ccd57fe5b6020026020010151828281518110610ce157fe5b6001600160a01b0390921660209283029190910190910152600101610c30565b5090505b92915050565b60608115610dab57600083815260066020526040902060010154610d415760405162461bcd60e51b815260040161077890612bd3565b60008381526006602090815260409182902080548351818402810184019094528084529091830182828015610d9f57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610d81575b50505050509050610d05565b600083815260056020526040902060010154610dd95760405162461bcd60e51b815260040161077890612b7a565b60008381526005602090815260409182902080548351818402810184019094528084529091830182828015610e3757602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610e19575b5050505050905092915050565b6001546001600160a01b031681565b60008115610e8657600380546000198101908110610e6d57fe5b9060005260206000209060030201600201549050610ead565b600480546000198101908110610e9857fe5b90600052602060002090600302016002015490505b919050565b60025481565b600154600160a01b900460ff1615610ee25760405162461bcd60e51b815260040161077890612ba9565b610eea611daf565b610ef383611432565b60808601526060850152604084015260ff9081166020840152168082526047148015610f265750806020015160ff166001145b610f2f57600080fd5b8251602084012060038054610fbf916001918491906000198101908110610f5257fe5b90600052602060002090600302016001015486600060028110610f7157fe5b60200201516020015187600060028110610f8757fe5b60200201516040015188516060810151608082015160a083015160c084015160e0850151610100909501519394929391929091611d2e565b610fc857600080fd5b600480546060916000916000198101908110610fe057fe5b9060005260206000209060030201600201549050808460800151116110175760405162461bcd60e51b8152600401610778906129ff565b806001018460800151141561104e57611047856001602002015160a0015160048054600019810190811061098657fe5b91506110c9565b60208086015160a0015160808601516000190160009081526005835260409081902080548251818602810186019093528083526110c6948301828280156109e4576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116109c6575050505050610be6565b91505b6020858101519081015160408201516060830151608084015160c085015160e09095015161111195899589959094909390929091908d60015b6020020151610100015161069c565b61111a57600080fd5b611128868560600151611389565b8551608081015160609091015160405192945060009261114c929190602001612929565b604051602081830303815290604052805190602001206040516020016111729190612920565b60408051601f19818403018152828252805160209182012060608401835286845288830151848301528383018190526080890151600090815260058352929092208351805193955090926111c99284920190611de4565b50602082015160018201556040918201516002909101556004548682015191517fc918d4d247e9af7d4a8ef8ac6f0078262006c87209d062a6d846f5b547dcff66926112159291612929565b60405180910390a150505050505050565b6000546001600160a01b031633146112505760405162461bcd60e51b815260040161077890612c02565b600154600160a01b900460ff161561127a5760405162461bcd60e51b815260040161077890612ba9565b600254421061129b5760405162461bcd60e51b815260040161077890612a87565b6001805460ff60a01b1916600160a01b1790556040517f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258906107e7903390612937565b6112e6611e49565b600482815481106112f357fe5b90600052602060002090600302016040518060600160405290816000820180548060200260200160405190810160405280929190818152602001828054801561136557602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611347575b50505050508152602001600182015481526020016002820154815250509050919050565b60608160200260620183511461139e57600080fd5b6060826001600160401b03811180156113b657600080fd5b506040519080825280602002602001820160405280156113e0578160200160208202803683370190505b5090506000805b84811015611428576020810260828701015191508183828151811061140857fe5b6001600160a01b03909216602092830291909101909101526001016113e7565b5090949350505050565b600080600080600060628651101561144957600080fd5b60008660008151811061145857fe5b602001015160f81c60f81b60f81c905060008760018151811061147757fe5b0160200151602289015160428a01516062909a0151939a60f89290921c9990985096509194509092505050565b6000546001600160a01b031633146114ce5760405162461bcd60e51b815260040161077890612c02565b60025442106114ef5760405162461bcd60e51b815260040161077890612a87565b61016e81106115105760405162461bcd60e51b815260040161077890612ace565b600280546201518083020190556040517f02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e89061154d908390612920565b60405180910390a150565b6004818154811061156557fe5b90600052602060002090600302016000915090508060010154908060020154905082565b6000546001600160a01b031633146115b35760405162461bcd60e51b815260040161077890612c02565b60025442106115d45760405162461bcd60e51b815260040161077890612a87565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b60035460609060009081908061160b57600080fd5b600019015b80821461165e5760006002600184840101049050856003828154811061163257fe5b9060005260206000209060030201600101541161165157809250611658565b6001810391505b50611610565b6003828154811061166b57fe5b906000526020600020906003020160000182818054806020026020016040519081016040528092919081815260200182805480156116d257602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116116b4575b50505050509150935093505050915091565b600154600160a01b900460ff161561170e5760405162461bcd60e51b815260040161077890612ba9565b61171781610e53565b8210156117365760405162461bcd60e51b815260040161077890612b2b565b60606117428383610d0b565b905060005b600281101561182e57606061177086836002811061176157fe5b602002015160a0015184610be6565b905061181c86836002811061178157fe5b602002015160000151805190602001208288856002811061179e57fe5b6020020151602001518986600281106117b357fe5b6020020151604001518a87600281106117c857fe5b6020020151606001518b88600281106117dd57fe5b6020020151608001518c89600281106117f257fe5b602002015160c001518d8a6002811061180757fe5b602002015160e001518e8b6002811061110257fe5b61182557600080fd5b50600101611747565b506000808061184487825b602002015151611c69565b919450925090506000808061185a8a6001611839565b925092509250600a818161186a57fe5b04600a85046001011461188f5760405162461bcd60e51b815260040161077890612c25565b8560ff1660491480156118a557508260ff166049145b6118c15760405162461bcd60e51b815260040161077890612aa8565b87156118eb576040805180820190915260008082526020909101869052600755600885905561190b565b6040805180820190915260008082526020909101869052600955600a8590555b7f045c41374537196136848ce62859f8aa22ae91b490103bd2bed75eac6a4d51808860405161193a9190612980565b60405180910390a150505050505050505050565b6005602052600090815260409020600181015460029091015482565b600154600160a01b900460ff16156119945760405162461bcd60e51b815260040161077890612ba9565b60085460009083156119b95760008581526006602052604090206002015491506119ce565b60008581526005602052604090206002015491505b6119e2828285600001518660200151610498565b6119eb57600080fd5b8315611b0f57600380548691906000198101908110611a0657fe5b90600052602060002090600302016002015460010114611a385760405162461bcd60e51b815260040161077890612a44565b60408051600087815260066020908152908390208054608092810284018301909452606083018481526003948493919291840182828015611aa257602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611a84575b505050918352505060008881526006602090815260408083206001908101548386015293018a905284549283018555938152839020825180519394600390930290910192611af39284920190611de4565b5060208201518160010155604082015181600201555050611c29565b600480548691906000198101908110611b2457fe5b90600052602060002090600302016002015460010114611b565760405162461bcd60e51b815260040161077890612a44565b60408051600087815260056020908152908390208054608092810284018301909452606083018481526004948493919291840182828015611bc057602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611ba2575b505050918352505060008881526005602090815260408083206001908101548386015293018a905284549283018555938152839020825180519394600390930290910192611c119284920190611de4565b50602082015181600101556040820151816002015550505b7f40906ca8849ddb84e9284fa5813086f804b1a6e4b236420725c43cbf663b61d58585604051611c5a929190612c8d565b60405180910390a15050505050565b6000806000604184511015611c7d57600080fd5b600084600081518110611c8c57fe5b0160200151602186015160419096015160f89190911c9690945092505050565b6003818154811061156557fe5b600454606090600090819080611cce57600080fd5b600019015b808214611d2157600060026001848401010490508560048281548110611cf557fe5b90600052602060002090600302016001015411611d1457809250611d1b565b6001810391505b50611cd3565b6004828154811061166b57fe5b6000606060008d15611d4d57611d438c6115f6565b9092509050611d5c565b611d568c611cb9565b90925090505b611d668783610be6565b9150611d798d838d8d8d8d8c8c8c61069c565b9e9d5050505050505050505050505050565b6000546001600160a01b031681565b611da2611e49565b600382815481106112f357fe5b6040518060a00160405280600060ff168152602001600060ff1681526020016000815260200160008152602001600081525090565b828054828255906000526020600020908101928215611e39579160200282015b82811115611e3957825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190611e04565b50611e45929150611e6a565b5090565b60405180606001604052806060815260200160008152602001600081525090565b611e8e91905b80821115611e455780546001600160a01b0319168155600101611e70565b90565b80356001600160a01b0381168114610d0557600080fd5b600082601f830112611eb8578081fd5b8135611ecb611ec682612d09565b612ce3565b818152915060208083019084810181840286018201871015611eec57600080fd5b60005b84811015611f1357611f018883611e91565b84529282019290820190600101611eef565b505050505092915050565b600082601f830112611f2e578081fd5b8135611f3c611ec682612d09565b818152915060208083019084810181840286018201871015611f5d57600080fd5b60005b84811015611f1357611f7288836120b0565b84529282019290820190600101611f60565b600082601f830112611f94578081fd5b8135611fa2611ec682612d09565b818152915060208083019084810181840286018201871015611fc357600080fd5b60005b84811015611f1357813584529282019290820190600101611fc6565b600082601f830112611ff2578081fd5b611ffc6040612ce3565b9050808260005b600281101561202e576120198683358701612128565b83526020928301929190910190600101612003565b50505092915050565b600082601f830112612047578081fd5b8135612055611ec682612d09565b81815291506020808301908481018184028601820187101561207657600080fd5b6000805b858110156120a457823560ff81168114612092578283fd5b8552938301939183019160010161207a565b50505050505092915050565b80358015158114610d0557600080fd5b600082601f8301126120d0578081fd5b81356001600160401b038111156120e5578182fd5b6120f8601f8201601f1916602001612ce3565b915080825283602082850101111561210f57600080fd5b8060208401602084013760009082016020015292915050565b600061012080838503121561213b578182fd5b61214481612ce3565b91505081356001600160401b038082111561215e57600080fd5b61216a858386016120c0565b8352602084013591508082111561218057600080fd5b61218c85838601611f84565b602084015260408401359150808211156121a557600080fd5b6121b185838601611f1e565b6040840152606084013560608401526080840135608084015260a08401359150808211156121de57600080fd5b6121ea85838601611f84565b60a084015260c084013591508082111561220357600080fd5b61220f85838601612037565b60c084015260e084013591508082111561222857600080fd5b61223485838601611f84565b60e08401526101009150818401358181111561224f57600080fd5b61225b86828701611f84565b8385015250505092915050565b600060208284031215612279578081fd5b6122838383611e91565b9392505050565b600080600080600060a086880312156122a1578081fd5b85356001600160401b03808211156122b7578283fd5b6122c389838a01611ea8565b96506020880135955060408801359150808211156122df578283fd5b6122eb89838a01612037565b94506060880135915080821115612300578283fd5b61230c89838a01611f84565b93506080880135915080821115612321578283fd5b5061232e88828901611f84565b9150509295509295909350565b60008060006060848603121561234f578283fd5b83356001600160401b03811115612364578384fd5b61237086828701611fe2565b93505060208401359150604084013561238881612d28565b809150509250925092565b600080604083850312156123a5578182fd5b82356001600160401b03808211156123bb578384fd5b6123c786838701611f84565b935060208501359150808211156123dc578283fd5b506123e985828601611ea8565b9150509250929050565b600060208284031215612404578081fd5b813561228381612d28565b60008060008060008060008060008060006101608c8e031215612430578889fd5b61243a8d8d6120b0565b9a5060208c0135995060408c013598506001600160401b038060608e01351115612462578687fd5b6124728e60608f01358f01611f84565b98508060808e01351115612484578687fd5b6124948e60808f01358f01611f1e565b975060a08d0135965060c08d013595508060e08e013511156124b4578485fd5b6124c48e60e08f01358f01611f84565b9450806101008e013511156124d7578384fd5b6124e88e6101008f01358f01612037565b9350806101208e013511156124fb578283fd5b61250c8e6101208f01358f01611f84565b9250806101408e0135111561251f578182fd5b506125318d6101408e01358e01611f84565b90509295989b509295989b9093969950565b60008060008060008060008060006101208a8c031215612561578283fd5b8935985060208a01356001600160401b038082111561257e578485fd5b61258a8d838e01611ea8565b995060408c013591508082111561259f578485fd5b6125ab8d838e01611f84565b985060608c01359150808211156125c0578485fd5b6125cc8d838e01611f1e565b975060808c0135965060a08c0135955060c08c01359150808211156125ef578485fd5b6125fb8d838e01612037565b945060e08c0135915080821115612610578384fd5b61261c8d838e01611f84565b93506101008c0135915080821115612632578283fd5b5061263f8c828d01611f84565b9150509295985092959850929598565b60008060008060808587031215612664578182fd5b843593506020850135925060408501356001600160401b0380821115612688578384fd5b61269488838901611f84565b935060608701359150808211156126a9578283fd5b506126b687828801611f1e565b91505092959194509250565b6000602082840312156126d3578081fd5b81356001600160401b038111156126e8578182fd5b6126f4848285016120c0565b949350505050565b6000806040838503121561270e578182fd5b82356001600160401b0380821115612724578384fd5b612730868387016120c0565b93506020850135915080821115612745578283fd5b506123e985828601611fe2565b60008060408385031215612764578182fd5b82356001600160401b038082111561277a578384fd5b612786868387016120c0565b9350602085013591508082111561279b578283fd5b506123e985828601612128565b600080604083850312156127ba578182fd5b82356001600160401b038111156127cf578283fd5b6127db858286016120c0565b95602094909401359450505050565b6000602082840312156127fb578081fd5b5035919050565b60008060408385031215612814578182fd5b82359150602083013561282681612d28565b809150509250929050565b600080600060608486031215612845578081fd5b83359250602084013561285781612d28565b915060408401356001600160401b0380821115612872578283fd5b81860160408189031215612884578384fd5b61288e6040612ce3565b925080358281111561289e578485fd5b6128aa89828401611f84565b8452506020810135828111156128be578485fd5b6128ca89828401611f1e565b6020850152505050809150509250925092565b6000815180845260208085019450808401835b838110156129155781516001600160a01b0316875295820195908201906001016128f0565b509495945050505050565b90815260200190565b918252602082015260400190565b6001600160a01b0391909116815260200190565b60006020825261228360208301846128dd565b60006040825261297160408301856128dd565b90508260208301529392505050565b901515815260200190565b93845260ff9290921660208401526040830152606082015260800190565b6020808252601490820152736e6f7420706175736564207269676874206e6f7760601b604082015260600190565b6020808252600e908201526d1cda59d2591e081a5b9d985b1a5960921b604082015260600190565b60208082526025908201527f63616e6e6f74207375626d69742063616e64696461746520666f72206f6c6420604082015264737761707360d81b606082015260800190565b60208082526023908201527f6d7573742070726f6d6f74652063616e6469646174652073657175656e7469616040820152626c6c7960e81b606082015260800190565b602080825260079082015266195e1c1a5c995960ca1b604082015260600190565b6020808252600c908201526b696e76616c6964206d65746160a01b604082015260600190565b6020808252601a908201527f63616e6e6f7420657874656e6420666f7220746f6f206c6f6e67000000000000604082015260600190565b6020808252600c908201526b1d5b985d5d1a1bdc9a5e995960a21b604082015260600190565b6020808252602f908201527f70726f6f66206d757374206265207369676e6564206279206e657720636f6d6d60408201526e69747465652f63616e64696461746560881b606082015260800190565b6020808252601590820152741a5b9d985b1a5908189c9a5919d9481cddd85c1251605a1b604082015260600190565b60208082526010908201526f706175736564207269676874206e6f7760801b604082015260600190565b6020808252601590820152741a5b9d985b1a590818995858dbdb881cddd85c1251605a1b604082015260600190565b6020808252600990820152683737ba1030b236b4b760b91b604082015260600190565b6020808252601390820152721c1c9bdc1bdcd9551a5b59481a5b9d985b1a59606a1b604082015260600190565b600060208252825160606020840152612c6e60808401826128dd565b6020850151604085015260408501516060850152809250505092915050565b9182521515602082015260400190565b60ff9390931683526020830191909152604082015260600190565b60ff958616815293909416602084015260408301919091526060820152608081019190915260a00190565b6040518181016001600160401b0381118282101715612d0157600080fd5b604052919050565b60006001600160401b03821115612d1e578081fd5b5060209081020190565b8015158114612d3657600080fd5b5056fea264697066735822122052a0436ba1ee80c643a50ac8236c1efe23018ee77847576f0b72bdfd5008647b64736f6c63430006060033"

// DeployIncognitoProxy deploys a new Ethereum contract, binding an instance of IncognitoProxy to it.
func DeployIncognitoProxy(auth *bind.TransactOpts, backend bind.ContractBackend, admin common.Address, beaconCommittee []common.Address, bridgeCommittee []common.Address) (common.Address, *types.Transaction, *IncognitoProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(IncognitoProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IncognitoProxyBin), backend, admin, beaconCommittee, bridgeCommittee)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IncognitoProxy{IncognitoProxyCaller: IncognitoProxyCaller{contract: contract}, IncognitoProxyTransactor: IncognitoProxyTransactor{contract: contract}, IncognitoProxyFilterer: IncognitoProxyFilterer{contract: contract}}, nil
}

// IncognitoProxy is an auto generated Go binding around an Ethereum contract.
type IncognitoProxy struct {
	IncognitoProxyCaller     // Read-only binding to the contract
	IncognitoProxyTransactor // Write-only binding to the contract
	IncognitoProxyFilterer   // Log filterer for contract events
}

// IncognitoProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type IncognitoProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncognitoProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IncognitoProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncognitoProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IncognitoProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncognitoProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IncognitoProxySession struct {
	Contract     *IncognitoProxy   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IncognitoProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IncognitoProxyCallerSession struct {
	Contract *IncognitoProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IncognitoProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IncognitoProxyTransactorSession struct {
	Contract     *IncognitoProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IncognitoProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type IncognitoProxyRaw struct {
	Contract *IncognitoProxy // Generic contract binding to access the raw methods on
}

// IncognitoProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IncognitoProxyCallerRaw struct {
	Contract *IncognitoProxyCaller // Generic read-only contract binding to access the raw methods on
}

// IncognitoProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IncognitoProxyTransactorRaw struct {
	Contract *IncognitoProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIncognitoProxy creates a new instance of IncognitoProxy, bound to a specific deployed contract.
func NewIncognitoProxy(address common.Address, backend bind.ContractBackend) (*IncognitoProxy, error) {
	contract, err := bindIncognitoProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IncognitoProxy{IncognitoProxyCaller: IncognitoProxyCaller{contract: contract}, IncognitoProxyTransactor: IncognitoProxyTransactor{contract: contract}, IncognitoProxyFilterer: IncognitoProxyFilterer{contract: contract}}, nil
}

// NewIncognitoProxyCaller creates a new read-only instance of IncognitoProxy, bound to a specific deployed contract.
func NewIncognitoProxyCaller(address common.Address, caller bind.ContractCaller) (*IncognitoProxyCaller, error) {
	contract, err := bindIncognitoProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IncognitoProxyCaller{contract: contract}, nil
}

// NewIncognitoProxyTransactor creates a new write-only instance of IncognitoProxy, bound to a specific deployed contract.
func NewIncognitoProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*IncognitoProxyTransactor, error) {
	contract, err := bindIncognitoProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IncognitoProxyTransactor{contract: contract}, nil
}

// NewIncognitoProxyFilterer creates a new log filterer instance of IncognitoProxy, bound to a specific deployed contract.
func NewIncognitoProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*IncognitoProxyFilterer, error) {
	contract, err := bindIncognitoProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IncognitoProxyFilterer{contract: contract}, nil
}

// bindIncognitoProxy binds a generic wrapper to an already deployed contract.
func bindIncognitoProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IncognitoProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IncognitoProxy *IncognitoProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IncognitoProxy.Contract.IncognitoProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IncognitoProxy *IncognitoProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.IncognitoProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IncognitoProxy *IncognitoProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.IncognitoProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IncognitoProxy *IncognitoProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IncognitoProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IncognitoProxy *IncognitoProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IncognitoProxy *IncognitoProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_IncognitoProxy *IncognitoProxyCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_IncognitoProxy *IncognitoProxySession) Admin() (common.Address, error) {
	return _IncognitoProxy.Contract.Admin(&_IncognitoProxy.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_IncognitoProxy *IncognitoProxyCallerSession) Admin() (common.Address, error) {
	return _IncognitoProxy.Contract.Admin(&_IncognitoProxy.CallOpts)
}

// BeaconCandidates is a free data retrieval call binding the contract method 0x62532738.
//
// Solidity: function beaconCandidates(uint256 ) view returns(uint256 startBlock, bytes32 beaconBlockHash)
func (_IncognitoProxy *IncognitoProxyCaller) BeaconCandidates(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StartBlock      *big.Int
	BeaconBlockHash [32]byte
}, error) {
	ret := new(struct {
		StartBlock      *big.Int
		BeaconBlockHash [32]byte
	})
	out := ret
	err := _IncognitoProxy.contract.Call(opts, out, "beaconCandidates", arg0)
	return *ret, err
}

// BeaconCandidates is a free data retrieval call binding the contract method 0x62532738.
//
// Solidity: function beaconCandidates(uint256 ) view returns(uint256 startBlock, bytes32 beaconBlockHash)
func (_IncognitoProxy *IncognitoProxySession) BeaconCandidates(arg0 *big.Int) (struct {
	StartBlock      *big.Int
	BeaconBlockHash [32]byte
}, error) {
	return _IncognitoProxy.Contract.BeaconCandidates(&_IncognitoProxy.CallOpts, arg0)
}

// BeaconCandidates is a free data retrieval call binding the contract method 0x62532738.
//
// Solidity: function beaconCandidates(uint256 ) view returns(uint256 startBlock, bytes32 beaconBlockHash)
func (_IncognitoProxy *IncognitoProxyCallerSession) BeaconCandidates(arg0 *big.Int) (struct {
	StartBlock      *big.Int
	BeaconBlockHash [32]byte
}, error) {
	return _IncognitoProxy.Contract.BeaconCandidates(&_IncognitoProxy.CallOpts, arg0)
}

// BeaconCommittees is a free data retrieval call binding the contract method 0xf203a5ed.
//
// Solidity: function beaconCommittees(uint256 ) view returns(uint256 startBlock, uint256 swapID)
func (_IncognitoProxy *IncognitoProxyCaller) BeaconCommittees(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StartBlock *big.Int
	SwapID     *big.Int
}, error) {
	ret := new(struct {
		StartBlock *big.Int
		SwapID     *big.Int
	})
	out := ret
	err := _IncognitoProxy.contract.Call(opts, out, "beaconCommittees", arg0)
	return *ret, err
}

// BeaconCommittees is a free data retrieval call binding the contract method 0xf203a5ed.
//
// Solidity: function beaconCommittees(uint256 ) view returns(uint256 startBlock, uint256 swapID)
func (_IncognitoProxy *IncognitoProxySession) BeaconCommittees(arg0 *big.Int) (struct {
	StartBlock *big.Int
	SwapID     *big.Int
}, error) {
	return _IncognitoProxy.Contract.BeaconCommittees(&_IncognitoProxy.CallOpts, arg0)
}

// BeaconCommittees is a free data retrieval call binding the contract method 0xf203a5ed.
//
// Solidity: function beaconCommittees(uint256 ) view returns(uint256 startBlock, uint256 swapID)
func (_IncognitoProxy *IncognitoProxyCallerSession) BeaconCommittees(arg0 *big.Int) (struct {
	StartBlock *big.Int
	SwapID     *big.Int
}, error) {
	return _IncognitoProxy.Contract.BeaconCommittees(&_IncognitoProxy.CallOpts, arg0)
}

// BeaconFinality is a free data retrieval call binding the contract method 0x5d7f86d9.
//
// Solidity: function beaconFinality() view returns(uint256 blockHeight, bytes32 rootHash)
func (_IncognitoProxy *IncognitoProxyCaller) BeaconFinality(opts *bind.CallOpts) (struct {
	BlockHeight *big.Int
	RootHash    [32]byte
}, error) {
	ret := new(struct {
		BlockHeight *big.Int
		RootHash    [32]byte
	})
	out := ret
	err := _IncognitoProxy.contract.Call(opts, out, "beaconFinality")
	return *ret, err
}

// BeaconFinality is a free data retrieval call binding the contract method 0x5d7f86d9.
//
// Solidity: function beaconFinality() view returns(uint256 blockHeight, bytes32 rootHash)
func (_IncognitoProxy *IncognitoProxySession) BeaconFinality() (struct {
	BlockHeight *big.Int
	RootHash    [32]byte
}, error) {
	return _IncognitoProxy.Contract.BeaconFinality(&_IncognitoProxy.CallOpts)
}

// BeaconFinality is a free data retrieval call binding the contract method 0x5d7f86d9.
//
// Solidity: function beaconFinality() view returns(uint256 blockHeight, bytes32 rootHash)
func (_IncognitoProxy *IncognitoProxyCallerSession) BeaconFinality() (struct {
	BlockHeight *big.Int
	RootHash    [32]byte
}, error) {
	return _IncognitoProxy.Contract.BeaconFinality(&_IncognitoProxy.CallOpts)
}

// BridgeCandidates is a free data retrieval call binding the contract method 0xe83d31ec.
//
// Solidity: function bridgeCandidates(uint256 ) view returns(uint256 startBlock, bytes32 beaconBlockHash)
func (_IncognitoProxy *IncognitoProxyCaller) BridgeCandidates(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StartBlock      *big.Int
	BeaconBlockHash [32]byte
}, error) {
	ret := new(struct {
		StartBlock      *big.Int
		BeaconBlockHash [32]byte
	})
	out := ret
	err := _IncognitoProxy.contract.Call(opts, out, "bridgeCandidates", arg0)
	return *ret, err
}

// BridgeCandidates is a free data retrieval call binding the contract method 0xe83d31ec.
//
// Solidity: function bridgeCandidates(uint256 ) view returns(uint256 startBlock, bytes32 beaconBlockHash)
func (_IncognitoProxy *IncognitoProxySession) BridgeCandidates(arg0 *big.Int) (struct {
	StartBlock      *big.Int
	BeaconBlockHash [32]byte
}, error) {
	return _IncognitoProxy.Contract.BridgeCandidates(&_IncognitoProxy.CallOpts, arg0)
}

// BridgeCandidates is a free data retrieval call binding the contract method 0xe83d31ec.
//
// Solidity: function bridgeCandidates(uint256 ) view returns(uint256 startBlock, bytes32 beaconBlockHash)
func (_IncognitoProxy *IncognitoProxyCallerSession) BridgeCandidates(arg0 *big.Int) (struct {
	StartBlock      *big.Int
	BeaconBlockHash [32]byte
}, error) {
	return _IncognitoProxy.Contract.BridgeCandidates(&_IncognitoProxy.CallOpts, arg0)
}

// BridgeCommittees is a free data retrieval call binding the contract method 0x9b30b637.
//
// Solidity: function bridgeCommittees(uint256 ) view returns(uint256 startBlock, uint256 swapID)
func (_IncognitoProxy *IncognitoProxyCaller) BridgeCommittees(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StartBlock *big.Int
	SwapID     *big.Int
}, error) {
	ret := new(struct {
		StartBlock *big.Int
		SwapID     *big.Int
	})
	out := ret
	err := _IncognitoProxy.contract.Call(opts, out, "bridgeCommittees", arg0)
	return *ret, err
}

// BridgeCommittees is a free data retrieval call binding the contract method 0x9b30b637.
//
// Solidity: function bridgeCommittees(uint256 ) view returns(uint256 startBlock, uint256 swapID)
func (_IncognitoProxy *IncognitoProxySession) BridgeCommittees(arg0 *big.Int) (struct {
	StartBlock *big.Int
	SwapID     *big.Int
}, error) {
	return _IncognitoProxy.Contract.BridgeCommittees(&_IncognitoProxy.CallOpts, arg0)
}

// BridgeCommittees is a free data retrieval call binding the contract method 0x9b30b637.
//
// Solidity: function bridgeCommittees(uint256 ) view returns(uint256 startBlock, uint256 swapID)
func (_IncognitoProxy *IncognitoProxyCallerSession) BridgeCommittees(arg0 *big.Int) (struct {
	StartBlock *big.Int
	SwapID     *big.Int
}, error) {
	return _IncognitoProxy.Contract.BridgeCommittees(&_IncognitoProxy.CallOpts, arg0)
}

// BridgeFinality is a free data retrieval call binding the contract method 0x2c1c9cac.
//
// Solidity: function bridgeFinality() view returns(uint256 blockHeight, bytes32 rootHash)
func (_IncognitoProxy *IncognitoProxyCaller) BridgeFinality(opts *bind.CallOpts) (struct {
	BlockHeight *big.Int
	RootHash    [32]byte
}, error) {
	ret := new(struct {
		BlockHeight *big.Int
		RootHash    [32]byte
	})
	out := ret
	err := _IncognitoProxy.contract.Call(opts, out, "bridgeFinality")
	return *ret, err
}

// BridgeFinality is a free data retrieval call binding the contract method 0x2c1c9cac.
//
// Solidity: function bridgeFinality() view returns(uint256 blockHeight, bytes32 rootHash)
func (_IncognitoProxy *IncognitoProxySession) BridgeFinality() (struct {
	BlockHeight *big.Int
	RootHash    [32]byte
}, error) {
	return _IncognitoProxy.Contract.BridgeFinality(&_IncognitoProxy.CallOpts)
}

// BridgeFinality is a free data retrieval call binding the contract method 0x2c1c9cac.
//
// Solidity: function bridgeFinality() view returns(uint256 blockHeight, bytes32 rootHash)
func (_IncognitoProxy *IncognitoProxyCallerSession) BridgeFinality() (struct {
	BlockHeight *big.Int
	RootHash    [32]byte
}, error) {
	return _IncognitoProxy.Contract.BridgeFinality(&_IncognitoProxy.CallOpts)
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() view returns(uint256)
func (_IncognitoProxy *IncognitoProxyCaller) Expire(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "expire")
	return *ret0, err
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() view returns(uint256)
func (_IncognitoProxy *IncognitoProxySession) Expire() (*big.Int, error) {
	return _IncognitoProxy.Contract.Expire(&_IncognitoProxy.CallOpts)
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() view returns(uint256)
func (_IncognitoProxy *IncognitoProxyCallerSession) Expire() (*big.Int, error) {
	return _IncognitoProxy.Contract.Expire(&_IncognitoProxy.CallOpts)
}

// ExtractCommitteeFromInstruction is a free data retrieval call binding the contract method 0x8eb60066.
//
// Solidity: function extractCommitteeFromInstruction(bytes inst, uint256 numVals) pure returns(address[])
func (_IncognitoProxy *IncognitoProxyCaller) ExtractCommitteeFromInstruction(opts *bind.CallOpts, inst []byte, numVals *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "extractCommitteeFromInstruction", inst, numVals)
	return *ret0, err
}

// ExtractCommitteeFromInstruction is a free data retrieval call binding the contract method 0x8eb60066.
//
// Solidity: function extractCommitteeFromInstruction(bytes inst, uint256 numVals) pure returns(address[])
func (_IncognitoProxy *IncognitoProxySession) ExtractCommitteeFromInstruction(inst []byte, numVals *big.Int) ([]common.Address, error) {
	return _IncognitoProxy.Contract.ExtractCommitteeFromInstruction(&_IncognitoProxy.CallOpts, inst, numVals)
}

// ExtractCommitteeFromInstruction is a free data retrieval call binding the contract method 0x8eb60066.
//
// Solidity: function extractCommitteeFromInstruction(bytes inst, uint256 numVals) pure returns(address[])
func (_IncognitoProxy *IncognitoProxyCallerSession) ExtractCommitteeFromInstruction(inst []byte, numVals *big.Int) ([]common.Address, error) {
	return _IncognitoProxy.Contract.ExtractCommitteeFromInstruction(&_IncognitoProxy.CallOpts, inst, numVals)
}

// ExtractDataFromBlockMerkleInstruction is a free data retrieval call binding the contract method 0xf179dc78.
//
// Solidity: function extractDataFromBlockMerkleInstruction(bytes inst) pure returns(uint8, bytes32, uint256)
func (_IncognitoProxy *IncognitoProxyCaller) ExtractDataFromBlockMerkleInstruction(opts *bind.CallOpts, inst []byte) (uint8, [32]byte, *big.Int, error) {
	var (
		ret0 = new(uint8)
		ret1 = new([32]byte)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _IncognitoProxy.contract.Call(opts, out, "extractDataFromBlockMerkleInstruction", inst)
	return *ret0, *ret1, *ret2, err
}

// ExtractDataFromBlockMerkleInstruction is a free data retrieval call binding the contract method 0xf179dc78.
//
// Solidity: function extractDataFromBlockMerkleInstruction(bytes inst) pure returns(uint8, bytes32, uint256)
func (_IncognitoProxy *IncognitoProxySession) ExtractDataFromBlockMerkleInstruction(inst []byte) (uint8, [32]byte, *big.Int, error) {
	return _IncognitoProxy.Contract.ExtractDataFromBlockMerkleInstruction(&_IncognitoProxy.CallOpts, inst)
}

// ExtractDataFromBlockMerkleInstruction is a free data retrieval call binding the contract method 0xf179dc78.
//
// Solidity: function extractDataFromBlockMerkleInstruction(bytes inst) pure returns(uint8, bytes32, uint256)
func (_IncognitoProxy *IncognitoProxyCallerSession) ExtractDataFromBlockMerkleInstruction(inst []byte) (uint8, [32]byte, *big.Int, error) {
	return _IncognitoProxy.Contract.ExtractDataFromBlockMerkleInstruction(&_IncognitoProxy.CallOpts, inst)
}

// ExtractMetaFromInstruction is a free data retrieval call binding the contract method 0x90500bae.
//
// Solidity: function extractMetaFromInstruction(bytes inst) pure returns(uint8, uint8, uint256, uint256, uint256)
func (_IncognitoProxy *IncognitoProxyCaller) ExtractMetaFromInstruction(opts *bind.CallOpts, inst []byte) (uint8, uint8, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(uint8)
		ret1 = new(uint8)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
		ret4 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _IncognitoProxy.contract.Call(opts, out, "extractMetaFromInstruction", inst)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// ExtractMetaFromInstruction is a free data retrieval call binding the contract method 0x90500bae.
//
// Solidity: function extractMetaFromInstruction(bytes inst) pure returns(uint8, uint8, uint256, uint256, uint256)
func (_IncognitoProxy *IncognitoProxySession) ExtractMetaFromInstruction(inst []byte) (uint8, uint8, *big.Int, *big.Int, *big.Int, error) {
	return _IncognitoProxy.Contract.ExtractMetaFromInstruction(&_IncognitoProxy.CallOpts, inst)
}

// ExtractMetaFromInstruction is a free data retrieval call binding the contract method 0x90500bae.
//
// Solidity: function extractMetaFromInstruction(bytes inst) pure returns(uint8, uint8, uint256, uint256, uint256)
func (_IncognitoProxy *IncognitoProxyCallerSession) ExtractMetaFromInstruction(inst []byte) (uint8, uint8, *big.Int, *big.Int, *big.Int, error) {
	return _IncognitoProxy.Contract.ExtractMetaFromInstruction(&_IncognitoProxy.CallOpts, inst)
}

// FilterSigners is a free data retrieval call binding the contract method 0x6ae474d8.
//
// Solidity: function filterSigners(uint256[] sigIdx, address[] signers) pure returns(address[])
func (_IncognitoProxy *IncognitoProxyCaller) FilterSigners(opts *bind.CallOpts, sigIdx []*big.Int, signers []common.Address) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "filterSigners", sigIdx, signers)
	return *ret0, err
}

// FilterSigners is a free data retrieval call binding the contract method 0x6ae474d8.
//
// Solidity: function filterSigners(uint256[] sigIdx, address[] signers) pure returns(address[])
func (_IncognitoProxy *IncognitoProxySession) FilterSigners(sigIdx []*big.Int, signers []common.Address) ([]common.Address, error) {
	return _IncognitoProxy.Contract.FilterSigners(&_IncognitoProxy.CallOpts, sigIdx, signers)
}

// FilterSigners is a free data retrieval call binding the contract method 0x6ae474d8.
//
// Solidity: function filterSigners(uint256[] sigIdx, address[] signers) pure returns(address[])
func (_IncognitoProxy *IncognitoProxyCallerSession) FilterSigners(sigIdx []*big.Int, signers []common.Address) ([]common.Address, error) {
	return _IncognitoProxy.Contract.FilterSigners(&_IncognitoProxy.CallOpts, sigIdx, signers)
}

// FindBeaconCommitteeFromHeight is a free data retrieval call binding the contract method 0xb600ffdb.
//
// Solidity: function findBeaconCommitteeFromHeight(uint256 blkHeight) view returns(address[], uint256)
func (_IncognitoProxy *IncognitoProxyCaller) FindBeaconCommitteeFromHeight(opts *bind.CallOpts, blkHeight *big.Int) ([]common.Address, *big.Int, error) {
	var (
		ret0 = new([]common.Address)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _IncognitoProxy.contract.Call(opts, out, "findBeaconCommitteeFromHeight", blkHeight)
	return *ret0, *ret1, err
}

// FindBeaconCommitteeFromHeight is a free data retrieval call binding the contract method 0xb600ffdb.
//
// Solidity: function findBeaconCommitteeFromHeight(uint256 blkHeight) view returns(address[], uint256)
func (_IncognitoProxy *IncognitoProxySession) FindBeaconCommitteeFromHeight(blkHeight *big.Int) ([]common.Address, *big.Int, error) {
	return _IncognitoProxy.Contract.FindBeaconCommitteeFromHeight(&_IncognitoProxy.CallOpts, blkHeight)
}

// FindBeaconCommitteeFromHeight is a free data retrieval call binding the contract method 0xb600ffdb.
//
// Solidity: function findBeaconCommitteeFromHeight(uint256 blkHeight) view returns(address[], uint256)
func (_IncognitoProxy *IncognitoProxyCallerSession) FindBeaconCommitteeFromHeight(blkHeight *big.Int) ([]common.Address, *big.Int, error) {
	return _IncognitoProxy.Contract.FindBeaconCommitteeFromHeight(&_IncognitoProxy.CallOpts, blkHeight)
}

// FindBridgeCommitteeFromHeight is a free data retrieval call binding the contract method 0xf5205fde.
//
// Solidity: function findBridgeCommitteeFromHeight(uint256 blkHeight) view returns(address[], uint256)
func (_IncognitoProxy *IncognitoProxyCaller) FindBridgeCommitteeFromHeight(opts *bind.CallOpts, blkHeight *big.Int) ([]common.Address, *big.Int, error) {
	var (
		ret0 = new([]common.Address)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _IncognitoProxy.contract.Call(opts, out, "findBridgeCommitteeFromHeight", blkHeight)
	return *ret0, *ret1, err
}

// FindBridgeCommitteeFromHeight is a free data retrieval call binding the contract method 0xf5205fde.
//
// Solidity: function findBridgeCommitteeFromHeight(uint256 blkHeight) view returns(address[], uint256)
func (_IncognitoProxy *IncognitoProxySession) FindBridgeCommitteeFromHeight(blkHeight *big.Int) ([]common.Address, *big.Int, error) {
	return _IncognitoProxy.Contract.FindBridgeCommitteeFromHeight(&_IncognitoProxy.CallOpts, blkHeight)
}

// FindBridgeCommitteeFromHeight is a free data retrieval call binding the contract method 0xf5205fde.
//
// Solidity: function findBridgeCommitteeFromHeight(uint256 blkHeight) view returns(address[], uint256)
func (_IncognitoProxy *IncognitoProxyCallerSession) FindBridgeCommitteeFromHeight(blkHeight *big.Int) ([]common.Address, *big.Int, error) {
	return _IncognitoProxy.Contract.FindBridgeCommitteeFromHeight(&_IncognitoProxy.CallOpts, blkHeight)
}

// GetBeaconCommittee is a free data retrieval call binding the contract method 0xfaea3167.
//
// Solidity: function getBeaconCommittee(uint256 i) view returns((address[],uint256,uint256))
func (_IncognitoProxy *IncognitoProxyCaller) GetBeaconCommittee(opts *bind.CallOpts, i *big.Int) (IncognitoProxyCommittee, error) {
	var (
		ret0 = new(IncognitoProxyCommittee)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "getBeaconCommittee", i)
	return *ret0, err
}

// GetBeaconCommittee is a free data retrieval call binding the contract method 0xfaea3167.
//
// Solidity: function getBeaconCommittee(uint256 i) view returns((address[],uint256,uint256))
func (_IncognitoProxy *IncognitoProxySession) GetBeaconCommittee(i *big.Int) (IncognitoProxyCommittee, error) {
	return _IncognitoProxy.Contract.GetBeaconCommittee(&_IncognitoProxy.CallOpts, i)
}

// GetBeaconCommittee is a free data retrieval call binding the contract method 0xfaea3167.
//
// Solidity: function getBeaconCommittee(uint256 i) view returns((address[],uint256,uint256))
func (_IncognitoProxy *IncognitoProxyCallerSession) GetBeaconCommittee(i *big.Int) (IncognitoProxyCommittee, error) {
	return _IncognitoProxy.Contract.GetBeaconCommittee(&_IncognitoProxy.CallOpts, i)
}

// GetBridgeCommittee is a free data retrieval call binding the contract method 0x8ceb69c3.
//
// Solidity: function getBridgeCommittee(uint256 i) view returns((address[],uint256,uint256))
func (_IncognitoProxy *IncognitoProxyCaller) GetBridgeCommittee(opts *bind.CallOpts, i *big.Int) (IncognitoProxyCommittee, error) {
	var (
		ret0 = new(IncognitoProxyCommittee)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "getBridgeCommittee", i)
	return *ret0, err
}

// GetBridgeCommittee is a free data retrieval call binding the contract method 0x8ceb69c3.
//
// Solidity: function getBridgeCommittee(uint256 i) view returns((address[],uint256,uint256))
func (_IncognitoProxy *IncognitoProxySession) GetBridgeCommittee(i *big.Int) (IncognitoProxyCommittee, error) {
	return _IncognitoProxy.Contract.GetBridgeCommittee(&_IncognitoProxy.CallOpts, i)
}

// GetBridgeCommittee is a free data retrieval call binding the contract method 0x8ceb69c3.
//
// Solidity: function getBridgeCommittee(uint256 i) view returns((address[],uint256,uint256))
func (_IncognitoProxy *IncognitoProxyCallerSession) GetBridgeCommittee(i *big.Int) (IncognitoProxyCommittee, error) {
	return _IncognitoProxy.Contract.GetBridgeCommittee(&_IncognitoProxy.CallOpts, i)
}

// InstructionApproved is a free data retrieval call binding the contract method 0xf65d2116.
//
// Solidity: function instructionApproved(bool isBeacon, bytes32 instHash, uint256 blkHeight, bytes32[] instPath, bool[] instPathIsLeft, bytes32 instRoot, bytes32 blkData, uint256[] sigIdx, uint8[] sigV, bytes32[] sigR, bytes32[] sigS) view returns(bool)
func (_IncognitoProxy *IncognitoProxyCaller) InstructionApproved(opts *bind.CallOpts, isBeacon bool, instHash [32]byte, blkHeight *big.Int, instPath [][32]byte, instPathIsLeft []bool, instRoot [32]byte, blkData [32]byte, sigIdx []*big.Int, sigV []uint8, sigR [][32]byte, sigS [][32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "instructionApproved", isBeacon, instHash, blkHeight, instPath, instPathIsLeft, instRoot, blkData, sigIdx, sigV, sigR, sigS)
	return *ret0, err
}

// InstructionApproved is a free data retrieval call binding the contract method 0xf65d2116.
//
// Solidity: function instructionApproved(bool isBeacon, bytes32 instHash, uint256 blkHeight, bytes32[] instPath, bool[] instPathIsLeft, bytes32 instRoot, bytes32 blkData, uint256[] sigIdx, uint8[] sigV, bytes32[] sigR, bytes32[] sigS) view returns(bool)
func (_IncognitoProxy *IncognitoProxySession) InstructionApproved(isBeacon bool, instHash [32]byte, blkHeight *big.Int, instPath [][32]byte, instPathIsLeft []bool, instRoot [32]byte, blkData [32]byte, sigIdx []*big.Int, sigV []uint8, sigR [][32]byte, sigS [][32]byte) (bool, error) {
	return _IncognitoProxy.Contract.InstructionApproved(&_IncognitoProxy.CallOpts, isBeacon, instHash, blkHeight, instPath, instPathIsLeft, instRoot, blkData, sigIdx, sigV, sigR, sigS)
}

// InstructionApproved is a free data retrieval call binding the contract method 0xf65d2116.
//
// Solidity: function instructionApproved(bool isBeacon, bytes32 instHash, uint256 blkHeight, bytes32[] instPath, bool[] instPathIsLeft, bytes32 instRoot, bytes32 blkData, uint256[] sigIdx, uint8[] sigV, bytes32[] sigR, bytes32[] sigS) view returns(bool)
func (_IncognitoProxy *IncognitoProxyCallerSession) InstructionApproved(isBeacon bool, instHash [32]byte, blkHeight *big.Int, instPath [][32]byte, instPathIsLeft []bool, instRoot [32]byte, blkData [32]byte, sigIdx []*big.Int, sigV []uint8, sigR [][32]byte, sigS [][32]byte) (bool, error) {
	return _IncognitoProxy.Contract.InstructionApproved(&_IncognitoProxy.CallOpts, isBeacon, instHash, blkHeight, instPath, instPathIsLeft, instRoot, blkData, sigIdx, sigV, sigR, sigS)
}

// InstructionApprovedBySigners is a free data retrieval call binding the contract method 0x3e379d05.
//
// Solidity: function instructionApprovedBySigners(bytes32 instHash, address[] signers, bytes32[] instPath, bool[] instPathIsLeft, bytes32 instRoot, bytes32 blkData, uint8[] sigV, bytes32[] sigR, bytes32[] sigS) view returns(bool)
func (_IncognitoProxy *IncognitoProxyCaller) InstructionApprovedBySigners(opts *bind.CallOpts, instHash [32]byte, signers []common.Address, instPath [][32]byte, instPathIsLeft []bool, instRoot [32]byte, blkData [32]byte, sigV []uint8, sigR [][32]byte, sigS [][32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "instructionApprovedBySigners", instHash, signers, instPath, instPathIsLeft, instRoot, blkData, sigV, sigR, sigS)
	return *ret0, err
}

// InstructionApprovedBySigners is a free data retrieval call binding the contract method 0x3e379d05.
//
// Solidity: function instructionApprovedBySigners(bytes32 instHash, address[] signers, bytes32[] instPath, bool[] instPathIsLeft, bytes32 instRoot, bytes32 blkData, uint8[] sigV, bytes32[] sigR, bytes32[] sigS) view returns(bool)
func (_IncognitoProxy *IncognitoProxySession) InstructionApprovedBySigners(instHash [32]byte, signers []common.Address, instPath [][32]byte, instPathIsLeft []bool, instRoot [32]byte, blkData [32]byte, sigV []uint8, sigR [][32]byte, sigS [][32]byte) (bool, error) {
	return _IncognitoProxy.Contract.InstructionApprovedBySigners(&_IncognitoProxy.CallOpts, instHash, signers, instPath, instPathIsLeft, instRoot, blkData, sigV, sigR, sigS)
}

// InstructionApprovedBySigners is a free data retrieval call binding the contract method 0x3e379d05.
//
// Solidity: function instructionApprovedBySigners(bytes32 instHash, address[] signers, bytes32[] instPath, bool[] instPathIsLeft, bytes32 instRoot, bytes32 blkData, uint8[] sigV, bytes32[] sigR, bytes32[] sigS) view returns(bool)
func (_IncognitoProxy *IncognitoProxyCallerSession) InstructionApprovedBySigners(instHash [32]byte, signers []common.Address, instPath [][32]byte, instPathIsLeft []bool, instRoot [32]byte, blkData [32]byte, sigV []uint8, sigR [][32]byte, sigS [][32]byte) (bool, error) {
	return _IncognitoProxy.Contract.InstructionApprovedBySigners(&_IncognitoProxy.CallOpts, instHash, signers, instPath, instPathIsLeft, instRoot, blkData, sigV, sigR, sigS)
}

// LeafInMerkleTree is a free data retrieval call binding the contract method 0x1217a3c7.
//
// Solidity: function leafInMerkleTree(bytes32 leaf, bytes32 root, bytes32[] path, bool[] left) pure returns(bool)
func (_IncognitoProxy *IncognitoProxyCaller) LeafInMerkleTree(opts *bind.CallOpts, leaf [32]byte, root [32]byte, path [][32]byte, left []bool) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "leafInMerkleTree", leaf, root, path, left)
	return *ret0, err
}

// LeafInMerkleTree is a free data retrieval call binding the contract method 0x1217a3c7.
//
// Solidity: function leafInMerkleTree(bytes32 leaf, bytes32 root, bytes32[] path, bool[] left) pure returns(bool)
func (_IncognitoProxy *IncognitoProxySession) LeafInMerkleTree(leaf [32]byte, root [32]byte, path [][32]byte, left []bool) (bool, error) {
	return _IncognitoProxy.Contract.LeafInMerkleTree(&_IncognitoProxy.CallOpts, leaf, root, path, left)
}

// LeafInMerkleTree is a free data retrieval call binding the contract method 0x1217a3c7.
//
// Solidity: function leafInMerkleTree(bytes32 leaf, bytes32 root, bytes32[] path, bool[] left) pure returns(bool)
func (_IncognitoProxy *IncognitoProxyCallerSession) LeafInMerkleTree(leaf [32]byte, root [32]byte, path [][32]byte, left []bool) (bool, error) {
	return _IncognitoProxy.Contract.LeafInMerkleTree(&_IncognitoProxy.CallOpts, leaf, root, path, left)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IncognitoProxy *IncognitoProxyCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IncognitoProxy *IncognitoProxySession) Paused() (bool, error) {
	return _IncognitoProxy.Contract.Paused(&_IncognitoProxy.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IncognitoProxy *IncognitoProxyCallerSession) Paused() (bool, error) {
	return _IncognitoProxy.Contract.Paused(&_IncognitoProxy.CallOpts)
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() view returns(address)
func (_IncognitoProxy *IncognitoProxyCaller) Successor(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "successor")
	return *ret0, err
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() view returns(address)
func (_IncognitoProxy *IncognitoProxySession) Successor() (common.Address, error) {
	return _IncognitoProxy.Contract.Successor(&_IncognitoProxy.CallOpts)
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() view returns(address)
func (_IncognitoProxy *IncognitoProxyCallerSession) Successor() (common.Address, error) {
	return _IncognitoProxy.Contract.Successor(&_IncognitoProxy.CallOpts)
}

// VerifySig is a free data retrieval call binding the contract method 0x3aacfdad.
//
// Solidity: function verifySig(address[] committee, bytes32 msgHash, uint8[] v, bytes32[] r, bytes32[] s) pure returns(bool)
func (_IncognitoProxy *IncognitoProxyCaller) VerifySig(opts *bind.CallOpts, committee []common.Address, msgHash [32]byte, v []uint8, r [][32]byte, s [][32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "verifySig", committee, msgHash, v, r, s)
	return *ret0, err
}

// VerifySig is a free data retrieval call binding the contract method 0x3aacfdad.
//
// Solidity: function verifySig(address[] committee, bytes32 msgHash, uint8[] v, bytes32[] r, bytes32[] s) pure returns(bool)
func (_IncognitoProxy *IncognitoProxySession) VerifySig(committee []common.Address, msgHash [32]byte, v []uint8, r [][32]byte, s [][32]byte) (bool, error) {
	return _IncognitoProxy.Contract.VerifySig(&_IncognitoProxy.CallOpts, committee, msgHash, v, r, s)
}

// VerifySig is a free data retrieval call binding the contract method 0x3aacfdad.
//
// Solidity: function verifySig(address[] committee, bytes32 msgHash, uint8[] v, bytes32[] r, bytes32[] s) pure returns(bool)
func (_IncognitoProxy *IncognitoProxyCallerSession) VerifySig(committee []common.Address, msgHash [32]byte, v []uint8, r [][32]byte, s [][32]byte) (bool, error) {
	return _IncognitoProxy.Contract.VerifySig(&_IncognitoProxy.CallOpts, committee, msgHash, v, r, s)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_IncognitoProxy *IncognitoProxyTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncognitoProxy.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_IncognitoProxy *IncognitoProxySession) Claim() (*types.Transaction, error) {
	return _IncognitoProxy.Contract.Claim(&_IncognitoProxy.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_IncognitoProxy *IncognitoProxyTransactorSession) Claim() (*types.Transaction, error) {
	return _IncognitoProxy.Contract.Claim(&_IncognitoProxy.TransactOpts)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_IncognitoProxy *IncognitoProxyTransactor) Extend(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _IncognitoProxy.contract.Transact(opts, "extend", n)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_IncognitoProxy *IncognitoProxySession) Extend(n *big.Int) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.Extend(&_IncognitoProxy.TransactOpts, n)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_IncognitoProxy *IncognitoProxyTransactorSession) Extend(n *big.Int) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.Extend(&_IncognitoProxy.TransactOpts, n)
}

// GetLatestSwapID is a paid mutator transaction binding the contract method 0x722a5d7b.
//
// Solidity: function getLatestSwapID(bool isBeacon) returns(uint256)
func (_IncognitoProxy *IncognitoProxyTransactor) GetLatestSwapID(opts *bind.TransactOpts, isBeacon bool) (*types.Transaction, error) {
	return _IncognitoProxy.contract.Transact(opts, "getLatestSwapID", isBeacon)
}

// GetLatestSwapID is a paid mutator transaction binding the contract method 0x722a5d7b.
//
// Solidity: function getLatestSwapID(bool isBeacon) returns(uint256)
func (_IncognitoProxy *IncognitoProxySession) GetLatestSwapID(isBeacon bool) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.GetLatestSwapID(&_IncognitoProxy.TransactOpts, isBeacon)
}

// GetLatestSwapID is a paid mutator transaction binding the contract method 0x722a5d7b.
//
// Solidity: function getLatestSwapID(bool isBeacon) returns(uint256)
func (_IncognitoProxy *IncognitoProxyTransactorSession) GetLatestSwapID(isBeacon bool) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.GetLatestSwapID(&_IncognitoProxy.TransactOpts, isBeacon)
}

// LoadCandidates is a paid mutator transaction binding the contract method 0x6e845d19.
//
// Solidity: function loadCandidates(uint256 swapID, bool isBeacon) returns(address[])
func (_IncognitoProxy *IncognitoProxyTransactor) LoadCandidates(opts *bind.TransactOpts, swapID *big.Int, isBeacon bool) (*types.Transaction, error) {
	return _IncognitoProxy.contract.Transact(opts, "loadCandidates", swapID, isBeacon)
}

// LoadCandidates is a paid mutator transaction binding the contract method 0x6e845d19.
//
// Solidity: function loadCandidates(uint256 swapID, bool isBeacon) returns(address[])
func (_IncognitoProxy *IncognitoProxySession) LoadCandidates(swapID *big.Int, isBeacon bool) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.LoadCandidates(&_IncognitoProxy.TransactOpts, swapID, isBeacon)
}

// LoadCandidates is a paid mutator transaction binding the contract method 0x6e845d19.
//
// Solidity: function loadCandidates(uint256 swapID, bool isBeacon) returns(address[])
func (_IncognitoProxy *IncognitoProxyTransactorSession) LoadCandidates(swapID *big.Int, isBeacon bool) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.LoadCandidates(&_IncognitoProxy.TransactOpts, swapID, isBeacon)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IncognitoProxy *IncognitoProxyTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncognitoProxy.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IncognitoProxy *IncognitoProxySession) Pause() (*types.Transaction, error) {
	return _IncognitoProxy.Contract.Pause(&_IncognitoProxy.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IncognitoProxy *IncognitoProxyTransactorSession) Pause() (*types.Transaction, error) {
	return _IncognitoProxy.Contract.Pause(&_IncognitoProxy.TransactOpts)
}

// PromoteCandidate is a paid mutator transaction binding the contract method 0xea200467.
//
// Solidity: function promoteCandidate(uint256 swapID, bool isBeacon, (bytes32[],bool[]) proof) returns()
func (_IncognitoProxy *IncognitoProxyTransactor) PromoteCandidate(opts *bind.TransactOpts, swapID *big.Int, isBeacon bool, proof IncognitoProxyMerkleProof) (*types.Transaction, error) {
	return _IncognitoProxy.contract.Transact(opts, "promoteCandidate", swapID, isBeacon, proof)
}

// PromoteCandidate is a paid mutator transaction binding the contract method 0xea200467.
//
// Solidity: function promoteCandidate(uint256 swapID, bool isBeacon, (bytes32[],bool[]) proof) returns()
func (_IncognitoProxy *IncognitoProxySession) PromoteCandidate(swapID *big.Int, isBeacon bool, proof IncognitoProxyMerkleProof) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.PromoteCandidate(&_IncognitoProxy.TransactOpts, swapID, isBeacon, proof)
}

// PromoteCandidate is a paid mutator transaction binding the contract method 0xea200467.
//
// Solidity: function promoteCandidate(uint256 swapID, bool isBeacon, (bytes32[],bool[]) proof) returns()
func (_IncognitoProxy *IncognitoProxyTransactorSession) PromoteCandidate(swapID *big.Int, isBeacon bool, proof IncognitoProxyMerkleProof) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.PromoteCandidate(&_IncognitoProxy.TransactOpts, swapID, isBeacon, proof)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_IncognitoProxy *IncognitoProxyTransactor) Retire(opts *bind.TransactOpts, _successor common.Address) (*types.Transaction, error) {
	return _IncognitoProxy.contract.Transact(opts, "retire", _successor)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_IncognitoProxy *IncognitoProxySession) Retire(_successor common.Address) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.Retire(&_IncognitoProxy.TransactOpts, _successor)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_IncognitoProxy *IncognitoProxyTransactorSession) Retire(_successor common.Address) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.Retire(&_IncognitoProxy.TransactOpts, _successor)
}

// SubmitBeaconCandidate is a paid mutator transaction binding the contract method 0x52d45239.
//
// Solidity: function submitBeaconCandidate(bytes inst, (bytes,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[]) instProof) returns()
func (_IncognitoProxy *IncognitoProxyTransactor) SubmitBeaconCandidate(opts *bind.TransactOpts, inst []byte, instProof IncognitoProxyInstructionProof) (*types.Transaction, error) {
	return _IncognitoProxy.contract.Transact(opts, "submitBeaconCandidate", inst, instProof)
}

// SubmitBeaconCandidate is a paid mutator transaction binding the contract method 0x52d45239.
//
// Solidity: function submitBeaconCandidate(bytes inst, (bytes,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[]) instProof) returns()
func (_IncognitoProxy *IncognitoProxySession) SubmitBeaconCandidate(inst []byte, instProof IncognitoProxyInstructionProof) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.SubmitBeaconCandidate(&_IncognitoProxy.TransactOpts, inst, instProof)
}

// SubmitBeaconCandidate is a paid mutator transaction binding the contract method 0x52d45239.
//
// Solidity: function submitBeaconCandidate(bytes inst, (bytes,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[]) instProof) returns()
func (_IncognitoProxy *IncognitoProxyTransactorSession) SubmitBeaconCandidate(inst []byte, instProof IncognitoProxyInstructionProof) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.SubmitBeaconCandidate(&_IncognitoProxy.TransactOpts, inst, instProof)
}

// SubmitBridgeCandidate is a paid mutator transaction binding the contract method 0x7e59f5bf.
//
// Solidity: function submitBridgeCandidate(bytes inst, (bytes,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])[2] instProofs) returns()
func (_IncognitoProxy *IncognitoProxyTransactor) SubmitBridgeCandidate(opts *bind.TransactOpts, inst []byte, instProofs [2]IncognitoProxyInstructionProof) (*types.Transaction, error) {
	return _IncognitoProxy.contract.Transact(opts, "submitBridgeCandidate", inst, instProofs)
}

// SubmitBridgeCandidate is a paid mutator transaction binding the contract method 0x7e59f5bf.
//
// Solidity: function submitBridgeCandidate(bytes inst, (bytes,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])[2] instProofs) returns()
func (_IncognitoProxy *IncognitoProxySession) SubmitBridgeCandidate(inst []byte, instProofs [2]IncognitoProxyInstructionProof) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.SubmitBridgeCandidate(&_IncognitoProxy.TransactOpts, inst, instProofs)
}

// SubmitBridgeCandidate is a paid mutator transaction binding the contract method 0x7e59f5bf.
//
// Solidity: function submitBridgeCandidate(bytes inst, (bytes,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])[2] instProofs) returns()
func (_IncognitoProxy *IncognitoProxyTransactorSession) SubmitBridgeCandidate(inst []byte, instProofs [2]IncognitoProxyInstructionProof) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.SubmitBridgeCandidate(&_IncognitoProxy.TransactOpts, inst, instProofs)
}

// SubmitFinalityProof is a paid mutator transaction binding the contract method 0xc37cfbf3.
//
// Solidity: function submitFinalityProof((bytes,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])[2] instProofs, uint256 swapID, bool isBeacon) returns()
func (_IncognitoProxy *IncognitoProxyTransactor) SubmitFinalityProof(opts *bind.TransactOpts, instProofs [2]IncognitoProxyInstructionProof, swapID *big.Int, isBeacon bool) (*types.Transaction, error) {
	return _IncognitoProxy.contract.Transact(opts, "submitFinalityProof", instProofs, swapID, isBeacon)
}

// SubmitFinalityProof is a paid mutator transaction binding the contract method 0xc37cfbf3.
//
// Solidity: function submitFinalityProof((bytes,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])[2] instProofs, uint256 swapID, bool isBeacon) returns()
func (_IncognitoProxy *IncognitoProxySession) SubmitFinalityProof(instProofs [2]IncognitoProxyInstructionProof, swapID *big.Int, isBeacon bool) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.SubmitFinalityProof(&_IncognitoProxy.TransactOpts, instProofs, swapID, isBeacon)
}

// SubmitFinalityProof is a paid mutator transaction binding the contract method 0xc37cfbf3.
//
// Solidity: function submitFinalityProof((bytes,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])[2] instProofs, uint256 swapID, bool isBeacon) returns()
func (_IncognitoProxy *IncognitoProxyTransactorSession) SubmitFinalityProof(instProofs [2]IncognitoProxyInstructionProof, swapID *big.Int, isBeacon bool) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.SubmitFinalityProof(&_IncognitoProxy.TransactOpts, instProofs, swapID, isBeacon)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IncognitoProxy *IncognitoProxyTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncognitoProxy.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IncognitoProxy *IncognitoProxySession) Unpause() (*types.Transaction, error) {
	return _IncognitoProxy.Contract.Unpause(&_IncognitoProxy.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IncognitoProxy *IncognitoProxyTransactorSession) Unpause() (*types.Transaction, error) {
	return _IncognitoProxy.Contract.Unpause(&_IncognitoProxy.TransactOpts)
}

// IncognitoProxyBeaconCommitteeSwappedIterator is returned from FilterBeaconCommitteeSwapped and is used to iterate over the raw logs and unpacked data for BeaconCommitteeSwapped events raised by the IncognitoProxy contract.
type IncognitoProxyBeaconCommitteeSwappedIterator struct {
	Event *IncognitoProxyBeaconCommitteeSwapped // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncognitoProxyBeaconCommitteeSwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncognitoProxyBeaconCommitteeSwapped)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncognitoProxyBeaconCommitteeSwapped)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncognitoProxyBeaconCommitteeSwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncognitoProxyBeaconCommitteeSwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncognitoProxyBeaconCommitteeSwapped represents a BeaconCommitteeSwapped event raised by the IncognitoProxy contract.
type IncognitoProxyBeaconCommitteeSwapped struct {
	Id          *big.Int
	StartHeight *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBeaconCommitteeSwapped is a free log retrieval operation binding the contract event 0xe15e1a9dec6ad906dd5985b062bfa5ee8bc5d5738e46e4deb8a2df2fbbbb59d1.
//
// Solidity: event BeaconCommitteeSwapped(uint256 id, uint256 startHeight)
func (_IncognitoProxy *IncognitoProxyFilterer) FilterBeaconCommitteeSwapped(opts *bind.FilterOpts) (*IncognitoProxyBeaconCommitteeSwappedIterator, error) {

	logs, sub, err := _IncognitoProxy.contract.FilterLogs(opts, "BeaconCommitteeSwapped")
	if err != nil {
		return nil, err
	}
	return &IncognitoProxyBeaconCommitteeSwappedIterator{contract: _IncognitoProxy.contract, event: "BeaconCommitteeSwapped", logs: logs, sub: sub}, nil
}

// WatchBeaconCommitteeSwapped is a free log subscription operation binding the contract event 0xe15e1a9dec6ad906dd5985b062bfa5ee8bc5d5738e46e4deb8a2df2fbbbb59d1.
//
// Solidity: event BeaconCommitteeSwapped(uint256 id, uint256 startHeight)
func (_IncognitoProxy *IncognitoProxyFilterer) WatchBeaconCommitteeSwapped(opts *bind.WatchOpts, sink chan<- *IncognitoProxyBeaconCommitteeSwapped) (event.Subscription, error) {

	logs, sub, err := _IncognitoProxy.contract.WatchLogs(opts, "BeaconCommitteeSwapped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncognitoProxyBeaconCommitteeSwapped)
				if err := _IncognitoProxy.contract.UnpackLog(event, "BeaconCommitteeSwapped", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconCommitteeSwapped is a log parse operation binding the contract event 0xe15e1a9dec6ad906dd5985b062bfa5ee8bc5d5738e46e4deb8a2df2fbbbb59d1.
//
// Solidity: event BeaconCommitteeSwapped(uint256 id, uint256 startHeight)
func (_IncognitoProxy *IncognitoProxyFilterer) ParseBeaconCommitteeSwapped(log types.Log) (*IncognitoProxyBeaconCommitteeSwapped, error) {
	event := new(IncognitoProxyBeaconCommitteeSwapped)
	if err := _IncognitoProxy.contract.UnpackLog(event, "BeaconCommitteeSwapped", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncognitoProxyCandidatePromotedIterator is returned from FilterCandidatePromoted and is used to iterate over the raw logs and unpacked data for CandidatePromoted events raised by the IncognitoProxy contract.
type IncognitoProxyCandidatePromotedIterator struct {
	Event *IncognitoProxyCandidatePromoted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncognitoProxyCandidatePromotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncognitoProxyCandidatePromoted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncognitoProxyCandidatePromoted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncognitoProxyCandidatePromotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncognitoProxyCandidatePromotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncognitoProxyCandidatePromoted represents a CandidatePromoted event raised by the IncognitoProxy contract.
type IncognitoProxyCandidatePromoted struct {
	SwapID   *big.Int
	IsBeacon bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCandidatePromoted is a free log retrieval operation binding the contract event 0x40906ca8849ddb84e9284fa5813086f804b1a6e4b236420725c43cbf663b61d5.
//
// Solidity: event CandidatePromoted(uint256 swapID, bool isBeacon)
func (_IncognitoProxy *IncognitoProxyFilterer) FilterCandidatePromoted(opts *bind.FilterOpts) (*IncognitoProxyCandidatePromotedIterator, error) {

	logs, sub, err := _IncognitoProxy.contract.FilterLogs(opts, "CandidatePromoted")
	if err != nil {
		return nil, err
	}
	return &IncognitoProxyCandidatePromotedIterator{contract: _IncognitoProxy.contract, event: "CandidatePromoted", logs: logs, sub: sub}, nil
}

// WatchCandidatePromoted is a free log subscription operation binding the contract event 0x40906ca8849ddb84e9284fa5813086f804b1a6e4b236420725c43cbf663b61d5.
//
// Solidity: event CandidatePromoted(uint256 swapID, bool isBeacon)
func (_IncognitoProxy *IncognitoProxyFilterer) WatchCandidatePromoted(opts *bind.WatchOpts, sink chan<- *IncognitoProxyCandidatePromoted) (event.Subscription, error) {

	logs, sub, err := _IncognitoProxy.contract.WatchLogs(opts, "CandidatePromoted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncognitoProxyCandidatePromoted)
				if err := _IncognitoProxy.contract.UnpackLog(event, "CandidatePromoted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCandidatePromoted is a log parse operation binding the contract event 0x40906ca8849ddb84e9284fa5813086f804b1a6e4b236420725c43cbf663b61d5.
//
// Solidity: event CandidatePromoted(uint256 swapID, bool isBeacon)
func (_IncognitoProxy *IncognitoProxyFilterer) ParseCandidatePromoted(log types.Log) (*IncognitoProxyCandidatePromoted, error) {
	event := new(IncognitoProxyCandidatePromoted)
	if err := _IncognitoProxy.contract.UnpackLog(event, "CandidatePromoted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncognitoProxyChainFinalizedIterator is returned from FilterChainFinalized and is used to iterate over the raw logs and unpacked data for ChainFinalized events raised by the IncognitoProxy contract.
type IncognitoProxyChainFinalizedIterator struct {
	Event *IncognitoProxyChainFinalized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncognitoProxyChainFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncognitoProxyChainFinalized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncognitoProxyChainFinalized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncognitoProxyChainFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncognitoProxyChainFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncognitoProxyChainFinalized represents a ChainFinalized event raised by the IncognitoProxy contract.
type IncognitoProxyChainFinalized struct {
	IsBeacon bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChainFinalized is a free log retrieval operation binding the contract event 0x045c41374537196136848ce62859f8aa22ae91b490103bd2bed75eac6a4d5180.
//
// Solidity: event ChainFinalized(bool isBeacon)
func (_IncognitoProxy *IncognitoProxyFilterer) FilterChainFinalized(opts *bind.FilterOpts) (*IncognitoProxyChainFinalizedIterator, error) {

	logs, sub, err := _IncognitoProxy.contract.FilterLogs(opts, "ChainFinalized")
	if err != nil {
		return nil, err
	}
	return &IncognitoProxyChainFinalizedIterator{contract: _IncognitoProxy.contract, event: "ChainFinalized", logs: logs, sub: sub}, nil
}

// WatchChainFinalized is a free log subscription operation binding the contract event 0x045c41374537196136848ce62859f8aa22ae91b490103bd2bed75eac6a4d5180.
//
// Solidity: event ChainFinalized(bool isBeacon)
func (_IncognitoProxy *IncognitoProxyFilterer) WatchChainFinalized(opts *bind.WatchOpts, sink chan<- *IncognitoProxyChainFinalized) (event.Subscription, error) {

	logs, sub, err := _IncognitoProxy.contract.WatchLogs(opts, "ChainFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncognitoProxyChainFinalized)
				if err := _IncognitoProxy.contract.UnpackLog(event, "ChainFinalized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChainFinalized is a log parse operation binding the contract event 0x045c41374537196136848ce62859f8aa22ae91b490103bd2bed75eac6a4d5180.
//
// Solidity: event ChainFinalized(bool isBeacon)
func (_IncognitoProxy *IncognitoProxyFilterer) ParseChainFinalized(log types.Log) (*IncognitoProxyChainFinalized, error) {
	event := new(IncognitoProxyChainFinalized)
	if err := _IncognitoProxy.contract.UnpackLog(event, "ChainFinalized", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncognitoProxyClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the IncognitoProxy contract.
type IncognitoProxyClaimIterator struct {
	Event *IncognitoProxyClaim // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncognitoProxyClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncognitoProxyClaim)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncognitoProxyClaim)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncognitoProxyClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncognitoProxyClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncognitoProxyClaim represents a Claim event raised by the IncognitoProxy contract.
type IncognitoProxyClaim struct {
	Claimer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_IncognitoProxy *IncognitoProxyFilterer) FilterClaim(opts *bind.FilterOpts) (*IncognitoProxyClaimIterator, error) {

	logs, sub, err := _IncognitoProxy.contract.FilterLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return &IncognitoProxyClaimIterator{contract: _IncognitoProxy.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_IncognitoProxy *IncognitoProxyFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *IncognitoProxyClaim) (event.Subscription, error) {

	logs, sub, err := _IncognitoProxy.contract.WatchLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncognitoProxyClaim)
				if err := _IncognitoProxy.contract.UnpackLog(event, "Claim", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaim is a log parse operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_IncognitoProxy *IncognitoProxyFilterer) ParseClaim(log types.Log) (*IncognitoProxyClaim, error) {
	event := new(IncognitoProxyClaim)
	if err := _IncognitoProxy.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncognitoProxyExtendIterator is returned from FilterExtend and is used to iterate over the raw logs and unpacked data for Extend events raised by the IncognitoProxy contract.
type IncognitoProxyExtendIterator struct {
	Event *IncognitoProxyExtend // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncognitoProxyExtendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncognitoProxyExtend)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncognitoProxyExtend)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncognitoProxyExtendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncognitoProxyExtendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncognitoProxyExtend represents a Extend event raised by the IncognitoProxy contract.
type IncognitoProxyExtend struct {
	Ndays *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterExtend is a free log retrieval operation binding the contract event 0x02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8.
//
// Solidity: event Extend(uint256 ndays)
func (_IncognitoProxy *IncognitoProxyFilterer) FilterExtend(opts *bind.FilterOpts) (*IncognitoProxyExtendIterator, error) {

	logs, sub, err := _IncognitoProxy.contract.FilterLogs(opts, "Extend")
	if err != nil {
		return nil, err
	}
	return &IncognitoProxyExtendIterator{contract: _IncognitoProxy.contract, event: "Extend", logs: logs, sub: sub}, nil
}

// WatchExtend is a free log subscription operation binding the contract event 0x02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8.
//
// Solidity: event Extend(uint256 ndays)
func (_IncognitoProxy *IncognitoProxyFilterer) WatchExtend(opts *bind.WatchOpts, sink chan<- *IncognitoProxyExtend) (event.Subscription, error) {

	logs, sub, err := _IncognitoProxy.contract.WatchLogs(opts, "Extend")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncognitoProxyExtend)
				if err := _IncognitoProxy.contract.UnpackLog(event, "Extend", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExtend is a log parse operation binding the contract event 0x02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8.
//
// Solidity: event Extend(uint256 ndays)
func (_IncognitoProxy *IncognitoProxyFilterer) ParseExtend(log types.Log) (*IncognitoProxyExtend, error) {
	event := new(IncognitoProxyExtend)
	if err := _IncognitoProxy.contract.UnpackLog(event, "Extend", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncognitoProxyPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the IncognitoProxy contract.
type IncognitoProxyPausedIterator struct {
	Event *IncognitoProxyPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncognitoProxyPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncognitoProxyPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncognitoProxyPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncognitoProxyPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncognitoProxyPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncognitoProxyPaused represents a Paused event raised by the IncognitoProxy contract.
type IncognitoProxyPaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address pauser)
func (_IncognitoProxy *IncognitoProxyFilterer) FilterPaused(opts *bind.FilterOpts) (*IncognitoProxyPausedIterator, error) {

	logs, sub, err := _IncognitoProxy.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &IncognitoProxyPausedIterator{contract: _IncognitoProxy.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address pauser)
func (_IncognitoProxy *IncognitoProxyFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *IncognitoProxyPaused) (event.Subscription, error) {

	logs, sub, err := _IncognitoProxy.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncognitoProxyPaused)
				if err := _IncognitoProxy.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address pauser)
func (_IncognitoProxy *IncognitoProxyFilterer) ParsePaused(log types.Log) (*IncognitoProxyPaused, error) {
	event := new(IncognitoProxyPaused)
	if err := _IncognitoProxy.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncognitoProxySubmittedBridgeCandidateIterator is returned from FilterSubmittedBridgeCandidate and is used to iterate over the raw logs and unpacked data for SubmittedBridgeCandidate events raised by the IncognitoProxy contract.
type IncognitoProxySubmittedBridgeCandidateIterator struct {
	Event *IncognitoProxySubmittedBridgeCandidate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncognitoProxySubmittedBridgeCandidateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncognitoProxySubmittedBridgeCandidate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncognitoProxySubmittedBridgeCandidate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncognitoProxySubmittedBridgeCandidateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncognitoProxySubmittedBridgeCandidateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncognitoProxySubmittedBridgeCandidate represents a SubmittedBridgeCandidate event raised by the IncognitoProxy contract.
type IncognitoProxySubmittedBridgeCandidate struct {
	Id          *big.Int
	StartHeight *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSubmittedBridgeCandidate is a free log retrieval operation binding the contract event 0xc918d4d247e9af7d4a8ef8ac6f0078262006c87209d062a6d846f5b547dcff66.
//
// Solidity: event SubmittedBridgeCandidate(uint256 id, uint256 startHeight)
func (_IncognitoProxy *IncognitoProxyFilterer) FilterSubmittedBridgeCandidate(opts *bind.FilterOpts) (*IncognitoProxySubmittedBridgeCandidateIterator, error) {

	logs, sub, err := _IncognitoProxy.contract.FilterLogs(opts, "SubmittedBridgeCandidate")
	if err != nil {
		return nil, err
	}
	return &IncognitoProxySubmittedBridgeCandidateIterator{contract: _IncognitoProxy.contract, event: "SubmittedBridgeCandidate", logs: logs, sub: sub}, nil
}

// WatchSubmittedBridgeCandidate is a free log subscription operation binding the contract event 0xc918d4d247e9af7d4a8ef8ac6f0078262006c87209d062a6d846f5b547dcff66.
//
// Solidity: event SubmittedBridgeCandidate(uint256 id, uint256 startHeight)
func (_IncognitoProxy *IncognitoProxyFilterer) WatchSubmittedBridgeCandidate(opts *bind.WatchOpts, sink chan<- *IncognitoProxySubmittedBridgeCandidate) (event.Subscription, error) {

	logs, sub, err := _IncognitoProxy.contract.WatchLogs(opts, "SubmittedBridgeCandidate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncognitoProxySubmittedBridgeCandidate)
				if err := _IncognitoProxy.contract.UnpackLog(event, "SubmittedBridgeCandidate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSubmittedBridgeCandidate is a log parse operation binding the contract event 0xc918d4d247e9af7d4a8ef8ac6f0078262006c87209d062a6d846f5b547dcff66.
//
// Solidity: event SubmittedBridgeCandidate(uint256 id, uint256 startHeight)
func (_IncognitoProxy *IncognitoProxyFilterer) ParseSubmittedBridgeCandidate(log types.Log) (*IncognitoProxySubmittedBridgeCandidate, error) {
	event := new(IncognitoProxySubmittedBridgeCandidate)
	if err := _IncognitoProxy.contract.UnpackLog(event, "SubmittedBridgeCandidate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncognitoProxyUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the IncognitoProxy contract.
type IncognitoProxyUnpausedIterator struct {
	Event *IncognitoProxyUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncognitoProxyUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncognitoProxyUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncognitoProxyUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncognitoProxyUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncognitoProxyUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncognitoProxyUnpaused represents a Unpaused event raised by the IncognitoProxy contract.
type IncognitoProxyUnpaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address pauser)
func (_IncognitoProxy *IncognitoProxyFilterer) FilterUnpaused(opts *bind.FilterOpts) (*IncognitoProxyUnpausedIterator, error) {

	logs, sub, err := _IncognitoProxy.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &IncognitoProxyUnpausedIterator{contract: _IncognitoProxy.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address pauser)
func (_IncognitoProxy *IncognitoProxyFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *IncognitoProxyUnpaused) (event.Subscription, error) {

	logs, sub, err := _IncognitoProxy.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncognitoProxyUnpaused)
				if err := _IncognitoProxy.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address pauser)
func (_IncognitoProxy *IncognitoProxyFilterer) ParseUnpaused(log types.Log) (*IncognitoProxyUnpaused, error) {
	event := new(IncognitoProxyUnpaused)
	if err := _IncognitoProxy.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}
