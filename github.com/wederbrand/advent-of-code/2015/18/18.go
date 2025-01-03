package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"time"
)

func main() {
	start := time.Now()
	inFile := GetFileContents("2015/18/input.txt", "\n")

	m := MakeChart(inFile, ".")
	part1 := doIt(m, false)
	fmt.Println("Part 1: ", part1, "in", time.Since(start))

	m = MakeChart(inFile, ".")
	m[Coord{0, 0}] = "#"
	m[Coord{0, 99}] = "#"
	m[Coord{99, 0}] = "#"
	m[Coord{99, 99}] = "#"

	part2 := doIt(m, true)
	fmt.Println("Part 2: ", part2, "in", time.Since(start))
}

func doIt(m Chart, stuckCorners bool) int {
	for i := 0; i < 100; i++ {
		newM := Chart{}
		for x := 0; x < 100; x++ {
			for y := 0; y < 100; y++ {
				if stuckCorners && (x == 0 || x == 99) && (y == 0 || y == 99) {
					newM[Coord{x, y}] = "#"
					continue
				}
				c := Coord{x, y}
				n := countNeighbours(c, m)
				if m[c] == "#" {
					// on
					if n == 2 || n == 3 {
						newM[c] = "#"
					}
				} else {
					// off
					if n == 3 {
						newM[c] = "#"
					}
				}
			}
		}
		m = newM
	}

	part1 := len(m)
	return part1
}

func countNeighbours(c Coord, m Chart) interface{} {
	count := 0
	for _, dir := range ALL_AND_DIAG {
		if m[c.Move(dir)] == "#" {
			count++
		}
	}
	return count
}
