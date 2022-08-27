package bridge

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"

	"github.com/stretchr/testify/require"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type AuroraTestSuite struct {
	*TradingTestSuite

	// token amounts for tests
	DepositingEther       float64
	AVAXBalanceAfterStep1 *big.Int
	USDTBalanceAfterStep2 *big.Int

	USDTTokenAddress string
	VaultAURORAAddr  common.Address
}

func NewAuroraTestSuite(tradingTestSuite *TradingTestSuite) *AuroraTestSuite {
	return &AuroraTestSuite{
		TradingTestSuite: tradingTestSuite,
	}
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (tradingSuite *AuroraTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")
	// Kovan env
	tradingSuite.IncUSDTTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000071"
	tradingSuite.IncEtherTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000073"
	tradingSuite.DepositingEther = float64(0.0001)
	tradingSuite.VaultAURORAAddr = common.HexToAddress("0x14D0cf3bC307aA15DA40Aa4c8cc2A2a81eF96B3a")
	tradingSuite.USDTTokenAddress = "0x30fb06E97a6CD370BCE994A88C428F9F3aB6Ec28"
}

func (tradingSuite *AuroraTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
	tradingSuite.AURORAClient.Close()
}

func (tradingSuite *AuroraTestSuite) SetupTest() {
	fmt.Println("Setting up the test...")
}

func (tradingSuite *AuroraTestSuite) TearDownTest() {
	fmt.Println("Tearing down the test...")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestAuroraTestSuite(t *testing.T) {
	fmt.Println("Starting entry point for aurora test suite...")

	tradingSuite := new(TradingTestSuite)
	suite.Run(t, tradingSuite)

	auroraSuite := NewAuroraTestSuite(tradingSuite)
	suite.Run(t, auroraSuite)

	fmt.Println("Finishing entry point for 0x test suite...")
}

func (tradingSuite *AuroraTestSuite) Test1Aurora() {
	fmt.Println("============ TEST SHIELD UNSHIELD NATIVE AURORA ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
	burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("------------ STEP 1: porting ETH to pETH --------------")
	txHash := tradingSuite.depositETHCompliance(
		tradingSuite.DepositingEther,
		tradingSuite.IncPaymentAddrStr,
		tradingSuite.VaultAURORAAddr,
		tradingSuite.AURORAClient,
	)

	fmt.Println("Waiting 50s for 5 blocks confirmation")
	time.Sleep(50 * time.Second)
	_, err := tradingSuite.callIssuingAURORAReq(
		tradingSuite.IncEtherTokenIDStr,
		txHash.String()[2:],
		"createandsendtxwithissuingaurorareq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(150 * time.Second)

	fmt.Println("------------ STEP 2: burning pETH to deposit ETH to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningPToken(
		tradingSuite.IncEtherTokenIDStr,
		burningPETH,
		pubKeyToAddrStr[2:],
		"createandsendburningaurorafordeposittoscrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForDepositToSC(
		burningTxID.(string),
		big.NewInt(int64(tradingSuite.ChainIDAURORA)),
		"getburnauroraprooffordeposittosc",
		tradingSuite.VaultAURORAAddr,
		tradingSuite.AURORAClient,
	)
	deposited := tradingSuite.getDepositedBalanceWithParams(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		pubKeyToAddrStr,
		tradingSuite.VaultAURORAAddr,
		tradingSuite.AURORAClient,
	)
	fmt.Println("deposited AURORA: ", deposited)

	fmt.Println("------------ step 3: withdrawing ETH from SC to pETH on Incognito --------------")
	txHash = tradingSuite.requestWithdrawCompliance(
		tradingSuite.EtherAddressStr,
		deposited,
		tradingSuite.AURORAClient,
		big.NewInt(int64(tradingSuite.ChainIDAURORA)),
		tradingSuite.VaultAURORAAddr,
		REQ_WITHDRAW_PREFIX,
	)

	fmt.Println("Waiting 50s for 5 blocks confirmation")
	time.Sleep(50 * time.Second)
	_, err = tradingSuite.callIssuingAURORAReq(
		tradingSuite.IncEtherTokenIDStr,
		txHash.String()[2:],
		"createandsendtxwithissuingaurorareq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(150 * time.Second)

	fmt.Println("------------ step 4: withdrawing pAVX from Incognito to AVAX --------------")
	withdrawingPAVAX := big.NewInt(0).Div(deposited, big.NewInt(1000000000))
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncEtherTokenIDStr,
		withdrawingPAVAX,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningaurorarequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForWithdrawal(
		burningTxID.(string),
		"getauroraburnproof",
		tradingSuite.VaultAURORAAddr,
		tradingSuite.AURORAClient,
		tradingSuite.ChainIDAURORA,
	)

	bal := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		tradingSuite.AURORAClient,
	)
	tradingSuite.AVAXBalanceAfterStep1 = bal
	fmt.Println("AVAX balance after step 1: ", tradingSuite.AVAXBalanceAfterStep1)
}

func (tradingSuite *AuroraTestSuite) Test2AuroraToken() {
	fmt.Println("============ TEST SHIELD UNSHIELD TOKEN AURORA ===========")
	fmt.Println("------------ step 0: declaration & initialization --------------")
	depositingUSDT := big.NewInt(3e6)
	burningPUSDT := depositingUSDT

	daibal := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.USDTTokenAddress),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		tradingSuite.AURORAClient,
	)
	fmt.Println("usdt balance of owner: ", daibal)

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("pubKeyToAddrStr: ", pubKeyToAddrStr)

	fmt.Println("------------ step 1: porting USDT to pUSDT --------------")
	txHash := tradingSuite.depositERC20ComplianceToBridge(
		depositingUSDT,
		common.HexToAddress(tradingSuite.USDTTokenAddress),
		tradingSuite.IncPaymentAddrStr,
		tradingSuite.VaultAURORAAddr,
		tradingSuite.AURORAClient,
		tradingSuite.ChainIDAURORA,
	)

	fmt.Println("Waiting 50s for 5 blocks confirmation")
	time.Sleep(50 * time.Second)

	_, err := tradingSuite.callIssuingAURORAReq(
		tradingSuite.IncUSDTTokenIDStr,
		txHash.String()[2:],
		"createandsendtxwithissuingaurorareq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(150 * time.Second)

	fmt.Println("------------ step 2: burning pUSDT to deposit USDT to SC --------------")

	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningPToken(
		tradingSuite.IncUSDTTokenIDStr,
		burningPUSDT,
		pubKeyToAddrStr[2:],
		"createandsendburningaurorafordeposittoscrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForDepositToSC(
		burningTxID.(string),
		big.NewInt(int64(tradingSuite.ChainIDAURORA)),
		"getburnauroraprooffordeposittosc",
		tradingSuite.VaultAURORAAddr,
		tradingSuite.AURORAClient,
	)
	deposited := tradingSuite.getDepositedBalanceWithParams(
		common.HexToAddress("tradingSuite.USDTTokenAddress"),
		pubKeyToAddrStr,
		tradingSuite.VaultAURORAAddr,
		tradingSuite.AURORAClient,
	)
	fmt.Println("deposited usdt: ", deposited)

	fmt.Println("------------ step 3: withdrawing USDT from SC to pUSDT on Incognito --------------")
	txHash = tradingSuite.requestWithdrawCompliance(
		tradingSuite.USDTTokenAddress,
		deposited,
		tradingSuite.AURORAClient,
		big.NewInt(int64(tradingSuite.ChainIDAURORA)),
		tradingSuite.VaultAURORAAddr,
		REQ_WITHDRAW_PREFIX,
	)

	fmt.Println("Waiting 50s for 5 blocks confirmation")
	time.Sleep(50 * time.Second)

	_, err = tradingSuite.callIssuingAURORAReq(
		tradingSuite.IncUSDTTokenIDStr,
		txHash.String()[2:],
		"createandsendtxwithissuingaurorareq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(150 * time.Second)

	fmt.Println("------------ step 4: withdrawing pUSDT from Incognito to USDT --------------")
	withdrawingPMRK := deposited
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncUSDTTokenIDStr,
		withdrawingPMRK,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningaurorarequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForWithdrawal(
		burningTxID.(string),
		"getauroraburnproof",
		tradingSuite.VaultAURORAAddr,
		tradingSuite.AURORAClient,
		tradingSuite.ChainIDAURORA,
	)

	bal := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.USDTTokenAddress),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		tradingSuite.AURORAClient,
	)
	tradingSuite.USDTBalanceAfterStep2 = bal
	fmt.Println("USDT balance after step 2: ", tradingSuite.USDTBalanceAfterStep2)
}
