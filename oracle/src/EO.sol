// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "openzeppelin-contracts/contracts/token/ERC20/ERC20.sol";
import "openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Capped.sol";
import "openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Burnable.sol";
import {Pausable} from "openzeppelin-contracts/contracts/utils/Pausable.sol";
import {Ownable} from "openzeppelin-contracts/contracts/access/Ownable.sol";

contract EO is ERC20Capped, Pausable, Ownable {

    uint internal constant OWNER_TOKENS = 20000000;
    uint internal constant TOKEN_SIZE = 10 ** 18;
    string internal constant TOKEN_NAME = "EsportOracleToken";
    string internal constant TOKEN_TICKER = "EO";
    mapping(address => bool) public AuthorizedContract;

    //EVENTS
    event ContractAuthorized(address indexed contractAddress, bool authorized);
    event TokensMinted(address indexed to, uint256 amount);

    constructor(uint256 cap) ERC20(TOKEN_NAME,  TOKEN_TICKER) ERC20Capped(cap * TOKEN_SIZE) Ownable(msg.sender) {
        _mint(msg.sender, OWNER_TOKENS * (TOKEN_SIZE));
    }

    function pause() external onlyOwner {
        _pause();
    }

    function unpause() external onlyOwner {
        _unpause();
    }

    function setAuthorizedContract(address addressContract, bool authorized) external onlyOwner {
        AuthorizedContract[addressContract] = authorized;
        emit ContractAuthorized(addressContract, authorized);
    }

    function mintTokens(address _to, uint256 _amount) external onlyOwner whenNotPaused {
        _mint(_to, _amount);
        emit TokensMinted(_to, _amount);
    }

    function _update(address from, address to, uint256 value) internal override {
        require(!paused(), "Token transfer while paused");
        super._update(from, to, value);
    }
}
