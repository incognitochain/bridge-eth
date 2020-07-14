package main

import (
	"bytes"
	"context"
	"encoding/json"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	ec "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/incognitochain/bridge-eth/blockchain"
	"github.com/incognitochain/bridge-eth/bridge/nextVault"
	"github.com/incognitochain/bridge-eth/bridge/vault"
	"github.com/incognitochain/bridge-eth/common"
	"github.com/incognitochain/bridge-eth/erc20"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

var DepositERC20Topic = "0x2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e"

func TestFixedFallback(t *testing.T) {
	p, _, err := setupFixedCommittee()
	assert.Nil(t, err)

	vr := vault.VaultRaw{p.v}
	auth.Value = big.NewInt(1000)
	_, err = vr.Transfer(auth)
	if assert.Nil(t, err) {
		p.sim.Commit()
		assert.Equal(t, big.NewInt(1000), p.getBalance(p.vAddr))
	}
	auth.Value = nil
}

func TestFixedUpdateIncognitoProxy(t *testing.T) {
	acc := newAccount()
	testCases := []struct {
		desc   string
		caller *account
		paused bool
		err    bool
	}{
		{
			desc:   "Success",
			caller: genesisAcc,
			paused: true,
		},
		{
			desc:   "Not admin",
			caller: acc,
			paused: true,
			err:    true,
		},
		{
			desc:   "Not paused",
			caller: genesisAcc,
			paused: false,
			err:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			p, _, err := setupFixedCommittee(tc.caller.Address)
			assert.Nil(t, err)

			if tc.paused {
				_, err = p.v.Pause(auth)
				assert.Nil(t, err)
				p.sim.Commit()
			}

			// Update
			txOpt := bind.NewKeyedTransactor(tc.caller.PrivateKey)
			txOpt.GasLimit = 1000000
			tx, err := p.v.UpdateIncognitoProxy(txOpt, acc.Address)
			p.sim.Commit()

			if tc.err {
				// Check error message != nil
				receipt, _ := bind.WaitMined(context.Background(), p.sim, tx)
				reason, _ := errorReason(context.Background(), p.sim, tx, receipt.BlockNumber, tc.caller.Address)
				// fmt.Println(reason)
				assert.True(t, len(reason) > 0)
			} else {
				assert.Nil(t, err)

				// Check new incognito proxy address
				inc, err := p.v.Incognito(nil)
				assert.Nil(t, err)
				assert.Equal(t, inc, acc.Address)
			}
		})
	}
}

// func TestFixedIsWithdrawedFalse(t *testing.T) {
// 	proof := getFixedBurnProofETH()

// 	p, _, err := setupFixedCommittee()
// 	assert.Nil(t, err)

// 	_, _, err = deposit(p, big.NewInt(int64(5e18)))
// 	assert.Nil(t, err)

// 	withdrawer := ec.HexToAddress("0xe722D8b71DCC0152D47D2438556a45D3357d631f")

// 	// First withdraw, must success
// 	_, err = Withdraw(p.v, auth, proof)
// 	assert.Nil(t, err)
// 	p.sim.Commit()
// 	bal := p.getBalance(withdrawer)
// 	assert.Equal(t, bal, big.NewInt(1000000000000))

// 	// Deploy new Vault
// 	prevVault := p.vAddr
// 	p.vAddr, _, p.v, err = vault.DeployVault(auth, p.sim, auth.From, p.incAddr, prevVault)
// 	assert.Nil(t, err)
// 	p.sim.Commit()

// 	// Deposit to new vault
// 	proof = getFixedBurnProofERC20()

// 	oldBalance, newBalance, err := lockSimERC20WithBalance(p, p.token, p.tokenAddr, big.NewInt(int64(1e9)))
// 	assert.Nil(t, err)
// 	assert.Equal(t, oldBalance.Add(oldBalance, big.NewInt(int64(1e9))), newBalance)

// 	// New withdraw, must success
// 	_, err = Withdraw(p.v, auth, proof)
// 	assert.Nil(t, err)
// 	p.sim.Commit()

// 	assert.Equal(t, big.NewInt(2000), getBalanceERC20(p.token, withdrawer))
// }

// func TestFixedIsWithdrawedTrue(t *testing.T) {
// 	proof := getFixedBurnProofETH()

// 	p, _, err := setupFixedCommittee()
// 	assert.Nil(t, err)

// 	_, _, err = deposit(p, big.NewInt(int64(5e18)))
// 	assert.Nil(t, err)

// 	withdrawer := ec.HexToAddress("0xe722D8b71DCC0152D47D2438556a45D3357d631f")

// 	// First withdraw, must success
// 	_, err = Withdraw(p.v, auth, proof)
// 	assert.Nil(t, err)
// 	p.sim.Commit()
// 	bal := p.getBalance(withdrawer)
// 	assert.Equal(t, bal, big.NewInt(1000000000000))

// 	// Deploy new Vault
// 	prevVault := p.vAddr
// 	p.vAddr, _, p.v, err = vault.DeployVault(auth, p.sim, auth.From, p.incAddr, prevVault)
// 	assert.Nil(t, err)
// 	p.sim.Commit()

// 	// Deposit to new vault
// 	_, _, err = deposit(p, big.NewInt(int64(5e18)))
// 	assert.Nil(t, err)

// 	// Withdraw with old proof, must fail
// 	_, err = Withdraw(p.v, auth, proof)
// 	assert.NotNil(t, err)
// 	assert.Equal(t, p.getBalance(withdrawer), big.NewInt(1000000000000))
// }

func TestFixedMoveERC20(t *testing.T) {
	p, _, _ := setupFixedCommittee() // New SimulatedBackend each time => ERC20 address is fixed
	erc20Addr := p.tokenAddr
	type initAsset struct {
		addr  ec.Address
		value int64
	}
	randAcc := newAccount()

	testCases := []struct {
		desc     string
		newVault ec.Address
		assets   []initAsset
		err      bool
	}{
		{
			desc:   "Success",
			assets: []initAsset{initAsset{erc20Addr, 1000}},
		},
		{
			desc:   "One asset failed",
			assets: []initAsset{initAsset{erc20Addr, 1000}, initAsset{randAcc.Address, 0}}, // Dummy address as erc20
			err:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			p, _, err := setupFixedCommittee()
			assert.Nil(t, err)

			// Deploy new vault to have updateAssets method
			newVault, _, _, err := nextVault.DeployNextVault(auth, p.sim, auth.From, p.incAddr, p.vAddr)
			assert.Nil(t, err)

			// Deposit to make sure there's ERC20 to move
			assets := []ec.Address{}
			for _, a := range tc.assets {
				assets = append(assets, a.addr)
				if a.value == 0 {
					continue
				}
				token, _ := erc20.NewErc20(a.addr, p.sim)
				oldBalance, newBalance, err := lockSimERC20WithBalance(p, token, a.addr, big.NewInt(a.value))
				assert.Nil(t, err)
				assert.Equal(t, oldBalance.Add(oldBalance, big.NewInt(a.value)), newBalance)
			}

			// Pause and migrate
			_, err = p.v.Pause(auth)
			assert.Nil(t, err)
			p.sim.Commit()
			_, err = p.v.Migrate(auth, newVault)
			assert.Nil(t, err)
			p.sim.Commit()

			// Move
			_, err = p.v.MoveAssets(auth, assets)
			p.sim.Commit()

			if tc.err {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				for _, a := range tc.assets {
					token, err := erc20.NewErc20(a.addr, p.sim)
					assert.Nil(t, err)
					assert.Equal(t, big.NewInt(a.value), getBalanceERC20(token, newVault))
				}
			}
		})
	}
}

