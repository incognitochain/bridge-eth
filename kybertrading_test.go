   package main

// import (
// 	"fmt"
// 	"math/big"
// 	"strings"
// 	"testing"
// 	"time"

// 	kbnmultiTrade "github.com/incognitochain/bridge-eth/bridge/kbnmultitrade"

// 	"github.com/incognitochain/bridge-eth/bridge/kbntrade"
// 	"github.com/incognitochain/bridge-eth/bridge/vault"
// 	"github.com/stretchr/testify/suite"

// 	"github.com/ethereum/go-ethereum/accounts/abi"
// 	"github.com/ethereum/go-ethereum/accounts/abi/bind"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/crypto"
// 	"github.com/ethereum/go-ethereum/params"

// 	"github.com/stretchr/testify/require"
// )

// // Define the suite, and absorb the built-in basic suite
// // functionality from testify - including assertion methods.
// type KyberTradingTestSuite struct {
// 	*TradingTestSuite

// 	KyberTradeDeployedAddr      common.Address
// 	KyberMultiTradeDeployedAddr common.Address
// 	KyberContractAddr           common.Address
// 	WETHAddr                    common.Address

// 	IncKBNTokenIDStr  string
// 	IncSALTTokenIDStr string
// 	IncOMGTokenIDStr  string
// 	IncSNTTokenIDStr  string

// 	EtherAddressStrKyber string
// 	KBNAddressStr        string
// 	SALTAddressStr       string
// 	OMGAddressStr        string
// 	SNTAddressStr        string

// 	// token amounts for tests
// 	DepositingEther       float64
// 	KBNBalanceAfterStep1  *big.Int
// 	SALTBalanceAfterStep2 *big.Int
// }

// func NewKyberTradingTestSuite(tradingTestSuite *TradingTestSuite) *KyberTradingTestSuite {
// 	return &KyberTradingTestSuite{
// 		TradingTestSuite: tradingTestSuite,
// 	}
// }

// // Make sure that VariableThatShouldStartAtFive is set to five
// // before each test
// func (tradingSuite *KyberTradingTestSuite) SetupSuite() {
// 	fmt.Println("Setting up the suite...")
// 	// Kovan env
// 	tradingSuite.EtherAddressStrKyber = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"

// 	tradingSuite.IncKBNTokenIDStr = "d6644f62d0ef0475335ae7bb6103f358979cbfcd2b85481e3bde2b82101a095c"
// 	tradingSuite.IncSALTTokenIDStr = "06ce44eae35daf57b9b8158ab95c0cddda9bac208fc380236a318ef40f3ac2ef"
// 	tradingSuite.IncOMGTokenIDStr = "27322fa7fce2c4d4d5a0022d595a0eec56d7735751a3ba8bc7f10b358ab938bc"
// 	tradingSuite.IncSNTTokenIDStr = "414a6459526e827321cedb6e574d2ba2eb267c5735b0a65991602a405fb753b7"
// 	tradingSuite.KBNAddressStr = "0xad67cB4d63C9da94AcA37fDF2761AaDF780ff4a2"                                    // kovan
// 	tradingSuite.SALTAddressStr = "0x6fEE5727EE4CdCBD91f3A873ef2966dF31713A04"                                   // kovan
// 	tradingSuite.OMGAddressStr = "0xdB7ec4E4784118D9733710e46F7C83fE7889596a"                                    // kovan
// 	tradingSuite.SNTAddressStr = "0x4c99B04682fbF9020Fcb31677F8D8d66832d3322"                                    // kovan
// 	tradingSuite.KyberTradeDeployedAddr = common.HexToAddress("0x3AfefD86A75f11da0Dc4537698Fa598Ae1b38470")      //kovan
// 	tradingSuite.KyberMultiTradeDeployedAddr = common.HexToAddress("0x77D6da98C726B935F1b5780C8B3FDBBBe5A20aAF") //kovan
// 	tradingSuite.DepositingEther = float64(0.01)
// 	tradingSuite.KyberContractAddr = common.HexToAddress("0x692f391bCc85cefCe8C237C01e1f636BbD70EA4D") // kovan
// }

// func (tradingSuite *KyberTradingTestSuite) TearDownSuite() {
// 	fmt.Println("Tearing down the suite...")
// 	tradingSuite.ETHClient.Close()
// }

// func (tradingSuite *KyberTradingTestSuite) SetupTest() {
// 	fmt.Println("Setting up the test...")
// }

