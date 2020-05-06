package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/incognitochain/bridge-eth/bridge/compoundAgent"
	"github.com/incognitochain/bridge-eth/bridge/zrxtrade"
	"github.com/pkg/errors"

	"github.com/incognitochain/bridge-eth/bridge/compoundLogic"

	"github.com/incognitochain/bridge-eth/bridge/vault"

	"github.com/incognitochain/bridge-eth/bridge/compound"

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
type CompoundTradingTestSuite struct {
	*TradingTestSuite

	CompoundDeployedAddr common.Address
	ZRXTradeDeployedAddr common.Address
	Quote0xUrl           string

	InccSAITokenIDStr string
	InccDAITokenIDStr string
	InccETHTokenIDStr string

	cSAIAddressStr string
	cDAIAddressStr string
	cETHAddressStr string
	cREPAddressStr string

	// token amounts for tests
	DepositingEther       float64
	KBNBalanceAfterStep1  *big.Int
	SALTBalanceAfterStep2 *big.Int
}

func NewCompoundTradingTestSuite(tradingTestSuite *TradingTestSuite) *CompoundTradingTestSuite {
	return &CompoundTradingTestSuite{
		TradingTestSuite: tradingTestSuite,
	}
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (tradingSuite *CompoundTradingTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")
	// Kovan env
	tradingSuite.InccDAITokenIDStr = "0000000000000000000000000000000000000000000000000000000000000061"
	tradingSuite.InccETHTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000052"
	tradingSuite.InccSAITokenIDStr = "0000000000000000000000000000000000000000000000000000000000000051"
	tradingSuite.cSAIAddressStr = "0x63c344bf8651222346dd870be254d4347c9359f7"                            // kovan
	tradingSuite.cDAIAddressStr = "0xe7bc397dbd069fc7d0109c0636d06888bb50668c"                            // kovan
	tradingSuite.cETHAddressStr = "0xf92fbe0d3c0dcdae407923b2ac17ec223b1084e4"                            // kovan
	tradingSuite.cREPAddressStr = "0xfd874be7e6733bdc6dca9c7cdd97c225ec235d39"                            // kovan
	tradingSuite.CompoundDeployedAddr = common.HexToAddress("0x2400eD5B1c9b61256e987f9731fFca987d3dB878") //kovan
	tradingSuite.ZRXTradeDeployedAddr = common.HexToAddress("0xEA087FdBBC526e2aB75eD3dE9197fE6f61C94fed") //kovan
	tradingSuite.Quote0xUrl = "https://kovan.api.0x.org/swap/v0/quote?sellToken=%v&buyToken=%v&sellAmount=%v"

	tradingSuite.DepositingEther = float64(0.1)
}

func (tradingSuite *CompoundTradingTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
	tradingSuite.ETHClient.Close()
}

func (tradingSuite *CompoundTradingTestSuite) SetupTest() {
	fmt.Println("Setting up the test...")
}

func (tradingSuite *CompoundTradingTestSuite) TearDownTest() {
	fmt.Println("Tearing down the test...")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestCompoundTradingTestSuite(t *testing.T) {
	fmt.Println("Starting entry point for compound test suite...")

	tradingSuite := new(TradingTestSuite)
	suite.Run(t, tradingSuite)

	kyberTradingSuite := NewCompoundTradingTestSuite(tradingSuite)
	suite.Run(t, kyberTradingSuite)

	fmt.Println("Finishing entry point for compound test suite...")
}

func (tradingSuite *CompoundTradingTestSuite) mintCoin(
	srcQty *big.Int,
	srcTokenIDStr string,
	destTokenIDStr string,
) {
	compounAgentLogicAbi, _ := abi.JSON(strings.NewReader(compoundLogic.CompoundLogicABI))

	srcToken := common.HexToAddress(srcTokenIDStr)
	destToken := common.HexToAddress(destTokenIDStr)
	inputAgent, _ := compounAgentLogicAbi.Pack("mint", destToken, srcQty)
	tradingSuite.callVault(srcToken, srcQty, destToken, inputAgent, "mint")
}

func (tradingSuite *CompoundTradingTestSuite) borrowCoin(
	collateralAmount *big.Int,
	borrowAmount *big.Int,
	cTokenBorrow string,
	srcTokenIDStr string,
	destTokenIDStr string,
	addCollaterals []common.Address,
) {
	compounAgentLogicAbi, _ := abi.JSON(strings.NewReader(compoundLogic.CompoundLogicABI))

	cBorrow := common.HexToAddress(cTokenBorrow)
	srcToken := common.HexToAddress(srcTokenIDStr)
	destToken := common.HexToAddress(destTokenIDStr)
	inputAgent, _ := compounAgentLogicAbi.Pack("borrow", cBorrow, borrowAmount, addCollaterals)
	tradingSuite.callVault(srcToken, collateralAmount, destToken, inputAgent, "borrow")
}

func (tradingSuite *CompoundTradingTestSuite) redeemCoin(
	amountSentFromVault *big.Int,
	amountRedeem *big.Int,
	srcTokenIDStr string,
	destTokenIDStr string,
	isUnderlyingToken bool,
	removecollateral common.Address,
) {
	compounAgentLogicAbi, _ := abi.JSON(strings.NewReader(compoundLogic.CompoundLogicABI))

	srcToken := common.HexToAddress(srcTokenIDStr)
	destToken := common.HexToAddress(destTokenIDStr)
	inputAgent, _ := compounAgentLogicAbi.Pack("redeem", srcToken, amountRedeem, isUnderlyingToken, removecollateral)
	tradingSuite.callVault(srcToken, amountSentFromVault, destToken, inputAgent, "redeem")
}

func (tradingSuite *CompoundTradingTestSuite) repayBorrow(
	srcQty *big.Int,
	srcTokenIDStr string,
	destTokenIDStr string,
) {
	compounAgentLogicAbi, _ := abi.JSON(strings.NewReader(compoundLogic.CompoundLogicABI))

	srcToken := common.HexToAddress(srcTokenIDStr)
	destToken := common.HexToAddress(destTokenIDStr)
	inputAgent, _ := compounAgentLogicAbi.Pack("repayBorrow", destToken, srcQty)
	tradingSuite.callVault(srcToken, srcQty, destToken, inputAgent, "repayBorrow")
}

func (tradingSuite *CompoundTradingTestSuite) liquidateBorrow(
	srcQty *big.Int,
	srcTokenIDStr string,
	destTokenIDStr string,
	borrower common.Address,
) {
	compounAgentLogicAbi, _ := abi.JSON(strings.NewReader(compoundLogic.CompoundLogicABI))

	srcToken := common.HexToAddress(srcTokenIDStr)
	destToken := common.HexToAddress(destTokenIDStr)
	inputAgent, _ := compounAgentLogicAbi.Pack("liquidateBorrow", destToken, borrower, srcQty, destToken)
	tradingSuite.callVault(srcToken, srcQty, destToken, inputAgent, "liquidateBorrow")
}

func (tradingSuite *CompoundTradingTestSuite) callVault(
	srcToken common.Address,
	srcQty *big.Int,
	destToken common.Address,
	inputAgent []byte,
	whocall string,
) {
	compounAgent, _ := abi.JSON(strings.NewReader(compoundAgent.CompoundAgentABI))
	inputAgent, _ = compounAgent.Pack("delegateCall", inputAgent)

	compounProxyAbi, _ := abi.JSON(strings.NewReader(compound.CompoundABI))
	timestamp := []byte(randomizeTimestamp())
	tempData := append(srcToken.Bytes(), inputAgent...)
	tempData1 := append(tempData, timestamp...)
	data := rawsha3(tempData1)
	// sign for proxy verification
	signBytes, _ := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)
	inputProxy, _ := compounProxyAbi.Pack("execute", srcToken, srcQty, inputAgent, timestamp, signBytes)
	tempData = append(tradingSuite.CompoundDeployedAddr[:], inputProxy...)
	tempData1 = append(tempData, timestamp...)
	tempData2 := append(tempData1, common.LeftPadBytes(srcQty.Bytes(), 32)...)
	data = rawsha3(tempData2)
	// sign for vault verification
	signBytes, _ = crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)
	// // Get contract instance
	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)
	auth := bind.NewKeyedTransactor(tradingSuite.ETHPrivKey)
	auth.GasPrice = big.NewInt(10000000000)
	// auth.GasLimit = 2500000
	tx, err := c.Execute(
		auth,
		srcToken,
		srcQty,
		destToken,
		tradingSuite.CompoundDeployedAddr,
		inputProxy,
		timestamp,
		signBytes,
	)
	require.Equal(tradingSuite.T(), nil, err)
	txHash := tx.Hash()
	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("%v token executed , txHash: %x\n", whocall, txHash[:])
}

