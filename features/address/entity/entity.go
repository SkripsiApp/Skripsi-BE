package entity

import (
	"time"

	"gorm.io/gorm"
)

type AddressCore struct {
	Id         string
	UserId     string
	Name       string
	Address    string
	City       string
	Subdistric string
	ZipCode    string
	Phone      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt   gorm.DeletedAt
}