func TestFixedMoveETH(t *testing.T) {
	acc := newAccount()
	testCases := []struct {
		desc     string
		mover    *account
		paused   bool
		newVault bool
		err      bool
	}{
		{
			desc:     "Success",
			mover:    genesisAcc,
			paused:   true,
			newVault: true,
		},
		{
			desc:     "Not admin",
			mover:    acc,
			paused:   true,
			newVault: true,
			err:      true,
		},
		{
			desc:     "Not paused",
			mover:    genesisAcc,
			paused:   false,
			newVault: true,
			err:      true,
		},
		{
			desc:     "Not migrated", // newVault = 0x0
			mover:    genesisAcc,
			paused:   true,
			newVault: false,
			err:      true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			p, _, err := setupFixedCommittee(tc.mover.Address)
			assert.Nil(t, err)

			newVault := ec.Address{}
			if tc.newVault {
				// Deploy new vault to have updateAssets method
				newVault, _, _, err = nextVault.DeployNextVault(auth, p.sim, auth.From, p.incAddr, p.vAddr)
				assert.Nil(t, err)
			}

			// Deposit to make sure there's ETH to move
			oldBalance, newBalance, err := deposit(p, big.NewInt(int64(1000)))
			assert.Nil(t, err)
			assert.Equal(t, oldBalance.Add(oldBalance, big.NewInt(1000)), newBalance)

			// Pause and migrate
			_, err = p.v.Pause(auth)
			assert.Nil(t, err)
			p.sim.Commit()
			if tc.newVault {
				_, err = p.v.Migrate(auth, newVault)
				assert.Nil(t, err)
				p.sim.Commit()
			}

			if !tc.paused {
				_, err = p.v.Unpause(auth)
				assert.Nil(t, err)
				p.sim.Commit()
			}

			// Move
			_, err = p.v.MoveAssets(
				bind.NewKeyedTransactor(tc.mover.PrivateKey),
				[]ec.Address{ec.Address{}},
			)
			p.sim.Commit()

			if tc.err {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, newBalance, p.getBalance(newVault))
			}
		})
	}
}

