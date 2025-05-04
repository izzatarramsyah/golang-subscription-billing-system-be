package server

import (
	"log"
	"os"
	"subscription-billing-system/api/v1/routes"
	"github.com/gin-gonic/gin"
)

// RunServer untuk setup dan menjalankan server
func RunServer() {
	// Setup Gin
	app := gin.Default()

	// Setup Routes
	routes.SetupRouter(app)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on port %s", port)
	app.Run(":" + port)
}
