package service

import (
	"log"
	"net/http"

	"github.com/chrisgardner402/bank-account-demo/application/request"
	"github.com/chrisgardner402/bank-account-demo/application/response"
	"github.com/chrisgardner402/bank-account-demo/domain/factory"
	"github.com/chrisgardner402/bank-account-demo/infra/persistence/repository"
	"github.com/labstack/echo/v4"
)

func ServiceDeposit(c echo.Context) error {
	// data binding
	depositRequest := new(request.DepositRequest)
	if err := c.Bind(&depositRequest); err != nil {
		log.Println(err)
		return err
	}

	// ----- business logic start -----
	// create account and search for
	account := factory.CreateAccountFromAccountid(depositRequest.Accountid)
	accountPersist, err := repository.SearchAccount(account)
	if isBad, errBadReq := handleBadReq(err, c); isBad {
		return errBadReq
	}
	// create deposit and validate request
	deposit := factory.CreateDeposit(accountPersist, depositRequest.Amount)
	err = deposit.ValidateDeposit()
	if isBad, errBadReq := handleBadReq(err, c); isBad {
		return errBadReq
	}
	// execute deposit
	err = repository.DepositAccount(deposit)
	if isBad, errBadReq := handleIntlSrvErr(err, c); isBad {
		return errBadReq
	}
	// create history and record
	history, err := factory.CreateDepositHistory(deposit)
	if isBad, errBadReq := handleIntlSrvErr(err, c); isBad {
		return errBadReq
	}
	err = repository.RecordHistory(history)
	if isBad, errBadReq := handleIntlSrvErr(err, c); isBad {
		return errBadReq
	}
	// ----- business logic end -----

	// data rendering
	depositResponse := new(response.DepositResponse)
	return c.JSON(http.StatusOK, depositResponse)
}
