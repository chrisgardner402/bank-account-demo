package factory

import "github.com/chrisgardner402/bank-account-demo/domain/valueobject"

func CreateMoney(value int) *valueobject.Money {
	money := valueobject.CreateMoney(value)
	return money
}
