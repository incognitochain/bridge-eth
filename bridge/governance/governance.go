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

// IGovernorCompatibilityBravoUpgradeableReceipt is an auto generated low-level Go binding around an user-defined struct.
type IGovernorCompatibilityBravoUpgradeableReceipt struct {
	HasVoted bool
	Support  uint8
	Votes    *big.Int
}

// GovernanceMetaData contains all meta data concerning the Governance contract.
var GovernanceMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"eta\",\"type\":\"uint256\"}],\"name\":\"ProposalQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldQuorumNumerator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newQuorumNumerator\",\"type\":\"uint256\"}],\"name\":\"QuorumNumeratorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"VoteCast\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"VoteCastWithParams\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BALLOT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COUNTING_MODE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXTENDED_BALLOT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROPOSAL_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"cancelBySig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"}],\"name\":\"castVote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"castVoteBySig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"castVoteWithReason\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"castVoteWithReasonAndParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"castVoteWithReasonAndParamsBySig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getActions\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_input\",\"type\":\"bytes32\"}],\"name\":\"getDataSign\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"getReceipt\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"hasVoted\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"uint96\",\"name\":\"votes\",\"type\":\"uint96\"}],\"internalType\":\"structIGovernorCompatibilityBravoUpgradeable.Receipt\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"getVotesWithParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"hashProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIVotesUpgradeable\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalDeadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposalEta\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"eta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"againstVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"abstainVotes\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"canceled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"propose\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"propose\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"proposeBySig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"queue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"queue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"quorum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorumDenominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"quorumNumerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorumNumerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorumVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"sigDataUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumIGovernorUpgradeable.ProposalState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timelock\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIVotesUpgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newQuorumNumerator\",\"type\":\"uint256\"}],\"name\":\"updateQuorumNumerator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b50614d91806100206000396000f3fe6080604052600436106102cd5760003560e01c80637d5e81e211610175578063c59057e4116100dc578063e23a9a5211610095578063f7d25dc01161006f578063f7d25dc014610a71578063f8ce560a14610a91578063fc0c546a14610ab1578063fe0d94c114610ad257600080fd5b8063e23a9a521461095a578063eb9019d414610a25578063f23a6e6114610a4557600080fd5b8063c59057e414610858578063d33219b414610878578063da95691a146108a0578063dd4e2ba5146108c0578063ddf0b00914610906578063deaaa7cc1461092657600080fd5b8063b58131b01161012e578063b58131b01461078e578063bc197c81146107a5578063bdc9b04c146107d1578063c01f9e3714610805578063c28bc2fa14610825578063c4d66de81461083857600080fd5b80637d5e81e2146106e45780638783ef441461070457806397c3d334146107245780639a802a6d14610738578063a7713a7014610758578063ab58fb8e1461076d57600080fd5b80632d63f6931161023457806340e58ee5116101ed57806356781388116101c757806356781388146106645780635f398a141461068457806360c4247f146106a45780637b3c71d3146106c457600080fd5b806340e58ee5146105cf57806343859632146105ef57806354fd4d501461063a57600080fd5b80632d63f693146104e95780632fe3e26114610509578063328dd9821461053d5780633932abb11461056d5780633bccf4fd146105825780633e4f49e6146105a257600080fd5b8063150b7a0211610286578063150b7a0214610407578063160cbed71461044b5780631ea1940e1461047057806320e297ac146104a157806324bc1a64146104c15780632656227d146104d657600080fd5b8063013cf08b146102db57806301ffc9a71461035657806302a251a31461038657806303420181146103a557806306f3f9e6146103c557806306fdde03146103e557600080fd5b366102d657005b005b600080fd5b3480156102e757600080fd5b506102fb6102f6366004613bf5565b610aea565b604080519a8b526001600160a01b0390991660208b0152978901969096526060880194909452608087019290925260a086015260c085015260e084015215156101008301521515610120820152610140015b60405180910390f35b34801561036257600080fd5b50610376610371366004613c0e565b610b89565b604051901515815260200161034d565b34801561039257600080fd5b5061b3cb5b60405190815260200161034d565b3480156103b157600080fd5b506103976103c0366004613d44565b610b9a565b3480156103d157600080fd5b506102d46103e0366004613bf5565b610c92565b3480156103f157600080fd5b506103fa610d15565b60405161034d9190613e3a565b34801561041357600080fd5b50610432610422366004613e62565b630a85bd0160e11b949350505050565b6040516001600160e01b0319909116815260200161034d565b34801561045757600080fd5b5061039761046636600461403e565b6000949350505050565b34801561047c57600080fd5b5061037661048b366004613bf5565b6101f96020526000908152604090205460ff1681565b3480156104ad57600080fd5b506102d46104bc3660046140cd565b610da8565b3480156104cd57600080fd5b50610397610e0a565b6103976104e436600461403e565b610e1a565b3480156104f557600080fd5b50610397610504366004613bf5565b610f47565b34801561051557600080fd5b506103977fb3b3f3b703cd84ce352197dcff232b1b5d3cfb2025ce47cf04742d0651f1af8881565b34801561054957600080fd5b5061055d610558366004613bf5565b610f7f565b60405161034d94939291906141d1565b34801561057957600080fd5b506119af610397565b34801561058e57600080fd5b5061039761059d36600461421e565b611211565b3480156105ae57600080fd5b506105c26105bd366004613bf5565b611287565b60405161034d9190614282565b3480156105db57600080fd5b506102d46105ea366004613bf5565b611292565b3480156105fb57600080fd5b5061037661060a3660046142aa565b6000828152610161602090815260408083206001600160a01b038516845260080190915290205460ff1692915050565b34801561064657600080fd5b506040805180820190915260018152603160f81b60208201526103fa565b34801561067057600080fd5b5061039761067f3660046142da565b6115af565b34801561069057600080fd5b5061039761069f366004614306565b6115df565b3480156106b057600080fd5b506103976106bf366004613bf5565b611630565b3480156106d057600080fd5b506103976106df366004614389565b6116cb565b3480156106f057600080fd5b506103976106ff3660046143e2565b611724565b34801561071057600080fd5b5061039761071f366004614482565b61173b565b34801561073057600080fd5b506064610397565b34801561074457600080fd5b50610397610753366004614552565b61184e565b34801561076457600080fd5b50610397611865565b34801561077957600080fd5b50610397610788366004613bf5565b50600090565b34801561079a57600080fd5b506305f5e100610397565b3480156107b157600080fd5b506104326107c03660046145aa565b63bc197c8160e01b95945050505050565b3480156107dd57600080fd5b506103977fb6d2dc83590271a7c0a5ab5fbf6a2dad418bbfd533c253e3d69a6772712809c781565b34801561081157600080fd5b50610397610820366004613bf5565b611892565b6102d461083336600461463d565b6118c2565b34801561084457600080fd5b506102d4610853366004614680565b6119c5565b34801561086457600080fd5b5061039761087336600461403e565b611b10565b34801561088457600080fd5b5060005b6040516001600160a01b03909116815260200161034d565b3480156108ac57600080fd5b506103976108bb36600461471c565b611b4a565b3480156108cc57600080fd5b5060408051808201909152601a81527f737570706f72743d627261766f2671756f72756d3d627261766f00000000000060208201526103fa565b34801561091257600080fd5b506102d4610921366004613bf5565b611b76565b34801561093257600080fd5b506103977f150214d74d59b7d1e90c73fc22ef3d991dd0a76b046543d4d80ab92d2a50328f81565b34801561096657600080fd5b506109f56109753660046142aa565b6040805160608101825260008082526020820181905291810191909152506000918252610161602090815260408084206001600160a01b0393909316845260089092018152918190208151606081018352905460ff8082161515835261010082041693820193909352620100009092046001600160601b03169082015290565b6040805182511515815260208084015160ff1690820152918101516001600160601b03169082015260600161034d565b348015610a3157600080fd5b50610397610a403660046147bd565b611de4565b348015610a5157600080fd5b50610432610a603660046147e9565b63f23a6e6160e01b95945050505050565b348015610a7d57600080fd5b50610397610a8c366004613bf5565b611e05565b348015610a9d57600080fd5b50610397610aac366004613bf5565b611e10565b348015610abd57600080fd5b5061019354610888906001600160a01b031681565b6102d4610ae0366004613bf5565b611e1b565b905090565b8060008080808080808080610afe8a610f47565b9650610b098b611892565b60008c81526101616020526040812080546005820154600683015460078401546001600160a01b039093169e50949a509850929650919450610b4a8d611287565b90506002816007811115610b6057610b6061426c565b1493506007816007811115610b7757610b7761426c565b14925050509193959799509193959799565b6000610b948261208a565b92915050565b600080610c3e610c367fb3b3f3b703cd84ce352197dcff232b1b5d3cfb2025ce47cf04742d0651f1af888c8c8c8c604051610bd6929190614851565b60405180910390208b80519060200120604051602001610c1b959493929190948552602085019390935260ff9190911660408401526060830152608082015260a00190565b604051602081830303815290604052805190602001206120f5565b868686612143565b9050610c848a828b8b8b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508d9250612161915050565b9a9950505050505050505050565b30610c9b6122c6565b6001600160a01b031614610cf15760405162461bcd60e51b8152602060048201526018602482015277476f7665726e6f723a206f6e6c79476f7665726e616e636560401b60448201526064015b60405180910390fd5b610d09565b80610d026101316122f0565b03610cf657505b610d128161236f565b50565b606061012f8054610d2590614861565b80601f0160208091040260200160405190810160405280929190818152602001828054610d5190614861565b8015610d9e5780601f10610d7357610100808354040283529160200191610d9e565b820191906000526020600020905b815481529060010190602001808311610d8157829003601f168201915b5050505050905090565b610dc9610dc185604051602001610c1b91815260200190565b848484612143565b6101f880546001600160a01b0319166001600160a01b0392909216919091179055610df384611292565b50506101f880546001600160a01b03191690555050565b6000610ae5610aac6001436148b1565b600080610e2986868686611b10565b90506000610e3682611287565b90506004816007811115610e4c57610e4c61426c565b1480610e6957506005816007811115610e6757610e6761426c565b145b610ebf5760405162461bcd60e51b815260206004820152602160248201527f476f7665726e6f723a2070726f706f73616c206e6f74207375636365737366756044820152601b60fa1b6064820152608401610ce8565b6000828152610130602052604090819020600201805460ff19166001179055517f712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f90610f0e9084815260200190565b60405180910390a1610f2382888888886124b7565b610f308288888888612545565b610f3d82888888886124b7565b5095945050505050565b6000818152610130602090815260408083208151928301909152546001600160401b0316908190525b6001600160401b031692915050565b606080606080600061016160008781526020019081526020016000209050806001018160020182600301836004018380548060200260200160405190810160405280929190818152602001828054801561100257602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610fe4575b505050505093508280548060200260200160405190810160405280929190818152602001828054801561105457602002820191906000526020600020905b815481526020019060010190808311611040575b5050505050925081805480602002602001604051908101604052809291908181526020016000905b8282101561112857838290600052602060002001805461109b90614861565b80601f01602080910402602001604051908101604052809291908181526020018280546110c790614861565b80156111145780601f106110e957610100808354040283529160200191611114565b820191906000526020600020905b8154815290600101906020018083116110f757829003601f168201915b50505050508152602001906001019061107c565b50505050915080805480602002602001604051908101604052809291908181526020016000905b828210156111fb57838290600052602060002001805461116e90614861565b80601f016020809104026020016040519081016040528092919081815260200182805461119a90614861565b80156111e75780601f106111bc576101008083540402835291602001916111e7565b820191906000526020600020905b8154815290600101906020018083116111ca57829003601f168201915b50505050508152602001906001019061114f565b5050505090509450945094509450509193509193565b604080517f150214d74d59b7d1e90c73fc22ef3d991dd0a76b046543d4d80ab92d2a50328f602082015290810186905260ff85166060820152600090819061125f90610c3690608001610c1b565b905061127c87828860405180602001604052806000815250612552565b979650505050505050565b6000610b9482612575565b60008181526101616020526040902080546001600160a01b03166112b46122c6565b6001600160a01b031614806112e557506305f5e10081546112e3906001600160a01b0316610a406001436148b1565b105b6113415760405162461bcd60e51b815260206004820152602760248201527f476f7665726e6f72427261766f3a2070726f706f7365722061626f76652074686044820152661c995cda1bdb1960ca1b6064820152608401610ce8565b6115aa8160010180548060200260200160405190810160405280929190818152602001828054801561139c57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161137e575b5050505050826002018054806020026020016040519081016040528092919081815260200182805480156113ef57602002820191906000526020600020905b8154815260200190600101908083116113db575b50505050506115a084600301805480602002602001604051908101604052809291908181526020016000905b828210156114c757838290600052602060002001805461143a90614861565b80601f016020809104026020016040519081016040528092919081815260200182805461146690614861565b80156114b35780601f10611488576101008083540402835291602001916114b3565b820191906000526020600020905b81548152906001019060200180831161149657829003601f168201915b50505050508152602001906001019061141b565b50505060048701805460408051602080840282018101909252828152935060009084015b8282101561159757838290600052602060002001805461150a90614861565b80601f016020809104026020016040519081016040528092919081815260200182805461153690614861565b80156115835780601f1061155857610100808354040283529160200191611583565b820191906000526020600020905b81548152906001019060200180831161156657829003601f168201915b5050505050815260200190600101906114eb565b50505050612689565b84600901546127bb565b505050565b6000806115ba6122c6565b90506115d784828560405180602001604052806000815250612552565b949350505050565b6000806115ea6122c6565b905061127c87828888888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508a9250612161915050565b6101c7546000908082036116495750506101c654919050565b60006101c76116596001846148b1565b81548110611669576116696148c4565b60009182526020918290206040805180820190915291015463ffffffff8116808352600160201b9091046001600160e01b031692820192909252915084106116bf57602001516001600160e01b03169392505050565b6115d76101c7856127c9565b6000806116d66122c6565b905061171a86828787878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061255292505050565b9695505050505050565b600061173285858585612889565b95945050505050565b60008061177a7fb6d2dc83590271a7c0a5ab5fbf6a2dad418bbfd533c253e3d69a6772712809c78a8a8a8a604051602001610c1b9594939291906148da565b60008181526101f9602052604090205490915060ff16156117dd5760405162461bcd60e51b815260206004820152601a60248201527f476f7665726e616e63653a207369676e617475726520757365640000000000006044820152606401610ce8565b60008181526101f960205260409020805460ff1916600117905561180381868686612143565b6101f880546001600160a01b0319166001600160a01b039290921691909117905561183089898989612889565b6101f880546001600160a01b03191690559998505050505050505050565b600061185b8484846128f8565b90505b9392505050565b6101c7546000901561188a5761187c6101c761296f565b6001600160e01b0316905090565b506101c65490565b6000818152610130602090815260408083208151928301909152600101546001600160401b031690819052610f70565b306118cb6122c6565b6001600160a01b03161461191c5760405162461bcd60e51b8152602060048201526018602482015277476f7665726e6f723a206f6e6c79476f7665726e616e636560401b6044820152606401610ce8565b611934565b8061192d6101316122f0565b0361192157505b600080856001600160a01b0316858585604051611952929190614851565b60006040518083038185875af1925050503d806000811461198f576040519150601f19603f3d011682016040523d82523d6000602084013e611994565b606091505b50915091506119bc8282604051806060016040528060288152602001614d0d602891396129a8565b50505050505050565b600054610100900460ff16158080156119e55750600054600160ff909116105b806119ff5750303b1580156119ff575060005460ff166001145b611a625760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610ce8565b6000805460ff191660011790558015611a85576000805461ff0019166101001790555b611ab26040518060400160405280600c81526020016b496e636f676e69746f44414f60a01b8152506129c1565b611abb82612a18565b611ac56032612a48565b8015611b0c576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020015b60405180910390a15b5050565b600084848484604051602001611b299493929190614939565b60408051601f19818403018152919052805160209091012095945050505050565b6000611b61611b576122c6565b8787878787612a78565b61171a8686611b708787612689565b85611724565b6000818152610161602090815260409182902060018101805484518185028101850190955280855291936115aa93909290830182828015611be057602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611bc2575b505050505082600201805480602002602001604051908101604052809291908181526020018280548015611c3357602002820191906000526020600020905b815481526020019060010190808311611c1f575b5050505050611ddb84600301805480602002602001604051908101604052809291908181526020016000905b82821015611d0b578382906000526020600020018054611c7e90614861565b80601f0160208091040260200160405190810160405280929190818152602001828054611caa90614861565b8015611cf75780601f10611ccc57610100808354040283529160200191611cf7565b820191906000526020600020905b815481529060010190602001808311611cda57829003601f168201915b505050505081526020019060010190611c5f565b50505060048701805460408051602080840282018101909252828152935060009084015b82821015611597578382906000526020600020018054611d4e90614861565b80601f0160208091040260200160405190810160405280929190818152602001828054611d7a90614861565b8015611dc75780601f10611d9c57610100808354040283529160200191611dc7565b820191906000526020600020905b815481529060010190602001808311611daa57829003601f168201915b505050505081526020019060010190611d2f565b50600092915050565b600061185e8383611e0060408051602081019091526000815290565b6128f8565b6000610b94826120f5565b6000610b9482612b36565b6000818152610161602090815260409182902060018101805484518185028101850190955280855291936115aa93909290830182828015611e8557602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611e67575b505050505082600201805480602002602001604051908101604052809291908181526020018280548015611ed857602002820191906000526020600020905b815481526020019060010190808311611ec4575b505050505061208084600301805480602002602001604051908101604052809291908181526020016000905b82821015611fb0578382906000526020600020018054611f2390614861565b80601f0160208091040260200160405190810160405280929190818152602001828054611f4f90614861565b8015611f9c5780601f10611f7157610100808354040283529160200191611f9c565b820191906000526020600020905b815481529060010190602001808311611f7f57829003601f168201915b505050505081526020019060010190611f04565b50505060048701805460408051602080840282018101909252828152935060009084015b82821015611597578382906000526020600020018054611ff390614861565b80601f016020809104026020016040519081016040528092919081815260200182805461201f90614861565b801561206c5780601f106120415761010080835404028352916020019161206c565b820191906000526020600020905b81548152906001019060200180831161204f57829003601f168201915b505050505081526020019060010190611fd4565b8460090154610e1a565b60006001600160e01b0319821663bf26d89760e01b14806120bb57506001600160e01b031982166379dd796f60e01b145b806120d657506001600160e01b03198216630271189760e51b145b80610b9457506301ffc9a760e01b6001600160e01b0319831614610b94565b6000610b94612102612bc5565b8360405161190160f01b6020820152602281018390526042810182905260009060620160405160208183030381529060405280519060200120905092915050565b600080600061215487878787612c40565b91509150610f3d81612d04565b600085815261013060205260408120600161217b88611287565b600781111561218c5761218c61426c565b146121e55760405162461bcd60e51b815260206004820152602360248201527f476f7665726e6f723a20766f7465206e6f742063757272656e746c792061637460448201526269766560e81b6064820152608401610ce8565b604080516020810190915281546001600160401b03169081905260009061220e908890866128f8565b905061221d8888888488612e4e565b835160000361227257866001600160a01b03167fb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4898884896040516122659493929190614984565b60405180910390a261127c565b866001600160a01b03167fe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb871289888489896040516122b39594939291906149ac565b60405180910390a2979650505050505050565b6101f8546000906001600160a01b03166122df57503390565b506101f8546001600160a01b031690565b600061230b8254600f81810b600160801b909204900b131590565b1561232957604051631ed9509560e11b815260040160405180910390fd5b508054600f0b6000818152600180840160205260408220805492905583546fffffffffffffffffffffffffffffffff191692016001600160801b03169190911790915590565b60648111156123f25760405162461bcd60e51b815260206004820152604360248201527f476f7665726e6f72566f74657351756f72756d4672616374696f6e3a2071756f60448201527f72756d4e756d657261746f72206f7665722071756f72756d44656e6f6d696e616064820152623a37b960e91b608482015260a401610ce8565b60006123fc611865565b9050801580159061240e57506101c754155b156124735760408051808201909152600081526101c7906020810161243284612fec565b6001600160e01b039081169091528254600181018455600093845260209384902083519490930151909116600160201b0263ffffffff909316929092179101555b61247f6101c783613059565b505060408051828152602081018490527f0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b46339979101611b03565b61253e565b845181101561253c57306001600160a01b03168582815181106124e1576124e16148c4565b60200260200101516001600160a01b03160361252c5761252c83828151811061250c5761250c6148c4565b60200260200101518051906020012061013161308f90919063ffffffff16565b612535816149e6565b90506124bc565b505b5050505050565b61253e85858585856130cb565b60006117328585858561257060408051602081019091526000815290565b612161565b600081815261013060205260408120600281015460ff161561259a5750600792915050565b6002810154610100900460ff16156125b55750600292915050565b60006125c084610f47565b9050806000036126125760405162461bcd60e51b815260206004820152601d60248201527f476f7665726e6f723a20756e6b6e6f776e2070726f706f73616c2069640000006044820152606401610ce8565b438110612623575060009392505050565b600061262e85611892565b905043811061264257506001949350505050565b61264b856131c1565b801561266e57506000858152610161602052604090206006810154600590910154115b1561267e57506004949350505050565b506003949350505050565b6060600082516001600160401b038111156126a6576126a6613c8f565b6040519080825280602002602001820160405280156126d957816020015b60608152602001906001900390816126c45790505b50905060005b84518110156127b3578481815181106126fa576126fa6148c4565b60200260200101515160001461276a5784818151811061271c5761271c6148c4565b60200260200101518051906020012084828151811061273d5761273d6148c4565b60200260200101516040516020016127569291906149ff565b604051602081830303815290604052612785565b83818151811061277c5761277c6148c4565b60200260200101515b828281518110612797576127976148c4565b6020026020010181905250806127ac906149e6565b90506126df565b509392505050565b6000611732858585856131ea565b600043821061281a5760405162461bcd60e51b815260206004820181905260248201527f436865636b706f696e74733a20626c6f636b206e6f7420796574206d696e65646044820152606401610ce8565b600061282583613309565b845490915060006128388684838561336e565b905080156128735761285d8661284f6001846148b1565b600091825260209091200190565b54600160201b90046001600160e01b0316612876565b60005b6001600160e01b03169695505050505050565b60006128ec6128966122c6565b868686516001600160401b038111156128b1576128b1613c8f565b6040519080825280602002602001820160405280156128e457816020015b60608152602001906001900390816128cf5790505b508787612a78565b611732858585856133c4565b61019354604051630748d63560e31b81526001600160a01b038581166004830152602482018590526000921690633a46b1a890604401602060405180830381865afa15801561294b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061185b9190614a30565b8054600090801561299f576129898361284f6001846148b1565b54600160201b90046001600160e01b031661185e565b60009392505050565b606083156129b757508161185e565b61185e838361368e565b600054610100900460ff166129e85760405162461bcd60e51b8152600401610ce890614a49565b612a0f81612a0a6040805180820190915260018152603160f81b602082015290565b6136b8565b610d12816136f9565b600054610100900460ff16612a3f5760405162461bcd60e51b8152600401610ce890614a49565b610d128161372d565b600054610100900460ff16612a6f5760405162461bcd60e51b8152600401610ce890614a49565b610d1281613777565b805160208201206000612a968787612a908888612689565b85611b10565b600081815261016160205260409020600981015491925090612b2b5780546001600160a01b0319166001600160a01b038a161781558751612ae090600183019060208b0190613a2c565b508651612af690600283019060208a0190613a8d565b508551612b0c9060038301906020890190613ac8565b508451612b229060048301906020880190613b1a565b50600981018390555b505050505050505050565b60006064612b4383611630565b61019354604051632394e7a360e21b8152600481018690526001600160a01b0390911690638e539e8c90602401602060405180830381865afa158015612b8d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612bb19190614a30565b612bbb9190614a94565b610b949190614aab565b6000610ae57f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f612bf460655490565b6066546040805160208101859052908101839052606081018290524660808201523060a082015260009060c0016040516020818303038152906040528051906020012090509392505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115612c775750600090506003612cfb565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015612ccb573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116612cf457600060019250925050612cfb565b9150600090505b94509492505050565b6000816004811115612d1857612d1861426c565b03612d205750565b6001816004811115612d3457612d3461426c565b03612d815760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610ce8565b6002816004811115612d9557612d9561426c565b03612de25760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610ce8565b6003816004811115612df657612df661426c565b03610d125760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610ce8565b6000858152610161602090815260408083206001600160a01b038816845260088101909252909120805460ff1615612ede5760405162461bcd60e51b815260206004820152602d60248201527f476f7665726e6f72436f6d7061746962696c697479427261766f3a20766f746560448201526c08185b1c9958591e4818d85cdd609a1b6064820152608401610ce8565b805460ff86166101000261ffff19909116176001178155612efe8461379e565b81546001600160601b039190911662010000026dffffffffffffffffffffffff00001990911617815560ff8516612f4e5783826006016000828254612f439190614acd565b909155506119bc9050565b60001960ff861601612f6e5783826005016000828254612f439190614acd565b60011960ff861601612f8e5783826007016000828254612f439190614acd565b60405162461bcd60e51b815260206004820152602d60248201527f476f7665726e6f72436f6d7061746962696c697479427261766f3a20696e766160448201526c6c696420766f7465207479706560981b6064820152608401610ce8565b60006001600160e01b038211156130555760405162461bcd60e51b815260206004820152602760248201527f53616665436173743a2076616c756520646f65736e27742066697420696e20326044820152663234206269747360c81b6064820152608401610ce8565b5090565b6000806130778461306943613309565b61307286612fec565b613806565b6001600160e01b0391821693501690505b9250929050565b8154600160801b90819004600f0b6000818152600180860160205260409091209390935583546001600160801b03908116939091011602179055565b6000604051806060016040528060278152602001614d3560279139905060005b85518110156119bc57600080878381518110613109576131096148c4565b60200260200101516001600160a01b031687848151811061312c5761312c6148c4565b6020026020010151878581518110613146576131466148c4565b602002602001015160405161315b9190614ae0565b60006040518083038185875af1925050503d8060008114613198576040519150601f19603f3d011682016040523d82523d6000602084013e61319d565b606091505b50915091506131ad8282866129a8565b505050806131ba906149e6565b90506130eb565b60008181526101616020526040812060058101546131e1610aac85610f47565b11159392505050565b6000806131f986868686611b10565b9050600061320682611287565b9050600281600781111561321c5761321c61426c565b1415801561323c575060068160078111156132395761323961426c565b14155b801561325a575060078160078111156132575761325761426c565b14155b6132a65760405162461bcd60e51b815260206004820152601d60248201527f476f7665726e6f723a2070726f706f73616c206e6f74206163746976650000006044820152606401610ce8565b6000828152610130602052604090819020600201805461ff001916610100179055517f789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c906132f79084815260200190565b60405180910390a15095945050505050565b600063ffffffff8211156130555760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203360448201526532206269747360d01b6064820152608401610ce8565b60005b818310156127b357600061338584846139a9565b60008781526020902090915063ffffffff86169082015463ffffffff1611156133b0578092506133be565b6133bb816001614acd565b93505b50613371565b60006305f5e1006133e16133d66122c6565b610a406001436148b1565b10156134495760405162461bcd60e51b815260206004820152603160248201527f476f7665726e6f723a2070726f706f73657220766f7465732062656c6f7720706044820152701c9bdc1bdcd85b081d1a1c995cda1bdb19607a1b6064820152608401610ce8565b600061345e8686868680519060200120611b10565b905084518651146134815760405162461bcd60e51b8152600401610ce890614afc565b83518651146134a25760405162461bcd60e51b8152600401610ce890614afc565b60008651116134f35760405162461bcd60e51b815260206004820152601860248201527f476f7665726e6f723a20656d7074792070726f706f73616c00000000000000006044820152606401610ce8565b60008181526101306020908152604091829020825191820190925281546001600160401b031690819052156135745760405162461bcd60e51b815260206004820152602160248201527f476f7665726e6f723a2070726f706f73616c20616c72656164792065786973746044820152607360f81b6064820152608401610ce8565b60006135816119af6139c4565b61358a436139c4565b6135949190614b3d565b905060006135a361b3cb6139c4565b6135ad9083614b3d565b835467ffffffffffffffff19166001600160401b038416178455905060018301805467ffffffffffffffff19166001600160401b0383161790557f7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e0846136116122c6565b8b8b8d516001600160401b0381111561362c5761362c613c8f565b60405190808252806020026020018201604052801561365f57816020015b606081526020019060019003908161364a5790505b508c88888e60405161367999989796959493929190614b64565b60405180910390a15091979650505050505050565b81511561369e5781518083602001fd5b8060405162461bcd60e51b8152600401610ce89190613e3a565b600054610100900460ff166136df5760405162461bcd60e51b8152600401610ce890614a49565b815160209283012081519190920120606591909155606655565b600054610100900460ff166137205760405162461bcd60e51b8152600401610ce890614a49565b61012f611b0c8282614c4d565b600054610100900460ff166137545760405162461bcd60e51b8152600401610ce890614a49565b61019380546001600160a01b0319166001600160a01b0392909216919091179055565b600054610100900460ff16610d095760405162461bcd60e51b8152600401610ce890614a49565b60006001600160601b038211156130555760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203960448201526536206269747360d01b6064820152608401610ce8565b82546000908190801561394c5760006138248761284f6001856148b1565b60408051808201909152905463ffffffff808216808452600160201b9092046001600160e01b0316602084015291925090871610156138a55760405162461bcd60e51b815260206004820152601760248201527f436865636b706f696e743a20696e76616c6964206b65790000000000000000006044820152606401610ce8565b805163ffffffff8088169116036138ed57846138c68861284f6001866148b1565b80546001600160e01b0392909216600160201b0263ffffffff90921691909117905561393c565b6040805180820190915263ffffffff80881682526001600160e01b0380881660208085019182528b54600181018d5560008d81529190912094519151909216600160201b029216919091179101555b6020015192508391506139a19050565b50506040805180820190915263ffffffff80851682526001600160e01b0380851660208085019182528854600181018a5560008a815291822095519251909316600160201b0291909316179201919091559050815b935093915050565b60006139b86002848418614aab565b61185e90848416614acd565b60006001600160401b038211156130555760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203660448201526534206269747360d01b6064820152608401610ce8565b828054828255906000526020600020908101928215613a81579160200282015b82811115613a8157825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190613a4c565b50613055929150613b6c565b828054828255906000526020600020908101928215613a81579160200282015b82811115613a81578251825591602001919060010190613aad565b828054828255906000526020600020908101928215613b0e579160200282015b82811115613b0e5782518290613afe9082614c4d565b5091602001919060010190613ae8565b50613055929150613b81565b828054828255906000526020600020908101928215613b60579160200282015b82811115613b605782518290613b509082614c4d565b5091602001919060010190613b3a565b50613055929150613b9e565b5b808211156130555760008155600101613b6d565b80821115613055576000613b958282613bbb565b50600101613b81565b80821115613055576000613bb28282613bbb565b50600101613b9e565b508054613bc790614861565b6000825580601f10613bd7575050565b601f016020900490600052602060002090810190610d129190613b6c565b600060208284031215613c0757600080fd5b5035919050565b600060208284031215613c2057600080fd5b81356001600160e01b03198116811461185e57600080fd5b803560ff81168114613c4957600080fd5b919050565b60008083601f840112613c6057600080fd5b5081356001600160401b03811115613c7757600080fd5b60208301915083602082850101111561308857600080fd5b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715613ccd57613ccd613c8f565b604052919050565b600082601f830112613ce657600080fd5b81356001600160401b03811115613cff57613cff613c8f565b613d12601f8201601f1916602001613ca5565b818152846020838601011115613d2757600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060008060008060e0898b031215613d6057600080fd5b88359750613d7060208a01613c38565b965060408901356001600160401b0380821115613d8c57600080fd5b613d988c838d01613c4e565b909850965060608b0135915080821115613db157600080fd5b50613dbe8b828c01613cd5565b945050613dcd60808a01613c38565b925060a0890135915060c089013590509295985092959890939650565b60005b83811015613e05578181015183820152602001613ded565b50506000910152565b60008151808452613e26816020860160208601613dea565b601f01601f19169290920160200192915050565b60208152600061185e6020830184613e0e565b6001600160a01b0381168114610d1257600080fd5b60008060008060808587031215613e7857600080fd5b8435613e8381613e4d565b93506020850135613e9381613e4d565b92506040850135915060608501356001600160401b03811115613eb557600080fd5b613ec187828801613cd5565b91505092959194509250565b60006001600160401b03821115613ee657613ee6613c8f565b5060051b60200190565b600082601f830112613f0157600080fd5b81356020613f16613f1183613ecd565b613ca5565b82815260059290921b84018101918181019086841115613f3557600080fd5b8286015b84811015613f59578035613f4c81613e4d565b8352918301918301613f39565b509695505050505050565b600082601f830112613f7557600080fd5b81356020613f85613f1183613ecd565b82815260059290921b84018101918181019086841115613fa457600080fd5b8286015b84811015613f595780358352918301918301613fa8565b600082601f830112613fd057600080fd5b81356020613fe0613f1183613ecd565b82815260059290921b84018101918181019086841115613fff57600080fd5b8286015b84811015613f595780356001600160401b038111156140225760008081fd5b6140308986838b0101613cd5565b845250918301918301614003565b6000806000806080858703121561405457600080fd5b84356001600160401b038082111561406b57600080fd5b61407788838901613ef0565b9550602087013591508082111561408d57600080fd5b61409988838901613f64565b945060408701359150808211156140af57600080fd5b506140bc87828801613fbf565b949793965093946060013593505050565b600080600080608085870312156140e357600080fd5b843593506140f360208601613c38565b93969395505050506040820135916060013590565b600081518084526020808501945080840160005b838110156141415781516001600160a01b03168752958201959082019060010161411c565b509495945050505050565b600081518084526020808501945080840160005b8381101561414157815187529582019590820190600101614160565b600081518084526020808501808196508360051b8101915082860160005b858110156141c45782840389526141b2848351613e0e565b9885019893509084019060010161419a565b5091979650505050505050565b6080815260006141e46080830187614108565b82810360208401526141f6818761414c565b9050828103604084015261420a818661417c565b9050828103606084015261127c818561417c565b600080600080600060a0868803121561423657600080fd5b8535945061424660208701613c38565b935061425460408701613c38565b94979396509394606081013594506080013592915050565b634e487b7160e01b600052602160045260246000fd5b60208101600883106142a457634e487b7160e01b600052602160045260246000fd5b91905290565b600080604083850312156142bd57600080fd5b8235915060208301356142cf81613e4d565b809150509250929050565b600080604083850312156142ed57600080fd5b823591506142fd60208401613c38565b90509250929050565b60008060008060006080868803121561431e57600080fd5b8535945061432e60208701613c38565b935060408601356001600160401b038082111561434a57600080fd5b61435689838a01613c4e565b9095509350606088013591508082111561436f57600080fd5b5061437c88828901613cd5565b9150509295509295909350565b6000806000806060858703121561439f57600080fd5b843593506143af60208601613c38565b925060408501356001600160401b038111156143ca57600080fd5b6143d687828801613c4e565b95989497509550505050565b600080600080608085870312156143f857600080fd5b84356001600160401b038082111561440f57600080fd5b61441b88838901613ef0565b9550602087013591508082111561443157600080fd5b61443d88838901613f64565b9450604087013591508082111561445357600080fd5b61445f88838901613fbf565b9350606087013591508082111561447557600080fd5b50613ec187828801613cd5565b600080600080600080600060e0888a03121561449d57600080fd5b87356001600160401b03808211156144b457600080fd5b6144c08b838c01613ef0565b985060208a01359150808211156144d657600080fd5b6144e28b838c01613f64565b975060408a01359150808211156144f857600080fd5b6145048b838c01613fbf565b965060608a013591508082111561451a57600080fd5b506145278a828b01613cd5565b94505061453660808901613c38565b925060a0880135915060c0880135905092959891949750929550565b60008060006060848603121561456757600080fd5b833561457281613e4d565b92506020840135915060408401356001600160401b0381111561459457600080fd5b6145a086828701613cd5565b9150509250925092565b600080600080600060a086880312156145c257600080fd5b85356145cd81613e4d565b945060208601356145dd81613e4d565b935060408601356001600160401b03808211156145f957600080fd5b61460589838a01613f64565b9450606088013591508082111561461b57600080fd5b61462789838a01613f64565b9350608088013591508082111561436f57600080fd5b6000806000806060858703121561465357600080fd5b843561465e81613e4d565b93506020850135925060408501356001600160401b038111156143ca57600080fd5b60006020828403121561469257600080fd5b813561185e81613e4d565b600082601f8301126146ae57600080fd5b813560206146be613f1183613ecd565b82815260059290921b840181019181810190868411156146dd57600080fd5b8286015b84811015613f595780356001600160401b038111156147005760008081fd5b61470e8986838b0101613cd5565b8452509183019183016146e1565b600080600080600060a0868803121561473457600080fd5b85356001600160401b038082111561474b57600080fd5b61475789838a01613ef0565b9650602088013591508082111561476d57600080fd5b61477989838a01613f64565b9550604088013591508082111561478f57600080fd5b61479b89838a0161469d565b945060608801359150808211156147b157600080fd5b61462789838a01613fbf565b600080604083850312156147d057600080fd5b82356147db81613e4d565b946020939093013593505050565b600080600080600060a0868803121561480157600080fd5b853561480c81613e4d565b9450602086013561481c81613e4d565b9350604086013592506060860135915060808601356001600160401b0381111561484557600080fd5b61437c88828901613cd5565b8183823760009101908152919050565b600181811c9082168061487557607f821691505b60208210810361489557634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b81810381811115610b9457610b9461489b565b634e487b7160e01b600052603260045260246000fd5b85815260a0602082015260006148f360a0830187614108565b8281036040840152614905818761414c565b90508281036060840152614919818661417c565b9050828103608084015261492d8185613e0e565b98975050505050505050565b60808152600061494c6080830187614108565b828103602084015261495e818761414c565b90508281036040840152614972818661417c565b91505082606083015295945050505050565b84815260ff8416602082015282604082015260806060820152600061171a6080830184613e0e565b85815260ff8516602082015283604082015260a0606082015260006149d460a0830185613e0e565b828103608084015261492d8185613e0e565b6000600182016149f8576149f861489b565b5060010190565b6001600160e01b0319831681528151600090614a22816004850160208701613dea565b919091016004019392505050565b600060208284031215614a4257600080fd5b5051919050565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b8082028115828204841417610b9457610b9461489b565b600082614ac857634e487b7160e01b600052601260045260246000fd5b500490565b80820180821115610b9457610b9461489b565b60008251614af2818460208701613dea565b9190910192915050565b60208082526021908201527f476f7665726e6f723a20696e76616c69642070726f706f73616c206c656e67746040820152600d60fb1b606082015260800190565b6001600160401b03818116838216019080821115614b5d57614b5d61489b565b5092915050565b8981526001600160a01b038916602082015261012060408201819052600090614b8f8382018b614108565b90508281036060840152614ba3818a61414c565b90508281036080840152614bb7818961417c565b905082810360a0840152614bcb818861417c565b6001600160401b0387811660c0860152861660e08501528381036101008501529050614bf78185613e0e565b9c9b505050505050505050505050565b601f8211156115aa57600081815260208120601f850160051c81016020861015614c2e5750805b601f850160051c820191505b8181101561253c57828155600101614c3a565b81516001600160401b03811115614c6657614c66613c8f565b614c7a81614c748454614861565b84614c07565b602080601f831160018114614caf5760008415614c975750858301515b600019600386901b1c1916600185901b17855561253c565b600085815260208120601f198616915b82811015614cde57888601518255948401946001909101908401614cbf565b5085821015614cfc5787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fe476f7665726e6f723a2072656c617920726576657274656420776974686f7574206d657373616765476f7665726e6f723a2063616c6c20726576657274656420776974686f7574206d657373616765a2646970667358221220e75bdb1dbc24b9e4dc8e9622434f42524fdea42dba1b84ef1f72f077c2e194fb64736f6c63430008110033",
}

