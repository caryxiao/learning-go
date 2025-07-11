//SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract SimpleStorage {
    address private owner;
    uint256 private number;
    bool private writeable = false;

    event NumberStored(address indexed sender, uint256 newValue);
    event OwnerSet(address indexed oldOwner, address indexed newOwner);

    modifier isOwner() {
        require(msg.sender == owner, "Caller is not owner");
        _;
    }

    modifier isWriteable() {
        require(writeable == true, "Contract is not writeable");
        _;
    }

    constructor() {
        owner = msg.sender;
        emit OwnerSet(address(0), owner);
    }

    function store(uint256 num) public isWriteable {
        number = num;
        emit NumberStored(msg.sender, num);
    }

    function retrieve() public view returns(uint256) {
        return number;
    }

    function setWriteable() public isOwner {
        writeable = true;
    }

    function setReadonly() public isOwner {
        writeable = false;
    }

    function isContractWriteable() public view returns(bool) {
        return writeable;
    }

    function getOwner() public view returns(address) {
        return owner;
    }
}