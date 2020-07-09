package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"math/big"
	"strings"
	"testing"

	ec "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/incognitochain/bridge-eth/blockchain"
	"github.com/incognitochain/bridge-eth/bridge/incognito_proxy"
	"github.com/incognitochain/bridge-eth/common"
	"github.com/incognitochain/bridge-eth/consensus/signatureschemes/bridgesig"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

// func TestCompareCommittee(t *testing.T) {
// 	// // Committee on chain
// 	// url := "https://mainnet.incognito.org/fullnode:443"
// 	// _, _, err := getCommittee(url)
// 	// if err != nil {
// 	// 	t.Fatal(err)
// 	// }

// 	// Committee on contract
// 	committeeIdx := big.NewInt(0)
// 	_, c := connectAndInstantiate(t)
// 	comm, err := c.inc.GetBeaconCommittee(nil, committeeIdx)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	fmt.Println("Start block: ", comm.StartBlock.String())
// 	for _, addr := range comm.Pubkeys {
// 		fmt.Printf("beaconOld: %s\n", addr)
// 	}

// 	comm, err = c.inc.GetBridgeCommittee(nil, committeeIdx)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	fmt.Println("Start block: ", comm.StartBlock.String())
// 	for _, addr := range comm.Pubkeys {
// 		fmt.Printf("bridgeOld: %s\n", addr)
// 	}
// }

// func TestGetBridgeCommitteeOnChain(t *testing.T) {
// 	url := "https://test-node.incognito.org:443"
// 	_, _, err := getCommittee(url)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// func TestFixedFindBeaconCommitteeFromHeight(t *testing.T) {
// 	testCases := []struct {
// 		desc   string
// 		length int
// 		find   int64
// 		pos    int64
// 	}{
// 		{
// 			desc:   "A lot of committees",
// 			length: 100,
// 			find:   50,
// 			pos:    50,
// 		},
// 		{

// 			desc:   "Only 1 committee",
// 			length: 0,
// 			find:   50,
// 			pos:    0,
// 		},
// 		{

// 			desc:   "2 committees, get left",
// 			length: 1,
// 			find:   0,
// 			pos:    0,
// 		},
// 		{

// 			desc:   "2 committees, get right",
// 			length: 1,
// 			find:   1,
// 			pos:    1,
// 		},
// 		{

