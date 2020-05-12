package main

// import (
// 	"context"
// 	"crypto/ecdsa"
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"math/big"

// 	"time"

// 	"github.com/incognitochain/bridge-eth/bridge/vault"
// 	"github.com/incognitochain/bridge-eth/common/base58"
// 	"github.com/incognitochain/bridge-eth/consensus/signatureschemes/bridgesig"
// 	"github.com/stretchr/testify/suite"

// 	"github.com/ethereum/go-ethereum/accounts/abi/bind"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/crypto"
// 	"github.com/ethereum/go-ethereum/ethclient"
// 	"github.com/ethereum/go-ethereum/params"

// 	"github.com/incognitochain/bridge-eth/erc20"
// 	"github.com/incognitochain/bridge-eth/rpccaller"
// 	"github.com/stretchr/testify/require"
// )

// type ZeroXIssuingETHRes struct {
// 	rpccaller.RPCBaseRes
// 	Result interface{} `json:"Result"`
// }

// type ZeroXBurningForDepositToSCRes struct {
// 	rpccaller.RPCBaseRes
// 	Result interface{} `json:"Result"`
// }

// type ZeroXBalanceIncAccount struct {
// 	Result interface{} `json:"Result"`
// }

// type ZeroXTxhashInC struct {
// 	Result interface{} `json:"Result"`
// }

// // Define the suite, and absorb the built-in basic suite
// // functionality from testify - including assertion methods.
// type ZeroXTradingTestSuite struct {
// 	suite.Suite
// 	IncBurningAddrStr string
// 	IncPrivKeyStr     string
// 	IncPaymentAddrStr string

// 	GeneratedPrivKeyForSC ecdsa.PrivateKey
// 	GeneratedPubKeyForSC  ecdsa.PublicKey

// 	IncEtherTokenIDStr string
// 	IncDAITokenIDStr   string
// 	IncSAITokenIDStr   string

// 	IncBridgeHost string
// 	IncRPCHost    string

// 	EtherAddressStr string
// 	DAIAddressStr   string
// 	SAIAddressStr   string

// 	ETHPrivKeyStr   string
// 	ETHOwnerAddrStr string

// 	ETHPrivKeyStrKyber   string
// 	ETHOwnerAddrStrKyber string

// 	ETHHost    string
// 	ETHPrivKey *ecdsa.PrivateKey
// 	ETHClient  *ethclient.Client

// 	VaultAddr            common.Address
// 	KBNTradeDeployedAddr common.Address

// 	KyberContractAddr common.Address
// 	MKRAddressStr     string
// 	BATAddressStr     string
// 	KNCAddressStr     string

// 	IncMKRTokenIDStr string
// 	IncBATTokenIDStr string
// 	IncKNCTokenIDStr string
// }

// // Make sure that VariableThatShouldStartAtFive is set to five
// // before each test
// func (tradingSuite *ZeroXTradingTestSuite) SetupSuite() {
// 	fmt.Println("Setting up the suite...")

// 	// 0x kovan env
// 	tradingSuite.IncBurningAddrStr = "15pABFiJVeh9D5uiQEhQX4SVibGGbdAVipQxBdxkmDqAJaoG1EdFKHBrNfs"
// 	tradingSuite.IncPrivKeyStr = "112t8roafGgHL1rhAP9632Yef3sx5k8xgp8cwK4MCJsCL1UWcxXvpzg97N4dwvcD735iKf31Q2ZgrAvKfVjeSUEvnzKJyyJD3GqqSZdxN4or"
// 	tradingSuite.IncPaymentAddrStr = "12S5Lrs1XeQLbqN4ySyKtjAjd2d7sBP2tjFijzmp6avrrkQCNFMpkXm3FPzj2Wcu2ZNqJEmh9JriVuRErVwhuQnLmWSaggobEWsBEci"
// 	// tradingSuite.GeneratedPubKeyForSCStr = "8224890Cd5A537792d1B8B56c95FAb8a1A5E98B1"

// 	tradingSuite.IncEtherTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000097"
// 	tradingSuite.IncDAITokenIDStr = "0000000000000000000000000000000000000000000000000000000000000098"
// 	tradingSuite.IncSAITokenIDStr = "0000000000000000000000000000000000000000000000000000000000000099"
// 	tradingSuite.IncMKRTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000091"
// 	tradingSuite.IncBATTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000092"
// 	tradingSuite.IncKNCTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000094"

