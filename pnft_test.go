package bridge

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/incognitochain/bridge-eth/bridge/blur"
	"github.com/incognitochain/bridge-eth/bridge/erc721"
	"github.com/incognitochain/bridge-eth/bridge/opensea"
	"github.com/incognitochain/bridge-eth/bridge/pnft"
	"github.com/incognitochain/bridge-eth/bridge/pnft/executiondelegate"
	policymanager "github.com/incognitochain/bridge-eth/bridge/pnft/policyManager"
	"github.com/incognitochain/bridge-eth/bridge/pnft/policyManager/standardPolicyERC721"
	"github.com/incognitochain/bridge-eth/bridge/pnft/proxy"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"math/big"
	"os/exec"
	"strings"
	"testing"
)

// // Define the suite, and absorb the built-in basic suite
// // functionality from testify - including assertion methods.
type PNFTTestSuite struct {
	*TradingTestSuite

	p                 *Platform
	c                 *committees
	standardProxy     common.Address
	executionDelegate *executiondelegate.Executiondelegate
	policyManager     *policymanager.Policymanager
	pnft              *pnft.BlurExchange
	pnftAddr          common.Address
	execDelegateAddr  common.Address
	incNFT            *erc721.Erc721
	Forwarder         *opensea.Opensea
	BlurProxy         *blur.Blur
	BlurProxyAddr     common.Address
	auth              *bind.TransactOpts
	EtherAddress      common.Address
	ETHHost           string
	ETHClient         *ethclient.Client
	IncHost           string
	ETHPrivKey        *ecdsa.PrivateKey
}

func NewPNFTTestSuite(tradingTestSuite *TradingTestSuite) *PNFTTestSuite {
	return &PNFTTestSuite{
		TradingTestSuite: tradingTestSuite,
	}
}

func (v2 *PNFTTestSuite) DeployContracts() {
	// deploy execution delegate contract
	execDelegateAddr, _, delegateInst, err := executiondelegate.DeployExecutiondelegate(auth, v2.p.sim)
	require.Equal(v2.T(), nil, err)
	v2.executionDelegate = delegateInst
	v2.execDelegateAddr = execDelegateAddr

	// deploy policy manager contract
	policyMangerAddr, _, policyManagerInst, err := policymanager.DeployPolicymanager(auth, v2.p.sim)
	require.Equal(v2.T(), nil, err)
	v2.policyManager = policyManagerInst

	// deploy pnft implementation
	pnftImplementation, _, _, err := pnft.DeployBlurExchange(auth, v2.p.sim)
	require.Equal(v2.T(), nil, err)

	// deploy pnft proxy
	nftABI, _ := abi.JSON(strings.NewReader(pnft.BlurExchangeMetaData.ABI))
	input, err := nftABI.Pack("initialize", execDelegateAddr, policyMangerAddr, auth2.From, big.NewInt(30))
	require.Equal(v2.T(), nil, err)
	proxyNftm, _, _, err := proxy.DeployProxy(auth, v2.p.sim, pnftImplementation, input)
	require.Equal(v2.T(), nil, err)
	pnftInst, err := pnft.NewBlurExchange(proxyNftm, v2.p.sim)
	require.Equal(v2.T(), nil, err)
	v2.pnft = pnftInst
	v2.pnftAddr = proxyNftm

	// call approve to delegate contract
	_, err = v2.executionDelegate.ApproveContract(auth, proxyNftm)
	// deploy and add new policy to policy manager
	standardPolicyErc721, _, _, err := standardPolicyERC721.DeployStandardPolicyERC721(auth, v2.p.sim)
	require.Equal(v2.T(), nil, err)
	_, err = v2.policyManager.AddPolicy(auth, standardPolicyErc721)
	require.Equal(v2.T(), nil, err)
	v2.standardProxy = standardPolicyErc721

	// deploy new erc721 contract to test
	_, _, erc721Inst, err := erc721.DeployErc721(auth, v2.p.sim, "incognito nft", "INFT", "")
	require.Equal(v2.T(), nil, err)
	v2.incNFT = erc721Inst
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

	// standard policy
	err = exec.Command("/bin/bash", "-c",
		"solc @openzeppelin/=node_modules/@openzeppelin/ --base-path=$(pwd)/bridge --bin --abi --overwrite bridge/contracts/pnft/polices/standardPolicyERC721/StandardPolicyERC721.sol -o bridge/pnft/policyManager/standardPolicyERC721").Run()
	require.Equal(v2.T(), nil, err)
	err = exec.Command("/bin/bash", "-c",
		"abigen --abi bridge/pnft/policyManager/standardPolicyERC721/StandardPolicyERC721.abi --bin bridge/pnft/policyManager/standardPolicyERC721/StandardPolicyERC721.bin --out bridge/pnft/policyManager/standardPolicyERC721/standardPolicyERC721.go --pkg standardPolicyERC721").Run()
	require.Equal(v2.T(), nil, err)

	// mock erc721
	err = exec.Command("/bin/bash", "-c",
		"solc @openzeppelin/=node_modules/@openzeppelin/ --base-path=$(pwd)/bridge --bin --abi --overwrite bridge/node_modules/@openzeppelin/contracts/token/ERC721/presets/ERC721PresetMinterPauserAutoId.sol -o bridge/erc721").Run()
	require.Equal(v2.T(), nil, err)
	err = exec.Command("/bin/bash", "-c",
		"abigen --abi bridge/erc721/ERC721PresetMinterPauserAutoId.abi --bin bridge/erc721/ERC721PresetMinterPauserAutoId.bin --out bridge/erc721/erc721.go --pkg erc721").Run()
	require.Equal(v2.T(), nil, err)

	p, c, err := setupFixedCommittee()
	require.Equal(v2.T(), nil, err)
	v2.p = p
	v2.c = c
	v2.DeployContracts()
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
	fmt.Println("Buy Single NFT token...")

	// approve all for delegate contract
	_, err := v2.incNFT.Mint(auth, auth.From)
	require.Equal(v2.T(), nil, err)
	_, err = v2.incNFT.SetApprovalForAll(auth, v2.execDelegateAddr, true)
	require.Equal(v2.T(), nil, err)

	// create new sell order

	// match order

	fmt.Println("Buy Multis NFT tokens...")
}
