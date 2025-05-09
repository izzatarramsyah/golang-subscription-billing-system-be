package usecase

import (
	"subscription-billing-system/models"
	"time"
)

type dashboardUseCase struct {
	repo models.DashboardRepository
}

func NewDashboardUseCase(repo models.DashboardRepository) models.DashboardUseCase {
	return &dashboardUseCase{repo}
}

func (u *dashboardUseCase) GetDashboardData() (models.Dashboard, error) {
	now := time.Now()
	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	today := now.Truncate(24 * time.Hour)

	subscribers, err := u.repo.CountActiveSubscribers()
	if err != nil {
		return models.Dashboard{}, err
	}

	orders, err := u.repo.CountAllOrders()
	if err != nil {
		return models.Dashboard{}, err
	}

	monthlySales, err := u.repo.CountMonthlySales(startOfMonth)
	if err != nil {
		return models.Dashboard{}, err
	}

	currentMonth := now.Month()
	currentYear := now.Year()

	monthlyRevenue, err := u.repo.GetRevenuePerMonth(currentMonth, currentYear)
	if err != nil {
		return models.Dashboard{}, err
	}

	todayRevenue, err := u.repo.GetTodayRevenue(today)
	if err != nil {
		return models.Dashboard{}, err
	}

	revenueByMonth := make(map[string]float64)
	for m := time.January; m <= time.December; m++ {
		label := time.Date(currentYear, m, 1, 0, 0, 0, 0, time.UTC).Format("Jan 2006")
		revenue, err := u.repo.GetRevenuePerMonth(m, currentYear)
		if err != nil {
			return models.Dashboard{}, err
		}
		revenueByMonth[label] = revenue
	}

	recentOrders, err := u.repo.GetRecentOrders(5)
	if err != nil {
		return models.Dashboard{}, err
	}

	targetRevenue := 1000000.0
	var percentage float64
	if targetRevenue > 0 {
		percentage = (float64(monthlyRevenue) / targetRevenue) * 100
	}

	return models.Dashboard{
		TotalSubscribers: int(subscribers),
		TotalOrders:      int(orders),
		MonthlySales:     int(monthlySales),
		MonthlyRevenue:   monthlyRevenue,
		TargetRevenue:    1000000.0,
		RevenuePerMonth:  revenueByMonth,
		TodayRevenue:     todayRevenue,
		RecentOrders:     recentOrders,
		MonthlyTargetPercentage: percentage, 
	}, nil
}
