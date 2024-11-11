package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	address "skripsi/features/address/repository"
	product "skripsi/features/product/repository"
	"skripsi/features/transaction/controller"
	transaction "skripsi/features/transaction/repository"
	"skripsi/features/transaction/service"
	user "skripsi/features/user/repository"
	voucher "skripsi/features/voucher/repository"

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
	transaction.GET("", transactionController.GetAllTransaction, jwt.JWTMiddleware())
	transaction.GET("/:id", transactionController.GetTransactionById, jwt.JWTMiddleware())
	transaction.PUT("/:id", transactionController.UpdateNoReceipt, jwt.JWTMiddleware())
	transaction.POST("/midtrans/notification", transactionController.HandleMidtransNotification)

	transaction.GET("/profile", transactionController.GetAllTransactionByUserId, jwt.JWTMiddleware())
}
