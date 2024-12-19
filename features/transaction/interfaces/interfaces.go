package interfaces

import (
	"skripsi/features/transaction/entity"
	"skripsi/utils/helper"
	"skripsi/utils/pagination"
)

type TransactionRepositoryInterface interface {
	CreateTransaction(data entity.TransactionCore) (entity.TransactionCore, error)
	GetAllTransaction(search string, page, limit int) ([]entity.TransactionCore, pagination.PageInfo, int, error)
	GetTransactionById(id string) (entity.TransactionCore, error)
	UpdateStatusTransactionById(id, status string) error
	GetAllTransactionByUserId(userId string, search string, page, limit int) ([]entity.TransactionCore, pagination.PageInfo, int, error)
	UpdateNoReceiptTransactionById(id, resi string) error
	UpdateStatusTransactionUserById(userId, transactionId, status string) error
	UpdatePaymentDetails(transactionId, paymentType, status string) error
}

type TransactionServiceInterface interface {
	CreateTransaction(data entity.TransactionCore) (entity.TransactionCore, string, error)
	GetAllTransaction(search string, page, limit int) ([]entity.TransactionCore, pagination.PageInfo, int, error)
	GetTransactionById(id string) (entity.TransactionCore, error)
	UpdateStatusTransactionUserById(userId, transactionId, status string) error
	HandleMidtransNotification(notification helper.MidtransNotificationPayload) error
	GetAllTransactionByUserId(userId string, search string, page, limit int) ([]entity.TransactionCore, pagination.PageInfo, int, error)
	UpdateNoReceiptTransactionById(id, resi string) error
}
