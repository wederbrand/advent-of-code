package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"time"
)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2016/06/input.txt", "\n")
	part1, part2 := doIt(inFile)
	fmt.Println("Part 1: ", part1, "in", time.Since(start))
	fmt.Println("Part 2: ", part2, "in", time.Since(start))
}

func doIt(inFile []string) (string, string) {
	letters := make([]map[string]int, len(inFile[0]))
	for i := range inFile[0] {
		letters[i] = make(map[string]int)
	}

	for _, line := range inFile {
		for i, r := range line {
			letters[i][string(r)]++
		}
	}

	part1 := ""
	part2 := ""

	for _, m := range letters {
		part1 += mostCommon(m)
		part2 += leastCommon(m)
	}

	return part1, part2
}

func mostCommon(m map[string]int) string {
	maxFound := 0
	maxLetter := ""
	for letter, count := range m {
		if count > maxFound {
			maxFound = count
			maxLetter = letter
		}
	}
	return maxLetter
}

func leastCommon(m map[string]int) string {
	minFound := 100000
	minLetter := ""
	for letter, count := range m {
		if count < minFound {
			minFound = count
			minLetter = letter
		}
	}
	return minLetter
}
