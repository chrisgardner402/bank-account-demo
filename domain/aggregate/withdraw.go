package aggregate

import (
	"github.com/chrisgardner402/bank-account-demo/domain/entity"
	"github.com/chrisgardner402/bank-account-demo/domain/valueobject"
)

type Withdraw struct {
	account entity.Account
	amount  valueobject.Money
}

func CreateWithdraw(account *entity.Account, amount *valueobject.Money) *Withdraw {
	return &Withdraw{account: *account, amount: *amount}
}

func (w *Withdraw) ValidateWithdraw() error {
	if w.amount.Value() <= 0 {
		return errNotPositive
	} else if w.account.Balance() < w.amount.Value() {
		return errNoMoney
	}
	return nil
}

func (w *Withdraw) Account() *entity.Account {
	return &w.account
}

func (w *Withdraw) Amount() *valueobject.Money {
	return &w.amount
}
