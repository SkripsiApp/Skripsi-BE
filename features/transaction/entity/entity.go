package entity

import "time"

type TransactionCore struct {
	Id                string
	UserId            string
	VoucherId         *string
	AddressId         string
	NoTransaction	 string
	OriginalPrice     int
	TotalPrice        int
	TotalPoint        int
	UsePoint          bool
	PointUsed         int
	VoucherDiscount   int
	DiscountAmount    int
	CourierName       string
	ShippingCost      int
	NoReceipt         string
	Status            string
	PaymentType       string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	TransactionDetail []TransactionDetailCore
}

type TransactionDetailCore struct {
	Id            string
	TransactionId string
	ProductId     string
	ProductName   string
	Image         string
	Size          string
	Quantity      int
	TotalPrice    int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
