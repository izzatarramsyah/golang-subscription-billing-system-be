package main

import (
	"log"
	"os"
	"subscription-billing-system/infra/database"
	"subscription-billing-system/models"
	"subscription-billing-system/server"
	"subscription-billing-system/api/v1/routes"
)

func main() {
	// Connect to the database
	database.ConnectDatabase()

	// Migrate models
	database.DB.AutoMigrate(&models.User{}, &models.Plan{}, &models.Subscription{}, &models.Payment{}, &models.Product{})

	// Set up and run the server
	server.RunServer()
}
