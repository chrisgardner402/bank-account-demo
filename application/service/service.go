package service

import (
	"log"
	"net/http"

	"github.com/chrisgardner402/bank-account-demo/application/response"
	"github.com/labstack/echo/v4"
)

func handleBadReq(err error, c echo.Context) (bool, error) {
	if errRes := returnErrRes(err); err != nil {
		log.Println(http.StatusBadRequest, errRes)
		return true, c.JSON(http.StatusBadRequest, errRes)
	}
	return false, nil
}

func handleIntlSrvErr(err error, c echo.Context) (bool, error) {
	if errRes := returnErrRes(err); err != nil {
		log.Println(http.StatusInternalServerError, errRes)
		return true, c.JSON(http.StatusInternalServerError, errRes)
	}
	return false, nil
}

func returnErrRes(err error) response.ErrorResponse {
	var errorResponse response.ErrorResponse
	if err != nil {
		errorResponse = response.ErrorResponse{Message: err.Error()}
	}
	return errorResponse
}
