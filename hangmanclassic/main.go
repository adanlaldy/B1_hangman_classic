package main

import (
	"fmt"
	"hangmanclassic/app"
	"strings"
)

func main() {
	data := &app.HangManData{}
	data.Randomword = strings.ToUpper(app.Randomword())
	data.TotalTries = 10
	data.NFormula = len(data.Randomword)/2 - 1
	data.Slice = make([]string, len(data.Randomword))
	data.SliceRandomword = make([]string, len(data.Randomword))
	app.PrintLettersInTheFullSlice(data.SliceRandomword, data.Randomword)
	app.Start(data.Randomword, data.Slice)
	app.PrintNLetters(data.Randomword, data.Slice, data.NFormula, data.Letter)
	for true {
		data.Boolean = false
		if app.IfZeroTry(data.TotalTries, data.Randomword) == true {
			return
		}
		fmt.Print("\nChoose: ")
		fmt.Scan(&data.Try)
		app.IfInputIsTheFullWord(data.Try, data.Letter, data.SliceRandomword, data.Randomword)
		app.IfInputIsTrue(data)
		if app.IfSliceIsFull(data.Randomword, data.Slice) == true {
			fmt.Println("\n\nWell played!")
			return
		}
	}
}
