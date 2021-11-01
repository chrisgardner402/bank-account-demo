package main

import (
	"fmt"

	"github.com/chrisgardner402/bank-account-demo/accounts"
)

func main() {
	var accountSlice []accounts.Account

	account1 := accounts.NewAccount("owner1")
	account2 := accounts.NewAccount("owner2")
	account3 := accounts.NewAccount("owner3")
	account4 := accounts.NewAccount("owner4")
	account5 := accounts.NewAccount("owner5")
	account6 := accounts.NewAccount("owner6")
	account7 := accounts.NewAccount("owner7")
	account8 := accounts.NewAccount("owner8")
	account9 := accounts.NewAccount("owner9")

	accountSlice = append(accountSlice, *account1, *account2, *account3, *account4, *account5, *account6, *account7, *account8, *account9)

	for _, account := range accountSlice {
		fmt.Println(account.String())
	}

	accountSlice = accounts.DepositAll(accountSlice, 1000)

	for _, account := range accountSlice {
		fmt.Println(account.String())
	}
}
