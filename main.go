package main

import (
	"fmt"

	"github.com/chrisgardner402/bank-account-demo/accounts"
)

func main() {
	account := accounts.NewAccount("cho")
	fmt.Println(account.String())
	account.Deposit(1000)
	fmt.Println(account.String())
	account.Withdraw(500)
	fmt.Println(account.String())
}
