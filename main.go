package main

import (
	"log"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var DB_URL string = os.Getenv("DB_URL")

	db := initDatabase(DB_URL)
	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("./client/build", true)))

	// redirect routing route
	// id is the id that corresponds to the id in the database
	r.GET("/to/:id", func(c *gin.Context) {

		var uri redirectUri // see request_models.go for struct

		if err := c.ShouldBindUri(&uri); err != nil {
			c.JSON(400, gin.H{"message": err})
		}
		var RedirectLoc string
		row := db.QueryRow("SELECT RedirectLoc FROM urls WHERE id = $1", uri.Id)
		err := row.Scan(&RedirectLoc)

		// will error out if the id is not found in the database
		if err != nil {
			c.JSON(404, gin.H{"message": err.Error()})
			return
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
