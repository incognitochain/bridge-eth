package main

import (
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/incognitochain/bridge-eth/bridge/kbntrade"
	"github.com/incognitochain/bridge-eth/bridge/vault"
	"github.com/stretchr/testify/suite"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"

	"github.com/stretchr/testify/require"

	kbnmultiTrade "github.com/incognitochain/bridge-eth/bridge/kbnmultitrade"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type KovanKyberTradingTestSuite struct {
	*TradingTestSuite

	KyberTradeDeployedAddr common.Address

	KyberContractAddr    common.Address
	WETHAddr             common.Address
	EtherAddressStrKyber string

	IncKBNTokenIDStr  string
	IncSALTTokenIDStr string
	IncOMGTokenIDStr  string
	IncSNTTokenIDStr  string
	IncPOLYTokenIDStr string
	IncZILTokenIDStr  string

	KBNAddressStr  string
	SALTAddressStr string
	OMGAddressStr  string
	SNTAddressStr  string
	POLYAddressStr string
	ZILAddressStr  string

	IncPrivKeyStrKb     string
	IncPaymentAddrStrKb string

	// token amounts for tests
	DepositingEther             float64
	OMGBalanceAfterStep1        *big.Int
	POLYBalanceAfterStep2       *big.Int
	KyberMultiTradeDeployedAddr common.Address
}

func NewKovanKyberTradingTestSuite(tradingTestSuite *TradingTestSuite) *KovanKyberTradingTestSuite {
	return &KovanKyberTradingTestSuite{
		TradingTestSuite: tradingTestSuite,
	}
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (tradingSuite *KovanKyberTradingTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")
	// Kovan env
	tradingSuite.IncPrivKeyStrKb = "112t8ro4yu78UE82jpto12rp3Cd8Z2Mse7fcavSyXXP82oApE1cg9y8hWq69Z74fFHGJrQyHz54vU8Mv1kx5qZ54cRQJPkF5Cb7DhNqL5Tgt"
	tradingSuite.IncPaymentAddrStrKb = "12RyGbTyktYkXe7mNwmZeD4rktqxtHMe3Tsyf4XiZdKVGFssEHaF1ZUTpXZmpFACuDotVr7a6FEw8v6FTn8DEMqpHNxZ8fJW3KNN3i1"

	tradingSuite.EtherAddressStrKyber = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
	tradingSuite.IncPOLYTokenIDStr = "d0379b8ccc25e4940d5b94ace07dcfa3656a20814279ddf2674f6d7180f65440"
	tradingSuite.IncOMGTokenIDStr = "27322fa7fce2c4d4d5a0022d595a0eec56d7735751a3ba8bc7f10b358ab938bc"
	tradingSuite.IncZILTokenIDStr = "3c115c066028bb682af410c594546b58026095ff149dc30c061749ee163d9051"
	tradingSuite.IncKBNTokenIDStr = "d6644f62d0ef0475335ae7bb6103f358979cbfcd2b85481e3bde2b82101a095c"
	tradingSuite.IncSALTTokenIDStr = "06ce44eae35daf57b9b8158ab95c0cddda9bac208fc380236a318ef40f3ac2ef"
	tradingSuite.IncSNTTokenIDStr = "414a6459526e827321cedb6e574d2ba2eb267c5735b0a65991602a405fb753b7"

	tradingSuite.ZILAddressStr = "0xAb74653cac23301066ABa8eba62b9Abd8a8c51d6"
	tradingSuite.POLYAddressStr = "0xd92266fd053161115163a7311067F0A4745070b5"
	tradingSuite.OMGAddressStr = "0xdB7ec4E4784118D9733710e46F7C83fE7889596a"
	tradingSuite.KBNAddressStr = "0xad67cB4d63C9da94AcA37fDF2761AaDF780ff4a2" // kovan
	tradingSuite.SALTAddressStr = "0x6fEE5727EE4CdCBD91f3A873ef2966dF31713A04"
	tradingSuite.SNTAddressStr = "0x4c99B04682fbF9020Fcb31677F8D8d66832d3322"

	tradingSuite.KyberTradeDeployedAddr = common.HexToAddress("0x434709450d861a16A75488D6b96b4b1c5492Ae17")      //kovan
	tradingSuite.KyberMultiTradeDeployedAddr = common.HexToAddress("0x07AE69CCAB2Ea4eE8323e4fc2a7d64322B27D188") //kovan
	tradingSuite.DepositingEther = float64(0.01)
	tradingSuite.KyberContractAddr = common.HexToAddress("0x692f391bCc85cefCe8C237C01e1f636BbD70EA4D") // kovan
}

func (tradingSuite *KovanKyberTradingTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
	tradingSuite.ETHClient.Close()
}

func (tradingSuite *KovanKyberTradingTestSuite) SetupTest() {
	fmt.Println("Setting up the test...")
}

func (tradingSuite *KovanKyberTradingTestSuite) TearDownTest() {
	fmt.Println("Tearing down the test...")
}
func (tradingSuite *KovanKyberTradingTestSuite) getExpectedRate(
	srcToken string,
	destToken string,
	srcQty *big.Int,
) *big.Int {
	if srcToken == tradingSuite.EtherAddressStr {
		srcToken = tradingSuite.EtherAddressStrKyber
	}
	if destToken == tradingSuite.EtherAddressStr {
		destToken = tradingSuite.EtherAddressStrKyber
	}
	c, err := kbntrade.NewKbntrade(tradingSuite.KyberTradeDeployedAddr, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)
	expectRate, slippageRate, err := c.GetConversionRates(nil, common.HexToAddress(srcToken), srcQty, common.HexToAddress(destToken))
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Printf("slippageRate value: %d\n", slippageRate)
	fmt.Printf("expectRate value: %d\n", expectRate)
	return expectRate
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestKovanKyberTradingTestSuite(t *testing.T) {
	fmt.Println("Starting entry point for Kyber test suite...")

	tradingSuite := new(TradingTestSuite)
	suite.Run(t, tradingSuite)

	kyberTradingSuite := NewKovanKyberTradingTestSuite(tradingSuite)
	suite.Run(t, kyberTradingSuite)

	fmt.Println("Finishing entry point for 0x test suite...")
}

func (tradingSuite *KovanKyberTradingTestSuite) executeWithKyber(
	srcQty *big.Int,
	srcTokenIDStr string,
	destTokenIDStr string,
) {
	tradeAbi, _ := abi.JSON(strings.NewReader(kbntrade.KbntradeABI))

	// Get contract instance
	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)
	auth := bind.NewKeyedTransactor(tradingSuite.ETHPrivKey)
	auth.GasPrice = big.NewInt(50000000000)
	auth.GasLimit = 2000000
	srcToken := common.HexToAddress(srcTokenIDStr)
	destToken := common.HexToAddress(destTokenIDStr)
	//auth.GasPrice = big.NewInt(50000000000)
	auth.GasPrice = big.NewInt(9800000000)
	auth.GasLimit = 2000000
	expectRate := tradingSuite.getExpectedRate(srcTokenIDStr, destTokenIDStr, srcQty)
	input, _ := tradeAbi.Pack("trade", srcToken, srcQty, destToken, expectRate)
	timestamp := []byte(randomizeTimestamp())
	tempData := append(tradingSuite.KyberTradeDeployedAddr[:], input...)
	tempData1 := append(tempData, timestamp...)
	tempData2 := append(tempData1, common.LeftPadBytes(srcQty.Bytes(), 32)...)
	data := rawsha3(tempData2)
	signBytes, _ := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)

	tx, err := c.Execute(
		auth,
		srcToken,
		srcQty,
		destToken,
		tradingSuite.KyberTradeDeployedAddr,
		input,
		timestamp,
		signBytes,
	)
	require.Equal(tradingSuite.T(), nil, err)
	txHash := tx.Hash()
	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("Kyber trade executed , txHash: %x\n", txHash[:])
}

func (tradingSuite *KovanKyberTradingTestSuite) executeMultiTradeWithKyber(
	srcQties []*big.Int,
	srcTokenIDStrs []string,
	destTokenIDStrs []string,
) {
	tradeAbi, _ := abi.JSON(strings.NewReader(kbnmultiTrade.KbnmultiTradeABI))
	auth := bind.NewKeyedTransactor(tradingSuite.ETHPrivKey)
	auth.GasPrice = big.NewInt(50000000000)
	auth.GasLimit = 2000000
	// Deploy kbnMultitrade
	// kbnMultiTradeAddr, tx, _, err := kbnmultiTrade.DeployKbnmultiTrade(auth, tradingSuite.ETHClient, tradingSuite.KyberContractAddr, tradingSuite.VaultAddr)
	// require.Equal(tradingSuite.T(), nil, err)
	// fmt.Println("deployed kbnMultitrade")
	// fmt.Printf("addr: %s\n", kbnMultiTradeAddr.Hex())
	// tradingSuite.KyberMultiTradeDeployedAddr = kbnMultiTradeAddr
	// Get contract instance
	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)
	auth.GasPrice = big.NewInt(50000000000)
	auth.GasLimit = 5000000
	sourceAddresses := make([]common.Address, 0)
	for _, p := range srcTokenIDStrs {
		sourceAddresses = append(sourceAddresses, common.HexToAddress(p))
	}
	destAddresses := make([]common.Address, 0)
	for _, p := range destTokenIDStrs {
		destAddresses = append(destAddresses, common.HexToAddress(p))
	}
	expectRates := make([]*big.Int, 0)
	for i := range destTokenIDStrs {
		expectRates = append(expectRates, tradingSuite.getExpectedRate(srcTokenIDStrs[i], destTokenIDStrs[i], srcQties[i]))
	}
	amounts := make([]byte, 0)
	for i := range srcQties {
		amounts = append(amounts, common.LeftPadBytes(srcQties[i].Bytes(), 32)...)
	}
	input, _ := tradeAbi.Pack("trade", sourceAddresses, srcQties, destAddresses, expectRates)
	timestamp := []byte(randomizeTimestamp())
	tempData := append(tradingSuite.KyberMultiTradeDeployedAddr[:], input...)
	tempData1 := append(tempData, timestamp...)
	tempData2 := append(tempData1, amounts...)
	data := rawsha3(tempData2)
	signBytes, _ := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)
	tx, err := c.ExecuteMulti(
		auth,
		sourceAddresses,
		srcQties,
		destAddresses,
		tradingSuite.KyberMultiTradeDeployedAddr,
		input,
		timestamp,
		signBytes,
	)
	require.Equal(tradingSuite.T(), nil, err)
	txHash := tx.Hash()
	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("Kyber multi trade executed , txHash: %x\n", txHash[:])
}


