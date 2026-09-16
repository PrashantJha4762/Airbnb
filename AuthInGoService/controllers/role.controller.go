package controllers

import (
	"AuthInGoService/dto"
	services "AuthInGoService/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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
	roleId := chi.URLParam(r, "id")
	if roleId == "" {
		http.Error(w, "Missing role ID", http.StatusBadRequest)
		return
	}
	roleid, err := strconv.Atoi(roleId)
	if err != nil {
		http.Error(w, "Invalid role ID", http.StatusBadRequest)
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

func (rc *RoleController) CreateRole(w http.ResponseWriter, r *http.Request) {
	var payload dto.CreateRoleRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil || payload.Name == "" || payload.Description == "" {
		http.Error(w, "Invalid role payload", http.StatusBadRequest)
		return
	}

	err := rc.roleService.CreateRole(payload.Name, payload.Description)
	if err != nil {
		http.Error(w, "Failed to create role", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Role created successfully"})
}
func (rc *RoleController) DeleteRoleById(w http.ResponseWriter, r *http.Request) {
	roleId := chi.URLParam(r, "id")
	if roleId == "" {
		http.Error(w, "Missing role ID", http.StatusBadRequest)
		return
	}
	roleid, err := strconv.Atoi(roleId)
	if err != nil {
		http.Error(w, "Invalid role ID", http.StatusBadRequest)
		return
	}
	err = rc.roleService.DeleteRoleById(roleid)
	if err != nil {
		http.Error(w, "Failed to delete role", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Role deleted successfully"})
}
func (rc *RoleController) AssignRoleToUser(w http.ResponseWriter, r *http.Request) {
	var payload dto.AssignRoleRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil || payload.RoleID < 1 {
		http.Error(w, "Invalid role assignment payload", http.StatusBadRequest)
		return
	}
	userIdStr := chi.URLParam(r, "userId")
	if userIdStr == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	err = rc.roleService.AssignRoleToUser(userId, int(payload.RoleID))
	if err != nil {
		http.Error(w, "Failed to assign role to user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Role assigned to user successfully"})
}
func (rc *RoleController) RemoveRoleFromUser(w http.ResponseWriter, r *http.Request) {
	var payload dto.RemoveRoleRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil || payload.RoleID < 1 {
		http.Error(w, "Invalid role removal payload", http.StatusBadRequest)
		return
	}
	userIdStr := chi.URLParam(r, "userId")
	if userIdStr == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	err = rc.roleService.RemoveRoleFromUser(int64(userId), int(payload.RoleID))
	if err != nil {
		http.Error(w, "Failed to remove role from user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Role removed from user successfully"})
}
