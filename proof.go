package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/incognitochain/bridge-eth/bridge/incognito_proxy"
	"github.com/incognitochain/bridge-eth/bridge/vault"
	"github.com/incognitochain/bridge-eth/common/base58"
	"github.com/incognitochain/bridge-eth/consensus/signatureschemes/bridgesig"
	"github.com/incognitochain/bridge-eth/erc20"
	"github.com/incognitochain/bridge-eth/erc20/bnb"
	"github.com/incognitochain/bridge-eth/jsonresult"
	"github.com/pkg/errors"
)

type contracts struct {
	v         *vault.Vault
	vAddr     common.Address
	inc       *incognito_proxy.IncognitoProxy
	incAddr   common.Address
	token     *erc20.Erc20
	tokenAddr common.Address

	tokens       map[int]tokenInfo       // mapping from decimal => token
	customErc20s map[string]*TokenerInfo // mapping from name => token
}

type TokenerInfo struct {
	addr common.Address
	c    Tokener
}

type Tokener interface {
	BalanceOf(*bind.CallOpts, common.Address) (*big.Int, error)
	Approve(*bind.TransactOpts, common.Address, *big.Int) (*types.Transaction, error)
}

var _ Tokener = (*bnb.Bnb)(nil)

type tokenInfo struct {
	c    *erc20.Erc20
	addr common.Address
}

type getProofResult struct {
	Result interface{}
	Error  struct {
		Code       int
		Message    string
		StackTrace string
	}
}

type CandidateProof struct {
	Instruction []byte
	InstProofs  [2]incognito_proxy.IncognitoProxyInstructionProof
}

type FinalityProof struct {
	IsBeacon     bool
	SwapID       *big.Int
	Instructions [2][]byte
	InstProofs   [2]incognito_proxy.IncognitoProxyInstructionProof
}

type PromoteProof struct {
	IsBeacon    bool
	SwapID      *big.Int
	MerkleProof incognito_proxy.IncognitoProxyMerkleProof
}

type decodedProof struct {
	Instruction []byte

	InstPaths [2][][32]byte
	InstIDs   [2]*big.Int
	BlkData   [2][32]byte
	SigIdxs   [2][]*big.Int
	SigVs     [2][]uint8
	SigRs     [2][][32]byte
	SigSs     [2][][32]byte

	instRoots [2][32]byte // Temporary storage
}

func GetAndDecodeBridgeCandidateProof(url string, block int) (*CandidateProof, error) {
	proof, err := parseGetInstructionProofBody(getBridgeSwapProof(url, block))
	if err != nil {
		return nil, err
	}
	printCandidateInstruction(proof.Instruction)
	return buildCandidateProof(proof)
}

func printCandidateInstruction(inst []byte) {
	fmt.Println("=============================")
	fmt.Println("Bridge candidate instruction:")
	fmt.Printf("meta: %d\n", inst[0])
	fmt.Printf("shard: %d\n", inst[1])
	fmt.Printf("height: %d\n", big.NewInt(0).SetBytes(inst[2:34]))
	fmt.Printf("numVals: %d\n", big.NewInt(0).SetBytes(inst[34:66]))
	fmt.Printf("swapID: %d\n", big.NewInt(0).SetBytes(inst[66:98]))
	fmt.Println("=============================")
}

func GetAndDecodeBeaconCandidateProof(url string, block int) (*CandidateProof, error) {
	proof, err := parseGetInstructionProofBody(getBeaconSwapProof(url, block))
	if err != nil {
		return nil, err
	}
	return buildCandidateProof(proof)
}

func parseGetInstructionProofBody(body string) (*decodedProof, error) {
	if len(body) < 1 {
		return nil, fmt.Errorf("no bridge swap proof found")
	}
	r := getProofResult{Result: &jsonresult.GetInstructionProof{}}
	if err := json.Unmarshal([]byte(body), &r); err != nil {
		return nil, err
	}
	proof, err := decodeGetInstructionProof(r.Result.(*jsonresult.GetInstructionProof))
	if err != nil {
		return nil, err
	}
	return proof, nil
}

func buildCandidateProof(proof *decodedProof) (*CandidateProof, error) {
	inst, instProofs := buildIncognitoProxyInstructionProof(proof)
	return &CandidateProof{
		Instruction: inst,
		InstProofs:  instProofs,
	}, nil
}

func GetAndDecodeFinalityProof(url string, isBeacon bool, block int) (*FinalityProof, error) {
	shardID := 1
	if isBeacon {
		shardID = -1
	}
	encodedProof, err := parseGetFinalityProofBody(getFinalityProof(url, block, shardID))
	if err != nil {
		return nil, err
	}

	proof, err := decodeGetFinalityProof(encodedProof)
	if err != nil {
		return nil, err
	}
	proof.IsBeacon = isBeacon
	return proof, nil
}

