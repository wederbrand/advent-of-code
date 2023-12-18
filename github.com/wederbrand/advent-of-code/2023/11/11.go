package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"time"
)

func main() {
	startTimer := time.Now()
	inFile := util.GetFileContents("2023/11/input.txt", "\n")

	// assume no stars
	xs := make(map[int]bool)
	ys := make(map[int]bool)
	for y, s := range inFile {
		for x, r := range s {
			if r == '#' {
				xs[x] = true
				ys[y] = true
			}
		}
	}

	m := make(map[string]bool)
	yOffset := 0
	for y, s := range inFile {
		expansion := 1000000
		if !ys[y] {
			yOffset += expansion - 1
		}
		xOffset := 0
		for x, r := range s {
			if !xs[x] {
				xOffset += expansion - 1
			}
			if r == '#' {
				m[util.IntKey(x+xOffset, y+yOffset)] = true
			}
		}
	}

	fmt.Println("parsing:", time.Since(startTimer))
	startTimer = time.Now()

	part2 := 0
	for a := range m {
		for b := range m {
			ax, ay := util.DeKey(a)
			bx, by := util.DeKey(b)
			dist := chart.Manhattan(chart.Coord{ax, ay}, chart.Coord{bx, by})
			part2 += dist
		}
	}

	part2 /= 2
	fmt.Println("part2: ", part2, "in", time.Since(startTimer))
}
