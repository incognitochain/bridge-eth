package bridge

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/incognitochain/bridge-eth/bridge/incognito_proxy"
	"github.com/incognitochain/bridge-eth/bridge/vault"
	"github.com/incognitochain/bridge-eth/bridge/vaultproxy"
	"github.com/incognitochain/bridge-eth/rpccaller"
	"github.com/pkg/errors"
)

type EstimateRes struct {
	rpccaller.RPCBaseRes
	Result interface{} `json:"Result"`
}

func TestGetNonceOfPendingTx(t *testing.T) {
	_, client, _, err := connect()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	tx := "0xcb086f05903f80f206cc3c3a1c9f47ccf9d40bd64ba85b065f8419456d0b8617"
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	a, _, err := client.TransactionByHash(ctx, common.HexToHash(tx))
	if err != nil {
		t.Fatalf("failed getting nonce: %+v", err)
	}
	fmt.Println(big.NewInt(int64(a.Nonce())))
}

func TestGetTxStatus(t *testing.T) {
	_, client, _, err := connect()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	tx := "0x635431d7a220fb21c14e152f79663b74bc6c97eda3b2f821ac7c1cdb1c60c3c8"
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	receipt, err := client.TransactionReceipt(ctx, common.HexToHash(tx))
	if err != nil {
		t.Fatalf("failed getting status: %+v", err)
	}
	fmt.Println(receipt.Status)
}

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

func TestSwapBeacon(t *testing.T) {
	// Get proof
	proof := getFixedSwapBeaconProof()

	// Connect to ETH
	privKey, client, chainID, err := connect()
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
	auth, err := bind.NewKeyedTransactorWithChainID(privKey, big.NewInt(chainID))
	if err != nil {
		t.Fatal(err)
	}
	tx, err := SwapBeacon(c, auth, proof)
	if err != nil {
		t.Fatal(err)
	}
	txHash := tx.Hash()
	fmt.Printf("swapped, txHash: %x\n", txHash[:])
}

func TestResendETH(t *testing.T) {
	privKey, client, _, err := connect()
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
	privKey, client, chainID, err := connect()
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
	auth, err := bind.NewKeyedTransactorWithChainID(privKey, big.NewInt(chainID))
	if err != nil {
		t.Fatal(err)
	}
	auth.GasPrice = gasPrice
	if nonce > 0 {
		auth.Nonce = big.NewInt(int64(nonce))
	}

	vaultAbi, err := abi.JSON(strings.NewReader(vault.VaultABI))
	if err != nil {
		t.Fatal(err)
	}
	input, err := vaultAbi.Pack(
		"withdraw",
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
		t.Fatal(err)
	}

	gasLimit, err := EstimateGas(ethereum.CallMsg{From: auth.From, To: &vaultAddr, GasPrice: auth.GasPrice, Value: auth.Value, Data: input})
	if err != nil {
		t.Fatal(err)
	}
	auth.GasLimit = gasLimit
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

	privKey, client, _, err := connect()
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
	proof, err := getAndDecodeBurnProof("40db51b1811fcf4d6b2220e83ec8b4743f0d56558da933791218f8a9dfe22e6f")
	if err != nil {
		t.Fatal(err)
	}
	instHash := crypto.Keccak256(proof.Instruction)

	// Connect to ETH
	_, client, _, err := connect()
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
	privKey, client, chainID, err := connect()
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
	auth, err := bind.NewKeyedTransactorWithChainID(privKey, big.NewInt(chainID))
	if err != nil {
		t.Fatal(err)
	}
	tx, err := Withdraw(c, auth, proof)
	if err != nil {
		t.Fatal(err)
	}
	txHash := tx.Hash()
	fmt.Printf("burned, txHash: %x\n", txHash[:])
}

