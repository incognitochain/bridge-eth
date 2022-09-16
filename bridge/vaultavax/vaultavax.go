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

// VaultBurnInstData is an auto generated low-level Go binding around an user-defined struct.
type VaultBurnInstData struct {
	Meta   uint8
	Shard  uint8
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Itx    [32]byte
}

// VaultRedepositOptions is an auto generated low-level Go binding around an user-defined struct.
type VaultRedepositOptions struct {
	RedepositToken      common.Address
	RedepositIncAddress []byte
	WithdrawAddress     common.Address
}

// VaultMetaData contains all meta data concerning the Vault contract.
var VaultMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositID\",\"type\":\"uint256\"}],\"name\":\"DepositV2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"phaseID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"errorData\",\"type\":\"bytes\"}],\"name\":\"ExecuteFnLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"redepositIncAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"itx\",\"type\":\"bytes32\"}],\"name\":\"Redeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newIncognitoProxy\",\"type\":\"address\"}],\"name\":\"UpdateIncognitoProxy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"UpdateTokenTotal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BURN_CALL_REQUEST_METADATA_TYPE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BURN_REQUEST_METADATA_TYPE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BURN_TO_CONTRACT_REQUEST_METADATA_TYPE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CURRENT_NETWORK_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETH_TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"externalCalldata\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"redepositToken\",\"type\":\"address\"}],\"name\":\"_callExternal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"_transferExternal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"txId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signData\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"txId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signData\",\"type\":\"bytes\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"txId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signData\",\"type\":\"bytes\"}],\"name\":\"depositERC20_V2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"txId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signData\",\"type\":\"bytes\"}],\"name\":\"deposit_V2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipientToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"exchangeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"timestamp\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signData\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"heights\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"instPaths\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"instPathIsLefts\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"instRoots\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"sigIdxs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"sigVs\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigRs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigSs\",\"type\":\"bytes32[]\"}],\"name\":\"executeWithBurnProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"executor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getDepositedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_prevVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_regulator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"isSigDataUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"isWithdrawed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"migration\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"notEntered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"}],\"name\":\"parseBurnInst\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"meta\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"shard\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"itx\",\"type\":\"bytes32\"}],\"internalType\":\"structVault.BurnInstData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"}],\"name\":\"parseCalldataFromBurnInst\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"meta\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"shard\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"itx\",\"type\":\"bytes32\"}],\"internalType\":\"structVault.BurnInstData\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"redepositToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"redepositIncAddress\",\"type\":\"bytes\"},{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"internalType\":\"structVault.RedepositOptions\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prevVault\",\"outputs\":[{\"internalType\":\"contractWithdrawable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"regulator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"timestamp\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"txId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"regulatorSig\",\"type\":\"bytes\"}],\"name\":\"requestWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_regulator\",\"type\":\"address\"}],\"name\":\"setRegulator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"sigDataUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"sigToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"storageLayoutVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"heights\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"instPaths\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"instPathIsLefts\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"instRoots\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"sigIdxs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"sigVs\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigRs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigSs\",\"type\":\"bytes32[]\"}],\"name\":\"submitBurnProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalDepositedToSCAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"updateAssets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_regulator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"upgradeVaultStorage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"heights\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"instPaths\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"instPathIsLefts\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"instRoots\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"sigIdxs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"sigVs\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigRs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigSs\",\"type\":\"bytes32[]\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawRequests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"withdrawed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506183dc80620000226000396000f3fe6080604052600436106102345760003560e01c8063995fac111161012e578063cde0a4f8116100ab578063dd8fee141161006f578063dd8fee14146108a3578063e4bd7074146108ce578063f75b98ce1461090b578063fa84702e14610948578063fee8efda146109735761023b565b8063cde0a4f8146107aa578063cf54aaa0146107d3578063d6a1fe3b14610810578063d7200eb11461083b578063dca40d9e146108665761023b565b8063bd835c42116100f2578063bd835c42146106d2578063bda9b509146106fd578063c0c53b8b1461073a578063c34c08e514610763578063c791d7051461078e5761023b565b8063995fac11146105ef578063a3f5d8cc1461062c578063a73b153214610657578063a807b5bb14610680578063b8237dbb146106a95761023b565b80636304541c116101bc57806373bf96511161018057806373bf965114610514578063749c5f861461053d5780637e16e6e11461057a57806384b3ac03146105b75780638588ccd6146105d35761023b565b80636304541c146103f357806365b5a00f1461043057806366945b311461046d5780636f2cbc48146104ac57806370a08231146104d75761023b565b8063392e53cd11610203578063392e53cd1461030c5780633ed1b376146103375780633fec6b4014610360578063568c04fd1461039d57806358bc8337146103c85761023b565b8063145e2a6b146102405780631beb7de2146102695780631ea1940e146102925780631ed4276d146102cf5761023b565b3661023b57005b600080fd5b34801561024c57600080fd5b5061026760048036038101906102629190616633565b61099c565b005b34801561027557600080fd5b50610290600480360381019061028b9190616c6d565b610b22565b005b34801561029e57600080fd5b506102b960048036038101906102b491906169a1565b6111a8565b6040516102c69190617af2565b60405180910390f35b3480156102db57600080fd5b506102f660048036038101906102f19190616903565b6111c8565b6040516103039190617af2565b60405180910390f35b34801561031857600080fd5b5061032161154e565b60405161032e9190617af2565b60405180910390f35b34801561034357600080fd5b5061035e60048036038101906103599190616a0f565b611561565b005b34801561036c57600080fd5b5061038760048036038101906103829190616c19565b611c44565b60405161039491906177fd565b60405180910390f35b3480156103a957600080fd5b506103b2611cd2565b6040516103bf9190617ee0565b60405180910390f35b3480156103d457600080fd5b506103dd611cd7565b6040516103ea91906177fd565b60405180910390f35b3480156103ff57600080fd5b5061041a60048036038101906104159190616509565b611cdc565b6040516104279190617ec5565b60405180910390f35b34801561043c57600080fd5b50610457600480360381019061045291906165f7565b611cf4565b6040516104649190617ec5565b60405180910390f35b34801561047957600080fd5b50610494600480360381019061048f91906169ca565b611d19565b6040516104a393929190617dd6565b60405180910390f35b3480156104b857600080fd5b506104c1612059565b6040516104ce9190617ee0565b60405180910390f35b3480156104e357600080fd5b506104fe60048036038101906104f99190616509565b61205e565b60405161050b9190617ec5565b60405180910390f35b34801561052057600080fd5b5061053b60048036038101906105369190616c6d565b6121af565b005b34801561054957600080fd5b50610564600480360381019061055f91906169a1565b6127e4565b6040516105719190617af2565b60405180910390f35b34801561058657600080fd5b506105a1600480360381019061059c9190616b97565b612927565b6040516105ae9190617dbb565b60405180910390f35b6105d160048036038101906105cc9190616f00565b612a42565b005b6105ed60048036038101906105e89190616760565b612bd8565b005b3480156105fb57600080fd5b50610616600480360381019061061191906165f7565b613381565b6040516106239190617af2565b60405180910390f35b34801561063857600080fd5b506106416133b0565b60405161064e9190617af2565b60405180910390f35b34801561066357600080fd5b5061067e600480360381019061067991906165f7565b6133c3565b005b34801561068c57600080fd5b506106a760048036038101906106a29190616853565b613541565b005b3480156106b557600080fd5b506106d060048036038101906106cb9190616853565b61390a565b005b3480156106de57600080fd5b506106e7613ce8565b6040516106f49190617ee0565b60405180910390f35b34801561070957600080fd5b50610724600480360381019061071f91906166d1565b613ced565b6040516107319190617ec5565b60405180910390f35b34801561074657600080fd5b50610761600480360381019061075c9190616682565b61406f565b005b34801561076f57600080fd5b50610778614270565b60405161078591906177fd565b60405180910390f35b6107a860048036038101906107a39190616f00565b614296565b005b3480156107b657600080fd5b506107d160048036038101906107cc9190616509565b614417565b005b3480156107df57600080fd5b506107fa60048036038101906107f59190616509565b61458a565b6040516108079190617ee0565b60405180910390f35b34801561081c57600080fd5b50610825614699565b6040516108329190617ec5565b60405180910390f35b34801561084757600080fd5b5061085061469f565b60405161085d9190617ee0565b60405180910390f35b34801561087257600080fd5b5061088d600480360381019061088891906169a1565b6146a4565b60405161089a9190617af2565b60405180910390f35b3480156108af57600080fd5b506108b86146c4565b6040516108c591906177fd565b60405180910390f35b3480156108da57600080fd5b506108f560048036038101906108f091906169a1565b6146ea565b6040516109029190617af2565b60405180910390f35b34801561091757600080fd5b50610932600480360381019061092d91906165f7565b61482e565b60405161093f9190617ec5565b60405180910390f35b34801561095457600080fd5b5061095d614ae2565b60405161096a9190617cbe565b60405180910390f35b34801561097f57600080fd5b5061099a60048036038101906109959190616df0565b614b08565b005b3373ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff16146109d56014614f9a565b90610a16576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a0d9190617cd9565b60405180910390fd5b50600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610a5b57610a56828261514a565b610b1d565b8273ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83836040518363ffffffff1660e01b8152600401610a9692919061786a565b600060405180830381600087803b158015610ab057600080fd5b505af1158015610ac4573d6000803e3d6000fd5b50505050610ad061523e565b610ada6004614f9a565b90610b1b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b129190617cd9565b60405180910390fd5b505b505050565b600560149054906101000a900460ff16610b3c6001614f9a565b90610b7d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b749190617cd9565b60405180910390fd5b506000600560146101000a81548160ff02191690831515021790555060828a511015610ba96006614f9a565b90610bea576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610be19190617cd9565b60405180910390fd5b50610bf3615f6f565b610bfc8b612927565b905060f160ff16816000015160ff16148015610c1f57506001816020015160ff16145b610c296006614f9a565b90610c6a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c619190617cd9565b60405180910390fd5b50610c788160a001516127e4565b15610c836005614f9a565b90610cc4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cbb9190617cd9565b60405180910390fd5b5060016000808360a00151815260200190815260200160002060006101000a81548160ff021916908315150217905550600073ffffffffffffffffffffffffffffffffffffffff16816040015173ffffffffffffffffffffffffffffffffffffffff161415610ddb57610d8760046000836040015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054826080015161527c90919063ffffffff16565b471015610d946007614f9a565b90610dd5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dcc9190617cd9565b60405180910390fd5b50610f59565b6000610dea826040015161458a565b905060098160ff161115610e2257610e1860098260ff1603600a0a836080015161531690919063ffffffff16565b8260800181815250505b610e7c60046000846040015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054836080015161527c90919063ffffffff16565b826040015173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401610eb99190617818565b60206040518083038186803b158015610ed157600080fd5b505afa158015610ee5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f099190616f89565b1015610f156007614f9a565b90610f56576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f4d9190617cd9565b60405180910390fd5b50505b610f6b8b8b8b8b8b8b8b8b8b8b6153b7565b600073ffffffffffffffffffffffffffffffffffffffff16816040015173ffffffffffffffffffffffffffffffffffffffff16141561106b576000816060015173ffffffffffffffffffffffffffffffffffffffff168260800151604051610fd2906177e8565b60006040518083038185875af1925050503d806000811461100f576040519150601f19603f3d011682016040523d82523d6000602084013e611014565b606091505b50509050806110236004614f9a565b90611064576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161105b9190617cd9565b60405180910390fd5b5050611139565b806040015173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb826060015183608001516040518363ffffffff1660e01b81526004016110b292919061786a565b600060405180830381600087803b1580156110cc57600080fd5b505af11580156110e0573d6000803e3d6000fd5b505050506110ec61523e565b6110f66004614f9a565b90611137576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161112e9190617cd9565b60405180910390fd5b505b7f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb816040015182606001518360800151604051611178939291906178f3565b60405180910390a1506001600560146101000a81548160ff02191690831515021790555050505050505050505050565b60016020528060005260406000206000915054906101000a900460ff1681565b60008073ffffffffffffffffffffffffffffffffffffffff16600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141580156112755750600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b61127f600c614f9a565b906112c0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112b79190617cd9565b60405180910390fd5b508282905085859050146112d4600a614f9a565b90611315576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161130c9190617cd9565b60405180910390fd5b50600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635c975abb6040518163ffffffff1660e01b815260040160206040518083038186803b15801561137e57600080fd5b505afa158015611392573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113b69190616978565b6113c0600d614f9a565b90611401576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113f89190617cd9565b60405180910390fd5b5060005b858590508110156115045761149484848381811061141f57fe5b905060200201356004600089898681811061143657fe5b905060200201602081019061144b9190616509565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461527c90919063ffffffff16565b600460008888858181106114a457fe5b90506020020160208101906114b99190616509565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508080600101915050611405565b507f6a7fbbcddfd518bb8c56b28ac6c7acb0f7ca093ed232eb3306e53d14e469895f8585858560405161153a9493929190617ab7565b60405180910390a160019050949350505050565b600560159054906101000a900460ff1681565b600560149054906101000a900460ff1661157b6001614f9a565b906115bc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115b39190617cd9565b60405180910390fd5b506000600560146101000a81548160ff0219169083151502179055506115e0615f6f565b6115e8615fda565b60606115f48e8e611d19565b9250925092506116078360a001516127e4565b156116126005614f9a565b90611653576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161164a9190617cd9565b60405180910390fd5b5060016000808560a00151815260200190815260200160002060006101000a81548160ff021916908315150217905550600073ffffffffffffffffffffffffffffffffffffffff16836040015173ffffffffffffffffffffffffffffffffffffffff16141561176a5761171660046000856040015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054846080015161527c90919063ffffffff16565b4710156117236007614f9a565b90611764576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161175b9190617cd9565b60405180910390fd5b506118e8565b6000611779846040015161458a565b905060098160ff1611156117b1576117a760098260ff1603600a0a856080015161531690919063ffffffff16565b8460800181815250505b61180b60046000866040015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054856080015161527c90919063ffffffff16565b846040015173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016118489190617818565b60206040518083038186803b15801561186057600080fd5b505afa158015611874573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118989190616f89565b10156118a46007614f9a565b906118e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118dc9190617cd9565b60405180910390fd5b50505b61193e8e8e8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508d8d8d8d8d8d8d8d8d6153b7565b3073ffffffffffffffffffffffffffffffffffffffff1663bda9b5098460400151856060015186608001518587600001516040518663ffffffff1660e01b815260040161198f95949392919061792a565b602060405180830381600087803b1580156119a957600080fd5b505af19250505080156119da57506040513d601f19601f820116820180604052508101906119d79190616f89565b60015b611a75573d8060008114611a0a576040519150601f19603f3d011682016040523d82523d6000602084013e611a0f565b606091505b50611a2c8460400151846020015186608001518760a001516154e2565b7fdbbb883f24557adf486292429863dcfd4ac5d4db168ae94921da8e3d9a95d4168460a00151600083604051611a6493929190617bfd565b60405180910390a150505050611c1c565b600073ffffffffffffffffffffffffffffffffffffffff16836040015173ffffffffffffffffffffffffffffffffffffffff161415611acb57611ac683600001518460200151838760a001516154e2565b611c17565b3073ffffffffffffffffffffffffffffffffffffffff1663145e2a6b84600001518560400151846040518463ffffffff1660e01b8152600401611b10939291906178bc565b600060405180830381600087803b158015611b2a57600080fd5b505af1925050508015611b3b575060015b611bd3573d8060008114611b6b576040519150601f19603f3d011682016040523d82523d6000602084013e611b70565b606091505b50611b8984600001518560200151848860a001516154e2565b7fdbbb883f24557adf486292429863dcfd4ac5d4db168ae94921da8e3d9a95d4168560a00151600183604051611bc193929190617c3b565b60405180910390a15050505050611c1c565b7f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb8360000151846040015183604051611c0e939291906178f3565b60405180910390a15b505050505b6001600560146101000a81548160ff0219169083151502179055505050505050505050505050565b6000806000806020860151915060408601519250601b86604081518110611c6757fe5b602001015160f81c60f81b60f81c01905060018582848660405160008152602001604052604051611c9b9493929190617c79565b6020604051602081039080840390855afa158015611cbd573d6000803e3d6000fd5b50505060206040510351935050505092915050565b60f181565b600081565b60046020528060005260406000206000915090505481565b6002602052816000526040600020602052806000526040600020600091509150505481565b611d21615f6f565b611d29615fda565b6060610128858590501015611d3e6013614f9a565b90611d7f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d769190617cd9565b60405180910390fd5b50611d88615f6f565b85856000818110611d9557fe5b9050013560f81c60f81b60f81c816000019060ff16908160ff168152505085856001818110611dc057fe5b9050013560f81c60f81b60f81c816020019060ff16908160ff1681525050600086866002818110611ded57fe5b9050013560f81c60f81b60f81c9050609e60ff16826000015160ff16148015611e1d57506001826020015160ff16145b8015611e2f5750600160ff168160ff16145b611e396013614f9a565b90611e7a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e719190617cd9565b60405180910390fd5b5050611e84615fda565b868660039060c392611e989392919061816b565b810190611ea59190616532565b8760400188606001896080018a60a0018a6000018b6040018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152508673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681525086815250868152508673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152508673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815250505050505050868660c39061012892611fa09392919061816b565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508160200181905250818188886101289080926120019392919061816b565b81818080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509050905094509450945050509250925092565b60f381565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561209c574790506121aa565b6120a582615649565b6120e4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120db90617d5b565b60405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b815260040161211d9190617818565b60206040518083038186803b15801561213557600080fd5b505afa92505050801561216657506040513d601f19601f820116820180604052508101906121639190616f89565b60015b6121a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161219c90617d9b565b60405180910390fd5b809150505b919050565b600560149054906101000a900460ff166121c96001614f9a565b9061220a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122019190617cd9565b60405180910390fd5b506000600560146101000a81548160ff02191690831515021790555060828a5110156122366006614f9a565b90612277576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161226e9190617cd9565b60405180910390fd5b50612280615f6f565b6122898b612927565b905060f360ff16816000015160ff161480156122ac57506001816020015160ff16145b6122b66006614f9a565b906122f7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122ee9190617cd9565b60405180910390fd5b506123058160a001516127e4565b156123106005614f9a565b90612351576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016123489190617cd9565b60405180910390fd5b5060016000808360a00151815260200190815260200160002060006101000a81548160ff021916908315150217905550600073ffffffffffffffffffffffffffffffffffffffff16816040015173ffffffffffffffffffffffffffffffffffffffff1614156124685761241460046000836040015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054826080015161527c90919063ffffffff16565b4710156124216007614f9a565b90612462576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016124599190617cd9565b60405180910390fd5b506125e6565b6000612477826040015161458a565b905060098160ff1611156124af576124a560098260ff1603600a0a836080015161531690919063ffffffff16565b8260800181815250505b61250960046000846040015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054836080015161527c90919063ffffffff16565b826040015173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016125469190617818565b60206040518083038186803b15801561255e57600080fd5b505afa158015612572573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125969190616f89565b10156125a26007614f9a565b906125e3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016125da9190617cd9565b60405180910390fd5b50505b6125f88b8b8b8b8b8b8b8b8b8b6153b7565b612693816080015160026000846060015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000846040015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461527c90919063ffffffff16565b60026000836060015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000836040015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550612775816080015160046000846040015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461527c90919063ffffffff16565b60046000836040015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550506001600560146101000a81548160ff02191690831515021790555050505050505050505050565b600080600083815260200190815260200160002060009054906101000a900460ff16156128145760019050612922565b600073ffffffffffffffffffffffffffffffffffffffff16600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156128745760009050612922565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663749c5f86836040518263ffffffff1660e01b81526004016128cf9190617be2565b60206040518083038186803b1580156128e757600080fd5b505afa1580156128fb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061291f9190616978565b90505b919050565b61292f615f6f565b612937615f6f565b8260008151811061294457fe5b602001015160f81c60f81b60f81c816000019060ff16908160ff16815250508260018151811061297057fe5b602001015160f81c60f81b60f81c816020019060ff16908160ff16815250506000806000806022870151935060428701519250606287015191506082870151905083856040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505082856060019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505081856080018181525050808560a00181815250508495505050505050919050565b600560149054906101000a900460ff16612a5c6001614f9a565b90612a9d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612a949190617cd9565b60405180910390fd5b506000600560146101000a81548160ff0219169083151502179055506b033b2e3c9fd0803ce8000000471115612ad36002614f9a565b90612b14576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b0b9190617cd9565b60405180910390fd5b50612b638383838080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505061565c565b7fd30df8040a1092415b49422a02dbd8cdd5915a596abcba02cd0f65dd86ab38516000868634612b93600661575c565b604051612ba4959493929190617a40565b60405180910390a1612bb6600661576a565b6001600560146101000a81548160ff0219169083151502179055505050505050565b600560149054906101000a900460ff16612bf26001614f9a565b90612c33576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c2a9190617cd9565b60405180910390fd5b506000600560146101000a81548160ff0219169083151502179055506000612cd3612c6260008d88888f615780565b8a8a8a8a604051602001612c7a959493929190617e1c565b60405160208183030381529060405284848080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050615825565b9050612cdf818c615949565b89600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015612d696008614f9a565b90612daa576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612da19190617cd9565b60405180910390fd5b50612dfd8a600460008e73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054615c8990919063ffffffff16565b600460008d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550612ecf8a600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008e73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054615c8990919063ffffffff16565b600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506000349050600073ffffffffffffffffffffffffffffffffffffffff168c73ffffffffffffffffffffffffffffffffffffffff161415612fa357612f9c8b8261527c90919063ffffffff16565b905061315f565b8a8c73ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401612fdd9190617818565b60206040518083038186803b158015612ff557600080fd5b505afa158015613009573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061302d9190616f89565b10156130396007614f9a565b9061307a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016130719190617cd9565b60405180910390fd5b508b73ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168d6040518363ffffffff1660e01b81526004016130d8929190617a8e565b600060405180830381600087803b1580156130f257600080fd5b505af1158015613106573d6000803e3d6000fd5b5050505061311261523e565b61311c6004614f9a565b9061315d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016131549190617cd9565b60405180910390fd5b505b60006131b18b838b8b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508d615d12565b905061324281600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008e73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461527c90919063ffffffff16565b600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061331481600460008e73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461527c90919063ffffffff16565b600460008d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050506001600560146101000a81548160ff02191690831515021790555050505050505050505050565b60036020528160005260406000206020528060005260406000206000915091509054906101000a900460ff1681565b600560149054906101000a900460ff1681565b6000600854146133d36012614f9a565b90613414576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161340b9190617cd9565b60405180910390fd5b506002600881905550600073ffffffffffffffffffffffffffffffffffffffff16600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146134796011614f9a565b906134ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016134b19190617cd9565b60405180910390fd5b5081600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b600560149054906101000a900460ff1661355b6001614f9a565b9061359c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016135939190617cd9565b60405180910390fd5b506000600560146101000a81548160ff0219169083151502179055506136068383838080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505061565c565b600087905060006136168961458a565b905060008273ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016136539190617818565b60206040518083038186803b15801561366b57600080fd5b505afa15801561367f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906136a39190616f89565b9050600081905060008a905060098460ff1611156136e65760098460ff1603600a0a81816136cd57fe5b04905060098460ff1603600a0a83816136e257fe5b0492505b670de0b6b3a764000081111580156137065750670de0b6b3a76400008311155b801561372c5750670de0b6b3a7640000613729848361527c90919063ffffffff16565b11155b6137366003614f9a565b90613777576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161376e9190617cd9565b60405180910390fd5b508473ffffffffffffffffffffffffffffffffffffffff166323b872dd33308e6040518463ffffffff1660e01b81526004016137b593929190617833565b600060405180830381600087803b1580156137cf57600080fd5b505af11580156137e3573d6000803e3d6000fd5b505050506137ef61523e565b6137f96004614f9a565b9061383a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016138319190617cd9565b60405180910390fd5b508a613857836138498f61205e565b615c8990919063ffffffff16565b14613862600a614f9a565b906138a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161389a9190617cd9565b60405180910390fd5b507f2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e8c8b8b846040516138d99493929190617a00565b60405180910390a150505050506001600560146101000a81548160ff02191690831515021790555050505050505050565b600560149054906101000a900460ff166139246001614f9a565b90613965576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161395c9190617cd9565b60405180910390fd5b506000600560146101000a81548160ff0219169083151502179055506139cf8383838080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505061565c565b600087905060006139df8961458a565b905060008273ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401613a1c9190617818565b60206040518083038186803b158015613a3457600080fd5b505afa158015613a48573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613a6c9190616f89565b9050600081905060008a905060098460ff161115613aaf5760098460ff1603600a0a8181613a9657fe5b04905060098460ff1603600a0a8381613aab57fe5b0492505b670de0b6b3a76400008111158015613acf5750670de0b6b3a76400008311155b8015613af55750670de0b6b3a7640000613af2848361527c90919063ffffffff16565b11155b613aff6003614f9a565b90613b40576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613b379190617cd9565b60405180910390fd5b508473ffffffffffffffffffffffffffffffffffffffff166323b872dd33308e6040518463ffffffff1660e01b8152600401613b7e93929190617833565b600060405180830381600087803b158015613b9857600080fd5b505af1158015613bac573d6000803e3d6000fd5b50505050613bb861523e565b613bc26004614f9a565b90613c03576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613bfa9190617cd9565b60405180910390fd5b508a613c2083613c128f61205e565b615c8990919063ffffffff16565b14613c2b600a614f9a565b90613c6c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613c639190617cd9565b60405180910390fd5b507fd30df8040a1092415b49422a02dbd8cdd5915a596abcba02cd0f65dd86ab38518c8b8b84613c9c600661575c565b604051613cad959493929190617a40565b60405180910390a1613cbf600661576a565b50505050506001600560146101000a81548160ff02191690831515021790555050505050505050565b609e81565b60003373ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614613d286014614f9a565b90613d69576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613d609190617cd9565b60405180910390fd5b506000613d758361205e565b9050606060008073ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff161415613db757869050613e9b565b8873ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16896040518363ffffffff1660e01b8152600401613e14929190617a8e565b600060405180830381600087803b158015613e2e57600080fd5b505af1158015613e42573d6000803e3d6000fd5b50505050613e4e61523e565b613e586004614f9a565b90613e99576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613e909190617cd9565b60405180910390fd5b505b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631cff79cd828a896040518463ffffffff1660e01b8152600401613ef9929190617984565b6000604051808303818588803b158015613f1257600080fd5b505af1158015613f26573d6000803e3d6000fd5b50505050506040513d6000823e3d601f19601f82011682018060405250810190613f509190616bd8565b91506040825114613f616009614f9a565b90613fa2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613f999190617cd9565b60405180910390fd5b5060008083806020019051810190613fba91906165bb565b915091508673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16148015614012575080614010866140028a61205e565b615c8990919063ffffffff16565b145b61401c6009614f9a565b9061405d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016140549190617cd9565b60405180910390fd5b50809550505050505095945050505050565b600560159054906101000a900460ff161561408a600f614f9a565b906140cb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016140c29190617cd9565b60405180910390fd5b5082600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001600560156101000a81548160ff0219169083151502179055506001600560146101000a81548160ff021916908315150217905550600073ffffffffffffffffffffffffffffffffffffffff16600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461419f6011614f9a565b906141e0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016141d79190617cd9565b60405180910390fd5b5081600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506002600881905550505050565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600560149054906101000a900460ff166142b06001614f9a565b906142f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016142e89190617cd9565b60405180910390fd5b506000600560146101000a81548160ff0219169083151502179055506b033b2e3c9fd0803ce80000004711156143276002614f9a565b90614368576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161435f9190617cd9565b60405180910390fd5b506143b78383838080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505061565c565b7f2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e60008686346040516143ed9493929190617a00565b60405180910390a16001600560146101000a81548160ff0219169083151502179055505050505050565b600073ffffffffffffffffffffffffffffffffffffffff16600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614806144c15750600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b80156144fa5750600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b6145046011614f9a565b90614545576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161453c9190617cd9565b60405180910390fd5b5080600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600061459582615649565b6145d4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016145cb90617d7b565b60405180910390fd5b60008290508073ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b815260040160206040518083038186803b15801561461f57600080fd5b505afa92505050801561465057506040513d601f19601f8201168201806040525081019061464d9190616f89565b60015b61468f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161468690617cfb565b60405180910390fd5b8092505050919050565b60085481565b600181565b60006020528060005260406000206000915054906101000a900460ff1681565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006001600083815260200190815260200160002060009054906101000a900460ff161561471b5760019050614829565b600073ffffffffffffffffffffffffffffffffffffffff16600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561477b5760009050614829565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e4bd7074836040518263ffffffff1660e01b81526004016147d69190617be2565b60206040518083038186803b1580156147ee57600080fd5b505afa158015614802573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906148269190616978565b90505b919050565b60008073ffffffffffffffffffffffffffffffffffffffff16600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141580156149155750600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16155b15614a5c57614a55600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f75b98ce85856040518363ffffffff1660e01b815260040161497a929190617893565b60206040518083038186803b15801561499257600080fd5b505afa1580156149a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906149ca9190616f89565b600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461527c90919063ffffffff16565b9050614adc565b600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490505b92915050565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600560149054906101000a900460ff16614b226001614f9a565b90614b63576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614b5a9190617cd9565b60405180910390fd5b506000600560146101000a81548160ff021916908315150217905550614bcd8383838080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505061565c565b6000614c4d614be060018c89898e615780565b8d8d604051602001614bf493929190617e71565b60405160208183030381529060405289898080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050615825565b9050614c59818b615949565b88600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015614ce36008614f9a565b90614d24576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614d1b9190617cd9565b60405180910390fd5b50614db489600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054615c8990919063ffffffff16565b600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550614e8689600460008d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054615c8990919063ffffffff16565b600460008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506000899050600073ffffffffffffffffffffffffffffffffffffffff168b73ffffffffffffffffffffffffffffffffffffffff1614614f33576000614f0d8c61458a565b905060098160ff161115614f315760098160ff1603600a0a8b81614f2d57fe5b0491505b505b7f2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e8b8e8e84604051614f689493929190617a00565b60405180910390a150506001600560146101000a81548160ff0219169083151502179055505050505050505050505050565b60606000826014811115614faa57fe5b90506000600a905060608167ffffffffffffffff81118015614fcb57600080fd5b506040519080825280601f01601f191660200182016040528015614ffe5781602001600182028036833780820191505090505b50905060005b60008460ff161461507f576000600a8560ff168161501e57fe5b069050600a8560ff168161502e57fe5b0494508060300160f81b83838060010194508151811061504a57fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535050615004565b60606001820167ffffffffffffffff8111801561509b57600080fd5b506040519080825280601f01601f1916602001820160405280156150ce5781602001600182028036833780820191505090505b50905060005b82811161513c5783818403815181106150e957fe5b602001015160f81c60f81b82828151811061510057fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080806001019150506150d4565b508095505050505050919050565b8047101561518d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161518490617d3b565b60405180910390fd5b60008273ffffffffffffffffffffffffffffffffffffffff16826040516151b3906177e8565b60006040518083038185875af1925050503d80600081146151f0576040519150601f19603f3d011682016040523d82523d6000602084013e6151f5565b606091505b5050905080615239576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161523090617d1b565b60405180910390fd5b505050565b600080600090503d6000811461525b576020811461526457615270565b60019150615270565b60206000803e60005191505b50600081141591505090565b60008082840190508381101580156152945750828110155b6040518060400160405280601281526020017f536166654d61746820657863657074696f6e00000000000000000000000000008152509061530b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016153029190617cd9565b60405180910390fd5b508091505092915050565b6000808284029050600084148061533557508284828161533257fe5b04145b6040518060400160405280601281526020017f536166654d61746820657863657074696f6e0000000000000000000000000000815250906153ac576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016153a39190617cd9565b60405180910390fd5b508091505092915050565b60008a8a6040516020016153cc9291906177c0565b6040516020818303038152906040528051906020012090506153ec615f3e565b73ffffffffffffffffffffffffffffffffffffffff1663f65d21166001838d8d8d8d8d8d8d8d8d6040518c63ffffffff1660e01b81526004016154399b9a99989796959493929190617b0d565b60206040518083038186803b15801561545157600080fd5b505afa158015615465573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906154899190616978565b6154936006614f9a565b906154d4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016154cb9190617cd9565b60405180910390fd5b505050505050505050505050565b6000829050600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16141561557d576b033b2e3c9fd0803ce80000004711156155366002614f9a565b90615577576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161556e9190617cd9565b60405180910390fd5b50615606565b60006155888661458a565b905060098160ff1611156155ac5760098160ff1603600a0a82816155a857fe5b0491505b670de0b6b3a76400008211156155c26003614f9a565b90615603576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016155fa9190617cd9565b60405180910390fd5b50505b7eb45d95b5117447e2fafe7f34def913ff3ba220e4b8688acf37ae2328af7a3d8585838560405161563a94939291906179b4565b60405180910390a15050505050565b600080823b905060008111915050919050565b60006156b88260405180604001604052803373ffffffffffffffffffffffffffffffffffffffff1681526020018681525060405160200161569d9190617eaa565b60405160208183030381529060405280519060200120611c44565b9050600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146157156010614f9a565b90615756576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161574d9190617cd9565b60405180910390fd5b50505050565b600081600001549050919050565b6001816000016000828254019250508190555050565b615788616027565b615790616027565b60405180608001604052808860078111156157a757fe5b81526020018773ffffffffffffffffffffffffffffffffffffffff16815260200186868080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505081526020018481525090508091505095945050505050565b6000808380519060200120905061583b816146ea565b156158466005614f9a565b90615887576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161587e9190617cd9565b60405180910390fd5b5060006158948483611c44565b9050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156158d16010614f9a565b90615912576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016159099190617cd9565b60405180910390fd5b50600180600084815260200190815260200160002060006101000a81548160ff021916908315150217905550809250505092915050565b600073ffffffffffffffffffffffffffffffffffffffff16600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614158015615a2f5750600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16155b15615c8557615b6f600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f75b98ce83856040518363ffffffff1660e01b8152600401615a94929190617893565b60206040518083038186803b158015615aac57600080fd5b505afa158015615ac0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615ae49190616f89565b600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461527c90919063ffffffff16565b600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b5050565b6000828211156040518060400160405280601281526020017f536166654d61746820657863657074696f6e000000000000000000000000000081525090615d06576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615cfd9190617cd9565b60405180910390fd5b50818303905092915050565b600080615d1e8661205e565b9050600073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff161415615d6b57615d683482615c8990919063ffffffff16565b90505b84471015615d796007614f9a565b90615dba576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615db19190617cd9565b60405180910390fd5b506060600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631cff79cd8786886040518463ffffffff1660e01b8152600401615e1b929190617984565b6000604051808303818588803b158015615e3457600080fd5b505af1158015615e48573d6000803e3d6000fd5b50505050506040513d6000823e3d601f19601f82011682018060405250810190615e729190616bd8565b905060008082806020019051810190615e8b91906165bb565b915091508873ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16148015615ee3575080615ee185615ed38c61205e565b615c8990919063ffffffff16565b145b615eed6009614f9a565b90615f2e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615f259190617cd9565b60405180910390fd5b5080945050505050949350505050565b6000807f62135fc083646fdb4e1a9d700e351b886a4a5a39da980650269edd1ade91ffd260001b9050805491505090565b6040518060c00160405280600060ff168152602001600060ff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600080191681525090565b6040518060600160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff1681525090565b60405180608001604052806000600781111561603f57fe5b8152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001600081525090565b60008135905061607f8161831c565b92915050565b60008135905061609481618333565b92915050565b6000815190506160a981618333565b92915050565b60008083601f8401126160c157600080fd5b8235905067ffffffffffffffff8111156160da57600080fd5b6020830191508360208202830111156160f257600080fd5b9250929050565b600082601f83011261610a57600080fd5b813561611d61611882617f28565b617efb565b9150818183526020840193506020810190508385602084028201111561614257600080fd5b60005b838110156161725781616158888261634f565b845260208401935060208301925050600181019050616145565b5050505092915050565b600082601f83011261618d57600080fd5b81356161a061619b82617f50565b617efb565b915081818352602084019350602081019050838560208402820111156161c557600080fd5b60005b838110156161f557816161db8882616379565b8452602084019350602083019250506001810190506161c8565b5050505092915050565b60008083601f84011261621157600080fd5b8235905067ffffffffffffffff81111561622a57600080fd5b60208301915083602082028301111561624257600080fd5b9250929050565b600082601f83011261625a57600080fd5b813561626d61626882617f78565b617efb565b9150818183526020840193506020810190508385602084028201111561629257600080fd5b60005b838110156162c257816162a888826164ca565b845260208401935060208301925050600181019050616295565b5050505092915050565b600082601f8301126162dd57600080fd5b81356162f06162eb82617fa0565b617efb565b9150818183526020840193506020810190508385602084028201111561631557600080fd5b60005b83811015616345578161632b88826164f4565b845260208401935060208301925050600181019050616318565b5050505092915050565b60008135905061635e8161834a565b92915050565b6000815190506163738161834a565b92915050565b60008135905061638881618361565b92915050565b60008083601f8401126163a057600080fd5b8235905067ffffffffffffffff8111156163b957600080fd5b6020830191508360018202830111156163d157600080fd5b9250929050565b600082601f8301126163e957600080fd5b81356163fc6163f782617fc8565b617efb565b9150808252602083016020830185838301111561641857600080fd5b6164238382846182b2565b50505092915050565b600082601f83011261643d57600080fd5b815161645061644b82617fc8565b617efb565b9150808252602083016020830185838301111561646c57600080fd5b6164778382846182c1565b50505092915050565b60008083601f84011261649257600080fd5b8235905067ffffffffffffffff8111156164ab57600080fd5b6020830191508360018202830111156164c357600080fd5b9250929050565b6000813590506164d981618378565b92915050565b6000815190506164ee81618378565b92915050565b6000813590506165038161838f565b92915050565b60006020828403121561651b57600080fd5b600061652984828501616070565b91505092915050565b60008060008060008060c0878903121561654b57600080fd5b600061655989828a01616085565b965050602061656a89828a01616085565b955050604061657b89828a016164ca565b945050606061658c89828a01616379565b935050608061659d89828a01616085565b92505060a06165ae89828a01616085565b9150509295509295509295565b600080604083850312156165ce57600080fd5b60006165dc8582860161609a565b92505060206165ed858286016164df565b9150509250929050565b6000806040838503121561660a57600080fd5b600061661885828601616070565b925050602061662985828601616070565b9150509250929050565b60008060006060848603121561664857600080fd5b600061665686828701616070565b935050602061666786828701616085565b9250506040616678868287016164ca565b9150509250925092565b60008060006060848603121561669757600080fd5b60006166a586828701616070565b93505060206166b686828701616070565b92505060406166c786828701616070565b9150509250925092565b600080600080600060a086880312156166e957600080fd5b60006166f788828901616070565b955050602061670888828901616070565b9450506040616719888289016164ca565b935050606086013567ffffffffffffffff81111561673657600080fd5b616742888289016163d8565b925050608061675388828901616070565b9150509295509295909350565b60008060008060008060008060008060e08b8d03121561677f57600080fd5b600061678d8d828e01616070565b9a5050602061679e8d828e016164ca565b99505060406167af8d828e01616070565b98505060606167c08d828e01616070565b97505060808b013567ffffffffffffffff8111156167dd57600080fd5b6167e98d828e0161638e565b965096505060a08b013567ffffffffffffffff81111561680857600080fd5b6168148d828e0161638e565b945094505060c08b013567ffffffffffffffff81111561683357600080fd5b61683f8d828e0161638e565b92509250509295989b9194979a5092959850565b600080600080600080600060a0888a03121561686e57600080fd5b600061687c8a828b01616070565b975050602061688d8a828b016164ca565b965050604088013567ffffffffffffffff8111156168aa57600080fd5b6168b68a828b01616480565b955095505060606168c98a828b01616379565b935050608088013567ffffffffffffffff8111156168e657600080fd5b6168f28a828b0161638e565b925092505092959891949750929550565b6000806000806040858703121561691957600080fd5b600085013567ffffffffffffffff81111561693357600080fd5b61693f878288016160af565b9450945050602085013567ffffffffffffffff81111561695e57600080fd5b61696a878288016161ff565b925092505092959194509250565b60006020828403121561698a57600080fd5b600061699884828501616364565b91505092915050565b6000602082840312156169b357600080fd5b60006169c184828501616379565b91505092915050565b600080602083850312156169dd57600080fd5b600083013567ffffffffffffffff8111156169f757600080fd5b616a038582860161638e565b92509250509250929050565b60008060008060008060008060008060006101408c8e031215616a3157600080fd5b60008c013567ffffffffffffffff811115616a4b57600080fd5b616a578e828f0161638e565b9b509b50506020616a6a8e828f016164ca565b99505060408c013567ffffffffffffffff811115616a8757600080fd5b616a938e828f0161617c565b98505060608c013567ffffffffffffffff811115616ab057600080fd5b616abc8e828f016160f9565b9750506080616acd8e828f01616379565b96505060a0616ade8e828f01616379565b95505060c08c013567ffffffffffffffff811115616afb57600080fd5b616b078e828f01616249565b94505060e08c013567ffffffffffffffff811115616b2457600080fd5b616b308e828f016162cc565b9350506101008c013567ffffffffffffffff811115616b4e57600080fd5b616b5a8e828f0161617c565b9250506101208c013567ffffffffffffffff811115616b7857600080fd5b616b848e828f0161617c565b9150509295989b509295989b9093969950565b600060208284031215616ba957600080fd5b600082013567ffffffffffffffff811115616bc357600080fd5b616bcf848285016163d8565b91505092915050565b600060208284031215616bea57600080fd5b600082015167ffffffffffffffff811115616c0457600080fd5b616c108482850161642c565b91505092915050565b60008060408385031215616c2c57600080fd5b600083013567ffffffffffffffff811115616c4657600080fd5b616c52858286016163d8565b9250506020616c6385828601616379565b9150509250929050565b6000806000806000806000806000806101408b8d031215616c8d57600080fd5b60008b013567ffffffffffffffff811115616ca757600080fd5b616cb38d828e016163d8565b9a50506020616cc48d828e016164ca565b99505060408b013567ffffffffffffffff811115616ce157600080fd5b616ced8d828e0161617c565b98505060608b013567ffffffffffffffff811115616d0a57600080fd5b616d168d828e016160f9565b9750506080616d278d828e01616379565b96505060a0616d388d828e01616379565b95505060c08b013567ffffffffffffffff811115616d5557600080fd5b616d618d828e01616249565b94505060e08b013567ffffffffffffffff811115616d7e57600080fd5b616d8a8d828e016162cc565b9350506101008b013567ffffffffffffffff811115616da857600080fd5b616db48d828e0161617c565b9250506101208b013567ffffffffffffffff811115616dd257600080fd5b616dde8d828e0161617c565b9150509295989b9194979a5092959850565b600080600080600080600080600080600060e08c8e031215616e1157600080fd5b60008c013567ffffffffffffffff811115616e2b57600080fd5b616e378e828f01616480565b9b509b50506020616e4a8e828f01616070565b9950506040616e5b8e828f016164ca565b98505060608c013567ffffffffffffffff811115616e7857600080fd5b616e848e828f0161638e565b975097505060808c013567ffffffffffffffff811115616ea357600080fd5b616eaf8e828f0161638e565b955095505060a0616ec28e828f01616379565b93505060c08c013567ffffffffffffffff811115616edf57600080fd5b616eeb8e828f0161638e565b92509250509295989b509295989b9093969950565b600080600080600060608688031215616f1857600080fd5b600086013567ffffffffffffffff811115616f3257600080fd5b616f3e88828901616480565b95509550506020616f5188828901616379565b935050604086013567ffffffffffffffff811115616f6e57600080fd5b616f7a8882890161638e565b92509250509295509295909350565b600060208284031215616f9b57600080fd5b6000616fa9848285016164df565b91505092915050565b6000616fbe8383617057565b60208301905092915050565b6000616fd683836172a2565b60208301905092915050565b6000616fee83836172c0565b60208301905092915050565b6000617006838361776d565b60208301905092915050565b600061701e83836177a2565b60208301905092915050565b61703381618222565b82525050565b617042816181b0565b82525050565b617051816181b0565b82525050565b6170608161819e565b82525050565b61706f8161819e565b82525050565b600061708183856180c1565b935061708c82617ff4565b8060005b858110156170c5576170a28284618154565b6170ac8882616fb2565b97506170b783618080565b925050600181019050617090565b5085925050509392505050565b60006170dd8261803e565b6170e781856180d2565b93506170f283617ffe565b8060005b8381101561712357815161710a8882616fca565b97506171158361808d565b9250506001810190506170f6565b5085935050505092915050565b600061713b82618049565b61714581856180e3565b93506171508361800e565b8060005b838110156171815781516171688882616fe2565b97506171738361809a565b925050600181019050617154565b5085935050505092915050565b600061719a83856180f4565b93507f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8311156171c957600080fd5b6020830292506171da8385846182b2565b82840190509392505050565b60006171f182618054565b6171fb81856180f4565b93506172068361801e565b8060005b8381101561723757815161721e8882616ffa565b9750617229836180a7565b92505060018101905061720a565b5085935050505092915050565b600061724f8261805f565b6172598185618105565b93506172648361802e565b8060005b8381101561729557815161727c8882617012565b9750617287836180b4565b925050600181019050617268565b5085935050505092915050565b6172ab816181c2565b82525050565b6172ba816181c2565b82525050565b6172c9816181ce565b82525050565b6172d8816181ce565b82525050565b60006172ea8385618127565b93506172f78385846182b2565b617300836182fe565b840190509392505050565b60006173168261806a565b6173208185618116565b93506173308185602086016182c1565b617339816182fe565b840191505092915050565b600061734f8261806a565b6173598185618127565b93506173698185602086016182c1565b617372816182fe565b840191505092915050565b60006173888261806a565b6173928185618138565b93506173a28185602086016182c1565b80840191505092915050565b6173b781618234565b82525050565b6173c681618258565b82525050565b6173d58161826a565b82525050565b6173e48161827c565b82525050565b60006173f68385618143565b93506174038385846182b2565b61740c836182fe565b840190509392505050565b600061742282618075565b61742c8185618143565b935061743c8185602086016182c1565b617445816182fe565b840191505092915050565b600061745d601883618143565b91507f67657420455243323020646563696d616c206661696c656400000000000000006000830152602082019050919050565b600061749d603a83618143565b91507f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260008301527f6563697069656e74206d617920686176652072657665727465640000000000006020830152604082019050919050565b6000617503601d83618143565b91507f416464726573733a20696e73756666696369656e742062616c616e63650000006000830152602082019050919050565b6000617543601683618143565b91507f62616c616e63654f66206e6f6e2d636f6e7472616374000000000000000000006000830152602082019050919050565b6000617583600083618138565b9150600082019050919050565b600061759d601883618143565b91507f676574446563696d616c73206e6f6e2d636f6e747261637400000000000000006000830152602082019050919050565b60006175dd601883618143565b91507f6765742045524332302062616c616e6365206661696c656400000000000000006000830152602082019050919050565b60c08201600082015161762660008501826177a2565b50602082015161763960208501826177a2565b50604082015161764c6040850182617057565b50606082015161765f6060850182617039565b506080820151617672608085018261776d565b5060a082015161768560a08501826172c0565b50505050565b60006080830160008301516176a360008601826173bd565b5060208301516176b66020860182617057565b50604083015184820360408601526176ce828261730b565b91505060608301516176e3606086018261776d565b508091505092915050565b60006060830160008301516177066000860182617057565b506020830151848203602086015261771e828261730b565b91505060408301516177336040860182617039565b508091505092915050565b6040820160008201516177546000850182617057565b50602082015161776760208501826172c0565b50505050565b6177768161820b565b82525050565b6177858161820b565b82525050565b61779c6177978261820b565b6182f4565b82525050565b6177ab81618215565b82525050565b6177ba81618215565b82525050565b60006177cc828561737d565b91506177d8828461778b565b6020820191508190509392505050565b60006177f382617576565b9150819050919050565b60006020820190506178126000830184617066565b92915050565b600060208201905061782d600083018461702a565b92915050565b6000606082019050617848600083018661702a565b617855602083018561702a565b617862604083018461777c565b949350505050565b600060408201905061787f600083018561702a565b61788c602083018461777c565b9392505050565b60006040820190506178a86000830185617066565b6178b56020830184617066565b9392505050565b60006060820190506178d16000830186617066565b6178de6020830185617048565b6178eb604083018461777c565b949350505050565b60006060820190506179086000830186617066565b617915602083018561702a565b617922604083018461777c565b949350505050565b600060a08201905061793f6000830188617066565b61794c602083018761702a565b617959604083018661777c565b818103606083015261796b8185617344565b905061797a6080830184617066565b9695505050505050565b60006040820190506179996000830185617066565b81810360208301526179ab8184617344565b90509392505050565b60006080820190506179c96000830187617066565b81810360208301526179db8186617344565b90506179ea604083018561777c565b6179f760608301846172cf565b95945050505050565b6000606082019050617a156000830187617066565b8181036020830152617a288185876173ea565b9050617a37604083018461777c565b95945050505050565b6000608082019050617a556000830188617066565b8181036020830152617a688186886173ea565b9050617a77604083018561777c565b617a84606083018461777c565b9695505050505050565b6000604082019050617aa36000830185617066565b617ab0602083018461777c565b9392505050565b60006040820190508181036000830152617ad2818688617075565b90508181036020830152617ae781848661718e565b905095945050505050565b6000602082019050617b0760008301846172b1565b92915050565b600061016082019050617b23600083018e6172b1565b617b30602083018d6172cf565b617b3d604083018c61777c565b8181036060830152617b4f818b617130565b90508181036080830152617b63818a6170d2565b9050617b7260a08301896172cf565b617b7f60c08301886172cf565b81810360e0830152617b9181876171e6565b9050818103610100830152617ba68186617244565b9050818103610120830152617bbb8185617130565b9050818103610140830152617bd08184617130565b90509c9b505050505050505050505050565b6000602082019050617bf760008301846172cf565b92915050565b6000606082019050617c1260008301866172cf565b617c1f60208301856173cc565b8181036040830152617c318184617344565b9050949350505050565b6000606082019050617c5060008301866172cf565b617c5d60208301856173db565b8181036040830152617c6f8184617344565b9050949350505050565b6000608082019050617c8e60008301876172cf565b617c9b60208301866177b1565b617ca860408301856172cf565b617cb560608301846172cf565b95945050505050565b6000602082019050617cd360008301846173ae565b92915050565b60006020820190508181036000830152617cf38184617417565b905092915050565b60006020820190508181036000830152617d1481617450565b9050919050565b60006020820190508181036000830152617d3481617490565b9050919050565b60006020820190508181036000830152617d54816174f6565b9050919050565b60006020820190508181036000830152617d7481617536565b9050919050565b60006020820190508181036000830152617d9481617590565b9050919050565b60006020820190508181036000830152617db4816175d0565b9050919050565b600060c082019050617dd06000830184617610565b92915050565b600061010082019050617dec6000830186617610565b81810360c0830152617dfe81856176ee565b905081810360e0830152617e128184617344565b9050949350505050565b60006080820190508181036000830152617e36818861768b565b9050617e456020830187617066565b617e526040830186617066565b8181036060830152617e658184866172de565b90509695505050505050565b60006040820190508181036000830152617e8b818661768b565b90508181036020830152617ea08184866173ea565b9050949350505050565b6000604082019050617ebf600083018461773e565b92915050565b6000602082019050617eda600083018461777c565b92915050565b6000602082019050617ef560008301846177b1565b92915050565b6000604051905081810181811067ffffffffffffffff82111715617f1e57600080fd5b8060405250919050565b600067ffffffffffffffff821115617f3f57600080fd5b602082029050602081019050919050565b600067ffffffffffffffff821115617f6757600080fd5b602082029050602081019050919050565b600067ffffffffffffffff821115617f8f57600080fd5b602082029050602081019050919050565b600067ffffffffffffffff821115617fb757600080fd5b602082029050602081019050919050565b600067ffffffffffffffff821115617fdf57600080fd5b601f19601f8301169050602081019050919050565b6000819050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b60006181636020840184616070565b905092915050565b6000808585111561817b57600080fd5b8386111561818857600080fd5b6001850283019150848603905094509492505050565b60006181a9826181eb565b9050919050565b60006181bb826181eb565b9050919050565b60008115159050919050565b6000819050919050565b60008190506181e68261830f565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b600061822d8261828e565b9050919050565b600061823f82618246565b9050919050565b6000618251826181eb565b9050919050565b6000618263826181d8565b9050919050565b60006182758261820b565b9050919050565b60006182878261820b565b9050919050565b6000618299826182a0565b9050919050565b60006182ab826181eb565b9050919050565b82818337600083830152505050565b60005b838110156182df5780820151818401526020810190506182c4565b838111156182ee576000848401525b50505050565b6000819050919050565b6000601f19601f8301169050919050565b6008811061831957fe5b50565b6183258161819e565b811461833057600080fd5b50565b61833c816181b0565b811461834757600080fd5b50565b618353816181c2565b811461835e57600080fd5b50565b61836a816181ce565b811461837557600080fd5b50565b6183818161820b565b811461838c57600080fd5b50565b61839881618215565b81146183a357600080fd5b5056fea264697066735822122078ac983421841d5be6218b852b5cf580cadd31e43643f21ca4354a4eaabf32f564736f6c634300060c0033",
}