// func (tradingSuite *KyberTradingTestSuite) TearDownTest() {
// 	fmt.Println("Tearing down the test...")
// }

// // In order for 'go test' to run this suite, we need to create
// // a normal test function and pass our suite to suite.Run
// func TestKyberTradingTestSuite(t *testing.T) {
// 	fmt.Println("Starting entry point for Kyber test suite...")

// 	tradingSuite := new(TradingTestSuite)
// 	suite.Run(t, tradingSuite)

// 	kyberTradingSuite := NewKyberTradingTestSuite(tradingSuite)
// 	suite.Run(t, kyberTradingSuite)

// 	fmt.Println("Finishing entry point for 0x test suite...")
// }

// func (tradingSuite *KyberTradingTestSuite) getExpectedRate(
// 	srcToken string,
// 	destToken string,
// 	srcQty *big.Int,
// ) *big.Int {
// 	if srcToken == tradingSuite.EtherAddressStr {
// 		srcToken = tradingSuite.EtherAddressStrKyber
// 	}
// 	if destToken == tradingSuite.EtherAddressStr {
// 		destToken = tradingSuite.EtherAddressStrKyber
// 	}
// 	c, err := kbntrade.NewKbntrade(tradingSuite.KyberTradeDeployedAddr, tradingSuite.ETHClient)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	expectRate, slippageRate, err := c.GetConversionRates(nil, common.HexToAddress(srcToken), srcQty, common.HexToAddress(destToken))
// 	require.Equal(tradingSuite.T(), nil, err)
// 	fmt.Printf("slippageRate value: %d\n", slippageRate)
// 	fmt.Printf("expectRate value: %d\n", expectRate)
// 	return expectRate
// }

// func (tradingSuite *KyberTradingTestSuite) executeWithKyber(
// 	srcQty *big.Int,
// 	srcTokenIDStr string,
// 	destTokenIDStr string,
// ) {
// 	tradeAbi, _ := abi.JSON(strings.NewReader(kbntrade.KbntradeABI))

// 	// Get contract instance
// 	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	auth := bind.NewKeyedTransactor(tradingSuite.ETHPrivKey)
// 	auth.GasPrice = big.NewInt(50000000000)
// 	auth.GasLimit = 2000000
// 	srcToken := common.HexToAddress(srcTokenIDStr)
// 	destToken := common.HexToAddress(destTokenIDStr)
// 	auth.GasPrice = big.NewInt(50000000000)
// 	auth.GasLimit = 2000000
// 	expectRate := tradingSuite.getExpectedRate(srcTokenIDStr, destTokenIDStr, srcQty)
// 	input, _ := tradeAbi.Pack("trade", srcToken, srcQty, destToken, expectRate)
// 	timestamp := []byte(randomizeTimestamp())
// 	tempData := append(tradingSuite.KyberTradeDeployedAddr[:], input...)
// 	tempData1 := append(tempData, timestamp...)
// 	data := rawsha3(tempData1)
// 	signBytes, _ := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)

// 	tx, err := c.Execute(
// 		auth,
// 		srcToken,
// 		srcQty,
// 		destToken,
// 		tradingSuite.KyberTradeDeployedAddr,
// 		input,
// 		timestamp,
// 		signBytes,
// 	)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	txHash := tx.Hash()
// 	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
// 		require.Equal(tradingSuite.T(), nil, err)
// 	}
// 	fmt.Printf("Kyber trade executed , txHash: %x\n", txHash[:])
// }

// func (tradingSuite *KyberTradingTestSuite) executeMultiTradeWithKyber(
// 	srcQties []*big.Int,
// 	srcTokenIDStrs []string,
// 	destTokenIDStrs []string,
// ) {
// 	tradeAbi, _ := abi.JSON(strings.NewReader(kbnmultiTrade.KbnmultiTradeABI))
// 	auth := bind.NewKeyedTransactor(tradingSuite.ETHPrivKey)
// 	auth.GasPrice = big.NewInt(50000000000)
// 	auth.GasLimit = 2000000
// 	// Deploy kbnMultitrade
// 	// kbnMultiTradeAddr, tx, _, err := kbnmultiTrade.DeployKbnmultiTrade(auth, tradingSuite.ETHClient, tradingSuite.KyberContractAddr, tradingSuite.VaultAddr)
// 	// require.Equal(tradingSuite.T(), nil, err)
// 	// fmt.Println("deployed kbnMultitrade")
// 	// fmt.Printf("addr: %s\n", kbnMultiTradeAddr.Hex())
// 	// tradingSuite.KyberMultiTradeDeployedAddr = kbnMultiTradeAddr

