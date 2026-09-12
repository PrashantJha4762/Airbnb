package db

import "database/sql"

type RoleRepository interface {
	GetRoleById(id int) (*models.Role,error)
}
type RoleRepositoryImpl struct {
	db *sql.DB
}
func NewRoleRepository(db *sql.DB) RoleRepository {
	return &RoleRepositoryImpl{
		db:db,
	}
}