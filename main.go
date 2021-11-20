package main

import (
	"github.com/chrisgardner402/bank-account-demo/application/rest"
	"github.com/chrisgardner402/bank-account-demo/infra/persistence/repository"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// db
	defer repository.Close()
	repository.Open()

	// echo
	e := echo.New()
	rest.ControlHealthCheck(e)
	rest.ControlDeposit(e)
	rest.ControlWithdraw(e)
	rest.ControlMassDeposit(e)
	e.Logger.Fatal(e.Start(":1323"))
}
