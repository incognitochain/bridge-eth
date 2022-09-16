package bridge

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/incognitochain/bridge-eth/bridge/dapp"
	"github.com/incognitochain/bridge-eth/bridge/vaultproxy"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	ec "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/incognitochain/bridge-eth/bridge/vault"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// // Define the suite, and absorb the built-in basic suite
// // functionality from testify - including assertion methods.
type VaulV2TestSuite struct {
	suite.Suite
	p            *Platform
	c            *committees
	v            *vault.Vault
	withdrawer   common.Address
	auth         *bind.TransactOpts
	EtherAddress common.Address
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (v2 *VaulV2TestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")
	v2.withdrawer = ec.HexToAddress("0xe722D8b71DCC0152D47D2438556a45D3357d631f")
	v2.EtherAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")
}

func (v2 *VaulV2TestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
}

func (v2 *VaulV2TestSuite) SetupTest() {
	fmt.Println("Setting up the test...")
	p, c, err := setupFixedCommittee()
	require.Equal(v2.T(), nil, err)
	v2.p = p
	v2.c = c
	require.Equal(v2.T(), nil, err)
}

func (v2 *VaulV2TestSuite) TearDownTest() {
	fmt.Println("Tearing down the test...")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestVaultV2(t *testing.T) {
	fmt.Println("Starting entry point for vault v2 test suite...")
	suite.Run(t, new(VaulV2TestSuite))

	fmt.Println("Finishing entry point for vault v2 test suite...")
}

func (v2 *VaulV2TestSuite) TestVaultV2SubmitBurnProof() {
	desc := "DAI"
	deposit := big.NewInt(int64(3e17))
	withdraw := big.NewInt(int64(1e8))
	// Deposit, must success
	tinfo := v2.p.customErc20s[desc]
	_, _, err := lockSimERC20WithTxs(v2.p, tinfo.c, tinfo.addr, deposit)
	require.Equal(v2.T(), nil, err)

	// wrong meta
	meta := 98
	shardID := 1
	proof := buildWithdrawTestcase(v2.c, meta, shardID, tinfo.addr, withdraw)

	auth.GasLimit = 0
	_, err = SubmitBurnProof(v2.p.v, auth2, proof)
	require.NotEqual(v2.T(), nil, err)
	v2.p.sim.Commit()

	// wrong shard
	meta = 243
	shardID = 2
	proof = buildWithdrawTestcase(v2.c, meta, shardID, tinfo.addr, withdraw)

	auth.GasLimit = 0
	_, err = SubmitBurnProof(v2.p.v, auth2, proof)
	require.NotEqual(v2.T(), nil, err)
	v2.p.sim.Commit()

	// able to submit proof
	shardID = 1
	proof = buildWithdrawTestcase(v2.c, meta, shardID, tinfo.addr, withdraw)

	auth.GasLimit = 0
	_, err = SubmitBurnProof(v2.p.v, auth2, proof)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()

	// Check balance
	require.Equal(v2.T(), nil, err)
	bal, err := v2.p.v.GetDepositedBalance(nil, tinfo.addr, v2.withdrawer)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), bal, big.NewInt(0).Mul(withdraw, big.NewInt(int64(1e9))))

	// use proof twice
	_, err = SubmitBurnProof(v2.p.v, auth2, proof)
	require.NotEqual(v2.T(), nil, err)
	v2.p.sim.Commit()

	// submit wrong proof
	proof = buildWithdrawTestcase(v2.c, meta, shardID, tinfo.addr, withdraw)
	proof.Instruction[0]++
	auth.GasLimit = 0
	_, err = SubmitBurnProof(v2.p.v, auth2, proof)
	require.NotEqual(v2.T(), nil, err)
	v2.p.sim.Commit()

	// submit wrong inst bytes arrays
	proof = buildWithdrawTestcase(v2.c, meta, shardID, tinfo.addr, withdraw)
	fmt.Println(hex.EncodeToString(proof.Instruction))
	proof.Instruction = proof.Instruction[1:]
	auth.GasLimit = 0
	_, err = SubmitBurnProof(v2.p.v, auth2, proof)
	require.NotEqual(v2.T(), nil, err)
	v2.p.sim.Commit()

	// pause and submitBurnProof
	_, err = v2.p.vp.Pause(auth)
	require.Equal(v2.T(), nil, err)
	proof = buildWithdrawTestcase(v2.c, meta, shardID, tinfo.addr, withdraw)
	auth.GasLimit = 0
	_, err = SubmitBurnProof(v2.p.v, auth2, proof)
	require.NotEqual(v2.T(), nil, err)
	v2.p.sim.Commit()
}

