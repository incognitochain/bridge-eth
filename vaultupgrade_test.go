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
	kbnmultiTrade "github.com/incognitochain/bridge-eth/bridge/kbnmultitrade"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type VaultUpgradeTestSuite struct {
	*TradingTestSuite

	KyberTradeDeployedAddr common.Address

	KyberContractAddr    common.Address
	WETHAddr             common.Address
	EtherAddressStrKyber string

	NewVaultAddress      common.Address
	IncProxyAddress      common.Address
	EtherAdminPrvKey     string

	IncKBNTokenIDStr  string
	IncSALTTokenIDStr string
	IncOMGTokenIDStr  string
	IncSNTTokenIDStr  string
	IncPOLYTokenIDStr string
	IncZILTokenIDStr  string

	KBNAddressStr  string
	SALTAddressStr string
	OMGAddressStr  string
	SNTAddressStr  string
	POLYAddressStr string
	ZILAddressStr  string

	IncPrivKeyStrKb     string
	IncPaymentAddrStrKb string

	// token amounts for tests
	DepositingEther             float64
	OMGBalanceAfterStep1        *big.Int
	POLYBalanceAfterStep2       *big.Int
	KyberMultiTradeDeployedAddr common.Address
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
	tradingSuite.IncPrivKeyStrKb = "112t8ro4yu78UE82jpto12rp3Cd8Z2Mse7fcavSyXXP82oApE1cg9y8hWq69Z74fFHGJrQyHz54vU8Mv1kx5qZ54cRQJPkF5Cb7DhNqL5Tgt"
	tradingSuite.IncPaymentAddrStrKb = "12RyGbTyktYkXe7mNwmZeD4rktqxtHMe3Tsyf4XiZdKVGFssEHaF1ZUTpXZmpFACuDotVr7a6FEw8v6FTn8DEMqpHNxZ8fJW3KNN3i1"
	tradingSuite.EtherAdminPrvKey = "eecbeb089a9a2fd1768f373b0cbeae6dea8b8a30dc0e798df69a8e9648c8f262"
	tradingSuite.IncProxyAddress = common.HexToAddress("0xcC9B51964ab39555275cd51d6F866949e0dF6efc") // kovan

	tradingSuite.EtherAddressStrKyber = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
	tradingSuite.IncPOLYTokenIDStr = "d0379b8ccc25e4940d5b94ace07dcfa3656a20814279ddf2674f6d7180f65440"
	tradingSuite.IncOMGTokenIDStr = "27322fa7fce2c4d4d5a0022d595a0eec56d7735751a3ba8bc7f10b358ab938bc"
	tradingSuite.IncZILTokenIDStr = "3c115c066028bb682af410c594546b58026095ff149dc30c061749ee163d9051"
	tradingSuite.IncKBNTokenIDStr = "d6644f62d0ef0475335ae7bb6103f358979cbfcd2b85481e3bde2b82101a095c"
	tradingSuite.IncSALTTokenIDStr = "06ce44eae35daf57b9b8158ab95c0cddda9bac208fc380236a318ef40f3ac2ef"
	tradingSuite.IncSNTTokenIDStr = "414a6459526e827321cedb6e574d2ba2eb267c5735b0a65991602a405fb753b7"

	tradingSuite.ZILAddressStr = "0xAb74653cac23301066ABa8eba62b9Abd8a8c51d6"
	tradingSuite.POLYAddressStr = "0xd92266fd053161115163a7311067F0A4745070b5"
	tradingSuite.OMGAddressStr = "0xdB7ec4E4784118D9733710e46F7C83fE7889596a"
	tradingSuite.KBNAddressStr = "0xad67cB4d63C9da94AcA37fDF2761AaDF780ff4a2" // kovan
	tradingSuite.SALTAddressStr = "0x6fEE5727EE4CdCBD91f3A873ef2966dF31713A04"
	tradingSuite.SNTAddressStr = "0x4c99B04682fbF9020Fcb31677F8D8d66832d3322"

	tradingSuite.KyberTradeDeployedAddr = common.HexToAddress("0x243ad46Be9697Dc9dc45e701D03D6F920f60c09d")      //kovan
	tradingSuite.KyberMultiTradeDeployedAddr = common.HexToAddress("0xa5BCDBF240CC2310BD71D8cAd03Bc5892Ba5f4Ee") //kovan
	tradingSuite.DepositingEther = float64(0.001)
	tradingSuite.KyberContractAddr = common.HexToAddress("0x692f391bCc85cefCe8C237C01e1f636BbD70EA4D") // kovan
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

func (tradingSuite *VaultUpgradeTestSuite) getExpectedRate(
	srcToken string,
	destToken string,
	srcQty *big.Int,
) *big.Int {
	if srcToken == tradingSuite.EtherAddressStr {
		srcToken = tradingSuite.EtherAddressStrKyber
	}
	if destToken == tradingSuite.EtherAddressStr {
		destToken = tradingSuite.EtherAddressStrKyber
	}
	c, err := kbntrade.NewKbntrade(tradingSuite.KyberTradeDeployedAddr, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)
	expectRate, slippageRate, err := c.GetConversionRates(nil, common.HexToAddress(srcToken), srcQty, common.HexToAddress(destToken))
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Printf("slippageRate value: %d\n", slippageRate)
	fmt.Printf("expectRate value: %d\n", expectRate)
	return expectRate
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run

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
	auth.GasPrice = big.NewInt(50000000000)
	auth.GasLimit = 2000000
	expectRate := tradingSuite.getExpectedRate(srcTokenIDStr, destTokenIDStr, srcQty)
	input, _ := tradeAbi.Pack("trade", srcToken, srcQty, destToken, expectRate)
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

func (tradingSuite *VaultUpgradeTestSuite) executeMultiTradeWithKyber(
	srcQties []*big.Int,
	srcTokenIDStrs []string,
	destTokenIDStrs []string,
) {
	tradeAbi, _ := abi.JSON(strings.NewReader(kbnmultiTrade.KbnmultiTradeABI))
	auth := bind.NewKeyedTransactor(tradingSuite.ETHPrivKey)
	auth.GasPrice = big.NewInt(50000000000)
	auth.GasLimit = 2000000
	// Deploy kbnMultitrade
	// kbnMultiTradeAddr, tx, _, err := kbnmultiTrade.DeployKbnmultiTrade(auth, tradingSuite.ETHClient, tradingSuite.KyberContractAddr, tradingSuite.VaultAddr)
	// require.Equal(tradingSuite.T(), nil, err)
	// fmt.Println("deployed kbnMultitrade")
	// fmt.Printf("addr: %s\n", kbnMultiTradeAddr.Hex())
	// tradingSuite.KyberMultiTradeDeployedAddr = kbnMultiTradeAddr

	// Get contract instance
	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)
	auth.GasPrice = big.NewInt(50000000000)
	auth.GasLimit = 5000000
	sourceAddresses := make([]common.Address, 0)
	for _, p := range srcTokenIDStrs {
		sourceAddresses = append(sourceAddresses, common.HexToAddress(p))
	}
	destAddresses := make([]common.Address, 0)
	for _, p := range destTokenIDStrs {
		destAddresses = append(destAddresses, common.HexToAddress(p))
	}
	expectRates := make([]*big.Int, 0)
	for i := range destTokenIDStrs {
		expectRates = append(expectRates, tradingSuite.getExpectedRate(srcTokenIDStrs[i], destTokenIDStrs[i], srcQties[i]))
	}

	input, _ := tradeAbi.Pack("trade", sourceAddresses, srcQties, destAddresses, expectRates)
	timestamp := []byte(randomizeTimestamp())
	tempData := append(tradingSuite.KyberMultiTradeDeployedAddr[:], input...)
	tempData1 := append(tempData, timestamp...)
	data := rawsha3(tempData1)
	signBytes, _ := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)

	tx, err := c.ExecuteMulti(
		auth,
		sourceAddresses,
		srcQties,
		destAddresses,
		tradingSuite.KyberMultiTradeDeployedAddr,
		input,
		timestamp,
		signBytes,
	)
	require.Equal(tradingSuite.T(), nil, err)
	txHash := tx.Hash()
	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("Kyber multi trade executed , txHash: %x\n", txHash[:])
}

