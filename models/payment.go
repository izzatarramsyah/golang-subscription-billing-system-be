package models

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payment struct {
    ID              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
    SubscriptionID  uuid.UUID
    Gateway         string // stripe, xendit
    Amount          float64
    Status          string // success, failed, pending
    PaidAt          time.Time
    CreatedAt       time.Time
}

func (Payment) TableName() string {
	return "subscription.payments"
}   