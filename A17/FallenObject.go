package A17

type coordinate map[int]int

type FallenObject interface {
	GetFiguredCoordinates() []coordinate

	SetFiguredCoordinates(coordinate)

	GetWidth() int

	GetHeight() int

	GetOffset() int
}
