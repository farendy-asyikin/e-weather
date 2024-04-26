package userrepository

import (
	"errors"
	"gorm.io/gorm"
	"main.go/models"
)

func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}

		return nil, err
	}

	return &user, nil
}

func (r *userRepository) CreateUser(user models.User) (*models.User, error) {

	err := r.db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetUserByID(ID string) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}

		return nil, err
	}

	return &user, nil
}

func (r *userRepository) UpdateUser(user models.User) (*models.User, error) {
	err := r.db.Model(&user).Updates(map[string]any{
		"name":     user.Name,
		"email":    user.Email,
		"password": user.Password,
	}).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) UpdateUserRoleBySuperuser(user models.User, userID uint64) (*models.User, error) {
	err := r.db.Model(&user).Updates(map[string]any{
		"role": user.Role,
	}).Where("id", userID).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) DeleteUserByID(ID string) error {
	var user models.User
	err := r.db.Model(&user).Where("id = ?", ID).Updates(map[string]interface{}{"is_active": false}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetUserBySpecificColumn(column string, value string) (*models.User, error) {
	var user models.User
	err := r.db.Where(column+" = ?", value).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}
