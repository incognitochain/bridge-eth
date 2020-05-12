package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"testing"
	"time"
	"context"

	"github.com/incognitochain/bridge-eth/bridge/vault"
	"github.com/incognitochain/bridge-eth/bridge/zrxtrade"
	"github.com/stretchr/testify/suite"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum"

	"github.com/stretchr/testify/require"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type ZRXTradingTestSuite struct {
	*TradingTestSuite

	ZRXTradeDeployedAddr common.Address

	ZRXContractAddr common.Address
	WETHAddr        common.Address

	Quote0xUrl string
	Quote0xUrlSlipPagePer string

	// token amounts for tests
	DepositingEther      float64
	DAIBalanceAfterStep1 *big.Int
	SAIBalanceAfterStep2 *big.Int
}

func NewZRXTradingTestSuite(tradingTestSuite *TradingTestSuite) *ZRXTradingTestSuite {
	return &ZRXTradingTestSuite{
		TradingTestSuite: tradingTestSuite,
	}
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (tradingSuite *ZRXTradingTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")
	// 0x kovan env
	tradingSuite.ZRXContractAddr = common.HexToAddress("0xf1ec01d6236d3cd881a0bf0130ea25fe4234003e")
	tradingSuite.WETHAddr = common.HexToAddress("0xd0a1e359811322d97991e03f863a0c30c2cf029c")

	tradingSuite.ZRXTradeDeployedAddr = common.HexToAddress("0x4C6e4aC9a49b78EDc7e885ad0281fE9D84dc8341")

	tradingSuite.Quote0xUrl = "https://kovan.api.0x.org/swap/v0/quote?sellToken=%v&buyToken=%v&sellAmount=%v&slippagePercentage=0.0001"
	tradingSuite.Quote0xUrlSlipPagePer = "https://kovan.api.0x.org/swap/v0/quote?sellToken=%v&buyToken=%v&sellAmount=%v&slippagePercentage=0.02"
	tradingSuite.DepositingEther = float64(0.001)
	tradingSuite.DAIBalanceAfterStep1 = big.NewInt(int64(5 * params.Ether))
	tradingSuite.SAIBalanceAfterStep2 = big.NewInt(int64(5 * params.Ether))

}

func (tradingSuite *ZRXTradingTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
	tradingSuite.ETHClient.Close()
}

func (tradingSuite *ZRXTradingTestSuite) SetupTest() {
	fmt.Println("Setting up the test...")
}

func (tradingSuite *ZRXTradingTestSuite) TearDownTest() {
	fmt.Println("Tearing down the test...")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestZRXTradingTestSuite(t *testing.T) {
	fmt.Println("Starting entry point for 0x test suite...")

	tradingSuite := new(TradingTestSuite)
	suite.Run(t, tradingSuite)

	zrxTradingSuite := NewZRXTradingTestSuite(tradingSuite)
	suite.Run(t, zrxTradingSuite)

	fmt.Println("Finishing entry point for 0x test suite...")
}

func (tradingSuite *ZRXTradingTestSuite) executeWith0x(
	srcQty *big.Int,
	srcTokenName string,
	srcTokenIDStr string,
	destTokenName string,
	destTokenIDStr string,
) (common.Hash, *big.Int) {
	tradeAbi, _ := abi.JSON(strings.NewReader(zrxtrade.ZrxtradeABI))

	// Get contract instance
	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)
	auth := bind.NewKeyedTransactor(tradingSuite.ETHPrivKey)
	auth.GasPrice = big.NewInt(50000000000)
	//auth.GasLimit = 2000000

	// quote
	srcToken := common.HexToAddress(srcTokenIDStr)
	destToken := common.HexToAddress(destTokenIDStr)

	quoteData, _ := quote0x(tradingSuite.Quote0xUrl, srcTokenName, destTokenName, srcQty)
	fmt.Println("quoteData :  ",quoteData)
	fmt.Println("forwarder : " ,quoteData["to"].(string))
	forwarder := common.HexToAddress(quoteData["to"].(string))
	dt := common.Hex2Bytes(quoteData["data"].(string)[2:])
	//auth.Value, _ = big.NewInt(0).SetString(quoteData["protocolFee"].(string), 10)
	//auth.Value, _ = big.NewInt(0).SetString("100000000000000000", 10)
	auth.GasPrice, _ = big.NewInt(0).SetString(quoteData["gasPrice"].(string), 10)
	//auth.Value = big.NewInt(0).Mul(auth.GasPrice, big.NewInt(int64(150000)))
	auth.Value = big.NewInt(0).Mul(auth.GasPrice, big.NewInt(int64(150000*len(quoteData["orders"].([]interface{})))))
	input, _ := tradeAbi.Pack("trade", srcToken, srcQty, destToken, dt, forwarder)
	timestamp := []byte(randomizeTimestamp())
	tempData := append(tradingSuite.ZRXTradeDeployedAddr[:], input...)
	tempData1 := append(tempData, timestamp...)
	tempData2 := append(tempData1, common.LeftPadBytes(srcQty.Bytes(), 32)...)
	data := rawsha3(tempData2)
	signBytes, _ := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)
	///////////
	executeAbi, _ := abi.JSON(strings.NewReader(vault.VaultABI))
	dataEst, _ := executeAbi.Pack("execute", srcToken,
		srcQty,
		destToken,
		tradingSuite.ZRXTradeDeployedAddr,
		input,
		timestamp,
		signBytes,
	)
	bgCtx := context.Background()
	msg := ethereum.CallMsg{From: common.HexToAddress(tradingSuite.ETHOwnerAddrStr), To: &tradingSuite.VaultAddr, Value: auth.Value, Data: dataEst}
	gasLimit, err := tradingSuite.ETHClient.EstimateGas(bgCtx, msg)
	if err != nil {
		fmt.Errorf("failed to estimate gas needed: %v", err)
	}
	fmt.Println("Gas estimated: ", gasLimit)
	auth.GasLimit = gasLimit + 100000

	/////
	tx, err := c.Execute(
		auth,
		srcToken,
		srcQty,
		destToken,
		tradingSuite.ZRXTradeDeployedAddr,
		input,
		timestamp,
		signBytes,
	)
	require.Equal(tradingSuite.T(), nil, err)
	txHash := tx.Hash()
	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("0x trade executed , txHash: %x\n", txHash[:])
	return txHash, auth.Value
}

func quote0x(
	quote0xUrl string,
	srcToken, destToken string,
	srcQty *big.Int,
	//	slippagePer float32,
) (map[string]interface{}, error) {
	var (
		err       error
		resp      *http.Response
		bodyBytes []byte
		result    interface{}
	)
	url := fmt.Sprintf(quote0xUrl, srcToken, destToken, srcQty.String())
	if resp, err = http.Get(url); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Request returns with  error!!!")
	}
	if bodyBytes, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(bodyBytes, &result); err != nil {
		return nil, err
	}
	return result.(map[string]interface{}), nil
}


func (tradingSuite *ZRXTradingTestSuite) Test1TradeEthForDaiWith0x() {
return
	fmt.Println("============ TEST 1 TRADE ETHER FOR DAI WITH 0X AGGREGATOR ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
	burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()

	// get info balance initialization
	balPrvInit, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance initialization : ", balPrvInit)

	balpEthInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance initialization : ", balpEthInit)

	balpDaiInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncDAITokenIDStr)
	fmt.Println("[INFO] pDAI balance initialization : ", balpDaiInit)

	balEthInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance initialization : ", balEthInit)

	balDaiInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.DAIAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] DAI balance initialization : ", balDaiInit)

	balEthScInit := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance initialization on SC : ", balEthScInit)

	balDaiScInit := tradingSuite.getDepositedBalance(
		tradingSuite.DAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] DAI balance initialization on SC : ", balDaiScInit)

	fmt.Println("------------ STEP 1: porting ETH to pETH --------------")
	fmt.Println("amount ETH deposit : ", (big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))))
	// Deposit to proof
	txHash := tradingSuite.depositETH(
		tradingSuite.DepositingEther,
		tradingSuite.IncPaymentAddrStr,
	)

	fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHash))
	// get ETH remain after depsit
	balEthAfDep := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain after deposit : ", balEthAfDep)

	require.Equal(tradingSuite.T(), balEthAfDep, big.NewInt(0).Sub(big.NewInt(0).Sub(balEthInit, big.NewInt(int64(tradingSuite.DepositingEther*params.Ether))), tradingSuite.getGasFeeETHbyTxhash(txHash)), "balance ETH incorrect")

	// get DepositProof
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	// wait to make sure 
	time.Sleep(100 * time.Second)

	// issuu pETH from incognito chain
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	// check call rpc issuu error
	require.Equal(tradingSuite.T(), nil, err)

	// wait to get block
	time.Sleep(120 * time.Second)

	// check PRV and token balance after issuing 
	balPrvAfIssS1, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after issuing step 1: ", balPrvAfIssS1)
	// dev accept fee_actual > fee_expect_of_rpc return
	//require.NotEqual(tradingSuite.T(), balPrvAfIssS1, (balPrvInit - tradingSuite.getFeePRVbyTxhashInC(issuuRes["TxID"].(string))), "Balance PRV remain incorrect after issuu step 1")

	balpEthAfIssS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance after issuing step 1 : ", balpEthAfIssS1)
	// check pToken balane affter issuu
	require.Equal(tradingSuite.T(), big.NewInt(int64(balpEthAfIssS1-balpEthInit)), big.NewInt(0).Div(big.NewInt(int64(tradingSuite.DepositingEther*params.Ether)), big.NewInt(1000000000)), " balnce pToken issuu incorrect")

	fmt.Println("------------ STEP 2: burning pETH to deposit ETH to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningPToken(
		tradingSuite.IncEtherTokenIDStr,
		burningPETH,
		pubKeyToAddrStr[2:],
		"createandsendburningfordeposittoscrequest",
	)
	// verify call prc success
	require.Equal(tradingSuite.T(), nil, err)

	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)

	time.Sleep(120 * time.Second)

	// check PRV and token balance after burning
	balPrvAfBurnS2, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after burning step 2: ", balPrvAfBurnS2)
	//require.NotEqual(tradingSuite.T(), balPrvAfBurnS2, (balPrvAfIssS1 - tradingSuite.getFeePRVbyTxhashInC(burningRes["TxID"].(string))), "Balance PRV remain incorrect after burn step 2")

	balpEthAfBurnS2, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance after burning step 2 : ", balpEthAfBurnS2)
	// TODO assert pETH balance issuing

	txHash2 := tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))
	fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHash2))
	// get ETH remain
	balEthAfDep2 := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain  : ", balEthAfDep2)
	// assert ETH balance affter pay fee to submitBurnProof
	require.Equal(tradingSuite.T(), balEthAfDep2, big.NewInt(0).Sub(balEthAfDep, tradingSuite.getGasFeeETHbyTxhash(txHash2)), "balance ETH incorrect")

	balEthScDepS2 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance after deposit on SC at step 2: ", balEthScDepS2)
	// TODO assert ETH balane on SC
	balDaiScS2 := tradingSuite.getDepositedBalance(
		tradingSuite.DAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] DAI balance on SC at step 2 : ", balDaiScS2)

	//require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPETH, big.NewInt(1000000000)), deposited)

	fmt.Println("------------ STEP 3: execute trade ETH for DAI through 0x aggregator --------------")
	txHash3, protoFee := tradingSuite.executeWith0x(
		tradeAmount,
		"ETH",
		tradingSuite.EtherAddressStr,
		"DAI",
		tradingSuite.DAIAddressStr,
	)
	fmt.Println("protocol fee :", protoFee)
	fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHash3))
	// get ETH remain
	balEthAfDep3 := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain  : ", balEthAfDep3)
	//require.Equal(tradingSuite.T(), balEthAfDep3, big.NewInt(0).Sub(big.NewInt(0).Sub(balEthAfDep2, tradingSuite.getGasFeeETHbyTxhash(txHash3)), protoFee), "balance ETH incorrect")
	balEthScTradeS3 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC after trade at step 3 : ", balEthScTradeS3)
	// TODO assert ETH balane on SC
	balDaiScTradeS3 := tradingSuite.getDepositedBalance(
		tradingSuite.DAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] DAI balance on SC after trade at step 3 : ", balDaiScTradeS3)
	//TODO assert DAI balane on SC
	require.NotEqual(tradingSuite.T(), balDaiScTradeS3, balDaiScS2, "trade failed")

	fmt.Println("------------ STEP 4: withdrawing DAI from SC to pDAI on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.DAIAddressStr,
		balDaiScTradeS3,
	)
	fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHashByEmittingWithdrawalReq))
	// get ETH remain
	balEthAfDep4 := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain  : ", balEthAfDep4)
	require.Equal(tradingSuite.T(), balEthAfDep4, big.NewInt(0).Sub(balEthAfDep3, tradingSuite.getGasFeeETHbyTxhash(txHashByEmittingWithdrawalReq)), "balance ETH incorrect")

	balDaiScDepS4 := tradingSuite.getDepositedBalance(
		tradingSuite.DAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] DAI balance on SC after withdraw at step 4 : ", balDaiScDepS4)
	// TODO assert DAI balane on SC
	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(100 * time.Second)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncDAITokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)

	balpDaiAfIssS4, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncDAITokenIDStr,
	)
	fmt.Println("[INFO] pDAI balance after issuing step 4 : ", balpDaiAfIssS4)
	// TODO assert pDai balance issuing
	balPrvAfIssS4, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after issuing step 4: ", balPrvAfIssS4)
	//require.NotEqual(tradingSuite.T(), balPrvAfIssS4, (balPrvAfBurnS2 - tradingSuite.getFeePRVbyTxhashInC(issuuRes["TxID"].(string))), "Balance PRV remain incorrect after issuu step 4")

	// fmt.Println("===================================================")

	// txHashByEmittingWithdrawalReq1 := tradingSuite.requestWithdraw(
	// 	tradingSuite.EtherAddressStr,
	// 	balEthScTradeS3,
	// )

	// balDaiScDepS41 := tradingSuite.getDepositedBalance(
	// 	tradingSuite.EtherAddressStr,
	// 	pubKeyToAddrStr,
	// )
	// fmt.Println("[INFO] ETH balance on SC after withdraw at step 4 : ", balDaiScDepS41)
	// // TODO assert DAI balane on SC
	// _, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq1)
	// require.Equal(tradingSuite.T(), nil, err)
	// fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	// fmt.Println("Waiting 90s for 15 blocks confirmation")
	// time.Sleep(100 * time.Second)

	// _, err = tradingSuite.callIssuingETHReq(
	// 	tradingSuite.IncEtherTokenIDStr,
	// 	ethDepositProof,
	// 	ethBlockHash,
	// 	ethTxIdx,
	// )
	// require.Equal(tradingSuite.T(), nil, err)
	// time.Sleep(120 * time.Second)

	// balpDaiAfIssS41, _ := tradingSuite.getBalanceTokenIncAccount(
	// 	tradingSuite.IncPrivKeyStr,
	// 	tradingSuite.IncEtherTokenIDStr,
	// )
	// fmt.Println("[INFO] pETH balance after issuing step 41 : ", balpDaiAfIssS41)

	// withdrawingPETH := big.NewInt(0).Div(balEthScTradeS3, big.NewInt(1000000000))
	// burningRes, err = tradingSuite.callBurningPToken(
	// 	tradingSuite.IncEtherTokenIDStr,
	// 	withdrawingPETH,
	// 	tradingSuite.ETHOwnerAddrStr,
	// 	"createandsendburningrequest",
	// )
	// require.Equal(tradingSuite.T(), nil, err)
	// burningTxID, found = burningRes["TxID"]
	// require.Equal(tradingSuite.T(), true, found)
	// time.Sleep(120 * time.Second)

	// balpEthAfBurnS51, _ := tradingSuite.getBalanceTokenIncAccount(
	// 	tradingSuite.IncPrivKeyStr,
	// 	tradingSuite.IncEtherTokenIDStr,
	// )
	// fmt.Println("[INFO] pETH balance after burning step 5 : ", balpEthAfBurnS51)

	// tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))

	// balEthAfDep51 := tradingSuite.getBalanceOnETHNet(
	// 	common.HexToAddress(tradingSuite.EtherAddressStr),
	// 	common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	// )
	// fmt.Println("[INFO] ETH balance remain  : ", balEthAfDep51)

	// balETH := tradingSuite.getBalanceOnETHNet(
	// 	common.HexToAddress(tradingSuite.EtherAddressStr),
	// 	common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	// )
	// fmt.Println("ETH balance after trade : ", balETH)

	// fmt.Println("===================================================")

	fmt.Println("------------ step 5: withdrawing pDAI from Incognito to DAI --------------")
	withdrawingPDAI := big.NewInt(0).Div(balDaiScTradeS3, big.NewInt(1000000000))
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

	balpDaiAfBurnS5, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncDAITokenIDStr,
	)
	fmt.Println("[INFO] pDAI balance after burning step 5 : ", balpDaiAfBurnS5)
	// TODO assert pDai balance issuing
	balPrvAfBrunS5, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after burning step 5: ", balPrvAfBrunS5)
	//require.NotEqual(tradingSuite.T(), balPrvAfBrunS5, (balPrvAfIssS4 - tradingSuite.getFeePRVbyTxhashInC(burningRes["TxID"].(string))), "Balance PRV remain incorrect after burn step 5")

	txHash5 := tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))
	fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHash5))
	// get ETH remain
	balEthAfDep5 := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain  : ", balEthAfDep5)
	//require.Equal(tradingSuite.T(),balEthAfDep5,big.NewInt(0).Sub(balEthAfDep4,tradingSuite.getGasFeeETHbyTxhash(txHash5)),"balance ETH incorrect")

	balDai := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.DAIAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	tradingSuite.DAIBalanceAfterStep1 = balDai
	fmt.Println("DAI balance after trade: ", balDai)
	// require.Equal(tradingSuite.T(), withdrawingPDAI.Uint64(), bal.Div(bal, big.NewInt(1000000000)).Uint64())
	balEth := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("ETH balance after trade: ", balEth)
}

