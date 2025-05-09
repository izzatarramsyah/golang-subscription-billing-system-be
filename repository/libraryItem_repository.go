package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"subscription-billing-system/models"
)

type libraryItemRepository struct {
	db *gorm.DB
}

func NewLibraryItemRepository(db *gorm.DB) models.LibraryItemRepository {
	return &libraryItemRepository{db}
}

func (r *libraryItemRepository) GetUserLibrary(userID uuid.UUID) ([]models.LibraryItem, error) {
	var items []models.LibraryItem

	// Eksekusi query dan simpan hasilnya dalam r.db
	err := r.db.
		Table("subscription.subscriptions").
		Select(`
			subscription.products.name as product_name,
			subscription.products.id as product_id,
			subscription.plans.name as plan_name,
			subscription.plans.price as plan_price,
			subscription.subscriptions.start_date as subscription_start,
			subscription.subscriptions.end_date as subscription_end,
			subscription.subscriptions.status as subscription_status,
			payment.status as payment_status
		`).
		Joins("JOIN subscription.plans ON subscription.subscriptions.plan_id = subscription.plans.id").
		Joins("JOIN subscription.products ON subscription.plans.product_id = subscription.products.id").
		Joins("LEFT JOIN subscription.payments AS payment ON payment.subscription_id = subscription.subscriptions.id").
		Where("subscription.subscriptions.user_id = ?", userID).
		Scan(&items)

	// Jika terjadi error, periksa field Error di r.db
	if err.Error != nil {
		// Kembalikan error yang terjadi
		return nil, err.Error
	}

	// Return hasil query jika berhasil
	return items, nil
}
