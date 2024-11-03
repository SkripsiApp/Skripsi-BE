package controller

import (
	"skripsi/features/user/dto/request"
	"skripsi/features/user/dto/response"
	"skripsi/features/user/interfaces"
	"skripsi/utils/constant"
	"skripsi/utils/helper"
	"strconv"

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

func (uh *userController) GetAll(e echo.Context) error {
	search := e.QueryParam("search")
	page, _ := strconv.Atoi(e.QueryParam("page"))
	limit, _ := strconv.Atoi(e.QueryParam("limit"))

	data, pageInfo, totalCount, err := uh.userService.GetAll(search, page, limit)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		return helper.ResponseError(200, "data belum tersedia")
	}

	response := response.ListUserCoreToUserResponse(data)

	return e.JSON(200, helper.ResponseSuccessWithPagnationAndCount(constant.SUCCESS_GET_DATA, response, pageInfo, totalCount))
}
