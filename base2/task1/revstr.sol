// SPDX-License-Identifier: MIT
pragma solidity ^0.8.10;
contract revstr {
    function reverstStr(string memory str) public pure returns(string memory) {
        bytes memory strBytes = bytes(str);
        for (uint i=0; i < strBytes.length / 2; i++) {
            bytes1 tmp = strBytes[i];
            strBytes[i] = strBytes[strBytes.length - i - 1];
            strBytes[strBytes.length - i - 1] = tmp;
        }
        return string(strBytes);
    }
}