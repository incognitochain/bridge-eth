package bridge

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/incognitochain/bridge-eth/bridge/blur"
	"github.com/incognitochain/bridge-eth/bridge/opensea"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"os/exec"
	"testing"
)

// // Define the suite, and absorb the built-in basic suite
// // functionality from testify - including assertion methods.
type PNFTTestSuite struct {
	*TradingTestSuite

	Forwarder     *opensea.Opensea
	BlurProxy     *blur.Blur
	BlurProxyAddr common.Address
	Blur1         *blur.BlurInterface
	Blur2         *blur.BlurInterface
	auth          *bind.TransactOpts
	EtherAddress  common.Address
	ETHHost       string
	ETHClient     *ethclient.Client
	IncHost       string
	ETHPrivKey    *ecdsa.PrivateKey
}

func NewPNFTTestSuite(tradingTestSuite *TradingTestSuite) *PNFTTestSuite {
	return &PNFTTestSuite{
		TradingTestSuite: tradingTestSuite,
	}
}

func (v2 *PNFTTestSuite) DeployContracts() {
	
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (v2 *PNFTTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")
	err := exec.Command("/bin/bash", "-c",
		"solc @openzeppelin/=node_modules/@openzeppelin/ --base-path=$(pwd)/bridge --bin --abi --optimize --overwrite bridge/contracts/pnft/executionDelegate.sol -o bridge/pnft/executiondelegate").Run()
	require.Equal(v2.T(), nil, err)
	err = exec.Command("/bin/bash", "-c",
		"abigen --abi bridge/pnft/executiondelegate/ExecutionDelegate.abi --bin bridge/pnft/executiondelegate/ExecutionDelegate.bin --out bridge/pnft/executiondelegate/executiondelegate.go --pkg executiondelegate").Run()

	err = exec.Command("/bin/bash", "-c",
		"solc @openzeppelin/=node_modules/@openzeppelin/ --base-path=$(pwd)/bridge --combined-json abi,bin --optimize --overwrite bridge/contracts/pnft/pnft.sol -o bridge/pnft/").Run()
	require.Equal(v2.T(), nil, err)
	err = exec.Command("/bin/bash", "-c",
		"abigen  --combined-json bridge/pnft/combined.json --out bridge/pnft/pnft.go --pkg pnft  --alias _execute=execute2").Run()
	require.Equal(v2.T(), nil, err)

	err = exec.Command("/bin/bash", "-c",
		"solc @openzeppelin/=node_modules/@openzeppelin/ --base-path=$(pwd)/bridge --bin --abi --overwrite bridge/contracts/pnft/policyManager.sol -o bridge/pnft/policyManager/").Run()
	require.Equal(v2.T(), nil, err)
	err = exec.Command("/bin/bash", "-c",
		"abigen --abi bridge/pnft/policyManager/PolicyManager.abi --bin bridge/pnft/policyManager/PolicyManager.bin --out bridge/pnft/policyManager/policymanager.go --pkg policymanager").Run()

	// proxy
	err = exec.Command("/bin/bash", "-c",
		"solc @openzeppelin/=node_modules/@openzeppelin/ --base-path=$(pwd)/bridge --bin --abi --overwrite bridge/node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol -o bridge/pnft/proxy").Run()
	require.Equal(v2.T(), nil, err)
	err = exec.Command("/bin/bash", "-c",
		"abigen --abi bridge/pnft/proxy/ERC1967Proxy.abi --bin bridge/pnft/proxy/ERC1967Proxy.bin --out bridge/pnft/proxy/proxy.go --pkg proxy").Run()
	require.Equal(v2.T(), nil, err)
}

func (v2 *PNFTTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
}

func (v2 *PNFTTestSuite) SetupTest() {
	fmt.Println("Setting up the test...")
}

func (v2 *PNFTTestSuite) TearDownTest() {
	fmt.Println("Tearing down the test...")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestPNFT(t *testing.T) {
	fmt.Println("Starting entry point for pnft test suite...")
	tradingSuite := new(TradingTestSuite)
	suite.Run(t, tradingSuite)

	uniswapTradingSuite := NewPNFTTestSuite(tradingSuite)
	suite.Run(t, uniswapTradingSuite)

	fmt.Println("Finishing entry point for pnft test suite...")
}

func (v2 *PNFTTestSuite) TestPBlurCreateProp() {

}
