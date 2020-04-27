package main

import (
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/incognitochain/bridge-eth/bridge/dapp"

	"github.com/incognitochain/bridge-eth/bridge/nextVault"

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
	v2.v, err = vault.NewVault(v2.p.vAddr, v2.p.sim)
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
	_, err = SubmitBurnProof(v2.p.v, auth, proof)
	require.NotEqual(v2.T(), nil, err)
	v2.p.sim.Commit()

	// wrong shard
	meta = 97
	shardID = 2
	proof = buildWithdrawTestcase(v2.c, meta, shardID, tinfo.addr, withdraw)

	auth.GasLimit = 0
	_, err = SubmitBurnProof(v2.p.v, auth, proof)
	require.NotEqual(v2.T(), nil, err)
	v2.p.sim.Commit()

	// able to submit proof
	shardID = 1
	proof = buildWithdrawTestcase(v2.c, meta, shardID, tinfo.addr, withdraw)

	auth.GasLimit = 0
	_, err = SubmitBurnProof(v2.p.v, auth, proof)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()

	// Check balance
	require.Equal(v2.T(), nil, err)
	bal, err := v2.v.GetDepositedBalance(nil, tinfo.addr, v2.withdrawer)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), bal, big.NewInt(0).Mul(withdraw, big.NewInt(int64(1e9))))

	// use proof twice
	_, err = SubmitBurnProof(v2.p.v, auth, proof)
	require.NotEqual(v2.T(), nil, err)
	v2.p.sim.Commit()

	// pause and submitBurnProof
	_, err = v2.v.Pause(auth)
	require.Equal(v2.T(), nil, err)
	proof = buildWithdrawTestcase(v2.c, meta, shardID, tinfo.addr, withdraw)
	auth.GasLimit = 0
	_, err = SubmitBurnProof(v2.p.v, auth, proof)
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

	proof := buildWithdrawTestcaseV2(v2.c, 97, 1, tinfo.addr, withdraw, address)
	auth.GasLimit = 0
	_, err = SubmitBurnProof(v2.p.v, auth, proof)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()

	// request amount bigger than balance
	timestamp := []byte(randomizeTimestamp())
	tempData := append([]byte(IncPaymentAddr), tinfo.addr[:]...)
	tempData1 := append(tempData, timestamp...)
	tempData2 := append(tempData1, common.LeftPadBytes(deposit.Bytes(), 32)...)
	data := rawsha3(tempData2)
	signBytes, _ := crypto.Sign(data, genesisAcc.PrivateKey)
	_, err = v2.v.RequestWithdraw(auth, IncPaymentAddr, tinfo.addr, deposit, signBytes, timestamp)
	require.NotEqual(v2.T(), nil, err)
	v2.p.sim.Commit()

	// able to request withdraw
	tempData2 = append(tempData1, common.LeftPadBytes(big.NewInt(0).Mul(redeposit, big.NewInt(int64(1e9))).Bytes(), 32)...)
	data = rawsha3(tempData2)
	signBytes, _ = crypto.Sign(data, genesisAcc.PrivateKey)
	_, err = v2.v.RequestWithdraw(auth, IncPaymentAddr, tinfo.addr, big.NewInt(0).Mul(redeposit, big.NewInt(int64(1e9))), signBytes, timestamp)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()

	// use signature twice
	_, err = v2.v.RequestWithdraw(auth, IncPaymentAddr, tinfo.addr, big.NewInt(0).Mul(redeposit, big.NewInt(int64(1e9))), signBytes, timestamp)
	require.NotEqual(v2.T(), nil, err)
	v2.p.sim.Commit()

	// check balance remain
	bal, err := v2.v.GetDepositedBalance(nil, tinfo.addr, address)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), bal, big.NewInt(0).Mul(redeposit, big.NewInt(int64(1e9))))

	// amount subtracted so can not request amount as amount at time withdraw from incognito
	timestamp = []byte(randomizeTimestamp())
	tempData1 = append(tempData, timestamp...)
	tempData2 = append(tempData1, common.LeftPadBytes(big.NewInt(0).Mul(withdraw, big.NewInt(int64(1e9))).Bytes(), 32)...)
	data = rawsha3(tempData2)
	signBytes, _ = crypto.Sign(data, genesisAcc.PrivateKey)
	_, err = v2.v.RequestWithdraw(auth, IncPaymentAddr, tinfo.addr, big.NewInt(0).Mul(withdraw, big.NewInt(int64(1e9))), signBytes, timestamp)
	require.NotEqual(v2.T(), nil, err)
	v2.p.sim.Commit()

	// update newVault then user must be able to request withdraw
	nextVaultAddr, _, nextVault, err := setupNextVault(auth, v2.p.sim, auth.From, v2.p.incAddr, v2.p.vAddr)
	require.Equal(v2.T(), nil, err)
	_, err = v2.v.Pause(auth)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()
	_, err = v2.v.Migrate(auth, nextVaultAddr)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()
	_, err = v2.v.MoveAssets(auth, []common.Address{tinfo.addr})
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()
	totalDeposit, err := v2.v.TotalDepositedToSCAmount(nil, tinfo.addr)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), 0, totalDeposit.Cmp(big.NewInt(0)))
	totalDeposit, err = nextVault.TotalDepositedToSCAmount(nil, tinfo.addr)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), big.NewInt(0).Mul(redeposit, big.NewInt(int64(1e9))), totalDeposit)

	// check balance from nextVault
	bal, err = nextVault.GetDepositedBalance(nil, tinfo.addr, address)

	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), bal, big.NewInt(0).Mul(redeposit, big.NewInt(int64(1e9))))

	tempData2 = append(tempData1, common.LeftPadBytes(big.NewInt(0).Mul(redeposit, big.NewInt(int64(1e9))).Bytes(), 32)...)
	data = rawsha3(tempData2)
	signBytes, err = crypto.Sign(data, genesisAcc.PrivateKey)
	require.Equal(v2.T(), nil, err)
	_, err = nextVault.RequestWithdraw(auth, IncPaymentAddr, tinfo.addr, big.NewInt(0).Mul(redeposit, big.NewInt(int64(1e9))), signBytes, timestamp)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()

	// check balance from nextVault after RequestWithdraw
	bal, err = nextVault.GetDepositedBalance(nil, tinfo.addr, address)
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
	address := crypto.PubkeyToAddress(genesisAcc.PrivateKey.PublicKey)
	dAddr, _, _, err := dapp.DeployDapp(auth, v2.p.sim, v2.p.vAddr)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()
	_, _, err = deposit(v2.p, big.NewInt(int64(5e18)))
	require.Equal(v2.T(), nil, err)
	proof := buildWithdrawTestcaseV2(v2.c, 97, 1, v2.EtherAddress, withdraw, address)
	auth.GasLimit = 0
	_, err = SubmitBurnProof(v2.p.v, auth, proof)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()
	dappAbi, err := abi.JSON(strings.NewReader(dapp.DappABI))

	// check balance remain
	bal, err := v2.v.GetDepositedBalance(nil, v2.EtherAddress, address)
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
			vault:     v2.v,
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
			vault:     v2.v,
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
			vault:     v2.v,
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
			vault:     v2.v,
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
			vault:     v2.v,
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
			vault:     v2.v,
			timestamp: []byte(randomizeTimestamp()),
			iserr:     true,
		},
	}

	for _, tc := range testCases {
		beforeExecute := v2.p.getBalance(v2.p.vAddr)
		_, err = runExecuteVault(auth, tc.dapp, tc.srcToken, tc.srcQty, tc.destToken, tc.input, tc.vault, tc.timestamp)
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
	proof := buildWithdrawTestcaseV2(v2.c, 97, 1, tinfo.addr, withdraw, address)
	auth.GasLimit = 0
	_, err = SubmitBurnProof(v2.p.v, auth, proof)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()
	dappAbi, err := abi.JSON(strings.NewReader(dapp.DappABI))

	// check balance remain
	bal, err := v2.v.GetDepositedBalance(nil, tinfo.addr, address)
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
			vault:     v2.v,
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
			vault:     v2.v,
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
			vault:     v2.v,
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
			vault:     v2.v,
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
			vault:     v2.v,
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
			vault:     v2.v,
			timestamp: []byte(randomizeTimestamp()),
			iserr:     true,
		},
	}

	for _, tc := range testCases {
		beforeExecute := getBalanceERC20(tinfo.c, v2.p.vAddr)
		_, err = runExecuteVault(auth, tc.dapp, tc.srcToken, tc.srcQty, tc.destToken, tc.input, tc.vault, tc.timestamp)
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

	proof := buildWithdrawTestcaseV2(v2.c, 97, 1, tinfo.addr, withdraw, address)
	auth.GasLimit = 0
	_, err = SubmitBurnProof(v2.p.v, auth, proof)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()

	// check balance remain
	bal, err := v2.v.GetDepositedBalance(nil, tinfo.addr, address)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), bal, big.NewInt(0).Mul(withdraw, big.NewInt(int64(1e9))))

	// update newVault then user must be able to request execute
	nextVaultAddr, _, _, err := setupNextVault(auth, v2.p.sim, auth.From, v2.p.incAddr, v2.p.vAddr)
	require.Equal(v2.T(), nil, err)
	_, err = v2.v.Pause(auth)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()
	_, err = v2.v.Migrate(auth, nextVaultAddr)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()
	_, err = v2.v.MoveAssets(auth, []common.Address{tinfo.addr})
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()

	wrapNextVault, err := vault.NewVault(nextVaultAddr, v2.p.sim)
	require.Equal(v2.T(), nil, err)
	// check balance on nextVault
	bal, err = wrapNextVault.GetDepositedBalance(nil, tinfo.addr, address)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), bal, big.NewInt(0).Mul(withdraw, big.NewInt(int64(1e9))))

	dAddr, _, _, err := dapp.DeployDapp(auth, v2.p.sim, nextVaultAddr)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()

	beforeExecute := getBalanceERC20(tinfo.c, nextVaultAddr)
	require.Equal(v2.T(), beforeExecute, deposit)

	dappAbi, err := abi.JSON(strings.NewReader(dapp.DappABI))
	_, err = runExecuteVault(
		auth,
		dAddr,
		tinfo.addr,
		executeAmount,
		v2.EtherAddress,
		v2.packInputData(dappAbi, "simpleCall", v2.EtherAddress),
		wrapNextVault,
		[]byte(randomizeTimestamp()),
	)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()
	require.NotEqual(v2.T(), beforeExecute, getBalanceERC20(tinfo.c, nextVaultAddr))
}