func TestDeposit(t *testing.T) {
	privKey, client, chainID, err := connect()
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
		chainID,
	)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRedeposit(t *testing.T) {
	// Set up client
	privKey, client, chainID, err := connect()
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
		chainID,
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
	chainID int64,
) (*types.Transaction, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(privKey, big.NewInt(chainID))
	if err != nil {
		return nil, err
	}

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
	txId := [32]byte{}
	sign, err := SignDataToShield(txId, genesisAcc2.PrivateKey, genesisAcc2.Address)
	if err != nil {
		return nil, err
	}
	tx, err := v.Deposit(auth, incPaymentAddr, txId, sign)
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
	privKey, client, chainID, err := connect()
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
	auth, err := bind.NewKeyedTransactorWithChainID(privKey, big.NewInt(chainID))
	if err != nil {
		t.Fatal(err)
	}

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
	vaultAddr, _, _, err := vault.DeployVault(auth, client)
	if err != nil {
		t.Fatal(err)
	}

	vaultAbi, _ := abi.JSON(strings.NewReader(vault.VaultABI))
	input, _ := vaultAbi.Pack("initialize", prevVault)
	vaultProxyAddr, _, _, err := vaultproxy.DeployTransparentUpgradeableProxy(auth, client, vaultAddr, admin, incAddr, input)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("deployed vault")
	fmt.Printf("addr: %s\n", vaultProxyAddr.Hex())
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

func connect() (*ecdsa.PrivateKey, *ethclient.Client, int64, error) {
	privKeyHex := os.Getenv("PRIVKEY")
	privKey, err := crypto.HexToECDSA(privKeyHex)
	if err != nil {
		return nil, nil, 0, err
	}
	fmt.Printf("Sign Txs with address: %s\n", crypto.PubkeyToAddress(privKey.PublicKey).Hex())

	ethNodeEndpoint := os.Getenv("ETHNODE")
	if len(ethNodeEndpoint) == 0 {
		return nil, nil, 0, errors.New("os env value of ETHNODE not set yet")
	}
	client, err := ethclient.Dial(ethNodeEndpoint)
	if err != nil {
		return nil, nil, 0, err
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, nil, 0, err
	}

	return privKey, client, chainID.Int64(), nil
}

func connect2(privateKeyHex, node string) (*ecdsa.PrivateKey, *ethclient.Client, int64, error) {
	privKeyHex := privateKeyHex
	privKey, err := crypto.HexToECDSA(privKeyHex)
	if err != nil {
		return nil, nil, 0, err
	}
	fmt.Printf("Sign Txs with address: %s\n", crypto.PubkeyToAddress(privKey.PublicKey).Hex())

	ethNodeEndpoint := node
	if len(ethNodeEndpoint) == 0 {
		return nil, nil, 0, errors.New("os env value of ETHNODE not set yet")
	}
	client, err := ethclient.Dial(ethNodeEndpoint)
	if err != nil {
		return nil, nil, 0, err
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, nil, 0, err
	}

	return privKey, client, chainID.Int64(), nil
}

func ToCallArg(msg ethereum.CallMsg) interface{} {
	arg := map[string]interface{}{
		"from": msg.From,
		"to":   msg.To,
	}
	if len(msg.Data) > 0 {
		arg["data"] = hexutil.Bytes(msg.Data)
	}
	if msg.Value != nil {
		arg["value"] = (*hexutil.Big)(msg.Value)
	}
	if msg.Gas != 0 {
		arg["gas"] = hexutil.Uint64(msg.Gas)
	}
	if msg.GasPrice != nil {
		arg["gasPrice"] = (*hexutil.Big)(msg.GasPrice)
	}
	return arg
}

func EstimateGas(msg ethereum.CallMsg) (uint64, error) {
	var resp EstimateRes
	rpcClient := rpccaller.NewRPCClient()
	params := []interface{}{
		ToCallArg(msg),
		"latest",
	}
	err := rpcClient.RPCCall(
		"",
		os.Getenv("ETHNODE"),
		"",
		"eth_estimateGas",
		params,
		&resp,
	)
	if err != nil {
		return 0, err
	}

	bb, _ := json.Marshal(resp)
	fmt.Println("Estimate gas res: ", string(bb))
	if resp.RPCError != nil {
		return 0, errors.New(resp.RPCError.Message)
	}
	hexString := resp.Result.(string)
	n := new(big.Int)
	n.SetString(hexString, 16)
	if n.Cmp(big.NewInt(0)) != 0 {
		return 0, errors.New("Call estimate gas got error")
	}
	return n.Uint64(), nil
}
