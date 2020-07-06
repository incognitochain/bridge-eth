pragma solidity ^0.6.6;
pragma experimental ABIEncoderV2;

import "./pause.sol";

/**
 * @dev Stores beacon and bridge committee members of Incognito Chain. Other
 * contracts can query this contract to check if an instruction is confimed on
 * Incognito
 */
contract IncognitoProxy is AdminPausable {
    struct MerkleProof {
        bytes32[] path;
        bool[] isLeft;
    }

    struct InstructionProof {
        bytes inst;
        bytes32[] path;
        bool[] isLeft;
        bytes32 root;
        bytes32 blkData;
        uint[] sigIdx;
        uint8[] sigV;
        bytes32[] sigR;
        bytes32[] sigS;
    }

    struct CommitteeMeta {
        uint8 meta;
        uint8 shard;
        uint startHeight;
        uint numVals;
        uint id;
    }

    struct Committee {
        address[] pubkeys; // ETH address of all members
        uint startBlock; // The block that the committee starts to work on
        uint swapID;
    }

    struct Candidate {
        address[] pubkeys;
        uint startBlock;
        bytes32 blockHash;
    }

    struct Finality {
        uint blockHeight; // TODO: remove if unnecessary
        bytes32 rootHash;
    }

    Committee[] public beaconCommittees; // All beacon committees from genesis block
    Committee[] public bridgeCommittees; // All bridge committees from genesis block
    mapping(uint => Candidate) public bridgeCandidates;
    mapping(uint => Candidate) public beaconCandidates;

    // Finality info for beacon and bridge
    Finality public beaconFinality;
    Finality public bridgeFinality;

    event SubmittedBridgeCandidate(uint id, uint startHeight);
    event BeaconCommitteeSwapped(uint id, uint startHeight);
    event BridgeCommitteeSwapped(uint id, uint startHeight);
    event ChainFinalized(bool isBeacon);
    event CandidatePromoted(uint swapID, bool isBeacon);

    /**
     * @dev Sets the genesis committees and the address of admin
     * @notice Admin is the one responsible for the contract in case of emergency
     * Here, they are authorized to Pause the contract, stopping new committees
     * from being added to the contract
     * Admin is authorized to Pause the contract at anytime for 1 year starting
     * from the moment the contract is deployed
     * Admin is also authorized to increase the expiration time if they need more
     * time to implement a more decentralized failsafe mechanism
     * @notice Admin can also be a smart contract implementing a DAO and making decisions through a voting system
     * @param admin: ETH address
     * @param beaconCommittee: genesis committee members of beacon chain
     * @param bridgeCommittee: genesis committee members of bridge
     */
    constructor(
        address admin,
        address[] memory beaconCommittee,
        address[] memory bridgeCommittee
    ) public AdminPausable(admin) {
        beaconCommittees.push(Committee({
            pubkeys: beaconCommittee,
            startBlock: 0,
            swapID: 0
        }));

        bridgeCommittees.push(Committee({
            pubkeys: bridgeCommittee,
            startBlock: 0,
            swapID: 0
        }));
    }

    /**
     * @dev Gets a beacon committee in the past
     * @notice We need to implement this because the autogenerated getter returns only the startBlock
     * @param i index of the committee to get
     * @return the committee and their startBlock
     */
    function getBeaconCommittee(uint i) public view returns(Committee memory) {
        return beaconCommittees[i];
    }

    /**
     * @dev Gets a bridge committee in the past
     * @notice the same as getBeaconCommittee but for bridge
     */
    function getBridgeCommittee(uint i) public view returns(Committee memory) {
        return bridgeCommittees[i];
    }

    /**
     * @dev Validates and stores a new group of bridge candidates
     * The candidates will become a committee when a finality proof is provided
     * @notice This function takes a swap instruction on Incognito Chain, checks for its validity and stores the candidates
     * @notice This only works when the contract is not Paused
     * @notice All params except inst are the list of 2 elements corresponding to the proof on beacon and bridge
     */

    function submitBridgeCandidate(
        bytes memory inst,
        InstructionProof[2] memory instProofs
    ) public isNotPaused {
        // Parse instruction and check metadata
        CommitteeMeta memory cm; // Temp var to by pass max local var in a function
        (cm.meta, cm.shard, cm.startHeight, cm.numVals, cm.id) = extractMetaFromInstruction(inst);
        require(cm.meta == 71 && cm.shard == 1);

        // Verify instruction on beacon
        // NOTE: assuming no swap candidate for beacon
        bytes32 instHash = keccak256(inst);
        require(instructionApproved(
            true,
            instHash,
            beaconCommittees[beaconCommittees.length-1].startBlock,
            instProofs[0].path,
            instProofs[0].isLeft,
            instProofs[0].root,
            instProofs[0].blkData,
            instProofs[0].sigIdx,
            instProofs[0].sigV,
            instProofs[0].sigR,
            instProofs[0].sigS
        ));

        // Verify instruction on bridge
        address[] memory signers;
        uint latestSwapID = bridgeCommittees[bridgeCommittees.length-1].swapID;
        if (latestSwapID + 1 < cm.id) {
            signers = filterSigners(instProofs[1].sigIdx, bridgeCandidates[cm.id-1].pubkeys);
        } else {
            signers = filterSigners(instProofs[1].sigIdx, bridgeCommittees[bridgeCommittees.length-1].pubkeys);
        }
        require(instructionApprovedBySigners(
            instHash,
            signers,
            instProofs[1].path,
            instProofs[1].isLeft,
            instProofs[1].root,
            instProofs[1].blkData,
            instProofs[1].sigV,
            instProofs[1].sigR,
            instProofs[1].sigS
        ));

        // Make sure 1 instruction can't be used twice (using startHeight)
        require(cm.startHeight > bridgeCommittees[bridgeCommittees.length-1].startBlock, "cannot change old committee");

        // Store candidates
        signers = extractCommitteeFromInstruction(inst, cm.numVals);
        bytes32 blk = keccak256(abi.encodePacked(keccak256(abi.encodePacked(instProofs[1].blkData, instProofs[1].root))));
        bridgeCandidates[cm.id] = Candidate({
            pubkeys: signers,
            startBlock: cm.startHeight,
            blockHash: blk
        });

        emit SubmittedBridgeCandidate(bridgeCommittees.length, cm.startHeight);
        // emit BridgeCommitteeSwapped(bridgeCommittees.length, cm.startHeight);
    }

    /**
     * @dev Updates the latest committee of the beacon chain
     * @notice This function takes a swap instruction on Incognito Chain, checks for its validity and stores the latest committee
     * @notice This only works when the contract is not Paused
     * @notice Swapping beacon committee doesn't require that the instruction is included in the bridge chain
     * @notice All params are the same as swapBridgeCommittee
     */
    function submitBeaconCandidate(
        bytes memory inst,
        InstructionProof memory instProof
    ) public isNotPaused {
        bytes32 instHash = keccak256(inst);

        // Verify instruction on beacon
        require(instructionApproved(
            true,
            instHash,
            beaconCommittees[beaconCommittees.length-1].startBlock,
            instProof.path,
            instProof.isLeft,
            instProof.root,
            instProof.blkData,
            instProof.sigIdx,
            instProof.sigV,
            instProof.sigR,
            instProof.sigS
        ));

        // Parse instruction and check metadata and shardID
        (uint8 meta, uint8 shard, uint startHeight, uint numVals, uint id) = extractMetaFromInstruction(inst);
        require(meta == 70 && shard == 1);

        // Make sure 1 instruction can't be used twice (using startHeight)
        require(startHeight > beaconCommittees[beaconCommittees.length-1].startBlock, "cannot change old committee");

        // Store candidates
        address[] memory pubkeys = extractCommitteeFromInstruction(inst, numVals);
        bytes32 blk = keccak256(abi.encodePacked(keccak256(abi.encodePacked(instProof.blkData, instProof.root))));
        beaconCandidates[id] = Candidate({
            pubkeys: pubkeys,
            startBlock: startHeight,
            blockHash: blk
        });

        emit SubmittedBridgeCandidate(beaconCommittees.length, startHeight);
        // emit BeaconCommitteeSwapped(beaconCommittees.length, startHeight);
    }

    // TODO: doc
    // TODO: split func
    function submitFinalityProof(
        InstructionProof[] memory instProofs,
        uint swapID,
        bool isCandidate,
        bool isBeacon
    ) public isNotPaused {
        require(instProofs.length == 2, "must provide 2 blocks for finality proof");

        // Extract the committee signed the instructions
        // Using the same committee for both blocks
        // We do not support two adjacent blocks with increasing timeslots but signed by 2 different committees
        address[] memory signers;
        if (isCandidate) {
            if (isBeacon) {
                signers = beaconCandidates[swapID].pubkeys;
            } else {
                signers = bridgeCandidates[swapID].pubkeys;
            }
        } else {
            if (isBeacon) {
                signers = beaconCommittees[beaconCommittees.length-1].pubkeys;
            } else {
                signers = bridgeCommittees[bridgeCommittees.length-1].pubkeys;
            }
        }

        // Check if both BlockMerkleRoot instructions are valid
        for (uint i = 0; i < 2; i++) {
            // Extract signers that signed this block (require sigIdx to be strictly increasing)
            address[] memory signersTmp = filterSigners(instProofs[i].sigIdx, signers);

            require(instructionApprovedBySigners(
                keccak256(instProofs[i].inst),
                signersTmp,
                instProofs[i].path,
                instProofs[i].isLeft,
                instProofs[i].root,
                instProofs[i].blkData,
                instProofs[i].sigV,
                instProofs[i].sigR,
                instProofs[i].sigS
            ));
        }

        // Validate instruction's data
        (uint8 meta0, bytes32 rootHash, uint proposeTime0) = extractDataFromBlockMerkleInstruction(instProofs[0].inst);
        (uint8 meta1, bytes32 _, uint proposeTime1) = extractDataFromBlockMerkleInstruction(instProofs[1].inst);
        require(proposeTime0 / 10 + 1 == proposeTime1 / 10, "proposeTime invalid");
        require(meta0 == 1 && meta1 == 1, "invalid meta");

        // Save the block merkle root
        if (isBeacon) {
            beaconFinality = Finality({
                blockHeight: 0,
                rootHash: rootHash
            });
        } else {
            bridgeFinality = Finality({
                blockHeight: 0,
                rootHash: rootHash
            });
        }
        emit ChainFinalized(isBeacon);
    }

    // TODO: doc
    function promoteCandidate(
        uint swapID,
        bool isBeacon,
        MerkleProof memory proof
    ) public isNotPaused {
        // Extract block data
        bytes32 blockHash;
        bytes32 blockRoot;
        if (isBeacon) {
            blockHash = beaconCandidates[swapID].blockHash;
            blockRoot = beaconFinality.rootHash;
        } else {
            blockHash = bridgeCandidates[swapID].blockHash;
            blockRoot = bridgeFinality.rootHash;
        }

        // Check that block is in merkle tree
        require(leafInMerkleTree(
            blockHash,
            blockRoot,
            proof.path,
            proof.isLeft
        ));

        // Move candidate to committee
        if (isBeacon) {
            // This must be the next swapID
            require(beaconCommittees[beaconCommittees.length-1].swapID + 1 == swapID, "must promote candidate sequentially");

            beaconCommittees.push(Committee({
                pubkeys: beaconCandidates[swapID].pubkeys,
                startBlock: beaconCandidates[swapID].startBlock,
                swapID: swapID
            }));
        } else {
            // This must be the next swapID
            require(bridgeCommittees[bridgeCommittees.length-1].swapID + 1 == swapID, "must promote candidate sequentially");

            bridgeCommittees.push(Committee({
                pubkeys: bridgeCandidates[swapID].pubkeys,
                startBlock: bridgeCandidates[swapID].startBlock,
                swapID: swapID
            }));
        }
        emit CandidatePromoted(swapID, isBeacon);
    }

    /**
     * @dev Checks if an instruction is confirmed on chain (beacon or bridge)
     * @notice A confirmation means that the instruction is included in a block
     * that has enough validators' signatures
     * @param isBeacon: check on beacon or bridge
     * @param instHash: keccak256 hash of the instruction's content
     * @param blkHeight: height of the block containing the instruction
     * @param instPath: merkle path of the instruction
     * @param instPathIsLeft: whether each node on the path is the left or right child
     * @param instRoot: root of the merkle tree contains all instructions
     * @param blkData: merkle has of the block body
     * @param sigIdx: indices of the validators who signed this block
     * @param sigV: part of the signatures of the validators
     * @param sigR: part of the signatures of the validators
     * @param sigS: part of the signatures of the validators
     * @return bool: whether the instruction is valid and confirmed
     */
    function instructionApproved(
        bool isBeacon,
        bytes32 instHash,
        uint blkHeight,
        bytes32[] memory instPath,
        bool[] memory instPathIsLeft,
        bytes32 instRoot,
        bytes32 blkData,
        uint[] memory sigIdx,
        uint8[] memory sigV,
        bytes32[] memory sigR,
        bytes32[] memory sigS
    ) public view returns (bool) {
        // Find committee in charge of this block
        address[] memory signers;
        uint _;
        if (isBeacon) {
            (signers, _) = findBeaconCommitteeFromHeight(blkHeight);
        } else {
            (signers, _) = findBridgeCommitteeFromHeight(blkHeight);
        }

        // Extract signers that signed this block (require sigIdx to be strictly increasing)
        signers = filterSigners(sigIdx, signers);

        return instructionApprovedBySigners(
            instHash,
            signers,
            instPath,
            instPathIsLeft,
            instRoot,
            blkData,
            sigV,
            sigR,
            sigS
        );
    }

    function instructionApprovedBySigners(
        bytes32 instHash,
        address[] memory signers,
        bytes32[] memory instPath,
        bool[] memory instPathIsLeft,
        bytes32 instRoot,
        bytes32 blkData,
        uint8[] memory sigV,
        bytes32[] memory sigR,
        bytes32[] memory sigS
    ) public view returns (bool) {
        // Get double block hash from instRoot and other data
        bytes32 blk = keccak256(abi.encodePacked(keccak256(abi.encodePacked(blkData, instRoot))));

        // Check if enough validators signed this block
        if (sigV.length <= signers.length * 2 / 3) {
            return false;
        }

        // Check that signature is correct
        require(verifySig(signers, blk, sigV, sigR, sigS));

        // Check that inst is in block
        require(leafInMerkleTree(
            instHash,
            instRoot,
            instPath,
            instPathIsLeft
        ));

        return true;
    }

    /**
     * @dev Finds the beacon committee in charge of signing a block height
     * @notice This functions does a binary search of all committees (since genesis block)
     * @param blkHeight: to search for
     * @return committee: address of the committee members
     * @return id: index of the committee
     */
    function findBeaconCommitteeFromHeight(uint blkHeight) public view returns (address[] memory, uint) {
        uint l = 0;
        uint r = beaconCommittees.length;
        require(r > 0);
        r = r - 1;
        while (l != r) {
            uint m = (l + r + 1) / 2;
            if (beaconCommittees[m].startBlock <= blkHeight) {
                l = m;
            } else {
                r = m - 1;
            }
        }
        return (beaconCommittees[l].pubkeys, l);
    }

    /**
     * @dev Finds the bridge committee in charge of signing a block height
     * @notice The same as findBeaconCommitteeFromHeight but for bridge chain
     */
    function findBridgeCommitteeFromHeight(uint blkHeight) public view returns (address[] memory, uint) {
        uint l = 0;
        uint r = bridgeCommittees.length;
        require(r > 0);
        r = r - 1;
        while (l != r) {
            uint m = (l + r + 1) / 2;
            if (bridgeCommittees[m].startBlock <= blkHeight) {
                l = m;
            } else {
                r = m - 1;
            }
        }
        return (bridgeCommittees[l].pubkeys, l);
    }

    /**
     * @dev Checks if a value is in a merkle tree
     * @param leaf: the value to check
     * @param root: of the merkle tree
     * @param path: merkle path of the value to check
     * @param left: whether each node on the path is the left or right child
     * @return bool: whether the value is in the merkle tree
     */
    function leafInMerkleTree(
        bytes32 leaf,
        bytes32 root,
        bytes32[] memory path,
        bool[] memory left
    ) public pure returns (bool) {
        require(left.length == path.length);
        bytes32 hash = leaf;
        for (uint i = 0; i < path.length; i++) {
            if (left[i]) {
                hash = keccak256(abi.encodePacked(path[i], hash));
            } else if (path[i] == 0x0) {
                hash = keccak256(abi.encodePacked(hash, hash));
            } else {
                hash = keccak256(abi.encodePacked(hash, path[i]));
            }
        }
        return hash == root;
    }

    /**
     * @dev Extracts the metadata of a swap instruction
     * @param inst: the full instruction, containing both metadata and body
     * @return meta: type of the instruction, 70 for swapping beacon and 71 for bridge
     * @return shard: ID of the Incognito shard containing the instruction, must be 1
     * @return height: the starting block that the committee is responsible for
     * @return numVals: number of validators in the new committee
     * @return id: id of the swap
     */
    function extractMetaFromInstruction(bytes memory inst) public pure returns(uint8, uint8, uint, uint, uint) {
        require(inst.length >= 0x42); // 0x02 bytes for meta and shard, 0x20 each for height and numVals
        uint8 meta = uint8(inst[0]);
        uint8 shard = uint8(inst[1]);
        uint height;
        uint numVals;
        uint id;
        assembly {
            // skip first 0x20 bytes (stored length of inst)
            height := mload(add(inst, 0x22)) // [2:34]
            numVals := mload(add(inst, 0x42)) // [34:66]
            id := mload(add(inst, 0x62)) // [66:98]
        }
        return (meta, shard, height, numVals, id);
    }

    /**
     * @dev Extracts the committee (body) from a swap instruction
     * @param inst: the full instruction, containing both metadata and body
     * @param numVals: number of validators in the new committee
     * @return committee: address of the committee members
     */
    function extractCommitteeFromInstruction(bytes memory inst, uint numVals) public pure returns (address[] memory) {
        require(inst.length == 0x62 + numVals * 0x20);
        address[] memory addr = new address[](numVals);
        address tmp;
        for (uint i = 0; i < numVals; i++) {
            assembly {
                // skip first 0x20 bytes (stored length of inst)
                // also, skip the next 0x62 bytes (stored metadata)
                tmp := mload(add(add(inst, 0x82), mul(i, 0x20))) // 98+i*32
            }
            addr[i] = tmp;
        }
        return addr;
    }

    function extractDataFromBlockMerkleInstruction(bytes memory inst) public pure returns (uint8, bytes32, uint) {
        require(inst.length >= 0x41); // 0x01 bytes for meta and shard, 0x20 each for rootHash and proposeTime
        uint8 meta = uint8(inst[0]);
        bytes32 rootHash;
        uint proposeTime;
        assembly {
            // skip first 0x20 bytes (stored length of inst)
            rootHash := mload(add(inst, 0x21)) // [1:33]
            proposeTime := mload(add(inst, 0x41)) // [33:65]
        }
        return (meta, rootHash, proposeTime);
    }

    // TODO: doc
    function filterSigners(uint[] memory sigIdx, address[] memory signers) public pure returns (address[] memory) {
        address[] memory signersTmp = new address[](signers.length);
        for (uint i = 0; i < sigIdx.length; i++) {
            if ((i > 0 && sigIdx[i] <= sigIdx[i-1]) || sigIdx[i] >= signers.length) {
                revert("sigIdx invalid");
            }
            signersTmp[i] = signers[sigIdx[i]];
        }
        return signersTmp;
    }

    /**
     * @dev Verifies that the signatures for a message are correct
     * @param msgHash: the message to be verify
     * @param v: part of the signatures
     * @param r: part of the signatures
     * @param s: part of the signatures
     * @return bool: whether all signatures are correct
     */
    function verifySig(
        address[] memory committee,
        bytes32 msgHash,
        uint8[] memory v,
        bytes32[] memory r,
        bytes32[] memory s
    ) public pure returns (bool) {
        require(v.length == r.length);
        require(v.length == s.length);
        for (uint i = 0; i < v.length; i++){
            if (ecrecover(msgHash, v[i], r[i], s[i]) != committee[i]) {
                return false;
            }
        }
        return true;
    }
}
