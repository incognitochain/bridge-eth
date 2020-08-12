package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/incognitochain/bridge-eth/bridge/locker"
	"github.com/incognitochain/bridge-eth/bridge/reentrancer"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestFixedLockerChangeVault(t *testing.T) {
	rand := newAccount()
	testCases := []struct {
		desc  string
		admin *account
		err   bool
	}{
		{
			desc:  "Success",
			admin: genesisAcc,
		},
		{
			desc:  "Not admin",
			admin: rand,
			err:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			p, err := setupLocker(tc.admin.Address, common.Address{})
			assert.Nil(t, err)

			_, err = p.locker.ChangeVault(auth, auth.From)
			isErr := err != nil
			if isErr != tc.err {
				t.Fatal(errors.Errorf("expect error = %t, got %v", tc.err, err))
			}
			if tc.err {
				return
			}
			p.sim.Commit()

			// Check vault address
			addr, _ := p.locker.Vault(nil)
			assert.NotEqual(t, common.Address{}, addr)
		})
	}
}

func TestFixedLockerSetNextLocker(t *testing.T) {
	rand := newAccount()
	testCases := []struct {
		desc  string
		admin *account
		pause bool
		err   bool
	}{
		{
			desc:  "Success",
			admin: genesisAcc,
			pause: true,
		},
		{
			desc:  "Not admin",
			admin: rand,
			pause: true,
			err:   true,
		},
		{
			desc:  "Not paused",
			admin: genesisAcc,
			pause: false,
			err:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			p, err := setupLocker(tc.admin.Address, common.Address{})
			assert.Nil(t, err)

			if tc.pause {
				auth := bind.NewKeyedTransactor(tc.admin.PrivateKey)
				_, err := p.locker.Pause(auth)
				assert.Nil(t, err)
			}

			_, err = p.locker.SetNextLocker(auth, auth.From)
			isErr := err != nil
			if isErr != tc.err {
				t.Fatal(errors.Errorf("expect error = %t, got %v", tc.err, err))
			}
			if tc.err {
				return
			}
			p.sim.Commit()

			// Check nextLocker address
			addr, _ := p.locker.NextLocker(nil)
			assert.Equal(t, auth.From, addr)
		})
	}
}

func TestFixedLockerUnlock(t *testing.T) {
	rand := newAccount()
	p, err := setupLocker(genesisAcc.Address, common.Address{}, rand.Address)
	assert.Nil(t, err)
	helper := &lockerHelper{p, t, auth}

	p.contracts.customErc20s = map[string]*TokenerInfo{}
	err = setupCustomTokens(p)
	assert.Nil(t, err)

	// Prepare: transfer tokens to Locker
	assetNames := []string{"BNB", "USDT", "DAI", "FAIL"}
	assets := []common.Address{}
	for _, name := range assetNames {
		_, err := p.contracts.customErc20s[name].c.Transfer(auth, p.lockerAddr, big.NewInt(1000))
		assert.Nil(t, err)
	}

	// Test: failed when NextLocker is 0x0
	helper.PauseOK()
	strContains(t, helper.UnlockWithErrorReason(auth, assets), "null nextLocker")

	// Test: failed when Locker is not paused
	helper.SetNextLockerOK(rand.Address)
	helper.UnpauseOK()
	strContains(t, helper.UnlockWithErrorReason(auth, assets), "not paused right now")

	// Test: failed when not called by admin
	helper.PauseOK()
	auth2 := bind.NewKeyedTransactor(rand.PrivateKey)
	strContains(t, helper.UnlockWithErrorReason(auth2, assets), "not admin")

	// Test: tx still success when 1 token failed
	helper.UnlockOK(
		[]common.Address{
			p.contracts.customErc20s["BNB"].addr,
			p.contracts.customErc20s["FAIL"].addr,
		},
	)
	helper.CheckBalance(rand.Address, "BNB", 1000)
	helper.CheckBalance(rand.Address, "FAIL", 1000)

	// Test: all tokens succeed
	helper.UnlockOK(
		[]common.Address{
			p.contracts.customErc20s["USDT"].addr,
			p.contracts.customErc20s["DAI"].addr,
		},
	)
	helper.CheckBalance(rand.Address, "USDT", 1000)
	helper.CheckBalance(rand.Address, "DAI", 1000)
}

