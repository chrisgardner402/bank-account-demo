package model

import (
	"github.com/google/uuid"
)

// History struct
type History struct {
	accountid string
	historyid string
	deposit   int
	withdraw  int
}

// CreateDepositHis creates an history
func CreateDepositHis(account *Account) (*History, error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return nil, nil
	}
	historyid := uuid.String()
	history := History{accountid: account.accountid, historyid: historyid}
	return &history, nil
}

// CreateWithdrawHis creates an history
func CreateWithdrawHis(account *Account) (*History, error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return nil, nil
	}
	historyid := uuid.String()
	history := History{accountid: account.accountid, historyid: historyid}
	return &history, nil
}

// CreateMassDepositHis creates mass histories
func CreateMassDepositHis(accountlist *[]Account) (*[]History, error) {
	var historylist []History
	for _, a := range *accountlist {
		history, err := CreateDepositHis(&a)
		if err != nil {
			return nil, err
		}
		historylist = append(historylist, *history)
	}
	return &historylist, nil
}

func (h *History) Accountid() string {
	return h.accountid
}

func (h *History) Historyid() string {
	return h.historyid
}

func (h *History) Deposit() int {
	return h.deposit
}

func (h *History) Withdraw() int {
	return h.withdraw
}
