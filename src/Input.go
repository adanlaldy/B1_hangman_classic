package main

import "fmt"

func ifInputIsTrue1(data *HangManData) {
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
			fmt.Print(data.letter)
			fmt.Print(" ")
		}
	}
}

func ifInputIsTheFullWord1(try string, letter string, sliceRandomword []string, randomword string) {
	if try == randomword {
		for j := 0; j < len(sliceRandomword); j++ {
			letter = sliceRandomword[j]
			fmt.Print(letter)
			fmt.Print(" ")
		}
		fmt.Println("\n\nCongrats!")
		return
	}
}
