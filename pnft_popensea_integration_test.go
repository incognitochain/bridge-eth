package bridge

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/incognitochain/bridge-eth/bridge/erc721"
	"github.com/incognitochain/bridge-eth/bridge/opensea"
	"github.com/incognitochain/bridge-eth/bridge/pnft"
	"github.com/incognitochain/bridge-eth/bridge/pnft/adapter"
	"github.com/incognitochain/bridge-eth/bridge/pnft/executiondelegate"
	policymanager "github.com/incognitochain/bridge-eth/bridge/pnft/policyManager"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"math/big"
	"os/exec"
	"strings"
	"testing"
)

// // Define the suite, and absorb the built-in basic suite
// // functionality from testify - including assertion methods.
type PNFTIntegrationTestSuite struct {
	*TradingTestSuite

	standardPolicy    common.Address
	pnftDeployer      *bind.TransactOpts
	executionDelegate *executiondelegate.Executiondelegate
	policyManager     *policymanager.Policymanager
	pnft              *pnft.BlurExchange
	pnftAddr          common.Address
	execDelegateAddr  common.Address
	incNFT            *erc721.Erc721
	incNFTAddr        common.Address
	adapter           *adapter.Adapter
	adapterAddr       common.Address
	Forwarder         *opensea.Opensea
	auth              *bind.TransactOpts
	EtherAddress      common.Address
	ETHHost           string
	ETHClient         *ethclient.Client
	IncHost           string
	ETHPrivKey        *ecdsa.PrivateKey
}

