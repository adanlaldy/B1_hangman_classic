package main

import (
	"fmt"
	"os"
	"strings"
)

func GetStringFromArray(data []string, start int, end int) string {
	result := ""
	for i := start; i < end-1 && i < len(data); i++ {
		result += data[i] + "\n"
	}
	return result
}
func PositionJose(life int) {
	data, err := os.ReadFile("hangmanpositions.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	content := strings.ReplaceAll(string(data), "\r", "")
	splitted := strings.Split(content, "\n")
	tries := 9 - life
	fmt.Print(GetStringFromArray(splitted, tries*8, (tries+1)*8))
}
