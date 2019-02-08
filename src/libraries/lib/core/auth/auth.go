package auth

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Jwt(c *gin.Context) {
	var response = Response{}
	var errors Errors

	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("rahasia"), nil
	})

	if token == nil || err != nil {
		errors.Code = 11 //TODO : let's create repository errors code
		errors.Message = err.Error()

		response.Set("failed", "Authentication Failed", nil, errors)
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
	}
}
