package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var DB_FILE string = "./app.db"

func main() {

	db := initDatabase(DB_FILE)
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// redirect routing route
	// id is the id that corresponds to the id in the database
	r.GET("/to/:id", func(c *gin.Context) {

		var uri redirectUri // see request_models.go for struct

		if err := c.ShouldBindUri(&uri); err != nil {
			c.JSON(400, gin.H{"message": err})
		}
		var RedirectLoc string
		row := db.QueryRow("SELECT RedirectLoc FROM urls WHERE id = ?", uri.Id)
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

		if err := c.Bind(&uri); err != nil {
			c.JSON(400, gin.H{"msg": err})
		}
		shorten_url(uri.Url)
		c.JSON(200, gin.H{
			"url": uri.Url,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
