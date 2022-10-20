package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

func main() {
	content, err := ioutil.ReadFile("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	rand.Seed(time.Now().UnixNano())
	strconvert := string(content)
	split := strings.Split(strconvert, "\n")
	rand := rand.Intn(len(split))
	result := split[rand]
	fmt.Print(result)
}