func (tradingSuite *CompoundTradingTestSuite) executeWith0xCompound(
	srcQty *big.Int,
	srcTokenName string,
	srcTokenIDStr string,
	destTokenName string,
	destTokenIDStr string,
) {
	tradeAbi, _ := abi.JSON(strings.NewReader(zrxtrade.ZrxtradeABI))

	// Get contract instance
	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)
	auth := bind.NewKeyedTransactor(tradingSuite.ETHPrivKey)
	auth.GasPrice = big.NewInt(50000000000)
	// quote
	srcToken := common.HexToAddress(srcTokenIDStr)
	destToken := common.HexToAddress(destTokenIDStr)

	quoteData, _ := quote0xCompound(tradingSuite.Quote0xUrl, srcTokenName, destTokenName, srcQty)
	forwarder := common.HexToAddress(quoteData["to"].(string))
	dt := common.Hex2Bytes(quoteData["data"].(string)[2:])
	auth.Value, _ = big.NewInt(0).SetString(quoteData["protocolFee"].(string), 10)
	auth.GasPrice, _ = big.NewInt(0).SetString(quoteData["gasPrice"].(string), 10)
	input, _ := tradeAbi.Pack("trade", srcToken, srcQty, destToken, dt, forwarder)
	timestamp := []byte(randomizeTimestamp())
	tempData := append(tradingSuite.ZRXTradeDeployedAddr[:], input...)
	tempData1 := append(tempData, timestamp...)
	tempData2 := append(tempData1, common.LeftPadBytes(srcQty.Bytes(), 32)...)
	data := rawsha3(tempData2)
	signBytes, _ := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)

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
}

