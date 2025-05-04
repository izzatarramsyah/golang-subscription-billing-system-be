package routes

import (
	"subscription-billing-system/api/v1/handlers"
	"subscription-billing-system/repository"
	"subscription-billing-system/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Inisialisasi layer
	userRepo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userAPI := handlers.NewUserAPI(userUseCase)

	// API Grouping versi 1
	v1 := router.Group("/api/v1")
	{
		v1.POST("/register", userAPI.Register)
		v1.POST("/login", userAPI.Login) 
	}

	return router
}
