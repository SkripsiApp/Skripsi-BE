package model

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id         string         `gorm:"primaryKey;not null" json:"id"`
	Name       string         `gorm:"not null" json:"name"`
	Username   string         `gorm:"unique;not null" json:"username"`
	Email      string         `gorm:"unique;not null" json:"email"`
	Password   string         `gorm:"not null" json:"password"`
	Point      int            `gorm:"default:0" json:"point"`
	Role       string         `json:"role"`
	Otp        string         `json:"otp"`
	OtpExpired string         `json:"otp_expired"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeleteAt   gorm.DeletedAt `gorm:"index"`
}
