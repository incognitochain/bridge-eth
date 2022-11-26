package bridge

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	ec "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/incognitochain/bridge-eth/bridge/governance"
	"github.com/incognitochain/bridge-eth/bridge/prveth"
	"github.com/incognitochain/bridge-eth/bridge/prvvote"
	"github.com/incognitochain/bridge-eth/bridge/vault"
	"github.com/incognitochain/bridge-eth/bridge/vaultproxy"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"math/big"
	"os/exec"
	"strings"
	"testing"
)

// // Define the suite, and absorb the built-in basic suite
// // functionality from testify - including assertion methods.
type PDaoTestSuite struct {
	suite.Suite
	p                *Platform
	c                *committees
	v                *vault.Vault
	governance       *governance.Governance
	governanceHelper *governance.GovernanceHelper
	governanceAddr   common.Address
	prvvote          *prvvote.Prvvote
	prvvoteAddr      common.Address
	withdrawer       common.Address
	auth             *bind.TransactOpts
	EtherAddress     common.Address
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (v2 *PDaoTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")
	v2.withdrawer = ec.HexToAddress("0xe722D8b71DCC0152D47D2438556a45D3357d631f")
	v2.EtherAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")

	err := exec.Command("/bin/bash", "-c",
		"solc @openzeppelin/=node_modules/@openzeppelin/ --base-path=$(pwd)/bridge --bin --abi --optimize --overwrite bridge/contracts/pdao/governance.sol -o bridge/governance").Run()
	require.Equal(v2.T(), nil, err)
	//err = exec.Command("/bin/bash", "-c",
	//	"abigen --abi bridge/governance/IncognitoDAO.abi --bin bridge/governance/IncognitoDAO.bin --out bridge/governance/governance.go --pkg governance").Run()
	//require.Equal(v2.T(), nil, err)
	//err = exec.Command("/bin/bash", "-c",
	//	"solc @openzeppelin/=node_modules/@openzeppelin/ --base-path=$(pwd)/bridge --bin --abi --optimize --overwrite bridge/contracts/pdao/governanceHelper.sol -o bridge/governance").Run()
	//require.Equal(v2.T(), nil, err)
	//err = exec.Command("/bin/bash", "-c",
	//	"abigen --abi bridge/governance/GovernanceHelper.abi --bin bridge/governance/GovernanceHelper.bin --out bridge/governance/governanceHelper.go --pkg governanceHelper").Run()
	require.Equal(v2.T(), nil, err)
	err = exec.Command("/bin/bash", "-c", "solc @openzeppelin/=node_modules/@openzeppelin/ --base-path=$(pwd)/bridge --bin --abi --overwrite bridge/contracts/pdao/prv_vote.sol -o bridge/prvvote").Run()
	require.Equal(v2.T(), nil, err)
	err = exec.Command("/bin/bash", "-c", "abigen --abi bridge/prvvote/PrvVoting.abi --bin bridge/prvvote/PrvVoting.bin --out bridge/prvvote/prvvote.go --pkg prvvote").Run()
	require.Equal(v2.T(), nil, err)

	p, c, err := setupFixedCommittee()
	require.Equal(v2.T(), nil, err)
	v2.p = p
	v2.c = c
	require.Equal(v2.T(), nil, err)

	// deploy governance helper
	_, _, gHelper, err := governance.DeployGovernanceHelper(auth2, v2.p.sim)
	require.Equal(v2.T(), nil, err)
	v2.governanceHelper = gHelper
	v2.p.sim.Commit()

	// deploy prv vote
	var prvVote common.Address
	prvVote, _, _, err = prvvote.DeployPrvvote(auth2, v2.p.sim)
	require.Equal(v2.T(), nil, err)
	prvAbi, _ := abi.JSON(strings.NewReader(prvvote.PrvvoteMetaData.ABI))
	input, _ := prvAbi.Pack("initialize", "Incognito", "PRV")
	v2.p.sim.Commit()

	prvVote, _, _, err = vaultproxy.DeployTransparentUpgradeableProxy(auth, v2.p.sim, prvVote, auth3.From, v2.p.incAddr, input)
	require.Equal(v2.T(), nil, err)
	v2.prvvote, err = prvvote.NewPrvvote(prvVote, v2.p.sim)
	require.Equal(v2.T(), nil, err)
	v2.prvvoteAddr = prvVote
	v2.p.sim.Commit()

	// deploy governance
	var governanceAddr common.Address
	governanceAddr, _, _, err = governance.DeployGovernance(auth2, v2.p.sim)
	require.Equal(v2.T(), nil, err)
	governanceAbi, _ := abi.JSON(strings.NewReader(governance.GovernanceMetaData.ABI))
	input, _ = governanceAbi.Pack("initialize", prvVote)
	v2.p.sim.Commit()

	// todo: update to simplified proxy contract
	governanceAddr, _, _, err = vaultproxy.DeployTransparentUpgradeableProxy(auth, v2.p.sim, governanceAddr, auth3.From, v2.p.incAddr, input)
	require.Equal(v2.T(), nil, err)
	v2.governance, err = governance.NewGovernance(governanceAddr, v2.p.sim)
	v2.governanceAddr = governanceAddr
	require.Equal(v2.T(), nil, err)
	v2.p.sim.Commit()
}

func (v2 *PDaoTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
}

func (v2 *PDaoTestSuite) SetupTest() {
	fmt.Println("Setting up the test...")
}

func (v2 *PDaoTestSuite) TearDownTest() {
	fmt.Println("Tearing down the test...")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestPDao(t *testing.T) {
	fmt.Println("Starting entry point for pdao test suite...")
	suite.Run(t, new(PDaoTestSuite))

	fmt.Println("Finishing entry point for pdao test suite...")
}

func (v2 *PDaoTestSuite) TestPDAOCreateProp() {
	// burn and submit proof to prv vote contract
	proof := buildWithdrawTestcaseV2(v2.c, 170, 1, v2.prvvoteAddr, big.NewInt(1e10), auth2.From)
	prvInst, err := prveth.NewPrveth(v2.prvvoteAddr, v2.p.sim)
	require.Equal(v2.T(), nil, err)
	_, err = SubmitMintPRVProof(prvInst, auth, proof)
	require.Equal(v2.T(), nil, err)
	GenNewBlocks(v2.p.sim, 1)

	// query checkpoint
	prvBal, _ := v2.prvvote.BalanceOf(nil, auth2.From)
	require.Equal(v2.T(), big.NewInt(1e10), prvBal)

	GenNewBlocks(v2.p.sim, 2)
	prvBalSnapshot, err := v2.prvvote.GetPastVotes(nil, auth2.From, big.NewInt(0).Sub(v2.p.sim.Blockchain().CurrentHeader().Number, big.NewInt(1)))
	require.Equal(v2.T(), nil, err)
	require.Equal(v2.T(), big.NewInt(1e10), prvBalSnapshot)

	// create new proposal
	n := new(big.Int)
	n, ok := n.SetString("1"+strings.Repeat("0", 18), 10)
	require.Equal(v2.T(), true, ok)

	// transfer native token to governance
	gr := governance.GovernanceRaw{Contract: v2.governance}
	auth2.Value = n
	_, err = gr.Transfer(auth2)
	require.Equal(v2.T(), nil, err)
	auth2.Value = big.NewInt(0)

	testAccount := newAccount()
	// case fail cause not enough PRV token to create proposal
	_, err = v2.governance.Propose(auth, []common.Address{testAccount.Address}, []*big.Int{big.NewInt(1e10)}, [][]byte{{0x0}}, "move funds")
	require.NotEqual(v2.T(), nil, err)

	// auth2 create new proposal
	_, err = v2.governance.Propose(auth2, []common.Address{testAccount.Address}, []*big.Int{big.NewInt(1e11)}, [][]byte{{0x0}}, "move funds")
	require.Equal(v2.T(), nil, err)
	GenNewBlocks(v2.p.sim, 1)

	// auth2 create new proposal by signature
	testAccount2 := newAccount()
	propEncode, _ := v2.governanceHelper.BuildSignProposalEncodeAbi(
		nil,
		[]common.Address{testAccount2.Address},
		[]*big.Int{big.NewInt(1e11)},
		[][]byte{{0x0}},
		"move funds",
	)
	signData, err := v2.governance.GetDataSign(nil, keccak256(propEncode))
	require.Equal(v2.T(), nil, err)
	signBytes, err := crypto.Sign(signData[:], genesisAcc2.PrivateKey)
	require.Equal(v2.T(), nil, err)
	_, err = v2.governance.ProposeBySig(
		auth2,
		[]common.Address{testAccount2.Address},
		[]*big.Int{big.NewInt(1e11)},
		[][]byte{{0x0}},
		"move funds",
		signBytes[64]+27,
		toByte32(signBytes[:32]),
		toByte32(signBytes[32:64]),
	)
	require.Equal(v2.T(), nil, err)
	GenNewBlocks(v2.p.sim, 1)

	// cancel proposal 1 by signature

	// vote proposal 2 by signature

	// burn before snapshot day can vote

	// burn after snapshot day can not vote

	// execute proposal 2
}

func GenNewBlocks(s *backends.SimulatedBackend, n int) {
	i := 0
	for i != n {
		s.Commit()
		i++
	}
}
