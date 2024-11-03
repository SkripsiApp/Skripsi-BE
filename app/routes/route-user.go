package routes

import (
	"skripsi/features/user/controller"
	"skripsi/features/user/repository"
	"skripsi/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteUser(e *echo.Group, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	e.POST("/register", userController.Register)
	e.GET("/users", userController.GetAll)
	e.GET("/users/:id", userController.GetById)
}
