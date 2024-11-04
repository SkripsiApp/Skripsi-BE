package request

import "skripsi/features/voucher/entity"

func VoucherRequestToVoucherCore(voucher VoucherRequest) entity.VoucherCore {
	return entity.VoucherCore{
		Id:          voucher.Id,
		AdminId:     voucher.AdminId,
		Name:        voucher.Name,
		Description: voucher.Description,
		Discount:    voucher.Discount,
	}
}