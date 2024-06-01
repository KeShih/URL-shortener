package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"URL-shortener/handler"
	"URL-shortener/store"
)

func main()  {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, URL Shortener!",
		})
	})

	r.POST("/create-short-url", handler.CreateShortUrl)
	r.GET("/:shortUrl", handler.RedirectShortUrl)

	store.InitStore()

	err := r.Run(":8080")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}