func (tradingSuite *VaultUpgradeTestSuite) unPause() {
	privKey, _ := crypto.HexToECDSA(tradingSuite.EtherAdminPrvKey)
	auth := bind.NewKeyedTransactor(privKey)
	auth.GasPrice = big.NewInt(50000000000)
	auth.GasLimit = 5000000
	// Unpause vault contract
	c, _ := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
	tx, _ := c.Unpause(auth)
	txHash := tx.Hash()
	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("unPause , txHash: %x\n", txHash[:])
	//time.Sleep(15 * time.Second)
}

func (tradingSuite *VaultUpgradeTestSuite) Pause() {
	privKey, _ := crypto.HexToECDSA(tradingSuite.EtherAdminPrvKey)
	auth := bind.NewKeyedTransactor(privKey)
	auth.GasPrice = big.NewInt(50000000000)
	auth.GasLimit = 5000000
	// pause vault contract
	c, _ := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
	tx, _ := c.Pause(auth)
	txHash := tx.Hash()
	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("Pause , txHash: %x\n", txHash[:])
	//time.Sleep(15 * time.Second)
}

func (tradingSuite *VaultUpgradeTestSuite) moveAssetsToNewVault() {
	privKey, err := crypto.HexToECDSA(tradingSuite.EtherAdminPrvKey)
	auth := bind.NewKeyedTransactor(privKey)
	auth.GasPrice = big.NewInt(50000000000)
	auth.GasLimit = 5000000
	admin := common.HexToAddress(Admin)
	prevVault := tradingSuite.VaultAddr
	vaultAddr, tx, _, err := vault.DeployVault(auth, tradingSuite.ETHClient, admin, tradingSuite.IncProxyAddress, prevVault)
	txHash := tx.Hash()
	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Println("deployed new vault: ", tradingSuite.VaultAddr.Hex())
	tradingSuite.VaultAddr = vaultAddr
	fmt.Printf("addr: %s\n", vaultAddr.Hex())
	kbnTradeAddr, tx, _, err := kbntrade.DeployKbntrade(auth, tradingSuite.ETHClient, tradingSuite.KyberContractAddr, vaultAddr)
	txHash = tx.Hash()
	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Println("deployed kbntrade")
	fmt.Printf("addr: %s\n", kbnTradeAddr.Hex())
	tradingSuite.KyberTradeDeployedAddr = kbnTradeAddr
	c, err := vault.NewVault(prevVault, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)
	//kbnTradeAddr, tx, _, err = kbntrade.DeployKbntrade(auth, tradingSuite.ETHClient, tradingSuite.KyberContractAddr, vaultAddr)
	// pause vault contract
	tx, err = c.Pause(auth)
	txHash = tx.Hash()
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
func (tradingSuite *VaultUpgradeTestSuite) Test1GetBalSmartContract() {
return
	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()

	// get info balance initialization
	balPrvInit, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance initialization : ", balPrvInit)
	fmt.Println("------------ ETH --------------")
	balEthInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance initialization : ", balEthInit)

	balEthScInit := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance initialization on SC : ", balEthScInit)

	balpEthInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance initialization : ", balpEthInit)

	fmt.Println("------------ OMG --------------")

	balOMGInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.OMGAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] OMG balance initialization : ", balOMGInit)

	balOMGScInit := tradingSuite.getDepositedBalance(
		tradingSuite.OMGAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] OMG balance initialization on SC : ", balOMGScInit)

	balpOMGInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncOMGTokenIDStr)
	fmt.Println("[INFO] pOMG balance initialization : ", balpOMGInit)


	fmt.Println("------------ ZIL --------------")

	balZILInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.ZILAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ZIL balance initialization : ", balZILInit)

	balZILScInit := tradingSuite.getDepositedBalance(
		tradingSuite.ZILAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ZIL balance initialization on SC : ", balZILScInit)

	balpZILInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncZILTokenIDStr)
	fmt.Println("[INFO] pZIL balance initialization : ", balpZILInit)

	fmt.Println("------------ KBN --------------")

	balKBNInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.KBNAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] KBN balance initialization : ", balKBNInit)

	balKBNScInit := tradingSuite.getDepositedBalance(
		tradingSuite.KBNAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] KBN balance initialization on SC : ", balKBNScInit)

	balpKBNInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncKBNTokenIDStr)
	fmt.Println("[INFO] pKBN balance initialization : ", balpKBNInit)

	fmt.Println("------------ SALT --------------")

	balSALTInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.SALTAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] SALT balance initialization : ", balSALTInit)

	balSALTScInit := tradingSuite.getDepositedBalance(
		tradingSuite.SALTAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] SALT balance initialization on SC : ", balSALTScInit)

	balpSALTInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncSALTTokenIDStr)
	fmt.Println("[INFO] pSALT balance initialization : ", balpSALTInit)
}

