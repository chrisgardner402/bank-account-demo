package rest

import (
	"github.com/chrisgardner402/bank-account-demo/application/service"
	"github.com/labstack/echo/v4"
)

func ControlDeposit(e *echo.Echo) *echo.Route {
	return e.POST("/account/deposit", service.ServiceDeposit)
}
