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

func ServiceWithdraw(c echo.Context) error {
	withdrawRequest := new(request.WithdrawRequest)
	// binding
	if err := c.Bind(withdrawRequest); err != nil {
		log.Println(err)
		return err
	}
	// search for an account
	account, err := repository.SearchAccount(withdrawRequest.Accountid)
	if isBad, errBadReq := handleBadRequest(err, c); isBad {
		return errBadReq
	}
	// validate request
	err = validate.ValidateWithdraw(account, withdrawRequest.Amount)
	if isBad, errBadReq := handleBadRequest(err, c); isBad {
		return errBadReq
	}
	// execute withdraw
	err = repository.WithdrawAccount(&account, withdrawRequest.Amount)
	if isBad, errBadReq := handleIntlSrvErr(err, c); isBad {
		return errBadReq
	}
	// save history
	err = repository.SaveWithdrawHis(&account, withdrawRequest.Amount)
	if isBad, errBadReq := handleIntlSrvErr(err, c); isBad {
		return errBadReq
	}
	// rendering
	withdrawResponse := response.WithdrawResponse{}
	return c.JSON(http.StatusOK, withdrawResponse)
}
