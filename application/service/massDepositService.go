package service

import (
	"log"
	"net/http"

	"github.com/chrisgardner402/bank-account-demo/application/request"
	"github.com/chrisgardner402/bank-account-demo/application/response"
	"github.com/chrisgardner402/bank-account-demo/domain/validate"
	"github.com/chrisgardner402/bank-account-demo/infra/persistence/repository"
	"github.com/labstack/echo/v4"
)

func ServiceMassDeposit(c echo.Context) error {
	massDepositRequest := new(request.MassDepositRequest)
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
	massDepositResponse := response.MassDepositReponse{}
	return c.JSON(http.StatusOK, massDepositResponse)
}
