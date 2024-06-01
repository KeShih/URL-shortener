package handler

import (
	"github.com/gin-gonic/gin"
	"URL-shortener/shortener"
	"URL-shortener/store"
	"net/http"
)

// Request model definition
type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortUrl(creationRequest.LongUrl, creationRequest.UserId)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl)

	host := "http://localhost:8080/"
	c.JSON(200, gin.H{
		"message"  : "Short URL created successfully",
		"short_url": host + shortUrl,
	})
}

func RedirectShortUrl(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	longUrl := store.GetLongUrl(shortUrl)
	c.Redirect(302, longUrl)
}