// 	tradingSuite.MKRAddressStr = "0xaaf64bfcc32d0f15873a02163e7e500671a4ffcd" // MKR
// 	tradingSuite.BATAddressStr = "0x9f8cfb61d3b2af62864408dd703f9c3beb55dff7" //BAT
// 	tradingSuite.KNCAddressStr = "0xad67cb4d63c9da94aca37fdf2761aadf780ff4a2" //KNC

// 	tradingSuite.EtherAddressStr = "0x0000000000000000000000000000000000000000"
// 	tradingSuite.DAIAddressStr = "0x4f96fe3b7a6cf9725f59d353f723c1bdb64ca6aa" // DAI
// 	tradingSuite.SAIAddressStr = "0xc4375b7de8af5a38a93548eb8453a498222c4ff2" //SAI

// 	tradingSuite.ETHPrivKeyStr = "A5AE26C7154410DF235BC8669FFD27C0FC9D3068C21E469A4CC68165C68CD5CB"
// 	tradingSuite.ETHOwnerAddrStr = "cE40cE511A5D084017DBee7e3fF3e455ea32D85c"

// 	tradingSuite.ETHPrivKeyStrKyber = "81C85096BC78372F258C804ADFF8CC0F16F477CC707C366DDA02F4A50DD4FE3E"
// 	tradingSuite.ETHOwnerAddrStrKyber = "7652a1D59E88159cf286A578B635B9237C84190c"

// 	tradingSuite.ETHHost = "https://kovan.infura.io/v3/93fe721349134964aa71071a713c5cef"
// 	tradingSuite.IncBridgeHost = "http://127.0.0.1:9338"
// 	tradingSuite.IncRPCHost = "http://127.0.0.1:9334"
// 	tradingSuite.VaultAddr = common.HexToAddress("0x220b909142f6173E6Be19Fa14F69F9CB447ccA84")
// 	//tradingSuite.VaultAddr = common.HexToAddress("0x37AC6f8aADb28f09D74D9eDE87fF1d3C94E4D59F")

// 	// generate a new keys pair for SC
// 	tradingSuite.genKeysPairForSC()

// 	// connect to ethereum network
// 	tradingSuite.connectToETH()
// }

// func (tradingSuite *ZeroXTradingTestSuite) TearDownSuite() {
// 	fmt.Println("Tearing down the suite...")
// 	tradingSuite.ETHClient.Close()
// }

// func (tradingSuite *ZeroXTradingTestSuite) SetupTest() {
// 	fmt.Println("Setting up the test...")
// }

// func (tradingSuite *ZeroXTradingTestSuite) TearDownTest() {
// 	fmt.Println("Tearing down the test...")
// }

// func (tradingSuite *ZeroXTradingTestSuite) TestTradingTestSuite() {
// 	fmt.Println("This is generic test suite")
// }

// func (tradingSuite *ZeroXTradingTestSuite) getBalanceOnETHNet(
// 	tokenAddr common.Address,
// 	ownerAddr common.Address,
// ) *big.Int {
// 	if tokenAddr.Hex() == tradingSuite.EtherAddressStr {
// 		balance, err := tradingSuite.ETHClient.BalanceAt(context.Background(), ownerAddr, nil)
// 		require.Equal(tradingSuite.T(), nil, err)
// 		return balance
// 	}
// 	// erc20 token
// 	instance, err := erc20.NewErc20(tokenAddr, tradingSuite.ETHClient)
// 	require.Equal(tradingSuite.T(), nil, err)

// 	balance, err := instance.BalanceOf(&bind.CallOpts{}, ownerAddr)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	return balance
// }

// func (tradingSuite *ZeroXTradingTestSuite) connectToETH() {
// 	privKeyHex := tradingSuite.ETHPrivKeyStr
// 	privKey, err := crypto.HexToECDSA(privKeyHex)
// 	require.Equal(tradingSuite.T(), nil, err)

// 	fmt.Printf("Sign Txs with address: %s\n", crypto.PubkeyToAddress(privKey.PublicKey).Hex())

// 	network := "development"
// 	fmt.Printf("Connecting to network %s\n", network)
// 	client, err := ethclient.Dial(tradingSuite.ETHHost)
// 	require.Equal(tradingSuite.T(), nil, err)

