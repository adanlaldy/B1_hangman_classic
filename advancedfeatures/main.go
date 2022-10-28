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
func ifsliceisfull(randomword string, slice []string) bool {
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
func positionjose(life int) {
	data, err := os.ReadFile("hangmanpositions") // lire le fichier text.txt
	if err != nil {
		fmt.Println(err)
		return
	}
	content := strings.ReplaceAll(string(data), "\r", "")
	splitted := strings.Split(content, "\n")
	tries := 9 - life
	fmt.Print(getStringFromArray(splitted, tries*8, (tries+1)*8))
}
func main() {
	rand.Seed(time.Now().UnixNano())
	randomword := strings.ToUpper(randomword())
	totaltry := 10
	try := ""
	letter := ""
	n := len(randomword)/2 - 1
	slice := make([]string, len(randomword))
	slicerandomword := make([]string, len(randomword))
	slicetry := []string{}
	// print letters in slicerandomword
	for i := 0; i < len(randomword); i++ {
		slicerandomword[i] = string(randomword[i])
	}
	// print the "_"
	for i := 0; i < len(randomword); i++ {
		slice[i] = "_"
	}
	fmt.Println("Good Luck, you have 10 attempts.")
	// print n number of letters in the slice
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
	// buckle of the program
	for i := false; i != true; {
		a := false
		b := false
		if totaltry == 0 {
			fmt.Println("\nThe word was:", randomword)
			fmt.Println("\nTry again!")
			return
		}
		fmt.Print("\nChoose: ")
		fmt.Scan(&try)
		if try == randomword {
			for j := 0; j < len(slicerandomword); j++ {
				letter = slicerandomword[j]
				fmt.Print(letter)
				fmt.Print(" ")
			}
			fmt.Println("\n\nCongrats!")
			return
		}
		// if the input is true, print in the slice the letter
		for i := 0; i < len(randomword); i++ {
			if try == slicerandomword[i] {
				slice[i] = try
				b = true
			}
		}
		// error message when the input is already entered
		for i := 0; i < len(slicetry); i++ {
			if try == slicetry[i] {
				fmt.Println("You can't enter the same letter twice, try again.")
				a = true
				b = false
			}
		}
		if b == true {
			slicetry = append(slicetry, try)
		}
		if b == false && a == false {
			slicetry = append(slicetry, try)
			totaltry--
			totaltry--
			fmt.Printf("Not present in the word, %d attempts remaining\n", totaltry)
			positionjose(totaltry)
		} else {
			for j := 0; j < len(slice); j++ {
				letter = slice[j]
				fmt.Print(letter)
				fmt.Print(" ")
			}
		}
		// if the slice contains the same letters than slicerandomword, i is true AND it return, else i is false
		if ifsliceisfull(randomword, slice) == true {
			fmt.Println("\n\nWell played!")
			return
		}
	}
}
