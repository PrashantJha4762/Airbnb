package controllers

import (
	services "AuthInGoService/services"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserController struct {
	userservice services.UserService
}

func NewUserController(_userservice services.UserService) *UserController {
	return &UserController{
		userservice: _userservice,
	}
}
func (u *UserController) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting user info")
	u.userservice.GetUserById()
	fmt.Fprintf(w, "User info retrieved successfully")
}
func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := u.userservice.CreateUser(request.Username, request.Email, request.Password); err != nil {
		http.Error(w, "unable to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "User created successfully")
}
