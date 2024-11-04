package routes

import (
	"skripsi/features/voucher/controller"
	"skripsi/features/voucher/repository"
	"skripsi/features/voucher/service"
	"skripsi/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteVoucher(e *echo.Group, db *gorm.DB) {
	voucherRepository := repository.NewVoucherRepository(db)
	voucherService := service.NewVoucherService(voucherRepository)
	voucherController := controller.NewVoucherController(voucherService)

	voucher := e.Group("/voucher", jwt.JWTMiddleware())
	voucher.POST("", voucherController.Create)
	voucher.GET("", voucherController.GetAll)
	voucher.GET("/:id", voucherController.GetById)
	voucher.PUT("/:id", voucherController.UpdateById)
	voucher.DELETE("/:id", voucherController.DeleteById)
}
