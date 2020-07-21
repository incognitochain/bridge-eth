package jsonresult

type GetInstructionProof struct {
	Instruction string // Hex-encoded swap inst

	BeaconInstPath []string // Hex encoded path of the inst in merkle tree
	BeaconBlkData  string   // Hex encoded hash of the block meta
	BeaconInstID   int64    // Index of the instruction
	BeaconSigs     []string // Hex encoded signature (r, s, v)
	BeaconSigIdxs  []int    // Idxs of signer

	BridgeInstPath []string
	BridgeInstID   int64
	BridgeBlkData  string
	BridgeSigs     []string
	BridgeSigIdxs  []int
}
