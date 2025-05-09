package models

import (
	"time"
	"github.com/google/uuid"
)

type LibraryItem struct {
	ProductID           uuid.UUID   `json:"product_id" db:"product_id"`
    ProductName         string      `json:"product_name" db:"product_name"`
    PlanName            string      `json:"plan_name" db:"plan_name"`
    PlanPrice           float64     `json:"plan_price" db:"plan_price"`
    SubscriptionStart   time.Time   `json:"subscription_start" db:"subscription_start"`
    SubscriptionEnd     time.Time   `json:"subscription_end" db:"subscription_end"`
    SubscriptionStatus  string      `json:"subscription_status" db:"subscription_status"`
    PaymentStatus       string      `json:"payment_status" db:"payment_status"`
}

type LibraryItemRepository interface {
	GetUserLibrary(userID uuid.UUID) ([]LibraryItem, error)
}

type LibraryItemUseCase interface {
	GetUserLibrary(userID uuid.UUID) ([]LibraryItem, error)
}
