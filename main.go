package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/snir-maatuf/URL-shortener/handler"
	"github.com/snir-maatuf/URL-shortener/store"
)

func main() {
	// Set Gin to release mode
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the URL Shortener API",
		})
	})

	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	// Note that store initialization happens here
	store.InitializeStore()

	err := r.Run(":6379")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}

}