func (v2 *VaulV2TestSuite) TestVaultV2RequestWithdraw() {
	desc := "BNB"
	deposit := big.NewInt(int64(3e17))
	withdraw := big.NewInt(int64(2e8))
	redeposit := big.NewInt(int64(1e8))
	address := crypto.PubkeyToAddress(genesisAcc.PrivateKey.PublicKey)
	// Deposit, must success
	tinfo := v2.p.customErc20s[desc]
	_, _, err := lockSimERC20WithTxs(v2.p, tinfo.c, tinfo.addr, deposit)
	require.Equal(v2.T(), nil, err)

	proof := buildWithdrawTestcaseV2(v2.c, 243, 1, tinfo.addr, withdraw, address)
	auth.GasLimit = 0
	_, err = SubmitBurnProof(v2.p.v, auth2, proof)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()

	// vaultAbi, _ := abi.JSON(strings.NewReader(vault.VaultABI))
	// request amount bigger than balance
	timestamp := []byte(randomizeTimestamp())
	vaultHelperAbi, _ := abi.JSON(strings.NewReader(vault.VaultHelperABI))
	psData := vault.VaultHelperPreSignData{
		Prefix:    REQ_WITHDRAW_PREFIX,
		Token:     tinfo.addr,
		Timestamp: timestamp,
		Amount:    deposit,
	}
	tempData, _ := vaultHelperAbi.Pack("_buildSignRequestWithdraw", psData, IncPaymentAddr)
	// tempData, _ := vaultAbi.Pack("withdrawBuildData", IncPaymentAddr, tinfo.addr, timestamp, deposit)
	data := rawsha3(tempData[4:])
	signBytes, _ := crypto.Sign(data, genesisAcc.PrivateKey)
	txId := [32]byte{}
	// todo: update genesisAcc2.PrivateKey to real regulator
	sign, err := SignDataToShield(txId, genesisAcc2.PrivateKey, auth2.From)
	require.Equal(v2.T(), nil, err)
	_, err = v2.p.v.RequestWithdraw(auth2, IncPaymentAddr, tinfo.addr, deposit, signBytes, timestamp, txId, sign)
	require.NotEqual(v2.T(), nil, err)
	v2.p.sim.Commit()

	// able to request withdraw
	psData = vault.VaultHelperPreSignData{
		Prefix:    REQ_WITHDRAW_PREFIX,
		Token:     tinfo.addr,
		Timestamp: timestamp,
		Amount:    big.NewInt(0).Mul(redeposit, big.NewInt(int64(1e9))),
	}
	tempData, _ = vaultHelperAbi.Pack("_buildSignRequestWithdraw", psData, IncPaymentAddr)
	// tempData, _ = vaultAbi.Pack("withdrawBuildData", IncPaymentAddr, tinfo.addr, timestamp, big.NewInt(0).Mul(redeposit, big.NewInt(int64(1e9))))
	data = rawsha3(tempData[4:])
	signBytes, _ = crypto.Sign(data, genesisAcc.PrivateKey)
	_, err = v2.p.v.RequestWithdraw(auth2, IncPaymentAddr, tinfo.addr, big.NewInt(0).Mul(redeposit, big.NewInt(int64(1e9))), signBytes, timestamp, txId, sign)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()

	// use signature twice
	_, err = v2.p.v.RequestWithdraw(auth2, IncPaymentAddr, tinfo.addr, big.NewInt(0).Mul(redeposit, big.NewInt(int64(1e9))), signBytes, timestamp, txId, sign)
	require.NotEqual(v2.T(), nil, err)
	v2.p.sim.Commit()

	// check balance remain
	bal, err := v2.p.v.GetDepositedBalance(nil, tinfo.addr, address)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), bal, big.NewInt(0).Mul(redeposit, big.NewInt(int64(1e9))))

	// amount decreased so can not request amount as amount at time withdraw from incognito
	timestamp = []byte(randomizeTimestamp())
	psData = vault.VaultHelperPreSignData{
		Prefix:    REQ_WITHDRAW_PREFIX,
		Token:     tinfo.addr,
		Timestamp: timestamp,
		Amount:    big.NewInt(0).Mul(withdraw, big.NewInt(int64(1e9))),
	}
	tempData, _ = vaultHelperAbi.Pack("_buildSignRequestWithdraw", psData, IncPaymentAddr)
	// tempData, _ = vaultAbi.Pack("withdrawBuildData", IncPaymentAddr, tinfo.addr, timestamp, big.NewInt(0).Mul(withdraw, big.NewInt(int64(1e9))))
	data = rawsha3(tempData[4:])
	signBytes, _ = crypto.Sign(data, genesisAcc.PrivateKey)
	_, err = v2.p.v.RequestWithdraw(auth2, IncPaymentAddr, tinfo.addr, big.NewInt(0).Mul(withdraw, big.NewInt(int64(1e9))), signBytes, timestamp, txId, sign)
	require.NotEqual(v2.T(), nil, err)
	v2.p.sim.Commit()

	// update delegator then user must be able to request withdraw
	nextVaultAddr, _, _, err := setupNextDelegator(auth, v2.p.sim)
	require.Equal(v2.T(), nil, err)
	_, err = v2.p.vp.Pause(auth)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()
	proxyVault, err := vaultproxy.NewTransparentUpgradeableProxy(v2.p.vAddr, v2.p.sim)
	_, err = proxyVault.UpgradeTo(auth, nextVaultAddr)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()
	_, err = v2.p.vp.Unpause(auth)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()
	// totalDeposit, err := v2.p.v.TotalDepositedToSCAmount(nil, tinfo.addr)
	// require.Equal(v2.T(), nil, err)
	// require.Equal(v2.T(), 0, totalDeposit.Cmp(big.NewInt(0)))
	totalDeposit, err := v2.p.v.TotalDepositedToSCAmount(nil, tinfo.addr)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), big.NewInt(0).Mul(redeposit, big.NewInt(int64(1e9))), totalDeposit)

	// check balance from nextVault
	bal, err = v2.p.v.GetDepositedBalance(nil, tinfo.addr, address)

	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), bal, big.NewInt(0).Mul(redeposit, big.NewInt(int64(1e9))))
	psData = vault.VaultHelperPreSignData{
		Prefix:    REQ_WITHDRAW_PREFIX,
		Token:     tinfo.addr,
		Timestamp: timestamp,
		Amount:    big.NewInt(0).Mul(redeposit, big.NewInt(int64(1e9))),
	}
	tempData, _ = vaultHelperAbi.Pack("_buildSignRequestWithdraw", psData, IncPaymentAddr)
	// tempData, _ = vaultAbi.Pack("withdrawBuildData", IncPaymentAddr, tinfo.addr, timestamp, big.NewInt(0).Mul(redeposit, big.NewInt(int64(1e9))))
	data = rawsha3(tempData[4:])
	signBytes, err = crypto.Sign(data, genesisAcc.PrivateKey)
	require.Equal(v2.T(), nil, err)
	_, err = v2.p.v.RequestWithdraw(auth2, IncPaymentAddr, tinfo.addr, big.NewInt(0).Mul(redeposit, big.NewInt(int64(1e9))), signBytes, timestamp, txId, sign)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()

	// check balance from nextVault after RequestWithdraw
	bal, err = v2.p.v.GetDepositedBalance(nil, tinfo.addr, address)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), 0, bal.Cmp(big.NewInt(0)))
}

