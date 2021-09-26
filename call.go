package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/incognitochain/bridge-eth/bridge/incognito_proxy"
	"github.com/incognitochain/bridge-eth/bridge/vault"
	"github.com/incognitochain/bridge-eth/bridge/prveth"
)

func Withdraw(v *vault.Vault, auth *bind.TransactOpts, proof *decodedProof) (*types.Transaction, error) {
	// auth.GasPrice = big.NewInt(20000000000)
	tx, err := v.Withdraw(
		auth,
		proof.Instruction,
		proof.Heights[0],

		proof.InstPaths[0],
		proof.InstPathIsLefts[0],
		proof.InstRoots[0],
		proof.BlkData[0],
		proof.SigIdxs[0],
		proof.SigVs[0],
		proof.SigRs[0],
		proof.SigSs[0],
	)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func SubmitBurnProof(v *vault.Vault, auth *bind.TransactOpts, proof *decodedProof) (*types.Transaction, error) {
	// auth.GasPrice = big.NewInt(20000000000)
	tx, err := v.SubmitBurnProof(
		auth,
		proof.Instruction,
		proof.Heights[0],

		proof.InstPaths[0],
		proof.InstPathIsLefts[0],
		proof.InstRoots[0],
		proof.BlkData[0],
		proof.SigIdxs[0],
		proof.SigVs[0],
		proof.SigRs[0],
		proof.SigSs[0],
	)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func SubmitMintPRVProof(v *prveth.Prveth, auth *bind.TransactOpts, proof *decodedProof) (*types.Transaction, error) {
	// auth.GasPrice = big.NewInt(20000000000)
	tx, err := v.Mint(
		auth,
		proof.Instruction,
		proof.Heights[0],

		proof.InstPaths[0],
		proof.InstPathIsLefts[0],
		proof.InstRoots[0],
		proof.BlkData[0],
		proof.SigIdxs[0],
		proof.SigVs[0],
		proof.SigRs[0],
		proof.SigSs[0],
	)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func SwapBeacon(inc *incognito_proxy.IncognitoProxy, auth *bind.TransactOpts, proof *decodedProof) (*types.Transaction, error) {
	// auth.GasPrice = big.NewInt(20000000000)
	tx, err := inc.SwapBeaconCommittee(
		auth,
		proof.Instruction,

		proof.InstPaths[0],
		proof.InstPathIsLefts[0],
		proof.InstRoots[0],
		proof.BlkData[0],
		proof.SigIdxs[0],
		proof.SigVs[0],
		proof.SigRs[0],
		proof.SigSs[0],
	)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func SwapBridge(inc *incognito_proxy.IncognitoProxy, auth *bind.TransactOpts, proof *decodedProof) (*types.Transaction, error) {
	// auth.GasPrice = big.NewInt(20000000000)
	tx, err := inc.SwapBridgeCommittee(
		auth,
		proof.Instruction,

		proof.InstPaths,
		proof.InstPathIsLefts,
		proof.InstRoots,
		proof.BlkData,
		proof.SigIdxs,
		proof.SigVs,
		proof.SigRs,
		proof.SigSs,
	)
	if err != nil {
		return nil, err
	}
	return tx, nil
}
