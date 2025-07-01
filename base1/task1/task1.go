package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	r1 := isPalindrome1(0)
	fmt.Println("isPalindrome1: 0 is", r1)
	r2 := isPalindrome1(10)
	fmt.Println("isPalindrome1: 10 is", r2)
	r3 := isPalindrome1(121)
	fmt.Println("isPalindrome1: 121 is", r3)
	r4 := isPalindrome1(-121)
	fmt.Println("isPalindrome1: -121 is", r4)
	r5 := isPalindrome1(2552)
	fmt.Println("isPalindrome1: 2552 is", r5)
	r6 := isPalindrome1(2345432)
	fmt.Println("isPalindrome1: 2345432 is", r6)
	r7 := isPalindrome1(1211)
	fmt.Println("isPalindrome1: 1211 is", r7)

	s1 := "[()]"
	s2 := "([})"
	s3 := "()[]{}"
	fmt.Println("isValid: ", s1, " is", isValid(s1))
	fmt.Println("isValid: ", s2, " is", isValid(s2))
	fmt.Println("isValid: ", s3, " is", isValid(s3))

	lstr1 := []string{"flower", "flow", "flight"}
	fmt.Println("longestCommonPrefix: ", lstr1, longestCommonPrefix(lstr1))

	plusOne1 := []int{1, 2, 3}
	plusOne2 := []int{1, 2, 9}
	plusOne3 := []int{9, 9, 9}
	fmt.Println("plusOne: ", plusOne1, plusOne(plusOne1))
	fmt.Println("plusOne: ", plusOne2, plusOne(plusOne2))
	fmt.Println("plusOne: ", plusOne3, plusOne(plusOne3))

	removeDuplicates1 := []int{
		0, 0, 1, 1, 1, 2, 2, 3, 3, 4,
	}
	removeDuplicates1Copy := append([]int{}, removeDuplicates1...)

	fmt.Println("removeDuplicates: ", removeDuplicates1Copy, " new length:", removeDuplicates(removeDuplicates1))

	m1 := [][]int{
		{2, 6},
		{8, 10},
		{15, 18},
		{1, 3},
	}

	fmt.Println("merge: ", m1, " new: ", merge(m1))

	t1 := []int{2, 11, 7, 15}
	t1Target := 9
	fmt.Println("twoSum: ", t1, " target: ", t1Target, " solution: ", twoSum(t1, t1Target))
	fmt.Println("twoSum2: ", t1, " target: ", t1Target, " solution: ", twoSum2(t1, t1Target))

}

/*
** 136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
** 找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
** 例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
 */
func isPalindrome1(x int) bool {
	if x == 0 {
		return true
	}

	// 小于0和大于0的时候且以0结尾的数都不是回文数
	if x < 0 || x%10 == 0 {
		return false
	}

	// int转换为字符串
	var str = strconv.Itoa(x)
	// 创建rune切片
	var sr = []rune(str)

	// 翻转切片元素
	for i, j := 0, len(sr)-1; i < j; i, j = i+1, j-1 {
		sr[i], sr[j] = sr[j], sr[i]
	}

	return str == string(sr)
}

/*
** 20. 有效的括号：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
** 有效字符串需满足：左括号必须用相同类型的右括号闭合。左括号必须以正确的顺序闭合。每个右括号都有一个对应的相同类型的左括号。
 */
func isValid(s string) bool {
	// 奇数个字符串肯定不会匹配,只有偶数个字符串才能成对出现
	if len(s)%2 != 0 {
		return false
	}

	strMap := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	var stack []byte

	for i := 0; i < len(s); i++ {
		// 如果遇到左括号则入栈
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else if len(stack) > 0 && stack[len(stack)-1] == strMap[s[i]] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}

/*
** 最长公共前缀
** 考察：字符串处理、循环嵌套
** 题目：查找字符串数组中的最长公共前缀
 */
func longestCommonPrefix(strs []string) string {
	var prefix []byte
	if len(strs) == 1 {
		return strs[0]
	}
outer:
	for i := 0; i >= 0; i++ {
		for j := 0; j < len(strs); j++ {
			// 超出长度跳出
			if i > len(strs[j])-1 {
				break outer
			} else if j > 0 && strs[j][i] != strs[j-1][i] { //跟之前的只要有一个不一样，则跳出
				break outer
			} else if j != 0 && j == len(strs)-1 { // 只有最后一个字符串跟前面的都匹配才能算是公共前缀
				prefix = append(prefix, strs[j][i])
			}
		}
	}
	return string(prefix)
}

/*
** 加一
** 难度：简单
** 考察：数组操作、进位处理
** 题目：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
 */
func plusOne(digits []int) []int {
	var slack []int
	var add int
	for i := len(digits) - 1; i >= 0; i-- {
		var cur []int
		if i == len(digits)-1 || add > 0 {
			if digits[i] == 9 {
				cur = append(cur, 0)
				add = 1
				if i == 0 {
					cur = append([]int{1, 0})
					add = 0
				}
			} else {
				cur = append(cur, digits[i]+1)
				add = 0
			}
			slack = append(cur, slack...)
		} else {
			cur = append(cur, digits[i])
			slack = append(cur, slack...)
		}
	}
	return slack
}

/*
*
** 26. 删除有序数组中的重复项：给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，
** 返回删除后数组的新长度。不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
** 可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。
 */
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for j := 1; j < len(nums); j++ {
		//fmt.Println(nums[i], nums[j])
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}

	//fmt.Println(nums)
	return i + 1 // 长度为索引+1
}

/*
** 合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
** 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
** 可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，
** 遍历排序后的区间数组，将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
 */
func merge(intervals [][]int) [][]int {
	// 先排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	newIntervals := append([][]int{}, intervals[0])
	for i := 1; i < len(intervals); i++ {
		last := newIntervals[len(newIntervals)-1:]
		if last[0][1] < intervals[i][0] || intervals[i][1] < last[0][0] {
			newIntervals = append(newIntervals, intervals[i])
		} else {
			// 合并区间
			last[0][0] = min(last[0][0], intervals[i][0])
			last[0][1] = max(last[0][1], intervals[i][1])
		}
	}
	return newIntervals
}

/*
** 两数之和
** 考察：数组遍历、map使用
** 题目：给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
 */
func twoSum(nums []int, target int) []int {
	// 双循环，比较耗费时间
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSum2(nums []int, target int) []int {
	hash := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if val, ok := hash[target-nums[i]]; ok {
			return []int{val, i}
		} else {
			hash[nums[i]] = i
		}
	}
	return []int{}
}
