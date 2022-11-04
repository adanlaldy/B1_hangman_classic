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
	SliceTry        []string
	Boolean         bool
	Boolean2        bool
}

func main() {
	data := &HangManData{}
	data.Randomword = strings.ToUpper(Randomword())
	data.TotalTries = 10
	data.NFormula = len(data.Randomword)/2 - 1
	data.Slice = make([]string, len(data.Randomword))
	data.SliceRandomword = make([]string, len(data.Randomword))
	PrintLettersInTheFullSlice(data)
	Start(data)
	PrintNLetters(data)
	for true {
		data.Boolean = false
		data.Boolean2 = false
		if IfZeroTry(data) == true {
			return
		}
		fmt.Print("\nChoose: ")
		fmt.Scan(&data.Try)
		data.Try = strings.ToUpper(data.Try)
		if IfInputIsTheFullWord(data) == true {
			return
		}
		IfInputIsTrue(data)
		if IfSliceIsFull(data) == true {
			fmt.Println("\n\nWell played!")
			return
		}
	}
}
