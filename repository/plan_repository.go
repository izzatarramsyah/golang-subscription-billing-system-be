package repository

import (
	"gorm.io/gorm"
	"subscription-billing-system/models"
	"log"
	"github.com/google/uuid"
)

type planRepository struct {
	db *gorm.DB
}

func NewPlanRepository(db *gorm.DB) models.PlanRepository {
	return &planRepository{
		db: db,
	}
}

func (r *planRepository) Create(plan *models.Plan) error {
	err := r.db.Create(plan).Error
	if err != nil {
		log.Printf("Failed to create plan: %v\n", err)
	} else {
		log.Printf("Successfully created plan with ID: %d\n", plan.ID)
	}
	return err
}

func (r *planRepository) GetAll() (*[]models.Plan, error) {
	var plans []models.Plan
	if err := r.db.Preload("Product").Find(&plans).Error; err != nil {
        return nil, err
    }
    return &plans, nil
}

func (r *planRepository) GetByID(id uuid.UUID) (*models.Plan, error) {
	var plan models.Plan
	err := r.db.First(&plan, id).Error
	if err != nil {
		log.Printf("Failed to get plan with ID %d: %v\n", id, err)
	} else {
		log.Printf("Successfully retrieved plan with ID: %d\n", id)
	}
	return &plan, err
}

func (r *planRepository) Update(plan *models.Plan) error {
	err := r.db.Save(plan).Error
	if err != nil {
		log.Printf("Failed to update plan with ID %d: %v\n", plan.ID, err)
	} else {
		log.Printf("Successfully updated plan with ID: %d\n", plan.ID)
	}
	return err
}

func (r *planRepository) Delete(id uuid.UUID) error {
	err := r.db.Delete(&models.Plan{}, id).Error
	if err != nil {
		log.Printf("Failed to delete plan with ID %d: %v\n", id, err)
	} else {
		log.Printf("Successfully deleted plan with ID: %d\n", id)
	}
	return err
}

func (r *planRepository) GetByProductID(productId uuid.UUID) (*[]models.Plan, error){
	log.Println("Get Plan:")
	var plan []models.Plan
	if err := r.db.Where("product_id = ?", productId).Find(&plan).Error; err != nil {
		log.Printf("Error get plan: %v\n", err)
		return nil, err
	}
	log.Println("get plan successfully")
	return &plan, nil
}