package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"time"
)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2024/04/input.txt", "\n")
	m := MakeChart(inFile, "")

	p1, p2 := doIt(m)

	fmt.Println("Part 1:", p1, "in", time.Since(start))
	fmt.Println("Part 2:", p2, "in", time.Since(start))
}

func doIt(m Chart) (int, int) {
	p1 := 0
	p2 := 0

	minC, maxC := GetChartMaxes(m)
	for y := minC.Y; y <= maxC.Y; y++ {
		for x := minC.X; x <= maxC.X; x++ {
			c := Coord{x, y}
			if m[c] == "X" {
				for _, d := range ALL_AND_DIAG {
					if isItFollowedByMAS(m, c, d) {
						p1++
					}
				}
			}
			if m[c] == "A" {
				if doesItHaveThoseMsAndSes(m, c) {
					p2++
				}
			}
		}
	}

	return p1, p2
}

func isItFollowedByMAS(m Chart, c Coord, d Dir) bool {
	if m[c.Move(d, 1)] != "M" {
		return false
	}

	if m[c.Move(d, 2)] != "A" {
		return false
	}

	if m[c.Move(d, 3)] != "S" {
		return false
	}

	return true
}

func doesItHaveThoseMsAndSes(m Chart, c Coord) bool {
	upLeft := m[c.Move(UPLEFT)]
	downRight := m[c.Move(DOWNRIGHT)]
	oneWay := (upLeft == "M" && downRight == "S") || (upLeft == "S" && downRight == "M")

	upRight := m[c.Move(UPRIGHT)]
	downLeft := m[c.Move(DOWNLEFT)]
	theOtherWay := (upRight == "M" && downLeft == "S") || (upRight == "S" && downLeft == "M")

	return oneWay && theOtherWay
}
