package main

import (
	"database/sql"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// load env variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file - using local sqlite3 db instead")
	}

	var DB_URL string = os.Getenv("DB_URL")
	var db *sql.DB

	if len(DB_URL) > 0 {
		db = initDatabase(DB_URL, "postgres")
	} else {
		// use a local sqlite3 database instead (development)
		log.Println("[WARNING] - using local sqlite db instead. This is okay for development but NOT production")
		// refresh data in db if it exists
		if _, err := os.Stat("app.db"); err == nil {

			// app.db exists - remove it
			err := os.Remove("app.db")

			if err != nil {
				log.Fatal(err.Error())
			}
		}
		// init new one
		db = initDatabase("app.db", "sqlite3")

		// create table
		db.Exec("CREATE TABLE \"urls\" (\"id\"	TEXT, \"redirectloc\"	TEXT, PRIMARY KEY(\"id\"));")
	}

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.Use(static.Serve("/", static.LocalFile("./client/build", true)))

	// redirect routing route
	// id is the id that corresponds to the id in the database
	r.GET("/to/:id", func(c *gin.Context) {

		var uri redirectUri // see request_models.go for struct

		if err := c.ShouldBindUri(&uri); err != nil {
			c.JSON(400, gin.H{"message": err.Error()})
		}
		var RedirectLoc string

		row := db.QueryRow("SELECT redirectloc FROM urls WHERE id = $1", uri.Id)
		err := row.Scan(&RedirectLoc)

		// will error out if the id is not found in the database
		if err != nil {
			c.HTML(404, "404.html", nil)
			return
		}

		// check that redirectloc contains http at the begining
		// check that AT LEAST http is in url string...
		// if so... verify it is at the begining!
		//		if not --> place at begining
		//		else --> proceed as normal
		// if NOT in url string at all
		//		place at begining
		indx := strings.Index(RedirectLoc, "http")

		if indx != 0 { // http not in correct location or not present at all
			// place at begining
			// this case catches url strings
			// that have no http at all
			// and those that have it somewhere in middle
			RedirectLoc = "http://" + RedirectLoc

		}

		// set the "Location" header for the 301 redirect
		c.Header("Location", RedirectLoc)
		c.JSON(301, gin.H{
			"message": "ok",
		})
	})

	r.GET("/shorten", func(c *gin.Context) {

		var uri shortenUri // see request_models.go for struct
		var shortUri string

		if err := c.Bind(&uri); err != nil {
			c.JSON(400, gin.H{"msg": err})
		}
		shortUri = shortenUrl(uri.Url, db)
		c.JSON(200, gin.H{
			"url":      uri.Url,
			"shortUrl": shortUri,
		})
	})

	r.GET("/random", func(c *gin.Context) {
		var randID string = generateRandId(6)
		c.JSON(200, gin.H{
			"ID": randID,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
