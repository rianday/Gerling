package main

import (
	"libraries/lib/core/auth"

	"github.com/gin-gonic/gin"
)

func main() {
	auth.Register()
	ping()
	//test.RunCrudExample()
}

func ping() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
