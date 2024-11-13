package interfaces

import (
	product "skripsi/features/product/entity"
	transaction "skripsi/features/transaction/entity"
	user "skripsi/features/user/entity"
	"skripsi/utils/dashboard"
)

type DashboardRepositoryInterface interface {
	CountTotalUser() ([]user.UsersCore, error)
	CountTotalTransaction() ([]transaction.TransactionCore, error)
	CountTotalProduct() ([]product.ProductCore, error)
	CountMonthlyTransaction() ([]dashboard.MonthlyTransaction, error)
	CountUserRanking() ([]dashboard.UserRanking, error)
	CountTotalRevenue() (int, error)
}

type DashboardServiceInterface interface {
	DashboardData() (dashboard.Dashboard, error)
}
