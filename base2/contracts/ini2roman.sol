// SPDX-License-Identifier: MIT
pragma solidity ^0.8.10;

contract ini2roman {
    function num2Roman(uint num) public pure returns(string memory) {
        if (num == 0) return "nulla";
        if (num > 3999) return "too big";
        string memory result = "";
        uint[13] memory values = [uint(1000), 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1];
        string[13] memory symbols = ["M", "CM","D","CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"];

        for (uint i = 0; i < 13; i++) {
            while (num >= values[i]) {
                result = string.concat(result, symbols[i]);
                num -= values[i];
            }
        }
        return result;
    }
}