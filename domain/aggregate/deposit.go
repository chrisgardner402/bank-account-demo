package aggregate

import (
	"github.com/chrisgardner402/bank-account-demo/domain/entity"
	"github.com/chrisgardner402/bank-account-demo/domain/valueobject"
)

type Deposit struct {
	account entity.Account
	amount  valueobject.Money
}

func CreateDeposit(account *entity.Account, amount *valueobject.Money) *Deposit {
	return &Deposit{account: *account, amount: *amount}
}

func (d *Deposit) ValidateDeposit() error {
	if d.amount.Value() <= 0 {
		return errNotPositive
	}
	return nil
}

func (d *Deposit) Account() *entity.Account {
	return &d.account
}

func (d *Deposit) Amount() *valueobject.Money {
	return &d.amount
}
