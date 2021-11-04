package handler

import (
	"log"
	"net/http"

	"github.com/chrisgardner402/bank-account-demo/jsondata"
	"github.com/chrisgardner402/bank-account-demo/repository"
	"github.com/chrisgardner402/bank-account-demo/validate"
	"github.com/labstack/echo/v4"
)

func handleMassDeposit(c echo.Context) error {
	massDepositRequest := new(jsondata.MassDepositRequest)
	// binding
	if err := c.Bind(&massDepositRequest); err != nil {
		log.Println(err)
		return err
	}
	// validate request
	err := validate.ValidateDeposit(massDepositRequest.Amount)
	if isBad, errBadReq := handleBadRequest(err, c); isBad {
		return errBadReq
	}
	// search for mass account
	accountlist, err := repository.SearchMassAccount(massDepositRequest.Accountidlist)
	if isBad, errBadReq := handleBadRequest(err, c); isBad {
		return errBadReq
	}
	// execute mass deposit
	err = repository.DepositMassAccount(accountlist, massDepositRequest.Amount)
	if isBad, errBadReq := handleIntlSrvErr(err, c); isBad {
		return errBadReq
	}
	// save history
	err = repository.SaveMassDepositHis(accountlist, massDepositRequest.Amount)
	if isBad, errBadReq := handleIntlSrvErr(err, c); isBad {
		return errBadReq
	}
	// rendering
	massDepositResponse := jsondata.MassDepositReponse{}
	return c.JSON(http.StatusOK, massDepositResponse)
}
