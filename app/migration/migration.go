package migration

import (
	"skripsi/features/user/model"

	"gorm.io/gorm"
)

func InitMigrationMysql(db *gorm.DB) {
	db.AutoMigrate(&model.Users{})
}