// GovernanceABI is the input ABI used to generate the binding from.
// Deprecated: Use GovernanceMetaData.ABI instead.
var GovernanceABI = GovernanceMetaData.ABI

// GovernanceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GovernanceMetaData.Bin instead.
var GovernanceBin = GovernanceMetaData.Bin

// DeployGovernance deploys a new Ethereum contract, binding an instance of Governance to it.
func DeployGovernance(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Governance, error) {
	parsed, err := GovernanceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GovernanceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Governance{GovernanceCaller: GovernanceCaller{contract: contract}, GovernanceTransactor: GovernanceTransactor{contract: contract}, GovernanceFilterer: GovernanceFilterer{contract: contract}}, nil
}

// Governance is an auto generated Go binding around an Ethereum contract.
type Governance struct {
	GovernanceCaller     // Read-only binding to the contract
	GovernanceTransactor // Write-only binding to the contract
	GovernanceFilterer   // Log filterer for contract events
}

// GovernanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovernanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovernanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovernanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovernanceSession struct {
	Contract     *Governance       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovernanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovernanceCallerSession struct {
	Contract *GovernanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// GovernanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovernanceTransactorSession struct {
	Contract     *GovernanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GovernanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovernanceRaw struct {
	Contract *Governance // Generic contract binding to access the raw methods on
}

// GovernanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovernanceCallerRaw struct {
	Contract *GovernanceCaller // Generic read-only contract binding to access the raw methods on
}

// GovernanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovernanceTransactorRaw struct {
	Contract *GovernanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGovernance creates a new instance of Governance, bound to a specific deployed contract.
func NewGovernance(address common.Address, backend bind.ContractBackend) (*Governance, error) {
	contract, err := bindGovernance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Governance{GovernanceCaller: GovernanceCaller{contract: contract}, GovernanceTransactor: GovernanceTransactor{contract: contract}, GovernanceFilterer: GovernanceFilterer{contract: contract}}, nil
}

// NewGovernanceCaller creates a new read-only instance of Governance, bound to a specific deployed contract.
func NewGovernanceCaller(address common.Address, caller bind.ContractCaller) (*GovernanceCaller, error) {
	contract, err := bindGovernance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceCaller{contract: contract}, nil
}

// NewGovernanceTransactor creates a new write-only instance of Governance, bound to a specific deployed contract.
func NewGovernanceTransactor(address common.Address, transactor bind.ContractTransactor) (*GovernanceTransactor, error) {
	contract, err := bindGovernance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceTransactor{contract: contract}, nil
}

// NewGovernanceFilterer creates a new log filterer instance of Governance, bound to a specific deployed contract.
func NewGovernanceFilterer(address common.Address, filterer bind.ContractFilterer) (*GovernanceFilterer, error) {
	contract, err := bindGovernance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovernanceFilterer{contract: contract}, nil
}