// 	tradingSuite.ETHClient = client
// 	tradingSuite.ETHPrivKey = privKey
// }

// func (tradingSuite *ZeroXTradingTestSuite) depositETH(
// 	amt float64,
// 	incPaymentAddrStr string,
// ) common.Hash {
// 	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
// 	require.Equal(tradingSuite.T(), nil, err)

// 	auth := bind.NewKeyedTransactor(tradingSuite.ETHPrivKey)
// 	auth.GasPrice = big.NewInt(50000000000)
// 	auth.Value = big.NewInt(int64(amt * params.Ether))
// 	tx, err := c.Deposit(auth, incPaymentAddrStr)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	txHash := tx.Hash()

// 	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
// 		require.Equal(tradingSuite.T(), nil, err)
// 	}
// 	fmt.Printf("Deposited to proof, txHash : %x\n", txHash[:])
// 	return txHash

// }

// func (tradingSuite *ZeroXTradingTestSuite) depositERC20ToBridge(
// 	amt *big.Int,
// 	tokenAddr common.Address,
// 	incPaymentAddrStr string,
// ) common.Hash {
// 	auth := bind.NewKeyedTransactor(tradingSuite.ETHPrivKey)
// 	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
// 	require.Equal(tradingSuite.T(), nil, err)

// 	erc20Token, _ := erc20.NewErc20(tokenAddr, tradingSuite.ETHClient)
// 	auth.GasPrice = big.NewInt(50000000000)
// 	auth.GasLimit = 1000000
// 	tx2, apprErr := erc20Token.Approve(auth, tradingSuite.VaultAddr, amt)
// 	tx2Hash := tx2.Hash()
// 	fmt.Printf("Approve tx, txHash: %x\n", tx2Hash[:])
// 	require.Equal(tradingSuite.T(), nil, apprErr)
// 	time.Sleep(15 * time.Second)
// 	auth.GasPrice = big.NewInt(50000000000)
// 	auth.GasLimit = 1000000
// 	fmt.Println("Starting deposit erc20 to vault contract")
// 	tx, err := c.DepositERC20(auth, tokenAddr, amt, incPaymentAddrStr)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	fmt.Println("Finished deposit erc20 to vault contract")
// 	txHash := tx.Hash()

// 	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
// 		require.Equal(tradingSuite.T(), nil, err)
// 	}
// 	fmt.Printf("deposited erc20 token to bridge, txHash: %x\n", txHash[:])
// 	return txHash
// }

// func (tradingSuite *ZeroXTradingTestSuite) callIssuingETHReq(
// 	incTokenIDStr string,
// 	ethDepositProof []string,
// 	ethBlockHash string,
// 	ethTxIdx uint,
// ) (map[string]interface{}, error) {
// 	rpcClient := rpccaller.NewRPCClient()
// 	meta := map[string]interface{}{
// 		"IncTokenID": incTokenIDStr,
// 		"BlockHash":  ethBlockHash,
// 		"ProofStrs":  ethDepositProof,
// 		"TxIndex":    ethTxIdx,
// 	}
// 	params := []interface{}{
// 		tradingSuite.IncPrivKeyStr,
// 		nil,
// 		5,
// 		-1,
// 		meta,
// 	}
// 	var res ZeroXIssuingETHRes
// 	err := rpcClient.RPCCall(
// 		"",
// 		tradingSuite.IncRPCHost,
// 		"",
// 		"createandsendtxwithissuingethreq",
// 		params,
// 		&res,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

// 	response, _ := json.Marshal(res)
// 	fmt.Println("get response", string(response))

// 	if res.RPCError != nil {
// 		return nil, errors.New(res.RPCError.Message)
// 	}
// 	return res.Result.(map[string]interface{}), nil
// }

