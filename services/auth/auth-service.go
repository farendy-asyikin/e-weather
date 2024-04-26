package authservice

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"main.go/schemas"
	"main.go/utils"
)

func (s *authService) Login(email string, password string) (*schemas.LoginResponse, error) {
	user, err := s.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	token, err := utils.GenerateToken(*user)
	if err != nil {
		return nil, err
	}

	response := schemas.LoginResponse{
		Token: *token,
		ID:    user.ID,
		Name:  user.Name,
		Role:  user.Role,
	}

	return &response, nil
}