func parseGetFinalityProofBody(body string) (*jsonresult.GetFinalityProof, error) {
	if len(body) < 1 {
		return nil, fmt.Errorf("no finality proof found")
	}
	r := getProofResult{Result: &jsonresult.GetFinalityProof{}}
	if err := json.Unmarshal([]byte(body), &r); err != nil {
		return nil, err
	}
	return r.Result.(*jsonresult.GetFinalityProof), nil
}

func GetAndDecodePromoteProof(
	url string,
	isBeacon bool,
	swapID *big.Int,
	ancestorBlock,
	anchorBlock int,
) (*PromoteProof, error) {
	shardID := -1 // Must prove finality on beacon
	encodedProof, err := parseGetAncestorProofBody(getAncestorProof(url, ancestorBlock, anchorBlock, shardID))
	if err != nil {
		return nil, err
	}

	ancestorProof, err := decodeGetAncestorProof(encodedProof)
	if err != nil {
		return nil, err
	}

	promoteProof := &PromoteProof{
		IsBeacon:    isBeacon,
		SwapID:      swapID,
		MerkleProof: ancestorProof,
	}
	return promoteProof, nil
}

func parseGetAncestorProofBody(body string) (*jsonresult.GetAncestorProof, error) {
	if len(body) < 1 {
		return nil, fmt.Errorf("no ancestor proof found")
	}
	r := getProofResult{Result: &jsonresult.GetAncestorProof{}}
	if err := json.Unmarshal([]byte(body), &r); err != nil {
		return nil, err
	}
	return r.Result.(*jsonresult.GetAncestorProof), nil
}

func buildIncognitoProxyInstructionProof(proof *decodedProof) ([]byte, [2]incognito_proxy.IncognitoProxyInstructionProof) {
	instProofs := [2]incognito_proxy.IncognitoProxyInstructionProof{}
	for i := 0; i < 2; i++ {
		instProofs[i] = incognito_proxy.IncognitoProxyInstructionProof{
			Path:    proof.InstPaths[i],
			Id:      proof.InstIDs[i],
			BlkData: proof.BlkData[i],
			SigIdx:  proof.SigIdxs[i],
			SigV:    proof.SigVs[i],
			SigR:    proof.SigRs[i],
			SigS:    proof.SigSs[i],
		}
	}
	return proof.Instruction, instProofs
}

func getFinalityProof(url string, block int, shardID int) string {
	payload := strings.NewReader(fmt.Sprintf("{\n    \"id\": 1,\n    \"jsonrpc\": \"1.0\",\n    \"method\": \"getfinalityproof\",\n    \"params\": [\n    \t%d,\n    \t%d\n]\n}", block, shardID))

	req, _ := http.NewRequest("POST", url, payload)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("err:", err)
		return ""
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}

