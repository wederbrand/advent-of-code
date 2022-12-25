package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	readFile, err := ioutil.ReadFile("15/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(strings.TrimSpace(string(readFile)), ",")

	turn := 0
	numbers := make(map[int]int)
	lastNumber := 0
	for _, s := range input {
		turn++
		atoi, _ := strconv.Atoi(s)
		numbers[atoi] = turn
		lastNumber = atoi
	}

	mem := 0
	for turn < 30_000_000 {
		turn++

		// fake mem if it was in the input data
		lastTurn := numbers[lastNumber]
		if lastTurn <= len(input) {
			mem = 0
		}

		// set the next number to the memory
		lastNumber = mem

		// calculate memory to be the current row - the old index for this number
		lastTimeSeen, found := numbers[lastNumber]
		if !found {
			// fake mem if we've never seen this number
			mem = 0
		} else {
			mem = turn - lastTimeSeen
		}

		// set the new index for this number
		numbers[lastNumber] = turn
		if turn == 2020 {
			fmt.Println(lastNumber)
		}
		if turn == 30_000_000 {
			fmt.Println(lastNumber)
		}
	}
}