// 	// Get contract instance
// 	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	auth.GasPrice = big.NewInt(50000000000)
// 	auth.GasLimit = 5000000
// 	sourceAddresses := make([]common.Address, 0)
// 	for _, p := range srcTokenIDStrs {
// 		sourceAddresses = append(sourceAddresses, common.HexToAddress(p))
// 	}
// 	destAddresses := make([]common.Address, 0)
// 	for _, p := range destTokenIDStrs {
// 		destAddresses = append(destAddresses, common.HexToAddress(p))
// 	}
// 	expectRates := make([]*big.Int, 0)
// 	for i := range destTokenIDStrs {
// 		expectRates = append(expectRates, tradingSuite.getExpectedRate(srcTokenIDStrs[i], destTokenIDStrs[i], srcQties[i]))
// 	}

// 	input, _ := tradeAbi.Pack("trade", sourceAddresses, srcQties, destAddresses, expectRates)
// 	timestamp := []byte(randomizeTimestamp())
// 	tempData := append(tradingSuite.KyberMultiTradeDeployedAddr[:], input...)
// 	tempData1 := append(tempData, timestamp...)
// 	data := rawsha3(tempData1)
// 	signBytes, _ := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)

// 	tx, err := c.ExecuteMulti(
// 		auth,
// 		sourceAddresses,
// 		srcQties,
// 		destAddresses,
// 		tradingSuite.KyberMultiTradeDeployedAddr,
// 		input,
// 		timestamp,
// 		signBytes,
// 	)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	txHash := tx.Hash()
// 	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
// 		require.Equal(tradingSuite.T(), nil, err)
// 	}
// 	fmt.Printf("Kyber multi trade executed , txHash: %x\n", txHash[:])
// }

// func (tradingSuite *KyberTradingTestSuite) Test1TradeEthForKBNWithKyber() {

// 	fmt.Println("============ TEST TRADE ETHER FOR KBN WITH Kyber AGGREGATOR ===========")
// 	fmt.Println("------------ STEP 0: declaration & initialization --------------")
// 	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
// 	burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))

// 	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
// 	fmt.Println("------------ STEP 1: porting ETH to pETH --------------")
// 	txHash := tradingSuite.depositETH(
// 		tradingSuite.DepositingEther,
// 		tradingSuite.IncPaymentAddrStr,
// 	)
// 	// time.Sleep(15 * time.Second)
// 	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

// 	fmt.Println("Waiting 90s for 15 blocks confirmation")
// 	time.Sleep(90 * time.Second)
// 	_, err = tradingSuite.callIssuingETHReq(
// 		tradingSuite.IncEtherTokenIDStr,
// 		ethDepositProof,
// 		ethBlockHash,
// 		ethTxIdx,
// 	)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	time.Sleep(120 * time.Second)

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
// 	time.Sleep(120 * time.Second)

// 	tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))
// 	deposited := tradingSuite.getDepositedBalance(
// 		tradingSuite.EtherAddressStr,
// 		pubKeyToAddrStr,
// 	)
// 	fmt.Println("deposited EHT: ", deposited)
// 	// require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPETH, big.NewInt(1000000000)), deposited)

// 	fmt.Println("------------ step 3: execute trade ETH for KBN through Kyber aggregator --------------")
// 	tradingSuite.executeWithKyber(
// 		tradeAmount,
// 		tradingSuite.EtherAddressStr,
// 		tradingSuite.KBNAddressStr,
// 	)
// 	time.Sleep(15 * time.Second)
// 	kbnTraded := tradingSuite.getDepositedBalance(
// 		tradingSuite.KBNAddressStr,
// 		pubKeyToAddrStr,
// 	)
// 	fmt.Println("kbnTraded: ", kbnTraded)

// 	fmt.Println("------------ step 4: withdrawing KBN from SC to pKBN on Incognito --------------")
// 	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
// 		tradingSuite.KBNAddressStr,
// 		kbnTraded,
// 	)
// 	time.Sleep(15 * time.Second)

// 	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

