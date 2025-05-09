package models

import (
	"time"
	"github.com/google/uuid"
)

type Reminder struct {
    ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
    Type         string
	Title		 string
    Description  string	
	ReminderDate time.Time
    CreatedAt    time.Time
}

func (Reminder) TableName() string {
	return "subscription.reminders"
}

type ReminderRepository interface {
	GetAll()  (*[]Reminder, error)
	GetByID(id uuid.UUID) (*[]Reminder, error)
	Create(r *Reminder) error
	Delete(id uuid.UUID) error
	Update(r Reminder) error
}

type ReminderUseCase interface {
	GetAll()  (*[]Reminder, error)
	GetByID(id uuid.UUID) (*[]Reminder, error)
	Create(reminderType string, title string, description string, reminderDate string) error
	Delete(id uuid.UUID) error
	Update(r Reminder) error
}