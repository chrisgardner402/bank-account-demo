package handler

import (
	"log"
	"net/http"

	"github.com/chrisgardner402/bank-account-demo/jsondata"
	"github.com/chrisgardner402/bank-account-demo/repository"
	"github.com/chrisgardner402/bank-account-demo/validate"
	"github.com/labstack/echo/v4"
)

func handleDeposit(c echo.Context) error {
	depositRequest := new(jsondata.DepositRequest)
	// binding
	if err := c.Bind(&depositRequest); err != nil {
		log.Println(err)
		return err
	}
	// validate request
	err := validate.ValidateDeposit(depositRequest.Amount)
	if isBad, errBadReq := handleBadRequest(err, c); isBad {
		return errBadReq
	}
	// search for an account
	account, err := repository.SearchAccount(depositRequest.Accountid)
	if isBad, errBadReq := handleBadRequest(err, c); isBad {
		return errBadReq
	}
	// execute deposit
	err = repository.DepositAccount(&account, depositRequest.Amount)
	if isBad, errBadReq := handleIntlSrvErr(err, c); isBad {
		return errBadReq
	}
	// rendering
	depositResponse := jsondata.DepositResponse{}
	return c.JSON(http.StatusOK, depositResponse)
}
