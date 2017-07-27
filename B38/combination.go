package B38

import "fmt"

type Combination struct {
	LegSum,
	HeadSum,
	CraneLeg,
	TurtleLeg int

	Validator Validator
}

func NewCombination(legSum, headSum, craneLeg, turtleLeg int) *Combination {
	c := new(Combination)

	c.LegSum = legSum
	c.HeadSum = headSum
	c.CraneLeg = craneLeg
	c.TurtleLeg = turtleLeg

	return c
}

func (c *Combination) Output() {
	if !c.Validator.CheckImportValue(*c) {
		fmt.Println("miss")
		return
	}

	// when crane and turtle legs are same
	if c.CraneLeg == c.TurtleLeg {
		// invalid Combination
		if c.LegSum != (c.TurtleLeg * c.HeadSum) {
			fmt.Println("miss")
			return
		}

		if c.HeadSum > 2 {
			fmt.Println("miss")
			return
		}
	}

	turtle := 0

	for ;; {
		if turtle > 100 {
			fmt.Println("miss")
			return
		}

		crane := c.HeadSum - turtle
		if c.LegSum == ((crane * c.CraneLeg) + (turtle * c.TurtleLeg)) {
			// crane and turtle must be over 1
			if turtle == 0 {
				turtle += 1
				continue
			}
			fmt.Printf("%d %d\n", crane, turtle)
			break
		}
		turtle += 1
	}
}
