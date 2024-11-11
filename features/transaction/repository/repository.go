package repository

import (
	"skripsi/features/transaction/entity"
	"skripsi/features/transaction/interfaces"
	"skripsi/features/transaction/mapping"
	"skripsi/features/transaction/model"

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
func (t *transactionRepository) GetAllTransaction(search string, page int, limit int) ([]entity.TransactionCore, int, error) {
	panic("unimplemented")
}

// GetTransactionById implements interfaces.TransactionRepositoryInterface.
func (t *transactionRepository) GetTransactionById(id string) (entity.TransactionCore, error) {
	data := model.Transaction{}

	tx := t.db.Where("id = ?", id).First(&data)
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
