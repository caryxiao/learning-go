package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name  string
	Age   int
	Grade string
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/leaning-go?charset=utf8mb4&parseTime=True&loc=UTC"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Student{})
	if err != nil {
		panic(err)
	}

	// 先清空数据
	db.Exec("TRUNCATE TABLE students")

	// add new record
	db.Model(&Student{}).Create(&Student{
		Name:  "张三",
		Age:   20,
		Grade: "三年级",
	})

	db.Model(&Student{}).Create(&Student{
		Name:  "李四",
		Age:   13,
		Grade: "一年级",
	})

	db.Model(&Student{}).Create(&Student{
		Name:  "王五",
		Age:   18,
		Grade: "二年级",
	})

	db.Model(&Student{}).Create(&Student{
		Name:  "旺财",
		Age:   21,
		Grade: "三年级",
	})

	// query all rows by age then 18
	var students []Student
	results := db.Where("age > ?", 18).Find(&students)
	if results.Error != nil {
		panic(results.Error.Error())
	}

	for _, row := range students {
		fmt.Println(fmt.Sprintf("%+v", row))
	}

	//update data
	student1 := Student{}
	db.Model(&Student{Name: "张三"}).First(&student1)
	student1.Grade = "四年级"
	db.Save(&student1)

	result := db.Model(&Student{
		Name: "张三",
	}).First(&student1)

	if result.Error != nil {
		panic(result.Error.Error())
	}

	fmt.Println(student1)

	// delete by age < 15
	db.Where("age < ?", 15).Delete(&Student{})
}
