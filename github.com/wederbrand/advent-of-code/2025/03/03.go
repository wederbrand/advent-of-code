package main

import (
	"fmt"
	"time"

	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
)

func main() {
	start := time.Now()
	input := GetFileContents("2025/03/input.txt", "\n")

	part1 := 0
	part2 := 0

	for _, line := range input {
		part1 += doLine(line, 2)
		part2 += doLine(line, 12)
	}

	fmt.Println("Part 1:", part1, "in", time.Since(start))
	fmt.Println("Part 2:", part2, "in", time.Since(start))
}

func doLine(line string, batteries int) int {
	jolts := 0
	for i := batteries - 1; i >= 0; i-- {
		s := line[0 : len(line)-i]
		pos, one := findHighest(s)
		line = line[pos+1:]
		jolts *= 10
		jolts += Atoi(one)
	}
	return jolts
}

func findHighest(s string) (int, string) {
	high := '0'
	pos := 0
	for p, r := range s {
		if r > high {
			high = r
			pos = p
		}
	}
	return pos, string(high)
}
