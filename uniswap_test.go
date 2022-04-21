package bridge

import (
	"context"
	"crypto/ecdsa"
	crand "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/incognitochain/bridge-eth/bridge/incognito_proxy"
	"github.com/incognitochain/bridge-eth/bridge/uniswap"
	"github.com/incognitochain/bridge-eth/bridge/vault"
	// "github.com/incognitochain/bridge-eth/bridge/vaultproxy"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// // Define the suite, and absorb the built-in basic suite
// // functionality from testify - including assertion methods.
type UniswapTestSuite struct {
	suite.Suite
	p             *Platform
	c             *committees
	v             *vault.Vault
	withdrawer    common.Address
	auth          *bind.TransactOpts
	EtherAddress  common.Address
	EthPrivateKey string
	EthHost       string
	ETHPrivKey    *ecdsa.PrivateKey
	ETHClient     *ethclient.Client
	UniswapProxy  common.Address
	VaultAddress  common.Address
	IncAddr       common.Address

	WETH                    common.Address
	DAIAddress              common.Address
	ETHUniswapRouterAddress common.Address
	MRKAddressStr           common.Address
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (v2 *UniswapTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")
	var err error
	v2.WETH = common.HexToAddress("0xd0a1e359811322d97991e03f863a0c30c2cf029c")
	v2.EtherAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")
	v2.ETHUniswapRouterAddress = common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D")
	v2.DAIAddress = common.HexToAddress("0x4f96fe3b7a6cf9725f59d353f723c1bdb64ca6aa")
	v2.MRKAddressStr = common.HexToAddress("0xef13c0c8abcaf5767160018d268f9697ae4f5375")
	v2.EthPrivateKey = "B8DB29A7A43FB88AD520F762C5FDF6F1B0155637FA1E5CB2C796AFE9E5C04E31"
	v2.VaultAddress = common.HexToAddress("0xF92c6bb5206BEFD3c8305ddd64d554322cC22aFc")
	v2.EthHost = "https://kovan.infura.io/v3/93fe721349134964aa71071a713c5cef"
	v2.UniswapProxy = common.HexToAddress("0xc7724A86ced6D95573d36fc4B98f9C0414B907a1")
	v2.IncAddr = common.HexToAddress("0xf295681641c170359E04Bbe2EA3985BaA4CF0baf")
	v2.connectToETH()
	v2.c = getFixedCommittee()
	v2.auth = bind.NewKeyedTransactor(v2.ETHPrivKey)
	v2.v, err = vault.NewVault(v2.VaultAddress, v2.ETHClient)
	require.Equal(v2.T(), nil, err)

	// uncomment to deploy new one on kovan
	// incAddr, tx, _, err := incognito_proxy.DeployIncognitoProxy(v2.auth, v2.ETHClient, v2.auth.From, v2.c.beacons, v2.c.bridges)
	// require.Equal(v2.T(), nil, err)
	// // Wait until tx is confirmed
	// err = wait(v2.ETHClient, tx.Hash())
	// require.Equal(v2.T(), nil, err)
	// v2.IncAddr = incAddr
	// fmt.Printf("Proxy address: %s\n", v2.IncAddr.Hex())

	// delegatorAddr, tx, _, err := vault.DeployVault(v2.auth, v2.ETHClient)
	// require.Equal(v2.T(), nil, err)
	// err = wait(v2.ETHClient, tx.Hash())
	// require.Equal(v2.T(), nil, err)
	// fmt.Printf("delegatorAddr address: %s\n", delegatorAddr.Hex())

	// vaultAbi, _ := abi.JSON(strings.NewReader(vault.VaultABI))
	// input, _ := vaultAbi.Pack("initialize", common.Address{})

	// v2.VaultAddress, tx, _, err = vaultproxy.DeployVaultproxy(v2.auth, v2.ETHClient, delegatorAddr, common.HexToAddress("0x0000000000000000000000000000000000000001"), v2.IncAddr, input)
	// require.Equal(v2.T(), nil, err)
	// err = wait(v2.ETHClient, tx.Hash())
	// require.Equal(v2.T(), nil, err)
	// fmt.Printf("vault proxy address: %s\n", v2.VaultAddress.Hex())

	// v2.UniswapProxy, tx, _, err = uniswap.DeployUniswapV2Trade(v2.auth, v2.ETHClient, v2.ETHUniswapRouterAddress)
	// require.Equal(v2.T(), nil, err)
	// err = wait(v2.ETHClient, tx.Hash())
	// require.Equal(v2.T(), nil, err)
	// fmt.Printf("Uniswap proxy address: %s\n", v2.UniswapProxy.Hex())
}

func (v2 *UniswapTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
}

func (v2 *UniswapTestSuite) SetupTest() {
	fmt.Println("Setting up the test...")
}

func (v2 *UniswapTestSuite) TearDownTest() {
	fmt.Println("Tearing down the test...")
}

func (v2 *UniswapTestSuite) connectToETH() {
	privKeyHex := v2.EthPrivateKey
	privKey, err := crypto.HexToECDSA(privKeyHex)
	require.Equal(v2.T(), nil, err)

	fmt.Printf("Sign Txs with address: %s\n", crypto.PubkeyToAddress(privKey.PublicKey).Hex())

	network := "development"
	fmt.Printf("Connecting to network %s\n", network)
	client, err := ethclient.Dial(v2.EthHost)
	require.Equal(v2.T(), nil, err)

	v2.ETHClient = client
	v2.ETHPrivKey = privKey
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestVaultV2Uniswap(t *testing.T) {
	fmt.Println("Starting entry point for vault v2 test suite...")
	suite.Run(t, new(UniswapTestSuite))

	fmt.Println("Finishing entry point for vault v2 test suite...")
}

func (v2 *UniswapTestSuite) TestUniswapTrade() {
	address := crypto.PubkeyToAddress(v2.ETHPrivKey.PublicKey)
	rand.Seed(time.Now().UnixNano())
	deposit := big.NewInt(int64(1e12))
	v2.auth.Value = deposit
	tx, err := v2.v.Deposit(v2.auth, "")
	require.Equal(v2.T(), nil, err)
	err = wait(v2.ETHClient, tx.Hash())
	require.Equal(v2.T(), nil, err)
	v2.auth.Value = big.NewInt(0)
	proof := buildWithdrawTestcaseV2Uniswap(v2.c, 243, 1, v2.EtherAddress, deposit, address)
	tx, err = SubmitBurnProof(v2.v, v2.auth, proof)
	require.Equal(v2.T(), nil, err)
	err = wait(v2.ETHClient, tx.Hash())
	require.Equal(v2.T(), nil, err)

	bal, err := v2.v.GetDepositedBalance(nil, v2.EtherAddress, address)
	require.Equal(v2.T(), nil, err)
	fmt.Printf("ETH depsited: %v", bal)

	// Trade ETH - ERC20
	expectedRate := v2.getExpectedRate(v2.WETH, v2.DAIAddress, bal)
	doubleExpectedRate := big.NewInt(0).Mul(big.NewInt(2), expectedRate)
	// trade with rate higher than market provided
	v2.executeWithUniswap(bal, v2.EtherAddress, v2.DAIAddress, doubleExpectedRate, true)
	// trade amount source token greater than available
	v2.executeWithUniswap(big.NewInt(0).Add(bal, big.NewInt(1)), v2.EtherAddress, v2.DAIAddress, expectedRate, true)
	// expect success on this case
	v2.executeWithUniswap(bal, v2.EtherAddress, v2.DAIAddress, expectedRate, false)
	bal, err = v2.v.GetDepositedBalance(nil, v2.DAIAddress, address)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), bal.Cmp(expectedRate) > -1, true)

	// Trade ERC20 - ERC20
	expectedRate = v2.getExpectedRate(v2.DAIAddress, v2.MRKAddressStr, bal)
	doubleExpectedRate = big.NewInt(0).Mul(big.NewInt(2), expectedRate)
	// trade with rate higher than market provided
	v2.executeWithUniswap(bal, v2.DAIAddress, v2.MRKAddressStr, doubleExpectedRate, true)
	// trade amount source token greater than available
	v2.executeWithUniswap(big.NewInt(0).Add(bal, big.NewInt(1)), v2.DAIAddress, v2.MRKAddressStr, expectedRate, true)
	// expect success on this case
	v2.executeWithUniswap(bal, v2.DAIAddress, v2.MRKAddressStr, expectedRate, false)
	bal, err = v2.v.GetDepositedBalance(nil, v2.MRKAddressStr, address)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), bal.Cmp(expectedRate) > -1, true)

	// Trade ERC20 - ETH
	expectedRate = v2.getExpectedRate(v2.MRKAddressStr, v2.WETH, bal)
	doubleExpectedRate = big.NewInt(0).Mul(big.NewInt(2), expectedRate)
	// trade with rate higher than market provided
	v2.executeWithUniswap(bal, v2.MRKAddressStr, v2.EtherAddress, doubleExpectedRate, true)
	// trade amount source token greater than available
	v2.executeWithUniswap(big.NewInt(0).Add(bal, big.NewInt(1)), v2.MRKAddressStr, v2.EtherAddress, expectedRate, true)
	// expect success on this case
	v2.executeWithUniswap(bal, v2.MRKAddressStr, v2.EtherAddress, expectedRate, false)
	bal, err = v2.v.GetDepositedBalance(nil, v2.EtherAddress, address)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), bal.Cmp(expectedRate) > -1, true)
}

