package models

import (
	"github.com/google/uuid"
    "time"
)

type RevenueReport struct {
	TotalRevenue   *float64  `json:"total_revenue"`
	Subscriptions  []Subscriptions `json:"subscriptions"`
}

type Subscriptions struct {
	SubscriberName     string    `json:"subscriber_name"` 
	Email              string    `json:"email"`         
	ProductName        string    `json:"product_name"`    
	PlanName           string    `json:"plan_name"`      
	PlanPrice          float64   `json:"plan_price"`      
	StartDate          time.Time `json:"start_date"`      
	EndDate            time.Time `json:"end_date"`        
	SubscriptionStatus string    `json:"subscription_status"`
}

type ReportRepository interface {
	GetRevenueReport(userID uuid.UUID, startDate, endDate time.Time) (*RevenueReport, error)
}

type ReportUseCase interface {
	GetRevenueReport(userID uuid.UUID, startDate, endDate time.Time) (*RevenueReport, error)
}