func (v2 *VaulV2TestSuite) TestVaultV2ExecuteRentranceAttack() {
	desc := "BNB"
	tinfo := v2.p.customErc20s[desc]
	withdraw := big.NewInt(int64(5e18))
	executeAmount := big.NewInt(int64(1e18))
	remain := big.NewInt(int64(3e18))
	srcToken := v2.EtherAddress
	destToken := tinfo.addr
	address := crypto.PubkeyToAddress(genesisAcc.PrivateKey.PublicKey)
	dAddr, _, dappInst, err := dapp.DeployDapp(auth, v2.p.sim, v2.p.vAddr)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()
	_, _, err = deposit(v2.p, big.NewInt(int64(5e18)))
	require.Equal(v2.T(), nil, err)
	proof := buildWithdrawTestcaseV2(v2.c, 97, 1, v2.EtherAddress, withdraw, address)
	auth.GasLimit = 0
	_, err = SubmitBurnProof(v2.p.v, auth, proof)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()

	// must be able to reentrance if enough amount
	callData, err := buildDataReentranceAttackData(srcToken, executeAmount, destToken, dAddr)
	require.Equal(v2.T(), nil, err)
	_, err = dappInst.ReEntranceAttack(auth, destToken, callData)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()
	bal, err := v2.v.GetDepositedBalance(nil, v2.EtherAddress, address)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), bal, remain)

	// reentrance transfer available amount twice must fail
	callData, err = buildDataReentranceAttackData(srcToken, remain, destToken, dAddr)
	require.Equal(v2.T(), nil, err)
	_, err = dappInst.ReEntranceAttack(auth, destToken, callData)
	require.NotEqual(v2.T(), nil, err)
	v2.p.sim.Commit()
	bal, err = v2.v.GetDepositedBalance(nil, v2.EtherAddress, address)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), bal, remain)
	require.Equal(v2.T(), remain, v2.p.getBalance(v2.p.vAddr))
}

