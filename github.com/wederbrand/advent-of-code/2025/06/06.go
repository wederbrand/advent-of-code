package main

import (
	"fmt"
	"strings"
	"time"

	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
)

func main() {
	start := time.Now()
	input := GetFileContents("2025/06/input.txt", "\n")

	operations := input[len(input)-1]
	numbers := input[:len(input)-1]

	part1 := 0
	part2 := 0

	first := 0
	lastOp := '+'
	for i, r := range operations {
		if r == '+' || r == '*' {
			if i > 0 {
				last := i - 1
				// calculate
				r1, r2 := doIt(first, last, numbers, lastOp)
				part1 += r1
				part2 += r2
			}
			first = i
			lastOp = r
		}
	}

	// final calculation
	maxLen := 0
	for _, row := range numbers {
		if len(row) > maxLen {
			maxLen = len(row)
		}
	}

	// calculate
	r1, r2 := doIt(first, maxLen, numbers, lastOp)
	part1 += r1
	part2 += r2

	fmt.Println("Part 1:", part1, "in", time.Since(start))
	fmt.Println("Part 2:", part2, "in", time.Since(start))
}

func doIt(first int, last int, numbers []string, r rune) (int, int) {
	r1 := 0
	r2 := 0
	if r == '*' {
		r1 = 1
		r2 = 1
	}

	transposed := make([]string, len(numbers))
	for _, row := range numbers {
		maxLen := min(len(row), last)
		cut := row[first:maxLen]
		atoi := Atoi(strings.TrimSpace(cut))
		if r == '+' {
			r1 += atoi
		} else {
			r1 *= atoi
		}

		for j, char := range cut {
			if char != ' ' {
				transposed[j] += string(char)
			}
		}
	}

	for _, v := range transposed {
		trimmed := strings.TrimSpace(v)
		if len(trimmed) == 0 {
			continue
		}
		atoi := Atoi(trimmed)
		if r == '+' {
			r2 += atoi
		} else {
			r2 *= atoi
		}
	}

	return r1, r2
}
