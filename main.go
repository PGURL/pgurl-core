package main

import (
	"net/http"

	"github.com/PGURL/pgurl-core/pgurl/short"
	"github.com/gin-gonic/gin"
)

type url struct {
	URL string `json:"url"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/short", func(c *gin.Context) {
		var postURL url
		c.BindJSON(&postURL)
		shortURL, shortErr := short.ShortURL(postURL.URL)
		if shortErr != nil {
			c.JSON(200, gin.H{
				"status":   "url error",
				"shorturl": "",
			})

		} else {
			c.JSON(200, gin.H{
				"status":   "success",
				"shorturl": shortURL,
			})
		}
	})
	r.GET("/q/:url_key", func(c *gin.Context) {
		urlKey := c.Param("url_key")
		reurl, ReErr := short.ReShort(urlKey)
		if ReErr != nil {
			c.JSON(200, gin.H{
				"status": "failed",
				"msg":    ReErr,
			})
		}
		c.Redirect(http.StatusMovedPermanently, reurl)
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}
