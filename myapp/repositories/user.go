package repositories

import (
	"errors"
	"myapp/database"
	"myapp/models"

	"gorm.io/gorm"
)

// UserRepository is a struct to represent the user repository.
type UserRepository struct{}

// CreateUser inserts a new user into the database.
func (r *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	if err := database.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByID retrieves a user by their ID from the database.
func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	var user models.User
	result := database.DB.First(&user, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// UpdateUser updates an existing user in the database.
func (r *UserRepository) UpdateUser(user *models.User) error {
	if err := database.DB.Save(user).Error; err != nil {
		return err
	}
	return nil
}

// DeleteUser removes a user from the database.
func (r *UserRepository) DeleteUser(id int) error {
	result := database.DB.Delete(&models.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no user found to delete")
	}
	return nil
}

// ListUsers retrieves all users from the database.
func (r *UserRepository) ListUsers() ([]models.User, error) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