// bindGovernance binds a generic wrapper to an already deployed contract.
func bindGovernance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Governance *GovernanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Governance.Contract.GovernanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Governance *GovernanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.Contract.GovernanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Governance *GovernanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Governance.Contract.GovernanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Governance *GovernanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Governance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Governance *GovernanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Governance *GovernanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Governance.Contract.contract.Transact(opts, method, params...)
}

// BALLOTTYPEHASH is a free data retrieval call binding the contract method 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (_Governance *GovernanceCaller) BALLOTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "BALLOT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BALLOTTYPEHASH is a free data retrieval call binding the contract method 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (_Governance *GovernanceSession) BALLOTTYPEHASH() ([32]byte, error) {
	return _Governance.Contract.BALLOTTYPEHASH(&_Governance.CallOpts)
}

// BALLOTTYPEHASH is a free data retrieval call binding the contract method 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (_Governance *GovernanceCallerSession) BALLOTTYPEHASH() ([32]byte, error) {
	return _Governance.Contract.BALLOTTYPEHASH(&_Governance.CallOpts)
}

// COUNTINGMODE is a free data retrieval call binding the contract method 0xdd4e2ba5.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (_Governance *GovernanceCaller) COUNTINGMODE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "COUNTING_MODE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// COUNTINGMODE is a free data retrieval call binding the contract method 0xdd4e2ba5.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (_Governance *GovernanceSession) COUNTINGMODE() (string, error) {
	return _Governance.Contract.COUNTINGMODE(&_Governance.CallOpts)
}

// COUNTINGMODE is a free data retrieval call binding the contract method 0xdd4e2ba5.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (_Governance *GovernanceCallerSession) COUNTINGMODE() (string, error) {
	return _Governance.Contract.COUNTINGMODE(&_Governance.CallOpts)
}

// EXTENDEDBALLOTTYPEHASH is a free data retrieval call binding the contract method 0x2fe3e261.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (_Governance *GovernanceCaller) EXTENDEDBALLOTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "EXTENDED_BALLOT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EXTENDEDBALLOTTYPEHASH is a free data retrieval call binding the contract method 0x2fe3e261.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (_Governance *GovernanceSession) EXTENDEDBALLOTTYPEHASH() ([32]byte, error) {
	return _Governance.Contract.EXTENDEDBALLOTTYPEHASH(&_Governance.CallOpts)
}

// EXTENDEDBALLOTTYPEHASH is a free data retrieval call binding the contract method 0x2fe3e261.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (_Governance *GovernanceCallerSession) EXTENDEDBALLOTTYPEHASH() ([32]byte, error) {
	return _Governance.Contract.EXTENDEDBALLOTTYPEHASH(&_Governance.CallOpts)
}

// PROPOSALHASH is a free data retrieval call binding the contract method 0xbdc9b04c.
//
// Solidity: function PROPOSAL_HASH() view returns(bytes32)
func (_Governance *GovernanceCaller) PROPOSALHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "PROPOSAL_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PROPOSALHASH is a free data retrieval call binding the contract method 0xbdc9b04c.
//
// Solidity: function PROPOSAL_HASH() view returns(bytes32)
func (_Governance *GovernanceSession) PROPOSALHASH() ([32]byte, error) {
	return _Governance.Contract.PROPOSALHASH(&_Governance.CallOpts)
}

// PROPOSALHASH is a free data retrieval call binding the contract method 0xbdc9b04c.
//
// Solidity: function PROPOSAL_HASH() view returns(bytes32)
func (_Governance *GovernanceCallerSession) PROPOSALHASH() ([32]byte, error) {
	return _Governance.Contract.PROPOSALHASH(&_Governance.CallOpts)
}

