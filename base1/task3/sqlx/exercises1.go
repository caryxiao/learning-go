package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type Employee struct {
	Id         uint64  `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/leaning-go?charset=utf8mb4&parseTime=True&loc=UTC"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	var employees []Employee
	err = db.Select(&employees, "select * from employees where department = ?", "技术部")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v\n", employees)

	var maxSalaryEmployee Employee
	err = db.Get(&maxSalaryEmployee, "SELECT * FROM employees ORDER BY salary DESC LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v\n", maxSalaryEmployee)
}
