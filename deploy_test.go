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

func TestDecodeSwapBridgeInst(t *testing.T) {
	// Get proof
	url := "https://mainnet.incognito.org/fullnode:433"
	block := 73853
	proof, err := getAndDecodeBridgeSwapProof(url, block)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("inst: %x\n", proof.Instruction)
	fmt.Printf("instType: %d\n", proof.Instruction[0])
	fmt.Printf("shard: %d\n", proof.Instruction[1])
	fmt.Printf("startHeight: %d\n", big.NewInt(0).SetBytes(proof.Instruction[2:34]))
	fmt.Printf("numVals: %d\n", big.NewInt(0).SetBytes(proof.Instruction[34:66]))
}

func TestSwapBridge(t *testing.T) {
	// Enter nonce here
	nonce := uint64(0)

	// Enter gasPrice here
	gasPrice := big.NewInt(5000000000) // 5 GWei

	// Enter block here
	block := 455353

	// Get proof
	url := "https://mainnet.incognito.org/fullnode:433"
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
	auth.GasPrice = gasPrice
	if nonce > 0 {
		auth.Nonce = big.NewInt(int64(nonce))
	}
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

func TestResendETH(t *testing.T) {
	privKey, client, err := connect()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	// Enter receiver address here
	addr := ""

	// Enter nonce here
	nonce := uint64(0)

	// Enter amount to send here
	value := big.NewInt(0.1 * params.Ether)

	// Enter gasLimit and gasPrice here
	gasLimit := uint64(30000)
	gasPrice := big.NewInt(5000000000) // 5 GWei

	txHash, err := transfer(client, privKey, addr, nonce, value, gasLimit, gasPrice)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("sent, txHash: %s\n", txHash)
}

func TestReburn(t *testing.T) {
	// Get proof
	proof, err := getAndDecodeBurnProof("59333c998a206e99621faf150f46588bbdfeb6279538266de893cc309e7cf4c5")
	if err != nil {
		t.Fatal(err)
	}

	// Connect to ETH
	privKey, client, err := connect()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	// Enter nonce here
	nonce := uint64(0)

	// Enter gasPrice here
	gasPrice := big.NewInt(5000000000) // 5 GWei

	// Get contract instance
	vaultAddr := common.HexToAddress(VaultAddress)
	c, err := vault.NewVault(vaultAddr, client)
	if err != nil {
		t.Fatal(err)
	}

	// Burn
	auth := bind.NewKeyedTransactor(privKey)
	auth.GasPrice = gasPrice
	if nonce > 0 {
		auth.Nonce = big.NewInt(int64(nonce))
	}
	tx, err := Withdraw(c, auth, proof)
	if err != nil {
		t.Fatal(err)
	}
	txHash := tx.Hash()
	fmt.Printf("burned, txHash: %x\n", txHash[:])
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

	value := big.NewInt(0.1 * params.Ether)
	gasLimit := uint64(30000)
	gasPrice := big.NewInt(20000000000)
	for i, addr := range addrs {
		txHash, err := transfer(client, privKey, addr, nonce+uint64(i), value, gasLimit, gasPrice)
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
	value *big.Int,
	gasLimit uint64,
	gasPrice *big.Int,
) (string, error) {
	toAddress := common.HexToAddress(to)
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

func TestInstructionUsed(t *testing.T) {
	proof, err := getAndDecodeBurnProof("4600db42637741fac7d8514f0ebabab431a67f7e4d2df55f3e2ce1e88e4b3fe9")
	if err != nil {
		t.Fatal(err)
	}
	instHash := crypto.Keccak256(proof.Instruction)

	// Connect to ETH
	_, client, err := connect()
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

	h := [32]byte{}
	copy(h[:], instHash)
	fmt.Println(c.IsWithdrawed(nil, h))
}

func TestBurn(t *testing.T) {
	// Get proof
	proof, err := getAndDecodeBurnProof("f9a347693a0c81168cfc12bd909455d1d53ee3d1527f2bf48d7d3615448ad862")
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
	v, err := vault.NewVault(vaultAddr, client)
	if err != nil {
		t.Fatal(err)
	}

	amount := big.NewInt(0.123 * params.Ether)
	_, err = depositDetail(
		privKey,
		v,
		amount,
		IncPaymentAddr,
		0,
		0,
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRedeposit(t *testing.T) {
	// Set up client
	privKey, client, err := connect()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	// Enter nonce here
	nonce := uint64(0)

	// Enter amount here
	amount := big.NewInt(int64(0))

	// Enter user's incognito address here
	incPaymentAddr := ""

	// Enter gasPrice here
	gasPrice := big.NewInt(5000000000) // 5 GWei

	// Get contract instance
	vaultAddr := common.HexToAddress(VaultAddress)
	v, err := vault.NewVault(vaultAddr, client)
	if err != nil {
		t.Fatal(err)
	}

	// Deposit
	_, err = depositDetail(
		privKey,
		v,
		amount,
		incPaymentAddr,
		nonce,
		0,
		gasPrice,
	)
	if err != nil {
		t.Fatal(err)
	}
}

func depositDetail(
	privKey *ecdsa.PrivateKey,
	v *vault.Vault,
	amount *big.Int,
	incPaymentAddr string,
	nonce uint64,
	gasLimit uint64,
	gasPrice *big.Int,
) (*types.Transaction, error) {
	auth := bind.NewKeyedTransactor(privKey)
	auth.Value = amount
	if gasLimit > 0 {
		auth.GasLimit = gasLimit
	}
	if gasPrice != nil {
		auth.GasPrice = gasPrice
	}
	if nonce > 0 {
		auth.Nonce = big.NewInt(int64(nonce))
	}

	// Deposit
	tx, err := v.Deposit(auth, incPaymentAddr)
	if err != nil {
		return nil, err
	}
	txHash := tx.Hash()
	fmt.Printf("deposited, txHash: %x\n", txHash[:])
	return tx, nil
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
