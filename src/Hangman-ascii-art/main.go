package main

import (
	"fmt"
	"strings"
)

type HangManData struct {
	Try             string
	Letter          string
	Randomword      string
	TotalTries      int
	NFormula        int
	Slice           []string
	SliceRandomword []string
	Boolean         bool
}

func main() {
	data := &HangManData{}
	data.Randomword = Randomword()
	data.TotalTries = 10
	data.NFormula = len(data.Randomword)/2 - 1
	data.Slice = make([]string, len(data.Randomword))
	data.SliceRandomword = make([]string, len(data.Randomword))
	PrintLettersInTheFullSlice(data)
	Start(data)
	PrintNLetters(data)
	for true {
		data.Boolean = false
		if IfZeroTry(data.TotalTries, data.Randomword) == true {
			return
		}
		fmt.Print("\nChoose: ")
		fmt.Scan(&data.Try)
		data.Try = strings.ToLower(data.Try)
		if IfInputIsTheFullWord(data) == true {
			return
		}
		IfInputIsTrue(data)
		if IfSliceIsFull(data) == true {
			fmt.Println("\n\nWell played!")

			return
		} else {
			res := ""
			for _, word := range data.Slice {
				res = res + word
			}
			PrintAscii(res)
		}
	}
}