func (v2 *UniswapTestSuite) TestUniswapProxyBadcases() {
	// anyone can call kyberproxy to trade directly
	deposit := big.NewInt(int64(8e9))
	v2.auth.Value = deposit
	expectedRate := v2.getExpectedRate(v2.EtherAddress, v2.DAIAddress, deposit)
	uniswap, err := uniswap.NewUniswapV2Trade(v2.UniswapProxy, v2.ETHClient)
	tx, err := uniswap.Trade(v2.auth, v2.EtherAddress, deposit, v2.DAIAddress, expectedRate)
	require.Equal(v2.T(), nil, err)
	err = wait(v2.ETHClient, tx.Hash())
	require.Equal(v2.T(), nil, err)

	// send eth to proxy directly
	amountSendToProxy := big.NewInt(int64(8e9))
	var data []byte
	gasLimit := uint64(40000)
	gasPrice, err := v2.ETHClient.SuggestGasPrice(context.Background())
	require.Equal(v2.T(), nil, err)
	nonce, err := v2.ETHClient.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(v2.ETHPrivKey.PublicKey))
	require.Equal(v2.T(), nil, err)
	tx = types.NewTransaction(nonce, v2.UniswapProxy, amountSendToProxy, gasLimit, gasPrice, data)
	chainID, err := v2.ETHClient.NetworkID(context.Background())
	require.Equal(v2.T(), nil, err)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), v2.ETHPrivKey)
	require.Equal(v2.T(), nil, err)
	err = v2.ETHClient.SendTransaction(context.Background(), signedTx)
	require.Equal(v2.T(), nil, err)
	err = wait(v2.ETHClient, signedTx.Hash())
	require.Equal(v2.T(), nil, err)
}

