package main

import (
	"log"
	"subscription-billing-system/infra/database"
	"subscription-billing-system/models"
	"subscription-billing-system/infra/server"
)

func main() {
	// Connect to the database

	log.Println("Connecting to database...")
	db := database.ConnectDatabase()
	log.Printf("DB: %+v\n", database.DB)

	// Migrate models
	err := db.AutoMigrate(
		&models.User{},
		&models.Plan{},
		&models.Subscription{},
		&models.Payment{},
		&models.Product{},
		&models.Reminder{},
		&models.Ebook{},
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	} else {
		log.Println("Migration executed successfully")
	}

	// Run the server
	server.RunServer(db)
}