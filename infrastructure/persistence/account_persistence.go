package persistence

import (
	"log"

	"github.com/chrisgardner402/bank-account-demo/domain/aggregate"
	"github.com/chrisgardner402/bank-account-demo/domain/entity"
	"github.com/chrisgardner402/bank-account-demo/domain/factory"
	"github.com/chrisgardner402/bank-account-demo/domain/repository"
)

func NewAccountPersistence() repository.AccountRepository {
	return &accountPersistence{}
}

type accountPersistence struct{}

func (ap accountPersistence) SearchAccount(account *entity.Account) (*entity.Account, error) {
	// execute a query
	row, err := db.Query("select userid, accountid, balance from account where accountid = ?", account.Accountid())
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

func (ap accountPersistence) SearchMassAccount(accountSlice *[]entity.Account) (*[]entity.Account, error) {
	var accountSlicePersist []entity.Account
	for _, account := range *accountSlice {
		accountPersist, err := ap.SearchAccount(&account)
		if err != nil {
			return nil, err
		}
		accountSlicePersist = append(accountSlicePersist, *accountPersist)
	}
	return &accountSlicePersist, nil
}

func (ap accountPersistence) DepositAccount(deposit *aggregate.Deposit) error {
	// create a statement
	stmt, err := db.Prepare("update account set balance = balance + ? where accountid = ?")
	if err != nil {
		return err
	}
	// execute a statement
	res, err := stmt.Exec(deposit.Amount().Value(), deposit.Account().Accountid())
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

func (ap accountPersistence) DepositMassAccount(depositSlice *[]aggregate.Deposit) error {
	for _, deposit := range *depositSlice {
		err := ap.DepositAccount(&deposit)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ap accountPersistence) WithdrawAccount(withdarw *aggregate.Withdraw) error {
	// create a statement
	stmt, err := db.Prepare("update account set balance = balance - ? where accountid = ?")
	if err != nil {
		return err
	}
	// execute a statement
	res, err := stmt.Exec(withdarw.Amount().Value(), withdarw.Account().Accountid())
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
