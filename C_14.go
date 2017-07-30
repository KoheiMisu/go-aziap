package main

import (
	"fmt"
	"strings"
	"strconv"
	"bufio"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main()  {
	var input string
	if sc.Scan() {
		input = sc.Text()
	}

	inputSplitted := strings.Split(input, " ")

	boxCount, _ := strconv.Atoi(inputSplitted[0])
	r, _ := strconv.Atoi(inputSplitted[1])
	diameter := r * 2

	for i := 1; i <= boxCount; i++ {
		var boxLine string
		if sc.Scan() {
			boxLine = sc.Text()
		}

		boxLineSplitted := strings.Split(boxLine, " ")

		boxLines := make([]int, len(boxLineSplitted))
		for i, v := range boxLineSplitted {
			val, _ := strconv.Atoi(v)
			boxLines[i] = val
		}

		if diameter <= boxLines[0] &&
			diameter <= boxLines[1] &&
			diameter <= boxLines[2] {
			fmt.Println(i)
		}
	}
}
