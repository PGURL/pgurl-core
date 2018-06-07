package main

import (
	"github.com/PGURL/pgurl-core/pgurl/db"
	"github.com/PGURL/pgurl-core/pgurl/encrypt"
	"github.com/PGURL/pgurl-core/pgurl/gen"
	"github.com/gin-gonic/gin"
	"net/http"
)

const key = "LKHlhb899Y09olUi"

func main() {
	r := gin.Default()
	aes_key := []byte(key)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/short", func(c *gin.Context) {
		url := c.PostForm("url")
		// value, store in db.
		hash_url, _ := encrypt.Encrypt(aes_key, url)

		// key, return to user.
		url_key := gen.GenRandomStr(6)
		status := db.Create(url_key, hash_url)

		c.JSON(200, gin.H{
			"status":   status,
			"shorturl": url_key,
		})
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
