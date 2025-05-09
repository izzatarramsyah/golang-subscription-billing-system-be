package server

import (
	"log"
	"os"
	"subscription-billing-system/api/v1/routes"
	"gorm.io/gorm"
)

// RunServer untuk setup dan menjalankan server
func RunServer(db *gorm.DB) {
	router := routes.SetupRouter(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	router.Run(":" + port)
}