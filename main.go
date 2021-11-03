package main

import (
	"github.com/chrisgardner402/bank-account-demo/handler"
	"github.com/chrisgardner402/bank-account-demo/repository"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// db
	defer repository.Close()
	repository.Open()

	// echo
	e := echo.New()
	e.GET("/health", handler.HandleHealthCheck())
	e.POST("/account/deposit", handler.HandleDeposit())
	e.POST("/account/withdraw", handler.HandleWithdraw())
	e.POST("/mass/deposit", handler.HandleMassDeposit())
	e.Logger.Fatal(e.Start(":1323"))
}
