package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"time"

	"github.com/incognitochain/bridge-eth/bridge/vault"
	"github.com/incognitochain/bridge-eth/common/base58"
	"github.com/incognitochain/bridge-eth/consensus/signatureschemes/bridgesig"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/sha3"

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

type IssuingETHRes struct {
	rpccaller.RPCBaseRes
	Result interface{} `json:"Result"`
}

type BurningForDepositToSCRes struct {
	rpccaller.RPCBaseRes
	Result interface{} `json:"Result"`
}

type BalanceIncAccount struct {
	Result interface{} `json:"Result"`
}

type TxhashInC struct {
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
	IncSAITokenIDStr   string

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

	VaultAddr            common.Address
	KBNTradeDeployedAddr common.Address

	KyberContractAddr common.Address
	MKRAddressStr     string
	BATAddressStr     string
	KNCAddressStr     string
	CKNAddressStr     string

	IncMKRTokenIDStr string
	IncBATTokenIDStr string
	IncKNCTokenIDStr string
	IncCKNTokenIDStr string
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (tradingSuite *TradingTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")

	// 0x kovan env
	tradingSuite.IncBurningAddrStr = "12RxahVABnAVCGP3LGwCn8jkQxgw7z1x14wztHzn455TTVpi1wBq9YGwkRMQg3J4e657AbAnCvYCJSdA9czBUNuCKwGSRQt55Xwz8WA"
	tradingSuite.IncPrivKeyStr = "112t8rnX3VTd3MTWMpfbYP8HGY4ToAaLjrmUYzfjJBrAcb8iPLkNqvVDXWrLNiFV5yb2NBpR3FDZj3VW8GcLUwRdQ61hPMWP3YrREZAZ1UbH"
	tradingSuite.IncPaymentAddrStr = "12S6R8HfTyL74bggg47LX88RSvBPaMPBMEMoo6yx9WQ4EgLiYERXXcE2Mv2HrCsFuKhBsTfrYMeH82Bus5MHQGt3xHwoxX4v2qM5jRE"
	// tradingSuite.GeneratedPubKeyForSCStr = "8224890Cd5A537792d1B8B56c95FAb8a1A5E98B1"

	tradingSuite.IncEtherTokenIDStr = "ffd8d42dc40a8d166ea4848baf8b5f6e9fe0e9c30d60062eb7d44a8df9e00854"
	tradingSuite.IncDAITokenIDStr = "c7545459764224a000a9b323850648acf271186238210ce474b505cd17cc93a0"
	tradingSuite.IncSAITokenIDStr = "1d74e5e225e1f09ae38c496d3102aef464dcbd04ad3ac071e6e44077b8a740c9"
	tradingSuite.IncBATTokenIDStr = "d1b4c73821edc76963fdeda2236fe89478249a1f7b952de2a7135c0bc0cbe6dc"
	tradingSuite.IncKNCTokenIDStr = "d6644f62d0ef0475335ae7bb6103f358979cbfcd2b85481e3bde2b82101a095c"
	tradingSuite.IncCKNTokenIDStr = "f3c421e4d7520936f3916a878ab361ef3fd6a831e81063ca3e7b80ab4d15a84e"

	tradingSuite.BATAddressStr = "0x9f8cfb61d3b2af62864408dd703f9c3beb55dff7" //BAT
	tradingSuite.KNCAddressStr = "0xad67cb4d63c9da94aca37fdf2761aadf780ff4a2" //KNC
	tradingSuite.CKNAddressStr = "0x1d59Ee76304338fAc3a0eA9AE06E618C760D6042" //CKN

	tradingSuite.EtherAddressStr = "0x0000000000000000000000000000000000000000"
	tradingSuite.DAIAddressStr = "0x4f96fe3b7a6cf9725f59d353f723c1bdb64ca6aa" // DAI
	tradingSuite.SAIAddressStr = "0xc4375b7de8af5a38a93548eb8453a498222c4ff2" //SAI

	tradingSuite.ETHPrivKeyStr = "A5AE26C7154410DF235BC8669FFD27C0FC9D3068C21E469A4CC68165C68CD5CB"
	tradingSuite.ETHOwnerAddrStr = "cE40cE511A5D084017DBee7e3fF3e455ea32D85c"

	tradingSuite.ETHHost = "https://kovan.infura.io/v3/93fe721349134964aa71071a713c5cef"
	tradingSuite.IncBridgeHost = "https://test-node.incognito.org"
	tradingSuite.IncRPCHost = "https://test-node.incognito.org"
	//tradingSuite.IncBridgeHost = "http://68.183.187.162:29334"
	//tradingSuite.IncRPCHost = "http://68.183.187.162:29334"
	tradingSuite.VaultAddr = common.HexToAddress("0x31F7293dEebCEd75d035De0843498D87B90a3eee")
	//tradingSuite.VaultAddr = common.HexToAddress("0xe77aBF10cC0c30Ab3Ac2d877add39553cA7a8654")

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
) *big.Int {
	if tokenAddr.Hex() == tradingSuite.EtherAddressStr {
		balance, err := tradingSuite.ETHClient.BalanceAt(context.Background(), ownerAddr, nil)
		require.Equal(tradingSuite.T(), nil, err)
		return balance
	}
	// erc20 token
	instance, err := erc20.NewErc20(tokenAddr, tradingSuite.ETHClient)
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
	tradingSuite.ETHPrivKey = privKey
}

func (tradingSuite *TradingTestSuite) depositETH(
	amt float64,
	incPaymentAddrStr string,
) common.Hash {
	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)

	auth := bind.NewKeyedTransactor(tradingSuite.ETHPrivKey)
	auth.GasPrice = big.NewInt(50000000000)
	auth.Value = big.NewInt(int64(amt * params.Ether))
	tx, err := c.Deposit(auth, incPaymentAddrStr)
	require.Equal(tradingSuite.T(), nil, err)
	txHash := tx.Hash()

	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("Deposited to proof, txHash : %x\n", txHash[:])
	return txHash

}

