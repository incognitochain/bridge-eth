package bridge

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/incognitochain/bridge-eth/bridge/incognito_proxy"
	"github.com/incognitochain/bridge-eth/bridge/vault"
	"github.com/incognitochain/bridge-eth/bridge/vaultproxy"
	"github.com/incognitochain/bridge-eth/erc20"
	"github.com/incognitochain/bridge-eth/erc20/bnb"
	"github.com/incognitochain/bridge-eth/erc20/dai"
	"github.com/incognitochain/bridge-eth/erc20/dless"
	"github.com/incognitochain/bridge-eth/erc20/fail"
	"github.com/incognitochain/bridge-eth/erc20/usdt"
	"github.com/pkg/errors"
)

// admin
var auth *bind.TransactOpts

// user
var auth2 *bind.TransactOpts

// user
var auth3 *bind.TransactOpts

var genesisAcc *account
var genesisAcc2 *account
var genesisAcc3 *account

type Platform struct {
	*contracts
	sim *backends.SimulatedBackend
}

func init() {
	fmt.Println("Initializing genesis account...")
	genesisAcc = loadAccount()
	var err error
	auth, err = bind.NewKeyedTransactorWithChainID(genesisAcc.PrivateKey, big.NewInt(1337))
	if err != nil {
		panic(err.Error())
	}
	genesisAcc2 = loadAccount()
	auth2, err = bind.NewKeyedTransactorWithChainID(genesisAcc2.PrivateKey, big.NewInt(1337))
	if err != nil {
		panic(err.Error())
	}
	genesisAcc3 = loadAccount()
	auth3, err = bind.NewKeyedTransactorWithChainID(genesisAcc3.PrivateKey, big.NewInt(1337))
	if err != nil {
		panic(err.Error())
	}
}

func TestSimulatedSwapBeacon(t *testing.T) {
	body := getBeaconSwapProof(20)
	if len(body) < 1 {
		t.Fatal(fmt.Errorf("empty beacon swap proof"))
	}

	r := getProofResult{}
	if err := json.Unmarshal([]byte(body), &r); err != nil {
		t.Fatalf("%+v", err)
	}
	if len(r.Result.Instruction) == 0 {
		t.Fatal("invalid swap proof")
	}
	proof, err := decodeProof(&r)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	// a, _ := json.Marshal(proof)
	// fmt.Printf("proof: %s\n", string(a))

	p, err := setupWithHardcodedCommittee()
	// p, err := setupWithLocalCommittee()
	if err != nil {
		t.Fatalf("%+v", err)
	}

	auth.GasLimit = 7000000
	fmt.Printf("inst len: %d\n", len(proof.Instruction))
	tx, err := SwapBeacon(p.inc, auth, proof)
	if err != nil {
		fmt.Println("err:", err)
	}
	p.sim.Commit()
	printReceipt(p.sim, tx)
}

func TestSimulatedBurnETH(t *testing.T) {
	proof, err := getAndDecodeBurnProof("d2e1685eec22f83455980e0a48eafc2a4a03d0b8c5e9aa3671698972fb8b1099")
	if err != nil {
		t.Fatal(err)
	}
	// a, _ := json.Marshal(proof)
	// fmt.Println(string(a))

	p, err := setupWithHardcodedCommittee()
	// p, err := setupWithLocalCommittee()
	if err != nil {
		t.Fatalf("%+v", err)
	}

	oldBalance, newBalance, err := deposit(p, big.NewInt(int64(5e18)))
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("deposit to vault: %d -> %d\n", oldBalance, newBalance)

	withdrawer := common.HexToAddress("0x65033F315F214834BD6A65Dce687Bcb0f32b0a5A")
	fmt.Printf("withdrawer init balance: %d\n", p.getBalance(withdrawer))

	tx, err := Withdraw(p.v, auth2, proof)
	if err != nil {
		fmt.Println("err:", err)
	}
	p.sim.Commit()
	printReceipt(p.sim, tx)

	fmt.Printf("withdrawer new balance: %d\n", p.getBalance(withdrawer))
}

