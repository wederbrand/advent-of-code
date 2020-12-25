package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	readFile, err := ioutil.ReadFile("25/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(strings.TrimSpace(string(readFile)), "\n")
	cardPubKey, _ := strconv.Atoi(input[0])
	doorPubKey, _ := strconv.Atoi(input[1])

	cardLoopSize := findLoop(cardPubKey)
	doorLoopSize := findLoop(doorPubKey)

	fmt.Println(cardLoopSize, doorLoopSize)

	enc := cardPubKey
	for i := 1; i < doorLoopSize; i++ {
		enc = transform(enc, cardPubKey)
	}

	fmt.Println(enc)
}

func findLoop(key int) int {
	loop := 1
	initialFind := 7
	find := 7
	for {
		loop++
		find = transform(find, initialFind)
		if find == key {
			return loop
		}
	}
}

func transform(find int, initialFind int) int {
	find *= initialFind
	find %= 20201227
	return find
}
