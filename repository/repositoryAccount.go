package repository

import (
	"log"

	"github.com/chrisgardner402/bank-account-demo/accounts"
)

// SearchAccount searches an account
func SearchAccount(owner string) (accounts.Account, error) {
	// execute a query
	row, err := db.Query("select owner, balance from account where owner=?", owner)
	if err != nil {
		return accounts.Account{}, err
	}
	// check row
	defer row.Close()
	if !row.Next() {
		return accounts.Account{}, errNotFound
	}
	// create account object
	var balance int
	row.Scan(&owner, &balance)
	account := *accounts.CreateAccount(owner, balance)
	return account, nil
}

// DepositAccount updates an account for deposit transaction
func DepositAccount(account *accounts.Account, amount int) error {
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
func WithdrawAccount(account *accounts.Account, amount int) error {
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
func updateAccount(account *accounts.Account) (int64, error) {
	// create a statement
	stmt, err := db.Prepare("update account set balance=? where owner=?")
	if err != nil {
		return 0, err
	}
	// execute a statement
	res, err := stmt.Exec(account.Balance(), account.Owner())
	if err != nil {
		return 0, err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return affect, nil
}
