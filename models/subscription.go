package models

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Subscription struct {
    ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
    UserID     uuid.UUID
    PlanID     uuid.UUID
    StartDate  time.Time
    EndDate    time.Time
    Status     string // active, expired, cancelled
    CreatedAt  time.Time
}

func (Subscription) TableName() string {
	return "subscription.subscriptions"
}
