package controller

import (
	"skripsi/features/dashboard/interfaces"
	"skripsi/utils/constant"
	"skripsi/utils/helper"
	"skripsi/utils/jwt"

	"github.com/labstack/echo/v4"
)

type dashboardController struct {
	dashboardService interfaces.DashboardServiceInterface
}

func NewDashboardController(dashboardService interfaces.DashboardServiceInterface) *dashboardController {
	return &dashboardController{
		dashboardService: dashboardService,
	}
}

func (d *dashboardController) DashboardData(e echo.Context) error {
	_, role, err := jwt.ExtractToken(e)
	if err != nil {
		return err
	}

	if role != constant.ADMIN {
		return helper.ResponseError(401, constant.ERROR_AKSES_ROLE)
	}

	data, err := d.dashboardService.DashboardData()
	if err != nil {
		return err
	}

	return e.JSON(200, helper.ResponseSuccessWithData(constant.SUCCESS_GET_DATA, data))
}
