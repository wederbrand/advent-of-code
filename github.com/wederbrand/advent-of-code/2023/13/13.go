package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"time"
)

func main() {
	startTimer := time.Now()
	inFile := util.GetFileContents("2023/13/input.txt", "\n")

	m := make(Chart)
	y := -1
	part1 := 0
	part2 := 0
	for _, s := range inFile {
		if s == "" {
			// analyze last pattern and start a new
			part1 += analyze(m, 0)
			part2 += analyze(m, 1)
			m = make(Chart)
			y = -1
			continue
		}

		y++
		for x, r := range s {
			if r == '#' {
				m[Coord{x, y}] = "#"
			}
		}
	}
	part1 += analyze(m, 0)
	part2 += analyze(m, 1)

	fmt.Println("part1: ", part1, "in", time.Since(startTimer))
	fmt.Println("part2: ", part2, "in", time.Since(startTimer))
}

func analyze(m Chart, nbrErrors int) int {
	// test potential vertical mirrors
	// between x and x+ (left to right)
	minC, maxC := GetChartMaxes(m)
	for x := minC.X; x < maxC.X; x++ {
		if isMirror(m, x, minC.X, minC.Y, maxC.X, maxC.Y, nbrErrors) {
			return x + 1
		}
	}

	// test potential horizontal mirrors by rotating the map counter clock wise
	// between y and y+ (left to right, former top to bottom)
	m = RotateCounterClockWise(m)
	minC, maxC = GetChartMaxes(m)
	for x := minC.X; x < maxC.X; x++ {
		if isMirror(m, x, minC.X, minC.Y, maxC.X, maxC.Y, nbrErrors) {
			return 100 * (x + 1)
		}
	}

	panic("ho ho")
}

func isMirror(m Chart, mirror int, minX int, minY int, maxX int, maxY int, errors int) bool {
	// start with both sides of mirror
	// walk outwards
	mirroredX := mirror
	for x := mirror; x >= minX; x-- {
		mirroredX++
		if mirroredX > maxX {
			// outside, all is good
			return errors == 0
		}

		// test entire column
		for y := minY; y <= maxY; y++ {
			_, a := m[Coord{x, y}]
			_, b := m[Coord{mirroredX, y}]
			if a != b {
				errors--
				if errors < 0 {
					// not identical enough, this x can't work
					return false
				}
			}
		}
	}
	// outside, all is good
	return errors == 0
}
