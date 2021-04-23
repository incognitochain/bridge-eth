// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package usdc

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

// UsdcABI is the input ABI used to generate the binding from.
const UsdcABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"unBlacklist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"removeMinter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_currency\",\"type\":\"string\"},{\"name\":\"_decimals\",\"type\":\"uint8\"},{\"name\":\"_masterMinter\",\"type\":\"address\"},{\"name\":\"_pauser\",\"type\":\"address\"},{\"name\":\"_blacklister\",\"type\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"masterMinter\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minter\",\"type\":\"address\"},{\"name\":\"minterAllowedAmount\",\"type\":\"uint256\"}],\"name\":\"configureMinter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newPauser\",\"type\":\"address\"}],\"name\":\"updatePauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"minterAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pauser\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newMasterMinter\",\"type\":\"address\"}],\"name\":\"updateMasterMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newBlacklister\",\"type\":\"address\"}],\"name\":\"updateBlacklister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"blacklister\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currency\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"blacklist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isBlacklisted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"minterAllowedAmount\",\"type\":\"uint256\"}],\"name\":\"MinterConfigured\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldMinter\",\"type\":\"address\"}],\"name\":\"MinterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"newMasterMinter\",\"type\":\"address\"}],\"name\":\"MasterMinterChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"Blacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"UnBlacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"newBlacklister\",\"type\":\"address\"}],\"name\":\"BlacklisterChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"PauserChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// UsdcBin is the compiled bytecode used for deploying new contracts.
var UsdcBin = "0x60806040526000600160146101000a81548160ff0219169083151502179055506000600b5561003c33610041640100000000026401000000009004565b610084565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b612f5b806100936000396000f30060806040526004361061018b576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306fdde0314610190578063095ea7b31461022057806318160ddd146102855780631a895266146102b057806323b872dd146102f35780633092afd514610378578063313ce567146103d35780633357162b1461040457806335d99f35146105865780633f4ba83a146105dd57806340c10f19146105f457806342966c68146106595780634e44d95614610686578063554bab3c146106eb5780635c975abb1461072e57806370a082311461075d5780638456cb59146107b45780638a6db9c3146107cb5780638da5cb5b1461082257806395d89b41146108795780639fd0506d14610909578063a9059cbb14610960578063aa20e1e4146109c5578063aa271e1a14610a08578063ad38bf2214610a63578063bd10243014610aa6578063dd62ed3e14610afd578063e5a6b10f14610b74578063f2fde38b14610c04578063f9f92be414610c47578063fe575a8714610c8a575b600080fd5b34801561019c57600080fd5b506101a5610ce5565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101e55780820151818401526020810190506101ca565b50505050905090810190601f1680156102125780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561022c57600080fd5b5061026b600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610d83565b604051808215151515815260200191505060405180910390f35b34801561029157600080fd5b5061029a610f53565b6040518082815260200191505060405180910390f35b3480156102bc57600080fd5b506102f1600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610f5d565b005b3480156102ff57600080fd5b5061035e600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611057565b604051808215151515815260200191505060405180910390f35b34801561038457600080fd5b506103b9600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611556565b604051808215151515815260200191505060405180910390f35b3480156103df57600080fd5b506103e861169d565b604051808260ff1660ff16815260200191505060405180910390f35b34801561041057600080fd5b50610584600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803560ff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506116b0565b005b34801561059257600080fd5b5061059b61190d565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156105e957600080fd5b506105f2611933565b005b34801561060057600080fd5b5061063f600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506119d8565b604051808215151515815260200191505060405180910390f35b34801561066557600080fd5b5061068460048036038101908080359060200190929190505050611d7a565b005b34801561069257600080fd5b506106d1600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611fe1565b604051808215151515815260200191505060405180910390f35b3480156106f757600080fd5b5061072c600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061214f565b005b34801561073a57600080fd5b50610743612275565b604051808215151515815260200191505060405180910390f35b34801561076957600080fd5b5061079e600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612288565b6040518082815260200191505060405180910390f35b3480156107c057600080fd5b506107c96122d1565b005b3480156107d757600080fd5b5061080c600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612375565b6040518082815260200191505060405180910390f35b34801561082e57600080fd5b506108376123be565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561088557600080fd5b5061088e6123e7565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156108ce5780820151818401526020810190506108b3565b50505050905090810190601f1680156108fb5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561091557600080fd5b5061091e612485565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561096c57600080fd5b506109ab600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506124ab565b604051808215151515815260200191505060405180910390f35b3480156109d157600080fd5b50610a06600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506127ae565b005b348015610a1457600080fd5b50610a49600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506128d4565b604051808215151515815260200191505060405180910390f35b348015610a6f57600080fd5b50610aa4600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061292a565b005b348015610ab257600080fd5b50610abb612a50565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610b0957600080fd5b50610b5e600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612a76565b6040518082815260200191505060405180910390f35b348015610b8057600080fd5b50610b89612afd565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610bc9578082015181840152602081019050610bae565b50505050905090810190601f168015610bf65780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b348015610c1057600080fd5b50610c45600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612b9b565b005b348015610c5357600080fd5b50610c88600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612cc2565b005b348015610c9657600080fd5b50610ccb600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612dbc565b604051808215151515815260200191505060405180910390f35b60048054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610d7b5780601f10610d5057610100808354040283529160200191610d7b565b820191906000526020600020905b815481529060010190602001808311610d5e57829003601f168201915b505050505081565b6000600160149054906101000a900460ff16151515610da157600080fd5b3360001515600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515141515610e0157600080fd5b8360001515600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515141515610e6157600080fd5b83600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925866040518082815260200191505060405180910390a360019250505092915050565b6000600b54905090565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610fb957600080fd5b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508073ffffffffffffffffffffffffffffffffffffffff167f117e3210bb9aa7d9baff172026820255c6f6c30ba8999d1c2fd88e2848137c4e60405160405180910390a250565b6000600160149054906101000a900460ff1615151561107557600080fd5b8260001515600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615151415156110d557600080fd5b3360001515600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151514151561113557600080fd5b8560001515600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151514151561119557600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16141515156111d157600080fd5b600960008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054851115151561121f57600080fd5b600a60008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205485111515156112aa57600080fd5b6112fc85600960008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054612e1290919063ffffffff16565b600960008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061139185600960008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054612e2b90919063ffffffff16565b600960008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061146385600a60008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054612e1290919063ffffffff16565b600a60008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef876040518082815260200191505060405180910390a3600193505050509392505050565b6000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156115b457600080fd5b6000600c60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506000600d60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff167fe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb6669260405160405180910390a260019050919050565b600660009054906101000a900460ff1681565b600860149054906101000a900460ff161515156116cc57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415151561170857600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415151561174457600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415151561178057600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156117bc57600080fd5b87600490805190602001906117d2929190612e8a565b5086600590805190602001906117e9929190612e8a565b508560079080519060200190611800929190612e8a565b5084600660006101000a81548160ff021916908360ff16021790555083600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506118e881612e47565b6001600860146101000a81548160ff0219169083151502179055505050505050505050565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561198f57600080fd5b6000600160146101000a81548160ff0219169083151502179055507f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b600080600160149054906101000a900460ff161515156119f757600080fd5b60011515600c60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515141515611a5657600080fd5b3360001515600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515141515611ab657600080fd5b8460001515600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515141515611b1657600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1614151515611b5257600080fd5b600085111515611b6157600080fd5b600d60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549250828511151515611bb257600080fd5b611bc785600b54612e2b90919063ffffffff16565b600b81905550611c1f85600960008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054612e2b90919063ffffffff16565b600960008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550611c758584612e1290919063ffffffff16565b600d60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508573ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fab8530f87dc9b59234c4623bf917212bb2536d647574c8e7e5da92c2ede0c9f8876040518082815260200191505060405180910390a38573ffffffffffffffffffffffffffffffffffffffff1660007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef876040518082815260200191505060405180910390a36001935050505092915050565b6000600160149054906101000a900460ff16151515611d9857600080fd5b60011515600c60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515141515611df757600080fd5b3360001515600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515141515611e5757600080fd5b600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549150600083111515611ea857600080fd5b828210151515611eb757600080fd5b611ecc83600b54612e1290919063ffffffff16565b600b81905550611ee58383612e1290919063ffffffff16565b600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055503373ffffffffffffffffffffffffffffffffffffffff167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5846040518082815260200191505060405180910390a2600073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef856040518082815260200191505060405180910390a3505050565b6000600160149054906101000a900460ff16151515611fff57600080fd5b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561205b57600080fd5b6001600c60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555081600d60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff167f46980fca912ef9bcdbd36877427b6b90e860769f604e89c0e67720cece530d20836040518082815260200191505060405180910390a26001905092915050565b6121576123be565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561219057600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156121cc57600080fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fb80482a293ca2e013eda8683c9bd7fc8347cfdaeea5ede58cba46df502c2a60460405160405180910390a250565b600160149054906101000a900460ff1681565b6000600960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561232d57600080fd5b60018060146101000a81548160ff0219169083151502179055507f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b6000600d60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60058054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561247d5780601f106124525761010080835404028352916020019161247d565b820191906000526020600020905b81548152906001019060200180831161246057829003601f168201915b505050505081565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600160149054906101000a900460ff161515156124c957600080fd5b3360001515600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151514151561252957600080fd5b8360001515600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151514151561258957600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16141515156125c557600080fd5b600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054841115151561261357600080fd5b61266584600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054612e1290919063ffffffff16565b600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506126fa84600960008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054612e2b90919063ffffffff16565b600960008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef866040518082815260200191505060405180910390a360019250505092915050565b6127b66123be565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156127ef57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561282b57600080fd5b80600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fdb66dfa9c6b8f5226fe9aac7e51897ae8ee94ac31dc70bb6c9900b2574b707e660405160405180910390a250565b6000600c60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b6129326123be565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561296b57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156129a757600080fd5b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fc67398012c111ce95ecb7429b933096c977380ee6c421175a71a4a4c6c88c06e60405160405180910390a250565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b60078054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015612b935780601f10612b6857610100808354040283529160200191612b93565b820191906000526020600020905b815481529060010190602001808311612b7657829003601f168201915b505050505081565b612ba36123be565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515612bdc57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515612c1857600080fd5b7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0612c416123be565b82604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a1612cbf81612e47565b50565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515612d1e57600080fd5b6001600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508073ffffffffffffffffffffffffffffffffffffffff167fffa4e6181777692565cf28528fc88fd1516ea86b56da075235fa575af6a4b85560405160405180910390a250565b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b6000828211151515612e2057fe5b818303905092915050565b60008183019050828110151515612e3e57fe5b80905092915050565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10612ecb57805160ff1916838001178555612ef9565b82800160010185558215612ef9579182015b82811115612ef8578251825591602001919060010190612edd565b5b509050612f069190612f0a565b5090565b612f2c91905b80821115612f28576000816000905550600101612f10565b5090565b905600a165627a7a723058206e88309bb38ccbd9bd656057426ebeeeb874e7c4fa16a1545a2674f6df2b7d430029"