func (tradingSuite *KovanKyberTradingTestSuite) Test1TradeEthForKNCWithKyber() {
	return
	fmt.Println("============ TEST 1 TRADE ETHER FOR KNC WITH Kyber AGGREGATOR ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
	burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()

	// get info balance initialization
	balPrvInit, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance initialization : ", balPrvInit)

	balpEthInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance initialization : ", balpEthInit)

	balpKNCInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncKNCTokenIDStr)
	fmt.Println("[INFO] pKNC balance initialization : ", balpKNCInit)

	balEthInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance initialization : ", balEthInit)

	balKNCInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.KNCAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] KNC balance initialization : ", balKNCInit)

	balEthScInit := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance initialization on SC : ", balEthScInit)

	balKNCScInit := tradingSuite.getDepositedBalance(
		tradingSuite.KNCAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] KNC balance initialization on SC : ", balKNCScInit)

	fmt.Println("------------ STEP 1: porting ETH to pETH --------------")
	txHash := tradingSuite.depositETH(
		tradingSuite.DepositingEther,
		tradingSuite.IncPaymentAddrStr,
	)
	time.Sleep(15 * time.Second)
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
	time.Sleep(80 * time.Second)
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
	balKNCScS2 := tradingSuite.getDepositedBalance(
		tradingSuite.KNCAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] KNC balance on SC at step 2 : ", balKNCScS2)

	//require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPETH, big.NewInt(1000000000)), deposited)
	fmt.Println("------------ step 3: execute trade ETH for KNC through Kyber aggregator --------------")
	tradingSuite.executeWithKyber(
		tradeAmount,
		tradingSuite.EtherAddressStr,
		tradingSuite.KNCAddressStr,
	)
	balEthScTradeS3 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC after trade at step 3 : ", balEthScTradeS3)
	// TODO assert ETH balane on SC
	balKNCScTradeS3 := tradingSuite.getDepositedBalance(
		tradingSuite.KNCAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] KNC balance on SC after trade at step 3 : ", balKNCScTradeS3)
	// TODO assert KNC balane on SC
	require.NotEqual(tradingSuite.T(), balKNCScTradeS3, balKNCScS2, "trade failed")

	fmt.Println("------------ step 4: withdrawing KNC from SC to pKNC on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.KNCAddressStr,
		balKNCScTradeS3,
	)
	time.Sleep(15 * time.Second)

	balKNCScDepS4 := tradingSuite.getDepositedBalance(
		tradingSuite.KNCAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] KNC balance on SC after withdraw at step 4 : ", balKNCScDepS4)
	// TODO assert KNC balane on SC
	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(80 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncKNCTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)

	balpEthAfIssS4, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncKNCTokenIDStr,
	)
	fmt.Println("[INFO] pKNC balance after issuing step 4 : ", balpEthAfIssS4)

	balPrvAfIssS4, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after issuing step 4: ", balPrvAfIssS4)
	// TODO assert PRV balance remain

	fmt.Println("------------ step 5: withdrawing pKNC from Incognito to KNC --------------")
	withdrawingPKNC := big.NewInt(0).Div(balKNCScTradeS3, big.NewInt(1000000000))
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncKNCTokenIDStr,
		withdrawingPKNC,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))

	balKNC := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.KNCAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	//tradingSuite.KNCBalanceAfterStep1 = balKNC
	fmt.Println("KNC balance after trade: ", balKNC)
	// require.Equal(tradingSuite.T(), withdrawingPKNC.Uint64(), bal.Div(bal, big.NewInt(1000000000)).Uint64())
	balEth := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("ETH balance after trade: ", balEth)
}

