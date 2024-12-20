package model

import (
	address "skripsi/features/address/model"
	product "skripsi/features/product/model"
	voucher "skripsi/features/voucher/model"
	"time"
)

type Transaction struct {
	Id                string  `gorm:"primaryKey;type:varchar(191);not null"`
	UserId            string  `gorm:"type:varchar(191);not null"`
	VoucherId         *string `gorm:"type:varchar(191);null"`
	AddressId         string  `gorm:"type:varchar(191);not null"`
	NoTransaction     string  `gorm:"type:varchar(191);not null;unique"`
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
	TransactionDetail []TransactionDetail `gorm:"foreignKey:TransactionId;references:Id;constraint:OnDelete:CASCADE"`
	Voucher           voucher.Voucher     `gorm:"foreignKey:VoucherId;references:Id"`
	Address           address.Address     `gorm:"foreignKey:AddressId;references:Id"`
}

type TransactionDetail struct {
	Id            string `gorm:"primaryKey;type:varchar(191);not null"`
	TransactionId string `gorm:"type:varchar(191);not null"`
	ProductId     string `gorm:"type:varchar(191);not null"`
	Size          string `gorm:"type:varchar(10);not null"`
	Quantity      int
	TotalPrice    int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Product       product.Product `gorm:"foreignKey:ProductId;references:Id"`
}
