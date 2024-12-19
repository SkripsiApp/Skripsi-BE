package service

import (
	"fmt"
	"log"
	"os"
	address "skripsi/features/address/interfaces"
	product "skripsi/features/product/interfaces"
	"skripsi/features/transaction/entity"
	transaction "skripsi/features/transaction/interfaces"
	userCore "skripsi/features/user/entity"
	user "skripsi/features/user/interfaces"
	voucher "skripsi/features/voucher/interfaces"
	"skripsi/utils/helper"
	"skripsi/utils/pagination"

	"github.com/joho/godotenv"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type transactionService struct {
	transactionRepository transaction.TransactionRepositoryInterface
	voucherRepository     voucher.VoucherRepositoryInterface
	addressRepository     address.AddressRepositoryInterface
	userRepository        user.UserRepositoryInterface
	productRepository     product.ProductRepositoryInterface
}

func NewTransactionService(transactionRepository transaction.TransactionRepositoryInterface, userRepository user.UserRepositoryInterface, voucherRepository voucher.VoucherRepositoryInterface, addressRepository address.AddressRepositoryInterface, productRepository product.ProductRepositoryInterface) transaction.TransactionServiceInterface {
	return &transactionService{
		transactionRepository: transactionRepository,
		voucherRepository:     voucherRepository,
		addressRepository:     addressRepository,
		userRepository:        userRepository,
		productRepository:     productRepository,
	}
}

// CreateTransaction implements interfaces.TransactionServiceInterface.
func (t *transactionService) CreateTransaction(data entity.TransactionCore) (entity.TransactionCore, string, error) {
	fmt.Println("Received Transaction Data:", data)
	fmt.Printf("Initial UsePoint: %v, PointUsed: %d\n", data.UsePoint, data.PointUsed)

	if data.UserId == "" {
		return entity.TransactionCore{}, "", helper.ResponseError(400, "user id tidak boleh kosong")
	}

	var discountAmount int
	if data.VoucherId != nil {
		voucher, err := t.voucherRepository.GetById(*data.VoucherId)
		if err != nil {
			return entity.TransactionCore{}, "", helper.ResponseError(400, "voucher tidak ditemukan")
		}

		discountAmount += voucher.Discount
		fmt.Println("Voucher Discount:", data.DiscountAmount)
	}

	if data.AddressId == "" {
		return entity.TransactionCore{}, "", helper.ResponseError(400, "alamat tidak boleh kosong")
	}

	address, err := t.addressRepository.GetById(data.AddressId, data.UserId)
	if err != nil {
		return entity.TransactionCore{}, "", helper.ResponseError(400, "alamat tidak ditemukan")
	}

	data.AddressId = address.Id

	var totalPrice int
	var originalPrice int
	var itemDetails []midtrans.ItemDetails
	for i, detail := range data.TransactionDetail {
		product, err := t.productRepository.GetById(detail.ProductId)
		if err != nil {
			return entity.TransactionCore{}, "", helper.ResponseError(400, "produk tidak ditemukan")
		}

		productSize, err := t.productRepository.GetProductIdAndSize(detail.ProductId, detail.Size)
		if err != nil {
			return entity.TransactionCore{}, "", helper.ResponseError(400, "ukuran produk tidak ditemukan")
		}

		if productSize.Stock < detail.Quantity {
			return entity.TransactionCore{}, "", helper.ResponseError(400, "stock produk tidak mencukupi")
		}

		if err := t.productRepository.DecreaseStock(productSize.Id, detail.Quantity); err != nil {
			return entity.TransactionCore{}, "", err
		}

		if err := t.productRepository.IncreaseSold(detail.ProductId, detail.Quantity); err != nil {
			return entity.TransactionCore{}, "", err
		}

		itemDetails = append(itemDetails, midtrans.ItemDetails{
			ID:    detail.ProductId,
			Price: int64(product.Price),
			Qty:   int32(detail.Quantity),
			Name:  fmt.Sprintf("%s - Size: %s", product.Name, detail.Size),
		})
		
		productTotalPrice := product.Price * detail.Quantity
		totalPrice += productTotalPrice

		originalPrice += productTotalPrice
		data.TransactionDetail[i].TotalPrice = productTotalPrice
	}

	var updatedPoint int
	user, err := t.userRepository.GetById(data.UserId)
	if err != nil {
		return entity.TransactionCore{}, "", helper.ResponseError(400, "user tidak ditemukan")
	}

	fmt.Printf("User's current points before transaction: %d\n", user.Point)

	if data.UsePoint && data.PointUsed > 0 {
		fmt.Printf("Attempting to use points - UsePoint: %v, PointUsed: %d\n", data.UsePoint, data.PointUsed)
		if user.Point < data.PointUsed {
			return entity.TransactionCore{}, "", helper.ResponseError(400, "point tidak cukup")
		}
		discountAmount += data.PointUsed
		updatedPoint = user.Point - data.PointUsed
		fmt.Println("Points Used:", data.PointUsed)
		fmt.Println("Updated Points After Deduction:", updatedPoint)
	} else {
		updatedPoint = user.Point
		fmt.Println("Not using points or PointUsed is 0")
	}

	if discountAmount > 0 {
		itemDetails = append(itemDetails, midtrans.ItemDetails{
			ID:    "DISCOUNT",
			Price: int64(-discountAmount),
			Qty:   1,
			Name:  "Discount",
		})
	}

	if data.ShippingCost > 0 {
		itemDetails = append(itemDetails, midtrans.ItemDetails{
			ID:    "SHIPPING",
			Price: int64(data.ShippingCost),
			Qty:   1,
			Name:  "Shipping Cost",
		})
	}

	totalPrice -= discountAmount
	if totalPrice < 0 {
		totalPrice = 0
	}

	totalPrice += data.ShippingCost
	data.OriginalPrice = originalPrice
	data.TotalPrice = totalPrice
	data.DiscountAmount = discountAmount
	data.TotalPoint = totalPrice / 100

	fmt.Printf("Final calculation - Total Price: %d, Discount Amount: %d, Total Point: %d\n", totalPrice, discountAmount, data.TotalPoint)

	// updatedPoint += data.TotalPoint

	data.Status = "Pending"
	transaction, err := t.transactionRepository.CreateTransaction(data)
	if err != nil {
		return entity.TransactionCore{}, "", err
	}

	godotenv.Load()
	midtransClient := snap.Client{}
	midtransClient.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transaction.Id,
			GrossAmt: int64(totalPrice),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			Email: user.Email,
		},
		EnabledPayments: snap.AllSnapPaymentType,
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		Items: &itemDetails,
	}

	snapResp, err := midtransClient.CreateTransaction(req)
	if snapResp == nil || snapResp.RedirectURL == "" {
		return entity.TransactionCore{}, "", helper.ResponseError(500, "gagal mendapatkan URL pembayaran dari Midtrans")
	}

	snapRedirectURL := snapResp.RedirectURL

	fmt.Println("Updating User Points:", updatedPoint)
	if err := t.userRepository.UpdatedPoint(user.Id, updatedPoint); err != nil {
		return entity.TransactionCore{}, "", err
	}

	fmt.Println("Transaction created successfully. Final transaction data:", transaction)

	return transaction, snapRedirectURL, nil
}

