package auth

import "libraries/models/sys"
import "golang.org/x/crypto/bcrypt"

type FormRegister struct {
	Username string `form:"username" json:"username" binding:"required"`
	Email    string `form:"email" json:"user" binding:"required,email"`
	Phone    string `form:"phone" json:"phone" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,min=6"`
}

func (reg *FormRegister) SetRegister() (error, interface{}) {
	var response = Response{}
	var errors Errors
	var user sys.User
	var tempPassword []byte

	user.Email = reg.Email
	tempPassword, err := bcrypt.GenerateFromPassword([]byte(reg.Password), 14)
	user.Password = string(tempPassword)
	user.Username = reg.Username
	user.Phone = reg.Phone

	err, result := user.Add()

	if err != nil {
		errors.Code = 11 //TODO : let's create repository errors code
		errors.Message = err.Error()

		response.Set("failed", "Register failed", nil, errors)
	} else {
		response.Set("success", "You have registered", result, errors)
	}

	return err, response
}
