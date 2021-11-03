package repository

import (
	"log"

	"github.com/chrisgardner402/bank-account-demo/accounts"
)

// SearchAccount searches an account
func SearchAccount(owner string) (accounts.Account, error) {
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

// UpdateAccount updates accounts for deposit and withdraw transactions
func UpdateAccount(account *accounts.Account) error {
	stmt, err := db.Prepare("update account set balance=? where owner=?")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(account.Balance(), account.Owner())
	if err != nil {
		return err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Println("affect:", affect)
	return nil
}
