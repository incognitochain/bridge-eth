package main

import (
	"crypto/ecdsa"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/incognitochain/bridge-eth/bridge/kbntrade"
	"github.com/incognitochain/bridge-eth/bridge/vault"
	"github.com/pkg/errors"
)

func TestRetireVaultAdmin(t *testing.T) {
	privKey, c := getVault(t)

	successor := common.HexToAddress(Successor)
	fmt.Println("Successor address:", successor.Hex())

	auth := bind.NewKeyedTransactor(privKey)
	_, err := c.Retire(auth, successor)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClaimVaultAdmin(t *testing.T) {
	privKey, c := getVault(t)
	auth := bind.NewKeyedTransactor(privKey)
	_, err := c.Claim(auth)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeployNewVaultToMigrate(t *testing.T) {
	privKey, client, err := connect()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	admin := common.HexToAddress(Admin)
	incAddr := common.HexToAddress(IncognitoProxyAddress)
	prevVaultAddr := common.HexToAddress(VaultAddress)
	fmt.Println("Admin address:", admin.Hex())
	fmt.Println("IncognitoProxy address:", incAddr.Hex())
	fmt.Println("PrevVault address:", prevVaultAddr.Hex())

	// Deploy vault
	auth := bind.NewKeyedTransactor(privKey)
	vaultAddr, _, _, err := vault.DeployVault(auth, client, admin, incAddr, prevVaultAddr)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("deployed vault")
	fmt.Printf("addr: %s\n", vaultAddr.Hex())
}

func TestDeployKBNTrade(t *testing.T) {
	privKey, client, err := connect()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	kbnProxy := common.HexToAddress(KBNProxy)
	vaultAddr := common.HexToAddress(VaultAddress)
	fmt.Println("KBNProxy address:", kbnProxy.Hex())
	fmt.Println("Vault address:", vaultAddr.Hex())

	// Deploy KBNTrade
	auth := bind.NewKeyedTransactor(privKey)
	kbnTradeAddr, _, _, err := kbntrade.DeployKbntrade(auth, client, kbnProxy, vaultAddr)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("deployed kbnTrade")
	fmt.Printf("addr: %s\n", kbnTradeAddr.Hex())
}

func TestPauseVault(t *testing.T) {
	privKey, c := getVault(t)

	// Pause vault
	auth := bind.NewKeyedTransactor(privKey)
	_, err := c.Pause(auth)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnpauseVault(t *testing.T) {
	privKey, c := getVault(t)

	// Pause vault
	auth := bind.NewKeyedTransactor(privKey)
	_, err := c.Unpause(auth)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMigrateVault(t *testing.T) {
	privKey, c := getVault(t)

	// Migrate vault
	newAddr := ""
	if len(newAddr) != 42 {
		t.Fatal(errors.New("invalid new vault's address"))
	}
	newVault := common.HexToAddress(newAddr)
	auth := bind.NewKeyedTransactor(privKey)
	_, err := c.Migrate(auth, newVault)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMoveAssetsVault(t *testing.T) {
	privKey, c := getVault(t)

	// Migrate vault
	assets := []common.Address{
		common.Address{}, // ETH
	}
	auth := bind.NewKeyedTransactor(privKey)
	_, err := c.MoveAssets(auth, assets)
	if err != nil {
		t.Fatal(err)
	}
}

func getVault(t *testing.T) (*ecdsa.PrivateKey, *vault.Vault) {
	privKey, client, err := connect()
	if err != nil {
		t.Fatal(err)
	}

	// Get vault instance
	vaultAddr := common.HexToAddress(VaultAddress)
	c, err := vault.NewVault(vaultAddr, client)
	if err != nil {
		t.Fatal(err)
	}
	return privKey, c
}