// 	fmt.Println("Waiting 90s for 15 blocks confirmation")
// 	time.Sleep(90 * time.Second)
// 	_, err = tradingSuite.callIssuingETHReq(
// 		tradingSuite.IncKBNTokenIDStr,
// 		ethDepositProof,
// 		ethBlockHash,
// 		ethTxIdx,
// 	)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	time.Sleep(120 * time.Second)

// 	fmt.Println("------------ step 5: withdrawing pKBN from Incognito to KBN --------------")
// 	withdrawingPKBN := big.NewInt(0).Div(kbnTraded, big.NewInt(1000000000))
// 	burningRes, err = tradingSuite.callBurningPToken(
// 		tradingSuite.IncKBNTokenIDStr,
// 		withdrawingPKBN,
// 		tradingSuite.ETHOwnerAddrStr,
// 		"createandsendburningrequest",
// 	)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	burningTxID, found = burningRes["TxID"]
// 	require.Equal(tradingSuite.T(), true, found)
// 	time.Sleep(120 * time.Second)

// 	tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))

// 	bal := tradingSuite.getBalanceOnETHNet(
// 		common.HexToAddress(tradingSuite.KBNAddressStr),
// 		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
// 	)
// 	tradingSuite.KBNBalanceAfterStep1 = bal
// 	fmt.Println("KBN balance after step 1: ", tradingSuite.KBNBalanceAfterStep1)
// 	// require.Equal(tradingSuite.T(), withdrawingPKBN.Uint64(), bal.Div(bal, big.NewInt(1000000000)).Uint64())
// }

// func (tradingSuite *KyberTradingTestSuite) Test2TradeKBNForSALTWithKyber() {
// 	fmt.Println("============ TEST TRADE KBN FOR SALT WITH KYBER AGGREGATOR ===========")
// 	fmt.Println("------------ step 0: declaration & initialization --------------")
// 	depositingKBN := tradingSuite.KBNBalanceAfterStep1
// 	burningPKBN := big.NewInt(0).Div(depositingKBN, big.NewInt(1000000000))
// 	tradeAmount := depositingKBN

// 	kbnbal := tradingSuite.getBalanceOnETHNet(
// 		common.HexToAddress(tradingSuite.KBNAddressStr),
// 		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
// 	)
// 	fmt.Println("kbn balance of owner: ", kbnbal)

// 	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
// 	fmt.Println("pubKeyToAddrStr: ", pubKeyToAddrStr)

// 	fmt.Println("------------ step 1: porting KBN to pKBN --------------")
// 	txHash := tradingSuite.depositERC20ToBridge(
// 		depositingKBN,
// 		common.HexToAddress(tradingSuite.KBNAddressStr),
// 		tradingSuite.IncPaymentAddrStr,
// 	)

// 	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

// 	fmt.Println("Waiting 90s for 15 blocks confirmation")
// 	time.Sleep(90 * time.Second)

// 	_, err = tradingSuite.callIssuingETHReq(
// 		tradingSuite.IncKBNTokenIDStr,
// 		ethDepositProof,
// 		ethBlockHash,
// 		ethTxIdx,
// 	)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	time.Sleep(120 * time.Second)
// 	// TODO: check the new balance on peth of the incognito account

// 	fmt.Println("------------ step 2: burning pKBN to deposit KBN to SC --------------")

// 	// make a burn tx to incognito chain as a result of deposit to SC
// 	burningRes, err := tradingSuite.callBurningPToken(
// 		tradingSuite.IncKBNTokenIDStr,
// 		burningPKBN,
// 		pubKeyToAddrStr[2:],
// 		"createandsendburningfordeposittoscrequest",
// 	)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	burningTxID, found := burningRes["TxID"]
// 	require.Equal(tradingSuite.T(), true, found)
// 	time.Sleep(120 * time.Second)

// 	tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))
// 	deposited := tradingSuite.getDepositedBalance(
// 		tradingSuite.KBNAddressStr,
// 		pubKeyToAddrStr,
// 	)
// 	fmt.Println("deposited kbn: ", deposited)
// 	// require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPKBN, big.NewInt(1000000000)), deposited)

// 	fmt.Println("------------ step 3: execute trade KBN for SALT through Kyber aggregator --------------")
// 	tradingSuite.executeWithKyber(
// 		tradeAmount,
// 		tradingSuite.KBNAddressStr,
// 		tradingSuite.SALTAddressStr,
// 	)
// 	time.Sleep(15 * time.Second)
// 	saltTraded := tradingSuite.getDepositedBalance(
// 		tradingSuite.SALTAddressStr,
// 		pubKeyToAddrStr,
// 	)
// 	fmt.Println("saltTraded: ", saltTraded)

