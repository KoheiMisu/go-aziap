package main

import (
	"./B27"
)

func main() {
	var cards []string
	cards = append(cards, "5 8 8 6 3")
	cards = append(cards, "3 6 3 3 5")

	var history []string
	history = append(history, "1 4 2 2")
	history = append(history, "1 3 2 1")
	history = append(history, "2 4 2 3")
	history = append(history, "1 3 1 5")
	history = append(history, "2 5 1 1")
	history = append(history, "2 1 1 2")
	history = append(history, "1 5 2 1")
	history = append(history, "1 2 1 3")

	players := make(B27.Players)
	players[0] = B27.NewPlayer()
	players[1] = B27.NewPlayer()
	players[2] = B27.NewPlayer()

	t := *B27.NewTable(2, 5, cards)

	g := B27.NewGame(t, players, history)
	g.Boot()
}