func (tradingSuite *TradingTestSuite) depositERC20ToBridge(
	amt *big.Int,
	tokenAddr common.Address,
	incPaymentAddrStr string,
) common.Hash {
	auth := bind.NewKeyedTransactor(tradingSuite.ETHPrivKey)
	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)

	erc20Token, _ := erc20.NewErc20(tokenAddr, tradingSuite.ETHClient)
	auth.GasPrice = big.NewInt(50000000000)
	auth.GasLimit = 1000000
	tx2, apprErr := erc20Token.Approve(auth, tradingSuite.VaultAddr, amt)
	tx2Hash := tx2.Hash()
	fmt.Printf("Approve tx, txHash: %x\n", tx2Hash[:])
	require.Equal(tradingSuite.T(), nil, apprErr)
	time.Sleep(15 * time.Second)
	auth.GasPrice = big.NewInt(50000000000)
	auth.GasLimit = 1000000
	fmt.Println("Starting deposit erc20 to vault contract")
	tx, err := c.DepositERC20(auth, tokenAddr, amt, incPaymentAddrStr)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("Finished deposit erc20 to vault contract")
	txHash := tx.Hash()

	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
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
		"createandsendtxwithissuingethreq",
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
		5,
		-1,
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
) common.Hash {
	proof, err := getAndDecodeBurnProofV2(tradingSuite.IncBridgeHost, burningTxIDStr, "getburnprooffordeposittosc")
	require.Equal(tradingSuite.T(), nil, err)

	// Get contract instance
	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)

	// Burn
	auth := bind.NewKeyedTransactor(tradingSuite.ETHPrivKey)
	// auth.GasPrice = big.NewInt(1000000)
	auth.GasLimit = 2000000
	auth.GasPrice = big.NewInt(50000000000)
	tx, err := SubmitBurnProof(c, auth, proof)
	require.Equal(tradingSuite.T(), nil, err)

	txHash := tx.Hash()
	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("burned, txHash: %x\n", txHash[:])
	return txHash
}

func (tradingSuite *TradingTestSuite) submitBurnProofForWithdrawal(
	burningTxIDStr string,
) common.Hash {
	proof, err := getAndDecodeBurnProofV2(tradingSuite.IncBridgeHost, burningTxIDStr, "getburnproof")
	require.Equal(tradingSuite.T(), nil, err)

	// Get contract instance
	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)

	// Burn
	auth := bind.NewKeyedTransactor(tradingSuite.ETHPrivKey)
	// auth.GasPrice = big.NewInt(1000000)
	auth.GasLimit = 2000000
	auth.GasPrice = big.NewInt(50000000000)

	tx, err := Withdraw(c, auth, proof)
	require.Equal(tradingSuite.T(), nil, err)

	txHash := tx.Hash()
	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("burned, txHash: %x\n", txHash[:])
	return txHash
}

