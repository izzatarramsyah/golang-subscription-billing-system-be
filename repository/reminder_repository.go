package repository

import (
	"gorm.io/gorm"
	"subscription-billing-system/models"
	"github.com/google/uuid"
)

type reminderRepository struct {
	db *gorm.DB
}

func NewReminderRepository(db *gorm.DB) models.ReminderRepository {
	return &reminderRepository{
		db: db,
	}
}

func (r *reminderRepository) GetByID(id uuid.UUID) (*[]models.Reminder, error) {
	var reminder []models.Reminder
	if err := r.db.Where("id = ?", id).Find(&reminder).Error; err != nil {
		return nil, err
	}
	return &reminder, nil
}

func (r *reminderRepository) Create(rm *models.Reminder) error {
	return r.db.Create(rm).Error
}

func (r *reminderRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Reminder{}, id).Error
}

func (r *reminderRepository) Update(rm models.Reminder) error {
	return r.db.Save(rm).Error
}

func (r *reminderRepository) GetAll() (*[]models.Reminder, error) {
	var reminders []models.Reminder
	err := r.db.Find(&reminders).Error
	return &reminders, err
}
