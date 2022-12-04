package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	points := make(map[rune]int)
	p := 1
	for r := 'a'; r <= 'z'; r++ {
		points[r] = p
		points[unicode.ToUpper(r)] = p + 26
		p++
	}

	readFile, err := os.ReadFile("04/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	inFile := strings.Split(strings.TrimSpace(string(readFile)), "\n")

	contained := 0
	overlap := 0
	for _, s := range inFile {
		r1Start := 0
		r1End := 0
		r2Start := 0
		r2End := 0
		fmt.Sscanf(s, "%d-%d,%d-%d", &r1Start, &r1End, &r2Start, &r2End)

		if r1Start <= r2Start && r1End >= r2End || r2Start <= r1Start && r2End >= r1End {
			contained++
		}

		if r1End < r2Start || r2End < r1Start {
			// no overlap
		} else {
			overlap++
		}
	}

	fmt.Println(contained)
	fmt.Println(overlap)
}
