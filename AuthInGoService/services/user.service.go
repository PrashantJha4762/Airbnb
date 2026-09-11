package services

import (
	config "AuthInGoService/config/env"
	db "AuthInGoService/db/repositories"
	"AuthInGoService/models"
	"AuthInGoService/utils"
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	GetUserById() (*models.User, error)
	CreateUser(username, email, password string) error
	LoginUser(email, password string) (string, error)
}

var ErrInvalidCredentials = errors.New("invalid credentials")

type UserServiceImpl struct {
	userRepository db.UserRepository //the service depend on the repo interface rather than struct so that we can easily mock the repo interface in the test cases and we can easily change the implementation of the repo interface without changing the service code.
}

func (u *UserServiceImpl) GetUserById() (*models.User, error) {
	user, err := u.userRepository.GetById()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserServiceImpl) CreateUser(username, email, password string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		fmt.Println("Not able to hash password")
	}
	errr := u.userRepository.Create(username, email, hashedPassword)
	if errr != nil {
		return errr
	}
	return nil
}
func (u *UserServiceImpl) LoginUser(email, password string) (string, error) {

	//step 1: get the user by email from the repo layer
	username, hashedpwd, err := u.userRepository.GetUserByEmail(email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrInvalidCredentials
		}
		return "", err
	}

	if !utils.CheckPassword(hashedpwd, password) {
		return "", ErrInvalidCredentials
	}

	// The credentials are valid, so create a token for this session.
	payload := jwt.MapClaims{
		"email":    email,
		"username": username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenstring, err := token.SignedString([]byte(config.GetString("JWT_SECRET_KEY", "token")))
	if err != nil {
		return "", err
	}
	return tokenstring, nil
}

// This fn is basically facilitating the dependency injection. we are providing the instance
// of the userRepository to the UserServiceImpl struct so that we can use it in the CreateUser fn externally
func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}