func (tradingSuite *VaultUpgradeTestSuite) Test2tradeOldSmartContract() {
return	
		fmt.Println("============ TEST 2 TRADE ON OLD SMART CONTRACT ===========")
		fmt.Println("------------ STEP 0: declaration & initialization --------------")
		depositETH := tradingSuite.DepositingEther * 6
		tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
		burningPETH := big.NewInt(0).Div(big.NewInt(int64(tradingSuite.DepositingEther * 6 * params.Ether)), big.NewInt(1000000000))
	
		pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
		fmt.Println(" publickey :", pubKeyToAddrStr)
	
		// get info balance initialization
		balPrvInit, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
		fmt.Println("[INFO] PRV balance initialization : ", balPrvInit)
		fmt.Println("------------ ETH --------------")
		balEthInit := tradingSuite.getBalanceOnETHNet(
			common.HexToAddress(tradingSuite.EtherAddressStr),
			common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		)
		fmt.Println("[INFO] ETH balance initialization : ", balEthInit)
	
		balEthScInit := tradingSuite.getDepositedBalance(
			tradingSuite.EtherAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] ETH balance initialization on SC : ", balEthScInit)
	
		balpEthInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
		fmt.Println("[INFO] pETH balance initialization : ", balpEthInit)
	
		fmt.Println("------------ OMG --------------")
	
		balOMGInit := tradingSuite.getBalanceOnETHNet(
			common.HexToAddress(tradingSuite.OMGAddressStr),
			common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		)
		fmt.Println("[INFO] OMG balance initialization : ", balOMGInit)
	
		balOMGScInit := tradingSuite.getDepositedBalance(
			tradingSuite.OMGAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] OMG balance initialization on SC : ", balOMGScInit)
	
		balpOMGInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncOMGTokenIDStr)
		fmt.Println("[INFO] pOMG balance initialization : ", balpOMGInit)
	
	
		fmt.Println("------------ ZIL --------------")
	
		balZILInit := tradingSuite.getBalanceOnETHNet(
			common.HexToAddress(tradingSuite.ZILAddressStr),
			common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		)
		fmt.Println("[INFO] ZIL balance initialization : ", balZILInit)
	
		balZILScInit := tradingSuite.getDepositedBalance(
			tradingSuite.ZILAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] ZIL balance initialization on SC : ", balZILScInit)
	
		balpZILInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncZILTokenIDStr)
		fmt.Println("[INFO] pZIL balance initialization : ", balpZILInit)
	
		fmt.Println("------------ KBN --------------")
	
		balKBNInit := tradingSuite.getBalanceOnETHNet(
			common.HexToAddress(tradingSuite.KBNAddressStr),
			common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		)
		fmt.Println("[INFO] KBN balance initialization : ", balKBNInit)
	
		balKBNScInit := tradingSuite.getDepositedBalance(
			tradingSuite.KBNAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] KBN balance initialization on SC : ", balKBNScInit)
	
		balpKBNInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncKBNTokenIDStr)
		fmt.Println("[INFO] pKBN balance initialization : ", balpKBNInit)
	
		fmt.Println("------------ SALT --------------")
	
		balSALTInit := tradingSuite.getBalanceOnETHNet(
			common.HexToAddress(tradingSuite.SALTAddressStr),
			common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		)
		fmt.Println("[INFO] SALT balance initialization : ", balSALTInit)
	
		balSALTScInit := tradingSuite.getDepositedBalance(
			tradingSuite.SALTAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] SALT balance initialization on SC : ", balSALTScInit)
	
		balpSALTInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncSALTTokenIDStr)
		fmt.Println("[INFO] pSALT balance initialization : ", balpSALTInit)
	
		//tradingSuite.unPause()

		fmt.Println("------------ STEP 1: porting ETH to pETH --------------")
		fmt.Println("[INFO] ETH deposit : ", burningPETH )
		txHash := tradingSuite.depositETH(
			depositETH,
			tradingSuite.IncPaymentAddrStr,
		)
		time.Sleep(15 * time.Second)
		_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
		require.Equal(tradingSuite.T(), nil, err)
		fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)
	
		fmt.Println("Waiting 90s for 15 blocks confirmation")
		time.Sleep(100 * time.Second)
		_, err = tradingSuite.callIssuingETHReq(
			tradingSuite.IncEtherTokenIDStr,
			ethDepositProof,
			ethBlockHash,
			ethTxIdx,
		)
		require.Equal(tradingSuite.T(), nil, err)
		time.Sleep(70 * time.Second)
		
		balpEthAfIssS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
		fmt.Println("[INFO] pETH balance after issuing step 1 : ", balpEthAfIssS1)
		// TODO assert pETH balance issuing
	
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
		time.Sleep(100 * time.Second)
	
		balpEthAfBurnS2, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
		fmt.Println("[INFO] pETH balance after burning step 2 : ", balpEthAfBurnS2)
		// TODO assert pETH balance issuing
	
		tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))
		balEthScDepS2 := tradingSuite.getDepositedBalance(
			tradingSuite.EtherAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] ETH balance after deposit on SC at step 2: ", balEthScDepS2)
		
		balOMGScS2 := tradingSuite.getDepositedBalance(
			tradingSuite.OMGAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] OMG balance on SC at step 2 : ", balOMGScS2)
	
		balZILScS2 := tradingSuite.getDepositedBalance(
			tradingSuite.ZILAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] ZIL balance on SC at step 2 : ", balZILScS2)
	
		balKBNScS2 := tradingSuite.getDepositedBalance(
			tradingSuite.KBNAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] KBN balance on SC at step 2 : ", balKBNScS2)
	
		balSALTScS2 := tradingSuite.getDepositedBalance(
			tradingSuite.SALTAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] SALT balance on SC at step 2 : ", balSALTScS2)
	
	
		//require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPETH, big.NewInt(1000000000)), deposited)
		fmt.Println("------------ step 3: execute trade ETH through Kyber aggregator --------------")

		fmt.Println("------------  execute trade ETH to OMG 3.1--------------")
		tradingSuite.executeWithKyber(
			tradeAmount,
			tradingSuite.EtherAddressStr,
			tradingSuite.OMGAddressStr,
		)
		fmt.Println("------------  execute trade ETH to ZIL 3.1--------------")
		tradingSuite.executeWithKyber(
			tradeAmount,
			tradingSuite.EtherAddressStr,
			tradingSuite.ZILAddressStr,
		)
		fmt.Println("------------  execute trade ETH to KBN 3.1 --------------")
		tradingSuite.executeWithKyber(
			tradeAmount,
			tradingSuite.EtherAddressStr,
			tradingSuite.KBNAddressStr,
		)
		fmt.Println("------------  execute trade ETH to SALT 3.1--------------")
		tradingSuite.executeWithKyber(
			tradeAmount,
			tradingSuite.EtherAddressStr,
			tradingSuite.SALTAddressStr,
		)
	
		balEthScTradeS3 := tradingSuite.getDepositedBalance(
			tradingSuite.EtherAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] ETH balance on SC after trade at step 3.1 : ", balEthScTradeS3)
	
		balOMGScTradeS3 := tradingSuite.getDepositedBalance(
			tradingSuite.OMGAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] OMG balance on SC after trade at step 3.1 : ", balOMGScTradeS3)
	
	
		balZILScS3 := tradingSuite.getDepositedBalance(
			tradingSuite.ZILAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] ZIL balance on SC after trade at step 3.1 : ", balZILScS3)
	
		balKBNScS3 := tradingSuite.getDepositedBalance(
			tradingSuite.KBNAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] KBN balance on SC after trade at step 3.1 : ", balKBNScS3)
	
		balSALTScS3 := tradingSuite.getDepositedBalance(
			tradingSuite.SALTAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] SALT balance on SC after trade at step 3.1 : ", balSALTScS3)
	
		fmt.Println("------------  execute trade OMG to SALT 3.2--------------")
		tradingSuite.executeWithKyber(
			balOMGScTradeS3,
			tradingSuite.OMGAddressStr,
			tradingSuite.SALTAddressStr,
		)

		balOMGScTradeS32 := tradingSuite.getDepositedBalance(
			tradingSuite.OMGAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] OMG balance on SC after trade at step 3.2 : ", balOMGScTradeS32)

		balSALTScS32 := tradingSuite.getDepositedBalance(
			tradingSuite.SALTAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] SALT balance on SC after trade at step 3.2 : ", balSALTScS32)

		balEthScTradeS32 := tradingSuite.getDepositedBalance(
			tradingSuite.EtherAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] ETH balance on SC after trade at step 3.2 : ", balEthScTradeS32)

		fmt.Println("------------  execute SALT OMG to Ether 3.3--------------")
		tradingSuite.executeWithKyber(
			balSALTScS32,
			tradingSuite.SALTAddressStr,
			tradingSuite.EtherAddressStr,
		)

		balSALTScS33 := tradingSuite.getDepositedBalance(
			tradingSuite.SALTAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] SALT balance on SC after trade at step 3.3 : ", balSALTScS33)

		balEthScTradeS33 := tradingSuite.getDepositedBalance(
			tradingSuite.EtherAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] ETH balance on SC after trade at step 3.3 : ", balEthScTradeS33)

		// fmt.Println("------------  execute multi trade Ether to OMG, SALT 3.4--------------")
		// tradingSuite.executeMultiTradeWithKyber(
		// 	[]*big.Int{tradeAmount, tradeAmount},
		// 	[]string{tradingSuite.EtherAddressStr, tradingSuite.EtherAddressStr},
		// 	[]string{tradingSuite.OMGAddressStr, tradingSuite.SALTAddressStr},
		// )

		balEthScTradeS34 := tradingSuite.getDepositedBalance(
			tradingSuite.EtherAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] ETH balance on SC after trade at step 3.4 : ", balEthScTradeS34)
	
		balOMGScTradeS34 := tradingSuite.getDepositedBalance(
			tradingSuite.OMGAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] OMG balance on SC after trade at step 3.4 : ", balOMGScTradeS34)
	
	
		balZILScS34 := tradingSuite.getDepositedBalance(
			tradingSuite.ZILAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] ZIL balance on SC after trade at step 3.4 : ", balZILScS34)
	
		balKBNScS34 := tradingSuite.getDepositedBalance(
			tradingSuite.KBNAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] KBN balance on SC after trade at step 3.4 : ", balKBNScS34)
	
		balSALTScS34 := tradingSuite.getDepositedBalance(
			tradingSuite.SALTAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] SALT balance on SC after trade at step 3.4 : ", balSALTScS34)

		fmt.Println("------------ step 4: withdrawing ETH from SC to pETH on Incognito --------------")
		txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
			tradingSuite.EtherAddressStr,
			tradeAmount,
		)
		time.Sleep(15 * time.Second)
		balEthScDepS4 := tradingSuite.getDepositedBalance(
			tradingSuite.EtherAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] ETH balance on SC after withdraw at step 4 : ", balEthScDepS4)
		// TODO assert ETH balane on SC

		_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq)
		require.Equal(tradingSuite.T(), nil, err)
		fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

		fmt.Println("Waiting 90s for 15 blocks confirmation")
		time.Sleep(100 * time.Second)

		_, err = tradingSuite.callIssuingETHReq(
			tradingSuite.IncEtherTokenIDStr,
			ethDepositProof,
			ethBlockHash,
			ethTxIdx,
		)
		require.Equal(tradingSuite.T(), nil, err)
		time.Sleep(140 * time.Second)

		balpEthAfIssS4, _ := tradingSuite.getBalanceTokenIncAccount(
			tradingSuite.IncPrivKeyStr,
			tradingSuite.IncEtherTokenIDStr,
		)
		fmt.Println("[INFO] pETH balance after issuing step 4 : ", balpEthAfIssS4)
		// TODO assert pETH balance issuing
		balPrvAfIssS4, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
		fmt.Println("[INFO] PRV balance after issuing step 4: ", balPrvAfIssS4)
		// TODO assert PRV balance remain

		fmt.Println("------------ step 5: withdrawing pETH from Incognito to ETH --------------")
		withdrawingPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))
		burningRes, err = tradingSuite.callBurningPToken(
			tradingSuite.IncEtherTokenIDStr,
			withdrawingPETH,
			tradingSuite.ETHOwnerAddrStr,
			"createandsendburningrequest",
		)
		require.Equal(tradingSuite.T(), nil, err)
		burningTxID, found = burningRes["TxID"]
		require.Equal(tradingSuite.T(), true, found)
		time.Sleep(140 * time.Second)

		tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))

		balETH := tradingSuite.getBalanceOnETHNet(
			common.HexToAddress(tradingSuite.EtherAddressStr),
			common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		)
		fmt.Println("Ether balance : ", balETH)
	
}


