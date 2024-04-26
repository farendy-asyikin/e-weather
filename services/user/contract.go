package userservice

import (
	"main.go/models"
	userrepository "main.go/repositories/users"
	"main.go/schemas"
)

type UserService interface {
	CreateUser(request schemas.CreateUserRequest) (*models.User, error)
	UpdateUser(request schemas.UpdateUserRequest, user models.User) (*models.User, error)
	DeleteUserByID(ID string) error
	GetUserByID(ID string) (*models.User, *schemas.DetailUserResponse, error)
}

type userService struct {
	userRepository userrepository.UserRepository
}

func NewUserService(
	userRepository userrepository.UserRepository,
) *userService {
	return &userService{
		userRepository: userRepository,
	}
}
