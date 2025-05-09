package models

import (
	"time"
    "github.com/google/uuid"
)

type User struct {
    ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
    Username  string    `gorm:"unique;not null"`
    Email     string    `gorm:"unique;not null"`
    Status    string    `gorm:"not null;default:'ACTIVE'"` 
    Password  string    `gorm:"not null"`
    Role      string    `gorm:"not null"` // admin, business_owner, subscriber
    CreatedAt time.Time
}

func (User) TableName() string {
	return "subscription.users"
}

type UserRepository interface {
	CreateUser(user User) (*User, error)
    GetListUsers() (*[]User, error)
    FindUserByUsername(email string) (*User, error)
    UpdateUser(user User) (*User, error)
    UpdateUserStatus(userID uuid.UUID, status string) (*bool, error)
    FindUserByID(userID uuid.UUID) (*User, error)
    FindUserByRole(role string) (*[]User, error)
}

type UserUseCase interface {
	RegisterUser(user User) (*User, error)
    Login(user User) (*User, error)
    GetListUsers() (*[]User, error)
    UpdateUser(user User) (*User, error)
    UpdateUserStatus(userID uuid.UUID, status string) (*bool, error)
    FindUserByID(userID uuid.UUID) (*User, error)
    FindUserByRole(role string) (*[]User, error)
}
