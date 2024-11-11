package repository

import (
	"skripsi/features/transaction/entity"
	"skripsi/features/transaction/interfaces"
	"skripsi/features/transaction/mapping"
	"skripsi/features/transaction/model"
	"skripsi/utils/pagination"

	"gorm.io/gorm"
)

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) interfaces.TransactionRepositoryInterface {
	return &transactionRepository{
		db: db,
	}
}

// CreateTransaction implements interfaces.TransactionRepositoryInterface.
func (t *transactionRepository) CreateTransaction(data entity.TransactionCore) (entity.TransactionCore, error) {
	request := mapping.TransactionCoreToTransactionModel(data)

	tx := t.db.Create(&request)
	if tx.Error != nil {
		return entity.TransactionCore{}, tx.Error
	}

	response := mapping.TransactionModelToTransactionCore(request)
	return response, nil
}

// GetAllTransaction implements interfaces.TransactionRepositoryInterface.
func (t *transactionRepository) GetAllTransaction(search string, page int, limit int) ([]entity.TransactionCore, pagination.PageInfo, int, error) {
	data := []model.Transaction{}

	offset := (page - 1) * limit
	query := t.db.Model(&model.Transaction{}).Preload("TransactionDetail")

	if search != "" {
		query = query.Where("id LIKE ? or user_id LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	var totalCount int64
	tx := query.Count(&totalCount).Find(&data)
	if tx.Error != nil {
		return nil, pagination.PageInfo{}, 0, tx.Error
	}

	query = query.Offset(offset).Limit(limit)

	tx = query.Find(&data)
	if tx.Error != nil {
		return nil, pagination.PageInfo{}, 0, tx.Error
	}

	response := mapping.ListTransactionModelToListTransactionCore(data)

	pageInfo := pagination.CalculateData(int(totalCount), limit, page)
	return response, pageInfo, int(totalCount), nil
}

// GetTransactionById implements interfaces.TransactionRepositoryInterface.
func (t *transactionRepository) GetTransactionById(id string) (entity.TransactionCore, error) {
	data := model.Transaction{}

	tx := t.db.Preload("TransactionDetail").Where("id = ?", id).First(&data)
	if tx.Error != nil {
		return entity.TransactionCore{}, tx.Error
	}

	response := mapping.TransactionModelToTransactionCore(data)
	return response, nil
}

// UpdateStatusTransactionById implements interfaces.TransactionRepositoryInterface.
func (t *transactionRepository) UpdateStatusTransactionById(id string, status string) error {
	data := model.Transaction{}

	tx := t.db.Model(&data).Where("id = ?", id).Update("status", status)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// GetAllTransactionByUserId implements interfaces.TransactionRepositoryInterface.
func (t *transactionRepository) GetAllTransactionByUserId(userId string, search string, page int, limit int) ([]entity.TransactionCore, pagination.PageInfo, int, error) {
	data := []model.Transaction{}

	offset := (page - 1) * limit
	query := t.db.Model(&model.Transaction{}).Preload("TransactionDetail").Where("user_id = ?", userId)

	if search != "" {
		query = query.Where("id LIKE ? or user_id LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	var totalCount int64
	tx := query.Count(&totalCount).Find(&data)
	if tx.Error != nil {
		return nil, pagination.PageInfo{}, 0, tx.Error
	}

	query = query.Offset(offset).Limit(limit)

	tx = query.Find(&data)
	if tx.Error != nil {
		return nil, pagination.PageInfo{}, 0, tx.Error
	}

	response := mapping.ListTransactionModelToListTransactionCore(data)

	pageInfo := pagination.CalculateData(int(totalCount), limit, page)
	return response, pageInfo, int(totalCount), nil
}

// UpdateNoResiTransactionById implements interfaces.TransactionRepositoryInterface.
func (t *transactionRepository) UpdateNoReceiptTransactionById(id, resi string) error {
	data := model.Transaction{}

	tx := t.db.Model(&data).Where("id = ?", id).Update("no_resi", resi)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