// GetActions is a free data retrieval call binding the contract method 0x328dd982.
//
// Solidity: function getActions(uint256 proposalId) view returns(address[] targets, uint256[] values, string[] signatures, bytes[] calldatas)
func (_Governance *GovernanceCaller) GetActions(opts *bind.CallOpts, proposalId *big.Int) (struct {
	Targets    []common.Address
	Values     []*big.Int
	Signatures []string
	Calldatas  [][]byte
}, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getActions", proposalId)

	outstruct := new(struct {
		Targets    []common.Address
		Values     []*big.Int
		Signatures []string
		Calldatas  [][]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Targets = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Values = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.Signatures = *abi.ConvertType(out[2], new([]string)).(*[]string)
	outstruct.Calldatas = *abi.ConvertType(out[3], new([][]byte)).(*[][]byte)

	return *outstruct, err

}

// GetActions is a free data retrieval call binding the contract method 0x328dd982.
//
// Solidity: function getActions(uint256 proposalId) view returns(address[] targets, uint256[] values, string[] signatures, bytes[] calldatas)
func (_Governance *GovernanceSession) GetActions(proposalId *big.Int) (struct {
	Targets    []common.Address
	Values     []*big.Int
	Signatures []string
	Calldatas  [][]byte
}, error) {
	return _Governance.Contract.GetActions(&_Governance.CallOpts, proposalId)
}

// GetActions is a free data retrieval call binding the contract method 0x328dd982.
//
// Solidity: function getActions(uint256 proposalId) view returns(address[] targets, uint256[] values, string[] signatures, bytes[] calldatas)
func (_Governance *GovernanceCallerSession) GetActions(proposalId *big.Int) (struct {
	Targets    []common.Address
	Values     []*big.Int
	Signatures []string
	Calldatas  [][]byte
}, error) {
	return _Governance.Contract.GetActions(&_Governance.CallOpts, proposalId)
}

// GetDataSign is a free data retrieval call binding the contract method 0xf7d25dc0.
//
// Solidity: function getDataSign(bytes32 _input) view returns(bytes32)
func (_Governance *GovernanceCaller) GetDataSign(opts *bind.CallOpts, _input [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getDataSign", _input)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDataSign is a free data retrieval call binding the contract method 0xf7d25dc0.
//
// Solidity: function getDataSign(bytes32 _input) view returns(bytes32)
func (_Governance *GovernanceSession) GetDataSign(_input [32]byte) ([32]byte, error) {
	return _Governance.Contract.GetDataSign(&_Governance.CallOpts, _input)
}

// GetDataSign is a free data retrieval call binding the contract method 0xf7d25dc0.
//
// Solidity: function getDataSign(bytes32 _input) view returns(bytes32)
func (_Governance *GovernanceCallerSession) GetDataSign(_input [32]byte) ([32]byte, error) {
	return _Governance.Contract.GetDataSign(&_Governance.CallOpts, _input)
}

// GetReceipt is a free data retrieval call binding the contract method 0xe23a9a52.
//
// Solidity: function getReceipt(uint256 proposalId, address voter) view returns((bool,uint8,uint96))
func (_Governance *GovernanceCaller) GetReceipt(opts *bind.CallOpts, proposalId *big.Int, voter common.Address) (IGovernorCompatibilityBravoUpgradeableReceipt, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getReceipt", proposalId, voter)

	if err != nil {
		return *new(IGovernorCompatibilityBravoUpgradeableReceipt), err
	}

	out0 := *abi.ConvertType(out[0], new(IGovernorCompatibilityBravoUpgradeableReceipt)).(*IGovernorCompatibilityBravoUpgradeableReceipt)

	return out0, err

}

// GetReceipt is a free data retrieval call binding the contract method 0xe23a9a52.
//
// Solidity: function getReceipt(uint256 proposalId, address voter) view returns((bool,uint8,uint96))
func (_Governance *GovernanceSession) GetReceipt(proposalId *big.Int, voter common.Address) (IGovernorCompatibilityBravoUpgradeableReceipt, error) {
	return _Governance.Contract.GetReceipt(&_Governance.CallOpts, proposalId, voter)
}

// GetReceipt is a free data retrieval call binding the contract method 0xe23a9a52.
//
// Solidity: function getReceipt(uint256 proposalId, address voter) view returns((bool,uint8,uint96))
func (_Governance *GovernanceCallerSession) GetReceipt(proposalId *big.Int, voter common.Address) (IGovernorCompatibilityBravoUpgradeableReceipt, error) {
	return _Governance.Contract.GetReceipt(&_Governance.CallOpts, proposalId, voter)
}

// GetVotes is a free data retrieval call binding the contract method 0xeb9019d4.
//
// Solidity: function getVotes(address account, uint256 blockNumber) view returns(uint256)
func (_Governance *GovernanceCaller) GetVotes(opts *bind.CallOpts, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getVotes", account, blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotes is a free data retrieval call binding the contract method 0xeb9019d4.
//
// Solidity: function getVotes(address account, uint256 blockNumber) view returns(uint256)
func (_Governance *GovernanceSession) GetVotes(account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return _Governance.Contract.GetVotes(&_Governance.CallOpts, account, blockNumber)
}

// GetVotes is a free data retrieval call binding the contract method 0xeb9019d4.
//
// Solidity: function getVotes(address account, uint256 blockNumber) view returns(uint256)
func (_Governance *GovernanceCallerSession) GetVotes(account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return _Governance.Contract.GetVotes(&_Governance.CallOpts, account, blockNumber)
}

// GetVotesWithParams is a free data retrieval call binding the contract method 0x9a802a6d.
//
// Solidity: function getVotesWithParams(address account, uint256 blockNumber, bytes params) view returns(uint256)
func (_Governance *GovernanceCaller) GetVotesWithParams(opts *bind.CallOpts, account common.Address, blockNumber *big.Int, params []byte) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getVotesWithParams", account, blockNumber, params)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotesWithParams is a free data retrieval call binding the contract method 0x9a802a6d.
//
// Solidity: function getVotesWithParams(address account, uint256 blockNumber, bytes params) view returns(uint256)
func (_Governance *GovernanceSession) GetVotesWithParams(account common.Address, blockNumber *big.Int, params []byte) (*big.Int, error) {
	return _Governance.Contract.GetVotesWithParams(&_Governance.CallOpts, account, blockNumber, params)
}

// GetVotesWithParams is a free data retrieval call binding the contract method 0x9a802a6d.
//
// Solidity: function getVotesWithParams(address account, uint256 blockNumber, bytes params) view returns(uint256)
func (_Governance *GovernanceCallerSession) GetVotesWithParams(account common.Address, blockNumber *big.Int, params []byte) (*big.Int, error) {
	return _Governance.Contract.GetVotesWithParams(&_Governance.CallOpts, account, blockNumber, params)
}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (_Governance *GovernanceCaller) HasVoted(opts *bind.CallOpts, proposalId *big.Int, account common.Address) (bool, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "hasVoted", proposalId, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (_Governance *GovernanceSession) HasVoted(proposalId *big.Int, account common.Address) (bool, error) {
	return _Governance.Contract.HasVoted(&_Governance.CallOpts, proposalId, account)
}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (_Governance *GovernanceCallerSession) HasVoted(proposalId *big.Int, account common.Address) (bool, error) {
	return _Governance.Contract.HasVoted(&_Governance.CallOpts, proposalId, account)
}

// HashProposal is a free data retrieval call binding the contract method 0xc59057e4.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (_Governance *GovernanceCaller) HashProposal(opts *bind.CallOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "hashProposal", targets, values, calldatas, descriptionHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HashProposal is a free data retrieval call binding the contract method 0xc59057e4.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (_Governance *GovernanceSession) HashProposal(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*big.Int, error) {
	return _Governance.Contract.HashProposal(&_Governance.CallOpts, targets, values, calldatas, descriptionHash)
}

// HashProposal is a free data retrieval call binding the contract method 0xc59057e4.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (_Governance *GovernanceCallerSession) HashProposal(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*big.Int, error) {
	return _Governance.Contract.HashProposal(&_Governance.CallOpts, targets, values, calldatas, descriptionHash)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Governance *GovernanceCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Governance *GovernanceSession) Name() (string, error) {
	return _Governance.Contract.Name(&_Governance.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Governance *GovernanceCallerSession) Name() (string, error) {
	return _Governance.Contract.Name(&_Governance.CallOpts)
}

// ProposalDeadline is a free data retrieval call binding the contract method 0xc01f9e37.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (_Governance *GovernanceCaller) ProposalDeadline(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "proposalDeadline", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalDeadline is a free data retrieval call binding the contract method 0xc01f9e37.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (_Governance *GovernanceSession) ProposalDeadline(proposalId *big.Int) (*big.Int, error) {
	return _Governance.Contract.ProposalDeadline(&_Governance.CallOpts, proposalId)
}

// ProposalDeadline is a free data retrieval call binding the contract method 0xc01f9e37.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (_Governance *GovernanceCallerSession) ProposalDeadline(proposalId *big.Int) (*big.Int, error) {
	return _Governance.Contract.ProposalDeadline(&_Governance.CallOpts, proposalId)
}

// ProposalEta is a free data retrieval call binding the contract method 0xab58fb8e.
//
// Solidity: function proposalEta(uint256 ) pure returns(uint256)
func (_Governance *GovernanceCaller) ProposalEta(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "proposalEta", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalEta is a free data retrieval call binding the contract method 0xab58fb8e.
//
// Solidity: function proposalEta(uint256 ) pure returns(uint256)
func (_Governance *GovernanceSession) ProposalEta(arg0 *big.Int) (*big.Int, error) {
	return _Governance.Contract.ProposalEta(&_Governance.CallOpts, arg0)
}

// ProposalEta is a free data retrieval call binding the contract method 0xab58fb8e.
//
// Solidity: function proposalEta(uint256 ) pure returns(uint256)
func (_Governance *GovernanceCallerSession) ProposalEta(arg0 *big.Int) (*big.Int, error) {
	return _Governance.Contract.ProposalEta(&_Governance.CallOpts, arg0)
}

// ProposalSnapshot is a free data retrieval call binding the contract method 0x2d63f693.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (_Governance *GovernanceCaller) ProposalSnapshot(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "proposalSnapshot", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalSnapshot is a free data retrieval call binding the contract method 0x2d63f693.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (_Governance *GovernanceSession) ProposalSnapshot(proposalId *big.Int) (*big.Int, error) {
	return _Governance.Contract.ProposalSnapshot(&_Governance.CallOpts, proposalId)
}

// ProposalSnapshot is a free data retrieval call binding the contract method 0x2d63f693.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (_Governance *GovernanceCallerSession) ProposalSnapshot(proposalId *big.Int) (*big.Int, error) {
	return _Governance.Contract.ProposalSnapshot(&_Governance.CallOpts, proposalId)
}

// ProposalThreshold is a free data retrieval call binding the contract method 0xb58131b0.
//
// Solidity: function proposalThreshold() pure returns(uint256)
func (_Governance *GovernanceCaller) ProposalThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "proposalThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalThreshold is a free data retrieval call binding the contract method 0xb58131b0.
//
// Solidity: function proposalThreshold() pure returns(uint256)
func (_Governance *GovernanceSession) ProposalThreshold() (*big.Int, error) {
	return _Governance.Contract.ProposalThreshold(&_Governance.CallOpts)
}

// ProposalThreshold is a free data retrieval call binding the contract method 0xb58131b0.
//
// Solidity: function proposalThreshold() pure returns(uint256)
func (_Governance *GovernanceCallerSession) ProposalThreshold() (*big.Int, error) {
	return _Governance.Contract.ProposalThreshold(&_Governance.CallOpts)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 proposalId) view returns(uint256 id, address proposer, uint256 eta, uint256 startBlock, uint256 endBlock, uint256 forVotes, uint256 againstVotes, uint256 abstainVotes, bool canceled, bool executed)
func (_Governance *GovernanceCaller) Proposals(opts *bind.CallOpts, proposalId *big.Int) (struct {
	Id           *big.Int
	Proposer     common.Address
	Eta          *big.Int
	StartBlock   *big.Int
	EndBlock     *big.Int
	ForVotes     *big.Int
	AgainstVotes *big.Int
	AbstainVotes *big.Int
	Canceled     bool
	Executed     bool
}, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "proposals", proposalId)

	outstruct := new(struct {
		Id           *big.Int
		Proposer     common.Address
		Eta          *big.Int
		StartBlock   *big.Int
		EndBlock     *big.Int
		ForVotes     *big.Int
		AgainstVotes *big.Int
		AbstainVotes *big.Int
		Canceled     bool
		Executed     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Proposer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Eta = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StartBlock = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.EndBlock = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.ForVotes = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.AgainstVotes = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.AbstainVotes = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Canceled = *abi.ConvertType(out[8], new(bool)).(*bool)
	outstruct.Executed = *abi.ConvertType(out[9], new(bool)).(*bool)

	return *outstruct, err

}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 proposalId) view returns(uint256 id, address proposer, uint256 eta, uint256 startBlock, uint256 endBlock, uint256 forVotes, uint256 againstVotes, uint256 abstainVotes, bool canceled, bool executed)
func (_Governance *GovernanceSession) Proposals(proposalId *big.Int) (struct {
	Id           *big.Int
	Proposer     common.Address
	Eta          *big.Int
	StartBlock   *big.Int
	EndBlock     *big.Int
	ForVotes     *big.Int
	AgainstVotes *big.Int
	AbstainVotes *big.Int
	Canceled     bool
	Executed     bool
}, error) {
	return _Governance.Contract.Proposals(&_Governance.CallOpts, proposalId)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 proposalId) view returns(uint256 id, address proposer, uint256 eta, uint256 startBlock, uint256 endBlock, uint256 forVotes, uint256 againstVotes, uint256 abstainVotes, bool canceled, bool executed)
func (_Governance *GovernanceCallerSession) Proposals(proposalId *big.Int) (struct {
	Id           *big.Int
	Proposer     common.Address
	Eta          *big.Int
	StartBlock   *big.Int
	EndBlock     *big.Int
	ForVotes     *big.Int
	AgainstVotes *big.Int
	AbstainVotes *big.Int
	Canceled     bool
	Executed     bool
}, error) {
	return _Governance.Contract.Proposals(&_Governance.CallOpts, proposalId)
}

// Queue is a free data retrieval call binding the contract method 0x160cbed7.
//
// Solidity: function queue(address[] , uint256[] , bytes[] , bytes32 ) pure returns(uint256)
func (_Governance *GovernanceCaller) Queue(opts *bind.CallOpts, arg0 []common.Address, arg1 []*big.Int, arg2 [][]byte, arg3 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "queue", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Queue is a free data retrieval call binding the contract method 0x160cbed7.
//
// Solidity: function queue(address[] , uint256[] , bytes[] , bytes32 ) pure returns(uint256)
func (_Governance *GovernanceSession) Queue(arg0 []common.Address, arg1 []*big.Int, arg2 [][]byte, arg3 [32]byte) (*big.Int, error) {
	return _Governance.Contract.Queue(&_Governance.CallOpts, arg0, arg1, arg2, arg3)
}

// Queue is a free data retrieval call binding the contract method 0x160cbed7.
//
// Solidity: function queue(address[] , uint256[] , bytes[] , bytes32 ) pure returns(uint256)
func (_Governance *GovernanceCallerSession) Queue(arg0 []common.Address, arg1 []*big.Int, arg2 [][]byte, arg3 [32]byte) (*big.Int, error) {
	return _Governance.Contract.Queue(&_Governance.CallOpts, arg0, arg1, arg2, arg3)
}

// Quorum is a free data retrieval call binding the contract method 0xf8ce560a.
//
// Solidity: function quorum(uint256 blockNumber) view returns(uint256)
func (_Governance *GovernanceCaller) Quorum(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "quorum", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quorum is a free data retrieval call binding the contract method 0xf8ce560a.
//
// Solidity: function quorum(uint256 blockNumber) view returns(uint256)
func (_Governance *GovernanceSession) Quorum(blockNumber *big.Int) (*big.Int, error) {
	return _Governance.Contract.Quorum(&_Governance.CallOpts, blockNumber)
}

// Quorum is a free data retrieval call binding the contract method 0xf8ce560a.
//
// Solidity: function quorum(uint256 blockNumber) view returns(uint256)
func (_Governance *GovernanceCallerSession) Quorum(blockNumber *big.Int) (*big.Int, error) {
	return _Governance.Contract.Quorum(&_Governance.CallOpts, blockNumber)
}

// QuorumDenominator is a free data retrieval call binding the contract method 0x97c3d334.
//
// Solidity: function quorumDenominator() view returns(uint256)
func (_Governance *GovernanceCaller) QuorumDenominator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "quorumDenominator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumDenominator is a free data retrieval call binding the contract method 0x97c3d334.
//
// Solidity: function quorumDenominator() view returns(uint256)
func (_Governance *GovernanceSession) QuorumDenominator() (*big.Int, error) {
	return _Governance.Contract.QuorumDenominator(&_Governance.CallOpts)
}

// QuorumDenominator is a free data retrieval call binding the contract method 0x97c3d334.
//
// Solidity: function quorumDenominator() view returns(uint256)
func (_Governance *GovernanceCallerSession) QuorumDenominator() (*big.Int, error) {
	return _Governance.Contract.QuorumDenominator(&_Governance.CallOpts)
}

// QuorumNumerator is a free data retrieval call binding the contract method 0x60c4247f.
//
// Solidity: function quorumNumerator(uint256 blockNumber) view returns(uint256)
func (_Governance *GovernanceCaller) QuorumNumerator(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "quorumNumerator", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumNumerator is a free data retrieval call binding the contract method 0x60c4247f.
//
// Solidity: function quorumNumerator(uint256 blockNumber) view returns(uint256)
func (_Governance *GovernanceSession) QuorumNumerator(blockNumber *big.Int) (*big.Int, error) {
	return _Governance.Contract.QuorumNumerator(&_Governance.CallOpts, blockNumber)
}

// QuorumNumerator is a free data retrieval call binding the contract method 0x60c4247f.
//
// Solidity: function quorumNumerator(uint256 blockNumber) view returns(uint256)
func (_Governance *GovernanceCallerSession) QuorumNumerator(blockNumber *big.Int) (*big.Int, error) {
	return _Governance.Contract.QuorumNumerator(&_Governance.CallOpts, blockNumber)
}

// QuorumNumerator0 is a free data retrieval call binding the contract method 0xa7713a70.
//
// Solidity: function quorumNumerator() view returns(uint256)
func (_Governance *GovernanceCaller) QuorumNumerator0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "quorumNumerator0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumNumerator0 is a free data retrieval call binding the contract method 0xa7713a70.
//
// Solidity: function quorumNumerator() view returns(uint256)
func (_Governance *GovernanceSession) QuorumNumerator0() (*big.Int, error) {
	return _Governance.Contract.QuorumNumerator0(&_Governance.CallOpts)
}

// QuorumNumerator0 is a free data retrieval call binding the contract method 0xa7713a70.
//
// Solidity: function quorumNumerator() view returns(uint256)
func (_Governance *GovernanceCallerSession) QuorumNumerator0() (*big.Int, error) {
	return _Governance.Contract.QuorumNumerator0(&_Governance.CallOpts)
}

// QuorumVotes is a free data retrieval call binding the contract method 0x24bc1a64.
//
// Solidity: function quorumVotes() view returns(uint256)
func (_Governance *GovernanceCaller) QuorumVotes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "quorumVotes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumVotes is a free data retrieval call binding the contract method 0x24bc1a64.
//
// Solidity: function quorumVotes() view returns(uint256)
func (_Governance *GovernanceSession) QuorumVotes() (*big.Int, error) {
	return _Governance.Contract.QuorumVotes(&_Governance.CallOpts)
}

// QuorumVotes is a free data retrieval call binding the contract method 0x24bc1a64.
//
// Solidity: function quorumVotes() view returns(uint256)
func (_Governance *GovernanceCallerSession) QuorumVotes() (*big.Int, error) {
	return _Governance.Contract.QuorumVotes(&_Governance.CallOpts)
}

// SigDataUsed is a free data retrieval call binding the contract method 0x1ea1940e.
//
// Solidity: function sigDataUsed(bytes32 ) view returns(bool)
func (_Governance *GovernanceCaller) SigDataUsed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "sigDataUsed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SigDataUsed is a free data retrieval call binding the contract method 0x1ea1940e.
//
// Solidity: function sigDataUsed(bytes32 ) view returns(bool)
func (_Governance *GovernanceSession) SigDataUsed(arg0 [32]byte) (bool, error) {
	return _Governance.Contract.SigDataUsed(&_Governance.CallOpts, arg0)
}

// SigDataUsed is a free data retrieval call binding the contract method 0x1ea1940e.
//
// Solidity: function sigDataUsed(bytes32 ) view returns(bool)
func (_Governance *GovernanceCallerSession) SigDataUsed(arg0 [32]byte) (bool, error) {
	return _Governance.Contract.SigDataUsed(&_Governance.CallOpts, arg0)
}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_Governance *GovernanceCaller) State(opts *bind.CallOpts, proposalId *big.Int) (uint8, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "state", proposalId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_Governance *GovernanceSession) State(proposalId *big.Int) (uint8, error) {
	return _Governance.Contract.State(&_Governance.CallOpts, proposalId)
}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_Governance *GovernanceCallerSession) State(proposalId *big.Int) (uint8, error) {
	return _Governance.Contract.State(&_Governance.CallOpts, proposalId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Governance *GovernanceCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Governance *GovernanceSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Governance.Contract.SupportsInterface(&_Governance.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Governance *GovernanceCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Governance.Contract.SupportsInterface(&_Governance.CallOpts, interfaceId)
}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() pure returns(address)
func (_Governance *GovernanceCaller) Timelock(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "timelock")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() pure returns(address)
func (_Governance *GovernanceSession) Timelock() (common.Address, error) {
	return _Governance.Contract.Timelock(&_Governance.CallOpts)
}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() pure returns(address)
func (_Governance *GovernanceCallerSession) Timelock() (common.Address, error) {
	return _Governance.Contract.Timelock(&_Governance.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Governance *GovernanceCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Governance *GovernanceSession) Token() (common.Address, error) {
	return _Governance.Contract.Token(&_Governance.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Governance *GovernanceCallerSession) Token() (common.Address, error) {
	return _Governance.Contract.Token(&_Governance.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Governance *GovernanceCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Governance *GovernanceSession) Version() (string, error) {
	return _Governance.Contract.Version(&_Governance.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Governance *GovernanceCallerSession) Version() (string, error) {
	return _Governance.Contract.Version(&_Governance.CallOpts)
}

// VotingDelay is a free data retrieval call binding the contract method 0x3932abb1.
//
// Solidity: function votingDelay() pure returns(uint256)
func (_Governance *GovernanceCaller) VotingDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "votingDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingDelay is a free data retrieval call binding the contract method 0x3932abb1.
//
// Solidity: function votingDelay() pure returns(uint256)
func (_Governance *GovernanceSession) VotingDelay() (*big.Int, error) {
	return _Governance.Contract.VotingDelay(&_Governance.CallOpts)
}

// VotingDelay is a free data retrieval call binding the contract method 0x3932abb1.
//
// Solidity: function votingDelay() pure returns(uint256)
func (_Governance *GovernanceCallerSession) VotingDelay() (*big.Int, error) {
	return _Governance.Contract.VotingDelay(&_Governance.CallOpts)
}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() pure returns(uint256)
func (_Governance *GovernanceCaller) VotingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "votingPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() pure returns(uint256)
func (_Governance *GovernanceSession) VotingPeriod() (*big.Int, error) {
	return _Governance.Contract.VotingPeriod(&_Governance.CallOpts)
}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() pure returns(uint256)
func (_Governance *GovernanceCallerSession) VotingPeriod() (*big.Int, error) {
	return _Governance.Contract.VotingPeriod(&_Governance.CallOpts)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 proposalId) returns()
func (_Governance *GovernanceTransactor) Cancel(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "cancel", proposalId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 proposalId) returns()
func (_Governance *GovernanceSession) Cancel(proposalId *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.Cancel(&_Governance.TransactOpts, proposalId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 proposalId) returns()
func (_Governance *GovernanceTransactorSession) Cancel(proposalId *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.Cancel(&_Governance.TransactOpts, proposalId)
}

// CancelBySig is a paid mutator transaction binding the contract method 0x20e297ac.
//
// Solidity: function cancelBySig(uint256 proposalId, uint8 v, bytes32 r, bytes32 s) returns()
func (_Governance *GovernanceTransactor) CancelBySig(opts *bind.TransactOpts, proposalId *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "cancelBySig", proposalId, v, r, s)
}

// CancelBySig is a paid mutator transaction binding the contract method 0x20e297ac.
//
// Solidity: function cancelBySig(uint256 proposalId, uint8 v, bytes32 r, bytes32 s) returns()
func (_Governance *GovernanceSession) CancelBySig(proposalId *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Governance.Contract.CancelBySig(&_Governance.TransactOpts, proposalId, v, r, s)
}

// CancelBySig is a paid mutator transaction binding the contract method 0x20e297ac.
//
// Solidity: function cancelBySig(uint256 proposalId, uint8 v, bytes32 r, bytes32 s) returns()
func (_Governance *GovernanceTransactorSession) CancelBySig(proposalId *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Governance.Contract.CancelBySig(&_Governance.TransactOpts, proposalId, v, r, s)
}

// CastVote is a paid mutator transaction binding the contract method 0x56781388.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (_Governance *GovernanceTransactor) CastVote(opts *bind.TransactOpts, proposalId *big.Int, support uint8) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "castVote", proposalId, support)
}

// CastVote is a paid mutator transaction binding the contract method 0x56781388.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (_Governance *GovernanceSession) CastVote(proposalId *big.Int, support uint8) (*types.Transaction, error) {
	return _Governance.Contract.CastVote(&_Governance.TransactOpts, proposalId, support)
}

// CastVote is a paid mutator transaction binding the contract method 0x56781388.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (_Governance *GovernanceTransactorSession) CastVote(proposalId *big.Int, support uint8) (*types.Transaction, error) {
	return _Governance.Contract.CastVote(&_Governance.TransactOpts, proposalId, support)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x3bccf4fd.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, uint8 v, bytes32 r, bytes32 s) returns(uint256)
func (_Governance *GovernanceTransactor) CastVoteBySig(opts *bind.TransactOpts, proposalId *big.Int, support uint8, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "castVoteBySig", proposalId, support, v, r, s)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x3bccf4fd.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, uint8 v, bytes32 r, bytes32 s) returns(uint256)
func (_Governance *GovernanceSession) CastVoteBySig(proposalId *big.Int, support uint8, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Governance.Contract.CastVoteBySig(&_Governance.TransactOpts, proposalId, support, v, r, s)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x3bccf4fd.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, uint8 v, bytes32 r, bytes32 s) returns(uint256)
func (_Governance *GovernanceTransactorSession) CastVoteBySig(proposalId *big.Int, support uint8, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Governance.Contract.CastVoteBySig(&_Governance.TransactOpts, proposalId, support, v, r, s)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x7b3c71d3.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (_Governance *GovernanceTransactor) CastVoteWithReason(opts *bind.TransactOpts, proposalId *big.Int, support uint8, reason string) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "castVoteWithReason", proposalId, support, reason)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x7b3c71d3.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (_Governance *GovernanceSession) CastVoteWithReason(proposalId *big.Int, support uint8, reason string) (*types.Transaction, error) {
	return _Governance.Contract.CastVoteWithReason(&_Governance.TransactOpts, proposalId, support, reason)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x7b3c71d3.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (_Governance *GovernanceTransactorSession) CastVoteWithReason(proposalId *big.Int, support uint8, reason string) (*types.Transaction, error) {
	return _Governance.Contract.CastVoteWithReason(&_Governance.TransactOpts, proposalId, support, reason)
}

// CastVoteWithReasonAndParams is a paid mutator transaction binding the contract method 0x5f398a14.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (_Governance *GovernanceTransactor) CastVoteWithReasonAndParams(opts *bind.TransactOpts, proposalId *big.Int, support uint8, reason string, params []byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "castVoteWithReasonAndParams", proposalId, support, reason, params)
}

// CastVoteWithReasonAndParams is a paid mutator transaction binding the contract method 0x5f398a14.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (_Governance *GovernanceSession) CastVoteWithReasonAndParams(proposalId *big.Int, support uint8, reason string, params []byte) (*types.Transaction, error) {
	return _Governance.Contract.CastVoteWithReasonAndParams(&_Governance.TransactOpts, proposalId, support, reason, params)
}

// CastVoteWithReasonAndParams is a paid mutator transaction binding the contract method 0x5f398a14.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (_Governance *GovernanceTransactorSession) CastVoteWithReasonAndParams(proposalId *big.Int, support uint8, reason string, params []byte) (*types.Transaction, error) {
	return _Governance.Contract.CastVoteWithReasonAndParams(&_Governance.TransactOpts, proposalId, support, reason, params)
}

// CastVoteWithReasonAndParamsBySig is a paid mutator transaction binding the contract method 0x03420181.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, string reason, bytes params, uint8 v, bytes32 r, bytes32 s) returns(uint256)
func (_Governance *GovernanceTransactor) CastVoteWithReasonAndParamsBySig(opts *bind.TransactOpts, proposalId *big.Int, support uint8, reason string, params []byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "castVoteWithReasonAndParamsBySig", proposalId, support, reason, params, v, r, s)
}

// CastVoteWithReasonAndParamsBySig is a paid mutator transaction binding the contract method 0x03420181.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, string reason, bytes params, uint8 v, bytes32 r, bytes32 s) returns(uint256)
func (_Governance *GovernanceSession) CastVoteWithReasonAndParamsBySig(proposalId *big.Int, support uint8, reason string, params []byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Governance.Contract.CastVoteWithReasonAndParamsBySig(&_Governance.TransactOpts, proposalId, support, reason, params, v, r, s)
}

// CastVoteWithReasonAndParamsBySig is a paid mutator transaction binding the contract method 0x03420181.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, string reason, bytes params, uint8 v, bytes32 r, bytes32 s) returns(uint256)
func (_Governance *GovernanceTransactorSession) CastVoteWithReasonAndParamsBySig(proposalId *big.Int, support uint8, reason string, params []byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Governance.Contract.CastVoteWithReasonAndParamsBySig(&_Governance.TransactOpts, proposalId, support, reason, params, v, r, s)
}

// Execute is a paid mutator transaction binding the contract method 0x2656227d.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (_Governance *GovernanceTransactor) Execute(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "execute", targets, values, calldatas, descriptionHash)
}

// Execute is a paid mutator transaction binding the contract method 0x2656227d.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (_Governance *GovernanceSession) Execute(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _Governance.Contract.Execute(&_Governance.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Execute is a paid mutator transaction binding the contract method 0x2656227d.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (_Governance *GovernanceTransactorSession) Execute(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _Governance.Contract.Execute(&_Governance.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Execute0 is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 proposalId) payable returns()
func (_Governance *GovernanceTransactor) Execute0(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "execute0", proposalId)
}

