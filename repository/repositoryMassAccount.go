package repository

import (
	"log"

	"github.com/chrisgardner402/bank-account-demo/model"
)

// SearchAccount searches an account
func SearchMassAccount(accountidlist []string) (*[]model.Account, error) {
	accountlist := []model.Account{}
	for _, accountid := range accountidlist {
		account, err := SearchAccount(accountid)
		if err != nil {
			return nil, err
		}
		accountlist = append(accountlist, account)
	}
	return &accountlist, nil
}

// UpdateAccount updates accounts for deposit and withdraw transactions
func DepositMassAccount(accountlist *[]model.Account, amount int) error {
	// execute mass deposit
	accountlist, err := model.MassDeposit(*accountlist, amount)
	if err != nil {
		return nil
	}
	errC := make(chan error)
	// goroutine
	for _, account := range *accountlist {
		go massDepositAccount(account, errC)
		if err != nil {
			return err
		}
	}
	// channel
	for i := 0; i < len(*accountlist); i++ {
		if err := <-errC; err != nil {
			return err
		}
	}
	return nil
}

func massDepositAccount(account model.Account, errC chan<- error) {
	// update the account
	affect, err := updateAccount(&account)
	if err != nil {
		errC <- err
	}
	log.Println("mass deposit account affect:", affect)
	errC <- nil
}
