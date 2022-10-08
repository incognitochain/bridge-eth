package bridge

// Basic imports
import (
	"fmt"
	"math/big"
	"os"
	"strings"
	"testing"

	"github.com/incognitochain/bridge-eth/bridge/incognito_proxy"
	pancakeproxy "github.com/incognitochain/bridge-eth/bridge/pancake"
	"github.com/incognitochain/bridge-eth/bridge/pcurve"
	"github.com/incognitochain/bridge-eth/bridge/pdexbsc"
	"github.com/incognitochain/bridge-eth/bridge/pdexeth"
	"github.com/incognitochain/bridge-eth/bridge/prvbsc"
	"github.com/incognitochain/bridge-eth/bridge/prveth"
	puniswap "github.com/incognitochain/bridge-eth/bridge/puniswapproxy" // uniswap for polygon
	vaultaurora "github.com/incognitochain/bridge-eth/bridge/vaultaurora"
	vaultavax "github.com/incognitochain/bridge-eth/bridge/vaultavax"
	"github.com/incognitochain/bridge-eth/bridge/vaultbsc"
	"github.com/incognitochain/bridge-eth/bridge/vaultftm"
	"github.com/incognitochain/bridge-eth/bridge/vaultplg"

	// "github.com/incognitochain/bridge-eth/bridge/kbntrade"
	// "github.com/incognitochain/bridge-eth/bridge/uniswap"
	"github.com/incognitochain/bridge-eth/bridge/vault"
	"github.com/incognitochain/bridge-eth/bridge/vaultproxy"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type TradingDeployTestSuite struct {
	*TradingTestSuite

	KyberContractAddr         common.Address
	ZRXContractAddr           common.Address
	WETHAddr                  common.Address
	PWETHAddr                 common.Address
	UniswapRouteContractAddr  common.Address
	PancakeRouterContractAddr common.Address
	PolygonUniswapRouteAddr   common.Address
}

func NewTradingDeployTestSuite(tradingTestSuite *TradingTestSuite) *TradingDeployTestSuite {
	return &TradingDeployTestSuite{
		TradingTestSuite: tradingTestSuite,
	}
}

func (tradingDeploySuite *TradingDeployTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")
	// 0x kovan env
	tradingDeploySuite.KyberContractAddr = common.HexToAddress("0x692f391bCc85cefCe8C237C01e1f636BbD70EA4D")
	tradingDeploySuite.ZRXContractAddr = common.HexToAddress("0xf1ec01d6236d3cd881a0bf0130ea25fe4234003e")
	tradingDeploySuite.WETHAddr = common.HexToAddress("0xd0a1e359811322d97991e03f863a0c30c2cf029c")
	// uniswap router v2
	tradingDeploySuite.UniswapRouteContractAddr = common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D")
	// pancake v02
	tradingDeploySuite.PancakeRouterContractAddr = common.HexToAddress("0x9Ac64Cc6e4415144C455BD8E4837Fea55603e5c3")
	// polygon
	tradingDeploySuite.PolygonUniswapRouteAddr = common.HexToAddress("0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45")
	tradingDeploySuite.PWETHAddr = common.HexToAddress("0x9c3c9283d3e44854697cd22d3faa240cfb032889")
}

func (tradingDeploySuite *TradingDeployTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
	tradingDeploySuite.ETHClient.Close()
}

func TestTradingDeployTestSuite(t *testing.T) {
	fmt.Println("Starting entry point...")
	tradingSuite := new(TradingTestSuite)
	suite.Run(t, tradingSuite)

	tradingDeploySuite := NewTradingDeployTestSuite(tradingSuite)
	suite.Run(t, tradingDeploySuite)
	fmt.Println("Finishing entry point...")
}

func (tradingDeploySuite *TradingDeployTestSuite) TestDeployAllContracts() {
	admin := common.HexToAddress(Admin)
	regulator := common.HexToAddress(Regulator)
	executor := common.HexToAddress(Executor)
	fmt.Println("Admin address:", admin.Hex())
	fmt.Println("Regulator address:", regulator.Hex())
	fmt.Println("Rxecutor address:", executor.Hex())

	// Genesis committee
	// LOCAL
	// beaconComm, bridgeComm, err := convertCommittees(
	// 	localBeaconCommitteePubKeys, localBridgeCommitteePubKeys,
	// )

	// TESTNET2
	beaconComm, bridgeComm, err := convertCommittees(
		mainnetBeaconCommitteePubKeys, testnet2BridgeCommitteePubKeys,
	)
	// NOTE: uncomment this block to get mainnet committees when deploying to mainnet env
	/*
		beaconComm, bridgeComm, err := convertCommittees(
			mainnetBeaconCommitteePubKeys, mainnetBridgeCommitteePubKeys,
		)
	*/

	require.Equal(tradingDeploySuite.T(), nil, err)

	// Deploy incognito_proxy
	auth, err := bind.NewKeyedTransactorWithChainID(tradingDeploySuite.ETHPrivKey, big.NewInt(int64(tradingDeploySuite.ChainIDETH)))
	require.Equal(tradingDeploySuite.T(), nil, err)
	auth.Value = big.NewInt(0)
	auth.GasPrice = big.NewInt(1000000000)

	/*
		network:
		0 - deploy all
		1 - ETH
		2 - BSC
		3 - POLYGON
		4 - FANTOM
		5 - AURORA
		6 - AVALANCHE
	*/
	network := os.Getenv("NETWORK")

	if network == "1" || network == "0" {
		fmt.Println("============== DEPLOY INCOGNITO CONTRACT ON ETH ===============")

		incAddr, tx, _, err := incognito_proxy.DeployIncognitoProxy(auth, tradingDeploySuite.ETHClient, admin, beaconComm, bridgeComm)
		require.Equal(tradingDeploySuite.T(), nil, err)

		// incAddr := common.HexToAddress(IncognitoProxyAddress)
		fmt.Println("deployed incognito_proxy")
		fmt.Printf("addr: %s\n", incAddr.Hex())

		// Wait until tx is confirmed
		err = wait(tradingDeploySuite.ETHClient, tx.Hash())
		require.Equal(tradingDeploySuite.T(), nil, err)

		// Deploy vault
		prevVault := common.Address{}
		vaultAddr, tx, _, err := vault.DeployVault(auth, tradingDeploySuite.ETHClient)
		require.Equal(tradingDeploySuite.T(), nil, err)
		fmt.Println("deployed vault")
		fmt.Printf("addr: %s\n", vaultAddr.Hex())

		// Wait until tx is confirmed
		err = wait(tradingDeploySuite.ETHClient, tx.Hash())
		require.Equal(tradingDeploySuite.T(), nil, err)

		vaultAbi, _ := abi.JSON(strings.NewReader(vault.VaultABI))
		input, _ := vaultAbi.Pack("initialize", prevVault)

		// Deploy vault proxy
		vaultAddr, tx, _, err = vaultproxy.DeployTransparentUpgradeableProxy(auth, tradingDeploySuite.ETHClient, vaultAddr, admin, incAddr, input)
		require.Equal(tradingDeploySuite.T(), nil, err)
		fmt.Println("deployed vault proxy")
		fmt.Printf("addr: %s\n", vaultAddr.Hex())

		err = wait(tradingDeploySuite.ETHClient, tx.Hash())
		require.Equal(tradingDeploySuite.T(), nil, err)

		prvToken, tx, _, err := prveth.DeployPrveth(auth, tradingDeploySuite.ETHClient, "Incognito", "PRV", incAddr, vaultAddr)
		require.Equal(tradingDeploySuite.T(), nil, err)
		fmt.Println("deployed prv erc20 token")
		fmt.Printf("addr: %s\n", prvToken.Hex())

		err = wait(tradingDeploySuite.ETHClient, tx.Hash())
		require.Equal(tradingDeploySuite.T(), nil, err)

		prvToken, tx, _, err = pdexeth.DeployPdexeth(auth, tradingDeploySuite.ETHClient, "Incognito", "PDEX", incAddr, vaultAddr)
		require.Equal(tradingDeploySuite.T(), nil, err)
		fmt.Println("deployed pdex erc20 token")
		fmt.Printf("addr: %s\n", prvToken.Hex())

		err = wait(tradingDeploySuite.ETHClient, tx.Hash())
		require.Equal(tradingDeploySuite.T(), nil, err)

		if network != "0" {
			return
		}
	}

	if network == "2" || network == "0" {
		fmt.Println("============== DEPLOY INCOGNITO CONTRACT ON BSC ===============")

		auth, err = bind.NewKeyedTransactorWithChainID(tradingDeploySuite.ETHPrivKey, big.NewInt(int64(tradingDeploySuite.ChainIDBSC)))
		require.Equal(tradingDeploySuite.T(), nil, err)
		auth.GasPrice = big.NewInt(10000000000)
		incAddrBSC, tx, _, err := incognito_proxy.DeployIncognitoProxy(auth, tradingDeploySuite.BSCClient, admin, beaconComm, bridgeComm)
		require.Equal(tradingDeploySuite.T(), nil, err)

		// incAddr := common.HexToAddress(IncognitoProxyAddress)
		fmt.Println("deployed incognito_proxy")
		fmt.Printf("addr: %s\n", incAddrBSC.Hex())

		// Wait until tx is confirmed
		err = wait(tradingDeploySuite.BSCClient, tx.Hash())
		require.Equal(tradingDeploySuite.T(), nil, err)

		// Deploy vault
		vaultAddrBSC, tx, _, err := vaultbsc.DeployVaultbsc(auth, tradingDeploySuite.BSCClient)
		require.Equal(tradingDeploySuite.T(), nil, err)
		fmt.Println("deployed vault")
		fmt.Printf("addr: %s\n", vaultAddrBSC.Hex())

		// Wait until tx is confirmed
		err = wait(tradingDeploySuite.BSCClient, tx.Hash())
		require.Equal(tradingDeploySuite.T(), nil, err)

		prevVaultBSC := common.Address{}
		vaultAbi, _ := abi.JSON(strings.NewReader(vault.VaultABI))
		input, _ := vaultAbi.Pack("initialize", prevVaultBSC)

		// Deploy vault proxy
		vaultAddrBSC, tx, _, err = vaultproxy.DeployTransparentUpgradeableProxy(auth, tradingDeploySuite.BSCClient, vaultAddrBSC, admin, incAddrBSC, input)
		require.Equal(tradingDeploySuite.T(), nil, err)
		fmt.Println("deployed vault proxy")
		fmt.Printf("addr: %s\n", vaultAddrBSC.Hex())

		err = wait(tradingDeploySuite.BSCClient, tx.Hash())
		require.Equal(tradingDeploySuite.T(), nil, err)

		prvTokenBSC, tx, _, err := prvbsc.DeployPrvbsc(auth, tradingDeploySuite.BSCClient, "Incognito", "PRV", incAddrBSC, vaultAddrBSC)
		require.Equal(tradingDeploySuite.T(), nil, err)
		fmt.Println("deployed prv bep20 token")
		fmt.Printf("addr: %s\n", prvTokenBSC.Hex())

		err = wait(tradingDeploySuite.BSCClient, tx.Hash())
		require.Equal(tradingDeploySuite.T(), nil, err)

		prvTokenBSC, tx, _, err = pdexbsc.DeployPdexbsc(auth, tradingDeploySuite.BSCClient, "Incognito", "PDEX", incAddrBSC, vaultAddrBSC)
		require.Equal(tradingDeploySuite.T(), nil, err)
		fmt.Println("deployed pdex bep20 token")
		fmt.Printf("addr: %s\n", prvTokenBSC.Hex())

		err = wait(tradingDeploySuite.BSCClient, tx.Hash())
		require.Equal(tradingDeploySuite.T(), nil, err)

		pancakeProxyAddr, tx, _, err := pancakeproxy.DeployPancakeproxy(auth, tradingDeploySuite.BSCClient, tradingDeploySuite.PancakeRouterContractAddr)
		require.Equal(tradingDeploySuite.T(), nil, err)
		fmt.Println("deployed pancake proxy")
		fmt.Printf("addr: %s\n", pancakeProxyAddr.Hex())

		err = wait(tradingDeploySuite.BSCClient, tx.Hash())
		require.Equal(tradingDeploySuite.T(), nil, err)

		if network != "0" {
			return
		}
	}

	if network == "3" || network == "0" {
		fmt.Println("============== DEPLOY INCOGNITO CONTRACT ON POLYGON ===============")

		auth, err = bind.NewKeyedTransactorWithChainID(tradingDeploySuite.ETHPrivKey, big.NewInt(int64(tradingDeploySuite.ChainIDPLG)))
		require.Equal(tradingDeploySuite.T(), nil, err)
		auth.GasPrice = big.NewInt(10000000000)
		incAddrPLG, tx, _, err := incognito_proxy.DeployIncognitoProxy(auth, tradingDeploySuite.PLGClient, admin, beaconComm, bridgeComm)
		require.Equal(tradingDeploySuite.T(), nil, err)

		// incAddr := common.HexToAddress(IncognitoProxyAddress)
		fmt.Println("deployed incognito_proxy")
		fmt.Printf("addr: %s\n", incAddrPLG.Hex())

		// Wait until tx is confirmed
		err = wait(tradingDeploySuite.PLGClient, tx.Hash())
		require.Equal(tradingDeploySuite.T(), nil, err)

		// Deploy vault
		vaultAddrPLG, tx, _, err := vaultplg.DeployVaultplg(auth, tradingDeploySuite.PLGClient)
		require.Equal(tradingDeploySuite.T(), nil, err)
		fmt.Println("deployed vault")
		fmt.Printf("addr: %s\n", vaultAddrPLG.Hex())

		// Wait until tx is confirmed
		err = wait(tradingDeploySuite.PLGClient, tx.Hash())
		require.Equal(tradingDeploySuite.T(), nil, err)

		prevVaultBSC := common.Address{}
		vaultAbi, _ := abi.JSON(strings.NewReader(vault.VaultABI))
		input, _ := vaultAbi.Pack("initialize", prevVaultBSC)

		// Deploy vault proxy
		vaultAddrPLG, tx, _, err = vaultproxy.DeployTransparentUpgradeableProxy(auth, tradingDeploySuite.PLGClient, vaultAddrPLG, admin, incAddrPLG, input)
		require.Equal(tradingDeploySuite.T(), nil, err)
		fmt.Println("deployed vault proxy")
		fmt.Printf("addr: %s\n", vaultAddrPLG.Hex())

		err = wait(tradingDeploySuite.PLGClient, tx.Hash())
		require.Equal(tradingDeploySuite.T(), nil, err)

		// Deploy puniswap
		uniswapProxy, tx, _, err := puniswap.DeployPuniswap(auth, tradingDeploySuite.PLGClient, tradingDeploySuite.PolygonUniswapRouteAddr)
		require.Equal(tradingDeploySuite.T(), nil, err)
		fmt.Println("deployed puniswap proxy")
		fmt.Printf("addr: %s\n", uniswapProxy.Hex())

		err = wait(tradingDeploySuite.PLGClient, tx.Hash())
		require.Equal(tradingDeploySuite.T(), nil, err)

		// Deploy pcurve
		curveProxy, tx, _, err := pcurve.DeployPcurve(auth, tradingDeploySuite.PLGClient)
		require.Equal(tradingDeploySuite.T(), nil, err)
		fmt.Println("deployed curve proxy")
		fmt.Printf("addr: %s\n", curveProxy.Hex())

		err = wait(tradingDeploySuite.PLGClient, tx.Hash())
		require.Equal(tradingDeploySuite.T(), nil, err)

		if network != "0" {
			return
		}
	}

	if network == "4" || network == "0" {
		fmt.Println("============== DEPLOY INCOGNITO CONTRACT ON FANTOM ===============")

		auth, err = bind.NewKeyedTransactorWithChainID(tradingDeploySuite.ETHPrivKey, big.NewInt(int64(tradingDeploySuite.ChainIDFTM)))
		require.Equal(tradingDeploySuite.T(), nil, err)
		// auth.GasPrice = big.NewInt(500000)
		auth.GasLimit = uint64(5000000)
		incAddrFTM, tx, _, err := incognito_proxy.DeployIncognitoProxy(auth, tradingDeploySuite.FTMClient, admin, beaconComm, bridgeComm)
		require.Equal(tradingDeploySuite.T(), nil, err)

		// incAddr := common.HexToAddress(IncognitoProxyAddress)
		fmt.Println("deployed incognito_proxy")
		fmt.Printf("addr: %s\n", incAddrFTM.Hex())

		// Wait until tx is confirmed
		err = wait(tradingDeploySuite.FTMClient, tx.Hash())
		require.Equal(tradingDeploySuite.T(), nil, err)

		// Deploy vault
		vaultAddrFTM, tx, _, err := vaultftm.DeployVaultftm(auth, tradingDeploySuite.FTMClient)
		require.Equal(tradingDeploySuite.T(), nil, err)
		fmt.Println("deployed vault")
		fmt.Printf("addr: %s\n", vaultAddrFTM.Hex())

		// Wait until tx is confirmed
		err = wait(tradingDeploySuite.FTMClient, tx.Hash())
		require.Equal(tradingDeploySuite.T(), nil, err)

		prevVaultFTM := common.Address{}
		vaultAbi, _ := abi.JSON(strings.NewReader(vault.VaultMetaData.ABI))
		input, _ := vaultAbi.Pack("initialize", prevVaultFTM)

		// Deploy vault proxy
		vaultAddrFTM, tx, _, err = vaultproxy.DeployTransparentUpgradeableProxy(auth, tradingDeploySuite.FTMClient, vaultAddrFTM, admin, incAddrFTM, input)
		require.Equal(tradingDeploySuite.T(), nil, err)
		fmt.Println("deployed vault proxy")
		fmt.Printf("addr: %s\n", vaultAddrFTM.Hex())

		err = wait(tradingDeploySuite.FTMClient, tx.Hash())
		require.Equal(tradingDeploySuite.T(), nil, err)

		// // Deploy puniswap
		// uniswapProxy, tx, _, err := puniswap.DeployPuniswap(auth, tradingDeploySuite.PLGClient, tradingDeploySuite.PolygonUniswapRouteAddr)
		// require.Equal(tradingDeploySuite.T(), nil, err)
		// fmt.Println("deployed puniswap proxy")
		// fmt.Printf("addr: %s\n", uniswapProxy.Hex())

		// err = wait(tradingDeploySuite.PLGClient, tx.Hash())
		// require.Equal(tradingDeploySuite.T(), nil, err)

		// // Deploy pcurve
		// curveProxy, tx, _, err := pcurve.DeployPcurve(auth, tradingDeploySuite.PLGClient, tradingDeploySuite.PolygonUniswapRouteAddr)
		// require.Equal(tradingDeploySuite.T(), nil, err)
		// fmt.Println("deployed curve proxy")
		// fmt.Printf("addr: %s\n", curveProxy.Hex())

		// err = wait(tradingDeploySuite.PLGClient, tx.Hash())
		// require.Equal(tradingDeploySuite.T(), nil, err)

		if network != "0" {
			return
		}
	}

	if network == "5" || network == "0" {
		fmt.Println("============== DEPLOY INCOGNITO CONTRACT ON AURORA ===============")

		auth, err = bind.NewKeyedTransactorWithChainID(tradingDeploySuite.ETHPrivKey, big.NewInt(int64(tradingDeploySuite.ChainIDAURORA)))
		require.Equal(tradingDeploySuite.T(), nil, err)
		// auth.GasPrice = big.NewInt(500000)
		auth.GasLimit = uint64(5000000)
		incAddrAURORA, tx, _, err := incognito_proxy.DeployIncognitoProxy(auth, tradingDeploySuite.AURORAClient, admin, beaconComm, bridgeComm)
		require.Equal(tradingDeploySuite.T(), nil, err)

		// incAddr := common.HexToAddress(IncognitoProxyAddress)
		fmt.Println("deployed incognito_proxy")
		fmt.Printf("addr: %s\n", incAddrAURORA.Hex())

		// Wait until tx is confirmed
		err = wait(tradingDeploySuite.AURORAClient, tx.Hash())
		require.Equal(tradingDeploySuite.T(), nil, err)

		// Deploy vault
		vaultAddrAURORA, tx, _, err := vaultaurora.DeployVault(auth, tradingDeploySuite.AURORAClient)
		require.Equal(tradingDeploySuite.T(), nil, err)
		fmt.Println("deployed vault")
		fmt.Printf("addr: %s\n", vaultAddrAURORA.Hex())

		// Wait until tx is confirmed
		err = wait(tradingDeploySuite.AURORAClient, tx.Hash())
		require.Equal(tradingDeploySuite.T(), nil, err)

		prevVaultAURORA := common.Address{}
		vaultAbi, _ := abi.JSON(strings.NewReader(vaultaurora.VaultMetaData.ABI))
		input, err := vaultAbi.Pack("initialize", prevVaultAURORA, regulator, executor)
		require.Equal(tradingDeploySuite.T(), nil, err)

		// Deploy vault proxy
		vaultAddrAURORA, tx, _, err = vaultproxy.DeployTransparentUpgradeableProxy(auth, tradingDeploySuite.AURORAClient, vaultAddrAURORA, admin, incAddrAURORA, input)
		require.Equal(tradingDeploySuite.T(), nil, err)
		fmt.Println("deployed vault proxy")
		fmt.Printf("addr: %s\n", vaultAddrAURORA.Hex())

		err = wait(tradingDeploySuite.AURORAClient, tx.Hash())
		require.Equal(tradingDeploySuite.T(), nil, err)

		if network != "0" {
			return
		}
	}

	if network == "6" || network == "0" {
		fmt.Println("============== DEPLOY INCOGNITO CONTRACT ON AVAX ===============")

		auth, err = bind.NewKeyedTransactorWithChainID(tradingDeploySuite.ETHPrivKey, big.NewInt(int64(tradingDeploySuite.ChainIDAVAX)))
		require.Equal(tradingDeploySuite.T(), nil, err)
		// auth.GasPrice = big.NewInt(500000)
		auth.GasLimit = uint64(5000000)
		incAddrAVAX, tx, _, err := incognito_proxy.DeployIncognitoProxy(auth, tradingDeploySuite.AVAXClient, admin, beaconComm, bridgeComm)
		require.Equal(tradingDeploySuite.T(), nil, err)

		// incAddr := common.HexToAddress(IncognitoProxyAddress)
		fmt.Println("deployed incognito_proxy")
		fmt.Printf("addr: %s\n", incAddrAVAX.Hex())

		// Wait until tx is confirmed
		err = wait(tradingDeploySuite.AVAXClient, tx.Hash())
		require.Equal(tradingDeploySuite.T(), nil, err)

		// Deploy vault
		vaultAddrAVAX, tx, _, err := vaultavax.DeployVault(auth, tradingDeploySuite.AVAXClient)
		require.Equal(tradingDeploySuite.T(), nil, err)
		fmt.Println("deployed vault")
		fmt.Printf("addr: %s\n", vaultAddrAVAX.Hex())

		// Wait until tx is confirmed
		err = wait(tradingDeploySuite.AVAXClient, tx.Hash())
		require.Equal(tradingDeploySuite.T(), nil, err)

		prevVaultAVAX := common.Address{}
		vaultAbi, _ := abi.JSON(strings.NewReader(vaultavax.VaultMetaData.ABI))
		input, _ := vaultAbi.Pack("initialize", prevVaultAVAX, regulator, executor)

		// Deploy vault proxy
		vaultAddrAVAX, tx, _, err = vaultproxy.DeployTransparentUpgradeableProxy(auth, tradingDeploySuite.AVAXClient, vaultAddrAVAX, admin, incAddrAVAX, input)
		require.Equal(tradingDeploySuite.T(), nil, err)
		fmt.Println("deployed vault proxy")
		fmt.Printf("addr: %s\n", vaultAddrAVAX.Hex())

		err = wait(tradingDeploySuite.AVAXClient, tx.Hash())
		require.Equal(tradingDeploySuite.T(), nil, err)

		if network != "0" {
			return
		}
	}

	fmt.Println("============== NETWORK ONLY IN RANGE 0 - 6 ===============")
}

func convertCommittees(
	beaconComms []string, brigeComms []string,
) ([]common.Address, []common.Address, error) {
	beaconOld := make([]common.Address, len(beaconComms))
	for i, pk := range beaconComms {
		cpk := &CommitteePublicKey{}
		cpk.FromString(pk)
		addr, err := convertPubkeyToAddress(*cpk)
		if err != nil {
			return nil, nil, err
		}
		beaconOld[i] = addr
		fmt.Printf("beaconOld: %s\n", addr.Hex())
	}

	bridgeOld := make([]common.Address, len(brigeComms))
	for i, pk := range brigeComms {
		cpk := &CommitteePublicKey{}
		cpk.FromString(pk)
		addr, err := convertPubkeyToAddress(*cpk)
		if err != nil {
			return nil, nil, err
		}
		bridgeOld[i] = addr
		fmt.Printf("bridgeOld: %s\n", addr.Hex())
	}
	return beaconOld, bridgeOld, nil
}

// func TestDisplayCommitteesMainnet(t *testing.T) {
// 	fmt.Println("Mainnet Committees: [")
// 	beaconComms := mainnetBeaconCommitteePubKeys
// 	brigeComms := mainnetBridgeCommitteePubKeys
// 	beaconOld := make([]common.Address, len(beaconComms))
// 	for i, pk := range beaconComms {
// 		cpk := &CommitteePublicKey{}
// 		cpk.FromString(pk)
// 		addr, err := convertPubkeyToAddress(*cpk)
// 		if err != nil {
// 			return
// 		}
// 		beaconOld[i] = addr
// 		fmt.Printf("  %s,\n", addr.Hex())
// 	}
// 	fmt.Println("]")

// 	bridgeOld := make([]common.Address, len(brigeComms))
// 	for i, pk := range brigeComms {
// 		cpk := &CommitteePublicKey{}
// 		cpk.FromString(pk)
// 		addr, err := convertPubkeyToAddress(*cpk)
// 		if err != nil {
// 			return
// 		}
// 		bridgeOld[i] = addr
// 		fmt.Printf("%s,\n", addr.Hex())
// 	}
// }

// func TestDisplayCommitteesTestnet(t *testing.T) {
// 	fmt.Println("Testnet Committees: [")
// 	beaconComms := testnet2BridgeCommitteePubKeys
// 	brigeComms := testnet2BeaconCommitteePubKeys
// 	beaconOld := make([]common.Address, len(beaconComms))
// 	for i, pk := range beaconComms {
// 		cpk := &CommitteePublicKey{}
// 		cpk.FromString(pk)
// 		addr, err := convertPubkeyToAddress(*cpk)
// 		if err != nil {
// 			return
// 		}
// 		beaconOld[i] = addr
// 		fmt.Printf("  %s,\n", addr.Hex())
// 	}
// 	fmt.Println("]")

// 	bridgeOld := make([]common.Address, len(brigeComms))
// 	for i, pk := range brigeComms {
// 		cpk := &CommitteePublicKey{}
// 		cpk.FromString(pk)
// 		addr, err := convertPubkeyToAddress(*cpk)
// 		if err != nil {
// 			return
// 		}
// 		bridgeOld[i] = addr
// 		fmt.Printf("%s,\n", addr.Hex())
// 	}
// }
