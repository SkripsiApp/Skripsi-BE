package model

import (
	"time"

	"gorm.io/gorm"
)

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