func (tradingSuite *KovanKyberTradingTestSuite) Test2TradeEthForOMGWithKyber() {
return
	fmt.Println("============ TEST 2 TRADE ETHER FOR OMG WITH Kyber AGGREGATOR ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
	burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()

	// get info balance initialization
	balPrvInit, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance initialization : ", balPrvInit)

	balpEthInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance initialization : ", balpEthInit)

	balpOMGInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncOMGTokenIDStr)
	fmt.Println("[INFO] pOMG balance initialization : ", balpOMGInit)

	balEthInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance initialization : ", balEthInit)

	balOMGInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.OMGAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] OMG balance initialization : ", balOMGInit)

	balEthScInit := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance initialization on SC : ", balEthScInit)

	balOMGScInit := tradingSuite.getDepositedBalance(
		tradingSuite.OMGAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] OMG balance initialization on SC : ", balOMGScInit)

	fmt.Println("------------ STEP 1: porting ETH to pETH --------------")
	txHash := tradingSuite.depositETH(
		tradingSuite.DepositingEther,
		tradingSuite.IncPaymentAddrStr,
	)
	time.Sleep(15 * time.Second)
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
	time.Sleep(80 * time.Second)
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
	balOMGScS2 := tradingSuite.getDepositedBalance(
		tradingSuite.OMGAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] OMG balance on SC at step 2 : ", balOMGScS2)

	//require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPETH, big.NewInt(1000000000)), deposited)
	fmt.Println("------------ step 3: execute trade ETH for OMG through Kyber aggregator --------------")
	tradingSuite.executeWithKyber(
		tradeAmount,
		tradingSuite.EtherAddressStr,
		tradingSuite.OMGAddressStr,
	)
	balEthScTradeS3 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC after trade at step 3 : ", balEthScTradeS3)
	// TODO assert ETH balane on SC
	balOMGScTradeS3 := tradingSuite.getDepositedBalance(
		tradingSuite.OMGAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] OMG balance on SC after trade at step 3 : ", balOMGScTradeS3)
	// TODO assert OMG balane on SC

	fmt.Println("------------ step 4: withdrawing OMG from SC to pOMG on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.OMGAddressStr,
		balOMGScTradeS3,
	)
	time.Sleep(15 * time.Second)

	balOMGScDepS4 := tradingSuite.getDepositedBalance(
		tradingSuite.OMGAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] OMG balance on SC after withdraw at step 4 : ", balOMGScDepS4)
	// TODO assert OMG balane on SC
	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(80 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncOMGTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)

	balpEthAfIssS4, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncOMGTokenIDStr,
	)
	fmt.Println("[INFO] pOMG balance after issuing step 4 : ", balpEthAfIssS4)

	balPrvAfIssS4, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after issuing step 4: ", balPrvAfIssS4)
	// TODO assert PRV balance remain

	fmt.Println("------------ step 5: withdrawing pOMG from Incognito to OMG --------------")
	withdrawingPOMG := big.NewInt(0).Div(balOMGScTradeS3, big.NewInt(1000000000))
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncOMGTokenIDStr,
		withdrawingPOMG,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))

	balOMG := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.OMGAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	tradingSuite.OMGBalanceAfterStep1 = balOMG
	fmt.Println("OMG balance after trade: ", balOMG)
	// require.Equal(tradingSuite.T(), withdrawingPOMG.Uint64(), bal.Div(bal, big.NewInt(1000000000)).Uint64())
	balEth := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("ETH balance after trade: ", balEth)
}

func (tradingSuite *KovanKyberTradingTestSuite) Test3TradeOMGForZILWithKyber() {
return
	fmt.Println("============ TEST 3 TRADE OMG FOR ZIL WITH KYBER AGGREGATOR ===========")
	fmt.Println("------------ step 0: declaration & initialization --------------")
	depositingOMG := tradingSuite.OMGBalanceAfterStep1
	burningPOMG := big.NewInt(0).Div(depositingOMG, big.NewInt(1000000000))
	tradeAmount := depositingOMG

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("pubKeyToAddrStr: ", pubKeyToAddrStr)

	// get info balance initialization
	balPrvInit, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance initialization : ", balPrvInit)

	balpZILInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncZILTokenIDStr)
	fmt.Println("[INFO] pZIL balance initialization : ", balpZILInit)

	balpOMGInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncOMGTokenIDStr)
	fmt.Println("[INFO] pOMG balance initialization : ", balpOMGInit)

	balZILInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.ZILAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ZIL balance initialization : ", balZILInit)

	balOMGInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.OMGAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] OMG balance initialization : ", balOMGInit)

	balZILScInit := tradingSuite.getDepositedBalance(
		tradingSuite.ZILAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ZIL balance initialization on SC : ", balZILScInit)

	balOMGScInit := tradingSuite.getDepositedBalance(
		tradingSuite.OMGAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] OMG balance initialization on SC : ", balOMGScInit)

	fmt.Println("------------ step 1: porting OMG to pOMG --------------")
	txHash := tradingSuite.depositERC20ToBridge(
		depositingOMG,
		common.HexToAddress(tradingSuite.OMGAddressStr),
		tradingSuite.IncPaymentAddrStr,
	)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	// get OMG remain after depsit
	balOMGAfDep := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.OMGAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] OMG balance remain after deposit : ", balOMGAfDep)
	// TODO : assert OMG balance

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(80 * time.Second)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncOMGTokenIDStr,
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
	balpOMGAfIssS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncOMGTokenIDStr)
	fmt.Println("[INFO] pOMG balance after issuing step 1 : ", balpOMGAfIssS1)
	// TODO assert pOMG balance issuing

	fmt.Println("------------ step 2: burning pOMG to deposit OMG to SC --------------")

	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningPToken(
		tradingSuite.IncOMGTokenIDStr,
		burningPOMG,
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
	balpOMGAfBurnS2, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncOMGTokenIDStr)
	fmt.Println("[INFO] pOMG balance after burning step 2 : ", balpOMGAfBurnS2)
	// TODO assert pOMG balance issuing

	tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))

	balOMGScDepS2 := tradingSuite.getDepositedBalance(
		tradingSuite.OMGAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] OMG balance after deposit on SC before trade at step 2: ", balOMGScDepS2)
	// TODO assert OMG balane on SC
	// require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPOMG, big.NewInt(1000000000)), deposited)
	balZILScS2 := tradingSuite.getDepositedBalance(
		tradingSuite.IncZILTokenIDStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ZIL balance on SC before trade at step 2 : ", balZILScS2)

	fmt.Println("------------ step 3: execute trade OMG for ZIL through Kyber aggregator --------------")
	tradingSuite.executeWithKyber(
		tradeAmount,
		tradingSuite.OMGAddressStr,
		tradingSuite.ZILAddressStr,
	)

	balZILScTradeS3 := tradingSuite.getDepositedBalance(
		tradingSuite.ZILAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ZIL balance on SC after trade at step 3 : ", balZILScTradeS3)
	// TODO assert ZIL balane on SC
	balOMGScTradeS3 := tradingSuite.getDepositedBalance(
		tradingSuite.OMGAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] OMG balance on SC after trade at step 3 : ", balOMGScTradeS3)
	// TODO assert OMG balane on SC

	fmt.Println("------------ step 4: withdrawing ZIL from SC to pZIL on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.ZILAddressStr,
		balZILScTradeS3,
	)

	balZILScDepS4 := tradingSuite.getDepositedBalance(
		tradingSuite.ZILAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] pZIL balance on SC after withdraw at step 4 : ", balZILScDepS4)
	// TODO assert ZIL balane on SC

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(80 * time.Second)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncZILTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)

	balpZILAfIssS4, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncZILTokenIDStr,
	)
	fmt.Println("[INFO] pZILbalance after issuing step 4 : ", balpZILAfIssS4)
	// TODO assert pZIL balance issuing
	balPrvAfIssS4, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after issuing step 4: ", balPrvAfIssS4)
	// TODO assert PRV balance remain

	fmt.Println("------------ step 5: withdrawing pZIL from Incognito to ZIL --------------")
	withdrawingPZIL := big.NewInt(0).Div(balZILScTradeS3, big.NewInt(1000))
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncZILTokenIDStr,
		withdrawingPZIL,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))

	balZIL := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.ZILAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	tradingSuite.POLYBalanceAfterStep2 = balZIL
	fmt.Println("ZIL balance after step 2: ", balZIL)
	// require.Equal(tradingSuite.T(), withdrawingPZIL.Uint64(), bal.Uint64())
	// require.Equal(tradingSuite.T(), withdrawingPZIL.Uint64(), bal.Div(bal, big.NewInt(1000000000)).Uint64())
	balOMG := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.OMGAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("OMG balance after trade: ", balOMG)
}