// VaultABI is the input ABI used to generate the binding from.
// Deprecated: Use VaultMetaData.ABI instead.
var VaultABI = VaultMetaData.ABI

// VaultBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VaultMetaData.Bin instead.
var VaultBin = VaultMetaData.Bin

// DeployVault deploys a new Ethereum contract, binding an instance of Vault to it.
func DeployVault(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Vault, error) {
	parsed, err := VaultMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VaultBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Vault{VaultCaller: VaultCaller{contract: contract}, VaultTransactor: VaultTransactor{contract: contract}, VaultFilterer: VaultFilterer{contract: contract}}, nil
}

// Vault is an auto generated Go binding around an Ethereum contract.
type Vault struct {
	VaultCaller     // Read-only binding to the contract
	VaultTransactor // Write-only binding to the contract
	VaultFilterer   // Log filterer for contract events
}

// VaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultSession struct {
	Contract     *Vault            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultCallerSession struct {
	Contract *VaultCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultTransactorSession struct {
	Contract     *VaultTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultRaw struct {
	Contract *Vault // Generic contract binding to access the raw methods on
}

// VaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultCallerRaw struct {
	Contract *VaultCaller // Generic read-only contract binding to access the raw methods on
}

// VaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultTransactorRaw struct {
	Contract *VaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVault creates a new instance of Vault, bound to a specific deployed contract.
func NewVault(address common.Address, backend bind.ContractBackend) (*Vault, error) {
	contract, err := bindVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vault{VaultCaller: VaultCaller{contract: contract}, VaultTransactor: VaultTransactor{contract: contract}, VaultFilterer: VaultFilterer{contract: contract}}, nil
}

// NewVaultCaller creates a new read-only instance of Vault, bound to a specific deployed contract.
func NewVaultCaller(address common.Address, caller bind.ContractCaller) (*VaultCaller, error) {
	contract, err := bindVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultCaller{contract: contract}, nil
}

// NewVaultTransactor creates a new write-only instance of Vault, bound to a specific deployed contract.
func NewVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultTransactor, error) {
	contract, err := bindVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultTransactor{contract: contract}, nil
}

