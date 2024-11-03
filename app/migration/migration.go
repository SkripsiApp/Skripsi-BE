package migration

import (
	user "skripsi/features/user/model"

	"gorm.io/gorm"
)

func InitMigrationMysql(db *gorm.DB) {
	db.AutoMigrate(&user.Users{})
	db.AutoMigrate(&user.Address{})
}
