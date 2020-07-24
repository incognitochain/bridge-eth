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

type GetFinalityProof struct {
	SwapID       uint64
	Instructions [2]string   // BlockMerkleRootMeta instruction
	InstPaths    [2][]string // Proofs that the above insts are in the blocks
	IDs          [2]int64    // Index of the instruction in the block

	BlkData [2]string
	Sigs    [2][]string // Hex encoded signature (r, s, v)
	SigIdxs [2][]int    // Idxs of signer
}
