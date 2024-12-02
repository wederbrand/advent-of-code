package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2024/02/input.txt", "\n")

	part1 := part1(inFile)
	part2 := part2(inFile)

	fmt.Println("Part 1:", part1, "in", time.Since(start))
	fmt.Println("Part 2:", part2, "in", time.Since(start))
}

func part1(inFile []string) int {
	safe := 0
	for _, line := range inFile {
		split := strings.Split(line, " ")
		if checkIt(split) {
			safe++
		}
	}
	return safe
}

func part2(inFile []string) int {
	safe := 0
	for _, line := range inFile {
		split := strings.Split(line, " ")
		if checkIt(split) {
			safe++
			continue
		}

		for i := range split {
			newSplit := util.CloneSliceDelete(split, func(j int, s string) bool {
				return i == j
			})

			if checkIt(newSplit) {
				safe++
				break
			}
		}
	}
	return safe
}

func checkIt(split []string) bool {
	cntNeg := 0
	cntPos := 0

	for i := 1; i < len(split); i++ {
		a := util.Atoi(split[i-1])
		b := util.Atoi(split[i])

		diff := a - b

		if util.IntAbs(diff) < 1 || util.IntAbs(diff) > 3 {
			return false
		}

		if diff < 0 {
			cntNeg++
		} else {
			cntPos++
		}
	}

	if cntNeg == 0 || cntPos == 0 {
		return true
	}

	return false
}
