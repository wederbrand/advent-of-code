package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"time"
)

func main() {
	startTimer := time.Now()
	inFile := util.GetFileContents("2023/14/input.txt", "\n")

	m := util.MakeMap(inFile)

	rollNorth(m)
	part1 := getWeight(m)
	fmt.Println("part1: ", part1, "in", time.Since(startTimer))

	seen := make(map[string]int)
	for i := 0; i < 1000000000; i++ {
		for j := 0; j < 4; j++ {
			// tilt and rotate 4 times
			rollNorth(m)
			m = util.RotateClockWise(m)
		}

		s := util.MapAsString(m)
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

func getWeight(m map[string]string) int {
	_, _, _, maxY := util.GetMapMaxes(m)
	weight := 0
	for key, rock := range m {
		_, y := util.DeKey(key)
		if rock == "O" {
			weight += maxY - y + 1
		}
	}
	return weight
}

func rollNorth(m map[string]string) {
	minX, minY, maxX, maxY := util.GetMapMaxes(m)
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			// roll forward until blocked
			key := util.IntKey(x, y)
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
			for newY = y; newY > minY; newY-- {
				_, filled := m[util.IntKey(x, newY-1)]
				if filled {
					break
				}
			}
			newKey := util.IntKey(x, newY)
			m[newKey] = rock
		}
	}
}
