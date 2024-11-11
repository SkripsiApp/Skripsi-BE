package mapping

import (
	"skripsi/features/transaction/entity"
	"skripsi/features/transaction/model"
)

// Mapping Transaction Core to Model Transaction
func TransactionCoreToTransactionModel(data entity.TransactionCore) model.Transaction {
	transactionModel := model.Transaction{
		Id:                data.Id,
		UserId:            data.UserId,
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
		TransactionDetail: ListTransactionDetailCoreToTransactionDetailModel(data.TransactionDetail),
	}
	return transactionModel
}

func ListTransactionCoreToListTransactionModel(data []entity.TransactionCore) []model.Transaction {
	var transactions []model.Transaction
	for _, v := range data {
		transactions = append(transactions, TransactionCoreToTransactionModel(v))
	}
	return transactions
}

// Mapping Transaction Model to Transaction Core
func TransactionModelToTransactionCore(data model.Transaction) entity.TransactionCore {
	transactionCore := entity.TransactionCore{
		Id:                data.Id,
		UserId:            data.UserId,
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
		TransactionDetail: ListTransactionDetailModelToListTransactionDetailCore(data.TransactionDetail),
	}
	return transactionCore
}

func ListTransactionModelToListTransactionCore(data []model.Transaction) []entity.TransactionCore {
	var transactions []entity.TransactionCore
	for _, v := range data {
		transactions = append(transactions, TransactionModelToTransactionCore(v))
	}
	return transactions
}

// Mapping Transaction Detail Core to Model Transaction Detail
func TransactionDetailCoreToTransactionDetailModel(data entity.TransactionDetailCore) model.TransactionDetail {
	transactionDetailModel := model.TransactionDetail{
		Id:            data.Id,
		TransactionId: data.TransactionId,
		ProductId:     data.ProductId,
		Size:          data.Size,
		Quantity:      data.Quantity,
		TotalPrice:    data.TotalPrice,
		CreatedAt:     data.CreatedAt,
		UpdatedAt:     data.UpdatedAt,
	}
	return transactionDetailModel
}

func ListTransactionDetailCoreToTransactionDetailModel(data []entity.TransactionDetailCore) []model.TransactionDetail {
	var transactionDetails []model.TransactionDetail
	for _, v := range data {
		transactionDetails = append(transactionDetails, TransactionDetailCoreToTransactionDetailModel(v))
	}
	return transactionDetails
}

// Mapping Transaction Detail Model to Transaction Detail Core
func TransactionDetailModelToTransactionDetailCore(data model.TransactionDetail) entity.TransactionDetailCore {
	transactionDetailCore := entity.TransactionDetailCore{
		Id:            data.Id,
		TransactionId: data.TransactionId,
		ProductId:     data.ProductId,
		Size:          data.Size,
		Quantity:      data.Quantity,
		TotalPrice:    data.TotalPrice,
		CreatedAt:     data.CreatedAt,
		UpdatedAt:     data.UpdatedAt,
	}
	return transactionDetailCore
}

func ListTransactionDetailModelToListTransactionDetailCore(data []model.TransactionDetail) []entity.TransactionDetailCore {
	var transactionDetails []entity.TransactionDetailCore
	for _, v := range data {
		transactionDetails = append(transactionDetails, TransactionDetailModelToTransactionDetailCore(v))
	}
	return transactionDetails
}
