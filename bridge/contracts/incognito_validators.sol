pragma solidity ^0.6.6;
pragma experimental ABIEncoderV2;

import "./operable.sol";

contract IncognitoValidators is Operable {

    // isBeacon => address[group 0], address[group 1], ...
    mapping(bool => address[][]) public groupAddrs;

    struct Groups {
        uint oldGID;
        uint newGID;
    }

    // isBeacon => swapID => (oldGID, newGID)
    mapping(bool => mapping(uint => Groups)) public validators;

    // isBeacon => (oldGID, newGID)
    mapping(bool => Groups) public committee; // Latest committee

    constructor(
        address _admin,
        address _operator,
        address[] memory beaconCommittee,
        address[] memory bridgeCommittee
    ) public Operable(_admin, _operator) {
    }

    function loadGroups(bool isBeacon, uint oldGID, uint newGID) public view returns(address[] memory) {
        uint nAddrs = 0;
        if (oldGID > 0) {
            require(oldGID < groupAddrs[isBeacon].length);
            nAddrs += groupAddrs[isBeacon][oldGID].length;
        }
        if (newGID > 0) {
            require(newGID < groupAddrs[isBeacon].length);
            nAddrs += groupAddrs[isBeacon][newGID].length;
        }

        address[] memory addrs = new address[](nAddrs);
        uint j = 0;
        if (oldGID > 0) {
            for (uint i = 0; i < groupAddrs[isBeacon][oldGID].length; i++) {
                addrs[j++] = groupAddrs[isBeacon][oldGID][i];
            }
        }
        if (newGID > 0) {
            for (uint i = 0; i < groupAddrs[isBeacon][newGID].length; i++) {
                addrs[j++] = groupAddrs[isBeacon][newGID][i];
            }
        }
        return addrs;
    }

    function loadCandidates(bool isBeacon, uint swapID) public view returns(address[] memory) {
        return loadGroups(isBeacon, validators[isBeacon][swapID].oldGID, validators[isBeacon][swapID].newGID);
    }

    function storeCandidates(bool isBeacon, uint swapID, uint oldGID, uint newGID, address[] memory newAddrs) public onlyOperator {
        require(oldGID < groupAddrs[isBeacon].length);
        require(newGID == 0 || newGID == groupAddrs[isBeacon].length);

        // Save new group
        if (newGID > 0) {
            groupAddrs[isBeacon].push(newAddrs);
        }

        validators[isBeacon][swapID] = Groups({
            oldGID: oldGID,
            newGID: newGID
        });
    }

    function loadCommittee(bool isBeacon) public view returns(address[] memory) {
        return loadGroups(isBeacon, committee[isBeacon].oldGID, committee[isBeacon].newGID);
    }

    function storeCommittee(bool isBeacon, uint gID1, uint gID2) public onlyOperator {
        require(gID1 < 0 || gID1 < groupAddrs[isBeacon].length);
        require(gID2 < 0 || gID2 < groupAddrs[isBeacon].length);

        committee[isBeacon] = Groups({
            oldGID: gID1,
            newGID: gID2
        });
    }
}