// GetAllTransaction implements interfaces.TransactionServiceInterface.
func (t *transactionService) GetAllTransaction(search string, page int, limit int) ([]entity.TransactionCore, pagination.PageInfo, int, error) {
	if limit > 10 {
		return nil, pagination.PageInfo{}, 0, helper.ResponseError(400, "limit tidak boleh lebih dari 10")
	}

	page, limit = helper.ValidateCountLimitAndPage(page, limit)

	dataTransaction, pageInfo, totalCount, err := t.transactionRepository.GetAllTransaction(search, page, limit)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}

	return dataTransaction, pageInfo, totalCount, nil
}

// GetTransactionById implements interfaces.TransactionServiceInterface.
func (t *transactionService) GetTransactionById(id string) (entity.TransactionCore, error) {
	if id == "" {
		return entity.TransactionCore{}, helper.ResponseError(400, "id transaksi tidak boleh kosong")
	}

	transaction, err := t.transactionRepository.GetTransactionById(id)
	if err != nil {
		return entity.TransactionCore{}, helper.ResponseError(404, "transaksi tidak ditemukan")
	}

	return transaction, nil
}

// UpdateStatusTransactionById implements interfaces.TransactionServiceInterface.
func (t *transactionService) UpdateStatusTransactionById(id string, status string) error {
	panic("unimplemented")
}

// HandleMidtransNotification implements interfaces.TransactionServiceInterface.
func (t *transactionService) HandleMidtransNotification(notification helper.MidtransNotificationPayload) error {
	transaction, err := t.transactionRepository.GetTransactionById(notification.OrderID)
	if err != nil {
		return helper.ResponseError(404, "transaksi tidak ditemukan")
	}

	// if transaction.Status != "Pending" {
	// 	return nil
	// }

	transaction.PaymentType = notification.PaymentType

	// Handle different transaction statuses
	switch notification.TransactionStatus {
	case "capture":
		if notification.FraudStatus == "challenge" {
			transaction.Status = "Challenge"
		} else if notification.FraudStatus == "accept" {
			return t.processSuccessfulPayment(transaction)
		}

	case "settlement":
		return t.processSuccessfulPayment(transaction)

	case "cancel", "deny", "expire", "failure":
		return t.processFailedPayment(transaction)

	default:
		return t.transactionRepository.UpdatePaymentDetails(
			transaction.Id,
			notification.PaymentType,
			transaction.Status,
		)
	}

	return nil
}

func (t *transactionService) processSuccessfulPayment(transaction entity.TransactionCore) error {
	// Update transaction status
	transaction.Status = "Paid"
	if err := t.transactionRepository.UpdateStatusTransactionById(transaction.Id, transaction.Status); err != nil {
		return err
	}

	// Get user data
	user, err := t.userRepository.GetById(transaction.UserId)
	if err != nil {
		return helper.ResponseError(404, "user tidak ditemukan")
	}

	// Calculate and update user points
	earnedPoints := transaction.TotalPoint
	updatedPoints := user.Point + earnedPoints

	// Update user points
	return t.userRepository.UpdateById(transaction.UserId, userCore.UsersCore{
		Point: updatedPoints,
	})
}

