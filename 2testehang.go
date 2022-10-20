package main

import (
	"fmt"
	"os"
	"strings"
)

func getStringFromArray(data []string, start int, end int) string {
	result := ""
	for i := start - 1; i < end; i++ {
		result = result + data[i] + "\n"
	}
	return result
}

func main() {
	data, err := os.ReadFile("hangmanposition.txt") // lire le fichier text.txt
	if err != nil {
		fmt.Println(err)
	}
	splitted := strings.Split(string(data), "\n")
	tries := 10
	fmt.Println(getStringFromArray(splitted, tries*7, (tries+1)*7))
}
