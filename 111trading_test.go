   package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"math/big"
// 	"net/http"
// 	"strings"
// 	"testing"
// 	"time"

// 	"github.com/incognitochain/bridge-eth/bridge/vault"
// 	"github.com/incognitochain/bridge-eth/bridge/zrxtrade"
// 	"github.com/pkg/errors"
// 	"github.com/stretchr/testify/require"
// 	"github.com/stretchr/testify/suite"

// 	"github.com/ethereum/go-ethereum/accounts/abi"
// 	"github.com/ethereum/go-ethereum/accounts/abi/bind"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/crypto"
// 	"github.com/ethereum/go-ethereum/params"
// )

// // Define the suite, and absorb the built-in basic suite
// // functionality from testify - including assertion methods.
// type ZRXTradingTestSuite struct {
// 	*TradingTestSuite

// 	ZRXTradeDeployedAddr common.Address

// 	ZRXContractAddr common.Address
// 	WETHAddr        common.Address

// 	Quote0xUrl string

// 	// token amounts for tests
// 	DepositingEther      float64
// 	DAIBalanceAfterStep1 *big.Int
// 	SAIBalanceAfterStep2 *big.Int
// }

// func NewZRXTradingTestSuite(tradingTestSuite *TradingTestSuite) *ZRXTradingTestSuite {
// 	return &ZRXTradingTestSuite{
// 		TradingTestSuite: tradingTestSuite,
// 	}
// }

// // Make sure that VariableThatShouldStartAtFive is set to five
// // before each test
// func (tradingSuite *ZRXTradingTestSuite) SetupSuite() {
// 	fmt.Println("Setting up the suite...")
// 	// 0x kovan env
// 	tradingSuite.ZRXContractAddr = common.HexToAddress("0xf1ec01d6236d3cd881a0bf0130ea25fe4234003e")
// 	tradingSuite.WETHAddr = common.HexToAddress("0xd0a1e359811322d97991e03f863a0c30c2cf029c")

// 	tradingSuite.ZRXTradeDeployedAddr = common.HexToAddress("0x2ABEF99B499396D5e722702A5800F4eb0C5Bc7C8")

// 	tradingSuite.Quote0xUrl = "https://kovan.api.0x.org/swap/v0/quote?sellToken=%v&buyToken=%v&sellAmount=%v"

// 	tradingSuite.DepositingEther = float64(0.05)
// }

// func (tradingSuite *ZRXTradingTestSuite) TearDownSuite() {
// 	fmt.Println("Tearing down the suite...")
// 	tradingSuite.ETHClient.Close()
// }

// func (tradingSuite *ZRXTradingTestSuite) SetupTest() {
// 	fmt.Println("Setting up the test...")
// }

// func (tradingSuite *ZRXTradingTestSuite) TearDownTest() {
// 	fmt.Println("Tearing down the test...")
// }

// // In order for 'go test' to run this suite, we need to create
// // a normal test function and pass our suite to suite.Run
// func TestZRXTradingTestSuite(t *testing.T) {
// 	fmt.Println("Starting entry point for 0x test suite...")

// 	tradingSuite := new(TradingTestSuite)
// 	suite.Run(t, tradingSuite)

// 	zrxTradingSuite := NewZRXTradingTestSuite(tradingSuite)
// 	suite.Run(t, zrxTradingSuite)

// 	fmt.Println("Finishing entry point for 0x test suite...")
// }

// func (tradingSuite *ZRXTradingTestSuite) executeWith0x(
// 	srcQty *big.Int,
// 	srcTokenName string,
// 	srcTokenIDStr string,
// 	destTokenName string,
// 	destTokenIDStr string,
// ) {
// 	tradeAbi, _ := abi.JSON(strings.NewReader(zrxtrade.ZrxtradeABI))

// 	// Get contract instance
// 	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	auth := bind.NewKeyedTransactor(tradingSuite.ETHPrivKey)
// 	auth.GasPrice = big.NewInt(50000000000)
// 	auth.GasLimit = 2000000
// 	// quote
// 	srcToken := common.HexToAddress(srcTokenIDStr)
// 	destToken := common.HexToAddress(destTokenIDStr)

