package repository

import "github.com/chrisgardner402/bank-account-demo/accounts"

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
func UpdateMassAccount(accountSlice *[]accounts.Account) error {
	for _, account := range *accountSlice {
		err := UpdateAccount(&account)
		if err != nil {
			return err
		}
	}
	return nil
}
