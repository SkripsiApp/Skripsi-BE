package service

import (
	"skripsi/features/dashboard/interfaces"
	"skripsi/utils/dashboard"
)

type dashboardService struct {
	dashboardRepository interfaces.DashboardRepositoryInterface
}

func NewDashboardService(dashboardRepo interfaces.DashboardRepositoryInterface) interfaces.DashboardServiceInterface {
	return &dashboardService{
		dashboardRepository: dashboardRepo,
	}
}

// DashboardData implements interfaces.DashboardServiceInterface.
func (d *dashboardService) DashboardData() (dashboard.Dashboard, error) {
	// Count total user
	totalUsers, err := d.dashboardRepository.CountTotalUser()
	if err != nil {
		return dashboard.Dashboard{}, err
	}

	// Count total transaction
	totalTransactions, err := d.dashboardRepository.CountTotalTransaction()
	if err != nil {
		return dashboard.Dashboard{}, err
	}

	// Count total product
	totalProducts, err := d.dashboardRepository.CountTotalProduct()
	if err != nil {
		return dashboard.Dashboard{}, err
	}

	// Count monthly transaction
	monthlyTransactions, err := d.dashboardRepository.CountMonthlyTransaction()
	if err != nil {
		return dashboard.Dashboard{}, err
	}

	// Count user ranking
	userRanking, err := d.dashboardRepository.CountUserRanking()
	if err != nil {
		return dashboard.Dashboard{}, err
	}

	// Count total revenue
	totalRevenue, err := d.dashboardRepository.CountTotalRevenue()
	if err != nil {
		return dashboard.Dashboard{}, err
	}

	// Map monthly transactions ke response format
    var monthlyStats []dashboard.MonthlyTransaction
    for _, t := range monthlyTransactions {
        monthlyStats = append(monthlyStats, dashboard.MonthlyTransaction{
            Month: t.Month,
            Total: t.Total,
        })
    }

	var userRankingResponse []dashboard.UserRanking
    for _, u := range userRanking {
        userRankingResponse = append(userRankingResponse, dashboard.UserRanking{
            Id:    u.Id,
            Name:  u.Name,
            Email: u.Email,
            Total: u.Total,
        })
    }


	response := dashboard.Dashboard{
		TotalUsers:          len(totalUsers),
		TotalTransactions:   len(totalTransactions),
		TotalProducts:       len(totalProducts),
		TotalRevenue:        totalRevenue,
		MonthlyTransactions: monthlyStats,
		UserRanking:         userRankingResponse,
	}

	return response, nil
}
