package usecase

import (
	"github.com/google/uuid"
	"subscription-billing-system/models"
	"time"
	"fmt"
)

type reminderUseCase struct {
	repo models.ReminderRepository
}

func NewReminderUseCase(repo models.ReminderRepository) models.ReminderUseCase {
	return &reminderUseCase{repo}
}

func (uc *reminderUseCase) GetByID(id uuid.UUID) (*[]models.Reminder, error) {
	return uc.repo.GetByID(id)
}

func (uc *reminderUseCase) Create(reminderType string, title string, description string, reminderDate string,) error {
	// Membuat instance Reminder sebagai pointer
	rm := models.Reminder{
		Type:         reminderType,
		Title:        title,
		Description:  description,
		CreatedAt:    time.Now(), // Set current time sebagai CreatedAt
	}

	// Parsing ReminderDate (jika ada)
	if reminderDate != "" {
		parsedReminderDate, err := time.Parse("2006-01-02", reminderDate)
		if err != nil {
			return fmt.Errorf("invalid ReminderDate format: %v", err)
		}
		rm.ReminderDate = parsedReminderDate
	} else {
		// Jika ReminderDate kosong, set ke current time sebagai fallback
		rm.ReminderDate = time.Now()
	}

	// Simpan reminder ke repository dengan pointer rm
	return uc.repo.Create(&rm)
}


func (uc *reminderUseCase) Delete(id uuid.UUID) error {
	return uc.repo.Delete(id)
}

func (uc *reminderUseCase) Update(rm models.Reminder) error {
	return uc.repo.Update(rm)
}

func (uc *reminderUseCase) GetAll() (*[]models.Reminder, error) {
	return uc.repo.GetAll()
}