// 			desc:   "Get first committee",
// 			length: 10,
// 			find:   0,
// 			pos:    0,
// 		},
// 		{
// 			desc:   "Get last committee",
// 			length: 10,
// 			find:   10000000,
// 			pos:    10,
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.desc, func(t *testing.T) {
// 			p, c, _ := setupFixedCommittee()
// 			for i := 0; i < tc.length; i++ {
// 				in := repeatSwapBeacon(c, i+1, 70, 1)
// 				_, err := p.inc.SwapBeaconCommittee(auth, in.Instruction, in.InstPaths[0], in.InstPathIsLefts[0], in.InstRoots[0], in.BlkData[0], in.SigIdxs[0], in.SigVs[0], in.SigRs[0], in.SigSs[0])
// 				if err != nil {
// 					t.Fatal(err)
// 				}
// 				p.sim.Commit()
// 			}
// 			_, pos, err := p.inc.FindBeaconCommitteeFromHeight(nil, big.NewInt(tc.find))
// 			if err != nil || pos.Int64() != tc.pos {
// 				t.Errorf("invalid committee position, expect %d, got %d, err = %v", tc.pos, pos, err)
// 			}
// 		})
// 	}
// }

func repeatSwapBeacon(c *committees, startBlock, swapID, meta, shard int) *decodedProof {
	addrs := []string{
		"A5301a0d25103967bf0e29db1576cba3408fD9bB",
		"9BC0faE7BB432828759B6e391e0cC99995057791",
		"6cbc2937FEe477bbda360A842EeEbF92c2FAb613",
		"cabF3DB93eB48a61d41486AcC9281B6240411403",
	}
	inst, mp, blkData, blkHash := buildSwapData(meta, shard, startBlock, swapID, addrs)
	ip := signAndReturnInstProof(c.beaconPrivs, true, mp, blkData, blkHash[:])
	return &decodedProof{
		Instruction: inst,

		InstPaths:       [2][][32]byte{ip.instPath},
		InstPathIsLefts: [2][]bool{ip.instPathIsLeft},
		InstRoots:       [2][32]byte{ip.instRoot},
		BlkData:         [2][32]byte{ip.blkData},
		SigIdxs:         [2][]*big.Int{ip.sigIdx},
		SigVs:           [2][]uint8{ip.sigV},
		SigRs:           [2][][32]byte{ip.sigR},
		SigSs:           [2][][32]byte{ip.sigS},
	}
}

// func TestFixedFindBridgeCommitteeFromHeight(t *testing.T) {
// 	testCases := []struct {
// 		desc   string
// 		length int
// 		find   int64
// 		pos    int64
// 	}{
// 		{
// 			desc:   "A lot of committees",
// 			length: 100,
// 			find:   50,
// 			pos:    50,
// 		},
// 		{

// 			desc:   "Only 1 committee",
// 			length: 0,
// 			find:   50,
// 			pos:    0,
// 		},
// 		{

// 			desc:   "2 committees, get left",
// 			length: 1,
// 			find:   0,
// 			pos:    0,
// 		},
// 		{

// 			desc:   "2 committees, get right",
// 			length: 1,
// 			find:   1,
// 			pos:    1,
// 		},
// 		{

// 			desc:   "Get first committee",
// 			length: 10,
// 			find:   0,
// 			pos:    0,
// 		},
// 		{
// 			desc:   "Get last committee",
// 			length: 10,
// 			find:   10000000,
// 			pos:    10,
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.desc, func(t *testing.T) {
// 			p, c, _ := setupFixedCommittee()
// 			for i := 0; i < tc.length; i++ {
// 				in := repeatSwapBridge(c, i+1, 71, 1)
// 				_, err := p.inc.SwapBridgeCommittee(auth, in.Instruction, in.InstPaths, in.InstPathIsLefts, in.InstRoots, in.BlkData, in.SigIdxs, in.SigVs, in.SigRs, in.SigSs)
// 				if err != nil {
// 					t.Fatal(err)
// 				}
// 				p.sim.Commit()
// 			}
// 			_, pos, err := p.inc.FindBridgeCommitteeFromHeight(nil, big.NewInt(tc.find))
// 			if err != nil || pos.Int64() != tc.pos {
// 				t.Errorf("invalid committee position, expect %d, got %d, err = %v", tc.pos, pos, err)
// 			}
// 		})
// 	}
// }

// func repeatSwapBridge(c *committees, startBlock, meta, shard int) *decodedProof {
// 	addrs := []string{
// 		"3c78124783E8e39D1E084FdDD0E097334ba2D945",
// 		"76E34d8a527961286E55532620Af5b84F3C6538F",
// 		"68686dB6874588D2404155D00A73F82a50FDd190",
// 		"1533ac4d2922C150551f2F5dc2b0c1eDE382b890",
// 	}
// 	inst, mp, blkData, blkHash := buildSwapData(meta, shard, startBlock, addrs)
// 	ipBeacon := signAndReturnInstProof(c.beaconPrivs, true, mp, blkData, blkHash[:])
// 	ipBridge := signAndReturnInstProof(c.bridgePrivs, false, mp, blkData, blkHash[:])
// 	return &decodedProof{
// 		Instruction: inst,

// 		InstPaths:       [2][][32]byte{ipBeacon.instPath, ipBridge.instPath},
// 		InstPathIsLefts: [2][]bool{ipBeacon.instPathIsLeft, ipBridge.instPathIsLeft},
// 		InstRoots:       [2][32]byte{ipBeacon.instRoot, ipBridge.instRoot},
// 		BlkData:         [2][32]byte{ipBeacon.blkData, ipBridge.blkData},
// 		SigIdxs:         [2][]*big.Int{ipBeacon.sigIdx, ipBridge.sigIdx},
// 		SigVs:           [2][]uint8{ipBeacon.sigV, ipBridge.sigV},
// 		SigRs:           [2][][32]byte{ipBeacon.sigR, ipBridge.sigR},
// 		SigSs:           [2][][32]byte{ipBeacon.sigS, ipBridge.sigS},
// 	}
// }

// TestFixedSubmitBridgeCandidatePaused makes sure swapping committee isn't allowed when contract is paused
func TestFixedSubmitBridgeCandidatePaused(t *testing.T) {
	p, c, _ := setupFixedCommittee()

	in := buildBridgeCandidateTestcase(c, 789, 1, 71, 1)

	// Pause first, must success
	_, err := p.inc.Pause(auth)
	if err != nil {
		t.Fatalf("%+v", errors.Errorf("expect error == nil, got %v", err))
	}

	// Must fail
	inst, instProofs := buildIncognitoProxyInstructionProof(in)
	_, err = p.inc.SubmitBridgeCandidate(auth, inst, instProofs)

	// Check tx
	if err == nil {
		t.Fatalf("%+v", errors.Errorf("expect error != nil, got %v", err))
	}
	p.sim.Commit()

	// New committee mustn't be inserted
	comm, err := p.inc.BridgeCandidates(nil, big.NewInt(1))
	if err != nil {
		t.Fatalf("%+v", errors.Errorf("expect error == nil, got %v", err))
	}
	assert.Equal(t, comm.StartBlock.Int64(), int64(0))
	assert.Equal(t, comm.BeaconBlockHash, [32]byte{})
}

func TestFixedSubmitBridgeCandidate(t *testing.T) {
	_, c, _ := setupFixedCommittee()

	testCases := []struct {
		desc string
		in   *decodedProof
		out  int
		err  bool
	}{
		{
			desc: "Valid bridge swap instruction",
			in:   buildBridgeCandidateTestcase(c, 789, 1, 71, 1),
			out:  789,
		},
		{
			desc: "Invalid beacon inst",
			in: func() *decodedProof {
				proof := buildBridgeCandidateTestcase(c, 789, 1, 71, 1)
				proof.BlkData[1][0] = proof.BlkData[1][0] + 1
				return proof
			}(),
			err: true,
		},
		{
			desc: "Invalid bridge inst",
			in: func() *decodedProof {
				proof := buildBridgeCandidateTestcase(c, 789, 1, 71, 1)
				proof.InstRoots[1][0] = proof.InstRoots[1][0] + 1
				return proof
			}(),
			err: true,
		},
		{
			desc: "Invalid meta",
			in:   buildBridgeCandidateTestcase(c, 789, 1, 70, 1),
			err:  true,
		},
		{
			desc: "Invalid shard",
			in:   buildBridgeCandidateTestcase(c, 789, 1, 71, 2),
			err:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			p, _, _ := setupFixedCommittee()
			inst, instProofs := buildIncognitoProxyInstructionProof(tc.in)
			_, err := p.inc.SubmitBridgeCandidate(auth, inst, instProofs)
			isErr := err != nil
			if isErr != tc.err {
				t.Fatal(errors.Errorf("expect error = %t, got %v", tc.err, err))
			}
			if tc.err {
				return
			}
			p.sim.Commit()

			comm, err := p.inc.BridgeCandidates(nil, big.NewInt(1))
			if err != nil {
				t.Fatal(err)
			}
			if comm.StartBlock.Int64() != int64(tc.out) {
				t.Errorf("swap bridge failed, expect %v, got %v", tc.out, comm.StartBlock.Int64())
			}
		})
	}
}

func TestFixedSubmitManyBridgeCandidates(t *testing.T) {
	p, c, _ := setupFixedCommittee()
	decodedProofs := []*decodedProof{}
	for i := 0; i < 5; i++ {
		proof := &decodedProof{}
		proof, c = buildRandomBridgeCandidate(c, 789+i*400, i+1, 71, 1)
		decodedProofs = append(decodedProofs, proof)
	}
	for _, proof := range decodedProofs {
		inst, instProofs := buildIncognitoProxyInstructionProof(proof)
		_, err := p.inc.SubmitBridgeCandidate(auth, inst, instProofs)
		if err != nil {
			t.Fatal(err)
		}
		p.sim.Commit()
	}
}

func buildIncognitoProxyInstructionProof(proof *decodedProof) ([]byte, [2]incognito_proxy.IncognitoProxyInstructionProof) {
	instProofs := [2]incognito_proxy.IncognitoProxyInstructionProof{}
	for i := 0; i < 2; i++ {
		instProofs[i] = incognito_proxy.IncognitoProxyInstructionProof{
			Path:    proof.InstPaths[i],
			IsLeft:  proof.InstPathIsLefts[i],
			Root:    proof.InstRoots[i],
			BlkData: proof.BlkData[i],
			SigIdx:  proof.SigIdxs[i],
			SigV:    proof.SigVs[i],
			SigR:    proof.SigRs[i],
			SigS:    proof.SigSs[i],
		}
	}
	return proof.Instruction, instProofs
}

func buildRandomBridgeCandidate(c *committees, startBlock, swapID, meta, shard int) (*decodedProof, *committees) {
	n := 4
	addrs := []string{}
	newComm := &committees{
		beaconPrivs: c.beaconPrivs,
		beacons:     c.beacons,
	}
	for i := 0; i < n; i++ {
		key, _ := crypto.GenerateKey()
		privKey := crypto.FromECDSA(key)
		addr := crypto.PubkeyToAddress(key.PublicKey)
		newComm.bridgePrivs = append(newComm.bridgePrivs, privKey)
		newComm.bridges = append(newComm.bridges, addr)
		addrs = append(addrs, addr.String()[2:]) // Remove 0x
	}

	inst, mp, blkData, blkHash := buildSwapData(meta, shard, startBlock, swapID, addrs)
	ipBeacon := signAndReturnInstProof(c.beaconPrivs, true, mp, blkData, blkHash[:])
	ipBridge := signAndReturnInstProof(c.bridgePrivs, false, mp, blkData, blkHash[:])
	return &decodedProof{
		Instruction: inst,

		InstPaths:       [2][][32]byte{ipBeacon.instPath, ipBridge.instPath},
		InstPathIsLefts: [2][]bool{ipBeacon.instPathIsLeft, ipBridge.instPathIsLeft},
		InstRoots:       [2][32]byte{ipBeacon.instRoot, ipBridge.instRoot},
		BlkData:         [2][32]byte{ipBeacon.blkData, ipBridge.blkData},
		SigIdxs:         [2][]*big.Int{ipBeacon.sigIdx, ipBridge.sigIdx},
		SigVs:           [2][]uint8{ipBeacon.sigV, ipBridge.sigV},
		SigRs:           [2][][32]byte{ipBeacon.sigR, ipBridge.sigR},
		SigSs:           [2][][32]byte{ipBeacon.sigS, ipBridge.sigS},
	}, newComm
}

func buildBridgeCandidateTestcase(c *committees, startBlock, swapID, meta, shard int) *decodedProof {
	p, _ := buildRandomBridgeCandidate(c, startBlock, swapID, meta, shard)
	return p
}

// TestFixedSubmitBeaconCandidatePaused makes sure swapping committee isn't allowed when contract is paused
func TestFixedSubmitBeaconCandidatePaused(t *testing.T) {
	p, c, _ := setupFixedCommittee()

	in := buildBeaconCandidateTestcase(c, 789, 1, 70, 1)

	// Pause first, must success
	_, err := p.inc.Pause(auth)
	if err != nil {
		t.Fatalf("%+v", errors.Errorf("expect error == nil, got %v", err))
	}

	// Must fail
	inst, instProofs := buildIncognitoProxyInstructionProof(in)
	_, err = p.inc.SubmitBeaconCandidate(auth, inst, instProofs[0])

	// Check tx
	if err == nil {
		t.Fatalf("%+v", errors.Errorf("expect error != nil, got %v", err))
	}
	p.sim.Commit()

	// New committee mustn't be inserted
	comm, err := p.inc.BeaconCandidates(nil, big.NewInt(1))
	if err != nil {
		t.Fatalf("%+v", errors.Errorf("expect error == nil, got %v", err))
	}
	assert.Equal(t, comm.StartBlock.Int64(), int64(0))
	assert.Equal(t, comm.BeaconBlockHash, [32]byte{})
}

func TestFixedSubmitBeaconCandidate(t *testing.T) {
	_, c, _ := setupFixedCommittee()

	testCases := []struct {
		desc string
		in   *decodedProof
		out  int
		err  bool
	}{
		{
			desc: "Valid beacon swap instruction",
			in:   buildBeaconCandidateTestcase(c, 789, 1, 70, 1),
			out:  789,
		},
		{
			desc: "Invalid meta",
			in:   buildBeaconCandidateTestcase(c, 789, 1, 71, 1),
			err:  true,
		},
		{
			desc: "Invalid shard",
			in:   buildBeaconCandidateTestcase(c, 789, 1, 70, 2),
			err:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			p, _, _ := setupFixedCommittee()
			inst, instProofs := buildIncognitoProxyInstructionProof(tc.in)
			_, err := p.inc.SubmitBeaconCandidate(auth, inst, instProofs[0])
			isErr := err != nil
			if isErr != tc.err {
				t.Fatal(errors.Errorf("expect error = %t, got %v", tc.err, err))
			}
			if tc.err {
				return
			}
			p.sim.Commit()

			comm, err := p.inc.BeaconCandidates(nil, big.NewInt(1))
			if err != nil {
				t.Fatal(err)
			}
			if comm.StartBlock.Int64() != int64(tc.out) {
				t.Errorf("swap bridge failed, expect %v, got %v", tc.out, comm.StartBlock.Int64())
			}
		})
	}
}