func (tradingSuite *KovanKyberTradingTestSuite) Test4TradeZILForEthWithKyber() {
	return
	fmt.Println("============ TEST TRADE ZIL FOR ETH WITH KYBER AGGREGATOR ===========")
	fmt.Println("------------ step 0: declaration & initialization --------------")
	depositingZIL := tradingSuite.POLYBalanceAfterStep2
	burningPZIL := big.NewInt(0).Div(depositingZIL, big.NewInt(1000))
	tradeAmount := depositingZIL

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("pubKeyToAddrStr: ", pubKeyToAddrStr)

	// get info balance initialization
	balPrvInit, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance initialization : ", balPrvInit)

	balpZILInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncZILTokenIDStr)
	fmt.Println("[INFO] pZIL balance initialization : ", balpZILInit)

	balpEthInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance initialization : ", balpEthInit)

	balZILInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.ZILAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ZIL balance initialization : ", balZILInit)

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

	balZILScInit := tradingSuite.getDepositedBalance(
		tradingSuite.ZILAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ZIL balance initialization on SC : ", balZILScInit)

	fmt.Println("------------ step 1: porting ZIL to pZIL --------------")
	txHash := tradingSuite.depositERC20ToBridge(
		depositingZIL,
		common.HexToAddress(tradingSuite.ZILAddressStr),
		tradingSuite.IncPaymentAddrStr,
	)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)
	// get ZIL remain after depsit
	balZILAfDep := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.ZILAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ZIL balance remain after deposit : ", balZILAfDep)
	// TODO : assert ZIL balance
	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(80 * time.Second)

	issuingRes, err := tradingSuite.callIssuingETHReq(
		tradingSuite.IncZILTokenIDStr,
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
	balpZILAfIssS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncZILTokenIDStr)
	fmt.Println("[INFO] pZIL balance after issuing step 1 : ", balpZILAfIssS1)
	// TODO assert pZIL balance issuing

	fmt.Println("------------ step 2: burning pZIL to deposit ZIL to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningPToken(
		tradingSuite.IncZILTokenIDStr,
		burningPZIL,
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
	balpZILAfBurnS2, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncZILTokenIDStr)
	fmt.Println("[INFO] pZILbalance after burning step 2 : ", balpZILAfBurnS2)
	// TODO assert pZIL balance issuing

	tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))

	balZILScDepS2 := tradingSuite.getDepositedBalance(
		tradingSuite.ZILAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ZIL balance after deposit on SC at step 2: ", balZILScDepS2)
	//require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPZIL, big.NewInt(1000000000)), balZILScDepS2)

	balEthScS2 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC at step 2 : ", balEthScS2)

	fmt.Println("------------ step 3: execute trade ZIL for ETH through Kyber aggregator --------------")
	tradingSuite.executeWithKyber(
		tradeAmount,
		tradingSuite.ZILAddressStr,
		tradingSuite.EtherAddressStr,
	)
	balEthScTradeS3 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC after trade at step 3 : ", balEthScTradeS3)
	// TODO assert ETH balane on SC
	balZILScTradeS3 := tradingSuite.getDepositedBalance(
		tradingSuite.ZILAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ZIL balance on SC after trade at step 3 : ", balZILScTradeS3)
	// TODO assert ZIL balane on SC
	fmt.Println("------------ step 4: withdrawing ETH from SC to pETH on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.EtherAddressStr,
		balEthScTradeS3,
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
	time.Sleep(80 * time.Second)

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
	time.Sleep(140 * time.Second)

	tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))

	balETH := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("Ether balance after trade: ", balETH)
	// require.Equal(tradingSuite.T(), withdrawingPETH.Uint64(), bal.Div(bal, big.NewInt(1000000000)).Uint64())
	balZIL := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.ZILAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("ZIL balance after trade: ", balZIL)
}

