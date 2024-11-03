package middleware

import (
	"net/http"
	"skripsi/utils/helper"

	"github.com/labstack/echo/v4"
)

func ErrorMiddleware(err error, c echo.Context) {
	var status int
	var message string

	if resErr, ok := err.(*helper.ErrorResponse); ok {
		status = resErr.StatusCode
		message = resErr.Message
	} else {
		status = http.StatusInternalServerError
		message = "Internal Server Error"
		c.Logger().Error(err)
	}

	c.JSON(status, helper.ResponseErrorMiddleware(message))
}