func (v2 *VaulV2TestSuite) TestVaultV2ExecuteETHtoERC20() {
	desc := "BNB"
	tinfo := v2.p.customErc20s[desc]
	withdraw := big.NewInt(int64(5e18))
	executeAmount := big.NewInt(int64(1e18))
	srcToken := v2.EtherAddress
	destToken := tinfo.addr
	dAddr, _, _, err := dapp.DeployDapp(auth, v2.p.sim, v2.p.vAddr)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()
	_, _, err = deposit(v2.p, big.NewInt(int64(5e18)))
	require.Equal(v2.T(), nil, err)
	proof := buildWithdrawTestcaseV2(v2.c, 243, 1, v2.EtherAddress, withdraw, auth2.From)
	auth2.GasLimit = 0
	_, err = SubmitBurnProof(v2.p.v, auth2, proof)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()
	dappAbi, err := abi.JSON(strings.NewReader(dapp.DappABI))

	// check balance remain
	bal, err := v2.p.v.GetDepositedBalance(nil, v2.EtherAddress, auth2.From)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), withdraw, bal)
	timestamp := []byte(randomizeTimestamp())

	testCases := []struct {
		srcToken  common.Address
		srcQty    *big.Int
		destToken common.Address
		dapp      common.Address
		input     []byte
		vault     *vault.Vault
		timestamp []byte
		iserr     bool
	}{
		{
			// must be able to run execute smoothly by calling simpleCall func
			srcToken:  srcToken,
			srcQty:    executeAmount,
			destToken: destToken,
			dapp:      dAddr,
			input:     v2.packInputData(dappAbi, "simpleCall", destToken),
			vault:     v2.p.v,
			timestamp: timestamp,
			iserr:     false,
		},
		{
			// use sig twice
			srcToken:  srcToken,
			srcQty:    executeAmount,
			destToken: destToken,
			dapp:      dAddr,
			input:     v2.packInputData(dappAbi, "simpleCall", destToken),
			vault:     v2.p.v,
			timestamp: timestamp,
			iserr:     true,
		},
		{
			// trade same token
			srcToken:  srcToken,
			srcQty:    executeAmount,
			destToken: srcToken,
			dapp:      dAddr,
			input:     v2.packInputData(dappAbi, "simpleCall", destToken),
			vault:     v2.p.v,
			timestamp: []byte(randomizeTimestamp()),
			iserr:     true,
		},
		{
			// trade amount greater than amount available
			srcToken:  srcToken,
			srcQty:    withdraw,
			destToken: destToken,
			dapp:      dAddr,
			input:     v2.packInputData(dappAbi, "simpleCall", destToken),
			vault:     v2.p.v,
			timestamp: []byte(randomizeTimestamp()),
			iserr:     true,
		},
		{
			// execute dapp get revert
			srcToken:  srcToken,
			srcQty:    executeAmount,
			destToken: destToken,
			dapp:      dAddr,
			input:     v2.packInputData(dappAbi, "revertCall", destToken),
			vault:     v2.p.v,
			timestamp: []byte(randomizeTimestamp()),
			iserr:     true,
		},
		{
			// dapp return value amount but didn't transfer dest coin
			srcToken:  srcToken,
			srcQty:    executeAmount,
			destToken: destToken,
			dapp:      dAddr,
			input:     v2.packInputData(dappAbi, "ReturnAmountWithoutTranfer", destToken),
			vault:     v2.p.v,
			timestamp: []byte(randomizeTimestamp()),
			iserr:     true,
		},
	}

	for _, tc := range testCases {
		beforeExecute := v2.p.getBalance(v2.p.vAddr)
		_, err = runExecuteVault(auth2, tc.dapp, tc.srcToken, tc.srcQty, tc.destToken, tc.input, tc.vault, tc.timestamp, genesisAcc2.PrivateKey)
		if tc.iserr {
			require.NotEqual(v2.T(), nil, err)
			v2.p.sim.Commit()
			require.Equal(v2.T(), beforeExecute, v2.p.getBalance(v2.p.vAddr))
		} else {
			require.Equal(v2.T(), nil, err)
			v2.p.sim.Commit()
			require.NotEqual(v2.T(), beforeExecute, v2.p.getBalance(v2.p.vAddr))
		}
	}
}

