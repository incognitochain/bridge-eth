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
	Result jsonresult.GetInstructionProof
	Error  struct {
		Code       int
		Message    string
		StackTrace string
	}
}

type decodedProof struct {
	Instruction []byte
	Heights     [2]*big.Int

	InstPaths       [2][][32]byte
	InstPathIsLefts [2][]bool
	InstRoots       [2][32]byte
	BlkData         [2][32]byte
	SigIdxs         [2][]*big.Int
	SigVs           [2][]uint8
	SigRs           [2][][32]byte
	SigSs           [2][][32]byte
}

func getAndDecodeBurnProof(txID string) (*decodedProof, error) {
	body := getBurnProof(txID)
	if len(body) < 1 {
		return nil, fmt.Errorf("burn proof not found")
	}

	r := getProofResult{}
	err := json.Unmarshal([]byte(body), &r)
	if err != nil {
		return nil, err
	}
	return decodeProof(&r)
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

	r := getProofResult{}
	err = json.Unmarshal([]byte(body), &r)
	if err != nil {
		return nil, err
	}
	return decodeProof(&r)
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
	url := "https://mainnet.incognito.org/fullnode"

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

func decodeProof(r *getProofResult) (*decodedProof, error) {
	inst := decode(r.Result.Instruction)
	fmt.Printf("inst: %d %x\n", len(inst), inst)
	fmt.Printf("instHash (isWithdrawed, without height): %x\n", keccak256(inst))

	// Block heights
	beaconHeight := big.NewInt(0).SetBytes(decode(r.Result.BeaconHeight))
	bridgeHeight := big.NewInt(0).SetBytes(decode(r.Result.BridgeHeight))
	heights := [2]*big.Int{beaconHeight, bridgeHeight}
	fmt.Printf("beaconHeight: %d\n", beaconHeight)
	fmt.Printf("bridgeHeight: %d\n", bridgeHeight)

	beaconInstRoot := decode32(r.Result.BeaconInstRoot)
	beaconInstPath := make([][32]byte, len(r.Result.BeaconInstPath))
	beaconInstPathIsLeft := make([]bool, len(r.Result.BeaconInstPath))
	for i, path := range r.Result.BeaconInstPath {
		beaconInstPath[i] = decode32(path)
		beaconInstPathIsLeft[i] = r.Result.BeaconInstPathIsLeft[i]
	}
	// fmt.Printf("beaconInstRoot: %x\n", beaconInstRoot)

	beaconBlkData := toByte32(decode(r.Result.BeaconBlkData))
	fmt.Printf("data: %s %s\n", r.Result.BeaconBlkData, r.Result.BeaconInstRoot)
	fmt.Printf("expected beaconBlkHash: %x\n", keccak256(beaconBlkData[:], beaconInstRoot[:]))

	beaconSigVs, beaconSigRs, beaconSigSs, err := decodeSigs(r.Result.BeaconSigs)
	if err != nil {
		return nil, err
	}

	beaconSigIdxs := []*big.Int{}
	for _, sIdx := range r.Result.BeaconSigIdxs {
		beaconSigIdxs = append(beaconSigIdxs, big.NewInt(int64(sIdx)))
	}

	// For bridge
	bridgeInstRoot := decode32(r.Result.BridgeInstRoot)
	bridgeInstPath := make([][32]byte, len(r.Result.BridgeInstPath))
	bridgeInstPathIsLeft := make([]bool, len(r.Result.BridgeInstPath))
	for i, path := range r.Result.BridgeInstPath {
		bridgeInstPath[i] = decode32(path)
		bridgeInstPathIsLeft[i] = r.Result.BridgeInstPathIsLeft[i]
	}
	// fmt.Printf("bridgeInstRoot: %x\n", bridgeInstRoot)
	bridgeBlkData := toByte32(decode(r.Result.BridgeBlkData))

	bridgeSigVs, bridgeSigRs, bridgeSigSs, err := decodeSigs(r.Result.BridgeSigs)
	if err != nil {
		return nil, err
	}

	bridgeSigIdxs := []*big.Int{}
	for _, sIdx := range r.Result.BridgeSigIdxs {
		bridgeSigIdxs = append(bridgeSigIdxs, big.NewInt(int64(sIdx)))
		// fmt.Printf("bridgeSigIdxs[%d]: %d\n", i, j)
	}

	// Merge beacon and bridge proof
	instPaths := [2][][32]byte{beaconInstPath, bridgeInstPath}
	instPathIsLefts := [2][]bool{beaconInstPathIsLeft, bridgeInstPathIsLeft}
	instRoots := [2][32]byte{beaconInstRoot, bridgeInstRoot}
	blkData := [2][32]byte{beaconBlkData, bridgeBlkData}
	sigIdxs := [2][]*big.Int{beaconSigIdxs, bridgeSigIdxs}
	sigVs := [2][]uint8{beaconSigVs, bridgeSigVs}
	sigRs := [2][][32]byte{beaconSigRs, bridgeSigRs}
	sigSs := [2][][32]byte{beaconSigSs, bridgeSigSs}

	return &decodedProof{
		Instruction:     inst,
		Heights:         heights,
		InstPaths:       instPaths,
		InstPathIsLefts: instPathIsLefts,
		InstRoots:       instRoots,
		BlkData:         blkData,
		SigIdxs:         sigIdxs,
		SigVs:           sigVs,
		SigRs:           sigRs,
		SigSs:           sigSs,
	}, nil
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