func buildBeaconCandidateTestcase(c *committees, startBlock, swapID, meta, shard int) *decodedProof {
	p, _ := buildRandomBridgeCandidate(c, startBlock, swapID, meta, shard)
	return p
}

func buildSwapData(meta, shard, startBlock, swapID int, addrs []string) ([]byte, *merklePath, []byte, []byte) {
	// Build instruction merkle tree
	numInst := 10
	startNodeID := 7
	inst := buildDecodedSwapConfirmInst(meta, shard, startBlock, swapID, addrs)
	data := randomMerkleHashes(numInst)
	data[startNodeID] = inst
	mp := buildInstructionMerklePath(data, startNodeID)

	// Generate random blkHash
	h := randomMerkleHashes(1)
	blkData := h[0]
	blkHash := common.Keccak256(blkData, mp.root[:])
	return inst, mp, blkData, blkHash[:]
}

func TestFixedSubmitFinalityProof(t *testing.T) {
	_, c, _ := setupFixedCommittee()

	testCases := []struct {
		desc            string
		isBeacon        bool
		swapID          int
		meta            []byte
		blockMerkleRoot [][]byte
		proposeTime     []int
		init            func([2]incognito_proxy.IncognitoProxyInstructionProof)
		err             bool
	}{
		{
			desc:            "Valid bridge proof",
			isBeacon:        false,
			swapID:          0,
			meta:            []byte{73, 73},
			blockMerkleRoot: [][]byte{randomMerkleHashes(1)[0], randomMerkleHashes(1)[0]},
			proposeTime:     []int{123456780, 123456790},
		},
		{
			desc:            "Valid beacon proof",
			isBeacon:        true,
			swapID:          0,
			meta:            []byte{73, 73},
			blockMerkleRoot: [][]byte{randomMerkleHashes(1)[0], randomMerkleHashes(1)[0]},
			proposeTime:     []int{123456780, 123456790},
		},
		{
			desc:            "Invalid first proof",
			isBeacon:        false,
			swapID:          0,
			meta:            []byte{72, 72},
			blockMerkleRoot: [][]byte{randomMerkleHashes(1)[0], randomMerkleHashes(1)[0]},
			proposeTime:     []int{123456780, 123456790},
			err:             true,
			init: func(instProofs [2]incognito_proxy.IncognitoProxyInstructionProof) {
				instProofs[0].BlkData[0] = instProofs[0].BlkData[0] + 1
			},
		},
		{
			desc:            "Invalid second proof",
			isBeacon:        false,
			swapID:          0,
			meta:            []byte{72, 72},
			blockMerkleRoot: [][]byte{randomMerkleHashes(1)[0], randomMerkleHashes(1)[0]},
			proposeTime:     []int{123456780, 123456790},
			err:             true,
			init: func(instProofs [2]incognito_proxy.IncognitoProxyInstructionProof) {
				instProofs[1].BlkData[0] = instProofs[1].BlkData[0] + 1
			},
		},
		{
			desc:            "Invalid meta",
			isBeacon:        false,
			swapID:          0,
			meta:            []byte{73, 72},
			blockMerkleRoot: [][]byte{randomMerkleHashes(1)[0], randomMerkleHashes(1)[0]},
			proposeTime:     []int{123456780, 123456790},
			err:             true,
		},
		{
			desc:            "Invalid swapID",
			isBeacon:        false,
			swapID:          4,
			meta:            []byte{73, 72},
			blockMerkleRoot: [][]byte{randomMerkleHashes(1)[0], randomMerkleHashes(1)[0]},
			proposeTime:     []int{123456780, 123456790},
			err:             true,
		},
		{
			desc:            "Invalid timeslot",
			isBeacon:        false,
			swapID:          4,
			meta:            []byte{73, 72},
			blockMerkleRoot: [][]byte{randomMerkleHashes(1)[0], randomMerkleHashes(1)[0]},
			proposeTime:     []int{123456780, 123456785},
			err:             true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			p, _, _ := setupFixedCommittee()
			insts, instProofs := buildFinalityTestcase(
				c,
				tc.isBeacon,
				tc.meta,
				tc.blockMerkleRoot,
				tc.proposeTime,
			)
			if tc.init != nil {
				tc.init(instProofs)
			}

			_, err := p.inc.SubmitFinalityProof(auth, insts, instProofs, big.NewInt(int64(tc.swapID)), tc.isBeacon)
			isErr := err != nil
			if isErr != tc.err {
				t.Fatal(errors.Errorf("expect error = %t, got %v", tc.err, err))
			}
			if tc.err {
				return
			}
			p.sim.Commit()

			// Check finality result
			assertFinalityUpdated(t, p, tc.isBeacon)
		})
	}
}

func assertFinalityUpdated(t *testing.T, p *Platform, isBeacon bool) {
	var rootHash [32]byte
	if isBeacon {
		s, err := p.inc.BeaconFinality(nil)
		if err != nil {
			t.Fatal(err)
		}
		rootHash = s.RootHash
	} else {
		s, err := p.inc.BridgeFinality(nil)
		if err != nil {
			t.Fatal(err)
		}
		rootHash = s.RootHash
	}

	assert.NotEqual(t, [32]byte{}, rootHash)
}

func buildFinalityTestcase(
	c *committees,
	isBeacon bool,
	meta []byte,
	blockMerkleRoot [][]byte,
	proposeTime []int,
) ([2][]byte, [2]incognito_proxy.IncognitoProxyInstructionProof) {
	privs := [][]byte{}
	if isBeacon {
		privs = c.beaconPrivs
	} else {
		privs = c.bridgePrivs
	}

	// First block
	inst0, instProof0 := buildFinalityInstructionProof(privs, meta[0], blockMerkleRoot[0], proposeTime[0])

	// Second block
	inst1, instProof1 := buildFinalityInstructionProof(privs, meta[1], blockMerkleRoot[1], proposeTime[1])

	return [2][]byte{inst0, inst1}, [2]incognito_proxy.IncognitoProxyInstructionProof{
		instProof0,
		instProof1,
	}
}