func (v2 *VaulV2TestSuite) TestVaultV2sigToAddress() {
	timestamp := []byte(randomizeTimestamp())
	var data32 [32]byte
	data := rawsha3(timestamp)
	address := crypto.PubkeyToAddress(genesisAcc.PrivateKey.PublicKey)

	// verify signer with right signature
	signBytes, _ := crypto.Sign(data, genesisAcc.PrivateKey)
	copy(data32[:], data)
	signer, err := v2.v.SigToAddress(nil, signBytes, data32)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), signer, address)

	// wrong signature
	signBytes[0]++
	signer, err = v2.v.SigToAddress(nil, signBytes, data32)
	require.Equal(v2.T(), nil, err)
	require.NotEqual(v2.T(), signer, address)
}

func (v2 *VaulV2TestSuite) TestVaultV2isSigDataUsed() {
	desc := "USDT"
	testAmount := big.NewInt(int64(0))
	var data32 [32]byte
	tinfo := v2.p.customErc20s[desc]
	timestamp := []byte(randomizeTimestamp())
	tempData := append([]byte(IncPaymentAddr), tinfo.addr[:]...)
	tempData1 := append(tempData, timestamp...)
	tempData2 := append(tempData1, common.LeftPadBytes(testAmount.Bytes(), 32)...)
	data := rawsha3(tempData2)

	copy(data32[:], data)
	signBytes, _ := crypto.Sign(data, genesisAcc.PrivateKey)
	_, err := v2.v.RequestWithdraw(auth, IncPaymentAddr, tinfo.addr, testAmount, signBytes, timestamp)
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()

	// datahash used must return true
	isUsed, err := v2.v.IsSigDataUsed(nil, data32)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), true, isUsed)

	// update vault to next version and check sig used
	_, _, nextVault, err := setupNextVault(auth, v2.p.sim, auth.From, v2.p.incAddr, v2.p.vAddr)
	require.Equal(v2.T(), nil, err)
	isUsed, err = nextVault.IsSigDataUsed(nil, data32)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), true, isUsed)

	// datahash unsed must return false on new vault
	data32[0]++
	isUsed, err = nextVault.IsSigDataUsed(nil, data32)
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), false, isUsed)
}

