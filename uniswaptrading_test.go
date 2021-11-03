package main

import (
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/incognitochain/bridge-eth/bridge/uniswap"

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
type UniswapTradingTestSuite struct {
	*TradingTestSuite

	UniswapTradeDeployedAddr common.Address
	UniswapRouteContractAddr common.Address

	IncMRKTokenIDStr string
	WETHAddr         string
	MRKAddressStr    string

	// token amounts for tests
	DepositingEther      float64
	DAIBalanceAfterStep1 *big.Int
	MRKBalanceAfterStep2 *big.Int
}

func NewUniswapTradingTestSuite(tradingTestSuite *TradingTestSuite) *UniswapTradingTestSuite {
	return &UniswapTradingTestSuite{
		TradingTestSuite: tradingTestSuite,
	}
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (tradingSuite *UniswapTradingTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")
	// Kovan env
	tradingSuite.IncMRKTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000062"
	tradingSuite.MRKAddressStr = "0xef13c0c8abcaf5767160018d268f9697ae4f5375"                                 // kovan
	tradingSuite.WETHAddr = "0xd0a1e359811322d97991e03f863a0c30c2cf029c"                                      // kovan
	tradingSuite.UniswapTradeDeployedAddr = common.HexToAddress("0x31DC0CAEdee768e0b699c6AdD7A11BB128E0Fb77") //kovan
	tradingSuite.UniswapRouteContractAddr = common.HexToAddress("0xf164fC0Ec4E93095b804a4795bBe1e041497b92a")
	tradingSuite.DepositingEther = float64(0.05)
}

func (tradingSuite *UniswapTradingTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
	tradingSuite.ETHClient.Close()
}

func (tradingSuite *UniswapTradingTestSuite) SetupTest() {
	fmt.Println("Setting up the test...")
}

func (tradingSuite *UniswapTradingTestSuite) TearDownTest() {
	fmt.Println("Tearing down the test...")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestUniswapTradingTestSuite(t *testing.T) {
	fmt.Println("Starting entry point for Uniswap test suite...")

	tradingSuite := new(TradingTestSuite)
	suite.Run(t, tradingSuite)

	uniswapTradingSuite := NewUniswapTradingTestSuite(tradingSuite)
	suite.Run(t, uniswapTradingSuite)

	fmt.Println("Finishing entry point for 0x test suite...")
}

func (tradingSuite *UniswapTradingTestSuite) getExpectedAmount(
	srcToken string,
	destToken string,
	srcQty *big.Int,
) *big.Int {
	if srcToken == tradingSuite.EtherAddressStr {
		srcToken = tradingSuite.WETHAddr
	}
	if destToken == tradingSuite.EtherAddressStr {
		destToken = tradingSuite.WETHAddr
	}
	c, err := uniswap.NewUniswapV2Trade(tradingSuite.UniswapTradeDeployedAddr, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)
	amounts, err := c.GetAmountsOut(nil, common.HexToAddress(srcToken), srcQty, common.HexToAddress(destToken))
	require.Equal(tradingSuite.T(), nil, err)
	require.Equal(tradingSuite.T(), 2, len(amounts))
	fmt.Printf("intput value: %d\n", amounts[0])
	fmt.Printf("output value: %d\n", amounts[1])
	return amounts[1]
}

func (tradingSuite *UniswapTradingTestSuite) executeWithUniswap(
	srcQty *big.Int,
	srcTokenIDStr string,
	destTokenIDStr string,
) {
	tradeAbi, _ := abi.JSON(strings.NewReader(uniswap.UniswapV2TradeABI))

	// Get contract instance
	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)
	auth := bind.NewKeyedTransactor(tradingSuite.ETHPrivKey)
	auth.GasPrice = big.NewInt(50000000000)
	// auth.GasLimit = 2000000
	srcToken := common.HexToAddress(srcTokenIDStr)
	destToken := common.HexToAddress(destTokenIDStr)
	expectOutputAmount := tradingSuite.getExpectedAmount(srcTokenIDStr, destTokenIDStr, srcQty)
	input, _ := tradeAbi.Pack("trade", srcToken, srcQty, destToken, expectOutputAmount)
	timestamp := []byte(randomizeTimestamp())
	vaultAbi, _ := abi.JSON(strings.NewReader(vault.VaultHelperABI))
	psData := vault.VaultHelperPreSignData{
		Prefix: EXECUTE_PREFIX,
		Token: srcToken,
		Timestamp: timestamp,
		Amount: srcQty,
	}
	tempData, _ := vaultAbi.Pack("_buildSignExecute", psData, destToken, tradingSuite.UniswapTradeDeployedAddr, input)
	data := rawsha3(tempData[4:])
	signBytes, _ := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)

	tx, err := c.Execute(
		auth,
		srcToken,
		srcQty,
		destToken,
		tradingSuite.UniswapTradeDeployedAddr,
		input,
		timestamp,
		signBytes,
	)
	require.Equal(tradingSuite.T(), nil, err)
	txHash := tx.Hash()
	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("Uniswap trade executed , txHash: %x\n", txHash[:])
}

func (tradingSuite *UniswapTradingTestSuite) Test1TradeEthForDAIWithUniswap() {
	fmt.Println("============ TEST TRADE ETHER FOR DAI WITH uniswap AGGREGATOR ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
	burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("------------ STEP 1: porting ETH to pETH --------------")
	txHash := tradingSuite.depositETH(
		tradingSuite.DepositingEther,
		tradingSuite.IncPaymentAddrStr,
		tradingSuite.VaultBSCAddr,
		tradingSuite.BSCClient,
	)
	// time.Sleep(15 * time.Second)
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(230 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingethreq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)

	fmt.Println("------------ STEP 2: burning pETH to deposit ETH to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningPToken(
		tradingSuite.IncEtherTokenIDStr,
		burningPETH,
		pubKeyToAddrStr[2:],
		"createandsendburningfordeposittoscrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForDepositToSC(
		burningTxID.(string),
		big.NewInt(int64(tradingSuite.ChainIDETH)),
		"getburnprooffordeposittosc",
		tradingSuite.VaultAddr,
		tradingSuite.ETHClient,
	)
	deposited := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("deposited EHT: ", deposited)
	// require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPETH, big.NewInt(1000000000)), deposited)

	fmt.Println("------------ step 3: execute trade ETH for DAI through Uniswap aggregator --------------")
	tradingSuite.executeWithUniswap(
		deposited,
		tradingSuite.EtherAddressStr,
		tradingSuite.DAIAddressStr,
	)
	time.Sleep(15 * time.Second)
	daiTraded := tradingSuite.getDepositedBalance(
		tradingSuite.DAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("daiTraded: ", daiTraded)

	fmt.Println("------------ step 4: withdrawing DAI from SC to pDAI on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.DAIAddressStr,
		daiTraded,
		tradingSuite.BSCClient,
		big.NewInt(int64(tradingSuite.ChainIDETH)),
		tradingSuite.VaultAddr,
		REQ_WITHDRAW_PREFIX,
	)
	time.Sleep(15 * time.Second)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(90 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncDAITokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingethreq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)

	fmt.Println("------------ step 5: withdrawing pDAI from Incognito to DAI --------------")
	withdrawingPDAI := big.NewInt(0).Div(daiTraded, big.NewInt(1000000000))
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncDAITokenIDStr,
		withdrawingPDAI,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForWithdrawal(
		burningTxID.(string),
		"getburnproof",
		tradingSuite.VaultAddr,
		tradingSuite.ETHClient,
		tradingSuite.ChainIDETH,
	)

	bal := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.DAIAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		tradingSuite.ETHClient,
	)
	tradingSuite.DAIBalanceAfterStep1 = bal
	fmt.Println("DAI balance after step 1: ", tradingSuite.DAIBalanceAfterStep1)
	// require.Equal(tradingSuite.T(), withdrawingPDAI.Uint64(), bal.Div(bal, big.NewInt(1000000000)).Uint64())
}

func (tradingSuite *UniswapTradingTestSuite) Test2TradeDAIForMRKWithUniswap() {
	fmt.Println("============ TEST TRADE DAI FOR MRK WITH UNISWAP AGGREGATOR ===========")
	fmt.Println("------------ step 0: declaration & initialization --------------")
	depositingDAI := tradingSuite.DAIBalanceAfterStep1
	burningPDAI := big.NewInt(0).Div(depositingDAI, big.NewInt(1000000000))
	tradeAmount := depositingDAI

	daibal := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.DAIAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		tradingSuite.ETHClient,
	)
	fmt.Println("dai balance of owner: ", daibal)

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("pubKeyToAddrStr: ", pubKeyToAddrStr)

	fmt.Println("------------ step 1: porting DAI to pDAI --------------")
	txHash := tradingSuite.depositERC20ToBridge(
		depositingDAI,
		common.HexToAddress(tradingSuite.DAIAddressStr),
		tradingSuite.IncPaymentAddrStr,
		tradingSuite.VaultAddr,
		tradingSuite.ETHClient,
		tradingSuite.ChainIDETH,
	)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(90 * time.Second)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncDAITokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingethreq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)
	// TODO: check the new balance on peth of the incognito account

	fmt.Println("------------ step 2: burning pDAI to deposit DAI to SC --------------")

	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningPToken(
		tradingSuite.IncDAITokenIDStr,
		burningPDAI,
		pubKeyToAddrStr[2:],
		"createandsendburningfordeposittoscrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForDepositToSC(
		burningTxID.(string),
		big.NewInt(int64(tradingSuite.ChainIDETH)),
		"getburnprooffordeposittosc",
		tradingSuite.VaultAddr,
		tradingSuite.ETHClient,
	)
	deposited := tradingSuite.getDepositedBalance(
		tradingSuite.DAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("deposited dai: ", deposited)
	// require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPDAI, big.NewInt(1000000000)), deposited)

	fmt.Println("------------ step 3: execute trade DAI for MRK through Uniswap aggregator --------------")
	tradingSuite.executeWithUniswap(
		tradeAmount,
		tradingSuite.DAIAddressStr,
		tradingSuite.MRKAddressStr,
	)
	time.Sleep(15 * time.Second)
	mrkTraded := tradingSuite.getDepositedBalance(
		tradingSuite.MRKAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("mrkTraded: ", mrkTraded)

	fmt.Println("------------ step 4: withdrawing MRK from SC to pMRK on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.MRKAddressStr,
		mrkTraded,
		tradingSuite.BSCClient,
		big.NewInt(int64(tradingSuite.ChainIDETH)),
		tradingSuite.VaultAddr,
		REQ_WITHDRAW_PREFIX,
	)
	time.Sleep(15 * time.Second)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(90 * time.Second)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncMRKTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingethreq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)

	fmt.Println("------------ step 5: withdrawing pMRK from Incognito to MRK --------------")
	withdrawingPMRK := big.NewInt(0).Div(mrkTraded, big.NewInt(1000000000)) // MRK decimal 18
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncMRKTokenIDStr,
		withdrawingPMRK,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForWithdrawal(
		burningTxID.(string),
		"getburnproof",
		tradingSuite.VaultAddr,
		tradingSuite.ETHClient,
		tradingSuite.ChainIDETH,
	)

	bal := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.MRKAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		tradingSuite.ETHClient,
	)
	tradingSuite.MRKBalanceAfterStep2 = bal
	fmt.Println("MRK balance after step 2: ", tradingSuite.MRKBalanceAfterStep2)
	// require.Equal(tradingSuite.T(), withdrawingPMRK.Uint64(), bal.Uint64())
	// require.Equal(tradingSuite.T(), withdrawingPMRK.Uint64(), bal.Div(bal, big.NewInt(1000000000)).Uint64())
}