// NewVaultFilterer creates a new log filterer instance of Vault, bound to a specific deployed contract.
func NewVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultFilterer, error) {
	contract, err := bindVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultFilterer{contract: contract}, nil
}

// bindVault binds a generic wrapper to an already deployed contract.
func bindVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.VaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transact(opts, method, params...)
}

// BURNCALLREQUESTMETADATATYPE is a free data retrieval call binding the contract method 0xbd835c42.
//
// Solidity: function BURN_CALL_REQUEST_METADATA_TYPE() view returns(uint8)
func (_Vault *VaultCaller) BURNCALLREQUESTMETADATATYPE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "BURN_CALL_REQUEST_METADATA_TYPE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BURNCALLREQUESTMETADATATYPE is a free data retrieval call binding the contract method 0xbd835c42.
//
// Solidity: function BURN_CALL_REQUEST_METADATA_TYPE() view returns(uint8)
func (_Vault *VaultSession) BURNCALLREQUESTMETADATATYPE() (uint8, error) {
	return _Vault.Contract.BURNCALLREQUESTMETADATATYPE(&_Vault.CallOpts)
}

// BURNCALLREQUESTMETADATATYPE is a free data retrieval call binding the contract method 0xbd835c42.
//
// Solidity: function BURN_CALL_REQUEST_METADATA_TYPE() view returns(uint8)
func (_Vault *VaultCallerSession) BURNCALLREQUESTMETADATATYPE() (uint8, error) {
	return _Vault.Contract.BURNCALLREQUESTMETADATATYPE(&_Vault.CallOpts)
}

