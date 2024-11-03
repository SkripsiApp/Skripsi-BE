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

	if strings.Contains(err.Error(), "missing or malformed jwt") || strings.Contains(err.Error(), "invalid or expired jwt") {
		status = http.StatusUnauthorized
		message = "invalid token"
	} else if resErr, ok := err.(*helper.ErrorResponse); ok {
		status = resErr.StatusCode
		message = resErr.Message
	} else {
		status = http.StatusInternalServerError
		message = "internal server error"
		c.Logger().Error(err)
	}

	c.JSON(status, helper.ResponseErrorMiddleware(message))
}
