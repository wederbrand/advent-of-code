package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"slices"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2024/02/input.txt", "\n")

	part1 := part1(inFile)
	part2 := part2(inFile)

	fmt.Println("Part 1:", part1, "in", time.Since(start))
	fmt.Println("Part 2:", part2, "in", time.Since(start))
}

func part1(inFile []string) int {
	safe := 0
	for _, line := range inFile {
		split := strings.Split(line, " ")
		if checkIt(split) {
			safe++
		}
	}
	return safe
}

func part2(inFile []string) int {
	safe := 0
	for _, line := range inFile {
		split := strings.Split(line, " ")
		if checkIt(split) {
			safe++
			continue
		}

		for i := range split {
			newSplit := util.CloneSliceDelete(split, func(j int, s string) bool {
				return i == j
			})

			if checkIt(newSplit) {
				safe++
				break
			}
		}
	}
	return safe
}

func checkIt(split []string) bool {
	values := make([]int, 0)
	sorted := make([]int, 0)
	for _, s := range split {
		value := util.Atoi(s)
		values = append(values, value)
		sorted = append(sorted, value)
	}

	slices.Sort(sorted)

	for i := 1; i < len(sorted); i++ {
		diff := util.IntAbs(sorted[i] - sorted[i-1])
		if diff < 1 || diff > 3 {
			return false
		}
	}

	if slices.Equal(sorted, values) {
		return true
	}

	slices.Reverse(sorted)

	if slices.Equal(sorted, values) {
		return true
	}

	return false
}
