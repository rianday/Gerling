package auth

import (
	"libraries/models/sys"

	"golang.org/x/crypto/bcrypt"
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

	err, result := user.Get()

	if err != nil {
		errors.Code = 11 //TODO : let's create repository errors code
		errors.Message = err.Error()

		response.Set("failed", "Login failed", nil, errors)
	} else {
		//fmt.Println(result.(*sys.User).Password)
		err := bcrypt.CompareHashAndPassword([]byte(result.(*sys.User).Password), []byte(login.Password))
		if err != nil {
			errors.Code = 11 //TODO : let's create repository errors code
			errors.Message = err.Error()

			response.Set("failed", "Your email and password doesn't match", nil, errors)
		} else {
			response.Set("success", "You are logged in", result, errors)
		}
	}

	return err, response
}
