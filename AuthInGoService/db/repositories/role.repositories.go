package db

import (
	"AuthInGoService/models"
	"database/sql"
)

type RoleRepository interface {
	GetRoleById(id int) (*models.Role,error)
}
type RoleRepositoryImpl struct {
	db *sql.DB
}

func (r *RoleRepositoryImpl) GetRoleById(id int) (*models.Role,error){
	query:="Select id, name,description, created_at, updated_at from roles where id=?"

	row:=r.db.QueryRow(query,id)

	role:=&models.Role{}

	err:=row.Scan(&role.Id,&role.Name,&role.Description,&role.CreatedAt,&role.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return role, nil
}
func NewRoleRepository(db *sql.DB) RoleRepository {
	return &RoleRepositoryImpl{
		db:db,
	}
}