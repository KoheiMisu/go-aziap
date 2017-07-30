package main

import (
	"strings"
	"strconv"
	"fmt"
	"bufio"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var input string
	if sc.Scan() {
		input = sc.Text()
	}

	inputSplitted := strings.Split(input, " ")

	stock, _ := strconv.Atoi(inputSplitted[0])
	firstSalesRate, _ := strconv.Atoi(inputSplitted[1])
	secondSalesRate, _ := strconv.Atoi(inputSplitted[2])

	unsold := (float64(stock) * ((100 - float64(firstSalesRate)) / 100)) * ((100 - float64(secondSalesRate)) / 100)

	fmt.Println(unsold)
}
