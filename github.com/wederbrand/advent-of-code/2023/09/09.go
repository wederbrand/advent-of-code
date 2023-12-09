package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"time"
)

func main() {
	startTimer := time.Now()
	inFile := util.GetFileContents("2023/09/input.txt", "\n")

	var all [][]int
	part1 := 0
	part2 := 0
	for _, s := range inFile {
		in := util.MatchingNumbersAfterSplitOnAny(s, ":")[0]
		all = solve(in)
		all = append(all, in)

		last := 0
		first := 0
		for _, row := range all {
			last = row[len(row)-1] + last
			first = row[0] - first
		}
		part1 += last
		part2 += first
	}

	fmt.Println("part1: ", part1, "in", time.Since(startTimer))
	fmt.Println("part2: ", part2, "in", time.Since(startTimer))
}

func solve(in []int) [][]int {
	zeroes := true
	result := make([]int, 0)

	for i := 1; i < len(in); i++ {
		val := in[i] - in[i-1]
		result = append(result, val)
		if val != 0 {
			zeroes = false
		}
	}

	output := make([][]int, 0)
	if !zeroes {
		output = solve(result)
	}
	output = append(output, result)
	return output
}
