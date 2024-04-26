package authservice

import (
	userrepository "main.go/repositories/users"
	"main.go/schemas"
)

type AuthService interface {
	Login(email string, password string) (*schemas.LoginResponse, error)
}

type authService struct {
	userRepository userrepository.UserRepository
}

func NewAuthService(userRepository userrepository.UserRepository) *authService {
	return &authService{userRepository}
}
