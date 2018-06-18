package main

import (
	"github.com/PGURL/pgurl-core/pgurl/db"
	"github.com/PGURL/pgurl-core/pgurl/encrypt"
	"github.com/PGURL/pgurl-core/pgurl/gen"
	"github.com/PGURL/pgurl-core/pgurl/url"
	"github.com/gin-gonic/gin"
	"net/http"
)

const key = "LKHlhb899Y09olUi"

type URL struct {
	URL string `json:"url"`
}

func main() {
	r := gin.Default()
	aes_key := []byte(key)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/short", func(c *gin.Context) {
		var post_url URL
		c.BindJSON(&post_url)
		if url.IsValidUrl(post_url.URL) {
			// value, store in db.
			hash_url, _ := encrypt.Encrypt(aes_key, post_url.URL)

			// key, return to user.
			url_key := gen.GenRandomStr(6)
			status := db.Create(url_key, hash_url)

			c.JSON(200, gin.H{
				"status":   status,
				"shorturl": url_key,
			})
		} else {
			c.JSON(200, gin.H{
				"status":   "url error",
				"shorturl": "",
			})

		}
	})
	r.GET("/q/:url_key", func(c *gin.Context) {
		url_key := c.Param("url_key")
		hash_url, status := db.Read(url_key)
		if status != "err" {
			url, _ := encrypt.Decrypt(aes_key, hash_url)
			c.Redirect(http.StatusMovedPermanently, url)
		}
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
