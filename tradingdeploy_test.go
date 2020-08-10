package main

// Basic imports
import (
	"fmt"
	"math/big"

	"testing"

	"github.com/incognitochain/bridge-eth/bridge/incognito_proxy"
	kbnmultiTrade "github.com/incognitochain/bridge-eth/bridge/kbnmultitrade"
	"github.com/incognitochain/bridge-eth/bridge/kbntrade"
	"github.com/incognitochain/bridge-eth/bridge/locker"
	"github.com/incognitochain/bridge-eth/bridge/vault"
	"github.com/incognitochain/bridge-eth/bridge/zrxtrade"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type TradingDeployTestSuite struct {
	*TradingTestSuite

	KyberContractAddr common.Address
	ZRXContractAddr   common.Address
	WETHAddr          common.Address
}

func NewTradingDeployTestSuite(tradingTestSuite *TradingTestSuite) *TradingDeployTestSuite {
	return &TradingDeployTestSuite{
		TradingTestSuite: tradingTestSuite,
	}
}

func (tradingDeploySuite *TradingDeployTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")
	// 0x kovan env
	tradingDeploySuite.KyberContractAddr = common.HexToAddress("0x692f391bCc85cefCe8C237C01e1f636BbD70EA4D")
	tradingDeploySuite.ZRXContractAddr = common.HexToAddress("0xf1ec01d6236d3cd881a0bf0130ea25fe4234003e")
	tradingDeploySuite.WETHAddr = common.HexToAddress("0xd0a1e359811322d97991e03f863a0c30c2cf029c")
}

func (tradingDeploySuite *TradingDeployTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
	tradingDeploySuite.ETHClient.Close()
}

func TestTradingDeployTestSuite(t *testing.T) {
	fmt.Println("Starting entry point...")
	tradingSuite := new(TradingTestSuite)
	suite.Run(t, tradingSuite)

	tradingDeploySuite := NewTradingDeployTestSuite(tradingSuite)
	suite.Run(t, tradingDeploySuite)
	fmt.Println("Finishing entry point...")
}

func (tradingDeploySuite *TradingDeployTestSuite) TestDeployAllContracts() {
	admin := common.HexToAddress(Admin)
	fmt.Println("Admin address:", admin.Hex())

	// Genesis committee
	// for testnet & local env
	beaconComm, _, err := convertCommittees(
		testnetBeaconCommitteePubKeys, testnetBridgeCommitteePubKeys,
	)
	// NOTE: uncomment this block to get mainnet committees when deploying to mainnet env
	/*
		beaconComm, bridgeComm, err := convertCommittees(
			mainnetBeaconCommitteePubKeys, mainnetBridgeCommitteePubKeys,
		)
	*/

	require.Equal(tradingDeploySuite.T(), nil, err)

	// Deploy incognito_proxy
	auth := bind.NewKeyedTransactor(tradingDeploySuite.ETHPrivKey)
	auth.Value = big.NewInt(0)
	// auth.GasPrice = big.NewInt(10000000000)
	// auth.GasLimit = 4000000
	incAddr, tx, _, err := incognito_proxy.DeployIncognitoProxy(auth, tradingDeploySuite.ETHClient, admin, beaconComm)
	require.Equal(tradingDeploySuite.T(), nil, err)

	// incAddr := common.HexToAddress(IncognitoProxyAddress)
	fmt.Println("deployed incognito_proxy")
	fmt.Printf("addr: %s\n", incAddr.Hex())

	// Wait until tx is confirmed
	err = wait(tradingDeploySuite.ETHClient, tx.Hash())
	require.Equal(tradingDeploySuite.T(), nil, err)

	lockerAddr, _, _, err := locker.DeployLocker(auth, tradingDeploySuite.ETHClient, admin, common.Address{})
	require.Equal(tradingDeploySuite.T(), nil, err)
	fmt.Printf("Locker address: %s\n", lockerAddr.Hex())

	// Wait until tx is confirmed
	err = wait(tradingDeploySuite.ETHClient, tx.Hash())
	require.Equal(tradingDeploySuite.T(), nil, err)

	// Deploy vault
	prevVault := common.Address{}
	vaultAddr, tx, _, err := vault.DeployVault(auth, tradingDeploySuite.ETHClient, admin, incAddr, prevVault, lockerAddr)
	require.Equal(tradingDeploySuite.T(), nil, err)
	fmt.Println("deployed vault")
	fmt.Printf("addr: %s\n", vaultAddr.Hex())

	// Wait until tx is confirmed
	err = wait(tradingDeploySuite.ETHClient, tx.Hash())
	require.Equal(tradingDeploySuite.T(), nil, err)

	// Deploy kbntrade
	kbnTradeAddr, tx, _, err := kbntrade.DeployKBNTrade(auth, tradingDeploySuite.ETHClient, tradingDeploySuite.KyberContractAddr, vaultAddr)
	require.Equal(tradingDeploySuite.T(), nil, err)
	fmt.Println("deployed kbntrade")
	fmt.Printf("addr: %s\n", kbnTradeAddr.Hex())

	// Wait until tx is confirmed
	err = wait(tradingDeploySuite.ETHClient, tx.Hash())
	require.Equal(tradingDeploySuite.T(), nil, err)

	// Deploy kbnMultitrade
	kbnMultiTradeAddr, tx, _, err := kbnmultiTrade.DeployKBNMultiTrade(auth, tradingDeploySuite.ETHClient, tradingDeploySuite.KyberContractAddr, vaultAddr)
	require.Equal(tradingDeploySuite.T(), nil, err)
	fmt.Println("deployed kbnMultitrade")
	fmt.Printf("addr: %s\n", kbnMultiTradeAddr.Hex())

	// Wait until tx is confirmed
	err = wait(tradingDeploySuite.ETHClient, tx.Hash())
	require.Equal(tradingDeploySuite.T(), nil, err)

	// Deploy 0xTrade
	zrxTradeAddr, tx, _, err := zrxtrade.DeployZrxtrade(auth, tradingDeploySuite.ETHClient, tradingDeploySuite.WETHAddr, tradingDeploySuite.ZRXContractAddr, vaultAddr)
	require.Equal(tradingDeploySuite.T(), nil, err)

	// Wait until tx is confirmed
	err = wait(tradingDeploySuite.ETHClient, tx.Hash())
	require.Equal(tradingDeploySuite.T(), nil, err)

	fmt.Println("deployed zrxTrade")
	fmt.Printf("addr: %s\n", zrxTradeAddr.Hex())
}

func convertCommittees(
	beaconComms []string, brigeComms []string,
) ([]common.Address, []common.Address, error) {
	beaconOld := make([]common.Address, len(beaconComms))
	for i, pk := range beaconComms {
		cpk := &CommitteePublicKey{}
		cpk.FromString(pk)
		addr, err := convertPubkeyToAddress(*cpk)
		if err != nil {
			return nil, nil, err
		}
		beaconOld[i] = addr
		fmt.Printf("beaconOld: %s\n", addr.Hex())
	}

	bridgeOld := make([]common.Address, len(brigeComms))
	for i, pk := range brigeComms {
		cpk := &CommitteePublicKey{}
		cpk.FromString(pk)
		addr, err := convertPubkeyToAddress(*cpk)
		if err != nil {
			return nil, nil, err
		}
		bridgeOld[i] = addr
		fmt.Printf("bridgeOld: %s\n", addr.Hex())
	}
	return beaconOld, bridgeOld, nil
}
