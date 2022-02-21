package bridge

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/params"
	pUniswapHelper "github.com/incognitochain/bridge-eth/bridge/puniswaphelper"
	puniswap "github.com/incognitochain/bridge-eth/bridge/puniswapproxy"
	"github.com/incognitochain/bridge-eth/bridge/vault"
	"github.com/stretchr/testify/suite"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type PolygonTestSuite struct {
	*TradingTestSuite

	UniswapDeployedAddr      common.Address
	UniswapRouteContractAddr common.Address
	UNiswapQuoteContractAddr common.Address

	IncBUSDTokenIDStr string
	WMATICAddr        common.Address
	WETHAddr          common.Address

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

const (
	LOW    = 500
	MEDIUM = 3000
	HIGH   = 10000
)

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (tradingSuite *PolygonTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")
	// Polygon testnet env
	tradingSuite.IncBUSDTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000062"
	tradingSuite.UniswapDeployedAddr = common.HexToAddress("0x3Aa036cF80684332aC40CcF64371f6a356570e63")
	tradingSuite.UniswapRouteContractAddr = common.HexToAddress("0x68b3465833fb72a70ecdf485e0e4c7bd8665fc45")
	tradingSuite.UNiswapQuoteContractAddr = common.HexToAddress("0x61ffe014ba17989e743c5f6cb21bf9697530b21e")

	// tokens
	tradingSuite.WMATICAddr = common.HexToAddress("0x9c3c9283d3e44854697cd22d3faa240cfb032889")
	tradingSuite.WETHAddr = common.HexToAddress("0xa6fa4fb5f76172d178d61b04b0ecd319c5d1c0aa")

	tradingSuite.DepositingEther = float64(0.2)
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
	srcQty *big.Int,
	paths []common.Address,
	fees []int64,
) *big.Int {
	c, err := pUniswapHelper.NewPUniswapHelper(tradingSuite.UNiswapQuoteContractAddr, tradingSuite.PLGClient)
	require.Equal(tradingSuite.T(), nil, err)
	var amountOut *big.Int
	var amountIn *big.Int
	if len(fees) > 1 {
		inputParam := &pUniswapHelper.IUinswpaHelperExactInputParams{
			Path:     tradingSuite.buildPath(paths, fees),
			AmountIn: srcQty,
		}
		result, err := c.QuoteExactInput(nil, inputParam.Path, inputParam.AmountIn)
		require.Equal(tradingSuite.T(), nil, err)
		amountIn = inputParam.AmountIn
		amountOut = result.AmountOut
	} else {
		inputSingleParam := pUniswapHelper.IUinswpaHelperQuoteExactInputSingleParams{
			TokenIn:           paths[0],
			TokenOut:          paths[len(paths)-1],
			Fee:               big.NewInt(fees[0]),
			AmountIn:          srcQty,
			SqrtPriceLimitX96: big.NewInt(0),
		}
		result, err := c.QuoteExactInputSingle(nil, inputSingleParam)
		require.Equal(tradingSuite.T(), nil, err)
		amountIn = inputSingleParam.AmountIn
		amountOut = result.AmountOut
	}
	fmt.Printf("intput value: %v\n", amountIn.String())
	fmt.Printf("output value: %v\n", amountOut.String())
	return amountOut
}

const MAX_PERCENT = 10000

func (tradingSuite *PolygonTestSuite) executeWithPUniswapMultiTrade(
	srcQty *big.Int,
	paths [][]common.Address,
	fees [][]int64,
	percents []int64,
	isNative bool,
	deadline int64,
	isTestMultiPath []bool,
) {
	require.Equal(tradingSuite.T(), true, len(fees) != 0)
	require.Equal(tradingSuite.T(), len(paths), len(fees))
	require.Equal(tradingSuite.T(), len(percents), len(fees))

	tradeAbi1, err := abi.JSON(strings.NewReader(pUniswapHelper.PUniswapHelperMetaData.ABI))
	require.Equal(tradingSuite.T(), nil, err)
	tradeAbi, err := abi.JSON(strings.NewReader(puniswap.PuniswapMetaData.ABI))
	require.Equal(tradingSuite.T(), nil, err)

	// Get contract instance
	c, err := vault.NewVault(tradingSuite.VaultPLGAddr, tradingSuite.PLGClient)
	require.Equal(tradingSuite.T(), nil, err)
	auth, err := bind.NewKeyedTransactorWithChainID(tradingSuite.ETHPrivKey, big.NewInt(int64(tradingSuite.ChainIDPLG)))
	require.Equal(tradingSuite.T(), nil, err)
	auth.GasPrice = big.NewInt(3e10)
	var calldata [][]byte
	for i := 0; i < len(fees); i++ {
		var agr interface{}
		amount := big.NewInt(0).Div(big.NewInt(0).Mul(srcQty, big.NewInt(percents[i])), big.NewInt(MAX_PERCENT))
		expectOutputAmount := tradingSuite.getExpectedAmount(
			amount,
			paths[i],
			fees[i],
		)
		recipient := tradingSuite.VaultPLGAddr
		if isNative && bytes.Compare(paths[i][len(paths[i])-1].Bytes(), tradingSuite.WMATICAddr.Bytes()) == 0 {
			recipient = tradingSuite.UniswapDeployedAddr
		}
		var input []byte
		if len(fees[i]) > 1 || isTestMultiPath[i] {
			agr = &pUniswapHelper.IUinswpaHelperExactInputParams{
				Path:             tradingSuite.buildPath(paths[i], fees[i]),
				Recipient:        recipient,
				AmountIn:         amount,
				AmountOutMinimum: expectOutputAmount,
			}
			input, err = tradeAbi1.Pack("exactInput", agr)
			require.Equal(tradingSuite.T(), nil, err)
		} else {
			agr = &pUniswapHelper.IUinswpaHelperExactInputSingleParams{
				TokenIn:           paths[i][0],
				TokenOut:          paths[i][len(paths[i])-1],
				Fee:               big.NewInt(fees[i][0]),
				Recipient:         recipient,
				AmountIn:          amount,
				SqrtPriceLimitX96: big.NewInt(0),
				AmountOutMinimum:  expectOutputAmount,
			}
			input, err = tradeAbi1.Pack("exactInputSingle", agr)
			require.Equal(tradingSuite.T(), nil, err)
		}
		calldata = append(calldata, input)
	}
	input, err := tradeAbi.Pack("multiTrades", big.NewInt(deadline), calldata, paths[0][0], paths[0][len(paths[0])-1], srcQty, isNative)
	require.Equal(tradingSuite.T(), nil, err)
	timestamp := []byte(randomizeTimestamp())
	vaultAbi, err := abi.JSON(strings.NewReader(vault.VaultHelperABI))
	require.Equal(tradingSuite.T(), nil, err)
	sourceToken := paths[0][0]
	// todo: compare pTokenID
	if paths[0][0].String() == tradingSuite.WMATICAddr.String() {
		sourceToken = common.HexToAddress(tradingSuite.EtherAddressStr)
	}
	destToken := paths[0][len(paths[0])-1]
	if paths[0][len(paths[0])-1].String() == tradingSuite.WMATICAddr.String() && isNative {
		destToken = common.HexToAddress(tradingSuite.EtherAddressStr)
	}
	psData := vault.VaultHelperPreSignData{
		Prefix:    PLG_EXECUTE_PREFIX,
		Token:     sourceToken,
		Timestamp: timestamp,
		Amount:    srcQty,
	}
	tempData, err := vaultAbi.Pack("_buildSignExecute", psData, destToken, tradingSuite.UniswapDeployedAddr, input)
	require.Equal(tradingSuite.T(), nil, err)
	data := rawsha3(tempData[4:])
	signBytes, err := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)
	require.Equal(tradingSuite.T(), nil, err)

	tx, err := c.Execute(
		auth,
		sourceToken,
		srcQty,
		destToken,
		tradingSuite.UniswapDeployedAddr,
		input,
		timestamp,
		signBytes,
	)
	require.Equal(tradingSuite.T(), nil, err)
	txHash := tx.Hash()
	if err := wait(tradingSuite.PLGClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("pUniswap trade executed , txHash: %x\n", txHash[:])
}

func (tradingSuite *PolygonTestSuite) executeWithPUniswap(
	srcQty *big.Int,
	paths []common.Address,
	fees []int64,
	isNative bool,
	isTestMultiPath bool,
) {
	require.Equal(tradingSuite.T(), true, len(fees) != 0)
	require.Equal(tradingSuite.T(), len(paths), len(fees)+1)

	tradeAbi, err := abi.JSON(strings.NewReader(puniswap.PuniswapMetaData.ABI))
	require.Equal(tradingSuite.T(), nil, err)

	// Get contract instance
	c, err := vault.NewVault(tradingSuite.VaultPLGAddr, tradingSuite.PLGClient)
	require.Equal(tradingSuite.T(), nil, err)
	auth, err := bind.NewKeyedTransactorWithChainID(tradingSuite.ETHPrivKey, big.NewInt(int64(tradingSuite.ChainIDPLG)))
	require.Equal(tradingSuite.T(), nil, err)
	auth.GasPrice = big.NewInt(3e10)
	var agr interface{}
	expectOutputAmount := tradingSuite.getExpectedAmount(
		srcQty,
		paths,
		fees,
	)
	recipient := tradingSuite.VaultPLGAddr
	if isNative && bytes.Compare(paths[len(paths)-1].Bytes(), tradingSuite.WMATICAddr.Bytes()) == 0 {
		recipient = tradingSuite.UniswapDeployedAddr
	}
	var input []byte
	if len(fees) > 1 || isTestMultiPath {
		agr = &pUniswapHelper.IUinswpaHelperExactInputParams{
			Path:             tradingSuite.buildPath(paths, fees),
			Recipient:        recipient,
			AmountIn:         srcQty,
			AmountOutMinimum: expectOutputAmount,
		}
		input, err = tradeAbi.Pack("tradeInput", agr, isNative)
	} else {
		agr = &pUniswapHelper.IUinswpaHelperExactInputSingleParams{
			TokenIn:           paths[0],
			TokenOut:          paths[len(paths)-1],
			Fee:               big.NewInt(fees[0]),
			Recipient:         recipient,
			AmountIn:          srcQty,
			SqrtPriceLimitX96: big.NewInt(0),
			AmountOutMinimum:  expectOutputAmount,
		}
		input, err = tradeAbi.Pack("tradeInputSingle", agr, isNative)
	}
	require.Equal(tradingSuite.T(), nil, err)
	timestamp := []byte(randomizeTimestamp())
	vaultAbi, err := abi.JSON(strings.NewReader(vault.VaultHelperABI))
	require.Equal(tradingSuite.T(), nil, err)
	sourceToken := paths[0]
	// todo: compare pTokenID
	if paths[0].String() == tradingSuite.WMATICAddr.String() {
		sourceToken = common.HexToAddress(tradingSuite.EtherAddressStr)
	}
	destToken := paths[len(paths)-1]
	if paths[len(paths)-1].String() == tradingSuite.WMATICAddr.String() && isNative {
		destToken = common.HexToAddress(tradingSuite.EtherAddressStr)
	}
	psData := vault.VaultHelperPreSignData{
		Prefix:    PLG_EXECUTE_PREFIX,
		Token:     sourceToken,
		Timestamp: timestamp,
		Amount:    srcQty,
	}
	tempData, err := vaultAbi.Pack("_buildSignExecute", psData, destToken, tradingSuite.UniswapDeployedAddr, input)
	require.Equal(tradingSuite.T(), nil, err)
	data := rawsha3(tempData[4:])
	signBytes, err := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)
	require.Equal(tradingSuite.T(), nil, err)

	tx, err := c.Execute(
		auth,
		sourceToken,
		srcQty,
		destToken,
		tradingSuite.UniswapDeployedAddr,
		input,
		timestamp,
		signBytes,
	)
	require.Equal(tradingSuite.T(), nil, err)
	txHash := tx.Hash()
	if err := wait(tradingSuite.PLGClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("pUniswap trade executed , txHash: %x\n", txHash[:])
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

	fmt.Println("------------ step 3: execute trade MATIC for USDT through Pancake --------------")
	tradingSuite.executeWithPUniswap(
		deposited,
		[]common.Address{tradingSuite.WMATICAddr, tradingSuite.WETHAddr},
		[]int64{LOW},
		false,
		false,
	)
	time.Sleep(15 * time.Second)
	daiTraded := tradingSuite.getDepositedBalancePLG(
		tradingSuite.WETHAddr,
		pubKeyToAddrStr,
	)

	testCrossPoolTrade := big.NewInt(0).Div(daiTraded, big.NewInt(4))
	tradingSuite.executeWithPUniswap(
		testCrossPoolTrade,
		[]common.Address{tradingSuite.WETHAddr, tradingSuite.WMATICAddr},
		[]int64{LOW},
		true,
		false,
	)

	tradingSuite.executeWithPUniswap(
		testCrossPoolTrade,
		[]common.Address{tradingSuite.WETHAddr, tradingSuite.WMATICAddr},
		[]int64{LOW},
		false,
		true,
	)

	tradingSuite.executeWithPUniswapMultiTrade(
		testCrossPoolTrade,
		[][]common.Address{{tradingSuite.WETHAddr, tradingSuite.WMATICAddr}, {tradingSuite.WETHAddr, tradingSuite.WMATICAddr}},
		[][]int64{{LOW}, {MEDIUM}},
		[]int64{70, 30},
		false,
		time.Now().Unix()+60000,
		[]bool{false, false},
	)

	daiTraded = tradingSuite.getDepositedBalancePLG(
		tradingSuite.WETHAddr,
		pubKeyToAddrStr,
	)

	fmt.Println("weth: ", daiTraded)

	fmt.Println("------------ step 3: withdrawing WETH from SC to pWETH on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.WETHAddr.String(),
		daiTraded,
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

	fmt.Println("------------ step 4: withdrawing pWETH from Incognito to WETH --------------")
	withdrawingPDAI := big.NewInt(0).Div(daiTraded, big.NewInt(1e9))
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
		tradingSuite.WETHAddr,
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		tradingSuite.PLGClient,
	)
	tradingSuite.DAIBalanceAfterStep1 = bal
	fmt.Println("WETH balance after step 1: ", tradingSuite.DAIBalanceAfterStep1)
}

func (tradingSuite *PolygonTestSuite) buildPath(paths []common.Address, fees []int64) []byte {
	var temp []byte
	for i := 0; i < len(fees); i++ {
		temp = append(temp, paths[i].Bytes()...)
		fee, err := hex.DecodeString(fmt.Sprintf("%06x", fees[i]))
		require.Equal(tradingSuite.T(), nil, err)
		temp = append(temp, fee...)
	}
	temp = append(temp, paths[len(paths)-1].Bytes()...)

	return temp
}
