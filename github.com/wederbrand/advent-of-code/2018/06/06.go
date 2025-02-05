package main

import (
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"time"

	"fmt"
	"math"
	"regexp"
	"strconv"
)

type Coord struct {
	x int
	y int
}

var world [500][500]Coord
var world2 [500][500]bool

const limit = 10000

var coords = make([]Coord, 0)

var COLLISION = Coord{-1, -1}

func closest(x, y int) Coord {
	min := float64(2 * len(world))
	var closest Coord
	for _, coord := range coords {
		dist := math.Abs(float64(coord.x-x)) + math.Abs(float64(coord.y-y))
		if dist < min {
			min = dist
			closest = coord
		} else if dist == min {
			closest = COLLISION
		}
	}

	return closest
}

func closeEnough(x, y int) bool {
	dist := 0
	for _, coord := range coords {
		dist += int(math.Abs(float64(coord.x-x)) + math.Abs(float64(coord.y-y)))
	}

	if dist < limit {
		return true
	} else {
		return false
	}
}

func main() {
	start := time.Now()
	inFile := GetFileContents("2018/06/input.txt", "\n")

	lineMatcher := regexp.MustCompile(`^(\d+), (\d+)$`)

	for _, s := range inFile {
		lineMatch := lineMatcher.FindStringSubmatch(s)
		x, _ := strconv.Atoi(lineMatch[1])
		y, _ := strconv.Atoi(lineMatch[2])
		coords = append(coords, Coord{x, y})
	}

	for x := range world {
		for y := range world[x] {
			world[x][y] = closest(x, y)
			world2[x][y] = closeEnough(x, y)
		}
	}

	// mark the ones on the edges as useless
	inf := make(map[Coord]bool, 0)

	totalSize := 0
	for x := range world {
		for y := range world[x] {
			if world2[x][y] {
				totalSize++
			}
			if x == 0 || x == len(world)-1 || y == 0 || y == len(world[x])-1 {
				inf[world[x][y]] = true
			}
		}
	}

	maxSize := 0
	for _, coord := range coords {
		if !inf[coord] {
			// calculate size
			size := 0
			for x := range world {
				for y := range world[x] {
					if world[x][y] == coord {
						size++
					}
				}
			}

			if size > maxSize {
				maxSize = size
			}
		}
	}

	fmt.Println("Part 1:", maxSize, "in", time.Since(start))
	fmt.Println("Part 2:", totalSize, "in", time.Since(start))
}
