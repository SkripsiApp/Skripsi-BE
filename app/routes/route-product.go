package routes

import (
	"skripsi/features/product/controller"
	"skripsi/features/product/repository"
	"skripsi/features/product/service"
	"skripsi/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteProduct(e *echo.Group, db *gorm.DB) {
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	product := e.Group("/product", jwt.JWTMiddleware())
	product.POST("", productController.Create)
}