func TestFixedLockerUnlockReentrance(t *testing.T) {
	p, reAddr, re, err := setupLockerWithReentrancer()
	assert.Nil(t, err)

	// Prepare: transfer eth to Locker
	assert.Nil(t, transferETH(p.sim, genesisAcc.PrivateKey, p.lockerAddr, big.NewInt(1000)))
	assert.Equal(t, big.NewInt(1000), p.getBalance(p.lockerAddr))

	// Pause locker
	_, err = re.PauseLocker(auth, p.lockerAddr)
	assert.Nil(t, err)
	p.sim.Commit()

	// Set NextLocker
	_, err = re.SetNextLocker(auth, p.lockerAddr, reAddr)
	assert.Nil(t, err)
	p.sim.Commit()

	// 1st unlock: no reentrancy, must success
	assets := []common.Address{common.Address{}} // ETH
	_, err = re.ReUnlock(auth, big.NewInt(0), big.NewInt(1), p.lockerAddr, assets)
	assert.Nil(t, err)
	p.sim.Commit()
	assert.Equal(t, big.NewInt(0), p.getBalance(p.lockerAddr))

	// 2nd unlock: reentrancy, must fail
	_, err = re.SetFallbackUnlockData(auth, assets) // Reentrancy the same asset
	assert.Nil(t, err)
	assert.Nil(t, transferETH(p.sim, genesisAcc.PrivateKey, p.lockerAddr, big.NewInt(1000))) // Prepare: transfer eth to Locker
	assert.Equal(t, big.NewInt(1000), p.getBalance(p.lockerAddr))

	auth.GasLimit = 1000000
	defer func() {
		auth.GasLimit = 0
	}()
	_, err = re.ReUnlock(auth, big.NewInt(0), big.NewInt(0), p.lockerAddr, assets)
	assert.Nil(t, err)
	p.sim.Commit()

	assert.Equal(t, big.NewInt(1000), p.getBalance(p.lockerAddr)) // Balance must not change
}

func TestFixedLockerGive(t *testing.T) {
	rand := newAccount()
	p, err := setupLocker(genesisAcc.Address, rand.Address)
	assert.Nil(t, err)
	helper := &lockerHelper{p, t, auth}

	p.contracts.customErc20s = map[string]*TokenerInfo{}
	err = setupCustomTokens(p)
	assert.Nil(t, err)

	// Prepare: transfer tokens to Locker
	assetNames := []string{"USDT", "DAI", "FAIL"}
	erc20s := p.contracts.customErc20s
	for _, name := range assetNames {
		_, err := erc20s[name].c.Transfer(auth, p.lockerAddr, big.NewInt(1000))
		assert.Nil(t, err)
	}
	assert.Nil(t, transferETH(p.sim, genesisAcc.PrivateKey, p.lockerAddr, big.NewInt(1000)))

	// Test: failed when not called by vault
	eth := common.Address{}
	strContains(t, helper.GiveWithErrorReason(auth, rand.Address, eth, big.NewInt(1000)), "not vault")
	bigEqual(t, big.NewInt(1000), p.getBalance(p.lockerAddr))

	// Test: failed when paused
	helper.PauseOK()
	_, err = p.locker.ChangeVault(auth, genesisAcc.Address)
	assert.Nil(t, err)

	strContains(t, helper.GiveWithErrorReason(auth, rand.Address, eth, big.NewInt(1000)), "paused right now")
	bigEqual(t, big.NewInt(1000), p.getBalance(p.lockerAddr))
	bigEqual(t, big.NewInt(0), p.getBalance(rand.Address))

	// Test: give ETH successfully
	helper.UnpauseOK()
	helper.GiveOK(rand.Address, eth, big.NewInt(1000))
	bigEqual(t, big.NewInt(0), p.getBalance(p.lockerAddr))
	bigEqual(t, big.NewInt(1000), p.getBalance(rand.Address))

	// Test: give ERC20 successfully
	name := "DAI"
	helper.GiveOK(rand.Address, erc20s[name].addr, big.NewInt(1000))
	bigEqual(t, big.NewInt(0), getBalanceERC20(erc20s[name].c, p.lockerAddr))
	bigEqual(t, big.NewInt(1000), getBalanceERC20(erc20s[name].c, rand.Address))

	// Test: success even for ERC20s with invalid transfer()
	name = "USDT"
	helper.GiveOK(rand.Address, erc20s[name].addr, big.NewInt(1000))
	bigEqual(t, big.NewInt(0), getBalanceERC20(erc20s[name].c, p.lockerAddr))
	bigEqual(t, big.NewInt(1000), getBalanceERC20(erc20s[name].c, rand.Address))

	// Test: failed for invalid token
	name = "FAIL"
	strContains(t, helper.GiveWithErrorReason(auth, rand.Address, erc20s[name].addr, big.NewInt(1000)), "failed to give erc20")
	bigEqual(t, big.NewInt(1000), getBalanceERC20(erc20s[name].c, p.lockerAddr))
	bigEqual(t, big.NewInt(0), getBalanceERC20(erc20s[name].c, rand.Address))
}

