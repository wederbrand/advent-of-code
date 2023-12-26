package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	part1 := 0
	part2 := 0
	for i := 109165; i <= 576723; i++ {
		if testPart1(i) {
			part1++
		}
		if testPart2(i) {
			part2++
		}
	}
	fmt.Println("part1: ", part1, "in", time.Since(start))
	fmt.Println("part2: ", part2, "in", time.Since(start))
}

func testPart1(in int) bool {
	s := strconv.Itoa(in)

	twoInARowFound := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			return false
		}

		if s[i] == s[i+1] {
			twoInARowFound++
		}
	}

	return twoInARowFound > 0
}

func testPart2(in int) bool {
	s := strconv.Itoa(in)

	// reject strings not incrementing
	var twoInARowFound [10]int
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			return false
		}

		if s[i] == s[i+1] {
			twoInARowFound[util.Atoi(string(s[i]))]++
		}
	}

	for _, cnt := range twoInARowFound {
		if cnt == 1 {
			return true
		}
	}
	return false
}