func buildFinalityInstructionProof(
	privs [][]byte,
	meta byte,
	blockMerkleRoot []byte,
	proposeTime int,
) ([]byte, incognito_proxy.IncognitoProxyInstructionProof) {
	inst, mp, blkData, blkHash := buildFinalityData(meta, blockMerkleRoot, proposeTime)
	proof := signAndReturnInstProof(privs, true, mp, blkData, blkHash[:])
	return inst, incognito_proxy.IncognitoProxyInstructionProof{
		Path:    proof.instPath,
		IsLeft:  proof.instPathIsLeft,
		Root:    proof.instRoot,
		BlkData: proof.blkData,
		SigIdx:  proof.sigIdx,
		SigV:    proof.sigV,
		SigR:    proof.sigR,
		SigS:    proof.sigS,
	}
}

func buildFinalityData(meta byte, blockMerkleRoot []byte, proposeTime int) ([]byte, *merklePath, []byte, []byte) {
	// Build instruction merkle tree
	numInst := 12
	startNodeID := 9
	inst := buildDecodedBlockMerkleInst(meta, blockMerkleRoot, proposeTime)
	data := randomMerkleHashes(numInst)
	data[startNodeID] = inst
	mp := buildInstructionMerklePath(data, startNodeID)

	// Generate random blkHash
	h := randomMerkleHashes(1)
	blkData := h[0]
	blkHash := common.Keccak256(blkData, mp.root[:])
	return inst, mp, blkData, blkHash[:]
}

func buildDecodedBlockMerkleInst(meta byte, blockMerkleRoot []byte, proposeTime int) []byte {
	decoded := []byte{byte(meta)}
	decoded = append(decoded, blockMerkleRoot...)
	decoded = append(decoded, toBytes32BigEndian(big.NewInt(int64(proposeTime)).Bytes())...)
	return decoded
}

type promoteCandidateTestcase struct {
	candidateProof *decodedProof
	finalityInsts  [2][]byte
	finalityProofs [2]incognito_proxy.IncognitoProxyInstructionProof
	promoteProof   incognito_proxy.IncognitoProxyMerkleProof
}

func TestFixedPromoteCandidate(t *testing.T) {
	_, c, _ := setupFixedCommittee()

	testCases := []struct {
		desc     string
		isBeacon bool
		swapID   int
		in       promoteCandidateTestcase
		err      bool
	}{
		{
			desc:     "Valid bridge promotion",
			isBeacon: false,
			swapID:   1,
			in:       buildPromoteCandidateTestcase(c, false, 1),
		},
		{
			desc:     "Valid beacon promotion",
			isBeacon: true,
			swapID:   1,
			in:       buildPromoteCandidateTestcase(c, true, 1),
		},
		{
			desc:     "Invalid proof path",
			isBeacon: false,
			swapID:   1,
			in: func() promoteCandidateTestcase {
				t := buildPromoteCandidateTestcase(c, false, 1)
				t.promoteProof.IsLeft[0] = !t.promoteProof.IsLeft[0]
				return t
			}(),
			err: true,
		},
		{
			desc:     "Promote old swapID",
			isBeacon: false,
			swapID:   0,
			in:       buildPromoteCandidateTestcase(c, false, 1),
			err:      true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			p, _, _ := setupFixedCommittee()

			// Submit candidate, must success
			var err error
			inst, instProofs := buildIncognitoProxyInstructionProof(tc.in.candidateProof)
			if tc.isBeacon {
				_, err = p.inc.SubmitBeaconCandidate(auth, inst, instProofs[0])
			} else {
				_, err = p.inc.SubmitBridgeCandidate(auth, inst, instProofs)
			}
			assert.Nil(t, err)
			p.sim.Commit()

			// Submit finality, must success
			_, err = p.inc.SubmitFinalityProof(auth, tc.in.finalityInsts, tc.in.finalityProofs, big.NewInt(int64(0)), true)
			assert.Nil(t, err)
			p.sim.Commit()

			// Promote candidate and check results
			_, err = p.inc.PromoteCandidate(auth, big.NewInt(int64(tc.swapID)), tc.isBeacon, tc.in.promoteProof)
			isErr := err != nil
			if isErr != tc.err {
				t.Fatal(errors.Errorf("expect error = %t, got %v", tc.err, err))
			}
			if tc.err {
				return
			}
			p.sim.Commit()

			var c incognito_proxy.IncognitoProxyCommittee
			if tc.isBeacon {
				c, err = p.inc.GetBeaconCommittee(nil, big.NewInt(1))
			} else {
				c, err = p.inc.GetBridgeCommittee(nil, big.NewInt(1))
			}
			assert.Greater(t, c.StartBlock.Int64(), int64(0))
			assert.Nil(t, err)
		})
	}
}

func TestFixedPromoteCandidateSkippingSwapID(t *testing.T) {
	p, c, _ := setupFixedCommittee()

	// Build candidate proof
	candidateProof, c2 := buildRandomBridgeCandidate(c, 123, 1, 71, 1)
	blkHash := common.Keccak256(candidateProof.BlkData[0][:], candidateProof.InstRoots[0][:])
	candidateProof2, _ := buildRandomBridgeCandidate(c2, 234, 2, 71, 1)
	blkHash2 := common.Keccak256(candidateProof2.BlkData[0][:], candidateProof2.InstRoots[0][:])

	// Submit both candidates, must success
	inst, instProofs := buildIncognitoProxyInstructionProof(candidateProof)
	_, err := p.inc.SubmitBridgeCandidate(auth, inst, instProofs)
	assert.Nil(t, err)
	p.sim.Commit()
	inst, instProofs = buildIncognitoProxyInstructionProof(candidateProof2)
	_, err = p.inc.SubmitBridgeCandidate(auth, inst, instProofs)
	assert.Nil(t, err)
	p.sim.Commit()

	// Build promotion proof
	numBlocks := 4567
	candidateBlockHeight := 456
	candidateBlockHeight2 := 789
	data := randomMerkleHashes(numBlocks)
	data[candidateBlockHeight] = blkHash[:]
	data[candidateBlockHeight2] = blkHash2[:]
	merkleProof := buildPromoteCandidateProof(data, candidateBlockHeight2)
	promoteProof := buildIncognitoProxyMerkleProof(merkleProof)

	// Build finality proof
	finalityMeta := []byte{73, 73}
	b := [32]byte{}
	blockMerkleRoot := [][]byte{merkleProof.root[:], b[:]}
	proposeTime := []int{123456780, 123456790}
	finalityInsts, finalityProofs := buildFinalityTestcase(c, true, finalityMeta, blockMerkleRoot, proposeTime) // finality proof must be on beacon

	// Submit finality, must success
	_, err = p.inc.SubmitFinalityProof(auth, finalityInsts, finalityProofs, big.NewInt(int64(0)), true)
	assert.Nil(t, err)
	p.sim.Commit()

	// Promote candidate and check results
	_, err = p.inc.PromoteCandidate(auth, big.NewInt(int64(2)), false, promoteProof)
	assert.NotNil(t, err)
}

