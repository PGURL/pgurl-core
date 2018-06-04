package main

import (
	"./src/db"
	"./src/gen"
        "net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/short", func(c *gin.Context) {
		url := c.PostForm("url")
		hash := gen.GenRandomStr(6)
		status := db.Create(hash, url)

		c.JSON(200, gin.H{
			"status":   status,
			"shorturl": hash,
		})
	})
        r.GET("/q/:hash", func(c *gin.Context) {
                hash := c.Param("hash")
                url, status := db.Read(hash)
                if (status != "err") {
		    c.Redirect(http.StatusMovedPermanently, url)
                }
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