// DeployUsdc deploys a new Ethereum contract, binding an instance of Usdc to it.
func DeployUsdc(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Usdc, error) {
	parsed, err := abi.JSON(strings.NewReader(UsdcABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UsdcBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Usdc{UsdcCaller: UsdcCaller{contract: contract}, UsdcTransactor: UsdcTransactor{contract: contract}, UsdcFilterer: UsdcFilterer{contract: contract}}, nil
}

// Usdc is an auto generated Go binding around an Ethereum contract.
type Usdc struct {
	UsdcCaller     // Read-only binding to the contract
	UsdcTransactor // Write-only binding to the contract
	UsdcFilterer   // Log filterer for contract events
}

// UsdcCaller is an auto generated read-only Go binding around an Ethereum contract.
type UsdcCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsdcTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UsdcTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsdcFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UsdcFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsdcSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UsdcSession struct {
	Contract     *Usdc             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UsdcCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UsdcCallerSession struct {
	Contract *UsdcCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UsdcTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UsdcTransactorSession struct {
	Contract     *UsdcTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UsdcRaw is an auto generated low-level Go binding around an Ethereum contract.
type UsdcRaw struct {
	Contract *Usdc // Generic contract binding to access the raw methods on
}

// UsdcCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UsdcCallerRaw struct {
	Contract *UsdcCaller // Generic read-only contract binding to access the raw methods on
}

// UsdcTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UsdcTransactorRaw struct {
	Contract *UsdcTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUsdc creates a new instance of Usdc, bound to a specific deployed contract.
func NewUsdc(address common.Address, backend bind.ContractBackend) (*Usdc, error) {
	contract, err := bindUsdc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Usdc{UsdcCaller: UsdcCaller{contract: contract}, UsdcTransactor: UsdcTransactor{contract: contract}, UsdcFilterer: UsdcFilterer{contract: contract}}, nil
}

// NewUsdcCaller creates a new read-only instance of Usdc, bound to a specific deployed contract.
func NewUsdcCaller(address common.Address, caller bind.ContractCaller) (*UsdcCaller, error) {
	contract, err := bindUsdc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UsdcCaller{contract: contract}, nil
}

// NewUsdcTransactor creates a new write-only instance of Usdc, bound to a specific deployed contract.
func NewUsdcTransactor(address common.Address, transactor bind.ContractTransactor) (*UsdcTransactor, error) {
	contract, err := bindUsdc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UsdcTransactor{contract: contract}, nil
}

// NewUsdcFilterer creates a new log filterer instance of Usdc, bound to a specific deployed contract.
func NewUsdcFilterer(address common.Address, filterer bind.ContractFilterer) (*UsdcFilterer, error) {
	contract, err := bindUsdc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UsdcFilterer{contract: contract}, nil
}

// bindUsdc binds a generic wrapper to an already deployed contract.
func bindUsdc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UsdcABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Usdc *UsdcRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Usdc.Contract.UsdcCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Usdc *UsdcRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Usdc.Contract.UsdcTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Usdc *UsdcRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Usdc.Contract.UsdcTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Usdc *UsdcCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Usdc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Usdc *UsdcTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Usdc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Usdc *UsdcTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Usdc.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Usdc *UsdcCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Usdc.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Usdc *UsdcSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Usdc.Contract.Allowance(&_Usdc.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Usdc *UsdcCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Usdc.Contract.Allowance(&_Usdc.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Usdc *UsdcCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Usdc.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Usdc *UsdcSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Usdc.Contract.BalanceOf(&_Usdc.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Usdc *UsdcCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Usdc.Contract.BalanceOf(&_Usdc.CallOpts, account)
}

// Blacklister is a free data retrieval call binding the contract method 0xbd102430.
//
// Solidity: function blacklister() view returns(address)
func (_Usdc *UsdcCaller) Blacklister(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Usdc.contract.Call(opts, &out, "blacklister")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Blacklister is a free data retrieval call binding the contract method 0xbd102430.
//
// Solidity: function blacklister() view returns(address)
func (_Usdc *UsdcSession) Blacklister() (common.Address, error) {
	return _Usdc.Contract.Blacklister(&_Usdc.CallOpts)
}

// Blacklister is a free data retrieval call binding the contract method 0xbd102430.
//
// Solidity: function blacklister() view returns(address)
func (_Usdc *UsdcCallerSession) Blacklister() (common.Address, error) {
	return _Usdc.Contract.Blacklister(&_Usdc.CallOpts)
}

// Currency is a free data retrieval call binding the contract method 0xe5a6b10f.
//
// Solidity: function currency() view returns(string)
func (_Usdc *UsdcCaller) Currency(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Usdc.contract.Call(opts, &out, "currency")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Currency is a free data retrieval call binding the contract method 0xe5a6b10f.
//
// Solidity: function currency() view returns(string)
func (_Usdc *UsdcSession) Currency() (string, error) {
	return _Usdc.Contract.Currency(&_Usdc.CallOpts)
}

// Currency is a free data retrieval call binding the contract method 0xe5a6b10f.
//
// Solidity: function currency() view returns(string)
func (_Usdc *UsdcCallerSession) Currency() (string, error) {
	return _Usdc.Contract.Currency(&_Usdc.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Usdc *UsdcCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Usdc.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Usdc *UsdcSession) Decimals() (uint8, error) {
	return _Usdc.Contract.Decimals(&_Usdc.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Usdc *UsdcCallerSession) Decimals() (uint8, error) {
	return _Usdc.Contract.Decimals(&_Usdc.CallOpts)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address _account) view returns(bool)
func (_Usdc *UsdcCaller) IsBlacklisted(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _Usdc.contract.Call(opts, &out, "isBlacklisted", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address _account) view returns(bool)
func (_Usdc *UsdcSession) IsBlacklisted(_account common.Address) (bool, error) {
	return _Usdc.Contract.IsBlacklisted(&_Usdc.CallOpts, _account)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address _account) view returns(bool)
func (_Usdc *UsdcCallerSession) IsBlacklisted(_account common.Address) (bool, error) {
	return _Usdc.Contract.IsBlacklisted(&_Usdc.CallOpts, _account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_Usdc *UsdcCaller) IsMinter(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Usdc.contract.Call(opts, &out, "isMinter", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_Usdc *UsdcSession) IsMinter(account common.Address) (bool, error) {
	return _Usdc.Contract.IsMinter(&_Usdc.CallOpts, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_Usdc *UsdcCallerSession) IsMinter(account common.Address) (bool, error) {
	return _Usdc.Contract.IsMinter(&_Usdc.CallOpts, account)
}

// MasterMinter is a free data retrieval call binding the contract method 0x35d99f35.
//
// Solidity: function masterMinter() view returns(address)
func (_Usdc *UsdcCaller) MasterMinter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Usdc.contract.Call(opts, &out, "masterMinter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MasterMinter is a free data retrieval call binding the contract method 0x35d99f35.
//
// Solidity: function masterMinter() view returns(address)
func (_Usdc *UsdcSession) MasterMinter() (common.Address, error) {
	return _Usdc.Contract.MasterMinter(&_Usdc.CallOpts)
}

// MasterMinter is a free data retrieval call binding the contract method 0x35d99f35.
//
// Solidity: function masterMinter() view returns(address)
func (_Usdc *UsdcCallerSession) MasterMinter() (common.Address, error) {
	return _Usdc.Contract.MasterMinter(&_Usdc.CallOpts)
}

// MinterAllowance is a free data retrieval call binding the contract method 0x8a6db9c3.
//
// Solidity: function minterAllowance(address minter) view returns(uint256)
func (_Usdc *UsdcCaller) MinterAllowance(opts *bind.CallOpts, minter common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Usdc.contract.Call(opts, &out, "minterAllowance", minter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinterAllowance is a free data retrieval call binding the contract method 0x8a6db9c3.
//
// Solidity: function minterAllowance(address minter) view returns(uint256)
func (_Usdc *UsdcSession) MinterAllowance(minter common.Address) (*big.Int, error) {
	return _Usdc.Contract.MinterAllowance(&_Usdc.CallOpts, minter)
}

// MinterAllowance is a free data retrieval call binding the contract method 0x8a6db9c3.
//
// Solidity: function minterAllowance(address minter) view returns(uint256)
func (_Usdc *UsdcCallerSession) MinterAllowance(minter common.Address) (*big.Int, error) {
	return _Usdc.Contract.MinterAllowance(&_Usdc.CallOpts, minter)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Usdc *UsdcCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Usdc.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Usdc *UsdcSession) Name() (string, error) {
	return _Usdc.Contract.Name(&_Usdc.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Usdc *UsdcCallerSession) Name() (string, error) {
	return _Usdc.Contract.Name(&_Usdc.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Usdc *UsdcCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Usdc.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Usdc *UsdcSession) Owner() (common.Address, error) {
	return _Usdc.Contract.Owner(&_Usdc.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Usdc *UsdcCallerSession) Owner() (common.Address, error) {
	return _Usdc.Contract.Owner(&_Usdc.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Usdc *UsdcCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Usdc.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Usdc *UsdcSession) Paused() (bool, error) {
	return _Usdc.Contract.Paused(&_Usdc.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Usdc *UsdcCallerSession) Paused() (bool, error) {
	return _Usdc.Contract.Paused(&_Usdc.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_Usdc *UsdcCaller) Pauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Usdc.contract.Call(opts, &out, "pauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_Usdc *UsdcSession) Pauser() (common.Address, error) {
	return _Usdc.Contract.Pauser(&_Usdc.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_Usdc *UsdcCallerSession) Pauser() (common.Address, error) {
	return _Usdc.Contract.Pauser(&_Usdc.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Usdc *UsdcCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Usdc.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Usdc *UsdcSession) Symbol() (string, error) {
	return _Usdc.Contract.Symbol(&_Usdc.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Usdc *UsdcCallerSession) Symbol() (string, error) {
	return _Usdc.Contract.Symbol(&_Usdc.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Usdc *UsdcCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Usdc.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Usdc *UsdcSession) TotalSupply() (*big.Int, error) {
	return _Usdc.Contract.TotalSupply(&_Usdc.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Usdc *UsdcCallerSession) TotalSupply() (*big.Int, error) {
	return _Usdc.Contract.TotalSupply(&_Usdc.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Usdc *UsdcTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Usdc.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Usdc *UsdcSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Usdc.Contract.Approve(&_Usdc.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Usdc *UsdcTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Usdc.Contract.Approve(&_Usdc.TransactOpts, _spender, _value)
}

// Blacklist is a paid mutator transaction binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist(address _account) returns()
func (_Usdc *UsdcTransactor) Blacklist(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Usdc.contract.Transact(opts, "blacklist", _account)
}

// Blacklist is a paid mutator transaction binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist(address _account) returns()
func (_Usdc *UsdcSession) Blacklist(_account common.Address) (*types.Transaction, error) {
	return _Usdc.Contract.Blacklist(&_Usdc.TransactOpts, _account)
}

// Blacklist is a paid mutator transaction binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist(address _account) returns()
func (_Usdc *UsdcTransactorSession) Blacklist(_account common.Address) (*types.Transaction, error) {
	return _Usdc.Contract.Blacklist(&_Usdc.TransactOpts, _account)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _amount) returns()
func (_Usdc *UsdcTransactor) Burn(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Usdc.contract.Transact(opts, "burn", _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _amount) returns()
func (_Usdc *UsdcSession) Burn(_amount *big.Int) (*types.Transaction, error) {
	return _Usdc.Contract.Burn(&_Usdc.TransactOpts, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _amount) returns()
func (_Usdc *UsdcTransactorSession) Burn(_amount *big.Int) (*types.Transaction, error) {
	return _Usdc.Contract.Burn(&_Usdc.TransactOpts, _amount)
}

// ConfigureMinter is a paid mutator transaction binding the contract method 0x4e44d956.
//
// Solidity: function configureMinter(address minter, uint256 minterAllowedAmount) returns(bool)
func (_Usdc *UsdcTransactor) ConfigureMinter(opts *bind.TransactOpts, minter common.Address, minterAllowedAmount *big.Int) (*types.Transaction, error) {
	return _Usdc.contract.Transact(opts, "configureMinter", minter, minterAllowedAmount)
}

// ConfigureMinter is a paid mutator transaction binding the contract method 0x4e44d956.
//
// Solidity: function configureMinter(address minter, uint256 minterAllowedAmount) returns(bool)
func (_Usdc *UsdcSession) ConfigureMinter(minter common.Address, minterAllowedAmount *big.Int) (*types.Transaction, error) {
	return _Usdc.Contract.ConfigureMinter(&_Usdc.TransactOpts, minter, minterAllowedAmount)
}

// ConfigureMinter is a paid mutator transaction binding the contract method 0x4e44d956.
//
// Solidity: function configureMinter(address minter, uint256 minterAllowedAmount) returns(bool)
func (_Usdc *UsdcTransactorSession) ConfigureMinter(minter common.Address, minterAllowedAmount *big.Int) (*types.Transaction, error) {
	return _Usdc.Contract.ConfigureMinter(&_Usdc.TransactOpts, minter, minterAllowedAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0x3357162b.
//
// Solidity: function initialize(string _name, string _symbol, string _currency, uint8 _decimals, address _masterMinter, address _pauser, address _blacklister, address _owner) returns()
func (_Usdc *UsdcTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, _currency string, _decimals uint8, _masterMinter common.Address, _pauser common.Address, _blacklister common.Address, _owner common.Address) (*types.Transaction, error) {
	return _Usdc.contract.Transact(opts, "initialize", _name, _symbol, _currency, _decimals, _masterMinter, _pauser, _blacklister, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x3357162b.
//
// Solidity: function initialize(string _name, string _symbol, string _currency, uint8 _decimals, address _masterMinter, address _pauser, address _blacklister, address _owner) returns()
func (_Usdc *UsdcSession) Initialize(_name string, _symbol string, _currency string, _decimals uint8, _masterMinter common.Address, _pauser common.Address, _blacklister common.Address, _owner common.Address) (*types.Transaction, error) {
	return _Usdc.Contract.Initialize(&_Usdc.TransactOpts, _name, _symbol, _currency, _decimals, _masterMinter, _pauser, _blacklister, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x3357162b.
//
// Solidity: function initialize(string _name, string _symbol, string _currency, uint8 _decimals, address _masterMinter, address _pauser, address _blacklister, address _owner) returns()
func (_Usdc *UsdcTransactorSession) Initialize(_name string, _symbol string, _currency string, _decimals uint8, _masterMinter common.Address, _pauser common.Address, _blacklister common.Address, _owner common.Address) (*types.Transaction, error) {
	return _Usdc.Contract.Initialize(&_Usdc.TransactOpts, _name, _symbol, _currency, _decimals, _masterMinter, _pauser, _blacklister, _owner)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_Usdc *UsdcTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Usdc.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_Usdc *UsdcSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Usdc.Contract.Mint(&_Usdc.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_Usdc *UsdcTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Usdc.Contract.Mint(&_Usdc.TransactOpts, _to, _amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Usdc *UsdcTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Usdc.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Usdc *UsdcSession) Pause() (*types.Transaction, error) {
	return _Usdc.Contract.Pause(&_Usdc.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Usdc *UsdcTransactorSession) Pause() (*types.Transaction, error) {
	return _Usdc.Contract.Pause(&_Usdc.TransactOpts)
}

// RemoveMinter is a paid mutator transaction binding the contract method 0x3092afd5.
//
// Solidity: function removeMinter(address minter) returns(bool)
func (_Usdc *UsdcTransactor) RemoveMinter(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error) {
	return _Usdc.contract.Transact(opts, "removeMinter", minter)
}

// RemoveMinter is a paid mutator transaction binding the contract method 0x3092afd5.
//
// Solidity: function removeMinter(address minter) returns(bool)
func (_Usdc *UsdcSession) RemoveMinter(minter common.Address) (*types.Transaction, error) {
	return _Usdc.Contract.RemoveMinter(&_Usdc.TransactOpts, minter)
}

// RemoveMinter is a paid mutator transaction binding the contract method 0x3092afd5.
//
// Solidity: function removeMinter(address minter) returns(bool)
func (_Usdc *UsdcTransactorSession) RemoveMinter(minter common.Address) (*types.Transaction, error) {
	return _Usdc.Contract.RemoveMinter(&_Usdc.TransactOpts, minter)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Usdc *UsdcTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Usdc.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Usdc *UsdcSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Usdc.Contract.Transfer(&_Usdc.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Usdc *UsdcTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Usdc.Contract.Transfer(&_Usdc.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Usdc *UsdcTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Usdc.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Usdc *UsdcSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Usdc.Contract.TransferFrom(&_Usdc.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Usdc *UsdcTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Usdc.Contract.TransferFrom(&_Usdc.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Usdc *UsdcTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Usdc.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Usdc *UsdcSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Usdc.Contract.TransferOwnership(&_Usdc.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Usdc *UsdcTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Usdc.Contract.TransferOwnership(&_Usdc.TransactOpts, newOwner)
}

// UnBlacklist is a paid mutator transaction binding the contract method 0x1a895266.
//
// Solidity: function unBlacklist(address _account) returns()
func (_Usdc *UsdcTransactor) UnBlacklist(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Usdc.contract.Transact(opts, "unBlacklist", _account)
}

// UnBlacklist is a paid mutator transaction binding the contract method 0x1a895266.
//
// Solidity: function unBlacklist(address _account) returns()
func (_Usdc *UsdcSession) UnBlacklist(_account common.Address) (*types.Transaction, error) {
	return _Usdc.Contract.UnBlacklist(&_Usdc.TransactOpts, _account)
}

// UnBlacklist is a paid mutator transaction binding the contract method 0x1a895266.
//
// Solidity: function unBlacklist(address _account) returns()
func (_Usdc *UsdcTransactorSession) UnBlacklist(_account common.Address) (*types.Transaction, error) {
	return _Usdc.Contract.UnBlacklist(&_Usdc.TransactOpts, _account)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Usdc *UsdcTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Usdc.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Usdc *UsdcSession) Unpause() (*types.Transaction, error) {
	return _Usdc.Contract.Unpause(&_Usdc.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Usdc *UsdcTransactorSession) Unpause() (*types.Transaction, error) {
	return _Usdc.Contract.Unpause(&_Usdc.TransactOpts)
}

// UpdateBlacklister is a paid mutator transaction binding the contract method 0xad38bf22.
//
// Solidity: function updateBlacklister(address _newBlacklister) returns()
func (_Usdc *UsdcTransactor) UpdateBlacklister(opts *bind.TransactOpts, _newBlacklister common.Address) (*types.Transaction, error) {
	return _Usdc.contract.Transact(opts, "updateBlacklister", _newBlacklister)
}

// UpdateBlacklister is a paid mutator transaction binding the contract method 0xad38bf22.
//
// Solidity: function updateBlacklister(address _newBlacklister) returns()
func (_Usdc *UsdcSession) UpdateBlacklister(_newBlacklister common.Address) (*types.Transaction, error) {
	return _Usdc.Contract.UpdateBlacklister(&_Usdc.TransactOpts, _newBlacklister)
}

// UpdateBlacklister is a paid mutator transaction binding the contract method 0xad38bf22.
//
// Solidity: function updateBlacklister(address _newBlacklister) returns()
func (_Usdc *UsdcTransactorSession) UpdateBlacklister(_newBlacklister common.Address) (*types.Transaction, error) {
	return _Usdc.Contract.UpdateBlacklister(&_Usdc.TransactOpts, _newBlacklister)
}

// UpdateMasterMinter is a paid mutator transaction binding the contract method 0xaa20e1e4.
//
// Solidity: function updateMasterMinter(address _newMasterMinter) returns()
func (_Usdc *UsdcTransactor) UpdateMasterMinter(opts *bind.TransactOpts, _newMasterMinter common.Address) (*types.Transaction, error) {
	return _Usdc.contract.Transact(opts, "updateMasterMinter", _newMasterMinter)
}

// UpdateMasterMinter is a paid mutator transaction binding the contract method 0xaa20e1e4.
//
// Solidity: function updateMasterMinter(address _newMasterMinter) returns()
func (_Usdc *UsdcSession) UpdateMasterMinter(_newMasterMinter common.Address) (*types.Transaction, error) {
	return _Usdc.Contract.UpdateMasterMinter(&_Usdc.TransactOpts, _newMasterMinter)
}

// UpdateMasterMinter is a paid mutator transaction binding the contract method 0xaa20e1e4.
//
// Solidity: function updateMasterMinter(address _newMasterMinter) returns()
func (_Usdc *UsdcTransactorSession) UpdateMasterMinter(_newMasterMinter common.Address) (*types.Transaction, error) {
	return _Usdc.Contract.UpdateMasterMinter(&_Usdc.TransactOpts, _newMasterMinter)
}

// UpdatePauser is a paid mutator transaction binding the contract method 0x554bab3c.
//
// Solidity: function updatePauser(address _newPauser) returns()
func (_Usdc *UsdcTransactor) UpdatePauser(opts *bind.TransactOpts, _newPauser common.Address) (*types.Transaction, error) {
	return _Usdc.contract.Transact(opts, "updatePauser", _newPauser)
}

// UpdatePauser is a paid mutator transaction binding the contract method 0x554bab3c.
//
// Solidity: function updatePauser(address _newPauser) returns()
func (_Usdc *UsdcSession) UpdatePauser(_newPauser common.Address) (*types.Transaction, error) {
	return _Usdc.Contract.UpdatePauser(&_Usdc.TransactOpts, _newPauser)
}

// UpdatePauser is a paid mutator transaction binding the contract method 0x554bab3c.
//
// Solidity: function updatePauser(address _newPauser) returns()
func (_Usdc *UsdcTransactorSession) UpdatePauser(_newPauser common.Address) (*types.Transaction, error) {
	return _Usdc.Contract.UpdatePauser(&_Usdc.TransactOpts, _newPauser)
}

// UsdcApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Usdc contract.
type UsdcApprovalIterator struct {
	Event *UsdcApproval // Event containing the contract specifics and raw log

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
func (it *UsdcApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdcApproval)
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
		it.Event = new(UsdcApproval)
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
func (it *UsdcApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdcApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdcApproval represents a Approval event raised by the Usdc contract.
type UsdcApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Usdc *UsdcFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*UsdcApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Usdc.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &UsdcApprovalIterator{contract: _Usdc.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Usdc *UsdcFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *UsdcApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Usdc.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdcApproval)
				if err := _Usdc.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Usdc *UsdcFilterer) ParseApproval(log types.Log) (*UsdcApproval, error) {
	event := new(UsdcApproval)
	if err := _Usdc.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UsdcBlacklistedIterator is returned from FilterBlacklisted and is used to iterate over the raw logs and unpacked data for Blacklisted events raised by the Usdc contract.
type UsdcBlacklistedIterator struct {
	Event *UsdcBlacklisted // Event containing the contract specifics and raw log

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
func (it *UsdcBlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdcBlacklisted)
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
		it.Event = new(UsdcBlacklisted)
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
func (it *UsdcBlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdcBlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdcBlacklisted represents a Blacklisted event raised by the Usdc contract.
type UsdcBlacklisted struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBlacklisted is a free log retrieval operation binding the contract event 0xffa4e6181777692565cf28528fc88fd1516ea86b56da075235fa575af6a4b855.
//
// Solidity: event Blacklisted(address indexed _account)
func (_Usdc *UsdcFilterer) FilterBlacklisted(opts *bind.FilterOpts, _account []common.Address) (*UsdcBlacklistedIterator, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _Usdc.contract.FilterLogs(opts, "Blacklisted", _accountRule)
	if err != nil {
		return nil, err
	}
	return &UsdcBlacklistedIterator{contract: _Usdc.contract, event: "Blacklisted", logs: logs, sub: sub}, nil
}

// WatchBlacklisted is a free log subscription operation binding the contract event 0xffa4e6181777692565cf28528fc88fd1516ea86b56da075235fa575af6a4b855.
//
// Solidity: event Blacklisted(address indexed _account)
func (_Usdc *UsdcFilterer) WatchBlacklisted(opts *bind.WatchOpts, sink chan<- *UsdcBlacklisted, _account []common.Address) (event.Subscription, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _Usdc.contract.WatchLogs(opts, "Blacklisted", _accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdcBlacklisted)
				if err := _Usdc.contract.UnpackLog(event, "Blacklisted", log); err != nil {
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

// ParseBlacklisted is a log parse operation binding the contract event 0xffa4e6181777692565cf28528fc88fd1516ea86b56da075235fa575af6a4b855.
//
// Solidity: event Blacklisted(address indexed _account)
func (_Usdc *UsdcFilterer) ParseBlacklisted(log types.Log) (*UsdcBlacklisted, error) {
	event := new(UsdcBlacklisted)
	if err := _Usdc.contract.UnpackLog(event, "Blacklisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UsdcBlacklisterChangedIterator is returned from FilterBlacklisterChanged and is used to iterate over the raw logs and unpacked data for BlacklisterChanged events raised by the Usdc contract.
type UsdcBlacklisterChangedIterator struct {
	Event *UsdcBlacklisterChanged // Event containing the contract specifics and raw log

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
func (it *UsdcBlacklisterChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdcBlacklisterChanged)
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
		it.Event = new(UsdcBlacklisterChanged)
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
func (it *UsdcBlacklisterChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdcBlacklisterChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdcBlacklisterChanged represents a BlacklisterChanged event raised by the Usdc contract.
type UsdcBlacklisterChanged struct {
	NewBlacklister common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlacklisterChanged is a free log retrieval operation binding the contract event 0xc67398012c111ce95ecb7429b933096c977380ee6c421175a71a4a4c6c88c06e.
//
// Solidity: event BlacklisterChanged(address indexed newBlacklister)
func (_Usdc *UsdcFilterer) FilterBlacklisterChanged(opts *bind.FilterOpts, newBlacklister []common.Address) (*UsdcBlacklisterChangedIterator, error) {

	var newBlacklisterRule []interface{}
	for _, newBlacklisterItem := range newBlacklister {
		newBlacklisterRule = append(newBlacklisterRule, newBlacklisterItem)
	}

	logs, sub, err := _Usdc.contract.FilterLogs(opts, "BlacklisterChanged", newBlacklisterRule)
	if err != nil {
		return nil, err
	}
	return &UsdcBlacklisterChangedIterator{contract: _Usdc.contract, event: "BlacklisterChanged", logs: logs, sub: sub}, nil
}

// WatchBlacklisterChanged is a free log subscription operation binding the contract event 0xc67398012c111ce95ecb7429b933096c977380ee6c421175a71a4a4c6c88c06e.
//
// Solidity: event BlacklisterChanged(address indexed newBlacklister)
func (_Usdc *UsdcFilterer) WatchBlacklisterChanged(opts *bind.WatchOpts, sink chan<- *UsdcBlacklisterChanged, newBlacklister []common.Address) (event.Subscription, error) {

	var newBlacklisterRule []interface{}
	for _, newBlacklisterItem := range newBlacklister {
		newBlacklisterRule = append(newBlacklisterRule, newBlacklisterItem)
	}

	logs, sub, err := _Usdc.contract.WatchLogs(opts, "BlacklisterChanged", newBlacklisterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdcBlacklisterChanged)
				if err := _Usdc.contract.UnpackLog(event, "BlacklisterChanged", log); err != nil {
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

// ParseBlacklisterChanged is a log parse operation binding the contract event 0xc67398012c111ce95ecb7429b933096c977380ee6c421175a71a4a4c6c88c06e.
//
// Solidity: event BlacklisterChanged(address indexed newBlacklister)
func (_Usdc *UsdcFilterer) ParseBlacklisterChanged(log types.Log) (*UsdcBlacklisterChanged, error) {
	event := new(UsdcBlacklisterChanged)
	if err := _Usdc.contract.UnpackLog(event, "BlacklisterChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UsdcBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Usdc contract.
type UsdcBurnIterator struct {
	Event *UsdcBurn // Event containing the contract specifics and raw log

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
func (it *UsdcBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdcBurn)
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
		it.Event = new(UsdcBurn)
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
func (it *UsdcBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdcBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdcBurn represents a Burn event raised by the Usdc contract.
type UsdcBurn struct {
	Burner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed burner, uint256 amount)
func (_Usdc *UsdcFilterer) FilterBurn(opts *bind.FilterOpts, burner []common.Address) (*UsdcBurnIterator, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _Usdc.contract.FilterLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return &UsdcBurnIterator{contract: _Usdc.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed burner, uint256 amount)
func (_Usdc *UsdcFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *UsdcBurn, burner []common.Address) (event.Subscription, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _Usdc.contract.WatchLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdcBurn)
				if err := _Usdc.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed burner, uint256 amount)
func (_Usdc *UsdcFilterer) ParseBurn(log types.Log) (*UsdcBurn, error) {
	event := new(UsdcBurn)
	if err := _Usdc.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UsdcMasterMinterChangedIterator is returned from FilterMasterMinterChanged and is used to iterate over the raw logs and unpacked data for MasterMinterChanged events raised by the Usdc contract.
type UsdcMasterMinterChangedIterator struct {
	Event *UsdcMasterMinterChanged // Event containing the contract specifics and raw log

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
func (it *UsdcMasterMinterChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdcMasterMinterChanged)
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
		it.Event = new(UsdcMasterMinterChanged)
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
func (it *UsdcMasterMinterChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdcMasterMinterChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdcMasterMinterChanged represents a MasterMinterChanged event raised by the Usdc contract.
type UsdcMasterMinterChanged struct {
	NewMasterMinter common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMasterMinterChanged is a free log retrieval operation binding the contract event 0xdb66dfa9c6b8f5226fe9aac7e51897ae8ee94ac31dc70bb6c9900b2574b707e6.
//
// Solidity: event MasterMinterChanged(address indexed newMasterMinter)
func (_Usdc *UsdcFilterer) FilterMasterMinterChanged(opts *bind.FilterOpts, newMasterMinter []common.Address) (*UsdcMasterMinterChangedIterator, error) {

	var newMasterMinterRule []interface{}
	for _, newMasterMinterItem := range newMasterMinter {
		newMasterMinterRule = append(newMasterMinterRule, newMasterMinterItem)
	}

	logs, sub, err := _Usdc.contract.FilterLogs(opts, "MasterMinterChanged", newMasterMinterRule)
	if err != nil {
		return nil, err
	}
	return &UsdcMasterMinterChangedIterator{contract: _Usdc.contract, event: "MasterMinterChanged", logs: logs, sub: sub}, nil
}

// WatchMasterMinterChanged is a free log subscription operation binding the contract event 0xdb66dfa9c6b8f5226fe9aac7e51897ae8ee94ac31dc70bb6c9900b2574b707e6.
//
// Solidity: event MasterMinterChanged(address indexed newMasterMinter)
func (_Usdc *UsdcFilterer) WatchMasterMinterChanged(opts *bind.WatchOpts, sink chan<- *UsdcMasterMinterChanged, newMasterMinter []common.Address) (event.Subscription, error) {

	var newMasterMinterRule []interface{}
	for _, newMasterMinterItem := range newMasterMinter {
		newMasterMinterRule = append(newMasterMinterRule, newMasterMinterItem)
	}

	logs, sub, err := _Usdc.contract.WatchLogs(opts, "MasterMinterChanged", newMasterMinterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdcMasterMinterChanged)
				if err := _Usdc.contract.UnpackLog(event, "MasterMinterChanged", log); err != nil {
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

// ParseMasterMinterChanged is a log parse operation binding the contract event 0xdb66dfa9c6b8f5226fe9aac7e51897ae8ee94ac31dc70bb6c9900b2574b707e6.
//
// Solidity: event MasterMinterChanged(address indexed newMasterMinter)
func (_Usdc *UsdcFilterer) ParseMasterMinterChanged(log types.Log) (*UsdcMasterMinterChanged, error) {
	event := new(UsdcMasterMinterChanged)
	if err := _Usdc.contract.UnpackLog(event, "MasterMinterChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UsdcMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Usdc contract.
type UsdcMintIterator struct {
	Event *UsdcMint // Event containing the contract specifics and raw log

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
func (it *UsdcMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdcMint)
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
		it.Event = new(UsdcMint)
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
func (it *UsdcMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdcMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdcMint represents a Mint event raised by the Usdc contract.
type UsdcMint struct {
	Minter common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0xab8530f87dc9b59234c4623bf917212bb2536d647574c8e7e5da92c2ede0c9f8.
//
// Solidity: event Mint(address indexed minter, address indexed to, uint256 amount)
func (_Usdc *UsdcFilterer) FilterMint(opts *bind.FilterOpts, minter []common.Address, to []common.Address) (*UsdcMintIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Usdc.contract.FilterLogs(opts, "Mint", minterRule, toRule)
	if err != nil {
		return nil, err
	}
	return &UsdcMintIterator{contract: _Usdc.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0xab8530f87dc9b59234c4623bf917212bb2536d647574c8e7e5da92c2ede0c9f8.
//
// Solidity: event Mint(address indexed minter, address indexed to, uint256 amount)
func (_Usdc *UsdcFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *UsdcMint, minter []common.Address, to []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Usdc.contract.WatchLogs(opts, "Mint", minterRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdcMint)
				if err := _Usdc.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0xab8530f87dc9b59234c4623bf917212bb2536d647574c8e7e5da92c2ede0c9f8.
//
// Solidity: event Mint(address indexed minter, address indexed to, uint256 amount)
func (_Usdc *UsdcFilterer) ParseMint(log types.Log) (*UsdcMint, error) {
	event := new(UsdcMint)
	if err := _Usdc.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UsdcMinterConfiguredIterator is returned from FilterMinterConfigured and is used to iterate over the raw logs and unpacked data for MinterConfigured events raised by the Usdc contract.
type UsdcMinterConfiguredIterator struct {
	Event *UsdcMinterConfigured // Event containing the contract specifics and raw log

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
func (it *UsdcMinterConfiguredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdcMinterConfigured)
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
		it.Event = new(UsdcMinterConfigured)
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
func (it *UsdcMinterConfiguredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdcMinterConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdcMinterConfigured represents a MinterConfigured event raised by the Usdc contract.
type UsdcMinterConfigured struct {
	Minter              common.Address
	MinterAllowedAmount *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMinterConfigured is a free log retrieval operation binding the contract event 0x46980fca912ef9bcdbd36877427b6b90e860769f604e89c0e67720cece530d20.
//
// Solidity: event MinterConfigured(address indexed minter, uint256 minterAllowedAmount)
func (_Usdc *UsdcFilterer) FilterMinterConfigured(opts *bind.FilterOpts, minter []common.Address) (*UsdcMinterConfiguredIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _Usdc.contract.FilterLogs(opts, "MinterConfigured", minterRule)
	if err != nil {
		return nil, err
	}
	return &UsdcMinterConfiguredIterator{contract: _Usdc.contract, event: "MinterConfigured", logs: logs, sub: sub}, nil
}

// WatchMinterConfigured is a free log subscription operation binding the contract event 0x46980fca912ef9bcdbd36877427b6b90e860769f604e89c0e67720cece530d20.
//
// Solidity: event MinterConfigured(address indexed minter, uint256 minterAllowedAmount)
func (_Usdc *UsdcFilterer) WatchMinterConfigured(opts *bind.WatchOpts, sink chan<- *UsdcMinterConfigured, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _Usdc.contract.WatchLogs(opts, "MinterConfigured", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdcMinterConfigured)
				if err := _Usdc.contract.UnpackLog(event, "MinterConfigured", log); err != nil {
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

// ParseMinterConfigured is a log parse operation binding the contract event 0x46980fca912ef9bcdbd36877427b6b90e860769f604e89c0e67720cece530d20.
//
// Solidity: event MinterConfigured(address indexed minter, uint256 minterAllowedAmount)
func (_Usdc *UsdcFilterer) ParseMinterConfigured(log types.Log) (*UsdcMinterConfigured, error) {
	event := new(UsdcMinterConfigured)
	if err := _Usdc.contract.UnpackLog(event, "MinterConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UsdcMinterRemovedIterator is returned from FilterMinterRemoved and is used to iterate over the raw logs and unpacked data for MinterRemoved events raised by the Usdc contract.
type UsdcMinterRemovedIterator struct {
	Event *UsdcMinterRemoved // Event containing the contract specifics and raw log

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
func (it *UsdcMinterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdcMinterRemoved)
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
		it.Event = new(UsdcMinterRemoved)
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
func (it *UsdcMinterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdcMinterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdcMinterRemoved represents a MinterRemoved event raised by the Usdc contract.
type UsdcMinterRemoved struct {
	OldMinter common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMinterRemoved is a free log retrieval operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed oldMinter)
func (_Usdc *UsdcFilterer) FilterMinterRemoved(opts *bind.FilterOpts, oldMinter []common.Address) (*UsdcMinterRemovedIterator, error) {

	var oldMinterRule []interface{}
	for _, oldMinterItem := range oldMinter {
		oldMinterRule = append(oldMinterRule, oldMinterItem)
	}

	logs, sub, err := _Usdc.contract.FilterLogs(opts, "MinterRemoved", oldMinterRule)
	if err != nil {
		return nil, err
	}
	return &UsdcMinterRemovedIterator{contract: _Usdc.contract, event: "MinterRemoved", logs: logs, sub: sub}, nil
}

// WatchMinterRemoved is a free log subscription operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed oldMinter)
func (_Usdc *UsdcFilterer) WatchMinterRemoved(opts *bind.WatchOpts, sink chan<- *UsdcMinterRemoved, oldMinter []common.Address) (event.Subscription, error) {

	var oldMinterRule []interface{}
	for _, oldMinterItem := range oldMinter {
		oldMinterRule = append(oldMinterRule, oldMinterItem)
	}

	logs, sub, err := _Usdc.contract.WatchLogs(opts, "MinterRemoved", oldMinterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdcMinterRemoved)
				if err := _Usdc.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
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

// ParseMinterRemoved is a log parse operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed oldMinter)
func (_Usdc *UsdcFilterer) ParseMinterRemoved(log types.Log) (*UsdcMinterRemoved, error) {
	event := new(UsdcMinterRemoved)
	if err := _Usdc.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UsdcOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Usdc contract.
type UsdcOwnershipTransferredIterator struct {
	Event *UsdcOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *UsdcOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdcOwnershipTransferred)
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
		it.Event = new(UsdcOwnershipTransferred)
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
func (it *UsdcOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdcOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdcOwnershipTransferred represents a OwnershipTransferred event raised by the Usdc contract.
type UsdcOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address previousOwner, address newOwner)
func (_Usdc *UsdcFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts) (*UsdcOwnershipTransferredIterator, error) {

	logs, sub, err := _Usdc.contract.FilterLogs(opts, "OwnershipTransferred")
	if err != nil {
		return nil, err
	}
	return &UsdcOwnershipTransferredIterator{contract: _Usdc.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address previousOwner, address newOwner)
func (_Usdc *UsdcFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UsdcOwnershipTransferred) (event.Subscription, error) {

	logs, sub, err := _Usdc.contract.WatchLogs(opts, "OwnershipTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdcOwnershipTransferred)
				if err := _Usdc.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address previousOwner, address newOwner)
func (_Usdc *UsdcFilterer) ParseOwnershipTransferred(log types.Log) (*UsdcOwnershipTransferred, error) {
	event := new(UsdcOwnershipTransferred)
	if err := _Usdc.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UsdcPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the Usdc contract.
type UsdcPauseIterator struct {
	Event *UsdcPause // Event containing the contract specifics and raw log

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
func (it *UsdcPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdcPause)
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
		it.Event = new(UsdcPause)
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
func (it *UsdcPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdcPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdcPause represents a Pause event raised by the Usdc contract.
type UsdcPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Usdc *UsdcFilterer) FilterPause(opts *bind.FilterOpts) (*UsdcPauseIterator, error) {

	logs, sub, err := _Usdc.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &UsdcPauseIterator{contract: _Usdc.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Usdc *UsdcFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *UsdcPause) (event.Subscription, error) {

	logs, sub, err := _Usdc.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdcPause)
				if err := _Usdc.contract.UnpackLog(event, "Pause", log); err != nil {
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

// ParsePause is a log parse operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Usdc *UsdcFilterer) ParsePause(log types.Log) (*UsdcPause, error) {
	event := new(UsdcPause)
	if err := _Usdc.contract.UnpackLog(event, "Pause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UsdcPauserChangedIterator is returned from FilterPauserChanged and is used to iterate over the raw logs and unpacked data for PauserChanged events raised by the Usdc contract.
type UsdcPauserChangedIterator struct {
	Event *UsdcPauserChanged // Event containing the contract specifics and raw log

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
func (it *UsdcPauserChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdcPauserChanged)
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
		it.Event = new(UsdcPauserChanged)
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
func (it *UsdcPauserChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdcPauserChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdcPauserChanged represents a PauserChanged event raised by the Usdc contract.
type UsdcPauserChanged struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPauserChanged is a free log retrieval operation binding the contract event 0xb80482a293ca2e013eda8683c9bd7fc8347cfdaeea5ede58cba46df502c2a604.
//
// Solidity: event PauserChanged(address indexed newAddress)
func (_Usdc *UsdcFilterer) FilterPauserChanged(opts *bind.FilterOpts, newAddress []common.Address) (*UsdcPauserChangedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Usdc.contract.FilterLogs(opts, "PauserChanged", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &UsdcPauserChangedIterator{contract: _Usdc.contract, event: "PauserChanged", logs: logs, sub: sub}, nil
}

// WatchPauserChanged is a free log subscription operation binding the contract event 0xb80482a293ca2e013eda8683c9bd7fc8347cfdaeea5ede58cba46df502c2a604.
//
// Solidity: event PauserChanged(address indexed newAddress)
func (_Usdc *UsdcFilterer) WatchPauserChanged(opts *bind.WatchOpts, sink chan<- *UsdcPauserChanged, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Usdc.contract.WatchLogs(opts, "PauserChanged", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdcPauserChanged)
				if err := _Usdc.contract.UnpackLog(event, "PauserChanged", log); err != nil {
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

// ParsePauserChanged is a log parse operation binding the contract event 0xb80482a293ca2e013eda8683c9bd7fc8347cfdaeea5ede58cba46df502c2a604.
//
// Solidity: event PauserChanged(address indexed newAddress)
func (_Usdc *UsdcFilterer) ParsePauserChanged(log types.Log) (*UsdcPauserChanged, error) {
	event := new(UsdcPauserChanged)
	if err := _Usdc.contract.UnpackLog(event, "PauserChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UsdcTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Usdc contract.
type UsdcTransferIterator struct {
	Event *UsdcTransfer // Event containing the contract specifics and raw log

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
func (it *UsdcTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdcTransfer)
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
		it.Event = new(UsdcTransfer)
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
func (it *UsdcTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdcTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdcTransfer represents a Transfer event raised by the Usdc contract.
type UsdcTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Usdc *UsdcFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*UsdcTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Usdc.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &UsdcTransferIterator{contract: _Usdc.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Usdc *UsdcFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *UsdcTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Usdc.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdcTransfer)
				if err := _Usdc.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Usdc *UsdcFilterer) ParseTransfer(log types.Log) (*UsdcTransfer, error) {
	event := new(UsdcTransfer)
	if err := _Usdc.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UsdcUnBlacklistedIterator is returned from FilterUnBlacklisted and is used to iterate over the raw logs and unpacked data for UnBlacklisted events raised by the Usdc contract.
type UsdcUnBlacklistedIterator struct {
	Event *UsdcUnBlacklisted // Event containing the contract specifics and raw log

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
func (it *UsdcUnBlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdcUnBlacklisted)
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
		it.Event = new(UsdcUnBlacklisted)
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
func (it *UsdcUnBlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdcUnBlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdcUnBlacklisted represents a UnBlacklisted event raised by the Usdc contract.
type UsdcUnBlacklisted struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnBlacklisted is a free log retrieval operation binding the contract event 0x117e3210bb9aa7d9baff172026820255c6f6c30ba8999d1c2fd88e2848137c4e.
//
// Solidity: event UnBlacklisted(address indexed _account)
func (_Usdc *UsdcFilterer) FilterUnBlacklisted(opts *bind.FilterOpts, _account []common.Address) (*UsdcUnBlacklistedIterator, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _Usdc.contract.FilterLogs(opts, "UnBlacklisted", _accountRule)
	if err != nil {
		return nil, err
	}
	return &UsdcUnBlacklistedIterator{contract: _Usdc.contract, event: "UnBlacklisted", logs: logs, sub: sub}, nil
}

// WatchUnBlacklisted is a free log subscription operation binding the contract event 0x117e3210bb9aa7d9baff172026820255c6f6c30ba8999d1c2fd88e2848137c4e.
//
// Solidity: event UnBlacklisted(address indexed _account)
func (_Usdc *UsdcFilterer) WatchUnBlacklisted(opts *bind.WatchOpts, sink chan<- *UsdcUnBlacklisted, _account []common.Address) (event.Subscription, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _Usdc.contract.WatchLogs(opts, "UnBlacklisted", _accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdcUnBlacklisted)
				if err := _Usdc.contract.UnpackLog(event, "UnBlacklisted", log); err != nil {
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

// ParseUnBlacklisted is a log parse operation binding the contract event 0x117e3210bb9aa7d9baff172026820255c6f6c30ba8999d1c2fd88e2848137c4e.
//
// Solidity: event UnBlacklisted(address indexed _account)
func (_Usdc *UsdcFilterer) ParseUnBlacklisted(log types.Log) (*UsdcUnBlacklisted, error) {
	event := new(UsdcUnBlacklisted)
	if err := _Usdc.contract.UnpackLog(event, "UnBlacklisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UsdcUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the Usdc contract.
type UsdcUnpauseIterator struct {
	Event *UsdcUnpause // Event containing the contract specifics and raw log

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
func (it *UsdcUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdcUnpause)
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
		it.Event = new(UsdcUnpause)
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
func (it *UsdcUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdcUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdcUnpause represents a Unpause event raised by the Usdc contract.
type UsdcUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Usdc *UsdcFilterer) FilterUnpause(opts *bind.FilterOpts) (*UsdcUnpauseIterator, error) {

	logs, sub, err := _Usdc.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &UsdcUnpauseIterator{contract: _Usdc.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Usdc *UsdcFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *UsdcUnpause) (event.Subscription, error) {

	logs, sub, err := _Usdc.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdcUnpause)
				if err := _Usdc.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// ParseUnpause is a log parse operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Usdc *UsdcFilterer) ParseUnpause(log types.Log) (*UsdcUnpause, error) {
	event := new(UsdcUnpause)
	if err := _Usdc.contract.UnpackLog(event, "Unpause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
