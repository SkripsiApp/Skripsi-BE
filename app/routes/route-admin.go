package routes

import (
	"skripsi/features/admin/controller"
	"skripsi/features/admin/repository"
	"skripsi/features/admin/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteAdmin(e *echo.Group, db *gorm.DB) {
	adminRepository := repository.NewAdminRepository(db)
	adminService := service.NewAdminService(adminRepository)
	adminController := controller.NewAdminController(adminService)

	admin := e.Group("/admin")
	admin.POST("/register", adminController.Register)
	admin.POST("/login", adminController.Login)
}

