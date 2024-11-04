package response

import "skripsi/features/voucher/entity"

func VoucherCoreToVoucherResponse(voucher entity.VoucherCore) VoucherResponse {
	return VoucherResponse{
		Id:          voucher.Id,
		Name:        voucher.Name,
		Description: voucher.Description,
		Discount:    voucher.Discount,
	}
}

func ListVoucherCoreToListVoucherResponse(vouchers []entity.VoucherCore) []VoucherResponse {
	var voucherResponses []VoucherResponse
	for _, voucher := range vouchers {
		voucherResponses = append(voucherResponses, VoucherCoreToVoucherResponse(voucher))
	}
	return voucherResponses
}