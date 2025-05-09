package usecase

import (
	"subscription-billing-system/models"
	"time"
	"github.com/google/uuid"
)

type planUseCase struct {
    planRepo   models.PlanRepository
}

func NewPlanUseCase(planRepo models.PlanRepository) models.PlanUseCase {
    return &planUseCase{planRepo}
}

func (uc *planUseCase) CreatePlan(plan *models.Plan) error {
	plan.CreatedAt = time.Now() // Set waktu saat ini
	return uc.planRepo.Create(plan)
}

func (uc *planUseCase) GetAllPlans() (*[]models.Plan, error) {
	return uc.planRepo.GetAll()
}

func (uc *planUseCase) GetPlanByID(id uuid.UUID) (*models.Plan, error) {
	return uc.planRepo.GetByID(id)
}

func (uc *planUseCase) UpdatePlan(plan *models.Plan) error {
	plan.CreatedAt = time.Now() 
	return uc.planRepo.Update(plan)
}

func (uc *planUseCase) DeletePlan(id uuid.UUID) error {
	return uc.planRepo.Delete(id)
}

func (uc *planUseCase) GetByProductID(productId uuid.UUID) (*[]models.Plan, error) {
	return uc.planRepo.GetByProductID(productId)
}