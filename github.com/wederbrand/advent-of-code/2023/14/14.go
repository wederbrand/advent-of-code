package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"time"
)

func main() {
	startTimer := time.Now()
	inFile := util.GetFileContents("2023/14/input.txt", "\n")

	m := MakeChart(inFile)

	rollNorth(m)
	part1 := getWeight(m)
	fmt.Println("part1: ", part1, "in", time.Since(startTimer))

	seen := make(map[string]int)
	for i := 0; i < 1000000000; i++ {
		for j := 0; j < 4; j++ {
			// tilt and rotate 4 times
			rollNorth(m)
			m = RotateClockWise(m)
		}

		s := ChartAsString(m)
		lastIndex, found := seen[s]
		if found {
			cycle := i - lastIndex
			for ; i < (1000000000 - cycle); i += cycle {
				// do nothing
			}
		} else {
			seen[s] = i
		}
	}

	part2 := getWeight(m)
	fmt.Println("part2: ", part2, "in", time.Since(startTimer))
}

func getWeight(m Chart) int {
	_, maxC := GetChartMaxes(m)
	weight := 0
	for key, rock := range m {
		if rock == "O" {
			weight += maxC.Y - key.Y + 1
		}
	}
	return weight
}

func rollNorth(m Chart) {
	minC, maxC := GetChartMaxes(m)
	for y := minC.Y; y <= maxC.Y; y++ {
		for x := minC.X; x <= maxC.X; x++ {
			// roll forward until blocked
			key := Coord{x, y}
			rock, found := m[key]
			if !found {
				continue
			}

			if rock == "#" {
				continue
			}

			// remove it from the map
			delete(m, key)

			// roll to lower y
			newY := y
			for newY = y; newY > minC.Y; newY-- {
				_, filled := m[Coord{x, newY - 1}]
				if filled {
					break
				}
			}
			newKey := Coord{x, newY}
			m[newKey] = rock
		}
	}
}