func (tradingSuite *KovanKyberTradingTestSuite) Test5TradeEthithKyber() {
	return
	fmt.Println("============ TEST 5 TRADE ETHER Kyber AGGREGATOR ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
	burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()

	// get info balance initialization
	balPrvInit, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance initialization : ", balPrvInit)

	balpEthInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance initialization : ", balpEthInit)

	balpOMGInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncOMGTokenIDStr)
	fmt.Println("[INFO] pOMG balance initialization : ", balpOMGInit)

	balEthInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance initialization : ", balEthInit)

	balOMGInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.OMGAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] OMG balance initialization : ", balOMGInit)

	balEthScInit := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance initialization on SC : ", balEthScInit)

	balOMGScInit := tradingSuite.getDepositedBalance(
		tradingSuite.OMGAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] OMG balance initialization on SC : ", balOMGScInit)

	fmt.Println("------------ STEP 1: porting ETH to pETH --------------")
	txHash := tradingSuite.depositETH(
		tradingSuite.DepositingEther,
		tradingSuite.IncPaymentAddrStr,
	)
	time.Sleep(15 * time.Second)
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
	time.Sleep(80 * time.Second)
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

	fmt.Println("------------ STEP 2: burning pETH to deposit ZIL to SC --------------")
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
	balOMGScS2 := tradingSuite.getDepositedBalance(
		tradingSuite.OMGAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] OMG balance on SC at step 2 : ", balOMGScS2)

	//require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPETH, big.NewInt(1000000000)), deposited)
	fmt.Println("------------ step 3: execute trade through Kyber aggregator --------------")
	fmt.Println("------------ step 3.1: execute trade ETH for KBN through Kyber aggregator --------------")
	tradingSuite.executeWithKyber(
		tradeAmount,
		tradingSuite.EtherAddressStr,
		tradingSuite.KBNAddressStr,
	)
	balEthScTradeS31 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC after trade at step 3.1 : ", balEthScTradeS31)
	// TODO assert ETH balane on SC
	balKBNScTradeS31 := tradingSuite.getDepositedBalance(
		tradingSuite.KBNAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] KBN balance on SC after trade at step 3.1 : ", balKBNScTradeS31)
	// TODO assert OMG balane on SC

	fmt.Println("------------ step 3.2: execute trade KBN for SALT through Kyber aggregator --------------")
	tradeAmount = balKBNScTradeS31
	tradingSuite.executeWithKyber(
		tradeAmount,
		tradingSuite.KBNAddressStr,
		tradingSuite.SALTAddressStr,
	)
	balKBNScTradeS32 := tradingSuite.getDepositedBalance(
		tradingSuite.KBNAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] OMG balance on SC after trade at step 3.2 : ", balKBNScTradeS32)
	// TODO assert OMG balane on SC

	balSALTScTradeS32 := tradingSuite.getDepositedBalance(
		tradingSuite.SALTAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] SALT balance on SC after trade at step 3.2 : ", balSALTScTradeS32)
	// TODO assert ETH balane on SC

	fmt.Println("------------ step 3.2: execute trade SALT for ETH through Kyber aggregator --------------")
	tradeAmount = balSALTScTradeS32
	tradingSuite.executeWithKyber(
		tradeAmount,
		tradingSuite.SALTAddressStr,
		tradingSuite.EtherAddressStr,
	)

	balSALTScTradeS33 := tradingSuite.getDepositedBalance(
		tradingSuite.SALTAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] SALT balance on SC after trade at step 3.3 : ", balSALTScTradeS33)
	// TODO assert ETH balane on SC

	balEthScTradeS33 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC after trade at step 3.3 : ", balEthScTradeS33)
	// TODO assert ETH balane on SC

	fmt.Println("------------ step 4: withdrawing ETH from SC to pETH on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.EtherAddressStr,
		balEthScTradeS33,
	)
	time.Sleep(15 * time.Second)
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
	time.Sleep(80 * time.Second)

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
	time.Sleep(140 * time.Second)

	tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))

	balETH := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("Ether balance after trade: ", balETH)
	// require.Equal(tradingSuite.T(), withdrawingPETH.Uint64(), bal.Div(bal, big.NewInt(1000000000)).Uint64())

}

func (tradingSuite *KovanKyberTradingTestSuite) Test6MultiForKNCWithKyber() {
	return
	fmt.Println("============ TEST 6 MULTI WITH Kyber AGGREGATOR ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
	AmountDeposit := big.NewInt(int64(tradingSuite.DepositingEther * 2 * params.Ether))
	burningPETH := big.NewInt(0).Div(AmountDeposit, big.NewInt(1000000000))

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()

	// get info balance initialization
	balPrvInit, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance initialization : ", balPrvInit)

	balpEthInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance initialization : ", balpEthInit)

	balpKBNInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncKBNTokenIDStr)
	fmt.Println("[INFO] pKBN balance initialization : ", balpKBNInit)

	balEthInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance initialization : ", balEthInit)

	balKBNInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.KBNAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] KBN balance initialization : ", balKBNInit)

	balEthScInit := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance initialization on SC : ", balEthScInit)

	balKBNScInit := tradingSuite.getDepositedBalance(
		tradingSuite.KBNAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] KBN balance initialization on SC : ", balKBNScInit)

	fmt.Println("------------ STEP 1: porting ETH to pETH --------------")
	fmt.Println("ETH deposit amount : ", tradingSuite.DepositingEther*2)
	txHash := tradingSuite.depositETH(
		tradingSuite.DepositingEther*2,
		tradingSuite.IncPaymentAddrStr,
	)
	time.Sleep(15 * time.Second)
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
	time.Sleep(80 * time.Second)
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
	balKNCScS2 := tradingSuite.getDepositedBalance(
		tradingSuite.KNCAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] KNC balance on SC at step 2 : ", balKNCScS2)

	//require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPETH, big.NewInt(1000000000)), deposited)
	fmt.Println("------------ step 3: execute trade ETH for KBN through Kyber aggregator --------------")
	tradingSuite.executeWithKyber(
		tradeAmount,
		tradingSuite.EtherAddressStr,
		tradingSuite.KBNAddressStr,
	)
	balEthScTradeS3 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC after trade at step 3 : ", balEthScTradeS3)
	// TODO assert ETH balane on SC
	balKBNScTradeS3 := tradingSuite.getDepositedBalance(
		tradingSuite.KBNAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] KBN balance on SC after trade at step 3 : ", balKBNScTradeS3)
	// TODO assert KBN balane on SC
	//require.NotEqual(tradingSuite.T(),balKNCScTradeS3,balKBNScS2,"trade failed" )

	fmt.Println("------------ step 4: execute trade ETH and KBN for OMG and SALT and through Kyber aggregator --------------")

	balOMGScTradeS4 := tradingSuite.getDepositedBalance(
		tradingSuite.OMGAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] OMG balance on SC before trade at step 4 : ", balOMGScTradeS4)
	// TODO assert ETH balane on SC
	balSALScTradeS4 := tradingSuite.getDepositedBalance(
		tradingSuite.SALTAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] SALT balance on SC before trade at step 4 : ", balSALScTradeS4)

	tradingSuite.executeMultiTradeWithKyber(
		[]*big.Int{balKBNScTradeS3, balEthScTradeS3},
		[]string{tradingSuite.KBNAddressStr, tradingSuite.EtherAddressStr},
		[]string{tradingSuite.OMGAddressStr, tradingSuite.SALTAddressStr},
	)
	time.Sleep(15 * time.Second)
	balEthScTradeS4 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC after trade at step 4 : ", balEthScTradeS4)
	// TODO assert ETH balane on SC
	balKBNScTradeS4 := tradingSuite.getDepositedBalance(
		tradingSuite.KBNAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] KBN balance on SC after trade at step 4 : ", balKBNScTradeS4)
	// TODO assert KBN balane on SC
	balOMGScTradeS41 := tradingSuite.getDepositedBalance(
		tradingSuite.OMGAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] OMG balance on SC after trade at step 4 : ", balOMGScTradeS41)
	// TODO assert ETH balane on SC
	balSALScTradeS41 := tradingSuite.getDepositedBalance(
		tradingSuite.SALTAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] SALT balance on SC after trade at step 4 : ", balSALScTradeS41)

	fmt.Println("------------ step 5: execute trade ETH and KBN for OMG and SALT and through Kyber aggregator --------------")

	tradingSuite.executeMultiTradeWithKyber(
		[]*big.Int{balOMGScTradeS41, balSALScTradeS41},
		[]string{tradingSuite.OMGAddressStr, tradingSuite.SALTAddressStr},
		[]string{tradingSuite.KBNAddressStr, tradingSuite.EtherAddressStr},
	)
	time.Sleep(15 * time.Second)
	balEthScTradeS5 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC after trade at step 5 : ", balEthScTradeS5)
	// TODO assert ETH balane on SC
	balKBNScTradeS5 := tradingSuite.getDepositedBalance(
		tradingSuite.KBNAddressStr,
		pubKeyToAddrStr,
	)

	fmt.Println("[INFO] KBN balance on SC after trade at step 5 : ", balKBNScTradeS5)
	// TODO assert KBN balane on SC
	balOMGScTradeS5 := tradingSuite.getDepositedBalance(
		tradingSuite.OMGAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] OMG balance on SC after trade at step 5 : ", balOMGScTradeS5)
	// TODO assert ETH balane on SC
	balSALScTradeS5 := tradingSuite.getDepositedBalance(
		tradingSuite.SALTAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] SALT balance on SC after trade at step 5 : ", balSALScTradeS5)

	fmt.Println("------------ step 6: withdrawing ETH from SC to pETH on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.EtherAddressStr,
		balEthScTradeS5,
	)
	time.Sleep(15 * time.Second)
	balEthScDepS4 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC after withdraw at step 6 : ", balEthScDepS4)
	// TODO assert ETH balane on SC

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(80 * time.Second)

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
	fmt.Println("[INFO] pETH balance after issuing step 6 : ", balpEthAfIssS4)
	// TODO assert pETH balance issuing
	balPrvAfIssS4, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after issuing step 6: ", balPrvAfIssS4)
	// TODO assert PRV balance remain

	fmt.Println("------------ step 7: withdrawing pETH from Incognito to ETH --------------")
	withdrawingPETH := big.NewInt(0).Div(balEthScTradeS5, big.NewInt(1000000000))
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

	tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))

	balETH := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("Ether balance after trade: ", balETH)

}

