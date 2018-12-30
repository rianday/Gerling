package auth

import (
	"libraries/models/sys"
)

type FormLogin struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (login *FormLogin) CheckLogin() (error, interface{}) {
	var response = Response{}
	var errors Errors

	var user sys.User
	user.Email = login.Email
	user.Password = login.Password

	err, result := user.Get()

	if err != nil {
		errors.Code = 11 //TODO : let's create repository errors code
		errors.Message = err.Error()

		response.Set("failed", "Login failed", nil, errors)
		return err, nil
	} else {
		response.Set("success", "You are logged in", result, errors)
	}

	return nil, response
}