func TestFixedPromoteManyCandidate(t *testing.T) {
	p, c, _ := setupFixedCommittee()
	numCandidates := 10
	numBlocks := 4567

	data := randomMerkleHashes(numBlocks)
	candidateBlockHeights := []int{}
	var candidateProof *decodedProof
	for i := 0; i < numCandidates; i++ {
		// Build candidate proof
		candidateProof, c = buildRandomBridgeCandidate(c, 123+i, i+1, 71, 1)
		blkHash := common.Keccak256(candidateProof.BlkData[0][:], candidateProof.InstRoots[0][:])
		data[456+i] = blkHash[:]
		candidateBlockHeights = append(candidateBlockHeights, 456+i)

		// Submit candidate, must success
		inst, instProofs := buildIncognitoProxyInstructionProof(candidateProof)
		_, err := p.inc.SubmitBridgeCandidate(auth, inst, instProofs)
		assert.Nil(t, err)
		p.sim.Commit()
	}

	// Build promotion proofs
	promoteProofs := []incognito_proxy.IncognitoProxyMerkleProof{}
	var root []byte
	for i := 0; i < numCandidates; i++ {
		merkleProof := buildPromoteCandidateProof(data, candidateBlockHeights[i])
		promoteProof := buildIncognitoProxyMerkleProof(merkleProof)
		promoteProofs = append(promoteProofs, promoteProof)
		root = merkleProof.root[:]
	}

	// Build finality proof
	finalityMeta := []byte{73, 73}
	b := [32]byte{}
	blockMerkleRoot := [][]byte{root, b[:]}
	proposeTime := []int{123456780, 123456790}
	finalityInsts, finalityProofs := buildFinalityTestcase(c, true, finalityMeta, blockMerkleRoot, proposeTime) // finality proof must be on beacon

	// Submit finality, must success
	_, err := p.inc.SubmitFinalityProof(auth, finalityInsts, finalityProofs, big.NewInt(int64(0)), true)
	assert.Nil(t, err)
	p.sim.Commit()

	// Promote candidates and check results
	for i, proof := range promoteProofs {
		_, err := p.inc.PromoteCandidate(auth, big.NewInt(int64(i+1)), false, proof)
		assert.Nil(t, err)
		p.sim.Commit()

		c, err := p.inc.GetBridgeCommittee(nil, big.NewInt(int64(i+1)))
		assert.Equal(t, len(c.Pubkeys), 4)
	}
}

func buildPromoteCandidateTestcase(c *committees, isBeacon bool, swapID int) promoteCandidateTestcase {
	// Build candidate proof
	startBlock := 123
	shard := 1
	meta := 71
	if isBeacon {
		meta = 70
	}
	candidateProof, _ := buildRandomBridgeCandidate(c, startBlock, swapID, meta, shard)
	blkHash := common.Keccak256(candidateProof.BlkData[0][:], candidateProof.InstRoots[0][:])

	// Build promotion proof
	numBlocks := 789
	candidateBlockHeight := 456
	data := randomMerkleHashes(numBlocks)
	data[candidateBlockHeight] = blkHash[:]
	merkleProof := buildPromoteCandidateProof(data, candidateBlockHeight)
	promoteProof := buildIncognitoProxyMerkleProof(merkleProof)

	// Build finality proof
	finalityMeta := []byte{73, 73}
	b := [32]byte{}
	blockMerkleRoot := [][]byte{merkleProof.root[:], b[:]}
	proposeTime := []int{123456780, 123456790}
	finalityInsts, finalityProofs := buildFinalityTestcase(c, true, finalityMeta, blockMerkleRoot, proposeTime) // finality proof must be on beacon

	return promoteCandidateTestcase{candidateProof, finalityInsts, finalityProofs, promoteProof}
}

func buildPromoteCandidateProof(blockHashes [][]byte, candidateBlockHeight int) *merklePath {
	mp := buildInstructionMerklePath(blockHashes, candidateBlockHeight)
	return mp
}

func buildIncognitoProxyMerkleProof(merkleProof *merklePath) incognito_proxy.IncognitoProxyMerkleProof {
	promoteProof := incognito_proxy.IncognitoProxyMerkleProof{
		Path:   merkleProof.path,
		IsLeft: merkleProof.left,
	}
	return promoteProof
}

// func TestFixedInstructionApproved(t *testing.T) {
// 	p, c, _ := setupFixedCommittee()

// 	testCases := []struct {
// 		desc string
// 		in   *instProof
// 		out  bool
// 		err  bool
// 	}{
// 		{
// 			desc: "Valid beacon swap instruction",
// 			in:   buildInstructionApprovedTestcase(true, c),
// 			out:  true,
// 		},
// 		{
// 			desc: "Valid bridge swap instruction",
// 			in:   buildInstructionApprovedTestcase(false, c),
// 			out:  true,
// 		},
// 		{
// 			desc: "SigIdx incorrect",
// 			in: func() *instProof {
// 				p := buildInstructionApprovedTestcase(true, c)
// 				p.sigIdx[0] = big.NewInt(20)
// 				return p
// 			}(),
// 			out: false,
// 		},
// 		{
// 			desc: "Not enough sigs",
// 			in: func() *instProof {
// 				p := buildInstructionApprovedTestcase(true, c)
// 				p.sigIdx = p.sigIdx[:2]
// 				return p
// 			}(),
// 			err: true,
// 		},
// 		{
// 			desc: "Duplicated sigIdx",
// 			in: func() *instProof {
// 				p := buildInstructionApprovedTestcase(true, c)
// 				p.sigIdx[0] = p.sigIdx[1]
// 				p.sigV[0] = p.sigV[1]
// 				p.sigR[0] = p.sigR[1]
// 				p.sigS[0] = p.sigS[1]
// 				return p
// 			}(),
// 			out: false,
// 		},
// 		{
// 			desc: "Padded sigIdx",
// 			in: func() *instProof {
// 				p := buildInstructionApprovedTestcase(true, c)
// 				p.sigV = p.sigV[:1]
// 				p.sigR = p.sigR[:1]
// 				p.sigS = p.sigS[:1]
// 				return p
// 			}(),
// 			err: true,
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.desc, func(t *testing.T) {
// 			res, err := p.inc.InstructionApproved(nil, tc.in.isBeacon, tc.in.instHash, tc.in.blkHeight, tc.in.instPath, tc.in.instPathIsLeft, tc.in.instRoot, tc.in.blkData, tc.in.sigIdx, tc.in.sigV, tc.in.sigR, tc.in.sigS)
// 			isErr := err != nil
// 			if isErr != tc.err {
// 				t.Error(errors.Errorf("expect error = %t, got %v", tc.err, err))
// 			}
// 			if tc.err {
// 				return
// 			}
// 			if res != tc.out {
// 				t.Errorf("verify inst approval failed, expect %v, got %v", tc.out, res)
// 			}
// 		})
// 	}
// }

func buildInstructionApprovedTestcase(isBeacon bool, c *committees) *instProof {
	mp := buildMerklePathTestcase(8, 6, 6)

	// Generate random blkHash
	h := randomMerkleHashes(1)
	blkData := h[0]
	blkHash := common.Keccak256(blkData, mp.root[:])

	// Get private keys
	privs := c.beaconPrivs
	if !isBeacon {
		privs = c.bridgePrivs
	}

	return signAndReturnInstProof(privs, isBeacon, mp, blkData, blkHash[:])
}

func signAndReturnInstProof(
	privs [][]byte,
	isBeacon bool,
	mp *merklePath,
	blkData []byte,
	blkHash []byte,
) *instProof {
	sigV := make([]uint8, len(privs))
	sigR := make([][32]byte, len(privs))
	sigS := make([][32]byte, len(privs))
	sigIdx := make([]*big.Int, len(privs))
	for i, p := range privs {
		sig, _ := bridgesig.Sign(p, blkHash)
		sigV[i] = uint8(sig[64] + 27)
		sigR[i] = toByte32(sig[:32])
		sigS[i] = toByte32(sig[32:64])
		sigIdx[i] = big.NewInt(int64(i))
	}

	return &instProof{
		isBeacon:       isBeacon,
		instHash:       mp.leaf,
		blkHeight:      big.NewInt(0),
		instPath:       mp.path,
		instPathIsLeft: mp.left,
		instRoot:       mp.root,
		blkData:        toByte32(blkData),
		sigIdx:         sigIdx,
		sigV:           sigV,
		sigR:           sigR,
		sigS:           sigS,
	}
}

type instProof struct {
	isBeacon       bool
	instHash       [32]byte
	blkHeight      *big.Int
	instPath       [][32]byte
	instPathIsLeft []bool
	instRoot       [32]byte
	blkData        [32]byte
	sigIdx         []*big.Int
	sigV           []uint8
	sigR           [][32]byte
	sigS           [][32]byte
}

func TestFixedVerifySig(t *testing.T) {
	p, _, _ := setupFixedCommittee()

	testCases := []struct {
		desc string
		in   *committeeSig
		out  bool
		err  bool
	}{
		{
			desc: "Valid sig",
			in:   getFixedCommitteeSig(),
			out:  true,
		},
		{
			desc: "Invalid committee",
			in: func() *committeeSig {
				sig := getFixedCommitteeSig()
				sig.addrs[1][2] = 123
				return sig
			}(),
			out: false,
		},
		{
			desc: "Invalid msgHash",
			in: func() *committeeSig {
				sig := getFixedCommitteeSig()
				sig.msgHash[0] = 123
				return sig
			}(),
			out: false,
		},
		{
			desc: "Invalid v",
			in: func() *committeeSig {
				sig := getFixedCommitteeSig()
				sig.v[1] = 123
				return sig
			}(),
			out: false,
		},
		{
			desc: "Invalid r",
			in: func() *committeeSig {
				sig := getFixedCommitteeSig()
				sig.r[2][3] = 123
				return sig
			}(),
			out: false,
		},
		{
			desc: "Invalid s",
			in: func() *committeeSig {
				sig := getFixedCommitteeSig()
				sig.s[3][4] = 123
				return sig
			}(),
			out: false,
		},
		{
			desc: "Not enough committee members",
			in: func() *committeeSig {
				sig := getFixedCommitteeSig()
				sig.addrs = sig.addrs[:2]
				return sig
			}(),
			err: true,
		},
		{
			desc: "Not enough v",
			in: func() *committeeSig {
				sig := getFixedCommitteeSig()
				sig.v = sig.v[:2]
				return sig
			}(),
			err: true,
		},
		{
			desc: "Not enough r",
			in: func() *committeeSig {
				sig := getFixedCommitteeSig()
				sig.r = sig.r[:2]
				return sig
			}(),
			err: true,
		},
		{
			desc: "Not enough s",
			in: func() *committeeSig {
				sig := getFixedCommitteeSig()
				sig.s = sig.s[:2]
				return sig
			}(),
			err: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			res, err := p.inc.VerifySig(nil, tc.in.addrs, tc.in.msgHash, tc.in.v, tc.in.r, tc.in.s)
			isErr := err != nil
			if isErr != tc.err {
				t.Error(errors.Errorf("expect error = %t, got %v", tc.err, err))
			}
			if tc.err {
				return
			}
			if res != tc.out {
				t.Errorf("verifySig failed, expect %v, got %v", tc.out, res)
			}
		})
	}
}

