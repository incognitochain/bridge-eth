package main

// Basic imports
import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"testing"

	"github.com/incognitochain/bridge-eth/bridge/kbntrade"
	"github.com/incognitochain/bridge-eth/bridge/vault"
	"github.com/incognitochain/bridge-eth/bridge/zrxtrade"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type TradingMainnetDeployTestSuite struct {
	suite.Suite
	VaultAddr         common.Address
	KyberContractAddr common.Address
	ZRXContractAddr   common.Address
	WETHAddr          common.Address

	Admin        common.Address
	IncProxyAddr common.Address
	PrevVault    common.Address

	ETHPrivKeyHex string
	ETHRelayer    string
}

func NewTradingMainnetDeployTestSuite() *TradingMainnetDeployTestSuite {
	return &TradingMainnetDeployTestSuite{}
}

func (tradingDeploySuite *TradingMainnetDeployTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")

	// tradingDeploySuite.KyberContractAddr = common.HexToAddress("0xF77eC7Ed5f5B9a5aee4cfa6FFCaC6A4C315BaC76")
	// tradingDeploySuite.ZRXContractAddr = common.HexToAddress("0xf1ec01d6236d3cd881a0bf0130ea25fe4234003e")
	// tradingDeploySuite.WETHAddr = common.HexToAddress("0xd0a1e359811322d97991e03f863a0c30c2cf029c")

	// tradingDeploySuite.Admin = common.HexToAddress("0x126748A0144968DD14b0187B906dE62378c59067")
	// tradingDeploySuite.IncProxyAddr = common.HexToAddress("0x6b8B83235f20875A797571f7bC0bD7dD54205213")
	// tradingDeploySuite.PrevVault = common.HexToAddress("0x0000000000000000000000000000000000000000")

	// tradingDeploySuite.ETHPrivKeyHex = "B8DB29A7A43FB88AD520F762C5FDF6F1B0155637FA1E5CB2C796AFE9E5C04E31"
	// tradingDeploySuite.ETHRelayer = "https://kovan.infura.io/v3/93fe721349134964aa71071a713c5cef"

	tradingDeploySuite.KyberContractAddr = common.HexToAddress("0x818E6FECD516Ecc3849DAf6845e3EC868087B755")
	tradingDeploySuite.ZRXContractAddr = common.HexToAddress("0x95e6f48254609a6ee006f7d493c8e5fb97094cef")
	tradingDeploySuite.WETHAddr = common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")

	tradingDeploySuite.Admin = common.HexToAddress("")
	tradingDeploySuite.IncProxyAddr = common.HexToAddress("")
	tradingDeploySuite.PrevVault = common.HexToAddress("0x0261DB5AfF8E5eC99fBc8FBBA5D4B9f8EcD44ec7")

	tradingDeploySuite.ETHPrivKeyHex = ""
	tradingDeploySuite.ETHRelayer = "https://mainnet.infura.io/v3/93fe721349134964aa71071a713c5cef"
}

func (tradingDeploySuite *TradingMainnetDeployTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
}

func TestTradingMainnetDeployTestSuite(t *testing.T) {
	fmt.Println("Starting entry point...")
	tradingDeploySuite := NewTradingMainnetDeployTestSuite()
	suite.Run(t, tradingDeploySuite)
	fmt.Println("Finishing entry point...")
}

func connectToETH(ethHost, privKeyHex string) (*ecdsa.PrivateKey, *ethclient.Client, error) {
	privKey, err := crypto.HexToECDSA(privKeyHex)
	if err != nil {
		return nil, nil, err
	}

	fmt.Printf("Sign Txs with address: %s\n", crypto.PubkeyToAddress(privKey.PublicKey).Hex())

	network := "mainnet"
	fmt.Printf("Connecting to network %s\n", network)
	client, err := ethclient.Dial(ethHost)
	if err != nil {
		return nil, nil, err
	}

	return privKey, client, nil
}

func (tradingDeploySuite *TradingMainnetDeployTestSuite) TestDeployAllMainnetContracts() {
	fmt.Println("Deploying all mainnet contracts for trading...")

	ethPrivKey, ethClient, err := connectToETH(tradingDeploySuite.ETHRelayer, tradingDeploySuite.ETHPrivKeyHex)
	require.Equal(tradingDeploySuite.T(), nil, err)

	auth := bind.NewKeyedTransactor(ethPrivKey)
	auth.Value = big.NewInt(0)
	// auth.GasPrice = big.NewInt(10000000000)
	// auth.GasLimit = 4000000

	// Deploy vault
	vaultAddr, tx, _, err := vault.DeployVault(
		auth,
		ethClient,
		tradingDeploySuite.Admin,
		tradingDeploySuite.IncProxyAddr,
		tradingDeploySuite.PrevVault,
	)
	require.Equal(tradingDeploySuite.T(), nil, err)

	// Wait until tx is confirmed
	err = wait(ethClient, tx.Hash())
	require.Equal(tradingDeploySuite.T(), nil, err)
	fmt.Println("deployed vault")
	fmt.Printf("addr: %s\n", vaultAddr.Hex())

	// Deploy kbntrade
	kbnTradeAddr, tx, _, err := kbntrade.DeployKBNTrade(auth, ethClient, tradingDeploySuite.KyberContractAddr)
	require.Equal(tradingDeploySuite.T(), nil, err)

	// Wait until tx is confirmed
	err = wait(ethClient, tx.Hash())
	require.Equal(tradingDeploySuite.T(), nil, err)
	fmt.Println("deployed kbntrade")
	fmt.Printf("addr: %s\n", kbnTradeAddr.Hex())

	// Deploy 0xTrade
	zrxTradeAddr, tx, _, err := zrxtrade.DeployZrxtrade(auth, ethClient, tradingDeploySuite.WETHAddr, tradingDeploySuite.ZRXContractAddr, vaultAddr)
	require.Equal(tradingDeploySuite.T(), nil, err)

	// Wait until tx is confirmed
	err = wait(ethClient, tx.Hash())
	require.Equal(tradingDeploySuite.T(), nil, err)
	fmt.Println("deployed zrxTrade")
	fmt.Printf("addr: %s\n", zrxTradeAddr.Hex())
}
