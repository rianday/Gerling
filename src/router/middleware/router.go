package middleware

import (
	"webapi/controllers/auth"

	"github.com/gin-gonic/gin"
)

func init() {
	router := gin.Default()

	v1 := router.Group("/auth")
	{
		v1.POST("/login", auth.LoginAction)
		v1.POST("/register", auth.RegisterAction)
		v1.POST("/createTable", auth.RecreateTable)
	}

	v2 := router.Group("/admin")
	{
		v2.POST("login", auth.LoginAction)
	}

	router.Run()
}

func Run() {
}