func getFixedCommitteeSig() *committeeSig {
	validationData := "{\"ProducerBLSSig\":\"D4sg/eVi8yI+rX9WOwCWBEG+4mWXjGNorl2m3ppRCvE=\",\"ProducerBriSig\":null,\"ValidatiorsIdx\":[0,1,2,3],\"AggSig\":\"AriRXDvXcPDqkMNAQjHR61f3xis6YLNskuYF7vQJNzE=\",\"BridgeSig\":[\"/tSXMa9s1PKAxDC9H6etSMPcnAOEqqQYum3TfWtOKQpyvHxA1jllDkLmB68M6pp54bTUWenqXMQVWW+2GAcBjgA=\",\"MJyhaCCm8B6uwK/w6/OMqr7AW1Szo1etRfTcru0ZenZUwea0LVXhPo2QRKeO+Q1n12J2yRv4sUkhRLLL9zw1SwE=\",\"DOpccVDrw6SbGqs4+YP/Ti1nx4gg/xpsuHB7DBuhO2RMl8hAaUz2TVZ6hv+r8z0YLiUw/k6FEFY+5dg/EjMRAQA=\",\"qPEXt4KgFR8ZMw7JelEeEwsWQ7gW/IrzWMpx++zjQ6dLdeXwKcGwxoaBWhWnEpma+MVVQw1LvzzuvtzIBGZDKgE=\"]}"
	d, _ := DecodeValidationData(validationData)
	vs := []byte{}
	rs := [][32]byte{}
	ss := [][32]byte{}
	for _, sig := range d.BridgeSig {
		v, r, s, _ := bridgesig.DecodeECDSASig(sig)
		vs = append(vs, v)
		rs = append(rs, toByte32(r))
		ss = append(ss, toByte32(s))
	}

	hash, _ := common.Hash{}.NewHashFromStr("cb53ba7574335ecfa0fddcb136b387330af322784fb759c80ca7bb790a1c0f9d")
	c := getFixedCommittee()
	msgHash := toByte32(crypto.Keccak256Hash(hash.GetBytes()).Bytes())
	return &committeeSig{
		addrs:   c.beacons,
		msgHash: msgHash,
		v:       vs,
		r:       rs,
		s:       ss,
	}
}

type committeeSig struct {
	addrs   []ec.Address
	msgHash [32]byte
	v       []uint8
	r       [][32]byte
	s       [][32]byte
}

type ValidationData struct {
	ProducerBLSSig []byte
	ProducerBriSig []byte
	ValidatiorsIdx []int
	AggSig         []byte
	BridgeSig      [][]byte
}

func DecodeValidationData(data string) (*ValidationData, error) {
	var valData ValidationData
	err := json.Unmarshal([]byte(data), &valData)
	if err != nil {
		return nil, err
	}
	return &valData, nil
}

// // TestFixedSwapBridgeFixedProof the same as the next test but for bridge
// func TestFixedSwapBridgeFixedProof(t *testing.T) {
// 	proof := getFixedSwapBridgeProof()

// 	// p, err := setupWithLocalCommittee()
// 	p, _, err := setupFixedCommittee()
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	tx, err := SwapBridge(p.inc, auth, proof)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	p.sim.Commit()
// 	printReceipt(p.sim, tx)
// }

// // TestFixedSwapBeaconFixedProof decodes a fixed proof and submit to make sure proof
// // format wasn't changed without updating bot
// func TestFixedSwapBeaconFixedProof(t *testing.T) {
// 	proof := getFixedSwapBeaconProof()

// 	p, _, err := setupFixedCommittee()
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	tx, err := SwapBeacon(p.inc, auth, proof)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	p.sim.Commit()
// 	printReceipt(p.sim, tx)
// }

