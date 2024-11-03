package controller

import (
	"skripsi/features/user/dto/request"
	"skripsi/features/user/interfaces"
	"skripsi/utils/constant"
	"skripsi/utils/helper"

	"github.com/labstack/echo/v4"
)

type userController struct {
	userService interfaces.UserServiceInterrace
}

func NewUserController(userService interfaces.UserServiceInterrace) *userController {
	return &userController{
		userService: userService,
	}
}

func (uh *userController) Register(e echo.Context) error {
	input := request.UserRegister{}
	errBind := e.Bind(&input)
	if errBind != nil {
		return helper.ResponseError(400, "invalid input data")
	}

	data := request.UserRegisterToUserCore(input)
	_, err := uh.userService.Register(data)
	if err != nil {
		return err
	}

	return e.JSON(200, helper.ResponseSuccess(constant.SUCCESS_CREATE_DATA))
}
