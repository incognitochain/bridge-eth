package main

import (
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/incognitochain/bridge-eth/bridge/kbntrade"
	"github.com/incognitochain/bridge-eth/bridge/vault"
	"github.com/stretchr/testify/suite"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"

	"github.com/stretchr/testify/require"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type VaultUpgradeTestSuite struct {
	*TradingTestSuite

	KyberTradeDeployedAddr common.Address

	KyberContractAddr    common.Address
	WETHAddr             common.Address
	EtherAddressStrKyber common.Address
	NewVaultAddress      common.Address
	IncProxyAddress      common.Address
	EtherAdminPrvKey     string

	IncOMGTokenIDStr  string
	IncPOLYTokenIDStr string

	OMGAddressStr  string
	POLYAddressStr string

	// token amounts for tests
	DepositingEther       float64
	OMGBalanceAfterStep1  *big.Int
	POLYBalanceAfterStep2 *big.Int
}

func NewVaultUpgradeTestSuite(tradingTestSuite *TradingTestSuite) *VaultUpgradeTestSuite {
	return &VaultUpgradeTestSuite{
		TradingTestSuite: tradingTestSuite,
	}
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (tradingSuite *VaultUpgradeTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")
	// Kovan env
	tradingSuite.EtherAddressStrKyber = common.HexToAddress("0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee")
	tradingSuite.EtherAdminPrvKey = "eecbeb089a9a2fd1768f373b0cbeae6dea8b8a30dc0e798df69a8e9648c8f262"
	tradingSuite.IncProxyAddress = common.HexToAddress("0x14171aDd648fB6f61C38D73055e2777ddC90603D") // rinkeby

	tradingSuite.IncOMGTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000082"
	tradingSuite.IncPOLYTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000081"
	tradingSuite.OMGAddressStr = "0x6FA355a7b6bD2D6bD8b927C489221BFBb6f1D7B2"                               // rinkeby
	tradingSuite.POLYAddressStr = "0x058832CA736AB027c12367e53915e34e87a6081B"                              // rinkeby
	tradingSuite.KyberTradeDeployedAddr = common.HexToAddress("0xd21C43eeAc6dE6a657EDAA5de0B79FfA4023897b") //rinkeby
	tradingSuite.DepositingEther = float64(0.1)
	tradingSuite.KyberContractAddr = common.HexToAddress("0xF77eC7Ed5f5B9a5aee4cfa6FFCaC6A4C315BaC76") // rinkeby
}

func (tradingSuite *VaultUpgradeTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
	tradingSuite.ETHClient.Close()
}

func (tradingSuite *VaultUpgradeTestSuite) SetupTest() {
	fmt.Println("Setting up the test...")
}

func (tradingSuite *VaultUpgradeTestSuite) TearDownTest() {
	fmt.Println("Tearing down the test...")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestVaultUpgradeTestSuite(t *testing.T) {
	fmt.Println("Starting entry point for Kyber test suite...")

	tradingSuite := new(TradingTestSuite)
	suite.Run(t, tradingSuite)

	kyberTradingSuite := NewVaultUpgradeTestSuite(tradingSuite)
	suite.Run(t, kyberTradingSuite)

	fmt.Println("Finishing entry point for 0x test suite...")
}

func (tradingSuite *VaultUpgradeTestSuite) executeWithKyber(
	srcQty *big.Int,
	srcTokenIDStr string,
	destTokenIDStr string,
) {
	tradeAbi, _ := abi.JSON(strings.NewReader(kbntrade.KbntradeABI))

	// Get contract instance
	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)
	auth := bind.NewKeyedTransactor(tradingSuite.ETHPrivKey)
	auth.GasPrice = big.NewInt(50000000000)
	auth.GasLimit = 2000000
	srcToken := common.HexToAddress(srcTokenIDStr)
	destToken := common.HexToAddress(destTokenIDStr)
	input, _ := tradeAbi.Pack("trade", srcToken, srcQty, destToken)
	timestamp := []byte(randomizeTimestamp())
	tempData := append(tradingSuite.KyberTradeDeployedAddr[:], input...)
	tempData1 := append(tempData, timestamp...)
	data := rawsha3(tempData1)
	signBytes, _ := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)

	tx, err := c.Execute(
		auth,
		srcToken,
		srcQty,
		destToken,
		tradingSuite.KyberTradeDeployedAddr,
		input,
		timestamp,
		signBytes,
	)
	require.Equal(tradingSuite.T(), nil, err)
	txHash := tx.Hash()
	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("Kyber trade executed , txHash: %x\n", txHash[:])
}

func (tradingSuite *VaultUpgradeTestSuite) unPause() {
	privKey, _ := crypto.HexToECDSA(tradingSuite.EtherAdminPrvKey)
	auth := bind.NewKeyedTransactor(privKey)
	auth.GasPrice = big.NewInt(50000000000)
	auth.GasLimit = 5000000
	// pause vault contract
	c, _ := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
	tx, _ := c.Unpause(auth)
	txHash := tx.Hash()
	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("unPause , txHash: %x\n", txHash[:])
	time.Sleep(15 * time.Second)
}

func (tradingSuite *VaultUpgradeTestSuite) moveAssetsToNewVault() {
	privKey, err := crypto.HexToECDSA(tradingSuite.EtherAdminPrvKey)
	auth := bind.NewKeyedTransactor(privKey)
	auth.GasPrice = big.NewInt(50000000000)
	auth.GasLimit = 5000000
	admin := common.HexToAddress(Admin)
	prevVault := tradingSuite.VaultAddr
	vaultAddr, tx, _, err := vault.DeployVault(auth, tradingSuite.ETHClient, admin, tradingSuite.IncProxyAddress, prevVault)
	tradingSuite.VaultAddr = vaultAddr
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("deployed new vault: ", tradingSuite.VaultAddr.Hex())
	fmt.Printf("addr: %s\n", vaultAddr.Hex())

	kbnTradeAddr, tx, _, err := kbntrade.DeployKbntrade(auth, tradingSuite.ETHClient, tradingSuite.KyberContractAddr, vaultAddr)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("deployed kbntrade")
	fmt.Printf("addr: %s\n", kbnTradeAddr.Hex())
	tradingSuite.KyberTradeDeployedAddr = kbnTradeAddr
	c, err := vault.NewVault(prevVault, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)

	// pause vault contract
	tx, err = c.Pause(auth)
	txHash := tx.Hash()
	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("Pause , txHash: %x\n", txHash[:])
	time.Sleep(15 * time.Second)

	// update new vault to old vault
	tx, err = c.Migrate(auth, vaultAddr)
	txHash = tx.Hash()
	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("Set newVault , txHash: %x\n", txHash[:])
	time.Sleep(15 * time.Second)

	cNew, err := vault.NewVault(vaultAddr, tradingSuite.ETHClient)
	fmt.Printf("Before Move assets ------")
	deposited, err := cNew.TotalDepositedToSCAmount(nil, common.HexToAddress(tradingSuite.OMGAddressStr))
	fmt.Println("OMG: ", deposited)
	deposited, err = cNew.TotalDepositedToSCAmount(nil, common.HexToAddress(tradingSuite.POLYAddressStr))
	fmt.Println("POLY: ", deposited)
	deposited, err = cNew.TotalDepositedToSCAmount(nil, common.HexToAddress(tradingSuite.EtherAddressStr))
	fmt.Println("ETH: ", deposited)
	assetAddresses := make([]common.Address, 0)
	assetAddresses = append(assetAddresses, common.HexToAddress(tradingSuite.OMGAddressStr))
	assetAddresses = append(assetAddresses, common.HexToAddress(tradingSuite.POLYAddressStr))
	assetAddresses = append(assetAddresses, common.HexToAddress(tradingSuite.EtherAddressStr))

	// move coin from old to new vault
	tx, err = c.MoveAssets(auth, assetAddresses)
	txHash = tx.Hash()
	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("Set move Assets, txHash: %x\n", txHash[:])
	time.Sleep(20 * time.Second)

	fmt.Printf("After Move assets ------")
	deposited, err = cNew.TotalDepositedToSCAmount(nil, common.HexToAddress(tradingSuite.OMGAddressStr))
	fmt.Println("OMG: ", deposited)
	deposited, err = cNew.TotalDepositedToSCAmount(nil, common.HexToAddress(tradingSuite.POLYAddressStr))
	fmt.Println("POLY: ", deposited)
	deposited, err = cNew.TotalDepositedToSCAmount(nil, common.HexToAddress(tradingSuite.EtherAddressStr))
	fmt.Println("ETH: ", deposited)
}

func (tradingSuite *VaultUpgradeTestSuite) Test1TradeEthForOMGWithKyber() {
	fmt.Println("============ TEST TRADE ETHER FOR OMG WITH Kyber AGGREGATOR ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	// tradingSuite.unPause()
	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
	burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("------------ STEP 1: porting ETH to pETH twice for upgrade purpose --------------")
	txHash := tradingSuite.depositETH(
		tradingSuite.DepositingEther,
		tradingSuite.IncPaymentAddrStr,
	)
	time.Sleep(10 * time.Second)

	txHash = tradingSuite.depositETH(
		tradingSuite.DepositingEther,
		tradingSuite.IncPaymentAddrStr,
	)
	time.Sleep(10 * time.Second)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(90 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)

	fmt.Println("------------ STEP 2: burning pETH to deposit ETH to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningPToken(
		tradingSuite.IncEtherTokenIDStr,
		burningPETH,
		pubKeyToAddrStr[2:],
		"createandsendburningfordeposittoscrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))
	deposited := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("deposited EHT: ", deposited)
	//require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPETH, big.NewInt(1000000000)), deposited)

	fmt.Println("------------ step 3: execute trade ETH for OMG through Kyber aggregator --------------")
	tradingSuite.executeWithKyber(
		tradeAmount,
		tradingSuite.EtherAddressStr,
		tradingSuite.OMGAddressStr,
	)
	time.Sleep(15 * time.Second)
	omgTraded := tradingSuite.getDepositedBalance(
		tradingSuite.OMGAddressStr,
		pubKeyToAddrStr,
	)
	tradingSuite.OMGBalanceAfterStep1 = omgTraded
	fmt.Println("omgTraded: ", omgTraded)
}

func (tradingSuite *VaultUpgradeTestSuite) Test2TradeOMGForPOLYWithKyber() {
	fmt.Println("============ TEST TRADE OMG FOR POLY WITH KYBER AGGREGATOR ===========")
	fmt.Println("------------ step 0: declaration & initialization --------------")
	depositingOMG := new(big.Int).Div(tradingSuite.OMGBalanceAfterStep1, big.NewInt(2))
	tradeAmount := depositingOMG
	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("------------ Execute trade OMG for POLY through Kyber aggregator --------------")
	tradingSuite.executeWithKyber(
		tradeAmount,
		tradingSuite.OMGAddressStr,
		tradingSuite.POLYAddressStr,
	)
	time.Sleep(15 * time.Second)
	polyTraded := tradingSuite.getDepositedBalance(
		tradingSuite.POLYAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("polyTraded: ", polyTraded)
	tradingSuite.POLYBalanceAfterStep2 = polyTraded
}

func (tradingSuite *VaultUpgradeTestSuite) Test3TradeOMGForPOLYWithKyberOnNewVault() {
	fmt.Println("============ TEST TRADE POLY FOR ETH WITH KYBER AGGREGATOR ===========")
	fmt.Println("--------- UPGRADE VAULT CONTRACT ------")
	depositingPOLY := tradingSuite.POLYBalanceAfterStep2
	tradeAmount := depositingPOLY
	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("pubKeyToAddrStr: ", pubKeyToAddrStr)
	// upgrade to new vault
	tradingSuite.moveAssetsToNewVault()
	fmt.Println("------------ Execute trade POLY for ETH through Kyber aggregator --------------")
	tradingSuite.executeWithKyber(
		tradeAmount,
		tradingSuite.POLYAddressStr,
		tradingSuite.EtherAddressStr,
	)
	etherTraded := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("etherTraded: ", etherTraded)
}
