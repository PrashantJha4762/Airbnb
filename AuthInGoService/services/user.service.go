package services

import (
	config "AuthInGoService/config/env"
	db "AuthInGoService/db/repositories"
	"AuthInGoService/models"
	"AuthInGoService/utils"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	GetUserById(id int) (*models.User, error)
	CreateUser(username, email, password string) error
	LoginUser(email, password string) (string, error)
}

var ErrInvalidCredentials = errors.New("invalid credentials")

type UserServiceImpl struct {
	userRepository db.UserRepository //the service depend on the repo interface rather than struct so that we can easily mock the repo interface in the test cases and we can easily change the implementation of the repo interface without changing the service code.
}

func (u *UserServiceImpl) GetUserById(id int) (*models.User, error) {
	user, err := u.userRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserServiceImpl) CreateUser(username, email, password string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}
	errr := u.userRepository.Create(username, email, hashedPassword)
	if errr != nil {
		return errr
	}
	return nil
}
func (u *UserServiceImpl) LoginUser(email, password string) (string, error) {

	//step 1: get the user by email from the repo layer
	userID, username, hashedpwd, err := u.userRepository.GetUserByEmail(email)
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
		"user_id":  userID,
		"email":    email,
		"username": username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
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
