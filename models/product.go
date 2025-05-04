package models

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
    ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
    Name        string
    Description string
    OwnerID     uuid.UUID // FK ke User
    CreatedAt   time.Time
}

func (Product) TableName() string {
	return "subscription.products"
}
