package router

import (
	controllers "AuthInGoService/controllers"
	db "AuthInGoService/db/repositories"
	middleware "AuthInGoService/middleware"

	"github.com/go-chi/chi/v5"
)

type RoleRouter struct {
	rolecontroller     *controllers.RoleController
	userRoleRepository db.UserRoleRepository
}

func NewRoleRouter(_rolecontroller *controllers.RoleController, userRoleRepository db.UserRoleRepository) Router {
	return &RoleRouter{
		rolecontroller:     _rolecontroller,
		userRoleRepository: userRoleRepository,
	}
}

func (rr *RoleRouter) Register(r *chi.Mux) {
	admin := middleware.RequireAllRolesWithRepository(rr.userRoleRepository, "admin")
	manageRoles := middleware.RequireAllPermissionsWithRepository(rr.userRoleRepository, "role:manage")
	r.With(middleware.JwtAuth, admin).Get("/roles/{id}", rr.rolecontroller.GetRoleById)
	r.With(middleware.JwtAuth, admin).Get("/roles", rr.rolecontroller.GetAllRoles)
	r.With(middleware.JwtAuth, admin, manageRoles).Post("/roles", rr.rolecontroller.CreateRole)
	r.With(middleware.JwtAuth, admin, manageRoles).Delete("/roles/{id}", rr.rolecontroller.DeleteRoleById)
	r.With(middleware.JwtAuth, admin, manageRoles).Post("/roles/assign/{userId}", rr.rolecontroller.AssignRoleToUser)
	r.With(middleware.JwtAuth, admin, manageRoles).Post("/roles/remove/{userId}", rr.rolecontroller.RemoveRoleFromUser)
}
