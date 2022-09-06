package bridge

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ec "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/incognitochain/bridge-eth/bridge/incognito_proxy"
	"github.com/incognitochain/bridge-eth/bridge/kbntrade"
	"github.com/incognitochain/bridge-eth/bridge/vault"
	"github.com/incognitochain/bridge-eth/bridge/vaultproxy"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// // Define the suite, and absorb the built-in basic suite
// // functionality from testify - including assertion methods.
type KyberTestSuite struct {
	suite.Suite
	p                 *Platform
	c                 *committees
	v                 *vault.Vault
	withdrawer        common.Address
	auth              *bind.TransactOpts
	EtherAddress      common.Address
	EthPrivateKey     string
	EthHost           string
	ETHPrivKey        *ecdsa.PrivateKey
	ETHClient         *ethclient.Client
	KyberProxy        common.Address
	KyberMultiProxy   common.Address
	VaultAddress      common.Address
	IncAddr           common.Address
	KyberContractAddr common.Address

	KBNAddress      common.Address
	ETHKyberAddress common.Address
	MANAAddress     common.Address
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (v2 *KyberTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")
	v2.withdrawer = ec.HexToAddress("0xe722D8b71DCC0152D47D2438556a45D3357d631f")
	v2.EtherAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")
	v2.ETHKyberAddress = common.HexToAddress("0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee")
	v2.KBNAddress = common.HexToAddress("0x8c13AFB7815f10A8333955854E6ec7503eD841B7")
	v2.MANAAddress = common.HexToAddress("0xe19Ec968c15f487E96f631Ad9AA54fAE09A67C8c")
	v2.EthPrivateKey = "1ABA488300A9D7297A315D127837BE4219107C62C61966ECDF7A75431D75CC61"
	v2.KyberContractAddr = common.HexToAddress("0xd3add19ee7e5287148a5866784aE3C55bd4E375A")
	v2.EthHost = "http://localhost:8545"
	var err error
	fmt.Println("Pulling image if not exist, please wait...")
	// remove container if already running
	exec.Command("/bin/sh", "-c", "docker rm -f kybertrade").Output()
	_, err = exec.Command("/bin/sh", "-c", "docker run -d -p 8545:8545 --name kybertrade bomtb/kybertrade").Output()
	require.Equal(v2.T(), nil, err)
	time.Sleep(10 * time.Second)
}

func (v2 *KyberTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
	_, err := exec.Command("/bin/sh", "-c", "docker rm -f kybertrade").Output()
	require.Equal(v2.T(), nil, err)
}

func (v2 *KyberTestSuite) SetupTest() {
	fmt.Println("Setting up the test...")
	var err error
	ETHPrivKey, ETHClient, err := ethInstance(v2.EthPrivateKey, v2.EthHost)
	require.Equal(v2.T(), nil, err)
	v2.ETHClient = ETHClient
	v2.ETHPrivKey = ETHPrivKey
	v2.c = getFixedCommittee()
	v2.auth, err = bind.NewKeyedTransactorWithChainID(ETHPrivKey, v2.p.sim.Blockchain().Config().ChainID)
	require.Equal(v2.T(), nil, err)
	v2.IncAddr, _, _, err = incognito_proxy.DeployIncognitoProxy(v2.auth, ETHClient, v2.auth.From, v2.c.beacons, v2.c.bridges)
	require.Equal(v2.T(), nil, err)
	fmt.Printf("Proxy address: %s\n", v2.IncAddr.Hex())

	vaultDelegator, _, _, err := vault.DeployVault(v2.auth, v2.ETHClient)
	require.Equal(v2.T(), nil, err)
	fmt.Printf("Vault delegator address: %s\n", vaultDelegator.Hex())

	vaultAbi, _ := abi.JSON(strings.NewReader(vault.VaultABI))
	input, _ := vaultAbi.Pack("initialize", common.Address{})

	v2.VaultAddress, _, _, err = vaultproxy.DeployTransparentUpgradeableProxy(v2.auth, v2.ETHClient, vaultDelegator, common.HexToAddress("0x0000000000000000000000000000000000000001"), v2.IncAddr, input)
	require.Equal(v2.T(), nil, err)
	fmt.Printf("Vault proxy address: %s\n", v2.VaultAddress.Hex())
	v2.v, err = vault.NewVault(v2.VaultAddress, v2.ETHClient)
	require.Equal(v2.T(), nil, err)
	v2.KyberProxy, _, _, err = kbntrade.DeployKBNTrade(v2.auth, v2.ETHClient, v2.KyberContractAddr)
	require.Equal(v2.T(), nil, err)
	fmt.Printf("Kyber proxy address: %s\n", v2.KyberProxy.Hex())
}

func (v2 *KyberTestSuite) TearDownTest() {
	fmt.Println("Tearing down the test...")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestVaultV2Kyber(t *testing.T) {
	fmt.Println("Starting entry point for vault v2 test suite...")
	suite.Run(t, new(KyberTestSuite))

	fmt.Println("Finishing entry point for vault v2 test suite...")
}

func (v2 *KyberTestSuite) TestKyberTrade() {
	deposit := big.NewInt(int64(1e18))
	v2.auth.Value = deposit
	address := crypto.PubkeyToAddress(v2.ETHPrivKey.PublicKey)
	_, err := v2.v.Deposit(v2.auth, "")
	require.Equal(v2.T(), nil, err)
	v2.auth.Value = big.NewInt(0)
	proof := buildWithdrawTestcaseV2(v2.c, 243, 1, v2.EtherAddress, deposit, address)
	_, err = SubmitBurnProof(v2.v, v2.auth, proof)
	require.Equal(v2.T(), nil, err)

	bal, err := v2.v.GetDepositedBalance(nil, v2.EtherAddress, address)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), deposit, big.NewInt(int64(1e18)))

	// Trade ETH - ERC20
	bignnum, ok := new(big.Int).SetString("549000000000000000000", 10)
	require.Equal(v2.T(), ok, true)
	v2.executeWithKyber(bal, v2.EtherAddress, v2.KBNAddress)
	bal, err = v2.v.GetDepositedBalance(nil, v2.KBNAddress, address)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), bal, bignnum)

	// Trade ERC20 - ETH
	bignnum, ok = new(big.Int).SetString("3045173263944290564713", 10)
	require.Equal(v2.T(), ok, true)
	v2.executeWithKyber(bal, v2.KBNAddress, v2.MANAAddress)
	bal, err = v2.v.GetDepositedBalance(nil, v2.MANAAddress, address)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), bal, bignnum)

	// Trade ERC20 - ERC20
	bignnum, ok = new(big.Int).SetString("990451423918783372", 10)
	require.Equal(v2.T(), ok, true)
	v2.executeWithKyber(bal, v2.MANAAddress, v2.EtherAddress)
	bal, err = v2.v.GetDepositedBalance(nil, v2.EtherAddress, address)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), bal, bignnum)
}

