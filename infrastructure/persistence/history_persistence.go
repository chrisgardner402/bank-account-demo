package persistence

import (
	"log"

	"github.com/chrisgardner402/bank-account-demo/domain/entity"
	"github.com/chrisgardner402/bank-account-demo/domain/repository"
)

func NewHistoryPersistence() repository.HistoryRepository {
	return &historyPersistence{}
}

type historyPersistence struct{}

func (hp historyPersistence) RecordHistory(history *entity.History) error {
	// create a statement
	stmt, err := db.Prepare("insert into history(accountid, historyid, deposit, withdraw) values(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	// execute a statement
	res, err := stmt.Exec(history.Accountid(), history.Historyid(), history.Deposit(), history.Withdraw())
	if err != nil {
		return err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Println("record history affect:", affect)
	return nil
}

func (hp historyPersistence) RecordMassHistory(historySlice *[]entity.History) error {
	for _, history := range *historySlice {
		err := hp.RecordHistory(&history)
		if err != nil {
			return err
		}
	}
	return nil
}
