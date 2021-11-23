package service

import (
	"log"
	"net/http"

	"github.com/chrisgardner402/bank-account-demo/application/request"
	"github.com/chrisgardner402/bank-account-demo/application/response"
	"github.com/chrisgardner402/bank-account-demo/domain/factory"
	"github.com/chrisgardner402/bank-account-demo/domain/repository"

	"github.com/labstack/echo/v4"
)

func NewWithdrawService(ar repository.AccountRepository, hr repository.HistoryRepository) WithdrawService {
	return &withdrawService{
		accountRepository: ar,
		historyRepository: hr}
}

type WithdrawService interface {
	ServiceWithdraw(c echo.Context) error
}

type withdrawService struct {
	accountRepository repository.AccountRepository
	historyRepository repository.HistoryRepository
}

func (ws withdrawService) ServiceWithdraw(c echo.Context) error {
	// data binding
	withdrawRequest := new(request.WithdrawRequest)
	if err := c.Bind(withdrawRequest); err != nil {
		log.Println(err)
		return err
	}

	// ----- business logic start -----
	// create account and search for
	account := factory.CreateAccountFromAccountid(withdrawRequest.Accountid)
	accountPersist, err := ws.accountRepository.SearchAccount(account)
	if isBad, errBadReq := handleBadReq(err, c); isBad {
		return errBadReq
	}
	// create withdraw and validate request
	withdraw := factory.CreateWithdraw(accountPersist, withdrawRequest.Amount)
	err = withdraw.ValidateWithdraw()
	if isBad, errBadReq := handleBadReq(err, c); isBad {
		return errBadReq
	}
	// execute withdraw
	err = ws.accountRepository.WithdrawAccount(withdraw)
	if isBad, errBadReq := handleIntlSrvErr(err, c); isBad {
		return errBadReq
	}
	// create history and record
	history, err := factory.CreateWithdrawHistory(withdraw)
	if isBad, errBadReq := handleIntlSrvErr(err, c); isBad {
		return errBadReq
	}
	err = ws.historyRepository.RecordHistory(history)
	if isBad, errBadReq := handleIntlSrvErr(err, c); isBad {
		return errBadReq
	}
	// ----- business logic end -----

	// data rendering
	withdrawResponse := new(response.WithdrawResponse)
	return c.JSON(http.StatusOK, withdrawResponse)
}
