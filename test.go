package main

import (
	"fmt"
	"os"
	"strings"
)

func hangmanposition(data []string, start int, end int) string {
	result := ""
	for i := start - 1; i < end; i++ {
		result += data[i] + "\n"
	}
	return result
}
func main() {
	data, err := os.ReadFile("hangmanpositions") // lire le fichier text.txt
	if err != nil {
		fmt.Println(err)
	}
	splitted := strings.Split(string(data), "\n")
	tries := 10
	fmt.Println(hangmanposition(splitted, tries*7, (tries+1)*7))
}
