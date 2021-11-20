package request

type WithdrawRequest struct {
	Accountid string `json:"accountid"`
	Amount    int    `json:"amount"`
}
