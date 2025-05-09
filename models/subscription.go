package models

import (
	"time"
	"github.com/google/uuid"
)

const (
	SubscriptionStatusActive   = "active"
	SubscriptionStatusCanceled = "canceled"
	SubscriptionStatusExpired = "expired"
)

type Subscription struct {
    ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
    UserID     uuid.UUID `gorm:"type:uuid"`
    PlanID     uuid.UUID `gorm:"type:uuid"`
    StartDate  time.Time
    EndDate    time.Time
    Status     string // active, expired, cancelled
    CreatedAt  time.Time
}

func (Subscription) TableName() string {
	return "subscription.subscriptions"
}

type SubscriptionRepository interface {
	Create(sub *Subscription) error
	GetActiveByUser(userID uuid.UUID) (*[]Subscription, error)
	UpdateStatus(subscriptionID uuid.UUID, status string) error
}

type SubscriptionUseCase interface {
	Subscribe(userID uuid.UUID, planID uuid.UUID) (*uuid.UUID, error)
	MySubscription(userID uuid.UUID) (*[]Subscription, error)
	Unsubscribe(subscriptionID uuid.UUID) error
    IsSubscriptionActive(userID uuid.UUID) (bool, error)
}