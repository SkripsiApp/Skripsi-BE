package controller

import (
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

	data := request.TransactionRequestToTransactionCore(input)
	data.UserId = id

	transaction, err := t.transactionService.CreateTransaction(data)
	if err != nil {
		return err
	}

	response := response.TransactionCoreToTransactionResponse(transaction)
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