func TestSimulatedBurnERC20(t *testing.T) {
	proof, err := getAndDecodeBurnProof("9d33bdb2d2c614c9576176f5e173515073cdda6c8b970b56ddab506bc3b774d2")
	if err != nil {
		t.Fatal(err)
	}
	// a, _ := json.Marshal(proof)
	// fmt.Println(string(a))

	p, err := setupWithHardcodedCommittee()
	// p, err := setupWithLocalCommittee()
	if err != nil {
		t.Fatalf("%+v", err)
	}

	fmt.Printf("token: %s\n", p.tokenAddr.String())
	oldBalance, newBalance, err := lockSimERC20WithBalance(p, p.token, p.tokenAddr, big.NewInt(int64(1e9)))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("deposit erc20 to vault: %d -> %d\n", oldBalance, newBalance)

	withdrawer := common.HexToAddress("0x65033F315F214834BD6A65Dce687Bcb0f32b0a5A")
	fmt.Printf("withdrawer init balance: %d\n", getBalanceERC20(p.token, withdrawer))

	tx, err := Withdraw(p.v, auth, proof)
	if err != nil {
		fmt.Println("err:", err)
	}
	p.sim.Commit()
	printReceipt(p.sim, tx)

	fmt.Printf("withdrawer new balance: %d\n", getBalanceERC20(p.token, withdrawer))
}

func lockSimERC20WithTxs(
	p *Platform,
	token Tokener,
	tokenAddr common.Address,
	amount *big.Int,
) (*types.Transaction, *types.Transaction, error) {
	txApprove, err := approveERC20(genesisAcc2.PrivateKey, p.vAddr, token, amount, p.sim.Blockchain().Config().ChainID.Int64())
	if err != nil {
		return nil, nil, err
	}
	p.sim.Commit()

	txDeposit, err := depositERC20(genesisAcc2.PrivateKey, p.v, tokenAddr, amount, p.sim.Blockchain().Config().ChainID.Int64())
	if err != nil {
		return txApprove, nil, err
	}
	p.sim.Commit()
	return txApprove, txDeposit, nil
}

func lockSimERC20WithBalance(
	p *Platform,
	token Tokener,
	tokenAddr common.Address,
	amount *big.Int,
) (*big.Int, *big.Int, error) {
	initBalance := getBalanceERC20(token, p.vAddr)
	// fmt.Printf("bal: %d\n", getBalanceERC20(token, genesisAcc.Address))
	if _, _, err := lockSimERC20WithTxs(p, token, tokenAddr, amount); err != nil {
		return nil, nil, err
	}
	newBalance := getBalanceERC20(token, p.vAddr)
	return initBalance, newBalance, nil
}

func getBalanceERC20(token Tokener, addr common.Address) *big.Int {
	bal, err := token.BalanceOf(nil, addr)
	if err != nil {
		return big.NewInt(-1)
	}
	return bal
}

func (p *Platform) getBalance(addr common.Address) *big.Int {
	bal, _ := p.sim.BalanceAt(context.Background(), addr, nil)
	return bal
}

