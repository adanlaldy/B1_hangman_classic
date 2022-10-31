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

func PrintAscii(txt string) {
	content, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("bip bip boop")
		return
	}
	split := strings.Split(string(content), "\n")
	for _, val := range txt {
		fmt.Println(getStringFromArray(split, int((val-32)*9), int((val-31)*9)))
	}
}

func main() {
	PrintAscii("hello")
}
