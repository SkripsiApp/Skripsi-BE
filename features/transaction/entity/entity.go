package entity

import "time"

type TransactionCore struct {
	Id             string
	UserId         string
	VoucherId      string
	AddressId      string
	TotalPrice     int
	TotalPoint     int
	UsePoint       bool
	PointUsed      int
	DiscountAmount int
	CourierName    string
	NoReceipt      string
	Status         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	TransactionDetail []TransactionDetailCore
}

type TransactionDetailCore struct {
	Id            string
	TransactionId string
	ProductId     string
	Quantity      int
	TotalPrice    int
	TotalPoint    int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