func (tradingSuite *ZRXTradingTestSuite) Test2TradeDaiForSaiWith0x() {
return
	fmt.Println("============ TEST 2  TRADE DAI FOR SAI WITH 0X AGGREGATOR ===========")
	fmt.Println("------------ step 0: declaration & initialization --------------")
	depositingDAI := tradingSuite.DAIBalanceAfterStep1
	burningPDAI := big.NewInt(0).Div(depositingDAI, big.NewInt(1000000000))
	tradeAmount := depositingDAI

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("pubKeyToAddrStr: ", pubKeyToAddrStr)

	balEthInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance initialization : ", balEthInit)
	// get info balance initialization
	balPrvInit, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance initialization : ", balPrvInit)

	balpSaiInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncSAITokenIDStr)
	fmt.Println("[INFO] pSAI balance initialization : ", balpSaiInit)

	balpDaiInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncDAITokenIDStr)
	fmt.Println("[INFO] pDAI balance initialization : ", balpDaiInit)

	balSaiInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.SAIAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] SAI balance initialization : ", balSaiInit)

	balDaiInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.DAIAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] DAI balance initialization : ", balDaiInit)

	balSaiScInit := tradingSuite.getDepositedBalance(
		tradingSuite.SAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] SAI balance initialization on SC : ", balSaiScInit)

	balDaiScInit := tradingSuite.getDepositedBalance(
		tradingSuite.DAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] DAI balance initialization on SC : ", balDaiScInit)

	fmt.Println("------------ step 1: porting DAI to pDAI --------------")
	txHash := tradingSuite.depositERC20ToBridge(
		depositingDAI,
		common.HexToAddress(tradingSuite.DAIAddressStr),
		tradingSuite.IncPaymentAddrStr,
	)

	fmt.Println("gas Fee transaction  :", tradingSuite.getGasFeeETHbyTxhash(txHash))
	// get ETH remain after depsit
	balEthAfDep := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain after deposit : ", balEthAfDep)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	// get DAI remain after depsit
	balDaiAfDep := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.DAIAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] DAI balance remain after deposit : ", balDaiAfDep)
	// TODO : assert DAI balance

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(100 * time.Second)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncDAITokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)

	// check PRV and token balance after issuing
	balPrvAfIssS1, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after issuing step 1: ", balPrvAfIssS1)
	// TODO assert PRV balance remain
	balpDaiAfIssS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncDAITokenIDStr)
	fmt.Println("[INFO] pDAI balance after issuing step 1 : ", balpDaiAfIssS1)
	// TODO assert pDai balance issuing

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

	// check PRV and token balance after burning
	balPrvAfBurnS2, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after burning step 2: ", balPrvAfBurnS2)
	// TODO assert PRV balance remain
	balpDaiAfBurnS2, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncDAITokenIDStr)
	fmt.Println("[INFO] pDAI balance after burning step 2 : ", balpDaiAfBurnS2)
	// TODO assert pDai balance issuing

	txHash2 := tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))
	fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHash2))
	// get ETH remain
	balEthAfDep2 := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain  : ", balEthAfDep2)
	// TODO : assert ETH balance
	require.Equal(tradingSuite.T(), balEthAfDep2, big.NewInt(0).Sub(balEthAfDep, tradingSuite.getGasFeeETHbyTxhash(txHash2)), "balance ETH incorrect")

	balDaiScDepS2 := tradingSuite.getDepositedBalance(
		tradingSuite.IncDAITokenIDStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] DAI balance after deposit on SC before trade at step 2: ", balDaiScDepS2)
	// TODO assert DAI balane on SC
	// require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPDAI, big.NewInt(1000000000)), deposited)
	balSaiScS2 := tradingSuite.getDepositedBalance(
		tradingSuite.IncSAITokenIDStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] SAI balance on SC before trade at step 2 : ", balSaiScS2)

	fmt.Println("------------ step 3: execute trade DAI for SAI through 0x aggregator --------------")
	txHash3, protoFee := tradingSuite.executeWith0x(
		tradeAmount,
		"DAI",
		tradingSuite.DAIAddressStr,
		"SAI",
		tradingSuite.SAIAddressStr,
	)
	fmt.Println("protocol fee :", protoFee)
	fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHash3))
	// get ETH remain
	balEthAfDep3 := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain  : ", balEthAfDep3)
	//require.Equal(tradingSuite.T(), balEthAfDep3, big.NewInt(0).Sub(big.NewInt(0).Sub(balEthAfDep2, tradingSuite.getGasFeeETHbyTxhash(txHash3)), protoFee), "balance ETH incorrect")

	balSaiScTradeS3 := tradingSuite.getDepositedBalance(
		tradingSuite.SAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] SAI balance on SC after trade at step 3 : ", balSaiScTradeS3)
	// TODO assert SAI balane on SC
	balDaiScTradeS3 := tradingSuite.getDepositedBalance(
		tradingSuite.DAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] DAI balance on SC after trade at step 3 : ", balDaiScTradeS3)
	// TODO assert DAI balane on SC

	fmt.Println("------------ step 4: withdrawing SAI from SC to pSAI on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.SAIAddressStr,
		balSaiScTradeS3,
	)

	fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHashByEmittingWithdrawalReq))
	// get ETH remain
	balEthAfDep4 := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain  : ", balEthAfDep4)
	require.Equal(tradingSuite.T(), balEthAfDep4, big.NewInt(0).Sub(balEthAfDep3, tradingSuite.getGasFeeETHbyTxhash(txHashByEmittingWithdrawalReq)), "balance ETH incorrect")
	balSaiScDepS4 := tradingSuite.getDepositedBalance(
		tradingSuite.SAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] SAI balance on SC after withdraw at step 4 : ", balSaiScDepS4)
	// TODO assert SAI balane on SC

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(100 * time.Second)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncSAITokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)

	balpSaiAfIssS4, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncSAITokenIDStr,
	)
	fmt.Println("[INFO] pSAI balance after issuing step 4 : ", balpSaiAfIssS4)
	// TODO assert pSai balance issuing
	balPrvAfIssS4, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after issuing step 4: ", balPrvAfIssS4)
	// TODO assert PRV balance remain

	fmt.Println("------------ step 5: withdrawing pSAI from Incognito to SAI --------------")
	withdrawingPSAI := big.NewInt(0).Div(balSaiScTradeS3, big.NewInt(1000000000))
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncSAITokenIDStr,
		withdrawingPSAI,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	balpSaiAfBurnS5, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncSAITokenIDStr,
	)
	fmt.Println("[INFO] pSAI balance after burning step 5 : ", balpSaiAfBurnS5)
	// TODO assert pSAI balance issuing
	balPrvAfBrunS5, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after burning step 5: ", balPrvAfBrunS5)
	// TODO assert PRV balance remain

	txHash5 := tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))
	fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHash5))
	// get ETH remain
	balEthAfDep5 := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain  : ", balEthAfDep5)
	require.Equal(tradingSuite.T(), balEthAfDep5, big.NewInt(0).Sub(balEthAfDep4, tradingSuite.getGasFeeETHbyTxhash(txHash5)), "balance ETH incorrect")

	balSAI := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.SAIAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	tradingSuite.SAIBalanceAfterStep2 = balSAI
	fmt.Println("SAI balance after trade: ", balSAI)
	// require.Equal(tradingSuite.T(), withdrawingPSAI.Uint64(), bal.Uint64())
	// require.Equal(tradingSuite.T(), withdrawingPSAI.Uint64(), bal.Div(bal, big.NewInt(1000000000)).Uint64())

	balDAI := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.DAIAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("DAI balance after trade: ", balDAI)
}

