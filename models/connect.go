package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var conn *sql.DB

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/todo_go")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database.")

	conn = db

	return db
}
