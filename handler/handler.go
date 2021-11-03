package handler

import (
	"log"
	"net/http"

	"github.com/chrisgardner402/bank-account-demo/accounts"
	"github.com/chrisgardner402/bank-account-demo/jsondata"
	"github.com/chrisgardner402/bank-account-demo/repository"
	"github.com/labstack/echo/v4"
)

// HandleHealthCheck handles health check
func HandleHealthCheck() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	}
}

// HandleDeposit handles deposit api
func HandleDeposit() echo.HandlerFunc {
	return func(c echo.Context) error {
		depositRequest := new(jsondata.DepositRequest)
		// binding
		if err := c.Bind(&depositRequest); err != nil {
			log.Println(err)
			return err
		}
		// search for an account
		account, err := repository.SearchAccount(depositRequest.Owner)
		if isBad, errBadReq := handleBadRequest(err, c); isBad {
			return errBadReq
		}
		// before deposit
		err = account.Deposit(depositRequest.Amount)
		if isBad, errBadReq := handleBadRequest(err, c); isBad {
			return errBadReq
		}
		// update ledger
		err = repository.UpdateAccount(&account)
		if isBad, errBadReq := handleIntlSrvErr(err, c); isBad {
			return errBadReq
		}
		// rendering
		depositResponse := jsondata.DepositResponse{}
		return c.JSON(http.StatusOK, depositResponse)
	}
}

// HandleWithdraw handles withdraw api
func HandleWithdraw() echo.HandlerFunc {
	return func(c echo.Context) error {
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
}

// HandleMassDeposit handles mass deposit api
func HandleMassDeposit() echo.HandlerFunc {
	return func(c echo.Context) error {
		massDepositRequest := new(jsondata.MassDepositRequest)
		// binding
		if err := c.Bind(&massDepositRequest); err != nil {
			log.Println(err)
			return err
		}
		// search for an account
		accountSlice := []accounts.Account{}
		for _, owner := range massDepositRequest.Owner {
			account, err := repository.SearchAccount(owner)
			if isBad, errBadReq := handleBadRequest(err, c); isBad {
				return errBadReq
			}
			accountSlice = append(accountSlice, account)
		}
		// execute mass deposit
		accountSlice, err := accounts.MassDeposit(accountSlice, massDepositRequest.Amount)
		if isBad, errBadReq := handleBadRequest(err, c); isBad {
			return errBadReq
		}
		// update ledger
		for _, account := range accountSlice {
			err = repository.UpdateAccount(&account)
			if isBad, errBadReq := handleIntlSrvErr(err, c); isBad {
				return errBadReq
			}
		}
		// rendering
		massDepositResponse := jsondata.MassDepositReponse{}
		return c.JSON(http.StatusOK, massDepositResponse)
	}
}

// handle bad request
func handleBadRequest(err error, c echo.Context) (bool, error) {
	if errRes := returnErrRe(err); err != nil {
		log.Println(http.StatusBadRequest, errRes)
		return true, c.JSON(http.StatusBadRequest, errRes)
	}
	return false, nil
}

// handle internal server error
func handleIntlSrvErr(err error, c echo.Context) (bool, error) {
	if errRes := returnErrRe(err); err != nil {
		log.Println(http.StatusInternalServerError, errRes)
		return true, c.JSON(http.StatusInternalServerError, errRes)
	}
	return false, nil
}

// return error response
func returnErrRe(err error) jsondata.ErrorResponse {
	var errorResponse jsondata.ErrorResponse
	if err != nil {
		errorResponse = jsondata.ErrorResponse{Message: err.Error()}
	}
	return errorResponse
}
