package bridge

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/cbergoon/merkletree"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
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
	"golang.org/x/crypto/sha3"
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
	standardPolicy    common.Address
	pnftDeployer      *bind.TransactOpts
	executionDelegate *executiondelegate.Executiondelegate
	policyManager     *policymanager.Policymanager
	pnft              *pnft.BlurExchange
	pnftAddr          common.Address
	execDelegateAddr  common.Address
	incNFT            *erc721.Erc721
	incNFTAddr        common.Address
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

	privKey, err := crypto.HexToECDSA("1193a43543fc11e37daa1a026ae8fae69d84c5dd1f3f933047ff2588778c5cca")
	require.Equal(v2.T(), nil, err)
	v2.pnftDeployer, err = bind.NewKeyedTransactorWithChainID(privKey, big.NewInt(1337))
	require.Equal(v2.T(), nil, err)
	v2.SendEth(v2.pnftDeployer.From)

	proxyNftm, _, _, err := proxy.DeployProxy(v2.pnftDeployer, v2.p.sim, pnftImplementation, input)
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
	v2.standardPolicy = standardPolicyErc721

	// deploy new erc721 contract to test
	erc721Addr, _, erc721Inst, err := erc721.DeployErc721(auth, v2.p.sim, "incognito nft", "INFT", "")
	require.Equal(v2.T(), nil, err)
	v2.incNFT = erc721Inst
	v2.incNFTAddr = erc721Addr

	// deploy proxy contract
	_, _, fwd, err := opensea.DeployOpensea(auth, v2.p.sim)
	require.Equal(v2.T(), nil, err)
	v2.Forwarder = fwd

	// deploy blur proxy
	bPAddr, _, bProxy, err := blur.DeployBlur(auth, v2.p.sim)
	fmt.Println(bPAddr.String())
	require.Equal(v2.T(), nil, err)
	v2.BlurProxy = bProxy
	v2.BlurProxyAddr = bPAddr

	v2.p.sim.Commit()
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

	// incognito proxy pnft
	err = exec.Command("/bin/bash", "-c",
		"solc @openzeppelin/=node_modules/@openzeppelin/ --base-path=$(pwd)/bridge --bin --abi --optimize --overwrite bridge/contracts/blur/orderStructs.sol -o bridge/blur").Run()
	require.Equal(v2.T(), nil, err)
	err = exec.Command("/bin/bash", "-c",
		"abigen --abi bridge/blur/BuyBlurProxy.abi --bin bridge/blur/BuyBlurProxy.bin --out bridge/blur/blur.go --pkg blur").Run()

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
	_, err := v2.incNFT.Mint(auth, auth.From) // mint token nft id 0
	require.Equal(v2.T(), nil, err)
	_, err = v2.incNFT.SetApprovalForAll(auth, v2.execDelegateAddr, true)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()

	fmt.Println()

	order := pnft.Order{
		Side:           1,
		MatchingPolicy: v2.standardPolicy,
		Collection:     v2.incNFTAddr,
		PaymentToken:   common.Address{},
		TokenId:        big.NewInt(0),
		Amount:         big.NewInt(1),
		Price:          big.NewInt(100000000),
		ListingTime:    big.NewInt(int64(v2.p.sim.Blockchain().CurrentHeader().Time)),
		ExpirationTime: big.NewInt(int64(v2.p.sim.Blockchain().CurrentHeader().Time) + 100),
		Fees:           []pnft.Fee{},
		Salt:           big.NewInt(1),
		ExtraParams:    []byte{},
	}

	// create new sell order
	sellOrder := order
	sellOrder.Trader = auth.From
	sellInput, err := v2.SignSingle(&sellOrder, genesisAcc.PrivateKey)
	require.Equal(v2.T(), nil, err)

	buyOrder := order
	buyOrder.Trader = v2.BlurProxyAddr
	buyOrder.Side = 0
	buyInput := pnft.Input{
		Order:       buyOrder,
		BlockNumber: big.NewInt(0),
	}

	// match buy order
	auth2.Value = order.Price
	blurProxy, _ := abi.JSON(strings.NewReader(blur.BlurMetaData.ABI))
	buyInput.Order.Trader = v2.BlurProxyAddr
	fmt.Println(v2.pnftAddr)
	callDataBlur, err := blurProxy.Pack("execute", sellInput, buyInput, auth2.From)
	require.Equal(v2.T(), nil, err)
	fmt.Println(callDataBlur)
	_, err = v2.Forwarder.Forward(auth2, v2.BlurProxyAddr, callDataBlur)
	require.Equal(v2.T(), nil, err)
	auth2.Value = big.NewInt(0)
	v2.p.sim.Commit()
	// check ownership
	newOwner, err := v2.incNFT.OwnerOf(nil, big.NewInt(0))
	require.Equal(v2.T(), newOwner, auth2.From)

	fmt.Println("Sell and Buy Bulk NFT tokens...")

	// mint nft
	v2.incNFT.Mint(auth, auth.From) // mint token nft id 1
	v2.incNFT.Mint(auth, auth.From) // mint token nft id 2
	v2.incNFT.Mint(auth, auth.From) // mint token nft id 3
	v2.incNFT.SetApprovalForAll(auth, v2.execDelegateAddr, true)
	v2.p.sim.Commit()

	// order 1
	order1 := pnft.Order{
		Trader:         auth.From,
		Side:           1,
		MatchingPolicy: v2.standardPolicy,
		Collection:     v2.incNFTAddr,
		PaymentToken:   common.Address{},
		TokenId:        big.NewInt(1),
		Amount:         big.NewInt(1),
		Price:          big.NewInt(1000000000),
		ListingTime:    big.NewInt(int64(v2.p.sim.Blockchain().CurrentHeader().Time)),
		ExpirationTime: big.NewInt(int64(v2.p.sim.Blockchain().CurrentHeader().Time) + 100),
		Fees:           []pnft.Fee{},
		Salt:           big.NewInt(1),
		ExtraParams:    []byte{},
	}

	order2 := order1
	order2.TokenId = big.NewInt(2)
	order2.Price = big.NewInt(2000000000)

	order3 := order1
	order3.TokenId = big.NewInt(3)
	order3.Price = big.NewInt(3000000000)

	// create signature
	executions, err := v2.SignBulk([]pnft.Order{order1, order2, order3}, genesisAcc.PrivateKey)
	require.Equal(v2.T(), nil, err)

	// create buy side
	for i, _ := range executions {
		_buyOrder := executions[i].Sell.Order
		_buyOrder.Trader = auth2.From
		_buyOrder.Side = 0
		_buyInput := pnft.Input{
			Order:       _buyOrder,
			BlockNumber: big.NewInt(0),
		}
		executions[i].Buy = _buyInput
	}

	currentOwner, _ := v2.incNFT.OwnerOf(nil, big.NewInt(1))
	require.Equal(v2.T(), currentOwner, auth.From)

	// match order
	auth2.Value = big.NewInt(1e18)
	_, err = v2.pnft.BulkExecute(auth2, executions)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()

	newOwner, _ = v2.incNFT.OwnerOf(nil, big.NewInt(1))
	require.Equal(v2.T(), newOwner, auth2.From)
}

