package main

import (
	"fmt"
	"os"
	"strings"
)

func getStringFromArray8(data []string, start int, end int) string {
	result := ""
	for i := start; i < end && i < len(data); i++ {
		result = result + data[i] + "\n"
	}
	return result
}

func main8(life int) {
	data, err := os.ReadFile("position-Hangman-ascii-art.txt") // read the text.txt file
	if err != nil {
		fmt.Println(err)
	}

	splitted := strings.Split(string(data), "\n")
	tries := 9 - life
	fmt.Println(getStringFromArray(splitted, tries*10, (tries+1)*10))
}