func (v2 *VaulV2TestSuite) TestVaultV2ExecuteERC20toETH() {
	desc := "USDT"
	tinfo := v2.p.customErc20s[desc]
	withdraw := big.NewInt(int64(5e8))
	executeAmount := big.NewInt(int64(1e8))
	srcToken := tinfo.addr
	destToken := v2.EtherAddress
	address := crypto.PubkeyToAddress(genesisAcc.PrivateKey.PublicKey)
	dAddr, _, _, err := dapp.DeployDapp(auth, v2.p.sim, v2.p.vAddr)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()
	_, _, err = lockSimERC20WithTxs(v2.p, tinfo.c, tinfo.addr, withdraw)
	require.Equal(v2.T(), nil, err)
	proof := buildWithdrawTestcaseV2(v2.c, 243, 1, tinfo.addr, withdraw, address)
	auth.GasLimit = 0
	_, err = SubmitBurnProof(v2.p.v, auth2, proof)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()
	dappAbi, err := abi.JSON(strings.NewReader(dapp.DappABI))

	// check balance remain
	bal, err := v2.p.v.GetDepositedBalance(nil, tinfo.addr, address)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), withdraw, bal)
	timestamp := []byte(randomizeTimestamp())

	testCases := []struct {
		srcToken  common.Address
		srcQty    *big.Int
		destToken common.Address
		dapp      common.Address
		input     []byte
		vault     *vault.Vault
		timestamp []byte
		iserr     bool
	}{
		{
			// must be able to run execute smoothly by calling simpleCall func
			srcToken:  srcToken,
			srcQty:    executeAmount,
			destToken: destToken,
			dapp:      dAddr,
			input:     v2.packInputData(dappAbi, "simpleCall", destToken),
			vault:     v2.p.v,
			timestamp: timestamp,
			iserr:     false,
		},
		{
			// use sig twice
			srcToken:  srcToken,
			srcQty:    executeAmount,
			destToken: destToken,
			dapp:      dAddr,
			input:     v2.packInputData(dappAbi, "simpleCall", destToken),
			vault:     v2.p.v,
			timestamp: timestamp,
			iserr:     true,
		},
		{
			// trade same token
			srcToken:  srcToken,
			srcQty:    executeAmount,
			destToken: srcToken,
			dapp:      dAddr,
			input:     v2.packInputData(dappAbi, "simpleCall", destToken),
			vault:     v2.p.v,
			timestamp: []byte(randomizeTimestamp()),
			iserr:     true,
		},
		{
			// trade amount greater than amount available
			srcToken:  srcToken,
			srcQty:    withdraw,
			destToken: destToken,
			dapp:      dAddr,
			input:     v2.packInputData(dappAbi, "simpleCall", destToken),
			vault:     v2.p.v,
			timestamp: []byte(randomizeTimestamp()),
			iserr:     true,
		},
		{
			// execute dapp get revert
			srcToken:  srcToken,
			srcQty:    executeAmount,
			destToken: destToken,
			dapp:      dAddr,
			input:     v2.packInputData(dappAbi, "revertCall", destToken),
			vault:     v2.p.v,
			timestamp: []byte(randomizeTimestamp()),
			iserr:     true,
		},
		{
			// dapp return value amount but didn't transfer dest coin
			srcToken:  srcToken,
			srcQty:    executeAmount,
			destToken: destToken,
			dapp:      dAddr,
			input:     v2.packInputData(dappAbi, "ReturnAmountWithoutTranfer", destToken),
			vault:     v2.p.v,
			timestamp: []byte(randomizeTimestamp()),
			iserr:     true,
		},
	}

	for _, tc := range testCases {
		beforeExecute := getBalanceERC20(tinfo.c, v2.p.vAddr)
		_, err = runExecuteVault(auth2, tc.dapp, tc.srcToken, tc.srcQty, tc.destToken, tc.input, tc.vault, tc.timestamp, genesisAcc.PrivateKey)
		if tc.iserr {
			require.NotEqual(v2.T(), nil, err)
			v2.p.sim.Commit()
			require.Equal(v2.T(), beforeExecute, getBalanceERC20(tinfo.c, v2.p.vAddr))
		} else {
			require.Equal(v2.T(), nil, err)
			v2.p.sim.Commit()
			require.NotEqual(v2.T(), beforeExecute, getBalanceERC20(tinfo.c, v2.p.vAddr))
		}
	}
}

