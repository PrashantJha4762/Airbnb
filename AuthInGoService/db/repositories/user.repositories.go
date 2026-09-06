package db

import "database/sql"

type UserRepository interface {
	Create() error
}

type UserRepositoryImpl struct { // Hme UserRepository interface ko implement krna tha uske lie ek
	db *sql.DB                  // ek struct chahiye thi so hmne ye banaya
}

func (u *UserRepositoryImpl) Create() error { //jaise hi hmne ye method banaya ye UserRepository interface ko implement krne lag gya
	return nil
}

func NewUserRepository(_db *sql.DB) UserRepository { //ye constructor function h jo UserRepositoryImpl ka instance create krke return krta h
	return &UserRepositoryImpl{
		db: _db,
	}
}