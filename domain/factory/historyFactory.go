package factory

import (
	"github.com/chrisgardner402/bank-account-demo/domain/aggregate"
	"github.com/chrisgardner402/bank-account-demo/domain/entity"
	"github.com/google/uuid"
)

func CreateDepositHistory(deposit *aggregate.Deposit) (*entity.History, error) {
	historyid, err := newHistoryid()
	if err != nil {
		return nil, err
	}
	history := entity.CreateHistory(deposit.Account().Accountid(), historyid, deposit.Amount().Value(), 0)
	return history, nil
}

func CreateDepositHistorySlice(depositSlice *[]aggregate.Deposit) (*[]entity.History, error) {
	var historySlice []entity.History
	for _, deposit := range *depositSlice {
		history, err := CreateDepositHistory(&deposit)
		if err != nil {
			return nil, err
		}
		historySlice = append(historySlice, *history)
	}
	return &historySlice, nil
}

func CreateWithdrawHistory(withdraw *aggregate.Withdraw) (*entity.History, error) {
	historyid, err := newHistoryid()
	if err != nil {
		return nil, err
	}
	history := entity.CreateHistory(withdraw.Account().Accountid(), historyid, 0, withdraw.Amount().Value())
	return history, nil
}

func newHistoryid() (string, error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}