func (tradingSuite *VaultUpgradeTestSuite) Test3PauseSmartContract() {
return
		fmt.Println("============ TEST 3 PAUSE SMART CONTACT ===========")
		fmt.Println("------------ STEP 0: declaration & initialization --------------")
		tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
		burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))
	
		pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	
		// get info balance initialization
		balPrvInit, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
		fmt.Println("[INFO] PRV balance initialization : ", balPrvInit)
	
		balpEthInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
		fmt.Println("[INFO] pETH balance initialization : ", balpEthInit)
	
		balpOMGInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncOMGTokenIDStr)
		fmt.Println("[INFO] pOMG balance initialization : ", balpOMGInit)
	
		balEthInit := tradingSuite.getBalanceOnETHNet(
			common.HexToAddress(tradingSuite.EtherAddressStr),
			common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		)
		fmt.Println("[INFO] ETH balance initialization : ", balEthInit)
	
		balOMGInit := tradingSuite.getBalanceOnETHNet(
			common.HexToAddress(tradingSuite.OMGAddressStr),
			common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		)
		fmt.Println("[INFO] OMG balance initialization : ", balOMGInit)
	
		balEthScInit := tradingSuite.getDepositedBalance(
			tradingSuite.EtherAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] ETH balance initialization on SC : ", balEthScInit)
	
		balOMGScInit := tradingSuite.getDepositedBalance(
			tradingSuite.OMGAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] OMG balance initialization on SC : ", balOMGScInit)

		tradingSuite.Pause()
	
		fmt.Println("------------ PAUSED SMARTCONTRACT --------------")

		fmt.Println("------------ STEP 1: porting ETH to pETH --------------")
		txHash := tradingSuite.depositETH(
			tradingSuite.DepositingEther,
			tradingSuite.IncPaymentAddrStr,
		)
		time.Sleep(15 * time.Second)
		_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
		require.Equal(tradingSuite.T(), nil, err)
		fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)
	
		// get ETH remain after depsit
		balEthAfDep := tradingSuite.getBalanceOnETHNet(
			common.HexToAddress(tradingSuite.EtherAddressStr),
			common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		)
		fmt.Println("[INFO] ETH balance remain after deposit : ", balEthAfDep)
		// TODO : assert ETH balance
	
		fmt.Println("Waiting 90s for 15 blocks confirmation")
		time.Sleep(10 * time.Second)
		_, err = tradingSuite.callIssuingETHReq(
			tradingSuite.IncEtherTokenIDStr,
			ethDepositProof,
			ethBlockHash,
			ethTxIdx,
		)
		require.Equal(tradingSuite.T(), nil, err)
		time.Sleep(70 * time.Second)
		// check PRV and token balance after issuing
		balPrvAfIssS1, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
		fmt.Println("[INFO] PRV balance after issuing step 1: ", balPrvAfIssS1)
		// TODO assert PRV balance remain
		balpEthAfIssS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
		fmt.Println("[INFO] pETH balance after issuing step 1 : ", balpEthAfIssS1)
		// TODO assert pETH balance issuing
	
		fmt.Println("------------ STEP 2: burning pETH to deposit ZIL to SC --------------")
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
		time.Sleep(70 * time.Second)
	
		// check PRV and token balance after burning
		balPrvAfBurnS2, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
		fmt.Println("[INFO] PRV balance after burning step 2: ", balPrvAfBurnS2)
		// TODO assert PRV balance remain
		balpEthAfBurnS2, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
		fmt.Println("[INFO] pETH balance after burning step 2 : ", balpEthAfBurnS2)
		// TODO assert pETH balance issuing
	
		tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))
		balEthScDepS2 := tradingSuite.getDepositedBalance(
			tradingSuite.EtherAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] ETH balance after deposit on SC at step 2: ", balEthScDepS2)
		// TODO assert ETH balane on SC
		balOMGScS2 := tradingSuite.getDepositedBalance(
			tradingSuite.KBNAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] KBN balance on SC at step 2 : ", balOMGScS2)
	
		//require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPETH, big.NewInt(1000000000)), deposited)
		fmt.Println("------------ step 3: execute trade through Kyber aggregator --------------")
		fmt.Println("------------ step 3.1: execute trade ETH for KBN through Kyber aggregator --------------")
		tradingSuite.executeWithKyber(
			tradeAmount,
			tradingSuite.EtherAddressStr,
			tradingSuite.KBNAddressStr,
		)
		balEthScTradeS31 := tradingSuite.getDepositedBalance(
			tradingSuite.EtherAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] ETH balance on SC after trade at step 3.1 : ", balEthScTradeS31)
		// TODO assert ETH balane on SC
		balKBNScTradeS31 := tradingSuite.getDepositedBalance(
			tradingSuite.KBNAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] KBN balance on SC after trade at step 3.1 : ", balKBNScTradeS31)
		// TODO assert OMG balane on SC
	
		fmt.Println("------------ step 3.2: execute trade KBN for SALT through Kyber aggregator --------------")
		tradeAmount = balKBNScTradeS31
		tradingSuite.executeWithKyber(
			tradeAmount,
			tradingSuite.KBNAddressStr,
			tradingSuite.SALTAddressStr,
		)
		balKBNScTradeS32 := tradingSuite.getDepositedBalance(
			tradingSuite.KBNAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] OMG balance on SC after trade at step 3.2 : ", balKBNScTradeS32)
		// TODO assert OMG balane on SC
	
		balSALTScTradeS32 := tradingSuite.getDepositedBalance(
			tradingSuite.SALTAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] SALT balance on SC after trade at step 3.2 : ", balSALTScTradeS32)
		// TODO assert ETH balane on SC
	
		fmt.Println("------------ step 3.2: execute trade SALT for ETH through Kyber aggregator --------------")
		tradeAmount = balSALTScTradeS32
		tradingSuite.executeWithKyber(
			tradeAmount,
			tradingSuite.SALTAddressStr,
			tradingSuite.EtherAddressStr,
		)
	
		balSALTScTradeS33 := tradingSuite.getDepositedBalance(
			tradingSuite.SALTAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] SALT balance on SC after trade at step 3.3 : ", balSALTScTradeS33)
		// TODO assert ETH balane on SC
	
		balEthScTradeS33 := tradingSuite.getDepositedBalance(
			tradingSuite.EtherAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] ETH balance on SC after trade at step 3.3 : ", balEthScTradeS33)
		// TODO assert ETH balane on SC
	
		fmt.Println("------------ step 4: withdrawing ETH from SC to pETH on Incognito --------------")
		txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
			tradingSuite.EtherAddressStr,
			tradeAmount,
		)
		time.Sleep(15 * time.Second)
		balEthScDepS4 := tradingSuite.getDepositedBalance(
			tradingSuite.EtherAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] ETH balance on SC after withdraw at step 4 : ", balEthScDepS4)
		// TODO assert ETH balane on SC
	
		_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq)
		require.Equal(tradingSuite.T(), nil, err)
		fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)
	
		fmt.Println("Waiting 90s for 15 blocks confirmation")
		time.Sleep(10 * time.Second)
	
		_, err = tradingSuite.callIssuingETHReq(
			tradingSuite.IncEtherTokenIDStr,
			ethDepositProof,
			ethBlockHash,
			ethTxIdx,
		)
		require.Equal(tradingSuite.T(), nil, err)
		time.Sleep(140 * time.Second)
	
		balpEthAfIssS4, _ := tradingSuite.getBalanceTokenIncAccount(
			tradingSuite.IncPrivKeyStr,
			tradingSuite.IncEtherTokenIDStr,
		)
		fmt.Println("[INFO] pETH balance after issuing step 4 : ", balpEthAfIssS4)
		// TODO assert pETH balance issuing
		balPrvAfIssS4, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
		fmt.Println("[INFO] PRV balance after issuing step 4: ", balPrvAfIssS4)
		// TODO assert PRV balance remain
	
		fmt.Println("------------ step 5: withdrawing pETH from Incognito to ETH --------------")
		withdrawingPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))
		burningRes, err = tradingSuite.callBurningPToken(
			tradingSuite.IncEtherTokenIDStr,
			withdrawingPETH,
			tradingSuite.ETHOwnerAddrStr,
			"createandsendburningrequest",
		)
		require.Equal(tradingSuite.T(), nil, err)
		burningTxID, found = burningRes["TxID"]
		require.Equal(tradingSuite.T(), true, found)
		time.Sleep(140 * time.Second)
	
		tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))
	
		balETH := tradingSuite.getBalanceOnETHNet(
			common.HexToAddress(tradingSuite.EtherAddressStr),
			common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		)
		fmt.Println("Ether balance after trade: ", balETH)
		// require.Equal(tradingSuite.T(), withdrawingPETH.Uint64(), bal.Div(bal, big.NewInt(1000000000)).Uint64())
	
}


