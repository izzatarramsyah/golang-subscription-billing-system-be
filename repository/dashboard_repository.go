package repository

import (
	"subscription-billing-system/models"
	"time"

	"gorm.io/gorm"
)

type dashboardRepository struct {
	db *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) models.DashboardRepository {
	return &dashboardRepository{db}
}

func (r *dashboardRepository) CountActiveSubscribers() (int64, error) {
	var count int64
	err := r.db.Table("subscription.subscriptions").Where("status = ?", "active").Count(&count).Error
	return count, err
}

func (r *dashboardRepository) CountAllOrders() (int64, error) {
	var count int64
	err := r.db.Table("subscription.subscriptions").Count(&count).Error
	return count, err
}

func (r *dashboardRepository) CountMonthlySales(startOfMonth time.Time) (int64, error) {
	var count int64
	err := r.db.Table("subscription.subscriptions").
		Where("created_at >= ?", startOfMonth).
		Count(&count).Error
	return count, err
}

func (r *dashboardRepository) GetTodayRevenue(date time.Time) (float64, error) {
	var revenue float64
	err := r.db.Table("subscription.subscriptions s").
		Select("COALESCE(SUM(p.price), 0)").
		Joins("JOIN subscription.plans p ON p.id = s.plan_id").
		Where("DATE(s.created_at) = DATE(?) AND s.status = ?", date, "active").
		Scan(&revenue).Error
	return revenue, err
}

func (r *dashboardRepository) GetRevenuePerMonth(month time.Month, year int) (float64, error) {
	var revenue float64
	err := r.db.Table("subscription.subscriptions s").
		Select("COALESCE(SUM(p.price), 0)").
		Joins("JOIN subscription.plans p ON p.id = s.plan_id").
		Where("EXTRACT(MONTH FROM s.created_at) = ? AND EXTRACT(YEAR FROM s.created_at) = ? AND s.status = ?", month, year, "active").
		Scan(&revenue).Error
	return revenue, err
}

func (r *dashboardRepository) GetRecentOrders(limit int) ([]models.OrderSummary, error) {
	var orders []models.OrderSummary
	err := r.db.Table("subscription.subscriptions s").
		Select("s.id as order_id, u.username as subscriber, pr.name as product_name, p.name as plan_name, p.price as plan_price, s.created_at, s.status").
		Joins("JOIN subscription.users u ON u.id = s.user_id").
		Joins("JOIN subscription.plans p ON p.id = s.plan_id").
		Joins("JOIN subscription.products pr ON pr.id = p.product_id").
		Order("s.created_at DESC").
		Limit(limit).
		Scan(&orders).Error
	return orders, err
}