func (v2 *PNFTTestSuite) SignSingle(order *pnft.Order, privKey *ecdsa.PrivateKey) (pnft.Input, error) {
	orderHash, err := v2.pnft.GetOrderHash(nil, *order)
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

func (v2 *PNFTTestSuite) SignBulk(orders []pnft.Order, privKey *ecdsa.PrivateKey) ([]pnft.Execution, error) {
	//Build list of Content to build tree
	var list []merkletree.Content
	for _, order := range orders {
		orderHash, err := v2.pnft.GetOrderHash(nil, order)
		fmt.Println(hex.EncodeToString(orderHash[:]))
		require.Equal(v2.T(), nil, err)
		list = append(list, TestContent{x: orderHash})
	}

	//Create a new Merkle Tree from the list of Content
	t, err := merkletree.NewTreeWithHashStrategySorted(list, sha3.NewLegacyKeccak256, true)
	require.Equal(v2.T(), nil, err)
	// sign root hash
	byte32, err := abi.NewType("bytes32", "", nil)
	require.Equal(v2.T(), nil, err)
	argumentsRoot := abi.Arguments{
		abi.Argument{
			Type: byte32,
		},
		abi.Argument{
			Type: byte32,
		},
	}
	bytesRootPack, err := argumentsRoot.Pack(
		crypto.Keccak256Hash([]byte("Root(bytes32 root)")),
		toByte32(t.MerkleRoot()),
	)
	require.Equal(v2.T(), nil, err)
	domainSeparator, _ := v2.pnft.DOMAINSEPARATOR(nil)
	hashToSign := crypto.Keccak256Hash(
		[]byte("\x19\x01"),
		domainSeparator[:],
		crypto.Keccak256Hash(bytesRootPack).Bytes(),
	)
	signBytes, err := crypto.Sign(hashToSign[:], privKey)
	require.Equal(v2.T(), nil, err)

	// build result
	var result []pnft.Execution
	for i, _ := range list {
		var tempProof [][32]byte
		proof, _, err := t.GetMerklePath(list[i])
		require.Equal(v2.T(), nil, err)
		for _, p := range proof {
			tempProof = append(tempProof, toByte32(p))
		}
		arrayByte32, err := abi.NewType("bytes32[]", "", nil)
		require.Equal(v2.T(), nil, err)
		arguments := abi.Arguments{
			abi.Argument{
				Type: arrayByte32,
			},
		}
		bytes, err := arguments.Pack(tempProof)
		require.Equal(v2.T(), nil, err)
		var execution pnft.Execution
		execution.Sell = pnft.Input{
			Order:            orders[i],
			V:                signBytes[64] + 27,
			R:                toByte32(signBytes[:32]),
			S:                toByte32(signBytes[32:64]),
			ExtraSignature:   bytes,
			SignatureVersion: 1,
			BlockNumber:      big.NewInt(0),
		}
		result = append(result, execution)
	}

	return result, nil
}

// TestContent implements the Content interface provided by merkletree and represents the content stored in the tree.
type TestContent struct {
	x [32]byte
}

// CalculateHash hashes the values of a TestContent
func (t TestContent) CalculateHash() ([]byte, error) {
	return t.x[:], nil
}

// Equals tests for equality of two Contents
func (t TestContent) Equals(other merkletree.Content) (bool, error) {
	otherTC, ok := other.(TestContent)
	if !ok {
		return false, errors.New("value is not of type TestContent")
	}
	return t.x == otherTC.x, nil
}

func (v2 *PNFTTestSuite) SendEth(receiver common.Address) {
	var data []byte
	tx := types.NewTransaction(6, receiver, big.NewInt(5e18), 21000, big.NewInt(1e9), data)
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, genesisAcc.PrivateKey)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), nil, v2.p.sim.SendTransaction(context.Background(), signedTx))
	v2.p.sim.Commit()
}
