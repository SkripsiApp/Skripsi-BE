package response

import "skripsi/features/transaction/entity"

func TransactionCoreToTransactionResponse(data entity.TransactionCore) TransactionResponse {
	return TransactionResponse{
		Id:                data.Id,
		VoucherId:         data.VoucherId,
		AddressId:         data.AddressId,
		OriginalPrice:     data.OriginalPrice,
		TotalPrice:        data.TotalPrice,
		TotalPoint:        data.TotalPoint,
		UsePoint:          data.UsePoint,
		PointUsed:         data.PointUsed,
		DiscountAmount:    data.DiscountAmount,
		CourierName:       data.CourierName,
		ShippingCost:      data.ShippingCost,
		NoReceipt:         data.NoReceipt,
		Status:            data.Status,
		PaymentURL:        data.PaymentURL,
		CreatedAt:         data.CreatedAt,
		UpdatedAt:         data.UpdatedAt,
		TransactionDetails: ListTransactionDetailCoreToTransactionDetailResponse(data.TransactionDetail),
	}
}

func ListTransactionCoreToTransactionResponse(data []entity.TransactionCore) []TransactionResponse {
	var transactions []TransactionResponse
	for _, v := range data {
		transactions = append(transactions, TransactionCoreToTransactionResponse(v))
	}
	return transactions
}

func TransactionDetailCoreToTransactionDetailResponse(data entity.TransactionDetailCore) TransactionDetailResponse {
	return TransactionDetailResponse{
		Id:         data.Id,
		ProductId:  data.ProductId,
		Size:       data.Size,
		Quantity:   data.Quantity,
		TotalPrice: data.TotalPrice,
	}
}

func ListTransactionDetailCoreToTransactionDetailResponse(data []entity.TransactionDetailCore) []TransactionDetailResponse {
	var transactionDetails []TransactionDetailResponse
	for _, v := range data {
		transactionDetails = append(transactionDetails, TransactionDetailCoreToTransactionDetailResponse(v))
	}
	return transactionDetails
}
