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

func ServiceDeposit(c echo.Context) error {
	depositRequest := new(request.DepositRequest)
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
	// save history
	err = repository.SaveDepositHis(&account, depositRequest.Amount)
	if isBad, errBadReq := handleIntlSrvErr(err, c); isBad {
		return errBadReq
	}
	// rendering
	depositResponse := response.DepositResponse{}
	return c.JSON(http.StatusOK, depositResponse)
}
