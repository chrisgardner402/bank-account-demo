package repository

import (
	"github.com/chrisgardner402/bank-account-demo/domain/aggregate"
	"github.com/chrisgardner402/bank-account-demo/domain/entity"
)

type AccountRepository interface {
	SearchAccount(account *entity.Account) (*entity.Account, error)
	SearchMassAccount(accountSlice *[]entity.Account) (*[]entity.Account, error)
	DepositAccount(deposit *aggregate.Deposit) error
	DepositMassAccount(depositSlice *[]aggregate.Deposit) error
	WithdrawAccount(withdarw *aggregate.Withdraw) error
}
