// SPDX-License-Identifier: MIT
pragma solidity ^0.8.10;
contract mergeSortedArray {
    function merge(uint[] memory arr1, uint[] memory arr2) public pure returns (uint[] memory) {
        uint len = arr1.length+arr2.length;
        uint[] memory newArr = new uint[](len);
        uint i=0;
        uint j=0;
        uint k=0;
        while(i < arr1.length && j < arr2.length) {
            if (arr1[i] <= arr2[j]) {
                newArr[k++] = arr1[i++];
            } else {
                newArr[k++] = arr2[j++];
            }
        }

        while (i < arr1.length) {
            newArr[k++] = arr1[i++];
        }

        while (j < arr2.length) {
            newArr[k++] = arr2[j++];
        }

        return newArr;
    }
}