package service

import (
	"log"
	"net/http"

	"github.com/chrisgardner402/bank-account-demo/application/response"
	"github.com/labstack/echo/v4"
)

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
func returnErrRes(err error) response.ErrorResponse {
	var errorResponse response.ErrorResponse
	if err != nil {
		errorResponse = response.ErrorResponse{Message: err.Error()}
	}
	return errorResponse
}