// 	quoteData, _ := quote0x(tradingSuite.Quote0xUrl, srcTokenName, destTokenName, srcQty)
// 	forwarder := common.HexToAddress(quoteData["to"].(string))
// 	dt := common.Hex2Bytes(quoteData["data"].(string)[2:])
// 	auth.Value, _ = big.NewInt(0).SetString(quoteData["protocolFee"].(string), 10)
// 	auth.GasPrice, _ = big.NewInt(0).SetString(quoteData["gasPrice"].(string), 10)
// 	input, _ := tradeAbi.Pack("trade", srcToken, srcQty, destToken, dt, forwarder)
// 	timestamp := []byte(randomizeTimestamp())
// 	tempData := append(tradingSuite.ZRXTradeDeployedAddr[:], input...)
// 	tempData1 := append(tempData, timestamp...)
// 	data := rawsha3(tempData1)
// 	signBytes, _ := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)

// 	tx, err := c.Execute(
// 		auth,
// 		srcToken,
// 		srcQty,
// 		destToken,
// 		tradingSuite.ZRXTradeDeployedAddr,
// 		input,
// 		timestamp,
// 		signBytes,
// 	)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	txHash := tx.Hash()
// 	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
// 		require.Equal(tradingSuite.T(), nil, err)
// 	}
// 	fmt.Printf("0x trade executed , txHash: %x\n", txHash[:])
// }

// func quote0x(
// 	quote0xUrl string,
// 	srcToken, destToken string,
// 	srcQty *big.Int,
// ) (map[string]interface{}, error) {
// 	var (
// 		err       error
// 		resp      *http.Response
// 		bodyBytes []byte
// 		result    interface{}
// 	)
// 	url := fmt.Sprintf(quote0xUrl, srcToken, destToken, srcQty.String())
// 	if resp, err = http.Get(url); err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()
// 	if resp.StatusCode != http.StatusOK {
// 		return nil, errors.New("Request returns with fucking error!!!")
// 	}
// 	if bodyBytes, err = ioutil.ReadAll(resp.Body); err != nil {
// 		return nil, err
// 	}
// 	if err = json.Unmarshal(bodyBytes, &result); err != nil {
// 		return nil, err
// 	}
// 	return result.(map[string]interface{}), nil
// }

// func (tradingSuite *ZRXTradingTestSuite) Test1TradeEthForDaiWith0x() {

// 	fmt.Println("============ TEST TRADE ETHER FOR DAI WITH 0X AGGREGATOR ===========")
// 	fmt.Println("------------ STEP 0: declaration & initialization --------------")
// 	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
// 	burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))

// 	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()

// 	fmt.Println("------------ STEP 1: porting ETH to pETH --------------")
// 	txHash := tradingSuite.depositETH(
// 		tradingSuite.DepositingEther,
// 		tradingSuite.IncPaymentAddrStr,
// 	)

// 	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

// 	fmt.Println("Waiting 90s for 15 blocks confirmation")
// 	time.Sleep(10 * time.Second)
// 	_, err = tradingSuite.callIssuingETHReq(
// 		tradingSuite.IncEtherTokenIDStr,
// 		ethDepositProof,
// 		ethBlockHash,
// 		ethTxIdx,
// 	)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	time.Sleep(90 * time.Second)

// 	fmt.Println("------------ STEP 2: burning pETH to deposit ETH to SC --------------")
// 	// make a burn tx to incognito chain as a result of deposit to SC
// 	burningRes, err := tradingSuite.callBurningPToken(
// 		tradingSuite.IncEtherTokenIDStr,
// 		burningPETH,
// 		pubKeyToAddrStr[2:],
// 		"createandsendburningfordeposittoscrequest",
// 	)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	burningTxID, found := burningRes["TxID"]
// 	require.Equal(tradingSuite.T(), true, found)
// 	time.Sleep(90 * time.Second)

