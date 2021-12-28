package bridge

import (
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	pancakeproxy "github.com/incognitochain/bridge-eth/bridge/pancake"

	"github.com/incognitochain/bridge-eth/bridge/vault"
	"github.com/stretchr/testify/suite"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"

	"github.com/stretchr/testify/require"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type PolygonTestSuite struct {
	*TradingTestSuite

	CurveDeployedAddr common.Address
	CurveRouteContractAddr common.Address

	IncBUSDTokenIDStr string
	WBNBAddr          common.Address
	WBTCAddr          common.Address
	WBUSDAddr         common.Address
	WXRPAddr          common.Address
	WUSDTAddr         common.Address

	// token amounts for tests
	DepositingEther      float64
	DAIBalanceAfterStep1 *big.Int
	MRKBalanceAfterStep2 *big.Int
}

func NewPolygonTestSuite(tradingTestSuite *TradingTestSuite) *PolygonTestSuite {
	return &PolygonTestSuite{
		TradingTestSuite: tradingTestSuite,
	}
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (tradingSuite *PolygonTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")
	// Polygon testnet env
	tradingSuite.IncBUSDTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000062"
	tradingSuite.CurveDeployedAddr = common.HexToAddress("0xD772475E9CFA0F1134aef4B3A5BA83De014DEc45")
	tradingSuite.CurveRouteContractAddr = common.HexToAddress("0x9Ac64Cc6e4415144C455BD8E4837Fea55603e5c3")

	// tokens
	tradingSuite.WBNBAddr = common.HexToAddress("0xae13d989dac2f0debff460ac112a837c89baa7cd")
	tradingSuite.WBTCAddr = common.HexToAddress("0x6ce8da28e2f864420840cf74474eff5fd80e65b8")
	tradingSuite.WUSDTAddr = common.HexToAddress("0x7ef95a0fee0dd31b22626fa2e10ee6a223f8a684") // USDT
	tradingSuite.WBUSDAddr = common.HexToAddress("0x78867bbeef44f2326bf8ddd1941a4439382ef2a7")
	tradingSuite.WXRPAddr = common.HexToAddress("0xa83575490d7df4e2f47b7d38ef351a2722ca45b9")

	tradingSuite.DepositingEther = float64(0.001)
}

func (tradingSuite *PolygonTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
	tradingSuite.PLGClient.Close()
}

func (tradingSuite *PolygonTestSuite) SetupTest() {
	fmt.Println("Setting up the test...")
}

func (tradingSuite *PolygonTestSuite) TearDownTest() {
	fmt.Println("Tearing down the test...")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestPolygonTestSuite(t *testing.T) {
	fmt.Println("Starting entry point for Uniswap test suite...")

	tradingSuite := new(TradingTestSuite)
	suite.Run(t, tradingSuite)

	uniswapTradingSuite := NewPolygonTestSuite(tradingSuite)
	suite.Run(t, uniswapTradingSuite)

	fmt.Println("Finishing entry point for 0x test suite...")
}

func (tradingSuite *PolygonTestSuite) getExpectedAmount(
	path []common.Address,
	srcQty *big.Int,
) *big.Int {
	c, err := pancakeproxy.NewPancakeproxyrouter(tradingSuite.CurveRouteContractAddr, tradingSuite.PLGClient)
	require.Equal(tradingSuite.T(), nil, err)
	amounts, err := c.GetAmountsOut(nil, srcQty, path)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Printf("intput value: %v\n", srcQty.String())
	fmt.Printf("output value: %v\n", amounts[len(amounts)-1].String())
	return amounts[1]
}

func (tradingSuite *PolygonTestSuite) executeWithPancake(
	srcQty *big.Int,
	path []common.Address,
	deadline uint,
	isNative bool,
) {
	require.NotEqual(tradingSuite.T(), 0, len(path))

	tradeAbi, err := abi.JSON(strings.NewReader(pancakeproxy.PancakeproxyMetaData.ABI))
	require.Equal(tradingSuite.T(), nil, err)

	// Get contract instance
	c, err := vault.NewVault(tradingSuite.VaultPLGAddr, tradingSuite.PLGClient)
	require.Equal(tradingSuite.T(), nil, err)
	auth, err := bind.NewKeyedTransactorWithChainID(tradingSuite.ETHPrivKey, big.NewInt(int64(tradingSuite.ChainIDPLG)))
	require.Equal(tradingSuite.T(), nil, err)
	auth.GasPrice = big.NewInt(1e10)
	expectOutputAmount := tradingSuite.getExpectedAmount(path, srcQty)
	input, err := tradeAbi.Pack("trade", path, srcQty, expectOutputAmount, big.NewInt(int64(deadline)), isNative)
	require.Equal(tradingSuite.T(), nil, err)
	timestamp := []byte(randomizeTimestamp())
	vaultAbi, err := abi.JSON(strings.NewReader(vault.VaultHelperABI))
	require.Equal(tradingSuite.T(), nil, err)
	sourceToken := path[0]
	if path[0].String() == tradingSuite.WBNBAddr.String() {
		sourceToken = common.HexToAddress(tradingSuite.EtherAddressStr)
	}
	destToken := path[len(path)-1]
	if path[len(path)-1].String() == tradingSuite.WBNBAddr.String() && isNative {
		destToken = common.HexToAddress(tradingSuite.EtherAddressStr)
	}
	psData := vault.VaultHelperPreSignData{
		Prefix:    PLG_EXECUTE_PREFIX,
		Token:     sourceToken,
		Timestamp: timestamp,
		Amount:    srcQty,
	}
	tempData, err := vaultAbi.Pack("_buildSignExecute", psData, destToken, tradingSuite.CurveDeployedAddr, input)
	require.Equal(tradingSuite.T(), nil, err)
	data := rawsha3(tempData[4:])
	signBytes, err := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)
	require.Equal(tradingSuite.T(), nil, err)

	tx, err := c.Execute(
		auth,
		sourceToken,
		srcQty,
		destToken,
		tradingSuite.CurveDeployedAddr,
		input,
		timestamp,
		signBytes,
	)
	require.Equal(tradingSuite.T(), nil, err)
	txHash := tx.Hash()
	if err := wait(tradingSuite.PLGClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("Pancake trade executed , txHash: %x\n", txHash[:])
}

func (tradingSuite *PolygonTestSuite) Test1TradeEthForDAIWithPancake() {
	fmt.Println("============ TEST SHIELD UNSHIELD POLYGON ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
	burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("------------ STEP 1: porting MATIC to pMATIC --------------")
	txHash := tradingSuite.depositETH(
		tradingSuite.DepositingEther,
		tradingSuite.IncPaymentAddrStr,
		tradingSuite.VaultPLGAddr,
		tradingSuite.PLGClient,
	)
	time.Sleep(15 * time.Second)
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.PLGHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(50 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingplgreq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)

	fmt.Println("------------ STEP 2: burning pBNB to deposit BNB to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningPToken(
		tradingSuite.IncEtherTokenIDStr,
		burningPETH,
		pubKeyToAddrStr[2:],
		"createandsendburningplgfordeposittoscrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForDepositToSC(
		burningTxID.(string),
		big.NewInt(int64(tradingSuite.ChainIDPLG)),
		"getburnplgprooffordeposittosc",
		tradingSuite.VaultPLGAddr,
		tradingSuite.PLGClient,
	)
	deposited := tradingSuite.getDepositedBalancePLG(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		pubKeyToAddrStr,
	)
	fmt.Printf("address own asset %v \n", pubKeyToAddrStr)

	// fmt.Println("------------ step 3: execute trade BNB for USDT through Pancake --------------")
	// tradingSuite.executeWithPancake(
	// 	deposited,
	// 	[]common.Address{
	// 		tradingSuite.WBNBAddr,
	// 		tradingSuite.WBUSDAddr,
	// 		tradingSuite.WUSDTAddr,
	// 	},
	// 	uint(time.Now().Unix()+60000),
	// 	false,
	// )
	// time.Sleep(15 * time.Second)
	// daiTraded := tradingSuite.getDepositedBalancePLG(
	// 	tradingSuite.WUSDTAddr,
	// 	pubKeyToAddrStr,
	// )
	// fmt.Println("usdtTraded: ", daiTraded)

	fmt.Println("------------ step 3: withdrawing MATIC from SC to pMATIC on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.WUSDTAddr.String(),
		deposited,
		tradingSuite.PLGClient,
		big.NewInt(int64(tradingSuite.ChainIDPLG)),
		tradingSuite.VaultPLGAddr,
		PLG_REQ_WITHDRAW_PREFIX,
	)
	time.Sleep(45 * time.Second)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.PLGHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(50 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncUSDTTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingplgreq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)

	fmt.Println("------------ step 4: withdrawing pMATIC from Incognito to MATIC --------------")
	withdrawingPDAI := big.NewInt(0).Div(deposited, big.NewInt(1e9))
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncUSDTTokenIDStr,
		withdrawingPDAI,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningplgrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForWithdrawal(
		burningTxID.(string),
		"getplgburnproof",
		tradingSuite.VaultPLGAddr,
		tradingSuite.PLGClient,
		tradingSuite.ChainIDPLG,
	)

	bal := tradingSuite.getBalanceOnETHNet(
		tradingSuite.WUSDTAddr,
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		tradingSuite.PLGClient,
	)
	tradingSuite.DAIBalanceAfterStep1 = bal
	fmt.Println("USDT balance after step 1: ", tradingSuite.DAIBalanceAfterStep1)
}