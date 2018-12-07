package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Coord struct {
	x int
	y int
}

var world [500][500]Coord
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

func main() {
	file, err := os.Open("2018_6.input");
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	lineMatcher := regexp.MustCompile(`^(\d+), (\d+)$`)

	for scanner.Scan() {
		lineMatch := lineMatcher.FindStringSubmatch(scanner.Text())
		x, _ := strconv.Atoi(lineMatch[1])
		y, _ := strconv.Atoi(lineMatch[2])
		coords = append(coords, Coord{x, y})
	}

	for x := range world {
		for y := range world[x] {
			world[x][y] = closest(x, y)
		}
	}

	// mark the ones on the edges as useless
	inf := make(map[Coord]bool, 0)

	for x := range world {
		for y := range world[x] {
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
						size++;
					}
				}
			}

			if size > maxSize {
				maxSize = size
			}
		}
	}

	fmt.Println(maxSize)
}
