package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"time"
)

type Region struct {
	size    int
	border  int
	corners int
}

func main() {
	start := time.Now()
	inFile := GetFileContents("2024/12/input.txt", "\n")

	m := MakeChart(inFile, "")

	regions := make([]Region, 0)
	seen := make(Chart)
	for c, s := range m {
		if _, found := seen[c]; !found {
			size, border, corners := doIt(m, c, seen, s)
			regions = append(regions, Region{size, border, corners})
		}
	}

	p1 := 0
	p2 := 0
	for _, region := range regions {
		p1 += region.size * region.border
		p2 += region.size * region.corners
	}
	fmt.Println("Part 1:", p1, "in", time.Since(start))
	fmt.Println("Part 2:", p2, "in", time.Since(start))
}

func doIt(m Chart, c Coord, seen Chart, plant string) (int, int, int) {
	if plant != m[c] {
		return 0, 0, 0
	}

	if _, found := seen[c]; found {
		return 0, 0, 0
	}

	seen[c] = plant

	size := 1
	border := 0
	corners := 0

	for _, dir := range ALL {
		// it's a border if two neighboring cells are different
		if m[c.Move(dir)] != plant && m[c.Move(dir.Right())] != plant {
			corners++
		}

		// it's also a corner if two neighboring cells are equal but the diagonal is different
		if m[c.Move(dir)] == plant && m[c.Move(dir.Right())] == plant && m[c.Move(dir).Move(dir.Right())] != plant {
			corners++
		}

		if m[c.Move(dir)] != plant {
			border++
		} else {
			nextSize, nextBorder, nextCorners := doIt(m, c.Move(dir), seen, plant)
			size += nextSize
			border += nextBorder
			corners += nextCorners
		}
	}

	return size, border, corners
}
