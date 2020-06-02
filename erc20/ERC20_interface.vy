
# External Contracts
contract Erc20:
    def totalSupply() -> uint256: constant
    def allowance(_owner: address, _spender: address) -> uint256: constant
    def transfer(_to: address, _value: uint256) -> bool: modifying
    def transferFrom(_from: address, _to: address, _value: uint256) -> bool: modifying
    def approve(_spender: address, _value: uint256) -> bool: modifying
    def name() -> string[64]: constant
    def symbol() -> string[32]: constant
    def decimals() -> uint256: constant
    def balanceOf(arg0: address) -> uint256: constant


