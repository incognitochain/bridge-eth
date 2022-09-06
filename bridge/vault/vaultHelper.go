// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vault

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

// VaultHelperPreSignData is an auto generated low-level Go binding around an user-defined struct.
type VaultHelperPreSignData struct {
	Prefix    uint8
	Token     common.Address
	Timestamp []byte
	Amount    *big.Int
}

// VaultHelperMetaData contains all meta data concerning the VaultHelper contract.
var VaultHelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"enumVaultHelper.Prefix\",\"name\":\"prefix\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"timestamp\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"_buildPreSignData\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumVaultHelper.Prefix\",\"name\":\"prefix\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"timestamp\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structVaultHelper.PreSignData\",\"name\":\"psd\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"recipientToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"exchangeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"name\":\"_buildSignExecute\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumVaultHelper.Prefix\",\"name\":\"prefix\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"timestamp\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structVaultHelper.PreSignData\",\"name\":\"psd\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"}],\"name\":\"_buildSignRequestWithdraw\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_senser\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_txId\",\"type\":\"bytes32\"}],\"name\":\"_buildSignShield\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506109b8806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80633cf6e1ca1461005157806341f59b6f14610081578063741d7a50146100b1578063e24ef4bd146100cd575b600080fd5b61006b600480360381019061006691906104e7565b6100fd565b6040516100789190610700565b60405180910390f35b61009b6004803603810190610096919061044f565b610131565b6040516100a89190610700565b60405180910390f35b6100cb60048036038101906100c691906103cf565b61016b565b005b6100e760048036038101906100e29190610393565b610172565b6040516100f49190610700565b60405180910390f35b60608084848460405160200161011593929190610777565b6040516020818303038152906040529050809150509392505050565b606080868686868660405160200161014d959493929190610722565b60405160208183030381529060405290508091505095945050505050565b5050505050565b60608060405180604001604052808573ffffffffffffffffffffffffffffffffffffffff168152602001848152506040516020016101b091906107b0565b60405160208183030381529060405290508091505092915050565b6000813590506101da8161092d565b92915050565b6000813590506101ef81610944565b92915050565b60008083601f84011261020757600080fd5b8235905067ffffffffffffffff81111561022057600080fd5b60208301915083600182028301111561023857600080fd5b9250929050565b600082601f83011261025057600080fd5b813561026361025e826107f8565b6107cb565b9150808252602083016020830185838301111561027f57600080fd5b61028a8382846108cd565b50505092915050565b6000813590506102a28161095b565b92915050565b60008083601f8401126102ba57600080fd5b8235905067ffffffffffffffff8111156102d357600080fd5b6020830191508360018202830111156102eb57600080fd5b9250929050565b60006080828403121561030457600080fd5b61030e60806107cb565b9050600061031e84828501610293565b6000830152506020610332848285016101cb565b602083015250604082013567ffffffffffffffff81111561035257600080fd5b61035e8482850161023f565b60408301525060606103728482850161037e565b60608301525092915050565b60008135905061038d8161096b565b92915050565b600080604083850312156103a657600080fd5b60006103b4858286016101cb565b92505060206103c5858286016101e0565b9150509250929050565b6000806000806000608086880312156103e757600080fd5b60006103f588828901610293565b9550506020610406888289016101cb565b945050604086013567ffffffffffffffff81111561042357600080fd5b61042f888289016101f5565b935093505060606104428882890161037e565b9150509295509295909350565b60008060008060006080868803121561046757600080fd5b600086013567ffffffffffffffff81111561048157600080fd5b61048d888289016102f2565b955050602061049e888289016101cb565b94505060406104af888289016101cb565b935050606086013567ffffffffffffffff8111156104cc57600080fd5b6104d8888289016101f5565b92509250509295509295909350565b6000806000604084860312156104fc57600080fd5b600084013567ffffffffffffffff81111561051657600080fd5b610522868287016102f2565b935050602084013567ffffffffffffffff81111561053f57600080fd5b61054b868287016102a8565b92509250509250925092565b61056081610862565b82525050565b61056f81610862565b82525050565b61057e81610874565b82525050565b60006105908385610840565b935061059d8385846108cd565b6105a68361090f565b840190509392505050565b60006105bc82610824565b6105c6818561082f565b93506105d68185602086016108dc565b6105df8161090f565b840191505092915050565b60006105f582610824565b6105ff8185610840565b935061060f8185602086016108dc565b6106188161090f565b840191505092915050565b61062c816108bb565b82525050565b600061063e8385610851565b935061064b8385846108cd565b6106548361090f565b840190509392505050565b60006080830160008301516106776000860182610623565b50602083015161068a6020860182610557565b50604083015184820360408601526106a282826105b1565b91505060608301516106b760608601826106f1565b508091505092915050565b6040820160008201516106d86000850182610557565b5060208201516106eb6020850182610575565b50505050565b6106fa816108b1565b82525050565b6000602082019050818103600083015261071a81846105ea565b905092915050565b6000608082019050818103600083015261073c818861065f565b905061074b6020830187610566565b6107586040830186610566565b818103606083015261076b818486610584565b90509695505050505050565b60006040820190508181036000830152610791818661065f565b905081810360208301526107a6818486610632565b9050949350505050565b60006040820190506107c560008301846106c2565b92915050565b6000604051905081810181811067ffffffffffffffff821117156107ee57600080fd5b8060405250919050565b600067ffffffffffffffff82111561080f57600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600061086d82610891565b9050919050565b6000819050919050565b600081905061088c82610920565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006108c68261087e565b9050919050565b82818337600083830152505050565b60005b838110156108fa5780820151818401526020810190506108df565b83811115610909576000848401525b50505050565b6000601f19601f8301169050919050565b6002811061092a57fe5b50565b61093681610862565b811461094157600080fd5b50565b61094d81610874565b811461095857600080fd5b50565b6002811061096857600080fd5b50565b610974816108b1565b811461097f57600080fd5b5056fea264697066735822122020b592d9f97c24bd806d1717bf1e687a827df2dfe5b51cf7a2e268bf41760a5464736f6c634300060c0033",
}

