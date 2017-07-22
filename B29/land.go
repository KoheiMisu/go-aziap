package B29

type Land struct {
	X int
	Y int
	Price int
}

func NewLand(x int, y int, price int) *Land {
	l := new(Land)

	l.X = x
	l.Y = y
	l.Price = price

	return l
}