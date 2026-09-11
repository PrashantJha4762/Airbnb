package db

import (
	"AuthInGoService/models"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	GetById() (*models.User, error)
	Create(username, email, password string) error
	Getall() ([]*models.User, error)
	DeleteById(id int64) error
	GetUserByEmail(email string) (int, string, string, error)
}

type UserRepositoryImpl struct { // Hme UserRepository interface ko implement krna tha uske lie ek
	db *sql.DB // ek struct chahiye thi so hmne ye banaya
}

func (u *UserRepositoryImpl) GetById() (*models.User, error) { //jaise hi hmne ye method banaya ye UserRepository interface ko implement krne lag gya
	//step 1: write the query
	query := "SELECT * FROM users WHERE id = ?"
	//step 2: execute the query
	row := u.db.QueryRow(query, 1) //ye query execute krke ek row return krta h

	//step 3: Process the query result
	user := &models.User{}

	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	fmt.Println("User retrieved successfully", user)
	return user, nil
}
func (u *UserRepositoryImpl) Create(username, email, password string) error {
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"
	result, err := u.db.Exec(query, username, email, password)
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
		fmt.Println("User created successfully")
	}
	return nil
}
func (u *UserRepositoryImpl) Getall() ([]*models.User, error) {
	query := "select * from users"
	rows, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
func (u *UserRepositoryImpl) DeleteById(id int64) error {
	query := "delete from users where id=?"

	result, err := u.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Something wrong happened")
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowAffected == 0 {
		return fmt.Errorf("user not deleted")
	}
	if rowAffected > 0 {
		fmt.Println("User deleted successfully")
	}
	return nil
}

func (u *UserRepositoryImpl) GetUserByEmail(email string) (int, string, string, error) {
	query := "select id, username, password from users where email=?"
	var id int
	var username string
	var hpwd string
	row := u.db.QueryRow(query, email)
	err := row.Scan(&id, &username, &hpwd)
	if err != nil {
		return 0, "", "", err
	}
	return id, username, hpwd, nil
}

func NewUserRepository(_db *sql.DB) UserRepository { //ye constructor function h jo UserRepositoryImpl ka instance create krke return krta h
	return &UserRepositoryImpl{
		db: _db,
	}
}
