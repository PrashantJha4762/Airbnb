package db

import (
	"AuthInGoService/models"
	"database/sql"
	"fmt"
)

type RoleRepository interface {
	GetRoleById(id int) (*models.Role,error)
	GetRoleByName(name string) (*models.Role,error)
	GetAllRoles() ([]*models.Role,error)
	CreateRole(name, description string) error
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
func (r *RoleRepositoryImpl) GetRoleByName(name string) (*models.Role,error){
	query:="Select id, name,description, created_at, updated_at from roles where name=?"

	row:=r.db.QueryRow(query,name)

	role:=&models.Role{}

	err:=row.Scan(&role.Id,&role.Name,&role.Description,&role.CreatedAt,&role.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return role, nil
}
func (r *RoleRepositoryImpl) GetAllRoles() ([]*models.Role,error){
	query:="Select id, name,description, created_at, updated_at from roles"
	rows,err:=r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	roles:=[]*models.Role{}
	for rows.Next(){
		role:=models.Role{}
		err:=rows.Scan(&role.Id,&role.Name,&role.Description,&role.CreatedAt,&role.UpdatedAt)
		if err != nil {
			return nil, err
		}
		roles = append(roles, &role)
	}
	return roles, nil
}
func (r *RoleRepositoryImpl) CreateRole(name, description string)error{
	query:="Insert into roles(name,description) values(?,?)"
	result,err:=r.db.Exec(query,name,description)
	if err != nil {
		return  err
	}
	rowsaffected,err:=result.RowsAffected()
	if err != nil {
		return  err
	}
	if rowsaffected==0{
		return fmt.Errorf("Role not created")
	}
	if rowsaffected>0{
		fmt.Println("Role created successfully")
	}
	return nil
}
func NewRoleRepository(db *sql.DB) RoleRepository {
	return &RoleRepositoryImpl{
		db:db,
	}
}