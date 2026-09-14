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
	r.Get("/roles/{id}",rr.rolecontroller.GetRoleById);
	r.Get("/roles",rr.rolecontroller.GetAllRoles);
	r.Post("/roles",rr.rolecontroller.CreateRole);
	r.Delete("/roles/{id}",rr.rolecontroller.DeleteRoleById);
	r.Post("/roles/assign",rr.rolecontroller.AssignRoleToUser);
	r.Post("/roles/remove",rr.rolecontroller.RemoveRoleFromUser);
}