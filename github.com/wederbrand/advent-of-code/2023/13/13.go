package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"math"
	"time"
)

func main() {
	startTimer := time.Now()
	inFile := util.GetFileContents("2023/13/input.txt", "\n")

	m := make(map[string]bool)
	y := -1
	part1 := 0
	part2 := 0
	for _, s := range inFile {
		if s == "" {
			// analyze last pattern and start a new
			part1 += analyze(m, 0)
			part2 += analyze(m, 1)
			m = make(map[string]bool)
			y = -1
			continue
		}

		y++
		for x, r := range s {
			if r == '#' {
				key := util.IntKey(x, y)
				m[key] = true
			}
		}
	}
	part1 += analyze(m, 0)
	part2 += analyze(m, 1)

	fmt.Println("part1: ", part1, "in", time.Since(startTimer))
	fmt.Println("part2: ", part2, "in", time.Since(startTimer))
}

func analyze(m map[string]bool, nbrErrors int) int {
	// find either vertical or horizontal line
	minX := math.MaxInt
	minY := math.MaxInt
	maxX := math.MinInt
	maxY := math.MinInt
	for k := range m {
		x, y := util.DeKey(k)
		minX = min(minX, x)
		minY = min(minY, y)
		maxX = max(maxX, x)
		maxY = max(maxY, y)
	}

	// test potential vertical mirrors
	// between x and x+
	for x := minX; x < maxX; x++ {
		if isXMirror(m, x, minX, minY, maxX, maxY, nbrErrors) {
			return x + 1
		}
	}

	// test potential horizontal mirrors
	// between y and y+
	for y := minY; y < maxY; y++ {
		if isYMirror(m, y, minX, minY, maxX, maxY, nbrErrors) {
			return 100 * (y + 1)
		}
	}

	panic("ho ho")
}

func isXMirror(m map[string]bool, mirror int, minX int, minY int, maxX int, maxY int, errors int) bool {
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
			_, a := m[util.IntKey(x, y)]
			_, b := m[util.IntKey(mirroredX, y)]
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

func isYMirror(m map[string]bool, mirror int, minX int, minY int, maxX int, maxY int, errors int) bool {
	// start with both sides of mirror
	// walk outwards
	mirroredY := mirror
	for y := mirror; y >= minY; y-- {
		mirroredY++
		if mirroredY > maxY {
			// outside, all is good
			return errors == 0
		}

		// test entire row
		for x := minX; x <= maxX; x++ {
			_, a := m[util.IntKey(x, y)]
			_, b := m[util.IntKey(x, mirroredY)]
			if a != b {
				errors--
				if errors < 0 {
					// not identical enough, this y can't work
					return false
				}
			}
		}
	}
	// outside, all is good
	return errors == 0
}
