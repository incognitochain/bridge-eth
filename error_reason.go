package main

import (
	"bytes"
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ec "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

func errorReason(ctx context.Context, b ethereum.ContractCaller, tx *types.Transaction, blockNum *big.Int, addr ec.Address) (string, error) {
	msg := ethereum.CallMsg{
		From:     addr,
		To:       tx.To(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Value:    tx.Value(),
		Data:     tx.Data(),
	}
	_, err := b.CallContract(ctx, msg, blockNum)
	if err != nil {
		errMessage := err.Error()
		if strings.Contains(errMessage, "execution reverted: ") {
			return strings.ReplaceAll(errMessage, "execution reverted: ", ""), nil
		}
		return "", errors.Wrap(err, "CallContract")
	}
	return "", errors.New("no error found")
}

var (
	errorSig     = []byte{0x08, 0xc3, 0x79, 0xa0} // Keccak256("Error(string)")[:4]
	abiString, _ = abi.NewType("string", "", nil)
)

func unpackError(result []byte) (string, error) {
	if !bytes.Equal(result[:4], errorSig) {
		return "<tx result not Error(string)>", errors.New("TX result not of type Error(string)")
	}
	vs, err := abi.Arguments{{Type: abiString}}.UnpackValues(result[4:])
	if err != nil {
		return "<invalid tx result>", errors.Wrap(err, "unpacking revert reason")
	}
	return vs[0].(string), nil
}