// 	tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))
// 	deposited := tradingSuite.getDepositedBalance(
// 		tradingSuite.EtherAddressStr,
// 		pubKeyToAddrStr,
// 	)
// 	fmt.Println("deposited EHT: ", deposited)
// 	//require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPETH, big.NewInt(1000000000)), deposited)

// 	fmt.Println("------------ step 3: execute trade ETH for DAI through 0x aggregator --------------")
// 	tradingSuite.executeWith0x(
// 		tradeAmount,
// 		"ETH",
// 		tradingSuite.EtherAddressStr,
// 		"DAI",
// 		tradingSuite.DAIAddressStr,
// 	)
// 	daiTraded := tradingSuite.getDepositedBalance(
// 		tradingSuite.DAIAddressStr,
// 		pubKeyToAddrStr,
// 	)
// 	fmt.Println("daiTraded: ", daiTraded)

// 	fmt.Println("------------ step 4: withdrawing DAI from SC to pDAI on Incognito --------------")
// 	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
// 		tradingSuite.DAIAddressStr,
// 		daiTraded,
// 	)

// 	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

// 	fmt.Println("Waiting 90s for 15 blocks confirmation")
// 	time.Sleep(10 * time.Second)
// 	_, err = tradingSuite.callIssuingETHReq(
// 		tradingSuite.IncDAITokenIDStr,
// 		ethDepositProof,
// 		ethBlockHash,
// 		ethTxIdx,
// 	)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	time.Sleep(90 * time.Second)

// 	fmt.Println("------------ step 5: withdrawing pDAI from Incognito to DAI --------------")
// 	withdrawingPDAI := big.NewInt(0).Div(daiTraded, big.NewInt(1000000000))
// 	burningRes, err = tradingSuite.callBurningPToken(
// 		tradingSuite.IncDAITokenIDStr,
// 		withdrawingPDAI,
// 		tradingSuite.ETHOwnerAddrStr,
// 		"createandsendburningrequest",
// 	)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	burningTxID, found = burningRes["TxID"]
// 	require.Equal(tradingSuite.T(), true, found)
// 	time.Sleep(90 * time.Second)

// 	tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))

// 	bal := tradingSuite.getBalanceOnETHNet(
// 		common.HexToAddress(tradingSuite.DAIAddressStr),
// 		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
// 	)
// 	tradingSuite.DAIBalanceAfterStep1 = bal
// 	fmt.Println("DAI balance after step 1: ", tradingSuite.DAIBalanceAfterStep1)
// 	// require.Equal(tradingSuite.T(), withdrawingPDAI.Uint64(), bal.Div(bal, big.NewInt(1000000000)).Uint64())
// }

// func (tradingSuite *ZRXTradingTestSuite) Test2TradeDaiForSaiWith0x() {
// 	fmt.Println("============ TEST TRADE DAI FOR SAI WITH 0X AGGREGATOR ===========")
// 	fmt.Println("------------ step 0: declaration & initialization --------------")
// 	// bal111 := tradingSuite.getBalanceOnETHNet(
// 	// 	common.HexToAddress(tradingSuite.DAIAddressStr),
// 	// 	common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
// 	// )
// 	depositingDAI := tradingSuite.DAIBalanceAfterStep1
// 	//depositingDAI := bal111
// 	burningPDAI := big.NewInt(0).Div(depositingDAI, big.NewInt(1000000000))
// 	tradeAmount := depositingDAI

// 	daibal := tradingSuite.getBalanceOnETHNet(
// 		common.HexToAddress(tradingSuite.DAIAddressStr),
// 		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
// 	)
// 	fmt.Println("dai balance of owner: ", daibal)

// 	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
// 	fmt.Println("pubKeyToAddrStr: ", pubKeyToAddrStr)

// 	fmt.Println("------------ step 1: porting DAI to pDAI --------------")
// 	txHash := tradingSuite.depositERC20ToBridge(
// 		depositingDAI,
// 		common.HexToAddress(tradingSuite.DAIAddressStr),
// 		tradingSuite.IncPaymentAddrStr,
// 	)

// 	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

// 	fmt.Println("Waiting 90s for 15 blocks confirmation")
// 	time.Sleep(10 * time.Second)

