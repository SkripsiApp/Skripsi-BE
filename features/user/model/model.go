package model

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id         string `gorm:"primaryKey;type:varchar(191);not null" json:"id"`
	Name       string `gorm:"not null"`
	Username   string `gorm:"unique;not null"`
	Email      string `gorm:"unique;not null"`
	Password   string `gorm:"not null"`
	Point      int    `gorm:"default:0"`
	Role       string `gorm:"default:user"`
	Otp        string
	OtpExpired string
	Address    []Address `gorm:"foreignKey:UserId;references:Id"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeleteAt   gorm.DeletedAt `gorm:"index"`
}
