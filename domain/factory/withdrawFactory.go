package factory

import (
	"github.com/chrisgardner402/bank-account-demo/domain/aggregate"
	"github.com/chrisgardner402/bank-account-demo/domain/entity"
)

func CreateWithdraw(account *entity.Account, amount int) *aggregate.Withdraw {
	amountVO := CreateMoney(amount)
	withdraw := aggregate.CreateWithdraw(account, amountVO)
	return withdraw
}
