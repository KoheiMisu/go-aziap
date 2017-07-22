package B29

import (
	"fmt"
	"math"
	"sort"
)

type Calculator struct {
	lands []Land
	targetLand Land
	averageDenominator int
	priceSum int
}

type Distance []float64

func NewCalculator(averageDenominator int, targetLand *Land, lands []Land) *Calculator {
	c := new(Calculator)

	c.averageDenominator = averageDenominator
	c.targetLand = *targetLand
	c.lands = lands

	return c
}

func (c *Calculator) CalculatePrice() {
	// 比較するlandの数と分母が同じの場合、全ての計算を行う
	if c.averageDenominator == len(c.lands) {
		for _, land := range c.lands {
			c.priceSum += land.Price
		}

		fmt.Printf("%v\n", c.priceSum / c.averageDenominator)
		return
	}

	var distance Distance

	for _, land := range c.lands {
		xDiff := c.targetLand.X - land.X
		yDiff := c.targetLand.Y - land.Y
		distance = append(distance, math.Sqrt(math.Pow(float64(xDiff), 2)) + math.Pow(float64(yDiff), 2))
	}

	sort.Float64s(distance)

	slicedDistance := distance[:c.averageDenominator]

	for key, _ := range slicedDistance {
		l := c.lands[key]
		c.priceSum += l.Price
	}

	fmt.Printf("%v\n", c.priceSum / c.averageDenominator)
}