// func TestFixedExtractMetaFromInstruction(t *testing.T) {
// 	p, _, _ := setupFixedCommittee()
// 	addrs := []string{
// 		"834f98e1b7324450b798359c9febba74fb1fd888",
// 		"1250ba2c592ac5d883a0b20112022f541898e65b",
// 		"2464c00eab37be5a679d6e5f7c8f87864b03bfce",
// 		"6d4850ab610be9849566c09da24b37c5cfa93e50",
// 	}
// 	testCases := []struct {
// 		desc    string
// 		inst    []byte
// 		meta    uint8
// 		shard   uint8
// 		height  int
// 		numVals int
// 		err     bool
// 	}{
// 		{
// 			desc:    "Extract swap beacon instruction",
// 			inst:    buildDecodedSwapConfirmInst(70, 1, 123, addrs),
// 			meta:    70,
// 			shard:   1,
// 			height:  123,
// 			numVals: len(addrs),
// 		},
// 		{
// 			desc:    "Extract swap bridge instruction",
// 			inst:    buildDecodedSwapConfirmInst(71, 2, 19827312, addrs[:2]),
// 			meta:    71,
// 			shard:   2,
// 			height:  19827312,
// 			numVals: 2,
// 		},
// 		{
// 			desc: "Instruction too short",
// 			inst: make([]byte, 65),
// 			err:  true,
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.desc, func(t *testing.T) {
// 			meta, shard, height, numVals, err := p.inc.ExtractMetaFromInstruction(nil, tc.inst)
// 			isErr := err != nil
// 			if isErr != tc.err {
// 				t.Error(errors.Errorf("expect error = %t, got %v", tc.err, err))
// 			}
// 			if tc.err {
// 				return
// 			}
// 			if meta != tc.meta {
// 				t.Errorf("invalid meta, expect %v, got %v", tc.meta, meta)
// 			}
// 			if shard != tc.shard {
// 				t.Errorf("invalid shard, expect %v, got %v", tc.shard, shard)
// 			}
// 			if height.Int64() != int64(tc.height) {
// 				t.Errorf("invalid height, expect %v, got %v", tc.height, height)
// 			}
// 			if numVals.Int64() != int64(tc.numVals) {
// 				t.Errorf("invalid numVals, expect %v, got %v", tc.numVals, numVals)
// 			}
// 		})
// 	}
// }

func TestFixedExtractCommitteeFromInstruction(t *testing.T) {
	p, _, _ := setupFixedCommittee()
	addrs := []string{
		"834f98e1b7324450b798359c9febba74fb1fd888",
		"1250ba2c592ac5d883a0b20112022f541898e65b",
		"2464c00eab37be5a679d6e5f7c8f87864b03bfce",
		"6d4850ab610be9849566c09da24b37c5cfa93e50",
	}
	testCases := []struct {
		desc    string
		inst    []byte
		numVals int
		out     []string
		err     bool
	}{
		{
			desc:    "Extract beacon committee",
			inst:    buildDecodedSwapConfirmInst(70, 1, 789, 1, addrs),
			numVals: len(addrs),
			out:     addrs,
		},
		{
			desc:    "Extract bridge committee",
			inst:    buildDecodedSwapConfirmInst(71, 1, 19827312, 1, addrs[:2]),
			numVals: 2,
			out:     addrs[:2],
		},
		{
			desc:    "Instruction too short",
			inst:    make([]byte, 97),
			numVals: 1,
			err:     true,
		},
		{
			desc:    "numVals too big",
			inst:    buildDecodedSwapConfirmInst(70, 1, 789, 1, addrs),
			numVals: 8,
			err:     true,
		},
		{
			desc:    "numVals too low",
			inst:    buildDecodedSwapConfirmInst(70, 1, 789, 1, addrs),
			numVals: 2,
			err:     true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			comm, err := p.inc.ExtractCommitteeFromInstruction(nil, tc.inst, big.NewInt(int64(tc.numVals)))
			isErr := err != nil
			if isErr != tc.err {
				t.Error(errors.Errorf("expect error = %t, got %v", tc.err, err))
			}
			if tc.err {
				return
			}

			for i, c := range comm {
				addr := c.Hex()
				addr = addr[2:] // ignore 0x
				if strings.ToLower(addr) != tc.out[i] {
					t.Errorf("invalid committee[%d], expect %v, got %v", i, tc.out[i], addr)
				}
			}
		})
	}
}

func TestFixedLeafInMerkleTree(t *testing.T) {
	p, _, _ := setupFixedCommittee()
	testCases := []struct {
		desc string
		in   *merklePath
		out  bool
		err  bool
	}{
		{
			desc: "In merkle tree",
			in:   buildMerklePathTestcase(8, 6, 6),
			out:  true,
		},
		{
			desc: "Not in merkle tree",
			in:   buildMerklePathTestcase(8, 6, 4),
			out:  false,
		},
		{
			desc: "Random leaf",
			in:   buildMerklePathTestcase(8, 6, -1),
			out:  false,
		},
		{
			desc: "Big tree",
			in:   buildMerklePathTestcase(100000, 12345, 12345),
			out:  true,
		},
		{
			desc: "Single node",
			in:   buildMerklePathTestcase(1, 0, 0),
			out:  true,
		},
		{
			desc: "Invalid left.length",
			in: func() *merklePath {
				mp := buildMerklePathTestcase(10, 9, 9)
				mp.left = mp.left[:len(mp.left)-2]
				return mp
			}(),
			err: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			isIn, err := p.inc.LeafInMerkleTree(nil, tc.in.leaf, tc.in.root, tc.in.path, tc.in.left)
			isErr := err != nil
			if isErr != tc.err {
				t.Error(errors.Errorf("expect error = %t, got %v", tc.err, err))
			}
			if tc.err {
				return
			}

			if tc.out != isIn {
				t.Errorf("check inst in merkle tree error, expect %v, got %v", tc.out, isIn)
			}
		})
	}
}

func buildMerklePathTestcase(numInst, startNodeID, leafID int) *merklePath {
	data := randomMerkleHashes(numInst)
	mp := buildInstructionMerklePath(data, startNodeID)
	if leafID < 0 {
		// Randomize 32 bytes
		h := randomMerkleHashes(1)
		mp.leaf = toByte32(h[0])
	} else {
		mp.leaf = toByte32(mp.merkles[leafID])
	}
	return mp
}

func buildInstructionMerklePath(data [][]byte, startNodeID int) *merklePath {
	merkles := blockchain.BuildKeccak256MerkleTree(data)
	p, l := blockchain.GetKeccak256MerkleProofFromTree(merkles, startNodeID)
	path := [][32]byte{}
	left := []bool{}
	for i, x := range p {
		path = append(path, toByte32(x))
		left = append(left, l[i])
	}

	return &merklePath{
		merkles: merkles,
		leaf:    toByte32(merkles[startNodeID]),
		root:    toByte32(merkles[len(merkles)-1]),
		path:    path,
		left:    left,
	}
}

func randomMerkleHashes(n int) [][]byte {
	h := [][]byte{}
	for i := 0; i < n; i++ {
		b := make([]byte, 32)
		rand.Read(b)
		h = append(h, b)
	}
	return h
}

type merklePath struct {
	merkles [][]byte
	leaf    [32]byte
	root    [32]byte
	path    [][32]byte
	left    []bool
}

func TestFixedIncognitoProxyConstructor(t *testing.T) {
	p, _, _ := setupFixedCommittee()
	comm, err := p.inc.BeaconCommittees(nil, big.NewInt(0))
	if err != nil {
		t.Error(err)
	}
	if comm.StartBlock.Uint64() != uint64(0) {
		t.Errorf("incorrect startBlock, expect 0, got %d", comm.StartBlock)
	}
	comm, err = p.inc.BridgeCommittees(nil, big.NewInt(0))
	if err != nil {
		t.Error(err)
	}
	if comm.StartBlock.Uint64() != uint64(0) {
		t.Errorf("incorrect startBlock, expect 0, got %d", comm.StartBlock)
	}
}

func buildDecodedSwapConfirmInst(meta, shard, height, swapID int, addrs []string) []byte {
	a := []byte{}
	for _, addr := range addrs {
		d, _ := hex.DecodeString(addr)
		a = append(a, toBytes32BigEndian(d)...)
	}
	decoded := []byte{byte(meta)}
	decoded = append(decoded, byte(shard))
	decoded = append(decoded, toBytes32BigEndian(big.NewInt(int64(height)).Bytes())...)
	decoded = append(decoded, toBytes32BigEndian(big.NewInt(int64(len(addrs))).Bytes())...)
	decoded = append(decoded, toBytes32BigEndian(big.NewInt(int64(swapID)).Bytes())...)
	decoded = append(decoded, a...)
	return decoded
}

