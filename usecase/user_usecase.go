package usecase

import (
	"subscription-billing-system/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strings"
	"errors"
	"github.com/google/uuid"
)

type userUseCase struct {
    userRepo   models.UserRepository
}

func NewUserUseCase(userRepo models.UserRepository) models.UserUseCase {
    return &userUseCase{userRepo}
}

func (uc *userUseCase) RegisterUser(input models.User) (*models.User, error) {
	log.Println("Registering new user:", input.Email)

	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password for user %s: %v\n", input.Email, err)
		return nil, err
	}
	input.Password = string(hashedPassword)
	input.Status = "ACTIVE"

	// Ambil username dari email
	if atIndex := strings.Index(input.Email, "@"); atIndex != -1 {
		input.Username = input.Email[:atIndex]
	} else {
		log.Printf("Invalid email format: %s\n", input.Email)
		return nil, errors.New("invalid email format")
	}

	// Call repository to save user
	return uc.userRepo.CreateUser(input)
}

func (uc *userUseCase) Login(input models.User) (*models.User, error) {
	log.Println("Login attempt for user:", input.Email)


	user, err := uc.userRepo.FindUserByUsername(input.Email)
	if err != nil {
		log.Printf("Error finding user: %v\n", err)
		return nil, err
	}

	// Compare Password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		log.Printf("Password mismatch for user %s: %v\n", input.Email, err)
        return nil, err
    }

	log.Println("User logged in successfully:", user.Email)
	return user,nil
}

func (uc *userUseCase) UpdateUser(input models.User) (*models.User, error) {
	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password for user %s: %v\n", input.Email, err)
		return nil, err
	}
	input.Password = string(hashedPassword)
	input.Status = "ACTIVE"

	// Ambil username dari email
	if atIndex := strings.Index(input.Email, "@"); atIndex != -1 {
		input.Username = input.Email[:atIndex]
	} else {
		log.Printf("Invalid email format: %s\n", input.Email)
		return nil, errors.New("invalid email format")
	}
	return uc.userRepo.UpdateUser(input)
}

func (uc *userUseCase) UpdateUserStatus(userID uuid.UUID, status string) (*bool, error) {
	return uc.userRepo.UpdateUserStatus(userID, status)
}

func (uc *userUseCase) GetListUsers() (*[]models.User, error) {
	return uc.userRepo.GetListUsers()
}

func (uc *userUseCase) FindUserByID(userID uuid.UUID) (*models.User, error) {
	return uc.userRepo.FindUserByID(userID)
}

func (uc *userUseCase) FindUserByRole(role string) (*[]models.User, error) {
	return uc.userRepo.FindUserByRole(role)
}
