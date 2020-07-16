package jsonresult

type GetInstructionProof struct {
	Instruction  string // Hex-encoded swap inst
	BeaconHeight string // Hex encoded height of the block contains the inst
	BridgeHeight string

	BeaconInstPath []string // Hex encoded path of the inst in merkle tree
	BeaconBlkData  string   // Hex encoded hash of the block meta
	BeaconSigs     []string // Hex encoded signature (r, s, v)
	BeaconSigIdxs  []int    // Idxs of signer

	BridgeInstPath []string
	BridgeBlkData  string
	BridgeSigs     []string
	BridgeSigIdxs  []int
}
