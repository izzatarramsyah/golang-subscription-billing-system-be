package models

import (
	"time"
	"github.com/google/uuid"
)

type Plan struct {
    ID             uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
    ProductID      uuid.UUID `gorm:"type:uuid"`
	Product        Product   `gorm:"foreignKey:ProductID;references:ID"` 
    Name           string
    Price          float64
    DurationMonths int // 1 (monthly), 12 (yearly), etc
    CreatedAt      time.Time
}

func (Plan) TableName() string {
	return "subscription.plans"
}   

type PlanRepository interface {
	Create(plan *Plan) error
	GetAll() (*[]Plan, error)
	GetByID(id uuid.UUID) (*Plan, error)
	GetByProductID(id uuid.UUID) (*[]Plan, error)
	Update(plan *Plan) error
	Delete(id uuid.UUID) error
}

type PlanUseCase interface {
	CreatePlan(plan *Plan) error
	GetAllPlans() (*[]Plan, error)
	GetPlanByID(id uuid.UUID) (*Plan, error)
	GetByProductID(id uuid.UUID) (*[]Plan, error)
	UpdatePlan(plan *Plan) error
	DeletePlan(id uuid.UUID) error
}