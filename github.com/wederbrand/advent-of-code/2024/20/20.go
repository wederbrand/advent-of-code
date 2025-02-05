package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"time"
)

func main() {
	start := time.Now()
	inFile := GetFileContents("2024/20/input.txt", "\n")

	m := MakeChart(inFile, "")

	startCoord, _ := m.FindLetter("S")
	endCoord, _ := m.FindLetter("E")

	longest := m.GetPath(startCoord, endCoord)

	saves := findSavedPathLengths(longest, 2)
	part1 := 0
	for _, save := range saves {
		if save >= 100 {
			part1++
		}
	}
	fmt.Println("Part 1:", part1, "in", time.Since(start))

	saves = findSavedPathLengths(longest, 20)
	part2 := 0
	for _, save := range saves {
		if save >= 100 {
			part2++
		}
	}
	fmt.Println("Part 2:", part2, "in", time.Since(start))
}

/**
 * Find all cheats.
 */
func findSavedPathLengths(longest []Coord, cheatDistance int) []int {
	saves := make([]int, 0)

	for i, startCoord := range longest {
		for j := i + 1; j < len(longest); j++ {
			targetCoord := longest[j]
			manhattan := Manhattan(startCoord, targetCoord)
			normalPathDist := j - i
			if manhattan < normalPathDist && manhattan <= cheatDistance {
				saves = append(saves, normalPathDist-manhattan)
			}
		}
	}

	return saves
}
