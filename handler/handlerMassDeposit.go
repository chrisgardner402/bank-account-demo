package handler

import (
	"log"
	"net/http"

	"github.com/chrisgardner402/bank-account-demo/accounts"
	"github.com/chrisgardner402/bank-account-demo/jsondata"
	"github.com/chrisgardner402/bank-account-demo/repository"
	"github.com/labstack/echo/v4"
)

func handleMassDeposit(c echo.Context) error {
	massDepositRequest := new(jsondata.MassDepositRequest)
	// binding
	if err := c.Bind(&massDepositRequest); err != nil {
		log.Println(err)
		return err
	}
	// search for mass accounts
	accountSlice, err := repository.SearchMassAccount(massDepositRequest.Owner)
	if isBad, errBadReq := handleBadRequest(err, c); isBad {
		return errBadReq
	}
	// execute mass deposit
	accountSlice, err = accounts.MassDeposit(*accountSlice, massDepositRequest.Amount)
	if isBad, errBadReq := handleBadRequest(err, c); isBad {
		return errBadReq
	}
	// update ledger
	err = repository.UpdateMassAccount(accountSlice)
	if isBad, errBadReq := handleIntlSrvErr(err, c); isBad {
		return errBadReq
	}
	// rendering
	massDepositResponse := jsondata.MassDepositReponse{}
	return c.JSON(http.StatusOK, massDepositResponse)
}
