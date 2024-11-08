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

	product := e.Group("/products")
	product.POST("", productController.Create, jwt.JWTMiddleware())
	product.GET("", productController.GetAll)
	product.GET("/:id", productController.GetById)
	product.PUT("/:id", productController.UpdateById, jwt.JWTMiddleware())
	product.DELETE("/:id", productController.DeleteById, jwt.JWTMiddleware())
}
