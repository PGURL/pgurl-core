package api

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "../db"
    "../encrypt"
)

func Redirection (c *gin.Context) {
		url_key := c.Param("url_key")
		hash_url, status := db.Read(url_key)
		if status != "err" {
			url, _ := encrypt.Decrypt(aes_key, hash_url)
			c.Redirect(http.StatusMovedPermanently, url)
		}
	}

