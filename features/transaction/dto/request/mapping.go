package request

import "skripsi/features/transaction/entity"

func TransactionRequestToTransactionCore(data TransactionRequest) entity.TransactionCore {
	transactionCore := entity.TransactionCore{
		UserId:            data.UserId,
		VoucherId:         data.VoucherId,
		AddressId:         data.AddressId,
		UsePoint:          data.UsePoint,
		PointUsed:         data.PointUsed,
		CourierName:       data.CourierName,
		ShippingCost:      data.ShippingCost,
		TransactionDetail: ListTransactionDetailRequestToTransactionDetailCore(data.TransactionDetailRequest),
	}
	return transactionCore
}

func ListTransactionRequestToTransactionCore(data []TransactionRequest) []entity.TransactionCore {
	listTransactionCore := []entity.TransactionCore{}
	for _, transactionRequest := range data {
		transactionCore := TransactionRequestToTransactionCore(transactionRequest)
		listTransactionCore = append(listTransactionCore, transactionCore)
	}
	return listTransactionCore
}

func TransactionDetailRequestToTransactionDetailCore(data TransactionDetailRequest) entity.TransactionDetailCore {
	transactionDetailCore := entity.TransactionDetailCore{
		ProductId:  data.ProductId,
		Size:       data.Size,
		Quantity:   data.Quantity,
	}
	return transactionDetailCore
}

func ListTransactionDetailRequestToTransactionDetailCore(data []TransactionDetailRequest) []entity.TransactionDetailCore {
	listTransactionDetailCore := []entity.TransactionDetailCore{}
	for _, transactionDetailRequest := range data {
		transactionDetailCore := TransactionDetailRequestToTransactionDetailCore(transactionDetailRequest)
		listTransactionDetailCore = append(listTransactionDetailCore, transactionDetailCore)
	}
	return listTransactionDetailCore
}
