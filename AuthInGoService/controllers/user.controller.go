package controllers

import (
	services "AuthInGoService/services"
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
func (u *UserController) Register( w http.ResponseWriter, r *http.Request) {
	    w.WriteHeader(http.StatusCreated)
}