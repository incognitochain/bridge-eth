package main

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/incognitochain/bridge-eth/bridge/incognito_proxy"
	"github.com/incognitochain/bridge-eth/bridge/vault"
	"github.com/incognitochain/bridge-eth/erc20"
	"github.com/pkg/errors"
)

func TestERC20Burn(t *testing.T) {
	// Get proof
	txID := ""
	proof, err := getAndDecodeBurnProof(txID)
	if err != nil {
		t.Fatal(err)
	}

	// Get contract instance
	privKey, c := connectAndInstantiate(t)

	// Burn
	auth := bind.NewKeyedTransactor(privKey)
	tx, err := Withdraw(c.v, auth, proof)
	if err != nil {
		t.Fatal(err)
	}
	txHash := tx.Hash()
	fmt.Printf("burned erc20, txHash: %x\n", txHash[:])
}

func TestERC20Lock(t *testing.T) {
	// Get contract instance
	privKey, c := connectAndInstantiate(t)

	// Approve
	amount := int64(1000)
	_, err := approveERC20(privKey, c.vAddr, c.token, big.NewInt(amount))
	if err != nil {
		t.Fatal(err)
	}

	// Deposit
	if _, err := depositERC20(privKey, c.v, c.tokenAddr, big.NewInt(amount)); err != nil {
		t.Fatal(err)
	}
}

func TestERC20Deposit(t *testing.T) {
	privKey, c := connectAndInstantiate(t)

	// Deposit
	amount := big.NewInt(int64(100))
	if _, err := depositERC20(privKey, c.v, c.tokenAddr, amount); err != nil {
		t.Fatal(err)
	}
}