func (tradingSuite *KovanKyberTradingTestSuite) Test7tradeOldSmartContract() {
	return
	fmt.Println("============ TEST 7 TRADE ON OLD SMART CONTRACT ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	depositETH := tradingSuite.DepositingEther * 6
	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
	burningPETH := big.NewInt(0).Div(big.NewInt(int64(tradingSuite.DepositingEther*6*params.Ether)), big.NewInt(1000000000))

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()

	// get info balance initialization
	balPrvInit, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance initialization : ", balPrvInit)
	fmt.Println("------------ ETH --------------")
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

	fmt.Println("------------ OMG --------------")

	balOMGInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.OMGAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] OMG balance initialization : ", balOMGInit)

	balOMGScInit := tradingSuite.getDepositedBalance(
		tradingSuite.OMGAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] OMG balance initialization on SC : ", balOMGScInit)

	balpOMGInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncOMGTokenIDStr)
	fmt.Println("[INFO] pOMG balance initialization : ", balpOMGInit)

	fmt.Println("------------ ZIL --------------")

	balZILInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.ZILAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ZIL balance initialization : ", balZILInit)

	balZILScInit := tradingSuite.getDepositedBalance(
		tradingSuite.ZILAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ZIL balance initialization on SC : ", balZILScInit)

	balpZILInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncZILTokenIDStr)
	fmt.Println("[INFO] pZIL balance initialization : ", balpZILInit)

	fmt.Println("------------ KBN --------------")

	balKBNInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.KBNAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] KBN balance initialization : ", balKBNInit)

	balKBNScInit := tradingSuite.getDepositedBalance(
		tradingSuite.KBNAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] KBN balance initialization on SC : ", balKBNScInit)

	balpKBNInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncKBNTokenIDStr)
	fmt.Println("[INFO] pKBN balance initialization : ", balpKBNInit)

	fmt.Println("------------ SALT --------------")

	balSALTInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.SALTAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] SALT balance initialization : ", balSALTInit)

	balSALTScInit := tradingSuite.getDepositedBalance(
		tradingSuite.SALTAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] SALT balance initialization on SC : ", balSALTScInit)

	balpSALTInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncSALTTokenIDStr)
	fmt.Println("[INFO] pSALT balance initialization : ", balpSALTInit)

	fmt.Println("------------ STEP 1: porting ETH to pETH --------------")
	fmt.Println("[INFO] ETH deposit input : ", burningPETH)
	txHash := tradingSuite.depositETH(
		depositETH,
		tradingSuite.IncPaymentAddrStr,
	)
	//time.Sleep(15 * time.Second)
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
	time.Sleep(100 * time.Second)

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
	time.Sleep(100 * time.Second)

	balpEthAfBurnS2, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance after burning step 2 : ", balpEthAfBurnS2)
	// TODO assert pETH balance issuing

	tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))
	balEthScDepS2 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance after deposit on SC at step 2: ", balEthScDepS2)

	balOMGScS2 := tradingSuite.getDepositedBalance(
		tradingSuite.OMGAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] OMG balance on SC at step 2 : ", balOMGScS2)

	balZILScS2 := tradingSuite.getDepositedBalance(
		tradingSuite.ZILAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ZIL balance on SC at step 2 : ", balZILScS2)

	balKBNScS2 := tradingSuite.getDepositedBalance(
		tradingSuite.KBNAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] KBN balance on SC at step 2 : ", balKBNScS2)

	balSALTScS2 := tradingSuite.getDepositedBalance(
		tradingSuite.SALTAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] SALT balance on SC at step 2 : ", balSALTScS2)

	//require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPETH, big.NewInt(1000000000)), deposited)
	fmt.Println("------------ step 3: execute trade ETH through Kyber aggregator --------------")

	// fmt.Println("------------  execute trade ETH to OMG 3.1--------------")
	// tradingSuite.executeWithKyber(
	//   tradeAmount,
	//   tradingSuite.EtherAddressStr,
	//   tradingSuite.OMGAddressStr,
	// )
	// fmt.Println("------------  execute trade ETH to ZIL 3.1--------------")
	// tradingSuite.executeWithKyber(
	//   tradeAmount,
	//   tradingSuite.EtherAddressStr,
	//   tradingSuite.ZILAddressStr,
	// )
	// fmt.Println("------------  execute trade ETH to KBN 3.1 --------------")
	// tradingSuite.executeWithKyber(
	//   tradeAmount,
	//   tradingSuite.EtherAddressStr,
	//   tradingSuite.KBNAddressStr,
	// )
	// fmt.Println("------------  execute trade ETH to SALT 3.1--------------")
	// tradingSuite.executeWithKyber(
	//   tradeAmount,
	//   tradingSuite.EtherAddressStr,
	//   tradingSuite.SALTAddressStr,
	// )
	tradingSuite.executeMultiTradeWithKyber(
		[]*big.Int{tradeAmount, tradeAmount, tradeAmount, tradeAmount},
		[]string{tradingSuite.EtherAddressStr, tradingSuite.EtherAddressStr, tradingSuite.EtherAddressStr, tradingSuite.EtherAddressStr},
		[]string{tradingSuite.OMGAddressStr, tradingSuite.SALTAddressStr, tradingSuite.ZILAddressStr, tradingSuite.KBNAddressStr},
	)

	balEthScTradeS3 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC after trade at step 3.1 : ", balEthScTradeS3)

	balOMGScTradeS3 := tradingSuite.getDepositedBalance(
		tradingSuite.OMGAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] OMG balance on SC after trade at step 3.1 : ", balOMGScTradeS3)

	balZILScS3 := tradingSuite.getDepositedBalance(
		tradingSuite.ZILAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ZIL balance on SC after trade at step 3.1 : ", balZILScS3)

	balKBNScS3 := tradingSuite.getDepositedBalance(
		tradingSuite.KBNAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] KBN balance on SC after trade at step 3.1 : ", balKBNScS3)

	balSALTScS3 := tradingSuite.getDepositedBalance(
		tradingSuite.SALTAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] SALT balance on SC after trade at step 3.1 : ", balSALTScS3)

	fmt.Println("------------  execute trade OMG to SALT 3.2--------------")
	tradingSuite.executeWithKyber(
		balOMGScTradeS3,
		tradingSuite.OMGAddressStr,
		tradingSuite.SALTAddressStr,
	)

	balOMGScTradeS32 := tradingSuite.getDepositedBalance(
		tradingSuite.OMGAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] OMG balance on SC after trade at step 3.2 : ", balOMGScTradeS32)

	balSALTScS32 := tradingSuite.getDepositedBalance(
		tradingSuite.SALTAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] SALT balance on SC after trade at step 3.2 : ", balSALTScS32)

	balEthScTradeS32 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC after trade at step 3.2 : ", balEthScTradeS32)

	fmt.Println("------------  execute SALT OMG to Ether 3.3--------------")
	tradingSuite.executeWithKyber(
		balSALTScS32,
		tradingSuite.SALTAddressStr,
		tradingSuite.EtherAddressStr,
	)

	balSALTScS33 := tradingSuite.getDepositedBalance(
		tradingSuite.SALTAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] SALT balance on SC after trade at step 3.3 : ", balSALTScS33)

	balEthScTradeS33 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC after trade at step 3.3 : ", balEthScTradeS33)

	fmt.Println("------------  execute multi trade Ether to OMG, SALT 3.4--------------")
	tradingSuite.executeMultiTradeWithKyber(
		[]*big.Int{tradeAmount, tradeAmount},
		[]string{tradingSuite.EtherAddressStr, tradingSuite.EtherAddressStr},
		[]string{tradingSuite.OMGAddressStr, tradingSuite.SALTAddressStr},
	)

	balEthScTradeS34 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC after trade at step 3.4 : ", balEthScTradeS34)

	balOMGScTradeS34 := tradingSuite.getDepositedBalance(
		tradingSuite.OMGAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] OMG balance on SC after trade at step 3.4 : ", balOMGScTradeS34)

	balZILScS34 := tradingSuite.getDepositedBalance(
		tradingSuite.ZILAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ZIL balance on SC after trade at step 3.4 : ", balZILScS34)

	balKBNScS34 := tradingSuite.getDepositedBalance(
		tradingSuite.KBNAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] KBN balance on SC after trade at step 3.4 : ", balKBNScS34)

	balSALTScS34 := tradingSuite.getDepositedBalance(
		tradingSuite.SALTAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] SALT balance on SC after trade at step 3.4 : ", balSALTScS34)

	fmt.Println("------------ step 4: withdrawing ETH from SC to pETH on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.EtherAddressStr,
		balEthScTradeS34,
	)

	txHashByEmittingWithdrawalReqOMG := tradingSuite.requestWithdraw(
		tradingSuite.OMGAddressStr,
		balOMGScTradeS34,
	)

	txHashByEmittingWithdrawalReqZIL := tradingSuite.requestWithdraw(
		tradingSuite.ZILAddressStr,
		balZILScS34,
	)

	txHashByEmittingWithdrawalReqKBN := tradingSuite.requestWithdraw(
		tradingSuite.KBNAddressStr,
		balKBNScS34,
	)

	txHashByEmittingWithdrawalReqSALT := tradingSuite.requestWithdraw(
		tradingSuite.SALTAddressStr,
		balSALTScS34,
	)
	//time.Sleep(15 * time.Second)
	balEthScDepS4 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC after withdraw at step 4 : ", balEthScDepS4)

	balOMGScDepS4 := tradingSuite.getDepositedBalance(
		tradingSuite.OMGAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] OMG balance on SC after withdraw at step 4 : ", balOMGScDepS4)

	balZILScDepS4 := tradingSuite.getDepositedBalance(
		tradingSuite.ZILAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ZIL balance on SC after withdraw at step 4 : ", balZILScDepS4)

	balKBNScDepS4 := tradingSuite.getDepositedBalance(
		tradingSuite.KBNAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] KBN balance on SC after withdraw at step 4 : ", balKBNScDepS4)

	balSALTScDepS4 := tradingSuite.getDepositedBalance(
		tradingSuite.SALTAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] SALT balance on SC after withdraw at step 4 : ", balSALTScDepS4)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req ETH : ", ethBlockHash, ethTxIdx, ethDepositProof)

	_, ethBlockHashOMG, ethTxIdxOMG, ethDepositProofOMG, err := getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReqOMG)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req OMG : ", ethBlockHashOMG, ethTxIdxOMG, ethDepositProofOMG)

	_, ethBlockHashZIL, ethTxIdxZIL, ethDepositProofZIL, err := getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReqZIL)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req ZIL : ", ethBlockHashZIL, ethTxIdxZIL, ethDepositProofZIL)

	_, ethBlockHashKBN, ethTxIdxKBN, ethDepositProofKBN, err := getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReqKBN)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req KBN : ", ethBlockHashKBN, ethTxIdxKBN, ethDepositProofKBN)

	_, ethBlockHashSALT, ethTxIdxSALT, ethDepositProofSALT, err := getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReqSALT)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req SALT: ", ethBlockHashSALT, ethTxIdxSALT, ethDepositProofSALT)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(90 * time.Second)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(60 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncOMGTokenIDStr,
		ethDepositProofOMG,
		ethBlockHashOMG,
		ethTxIdxOMG,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(60 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncZILTokenIDStr,
		ethDepositProofZIL,
		ethBlockHashZIL,
		ethTxIdxZIL,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(60 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncKBNTokenIDStr,
		ethDepositProofKBN,
		ethBlockHashKBN,
		ethTxIdxKBN,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(60 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncSALTTokenIDStr,
		ethDepositProofSALT,
		ethBlockHashSALT,
		ethTxIdxSALT,
	)
	require.Equal(tradingSuite.T(), nil, err)

	time.Sleep(100 * time.Second)

	balpEthAfIssS4, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncEtherTokenIDStr,
	)
	fmt.Println("[INFO] pETH balance after issuing step 4 : ", balpEthAfIssS4)

	balpOMGAfIssS4, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncOMGTokenIDStr,
	)
	fmt.Println("[INFO] pOMG balance after issuing step 4 : ", balpOMGAfIssS4)

	balpZILAfIssS4, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncZILTokenIDStr,
	)
	fmt.Println("[INFO] pZIL balance after issuing step 4 : ", balpZILAfIssS4)

	balpKBNAfIssS4, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncKBNTokenIDStr,
	)
	fmt.Println("[INFO] pKBN balance after issuing step 4 : ", balpKBNAfIssS4)

	balpSALTAfIssS4, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncSALTTokenIDStr,
	)
	fmt.Println("[INFO] pSALT balance after issuing step 4 : ", balpSALTAfIssS4)

	balPrvAfIssS4, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after issuing step 4: ", balPrvAfIssS4)
	// TODO assert PRV balance remain

	fmt.Println("------------ step 5: withdrawing --------------")
	withdrawingPETH := big.NewInt(0).Div(balEthScTradeS34, big.NewInt(1000000000))
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncEtherTokenIDStr,
		withdrawingPETH,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(60 * time.Second)

	withdrawingPOMG := big.NewInt(0).Div(balOMGScTradeS34, big.NewInt(1000000000))
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncOMGTokenIDStr,
		withdrawingPOMG,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxIDOMG, foundOMG := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, foundOMG)
	time.Sleep(60 * time.Second)

	withdrawingPZIL := big.NewInt(0).Div(balZILScS34, big.NewInt(1000))
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncZILTokenIDStr,
		withdrawingPZIL,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxIDZIL, foundZIL := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, foundZIL)
	time.Sleep(60 * time.Second)

	withdrawingPKBN := big.NewInt(0).Div(balKBNScS34, big.NewInt(1000000000))
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncKBNTokenIDStr,
		withdrawingPKBN,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxIDKBN, foundKBN := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, foundKBN)
	time.Sleep(60 * time.Second)

	withdrawingPSALT := balSALTScS34
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncSALTTokenIDStr,
		withdrawingPSALT,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxIDSALT, foundSALT := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, foundSALT)
	time.Sleep(100 * time.Second)

	tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))
	tradingSuite.submitBurnProofForWithdrawal(burningTxIDOMG.(string))
	tradingSuite.submitBurnProofForWithdrawal(burningTxIDZIL.(string))
	tradingSuite.submitBurnProofForWithdrawal(burningTxIDKBN.(string))
	tradingSuite.submitBurnProofForWithdrawal(burningTxIDSALT.(string))

	fmt.Println("------------ ETH --------------")
	balEth := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance  : ", balEth)

	balEthSc := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC : ", balEthSc)

	balpEth, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance  : ", balpEth)

	fmt.Println("------------ OMG --------------")

	balOMG := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.OMGAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] OMG balance  : ", balOMG)

	balOMGSc := tradingSuite.getDepositedBalance(
		tradingSuite.OMGAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] OMG balance  on SC : ", balOMGSc)

	balpOMG, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncOMGTokenIDStr)
	fmt.Println("[INFO] pOMG balance  : ", balpOMG)

	fmt.Println("------------ ZIL --------------")

	balZIL := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.ZILAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ZIL balance  : ", balZIL)

	balZILSc := tradingSuite.getDepositedBalance(
		tradingSuite.ZILAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ZIL balance  on SC : ", balZILSc)

	balpZIL, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncZILTokenIDStr)
	fmt.Println("[INFO] pZIL balance  : ", balpZIL)

	fmt.Println("------------ KBN --------------")

	balKBN := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.KBNAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] KBN balance  : ", balKBN)

	balKBNSc := tradingSuite.getDepositedBalance(
		tradingSuite.KBNAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] KBN balance  on SC : ", balKBNSc)

	balpKBN, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncKBNTokenIDStr)
	fmt.Println("[INFO] pKBN balance  : ", balpKBN)

	fmt.Println("------------ SALT --------------")

	balSALT := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.SALTAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] SALT balance  : ", balSALT)

	balSALTSc := tradingSuite.getDepositedBalance(
		tradingSuite.SALTAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] SALT balance  on SC : ", balSALTSc)

	balpSALT, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncSALTTokenIDStr)
	fmt.Println("[INFO] pSALT balance  : ", balpSALT)

}

