package services

import (
	db "AuthInGoService/db/repositories"
	"AuthInGoService/models"
	"AuthInGoService/utils"
	"fmt"
)

type UserService interface {
	GetUserById() (*models.User,error) 
	CreateUser(username, email, password string) (error)
}
type UserServiceImpl struct {
	userRepository db.UserRepository //the service depend on the repo interface rather than struct so that we can easily mock the repo interface in the test cases and we can easily change the implementation of the repo interface without changing the service code.
}

func (u *UserServiceImpl) GetUserById() (*models.User,error) {
	user,err:=u.userRepository.GetById()
	if err!=nil{
		return nil,err
	}
	return user,nil
}

func (u *UserServiceImpl) CreateUser(username, email, password string) error {
	hashedPassword, err :=utils.HashPassword(password)
	if err!=nil{
		fmt.Println("Not able to hash password");
	}
	errr:=u.userRepository.Create(username, email, hashedPassword)
	if errr!=nil{
		return err
	}
	return nil
}
//This fn is basically facilitating the dependency injection. we are providing the instance 
//of the userRepository to the UserServiceImpl struct so that we can use it in the CreateUser fn externally
func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}