func (v2 *UniswapTestSuite) getExpectedRate(
	srcToken common.Address,
	destToken common.Address,
	srcQty *big.Int,
) *big.Int {
	if srcToken == v2.EtherAddress {
		srcToken = v2.WETH
	}
	if destToken == v2.EtherAddress {
		destToken = v2.WETH
	}
	c, err := uniswap.NewUniswapV2Trade(v2.UniswapProxy, v2.ETHClient)
	require.Equal(v2.T(), nil, err)
	amounts, err := c.GetAmountsOut(nil, srcToken, srcQty, destToken)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), 2, len(amounts))
	fmt.Printf("intput value: %d\n", amounts[0])
	fmt.Printf("output value: %d\n", amounts[1])

	return amounts[1]
}

func (v2 *UniswapTestSuite) executeWithUniswap(
	srcQty *big.Int,
	srcToken common.Address,
	destToken common.Address,
	expectRate *big.Int,
	isErrorExpected bool,
) {
	tradeAbi, _ := abi.JSON(strings.NewReader(uniswap.UniswapV2TradeABI))
	input, _ := tradeAbi.Pack("trade", srcToken, srcQty, destToken, expectRate)
	tx, err := runExecuteVault(v2.auth, v2.UniswapProxy, srcToken, srcQty, destToken, input, v2.v, []byte(randomizeTimestamp()), v2.ETHPrivKey)
	if !isErrorExpected {
		require.Equal(v2.T(), nil, err)
		err = wait(v2.ETHClient, tx.Hash())
		require.Equal(v2.T(), nil, err)
		fmt.Printf("Uniswap trade executed , txHash: %x\n", tx.Hash())
	} else {
		require.NotEqual(v2.T(), nil, err)
	}
}

