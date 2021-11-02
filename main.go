package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/chrisgardner402/bank-account-demo/accounts"
	"github.com/labstack/echo/v4"
)

type OpenRequest struct {
	Owner string `json:"owner"`
}

type OpenResponse struct {
	Owner   string `json:"owner"`
	Balance int    `json:"balance"`
}

type AccountRequest struct {
	Owner string `json:"owner"`
}

type AccountResponse struct {
	Owner   string `json:"owner"`
	Balance int    `json:"balance"`
}

type Ledger map[string]accounts.Account

var myLedger Ledger

var (
	errNotFound    = errors.New("not found")
	errOwnerExists = errors.New("that owner already exists")
)

// Search for an account
func (l Ledger) Search(owner string) (accounts.Account, error) {
	account, exists := l[owner]
	if exists {
		return account, nil
	}
	return accounts.Account{}, errNotFound
}

// Add an account to the ledger
func (l Ledger) Add(account accounts.Account) error {
	_, err := l.Search(account.Owner())
	switch err {
	case nil:
		return errOwnerExists
	case errNotFound:
		l[account.Owner()] = account
	}
	return nil
}

func main() {
	myLedger = Ledger{}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/open", openAccount)
	e.POST("/account", getAccount)
	e.Logger.Fatal(e.Start(":1323"))
}

// open an account
func openAccount(c echo.Context) error {
	openRequest := new(OpenRequest)
	if err := c.Bind(&openRequest); err != nil {
		return err
	}
	account := accounts.NewAccount(openRequest.Owner)
	myLedger.Add(*account)

	fmt.Println(myLedger)

	openResponse := OpenResponse{Owner: account.Owner(), Balance: account.Balance()}
	return c.JSON(http.StatusCreated, openResponse)
}

// get an account
func getAccount(c echo.Context) error {
	accountRequest := new(AccountRequest)
	if err := c.Bind(&accountRequest); err != nil {
		return err
	}

	fmt.Println(accountRequest.Owner)
	fmt.Println(myLedger)

	account, err := myLedger.Search(accountRequest.Owner)
	// TODO error handling
	if err != nil {
		return err
	}

	fmt.Println(account.String())

	OpenResponse := OpenResponse{Owner: account.Owner(), Balance: account.Balance()}
	return c.JSON(http.StatusOK, OpenResponse)
}
