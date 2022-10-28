package main

import (
	"fmt"
	"os"
	"strings"
)

func getStringFromArray(data []string, start int, end int) string {
	result := ""
	for i := start; i < end && i < len(data); i++ {
		result = result + data[i] + "\n"
	}
	return result
}

func main8(life int) {
	data, err := os.ReadFile("thinkertoy.txt") // read the text.txt file
	if err != nil {
		fmt.Println(err)
	}

	splitted := strings.Split(string(data), "\n")
	tries := 9 - life
	fmt.Println(getStringFromArray(splitted, tries*8, (tries+1)*8))
}
