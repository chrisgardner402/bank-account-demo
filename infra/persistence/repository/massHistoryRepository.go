package repository

import "github.com/chrisgardner402/bank-account-demo/domain/model"

// SaveMassDepositHis saves histories for deposit transactions
func SaveMassDepositHis(accountlist *[]model.Account, amount int) error {
	for _, account := range *accountlist {
		err := SaveDepositHis(&account, amount)
		if err != nil {
			return err
		}
	}
	return nil
}