// BURNREQUESTMETADATATYPE is a free data retrieval call binding the contract method 0x568c04fd.
//
// Solidity: function BURN_REQUEST_METADATA_TYPE() view returns(uint8)
func (_Vault *VaultCaller) BURNREQUESTMETADATATYPE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "BURN_REQUEST_METADATA_TYPE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BURNREQUESTMETADATATYPE is a free data retrieval call binding the contract method 0x568c04fd.
//
// Solidity: function BURN_REQUEST_METADATA_TYPE() view returns(uint8)
func (_Vault *VaultSession) BURNREQUESTMETADATATYPE() (uint8, error) {
	return _Vault.Contract.BURNREQUESTMETADATATYPE(&_Vault.CallOpts)
}

// BURNREQUESTMETADATATYPE is a free data retrieval call binding the contract method 0x568c04fd.
//
// Solidity: function BURN_REQUEST_METADATA_TYPE() view returns(uint8)
func (_Vault *VaultCallerSession) BURNREQUESTMETADATATYPE() (uint8, error) {
	return _Vault.Contract.BURNREQUESTMETADATATYPE(&_Vault.CallOpts)
}

// BURNTOCONTRACTREQUESTMETADATATYPE is a free data retrieval call binding the contract method 0x6f2cbc48.
//
// Solidity: function BURN_TO_CONTRACT_REQUEST_METADATA_TYPE() view returns(uint8)
func (_Vault *VaultCaller) BURNTOCONTRACTREQUESTMETADATATYPE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "BURN_TO_CONTRACT_REQUEST_METADATA_TYPE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BURNTOCONTRACTREQUESTMETADATATYPE is a free data retrieval call binding the contract method 0x6f2cbc48.
//
// Solidity: function BURN_TO_CONTRACT_REQUEST_METADATA_TYPE() view returns(uint8)
func (_Vault *VaultSession) BURNTOCONTRACTREQUESTMETADATATYPE() (uint8, error) {
	return _Vault.Contract.BURNTOCONTRACTREQUESTMETADATATYPE(&_Vault.CallOpts)
}

