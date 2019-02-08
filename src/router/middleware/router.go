package middleware

import (
	authAct "webapi/controllers/auth"

	auth "libraries/lib/core/auth"

	"github.com/gin-gonic/gin"
)

func init() {
	router := gin.Default()

	v1 := router.Group("/auth")
	{
		v1.POST("/login", authAct.LoginAction)
		v1.POST("/register", authAct.RegisterAction)
		v1.POST("/createTable", authAct.RecreateTable)
	}

	v2 := router.Group("/admin")
	{
		v2.POST("login", authAct.LoginAction)
		v2.POST("home", auth.Jwt, authAct.AdminAction)
	}

	router.Run()
}

func Run() {
}
