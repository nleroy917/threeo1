package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func initDatabase(file string) *sql.DB {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func shorten_url(url string) (id int) {
	fmt.Println("Shortening:", url)
	return
}