// Execute0 is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 proposalId) payable returns()
func (_Governance *GovernanceSession) Execute0(proposalId *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.Execute0(&_Governance.TransactOpts, proposalId)
}

// Execute0 is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 proposalId) payable returns()
func (_Governance *GovernanceTransactorSession) Execute0(proposalId *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.Execute0(&_Governance.TransactOpts, proposalId)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _token) returns()
func (_Governance *GovernanceTransactor) Initialize(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "initialize", _token)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _token) returns()
func (_Governance *GovernanceSession) Initialize(_token common.Address) (*types.Transaction, error) {
	return _Governance.Contract.Initialize(&_Governance.TransactOpts, _token)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _token) returns()
func (_Governance *GovernanceTransactorSession) Initialize(_token common.Address) (*types.Transaction, error) {
	return _Governance.Contract.Initialize(&_Governance.TransactOpts, _token)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Governance *GovernanceTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Governance *GovernanceSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Governance.Contract.OnERC1155BatchReceived(&_Governance.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Governance *GovernanceTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Governance.Contract.OnERC1155BatchReceived(&_Governance.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Governance *GovernanceTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Governance *GovernanceSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Governance.Contract.OnERC1155Received(&_Governance.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Governance *GovernanceTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Governance.Contract.OnERC1155Received(&_Governance.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Governance *GovernanceTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Governance *GovernanceSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Governance.Contract.OnERC721Received(&_Governance.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Governance *GovernanceTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Governance.Contract.OnERC721Received(&_Governance.TransactOpts, arg0, arg1, arg2, arg3)
}

// Propose is a paid mutator transaction binding the contract method 0x7d5e81e2.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (_Governance *GovernanceTransactor) Propose(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "propose", targets, values, calldatas, description)
}

// Propose is a paid mutator transaction binding the contract method 0x7d5e81e2.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (_Governance *GovernanceSession) Propose(targets []common.Address, values []*big.Int, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _Governance.Contract.Propose(&_Governance.TransactOpts, targets, values, calldatas, description)
}

// Propose is a paid mutator transaction binding the contract method 0x7d5e81e2.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (_Governance *GovernanceTransactorSession) Propose(targets []common.Address, values []*big.Int, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _Governance.Contract.Propose(&_Governance.TransactOpts, targets, values, calldatas, description)
}

// Propose0 is a paid mutator transaction binding the contract method 0xda95691a.
//
// Solidity: function propose(address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, string description) returns(uint256)
func (_Governance *GovernanceTransactor) Propose0(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, signatures []string, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "propose0", targets, values, signatures, calldatas, description)
}

// Propose0 is a paid mutator transaction binding the contract method 0xda95691a.
//
// Solidity: function propose(address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, string description) returns(uint256)
func (_Governance *GovernanceSession) Propose0(targets []common.Address, values []*big.Int, signatures []string, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _Governance.Contract.Propose0(&_Governance.TransactOpts, targets, values, signatures, calldatas, description)
}

// Propose0 is a paid mutator transaction binding the contract method 0xda95691a.
//
// Solidity: function propose(address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, string description) returns(uint256)
func (_Governance *GovernanceTransactorSession) Propose0(targets []common.Address, values []*big.Int, signatures []string, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _Governance.Contract.Propose0(&_Governance.TransactOpts, targets, values, signatures, calldatas, description)
}

// ProposeBySig is a paid mutator transaction binding the contract method 0x8783ef44.
//
// Solidity: function proposeBySig(address[] targets, uint256[] values, bytes[] calldatas, string description, uint8 v, bytes32 r, bytes32 s) returns(uint256)
func (_Governance *GovernanceTransactor) ProposeBySig(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, description string, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "proposeBySig", targets, values, calldatas, description, v, r, s)
}

