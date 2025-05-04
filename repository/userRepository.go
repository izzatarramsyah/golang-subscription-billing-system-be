package repository

import (
	"gorm.io/gorm"
	"subscription-billing-system/models"
)

type userRepository struct {
	db     *gorm.DB
}

func NewUserRepository(db *gorm.DB) models.UserRepository {
	return &userRepository{
		DB:     db,
	}
}

func (r *userRepository) CreateUser(user *models.User) (*models.User, error) {
	err := r.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *UserController) FindUserByUsername(username string) (*models.User, error){
	var user models.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}