// Helper method to process failed payments
func (t *transactionService) processFailedPayment(transaction entity.TransactionCore) error {
	// Update transaction status
	transaction.Status = "Failed"
	if err := t.transactionRepository.UpdateStatusTransactionById(transaction.Id, transaction.Status); err != nil {
		return err
	}

	// Restore product stock
	if len(transaction.TransactionDetail) == 0 {
		log.Printf("Transaction with ID %v has no details", transaction.Id)
	} else {
		for _, detail := range transaction.TransactionDetail {
			log.Printf("Processing product with ID: %v, Size: %s, Quantity: %d", detail.ProductId, detail.Size, detail.Quantity)

			productSize, err := t.productRepository.GetProductIdAndSize(detail.ProductId, detail.Size)
			if err != nil {
				log.Printf("Error getting product size - ProductID: %v, Size: %s", detail.ProductId, detail.Size)
				return helper.ResponseError(400, "ukuran produk tidak ditemukan")
			}

			log.Printf("Found product size with ID: %v", productSize.Id)

			// Mengembalikan stock produk yang dibatalkan
			if err := t.productRepository.IncreaseStock(productSize.Id, detail.Quantity); err != nil {
				log.Printf("Error increasing stock - ProductSizeID: %v, Quantity: %d", productSize.Id, detail.Quantity)
				return err
			}

			// Mengurangi jumlah produk yang terjual
			if err := t.productRepository.DecreaseSold(detail.ProductId, detail.Quantity); err != nil {
				log.Printf("Error decreasing sold - ProductID: %v, Quantity: %d", detail.ProductId, detail.Quantity)
				return err
			}
		}
	}

	// If points were used, restore them
	if transaction.UsePoint && transaction.PointUsed > 0 {
		user, err := t.userRepository.GetById(transaction.UserId)
		if err != nil {
			return helper.ResponseError(404, "user tidak ditemukan")
		}

		restoredPoints := user.Point + transaction.PointUsed
		return t.userRepository.UpdateById(transaction.UserId, userCore.UsersCore{
			Point: restoredPoints,
		})
	}

	return nil
}

// GetAllTransactionByUserId implements interfaces.TransactionServiceInterface.
func (t *transactionService) GetAllTransactionByUserId(userId string, search string, page int, limit int) ([]entity.TransactionCore, pagination.PageInfo, int, error) {
	if userId == "" {
		return nil, pagination.PageInfo{}, 0, helper.ResponseError(400, "user id tidak boleh kosong")
	}

	if limit > 10 {
		return nil, pagination.PageInfo{}, 0, helper.ResponseError(400, "limit tidak boleh lebih dari 10")
	}

	page, limit = helper.ValidateCountLimitAndPage(page, limit)

	dataTransaction, pageInfo, totalCount, err := t.transactionRepository.GetAllTransactionByUserId(userId, search, page, limit)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}

	return dataTransaction, pageInfo, totalCount, nil
}

// UpdateNoResiTransactionById implements interfaces.TransactionServiceInterface.
func (t *transactionService) UpdateNoReceiptTransactionById(id string, resi string) error {
	if id == "" {
		return helper.ResponseError(400, "id transaksi tidak boleh kosong")
	}

	if resi == "" {
		return helper.ResponseError(400, "no resi tidak boleh kosong")
	}

	transaction, err := t.transactionRepository.GetTransactionById(id)
	if err != nil {
		return helper.ResponseError(404, "transaksi tidak ditemukan")
	}

	if transaction.Status != "Paid" {
		return helper.ResponseError(400, "status transaksi harus Paid")
	}

	transaction.Status = "Shipped"

	err = t.transactionRepository.UpdateNoReceiptTransactionById(id, resi)
	if err != nil {
		return err
	}

	return nil
}

// UpdateStatusTransactionUserById implements interfaces.TransactionServiceInterface.
func (t *transactionService) UpdateStatusTransactionUserById(userId string, transactionId string, status string) error {
	if userId == "" {
		return helper.ResponseError(400, "user id tidak boleh kosong")
	}

	if transactionId == "" {
		return helper.ResponseError(400, "id transaksi tidak boleh kosong")
	}

	if status == "" {
		return helper.ResponseError(400, "status tidak boleh kosong")
	}

	if status != "Done" {
		return helper.ResponseError(400, "status hanya bisa diupdate ke done")
	}

	transaction, err := t.transactionRepository.GetTransactionById(transactionId)
	if err != nil {
		return helper.ResponseError(404, "transaksi tidak ditemukan")
	}

	if transaction.UserId != userId {
		return helper.ResponseError(400, "transaksi tidak ditemukan")
	}

	if transaction.Status != "Shipped" {
		return helper.ResponseError(400, "status transaksi harus shipped")
	}

	err = t.transactionRepository.UpdateStatusTransactionUserById(userId, transactionId, status)
	if err != nil {
		return err
	}

	return nil
}