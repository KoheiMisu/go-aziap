package B27

type Card struct {
	number int
}

func NewCard(number int) *Card {
	c := new(Card)

	c.number = number

	return c
}
