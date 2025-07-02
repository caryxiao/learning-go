package main

import (
	"fmt"
	"time"
)

func addtTen(iVal *int) {
	*iVal += 10
}

func sliceValMul2(slice []int) {
	for i := 0; i < len(slice); i++ {
		slice[i] *= 2
	}
}

func printOddEven(mode string) {
	for i := 1; i < 10; i++ {
		if mode == "odd" && i%2 == 1 {
			fmt.Println(i, " is odd")
		} else if mode == "even" && i%2 == 0 {
			fmt.Println(i, " is even")
		}
	}
}

func main() {
	// 题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
	//考察点 ：指针的使用、值传递与引用传递的区别。
	var iVal int
	addtTen(&iVal)
	fmt.Println("iVal: original value =>0, new value => ", iVal)

	//题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
	//考察点 ：指针运算、切片操作。
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sliceValMul2(slice1)
	fmt.Println("slice value x 2: ", slice1)

	go printOddEven("odd")
	go printOddEven("even")

	time.Sleep(2 * time.Second)
}
