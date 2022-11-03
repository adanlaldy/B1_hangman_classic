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

func Start(data *HangManData) {
	for i := 0; i < len(data.Randomword); i++ {
		data.Slice[i] = "_"
	}
	fmt.Println("Good Luck, you have 10 attempts.")
}
