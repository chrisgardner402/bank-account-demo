package validate

import "github.com/chrisgardner402/bank-account-demo/domain/model"

// ValidateDeposit validates deposit transaction
func ValidateDeposit(amount int) error {
	// before deposit
	if err := model.CheckDeposit(amount); err != nil {
		return err
	}
	return nil
}

// ValidateWithdraw validates withdraw transaction
func ValidateWithdraw(account model.Account, amount int) error {
	// before withdraw
	if err := account.CheckWithdraw(amount); err != nil {
		return err
	}
	return nil
}
