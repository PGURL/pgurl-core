package api

import (
	"github.com/gin-gonic/gin"
        "../encrypt"
        "../db"
        "../gen"
)

func GenShort(c *gin.Context) {
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
}
