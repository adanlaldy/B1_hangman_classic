package main

import (
	"os"
	"strings"
)

func PrintAscii(txt string) {
	content, err := os.ReadFile("standard.txt")
	if err != nil {
		os.Exit(1)
	}
	split := strings.Split(string(content), "\n")
	for d := 0; d < 9; d++ {
		for index := 0; index < len(txt); index++ {
			print(split[(int(txt[index]-32)*9)+d] + "")
		}
		println()
	}
}
func main() {
	PrintAscii("hello")
}