func (v2 *KyberTestSuite) TestKyberProxyBadcases() {
	deposit := big.NewInt(int64(2e18))
	tradeamount := big.NewInt(int64(1e18))
	v2.auth.Value = deposit
	address := crypto.PubkeyToAddress(v2.ETHPrivKey.PublicKey)
	_, err := v2.v.Deposit(v2.auth, "")
	require.Equal(v2.T(), nil, err)
	v2.auth.Value = big.NewInt(0)
	proof := buildWithdrawTestcaseV2(v2.c, 243, 1, v2.EtherAddress, deposit, address)
	_, err = SubmitBurnProof(v2.v, v2.auth, proof)
	require.Equal(v2.T(), nil, err)

	bal, err := v2.v.GetDepositedBalance(nil, v2.EtherAddress, address)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), bal, deposit)

	srcToken := v2.EtherAddress
	destToken := v2.KBNAddress

	// Trade with the srcQty from vault less than srcQty in kyber proxy
	bal, err = v2.v.GetDepositedBalance(nil, v2.KBNAddress, address)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), 0, bal.Cmp(big.NewInt(0)))

	tradeAbi, _ := abi.JSON(strings.NewReader(kbntrade.KBNTradeABI))
	expectRate := v2.getExpectedRate(srcToken, destToken, deposit)
	input, _ := tradeAbi.Pack("trade", srcToken, deposit, destToken, expectRate)
	_, err = runExecuteVault(v2.auth, v2.KyberProxy, srcToken, tradeamount, destToken, input, v2.v, []byte(randomizeTimestamp()), v2.ETHPrivKey)
	require.NotEqual(v2.T(), nil, err)

	bal, err = v2.v.GetDepositedBalance(nil, v2.KBNAddress, address)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), 0, bal.Cmp(big.NewInt(0)))

	// Trade with minconversionRate larger than expecRate
	expectRate = v2.getExpectedRate(srcToken, destToken, tradeamount)
	input, _ = tradeAbi.Pack("trade", srcToken, tradeamount, destToken, big.NewInt(0).Add(expectRate, big.NewInt(int64(1))))
	_, err = runExecuteVault(v2.auth, v2.KyberProxy, srcToken, tradeamount, destToken, input, v2.v, []byte(randomizeTimestamp()), v2.ETHPrivKey)
	require.NotEqual(v2.T(), nil, err)
	bal, err = v2.v.GetDepositedBalance(nil, v2.KBNAddress, address)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), 0, bal.Cmp(big.NewInt(0)))

	// anyone can call kyberproxy to trade directly
	v2.auth.Value = deposit
	expectRate = v2.getExpectedRate(srcToken, destToken, tradeamount)
	kybertrade, err := kbntrade.NewKBNTrade(v2.KyberProxy, v2.ETHClient)
	_, err = kybertrade.Trade(v2.auth, srcToken, deposit, destToken, expectRate)
	require.Equal(v2.T(), nil, err)
}

