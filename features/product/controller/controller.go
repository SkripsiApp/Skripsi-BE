package controller

import (
	"skripsi/features/product/dto/request"
	"skripsi/features/product/interfaces"
	"skripsi/utils/constant"
	"skripsi/utils/helper"
	"skripsi/utils/jwt"

	"github.com/labstack/echo/v4"
)

type productController struct {
	productService interfaces.ProductServiceInterface
}

func NewProductController(productService interfaces.ProductServiceInterface) *productController {
	return &productController{
		productService: productService,
	}
}

func (p *productController) Create(e echo.Context) error {
	_, role, errExtract := jwt.ExtractToken(e)
	if role != "admin" {
		return helper.ResponseError(401, constant.ERROR_AKSES_ROLE)
	}

	if errExtract != nil {
		return errExtract
	}

	input := request.ProductRequest{}

	errBind := e.Bind(&input)
	if errBind != nil {
		return helper.ResponseError(400, "invalid input data")
	}

	data := request.ProductRequestToProductCore(input)

	file, err := e.FormFile("image")
	if err != nil {
		return helper.ResponseError(400, "gambar tidak ditemukan")
	}
	
	_, err = p.productService.Create(file, data)
	if err != nil {
		return err
	}

	return e.JSON(200, helper.ResponseSuccess(constant.SUCCESS_CREATE_DATA))
}
