package B27

import (
	"strings"
	"fmt"
)

type Players map[int]*Player

type Game struct {
	table Table
	players Players
	history []string
}

func NewGame(table Table, players Players, history []string) *Game {
	g := new(Game)

	g.table = table
	g.players = players
	g.history = history

	return g
}

func (g *Game) Boot() {
	playerNumber := 0

	for _, historyLine := range g.history {
		values := strings.Split(historyLine, " ")
		if g.table.isCardSame(values) {
			player := g.players[playerNumber]
			player.addCardCount(2)
		} else {
			playerNumber = g.refreshPlayerNumber(playerNumber)
		}
	}

	for _, player := range g.players {
		fmt.Println(player.cardCount)
	}
}

func (g *Game) refreshPlayerNumber(playerNumber int) int {
	playerNumber += 1

	// array does not have this behavior
	if _, ok := g.players[playerNumber]; ok {
		return playerNumber
	}

	return 0
}