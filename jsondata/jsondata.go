package jsondata

type ErrorResponse struct {
	Message string `json:"message"`
}

type OpenRequest struct {
	Owner string `json:"owner"`
}

type OpenResponse struct {
	Owner   string `json:"owner"`
	Balance int    `json:"balance"`
}

type AccountRequest struct {
	Owner string `json:"owner"`
}

type AccountResponse struct {
	Owner   string `json:"owner"`
	Balance int    `json:"balance"`
}

type DepositRequest struct {
	Owner  string `json:"owner"`
	Amount int    `json:"amount"`
}

type DepositResponse struct {
}

type WithdrawRequest struct {
	Owner  string `json:"owner"`
	Amount int    `json:"amount"`
}

type WithdrawResponse struct {
}
