package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	readFile, err := ioutil.ReadFile("5/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	all := strings.Split(strings.TrimSpace(string(readFile)), "\n")

	seats := make(map[int]bool, 0)

	for _, s := range all {
		rowMin := 0
		rowMax := 127
		colMin := 0
		colMax := 7
		for _, r := range s {
			switch r {
			case 'F':
				rowMax -= (rowMax - rowMin + 1) / 2
			case 'B':
				rowMin += (rowMax - rowMin + 1) / 2
			case 'L':
				colMax -= (colMax - colMin + 1) / 2
			case 'R':
				colMin += (colMax - colMin + 1) / 2
			}
		}
		id := rowMin*8 + colMin
		seats[id] = true // someone sits here
	}

	i := 0
	for {
		i++
		if !seats[i] && seats[i-1] && seats[i+1] {
			fmt.Println(i)
			return
		}
	}
}