func (tradingSuite *VaultUpgradeTestSuite) Test4UnPauseSmartContract() {
return
			fmt.Println("============ TEST 4 unPAUSE SMART CONTACT ===========")
			fmt.Println("------------ STEP 0: declaration & initialization --------------")
			tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
			burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))
		
			pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
		
			// get info balance initialization
			balPrvInit, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
			fmt.Println("[INFO] PRV balance initialization : ", balPrvInit)
		
			balpEthInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
			fmt.Println("[INFO] pETH balance initialization : ", balpEthInit)
		
			balpOMGInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncOMGTokenIDStr)
			fmt.Println("[INFO] pOMG balance initialization : ", balpOMGInit)
		
			balEthInit := tradingSuite.getBalanceOnETHNet(
				common.HexToAddress(tradingSuite.EtherAddressStr),
				common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
			)
			fmt.Println("[INFO] ETH balance initialization : ", balEthInit)
		
			balOMGInit := tradingSuite.getBalanceOnETHNet(
				common.HexToAddress(tradingSuite.OMGAddressStr),
				common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
			)
			fmt.Println("[INFO] OMG balance initialization : ", balOMGInit)
		
			balEthScInit := tradingSuite.getDepositedBalance(
				tradingSuite.EtherAddressStr,
				pubKeyToAddrStr,
			)
			fmt.Println("[INFO] ETH balance initialization on SC : ", balEthScInit)
		
			balOMGScInit := tradingSuite.getDepositedBalance(
				tradingSuite.OMGAddressStr,
				pubKeyToAddrStr,
			)
			fmt.Println("[INFO] OMG balance initialization on SC : ", balOMGScInit)
	
			tradingSuite.Pause()
		
			fmt.Println("------------ PAUSED SMARTCONTRACT --------------")
			tradingSuite.unPause()
	
			fmt.Println("------------ UNPAUSED SMARTCONTRACT --------------")
	
			fmt.Println("------------ STEP 1: porting ETH to pETH --------------")
			txHash := tradingSuite.depositETH(
				tradingSuite.DepositingEther,
				tradingSuite.IncPaymentAddrStr,
			)
			time.Sleep(15 * time.Second)
			_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
			require.Equal(tradingSuite.T(), nil, err)
			fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)
		
			// get ETH remain after depsit
			balEthAfDep := tradingSuite.getBalanceOnETHNet(
				common.HexToAddress(tradingSuite.EtherAddressStr),
				common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
			)
			fmt.Println("[INFO] ETH balance remain after deposit : ", balEthAfDep)
			// TODO : assert ETH balance
		
			fmt.Println("Waiting 90s for 15 blocks confirmation")
			time.Sleep(10 * time.Second)
			_, err = tradingSuite.callIssuingETHReq(
				tradingSuite.IncEtherTokenIDStr,
				ethDepositProof,
				ethBlockHash,
				ethTxIdx,
			)
			require.Equal(tradingSuite.T(), nil, err)
			time.Sleep(70 * time.Second)
			// check PRV and token balance after issuing
			balPrvAfIssS1, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
			fmt.Println("[INFO] PRV balance after issuing step 1: ", balPrvAfIssS1)
			// TODO assert PRV balance remain
			balpEthAfIssS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
			fmt.Println("[INFO] pETH balance after issuing step 1 : ", balpEthAfIssS1)
			// TODO assert pETH balance issuing
		
			fmt.Println("------------ STEP 2: burning pETH to deposit ZIL to SC --------------")
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
			time.Sleep(70 * time.Second)
		
			// check PRV and token balance after burning
			balPrvAfBurnS2, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
			fmt.Println("[INFO] PRV balance after burning step 2: ", balPrvAfBurnS2)
			// TODO assert PRV balance remain
			balpEthAfBurnS2, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
			fmt.Println("[INFO] pETH balance after burning step 2 : ", balpEthAfBurnS2)
			// TODO assert pETH balance issuing
		
			tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))
			balEthScDepS2 := tradingSuite.getDepositedBalance(
				tradingSuite.EtherAddressStr,
				pubKeyToAddrStr,
			)
			fmt.Println("[INFO] ETH balance after deposit on SC at step 2: ", balEthScDepS2)
			// TODO assert ETH balane on SC
			balOMGScS2 := tradingSuite.getDepositedBalance(
				tradingSuite.KBNAddressStr,
				pubKeyToAddrStr,
			)
			fmt.Println("[INFO] KBN balance on SC at step 2 : ", balOMGScS2)
		
			//require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPETH, big.NewInt(1000000000)), deposited)
			fmt.Println("------------ step 3: execute trade through Kyber aggregator --------------")
			fmt.Println("------------ step 3.1: execute trade ETH for KBN through Kyber aggregator --------------")
			tradingSuite.executeWithKyber(
				tradeAmount,
				tradingSuite.EtherAddressStr,
				tradingSuite.KBNAddressStr,
			)
			balEthScTradeS31 := tradingSuite.getDepositedBalance(
				tradingSuite.EtherAddressStr,
				pubKeyToAddrStr,
			)
			fmt.Println("[INFO] ETH balance on SC after trade at step 3.1 : ", balEthScTradeS31)
			// TODO assert ETH balane on SC
			balKBNScTradeS31 := tradingSuite.getDepositedBalance(
				tradingSuite.KBNAddressStr,
				pubKeyToAddrStr,
			)
			fmt.Println("[INFO] KBN balance on SC after trade at step 3.1 : ", balKBNScTradeS31)
			// TODO assert OMG balane on SC
		
			fmt.Println("------------ step 3.2: execute trade KBN for SALT through Kyber aggregator --------------")
			tradeAmount = balKBNScTradeS31
			tradingSuite.executeWithKyber(
				tradeAmount,
				tradingSuite.KBNAddressStr,
				tradingSuite.SALTAddressStr,
			)
			balKBNScTradeS32 := tradingSuite.getDepositedBalance(
				tradingSuite.KBNAddressStr,
				pubKeyToAddrStr,
			)
			fmt.Println("[INFO] OMG balance on SC after trade at step 3.2 : ", balKBNScTradeS32)
			// TODO assert OMG balane on SC
		
			balSALTScTradeS32 := tradingSuite.getDepositedBalance(
				tradingSuite.SALTAddressStr,
				pubKeyToAddrStr,
			)
			fmt.Println("[INFO] SALT balance on SC after trade at step 3.2 : ", balSALTScTradeS32)
			// TODO assert ETH balane on SC
		
			fmt.Println("------------ step 3.2: execute trade SALT for ETH through Kyber aggregator --------------")
			tradeAmount = balSALTScTradeS32
			tradingSuite.executeWithKyber(
				tradeAmount,
				tradingSuite.SALTAddressStr,
				tradingSuite.EtherAddressStr,
			)
		
			balSALTScTradeS33 := tradingSuite.getDepositedBalance(
				tradingSuite.SALTAddressStr,
				pubKeyToAddrStr,
			)
			fmt.Println("[INFO] SALT balance on SC after trade at step 3.3 : ", balSALTScTradeS33)
			// TODO assert ETH balane on SC
		
			balEthScTradeS33 := tradingSuite.getDepositedBalance(
				tradingSuite.EtherAddressStr,
				pubKeyToAddrStr,
			)
			fmt.Println("[INFO] ETH balance on SC after trade at step 3.3 : ", balEthScTradeS33)
			// TODO assert ETH balane on SC
		
			fmt.Println("------------ step 4: withdrawing ETH from SC to pETH on Incognito --------------")
			txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
				tradingSuite.EtherAddressStr,
				tradeAmount,
			)
			time.Sleep(15 * time.Second)
			balEthScDepS4 := tradingSuite.getDepositedBalance(
				tradingSuite.EtherAddressStr,
				pubKeyToAddrStr,
			)
			fmt.Println("[INFO] ETH balance on SC after withdraw at step 4 : ", balEthScDepS4)
			// TODO assert ETH balane on SC
		
			_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq)
			require.Equal(tradingSuite.T(), nil, err)
			fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)
		
			fmt.Println("Waiting 90s for 15 blocks confirmation")
			time.Sleep(10 * time.Second)
		
			_, err = tradingSuite.callIssuingETHReq(
				tradingSuite.IncEtherTokenIDStr,
				ethDepositProof,
				ethBlockHash,
				ethTxIdx,
			)
			require.Equal(tradingSuite.T(), nil, err)
			time.Sleep(140 * time.Second)
		
			balpEthAfIssS4, _ := tradingSuite.getBalanceTokenIncAccount(
				tradingSuite.IncPrivKeyStr,
				tradingSuite.IncEtherTokenIDStr,
			)
			fmt.Println("[INFO] pETH balance after issuing step 4 : ", balpEthAfIssS4)
			// TODO assert pETH balance issuing
			balPrvAfIssS4, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
			fmt.Println("[INFO] PRV balance after issuing step 4: ", balPrvAfIssS4)
			// TODO assert PRV balance remain
		
			fmt.Println("------------ step 5: withdrawing pETH from Incognito to ETH --------------")
			withdrawingPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))
			burningRes, err = tradingSuite.callBurningPToken(
				tradingSuite.IncEtherTokenIDStr,
				withdrawingPETH,
				tradingSuite.ETHOwnerAddrStr,
				"createandsendburningrequest",
			)
			require.Equal(tradingSuite.T(), nil, err)
			burningTxID, found = burningRes["TxID"]
			require.Equal(tradingSuite.T(), true, found)
			time.Sleep(140 * time.Second)
		
			tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))
		
			balETH := tradingSuite.getBalanceOnETHNet(
				common.HexToAddress(tradingSuite.EtherAddressStr),
				common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
			)
			fmt.Println("Ether balance after trade: ", balETH)
			// require.Equal(tradingSuite.T(), withdrawingPETH.Uint64(), bal.Div(bal, big.NewInt(1000000000)).Uint64())
}
	
