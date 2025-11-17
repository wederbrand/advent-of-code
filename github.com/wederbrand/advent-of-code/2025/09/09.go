package main

import (
	"fmt"
	"sort"
	"time"

	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/priorityqueue"
	"golang.org/x/exp/maps"
)

func main() {
	start := time.Now()
	input := GetFileContents("2025/09/input.txt", "\n")
	input = append(input, input[0]) // close the loop

	bigCoords := make([]Coord, 0)

	for _, s := range input {
		c := Coord{}
		fmt.Sscanf(s, "%d,%d", &c.X, &c.Y)
		bigCoords = append(bigCoords, c)
	}

	translateBigToSmall := make(map[Coord]Coord)
	translateSmallToBig := make(map[Coord]Coord)

	xs := make(map[int]bool)
	ys := make(map[int]bool)
	for _, coord := range bigCoords {
		xs[coord.X] = true
		ys[coord.Y] = true
	}

	xKeys := maps.Keys(xs)
	yKeys := maps.Keys(ys)

	sort.Ints(xKeys)
	sort.Ints(yKeys)

	xTranslateBigToSmall := make(map[int]int)
	yTranslateBigToSmall := make(map[int]int)

	for newX, x := range xKeys {
		xTranslateBigToSmall[x] = newX + 1
	}

	for newY, y := range yKeys {
		yTranslateBigToSmall[y] = newY + 1
	}

	for _, coord := range bigCoords {
		newCoord := Coord{X: xTranslateBigToSmall[coord.X], Y: yTranslateBigToSmall[coord.Y]}
		translateBigToSmall[coord] = newCoord
		translateSmallToBig[newCoord] = coord
	}

	m := Chart{}
	lastSmallCoord := Coord{}
	smallCoords := make([]Coord, 0)
	for _, coord := range bigCoords {
		smallCoord := translateBigToSmall[coord]
		smallCoords = append(smallCoords, smallCoord)
		if lastSmallCoord != (Coord{}) {
			for x := min(lastSmallCoord.X, smallCoord.X); x <= max(lastSmallCoord.X, smallCoord.X); x++ {
				for y := min(lastSmallCoord.Y, smallCoord.Y); y <= max(lastSmallCoord.Y, smallCoord.Y); y++ {
					m[Coord{X: x, Y: y}] = "#"
				}
			}
		}
		lastSmallCoord = smallCoord
	}

	minC, maxC := GetChartMaxes(m)

	// "flood fill" the outside
	q := NewQueue()
	initial := State{
		Data:     Coord{X: 0, Y: 0},
		Priority: 0,
	}
	q.Add(&initial)
	i := 0
	for q.HasNext() {
		qs := q.Next()
		c := qs.Data.(Coord)
		if _, exists := m[c]; exists {
			continue
		}
		if c.X < minC.X-1 || c.X > maxC.X+1 || c.Y < minC.Y-1 || c.Y > maxC.Y+1 {
			continue
		}
		i++
		m[c] = "O" // outside
		for _, dir := range ALL {
			q.Add(&State{
				c.Move(dir),
				0,
			})
		}
	}

	part1Area := 0
	part2Area := 0
	for _, smallA := range smallCoords {
		for _, smallB := range smallCoords {
			if smallA == smallB {
				continue
			}
			bigA := translateSmallToBig[smallA]
			bigB := translateSmallToBig[smallB]
			area := (IntAbs(bigA.X-bigB.X) + 1) * (IntAbs(bigA.Y-bigB.Y) + 1)
			if area > part1Area {
				part1Area = area
			}
			if area > part2Area {
				// check if box is inside
				if isBoxInside(smallA, smallB, m) {
					part2Area = area
				}
			}
		}
	}

	fmt.Println("Part 1:", part1Area, "in", time.Since(start))
	fmt.Println("Part 2:", part2Area, "in", time.Since(start))
}

func isBoxInside(smallA Coord, smallB Coord, m Chart) bool {
	minX := min(smallA.X, smallB.X)
	maxX := max(smallA.X, smallB.X)
	minY := min(smallA.Y, smallB.Y)
	maxY := max(smallA.Y, smallB.Y)
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			if m[Coord{X: x, Y: y}] == "O" {
				// outside found, not valid
				return false
			}
		}
	}

	return true
}
