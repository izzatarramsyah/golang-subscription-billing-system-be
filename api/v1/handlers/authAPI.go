// handlers/user_handlers.go
package handlers

import (
	"net/http"
	"subscription-billing-system/models"
	"subscription-billing-system/middleware"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type UserController struct {
	userUseCase models.UserUseCase
}

func NewUserAPI(userUseCase models.UserUseCase) *UserController {
	return &UserController{userUseCase}
}

// Register untuk register user baru
func (uc *UserController) Register(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := uc.userUseCase.RegisterUser(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": createdUser})
}

// Login untuk login dan menghasilkan JWT token
func (uc *UserController) Login(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.userUseCase.Login(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	// Generate JWT
	token, err := middleware.GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
