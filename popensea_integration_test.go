package bridge

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/incognitochain/bridge-eth/bridge/opensea"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"math/big"
	"strings"
	"testing"
)

// // Define the suite, and absorb the built-in basic suite
// // functionality from testify - including assertion methods.
type POpenseaIntegrationTestSuite struct {
	*TradingTestSuite

	OpenseaProxy  *opensea.Opensea
	withdrawer    common.Address
	auth          *bind.TransactOpts
	EtherAddress  common.Address
	ETHHost       string
	ETHClient     *ethclient.Client
	IncHost       string
	ETHPrivKeyStr string
	ETHPrivKey    *ecdsa.PrivateKey
}

func NewPOpenseaIntegrationTestSuite(tradingTestSuite *TradingTestSuite) *POpenseaIntegrationTestSuite {
	return &POpenseaIntegrationTestSuite{
		TradingTestSuite: tradingTestSuite,
	}
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (v2 *POpenseaIntegrationTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")

	v2.IncHost = "http://127.0.0.1:9338"
	//	deployed prv vote
	//addr: 0xa0904E2F05D1108063Ac7CfB7719bD6518FDBDF4
	//	deployed governance
	//addr: 0xD58d36a9053BbB69a75C4F9AF6864164dcbD2Cdb
	v2.ETHHost = "https://goerli.infura.io/v3/1138a1e99b154b10bae5c382ad894361"
	client, err := ethclient.Dial(v2.ETHHost)
	require.Equal(v2.T(), nil, err)
	v2.ETHClient = client

	ops, err := opensea.NewOpensea(common.HexToAddress("0x1A2944620a2b8B899bbC5A178A578146C3231a00"), v2.ETHClient)
	require.Equal(v2.T(), nil, err)
	v2.OpenseaProxy = ops

	v2.ETHPrivKeyStr = "1193a43543fc11e37daa1a026ae8fae69d84c5dd1f3f933047ff2588778c5cca"
	privKey, err := crypto.HexToECDSA(v2.ETHPrivKeyStr)
	require.Equal(v2.T(), nil, err)
	v2.ETHPrivKey = privKey
}

func (v2 *POpenseaIntegrationTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
}

func (v2 *POpenseaIntegrationTestSuite) SetupTest() {
	fmt.Println("Setting up the test...")
}

func (v2 *POpenseaIntegrationTestSuite) TearDownTest() {
	fmt.Println("Tearing down the test...")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestPOpenseaIntegration(t *testing.T) {
	fmt.Println("Starting entry point for pdao test suite...")
	tradingSuite := new(TradingTestSuite)
	suite.Run(t, tradingSuite)

	uniswapTradingSuite := NewPOpenseaIntegrationTestSuite(tradingSuite)
	suite.Run(t, uniswapTradingSuite)

	fmt.Println("Finishing entry point for pdao test suite...")
}

func (v2 *POpenseaIntegrationTestSuite) TestPDAOCreateProp() {
	// burn prv token
	fmt.Println("------------ STEP 1: burning pPRV --------------")
	auth, err := bind.NewKeyedTransactorWithChainID(v2.ETHPrivKey, big.NewInt(int64(v2.ChainIDETH)))
	require.Equal(v2.T(), nil, err)
	value, ok := big.NewInt(0).SetString("1000000000000000", 10)
	require.Equal(v2.T(), true, ok)
	auth.Value = value
	salt, ok := big.NewInt(0).SetString("0x360c6ebe0000000000000000000000000000000000000000f6d621abfb415587"[2:], 16)
	require.Equal(v2.T(), true, ok)
	offerId, ok := big.NewInt(0).SetString("7", 10)
	require.Equal(v2.T(), true, ok)
	signature, err := hex.DecodeString("0x9245f238efb011fb78fa25af772102f7e70b66b003691277d0a60795de549e290e67be3cd98dac444c8ba3cc2e51d4076c34c35d4dc4b4278806a0be631f13da1b"[2:])
	require.Equal(v2.T(), nil, err)

	basicOrder := opensea.BasicOrderParameters{
		ConsiderationToken:                common.HexToAddress("0x0000000000000000000000000000000000000000"),
		ConsiderationIdentifier:           big.NewInt(0),
		ConsiderationAmount:               big.NewInt(975000000000000),
		Offerer:                           common.HexToAddress("0x0fbd0254d34b9f82c054ceb77515e6eacd80e3f9"),
		Zone:                              common.HexToAddress("0x0000000000000000000000000000000000000000"),
		OfferToken:                        common.HexToAddress("0xA04B2E948D96CE9AAf78b6C4bdd2a6eC54DB8279"),
		OfferIdentifier:                   offerId,
		OfferAmount:                       big.NewInt(1),
		BasicOrderType:                    0,
		StartTime:                         big.NewInt(1671852986),
		EndTime:                           big.NewInt(1674531386),
		ZoneHash:                          [32]byte{},
		Salt:                              salt,
		OffererConduitKey:                 toByte32(common.HexToHash("0x0000007b02230091a7ed01230072f7006a004d60a8d4e71d599b8104250f0000").Bytes()),
		FulfillerConduitKey:               toByte32(common.HexToHash("0x0000007b02230091a7ed01230072f7006a004d60a8d4e71d599b8104250f0000").Bytes()),
		TotalOriginalAdditionalRecipients: big.NewInt(1),
		AdditionalRecipients: []opensea.AdditionalRecipient{
			{
				Amount:    big.NewInt(25000000000000),
				Recipient: common.HexToAddress("0x0000a26b00c1F0DF003000390027140000fAa719"),
			},
		},
		Signature: signature,
	}
	tradeAbi, _ := abi.JSON(strings.NewReader(opensea.IopenseaMetaData.ABI))
	calldata, err := tradeAbi.Pack("fulfillBasicOrder", basicOrder)
	require.Equal(v2.T(), nil, err)

	tx, err := v2.OpenseaProxy.Forward(auth, calldata)
	require.Equal(v2.T(), nil, err)

	//openseaAddr := common.HexToAddress("0x00000000006c3852cbEf3e08E8dF289169EdE581")
	//iopensea, _ := opensea.NewIopensea(openseaAddr, v2.ETHClient)
	//auth.GasLimit = 1000000
	//tx, err := iopensea.FulfillBasicOrder(auth, basicOrder)
	//require.Equal(v2.T(), nil, err)

	fmt.Println(tx.Hash().String())
}
