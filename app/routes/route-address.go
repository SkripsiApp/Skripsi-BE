package routes

import (
	"skripsi/features/address/controller"
	"skripsi/features/address/repository"
	"skripsi/features/address/service"
	"skripsi/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteAddress(e *echo.Group, db *gorm.DB) {
	addressRepository := repository.NewAddressRepository(db)
	addressService := service.NewAddressService(addressRepository)
	addressController := controller.NewAddressController(addressService)

	adress := e.Group("/address", jwt.JWTMiddleware())
	adress.POST("", addressController.Create)
	adress.GET("/:id", addressController.GetById)
	adress.DELETE("/:id", addressController.DeleteById)
}
