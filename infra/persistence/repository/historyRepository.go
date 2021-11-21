package repository

import (
	"log"

	"github.com/chrisgardner402/bank-account-demo/domain/entity"
	"github.com/chrisgardner402/bank-account-demo/infra/persistence/mapper"
)

func RecordHistory(history *entity.History) error {
	// create and execute a statement
	res, err := mapper.InsertHistory(db, history)
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

func RecordMassHistory(historySlice *[]entity.History) error {
	for _, history := range *historySlice {
		err := RecordHistory(&history)
		if err != nil {
			return err
		}
	}
	return nil
}
