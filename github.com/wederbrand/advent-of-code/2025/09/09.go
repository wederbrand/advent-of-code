package main

import (
	"fmt"
	"time"

	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
)

func main() {
	start := time.Now()
	input := GetFileContents("2025/09/input.txt", "\n")
	input = append(input, input[0]) // close the loop

	coords := make([]Coord, 0)

	for _, s := range input {
		c := Coord{}
		fmt.Sscanf(s, "%d,%d", &c.X, &c.Y)
		coords = append(coords, c)
	}

	vertices := make([][2]Coord, 0)
	lastCoord := Coord{}
	for _, coord := range coords {
		if lastCoord == (Coord{}) {
			lastCoord = coord
			continue
		}

		minX := min(lastCoord.X, coord.X)
		maxX := max(lastCoord.X, coord.X)
		minY := min(lastCoord.Y, coord.Y)
		maxY := max(lastCoord.Y, coord.Y)

		vertices = append(vertices, [2]Coord{{X: minX, Y: minY}, {X: maxX, Y: maxY}})
		lastCoord = coord
	}

	part1Area := 0
	part2Area := 0
	for _, a := range coords {
		for _, b := range coords {
			if a == b {
				continue
			}
			area := (IntAbs(a.X-b.X) + 1) * (IntAbs(a.Y-b.Y) + 1)
			if area > part1Area {
				part1Area = area
			}
			if area > part2Area {
				// check if box is inside
				if !isBoxIntersected(a, b, vertices) {
					part2Area = area
				}
			}
		}
	}

	fmt.Println("Part 1:", part1Area, "in", time.Since(start))
	fmt.Println("Part 2:", part2Area, "in", time.Since(start))
}

func isBoxIntersected(a Coord, b Coord, vertices [][2]Coord) bool {
	inMinX := min(a.X, b.X)
	inMaxX := max(a.X, b.X)
	inMinY := min(a.Y, b.Y)
	inMaxY := max(a.Y, b.Y)

	for _, vertex := range vertices {
		// vertices are already pointing down/right
		vMinX := vertex[0].X
		vMaxX := vertex[1].X
		vMinY := vertex[0].Y
		vMaxY := vertex[1].Y

		if inMinX < vMaxX && inMaxX > vMinX && inMinY < vMaxY && inMaxY > vMinY {
			// vertex intersects the box
			return true
		}
	}
	return false
}
