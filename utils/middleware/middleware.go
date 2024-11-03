package middleware

import (
	"net/http"
	"skripsi/utils/helper"
	"strings"

	"github.com/labstack/echo/v4"
)

func ErrorMiddleware(err error, c echo.Context) {
	var status int
	var message string

	if strings.Contains(err.Error(), "token is malformed") || strings.Contains(err.Error(), "invalid or expired jwt") {
		status = http.StatusUnauthorized
		message = "invalid token"
	} else if resErr, ok := err.(*helper.ErrorResponse); ok {
		status = resErr.StatusCode
		message = resErr.Message
	} else {
		// Untuk error lainnya, berikan response 500 dan log error
		status = http.StatusInternalServerError
		message = "internal server error"
		c.Logger().Error(err)
	}

	c.JSON(status, helper.ResponseErrorMiddleware(message))
}