func (tradingSuite *ZRXTradingTestSuite) Test3TradeSaiForEthWith0x() {
return
	fmt.Println("============ TEST TRADE SAI FOR ETH WITH 0X AGGREGATOR ===========")
	fmt.Println("------------ step 0: declaration & initialization --------------")
	depositingSAI := tradingSuite.SAIBalanceAfterStep2
	burningPSAI := big.NewInt(0).Div(depositingSAI, big.NewInt(1000000000))
	tradeAmount := depositingSAI

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("pubKeyToAddrStr: ", pubKeyToAddrStr)

	// get info balance initialization
	balPrvInit, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance initialization : ", balPrvInit)

	balpSAIInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncSAITokenIDStr)
	fmt.Println("[INFO] pSAI balance initialization : ", balpSAIInit)

	balpEthInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance initialization : ", balpEthInit)

	balSaiInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.SAIAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] SAI balance initialization : ", balSaiInit)

	balEthInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance initialization : ", balEthInit)

	balEthScInit := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance initialization on SC : ", balEthScInit)

	balSaiScInit := tradingSuite.getDepositedBalance(
		tradingSuite.SAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] SAI balance initialization on SC : ", balSaiScInit)

	fmt.Println("------------ step 1: porting SAI to pSAI --------------")
	txHash := tradingSuite.depositERC20ToBridge(
		depositingSAI,
		common.HexToAddress(tradingSuite.SAIAddressStr),
		tradingSuite.IncPaymentAddrStr,
	)

	fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHash))
	// get ETH remain after depsit
	balEthAfDep := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain after deposit : ", balEthAfDep)
	//require.Equal(tradingSuite.T(),balEthAfDep,big.NewInt(0).Sub(big.NewInt(0).Sub(balEthInit,tradingSuite.getGasFeeETHbyTxhash(txHash)),tradingSuite.getGasFeeETHbyTxhash(apTxHash)),"balance ETH incorrect")
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)
	// get SAI remain after depsit
	balSaiAfDep := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.SAIAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] SAI balance remain after deposit : ", balSaiAfDep)
	// TODO : assert SAI balance

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(100 * time.Second)

	issuingRes, err := tradingSuite.callIssuingETHReq(
		tradingSuite.IncSAITokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("issuingRes: ", issuingRes)
	time.Sleep(120 * time.Second)
	// check PRV and token balance after issuing
	balPrvAfIssS1, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after issuing step 1: ", balPrvAfIssS1)
	// TODO assert PRV balance remain
	balpSaiAfIssS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncSAITokenIDStr)
	fmt.Println("[INFO] pSAI balance after issuing step 1 : ", balpSaiAfIssS1)
	// TODO assert pSAI balance issuing

	fmt.Println("------------ step 2: burning pSAI to deposit SAI to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningPToken(
		tradingSuite.IncSAITokenIDStr,
		burningPSAI,
		pubKeyToAddrStr[2:],
		"createandsendburningfordeposittoscrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(140 * time.Second)

	// check PRV and token balance after burning
	balPrvAfBurnS2, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after burning step 2: ", balPrvAfBurnS2)
	// TODO assert PRV balance remain
	balpSaiAfBurnS2, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncSAITokenIDStr)
	fmt.Println("[INFO] pSAI balance after burning step 2 : ", balpSaiAfBurnS2)
	// TODO assert pSAI balance issuing

	txHash2 := tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))
	fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHash2))
	// get ETH remain
	balEthAfDep2 := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain  : ", balEthAfDep2)
	// TODO : assert ETH balance
	require.Equal(tradingSuite.T(), balEthAfDep2, big.NewInt(0).Sub(balEthAfDep, tradingSuite.getGasFeeETHbyTxhash(txHash2)), "balance ETH incorrect")

	balSaiScDepS2 := tradingSuite.getDepositedBalance(
		tradingSuite.SAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] SAI balance after deposit on SC at step 2: ", balSaiScDepS2)
	require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPSAI, big.NewInt(1000000000)), balSaiScDepS2)

	balEthScS2 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC at step 2 : ", balEthScS2)

	fmt.Println("------------ step 3: execute trade SAI for ETH through 0x aggregator --------------")
	txHash3, protoFee := tradingSuite.executeWith0x(
		tradeAmount,
		"SAI",
		tradingSuite.SAIAddressStr,
		"ETH",
		tradingSuite.EtherAddressStr,
	)
	fmt.Println("protocol fee :", protoFee)
	fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHash3))
	// get ETH remain
	balEthAfDep3 := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain  : ", balEthAfDep3)
	//require.Equal(tradingSuite.T(), balEthAfDep3, big.NewInt(0).Sub(big.NewInt(0).Sub(balEthAfDep2, tradingSuite.getGasFeeETHbyTxhash(txHash3)), protoFee), "balance ETH incorrect")

	balEthScTradeS3 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC after trade at step 3 : ", balEthScTradeS3)
	// TODO assert ETH balane on SC
	balSaiScTradeS3 := tradingSuite.getDepositedBalance(
		tradingSuite.SAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] SAI balance on SC after trade at step 3 : ", balSaiScTradeS3)
	// TODO assert SAI balane on SC

	fmt.Println("------------ step 4: withdrawing ETH from SC to pETH on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.EtherAddressStr,
		balEthScTradeS3,
	)
	fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHashByEmittingWithdrawalReq))
	// get ETH remain
	balEthAfDep4 := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain  : ", balEthAfDep4)
	require.Equal(tradingSuite.T(), balEthAfDep4, big.NewInt(0).Sub(balEthAfDep3, tradingSuite.getGasFeeETHbyTxhash(txHashByEmittingWithdrawalReq)), "balance ETH incorrect")

	balEthScDepS4 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC after withdraw at step 4 : ", balEthScDepS4)
	// TODO assert ETH balane on SC

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(120 * time.Second)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(140 * time.Second)

	balpEthAfIssS4, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncEtherTokenIDStr,
	)
	fmt.Println("[INFO] pETH balance after issuing step 4 : ", balpEthAfIssS4)
	// TODO assert pETH balance issuing
	balPrvAfIssS4, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after issuing step 4: ", balPrvAfIssS4)
	// TODO assert PRV balance remain

	fmt.Println("------------ step 5: withdrawing pETH from Incognito to ETH --------------")
	withdrawingPETH := big.NewInt(0).Div(balEthScTradeS3, big.NewInt(1000000000))
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncEtherTokenIDStr,
		withdrawingPETH,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	balpEthAfBurnS5, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncEtherTokenIDStr,
	)
	fmt.Println("[INFO] pETH balance after burning step 5 : ", balpEthAfBurnS5)
	// TODO assert pETH balance issuing
	balPrvAfBrunS5, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after burning step 5: ", balPrvAfBrunS5)
	// TODO assert PRV balance remain

	txHash5 := tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))
	fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHash5))
	// get ETH remain
	balEthAfDep5 := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain  : ", balEthAfDep5)
	require.Equal(tradingSuite.T(), balEthAfDep5, big.NewInt(0).Sub(balEthAfDep4, tradingSuite.getGasFeeETHbyTxhash(txHash5)), "balance ETH incorrect")

	balETH := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("ETH balance after trade : ", balETH)
	// require.Equal(tradingSuite.T(), withdrawingPETH.Uint64(), bal.Div(bal, big.NewInt(1000000000)).Uint64())

	balSai := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.SAIAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("SAI balance after trade: ", balSai)
}

func (tradingSuite *ZRXTradingTestSuite) Test4TradeEthForEthWith0x() {
return
	fmt.Println("============ TEST 4 TRADE WITH 0X AGGREGATOR ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
	burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()

	// get info balance initialization
	balPrvInit, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance initialization : ", balPrvInit)

	balpEthInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance initialization : ", balpEthInit)

	balpDaiInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncDAITokenIDStr)
	fmt.Println("[INFO] pDAI balance initialization : ", balpDaiInit)

	balEthInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance initialization : ", balEthInit)

	balDaiInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.DAIAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] DAI balance initialization : ", balDaiInit)

	balEthScInit := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance initialization on SC : ", balEthScInit)

	balDaiScInit := tradingSuite.getDepositedBalance(
		tradingSuite.DAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] DAI balance initialization on SC : ", balDaiScInit)

	fmt.Println("------------ STEP 1: porting ETH to pETH --------------")
	// Deposit to proof
	txHash := tradingSuite.depositETH(
		tradingSuite.DepositingEther,
		tradingSuite.IncPaymentAddrStr,
	)
	// Proof
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)
	// get ETH remain after depsit
	balEthAfDep := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain after deposit : ", balEthAfDep)
	// TODO : assert ETH balance

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(100 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)
	// check PRV and token balance after issuing
	balPrvAfIssS1, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after issuing step 1: ", balPrvAfIssS1)
	// TODO assert PRV balance remain
	balpEthAfIssS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance after issuing step 1 : ", balpEthAfIssS1)
	// TODO assert pETH balance issuing

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

	// check PRV and token balance after burning
	balPrvAfBurnS2, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after burning step 2: ", balPrvAfBurnS2)
	// TODO assert PRV balance remain
	balpEthAfBurnS2, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance after burning step 2 : ", balpEthAfBurnS2)
	// TODO assert pETH balance issuing

	tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))

	balEthScDepS2 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance after deposit on SC at step 2: ", balEthScDepS2)
	// TODO assert ETH balane on SC
	balDaiScS2 := tradingSuite.getDepositedBalance(
		tradingSuite.DAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] DAI balance on SC at step 2 : ", balDaiScS2)

	//require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPETH, big.NewInt(1000000000)), deposited)

	fmt.Println("------------ STEP 3: execute trade through 0x aggregator --------------")
	fmt.Println("------------ STEP 3.1: execute trade ETH for DAI through 0x aggregator --------------")
	tradingSuite.executeWith0x(
		tradeAmount,
		"ETH",
		tradingSuite.EtherAddressStr,
		"DAI",
		tradingSuite.DAIAddressStr,
	)
	balEthScTradeS31 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC after trade at step 3.1 : ", balEthScTradeS31)
	// TODO assert ETH balane on SC
	balDaiScTradeS31 := tradingSuite.getDepositedBalance(
		tradingSuite.DAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] DAI balance on SC after trade at step 3.1 : ", balDaiScTradeS31)
	// TODO assert DAI balane on SC
	fmt.Println("------------ STEP 3.2: execute trade DAI for SAI through 0x aggregator --------------")
	tradeAmount = balDaiScTradeS31
	tradingSuite.executeWith0x(
		tradeAmount,
		"DAI",
		tradingSuite.DAIAddressStr,
		"SAI",
		tradingSuite.SAIAddressStr,
	)
	balDAIScTradeS32 := tradingSuite.getDepositedBalance(
		tradingSuite.DAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] DAI balance on SC after trade at step 3.2 : ", balDAIScTradeS32)
	balSAIScTradeS32 := tradingSuite.getDepositedBalance(
		tradingSuite.SAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] SAI balance on SC after trade at step 3.2 : ", balSAIScTradeS32)

	fmt.Println("------------ STEP 3.3: execute trade SAI for ETH through 0x aggregator --------------")
	tradeAmount = balSAIScTradeS32
	tradingSuite.executeWith0x(
		tradeAmount,
		"SAI",
		tradingSuite.SAIAddressStr,
		"WETH",
		tradingSuite.EtherAddressStr,
	)
	balSAIScTradeS33 := tradingSuite.getDepositedBalance(
		tradingSuite.SAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] SAI balance on SC after trade at step 3.3 : ", balSAIScTradeS33)
	balEthScTradeS33 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC after trade at step 3.3 : ", balEthScTradeS33)

	fmt.Println("------------ step 4: withdrawing ETH from SC to pETH on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.EtherAddressStr,
		balEthScTradeS33,
	)

	balEthScDepS4 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC after withdraw at step 4 : ", balEthScDepS4)
	// TODO assert ETH balane on SC

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(120 * time.Second)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(140 * time.Second)

	balpEthAfIssS4, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncEtherTokenIDStr,
	)
	fmt.Println("[INFO] pETH balance after issuing step 4 : ", balpEthAfIssS4)
	// TODO assert pETH balance issuing
	balPrvAfIssS4, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after issuing step 4: ", balPrvAfIssS4)
	// TODO assert PRV balance remain

	fmt.Println("------------ step 5: withdrawing pETH from Incognito to ETH --------------")
	withdrawingPETH := big.NewInt(0).Div(balEthScTradeS33, big.NewInt(1000000000))
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncEtherTokenIDStr,
		withdrawingPETH,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	balpEthAfBurnS5, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncEtherTokenIDStr,
	)
	fmt.Println("[INFO] pETH balance after burning step 5 : ", balpEthAfBurnS5)
	// TODO assert pETH balance issuing
	balPrvAfBrunS5, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after burning step 5: ", balPrvAfBrunS5)
	// TODO assert PRV balance remain

	tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))

	balETH := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("ETH balance after trade : ", balETH)
	// require.Equal(tradingSuite.T(), withdrawingPETH.Uint64(), bal.Div(bal, big.NewInt(1000000000)).Uint64())

	balSai := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.SAIAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("SAI balance after trade: ", balSai)
}