// BURNTOCONTRACTREQUESTMETADATATYPE is a free data retrieval call binding the contract method 0x6f2cbc48.
//
// Solidity: function BURN_TO_CONTRACT_REQUEST_METADATA_TYPE() view returns(uint8)
func (_Vault *VaultCallerSession) BURNTOCONTRACTREQUESTMETADATATYPE() (uint8, error) {
	return _Vault.Contract.BURNTOCONTRACTREQUESTMETADATATYPE(&_Vault.CallOpts)
}

// CURRENTNETWORKID is a free data retrieval call binding the contract method 0xd7200eb1.
//
// Solidity: function CURRENT_NETWORK_ID() view returns(uint8)
func (_Vault *VaultCaller) CURRENTNETWORKID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "CURRENT_NETWORK_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CURRENTNETWORKID is a free data retrieval call binding the contract method 0xd7200eb1.
//
// Solidity: function CURRENT_NETWORK_ID() view returns(uint8)
func (_Vault *VaultSession) CURRENTNETWORKID() (uint8, error) {
	return _Vault.Contract.CURRENTNETWORKID(&_Vault.CallOpts)
}

// CURRENTNETWORKID is a free data retrieval call binding the contract method 0xd7200eb1.
//
// Solidity: function CURRENT_NETWORK_ID() view returns(uint8)
func (_Vault *VaultCallerSession) CURRENTNETWORKID() (uint8, error) {
	return _Vault.Contract.CURRENTNETWORKID(&_Vault.CallOpts)
}

// ETHTOKEN is a free data retrieval call binding the contract method 0x58bc8337.
//
// Solidity: function ETH_TOKEN() view returns(address)
func (_Vault *VaultCaller) ETHTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "ETH_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ETHTOKEN is a free data retrieval call binding the contract method 0x58bc8337.
//
// Solidity: function ETH_TOKEN() view returns(address)
func (_Vault *VaultSession) ETHTOKEN() (common.Address, error) {
	return _Vault.Contract.ETHTOKEN(&_Vault.CallOpts)
}

// ETHTOKEN is a free data retrieval call binding the contract method 0x58bc8337.
//
// Solidity: function ETH_TOKEN() view returns(address)
func (_Vault *VaultCallerSession) ETHTOKEN() (common.Address, error) {
	return _Vault.Contract.ETHTOKEN(&_Vault.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address token) view returns(uint256)
func (_Vault *VaultCaller) BalanceOf(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "balanceOf", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address token) view returns(uint256)
func (_Vault *VaultSession) BalanceOf(token common.Address) (*big.Int, error) {
	return _Vault.Contract.BalanceOf(&_Vault.CallOpts, token)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address token) view returns(uint256)
func (_Vault *VaultCallerSession) BalanceOf(token common.Address) (*big.Int, error) {
	return _Vault.Contract.BalanceOf(&_Vault.CallOpts, token)
}

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_Vault *VaultCaller) Executor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "executor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_Vault *VaultSession) Executor() (common.Address, error) {
	return _Vault.Contract.Executor(&_Vault.CallOpts)
}

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_Vault *VaultCallerSession) Executor() (common.Address, error) {
	return _Vault.Contract.Executor(&_Vault.CallOpts)
}

