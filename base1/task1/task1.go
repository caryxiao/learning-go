package main

import (
	"fmt"
	"strconv"
)

func main() {
	r1_1 := isPalindrome1(0)
	fmt.Println("isPalindrome1: 0 is", r1_1)
	r1_2 := isPalindrome1(10)
	fmt.Println("isPalindrome1: 10 is", r1_2)
	r1_3 := isPalindrome1(121)
	fmt.Println("isPalindrome1: 121 is", r1_3)
	r1_4 := isPalindrome1(-121)
	fmt.Println("isPalindrome1: -121 is", r1_4)
	r1_5 := isPalindrome1(2552)
	fmt.Println("isPalindrome1: 2552 is", r1_5)
	r1_6 := isPalindrome1(2345432)
	fmt.Println("isPalindrome1: 2345432 is", r1_6)
	r1_7 := isPalindrome1(1211)
	fmt.Println("isPalindrome1: 1211 is", r1_7)

	result1 := isPalindrome2(0)
	fmt.Println("isPalindrome2: 0 is", result1)
	result2 := isPalindrome2(10)
	fmt.Println("isPalindrome2: 10 is", result2)
	result3 := isPalindrome2(121)
	fmt.Println("isPalindrome2: 121 is", result3)
	result4 := isPalindrome2(-121)
	fmt.Println("isPalindrome2: -121 is", result4)
	result5 := isPalindrome2(2552)
	fmt.Println("isPalindrome2: 2552 is", result5)
	result6 := isPalindrome2(2345432)
	fmt.Println("isPalindrome2: 2345432 is", result6)
	result7 := isPalindrome2(1211)
	fmt.Println("isPalindrome2: 1211 is", result7)
}

func isPalindrome1(x int) bool {
	if x == 0 {
		return true
	}

	if x < 0 {
		return false
	}

	if x%10 == 0 {
		return false
	}

	var str = strconv.Itoa(x)
	var sr = []rune(str)

	for i, j := 0, len(sr)-1; i < j; i, j = i+1, j-1 {
		sr[i], sr[j] = sr[j], sr[i]
	}

	return str == string(sr)
}

func isPalindrome2(x int) bool {
	var str = strconv.Itoa(x)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}
