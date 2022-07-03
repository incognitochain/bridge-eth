package bridge

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/suite"
)

type PUnifiedTestSuite struct {
	*TradingTestSuite
}

const (
	ETHNetworkID = 1
	BSCNetworkID = 2
	PLGNetworkID = 3
	FTMNetworkID = 4
)

type EVMProof struct {
	BlockHash common.Hash `json:"BlockHash"`
	TxIndex   uint        `json:"TxIndex"`
	Proof     []string    `json:"Proof"`
}

func NewPUnifiedTestSuite(tradingTestSuite *TradingTestSuite) *PUnifiedTestSuite {
	return &PUnifiedTestSuite{
		TradingTestSuite: tradingTestSuite,
	}
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (pUnifiedTestSuite *PUnifiedTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")
}

func (pUnifiedTestSuite *PUnifiedTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
	// pUnifiedTestSuite.BSCClient.Close()
}

func (pUnifiedTestSuite *PUnifiedTestSuite) SetupTest() {
	fmt.Println("Setting up the test...")
}

func (pUnifiedTestSuite *PUnifiedTestSuite) TearDownTest() {
	fmt.Println("Tearing down the test...")
}

func (pUnifiedTestSuite *PUnifiedTestSuite) BuildEVMShieldProof(networkID int, txHash common.Hash) ([]byte, error) {
	host := ""
	switch networkID {
	case ETHNetworkID:
		{
			host = pUnifiedTestSuite.ETHHost
		}
	case BSCNetworkID:
		{
			host = pUnifiedTestSuite.BSCHost
		}
	case PLGNetworkID:
		{
			host = pUnifiedTestSuite.PLGHost
		}
	case FTMNetworkID:
		{
			host = pUnifiedTestSuite.FTMHost
		}

	}

	timeStart := time.Now()
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(host, txHash)
	timeEnd := time.Since(timeStart)
	fmt.Printf("timeEnd: %v\n", timeEnd)
	fmt.Printf("ethBlockHash, ethTxIdx, ethDepositProof: %v, %v, %v\n", ethBlockHash, ethTxIdx, ethDepositProof)
	fmt.Printf("Proof: %#v\n", ethDepositProof)

	proof := EVMProof{
		BlockHash: common.HexToHash(ethBlockHash),
		TxIndex:   ethTxIdx,
		Proof:     ethDepositProof,
	}
	proofData, err := json.Marshal(proof)
	if err != nil {
		return nil, fmt.Errorf("Can not marshal proof: %v", err)
	}
	fmt.Printf("proofData: %v\n", proofData)

	return proofData, nil
}

func (pUnifiedTestSuite *PUnifiedTestSuite) TestBuildEVMProof() {
	txHash := common.HexToHash("0x9195a8cd8b55dcab39cccdded5c2ac51cf1893aef7b61828b8183bb7b27de2fc")
	networkID := BSCNetworkID
	pUnifiedTestSuite.BuildEVMShieldProof(networkID, txHash)

}

func TestPUnifiedTokenSuite(t *testing.T) {
	fmt.Println("Starting entry point for pUnified token test suite...")

	tradingSuite := new(TradingTestSuite)
	suite.Run(t, tradingSuite)

	pUnifiedSuite := NewPUnifiedTestSuite(tradingSuite)
	suite.Run(t, pUnifiedSuite)

	fmt.Println("Finishing entry point for pUnified token test suite...")
}