func TestERC20Redeposit(t *testing.T) {
	// Set up client
	privKey, client, err := connect()
	if err != nil {
		t.Fatal(err)
	}

	// Enter nonce here
	nonce := uint64(0)

	// Enter amount here
	amount := big.NewInt(int64(0))

	// Enter user's incognito address here
	incPaymentAddr := ""

	// Enter gasPrice here
	gasPrice := big.NewInt(5000000000) // 5 GWei

	// Enter token here
	tokenAddr := ""

	// Get contract instances
	c, err := instantiate(client, IncognitoProxyAddress, VaultAddress, tokenAddr)
	if err != nil {
		t.Fatal(err)
	}

	// Deposit
	_, err = depositERC20Detail(
		privKey,
		c.v,
		c.tokenAddr,
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

func TestERC20Approve(t *testing.T) {
	// Get contract instances
	privKey, c := connectAndInstantiate(t)

	// Approve
	amount := big.NewInt(int64(100))
	_, err := approveERC20(privKey, c.vAddr, c.token, amount)
	if err != nil {
		t.Fatal(err)
	}
}

func TestERC20Reapprove(t *testing.T) {
	// Set up client
	privKey, client, err := connect()
	if err != nil {
		t.Fatal(err)
	}

	// Enter nonce here
	nonce := uint64(0)

	// Enter amount here
	amount := big.NewInt(int64(0))

	// Enter gasPrice here
	gasPrice := big.NewInt(5000000000) // 5 GWei

	// Enter token here
	tokenAddr := ""

	// Get contract instances
	c, err := instantiate(client, IncognitoProxyAddress, VaultAddress, tokenAddr)
	if err != nil {
		t.Fatal(err)
	}

	// Approve
	_, err = approveERC20Detail(
		privKey,
		c.vAddr,
		c.token,
		amount,
		nonce,
		0,
		gasPrice,
	)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetAllowance(t *testing.T) {
	// Get contract instances
	privKey, c := connectAndInstantiate(t)

	// Allowance
	userAddr := crypto.PubkeyToAddress(privKey.PublicKey)
	allow, err := c.token.Allowance(nil, userAddr, c.vAddr)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("allowance: %d\n", allow)

	bal, _ := c.token.BalanceOf(nil, userAddr)
	fmt.Printf("balanceOf owner: %d\n", bal)
	bal, _ = c.token.BalanceOf(nil, c.vAddr)
	fmt.Printf("balanceOf spender: %d\n", bal)
}

func TestERC20Deploy(t *testing.T) {
	privKey, client, err := connect()
	if err != nil {
		t.Error(err)
	}
	defer client.Close()

	// Deploy incognito_proxy
	auth := bind.NewKeyedTransactor(privKey)
	name := "CHICKEN"
	symbol := "CKN"
	decimals := big.NewInt(8)
	supply := big.NewInt(int64(2e18))
	addr, _, _, err := erc20.DeployErc20(auth, client, name, symbol, decimals, supply)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("deployed erc20")
	fmt.Printf("addr: %s\n", addr.Hex())
}

func connectAndInstantiate(t *testing.T) (*ecdsa.PrivateKey, *contracts) {
	privKey, client, err := connect()
	if err != nil {
		t.Fatal(err)
	}

	// Get contract instance
	c, err := instantiate(client, IncognitoProxyAddress, VaultAddress, TokenAddress)
	if err != nil {
		t.Fatal(err)
	}
	return privKey, c
}

func depositERC20Detail(
	privKey *ecdsa.PrivateKey,
	v *vault.Vault,
	tokenAddr common.Address,
	amount *big.Int,
	incPaymentAddr string,
	nonce uint64,
	gasLimit uint64,
	gasPrice *big.Int,
) (*types.Transaction, error) {
	auth := bind.NewKeyedTransactor(privKey)
	if gasLimit > 0 {
		auth.GasLimit = gasLimit
	}
	if gasPrice != nil {
		auth.GasPrice = gasPrice
	}
	if nonce > 0 {
		auth.Nonce = big.NewInt(int64(nonce))
	}

	tx, err := v.DepositERC20(auth, tokenAddr, amount, incPaymentAddr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	txHash := tx.Hash()
	fmt.Printf("erc20 deposited, txHash: %x\n", txHash[:])
	return tx, nil
}

func depositERC20(
	privKey *ecdsa.PrivateKey,
	v *vault.Vault,
	tokenAddr common.Address,
	amount *big.Int,
) (*types.Transaction, error) {
	return depositERC20Detail(
		privKey,
		v,
		tokenAddr,
		amount,
		IncPaymentAddr,
		0,
		0,
		nil,
	)
}

func approveERC20Detail(
	privKey *ecdsa.PrivateKey,
	spender common.Address,
	token Tokener,
	amount *big.Int,
	nonce uint64,
	gasLimit uint64,
	gasPrice *big.Int,
) (*types.Transaction, error) {
	// Check balance
	userAddr := crypto.PubkeyToAddress(privKey.PublicKey)
	bal, _ := token.BalanceOf(nil, userAddr)
	fmt.Printf("erc20 balance: %d\n", bal)

	// Approve
	auth := bind.NewKeyedTransactor(privKey)
	if gasLimit > 0 {
		auth.GasLimit = gasLimit
	}
	if gasPrice != nil {
		auth.GasPrice = gasPrice
	}
	if nonce > 0 {
		auth.Nonce = big.NewInt(int64(nonce))
	}

	tx, err := token.Approve(auth, spender, amount)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	txHash := tx.Hash()
	fmt.Printf("erc20 approved, txHash: %x\n", txHash[:])
	return tx, nil
}

func approveERC20(privKey *ecdsa.PrivateKey, spender common.Address, token Tokener, amount *big.Int) (*types.Transaction, error) {
	return approveERC20Detail(
		privKey,
		spender,
		token,
		amount,
		0,
		0,
		nil,
	)
}

func instantiate(client *ethclient.Client, incAddr, vAddr, tokenAddr string) (*contracts, error) {
	// Get contract instance
	var err error
	c := &contracts{}
	c.incAddr = common.HexToAddress(incAddr)
	c.inc, err = incognito_proxy.NewIncognitoProxy(c.incAddr, client)
	if err != nil {
		return nil, err
	}

	// Vault
	c.vAddr = common.HexToAddress(vAddr)
	c.v, err = vault.NewVault(c.vAddr, client)
	if err != nil {
		return nil, err
	}

	// ERC20 token
	c.tokenAddr = common.HexToAddress(tokenAddr)
	c.token, err = erc20.NewErc20(c.tokenAddr, client)
	if err != nil {
		return nil, err
	}
	return c, nil
}
