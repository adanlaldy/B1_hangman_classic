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
	Boolean         bool
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
func IfSliceIsFull(randomword string, slice []string) bool {
	for i := 0; i < len(randomword); i++ {
		if slice[i] != string(randomword[i]) {
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
func PrintLettersInTheFullSlice(sliceRandomword []string, randomword string) {
	for i := 0; i < len(randomword); i++ {
		sliceRandomword[i] = string(randomword[i])
	}
}
func PrintNLetters(randomword string, slice []string, n int, letter string) {
	for i := 0; i < n; i++ {
		c := randomword[rand.Intn(len(randomword))]
		for j := 0; j < len(randomword); j++ {
			if randomword[j] == c {
				slice[j] = string(c)
			}
		}
	}
	for j := 0; j < len(slice); j++ {
		letter = slice[j]
		fmt.Print(letter)
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
func IfInputIsTheFullWord(try string, letter string, sliceRandomword []string, randomword string) {
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
func Start(randomword string, slice []string) {
	for i := 0; i < len(randomword); i++ {
		slice[i] = "_"
	}
	fmt.Println("Good Luck, you have 10 attempts.")
}
func IfZeroTry(totaltry int, randomword string) bool {
	boolean := false
	if totaltry == 0 {
		fmt.Println("\nThe word was:", randomword)
		fmt.Println("\nTry again!")
		boolean = true
		return boolean
	}
	return boolean
}
