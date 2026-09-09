package controllers

import (
	services "AuthInGoService/services"
	"fmt"
	"net/http"
)

type UserController struct {
	userservice services.UserService
}

func NewUserController(_userservice services.UserService)* UserController{
	return &UserController{
		userservice: _userservice,
	}
}
func (u *UserController) GetUserInfo( w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting user info")
	u.userservice.GetUserById()
	fmt.Fprintf(w, "User info retrieved successfully")
}