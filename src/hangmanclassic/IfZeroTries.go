package main

import "fmt"

func IfZeroTry(data *HangManData) bool {
	boolean := false
	if data.TotalTries == 0 {
		fmt.Println("\nThe word was:", data.Randomword)
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
			fmt.Print(data.Letter)
			fmt.Print(" ")
		}
	}
}
func IfInputIsTheFullWord(data *HangManData) {
	if data.Try == data.Randomword {
		for j := 0; j < len(data.SliceRandomword); j++ {
			data.Letter = data.SliceRandomword[j]
			fmt.Print(data.Letter)
			fmt.Print(" ")
		}
		fmt.Println("\n\nCongrats!")
		return
	}
}