// func (tradingSuite *ZeroXTradingTestSuite) callBurningPToken(
// 	incTokenIDStr string,
// 	amount *big.Int,
// 	remoteAddrStr string,
// 	burningMethod string,
// ) (map[string]interface{}, error) {
// 	rpcClient := rpccaller.NewRPCClient()
// 	meta := map[string]interface{}{
// 		"TokenID":     incTokenIDStr,
// 		"TokenTxType": 1,
// 		"TokenName":   "",
// 		"TokenSymbol": "",
// 		"TokenAmount": amount.Uint64(),
// 		"TokenReceivers": map[string]uint64{
// 			tradingSuite.IncBurningAddrStr: amount.Uint64(),
// 		},
// 		"RemoteAddress": remoteAddrStr,
// 		"Privacy":       true,
// 		"TokenFee":      0,
// 	}
// 	params := []interface{}{
// 		tradingSuite.IncPrivKeyStr,
// 		nil,
// 		5,
// 		-1,
// 		meta,
// 		"",
// 		0,
// 	}
// 	var res ZeroXBurningForDepositToSCRes
// 	err := rpcClient.RPCCall(
// 		"",
// 		tradingSuite.IncRPCHost,
// 		"",
// 		burningMethod,
// 		params,
// 		&res,
// 	)
// 	if err != nil {
// 		fmt.Println("calling burning ptokens err: ", err)
// 		return nil, err
// 	}
// 	bb, _ := json.Marshal(res)
// 	fmt.Println("calling burning ptokens res: ", string(bb))
// 	if res.RPCError != nil {
// 		return nil, errors.New(res.RPCError.Message)
// 	}
// 	return res.Result.(map[string]interface{}), nil
// }

// func (tradingSuite *ZeroXTradingTestSuite) submitBurnProofForDepositToSC(
// 	burningTxIDStr string,
// ) common.Hash {
// 	proof, err := getAndDecodeBurnProofV2(tradingSuite.IncBridgeHost, burningTxIDStr, "getburnprooffordeposittosc")
// 	require.Equal(tradingSuite.T(), nil, err)

// 	// Get contract instance
// 	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
// 	require.Equal(tradingSuite.T(), nil, err)

// 	// Burn
// 	auth := bind.NewKeyedTransactor(tradingSuite.ETHPrivKey)
// 	// auth.GasPrice = big.NewInt(1000000)
// 	auth.GasLimit = 10000000
// 	auth.GasPrice = big.NewInt(50000000000)
// 	tx, err := SubmitBurnProof(c, auth, proof)
// 	require.Equal(tradingSuite.T(), nil, err)

// 	txHash := tx.Hash()
// 	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
// 		require.Equal(tradingSuite.T(), nil, err)
// 	}
// 	fmt.Printf("burned, txHash: %x\n", txHash[:])
// 	return txHash
// }

// func (tradingSuite *ZeroXTradingTestSuite) submitBurnProofForWithdrawal(
// 	burningTxIDStr string,
// ) common.Hash {
// 	proof, err := getAndDecodeBurnProofV2(tradingSuite.IncBridgeHost, burningTxIDStr, "getburnproof")
// 	require.Equal(tradingSuite.T(), nil, err)

// 	// Get contract instance
// 	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
// 	require.Equal(tradingSuite.T(), nil, err)

// 	// Burn
// 	auth := bind.NewKeyedTransactor(tradingSuite.ETHPrivKey)
// 	// auth.GasPrice = big.NewInt(1000000)
// 	// auth.GasLimit = 4000000
// 	auth.GasPrice = big.NewInt(50000000000)
// 	tx, err := Withdraw(c, auth, proof)
// 	require.Equal(tradingSuite.T(), nil, err)

// 	txHash := tx.Hash()
// 	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
// 		require.Equal(tradingSuite.T(), nil, err)
// 	}
// 	fmt.Printf("burned, txHash: %x\n", txHash[:])
// 	return txHash
// }

// func (tradingSuite *ZeroXTradingTestSuite) genKeysPairForSC() {
// 	incPriKeyBytes, _, err := base58.Base58Check{}.Decode(tradingSuite.IncPrivKeyStr)
// 	require.Equal(tradingSuite.T(), nil, err)

// 	tradingSuite.GeneratedPrivKeyForSC, tradingSuite.GeneratedPubKeyForSC = bridgesig.KeyGen(incPriKeyBytes)
// }

// func (tradingSuite *ZeroXTradingTestSuite) getDepositedBalance(
// 	ethTokenAddrStr string,
// 	ownerAddrStr string,
// ) *big.Int {
// 	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	token := common.HexToAddress(ethTokenAddrStr)
// 	owner := common.HexToAddress(ownerAddrStr)
// 	bal, err := c.GetDepositedBalance(nil, token, owner)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	//fmt.Printf("deposited balance on SC: %d\n", bal)
// 	return bal
// }

