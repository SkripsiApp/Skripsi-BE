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

type Address struct {
	Id         string `gorm:"primaryKey;not null" json:"id"`
	UserId     string `gorm:"type:varchar(191);not null"`
	Name       string `gorm:"not null"`
	Address    string `gorm:"not null"`
	City       string `gorm:"not null"`
	Subdistric string `gorm:"not null"`
	ZipCode    string `gorm:"not null"`
	Phone      string `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeleteAt   gorm.DeletedAt `gorm:"index"`
}
