package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoutes(e *echo.Echo, db *gorm.DB) {
	base := e.Group("")

	RouteUser(base, db)
	RouteAddress(base, db)
	RouteAdmin(base, db)
}