func (tradingSuite *KovanKyberTradingTestSuite) Test8DepositAndWithdrwaERC20tokenOMG() {
return
	fmt.Println("============ TEST 8 DEPOSIT AND WITHDRAW ERC20 TOKEN (OMG) ===========")
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

	balpBATInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncOMGTokenIDStr)
	fmt.Println("[INFO] pOMG balance initialization : ", balpBATInit)

	balBATInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.OMGAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] OMG balance initialization : ", balBATInit)

	balEthScInit := tradingSuite.getDepositedBalance(
		tradingSuite.OMGAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] OMG balance initialization on SC : ", balEthScInit)

	fmt.Println("------------ STEP 1: porting OMG to pOMG --------------")

	fmt.Println("amount OMG deposit input : ", (big.NewInt(int64(0.5 * params.Ether))))
	deposit := big.NewInt(int64(0.5 * params.Ether))
	burningPETH := big.NewInt(0).Div(deposit, big.NewInt(1000000000))
	// Deposit to proof
	txHash := tradingSuite.depositERC20ToBridge(
		deposit,
		common.HexToAddress(tradingSuite.OMGAddressStr),
		tradingSuite.IncPaymentAddrStr,
	)

	balOMGAfDep := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.OMGAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] OMG balance remain after deposit : ", balOMGAfDep)

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
		tradingSuite.IncOMGTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)

	time.Sleep(120 * time.Second)

	balpBATAfIssS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncOMGTokenIDStr)
	fmt.Println("[INFO] pOMG balance after issuing step 1 : ", balpBATAfIssS1)
	// check PRV and token balance after issuing
	balPrvAfIssS1, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after issuing step 1: ", balPrvAfIssS1)

	fmt.Println("------------ STEP 2: burning pOMG to deposit OMG to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningPToken(
		tradingSuite.IncOMGTokenIDStr,
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

	balpBATAfBurnS2, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncOMGTokenIDStr)
	fmt.Println("[INFO] pOMG balance after burning step 2 : ", balpBATAfBurnS2)
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
		tradingSuite.OMGAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] OMG balance after deposit on SC at step 2: ", balBATScDepS2)
	// TODO assert ETH balane on SC

	fmt.Println("------------ STEP 3: withdraw OMG to deposit pOMG to Incognito  --------------")

	txHashByEmittingWithdrawalReq1 := tradingSuite.requestWithdraw(
		tradingSuite.OMGAddressStr,
		balBATScDepS2,
	)

	balBATScDepS41 := tradingSuite.getDepositedBalance(
		tradingSuite.OMGAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] OMG balance on SC after withdraw at step 3 : ", balBATScDepS41)
	// TODO assert BAT balane on SC

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq1)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(100 * time.Second)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncOMGTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(100 * time.Second)

	balpBATAfIssS41, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncOMGTokenIDStr,
	)
	fmt.Println("[INFO] pOMG balance after issuing step 3 : ", balpBATAfIssS41)

	fmt.Println("------------ STEP 4: withdraw pOMG to deposit OMG   --------------")

	withdrawingPBAT := big.NewInt(0).Div(balBATScDepS2, big.NewInt(1000000000))

	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncOMGTokenIDStr,
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
		tradingSuite.IncOMGTokenIDStr,
	)
	fmt.Println("[INFO] pOMG balance after burning step 4 : ", balpBATAfBurnS51)

	tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))

	balBATAfDep51 := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.OMGAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] OMG balance after withdraw  : ", balBATAfDep51)

}
