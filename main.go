package main

import (
	"github.com/gin-gonic/gin"
)

type shortenUri struct {
	Url string `form:"url"`
}

type redirectUri struct {
	Id string `uri:"id" binding:"required"`
}

var DB_FILE string = "./app.db"

func main() {

	db := initDatabase(DB_FILE)
	r := gin.Default()

	r.GET("/to/:id", func(c *gin.Context) {
		var uri redirectUri
		if err := c.ShouldBindUri(&uri); err != nil {
			c.JSON(400, gin.H{"message": err})
		}
		var RedirectLoc string
		row := db.QueryRow("SELECT RedirectLoc FROM urls WHERE id = ?", uri.Id)
		err := row.Scan(&RedirectLoc)
		if err != nil {
			c.JSON(404, gin.H{"message": err.Error()})
			return
		}
		c.Header("Location", RedirectLoc)
		c.JSON(301, gin.H{
			"message": "ok",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/shorten", func(c *gin.Context) {
		var uri shortenUri
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
