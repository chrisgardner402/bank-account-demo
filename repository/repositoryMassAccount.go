package repository

import (
	"github.com/chrisgardner402/bank-account-demo/model"
)

// SearchMassAccount searches accounts
func SearchMassAccount(accountidlist []string) (*[]model.Account, error) {
	var accountlist []model.Account
	for _, accountid := range accountidlist {
		account, err := SearchAccount(accountid)
		if err != nil {
			return nil, err
		}
		accountlist = append(accountlist, account)
	}
	return &accountlist, nil
}

// DepositMassAccount updates accounts for deposit transactions
func DepositMassAccount(accountlist *[]model.Account, amount int) error {
	for _, account := range *accountlist {
		err := DepositAccount(&account, amount)
		if err != nil {
			return err
		}
	}
	return nil
}
