package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

var (
	errNoMoney     = errors.New("can't withdraw. you are poor")
	errNotPositive = errors.New("amount must be positive")
)

// CreateAccount creates an account
func CreateAccount(owner string, balance int) *Account {
	account := Account{owner: owner, balance: balance}
	return &account
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) error {
	if amount <= 0 {
		return errNotPositive
	}
	a.balance += amount
	return nil
}

func (a *Account) deposit(amount int, c chan<- Account) {
	a.balance += amount
	c <- *a
}

// MassDeposit x amount to all accounts
func MassDeposit(accounts []Account, amount int) ([]Account, error) {
	if amount <= 0 {
		return nil, errNotPositive
	}
	var accountSlice []Account
	c := make(chan Account)
	for i := 0; i < len(accounts); i++ {
		go accounts[i].deposit(amount, c)
	}
	for i := 0; i < len(accounts); i++ {
		account := <-c
		accountSlice = append(accountSlice, account)
	}
	return accountSlice, nil
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	} else if amount <= 0 {
		return errNotPositive
	}
	a.balance -= amount
	return nil
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account has: ", a.Balance())
}
