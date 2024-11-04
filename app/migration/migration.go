package migration

import (
	address "skripsi/features/address/model"
	admin "skripsi/features/admin/model"
	user "skripsi/features/user/model"

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
}
