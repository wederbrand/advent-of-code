package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"time"
)

func main() {
	start := time.Now()
	inFile := GetFileContents("2024/10/input.txt", "\n")

	m := MakeChart(inFile, "")

	p1 := 0
	p2 := 0

	for c, s := range m {
		if s == "0" {
			nines := make(map[Coord]bool)
			paths := findPaths(m, c, nines, 0)
			p1 += len(nines)
			p2 += paths
		}
	}

	fmt.Println("Part 1:", p1, "in", time.Since(start))
	fmt.Println("Part 2:", p2, "in", time.Since(start))
}

func findPaths(m Chart, c Coord, nines map[Coord]bool, h int) int {
	if h == 9 {
		nines[c] = true
		return 1
	}

	// walk in all directions that are one higher
	// return the found nines and paths that lead there
	paths := 0
	for _, dir := range ALL {
		next := c.Move(dir)
		if Atoi(m[next]) == h+1 {
			paths += findPaths(m, next, nines, h+1)
		}
	}
	return paths
}
