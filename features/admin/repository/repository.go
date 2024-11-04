package repository

import (
	"skripsi/features/admin/entity"
	"skripsi/features/admin/interfaces"
	"skripsi/features/admin/mapping"

	"gorm.io/gorm"
)

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) interfaces.AdminRepositoryInterface {
	return &adminRepository{
		db: db,
	}
}

// Register implements interfaces.AdminRepositoryInterface.
func (a *adminRepository) Register(data entity.AdminCore) (entity.AdminCore, error) {
	request := mapping.AdminCoreToAdminModel(data)

	tx := a.db.Create(&request)
	if tx.Error != nil {
		return entity.AdminCore{}, tx.Error
	}

	dataResponse := mapping.AdminModelToAdminCore(request)
	return dataResponse, nil
}

// FindByEmail implements interfaces.AdminRepositoryInterface.
func (a *adminRepository) FindByEmail(email string) (entity.AdminCore, error) {
	dataAdmin := entity.AdminCore{}

	tx := a.db.Where("email = ?", email).First(&dataAdmin)
	if tx.RowsAffected == 0 {
		return entity.AdminCore{}, tx.Error
	}

	if tx.Error != nil {
		return entity.AdminCore{}, tx.Error
	}

	return dataAdmin, nil
}
