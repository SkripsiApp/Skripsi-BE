package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (admin *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	admin.Id = newUuid.String()

	return nil
}
