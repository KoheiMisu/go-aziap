package A17

import (
	"fmt"
	"strings"
)

type Simulation struct {
	screen *Screen
	fallenObjects []FallenObject
}

func NewSimulation(screen *Screen, fallenObjects []FallenObject) *Simulation {
	s := new(Simulation)

	s.screen = screen
	s.fallenObjects = fallenObjects

	return s
}

func (s *Simulation) Output()  {
	s.screen.Operate(s.fallenObjects)

	s.screen.TransformCoordinates(s.fallenObjects)

	result := s.screen.Coordinates

	start := len(result) -1

	for i := start; i >= 0; i-- {
		fmt.Println(strings.Join(result[i], ""))
	}
}
