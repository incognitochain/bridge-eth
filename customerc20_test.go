package bridge

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/incognitochain/bridge-eth/erc20/dai"
	"github.com/incognitochain/bridge-eth/erc20/usdt"
)

func TestUSDTDeploy(t *testing.T) {
	privKey, client, chainID, err := connect()
	if err != nil {
		t.Error(err)
	}
	defer client.Close()

	// Deploy USDT
	auth, err := bind.NewKeyedTransactorWithChainID(privKey, big.NewInt(chainID))
	if err != nil {
		t.Fatal(err)
	}
	bal, _ := big.NewInt(1).SetString("100000000000", 10)
	addr, _, _, err := usdt.DeployUsdt(auth, client, bal, "Tether USD", "USDT", big.NewInt(6))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("deployed usdt")
	fmt.Printf("addr: %s\n", addr.Hex())
}

func TestDAIDeploy(t *testing.T) {
	privKey, client, chainID, err := connect()
	if err != nil {
		t.Error(err)
	}
	defer client.Close()

	// Deploy DAI
	auth, err := bind.NewKeyedTransactorWithChainID(privKey, big.NewInt(chainID))
	if err != nil {
		t.Fatal(err)
	}
	symbol := [32]byte{'D', 'A', 'I'}
	addr, tx, d, err := dai.DeployDai(auth, client, symbol)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("deployed dai")
	fmt.Printf("addr: %s\n", addr.Hex())

	// Wait till tx is done
	err = wait(client, tx.Hash())
	if err != nil {
		t.Fatal(err)
	}

	// Mint DAI
	bal, _ := big.NewInt(1).SetString("1000000000000000000000000000", 10)
	tx, err = d.Mint(auth, bal)

	// daiABI, _ := abi.JSON(strings.NewReader(dai.DaiABI))
	// fmt.Printf("%+v\n", daiABI.Methods)
	// fmt.Printf("%+v\n", daiABI.Methods["mint"])
	// contract := bind.NewBoundContract(address, daiABI, client, client, client)
	// dtr := dai.DaiTransactor{contract: contract}
	// tx, err = dtr.Mint(auth, crypto.PubkeyToAddress(privKey.PublicKey), bal)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("minted dai, tx: %s\n", tx.Hash().Hex())
}
