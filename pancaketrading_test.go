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
type PancakeTradingTestSuite struct {
	*TradingTestSuite

	PanackeTradeDeployedAddr common.Address
	PanackeRouteContractAddr common.Address

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

func NewPanckeTradingTestSuite(tradingTestSuite *TradingTestSuite) *PancakeTradingTestSuite {
	return &PancakeTradingTestSuite{
		TradingTestSuite: tradingTestSuite,
	}
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (tradingSuite *PancakeTradingTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")
	// Bsc testnet env
	tradingSuite.IncBUSDTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000062"
	tradingSuite.PanackeTradeDeployedAddr = common.HexToAddress("0xD772475E9CFA0F1134aef4B3A5BA83De014DEc45")
	tradingSuite.PanackeRouteContractAddr = common.HexToAddress("0x9Ac64Cc6e4415144C455BD8E4837Fea55603e5c3")

	// tokens
	tradingSuite.WBNBAddr = common.HexToAddress("0xae13d989dac2f0debff460ac112a837c89baa7cd")
	tradingSuite.WBTCAddr = common.HexToAddress("0x6ce8da28e2f864420840cf74474eff5fd80e65b8")
	tradingSuite.WUSDTAddr = common.HexToAddress("0x7ef95a0fee0dd31b22626fa2e10ee6a223f8a684") // USDT
	tradingSuite.WBUSDAddr = common.HexToAddress("0x78867bbeef44f2326bf8ddd1941a4439382ef2a7")
	tradingSuite.WXRPAddr = common.HexToAddress("0xa83575490d7df4e2f47b7d38ef351a2722ca45b9")

	tradingSuite.DepositingEther = float64(0.01)
}

func (tradingSuite *PancakeTradingTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
	tradingSuite.BSCClient.Close()
}

func (tradingSuite *PancakeTradingTestSuite) SetupTest() {
	fmt.Println("Setting up the test...")
}

func (tradingSuite *PancakeTradingTestSuite) TearDownTest() {
	fmt.Println("Tearing down the test...")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestPancakeTradingTestSuite(t *testing.T) {
	fmt.Println("Starting entry point for Uniswap test suite...")

	tradingSuite := new(TradingTestSuite)
	suite.Run(t, tradingSuite)

	uniswapTradingSuite := NewPanckeTradingTestSuite(tradingSuite)
	suite.Run(t, uniswapTradingSuite)

	fmt.Println("Finishing entry point for 0x test suite...")
}

func (tradingSuite *PancakeTradingTestSuite) getExpectedAmount(
	path []common.Address,
	srcQty *big.Int,
) *big.Int {
	c, err := pancakeproxy.NewPancakeproxyrouter(tradingSuite.PanackeRouteContractAddr, tradingSuite.BSCClient)
	require.Equal(tradingSuite.T(), nil, err)
	amounts, err := c.GetAmountsOut(nil, srcQty, path)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Printf("intput value: %v\n", srcQty.String())
	fmt.Printf("output value: %v\n", amounts[len(amounts)-1].String())
	return amounts[1]
}

func (tradingSuite *PancakeTradingTestSuite) executeWithPancake(
	srcQty *big.Int,
	path []common.Address,
	deadline uint,
	isNative bool,
) {
	require.NotEqual(tradingSuite.T(), 0, len(path))

	tradeAbi, err := abi.JSON(strings.NewReader(pancakeproxy.PancakeproxyMetaData.ABI))
	require.Equal(tradingSuite.T(), nil, err)

	// Get contract instance
	c, err := vault.NewVault(tradingSuite.VaultBSCAddr, tradingSuite.BSCClient)
	require.Equal(tradingSuite.T(), nil, err)
	auth, err := bind.NewKeyedTransactorWithChainID(tradingSuite.ETHPrivKey, big.NewInt(int64(tradingSuite.ChainIDBSC)))
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
		Prefix:    BSC_EXECUTE_PREFIX,
		Token:     sourceToken,
		Timestamp: timestamp,
		Amount:    srcQty,
	}
	tempData, err := vaultAbi.Pack("_buildSignExecute", psData, destToken, tradingSuite.PanackeTradeDeployedAddr, input)
	require.Equal(tradingSuite.T(), nil, err)
	data := rawsha3(tempData[4:])
	signBytes, err := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)
	require.Equal(tradingSuite.T(), nil, err)

	tx, err := c.Execute(
		auth,
		sourceToken,
		srcQty,
		destToken,
		tradingSuite.PanackeTradeDeployedAddr,
		input,
		timestamp,
		signBytes,
	)
	require.Equal(tradingSuite.T(), nil, err)
	txHash := tx.Hash()
	if err := wait(tradingSuite.BSCClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("Pancake trade executed , txHash: %x\n", txHash[:])
}

func (tradingSuite *PancakeTradingTestSuite) Test1TradeEthForDAIWithPancake() {
	fmt.Println("============ TEST TRADE BNB FOR USDT WITH pancake ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
	burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("------------ STEP 1: porting BNB to pBNB --------------")
	txHash := tradingSuite.depositETH(
		tradingSuite.DepositingEther,
		tradingSuite.IncPaymentAddrStr,
		tradingSuite.VaultBSCAddr,
		tradingSuite.BSCClient,
	)
	time.Sleep(15 * time.Second)
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.BSCHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(50 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingbscreq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)

	fmt.Println("------------ STEP 2: burning pBNB to deposit BNB to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningPToken(
		tradingSuite.IncEtherTokenIDStr,
		burningPETH,
		pubKeyToAddrStr[2:],
		"createandsendburningpbscfordeposittoscrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForDepositToSC(
		burningTxID.(string),
		big.NewInt(int64(tradingSuite.ChainIDBSC)),
		"getburnpbscprooffordeposittosc",
		tradingSuite.VaultBSCAddr,
		tradingSuite.BSCClient,
	)
	deposited := tradingSuite.getDepositedBalanceBSC(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		pubKeyToAddrStr,
	)
	fmt.Printf("address own asset %v \n", pubKeyToAddrStr)

	fmt.Println("------------ step 3: execute trade BNB for USDT through Pancake --------------")
	tradingSuite.executeWithPancake(
		deposited,
		[]common.Address{
			tradingSuite.WBNBAddr,
			tradingSuite.WBUSDAddr,
			tradingSuite.WUSDTAddr,
		},
		uint(time.Now().Unix()+60000),
		false,
	)
	time.Sleep(15 * time.Second)
	daiTraded := tradingSuite.getDepositedBalanceBSC(
		tradingSuite.WUSDTAddr,
		pubKeyToAddrStr,
	)
	fmt.Println("usdtTraded: ", daiTraded)

	fmt.Println("------------ step 4: withdrawing USDT from SC to pUSDT on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.WUSDTAddr.String(),
		daiTraded,
		tradingSuite.BSCClient,
		big.NewInt(int64(tradingSuite.ChainIDBSC)),
		tradingSuite.VaultBSCAddr,
		BSC_REQ_WITHDRAW_PREFIX,
	)
	time.Sleep(45 * time.Second)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.BSCHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(50 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncUSDTTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingbscreq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)

	fmt.Println("------------ step 5: withdrawing pUSDT from Incognito to USDT --------------")
	withdrawingPDAI := big.NewInt(0).Div(daiTraded, big.NewInt(1e9))
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncUSDTTokenIDStr,
		withdrawingPDAI,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningbscrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForWithdrawal(
		burningTxID.(string),
		"getbscburnproof",
		tradingSuite.VaultBSCAddr,
		tradingSuite.BSCClient,
		tradingSuite.ChainIDBSC,
	)

	bal := tradingSuite.getBalanceOnETHNet(
		tradingSuite.WUSDTAddr,
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		tradingSuite.BSCClient,
	)
	tradingSuite.DAIBalanceAfterStep1 = bal
	fmt.Println("USDT balance after step 1: ", tradingSuite.DAIBalanceAfterStep1)
}

func (tradingSuite *PancakeTradingTestSuite) Test2TradeDAIForMRKWithUniswap() {
	fmt.Println("============ TEST TRADE USDT FOR BUSD WITH PANCAKE ===========")
	fmt.Println("------------ step 0: declaration & initialization --------------")
	depositingDAI := tradingSuite.DAIBalanceAfterStep1
	burningPDAI := big.NewInt(0).Div(depositingDAI, big.NewInt(1000000000))

	daibal := tradingSuite.getBalanceOnETHNet(
		tradingSuite.WUSDTAddr,
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		tradingSuite.BSCClient,
	)
	fmt.Println("dai balance of owner: ", daibal)

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("pubKeyToAddrStr: ", pubKeyToAddrStr)

	fmt.Println("------------ step 1: porting USDT to pUSDT --------------")
	txHash := tradingSuite.depositERC20ToBridge(
		depositingDAI,
		tradingSuite.WUSDTAddr,
		tradingSuite.IncPaymentAddrStr,
		tradingSuite.VaultBSCAddr,
		tradingSuite.BSCClient,
		tradingSuite.ChainIDBSC,
	)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.BSCHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(50 * time.Second)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncUSDTTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingbscreq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)
	// TODO: check the new balance on peth of the incognito account

	fmt.Println("------------ step 2: burning pUSDT to deposit USDT to SC --------------")

	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningPToken(
		tradingSuite.IncUSDTTokenIDStr,
		burningPDAI,
		pubKeyToAddrStr[2:],
		"createandsendburningpbscfordeposittoscrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForDepositToSC(
		burningTxID.(string),
		big.NewInt(int64(tradingSuite.ChainIDBSC)),
		"getburnpbscprooffordeposittosc",
		tradingSuite.VaultBSCAddr,
		tradingSuite.BSCClient,
	)

	deposited := tradingSuite.getDepositedBalanceBSC(
		tradingSuite.WUSDTAddr,
		pubKeyToAddrStr,
	)
	fmt.Println("deposited dai: ", deposited)
	// require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPDAI, big.NewInt(1000000000)), deposited)

	fmt.Println("------------ step 3: execute trade DAI for MRK through Uniswap aggregator --------------")
	tradingSuite.executeWithPancake(
		deposited,
		[]common.Address{
			tradingSuite.WUSDTAddr,
			tradingSuite.WBUSDAddr,
		},
		uint(time.Now().Unix()+60000),
		false,
	)
	time.Sleep(20 * time.Second)
	busdTraded := tradingSuite.getDepositedBalanceBSC(
		tradingSuite.WBUSDAddr,
		pubKeyToAddrStr,
	)
	fmt.Println("busd Traded: ", busdTraded)

	fmt.Println("------------ step 4: withdrawing BUSD from SC to pBUSD on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.WBUSDAddr.String(),
		busdTraded,
		tradingSuite.BSCClient,
		big.NewInt(int64(tradingSuite.ChainIDBSC)),
		tradingSuite.VaultBSCAddr,
		BSC_REQ_WITHDRAW_PREFIX,
	)
	time.Sleep(15 * time.Second)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.BSCHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(90 * time.Second)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncBUSDTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingbscreq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)

	fmt.Println("------------ step 5: withdrawing pMRK from Incognito to MRK --------------")
	withdrawingPBUSD := big.NewInt(0).Div(busdTraded, big.NewInt(1000000000)) // MRK decimal 18
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncBUSDTokenIDStr,
		withdrawingPBUSD,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningbscrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForWithdrawal(
		burningTxID.(string),
		"getbscburnproof",
		tradingSuite.VaultBSCAddr,
		tradingSuite.BSCClient,
		tradingSuite.ChainIDBSC,
	)

	bal := tradingSuite.getBalanceOnETHNet(
		tradingSuite.WBUSDAddr,
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		tradingSuite.BSCClient,
	)
	tradingSuite.MRKBalanceAfterStep2 = bal
	fmt.Println("BUSD balance after step 2: ", tradingSuite.MRKBalanceAfterStep2)
}

func (tradingSuite *PancakeTradingTestSuite) Test3TradeMRKForEthWithUniswap() {
	fmt.Println("============ TEST TRADE BUSD FOR BNB WITH PANCAKE ===========")
	fmt.Println("------------ step 0: declaration & initialization --------------")
	depositingMRK := tradingSuite.MRKBalanceAfterStep2
	burningPMRK := big.NewInt(0).Div(depositingMRK, big.NewInt(1000000000))

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("pubKeyToAddrStr: ", pubKeyToAddrStr)

	fmt.Println("------------ step 1: porting BUSD to pBUSD --------------")
	txHash := tradingSuite.depositERC20ToBridge(
		depositingMRK,
		tradingSuite.WBUSDAddr,
		tradingSuite.IncPaymentAddrStr,
		tradingSuite.VaultBSCAddr,
		tradingSuite.BSCClient,
		tradingSuite.ChainIDBSC,
	)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.BSCHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(50 * time.Second)

	issuingRes, err := tradingSuite.callIssuingETHReq(
		tradingSuite.IncBUSDTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingbscreq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("issuingRes: ", issuingRes)
	time.Sleep(120 * time.Second)

	fmt.Println("------------ step 2: burning pBUSD to deposit BUSD to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningPToken(
		tradingSuite.IncBUSDTokenIDStr,
		burningPMRK,
		pubKeyToAddrStr[2:],
		"createandsendburningpbscfordeposittoscrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForDepositToSC(
		burningTxID.(string),
		big.NewInt(int64(tradingSuite.ChainIDBSC)),
		"getburnpbscprooffordeposittosc",
		tradingSuite.VaultBSCAddr,
		tradingSuite.BSCClient,
	)

	deposited := tradingSuite.getDepositedBalanceBSC(
		tradingSuite.WBUSDAddr,
		pubKeyToAddrStr,
	)
	fmt.Println("deposited WBUSDAddr: ", deposited)
	// require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPMRK, big.NewInt(1000000000)), deposited)

	fmt.Println("------------ step 3: execute trade BUSD for BNB through Pancake aggregator --------------")
	tradingSuite.executeWithPancake(
		big.NewInt(0).Div(deposited, big.NewInt(2)),
		[]common.Address{
			tradingSuite.WBUSDAddr,
			tradingSuite.WBNBAddr,
		},
		uint(time.Now().Unix()+60000),
		true,
	)
	etherTraded := tradingSuite.getDepositedBalanceBSC(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		pubKeyToAddrStr,
	)
	fmt.Println("bnb Traded: ", etherTraded)
	require.NotEqual(tradingSuite.T(), 0, etherTraded)
	time.Sleep(5 * time.Second)
	tradingSuite.executeWithPancake(
		big.NewInt(0).Div(deposited, big.NewInt(2)),
		[]common.Address{
			tradingSuite.WBUSDAddr,
			tradingSuite.WBNBAddr,
		},
		uint(time.Now().Unix()+60000),
		false,
	)
	wBNBTraded := tradingSuite.getDepositedBalanceBSC(
		tradingSuite.WBNBAddr,
		pubKeyToAddrStr,
	)
	fmt.Println("wBNB Traded: ", wBNBTraded)
	require.NotEqual(tradingSuite.T(), 0, wBNBTraded)

	fmt.Println("------------ step 4: withdrawing ETH from SC to pETH on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.EtherAddressStr,
		etherTraded,
		tradingSuite.BSCClient,
		big.NewInt(int64(tradingSuite.ChainIDBSC)),
		tradingSuite.VaultBSCAddr,
		BSC_REQ_WITHDRAW_PREFIX,
	)
	time.Sleep(50 * time.Second)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.BSCHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(90 * time.Second)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingbscreq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(140 * time.Second)

	fmt.Println("------------ step 5: withdrawing pETH from Incognito to ETH --------------")
	withdrawingPETH := big.NewInt(0).Div(etherTraded, big.NewInt(1000000000))
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncEtherTokenIDStr,
		withdrawingPETH,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningbscrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(140 * time.Second)

	tradingSuite.submitBurnProofForWithdrawal(
		burningTxID.(string),
		"getbscburnproof",
		tradingSuite.VaultBSCAddr,
		tradingSuite.BSCClient,
		tradingSuite.ChainIDBSC,
	)

	bal := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		tradingSuite.BSCClient,
	)
	fmt.Println("BNB balance after step 3: ", bal)
}