func (tradingSuite *TradingTestSuite) genKeysPairForSC() {
	incPriKeyBytes, _, err := base58.Base58Check{}.Decode(tradingSuite.IncPrivKeyStr)
	require.Equal(tradingSuite.T(), nil, err)

	tradingSuite.GeneratedPrivKeyForSC, tradingSuite.GeneratedPubKeyForSC = bridgesig.KeyGen(incPriKeyBytes)
}

// func  testgenKeysPairForSC(priKey string) (ecdsa.PublicKey){
// 	incPriKeyBytes, _, _ := base58.Base58Check{}.Decode(priKey)
// 	//require.Equal(tradingSuite.T(), nil, err)

// 	_, pubKey := bridgesig.KeyGen(incPriKeyBytes)

// 	return pubKey

// }

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
	//fmt.Printf("deposited balance on SC: %d\n", bal)
	return bal
}

func (tradingSuite *TradingTestSuite) requestWithdraw(
	withdrawalETHTokenIDStr string,
	amount *big.Int,
) common.Hash {
	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)
	auth := bind.NewKeyedTransactor(tradingSuite.ETHPrivKey)

	token := common.HexToAddress(withdrawalETHTokenIDStr)
	// amount := big.NewInt(0.1 * params.Ether)
	timestamp := []byte(randomizeTimestamp())
	tempData := append([]byte(tradingSuite.IncPaymentAddrStr), token[:]...)
	tempData1 := append(tempData, timestamp...)
	tempData2 := append(tempData1, common.LeftPadBytes(amount.Bytes(), 32)...)
	data := rawsha3(tempData2)
	signBytes, _ := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)
	auth.GasPrice = big.NewInt(50000000000)
	//auth.GasLimit = 10000000
	auth.GasLimit = 2000000

	tx, err := c.RequestWithdraw(auth, tradingSuite.IncPaymentAddrStr, token, amount, signBytes, timestamp)
	require.Equal(tradingSuite.T(), nil, err)

	txHash := tx.Hash()
	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("request withdrawal, txHash: %x\n", txHash[:])
	return txHash
}

func (tradingSuite *TradingTestSuite) getBalanceTokenIncAccount(
	IncPrivKeyStr string,
	ethTokenAddrStr string,
) (uint64, error) {
	rpcClient := rpccaller.NewRPCClient()
	params := []interface{}{
		IncPrivKeyStr,
		ethTokenAddrStr,
	}
	var res BalanceIncAccount
	err := rpcClient.RPCCall(
		"",
		tradingSuite.IncRPCHost,
		"",
		"getbalanceprivacycustomtoken",
		params,
		&res,
	)
	if err != nil {
		return 0, err
	}
	//fmt.Println(res.Result)
	return uint64(res.Result.(float64)), nil
}

func (tradingSuite *TradingTestSuite) getBalancePrvIncAccount(
	IncPrivKeyStr string,
) (uint64, error) {
	rpcClient := rpccaller.NewRPCClient()
	params := []interface{}{
		IncPrivKeyStr,
	}
	var res BalanceIncAccount
	err := rpcClient.RPCCall(
		"",
		tradingSuite.IncRPCHost,
		"",
		"getbalancebyprivatekey",
		params,
		&res,
	)
	if err != nil {
		return 0, err
	}
	//fmt.Println(res.Result)
	return uint64(res.Result.(float64)), nil
}

func (tradingSuite *TradingTestSuite) getFeePRVbyTxhashInC(
	txhash string,
) uint64 {
	rpcClient := rpccaller.NewRPCClient()
	params := []interface{}{
		txhash,
	}
	var res TxhashInC
	err := rpcClient.RPCCall(
		"",
		tradingSuite.IncRPCHost,
		"",
		"gettransactionbyhash",
		params,
		&res,
	)
	if err != nil {
		return 0
	}
	return uint64(res.Result.(map[string]interface{})["Fee"].(float64))
}

func (tradingSuite *TradingTestSuite) getGasFeeETHbyTxhash(txHash common.Hash) *big.Int {
	reciept, _ := getETHTransactionReceipt(tradingSuite.ETHHost, txHash)
	//fmt.Println("reciept gas fee : ", reciept.GasUsed)
	gasPrice := big.NewInt(50000000000)
	gasFee := big.NewInt(0).Div(big.NewInt(0).Mul(big.NewInt(int64(reciept.GasUsed)), gasPrice), big.NewInt(int64(1)))
	return gasFee

}
