package repository

import "github.com/chrisgardner402/bank-account-demo/domain/entity"

type HistoryRepository interface {
	RecordHistory(history *entity.History) error
	RecordMassHistory(historySlice *[]entity.History) error
}