func buildWithdrawTestcaseV2(c *committees, meta, shard int, tokenID ec.Address, amount *big.Int, withdrawer common.Address) *decodedProof {
	inst, mp, blkData, blkHash := buildWithdrawDataV2(meta, shard, tokenID, amount, withdrawer)
	ipBeacon := signAndReturnInstProof(c.beaconPrivs, true, mp, blkData, blkHash[:])
	ipBridge := signAndReturnInstProof(c.bridgePrivs, false, mp, blkData, blkHash[:])
	return &decodedProof{
		Instruction: inst,
		Heights:     [2]*big.Int{big.NewInt(1), big.NewInt(1)},

		InstPaths:       [2][][32]byte{ipBeacon.instPath, ipBridge.instPath},
		InstPathIsLefts: [2][]bool{ipBeacon.instPathIsLeft, ipBridge.instPathIsLeft},
		InstRoots:       [2][32]byte{ipBeacon.instRoot, ipBridge.instRoot},
		BlkData:         [2][32]byte{ipBeacon.blkData, ipBridge.blkData},
		SigIdxs:         [2][]*big.Int{ipBeacon.sigIdx, ipBridge.sigIdx},
		SigVs:           [2][]uint8{ipBeacon.sigV, ipBridge.sigV},
		SigRs:           [2][][32]byte{ipBeacon.sigR, ipBridge.sigR},
		SigSs:           [2][][32]byte{ipBeacon.sigS, ipBridge.sigS},
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

func setupNextVault(
	auth *bind.TransactOpts,
	backend *backends.SimulatedBackend,
	admin, incAddr, prevVault common.Address,
) (common.Address, *types.Transaction, *nextVault.NextVault, error) {
	addr, tx, v, err := nextVault.DeployNextVault(auth, backend, admin, incAddr, prevVault)
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
	vault *vault.Vault,
	timestamp []byte,
) (*types.Transaction, error) {
	tempData := append(dapp[:], input...)
	tempData1 := append(tempData, timestamp...)
	tempData2 := append(tempData1, common.LeftPadBytes(srcQty.Bytes(), 32)...)
	data := rawsha3(tempData2)
	signBytes, err := crypto.Sign(data, genesisAcc.PrivateKey)
	if err != nil {
		return nil, err
	}
	tx, err := vault.Execute(
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
	tempData := append(dAddr[:], input1...)
	tempData1 := append(tempData, timestamp...)
	tempData2 := append(tempData1, common.LeftPadBytes(executeAmount.Bytes(), 32)...)
	data := rawsha3(tempData2)
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
	tempData = append(dAddr[:], input3...)
	tempData1 = append(tempData, timestamp...)
	tempData2 = append(tempData1, common.LeftPadBytes(executeAmount.Bytes(), 32)...)
	data = rawsha3(tempData2)
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
