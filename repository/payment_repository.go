package repository

import (
	"subscription-billing-system/models"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) models.PaymentRepository {
	return &paymentRepository{db}
}

func (r *paymentRepository) Create(payment *models.Payment) error {
	return r.db.Create(payment).Error
}

func (r *paymentRepository) GetByUser(userID uuid.UUID) (*[]models.Payment, error) {
	var payments []models.Payment
	err := r.db.Where("user_id = ?", userID).Find(&payments).Error
	return &payments, err
}

func (r *paymentRepository) GetAllPaymentDetails() ([]models.PaymentDetail, error) {
    var details []models.PaymentDetail

    err := r.db.
        Table("subscription.payments").
        Select(`
            payments.id as payment_id,
            payments.gateway,
            payments.amount,
            payments.status,
            payments.paid_at,
            users.id as user_id,
            users.username,
            users.email,
            subscriptions.id as subscription_id,
            subscriptions.status as subscription_status,
            subscriptions.start_date,
            subscriptions.end_date
        `).
        Joins("JOIN subscription.subscriptions ON payments.subscription_id = subscriptions.id").
        Joins("JOIN subscription.users ON payments.user_id = users.id").
        Scan(&details).Error

    if err != nil {
        return nil, err
    }

    return details, nil
}

func (r *paymentRepository) UpdatePaymentStatus(userID uuid.UUID, status string) (*bool, error) {

	err := r.db.Model(&models.Payment{}).Where("id = ?", userID).Update("status", status).Error
	if err != nil {
		result := false
		return &result, nil
	}

	result := true
	return &result, nil
}