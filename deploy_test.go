package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/incognitochain/bridge-eth/bridge/incognito_proxy"
	"github.com/incognitochain/bridge-eth/bridge/vault"
	"github.com/pkg/errors"
)

func TestSwapBridge(t *testing.T) {
	// Get proof
	url := "http://54.39.158.106:19032"
	block := 54
	proof, err := getAndDecodeBridgeSwapProof(url, block)
	if err != nil {
		t.Fatal(err)
	}

	// Connect to ETH
	privKey, client, err := connect()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	// Get contract instance
	incAddr := common.HexToAddress(IncognitoProxyAddress)
	c, err := incognito_proxy.NewIncognitoProxy(incAddr, client)
	if err != nil {
		t.Fatal(err)
	}

	// Swap
	auth := bind.NewKeyedTransactor(privKey)
	tx, err := SwapBridge(c, auth, proof)
	if err != nil {
		t.Fatal(err)
	}
	txHash := tx.Hash()
	fmt.Printf("swapped, txHash: %x\n", txHash[:])
}

func TestSwapBeacon(t *testing.T) {
	// Get proof
	proof := getFixedSwapBeaconProof()

	// Connect to ETH
	privKey, client, err := connect()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	// Get contract instance
	incAddr := common.HexToAddress(IncognitoProxyAddress)
	c, err := incognito_proxy.NewIncognitoProxy(incAddr, client)
	if err != nil {
		t.Fatal(err)
	}

	// Swap
	auth := bind.NewKeyedTransactor(privKey)
	tx, err := SwapBeacon(c, auth, proof)
	if err != nil {
		t.Fatal(err)
	}
	txHash := tx.Hash()
	fmt.Printf("swapped, txHash: %x\n", txHash[:])
}

func TestMassSend(t *testing.T) {
	addrs := []string{
		"0xDF1A9BE4dA9Ed6CDC28bea3c23E81B8a3e4480Ae",
		"0x354e2c1ee8f254f379A17679Dd14e3afa61c0346",
		"0x9a6A22438307C68A794485b86Faa6b72Aa67Ded7",
		"0x7A279AEe9cc310B64F0F159904271c0a68014082",
	}

	privKey, client, err := connect()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	// Deposit
	nonce, err := client.NonceAt(context.Background(), crypto.PubkeyToAddress(privKey.PublicKey), nil)
	if err != nil {
		t.Fatal(err)
	}

	for i, addr := range addrs {
		txHash, err := transfer(client, privKey, addr, nonce+uint64(i))
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("sent, txHash: %s\n", txHash)
	}
}

