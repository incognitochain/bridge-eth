// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package governance

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

// GovernanceHelperMetaData contains all meta data concerning the GovernanceHelper contract.
var GovernanceHelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PROPOSAL_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"_buildSignProposal\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"_buildSignProposalEncodeAbi\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600054610100900460ff16158080156100315750600054600160ff909116105b8061005c575061004a3061018060201b6101691760201c565b15801561005c575060005460ff166001145b6100c45760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff1916600117905580156100e7576000805461ff0019166101001790555b6101346040518060400160405280600c81526020016b496e636f676e69746f44414f60a01b815250604051806040016040528060018152602001603160f81b81525061018f60201b60201c565b801561017a576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50610269565b6001600160a01b03163b151590565b600054610100900460ff166101e85760405162461bcd60e51b815260206004820152602b602482015260008051602061095e83398151915260448201526a6e697469616c697a696e6760a81b60648201526084016100bb565b6101f282826101f6565b5050565b600054610100900460ff1661024f5760405162461bcd60e51b815260206004820152602b602482015260008051602061095e83398151915260448201526a6e697469616c697a696e6760a81b60648201526084016100bb565b815160209283012081519190920120600191909155600255565b6106e6806102786000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063bdc9b04c14610046578063d2197bac14610080578063f5bf899f146100a0575b600080fd5b61006d7fb6d2dc83590271a7c0a5ab5fbf6a2dad418bbfd533c253e3d69a6772712809c781565b6040519081526020015b60405180910390f35b61009361008e366004610436565b6100b3565b6040516100779190610594565b61006d6100ae366004610436565b610107565b60607fb6d2dc83590271a7c0a5ab5fbf6a2dad418bbfd533c253e3d69a6772712809c7858585856040516020016100ee9594939291906105f9565b6040516020818303038152906040529050949350505050565b60006101607fb6d2dc83590271a7c0a5ab5fbf6a2dad418bbfd533c253e3d69a6772712809c7868686866040516020016101459594939291906105f9565b60405160208183030381529060405280519060200120610178565b95945050505050565b6001600160a01b03163b151590565b60006101c66101856101cc565b8360405161190160f01b6020820152602281018390526042810182905260009060620160405160208183030381529060405280519060200120905092915050565b92915050565b60006102477f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6101fb60015490565b6002546040805160208101859052908101839052606081018290524660808201523060a082015260009060c0016040516020818303038152906040528051906020012090509392505050565b905090565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561028b5761028b61024c565b604052919050565b600067ffffffffffffffff8211156102ad576102ad61024c565b5060051b60200190565b600082601f8301126102c857600080fd5b813560206102dd6102d883610293565b610262565b82815260059290921b840181019181810190868411156102fc57600080fd5b8286015b848110156103175780358352918301918301610300565b509695505050505050565b600067ffffffffffffffff83111561033c5761033c61024c565b61034f601f8401601f1916602001610262565b905082815283838301111561036357600080fd5b828260208301376000602084830101529392505050565b600082601f83011261038b57600080fd5b8135602061039b6102d883610293565b82815260059290921b840181019181810190868411156103ba57600080fd5b8286015b8481101561031757803567ffffffffffffffff8111156103de5760008081fd5b8701603f810189136103f05760008081fd5b610401898683013560408401610322565b8452509183019183016103be565b600082601f83011261042057600080fd5b61042f83833560208501610322565b9392505050565b6000806000806080858703121561044c57600080fd5b843567ffffffffffffffff8082111561046457600080fd5b818701915087601f83011261047857600080fd5b813560206104886102d883610293565b82815260059290921b8401810191818101908b8411156104a757600080fd5b948201945b838610156104db5785356001600160a01b03811681146104cc5760008081fd5b825294820194908201906104ac565b985050880135925050808211156104f157600080fd5b6104fd888389016102b7565b9450604087013591508082111561051357600080fd5b61051f8883890161037a565b9350606087013591508082111561053557600080fd5b506105428782880161040f565b91505092959194509250565b6000815180845260005b8181101561057457602081850181015186830182015201610558565b506000602082860101526020601f19601f83011685010191505092915050565b60208152600061042f602083018461054e565b6000815180845260208085019450848260051b860182860160005b858110156105ec5783830389526105da83835161054e565b988501989250908401906001016105c2565b5090979650505050505050565b600060a08201878352602060a08185015281885180845260c086019150828a01935060005b818110156106435784516001600160a01b03168352938301939183019160010161061e565b50508481036040860152875180825290820192508188019060005b8181101561067a5782518552938301939183019160010161065e565b50505050828103606084015261069081866105a7565b905082810360808401526106a4818561054e565b9897505050505050505056fea2646970667358221220caa02b5a035981209fe0b7b3456f15613fcc80734d2d70a221e58d983c86b97564736f6c63430008110033496e697469616c697a61626c653a20636f6e7472616374206973206e6f742069",
}