func (tradingSuite *ZRXTradingTestSuite) Test5TradeEthWith0x() {
return
	fmt.Println("============ TEST 5 TRADE WITH 0X AGGREGATOR ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
	burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()

	// get info balance initialization
	balPrvInit, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance initialization : ", balPrvInit)

	balpEthInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance initialization : ", balpEthInit)

	balpDaiInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncDAITokenIDStr)
	fmt.Println("[INFO] pDAI balance initialization : ", balpDaiInit)

	balEthInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance initialization : ", balEthInit)

	balDaiInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.DAIAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] DAI balance initialization : ", balDaiInit)

	balEthScInit := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance initialization on SC : ", balEthScInit)

	balDaiScInit := tradingSuite.getDepositedBalance(
		tradingSuite.DAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] DAI balance initialization on SC : ", balDaiScInit)

	fmt.Println("------------ STEP 1: porting ETH to pETH --------------")
	// Deposit to proof
	txHash := tradingSuite.depositETH(
		tradingSuite.DepositingEther,
		tradingSuite.IncPaymentAddrStr,
	)
	// Proof
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)
	// get ETH remain after depsit
	balEthAfDep := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain after deposit : ", balEthAfDep)
	// TODO : assert ETH balance

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(100 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)
	// check PRV and token balance after issuing
	balPrvAfIssS1, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after issuing step 1: ", balPrvAfIssS1)
	// TODO assert PRV balance remain
	balpEthAfIssS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance after issuing step 1 : ", balpEthAfIssS1)
	// TODO assert pETH balance issuing

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

	// check PRV and token balance after burning
	balPrvAfBurnS2, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after burning step 2: ", balPrvAfBurnS2)
	// TODO assert PRV balance remain
	balpEthAfBurnS2, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance after burning step 2 : ", balpEthAfBurnS2)
	// TODO assert pETH balance issuing

	tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))

	balEthScDepS2 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance after deposit on SC at step 2: ", balEthScDepS2)
	// TODO assert ETH balane on SC
	balDaiScS2 := tradingSuite.getDepositedBalance(
		tradingSuite.DAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] DAI balance on SC at step 2 : ", balDaiScS2)

	//require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPETH, big.NewInt(1000000000)), deposited)

	// get ETH remain
	balEthAfDep2 := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain  : ", balEthAfDep2)

	fmt.Println("------------ STEP 3: execute trade through 0x aggregator --------------")
	fmt.Println("------------ STEP 3.1: execute trade ETH for DAI through 0x aggregator --------------")
	txHash31, protoFee31 := tradingSuite.executeWith0x(
		tradeAmount,
		"ETH",
		tradingSuite.EtherAddressStr,
		"DAI",
		tradingSuite.DAIAddressStr,
	)
	fmt.Println("protocol fee :", protoFee31)
	fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHash31))
	// get ETH remain
	balEthAfDep31 := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain s3.1  : ", balEthAfDep31)
	//require.Equal(tradingSuite.T(), balEthAfDep31, big.NewInt(0).Sub(big.NewInt(0).Sub(balEthAfDep2, tradingSuite.getGasFeeETHbyTxhash(txHash31)), protoFee31), "balance ETH incorrect")

	balEthScTradeS31 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC after trade at step 3.1 : ", balEthScTradeS31)
	// TODO assert ETH balane on SC
	balDaiScTradeS31 := tradingSuite.getDepositedBalance(
		tradingSuite.DAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] DAI balance on SC after trade at step 3.1 : ", balDaiScTradeS31)
	// TODO assert DAI balane on SC
	fmt.Println("------------ STEP 3.2: execute trade DAI for SAI through 0x aggregator --------------")
	tradeAmount = balDaiScTradeS31
	txHash32, protoFee32 := tradingSuite.executeWith0x(
		tradeAmount,
		"DAI",
		tradingSuite.DAIAddressStr,
		"SAI",
		tradingSuite.SAIAddressStr,
	)
	fmt.Println("protocol fee :", protoFee32)
	fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHash32))
	// get ETH remain
	balEthAfDep32 := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain s3.2  : ", balEthAfDep32)
	//require.Equal(tradingSuite.T(), balEthAfDep32, big.NewInt(0).Sub(big.NewInt(0).Sub(balEthAfDep31, tradingSuite.getGasFeeETHbyTxhash(txHash32)), protoFee32), "balance ETH incorrect")

	balDAIScTradeS32 := tradingSuite.getDepositedBalance(
		tradingSuite.DAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] DAI balance on SC after trade at step 3.2 : ", balDAIScTradeS32)
	balSAIScTradeS32 := tradingSuite.getDepositedBalance(
		tradingSuite.SAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] SAI balance on SC after trade at step 3.2 : ", balSAIScTradeS32)

	fmt.Println("------------ STEP 3.3: execute trade SAI for ETH through 0x aggregator --------------")
	tradeAmount = balSAIScTradeS32
	txHash33, protoFee33 := tradingSuite.executeWith0x(
		tradeAmount,
		"SAI",
		tradingSuite.SAIAddressStr,
		"ETH",
		tradingSuite.EtherAddressStr,
	)
	fmt.Println("protocol fee :", protoFee33)
	fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHash33))
	// get ETH remain
	balEthAfDep33 := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain s3.3  : ", balEthAfDep33)
	//require.Equal(tradingSuite.T(), balEthAfDep33, big.NewInt(0).Sub(big.NewInt(0).Sub(balEthAfDep32, tradingSuite.getGasFeeETHbyTxhash(txHash33)), protoFee33), "balance ETH incorrect")

	balSAIScTradeS33 := tradingSuite.getDepositedBalance(
		tradingSuite.SAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] SAI balance on SC after trade at step 3.3 : ", balSAIScTradeS33)
	balEthScTradeS33 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC after trade at step 3.3 : ", balEthScTradeS33)

	fmt.Println("------------ STEP 3.4: execute trade ETH for KNC through 0x aggregator --------------")
	tradeAmount = balEthScTradeS33
	txHash34, protoFee34 := tradingSuite.executeWith0x(
		tradeAmount,
		"ETH",
		tradingSuite.EtherAddressStr,
		"KNC",
		tradingSuite.KNCAddressStr,
	)
	fmt.Println("protocol fee :", protoFee34)
	fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHash34))
	// get ETH remain
	balEthAfDep34 := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain s3.4  : ", balEthAfDep34)
	//require.Equal(tradingSuite.T(),balEthAfDep34,big.NewInt(0).Sub(big.NewInt(0).Sub(balEthAfDep33,tradingSuite.getGasFeeETHbyTxhash(txHash34)),protoFee34),"balance ETH incorrect")

	balEthScTradeS34 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC after trade at step 3.4 : ", balEthScTradeS34)
	balKNCScTradeS33 := tradingSuite.getDepositedBalance(
		tradingSuite.KNCAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] KNC balance on SC after trade at step 3.4 : ", balKNCScTradeS33)

	fmt.Println("------------ STEP 3.5: execute trade KNC for BAT through 0x aggregator --------------")
	tradeAmount = balKNCScTradeS33
	txHash35, protoFee35 := tradingSuite.executeWith0x(
		tradeAmount,
		"KNC",
		tradingSuite.KNCAddressStr,
		"BAT",
		tradingSuite.BATAddressStr,
	)
	fmt.Println("protocol fee :", protoFee35)
	fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHash35))
	// get ETH remain
	balEthAfDep35 := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain s3.5  : ", balEthAfDep35)
	//require.Equal(tradingSuite.T(),balEthAfDep35,big.NewInt(0).Sub(big.NewInt(0).Sub(balEthAfDep34,tradingSuite.getGasFeeETHbyTxhash(txHash35)),protoFee35),"balance ETH incorrect")

	balKNCScTradeS35 := tradingSuite.getDepositedBalance(
		tradingSuite.KNCAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] KNC balance on SC after trade at step 3.5 : ", balKNCScTradeS35)
	balBATScTradeS35 := tradingSuite.getDepositedBalance(
		tradingSuite.BATAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] BAT balance on SC after trade at step 3.5 : ", balBATScTradeS35)

	fmt.Println("------------ step 4: withdrawing BAT from SC to pBAT on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.BATAddressStr,
		balBATScTradeS35,
	)

	balBATScDepS4 := tradingSuite.getDepositedBalance(
		tradingSuite.BATAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] BAT balance on SC after withdraw at step 4 : ", balBATScDepS4)
	// TODO assert KNC balane on SC

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(120 * time.Second)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncBATTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(140 * time.Second)

	balpBATAfIssS4, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncBATTokenIDStr,
	)
	fmt.Println("[INFO] pBAT balance after issuing step 4 : ", balpBATAfIssS4)
	// TODO assert pBAT balance issuing
	balPrvAfIssS4, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after issuing step 4: ", balPrvAfIssS4)
	// TODO assert PRV balance remain

	fmt.Println("------------ step 5: withdrawing pBAT from Incognito to BAT --------------")
	withdrawingPETH := big.NewInt(0).Div(balBATScTradeS35, big.NewInt(1000000000))
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncBATTokenIDStr,
		withdrawingPETH,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	balpBATAfBurnS5, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncBATTokenIDStr,
	)
	fmt.Println("[INFO] pBAT balance after burning step 5 : ", balpBATAfBurnS5)
	// TODO assert pKNC balance issuing
	balPrvAfBrunS5, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after burning step 5: ", balPrvAfBrunS5)
	// TODO assert PRV balance remain

	tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))

	balBAT := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.BATAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("BAT balance after trade : ", balBAT)
	// require.Equal(tradingSuite.T(), withdrawingPETH.Uint64(), bal.Div(bal, big.NewInt(1000000000)).Uint64())

	balETH := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("ETH balance after trade: ", balETH)
}

func (tradingSuite *ZRXTradingTestSuite) Test6VerifyIssuuTokenWait15Block() {
return
	fmt.Println("============ TEST 6 VERIFY ISSUU TOKEN WAIT 15 blocks ===========")

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()

	// get info balance initialization
	balPrvInit, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance initialization : ", balPrvInit)

	balpEthInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance initialization : ", balpEthInit)

	balEthInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance initialization : ", balEthInit)

	balEthScInit := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance initialization on SC : ", balEthScInit)

	// Deposit to proof
	txHash := tradingSuite.depositETH(
		tradingSuite.DepositingEther,
		tradingSuite.IncPaymentAddrStr,
	)

	// // get ETH remain after depsit
	// balEthAfDep := tradingSuite.getBalanceOnETHNet(
	// 	common.HexToAddress(tradingSuite.EtherAddressStr),
	// 	common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	// )
	// fmt.Println("[INFO] ETH balance remain after deposit 1 : ", balEthAfDep)
	// require.NotEqual(tradingSuite.T(), balEthAfDep, balEthInit, "balane ETH incorect after deposit")

	// Proof
	_, ethBlockHash, ethTxIdx, ethDepositProof, err1 := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err1)
	fmt.Println("depositProof 1 ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	// fmt.Println("Issuing ignore waiting for 15 blocks confirmation")
	// _, err2 := tradingSuite.callIssuingETHReq(
	// 	tradingSuite.IncEtherTokenIDStr,
	// 	ethDepositProof,
	// 	ethBlockHash,
	// 	ethTxIdx,
	// )
	// //fmt.Println(err2)
	// require.Equal(tradingSuite.T(), nil, err2, "issuu must failed")
	// time.Sleep(120 * time.Second)
	// // check PRV after issuing
	// balPrvAfIssS1, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	// fmt.Println("[INFO] PRV balance after issuing (don't wait): ", balPrvAfIssS1)
	// require.NotEqual(tradingSuite.T(), balPrvAfIssS1, balPrvInit, "Balance PRV remain incorrect after issuu failed")
	// // check token balance after issuing
	// balpEthAfIssS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	// fmt.Println("[INFO] pETH balance after issuing (don't wait) : ", balpEthAfIssS1)
	// require.Equal(tradingSuite.T(), balpEthAfIssS1, balpEthInit, "must issuing failed")

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(100 * time.Second)
	_, err := tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)

	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)
	// check PRV after issuing
	// balPrvAfIssS1, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	// fmt.Println("[INFO] PRV balance after issuing (waiting 15 blocks confirmation): ", balPrvAfIssS1)
	// require.NotEqual(tradingSuite.T(), balPrvAfIssS1, (balPrvInit - tradingSuite.getFeePRVbyTxhashInC(issuuRes["TxID"].(string))), "Balance PRV remain incorrect after issuu success")
	// check token balance after issuing
	balpEthAfIssS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance after issuing (waiting 15 block confirmation) : ", balpEthAfIssS1)
	require.NotEqual(tradingSuite.T(), balpEthAfIssS1, balpEthInit, "must issuing success")
	// TODO assert pETH balance issuing

}

