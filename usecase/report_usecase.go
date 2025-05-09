package usecase

import (
	"github.com/google/uuid"
	"time"
	"subscription-billing-system/models"
)

type reportUseCase struct {
	repo models.ReportRepository
}

func NewReportUseCase(repo models.ReportRepository) models.ReportUseCase {
	return &reportUseCase{repo}
}

func (uc *reportUseCase) GetRevenueReport(userID uuid.UUID, startDate, endDate time.Time) (*models.RevenueReport, error) {
	return uc.repo.GetRevenueReport(userID, startDate, endDate)
}
