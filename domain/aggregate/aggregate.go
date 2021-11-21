package aggregate

import "errors"

var (
	errNoMoney     = errors.New("can't withdraw. you are poor")
	errNotPositive = errors.New("amount must be positive")
)
