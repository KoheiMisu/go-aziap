package A17

import (
	"../util"
)

type screenLine []string

type Screen struct {
	Coordinates []screenLine
	history *History
}

func NewScreen(screenHeight, screenWidth int, history *History) *Screen {
	s := new(Screen)

	for i := 0; i < screenHeight; i++ {
		// []interface{}から[]stringにする際もそれぞれの要素について型アサーションする必要がある
		fillValues := util.ArrayFill(0, screenWidth, ".")
		screenLine := make(screenLine, len(fillValues))
		for i, v := range fillValues {
			screenLine[i] = v.(string)
		}

		s.Coordinates = append(s.Coordinates, screenLine)
	}

	s.history = history

	return s
}

func (s *Screen) Operate(fallenObjects []FallenObject) {
	for _, fallenObject := range fallenObjects {
		origin := s.calcOrigin(fallenObject)
		fallenObject.SetFiguredCoordinates(origin)
	}
}

func (s *Screen) TransformCoordinates(fallenObjects []FallenObject) {
	for _, fallenObject := range fallenObjects {
		for _, coordinate := range fallenObject.GetFiguredCoordinates() {
			var x int
			for k := range coordinate {
				x = k
			}

			y := coordinate[x]

			s.Coordinates[y][x] = "#"
		}
	}
}

func (s *Screen) calcOrigin(fallenObject FallenObject) coordinate {
	y := s.history.GetHeight(fallenObject)

	s.history.SetHeight(fallenObject)

	coordinate := coordinate{fallenObject.GetOffset():y}

	return coordinate
}
