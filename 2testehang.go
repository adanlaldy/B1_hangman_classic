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

func main() {
	data, err := os.ReadFile("hangman.txt") // lire le fichier text.txt
	if err != nil {
		fmt.Println(err)
	}

	splitted := strings.Split(string(data), "\n")
	tries := 9
	fmt.Println(getStringFromArray(splitted, tries*8, (tries+1)*8))
}
