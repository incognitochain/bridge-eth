pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/token/ERC20/extensions/ERC20VotesCompUpgradeable.sol";

contract PrvVoting is ERC20VotesCompUpgradeable {

    function initialize(string memory name_, string memory symbol_) external initializer {
        __ERC20VotesComp_init();
        __ERC20_init(name_, symbol_);
    }

}
