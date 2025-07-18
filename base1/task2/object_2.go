package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (e *Employee) PrintInfo() {
	fmt.Println("Employee Information => Employee ID:", e.EmployeeID, ",Name:", e.Name, ",Age:", e.Age)
}

func main() {
	//题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
	//考察点 ：组合的使用、方法接收者。
	employee1 := &Employee{
		Person: Person{
			Name: "Tom",
			Age:  20,
		},
		EmployeeID: 1000001,
	}

	employee1.PrintInfo()
}
