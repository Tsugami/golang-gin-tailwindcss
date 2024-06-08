package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("assets/templates/*")
	r.Static("/assets", "./assets/dist")

	r.GET("/ping", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", func(c *gin.Context) {
		now := time.Now()
		currentTime := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
			now.Year(), now.Month(), now.Day(),
			now.Hour(), now.Minute(), now.Second())

		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":       "Main website",
			"currentTime": currentTime,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
