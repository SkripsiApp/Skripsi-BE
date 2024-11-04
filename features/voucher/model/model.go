package model

import (
	"time"

	"gorm.io/gorm"
)

type Voucher struct {
	Id          string `gorm:"primaryKey;type:varchar(191);not null"`
	AdminId     string `gorm:"type:varchar(191);not null"`
	Name        string `gorm:"not null"`
	Description string 
	Discount    int    `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
