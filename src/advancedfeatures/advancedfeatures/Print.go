package main

import (
	"fmt"
	"math/rand"
)

func PrintLettersInTheFullSlice(data *HangManData) {
	for i := 0; i < len(data.Randomword); i++ {
		data.SliceRandomword[i] = string(data.Randomword[i])
	}
}
func PrintNLetters(data *HangManData) {
	for i := 0; i < data.NFormula; i++ {
		c := data.Randomword[rand.Intn(len(data.Randomword))]
		for j := 0; j < len(data.Randomword); j++ {
			if data.Randomword[j] == c {
				data.Slice[j] = string(c)
			}
		}
	}
	for j := 0; j < len(data.Slice); j++ {
		data.Letter = data.Slice[j]
		fmt.Print(data.Letter)
		fmt.Print(" ")
	}
	fmt.Print("\n")
}