// 	_, err = tradingSuite.callIssuingETHReq(
// 		tradingSuite.IncDAITokenIDStr,
// 		ethDepositProof,
// 		ethBlockHash,
// 		ethTxIdx,
// 	)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	time.Sleep(90 * time.Second)
// 	// TODO: check the new balance on peth of the incognito account

// 	fmt.Println("------------ step 2: burning pDAI to deposit DAI to SC --------------")

// 	// make a burn tx to incognito chain as a result of deposit to SC
// 	burningRes, err := tradingSuite.callBurningPToken(
// 		tradingSuite.IncDAITokenIDStr,
// 		burningPDAI,
// 		pubKeyToAddrStr[2:],
// 		"createandsendburningfordeposittoscrequest",
// 	)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	burningTxID, found := burningRes["TxID"]
// 	require.Equal(tradingSuite.T(), true, found)
// 	time.Sleep(90 * time.Second)

// 	tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))
// 	deposited := tradingSuite.getDepositedBalance(
// 		tradingSuite.DAIAddressStr,
// 		pubKeyToAddrStr,
// 	)
// 	fmt.Println("deposited dai: ", deposited)
// 	// require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPDAI, big.NewInt(1000000000)), deposited)

// 	fmt.Println("------------ step 3: execute trade DAI for SAI through 0x aggregator --------------")
// 	tradingSuite.executeWith0x(
// 		tradeAmount,
// 		"DAI",
// 		tradingSuite.DAIAddressStr,
// 		"SAI",
// 		tradingSuite.SAIAddressStr,
// 	)
// 	saiTraded := tradingSuite.getDepositedBalance(
// 		tradingSuite.SAIAddressStr,
// 		pubKeyToAddrStr,
// 	)
// 	fmt.Println("saiTraded: ", saiTraded)

// 	fmt.Println("------------ step 4: withdrawing SAI from SC to pSAI on Incognito --------------")
// 	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
// 		tradingSuite.SAIAddressStr,
// 		saiTraded,
// 	)

// 	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

// 	fmt.Println("Waiting 90s for 15 blocks confirmation")
// 	time.Sleep(10 * time.Second)

// 	_, err = tradingSuite.callIssuingETHReq(
// 		tradingSuite.IncSAITokenIDStr,
// 		ethDepositProof,
// 		ethBlockHash,
// 		ethTxIdx,
// 	)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	time.Sleep(90 * time.Second)

// 	fmt.Println("------------ step 5: withdrawing pSAI from Incognito to SAI --------------")
// 	withdrawingPSAI := big.NewInt(0).Div(saiTraded, big.NewInt(1000000000))
// 	burningRes, err = tradingSuite.callBurningPToken(
// 		tradingSuite.IncSAITokenIDStr,
// 		withdrawingPSAI,
// 		tradingSuite.ETHOwnerAddrStr,
// 		"createandsendburningrequest",
// 	)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	burningTxID, found = burningRes["TxID"]
// 	require.Equal(tradingSuite.T(), true, found)
// 	time.Sleep(90 * time.Second)

// 	tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))

// 	bal := tradingSuite.getBalanceOnETHNet(
// 		common.HexToAddress(tradingSuite.SAIAddressStr),
// 		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
// 	)
// 	tradingSuite.SAIBalanceAfterStep2 = bal
// 	fmt.Println("SAI balance after step 2: ", tradingSuite.SAIBalanceAfterStep2)
// 	// require.Equal(tradingSuite.T(), withdrawingPSAI.Uint64(), bal.Uint64())
// 	// require.Equal(tradingSuite.T(), withdrawingPSAI.Uint64(), bal.Div(bal, big.NewInt(1000000000)).Uint64())
// }

// func (tradingSuite *ZRXTradingTestSuite) Test3TradeSaiForEthWith0x() {
// return
// 	fmt.Println("============ TEST TRADE SAI FOR ETH WITH 0X AGGREGATOR ===========")
// 	fmt.Println("------------ step 0: declaration & initialization --------------")
// 	depositingSAI := tradingSuite.SAIBalanceAfterStep2
// 	burningPSAI := big.NewInt(0).Div(depositingSAI, big.NewInt(1000000000))
// 	tradeAmount := depositingSAI

