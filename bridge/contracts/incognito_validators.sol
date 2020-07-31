pragma solidity ^0.6.6;
pragma experimental ABIEncoderV2;

import "./operable.sol";

contract IncognitoValidators is Operable {

    // isBeacon => address[group 0], address[group 1], ...
    mapping(bool => address[][]) public groupAddrs;

    struct Groups {
        int oldGID;
        int newGID;
    }

    // isBeacon => swapID => (oldGID, newGID)
    mapping(bool => mapping(uint => Groups)) public validators;

    // isBeacon => (oldGID, newGID)
    mapping(bool => Groups) public committees; // Latest committee

    constructor(
        address _admin,
        address _operator,
        address[] memory beaconCommittee,
        address[] memory bridgeCommittee
    ) public Operable(_admin, _operator) {
    }

    function loadGroups(bool isBeacon, int oldGID, int newGID) public view returns(address[] memory) {
        address[] memory addrs;
        if (oldGID >= 0) {
            require(oldGID < groupAddrs[isBeacon].length);
            for (i := 0; i < groupAddrs[isBeacon][oldGID].length; i++) {
                addrs.push(groupAddrs[isBeacon][oldGID][i]);
            }
        }
        if (newGID >= 0) {
            require(newGID < groupAddrs[isBeacon].length);
            for (i := 0; i < groupAddrs[isBeacon][newGID].length; i++) {
                addrs.push(groupAddrs[isBeacon][newGID][i]);
            }
        }
        return addrs
    }

    function loadCandidates(bool isBeacon, uint swapID) public view returns(address[] memory) {
        return loadGroups(validators[isBeacon][swapID].oldGID, validators[isBeacon][swapID].newGID);
    }

    function storeCandidates(bool isBeacon, uint swapID, int oldGID, int newGID, address[] newAddrs) public onlyOperator {
        require(oldGID < 0 || oldGID < groupAddrs[isBeacon].length);
        require(newGID < 0 || newGID == groupAddrs[isBeacon].length);

        // Save new group
        if (newGID >= 0) {
            groupAddrs[isBeacon].push(newAddrs);
        }

        validators[isBeacon][swapID] = Groups{
            oldGID: oldGID,
            newGID: newGID
        };
    }

    function loadCommittee(bool isBeacon) public view returns(address[] memory) onlyOperator {
        return loadGroups(isBeacon, committee[isBeacon].oldGID, committee[isBeacon].newGID);
    }

    function storeCommittee(bool isBeacon, int gID1, int gID2) public onlyOperator {
        require(gID1 < 0 || gID1 < groupAddrs[isBeacon].length);
        require(gID2 < 0 || gID2 < groupAddrs[isBeacon].length);

        committee[isBeacon] = Groups{
            oldGID: gID1,
            newGID: gID2
        };
    }
}