func setup(
	beaconComm []common.Address,
	bridgeComm []common.Address,
	decimals []int,
	accs ...common.Address,
) (*Platform, error) {
	alloc := make(core.GenesisAlloc)
	balance, _ := big.NewInt(1).SetString("1000000000000000000000000000000", 10) // 1E30 wei
	alloc[auth.From] = core.GenesisAccount{Balance: balance}
	alloc[auth2.From] = core.GenesisAccount{Balance: balance}
	alloc[auth3.From] = core.GenesisAccount{Balance: balance}
	for _, acc := range accs {
		alloc[acc] = core.GenesisAccount{Balance: balance}
	}
	sim := backends.NewSimulatedBackend(alloc, 16000000)
	p := &Platform{
		sim: sim,
		contracts: &contracts{
			tokens:       map[int]tokenInfo{},
			customErc20s: map[string]*TokenerInfo{},
		},
	}

	var err error
	var tx *types.Transaction
	_ = tx

	// ERC20: always deploy first so its address is fixed
	p.tokenAddr, tx, p.token, err = erc20.DeployErc20(auth2, sim, "MyErc20", "ERC", big.NewInt(0), big.NewInt(int64(1e18)))
	if err != nil {
		return nil, fmt.Errorf("failed to deploy ERC20 contract: %v", err)
	}
	// fmt.Printf("token addr: %s\n", p.tokenAddr.Hex())
	sim.Commit()

	// Custom tokens
	err = setupCustomTokens(p)
	if err != nil {
		return nil, err
	}

	// Deploy erc20s with different decimals to test
	ercBal := big.NewInt(20)
	ercBal = ercBal.Mul(ercBal, big.NewInt(int64(1e18)))
	ercBal = ercBal.Mul(ercBal, big.NewInt(int64(1e18)))
	for _, d := range decimals {
		tokenAddr, _, token, err := erc20.DeployErc20(auth2, sim, "MyErc20", "ERC", big.NewInt(int64(d)), ercBal)
		if err != nil {
			return nil, fmt.Errorf("failed to deploy ERC20 contract: %v", err)
		}
		sim.Commit()
		p.tokens[d] = tokenInfo{c: token, addr: tokenAddr}
	}

	// IncognitoProxy
	admin := auth.From
	p.incAddr, tx, p.inc, err = incognito_proxy.DeployIncognitoProxy(auth, sim, admin, beaconComm, bridgeComm)
	if err != nil {
		return nil, fmt.Errorf("failed to deploy IncognitoProxy contract: %v", err)
	}
	sim.Commit()
	// fmt.Printf("deployed bridge, addr: %x ", p.incAddr)
	// printReceipt(sim, tx)

	// Vault
	prevVault := common.Address{}
	addr, _, v, err := setupVault(auth, sim, admin, p.incAddr, prevVault)
	if err != nil {
		return nil, err
	}
	p.vAddr = addr
	p.v = v
	vp, _ := vaultproxy.NewTransparentUpgradeableProxy(addr, sim)
	p.vp = vp
	// fmt.Printf("deployed vault, addr: %x ", p.vAddr)
	// printReceipt(sim, tx)

	return p, nil
}

func setupVault(
	auth *bind.TransactOpts,
	backend *backends.SimulatedBackend,
	admin, incAddr, prevVault common.Address,
) (common.Address, *types.Transaction, *vault.Vault, error) {
	addr, _, _, err := vault.DeployVault(auth, backend)
	if err != nil {
		return common.Address{}, nil, nil, fmt.Errorf("failed to deploy Vault contract: %v", err)
	}
	backend.Commit()

	vaultAbi, _ := abi.JSON(strings.NewReader(vault.VaultABI))
	input, _ := vaultAbi.Pack("initialize", prevVault)

	proxyAddr, tx, _, err := vaultproxy.DeployTransparentUpgradeableProxy(auth, backend, addr, admin, incAddr, input)
	if err != nil {
		return common.Address{}, nil, nil, fmt.Errorf("failed to deploy Vault Proxy contract: %v", err)
	}
	backend.Commit()

	v, err := vault.NewVault(proxyAddr, backend)
	if err != nil {
		return common.Address{}, nil, nil, fmt.Errorf("failed create Vault instance: %v", err)
	}

	return proxyAddr, tx, v, nil
}