func NewPNFTIntegrationTestSuite(tradingTestSuite *TradingTestSuite) *PNFTIntegrationTestSuite {
	return &PNFTIntegrationTestSuite{
		TradingTestSuite: tradingTestSuite,
	}
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (v2 *PNFTIntegrationTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")

	err := exec.Command("/bin/bash", "-c",
		"solc @openzeppelin/=node_modules/@openzeppelin/ --base-path=$(pwd)/bridge --bin --abi --optimize --overwrite bridge/contracts/pnft/adapter.sol -o bridge/pnft/adapter").Run()
	require.Equal(v2.T(), nil, err)
	err = exec.Command("/bin/bash", "-c",
		"abigen --abi bridge/pnft/adapter/PNFTAdapter.abi --bin bridge/pnft/adapter/PNFTAdapter.bin --out bridge/pnft/adapter/adapter.go --pkg adapter").Run()

	v2.ETHHost = "https://eth-goerli.g.alchemy.com/v2/uaJphkFTwcgwaLUWLB8fEen0FqoXVj1N"
	client, err := ethclient.Dial(v2.ETHHost)
	require.Equal(v2.T(), nil, err)
	v2.ETHClient = client

	privKey, err := crypto.HexToECDSA(v2.ETHPrivKeyStr)
	require.Equal(v2.T(), nil, err)
	v2.ETHPrivKey = privKey

	// deploy adapter
	v2.adapterAddr = common.HexToAddress("0x82A65855a1c1CE7CAb090868e5217Db49ddd02b9")
	adapterInst, err := adapter.NewAdapter(v2.adapterAddr, v2.ETHClient)
	require.Equal(v2.T(), nil, err)
	v2.adapter = adapterInst
	v2.execDelegateAddr = common.HexToAddress("0xE11D03730472681FcD61e9aFc29A4dAa4f93f509")
	v2.standardPolicy = common.HexToAddress("0x7DFd97E5Add5A69a59d3eCD73Ef5c82b3DA3c8AC")
	v2.incNFTAddr = common.HexToAddress("0x1f419B9469D641D333805C4054CA3b65Af54d315")

	pnftInst, err := pnft.NewBlurExchange(common.HexToAddress("0x87E5Ffa37503487691c75359401080B1e2FBdE5E"), v2.ETHClient)
	require.Equal(v2.T(), nil, err)
	v2.pnft = pnftInst
}

func (v2 *PNFTIntegrationTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
}

func (v2 *PNFTIntegrationTestSuite) SetupTest() {
	fmt.Println("Setting up the test...")
}

func (v2 *PNFTIntegrationTestSuite) TearDownTest() {
	fmt.Println("Tearing down the test...")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestPNFTIntegration(t *testing.T) {
	fmt.Println("Starting entry point for pdao test suite...")
	tradingSuite := new(TradingTestSuite)
	suite.Run(t, tradingSuite)

	uniswapTradingSuite := NewPNFTIntegrationTestSuite(tradingSuite)
	suite.Run(t, uniswapTradingSuite)

	fmt.Println("Finishing entry point for pdao test suite...")
}

func (v2 *PNFTIntegrationTestSuite) TestPBlurCreateProp() {
	privKe2, err := crypto.HexToECDSA("15681448451d0a925d17408e6c3f33a3e1b5a60b89ab5096c2381a37ab58e234")
	require.Equal(v2.T(), nil, err)
	auth2, err := bind.NewKeyedTransactorWithChainID(privKe2, big.NewInt(int64(v2.ChainIDETH)))
	require.Equal(v2.T(), nil, err)

	auth, err := bind.NewKeyedTransactorWithChainID(v2.ETHPrivKey, big.NewInt(int64(v2.ChainIDETH)))
	require.Equal(v2.T(), nil, err)

	// get timestamp
	lastBlock, err := v2.ETHClient.BlockNumber(context.Background())
	require.Equal(v2.T(), nil, err)
	blockHeader, err := v2.ETHClient.BlockByNumber(context.Background(), big.NewInt(int64(lastBlock)))
	require.Equal(v2.T(), nil, err)

	// pnft data to swap
	order := pnft.Order{
		Side:           1,
		MatchingPolicy: v2.standardPolicy,
		Collection:     v2.incNFTAddr,
		PaymentToken:   common.Address{},
		TokenId:        big.NewInt(1520),
		Amount:         big.NewInt(1),
		Price:          big.NewInt(100000000),
		ListingTime:    big.NewInt(int64(blockHeader.Time() - 1)),
		ExpirationTime: big.NewInt(int64(blockHeader.Time()) + 10000),
		Fees:           []pnft.Fee{},
		Salt:           big.NewInt(1),
		ExtraParams:    []byte{},
	}

	// create new sell order
	sellOrder := order
	sellOrder.Trader = auth2.From
	sellInput, err := v2.SignSingle(&sellOrder, privKe2)
	require.Equal(v2.T(), nil, err)

	buyOrder := order
	buyOrder.Trader = v2.adapterAddr
	buyOrder.Side = 0
	buyInput := pnft.Input{
		Order:       buyOrder,
		BlockNumber: big.NewInt(0),
	}

	// match buy order
	adapterProxy, _ := abi.JSON(strings.NewReader(adapter.AdapterMetaData.ABI))
	pnftCallData, err := adapterProxy.Pack("execute", sellInput, buyInput, auth.From)
	require.Equal(v2.T(), nil, err)
	fmt.Println(hex.EncodeToString(pnftCallData))

	// buy opensea and pnft platform
	// prepare data
	// opensea buy order
	// fulfillAvailableAdvancedOrders
	openseaData, err := hex.DecodeString("87201b4100000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000000b800000000000000000000000000000000000000000000000000000000000000bc00000000000000000000000000000000000000000000000000000000000000ce00000007b02230091a7ed01230072f7006a004d60a8d4e71d599b8104250f000000000000000000000000000054b3dba467c9dbb916ef4d6aedafa19c4fef8258000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000056000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000046000000000000000000000000000000000000000000000000000000000000004e0000000000000000000000000bfb53a2c470cdb4ff32ee4f18a93b98f9f55d0e100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000160000000000000000000000000000000000000000000000000000000000000022000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000063ceb9e70000000000000000000000000000000000000000000000000000000063f798670000000000000000000000000000000000000000000000000000000000000000360c6ebe0000000000000000000000000000000000000000d8fe0a9ac1fb9b4c0000007b02230091a7ed01230072f7006a004d60a8d4e71d599b8104250f00000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000020000000000000000000000001f419b9469d641d333805c4054ca3b65af54d315000000000000000000000000000000000000000000000000000000000000065c000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000376c1e0a7f000000000000000000000000000000000000000000000000000000376c1e0a7f000000000000000000000000000bfb53a2c470cdb4ff32ee4f18a93b98f9f55d0e1000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000016bcc41e9000000000000000000000000000000000000000000000000000000016bcc41e90000000000000000000000000000000a26b00c1f0df003000390027140000faa71900000000000000000000000000000000000000000000000000000000000000418ea33cbdee6dbe18fdea621784a4df35124404fa88cb36a0318ec17751224d0f1269361993eb17a603341feee8724cc4e201af3a81e82083904e32b46dd3b3971c000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000046000000000000000000000000000000000000000000000000000000000000004e0000000000000000000000000bfb53a2c470cdb4ff32ee4f18a93b98f9f55d0e100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000160000000000000000000000000000000000000000000000000000000000000022000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000063ceba100000000000000000000000000000000000000000000000000000000063f798900000000000000000000000000000000000000000000000000000000000000000360c6ebe0000000000000000000000000000000000000000e7e66975d60048bd0000007b02230091a7ed01230072f7006a004d60a8d4e71d599b8104250f00000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000020000000000000000000000001f419b9469d641d333805c4054ca3b65af54d315000000000000000000000000000000000000000000000000000000000000065f000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000376c1e0a7f000000000000000000000000000000000000000000000000000000376c1e0a7f000000000000000000000000000bfb53a2c470cdb4ff32ee4f18a93b98f9f55d0e1000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000016bcc41e9000000000000000000000000000000000000000000000000000000016bcc41e90000000000000000000000000000000a26b00c1f0df003000390027140000faa7190000000000000000000000000000000000000000000000000000000000000041de0e0247e749e0fef3c3dee13f863451e9b70b9703b03262e3061043e325fe6a328c8633b00a2b66d09adaa8a7315365d15cf3dd5ebfbdaffe1223165fb606601c0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001ed549c1f360c6ebe")
	require.Equal(v2.T(), nil, err)

	auth.Value = big.NewInt(0).Add(big.NewInt(2000000000000000), order.Price)
	//tx, err := v2.adapter.BuyBatchETH(auth, []uint8{0, 1}, [][]byte{pnftCallData, openseaData})
	//require.Equal(v2.T(), nil, err)
	//fmt.Printf("buy nft tx: %s \n", tx.Hash().String())
	var salt = big.NewInt(0)
	salt.SetString("24446860302761739304752683030156737591518664810215442929816018635051325168460", 10)
	var salt2 = big.NewInt(0)
	salt2.SetString("24446860302761739304752683030156737591518664810215442929817092847857479731389", 10)
	advancedOrders := []opensea.AdvancedOrder{
		{
			Parameters: opensea.OrderParameters{
				Offerer: common.HexToAddress("0xbFb53a2c470cdb4FF32eE4F18A93B98F9f55D0E1"),
				Zone:    common.HexToAddress("0x0000000000000000000000000000000000000000"),
				Offer: []opensea.OfferItem{
					{
						ItemType:             2,
						Token:                common.HexToAddress("0x1f419B9469D641D333805C4054CA3b65Af54d315"),
						IdentifierOrCriteria: big.NewInt(1628),
						StartAmount:          big.NewInt(1),
						EndAmount:            big.NewInt(1),
					},
				},
				Consideration: []opensea.ConsiderationItem{
					{
						ItemType:             0,
						Token:                common.HexToAddress("0x0000000000000000000000000000000000000000"),
						IdentifierOrCriteria: big.NewInt(0),
						StartAmount:          big.NewInt(975000000000000),
						EndAmount:            big.NewInt(975000000000000),
						Recipient:            common.HexToAddress("0xbFb53a2c470cdb4FF32eE4F18A93B98F9f55D0E1"),
					},
					{
						ItemType:             1,
						Token:                common.HexToAddress("0x0000000000000000000000000000000000000000"),
						IdentifierOrCriteria: big.NewInt(0),
						StartAmount:          big.NewInt(25000000000000),
						EndAmount:            big.NewInt(25000000000000),
						Recipient:            common.HexToAddress("0x0000a26b00c1F0DF003000390027140000fAa719"),
					},
				},
				OrderType:                       0,
				StartTime:                       big.NewInt(1674492391),
				EndTime:                         big.NewInt(1677170791),
				ZoneHash:                        common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
				Salt:                            salt,
				ConduitKey:                      toByte32([]byte{0, 0, 0, 123, 2, 35, 0, 145, 167, 237, 1, 35, 0, 114, 247, 0, 106, 0, 77, 96, 168, 212, 231, 29, 89, 155, 129, 4, 37, 15, 0, 0}),
				TotalOriginalConsiderationItems: big.NewInt(2),
			},
			Numerator:   big.NewInt(1),
			Denominator: big.NewInt(1),
			Signature:   []byte{142, 163, 60, 189, 238, 109, 190, 24, 253, 234, 98, 23, 132, 164, 223, 53, 18, 68, 4, 250, 136, 203, 54, 160, 49, 142, 193, 119, 81, 34, 77, 15, 18, 105, 54, 25, 147, 235, 23, 166, 3, 52, 31, 238, 232, 114, 76, 196, 226, 1, 175, 58, 129, 232, 32, 131, 144, 78, 50, 180, 109, 211, 179, 151, 28},
			ExtraData:   []byte{},
		},
		{
			Parameters: opensea.OrderParameters{
				Offerer: common.HexToAddress("0xbFb53a2c470cdb4FF32eE4F18A93B98F9f55D0E1"),
				Zone:    common.HexToAddress("0x0000000000000000000000000000000000000000"),
				Offer: []opensea.OfferItem{
					{
						ItemType:             2,
						Token:                common.HexToAddress("0x1f419B9469D641D333805C4054CA3b65Af54d315"),
						IdentifierOrCriteria: big.NewInt(1631),
						StartAmount:          big.NewInt(1),
						EndAmount:            big.NewInt(1),
					},
				},
				Consideration: []opensea.ConsiderationItem{
					{
						ItemType:             0,
						Token:                common.HexToAddress("0x0000000000000000000000000000000000000000"),
						IdentifierOrCriteria: big.NewInt(0),
						StartAmount:          big.NewInt(975000000000000),
						EndAmount:            big.NewInt(975000000000000),
						Recipient:            common.HexToAddress("0xbFb53a2c470cdb4FF32eE4F18A93B98F9f55D0E1"),
					},
					{
						ItemType:             1,
						Token:                common.HexToAddress("0x0000000000000000000000000000000000000000"),
						IdentifierOrCriteria: big.NewInt(0),
						StartAmount:          big.NewInt(25000000000000),
						EndAmount:            big.NewInt(25000000000000),
						Recipient:            common.HexToAddress("0x0000a26b00c1F0DF003000390027140000fAa719"),
					},
				},
				OrderType:                       0,
				StartTime:                       big.NewInt(1674492432),
				EndTime:                         big.NewInt(1677170832),
				ZoneHash:                        common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
				Salt:                            salt2,
				ConduitKey:                      toByte32([]byte{0, 0, 0, 123, 2, 35, 0, 145, 167, 237, 1, 35, 0, 114, 247, 0, 106, 0, 77, 96, 168, 212, 231, 29, 89, 155, 129, 4, 37, 15, 0, 0}),
				TotalOriginalConsiderationItems: big.NewInt(2),
			},
			Numerator:   big.NewInt(1),
			Denominator: big.NewInt(1),
			Signature:   []byte{222, 14, 2, 71, 231, 73, 224, 254, 243, 195, 222, 225, 63, 134, 52, 81, 233, 183, 11, 151, 3, 176, 50, 98, 227, 6, 16, 67, 227, 37, 254, 106, 50, 140, 134, 51, 176, 10, 43, 102, 208, 154, 218, 168, 167, 49, 83, 101, 209, 92, 243, 221, 94, 191, 189, 175, 254, 18, 35, 22, 95, 182, 6, 96, 28},
			ExtraData:   []byte{},
		},
	}
	criteriaResolvers := []opensea.CriteriaResolver{}
	offerFulfillments := [][]opensea.FulfillmentComponent{
		{
			{
				OrderIndex: big.NewInt(0),
				ItemIndex:  big.NewInt(0),
			},
			{
				OrderIndex: big.NewInt(1),
				ItemIndex:  big.NewInt(0),
			},
		},
	}
	considerationFulfillments := [][]opensea.FulfillmentComponent{
		{
			{
				OrderIndex: big.NewInt(0),
				ItemIndex:  big.NewInt(0),
			},
			{
				OrderIndex: big.NewInt(1),
				ItemIndex:  big.NewInt(0),
			},
		},
		{
			{
				OrderIndex: big.NewInt(0),
				ItemIndex:  big.NewInt(1),
			},
			{
				OrderIndex: big.NewInt(1),
				ItemIndex:  big.NewInt(1),
			},
		},
	}
	fulfillerConduitKey := toByte32([]byte{0, 0, 0, 123, 2, 35, 0, 145, 167, 237, 1, 35, 0, 114, 247, 0, 106, 0, 77, 96, 168, 212, 231, 29, 89, 155, 129, 4, 37, 15, 0, 0})
	recipient := common.HexToAddress("0x54b3DBA467C9Dbb916EF4D6AedaFa19C4Fef8258")
	maximumFulfilled := big.NewInt(2)
	iopenseaAbi, _ := abi.JSON(strings.NewReader(opensea.IopenseaMetaData.ABI))
	openseaCalldataBuild, err := iopenseaAbi.Pack("fulfillAvailableAdvancedOrders", advancedOrders, criteriaResolvers, offerFulfillments, considerationFulfillments, fulfillerConduitKey, recipient, maximumFulfilled)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), openseaCalldataBuild, openseaData)

	pnftCallData, err = adapterProxy.Pack("buyBatchETH", []uint8{0, 1}, [][]byte{pnftCallData, openseaData})
	require.Equal(v2.T(), nil, err)
	fmt.Printf("final call data: %v \n", hex.EncodeToString(pnftCallData))
}

func (v2 *PNFTIntegrationTestSuite) SignSingle(order *pnft.Order, privKey *ecdsa.PrivateKey) (pnft.Input, error) {
	nonce, _ := v2.pnft.Nonces(nil, order.Trader)
	orderHash, err := v2.pnft.GetOrderHash(nil, *order, nonce)
	require.Equal(v2.T(), nil, err)

	domainSeparator, _ := v2.pnft.DOMAINSEPARATOR(nil)
	hashToSign := crypto.Keccak256Hash([]byte("\x19\x01"), domainSeparator[:], orderHash[:])
	signBytes, err := crypto.Sign(hashToSign[:], privKey)
	require.Equal(v2.T(), nil, err)

	return pnft.Input{
		Order:       *order,
		V:           signBytes[64] + 27,
		R:           toByte32(signBytes[:32]),
		S:           toByte32(signBytes[32:64]),
		BlockNumber: big.NewInt(0),
	}, nil
}
