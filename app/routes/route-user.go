package routes

import (
	"skripsi/features/user/controller"
	"skripsi/features/user/repository"
	"skripsi/features/user/service"
	"skripsi/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteUser(e *echo.Group, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	// User
	e.POST("/register", userController.Register)
	e.POST("/login", userController.Login)
	e.GET("/profile", userController.GetById, jwt.JWTMiddleware())
	e.PUT("/profile", userController.UpdateById, jwt.JWTMiddleware())

	// Admin
	e.GET("/users", userController.GetAll, jwt.JWTMiddleware())

}
