package models

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Plan struct {
    ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
    ProductID     uuid.UUID // FK ke Product
    Name          string
    Price         float64
    DurationMonths int // 1 (monthly), 12 (yearly), etc
    CreatedAt     time.Time
}

func (Plan) TableName() string {
	return "subscription.plans"
}   