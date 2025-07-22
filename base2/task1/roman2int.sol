// SPDX-License-Identifier: MIT
pragma solidity ^0.8.10;
contract roman2int {
    mapping(bytes1 => uint) public r2iMap;
    constructor() {
        r2iMap["I"] = 1;
        r2iMap["V"] = 5;
        r2iMap["X"] = 10;
        r2iMap["L"] = 50;
        r2iMap["C"] = 100;
        r2iMap["D"] = 500;
        r2iMap["M"] = 1000;
    }

    function romanToInt(string memory romanStr) public view returns(uint) {
        bytes memory romanStrArr = bytes(romanStr);
        uint val = 0;
        for (uint i=0; i<romanStrArr.length; i++) {
            bytes1 nextb = i == romanStrArr.length - 1 ? bytes1("") : romanStrArr[i + 1];
            if (romanStrArr[i] == 'I' && nextb != bytes1("") && (nextb == bytes1("V") || nextb == bytes1("X"))) {
                val += nextb == bytes1("V") ? 4 : 9;
                i++;
                continue;
            }

            if (romanStrArr[i] == 'X' && nextb != bytes1("") && (nextb == bytes1("L") || nextb == bytes1("C"))) {
                val += nextb == bytes1("L") ? 40 : 90;
                i++;
                continue;
            }

            if (romanStrArr[i] == 'C' && nextb != bytes1("") && (nextb == bytes1("D") || nextb == bytes1("M"))) {
                val += nextb == bytes1("D") ? 400 : 900;
                i++;
                continue;
            }

            val += r2iMap[romanStrArr[i]];
        }
        return val;
    }
}