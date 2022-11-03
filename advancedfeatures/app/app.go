package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
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

func Randomword() string {
	rand.Seed(time.Now().UnixNano())
	content, err := ioutil.ReadFile("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	strconvert := string(content)
	split := strings.Split(strconvert, "\r\n")
	rand := rand.Intn(len(split))
	result := split[rand]
	return result
}
func IfSliceIsFull(data *HangManData) bool {
	for i := 0; i < len(data.Randomword); i++ {
		if data.Slice[i] != string(data.Randomword[i]) {
			return false
		}
	}
	return true
}
func GetStringFromArray(data []string, start int, end int) string {
	result := ""
	for i := start; i < end-1 && i < len(data); i++ {
		result += data[i] + "\n"
	}
	return result
}
func PositionJose(life int) {
	data, err := os.ReadFile("hangmanpositions.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	content := strings.ReplaceAll(string(data), "\r", "")
	splitted := strings.Split(content, "\n")
	tries := 9 - life
	fmt.Print(GetStringFromArray(splitted, tries*8, (tries+1)*8))
}
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
func IfInputIsTrue(data *HangManData) {
	for i := 0; i < len(data.Randomword); i++ {
		if data.Try == data.SliceRandomword[i] {
			data.Slice[i] = data.Try
			data.Boolean = true
		}
	}
	// error message when the input is already entered
	for i := 0; i < len(data.SliceTry); i++ {
		if data.Try == data.SliceTry[i] {
			fmt.Println("You can't enter the same letter twice, try again.")
			data.Boolean2 = true
			data.Boolean = false
		}
	}
	if data.Boolean == true {
		data.SliceTry = append(data.SliceTry, data.Try)
	}
	if data.Boolean == false && data.Boolean2 == false {
		data.SliceTry = append(data.SliceTry, data.Try)
		data.TotalTries--
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
func Start(data *HangManData) {
	for i := 0; i < len(data.Randomword); i++ {
		data.Slice[i] = "_"
	}
	fmt.Println("Good Luck, you have 10 attempts.")
}
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
