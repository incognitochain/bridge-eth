package bridge

// Basic imports
import (
	"fmt"
	"math/big"
	"os"
	"strings"
	"testing"

	"github.com/incognitochain/bridge-eth/bridge/vault"
	"github.com/incognitochain/bridge-eth/bridge/vaultbsc"
	"github.com/incognitochain/bridge-eth/bridge/vaultftm"
	"github.com/incognitochain/bridge-eth/bridge/vaultplg"
	"github.com/incognitochain/bridge-eth/bridge/vaultproxy"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type TradingComplianceDeployTestSuite struct {
	*TradingTestSuite
}

func NewTradingComplianceDeployTestSuite(tradingTestSuite *TradingTestSuite) *TradingComplianceDeployTestSuite {
	return &TradingComplianceDeployTestSuite{
		TradingTestSuite: tradingTestSuite,
	}
}

func (tradingDeploySuite *TradingComplianceDeployTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")
}

func (tradingDeploySuite *TradingComplianceDeployTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
	tradingDeploySuite.ETHClient.Close()
}

func TestComplianceDeployTestSuite(t *testing.T) {
	fmt.Println("Starting entry point...")
	tradingSuite := new(TradingTestSuite)
	suite.Run(t, tradingSuite)

	tradingDeploySuite := NewTradingComplianceDeployTestSuite(tradingSuite)
	suite.Run(t, tradingDeploySuite)
	fmt.Println("Finishing entry point...")
}

func (tradingDeploySuite *TradingComplianceDeployTestSuite) TestDeployAllContracts() {
	regulator := common.HexToAddress(Regulator)
	fmt.Println("Regulator address:", regulator.Hex())

	// Deploy incognito_proxy
	auth, err := bind.NewKeyedTransactorWithChainID(tradingDeploySuite.ETHPrivKey, big.NewInt(int64(tradingDeploySuite.ChainIDETH)))
	require.Equal(tradingDeploySuite.T(), nil, err)
	auth.Value = big.NewInt(0)
	auth.GasPrice = big.NewInt(1000000000)

	/*
		network:
		0 - deploy all
		1 - ETH
		2 - BSC
		3 - POLYGON
		4 - FANTOM
	*/
	network := os.Getenv("NETWORK")
	complianceVaultAbi, _ := abi.JSON(strings.NewReader(vault.ComplianceVaultMetaData.ABI))
	input, err := complianceVaultAbi.Pack("setRegulator", regulator)

	if network == "1" || network == "0" {
		fmt.Println("============== DEPLOY INCOGNITO CONTRACT ON ETH ===============")

		// Deploy vault
		vaultAddr, tx, _, err := vault.DeployComplianceVault(auth, tradingDeploySuite.ETHClient)
		require.Equal(tradingDeploySuite.T(), nil, err)
		fmt.Println("deployed vault")
		fmt.Printf("addr: %s\n", vaultAddr.Hex())

		// Wait until tx is confirmed
		err = wait(tradingDeploySuite.ETHClient, tx.Hash())
		require.Equal(tradingDeploySuite.T(), nil, err)

		// Deploy vault proxy
		vaultProxy, err := vaultproxy.NewTransparentUpgradeableProxy(tradingDeploySuite.VaultAddr, tradingDeploySuite.ETHClient)
		require.Equal(tradingDeploySuite.T(), nil, err)
		tx, err = vaultProxy.UpgradeToAndCall(auth, vaultAddr, input)
		require.Equal(tradingDeploySuite.T(), nil, err)

		err = wait(tradingDeploySuite.ETHClient, tx.Hash())
		require.Equal(tradingDeploySuite.T(), nil, err)

		tradingDeploySuite.ETHClient.Close()
		if network != "0" {
			return
		}
	}

	if network == "2" || network == "0" {
		fmt.Println("============== DEPLOY INCOGNITO CONTRACT ON BSC ===============")

		auth, err = bind.NewKeyedTransactorWithChainID(tradingDeploySuite.ETHPrivKey, big.NewInt(int64(tradingDeploySuite.ChainIDBSC)))
		require.Equal(tradingDeploySuite.T(), nil, err)
		auth.GasPrice = big.NewInt(10000000000)

		// Deploy vault
		vaultAddrBSC, tx, _, err := vaultbsc.DeployVaultbsc(auth, tradingDeploySuite.BSCClient)
		require.Equal(tradingDeploySuite.T(), nil, err)
		fmt.Println("deployed vault")
		fmt.Printf("addr: %s\n", vaultAddrBSC.Hex())

		// Wait until tx is confirmed
		err = wait(tradingDeploySuite.BSCClient, tx.Hash())
		require.Equal(tradingDeploySuite.T(), nil, err)

		// Deploy vault proxy
		vaultProxy, err := vaultproxy.NewTransparentUpgradeableProxy(tradingDeploySuite.VaultBSCAddr, tradingDeploySuite.BSCClient)
		require.Equal(tradingDeploySuite.T(), nil, err)
		tx, err = vaultProxy.UpgradeToAndCall(auth, vaultAddrBSC, input)
		require.Equal(tradingDeploySuite.T(), nil, err)

		err = wait(tradingDeploySuite.BSCClient, tx.Hash())
		require.Equal(tradingDeploySuite.T(), nil, err)
		tradingDeploySuite.BSCClient.Close()
		if network != "0" {
			return
		}
	}

	if network == "3" || network == "0" {
		fmt.Println("============== DEPLOY INCOGNITO CONTRACT ON POLYGON ===============")

		auth, err = bind.NewKeyedTransactorWithChainID(tradingDeploySuite.ETHPrivKey, big.NewInt(int64(tradingDeploySuite.ChainIDPLG)))
		require.Equal(tradingDeploySuite.T(), nil, err)
		auth.GasPrice = big.NewInt(10000000000)

		// Deploy vault
		vaultAddrPLG, tx, _, err := vaultplg.DeployVaultplg(auth, tradingDeploySuite.PLGClient)
		require.Equal(tradingDeploySuite.T(), nil, err)
		fmt.Println("deployed vault")
		fmt.Printf("addr: %s\n", vaultAddrPLG.Hex())

		// Wait until tx is confirmed
		err = wait(tradingDeploySuite.PLGClient, tx.Hash())
		require.Equal(tradingDeploySuite.T(), nil, err)

		// Deploy vault proxy
		vaultProxy, err := vaultproxy.NewTransparentUpgradeableProxy(tradingDeploySuite.VaultPLGAddr, tradingDeploySuite.PLGClient)
		require.Equal(tradingDeploySuite.T(), nil, err)
		tx, err = vaultProxy.UpgradeToAndCall(auth, vaultAddrPLG, input)
		require.Equal(tradingDeploySuite.T(), nil, err)

		err = wait(tradingDeploySuite.PLGClient, tx.Hash())
		require.Equal(tradingDeploySuite.T(), nil, err)
		tradingDeploySuite.PLGClient.Close()
		if network != "0" {
			return
		}
	}

	if network == "4" || network == "0" {
		fmt.Println("============== DEPLOY INCOGNITO CONTRACT ON FANTOM ===============")

		auth, err = bind.NewKeyedTransactorWithChainID(tradingDeploySuite.ETHPrivKey, big.NewInt(int64(tradingDeploySuite.ChainIDFTM)))
		require.Equal(tradingDeploySuite.T(), nil, err)
		auth.GasLimit = uint64(5000000)

		// Deploy vault
		vaultAddrFTM, tx, _, err := vaultftm.DeployVaultftm(auth, tradingDeploySuite.FTMClient)
		require.Equal(tradingDeploySuite.T(), nil, err)
		fmt.Println("deployed vault")
		fmt.Printf("addr: %s\n", vaultAddrFTM.Hex())

		// Wait until tx is confirmed
		err = wait(tradingDeploySuite.FTMClient, tx.Hash())
		require.Equal(tradingDeploySuite.T(), nil, err)

		// Deploy vault proxy
		vaultProxy, err := vaultproxy.NewTransparentUpgradeableProxy(tradingDeploySuite.VaultFTMAddr, tradingDeploySuite.FTMClient)
		require.Equal(tradingDeploySuite.T(), nil, err)
		tx, err = vaultProxy.UpgradeToAndCall(auth, vaultAddrFTM, input)
		require.Equal(tradingDeploySuite.T(), nil, err)

		err = wait(tradingDeploySuite.FTMClient, tx.Hash())
		require.Equal(tradingDeploySuite.T(), nil, err)
		tradingDeploySuite.FTMClient.Close()
		if network != "0" {
			return
		}
	}

	fmt.Println("============== NETWORK ONLY IN RANGE 0 - 6 ===============")
}
