package bridge

import (
	"fmt"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type AvaxTestSuite struct {
	*TradingTestSuite

	// token amounts for tests
	DepositingEther       float64
	AVAXBalanceAfterStep1 *big.Int
	USDTBalanceAfterStep2 *big.Int

	USDTTokenAddress string
	VaultAVAXAddr    common.Address
}

func NewAvaxTestSuite(tradingTestSuite *TradingTestSuite) *AvaxTestSuite {
	return &AvaxTestSuite{
		TradingTestSuite: tradingTestSuite,
	}
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (tradingSuite *AvaxTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")
	// Kovan env
	tradingSuite.IncUSDTTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000062"
	tradingSuite.DepositingEther = float64(0.0001)
	tradingSuite.VaultAVAXAddr = common.HexToAddress("0x30fb06E97a6CD370BCE994A88C428F9F3aB6Ec28")
	tradingSuite.USDTTokenAddress = "0x48601da98F729aE3B4d7DeD8F29777DF102a167d"
}

func (tradingSuite *AvaxTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
	tradingSuite.AVAXClient.Close()
}

func (tradingSuite *AvaxTestSuite) SetupTest() {
	fmt.Println("Setting up the test...")
}

func (tradingSuite *AvaxTestSuite) TearDownTest() {
	fmt.Println("Tearing down the test...")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestAvaxTestSuite(t *testing.T) {
	fmt.Println("Starting entry point for avax test suite...")

	tradingSuite := new(TradingTestSuite)
	suite.Run(t, tradingSuite)

	avaxSuite := NewAvaxTestSuite(tradingSuite)
	suite.Run(t, avaxSuite)

	fmt.Println("Finishing entry point for 0x test suite...")
}

func (tradingSuite *AvaxTestSuite) Test1Avax() {
	fmt.Println("============ TEST SHIELD UNSHIELD NATIVE AVAX ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
	burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("------------ STEP 1: porting ETH to pETH --------------")
	txHash := tradingSuite.depositETHCompliance(
		tradingSuite.DepositingEther,
		tradingSuite.IncPaymentAddrStr,
		tradingSuite.VaultAVAXAddr,
		tradingSuite.AVAXClient,
	)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.AVAXHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(230 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingavaxreq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)

	fmt.Println("------------ STEP 2: burning pETH to deposit ETH to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningPToken(
		tradingSuite.IncEtherTokenIDStr,
		burningPETH,
		pubKeyToAddrStr[2:],
		"createandsendburningavaxfordeposittoscrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForDepositToSC(
		burningTxID.(string),
		big.NewInt(int64(tradingSuite.ChainIDAVAX)),
		"getburnavaxprooffordeposittosc",
		tradingSuite.VaultAVAXAddr,
		tradingSuite.AVAXClient,
	)
	deposited := tradingSuite.getDepositedBalanceWithParams(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		pubKeyToAddrStr,
		tradingSuite.VaultAVAXAddr,
		tradingSuite.AVAXClient,
	)
	fmt.Println("deposited AVAX: ", deposited)

	fmt.Println("------------ step 3: withdrawing ETH from SC to pETH on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdrawCompliance(
		tradingSuite.EtherAddressStr,
		deposited,
		tradingSuite.AVAXClient,
		big.NewInt(int64(tradingSuite.ChainIDAVAX)),
		tradingSuite.VaultAVAXAddr,
		REQ_WITHDRAW_PREFIX,
	)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.AVAXHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(230 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingavaxreq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)

	fmt.Println("------------ step 4: withdrawing pAVX from Incognito to AVAX --------------")
	withdrawingPAVAX := big.NewInt(0).Div(deposited, big.NewInt(1000000000))
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncEtherTokenIDStr,
		withdrawingPAVAX,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningavaxrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForWithdrawal(
		burningTxID.(string),
		"getavaxburnproof",
		tradingSuite.VaultAVAXAddr,
		tradingSuite.AVAXClient,
		tradingSuite.ChainIDAVAX,
	)

	bal := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		tradingSuite.AVAXClient,
	)
	tradingSuite.AVAXBalanceAfterStep1 = bal
	fmt.Println("AVAX balance after step 1: ", tradingSuite.AVAXBalanceAfterStep1)
}

func (tradingSuite *AvaxTestSuite) Test2AvaxTokeb() {
	fmt.Println("============ TEST SHIELD UNSHIELD TOKEN AVAX ===========")
	fmt.Println("------------ step 0: declaration & initialization --------------")
	depositingUSDT := big.NewInt(3e6)
	burningPUSDT := depositingUSDT

	daibal := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.USDTTokenAddress),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		tradingSuite.AVAXClient,
	)
	fmt.Println("balance balance of owner: ", daibal)

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("pubKeyToAddrStr: ", pubKeyToAddrStr)

	fmt.Println("------------ step 1: porting USDT to pUSDT --------------")
	txHash := tradingSuite.depositERC20ComplianceToBridge(
		depositingUSDT,
		common.HexToAddress(tradingSuite.USDTTokenAddress),
		tradingSuite.IncPaymentAddrStr,
		tradingSuite.VaultAVAXAddr,
		tradingSuite.AVAXClient,
		tradingSuite.ChainIDAVAX,
	)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.AVAXHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(230 * time.Second)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncUSDTTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingavaxreq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)

	fmt.Println("------------ step 2: burning pUSDT to deposit USDT to SC --------------")

	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningPToken(
		tradingSuite.IncUSDTTokenIDStr,
		burningPUSDT,
		pubKeyToAddrStr[2:],
		"createandsendburningavaxfordeposittoscrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForDepositToSC(
		burningTxID.(string),
		big.NewInt(int64(tradingSuite.ChainIDAVAX)),
		"getburnavaxprooffordeposittosc",
		tradingSuite.VaultAVAXAddr,
		tradingSuite.AVAXClient,
	)
	deposited := tradingSuite.getDepositedBalanceWithParams(
		common.HexToAddress(tradingSuite.USDTTokenAddress),
		pubKeyToAddrStr,
		tradingSuite.VaultAVAXAddr,
		tradingSuite.AVAXClient,
	)
	fmt.Println("deposited dai: ", deposited)

	fmt.Println("------------ step 3: withdrawing USDT from SC to pUSDT on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdrawCompliance(
		tradingSuite.USDTTokenAddress,
		deposited,
		tradingSuite.AVAXClient,
		big.NewInt(int64(tradingSuite.ChainIDAVAX)),
		tradingSuite.VaultAVAXAddr,
		REQ_WITHDRAW_PREFIX,
	)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.AVAXHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(230 * time.Second)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncUSDTTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingavaxreq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)

	fmt.Println("------------ step 4: withdrawing pUSDT from Incognito to USDT --------------")
	withdrawingPMRK := deposited
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncUSDTTokenIDStr,
		withdrawingPMRK,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningavaxrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForWithdrawal(
		burningTxID.(string),
		"getavaxburnproof",
		tradingSuite.VaultAVAXAddr,
		tradingSuite.AVAXClient,
		tradingSuite.ChainIDAVAX,
	)

	bal := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.USDTTokenAddress),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		tradingSuite.AVAXClient,
	)
	tradingSuite.USDTBalanceAfterStep2 = bal
	fmt.Println("USDT balance after step 2: ", tradingSuite.USDTBalanceAfterStep2)
}
