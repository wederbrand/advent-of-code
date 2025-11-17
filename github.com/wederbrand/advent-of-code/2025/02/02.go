package main

import (
	"fmt"
	"strconv"
	"time"

	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
)

func main() {
	start := time.Now()
	input := GetFileContents("2025/02/input.txt", ",")

	part1 := 0
	part2 := 0
	for _, s := range input {
		var a, b int
		fmt.Sscanf(s, "%d-%d", &a, &b)

		for test := a; test <= b; test++ {
			candidate := strconv.Itoa(test)
			if p1(candidate) {
				part1 += test
			}
			if p2(candidate) {
				part2 += test
			}
		}
	}

	fmt.Println("Part 1:", part1, "in", time.Since(start))
	fmt.Println("Part 2:", part2, "in", time.Since(start))
}

func p1(candidate string) bool {
	mid := len(candidate) / 2
	if mid*2 != len(candidate) {
		// odd length, can't be split evenly
		return false
	}
	if testIt(candidate, mid) {
		return true
	}
	return false
}

func p2(candidate string) bool {
	mid := len(candidate) / 2

	for i := 1; i <= mid; i++ {
		if testIt(candidate, i) {
			return true
		}
	}

	// no matches found
	return false
}

func testIt(input string, subLen int) bool {
	// input with subLen length substrings
	if subLen == 0 || len(input)%subLen != 0 {
		// not evenly divisible
		return false
	}

	left := input[:subLen]
	for i := subLen; i < len(input); i += subLen {
		right := input[i : i+subLen]
		if left != right {
			return false
		}
	}

	// nothing broke, we have a match
	return true
}
