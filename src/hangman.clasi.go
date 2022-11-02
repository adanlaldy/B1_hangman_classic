package main

import (
	"fmt"
	"strings"
)

type HangManData struct {
	try             string
	letter          string
	randomword      string
	totalTries      int
	nFormula        int
	slice           []string
	sliceRandomword []string
	boolean         bool
}

func main() {
	data := &HangManData{}
	data.randomword = strings.ToUpper(randomword())
	data.totalTries = 10
	data.nFormula = len(data.randomword)/2 - 1
	data.slice = make([]string, len(data.randomword))
	data.sliceRandomword = make([]string, len(data.randomword))
	printLettersInTheFullSlice(data.sliceRandomword, data.randomword)
	start(data.randomword, data.slice)
	printNLetters(data.randomword, data.slice, data.nFormula, data.letter)
	for true {
		data.boolean = false
		if ifZeroTry(data.totalTries, data.randomword) == true {
			return
		}
		fmt.Print("\nChoose: ")
		fmt.Scan(&data.try)
		ifInputIsTheFullWord(data.try, data.letter, data.sliceRandomword, data.randomword)
		ifInputIsTrue(data)
		if ifSliceIsFull(data.randomword, data.slice) == true {
			fmt.Println("\n\nWell played!")
			return
		}
	}
}
