package request

type DepositRequest struct {
	Accountid string `json:"accountid"`
	Amount    int    `json:"amount"`
}