// VaultHelperABI is the input ABI used to generate the binding from.
// Deprecated: Use VaultHelperMetaData.ABI instead.
var VaultHelperABI = VaultHelperMetaData.ABI

// VaultHelperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VaultHelperMetaData.Bin instead.
var VaultHelperBin = VaultHelperMetaData.Bin

// DeployVaultHelper deploys a new Ethereum contract, binding an instance of VaultHelper to it.
func DeployVaultHelper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VaultHelper, error) {
	parsed, err := VaultHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VaultHelperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VaultHelper{VaultHelperCaller: VaultHelperCaller{contract: contract}, VaultHelperTransactor: VaultHelperTransactor{contract: contract}, VaultHelperFilterer: VaultHelperFilterer{contract: contract}}, nil
}

// VaultHelper is an auto generated Go binding around an Ethereum contract.
type VaultHelper struct {
	VaultHelperCaller     // Read-only binding to the contract
	VaultHelperTransactor // Write-only binding to the contract
	VaultHelperFilterer   // Log filterer for contract events
}

// VaultHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultHelperSession struct {
	Contract     *VaultHelper      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultHelperCallerSession struct {
	Contract *VaultHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// VaultHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultHelperTransactorSession struct {
	Contract     *VaultHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// VaultHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultHelperRaw struct {
	Contract *VaultHelper // Generic contract binding to access the raw methods on
}

// VaultHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultHelperCallerRaw struct {
	Contract *VaultHelperCaller // Generic read-only contract binding to access the raw methods on
}

// VaultHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultHelperTransactorRaw struct {
	Contract *VaultHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVaultHelper creates a new instance of VaultHelper, bound to a specific deployed contract.
func NewVaultHelper(address common.Address, backend bind.ContractBackend) (*VaultHelper, error) {
	contract, err := bindVaultHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VaultHelper{VaultHelperCaller: VaultHelperCaller{contract: contract}, VaultHelperTransactor: VaultHelperTransactor{contract: contract}, VaultHelperFilterer: VaultHelperFilterer{contract: contract}}, nil
}

