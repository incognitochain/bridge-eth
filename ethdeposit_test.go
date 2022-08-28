package bridge

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
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
	// depositingEther := float64(0.05)

	// txHash := ethDepositSuite.depositETH(
	// 	depositingEther,
	// 	ethDepositSuite.IncPaymentAddrStr,
	// 	ethDepositSuite.VaultAddr,
	// 	ethDepositSuite.ETHClient,
	// )

	txHash := common.HexToHash("0x1b4fa59cd2c026993f22375b533b43e6dbb167901eddfa7df8d15d281a3a118f")
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(ethDepositSuite.ETHHost, txHash)
	// time.Sleep(200 * time.Second)
	require.Equal(ethDepositSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	_, err = ethDepositSuite.callIssuingETHReq(
		ethDepositSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingethreq",
	)
	// // expected an error is returned due to not meet 15 block confirmations
	// require.NotEqual(ethDepositSuite.T(), nil, err)

	// fmt.Println("Waiting 90s for 15 block confirmations")
	// time.Sleep(90 * time.Second)

	// // retry again
	// _, err = ethDepositSuite.callIssuingETHReq(
	// 	ethDepositSuite.IncEtherTokenIDStr,
	// 	ethDepositProof,
	// 	ethBlockHash,
	// 	ethTxIdx,
	// 	"createandsendtxwithissuingethreq",
	// )
	// require.Equal(ethDepositSuite.T(), nil, err)
}

// func (ethDepositSuite *ETHDepositTestSuite) TestDepositERC20() {
// 	fmt.Println("Running deposit erc20 test...")
// 	// depositingDAI, _ := big.NewInt(0).SetString("10000000000000000", 10) // 0.01 DAI

// 	// txHash := ethDepositSuite.depositERC20ToBridge(
// 	// 	depositingDAI,
// 	// 	common.HexToAddress("0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984"), // iuni
// 	// 	ethDepositSuite.IncPaymentAddrStr,
// 	// 	ethDepositSuite.VaultAddr,
// 	// 	ethDepositSuite.ETHClient,
// 	// 	ethDepositSuite.ChainIDETH,
// 	// )

// 	// fmt.Println("Waiting 200s for 15 block confirmations")
// 	// time.Sleep(200 * time.Second)

// 	txHash := common.HexToHash("0xd66dc09d35c23beb890726346824ae82c3008778290c24c2bbcf6653ff39e68a")

// 	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(ethDepositSuite.ETHHost, txHash)
// 	require.Equal(ethDepositSuite.T(), nil, err)
// 	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

// 	_, err = ethDepositSuite.callIssuingETHReq(
// 		"779af92f61938356cc7e5f80eb4bc207aa39129e469ed8d9fd945c062a11582b",
// 		ethDepositProof,
// 		ethBlockHash,
// 		ethTxIdx,
// 		"createandsendtxwithissuingethreq",
// 	)
// 	// expected an error is returned due to not meet 15 block confirmations
// 	require.NotEqual(ethDepositSuite.T(), nil, err)
// }

// func (ethDepositSuite *ETHDepositTestSuite) TestWithdrawEther() {
// 	withdrawAmt, _ := big.NewInt(0).SetString("10000000", 10) // 0.01 ETH
// 	burningRes, err := ethDepositSuite.callBurningPToken(
// 		ethDepositSuite.IncEtherTokenIDStr,
// 		withdrawAmt,
// 		ethDepositSuite.ETHOwnerAddrStr,
// 		"createandsendburningrequest",
// 	)
// 	require.Equal(ethDepositSuite.T(), nil, err)
// 	burningTxID, found := burningRes["TxID"]
// 	require.Equal(ethDepositSuite.T(), true, found)
// 	time.Sleep(120 * time.Second)

// 	ethDepositSuite.submitBurnProofForWithdrawal(
// 		burningTxID.(string),
// 		"getburnproof",
// 		ethDepositSuite.VaultAddr,
// 		ethDepositSuite.ETHClient,
// 		ethDepositSuite.ChainIDETH,
// 	)
// }

func (ethDepositSuite *ETHDepositTestSuite) TestWithdrawERC20() {
	withdrawAmt, _ := big.NewInt(0).SetString("5000000", 10) // 0.005 UNI
	burningRes, err := ethDepositSuite.callBurningPToken(
		"779af92f61938356cc7e5f80eb4bc207aa39129e469ed8d9fd945c062a11582b",
		withdrawAmt,
		ethDepositSuite.ETHOwnerAddrStr,
		"createandsendburningrequest",
	)
	require.Equal(ethDepositSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(ethDepositSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	ethDepositSuite.submitBurnProofForWithdrawal(
		burningTxID.(string),
		"getburnproof",
		ethDepositSuite.VaultAddr,
		ethDepositSuite.ETHClient,
		ethDepositSuite.ChainIDETH,
	)
}
