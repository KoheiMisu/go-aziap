package main

import (
	"./B38"
)

func main()  {
	//32, 10, 2, 4
	//52, 65, 18, 76
	//8, 90, 46, 47
	//12, 4, 3, 3
	c := B38.NewCombination(2, 2, 1, 1)
	c.Output()
}