func getFixedSwapBridgeProof() *decodedProof {
	proofMarshalled := `{"Instruction":"RwEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFAAAAAAAAAAAAAAAAPHgSR4Po450eCE/d0OCXM0ui2UUAAAAAAAAAAAAAAAB2402KUnlhKG5VUyYgr1uE88ZTjwAAAAAAAAAAAAAAAGhobbaHRYjSQEFV0Apz+CpQ/dGQAAAAAAAAAAAAAAAAFTOsTSkiwVBVHy9dwrDB7eOCuJAAAAAAAAAAAAAAAAApvTTlHqui8K9X+Wq3yAn5EP+ItA==","BeaconHeight":null,"BridgeHeight":null,"InstPaths":[[[88,140,70,183,2,23,101,172,69,86,34,176,184,117,64,160,85,82,228,181,203,117,123,125,103,133,131,40,37,146,221,45],[255,180,222,138,116,33,250,156,148,107,1,251,2,193,86,175,149,35,124,138,118,164,80,32,160,18,232,179,23,153,29,138]],[[77,131,10,179,224,59,122,235,216,18,200,37,72,214,28,194,124,156,152,68,191,168,120,213,233,125,18,1,7,231,120,29]]],"InstPathIsLefts":[[false,false],[true]],"InstRoots":[[246,33,171,205,186,148,38,13,162,221,242,161,124,79,164,114,243,145,134,185,213,231,248,31,146,95,157,166,182,119,207,159],[180,112,69,218,29,1,113,71,57,188,201,181,133,160,215,199,181,199,206,124,81,228,14,17,103,32,170,202,142,158,137,187]],"BlkData":[[162,231,15,82,145,236,218,38,82,74,53,255,21,106,108,255,49,25,186,18,220,223,225,21,148,115,60,223,174,153,25,219],[163,34,83,153,105,147,63,92,129,38,131,191,118,215,180,33,183,78,159,54,169,216,222,92,116,242,40,145,183,176,121,195]],"SigIdxs":[[0,1,2,3],[0,1,2,3]],"SigVs":["HBsbGw==","GxwcHA=="],"SigRs":[[[122,158,222,191,249,77,177,45,79,125,211,76,100,66,99,84,123,188,206,22,43,136,245,204,4,196,103,106,137,193,194,240],[171,198,85,94,254,163,51,184,184,83,82,33,95,89,35,217,159,7,246,236,26,235,149,157,239,229,54,56,159,4,243,218],[40,72,83,120,178,178,158,218,12,198,235,219,76,120,230,107,242,60,97,120,232,250,22,63,111,108,212,236,112,37,112,56],[92,64,94,182,180,147,220,249,220,98,20,57,24,89,185,24,137,222,52,0,45,58,132,236,144,194,99,55,225,229,171,176]],[[241,220,136,208,241,17,4,61,189,221,87,249,132,229,40,147,190,105,44,154,227,71,37,136,244,248,165,118,105,28,234,34],[230,19,246,136,131,47,99,248,182,122,208,89,128,253,194,149,35,224,127,70,166,54,251,38,113,110,148,60,52,143,40,138],[147,197,97,112,153,58,202,80,138,238,197,91,40,110,21,187,118,8,131,56,189,13,216,122,97,105,172,133,6,72,12,107],[143,164,230,153,92,129,175,244,143,250,101,62,6,168,174,128,108,202,159,46,52,209,248,52,29,89,203,204,80,103,37,194]]],"SigSs":[[[71,191,26,215,147,217,74,139,212,164,210,176,198,29,231,135,208,189,53,131,126,137,22,22,126,39,164,218,11,114,211,103],[81,14,15,166,105,139,180,202,32,164,27,13,132,108,13,221,82,136,57,91,20,76,104,33,160,102,185,19,99,62,122,28],[63,124,97,144,229,31,134,239,62,254,141,43,14,251,84,64,217,79,125,200,72,98,67,244,197,194,124,16,232,51,126,101],[121,206,180,155,162,60,187,253,97,42,198,174,80,56,130,237,25,237,198,218,151,85,176,95,2,222,208,65,51,89,50,179]],[[26,22,150,125,176,0,110,250,194,183,122,233,35,161,235,66,250,71,158,154,141,82,127,83,44,192,190,199,119,109,52,62],[85,175,185,87,227,57,252,47,170,105,170,147,215,245,102,143,143,251,163,48,126,18,26,68,200,17,223,178,63,59,110,190],[114,53,172,67,216,64,249,23,254,202,14,50,56,181,72,183,196,47,245,38,237,184,212,57,62,246,109,241,135,42,234,41],[13,119,178,123,51,14,72,185,202,141,83,32,153,165,213,118,169,66,113,118,47,15,204,26,142,221,70,252,132,249,73,78]]]}`
	proof := &decodedProof{}
	json.Unmarshal([]byte(proofMarshalled), proof)
	return proof
}

func getFixedSwapBeaconProof() *decodedProof {
	proofMarshalled := `{"Instruction":"RgEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAm8D657tDKCh1m245HgzJmZUFd5EAAAAAAAAAAAAAAABsvCk3/uR3u9o2CoQu7r+Swvq2EwAAAAAAAAAAAAAAAMq/Pbk+tIph1BSGrMkoG2JAQRQDAAAAAAAAAAAAAAAAKb005R6rovCvV/lqt8gJ+RD/iLQ=","BeaconHeight":null,"BridgeHeight":null,"InstPaths":[[[190,127,252,151,179,86,125,52,148,38,88,71,42,122,97,139,203,231,174,68,80,117,88,56,239,128,172,245,206,95,53,178],[162,228,131,39,165,165,2,156,136,240,223,91,73,46,190,153,146,202,26,238,236,18,158,126,177,82,79,156,212,10,53,210]],[]],"InstPathIsLefts":[[true,true],[]],"InstRoots":[[249,163,147,109,61,102,214,159,153,210,254,167,250,0,106,47,239,2,166,192,182,187,8,184,238,217,70,60,130,211,61,59],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]],"BlkData":[[10,7,159,166,138,29,232,54,106,83,29,231,41,186,189,17,178,237,217,254,243,137,208,33,50,61,20,187,191,149,70,231],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]],"SigIdxs":[[0,1,2,3],[]],"SigVs":["GxsbHA==",""],"SigRs":[[[165,220,162,117,239,247,151,155,228,191,90,68,192,6,196,77,198,36,121,66,99,137,213,29,38,37,157,36,124,206,194,247],[36,68,143,168,167,222,194,213,53,167,221,217,46,130,92,182,182,202,148,83,193,157,119,206,248,168,61,177,215,228,79,180],[189,188,89,137,208,221,155,133,233,186,109,105,240,42,87,56,94,213,128,242,157,110,147,69,203,139,70,238,166,130,168,98],[6,116,33,87,1,135,160,111,69,195,162,75,72,156,177,129,167,180,20,219,197,192,83,135,121,80,116,164,240,63,214,94]],[]],"SigSs":[[[60,52,159,94,158,253,102,75,1,176,21,223,186,241,110,246,36,174,167,155,202,221,109,152,104,223,34,191,173,116,189,149],[63,171,88,18,252,67,227,4,95,243,181,130,226,53,225,227,220,209,31,72,139,127,35,237,164,214,248,105,22,231,118,134],[13,203,155,190,14,195,49,223,120,18,255,159,32,52,195,35,136,77,222,27,85,179,41,183,126,55,52,223,50,195,85,106],[85,192,199,155,76,34,126,82,252,53,138,243,213,220,183,150,27,70,88,199,90,192,56,200,25,201,112,71,87,207,216,171]],[]]}`
	proof := &decodedProof{}
	json.Unmarshal([]byte(proofMarshalled), proof)
	return proof
}