// GovernanceHelperABI is the input ABI used to generate the binding from.
// Deprecated: Use GovernanceHelperMetaData.ABI instead.
var GovernanceHelperABI = GovernanceHelperMetaData.ABI

// GovernanceHelperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GovernanceHelperMetaData.Bin instead.
var GovernanceHelperBin = GovernanceHelperMetaData.Bin

// DeployGovernanceHelper deploys a new Ethereum contract, binding an instance of GovernanceHelper to it.
func DeployGovernanceHelper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GovernanceHelper, error) {
	parsed, err := GovernanceHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GovernanceHelperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GovernanceHelper{GovernanceHelperCaller: GovernanceHelperCaller{contract: contract}, GovernanceHelperTransactor: GovernanceHelperTransactor{contract: contract}, GovernanceHelperFilterer: GovernanceHelperFilterer{contract: contract}}, nil
}

// GovernanceHelper is an auto generated Go binding around an Ethereum contract.
type GovernanceHelper struct {
	GovernanceHelperCaller     // Read-only binding to the contract
	GovernanceHelperTransactor // Write-only binding to the contract
	GovernanceHelperFilterer   // Log filterer for contract events
}

// GovernanceHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovernanceHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovernanceHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovernanceHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovernanceHelperSession struct {
	Contract     *GovernanceHelper // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovernanceHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovernanceHelperCallerSession struct {
	Contract *GovernanceHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// GovernanceHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovernanceHelperTransactorSession struct {
	Contract     *GovernanceHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// GovernanceHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovernanceHelperRaw struct {
	Contract *GovernanceHelper // Generic contract binding to access the raw methods on
}

// GovernanceHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovernanceHelperCallerRaw struct {
	Contract *GovernanceHelperCaller // Generic read-only contract binding to access the raw methods on
}

// GovernanceHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovernanceHelperTransactorRaw struct {
	Contract *GovernanceHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGovernanceHelper creates a new instance of GovernanceHelper, bound to a specific deployed contract.
func NewGovernanceHelper(address common.Address, backend bind.ContractBackend) (*GovernanceHelper, error) {
	contract, err := bindGovernanceHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GovernanceHelper{GovernanceHelperCaller: GovernanceHelperCaller{contract: contract}, GovernanceHelperTransactor: GovernanceHelperTransactor{contract: contract}, GovernanceHelperFilterer: GovernanceHelperFilterer{contract: contract}}, nil
}

// NewGovernanceHelperCaller creates a new read-only instance of GovernanceHelper, bound to a specific deployed contract.
func NewGovernanceHelperCaller(address common.Address, caller bind.ContractCaller) (*GovernanceHelperCaller, error) {
	contract, err := bindGovernanceHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceHelperCaller{contract: contract}, nil
}

// NewGovernanceHelperTransactor creates a new write-only instance of GovernanceHelper, bound to a specific deployed contract.
func NewGovernanceHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*GovernanceHelperTransactor, error) {
	contract, err := bindGovernanceHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceHelperTransactor{contract: contract}, nil
}

// NewGovernanceHelperFilterer creates a new log filterer instance of GovernanceHelper, bound to a specific deployed contract.
func NewGovernanceHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*GovernanceHelperFilterer, error) {
	contract, err := bindGovernanceHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovernanceHelperFilterer{contract: contract}, nil
}

// bindGovernanceHelper binds a generic wrapper to an already deployed contract.
func bindGovernanceHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernanceHelperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovernanceHelper *GovernanceHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GovernanceHelper.Contract.GovernanceHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovernanceHelper *GovernanceHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernanceHelper.Contract.GovernanceHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovernanceHelper *GovernanceHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovernanceHelper.Contract.GovernanceHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovernanceHelper *GovernanceHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GovernanceHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovernanceHelper *GovernanceHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernanceHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovernanceHelper *GovernanceHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovernanceHelper.Contract.contract.Transact(opts, method, params...)
}

