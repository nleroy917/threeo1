package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"

	_ "github.com/lib/pq"
)

func initDatabase(file string) *sql.DB {
	db, err := sql.Open("postgres", file)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func shortenUrl(url string, db *sql.DB) string {

	var shortUriBase string = "https://threeo1.com/to/"
	const ID_LENGTH = 6

	id := generateRandId(ID_LENGTH)

	for idExists(id, db) {
		id = generateRandId(ID_LENGTH)
	}

	insertNewRedirect(url, id, db)

	return (shortUriBase + id)
}

func idExists(potentialId string, db *sql.DB) bool {

	var id string

	row := db.QueryRow("SELECT id FROM URLS WHERE id = $1", potentialId)
	err := row.Scan(&id)

	// will error out if the id is not found in the database
	// thus, return false if it can't find it
	if err != nil {
		return false
	}

	// if err returned, id was in db, return true
	return true
}

func insertNewRedirect(url string, id string, db *sql.DB) error {
	var err error = nil

	_, dberr := db.Exec("INSERT INTO URLS(id, RedirectLoc) VALUES ($1, $2)", id, url)
	if dberr != nil {
		fmt.Println("Error")
		log.Println(dberr.Error())
		return dberr
	}
	return err
}

func generateRandId(N int) string {
	char_map := [...]byte{'a', 'b', 'c', 'd',
		'e', 'f', 'g', 'h', 'i',
		'j', 'k', 'l', 'm', 'n',
		'o', 'p', 'q', 'r', 's',
		't', 'u', 'v', 'w', 'x',
		'y', 'z', '0', '1', '2',
		'3', '4', '5', '6', '7',
		'8', '9'}
	var ID string = ""
	for i := 0; i < N; i++ {
		indx := rand.Intn(len(char_map))
		pick := char_map[indx]
		ID += string(pick)
	}
	return ID
}