func setupCustomTokens(p *Platform) error {
	// Deploy BNB
	bal, _ := big.NewInt(1).SetString("200000000000000000000000000", 10)
	addr, _, bnb, err := bnb.DeployBnb(auth2, p.sim, bal, "BNB", uint8(18), "BNB")
	if err != nil {
		return errors.Errorf("failed to deploy BNB contract: %v", err)
	}
	p.sim.Commit()
	p.contracts.customErc20s["BNB"] = &TokenerInfo{addr: addr, c: bnb}

	// Deploy USDT
	bal, _ = big.NewInt(1).SetString("100000000000", 10)
	addr, _, usdt, err := usdt.DeployUsdt(auth2, p.sim, bal, "Tether USD", "USDT", big.NewInt(6))
	if err != nil {
		return errors.Errorf("failed to deploy USDT contract: %v", err)
	}
	p.sim.Commit()
	p.contracts.customErc20s["USDT"] = &TokenerInfo{addr: addr, c: usdt}

	// Deploy DAI
	symbol := [32]byte{'D', 'A', 'I'}
	addr, _, d, err := dai.DeployDai(auth2, p.sim, symbol)
	if err != nil {
		return errors.Errorf("failed to deploy DAI contract: %v", err)
	}
	p.sim.Commit()
	p.contracts.customErc20s["DAI"] = &TokenerInfo{addr: addr, c: d}

	// Mint DAI
	bal, _ = big.NewInt(1).SetString("1000000000000000000000000000", 10)
	_, err = d.Mint(auth2, bal)
	if err != nil {
		return errors.Errorf("failed to mint DAI: %v", err)
	}
	p.sim.Commit()

	// // Deploy USDC
	// // symbol := [32]byte{'D', 'A', 'I'}
	// addr, _, dc, err := usdc.DeployUsdc(auth, p.sim)
	// if err != nil {
	// 	fmt.Println("ASDASD", err)
	// 	return errors.Errorf("failed to deploy USDC contract: %v", err)
	// }
	// p.sim.Commit()
	// p.contracts.customErc20s["USDC"] = &TokenerInfo{c: dc}

	// // Deploy USDC wrapper
	// addr, _, _, err = usdc_wrap.DeployUsdcWrap(auth, p.sim, addr)
	// if err != nil {
	// 	fmt.Println("!@(*#&!@*(#&", err)
	// 	return errors.Errorf("failed to deploy USDCWrap contract: %v", err)
	// }
	// p.sim.Commit()
	// p.contracts.customErc20s["USDC"].addr = addr

	// Deploy FAIL token
	bal, _ = big.NewInt(1).SetString("1000000000000000000", 10)
	addr, _, fail, err := fail.DeployFAIL(auth2, p.sim, bal, "FAIL", 6, "FAIL")
	if err != nil {
		return errors.Errorf("failed to deploy FAIL contract: %v", err)
	}
	p.sim.Commit()
	p.contracts.customErc20s["FAIL"] = &TokenerInfo{addr: addr, c: fail}

	// Deploy DLESS token
	bal, _ = big.NewInt(1).SetString("1000000000000000000", 10)
	addr, _, dless, err := dless.DeployDLESS(auth2, p.sim, bal, "DLESS", "DLESS")
	if err != nil {
		return errors.Errorf("failed to deploy DLESS contract: %v", err)
	}
	p.sim.Commit()
	p.contracts.customErc20s["DLESS"] = &TokenerInfo{addr: addr, c: dless}

	return nil
}

func setupWithLocalCommittee() (*Platform, error) {
	url := "http://127.0.0.1:9334"
	beaconOld, bridgeOld, err := getCommittee(url)
	if err != nil {
		return nil, err
	}
	return setup(beaconOld, bridgeOld, []int{})
}

func setupWithHardcodedCommittee() (*Platform, error) {
	cmtee := getCommitteeHardcoded()
	return setup(cmtee.beacons, cmtee.bridges, []int{})
}

type account struct {
	PrivateKey *ecdsa.PrivateKey
	Address    common.Address
}

func loadAccount() *account {
	key, err := crypto.LoadECDSA("genesisKey.hex")
	if err != nil {
		return newAccount()
	}
	return &account{
		PrivateKey: key,
		Address:    crypto.PubkeyToAddress(key.PublicKey),
	}
}

func newAccount() *account {
	key, _ := crypto.GenerateKey()
	// crypto.SaveECDSA("genesisKey.hex", key)
	return &account{
		PrivateKey: key,
		Address:    crypto.PubkeyToAddress(key.PublicKey),
	}
}

func retrieveEvents(sim *backends.SimulatedBackend, tx *types.Transaction) (*types.Receipt, map[string][]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	receipt, err := sim.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	if len(receipt.Logs) == 0 {
		fmt.Println("empty log")
		return receipt, nil, nil
	}

	events := map[string][]byte{}
	for _, log := range receipt.Logs {
		events[log.Topics[0].Hex()] = log.Data
	}
	return receipt, events, nil
}

