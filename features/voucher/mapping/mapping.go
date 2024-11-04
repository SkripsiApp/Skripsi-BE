package mapping

import (
	"skripsi/features/voucher/entity"
	"skripsi/features/voucher/model"
)

// Mapping Core to Model
func VoucherCoreToVoucherModel(data entity.VoucherCore) model.Voucher {
	voucherModel := model.Voucher{
		Id:          data.Id,
		AdminId:     data.AdminId,
		Name:        data.Name,
		Description: data.Description,
		Discount:    data.Discount,
	}
	return voucherModel
}

func ListVoucherCoreToVoucherModel(data []entity.VoucherCore) []model.Voucher {
	listVoucherModel := []model.Voucher{}
	for _, voucherCore := range data {
		voucherModel := VoucherCoreToVoucherModel(voucherCore)
		listVoucherModel = append(listVoucherModel, voucherModel)
	}
	return listVoucherModel
}

// Mapping Model to Core
func VoucherModelToVoucherCore(data model.Voucher) entity.VoucherCore {
	voucherCore := entity.VoucherCore{
		Id:          data.Id,
		AdminId:     data.AdminId,
		Name:        data.Name,
		Description: data.Description,
		Discount:    data.Discount,
	}
	return voucherCore
}

func ListVoucherModelToVoucherCore(data []model.Voucher) []entity.VoucherCore {
	listVoucherCore := []entity.VoucherCore{}
	for _, voucherModel := range data {
		voucherCore := VoucherModelToVoucherCore(voucherModel)
		listVoucherCore = append(listVoucherCore, voucherCore)
	}
	return listVoucherCore
}

