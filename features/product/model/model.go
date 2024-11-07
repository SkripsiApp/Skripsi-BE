package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	Id          string        `gorm:"primaryKey;type:varchar(191);not null"`
	Name        string        `gorm:"not null"`
	Description string        `gorm:"type:text"`
	Price       int           `gorm:"not null"`
	Category    string        `gorm:"not null"`
	Sold        int           `gorm:"default:0"`
	Image       string        `gorm:"not null"`
	ProductSize []ProductSize `gorm:"foreignKey:ProductId;references:Id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type ProductSize struct {
	Id        string `gorm:"primaryKey;type:varchar(191);not null"`
	ProductId string `gorm:"type:varchar(191);not null"`
	Size      string `gorm:"type:varchar(10);not null"`
	Stock     int    `gorm:"not null;default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
