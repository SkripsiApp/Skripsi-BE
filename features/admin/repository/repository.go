package repository

import (
	"skripsi/features/admin/entity"
	"skripsi/features/admin/interfaces"

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

// Login implements interfaces.AdminRepositoryInterface.
func (a *adminRepository) Login(username string, password string) (entity.AdminCore, string, error) {
	panic("unimplemented")
}

// Register implements interfaces.AdminRepositoryInterface.
func (a *adminRepository) Register(data entity.AdminCore) (entity.AdminCore, error) {
	panic("unimplemented")
}