package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"math"
	"strings"
	"time"
)

type Range struct {
	from int
	to   int
}

func (r Range) blockAndSplit(from int, to int) []Range {
	// if outside just return the original one
	if from > r.to || to < r.from {
		return []Range{r}
	}

	// if covered the inputted one splits this one into two bordering
	if from >= r.from && to <= r.to {
		return []Range{{r.from, from - 1}, Range{to + 1, r.to}}
	}

	// if covering the inputted one completely wipes this one
	if from <= r.from && to >= r.to {
		return []Range{}
	}

	// if from is inside a smaller range
	if from >= r.from && from <= r.to {
		return []Range{{r.from, from - 1}}
	}

	// if to is inside a smaller range
	if to >= r.from && to <= r.to {
		return []Range{{to + 1, r.to}}
	}

	panic("Should not happen")
}

func main() {
	start := time.Now()
	inFile := GetFileContents("2016/20/input.txt", "\n")

	// the world is open
	allowedRanges := make([]Range, 0)
	allowedRanges = append(allowedRanges, Range{0, 4294967295})

	for _, line := range inFile {
		// find ranges that this new one splits
		// don't care about merging, just split
		split := strings.Split(line, "-")
		from := Atoi(split[0])
		to := Atoi(split[1])

		newRanges := make([]Range, 0)
		for _, r := range allowedRanges {
			splitRanges := r.blockAndSplit(from, to)
			for _, splitRange := range splitRanges {
				if splitRange.from <= splitRange.to {
					newRanges = append(newRanges, splitRange)
				}
			}
		}
		allowedRanges = newRanges
	}

	part1 := math.MaxInt
	part2 := 0
	for _, allowedRange := range allowedRanges {
		if allowedRange.from < part1 {
			part1 = allowedRange.from
		}
		part2 += allowedRange.to - allowedRange.from + 1
	}

	fmt.Println("Part 1: ", part1, "in", time.Since(start))
	fmt.Println("Part 2: ", part2, "in", time.Since(start))
}
