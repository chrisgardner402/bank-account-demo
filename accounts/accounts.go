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

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
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

func (a *Account) deposit(amount int, c chan<- Account) error {
	if amount <= 0 {
		return errNotPositive
	}
	a.balance += amount
	c <- *a
	return nil
}

// Deposit x amount to all accounts
func DepositAll(accounts []Account, amount int) []Account {
	var accountSlice []Account
	c := make(chan Account)
	for i := 0; i < len(accounts); i++ {
		go accounts[i].deposit(amount, c)
	}
	for i := 0; i < len(accounts); i++ {
		account := <-c
		accountSlice = append(accountSlice, account)
	}
	return accountSlice
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
