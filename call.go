package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/incognitochain/bridge-eth/bridge/incognito_proxy"
	"github.com/incognitochain/bridge-eth/bridge/vault"
)

func Withdraw(v *vault.Vault, auth *bind.TransactOpts, proof *decodedProof) (*types.Transaction, error) {
	return nil, nil
}

// TODO: replace with Withdraw
func WithdrawNew(
	v *vault.Vault,
	auth *bind.TransactOpts,
	inst []byte,
	heights [2]*big.Int,
	minedProofs [2]vault.VaultMinedProof,
	ancestorProofs [2]vault.VaultMerkleProof,
) (*types.Transaction, error) {
	tx, err := v.Withdraw(
		auth,
		inst,
		heights,
		minedProofs,
		ancestorProofs,
	)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func SubmitBurnProof(v *vault.Vault, auth *bind.TransactOpts, proof *decodedProof) (*types.Transaction, error) {
	return nil, nil
}

func SubmitBurnProofNew(
	v *vault.Vault,
	auth *bind.TransactOpts,
	inst []byte,
	heights [2]*big.Int,
	minedProofs [2]vault.VaultMinedProof,
	ancestorProofs [2]vault.VaultMerkleProof,
) (*types.Transaction, error) {
	tx, err := v.SubmitBurnProof(
		auth,
		inst,
		heights,
		minedProofs,
		ancestorProofs,
	)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func SubmitBridgeCandidate(
	inc *incognito_proxy.IncognitoProxy,
	auth *bind.TransactOpts,
	proof *CandidateProof,
) (*types.Transaction, error) {
	tx, err := inc.SubmitBridgeCandidate(auth, proof.Instruction, proof.InstProofs)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func SubmitBeaconCandidate(
	inc *incognito_proxy.IncognitoProxy,
	auth *bind.TransactOpts,
	proof *CandidateProof,
) (*types.Transaction, error) {
	tx, err := inc.SubmitBeaconCandidate(auth, proof.Instruction, proof.InstProofs[0])
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func SubmitFinalityProof(inc *incognito_proxy.IncognitoProxy, auth *bind.TransactOpts, proof *FinalityProof) (*types.Transaction, error) {
	tx, err := inc.SubmitFinalityProof(auth, proof.Instructions, proof.InstProofs, proof.SwapID, proof.IsBeacon)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// TODO: update for deploy_test
func SwapBridge(inc *incognito_proxy.IncognitoProxy, auth *bind.TransactOpts, proof *decodedProof) (*types.Transaction, error) {
	auth.GasPrice = big.NewInt(20000000000)
	return nil, nil
	// tx, err := inc.SwapBridgeCommittee(
	// 	auth,
	// 	proof.Instruction,

	// 	proof.InstPaths,
	// 	proof.InstPathIsLefts,
	// 	proof.InstRoots,
	// 	proof.BlkData,
	// 	proof.SigIdxs,
	// 	proof.SigVs,
	// 	proof.SigRs,
	// 	proof.SigSs,
	// )
	// if err != nil {
	// 	return nil, err
	// }
	// return tx, nil
}

func SwapBeacon(inc *incognito_proxy.IncognitoProxy, auth *bind.TransactOpts, proof *decodedProof) (*types.Transaction, error) {
	auth.GasPrice = big.NewInt(20000000000)
	return nil, nil
	// tx, err := inc.SwapBeaconCommittee(
	// 	auth,
	// 	proof.Instruction,

	// 	proof.InstPaths[0],
	// 	proof.InstPathIsLefts[0],
	// 	proof.InstRoots[0],
	// 	proof.BlkData[0],
	// 	proof.SigIdxs[0],
	// 	proof.SigVs[0],
	// 	proof.SigRs[0],
	// 	proof.SigSs[0],
	// )
	// if err != nil {
	// 	return nil, err
	// }
	// return tx, nil
}
