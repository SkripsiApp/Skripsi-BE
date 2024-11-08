package entity

import (
	"time"

	"gorm.io/gorm"
)

type ProductCore struct {
	Id          string
	Name        string
	Description string
	Price       int
	Category    string
	Sold        int
	Image       string
	ProductSize []ProductSizeCore
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

type ProductSizeCore struct {
	Id        string
	ProductId string
	Size      string
	Stock     int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