func (tradingSuite *ZRXTradingTestSuite) Test7VerifyIssuuTokenWithdrawWait15Block() {
return
	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()

	fmt.Println("============ TEST 7 VERIFY ISSUU TOKEN WITHDRAW WAIT 15 BLOCKS ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")

	// get info balance initialization
	balPrvInit, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance initialization : ", balPrvInit)

	balpEthInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance initialization : ", balpEthInit)

	balEthInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance initialization : ", balEthInit)

	balEthScInit := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance initialization on SC : ", balEthScInit)

	// Deposit to proof
	txHash := tradingSuite.depositETH(
		tradingSuite.DepositingEther,
		tradingSuite.IncPaymentAddrStr,
	)

	// // get ETH remain after depsit
	// balEthAfDep := tradingSuite.getBalanceOnETHNet(
	// 	common.HexToAddress(tradingSuite.EtherAddressStr),
	// 	common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	// )
	// fmt.Println("[INFO] ETH balance remain after deposit 1 : ", balEthAfDep)
	// require.NotEqual(tradingSuite.T(), balEthAfDep, balEthInit, "balane ETH incorect after deposit")

	// Proof
	_, ethBlockHash, ethTxIdx, ethDepositProof, err1 := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err1)
	fmt.Println("depositProof 1 ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	// fmt.Println("Issuing ignore waiting for 15 blocks confirmation")
	// _, err2 := tradingSuite.callIssuingETHReq(
	// 	tradingSuite.IncEtherTokenIDStr,
	// 	ethDepositProof,
	// 	ethBlockHash,
	// 	ethTxIdx,
	// )
	// //fmt.Println(err2)
	// require.Equal(tradingSuite.T(), nil, err2, "issuu must failed")
	// time.Sleep(120 * time.Second)
	// // check PRV after issuing
	// balPrvAfIssS1, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	// fmt.Println("[INFO] PRV balance after issuing (don't wait): ", balPrvAfIssS1)
	// require.NotEqual(tradingSuite.T(), balPrvAfIssS1, balPrvInit, "Balance PRV remain incorrect after issuu failed")
	// // check token balance after issuing
	// balpEthAfIssS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	// fmt.Println("[INFO] pETH balance after issuing (don't wait) : ", balpEthAfIssS1)
	// require.Equal(tradingSuite.T(), balpEthAfIssS1, balpEthInit, "must issuing failed")

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(100 * time.Second)
	_, err := tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)

	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)

	balpEthAfIssS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance after issuing (waiting 15 block confirmation) : ", balpEthAfIssS1)
	require.NotEqual(tradingSuite.T(), balpEthAfIssS1, balpEthInit, "must issuing success")
	// TODO assert pETH balance issuing

	fmt.Println("[INFO] ETH balance initialization on SC : ", balEthScInit)
	fmt.Println("------------  withdrawing ETH from SC to pETH on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.EtherAddressStr,
		big.NewInt(0).Div(balEthScInit, big.NewInt(int64(4))),
	)

	balEthScDepS4 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC after withdraw  : ", balEthScDepS4)
	// TODO assert ETH balane on SC

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(100 * time.Second)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(140 * time.Second)

	balpEthAfIssS4, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncEtherTokenIDStr,
	)
	fmt.Println("[INFO] pETH balance after issuing step 4 : ", balpEthAfIssS4)

}

func (tradingSuite *ZRXTradingTestSuite) Test7Verify1DepositButIssuu3Times() {
return
	fmt.Println("============ TEST 7 VERIFY DEPOSTI 1 PROOF BUT ISSUU 3 TIMES ===========")

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()

	// get info balance initialization
	balPrvInit, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance initialization : ", balPrvInit)

	balpEthInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance initialization : ", balpEthInit)

	balEthInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance initialization : ", balEthInit)

	balEthScInit := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance initialization on SC : ", balEthScInit)

	// Deposit to proof
	txHash := tradingSuite.depositETH(
		tradingSuite.DepositingEther,
		tradingSuite.IncPaymentAddrStr,
	)

	// get ETH remain after depsit
	balEthAfDep := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain after deposit 1 : ", balEthAfDep)
	require.NotEqual(tradingSuite.T(), balEthAfDep, balEthInit, "balane ETH incorect after deposit")

	// Proof
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof 1 ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(100 * time.Second)
	//"----The first issuu ------"
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err, "Issuu must success")
	//"----The second issuu ------"
	_, err2 := tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	fmt.Println(err2)
	require.NotEqual(tradingSuite.T(), nil, err2, "Issuu must failed")

	time.Sleep(120 * time.Second)
	// check PRV after issuing
	balPrvAfIssS1, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after issuing (waiting 15 blocks confirmation): ", balPrvAfIssS1)
	//require.NotEqual(tradingSuite.T(), balPrvAfIssS1, (balPrvInit - tradingSuite.getFeePRVbyTxhashInC(issuuRes["TxID"].(string))), "Balance PRV remain incorrect after issuu success")
	// check token balance after issuing
	balpEthAfIssS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance after issuing (waiting 15 block confirmation) : ", balpEthAfIssS1)
	require.NotEqual(tradingSuite.T(), balpEthAfIssS1, balpEthInit, "must issuing success")
	// TODO assert pETH balance issuing
	//"----The third issuu ------"
	_, err3 := tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.NotEqual(tradingSuite.T(), nil, err3, "Issuu must failed")

}

func (tradingSuite *ZRXTradingTestSuite) Test8Verify1DepositManyCoin1Time() {
return
	fmt.Println("============ TEST 8 VERIFY DEPOSTI MANY COINS AT SAME TIME ===========")

	// get info balance initialization
	balEthInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance initialization : ", balEthInit)

	balBATInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.BATAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] BAT balance initialization : ", balBATInit)

	balpEthInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance initialization : ", balpEthInit)

	balpBATInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncBATTokenIDStr)
	fmt.Println("[INFO] pBATbalance initialization : ", balpBATInit)

	// Deposit to proof
	txHash := tradingSuite.depositETH(
		tradingSuite.DepositingEther,
		tradingSuite.IncPaymentAddrStr,
	)

	txHashBAT := tradingSuite.depositERC20ToBridge(
		tradingSuite.DAIBalanceAfterStep1,
		common.HexToAddress(tradingSuite.BATAddressStr),
		tradingSuite.IncPaymentAddrStr,
	)

	balEth := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain : ", balEth)

	balBAT := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.BATAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] BAT balance remain : ", balBAT)

	// Proof
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof 1 ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	_, ethBlockHash2, ethTxIdx2, ethDepositProof2, err := getETHDepositProof(tradingSuite.ETHHost, txHashBAT)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof 2 ---- : ", ethBlockHash2, ethTxIdx2, ethDepositProof2)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(100 * time.Second)
	//"----The first issuu ------"
	_, err1 := tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err1, "Issuu must success")
	//"----The second issuu ------"
	_, err2 := tradingSuite.callIssuingETHReq(
		tradingSuite.BATAddressStr,
		ethDepositProof2,
		ethBlockHash2,
		ethTxIdx2,
	)
	require.Equal(tradingSuite.T(), nil, err2, "Issuu success")

	time.Sleep(120 * time.Second)

	// check token balance after issuing
	balpEthAfIssS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance after issuing (waiting 15 block confirmation) : ", balpEthAfIssS1)
	require.NotEqual(tradingSuite.T(), balpEthAfIssS1, balpEthInit, "must issuing success")

	balpBATAfIssS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncBATTokenIDStr)
	fmt.Println("[INFO] pETH balance after issuing (waiting 15 block confirmation) : ", balpBATAfIssS1)
	require.NotEqual(tradingSuite.T(), balpEthAfIssS1, balpEthInit, "must issuing success")

}

func (tradingSuite *ZRXTradingTestSuite) Test9DuplicateActionSumitProof() {
return
	fmt.Println("============ TEST 1 TRADE ETHER FOR DAI WITH 0X AGGREGATOR ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
	burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()

	fmt.Println("------------ STEP 1: porting ETH to pETH --------------")
	fmt.Println("amount ETH deposit : ", (big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))))
	// Deposit to proof
	txHash := tradingSuite.depositETH(
		tradingSuite.DepositingEther,
		tradingSuite.IncPaymentAddrStr,
	)

	balEthScDepS1 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance after deposit on SC at step 1: ", balEthScDepS1)

	// Proof
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(100 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)

	_, err2 := tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.NotEqual(tradingSuite.T(), nil, err2)

	time.Sleep(120 * time.Second)

	fmt.Println("------------ STEP 2: burning pETH to deposit ETH to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes3, err3 := tradingSuite.callBurningPToken(
		tradingSuite.IncEtherTokenIDStr,
		burningPETH,
		pubKeyToAddrStr[2:],
		"createandsendburningfordeposittoscrequest",
	)
	require.Equal(tradingSuite.T(), nil, err3)

	_, err4 := tradingSuite.callBurningPToken(
		tradingSuite.IncEtherTokenIDStr,
		burningPETH,
		pubKeyToAddrStr[2:],
		"createandsendburningfordeposittoscrequest",
	)
	require.NotEqual(tradingSuite.T(), nil, err4)

	burningTxID, found := burningRes3["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))

	balEthScDepS2 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance after deposit on SC at step 2: ", balEthScDepS2)

	tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))

	fmt.Println("------------ STEP 3: withdrawing DAI from SC to pDAI on Incognito --------------")
	txHashByEmittingWithdrawalReq1 := tradingSuite.requestWithdraw(
		tradingSuite.EtherAddressStr,
		balEthScDepS2,
	)

	tradingSuite.requestWithdraw(
		tradingSuite.EtherAddressStr,
		balEthScDepS2,
	)

	balEthScDepS3 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance after deposit on SC : ", balEthScDepS3)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq1)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(100 * time.Second)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)

	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)

	balpDaiAfIssS41, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncEtherTokenIDStr,
	)
	fmt.Println("[INFO] pETH balance after issuing  : ", balpDaiAfIssS41)

	withdrawingPETH := big.NewInt(0).Div(balEthScDepS2, big.NewInt(1000000000))

	burningRes, err := tradingSuite.callBurningPToken(
		tradingSuite.IncEtherTokenIDStr,
		withdrawingPETH,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	balpEthAfBurnS51, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncEtherTokenIDStr,
	)
	fmt.Println("[INFO] pETH balance after burning  : ", balpEthAfBurnS51)

	tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))

	tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))

	balEthAfDep51 := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain  : ", balEthAfDep51)

}