func TestFixedMigrate(t *testing.T) {
	acc := newAccount()
	testCases := []struct {
		desc     string
		newVault ec.Address
		migrator *account
		paused   bool
		err      bool
	}{
		{
			desc:     "Success",
			newVault: genesisAcc.Address,
			migrator: genesisAcc,
			paused:   true,
		},
		{
			desc:     "Not admin",
			newVault: genesisAcc.Address,
			migrator: acc,
			paused:   true,
			err:      true,
		},
		{
			desc:     "Not paused",
			newVault: genesisAcc.Address,
			migrator: genesisAcc,
			paused:   false,
			err:      true,
		},
		{
			desc:     "Migrate to zero",
			migrator: genesisAcc,
			paused:   true,
			err:      true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			p, _, err := setupFixedCommittee(tc.migrator.Address)
			assert.Nil(t, err)

			if tc.paused {
				_, err = p.v.Pause(auth)
				assert.Nil(t, err)
				p.sim.Commit()
			}

			// Migrate
			_, err = p.v.Migrate(bind.NewKeyedTransactor(tc.migrator.PrivateKey), tc.newVault)
			p.sim.Commit()

			if tc.err {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestFixedDepositETH(t *testing.T) {
	p, _, err := setupFixedCommittee()
	assert.Nil(t, err)

	oldBalance, newBalance, err := deposit(p, big.NewInt(int64(5e18)))
	assert.Nil(t, err)

	assert.Equal(t, oldBalance.Add(oldBalance, big.NewInt(int64(5e18))), newBalance)
}

func TestFixedDepositOverbalanceETH(t *testing.T) {
	p, _, err := setupFixedCommittee()
	assert.Nil(t, err)

	// First deposit, success
	amount := big.NewInt(1).Exp(big.NewInt(10), big.NewInt(27), nil)
	oldBalance, newBalance, err := deposit(p, amount)
	if assert.Nil(t, err) {
		assert.Equal(t, oldBalance.Add(oldBalance, amount), newBalance)
	}

	// Second deposit, overbalance, fail
	oldBalance, newBalance, err = deposit(p, big.NewInt(1))
	if assert.NotNil(t, err) {
		assert.Equal(t, oldBalance, newBalance)
	}
}

func TestFixedDepositETHPaused(t *testing.T) {
	p, _, err := setupFixedCommittee()
	assert.Nil(t, err)

	// Pause first
	_, err = p.v.Pause(auth)
	assert.Nil(t, err)
	p.sim.Commit()

	oldBalance, newBalance, err := deposit(p, big.NewInt(int64(5e18)))
	assert.NotNil(t, err)

	assert.Equal(t, oldBalance, newBalance)
}

func TestFixedDepositERC20Once(t *testing.T) {
	p, _, err := setupFixedCommittee()
	assert.Nil(t, err)

	oldBalance, newBalance, err := lockSimERC20WithBalance(p, p.token, p.tokenAddr, big.NewInt(int64(1e9)))
	assert.Nil(t, err)

	assert.Equal(t, oldBalance.Add(oldBalance, big.NewInt(int64(1e9))), newBalance)
}

func TestFixedDepositERC20Decimals(t *testing.T) {
	b2e27, _ := big.NewInt(1).SetString("2000000000000000000000000000", 10)
	testCases := []struct {
		desc    string
		decimal int
		value   *big.Int
		emit    *big.Int
		err     bool
	}{
		{
			desc:    "DAI (d=18)",
			decimal: 18,
			value:   big.NewInt(int64(5e18)),
			emit:    big.NewInt(int64(5e9)),
		},
		{
			desc:    "ZIL (d=12)",
			decimal: 12,
			value:   big.NewInt(int64(3e12)),
			emit:    big.NewInt(int64(3e9)),
		},
		{
			desc:    "ABC (d=27)",
			decimal: 27,
			value:   b2e27,
			emit:    big.NewInt(int64(2e9)),
		},
		{
			desc:    "XYZ (d=9)",
			decimal: 9,
			value:   big.NewInt(int64(4e9)),
			emit:    big.NewInt(int64(4e9)),
		},
		{
			desc:    "USDT (d=6)",
			decimal: 6,
			value:   big.NewInt(int64(8e6)),
			emit:    big.NewInt(int64(8e6)),
		},
		{
			desc:    "IJK (d=0)",
			decimal: 0,
			value:   big.NewInt(9),
			emit:    big.NewInt(9),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			decimals := []int{tc.decimal}
			p, _, err := setupFixedERC20s(decimals)
			assert.Nil(t, err)

			tinfo := p.tokens[tc.decimal]
			oldBalance := getBalanceERC20(tinfo.c, p.vAddr)
			_, tx, err := lockSimERC20WithTxs(p, tinfo.c, tinfo.addr, tc.value)
			if assert.Nil(t, err) {
				newBalance := getBalanceERC20(tinfo.c, p.vAddr)
				assert.Equal(t, oldBalance.Add(oldBalance, tc.value), newBalance)

				emitted, err := extractAmountInDepositERC20Event(p.sim, tx)
				if assert.Nil(t, err) {
					assert.Equal(t, tc.emit, emitted)
				}
			}
		})
	}
}

func extractAmountInDepositERC20Event(sim *backends.SimulatedBackend, tx *types.Transaction) (*big.Int, error) {
	_, events, err := retrieveEvents(sim, tx)
	if err != nil {
		return nil, err
	}
	data, ok := events[DepositERC20Topic]
	if !ok {
		return nil, errors.Errorf("no erc20 deposit event found in tx %v", tx.Hash().Hex())
	}
	cAbi, err := abi.JSON(strings.NewReader(vault.VaultABI))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	e := struct {
		Token            ec.Address
		IncognitoAddress string
		Amount           *big.Int
	}{}
	// fmt.Printf("%+v\n", cAbi.Events["Deposit"])
	err = cAbi.Unpack(&e, "Deposit", data)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return e.Amount, nil
}

func TestFixedDepositOverbalanceERC20(t *testing.T) {
	p, _, err := setupFixedCommittee()
	assert.Nil(t, err)

	// First deposit, success
	amount := int64(1e18)
	oldBalance, newBalance, err := lockSimERC20WithBalance(p, p.token, p.tokenAddr, big.NewInt(amount))
	if assert.Nil(t, err) {
		assert.Equal(t, oldBalance.Add(oldBalance, big.NewInt(amount)), newBalance)
	}

	// Second deposit, overbalance, fail
	oldBalance, newBalance, err = lockSimERC20WithBalance(p, p.token, p.tokenAddr, big.NewInt(1))
	if assert.NotNil(t, err) {
		assert.Equal(t, oldBalance, newBalance)
	}
}

func TestFixedDepositERC20Paused(t *testing.T) {
	p, _, err := setupFixedCommittee()
	assert.Nil(t, err)

	// Pause first
	_, err = p.v.Pause(auth)
	assert.Nil(t, err)
	p.sim.Commit()

	oldBalance, newBalance, err := lockSimERC20WithBalance(p, p.token, p.tokenAddr, big.NewInt(int64(1e9)))
	assert.NotNil(t, err)

	assert.Equal(t, oldBalance, newBalance)
}

func TestFixedDepositCustomERC20s(t *testing.T) {
	testCases := []struct {
		desc    string
		decimal uint8
		value   *big.Int
		emit    *big.Int
		err     bool
	}{
		{
			desc:    "USDT",
			decimal: 6,
			value:   big.NewInt(int64(5e8)),
			emit:    big.NewInt(int64(5e8)),
		},
		{
			desc:    "BNB",
			decimal: 18,
			value:   big.NewInt(int64(3e18)),
			emit:    big.NewInt(int64(3e9)),
		},
		{
			desc:    "DAI",
			decimal: 18,
			value:   big.NewInt(int64(8e15)),
			emit:    big.NewInt(int64(8e6)),
		},
		{
			desc:    "FAIL",
			decimal: 6,
			value:   big.NewInt(int64(1e9)),
			err:     true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			p, _, err := setupFixedCommittee()
			assert.Nil(t, err)

			// Check decimal
			c := p.customErc20s[tc.desc].c
			addr := p.customErc20s[tc.desc].addr
			decimal, err := p.v.GetDecimals(nil, addr)
			if assert.Nil(t, err) {
				assert.Equal(t, tc.decimal, decimal)
			}

			oldBalance := getBalanceERC20(c, p.vAddr)
			_, tx, err := lockSimERC20WithTxs(p, c, addr, tc.value)
			if !tc.err {
				if assert.Nil(t, err) {
					newBalance := getBalanceERC20(c, p.vAddr)
					assert.Equal(t, oldBalance.Add(oldBalance, tc.value), newBalance)

					emitted, err := extractAmountInDepositERC20Event(p.sim, tx)
					if assert.Nil(t, err) {
						assert.Equal(t, tc.emit, emitted)
					}
				}
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

// func TestFixedWithdrawAfterSwap(t *testing.T) {
// 	p, _, err := setupFixedCommittee()
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	burnProofs := getFixedBurnProofAfterSwap()
// 	swapProofs := getFixedSwapProofsToBurn()

// 	oldBalance, newBalance, err := deposit(p, big.NewInt(int64(5e18)))
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	fmt.Printf("deposit to vault: %d -> %d\n", oldBalance, newBalance)

// 	swapOnce(t, p, swapProofs[0])
// 	bal := withdrawOnce(t, p, burnProofs[0])
// 	if bal.Uint64() != uint64(1000000000000) {
// 		t.Fatalf("incorrect balance after withdrawing, expect %d, got %d", 1000000000000, bal)
// 	}
// 	swapOnce(t, p, swapProofs[1])
// 	bal = withdrawOnce(t, p, burnProofs[1])
// 	if bal.Uint64() != uint64(3000000000000) {
// 		t.Fatalf("incorrect balance after withdrawing, expect %d, got %d", 3000000000000, bal)
// 	}
// }

// func swapOnce(t *testing.T, p *Platform, swapProof *decodedProof) {
// 	tx, err := SwapBridge(p.inc, auth, swapProof)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	p.sim.Commit()
// 	printReceipt(p.sim, tx)
// }

// func withdrawOnce(t *testing.T, p *Platform, burnProof *decodedProof) *big.Int {
// 	withdrawer := ec.HexToAddress("0xe722D8b71DCC0152D47D2438556a45D3357d631f")
// 	fmt.Printf("withdrawer init balance: %d\n", p.getBalance(withdrawer))

// 	tx, err := Withdraw(p.v, auth, burnProof)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	p.sim.Commit()
// 	printReceipt(p.sim, tx)

// 	fmt.Printf("withdrawer new balance: %d\n", p.getBalance(withdrawer))
// 	return p.getBalance(withdrawer)
// }

func TestFixedParseBurnInst(t *testing.T) {
	token := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 3}
	to := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 5, 6}
	amount := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 7, 8, 9}
	in := &burnInst{
		meta:   72,
		shard:  1,
		token:  ec.BytesToAddress(token),
		to:     ec.BytesToAddress(to),
		amount: big.NewInt(0).SetBytes(amount),
	}
	data := []byte{in.meta, in.shard}
	data = append(data, token[:]...)
	data = append(data, to[:]...)
	data = append(data, amount[:]...)

	p, _, err := setupFixedCommittee()
	if err != nil {
		t.Error(err)
	}
	resMeta, resShard, resToken, resTo, resAmount, err := p.v.ParseBurnInst(nil, data)
	out := &burnInst{
		meta:   resMeta,
		shard:  resShard,
		token:  resToken,
		to:     resTo,
		amount: resAmount,
	}
	if err != nil {
		t.Error(err)
	}
	checkBurnInst(t, in, out)
}

func checkBurnInst(t *testing.T, in, out *burnInst) {
	if in.meta != out.meta {
		t.Error(errors.Errorf("incorrect meta: expect %x, got %x", out.meta, in.meta))
	}
	if in.shard != out.shard {
		t.Error(errors.Errorf("incorrect shard: expect %x, got %x", out.shard, in.shard))
	}
	if !bytes.Equal(in.token[:], out.token[:]) {
		t.Error(errors.Errorf("incorrect token: expect %x, got %x", out.token, in.token))
	}
	if !bytes.Equal(in.to[:], out.to[:]) {
		t.Error(errors.Errorf("incorrect to: expect %x, got %x", out.to, in.to))
	}
	if in.amount.Cmp(out.amount) != 0 {
		t.Error(errors.Errorf("incorrect amount: expect %x, got %x", out.amount, in.amount))
	}
}

type burnInst struct {
	meta   uint8
	shard  uint8
	token  ec.Address
	to     ec.Address
	amount *big.Int
}

// func TestFixedWithdrawTwice(t *testing.T) {
// 	proof := getFixedBurnProofETH()

// 	p, _, err := setupFixedCommittee()
// 	assert.Nil(t, err)

// 	_, _, err = deposit(p, big.NewInt(int64(5e18)))
// 	assert.Nil(t, err)

// 	withdrawer := ec.HexToAddress("0xe722D8b71DCC0152D47D2438556a45D3357d631f")

// 	// First withdraw, must success
// 	_, err = Withdraw(p.v, auth, proof)
// 	assert.Nil(t, err)
// 	p.sim.Commit()
// 	bal := p.getBalance(withdrawer)
// 	assert.Equal(t, bal, big.NewInt(1000000000000))

// 	// Second withdraw, must fail
// 	_, err = Withdraw(p.v, auth, proof)
// 	assert.NotNil(t, err)
// }

// func TestFixedWithdrawETH(t *testing.T) {
// 	proof := getFixedBurnProofETH()

// 	p, _, err := setupFixedCommittee()
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	oldBalance, newBalance, err := deposit(p, big.NewInt(int64(5e18)))
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	fmt.Printf("deposit to vault: %d -> %d\n", oldBalance, newBalance)

// 	withdrawer := ec.HexToAddress("0xe722D8b71DCC0152D47D2438556a45D3357d631f")
// 	fmt.Printf("withdrawer init balance: %d\n", p.getBalance(withdrawer))

// 	tx, err := Withdraw(p.v, auth, proof)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	p.sim.Commit()
// 	printReceipt(p.sim, tx)

// 	bal := p.getBalance(withdrawer)
// 	fmt.Printf("withdrawer new balance: %d\n", bal)
// 	if bal.Int64() != int64(1000000000000) {
// 		t.Fatalf("incorrect balance after withdrawing, expect %d, got %d", 1000000000, bal)
// 	}
// }

// func TestFixedWithdrawPaused(t *testing.T) {
// 	proof := getFixedBurnProofETH()

// 	p, _, err := setupFixedCommittee()
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	oldBalance, newBalance, err := deposit(p, big.NewInt(int64(5e18)))
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	fmt.Printf("deposit to vault: %d -> %d\n", oldBalance, newBalance)

// 	// Pause first
// 	_, err = p.v.Pause(auth)
// 	assert.Nil(t, err)

// 	withdrawer := ec.HexToAddress("0xe722D8b71DCC0152D47D2438556a45D3357d631f")
// 	_, err = Withdraw(p.v, auth, proof)
// 	assert.NotNil(t, err)

// 	bal := p.getBalance(withdrawer)
// 	assert.Zero(t, bal.Int64())
// }

func TestFixedWithdrawERC20Decimals(t *testing.T) {
	b2e27, _ := big.NewInt(1).SetString("2000000000000000000000000000", 10)
	testCases := []struct {
		desc     string
		decimal  int
		deposit  *big.Int
		withdraw *big.Int
		remain   *big.Int
		err      bool
	}{
		{
			desc:     "DAI (d=18)",
			decimal:  18,
			deposit:  big.NewInt(int64(5e18)),
			withdraw: big.NewInt(int64(4e9)),
			remain:   big.NewInt(int64(1e18)),
		},
		{
			desc:     "ZIL (d=12)",
			decimal:  12,
			deposit:  big.NewInt(int64(3e12)),
			withdraw: big.NewInt(int64(3e8)),
			remain:   big.NewInt(int64(2.7e12)),
		},
		{
			desc:     "ABC (d=27)",
			decimal:  27,
			deposit:  b2e27,
			withdraw: big.NewInt(int64(2e9)),
			remain:   big.NewInt(int64(0)),
		},
		{
			desc:     "XYZ (d=9)",
			decimal:  9,
			deposit:  big.NewInt(int64(4e9)),
			withdraw: big.NewInt(int64(1)),
			remain:   big.NewInt(int64(4e9 - 1)),
		},
		{
			desc:     "USDT (d=6)",
			decimal:  6,
			deposit:  big.NewInt(int64(8e6)),
			withdraw: big.NewInt(int64(7e6)),
			remain:   big.NewInt(int64(1e6)),
		},
		{
			desc:     "IJK (d=0)",
			decimal:  0,
			deposit:  big.NewInt(9),
			withdraw: big.NewInt(2),
			remain:   big.NewInt(7),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			decimals := []int{tc.decimal}
			p, c, err := setupFixedERC20s(decimals)
			assert.Nil(t, err)
			tinfo := p.tokens[tc.decimal]

			// Deposit, must success
			_, _, err = lockSimERC20WithTxs(p, tinfo.c, tinfo.addr, tc.deposit)
			assert.Nil(t, err)

			meta := 72
			shardID := 1
			data := buildWithdrawTestcase(c, meta, shardID, tinfo.addr, tc.withdraw)

			// The blocks with withdraw instructions are final
			assert.Nil(t, setupFinalityWithBlocks(p, true, c, data.beaconHashes))
			assert.Nil(t, setupFinalityWithBlocks(p, false, c, data.bridgeHashes))

			_, err = WithdrawNew(p.v, auth, data.inst, data.heights, data.minedProofs, data.ancestorProofs)
			if assert.Nil(t, err) {
				p.sim.Commit()

				// Check balance
				bal := getBalanceERC20(tinfo.c, p.vAddr)
				assert.Zero(t, tc.remain.Cmp(bal))
			}
		})
	}

}

// func TestFixedWithdrawERC20(t *testing.T) {
// 	proof := getFixedBurnProofERC20()

// 	p, _, err := setupFixedCommittee()
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	oldBalance, newBalance, err := lockSimERC20WithBalance(p, p.token, p.tokenAddr, big.NewInt(int64(1e9)))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	fmt.Printf("deposit erc20 to vault: %d -> %d\n", oldBalance, newBalance)

// 	withdrawer := ec.HexToAddress("0xe722D8b71DCC0152D47D2438556a45D3357d631f")
// 	fmt.Printf("withdrawer init balance: %d\n", getBalanceERC20(p.token, withdrawer))

// 	auth.GasLimit = 8000000
// 	tx, err := Withdraw(p.v, auth, proof)
// 	if err != nil {
// 		fmt.Println("err:", err)
// 	}
// 	p.sim.Commit()
// 	printReceipt(p.sim, tx)

// 	bal := getBalanceERC20(p.token, withdrawer)
// 	fmt.Printf("withdrawer new balance: %d\n", bal)
// 	if bal.Int64() != int64(2000) {
// 		t.Fatalf("incorrect balance after withdrawing, expect %d, got %d", 2000, bal)
// 	}
// }

func TestFixedWithdrawCustomERC20s(t *testing.T) {
	testCases := []struct {
		desc     string
		deposit  *big.Int
		withdraw *big.Int
		remain   *big.Int
		err      bool
	}{
		{
			desc:     "USDT",
			deposit:  big.NewInt(int64(5e8)),
			withdraw: big.NewInt(int64(4e8)),
			remain:   big.NewInt(int64(1e8)),
		},
		{
			desc:     "BNB",
			deposit:  big.NewInt(int64(3e18)),
			withdraw: big.NewInt(int64(3e9)),
			remain:   big.NewInt(int64(0)),
		},
		{
			desc:     "DAI",
			deposit:  big.NewInt(int64(3e17)),
			withdraw: big.NewInt(int64(2e8)),
			remain:   big.NewInt(int64(1e17)),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			p, c, err := setupFixedCommittee()
			assert.Nil(t, err)

			// Deposit, must success
			tinfo := p.customErc20s[tc.desc]
			_, _, err = lockSimERC20WithTxs(p, tinfo.c, tinfo.addr, tc.deposit)
			assert.Nil(t, err)

			meta := 72
			shardID := 1
			data := buildWithdrawTestcase(c, meta, shardID, tinfo.addr, tc.withdraw)

			// The blocks with withdraw instructions are final
			assert.Nil(t, setupFinalityWithBlocks(p, true, c, data.beaconHashes))
			assert.Nil(t, setupFinalityWithBlocks(p, false, c, data.bridgeHashes))

			auth.GasLimit = 0
			_, err = WithdrawNew(p.v, auth, data.inst, data.heights, data.minedProofs, data.ancestorProofs)
			if assert.Nil(t, err) {
				p.sim.Commit()

				// Check balance
				bal := getBalanceERC20(tinfo.c, p.vAddr)
				assert.Zero(t, tc.remain.Cmp(bal))
			}
		})
	}

}

func TestFixedWithdrawModifiedProof(t *testing.T) {
	p, c, err := setupFixedCommittee()
	desc := "DAI"
	deposit := big.NewInt(int64(3e17))
	withdraw := big.NewInt(int64(1e8))
	// Deposit, must success
	tinfo := p.customErc20s[desc]
	_, _, err = lockSimERC20WithTxs(p, tinfo.c, tinfo.addr, deposit)
	assert.Nil(t, err)

	// wrong meta
	meta := 98
	shardID := 1
	data := buildWithdrawTestcase(c, meta, shardID, tinfo.addr, withdraw)

	// The blocks with withdraw instructions are final
	assert.Nil(t, setupFinalityWithBlocks(p, true, c, data.beaconHashes))
	assert.Nil(t, setupFinalityWithBlocks(p, false, c, data.bridgeHashes))

	auth.GasLimit = 0
	_, err = WithdrawNew(p.v, auth, data.inst, data.heights, data.minedProofs, data.ancestorProofs)
	assert.NotNil(t, err)
	p.sim.Commit()

	// wrong shard
	meta = 97
	shardID = 2
	auth.GasLimit = 0
	_, err = WithdrawNew(p.v, auth, data.inst, data.heights, data.minedProofs, data.ancestorProofs)
	assert.NotNil(t, err)
	p.sim.Commit()
}

func setupFixedCommittee(accs ...ec.Address) (*Platform, *committees, error) {
	c := getFixedCommittee()
	p, err := setup(c.beacons, c.bridges, []int{}, accs...)
	return p, c, err
}

func setupFixedERC20s(decimals []int) (*Platform, *committees, error) {
	c := getFixedCommittee()
	p, err := setup(c.beacons, c.bridges, decimals)
	return p, c, err
}

func setupFinalityWithBlocks(p *Platform, isBeacon bool, c *committees, blkHashes [][]byte) error {
	root := blockchain.GetKeccak256MerkleRoot(blkHashes)
	finalityInsts, finalityProofs := buildFinalityTestcase(c, isBeacon, []byte{73, 73}, [][]byte{root, root}, []int{123456780, 123456790})
	_, err := p.inc.SubmitFinalityProof(auth, finalityInsts, finalityProofs, big.NewInt(int64(0)), isBeacon)
	if err == nil {
		p.sim.Commit()
	}
	return err
}

type withdrawTestcase struct {
	inst           []byte
	heights        [2]*big.Int
	minedProofs    [2]vault.VaultMinedProof
	ancestorProofs [2]vault.VaultMerkleProof
	beaconHashes   [][]byte
	bridgeHashes   [][]byte
}

func buildWithdrawTestcase(
	c *committees,
	meta, shard int,
	tokenID ec.Address,
	amount *big.Int,
) withdrawTestcase {
	// Proofs on beacon
	inst, mp, blkData, blkHash, withdrawHeight := buildWithdrawData(meta, shard, tokenID, amount)
	minedBeacon := buildVaultMinedProof(signAndReturnInstProof(c.beaconPrivs, true, mp, blkData, blkHash[:]))
	beaconHashes := randomMerkleHashes(1234)
	beaconHashes[withdrawHeight] = blkHash[:]
	ancestorBeacon := buildVaultAncestorProof(buildInstructionMerklePath(beaconHashes, withdrawHeight))

	// Proofs on bridge
	minedBridge := buildVaultMinedProof(signAndReturnInstProof(c.bridgePrivs, false, mp, blkData, blkHash[:]))
	bridgeHashes := randomMerkleHashes(4567)
	bridgeHashes[withdrawHeight] = blkHash[:] // Use the same block hash
	ancestorBridge := buildVaultAncestorProof(buildInstructionMerklePath(bridgeHashes, withdrawHeight))

	minedProofs := [2]vault.VaultMinedProof{minedBeacon, minedBridge}
	ancestorProofs := [2]vault.VaultMerkleProof{ancestorBeacon, ancestorBridge}
	return withdrawTestcase{
		inst:           inst,
		heights:        [2]*big.Int{big.NewInt(int64(withdrawHeight)), big.NewInt(int64(withdrawHeight))},
		minedProofs:    minedProofs,
		ancestorProofs: ancestorProofs,
		beaconHashes:   beaconHashes,
		bridgeHashes:   bridgeHashes,
	}
}

func buildVaultMinedProof(proof *instProof) vault.VaultMinedProof {
	return vault.VaultMinedProof{
		Path:    proof.instPath,
		IsLeft:  proof.instPathIsLeft,
		Root:    proof.instRoot,
		BlkData: proof.blkData,
	}
}

func buildVaultAncestorProof(proof *merklePath) vault.VaultMerkleProof {
	return vault.VaultMerkleProof{
		Path:   proof.path,
		IsLeft: proof.left,
	}
}

func buildWithdrawData(meta, shard int, tokenID ec.Address, amount *big.Int) ([]byte, *merklePath, []byte, []byte, int) {
	// Build instruction merkle tree
	numInst := 10
	startNodeID := 7
	height := 1
	withdrawer := ec.HexToAddress("0xe722D8b71DCC0152D47D2438556a45D3357d631f")
	inst := buildDecodedWithdrawInst(meta, shard, tokenID, withdrawer, amount)
	instWithHeight := append(inst, toBytes32BigEndian(big.NewInt(int64(height)).Bytes())...)
	data := randomMerkleHashes(numInst)
	data[startNodeID] = instWithHeight
	mp := buildInstructionMerklePath(data, startNodeID)

	// Generate random blkHash
	h := randomMerkleHashes(1)
	blkData := h[0]
	blkHash := common.Keccak256(blkData, mp.root[:])
	return inst, mp, blkData, blkHash[:], height
}

func buildDecodedWithdrawInst(meta, shard int, tokenID, withdrawer ec.Address, amount *big.Int) []byte {
	decoded := []byte{byte(meta)}
	decoded = append(decoded, byte(shard))
	decoded = append(decoded, toBytes32BigEndian(tokenID[:])...)
	decoded = append(decoded, toBytes32BigEndian(withdrawer[:])...)
	decoded = append(decoded, toBytes32BigEndian(amount.Bytes())...)
	decoded = append(decoded, toBytes32BigEndian(make([]byte, 32))...) // txID
	decoded = append(decoded, make([]byte, 16)...)                     // incTokenID, variable length
	return decoded
}

type committees struct {
	beacons     []ec.Address
	bridges     []ec.Address
	beaconPrivs [][]byte
	bridgePrivs [][]byte
}

func getFixedBurnProofETH() *decodedProof {
	proofMarshalled := `{"Instruction":"SAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAOci2LcdzAFS1H0kOFVqRdM1fWMfAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA6NSlEAAQph4uU2jzYJA/wwowaQNkV1YDSZxF3b7zhinToad2CgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","Heights":[13,15],"InstPaths":[[[85,90,80,85,102,77,35,181,79,15,218,226,21,248,182,95,46,161,151,187,156,132,212,157,64,83,229,28,33,225,60,39],[90,120,63,154,71,105,58,60,79,100,14,227,126,162,217,227,120,234,73,47,138,250,88,120,245,244,79,86,210,74,95,107]],[]],"InstPathIsLefts":[[false,false],[]],"InstRoots":[[232,220,37,58,255,80,189,228,5,6,217,133,81,89,207,211,88,46,172,168,207,11,240,31,39,121,213,208,161,87,191,205],[189,164,48,17,53,82,33,194,160,26,125,127,160,224,216,173,239,49,214,141,169,214,192,76,85,66,30,230,107,198,183,59]],"BlkData":[[85,105,22,139,166,247,192,177,35,127,254,184,33,212,88,184,142,32,141,79,0,92,62,215,250,238,97,155,250,103,212,226],[43,69,182,160,171,28,79,48,143,171,86,32,244,96,9,226,26,232,37,254,92,251,3,56,247,81,147,139,97,238,181,249]],"SigIdxs":[[0,2,3],[0,1,2,3]],"SigVs":["HBsb","GxwcHA=="],"SigRs":[[[64,129,127,189,84,17,9,178,133,57,35,78,163,181,179,201,225,45,218,234,194,20,180,7,17,207,92,110,26,73,205,24],[195,54,155,40,185,122,123,141,83,10,38,226,70,217,107,128,172,33,9,9,218,212,254,215,89,20,84,138,200,144,225,255],[46,36,148,21,139,126,174,64,59,164,229,133,213,79,228,67,35,172,234,237,175,125,35,138,2,114,186,151,184,76,242,134]],[[78,44,39,129,74,56,101,155,93,226,145,64,149,247,158,112,186,70,242,82,115,255,142,140,101,38,183,206,197,186,152,138],[119,42,151,249,112,98,99,187,116,121,228,46,126,106,38,166,45,217,156,48,83,103,34,168,246,123,178,82,208,197,9,8],[190,186,249,188,241,169,26,152,54,224,240,95,81,179,223,224,224,81,209,193,57,97,3,165,197,40,229,2,208,20,113,131],[80,12,25,116,195,24,37,40,247,144,182,31,57,127,236,239,190,90,236,59,218,97,205,16,88,102,235,156,220,144,198,234]]],"SigSs":[[[118,52,192,85,239,241,154,66,156,98,190,204,10,253,210,148,162,165,195,201,212,206,135,182,178,209,193,172,240,180,99,110],[61,45,9,108,235,184,112,73,21,154,9,116,162,95,175,207,11,22,215,160,46,209,83,59,118,230,83,209,241,17,168,127],[59,15,150,127,223,9,234,160,9,170,195,109,79,118,137,45,250,156,249,145,226,175,235,60,65,86,182,85,224,162,235,145]],[[101,235,22,155,225,149,55,249,33,205,164,198,58,228,75,33,48,196,171,76,14,181,90,194,239,204,40,159,26,179,46,3],[121,162,14,137,255,11,150,184,16,237,220,66,226,26,240,69,154,30,136,46,68,254,217,190,122,100,41,254,221,122,158,66],[84,26,159,168,124,109,132,10,52,124,177,99,8,53,85,52,194,213,247,52,147,255,17,146,245,114,252,139,176,126,82,203],[39,215,56,10,60,107,131,198,45,193,242,191,111,195,106,238,112,77,177,198,3,88,157,118,76,170,148,65,187,106,118,172]]]}`
	proof := &decodedProof{}
	json.Unmarshal([]byte(proofMarshalled), proof)

	return proof
}

func getFixedBurnProofERC20() *decodedProof {
	proofMarshalled := `{"Instruction":"SAEAAAAAAAAAAAAAAAA+xcl8giIrHqZzLSXw/CizXak4DgAAAAAAAAAAAAAAAOci2LcdzAFS1H0kOFVqRdM1fWMfAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAB9ATWg+enEgVZfeVa3dN1q9UUIeUGdRvAV/yeF06Qe6iXQAAAAAAAAAAAAAAAEfiaoY5HjglF+JXeXDguanESiMB","Heights":[125,127],"InstPaths":[[[17,221,161,166,191,238,13,140,202,175,128,24,145,162,251,223,153,109,242,78,201,65,81,171,124,219,161,20,197,170,219,0],[84,130,181,135,17,33,19,11,216,117,72,234,53,166,213,97,82,240,41,81,39,156,166,254,197,254,129,111,208,88,140,22],[198,35,36,192,5,218,31,215,52,199,187,191,144,207,19,4,228,179,194,5,13,198,213,34,151,75,136,6,173,172,50,134]],[]],"InstPathIsLefts":[[false,false,false],[]],"InstRoots":[[81,96,141,86,165,75,80,243,103,191,248,250,215,87,76,175,0,82,106,87,133,37,47,247,120,243,226,119,211,128,157,135],[9,70,219,61,171,231,41,225,224,130,248,32,95,180,80,179,67,237,159,252,91,57,113,91,222,166,1,243,99,121,54,122]],"BlkData":[[82,27,63,139,96,16,9,27,44,28,16,224,177,26,73,253,215,199,165,214,1,216,35,197,246,35,46,105,196,216,64,173],[172,39,66,38,128,89,242,151,1,14,54,207,142,15,59,110,126,126,195,234,164,117,41,82,150,155,183,4,231,148,53,42]],"SigIdxs":[[0,1,2,3],[0,1,2,3]],"SigVs":["GxwbHA==","GxwbHA=="],"SigRs":[[[6,196,6,82,61,137,152,169,215,249,81,205,46,199,144,13,111,180,67,100,66,170,215,133,234,146,207,26,48,163,121,86],[87,58,198,80,229,65,170,168,50,112,63,182,254,185,229,98,126,66,61,87,101,207,253,224,14,42,160,154,185,6,221,139],[175,225,221,74,177,235,142,81,72,131,11,51,82,58,102,93,253,25,184,45,91,44,225,138,198,114,229,149,210,31,248,98],[83,79,139,76,91,47,237,200,132,49,208,84,126,196,194,90,142,170,239,92,195,194,107,251,141,140,238,198,34,39,111,231]],[[8,49,248,38,235,197,250,7,179,42,42,4,65,128,206,235,148,32,5,145,41,30,229,184,232,38,158,50,133,22,105,104],[3,201,1,66,25,137,91,9,137,46,255,5,212,155,24,232,41,215,229,171,255,132,192,51,52,78,182,144,165,193,62,228],[46,245,113,58,17,213,226,210,86,42,144,229,190,49,65,7,184,173,14,89,6,81,11,235,145,209,31,127,60,156,74,163],[153,65,194,100,24,19,188,17,150,74,179,196,155,50,166,145,30,152,85,17,73,64,144,168,86,185,154,191,254,190,81,79]]],"SigSs":[[[26,192,98,189,190,254,107,245,67,17,148,121,223,22,16,130,126,17,223,243,167,153,85,119,111,176,228,2,182,0,165,166],[110,244,64,23,134,72,133,224,168,45,123,139,249,121,158,187,230,215,194,214,3,31,53,5,237,163,58,155,114,92,203,78],[33,148,100,187,251,227,27,42,194,39,163,207,210,141,89,33,183,188,206,199,129,147,174,64,18,31,29,87,125,248,109,132],[127,66,101,105,146,1,150,63,202,148,228,90,191,0,43,114,23,166,171,113,9,70,95,140,229,84,40,64,47,253,253,201]],[[39,250,172,17,160,109,49,223,126,154,191,233,3,145,215,96,252,82,76,178,212,176,10,193,189,215,191,181,79,39,167,236],[69,121,234,127,246,5,236,92,233,240,149,49,224,212,85,94,50,13,183,7,214,156,216,9,28,159,85,18,207,236,165,22],[96,179,216,187,70,50,194,18,46,253,255,169,118,225,36,52,56,29,19,178,119,149,106,54,187,101,38,134,190,70,229,103],[4,147,185,79,168,172,162,20,197,235,205,10,11,59,52,83,18,206,41,102,231,208,89,64,55,41,78,252,76,12,136,123]]]}`
	proof := &decodedProof{}
	json.Unmarshal([]byte(proofMarshalled), proof)

	return proof
}

func getFixedSwapProofsToBurn() []*decodedProof {
	m := []string{
		`{"Instruction":"RwEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFAAAAAAAAAAAAAAAAPHgSR4Po450eCE/d0OCXM0ui2UUAAAAAAAAAAAAAAAB2402KUnlhKG5VUyYgr1uE88ZTjwAAAAAAAAAAAAAAAGhobbaHRYjSQEFV0Apz+CpQ/dGQAAAAAAAAAAAAAAAAFTOsTSkiwVBVHy9dwrDB7eOCuJAAAAAAAAAAAAAAAAAlm6JOAlKTsmvBKNtKZA2NdyNvBA==","Heights":[0,0],"InstPaths":[[[245,88,13,99,179,59,250,20,46,12,181,171,100,77,39,224,160,202,97,176,128,114,130,243,120,213,30,38,55,222,240,41],[81,62,111,249,189,249,62,211,109,5,111,114,13,10,247,73,18,2,15,184,115,138,45,6,103,94,191,65,17,217,177,0]],[[168,50,153,62,48,142,210,183,222,118,66,243,27,72,27,10,36,27,99,142,207,166,7,170,195,169,161,198,18,120,204,95]]],"InstPathIsLefts":[[false,false],[true]],"InstRoots":[[66,117,186,245,197,16,60,67,38,136,60,223,81,179,86,69,202,243,142,194,238,113,14,92,124,66,85,137,47,32,96,115],[208,183,249,148,42,16,121,249,129,191,146,178,85,213,16,209,57,33,172,123,190,24,192,144,189,184,105,84,173,62,215,120]],"BlkData":[[58,91,247,136,117,57,136,110,62,91,46,223,203,0,200,20,191,19,138,153,231,84,60,70,95,10,127,27,169,163,187,117],[154,189,24,191,205,43,104,36,163,101,27,25,225,117,15,199,69,163,50,111,82,142,12,173,14,3,112,132,246,168,249,184]],"SigIdxs":[[0,1,2,3],[0,1,2,3]],"SigVs":["HBwbHA==","GxwcGw=="],"SigRs":[[[149,90,252,212,61,138,129,182,118,144,75,135,220,43,232,163,160,69,185,248,251,9,93,115,53,252,9,107,90,181,89,247],[59,174,234,98,170,126,227,77,232,217,74,93,119,26,255,122,122,142,53,148,112,168,60,112,191,25,40,131,234,117,46,151],[123,169,251,7,162,18,102,26,210,217,254,17,229,105,177,189,99,62,31,123,146,209,222,224,28,76,42,1,114,162,36,38],[211,86,152,241,169,184,229,242,88,232,152,137,187,157,16,202,114,177,67,93,185,106,209,57,74,239,46,138,54,104,126,214]],[[155,181,233,82,115,19,60,220,94,84,113,4,255,4,30,20,163,216,85,55,199,225,206,52,146,107,155,154,130,213,193,170],[158,173,73,40,10,249,238,117,55,210,48,0,174,21,210,157,20,243,206,195,2,94,54,235,140,31,97,115,114,147,3,100],[254,69,34,39,232,77,57,202,86,120,141,203,166,132,152,137,226,72,86,160,159,169,230,51,16,113,132,169,21,178,61,141],[121,85,221,32,221,217,95,244,83,143,145,158,1,171,234,113,168,254,159,8,60,158,38,167,57,28,249,82,240,231,32,202]]],"SigSs":[[[35,72,66,217,196,218,146,58,88,63,207,152,62,136,200,12,238,70,230,70,47,0,142,34,249,110,163,69,148,176,159,21],[33,21,180,17,81,177,33,187,62,7,48,173,11,116,211,229,94,47,13,153,79,58,93,43,140,21,111,245,74,92,52,22],[2,93,0,59,22,186,250,251,86,20,133,67,213,170,37,62,166,182,61,16,117,136,137,218,189,100,5,66,169,237,147,133],[7,237,7,91,127,140,205,98,73,118,142,42,37,41,26,167,215,123,51,199,141,144,79,161,222,123,76,254,112,40,55,181]],[[99,188,138,66,234,218,10,69,184,81,74,75,251,97,165,227,173,156,57,17,10,120,143,17,151,163,247,118,188,206,218,189],[26,152,153,72,71,143,144,242,217,209,109,185,16,83,146,34,124,102,250,86,233,202,109,175,99,176,238,241,53,134,98,181],[79,252,92,107,188,70,199,131,4,152,34,81,160,11,13,211,123,210,57,219,220,118,220,202,213,159,154,54,195,26,249,116],[38,166,150,148,116,138,180,156,43,169,27,90,14,167,62,196,162,24,85,70,165,129,78,22,187,104,32,149,239,49,234,188]]]}`,
		`{"Instruction":"RwEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAaAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFAAAAAAAAAAAAAAAAduNNilJ5YShuVVMmIK9bhPPGU48AAAAAAAAAAAAAAABoaG22h0WI0kBBVdAKc/gqUP3RkAAAAAAAAAAAAAAAABUzrE0pIsFQVR8vXcKwwe3jgriQAAAAAAAAAAAAAAAAJZuiTgJSk7JrwSjbSmQNjXcjbwQAAAAAAAAAAAAAAAAwHIRBG3slTfCnXoLFBR0aiGfdQg==","Heights":[0,0],"InstPaths":[[[157,63,81,177,76,247,0,3,96,134,191,35,81,234,16,167,215,70,107,151,53,160,143,55,39,32,99,208,230,207,94,152],[106,236,245,106,219,98,91,69,84,61,238,124,163,34,58,216,164,91,206,255,7,107,55,254,17,133,66,225,243,120,131,190]],[[22,183,77,168,123,88,10,139,129,54,23,227,168,45,0,179,7,56,192,166,88,175,50,26,127,239,129,113,241,202,19,182]]],"InstPathIsLefts":[[false,false],[true]],"InstRoots":[[170,40,75,16,34,244,230,138,66,167,28,151,188,131,140,254,37,251,127,79,118,56,52,214,16,124,17,148,47,127,28,70],[237,177,221,7,97,163,98,205,110,152,179,184,183,18,11,3,33,241,178,23,179,232,170,45,231,90,210,121,150,184,61,167]],"BlkData":[[222,153,45,19,12,194,102,218,165,140,140,91,240,66,148,245,117,123,247,218,239,173,202,209,218,23,94,223,74,246,55,148],[192,197,17,69,11,161,70,6,106,1,219,207,73,11,172,222,93,217,117,235,12,186,150,231,119,65,132,133,59,9,83,33]],"SigIdxs":[[0,1,2,3],[0,1,2,3,4]],"SigVs":["HBsbGw==","HBscHBw="],"SigRs":[[[130,176,20,252,105,194,231,35,94,123,22,66,38,121,233,219,198,148,250,124,230,119,165,120,203,200,75,11,96,199,116,105],[174,102,21,28,215,237,86,177,241,229,98,114,221,175,252,192,5,59,14,212,82,109,115,149,129,161,103,61,240,245,126,71],[74,143,140,200,192,31,254,191,252,255,114,13,81,6,114,177,255,2,122,157,199,241,83,74,42,242,89,134,111,84,228,212],[224,104,70,164,239,226,193,128,69,123,24,94,244,91,214,147,125,156,110,48,48,111,30,195,200,238,5,151,32,64,86,85]],[[248,27,87,86,236,205,200,57,40,231,26,25,210,251,211,20,173,66,28,191,193,22,218,30,177,28,173,90,226,169,245,140],[175,59,147,143,156,216,165,174,37,189,236,74,229,13,235,235,229,45,222,166,184,178,171,114,92,109,19,140,80,236,113,168],[37,166,55,217,175,83,168,197,21,28,69,158,105,230,153,113,133,76,210,178,208,20,148,72,137,245,86,153,170,99,131,112],[30,19,65,145,18,7,70,215,14,48,31,85,57,137,159,242,79,73,250,180,99,237,91,80,193,103,0,68,41,165,111,117],[149,125,220,32,178,240,63,88,7,36,98,12,141,73,47,174,80,196,74,119,42,127,233,136,58,121,157,200,51,6,26,133]]],"SigSs":[[[5,167,157,42,157,39,234,228,16,84,38,116,113,77,161,247,225,89,5,89,234,222,132,41,197,106,158,177,46,252,7,155],[39,67,205,63,111,58,253,13,47,125,15,135,247,21,233,174,7,14,101,226,251,98,25,149,78,190,249,151,117,47,60,140],[56,235,154,48,161,68,198,21,231,76,39,29,210,102,29,64,251,208,66,66,121,5,100,253,251,0,201,45,109,108,21,216],[114,177,229,31,43,147,213,108,145,170,210,29,80,44,225,81,30,71,143,145,179,189,176,234,245,198,212,67,220,148,191,231]],[[58,172,173,194,244,235,146,17,7,12,175,170,13,107,223,172,163,181,21,233,225,228,219,3,34,74,83,176,101,16,206,234],[22,14,182,44,199,174,104,220,227,26,222,123,161,231,50,123,246,227,148,246,190,17,131,21,164,149,199,80,60,208,96,32],[50,244,60,98,241,56,251,122,20,177,149,102,110,139,190,192,197,174,132,171,76,151,177,107,57,171,205,58,238,54,180,219],[71,215,4,123,31,39,131,61,71,180,117,6,82,172,48,63,123,126,44,87,136,178,249,109,200,229,175,176,3,130,60,57],[119,100,57,116,141,245,39,243,186,96,32,151,190,201,216,9,120,239,130,188,131,214,254,58,119,213,1,152,107,100,82,90]]]}`,
	}
	proofs := make([]*decodedProof, len(m))
	for i, p := range m {
		proof := &decodedProof{}
		json.Unmarshal([]byte(p), proof)
		proofs[i] = proof
	}
	return proofs
}

func getFixedBurnProofAfterSwap() []*decodedProof {
	m := []string{
		`{"Instruction":"SAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAOci2LcdzAFS1H0kOFVqRdM1fWMfAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA6NSlEADsZ2jiv3iTkZgXoj3VgJEkYN3H071Pj9/2uB2QMoUfnAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","Heights":[68,61],"InstPaths":[[[8,65,138,72,137,30,24,114,197,142,55,170,176,246,231,177,116,29,192,233,180,12,58,144,248,129,168,165,118,168,163,22],[223,80,70,245,216,143,223,177,86,153,255,63,244,137,103,188,162,59,45,141,113,242,52,200,196,185,27,108,233,9,147,13]],[]],"InstPathIsLefts":[[false,false],[]],"InstRoots":[[106,229,28,29,4,209,140,102,54,159,154,169,176,193,182,129,100,85,62,197,58,98,139,192,190,233,12,44,118,95,56,83],[139,202,35,86,234,158,143,225,134,116,19,70,112,181,105,74,195,131,20,230,166,31,109,38,192,160,170,107,51,53,101,120]],"BlkData":[[14,226,168,223,83,228,87,191,70,202,104,120,248,106,30,169,12,56,88,98,101,76,203,13,62,214,47,99,3,102,157,106],[183,183,247,53,39,71,28,2,117,157,15,119,149,45,45,113,201,59,219,147,24,87,109,164,159,242,145,45,93,143,73,143]],"SigIdxs":[[0,1,2],[0,1,2,3,4]],"SigVs":["HBwb","HBscHBs="],"SigRs":[[[61,24,65,50,76,9,146,83,142,80,145,46,84,123,152,97,209,131,105,172,17,29,33,79,215,109,35,114,137,139,66,159],[113,151,138,61,255,74,222,73,213,140,178,11,117,165,255,49,80,241,4,87,113,206,168,108,149,78,7,77,29,240,82,147],[11,72,83,105,30,205,30,196,21,198,171,143,128,34,17,123,41,5,173,132,196,255,119,230,170,195,218,21,14,184,253,34]],[[61,109,51,252,17,240,15,151,232,216,220,234,103,42,249,5,110,31,53,82,32,66,15,132,144,58,156,13,183,51,88,224],[15,142,66,126,113,42,38,234,250,155,202,185,59,105,108,104,112,29,59,105,222,124,19,231,31,120,13,231,13,119,84,14],[98,62,10,228,154,168,172,245,251,230,61,73,118,24,18,177,216,63,1,91,229,248,88,174,15,111,137,134,148,2,159,138],[205,112,250,244,35,101,86,73,12,15,127,61,146,197,22,104,53,6,104,223,29,149,191,111,96,146,110,35,151,144,246,97],[19,61,97,242,155,33,45,95,56,248,148,138,90,63,254,44,86,240,219,196,22,182,21,96,87,227,52,91,238,60,101,4]]],"SigSs":[[[55,21,83,223,148,220,104,138,97,152,208,26,252,223,128,248,73,105,3,40,112,36,68,171,57,180,19,91,165,146,20,141],[15,219,148,117,55,142,146,253,60,74,189,232,194,179,52,218,251,94,170,159,7,169,189,89,75,144,38,161,145,0,27,127],[92,114,32,54,187,135,208,67,73,247,13,112,187,251,182,189,146,94,55,44,57,80,28,6,43,147,70,226,95,138,3,181]],[[73,197,26,120,169,12,199,137,65,251,15,9,26,14,49,127,37,207,171,76,171,130,36,113,254,6,226,132,183,218,25,174],[14,253,179,22,24,170,140,159,22,87,140,21,162,245,127,143,222,108,131,211,113,81,255,241,199,32,107,68,8,254,6,130],[17,251,128,69,193,108,39,110,118,219,249,220,179,45,29,85,148,236,4,229,184,116,130,55,254,154,190,44,67,248,59,155],[31,127,243,109,205,14,122,103,1,223,162,186,67,92,58,69,186,1,0,254,119,222,14,17,221,241,56,247,50,233,77,58],[58,2,46,230,170,155,243,197,200,115,58,233,32,40,108,75,167,102,57,49,116,188,152,8,145,223,213,171,183,37,118,125]]]}`,
		`{"Instruction":"SAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAOci2LcdzAFS1H0kOFVqRdM1fWMfAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAB0alKIACaEK1Sd5T55A8zxwStPiZEi9swxVrsodFvRkQgUM1afwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","Heights":[143,126],"InstPaths":[[[246,128,168,118,179,160,93,6,114,160,32,222,226,67,111,83,253,181,98,174,168,225,12,108,132,67,78,147,30,38,133,179],[197,238,72,100,241,26,189,160,1,232,154,26,160,123,167,203,231,124,16,4,107,200,49,214,51,43,228,200,39,64,14,209]],[]],"InstPathIsLefts":[[false,false],[]],"InstRoots":[[35,212,232,66,202,215,57,160,254,183,83,184,253,157,133,10,136,67,136,236,188,176,93,216,32,73,44,64,229,28,226,78],[166,43,46,98,30,225,108,196,97,127,84,20,30,215,182,91,33,127,179,72,102,120,146,239,254,72,44,40,26,202,151,122]],"BlkData":[[109,203,122,167,33,201,172,153,5,47,133,134,90,89,189,214,6,58,231,5,241,86,21,196,137,90,232,250,87,172,147,65],[129,167,95,76,239,37,214,145,168,238,142,130,109,203,245,168,198,202,242,251,203,77,126,69,179,246,206,201,30,206,179,217]],"SigIdxs":[[0,1,2,3],[0,1,2,3,4]],"SigVs":["GxscHA==","HBwbGxs="],"SigRs":[[[186,170,226,213,84,218,168,158,116,145,213,129,200,255,183,61,193,83,181,242,113,60,17,220,95,133,102,247,98,159,117,151],[60,188,234,32,219,165,102,152,228,98,164,4,154,143,15,190,195,78,174,50,195,44,186,241,167,250,122,133,57,255,110,7],[71,125,189,173,113,146,4,120,16,136,254,253,222,231,187,137,150,247,229,3,175,87,6,230,23,230,6,101,96,140,20,58],[74,180,84,157,145,193,11,54,91,134,151,190,128,164,27,199,66,15,173,213,180,150,51,10,22,254,66,145,94,63,97,83]],[[198,218,89,147,94,59,184,68,136,43,27,71,138,131,73,183,55,222,137,175,32,12,216,202,149,69,153,210,17,111,253,132],[255,106,145,208,250,153,61,29,209,133,44,215,62,191,66,63,81,207,16,146,158,193,92,254,14,177,120,189,137,142,201,245],[156,205,248,131,5,43,2,172,254,97,31,96,157,159,0,185,49,42,165,126,34,169,16,68,190,213,237,152,124,206,233,85],[157,126,233,178,210,111,116,227,154,56,27,211,37,184,70,173,200,252,135,164,111,74,217,231,176,11,40,227,210,33,193,187],[22,198,74,196,37,45,7,144,61,27,217,5,82,95,6,186,151,21,183,253,42,225,81,221,75,189,116,51,12,73,160,73]]],"SigSs":[[[32,193,55,190,251,255,43,244,32,90,218,86,197,24,214,180,43,220,55,31,197,161,190,214,33,166,195,115,99,51,41,178],[67,223,197,60,89,83,98,65,47,190,225,205,48,193,247,110,29,8,21,64,236,218,94,181,151,27,254,161,118,1,10,62],[100,206,2,22,183,9,161,1,224,44,122,244,47,246,117,215,215,238,224,246,57,134,9,97,138,84,45,74,97,183,125,100],[37,181,133,255,247,33,5,136,109,62,114,144,217,86,80,174,41,99,224,219,136,113,229,185,197,185,113,25,87,1,224,12]],[[49,178,67,134,50,110,78,70,152,75,175,53,23,1,162,53,216,224,68,152,16,146,10,190,109,244,215,31,55,1,119,171],[22,183,202,142,133,242,225,91,174,24,43,114,229,203,114,223,104,163,247,122,244,83,30,153,123,109,190,49,43,50,121,21],[40,9,243,182,118,152,112,8,34,110,106,246,242,167,185,8,175,53,167,211,222,210,121,138,39,97,30,163,205,173,7,1],[29,62,51,243,247,251,14,54,91,100,135,83,111,203,246,41,221,46,202,75,19,84,218,253,240,34,57,250,255,64,201,148],[16,39,8,101,163,189,8,183,61,177,169,50,234,124,85,30,182,223,54,177,207,158,135,65,253,28,5,25,15,63,238,175]]]}`,
	}
	proofs := make([]*decodedProof, len(m))
	for i, p := range m {
		proof := &decodedProof{}
		json.Unmarshal([]byte(p), proof)
		proofs[i] = proof
	}
	return proofs
}

func toBytes32BigEndian(b []byte) []byte {
	a := copyToBytes32(b)
	return a[:]
}

func copyToBytes32(b []byte) [32]byte {
	a := [32]byte{}
	copy(a[32-len(b):], b)
	return a
}
