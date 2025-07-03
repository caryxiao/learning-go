package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type Book struct {
	Id     uint64  `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
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

	var books []Book
	err = db.Select(&books, "SELECT * FROM books WHERE price > 50")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(books)
}
