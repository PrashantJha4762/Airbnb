package router

import (
	controllers "AuthInGoService/controllers"

	"github.com/go-chi/chi/v5"
)

type UserRouter struct{
	usercontroller controllers.UserController
}

func NewUserRouter(_usercontroller controllers.UserController) Router{
	return &UserRouter{
		usercontroller: _usercontroller,
	}
}
func (u *UserRouter) Register(r *chi.Mux) {
	r.Post("/register",u.usercontroller.Register)
}