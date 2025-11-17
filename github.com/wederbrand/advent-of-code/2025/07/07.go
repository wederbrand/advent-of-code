package main

import (
	"fmt"
	"time"

	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
)

func main() {
	start := time.Now()
	input := GetFileContents("2025/07/input.txt", "\n")
	m := MakeChart(input, "")
	sInput, _ := m.FindLetter("S")

	part2 := doIt(sInput, m)
	part1 := len(seen)
	fmt.Println("Part 1:", part1, "in", time.Since(start))
	fmt.Println("Part 2:", part2, "in", time.Since(start))
}

var cache = make(map[Coord]int)
var seen = make(map[Coord]bool)

func doIt(sInput Coord, m Chart) int {
	result := 0
	found := false
	if result, found = cache[sInput]; found {
		return result
	}

	below, found := m[sInput.Move(DOWN)]
	if !found {
		// outside
		result = 1
	} else if below == "." {
		// can fall down
		result = doIt(sInput.Move(DOWN), m)
	} else if below == "^" {
		seen[sInput] = true
		// split
		result = doIt(sInput.Move(DOWNLEFT), m) + doIt(sInput.Move(DOWNRIGHT), m)
	} else {
		panic("ho ho ho")
	}
	cache[sInput] = result
	return result
}
