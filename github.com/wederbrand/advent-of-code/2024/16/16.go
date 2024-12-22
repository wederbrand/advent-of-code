package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"math"
	"time"
)

type state struct {
	// I can get here at this price
	c     Coord
	d     Dir
	score int
	path  []Coord
}

type seenKey struct {
	c Coord
	d Dir
}

func main() {
	start := time.Now()
	inFile := GetFileContents("2024/16/input.txt", "\n")

	m := MakeChart(inFile, ".")

	s := m.FindLetter("S")
	e := m.FindLetter("E")

	p1, p2 := doIt(s, e, m)

	fmt.Println("Part 1:", p1, "in", time.Since(start))
	fmt.Println("Part 2:", p2, "in", time.Since(start))
}

func doIt(s Coord, e Coord, m Chart) (int, int) {
	ch := make(chan []Coord)

	paths := make(Chart)
	paths[s] = "O"
	paths[e] = "O"

	go m.GetAllPaths(ch, s, e, true)

	minFound := math.MaxInt
	for path := range ch {
		score := 0

		// GetAllPaths returns the start coord as the first element
		path = path[1:]

		curr := s
		lastDir := E
		for _, coord := range path {
			score++
			if curr.DirectionTo(coord) != lastDir {
				// we turned
				score += 1000
			}

			lastDir = curr.DirectionTo(coord)
			curr = coord

			paths[coord] = "O"
		}

		if score < minFound {
			minFound = score
		}
	}

	return minFound, len(paths)
}
