package main

import (
	"fmt"
	"os"
	"strings"
)

func IfZeroTry(totaltry int, randomword string) bool {
	boolean := false
	if totaltry == 0 {
		fmt.Println("\nThe word was:", randomword)
		fmt.Println("\nTry again!")
		boolean = true
		return boolean
	}
	return boolean
}
func IfSliceIsFull(data *HangManData) bool {
	for i := 0; i < len(data.Randomword); i++ {
		if data.Slice[i] != string(data.Randomword[i]) {
			return false
		}
	}
	return true
}
func IfInputIsTrue(data *HangManData) {
	for i := 0; i < len(data.Randomword); i++ {
		if data.Try == data.SliceRandomword[i] {
			data.Slice[i] = data.Try
			data.Boolean = true
		}
	}
	if data.Boolean == false {
		data.TotalTries--
		fmt.Printf("Not present in the word, %d attempts remaining\n", data.TotalTries)
		PositionJose(data.TotalTries)
	} else {
		for j := 0; j < len(data.Slice); j++ {
			data.Letter = data.Slice[j]
		}
	}
}
func PrintAscii(txt string) {
	content, err := os.ReadFile("parameters/standard.txt")
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
func IfInputIsTheFullWord(data *HangManData) {
	if data.Try == data.Randomword {
		res := ""
		for j := 0; j < len(data.SliceRandomword); j++ {
			data.Letter = data.SliceRandomword[j]
			res = res + data.Letter + " "
		}
		return
	}
}
