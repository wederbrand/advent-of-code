package main

import (
	"fmt"
	"slices"
	"strings"
	"time"

	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
)

func main() {
	start := time.Now()
	input := GetFileContents("2025/05/input.txt", "\n")

	ranges := [][]int{}
	ingredients := []int{}
	for _, s := range input {
		// range
		if strings.Contains(s, "-") {
			var a, b int
			fmt.Sscanf(s, "%d-%d", &a, &b)
			ranges = append(ranges, []int{a, b})
		} else if s != "" {
			// ingredient
			ingredients = append(ingredients, Atoi(s))
		}
	}

	part1 := 0
	for _, ingredient := range ingredients {
		spoiled := true
		for _, r := range ranges {
			if ingredient >= r[0] && ingredient <= r[1] {
				spoiled = false
				break
			}
		}
		if !spoiled {
			part1++
		}
	}
	fmt.Println("Part 1:", part1, "in", time.Since(start))

	// sort the ranges
	slices.SortFunc(ranges, func(a, b []int) int {
		return a[0] - b[0]
	})

	// merge overlapping ranges
	for i := range ranges {
		if i == 0 {
			continue
		}

		if ranges[i][0] <= ranges[i-1][1] {
			// merge into the second one
			ranges[i][0] = min(ranges[i-1][0], ranges[i][0])
			ranges[i][1] = max(ranges[i-1][1], ranges[i][1])
			// mark the first for deletion
			ranges[i-1] = []int{-1, -1}
		}
	}

	part2 := 0
	// count the remaining range
	for _, r := range ranges {
		if r[0] == -1 {
			continue
		}

		part2 += r[1] - r[0] + 1
	}

	fmt.Println("Part 2:", part2, "in", time.Since(start))
}