func (tradingSuite *VaultUpgradeTestSuite) Test5TradeOnNewVault() {
	return	
		fmt.Println("============ TEST TRADE ON NEW VAULT ===========")
		pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	
		// get info balance initialization
		balPrvInit, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
		fmt.Println("[INFO] PRV balance initialization : ", balPrvInit)
		fmt.Println("------------ ETH --------------")
		balEthInit := tradingSuite.getBalanceOnETHNet(
			common.HexToAddress(tradingSuite.EtherAddressStr),
			common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		)
		fmt.Println("[INFO] ETH balance initialization : ", balEthInit)
	
		balEthScInit := tradingSuite.getDepositedBalance(
			tradingSuite.EtherAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] ETH balance initialization on SC : ", balEthScInit)
	
		balpEthInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
		fmt.Println("[INFO] pETH balance initialization : ", balpEthInit)
	
		fmt.Println("------------ OMG --------------")
	
		balOMGInit := tradingSuite.getBalanceOnETHNet(
			common.HexToAddress(tradingSuite.OMGAddressStr),
			common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		)
		fmt.Println("[INFO] OMG balance initialization : ", balOMGInit)
	
		balOMGScInit := tradingSuite.getDepositedBalance(
			tradingSuite.OMGAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] OMG balance initialization on SC : ", balOMGScInit)
	
		balpOMGInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncOMGTokenIDStr)
		fmt.Println("[INFO] pOMG balance initialization : ", balpOMGInit)
	
	
		fmt.Println("------------ ZIL --------------")
	
		balZILInit := tradingSuite.getBalanceOnETHNet(
			common.HexToAddress(tradingSuite.ZILAddressStr),
			common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		)
		fmt.Println("[INFO] ZIL balance initialization : ", balZILInit)
	
		balZILScInit := tradingSuite.getDepositedBalance(
			tradingSuite.ZILAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] ZIL balance initialization on SC : ", balZILScInit)
	
		balpZILInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncZILTokenIDStr)
		fmt.Println("[INFO] pZIL balance initialization : ", balpZILInit)
	
		fmt.Println("------------ KBN --------------")
	
		balKBNInit := tradingSuite.getBalanceOnETHNet(
			common.HexToAddress(tradingSuite.KBNAddressStr),
			common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		)
		fmt.Println("[INFO] KBN balance initialization : ", balKBNInit)
	
		balKBNScInit := tradingSuite.getDepositedBalance(
			tradingSuite.KBNAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] KBN balance initialization on SC : ", balKBNScInit)
	
		balpKBNInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncKBNTokenIDStr)
		fmt.Println("[INFO] pKBN balance initialization : ", balpKBNInit)
	
		fmt.Println("------------ SALT --------------")
	
		balSALTInit := tradingSuite.getBalanceOnETHNet(
			common.HexToAddress(tradingSuite.SALTAddressStr),
			common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		)
		fmt.Println("[INFO] SALT balance initialization : ", balSALTInit)
	
		balSALTScInit := tradingSuite.getDepositedBalance(
			tradingSuite.SALTAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] SALT balance initialization on SC : ", balSALTScInit)
	
	
		balpSALTInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncSALTTokenIDStr)
		fmt.Println("[INFO] pSALT balance initialization : ", balpSALTInit)
	
		tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
		//burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))
	
		fmt.Println("------------ Update new vault --------------")
		tradingSuite.unPause()
		tradingSuite.moveAssetsToNewVault()
	
		fmt.Println("------------ Execute trade POLY for ETH through Kyber aggregator --------------")
		tradingSuite.executeWithKyber(
			tradeAmount,
			tradingSuite.OMGAddressStr,
			tradingSuite.EtherAddressStr,
		)
	
		fmt.Println("------------AFTER TRADE--------------")
		// get info balance after trade
		balPrvInit, _ = tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
		fmt.Println("[INFO] PRV balance  : ", balPrvInit)
		fmt.Println("------------ ETH --------------")
		balEthInit = tradingSuite.getBalanceOnETHNet(
			common.HexToAddress(tradingSuite.EtherAddressStr),
			common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		)
		fmt.Println("[INFO] ETH balance  : ", balEthInit)
	
		balEthScInit = tradingSuite.getDepositedBalance(
			tradingSuite.EtherAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] ETH balance  on SC : ", balEthScInit)
	
		balpEthInit, _ = tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
		fmt.Println("[INFO] pETH balance  : ", balpEthInit)
	
		fmt.Println("------------ OMG --------------")
	
		balOMGInit = tradingSuite.getBalanceOnETHNet(
			common.HexToAddress(tradingSuite.OMGAddressStr),
			common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		)
		fmt.Println("[INFO] OMG balance  : ", balOMGInit)
	
		balOMGScInit = tradingSuite.getDepositedBalance(
			tradingSuite.OMGAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] OMG balance  on SC : ", balOMGScInit)
	
		balpOMGInit, _ = tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncOMGTokenIDStr)
		fmt.Println("[INFO] pOMG balance  : ", balpOMGInit)
	
	
		fmt.Println("------------ ZIL --------------")
	
		balZILInit = tradingSuite.getBalanceOnETHNet(
			common.HexToAddress(tradingSuite.ZILAddressStr),
			common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		)
		fmt.Println("[INFO] ZIL balance  : ", balZILInit)
	
		balZILScInit = tradingSuite.getDepositedBalance(
			tradingSuite.ZILAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] ZIL balance  on SC : ", balZILScInit)
	
		balpZILInit, _ = tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncZILTokenIDStr)
		fmt.Println("[INFO] pZIL balance  : ", balpZILInit)
	
		fmt.Println("------------ KBN --------------")
	
		balKBNInit = tradingSuite.getBalanceOnETHNet(
			common.HexToAddress(tradingSuite.KBNAddressStr),
			common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		)
		fmt.Println("[INFO] KBN balance  : ", balKBNInit)
	
		balKBNScInit = tradingSuite.getDepositedBalance(
			tradingSuite.KBNAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] KBN balance  on SC : ", balKBNScInit)
	
		balpKBNInit, _ = tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncKBNTokenIDStr)
		fmt.Println("[INFO] pKBN balance  : ", balpKBNInit)
	
		fmt.Println("------------ SALT --------------")
	
		balSALTInit = tradingSuite.getBalanceOnETHNet(
			common.HexToAddress(tradingSuite.SALTAddressStr),
			common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		)
		fmt.Println("[INFO] SALT balance  : ", balSALTInit)
	
		balSALTScInit = tradingSuite.getDepositedBalance(
			tradingSuite.SALTAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] SALT balance  on SC : ", balSALTScInit)
	
	
		balpSALTInit, _ = tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncSALTTokenIDStr)
		fmt.Println("[INFO] pSALT balance  : ", balpSALTInit)
	}

