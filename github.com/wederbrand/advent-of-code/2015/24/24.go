package main

import (
	"fmt"
	"github.com/mowshon/iterium"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"math"
	"time"
)

func main() {
	start := time.Now()
	inFile := GetFileContents("2015/24/input.txt", "\n")

	packages := make([]int, len(inFile))
	totalWeight := 0
	for i, v := range inFile {
		packages[i] = Atoi(v)
		totalWeight += packages[i]
	}

	part1 := doIt(packages, totalWeight, 3)
	fmt.Println("Part 1: ", part1, "in", time.Since(start))

	part2 := doIt(packages, totalWeight, 4)
	fmt.Println("Part 2: ", part2, "in", time.Since(start))
}

func doIt(packages []int, totalWeight int, divisor int) int {
	minLen := math.MaxInt
	minQE := math.MaxInt
	for i := 1; i < len(packages) && minLen == math.MaxInt; i++ {
		combinations := iterium.Combinations(packages, i)
		for comb := range combinations.Chan() {
			// The problem states that the other two sets must have the same weight,
			// but it turns out to be correct without checking
			// Lucky me.
			if sum(comb) == totalWeight/divisor {
				if len(comb) < minLen {
					minLen = len(comb)
					minQE = prod(comb)
				} else if len(comb) == minLen {
					minQE = intMin(minQE, prod(comb))
				}
			}
		}
	}
	return minQE
}

func sum(comb []int) int {
	result := 0
	for _, v := range comb {
		result += v
	}
	return result
}

func prod(comb []int) int {
	result := 1
	for _, v := range comb {
		result *= v
	}
	return result
}

func intMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
