package main

import (
	"./A17"

	"strings"
	"strconv"
)

func main() {
	screenHeight := 10
	screenWidth := 10

	var data []string
	data = append(data, "1 8 1")
	data = append(data, "4 1 5")
	data = append(data, "1 6 2")
	data = append(data, "2 2 0")

	//data = append(data, "2 2 0")
	//data = append(data, "1 6 2")
	//data = append(data, "4 1 5")
	//data = append(data, "1 8 1")

	var fallenObjects []A17.FallenObject

	for _, str := range data {
		objData := strings.Split(str, " ")

		height, _ := strconv.Atoi(objData[0])
		width, _ := strconv.Atoi(objData[1])
		offset, _ := strconv.Atoi(objData[2])

		var fallenObject A17.FallenObject = A17.NewRectangle(height, width, offset)

		fallenObjects = append(fallenObjects, fallenObject)
	}

	h := A17.NewHistory(screenWidth)
	s := A17.NewScreen(screenHeight, screenWidth, h)

	simulation := A17.NewSimulation(s, fallenObjects)
	simulation.Output()
}