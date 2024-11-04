package controller

import (
	"skripsi/features/admin/dto/request"
	"skripsi/features/admin/dto/response"
	"skripsi/features/admin/interfaces"
	"skripsi/utils/constant"
	"skripsi/utils/helper"

	"github.com/labstack/echo/v4"
)

type adminController struct {
	adminService interfaces.AdminServiceInterface
}

func NewAdminController(adminService interfaces.AdminServiceInterface) *adminController {
	return &adminController{
		adminService: adminService,
	}
}

func (a *adminController) Register(e echo.Context) error {
	input := request.AdminRegister{}
	errBind := e.Bind(&input)
	if errBind != nil {
		return helper.ResponseError(400, "invalid input data")
	}

	data := request.AdminRegisterToAdminCore(input)
	_, err := a.adminService.Register(data)
	if err != nil {
		return err
	}

	return e.JSON(200, helper.ResponseSuccess(constant.SUCCESS_CREATE_DATA))
}

func (a *adminController) Login(e echo.Context) error {
	input := request.AdminLogin{}
	errBind := e.Bind(&input)
	if errBind != nil {
		return helper.ResponseError(400, "invalid input data")
	}

	data, token, err := a.adminService.Login(input.Email, input.Password)
	if err != nil {
		return err
	}

	response := response.AdminCoreToLoginResponse(data, token)

	return e.JSON(200, helper.ResponseSuccessWithData(constant.SUCCESS_LOGIN, response))
}