// func (tradingSuite *ZeroXTradingTestSuite) requestWithdraw(
// 	withdrawalETHTokenIDStr string,
// 	amount *big.Int,
// ) common.Hash {
// 	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	auth := bind.NewKeyedTransactor(tradingSuite.ETHPrivKey)

// 	token := common.HexToAddress(withdrawalETHTokenIDStr)
// 	// amount := big.NewInt(0.1 * params.Ether)
// 	timestamp := []byte(randomizeTimestamp())
// 	tempData := append([]byte(tradingSuite.IncPaymentAddrStr), token[:]...)
// 	tempData1 := append(tempData, timestamp...)
// 	data := rawsha3(tempData1)
// 	signBytes, _ := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)
// 	auth.GasPrice = big.NewInt(50000000000)
// 	auth.GasLimit = 10000000
// 	//auth.GasLimit = 2000000

// 	tx, err := c.RequestWithdraw(auth, tradingSuite.IncPaymentAddrStr, token, amount, signBytes, timestamp)
// 	require.Equal(tradingSuite.T(), nil, err)

// 	txHash := tx.Hash()
// 	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
// 		require.Equal(tradingSuite.T(), nil, err)
// 	}
// 	fmt.Printf("request withdrawal, txHash: %x\n", txHash[:])
// 	return txHash
// }

// func (tradingSuite *ZeroXTradingTestSuite) getBalanceTokenIncAccount(
// 	IncPrivKeyStr string,
// 	ethTokenAddrStr string,
// ) (uint64, error) {
// 	rpcClient := rpccaller.NewRPCClient()
// 	params := []interface{}{
// 		IncPrivKeyStr,
// 		ethTokenAddrStr,
// 	}
// 	var res ZeroXBalanceIncAccount
// 	err := rpcClient.RPCCall(
// 		"",
// 		tradingSuite.IncRPCHost,
// 		"",
// 		"getbalanceprivacycustomtoken",
// 		params,
// 		&res,
// 	)
// 	if err != nil {
// 		return 0, err
// 	}
// 	//fmt.Println(res.Result)
// 	return uint64(res.Result.(float64)), nil
// }

// func (tradingSuite *ZeroXTradingTestSuite) getBalancePrvIncAccount(
// 	IncPrivKeyStr string,
// ) (uint64, error) {
// 	rpcClient := rpccaller.NewRPCClient()
// 	params := []interface{}{
// 		IncPrivKeyStr,
// 	}
// 	var res ZeroXBalanceIncAccount
// 	err := rpcClient.RPCCall(
// 		"",
// 		tradingSuite.IncRPCHost,
// 		"",
// 		"getbalancebyprivatekey",
// 		params,
// 		&res,
// 	)
// 	if err != nil {
// 		return 0, err
// 	}
// 	//fmt.Println(res.Result)
// 	return uint64(res.Result.(float64)), nil
// }

// func (tradingSuite *ZeroXTradingTestSuite) getFeePRVbyTxhashInC(
// 	txhash string,
// ) uint64 {
// 	rpcClient := rpccaller.NewRPCClient()
// 	params := []interface{}{
// 		txhash,
// 	}
// 	var res ZeroXTxhashInC
// 	err := rpcClient.RPCCall(
// 		"",
// 		tradingSuite.IncRPCHost,
// 		"",
// 		"gettransactionbyhash",
// 		params,
// 		&res,
// 	)
// 	if err != nil {
// 		return 0
// 	}
// 	return uint64(res.Result.(map[string]interface{})["Fee"].(float64))
// }

// func (tradingSuite *ZeroXTradingTestSuite) getGasFeeETHbyTxhash(txHash common.Hash) *big.Int {
// 	reciept, _ := getETHTransactionReceipt(tradingSuite.ETHHost, txHash)
// 	//fmt.Println("reciept gas fee : ", reciept.GasUsed)
// 	gasPrice := big.NewInt(50000000000)
// 	gasFee := big.NewInt(0).Div(big.NewInt(0).Mul(big.NewInt(int64(reciept.GasUsed)), gasPrice), big.NewInt(int64(1)))
// 	return gasFee

// }