func printReceipt(sim *backends.SimulatedBackend, tx *types.Transaction) {
	receipt, events, err := retrieveEvents(sim, tx)
	if err != nil {
		fmt.Printf("receipt err: %+v", err)
		return
	}

	fmt.Printf("tx gas used: %v\n", receipt.CumulativeGasUsed)

	for topic, data := range events {
		var d interface{}
		d = data

		format := "%+v"
		switch topic {
		case "0x8b1126c8e4087477c3efd9e3785935b29c778491c70e249de774345f7ca9b7f9", "0xa95e6e2a182411e7a6f9ed114a85c3761d87f9b8f453d842c71235aa64fff99f": // bytes32
			format = "%s"
		case "0xb42152598f9b870207037767fd41b627a327c9434c796b2ee501d68acec68d1b", "0x009fd52f05c0ded31d6fb0ee580b923f85e99cf1a5a1da342f25e73c45829c83":
			format = "%x"
		case "0x6c8f06ff564112a969115be5f33d4a0f87ba918c9c9bc3090fe631968e818be4": // bool
			format = "%t"
			d = data[len(data)-1] > 0
		case "0x8e2fc7b10a4f77a18c553db9a8f8c24d9e379da2557cb61ad4cc513a2f992cbd", "0x0ac68d08c5119b8cdb4058edbf0d4168f208ec3935d26a8f1f0d92eb9d4de8bf": // uint
			format = "%s"
			d = big.NewInt(int64(0)).SetBytes(data)
		case "0x0ac6e167e94338a282ec23bdd86f338fc787bd67f48b3ade098144aac3fcd86e", "0xb123f68b8ba02b447d91a6629e121111b7dd6061ff418a60139c8bf00522a284": // address
			format = "%x"
			d = data[12:]
		}

		fmt.Printf(fmt.Sprintf("logs: %s\n", format), d)
		// fmt.Println(topic)
	}
}

func getAndDecodeBridgeSwapProof(url string, block int) (*DecodedProof, error) {
	body := getBridgeSwapProof(url, block)
	if len(body) < 1 {
		return nil, fmt.Errorf("no bridge swap proof found")
	}
	r := getProofResult{}
	if err := json.Unmarshal([]byte(body), &r); err != nil {
		return nil, err
	}
	if len(r.Result.Instruction) == 0 {
		return nil, fmt.Errorf("invalid swap proof")
	}
	proof, err := decodeProof(&r)
	if err != nil {
		return nil, err
	}
	return proof, nil
}

func getBridgeSwapProof(url string, block int) string {
	payload := strings.NewReader(fmt.Sprintf("{\n    \"id\": 1,\n    \"jsonrpc\": \"1.0\",\n    \"method\": \"getbridgeswapproof\",\n    \"params\": [\n    \t%d\n    ]\n}", block))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Host", "127.0.0.1:9338")
	req.Header.Add("accept-encoding", "gzip, deflate")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("err:", err)
		return ""
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(string(body))
	return string(body)
}

func getBeaconSwapProof(block int) string {
	url := "http://127.0.0.1:9344"

	payload := strings.NewReader(fmt.Sprintf("{\n    \"id\": 1,\n    \"jsonrpc\": \"1.0\",\n    \"method\": \"getbeaconswapproof\",\n    \"params\": [\n    \t%d\n    ]\n}", block))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Host", "127.0.0.1:9338")
	req.Header.Add("accept-encoding", "gzip, deflate")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("err:", err)
		return ""
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(string(body))
	return string(body)
}

func deposit(p *Platform, amount *big.Int) (*big.Int, *big.Int, error) {
	initBalance := p.getBalance(p.vAddr)
	auth, err := bind.NewKeyedTransactorWithChainID(genesisAcc2.PrivateKey, p.sim.Blockchain().Config().ChainID)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	auth.GasLimit = 0
	auth.Value = amount
	txId := [32]byte{}
	sign, err := SignDataToShield(txId, genesisAcc2.PrivateKey, genesisAcc2.Address)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	_, err = p.v.Deposit(auth, "", txId, sign)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	p.sim.Commit()
	newBalance := p.getBalance(p.vAddr)
	return initBalance, newBalance, nil
}