func (tradingSuite *VaultUpgradeTestSuite) Test12DepositAndWithdrwaEther() {
return
		fmt.Println("============ TEST 12 DEPOSIT AND WITHDRAW ETHER ===========")
		fmt.Println("------------ STEP 0: declaration & initialization --------------")
		tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
		burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))
	
		pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	
		// get info balance initialization
		balPrvInit, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
		fmt.Println("[INFO] PRV balance initialization : ", balPrvInit)
	
		balpEthInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
		fmt.Println("[INFO] pETH balance initialization : ", balpEthInit)
	
		balEthInit := tradingSuite.getBalanceOnETHNet(
			common.HexToAddress(tradingSuite.EtherAddressStr),
			common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		)
		fmt.Println("[INFO] ETH balance initialization : ", balEthInit)
	
		balEthScInit := tradingSuite.getDepositedBalance(
			tradingSuite.EtherAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] ETH balance initialization on SC : ", balEthScInit)
	
		fmt.Println("------------ STEP 1: porting ETH to pETH --------------")
	
		fmt.Println("amount ETH deposit : ", (big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))))
	
		// Deposit to proof
		txHash := tradingSuite.depositETH(
			tradingSuite.DepositingEther,
			tradingSuite.IncPaymentAddrStr,
		)
	
		fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHash))
		// get ETH remain after depsit
		balEthAfDep := tradingSuite.getBalanceOnETHNet(
			common.HexToAddress(tradingSuite.EtherAddressStr),
			common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		)
		fmt.Println("[INFO] ETH balance remain after deposit : ", balEthAfDep)
	
		//	require.Equal(tradingSuite.T(), balEthAfDep, big.NewInt(0).Sub(big.NewInt(0).Sub(balEthInit, big.NewInt(int64(tradingSuite.DepositingEther*params.Ether))), tradingSuite.getGasFeeETHbyTxhash(txHash)), "balance ETH incorrect")
	
		// Proof
		_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
		require.Equal(tradingSuite.T(), nil, err)
		fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)
	
		fmt.Println("Waiting 90s for 15 blocks confirmation")
		time.Sleep(100 * time.Second)
		_, err = tradingSuite.callIssuingETHReq(
			tradingSuite.IncEtherTokenIDStr,
			ethDepositProof,
			ethBlockHash,
			ethTxIdx,
		)
		require.Equal(tradingSuite.T(), nil, err)
	
		time.Sleep(120 * time.Second)
		// check PRV and token balance after issuing
		balPrvAfIssS1, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
		fmt.Println("[INFO] PRV balance after issuing step 1: ", balPrvAfIssS1)
		//require.NotEqual(tradingSuite.T(), balPrvAfIssS1, (balPrvInit - tradingSuite.getFeePRVbyTxhashInC(issuuRes["TxID"].(string))), "Balance PRV remain incorrect after issuu step 1")
	
		balpEthAfIssS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
		fmt.Println("[INFO] pETH balance after issuing step 1 : ", balpEthAfIssS1)
		//require.Equal(tradingSuite.T(), big.NewInt(int64(balpEthAfIssS1-balpEthInit)), big.NewInt(0).Div(big.NewInt(int64(tradingSuite.DepositingEther*params.Ether)), big.NewInt(1000000000)), " balnce pToken issuu incorrect")
	
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
	
		// check PRV and token balance after burning
		balPrvAfBurnS2, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
		fmt.Println("[INFO] PRV balance after burning step 2: ", balPrvAfBurnS2)
		//require.NotEqual(tradingSuite.T(), balPrvAfBurnS2, (balPrvAfIssS1 - tradingSuite.getFeePRVbyTxhashInC(burningRes["TxID"].(string))), "Balance PRV remain incorrect after burn step 2")
	
		balpEthAfBurnS2, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
		fmt.Println("[INFO] pETH balance after burning step 2 : ", balpEthAfBurnS2)
		// TODO assert pETH balance issuing
	
		txHash2 := tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))
		fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHash2))
		// get ETH remain
		balEthAfDep2 := tradingSuite.getBalanceOnETHNet(
			common.HexToAddress(tradingSuite.EtherAddressStr),
			common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		)
		fmt.Println("[INFO] ETH balance remain  : ", balEthAfDep2)
		// TODO : assert ETH balance
		require.Equal(tradingSuite.T(), balEthAfDep2, big.NewInt(0).Sub(balEthAfDep, tradingSuite.getGasFeeETHbyTxhash(txHash2)), "balance ETH incorrect")
	
		balEthScDepS2 := tradingSuite.getDepositedBalance(
			tradingSuite.EtherAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] ETH balance after deposit on SC at step 2: ", balEthScDepS2)
		// TODO assert ETH balane on SC
	
		fmt.Println("------------ STEP 3: withdraw ETH to deposit pETH to Incognito  --------------")
	
		txHashByEmittingWithdrawalReq1 := tradingSuite.requestWithdraw(
			tradingSuite.EtherAddressStr,
			balEthScDepS2,
		)
	
		balDaiScDepS41 := tradingSuite.getDepositedBalance(
			tradingSuite.EtherAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] ETH balance on SC after withdraw at step 3 : ", balDaiScDepS41)
		// TODO assert DAI balane on SC
	
		_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq1)
		require.Equal(tradingSuite.T(), nil, err)
		fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)
	
		fmt.Println("Waiting 90s for 15 blocks confirmation")
		time.Sleep(100 * time.Second)
	
		_, err = tradingSuite.callIssuingETHReq(
			tradingSuite.IncEtherTokenIDStr,
			ethDepositProof,
			ethBlockHash,
			ethTxIdx,
		)
		require.Equal(tradingSuite.T(), nil, err)
		time.Sleep(120 * time.Second)
	
		balpDaiAfIssS41, _ := tradingSuite.getBalanceTokenIncAccount(
			tradingSuite.IncPrivKeyStr,
			tradingSuite.IncEtherTokenIDStr,
		)
		fmt.Println("[INFO] pETH balance after issuing step 3 : ", balpDaiAfIssS41)
	
		fmt.Println("------------ STEP 4: withdraw pETH to deposit ETH   --------------")
	
		withdrawingPETH := big.NewInt(0).Div(balEthScDepS2, big.NewInt(1000000000))
	
		burningRes, err = tradingSuite.callBurningPToken(
			tradingSuite.IncEtherTokenIDStr,
			withdrawingPETH,
			tradingSuite.ETHOwnerAddrStr,
			"createandsendburningrequest",
		)
		require.Equal(tradingSuite.T(), nil, err)
		burningTxID, found = burningRes["TxID"]
		require.Equal(tradingSuite.T(), true, found)
		time.Sleep(120 * time.Second)
	
		balpEthAfBurnS51, _ := tradingSuite.getBalanceTokenIncAccount(
			tradingSuite.IncPrivKeyStr,
			tradingSuite.IncEtherTokenIDStr,
		)
		fmt.Println("[INFO] pETH balance after burning step 4 : ", balpEthAfBurnS51)
	
		tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))
	
		balEthAfDep51 := tradingSuite.getBalanceOnETHNet(
			common.HexToAddress(tradingSuite.EtherAddressStr),
			common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		)
		fmt.Println("[INFO] ETH balance after withdraw  : ", balEthAfDep51)
	
	}
	




