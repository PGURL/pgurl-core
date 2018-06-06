package main

import (
	"./pgurl/api"
	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()
	r.GET("/ping", api.Ping)
	r.POST("/short", api.GenShort)
	r.GET("/q/:url_key", api.Redirection)
	r.Run() // listen and serve on 0.0.0.0:8080
}
