package bridge

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/incognitochain/bridge-eth/bridge/incognito_proxy"
	"github.com/incognitochain/bridge-eth/bridge/pcurve"
	"github.com/incognitochain/bridge-eth/bridge/uniswap"
	"github.com/incognitochain/bridge-eth/bridge/vault"
	"github.com/incognitochain/bridge-eth/bridge/vaultplg"
	"github.com/incognitochain/bridge-eth/bridge/vaultproxy"
	"github.com/incognitochain/bridge-eth/erc20"
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
type PolygonCurveTestSuite struct {
	*TradingTestSuite

	PcurveDeployedAddr       common.Address
	UNiswapQuoteContractAddr common.Address

	IncBUSDTokenIDStr string
	WMATICAddr        common.Address
	MATICAddr         common.Address
	USDCAddress       common.Address
	USDTAddress       common.Address

	QuickSwap    common.Address
	PolygonPool1 common.Address
	PolygonPool2 common.Address
	c            *committees
	v            *vault.Vault
	// token amounts for tests
	DepositingEther      float64
	DAIBalanceAfterStep1 *big.Int
	MRKBalanceAfterStep2 *big.Int
}

func NewPolygonCurveTestSuite(tradingTestSuite *TradingTestSuite) *PolygonCurveTestSuite {
	return &PolygonCurveTestSuite{
		TradingTestSuite: tradingTestSuite,
	}
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (tradingSuite *PolygonCurveTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")
	// Polygon testnet env
	tradingSuite.IncBUSDTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000062"
	tradingSuite.UNiswapQuoteContractAddr = common.HexToAddress("0x61ffe014ba17989e743c5f6cb21bf9697530b21e")
	tradingSuite.PLGHost = "http://127.0.0.1:8545"
	tradingSuite.PolygonPool1 = common.HexToAddress("0x3FCD5De6A9fC8A99995c406c77DDa3eD7E406f81")
	tradingSuite.PolygonPool2 = common.HexToAddress("0x751B1e21756bDbc307CBcC5085c042a0e9AaEf36")
	//tradingSuite.VaultPLGAddr = common.HexToAddress("0x43D037A562099A4C2c95b1E2120cc43054450629")
	client, err := ethclient.Dial(tradingSuite.PLGHost)
	require.Equal(tradingSuite.T(), nil, err)
	tradingSuite.PLGClient = client
	require.Equal(tradingSuite.T(), nil, err)
	//cmd := exec.Command("ganache-cli --fork https://polygon-mainnet.infura.io/v3/9bc873177cf74a03a35739e45755a9ac@25006800 --secure -u 0x1c175357c46Fece0e4c1edC72E544a5278b07963 --account="0xaad53b70ad9ed01b75238533dd6b395f4d300427da0165aafbd42ea7a606601f,1000000000000000000000000" --networkId 137")
	//err = cmd.Run()
	//require.Equal(tradingSuite.T(), nil, err)

	// tokens
	tradingSuite.WMATICAddr = common.HexToAddress("0x0d500b1d8e8ef31e21c99d1db9a6444d3adf1270")
	tradingSuite.MATICAddr = common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE")
	tradingSuite.USDCAddress = common.HexToAddress("0x2791bca1f2de4661ed88a30c99a7a9449aa84174")
	tradingSuite.USDTAddress = common.HexToAddress("0xc2132d05d31c914a87c6611c10748aeb04b58e8f")
	tradingSuite.QuickSwap = common.HexToAddress("0xa5E0829CaCEd8fFDD4De3c43696c57F7D7A678ff")
	tradingSuite.DepositingEther = float64(10)
	tradingSuite.c = getFixedCommittee()
	tradingSuite.ChainIDPLG = 137

	auth, err = bind.NewKeyedTransactorWithChainID(tradingSuite.ETHPrivKey, big.NewInt(int64(tradingSuite.ChainIDPLG)))
	require.Equal(tradingSuite.T(), nil, err)
	auth.GasPrice = big.NewInt(10000000000)

	incAddrPLG, tx, _, err := incognito_proxy.DeployIncognitoProxy(auth, tradingSuite.PLGClient, tradingSuite.WMATICAddr, tradingSuite.c.beacons, tradingSuite.c.bridges)
	require.Equal(tradingSuite.T(), nil, err)

	// incAddr := common.HexToAddress(IncognitoProxyAddress)
	fmt.Println("deployed incognito_proxy")
	fmt.Printf("addr: %s\n", incAddrPLG.Hex())

	// Wait until tx is confirmed
	err = wait(tradingSuite.PLGClient, tx.Hash())
	require.Equal(tradingSuite.T(), nil, err)

	// Deploy vault
	vaultAddrPLG, tx, _, err := vaultplg.DeployVaultplg(auth, tradingSuite.PLGClient)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("deployed vault")
	fmt.Printf("addr: %s\n", vaultAddrPLG.Hex())

	// Wait until tx is confirmed
	err = wait(tradingSuite.PLGClient, tx.Hash())
	require.Equal(tradingSuite.T(), nil, err)

	prevVaultPLG := common.Address{}
	vaultAbi, _ := abi.JSON(strings.NewReader(vault.VaultABI))
	input, _ := vaultAbi.Pack("initialize", prevVaultPLG)
	tradingSuite.VaultPLGAddr, tx, _, err = vaultproxy.DeployTransparentUpgradeableProxy(auth, tradingSuite.PLGClient, vaultAddrPLG, common.HexToAddress("0x868C20050BDD1b0FA90C17b6329051A45f0031bc"), incAddrPLG, input)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("deployed vault proxy")
	fmt.Printf("addr: %s\n", tradingSuite.VaultPLGAddr.Hex())

	err = wait(tradingSuite.PLGClient, tx.Hash())
	require.Equal(tradingSuite.T(), nil, err)

	tradingSuite.PcurveDeployedAddr, tx, _, err = pcurve.DeployPcurve(auth, tradingSuite.PLGClient)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("deployed pcurve proxy")
	fmt.Printf("addr: %s\n", tradingSuite.PcurveDeployedAddr.Hex())

	err = wait(tradingSuite.PLGClient, tx.Hash())
	require.Equal(tradingSuite.T(), nil, err)

	tradingSuite.v, err = vault.NewVault(tradingSuite.VaultPLGAddr, tradingSuite.PLGClient)
	require.Equal(tradingSuite.T(), nil, err)
}

func (tradingSuite *PolygonCurveTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
	tradingSuite.PLGClient.Close()
}

func (tradingSuite *PolygonCurveTestSuite) SetupTest() {
	fmt.Println("Setting up the test...")
}

func (tradingSuite *PolygonCurveTestSuite) TearDownTest() {
	fmt.Println("Tearing down the test...")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestPolygonCurveTestSuite(t *testing.T) {
	fmt.Println("Starting entry point for Pcurve test suite...")

	tradingSuite := new(TradingTestSuite)
	suite.Run(t, tradingSuite)

	uniswapTradingSuite := NewPolygonCurveTestSuite(tradingSuite)
	suite.Run(t, uniswapTradingSuite)

	fmt.Println("Finishing entry point for 0x test suite...")
}

func (tradingSuite *PolygonCurveTestSuite) getExpectedAmount(
	srcQty *big.Int,
	i *big.Int,
	j *big.Int,
	curvePool common.Address,
) *big.Int {
	c, err := pcurve.NewPcurvehelper(curvePool, tradingSuite.PLGClient)
	require.Equal(tradingSuite.T(), nil, err)
	amountOut, err := c.GetDy(nil, i, j, srcQty)
	require.Equal(tradingSuite.T(), nil, err)

	fmt.Printf("intput value: %v\n", srcQty.String())
	fmt.Printf("output value: %v\n", amountOut.String())
	return amountOut
}

func (tradingSuite *PolygonCurveTestSuite) getExpectedUnderlyingAmount(
	srcQty *big.Int,
	i *big.Int,
	j *big.Int,
	curvePool common.Address,
) *big.Int {
	c, err := pcurve.NewPcurvehelper(curvePool, tradingSuite.PLGClient)
	require.Equal(tradingSuite.T(), nil, err)
	amountOut, err := c.GetDyUnderlying(nil, i, j, srcQty)
	require.Equal(tradingSuite.T(), nil, err)

	fmt.Printf("intput value: %v\n", srcQty.String())
	fmt.Printf("output underlying value: %v\n", amountOut.String())
	return amountOut
}

func (tradingSuite *PolygonCurveTestSuite) executeWithPcurve(
	srcQty *big.Int,
	i *big.Int,
	j *big.Int,
	curvePool common.Address,
	isUnderlying bool,
) {
	tradeAbi, err := abi.JSON(strings.NewReader(pcurve.PcurveMetaData.ABI))
	require.Equal(tradingSuite.T(), nil, err)

	// Get contract instance
	c, err := vault.NewVault(tradingSuite.VaultPLGAddr, tradingSuite.PLGClient)
	require.Equal(tradingSuite.T(), nil, err)
	auth, err := bind.NewKeyedTransactorWithChainID(tradingSuite.ETHPrivKey, big.NewInt(int64(tradingSuite.ChainIDPLG)))
	require.Equal(tradingSuite.T(), nil, err)
	auth.GasPrice = big.NewInt(3e10)
	expectOutputAmount := big.NewInt(0)
	destToken := common.Address{}
	sourceToken := common.Address{}
	cHelper, err := pcurve.NewPcurvehelper(curvePool, tradingSuite.PLGClient)
	require.Equal(tradingSuite.T(), nil, err)
	if isUnderlying {
		expectOutputAmount = tradingSuite.getExpectedUnderlyingAmount(
			srcQty,
			i,
			j,
			curvePool,
		)
		destToken, err = cHelper.UnderlyingCoins(nil, j)
		require.Equal(tradingSuite.T(), nil, err)
		sourceToken, err = cHelper.UnderlyingCoins(nil, i)
	} else {
		expectOutputAmount = tradingSuite.getExpectedAmount(
			srcQty,
			i,
			j,
			curvePool,
		)
		destToken, err = cHelper.Coins(nil, j)
		require.Equal(tradingSuite.T(), nil, err)
		sourceToken, err = cHelper.Coins(nil, i)
	}
	require.Equal(tradingSuite.T(), nil, err)

	recipient := destToken
	if isUnderlying && bytes.Compare(destToken.Bytes(), tradingSuite.MATICAddr.Bytes()) == 0 {
		recipient = common.HexToAddress(tradingSuite.EtherAddressStr)
	}
	var input []byte
	if isUnderlying {
		input, err = tradeAbi.Pack("exchange", i, j, srcQty, expectOutputAmount, curvePool)
	} else {
		input, err = tradeAbi.Pack("exchangeUnderlying", i, j, srcQty, expectOutputAmount, curvePool)
	}
	require.Equal(tradingSuite.T(), nil, err)

	timestamp := []byte(randomizeTimestamp())
	vaultAbi, err := abi.JSON(strings.NewReader(vault.VaultHelperABI))
	require.Equal(tradingSuite.T(), nil, err)
	// todo: compare pTokenID
	if bytes.Compare(sourceToken.Bytes(), tradingSuite.MATICAddr.Bytes()) == 0 {
		sourceToken = common.HexToAddress(tradingSuite.EtherAddressStr)
	}
	psData := vault.VaultHelperPreSignData{
		Prefix:    PLG_EXECUTE_PREFIX,
		Token:     sourceToken,
		Timestamp: timestamp,
		Amount:    srcQty,
	}
	tempData, err := vaultAbi.Pack("_buildSignExecute", psData, recipient, tradingSuite.PcurveDeployedAddr, input)
	require.Equal(tradingSuite.T(), nil, err)
	data := rawsha3(tempData[4:])
	signBytes, err := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)
	require.Equal(tradingSuite.T(), nil, err)

	tx, err := c.Execute(
		auth,
		sourceToken,
		srcQty,
		recipient,
		tradingSuite.PcurveDeployedAddr,
		input,
		timestamp,
		signBytes,
	)
	require.Equal(tradingSuite.T(), nil, err)
	txHash := tx.Hash()
	if err := wait(tradingSuite.PLGClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("Pcurve trade executed , txHash: %x\n", txHash[:])
}

func (tradingSuite *PolygonCurveTestSuite) Test1TradeEthForDAIWithPancake() {
	fmt.Println("============ TEST SHIELD UNSHIELD POLYGON ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC)
	// swap from matic to usdc
	quickSwap, err := uniswap.NewUniswapV2(tradingSuite.QuickSwap, tradingSuite.PLGClient)
	require.Equal(tradingSuite.T(), nil, err)
	auth, err = bind.NewKeyedTransactorWithChainID(tradingSuite.ETHPrivKey, big.NewInt(int64(tradingSuite.ChainIDPLG)))
	require.Equal(tradingSuite.T(), nil, err)
	value, ok := big.NewInt(0).SetString("10000000000000000000", 10)
	require.Equal(tradingSuite.T(), true, ok)
	auth.Value = value
	auth.GasPrice = big.NewInt(10000000000)
	auth.GasLimit = 500000
	tx, err := quickSwap.SwapExactETHForTokens(
		auth,
		big.NewInt(0),
		[]common.Address{tradingSuite.WMATICAddr, tradingSuite.USDCAddress},
		auth.From,
		big.NewInt(3299046487),
	)
	require.Equal(tradingSuite.T(), nil, err)
	err = wait(tradingSuite.PLGClient, tx.Hash())
	require.Equal(tradingSuite.T(), nil, err)
	auth.Value = big.NewInt(0)
	erc20Token, err := erc20.NewErc20(tradingSuite.USDCAddress, tradingSuite.PLGClient)
	require.Equal(tradingSuite.T(), nil, err)
	usdcBalance, err := erc20Token.BalanceOf(nil, auth.From)
	require.Equal(tradingSuite.T(), nil, err)
	require.Equal(tradingSuite.T(), 1, usdcBalance.Cmp(big.NewInt(0)))
	auth.GasLimit = 0
	// assumed the shielding flow is working
	// deposite erc20 token to vault
	tradingSuite.depositERC20ToBridge(
		usdcBalance,
		tradingSuite.USDCAddress,
		tradingSuite.IncPaymentAddrStr,
		tradingSuite.VaultPLGAddr,
		tradingSuite.PLGClient,
		tradingSuite.ChainIDPLG,
	)

	fmt.Println("------------ STEP 2: burning pBNB to deposit BNB to SC --------------")
	burnUSDCAmount := big.NewInt(0).Div(usdcBalance, big.NewInt(2))
	proof := buildWithdrawTestcaseV2Uniswap(tradingSuite.c, 154, 1, tradingSuite.USDCAddress, burnUSDCAmount, pubKeyToAddrStr)
	_, err = SubmitBurnProof(tradingSuite.v, auth, proof)
	require.Equal(tradingSuite.T(), nil, err)
	deposited := tradingSuite.getDepositedBalancePLG(
		tradingSuite.USDCAddress,
		pubKeyToAddrStr.Hex(),
	)
	fmt.Printf("address own asset %v \n", pubKeyToAddrStr)
	fmt.Printf("deposited amount %v \n", deposited.String())

	fmt.Println("------------ step 3: execute trade From USDC for USDT through curve finance --------------")
	tradingSuite.executeWithPcurve(
		deposited,
		big.NewInt(1),
		big.NewInt(2),
		tradingSuite.PolygonPool1,
		true,
	)
	time.Sleep(15 * time.Second)
	usdtTraded := tradingSuite.getDepositedBalancePLG(
		tradingSuite.USDTAddress,
		pubKeyToAddrStr.Hex(),
	)

	fmt.Printf("swap to USDT amount %v \n", usdtTraded.String())
}

func (tradingSuite *PolygonCurveTestSuite) buildPath(paths []common.Address, fees []int64) []byte {
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
