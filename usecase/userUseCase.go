package usecase

import (
	"subscription-billing-system/models"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
    userRepo   models.UserRepository
}

func NewUserUseCase(userRepo models.UserRepository) models.UserUseCase {
    return &userUseCase{userRepo}
}

func (uc *userUseCase) RegisterUser(input *models.User) (*models.User, error) {
	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)

	// Call repository to save user
	return uc.userRepo.CreateUser(user)
}

func (uc *userUseCase) Login(input *models.User) (*models.User, error) {
	user, err := uc.userRepo.FindUserByUsername(user.username)
	if err != nil {
		return nil, err
	}

	// Compare Password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        return nil, err
    }

	return user,nill
}