// GetDecimals is a free data retrieval call binding the contract method 0xcf54aaa0.
//
// Solidity: function getDecimals(address token) view returns(uint8)
func (_Vault *VaultCaller) GetDecimals(opts *bind.CallOpts, token common.Address) (uint8, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getDecimals", token)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetDecimals is a free data retrieval call binding the contract method 0xcf54aaa0.
//
// Solidity: function getDecimals(address token) view returns(uint8)
func (_Vault *VaultSession) GetDecimals(token common.Address) (uint8, error) {
	return _Vault.Contract.GetDecimals(&_Vault.CallOpts, token)
}

// GetDecimals is a free data retrieval call binding the contract method 0xcf54aaa0.
//
// Solidity: function getDecimals(address token) view returns(uint8)
func (_Vault *VaultCallerSession) GetDecimals(token common.Address) (uint8, error) {
	return _Vault.Contract.GetDecimals(&_Vault.CallOpts, token)
}

// GetDepositedBalance is a free data retrieval call binding the contract method 0xf75b98ce.
//
// Solidity: function getDepositedBalance(address token, address owner) view returns(uint256)
func (_Vault *VaultCaller) GetDepositedBalance(opts *bind.CallOpts, token common.Address, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getDepositedBalance", token, owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDepositedBalance is a free data retrieval call binding the contract method 0xf75b98ce.
//
// Solidity: function getDepositedBalance(address token, address owner) view returns(uint256)
func (_Vault *VaultSession) GetDepositedBalance(token common.Address, owner common.Address) (*big.Int, error) {
	return _Vault.Contract.GetDepositedBalance(&_Vault.CallOpts, token, owner)
}

// GetDepositedBalance is a free data retrieval call binding the contract method 0xf75b98ce.
//
// Solidity: function getDepositedBalance(address token, address owner) view returns(uint256)
func (_Vault *VaultCallerSession) GetDepositedBalance(token common.Address, owner common.Address) (*big.Int, error) {
	return _Vault.Contract.GetDepositedBalance(&_Vault.CallOpts, token, owner)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_Vault *VaultCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_Vault *VaultSession) IsInitialized() (bool, error) {
	return _Vault.Contract.IsInitialized(&_Vault.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_Vault *VaultCallerSession) IsInitialized() (bool, error) {
	return _Vault.Contract.IsInitialized(&_Vault.CallOpts)
}

// IsSigDataUsed is a free data retrieval call binding the contract method 0xe4bd7074.
//
// Solidity: function isSigDataUsed(bytes32 hash) view returns(bool)
func (_Vault *VaultCaller) IsSigDataUsed(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "isSigDataUsed", hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSigDataUsed is a free data retrieval call binding the contract method 0xe4bd7074.
//
// Solidity: function isSigDataUsed(bytes32 hash) view returns(bool)
func (_Vault *VaultSession) IsSigDataUsed(hash [32]byte) (bool, error) {
	return _Vault.Contract.IsSigDataUsed(&_Vault.CallOpts, hash)
}

// IsSigDataUsed is a free data retrieval call binding the contract method 0xe4bd7074.
//
// Solidity: function isSigDataUsed(bytes32 hash) view returns(bool)
func (_Vault *VaultCallerSession) IsSigDataUsed(hash [32]byte) (bool, error) {
	return _Vault.Contract.IsSigDataUsed(&_Vault.CallOpts, hash)
}

// IsWithdrawed is a free data retrieval call binding the contract method 0x749c5f86.
//
// Solidity: function isWithdrawed(bytes32 hash) view returns(bool)
func (_Vault *VaultCaller) IsWithdrawed(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "isWithdrawed", hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWithdrawed is a free data retrieval call binding the contract method 0x749c5f86.
//
// Solidity: function isWithdrawed(bytes32 hash) view returns(bool)
func (_Vault *VaultSession) IsWithdrawed(hash [32]byte) (bool, error) {
	return _Vault.Contract.IsWithdrawed(&_Vault.CallOpts, hash)
}

// IsWithdrawed is a free data retrieval call binding the contract method 0x749c5f86.
//
// Solidity: function isWithdrawed(bytes32 hash) view returns(bool)
func (_Vault *VaultCallerSession) IsWithdrawed(hash [32]byte) (bool, error) {
	return _Vault.Contract.IsWithdrawed(&_Vault.CallOpts, hash)
}

// Migration is a free data retrieval call binding the contract method 0x995fac11.
//
// Solidity: function migration(address , address ) view returns(bool)
func (_Vault *VaultCaller) Migration(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "migration", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Migration is a free data retrieval call binding the contract method 0x995fac11.
//
// Solidity: function migration(address , address ) view returns(bool)
func (_Vault *VaultSession) Migration(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Vault.Contract.Migration(&_Vault.CallOpts, arg0, arg1)
}

// Migration is a free data retrieval call binding the contract method 0x995fac11.
//
// Solidity: function migration(address , address ) view returns(bool)
func (_Vault *VaultCallerSession) Migration(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Vault.Contract.Migration(&_Vault.CallOpts, arg0, arg1)
}

// NotEntered is a free data retrieval call binding the contract method 0xa3f5d8cc.
//
// Solidity: function notEntered() view returns(bool)
func (_Vault *VaultCaller) NotEntered(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "notEntered")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NotEntered is a free data retrieval call binding the contract method 0xa3f5d8cc.
//
// Solidity: function notEntered() view returns(bool)
func (_Vault *VaultSession) NotEntered() (bool, error) {
	return _Vault.Contract.NotEntered(&_Vault.CallOpts)
}

// NotEntered is a free data retrieval call binding the contract method 0xa3f5d8cc.
//
// Solidity: function notEntered() view returns(bool)
func (_Vault *VaultCallerSession) NotEntered() (bool, error) {
	return _Vault.Contract.NotEntered(&_Vault.CallOpts)
}

// ParseBurnInst is a free data retrieval call binding the contract method 0x7e16e6e1.
//
// Solidity: function parseBurnInst(bytes inst) pure returns((uint8,uint8,address,address,uint256,bytes32))
func (_Vault *VaultCaller) ParseBurnInst(opts *bind.CallOpts, inst []byte) (VaultBurnInstData, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "parseBurnInst", inst)

	if err != nil {
		return *new(VaultBurnInstData), err
	}

	out0 := *abi.ConvertType(out[0], new(VaultBurnInstData)).(*VaultBurnInstData)

	return out0, err

}

// ParseBurnInst is a free data retrieval call binding the contract method 0x7e16e6e1.
//
// Solidity: function parseBurnInst(bytes inst) pure returns((uint8,uint8,address,address,uint256,bytes32))
func (_Vault *VaultSession) ParseBurnInst(inst []byte) (VaultBurnInstData, error) {
	return _Vault.Contract.ParseBurnInst(&_Vault.CallOpts, inst)
}

// ParseBurnInst is a free data retrieval call binding the contract method 0x7e16e6e1.
//
// Solidity: function parseBurnInst(bytes inst) pure returns((uint8,uint8,address,address,uint256,bytes32))
func (_Vault *VaultCallerSession) ParseBurnInst(inst []byte) (VaultBurnInstData, error) {
	return _Vault.Contract.ParseBurnInst(&_Vault.CallOpts, inst)
}

// ParseCalldataFromBurnInst is a free data retrieval call binding the contract method 0x66945b31.
//
// Solidity: function parseCalldataFromBurnInst(bytes inst) pure returns((uint8,uint8,address,address,uint256,bytes32), (address,bytes,address), bytes)
func (_Vault *VaultCaller) ParseCalldataFromBurnInst(opts *bind.CallOpts, inst []byte) (VaultBurnInstData, VaultRedepositOptions, []byte, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "parseCalldataFromBurnInst", inst)

	if err != nil {
		return *new(VaultBurnInstData), *new(VaultRedepositOptions), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(VaultBurnInstData)).(*VaultBurnInstData)
	out1 := *abi.ConvertType(out[1], new(VaultRedepositOptions)).(*VaultRedepositOptions)
	out2 := *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return out0, out1, out2, err

}

// ParseCalldataFromBurnInst is a free data retrieval call binding the contract method 0x66945b31.
//
// Solidity: function parseCalldataFromBurnInst(bytes inst) pure returns((uint8,uint8,address,address,uint256,bytes32), (address,bytes,address), bytes)
func (_Vault *VaultSession) ParseCalldataFromBurnInst(inst []byte) (VaultBurnInstData, VaultRedepositOptions, []byte, error) {
	return _Vault.Contract.ParseCalldataFromBurnInst(&_Vault.CallOpts, inst)
}

// ParseCalldataFromBurnInst is a free data retrieval call binding the contract method 0x66945b31.
//
// Solidity: function parseCalldataFromBurnInst(bytes inst) pure returns((uint8,uint8,address,address,uint256,bytes32), (address,bytes,address), bytes)
func (_Vault *VaultCallerSession) ParseCalldataFromBurnInst(inst []byte) (VaultBurnInstData, VaultRedepositOptions, []byte, error) {
	return _Vault.Contract.ParseCalldataFromBurnInst(&_Vault.CallOpts, inst)
}

// PrevVault is a free data retrieval call binding the contract method 0xfa84702e.
//
// Solidity: function prevVault() view returns(address)
func (_Vault *VaultCaller) PrevVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "prevVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrevVault is a free data retrieval call binding the contract method 0xfa84702e.
//
// Solidity: function prevVault() view returns(address)
func (_Vault *VaultSession) PrevVault() (common.Address, error) {
	return _Vault.Contract.PrevVault(&_Vault.CallOpts)
}

// PrevVault is a free data retrieval call binding the contract method 0xfa84702e.
//
// Solidity: function prevVault() view returns(address)
func (_Vault *VaultCallerSession) PrevVault() (common.Address, error) {
	return _Vault.Contract.PrevVault(&_Vault.CallOpts)
}

// Regulator is a free data retrieval call binding the contract method 0xdd8fee14.
//
// Solidity: function regulator() view returns(address)
func (_Vault *VaultCaller) Regulator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "regulator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Regulator is a free data retrieval call binding the contract method 0xdd8fee14.
//
// Solidity: function regulator() view returns(address)
func (_Vault *VaultSession) Regulator() (common.Address, error) {
	return _Vault.Contract.Regulator(&_Vault.CallOpts)
}

// Regulator is a free data retrieval call binding the contract method 0xdd8fee14.
//
// Solidity: function regulator() view returns(address)
func (_Vault *VaultCallerSession) Regulator() (common.Address, error) {
	return _Vault.Contract.Regulator(&_Vault.CallOpts)
}

// SigDataUsed is a free data retrieval call binding the contract method 0x1ea1940e.
//
// Solidity: function sigDataUsed(bytes32 ) view returns(bool)
func (_Vault *VaultCaller) SigDataUsed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "sigDataUsed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SigDataUsed is a free data retrieval call binding the contract method 0x1ea1940e.
//
// Solidity: function sigDataUsed(bytes32 ) view returns(bool)
func (_Vault *VaultSession) SigDataUsed(arg0 [32]byte) (bool, error) {
	return _Vault.Contract.SigDataUsed(&_Vault.CallOpts, arg0)
}

// SigDataUsed is a free data retrieval call binding the contract method 0x1ea1940e.
//
// Solidity: function sigDataUsed(bytes32 ) view returns(bool)
func (_Vault *VaultCallerSession) SigDataUsed(arg0 [32]byte) (bool, error) {
	return _Vault.Contract.SigDataUsed(&_Vault.CallOpts, arg0)
}

// SigToAddress is a free data retrieval call binding the contract method 0x3fec6b40.
//
// Solidity: function sigToAddress(bytes signData, bytes32 hash) pure returns(address)
func (_Vault *VaultCaller) SigToAddress(opts *bind.CallOpts, signData []byte, hash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "sigToAddress", signData, hash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SigToAddress is a free data retrieval call binding the contract method 0x3fec6b40.
//
// Solidity: function sigToAddress(bytes signData, bytes32 hash) pure returns(address)
func (_Vault *VaultSession) SigToAddress(signData []byte, hash [32]byte) (common.Address, error) {
	return _Vault.Contract.SigToAddress(&_Vault.CallOpts, signData, hash)
}

// SigToAddress is a free data retrieval call binding the contract method 0x3fec6b40.
//
// Solidity: function sigToAddress(bytes signData, bytes32 hash) pure returns(address)
func (_Vault *VaultCallerSession) SigToAddress(signData []byte, hash [32]byte) (common.Address, error) {
	return _Vault.Contract.SigToAddress(&_Vault.CallOpts, signData, hash)
}

// StorageLayoutVersion is a free data retrieval call binding the contract method 0xd6a1fe3b.
//
// Solidity: function storageLayoutVersion() view returns(uint256)
func (_Vault *VaultCaller) StorageLayoutVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "storageLayoutVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StorageLayoutVersion is a free data retrieval call binding the contract method 0xd6a1fe3b.
//
// Solidity: function storageLayoutVersion() view returns(uint256)
func (_Vault *VaultSession) StorageLayoutVersion() (*big.Int, error) {
	return _Vault.Contract.StorageLayoutVersion(&_Vault.CallOpts)
}

// StorageLayoutVersion is a free data retrieval call binding the contract method 0xd6a1fe3b.
//
// Solidity: function storageLayoutVersion() view returns(uint256)
func (_Vault *VaultCallerSession) StorageLayoutVersion() (*big.Int, error) {
	return _Vault.Contract.StorageLayoutVersion(&_Vault.CallOpts)
}

// TotalDepositedToSCAmount is a free data retrieval call binding the contract method 0x6304541c.
//
// Solidity: function totalDepositedToSCAmount(address ) view returns(uint256)
func (_Vault *VaultCaller) TotalDepositedToSCAmount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "totalDepositedToSCAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDepositedToSCAmount is a free data retrieval call binding the contract method 0x6304541c.
//
// Solidity: function totalDepositedToSCAmount(address ) view returns(uint256)
func (_Vault *VaultSession) TotalDepositedToSCAmount(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.TotalDepositedToSCAmount(&_Vault.CallOpts, arg0)
}

// TotalDepositedToSCAmount is a free data retrieval call binding the contract method 0x6304541c.
//
// Solidity: function totalDepositedToSCAmount(address ) view returns(uint256)
func (_Vault *VaultCallerSession) TotalDepositedToSCAmount(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.TotalDepositedToSCAmount(&_Vault.CallOpts, arg0)
}

// WithdrawRequests is a free data retrieval call binding the contract method 0x65b5a00f.
//
// Solidity: function withdrawRequests(address , address ) view returns(uint256)
func (_Vault *VaultCaller) WithdrawRequests(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "withdrawRequests", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawRequests is a free data retrieval call binding the contract method 0x65b5a00f.
//
// Solidity: function withdrawRequests(address , address ) view returns(uint256)
func (_Vault *VaultSession) WithdrawRequests(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Vault.Contract.WithdrawRequests(&_Vault.CallOpts, arg0, arg1)
}

// WithdrawRequests is a free data retrieval call binding the contract method 0x65b5a00f.
//
// Solidity: function withdrawRequests(address , address ) view returns(uint256)
func (_Vault *VaultCallerSession) WithdrawRequests(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Vault.Contract.WithdrawRequests(&_Vault.CallOpts, arg0, arg1)
}

// Withdrawed is a free data retrieval call binding the contract method 0xdca40d9e.
//
// Solidity: function withdrawed(bytes32 ) view returns(bool)
func (_Vault *VaultCaller) Withdrawed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "withdrawed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Withdrawed is a free data retrieval call binding the contract method 0xdca40d9e.
//
// Solidity: function withdrawed(bytes32 ) view returns(bool)
func (_Vault *VaultSession) Withdrawed(arg0 [32]byte) (bool, error) {
	return _Vault.Contract.Withdrawed(&_Vault.CallOpts, arg0)
}

// Withdrawed is a free data retrieval call binding the contract method 0xdca40d9e.
//
// Solidity: function withdrawed(bytes32 ) view returns(bool)
func (_Vault *VaultCallerSession) Withdrawed(arg0 [32]byte) (bool, error) {
	return _Vault.Contract.Withdrawed(&_Vault.CallOpts, arg0)
}

// CallExternal is a paid mutator transaction binding the contract method 0xbda9b509.
//
// Solidity: function _callExternal(address token, address to, uint256 amount, bytes externalCalldata, address redepositToken) returns(uint256)
func (_Vault *VaultTransactor) CallExternal(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int, externalCalldata []byte, redepositToken common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "_callExternal", token, to, amount, externalCalldata, redepositToken)
}

// CallExternal is a paid mutator transaction binding the contract method 0xbda9b509.
//
// Solidity: function _callExternal(address token, address to, uint256 amount, bytes externalCalldata, address redepositToken) returns(uint256)
func (_Vault *VaultSession) CallExternal(token common.Address, to common.Address, amount *big.Int, externalCalldata []byte, redepositToken common.Address) (*types.Transaction, error) {
	return _Vault.Contract.CallExternal(&_Vault.TransactOpts, token, to, amount, externalCalldata, redepositToken)
}

// CallExternal is a paid mutator transaction binding the contract method 0xbda9b509.
//
// Solidity: function _callExternal(address token, address to, uint256 amount, bytes externalCalldata, address redepositToken) returns(uint256)
func (_Vault *VaultTransactorSession) CallExternal(token common.Address, to common.Address, amount *big.Int, externalCalldata []byte, redepositToken common.Address) (*types.Transaction, error) {
	return _Vault.Contract.CallExternal(&_Vault.TransactOpts, token, to, amount, externalCalldata, redepositToken)
}

// TransferExternal is a paid mutator transaction binding the contract method 0x145e2a6b.
//
// Solidity: function _transferExternal(address token, address to, uint256 amount) returns()
func (_Vault *VaultTransactor) TransferExternal(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "_transferExternal", token, to, amount)
}

// TransferExternal is a paid mutator transaction binding the contract method 0x145e2a6b.
//
// Solidity: function _transferExternal(address token, address to, uint256 amount) returns()
func (_Vault *VaultSession) TransferExternal(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.TransferExternal(&_Vault.TransactOpts, token, to, amount)
}

// TransferExternal is a paid mutator transaction binding the contract method 0x145e2a6b.
//
// Solidity: function _transferExternal(address token, address to, uint256 amount) returns()
func (_Vault *VaultTransactorSession) TransferExternal(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.TransferExternal(&_Vault.TransactOpts, token, to, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xc791d705.
//
// Solidity: function deposit(string incognitoAddress, bytes32 txId, bytes signData) payable returns()
func (_Vault *VaultTransactor) Deposit(opts *bind.TransactOpts, incognitoAddress string, txId [32]byte, signData []byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "deposit", incognitoAddress, txId, signData)
}

// Deposit is a paid mutator transaction binding the contract method 0xc791d705.
//
// Solidity: function deposit(string incognitoAddress, bytes32 txId, bytes signData) payable returns()
func (_Vault *VaultSession) Deposit(incognitoAddress string, txId [32]byte, signData []byte) (*types.Transaction, error) {
	return _Vault.Contract.Deposit(&_Vault.TransactOpts, incognitoAddress, txId, signData)
}

// Deposit is a paid mutator transaction binding the contract method 0xc791d705.
//
// Solidity: function deposit(string incognitoAddress, bytes32 txId, bytes signData) payable returns()
func (_Vault *VaultTransactorSession) Deposit(incognitoAddress string, txId [32]byte, signData []byte) (*types.Transaction, error) {
	return _Vault.Contract.Deposit(&_Vault.TransactOpts, incognitoAddress, txId, signData)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xa807b5bb.
//
// Solidity: function depositERC20(address token, uint256 amount, string incognitoAddress, bytes32 txId, bytes signData) returns()
func (_Vault *VaultTransactor) DepositERC20(opts *bind.TransactOpts, token common.Address, amount *big.Int, incognitoAddress string, txId [32]byte, signData []byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "depositERC20", token, amount, incognitoAddress, txId, signData)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xa807b5bb.
//
// Solidity: function depositERC20(address token, uint256 amount, string incognitoAddress, bytes32 txId, bytes signData) returns()
func (_Vault *VaultSession) DepositERC20(token common.Address, amount *big.Int, incognitoAddress string, txId [32]byte, signData []byte) (*types.Transaction, error) {
	return _Vault.Contract.DepositERC20(&_Vault.TransactOpts, token, amount, incognitoAddress, txId, signData)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xa807b5bb.
//
// Solidity: function depositERC20(address token, uint256 amount, string incognitoAddress, bytes32 txId, bytes signData) returns()
func (_Vault *VaultTransactorSession) DepositERC20(token common.Address, amount *big.Int, incognitoAddress string, txId [32]byte, signData []byte) (*types.Transaction, error) {
	return _Vault.Contract.DepositERC20(&_Vault.TransactOpts, token, amount, incognitoAddress, txId, signData)
}

// DepositERC20V2 is a paid mutator transaction binding the contract method 0xb8237dbb.
//
// Solidity: function depositERC20_V2(address token, uint256 amount, string incognitoAddress, bytes32 txId, bytes signData) returns()
func (_Vault *VaultTransactor) DepositERC20V2(opts *bind.TransactOpts, token common.Address, amount *big.Int, incognitoAddress string, txId [32]byte, signData []byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "depositERC20_V2", token, amount, incognitoAddress, txId, signData)
}

// DepositERC20V2 is a paid mutator transaction binding the contract method 0xb8237dbb.
//
// Solidity: function depositERC20_V2(address token, uint256 amount, string incognitoAddress, bytes32 txId, bytes signData) returns()
func (_Vault *VaultSession) DepositERC20V2(token common.Address, amount *big.Int, incognitoAddress string, txId [32]byte, signData []byte) (*types.Transaction, error) {
	return _Vault.Contract.DepositERC20V2(&_Vault.TransactOpts, token, amount, incognitoAddress, txId, signData)
}

// DepositERC20V2 is a paid mutator transaction binding the contract method 0xb8237dbb.
//
// Solidity: function depositERC20_V2(address token, uint256 amount, string incognitoAddress, bytes32 txId, bytes signData) returns()
func (_Vault *VaultTransactorSession) DepositERC20V2(token common.Address, amount *big.Int, incognitoAddress string, txId [32]byte, signData []byte) (*types.Transaction, error) {
	return _Vault.Contract.DepositERC20V2(&_Vault.TransactOpts, token, amount, incognitoAddress, txId, signData)
}

// DepositV2 is a paid mutator transaction binding the contract method 0x84b3ac03.
//
// Solidity: function deposit_V2(string incognitoAddress, bytes32 txId, bytes signData) payable returns()
func (_Vault *VaultTransactor) DepositV2(opts *bind.TransactOpts, incognitoAddress string, txId [32]byte, signData []byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "deposit_V2", incognitoAddress, txId, signData)
}

// DepositV2 is a paid mutator transaction binding the contract method 0x84b3ac03.
//
// Solidity: function deposit_V2(string incognitoAddress, bytes32 txId, bytes signData) payable returns()
func (_Vault *VaultSession) DepositV2(incognitoAddress string, txId [32]byte, signData []byte) (*types.Transaction, error) {
	return _Vault.Contract.DepositV2(&_Vault.TransactOpts, incognitoAddress, txId, signData)
}

// DepositV2 is a paid mutator transaction binding the contract method 0x84b3ac03.
//
// Solidity: function deposit_V2(string incognitoAddress, bytes32 txId, bytes signData) payable returns()
func (_Vault *VaultTransactorSession) DepositV2(incognitoAddress string, txId [32]byte, signData []byte) (*types.Transaction, error) {
	return _Vault.Contract.DepositV2(&_Vault.TransactOpts, incognitoAddress, txId, signData)
}

// Execute is a paid mutator transaction binding the contract method 0x8588ccd6.
//
// Solidity: function execute(address token, uint256 amount, address recipientToken, address exchangeAddress, bytes callData, bytes timestamp, bytes signData) payable returns()
func (_Vault *VaultTransactor) Execute(opts *bind.TransactOpts, token common.Address, amount *big.Int, recipientToken common.Address, exchangeAddress common.Address, callData []byte, timestamp []byte, signData []byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "execute", token, amount, recipientToken, exchangeAddress, callData, timestamp, signData)
}

// Execute is a paid mutator transaction binding the contract method 0x8588ccd6.
//
// Solidity: function execute(address token, uint256 amount, address recipientToken, address exchangeAddress, bytes callData, bytes timestamp, bytes signData) payable returns()
func (_Vault *VaultSession) Execute(token common.Address, amount *big.Int, recipientToken common.Address, exchangeAddress common.Address, callData []byte, timestamp []byte, signData []byte) (*types.Transaction, error) {
	return _Vault.Contract.Execute(&_Vault.TransactOpts, token, amount, recipientToken, exchangeAddress, callData, timestamp, signData)
}

// Execute is a paid mutator transaction binding the contract method 0x8588ccd6.
//
// Solidity: function execute(address token, uint256 amount, address recipientToken, address exchangeAddress, bytes callData, bytes timestamp, bytes signData) payable returns()
func (_Vault *VaultTransactorSession) Execute(token common.Address, amount *big.Int, recipientToken common.Address, exchangeAddress common.Address, callData []byte, timestamp []byte, signData []byte) (*types.Transaction, error) {
	return _Vault.Contract.Execute(&_Vault.TransactOpts, token, amount, recipientToken, exchangeAddress, callData, timestamp, signData)
}

// ExecuteWithBurnProof is a paid mutator transaction binding the contract method 0x3ed1b376.
//
// Solidity: function executeWithBurnProof(bytes inst, uint256 heights, bytes32[] instPaths, bool[] instPathIsLefts, bytes32 instRoots, bytes32 blkData, uint256[] sigIdxs, uint8[] sigVs, bytes32[] sigRs, bytes32[] sigSs) returns()
func (_Vault *VaultTransactor) ExecuteWithBurnProof(opts *bind.TransactOpts, inst []byte, heights *big.Int, instPaths [][32]byte, instPathIsLefts []bool, instRoots [32]byte, blkData [32]byte, sigIdxs []*big.Int, sigVs []uint8, sigRs [][32]byte, sigSs [][32]byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "executeWithBurnProof", inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// ExecuteWithBurnProof is a paid mutator transaction binding the contract method 0x3ed1b376.
//
// Solidity: function executeWithBurnProof(bytes inst, uint256 heights, bytes32[] instPaths, bool[] instPathIsLefts, bytes32 instRoots, bytes32 blkData, uint256[] sigIdxs, uint8[] sigVs, bytes32[] sigRs, bytes32[] sigSs) returns()
func (_Vault *VaultSession) ExecuteWithBurnProof(inst []byte, heights *big.Int, instPaths [][32]byte, instPathIsLefts []bool, instRoots [32]byte, blkData [32]byte, sigIdxs []*big.Int, sigVs []uint8, sigRs [][32]byte, sigSs [][32]byte) (*types.Transaction, error) {
	return _Vault.Contract.ExecuteWithBurnProof(&_Vault.TransactOpts, inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// ExecuteWithBurnProof is a paid mutator transaction binding the contract method 0x3ed1b376.
//
// Solidity: function executeWithBurnProof(bytes inst, uint256 heights, bytes32[] instPaths, bool[] instPathIsLefts, bytes32 instRoots, bytes32 blkData, uint256[] sigIdxs, uint8[] sigVs, bytes32[] sigRs, bytes32[] sigSs) returns()
func (_Vault *VaultTransactorSession) ExecuteWithBurnProof(inst []byte, heights *big.Int, instPaths [][32]byte, instPathIsLefts []bool, instRoots [32]byte, blkData [32]byte, sigIdxs []*big.Int, sigVs []uint8, sigRs [][32]byte, sigSs [][32]byte) (*types.Transaction, error) {
	return _Vault.Contract.ExecuteWithBurnProof(&_Vault.TransactOpts, inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _prevVault, address _regulator, address _executor) returns()
func (_Vault *VaultTransactor) Initialize(opts *bind.TransactOpts, _prevVault common.Address, _regulator common.Address, _executor common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "initialize", _prevVault, _regulator, _executor)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _prevVault, address _regulator, address _executor) returns()
func (_Vault *VaultSession) Initialize(_prevVault common.Address, _regulator common.Address, _executor common.Address) (*types.Transaction, error) {
	return _Vault.Contract.Initialize(&_Vault.TransactOpts, _prevVault, _regulator, _executor)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _prevVault, address _regulator, address _executor) returns()
func (_Vault *VaultTransactorSession) Initialize(_prevVault common.Address, _regulator common.Address, _executor common.Address) (*types.Transaction, error) {
	return _Vault.Contract.Initialize(&_Vault.TransactOpts, _prevVault, _regulator, _executor)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0xfee8efda.
//
// Solidity: function requestWithdraw(string incognitoAddress, address token, uint256 amount, bytes signData, bytes timestamp, bytes32 txId, bytes regulatorSig) returns()
func (_Vault *VaultTransactor) RequestWithdraw(opts *bind.TransactOpts, incognitoAddress string, token common.Address, amount *big.Int, signData []byte, timestamp []byte, txId [32]byte, regulatorSig []byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "requestWithdraw", incognitoAddress, token, amount, signData, timestamp, txId, regulatorSig)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0xfee8efda.
//
// Solidity: function requestWithdraw(string incognitoAddress, address token, uint256 amount, bytes signData, bytes timestamp, bytes32 txId, bytes regulatorSig) returns()
func (_Vault *VaultSession) RequestWithdraw(incognitoAddress string, token common.Address, amount *big.Int, signData []byte, timestamp []byte, txId [32]byte, regulatorSig []byte) (*types.Transaction, error) {
	return _Vault.Contract.RequestWithdraw(&_Vault.TransactOpts, incognitoAddress, token, amount, signData, timestamp, txId, regulatorSig)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0xfee8efda.
//
// Solidity: function requestWithdraw(string incognitoAddress, address token, uint256 amount, bytes signData, bytes timestamp, bytes32 txId, bytes regulatorSig) returns()
func (_Vault *VaultTransactorSession) RequestWithdraw(incognitoAddress string, token common.Address, amount *big.Int, signData []byte, timestamp []byte, txId [32]byte, regulatorSig []byte) (*types.Transaction, error) {
	return _Vault.Contract.RequestWithdraw(&_Vault.TransactOpts, incognitoAddress, token, amount, signData, timestamp, txId, regulatorSig)
}

// SetRegulator is a paid mutator transaction binding the contract method 0xcde0a4f8.
//
// Solidity: function setRegulator(address _regulator) returns()
func (_Vault *VaultTransactor) SetRegulator(opts *bind.TransactOpts, _regulator common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setRegulator", _regulator)
}

// SetRegulator is a paid mutator transaction binding the contract method 0xcde0a4f8.
//
// Solidity: function setRegulator(address _regulator) returns()
func (_Vault *VaultSession) SetRegulator(_regulator common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetRegulator(&_Vault.TransactOpts, _regulator)
}

// SetRegulator is a paid mutator transaction binding the contract method 0xcde0a4f8.
//
// Solidity: function setRegulator(address _regulator) returns()
func (_Vault *VaultTransactorSession) SetRegulator(_regulator common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetRegulator(&_Vault.TransactOpts, _regulator)
}

// SubmitBurnProof is a paid mutator transaction binding the contract method 0x73bf9651.
//
// Solidity: function submitBurnProof(bytes inst, uint256 heights, bytes32[] instPaths, bool[] instPathIsLefts, bytes32 instRoots, bytes32 blkData, uint256[] sigIdxs, uint8[] sigVs, bytes32[] sigRs, bytes32[] sigSs) returns()
func (_Vault *VaultTransactor) SubmitBurnProof(opts *bind.TransactOpts, inst []byte, heights *big.Int, instPaths [][32]byte, instPathIsLefts []bool, instRoots [32]byte, blkData [32]byte, sigIdxs []*big.Int, sigVs []uint8, sigRs [][32]byte, sigSs [][32]byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "submitBurnProof", inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// SubmitBurnProof is a paid mutator transaction binding the contract method 0x73bf9651.
//
// Solidity: function submitBurnProof(bytes inst, uint256 heights, bytes32[] instPaths, bool[] instPathIsLefts, bytes32 instRoots, bytes32 blkData, uint256[] sigIdxs, uint8[] sigVs, bytes32[] sigRs, bytes32[] sigSs) returns()
func (_Vault *VaultSession) SubmitBurnProof(inst []byte, heights *big.Int, instPaths [][32]byte, instPathIsLefts []bool, instRoots [32]byte, blkData [32]byte, sigIdxs []*big.Int, sigVs []uint8, sigRs [][32]byte, sigSs [][32]byte) (*types.Transaction, error) {
	return _Vault.Contract.SubmitBurnProof(&_Vault.TransactOpts, inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// SubmitBurnProof is a paid mutator transaction binding the contract method 0x73bf9651.
//
// Solidity: function submitBurnProof(bytes inst, uint256 heights, bytes32[] instPaths, bool[] instPathIsLefts, bytes32 instRoots, bytes32 blkData, uint256[] sigIdxs, uint8[] sigVs, bytes32[] sigRs, bytes32[] sigSs) returns()
func (_Vault *VaultTransactorSession) SubmitBurnProof(inst []byte, heights *big.Int, instPaths [][32]byte, instPathIsLefts []bool, instRoots [32]byte, blkData [32]byte, sigIdxs []*big.Int, sigVs []uint8, sigRs [][32]byte, sigSs [][32]byte) (*types.Transaction, error) {
	return _Vault.Contract.SubmitBurnProof(&_Vault.TransactOpts, inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// UpdateAssets is a paid mutator transaction binding the contract method 0x1ed4276d.
//
// Solidity: function updateAssets(address[] assets, uint256[] amounts) returns(bool)
func (_Vault *VaultTransactor) UpdateAssets(opts *bind.TransactOpts, assets []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "updateAssets", assets, amounts)
}

// UpdateAssets is a paid mutator transaction binding the contract method 0x1ed4276d.
//
// Solidity: function updateAssets(address[] assets, uint256[] amounts) returns(bool)
func (_Vault *VaultSession) UpdateAssets(assets []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Vault.Contract.UpdateAssets(&_Vault.TransactOpts, assets, amounts)
}

// UpdateAssets is a paid mutator transaction binding the contract method 0x1ed4276d.
//
// Solidity: function updateAssets(address[] assets, uint256[] amounts) returns(bool)
func (_Vault *VaultTransactorSession) UpdateAssets(assets []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Vault.Contract.UpdateAssets(&_Vault.TransactOpts, assets, amounts)
}

// UpgradeVaultStorage is a paid mutator transaction binding the contract method 0xa73b1532.
//
// Solidity: function upgradeVaultStorage(address _regulator, address _executor) returns()
func (_Vault *VaultTransactor) UpgradeVaultStorage(opts *bind.TransactOpts, _regulator common.Address, _executor common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "upgradeVaultStorage", _regulator, _executor)
}

// UpgradeVaultStorage is a paid mutator transaction binding the contract method 0xa73b1532.
//
// Solidity: function upgradeVaultStorage(address _regulator, address _executor) returns()
func (_Vault *VaultSession) UpgradeVaultStorage(_regulator common.Address, _executor common.Address) (*types.Transaction, error) {
	return _Vault.Contract.UpgradeVaultStorage(&_Vault.TransactOpts, _regulator, _executor)
}

// UpgradeVaultStorage is a paid mutator transaction binding the contract method 0xa73b1532.
//
// Solidity: function upgradeVaultStorage(address _regulator, address _executor) returns()
func (_Vault *VaultTransactorSession) UpgradeVaultStorage(_regulator common.Address, _executor common.Address) (*types.Transaction, error) {
	return _Vault.Contract.UpgradeVaultStorage(&_Vault.TransactOpts, _regulator, _executor)
}

// Withdraw is a paid mutator transaction binding the contract method 0x1beb7de2.
//
// Solidity: function withdraw(bytes inst, uint256 heights, bytes32[] instPaths, bool[] instPathIsLefts, bytes32 instRoots, bytes32 blkData, uint256[] sigIdxs, uint8[] sigVs, bytes32[] sigRs, bytes32[] sigSs) returns()
func (_Vault *VaultTransactor) Withdraw(opts *bind.TransactOpts, inst []byte, heights *big.Int, instPaths [][32]byte, instPathIsLefts []bool, instRoots [32]byte, blkData [32]byte, sigIdxs []*big.Int, sigVs []uint8, sigRs [][32]byte, sigSs [][32]byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "withdraw", inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// Withdraw is a paid mutator transaction binding the contract method 0x1beb7de2.
//
// Solidity: function withdraw(bytes inst, uint256 heights, bytes32[] instPaths, bool[] instPathIsLefts, bytes32 instRoots, bytes32 blkData, uint256[] sigIdxs, uint8[] sigVs, bytes32[] sigRs, bytes32[] sigSs) returns()
func (_Vault *VaultSession) Withdraw(inst []byte, heights *big.Int, instPaths [][32]byte, instPathIsLefts []bool, instRoots [32]byte, blkData [32]byte, sigIdxs []*big.Int, sigVs []uint8, sigRs [][32]byte, sigSs [][32]byte) (*types.Transaction, error) {
	return _Vault.Contract.Withdraw(&_Vault.TransactOpts, inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// Withdraw is a paid mutator transaction binding the contract method 0x1beb7de2.
//
// Solidity: function withdraw(bytes inst, uint256 heights, bytes32[] instPaths, bool[] instPathIsLefts, bytes32 instRoots, bytes32 blkData, uint256[] sigIdxs, uint8[] sigVs, bytes32[] sigRs, bytes32[] sigSs) returns()
func (_Vault *VaultTransactorSession) Withdraw(inst []byte, heights *big.Int, instPaths [][32]byte, instPathIsLefts []bool, instRoots [32]byte, blkData [32]byte, sigIdxs []*big.Int, sigVs []uint8, sigRs [][32]byte, sigSs [][32]byte) (*types.Transaction, error) {
	return _Vault.Contract.Withdraw(&_Vault.TransactOpts, inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Vault *VaultTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Vault *VaultSession) Receive() (*types.Transaction, error) {
	return _Vault.Contract.Receive(&_Vault.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Vault *VaultTransactorSession) Receive() (*types.Transaction, error) {
	return _Vault.Contract.Receive(&_Vault.TransactOpts)
}

// VaultDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Vault contract.
type VaultDepositIterator struct {
	Event *VaultDeposit // Event containing the contract specifics and raw log

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
func (it *VaultDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultDeposit)
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
		it.Event = new(VaultDeposit)
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
func (it *VaultDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultDeposit represents a Deposit event raised by the Vault contract.
type VaultDeposit struct {
	Token            common.Address
	IncognitoAddress string
	Amount           *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e.
//
// Solidity: event Deposit(address token, string incognitoAddress, uint256 amount)
func (_Vault *VaultFilterer) FilterDeposit(opts *bind.FilterOpts) (*VaultDepositIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &VaultDepositIterator{contract: _Vault.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e.
//
// Solidity: event Deposit(address token, string incognitoAddress, uint256 amount)
func (_Vault *VaultFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *VaultDeposit) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultDeposit)
				if err := _Vault.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e.
//
// Solidity: event Deposit(address token, string incognitoAddress, uint256 amount)
func (_Vault *VaultFilterer) ParseDeposit(log types.Log) (*VaultDeposit, error) {
	event := new(VaultDeposit)
	if err := _Vault.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultDepositV2Iterator is returned from FilterDepositV2 and is used to iterate over the raw logs and unpacked data for DepositV2 events raised by the Vault contract.
type VaultDepositV2Iterator struct {
	Event *VaultDepositV2 // Event containing the contract specifics and raw log

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
func (it *VaultDepositV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultDepositV2)
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
		it.Event = new(VaultDepositV2)
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
func (it *VaultDepositV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultDepositV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultDepositV2 represents a DepositV2 event raised by the Vault contract.
type VaultDepositV2 struct {
	Token            common.Address
	IncognitoAddress string
	Amount           *big.Int
	DepositID        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDepositV2 is a free log retrieval operation binding the contract event 0xd30df8040a1092415b49422a02dbd8cdd5915a596abcba02cd0f65dd86ab3851.
//
// Solidity: event DepositV2(address token, string incognitoAddress, uint256 amount, uint256 depositID)
func (_Vault *VaultFilterer) FilterDepositV2(opts *bind.FilterOpts) (*VaultDepositV2Iterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "DepositV2")
	if err != nil {
		return nil, err
	}
	return &VaultDepositV2Iterator{contract: _Vault.contract, event: "DepositV2", logs: logs, sub: sub}, nil
}

// WatchDepositV2 is a free log subscription operation binding the contract event 0xd30df8040a1092415b49422a02dbd8cdd5915a596abcba02cd0f65dd86ab3851.
//
// Solidity: event DepositV2(address token, string incognitoAddress, uint256 amount, uint256 depositID)
func (_Vault *VaultFilterer) WatchDepositV2(opts *bind.WatchOpts, sink chan<- *VaultDepositV2) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "DepositV2")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultDepositV2)
				if err := _Vault.contract.UnpackLog(event, "DepositV2", log); err != nil {
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

// ParseDepositV2 is a log parse operation binding the contract event 0xd30df8040a1092415b49422a02dbd8cdd5915a596abcba02cd0f65dd86ab3851.
//
// Solidity: event DepositV2(address token, string incognitoAddress, uint256 amount, uint256 depositID)
func (_Vault *VaultFilterer) ParseDepositV2(log types.Log) (*VaultDepositV2, error) {
	event := new(VaultDepositV2)
	if err := _Vault.contract.UnpackLog(event, "DepositV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultExecuteFnLogIterator is returned from FilterExecuteFnLog and is used to iterate over the raw logs and unpacked data for ExecuteFnLog events raised by the Vault contract.
type VaultExecuteFnLogIterator struct {
	Event *VaultExecuteFnLog // Event containing the contract specifics and raw log

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
func (it *VaultExecuteFnLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultExecuteFnLog)
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
		it.Event = new(VaultExecuteFnLog)
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
func (it *VaultExecuteFnLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultExecuteFnLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultExecuteFnLog represents a ExecuteFnLog event raised by the Vault contract.
type VaultExecuteFnLog struct {
	Id        [32]byte
	PhaseID   *big.Int
	ErrorData []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExecuteFnLog is a free log retrieval operation binding the contract event 0xdbbb883f24557adf486292429863dcfd4ac5d4db168ae94921da8e3d9a95d416.
//
// Solidity: event ExecuteFnLog(bytes32 id, uint256 phaseID, bytes errorData)
func (_Vault *VaultFilterer) FilterExecuteFnLog(opts *bind.FilterOpts) (*VaultExecuteFnLogIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "ExecuteFnLog")
	if err != nil {
		return nil, err
	}
	return &VaultExecuteFnLogIterator{contract: _Vault.contract, event: "ExecuteFnLog", logs: logs, sub: sub}, nil
}

// WatchExecuteFnLog is a free log subscription operation binding the contract event 0xdbbb883f24557adf486292429863dcfd4ac5d4db168ae94921da8e3d9a95d416.
//
// Solidity: event ExecuteFnLog(bytes32 id, uint256 phaseID, bytes errorData)
func (_Vault *VaultFilterer) WatchExecuteFnLog(opts *bind.WatchOpts, sink chan<- *VaultExecuteFnLog) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "ExecuteFnLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultExecuteFnLog)
				if err := _Vault.contract.UnpackLog(event, "ExecuteFnLog", log); err != nil {
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

// ParseExecuteFnLog is a log parse operation binding the contract event 0xdbbb883f24557adf486292429863dcfd4ac5d4db168ae94921da8e3d9a95d416.
//
// Solidity: event ExecuteFnLog(bytes32 id, uint256 phaseID, bytes errorData)
func (_Vault *VaultFilterer) ParseExecuteFnLog(log types.Log) (*VaultExecuteFnLog, error) {
	event := new(VaultExecuteFnLog)
	if err := _Vault.contract.UnpackLog(event, "ExecuteFnLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultRedepositIterator is returned from FilterRedeposit and is used to iterate over the raw logs and unpacked data for Redeposit events raised by the Vault contract.
type VaultRedepositIterator struct {
	Event *VaultRedeposit // Event containing the contract specifics and raw log

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
func (it *VaultRedepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultRedeposit)
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
		it.Event = new(VaultRedeposit)
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
func (it *VaultRedepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultRedepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultRedeposit represents a Redeposit event raised by the Vault contract.
type VaultRedeposit struct {
	Token               common.Address
	RedepositIncAddress []byte
	Amount              *big.Int
	Itx                 [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterRedeposit is a free log retrieval operation binding the contract event 0x00b45d95b5117447e2fafe7f34def913ff3ba220e4b8688acf37ae2328af7a3d.
//
// Solidity: event Redeposit(address token, bytes redepositIncAddress, uint256 amount, bytes32 itx)
func (_Vault *VaultFilterer) FilterRedeposit(opts *bind.FilterOpts) (*VaultRedepositIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Redeposit")
	if err != nil {
		return nil, err
	}
	return &VaultRedepositIterator{contract: _Vault.contract, event: "Redeposit", logs: logs, sub: sub}, nil
}

// WatchRedeposit is a free log subscription operation binding the contract event 0x00b45d95b5117447e2fafe7f34def913ff3ba220e4b8688acf37ae2328af7a3d.
//
// Solidity: event Redeposit(address token, bytes redepositIncAddress, uint256 amount, bytes32 itx)
func (_Vault *VaultFilterer) WatchRedeposit(opts *bind.WatchOpts, sink chan<- *VaultRedeposit) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Redeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultRedeposit)
				if err := _Vault.contract.UnpackLog(event, "Redeposit", log); err != nil {
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

// ParseRedeposit is a log parse operation binding the contract event 0x00b45d95b5117447e2fafe7f34def913ff3ba220e4b8688acf37ae2328af7a3d.
//
// Solidity: event Redeposit(address token, bytes redepositIncAddress, uint256 amount, bytes32 itx)
func (_Vault *VaultFilterer) ParseRedeposit(log types.Log) (*VaultRedeposit, error) {
	event := new(VaultRedeposit)
	if err := _Vault.contract.UnpackLog(event, "Redeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultUpdateIncognitoProxyIterator is returned from FilterUpdateIncognitoProxy and is used to iterate over the raw logs and unpacked data for UpdateIncognitoProxy events raised by the Vault contract.
type VaultUpdateIncognitoProxyIterator struct {
	Event *VaultUpdateIncognitoProxy // Event containing the contract specifics and raw log

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
func (it *VaultUpdateIncognitoProxyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultUpdateIncognitoProxy)
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
		it.Event = new(VaultUpdateIncognitoProxy)
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
func (it *VaultUpdateIncognitoProxyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultUpdateIncognitoProxyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultUpdateIncognitoProxy represents a UpdateIncognitoProxy event raised by the Vault contract.
type VaultUpdateIncognitoProxy struct {
	NewIncognitoProxy common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUpdateIncognitoProxy is a free log retrieval operation binding the contract event 0x204252dfe190ad6ef63db40a490f048b39f661de74628408f13cd0bb2d4c3446.
//
// Solidity: event UpdateIncognitoProxy(address newIncognitoProxy)
func (_Vault *VaultFilterer) FilterUpdateIncognitoProxy(opts *bind.FilterOpts) (*VaultUpdateIncognitoProxyIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "UpdateIncognitoProxy")
	if err != nil {
		return nil, err
	}
	return &VaultUpdateIncognitoProxyIterator{contract: _Vault.contract, event: "UpdateIncognitoProxy", logs: logs, sub: sub}, nil
}

// WatchUpdateIncognitoProxy is a free log subscription operation binding the contract event 0x204252dfe190ad6ef63db40a490f048b39f661de74628408f13cd0bb2d4c3446.
//
// Solidity: event UpdateIncognitoProxy(address newIncognitoProxy)
func (_Vault *VaultFilterer) WatchUpdateIncognitoProxy(opts *bind.WatchOpts, sink chan<- *VaultUpdateIncognitoProxy) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "UpdateIncognitoProxy")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultUpdateIncognitoProxy)
				if err := _Vault.contract.UnpackLog(event, "UpdateIncognitoProxy", log); err != nil {
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

// ParseUpdateIncognitoProxy is a log parse operation binding the contract event 0x204252dfe190ad6ef63db40a490f048b39f661de74628408f13cd0bb2d4c3446.
//
// Solidity: event UpdateIncognitoProxy(address newIncognitoProxy)
func (_Vault *VaultFilterer) ParseUpdateIncognitoProxy(log types.Log) (*VaultUpdateIncognitoProxy, error) {
	event := new(VaultUpdateIncognitoProxy)
	if err := _Vault.contract.UnpackLog(event, "UpdateIncognitoProxy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultUpdateTokenTotalIterator is returned from FilterUpdateTokenTotal and is used to iterate over the raw logs and unpacked data for UpdateTokenTotal events raised by the Vault contract.
type VaultUpdateTokenTotalIterator struct {
	Event *VaultUpdateTokenTotal // Event containing the contract specifics and raw log

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
func (it *VaultUpdateTokenTotalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultUpdateTokenTotal)
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
		it.Event = new(VaultUpdateTokenTotal)
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
func (it *VaultUpdateTokenTotalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultUpdateTokenTotalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultUpdateTokenTotal represents a UpdateTokenTotal event raised by the Vault contract.
type VaultUpdateTokenTotal struct {
	Assets  []common.Address
	Amounts []*big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateTokenTotal is a free log retrieval operation binding the contract event 0x6a7fbbcddfd518bb8c56b28ac6c7acb0f7ca093ed232eb3306e53d14e469895f.
//
// Solidity: event UpdateTokenTotal(address[] assets, uint256[] amounts)
func (_Vault *VaultFilterer) FilterUpdateTokenTotal(opts *bind.FilterOpts) (*VaultUpdateTokenTotalIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "UpdateTokenTotal")
	if err != nil {
		return nil, err
	}
	return &VaultUpdateTokenTotalIterator{contract: _Vault.contract, event: "UpdateTokenTotal", logs: logs, sub: sub}, nil
}

// WatchUpdateTokenTotal is a free log subscription operation binding the contract event 0x6a7fbbcddfd518bb8c56b28ac6c7acb0f7ca093ed232eb3306e53d14e469895f.
//
// Solidity: event UpdateTokenTotal(address[] assets, uint256[] amounts)
func (_Vault *VaultFilterer) WatchUpdateTokenTotal(opts *bind.WatchOpts, sink chan<- *VaultUpdateTokenTotal) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "UpdateTokenTotal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultUpdateTokenTotal)
				if err := _Vault.contract.UnpackLog(event, "UpdateTokenTotal", log); err != nil {
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

// ParseUpdateTokenTotal is a log parse operation binding the contract event 0x6a7fbbcddfd518bb8c56b28ac6c7acb0f7ca093ed232eb3306e53d14e469895f.
//
// Solidity: event UpdateTokenTotal(address[] assets, uint256[] amounts)
func (_Vault *VaultFilterer) ParseUpdateTokenTotal(log types.Log) (*VaultUpdateTokenTotal, error) {
	event := new(VaultUpdateTokenTotal)
	if err := _Vault.contract.UnpackLog(event, "UpdateTokenTotal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Vault contract.
type VaultWithdrawIterator struct {
	Event *VaultWithdraw // Event containing the contract specifics and raw log

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
func (it *VaultWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultWithdraw)
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
		it.Event = new(VaultWithdraw)
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
func (it *VaultWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultWithdraw represents a Withdraw event raised by the Vault contract.
type VaultWithdraw struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address token, address to, uint256 amount)
func (_Vault *VaultFilterer) FilterWithdraw(opts *bind.FilterOpts) (*VaultWithdrawIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &VaultWithdrawIterator{contract: _Vault.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address token, address to, uint256 amount)
func (_Vault *VaultFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *VaultWithdraw) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultWithdraw)
				if err := _Vault.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address token, address to, uint256 amount)
func (_Vault *VaultFilterer) ParseWithdraw(log types.Log) (*VaultWithdraw, error) {
	event := new(VaultWithdraw)
	if err := _Vault.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
