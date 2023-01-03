// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/governance/GovernorUpgradeable.sol";
import "./GovernorCompatibilityBravoUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/governance/extensions/GovernorVotesUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/governance/extensions/GovernorVotesQuorumFractionUpgradeable.sol";

contract IncognitoDAO is GovernorUpgradeable, GovernorCompatibilityBravoUpgradeable, GovernorVotesUpgradeable, GovernorVotesQuorumFractionUpgradeable {

    bytes32 public constant PROPOSAL_HASH = keccak256("proposal");

    // storage layout
    address private tempAddress;

    modifier senderWithSig() {
        _;
        tempAddress = address(0x0);
    }

    function initialize(IVotesUpgradeable _token) external initializer {
        __Governor_init("IncognitoDAO");
        __GovernorVotes_init(_token);
        __GovernorVotesQuorumFraction_init(50);
    }

    function votingDelay() public pure override returns (uint256) {
        return 50400; // 7 days
    }

    function votingPeriod() public pure override returns (uint256) {
        return 7200; // 1 day
    }

    function proposalThreshold() public pure override returns (uint256) {
        return 62500 * 1e9; // 0.25% of 25M PRV
    }

    function quorumVotes() public pure override returns (uint256) {
        return 500000 * 1e9; // 2% of 25M PRV
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
    override(GovernorUpgradeable, IGovernorUpgradeable)
    returns (ProposalState)
    {
        return super.state(proposalId);
    }

    function propose(address[] memory targets, uint256[] memory values, bytes[] memory calldatas, string memory description)
    public
    override(GovernorUpgradeable, GovernorCompatibilityBravoUpgradeable)
    returns (uint256)
    {
        return super.propose(targets, values, calldatas, description);
    }

    function _execute(uint256 proposalId, address[] memory targets, uint256[] memory values, bytes[] memory calldatas, bytes32 descriptionHash)
    internal
    override(GovernorUpgradeable)
    {
        super._execute(proposalId, targets, values, calldatas, descriptionHash);
    }

    function _cancel(address[] memory targets, uint256[] memory values, bytes[] memory calldatas, bytes32 descriptionHash)
    internal
    override(GovernorUpgradeable)
    returns (uint256)
    {
        return super._cancel(targets, values, calldatas, descriptionHash);
    }

    function _executor()
    internal
    view
    override(GovernorUpgradeable)
    returns (address)
    {
        return super._executor();
    }

    function supportsInterface(bytes4 interfaceId)
    public
    view
    override(GovernorUpgradeable, IERC165Upgradeable)
    returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }

    function _msgSender() internal view override returns (address) {
        return tempAddress != address(0x0) ? tempAddress : msg.sender;
    }

    // pdapp handle with user signature
    function proposeBySig(
        address[] memory targets,
        uint256[] memory values,
        bytes[] memory calldatas,
        string memory description,
        uint8 v,
        bytes32 r,
        bytes32 s
    )
    public
    senderWithSig
    returns (uint256)
    {
        bytes32 signData = _hashTypedDataV4(keccak256(abi.encode(PROPOSAL_HASH, targets, values, calldatas, description)));
        tempAddress = ECDSAUpgradeable.recover(
            signData,
            v,
            r,
            s
        );
        return super.propose(targets, values, calldatas, description);
    }

    function cancelBySig(
        uint256 proposalId,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) public senderWithSig {
        tempAddress = ECDSAUpgradeable.recover(
            _hashTypedDataV4(keccak256(abi.encode(proposalId))),
            v,
            r,
            s
        );

        cancel(proposalId);
    }

    // override unused functions
    function proposalEta(uint256) public pure override returns (uint256) {
        return 0;
    }

    function timelock() public pure override returns (address) {
        return address(0x0);
    }

    function _quorumReached(uint256 proposalId)
    internal
    view
    virtual
    override(GovernorCompatibilityBravoUpgradeable, GovernorUpgradeable)
    returns (bool)
    {
        return super._forVotes(proposalId) >= quorumVotes();
    }

    function queue(
        address[] memory,
        uint256[] memory,
        bytes[] memory,
        bytes32
    ) public pure override returns (uint256) {
        return 0;
    }

    function getDataSign(bytes32 _input) external view returns (bytes32) {
        return _hashTypedDataV4(_input);
    }
}