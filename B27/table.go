package B27

import (
	"fmt"
)

type Cards map[int]Card

type FieldOfCards map[int]Cards

//type Table struct {
//	 field []FieldOfCards
//}

type Table struct {
	field Cards
}

func NewTable(length int, side int, cards []string) *Table {
	t := new(Table)
	t.field = make(Cards)

	for i := 1; i <= length; i++ {
		//sideValues := strings.Split(cards[i-1], " ")

		for j := 1; j <= side; j++ {
			//cards := make(Cards)


			//field := Cards{j: *NewCard(1)}
			//intVal, _ := strconv.Atoi(sideValues[j-1])
			t.field[i] = *NewCard(1)
			//fmt.Printf("%v\n", t.field)
			//t.field[i] = make(FieldOfCards)
			//t.field[i][j] = *NewCard(intVal)
			//fmt.Printf("%v\n", field)
			//t.field[i] = field
		}
	}

	fmt.Printf("%v\n", t.field)

	return t
}
