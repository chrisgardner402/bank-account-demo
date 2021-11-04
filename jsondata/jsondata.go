package jsondata

type ErrorResponse struct {
	Message string `json:"message"`
}

type DepositRequest struct {
	Accountid string `json:"accountid"`
	Amount    int    `json:"amount"`
}

type DepositResponse struct {
}

type WithdrawRequest struct {
	Accountid string `json:"accountid"`
	Amount    int    `json:"amount"`
}

type WithdrawResponse struct {
}

type MassDepositRequest struct {
	Accountidlist []string `json:"accountidlist"`
	Amount        int      `json:"amount"`
}

type MassDepositReponse struct {
}
