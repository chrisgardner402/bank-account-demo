package repository

import (
	"database/sql"
	"log"

	"github.com/chrisgardner402/bank-account-demo/model"
)

// SaveDepositHis saves deposit history
func SaveDepositHis(account *model.Account, amount int) error {
	// create a history
	history, err := model.CreateDepositHis(account)
	if err != nil {
		return err
	}
	// create a statement
	stmt, err := db.Prepare("insert into history(accountid, historyid, deposit) values(?, ?, ?)")
	if err != nil {
		return err
	}
	// execute a statement
	affect, err := execTran(stmt, history, amount)
	if err != nil {
		return err
	}
	log.Println("save deposit history affect:", affect)
	return nil
}

// SaveWithdrawHis saves withdraw history
func SaveWithdrawHis(account *model.Account, amount int) error {
	// create a history
	history, err := model.CreateWithdrawHis(account)
	if err != nil {
		return err
	}
	// create a statement
	stmt, err := db.Prepare("insert into history(accountid, historyid, withdraw) values(?, ?, ?)")
	if err != nil {
		return err
	}
	// execute a statement
	affect, err := execTran(stmt, history, amount)
	if err != nil {
		return err
	}
	log.Println("save withdraw history affect:", affect)
	return nil
}

func execTran(stmt *sql.Stmt, history *model.History, amount int) (int64, error) {
	res, err := stmt.Exec(history.Accountid(), history.Historyid(), amount)
	if err != nil {
		return 0, err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return affect, nil
}
