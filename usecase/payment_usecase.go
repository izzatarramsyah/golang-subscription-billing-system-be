package usecase

import (
	"time"
	"github.com/google/uuid"

	"subscription-billing-system/models"
)

type paymentUseCase struct {
	paymentRepository models.PaymentRepository
}

func NewPaymentUseCase(r models.PaymentRepository) models.PaymentUseCase {
	return &paymentUseCase{r}
}

func (u *paymentUseCase) ProcessPayment(userID uuid.UUID, subscriptionID uuid.UUID, amount float64, paymentMethod string) (*models.Payment, error) {
	payment := &models.Payment{
		UserID:         userID,
		SubscriptionID: subscriptionID,
		Amount:         amount,
		Gateway:        paymentMethod,
		Status:         "pending",
		PaidAt:         time.Now(),
		CreatedAt:      time.Now(),
	}

	err := u.paymentRepository.Create(payment)
	return payment, err
}

func (u *paymentUseCase) GetPayments(userID uuid.UUID) (*[]models.Payment, error) {
	return u.paymentRepository.GetByUser(userID)
}

func (u *paymentUseCase) GetAllPaymentDetails() ([]models.PaymentDetail, error) {
	return u.paymentRepository.GetAllPaymentDetails()
}

func (u *paymentUseCase) UpdatePaymentStatus(paymentID uuid.UUID, status string) (*bool, error) {
	return u.paymentRepository.UpdatePaymentStatus(paymentID, status)
}