// ProposeBySig is a paid mutator transaction binding the contract method 0x8783ef44.
//
// Solidity: function proposeBySig(address[] targets, uint256[] values, bytes[] calldatas, string description, uint8 v, bytes32 r, bytes32 s) returns(uint256)
func (_Governance *GovernanceSession) ProposeBySig(targets []common.Address, values []*big.Int, calldatas [][]byte, description string, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Governance.Contract.ProposeBySig(&_Governance.TransactOpts, targets, values, calldatas, description, v, r, s)
}

// ProposeBySig is a paid mutator transaction binding the contract method 0x8783ef44.
//
// Solidity: function proposeBySig(address[] targets, uint256[] values, bytes[] calldatas, string description, uint8 v, bytes32 r, bytes32 s) returns(uint256)
func (_Governance *GovernanceTransactorSession) ProposeBySig(targets []common.Address, values []*big.Int, calldatas [][]byte, description string, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Governance.Contract.ProposeBySig(&_Governance.TransactOpts, targets, values, calldatas, description, v, r, s)
}

// Queue0 is a paid mutator transaction binding the contract method 0xddf0b009.
//
// Solidity: function queue(uint256 proposalId) returns()
func (_Governance *GovernanceTransactor) Queue0(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "queue0", proposalId)
}

// Queue0 is a paid mutator transaction binding the contract method 0xddf0b009.
//
// Solidity: function queue(uint256 proposalId) returns()
func (_Governance *GovernanceSession) Queue0(proposalId *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.Queue0(&_Governance.TransactOpts, proposalId)
}

// Queue0 is a paid mutator transaction binding the contract method 0xddf0b009.
//
// Solidity: function queue(uint256 proposalId) returns()
func (_Governance *GovernanceTransactorSession) Queue0(proposalId *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.Queue0(&_Governance.TransactOpts, proposalId)
}

// Relay is a paid mutator transaction binding the contract method 0xc28bc2fa.
//
// Solidity: function relay(address target, uint256 value, bytes data) payable returns()
func (_Governance *GovernanceTransactor) Relay(opts *bind.TransactOpts, target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "relay", target, value, data)
}

// Relay is a paid mutator transaction binding the contract method 0xc28bc2fa.
//
// Solidity: function relay(address target, uint256 value, bytes data) payable returns()
func (_Governance *GovernanceSession) Relay(target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Governance.Contract.Relay(&_Governance.TransactOpts, target, value, data)
}

// Relay is a paid mutator transaction binding the contract method 0xc28bc2fa.
//
// Solidity: function relay(address target, uint256 value, bytes data) payable returns()
func (_Governance *GovernanceTransactorSession) Relay(target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Governance.Contract.Relay(&_Governance.TransactOpts, target, value, data)
}

// UpdateQuorumNumerator is a paid mutator transaction binding the contract method 0x06f3f9e6.
//
// Solidity: function updateQuorumNumerator(uint256 newQuorumNumerator) returns()
func (_Governance *GovernanceTransactor) UpdateQuorumNumerator(opts *bind.TransactOpts, newQuorumNumerator *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "updateQuorumNumerator", newQuorumNumerator)
}

// UpdateQuorumNumerator is a paid mutator transaction binding the contract method 0x06f3f9e6.
//
// Solidity: function updateQuorumNumerator(uint256 newQuorumNumerator) returns()
func (_Governance *GovernanceSession) UpdateQuorumNumerator(newQuorumNumerator *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.UpdateQuorumNumerator(&_Governance.TransactOpts, newQuorumNumerator)
}

