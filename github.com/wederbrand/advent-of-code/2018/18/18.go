package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"time"
)

func main() {
	start := time.Now()
	inFile := GetFileContents("2018/18/input.txt", "\n")

	m := MakeChart(inFile, "")

	for i := 0; i < 10; i++ {
		_, m = doIt(m, i)
	}

	part1 := countTreesAndLumberYards(m)
	fmt.Println("Part 1:", part1, "in", time.Since(start))

	warp := 0
	for i := 10; i < 1000000000; i++ {
		warp, m = doIt(m, i)
		if warp != 0 {
			warpSpeed := i - warp
			i += warpSpeed * ((1000000000 - i) / warpSpeed)
		}
	}

	part2 := countTreesAndLumberYards(m)
	fmt.Println("Part 2:", part2, "in", time.Since(start))
}

func countTreesAndLumberYards(m Chart) int {
	trees := 0
	lumberyards := 0
	for _, s := range m {
		if s == "|" {
			trees++
		} else if s == "#" {
			lumberyards++
		}
	}
	return trees * lumberyards
}

var cache = map[string]int{}

func doIt(m Chart, i int) (int, Chart) {
	newM := Chart{}
	for c, s := range m {
		switch s {
		case ".":
			// An open acre will become filled with trees if three or more adjacent acres contained trees. Otherwise, nothing happens.
			if count(m, c, "|") >= 3 {
				newM[c] = "|"
			} else {
				newM[c] = "."
			}
		case "|":
			// An acre filled with trees will become a lumberyard if three or more adjacent acres were lumberyards. Otherwise, nothing happens.
			if count(m, c, "#") >= 3 {
				newM[c] = "#"
			} else {
				newM[c] = "|"
			}
		case "#":
			// An acre containing a lumberyard will remain a lumberyard if it was adjacent to at least one other lumberyard and at least one acre containing trees. Otherwise, it becomes open.
			if count(m, c, "#") >= 1 && count(m, c, "|") >= 1 {
				newM[c] = "#"
			} else {
				newM[c] = "."
			}
		}
	}

	key := AsString(m)
	if v, ok := cache[key]; ok {
		// warp found
		return v, newM
	}

	cache[key] = i
	return 0, newM
}

func count(m Chart, c Coord, s string) int {
	result := 0
	for _, d := range ALL_AND_DIAG {
		if m[c.Move(d)] == s {
			result++
		}
	}
	return result
}
