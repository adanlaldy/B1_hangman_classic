package main

import (
	"os"
	"strings"
)

func Print(txt string) {
	content, err := os.ReadFile("res/standard.txt")
	if err != nil {
		os.Exit(1)
	}
	split := strings.Split(strings.ReplaceAll(string(content), "\r", ""), "\n")
	for d := 0; d < 9; d++ {
		for index := 0; index < len(txt); index++ {
			print(split[(int(txt[index]-32)*9)+d] + "")
		}
		println()
	}
}

func main() {
	Print("hello")
}