func (v2 *VaulV2TestSuite) TestVaultV2UpdateVaultThenTryExecute() {
	desc := "BNB"
	deposit := big.NewInt(int64(3e17))
	withdraw := big.NewInt(int64(2e8))
	executeAmount := big.NewInt(int64(1e17))
	address := crypto.PubkeyToAddress(genesisAcc.PrivateKey.PublicKey)
	// Deposit, must success
	tinfo := v2.p.customErc20s[desc]
	_, _, err := lockSimERC20WithTxs(v2.p, tinfo.c, tinfo.addr, deposit)
	require.Equal(v2.T(), nil, err)

	proof := buildWithdrawTestcaseV2(v2.c, 243, 1, tinfo.addr, withdraw, address)
	_, err = SubmitBurnProof(v2.p.v, auth2, proof)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()

	// check balance remain
	bal, err := v2.p.v.GetDepositedBalance(nil, tinfo.addr, address)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), bal, big.NewInt(0).Mul(withdraw, big.NewInt(int64(1e9))))

	// update newVault then user must be able to request execute
	nextVaultAddr, _, _, err := setupNextDelegator(auth, v2.p.sim)
	require.Equal(v2.T(), nil, err)
	_, err = v2.p.vp.Pause(auth)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()
	proxyVault, err := vaultproxy.NewTransparentUpgradeableProxy(v2.p.vAddr, v2.p.sim)
	require.Equal(v2.T(), nil, err)
	_, err = proxyVault.UpgradeTo(auth, nextVaultAddr)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()
	_, err = v2.p.vp.Unpause(auth)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()
	require.Equal(v2.T(), nil, err)

	require.Equal(v2.T(), nil, err)
	// check balance on vault with new delegator
	bal, err = v2.p.v.GetDepositedBalance(nil, tinfo.addr, address)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), bal, big.NewInt(0).Mul(withdraw, big.NewInt(int64(1e9))))

	dAddr, _, _, err := dapp.DeployDapp(auth, v2.p.sim, v2.p.vAddr)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()

	beforeExecute := getBalanceERC20(tinfo.c, v2.p.vAddr)
	require.Equal(v2.T(), beforeExecute, deposit)

	dappAbi, err := abi.JSON(strings.NewReader(dapp.DappABI))
	_, err = runExecuteVault(
		auth2,
		dAddr,
		tinfo.addr,
		executeAmount,
		v2.EtherAddress,
		v2.packInputData(dappAbi, "simpleCall", v2.EtherAddress),
		v2.p.v,
		[]byte(randomizeTimestamp()),
		genesisAcc.PrivateKey,
	)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()
	require.NotEqual(v2.T(), beforeExecute, getBalanceERC20(tinfo.c, v2.p.vAddr))
}