func (tradingSuite *ZRXTradingTestSuite) Test10verifyburningAmount() {
return
	fmt.Println("============ TEST 9 TRADE ETHER FOR DAI WITH 0X AGGREGATOR ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
	burningPETH1 := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))
	//burningPETH := big.NewInt(0).Add(big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000)), big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000)))
	IncPaymentAddrStr2 := "12S6R8HfTyL74bggg47LX88RSvBPaMPBMEMoo6yx9WQ4EgLiYERXXcE2Mv2HrCsFuKhBsTfrYMeH82Bus5MHQGt3xHwoxX4v2qM5jRE"
	IncPriAddrStr2 := "112t8rnX3VTd3MTWMpfbYP8HGY4ToAaLjrmUYzfjJBrAcb8iPLkNqvVDXWrLNiFV5yb2NBpR3FDZj3VW8GcLUwRdQ61hPMWP3YrREZAZ1UbH"

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()

	// get info balance initialization
	balPrvInit, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance initialization : ", balPrvInit)

	balpEthInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance initialization add1 : ", balpEthInit)

	balpEthInit2, _ := tradingSuite.getBalanceTokenIncAccount(IncPriAddrStr2, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance initialization add2 : ", balpEthInit2)

	balEthInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance initialization : ", balEthInit)

	balEthScInit := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance initialization on SC : ", balEthScInit)

	fmt.Println("------------ STEP 1: porting ETH to pETH --------------")
	fmt.Println("amount ETH deposit : ", (big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))))
	// Deposit to proof
	txHash := tradingSuite.depositETH(
		tradingSuite.DepositingEther,
		tradingSuite.IncPaymentAddrStr,
	)

	// fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHash))
	// // get ETH remain after depsit
	// balEthAfDep := tradingSuite.getBalanceOnETHNet(
	// 	common.HexToAddress(tradingSuite.EtherAddressStr),
	// 	common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	// )
	// fmt.Println("[INFO] ETH balance remain after deposit : ", balEthAfDep)

	// require.Equal(tradingSuite.T(), balEthAfDep, big.NewInt(0).Sub(big.NewInt(0).Sub(balEthInit, big.NewInt(int64(tradingSuite.DepositingEther*params.Ether))), tradingSuite.getGasFeeETHbyTxhash(txHash)), "balance ETH incorrect")
	txHash2 := tradingSuite.depositETH(
		tradingSuite.DepositingEther,
		IncPaymentAddrStr2,
	)

	// Proof
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	_, ethBlockHash2, ethTxIdx2, ethDepositProof2, err2 := getETHDepositProof(tradingSuite.ETHHost, txHash2)
	require.Equal(tradingSuite.T(), nil, err2)
	fmt.Println("depositProof ---- : ", ethBlockHash2, ethTxIdx2, ethDepositProof2)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(100 * time.Second)

	_, err3 := tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err3)
	time.Sleep(100 * time.Second)

	_, err4 := tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof2,
		ethBlockHash2,
		ethTxIdx2,
	)
	require.Equal(tradingSuite.T(), nil, err4)

	time.Sleep(120 * time.Second)

	balpEthAfIssS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance after issuing add 1 : ", balpEthAfIssS1)
	//require.Equal(tradingSuite.T(), big.NewInt(int64(balpEthAfIssS1-balpEthInit)), big.NewInt(0).Div(big.NewInt(int64(tradingSuite.DepositingEther*params.Ether)), big.NewInt(1000000000)), " balnce pToken issuu incorrect")
	balpEthAfIssS2, _ := tradingSuite.getBalanceTokenIncAccount(IncPriAddrStr2, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance after issuing add 2 : ", balpEthAfIssS2)

	fmt.Println("------------ STEP 2: burning pETH to deposit ETH to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	_, err = tradingSuite.callBurningPToken(
		tradingSuite.IncEtherTokenIDStr,
		big.NewInt(0).Add(big.NewInt(int64(balpEthAfIssS1)), burningPETH1),
		pubKeyToAddrStr[2:],
		"createandsendburningfordeposittoscrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	//burningTxID, found := burningRes["TxID"]
	//require.Equal(tradingSuite.T(), true, found)
	time.Sleep(20 * time.Second)
	balpEthAfIssS12, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance after burn add 1 : ", balpEthAfIssS12)
	//require.Equal(tradingSuite.T(), big.NewInt(int64(balpEthAfIssS1-balpEthInit)), big.NewInt(0).Div(big.NewInt(int64(tradingSuite.DepositingEther*params.Ether)), big.NewInt(1000000000)), " balnce pToken issuu incorrect")
	balpEthAfIssS22, _ := tradingSuite.getBalanceTokenIncAccount(IncPriAddrStr2, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance after burn add 2 : ", balpEthAfIssS22)
}

func (tradingSuite *ZRXTradingTestSuite) Test12DepositAndWithdrwaEther() {
return
	fmt.Println("============ TEST 12 DEPOSIT AND WITHDRAW ETHER ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
	burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()

	// get info balance initialization
	balPrvInit, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance initialization : ", balPrvInit)

	balpEthInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance initialization : ", balpEthInit)

	balEthInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance initialization : ", balEthInit)

	balEthScInit := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance initialization on SC : ", balEthScInit)

	fmt.Println("------------ STEP 1: porting ETH to pETH --------------")

	fmt.Println("amount ETH deposit : ", (big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))))

	// Deposit to proof
	txHash := tradingSuite.depositETH(
		tradingSuite.DepositingEther,
		tradingSuite.IncPaymentAddrStr,
	)

	fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHash))
	// get ETH remain after depsit
	balEthAfDep := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain after deposit : ", balEthAfDep)

	//	require.Equal(tradingSuite.T(), balEthAfDep, big.NewInt(0).Sub(big.NewInt(0).Sub(balEthInit, big.NewInt(int64(tradingSuite.DepositingEther*params.Ether))), tradingSuite.getGasFeeETHbyTxhash(txHash)), "balance ETH incorrect")

	// Proof
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(100 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)

	time.Sleep(120 * time.Second)
	// check PRV and token balance after issuing
	balPrvAfIssS1, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after issuing step 1: ", balPrvAfIssS1)
	//require.NotEqual(tradingSuite.T(), balPrvAfIssS1, (balPrvInit - tradingSuite.getFeePRVbyTxhashInC(issuuRes["TxID"].(string))), "Balance PRV remain incorrect after issuu step 1")

	balpEthAfIssS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance after issuing step 1 : ", balpEthAfIssS1)
	//require.Equal(tradingSuite.T(), big.NewInt(int64(balpEthAfIssS1-balpEthInit)), big.NewInt(0).Div(big.NewInt(int64(tradingSuite.DepositingEther*params.Ether)), big.NewInt(1000000000)), " balnce pToken issuu incorrect")

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

	// check PRV and token balance after burning
	balPrvAfBurnS2, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after burning step 2: ", balPrvAfBurnS2)
	//require.NotEqual(tradingSuite.T(), balPrvAfBurnS2, (balPrvAfIssS1 - tradingSuite.getFeePRVbyTxhashInC(burningRes["TxID"].(string))), "Balance PRV remain incorrect after burn step 2")

	balpEthAfBurnS2, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance after burning step 2 : ", balpEthAfBurnS2)
	// TODO assert pETH balance issuing

	txHash2 := tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))
	fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHash2))
	// get ETH remain
	balEthAfDep2 := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain  : ", balEthAfDep2)
	// TODO : assert ETH balance
	//require.Equal(tradingSuite.T(), balEthAfDep2, big.NewInt(0).Sub(balEthAfDep, tradingSuite.getGasFeeETHbyTxhash(txHash2)), "balance ETH incorrect")

	balEthScDepS2 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance after deposit on SC at step 2: ", balEthScDepS2)
	// TODO assert ETH balane on SC

	fmt.Println("------------ STEP 3: withdraw ETH to deposit pETH to Incognito  --------------")

	txHashByEmittingWithdrawalReq1 := tradingSuite.requestWithdraw(
		tradingSuite.EtherAddressStr,
		balEthScDepS2,
	)

	balDaiScDepS41 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC after withdraw at step 3 : ", balDaiScDepS41)
	// TODO assert DAI balane on SC

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq1)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(100 * time.Second)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)

	balpDaiAfIssS41, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncEtherTokenIDStr,
	)
	fmt.Println("[INFO] pETH balance after issuing step 3 : ", balpDaiAfIssS41)

	fmt.Println("------------ STEP 4: withdraw pETH to deposit ETH   --------------")

	withdrawingPETH := big.NewInt(0).Div(balEthScDepS2, big.NewInt(1000000000))

	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncEtherTokenIDStr,
		withdrawingPETH,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	balpEthAfBurnS51, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncEtherTokenIDStr,
	)
	fmt.Println("[INFO] pETH balance after burning step 4 : ", balpEthAfBurnS51)

	tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))

	balEthAfDep51 := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance after withdraw  : ", balEthAfDep51)

}

func (tradingSuite *ZRXTradingTestSuite) Test13DepositAndWithdrwaERC20token() {
return
	fmt.Println("============ TEST 13 DEPOSIT AND WITHDRAW ERC20 TOKEN (BAT) ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()


	// get info balance initialization
	balPrvInit, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance initialization : ", balPrvInit)

	balEthInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance initialization : ", balEthInit)

	balpBATInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncBATTokenIDStr)
	fmt.Println("[INFO] pBATbalance initialization : ", balpBATInit)

	balBATInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.BATAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] BAT balance initialization : ", balBATInit)

	balEthScInit := tradingSuite.getDepositedBalance(
		tradingSuite.BATAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] BAT balance initialization on SC : ", balEthScInit)

	fmt.Println("------------ STEP 1: porting BAT to pBAT --------------")

	fmt.Println("amount BAT deposit : ", (big.NewInt(int64(0.5 * params.Ether))))
	deposit := big.NewInt(int64(0.1 * params.Ether))
	burningPETH := big.NewInt(0).Div(deposit, big.NewInt(1000000000))
	// Deposit to proof
	txHash := tradingSuite.depositERC20ToBridge(
		deposit,
		common.HexToAddress(tradingSuite.BATAddressStr),
		tradingSuite.IncPaymentAddrStr,
	)

	balBATAfDep := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.BATAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] BAT balance remain after deposit : ", balBATAfDep)

	fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHash))
	// get ETH remain after depsit
	balEthAfDep := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain after deposit : ", balEthAfDep)

	//require.Equal(tradingSuite.T(), balEthAfDep, big.NewInt(0).Sub(big.NewInt(0).Sub(balEthInit, big.NewInt(int64(tradingSuite.DepositingEther*params.Ether))), tradingSuite.getGasFeeETHbyTxhash(txHash)), "balance ETH incorrect")

	// Proof
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(100 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncBATTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)

	time.Sleep(120 * time.Second)

	balpBATAfIssS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncBATTokenIDStr)
	fmt.Println("[INFO] pBAT balance after issuing step 1 : ", balpBATAfIssS1)
	// check PRV and token balance after issuing
	balPrvAfIssS1, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after issuing step 1: ", balPrvAfIssS1)

	fmt.Println("------------ STEP 2: burning pBAT to deposit BAT to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningPToken(
		tradingSuite.IncBATTokenIDStr,
		burningPETH,
		pubKeyToAddrStr[2:],
		"createandsendburningfordeposittoscrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	// check PRV and token balance after burning
	balPrvAfBurnS2, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after burning step 2: ", balPrvAfBurnS2)

	balpBATAfBurnS2, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncBATTokenIDStr)
	fmt.Println("[INFO] pBAT balance after burning step 2 : ", balpBATAfBurnS2)
	// TODO assert pETH balance issuing

	txHash2 := tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))

	fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHash2))
	// get ETH remain
	balEthAfDep2 := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain  : ", balEthAfDep2)
	// TODO : assert ETH balance
	//require.Equal(tradingSuite.T(), balEthAfDep2, big.NewInt(0).Sub(balEthAfDep, tradingSuite.getGasFeeETHbyTxhash(txHash2)), "balance ETH incorrect")

	balBATScDepS2 := tradingSuite.getDepositedBalance(
		tradingSuite.BATAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] BAT balance after deposit on SC at step 2: ", balBATScDepS2)
	// TODO assert ETH balane on SC

	fmt.Println("------------ STEP 3: withdraw BAT to deposit pBAT to Incognito  --------------")

	txHashByEmittingWithdrawalReq1 := tradingSuite.requestWithdraw(
		tradingSuite.BATAddressStr,
		balBATScDepS2,
	)

	balBATScDepS41 := tradingSuite.getDepositedBalance(
		tradingSuite.BATAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] BAT balance on SC after withdraw at step 3 : ", balBATScDepS41)
	// TODO assert BAT balane on SC

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq1)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(100 * time.Second)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncBATTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(100 * time.Second)

	balpBATAfIssS41, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncBATTokenIDStr,
	)
	fmt.Println("[INFO] pBAT balance after issuing step 3 : ", balpBATAfIssS41)

	fmt.Println("------------ STEP 4: withdraw pBAT to deposit BAT   --------------")

	withdrawingPBAT := big.NewInt(0).Div(balBATScDepS2, big.NewInt(1000000000))

	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncBATTokenIDStr,
		withdrawingPBAT,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	balpBATAfBurnS51, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncBATTokenIDStr,
	)
	fmt.Println("[INFO] pBAT balance after burning step 4 : ", balpBATAfBurnS51)

	tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))

	balBATAfDep51 := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.BATAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] BAT balance after withdraw  : ", balBATAfDep51)

}

