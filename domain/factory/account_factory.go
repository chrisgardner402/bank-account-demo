package factory

import "github.com/chrisgardner402/bank-account-demo/domain/entity"

func CreateEmptyAccount() *entity.Account {
	return new(entity.Account)
}

func CreateAccount(userid string, accountid string, balance int) *entity.Account {
	account := entity.CreateAccount(userid, accountid, balance)
	return account
}

func CreateAccountFromAccountid(accountid string) *entity.Account {
	account := entity.CreateAccount("", accountid, 0)
	return account
}

func CreateAccountSliceFromAccountid(accountidList []string) *[]entity.Account {
	var accountSlice []entity.Account
	for _, accountid := range accountidList {
		account := entity.CreateAccount("", accountid, 0)
		accountSlice = append(accountSlice, *account)
	}
	return &accountSlice
}
