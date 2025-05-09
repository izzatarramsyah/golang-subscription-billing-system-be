package repository

import (
	"github.com/google/uuid"
	"time"
	"gorm.io/gorm"
	"subscription-billing-system/models"
)

type reportRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) models.ReportRepository {
	return &reportRepository{db}
}

func (r *reportRepository) GetRevenueReport(userID uuid.UUID, startDate, endDate time.Time) (*models.RevenueReport, error) {
	var report models.RevenueReport

	// Mengambil total pendapatan
	err := r.db.Table("subscription.subscriptions").
		Select("SUM(subscription.plans.price) AS total_revenue").
		Joins("JOIN subscription.plans ON subscription.plans.id = subscription.subscriptions.plan_id").
		Joins("JOIN subscription.products ON subscription.products.id = subscription.plans.product_id").
		Where("subscription.products.owner_id = ?", userID).
		// Where("subscription.subscriptions.start_date >= ? AND subscription.subscriptions.end_date <= ?", startDate, endDate).
		Scan(&report.TotalRevenue).Error

	// Cek apakah terjadi error atau hasil NULL
	if err != nil {
		return nil, err
	}

	// Jika TotalRevenue masih nil (NULL), inisialisasi ke 0
	if report.TotalRevenue == nil {
		defaultRevenue := 0.0
		report.TotalRevenue = &defaultRevenue
	}

	// Mengambil detail langganan
	err = r.db.Table("subscription.subscriptions").
		Select("users.username AS subscriber_name, users.email, subscription.products.name AS product_name, subscription.plans.name AS plan_name, subscription.plans.price AS plan_price, subscription.subscriptions.start_date, subscription.subscriptions.end_date, subscription.subscriptions.status AS subscription_status").
		Joins("JOIN subscription.plans ON subscription.plans.id = subscription.subscriptions.plan_id").
		Joins("JOIN subscription.products ON subscription.products.id = subscription.plans.product_id").
		Joins("JOIN subscription.users ON subscription.users.id = subscription.subscriptions.user_id"). // Menambahkan join ke tabel users
		Where("subscription.products.owner_id = ?", userID).
		// Where("subscription.subscriptions.start_date >= ? AND subscription.subscriptions.end_date <= ?", startDate, endDate).
		Scan(&report.Subscriptions).Error

	if err != nil {
		return nil, err
	}

	return &report, nil
}