func (tradingSuite *ZRXTradingTestSuite) Test999GetBalance() {
return
	fmt.Println("============ TEST 14 GET BALANCE (BAT) ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()

	// get info balance initialization
	balPrvInit, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance initialization : ", balPrvInit)

	balEthInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance initialization : ", balEthInit)

	balEthScInit := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance initialization on SC : ", balEthScInit)

	balpETHInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance initialization : ", balpETHInit)

	balBATInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.BATAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] BAT balance initialization : ", balBATInit)

	balBATScInit := tradingSuite.getDepositedBalance(
		tradingSuite.BATAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] BAT balance initialization on SC : ", balBATScInit)

	balpBATInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncBATTokenIDStr)
	fmt.Println("[INFO] pBAT balance initialization : ", balpBATInit)

	balDAIInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.DAIAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] DAI balance initialization : ", balDAIInit)

	balDAIScInit := tradingSuite.getDepositedBalance(
		tradingSuite.DAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] DAI balance initialization on SC : ", balDAIScInit)

	balpDAIInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncDAITokenIDStr)
	fmt.Println("[INFO] pDAI balance initialization : ", balpDAIInit)

	balSAIInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.SAIAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] SAI balance initialization : ", balSAIInit)

	balSAIScInit := tradingSuite.getDepositedBalance(
		tradingSuite.SAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] SAI balance initialization on SC : ", balSAIScInit)

	balpSAIInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncSAITokenIDStr)
	fmt.Println("[INFO] pSAI balance initialization : ", balpSAIInit)
}

func (tradingSuite *ZRXTradingTestSuite) Test15DepositAndWithdrwaERC20token() {
return
	fmt.Println("============ TEST 15 DEPOSIT AND WITHDRAW ERC20 TOKEN (DAI) ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()

	// get info balance initialization
	balPrvInit, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance initialization : ", balPrvInit)

	balEthInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance initialization : ", balEthInit)

	balpBATInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncDAITokenIDStr)
	fmt.Println("[INFO] pDAI balance initialization : ", balpBATInit)

	balBATInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.DAIAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] DAI balance initialization : ", balBATInit)

	balEthScInit := tradingSuite.getDepositedBalance(
		tradingSuite.DAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] DAI balance initialization on SC : ", balEthScInit)

	fmt.Println("------------ STEP 1: porting DAI to pDAI --------------")

	fmt.Println("amount DAI deposit : ", (big.NewInt(int64(0.012 * params.Ether))))
	deposit := big.NewInt(int64(0.002 * params.Ether))
	burningPETH := big.NewInt(0).Div(deposit, big.NewInt(1000000000))
	// Deposit to proof
	txHash := tradingSuite.depositERC20ToBridge(
		deposit,
		common.HexToAddress(tradingSuite.DAIAddressStr),
		tradingSuite.IncPaymentAddrStr,
	)

	fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHash))
	// get ETH remain after depsit
	balEthAfDep := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain after deposit : ", balEthAfDep)

	//require.Equal(tradingSuite.T(), balEthAfDep, big.NewInt(0).Sub(big.NewInt(0).Sub(balEthInit, big.NewInt(int64(tradingSuite.DepositingEther*params.Ether))), tradingSuite.getGasFeeETHbyTxhash(txHash)), "balance ETH incorrect")

	// Proof
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(100 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncDAITokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)

	time.Sleep(120 * time.Second)

	balpBATAfIssS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncDAITokenIDStr)
	fmt.Println("[INFO] pDAI balance after issuing step 1 : ", balpBATAfIssS1)
	// check PRV and token balance after issuing
	balPrvAfIssS1, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after issuing step 1: ", balPrvAfIssS1)

	fmt.Println("------------ STEP 2: burning pDAI to deposit DAI to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningPToken(
		tradingSuite.IncDAITokenIDStr,
		burningPETH,
		pubKeyToAddrStr[2:],
		"createandsendburningfordeposittoscrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	// check PRV and token balance after burning
	balPrvAfBurnS2, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after burning step 2: ", balPrvAfBurnS2)

	balpBATAfBurnS2, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncDAITokenIDStr)
	fmt.Println("[INFO] pDAI balance after burning step 2 : ", balpBATAfBurnS2)
	// TODO assert pETH balance issuing

	txHash2 := tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHash2)
	require.Equal(tradingSuite.T(), nil, err)

	fmt.Println("Burn Proof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHash2))
	// get ETH remain
	balEthAfDep2 := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain  : ", balEthAfDep2)
	// TODO : assert ETH balance
	//require.Equal(tradingSuite.T(), balEthAfDep2, big.NewInt(0).Sub(balEthAfDep, tradingSuite.getGasFeeETHbyTxhash(txHash2)), "balance ETH incorrect")

	balBATScDepS2 := tradingSuite.getDepositedBalance(
		tradingSuite.DAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] DAI balance after deposit on SC at step 2: ", balBATScDepS2)
	// TODO assert ETH balane on SC

	fmt.Println("------------ STEP 3: withdraw DAI to deposit pDAI to Incognito  --------------")

	txHashByEmittingWithdrawalReq1 := tradingSuite.requestWithdraw(
		tradingSuite.DAIAddressStr,
		balBATScDepS2,
	)

	balBATScDepS41 := tradingSuite.getDepositedBalance(
		tradingSuite.DAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] DAI balance on SC after withdraw at step 3 : ", balBATScDepS41)
	// TODO assert BAT balane on SC

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq1)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(120 * time.Second)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncDAITokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(100 * time.Second)

	balpBATAfIssS41, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncDAITokenIDStr,
	)
	fmt.Println("[INFO] pDAI balance after issuing step 3 : ", balpBATAfIssS41)

	fmt.Println("------------ STEP 4: withdraw pDAI to deposit DAI   --------------")

	withdrawingPBAT := big.NewInt(0).Div(balBATScDepS2, big.NewInt(1000000000))

	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncDAITokenIDStr,
		withdrawingPBAT,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	balpBATAfBurnS51, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncDAITokenIDStr,
	)
	fmt.Println("[INFO] pDAI balance after burning step 4 : ", balpBATAfBurnS51)

	txhash3 := tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))
	fmt.Println("txHasd burn ")

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txhash3)
	require.Equal(tradingSuite.T(), nil, err)

	fmt.Println("Burn Proof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHash2))

	balBATAfDep51 := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.DAIAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] DAI balance after withdraw  : ", balBATAfDep51)
}

func (tradingSuite *ZRXTradingTestSuite) Test16DepositAndWithdrwaERC20tokenDecimal8() {
//return
	fmt.Println("============ TEST 15 DEPOSIT AND WITHDRAW ERC20 TOKEN DECIMAL 8 (CKN) ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()

	// get info balance initialization
	balPrvInit, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance initialization : ", balPrvInit)

	balEthInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance initialization : ", balEthInit)

	balpBATInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncCKNTokenIDStr)
	fmt.Println("[INFO] pCKN balance initialization : ", balpBATInit)

	balBATInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.CKNAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] CKN balance initialization : ", balBATInit)

	balEthScInit := tradingSuite.getDepositedBalance(
		tradingSuite.CKNAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] CKN balance initialization on SC : ", balEthScInit)

	fmt.Println("------------ STEP 1: porting CKN to pCKN --------------")
	deposit := big.NewInt(int64(1000000))
	fmt.Println("amount CKN deposit : ", deposit)

	burningPETH := deposit
	// Deposit to proof
	txHash := tradingSuite.depositERC20ToBridge(
		deposit,
		common.HexToAddress(tradingSuite.CKNAddressStr),
		tradingSuite.IncPaymentAddrStr,
	)

	fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHash))
	// get ETH remain after depsit
	balEthAfDep := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain after deposit : ", balEthAfDep)

	//require.Equal(tradingSuite.T(), balEthAfDep, big.NewInt(0).Sub(big.NewInt(0).Sub(balEthInit, big.NewInt(int64(tradingSuite.DepositingEther*params.Ether))), tradingSuite.getGasFeeETHbyTxhash(txHash)), "balance ETH incorrect")

	// Proof
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(100 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncCKNTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)

	time.Sleep(120 * time.Second)

	balpBATAfIssS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncCKNTokenIDStr)
	fmt.Println("[INFO] pCKN balance after issuing step 1 : ", balpBATAfIssS1)
	// check PRV and token balance after issuing
	balPrvAfIssS1, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after issuing step 1: ", balPrvAfIssS1)

	fmt.Println("------------ STEP 2: burning pCKN to deposit CKN to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningPToken(
		tradingSuite.IncCKNTokenIDStr,
		burningPETH,
		pubKeyToAddrStr[2:],
		"createandsendburningfordeposittoscrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	// check PRV and token balance after burning
	balPrvAfBurnS2, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after burning step 2: ", balPrvAfBurnS2)

	balpBATAfBurnS2, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncCKNTokenIDStr)
	fmt.Println("[INFO] pCKN balance after burning step 2 : ", balpBATAfBurnS2)
	// TODO assert pETH balance issuing

	txHash2 := tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))

	fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHash2))
	// get ETH remain
	balEthAfDep2 := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain  : ", balEthAfDep2)
	// TODO : assert ETH balance
	//require.Equal(tradingSuite.T(), balEthAfDep2, big.NewInt(0).Sub(balEthAfDep, tradingSuite.getGasFeeETHbyTxhash(txHash2)), "balance ETH incorrect")

	balBATScDepS2 := tradingSuite.getDepositedBalance(
		tradingSuite.CKNAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] CKN balance after deposit on SC at step 2: ", balBATScDepS2)
	// TODO assert ETH balane on SC

	fmt.Println("------------ STEP 3: withdraw CKN to deposit pCKN to Incognito  --------------")

	txHashByEmittingWithdrawalReq1 := tradingSuite.requestWithdraw(
		tradingSuite.CKNAddressStr,
		balBATScDepS2,
	)

	balBATScDepS41 := tradingSuite.getDepositedBalance(
		tradingSuite.CKNAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] CKN balance on SC after withdraw at step 3 : ", balBATScDepS41)
	// TODO assert BAT balane on SC

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq1)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(100 * time.Second)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncCKNTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(100 * time.Second)

	balpBATAfIssS41, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncCKNTokenIDStr,
	)
	fmt.Println("[INFO] pCKN balance after issuing step 3 : ", balpBATAfIssS41)

	fmt.Println("------------ STEP 4: withdraw pCKN to deposit CKN   --------------")

	withdrawingPBAT := balBATScDepS2

	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncCKNTokenIDStr,
		withdrawingPBAT,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	balpBATAfBurnS51, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncCKNTokenIDStr,
	)
	fmt.Println("[INFO] pCKN balance after burning step 4 : ", balpBATAfBurnS51)

	tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))

	balBATAfDep51 := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.CKNAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] CKN balance after withdraw  : ", balBATAfDep51)
}

