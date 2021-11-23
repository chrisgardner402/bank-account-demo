package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewHealthCheckService() HealthCheckService {
	return &healthCheckService{}
}

type HealthCheckService interface {
	ServiceHealthCheck(c echo.Context) error
}

type healthCheckService struct{}

func (hcs healthCheckService) ServiceHealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
