package repository

import (
	"log"

	"github.com/chrisgardner402/bank-account-demo/domain/aggregate"
	"github.com/chrisgardner402/bank-account-demo/domain/entity"
	"github.com/chrisgardner402/bank-account-demo/domain/factory"
	"github.com/chrisgardner402/bank-account-demo/infra/persistence/mapper"
)

func SearchAccount(account *entity.Account) (*entity.Account, error) {
	// execute a query
	row, err := mapper.SelectAccount(db, account.Accountid())
	// check error
	if err != nil {
		return factory.CreateEmptyAccount(), err
	}
	// check row
	defer row.Close()
	if !row.Next() {
		return factory.CreateEmptyAccount(), errNotFound
	}
	// create account object
	var userid string
	var accountid string
	var balance int
	err = row.Scan(&userid, &accountid, &balance)
	if err != nil {
		return factory.CreateEmptyAccount(), err
	}
	accountPersist := factory.CreateAccount(userid, accountid, balance)
	return accountPersist, nil
}

func SearchMassAccount(accountSlice *[]entity.Account) (*[]entity.Account, error) {
	var accountSlicePersist []entity.Account
	for _, account := range *accountSlice {
		accountPersist, err := SearchAccount(&account)
		if err != nil {
			return nil, err
		}
		accountSlicePersist = append(accountSlicePersist, *accountPersist)
	}
	return &accountSlicePersist, nil
}

func DepositAccount(deposit *aggregate.Deposit) error {
	// create and execute a statement
	res, err := mapper.UpdateAccountForDeposit(db, deposit.Amount(), deposit.Account())
	if err != nil {
		return err
	}
	// get the number of rows affected
	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Println("deposit account affect:", affect)
	return nil
}

func DepositMassAccount(depositSlice *[]aggregate.Deposit) error {
	for _, deposit := range *depositSlice {
		err := DepositAccount(&deposit)
		if err != nil {
			return err
		}
	}
	return nil
}

func WithdrawAccount(withdarw *aggregate.Withdraw) error {
	// create and execute a statement
	res, err := mapper.UpdateAccountForWithdraw(db, withdarw.Amount(), withdarw.Account())
	if err != nil {
		return err
	}
	// get the number of rows affected
	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Println("withdraw account affect:", affect)
	return nil
}
