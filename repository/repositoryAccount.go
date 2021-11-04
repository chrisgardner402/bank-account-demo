package repository

import (
	"log"

	"github.com/chrisgardner402/bank-account-demo/model"
)

// SearchAccount searches an account
func SearchAccount(accountid string) (model.Account, error) {
	// execute a query
	row, err := db.Query("select accountid, balance from account where accountid=?", accountid)
	if err != nil {
		return model.Account{}, err
	}
	// check row
	defer row.Close()
	if !row.Next() {
		return model.Account{}, errNotFound
	}
	// create account object
	var balance int
	row.Scan(&accountid, &balance)
	account := *model.CreateAccount(accountid, balance)
	return account, nil
}

// DepositAccount updates an account for deposit transaction
func DepositAccount(account *model.Account, amount int) error {
	// before deposit
	err := account.Deposit(amount)
	if err != nil {
		return err
	}
	// update the account
	affect, err := updateAccount(account)
	if err != nil {
		return err
	}
	log.Println("deposit account affect:", affect)
	return nil
}

// WithdrawAccount updates an account for withdraw transaction
func WithdrawAccount(account *model.Account, amount int) error {
	// before withdraw
	err := account.Withdraw(amount)
	if err != nil {
		return err
	}
	// update the account
	affect, err := updateAccount(account)
	if err != nil {
		return err
	}
	log.Println("withdraw account affect:", affect)
	return nil
}

// update an account
func updateAccount(account *model.Account) (int64, error) {
	// create a statement
	stmt, err := db.Prepare("update account set balance=? where accountid=?")
	if err != nil {
		return 0, err
	}
	// execute a statement
	res, err := stmt.Exec(account.Balance(), account.Accountid())
	if err != nil {
		return 0, err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return affect, nil
}
