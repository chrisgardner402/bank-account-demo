package ledger

import (
	"errors"
	"fmt"

	"github.com/chrisgardner402/bank-account-demo/accounts"
)

type Ledger map[string]accounts.Account

var (
	errNotFound    = errors.New("account not found")
	errOwnerExists = errors.New("that owner already exists")
)

// Search for an account
func (l Ledger) Search(owner string) (accounts.Account, error) {
	account, exists := l[owner]
	if exists {
		return account, nil
	}
	return accounts.Account{}, errNotFound
}

// Add an account to the ledger
func (l Ledger) Add(account accounts.Account) error {
	_, err := l.Search(account.Owner())
	switch err {
	case nil:
		return errOwnerExists
	case errNotFound:
		l[account.Owner()] = account
	}
	return nil
}

func (l Ledger) Print() {
	fmt.Println(">>>>> ledger <<<<<")
	for _, account := range l {
		fmt.Println("*", account.String())
	}
}
