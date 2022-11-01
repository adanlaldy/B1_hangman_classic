package main

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
	try             string
	letter          string
	randomword      string
	totalTries      int
	nFormula        int
	slice           []string
	sliceRandomword []string
	boolean         bool
}

func randomword() string {
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
func ifSliceIsFull(randomword string, slice []string) bool {
	for i := 0; i < len(randomword); i++ {
		if slice[i] != string(randomword[i]) {
			return false
		}
	}
	return true
}
func getStringFromArray(data []string, start int, end int) string {
	result := ""
	for i := start; i < end-1 && i < len(data); i++ {
		result += data[i] + "\n"
	}
	return result
}
func positionJose(life int) {
	data, err := os.ReadFile("hangman.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	content := strings.ReplaceAll(string(data), "\r", "")
	splitted := strings.Split(content, "\n")
	tries := 9 - life
	fmt.Print(getStringFromArray(splitted, tries*8, (tries+1)*8))
}
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
	for j := 0; j < len(slice); j++ {
		letter = slice[j]
		fmt.Print(letter)
		fmt.Print(" ")
	}
	fmt.Print("\n")
}
func ifInputIsTrue(data *HangManData) {
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
func ifInputIsTheFullWord(try string, letter string, sliceRandomword []string, randomword string) {
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
func start(randomword string, slice []string) {
	for i := 0; i < len(randomword); i++ {
		slice[i] = "_"
	}
	fmt.Println("Good Luck, you have 10 attempts.")
}
func ifZeroTry(totaltry int, randomword string) bool {
	boolean := false
	if totaltry == 0 {
		fmt.Println("\nThe word was:", randomword)
		fmt.Println("\nTry again!")
		boolean = true
		return boolean
	}
	return boolean
}

func main() {
	data := &HangManData{}
	data.randomword = strings.ToUpper(randomword())
	data.totalTries = 10
	data.nFormula = len(data.randomword)/2 - 1
	data.slice = make([]string, len(data.randomword))
	data.sliceRandomword = make([]string, len(data.randomword))
	printLettersInTheFullSlice(data.sliceRandomword, data.randomword)
	start(data.randomword, data.slice)
	printNLetters(data.randomword, data.slice, data.nFormula, data.letter)
	for true {
		data.boolean = false
		if ifZeroTry(data.totalTries, data.randomword) == true {
			return
		}
		fmt.Print("\nChoose: ")
		fmt.Scan(&data.try)
		ifInputIsTheFullWord(data.try, data.letter, data.sliceRandomword, data.randomword)
		ifInputIsTrue(data)
		if ifSliceIsFull(data.randomword, data.slice) == true {
			fmt.Println("\n\nWell played!")
			return
		}
	}
}
