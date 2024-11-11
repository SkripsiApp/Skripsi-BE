package interfaces

import (
	"skripsi/features/transaction/entity"
	"skripsi/utils/helper"
)

type TransactionRepositoryInterface interface {
	CreateTransaction(data entity.TransactionCore) (entity.TransactionCore, error)
	GetAllTransaction(search string, page, limit int) ([]entity.TransactionCore, int, error)
	GetTransactionById(id string) (entity.TransactionCore, error)
	UpdateStatusTransactionById(id, status string) error
}

type TransactionServiceInterface interface {
	CreateTransaction(data entity.TransactionCore) (entity.TransactionCore, error)
	GetAllTransaction(search string, page, limit int) ([]entity.TransactionCore, int, error)
	GetTransactionById(id string) (entity.TransactionCore, error)
	UpdateStatusTransactionById(id, status string) error
	HandleMidtransNotification(notification helper.MidtransNotificationPayload) error
}
