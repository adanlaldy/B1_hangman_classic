package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
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

func hangmanposition(data []string, start int, end int) string {
	result := ""
	for i := start - 1; i < end; i++ {
		result = result + data[i] + "\n"
	}
	return result
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
	// choisir une lettre
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

	for z := false; z != true; {
		// break si la slice contient toutes les lettres
		k := false
		for j := 0; j < len(slicerandomword); j++ {
			if slice[j] == slicerandomword[j] {
				k = true
			} else {
				k = false
			}
		}
		if k == true {
			fmt.Print("\n")
			fmt.Println("Congrats !")
			break
		}
		fmt.Print("Choose: ")
		fmt.Scan(&try)
		if try == randomword {
			for j := 0; j < len(slice); j++ {
				letter = slicerandomword[j]
				fmt.Print(letter)
				fmt.Print(" ")
			}
			fmt.Println("\n")
			fmt.Println("Congrats !")
			break
		}
		// slice[i] prend la valeur de la lettre[i]
		f := false // A CHECKER
		for i := 0; i < len(randomword); i++ {
			if try == string(randomword[i]) {
				slice[i] = try
				f = true
			} else {
				f = false
			}
		}
		for j := 0; j < len(slice); j++ {
			letter = slice[j]
			fmt.Print(letter)
			fmt.Print(" ")
		}
		fmt.Print("\n")
		if f == false {
			totaltry--
			fmt.Printf("Not present in the word, %d attempts remaining\n", totaltry)
		}
		// conditions
		if totaltry == 1 {
			totaltry--
			fmt.Printf("Not present in the word, %d attempts remaining\n", totaltry)
			fmt.Println("The word was: ", randomword)
			break
		}
	}
}