func (tradingSuite *UniswapTradingTestSuite) Test3TradeMRKForEthWithUniswap() {
	fmt.Println("============ TEST TRADE SALT FOR ETH WITH UNISWAP AGGREGATOR ===========")
	fmt.Println("------------ step 0: declaration & initialization --------------")
	depositingMRK := tradingSuite.MRKBalanceAfterStep2
	burningPMRK := big.NewInt(0).Div(depositingMRK, big.NewInt(1000000000))
	tradeAmount := depositingMRK

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("pubKeyToAddrStr: ", pubKeyToAddrStr)

	fmt.Println("------------ step 1: porting MRK to pMRK --------------")
	txHash := tradingSuite.depositERC20ToBridge(
		depositingMRK,
		common.HexToAddress(tradingSuite.MRKAddressStr),
		tradingSuite.IncPaymentAddrStr,
		tradingSuite.VaultAddr,
		tradingSuite.ETHClient,
		tradingSuite.ChainIDETH,
	)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(90 * time.Second)

	issuingRes, err := tradingSuite.callIssuingETHReq(
		tradingSuite.IncMRKTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingethreq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("issuingRes: ", issuingRes)
	time.Sleep(120 * time.Second)

	fmt.Println("------------ step 2: burning pMRK to deposit MRK to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningPToken(
		tradingSuite.IncMRKTokenIDStr,
		burningPMRK,
		pubKeyToAddrStr[2:],
		"createandsendburningfordeposittoscrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(140 * time.Second)

	tradingSuite.submitBurnProofForDepositToSC(
		burningTxID.(string), 
		big.NewInt(int64(tradingSuite.ChainIDETH)),
		"getburnprooffordeposittosc",
		tradingSuite.VaultAddr,
		tradingSuite.ETHClient,
	)
	deposited := tradingSuite.getDepositedBalance(
		tradingSuite.MRKAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("deposited mrk: ", deposited)
	// require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPMRK, big.NewInt(1000000000)), deposited)

	fmt.Println("------------ step 3: execute trade MRK for ETH through Uniswap aggregator --------------")
	tradingSuite.executeWithUniswap(
		tradeAmount,
		tradingSuite.MRKAddressStr,
		tradingSuite.EtherAddressStr,
	)
	etherTraded := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("etherTraded: ", etherTraded)

	fmt.Println("------------ step 4: withdrawing ETH from SC to pETH on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.EtherAddressStr,
		etherTraded,
		tradingSuite.BSCClient,
		big.NewInt(int64(tradingSuite.ChainIDETH)),
		tradingSuite.VaultAddr,
		REQ_WITHDRAW_PREFIX,
	)
	time.Sleep(15 * time.Second)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(90 * time.Second)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingethreq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(140 * time.Second)

	fmt.Println("------------ step 5: withdrawing pETH from Incognito to ETH --------------")
	withdrawingPETH := big.NewInt(0).Div(etherTraded, big.NewInt(1000000000))
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncEtherTokenIDStr,
		withdrawingPETH,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(140 * time.Second)

	tradingSuite.submitBurnProofForWithdrawal(
		burningTxID.(string),
		"getburnproof",
		tradingSuite.VaultAddr,
		tradingSuite.ETHClient,
		tradingSuite.ChainIDETH,
	)

	bal := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		tradingSuite.ETHClient,
	)
	fmt.Println("Ether balance after step 3: ", bal)
	// require.Equal(tradingSuite.T(), withdrawingPETH.Uint64(), bal.Div(bal, big.NewInt(1000000000)).Uint64())
}