func TestFixedLockerGiveReentrance(t *testing.T) {
	rand := newAccount()
	p, reAddr, re, err := setupLockerWithReentrancer()
	assert.Nil(t, err)

	// Prepare: transfer eth to Locker
	assert.Nil(t, transferETH(p.sim, genesisAcc.PrivateKey, p.lockerAddr, big.NewInt(1000)))
	assert.Equal(t, big.NewInt(1000), p.getBalance(p.lockerAddr))

	// 1st give: no reentrancy, must success
	eth := common.Address{}
	_, err = re.ReGive(auth, big.NewInt(1), big.NewInt(1), p.lockerAddr, rand.Address, eth, big.NewInt(300))
	assert.Nil(t, err)
	p.sim.Commit()
	bigEqual(t, big.NewInt(700), p.getBalance(p.lockerAddr))
	bigEqual(t, big.NewInt(300), p.getBalance(rand.Address))

	// 2nd give: reentrancy, must fail
	_, err = re.SetFallbackGiveData(auth, reAddr, eth, big.NewInt(200)) // Reentrancy data
	assert.Nil(t, err)

	auth.GasLimit = 1000000
	defer func() {
		auth.GasLimit = 0
	}()
	tx, err := re.ReGive(auth, big.NewInt(1), big.NewInt(0), p.lockerAddr, reAddr, eth, big.NewInt(200))
	assert.Nil(t, err)
	p.sim.Commit()

	bigEqual(t, big.NewInt(700), p.getBalance(p.lockerAddr)) // Balance must not change
	reason, _ := errorReason(context.Background(), p.sim, tx, nil, auth.From)
	strContains(t, reason, "failed to give eth")
}

func setupBackend(accs ...common.Address) *backends.SimulatedBackend {
	alloc := make(core.GenesisAlloc)
	balance, _ := big.NewInt(1).SetString("100000000000000000000", 10) // 100 eth
	alloc[auth.From] = core.GenesisAccount{Balance: balance}
	for _, acc := range accs {
		alloc[acc] = core.GenesisAccount{Balance: balance}
	}
	return backends.NewSimulatedBackend(alloc, 8000000)
}

func setupLockerContract(sim *backends.SimulatedBackend, admin, vault common.Address) (common.Address, *locker.Locker, error) {
	lockerAddr, _, l, err := locker.DeployLocker(auth, sim, admin, vault)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("failed to deploy Locker contract: %v", err)
	}
	sim.Commit()
	return lockerAddr, l, nil
}

