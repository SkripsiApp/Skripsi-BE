package controller

import (
	"fmt"
	"log"
	"net/http"
	"skripsi/features/product/dto/request"
	"skripsi/features/product/dto/response"
	"skripsi/features/product/interfaces"
	"skripsi/utils/constant"
	"skripsi/utils/helper"
	"skripsi/utils/jwt"
	"strconv"

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

	// Log the input to check if product_size is bound correctly
	log.Printf("ProductRequest before binding product_size: %+v", input)

	// Manually bind product_size
	productSizes := []request.ProductSizeRequest{}
	for i := 0; ; i++ {
		size := e.FormValue(fmt.Sprintf("product_size[%d].size", i))
		if size == "" {
			break
		}
		stock, err := strconv.Atoi(e.FormValue(fmt.Sprintf("product_size[%d].stock", i)))
		if err != nil {
			return helper.ResponseError(400, "invalid nilai stock")
		}
		productSizes = append(productSizes, request.ProductSizeRequest{
			Size:  size,
			Stock: stock,
		})
	}
	input.ProductSize = productSizes

	// Log the input to check if product_size is bound correctly
	log.Printf("ProductRequest after binding product_size: %+v", input)

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

func (p *productController) GetAll(e echo.Context) error {
	search := e.QueryParam("search")
	page, _ := strconv.Atoi(e.QueryParam("page"))
	limit, _ := strconv.Atoi(e.QueryParam("limit"))

	data, pageInfo, totalCount, err := p.productService.GetAll(search, page, limit)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		return helper.ResponseError(200, "data belum tersedia")
	}

	response := response.ListProductCoreToListProductResponse(data)

	return e.JSON(200, helper.ResponseSuccessWithPagnationAndCount(constant.SUCCESS_GET_DATA, response, pageInfo, totalCount))
}

func (p *productController) GetById(e echo.Context) error {
	id := e.Param("id")

	data, err := p.productService.GetById(id)
	if err != nil {
		return err
	}

	response := response.ProductCoreToProductResponse(data)

	return e.JSON(200, helper.ResponseSuccessWithData(constant.SUCCESS_GET_DATA, response))
}

func (p *productController) UpdateById(e echo.Context) error {
	_, role, errExtract := jwt.ExtractToken(e)
    if role != "admin" {
        return helper.ResponseError(401, constant.ERROR_AKSES_ROLE)
    }

    if errExtract != nil {
        return errExtract
    }

    id := e.Param("id")

    input := request.ProductRequest{}

    errBind := e.Bind(&input)
    if errBind != nil {
        return helper.ResponseError(400, "invalid input data")
    }

    // Manually bind product_size
    productSizes := []request.ProductSizeRequest{}
    for i := 0; ; i++ {
        size := e.FormValue(fmt.Sprintf("product_size[%d].size", i))
        if size == "" {
            break
        }
        stock, err := strconv.Atoi(e.FormValue(fmt.Sprintf("product_size[%d].stock", i)))
        if err != nil {
            return helper.ResponseError(400, "invalid nilai stock")
        }
        productSizes = append(productSizes, request.ProductSizeRequest{
            Size:  size,
            Stock: stock,
        })
    }
    input.ProductSize = productSizes

    data := request.ProductRequestToProductCore(input)

    file, err := e.FormFile("image")
    if err != nil && err != http.ErrMissingFile {
        return helper.ResponseError(400, "gambar tidak ditemukan")
    }

    err = p.productService.UpdateById(id, file, data)
    if err != nil {
        return err
    }

    return e.JSON(200, helper.ResponseSuccess(constant.SUCCESS_UPDATE_DATA))
}

func (p *productController) DeleteById(e echo.Context) error {
	_, role, errExtract := jwt.ExtractToken(e)
	if role != "admin" {
		return helper.ResponseError(401, constant.ERROR_AKSES_ROLE)
	}

	if errExtract != nil {
		return errExtract
	}

	id := e.Param("id")

	err := p.productService.DeleteById(id)
	if err != nil {
		return err
	}

	return e.JSON(200, helper.ResponseSuccess(constant.SUCCESS_DELETE_DATA))
}
