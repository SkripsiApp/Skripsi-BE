package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (product *Product) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	product.Id = newUuid.String()

	return nil
}

func (product *ProductSize) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	product.Id = newUuid.String()

	return nil
}

