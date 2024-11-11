package controller

import (
	"skripsi/features/transaction/dto/request"
	"skripsi/features/transaction/interfaces"
	"skripsi/utils/helper"
	"skripsi/utils/jwt"

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

	return e.JSON(200, transaction)
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

	return e.JSON(200, helper.ResponseSuccess("notification processed successfully"))
}
