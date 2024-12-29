package response

import "skripsi/features/transaction/entity"

func TransactionCoreToTransactionResponse(data entity.TransactionCore) TransactionResponse {
	return TransactionResponse{
		Id:                 data.Id,
		UserId:             data.UserId,
		VoucherId:          data.VoucherId,
		AddressId:          data.AddressId,
		AddressName:        data.AddressName,
		NoTransaction:      data.NoTransaction,
		OriginalPrice:      data.OriginalPrice,
		TotalPrice:         data.TotalPrice,
		TotalPoint:         data.TotalPoint,
		UsePoint:           data.UsePoint,
		PointUsed:          data.PointUsed,
		VoucherDiscount:    data.VoucherDiscount,
		DiscountAmount:     data.DiscountAmount,
		CourierName:        data.CourierName,
		ShippingCost:       data.ShippingCost,
		NoReceipt:          data.NoReceipt,
		Status:             data.Status,
		CreatedAt:          data.CreatedAt,
		UpdatedAt:          data.UpdatedAt,
		TransactionDetails: ListTransactionDetailCoreToTransactionDetailResponse(data.TransactionDetail),
	}
}

func TransactionCoreToTransactionRequestResponse(data entity.TransactionCore, snap string) TransactionRequestResponse {
	return TransactionRequestResponse{
		Id:            data.Id,
		NoTransaction: data.NoTransaction,
		PaymentURL:    snap,
	}
}

// func TransactionCoreToTransactionResponseWithSnap(data entity.TransactionCore, snap string) TransactionResponse {
// 	return TransactionResponse{
// 		Id:                 data.Id,
// 		UserId:             data.UserId,
// 		VoucherId:          data.VoucherId,
// 		AddressId:          data.AddressId,
// 		NoTransaction:      data.NoTransaction,
// 		OriginalPrice:      data.OriginalPrice,
// 		TotalPrice:         data.TotalPrice,
// 		TotalPoint:         data.TotalPoint,
// 		UsePoint:           data.UsePoint,
// 		PointUsed:          data.PointUsed,
// 		VoucherDiscount:    data.VoucherDiscount,
// 		DiscountAmount:     data.DiscountAmount,
// 		CourierName:        data.CourierName,
// 		ShippingCost:       data.ShippingCost,
// 		NoReceipt:          data.NoReceipt,
// 		Status:             data.Status,
// 		CreatedAt:          data.CreatedAt,
// 		UpdatedAt:          data.UpdatedAt,
// 		PaymentURL:         snap,
// 		TransactionDetails: ListTransactionDetailCoreToTransactionDetailResponse(data.TransactionDetail),
// 	}
// }

func ListTransactionCoreToTransactionResponse(data []entity.TransactionCore) []TransactionResponse {
	var transactions []TransactionResponse
	for _, v := range data {
		transactions = append(transactions, TransactionCoreToTransactionResponse(v))
	}
	return transactions
}

func TransactionDetailCoreToTransactionDetailResponse(data entity.TransactionDetailCore) TransactionDetailResponse {
	return TransactionDetailResponse{
		Id:          data.Id,
		ProductId:   data.ProductId,
		ProductName: data.ProductName,
		Image:       data.Image,
		Size:        data.Size,
		Quantity:    data.Quantity,
		TotalPrice:  data.TotalPrice,
	}
}

func ListTransactionDetailCoreToTransactionDetailResponse(data []entity.TransactionDetailCore) []TransactionDetailResponse {
	var transactionDetails []TransactionDetailResponse
	for _, v := range data {
		transactionDetails = append(transactionDetails, TransactionDetailCoreToTransactionDetailResponse(v))
	}
	return transactionDetails
}
