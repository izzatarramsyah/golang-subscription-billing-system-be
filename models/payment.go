package models

import (
	"time"
	"github.com/google/uuid"
)

type Payment struct {
    ID 				uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID 			uuid.UUID `gorm:"type:uuid"`
    SubscriptionID  uuid.UUID `gorm:"type:uuid"`
    Gateway         string // stripe, xendit
    Amount          float64
    Status          string // success, failed, pending
    PaidAt          time.Time
    CreatedAt       time.Time
}

type PaymentDetail struct {
    PaymentID        uuid.UUID
    Gateway          string
    Amount           float64
    Status           string
    PaidAt           time.Time
    UserID           uuid.UUID
    Username         string
    Email            string
    SubscriptionID   uuid.UUID
    SubscriptionStatus string
    StartDate        time.Time
    EndDate          time.Time
}


func (Payment) TableName() string {
	return "subscription.payments"
}   

type PaymentRepository interface {
	Create(payment *Payment) error
	GetByUser(userID uuid.UUID) (*[]Payment, error)
	GetAllPaymentDetails() ([]PaymentDetail, error)
    UpdatePaymentStatus(userID uuid.UUID, status string) (*bool, error)
}

type PaymentUseCase interface {
	ProcessPayment(userID, subscriptionID uuid.UUID, amount float64, paymentMethod string) (*Payment, error)
	GetPayments(userID uuid.UUID) (*[]Payment, error)
	GetAllPaymentDetails() ([]PaymentDetail, error)
    UpdatePaymentStatus(userID uuid.UUID, status string) (*bool, error)
}
