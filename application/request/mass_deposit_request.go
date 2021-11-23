package request

type MassDepositRequest struct {
	Accountidlist []string `json:"accountidlist"`
	Amount        int      `json:"amount"`
}
