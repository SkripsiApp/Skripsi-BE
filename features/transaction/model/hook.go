package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (transaction *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	transaction.Id = newUuid.String()

	return nil
}

func (transactionDetail *TransactionDetail) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	transactionDetail.Id = newUuid.String()

	return nil
}