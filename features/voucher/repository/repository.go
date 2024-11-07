package repository

import (
	"errors"
	"skripsi/features/voucher/entity"
	"skripsi/features/voucher/interfaces"
	"skripsi/features/voucher/mapping"
	"skripsi/features/voucher/model"
	"skripsi/utils/constant"
	"skripsi/utils/helper"
	"skripsi/utils/pagination"

	"gorm.io/gorm"
)

type voucherRepository struct {
	db *gorm.DB
}

func NewVoucherRepository(db *gorm.DB) interfaces.VoucherRepositoryInterface {
	return &voucherRepository{
		db: db,
	}
}

// Create implements interfaces.VoucherRepositoryInterface.
func (v *voucherRepository) Create(data entity.VoucherCore) (entity.VoucherCore, error) {
	request := mapping.VoucherCoreToVoucherModel(data)

	tx := v.db.Create(&request)
	if tx.Error != nil {
		return entity.VoucherCore{}, tx.Error
	}

	response := mapping.VoucherModelToVoucherCore(request)
	return response, nil
}

// DeleteById implements interfaces.VoucherRepositoryInterface.
func (v *voucherRepository) DeleteById(id string) error {
	data := model.Voucher{}

	tx := v.db.Where("id = ?", id).Delete(&data)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return helper.ResponseError(404, constant.ERROR_DATA_NOT_FOUND)
		}
		return tx.Error
	}

	return nil
}

// GetAll implements interfaces.VoucherRepositoryInterface.
func (v *voucherRepository) GetAll(search string, page int, limit int) ([]entity.VoucherCore, pagination.PageInfo, int, error) {
	data := []model.Voucher{}

	offset := (page - 1) * limit
	query := v.db.Model(&model.Voucher{})

	if search != "" {
		query = query.Where("name LIKE ? or id LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	var totalCount int64
	tx := query.Count(&totalCount).Find(&data)
	if tx.Error != nil {
		return nil, pagination.PageInfo{}, 0, tx.Error
	}

	query = query.Offset(offset).Limit(limit)

	tx = query.Find(&data)
	if tx.Error != nil {
		return nil, pagination.PageInfo{}, 0, tx.Error
	}

	dataResponse := mapping.ListVoucherModelToVoucherCore(data)

	pageInfo := pagination.CalculateData(int(totalCount), limit, page)
	return dataResponse, pageInfo, int(totalCount), nil
}

// GetById implements interfaces.VoucherRepositoryInterface.
func (v *voucherRepository) GetById(id string) (entity.VoucherCore, error) {
	data := model.Voucher{}

	tx := v.db.Where("id = ?", id).First(&data)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return entity.VoucherCore{}, helper.ResponseError(404, constant.ERROR_DATA_NOT_FOUND)
		}
		return entity.VoucherCore{}, tx.Error
	}

	dataResponse := mapping.VoucherModelToVoucherCore(data)
	return dataResponse, nil
}

// UpdateById implements interfaces.VoucherRepositoryInterface.
func (v *voucherRepository) UpdateById(id string, data entity.VoucherCore) error {
	request := mapping.VoucherCoreToVoucherModel(data)

	tx := v.db.Where("id = ?", id).Updates(&request)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return helper.ResponseError(404, constant.ERROR_DATA_NOT_FOUND)
		}
		return tx.Error
	}

	return nil
}
