// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract EsportOracle {
    address public _owner;

    constructor() {
        _owner = msg.sender;
    }

    modifier onlyOwner() {
        require(msg.sender == _owner, "Not the contract owner");
        _;
    }

    function setOwner(address newOwner) public onlyOwner {
        require(newOwner != address(0), "New owner cannot be zero address");
        _owner = newOwner;
    }
}