func (v2 *VaulV2TestSuite) TestVaultV2ExecuteRentranceAttack() {
	desc := "BNB"
	tinfo := v2.p.customErc20s[desc]
	withdraw := big.NewInt(int64(5e18))
	executeAmount := big.NewInt(int64(1e18))
	srcToken := v2.EtherAddress
	destToken := tinfo.addr
	address := crypto.PubkeyToAddress(genesisAcc.PrivateKey.PublicKey)
	dAddr, _, dappInst, err := dapp.DeployDapp(auth, v2.p.sim, v2.p.vAddr)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()
	_, _, err = deposit(v2.p, big.NewInt(int64(5e18)))
	require.Equal(v2.T(), nil, err)
	proof := buildWithdrawTestcaseV2(v2.c, 243, 1, v2.EtherAddress, withdraw, address)
	_, err = SubmitBurnProof(v2.p.v, auth2, proof)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()

	// must not be able to reentrance in any case
	callData, err := buildDataReentranceAttackData(srcToken, executeAmount, destToken, dAddr)
	require.Equal(v2.T(), nil, err)
	_, err = dappInst.ReEntranceAttack(auth2, destToken, callData)
	require.NotEqual(v2.T(), nil, err)
	v2.p.sim.Commit()
	bal, err := v2.p.v.GetDepositedBalance(nil, v2.EtherAddress, address)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), bal, withdraw)

	// reentrance transfer available amount twice must fail
	callData, err = buildDataReentranceAttackData(srcToken, executeAmount, destToken, dAddr)
	require.Equal(v2.T(), nil, err)
	_, err = dappInst.ReEntranceAttack(auth2, destToken, callData)
	require.NotEqual(v2.T(), nil, err)
	v2.p.sim.Commit()
	bal, err = v2.p.v.GetDepositedBalance(nil, v2.EtherAddress, address)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), bal, withdraw)
	require.Equal(v2.T(), withdraw, v2.p.getBalance(v2.p.vAddr))
}

