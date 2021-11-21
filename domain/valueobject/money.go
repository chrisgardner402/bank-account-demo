package valueobject

type Money struct {
	value int
}

func CreateMoney(value int) *Money {
	return &Money{value: value}
}

func (m *Money) Value() int {
	return m.value
}
