package mapper

import (
	"database/sql"

	"github.com/chrisgardner402/bank-account-demo/domain/entity"
	"github.com/chrisgardner402/bank-account-demo/domain/valueobject"
)

func SelectAccount(db *sql.DB, accountid string) (*sql.Rows, error) {
	row, err := db.Query("select userid, accountid, balance from account where accountid = ?", accountid)
	return row, err
}

func UpdateAccountForDeposit(db *sql.DB, amount *valueobject.Money, account *entity.Account) (sql.Result, error) {
	stmt, err := db.Prepare("update account set balance = balance + ? where accountid = ?")
	if err != nil {
		return nil, err
	}
	res, err := stmt.Exec(amount.Value(), account.Accountid())
	if err != nil {
		return nil, err
	}
	return res, nil
}

func UpdateAccountForWithdraw(db *sql.DB, amount *valueobject.Money, account *entity.Account) (sql.Result, error) {
	stmt, err := db.Prepare("update account set balance = balance - ? where accountid = ?")
	if err != nil {
		return nil, err
	}
	res, err := stmt.Exec(amount.Value(), account.Accountid())
	return res, err
}
