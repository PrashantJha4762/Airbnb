package services

import db "AuthInGoService/db/repositories"

type UserService interface {
	CreateUser() error
}
type UserServiceImpl struct {
	userRepository db.UserRepository //the service depend on the repo interface rather than struct so that we can easily mock the repo interface in the test cases and we can easily change the implementation of the repo interface without changing the service code.
}

func (u *UserServiceImpl) CreateUser() error {
	return nil
}
//This fn is basically fecilitiating the dependency injection. we are providing the instance 
//od the userRepository to the UserServiceImpl struct so that we can use it in the CreateUser fn externally
func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}
