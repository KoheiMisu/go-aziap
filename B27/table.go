package B27

import (
	"strings"
	"strconv"
)

type Cards map[int]Card

type FieldOfCards map[int]Cards

type Table struct {
	field FieldOfCards
}

func NewTable(length int, side int, cards []string) *Table {
	t := new(Table)
	t.field = make(FieldOfCards)

	for i := 1; i <= length; i++ {
		t.field[i] = make(Cards)
		sideValues := strings.Split(cards[i-1], " ")

		for j := 1; j <= side; j++ {
			intVal, _ := strconv.Atoi(sideValues[j-1]) // string to int
			t.field[i][j] = *NewCard(intVal)
		}
	}

	return t
}

func (t *Table) isCardSame(content []string) bool {
	if t.getCardNumber(content[0], content[1]) == t.getCardNumber(content[2], content[3]) {
		return true
	}

	return false
}

func (t *Table) getCardNumber(length, side string) int  {
	intLength, _ := strconv.Atoi(length) // string to int
	intSide, _ := strconv.Atoi(side) // string to int

	card := t.field[intLength][intSide]

	return card.number
}
