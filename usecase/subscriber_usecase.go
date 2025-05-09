package usecase

import (
	"github.com/google/uuid"

	"subscription-billing-system/models"
)

type subscriberUseCase struct {
	repo models.SubscriberRepository
}

func NewSubscriberUseCase(repo models.SubscriberRepository) models.SubscriberUseCase {
	return &subscriberUseCase{repo}
}

func (uc *subscriberUseCase) GetSubscribersByOwner(ownerID uuid.UUID) (*[]models.Subscriber, error) {
	return uc.repo.GetSubscribersByOwner(ownerID)
}
