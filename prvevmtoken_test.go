package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type PrvEvmTokenTestSuite struct {
	*TradingTestSuite

	// token amounts for tests
	DepositingPRV      float64
	BurnPRV	           float64
	PRVDecimal         int
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
	tradingSuite.DepositingPRV = float64(0.05)
	tradingSuite.BurnPRV = float64(10)
	tradingSuite.PRVDecimal = 9
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

func (tradingSuite *PrvEvmTokenTestSuite) Test1BurnAndMintPRV() {
	fmt.Println("============ TEST TRADE ETHER FOR DAI WITH uniswap AGGREGATOR ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	// decimal := tradingSuite.PRVDecimal
	// burnPRV := big.NewInt(int64(tradingSuite.BurnPRV * math.Pow(10, float64(decimal))))

	fmt.Println("------------ STEP 1: burning PRV to mint PRV to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	// burningRes, err := tradingSuite.callBurningPRV(
	// 	burnPRV,
	// 	tradingSuite.ETHOwnerAddrStr,
	// 	"createandsendburningprverc20request",
	// )
	// require.Equal(tradingSuite.T(), nil, err)
	// burningTxID, found := burningRes["TxID"]
	// require.Equal(tradingSuite.T(), true, found)
	// time.Sleep(30 * time.Second)

	tradingSuite.submitBurnProofForMintPRV("85985883ebc51adbe3b20a9080c9ff0f676a66bed95667a922b718e88e59eda5")
	// deposited := tradingSuite.getDepositedBalance(
	// 	tradingSuite.EtherAddressStr,
	// 	pubKeyToAddrStr,
	// )
	// fmt.Println("deposited EHT: ", deposited)
	// // require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPETH, big.NewInt(1000000000)), deposited)

	// fmt.Println("------------ step 2: withdrawing pDAI from Incognito to DAI --------------")
	// withdrawingPDAI := big.NewInt(0).Div(daiTraded, big.NewInt(1000000000))
	// burningRes, err = tradingSuite.callBurningPToken(
	// 	tradingSuite.IncDAITokenIDStr,
	// 	withdrawingPDAI,
	// 	tradingSuite.ETHOwnerAddrStr,
	// 	"createandsendburningrequest",
	// )
	// require.Equal(tradingSuite.T(), nil, err)
	// burningTxID, found = burningRes["TxID"]
	// require.Equal(tradingSuite.T(), true, found)
	// time.Sleep(120 * time.Second)

	// pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("------------ STEP 2: porting ETH to pETH --------------")
	// txHash := tradingSuite.depositETH(
	// 	tradingSuite.DepositingEther,
	// 	tradingSuite.IncPaymentAddrStr,
	// )
    // txHash := common.HexToHash("0xbb1d4535fa5a2ad294bca45e5c5b72dfa5c0feb6f05361a1ef7f8cb521d4193a")
	// // time.Sleep(15 * time.Second)
	// _, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	// require.Equal(tradingSuite.T(), nil, err)
	// fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	// fmt.Println("Waiting 90s for 15 blocks confirmation")
	// // time.Sleep(230 * time.Second)
	// _, err = tradingSuite.callIssuingETHReq(
	// 	tradingSuite.IncEtherTokenIDStr,
	// 	ethDepositProof,
	// 	ethBlockHash,
	// 	ethTxIdx,
	// )
	// require.Equal(tradingSuite.T(), nil, err)
	// time.Sleep(120 * time.Second)

	// tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))
}
