package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"time"
)

func main() {
	part1()
	part2()
}

func part1() {
	start := time.Now()
	inFile := util.GetFileContents("2016/02/input.txt", "\n")

	m := chart.Chart{}
	m[chart.Coord{0, 0}] = "1"
	m[chart.Coord{1, 0}] = "2"
	m[chart.Coord{2, 0}] = "3"

	m[chart.Coord{0, 1}] = "4"
	m[chart.Coord{1, 1}] = "5"
	m[chart.Coord{2, 1}] = "6"

	m[chart.Coord{0, 2}] = "7"
	m[chart.Coord{1, 2}] = "8"
	m[chart.Coord{2, 2}] = "9"

	curr := chart.Coord{1, 1} // starting at 5

	code := ""
	for _, line := range inFile {
		for _, r := range line {
			switch r {
			case 'U':
				if m[curr.Move(chart.UP)] != "" {
					curr = curr.Move(chart.UP)
				}
			case 'D':
				if m[curr.Move(chart.DOWN)] != "" {
					curr = curr.Move(chart.DOWN)
				}
			case 'L':
				if m[curr.Move(chart.LEFT)] != "" {
					curr = curr.Move(chart.LEFT)
				}
			case 'R':
				if m[curr.Move(chart.RIGHT)] != "" {
					curr = curr.Move(chart.RIGHT)
				}
			}
		}
		code += m[curr]
	}

	fmt.Println("Part 1:", code, "in", time.Since(start))
}

func part2() {
	start := time.Now()
	inFile := util.GetFileContents("2016/02/input.txt", "\n")

	m := chart.Chart{}
	m[chart.Coord{2, 0}] = "1"

	m[chart.Coord{1, 1}] = "2"
	m[chart.Coord{2, 1}] = "3"
	m[chart.Coord{3, 1}] = "4"

	m[chart.Coord{0, 2}] = "5"
	m[chart.Coord{1, 2}] = "6"
	m[chart.Coord{2, 2}] = "7"
	m[chart.Coord{3, 2}] = "8"
	m[chart.Coord{4, 2}] = "9"

	m[chart.Coord{1, 3}] = "A"
	m[chart.Coord{2, 3}] = "B"
	m[chart.Coord{3, 3}] = "C"

	m[chart.Coord{2, 4}] = "D"

	curr := chart.Coord{2, 2} // starting at 5

	code := ""
	for _, line := range inFile {
		for _, r := range line {
			switch r {
			case 'U':
				if m[curr.Move(chart.UP)] != "" {
					curr = curr.Move(chart.UP)
				}
			case 'D':
				if m[curr.Move(chart.DOWN)] != "" {
					curr = curr.Move(chart.DOWN)
				}
			case 'L':
				if m[curr.Move(chart.LEFT)] != "" {
					curr = curr.Move(chart.LEFT)
				}
			case 'R':
				if m[curr.Move(chart.RIGHT)] != "" {
					curr = curr.Move(chart.RIGHT)
				}
			}
		}
		code += m[curr]
	}

	fmt.Println("Part 2:", code, "in", time.Since(start))
}
