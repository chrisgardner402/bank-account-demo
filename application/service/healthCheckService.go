package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ServiceHealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
