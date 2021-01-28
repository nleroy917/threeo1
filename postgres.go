package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// init the database - returns sql.DB struct pointer
func initDatabase(file string, engine string) *sql.DB {
	db, err := sql.Open(engine, file)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// main function to shorten url's ----
// first, the function checks if the url
// exists in the database already...
// 	if it does, it simply returns the short url
// 	associated with that url.
//	else it will create a new entry --
//
// To create a new entry, genreate a potential id,
// while this id exists in the database, generate a
// a new one. Once a unique id is found, then create
// a new entry in the database
func shortenUrl(url string, db *sql.DB) string {

	var shortUriBase string = "https://www.threeo.one/to/"
	const ID_LENGTH = 6

	// check if url already exists in db
	// if it does return the short url for it
	// shorturl == "" if url doesnt exist
	shorturl := urlExists(url, db)
	if len(shorturl) > 0 {
		return shorturl
	}

	// genreate potential id
	id := generateRandId(ID_LENGTH)

	// while id exists in the database,
	// generate a new one
	for idExists(id, db) {
		id = generateRandId(ID_LENGTH)
	}

	// insert the new entry
	insertNewRedirect(url, id, db)

	return (shortUriBase + id)
}

// check if a url exists - if it does, return short url,
// else return an empty string
func urlExists(potentialUrl string, db *sql.DB) string {

	var id string
	var url string
	var shortUriBase string = "https://www.threeo.one/to/"

	// log.Println(potentialUrl)
	row := db.QueryRow("SELECT id, redirectloc FROM URLS WHERE redirectloc= $1", potentialUrl)
	err := row.Scan(&id, &url)

	// will error out if the url is not found in the database
	// thus, return empty string
	if err != nil {
		// log.Println("Not found")
		return url
	}

	// else ...
	// return the short id genreated from
	// base + id
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

	_, dberr := db.Exec("INSERT INTO URLS(id, redirectloc) VALUES ($1, $2)", id, url)
	if dberr != nil {
		fmt.Println("Error")
		log.Println(dberr.Error())
		return dberr
	}
	return err
}

// function to create random alphanumeric
// id's of length N.
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
