package mapper

import (
	"database/sql"

	"github.com/chrisgardner402/bank-account-demo/domain/entity"
)

func InsertHistory(db *sql.DB, history *entity.History) (sql.Result, error) {
	stmt, err := db.Prepare("insert into history(accountid, historyid, deposit, withdraw) values(?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	res, err := stmt.Exec(history.Accountid(), history.Historyid(), history.Deposit(), history.Withdraw())
	return res, err
}
