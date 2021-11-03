package validate

import (
	"github.com/chrisgardner402/bank-account-demo/accounts"
)

// ValidateDeposit validates deposit transaction
func ValidateDeposit(amount int) error {
	// before deposit
	if err := accounts.CheckDeposit(amount); err != nil {
		return err
	}
	return nil
}

// ValidateWithdraw validates withdraw transaction
func ValidateWithdraw(account accounts.Account, amount int) error {
	// before withdraw
	if err := account.CheckWithdraw(amount); err != nil {
		return err
	}
	return nil
}
