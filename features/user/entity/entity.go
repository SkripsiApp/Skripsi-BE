package entity

import (
	"time"

	"gorm.io/gorm"
)

type UsersCore struct {
	Id              string
	Name            string
	Username        string
	Email           string
	Password        string
	ConfirmPassword string
	Point           int
	Role            string
	Otp             string
	OtpExpired      int64
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}
