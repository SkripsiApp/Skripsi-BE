package migration

import (
	address "skripsi/features/address/model"
	admin "skripsi/features/admin/model"
	product "skripsi/features/product/model"
	user "skripsi/features/user/model"
	voucher "skripsi/features/voucher/model"

	"gorm.io/gorm"
)

func InitMigrationMysql(db *gorm.DB) {
	db.AutoMigrate(&user.Users{})
	db.AutoMigrate(&address.Address{})
}

func InitMigrationPostgre(db *gorm.DB) {
	db.AutoMigrate(&user.Users{})
	db.AutoMigrate(&address.Address{})
	db.AutoMigrate(&admin.Admin{})
	db.AutoMigrate(&voucher.Voucher{})
	db.AutoMigrate(&product.Product{}, &product.ProductSize{})
}
