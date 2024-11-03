package controller

import (
	"skripsi/features/user/dto/request"
	"skripsi/features/user/dto/response"
	"skripsi/features/user/interfaces"
	"skripsi/utils/constant"
	"skripsi/utils/helper"
	"skripsi/utils/jwt"
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
	_, role, err := jwt.ExtractToken(e)
	if role != constant.ADMIN {
		return helper.ResponseError(401, constant.ERROR_AKSES_ROLE)
	}

	if err != nil {
		return helper.ResponseError(401, "invalid token")
	}

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

func (uh *userController) GetById(e echo.Context) error {
	id, _, errExtract := jwt.ExtractToken(e)
	if errExtract != nil {
		return helper.ResponseError(401, "invalid token")
	}

	data, err := uh.userService.GetById(id)
	if err != nil {
		return err
	}

	response := response.UserCoreToUserResponse(data)

	return e.JSON(200, helper.ResponseSuccessWithData(constant.SUCCESS_GET_DATA, response))
}

func (uh *userController) UpdateById(e echo.Context) error {
	id, _, errExtract := jwt.ExtractToken(e)
	if errExtract != nil {
		return helper.ResponseError(401, "invalid token")
	}

	input := request.UserUpdate{}
	errBind := e.Bind(&input)
	if errBind != nil {
		return helper.ResponseError(400, "invalid input data")
	}

	data := request.UserUpdateToUserCore(input)
	err := uh.userService.UpdateById(id, data)
	if err != nil {
		return err
	}

	return e.JSON(200, helper.ResponseSuccess(constant.SUCCESS_UPDATE_DATA))
}

func (uh *userController) Login(e echo.Context) error {
	input := request.UserLogin{}
	errBind := e.Bind(&input)
	if errBind != nil {
		return helper.ResponseError(400, "invalid input data")
	}

	data, token, err := uh.userService.Login(input.Email, input.Password)
	if err != nil {
		return err
	}

	response := response.UserCoreToLoginResponse(data, token)

	return e.JSON(200, helper.ResponseSuccessWithData(constant.SUCCESS_GET_DATA, response))
}