func quote0xCompound(
	quote0xUrl string,
	srcToken, destToken string,
	srcQty *big.Int,
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
		return nil, errors.New("Request returns with fucking error!!!")
	}
	if bodyBytes, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(bodyBytes, &result); err != nil {
		return nil, err
	}
	return result.(map[string]interface{}), nil
}

func (tradingSuite *CompoundTradingTestSuite) executeMultiCollateralCompound(
	srcQties []*big.Int,
	srcTokenIDStrs []string,
	destTokenIDStrs []string,
	borrowAmount *big.Int,
	cTokenBorrow string,
	addCollaterals []common.Address,
) {
	compounAgentLogicAbi, _ := abi.JSON(strings.NewReader(compoundLogic.CompoundLogicABI))
	compounAgent, _ := abi.JSON(strings.NewReader(compoundAgent.CompoundAgentABI))
	compounProxyAbi, _ := abi.JSON(strings.NewReader(compound.CompoundABI))
	cBorrow := common.HexToAddress(cTokenBorrow)
	inputAgent, _ := compounAgentLogicAbi.Pack("borrowByMultiCollateral", cBorrow, borrowAmount, addCollaterals)
	inputAgent, _ = compounAgent.Pack("delegateCall", inputAgent)
	auth := bind.NewKeyedTransactor(tradingSuite.ETHPrivKey)
	auth.GasPrice = big.NewInt(50000000000)
	// auth.GasLimit = 2500000
	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)

	sourceAddresses := make([]common.Address, 0)
	for _, p := range srcTokenIDStrs {
		sourceAddresses = append(sourceAddresses, common.HexToAddress(p))
	}
	destAddresses := make([]common.Address, 0)
	for _, p := range destTokenIDStrs {
		destAddresses = append(destAddresses, common.HexToAddress(p))
	}
	amounts := make([]byte, 0)
	for i := range srcQties {
		amounts = append(amounts, common.LeftPadBytes(srcQties[i].Bytes(), 32)...)
	}

	timestamp := []byte(randomizeTimestamp())
	tempData1 := append(inputAgent, timestamp...)
	data := rawsha3(tempData1)
	// // sign for proxy verification
	signBytes, _ := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)
	// function executeMulti(
	//     IERC20[] memory srcTokens,
	//     uint[] memory amounts,
	//     bytes memory callData,
	//     bytes memory timestamp,
	//     bytes memory signData
	inputProxy, _ := compounProxyAbi.Pack("executeMulti", sourceAddresses, srcQties, inputAgent, timestamp, signBytes)
	tempData := append(tradingSuite.CompoundDeployedAddr[:], inputProxy...)
	tempData1 = append(tempData, timestamp...)
	tempData2 := append(tempData1, amounts...)
	data = rawsha3(tempData2)
	signBytes, _ = crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)

	tx, err := c.ExecuteMulti(
		auth,
		sourceAddresses,
		srcQties,
		destAddresses,
		tradingSuite.CompoundDeployedAddr,
		inputProxy,
		timestamp,
		signBytes,
	)
	require.Equal(tradingSuite.T(), nil, err)
	txHash := tx.Hash()
	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("Multi collaterals executed , txHash: %x\n", txHash[:])
}

