package main

import (
	"./B29"
)

type Lands []B29.Land

func main()  {
	targetLand := B29.NewLand(0, 0, 0)

	var lands Lands

	lands = append(lands, *B29.NewLand(0, 2, 4))
	lands = append(lands, *B29.NewLand(1, 1, 5))
	lands = append(lands, *B29.NewLand(1, 2, 6))
	lands = append(lands, *B29.NewLand(1, 2, 6))

	c := B29.NewCalculator(3, targetLand, lands)
	c.CalculatePrice()
}
