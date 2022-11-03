package main

import (
	"advancedfeatures/app"
	"fmt"
	"strings"
)

func main() {
	data := &app.HangManData{}
	data.Randomword = strings.ToUpper(app.Randomword())
	data.TotalTries = 10
	data.NFormula = len(data.Randomword)/2 - 1
	data.Slice = make([]string, len(data.Randomword))
	data.SliceRandomword = make([]string, len(data.Randomword))
	app.PrintLettersInTheFullSlice(data)
	app.Start(data)
	app.PrintNLetters(data)
	for true {
		data.Boolean = false
		data.Boolean2 = false
		if app.IfZeroTry(data) == true {
			return
		}
		fmt.Print("\nChoose: ")
		fmt.Scan(&data.Try)
		app.IfInputIsTheFullWord(data)
		app.IfInputIsTrue(data)
		if app.IfSliceIsFull(data) == true {
			fmt.Println("\n\nWell played!")
			return
		}
	}
}