// 	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
// 	fmt.Println("pubKeyToAddrStr: ", pubKeyToAddrStr)

// 	fmt.Println("------------ step 1: porting SAI to pSAI --------------")
// 	txHash := tradingSuite.depositERC20ToBridge(
// 		depositingSAI,
// 		common.HexToAddress(tradingSuite.SAIAddressStr),
// 		tradingSuite.IncPaymentAddrStr,
// 	)

// 	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

// 	fmt.Println("Waiting 90s for 15 blocks confirmation")
// 	time.Sleep(10 * time.Second)

// 	issuingRes, err := tradingSuite.callIssuingETHReq(
// 		tradingSuite.IncSAITokenIDStr,
// 		ethDepositProof,
// 		ethBlockHash,
// 		ethTxIdx,
// 	)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	fmt.Println("issuingRes: ", issuingRes)
// 	time.Sleep(90 * time.Second)

// 	fmt.Println("------------ step 2: burning pSAI to deposit SAI to SC --------------")
// 	// make a burn tx to incognito chain as a result of deposit to SC
// 	burningRes, err := tradingSuite.callBurningPToken(
// 		tradingSuite.IncSAITokenIDStr,
// 		burningPSAI,
// 		pubKeyToAddrStr[2:],
// 		"createandsendburningfordeposittoscrequest",
// 	)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	burningTxID, found := burningRes["TxID"]
// 	require.Equal(tradingSuite.T(), true, found)
// 	time.Sleep(90 * time.Second)

// 	tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))
// 	deposited := tradingSuite.getDepositedBalance(
// 		tradingSuite.SAIAddressStr,
// 		pubKeyToAddrStr,
// 	)
// 	require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPSAI, big.NewInt(1000000000)), deposited)

// 	fmt.Println("------------ step 3: execute trade SAI for ETH through 0x aggregator --------------")
// 	tradingSuite.executeWith0x(
// 		tradeAmount,
// 		"SAI",
// 		tradingSuite.SAIAddressStr,
// 		"ETH",
// 		tradingSuite.EtherAddressStr,
// 	)
// 	etherTraded := tradingSuite.getDepositedBalance(
// 		tradingSuite.EtherAddressStr,
// 		pubKeyToAddrStr,
// 	)
// 	fmt.Println("etherTraded: ", etherTraded)

// 	fmt.Println("------------ step 4: withdrawing ETH from SC to pETH on Incognito --------------")
// 	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
// 		tradingSuite.EtherAddressStr,
// 		etherTraded,
// 	)

// 	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

// 	fmt.Println("Waiting 90s for 15 blocks confirmation")
// 	time.Sleep(10 * time.Second)

// 	_, err = tradingSuite.callIssuingETHReq(
// 		tradingSuite.IncEtherTokenIDStr,
// 		ethDepositProof,
// 		ethBlockHash,
// 		ethTxIdx,
// 	)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	time.Sleep(90 * time.Second)

// 	fmt.Println("------------ step 5: withdrawing pETH from Incognito to ETH --------------")
// 	withdrawingPETH := big.NewInt(0).Div(etherTraded, big.NewInt(1000000000))
// 	burningRes, err = tradingSuite.callBurningPToken(
// 		tradingSuite.IncEtherTokenIDStr,
// 		withdrawingPETH,
// 		tradingSuite.ETHOwnerAddrStr,
// 		"createandsendburningrequest",
// 	)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	burningTxID, found = burningRes["TxID"]
// 	require.Equal(tradingSuite.T(), true, found)
// 	time.Sleep(90 * time.Second)

// 	tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))

// 	bal := tradingSuite.getBalanceOnETHNet(
// 		common.HexToAddress(tradingSuite.EtherAddressStr),
// 		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
// 	)
// 	fmt.Println("Ether balance after step 3: ", bal)
// 	// require.Equal(tradingSuite.T(), withdrawingPETH.Uint64(), bal.Div(bal, big.NewInt(1000000000)).Uint64())
// }
