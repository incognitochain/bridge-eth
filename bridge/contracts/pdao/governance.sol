// SPDX-License-Identifier: MIT
pragma solidity ^0.8.2;

import "@openzeppelin/contracts-upgradeable/governance/GovernorUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/governance/compatibility/GovernorCompatibilityBravoUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/governance/extensions/GovernorVotesUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/governance/extensions/GovernorVotesQuorumFractionUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/governance/extensions/GovernorTimelockControlUpgradeable.sol";

contract IncognitoDAO is GovernorUpgradeable, GovernorCompatibilityBravoUpgradeable, GovernorVotesUpgradeable, GovernorVotesQuorumFractionUpgradeable, GovernorTimelockControlUpgradeable {

    // storage layout
    address private tempAddress;
    mapping(bytes32 => bool) public sigDataUsed;

    modifier senderWithSig() {
        _;
        tempAddress = address(0x0);
    }

    function initialize(IVotesUpgradeable _token, TimelockControllerUpgradeable _timelock) external initializer {
        __Governor_init("IncognitoDAO");
        __GovernorVotes_init(_token);
        __GovernorVotesQuorumFraction_init(100);
        __GovernorTimelockControl_init(_timelock);
    }

    function votingDelay() public pure override returns (uint256) {
        return 6575; // 1 day
    }

    function votingPeriod() public pure override returns (uint256) {
        return 46027; // 1 week
    }

    function proposalThreshold() public pure override returns (uint256) {
        return 0;
    }

    // The functions below are overrides required by Solidity.

    function quorum(uint256 blockNumber)
    public
    view
    override(IGovernorUpgradeable, GovernorVotesQuorumFractionUpgradeable)
    returns (uint256)
    {
        return super.quorum(blockNumber);
    }

    function state(uint256 proposalId)
    public
    view
    override(GovernorUpgradeable, IGovernorUpgradeable, GovernorTimelockControlUpgradeable)
    returns (ProposalState)
    {
        return super.state(proposalId);
    }

    function propose(address[] memory targets, uint256[] memory values, bytes[] memory calldatas, string memory description)
    public
    override(GovernorUpgradeable, GovernorCompatibilityBravoUpgradeable, IGovernorUpgradeable)
    returns (uint256)
    {
        return super.propose(targets, values, calldatas, description);
    }

    function _execute(uint256 proposalId, address[] memory targets, uint256[] memory values, bytes[] memory calldatas, bytes32 descriptionHash)
    internal
    override(GovernorUpgradeable, GovernorTimelockControlUpgradeable)
    {
        super._execute(proposalId, targets, values, calldatas, descriptionHash);
    }

    function _cancel(address[] memory targets, uint256[] memory values, bytes[] memory calldatas, bytes32 descriptionHash)
    internal
    override(GovernorUpgradeable, GovernorTimelockControlUpgradeable)
    returns (uint256)
    {
        return super._cancel(targets, values, calldatas, descriptionHash);
    }

    function _executor()
    internal
    view
    override(GovernorUpgradeable, GovernorTimelockControlUpgradeable)
    returns (address)
    {
        return super._executor();
    }

    function supportsInterface(bytes4 interfaceId)
    public
    view
    override(GovernorUpgradeable, IERC165Upgradeable, GovernorTimelockControlUpgradeable)
    returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }

    function _msgSender() internal view override returns (address) {
        return tempAddress != address(0x0) ? tempAddress : msg.sender;
    }

    //    function castVoteBySig(
    //        uint256 proposalId,
    //        uint8 support,
    //        uint8 v,
    //        bytes32 r,
    //        bytes32 s
    //    ) public virtual override returns (uint256) {
    //        address voter = ECDSAUpgradeable.recover(
    //            _hashTypedDataV4(keccak256(abi.encode(BALLOT_TYPEHASH, proposalId, support))),
    //            v,
    //            r,
    //            s
    //        );
    //        return _castVote(proposalId, voter, support, "");
    //    }

    /// @dev: pdapp handle with user signature

    // todo: add function propose by sig

    // todo: add function cancel by sig

}