package controller

import (
	"skripsi/features/address/dto/request"
	"skripsi/features/address/interfaces"
	"skripsi/utils/constant"
	"skripsi/utils/helper"
	"skripsi/utils/jwt"

	"github.com/labstack/echo/v4"
)

type addressController struct {
	addressService interfaces.AddressServiceInterface
}

func NewAddressController(addressService interfaces.AddressServiceInterface) *addressController {
	return &addressController{
		addressService: addressService,
	}
}

func (a *addressController) Create(e echo.Context) error {
	id, _, errExtract := jwt.ExtractToken(e)
	if errExtract != nil {
		return errExtract
	}

	input := request.AddressRequest{}
	input.UserId = id

	errBind := e.Bind(&input)
	if errBind != nil {
		return helper.ResponseError(400, constant.ERROR_INVALID_INPUT)
	}

	data := request.AddressRequestToAddressCore(input)
	_, err := a.addressService.Create(data)
	if err != nil {
		return err
	}

	return e.JSON(201, helper.ResponseSuccess(constant.SUCCESS_CREATE_DATA))
}
