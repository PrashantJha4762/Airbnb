package controllers

import (
	"AuthInGoService/dto"
	services "AuthInGoService/services"
	"AuthInGoService/utils"
	"errors"
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

	if err := utils.ReadJsonResponse(r, &request); err != nil {
		_ = utils.WriteJsonResponse(w, http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
		return
	}

	if err := u.userservice.CreateUser(request.Username, request.Email, request.Password); err != nil {
		_ = utils.WriteJsonResponse(w, http.StatusInternalServerError, map[string]string{
			"error": "unable to create user",
		})
		return
	}

	_ = utils.WriteJsonResponse(w, http.StatusCreated, map[string]string{
		"message": "User created successfully",
	})
}
func (u *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	var request dto.LoginRequest
	if err := utils.ReadJsonResponse(r, &request); err != nil {
		_ = utils.WriteJsonResponse(w, http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
		return
	}

	if request.Email == "" || request.Password == "" {
		_ = utils.WriteJsonResponse(w, http.StatusBadRequest, map[string]string{
			"error": "email and password are required",
		})
		return
	}

	token, err := u.userservice.LoginUser(request.Email, request.Password)
	if err != nil {
		if errors.Is(err, services.ErrInvalidCredentials) {
			_ = utils.WriteJsonResponse(w, http.StatusUnauthorized, map[string]string{
				"error": "invalid email or password",
			})
			return
		}

		_ = utils.WriteJsonResponse(w, http.StatusInternalServerError, map[string]string{
			"error": "unable to login user",
		})
		return
	}

	_ = utils.WriteJsonResponse(w, http.StatusOK, map[string]string{
		"message": "User logged in successfully",
		"token":   token,
	})
}
