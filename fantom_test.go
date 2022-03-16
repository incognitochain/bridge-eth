package bridge

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type FantomTestSuite struct {
	*TradingTestSuite

	// UniswapDeployedAddr      common.Address
	// UniswapRouteContractAddr common.Address
	// UNiswapQuoteContractAddr common.Address

	IncBUSDTokenIDStr string
	WMATICAddr        common.Address
	WETHAddr          common.Address

	// token amounts for tests
	DepositingFantom  float64
	WithdrawingFantom float64
}

func NewFantomTestSuite(tradingTestSuite *TradingTestSuite) *FantomTestSuite {
	return &FantomTestSuite{
		TradingTestSuite: tradingTestSuite,
	}
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (fantomSuite *FantomTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")
	// Fantom testnet env
	fantomSuite.IncBUSDTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000062"
	// fantomSuite.UniswapDeployedAddr = common.HexToAddress("0x3Aa036cF80684332aC40CcF64371f6a356570e63")
	// fantomSuite.UniswapRouteContractAddr = common.HexToAddress("0x68b3465833fb72a70ecdf485e0e4c7bd8665fc45")
	// fantomSuite.UNiswapQuoteContractAddr = common.HexToAddress("0x61ffe014ba17989e743c5f6cb21bf9697530b21e")

	// tokens
	fantomSuite.WMATICAddr = common.HexToAddress("0x9c3c9283d3e44854697cd22d3faa240cfb032889")
	fantomSuite.WETHAddr = common.HexToAddress("0xa6fa4fb5f76172d178d61b04b0ecd319c5d1c0aa")

	fantomSuite.DepositingFantom = float64(10)
	fantomSuite.WithdrawingFantom = float64(1)
}

func (fantomSuite *FantomTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
	fantomSuite.FTMClient.Close()
}

func (fantomSuite *FantomTestSuite) SetupTest() {
	fmt.Println("Setting up the test...")
}

func (fantomSuite *FantomTestSuite) TearDownTest() {
	fmt.Println("Tearing down the test...")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestFantomTestSuite(t *testing.T) {
	fmt.Println("Starting entry point for Fantom bridge test suite...")

	tradingSuite := new(TradingTestSuite)
	suite.Run(t, tradingSuite)

	fantomTestSuite := NewFantomTestSuite(tradingSuite)
	suite.Run(t, fantomTestSuite)

	fmt.Println("Finishing entry point for 0x test suite...")
}

// func (tradingSuite *FantomTestSuite) getExpectedAmount(
// 	srcQty *big.Int,
// 	paths []common.Address,
// 	fees []int64,
// ) *big.Int {
// 	c, err := pUniswapHelper.NewPUniswapHelper(tradingSuite.UNiswapQuoteContractAddr, tradingSuite.PLGClient)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	var amountOut *big.Int
// 	var amountIn *big.Int
// 	if len(fees) > 1 {
// 		inputParam := &pUniswapHelper.IUinswpaHelperExactInputParams{
// 			Path:     tradingSuite.buildPath(paths, fees),
// 			AmountIn: srcQty,
// 		}
// 		result, err := c.QuoteExactInput(nil, inputParam.Path, inputParam.AmountOutMinimum)
// 		require.Equal(tradingSuite.T(), nil, err)
// 		amountIn = inputParam.AmountIn
// 		amountOut = result.AmountOut
// 	} else {
// 		inputSingleParam := pUniswapHelper.IUinswpaHelperQuoteExactInputSingleParams{
// 			TokenIn:           paths[0],
// 			TokenOut:          paths[len(paths)-1],
// 			Fee:               big.NewInt(fees[0]),
// 			AmountIn:          srcQty,
// 			SqrtPriceLimitX96: big.NewInt(0),
// 		}
// 		result, err := c.QuoteExactInputSingle(nil, inputSingleParam)
// 		require.Equal(tradingSuite.T(), nil, err)
// 		amountIn = inputSingleParam.AmountIn
// 		amountOut = result.AmountOut
// 	}
// 	fmt.Printf("intput value: %v\n", amountIn.String())
// 	fmt.Printf("output value: %v\n", amountOut.String())
// 	return amountOut
// }
//
// const MAX_PERCENT = 10000
//
// func (tradingSuite *FantomTestSuite) executeWithPUniswapMultiTrade(
// 	srcQty *big.Int,
// 	paths [][]common.Address,
// 	fees [][]int64,
// 	percents []int64,
// 	isNative bool,
// 	deadline int64,
// 	isTestMultiPath []bool,
// ) {
// 	require.Equal(tradingSuite.T(), true, len(fees) != 0)
// 	require.Equal(tradingSuite.T(), len(paths), len(fees))
// 	require.Equal(tradingSuite.T(), len(percents), len(fees))
//
// 	tradeAbi1, err := abi.JSON(strings.NewReader(pUniswapHelper.PUniswapHelperMetaData.ABI))
// 	require.Equal(tradingSuite.T(), nil, err)
// 	tradeAbi, err := abi.JSON(strings.NewReader(puniswap.PuniswapMetaData.ABI))
// 	require.Equal(tradingSuite.T(), nil, err)
//
// 	// Get contract instance
// 	c, err := vault.NewVault(tradingSuite.VaultPLGAddr, tradingSuite.PLGClient)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	auth, err := bind.NewKeyedTransactorWithChainID(tradingSuite.ETHPrivKey, big.NewInt(int64(tradingSuite.ChainIDPLG)))
// 	require.Equal(tradingSuite.T(), nil, err)
// 	auth.GasPrice = big.NewInt(3e10)
// 	var calldata [][]byte
// 	for i := 0; i < len(fees); i++ {
// 		var agr interface{}
// 		amount := big.NewInt(0).Div(big.NewInt(0).Mul(srcQty, big.NewInt(percents[i])), big.NewInt(MAX_PERCENT))
// 		expectOutputAmount := tradingSuite.getExpectedAmount(
// 			amount,
// 			paths[i],
// 			fees[i],
// 		)
// 		recipient := tradingSuite.VaultPLGAddr
// 		if isNative && bytes.Compare(paths[i][len(paths[i])-1].Bytes(), tradingSuite.WMATICAddr.Bytes()) == 0 {
// 			recipient = tradingSuite.UniswapDeployedAddr
// 		}
// 		var input []byte
// 		if len(fees[i]) > 1 || isTestMultiPath[i] {
// 			agr = &pUniswapHelper.IUinswpaHelperExactInputParams{
// 				Path:             tradingSuite.buildPath(paths[i], fees[i]),
// 				Recipient:        recipient,
// 				AmountIn:         amount,
// 				AmountOutMinimum: expectOutputAmount,
// 			}
// 			input, err = tradeAbi1.Pack("exactInput", agr)
// 			require.Equal(tradingSuite.T(), nil, err)
// 		} else {
// 			agr = &pUniswapHelper.IUinswpaHelperExactInputSingleParams{
// 				TokenIn:           paths[i][0],
// 				TokenOut:          paths[i][len(paths[i])-1],
// 				Fee:               big.NewInt(fees[i][0]),
// 				Recipient:         recipient,
// 				AmountIn:          amount,
// 				SqrtPriceLimitX96: big.NewInt(0),
// 				AmountOutMinimum:  expectOutputAmount,
// 			}
// 			input, err = tradeAbi1.Pack("exactInputSingle", agr)
// 			require.Equal(tradingSuite.T(), nil, err)
// 		}
// 		calldata = append(calldata, input)
// 	}
// 	input, err := tradeAbi.Pack("multiTrades", big.NewInt(deadline), calldata, paths[0][0], paths[0][len(paths[0])-1], srcQty, isNative)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	timestamp := []byte(randomizeTimestamp())
// 	vaultAbi, err := abi.JSON(strings.NewReader(vault.VaultHelperABI))
// 	require.Equal(tradingSuite.T(), nil, err)
// 	sourceToken := paths[0][0]
// 	// todo: compare pTokenID
// 	if paths[0][0].String() == tradingSuite.WMATICAddr.String() {
// 		sourceToken = common.HexToAddress(tradingSuite.EtherAddressStr)
// 	}
// 	destToken := paths[0][len(paths[0])-1]
// 	if paths[0][len(paths[0])-1].String() == tradingSuite.WMATICAddr.String() && isNative {
// 		destToken = common.HexToAddress(tradingSuite.EtherAddressStr)
// 	}
// 	psData := vault.VaultHelperPreSignData{
// 		Prefix:    PLG_EXECUTE_PREFIX,
// 		Token:     sourceToken,
// 		Timestamp: timestamp,
// 		Amount:    srcQty,
// 	}
// 	tempData, err := vaultAbi.Pack("_buildSignExecute", psData, destToken, tradingSuite.UniswapDeployedAddr, input)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	data := rawsha3(tempData[4:])
// 	signBytes, err := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)
// 	require.Equal(tradingSuite.T(), nil, err)
//
// 	tx, err := c.Execute(
// 		auth,
// 		sourceToken,
// 		srcQty,
// 		destToken,
// 		tradingSuite.UniswapDeployedAddr,
// 		input,
// 		timestamp,
// 		signBytes,
// 	)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	txHash := tx.Hash()
// 	if err := wait(tradingSuite.PLGClient, txHash); err != nil {
// 		require.Equal(tradingSuite.T(), nil, err)
// 	}
// 	fmt.Printf("pUniswap trade executed , txHash: %x\n", txHash[:])
// }
//
// func (tradingSuite *FantomTestSuite) executeWithPUniswap(
// 	srcQty *big.Int,
// 	paths []common.Address,
// 	fees []int64,
// 	isNative bool,
// 	isTestMultiPath bool,
// ) {
// 	require.Equal(tradingSuite.T(), true, len(fees) != 0)
// 	require.Equal(tradingSuite.T(), len(paths), len(fees)+1)
//
// 	tradeAbi, err := abi.JSON(strings.NewReader(puniswap.PuniswapMetaData.ABI))
// 	require.Equal(tradingSuite.T(), nil, err)
//
// 	// Get contract instance
// 	c, err := vault.NewVault(tradingSuite.VaultPLGAddr, tradingSuite.PLGClient)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	auth, err := bind.NewKeyedTransactorWithChainID(tradingSuite.ETHPrivKey, big.NewInt(int64(tradingSuite.ChainIDPLG)))
// 	require.Equal(tradingSuite.T(), nil, err)
// 	auth.GasPrice = big.NewInt(3e10)
// 	var agr interface{}
// 	expectOutputAmount := tradingSuite.getExpectedAmount(
// 		srcQty,
// 		paths,
// 		fees,
// 	)
// 	recipient := tradingSuite.VaultPLGAddr
// 	if isNative && bytes.Compare(paths[len(paths)-1].Bytes(), tradingSuite.WMATICAddr.Bytes()) == 0 {
// 		recipient = tradingSuite.UniswapDeployedAddr
// 	}
// 	var input []byte
// 	if len(fees) > 1 || isTestMultiPath {
// 		agr = &pUniswapHelper.IUinswpaHelperExactInputParams{
// 			Path:             tradingSuite.buildPath(paths, fees),
// 			Recipient:        recipient,
// 			AmountIn:         srcQty,
// 			AmountOutMinimum: expectOutputAmount,
// 		}
// 		input, err = tradeAbi.Pack("tradeInput", agr, isNative)
// 	} else {
// 		agr = &pUniswapHelper.IUinswpaHelperExactInputSingleParams{
// 			TokenIn:           paths[0],
// 			TokenOut:          paths[len(paths)-1],
// 			Fee:               big.NewInt(fees[0]),
// 			Recipient:         recipient,
// 			AmountIn:          srcQty,
// 			SqrtPriceLimitX96: big.NewInt(0),
// 			AmountOutMinimum:  expectOutputAmount,
// 		}
// 		input, err = tradeAbi.Pack("tradeInputSingle", agr, isNative)
// 	}
// 	require.Equal(tradingSuite.T(), nil, err)
// 	timestamp := []byte(randomizeTimestamp())
// 	vaultAbi, err := abi.JSON(strings.NewReader(vault.VaultHelperABI))
// 	require.Equal(tradingSuite.T(), nil, err)
// 	sourceToken := paths[0]
// 	// todo: compare pTokenID
// 	if paths[0].String() == tradingSuite.WMATICAddr.String() {
// 		sourceToken = common.HexToAddress(tradingSuite.EtherAddressStr)
// 	}
// 	destToken := paths[len(paths)-1]
// 	if paths[len(paths)-1].String() == tradingSuite.WMATICAddr.String() && isNative {
// 		destToken = common.HexToAddress(tradingSuite.EtherAddressStr)
// 	}
// 	psData := vault.VaultHelperPreSignData{
// 		Prefix:    PLG_EXECUTE_PREFIX,
// 		Token:     sourceToken,
// 		Timestamp: timestamp,
// 		Amount:    srcQty,
// 	}
// 	tempData, err := vaultAbi.Pack("_buildSignExecute", psData, destToken, tradingSuite.UniswapDeployedAddr, input)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	data := rawsha3(tempData[4:])
// 	signBytes, err := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)
// 	require.Equal(tradingSuite.T(), nil, err)
//
// 	tx, err := c.Execute(
// 		auth,
// 		sourceToken,
// 		srcQty,
// 		destToken,
// 		tradingSuite.UniswapDeployedAddr,
// 		input,
// 		timestamp,
// 		signBytes,
// 	)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	txHash := tx.Hash()
// 	if err := wait(tradingSuite.PLGClient, txHash); err != nil {
// 		require.Equal(tradingSuite.T(), nil, err)
// 	}
// 	fmt.Printf("pUniswap trade executed , txHash: %x\n", txHash[:])
// }

func (fantomSuite *FantomTestSuite) Test1TradeEthForDAIWithPancake() {
	fmt.Println("============ TEST SHIELD UNSHIELD FANTOM ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	// depositAmount := big.NewInt(int64(fantomSuite.DepositingFantom * params.Ether))
	withdrawAmount := big.NewInt(int64(fantomSuite.WithdrawingFantom * params.Ether))
	withdrawingAmt := big.NewInt(0).Div(withdrawAmount, big.NewInt(1e9))

	pubKeyToAddrStr := crypto.PubkeyToAddress(fantomSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("------------ STEP 1: porting FTM to pFTM --------------")
	// create Fantom tx to send FTM to Vault FTM
	txHash := fantomSuite.depositETH(
		fantomSuite.DepositingFantom,
		fantomSuite.IncPaymentAddrStr,
		fantomSuite.VaultFTMAddr,
		fantomSuite.FTMClient,
	)
	// txHash := common.HexToHash("0x4e46b6144a84e55b9f8b5db3ed58f2971797958f00b7bd858eebb6c05d76f73e")
	time.Sleep(15 * time.Second)
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(fantomSuite.FTMHost, txHash)
	require.Equal(fantomSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 3s for 1 block confirmation")
	time.Sleep(3 * time.Second)
	_, err = fantomSuite.callIssuingETHReq(
		fantomSuite.IncFTMTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingfantomreq",
	)
	require.Equal(fantomSuite.T(), nil, err)
	time.Sleep(120 * time.Second)

	fmt.Println("------------ STEP 2: burning pFTM to deposit FTM to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := fantomSuite.callBurningPToken(
		fantomSuite.IncFTMTokenIDStr,
		withdrawingAmt,
		pubKeyToAddrStr[2:],
		"createandsendburningfantomfordeposittoscrequest",
	)
	require.Equal(fantomSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(fantomSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	fantomSuite.submitBurnProofForDepositToSC(
		burningTxID.(string),
		big.NewInt(int64(fantomSuite.ChainIDFTM)),
		"getburnftmprooffordeposittosc",
		fantomSuite.VaultFTMAddr,
		fantomSuite.FTMClient,
	)
	deposited := fantomSuite.getDepositedBalanceFTM(
		common.HexToAddress(fantomSuite.EtherAddressStr),
		pubKeyToAddrStr,
	)
	fmt.Printf("address own asset %v \n", pubKeyToAddrStr)
	//
	// fmt.Println("------------ step 3: execute trade MATIC for USDT through Pancake --------------")
	// fantomSuite.executeWithPUniswap(
	// 	deposited,
	// 	[]common.Address{fantomSuite.WMATICAddr, fantomSuite.WETHAddr},
	// 	[]int64{LOW},
	// 	false,
	// 	false,
	// )
	// time.Sleep(15 * time.Second)
	// daiTraded := fantomSuite.getDepositedBalancePLG(
	// 	fantomSuite.WETHAddr,
	// 	pubKeyToAddrStr,
	// )
	//
	// testCrossPoolTrade := big.NewInt(0).Div(daiTraded, big.NewInt(4))
	// fantomSuite.executeWithPUniswap(
	// 	testCrossPoolTrade,
	// 	[]common.Address{fantomSuite.WETHAddr, fantomSuite.WMATICAddr},
	// 	[]int64{LOW},
	// 	true,
	// 	false,
	// )
	//
	// fantomSuite.executeWithPUniswap(
	// 	testCrossPoolTrade,
	// 	[]common.Address{fantomSuite.WETHAddr, fantomSuite.WMATICAddr},
	// 	[]int64{LOW},
	// 	false,
	// 	true,
	// )
	//
	// fantomSuite.executeWithPUniswapMultiTrade(
	// 	testCrossPoolTrade,
	// 	[][]common.Address{{fantomSuite.WETHAddr, fantomSuite.WMATICAddr}, {fantomSuite.WETHAddr, fantomSuite.WMATICAddr}},
	// 	[][]int64{{LOW}, {MEDIUM}},
	// 	[]int64{70, 30},
	// 	false,
	// 	time.Now().Unix()+60000,
	// 	[]bool{false, false},
	// )
	//
	// daiTraded = fantomSuite.getDepositedBalancePLG(
	// 	fantomSuite.WETHAddr,
	// 	pubKeyToAddrStr,
	// )
	//
	// fmt.Println("weth: ", daiTraded)
	//
	fmt.Println("------------ step 3: withdrawing WETH from SC to pWETH on Incognito --------------")
	txHashByEmittingWithdrawalReq := fantomSuite.requestWithdraw(
		fantomSuite.EtherAddressStr,
		deposited,
		fantomSuite.FTMClient,
		big.NewInt(int64(fantomSuite.ChainIDFTM)),
		fantomSuite.VaultFTMAddr,
		FTM_REQ_WITHDRAW_PREFIX,
	)
	time.Sleep(45 * time.Second)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(fantomSuite.FTMHost, txHashByEmittingWithdrawalReq)
	require.Equal(fantomSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(50 * time.Second)
	_, err = fantomSuite.callIssuingETHReq(
		fantomSuite.IncFTMTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingfantomreq",
	)
	require.Equal(fantomSuite.T(), nil, err)
	time.Sleep(120 * time.Second)

	fmt.Println("------------ step 4: withdrawing pFTM from Incognito to FTM --------------")
	burningRes, err = fantomSuite.callBurningPToken(
		fantomSuite.IncFTMTokenIDStr,
		withdrawingAmt,
		fantomSuite.ETHOwnerAddrStr,
		"createandsendburningfantomrequest",
	)
	require.Equal(fantomSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(fantomSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	fantomSuite.submitBurnProofForWithdrawal(
		burningTxID.(string),
		"getftmburnproof",
		fantomSuite.VaultFTMAddr,
		fantomSuite.FTMClient,
		fantomSuite.ChainIDFTM,
	)

	bal := fantomSuite.getBalanceOnETHNet(
		common.HexToAddress(fantomSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", fantomSuite.ETHOwnerAddrStr)),
		fantomSuite.FTMClient,
	)
	fmt.Println("FTM balance: ", bal)
}

// func (fantomSuite *FantomTestSuite) buildPath(paths []common.Address, fees []int64) []byte {
// 	var temp []byte
// 	for i := 0; i < len(fees); i++ {
// 		temp = append(temp, paths[i].Bytes()...)
// 		fee, err := hex.DecodeString(fmt.Sprintf("%06x", fees[i]))
// 		require.Equal(fantomSuite.T(), nil, err)
// 		temp = append(temp, fee...)
// 	}
// 	temp = append(temp, paths[len(paths)-1].Bytes()...)

// 	return temp
// }
