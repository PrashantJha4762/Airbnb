package db

import (
	"AuthInGoService/models"
	"database/sql"
	"fmt"
)

type PermissionRepository interface {
	GetPermissionById(id int) (*models.Permission,error)
	GetPermissionByName(name string) (*models.Permission,error)
	GetAllPermissions() ([]*models.Permission,error)
	CreatePermission(name, description, resource, action string) error
	DeletePermissionById(id int) error
}

type PermissionRepositoryImpl struct {
	db *sql.DB
}
func (p *PermissionRepositoryImpl) GetPermissionById(id int) (*models.Permission,error){
	query:="Select id, name,description, resource, action, created_at, updated_at from permissions where id=?"
	row:=p.db.QueryRow(query,id)
	permission:=&models.Permission{}
	err:=row.Scan(&permission.Id,&permission.Name,&permission.Description,&permission.Resource,&permission.Action,&permission.CreatedAt,&permission.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return permission, nil
}
func (p *PermissionRepositoryImpl) GetPermissionByName(name string) (*models.Permission,error){
	query:="Select id, name,description, resource, action, created_at, updated_at from permissions where name=?"
	row:=p.db.QueryRow(query,name)
	permission:=&models.Permission{}
	err:=row.Scan(&permission.Id,&permission.Name,&permission.Description,&permission.Resource,&permission.Action,&permission.CreatedAt,&permission.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return permission, nil
}
func (p *PermissionRepositoryImpl) GetAllPermissions() ([]*models.Permission,error){
	query:="Select id, name,description, resource, action, created_at, updated_at from permissions"
	rows,err:=p.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	permissions:=[]*models.Permission{}
	for rows.Next(){
		permission:=models.Permission{}
		err:=rows.Scan(&permission.Id,&permission.Name,&permission.Description,&permission.Resource,&permission.Action,&permission.CreatedAt,&permission.UpdatedAt)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, &permission)
	}
	return permissions, nil
}
func (p *PermissionRepositoryImpl) CreatePermission(name, description, resource, action string) error {
	query:="INSERT INTO permissions (name, description, resource, action) VALUES (?, ?, ?, ?)"
	result, err := p.db.Exec(query, name, description, resource, action)
	if err != nil {
		return err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowAffected == 0 {
		return fmt.Errorf("no rows affected")
	}
	if rowAffected > 0 {
		fmt.Println("Permission created successfully")
	}
	return nil
}
func (p *PermissionRepositoryImpl) DeletePermissionById(id int) error {
	query:="DELETE FROM permissions WHERE id=?"
	rows,err:=p.db.Exec(query,id)
	if err != nil {
		return err
	}
	rowAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}	
	if rowAffected == 0 {
		return fmt.Errorf("no rows affected")
	}
	return nil
}
func NewPermissionRepository(db *sql.DB) PermissionRepository {
	return &PermissionRepositoryImpl{
		db: db,
	}
}