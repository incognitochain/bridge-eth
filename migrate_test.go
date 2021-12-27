package bridge

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/incognitochain/bridge-eth/bridge/kbntrade"
	"github.com/incognitochain/bridge-eth/bridge/vault"
	"github.com/incognitochain/bridge-eth/bridge/vaultproxy"
	"github.com/pkg/errors"
)

func TestRetireVaultAdmin(t *testing.T) {
	privKey, c := getVaultProxy(t)

	successor := common.HexToAddress(Successor)
	fmt.Println("Successor address:", successor.Hex())

	auth := bind.NewKeyedTransactor(privKey)
	_, err := c.Retire(auth, successor)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClaimVaultAdmin(t *testing.T) {
	privKey, c := getVaultProxy(t)
	auth := bind.NewKeyedTransactor(privKey)
	_, err := c.Claim(auth)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateImplementation(t *testing.T) {
	privKey, client, chainID, err := connect()
	if err != nil {
		t.Fatal(err)
	}

	// Get vault instance
	c, err := vaultproxy.NewTransparentUpgradeableProxy(common.HexToAddress(BSCVaultAddress), client)
	if err != nil {
		t.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privKey, big.NewInt(chainID))
	if err != nil {
		t.Fatal(err)
	}
	auth.GasPrice = big.NewInt(5e9)

	tx, err := c.UpgradeTo(auth, common.HexToAddress(BSCNewImplementation))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Created transaction: %s\n", tx.Hash().String())
}

func TestClaimAllVaultAdmin(t *testing.T) {
	privKey, client, chainID, err := connect()
	if err != nil {
		t.Fatal(err)
	}

	// Get vault instance
	c, err := vaultproxy.NewTransparentUpgradeableProxy(common.HexToAddress(VaultAddress), client)
	if err != nil {
		t.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privKey, big.NewInt(chainID))
	if err != nil {
		t.Fatal(err)
	}
	tx, err := c.Claim(auth)
	if err != nil {
		t.Fatal(err)
	}

	// Get prev vault instance
	c, err = vaultproxy.NewTransparentUpgradeableProxy(common.HexToAddress(PrevVault), client)
	if err != nil {
		t.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(tx.Nonce()) + 1)

	_, err = c.Claim(auth)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeployNewVaultToMigrate(t *testing.T) {
	privKey, client, chainID, err := connect()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	admin := common.HexToAddress(Admin)
	incAddr := common.HexToAddress(IncognitoProxyAddress)
	prevVaultAddr := common.HexToAddress(VaultAddress)
	fmt.Println("Admin address:", admin.Hex())
	fmt.Println("IncognitoProxy address:", incAddr.Hex())
	fmt.Println("PrevVault address:", prevVaultAddr.Hex())

	// Deploy vault
	auth, err := bind.NewKeyedTransactorWithChainID(privKey, big.NewInt(chainID))
	if err != nil {
		t.Fatal(err)
	}
	auth.GasPrice = big.NewInt(int64(25000000000))
	vaultDelegatorAddr, _, _, err := vault.DeployVault(auth, client)
	if err != nil {
		t.Fatal(err)
	}

	vaultAbi, _ := abi.JSON(strings.NewReader(vault.VaultABI))
	input, _ := vaultAbi.Pack("initialize", common.Address{})

	vaultProxy, _, _, err := vaultproxy.DeployTransparentUpgradeableProxy(auth, client, vaultDelegatorAddr, admin, incAddr, input)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("deployed vault")
	fmt.Printf("addr: %s\n", vaultProxy.Hex())
}

func TestDeployKBNTrade(t *testing.T) {
	privKey, client, chainID, err := connect()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	newAddr := NewVaultTmp
	if len(newAddr) != 42 {
		t.Fatal(errors.New("invalid new vault's address"))
	}
	newVault := common.HexToAddress(newAddr)

	kbnProxy := common.HexToAddress(KBNProxy)
	fmt.Println("KBNProxy address:", kbnProxy.Hex())
	fmt.Println("New vault address:", newVault.Hex())

	// Deploy KBNTrade
	auth, err := bind.NewKeyedTransactorWithChainID(privKey, big.NewInt(chainID))
	if err != nil {
		t.Fatal(err)
	}
	auth.GasPrice = big.NewInt(int64(23000000000))
	kbnTradeAddr, _, _, err := kbntrade.DeployKBNTrade(auth, client, kbnProxy)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("deployed kbnTrade")
	fmt.Printf("addr: %s\n", kbnTradeAddr.Hex())
}

// func TestPauseVault(t *testing.T) {
// 	privKey, c := getVault(t)

// 	// Pause vault
// 	auth := bind.NewKeyedTransactor(privKey)
// 	_, err := c.Pause(auth)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

// func TestUnpauseVault(t *testing.T) {
// 	privKey, c := getVault(t)

// 	// Pause vault
// 	auth := bind.NewKeyedTransactor(privKey)
// 	_, err := c.Unpause(auth)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

// func TestMigrateVault(t *testing.T) {
// 	privKey, c := getVault(t)

// 	// Migrate vault
// 	newAddr := NewVaultTmp
// 	if len(newAddr) != 42 {
// 		t.Fatal(errors.New("invalid new vault's address"))
// 	}
// 	newVault := common.HexToAddress(newAddr)
// 	auth := bind.NewKeyedTransactor(privKey)
// 	_, err := c.Migrate(auth, newVault)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

func TestMoveAssetsVault(t *testing.T) {
	// privKey, c := getVault(t)

	// KNOWN FAILED
	// common.HexToAddress("0xc12d1c73ee7dc3615ba4e37e4abfdbddfa38907e"), // KickToken
	// common.HexToAddress("0xbddab785b306bcd9fb056da189615cc8ece1d823"), // Ebakus
	// common.HexToAddress("0x0e536b7831c7A7527FaD55da433986853d21A0c7"), // Harpoon
	// common.HexToAddress("0xc92e74b131D7b1D46E60e07F3FaE5d8877Dd03F0"), // Minereum
	// common.HexToAddress("0x426CA1eA2406c07d75Db9585F22781c096e3d0E0"), // MINEREUM

	// Migrate vault
	// assets := []common.Address{}

	// DONE
	// common.HexToAddress("0x2fB419E7023b32201e9aB3aba947f5c101a5C30e"), // Synth sEUR
	// common.HexToAddress("0x2987252148e34863612ac7f4ef3260de0c2a68f7"), // Synthetix Ne
	// common.HexToAddress("0xecF51a98B71f0421151a1d45E033Ab8B88665221"), // VAYLA
	// common.HexToAddress("0xD7F8032777C50aFD2e7AFa41912a4d8038127271"), // MUAN
	// common.HexToAddress("0x426CA1eA2406c07d75Db9585F22781c096e3d0E0"), // Minereum
	// common.HexToAddress("0x70Ec7702ADA8530d8f7332f7f3700099553D772D"), // Indonesian P
	// common.HexToAddress("0xA587469eE454A0911C5adF761754320319E7bD2F"), // easy ieo
	// common.HexToAddress("0x2505a3c035291c05cb78cb43cff39564637e1dd9"), // EVEN
	// common.HexToAddress("0xB351dA6ffEbd5DddD1dA037929FCf334d6B4A8D5"), // Flit Token
	// common.HexToAddress("0x2f141ce366a2462f02cea3d12cf93e4dca49e4fd"), // Free Coin
	// common.HexToAddress("0xf34845b76015d2952b6e39436bc59cae3c9ba17d"), // GIGS
	// common.HexToAddress("0xF1d50e8299687FC17378c6D5e2B25a7a0e07DcB4"), // BulkHead
	// common.HexToAddress("0x5d3dc0fdba0477b906ad4a36f95035b252060434"), // BulkHead
	// common.HexToAddress("0xfF95Ea9eBeFf204A95954bb1Ed76175354914Ea1"), // Cole Coin
	// common.HexToAddress("0xe0b9bcd54bf8a730ea5d3f1ffce0885e911a502c"), // Zum Token
	// common.HexToAddress("0x08130635368AA28b217a4dfb68E1bF8dC525621C"), // AfroDex
	// common.HexToAddress("0xc3761eb917cd790b30dad99f6cc5b4ff93c4f9ea"), // ERC20
	// common.HexToAddress("0x2b591e99afe9f32eaa6214f7b7629768c40eeb39"), // HEX
	// common.HexToAddress("0x5228a22e72ccc52d415ecfd199f99d0665e7733b"), // pTokens BTC
	// common.HexToAddress("0x16ea01acb4b0bca2000ee5473348b6937ee6f72f"), // Enecuum
	// common.HexToAddress("0x66fd97a78d8854fec445cd1c80a07896b0b4851f"), // Lunch Money
	// common.HexToAddress("0x2d184014b5658c453443aa87c8e9c4d57285620b"), // JSE Token
	// common.Address{}, // ETH
	// common.HexToAddress("0x4cc19356f2d37338b9802aa8e8fc58b0373296e7"), // SelfKey
	// common.HexToAddress("0x4f9254c83eb525f9fcf346490bbb3ed28a81c667"), // CelerToken
	// common.HexToAddress("0xead7f3ae4e0bb0d8785852cc37cc9d0b5e75c06a"), // 3X Long EOS
	// common.HexToAddress("0x4575f41308EC1483f3d399aa9a2826d74Da13Deb"), // Orchid
	// common.HexToAddress("0x799a4202c12ca952cb311598a024c80ed371a41e"), // HarmonyOne
	// common.HexToAddress("0x607F4C5BB672230e8672085532f7e901544a7375"), // RLC
	// common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7"), // Tether USD
	// common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"), // USD Coin
	// common.HexToAddress("0x6b175474e89094c44da98b954eedeac495271d0f"), // Dai Stableco
	// common.HexToAddress("0x7d1afa7b718fb893db30a3abc0cfc608aacfebb0"), // Matic Token
	// common.HexToAddress("0x4fabb145d64652a948d72533023f6e7a623c7c53"), // Binance USD
	// common.HexToAddress("0x0d8775f648430679a709e98d2b0cb6250d2887ef"), // BAT
	// common.HexToAddress("0x514910771af9ca656af840dff83e8264ecf986ca"), // ChainLink To
	// common.HexToAddress("0x89d24a6b4ccb1b6faa2625fe562bdd9a23260359"), // Sai Stableco
	// common.HexToAddress("0xb7cb1c96db6b22b0d3d9536e0108d062bd488f74"), // Walton
	// common.HexToAddress("0x8762db106b2c2a0bccb3a80d1ed41273552616e8"), // Reserve Righ
	// common.HexToAddress("0x408e41876cccdc0f92210600ef50372656052a38"), // Republic
	// common.HexToAddress("0xf0ee6b27b759c9893ce4f094b49ad28fd15a23e4"), // Enigma
	// common.HexToAddress("0xb63b606ac810a52cca15e44bb630fd42d8d1d83d"), // MCO
	// common.HexToAddress("0xdd974d5c2e2928dea5f71b9825b8b646686bd200"), // KyberNetwork
	// common.HexToAddress("0x8400d94a5cb0fa0d041a3788e395285d61c9ee5e"), // UniBright
	// common.HexToAddress("0x998FFE1E43fAcffb941dc337dD0468d52bA5b48A"), // Rupiah Token
	// common.HexToAddress("0x056fd409e1d7a124bd7017459dfea2f387b6d5cd"), // Gemini dolla
	// common.HexToAddress("0xae746520FfDB15d0505e32f1d6e9a2b4ab866572"), // MorCrypto Co
	// common.HexToAddress("0x41e5560054824ea6b0732e656e3ad64e20e94e45"), // Civic
	// common.HexToAddress("0x5432c580e34f590f4dd901b825ddeb92e905e826"), // TradeX Token
	// common.HexToAddress("0xb683d83a532e2cb7dfa5275eed3698436371cc9f"), // BTU Protocol
	// common.HexToAddress("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"), // OMG Network
	// common.HexToAddress("0x286BDA1413a2Df81731D4930ce2F862a35A609fE"), // WaBi
	// common.HexToAddress("0xf629cbd94d3791c9250152bd8dfbdf380e2a3b9c"), // EnjinCoin
	// common.HexToAddress("0x595832f8fc6bf59c85c527fec3740a1b7a361269"), // PowerLedger
	// common.HexToAddress("0x1d287cc25dad7ccaf76a26bc660c5f7c8e2a05bd"), // Fetch
	// common.HexToAddress("0xba11d00c5f74255f56a5e366f4f77f5a186d7f55"), // BandToken
	// common.HexToAddress("0x6c6ee5e31d828de241282b9606c8e98ea48526e2"), // HoloToken
	// common.HexToAddress("0x8e870d67f660d95d5be530380d0ec0bd388289e1"), // Paxos Standa
	// common.HexToAddress("0x1f573d6fb3f13d689ff844b4ce37794d79a7ff1c"), // Bancor
	// common.HexToAddress("0x50d1c9771902476076ecfc8b2a83ad6b9355a4c9"), // FTT
	// common.HexToAddress("0x4f3afec4e5a3f2a6a1a411def7d7dfe50ee057bf"), // Digix Gold T
	// common.HexToAddress("0xbbbbca6a901c926f240b89eacb641d8aec7aeafd"), // LoopringCoin
	// common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498"), // ZRX
	// common.HexToAddress("0x0000000000085d4780B73119b644AE5ecd22b376"), // TrueUSD
	// common.HexToAddress("0x1c5857e110cd8411054660f60b5de6a6958cfae2"), // Reserve
	// common.HexToAddress("0x5d3a536E4D6DbD6114cc1Ead35777bAB948E3643"), // Compound Dai
	// common.HexToAddress("0x5C406D99E04B8494dc253FCc52943Ef82bcA7D75"), // cUSD Currenc
	// common.HexToAddress("0x55296f69f40ea6d20e478533c15a6b08b654e758"), // XY Oracle
	// common.HexToAddress("0xa0b73e1ff0b80914ab6fe0444e65848c4c34450b"), // Crypto.com C
	// common.HexToAddress("0x4e15361fd6b4bb609fa63c81a2be19d873717870"), // Fantom Token
	// common.HexToAddress("0xAC8491258d2D93228E8b49aAC2e332A96f04E56C"), // 0xETH Cash
	// common.HexToAddress("0x8E9c3D1F30904E91764B7b8bBFDa3a429b886874"), // 0xETH SV
	// common.HexToAddress("0x71E4B8DE109428f999391eB3424D2CC87192e8bA"), // 0xETH Classi
	// common.HexToAddress("0x91f5f9c36B093907B0F99d1dBcf41aAF1db28916"), // 0xETH Diamon
	// common.HexToAddress("0x035df12e0f3ac6671126525f1015e47d79dfeddf"), // 0xMonero
	// common.HexToAddress("0x7db5454f3500f28171d1f9c7a38527c9cf94e6b2"), // Silver Stand
	// common.HexToAddress("0xAbf14EAc02407842a0140AD012239a03F8985404"), // betbeb.com A
	// common.HexToAddress("0xf222Ba8Af81d799C565241B0D3eEDF9Bdc4Fc462"), // betbeb.com空投
	// common.HexToAddress("0xc30951ff31c04a47b26ce496b0510a3b2d709e92"), // 启动公链
	// common.HexToAddress("0xcea83bc11a51cf07eea1286eee099871b428e613"), // Bitcoin 3
	// common.HexToAddress("0x3917E933bd430C08304cae2AA6d9746b806406c2"), // Bitcoin EVO
	// common.HexToAddress("0xe172f366678ec7b559f6c2913a437baadfd4e6c8"), // Kauri
	// common.HexToAddress("0xeD79E6dd91324F6Af138f01967BD24233d642a24"), // KING MAKER C
	// common.HexToAddress("0x1e24f42cb509e2af5675c0f5b529fc0b4c1a112a"), // Marek Adam W
	// common.HexToAddress("0x9D494A823Fc3E852f8fF92f36A05662A46de0381"), // Paybchain
	// common.HexToAddress("0x159A1dFAe19057de57dFfFcbB3DA1aE784678965"), // Reflex
	// common.HexToAddress("0x4df76a9dab9bb8310e4ad3dc4336a8e26ed24ebb"), // Sappchain
	// common.HexToAddress("0xaa7fb1c8ce6f18d4fd4aabb61a2193d4d441c54f"), // ShitCoin
	// common.HexToAddress("0xaC9Bb427953aC7FDDC562ADcA86CF42D988047Fd"), // Scatter.cx
	// common.HexToAddress("0xec91fcca41e8ab83dd5bc2bbcc2ffb71e314ba25"), // TaolCash
	// common.HexToAddress("0xf816725650497630642b52dbc3a734e118cf2ed2"), // Orius Ventur
	// common.HexToAddress("0xc00b89fc3a7ee7043629d8f9a79abfef55634960"), // Unicorn Coin
	// common.HexToAddress("0x3e3Aafa44d6E122b07d329b992F0DF62CF82b1e7"), // BEB公有链VPOW投票
	// common.HexToAddress("0x86C31E6da2190a1FFd39A36990a44174D0A8be15"), // VianeX
	// common.HexToAddress("0xdf413690fb785e56895551cc21086a15b6e90386"), // Vincoin Cash
	// common.HexToAddress("0x719035ac096b12aad44578a2db8a352ad874892d"), // VincoinCash
	// common.HexToAddress("0x0f90842e0b39fdc014dfe9daf5835c54ef894bf0"), // https://VPN.
	// common.HexToAddress("0x7b53B2C4B2F495d843a4e92e5c5511034d32bd15"), // VAYLA Token
	// common.HexToAddress("0x716523231368d43BDfe1F06AfE1C62930731aB13"), // Wrapped 0xEt
	// common.HexToAddress("0x767EE3150Ac31f982190Ef41728Cf9a969355286"), // Xamatek

	// auth := bind.NewKeyedTransactor(privKey)
	// _, err := c.MoveAssets(auth, assets)
	// if err != nil {
	// 	t.Fatal(err)
	// }
}

func getVault(t *testing.T) (*ecdsa.PrivateKey, *vault.Vault) {
	privKey, client, _, err := connect()
	if err != nil {
		t.Fatal(err)
	}

	// Get vault instance
	vaultAddr := common.HexToAddress(VaultAddress)
	c, err := vault.NewVault(vaultAddr, client)
	if err != nil {
		t.Fatal(err)
	}
	return privKey, c
}

func getVaultProxy(t *testing.T) (*ecdsa.PrivateKey, *vaultproxy.TransparentUpgradeableProxy) {
	privKey, client, _, err := connect()
	if err != nil {
		t.Fatal(err)
	}

	// Get vault instance
	vaultAddr := common.HexToAddress(VaultAddress)
	c, err := vaultproxy.NewTransparentUpgradeableProxy(vaultAddr, client)
	if err != nil {
		t.Fatal(err)
	}
	return privKey, c
}