func (tradingSuite *VaultUpgradeTestSuite) Test8DepositAndWithdrwaERC20tokenOMG() {
//return
		fmt.Println("============ TEST 8 DEPOSIT AND WITHDRAW ERC20 TOKEN (OMG) ===========")
		fmt.Println("------------ STEP 0: declaration & initialization --------------")
	
		pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	
		// get info balance initialization
		balPrvInit, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
		fmt.Println("[INFO] PRV balance initialization : ", balPrvInit)
	
		balEthInit := tradingSuite.getBalanceOnETHNet(
			common.HexToAddress(tradingSuite.EtherAddressStr),
			common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		)
		fmt.Println("[INFO] ETH balance initialization : ", balEthInit)
	
		balpBATInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncOMGTokenIDStr)
		fmt.Println("[INFO] pOMG balance initialization : ", balpBATInit)
	
		balBATInit := tradingSuite.getBalanceOnETHNet(
			common.HexToAddress(tradingSuite.OMGAddressStr),
			common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		)
		fmt.Println("[INFO] OMG balance initialization : ", balBATInit)
	
		balEthScInit := tradingSuite.getDepositedBalance(
			tradingSuite.OMGAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] OMG balance initialization on SC : ", balEthScInit)
	
		fmt.Println("------------ STEP 1: porting OMG to pOMG --------------")
	
		fmt.Println("amount OMG deposit input : ", (big.NewInt(int64(0.01 * params.Ether))))
		deposit := big.NewInt(int64(0.01 * params.Ether))
		burningPETH := big.NewInt(0).Div(deposit, big.NewInt(1000000000))
		// Deposit to proof
		txHash := tradingSuite.depositERC20ToBridge(
			deposit,
			common.HexToAddress(tradingSuite.OMGAddressStr),
			tradingSuite.IncPaymentAddrStr,
		)
	
		balOMGAfDep := tradingSuite.getBalanceOnETHNet(
			common.HexToAddress(tradingSuite.OMGAddressStr),
			common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		)
		fmt.Println("[INFO] OMG balance remain after deposit : ", balOMGAfDep)
	
		fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHash))
		// get ETH remain after depsit
		balEthAfDep := tradingSuite.getBalanceOnETHNet(
			common.HexToAddress(tradingSuite.EtherAddressStr),
			common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		)
		fmt.Println("[INFO] ETH balance remain after deposit : ", balEthAfDep)
	
		//require.Equal(tradingSuite.T(), balEthAfDep, big.NewInt(0).Sub(big.NewInt(0).Sub(balEthInit, big.NewInt(int64(tradingSuite.DepositingEther*params.Ether))), tradingSuite.getGasFeeETHbyTxhash(txHash)), "balance ETH incorrect")
	
		// Proof
		_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
		require.Equal(tradingSuite.T(), nil, err)
		fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)
	
		fmt.Println("Waiting 90s for 15 blocks confirmation")
		time.Sleep(100 * time.Second)
		_, err = tradingSuite.callIssuingETHReq(
			tradingSuite.IncOMGTokenIDStr,
			ethDepositProof,
			ethBlockHash,
			ethTxIdx,
		)
		require.Equal(tradingSuite.T(), nil, err)
	
		time.Sleep(120 * time.Second)
	
		balpBATAfIssS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncOMGTokenIDStr)
		fmt.Println("[INFO] pOMG balance after issuing step 1 : ", balpBATAfIssS1)
		// check PRV and token balance after issuing
		balPrvAfIssS1, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
		fmt.Println("[INFO] PRV balance after issuing step 1: ", balPrvAfIssS1)
	
		fmt.Println("------------ STEP 2: burning pOMG to deposit OMG to SC --------------")
		// make a burn tx to incognito chain as a result of deposit to SC
		burningRes, err := tradingSuite.callBurningPToken(
			tradingSuite.IncOMGTokenIDStr,
			burningPETH,
			pubKeyToAddrStr[2:],
			"createandsendburningfordeposittoscrequest",
		)
		require.Equal(tradingSuite.T(), nil, err)
		burningTxID, found := burningRes["TxID"]
		require.Equal(tradingSuite.T(), true, found)
		time.Sleep(120 * time.Second)
	
		// check PRV and token balance after burning
		balPrvAfBurnS2, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
		fmt.Println("[INFO] PRV balance after burning step 2: ", balPrvAfBurnS2)
	
		balpBATAfBurnS2, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncOMGTokenIDStr)
		fmt.Println("[INFO] pOMG balance after burning step 2 : ", balpBATAfBurnS2)
		// TODO assert pETH balance issuing
	
		txHash2 := tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))
	
		fmt.Println("gas Fee transaction :", tradingSuite.getGasFeeETHbyTxhash(txHash2))
		// get ETH remain
		balEthAfDep2 := tradingSuite.getBalanceOnETHNet(
			common.HexToAddress(tradingSuite.EtherAddressStr),
			common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		)
		fmt.Println("[INFO] ETH balance remain  : ", balEthAfDep2)
		// TODO : assert ETH balance
		//require.Equal(tradingSuite.T(), balEthAfDep2, big.NewInt(0).Sub(balEthAfDep, tradingSuite.getGasFeeETHbyTxhash(txHash2)), "balance ETH incorrect")
	
		balBATScDepS2 := tradingSuite.getDepositedBalance(
			tradingSuite.OMGAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] OMG balance after deposit on SC at step 2: ", balBATScDepS2)
		// TODO assert ETH balane on SC
	
		fmt.Println("------------ STEP 3: withdraw OMG to deposit pOMG to Incognito  --------------")
	
		txHashByEmittingWithdrawalReq1 := tradingSuite.requestWithdraw(
			tradingSuite.OMGAddressStr,
			balBATScDepS2,
		)
	
		balBATScDepS41 := tradingSuite.getDepositedBalance(
			tradingSuite.OMGAddressStr,
			pubKeyToAddrStr,
		)
		fmt.Println("[INFO] OMG balance on SC after withdraw at step 3 : ", balBATScDepS41)
		// TODO assert BAT balane on SC
	
		_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq1)
		require.Equal(tradingSuite.T(), nil, err)
		fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)
	
		fmt.Println("Waiting 90s for 15 blocks confirmation")
		time.Sleep(100 * time.Second)
	
		_, err = tradingSuite.callIssuingETHReq(
			tradingSuite.IncOMGTokenIDStr,
			ethDepositProof,
			ethBlockHash,
			ethTxIdx,
		)
		require.Equal(tradingSuite.T(), nil, err)
		time.Sleep(100 * time.Second)
	
		balpBATAfIssS41, _ := tradingSuite.getBalanceTokenIncAccount(
			tradingSuite.IncPrivKeyStr,
			tradingSuite.IncOMGTokenIDStr,
		)
		fmt.Println("[INFO] pOMG balance after issuing step 3 : ", balpBATAfIssS41)
	
		fmt.Println("------------ STEP 4: withdraw pOMG to deposit OMG   --------------")
	
		withdrawingPBAT := big.NewInt(0).Div(balBATScDepS2, big.NewInt(1000000000))
	
		burningRes, err = tradingSuite.callBurningPToken(
			tradingSuite.IncOMGTokenIDStr,
			withdrawingPBAT,
			tradingSuite.ETHOwnerAddrStr,
			"createandsendburningrequest",
		)
		require.Equal(tradingSuite.T(), nil, err)
		burningTxID, found = burningRes["TxID"]
		require.Equal(tradingSuite.T(), true, found)
		time.Sleep(120 * time.Second)
	
		balpBATAfBurnS51, _ := tradingSuite.getBalanceTokenIncAccount(
			tradingSuite.IncPrivKeyStr,
			tradingSuite.IncOMGTokenIDStr,
		)
		fmt.Println("[INFO] pOMG balance after burning step 4 : ", balpBATAfBurnS51)
	
		tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))
	
		balBATAfDep51 := tradingSuite.getBalanceOnETHNet(
			common.HexToAddress(tradingSuite.OMGAddressStr),
			common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		)
		fmt.Println("[INFO] OMG balance after withdraw  : ", balBATAfDep51)
	
	}