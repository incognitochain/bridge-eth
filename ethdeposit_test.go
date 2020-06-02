package main

import (
	"fmt"
	"time"
	"testing"
	"math/big"

	"github.com/stretchr/testify/suite"
	"github.com/stretchr/testify/require"
	"github.com/ethereum/go-ethereum/common"
)

type ETHDepositTestSuite struct {
	*TradingTestSuite
}

func NewETHDepositTestSuite(tradingTestSuite *TradingTestSuite) *ETHDepositTestSuite {
	return &ETHDepositTestSuite{
		TradingTestSuite: tradingTestSuite,
	}
}

func (ethDepositSuite *ETHDepositTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")
}

func (ethDepositSuite *ETHDepositTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
	ethDepositSuite.ETHClient.Close()
}

func TestETHDepositTestSuite(t *testing.T) {
	fmt.Println("Starting entry point...")
	tradingSuite := new(TradingTestSuite)
	suite.Run(t, tradingSuite)

	ethDepositSuite := NewETHDepositTestSuite(tradingSuite)
	suite.Run(t, ethDepositSuite)
	fmt.Println("Finishing entry point...")
}

func (ethDepositSuite *ETHDepositTestSuite) TestDepositEther() {
	fmt.Println("Running deposit ether test...")
	depositingEther := float64(0.05)

	txHash := ethDepositSuite.depositETH(
		depositingEther,
		ethDepositSuite.IncPaymentAddrStr,
	)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(ethDepositSuite.ETHHost, txHash)
	require.Equal(ethDepositSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	_, err = ethDepositSuite.callIssuingETHReq(
		ethDepositSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	// expected an error is returned due to not meet 15 block confirmations
	require.NotEqual(ethDepositSuite.T(), nil, err)

	fmt.Println("Waiting 90s for 15 block confirmations")
	time.Sleep(90 * time.Second)

	// retry again
	_, err = ethDepositSuite.callIssuingETHReq(
		ethDepositSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(ethDepositSuite.T(), nil, err)
}

func (ethDepositSuite *ETHDepositTestSuite) TestDepositERC20() {
	fmt.Println("Running deposit erc20 test...")
	depositingDAI, _ := big.NewInt(0).SetString("1000000000000000000", 10) // 1 DAI

	txHash := ethDepositSuite.depositERC20ToBridge(
		depositingDAI,
		common.HexToAddress(ethDepositSuite.DAIAddressStr),
		ethDepositSuite.IncPaymentAddrStr,
	)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(ethDepositSuite.ETHHost, txHash)
	require.Equal(ethDepositSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	_, err = ethDepositSuite.callIssuingETHReq(
		ethDepositSuite.IncDAITokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	// expected an error is returned due to not meet 15 block confirmations
	require.NotEqual(ethDepositSuite.T(), nil, err)

	fmt.Println("Waiting 90s for 15 block confirmations")
	time.Sleep(90 * time.Second)

	// retry again
	_, err = ethDepositSuite.callIssuingETHReq(
		ethDepositSuite.IncDAITokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(ethDepositSuite.T(), nil, err)
}

