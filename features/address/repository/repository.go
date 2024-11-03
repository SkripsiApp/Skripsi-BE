package repository

import (
	"errors"
	"skripsi/features/address/entity"
	"skripsi/features/address/interfaces"
	"skripsi/features/address/mapping"
	"skripsi/features/address/model"
	"skripsi/utils/constant"
	"skripsi/utils/helper"

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
func (a *addressRepository) DeleteById(id string, userId string) error {
	data := model.Address{}

	tx := a.db.Where("id = ? AND user_id = ?", id, userId).Delete(&data)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return helper.ResponseError(404, constant.ERROR_DATA_NOT_FOUND)
		}
		return tx.Error
	}

	return nil
}

// GetAll implements interfaces.AddressRepositoryInterface.
func (a *addressRepository) GetAll(idUser string) ([]entity.AddressCore, error) {
	data := []model.Address{}

	tx := a.db.Where("user_id = ?", idUser).Find(&data)
	if tx.Error != nil {
		return nil, tx.Error
	}

	response := mapping.ListAddressModelToAddressCore(data)
	return response, nil
}

// GetById implements interfaces.AddressRepositoryInterface.
func (a *addressRepository) GetById(id string, userId string) (entity.AddressCore, error) {
	data := model.Address{}

	tx := a.db.Where("id = ? AND user_id = ?", id, userId).First(&data)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return entity.AddressCore{}, helper.ResponseError(404, constant.ERROR_DATA_NOT_FOUND)
		}
		return entity.AddressCore{}, tx.Error
	}

	response := mapping.AddressModelToAddressCore(data)
	return response, nil
}

// UpdateById implements interfaces.AddressRepositoryInterface.
func (a *addressRepository) UpdateById(id string, userId string, data entity.AddressCore) error {
	request := mapping.AddressCoreToAddressModel(data)

	tx := a.db.Where("id = ? AND user_id = ?", id, userId).Updates(&request)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return helper.ResponseError(404, constant.ERROR_DATA_NOT_FOUND)
		}
		return tx.Error
	}

	return nil
}

// FindById implements interfaces.AddressRepositoryInterface.
func (a *addressRepository) FindById(id string) (entity.AddressCore, error) {
	data := model.Address{}

	tx := a.db.Where("id = ?", id).First(&data)
	if tx.Error != nil {
		return entity.AddressCore{}, tx.Error
	}

	response := mapping.AddressModelToAddressCore(data)
	return response, nil
}
