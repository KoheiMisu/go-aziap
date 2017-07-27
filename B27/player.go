package B27

type Player struct {
	cardCount int
}

func NewPlayer() *Player {
	p := new(Player)

	return p
}

func (p *Player) addCardCount(count int)  {
	p.cardCount += count
}

func (p *Player) getCardCount() int {
	return p.cardCount
}