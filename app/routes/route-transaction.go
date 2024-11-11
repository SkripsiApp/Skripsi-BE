package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"skripsi/features/transaction/controller"
	transaction "skripsi/features/transaction/repository"
	user "skripsi/features/user/repository"
	voucher "skripsi/features/voucher/repository"
	address "skripsi/features/address/repository"
	product "skripsi/features/product/repository"
	"skripsi/features/transaction/service"

	

	"skripsi/utils/jwt"
)

func RouteTransaction(e *echo.Group, db *gorm.DB) {
	transactionRepository := transaction.NewTransactionRepository(db)
	userRepository := user.NewUserRepository(db)
	addressRepository := address.NewAddressRepository(db)
	voucherRepository := voucher.NewVoucherRepository(db)
	productRepository := product.NewProductRepository(db)
	transactionService := service.NewTransactionService(transactionRepository, userRepository, voucherRepository, addressRepository, productRepository)
	transactionController := controller.NewTransactionController(transactionService)

	transaction := e.Group("/transaction")
	transaction.POST("", transactionController.CreateTransaction, jwt.JWTMiddleware())
	transaction.POST("/midtrans/notification", transactionController.HandleMidtransNotification)
}