package handler

import (
	"log"
	"net/http"

	"github.com/chrisgardner402/bank-account-demo/jsondata"
	"github.com/labstack/echo/v4"
)

// HandleHealthCheck handles health check
func HandleHealthCheck() echo.HandlerFunc {
	return handleHealthCheck
}

// HandleDeposit handles deposit api
func HandleDeposit() echo.HandlerFunc {
	return handleDeposit
}

// HandleWithdraw handles withdraw api
func HandleWithdraw() echo.HandlerFunc {
	return handleWithdraw
}

// HandleMassDeposit handles mass deposit api
func HandleMassDeposit() echo.HandlerFunc {
	return handleMassDeposit
}

// handle bad request
func handleBadRequest(err error, c echo.Context) (bool, error) {
	if errRes := returnErrRes(err); err != nil {
		log.Println(http.StatusBadRequest, errRes)
		return true, c.JSON(http.StatusBadRequest, errRes)
	}
	return false, nil
}

// handle internal server error
func handleIntlSrvErr(err error, c echo.Context) (bool, error) {
	if errRes := returnErrRes(err); err != nil {
		log.Println(http.StatusInternalServerError, errRes)
		return true, c.JSON(http.StatusInternalServerError, errRes)
	}
	return false, nil
}

// return error response
func returnErrRes(err error) jsondata.ErrorResponse {
	var errorResponse jsondata.ErrorResponse
	if err != nil {
		errorResponse = jsondata.ErrorResponse{Message: err.Error()}
	}
	return errorResponse
}
