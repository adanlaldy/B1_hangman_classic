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
func test(randomword string, slice []string) bool {
	for i := 0; i < len(randomword); i++ {
		if slice[i] != string(randomword[i]) {
			return false
		}
	}
	return true
}

//	func hangmanposition(data []string, start int, end int) string {
//		result := ""
//		for i := start - 1; i < end; i++ {
//			result = result + data[i] + "\n"
//		}
//		return result
//	}
func main() {
	rand.Seed(time.Now().UnixNano())
	randomword := strings.ToUpper(randomword())
	totaltry := 10
	try := ""
	letter := ""
	n := len(randomword)/2 - 1
	slice := make([]string, len(randomword))
	slicerandomword := make([]string, len(randomword))
	// imprimer les lettres dans la slicerandomword
	for i := 0; i < len(randomword); i++ {
		slicerandomword[i] = string(randomword[i])
	}
	// affichage des _
	for i := 0; i < len(randomword); i++ {
		slice[i] = "_"
	}
	fmt.Println("Good Luck, you have 10 attempts.")
	//affichage de n nombres de lettres dans la slice
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
	// boucle du programme
	for i := false; i != true; {
		b := false
		if totaltry == 0 {
			fmt.Println("\n")
			fmt.Println("Try again!")
			return
		}
		fmt.Print("Choose: ")
		fmt.Scan(&try)
		if try == randomword {
			for j := 0; j < len(slicerandomword); j++ {
				letter = slicerandomword[j]
				fmt.Print(letter)
				fmt.Print(" ")
			}
			fmt.Println("\n")
			fmt.Println("Congrats!")
			break
		}
		// si lettre entrée juste, imprimer dans la slice la lettre
		for i := 0; i < len(randomword); i++ {
			if try == slicerandomword[i] {
				slice[i] = try
				b = true
			}
		}
		if b == false {
			totaltry--
			fmt.Printf("Not present in the word, %d attempts remaining\n", totaltry)
			main2(totaltry)

		}
		for j := 0; j < len(slice); j++ {
			letter = slice[j]
			fmt.Print(letter)
			fmt.Print(" ")
		}
		fmt.Print("\n")
		// si la slice contient les mêmes lettres que la slicerandomword, i est vraie ET ça break, sinon i est faux
		if test(randomword, slice) == true {
			fmt.Print("\n")
			fmt.Println("Well played!")
			return
		}
	}
}

func getStringFromArray(data []string, start int, end int) string {
	result := ""
	for i := start; i < end && i < len(data); i++ {
		result = result + data[i] + "\n"
	}
	return result
}

func main2(life int) {
	data, err := os.ReadFile("hangman.txt") // lire le fichier text.txt
	if err != nil {
		fmt.Println(err)
	}

	splitted := strings.Split(string(data), "\n")
	tries := life
	fmt.Println(getStringFromArray(splitted, tries*8, (tries+1)*8))
}
