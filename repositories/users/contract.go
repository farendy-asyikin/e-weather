package userrepository

import (
	"gorm.io/gorm"
	"main.go/models"
)

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user models.User) (*models.User, error)
	GetUserByID(ID string) (*models.User, error)
	GetUserBySpecificColumn(column string, value string) (*models.User, error)
	UpdateUserRoleBySuperuser(user models.User, userID uint64) (*models.User, error)
	UpdateUser(user models.User) (*models.User, error)
	DeleteUserByID(ID string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}
