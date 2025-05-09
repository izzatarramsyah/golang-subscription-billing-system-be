package usecase

import (
	"github.com/google/uuid"
	"subscription-billing-system/models"
)

type libraryItemUseCase struct {
	repo models.LibraryItemRepository
}

func NewLibraryItemUseCase(repo models.LibraryItemRepository) models.LibraryItemUseCase {
	return &libraryItemUseCase{repo}
}

func (uc *libraryItemUseCase) GetUserLibrary(userID uuid.UUID) ([]models.LibraryItem, error) {
	return uc.repo.GetUserLibrary(userID)
}
