package main

import (
	"fmt"
	"math"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type PrvEvmTokenTestSuite struct {
	*TradingTestSuite

	// token amounts for tests
	DepositingPRV float64
	BurnPRV       float64
	PRVDecimal    int

	ETHChainID int64
	BSCChainID int64
}

func NewPrvEvmTokenTestSuite(tradingTestSuite *TradingTestSuite) *PrvEvmTokenTestSuite {
	return &PrvEvmTokenTestSuite{
		TradingTestSuite: tradingTestSuite,
	}
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (tradingSuite *PrvEvmTokenTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")
	// Kovan env
	tradingSuite.DepositingPRV = float64(100)
	tradingSuite.BurnPRV = float64(50)
	tradingSuite.PRVDecimal = 9
	tradingSuite.ETHChainID = 42
	tradingSuite.BSCChainID = 97
}

func (tradingSuite *PrvEvmTokenTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
	tradingSuite.ETHClient.Close()
}

func (tradingSuite *PrvEvmTokenTestSuite) SetupTest() {
	fmt.Println("Setting up the test...")
}

func (tradingSuite *PrvEvmTokenTestSuite) TearDownTest() {
	fmt.Println("Tearing down the test...")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestPrvEvmTokenTestSuite(t *testing.T) {
	fmt.Println("Starting entry point for prv evm test suite...")

	tradingSuite := new(TradingTestSuite)
	suite.Run(t, tradingSuite)

	prvEvmTokenTestSuite := NewPrvEvmTokenTestSuite(tradingSuite)
	suite.Run(t, prvEvmTokenTestSuite)

	fmt.Println("Finishing entry point for 0x test suite...")
}

func (tradingSuite *PrvEvmTokenTestSuite) Test1BurnAndMintPRVERC20() {
	fmt.Println("============ TEST TRADE ETHER FOR DAI WITH uniswap AGGREGATOR ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	decimal := tradingSuite.PRVDecimal
	burnPRV := big.NewInt(int64(tradingSuite.BurnPRV * math.Pow(10, float64(decimal))))

	fmt.Println("------------ STEP 1: burning PRV to mint PRV to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningPRV(
		burnPRV,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningprverc20request",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(60 * time.Second)

	tradingSuite.submitBurnProofForMintPRV(burningTxID.(string), tradingSuite.PRVERC20Addr, "getprverc20burnproof", tradingSuite.ETHClient, tradingSuite.ETHChainID)
	time.Sleep(30 * time.Second)

	// pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("------------ STEP 2: burn PRV from SC to mint PRV on incognito --------------")
	txHash := tradingSuite.burnPRV(
		tradingSuite.BurnPRV,
		tradingSuite.IncPaymentAddrStr,
		tradingSuite.PRVERC20Addr,
		tradingSuite.ETHClient,
		tradingSuite.ETHChainID,
	)
	// txHash := common.HexToHash("0xbb1d4535fa5a2ad294bca45e5c5b72dfa5c0feb6f05361a1ef7f8cb521d4193a")
	time.Sleep(30 * time.Second)
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(150 * time.Second)
	_, err = tradingSuite.callIssuingPRVReq(
		tradingSuite.IncPRVTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingprverc20req",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)
}

func (tradingSuite *PrvEvmTokenTestSuite) Test2BurnAndMintPRVBEP20() {
	fmt.Println("============ TEST TRADE ETHER FOR DAI WITH uniswap AGGREGATOR ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	decimal := tradingSuite.PRVDecimal
	burnPRV := big.NewInt(int64(tradingSuite.BurnPRV * math.Pow(10, float64(decimal))))

	fmt.Println("------------ STEP 1: burning PRV to mint PRV to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningPRV(
		burnPRV,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningprvbep20request",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(60 * time.Second)

	tradingSuite.submitBurnProofForMintPRV(burningTxID.(string), tradingSuite.PRVBEP20Addr, "getprvbep20burnproof", tradingSuite.BSCClient, tradingSuite.BSCChainID)
	time.Sleep(30 * time.Second)

	// pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("------------ STEP 2: burn PRV from SC to mint PRV on incognito --------------")
	txHash := tradingSuite.burnPRV(
		tradingSuite.BurnPRV,
		tradingSuite.IncPaymentAddrStr,
		tradingSuite.PRVBEP20Addr,
		tradingSuite.BSCClient,
		tradingSuite.BSCChainID,
	)
	// txHash := common.HexToHash("0xbb1d4535fa5a2ad294bca45e5c5b72dfa5c0feb6f05361a1ef7f8cb521d4193a")
	time.Sleep(30 * time.Second)
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(60 * time.Second)
	_, err = tradingSuite.callIssuingPRVReq(
		tradingSuite.IncPRVTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingprvbep20req",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)
}

func (tradingSuite *PrvEvmTokenTestSuite) Test3BurnAndMintPDEXERC20() {
	fmt.Println("============ BURN AND MINT PDEX TOKEN ON ETH ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	decimal := tradingSuite.PRVDecimal
	burnPRV := big.NewInt(int64(tradingSuite.BurnPRV * math.Pow(10, float64(decimal))))

	fmt.Println("------------ STEP 1: burning PRV to mint PRV to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningPToken(
		tradingSuite.IncPDEXTokenIDStr,
		burnPRV,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningpdexerc20request",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(60 * time.Second)

	tradingSuite.submitBurnProofForMintPRV(burningTxID.(string), tradingSuite.PDEXERC20Addr, "getpdexerc20burnproof", tradingSuite.ETHClient, 42)
	time.Sleep(30 * time.Second)

	// pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	// fmt.Println("------------ STEP 2: burn PRV from SC to mint PRV on incognito --------------")
	// txHash := tradingSuite.burnPRV(
	// 	tradingSuite.BurnPRV,
	// 	tradingSuite.IncPaymentAddrStr,
	// 	tradingSuite.PRVBEP20Addr,
	// )
	// // txHash := common.HexToHash("0xbb1d4535fa5a2ad294bca45e5c5b72dfa5c0feb6f05361a1ef7f8cb521d4193a")
	// time.Sleep(30 * time.Second)
	// _, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	// require.Equal(tradingSuite.T(), nil, err)
	// fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	// fmt.Println("Waiting 90s for 15 blocks confirmation")
	// time.Sleep(60 * time.Second)
	// _, err = tradingSuite.callIssuingPRVReq(
	// 	tradingSuite.IncPRVTokenIDStr,
	// 	ethDepositProof,
	// 	ethBlockHash,
	// 	ethTxIdx,
	// 	"createandsendtxwithissuingprvbep20req",
	// )
	// require.Equal(tradingSuite.T(), nil, err)
	// time.Sleep(120 * time.Second)
}

func (tradingSuite *PrvEvmTokenTestSuite) Test4BurnAndMintPDEXBEP20() {
	fmt.Println("============ BURN AND MINT PDEX TOKEN ON BSC ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	decimal := tradingSuite.PRVDecimal
	burnPRV := big.NewInt(int64(tradingSuite.BurnPRV * math.Pow(10, float64(decimal))))

	fmt.Println("------------ STEP 1: burning PRV to mint PRV to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningPToken(
		tradingSuite.IncPDEXTokenIDStr,
		burnPRV,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningpdexbep20request",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(60 * time.Second)

	tradingSuite.submitBurnProofForMintPRV(burningTxID.(string), tradingSuite.PDEXBEP20Addr, "getpdexbep20burnproof", tradingSuite.BSCClient, 97)
	time.Sleep(30 * time.Second)

	// pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	// fmt.Println("------------ STEP 2: burn PRV from SC to mint PRV on incognito --------------")
	// txHash := tradingSuite.burnPRV(
	// 	tradingSuite.BurnPRV,
	// 	tradingSuite.IncPaymentAddrStr,
	// 	tradingSuite.PRVBEP20Addr,
	// )
	// // txHash := common.HexToHash("0xbb1d4535fa5a2ad294bca45e5c5b72dfa5c0feb6f05361a1ef7f8cb521d4193a")
	// time.Sleep(30 * time.Second)
	// _, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	// require.Equal(tradingSuite.T(), nil, err)
	// fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	// fmt.Println("Waiting 90s for 15 blocks confirmation")
	// time.Sleep(60 * time.Second)
	// _, err = tradingSuite.callIssuingPRVReq(
	// 	tradingSuite.IncPRVTokenIDStr,
	// 	ethDepositProof,
	// 	ethBlockHash,
	// 	ethTxIdx,
	// 	"createandsendtxwithissuingprvbep20req",
	// )
	// require.Equal(tradingSuite.T(), nil, err)
	// time.Sleep(120 * time.Second)
}
