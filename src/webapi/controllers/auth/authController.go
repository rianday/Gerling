package auth

import (
	"libraries/models/sys"
	"net/http"

	"libraries/lib/core/auth"

	"github.com/gin-gonic/gin"
)

func RegisterAction(c *gin.Context) {
	var form auth.FormRegister
	var response auth.Response
	var errors auth.Errors

	if err := c.ShouldBind(&form); err != nil {
		errors.Code = 12
		errors.Message = err.Error()
		response.Set("failed", "Please check your input", nil, errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err, result := form.SetRegister()
	if err != nil {
		errors.Code = 12
		errors.Message = err.Error()
		response.Set("failed", "Please check your input", nil, errors)
		//c.JSON(http.StatusBadRequest, response)
	} else {
		if result == nil {
			errors.Code = 12
			errors.Message = "No Response System"
			response.Set("failed", "No Response", nil, errors)
		} else {
			response = result.(auth.Response)
		}
	}

	c.JSON(http.StatusOK, response)
}

func LoginAction(c *gin.Context) {
	var form auth.FormLogin
	var response auth.Response
	var errors auth.Errors

	if err := c.ShouldBind(&form); err != nil {
		errors.Code = 12
		errors.Message = err.Error()
		response.Set("failed", "Please check your input", nil, errors)
		c.JSON(http.StatusBadRequest, response)
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err, result := form.CheckLogin()
	if err != nil {
		if result == nil {
			errors.Code = 12
			errors.Message = err.Error()
			response.Set("failed", "Please check your input", nil, errors)
		} else {
			response = result.(auth.Response)
			//fmt.Println(result)
		}
	}

	c.JSON(http.StatusOK, result)
}

func RecreateTable(c *gin.Context) {
	var user sys.User

	user.CreateTable()
}

func AdminAction(c *gin.Context) {
	var response auth.Response
	var errors auth.Errors

	response.Set("success", "You are logged in", nil, errors)

	c.JSON(http.StatusOK, response)
}