// NewVaultHelperCaller creates a new read-only instance of VaultHelper, bound to a specific deployed contract.
func NewVaultHelperCaller(address common.Address, caller bind.ContractCaller) (*VaultHelperCaller, error) {
	contract, err := bindVaultHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultHelperCaller{contract: contract}, nil
}

// NewVaultHelperTransactor creates a new write-only instance of VaultHelper, bound to a specific deployed contract.
func NewVaultHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultHelperTransactor, error) {
	contract, err := bindVaultHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultHelperTransactor{contract: contract}, nil
}

// NewVaultHelperFilterer creates a new log filterer instance of VaultHelper, bound to a specific deployed contract.
func NewVaultHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultHelperFilterer, error) {
	contract, err := bindVaultHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultHelperFilterer{contract: contract}, nil
}

// bindVaultHelper binds a generic wrapper to an already deployed contract.
func bindVaultHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultHelperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultHelper *VaultHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultHelper.Contract.VaultHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultHelper *VaultHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultHelper.Contract.VaultHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultHelper *VaultHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultHelper.Contract.VaultHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultHelper *VaultHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultHelper *VaultHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultHelper *VaultHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultHelper.Contract.contract.Transact(opts, method, params...)
}

// BuildPreSignData is a free data retrieval call binding the contract method 0x741d7a50.
//
// Solidity: function _buildPreSignData(uint8 prefix, address token, bytes timestamp, uint256 amount) pure returns()
func (_VaultHelper *VaultHelperCaller) BuildPreSignData(opts *bind.CallOpts, prefix uint8, token common.Address, timestamp []byte, amount *big.Int) error {
	var out []interface{}
	err := _VaultHelper.contract.Call(opts, &out, "_buildPreSignData", prefix, token, timestamp, amount)

	if err != nil {
		return err
	}

	return err

}

// BuildPreSignData is a free data retrieval call binding the contract method 0x741d7a50.
//
// Solidity: function _buildPreSignData(uint8 prefix, address token, bytes timestamp, uint256 amount) pure returns()
func (_VaultHelper *VaultHelperSession) BuildPreSignData(prefix uint8, token common.Address, timestamp []byte, amount *big.Int) error {
	return _VaultHelper.Contract.BuildPreSignData(&_VaultHelper.CallOpts, prefix, token, timestamp, amount)
}

// BuildPreSignData is a free data retrieval call binding the contract method 0x741d7a50.
//
// Solidity: function _buildPreSignData(uint8 prefix, address token, bytes timestamp, uint256 amount) pure returns()
func (_VaultHelper *VaultHelperCallerSession) BuildPreSignData(prefix uint8, token common.Address, timestamp []byte, amount *big.Int) error {
	return _VaultHelper.Contract.BuildPreSignData(&_VaultHelper.CallOpts, prefix, token, timestamp, amount)
}