// 	fmt.Println("------------ step 4: withdrawing SALT from SC to pSALT on Incognito --------------")
// 	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
// 		tradingSuite.SALTAddressStr,
// 		saltTraded,
// 	)
// 	time.Sleep(15 * time.Second)

// 	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

// 	fmt.Println("Waiting 90s for 15 blocks confirmation")
// 	time.Sleep(90 * time.Second)

// 	_, err = tradingSuite.callIssuingETHReq(
// 		tradingSuite.IncSALTTokenIDStr,
// 		ethDepositProof,
// 		ethBlockHash,
// 		ethTxIdx,
// 	)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	time.Sleep(120 * time.Second)

// 	fmt.Println("------------ step 5: withdrawing pSALT from Incognito to SALT --------------")
// 	withdrawingPSALT := saltTraded // SALT decimal 8
// 	burningRes, err = tradingSuite.callBurningPToken(
// 		tradingSuite.IncSALTTokenIDStr,
// 		withdrawingPSALT,
// 		tradingSuite.ETHOwnerAddrStr,
// 		"createandsendburningrequest",
// 	)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	burningTxID, found = burningRes["TxID"]
// 	require.Equal(tradingSuite.T(), true, found)
// 	time.Sleep(120 * time.Second)

// 	tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))

// 	bal := tradingSuite.getBalanceOnETHNet(
// 		common.HexToAddress(tradingSuite.SALTAddressStr),
// 		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
// 	)
// 	tradingSuite.SALTBalanceAfterStep2 = bal
// 	fmt.Println("SALT balance after step 2: ", tradingSuite.SALTBalanceAfterStep2)
// 	// require.Equal(tradingSuite.T(), withdrawingPSALT.Uint64(), bal.Uint64())
// 	// require.Equal(tradingSuite.T(), withdrawingPSALT.Uint64(), bal.Div(bal, big.NewInt(1000000000)).Uint64())
// }

// func (tradingSuite *KyberTradingTestSuite) Test3TradeSALTForEthWithKyber() {
// 	fmt.Println("============ TEST TRADE SALT FOR ETH WITH KYBER AGGREGATOR ===========")
// 	fmt.Println("------------ step 0: declaration & initialization --------------")
// 	depositingSALT := tradingSuite.SALTBalanceAfterStep2
// 	burningPSALT := depositingSALT // SALT decimal 8
// 	tradeAmount := depositingSALT

// 	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
// 	fmt.Println("pubKeyToAddrStr: ", pubKeyToAddrStr)

// 	fmt.Println("------------ step 1: porting SALT to pSALT --------------")
// 	txHash := tradingSuite.depositERC20ToBridge(
// 		depositingSALT,
// 		common.HexToAddress(tradingSuite.SALTAddressStr),
// 		tradingSuite.IncPaymentAddrStr,
// 	)

// 	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

// 	fmt.Println("Waiting 90s for 15 blocks confirmation")
// 	time.Sleep(90 * time.Second)

// 	issuingRes, err := tradingSuite.callIssuingETHReq(
// 		tradingSuite.IncSALTTokenIDStr,
// 		ethDepositProof,
// 		ethBlockHash,
// 		ethTxIdx,
// 	)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	fmt.Println("issuingRes: ", issuingRes)
// 	time.Sleep(120 * time.Second)

// 	fmt.Println("------------ step 2: burning pSALT to deposit SALT to SC --------------")
// 	// make a burn tx to incognito chain as a result of deposit to SC
// 	burningRes, err := tradingSuite.callBurningPToken(
// 		tradingSuite.IncSALTTokenIDStr,
// 		burningPSALT,
// 		pubKeyToAddrStr[2:],
// 		"createandsendburningfordeposittoscrequest",
// 	)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	burningTxID, found := burningRes["TxID"]
// 	require.Equal(tradingSuite.T(), true, found)
// 	time.Sleep(140 * time.Second)

// 	tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))
// 	deposited := tradingSuite.getDepositedBalance(
// 		tradingSuite.SALTAddressStr,
// 		pubKeyToAddrStr,
//   )
//   fmt.Println("SAL balance in SC :", deposited)
// 	//require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPSALT, big.NewInt(1000000000)), deposited)

