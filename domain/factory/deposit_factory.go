package factory

import (
	"github.com/chrisgardner402/bank-account-demo/domain/aggregate"
	"github.com/chrisgardner402/bank-account-demo/domain/entity"
)

func CreateDeposit(account *entity.Account, amount int) *aggregate.Deposit {
	amountVO := CreateMoney(amount)
	deposit := aggregate.CreateDeposit(account, amountVO)
	return deposit
}

func CreateDepositSlice(accountSlice *[]entity.Account, amount int) *[]aggregate.Deposit {
	var depositSlice []aggregate.Deposit
	for _, account := range *accountSlice {
		deposit := CreateDeposit(&account, amount)
		depositSlice = append(depositSlice, *deposit)
	}
	return &depositSlice
}
