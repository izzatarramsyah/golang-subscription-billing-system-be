package models

import (
	"time"
    "github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
    ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
    Email     string    `gorm:"unique;not null"`
    Password  string    `gorm:"not null"`
    Role      string    `gorm:"not null"` // admin, business_owner, subscriber
    CreatedAt time.Time
}

func (User) TableName() string {
	return "subscription.users"
}

type UserRepository interface {
	CreateUser(user User) (*User, error)
    FindUserByUsername(username string) (*User, error)
}

type UserUseCase interface {
	RegisterUser(user models.User) (User, error)
    Login(user models.User) (User, error)
}
