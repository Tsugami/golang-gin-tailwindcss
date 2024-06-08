package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func getCurrentTime() string {
	now := time.Now()
	currentTime := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())

	return currentTime
}

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

		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":       "Main website",
			"currentTime": getCurrentTime(),
		})
	})

	r.GET("/current-time", func(c *gin.Context) {
		c.HTML(http.StatusOK, "current-time.html", gin.H{
			"currentTime": getCurrentTime(),
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