func (v2 *VaulV2TestSuite) TestVaultV2sigToAddress() {
	timestamp := []byte(randomizeTimestamp())
	var data32 [32]byte
	data := rawsha3(timestamp)
	address := crypto.PubkeyToAddress(genesisAcc2.PrivateKey.PublicKey)

	// verify signer with right signature
	signBytes, _ := crypto.Sign(data, genesisAcc2.PrivateKey)
	copy(data32[:], data)
	signer, err := v2.p.v.SigToAddress(nil, signBytes, data32)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), signer, address)

	// wrong signature
	signBytes[0]++
	signer, err = v2.p.v.SigToAddress(nil, signBytes, data32)
	require.Equal(v2.T(), nil, err)
	require.NotEqual(v2.T(), signer, address)
}

func (v2 *VaulV2TestSuite) TestVaultV2isSigDataUsed() {
	desc := "USDT"
	testAmount := big.NewInt(int64(0))
	var data32 [32]byte
	tinfo := v2.p.customErc20s[desc]
	timestamp := []byte(randomizeTimestamp())
	// vaultAbi, _ := abi.JSON(strings.NewReader(vault.VaultABI))
	vaultHelperAbi, _ := abi.JSON(strings.NewReader(vault.VaultHelperABI))
	psData := vault.VaultHelperPreSignData{
		Prefix:    REQ_WITHDRAW_PREFIX,
		Token:     tinfo.addr,
		Timestamp: timestamp,
		Amount:    testAmount,
	}
	tempData, _ := vaultHelperAbi.Pack("_buildSignRequestWithdraw", psData, IncPaymentAddr)
	// tempData, _ := vaultAbi.Pack("withdrawBuildData", IncPaymentAddr, tinfo.addr, timestamp, testAmount)
	data := rawsha3(tempData[4:])

	copy(data32[:], data)
	signBytes, _ := crypto.Sign(data, genesisAcc.PrivateKey)
	txId := [32]byte{}
	// todo: update genesisAcc2.PrivateKey to real regulator
	sign, err := SignDataToShield(txId, genesisAcc2.PrivateKey, auth2.From)
	require.Equal(v2.T(), nil, err)
	_, err = v2.p.v.RequestWithdraw(auth2, IncPaymentAddr, tinfo.addr, testAmount, signBytes, timestamp, txId, sign)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()

	// datahash used must return true
	isUsed, err := v2.p.v.IsSigDataUsed(nil, data32)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), true, isUsed)

	// update vault to next version and check sig used
	nextDelegator, _, _, err := setupNextDelegator(auth, v2.p.sim)
	require.Equal(v2.T(), nil, err)
	_, err = v2.p.vp.Pause(auth)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()
	proxyVault, err := vaultproxy.NewTransparentUpgradeableProxy(v2.p.vAddr, v2.p.sim)
	require.Equal(v2.T(), nil, err)
	_, err = proxyVault.UpgradeTo(auth, nextDelegator)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()
	_, err = v2.p.vp.Unpause(auth)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()
	require.Equal(v2.T(), nil, err)

	isUsed, err = v2.p.v.IsSigDataUsed(nil, data32)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), true, isUsed)

	// datahash unsed must return false on new vault
	data32[0]++
	isUsed, err = v2.p.v.IsSigDataUsed(nil, data32)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), false, isUsed)
}