func transfer(
	client *ethclient.Client,
	privKey *ecdsa.PrivateKey,
	to string,
	nonce uint64,
) (string, error) {
	toAddress := common.HexToAddress(to)
	value := big.NewInt(0.1 * params.Ether)
	gasLimit := uint64(30000)
	gasPrice := big.NewInt(20000000000)
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", errors.WithStack(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privKey)
	if err != nil {
		return "", errors.WithStack(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return signedTx.Hash().String(), nil
}

func TestBurn(t *testing.T) {
	// Get proof
	proof, err := getAndDecodeBurnProof("59333c998a206e99621faf150f46588bbdfeb6279538266de893cc309e7cf4c5")
	if err != nil {
		t.Fatal(err)
	}
	// return

	// Connect to ETH
	privKey, client, err := connect()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	// Get contract instance
	vaultAddr := common.HexToAddress(VaultAddress)
	c, err := vault.NewVault(vaultAddr, client)
	if err != nil {
		t.Fatal(err)
	}

	// Burn
	auth := bind.NewKeyedTransactor(privKey)
	tx, err := Withdraw(c, auth, proof)
	if err != nil {
		t.Fatal(err)
	}
	txHash := tx.Hash()
	fmt.Printf("burned, txHash: %x\n", txHash[:])
}

func TestDeposit(t *testing.T) {
	privKey, client, err := connect()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	// Get contract instance
	vaultAddr := common.HexToAddress(VaultAddress)
	c, err := vault.NewVault(vaultAddr, client)
	if err != nil {
		t.Fatal(err)
	}

	// Deposit
	auth := bind.NewKeyedTransactor(privKey)
	// auth.GasPrice = big.NewInt(20000000000)
	auth.Value = big.NewInt(0.123 * params.Ether)
	tx, err := c.Deposit(auth, IncPaymentAddr)
	if err != nil {
		t.Fatal(err)
	}
	txHash := tx.Hash()
	fmt.Printf("deposited, txHash: %x\n", txHash[:])
}

// func TestGetCommittee(t *testing.T) {
// 	_, c := connectAndInstantiate(t)
// 	beaconBlk, _ := c.inc.LatestBeaconBlk(nil)
// 	for {
// 		for i := 0; i < comm_size; i++ {
// 			pubkeys, err := c.inc.BeaconCommPubkeys(nil, beaconBlk, big.NewInt(int64(i)))
// 			if err != nil {
// 				t.Fatal(err)
// 			}
// 			fmt.Printf("beaconBlk: %d %x\n", beaconBlk, pubkeys)
// 		}

// 		prevBeaconBlk, err := c.inc.BeaconCommPrevBlk(nil, beaconBlk)
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		if beaconBlk.Uint64() == 0 {
// 			break
// 		}
// 		beaconBlk = prevBeaconBlk
// 	}
// 	bridgeBlk, _ := c.inc.LatestBridgeBlk(nil)
// 	for {
// 		for i := 0; i < comm_size; i++ {
// 			pubkeys, err := c.inc.BridgeCommPubkeys(nil, beaconBlk, big.NewInt(int64(i)))
// 			if err != nil {
// 				t.Fatal(err)
// 			}
// 			fmt.Printf("bridgeBlk: %d %x\n", bridgeBlk, pubkeys)
// 		}

// 		prevBridgeBlk, err := c.inc.BridgeCommPrevBlk(nil, bridgeBlk)
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		if bridgeBlk.Uint64() == 0 {
// 			break
// 		}
// 		bridgeBlk = prevBridgeBlk
// 	}
// }

func TestDeployProxyAndVault(t *testing.T) {
	privKey, client, err := connect()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	admin := common.HexToAddress(Admin)
	fmt.Println("Admin address:", admin.Hex())

	// Genesis committee
	cmtee := getCommitteeHardcoded()
	// beaconComm, bridgeComm, err := getCommittee("http://54.39.158.106:19032/")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// Deploy incognito_proxy
	auth := bind.NewKeyedTransactor(privKey)
	auth.Value = big.NewInt(0)
	// auth.GasPrice = big.NewInt(10000000000)
	// auth.GasLimit = 4000000
	incAddr, tx, _, err := incognito_proxy.DeployIncognitoProxy(auth, client, admin, cmtee.beacons, cmtee.bridges)
	if err != nil {
		t.Fatal(err)
	}
	// incAddr := common.HexToAddress(IncognitoProxyAddress)
	fmt.Println("deployed incognito_proxy")
	fmt.Printf("addr: %s\n", incAddr.Hex())

	// Wait until tx is confirmed
	if err := wait(client, tx.Hash()); err != nil {
		t.Fatal(err)
	}

	// Deploy vault
	prevVault := common.Address{}
	vaultAddr, _, _, err := vault.DeployVault(auth, client, admin, incAddr, prevVault)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("deployed vault")
	fmt.Printf("addr: %s\n", vaultAddr.Hex())
}

func wait(client *ethclient.Client, tx common.Hash) error {
	ctx := context.Background()
	for range time.Tick(10 * time.Second) {
		_, err := client.TransactionReceipt(ctx, tx)
		if err == nil {
			break
		} else if err == ethereum.NotFound {
			continue
		} else {
			return err
		}
	}
	return nil
}

func connect() (*ecdsa.PrivateKey, *ethclient.Client, error) {
	privKeyHex := os.Getenv("PRIVKEY")
	privKey, err := crypto.HexToECDSA(privKeyHex)
	if err != nil {
		return nil, nil, err
	}
	fmt.Printf("Sign Txs with address: %s\n", crypto.PubkeyToAddress(privKey.PublicKey).Hex())

	network := "mainnet"
	fmt.Printf("Connecting to network %s\n", network)
	client, err := ethclient.Dial(fmt.Sprintf("https://%s.infura.io/v3/29fead42346b4bfa88dd5fd7e56b6406", network))
	if err != nil {
		return nil, nil, err
	}

	return privKey, client, nil
}
