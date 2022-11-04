package main

import (
	"fmt"
	"os"
	"strings"
)

func positionJose(life int) {
	data, err := os.ReadFile("positions/hangman.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	content := strings.ReplaceAll(string(data), "\r", "")
	splitted := strings.Split(content, "\n")
	tries := 9 - life
	fmt.Print(getStringFromArray2(splitted, tries*8, (tries+1)*8))
}
func getStringFromArray2(data []string, start int, end int) string {
	result := ""
	for i := start; i < end-1 && i < len(data); i++ {
		result += data[i] + "\n"
	}
	return result
}
