package repository

import (
	"log"

	"github.com/chrisgardner402/bank-account-demo/accounts"
)

// SearchAccount searches an account
func SearchMassAccount(owners []string) (*[]accounts.Account, error) {
	accountSlice := []accounts.Account{}
	for _, owner := range owners {
		account, err := SearchAccount(owner)
		if err != nil {
			return nil, err
		}
		accountSlice = append(accountSlice, account)
	}
	return &accountSlice, nil
}

// UpdateAccount updates accounts for deposit and withdraw transactions
func DepositMassAccount(accountSlice *[]accounts.Account, amount int) error {
	// execute mass deposit
	accountSlice, err := accounts.MassDeposit(*accountSlice, amount)
	if err != nil {
		return nil
	}
	errC := make(chan error)
	// goroutine
	for _, account := range *accountSlice {
		go massDepositAccount(account, errC)
		if err != nil {
			return err
		}
	}
	// channel
	for i := 0; i < len(*accountSlice); i++ {
		if err := <-errC; err != nil {
			return err
		}
	}
	return nil
}

func massDepositAccount(account accounts.Account, errC chan<- error) {
	// update the account
	affect, err := updateAccount(&account)
	if err != nil {
		errC <- err
	}
	log.Println("mass deposit account affect:", affect)
	errC <- nil
}
