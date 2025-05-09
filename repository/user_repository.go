package repository

import (
	"gorm.io/gorm"
	"subscription-billing-system/models"
	"log"
	"github.com/google/uuid"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) models.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(user models.User) (*models.User, error) {
	log.Println("Creating new user with email:", user.Email)
	err := r.db.Create(&user).Error
	if err != nil {
		log.Printf("Error creating user: %v\n", err)
		return nil, err
	}
	log.Println("User created successfully:", user.Email)
	return &user, nil
}

func (r *userRepository) FindUserByUsername(email string) (*models.User, error){
	log.Println("Finding user by email:", email)
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		log.Printf("Error finding user: %v\n", err)
		return nil, err
	}
	log.Println("User found:", user.Email)
	return &user, nil
}

func (r *userRepository) UpdateUser(user models.User) (*models.User, error) {
	log.Println("Updating user with email:", user.Email)

	err := r.db.Save(&user).Error 
	if err != nil {
		log.Printf("Error updating user: %v\n", err)
		return nil, err
	}

	// Re-fetch updated user from DB to ensure fresh data
	var updatedUser models.User
	if err := r.db.Where("email = ?", user.Email).First(&updatedUser).Error; err != nil {
		log.Printf("Error fetching updated user: %v\n", err)
		return nil, err
	}

	log.Println("User updated successfully:", updatedUser.Email)
	return &updatedUser, nil
}

func (r *userRepository) UpdateUserStatus(userID uuid.UUID, status string) (*bool, error) {

	err := r.db.Model(&models.User{}).Where("id = ?", userID).Update("status", status).Error
	if err != nil {
		result := false
		return &result, nil
	}

	result := true
	return &result, nil
}

func (r *userRepository) GetListUsers() (*[]models.User, error){
	log.Println("Get List User:")
	var users []models.User
	if err := r.db.Where("role = ?", "subscriber").Find(&users).Error; err != nil {
		log.Printf("Error get list user: %v\n", err)
		return nil, err
	}
	log.Println("get list user successfully")
	return &users, nil
}

func (r *userRepository) FindUserByID(userID uuid.UUID) (*models.User, error){
	log.Println("Get User:")
	var users models.User
	if err := r.db.Where("id = ?", userID).Find(&users).Error; err != nil {
		log.Printf("Error get user: %v\n", err)
		return nil, err
	}
	log.Println("get user successfully")
	return &users, nil
}

func (r *userRepository) FindUserByRole(role string) (*[]models.User, error){
	log.Println("Get User:")
	var users []models.User
	if err := r.db.Where("role = ?", role).Find(&users).Error; err != nil {
		log.Printf("Error get user: %v\n", err)
		return nil, err
	}
	log.Println("get user successfully")
	return &users, nil
}