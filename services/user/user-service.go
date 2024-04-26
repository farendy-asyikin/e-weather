package userservice

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"main.go/models"
	"main.go/schemas"
	"main.go/utils"
)

func (s *userService) CreateUser(request schemas.CreateUserRequest) (*models.User, error) {
	if request.Password != request.ConfirmPassword {
		return nil, errors.New("password and confirm password must be same")
	}

	if utils.IsContainsNumber(request.Name) {
		return nil, errors.New("name must be string")
	}

	isEmailExist, err := s.userRepository.GetUserBySpecificColumn("email", request.Email)
	if err != nil {
		return nil, err
	}
	if isEmailExist != nil {
		return nil, errors.New("email already exist")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: string(passwordHash),
	}

	result, err := s.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *userService) GetUserByID(ID string) (*models.User, *schemas.DetailUserResponse, error) {
	user, err := s.userRepository.GetUserByID(ID)
	if err != nil {
		return nil, nil, err
	}

	response := schemas.DetailUserResponse{
		Name:  user.Name,
		Email: user.Email,
	}

	return user, &response, nil
}

func (s *userService) UpdateUser(request schemas.UpdateUserRequest, user models.User) (*models.User, error) {
	if request.Name != nil {
		if utils.IsContainsNumber(*request.Name) {
			return nil, errors.New("name must be string")
		}
		user.Name = *request.Name
	}

	if request.Email != nil {

		isEmailExist, err := s.userRepository.GetUserBySpecificColumn("email", *request.Email)
		if err != nil {
			return nil, err
		}
		if isEmailExist != nil && isEmailExist.ID != user.ID {
			return nil, errors.New("email already exist")
		}
		user.Email = *request.Email
	}

	result, err := s.userRepository.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *userService) DeleteUserByID(ID string) error {
	err := s.userRepository.DeleteUserByID(ID)
	if err != nil {
		return err
	}

	return nil
}
