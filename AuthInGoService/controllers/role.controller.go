package controllers

import (
	services "AuthInGoService/services"
	"encoding/json"
	"net/http"
	"strconv"
)

type RoleController struct {
	roleService services.RoleService
}

func NewRoleController(roleService services.RoleService) *RoleController {
	return &RoleController{
		roleService: roleService,
	}
}

func (rc *RoleController) GetRoleById(w http.ResponseWriter, r *http.Request) {
	roleId := r.URL.Query().Get("id")
	roleid,err:= strconv.Atoi(roleId)
	if err != nil {
		http.Error(w, "Invalid role ID", http.StatusBadRequest)
		return
	}
	if roleId == "" {
		http.Error(w, "Missing role ID", http.StatusBadRequest)
		return
	}
	role, err := rc.roleService.GetRoleById((roleid))
	if err != nil {
		http.Error(w, "Role not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(role)
}
func (rc *RoleController) GetAllRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := rc.roleService.GetAllRoles()
	if err != nil {
		http.Error(w, "Failed to fetch roles", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(roles)
}