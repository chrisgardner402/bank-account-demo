package main

import (
	"github.com/chrisgardner402/bank-account-demo/application/service"
	"github.com/chrisgardner402/bank-account-demo/infrastructure/persistence"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// ----- db -----
	defer persistence.CloseSqlite3()
	persistence.OpenSqlite3()

	// ----- architecture -----
	// infrastructure layer
	accountRepository := persistence.NewAccountPersistence()
	historyRepository := persistence.NewHistoryPersistence()
	// application layer
	healthCheckService := service.NewHealthCheckService()
	depositService := service.NewDepositService(accountRepository, historyRepository)
	withdrawService := service.NewWithdrawService(accountRepository, historyRepository)
	massDepositService := service.NewMassDepositService(accountRepository, historyRepository)

	// ----- web server -----
	e := echo.New()
	e.GET("/health", healthCheckService.ServiceHealthCheck)
	e.POST("/account/deposit", depositService.ServiceDeposit)
	e.POST("/account/withdraw", withdrawService.ServiceWithdraw)
	e.POST("/mass/deposit", massDepositService.ServiceMassDeposit)
	e.Logger.Fatal(e.Start(":1323"))
}
