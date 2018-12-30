package auth

import "libraries/models/sys"

type FormRegister struct {
	Username string `form:"username" json:"username" binding:"required"`
	Email    string `form:"email" json:"user" binding:"required"`
	Phone    string `form:"phone" json:"phone" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (reg *FormRegister) SetRegister() (error, interface{}) {
	var response = Response{}
	var errors Errors
	var user sys.User

	user.Email = reg.Email
	user.Password = reg.Password
	user.Username = reg.Username
	user.Phone = reg.Phone

	err, result := user.Add()

	if err != nil {
		errors.Code = 11 //TODO : let's create repository errors code
		errors.Message = err.Error()

		response.Set("failed", "Register failed", nil, errors)
		return err, nil
	} else {
		response.Set("success", "You have registered", result, errors)
	}

	return nil, response
}
