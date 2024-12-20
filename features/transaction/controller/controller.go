package controller

import (
	"fmt"
	"skripsi/features/transaction/dto/request"
	"skripsi/features/transaction/dto/response"
	"skripsi/features/transaction/interfaces"
	"skripsi/utils/constant"
	"skripsi/utils/helper"
	"skripsi/utils/jwt"
	"strconv"

	"github.com/labstack/echo/v4"
)

type transactionController struct {
	transactionService interfaces.TransactionServiceInterface
}

func NewTransactionController(transactionService interfaces.TransactionServiceInterface) *transactionController {
	return &transactionController{
		transactionService: transactionService,
	}
}

func (t *transactionController) CreateTransaction(e echo.Context) error {
	id, _, errExtract := jwt.ExtractToken(e)
	if errExtract != nil {
		return errExtract
	}

	input := request.TransactionRequest{}

	errBind := e.Bind(&input)
	if errBind != nil {
		return helper.ResponseError(400, "invalid input data")
	}

	fmt.Println("data input", input)

	data := request.TransactionRequestToTransactionCore(input)
	data.UserId = id

	transaction, snap, err := t.transactionService.CreateTransaction(data)
	if err != nil {
		return err
	}

	response := response.TransactionCoreToTransactionRequestResponse(transaction, snap)

	return e.JSON(200, helper.ResponseSuccessWithData("transaksi berhasil dibuat", response))
}

func (t *transactionController) HandleMidtransNotification(e echo.Context) error {
	var notification helper.MidtransNotificationPayload
	errBind := e.Bind(&notification)
	if errBind != nil {
		return helper.ResponseError(400, "invalid input data")
	}

	err := t.transactionService.HandleMidtransNotification(notification)
	if err != nil {
		return err
	}

	return e.JSON(200, helper.ResponseSuccess("notification berhasil dihandle"))
}

func (t *transactionController) GetAllTransaction(e echo.Context) error {
	_, role, errExtract := jwt.ExtractToken(e)
	if role != constant.ADMIN {
		return helper.ResponseError(401, constant.ERROR_AKSES_ROLE)
	}

	if errExtract != nil {
		return errExtract
	}

	search := e.QueryParam("search")
	page, _ := strconv.Atoi(e.QueryParam("page"))
	limit, _ := strconv.Atoi(e.QueryParam("limit"))

	data, pageInfo, totalCount, err := t.transactionService.GetAllTransaction(search, page, limit)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		return helper.ResponseError(200, "data belum tersedia")
	}

	response := response.ListTransactionCoreToTransactionResponse(data)

	return e.JSON(200, helper.ResponseSuccessWithPagnationAndCount(constant.SUCCESS_GET_DATA, response, pageInfo, totalCount))
}

func (t *transactionController) GetTransactionById(e echo.Context) error {
	_, role, errExtract := jwt.ExtractToken(e)
	if role != constant.ADMIN {
		return helper.ResponseError(401, constant.ERROR_AKSES_ROLE)
	}

	if errExtract != nil {
		return errExtract
	}

	id := e.Param("id")

	data, err := t.transactionService.GetTransactionById(id)
	if err != nil {
		return err
	}

	response := response.TransactionCoreToTransactionResponse(data)

	return e.JSON(200, helper.ResponseSuccessWithData(constant.SUCCESS_GET_DATA, response))
}

func (t *transactionController) GetAllTransactionByUserId (e echo.Context) error {
	id, _, errExtract := jwt.ExtractToken(e)
	if errExtract != nil {
		return errExtract
	}

	search := e.QueryParam("search")
	page, _ := strconv.Atoi(e.QueryParam("page"))
	limit, _ := strconv.Atoi(e.QueryParam("limit"))

	data, pageInfo, totalCount, err := t.transactionService.GetAllTransactionByUserId(id, search, page, limit)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		return helper.ResponseError(200, "data belum tersedia")
	}

	response := response.ListTransactionCoreToTransactionResponse(data)

	return e.JSON(200, helper.ResponseSuccessWithPagnationAndCount(constant.SUCCESS_GET_DATA, response, pageInfo, totalCount))
}

func (t *transactionController) UpdateNoReceipt(e echo.Context) error {
	_, role, errExtract := jwt.ExtractToken(e)
	if role != constant.ADMIN {
		return helper.ResponseError(401, constant.ERROR_AKSES_ROLE)
	}

	if errExtract != nil {
		return errExtract
	}

	id := e.Param("id")

	input := request.NoReceiptRequest{}
	errBind := e.Bind(&input)
	if errBind != nil {
		return helper.ResponseError(400, "invalid input data")
	}

	err := t.transactionService.UpdateNoReceiptTransactionById(id, input.NoReceipt)
	if err != nil {
		return err
	}

	return e.JSON(200, helper.ResponseSuccess("no resi berhasil diupdate"))
}

func (t *transactionController) UpdateStatusTransactionUserById(e echo.Context) error {
	id, _, errExtract := jwt.ExtractToken(e)
	if errExtract != nil {
		return errExtract
	}

	transactionId := e.Param("id")

	input := request.UpdateStatusRequest{}
	errBind := e.Bind(&input)
	if errBind != nil {
		return helper.ResponseError(400, "invalid input data")
	}

	err := t.transactionService.UpdateStatusTransactionUserById(id, transactionId, input.Status)
	if err != nil {
		return err
	}

	return e.JSON(200, helper.ResponseSuccess("status transaksi berhasil diupdate"))
}