// 	fmt.Println("------------ step 3: execute trade SALT for ETH through Kyber aggregator --------------")
// 	tradingSuite.executeWithKyber(
// 		tradeAmount,
// 		tradingSuite.SALTAddressStr,
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
// 	time.Sleep(15 * time.Second)

// 	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

// 	fmt.Println("Waiting 90s for 15 blocks confirmation")
// 	time.Sleep(90 * time.Second)

// 	_, err = tradingSuite.callIssuingETHReq(
// 		tradingSuite.IncEtherTokenIDStr,
// 		ethDepositProof,
// 		ethBlockHash,
// 		ethTxIdx,
// 	)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	time.Sleep(140 * time.Second)

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
// 	time.Sleep(140 * time.Second)

// 	tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))

// 	bal := tradingSuite.getBalanceOnETHNet(
// 		common.HexToAddress(tradingSuite.EtherAddressStr),
// 		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
// 	)
// 	fmt.Println("Ether balance after step 3: ", bal)
// 	// require.Equal(tradingSuite.T(), withdrawingPETH.Uint64(), bal.Div(bal, big.NewInt(1000000000)).Uint64())
// }

// func (tradingSuite *KyberTradingTestSuite) Test4MultiTradeWithKyber() {
// 	fmt.Println("============ TEST TRADE ETHER FOR KBN WITH Kyber AGGREGATOR ===========")
// 	fmt.Println("------------ STEP 0: declaration & initialization --------------")
// 	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether * 2))
// 	tradeHaftAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
// 	burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))

// 	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
// 	fmt.Println("------------ STEP 1: porting ETH to pETH --------------")
// 	txHash := tradingSuite.depositETH(
// 		tradingSuite.DepositingEther*2,
// 		tradingSuite.IncPaymentAddrStr,
// 	)
// 	time.Sleep(15 * time.Second)
// 	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

// 	fmt.Println("Waiting 90s for 15 blocks confirmation")
// 	time.Sleep(90 * time.Second)
// 	_, err = tradingSuite.callIssuingETHReq(
// 		tradingSuite.IncEtherTokenIDStr,
// 		ethDepositProof,
// 		ethBlockHash,
// 		ethTxIdx,
// 	)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	time.Sleep(120 * time.Second)

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
// 	time.Sleep(120 * time.Second)

// 	tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))
// 	deposited := tradingSuite.getDepositedBalance(
// 		tradingSuite.EtherAddressStr,
// 		pubKeyToAddrStr,
// 	)
// 	fmt.Println("deposited EHT: ", deposited)

// 	fmt.Println("------------ step 3: execute trade ETH for KBN through Kyber aggregator --------------")
// 	tradingSuite.executeWithKyber(
// 		tradeHaftAmount,
// 		tradingSuite.EtherAddressStr,
// 		tradingSuite.KBNAddressStr,
// 	)
// 	time.Sleep(20 * time.Second)
// 	kbnTraded := tradingSuite.getDepositedBalance(
// 		tradingSuite.KBNAddressStr,
// 		pubKeyToAddrStr,
// 	)
// 	tradingSuite.getDepositedBalance(
// 		tradingSuite.EtherAddressStr,
// 		pubKeyToAddrStr,
// 	)
// 	fmt.Println("kbnTraded: ", kbnTraded)

// 	fmt.Println("------------ step 4: execute trade ETH and KBN for OMG and SALT and through Kyber aggregator --------------")

// 	tradingSuite.executeMultiTradeWithKyber(
// 		[]*big.Int{kbnTraded, tradeHaftAmount},
// 		[]string{tradingSuite.KBNAddressStr, tradingSuite.EtherAddressStr},
// 		[]string{tradingSuite.OMGAddressStr, tradingSuite.SALTAddressStr},
// 	)
// 	time.Sleep(15 * time.Second)

// 	omgTraded := tradingSuite.getDepositedBalance(
// 		tradingSuite.OMGAddressStr,
// 		pubKeyToAddrStr,
// 	)
// 	fmt.Println("omgTraded: ", omgTraded)

// 	saltTraded := tradingSuite.getDepositedBalance(
// 		tradingSuite.SALTAddressStr,
// 		pubKeyToAddrStr,
// 	)
// 	fmt.Println("saltTraded: ", saltTraded)
// }
