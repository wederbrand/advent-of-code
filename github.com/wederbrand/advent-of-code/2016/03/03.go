package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"time"
)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2016/03/input.txt", "\n")
	part1 := part1(inFile)
	fmt.Println("Part 1: ", part1, "in", time.Since(start))
	part2 := part2(inFile)
	fmt.Println("Part 2: ", part2, "in", time.Since(start))
}

func part1(inFile []string) int {
	valid := 0
	for _, line := range inFile {
		result := util.MatchingNumbersAfterSplitOnAny(line, " ", " ")
		a := result[0][0]
		b := result[1][0]
		c := result[2][0]

		if a+b > c && a+c > b && b+c > a {
			valid++
		}
	}

	return valid
}

func part2(inFile []string) int {
	valid := 0
	for i := 0; i < len(inFile); i += 3 {
		for j := 0; j < 3; j++ {
			aResult := util.MatchingNumbersAfterSplitOnAny(inFile[i+0], " ", " ")
			bResult := util.MatchingNumbersAfterSplitOnAny(inFile[i+1], " ", " ")
			cResult := util.MatchingNumbersAfterSplitOnAny(inFile[i+2], " ", " ")

			a := aResult[j][0]
			b := bResult[j][0]
			c := cResult[j][0]

			if a+b > c && a+c > b && b+c > a {
				valid++
			}
		}
	}

	return valid
}