func buildWithdrawTestcaseV2Uniswap(c *committees, meta, shard int, tokenID common.Address, amount *big.Int, withdrawer common.Address) *DecodedProof {
	inst, mp, blkData, blkHash := buildWithdrawDataV2Uniswap(meta, shard, tokenID, amount, withdrawer)
	ipBeacon := signAndReturnInstProof(c.beaconPrivs, true, mp, blkData, blkHash[:])
	return &DecodedProof{
		Instruction: inst,
		Heights:     [2]*big.Int{big.NewInt(1), big.NewInt(1)},

		InstPaths:       [2][][32]byte{ipBeacon.instPath},
		InstPathIsLefts: [2][]bool{ipBeacon.instPathIsLeft},
		InstRoots:       [2][32]byte{ipBeacon.instRoot},
		BlkData:         [2][32]byte{ipBeacon.blkData},
		SigIdxs:         [2][]*big.Int{ipBeacon.sigIdx},
		SigVs:           [2][]uint8{ipBeacon.sigV},
		SigRs:           [2][][32]byte{ipBeacon.sigR},
		SigSs:           [2][][32]byte{ipBeacon.sigS},
	}
}

func buildWithdrawDataV2Uniswap(meta, shard int, tokenID common.Address, amount *big.Int, withdrawer common.Address) ([]byte, *merklePath, []byte, []byte) {
	// Build instruction merkle tree
	numInst := 10
	startNodeID := 7
	height := big.NewInt(1)
	inst := buildDecodedWithdrawInstUniswap(meta, shard, tokenID, withdrawer, amount)
	instWithHeight := append(inst, toBytes32BigEndian(height.Bytes())...)
	data := randomMerkleHashes(numInst)
	data[startNodeID] = instWithHeight
	mp := buildInstructionMerklePath(data, numInst, startNodeID)

	// Generate random blkHash
	h := randomMerkleHashes(1)
	blkData := h[0]
	blkHash := rawsha3(append(blkData, mp.root[:]...))
	return inst, mp, blkData, blkHash[:]
}

func buildDecodedWithdrawInstUniswap(meta, shard int, tokenID, withdrawer common.Address, amount *big.Int) []byte {
	decoded := []byte{byte(meta)}
	decoded = append(decoded, byte(shard))
	decoded = append(decoded, toBytes32BigEndian(tokenID[:])...)
	decoded = append(decoded, toBytes32BigEndian(withdrawer[:])...)
	decoded = append(decoded, toBytes32BigEndian(amount.Bytes())...)
	txId := make([]byte, 32)
	crand.Read(txId)
	decoded = append(decoded, toBytes32BigEndian(txId)...) // txID
	decoded = append(decoded, make([]byte, 16)...)         // incTokenID, variable length
	return decoded
}
