package main

import (
	"fmt"
	"time"

	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
)

func main() {
	start := time.Now()
	input := GetFileContents("2025/04/input.txt", "\n")
	m := MakeChart(input, ".")

	free := liftable(m)
	part1 := len(free)
	fmt.Println("Part 1:", part1, "in", time.Since(start))

	part2 := 0
	queue := append([]Coord(nil), free...)
	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]

		if _, exists := m[c]; !exists {
			continue
		}

		a := getAdjacent(m, c)

		if len(a) < 4 {
			delete(m, c)
			part2++
			queue = append(queue, a...)
		}
	}

	fmt.Println("Part 2:", part2, "in", time.Since(start))
}

func liftable(m Chart) []Coord {
	var result []Coord
	for c := range m {
		a := getAdjacent(m, c)
		if len(a) < 4 {
			result = append(result, c)
		}
	}

	return result
}

func getAdjacent(m Chart, c Coord) []Coord {
	var adjacent []Coord
	for _, d := range ALL_AND_DIAG {
		_, found := m[c.Move(d)]
		if found {
			adjacent = append(adjacent, c.Move(d))
		}
	}
	return adjacent
}
