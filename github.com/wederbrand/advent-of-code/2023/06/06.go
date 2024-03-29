package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2023/06/input.txt", "\n")

	times := util.MatchingNumbersAfterSplitOnAny(inFile[0], ":", " ")[1]
	dists := util.MatchingNumbersAfterSplitOnAny(inFile[1], ":", " ")[1]

	part1 := solve(times, dists)
	fmt.Println("part1: ", part1, "in", time.Since(start))

	times = util.MatchingNumbersAfterSplitOnAny(strings.ReplaceAll(inFile[0], " ", ""), ":", " ")[1]
	dists = util.MatchingNumbersAfterSplitOnAny(strings.ReplaceAll(inFile[1], " ", ""), ":", " ")[1]
	part2 := solve(times, dists)
	fmt.Println("part2: ", part2, "in", time.Since(start))
}

func solve(times []int, dists []int) int {
	result := 1
	for i := range times {
		t := times[i]
		d := dists[i]
		_, first := util.BinarySearch(0, t/2, func(i int) bool {
			return i*(t-i) > d
		})
		result *= t - 2*first + 1
	}
	return result
}
