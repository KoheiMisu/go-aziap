package main

import (
	"fmt"
	"strings"
)

var translatePairs map[string]string = map[string]string{
	"A": "4",
	"E": "3",
	"G": "6",
	"I": "1",
	"O": "0",
	"S": "5",
	"Z": "2",
}

func main()  {
	var word string
	fmt.Scan(&word)

	wordSplitted := strings.Split(word, "")

	result := make([]string, len(wordSplitted))


	for _, string := range wordSplitted {
		if val, ok := translatePairs[string]; ok {
			result = append(result, val)
			continue
		}
		result = append(result, string)
	}

	fmt.Println(strings.Join(result, ""))
}
