package services

import (
	db "AuthInGoService/db/repositories"
	"AuthInGoService/models"
)

type RoleService interface {
	GetRoleById(id int) (*models.Role, error)
	GetRoleByName(name string) (*models.Role, error)
	GetAllRoles() ([]*models.Role, error)
	CreateRole(name, description string) error
	DeleteRoleById(id int) error
	AssignRoleToUser(userId, roleId int) error
	RemoveRoleFromUser(userId int64, roleId int) error
}

type RoleServiceImpl struct {
	roleRepository     db.RoleRepository
	userRoleRepository db.UserRoleRepository
}

func (r *RoleServiceImpl) GetRoleById(id int) (*models.Role, error) {
	return r.roleRepository.GetRoleById(id)
}

func (r *RoleServiceImpl) GetRoleByName(name string) (*models.Role, error) {
	return r.roleRepository.GetRoleByName(name)
}

func (r *RoleServiceImpl) GetAllRoles() ([]*models.Role, error) {
	return r.roleRepository.GetAllRoles()
}

func (r *RoleServiceImpl) CreateRole(name, description string) error {
	return r.roleRepository.CreateRole(name, description)
}

func (r *RoleServiceImpl) DeleteRoleById(id int) error {
	return r.roleRepository.DeleteRoleById(id)
}

func (r *RoleServiceImpl) AssignRoleToUser(userId, roleId int) error {
	return r.userRoleRepository.AssignRoleToUser(userId, roleId)
}

func (r *RoleServiceImpl) RemoveRoleFromUser(userId int64, roleId int) error {
	return r.userRoleRepository.RemoveRoleFromUser(userId, roleId)
}

// NewRoleService creates a role service with its repository dependencies.
func NewRoleService(roleRepository db.RoleRepository) RoleService {
	return &RoleServiceImpl{
		roleRepository:     roleRepository,
	}
}
