package main

import (
	"fmt"
	"math/rand"
)

func printLettersInTheFullSlice(sliceRandomword []string, randomword string) {
	for i := 0; i < len(randomword); i++ {
		sliceRandomword[i] = string(randomword[i])
	}
}
func printNLetters(randomword string, slice []string, n int, letter string) {
	for i := 0; i < n; i++ {
		c := randomword[rand.Intn(len(randomword))]
		for j := 0; j < len(randomword); j++ {
			if randomword[j] == c {
				slice[j] = string(c)
			}
		}
	}
	res := ""
	for j := 0; j < len(slice); j++ {
		letter = slice[j]
		res = res + (letter + " ")
	}
	PrintAscii(res)
	fmt.Print("\n")
}
