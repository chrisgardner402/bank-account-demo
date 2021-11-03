package handler

import (
	"log"
	"net/http"

	"github.com/chrisgardner402/bank-account-demo/jsondata"
	"github.com/chrisgardner402/bank-account-demo/repository"
	"github.com/labstack/echo/v4"
)

func handleWithdraw(c echo.Context) error {
	withdrawRequest := new(jsondata.WithdrawRequest)
	// binding
	if err := c.Bind(withdrawRequest); err != nil {
		log.Println(err)
		return err
	}
	// search for an account
	account, err := repository.SearchAccount(withdrawRequest.Owner)
	if isBad, errBadReq := handleBadRequest(err, c); isBad {
		return errBadReq
	}
	// before withdraw
	err = account.Withdraw(withdrawRequest.Amount)
	if isBad, errBadReq := handleBadRequest(err, c); isBad {
		return errBadReq
	}
	// update ledger
	err = repository.UpdateAccount(&account)
	if isBad, errBadReq := handleIntlSrvErr(err, c); isBad {
		return errBadReq
	}
	// rendering
	withdrawResponse := jsondata.WithdrawResponse{}
	return c.JSON(http.StatusOK, withdrawResponse)
}
