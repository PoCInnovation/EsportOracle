// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract EsportOracle {
    struct Personn {
        uint256 id;
        string name;
    }
    mapping(address => Personn) test;
    function setPerson(address _to, uint256 _id, string memory _name) public {
        test[_to].id = _id;
        test[_to].name = _name;
    }
    function getPerson(address _to) public view returns (uint256, string memory) {
        return (test[_to].id, test[_to].name);
    }
}