// UpdateQuorumNumerator is a paid mutator transaction binding the contract method 0x06f3f9e6.
//
// Solidity: function updateQuorumNumerator(uint256 newQuorumNumerator) returns()
func (_Governance *GovernanceTransactorSession) UpdateQuorumNumerator(newQuorumNumerator *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.UpdateQuorumNumerator(&_Governance.TransactOpts, newQuorumNumerator)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Governance *GovernanceTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Governance *GovernanceSession) Receive() (*types.Transaction, error) {
	return _Governance.Contract.Receive(&_Governance.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Governance *GovernanceTransactorSession) Receive() (*types.Transaction, error) {
	return _Governance.Contract.Receive(&_Governance.TransactOpts)
}

// GovernanceInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Governance contract.
type GovernanceInitializedIterator struct {
	Event *GovernanceInitialized // Event containing the contract specifics and raw log

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
func (it *GovernanceInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceInitialized)
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
		it.Event = new(GovernanceInitialized)
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
func (it *GovernanceInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceInitialized represents a Initialized event raised by the Governance contract.
type GovernanceInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Governance *GovernanceFilterer) FilterInitialized(opts *bind.FilterOpts) (*GovernanceInitializedIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GovernanceInitializedIterator{contract: _Governance.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Governance *GovernanceFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GovernanceInitialized) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceInitialized)
				if err := _Governance.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Governance *GovernanceFilterer) ParseInitialized(log types.Log) (*GovernanceInitialized, error) {
	event := new(GovernanceInitialized)
	if err := _Governance.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceProposalCanceledIterator is returned from FilterProposalCanceled and is used to iterate over the raw logs and unpacked data for ProposalCanceled events raised by the Governance contract.
type GovernanceProposalCanceledIterator struct {
	Event *GovernanceProposalCanceled // Event containing the contract specifics and raw log

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
func (it *GovernanceProposalCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalCanceled)
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
		it.Event = new(GovernanceProposalCanceled)
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
func (it *GovernanceProposalCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalCanceled represents a ProposalCanceled event raised by the Governance contract.
type GovernanceProposalCanceled struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalCanceled is a free log retrieval operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_Governance *GovernanceFilterer) FilterProposalCanceled(opts *bind.FilterOpts) (*GovernanceProposalCanceledIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalCanceled")
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalCanceledIterator{contract: _Governance.contract, event: "ProposalCanceled", logs: logs, sub: sub}, nil
}

// WatchProposalCanceled is a free log subscription operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_Governance *GovernanceFilterer) WatchProposalCanceled(opts *bind.WatchOpts, sink chan<- *GovernanceProposalCanceled) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalCanceled)
				if err := _Governance.contract.UnpackLog(event, "ProposalCanceled", log); err != nil {
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

// ParseProposalCanceled is a log parse operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_Governance *GovernanceFilterer) ParseProposalCanceled(log types.Log) (*GovernanceProposalCanceled, error) {
	event := new(GovernanceProposalCanceled)
	if err := _Governance.contract.UnpackLog(event, "ProposalCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the Governance contract.
type GovernanceProposalCreatedIterator struct {
	Event *GovernanceProposalCreated // Event containing the contract specifics and raw log

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
func (it *GovernanceProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalCreated)
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
		it.Event = new(GovernanceProposalCreated)
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
func (it *GovernanceProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalCreated represents a ProposalCreated event raised by the Governance contract.
type GovernanceProposalCreated struct {
	ProposalId  *big.Int
	Proposer    common.Address
	Targets     []common.Address
	Values      []*big.Int
	Signatures  []string
	Calldatas   [][]byte
	StartBlock  *big.Int
	EndBlock    *big.Int
	Description string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0x7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e0.
//
// Solidity: event ProposalCreated(uint256 proposalId, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 startBlock, uint256 endBlock, string description)
func (_Governance *GovernanceFilterer) FilterProposalCreated(opts *bind.FilterOpts) (*GovernanceProposalCreatedIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalCreatedIterator{contract: _Governance.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0x7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e0.
//
// Solidity: event ProposalCreated(uint256 proposalId, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 startBlock, uint256 endBlock, string description)
func (_Governance *GovernanceFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *GovernanceProposalCreated) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalCreated)
				if err := _Governance.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
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

// ParseProposalCreated is a log parse operation binding the contract event 0x7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e0.
//
// Solidity: event ProposalCreated(uint256 proposalId, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 startBlock, uint256 endBlock, string description)
func (_Governance *GovernanceFilterer) ParseProposalCreated(log types.Log) (*GovernanceProposalCreated, error) {
	event := new(GovernanceProposalCreated)
	if err := _Governance.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the Governance contract.
type GovernanceProposalExecutedIterator struct {
	Event *GovernanceProposalExecuted // Event containing the contract specifics and raw log

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
func (it *GovernanceProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalExecuted)
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
		it.Event = new(GovernanceProposalExecuted)
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
func (it *GovernanceProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalExecuted represents a ProposalExecuted event raised by the Governance contract.
type GovernanceProposalExecuted struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_Governance *GovernanceFilterer) FilterProposalExecuted(opts *bind.FilterOpts) (*GovernanceProposalExecutedIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalExecutedIterator{contract: _Governance.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_Governance *GovernanceFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *GovernanceProposalExecuted) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalExecuted)
				if err := _Governance.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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

// ParseProposalExecuted is a log parse operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_Governance *GovernanceFilterer) ParseProposalExecuted(log types.Log) (*GovernanceProposalExecuted, error) {
	event := new(GovernanceProposalExecuted)
	if err := _Governance.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceProposalQueuedIterator is returned from FilterProposalQueued and is used to iterate over the raw logs and unpacked data for ProposalQueued events raised by the Governance contract.
type GovernanceProposalQueuedIterator struct {
	Event *GovernanceProposalQueued // Event containing the contract specifics and raw log

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
func (it *GovernanceProposalQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalQueued)
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
		it.Event = new(GovernanceProposalQueued)
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
func (it *GovernanceProposalQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalQueued represents a ProposalQueued event raised by the Governance contract.
type GovernanceProposalQueued struct {
	ProposalId *big.Int
	Eta        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalQueued is a free log retrieval operation binding the contract event 0x9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892.
//
// Solidity: event ProposalQueued(uint256 proposalId, uint256 eta)
func (_Governance *GovernanceFilterer) FilterProposalQueued(opts *bind.FilterOpts) (*GovernanceProposalQueuedIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalQueued")
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalQueuedIterator{contract: _Governance.contract, event: "ProposalQueued", logs: logs, sub: sub}, nil
}

// WatchProposalQueued is a free log subscription operation binding the contract event 0x9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892.
//
// Solidity: event ProposalQueued(uint256 proposalId, uint256 eta)
func (_Governance *GovernanceFilterer) WatchProposalQueued(opts *bind.WatchOpts, sink chan<- *GovernanceProposalQueued) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalQueued)
				if err := _Governance.contract.UnpackLog(event, "ProposalQueued", log); err != nil {
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

// ParseProposalQueued is a log parse operation binding the contract event 0x9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892.
//
// Solidity: event ProposalQueued(uint256 proposalId, uint256 eta)
func (_Governance *GovernanceFilterer) ParseProposalQueued(log types.Log) (*GovernanceProposalQueued, error) {
	event := new(GovernanceProposalQueued)
	if err := _Governance.contract.UnpackLog(event, "ProposalQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceQuorumNumeratorUpdatedIterator is returned from FilterQuorumNumeratorUpdated and is used to iterate over the raw logs and unpacked data for QuorumNumeratorUpdated events raised by the Governance contract.
type GovernanceQuorumNumeratorUpdatedIterator struct {
	Event *GovernanceQuorumNumeratorUpdated // Event containing the contract specifics and raw log

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
func (it *GovernanceQuorumNumeratorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceQuorumNumeratorUpdated)
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
		it.Event = new(GovernanceQuorumNumeratorUpdated)
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
func (it *GovernanceQuorumNumeratorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceQuorumNumeratorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceQuorumNumeratorUpdated represents a QuorumNumeratorUpdated event raised by the Governance contract.
type GovernanceQuorumNumeratorUpdated struct {
	OldQuorumNumerator *big.Int
	NewQuorumNumerator *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterQuorumNumeratorUpdated is a free log retrieval operation binding the contract event 0x0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997.
//
// Solidity: event QuorumNumeratorUpdated(uint256 oldQuorumNumerator, uint256 newQuorumNumerator)
func (_Governance *GovernanceFilterer) FilterQuorumNumeratorUpdated(opts *bind.FilterOpts) (*GovernanceQuorumNumeratorUpdatedIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "QuorumNumeratorUpdated")
	if err != nil {
		return nil, err
	}
	return &GovernanceQuorumNumeratorUpdatedIterator{contract: _Governance.contract, event: "QuorumNumeratorUpdated", logs: logs, sub: sub}, nil
}

// WatchQuorumNumeratorUpdated is a free log subscription operation binding the contract event 0x0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997.
//
// Solidity: event QuorumNumeratorUpdated(uint256 oldQuorumNumerator, uint256 newQuorumNumerator)
func (_Governance *GovernanceFilterer) WatchQuorumNumeratorUpdated(opts *bind.WatchOpts, sink chan<- *GovernanceQuorumNumeratorUpdated) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "QuorumNumeratorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceQuorumNumeratorUpdated)
				if err := _Governance.contract.UnpackLog(event, "QuorumNumeratorUpdated", log); err != nil {
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

// ParseQuorumNumeratorUpdated is a log parse operation binding the contract event 0x0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997.
//
// Solidity: event QuorumNumeratorUpdated(uint256 oldQuorumNumerator, uint256 newQuorumNumerator)
func (_Governance *GovernanceFilterer) ParseQuorumNumeratorUpdated(log types.Log) (*GovernanceQuorumNumeratorUpdated, error) {
	event := new(GovernanceQuorumNumeratorUpdated)
	if err := _Governance.contract.UnpackLog(event, "QuorumNumeratorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceVoteCastIterator is returned from FilterVoteCast and is used to iterate over the raw logs and unpacked data for VoteCast events raised by the Governance contract.
type GovernanceVoteCastIterator struct {
	Event *GovernanceVoteCast // Event containing the contract specifics and raw log

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
func (it *GovernanceVoteCastIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceVoteCast)
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
		it.Event = new(GovernanceVoteCast)
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
func (it *GovernanceVoteCastIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceVoteCastIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceVoteCast represents a VoteCast event raised by the Governance contract.
type GovernanceVoteCast struct {
	Voter      common.Address
	ProposalId *big.Int
	Support    uint8
	Weight     *big.Int
	Reason     string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteCast is a free log retrieval operation binding the contract event 0xb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason)
func (_Governance *GovernanceFilterer) FilterVoteCast(opts *bind.FilterOpts, voter []common.Address) (*GovernanceVoteCastIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "VoteCast", voterRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceVoteCastIterator{contract: _Governance.contract, event: "VoteCast", logs: logs, sub: sub}, nil
}

// WatchVoteCast is a free log subscription operation binding the contract event 0xb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason)
func (_Governance *GovernanceFilterer) WatchVoteCast(opts *bind.WatchOpts, sink chan<- *GovernanceVoteCast, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "VoteCast", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceVoteCast)
				if err := _Governance.contract.UnpackLog(event, "VoteCast", log); err != nil {
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

// ParseVoteCast is a log parse operation binding the contract event 0xb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason)
func (_Governance *GovernanceFilterer) ParseVoteCast(log types.Log) (*GovernanceVoteCast, error) {
	event := new(GovernanceVoteCast)
	if err := _Governance.contract.UnpackLog(event, "VoteCast", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceVoteCastWithParamsIterator is returned from FilterVoteCastWithParams and is used to iterate over the raw logs and unpacked data for VoteCastWithParams events raised by the Governance contract.
type GovernanceVoteCastWithParamsIterator struct {
	Event *GovernanceVoteCastWithParams // Event containing the contract specifics and raw log

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
func (it *GovernanceVoteCastWithParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceVoteCastWithParams)
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
		it.Event = new(GovernanceVoteCastWithParams)
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
func (it *GovernanceVoteCastWithParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceVoteCastWithParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceVoteCastWithParams represents a VoteCastWithParams event raised by the Governance contract.
type GovernanceVoteCastWithParams struct {
	Voter      common.Address
	ProposalId *big.Int
	Support    uint8
	Weight     *big.Int
	Reason     string
	Params     []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteCastWithParams is a free log retrieval operation binding the contract event 0xe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb8712.
//
// Solidity: event VoteCastWithParams(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason, bytes params)
func (_Governance *GovernanceFilterer) FilterVoteCastWithParams(opts *bind.FilterOpts, voter []common.Address) (*GovernanceVoteCastWithParamsIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "VoteCastWithParams", voterRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceVoteCastWithParamsIterator{contract: _Governance.contract, event: "VoteCastWithParams", logs: logs, sub: sub}, nil
}

// WatchVoteCastWithParams is a free log subscription operation binding the contract event 0xe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb8712.
//
// Solidity: event VoteCastWithParams(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason, bytes params)
func (_Governance *GovernanceFilterer) WatchVoteCastWithParams(opts *bind.WatchOpts, sink chan<- *GovernanceVoteCastWithParams, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "VoteCastWithParams", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceVoteCastWithParams)
				if err := _Governance.contract.UnpackLog(event, "VoteCastWithParams", log); err != nil {
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

// ParseVoteCastWithParams is a log parse operation binding the contract event 0xe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb8712.
//
// Solidity: event VoteCastWithParams(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason, bytes params)
func (_Governance *GovernanceFilterer) ParseVoteCastWithParams(log types.Log) (*GovernanceVoteCastWithParams, error) {
	event := new(GovernanceVoteCastWithParams)
	if err := _Governance.contract.UnpackLog(event, "VoteCastWithParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
