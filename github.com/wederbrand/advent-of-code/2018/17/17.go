package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"time"
)

const EMPTY = ""

func main() {
	start := time.Now()
	inFile := GetFileContents("2018/17/input.txt", "\n")

	m := Chart{}
	for _, line := range inFile {
		var a, c string
		var b, d, e int
		fmt.Sscanf(line, "%1s=%d, %1s=%d..%d", &a, &b, &c, &d, &e)
		if a == "x" {
			for y := d; y <= e; y++ {
				m[Coord{b, y}] = "#"
			}
		} else {
			for x := d; x <= e; x++ {
				m[Coord{x, b}] = "#"
			}
		}
	}

	minC, maxC := GetChartMaxes(m)

	fill(m, maxC.Y, Coord{X: 500, Y: minC.Y})

	part1 := 0
	part2 := 0
	for _, v := range m {
		if v == "~" {
			part1++
			part2++
		}
		if v == "|" {
			part1++
		}
	}

	fmt.Println("Part 1:", part1, "in", time.Since(start))
	fmt.Println("Part 2:", part2, "in", time.Since(start))
}

func fill(m Chart, maxY int, coord Coord) {
	m[coord] = "|"

	if coord.Y == maxY {
		return
	}

	// if down is unknown, fill down first
	if m[coord.Move(DOWN)] == EMPTY {
		fill(m, maxY, coord.Move(DOWN))
	}

	switch m[coord.Move(DOWN)] {
	case "#", "~":
		m[coord] = "~"

		if m[coord.Move(LEFT)] == EMPTY {
			fill(m, maxY, coord.Move(LEFT))
			if m[coord.Move(LEFT)] == "|" {
				m[coord] = "|"
			}
		}
		if m[coord.Move(LEFT)] == "|" {
			m[coord] = "|"
		}
		if m[coord.Move(RIGHT)] == EMPTY {
			fill(m, maxY, coord.Move(RIGHT))
			if m[coord.Move(RIGHT)] == "|" {
				m[coord] = "|"
				// weird bonus: if we find | to the right we need to change the ~ to | to the left
				weirdX := coord.X - 1
				for m[Coord{weirdX, coord.Y}] == "~" {
					m[Coord{weirdX, coord.Y}] = "|"
					weirdX--
				}
			}
		}
	case "|":
		// this is also |
		m[coord] = "|"
	}
}
