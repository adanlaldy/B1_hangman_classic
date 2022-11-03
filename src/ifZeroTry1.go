package main

import (
	"fmt"
	"os"
	"strings"
)

func ifZeroTry(totaltry int, randomword string) bool {
	boolean := false
	if totaltry == 0 {
		fmt.Println("\nThe word was:", randomword)
		fmt.Println("\nTry again!")
		boolean = true
		return boolean
	}
	return boolean
}
func ifSliceIsFull(randomword string, slice []string) bool {
	for i := 0; i < len(randomword); i++ {
		if slice[i] != string(randomword[i]) {
			return false
		}
	}
	return true
}
func ifInputIsTrue(data *HangManData) {
	for i := 0; i < len(data.randomword); i++ {
		if data.try == data.sliceRandomword[i] {
			data.slice[i] = data.try
			data.boolean = true
		}
	}
	if data.boolean == false {
		data.totalTries--
		fmt.Printf("Not present in the word, %d attempts remaining\n", data.totalTries)
		positionJose(data.totalTries)
	} else {
		for j := 0; j < len(data.slice); j++ {
			data.letter = data.slice[j]
		}
	}
}

func PrintAscii(txt string) {
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

func ifInputIsTheFullWord(try string, letter string, sliceRandomword []string, randomword string) {
	if try == randomword {
		res := ""
		for j := 0; j < len(sliceRandomword); j++ {
			letter = sliceRandomword[j]
			res = res + letter + " "
		}
		//PrintAscii(res)
		//fmt.Println("\n\nCongrats!")
		return
	}
}
