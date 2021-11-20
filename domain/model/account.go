package model

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	accountid string
	balance   int
}

var (
	errNoMoney     = errors.New("can't withdraw. you are poor")
	errNotPositive = errors.New("amount must be positive")
)

// CreateAccount creates an account
func CreateAccount(accountid string, balance int) *Account {
	account := Account{accountid: accountid, balance: balance}
	return &account
}

// Accountid of the account
func (a Account) Accountid() string {
	return a.accountid
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) error {
	CheckDeposit(amount)
	a.balance += amount
	return nil
}

func (a *Account) deposit(amount int, accountC chan<- Account) {
	a.balance += amount
	accountC <- *a
}

// MassDeposit x amount to all accounts
func MassDeposit(accountlist []Account, amount int) (*[]Account, error) {
	if err := checkAmount(amount); err != nil {
		return nil, err
	}
	var accountlistReturn []Account
	accountC := make(chan Account)
	for i := 0; i < len(accountlist); i++ {
		go accountlist[i].deposit(amount, accountC)
	}
	for i := 0; i < len(accountlist); i++ {
		account := <-accountC
		accountlistReturn = append(accountlistReturn, account)
	}
	return &accountlistReturn, nil
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	a.CheckWithdraw(amount)
	a.balance -= amount
	return nil
}

func (a Account) String() string {
	return fmt.Sprint(a.Accountid(), " has: ", a.Balance())
}

func CheckDeposit(amount int) error {
	if err := checkAmount(amount); err != nil {
		return err
	}
	return nil
}

func (a *Account) CheckWithdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	} else if err := checkAmount(amount); err != nil {
		return err
	}
	return nil
}

func checkAmount(amount int) error {
	if amount <= 0 {
		return errNotPositive
	}
	return nil
}