func (tradingSuite *ZRXTradingTestSuite) Test17TradeEthForDaiSaiBatWith0x() {
return
	fmt.Println("============ TEST 17 TRADE ETHER FOR DAI SAI BAT WITH 0X AGGREGATOR ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()

	// get info balance initialization
	balPrvInit, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance initialization : ", balPrvInit)

	balEthInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance initialization : ", balEthInit)

	balEthScInit := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance initialization on SC : ", balEthScInit)

	balpEthInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance initialization : ", balpEthInit)

	balDaiInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.DAIAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] DAI balance initialization : ", balDaiInit)

	balDaiScInit := tradingSuite.getDepositedBalance(
		tradingSuite.DAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] DAI balance initialization on SC : ", balDaiScInit)

	balpDaiInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncDAITokenIDStr)
	fmt.Println("[INFO] pDAI balance initialization : ", balpDaiInit)

	balSaiInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.SAIAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] SAI balance initialization : ", balSaiInit)

	balSaiScInit := tradingSuite.getDepositedBalance(
		tradingSuite.SAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] SAI balance initialization on SC : ", balSaiScInit)

	balpSaiInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncSAITokenIDStr)
	fmt.Println("[INFO] pSAI balance initialization : ", balpSaiInit)

	balBatInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.BATAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] BAT balance initialization : ", balBatInit)

	balBATScInit := tradingSuite.getDepositedBalance(
		tradingSuite.BATAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] BAT balance initialization on SC : ", balBATScInit)

	balpBATInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncBATTokenIDStr)
	fmt.Println("[INFO] pBAT balance initialization : ", balpBATInit)

	fmt.Println("------------ STEP 1: porting ETH to pETH --------------")
	fmt.Println("amount ETH deposit : ", (big.NewInt(int64(tradingSuite.DepositingEther * 3 * params.Ether))))
	depositEth := big.NewInt(int64(tradingSuite.DepositingEther * 3 * params.Ether))
	burningETH := big.NewInt(0).Div(depositEth, big.NewInt(1000000000))
	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))

	// Deposit to proof
	txHash := tradingSuite.depositETH(
		tradingSuite.DepositingEther*3,
		tradingSuite.IncPaymentAddrStr,
	)

	fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHash))
	// get ETH remain after depsit
	balEthAfDep := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain after deposit : ", balEthAfDep)

	//require.Equal(tradingSuite.T(), balEthAfDep, big.NewInt(0).Sub(big.NewInt(0).Sub(balEthInit, big.NewInt(int64(tradingSuite.DepositingEther*params.Ether))), tradingSuite.getGasFeeETHbyTxhash(txHash)), "balance ETH incorrect")

	// Proof
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(120 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)

	time.Sleep(100 * time.Second)
	// check PRV and token balance after issuing
	balPrvAfIssS1, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after issuing step 1: ", balPrvAfIssS1)
	//require.NotEqual(tradingSuite.T(), balPrvAfIssS1, (balPrvInit - tradingSuite.getFeePRVbyTxhashInC(issuuRes["TxID"].(string))), "Balance PRV remain incorrect after issuu step 1")

	balpEthAfIssS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance after issuing step 1 : ", balpEthAfIssS1)
	//require.Equal(tradingSuite.T(), big.NewInt(int64(balpEthAfIssS1-balpEthInit)), big.NewInt(0).Div(big.NewInt(int64(tradingSuite.DepositingEther*params.Ether)), big.NewInt(1000000000)), " balnce pToken issuu incorrect")

	fmt.Println("------------ STEP 2: burning pETH to deposit ETH to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningPToken(
		tradingSuite.IncEtherTokenIDStr,
		burningETH,
		pubKeyToAddrStr[2:],
		"createandsendburningfordeposittoscrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(100 * time.Second)

	// check PRV and token balance after burning
	balPrvAfBurnS2, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after burning step 2: ", balPrvAfBurnS2)
	//require.NotEqual(tradingSuite.T(), balPrvAfBurnS2, (balPrvAfIssS1 - tradingSuite.getFeePRVbyTxhashInC(burningRes["TxID"].(string))), "Balance PRV remain incorrect after burn step 2")

	balpEthAfBurnS2, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance after burning step 2 : ", balpEthAfBurnS2)
	// TODO assert pETH balance issuing

	txHash2 := tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))
	fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHash2))
	// get ETH remain
	balEthAfDep2 := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain  : ", balEthAfDep2)
	// TODO : assert ETH balance
	//require.Equal(tradingSuite.T(), balEthAfDep2, big.NewInt(0).Sub(balEthAfDep, tradingSuite.getGasFeeETHbyTxhash(txHash2)), "balance ETH incorrect")

	balEthScDepS2 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance after deposit on SC at step 2: ", balEthScDepS2)
	// TODO assert ETH balane on SC
	balDaiScS2 := tradingSuite.getDepositedBalance(
		tradingSuite.DAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] DAI balance on SC at step 2 : ", balDaiScS2)

	balSaiScS2 := tradingSuite.getDepositedBalance(
		tradingSuite.SAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] SAI balance on SC at step 2 : ", balSaiScS2)

	balBatScS2 := tradingSuite.getDepositedBalance(
		tradingSuite.BATAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] BAT balance on SC at step 2 : ", balBatScS2)

	//require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPETH, big.NewInt(1000000000)), deposited)

	fmt.Println("------------ STEP 3: execute trade ETH for DAI through 0x aggregator --------------")
	tradingSuite.executeWith0x(
		tradeAmount,
		"ETH",
		tradingSuite.EtherAddressStr,
		"DAI",
		tradingSuite.DAIAddressStr,
	)

	tradingSuite.executeWith0x(
		tradeAmount,
		"ETH",
		tradingSuite.EtherAddressStr,
		"SAI",
		tradingSuite.SAIAddressStr,
	)

	tradingSuite.executeWith0x(
		tradeAmount,
		"ETH",
		tradingSuite.EtherAddressStr,
		"BAT",
		tradingSuite.BATAddressStr,
	)
	// fmt.Println("protocol fee :", protoFee)
	// fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHash3))
	// get ETH remain
	// balEthAfDep3 := tradingSuite.getBalanceOnETHNet(
	// 	common.HexToAddress(tradingSuite.EtherAddressStr),
	// 	common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	// )
	// fmt.Println("[INFO] ETH balance remain  : ", balEthAfDep3)
	//require.Equal(tradingSuite.T(), balEthAfDep3, big.NewInt(0).Sub(big.NewInt(0).Sub(balEthAfDep2, tradingSuite.getGasFeeETHbyTxhash(txHash3)), protoFee), "balance ETH incorrect")
	balEthScTradeS3 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC after trade at step 3 : ", balEthScTradeS3)
	// TODO assert ETH balane on SC
	balDaiScTradeS3 := tradingSuite.getDepositedBalance(
		tradingSuite.DAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] DAI balance on SC after trade at step 3 : ", balDaiScTradeS3)
	// TODO assert DAI balane on SC
	//require.NotEqual(tradingSuite.T(), balDaiScTradeS3, balDaiScS2, "trade failed")

	balSaiScS3 := tradingSuite.getDepositedBalance(
		tradingSuite.SAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] SAI balance on SC after trade at step 3: ", balSaiScS3)

	balBatScS3 := tradingSuite.getDepositedBalance(
		tradingSuite.BATAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] BAT balance on SC after trade at step 3 : ", balBatScS3)

	///////////////////////////

	// tradingSuite.executeWith0x(
	// 	tradeAmount,
	// 	"BAT",
	// 	tradingSuite.BATAddressStr,
	// 	"ETH",
	// 	tradingSuite.EtherAddressStr,	
	// )

	// balEthScTradeS3 = tradingSuite.getDepositedBalance(
	// 	tradingSuite.EtherAddressStr,
	// 	pubKeyToAddrStr,
	// )
	// fmt.Println("[INFO] ETH balance on SC after trade at step 3.1111 : ", balEthScTradeS3)

	// balBatScS3 = tradingSuite.getDepositedBalance(
	// 	tradingSuite.BATAddressStr,
	// 	pubKeyToAddrStr,
	// )
	// fmt.Println("[INFO] BAT balance on SC after trade at step 3.111 : ", balBatScS3)

	///////////////////////////

	fmt.Println("------------ STEP 4: withdrawing to pToken on Incognito --------------")
	txHashByEmittingWithdrawalReqDAI := tradingSuite.requestWithdraw(
		tradingSuite.DAIAddressStr,
		balDaiScTradeS3,
	)

	txHashByEmittingWithdrawalReqSAI := tradingSuite.requestWithdraw(
		tradingSuite.SAIAddressStr,
		balSaiScS3,
	)

	txHashByEmittingWithdrawalReqBAT := tradingSuite.requestWithdraw(
		tradingSuite.BATAddressStr,
		balBatScS3,
	)

	//fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHashByEmittingWithdrawalReqDAI))
	// get ETH remain
	balEthAfDep4 := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain  : ", balEthAfDep4)
	//require.Equal(tradingSuite.T(), balEthAfDep4, big.NewInt(0).Sub(balEthAfDep3, tradingSuite.getGasFeeETHbyTxhash(txHashByEmittingWithdrawalReq)), "balance ETH incorrect")

	balDaiScDepS4 := tradingSuite.getDepositedBalance(
		tradingSuite.DAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] DAI balance on SC after withdraw at step 4 : ", balDaiScDepS4)

	balSaiScDepS4 := tradingSuite.getDepositedBalance(
		tradingSuite.SAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] SAI balance on SC after withdraw at step 4 : ", balSaiScDepS4)

	balBATScDepS4 := tradingSuite.getDepositedBalance(
		tradingSuite.BATAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] BAT balance on SC after withdraw at step 4 : ", balBATScDepS4)

	// TODO assert DAI balane on SC
	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReqDAI)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	_, ethBlockHashS, ethTxIdxS, ethDepositProofS, err := getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReqSAI)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHashS, ethTxIdxS, ethDepositProofS)

	_, ethBlockHashB, ethTxIdxB, ethDepositProofB, err := getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReqBAT)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHashB, ethTxIdxB, ethDepositProofB)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(100 * time.Second)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncDAITokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(60 * time.Second)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncSAITokenIDStr,
		ethDepositProofS,
		ethBlockHashS,
		ethTxIdxS,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(60 * time.Second)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncBATTokenIDStr,
		ethDepositProofB,
		ethBlockHashB,
		ethTxIdxB,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(100 * time.Second)

	balpDaiAfIssS4, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncDAITokenIDStr,
	)
	fmt.Println("[INFO] pDAI balance after issuing step 4 : ", balpDaiAfIssS4)

	balpSaiAfIssS4, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncSAITokenIDStr,
	)
	fmt.Println("[INFO] pDAI balance after issuing step 4 : ", balpSaiAfIssS4)

	balpBatAfIssS4, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncBATTokenIDStr,
	)
	fmt.Println("[INFO] pBAT balance after issuing step 4 : ", balpBatAfIssS4)
	// TODO assert pDai balance issuing
	balPrvAfIssS4, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after issuing step 4: ", balPrvAfIssS4)

	fmt.Println("------------ step 5: withdrawing pDAI from Incognito to DAI --------------")
	withdrawingPDAI := big.NewInt(0).Div(balDaiScTradeS3, big.NewInt(1000000000))
	withdrawingPSAI := big.NewInt(0).Div(balSaiScS3, big.NewInt(1000000000))
	withdrawingPBAT := big.NewInt(0).Div(balBatScS3, big.NewInt(1000000000))

	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncDAITokenIDStr,
		withdrawingPDAI,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningrequest",
	)
	time.Sleep(60 * time.Second)

	burningResS, err := tradingSuite.callBurningPToken(
		tradingSuite.IncSAITokenIDStr,
		withdrawingPSAI,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningrequest",
	)
	time.Sleep(60 * time.Second)

	burningResB, err := tradingSuite.callBurningPToken(
		tradingSuite.IncBATTokenIDStr,
		withdrawingPBAT,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningrequest",
	)

	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	burningTxIDS, found := burningResS["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	burningTxIDB, found := burningResB["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(100 * time.Second)

	balpDaiAfBurnS5, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncDAITokenIDStr,
	)
	fmt.Println("[INFO] pDAI balance after burning step 5 : ", balpDaiAfBurnS5)

	balpSaiAfBurnS5, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncSAITokenIDStr,
	)
	fmt.Println("[INFO] pSAI balance after burning step 5 : ", balpSaiAfBurnS5)

	balpBatAfBurnS5, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncBATTokenIDStr,
	)
	fmt.Println("[INFO] pBAT balance after burning step 5 : ", balpBatAfBurnS5)
	// TODO assert pDai balance issuing
	balPrvAfBrunS5, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after burning step 5: ", balPrvAfBrunS5)
	//require.NotEqual(tradingSuite.T(), balPrvAfBrunS5, (balPrvAfIssS4 - tradingSuite.getFeePRVbyTxhashInC(burningRes["TxID"].(string))), "Balance PRV remain incorrect after burn step 5")

	tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))
	tradingSuite.submitBurnProofForWithdrawal(burningTxIDS.(string))
	tradingSuite.submitBurnProofForWithdrawal(burningTxIDB.(string))
	// fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHash5))
	// // get ETH remain
	// balEthAfDep5 := tradingSuite.getBalanceOnETHNet(
	// 	common.HexToAddress(tradingSuite.EtherAddressStr),
	// 	common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	// )
	// fmt.Println("[INFO] ETH balance remain  : ", balEthAfDep5)
	//require.Equal(tradingSuite.T(),balEthAfDep5,big.NewInt(0).Sub(balEthAfDep4,tradingSuite.getGasFeeETHbyTxhash(txHash5)),"balance ETH incorrect")

	balDai := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.DAIAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("DAI balance after trade: ", balDai)

	balSai := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.SAIAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("SAI balance after trade: ", balSai)

	balBat := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.BATAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("BAT balance after trade: ", balBat)
	// require.Equal(tradingSuite.T(), withdrawingPDAI.Uint64(), bal.Div(bal, big.NewInt(1000000000)).Uint64())
	balEth := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("ETH balance after trade: ", balEth)
}

func (tradingSuite *ZRXTradingTestSuite) TestXXXXXXX() {
return
	//     fmt.Println("----------test txhash -----")
	//     pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()

	//     fmt.Println("key 1 :", pubKeyToAddrStr)

	//    // IncPaymentAddrStr2 := "12S6R8HfTyL74bggg47LX88RSvBPaMPBMEMoo6yx9WQ4EgLiYERXXcE2Mv2HrCsFuKhBsTfrYMeH82Bus5MHQGt3xHwoxX4v2qM5jRE"
	//   IncPriAddrStr2 := "112t8rnX3VTd3MTWMpfbYP8HGY4ToAaLjrmUYzfjJBrAcb8iPLkNqvVDXWrLNiFV5yb2NBpR3FDZj3VW8GcLUwRdQ61hPMWP3YrREZAZ1UbH"

	//   key2 :=testgenKeysPairForSC(IncPriAddrStr2)

	//   fmt.Println("key 2 :", key2)

	txHash := common.HexToHash("b063bb84a6aa3b38682dc7c73184825a53e3e129af01d406cdd6d3e111e0fbeb")

	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("Proof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

}
