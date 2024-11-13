package repository

import (
	"skripsi/features/dashboard/interfaces"
	product "skripsi/features/product/entity"
	mappingProduct "skripsi/features/product/mapping"
	modelProduct "skripsi/features/product/model"
	transaction "skripsi/features/transaction/entity"
	mappingTransaction "skripsi/features/transaction/mapping"
	modelTransaction "skripsi/features/transaction/model"
	user "skripsi/features/user/entity"
	mappingUser "skripsi/features/user/mapping"
	modelUser "skripsi/features/user/model"
	"skripsi/utils/dashboard"
	"time"

	"gorm.io/gorm"
)

type dashboardRepository struct {
	db *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) interfaces.DashboardRepositoryInterface {
	return &dashboardRepository{
		db: db,
	}
}

// CountMonthlyTransaction implements interfaces.DashboardRepositoryInterface.
func (d *dashboardRepository) CountMonthlyTransaction() ([]dashboard.MonthlyTransaction, error) {
	months := []string{
		"Januari", "Februari", "Maret", "April", "Mei", "Juni",
		"Juli", "Agustus", "September", "Oktober", "November", "Desember",
	}

	data := []struct {
		Month int `json:"month"`
		Total int `json:"total"`
	}{}

	err := d.db.
		Table("transactions").
		Select("EXTRACT(MONTH FROM created_at) AS month, COUNT(id) AS total").
		Where("EXTRACT(YEAR FROM created_at) = ?", time.Now().Year()).
		Group("month").
		Order("month").
		Scan(&data).Error
	if err != nil {
		return nil, err
	}

	// Simpan transaksi per bulan ke map
	monthlyMap := make(map[int]int)
	for _, item := range data {
		monthlyMap[item.Month] = item.Total
	}

	// Tampilkan semua bulan, jika tidak ada transaksi maka total = 0
	var result []dashboard.MonthlyTransaction
	for i := 1; i <= 12; i++ {
		result = append(result, dashboard.MonthlyTransaction{
			Month: months[i-1],
			Total: monthlyMap[i],
		})
	}

	return result, nil
}

// CountTotalProduct implements interfaces.DashboardRepositoryInterface.
func (d *dashboardRepository) CountTotalProduct() ([]product.ProductCore, error) {
	model := []modelProduct.Product{}

	err := d.db.Table("products").Find(&model).Error

	if err != nil {
		return nil, err
	}

	results := mappingProduct.ListProductModelToProductCore(model)
	return results, nil
}

// CountTotalTransaction implements interfaces.DashboardRepositoryInterface.
func (d *dashboardRepository) CountTotalTransaction() ([]transaction.TransactionCore, error) {
	model := []modelTransaction.Transaction{}

	err := d.db.Table("transactions").Find(&model).Error

	if err != nil {
		return nil, err
	}

	results := mappingTransaction.ListTransactionModelToListTransactionCore(model)
	return results, nil
}

// CountTotalUser implements interfaces.DashboardRepositoryInterface.
func (d *dashboardRepository) CountTotalUser() ([]user.UsersCore, error) {
	model := []modelUser.Users{}

	err := d.db.Table("users").Find(&model).Error

	if err != nil {
		return nil, err
	}

	results := mappingUser.ListUserModelToUserCore(model)
	return results, nil
}

// CountUserRanking implements interfaces.DashboardRepositoryInterface.
func (d *dashboardRepository) CountUserRanking() ([]dashboard.UserRanking, error) {
	data := []dashboard.UserRanking{}

	err := d.db.
		Table("users").
		Select("users.id, users.name, users.email, COUNT(transactions.id) as total").
		Joins("LEFT JOIN transactions ON users.id = transactions.user_id").
		Group("users.id, users.name, users.email").
		Order("total DESC").
		Scan(&data).Error

	if err != nil {
		return nil, err
	}
	return data, nil
}

// CountTotalRevenue implements interfaces.DashboardRepositoryInterface.
func (d *dashboardRepository) CountTotalRevenue() (int, error) {
	var totalRevenue int

	err := d.db.Table("transactions").
		Where("status = ?", "Paid").
		Select("SUM(total_price)").
		Scan(&totalRevenue).Error
	if err != nil {
		return 0, err
	}

	return totalRevenue, nil
}
