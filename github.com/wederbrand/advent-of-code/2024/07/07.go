package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2024/07/input.txt", "\n")

	p1 := 0
	p2 := 0
	for _, line := range inFile {
		split := strings.Split(line, ":")

		test := util.Atoi(split[0])
		matches := util.MatchingNumbersAfterSplitOnAny(split[1], "", " ")[0]

		if canDo(test, 0, matches, false) {
			p1 += test
			p2 += test
		} else if canDo(test, 0, matches, true) {
			p2 += test
		}
	}

	fmt.Println("Part 1:", p1, "in", time.Since(start))
	fmt.Println("Part 2:", p2, "in", time.Since(start))
}

func canDo(test int, i int, matches []int, withCon bool) bool {
	if i > test {
		return false
	}

	if i == test && len(matches) == 0 {
		return true
	}

	if len(matches) == 0 {
		return false
	}

	return canDo(test, i+matches[0], matches[1:], withCon) ||
		canDo(test, i*matches[0], matches[1:], withCon) ||
		(withCon && canDo(test, concat(i, matches[0]), matches[1:], withCon))
}

func concat(a int, b int) int {
	switch {
	case b < 10:
		return a*10 + b
	case b < 100:
		return a*100 + b
	case b < 1000:
		return a*1000 + b
	}
	return util.Atoi(fmt.Sprintf("%d%d", a, b))
}
