// handlers/user_handlers.go
package handlers

import (
	"net/http"
	"subscription-billing-system/models"
	"subscription-billing-system/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

type AuthAPI struct {
	userUseCase models.UserUseCase
}

func NewAuthAPI(userUseCase models.UserUseCase) *AuthAPI {
	return &AuthAPI{userUseCase}
}

// Register untuk register user baru
func (uc *AuthAPI) Register(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("Calling RegisterUser in use case...")
	createdUser, err := uc.userUseCase.RegisterUser(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": createdUser})
}

// Login untuk login dan menghasilkan JWT token
func (uc *AuthAPI) Login(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("Calling Login in use case...")
	user, err := uc.userUseCase.Login(input)
	if err != nil {
		log.Printf("Error logging in user: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Email or Password did'nt match"})
		return
	}

	// Generate JWT
	log.Println("Generate JWT...")
	token, err := middleware.GenerateJWT(user)
	if err != nil {
		log.Printf("Error generating JWT for user %s: %v\n", user.Email, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "data":user})
}