func (tradingSuite *CompoundTradingTestSuite) getAgent(userAddress common.Address) (common.Address, error) {
	compound, err := compound.NewCompound(tradingSuite.CompoundDeployedAddr, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)
	return compound.GetAgentAddress(nil, userAddress)
}

func (tradingSuite *CompoundTradingTestSuite) Test1MintBorrowRedeemCETHbyCompound() {
	fmt.Println("============ TEST COMPOUND ETHER FOR CETH WITH COMPOUND PROXY ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
	burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("------------ STEP 1: porting ETH to pETH --------------")
	txHash := tradingSuite.depositETH(
		tradingSuite.DepositingEther,
		tradingSuite.IncPaymentAddrStr,
	)
	// time.Sleep(15 * time.Second)
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 40s for 15 blocks confirmation")
	time.Sleep(40 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
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

	tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))
	// require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPETH, big.NewInt(1000000000)), deposited)

	deposited := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("deposited EHT: ", deposited, big.NewInt(0).Div(tradeAmount, big.NewInt(int64(3))))

	fmt.Println("------------ step 3: mint cETH through Compound  --------------")
	tradingSuite.mintCoin(
		big.NewInt(0).Div(tradeAmount, big.NewInt(int64(3))),
		tradingSuite.EtherAddressStr,
		tradingSuite.cETHAddressStr,
	)

	deposited = tradingSuite.getDepositedBalance(
		tradingSuite.cETHAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("After: minted cEHT: ", deposited)

	fmt.Println("------------ step 4: collateral cETH and borrow DAI through Compound  --------------")
	tradingSuite.borrowCoin(
		big.NewInt(0).Div(deposited, big.NewInt(int64(3))),
		big.NewInt(10000000),
		tradingSuite.cDAIAddressStr,
		tradingSuite.cETHAddressStr,
		tradingSuite.DAIAddressStr,
		[]common.Address{common.HexToAddress(tradingSuite.cETHAddressStr)},
	)

	DAIBorrowed := tradingSuite.getDepositedBalance(
		tradingSuite.DAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("DAI borrowed: ", DAIBorrowed)

	fmt.Println("------------ step 5: Redeem cETH through Compound  --------------")
	tradingSuite.redeemCoin(
		big.NewInt(0).Div(deposited, big.NewInt(int64(3))),
		big.NewInt(0).Div(deposited, big.NewInt(int64(3))),
		tradingSuite.cETHAddressStr,
		tradingSuite.EtherAddressStr,
		false,
		common.Address{},
	)

	ETHredeem := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("ETH after call redeem: ", ETHredeem)

	fmt.Println("------------ step 6: Borrow ETH and Repay ETH through Compound  --------------")
	// borrow to repay
	tradingSuite.borrowCoin(
		big.NewInt(0).Div(deposited, big.NewInt(int64(3))),
		big.NewInt(10000000),
		tradingSuite.cETHAddressStr,
		tradingSuite.cETHAddressStr,
		tradingSuite.EtherAddressStr,
		[]common.Address{},
	)

	tradingSuite.repayBorrow(
		big.NewInt(10000000),
		tradingSuite.EtherAddressStr,
		tradingSuite.cETHAddressStr,
	)

	fmt.Println("------------ step 7: Liquidate Borrow through Compound  --------------")
}

func (tradingSuite *CompoundTradingTestSuite) Test2MintBorrowRedeemCDAIbyCompound() {
	fmt.Println("============ TEST COMPOUND DAI FOR CDAI WITH COMPOUND PROXY AGGREGATOR ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()

	fmt.Println("------------ step 1: execute trade ETH for DAI through 0x aggregator --------------")
	tradingSuite.executeWith0xCompound(
		big.NewInt(0).Div(tradeAmount, big.NewInt(int64(3))),
		"ETH",
		tradingSuite.EtherAddressStr,
		"DAI",
		tradingSuite.DAIAddressStr,
	)
	daiTraded := tradingSuite.getDepositedBalance(
		tradingSuite.DAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("daiTraded: ", daiTraded)

	fmt.Println("------------ step 2: mint cDAI through Compound  --------------")
	tradingSuite.mintCoin(
		big.NewInt(0).Div(daiTraded, big.NewInt(int64(3))),
		tradingSuite.DAIAddressStr,
		tradingSuite.cDAIAddressStr,
	)

	deposited := tradingSuite.getDepositedBalance(
		tradingSuite.cDAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("After: minted cDAI: ", deposited)

	fmt.Println("------------ step 3: collateral cDAI and borrow SAI through Compound  --------------")
	tradingSuite.borrowCoin(
		big.NewInt(0).Div(deposited, big.NewInt(int64(3))),
		big.NewInt(100000000),
		tradingSuite.cSAIAddressStr,
		tradingSuite.cDAIAddressStr,
		tradingSuite.SAIAddressStr,
		[]common.Address{common.HexToAddress(tradingSuite.cDAIAddressStr)},
	)

	SAIBorrowed := tradingSuite.getDepositedBalance(
		tradingSuite.SAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("SAI borrowed: ", SAIBorrowed)

	fmt.Println("------------ step 4: Redeem cDAI through Compound --------------")
	tradingSuite.redeemCoin(
		big.NewInt(0).Div(deposited, big.NewInt(int64(3))),
		big.NewInt(0).Div(deposited, big.NewInt(int64(3))),
		tradingSuite.cDAIAddressStr,
		tradingSuite.DAIAddressStr,
		false,
		common.Address{},
	)

	DAIredeem := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("DAI after call redeem: ", DAIredeem)

	fmt.Println("------------ step 5: Borrow DAI and Repay DAI through Compound aggregator --------------")
	// borrow to repay
	tradingSuite.borrowCoin(
		big.NewInt(0).Div(deposited, big.NewInt(int64(3))),
		big.NewInt(10000000),
		tradingSuite.cDAIAddressStr,
		tradingSuite.cDAIAddressStr,
		tradingSuite.DAIAddressStr,
		[]common.Address{},
	)

	tradingSuite.repayBorrow(
		big.NewInt(10000000),
		tradingSuite.DAIAddressStr,
		tradingSuite.cDAIAddressStr,
	)

	fmt.Println("------------ step 6: Collateral both cETH and cDAI to Borrow SAI through Compound aggregator --------------")
	depositedETH := tradingSuite.getDepositedBalance(
		tradingSuite.cETHAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("ETH deposited: ", depositedETH)
	tradingSuite.executeMultiCollateralCompound(
		[]*big.Int{big.NewInt(0).Div(depositedETH, big.NewInt(int64(3))), big.NewInt(0).Div(deposited, big.NewInt(int64(3)))},
		[]string{tradingSuite.cETHAddressStr, tradingSuite.cDAIAddressStr},
		[]string{tradingSuite.SAIAddressStr},
		big.NewInt(10000000),
		tradingSuite.cSAIAddressStr,
		[]common.Address{common.HexToAddress(tradingSuite.cETHAddressStr), common.HexToAddress(tradingSuite.cDAIAddressStr)},
	)

	SAIBorrowed = tradingSuite.getDepositedBalance(
		tradingSuite.SAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("SAI borrowed after execute multi collateral: ", SAIBorrowed)
	fmt.Println("------------ step 7: Liquidate Borrow DAI through Compound  --------------")
}
