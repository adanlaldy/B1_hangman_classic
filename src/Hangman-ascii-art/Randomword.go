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

func Randomword() string {
	arg := os.Args[1]
	rand.Seed(time.Now().UnixNano())
	content, err := ioutil.ReadFile(arg)
	if err != nil {
		log.Fatal(err)
	}
	strconvert := string(content)
	split := strings.Split(strconvert, "\r\n")
	rand := rand.Intn(len(split))
	result := split[rand]
	return result
}
func Print(txt string) {
	content, err := os.ReadFile("res/parameters.txt")
	if err != nil {
		os.Exit(1)
	}
	split := strings.Split(strings.ReplaceAll(string(content), "\r", ""), "\n")
	for d := 0; d < 9; d++ {
		for index := 0; index < len(txt); index++ {
			print(split[(int(txt[index]-32)*9)+d] + "")
		}
		println()
	}
}
func Start(data *HangManData) {
	for i := 0; i < len(data.Randomword); i++ {
		data.Slice[i] = "_"
	}
	fmt.Println("Good Luck, you have 10 attempts.")
}
