package main

import (
	"fmt"
	"bytes"
	"encoding/base64"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/light"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/pkg/errors"

	"github.com/incognitochain/bridge-eth/rpccaller"

)

type Receipt struct {
	Result *types.Receipt `json:"result"`
}

type NormalResult struct {
	Result interface{} `json:"result"`
}

// getTransactionByHashToInterface returns the transaction as a map[string]interface{} type
func getETHTransactionByHash(
	url string,
	tx common.Hash,
) (map[string]interface{}, error) {
	rpcClient := rpccaller.NewRPCClient()
	params := []interface{}{tx.String()}
	var res NormalResult
	err := rpcClient.RPCCall(
		"",
		url,
		"",
		"eth_getTransactionByHash",
		params,
		&res,
	)
	if err != nil {
		return nil, err
	}
	if res.Result == nil {
		return nil, errors.New("eth tx by hash not found")
	}
	return res.Result.(map[string]interface{}), nil
}

func getETHBlockByHash(
	url string,
	blockHash string,
) (map[string]interface{}, error) {
	rpcClient := rpccaller.NewRPCClient()
	params := []interface{}{blockHash, false}
	var res NormalResult
	err := rpcClient.RPCCall(
		"",
		url,
		"",
		"eth_getBlockByHash",
		params,
		&res,
	)
	if err != nil {
		return nil, err
	}
	return res.Result.(map[string]interface{}), nil
}

func getETHTransactionReceipt(url string, txHash common.Hash) (*types.Receipt, error) {
	rpcClient := rpccaller.NewRPCClient()
	params := []interface{}{txHash.String()}
	var res Receipt
	err := rpcClient.RPCCall(
		"",
		url,
		"",
		"eth_getTransactionReceipt",
		params,
		&res,
	)
	if err != nil {
		return nil, err
	}
	return res.Result, nil
}

func getETHDepositProof(
	url string,
	txHash common.Hash,
) (*big.Int, string, uint, []string, error) {
	// Get tx content
	txContent, err := getETHTransactionByHash(url, txHash)
	if err != nil {
		fmt.Println("fuck fuck : ", err)
		return nil, "", 0, nil, err
	}
	blockHash := txContent["blockHash"].(string)
	if err != nil {
		return nil, "", 0, nil, err
	}
	txIndexStr, success := txContent["transactionIndex"].(string)
	if !success {
		return nil, "", 0, nil, errors.New("Cannot find transactionIndex field")
	}
	txIndex, err := strconv.ParseUint(txIndexStr[2:], 16, 64)
	if err != nil {
		return nil, "", 0, nil, err
	}

	// Get tx's block for constructing receipt trie
	blockNumString, success := txContent["blockNumber"].(string)
	if !success {
		return nil, "", 0, nil, errors.New("Cannot find blockNumber field")
	}
	blockNumber := new(big.Int)
	_, success = blockNumber.SetString(blockNumString[2:], 16)
	if !success {
		return nil, "", 0, nil, errors.New("Cannot convert blockNumber into integer")
	}
	blockHeader, err := getETHBlockByHash(url, blockHash)
	if err != nil {
		return nil, "", 0, nil, err
	}

	// Get all sibling Txs
	siblingTxs, success := blockHeader["transactions"].([]interface{})
	if !success {
		return nil, "", 0, nil, errors.New("Cannot find transactions field")
	}

	// Constructing the receipt trie (source: go-ethereum/core/types/derive_sha.go)
	keybuf := new(bytes.Buffer)
	receiptTrie := new(trie.Trie)
	for i, tx := range siblingTxs {
		siblingReceipt, err := getETHTransactionReceipt(url, common.HexToHash(tx.(string)))
		if err != nil {
			return nil, "", 0, nil, err
		}
		keybuf.Reset()
		rlp.Encode(keybuf, uint(i))
		encodedReceipt, err := rlp.EncodeToBytes(siblingReceipt)
		if err != nil {
			return nil, "", 0, nil, err
		}
		receiptTrie.Update(keybuf.Bytes(), encodedReceipt)
	}

	// Constructing the proof for the current receipt (source: go-ethereum/trie/proof.go)
	proof := light.NewNodeSet()
	keybuf.Reset()
	rlp.Encode(keybuf, uint(txIndex))
	err = receiptTrie.Prove(keybuf.Bytes(), 0, proof)
	if err != nil {
		return nil, "", 0, nil, err
	}

	nodeList := proof.NodeList()
	encNodeList := make([]string, 0)
	for _, node := range nodeList {
		str := base64.StdEncoding.EncodeToString(node)
		encNodeList = append(encNodeList, str)
	}

	return blockNumber, blockHash, uint(txIndex), encNodeList, nil
}