// PROPOSALHASH is a free data retrieval call binding the contract method 0xbdc9b04c.
//
// Solidity: function PROPOSAL_HASH() view returns(bytes32)
func (_GovernanceHelper *GovernanceHelperCaller) PROPOSALHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovernanceHelper.contract.Call(opts, &out, "PROPOSAL_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PROPOSALHASH is a free data retrieval call binding the contract method 0xbdc9b04c.
//
// Solidity: function PROPOSAL_HASH() view returns(bytes32)
func (_GovernanceHelper *GovernanceHelperSession) PROPOSALHASH() ([32]byte, error) {
	return _GovernanceHelper.Contract.PROPOSALHASH(&_GovernanceHelper.CallOpts)
}

// PROPOSALHASH is a free data retrieval call binding the contract method 0xbdc9b04c.
//
// Solidity: function PROPOSAL_HASH() view returns(bytes32)
func (_GovernanceHelper *GovernanceHelperCallerSession) PROPOSALHASH() ([32]byte, error) {
	return _GovernanceHelper.Contract.PROPOSALHASH(&_GovernanceHelper.CallOpts)
}

// BuildSignProposal is a free data retrieval call binding the contract method 0xf5bf899f.
//
// Solidity: function _buildSignProposal(address[] targets, uint256[] values, bytes[] calldatas, string description) view returns(bytes32)
func (_GovernanceHelper *GovernanceHelperCaller) BuildSignProposal(opts *bind.CallOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, description string) ([32]byte, error) {
	var out []interface{}
	err := _GovernanceHelper.contract.Call(opts, &out, "_buildSignProposal", targets, values, calldatas, description)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BuildSignProposal is a free data retrieval call binding the contract method 0xf5bf899f.
//
// Solidity: function _buildSignProposal(address[] targets, uint256[] values, bytes[] calldatas, string description) view returns(bytes32)
func (_GovernanceHelper *GovernanceHelperSession) BuildSignProposal(targets []common.Address, values []*big.Int, calldatas [][]byte, description string) ([32]byte, error) {
	return _GovernanceHelper.Contract.BuildSignProposal(&_GovernanceHelper.CallOpts, targets, values, calldatas, description)
}

// BuildSignProposal is a free data retrieval call binding the contract method 0xf5bf899f.
//
// Solidity: function _buildSignProposal(address[] targets, uint256[] values, bytes[] calldatas, string description) view returns(bytes32)
func (_GovernanceHelper *GovernanceHelperCallerSession) BuildSignProposal(targets []common.Address, values []*big.Int, calldatas [][]byte, description string) ([32]byte, error) {
	return _GovernanceHelper.Contract.BuildSignProposal(&_GovernanceHelper.CallOpts, targets, values, calldatas, description)
}

// BuildSignProposalEncodeAbi is a free data retrieval call binding the contract method 0xd2197bac.
//
// Solidity: function _buildSignProposalEncodeAbi(address[] targets, uint256[] values, bytes[] calldatas, string description) pure returns(bytes)
func (_GovernanceHelper *GovernanceHelperCaller) BuildSignProposalEncodeAbi(opts *bind.CallOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, description string) ([]byte, error) {
	var out []interface{}
	err := _GovernanceHelper.contract.Call(opts, &out, "_buildSignProposalEncodeAbi", targets, values, calldatas, description)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// BuildSignProposalEncodeAbi is a free data retrieval call binding the contract method 0xd2197bac.
//
// Solidity: function _buildSignProposalEncodeAbi(address[] targets, uint256[] values, bytes[] calldatas, string description) pure returns(bytes)
func (_GovernanceHelper *GovernanceHelperSession) BuildSignProposalEncodeAbi(targets []common.Address, values []*big.Int, calldatas [][]byte, description string) ([]byte, error) {
	return _GovernanceHelper.Contract.BuildSignProposalEncodeAbi(&_GovernanceHelper.CallOpts, targets, values, calldatas, description)
}

// BuildSignProposalEncodeAbi is a free data retrieval call binding the contract method 0xd2197bac.
//
// Solidity: function _buildSignProposalEncodeAbi(address[] targets, uint256[] values, bytes[] calldatas, string description) pure returns(bytes)
func (_GovernanceHelper *GovernanceHelperCallerSession) BuildSignProposalEncodeAbi(targets []common.Address, values []*big.Int, calldatas [][]byte, description string) ([]byte, error) {
	return _GovernanceHelper.Contract.BuildSignProposalEncodeAbi(&_GovernanceHelper.CallOpts, targets, values, calldatas, description)
}

// GovernanceHelperInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the GovernanceHelper contract.
type GovernanceHelperInitializedIterator struct {
	Event *GovernanceHelperInitialized // Event containing the contract specifics and raw log

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
func (it *GovernanceHelperInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceHelperInitialized)
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
		it.Event = new(GovernanceHelperInitialized)
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
func (it *GovernanceHelperInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceHelperInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceHelperInitialized represents a Initialized event raised by the GovernanceHelper contract.
type GovernanceHelperInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GovernanceHelper *GovernanceHelperFilterer) FilterInitialized(opts *bind.FilterOpts) (*GovernanceHelperInitializedIterator, error) {

	logs, sub, err := _GovernanceHelper.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GovernanceHelperInitializedIterator{contract: _GovernanceHelper.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GovernanceHelper *GovernanceHelperFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GovernanceHelperInitialized) (event.Subscription, error) {

	logs, sub, err := _GovernanceHelper.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceHelperInitialized)
				if err := _GovernanceHelper.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GovernanceHelper *GovernanceHelperFilterer) ParseInitialized(log types.Log) (*GovernanceHelperInitialized, error) {
	event := new(GovernanceHelperInitialized)
	if err := _GovernanceHelper.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
