package controller

import (
	"skripsi/features/voucher/dto/request"
	"skripsi/features/voucher/dto/response"
	"skripsi/features/voucher/interfaces"
	"skripsi/utils/constant"
	"skripsi/utils/helper"
	"skripsi/utils/jwt"
	"strconv"

	"github.com/labstack/echo/v4"
)

type voucherController struct {
	voucherService interfaces.VoucherServiceInterface
}

func NewVoucherController(voucherService interfaces.VoucherServiceInterface) *voucherController {
	return &voucherController{
		voucherService: voucherService,
	}
}

func (v *voucherController) Create(e echo.Context) error {
	id, _, errExtract := jwt.ExtractToken(e)
	if errExtract != nil {
		return errExtract
	}

	input := request.VoucherRequest{}
	input.AdminId = id

	errBind := e.Bind(&input)
	if errBind != nil {
		return helper.ResponseError(400, "invalid input data")
	}

	data := request.VoucherRequestToVoucherCore(input)
	_, err := v.voucherService.Create(data)
	if err != nil {
		return err
	}

	return e.JSON(200, helper.ResponseSuccess(constant.SUCCESS_CREATE_DATA))
}

func (v *voucherController) DeleteById(e echo.Context) error {
	id := e.Param("id")
	err := v.voucherService.DeleteById(id)
	if err != nil {
		return err
	}

	return e.JSON(200, helper.ResponseSuccess(constant.SUCCESS_DELETE_DATA))
}

func (v *voucherController) GetAll(e echo.Context) error {
	search := e.QueryParam("search")
	page, _ := strconv.Atoi(e.QueryParam("page"))
	limit, _ := strconv.Atoi(e.QueryParam("limit"))

	data, pageInfo, totalCount, err := v.voucherService.GetAll(search, page, limit)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		return helper.ResponseError(200, "data belum tersedia")
	}

	response := response.ListVoucherCoreToListVoucherResponse(data)

	return e.JSON(200, helper.ResponseSuccessWithPagnationAndCount(constant.SUCCESS_GET_DATA, response, pageInfo, totalCount))
}

func (v *voucherController) GetById(e echo.Context) error {
	id := e.Param("id")

	data, err := v.voucherService.GetById(id)
	if err != nil {
		return err
	}

	response := response.VoucherCoreToVoucherResponse(data)

	return e.JSON(200, helper.ResponseSuccessWithData(constant.SUCCESS_GET_DATA, response))
}

func (v *voucherController) UpdateById(e echo.Context) error {
	id := e.Param("id")

	input := request.VoucherRequest{}
	errBind := e.Bind(&input)
	if errBind != nil {
		return helper.ResponseError(400, "invalid input data")
	}

	data := request.VoucherRequestToVoucherCore(input)
	err := v.voucherService.UpdateById(id, data)
	if err != nil {
		return err
	}

	return e.JSON(200, helper.ResponseSuccess(constant.SUCCESS_UPDATE_DATA))
}