package models

import "time"

type Dashboard struct {
	TotalSubscribers int                `json:"total_subscribers"`
	TotalOrders      int                `json:"total_orders"`
	MonthlySales     int                `json:"monthly_sales"`
	MonthlyTargetPercentage float64     `json:"monthly_percentage"`
	TargetRevenue    float64            `json:"target_revenue"`
	RevenuePerMonth  map[string]float64 `json:"revenue_per_month"`
	TodayRevenue     float64            `json:"today_revenue"`
	MonthlyRevenue   float64            `json:"monthly_revenue"`
	RecentOrders     []OrderSummary     `json:"recent_orders"`
}

type OrderSummary struct {
	OrderID     string    `json:"order_id"`
	Subscriber  string    `json:"subscriber"`
	ProductName string    `json:"product_name"`
	PlanName    string    `json:"plan_name"`
	PlanPrice   float64   `json:"plan_price"`
	CreatedAt   time.Time `json:"created_at"`
	Status      string    `json:"status"`
}

type DashboardRepository interface {
	CountActiveSubscribers() (int64, error)
	CountAllOrders() (int64, error)
	CountMonthlySales(startOfMonth time.Time) (int64, error)
	GetTodayRevenue(today time.Time) (float64, error)
	GetRevenuePerMonth(month time.Month, year int) (float64, error)
	GetRecentOrders(limit int) ([]OrderSummary, error)
}

type DashboardUseCase interface {
	GetDashboardData() (Dashboard, error)
}
