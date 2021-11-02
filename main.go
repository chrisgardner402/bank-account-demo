package main

import (
	"fmt"
	"net/http"

	"github.com/chrisgardner402/bank-account-demo/accounts"
	"github.com/chrisgardner402/bank-account-demo/jsondata"
	"github.com/chrisgardner402/bank-account-demo/ledger"
	"github.com/labstack/echo/v4"
)

// temporary data for storage
var myLedger ledger.Ledger

func main() {
	// TODO remove
	myLedger = ledger.Ledger{}
	dali := accounts.NewAccount("dali")
	van := accounts.NewAccount("van")
	picasso := accounts.NewAccount("picasso")
	myLedger.Add(*dali)
	myLedger.Add(*van)
	myLedger.Add(*picasso)
	myLedger.Print()

	e := echo.New()
	e.GET("/health", handleHealthCheck)
	e.POST("/account/deposit", handleDeposit)
	e.POST("/account/withdraw", handleWithdraw)
	e.Logger.Fatal(e.Start(":1323"))
}

// handle health check
func handleHealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

// handle deposit api
func handleDeposit(c echo.Context) error {
	depositRequest := new(jsondata.DepositRequest)
	// binding
	if err := c.Bind(&depositRequest); err != nil {
		fmt.Println(err)
		return err
	}
	// search for an account
	account, err := myLedger.Search(depositRequest.Owner)
	// handle error
	if errRes := returnErrRe(err); err != nil {
		return c.JSON(http.StatusBadRequest, errRes)
	}
	// execute deposit
	err = account.Deposit(depositRequest.Amount)
	// handle error
	if errRes := returnErrRe(err); err != nil {
		return c.JSON(http.StatusBadRequest, errRes)
	}
	// update ledger
	myLedger[account.Owner()] = account
	myLedger.Print() // TODO remove
	// rendering
	depositResponse := jsondata.DepositResponse{}
	return c.JSON(http.StatusOK, depositResponse)
}

// handle withdraw api
func handleWithdraw(c echo.Context) error {
	withdrawRequest := new(jsondata.WithdrawRequest)
	// binding
	if err := c.Bind(withdrawRequest); err != nil {
		fmt.Println(err)
		return err
	}
	// search for an account
	account, err := myLedger.Search(withdrawRequest.Owner)
	// handle error
	if errRes := returnErrRe(err); err != nil {
		return c.JSON(http.StatusBadRequest, errRes)
	}
	// execute withdraw
	err = account.Withdraw(withdrawRequest.Amount)
	// handle error
	if errRes := returnErrRe(err); err != nil {
		return c.JSON(http.StatusBadRequest, errRes)
	}
	// update ledger
	myLedger[account.Owner()] = account
	myLedger.Print() // TODO remove
	// rendering
	withdrawResponse := jsondata.WithdrawResponse{}
	return c.JSON(http.StatusOK, withdrawResponse)
}

// return error response
func returnErrRe(err error) jsondata.ErrorResponse {
	var errorResponse jsondata.ErrorResponse
	if err != nil {
		errorResponse = jsondata.ErrorResponse{Message: err.Error()}
	}
	return errorResponse
}
