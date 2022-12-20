package bridge

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/incognitochain/bridge-eth/bridge/governance"
	"github.com/incognitochain/bridge-eth/bridge/prvvote"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"math/big"
	"testing"
	"time"
)

// // Define the suite, and absorb the built-in basic suite
// // functionality from testify - including assertion methods.
type PDaoIntegrationTestSuite struct {
	*TradingTestSuite

	governance       *governance.Governance
	governanceHelper *governance.GovernanceHelper
	governanceAddr   common.Address
	prvvote          *prvvote.Prvvote
	prvvoteAddr      common.Address
	withdrawer       common.Address
	auth             *bind.TransactOpts
	EtherAddress     common.Address
	ETHHost          string
	ETHClient        *ethclient.Client
	IncHost          string
	ETHPrivKeyStr    string
	ETHPrivKey       *ecdsa.PrivateKey
}

func NewPDaoIntegrationTestSuite(tradingTestSuite *TradingTestSuite) *PDaoIntegrationTestSuite {
	return &PDaoIntegrationTestSuite{
		TradingTestSuite: tradingTestSuite,
	}
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (v2 *PDaoIntegrationTestSuite) SetupSuite() {
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

	gv, err := governance.NewGovernance(common.HexToAddress("0xD58d36a9053BbB69a75C4F9AF6864164dcbD2Cdb"), v2.ETHClient)
	require.Equal(v2.T(), nil, err)
	v2.governance = gv

	v2.PRVERC20Addr = common.HexToAddress("0xa0904E2F05D1108063Ac7CfB7719bD6518FDBDF4")
	prv, err := prvvote.NewPrvvote(v2.PRVERC20Addr, v2.ETHClient)
	require.Equal(v2.T(), nil, err)
	v2.prvvote = prv

	v2.ETHPrivKeyStr = "1193a43543fc11e37daa1a026ae8fae69d84c5dd1f3f933047ff2588778c5cca"
	privKey, err := crypto.HexToECDSA(v2.ETHPrivKeyStr)
	require.Equal(v2.T(), nil, err)
	v2.ETHPrivKey = privKey
}

func (v2 *PDaoIntegrationTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
}

func (v2 *PDaoIntegrationTestSuite) SetupTest() {
	fmt.Println("Setting up the test...")
}

func (v2 *PDaoIntegrationTestSuite) TearDownTest() {
	fmt.Println("Tearing down the test...")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestPDaoIntegration(t *testing.T) {
	fmt.Println("Starting entry point for pdao test suite...")
	tradingSuite := new(TradingTestSuite)
	suite.Run(t, tradingSuite)

	uniswapTradingSuite := NewPDaoIntegrationTestSuite(tradingSuite)
	suite.Run(t, uniswapTradingSuite)

	fmt.Println("Finishing entry point for pdao test suite...")
}

func (v2 *PDaoIntegrationTestSuite) TestPDAOCreateProp() {
	// burn prv token
	fmt.Println("------------ STEP 1: burning pPRV --------------")
	auth, err := bind.NewKeyedTransactorWithChainID(v2.ETHPrivKey, big.NewInt(int64(v2.ChainIDETH)))
	require.Equal(v2.T(), nil, err)

	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := v2.callBurningPRV(
		big.NewInt(1e9),
		v2.ETHOwnerAddrStr,
		"createandsendburningprvrequest",
	)
	require.Equal(v2.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(v2.T(), true, found)
	time.Sleep(60 * time.Second)

	// submit burn proof
	proof := v2.submitBurnProofForMintPRV(burningTxID.(string), v2.PRVERC20Addr, "getprvburnproof", v2.ETHClient, int64(v2.ChainIDETH))
	time.Sleep(60 * time.Second)

	// sign to reshield
	fmt.Println("Encode unshield tx: " + hex.EncodeToString(proof.Instruction[98:130]))
	signBytes, err := crypto.Sign(proof.Instruction[98:130], v2.ETHPrivKey)
	require.Equal(v2.T(), nil, err)
	tx, err := v2.prvvote.BurnBySignUnShieldTx(
		auth,
		toByte32(proof.Instruction[98:130]),
		signBytes[64]+27,
		toByte32(signBytes[:32]),
		toByte32(signBytes[32:64]),
	)
	require.Equal(v2.T(), nil, err)
	if err := wait(v2.ETHClient, tx.Hash()); err != nil {
		require.Equal(v2.T(), nil, err)
	}

	// shield prv to chain
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(v2.ETHHost, tx.Hash())
	require.Equal(v2.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)
	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(200 * time.Second)
	_, err = v2.callIssuingPRVReq(
		v2.IncPRVTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingprverc20req",
	)
	require.Equal(v2.T(), nil, err)
}
