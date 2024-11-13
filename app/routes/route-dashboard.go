package routes

import (
	"skripsi/features/dashboard/controller"
	"skripsi/features/dashboard/repository"
	"skripsi/features/dashboard/service"
	"skripsi/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteDashboard(e *echo.Group, db *gorm.DB) {
	dashboardRepository := repository.NewDashboardRepository(db)
	dashboardService := service.NewDashboardService(dashboardRepository)
	dashboardController := controller.NewDashboardController(dashboardService)

	e.GET("/dashboard", dashboardController.DashboardData, jwt.JWTMiddleware())
}
