package router

import (
	controllers "AuthInGoService/controllers"
	db "AuthInGoService/db/repositories"
	middleware "AuthInGoService/middleware"

	"github.com/go-chi/chi/v5"
)

type UserRouter struct {
	usercontroller     *controllers.UserController
	userRoleRepository db.UserRoleRepository
}

func NewUserRouter(_usercontroller *controllers.UserController, userRoleRepository db.UserRoleRepository) Router {
	return &UserRouter{
		usercontroller:     _usercontroller,
		userRoleRepository: userRoleRepository,
	}
}
func (u *UserRouter) Register(r *chi.Mux) {
	r.With(middleware.JwtAuth, middleware.RequireAllPermissionsWithRepository(u.userRoleRepository, "user:read")).Get("/profile", u.usercontroller.GetUserInfo)
	r.Post("/signup", u.usercontroller.CreateUser)
	r.Post("/login", u.usercontroller.LoginUser)
}
