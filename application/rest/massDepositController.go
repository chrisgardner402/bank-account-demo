package rest

import (
	"github.com/chrisgardner402/bank-account-demo/application/service"
	"github.com/labstack/echo/v4"
)

func ControlMassDeposit(e *echo.Echo) *echo.Route {
	return e.POST("/mass/deposit", service.ServiceMassDeposit)
}
