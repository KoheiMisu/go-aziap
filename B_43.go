package main

import (
	"strings"
	"bufio"
	"os"
	"strconv"
	"fmt"
)

const (
	wall = "#"
	enemy = "X"
)

type Point map[int]interface{}

type Screen struct {
	height int
	width int
	points map[int]Point
}

var sc = bufio.NewScanner(os.Stdin)

func main() {
	s := new(Screen)
	s.points = make(map[int]Point)

	if sc.Scan() {
		inputs := strings.Split(sc.Text(), " ")
		s.height, _ = strconv.Atoi(inputs[0])
		s.width, _ = strconv.Atoi(inputs[1])
	}

	for i := 0; i < s.height; i++ {
		if sc.Scan() {
			inputs := strings.Split(sc.Text(), "")
			p := make(Point)
			for index, val := range inputs {
				p[index] = val
			}
			s.points[i] = p
		}
	}
}
