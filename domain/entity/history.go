package entity

type History struct {
	accountid   string
	historyid   string
	deposit     int
	withdraw    int
	historytime string
}

func CreateHistory(accountid string, historyid string, deposit int, withdraw int) *History {
	return &History{accountid: accountid, historyid: historyid, deposit: deposit, withdraw: withdraw}
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

func (h *History) Historytime() string {
	return h.historytime
}
