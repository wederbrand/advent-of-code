package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"time"
)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2024/06/input.txt", "\n")
	m := MakeChart(inFile, ".")

	_, startC := Find(m, "^")
	minC, maxC := GetChartMaxes(m)

	_, seen := doIt(m, minC, maxC, startC, UP)
	fmt.Println("Part 1:", len(seen), "in", time.Since(start))

	p2 := 0
	for c, _ := range seen {
		if m[c] != "#" {
			m[c] = "#"
			loop, _ := doIt(m, minC, maxC, startC, UP)
			if loop {
				p2++
			}
			// reset
			delete(m, c)
		}
	}

	fmt.Println("Part 2:", p2, "in", time.Since(start))
}

type CoordDir struct {
	c Coord
	d Dir
}

func doIt(m Chart, minC Coord, maxC Coord, c Coord, d Dir) (bool, map[Coord]bool) {
	// walk until out of bounds
	loop := make(map[CoordDir]bool)
	seen := make(map[Coord]bool)
	for c.X <= maxC.X && c.Y <= maxC.Y && c.X >= minC.X && c.Y >= minC.Y {
		_, found := loop[CoordDir{c, d}]
		if found {
			return true, nil
		}
		seen[c] = true
		loop[CoordDir{c, d}] = true
		if m[c.Move(d)] == "#" {
			d = d.Right()
		} else {
			c = c.Move(d)
		}
	}

	return false, seen
}