func getAncestorProof(url string, ancestorBlock, anchorBlock int, shardID int) string {
	payload := strings.NewReader(fmt.Sprintf("{\n    \"id\": 1,\n    \"jsonrpc\": \"1.0\",\n    \"method\": \"getancestorproof\",\n    \"params\": [\n    \t%d,\n    \t%d,\n\t%d]\n}", shardID, ancestorBlock, anchorBlock))

	req, _ := http.NewRequest("POST", url, payload)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("err:", err)
		return ""
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
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

func getBeaconSwapProof(url string, block int) string {
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

func getAndDecodeBurnProof(txID string) (*decodedProof, error) {
	body := getBurnProof(txID)
	if len(body) < 1 {
		return nil, fmt.Errorf("burn proof not found")
	}

	r := getProofResult{Result: &jsonresult.GetInstructionProof{}}
	err := json.Unmarshal([]byte(body), &r)
	if err != nil {
		return nil, err
	}
	return decodeGetInstructionProof(r.Result.(*jsonresult.GetInstructionProof))
}

func getAndDecodeBurnProofV2(
	incBridgeHost string,
	txID string,
	rpcMethod string,
) (*decodedProof, error) {
	body, err := getBurnProofV2(incBridgeHost, txID, rpcMethod)
	if err != nil {
		return nil, err
	}
	if len(body) < 1 {
		return nil, fmt.Errorf("burn proof for deposit to SC not found")
	}

	r := getProofResult{Result: &jsonresult.GetInstructionProof{}}
	err = json.Unmarshal([]byte(body), &r)
	if err != nil {
		return nil, err
	}
	return decodeGetInstructionProof(r.Result.(*jsonresult.GetInstructionProof))
}

func getCommittee(url string) ([]common.Address, []common.Address, error) {
	payload := strings.NewReader("{\n    \"id\": 1,\n    \"jsonrpc\": \"1.0\",\n    \"method\": \"getbeaconbeststate\",\n    \"params\": []\n}")

	req, _ := http.NewRequest("POST", url, payload)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	type beaconBestStateResult struct {
		BeaconCommittee []string
		ShardCommittee  map[string][]string
	}

	type getBeaconBestStateResult struct {
		Result beaconBestStateResult
		Error  string
		Id     int
	}

	r := getBeaconBestStateResult{}
	err = json.Unmarshal([]byte(body), &r)
	if err != nil {
		return nil, nil, err
	}

	// Genesis committee
	beaconOld := make([]common.Address, len(r.Result.BeaconCommittee))
	for i, pk := range r.Result.BeaconCommittee {
		cpk := &CommitteePublicKey{}
		cpk.FromString(pk)
		addr, err := convertPubkeyToAddress(*cpk)
		if err != nil {
			return nil, nil, err
		}
		beaconOld[i] = addr
		fmt.Printf("beaconOld: %s\n", addr.Hex())
	}

	bridgeOld := make([]common.Address, len(r.Result.ShardCommittee["1"]))
	for i, pk := range r.Result.ShardCommittee["1"] {
		cpk := &CommitteePublicKey{}
		cpk.FromString(pk)
		addr, err := convertPubkeyToAddress(*cpk)
		if err != nil {
			return nil, nil, err
		}
		bridgeOld[i] = addr
		fmt.Printf("bridgeOld: %s\n", addr.Hex())
	}

	return beaconOld, bridgeOld, nil
}

func getBurnProof(txID string) string {
	// url := "http://127.0.0.1:9344"
	// url := "https://dev-test-node.incognito.org/"
	url := "http://51.83.36.184:9334"

	if len(txID) == 0 {
		txID = "87c89c1c19cec3061eff9cfefdcc531d9456ac48de568b3974c5b0a88d5f3834"
	}
	payload := strings.NewReader(fmt.Sprintf("{\n    \"id\": 1,\n    \"jsonrpc\": \"1.0\",\n    \"method\": \"getburnproof\",\n    \"params\": [\n    \t\"%s\"\n    ]\n}", txID))

	req, _ := http.NewRequest("POST", url, payload)

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

func getBurnProofV2(
	incBridgeHost string,
	txID string,
	rpcMethod string,
) (string, error) {
	if len(txID) == 0 {
		txID = "87c89c1c19cec3061eff9cfefdcc531d9456ac48de568b3974c5b0a88d5f3834"
	}
	payload := strings.NewReader(fmt.Sprintf("{\n    \"id\": 1,\n    \"jsonrpc\": \"1.0\",\n    \"method\": \"%s\",\n    \"params\": [\n    \t\"%s\"\n    ]\n}", rpcMethod, txID))

	req, _ := http.NewRequest("POST", incBridgeHost, payload)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func decodeGetInstructionProof(result *jsonresult.GetInstructionProof) (*decodedProof, error) {
	inst := decode(result.Instruction)
	fmt.Printf("inst: %d %x\n", len(inst), inst)
	// fmt.Printf("instHash (isWithdrawed, without height): %x\n", keccak256(inst))

	// Decode beacon data
	beaconInstPath := make([][32]byte, len(result.BeaconInstPath))
	for i, path := range result.BeaconInstPath {
		beaconInstPath[i] = decode32(path)
	}
	// fmt.Printf("beaconInstRoot: %x\n", beaconInstRoot)

	beaconInstID := big.NewInt(result.BeaconInstID)

	beaconBlkData := toByte32(decode(result.BeaconBlkData))

	beaconSigVs, beaconSigRs, beaconSigSs, err := decodeSigs(result.BeaconSigs)
	if err != nil {
		return nil, err
	}

	beaconSigIdxs := []*big.Int{}
	for _, sIdx := range result.BeaconSigIdxs {
		beaconSigIdxs = append(beaconSigIdxs, big.NewInt(int64(sIdx)))
	}

	// For bridge
	bridgeInstPath := make([][32]byte, len(result.BridgeInstPath))
	for i, path := range result.BridgeInstPath {
		bridgeInstPath[i] = decode32(path)
	}

	bridgeInstID := big.NewInt(result.BridgeInstID)

	// fmt.Printf("bridgeInstRoot: %x\n", bridgeInstRoot)
	bridgeBlkData := toByte32(decode(result.BridgeBlkData))

	bridgeSigVs, bridgeSigRs, bridgeSigSs, err := decodeSigs(result.BridgeSigs)
	if err != nil {
		return nil, err
	}

	bridgeSigIdxs := []*big.Int{}
	for _, sIdx := range result.BridgeSigIdxs {
		bridgeSigIdxs = append(bridgeSigIdxs, big.NewInt(int64(sIdx)))
		// fmt.Printf("bridgeSigIdxs[%d]: %d\n", i, j)
	}

	// Merge beacon and bridge proof
	instPaths := [2][][32]byte{beaconInstPath, bridgeInstPath}
	instIDs := [2]*big.Int{beaconInstID, bridgeInstID}
	blkData := [2][32]byte{beaconBlkData, bridgeBlkData}
	sigIdxs := [2][]*big.Int{beaconSigIdxs, bridgeSigIdxs}
	sigVs := [2][]uint8{beaconSigVs, bridgeSigVs}
	sigRs := [2][][32]byte{beaconSigRs, bridgeSigRs}
	sigSs := [2][][32]byte{beaconSigSs, bridgeSigSs}

	return &decodedProof{
		Instruction: inst,
		InstPaths:   instPaths,
		InstIDs:     instIDs,
		BlkData:     blkData,
		SigIdxs:     sigIdxs,
		SigVs:       sigVs,
		SigRs:       sigRs,
		SigSs:       sigSs,
	}, nil
}

func decodeGetFinalityProof(result *jsonresult.GetFinalityProof) (*FinalityProof, error) {
	insts := [2][]byte{}
	for i, inst := range result.Instructions {
		insts[i] = decode(inst)
	}

	instProofs := [2]incognito_proxy.IncognitoProxyInstructionProof{}
	for i := 0; i < 2; i++ {
		instPath := make([][32]byte, len(result.InstPaths[i]))
		for j, path := range result.InstPaths[i] {
			instPath[j] = decode32(path)
		}

		instID := big.NewInt(result.IDs[i])
		blkData := toByte32(decode(result.BlkData[i]))

		sigV, sigR, sigS, err := decodeSigs(result.Sigs[i])
		if err != nil {
			return nil, err
		}

		sigIdx := []*big.Int{}
		for _, sIdx := range result.SigIdxs[i] {
			sigIdx = append(sigIdx, big.NewInt(int64(sIdx)))
		}

		instProofs[i] = incognito_proxy.IncognitoProxyInstructionProof{
			Path:    instPath,
			Id:      instID,
			BlkData: blkData,
			SigIdx:  sigIdx,
			SigV:    sigV,
			SigR:    sigR,
			SigS:    sigS,
		}
	}

	swapID := big.NewInt(int64(result.SwapID))
	return &FinalityProof{
		SwapID:       swapID,
		Instructions: insts,
		InstProofs:   instProofs,
	}, nil
}

func decodeGetAncestorProof(result *jsonresult.GetAncestorProof) (incognito_proxy.IncognitoProxyMerkleProof, error) {
	proof := incognito_proxy.IncognitoProxyMerkleProof{
		Id:   big.NewInt(result.ID),
		Path: make([][32]byte, len(result.Path)),
	}

	for i, p := range result.Path {
		proof.Path[i] = decode32(p)
	}

	return proof, nil
}
func decodeSigs(sigs []string) (
	sigVs []uint8,
	sigRs [][32]byte,
	sigSs [][32]byte,
	err error,
) {
	sigVs = make([]uint8, len(sigs))
	sigRs = make([][32]byte, len(sigs))
	sigSs = make([][32]byte, len(sigs))
	for i, sig := range sigs {
		v, r, s, e := bridgesig.DecodeECDSASig(decode(sig))
		if e != nil {
			err = e
			return
		}
		sigVs[i] = uint8(v)
		copy(sigRs[i][:], r)
		copy(sigSs[i][:], s)
	}
	return
}

func toByte32(s []byte) [32]byte {
	a := [32]byte{}
	copy(a[:], s)
	return a
}

func decode(s string) []byte {
	d, _ := hex.DecodeString(s)
	return d
}

func decode32(s string) [32]byte {
	return toByte32(decode(s))
}

func keccak256(b ...[]byte) [32]byte {
	h := crypto.Keccak256(b...)
	r := [32]byte{}
	copy(r[:], h)
	return r
}

func convertPubkeyToAddress(cKey CommitteePublicKey) (common.Address, error) {
	pk, err := crypto.DecompressPubkey(cKey.MiningPubKey[BRI_CONSENSUS])
	if err != nil {
		return common.Address{}, errors.Wrapf(err, "cKey: %+v", cKey)
	}
	address := crypto.PubkeyToAddress(*pk)
	return address, nil
}

var BRI_CONSENSUS = "dsa"

type CommitteePublicKey struct {
	IncPubKey    []byte
	MiningPubKey map[string][]byte
}

func (pubKey *CommitteePublicKey) FromString(keyString string) error {
	keyBytes, ver, err := base58.Base58Check{}.Decode(keyString)
	if (ver != 0) || (err != nil) {
		return errors.New("Wrong input")
	}
	return json.Unmarshal(keyBytes, pubKey)
}
