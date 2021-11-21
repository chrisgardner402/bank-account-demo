package entity

type Account struct {
	userid    string
	accountid string
	balance   int
}

func CreateAccount(userid string, accountid string, balance int) *Account {
	return &Account{userid: userid, accountid: accountid, balance: balance}
}

func (a *Account) Userid() string {
	return a.userid
}

func (a *Account) Accountid() string {
	return a.accountid
}

func (a *Account) Balance() int {
	return a.balance
}
