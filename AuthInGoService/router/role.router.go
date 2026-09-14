package router

import (
	controllers "AuthInGoService/controllers"

	"github.com/go-chi/chi/v5"
)

type RoleRouter struct {
	rolecontroller *controllers.RoleController
}

func NewRoleRouter(_rolecontroller *controllers.RoleController) Router {
	return &RoleRouter{
		rolecontroller: _rolecontroller,
	}
}

func (rr *RoleRouter) Register(r *chi.Mux) {
	r.Get("/role",rr.rolecontroller.GetRoleById);
}