func (v2 *KyberTestSuite) getExpectedRate(
	srcToken common.Address,
	destToken common.Address,
	srcQty *big.Int,
) *big.Int {
	if srcToken == v2.EtherAddress {
		srcToken = v2.ETHKyberAddress
	}
	if destToken == v2.EtherAddress {
		destToken = v2.ETHKyberAddress
	}
	c, err := kbntrade.NewKBNTrade(v2.KyberProxy, v2.ETHClient)
	require.Equal(v2.T(), nil, err)
	expectRate, _, err := c.GetConversionRates(nil, srcToken, srcQty, destToken)
	require.Equal(v2.T(), nil, err)
	return expectRate
}

func (v2 *KyberTestSuite) executeWithKyber(
	srcQty *big.Int,
	srcToken common.Address,
	destToken common.Address,
) {
	tradeAbi, _ := abi.JSON(strings.NewReader(kbntrade.KBNTradeABI))
	expectRate := v2.getExpectedRate(srcToken, destToken, srcQty)
	input, _ := tradeAbi.Pack("trade", srcToken, srcQty, destToken, expectRate)
	tx, err := runExecuteVault(v2.auth, v2.KyberProxy, srcToken, srcQty, destToken, input, v2.v, []byte(randomizeTimestamp()), v2.ETHPrivKey)
	require.Equal(v2.T(), nil, err)
	fmt.Printf("Kyber trade executed , txHash: %x\n", tx.Hash())
}

func ethInstance(ethPrivate string, ethEnpoint string) (*ecdsa.PrivateKey, *ethclient.Client, error) {
	privKey, err := crypto.HexToECDSA(ethPrivate)
	if err != nil {
		return nil, nil, err
	}
	fmt.Printf("Sign Txs with address: %s\n", crypto.PubkeyToAddress(privKey.PublicKey).Hex())

	network := "development"
	fmt.Printf("Connecting to network %s\n", network)
	client, err := ethclient.Dial(ethEnpoint)
	if err != nil {
		return nil, nil, err
	}
	return privKey, client, nil
}
