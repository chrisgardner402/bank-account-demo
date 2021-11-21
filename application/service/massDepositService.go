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

func ServiceMassDeposit(c echo.Context) error {
	// data binding
	massDepositRequest := new(request.MassDepositRequest)
	if err := c.Bind(&massDepositRequest); err != nil {
		log.Println(err)
		return err
	}

	// ----- business logic start -----
	// create account slice and search for
	accountSlice := factory.CreateAccountSliceFromAccountid(massDepositRequest.Accountidlist)
	accountSlicePersist, err := repository.SearchMassAccount(accountSlice)
	if isBad, errBadReq := handleBadReq(err, c); isBad {
		return errBadReq
	}
	// create deposit slice and validate request
	depositSlice := factory.CreateDepositSlice(accountSlicePersist, massDepositRequest.Amount)
	for _, deposit := range *depositSlice {
		err := deposit.ValidateDeposit()
		if isBad, errBadReq := handleBadReq(err, c); isBad {
			return errBadReq
		}
	}
	// execute deposit slice
	err = repository.DepositMassAccount(depositSlice)
	if isBad, errBadReq := handleIntlSrvErr(err, c); isBad {
		return errBadReq
	}
	// create history slice and record
	historySlice, err := factory.CreateDepositHistorySlice(depositSlice)
	if isBad, errBadReq := handleIntlSrvErr(err, c); isBad {
		return errBadReq
	}
	err = repository.RecordMassHistory(historySlice)
	if isBad, errBadReq := handleIntlSrvErr(err, c); isBad {
		return errBadReq
	}
	// ----- business logic end -----

	// data rendering
	massDepositResponse := new(response.MassDepositReponse)
	return c.JSON(http.StatusOK, massDepositResponse)
}