func setupLocker(admin, vault common.Address, accs ...common.Address) (*Platform, error) {
	p := &Platform{sim: setupBackend(append(accs, admin)...), contracts: &contracts{}}
	var err error
	p.lockerAddr, p.locker, err = setupLockerContract(p.sim, admin, vault)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func setupLockerWithReentrancer() (*Platform, common.Address, *reentrancer.Reentrancer, error) {
	p := &Platform{sim: setupBackend(), contracts: &contracts{}}
	reAddr, _, re, err := reentrancer.DeployReentrancer(auth, p.sim)
	if err != nil {
		return nil, common.Address{}, nil, err
	}
	p.sim.Commit()

	p.lockerAddr, p.locker, err = setupLockerContract(p.sim, reAddr, reAddr) // Reentrancer as both admin and vault
	if err != nil {
		return nil, common.Address{}, nil, err
	}
	return p, reAddr, re, nil
}

type lockerHelper struct {
	p    *Platform
	t    *testing.T
	auth *bind.TransactOpts
}

func (helper *lockerHelper) PauseOK() {
	_, err := helper.p.locker.Pause(helper.auth)
	assert.Nil(helper.t, err)
	helper.p.sim.Commit()
}

func (helper *lockerHelper) UnpauseOK() {
	_, err := helper.p.locker.Unpause(helper.auth)
	assert.Nil(helper.t, err)
	helper.p.sim.Commit()
}

func (helper *lockerHelper) SetNextLockerOK(addr common.Address) {
	_, err := helper.p.locker.SetNextLocker(helper.auth, addr)
	assert.Nil(helper.t, err)
	helper.p.sim.Commit()
}

func (helper *lockerHelper) UnlockOK(assets []common.Address) {
	helper.auth.GasLimit = 0
	_, err := helper.p.locker.Unlock(helper.auth, assets)
	assert.Nil(helper.t, err)
	helper.p.sim.Commit()
}

func (helper *lockerHelper) CheckBalance(addr common.Address, asset string, bal int) {
	val, _ := helper.p.contracts.customErc20s[asset].c.BalanceOf(nil, addr)
	assert.Equal(helper.t, int64(bal), val.Int64())
}

func (helper *lockerHelper) UnlockWithErrorReason(auth *bind.TransactOpts, assets []common.Address) string {
	auth.GasLimit = 1000000
	defer func() {
		auth.GasLimit = 0
	}()

	tx, err := helper.p.locker.Unlock(auth, assets)
	assert.Nil(helper.t, err)
	helper.p.sim.Commit()
	reason, _ := errorReason(context.Background(), helper.p.sim, tx, nil, auth.From)
	return reason
}

func (helper *lockerHelper) GiveOK(to, token common.Address, amount *big.Int) {
	helper.auth.GasLimit = 0
	_, err := helper.p.locker.Give(helper.auth, to, token, amount)
	assert.Nil(helper.t, err)
	helper.p.sim.Commit()
}

func (helper *lockerHelper) GiveWithErrorReason(auth *bind.TransactOpts, to, token common.Address, amount *big.Int) string {
	auth.GasLimit = 1000000
	defer func() {
		auth.GasLimit = 0
	}()

	tx, err := helper.p.locker.Give(auth, to, token, amount)
	assert.Nil(helper.t, err)
	helper.p.sim.Commit()
	reason, _ := errorReason(context.Background(), helper.p.sim, tx, nil, auth.From)
	return reason
}

func transferETH(
	sim *backends.SimulatedBackend,
	privKey *ecdsa.PrivateKey,
	toAddress common.Address,
	value *big.Int,
) error {
	ctx := context.Background()
	nonce, err := sim.PendingNonceAt(ctx, crypto.PubkeyToAddress(privKey.PublicKey))
	if err != nil {
		return err
	}

	tx := types.NewTransaction(nonce, toAddress, value, 30000, nil, nil)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(params.AllEthashProtocolChanges.ChainID), privKey)
	if err != nil {
		return err
	}

	err = sim.SendTransaction(ctx, signedTx)
	if err != nil {
		return err
	}

	sim.Commit()
	return nil
}

func bigEqual(t *testing.T, x, y *big.Int) {
	assert.True(t, x.Cmp(y) == 0, "expected %v, got %v", x, y)
}

func strContains(t *testing.T, container, substr string) {
	assert.True(t, strings.Contains(container, substr), "expected to contain `%v`, got `%v`", substr, container)
}
