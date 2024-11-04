package model

import "skripsi/features/voucher/model"

type Admin struct {
	Id       string          `gorm:"primaryKey;type:varchar(191);not null" json:"id"`
	Email    string          `gorm:"unique;not null"`
	Password string          `json:"password"`
	Role     string          `gorm:"default:admin"`
	Voucher  []model.Voucher `gorm:"foreignKey:AdminId;references:Id"`
}
