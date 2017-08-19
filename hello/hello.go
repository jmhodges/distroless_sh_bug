package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalln("Open:", err)
	}
	_, err = db.Exec("create table foo (name text)")
	if err != nil {
		log.Fatalln("Exec:", err)
	}
	defer db.Close()
	fmt.Println("Hello, World")
}
