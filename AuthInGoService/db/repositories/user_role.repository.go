package db

import (
	"AuthInGoService/models"
	"database/sql"
)

type UserRoleRepository interface {
	GetUserRoles(userId int) ([]*models.Role, error)
	AssignRoleToUser(userId, roleId int) error
	RemoveRoleFromUser(userId int64, roleId int) error
	GetUserPermissions(userId int) ([]*models.Permission, error)
	HasPermission(userId int, permissionName string) (bool, error)
	HasRole(userId int, roleName string) (bool, error)
}
type UserRoleRepositoryImpl struct {
	db *sql.DB
}

func (ur *UserRoleRepositoryImpl) GetUserRoles(userId int) ([]*models.Role, error) {
	query := "SELECT r.id, r.name, r.description, r.created_at, r.updated_at FROM roles r INNER JOIN user_roles ur ON r.id = ur.role_id WHERE ur.user_id = ?"
	rows, err := ur.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	roles := []*models.Role{}
	for rows.Next() {
		role := &models.Role{}
		err := rows.Scan(&role.Id, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}
func (ur *UserRoleRepositoryImpl) AssignRoleToUser(userId, roleId int) error {
	query := "INSERT INTO user_roles (user_id, role_id) VALUES (?, ?)"
	_, err := ur.db.Exec(query, userId, roleId)
	if err != nil {
		return err
	}
	return nil
}
func (ur *UserRoleRepositoryImpl) RemoveRoleFromUser(userId int64, roleId int) error {
	query := "DELETE FROM user_roles WHERE user_id = ? AND role_id = ?"
	_, err := ur.db.Exec(query, userId, roleId)
	if err != nil {
		return err
	}	
	return nil
}
func (ur *UserRoleRepositoryImpl) GetUserPermissions(userId int) ([]*models.Permission, error) {
	query := `
		SELECT p.id, p.name, p.description, p.resource, p.action, p.created_at, p.updated_at
		FROM permissions p
		INNER JOIN role_permissions rp ON p.id = rp.permission_id
		INNER JOIN user_roles ur ON rp.role_id = ur.role_id
		WHERE ur.user_id = ?`
	rows, err := ur.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	permissions := []*models.Permission{}
	for rows.Next() {
		permission := &models.Permission{}
		err := rows.Scan(&permission.Id, &permission.Name, &permission.Description, &permission.Resource, &permission.Action, &permission.CreatedAt, &permission.UpdatedAt)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}

	return permissions, nil
}
func (ur *UserRoleRepositoryImpl) HasPermission(userId int, permissionName string) (bool, error) {
	query := `
		SELECT COUNT(*) FROM permissions p
		INNER JOIN role_permissions rp ON p.id = rp.permission_id
		INNER JOIN user_roles ur ON rp.role_id = ur.role_id
		WHERE ur.user_id = ? AND p.name = ?`
	var count int
	err := ur.db.QueryRow(query, userId, permissionName).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
func (ur *UserRoleRepositoryImpl) HasRole(userId int, roleName string) (bool, error) {
	query := `
		SELECT COUNT(*) FROM roles r
		INNER JOIN user_roles ur ON r.id = ur.role_id
		WHERE ur.user_id = ? AND r.name = ?`
	var count int
	err := ur.db.QueryRow(query, userId, roleName).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func NewUserRoleRepository(db *sql.DB) UserRoleRepository {
	return &UserRoleRepositoryImpl{
		db: db,
	}
}