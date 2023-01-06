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
	value, ok := big.NewInt(0).SetString("10000000000000000", 10)
	require.Equal(v2.T(), true, ok)
	auth.Value = value
	salt, ok := big.NewInt(0).SetString("0x360c6ebe0000000000000000000000000000000000000000b73405421208f37d"[2:], 16)
	require.Equal(v2.T(), true, ok)
	offerId, ok := big.NewInt(0).SetString("0", 10)
	require.Equal(v2.T(), true, ok)
	signature, err := hex.DecodeString("0xb6e32411b8d2419a67856fec922915a4c11742095410a732d229a699beac6cd25e4cd13d1605a7b3912d5e54a6e641db8832612a9d8ad8841f160bb42a081e7b1b"[2:])
	require.Equal(v2.T(), nil, err)

	//basicOrder := opensea.BasicOrderParameters{
	//	ConsiderationToken:                common.HexToAddress("0x0000000000000000000000000000000000000000"),
	//	ConsiderationIdentifier:           big.NewInt(0),
	//	ConsiderationAmount:               big.NewInt(9750000000000000),
	//	Offerer:                           common.HexToAddress("0x50e5ca66d250bb3b1e52f7eaade8af016ad2e147"),
	//	Zone:                              common.HexToAddress("0x0000000000000000000000000000000000000000"),
	//	OfferToken:                        common.HexToAddress("0x8b0e17927a58392BBc42faeD1Cb41abE3A43A50C"),
	//	OfferIdentifier:                   offerId,
	//	OfferAmount:                       big.NewInt(1),
	//	BasicOrderType:                    0,
	//	StartTime:                         big.NewInt(1671919307),
	//	EndTime:                           big.NewInt(1687471307),
	//	ZoneHash:                          [32]byte{},
	//	Salt:                              salt,
	//	OffererConduitKey:                 toByte32(common.HexToHash("0x0000007b02230091a7ed01230072f7006a004d60a8d4e71d599b8104250f0000").Bytes()),
	//	FulfillerConduitKey:               toByte32(common.HexToHash("0x0000007b02230091a7ed01230072f7006a004d60a8d4e71d599b8104250f0000").Bytes()),
	//	TotalOriginalAdditionalRecipients: big.NewInt(1),
	//	AdditionalRecipients: []opensea.AdditionalRecipient{
	//		{
	//			Amount:    big.NewInt(250000000000000),
	//			Recipient: common.HexToAddress("0x0000a26b00c1F0DF003000390027140000fAa719"),
	//		},
	//	},
	//	Signature: signature,
	//}
	advanceOrder := opensea.AdvancedOrder{
		Parameters: opensea.OrderParameters{
			Offerer: common.HexToAddress("0x50e5ca66d250bb3b1e52f7eaade8af016ad2e147"),
			Zone:    common.HexToAddress("0x0000000000000000000000000000000000000000"),
			Offer: []opensea.OfferItem{
				{
					ItemType:             2,
					Token:                common.HexToAddress("0x8b0e17927a58392BBc42faeD1Cb41abE3A43A50C"),
					IdentifierOrCriteria: offerId,
					StartAmount:          big.NewInt(1),
					EndAmount:            big.NewInt(1),
				},
			},
			Consideration: []opensea.ConsiderationItem{
				{
					ItemType:             0,
					Token:                common.HexToAddress("0x0000000000000000000000000000000000000000"),
					IdentifierOrCriteria: big.NewInt(0),
					StartAmount:          big.NewInt(9750000000000000),
					EndAmount:            big.NewInt(9750000000000000),
					Recipient:            common.HexToAddress("0x50E5ca66d250BB3B1e52F7EAAde8aF016AD2E147"),
				},
				{
					ItemType:             0,
					Token:                common.HexToAddress("0x0000000000000000000000000000000000000000"),
					IdentifierOrCriteria: big.NewInt(0),
					StartAmount:          big.NewInt(250000000000000),
					EndAmount:            big.NewInt(250000000000000),
					Recipient:            common.HexToAddress("0x0000a26b00c1F0DF003000390027140000fAa719"),
				},
			},
			OrderType:                       0,
			StartTime:                       big.NewInt(1671919307),
			EndTime:                         big.NewInt(1687471307),
			ZoneHash:                        [32]byte{},
			Salt:                            salt,
			ConduitKey:                      toByte32(common.HexToHash("0x0000007b02230091a7ed01230072f7006a004d60a8d4e71d599b8104250f0000").Bytes()),
			TotalOriginalConsiderationItems: big.NewInt(2),
		},
		Numerator:   big.NewInt(1),
		Denominator: big.NewInt(1),
		Signature:   signature,
		ExtraData:   []byte{},
	}
	fulfillerConduitKey := toByte32(common.HexToHash("0x0000007b02230091a7ed01230072f7006a004d60a8d4e71d599b8104250f0000").Bytes())
	// address will receive nft
	reciepentAddr := common.HexToAddress("0x126748A0144968DD14b0187B906dE62378c59067")

	iopenseaAbi, _ := abi.JSON(strings.NewReader(opensea.IopenseaMetaData.ABI))
	calldata, err := iopenseaAbi.Pack("fulfillAdvancedOrder", advanceOrder, []opensea.CriteriaResolver{}, fulfillerConduitKey, reciepentAddr)
	require.Equal(v2.T(), nil, err)
	fmt.Println("build opensea calldata: " + hex.EncodeToString(calldata) + "\n")
	//tx, err := v2.OpenseaProxy.Forward(auth, calldata)
	//require.Equal(v2.T(), nil, err)

	//openseaAddr := common.HexToAddress("0x00000000006c3852cbEf3e08E8dF289169EdE581")
	//iopensea, _ := opensea.NewIopensea(openseaAddr, v2.ETHClient)
	//auth.GasLimit = 1000000
	//tx, err := iopensea.FulfillBasicOrder(auth, basicOrder)
	//require.Equal(v2.T(), nil, err)

	//fmt.Println(tx.Hash().String())

	// final burncalldata
	openseaProxyAbi, _ := abi.JSON(strings.NewReader(opensea.OpenseaMetaData.ABI))
	calldata, err = openseaProxyAbi.Pack("forward", calldata)
	require.Equal(v2.T(), nil, err)
	fmt.Println("final burn calldata: " + hex.EncodeToString(calldata) + "\n")
}