// BuildSignExecute is a free data retrieval call binding the contract method 0x41f59b6f.
//
// Solidity: function _buildSignExecute((uint8,address,bytes,uint256) psd, address recipientToken, address exchangeAddress, bytes callData) pure returns(bytes)
func (_VaultHelper *VaultHelperCaller) BuildSignExecute(opts *bind.CallOpts, psd VaultHelperPreSignData, recipientToken common.Address, exchangeAddress common.Address, callData []byte) ([]byte, error) {
	var out []interface{}
	err := _VaultHelper.contract.Call(opts, &out, "_buildSignExecute", psd, recipientToken, exchangeAddress, callData)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// BuildSignExecute is a free data retrieval call binding the contract method 0x41f59b6f.
//
// Solidity: function _buildSignExecute((uint8,address,bytes,uint256) psd, address recipientToken, address exchangeAddress, bytes callData) pure returns(bytes)
func (_VaultHelper *VaultHelperSession) BuildSignExecute(psd VaultHelperPreSignData, recipientToken common.Address, exchangeAddress common.Address, callData []byte) ([]byte, error) {
	return _VaultHelper.Contract.BuildSignExecute(&_VaultHelper.CallOpts, psd, recipientToken, exchangeAddress, callData)
}

// BuildSignExecute is a free data retrieval call binding the contract method 0x41f59b6f.
//
// Solidity: function _buildSignExecute((uint8,address,bytes,uint256) psd, address recipientToken, address exchangeAddress, bytes callData) pure returns(bytes)
func (_VaultHelper *VaultHelperCallerSession) BuildSignExecute(psd VaultHelperPreSignData, recipientToken common.Address, exchangeAddress common.Address, callData []byte) ([]byte, error) {
	return _VaultHelper.Contract.BuildSignExecute(&_VaultHelper.CallOpts, psd, recipientToken, exchangeAddress, callData)
}

// BuildSignRequestWithdraw is a free data retrieval call binding the contract method 0x3cf6e1ca.
//
// Solidity: function _buildSignRequestWithdraw((uint8,address,bytes,uint256) psd, string incognitoAddress) pure returns(bytes)
func (_VaultHelper *VaultHelperCaller) BuildSignRequestWithdraw(opts *bind.CallOpts, psd VaultHelperPreSignData, incognitoAddress string) ([]byte, error) {
	var out []interface{}
	err := _VaultHelper.contract.Call(opts, &out, "_buildSignRequestWithdraw", psd, incognitoAddress)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// BuildSignRequestWithdraw is a free data retrieval call binding the contract method 0x3cf6e1ca.
//
// Solidity: function _buildSignRequestWithdraw((uint8,address,bytes,uint256) psd, string incognitoAddress) pure returns(bytes)
func (_VaultHelper *VaultHelperSession) BuildSignRequestWithdraw(psd VaultHelperPreSignData, incognitoAddress string) ([]byte, error) {
	return _VaultHelper.Contract.BuildSignRequestWithdraw(&_VaultHelper.CallOpts, psd, incognitoAddress)
}

// BuildSignRequestWithdraw is a free data retrieval call binding the contract method 0x3cf6e1ca.
//
// Solidity: function _buildSignRequestWithdraw((uint8,address,bytes,uint256) psd, string incognitoAddress) pure returns(bytes)
func (_VaultHelper *VaultHelperCallerSession) BuildSignRequestWithdraw(psd VaultHelperPreSignData, incognitoAddress string) ([]byte, error) {
	return _VaultHelper.Contract.BuildSignRequestWithdraw(&_VaultHelper.CallOpts, psd, incognitoAddress)
}

// BuildSignShield is a free data retrieval call binding the contract method 0xe24ef4bd.
//
// Solidity: function _buildSignShield(address _senser, bytes32 _txId) pure returns(bytes)
func (_VaultHelper *VaultHelperCaller) BuildSignShield(opts *bind.CallOpts, _senser common.Address, _txId [32]byte) ([]byte, error) {
	var out []interface{}
	err := _VaultHelper.contract.Call(opts, &out, "_buildSignShield", _senser, _txId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// BuildSignShield is a free data retrieval call binding the contract method 0xe24ef4bd.
//
// Solidity: function _buildSignShield(address _senser, bytes32 _txId) pure returns(bytes)
func (_VaultHelper *VaultHelperSession) BuildSignShield(_senser common.Address, _txId [32]byte) ([]byte, error) {
	return _VaultHelper.Contract.BuildSignShield(&_VaultHelper.CallOpts, _senser, _txId)
}

// BuildSignShield is a free data retrieval call binding the contract method 0xe24ef4bd.
//
// Solidity: function _buildSignShield(address _senser, bytes32 _txId) pure returns(bytes)
func (_VaultHelper *VaultHelperCallerSession) BuildSignShield(_senser common.Address, _txId [32]byte) ([]byte, error) {
	return _VaultHelper.Contract.BuildSignShield(&_VaultHelper.CallOpts, _senser, _txId)
}
