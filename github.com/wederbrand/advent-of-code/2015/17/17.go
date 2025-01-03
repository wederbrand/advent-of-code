package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"math"
	"time"
)

func main() {
	start := time.Now()
	inFile := GetFileContents("2015/17/input.txt", "\n")

	containers := make([]int, len(inFile))
	for i, s := range inFile {
		containers[i] = Atoi(s)
	}

	part1, minUsed := doIt(containers, 150, 0, 0, math.MaxInt, math.MaxInt)
	fmt.Println("Part 1: ", part1, "in", time.Since(start))

	part2, _ := doIt(containers, 150, 0, 0, math.MaxInt, minUsed)
	fmt.Println("Part 2: ", part2, "in", time.Since(start))
}

func doIt(containers []int, target int, startIndex int, used int, minUsed int, countLimit int) (int, int) {
	ways := 0
	used++
	for i := startIndex; i < len(containers); i++ {
		c := containers[i]
		if c == target {
			if used <= countLimit {
				ways++
			}
			if used < minUsed {
				minUsed = used
			}
		} else if c < target {
			w, mU := doIt(containers, target-c, i+1, used, minUsed, countLimit)
			ways += w
			minUsed = mU
		}
	}

	return ways, minUsed
}
