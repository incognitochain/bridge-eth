package bridge

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"strings"
	"time"

	"github.com/incognitochain/bridge-eth/bridge/prveth"
	"github.com/incognitochain/bridge-eth/bridge/vault"
	"github.com/incognitochain/bridge-eth/common/base58"
	"github.com/incognitochain/bridge-eth/consensus/signatureschemes/bridgesig"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/sha3"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/incognitochain/bridge-eth/erc20"
	"github.com/incognitochain/bridge-eth/rpccaller"
	"github.com/stretchr/testify/require"
)

const (
	EXECUTE_PREFIX          = 0
	REQ_WITHDRAW_PREFIX     = 1
	BSC_EXECUTE_PREFIX      = 2
	BSC_REQ_WITHDRAW_PREFIX = 3
	PLG_EXECUTE_PREFIX      = 4
	PLG_REQ_WITHDRAW_PREFIX = 5
	FTM_EXECUTE_PREFIX      = 6
	FTM_REQ_WITHDRAW_PREFIX = 7
)

type IssuingETHRes struct {
	rpccaller.RPCBaseRes
	Result interface{} `json:"Result"`
}

type BurningForDepositToSCRes struct {
	rpccaller.RPCBaseRes
	Result interface{} `json:"Result"`
}

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type TradingTestSuite struct {
	suite.Suite
	IncBurningAddrStr string
	IncPrivKeyStr     string
	IncPaymentAddrStr string

	GeneratedPrivKeyForSC ecdsa.PrivateKey
	GeneratedPubKeyForSC  ecdsa.PublicKey

	IncEtherTokenIDStr string
	IncDAITokenIDStr   string
	IncUSDTTokenIDStr  string
	IncSAITokenIDStr   string
	IncPRVTokenIDStr   string
	IncPDEXTokenIDStr  string
	IncFTMTokenIDStr   string

	IncBridgeHost string
	IncRPCHost    string

	EtherAddressStr string
	DAIAddressStr   string
	SAIAddressStr   string

	ETHPrivKeyStr   string
	ETHOwnerAddrStr string

	ETHHost    string
	ETHPrivKey *ecdsa.PrivateKey
	ETHClient  *ethclient.Client

	BSCHost   string
	BSCClient *ethclient.Client

	PLGHost   string
	PLGClient *ethclient.Client

	FTMHost   string
	FTMClient *ethclient.Client

	AURORAHost   string
	AURORAClient *ethclient.Client

	AVAXHost   string
	AVAXClient *ethclient.Client

	ChainIDETH    uint
	ChainIDBSC    uint
	ChainIDPLG    uint
	ChainIDFTM    uint
	ChainIDAURORA uint
	ChainIDAVAX   uint

	VaultAddr            common.Address
	VaultBSCAddr         common.Address
	VaultPLGAddr         common.Address
	VaultFTMAddr         common.Address
	KBNTradeDeployedAddr common.Address
	PRVERC20Addr         common.Address
	PRVBEP20Addr         common.Address
	PDEXERC20Addr        common.Address
	PDEXBEP20Addr        common.Address

	KyberContractAddr common.Address
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (tradingSuite *TradingTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")

	// 0x kovan env
	tradingSuite.IncBurningAddrStr = "12RxahVABnAVCGP3LGwCn8jkQxgw7z1x14wztHzn455TTVpi1wBq9YGwkRMQg3J4e657AbAnCvYCJSdA9czBUNuCKwGSRQt55Xwz8WA"
	tradingSuite.IncPrivKeyStr = "112t8roafGgHL1rhAP9632Yef3sx5k8xgp8cwK4MCJsCL1UWcxXvpzg97N4dwvcD735iKf31Q2ZgrAvKfVjeSUEvnzKJyyJD3GqqSZdxN4or"
	tradingSuite.IncPaymentAddrStr = "12svfkP6w5UDJDSCwqH978PvqiqBxKmUnA9em9yAYWYJVRv7wuXY1qhhYpPAm4BDz2mLbFrRmdK3yRhnTqJCZXKHUmoi7NV83HCH2YFpctHNaDdkSiQshsjw2UFUuwdEvcidgaKmF3VJpY5f8RdN"

	tradingSuite.IncEtherTokenIDStr = "b366fa400c36e6bbcf24ac3e99c90406ddc64346ab0b7ba21e159b83d938812d"
	tradingSuite.IncUSDTTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000098"
	tradingSuite.IncSAITokenIDStr = "0000000000000000000000000000000000000000000000000000000000000097"
	tradingSuite.IncDAITokenIDStr = "0000000000000000000000000000000000000000000000000000000000000096"
	tradingSuite.IncPRVTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000004"
	tradingSuite.IncPDEXTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000006"
	tradingSuite.IncFTMTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000100"

	tradingSuite.EtherAddressStr = "0x0000000000000000000000000000000000000000"
	tradingSuite.DAIAddressStr = "0x4f96fe3b7a6cf9725f59d353f723c1bdb64ca6aa"
	tradingSuite.SAIAddressStr = "0xc4375b7de8af5a38a93548eb8453a498222c4ff2"

	tradingSuite.ETHPrivKeyStr = "aad53b70ad9ed01b75238533dd6b395f4d300427da0165aafbd42ea7a606601f"
	tradingSuite.ETHOwnerAddrStr = "0xD7d93b7fa42b60b6076f3017fCA99b69257A912D"

	tradingSuite.ETHHost = "https://goerli.infura.io/v3/1138a1e99b154b10bae5c382ad894361"
	tradingSuite.BSCHost = "https://data-seed-prebsc-1-s1.binance.org:8545"
	tradingSuite.PLGHost = "https://polygon-mumbai.infura.io/v3/9bc873177cf74a03a35739e45755a9ac"
	tradingSuite.FTMHost = "https://rpc.testnet.fantom.network"
	tradingSuite.AVAXHost = "https://api.avax-test.network/ext/C/rpc"
	tradingSuite.AURORAHost = "https://testnet.aurora.dev"

	tradingSuite.IncBridgeHost = "http://127.0.0.1:9338"
	// tradingSuite.IncBridgeHost = "http://127.0.0.1:9350" // 0xkraken
	tradingSuite.IncRPCHost = "http://127.0.0.1:9334"

	// testnet 2 vaults
	tradingSuite.VaultAddr = common.HexToAddress("0x9cb4baf1b60dabb6b22bcff07cc0e10395423aed")
	tradingSuite.VaultBSCAddr = common.HexToAddress("0x2f6F03F1b43Eab22f7952bd617A24AB46E970dF7")
	tradingSuite.VaultPLGAddr = common.HexToAddress("0x4fF5c88cD1FD773436C2aBcFE175fe4ba6a2eB68")
	tradingSuite.VaultFTMAddr = common.HexToAddress("0xDdF0C1B9f513a8D0a26732B18478CC295d204320")

	tradingSuite.PRVERC20Addr = common.HexToAddress("0xf4933b0288644778f6f2264EaB009fD04fF669a1")
	tradingSuite.PRVBEP20Addr = common.HexToAddress("0x5A15626f6beA715870D46f43f50bE9821368963f")
	tradingSuite.PDEXERC20Addr = common.HexToAddress("0x9c59b98fcC33f2859A2aB11BC2aAfDcf513b6c33")
	tradingSuite.PDEXBEP20Addr = common.HexToAddress("0xa43F2911dF4a560A1F687Eba359D047753Cd9BD9")

	tradingSuite.ChainIDBSC = 97
	tradingSuite.ChainIDETH = 5
	tradingSuite.ChainIDPLG = 80001
	tradingSuite.ChainIDFTM = 4002
	tradingSuite.ChainIDAURORA = 1313161555
	tradingSuite.ChainIDAVAX = 43113

	// generate a new keys pair for SC
	tradingSuite.genKeysPairForSC()

	// connect to ethereum network
	tradingSuite.connectToETH()
}

func (tradingSuite *TradingTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
	tradingSuite.ETHClient.Close()
}

func (tradingSuite *TradingTestSuite) SetupTest() {
	fmt.Println("Setting up the test...")
}

func (tradingSuite *TradingTestSuite) TearDownTest() {
	fmt.Println("Tearing down the test...")
}

func (tradingSuite *TradingTestSuite) TestTradingTestSuite() {
	fmt.Println("This is generic test suite")
}

func (tradingSuite *TradingTestSuite) getBalanceOnETHNet(
	tokenAddr common.Address,
	ownerAddr common.Address,
	client *ethclient.Client,
) *big.Int {
	if tokenAddr.Hex() == tradingSuite.EtherAddressStr {
		balance, err := client.BalanceAt(context.Background(), ownerAddr, nil)
		require.Equal(tradingSuite.T(), nil, err)
		return balance
	}
	// erc20 token
	instance, err := erc20.NewErc20(tokenAddr, client)
	require.Equal(tradingSuite.T(), nil, err)

	balance, err := instance.BalanceOf(&bind.CallOpts{}, ownerAddr)
	require.Equal(tradingSuite.T(), nil, err)
	return balance
}

func (tradingSuite *TradingTestSuite) connectToETH() {
	privKeyHex := tradingSuite.ETHPrivKeyStr
	privKey, err := crypto.HexToECDSA(privKeyHex)
	require.Equal(tradingSuite.T(), nil, err)

	fmt.Printf("Sign Txs with address: %s\n", crypto.PubkeyToAddress(privKey.PublicKey).Hex())

	network := "development"
	fmt.Printf("Connecting to network %s\n", network)
	client, err := ethclient.Dial(tradingSuite.ETHHost)
	require.Equal(tradingSuite.T(), nil, err)
	tradingSuite.ETHClient = client

	client, err = ethclient.Dial(tradingSuite.BSCHost)
	require.Equal(tradingSuite.T(), nil, err)
	tradingSuite.BSCClient = client

	client, err = ethclient.Dial(tradingSuite.PLGHost)
	require.Equal(tradingSuite.T(), nil, err)
	tradingSuite.PLGClient = client

	client, err = ethclient.Dial(tradingSuite.FTMHost)
	require.Equal(tradingSuite.T(), nil, err)
	tradingSuite.FTMClient = client

	client, err = ethclient.Dial(tradingSuite.AURORAHost)
	require.Equal(tradingSuite.T(), nil, err)
	tradingSuite.AURORAClient = client

	client, err = ethclient.Dial(tradingSuite.AVAXHost)
	require.Equal(tradingSuite.T(), nil, err)
	tradingSuite.AVAXClient = client

	tradingSuite.ETHPrivKey = privKey
}

func (tradingSuite *TradingTestSuite) depositETH(
	amt float64,
	incPaymentAddrStr string,
	vaultAddr common.Address,
	client *ethclient.Client,
) common.Hash {
	c, err := vault.NewVault(vaultAddr, client)
	require.Equal(tradingSuite.T(), nil, err)
	chainID, err := client.ChainID(auth.Context)
	require.Equal(tradingSuite.T(), nil, err)
	auth, err := bind.NewKeyedTransactorWithChainID(tradingSuite.ETHPrivKey, chainID)
	require.Equal(tradingSuite.T(), nil, err)
	auth.Value = new(big.Int).SetUint64(uint64(amt * params.Ether))
	txId := [32]byte{}
	// todo: update genesisAcc2.PrivateKey to real regulator
	sign, err := SignDataToShield(txId, genesisAcc2.PrivateKey, auth.From)
	require.Equal(tradingSuite.T(), nil, err)
	tx, err := c.Deposit(auth, incPaymentAddrStr, txId, sign)
	require.Equal(tradingSuite.T(), nil, err)
	txHash := tx.Hash()

	if err := wait(client, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("deposited, txHash: %x\n", txHash[:])
	return txHash
}

func (tradingSuite *TradingTestSuite) depositERC20ToBridge(
	amt *big.Int,
	tokenAddr common.Address,
	incPaymentAddrStr string,
	vaultAddr common.Address,
	client *ethclient.Client,
	chainID uint,
) common.Hash {
	auth, err := bind.NewKeyedTransactorWithChainID(tradingSuite.ETHPrivKey, big.NewInt(int64(chainID)))
	require.Equal(tradingSuite.T(), nil, err)
	c, err := vault.NewVault(vaultAddr, client)
	require.Equal(tradingSuite.T(), nil, err)

	erc20Token, _ := erc20.NewErc20(tokenAddr, client)
	auth.GasPrice = big.NewInt(1e10)
	tx2, apprErr := erc20Token.Approve(auth, vaultAddr, amt)
	require.Equal(tradingSuite.T(), nil, apprErr)
	tx2Hash := tx2.Hash()
	fmt.Printf("Approve tx, txHash: %x\n", tx2Hash[:])
	time.Sleep(15 * time.Second)
	auth.GasPrice = big.NewInt(1e10)

	fmt.Println("Starting deposit erc20 to vault contract")
	txId := [32]byte{}
	// todo: update genesisAcc2.PrivateKey to real regulator
	sign, err := SignDataToShield(txId, genesisAcc2.PrivateKey, auth.From)
	require.Equal(tradingSuite.T(), nil, err)
	tx, err := c.DepositERC20(auth, tokenAddr, amt, incPaymentAddrStr, txId, sign)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("Finished deposit erc20 to vault contract")
	txHash := tx.Hash()

	if err := wait(client, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("deposited erc20 token to bridge, txHash: %x\n", txHash[:])
	return txHash
}

func (tradingSuite *TradingTestSuite) callIssuingETHReq(
	incTokenIDStr string,
	ethDepositProof []string,
	ethBlockHash string,
	ethTxIdx uint,
	method string,
) (map[string]interface{}, error) {
	rpcClient := rpccaller.NewRPCClient()
	meta := map[string]interface{}{
		"IncTokenID": incTokenIDStr,
		"BlockHash":  ethBlockHash,
		"ProofStrs":  ethDepositProof,
		"TxIndex":    ethTxIdx,
	}
	params := []interface{}{
		tradingSuite.IncPrivKeyStr,
		nil,
		5,
		-1,
		meta,
	}
	var res IssuingETHRes
	err := rpcClient.RPCCall(
		"",
		tradingSuite.IncRPCHost,
		"",
		method,
		params,
		&res,
	)
	if err != nil {
		return nil, err
	}

	response, _ := json.Marshal(res)
	fmt.Println("get response", string(response))

	if res.RPCError != nil {
		return nil, errors.New(res.RPCError.Message)
	}
	return res.Result.(map[string]interface{}), nil
}

func (tradingSuite *TradingTestSuite) callBurningPToken(
	incTokenIDStr string,
	amount *big.Int,
	remoteAddrStr string,
	burningMethod string,
) (map[string]interface{}, error) {
	rpcClient := rpccaller.NewRPCClient()
	meta := map[string]interface{}{
		"TokenID":     incTokenIDStr,
		"TokenTxType": 1,
		"TokenName":   "",
		"TokenSymbol": "",
		"TokenAmount": amount.Uint64(),
		"TokenReceivers": map[string]uint64{
			tradingSuite.IncBurningAddrStr: amount.Uint64(),
		},
		"RemoteAddress": remoteAddrStr,
		"Privacy":       true,
		"TokenFee":      0,
	}
	params := []interface{}{
		tradingSuite.IncPrivKeyStr,
		nil,
		-1,
		0,
		meta,
		"",
		0,
	}
	var res BurningForDepositToSCRes
	err := rpcClient.RPCCall(
		"",
		tradingSuite.IncRPCHost,
		"",
		burningMethod,
		params,
		&res,
	)
	if err != nil {
		fmt.Println("calling burning ptokens err: ", err)
		return nil, err
	}
	bb, _ := json.Marshal(res)
	fmt.Println("calling burning ptokens res: ", string(bb))
	if res.RPCError != nil {
		return nil, errors.New(res.RPCError.Message)
	}
	return res.Result.(map[string]interface{}), nil
}

func (tradingSuite *TradingTestSuite) callBurningPRV(
	amount *big.Int,
	remoteAddrStr string,
	burningMethod string,
) (map[string]interface{}, error) {
	rpcClient := rpccaller.NewRPCClient()
	meta := map[string]interface{}{
		"TokenID":     tradingSuite.IncPRVTokenIDStr,
		"TokenTxType": 1,
		"TokenName":   "",
		"TokenSymbol": "",
		"TokenAmount": amount.Uint64(),
		"TokenReceivers": map[string]uint64{
			tradingSuite.IncBurningAddrStr: amount.Uint64(),
		},
		"RemoteAddress": remoteAddrStr,
		"Privacy":       true,
		"TokenFee":      0,
	}
	params := []interface{}{
		tradingSuite.IncPrivKeyStr,
		map[string]interface{}{
			tradingSuite.IncBurningAddrStr: amount.Uint64(),
		},
		-1,
		0,
		meta,
		"",
		0,
	}
	var res BurningForDepositToSCRes
	err := rpcClient.RPCCall(
		"",
		tradingSuite.IncRPCHost,
		"",
		burningMethod,
		params,
		&res,
	)
	if err != nil {
		fmt.Println("calling burning ptokens err: ", err)
		return nil, err
	}
	bb, _ := json.Marshal(res)
	fmt.Println("calling burning ptokens res: ", string(bb))
	if res.RPCError != nil {
		return nil, errors.New(res.RPCError.Message)
	}
	return res.Result.(map[string]interface{}), nil
}

func (tradingSuite *TradingTestSuite) submitBurnProofForDepositToSC(
	burningTxIDStr string,
	chainID *big.Int,
	method string,
	vaultAddr common.Address,
	client *ethclient.Client,
) {
	proof, err := getAndDecodeBurnProofV2(tradingSuite.IncBridgeHost, burningTxIDStr, method)
	require.Equal(tradingSuite.T(), nil, err)

	// Get contract instance
	c, err := vault.NewVault(vaultAddr, client)
	require.Equal(tradingSuite.T(), nil, err)

	// Burn
	auth, err := bind.NewKeyedTransactorWithChainID(tradingSuite.ETHPrivKey, chainID)
	require.Equal(tradingSuite.T(), nil, err)
	auth.GasPrice = big.NewInt(1e10)
	// auth.GasLimit = uint64(5000000) // for FTM testnet
	tx, err := SubmitBurnProof(c, auth, proof)
	require.Equal(tradingSuite.T(), nil, err)

	txHash := tx.Hash()
	if err := wait(client, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("burned, txHash: %x\n", txHash[:])
}

func (tradingSuite *TradingTestSuite) submitBurnProofForWithdrawal(
	burningTxIDStr string,
	method string,
	vaultAddr common.Address,
	client *ethclient.Client,
	chainID uint,
) {
	proof, err := getAndDecodeBurnProofV2(tradingSuite.IncBridgeHost, burningTxIDStr, method)
	require.Equal(tradingSuite.T(), nil, err)

	// Get contract instance
	c, err := vault.NewVault(vaultAddr, client)
	require.Equal(tradingSuite.T(), nil, err)

	// Burn
	auth, err := bind.NewKeyedTransactorWithChainID(tradingSuite.ETHPrivKey, big.NewInt(int64(chainID)))
	require.Equal(tradingSuite.T(), nil, err)
	auth.GasPrice = big.NewInt(1e9)
	// auth.GasLimit = uint64(5000000) // for FTM testnet
	tx, err := Withdraw(c, auth, proof)
	require.Equal(tradingSuite.T(), nil, err)

	txHash := tx.Hash()
	if err := wait(client, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("burned, txHash: %x\n", txHash[:])
}

func (tradingSuite *TradingTestSuite) submitBurnProofForMintPRV(
	burningTxIDStr string,
	contractAddress common.Address,
	method string,
	clientInst *ethclient.Client,
	chainID int64,
) {
	proof, err := getAndDecodeBurnProofV2(tradingSuite.IncBridgeHost, burningTxIDStr, method)
	require.Equal(tradingSuite.T(), nil, err)

	// Get contract instance
	c, err := prveth.NewPrveth(contractAddress, clientInst)
	require.Equal(tradingSuite.T(), nil, err)

	// Burn
	auth, err := bind.NewKeyedTransactorWithChainID(tradingSuite.ETHPrivKey, big.NewInt(chainID))
	require.Equal(tradingSuite.T(), nil, err)
	auth.GasPrice = big.NewInt(1e10)
	tx, err := SubmitMintPRVProof(c, auth, proof)
	require.Equal(tradingSuite.T(), nil, err)

	txHash := tx.Hash()
	if err := wait(clientInst, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("mint evm prv, txHash: %x\n", txHash[:])
}

func (tradingSuite *TradingTestSuite) genKeysPairForSC() {
	incPriKeyBytes, _, err := base58.Base58Check{}.Decode(tradingSuite.IncPrivKeyStr)
	require.Equal(tradingSuite.T(), nil, err)

	tradingSuite.GeneratedPrivKeyForSC, tradingSuite.GeneratedPubKeyForSC = bridgesig.KeyGen(incPriKeyBytes)
}

func randomizeTimestamp() string {
	randomTime := rand.Int63n(time.Now().Unix()-94608000) + 94608000
	randomNow := time.Unix(randomTime, 0)
	return randomNow.String()
}

func rawsha3(b []byte) []byte {
	hashF := sha3.NewLegacyKeccak256()
	hashF.Write(b)
	buf := hashF.Sum(nil)
	return buf
}

func rlpHash(x interface{}) (h common.Hash) {
	hw := sha3.NewLegacyKeccak256()
	rlp.Encode(hw, x)
	hw.Sum(h[:0])
	return h
}

func (tradingSuite *TradingTestSuite) getDepositedBalance(
	ethTokenAddrStr string,
	ownerAddrStr string,
) *big.Int {
	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)
	token := common.HexToAddress(ethTokenAddrStr)
	owner := common.HexToAddress(ownerAddrStr)
	bal, err := c.GetDepositedBalance(nil, token, owner)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Printf("deposited balance: %d\n", bal)
	return bal
}

func (tradingSuite *TradingTestSuite) getDepositedBalanceBSC(
	token common.Address,
	ownerAddrStr string,
) *big.Int {
	c, err := vault.NewVault(tradingSuite.VaultBSCAddr, tradingSuite.BSCClient)
	require.Equal(tradingSuite.T(), nil, err)
	owner := common.HexToAddress(ownerAddrStr)
	bal, err := c.GetDepositedBalance(nil, token, owner)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Printf("deposited balance: %d\n", bal)
	return bal
}

func (tradingSuite *TradingTestSuite) getDepositedBalancePLG(
	token common.Address,
	ownerAddrStr string,
) *big.Int {
	c, err := vault.NewVault(tradingSuite.VaultPLGAddr, tradingSuite.PLGClient)
	require.Equal(tradingSuite.T(), nil, err)
	owner := common.HexToAddress(ownerAddrStr)
	bal, err := c.GetDepositedBalance(nil, token, owner)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Printf("deposited balance: %d\n", bal)
	return bal
}

func (tradingSuite *TradingTestSuite) getDepositedBalanceFTM(
	token common.Address,
	ownerAddrStr string,
) *big.Int {
	c, err := vault.NewVault(tradingSuite.VaultFTMAddr, tradingSuite.FTMClient)
	require.Equal(tradingSuite.T(), nil, err)
	owner := common.HexToAddress(ownerAddrStr)
	bal, err := c.GetDepositedBalance(nil, token, owner)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Printf("deposited balance: %d\n", bal)
	return bal
}

func (tradingSuite *TradingTestSuite) requestWithdraw(
	withdrawalETHTokenIDStr string,
	amount *big.Int,
	client *ethclient.Client,
	chainID *big.Int,
	vaultAddrr common.Address,
	signaturePrefix uint8,
) common.Hash {
	c, err := vault.NewVault(vaultAddrr, client)
	require.Equal(tradingSuite.T(), nil, err)
	auth, err := bind.NewKeyedTransactorWithChainID(tradingSuite.ETHPrivKey, chainID)
	require.Equal(tradingSuite.T(), nil, err)
	token := common.HexToAddress(withdrawalETHTokenIDStr)
	timestamp := []byte(randomizeTimestamp())
	vaultAbi, _ := abi.JSON(strings.NewReader(vault.VaultHelperABI))
	psData := vault.VaultHelperPreSignData{
		Prefix:    signaturePrefix,
		Token:     token,
		Timestamp: timestamp,
		Amount:    amount,
	}
	tempData, _ := vaultAbi.Pack("_buildSignRequestWithdraw", psData, tradingSuite.IncPaymentAddrStr)
	data := rawsha3(tempData[4:])
	signBytes, _ := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)
	auth.GasPrice = big.NewInt(1e10)
	// auth.GasLimit = uint64(5000000) // for FTM testnet

	txId := [32]byte{}
	// todo: update genesisAcc2.PrivateKey to real regulator
	sign, err := SignDataToShield(txId, genesisAcc2.PrivateKey, auth.From)
	require.Equal(tradingSuite.T(), nil, err)
	tx, err := c.RequestWithdraw(auth, tradingSuite.IncPaymentAddrStr, token, amount, signBytes, timestamp, txId, sign)
	require.Equal(tradingSuite.T(), nil, err)

	txHash := tx.Hash()
	if err := wait(client, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("request withdrawal, txHash: %x\n", txHash[:])
	return txHash
}

func (tradingSuite *TradingTestSuite) burnPRV(
	amt float64,
	incPaymentAddrStr string,
	contractAddress common.Address,
	clientInst *ethclient.Client,
	chainID int64,
) common.Hash {
	c, err := prveth.NewPrveth(contractAddress, clientInst)
	require.Equal(tradingSuite.T(), nil, err)
	auth, err := bind.NewKeyedTransactorWithChainID(tradingSuite.ETHPrivKey, big.NewInt(chainID))
	require.Equal(tradingSuite.T(), nil, err)
	auth.GasPrice = big.NewInt(1e10)

	tx, err := c.Burn(auth, incPaymentAddrStr, big.NewInt(int64(amt*float64(1e9))))
	require.Equal(tradingSuite.T(), nil, err)
	txHash := tx.Hash()

	if err := wait(clientInst, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("burn prv token, txHash: %x\n", txHash[:])
	return txHash
}

func (tradingSuite *TradingTestSuite) callIssuingPRVReq(
	incTokenIDStr string,
	ethDepositProof []string,
	ethBlockHash string,
	ethTxIdx uint,
	methodName string,
) (map[string]interface{}, error) {
	rpcClient := rpccaller.NewRPCClient()
	meta := map[string]interface{}{
		"IncTokenID": incTokenIDStr,
		"BlockHash":  ethBlockHash,
		"ProofStrs":  ethDepositProof,
		"TxIndex":    ethTxIdx,
	}
	params := []interface{}{
		tradingSuite.IncPrivKeyStr,
		nil,
		5,
		-1,
		meta,
	}
	var res IssuingETHRes
	err := rpcClient.RPCCall(
		"",
		tradingSuite.IncRPCHost,
		"",
		methodName,
		params,
		&res,
	)
	if err != nil {
		return nil, err
	}

	response, _ := json.Marshal(res)
	fmt.Println("get response", string(response))

	if res.RPCError != nil {
		return nil, errors.New(res.RPCError.Message)
	}
	return res.Result.(map[string]interface{}), nil
}
