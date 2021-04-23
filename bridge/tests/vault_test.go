package tests

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/incognitochain/bridge-eth/bridge/vault"
	"github.com/incognitochain/bridge-eth/bridge/vaultproxy"
	"github.com/incognitochain/bridge-eth/erc20"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type VaultTestSuite struct {
	suite.Suite
	auth           *bind.TransactOpts
	address        common.Address
	vaultAddress   common.Address
	gAlloc         core.GenesisAlloc
	sim            *backends.SimulatedBackend
	erc20TokenAddr common.Address
}

func TestVaultTestSuite(t *testing.T) {
	suite.Run(t, new(VaultTestSuite))
}

func (s *VaultTestSuite) SetupTest() {
	key, _ := crypto.GenerateKey()
	s.auth = bind.NewKeyedTransactor(key)

	s.address = s.auth.From
	s.gAlloc = map[common.Address]core.GenesisAccount{
		s.address: {Balance: big.NewInt(10000000000)},
	}

	s.sim = backends.NewSimulatedBackend(s.gAlloc, 10000000000)

	// deploy vault contract
	incognitoProxyAddress := common.Address{}
	preVaultAddress := common.Address{}
	admin := common.Address{}
	vaultAddr, deployedTx, _, e := vault.DeployVault(s.auth, s.sim)
	fmt.Println("vault delegator contract deployment gas: ", deployedTx.Gas())

	vaultAbi, _ := abi.JSON(strings.NewReader(vault.VaultABI))
	input, _ := vaultAbi.Pack("initialize", preVaultAddress)
	vaultProxyAddr, deployedTx, _, e := vaultproxy.DeployTransparentUpgradeableProxy(s.auth, s.sim, vaultAddr, admin, incognitoProxyAddress, input)
	fmt.Println("vault proxy contract deployment gas: ", deployedTx.Gas())

	s.vaultAddress = vaultProxyAddr
	s.Nil(e)
	s.sim.Commit()

	// deploy erc20 token
	_, erc20TokenAddr, e := deployErc20(s.sim, s.auth)
	s.erc20TokenAddr = erc20TokenAddr
	s.Nil(e)
}

func (s *VaultTestSuite) TestDeposit() {
	vault, _ := vault.NewVault(s.vaultAddress, s.sim)
	depositedAmt := big.NewInt(3000)
	beforeDepositBal, _ := s.sim.BalanceAt(context.Background(), s.address, nil)
	s.auth.Value = depositedAmt
	tx, _ := vault.Deposit(s.auth, "0xfakeaddress")

	fmt.Println("Deposit Ether gas: ", tx.Gas())
	assert.NotEqual(s.T(), nil, tx)
	s.sim.Commit()

	afterDepositBal, _ := s.sim.BalanceAt(context.Background(), s.address, nil)
	scBal, _ := s.sim.BalanceAt(context.Background(), s.vaultAddress, nil)
	assert.Equal(s.T(), depositedAmt, scBal)
	total := big.NewInt(0)
	total.Add(scBal, big.NewInt(int64(tx.Gas())))
	total.Add(total, afterDepositBal)

	assert.Equal(s.T(), beforeDepositBal, total)
}

func deployErc20(sim *backends.SimulatedBackend, auth *bind.TransactOpts) (*erc20.Erc20, common.Address, error) {
	name := "Super duper erc20"
	symbol := "="
	decimals := big.NewInt(0)
	supply := big.NewInt(1000000000000000000)
	addr, deployedTx, c, err := erc20.DeployErc20(auth, sim, name, symbol, decimals, supply)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("failed to deploy Erc20 contract: %v", err)
	}
	fmt.Println("erc20 deployment gas: ", deployedTx.Gas())
	sim.Commit()
	fmt.Printf("deployed erc20, addr: %x\n", addr)
	erc20Bal, _ := c.BalanceOf(nil, auth.From)
	fmt.Printf("genesis address erc20 balance: %d\n", erc20Bal)
	return c, addr, nil
}

func (s *VaultTestSuite) TestDepositERC20() {
	depositedAmt := big.NewInt(5000)
	vault, _ := vault.NewVault(s.vaultAddress, s.sim)
	erc20Token, _ := erc20.NewErc20(s.erc20TokenAddr, s.sim)

	apprTx, apprErr := erc20Token.Approve(s.auth, s.vaultAddress, depositedAmt)
	if apprErr != nil {
		fmt.Println("Failed to approve erc20 deposit: ", apprErr)
		return
	}
	fmt.Println("ERC20 approval gas: ", apprTx.Gas())
	s.sim.Commit()

	beforeDepositBal, _ := erc20Token.BalanceOf(nil, s.address)
	depoTx, depoErr := vault.DepositERC20(s.auth, s.erc20TokenAddr, depositedAmt, "0xfakeaddress")
	if depoErr != nil {
		fmt.Println("Failed to deposit erc20: ", depoErr)
		return
	}
	fmt.Println("Deposit erc20 token gas: ", depoTx.Gas())
	s.sim.Commit()

	afterDepositBal, _ := erc20Token.BalanceOf(nil, s.address)
	scBal, _ := erc20Token.BalanceOf(nil, s.vaultAddress)
	assert.Equal(s.T(), depositedAmt, scBal)
	total := big.NewInt(0)
	total.Add(afterDepositBal, depositedAmt)
	assert.Equal(s.T(), beforeDepositBal, total)
}