func buildWithdrawTestcaseV2(c *committees, meta, shard int, tokenID ec.Address, amount *big.Int, withdrawer common.Address) *DecodedProof {
	inst, mp, blkData, blkHash := buildWithdrawDataV2(meta, shard, tokenID, amount, withdrawer)
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

func buildWithdrawDataV2(meta, shard int, tokenID ec.Address, amount *big.Int, withdrawer common.Address) ([]byte, *merklePath, []byte, []byte) {
	// Build instruction merkle tree
	numInst := 10
	startNodeID := 7
	height := big.NewInt(1)
	inst := buildDecodedWithdrawInst(meta, shard, tokenID, withdrawer, amount)
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

func setupNextDelegator(
	auth *bind.TransactOpts,
	backend *backends.SimulatedBackend,
) (common.Address, *types.Transaction, *vault.Vault, error) {
	addr, tx, v, err := vault.DeployVault(auth, backend)
	if err != nil {
		return common.Address{}, nil, nil, fmt.Errorf("failed to deploy Vault contract: %v", err)
	}
	backend.Commit()
	return addr, tx, v, nil
}

func runExecuteVault(
	auth *bind.TransactOpts,
	dapp common.Address,
	srcToken common.Address,
	srcQty *big.Int,
	destoken common.Address,
	input []byte,
	vaultInput *vault.Vault,
	timestamp []byte,
	signer *ecdsa.PrivateKey,
) (*types.Transaction, error) {
	vaultAbi, _ := abi.JSON(strings.NewReader(vault.VaultHelperABI))
	psData := vault.VaultHelperPreSignData{
		Prefix:    EXECUTE_PREFIX,
		Token:     srcToken,
		Timestamp: timestamp,
		Amount:    srcQty,
	}
	tempData, err := vaultAbi.Pack("_buildSignExecute", psData, destoken, dapp, input)
	// tempData, _ := vaultAbi.Pack("executeBuildData", dapp, input, timestamp, srcQty)
	data := rawsha3(tempData[4:])
	signBytes, err := crypto.Sign(data, signer)
	if err != nil {
		return nil, err
	}
	tx, err := vaultInput.Execute(
		auth,
		srcToken,
		srcQty,
		destoken,
		dapp,
		input,
		timestamp,
		signBytes,
	)
	if err != nil {
		return nil, err
	}
	return tx, err
}

func (v2 *VaulV2TestSuite) packInputData(abi abi.ABI, method string, dest common.Address) []byte {
	callData, err := abi.Pack(method, dest)
	if err != nil {
		require.Equal(v2.T(), nil, err)
	}
	return callData
}

func buildDataReentranceAttackData(
	srcToken common.Address,
	executeAmount *big.Int,
	destToken common.Address,
	dAddr common.Address,
) ([]byte, error) {
	dappAbi, err := abi.JSON(strings.NewReader(dapp.DappABI))
	if err != nil {
		return nil, err
	}
	vaultAbi, err := abi.JSON(strings.NewReader(vault.VaultABI))
	if err != nil {
		return nil, err
	}
	timestamp := []byte(randomizeTimestamp())
	input1, err := dappAbi.Pack("simpleCall", destToken)
	if err != nil {
		return nil, err
	}
	vaultHelperAbi, _ := abi.JSON(strings.NewReader(vault.VaultHelperABI))
	psData := vault.VaultHelperPreSignData{
		Prefix:    EXECUTE_PREFIX,
		Token:     srcToken,
		Timestamp: timestamp,
		Amount:    executeAmount,
	}
	tempData, err := vaultHelperAbi.Pack("_buildSignExecute", psData, destToken, dAddr, input1)
	// tempData, _ := vaultAbi.Pack("executeBuildData", dAddr, input1, timestamp, executeAmount)

	data := rawsha3(tempData[4:])
	signBytes, err := crypto.Sign(data, genesisAcc.PrivateKey)
	if err != nil {
		return nil, err
	}
	input2, err := vaultAbi.Pack(
		"execute",
		srcToken,
		executeAmount,
		destToken,
		dAddr,
		input1,
		timestamp,
		signBytes,
	)
	if err != nil {
		return nil, err
	}
	input3, err := dappAbi.Pack("ReEntranceAttack", destToken, input2)
	if err != nil {
		return nil, err
	}
	psData = vault.VaultHelperPreSignData{
		Prefix:    EXECUTE_PREFIX,
		Token:     srcToken,
		Timestamp: timestamp,
		Amount:    executeAmount,
	}
	tempData, err = vaultHelperAbi.Pack("_buildSignExecute", psData, destToken, dAddr, input3)
	// tempData, _ = vaultAbi.Pack("executeBuildData", dAddr, input3, timestamp, executeAmount)
	data = rawsha3(tempData[4:])
	signBytes, err = crypto.Sign(data, genesisAcc.PrivateKey)
	if err != nil {
		return nil, err
	}
	input4, err := vaultAbi.Pack(
		"execute",
		srcToken,
		executeAmount,
		destToken,
		dAddr,
		input3,
		timestamp,
		signBytes,
	)
	if err != nil {
		return nil, err
	}
	return input4, nil
}
