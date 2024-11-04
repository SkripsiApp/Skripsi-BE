package service

import (
	"skripsi/features/voucher/entity"
	"skripsi/features/voucher/interfaces"
	"skripsi/utils/constant"
	"skripsi/utils/helper"
	"skripsi/utils/pagination"
)

type voucherService struct {
	voucherRepository interfaces.VoucherRepositoryInterface
}

func NewVoucherService(voucherRepository interfaces.VoucherRepositoryInterface) interfaces.VoucherServiceInterface {
	return &voucherService{
		voucherRepository: voucherRepository,
	}
}

// Create implements interfaces.VoucherServiceInterface.
func (v *voucherService) Create(data entity.VoucherCore) (entity.VoucherCore, error) {
	if data.AdminId == "" {
		return entity.VoucherCore{}, helper.ResponseError(400, constant.ERROR_DATA_ID)
	}

	if data.Id == "" || data.Name == "" || data.Description == "" || data.Discount == 0 {
		return entity.VoucherCore{}, helper.ResponseError(400, constant.ERROR_EMPTY)
	}

	response, err := v.voucherRepository.Create(data)
	if err != nil {
		return entity.VoucherCore{}, err
	}

	return response, nil
}

// DeleteById implements interfaces.VoucherServiceInterface.
func (v *voucherService) DeleteById(id string) error {
	if id == "" {
		return helper.ResponseError(400, constant.ERROR_DATA_ID)
	}

	err := v.voucherRepository.DeleteById(id)
	if err != nil {
		return err
	}

	return nil
}

// GetAll implements interfaces.VoucherServiceInterface.
func (v *voucherService) GetAll(search string, page int, limit int) ([]entity.VoucherCore, pagination.PageInfo, int, error) {
	if limit > 10 {
		return nil, pagination.PageInfo{}, 0, helper.ResponseError(400, "Limit tidak boleh lebih dari 10")
	}

	page, limit = helper.ValidateCountLimitAndPage(page, limit)

	dataVoucher, pageInfo, totalCount, err := v.voucherRepository.GetAll(search, page, limit)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}

	return dataVoucher, pageInfo, totalCount, nil
}

// GetById implements interfaces.VoucherServiceInterface.
func (v *voucherService) GetById(id string) (entity.VoucherCore, error) {
	if id == "" {
		return entity.VoucherCore{}, helper.ResponseError(400, constant.ERROR_DATA_ID)
	}

	data, err := v.voucherRepository.GetById(id)
	if err != nil {
		return entity.VoucherCore{}, err
	}

	return data, nil
}

// UpdateById implements interfaces.VoucherServiceInterface.
func (v *voucherService) UpdateById(id string, data entity.VoucherCore) error {
	if id == "" {
		return helper.ResponseError(400, constant.ERROR_DATA_ID)
	}

	if data.Id == "" || data.Name == "" || data.Description == "" || data.Discount == 0 {
		return helper.ResponseError(400, constant.ERROR_EMPTY)
	}

	err := v.voucherRepository.UpdateById(id, data)
	if err != nil {
		return err
	}

	return nil
}
