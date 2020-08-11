package main

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/incognitochain/bridge-eth/bridge/locker"
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
			p, err := setupLockerContract(tc.admin.Address)
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

func setupLockerContract(admin common.Address) (*Platform, error) {
	alloc := make(core.GenesisAlloc)
	balance, _ := big.NewInt(1).SetString("100000000000000000000", 10) // 100 eth
	alloc[auth.From] = core.GenesisAccount{Balance: balance}
	alloc[admin] = core.GenesisAccount{Balance: balance}
	sim := backends.NewSimulatedBackend(alloc, 8000000)

	var err error
	p := &Platform{sim: sim, contracts: &contracts{}}
	p.lockerAddr, _, p.locker, err = locker.DeployLocker(auth, sim, admin, common.Address{})
	if err != nil {
		return nil, fmt.Errorf("failed to deploy Locker contract: %v", err)
	}
	sim.Commit()
	return p, nil
}
