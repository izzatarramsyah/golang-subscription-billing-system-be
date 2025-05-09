package repository

import (
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"subscription-billing-system/models"
)

type subscriptionRepository struct {
	db *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) models.SubscriptionRepository {
	return &subscriptionRepository{db}
}

func (r *subscriptionRepository) Create(sub *models.Subscription) error {
	log.Println("Creating new subscription")
	if err := r.db.Create(sub).Error; err != nil {
		log.Printf("Error creating subscription: %v\n", err)
		return err
	}
	log.Println("Subscription created successfully")
	return nil
}

func (r *subscriptionRepository) GetActiveByUser(userID uuid.UUID) (*[]models.Subscription, error) {
	log.Println("Get active subscriptions by user ID:", userID)
	var subs []models.Subscription
	err := r.db.Where("user_id = ? AND status = ?", userID, "active").Find(&subs).Error
	if err != nil {
		log.Printf("Error getting active subscriptions: %v\n", err)
		return nil, err
	}
	return &subs, nil
}

func (r *subscriptionRepository) UpdateStatus(subscriptionID uuid.UUID, status string) error {
	log.Println("Updating status of subscription ID:", subscriptionID)
	err := r.db.Model(&models.Subscription{}).
		Where("id = ?", subscriptionID).
		Update("status", status).Error

	if err != nil {
		log.Printf("Failed to update status for subscription ID %s: %v\n", subscriptionID, err)
	} else {
		log.Printf("Successfully updated status for subscription ID: %s\n", subscriptionID)
	}
	return err
}
