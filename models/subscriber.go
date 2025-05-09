package models

import (
	"github.com/google/uuid"
)


type Subscriber struct {
	UserName           string `json:"user_name"`
	Email              string `json:"email"`
	ProductName        string `json:"product_name"`
	PlanName           string `json:"plan_name"`
	PlanPrice		   string `json:"plan_price"`
	StartDate          string `json:"start_date"`
	EndDate            string `json:"end_date"`
}

type SubscriberRepository interface {
	GetSubscribersByOwner(ownerID uuid.UUID) (*[]Subscriber, error)
}

type SubscriberUseCase interface {
	GetSubscribersByOwner(ownerID uuid.UUID) (*[]Subscriber, error)
}
