package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

// func reveal(string) {
// randomword := "HELLO"
// //n := len(randomword)/2 - 1
//
//	for i := 0; i < len(randomword); i++ {
//	   fmt.Print("_")
//	   fmt.Print(" ")
//	}
//
// }
// func RemoveIndex(s []string, index int) []string {
// return append(s[:index], s[index+1:]...)
// }
func main() {
	rand.Seed(time.Now().UnixNano())
	randomword := "HELLLO"
	totaltry := 10
	try := ""
	slice := make([]string, len(randomword))
	for i := 0; i < len(randomword); i++ {
		slice[i] = "_"
	}
	letter := ""
	n := len(randomword)/2 - 1
	fmt.Println("Good Luck, you have 10 attempts.")
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

	for i := 0; i < 10; i++ {
		fmt.Print("Choose an uppercase letter : ")
		fmt.Scan(&try)
		if try == randomword {
			fmt.Println("Congrats !")
			break
		} else if try == string(randomword[i]) {
			slice[i] = try
			for j := 0; j < len(slice); j++ {
				letter = slice[j]
				fmt.Print(letter)
				fmt.Print(" ")
			}
			fmt.Print("\n")
		} else if try != randomword {
			totaltry--
			fmt.Printf("Not present in the word, %d attempts remaining\n", totaltry)
			if totaltry == 0 {
				{
					{
						data, err := ioutil.ReadFile("hangmanposition.txt") // lire le fichier text.txt
						if err != nil {
							fmt.Println(err)
						}

						fmt.Println(string(data)) // conversion de byte en string
					}
				}

			}

		}
	}
}
