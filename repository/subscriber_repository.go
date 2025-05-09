package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"subscription-billing-system/models"
)

type subscriberRepository struct {
	db *gorm.DB
}

func NewSubscriberRepository(db *gorm.DB) models.SubscriberRepository {
	return &subscriberRepository{db}
}

func (r *subscriberRepository) GetSubscribersByOwner(ownerID uuid.UUID) (*[]models.Subscriber, error) {
	var subscribers []models.Subscriber

	err := r.db.
	Table("subscription.subscriptions").
	Select(`users.username AS user_name, 
            users.email, 
            products.name AS product_name, 
            plans.name AS plan_name, 
			plans.price as plan_price, 
            subscriptions.start_date, 
            subscriptions.end_date`).
	Joins("JOIN subscription.users ON subscription.users.id = subscription.subscriptions.user_id").
	Joins("JOIN subscription.plans ON subscription.plans.id = subscription.subscriptions.plan_id").
	Joins("JOIN subscription.products ON subscription.products.id = subscription.plans.product_id").
	Where("subscription.products.owner_id = ?", ownerID).
	Scan(&subscribers).Error

	if err != nil {
		return nil, err
	}

	return &subscribers, nil
}
