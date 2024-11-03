package repository

import (
	"skripsi/features/address/entity"
	"skripsi/features/address/interfaces"
	"skripsi/features/address/mapping"
	"skripsi/utils/pagination"

	"gorm.io/gorm"
)

type addressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) interfaces.AddressRepositoryInterface {
	return &addressRepository{
		db: db,
	}
}

// Create implements interfaces.AddressRepositoryInterface.
func (a *addressRepository) Create(data entity.AddressCore) (entity.AddressCore, error) {
	request := mapping.AddressCoreToAddressModel(data)

	tx := a.db.Create(&request)
	if tx.Error != nil {
		return entity.AddressCore{}, tx.Error
	}

	response := mapping.AddressModelToAddressCore(request)
	return response, nil
}

// DeleteById implements interfaces.AddressRepositoryInterface.
func (a *addressRepository) DeleteById(id string) error {
	panic("unimplemented")
}

// GetAll implements interfaces.AddressRepositoryInterface.
func (a *addressRepository) GetAll(search string, page int, limit int) ([]entity.AddressCore, pagination.PageInfo, int, error) {
	panic("unimplemented")
}

// GetById implements interfaces.AddressRepositoryInterface.
func (a *addressRepository) GetById(id string) (entity.AddressCore, error) {
	panic("unimplemented")
}

// UpdateById implements interfaces.AddressRepositoryInterface.
func (a *addressRepository) UpdateById(id string, data entity.AddressCore) error {
	panic("unimplemented")
}